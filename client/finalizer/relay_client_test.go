package finalizer

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/utils/chain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// TestRelayedKey_LookupMatchesWithUnrelatedQueueItemFields pins the
// fix for the duplicate-suppression bug in processDelayedQueue: the
// lookup key built from a ProtocolMessageRelayed event must match the
// lookup performed against a queueItem regardless of that item's seed
// (*big.Int, pointer-compared by maps) and msgHash.
//
// Pre-fix, ProtocolMessageRelayed used queueItem as the map key with
// only protocolID + votingRoundID populated, so the lookup with a
// fully-populated queueItem never matched and every delayed item was
// re-sent (the relay contract returns the non-fatal "Already relayed"
// error, so the bug only burned gas).
func TestRelayedKey_LookupMatchesWithUnrelatedQueueItemFields(t *testing.T) {
	relayed := map[relayedKey]bool{
		{protocolID: 100, votingRoundID: 42}: true,
		{protocolID: 200, votingRoundID: 99}: true,
	}

	item := &queueItem{
		seed:          big.NewInt(0xdeadbeef), // non-nil pointer — would have broken the old lookup
		votingRoundID: 42,
		protocolID:    100,
		msgHash:       common.HexToHash("0xabc"), // non-zero — would have broken the old lookup
	}

	require.True(
		t,
		relayed[relayedKey{protocolID: item.protocolID, votingRoundID: item.votingRoundID}],
		"already-relayed lookup must match by (protocolID, votingRoundID) regardless of seed/msgHash",
	)

	miss := &queueItem{
		seed:          big.NewInt(1),
		votingRoundID: 43, // different round
		protocolID:    100,
		msgHash:       common.HexToHash("0xabc"),
	}
	require.False(
		t,
		relayed[relayedKey{protocolID: miss.protocolID, votingRoundID: miss.votingRoundID}],
		"different (protocolID, votingRoundID) must not match",
	)
}

// TestRelayedKey_TwoNonNilSeedsDoNotCollide makes the failure mode of
// the pre-fix code explicit: under the old shape, two queueItems
// holding distinct *big.Int seeds compared as DIFFERENT map keys even
// when their (protocolID, votingRoundID) were equal — that's exactly
// why the producer-built map (seed=nil) never matched the
// consumer-side lookup (seed != nil).
func TestRelayedKey_TwoNonNilSeedsDoNotCollide(t *testing.T) {
	a := queueItem{seed: big.NewInt(1), votingRoundID: 1, protocolID: 1}
	b := queueItem{seed: big.NewInt(1), votingRoundID: 1, protocolID: 1}

	m := map[queueItem]bool{a: true}
	require.False(t, m[b], "*big.Int is pointer-compared in map keys — distinct allocations of equal values DO NOT match")

	// And the fix: relayedKey excludes the pointer field entirely.
	ra := relayedKey{protocolID: 1, votingRoundID: 1}
	rb := relayedKey{protocolID: 1, votingRoundID: 1}
	rm := map[relayedKey]bool{ra: true}
	require.True(t, rm[rb], "relayedKey is pure value type — equal values match")
}

type scriptedRelayClient struct {
	results   []chain.SendResult
	nonces    []uint64
	receipts  map[common.Hash]*types.Receipt
	reverts   map[common.Hash]string
	nonceFail int // number of leading Nonce calls that fail before succeeding

	sentNonces []uint64
	sendIdx    int
	nonceIdx   int
}

func (c *scriptedRelayClient) SendRawTx(_ context.Context, _ *ecdsa.PrivateKey, nonce uint64, _ common.Address, _ []byte, _ *config.Gas, _ time.Duration, _ bool) chain.SendResult {
	c.sentNonces = append(c.sentNonces, nonce)
	r := c.results[min(c.sendIdx, len(c.results)-1)]
	c.sendIdx++
	return r
}
func (c *scriptedRelayClient) Nonce(context.Context, *ecdsa.PrivateKey, time.Duration) (uint64, error) {
	i := c.nonceIdx
	c.nonceIdx++
	if i < c.nonceFail {
		return 0, errors.New("rpc down")
	}
	return c.nonces[min(i-c.nonceFail, len(c.nonces)-1)], nil
}
func (c *scriptedRelayClient) Receipt(_ context.Context, h common.Hash, _ time.Duration) (*types.Receipt, error) {
	return c.receipts[h], nil
}
func (c *scriptedRelayClient) RevertReason(_ context.Context, _ common.Address, h common.Hash, _ time.Duration) (string, error) {
	return c.reverts[h], nil
}

func testRelayClient(t *testing.T, cc chain.Client) *relayContractClient {
	t.Helper()
	pk, err := crypto.HexToECDSA(testPrivateKeyHex)
	require.NoError(t, err)
	return &relayContractClient{
		chainClient:   cc,
		gasConfig:     &config.Gas{TxType: 0, GasPriceFixed: common.Big0},
		privateKey:    pk,
		address:       relayContractAddress,
		senderAddress: crypto.PubkeyToAddress(pk.PublicKey),
	}
}

