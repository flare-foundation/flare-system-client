package chain

import (
	"crypto/ecdsa"
	"errors"

	"github.com/flare-foundation/flare-system-client/client/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client interface {
	SendRawTx(privateKey *ecdsa.PrivateKey, to common.Address, payload []byte, gasConfig *config.Gas) error
}

type ClientImpl struct {
	EthClient *ethclient.Client
}

// SendRawTx sends a transaction with payload signed by privateKey to to address.
func (c ClientImpl) SendRawTx(privateKey *ecdsa.PrivateKey, to common.Address, payload []byte, gasConfig *config.Gas) error {
	switch gasConfig.TxType {
	case 0:
		return SendRawTx(c.EthClient, privateKey, to, payload, true, gasConfig)
	case 2:
		return SendRawType2Tx(c.EthClient, privateKey, to, payload, true, gasConfig)
	default:
		return errors.New("unsupported tx type: set TxType to 0 or 2")
	}
}
