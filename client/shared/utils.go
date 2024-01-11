package shared

import (
	"encoding/hex"
	"flare-tlc/database"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

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
