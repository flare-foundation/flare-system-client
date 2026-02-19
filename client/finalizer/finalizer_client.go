package finalizer

import (
	"context"
	"fmt"
	"time"

	clientContext "github.com/flare-foundation/flare-system-client/client/context"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/config"
	"github.com/flare-foundation/flare-system-client/utils/credentials"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/flare-foundation/go-flare-common/pkg/policy"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
)

const minRoundsStored uint32 = 10

// client manages finalization tasks:
//   - collects messages and signatures
//   - prepares and submits finalization transactions
type client struct {
	db finalizerDB

	relayClient          *relayContractClient          // reading and writing to the relay contract for the finalization
	submissionListener   *submissionListener           // listening for new transactions on the submission contract
	signingPolicyStorage *policy.Storage               // storing data about participants in the protocol
	messages             <-chan shared.ProtocolMessage // channel to receive data from the submitter (protocol package)
	finalizationStorage  *finalizationStorage          // storing data for the finalization
	queueProcessor       *finalizerQueueProcessor      // implementation of a processor finalizing data from a queue

	finalizerContext *finalizerContext
}

// NewClient creates a new client that manages finalizations.
//
// messageChannel is used to receive messages from protocol.client.
func NewClient(ctx clientContext.ClientContext, messageChannel <-chan shared.ProtocolMessage) (*client, error) {
	cfg := ctx.Config()
	if !cfg.Clients.EnabledFinalizer {
		return nil, nil
	}

	chainCfg := cfg.ChainConfig()
	ethClient, err := chainCfg.DialETH()
	if err != nil {
		return nil, err
	}

	relayContract, err := relay.NewRelay(cfg.ContractAddresses.Relay, ethClient)
	if err != nil {
		return nil, errors.Wrap(err, "error creating relay contract")
	}
	finalizerContext, err := newFinalizerContext(cfg, relayContract)
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
		&cfg.RelayGas,
	)
	if err != nil {
		return nil, err
	}
	submissionListener := NewSubmissionListener(cfg.ContractAddresses.Submission)
	finalizationStorage := newFinalizationStorage()

	db := finalizerDBImpl{client: ctx.DB()}

	return &client{
		db:                   db,
		relayClient:          relayClient,
		signingPolicyStorage: policy.NewStorage(),
		messages:             messageChannel,
		finalizationStorage:  finalizationStorage,
		submissionListener:   submissionListener,
		queueProcessor:       newFinalizerQueueProcessor(db, finalizationStorage, relayClient, finalizerContext),
		finalizerContext:     finalizerContext,
	}, nil
}

// Run runs the client. Should be called in a goroutine.
func (c *client) Run(ctx context.Context) error {
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

func (c *client) fetchExistingSigningPolicies(
	_ context.Context, startTime time.Time,
) (time.Time, error) {
	// Read current signing policies from the database and add them to the storage
	spList, err := c.relayClient.FetchSigningPolicies(c.db, startTime.Unix(), time.Now().Unix())
	if err != nil {
		return startTime, err
	}
	for _, sp := range spList {
		newPolicy := policy.NewSigningPolicy(sp.policyData, nil)
		if newPolicy.RewardEpochID < c.finalizerContext.startingRewardEpoch {
			continue
		}
		if err := c.signingPolicyStorage.Add(newPolicy); err != nil {
			return startTime, err
		}
	}
	logger.Infof("Added %d signing policies", len(spList))

	if len(spList) > 0 {
		return time.Unix(spList[len(spList)-1].timestamp, 0), nil
	}

	return startTime, nil
}

func (c *client) runSigningPolicyInitializedListener(ctx context.Context, startTime time.Time) error {
	spListener := c.relayClient.SigningPolicyInitializedListener(c.db, startTime)
	for {
		var dbPolicy signingPolicyListenerResponse
		select {
		case dbPolicy = <-spListener:
		case <-ctx.Done():
			logger.Infof("Signing policy initialized listener stopped")
			return ctx.Err()
		}

		policy := policy.NewSigningPolicy(dbPolicy.policyData, nil)
		if policy.RewardEpochID < c.finalizerContext.startingRewardEpoch {
			continue
		}
		if err := c.signingPolicyStorage.Add(policy); err != nil {
			logger.Warnf("Error adding signing policy %v", err)
		}
		logger.Infof("New signing policy received for epoch %v", policy.RewardEpochID)

		c.signingPolicyStorage.RemoveBefore(c.finalizationStorage.lowestRoundStored) // remove signingPolicies that will never be used again
	}
}

// signingPolicyData returns signing policy and voting threshold for the given votingRoundID.
//
// If the signing policy was expected to end before votingRoundID but it was prolonged, the threshold is raised to 60% of total weight.
func (c *client) signingPolicyData(votingRoundID uint32) (*policy.SigningPolicy, uint16) {
	sp, last := c.signingPolicyStorage.ForVotingRound(votingRoundID)
	if sp == nil {
		return nil, 0
	}
	if !last {
		return sp, sp.Threshold
	}
	expectedEnd := c.finalizerContext.rewardEpoch.EndEpoch(sp.RewardEpochID)

	if int64(votingRoundID) < expectedEnd {
		return sp, sp.Threshold
	} else {
		return sp, uint16((uint32(sp.Voters.TotalWeight) * 60) / 100) // if the rewardEpoch extends beyond the expected end, the threshold is raised to 60%.
	}
}

// checkVotingRoundTime returns true if votingRoundID is not in the future, i.e., is <= the current voting round
func (c *client) checkVotingRoundTime(votingRoundID uint32) bool {
	currentEpochID := c.finalizerContext.votingRoundTiming.EpochIndex(time.Now())
	return votingRoundID <= uint32(currentEpochID)
}

// messagesChannelListener listens to the messages from the protocol message channel and adds them to the finalizationStorage.
func (c *client) messagesChannelListener(ctx context.Context) error {
	for {
		var protocolMessage shared.ProtocolMessage

		select {
		case protocolMessage = <-c.messages:
		case <-ctx.Done():
			logger.Infof("Message Channel Listener stopped")
			return ctx.Err()
		}

		sp, threshold := c.signingPolicyData(protocolMessage.VotingRoundID)

		if sp == nil {
			oldestSP := c.signingPolicyStorage.OldestStored()
			if oldestSP != nil && protocolMessage.VotingRoundID < oldestSP.StartVotingRoundID {
				// This is a submission for an old voting round, skip it
				logger.Debugf("Ignoring message for voting round %d, protocolID  %d - before policy startVotingRoundID", protocolMessage.VotingRoundID, protocolMessage.ProtocolID)
				continue
			}
			return fmt.Errorf("no signing policy found for voting round %d", protocolMessage.VotingRoundID) // this stops the whole fsp client
		}
		finalizationReady, err := c.finalizationStorage.AddMessage(&protocolMessage, sp, threshold)

		if err != nil {
			logger.Debugf("Ignoring submitted message for protocol %d, round %d: %v", protocolMessage.ProtocolID, protocolMessage.VotingRoundID, err)
			continue
		}

		if finalizationReady.thresholdReached {
			logger.Infof("Threshold reached for protocol %d in voting round %d with hash %v", finalizationReady.protocolID, finalizationReady.votingRoundID)
			c.queueProcessor.Add(&finalizationReady, sp.Seed)
		}
	}
}
