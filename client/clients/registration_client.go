package clients

import (
	"flare-tlc/client/context"
	"flare-tlc/logger"

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

	signingPolicyTopic0      string
	voterRegistrationAddress string
	votePowerBlockTopic0     string
}

func NewRegistratinClient(ctx context.ClientContext) (*registrationClient, error) {
	config := ctx.Config()
	chainConfig := config.ChainConfig()
	ethClient, err := chainConfig.DialETH()
	if err != nil {
		return nil, err
	}

	pkString, err := chainConfig.GetPrivateKey()
	if err != nil {
		return nil, err
	}

	systemManagerClient, err := NewSystemManagerClient(
		chainConfig.ChainID,
		ethClient,
		config.ContractAddresses.SystemManager,
		pkString,
	)
	if err != nil {
		return nil, err
	}

	relayClient, err := NewRelayContractClient(
		chainConfig.ChainID,
		ethClient,
		config.ContractAddresses.Relay,
		pkString,
	)
	if err != nil {
		return nil, err
	}

	registryClient, err := NewRegistryContractClient(
		chainConfig.ChainID,
		ethClient,
		config.ContractAddresses.VoterRegistry,
		pkString,
	)
	if err != nil {
		return nil, err
	}

	return &registrationClient{
		db:                  ctx.DB(),
		systemManagerClient: systemManagerClient,
		relayClient:         relayClient,
		registryClient:      registryClient,

		signingPolicyTopic0:      config.SigningPolicy.Topic0,
		voterRegistrationAddress: config.VoterRegistration.Address,
		votePowerBlockTopic0:     config.VotePowerBlock.Topic0,
	}, nil
}

// Run runs the registration client, should be called in a goroutine
func (c *registrationClient) Run() error {

	epoch, err := c.systemManagerClient.EpochFromChain()
	if err != nil {
		return err
	}

	for {
		// Wait until VotePowerBlockSelected (enabled voter registration) event is emitted
		powerBlockData := <-c.systemManagerClient.VotePowerBlockSelectedListener(c.db, epoch, c.votePowerBlockTopic0)
		logger.Info("VotePowerBlockSelected event emitted for epoch %v", powerBlockData.RewardEpochId)

		// Call RegisterVoter function on VoterRegistry
		registerResult := <-c.registryClient.RegisterVoter(powerBlockData.RewardEpochId, c.voterRegistrationAddress)
		if !registerResult.Success {
			logger.Error("RegisterVoter failed %s", registerResult.Message)
			continue
		}

		// Wait until we get voter registered event
		// Already in RegisterVoter

		// Wait until SigningPolicyInitialized event is emitted
		signingPolicy := <-c.relayClient.SigningPolicyInitializedListener(c.db, powerBlockData.Timestamp, c.signingPolicyTopic0)

		// Call signNewSigningPolicy
		signingResult := <-c.systemManagerClient.SignNewSigningPolicy(signingPolicy.RewardEpochId, signingPolicy.SigningPolicyBytes)
		if !signingResult.Success {
			logger.Error("SignNewSigningPolicy failed %s", signingResult.Message)
			continue
		}

	}
}
