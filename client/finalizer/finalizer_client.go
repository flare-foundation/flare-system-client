package finalizer

import (
	"context"
	"fmt"
	"math"
	"time"

	clientContext "github.com/flare-foundation/flare-system-client/client/context"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/config"
	"github.com/flare-foundation/flare-system-client/utils/credentials"

	"github.com/ethereum/go-ethereum/common"

	"golang.org/x/sync/errgroup"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/flare-foundation/go-flare-common/pkg/policy"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
)

const (
	minRoundsStored uint32 = 10

	// Relay.sol's THRESHOLD_BIPS, the divisor of thresholdIncreaseBIPS
	relayThresholdBIPS = 10000

	// randomNumber ‖ proof, capped at depth 32 (2³² leaves) — the Relay sets no bound
	maxFinalizationDataLength = 33 * common.HashLength
)

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
	relayCutover     *shared.RelayCutover
}

// NewClient creates a new client that manages finalizations.
//
// messageChannel is used to receive messages from protocol.client.
func NewClient(ctx clientContext.ClientContext, messageChannel <-chan shared.ProtocolMessage, relayCutover *shared.RelayCutover) (*client, error) {
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
		return nil, fmt.Errorf("creating relay contract: %w", err)
	}
	finalizerContext, err := newFinalizerContext(cfg, relayContract)
	if err != nil {
		return nil, err
	}

	senderPkString, err := config.PrivateKeyFromConfig(cfg.Credentials.SigningPolicyPrivateKeyFile,
		cfg.Credentials.SigningPolicyPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("reading sender private key: %w", err)
	}
	txOpts, senderPk, err := credentials.CredentialsFromPrivateKey(senderPkString, chainCfg.ChainID)
	if err != nil {
		return nil, fmt.Errorf("creating sender register tx opts: %w", err)
	}
	relayClient, err := NewRelayContractClient(
		ethClient,
		cfg.ContractAddresses.Relay,
		senderPk,
		txOpts.From,
		&cfg.RelayGas,
		chainCfg.ChainID,
		relayCutover,
	)
	if err != nil {
		return nil, err
	}
	// only the new Relay demands the random number and Merkle proof, and they ride on the random
	// protocol's own message, so that protocol must be one the submitter queries
	if relayCutover.Scheduled() && !randomProtocolConfigured(cfg.Protocol, finalizerContext.randomNumberProtocolID) {
		return nil, fmt.Errorf("a relay cutover is scheduled but protocol %d, whose provider serves the random number and Merkle proof the new Relay needs to finalize it, is not configured", finalizerContext.randomNumberProtocolID)
	}

	submissionListener := NewSubmissionListener(cfg.ContractAddresses.Submission)
	finalizationStorage := newFinalizationStorage(relayCutover)

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
		relayCutover:         relayCutover,
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
	ctx context.Context, startTime time.Time,
) (time.Time, error) {
	// Read current signing policies from the database and add them to the storage
	spList, err := c.relayClient.FetchSigningPolicies(ctx, c.db, startTime.Unix(), time.Now().Unix())
	if err != nil {
		return startTime, err
	}
	for _, sp := range spList {
		newPolicy := policy.NewSigningPolicy(sp.policyData, nil)
		c.relayCutover.ObserveSigningPolicy(newPolicy.RewardEpochID, newPolicy.StartVotingRoundID)
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
	spListener := c.relayClient.SigningPolicyInitializedListener(ctx, c.db, startTime)
	for {
		var dbPolicy signingPolicyListenerResponse
		select {
		case dbPolicy = <-spListener:
		case <-ctx.Done():
			logger.Infof("Signing policy initialized listener stopped")
			return ctx.Err()
		}

		policy := policy.NewSigningPolicy(dbPolicy.policyData, nil)
		c.relayCutover.ObserveSigningPolicy(policy.RewardEpochID, policy.StartVotingRoundID)
		if policy.RewardEpochID < c.finalizerContext.startingRewardEpoch {
			continue
		}
		if err := c.signingPolicyStorage.Add(policy); err != nil {
			logger.Warnf("Error adding signing policy %v", err)
		}

		logger.Infof("New signing policy received for epoch %v", policy.RewardEpochID)

		c.signingPolicyStorage.RemoveBefore(c.finalizationStorage.LowestRoundStored()) // remove signingPolicies that will never be used again
	}
}

// signingPolicyData returns signing policy and voting threshold for the given votingRoundID.
//
// If the signing policy was expected to end before votingRoundID but was prolonged, the threshold
// is raised the way the Relay raises it: the policy's own threshold scaled by thresholdIncreaseBIPS.
// Both Relays scale identically (Relay.sol:1165-1180, RelayMainDeployed.sol:914-926), so the raise
// is not gated on the cutover.
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
	}
	// mirrors Relay.sol:1168-1180; clamp, not wrap — the contract's uint256 threshold is unreachable
	raised := uint64(sp.Threshold) * uint64(c.finalizerContext.thresholdIncreaseBIPS) / relayThresholdBIPS
	return sp, uint16(min(raised, math.MaxUint16))
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
				logger.Warnf("Ignoring message for voting round %d, protocolID  %d - before policy startVotingRoundID", protocolMessage.VotingRoundID, protocolMessage.ProtocolID)
				continue
			}
			logger.Panicf("messagesChannelListener: no signing policy found for voting round %d. Storage is empty: %v", protocolMessage.VotingRoundID, c.signingPolicyStorage.OldestStored() == nil) // this stops the whole fsp client, it only happens if there is no signing policy in the storage.
		}
		protocolMessage.FinalizationData = c.finalizationDataToStore(&protocolMessage, sp)

		finalizationReady, err := c.finalizationStorage.AddMessage(&protocolMessage, sp, threshold)

		if err != nil {
			logger.Debugf("Ignoring submitted message for protocol %d, round %d: %v", protocolMessage.ProtocolID, protocolMessage.VotingRoundID, err)
			continue
		}

		if finalizationReady.thresholdReached {
			logger.Infof("Threshold reached for protocol %d in voting round %d", finalizationReady.protocolID, finalizationReady.votingRoundID)
			c.onThresholdReached(&finalizationReady, sp)
		}
	}
}

