package clients

import (
	"time"

	"flare-tlc/logger"
)

const (
	ListenerInterval time.Duration = 10 * time.Second
	MaxTxSendRetries int           = 5
	TxRetryInterval  time.Duration = 5 * time.Second
)

type ExecuteStatus struct {
	Success bool
	Message string
}

func ExecuteWithRetry(f func() error, maxRetries int) <-chan ExecuteStatus {
	out := make(chan ExecuteStatus)
	go func() {
		for ri := 0; ri < maxRetries; ri++ {
			err := f()
			if err == nil {
				out <- ExecuteStatus{Success: true}
				return
			} else {
				logger.Error("error executing in retry no. %d: %w", ri, err)
			}
			time.Sleep(TxRetryInterval)
		}
		out <- ExecuteStatus{Success: false, Message: "max retries reached"}
	}()
	return out
}
