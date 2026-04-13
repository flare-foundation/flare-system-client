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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/flare-foundation/go-flare-common/pkg/logger"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/preregistry"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/registry"
)

// TODO configure addresses and breaking change

const (
	chainIDCoston  = 16
	chainIDCoston2 = 114

	breakingEpochCoston = 5451
)

const (
	newRegistryCostonAddr = "0x4C797636FC2410e1BbA7CF4bf2e397d94e65DfB8"
	oldRegistryCostonAddr = "0xE2c06DF29d175Aa0EcfcD10134eB96f8C94448A3"

	newPreRegistryCostonAddr = "0x8F8E788d3A019bA0F05aa7d6f596c564EEbeea7c"
	oldPreRegistryCostonAddr = "0xD9e8Ee56CB5A06f2070A2793bf76C5dFb3fc3D52"
)

var (
	newRegistryCoston = common.HexToAddress(newRegistryCostonAddr)
	oldRegistryCoston = common.HexToAddress(oldRegistryCostonAddr)

	newPreRegistryCoston = common.HexToAddress(newPreRegistryCostonAddr)
	oldPreRegistryCoston = common.HexToAddress(oldPreRegistryCostonAddr)
)

var (
	registratorArguments    abi.Arguments
	registratorArgumentsNew abi.Arguments

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
		{ // nextRewardEpochIDs
			Type: uint32Ty,
		},
		{ // address
			Type: addressTy,
		},
	}

	registratorArgumentsNew = abi.Arguments{
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
	oldRegistry        *registry.Registry
	preregistry        *preregistry.Preregistry
	oldPreregistry     *preregistry.Preregistry
	senderTxOpts       *bind.TransactOpts
	gasCfg             *config.Gas
	txVerifier         *chain.TxVerifier
	signerPrivateKey   *ecdsa.PrivateKey
	chainID            int
}

func NewRegistryContractClient(
	ethClient *ethclient.Client,
	gasCfg *config.Gas,
	registryAddress common.Address,
	preregistryAddress common.Address,
	senderTxOpts *bind.TransactOpts,
	signerPk *ecdsa.PrivateKey,
	chainID int,
) (*registryContractClientImpl, error) {
	registryBinding, err := registry.NewRegistry(registryAddress, ethClient)
	if err != nil {
		return nil, err
	}
	preregistryBinding, err := preregistry.NewPreregistry(preregistryAddress, ethClient)
	if err != nil {
		return nil, err
	}

	var oldRegistryBinding *registry.Registry
	if registryAddress == newRegistryCoston {
		oldRegistryBinding, err = registry.NewRegistry(oldRegistryCoston, ethClient)
		if err != nil {
			return nil, err
		}
	}

	var oldPreRegistryBinding *preregistry.Preregistry
	if preregistryAddress == newPreRegistryCoston {
		oldPreRegistryBinding, err = preregistry.NewPreregistry(oldPreRegistryCoston, ethClient)
		if err != nil {
			return nil, err
		}
	}

	return &registryContractClientImpl{
		ethClient:          ethClient,
		registryAddress:    registryAddress,
		preregistryAddress: preregistryAddress,
		registry:           registryBinding,
		oldRegistry:        oldRegistryBinding,
		preregistry:        preregistryBinding,
		oldPreregistry:     oldPreRegistryBinding,
		senderTxOpts:       senderTxOpts,
		gasCfg:             gasCfg,
		txVerifier:         chain.NewTxVerifier(ethClient),
		signerPrivateKey:   signerPk,
		chainID:            chainID,
	}, nil
}

// RegisterVoter tries to register voter on VoterRegistry smart contract.
func (r *registryContractClientImpl) RegisterVoter(ctx context.Context, nextRewardEpochID *big.Int, address common.Address) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetryChan(func() (any, error) {
		err := r.sendRegisterVoter(ctx, nextRewardEpochID, address)
		if err != nil {
			if shared.ExistsAsSubstring(nonFatalRegisterErrors, err.Error()) {
				logger.Debugf("Non fatal error sending register voter: %v", err)
			} else {
				return nil, errors.Wrap(err, "error sending register voter")
			}
		}
		return nil, nil
	}, shared.MaxTxSendRetriesLong, shared.TxRetryIntervalLong)
}

