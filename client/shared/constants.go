package shared

import "time"

const (
	ListenerInterval      time.Duration = 2 * time.Second
	EventListenerInterval time.Duration = 5 * time.Second
	MaxTxSendRetries      int           = 10
	MaxTxSendRetriesLong  int           = 20
	TxRetryInterval       time.Duration = 5 * time.Second
	TxRetryIntervalLong   time.Duration = 20 * time.Second
)
