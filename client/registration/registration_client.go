package registration

import (
	"context"
	clientConfig "flare-tlc/client/config"
	flarectx "flare-tlc/client/context"
	"flare-tlc/config"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/system"
	"flare-tlc/utils/credentials"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"math/big"
)

// Start tic voter registration & signing policy voter client 2 hours
// before end of epoch (reward epoch 3.5 days)
//  1. Listen until VotePowerBlockSelected (enabled voter registration) event is emitted
//  2. Call RegisterVoter function on VoterRegistry
//  3. Wait until we get voter registered event
//  4. Wait until SigningPolicyInitialized is emitted
//  5. Call signNewSigningPolicy
//  6. Wait until SigningPolicySigned is emitted (for the voter)

type registrationClient struct {
	db registrationClientDB

	systemsManagerClient systemsManagerContractClient
	relayClient          relayContractClient
	registryClient       registryContractClient

	identityAddress common.Address
	rewardsConfig   *clientConfig.RewardsConfig
}

type registrationClientDB interface {
	FetchLogsByAddressAndTopic0(common.Address, string, int64, int64) ([]database.Log, error)
}

type registrationClientDBGorm struct {
	db *gorm.DB
}

func (g registrationClientDBGorm) FetchLogsByAddressAndTopic0(
	address common.Address, topic0 string, fromBlock int64, toBlock int64,
) ([]database.Log, error) {
	return database.FetchLogsByAddressAndTopic0(g.db, address.Hex(), topic0, fromBlock, toBlock)
}

func NewRegistrationClient(ctx flarectx.ClientContext) (*registrationClient, error) {
	cfg := ctx.Config()
	if !cfg.Clients.EnabledRegistration {
		return nil, nil
	}

	chainCfg := cfg.ChainConfig()
	ethClient, err := chainCfg.DialETH()
	if err != nil {
		return nil, err
	}

	senderPk, err := config.PrivateKeyFromConfig(cfg.Credentials.SystemClientSenderPrivateKeyFile,
		cfg.Credentials.SystemClientSenderPrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "error reading sender private key")
	}
	senderTxOpts, _, err := credentials.CredentialsFromPrivateKey(senderPk, chainCfg.ChainID)
	if err != nil {
		return nil, errors.Wrap(err, "error creating sender register tx opts")
	}

	signerPk, err := config.PrivateKeyFromConfig(cfg.Credentials.SigningPolicyPrivateKeyFile,
		cfg.Credentials.SigningPolicyPrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "error creating signer private key")
	}

	systemsManagerClient, err := NewSystemsManagerClient(ethClient, cfg.ContractAddresses.SystemsManager, senderTxOpts, signerPk, chainCfg.ChainID)
	if err != nil {
		return nil, err
	}

	relayClient, err := NewRelayContractClient(
		ethClient,
		cfg.ContractAddresses.Relay,
	)
	if err != nil {
		return nil, err
	}

	registryClient, err := NewRegistryContractClient(
		ethClient,
		&cfg.RegisterGas,
		cfg.ContractAddresses.VoterRegistry,
		senderTxOpts,
		signerPk,
	)
	if err != nil {
		return nil, err
	}

	identityAddress := cfg.Identity.Address
	if identityAddress == chain.EmptyAddress {
		return nil, errors.New("no identity address provided")
	}
	logger.Debug("Identity addr %v", identityAddress)

	db := registrationClientDBGorm{db: ctx.DB()}
	return &registrationClient{
		db:                   db,
		systemsManagerClient: systemsManagerClient,
		relayClient:          relayClient,
		registryClient:       registryClient,
		identityAddress:      identityAddress,
		rewardsConfig:        &cfg.Rewards,
	}, nil
}

// Run runs the registration client, should be called in a goroutine
func (c *registrationClient) Run(ctx context.Context) error {
	return c.RunContext(ctx)
}