func (r *registryContractClientImpl) sendRegisterVoter(ctx context.Context, nextRewardEpochID *big.Int, address common.Address) error {
	epochID := uint32(nextRewardEpochID.Uint64())
	signature, err := r.signature(epochID, address)
	if err != nil {
		return fmt.Errorf("signature: %w", err)
	}

	vrsSignature := registry.IVoterRegistrySignature{
		R: [32]byte(signature[0:32]),
		S: [32]byte(signature[32:64]),
		V: signature[64] + 27,
	}

	err = SetGas(ctx, r.senderTxOpts, r.ethClient, r.gasCfg)
	if err != nil {
		return fmt.Errorf("setting gas: %w", err)
	}

	estimatedGasLimit, err := chain.DryRunTxAbi(
		ctx,
		r.ethClient,
		chain.DefaultTxTimeout,
		r.senderTxOpts.From,
		r.registryAddress,
		common.Big0,
		registryAbi,
		"registerVoter",
		address,
		vrsSignature,
	)
	if err != nil {
		return fmt.Errorf("dry run failed: %w", err)
	}

	if r.gasCfg.GasLimit != 0 {
		r.senderTxOpts.GasLimit = uint64(r.gasCfg.GasLimit)
	} else {
		r.senderTxOpts.GasLimit = estimatedGasLimit
	}

	var tx *types.Transaction

	if shouldUseOldRegistry(epochID, r.registryAddress) {
		tx, err = r.oldRegistry.RegisterVoter(r.senderTxOpts, address, vrsSignature)
		if err != nil {
			return fmt.Errorf("sending registry old tx: %w", err)
		}
	} else {
		tx, err = r.registry.RegisterVoter(r.senderTxOpts, address, vrsSignature)
		if err != nil {
			return fmt.Errorf("sending registry tx: %w", err)
		}
	}

	err = r.txVerifier.WaitUntilMined(ctx, r.senderTxOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Infof("Voter %s registered for epoch %v", address, nextRewardEpochID)
	return nil
}

// PreregisterVoter tries to pre-register voter on VoterPreRegistry smart contract.
func (r *registryContractClientImpl) PreregisterVoter(ctx context.Context, nextRewardEpochId *big.Int, address common.Address) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetryChan(func() (any, error) {
		err := r.sendPreRegisterVoter(ctx, nextRewardEpochId, address)
		if err != nil {
			if shared.ExistsAsSubstring(nonFatalPreregisterErrors, err.Error()) {
				logger.Debugf("Non fatal error sending pre-register voter: %v", err)
			} else {
				return nil, errors.Wrap(err, "error sending pre-register voter")
			}
		}
		return nil, nil
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)
}

func (r *registryContractClientImpl) sendPreRegisterVoter(ctx context.Context, nextRewardEpochID *big.Int, address common.Address) error {
	epochID := uint32(nextRewardEpochID.Uint64())
	signature, err := r.signature(epochID, address)
	if err != nil {
		return fmt.Errorf("signature: %w", err)
	}

	vrsSignature := preregistry.IVoterRegistrySignature{
		R: [32]byte(signature[0:32]),
		S: [32]byte(signature[32:64]),
		V: signature[64] + 27,
	}

	err = SetGas(ctx, r.senderTxOpts, r.ethClient, r.gasCfg)
	if err != nil {
		return fmt.Errorf("setting gas pre registry: %w", err)
	}

	estimatedGasLimit, err := chain.DryRunTxAbi(
		ctx,
		r.ethClient,
		chain.DefaultTxTimeout,
		r.senderTxOpts.From,
		r.preregistryAddress,
		common.Big0,
		preregistryAbi,
		"preRegisterVoter",
		address,
		vrsSignature,
	)
	if err != nil {
		return fmt.Errorf("dry run failed, %w", err)
	}

	if r.gasCfg.GasLimit != 0 {
		r.senderTxOpts.GasLimit = uint64(r.gasCfg.GasLimit)
	} else {
		r.senderTxOpts.GasLimit = estimatedGasLimit
	}

	var tx *types.Transaction

	if shouldUseOldPreRegistry(epochID, r.preregistryAddress) {
		tx, err = r.oldPreregistry.PreRegisterVoter(r.senderTxOpts, address, vrsSignature)
		if err != nil {
			return fmt.Errorf("sending preregistry tx: %w", err)
		}
	} else {
		tx, err = r.preregistry.PreRegisterVoter(r.senderTxOpts, address, vrsSignature)
		if err != nil {
			return fmt.Errorf("sending preregistry tx: %w", err)
		}
	}

	err = r.txVerifier.WaitUntilMined(ctx, r.senderTxOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Infof("Voter %s pre-registered for epoch %v", address, nextRewardEpochID)
	return nil
}

// createSignature creates ECDSA message signature keccak256(abi.encode(nextRewardEpochID, address)) with signerPrivateKey
func (r *registryContractClientImpl) createSignature(nextRewardEpochID uint32, address common.Address) ([]byte, error) {
	message, err := registratorArguments.Pack(nextRewardEpochID, address)
	if err != nil {
		return nil, err
	}
	messageHash := crypto.Keccak256(message)
	return crypto.Sign(accounts.TextHash(messageHash), r.signerPrivateKey)
}

// createSignatureNew creates ECDSA message signature keccak256(abi.encode(chainID, nextRewardEpochID, address)) with signerPrivateKey
func (r *registryContractClientImpl) createSignatureNew(chainID int, nextRewardEpochID uint32, address common.Address) ([]byte, error) {
	chainIDB := big.NewInt(int64(chainID))

	message, err := registratorArgumentsNew.Pack(chainIDB, nextRewardEpochID, address)
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
		feeCtx, cancelFunc := context.WithTimeout(ctx, chain.DefaultTxTimeout)
		baseFeePerGas, err := chain.BaseFee(feeCtx, client)
		cancelFunc()

		if err != nil {
			logger.Debug("Error getting baseFee, %v", err)
			return err
		}

		gasFeeCap := new(big.Int)
		if gasConfig.BaseFeePerGasCap != nil && gasConfig.BaseFeePerGasCap.Cmp(big.NewInt(0)) == 1 {
			gasFeeCap.Set(gasConfig.BaseFeePerGasCap)
		} else {
			if gasConfig.BaseFeeMultiplier != nil && gasConfig.BaseFeeMultiplier.Cmp(common.Big0) == 1 {
				gasFeeCap = gasFeeCap.Mul(baseFeePerGas, gasConfig.BaseFeeMultiplier)
			} else {
				gasFeeCap = gasFeeCap.Mul(baseFeePerGas, chain.DefaultBaseFeeCapMultiplier)
			}
		}

		tipCap := new(big.Int)
		if gasConfig.MaxPriorityMultiplier != nil && gasConfig.MaxPriorityMultiplier.Cmp(big.NewInt(0)) == 1 {
			tipCap.Mul(baseFeePerGas, gasConfig.MaxPriorityMultiplier)
		} else {
			tipCap.Mul(baseFeePerGas, chain.DefaultTipMultiplier)
		}
		gasFeeCap = gasFeeCap.Add(gasFeeCap, tipCap)

		txOptions.GasFeeCap = gasFeeCap
		txOptions.GasTipCap = tipCap
		return nil
	default:
		// should never happen. txType is checked when config is read from toml file.
		return fmt.Errorf("unsupported tx type: %d", gasConfig.TxType)
	}
}

func shouldUseOldRegistry(epochID uint32, address common.Address) bool {
	if address == newRegistryCoston && epochID < breakingEpochCoston {
		return true
	}

	return false
}

func shouldUseOldPreRegistry(epochID uint32, address common.Address) bool {
	if address == newPreRegistryCoston && epochID < breakingEpochCoston {
		return true
	}

	return false
}

func (r *registryContractClientImpl) signature(epochID uint32, address common.Address) ([]byte, error) {
	var (
		signature []byte
		err       error
	)
	switch {
	case r.chainID == chainIDCoston2,
		r.chainID == chainIDCoston && epochID >= breakingEpochCoston:
		signature, err = r.createSignatureNew(r.chainID, epochID, address)
		if err != nil {
			return nil, fmt.Errorf("creating pre registry signature new: %w", err)
		}
	default:
		signature, err = r.createSignature(epochID, address)
		if err != nil {
			return nil, fmt.Errorf("creating pre registry signature old: %w", err)
		}
	}

	return signature, nil
}
