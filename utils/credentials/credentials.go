package credentials

import (
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

func TransactOptsFromPrivateKey(pk *ecdsa.PrivateKey, chainID int) (*bind.TransactOpts, error) {
	opts, _, err := CredentialsFromPrivateKey(pk, chainID)
	return opts, err
}

func CredentialsFromPrivateKey(pk *ecdsa.PrivateKey, chainID int) (*bind.TransactOpts, *ecdsa.PrivateKey, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		pk, big.NewInt(int64(chainID)),
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	// bind.N
	return opts, pk, nil
}

func PrivateKeyFromHex(privateKey string) (*ecdsa.PrivateKey, error) {
	if len(privateKey) < 2 {
		return nil, errors.New("privateKey is too short")
	}

	privateKey = strings.TrimPrefix(privateKey, "0x")

	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "crypto.HexToECDSA")
	}
	return pk, nil
}
