package registration

import (
	"crypto/ecdsa"
	"flare-tlc/client/shared"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/registry"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

var (
	registratorArguments abi.Arguments
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

type RegistryContractClient struct {
	address          common.Address
	registry         *registry.Registry
	senderTxOpts     *bind.TransactOpts
	txVerifier       *chain.TxVerifier
	signerPrivateKey *ecdsa.PrivateKey
}

func NewRegistryContractClient(
	ethClient *ethclient.Client,
	address common.Address,
	senderTxOpts *bind.TransactOpts,
	signerPk *ecdsa.PrivateKey,
) (*RegistryContractClient, error) {
	registry, err := registry.NewRegistry(address, ethClient)
	if err != nil {
		return nil, err
	}
	return &RegistryContractClient{
		address:          address,
		registry:         registry,
		senderTxOpts:     senderTxOpts,
		txVerifier:       chain.NewTxVerifier(ethClient),
		signerPrivateKey: signerPk,
	}, nil

}

func (r *RegistryContractClient) RegisterVoter(nextRewardEpochId *big.Int, address string) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetry(func() (any, error) {
		err := r.sendRegisterVoter(nextRewardEpochId, address)
		if err != nil {
			return nil, errors.Wrap(err, "error sending register voter")
		}
		return nil, nil
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)
}

func (r *RegistryContractClient) sendRegisterVoter(nextRewardEpochId *big.Int, addressString string) error {
	epochId := uint32(nextRewardEpochId.Uint64())
	address := common.HexToAddress(addressString)
	signature, err := r.createSignature(epochId, address)
	if err != nil {
		return err
	}
	vrsSignature := registry.VoterRegistrySignature{
		R: [32]byte(signature[0:32]),
		S: [32]byte(signature[32:64]),
		V: signature[64] + 27,
	}
	r.senderTxOpts.GasPrice = big.NewInt(50 * 1e9)
	tx, err := r.registry.RegisterVoter(r.senderTxOpts, address, vrsSignature)
	if err != nil {
		return err
	}
	err = r.txVerifier.WaitUntilMined(r.senderTxOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Info("Voter %s registered for epoch %v", address, nextRewardEpochId)
	return nil
}

func (r *RegistryContractClient) createSignature(nextRewardEpochId uint32, address common.Address) ([]byte, error) {
	message, err := registratorArguments.Pack(nextRewardEpochId, address)
	if err != nil {
		return nil, err
	}
	messageHash := crypto.Keccak256(message)
	return crypto.Sign(accounts.TextHash(messageHash), r.signerPrivateKey)
}
