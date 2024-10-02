package shared

import (
	"flare-tlc/logger"
	"strings"
	"time"
)

type ExecuteStatus[T any] struct {
	Success bool
	Message string
	Value   T
}

func ExecuteWithRetry[T any](f func() (T, error), maxRetries int, delay time.Duration) <-chan ExecuteStatus[T] {
	out := make(chan ExecuteStatus[T])
	go func() {
		for ri := 0; ri < maxRetries; ri++ {
			result, err := f()
			if err == nil {
				out <- ExecuteStatus[T]{Success: true, Value: result}
				return
			} else {
				logger.Error("error executing in retry no. %d: %v", ri, err)
			}
			time.Sleep(delay)
		}
		out <- ExecuteStatus[T]{Success: false, Message: "max retries reached"}
	}()
	return out
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
				logger.Error("error executing in retry no. %d: %v", ri, err)
			}
			time.Sleep(delay)
		}
		out <- ExecuteStatus[T]{Success: false, Message: "max retries reached"}
	}()
	return out
}

// ExistsAsSubstring returns true if any of the strings in the slice is a substring of s
func ExistsAsSubstring(slice []string, s string) bool {
	for _, item := range slice {
		if strings.Contains(s, item) {
			return true
		}
	}
	return false
}
