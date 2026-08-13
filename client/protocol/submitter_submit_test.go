package protocol

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"testing"
	"time"

	clientConfig "github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/utils/chain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// scriptedChainClient returns a pre-scripted SendResult per SendRawTx attempt
// (clamped to the last), a pre-scripted nonce per Nonce call (clamped), and
// receipts/reverts keyed by tx hash. It records the nonce seen on each send so
// tests can assert nonce reuse vs. refetch.
type scriptedChainClient struct {
	results    []chain.SendResult
	nonces     []uint64
	receipts   map[common.Hash]*types.Receipt
	reverts    map[common.Hash]string
	receiptErr map[common.Hash]bool // hashes whose Receipt lookup fails (RPC error)

	sentNonces []uint64
	sendIdx    int
	nonceIdx   int
}

func (c *scriptedChainClient) SendRawTx(_ context.Context, _ *ecdsa.PrivateKey, nonce uint64, _ common.Address, _ []byte, _ *clientConfig.Gas, _ time.Duration, _ bool) chain.SendResult {
	c.sentNonces = append(c.sentNonces, nonce)
	r := c.results[min(c.sendIdx, len(c.results)-1)]
	c.sendIdx++
	return r
}

func (c *scriptedChainClient) Nonce(_ context.Context, _ *ecdsa.PrivateKey, _ time.Duration) (uint64, error) {
	n := c.nonces[min(c.nonceIdx, len(c.nonces)-1)]
	c.nonceIdx++
	return n, nil
}

func (c *scriptedChainClient) Receipt(_ context.Context, hash common.Hash, _ time.Duration) (*types.Receipt, error) {
	if c.receiptErr[hash] {
		return nil, errors.New("rpc down")
	}
	return c.receipts[hash], nil
}

func (c *scriptedChainClient) RevertReason(_ context.Context, _ common.Address, hash common.Hash, _ time.Duration) (string, error) {
	return c.reverts[hash], nil
}

func testSubmitterBase(t *testing.T, cc chain.Client, retries int) *SubmitterBase {
	t.Helper()
	pk, err := crypto.HexToECDSA(testPrivateKeyHex)
	require.NoError(t, err)
	return &SubmitterBase{
		chainClient:      cc,
		gasConfig:        &clientConfig.Gas{TxType: 0, GasPriceFixed: common.Big0},
		protocolContext:  &protocolContext{submitContractAddress: common.HexToAddress(submitContractAddress)},
		name:             "test",
		submitRetries:    retries,
		submitTimeout:    time.Second,
		retryDelay:       time.Millisecond,
		submitPrivateKey: pk,
	}
}

var (
	hash0 = common.HexToHash("0x01")
	hash1 = common.HexToHash("0x02")
)

// A post-broadcast timeout followed by "nonce too low" whose earlier broadcast
// was mined must be reported as success WITHOUT bumping the nonce (the tx that
// timed out is the one that got mined).
func TestSubmitPostBroadcastTimeoutThenNonceTooLowAccepted(t *testing.T) {
	cc := &scriptedChainClient{
		nonces: []uint64{10},
		results: []chain.SendResult{
			{Hash: hash0, Broadcast: true, Err: context.DeadlineExceeded},     // post-broadcast timeout
			{Hash: hash1, Broadcast: false, Err: errors.New("nonce too low")}, // rejected: hash0 won the nonce
		},
		receipts: map[common.Hash]*types.Receipt{hash0: {Status: types.ReceiptStatusSuccessful}},
	}
	base := testSubmitterBase(t, cc, 3)

	require.True(t, base.submit(context.Background(), 100, make([]byte, 40)))
	// Nonce reused on the second attempt (post-broadcast timeout keeps the nonce).
	require.Equal(t, []uint64{10, 10}, cc.sentNonces)
}

// "nonce too low" where none of our broadcasts made it on chain means a foreign
// tx took the nonce: bump the nonce and resend, then succeed.
func TestSubmitNonceTooLowForeignBumpsNonce(t *testing.T) {
	cc := &scriptedChainClient{
		nonces: []uint64{10, 11}, // initial fetch, then refetch after bump
		results: []chain.SendResult{
			{Hash: hash0, Broadcast: false, Err: errors.New("nonce too low")}, // stale nonce, not broadcast
			{Hash: hash1, Broadcast: true, Err: nil},                          // succeeds at bumped nonce
		},
		receipts: map[common.Hash]*types.Receipt{}, // nothing of ours mined
	}
	base := testSubmitterBase(t, cc, 3)

	require.True(t, base.submit(context.Background(), 100, make([]byte, 40)))
	require.Equal(t, []uint64{10, 11}, cc.sentNonces) // second attempt used the bumped nonce
}

