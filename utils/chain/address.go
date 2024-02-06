package chain

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

var (
	EmptyAddress = common.Address{}
)

func PrivateKeyToEthAddress(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("error casting public key to ECDSA")
	}

	return ethCrypto.PubkeyToAddress(*publicKeyECDSA), nil
}
