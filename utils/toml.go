package utils

import (
	"errors"
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
}

func (t *Timestamp) UnmarshalText(text []byte) error {
	// Try to parse as RFC3339
	parsed, err := time.Parse(time.RFC3339, string(text))
	if err == nil {
		t.Time = parsed
		return nil
	}
	// Try to parse as Unix timestamp (as integer, in seconds)
	unixTimestamp, err := strconv.ParseInt(string(text), 10, 64)
	if err == nil {
		t.Time = time.Unix(unixTimestamp, 0)
		return nil
	}
	return errors.New("timestamp must be in RFC3339 or Unix timestamp format")
}
