package epoch

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils/chain"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/flare-foundation/go-flare-common/pkg/logger"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/preregistry"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/registry"
)

var (
	registratorArguments abi.Arguments

	registryAbi    *abi.ABI
	preregistryAbi *abi.ABI

	fallbackGasPrice = big.NewInt(50 * 1e9) // 50 GWei
)

var (
	nonFatalRegisterErrors = []string{
		"already registered",
		"voter registration not enabled",
	}
	nonFatalPreregisterErrors = []string{
		"voter already pre-registered",
		"voter currently not registered",
		"pre-registration not opened anymore",
	}
)

func init() {
	uint32Ty, err := abi.NewType("uint32", "uint32", nil)
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}

	addressTy, err := abi.NewType("address", "address", nil)
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}

	uint256Ty, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}

	registratorArguments = abi.Arguments{
		{ // chainID
			Type: uint256Ty,
		},
		{ // nextRewardEpochID
			Type: uint32Ty,
		},
		{ // address
			Type: addressTy,
		},
	}

	registryAbi, err = registry.RegistryMetaData.GetAbi()
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	preregistryAbi, err = preregistry.PreregistryMetaData.GetAbi()
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
}

type registryContractClient interface {
	RegisterVoter(ctx context.Context, nextRewardEpochId *big.Int, address common.Address) <-chan shared.ExecuteStatus[any]
	PreregisterVoter(ctx context.Context, nextRewardEpochId *big.Int, address common.Address) <-chan shared.ExecuteStatus[any]
}

type registryContractClientImpl struct {
	ethClient          *ethclient.Client
	registryAddress    common.Address
	preregistryAddress common.Address
	registry           *registry.Registry
	preregistry        *preregistry.Preregistry
	senderTxOpts       *bind.TransactOpts
	gasCfg             *config.Gas
	txVerifier         *chain.TxVerifier
	signerPrivateKey   *ecdsa.PrivateKey
	chainID            int64
}

func NewRegistryContractClient(
	ethClient *ethclient.Client,
	gasCfg *config.Gas,
	registryAddress common.Address,
	preregistryAddress common.Address,
	senderTxOpts *bind.TransactOpts,
	signerPk *ecdsa.PrivateKey,
	chainID int64,
) (*registryContractClientImpl, error) {
	registryBinding, err := registry.NewRegistry(registryAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("registry binding: %w", err)
	}
	preregistryBinding, err := preregistry.NewPreregistry(preregistryAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("pre registry binding: %w", err)
	}

	return &registryContractClientImpl{
		ethClient:          ethClient,
		registryAddress:    registryAddress,
		preregistryAddress: preregistryAddress,
		registry:           registryBinding,
		preregistry:        preregistryBinding,
		senderTxOpts:       senderTxOpts,
		gasCfg:             gasCfg,
		txVerifier:         chain.NewTxVerifier(ethClient),
		signerPrivateKey:   signerPk,
		chainID:            chainID,
	}, nil
}

// RegisterVoter tries to register voter on VoterRegistry smart contract.
func (r *registryContractClientImpl) RegisterVoter(ctx context.Context, nextRewardEpochID *big.Int, address common.Address) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetryChan(ctx, func() (any, error) {
		err := r.sendRegisterVoter(ctx, nextRewardEpochID, address)
		if err != nil {
			if shared.ExistsAsSubstring(nonFatalRegisterErrors, err.Error()) {
				logger.Debugf("Non fatal error sending register voter: %v", err)
			} else {
				return nil, fmt.Errorf("sending register voter: %w", err)
			}
		}
		return nil, nil
	}, shared.MaxTxSendRetriesLong, shared.TxRetryIntervalLong)
}

