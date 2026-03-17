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
	"github.com/pkg/errors"

	"github.com/flare-foundation/go-flare-common/pkg/logger"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/preregistry"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/registry"
)

const (
	chainIDCoston2 = 114
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
	RegisterVoter(nextRewardEpochId *big.Int, address common.Address) <-chan shared.ExecuteStatus[any]
	PreregisterVoter(nextRewardEpochId *big.Int, address common.Address) <-chan shared.ExecuteStatus[any]
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
func (r *registryContractClientImpl) RegisterVoter(nextRewardEpochID *big.Int, address common.Address) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetryChan(func() (any, error) {
		err := r.sendRegisterVoter(nextRewardEpochID, address)
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

func (r *registryContractClientImpl) sendRegisterVoter(nextRewardEpochID *big.Int, address common.Address) error {
	epochID := uint32(nextRewardEpochID.Uint64())

	var (
		signature []byte
		err       error
	)

	if r.chainID != chainIDCoston2 {
		signature, err = r.createSignatureOld(epochID, address)
		if err != nil {
			return fmt.Errorf("creating registry signature old: %w", err)
		}
	} else {
		signature, err = r.createSignatureNew(r.chainID, epochID, address)
		if err != nil {
			return fmt.Errorf("creating pre registry signature new: %w", err)
		}
	}

	vrsSignature := registry.IVoterRegistrySignature{
		R: [32]byte(signature[0:32]),
		S: [32]byte(signature[32:64]),
		V: signature[64] + 27,
	}

	err = SetGas(r.senderTxOpts, r.ethClient, r.gasCfg)
	if err != nil {
		return err
	}

	estimatedGasLimit, err := chain.DryRunTxAbi(
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
		return errors.Wrap(err, "Dry run failed")
	}

	if r.gasCfg.GasLimit != 0 {
		r.senderTxOpts.GasLimit = uint64(r.gasCfg.GasLimit)
	} else {
		r.senderTxOpts.GasLimit = estimatedGasLimit
	}

	tx, err := r.registry.RegisterVoter(r.senderTxOpts, address, vrsSignature)
	if err != nil {
		return fmt.Errorf("sending registry tx: %w", err)
	}

	err = r.txVerifier.WaitUntilMined(r.senderTxOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Infof("Voter %s registered for epoch %v", address, nextRewardEpochID)
	return nil
}

// PreregisterVoter tries to pre-register voter on VoterPreRegistry smart contract.
func (r *registryContractClientImpl) PreregisterVoter(nextRewardEpochId *big.Int, address common.Address) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetryChan(func() (any, error) {
		err := r.sendPreRegisterVoter(nextRewardEpochId, address)
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

func (r *registryContractClientImpl) sendPreRegisterVoter(nextRewardEpochID *big.Int, address common.Address) error {
	epochID := uint32(nextRewardEpochID.Uint64())

	var (
		signature []byte
		err       error
	)

	if r.chainID != chainIDCoston2 {
		signature, err = r.createSignatureOld(epochID, address)
		if err != nil {
			return fmt.Errorf("creating registry signature old: %w", err)
		}
	} else {
		signature, err = r.createSignatureNew(r.chainID, epochID, address)
		if err != nil {
			return fmt.Errorf("creating pre registry signature new: %w", err)
		}
	}

	vrsSignature := preregistry.IVoterRegistrySignature{
		R: [32]byte(signature[0:32]),
		S: [32]byte(signature[32:64]),
		V: signature[64] + 27,
	}

	err = SetGas(r.senderTxOpts, r.ethClient, r.gasCfg)
	if err != nil {
		return fmt.Errorf("setting gas pre registry:%w", err)
	}

	estimatedGasLimit, err := chain.DryRunTxAbi(
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
		return errors.Wrap(err, "Dry run failed")
	}

	if r.gasCfg.GasLimit != 0 {
		r.senderTxOpts.GasLimit = uint64(r.gasCfg.GasLimit)
	} else {
		r.senderTxOpts.GasLimit = estimatedGasLimit
	}

	tx, err := r.preregistry.PreRegisterVoter(r.senderTxOpts, address, vrsSignature)
	if err != nil {
		return fmt.Errorf("sending preregistry tx: %w", err)
	}

	err = r.txVerifier.WaitUntilMined(r.senderTxOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Infof("Voter %s pre-registered for epoch %v", address, nextRewardEpochID)
	return nil
}

// createSignatureOld creates ECDSA message signature keccak256(abi.encode(nextRewardEpochID, address)) with signerPrivateKey
func (r *registryContractClientImpl) createSignatureOld(nextRewardEpochID uint32, address common.Address) ([]byte, error) {
	message, err := registratorArguments.Pack(nextRewardEpochID, address)
	if err != nil {
		return nil, err
	}
	messageHash := crypto.Keccak256(message)
	return crypto.Sign(accounts.TextHash(messageHash), r.signerPrivateKey)
}

// createSignatureOld creates ECDSA message signature keccak256(abi.encode(chainID, nextRewardEpochID, address)) with signerPrivateKey
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
func SetGas(txOptions *bind.TransactOpts, client *ethclient.Client, gasConfig *config.Gas) error {
	switch gasConfig.TxType {
	case 0:
		gasPrice, err := chain.GetGasPrice(gasConfig, client, chain.DefaultTxTimeout)
		if err != nil {
			logger.Warnf("Unable to obtain gas price: %v, using fallback %d", err, fallbackGasPrice)
			gasPrice = new(big.Int).Set(fallbackGasPrice)
		}
		txOptions.GasPrice = gasPrice
		return nil
	case 2:
		ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
		baseFeePerGas, err := chain.BaseFee(ctx, client)
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