// onThresholdReached queues the finalization and prunes rounds too old to finalize. Both intake
// paths call it: a buffered payload crosses the threshold inside AddMessage.
func (c *client) onThresholdReached(ready *FinalizationReady, sp *policy.SigningPolicy) {
	c.queueProcessor.Add(ready, sp.Seed)

	if ready.votingRoundID > minRoundsStored {
		c.finalizationStorage.RemoveRoundsBefore(ready.votingRoundID - minRoundsStored)
	}
}

// finalizationDataToStore returns what relay() needs appended: the random number and Merkle proof
// for the random protocol on the new Relay, nothing elsewhere. Their meaning is the Relay's to
// check; only the shape is — non-empty, word-aligned, capped — so a bad provider is flagged before a send.
func (c *client) finalizationDataToStore(m *shared.ProtocolMessage, sp *policy.SigningPolicy) []byte {
	expected := m.ProtocolID == c.finalizerContext.randomNumberProtocolID &&
		c.relayCutover.NewRelayFromRewardEpoch(sp.RewardEpochID)
	if !expected {
		if len(m.FinalizationData) > 0 {
			logger.Debugf("Ignoring finalization data for protocol %d in voting round %d, none is needed", m.ProtocolID, m.VotingRoundID)
		}
		return nil
	}

	if len(m.FinalizationData) == 0 || len(m.FinalizationData)%common.HashLength != 0 {
		logger.Errorf("Protocol %d served %d bytes of finalization data with its message for voting round %d, which cannot be finalized without the random number and whole Merkle proof nodes",
			m.ProtocolID, len(m.FinalizationData), m.VotingRoundID)
		return nil
	}
	if len(m.FinalizationData) > maxFinalizationDataLength {
		logger.Errorf("Protocol %d served %d bytes of finalization data with its message for voting round %d, more than the %d a random number and Merkle proof take",
			m.ProtocolID, len(m.FinalizationData), m.VotingRoundID, maxFinalizationDataLength)
		return nil
	}
	return m.FinalizationData
}
