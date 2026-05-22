package finalizer

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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
