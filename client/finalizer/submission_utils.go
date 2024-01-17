package finalizer

import (
	"encoding/hex"
	"flare-tlc/client/shared"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type submissionContractClient struct {
	address common.Address
}

type submissionListenerResponse struct {
	payload   []*submitterPayloadItem
	timestamp int64
}

func NewSubmissionContractClient(address common.Address) *submissionContractClient {
	return &submissionContractClient{
		address: address,
	}
}

func (s *submissionContractClient) SubmissionTxListener(db *gorm.DB, startTime time.Time) <-chan submissionListenerResponse {
	out := make(chan submissionListenerResponse, listenerBufferSize)
	selectorBytes := chain.FunctionSelector("submitSignatures()")
	selector := hex.EncodeToString(selectorBytes[:])
	go func() {
		ticker := time.NewTicker(shared.ListenerInterval)
		eventRangeStart := startTime.Unix()
		for {
			<-ticker.C
			now := time.Now().Unix()
			txs, err := database.FetchTransactionsByAddressAndSelector(db, s.address.Hex(), selector, eventRangeStart, now)
			if err != nil {
				logger.Error("Error fetching transactions %v", err)
				continue
			}
			for _, tx := range txs {
				inputBytes, err := hex.DecodeString(tx.Input)
				if err != nil {
					logger.Error("Error decoding input %v", err)
					break
				}
				payload, err := DecodeSubmitterPayload(inputBytes)
				if err != nil {
					logger.Error("Error parsing payload %v", err)
					break
				}
				out <- submissionListenerResponse{
					payload:   payload,
					timestamp: int64(tx.Timestamp),
				}
				eventRangeStart = int64(tx.Timestamp)
			}
		}
	}()
	return out
}
