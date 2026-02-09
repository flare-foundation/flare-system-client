package epoch

import (
	"context"
	"math/big"

	clientConfig "github.com/flare-foundation/flare-system-client/client/config"
	flarectx "github.com/flare-foundation/flare-system-client/client/context"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/config"
	"github.com/flare-foundation/flare-system-client/utils/credentials"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/flare-foundation/go-flare-common/pkg/logger"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/system"
)

// client performs reward epoch registration and signing actions, triggered on SystemsManager contract events:
//   - Voter registration (on VoterPowerBlockSelected)
//   - Signing new signing policy (on SigningPolicyInitialized)
//   - Signing uptime vote (on SignUptimeVoteEnabled)
//   - Signing rewards (on UptimeVoteSigned with threshold reached)
type client struct {
	db epochClientDB

	systemsManagerClient systemsManagerContractClient
	relayClient          relayContractClient
	registryClient       registryContractClient

	identityAddress common.Address

	preregistrationEnabled bool
	registrationEnabled    bool
	uptimeVotingEnabled    bool
	rewardsSigningEnabled  bool

	rewardsConfig *clientConfig.RewardsConfig
}

// NewClient creates a client that manages reward epoch tasks.
func NewClient(ctx flarectx.ClientContext) (*client, error) {
	cfg := ctx.Config()
	if !cfg.Clients.EpochClientEnabled() {
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

	signerPk, err := config.PrivateKeyFromConfig(
		cfg.Credentials.SigningPolicyPrivateKeyFile,
		cfg.Credentials.SigningPolicyPrivateKey,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error creating signer private key")
	}

	systemsManagerClient, err := NewSystemsManagerClient(
		ethClient,
		&cfg.SystemManagerGas,
		cfg.ContractAddresses.SystemsManager,
		senderTxOpts,
		signerPk,
		chainCfg.ChainID,
	)
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
		cfg.ContractAddresses.VoterPreRegistry,
		senderTxOpts,
		signerPk,
	)
	if err != nil {
		return nil, err
	}

	identityAddress := cfg.Identity.Address
	if identityAddress == (common.Address{}) {
		return nil, errors.New("no identity address provided")
	}
	logger.Debugf("Identity address %v", identityAddress)

	db := epochClientDBGorm{db: ctx.DB()}
	return &client{
		db:                     db,
		systemsManagerClient:   systemsManagerClient,
		relayClient:            relayClient,
		registryClient:         registryClient,
		identityAddress:        identityAddress,
		preregistrationEnabled: cfg.Clients.EnabledPreregistration,
		registrationEnabled:    cfg.Clients.EnabledRegistration,
		uptimeVotingEnabled:    cfg.Clients.EnabledUptimeVoting,
		rewardsSigningEnabled:  cfg.Clients.EnabledRewardSigning,
		rewardsConfig:          &cfg.Rewards,
	}, nil
}

