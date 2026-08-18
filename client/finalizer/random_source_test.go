package finalizer

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/policy"
	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

const randomProtocolID = uint8(100)

// The data arrives through the random protocol's own message, so that protocol must be
// one the submitter queries; anything else must fail startup.
func TestRandomProtocolConfigured(t *testing.T) {
	protocols := map[string]config.ProtocolConfig{
		"ftso": {ID: randomProtocolID, APIURL: "https://ftso.example/api"},
		"fdc":  {ID: 200, APIURL: "https://fdc.example"},
	}
	require.True(t, randomProtocolConfigured(protocols, randomProtocolID))
	require.False(t, randomProtocolConfigured(protocols, 42))
	require.False(t, randomProtocolConfigured(nil, randomProtocolID))
}

// stubRandomSource answers Get with fixed data so the gate and the trailer can be observed.
type stubRandomSource struct {
	data      randomData
	err       error
	getCalled atomic.Int32
}

func (s *stubRandomSource) Put(uint32, randomData) {}

func (s *stubRandomSource) Get(context.Context, uint32) (randomData, error) {
	s.getCalled.Add(1)
	return s.data, s.err
}

func (s *stubRandomSource) RemoveBefore(uint32) {}

func queueProcessorForTest(t *testing.T, src randomSource, cutover *shared.RelayCutover) *finalizerQueueProcessor {
	t.Helper()
	return &finalizerQueueProcessor{
		queue:               newFinalizerQueue(),
		finalizationStorage: newFinalizationStorage(cutover),
		relayClient:         &relayContractClient{address: oldRelayAddress, relayCutover: cutover},
		finalizerContext:    &finalizerContext{randomNumberProtocolID: randomProtocolID},
		randomSource:        src,
	}
}

// storeCollection registers a signature collection so needsRandomTrailer can read
// the round's signing policy, and returns the item that addresses it.
func storeCollection(t *testing.T, p *finalizerQueueProcessor, protocolID uint8, votingRoundID uint32, rewardEpoch int64, message shared.Message) *queueItem {
	t.Helper()

	sp := &policy.SigningPolicy{
		RewardEpochID: rewardEpoch,
		Voters:        voters.NewSet([]common.Address{{}}, []uint16{1}, nil),
	}
	_, err := p.finalizationStorage.AddMessage(
		&shared.ProtocolMessage{ProtocolID: protocolID, VotingRoundID: votingRoundID, Message: message}, sp, 1)
	require.NoError(t, err)

	// the collection is keyed by the digest, which depends on the policy epoch
	msgHash := common.Hash(p.finalizationStorage.relayCutover.DigestForRewardEpoch(message, rewardEpoch))

	return &queueItem{protocolID: protocolID, votingRoundID: votingRoundID, msgHash: msgHash}
}

// The trailer exists only for the new Relay: it is what verifies and stores it. A
// pre-cutover round still goes to the old Relay, which knows nothing about it.
func TestNeedsRandomTrailerOnlyFromTheBreakingEpoch(t *testing.T) {
	cutover := cutoverForTest()
	message := buildMessage(randomProtocolID, 7, 1, common.Hash{})

	cases := []struct {
		name        string
		protocolID  uint8
		rewardEpoch int64
		want        bool
	}{
		{"random protocol before the breaking epoch", randomProtocolID, testBreakingEpoch - 1, false},
		{"random protocol at the breaking epoch", randomProtocolID, testBreakingEpoch, true},
		{"random protocol after the breaking epoch", randomProtocolID, testBreakingEpoch + 1, true},
		{"another protocol at the breaking epoch", 200, testBreakingEpoch, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p := queueProcessorForTest(t, &stubRandomSource{}, cutover)
			item := storeCollection(t, p, c.protocolID, 7, c.rewardEpoch, message)
			require.Equal(t, c.want, p.needsRandomTrailer(item))
		})
	}

	// with no cutover scheduled at all, no round ever needs one
	plain := queueProcessorForTest(t, &stubRandomSource{}, shared.NewRelayCutover(testChainID))
	item := storeCollection(t, plain, randomProtocolID, 7, testBreakingEpoch+1, message)
	require.False(t, plain.needsRandomTrailer(item))
}

