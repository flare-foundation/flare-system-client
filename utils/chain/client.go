package chain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client interface {
	// SendRawTx signs, broadcasts and waits for a transaction; SendResult
	// classifies the outcome (pre-broadcast failure vs post-broadcast timeout).
	SendRawTx(ctx context.Context, privateKey *ecdsa.PrivateKey, nonce uint64, to common.Address, payload []byte, gasConfig *config.Gas, timeout time.Duration, dryRun bool) SendResult

	// Nonce returns the latest-block nonce of the privateKey's address.
	Nonce(ctx context.Context, privateKey *ecdsa.PrivateKey, timeout time.Duration) (uint64, error)

	// Receipt returns the tx's receipt, or nil (nil error) if not yet mined.
	Receipt(ctx context.Context, hash common.Hash, timeout time.Duration) (*types.Receipt, error)

	// RevertReason returns the revert reason of a mined-but-reverted tx.
	RevertReason(ctx context.Context, from common.Address, hash common.Hash, timeout time.Duration) (string, error)
}

type ClientImpl struct {
	EthClient *ethclient.Client
	ChainID   *big.Int // EIP-155 chain id for tx signing, from config — immutable, so never refetched per send
}

// NewClientImpl returns a ClientImpl signing txs with the EIP-155 chainID.
func NewClientImpl(ethClient *ethclient.Client, chainID int64) ClientImpl {
	return ClientImpl{EthClient: ethClient, ChainID: big.NewInt(chainID)}
}

// SendRawTx sends a transaction with payload signed by privateKey to to address.
func (c ClientImpl) SendRawTx(ctx context.Context, privateKey *ecdsa.PrivateKey, nonce uint64, to common.Address, payload []byte, gasConfig *config.Gas, timeout time.Duration, dryRun bool) SendResult {
	return SendRawTx(ctx, c.EthClient, privateKey, c.ChainID, nonce, to, payload, dryRun, gasConfig, timeout)
}

// Nonce returns the nonce of the address corresponding to the privateKey from the latest known block.
func (c ClientImpl) Nonce(ctx context.Context, privateKey *ecdsa.PrivateKey, timeout time.Duration) (uint64, error) {
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonceCtx, cancelFunc := context.WithTimeout(ctx, timeout)
	nonce, err := c.EthClient.NonceAt(nonceCtx, address, nil)
	cancelFunc()
	if err != nil {
		return 0, err
	}

	return nonce, nil
}

// Receipt returns the tx's receipt. A pending/unknown tx returns nil, nil; an
// error is returned only on an actual RPC failure.
func (c ClientImpl) Receipt(ctx context.Context, hash common.Hash, timeout time.Duration) (*types.Receipt, error) {
	receiptCtx, cancelFunc := context.WithTimeout(ctx, timeout)
	defer cancelFunc()

	receipt, err := c.EthClient.TransactionReceipt(receiptCtx, hash)
	if errors.Is(err, ethereum.NotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

// RevertReason returns the revert reason of a mined tx by replaying it against
// the state of the block it was mined in.
func (c ClientImpl) RevertReason(ctx context.Context, from common.Address, hash common.Hash, timeout time.Duration) (string, error) {
	reasonCtx, cancelFunc := context.WithTimeout(ctx, timeout)
	defer cancelFunc()

	tx, isPending, err := c.EthClient.TransactionByHash(reasonCtx, hash)
	if err != nil {
		return "", err
	}
	if isPending {
		return "", errors.New("transaction still pending")
	}

	receipt, err := c.EthClient.TransactionReceipt(reasonCtx, hash)
	if err != nil {
		return "", err
	}

	return errorReason(reasonCtx, c.EthClient, from, tx, receipt.BlockNumber)
}
