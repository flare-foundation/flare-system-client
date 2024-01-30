package shared

import "time"

const (
	ListenerInterval time.Duration = 2 * time.Second
	MaxTxSendRetries int           = 4
	TxRetryInterval  time.Duration = 5 * time.Second
)
