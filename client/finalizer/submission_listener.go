package finalizer

import (
	"context"
	"flare-fsc/client/shared"
	"flare-fsc/utils/contracts/submission"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/database"

	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/logger"
)

type submissionListener struct {
	address common.Address
}

func NewSubmissionListener(address common.Address) *submissionListener {
	return &submissionListener{
		address: address,
	}
}

func (s *submissionListener) SubmissionTxListen(
	ctx context.Context,
	db finalizerDB,
	startTime time.Time,
	processor submissionProcessor,
) error {
	submissionABI, err := submission.SubmissionMetaData.GetAbi()
	if err != nil {
		// Should not happen, unhandled errors will cause a panic further up.
		return err
	}

	selector := submissionABI.Methods["submitSignatures"].ID
	lastBlockChecked := uint64(0)
	var txs []database.Transaction

	ticker := time.NewTicker(shared.ListenerInterval)
	// a for loop to guarantee to get at least one transaction
	for {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			logger.Info("Submission tx listener stopped")
			return ctx.Err()
		}
		txs, err = db.FetchTransactionsByAddressAndSelector(s.address, selector, startTime.Unix(), time.Now().Unix())
		if err != nil {
			logger.Errorf("Error fetching transactions %v", err)
		}
		if len(txs) > 0 {
			break
		}
	}

	for _, tx := range txs {
		err := processor.ProcessTransaction(tx)
		if err != nil {
			logger.Warnf("Error processing submitSignatures payload sent by %s: %v", tx.FromAddress, err)
		}

		if tx.BlockNumber > uint64(lastBlockChecked) {
			lastBlockChecked = tx.BlockNumber
		}
	}

	ticker = time.NewTicker(shared.ListenerInterval)
	for {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			logger.Info("Submission tx listener stopped")
			return ctx.Err()
		}

		txs, err := db.FetchTransactionsByAddressAndSelectorFromBlockNumber(s.address, selector, int64(lastBlockChecked))
		if err != nil {
			logger.Errorf("Error fetching transactions %v", err)
			continue
		}
		for _, tx := range txs {
			err := processor.ProcessTransaction(tx)
			if err != nil {
				logger.Warnf("Error processing submitSignatures payload sent by %s: %v", tx.FromAddress, err)
			}
			if tx.BlockNumber > uint64(lastBlockChecked) {
				lastBlockChecked = tx.BlockNumber
			}
		}
	}
}