// Run runs the client. Should be called in a goroutine.
func (c *client) Run(ctx context.Context) error {
	logger.Info("Starting reward epoch client")

	rewardEpochTiming, err := c.systemsManagerClient.RewardEpochTimingFromChain()
	if err != nil {
		return err
	}

	var epochStartedListener <-chan *system.FlareSystemsManagerRewardEpochStarted
	var vpbsListener <-chan *system.FlareSystemsManagerVotePowerBlockSelected
	var policyListener <-chan *relay.RelaySigningPolicyInitialized
	var uptimeEnabledListener <-chan *system.FlareSystemsManagerSignUptimeVoteEnabled
	var uptimeSignedListener <-chan *system.FlareSystemsManagerUptimeVoteSigned

	if c.preregistrationEnabled {
		epochStartedListener = c.systemsManagerClient.RewardEpochStartedListener(c.db, rewardEpochTiming)
	}
	if c.registrationEnabled {
		logger.Info("Waiting for VotePowerBlockSelected event to start registration")
		vpbsListener = c.systemsManagerClient.VotePowerBlockSelectedListener(c.db, rewardEpochTiming)
		policyListener = c.relayClient.SigningPolicyInitializedListener(c.db, rewardEpochTiming)
	}
	if c.uptimeVotingEnabled {
		logger.Info("Waiting for SignUptimeVoteEnabled event to start uptime vote signing")
		uptimeEnabledListener = c.systemsManagerClient.SignUptimeVoteEnabledListener(c.db, rewardEpochTiming)
	}
	if c.rewardsSigningEnabled {
		logger.Info("Waiting for UptimeVoteSigned event to start rewards signing")
		uptimeSignedListener = c.systemsManagerClient.UptimeVoteSignedListener(c.db, rewardEpochTiming)
	}

	for {
		select {
		case rewardEpochStarted := <-epochStartedListener:
			logger.Debugf("RewardEpochStarted event emitted for epoch %v", rewardEpochStarted.RewardEpochId)
			c.preregisterVoter(new(big.Int).Add(rewardEpochStarted.RewardEpochId, big.NewInt(1)))
		case powerBlockData := <-vpbsListener:
			logger.Debugf("VotePowerBlockSelected event emitted for epoch %v", powerBlockData.RewardEpochId)
			c.registerVoter(powerBlockData.RewardEpochId)
		case signingPolicy := <-policyListener:
			logger.Debugf("SigningPolicyInitialized event emitted for epoch %v", signingPolicy.RewardEpochId)
			c.signPolicy(signingPolicy.RewardEpochId, signingPolicy.SigningPolicyBytes)
		case uptimeVoteEnabled := <-uptimeEnabledListener:
			logger.Debugf("SignUptimeVoteEnabled event emitted for epoch %v", uptimeVoteEnabled.RewardEpochId)
			c.signUptimeVote(uptimeVoteEnabled.RewardEpochId)
		case uptimeVoteSigned := <-uptimeSignedListener:
			logger.Infof("Uptime vote threshold reached for epoch %v, signing rewards", uptimeVoteSigned.RewardEpochId)
			c.signRewards(uptimeVoteSigned.RewardEpochId)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (c *client) registerVoter(epochID *big.Int) {
	if !c.isFutureEpoch(epochID) {
		logger.Debugf("Skipping registration process for old epoch %v", epochID)
		return
	}

	logger.Infof("VotePowerBlockSelected event emitted for next epoch %v, starting registration", epochID)
	registerResult := <-c.registryClient.RegisterVoter(epochID, c.identityAddress)
	if !registerResult.Success {
		logger.Errorf("RegisterVoter failed %s", registerResult.Message)
	}
}

func (c *client) preregisterVoter(epochID *big.Int) {
	if !c.isFutureEpoch(epochID) {
		logger.Debugf("Skipping pre-registration process for old epoch %v", epochID)
		return
	}

	registerResult := <-c.registryClient.PreregisterVoter(epochID, c.identityAddress)
	if !registerResult.Success {
		logger.Errorf("PreregisterVoter failed %s", registerResult.Message)
	}
}

func (c *client) signPolicy(epochID *big.Int, policy []byte) {
	if !c.isFutureEpoch(epochID) {
		logger.Debugf("Skipping policy signing for old epoch %v", epochID)
		return
	}

	logger.Infof("SigningPolicyInitialized event emitted for next epoch %v, signing new policy", epochID)
	signingResult := <-c.systemsManagerClient.SignNewSigningPolicy(epochID, policy)
	if signingResult.Success {
		logger.Info("SignNewSigningPolicy success")
	} else {
		logger.Errorf("SignNewSigningPolicy failed %s", signingResult.Message)
		return
	}
}

func (c *client) signUptimeVote(epochId *big.Int) {
	signUptimeVoteResult := <-c.systemsManagerClient.SignUptimeVote(epochId)
	if signUptimeVoteResult.Success {
		logger.Info("SignUptimeVote completed")
	} else {
		logger.Errorf("SignUptimeVote failed %s", signUptimeVoteResult.Message)
		return
	}
}

func (c *client) isFutureEpoch(epochID *big.Int) bool {
	epochIDResult := <-c.systemsManagerClient.GetCurrentRewardEpochID()
	if !epochIDResult.Success {
		logger.Errorf("GetCurrentRewardEpochId failed %s", epochIDResult.Message)
		return false
	}
	currentEpochId := epochIDResult.Value
	if epochID.Cmp(currentEpochId) <= 0 {
		logger.Debugf("Epoch in the past: current %v >= next %v", currentEpochId, epochID)
		return false
	}
	return true
}

// signRewards signs the reward claim Merkle root for the given epoch.
//
// Once uptime signing is completed, data providers can start signing the reward hash for the epoch.
// The end goal is that every data provider calculates reward hash by themselves and signs it. While rewarding logic
// is still in flux, there is an interim solution where reward data is calculated & published centrally, and data
// providers independently verify and sign the hash.
//
// Here signRewards first retrieves the reward claim data file from the configured URL. It verifies the provided
// Merkle root matches the list of claims, checks there is a reward claim for the identity address of this provider,
// with reward amount within expected bounds.
//
// If all checks pass, it signs the Merkle root and sends the signature to the SystemsManager contract.
//
// Since reward claim data is currently published manually, and it might take a day or so for the data to be available,
// a retry mechanism is employed with a large retry interval (configurable).
func (c *client) signRewards(epochId *big.Int) {
	res := shared.ExecuteWithRetryAttempts(func(i int) (*struct{}, error) {
		if c.systemsManagerClient.IsRewardHashSigned(epochId) {
			return nil, nil
		}

		logger.Infof("Signing rewards for epoch %v, attempt %d", epochId, i)

		data, err := fetchRewardData(epochId, c.rewardsConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to fetch reward data for epoch %d", epochId)
		}
		if data == nil {
			return nil, errors.New("no reward data found")
		}
		hash, weightClaims, err := verifyRewardData(epochId, c.identityAddress, data, c.rewardsConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "reward data verification for epoch %d failed", epochId)
		}
		signingResult := <-c.systemsManagerClient.SignRewards(epochId, hash, weightClaims)
		if !signingResult.Success {
			return nil, errors.Errorf("unable to send reward signature")
		}
		return nil, nil
	}, c.rewardsConfig.Retries, c.rewardsConfig.RetryInterval)

	// The retry loop may run four hours until the reward data is published, so we don't block for result here.
	go func() {
		status := <-res
		if status.Success {
			logger.Infof("Signing rewards for epoch %v completed", epochId)
		} else {
			logger.Infof("Signing rewards for epoch %v failed: %s", epochId, status.Message)
		}
	}()
}
