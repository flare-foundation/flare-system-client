package shared

import (
	"context"
	"strings"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

type ExecuteStatus[T any] struct {
	Success bool
	Message string
	Value   T
}

// ExecuteWithRetryChan executes function with retry until success or maxRetries. Between each retries there is a delay.
func ExecuteWithRetryChan[T any](f func() (T, error), maxRetries int, delay time.Duration) <-chan ExecuteStatus[T] {
	out := make(chan ExecuteStatus[T])
	go func() {
		for ri := 0; ri <= maxRetries; ri++ {
			result, err := f()
			if err == nil {
				out <- ExecuteStatus[T]{Success: true, Value: result}
				return
			} else {
				logger.Errorf("error executing in retry no. %d: %v", ri, err)
			}
			time.Sleep(delay)
		}
		out <- ExecuteStatus[T]{Success: false, Message: "max retries reached"}
	}()
	return out
}

// ExecuteWithRetryWithContext retries function f until success or ctx is canceled. Between each retries there is at least minimalDuration time.
func ExecuteWithRetryWithContext[T any](ctx context.Context, f func() (T, error), minimalDuration time.Duration) ExecuteStatus[T] {
	for {
		timer := time.NewTimer(minimalDuration)

		select {
		case <-ctx.Done():
			return ExecuteStatus[T]{Success: false, Message: "context canceled"}
		default:
		}

		result, err := f()
		if err == nil {
			return ExecuteStatus[T]{Success: true, Value: result}
		} else {
			<-timer.C
		}
	}
}

func ExecuteWithRetryAttempts[T any](f func(int) (T, error), maxRetries int, delay time.Duration) <-chan ExecuteStatus[T] {
	out := make(chan ExecuteStatus[T])
	go func() {
		for ri := 0; ri < maxRetries; ri++ {
			result, err := f(ri)
			if err == nil {
				out <- ExecuteStatus[T]{Success: true, Value: result}
				return
			} else {
				logger.Errorf("error executing in retry no. %d: %v", ri, err)
			}
			time.Sleep(delay)
		}
		out <- ExecuteStatus[T]{Success: false, Message: "max retries reached"}
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
