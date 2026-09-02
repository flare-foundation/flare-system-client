package finalizer

import (
	"crypto/ecdsa"
	"crypto/rand"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/flare-foundation/go-flare-common/pkg/policy"
	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

func bufferPayload(t *testing.T, pc *protocolCollection, sender common.Address) {
	t.Helper()
	_, err := pc.addPayload(&submitSignaturesPayload{sender: sender})
	require.NoError(t, err)
}

// Before the message, a sender may buffer at most one payload per (round,
// protocol); extras are dropped to bound the ECDSA-recovery burst on drain (DOS-01).
func TestProtocolCollectionBuffersOnePayloadPerSender(t *testing.T) {
	pc := &protocolCollection{relayCutover: testCutover}
	sender := common.HexToAddress("0x1111111111111111111111111111111111111111")

	for range 100 {
		bufferPayload(t, pc, sender)
	}

	require.Len(t, pc.unprocessedPayloads, 1, "only the first payload per sender should be buffered")
}

func TestProtocolCollectionBuffersPerSenderIndependently(t *testing.T) {
	pc := &protocolCollection{relayCutover: testCutover}
	a := common.HexToAddress("0xaaaa000000000000000000000000000000000001")
	b := common.HexToAddress("0xbbbb000000000000000000000000000000000002")

	bufferPayload(t, pc, a)
	bufferPayload(t, pc, a) // dropped
	bufferPayload(t, pc, b)
	bufferPayload(t, pc, b) // dropped

	require.Len(t, pc.unprocessedPayloads, 2, "one buffered payload per distinct sender")
}

// Only the local message creates a collection. Payloads that arrive first, whatever their type,
// wait in the buffer and are verified against that message when it lands.
func TestPayloadsWaitForTheMessage(t *testing.T) {
	const round = uint32(50)
	priv, addr := newKeyAndAddress(t)
	sp := &policy.SigningPolicy{Voters: voters.NewSet([]common.Address{addr}, []uint16{2}, nil)}

	message := make(shared.Message, 38)
	_, err := rand.Read(message)
	require.NoError(t, err)

	s := newFinalizationStorage(testCutover)
	ready, err := s.addPayload(&submitSignaturesPayload{
		sender: addr, votingRoundID: round, protocolID: 1,
		signature: signVRS(t, testDigest(message), priv),
	}, sp, 1)
	require.NoError(t, err)
	require.False(t, ready.thresholdReached)
	_, exists := s.get(round, 1)
	require.False(t, exists, "no collection before the message")

	ready, err = s.AddMessage(&shared.ProtocolMessage{ProtocolID: 1, VotingRoundID: round, Message: message}, sp, 1)
	require.NoError(t, err)
	require.True(t, ready.thresholdReached, "the buffered signature counts once the message is known")
	sc, exists := s.get(round, 1)
	require.True(t, exists)
	require.Equal(t, message, sc.message)
}

// The first local message wins: a re-served one must not displace the root, drop the counted
// weight or refresh finalizationData.
func TestASecondMessageCannotDisplaceTheFirst(t *testing.T) {
	const round = uint32(60)
	priv, addr := newKeyAndAddress(t)
	sp := &policy.SigningPolicy{Voters: voters.NewSet([]common.Address{addr}, []uint16{2}, nil)}

	first := buildMessage(1, round, randomValue)
	firstData := words(randomValue)

	s := newFinalizationStorage(testCutover)
	_, err := s.AddMessage(&shared.ProtocolMessage{
		ProtocolID: 1, VotingRoundID: round, Message: first, FinalizationData: firstData,
	}, sp, 1)
	require.NoError(t, err)

	ready, err := s.addPayload(&submitSignaturesPayload{
		sender: addr, votingRoundID: round, protocolID: 1, signature: signVRS(t, testDigest(first), priv),
	}, sp, 1)
	require.NoError(t, err)
	require.True(t, ready.thresholdReached)

	ready, secondErr := s.AddMessage(&shared.ProtocolMessage{
		ProtocolID: 1, VotingRoundID: round, Message: buildMessage(1, round, proofNodeA),
		FinalizationData: words(proofNodeA),
	}, sp, 1)

	sc, exists := s.get(round, 1)
	require.True(t, exists)
	require.Equal(t, first, sc.message, "the root the counted signature covers")
	require.Equal(t, firstData, sc.finalizationData)
	require.Equal(t, uint16(2), sc.weight, "the counted signature survives")
	require.True(t, sc.thresholdReached)
	require.ErrorContains(t, secondErr, "message added twice")
	require.False(t, ready.thresholdReached, "one round and protocol queue at most once")
}

// bufferRoundPayload stores a buffered payload so the round exists in the storage.
func bufferRoundPayload(t *testing.T, s *finalizationStorage, round uint32, sp *policy.SigningPolicy) error {
	t.Helper()
	sender := common.HexToAddress("0x1111111111111111111111111111111111111111")
	_, err := s.addPayload(&submitSignaturesPayload{sender: sender, votingRoundID: round, protocolID: 1}, sp, 1)
	return err
}

// Regression: RemoveRoundsBefore used to set lowestRoundStored to
// votingRoundID+1 while only deleting rounds < votingRoundID, permanently
// leaking one roundCollection per cleanup and rejecting new payloads for a
// round that was never deleted.
func TestRemoveRoundsBeforeKeepsTargetRound(t *testing.T) {
	s := newFinalizationStorage(testCutover)
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
	digest := testDigest(message)

	s := newFinalizationStorage(testCutover)

	var wg sync.WaitGroup
	// payload listeners: one valid signature per voter
	for i := range voterCount {
		wg.Go(func() {
			p := &submitSignaturesPayload{
				sender:        addrs[i],
				votingRoundID: round,
				protocolID:    protocolID,
				signature:     signVRS(t, digest, privs[i]),
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
				if sc, exists := s.get(round, protocolID); exists {
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
	sc, exists := s.get(round, protocolID)
	require.True(t, exists)
	require.True(t, sc.thresholdReached)
	require.Equal(t, uint16(voterCount), sc.weight)
}

// TestBadPayloadClassification pins that the four payload-caused rejections wrap
// errBadPayload: the Debug-vs-Error log split in ProcessSubmissionData and
// addMessage's drain relies on errors.Is.
func TestBadPayloadClassification(t *testing.T) {
	hash := common.HexToHash("0x0102030405060708091011121314151617181920212223242526272829303132").Bytes()

	t.Run("signature recovery failure", func(t *testing.T) {
		priv, _ := newKeyAndAddress(t)
		sig := signVRS(t, hash, priv)
		sig[0] = 99 // invalid V
		pld := &submitSignaturesPayload{signature: sig, voterIndex: -1}
		set := voters.NewSet([]common.Address{{}}, []uint16{1}, nil)
		require.ErrorIs(t, pld.AddSigner(hash, set), errBadPayload)
	})

	t.Run("unregistered signer", func(t *testing.T) {
		priv, _ := newKeyAndAddress(t)
		pld := &submitSignaturesPayload{signature: signVRS(t, hash, priv), voterIndex: -1}
		stranger := common.HexToAddress("0x2222222222222222222222222222222222222222")
		set := voters.NewSet([]common.Address{stranger}, []uint16{1}, nil)
		require.ErrorIs(t, pld.AddSigner(hash, set), errBadPayload)
	})

	t.Run("duplicate signature", func(t *testing.T) {
		_, addr := newKeyAndAddress(t)
		sp := &policy.SigningPolicy{Voters: voters.NewSet([]common.Address{addr}, []uint16{1}, nil)}
		sc := NewSignatureCollection(shared.Message{}, sp, 100)
		pld := &submitSignaturesPayload{signature: []byte{1}, voterIndex: 0, signer: addr}
		_, err := sc.addSignature(pld)
		require.NoError(t, err)
		_, err = sc.addSignature(pld)
		require.ErrorIs(t, err, errBadPayload)
	})

	t.Run("round below lowest stored", func(t *testing.T) {
		s := newFinalizationStorage(testCutover)
		s.lowestRoundStored = 5
		sp := &policy.SigningPolicy{Voters: voters.NewSet([]common.Address{{}}, []uint16{1}, nil)}
		_, err := s.addPayload(&submitSignaturesPayload{votingRoundID: 1}, sp, 100)
		require.ErrorIs(t, err, errBadPayload)
	})
}

// captureWarnings redirects the process logger to a file for the duration of the test.
// Tests using it must not be parallel: the logger is global.
func captureWarnings(t *testing.T) func() string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "warn.log")
	logger.Set(logger.Config{Level: "WARN", File: path, MaxFileSize: 1})
	t.Cleanup(func() { logger.Set(logger.DefaultConfig()) })

	return func() string {
		logger.SyncFileLogger()
		b, err := os.ReadFile(path)
		if err != nil {
			return ""
		}
		return string(b)
	}
}

// Giving up on a round is the only Warn-level signal; protocols we never served a message for are skipped.
func TestPruneWarnsAboutAnUnfinalizedRound(t *testing.T) {
	warnings := captureWarnings(t)
	sp := &policy.SigningPolicy{Voters: voters.NewSet([]common.Address{{}}, []uint16{1}, nil)}
	message := make(shared.Message, shared.RelayMessageLength)

	s := newFinalizationStorage(testCutover)
	// protocol 1: we served the message, nothing valid signed it
	_, err := s.AddMessage(&shared.ProtocolMessage{ProtocolID: 1, VotingRoundID: 40, Message: message}, sp, 100)
	require.NoError(t, err)
	// protocol 200: peers submit, we never served a message
	_, err = s.addPayload(&submitSignaturesPayload{
		sender: common.HexToAddress("0x1111111111111111111111111111111111111111"), votingRoundID: 40, protocolID: 200,
	}, sp, 100)
	require.NoError(t, err)

	s.RemoveRoundsBefore(41)

	out := warnings()
	require.Contains(t, out, "Discarding round 40 for protocol 1 unfinalized")
	require.NotContains(t, out, "protocol 200", "a protocol this node does not answer for is not its business")
}

// A round that did reach the threshold is discarded quietly.
func TestPruneIsQuietForAFinalizedRound(t *testing.T) {
	warnings := captureWarnings(t)
	priv, addr := newKeyAndAddress(t)
	sp := &policy.SigningPolicy{Voters: voters.NewSet([]common.Address{addr}, []uint16{2}, nil)}
	message := make(shared.Message, shared.RelayMessageLength)

	s := newFinalizationStorage(testCutover)
	_, err := s.AddMessage(&shared.ProtocolMessage{ProtocolID: 1, VotingRoundID: 40, Message: message}, sp, 1)
	require.NoError(t, err)
	ready, err := s.addPayload(&submitSignaturesPayload{
		sender: addr, votingRoundID: 40, protocolID: 1, signature: signVRS(t, testDigest(message), priv),
	}, sp, 1)
	require.NoError(t, err)
	require.True(t, ready.thresholdReached)

	s.RemoveRoundsBefore(41)

	require.NotContains(t, warnings(), "round 40")
}
