package finalizer

import (
	clientContext "flare-tlc/client/context"
	"flare-tlc/config"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	startOffset = 7 * 24 * time.Hour // 7 days
)

type finalizerClient struct {
	db *gorm.DB

	relayClient          *relayContractClient
	submissionClient     *submissionContractClient
	signingPolicyStorage *signingPolicyStorage
	submissionStorage    *submissionStorage
	finalizerQueue       *finalizerQueue
}

func NewFinalizerClient(ctx clientContext.ClientContext) (*finalizerClient, error) {
	cfg := ctx.Config()
	if !cfg.Voting.EnabledFinalizer {
		return nil, nil
	}

	chainCfg := cfg.ChainConfig()
	ethClient, err := chainCfg.DialETH()
	if err != nil {
		return nil, err
	}

	senderPk, err := config.ReadFileToString(cfg.Credentials.SigningPolicyPrivateKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "error reading sender private key")
	}
	senderTxOpts, _, err := chain.CredentialsFromPrivateKey(senderPk, chainCfg.ChainID)
	if err != nil {
		return nil, errors.Wrap(err, "error creating sender register tx opts")
	}
	relayClient, err := NewRelayContractClient(
		ethClient,
		cfg.ContractAddresses.Relay,
		senderTxOpts,
	)
	if err != nil {
		return nil, err
	}
	submissionClient := NewSubmissionContractClient(cfg.ContractAddresses.Submission)

	return &finalizerClient{
		db:                   ctx.DB(),
		relayClient:          relayClient,
		signingPolicyStorage: newSigningPolicyStorage(),
		submissionClient:     submissionClient,
		submissionStorage:    newSubmissionStorage(),
		finalizerQueue:       newFinalizerQueue(),
	}, nil
}

func (c *finalizerClient) Run() error {
	startTime := time.Now().Add(-startOffset)
	spListener := c.relayClient.SigningPolicyInitializedListener(c.db, startTime)
	go func() {
		for {
			dbPolicy := <-spListener
			// Todo: move to listener to avoid creating a channel and a goroutine
			// Todo: synchronize with the epoch
			c.signingPolicyStorage.Add(newSigningPolicy(dbPolicy.policyData))
		}
	}()

	txListener := c.submissionClient.SubmissionTxListener(c.db, startTime)
	go func() {
		for {
			submResponse := <-txListener

			for _, payloadItem := range submResponse.payload {
				sp := c.signingPolicyStorage.GetForVotingRound(payloadItem.votingRoundId)
				if sp == nil {
					logger.Error("No signing policy found for voting round %d", payloadItem.votingRoundId)
					continue
				}
				addResult, err := c.submissionStorage.Add(payloadItem.payload, sp)
				if err != nil {
					logger.Error("Error adding submission %v", err)
					continue
				}
				if addResult.thresholdReached {
					logger.Info("Threshold reached for voting round %d", payloadItem.votingRoundId)
					c.finalizerQueue.Add(payloadItem.votingRoundId, payloadItem.protocolId, payloadItem.payload.messageHash)
				}
			}
		}
	}()
	return nil
}
