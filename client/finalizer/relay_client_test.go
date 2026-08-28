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
// (*big.Int, pointer-compared by maps) and digest.
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
		digest:        common.HexToHash("0xabc"), // non-zero — would have broken the old lookup
	}

	require.True(
		t,
		relayed[relayedKey{protocolID: item.protocolID, votingRoundID: item.votingRoundID}],
		"already-relayed lookup must match by (protocolID, votingRoundID) regardless of seed/digest",
	)

	miss := &queueItem{
		seed:          big.NewInt(1),
		votingRoundID: 43, // different round
		protocolID:    100,
		digest:        common.HexToHash("0xabc"),
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

	sentNonces    []uint64
	sentTo        []common.Address
	sendRemaining []time.Duration // per send: time left on its ctx deadline (0 = none)
	sendIdx       int
	nonceIdx      int
}

func (c *scriptedRelayClient) SendRawTx(ctx context.Context, _ *ecdsa.PrivateKey, nonce uint64, to common.Address, _ []byte, _ *config.Gas, _ time.Duration, _ bool) chain.SendResult {
	c.sentNonces = append(c.sentNonces, nonce)
	c.sentTo = append(c.sentTo, to)
	remaining := time.Duration(0)
	if dl, ok := ctx.Deadline(); ok {
		remaining = time.Until(dl)
	}
	c.sendRemaining = append(c.sendRemaining, remaining)
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
		retryDelay:    time.Millisecond,
	}
}

var (
	relayHash0 = common.HexToHash("0x11")
	relayHash1 = common.HexToHash("0x12")
)

// The new Relay reverts with a bare custom-error selector, which chain decodes to
// its signature. Both relays' reasons must stay non-fatal: the old one still serves
// every reward epoch before the cutover.
func TestRelayAlreadyRelayedReasons(t *testing.T) {
	for _, reason := range []string{"Already relayed", "AlreadyRelayed()"} {
		t.Run(reason+" on the current attempt", func(t *testing.T) {
			cc := &scriptedRelayClient{
				nonces:  []uint64{10},
				results: []chain.SendResult{{Err: errors.New("preparing tx: dry run: execution reverted: " + reason)}},
			}
			r := testRelayClient(t, cc)

			r.SubmitPayloads(context.Background(), relayContractAddress, make([]byte, 40), false, 1, 100)
			require.Len(t, cc.sentNonces, 1) // no retry, treated as success
		})
	}
}

// pins the per-attempt ctx cap — without it one slow attempt eats later attempts' slices
func TestRelaySendAttemptsAreDeadlineScoped(t *testing.T) {
	failing := func() *scriptedRelayClient {
		return &scriptedRelayClient{
			nonces:  []uint64{10},
			results: []chain.SendResult{{Broadcast: false, Err: errors.New("send failed")}},
		}
	}

	t.Run("unbounded path caps each attempt at the default timeout", func(t *testing.T) {
		cc := failing()
		r := testRelayClient(t, cc)

		r.SubmitPayloads(context.Background(), relayContractAddress, make([]byte, 40), false, 1, 100)

		require.NotEmpty(t, cc.sendRemaining)
		for _, remaining := range cc.sendRemaining {
			require.Greater(t, remaining, chain.DefaultTxTimeout-5*time.Second)
			require.LessOrEqual(t, remaining, chain.DefaultTxTimeout)
		}
	})

	t.Run("bounded path caps each attempt at its perAttempt slice", func(t *testing.T) {
		cc := failing()
		r := testRelayClient(t, cc)
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		r.SubmitPayloads(ctx, relayContractAddress, make([]byte, 40), false, 1, 100)

		require.NotEmpty(t, cc.sendRemaining)
		// perAttempt = (remaining − 2×retryDelay)/3 ≈ 16.7s, well below the 50s parent
		for _, remaining := range cc.sendRemaining {
			require.Greater(t, remaining, 10*time.Second)
			require.Less(t, remaining, 18*time.Second)
		}
	})
}

