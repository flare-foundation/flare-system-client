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

var (
	registratorArguments abi.Arguments
	registryAbi          *abi.ABI
	preregistryAbi       *abi.ABI

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
	registratorArguments = abi.Arguments{
		{ // nextRewardEpochId
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
	address            common.Address
	preregistryAddress common.Address
	registry           *registry.Registry
	preregistry        *preregistry.Preregistry
	senderTxOpts       *bind.TransactOpts
	gasCfg             *config.Gas
	txVerifier         *chain.TxVerifier
	signerPrivateKey   *ecdsa.PrivateKey
}

func NewRegistryContractClient(
	ethClient *ethclient.Client,
	gasCfg *config.Gas,
	registryAddress common.Address,
	preregistryAddress common.Address,
	senderTxOpts *bind.TransactOpts,
	signerPk *ecdsa.PrivateKey,
) (*registryContractClientImpl, error) {
	registry, err := registry.NewRegistry(registryAddress, ethClient)
	if err != nil {
		return nil, err
	}
	preregistry, err := preregistry.NewPreregistry(preregistryAddress, ethClient)
	if err != nil {
		return nil, err
	}
	return &registryContractClientImpl{
		ethClient:          ethClient,
		address:            registryAddress,
		preregistryAddress: preregistryAddress,
		registry:           registry,
		preregistry:        preregistry,
		senderTxOpts:       senderTxOpts,
		gasCfg:             gasCfg,
		txVerifier:         chain.NewTxVerifier(ethClient),
		signerPrivateKey:   signerPk,
	}, nil
}

// RegisterVoter tries to register voter on VoterRegistry smart contract.
func (r *registryContractClientImpl) RegisterVoter(nextRewardEpochId *big.Int, address common.Address) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetryChan(func() (any, error) {
		err := r.sendRegisterVoter(nextRewardEpochId, address)
		if err != nil {
			if shared.ExistsAsSubstring(nonFatalRegisterErrors, err.Error()) {
				logger.Debugf("Non fatal error sending register voter: %v", err)
			} else {
				return nil, errors.Wrap(err, "error sending register voter")
			}
		}
		return nil, nil
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)
}

func (r *registryContractClientImpl) sendRegisterVoter(nextRewardEpochId *big.Int, address common.Address) error {
	epochId := uint32(nextRewardEpochId.Uint64())
	signature, err := r.createSignature(epochId, address)
	if err != nil {
		return err
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
		r.address,
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
		return err
	}
	err = r.txVerifier.WaitUntilMined(r.senderTxOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Infof("Voter %s registered for epoch %v", address, nextRewardEpochId)
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

func (r *registryContractClientImpl) sendPreRegisterVoter(nextRewardEpochId *big.Int, address common.Address) error {
	epochId := uint32(nextRewardEpochId.Uint64())
	signature, err := r.createSignature(epochId, address)
	if err != nil {
		return err
	}
	vrsSignature := preregistry.IVoterRegistrySignature{
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
		return err
	}
	err = r.txVerifier.WaitUntilMined(r.senderTxOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Infof("Voter %s pre-registered for epoch %v", address, nextRewardEpochId)
	return nil
}

// createSignature creates ECDSA message signature keccak256(abi.encode(nextRewardEpochId, address)) with signerPrivateKey
func (r *registryContractClientImpl) createSignature(nextRewardEpochId uint32, address common.Address) ([]byte, error) {
	message, err := registratorArguments.Pack(nextRewardEpochId, address)
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
		gasFeeCap := new(big.Int)
		if gasConfig.BaseFeePerGasCap != nil && gasConfig.BaseFeePerGasCap.Cmp(big.NewInt(0)) == 1 {
			gasFeeCap.Set(gasConfig.BaseFeePerGasCap)
		} else {
			ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
			baseFeePerGas, err := chain.BaseFee(ctx, client)
			cancelFunc()

			if err != nil {
				logger.Debug("Error getting baseFee, %v", err)
				return err
			}

			if gasConfig.BaseFeeMultiplier != nil && gasConfig.BaseFeeMultiplier.Cmp(common.Big0) == 1 {
				gasFeeCap = gasFeeCap.Mul(baseFeePerGas, gasConfig.BaseFeeMultiplier)
			} else {
				gasFeeCap = gasFeeCap.Mul(baseFeePerGas, chain.DefaultBaseFeeCapMultiplier)
			}
		}

		tipCap := new(big.Int)
		if gasConfig.MaxPriorityFeePerGas != nil && gasConfig.MaxPriorityFeePerGas.Cmp(big.NewInt(0)) == 1 {
			tipCap.Set(gasConfig.MaxPriorityFeePerGas)
		} else {
			tipCap.Set(chain.DefaultTipCap)
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
