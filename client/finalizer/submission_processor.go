package finalizer

import (
	"encoding/hex"
	"flare-fsc/database"
	"flare-fsc/logger"
	"fmt"
)

type submissionProcessor interface {
	// ProcessTransaction expects submitSignatures transaction. It extracts and processes signature payloads.
	// Return error if the transaction was not processed and needs a retry.
	//
	// Must be able to handle duplicates.
	ProcessTransaction(database.Transaction) error
}

// ProcessTransaction
func (c *finalizerClient) ProcessTransaction(tx database.Transaction) error {
	inputBytes, err := hex.DecodeString(tx.Input)
	if err != nil {
		logger.Info("Invalid submitSignatures tx sent by %s: %v, skipping", tx.FromAddress, err)
	}
	payloads, err := ExtractPayloads(inputBytes)
	if err != nil {
		// if input cannot be decoded, it is not a valid submission and should be skipped
		logger.Info("Invalid submitSignatures input sent by %s: %v, skipping", tx.FromAddress, err)
	}

	signaturePayloads := []*submitSignaturesPayload{}
	for i := range payloads {
		signaturePayload := new(submitSignaturesPayload)
		err := signaturePayload.FromSignedPayload(payloads[i])
		if err != nil {
			// if input cannot be decoded, it is not a valid submission and should be skipped
			logger.Info("Invalid submitSignatures payload sent by %s: %v, skipping", tx.FromAddress, err)

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

func (c *finalizerClient) ProcessSubmissionData(payloads []*submitSignaturesPayload) error {
	for _, payloadItem := range payloads {
		if payloadItem.votingRoundID < c.finalizerContext.startingVotingRound {
			logger.Debug("Ignoring submitted signature for voting round %d, protocolID  %d - before startingVotingRound %d", payloadItem.votingRoundID, payloadItem.protocolID, c.finalizerContext.startingVotingRound)
			continue
		}

		// Skip if voting round is in the future
		if !c.checkVotingRoundTime(payloadItem.votingRoundID) {
			logger.Debug("Ignoring submitted signature for voting round %d, protocolID  %d - round not started yet", payloadItem.votingRoundID, payloadItem.protocolID)

			continue
		}
		sp, threshold := c.signingPolicyData(payloadItem.votingRoundID)
		if sp == nil {
			first := c.signingPolicyStorage.First()
			if first != nil && payloadItem.votingRoundID < first.startVotingRoundID {
				// This is a submission for an old voting round, skip it
				logger.Debug("Ignoring submitted signature for voting round %d, protocolID  %d - before policy startVotingRoundID", payloadItem.votingRoundID, payloadItem.protocolID)
				continue
			}
			return fmt.Errorf("no signing policy found for voting round %d", payloadItem.votingRoundID)
		}
		finalizationReady, err := c.finalizationStorage.addPayload(payloadItem, sp, threshold)
		if err != nil {
			// Error is non-fatal, skip this submission
			logger.Debug("Ignoring submitted signature for voting round %d, protocolID  %d - %v", payloadItem.votingRoundID, payloadItem.protocolID, err)
			continue
		}

		if finalizationReady.thresholdReached {
			logger.Info("Threshold reached for protocol %d in voting round %d", finalizationReady.protocolID, finalizationReady.votingRoundID)
			c.queueProcessor.Add(&finalizationReady, sp.seed)

			//clean old rounds
			if finalizationReady.votingRoundID > minRoundsStored {
				c.finalizationStorage.RemoveRoundsBefore(finalizationReady.votingRoundID - minRoundsStored) // remove that are at least minRoundStored + 1 older then the one that has been finalized
			}
		}
	}
	return nil
}