// A prior broadcast mined but reverted with an allowed reason — either relay's
// spelling — must reconcile a subsequent "nonce too low" as success (no duplicate).
func TestRelayReconcilesMinedRevertedAllowed(t *testing.T) {
	for _, reason := range []string{"Already relayed", "AlreadyRelayed()"} {
		t.Run(reason, func(t *testing.T) {
			cc := &scriptedRelayClient{
				nonces: []uint64{10},
				results: []chain.SendResult{
					{Hash: relayHash0, Broadcast: true, Err: context.DeadlineExceeded}, // post-broadcast timeout
					{Hash: relayHash1, Broadcast: false, Err: errors.New("nonce too low")},
				},
				receipts: map[common.Hash]*types.Receipt{relayHash0: {Status: types.ReceiptStatusFailed}},
				reverts:  map[common.Hash]string{relayHash0: reason},
			}
			r := testRelayClient(t, cc)

			r.SubmitPayloads(context.Background(), relayContractAddress, make([]byte, 40), false, 1, 100)
			require.Equal(t, []uint64{10, 10}, cc.sentNonces) // reconciled, nonce not bumped
		})
	}
}

// A mined-but-reverted relay tx with a fatal (non-"Already relayed") reason is
// deterministic, so it must be terminal: no retry, no nonce refresh.
func TestRelayMinedRevertedFatalIsTerminal(t *testing.T) {
	cc := &scriptedRelayClient{
		nonces:  []uint64{10, 11}, // 11 would appear if the nonce were refreshed
		results: []chain.SendResult{{Hash: relayHash0, Broadcast: true, Mined: true, Err: errors.New("tx mined but reverted: Not enough weight")}},
	}
	r := testRelayClient(t, cc)

	r.SubmitPayloads(context.Background(), relayContractAddress, make([]byte, 40), false, 1, 100)
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

	r.SubmitPayloads(context.Background(), relayContractAddress, make([]byte, 40), false, 1, 100)
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

	r.SubmitPayloads(context.Background(), relayContractAddress, make([]byte, 40), false, 1, 100)
	require.Equal(t, []uint64{10, 10}, cc.sentNonces) // reconciled to terminal: no bump, no attempt 3
}

// Reconciliation that cannot find the prior broadcast's receipt refreshes the
// nonce and resends instead of retrying the dead nonce (a duplicate reverts
// non-fatally with "Already relayed").
func TestRelayNonceTooLowUndeterminedRefetchesNonce(t *testing.T) {
	cc := &scriptedRelayClient{
		nonces: []uint64{10, 11},
		results: []chain.SendResult{
			{Hash: relayHash0, Broadcast: true, Err: context.DeadlineExceeded},     // broadcast at 10, times out
			{Hash: relayHash1, Broadcast: false, Err: errors.New("nonce too low")}, // consumed, relayHash0 not found
			{Hash: relayHash1, Broadcast: true, Err: nil},                          // succeeds at the refreshed nonce
		},
	}
	r := testRelayClient(t, cc)

	r.SubmitPayloads(context.Background(), relayContractAddress, make([]byte, 40), false, 1, 100)
	require.Equal(t, []uint64{10, 10, 11}, cc.sentNonces)
}

// hangingRelayClient blocks every send until ctx is done, emulating a stuck tx.
type hangingRelayClient struct {
	scriptedRelayClient
}

func (c *hangingRelayClient) SendRawTx(ctx context.Context, _ *ecdsa.PrivateKey, nonce uint64, _ common.Address, _ []byte, _ *config.Gas, _ time.Duration, _ bool) chain.SendResult {
	c.sentNonces = append(c.sentNonces, nonce)
	<-ctx.Done()
	return chain.SendResult{Hash: relayHash0, Broadcast: true, Err: ctx.Err()}
}

// The queue processor bounds each item's send with a deadline ctx (see
// processItemBounded); SubmitPayloads must honor it instead of grinding the
// full retry budget.
func TestSubmitPayloadsHonorsContextDeadline(t *testing.T) {
	cc := &hangingRelayClient{scriptedRelayClient{nonces: []uint64{10}}}
	r := testRelayClient(t, cc)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	start := time.Now()
	r.SubmitPayloads(ctx, relayContractAddress, make([]byte, 40), false, 1, 100)
	require.Less(t, time.Since(start), 5*time.Second) // deadline abort, not the ~100s budget
	require.LessOrEqual(t, len(cc.sentNonces), 2)
}

// A failing nonce fetch with a canceled ctx must abort promptly without sending.
func TestRelayNonceFetchAbortsOnCanceledCtx(t *testing.T) {
	cc := &scriptedRelayClient{nonceFail: 100, nonces: []uint64{10}}
	r := testRelayClient(t, cc)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // canceled: the retry loop must bail instead of grinding the budget
	r.SubmitPayloads(ctx, relayContractAddress, make([]byte, 40), false, 1, 100)
	require.Empty(t, cc.sentNonces) // never reached the send loop
}