func TestRandomTrailerVerifiesAgainstTheSignedMessage(t *testing.T) {
	leaf := randomLeaf(7, randomValue, true)
	root := sortedPair(leaf, proofNodeA)
	message := buildMessage(randomProtocolID, 7, 1, root)

	good := randomData{Value: randomValue, Proof: []common.Hash{proofNodeA}}
	p := queueProcessorForTest(t, &stubRandomSource{data: good}, cutoverForTest())

	trailer, err := p.randomTrailer(context.Background(), message, 7)
	require.NoError(t, err)
	require.Equal(t, good.trailer(), trailer)

	// a proof that cannot finalize must be caught here, not on chain
	bad := randomData{Value: common.HexToHash("0x02"), Proof: []common.Hash{proofNodeA}}
	p = queueProcessorForTest(t, &stubRandomSource{data: bad}, cutoverForTest())
	_, err = p.randomTrailer(context.Background(), message, 7)
	require.ErrorContains(t, err, "folds to")

	// and a message for a different round than the item is a bug, not a proof failure
	p = queueProcessorForTest(t, &stubRandomSource{data: good}, cutoverForTest())
	_, err = p.randomTrailer(context.Background(), message, 8)
	require.ErrorContains(t, err, "message is for round 7")
}

// The trailer only reaches the tx input when the finalization sets one.
func TestPrepareFinalizationTxInputAppendsTheTrailer(t *testing.T) {
	sp := &policy.SigningPolicy{Voters: voters.NewSet([]common.Address{{}}, []uint16{1}, nil)}
	message := buildMessage(randomProtocolID, 7, 1, common.Hash{})
	trailer := randomData{Value: randomValue, Proof: []common.Hash{proofNodeA}}.trailer()

	withTrailer := FinalizationResult{
		message:       message,
		signingPolicy: sp,
		signatures:    []IndexedSignature{{index: 0, signature: make([]byte, 65)}},
		randomTrailer: trailer,
	}
	without := FinalizationResult{
		message:       message,
		signingPolicy: sp,
		signatures:    []IndexedSignature{{index: 0, signature: make([]byte, 65)}},
	}

	withInput, err := withTrailer.PrepareFinalizationTxInput()
	require.NoError(t, err)
	withoutInput, err := without.PrepareFinalizationTxInput()
	require.NoError(t, err)

	require.Equal(t, append(append([]byte{}, withoutInput...), trailer...), withInput)

	// a trailer that is not a whole number of words would revert on chain
	broken := withTrailer
	broken.randomTrailer = trailer[:len(trailer)-1]
	_, err = broken.PrepareFinalizationTxInput()
	require.ErrorContains(t, err, "not a multiple of")
}

// The provider serves the trailer bytes themselves; only their shape is checked here,
// the fold checks their meaning.
func TestParseRandomData(t *testing.T) {
	value := randomValue.Bytes()
	twoWords := append(append([]byte{}, value...), proofNodeA.Bytes()...)

	data, err := parseRandomData(value)
	require.NoError(t, err)
	require.Equal(t, randomData{Value: randomValue}, data, "a single word is a single-leaf tree")

	data, err = parseRandomData(twoWords)
	require.NoError(t, err)
	require.Equal(t, randomData{Value: randomValue, Proof: []common.Hash{proofNodeA}}, data)

	_, err = parseRandomData(nil)
	require.ErrorContains(t, err, "0 bytes")
	_, err = parseRandomData(twoWords[:48])
	require.ErrorContains(t, err, "48 bytes")
}

// The store is a rendezvous between the message listener and the send: whichever
// side comes first, the send gets the data or a bounded wait.
func TestRandomStoreHandsOverInEitherOrder(t *testing.T) {
	good := randomData{Value: randomValue, Proof: []common.Hash{proofNodeA}}

	store := newRandomStore()
	store.Put(7, good)
	data, err := store.Get(context.Background(), 7)
	require.NoError(t, err)
	require.Equal(t, good, data)

	// a Get already waiting is woken by the Put
	store = newRandomStore()
	got := make(chan randomData, 1)
	go func() {
		data, err := store.Get(context.Background(), 7)
		require.NoError(t, err)
		got <- data
	}()
	time.Sleep(10 * time.Millisecond)
	store.Put(7, good)
	select {
	case data := <-got:
		require.Equal(t, good, data)
	case <-time.After(time.Second):
		t.Fatal("Get was not woken by Put")
	}

	// the first Put wins; a later one for the same round is ignored
	store.Put(7, randomData{Value: common.HexToHash("0x02")})
	data, err = store.Get(context.Background(), 7)
	require.NoError(t, err)
	require.Equal(t, good, data)
}

