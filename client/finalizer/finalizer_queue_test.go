package finalizer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare-system-client/utils"
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
