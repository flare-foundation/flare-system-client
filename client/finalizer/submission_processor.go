package finalizer

import (
	"encoding/hex"

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
		return nil
	}
	payloads, err := ExtractPayloads(inputBytes)
	if err != nil {
		// if input cannot be decoded, it is not a valid submission and should be skipped
		logger.Infof("Invalid submitSignatures input sent by %s: %v, skipping", tx.FromAddress, err)
		return nil
	}

	signaturePayloads := []*submitSignaturesPayload{}
	for i := range payloads {
		signaturePayload := new(submitSignaturesPayload)
		err := signaturePayload.FromSignedPayload(payloads[i])
		if err != nil {
			// if input cannot be decoded, it is not a valid submission and should be skipped
			logger.Infof("Invalid submitSignatures payload sent by %s: %v, skipping", tx.FromAddress, err)
		} else {
			signaturePayloads = append(signaturePayloads, signaturePayload)
		}
	}

	if len(signaturePayloads) > 0 {
		err = c.ProcessSubmissionData(signaturePayloads)
		if err != nil {
			// retry the full range, error occurs when the corresponding signing policy
			// is not yet available
			return err
		}
	}
	return nil
}

func (c *client) ProcessSubmissionData(payloads []*submitSignaturesPayload) error {
	for _, payloadItem := range payloads {
		if payloadItem.votingRoundID < c.finalizerContext.startingVotingRound {
			continue
		}

		// Skip if voting round is in the future
		if !c.checkVotingRoundTime(payloadItem.votingRoundID) {
			logger.Debugf("ProcessSubmissionData: Ignoring submitted signature for voting round %d, protocolID  %d - round not started yet", payloadItem.votingRoundID, payloadItem.protocolID)
			continue
		}
		sp, threshold := c.signingPolicyData(payloadItem.votingRoundID)
		if sp == nil {
			oldestSP := c.signingPolicyStorage.OldestStored()
			if oldestSP != nil && payloadItem.votingRoundID < oldestSP.StartVotingRoundID {
				// This is a submission for an old voting round, skip it
				logger.Warnf("ProcessSubmissionData: Ignoring submitted signature for voting round %d, protocolID  %d - before policy startVotingRoundID", payloadItem.votingRoundID, payloadItem.protocolID)
				continue
			}
			logger.Panicf("ProcessSubmissionData: no signing policy found for voting round %d. Storage is empty: %v", payloadItem.votingRoundID, c.signingPolicyStorage.OldestStored() == nil) // this stops the whole fsp client, it only happens if there is no signing policy in the storage.
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
