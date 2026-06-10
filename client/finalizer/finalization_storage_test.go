package finalizer

import (
	"crypto/ecdsa"
	"crypto/rand"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/flare-foundation/go-flare-common/pkg/policy"
	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

func bufferPayload(t *testing.T, pc *protocolCollection, sender common.Address) {
	t.Helper()
	_, _, err := pc.addPayload(&submitSignaturesPayload{typeID: 1, sender: sender})
	require.NoError(t, err)
}

// Before messageAdded, a sender may buffer at most one payload per (round,
// protocol); extras are dropped to bound the ECDSA-recovery burst on drain (DOS-01).
func TestProtocolCollectionBuffersOnePayloadPerSender(t *testing.T) {
	pc := &protocolCollection{signatureCollection: map[common.Hash]*signaturesCollection{}}
	sender := common.HexToAddress("0x1111111111111111111111111111111111111111")

	for range 100 {
		bufferPayload(t, pc, sender)
	}

	require.Len(t, pc.unprocessedPayloads, 1, "only the first payload per sender should be buffered")
}

func TestProtocolCollectionBuffersPerSenderIndependently(t *testing.T) {
	pc := &protocolCollection{signatureCollection: map[common.Hash]*signaturesCollection{}}
	a := common.HexToAddress("0xaaaa000000000000000000000000000000000001")
	b := common.HexToAddress("0xbbbb000000000000000000000000000000000002")

	bufferPayload(t, pc, a)
	bufferPayload(t, pc, a) // dropped
	bufferPayload(t, pc, b)
	bufferPayload(t, pc, b) // dropped

	require.Len(t, pc.unprocessedPayloads, 2, "one buffered payload per distinct sender")
}

// TypeID-0 payloads carry the message inline, so an attacker controls the
// messageHash key per payload. Without the cap, each payload allocates a fresh
// signaturesCollection before any signer check, giving an unbounded memory
// growth vector. The cap limits allocations to one per sender per (round, protocol).
func TestProtocolCollectionCapsTypeZeroAllocations(t *testing.T) {
	sender := common.HexToAddress("0x1111111111111111111111111111111111111111")
	pc := &protocolCollection{
		signatureCollection: map[common.Hash]*signaturesCollection{},
		signingPolicy: &policy.SigningPolicy{
			Voters: voters.NewSet([]common.Address{sender}, []uint16{1}, nil),
		},
	}

	for range 100 {
		// fresh 38-byte message per call → unique hash per payload
		msg := make(shared.Message, 38)
		_, err := rand.Read(msg)
		require.NoError(t, err)
		// Errors from AddSigner (bad signature) are expected and irrelevant here;
		// what we assert is that no further allocations occur past the cap.
		_, _, _ = pc.addPayload(&submitSignaturesPayload{
			typeID:    0,
			sender:    sender,
			message:   msg,
			signature: make([]byte, 65), // satisfy AddSigner length; recovery will fail
		})
	}

	require.LessOrEqual(t, len(pc.signatureCollection), 1,
		"only the first type-0 payload per sender may allocate a signaturesCollection")
}

// bufferRoundPayload stores a buffered (typeID-1) payload so the round exists in the storage.
func bufferRoundPayload(t *testing.T, s *finalizationStorage, round uint32, sp *policy.SigningPolicy) error {
	t.Helper()
	sender := common.HexToAddress("0x1111111111111111111111111111111111111111")
	_, err := s.addPayload(&submitSignaturesPayload{typeID: 1, sender: sender, votingRoundID: round, protocolID: 1}, sp, 1)
	return err
}

// Regression: RemoveRoundsBefore used to set lowestRoundStored to
// votingRoundID+1 while only deleting rounds < votingRoundID, permanently
// leaking one roundCollection per cleanup and rejecting new payloads for a
// round that was never deleted.
func TestRemoveRoundsBeforeKeepsTargetRound(t *testing.T) {
	s := newFinalizationStorage()
	sp := &policy.SigningPolicy{Voters: voters.NewSet([]common.Address{{}}, []uint16{1}, nil)}

	for round := uint32(100); round <= 105; round++ {
		require.NoError(t, bufferRoundPayload(t, s, round, sp))
	}

	s.RemoveRoundsBefore(103)

	require.Equal(t, uint32(103), s.LowestRoundStored())
	for round := uint32(100); round < 103; round++ {
		require.NotContains(t, s.stg, round, "round %d should have been deleted", round)
	}
	for round := uint32(103); round <= 105; round++ {
		require.Contains(t, s.stg, round, "round %d should still be stored", round)
	}

	// the round the cleanup was anchored on must still accept payloads
	require.NoError(t, bufferRoundPayload(t, s, 103, sp))
	// removed rounds must not
	require.Error(t, bufferRoundPayload(t, s, 102, sp))
}

// TestFinalizationStorageConcurrentAccess runs the three real access patterns
// against the storage at once: listener goroutines adding payloads/messages,
// the queue processor reading collections via Get and preparing finalization
// results, and the cleanup path. Run with -race; it pins that Get returns a
// safe snapshot and that RemoveRoundsBefore/LowestRoundStored take the lock.
func TestFinalizationStorageConcurrentAccess(t *testing.T) {
	const round = uint32(50)
	const protocolID = uint8(1)
	const voterCount = 8

	privs := make([]*ecdsa.PrivateKey, voterCount)
	addrs := make([]common.Address, voterCount)
	weights := make([]uint16, voterCount)
	for i := range voterCount {
		privs[i], addrs[i] = newKeyAndAddress(t)
		weights[i] = 1
	}
	sp := &policy.SigningPolicy{Voters: voters.NewSet(addrs, weights, nil)}
	threshold := uint16(voterCount / 2)

	message := make(shared.Message, 38)
	_, err := rand.Read(message)
	require.NoError(t, err)
	msgHash := message.Hash()

	s := newFinalizationStorage()

	var wg sync.WaitGroup
	// payload listeners: one valid signature per voter
	for i := range voterCount {
		wg.Go(func() {
			p := &submitSignaturesPayload{
				typeID:        0,
				sender:        addrs[i],
				votingRoundID: round,
				protocolID:    protocolID,
				message:       message,
				signature:     signVRS(t, msgHash, privs[i]),
			}
			_, err := s.addPayload(p, sp, threshold)
			require.NoError(t, err)
		})
	}
	// message listener
	wg.Go(func() {
		_, _ = s.AddMessage(&shared.ProtocolMessage{ProtocolID: protocolID, VotingRoundID: round, Message: message}, sp, threshold)
	})
	// queue processor: read collections and prepare finalization results
	for range 4 {
		wg.Go(func() {
			for range 200 {
				if sc, exists := s.get(round, protocolID, common.Hash(msgHash)); exists {
					_, _ = PrepareFinalizationResults(sc)
				}
			}
		})
	}
	// cleanup path
	wg.Go(func() {
		for range 50 {
			s.RemoveRoundsBefore(round - 20) // keeps the round under test
			_ = s.LowestRoundStored()
		}
	})
	wg.Wait()

	// all signatures landed and the final state is consistent
	sc, exists := s.get(round, protocolID, common.Hash(msgHash))
	require.True(t, exists)
	require.True(t, sc.thresholdReached)
	require.Equal(t, uint16(voterCount), sc.weight)
}
