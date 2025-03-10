package chain

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"math"
	"math/big"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

const (
	// default timeout for waiting for a tx to be mined.
	DefaultTxTimeout    = 60 * time.Second
	DefaultGasLimit     = 2_500_000
	DefaultTipPerGasCap = 20_000_000_000 //20 GWei

	baseFeeCapMultiplier = 4

	retryTipMultiplierTimes10 int64 = 12 // 1,2 * 10
)

var (
	DefaultTipCap               = big.NewInt(DefaultTipPerGasCap)  // 20 GWei
	DefaultBaseFeeCapMultiplier = big.NewInt(baseFeeCapMultiplier) // 4
)

type TxVerifier struct {
	eth *ethclient.Client
}

func NewTxVerifier(eth *ethclient.Client) *TxVerifier {
	return &TxVerifier{eth: eth}
}

func (t TxVerifier) WaitUntilMined(from common.Address, tx *types.Transaction, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, t.eth, tx)
	if err != nil {
		return errors.Wrap(err, "bind.WaitMined")
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		reason, err := errorReason(ctx, t.eth, from, tx, receipt.BlockNumber)
		if err != nil {
			return err
		}
		return errors.Errorf("tx failed: %s", reason)
	}
	return nil
}

// Taken from: https://ethereum.stackexchange.com/questions/48383/how-to-retrieve-revert-reason-for-past-transactions
func errorReason(ctx context.Context, b ethereum.ContractCaller, from common.Address, tx *types.Transaction, blockNum *big.Int) (string, error) {
	msg := ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}
	res, err := b.CallContract(ctx, msg, blockNum)
	if err != nil {
		return "", errors.Wrap(err, "CallContract")
	}
	return unpackError(res)
}

var (
	errorSig     = []byte{0x08, 0xc3, 0x79, 0xa0} // Keccak256("Error(string)")[:4]
	abiString, _ = abi.NewType("string", "", nil)
)

func unpackError(result []byte) (string, error) {
	if len(result) < 4 || !bytes.Equal(result[:4], errorSig) {
		return "<tx result not Error(string)>", errors.New("tx result not of type Error(string)")
	}
	vs, err := abi.Arguments{{Type: abiString}}.UnpackValues(result[4:])
	if err != nil {
		return "<invalid tx result>", errors.Wrap(err, "unpacking revert reason")
	}
	return vs[0].(string), nil
}

// BaseFee returns the BaseFee per gas if the block was mined immediately.
//
// WORKS ONLY ON AVALANCHE C-CHAIN LIKE CHAINS
func BaseFee(ctx context.Context, client *ethclient.Client) (*big.Int, error) {
	var result hexutil.Big
	err := client.Client().CallContext(ctx, &result, "eth_baseFee")
	return (*big.Int)(&result), err
}

