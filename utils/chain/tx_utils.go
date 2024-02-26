package chain

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"flare-tlc/client/config"
	"flare-tlc/logger"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

const (
	// default timeout for waiting for a tx to be mined.
	DefaultTxTimeout = 60 * time.Second
	DefaultGasLimit  = 2_500_000
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
	if !bytes.Equal(result[:4], errorSig) {
		return "<tx result not Error(string)>", errors.New("tx result not of type Error(string)")
	}
	vs, err := abi.Arguments{{Type: abiString}}.UnpackValues(result[4:])
	if err != nil {
		return "<invalid tx result>", errors.Wrap(err, "unpacking revert reason")
	}
	return vs[0].(string), nil
}

func TransactOptsFromPrivateKey(privateKey string, chainID int) (*bind.TransactOpts, error) {
	opts, _, err := CredentialsFromPrivateKey(privateKey, chainID)
	return opts, err
}

func CredentialsFromPrivateKey(privateKey string, chainID int) (*bind.TransactOpts, *ecdsa.PrivateKey, error) {
	pk, err := PrivateKeyFromHex(privateKey)
	if err != nil {
		return nil, nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(
		pk, big.NewInt(int64(chainID)),
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	// bind.N
	return opts, pk, nil
}

func PrivateKeyFromHex(privateKey string) (*ecdsa.PrivateKey, error) {
	if len(privateKey) < 2 {
		return nil, errors.New("privateKey is too short")
	}

	privateKey = strings.TrimPrefix(privateKey, "0x")

	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "crypto.HexToECDSA")
	}
	return pk, nil
}

func SendRawTx(client *ethclient.Client, privateKey *ecdsa.PrivateKey, toAddress common.Address, data []byte, gasConfig *config.GasConfig) error {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	value := big.NewInt(0) // in wei (1 eth)

	gasLimit := getGasLimit(gasConfig, client, fromAddress, toAddress, value, data)
	gasPrice, err := GetGasPrice(gasConfig, client)
	if err != nil {
		return err
	}

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return err
	}

	logger.Debug("Sending signed tx: %s", signedTx.Hash().Hex())
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	verifier := NewTxVerifier(client)

	logger.Debug("Waiting for tx to be mined...")
	err = verifier.WaitUntilMined(fromAddress, signedTx, DefaultTxTimeout)
	if err != nil {
		return err
	}

	logger.Debug("Tx mined, getting receipt %s", signedTx.Hash().Hex())
	rec, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		return err
	}
	logger.Debug("Receipt status: %v", rec.Status)
	return nil
}

func getGasLimit(gasConfig *config.GasConfig, client *ethclient.Client, fromAddress common.Address, toAddress common.Address, value *big.Int, data []byte) uint64 {
	var gasLimit uint64
	if gasConfig.GasLimit == 0 {
		estimatedGas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			From:  fromAddress,
			To:    &toAddress,
			Value: value,
			Data:  data,
		})
		if err != nil {
			logger.Warn("Unable to estimate gas: %v, using default gas limit: %d", err, DefaultGasLimit)
			gasLimit = DefaultGasLimit
		} else {
			logger.Debug("Estimated gas: %d", estimatedGas)
			gasLimit = estimatedGas
		}
	} else {
		gasLimit = uint64(gasConfig.GasLimit)
	}
	return gasLimit
}

func GetGasPrice(gasConfig *config.GasConfig, client *ethclient.Client) (*big.Int, error) {
	var gasPrice *big.Int
	if gasConfig.GasPriceFixed.Cmp(common.Big0) != 0 {
		gasPrice = gasConfig.GasPriceFixed
	} else {
		suggestedPrice, err := client.SuggestGasPrice(context.Background())
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
