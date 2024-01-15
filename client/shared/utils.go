package shared

import (
	"encoding/hex"
	"flare-tlc/database"
	"flare-tlc/logger"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

// ExistsAsSubstring returns true if any of the strings in the slice is a substring of s
func ExistsAsSubstring(slice []string, s string) bool {
	for _, item := range slice {
		if strings.Contains(s, item) {
			return true
		}
	}
	return false
}

// DBLogToChainLog converts a database log to a chain log for use in the log decoder
// It only converts the fields used by the log decoder (Topics and Data)
func ConvertDatabaseLogToChainLog(dbLog database.Log) (*types.Log, error) {
	data, err := hex.DecodeString(dbLog.Data)
	if err != nil {
		return nil, err
	}

	var topics []common.Hash

	if dbLog.Topic0 != "NULL" {
		topics = append(topics, common.HexToHash(dbLog.Topic0))
	}
	if dbLog.Topic1 != "NULL" {
		topics = append(topics, common.HexToHash(dbLog.Topic1))
	}
	if dbLog.Topic2 != "NULL" {
		topics = append(topics, common.HexToHash(dbLog.Topic2))
	}
	if dbLog.Topic3 != "NULL" {
		topics = append(topics, common.HexToHash(dbLog.Topic3))
	}
	return &types.Log{
		Topics: topics,
		Data:   data,
		// Other fields are not used by log decoder
	}, nil
}
