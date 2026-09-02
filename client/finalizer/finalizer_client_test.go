package finalizer

import (
	"context"
	"errors"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/flare-foundation/go-flare-common/pkg/policy"
	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFinalizerClientType0(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clients, err := setupTest(0)
	require.NoError(t, err)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	require.Eventually(
		t, clients.eth.hasAnyCalls, 10*time.Second, 100*time.Millisecond,
	)

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Len(t, clients.eth.sentTxs, 1)

	cupaloy.SnapshotT(t, clients.eth.sentTxs[0])
}

func TestFinalizerClientType1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clients, err := setupTest(1)
	require.NoError(t, err)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	require.Eventually(
		t, clients.eth.hasAnyCalls, 10*time.Second, 100*time.Millisecond,
	)

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Len(t, clients.eth.sentTxs, 1)

	cupaloy.SnapshotT(t, clients.eth.sentTxs[0])
}

func TestFinalizerClientSendTxErr(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	clients, err := setupTest(0)
	require.NoError(t, err)

	clients.eth.sendTxErr = errors.New("sendRawTx error")

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	require.Eventually(
		t, clients.eth.hasAnyCalls, 10*time.Second, 100*time.Millisecond,
	)

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Empty(t, clients.eth.sentTxs)
}

func TestFinalizerClientFetchTxsErr(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	clients, err := setupTest(0)
	require.NoError(t, err)

	clients.db.fetchTxsErr = errors.New("fetchTxs error")

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	require.Eventually(
		t, clients.db.hasAnyFetchTxsCalls, 10*time.Second, 100*time.Millisecond,
	)

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Empty(t, clients.eth.sentTxs)
}

func TestFinalizerClientFetchLogsErr(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clients, err := setupTest(0)
	require.NoError(t, err)

	clients.db.fetchLogsErr = errors.New("fetchLogs error")

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	err = eg.Wait()
	require.True(t, errors.Is(err, clients.db.fetchLogsErr), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Empty(t, clients.eth.sentTxs)
}

// A payload buffered before the message crosses the threshold on the message path, so that
// path must prune old rounds too, or a late provider would starve the cleanup.
func TestThresholdOnTheMessagePathPrunesStorage(t *testing.T) {
	const round = uint32(50)
	priv, addr := newKeyAndAddress(t)
	sp := &policy.SigningPolicy{
		RewardEpochID: 1, StartVotingRoundID: 1, Threshold: 1, Seed: big.NewInt(1),
		Voters: voters.NewSet([]common.Address{addr}, []uint16{2}, nil),
	}
	policies := policy.NewStorage()
	require.NoError(t, policies.Add(sp))

	message, err := encodeMessage(1, round, true, make([]byte, 32))
	require.NoError(t, err)

	storage := newFinalizationStorage(testCutover)
	for old := uint32(20); old < 25; old++ {
		require.NoError(t, bufferRoundPayload(t, storage, old, sp))
	}
	_, err = storage.addPayload(&submitSignaturesPayload{
		sender: addr, votingRoundID: round, protocolID: 1, signature: signVRS(t, testDigest(message), priv),
	}, sp, sp.Threshold)
	require.NoError(t, err)

	fCtx := &finalizerContext{
		votingRoundTiming: &utils.EpochTimingConfig{Start: time.Unix(0, 0), Period: time.Hour},
		rewardEpoch:       &utils.RewardEpochConfig{Start: 0, Period: 100},
	}
	messages := make(chan shared.ProtocolMessage, 1)
	messages <- shared.ProtocolMessage{ProtocolID: 1, VotingRoundID: round, Message: message}
	c := &client{
		signingPolicyStorage: policies,
		finalizationStorage:  storage,
		queueProcessor:       newFinalizerQueueProcessor(&testDB{}, storage, nil, fCtx),
		finalizerContext:     fCtx,
		relayCutover:         testCutover,
		messages:             messages,
	}

	go func() { _ = c.messagesChannelListener(t.Context()) }()

	require.Eventually(t, func() bool { return storage.LowestRoundStored() == round-minRoundsStored },
		5*time.Second, 10*time.Millisecond)
	require.NotNil(t, c.queueProcessor.queue.Pop(), "and the finalization is queued")
	for old := uint32(20); old < 25; old++ {
		require.NotContains(t, storage.stg, old)
	}
}

// A protocol the submitter does not query never gets a local message, so its payloads are not stored.
func TestOnlyConfiguredProtocolsAreStored(t *testing.T) {
	const round = uint32(50)
	sp := &policy.SigningPolicy{
		RewardEpochID: 1, StartVotingRoundID: 1, Threshold: 100,
		Voters: voters.NewSet([]common.Address{{}}, []uint16{1}, nil),
	}
	policies := policy.NewStorage()
	require.NoError(t, policies.Add(sp))

	storage := newFinalizationStorage(testCutover)
	c := &client{
		signingPolicyStorage: policies,
		finalizationStorage:  storage,
		finalizerContext: &finalizerContext{
			votingRoundTiming: &utils.EpochTimingConfig{Start: time.Unix(0, 0), Period: time.Hour},
			rewardEpoch:       &utils.RewardEpochConfig{Start: 0, Period: 100},
			protocolIDs:       map[uint8]struct{}{1: {}},
		},
	}

	sender := common.HexToAddress("0x1111111111111111111111111111111111111111")
	require.NoError(t, c.ProcessSubmissionData([]*submitSignaturesPayload{
		{sender: sender, votingRoundID: round, protocolID: 1},
		{sender: sender, votingRoundID: round, protocolID: 200},
	}))

	rc, exists := storage.stg[round]
	require.True(t, exists)
	require.Contains(t, rc.protocolCollections, uint8(1))
	require.NotContains(t, rc.protocolCollections, uint8(200), "an unconfigured protocol must not be buffered")
}
