package clients

import (
	"strings"
	"time"

	"flare-tlc/logger"
)

const (
	ListenerInterval time.Duration = 2 * time.Second // TODO: change to 10 seconds or read from config
	MaxTxSendRetries int           = 4
	TxRetryInterval  time.Duration = 5 * time.Second
)

type ExecuteStatus[T any] struct {
	Success bool
	Message string
	Value   T
}

func ExecuteWithRetry[T any](f func() (T, error), maxRetries int) <-chan ExecuteStatus[T] {
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
			time.Sleep(TxRetryInterval)
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
