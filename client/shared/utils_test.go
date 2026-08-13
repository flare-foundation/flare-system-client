package shared

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// A canceled ctx must stop the retry loop promptly instead of grinding through
// the full budget of sleeps.
func TestExecuteWithRetryAttemptsStopsOnCanceledCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // already canceled before the first attempt

	start := time.Now()
	res := <-ExecuteWithRetryAttempts(ctx, func(int) (int, error) {
		return 0, errors.New("boom")
	}, 10, 5*time.Second)

	require.False(t, res.Success)
	require.Less(t, time.Since(start), time.Second, "must not grind ~50s after cancellation")
}

// Cancellation during the inter-attempt delay must interrupt the sleep.
func TestExecuteWithRetryChanStopsWhenCanceledDuringDelay(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	start := time.Now()
	res := <-ExecuteWithRetryChan(ctx, func() (int, error) {
		return 0, errors.New("boom")
	}, 10, 5*time.Second)

	require.False(t, res.Success)
	require.Less(t, time.Since(start), time.Second, "delay must be interruptible by ctx")
}

// A live ctx must not change the normal success / exhaustion behavior.
func TestExecuteWithRetryLiveCtxUnaffected(t *testing.T) {
	res := <-ExecuteWithRetryChan(context.Background(), func() (int, error) {
		return 42, nil
	}, 3, time.Millisecond)
	require.True(t, res.Success)
	require.Equal(t, 42, res.Value)

	calls := 0
	res = <-ExecuteWithRetryAttempts(context.Background(), func(int) (int, error) {
		calls++
		return 0, errors.New("boom")
	}, 3, time.Millisecond)
	require.False(t, res.Success)
	require.Equal(t, 3, calls) // full budget used when ctx stays live
}
