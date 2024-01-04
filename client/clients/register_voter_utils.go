package clients

import (
	"crypto/ecdsa"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/registry"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

const (
	maxRetries = 10
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

type voterRegistrator struct {
	txOpts        *bind.TransactOpts
	voterRegistry *registry.Registry
	privateKey    *ecdsa.PrivateKey
	txVerifier    *chain.TxVerifier
}

func NewVoterRegistrator(
	voterRegistry *registry.Registry,
	txOpts *bind.TransactOpts,
	txVerifier *chain.TxVerifier,
	privateKey *ecdsa.PrivateKey,
) *voterRegistrator {
	return &voterRegistrator{
		voterRegistry: voterRegistry,
		txOpts:        txOpts,
		privateKey:    privateKey,
		txVerifier:    txVerifier,
	}
}

func (v *voterRegistrator) RegisterVoter(nextRewardEpochId *big.Int, address string) <-chan bool {
	out := make(chan bool)
	go func() {
		for retry := 0; retry < maxRetries; retry++ {
			// Todo, check if voter is already registered
			err := v.SendRegisterVoter(nextRewardEpochId, address)
			if err != nil {
				logger.Error("SendRegisterVoter: %w", err)
			} else {
				out <- true
				return
			}
		}
		logger.Error("SendRegisterVoter: max retries reached")
		out <- false
	}()
	return out
}

func (v *voterRegistrator) createSignature(nextRewardEpochId *big.Int, address string) ([]byte, error) {
	message, err := registratorArguments.Pack(nextRewardEpochId, address)
	if err != nil {
		return nil, err
	}
	messageHash := crypto.Keccak256(message)
	return crypto.Sign(messageHash, v.privateKey)
}

func (v *voterRegistrator) SendRegisterVoter(nextRewardEpochId *big.Int, address string) error {
	signature, err := v.createSignature(nextRewardEpochId, address)
	if err != nil {
		return errors.Wrap(err, "createSignature")
	}
	vrsSignature := registry.VoterRegistrySignature{
		V: signature[0],
		R: [32]byte(signature[1:33]),
		S: [32]byte(signature[33:65]),
	}
	tx, err := v.voterRegistry.RegisterVoter(v.txOpts, v.txOpts.From, vrsSignature)
	if err != nil {
		return err
	}
	err = v.txVerifier.WaitUntilMined(v.txOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Info("Voter %s registered for epoch %v", address, nextRewardEpochId)
	return nil
}
