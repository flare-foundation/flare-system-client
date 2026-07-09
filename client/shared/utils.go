package shared

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

type ExecuteStatus[T any] struct {
	Success bool
	Message string
	Value   T
}

// ExecuteWithRetryChan retries f until success, maxRetries, or ctx cancellation
// (the delay between attempts is interruptible, so shutdown is not blocked).
func ExecuteWithRetryChan[T any](ctx context.Context, f func() (T, error), maxRetries int, delay time.Duration) <-chan ExecuteStatus[T] {
	out := make(chan ExecuteStatus[T])
	go func() {
		var finalError error
		for ri := range maxRetries {
			if err := ctx.Err(); err != nil {
				out <- ExecuteStatus[T]{Success: false, Message: fmt.Sprintf("context done: %v (last error: %v)", err, finalError)}
				return
			}
			result, err := f()
			if err == nil {
				out <- ExecuteStatus[T]{Success: true, Value: result}
				return
			}
			logger.Debugf("executing in retry no. %d: %v", ri, err)
			finalError = err
			select {
			case <-time.After(delay):
			case <-ctx.Done():
				out <- ExecuteStatus[T]{Success: false, Message: fmt.Sprintf("context done: %v (last error: %v)", ctx.Err(), finalError)}
				return
			}
		}
		out <- ExecuteStatus[T]{Success: false, Message: fmt.Sprintf("max retries reached: %v", finalError)}
	}()
	return out
}

// ExecuteWithRetryWithContext retries function f until success or ctx is canceled.
// Between starts of each retries there is at least minimalDuration time.
func ExecuteWithRetryWithContext[T any](ctx context.Context, f func() (T, error), minimalDuration time.Duration) ExecuteStatus[T] {
	var err error
	var result T

	for {
		timer := time.NewTimer(minimalDuration)

		select {
		case <-ctx.Done():
			return ExecuteStatus[T]{Success: false, Message: fmt.Sprintf("context closed, final error: %v", err)}
		default:
		}

		result, err = f()
		if err == nil {
			return ExecuteStatus[T]{Success: true, Value: result}
		} else {
			<-timer.C
		}
	}
}

// ExecuteWithRetryAttempts is ExecuteWithRetryChan where f receives the attempt
// number; same cancellation semantics.
func ExecuteWithRetryAttempts[T any](ctx context.Context, f func(int) (T, error), maxRetries int, delay time.Duration) <-chan ExecuteStatus[T] {
	out := make(chan ExecuteStatus[T])
	go func() {
		var finalError error
		for ri := range maxRetries {
			if err := ctx.Err(); err != nil {
				out <- ExecuteStatus[T]{Success: false, Message: fmt.Sprintf("context done: %v (last error: %v)", err, finalError)}
				return
			}
			result, err := f(ri)
			if err == nil {
				out <- ExecuteStatus[T]{Success: true, Value: result}
				return
			}
			logger.Debugf("executing in retry no. %d: %v", ri, err)
			finalError = err
			select {
			case <-time.After(delay):
			case <-ctx.Done():
				out <- ExecuteStatus[T]{Success: false, Message: fmt.Sprintf("context done: %v (last error: %v)", ctx.Err(), finalError)}
				return
			}
		}
		out <- ExecuteStatus[T]{Success: false, Message: fmt.Sprintf("max retries reached: %v", finalError)}
	}()
	return out
}

// ExistsAsSubstring returns true if any of the strings in the slice is a substring of s.
func ExistsAsSubstring(slice []string, s string) bool {
	for _, item := range slice {
		if strings.Contains(s, item) {
			return true
		}
	}
	return false
}
