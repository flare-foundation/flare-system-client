package shared

import (
	"encoding/hex"
	"flare-tlc/database"
	"flare-tlc/utils/chain"

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
	return &types.Log{
		Topics: []common.Hash{
			chain.ParseTopic(dbLog.Topic0),
			chain.ParseTopic(dbLog.Topic1),
			chain.ParseTopic(dbLog.Topic2),
			chain.ParseTopic(dbLog.Topic3),
		},
		Data: data,
		// Other fields are not used by log decoder
	}, nil
}
