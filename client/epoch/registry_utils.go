package epoch

import (
	"crypto/ecdsa"
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

	"github.com/flare-foundation/go-flare-common/pkg/contracts/registry"
)

var (
	registratorArguments abi.Arguments
	fallbackGasPrice     = big.NewInt(50 * 1e9)
)

func init() {
	uint32Ty, _ := abi.NewType("uint32", "uint32", nil)
	addressTy, _ := abi.NewType("address", "address", nil)
	registratorArguments = abi.Arguments{
		{ // nextRewardEpochId
			Type: uint32Ty,
		},
		{ // address
			Type: addressTy,
		},
	}
}

type registryContractClient interface {
	RegisterVoter(nextRewardEpochId *big.Int, address common.Address) <-chan shared.ExecuteStatus[any]
}

type registryContractClientImpl struct {
	ethClient        *ethclient.Client
	address          common.Address
	registry         *registry.Registry
	senderTxOpts     *bind.TransactOpts
	gasCfg           *config.Gas
	txVerifier       *chain.TxVerifier
	signerPrivateKey *ecdsa.PrivateKey
}

func NewRegistryContractClient(
	ethClient *ethclient.Client,
	gasCfg *config.Gas,
	address common.Address,
	senderTxOpts *bind.TransactOpts,
	signerPk *ecdsa.PrivateKey,
) (*registryContractClientImpl, error) {
	registry, err := registry.NewRegistry(address, ethClient)
	if err != nil {
		return nil, err
	}
	return &registryContractClientImpl{
		ethClient:        ethClient,
		address:          address,
		registry:         registry,
		senderTxOpts:     senderTxOpts,
		gasCfg:           gasCfg,
		txVerifier:       chain.NewTxVerifier(ethClient),
		signerPrivateKey: signerPk,
	}, nil

}

func (r *registryContractClientImpl) RegisterVoter(nextRewardEpochId *big.Int, address common.Address) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetryChan(func() (any, error) {
		err := r.sendRegisterVoter(nextRewardEpochId, address)
		if err != nil {
			return nil, errors.Wrap(err, "error sending register voter")
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

	gasPrice, err := chain.GetGasPrice(r.gasCfg, r.ethClient, chain.DefaultTxTimeout)
	if err != nil {
		logger.Warnf("Unable to obtain gas price: %v, using fallback %d", err, fallbackGasPrice)
		gasPrice = fallbackGasPrice
	}

	if r.gasCfg.GasLimit != 0 {
		r.senderTxOpts.GasLimit = uint64(r.gasCfg.GasLimit)
	}
	r.senderTxOpts.GasPrice = gasPrice
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

func (r *registryContractClientImpl) createSignature(nextRewardEpochId uint32, address common.Address) ([]byte, error) {
	message, err := registratorArguments.Pack(nextRewardEpochId, address)
	if err != nil {
		return nil, err
	}
	messageHash := crypto.Keccak256(message)
	return crypto.Sign(accounts.TextHash(messageHash), r.signerPrivateKey)
}