// A message that never comes must not hold the send: the wait is bounded, and so is
// the caller's own deadline.
func TestRandomStoreGetIsBounded(t *testing.T) {
	store := newRandomStore()
	store.wait = 20 * time.Millisecond

	_, err := store.Get(context.Background(), 9)
	require.ErrorContains(t, err, "no finalization data for round 9")

	store.wait = time.Minute
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = store.Get(ctx, 9)
	require.ErrorIs(t, err, context.Canceled)
}

func TestRandomStoreRemoveBefore(t *testing.T) {
	store := newRandomStore()
	store.Put(5, randomData{Value: randomValue})
	store.Put(7, randomData{Value: randomValue})
	store.wait = time.Millisecond
	_, _ = store.Get(context.Background(), 6) // leaves a pending entry

	store.RemoveBefore(7)
	require.Len(t, store.entries, 1)
	_, ok := store.entries[7]
	require.True(t, ok)
}

func clientForFinalizationData(cutover *shared.RelayCutover) (*client, *randomStore) {
	store := newRandomStore()
	store.wait = time.Millisecond
	c := &client{
		finalizerContext: &finalizerContext{randomNumberProtocolID: randomProtocolID},
		relayCutover:     cutover,
		queueProcessor:   &finalizerQueueProcessor{randomSource: store},
	}
	return c, store
}

// What arrives with the message is checked there and then: only data that folds to the
// message's own root is kept, and only for the protocol and rounds that need it.
func TestStoreFinalizationData(t *testing.T) {
	leaf := randomLeaf(7, randomValue, true)
	root := sortedPair(leaf, proofNodeA)
	message := buildMessage(randomProtocolID, 7, 1, root)
	good := randomData{Value: randomValue, Proof: []common.Hash{proofNodeA}}
	policyAt := func(epoch int64) *policy.SigningPolicy { return &policy.SigningPolicy{RewardEpochID: epoch} }

	cases := []struct {
		name       string
		protocolID uint8
		epoch      int64
		data       []byte
		stored     bool
	}{
		{"kept for the random protocol from the breaking epoch", randomProtocolID, testBreakingEpoch, good.trailer(), true},
		{"a single-leaf tree is kept", randomProtocolID, testBreakingEpoch, nil, false}, // placeholder, replaced below
		{"missing data is reported, not stored", randomProtocolID, testBreakingEpoch, nil, false},
		{"a truncated word is not stored", randomProtocolID, testBreakingEpoch, good.trailer()[:40], false},
		{"data that does not fold is not stored", randomProtocolID, testBreakingEpoch,
			randomData{Value: common.HexToHash("0x02"), Proof: []common.Hash{proofNodeA}}.trailer(), false},
		{"ignored before the breaking epoch", randomProtocolID, testBreakingEpoch - 1, good.trailer(), false},
		{"ignored for another protocol", 200, testBreakingEpoch, good.trailer(), false},
	}
	// the single-leaf case needs its own message, whose root is the leaf itself
	singleLeaf := buildMessage(randomProtocolID, 7, 1, leaf)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cl, store := clientForFinalizationData(cutoverForTest())
			msg := &shared.ProtocolMessage{ProtocolID: c.protocolID, VotingRoundID: 7, Message: message, FinalizationData: c.data}
			want := good
			if c.name == "a single-leaf tree is kept" {
				msg.Message, msg.FinalizationData = singleLeaf, randomValue.Bytes()
				want, c.stored = randomData{Value: randomValue}, true
			}

			cl.storeFinalizationData(msg, policyAt(c.epoch))

			data, err := store.Get(context.Background(), 7)
			if !c.stored {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, want, data)
		})
	}

	// with no cutover scheduled nothing is expected, and the stand-in source drops it
	cl, _ := clientForFinalizationData(shared.NewRelayCutover(testChainID))
	cl.queueProcessor.randomSource = unconfiguredRandomSource{}
	cl.storeFinalizationData(&shared.ProtocolMessage{ProtocolID: randomProtocolID, VotingRoundID: 7, Message: message, FinalizationData: good.trailer()}, policyAt(testBreakingEpoch))
}
