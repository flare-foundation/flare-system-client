package finalizer

import (
	"context"
	"encoding/hex"
	"flare-tlc/client/shared"
	"flare-tlc/logger"
	"flare-tlc/utils/contracts/submission"
	"time"

	"github.com/ethereum/go-ethereum/common"
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
	ctx context.Context,
	db finalizerDB,
	startTime time.Time,
	processor submitterItemProcessor,
) error {
	submissionABI, err := submission.SubmissionMetaData.GetAbi()
	if err != nil {
		// Should not happen, unhandled errors will cause a panic further up.
		return err
	}

	selector := submissionABI.Methods["submitSignatures"].ID
	ticker := time.NewTicker(shared.ListenerInterval)
	eventRangeStart := startTime.Unix()
	for {
		select {
		case <-ticker.C:
			break

		case <-ctx.Done():
			logger.Info("Submission tx listener stopped")
			return ctx.Err()
		}
		now := time.Now().Unix()
		txs, err := db.FetchTransactionsByAddressAndSelector(s.address, selector, eventRangeStart, now)
		if err != nil {
			logger.Error("Error fetching transactions %v", err)
			continue
		}
		for _, tx := range txs {
			inputBytes, err := hex.DecodeString(tx.Input)
			if err != nil {
				logger.Info("Invalid submitSignatures tx sent by %s: %v, skipping", tx.FromAddress, err)
			}
			payload, err := DecodeSubmitterPayload(inputBytes)
			if err != nil {
				// if input cannot be decoded, it is not a valid submission and should be skipped
				logger.Info("Invalid submitSignatures payload sent by %s: %v, skipping", tx.FromAddress, err)
			}
			if len(payload) > 0 {
				err = processor.ProcessSubmissionData(submissionListenerResponse{
					payload:   payload,
					timestamp: int64(tx.Timestamp),
				})
				if err != nil {
					// retry the full range, error occurs when the corresponding signing policy
					// is not yet available
					logger.Warn("Error processing submitSignatures payload sent by %s: %v, retrying", tx.FromAddress, err)
					break
				}
			}
			// -1 for overlap in case of an error and retry above
			// processor should be able to handle duplicates
			eventRangeStart = int64(tx.Timestamp) - 1
		}
	}
}
