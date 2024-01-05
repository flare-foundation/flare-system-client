package clients

import (
	"crypto/ecdsa"
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
	uint256Ty, _ := abi.NewType("uint256", "uint256", nil)
	addressTy, _ := abi.NewType("address", "adddress", nil)
	registratorArguments = abi.Arguments{
		{ // nextRewardEpochId
			Type: uint256Ty,
		},
		{ // address
			Type: addressTy,
		},
	}
}

type RegistryContractClient struct {
	address    common.Address
	registry   *registry.Registry
	txOpts     *bind.TransactOpts
	txVerifier *chain.TxVerifier
	privateKey *ecdsa.PrivateKey
}

func NewRegistryContractClient(
	chainID int,
	ethClient *ethclient.Client,
	address common.Address,
	privateKeyString string,
) (*RegistryContractClient, error) {
	txOpts, privateKey, err := chain.CredentialsFromPrivateKey(privateKeyString, chainID)
	if err != nil {
		return nil, err
	}
	registry, err := registry.NewRegistry(address, ethClient)
	if err != nil {
		return nil, err
	}
	return &RegistryContractClient{
		address:    address,
		registry:   registry,
		txOpts:     txOpts,
		txVerifier: chain.NewTxVerifier(ethClient),
		privateKey: privateKey,
	}, nil

}

func (r *RegistryContractClient) RegisterVoter(nextRewardEpochId *big.Int, address string) <-chan ExecuteStatus {
	return ExecuteWithRetry(func() error {
		err := r.sendRegisterVoter(nextRewardEpochId, address)
		if err != nil {
			return errors.Wrap(err, "error sending register voter")
		}
		return nil
	}, MaxTxSendRetries)
}

func (r *RegistryContractClient) sendRegisterVoter(nextRewardEpochId *big.Int, address string) error {
	signature, err := r.createSignature(nextRewardEpochId, address)
	if err != nil {
		return err
	}
	vrsSignature := registry.VoterRegistrySignature{
		V: signature[0],
		R: [32]byte(signature[1:33]),
		S: [32]byte(signature[33:65]),
	}
	tx, err := r.registry.RegisterVoter(r.txOpts, r.txOpts.From, vrsSignature)
	if err != nil {
		return err
	}
	err = r.txVerifier.WaitUntilMined(r.txOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Info("Voter %s registered for epoch %v", address, nextRewardEpochId)
	return nil
}

func (r *RegistryContractClient) createSignature(nextRewardEpochId *big.Int, address string) ([]byte, error) {
	message, err := registratorArguments.Pack(nextRewardEpochId, address)
	if err != nil {
		return nil, err
	}
	messageHash := crypto.Keccak256(message)
	return crypto.Sign(accounts.TextHash(messageHash), r.privateKey)
}
