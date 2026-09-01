package finalizer

import (
	"bytes"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/flare-foundation/go-flare-common/pkg/policy"
	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

func TestDelayedRetryTime(t *testing.T) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	p := &finalizerQueueProcessor{
		finalizerContext: &finalizerContext{
			votingRoundTiming: &utils.EpochTimingConfig{
				Start:  start,
				Period: 90 * time.Second,
			},
			gracePeriodEndOffset: 65 * time.Second,
		},
	}

	// Grace period end of round 2 (round 2 starts at start + 2*90s).
	graceEnd := start.Add(2*90*time.Second + 65*time.Second)

	t.Run("future target is kept", func(t *testing.T) {
		now := graceEnd.Add(-10 * time.Second)
		require.Equal(t, graceEnd, p.delayedRetryTime(1, now))
	})

	t.Run("past target is rescheduled ahead of now", func(t *testing.T) {
		// A send cut off by queueSendTimeout finishes past the grace-period
		// end; the retry must not be scheduled in the past (the delayed queue
		// drops past-time targets).
		now := graceEnd.Add(30 * time.Second)
		require.Equal(t, now.Add(delayedRetryDelay), p.delayedRetryTime(1, now))
	})

	t.Run("target equal to now is rescheduled", func(t *testing.T) {
		require.Equal(t, graceEnd.Add(delayedRetryDelay), p.delayedRetryTime(1, graceEnd))
	})

	t.Run("clamped target is truncated to the second so retries batch", func(t *testing.T) {
		// timeMap keys compare time.Time values exactly: sub-second and monotonic
		// components would give every clamped retry its own timer and DB query.
		now := graceEnd.Add(30*time.Second + 123*time.Millisecond)
		got := p.delayedRetryTime(1, now)
		require.Equal(t, now.Add(delayedRetryDelay).Truncate(time.Second), got)
		require.Equal(t, got, p.delayedRetryTime(1, now.Add(500*time.Millisecond)))

		// a monotonic reading must be stripped too; require.Equal compares the
		// struct exactly, unlike time.Time.Equal which ignores monotonic parts
		live := p.delayedRetryTime(1, time.Now())
		require.Equal(t, live.Round(0), live)
	})
}

// a failed already-relayed lookup must not drop the batch — items are already off the queue
func TestProcessDelayedQueueSurvivesDBError(t *testing.T) {
	privateKey, err := crypto.HexToECDSA(testPrivateKeyHex)
	require.NoError(t, err)
	sender := crypto.PubkeyToAddress(privateKey.PublicKey)

	msg, err := encodeMessage(1, 1, true, bytes.Repeat([]byte{0xff}, 32))
	require.NoError(t, err)
	sig, err := signMessage(msg, privateKey)
	require.NoError(t, err)

	sp := &policy.SigningPolicy{Voters: voters.NewSet([]common.Address{sender}, []uint16{2}, nil)}
	storage := newFinalizationStorage(testCutover)
	ready, err := storage.addPayload(&submitSignaturesPayload{
		protocolID: 1, votingRoundID: 1, typeID: 0, message: msg, signature: sig, sender: sender,
	}, sp, 1)
	require.NoError(t, err)
	require.True(t, ready.thresholdReached)

	eth := new(testEthClient)
	relayClient, err := NewRelayContractClient(nil, relayContractAddress, privateKey, sender, &config.Gas{}, testChainID, testCutover)
	require.NoError(t, err)
	relayClient.chainClient = eth
	relayClient.retryDelay = time.Millisecond

	qp := newFinalizerQueueProcessor(
		&testDB{fetchLogsErr: errors.New("db down")},
		storage,
		relayClient,
		&finalizerContext{votingRoundTiming: &utils.EpochTimingConfig{Start: time.Unix(0, 0), Period: time.Hour}},
	)

	err = qp.processDelayedQueue(context.Background(), []*queueItem{
		{votingRoundID: ready.votingRoundID, protocolID: ready.protocolID, digest: ready.digest},
	})
	require.NoError(t, err)
	require.Len(t, eth.sentTxs, 1, "item must be sent despite the failed dedup query")
}