// prepareAndSignType0 prepares a type 0 (legacy) transaction and signs it.
func prepareAndSignType0(client *ethclient.Client, gasConfig *config.Gas, privateKey *ecdsa.PrivateKey, chainID *big.Int, nonce uint64, gasLimit uint64, toAddress common.Address, value *big.Int, data []byte, timeout time.Duration) (*types.Transaction, error) {
	gasPrice, err := GetGasPrice(gasConfig, client, timeout)
	if err != nil {
		return nil, err
	}

	txData := types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		Data:     data,
	}

	tx := types.NewTx(&txData)
	signedTx, err := types.SignTx(tx, types.NewCancunSigner(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

// prepareAndSignType0 prepares a type 2 (eip 1559) transaction and signs it.
func prepareAndSignType2(client *ethclient.Client, gasConfig *config.Gas, privateKey *ecdsa.PrivateKey, chainID *big.Int, nonce uint64, gasLimit uint64, toAddress common.Address, value *big.Int, data []byte, timeout time.Duration) (*types.Transaction, error) {
	gasFeeCap := new(big.Int)
	if gasConfig.BaseFeePerGasCap != nil && gasConfig.BaseFeePerGasCap.Cmp(big.NewInt(0)) == 1 {
		gasFeeCap.Set(gasConfig.BaseFeePerGasCap)
	} else {
		ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
		baseFeePerGas, err := BaseFee(ctx, client)
		cancelFunc()

		if err != nil {
			logger.Debug("Error getting baseFee, %v", err)
			return nil, err
		}

		if gasConfig.BaseFeeMultiplier != nil && gasConfig.BaseFeeMultiplier.Cmp(common.Big0) == 1 {
			gasFeeCap = gasFeeCap.Mul(baseFeePerGas, gasConfig.BaseFeeMultiplier)
		} else {
			gasFeeCap = gasFeeCap.Mul(baseFeePerGas, big.NewInt(baseFeeCapMultiplier))
		}
	}

	tipCap := new(big.Int)
	if gasConfig.MaxPriorityFeePerGas != nil && gasConfig.MaxPriorityFeePerGas.Cmp(big.NewInt(0)) == 1 {
		tipCap.Set(gasConfig.MaxPriorityFeePerGas)
	} else {
		tipCap.Set(DefaultTipCap)
	}

	gasFeeCap = gasFeeCap.Add(gasFeeCap, tipCap)

	txData := types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: tipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     value,
		Data:      data,
	}

	tx := types.NewTx(&txData)
	signedTx, err := types.SignTx(tx, types.NewCancunSigner(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

// SendRawTx sends a transaction to toAddress with input data with prescribed nonce and gasConfig.
func SendRawTx(client *ethclient.Client, privateKey *ecdsa.PrivateKey, nonce uint64, toAddress common.Address, data []byte, dryRun bool, gasConfig *config.Gas, timeout time.Duration) error {
	var err error

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	value := big.NewInt(0)

	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	chainID, err := client.NetworkID(ctx)
	cancelFunc()
	if err != nil {
		return err
	}

	var gasLimit uint64
	if dryRun && gasConfig.GasLimit > 0 {
		gasLimit = uint64(gasConfig.GasLimit)
		_, err = DryRunTx(client, fromAddress, toAddress, value, data, timeout)
		if err != nil {
			return errors.Wrap(err, "dry run failed")
		}
	} else if dryRun {
		gasLimit, err = DryRunTx(client, fromAddress, toAddress, value, data, timeout)
		if err != nil {
			return errors.Wrap(err, "dry run failed")
		}
	} else {
		gasLimit = getGasLimit(gasConfig, client, fromAddress, toAddress, value, data, timeout)
	}

	var signedTx *types.Transaction
	switch gasConfig.TxType {
	case 0:
		signedTx, err = prepareAndSignType0(client, gasConfig, privateKey, chainID, nonce, gasLimit, toAddress, value, data, timeout)
	case 2:
		signedTx, err = prepareAndSignType2(client, gasConfig, privateKey, chainID, nonce, gasLimit, toAddress, value, data, timeout)
	default:
		return errors.New("unsupported tx type: set TxType to 0 or 2")
	}
	if err != nil {
		return errors.Wrap(err, "preparing tx")
	}

	logger.Debugf("Sending signed tx: %s, nonce: %d", signedTx.Hash().Hex(), nonce)
	ctx, cancelFunc = context.WithTimeout(context.Background(), timeout)
	err = client.SendTransaction(ctx, signedTx)
	cancelFunc()
	if err != nil {
		return err
	}

	verifier := NewTxVerifier(client)

	err = verifier.WaitUntilMined(fromAddress, signedTx, timeout)
	if err != nil {
		return err
	}
	logger.Debugf("Successful tx: %s", signedTx.Hash().Hex())

	return nil
}

// DryRunTx locally executes a transaction with the current state of blockchain and returns estimated Gas multiplied with 1.5 and potential errors.
func DryRunTx(client *ethclient.Client, fromAddress common.Address, toAddress common.Address, value *big.Int, data []byte, timeout time.Duration) (uint64, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	estimatedGas, err := estimateGas(ctx, client, fromAddress, toAddress, value, data)
	cancelFunc()
	return 3 * estimatedGas / 2, err
}

// DryRunTxAbi locally executes a transaction to method with arguments with the current state of blockchain and returns estimated Gas multiplied with 1.5 and potential errors.
func DryRunTxAbi(client *ethclient.Client, timeout time.Duration, fromAddress common.Address, toAddress common.Address, value *big.Int, abi *abi.ABI, method string, arguments ...any) (uint64, error) {
	data, err := abi.Pack(method, arguments...)
	if err != nil {
		return 0, errors.Wrap(err, "DryRunTxAbi packing")
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	estimatedGas, err := estimateGas(ctx, client, fromAddress, toAddress, value, data)
	cancelFunc()
	return 3 * estimatedGas / 2, err
}

func estimateGas(ctx context.Context, client *ethclient.Client, from, to common.Address, value *big.Int, data []byte) (uint64, error) {
	return client.EstimateGas(ctx, ethereum.CallMsg{
		From:  from,
		To:    &to,
		Value: value,
		Data:  data,
	})
}

func getGasLimit(gasConfig *config.Gas, client *ethclient.Client, fromAddress common.Address, toAddress common.Address, value *big.Int, data []byte, timeout time.Duration) uint64 {
	var gasLimit uint64
	if gasConfig.GasLimit == 0 {
		ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
		estimatedGas, err := client.EstimateGas(ctx, ethereum.CallMsg{
			From:  fromAddress,
			To:    &toAddress,
			Value: value,
			Data:  data,
		})
		cancelFunc()
		if err != nil {
			logger.Warnf("Unable to estimate gas: %v, using default gas limit: %d", err, DefaultGasLimit)
			gasLimit = DefaultGasLimit
		} else {
			gasLimit = 3 * estimatedGas / 2
			logger.Debugf("Gas limit: %d", gasLimit)
		}
	} else {
		gasLimit = uint64(gasConfig.GasLimit)
	}
	return gasLimit
}

func GetGasPrice(gasConfig *config.Gas, client *ethclient.Client, timeout time.Duration) (*big.Int, error) {
	var gasPrice *big.Int
	if gasConfig.GasPriceFixed != nil && gasConfig.GasPriceFixed.Cmp(common.Big0) == 1 {
		gasPrice = gasConfig.GasPriceFixed
	} else {
		ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
		suggestedPrice, err := client.SuggestGasPrice(ctx)
		cancelFunc()
		if err != nil {
			return nil, errors.Wrap(err, "Unable to estimate gas price")
		}
		if gasConfig.GasPriceMultiplier != 0 {
			gasPriceFloat := new(big.Float).SetInt(suggestedPrice)
			gasPriceMultiplierFloat := new(big.Float).SetFloat64(float64(gasConfig.GasPriceMultiplier))
			gasPriceFloat.Mul(gasPriceFloat, gasPriceMultiplierFloat)
			gasPrice, _ = gasPriceFloat.Int(nil)
		} else {
			gasPrice = suggestedPrice
		}
	}
	return gasPrice, nil
}

// GasConfigForAttempt sets gas config for a retry attempt.
//
// For type 0 transaction, it bumps up GasPriceMultiplier for each retry attempt by 50%,
// up to a maximum of 10x the original value.
// If GasPriceFixed is used, the retry multiplier will not be applied.
//
// For type 2 transaction, MaxPriorityFeePerGas on the n-the attempt is n-th power of 1,2 times the MaxPriorityFeePerGas of the initial attempt.
func GasConfigForAttempt(cfg *config.Gas, attempt int) *config.Gas {
	if cfg.TxType == 0 {
		if cfg.GasPriceFixed != nil && cfg.GasPriceFixed.Cmp(common.Big0) != 0 {
			return cfg
		}

		retryMultiplier := min(10.0, math.Pow(1.5, float64(attempt)))

		return &config.Gas{
			TxType:   0,
			GasLimit: cfg.GasLimit,

			GasPriceMultiplier: max(1.0, cfg.GasPriceMultiplier) * float32(retryMultiplier),
			GasPriceFixed:      cfg.GasPriceFixed,
		}
	} else if cfg.TxType == 2 {
		tipCap := new(big.Int)
		if cfg.MaxPriorityFeePerGas != nil && cfg.MaxPriorityFeePerGas.Cmp(big.NewInt(0)) == 1 {
			tipCap.Set(cfg.MaxPriorityFeePerGas)
		} else {
			tipCap.Set(DefaultTipCap)
		}

		attemptBig := big.NewInt(int64(attempt))

		multiplierTemp := big.NewInt(retryTipMultiplierTimes10)
		multiplierTemp = multiplierTemp.Exp(multiplierTemp, attemptBig, nil)

		normalizer := new(big.Int).Exp(big.NewInt(10), attemptBig, nil)

		tipCap = tipCap.Mul(tipCap, multiplierTemp)
		tipCap = tipCap.Div(tipCap, normalizer)

		baseFeeMultiplier := new(big.Int)
		if cfg.BaseFeeMultiplier != nil && cfg.BaseFeeMultiplier.Cmp(common.Big0) == 1 {
			baseFeeMultiplier = baseFeeMultiplier.Set(cfg.BaseFeeMultiplier)
			baseFeeMultiplier = baseFeeMultiplier.Add(baseFeeMultiplier, attemptBig)
		} else {
			baseFeeMultiplier = big.NewInt(3 + int64(attempt))
		}
		return &config.Gas{
			TxType:   2,
			GasLimit: cfg.GasLimit,

			BaseFeeMultiplier:    baseFeeMultiplier,
			MaxPriorityFeePerGas: tipCap,
			BaseFeePerGasCap:     cfg.BaseFeePerGasCap,
		}
	} else {
		return cfg
	}
}
