package chain

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client interface {
	SendRawTx(privateKey *ecdsa.PrivateKey, nonce uint64, to common.Address, payload []byte, gasConfig *config.Gas, timeout time.Duration, dryRun bool) error

	Nonce(privateKey *ecdsa.PrivateKey, timeout time.Duration) (uint64, error)
}

type ClientImpl struct {
	EthClient *ethclient.Client
}

// SendRawTx sends a transaction with payload signed by privateKey to to address.
func (c ClientImpl) SendRawTx(privateKey *ecdsa.PrivateKey, nonce uint64, to common.Address, payload []byte, gasConfig *config.Gas, timeout time.Duration, dryRun bool) error {
	return SendRawTx(c.EthClient, privateKey, nonce, to, payload, dryRun, gasConfig, timeout)
}

// Nonce returns of the address corresponding to the privateKey from the latest known block.
func (c ClientImpl) Nonce(privateKey *ecdsa.PrivateKey, timeout time.Duration) (uint64, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return 0, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	nonce, err := c.EthClient.NonceAt(ctx, address, nil)
	cancelFunc()
	if err != nil {
		return 0, err
	}

	return nonce, nil
}
