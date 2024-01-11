package clients

import (
	"flare-tlc/client/context"
	"flare-tlc/config"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"

	"github.com/pkg/errors"
	"gorm.io/gorm"
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
	db *gorm.DB

	systemManagerClient *SystemManagerContractClient
	relayClient         *RelayContractClient
	registryClient      *RegistryContractClient

	identityAddress string
}

func NewRegistrationClient(ctx context.ClientContext) (*registrationClient, error) {
	cfg := ctx.Config()
	if !cfg.Voting.EnabledRegistration {
		return nil, nil
	}

	chainCfg := cfg.ChainConfig()
	ethClient, err := chainCfg.DialETH()
	if err != nil {
		return nil, err
	}

	senderPk, err := config.ReadFileToString(cfg.Credentials.SystemManagerSenderPrivateKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "error reading sender private key")
	}
	senderTxOpts, _, err := chain.CredentialsFromPrivateKey(senderPk, chainCfg.ChainID)
	if err != nil {
		return nil, errors.Wrap(err, "error creating sender register tx opts")
	}

	signerPkString, err := config.ReadFileToString(cfg.Credentials.SigningPolicyPrivateKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "error reading signer private key")
	}
	signerPk, err := chain.PrivateKeyFromHex(signerPkString)
	if err != nil {
		return nil, errors.Wrap(err, "error creating signer private key")
	}

	systemManagerClient, err := NewSystemManagerClient(
		ethClient,
		cfg.ContractAddresses.SystemManager,
		senderTxOpts,
		signerPk,
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
		cfg.ContractAddresses.VoterRegistry,
		senderTxOpts,
		signerPk,
	)
	if err != nil {
		return nil, err
	}

	return &registrationClient{
		db:                  ctx.DB(),
		systemManagerClient: systemManagerClient,
		relayClient:         relayClient,
		registryClient:      registryClient,
		identityAddress:     cfg.Credentials.IdentityAddress,
	}, nil
}

// Run runs the registration client, should be called in a goroutine
func (c *registrationClient) Run() error {
	epoch, err := c.systemManagerClient.EpochFromChain()
	if err != nil {
		return err
	}
	vpbsListener := c.systemManagerClient.VotePowerBlockSelectedListener(c.db, epoch)

	for {
		// Wait until VotePowerBlockSelected (enabled voter registration) event is emitted
		logger.Debug("Waiting for VotePowerBlockSelected event")
		powerBlockData := <-vpbsListener
		logger.Info("VotePowerBlockSelected event emitted for epoch %v", powerBlockData.RewardEpochId)

		id, _ := c.systemManagerClient.flareSystemManager.GetCurrentRewardEpochId(nil)
		logger.Debug("Current reward epoch id %v", id)
		logger.Debug("Reward epoch id %v", powerBlockData.RewardEpochId)

		// Call RegisterVoter function on VoterRegistry
		registerResult := <-c.registryClient.RegisterVoter(powerBlockData.RewardEpochId, c.identityAddress)
		if !registerResult.Success {
			logger.Error("RegisterVoter failed %s", registerResult.Message)
			continue
		}

		// Wait until we get voter registered event
		// Already in RegisterVoter

		// Wait until SigningPolicyInitialized event is emitted
		signingPolicy := <-c.relayClient.SigningPolicyInitializedListener(c.db, powerBlockData.Timestamp)
		logger.Info("SigningPolicyInitialized event emitted for epoch %v", signingPolicy.RewardEpochId)

		// Call signNewSigningPolicy
		signingResult := <-c.systemManagerClient.SignNewSigningPolicy(signingPolicy.RewardEpochId, signingPolicy.SigningPolicyBytes)
		if !signingResult.Success {
			logger.Error("SignNewSigningPolicy failed %s", signingResult.Message)
			continue
		}

	}
}
