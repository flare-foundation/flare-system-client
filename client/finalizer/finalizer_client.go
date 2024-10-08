package finalizer

import (
	"context"
	clientContext "flare-fsc/client/context"
	"flare-fsc/client/shared"
	"flare-fsc/config"
	"flare-fsc/logger"
	"flare-fsc/utils/contracts/relay"
	"flare-fsc/utils/credentials"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

const minRoundsStored uint32 = 10

type finalizerClient struct {
	db finalizerDB

	relayClient          *relayContractClient          // reading and writing to the relay contract for the finalization
	submissionListener   *submissionListener           // listening for new transactions on the submission contract
	signingPolicyStorage *signingPolicyStorage         // storing data about participants in the protocol
	messages             <-chan shared.ProtocolMessage // channel to receive data from the submitter (protocol package)
	finalizationStorage  *finalizationStorage          // storing data for the finalization
	queueProcessor       *finalizerQueueProcessor      // implementation of a processor finalizing data from a queue

	finalizerContext *finalizerContext
}

func NewFinalizerClient(ctx clientContext.ClientContext, messageChannel <-chan shared.ProtocolMessage) (*finalizerClient, error) {
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
	submissionListener := NewSubmissionListener(cfg.ContractAddresses.Submission)
	finalizationStorage := newFinalizationStorage()

	db := finalizerDBImpl{client: ctx.DB()}

	return &finalizerClient{
		db:                   db,
		relayClient:          relayClient,
		signingPolicyStorage: newSigningPolicyStorage(),
		messages:             messageChannel,
		finalizationStorage:  finalizationStorage,
		submissionListener:   submissionListener,
		queueProcessor:       newFinalizerQueueProcessor(db, finalizationStorage, relayClient, finalizerContext),
		finalizerContext:     finalizerContext,
	}, nil
}

func (c *finalizerClient) Run(ctx context.Context) error {
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
		return c.submissionListener.SubmissionTxListen(ctx, c.db, startTime, c)
	})
	eg.Go(func() error {
		return c.queueProcessor.Run(ctx)
	})
	eg.Go(func() error {
		return c.messagesChannelListener(ctx)
	})

	return eg.Wait()
}

func (c *finalizerClient) fetchExistingSigningPolicies(
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

func (c *finalizerClient) runSigningPolicyInitializedListener(ctx context.Context, startTime time.Time) error {
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
		c.signingPolicyStorage.RemoveBefore(c.finalizationStorage.lowestRoundStored) // remove signingPolicies that will never be used again
	}
}

// signingPolicyData return signing policy and voting threshold for the given votingRoundID.
func (c *finalizerClient) signingPolicyData(votingRoundID uint32) (*signingPolicy, uint16) {
	sp, last := c.signingPolicyStorage.GetForVotingRound(votingRoundID)
	if sp == nil {
		return nil, 0
	}
	if !last {
		return sp, sp.threshold
	}
	expectedEnd := c.finalizerContext.rewardEpoch.EndEpoch(sp.rewardEpochID)

	if int64(votingRoundID) < expectedEnd {
		return sp, sp.threshold
	} else {
		return sp, uint16((uint32(sp.voters.TotalWeight()) * 60) / 100) // if the rewardEpoch extends beyond the expected end, the threshold is raised to 60%.
	}
}

// Return true if voting round is not in the future, i.e., is <= the current voting round
func (c *finalizerClient) checkVotingRoundTime(votingRoundID uint32) bool {
	currentEpochID := c.finalizerContext.votingEpoch.EpochIndex(time.Now())
	return votingRoundID <= uint32(currentEpochID)
}

func (c *finalizerClient) messagesChannelListener(ctx context.Context) error {
	for {
		var protocolMessage shared.ProtocolMessage

		select {
		case protocolMessage = <-c.messages:
		case <-ctx.Done():
			logger.Info("Message Channel Listener stopped")
			return ctx.Err()
		}

		sp, threshold := c.signingPolicyData(protocolMessage.VotingRoundID)

		// TODO check this
		if sp == nil {
			first := c.signingPolicyStorage.First()
			if first != nil && protocolMessage.VotingRoundID < first.startVotingRoundID {
				// This is a submission for an old voting round, skip it
				logger.Debug("Ignoring message for voting round %d, protocolID  %d - before policy startVotingRoundID", protocolMessage.VotingRoundID, protocolMessage.ProtocolID)
				continue
			}
			return fmt.Errorf("no signing policy found for voting round %d", protocolMessage.VotingRoundID) // should this really return?
		}
		finalizationReady, err := c.finalizationStorage.AddMessage(&protocolMessage, sp, threshold)

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
