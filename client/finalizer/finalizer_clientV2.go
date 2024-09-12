package finalizer

import (
	"context"
	"encoding/hex"
	clientContext "flare-fsc/client/context"
	"flare-fsc/client/shared"
	"flare-fsc/config"
	"flare-fsc/database"
	"flare-fsc/logger"
	"flare-fsc/utils/contracts/relay"
	"flare-fsc/utils/credentials"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

const minRoundsStored uint32 = 10

type finalizerClientV2 struct {
	db finalizerDB

	relayClient          *relayContractClient
	submissionClient     *submissionContractClient
	signingPolicyStorage *signingPolicyStorage
	messages             <-chan shared.ProtocolMessage
	finalizationStorage  *finalizationStorage
	queueProcessor       *finalizerQueueProcessorV2

	finalizerContext *finalizerContext
}

func NewFinalizerClientV2(ctx clientContext.ClientContext, messageChannel <-chan shared.ProtocolMessage) (*finalizerClientV2, error) {
	cfg := ctx.Config()
	if !cfg.Clients.EnabledFinalizer {
		return nil, nil
	}

	chainCfg := cfg.ChainConfig()
	ethClient, err := chainCfg.DialETH()
	if err != nil {
		return nil, err
	}

	relay, err := relay.NewRelay(cfg.ContractAddresses.Relay, ethClient)
	if err != nil {
		return nil, errors.Wrap(err, "error creating relay contract")
	}
	finalizerContext, err := newFinalizerContext(cfg, relay)
	if err != nil {
		return nil, err
	}

	senderPkString, err := config.PrivateKeyFromConfig(cfg.Credentials.SigningPolicyPrivateKeyFile,
		cfg.Credentials.SigningPolicyPrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "error reading sender private key")
	}
	txOpts, senderPk, err := credentials.CredentialsFromPrivateKey(senderPkString, chainCfg.ChainID)
	if err != nil {
		return nil, errors.Wrap(err, "error creating sender register tx opts")
	}
	relayClient, err := NewRelayContractClient(
		ethClient,
		cfg.ContractAddresses.Relay,
		senderPk,
		txOpts.From,
	)
	if err != nil {
		return nil, err
	}
	submissionClient := NewSubmissionContractClient(cfg.ContractAddresses.Submission)
	finalizationStorage := newFinalizationStorage()

	db := finalizerDBImpl{client: ctx.DB()}

	return &finalizerClientV2{
		db:                   db,
		relayClient:          relayClient,
		signingPolicyStorage: newSigningPolicyStorage(),
		messages:             messageChannel,
		finalizationStorage:  finalizationStorage,
		submissionClient:     submissionClient,
		queueProcessor:       newFinalizerQueueProcessorV2(db, finalizationStorage, relayClient, finalizerContext),
		finalizerContext:     finalizerContext,
	}, nil
}

func (c *finalizerClientV2) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)

	startTime := time.Now().Add(-c.finalizerContext.startTimeOffset)
	startTime, err := c.fetchExistingSigningPolicies(ctx, startTime)
	if err != nil {
		return err
	}

	eg.Go(func() error {
		return c.runSigningPolicyInitializedListener(ctx, startTime)
	})
	eg.Go(func() error {
		return c.submissionClient.SubmissionTxListenerV2(ctx, c.db, startTime, c)
	})
	eg.Go(func() error {
		return c.queueProcessor.Run(ctx)
	})
	eg.Go(func() error {
		return c.messagesChannelListener(ctx)
	})

	return eg.Wait()
}

func (c *finalizerClientV2) fetchExistingSigningPolicies(
	_ context.Context, startTime time.Time,
) (time.Time, error) {
	// Read current signing policies from the database and add them to the storage
	spList, err := c.relayClient.FetchSigningPolicies(c.db, startTime.Unix(), time.Now().Unix())
	if err != nil {
		return startTime, err
	}
	for _, sp := range spList {
		policy := newSigningPolicy(sp.policyData)
		if policy.rewardEpochID < c.finalizerContext.startingRewardEpoch {
			continue
		}
		if err := c.signingPolicyStorage.Add(policy); err != nil {
			return startTime, err
		}
	}
	logger.Info("Added %d signing policies", len(spList))

	if len(spList) > 0 {
		return time.Unix(spList[len(spList)-1].timestamp, 0), nil
	}

	return startTime, nil
}

