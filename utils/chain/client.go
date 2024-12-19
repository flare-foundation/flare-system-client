package chain

import (
	"crypto/ecdsa"
	"errors"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client interface {
	SendRawTx(privateKey *ecdsa.PrivateKey, to common.Address, payload []byte, gasConfig *config.Gas, timeout time.Duration, dryRun bool) error
}

type ClientImpl struct {
	EthClient *ethclient.Client
}

// SendRawTx sends a transaction with payload signed by privateKey to to address.
func (c ClientImpl) SendRawTx(privateKey *ecdsa.PrivateKey, to common.Address, payload []byte, gasConfig *config.Gas, timeout time.Duration, dryRun bool) error {
	switch gasConfig.TxType {
	case 0:
		return SendRawTx(c.EthClient, privateKey, to, payload, dryRun, gasConfig, timeout)
	case 2:
		return SendRawType2Tx(c.EthClient, privateKey, to, payload, dryRun, gasConfig, timeout)
	default:
		return errors.New("unsupported tx type: set TxType to 0 or 2")
	}
}
