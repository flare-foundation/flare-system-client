package finalizer

import (
	"encoding/hex"
	"flare-tlc/client/shared"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils/contracts/submission"
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

type submitterItemProcessor interface {
	// Return error if the submission was not processed and needs a retry
	// Should be able to handle duplicates
	ProcessSubmissionData(submissionListenerResponse) error
}

func NewSubmissionContractClient(address common.Address) *submissionContractClient {
	return &submissionContractClient{
		address: address,
	}
}

func (s *submissionContractClient) SubmissionTxListener(
	db *gorm.DB,
	startTime time.Time,
	processor submitterItemProcessor,
) {
	submissionABI, err := submission.SubmissionMetaData.GetAbi()
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	selectorBytes := submissionABI.Methods["submitSignatures"].ID
	selector := hex.EncodeToString(selectorBytes)
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
				logger.Error("Error decoding input %v, retrying", err) // continue?
				break
			}
			payload, err := DecodeSubmitterPayload(inputBytes)
			if err != nil {
				logger.Error("Error parsing payload %v, retrying", err) // continue?
				break
			}
			err = processor.ProcessSubmissionData(submissionListenerResponse{
				payload:   payload,
				timestamp: int64(tx.Timestamp),
			})
			if err != nil {
				logger.Warn("Error processing payload %v, retrying", err)
				break
			}
			// -1 for a bit of overlap in case of an error and break above
			// processor should be able to handle duplicates
			eventRangeStart = int64(tx.Timestamp) - 1
		}
	}
}