func (c *finalizerClientV2) runSigningPolicyInitializedListener(ctx context.Context, startTime time.Time) error {
	spListener := c.relayClient.SigningPolicyInitializedListener(c.db, startTime)
	for {
		var dbPolicy signingPolicyListenerResponse
		select {
		case dbPolicy = <-spListener:
		case <-ctx.Done():
			logger.Info("Signing policy initialized listener stopped")
			return ctx.Err()
		}

		policy := newSigningPolicy(dbPolicy.policyData)
		if policy.rewardEpochID < c.finalizerContext.startingRewardEpoch {
			continue
		}
		if err := c.signingPolicyStorage.Add(policy); err != nil {
			logger.Warn("Error adding signing policy %v", err)
		}
		logger.Info("New signing policy received for epoch %v", policy.rewardEpochID)
		c.signingPolicyStorage.RemoveBefore(c.finalizationStorage.lowestRoundStored)
	}
}

func (c *finalizerClientV2) ProcessSubmissionData(slr submissionListenerResponseV2) error {
	for _, payloadItem := range slr.payloads {
		if payloadItem.votingRoundID < c.finalizerContext.startingVotingRound {
			logger.Debug("Ignoring submitted signature for voting round %d, protocolID  %d - before startingVotingRound", payloadItem.votingRoundID, payloadItem.protocolID)
			continue
		}

		// Skip if voting round is in the future
		if !c.checkVotingRoundTime(payloadItem.votingRoundID) {
			continue
		}
		sp, _ := c.signingPolicyData(payloadItem.votingRoundID)
		if sp == nil {
			first := c.signingPolicyStorage.First()
			if first != nil && payloadItem.votingRoundID < first.startVotingRoundID {
				// This is a submission for an old voting round, skip it
				logger.Debug("Ignoring submitted signature for voting round %d, protocolID  %d - before policy startVotingRoundID", payloadItem.votingRoundID, payloadItem.protocolID)
				continue
			}
			return fmt.Errorf("no signing policy found for voting round %d", payloadItem.votingRoundID)
		}
		finalizationReady, err := c.finalizationStorage.addPayload(payloadItem, sp)
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
				c.finalizationStorage.RemoveRoundsBefore(finalizationReady.votingRoundID - minRoundsStored)
			}
		}
	}
	return nil
}

func (c *finalizerClientV2) ProcessTransaction(tx database.Transaction) error {
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
		signaturePayload, err := decodeSignedPayloadV2(payloads[i])

		if err != nil {
			// if input cannot be decoded, it is not a valid submission and should be skipped
			logger.Info("Invalid submitSignatures payload sent by %s: %v, skipping", tx.FromAddress, err)

		}
		signaturePayloads = append(signaturePayloads, &signaturePayload)
	}

	if len(signaturePayloads) > 0 {
		err = c.ProcessSubmissionData(submissionListenerResponseV2{
			payloads:  signaturePayloads,
			timestamp: int64(tx.Timestamp),
		})
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

// return signing policy and voting threshold for the given voting round
func (c *finalizerClientV2) signingPolicyData(votingRoundID uint32) (*signingPolicy, uint16) {
	sp, last := c.signingPolicyStorage.GetForVotingRound(votingRoundID)
	if sp == nil {
		return nil, 0
	}
	if !last {
		return sp, sp.threshold
	}
	endVotingEpoch := c.finalizerContext.rewardEpoch.EndEpoch(sp.rewardEpochID)
	end := c.finalizerContext.votingEpoch.EndTime(endVotingEpoch)

	if time.Now().Before(end) {
		return sp, sp.threshold
	} else {
		return sp, uint16((uint32(sp.voters.TotalWeight()) * 60) / 100)
	}
}

// Return true if voting round is not in the future, i.e., is <= the current voting round
func (c *finalizerClientV2) checkVotingRoundTime(votingRoundID uint32) bool {
	currentEpochID := c.finalizerContext.votingEpoch.EpochIndex(time.Now())
	return votingRoundID <= uint32(currentEpochID)
}

func (c *finalizerClientV2) messagesChannelListener(ctx context.Context) error {
	for {
		var protocolMessage shared.ProtocolMessage

		select {
		case protocolMessage = <-c.messages:
		case <-ctx.Done():
			logger.Info("Message Channel Listener stopped")
			return ctx.Err()
		}

		sp, _ := c.signingPolicyData(protocolMessage.VotingRoundID)

		finalizationReady, err := c.finalizationStorage.AddMessage(&protocolMessage, sp)

		if err != nil {
			logger.Debug("Ignoring submitted message for protocol %d, round %d: %v", protocolMessage.ProtocolID, protocolMessage.VotingRoundID, err)
			continue
		}

		if finalizationReady.thresholdReached {
			logger.Info("Threshold reached for protocol %d in voting round %d with hash %v", finalizationReady.protocolID, finalizationReady.votingRoundID)
			c.queueProcessor.Add(&finalizationReady, sp.seed)
		}
	}
}