var (
	relayHash0 = common.HexToHash("0x11")
	relayHash1 = common.HexToHash("0x12")
)

// A relay tx reverting with "Already relayed" on the current attempt is a
// non-fatal success (someone else finalized the round).
func TestRelayAlreadyRelayedIsNonFatal(t *testing.T) {
	cc := &scriptedRelayClient{
		nonces:  []uint64{10},
		results: []chain.SendResult{{Broadcast: false, Err: errors.New("execution reverted: Already relayed")}},
	}
	r := testRelayClient(t, cc)

	r.SubmitPayloads(context.Background(), make([]byte, 40), false, 1)
	require.Len(t, cc.sentNonces, 1) // no retry, treated as success
}

// A prior broadcast mined but reverted with the allowed "Already relayed" reason
// must reconcile a subsequent "nonce too low" as success (no duplicate).
func TestRelayReconcilesMinedRevertedAllowed(t *testing.T) {
	cc := &scriptedRelayClient{
		nonces: []uint64{10},
		results: []chain.SendResult{
			{Hash: relayHash0, Broadcast: true, Err: context.DeadlineExceeded}, // post-broadcast timeout
			{Hash: relayHash1, Broadcast: false, Err: errors.New("nonce too low")},
		},
		receipts: map[common.Hash]*types.Receipt{relayHash0: {Status: types.ReceiptStatusFailed}},
		reverts:  map[common.Hash]string{relayHash0: "Already relayed"},
	}
	r := testRelayClient(t, cc)

	r.SubmitPayloads(context.Background(), make([]byte, 40), false, 1)
	require.Equal(t, []uint64{10, 10}, cc.sentNonces) // reconciled, nonce not bumped
}

// A mined-but-reverted relay tx with a fatal (non-"Already relayed") reason is
// deterministic, so it must be terminal: no retry, no nonce refresh.
func TestRelayMinedRevertedFatalIsTerminal(t *testing.T) {
	cc := &scriptedRelayClient{
		nonces:  []uint64{10, 11}, // 11 would appear if the nonce were refreshed
		results: []chain.SendResult{{Hash: relayHash0, Broadcast: true, Mined: true, Err: errors.New("tx mined but reverted: Not enough weight")}},
	}
	r := testRelayClient(t, cc)

	r.SubmitPayloads(context.Background(), make([]byte, 40), false, 1)
	require.Equal(t, []uint64{10}, cc.sentNonces) // exactly one send: no retry, no bump
}

// A mined-but-reverted relay tx whose reason is the allowed "Already relayed" is
// treated as a non-fatal success (not the terminal-revert case), so no retry.
func TestRelayMinedRevertedAllowedIsSuccess(t *testing.T) {
	cc := &scriptedRelayClient{
		nonces:  []uint64{10},
		results: []chain.SendResult{{Hash: relayHash0, Broadcast: true, Mined: true, Err: errors.New("tx mined but reverted: Already relayed")}},
	}
	r := testRelayClient(t, cc)

	r.SubmitPayloads(context.Background(), make([]byte, 40), false, 1)
	require.Len(t, cc.sentNonces, 1) // non-fatal: success, no retry
}

// A prior broadcast that timed out and then mined-reverted with a fatal reason is
// discovered via nonce-too-low reconciliation. It must be terminal (no nonce bump,
// no re-broadcast of the deterministically-reverting input).
func TestRelayReconciledMinedRevertedFatalIsTerminal(t *testing.T) {
	cc := &scriptedRelayClient{
		nonces: []uint64{10, 11}, // 11 would appear only if the nonce were bumped
		results: []chain.SendResult{
			{Hash: relayHash0, Broadcast: true, Err: context.DeadlineExceeded},     // broadcast at 10, times out
			{Hash: relayHash1, Broadcast: false, Err: errors.New("nonce too low")}, // resend at 10 rejected
		},
		receipts: map[common.Hash]*types.Receipt{relayHash0: {Status: types.ReceiptStatusFailed}},
		reverts:  map[common.Hash]string{relayHash0: "Not enough weight"}, // fatal (not "Already relayed")
	}
	r := testRelayClient(t, cc)

	r.SubmitPayloads(context.Background(), make([]byte, 40), false, 1)
	require.Equal(t, []uint64{10, 10}, cc.sentNonces) // reconciled to terminal: no bump, no attempt 3
}

// A failing nonce fetch with a canceled ctx must abort promptly without sending.
func TestRelayNonceFetchAbortsOnCanceledCtx(t *testing.T) {
	cc := &scriptedRelayClient{nonceFail: 100, nonces: []uint64{10}}
	r := testRelayClient(t, cc)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // canceled: the retry loop must bail instead of grinding the budget
	r.SubmitPayloads(ctx, make([]byte, 40), false, 1)
	require.Empty(t, cc.sentNonces) // never reached the send loop
}
