package chain

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

const (
	// default timeout for waiting for a tx to be mined.
	DefaultTxTimeout = 60 * time.Second
)

type TxVerifier struct {
	eth *ethclient.Client
}

func NewTxVerifier(eth *ethclient.Client) *TxVerifier {
	return &TxVerifier{eth: eth}
}

func (t TxVerifier) WaitUntilMined(from common.Address, tx *types.Transaction, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, t.eth, tx)
	if err != nil {
		return errors.Wrap(err, "bind.WaitMined")
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		reason, err := errorReason(ctx, t.eth, from, tx, receipt.BlockNumber)
		if err != nil {
			return err
		}
		return errors.Errorf("tx failed: %s", reason)
	}
	return nil
}

// Taken from: https://ethereum.stackexchange.com/questions/48383/how-to-retrieve-revert-reason-for-past-transactions
func errorReason(ctx context.Context, b ethereum.ContractCaller, from common.Address, tx *types.Transaction, blockNum *big.Int) (string, error) {
	msg := ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}
	res, err := b.CallContract(ctx, msg, blockNum)
	if err != nil {
		return "", errors.Wrap(err, "CallContract")
	}
	return unpackError(res)
}

var (
	errorSig     = []byte{0x08, 0xc3, 0x79, 0xa0} // Keccak256("Error(string)")[:4]
	abiString, _ = abi.NewType("string", "", nil)
)

func unpackError(result []byte) (string, error) {
	if !bytes.Equal(result[:4], errorSig) {
		return "<tx result not Error(string)>", errors.New("tx result not of type Error(string)")
	}
	vs, err := abi.Arguments{{Type: abiString}}.UnpackValues(result[4:])
	if err != nil {
		return "<invalid tx result>", errors.Wrap(err, "unpacking revert reason")
	}
	return vs[0].(string), nil
}

func TransactOptsFromPrivateKey(privateKey string, chainID int) (*bind.TransactOpts, error) {
	if len(privateKey) < 2 {
		return nil, errors.New("privateKey is too short")
	}

	if privateKey[:2] == "0x" {
		privateKey = privateKey[2:]
	}

	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "crypto.HexToECDSA")
	}

	opts, err := bind.NewKeyedTransactorWithChainID(
		pk, big.NewInt(int64(chainID)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	// bind.N
	return opts, nil
}

func SendRawTx(client ethclient.Client, privateKeyHex string, toAddress common.Address, data []byte) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	value := big.NewInt(0)     // in wei (1 eth)
	gasLimit := uint64(210000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	// TODO: NewTransaction is deprecated
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sending signed tx: ", signedTx.Hash().Hex())

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err)
	}

	rec, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		panic(err)
	}
	fmt.Println("Receipt: ", rec.Status)

	verifier := NewTxVerifier(&client)

	fmt.Println("Waiting for tx to be mined...", time.Now())
	err = verifier.WaitUntilMined(fromAddress, signedTx, 10*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tx mined ", time.Now())
}