func (r *registryContractClientImpl) sendRegisterVoter(ctx context.Context, nextRewardEpochID *big.Int, address common.Address) error {
	epochID := uint32(nextRewardEpochID.Uint64())
	signature, err := r.createSignature(epochID, address)
	if err != nil {
		return fmt.Errorf("signature: %w", err)
	}

	vrsSignature := registry.IVoterRegistrySignature{
		R: [32]byte(signature[0:32]),
		S: [32]byte(signature[32:64]),
		V: signature[64] + 27,
	}

	// senderTxOpts is shared between concurrent send paths, so gas settings are applied to a copy
	txOpts := chain.CopyTxOpts(r.senderTxOpts)

	err = SetGas(ctx, txOpts, r.ethClient, r.gasCfg)
	if err != nil {
		return fmt.Errorf("setting gas: %w", err)
	}

	estimatedGasLimit, err := chain.DryRunTxAbi(
		ctx,
		r.ethClient,
		chain.DefaultTxTimeout,
		txOpts.From,
		r.registryAddress,
		common.Big0,
		registryAbi,
		"registerVoter",
		address,
		vrsSignature,
	)
	if err != nil {
		return fmt.Errorf("dry run: %w", err)
	}

	if r.gasCfg.GasLimit != 0 {
		txOpts.GasLimit = uint64(r.gasCfg.GasLimit)
	} else {
		txOpts.GasLimit = estimatedGasLimit
	}

	tx, err := r.registry.RegisterVoter(txOpts, address, vrsSignature)
	if err != nil {
		return fmt.Errorf("sending registry tx: %w", err)
	}

	err = r.txVerifier.WaitUntilMined(ctx, txOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Infof("Voter %s registered for epoch %v", address, nextRewardEpochID)
	return nil
}

// PreregisterVoter tries to pre-register voter on VoterPreRegistry smart contract.
func (r *registryContractClientImpl) PreregisterVoter(ctx context.Context, nextRewardEpochId *big.Int, address common.Address) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetryChan(ctx, func() (any, error) {
		err := r.sendPreRegisterVoter(ctx, nextRewardEpochId, address)
		if err != nil {
			if shared.ExistsAsSubstring(nonFatalPreregisterErrors, err.Error()) {
				logger.Debugf("Non fatal error sending pre-register voter: %v", err)
			} else {
				return nil, fmt.Errorf("sending pre-register voter: %w", err)
			}
		}
		return nil, nil
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)
}

func (r *registryContractClientImpl) sendPreRegisterVoter(ctx context.Context, nextRewardEpochID *big.Int, address common.Address) error {
	epochID := uint32(nextRewardEpochID.Uint64())
	signature, err := r.createSignature(epochID, address)
	if err != nil {
		return fmt.Errorf("signature: %w", err)
	}

	vrsSignature := preregistry.IVoterRegistrySignature{
		R: [32]byte(signature[0:32]),
		S: [32]byte(signature[32:64]),
		V: signature[64] + 27,
	}

	// senderTxOpts is shared between concurrent send paths, so gas settings are applied to a copy
	txOpts := chain.CopyTxOpts(r.senderTxOpts)

	err = SetGas(ctx, txOpts, r.ethClient, r.gasCfg)
	if err != nil {
		return fmt.Errorf("setting gas pre registry: %w", err)
	}

	estimatedGasLimit, err := chain.DryRunTxAbi(
		ctx,
		r.ethClient,
		chain.DefaultTxTimeout,
		txOpts.From,
		r.preregistryAddress,
		common.Big0,
		preregistryAbi,
		"preRegisterVoter",
		address,
		vrsSignature,
	)
	if err != nil {
		return fmt.Errorf("dry run failed: %w", err)
	}

	if r.gasCfg.GasLimit != 0 {
		txOpts.GasLimit = uint64(r.gasCfg.GasLimit)
	} else {
		txOpts.GasLimit = estimatedGasLimit
	}

	tx, err := r.preregistry.PreRegisterVoter(txOpts, address, vrsSignature)
	if err != nil {
		return fmt.Errorf("sending preregistry tx: %w", err)
	}

	err = r.txVerifier.WaitUntilMined(ctx, txOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Infof("Voter %s pre-registered for epoch %v", address, nextRewardEpochID)
	return nil
}

// createSignature creates ECDSA message signature keccak256(abi.encode(chainID, nextRewardEpochID, address)) with signerPrivateKey
func (r *registryContractClientImpl) createSignature(nextRewardEpochID uint32, address common.Address) ([]byte, error) {
	chainIDB := big.NewInt(r.chainID)

	message, err := registratorArguments.Pack(chainIDB, nextRewardEpochID, address)
	if err != nil {
		return nil, err
	}
	messageHash := crypto.Keccak256(message)
	return crypto.Sign(accounts.TextHash(messageHash), r.signerPrivateKey)
}

// SetGas sets gas parameters in txOptions according to the gasConfig.
func SetGas(ctx context.Context, txOptions *bind.TransactOpts, client *ethclient.Client, gasConfig *config.Gas) error {
	switch gasConfig.TxType {
	case 0:
		gasPrice, err := chain.GetGasPrice(ctx, gasConfig, client, chain.DefaultTxTimeout)
		if err != nil {
			logger.Warnf("Unable to obtain gas price: %v, using fallback %d", err, fallbackGasPrice)
			gasPrice = new(big.Int).Set(fallbackGasPrice)
		}
		txOptions.GasPrice = gasPrice
		return nil
	case 2:
		// Default unset fields so EnforceMaxPriorityFeeCaps never sees a nil cap,
		// matching the SendRawTx path.
		gasConfig := gasConfig.CopyAndDefault()

		feeCtx, cancelFunc := context.WithTimeout(ctx, chain.DefaultTxTimeout)
		baseFeePerGas, err := chain.BaseFee(feeCtx, client)
		cancelFunc()

		if err != nil {
			logger.Debugf("Error getting baseFee: %v", err)
			return err
		}

		gasFeeCap := new(big.Int)
		if gasConfig.BaseFeePerGasCap != nil && gasConfig.BaseFeePerGasCap.Sign() == 1 {
			gasFeeCap.Set(gasConfig.BaseFeePerGasCap)
		} else {
			gasFeeCap = chain.MultiplyWithFloat(baseFeePerGas, float64(gasConfig.BaseFeeMultiplier), gasFeeCap)
		}

		tipCap := chain.MultiplyWithFloat(baseFeePerGas, float64(gasConfig.MaxPriorityMultiplier), nil)
		tipCap = gasConfig.EnforceMaxPriorityFeeCaps(tipCap)

		gasFeeCap.Add(gasFeeCap, tipCap)

		txOptions.GasFeeCap = gasFeeCap
		txOptions.GasTipCap = tipCap

		return nil
	default:
		// should never happen. txType is checked when config is read from toml file.
		return fmt.Errorf("unsupported tx type: %d", gasConfig.TxType)
	}
}
