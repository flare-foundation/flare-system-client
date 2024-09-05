package finalizer

import (
	"context"
	"encoding/hex"
	clientContext "flare-fsc/client/context"
	"flare-fsc/config"
	"flare-fsc/database"
	"flare-fsc/logger"
	"flare-fsc/utils/contracts/relay"
	"flare-fsc/utils/credentials"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type finalizerClient struct {
	db finalizerDB

	relayClient          *relayContractClient
	submissionClient     *submissionContractClient
	signingPolicyStorage *signingPolicyStorage
	submissionStorage    *submissionStorage
	queueProcessor       *finalizerQueueProcessor

	finalizerContext *finalizerContext
}

type finalizerDB interface {
	FetchTransactionsByAddressAndSelector(
		common.Address, []byte, int64, int64,
	) ([]database.Transaction, error)
	FetchLogsByAddressAndTopic0(common.Address, string, int64, int64) ([]database.Log, error)
}

type finalizerDBImpl struct {
	client *gorm.DB
}

func (db finalizerDBImpl) FetchTransactionsByAddressAndSelector(
	address common.Address, selector []byte, from, to int64,
) ([]database.Transaction, error) {
	hexSelector := hex.EncodeToString(selector)
	return database.FetchTransactionsByAddressAndSelector(db.client, address.Hex(), hexSelector, from, to)
}

func (db finalizerDBImpl) FetchLogsByAddressAndTopic0(
	address common.Address, topic0 string, from, to int64,
) ([]database.Log, error) {
	return database.FetchLogsByAddressAndTopic0(db.client, address.Hex(), topic0, from, to)
}

func NewFinalizerClient(ctx clientContext.ClientContext) (*finalizerClient, error) {
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
	submissionStorage := newSubmissionStorage()

	db := finalizerDBImpl{client: ctx.DB()}

	return &finalizerClient{
		db:                   db,
		relayClient:          relayClient,
		signingPolicyStorage: newSigningPolicyStorage(),
		submissionStorage:    submissionStorage,
		submissionClient:     submissionClient,
		queueProcessor:       newFinalizerQueueProcessor(db, submissionStorage, relayClient, finalizerContext),
		finalizerContext:     finalizerContext,
	}, nil
}

func (c *finalizerClient) Run(ctx context.Context) error {
	return c.RunContext(ctx)
}

func (c *finalizerClient) RunContext(ctx context.Context) error {
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
		return c.submissionClient.SubmissionTxListener(ctx, c.db, startTime, c)
	})
	eg.Go(func() error {
		return c.queueProcessor.Run(ctx)
	})

	return eg.Wait()
}

func (c *finalizerClient) fetchExistingSigningPolicies(
	ctx context.Context, startTime time.Time,
) (time.Time, error) {
	// Read current signing policies from the database and add them to the storage
	spList, err := c.relayClient.FetchSigningPolicies(c.db, startTime.Unix(), time.Now().Unix())
	if err != nil {
		return startTime, err
	}
	for _, sp := range spList {
		policy := newSigningPolicy(sp.policyData)
		if policy.rewardEpochId < c.finalizerContext.startingRewardEpoch {
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
			break

		case <-ctx.Done():
			logger.Info("Signing policy initialized listener stopped")
			return ctx.Err()
		}

		policy := newSigningPolicy(dbPolicy.policyData)
		if policy.rewardEpochId < c.finalizerContext.startingRewardEpoch {
			continue
		}
		if err := c.signingPolicyStorage.Add(policy); err != nil {
			logger.Warn("Error adding signing policy %v", err)
		}
		logger.Info("New signing policy received for epoch %v", policy.rewardEpochId)
		c.rewardEpochCleanup()
	}
}

func (c *finalizerClient) ProcessSubmissionData(slr submissionListenerResponse) error {
	for _, payloadItem := range slr.payload {
		if payloadItem.votingRoundId < c.finalizerContext.startingVotingRound {
			logger.Debug("Ignoring submitted signature for voting round %d - before startingVotingRound", payloadItem.votingRoundId)
			continue
		}

		// Skip if voting round is in the future
		if !c.checkVotingRoundTime(payloadItem.votingRoundId) {
			continue
		}
		sp, threshold := c.signingPolicyData(payloadItem.votingRoundId)
		if sp == nil {
			first := c.signingPolicyStorage.First()
			if first != nil && payloadItem.votingRoundId < first.startVotingRoundId {
				// This is a submission for an old voting round, skip it
				logger.Debug("Ignoring submitted signature for voting round %d - before policy startVotingRoundId", payloadItem.votingRoundId)
				continue
			}
			return fmt.Errorf("no signing policy found for voting round %d", payloadItem.votingRoundId)
		}
		addResult, err := c.submissionStorage.Add(payloadItem.payload, sp, threshold)
		if err != nil {
			// Error is non-fatal, skip this submission
			logger.Debug("Ignoring submitted signature: %v", err)
			continue
		}
		if addResult.thresholdReached {
			logger.Info("Threshold reached for protocol %d in voting round %d with hash %v", payloadItem.protocolId, payloadItem.votingRoundId, payloadItem.payload.messageHash)
			c.queueProcessor.Add(payloadItem, sp.seed)
		}
	}
	return nil
}

// return signing policy and voting threshold for the given voting round
func (c *finalizerClient) signingPolicyData(votingRoundId uint32) (*signingPolicy, uint16) {
	sp, last := c.signingPolicyStorage.GetForVotingRound(votingRoundId)
	if sp == nil {
		return nil, 0
	}
	if !last {
		return sp, sp.threshold
	}
	endVotingEpoch := c.finalizerContext.rewardEpoch.EndEpoch(sp.rewardEpochId)
	end := c.finalizerContext.votingEpoch.EndTime(endVotingEpoch)

	if time.Now().Before(end) {
		return sp, sp.threshold
	} else {
		return sp, uint16((uint32(sp.voters.TotalWeight()) * 60) / 100)
	}
}

// Return true if voting round is not in the future, i.e., is <= the current voting round
func (c *finalizerClient) checkVotingRoundTime(votingRoundId uint32) bool {
	currentEpochId := c.finalizerContext.votingEpoch.EpochIndex(time.Now())
	return votingRoundId <= uint32(currentEpochId)
}

func (c *finalizerClient) rewardEpochCleanup() {
	cleanupTime := time.Now().Add(-2 * c.finalizerContext.startTimeOffset)
	cleanupVotingRoundId := c.finalizerContext.votingEpoch.EpochIndex(cleanupTime)
	if cleanupVotingRoundId < 0 {
		return
	}
	removedEpochIds := c.signingPolicyStorage.RemoveByVotingRound(uint32(cleanupVotingRoundId))
	c.submissionStorage.RemoveVotingRoundIds(removedEpochIds)
	if len(removedEpochIds) > 0 {
		logger.Info("Removed signing policies and submissions with reward epoch <= %d", removedEpochIds[len(removedEpochIds)-1])
	}
}
