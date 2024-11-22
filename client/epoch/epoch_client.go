package epoch

import (
	"context"
	clientConfig "flare-tlc/client/config"
	flarectx "flare-tlc/client/context"
	"flare-tlc/client/shared"
	"flare-tlc/config"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/relay"
	"flare-tlc/utils/contracts/system"
	"flare-tlc/utils/credentials"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

// EpochClient performs reward epoch registration and signing actions, triggered on SystemsManager contract events:
// - Voter registration (on VoterPowerBlockSelected)
// - Signing new signing policy (on SigningPolicyInitialized)
// - Signing uptime vote (on SignUptimeVoteEnabled)
// - Signing rewards (on UptimeVoteSigned with threshold reached)
type EpochClient struct {
	db epochClientDB

	systemsManagerClient systemsManagerContractClient
	relayClient          relayContractClient
	registryClient       registryContractClient

	identityAddress common.Address

	registrationEnabled   bool
	uptimeVotingEnabled   bool
	rewardsSigningEnabled bool

	rewardsConfig *clientConfig.RewardsConfig
}

func NewEpochClient(ctx flarectx.ClientContext) (*EpochClient, error) {
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

	db := epochClientDBGorm{db: ctx.DB()}
	return &EpochClient{
		db:                    db,
		systemsManagerClient:  systemsManagerClient,
		relayClient:           relayClient,
		registryClient:        registryClient,
		identityAddress:       identityAddress,
		registrationEnabled:   cfg.Clients.EnabledRegistration,
		uptimeVotingEnabled:   cfg.Clients.EnabledUptimeVoting,
		rewardsSigningEnabled: cfg.Clients.EnabledRewardSigning,
		rewardsConfig:         &cfg.Rewards,
	}, nil
}

// Run runs the  client, should be called in a goroutine
func (c *EpochClient) Run(ctx context.Context) error {
	return c.RunContext(ctx)
}

func (c *EpochClient) RunContext(ctx context.Context) error {
	logger.Info("Starting reward epoch client")

	epoch, err := c.systemsManagerClient.RewardEpochFromChain()
	if err != nil {
		return err
	}

	var vpbsListener <-chan *system.FlareSystemsManagerVotePowerBlockSelected
	var policyListener <-chan *relay.RelaySigningPolicyInitialized
	var uptimeEnabledListener <-chan *system.FlareSystemsManagerSignUptimeVoteEnabled
	var uptimeSignedListener <-chan *system.FlareSystemsManagerUptimeVoteSigned

	if c.registrationEnabled {
		logger.Info("Waiting for VotePowerBlockSelected event to start registration")
		vpbsListener = c.systemsManagerClient.VotePowerBlockSelectedListener(c.db, epoch)
		policyListener = c.relayClient.SigningPolicyInitializedListener(c.db, epoch)
	}
	if c.uptimeVotingEnabled {
		logger.Info("Waiting for SignUptimeVoteEnabled event to start uptime vote signing")
		uptimeEnabledListener = c.systemsManagerClient.SignUptimeVoteEnabledListener(c.db, epoch)
	}
	if c.rewardsSigningEnabled {
		logger.Info("Waiting for UptimeVoteSigned event to start rewards signing")
		uptimeSignedListener = c.systemsManagerClient.UptimeVoteSignedListener(c.db, epoch)
	}

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
			logger.Info("Uptime vote threshold reached for epoch %v, signing rewards", uptimeVoteSigned.RewardEpochId)
			c.signRewards(uptimeVoteSigned.RewardEpochId)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (c *EpochClient) registerVoter(epochId *big.Int) {
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

func (c *EpochClient) signPolicy(epochId *big.Int, policy []byte) {
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

func (c *EpochClient) signUptimeVote(epochId *big.Int) {
	signUptimeVoteResult := <-c.systemsManagerClient.SignUptimeVote(epochId)
	if signUptimeVoteResult.Success {
		logger.Info("SignUptimeVote completed")
	} else {
		logger.Error("SignUptimeVote failed %s", signUptimeVoteResult.Message)
		return
	}
}

func (c *EpochClient) isFutureEpoch(epochId *big.Int) bool {
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

// signRewards signs the reward claim merkle root for the given epoch.
//
// Once uptime signing is completed, data providers can start signing the reward hash for the epoch.
// The end goal is that every data provider calculates reward hash by themselves and signs it. While rewarding logic
// is still in flux, there is an interim solution where reward data is calculated & published centrally, and data
// providers independently verify and sign the hash.
//
// Here signRewards first retrieves the reward claim data file from the configured URL. It verifies the provided
// merkle root matches the list of claims, checks there is a reward claim for the identity address of this provider,
// with reward amount within expected bounds.
//
// If all checks pass, it signs the merkle root and sends the signature to the SystemsManager contract.
//
// Since reward claim data is currently published manually, and it might take a day or so for the data to be available,
// a retry mechanism is employed with a large retry interval (configurable).
func (c *EpochClient) signRewards(epochId *big.Int) {
	res := shared.ExecuteWithRetryAttempts(func(i int) (*struct{}, error) {
		if c.systemsManagerClient.IsRewardHashSigned(epochId) {
			return nil, nil
		}

		logger.Info("Signing rewards for epoch %v, attempt %d", epochId, i)

		data, err := fetchRewardData(epochId, c.rewardsConfig)
		if data == nil {
			return nil, errors.New("no reward data found")
		}
		if err != nil {
			return nil, errors.Wrapf(err, "unable to fetch reward data for epoch %d", epochId)
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
			logger.Info("Signing rewards for epoch %v completed", epochId)
		} else {
			logger.Info("Signing rewards for epoch %v failed: %s", epochId, status.Message)
		}
	}()
}
