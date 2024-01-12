package shared

import "time"

const (
	ListenerInterval time.Duration = 2 * time.Second // TODO: change to 10 seconds or read from config
	MaxTxSendRetries int           = 4
	TxRetryInterval  time.Duration = 5 * time.Second
)