// A broadcast hash whose receipt reads "not found" leaves the outcome unknown:
// the nonce is refreshed and the payload resent (a duplicate submit is
// idempotent; retrying a nonce consumed by another tx would drop the round).
func TestSubmitNonceTooLowNotFoundRefetchesNonce(t *testing.T) {
	cc := &scriptedChainClient{
		nonces: []uint64{10, 11}, // initial fetch, then refresh after Undetermined
		results: []chain.SendResult{
			{Hash: hash0, Broadcast: true, Err: context.DeadlineExceeded},     // broadcast at 10, times out
			{Hash: hash1, Broadcast: false, Err: errors.New("nonce too low")}, // nonce consumed, hash0 not found
			{Hash: hash1, Broadcast: true, Err: nil},                          // succeeds at the refreshed nonce
		},
		receipts: map[common.Hash]*types.Receipt{},
	}
	base := testSubmitterBase(t, cc, 3)

	require.True(t, base.submit(context.Background(), 100, make([]byte, 40)))
	require.Equal(t, []uint64{10, 10, 11}, cc.sentNonces) // resent at the refreshed nonce
}

// A pre-broadcast timeout never sent the tx, so the nonce must be refetched
// (contrast with the post-broadcast case which reuses it).
func TestSubmitPreBroadcastTimeoutRefetchesNonce(t *testing.T) {
	cc := &scriptedChainClient{
		nonces: []uint64{10, 11},
		results: []chain.SendResult{
			{Broadcast: false, Err: context.DeadlineExceeded}, // pre-broadcast timeout, never signed
			{Hash: hash1, Broadcast: true, Err: nil},
		},
	}
	base := testSubmitterBase(t, cc, 3)

	require.True(t, base.submit(context.Background(), 100, make([]byte, 40)))
	require.Equal(t, []uint64{10, 11}, cc.sentNonces) // refetched, not reused
}

// When reconciliation can't determine whether our broadcast landed (RPC error),
// the nonce is refreshed and the payload resent rather than retrying a possibly
// dead nonce until the budget is exhausted.
func TestSubmitNonceTooLowUndeterminedRefetchesNonce(t *testing.T) {
	cc := &scriptedChainClient{
		nonces: []uint64{10, 11},
		results: []chain.SendResult{
			{Hash: hash0, Broadcast: true, Err: context.DeadlineExceeded},     // post-broadcast timeout, keep nonce
			{Hash: hash1, Broadcast: false, Err: errors.New("nonce too low")}, // consumed, hash0's fate unknown
			{Hash: hash1, Broadcast: true, Err: nil},                          // succeeds at the refreshed nonce
		},
		receiptErr: map[common.Hash]bool{hash0: true},
	}
	base := testSubmitterBase(t, cc, 3)

	require.True(t, base.submit(context.Background(), 100, make([]byte, 40)))
	require.Equal(t, []uint64{10, 10, 11}, cc.sentNonces) // resent at the refreshed nonce
}

// A pre-broadcast failure on a retry, when a prior attempt already broadcast,
// must keep the nonce (the outstanding tx is at it) rather than advancing it and
// resending a duplicate at a new nonce.
func TestSubmitPreBroadcastFailureAfterBroadcastKeepsNonce(t *testing.T) {
	cc := &scriptedChainClient{
		nonces: []uint64{10, 999}, // 999 would appear if the nonce were wrongly refetched/bumped
		results: []chain.SendResult{
			{Hash: hash0, Broadcast: true, Err: context.DeadlineExceeded},      // broadcast at 10, keep nonce
			{Broadcast: false, Err: errors.New("preparing tx: baseFee: boom")}, // pre-broadcast failure, nothing new sent
			{Hash: hash1, Broadcast: true, Err: nil},                           // finally confirms at 10
		},
	}
	base := testSubmitterBase(t, cc, 3)

	require.True(t, base.submit(context.Background(), 100, make([]byte, 40)))
	require.Equal(t, []uint64{10, 10, 10}, cc.sentNonces) // nonce held across the pre-broadcast failure
}

// Exhausting retries on repeated failures reports failure.
func TestSubmitExhaustsRetries(t *testing.T) {
	cc := &scriptedChainClient{
		nonces:  []uint64{10},
		results: []chain.SendResult{{Broadcast: true, Err: context.DeadlineExceeded}},
	}
	base := testSubmitterBase(t, cc, 2)

	require.False(t, base.submit(context.Background(), 100, make([]byte, 40)))
	require.Len(t, cc.sentNonces, 2) // exactly submitRetries attempts
}
