package finalizer

import (
	"context"
	"flare-fsc/client/shared"
	"flare-fsc/logger"
	"flare-fsc/utils/contracts/submission"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type submissionContractClient struct {
	address common.Address
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
	processor submitterProcessor,
) error {
	submissionABI, err := submission.SubmissionMetaData.GetAbi()
	if err != nil {
		// Should not happen, unhandled errors will cause a panic further up.
		return err
	}

	selector := submissionABI.Methods["submitSignatures"].ID
	ticker := time.NewTicker(shared.ListenerInterval)

	txs, err := db.FetchTransactionsByAddressAndSelector(s.address, selector, startTime.Unix(), time.Now().Unix())
	if err != nil {
		logger.Error("Error fetching transactions %v", err)
	}

	lastBlockChecked := uint64(0)

	for _, tx := range txs {
		err := processor.ProcessTransaction(tx)
		if err != nil {
			logger.Warn("Error processing submitSignatures payload sent by %s: %v", tx.FromAddress, err)
		}

		if tx.BlockNumber > uint64(lastBlockChecked) {
			lastBlockChecked = tx.BlockNumber
		}
	}

	for {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			logger.Info("Submission tx listener stopped")
			return ctx.Err()
		}

		txs, err := db.FetchTransactionsByAddressAndSelectorFromBlockNumber(s.address, selector, int64(lastBlockChecked))
		if err != nil {
			logger.Error("Error fetching transactions %v", err)
			continue
		}
		for _, tx := range txs {
			err := processor.ProcessTransaction(tx)
			if err != nil {
				logger.Warn("Error processing submitSignatures payload sent by %s: %v", tx.FromAddress, err)
			}
			if tx.BlockNumber > uint64(lastBlockChecked) {
				lastBlockChecked = tx.BlockNumber
			}
		}
	}
}
