package finalizer

import (
	clientContext "flare-tlc/client/context"
	"flare-tlc/config"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/system"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type finalizerClient struct {
	db *gorm.DB

	relayClient          *relayContractClient
	submissionClient     *submissionContractClient
	signingPolicyStorage *signingPolicyStorage
	submissionStorage    *submissionStorage
	queueProcessor       *finalizerQueueProcessor

	finalizerContext *finalizerContext
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

	systemManager, err := system.NewFlareSystemManager(cfg.ContractAddresses.SystemManager, ethClient)
	if err != nil {
		return nil, errors.Wrap(err, "error creating system manager contract")
	}
	finalizerContext, err := newFinalizerContext(cfg, systemManager)
	if err != nil {
		return nil, err
	}

	senderPkString, err := config.ReadFileToString(cfg.Credentials.SigningPolicyPrivateKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "error reading sender private key")
	}
	senderPk, err := chain.PrivateKeyFromHex(senderPkString)
	if err != nil {
		return nil, errors.Wrap(err, "error creating sender register tx opts")
	}
	relayClient, err := NewRelayContractClient(
		ethClient,
		cfg.ContractAddresses.Relay,
		senderPk,
	)
	if err != nil {
		return nil, err
	}
	submissionClient := NewSubmissionContractClient(cfg.ContractAddresses.Submission)
	submissionStorage := newSubmissionStorage()

	return &finalizerClient{
		db:                   ctx.DB(),
		relayClient:          relayClient,
		signingPolicyStorage: newSigningPolicyStorage(),
		submissionStorage:    submissionStorage,
		submissionClient:     submissionClient,
		queueProcessor:       newFinalizerQueueProcessor(submissionStorage, relayClient),
		finalizerContext:     finalizerContext,
	}, nil
}

func (c *finalizerClient) Run() error {
	startTime := time.Now().Add(-c.finalizerContext.startTimeOffset)
	err := c.startSigningPolicyInitializedListener(startTime)
	if err != nil {
		return err
	}
	go func() {
		c.submissionClient.SubmissionTxListener(c.db, startTime, c)
	}()
	go func() {
		c.queueProcessor.Run()
	}()
	return nil
}

func (c *finalizerClient) startSigningPolicyInitializedListener(startTime time.Time) error {
	// Read current signing policies from the database and add them to the storage
	spList, err := c.relayClient.FetchSigningPolicies(c.db, startTime.Unix(), time.Now().Unix())
	if err != nil {
		return err
	}
	for _, sp := range spList {
		policy := newSigningPolicy(sp.policyData)
		if policy.rewardEpochId < c.finalizerContext.startingRewardEpoch {
			continue
		}
		c.signingPolicyStorage.Add(policy)
	}
	logger.Info("Added %d signing policies", len(spList))

	go func() {
		if len(spList) > 0 {
			startTime = time.Unix(spList[len(spList)-1].timestamp, 0)
		}
		spListener := c.relayClient.SigningPolicyInitializedListener(c.db, startTime)
		for {
			dbPolicy := <-spListener
			policy := newSigningPolicy(dbPolicy.policyData)
			if policy.rewardEpochId < c.finalizerContext.startingRewardEpoch {
				continue
			}
			c.signingPolicyStorage.Add(policy)
			logger.Info("New signing policy received for epoch %v", policy.rewardEpochId)
			c.rewardEpochCleanup()
		}
	}()
	return nil
}

func (c *finalizerClient) ProcessSubmissionData(slr submissionListenerResponse) error {
	for _, payloadItem := range slr.payload {
		if payloadItem.votingRoundId < c.finalizerContext.startingVotingRound {
			continue
		}

		sp := c.signingPolicyStorage.GetForVotingRound(payloadItem.votingRoundId)
		if sp == nil {
			first := c.signingPolicyStorage.First()
			if first != nil && payloadItem.votingRoundId < first.startVotingRoundId {
				// This is a submission for an old voting round, skip it
				continue
			}
			return fmt.Errorf("no signing policy found for voting round %d", payloadItem.votingRoundId)
		}

		addResult, err := c.submissionStorage.Add(payloadItem.payload, sp)
		if err != nil {
			// Error is non-fatal, skip this submission
			logger.Warn("Error adding submission %v", err)
			continue
		}
		if addResult.thresholdReached {
			logger.Info("Threshold reached for voting round %d and hash %v", payloadItem.votingRoundId, payloadItem.payload.messageHash)
			c.queueProcessor.Add(payloadItem)
		}
	}
	return nil
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
