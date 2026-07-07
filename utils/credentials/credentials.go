package credentials

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
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
		return nil, nil, fmt.Errorf("bind.NewKeyedTransactorWithChainID: %w", err)
	}
	// bind.N
	return opts, pk, nil
}

func PrivateKeyFromHex(privateKey string) (*ecdsa.PrivateKey, error) {
	privateKey = strings.TrimPrefix(privateKey, "0x")

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("crypto.HexToECDSA: %w", err)
	}
	return privKey, nil
}