func (c *registrationClient) RunContext(ctx context.Context) error {
	epoch, err := c.systemsManagerClient.RewardEpochFromChain()
	if err != nil {
		return err
	}
	vpbsListener := c.systemsManagerClient.VotePowerBlockSelectedListener(c.db, epoch)
	policyListener := c.relayClient.SigningPolicyInitializedListener(c.db, epoch)
	uptimeEnabledListener := c.systemsManagerClient.SignUptimeVoteEnabledListener(c.db, epoch)

	var uptimeSignedListener <-chan *system.FlareSystemsManagerUptimeVoteSigned
	if c.rewardsConfig.SigningEnabled {
		uptimeSignedListener = c.systemsManagerClient.UptimeVoteSignedListener(c.db, epoch)
	} else {
		uptimeSignedListener = make(chan *system.FlareSystemsManagerUptimeVoteSigned)
	}
	rewardsSigned := map[int64]bool{}

	// Wait until VotePowerBlockSelected (enabled voter registration) event is emitted
	logger.Info("Waiting for VotePowerBlockSelected event to start registration")

	for {
		select {
		case powerBlockData := <-vpbsListener:
			logger.Debug("VotePowerBlockSelected event emitted for epoch %v", powerBlockData.RewardEpochId)
			c.registerVoter(powerBlockData.RewardEpochId)
		case signingPolicy := <-policyListener:
			logger.Debug("SigningPolicyInitialized event emitted for epoch %v", signingPolicy.RewardEpochId)
			c.signPolicy(signingPolicy.RewardEpochId, signingPolicy.SigningPolicyBytes)
		case uptimeVoteEnabled := <-uptimeEnabledListener:
			logger.Debug("SignUptimeVoteEnabled event emitted for epoch %v", uptimeVoteEnabled.RewardEpochId)
			c.signUptimeVote(uptimeVoteEnabled.RewardEpochId)
		case uptimeVoteSigned := <-uptimeSignedListener:
			logger.Debug("UptimeVoteSigned event emitted for epoch %v, signer %s", uptimeVoteSigned.RewardEpochId, uptimeVoteSigned.SigningPolicyAddress.Hex())
			if uptimeVoteSigned.ThresholdReached {
				epochIdInt := uptimeVoteSigned.RewardEpochId.Int64()
				// There might be multiple events fired for the same epoch, only attempt to sign once
				if rewardsSigned[epochIdInt] {
					logger.Debug("Already attempted reward signing for epoch %d", epochIdInt)
					continue
				}

				logger.Info("Uptime vote threshold reached for epoch %v, signing rewards", epochIdInt)
				rewardsSigned[epochIdInt] = true
				c.signRewards(uptimeVoteSigned.RewardEpochId)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (c *registrationClient) registerVoter(epochId *big.Int) {
	if !c.isFutureEpoch(epochId) {
		logger.Debug("Skipping registration process for old epoch %v", epochId)
		return
	}

	logger.Info("VotePowerBlockSelected event emitted for next epoch %v, starting registration", epochId)
	registerResult := <-c.registryClient.RegisterVoter(epochId, c.identityAddress)
	if registerResult.Success {
		logger.Info("RegisterVoter success")
	} else {
		logger.Error("RegisterVoter failed %s", registerResult.Message)
	}
}

func (c *registrationClient) signPolicy(epochId *big.Int, policy []byte) {
	if !c.isFutureEpoch(epochId) {
		logger.Debug("Skipping policy signing for old epoch %v", epochId)
		return
	}

	logger.Info("SigningPolicyInitialized event emitted for next epoch %v, signing new policy", epochId)
	signingResult := <-c.systemsManagerClient.SignNewSigningPolicy(epochId, policy)
	if signingResult.Success {
		logger.Info("SignNewSigningPolicy success")
	} else {
		logger.Error("SignNewSigningPolicy failed %s", signingResult.Message)
		return
	}
}

func (c *registrationClient) signUptimeVote(epochId *big.Int) {
	logger.Info("SignUptimeVoteEnabled event emitted for epoch %v, signing uptime vote", epochId)
	signUptimeVoteResult := <-c.systemsManagerClient.SignUptimeVote(epochId)
	if signUptimeVoteResult.Success {
		logger.Info("SignUptimeVote completed")
	} else {
		logger.Error("SignUptimeVote failed %s", signUptimeVoteResult.Message)
		return
	}
}

func (c *registrationClient) isFutureEpoch(epochId *big.Int) bool {
	epochIdResult := <-c.systemsManagerClient.GetCurrentRewardEpochId()
	if !epochIdResult.Success {
		logger.Error("GetCurrentRewardEpochId failed %s", epochIdResult.Message)
		return false
	}
	currentEpochId := epochIdResult.Value
	if epochId.Cmp(currentEpochId) <= 0 {
		logger.Debug("Epoch in the past: current %v >= next %v", currentEpochId, epochId)
		return false
	}
	return true
}

func (c *registrationClient) signRewards(epochId *big.Int) {
	logger.Info("Signing rewards for epoch %v", epochId)
	hash, weightClaims, err := getRewardsHash(epochId, c.rewardsConfig)
	if err != nil {
		logger.Error("error obtaining reward hash data for epoch %v, restart client to retry: %s", epochId, err)
		return
	}
	signingResult := <-c.systemsManagerClient.SignRewards(epochId, hash, weightClaims)
	if signingResult.Success {
		logger.Info("SignRewards completed")
	} else {
		logger.Error("SignRewards failed %s", signingResult.Message)
	}
}
