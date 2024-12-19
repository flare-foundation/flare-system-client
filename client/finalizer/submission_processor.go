package finalizer

import (
	"encoding/hex"
	"fmt"

	"github.com/flare-foundation/go-flare-common/pkg/database"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

type submissionProcessor interface {
	// ProcessTransaction expects submitSignatures transaction. It extracts and processes signature payloads.
	// Return error if the transaction was not processed and needs a retry.
	//
	// Must be able to handle duplicates.
	ProcessTransaction(database.Transaction) error
}

// ProcessTransaction
func (c *client) ProcessTransaction(tx database.Transaction) error {
	inputBytes, err := hex.DecodeString(tx.Input)
	if err != nil {
		logger.Infof("Invalid submitSignatures tx sent by %s: %v, skipping", tx.FromAddress, err)
	}
	payloads, err := ExtractPayloads(inputBytes)
	if err != nil {
		// if input cannot be decoded, it is not a valid submission and should be skipped
		logger.Infof("Invalid submitSignatures input sent by %s: %v, skipping", tx.FromAddress, err)
	}

	signaturePayloads := []*submitSignaturesPayload{}
	for i := range payloads {
		signaturePayload := new(submitSignaturesPayload)
		err := signaturePayload.FromSignedPayload(payloads[i])
		if err != nil {
			// if input cannot be decoded, it is not a valid submission and should be skipped
			logger.Infof("Invalid submitSignatures payload sent by %s: %v, skipping", tx.FromAddress, err)

		}
		signaturePayloads = append(signaturePayloads, signaturePayload)
	}

	if len(signaturePayloads) > 0 {
		err = c.ProcessSubmissionData(signaturePayloads)
		if err != nil {
			// retry the full range, error occurs when the corresponding signing policy
			// is not yet available
			return err
		}
	}
	// -1 for overlap in case of an error and retry above
	// processor should be able to handle duplicates
	return nil
}

func (c *client) ProcessSubmissionData(payloads []*submitSignaturesPayload) error {
	for _, payloadItem := range payloads {
		if payloadItem.votingRoundID < c.finalizerContext.startingVotingRound {
			continue
		}

		// Skip if voting round is in the future
		if !c.checkVotingRoundTime(payloadItem.votingRoundID) {
			logger.Debugf("Ignoring submitted signature for voting round %d, protocolID  %d - round not started yet", payloadItem.votingRoundID, payloadItem.protocolID)
			continue
		}
		sp, threshold := c.signingPolicyData(payloadItem.votingRoundID)
		if sp == nil {
			oldestSP := c.signingPolicyStorage.OldestStored()
			if oldestSP != nil && payloadItem.votingRoundID < oldestSP.StartVotingRoundID {
				// This is a submission for an old voting round, skip it
				logger.Debugf("Ignoring submitted signature for voting round %d, protocolID  %d - before policy startVotingRoundID", payloadItem.votingRoundID, payloadItem.protocolID)
				continue
			}
			return fmt.Errorf("no signing policy found for voting round %d", payloadItem.votingRoundID) // this stops the whole fsp client
		}
		finalizationReady, err := c.finalizationStorage.addPayload(payloadItem, sp, threshold)
		if err != nil {
			// Error is non-fatal, skip this submission
			logger.Debugf("Ignoring submitted signature for voting round %d, protocolID  %d - %v", payloadItem.votingRoundID, payloadItem.protocolID, err)
			continue
		}

		if finalizationReady.thresholdReached {
			logger.Infof("Threshold reached for protocol %d in voting round %d", finalizationReady.protocolID, finalizationReady.votingRoundID)
			c.queueProcessor.Add(&finalizationReady, sp.Seed)

			//clean old rounds
			if finalizationReady.votingRoundID > minRoundsStored {
				c.finalizationStorage.RemoveRoundsBefore(finalizationReady.votingRoundID - minRoundsStored) // remove that are at least minRoundStored + 1 older then the one that has been finalized
			}
		}
	}
	return nil
}
