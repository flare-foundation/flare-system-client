package clients

import (
	"flare-tlc/client/context"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/contracts/registry"
	"flare-tlc/utils/contracts/system"

	"gorm.io/gorm"
)

// Start tic voter registration & signing policy voter client 2 hours
// before end of epoch (reward epoch 3.5 days)
//  1. Listen until VotePowerBlockSelected (enabled voter registration) event is emitted
//  2. Call RegisterVoter function on VoterRegistry
//  3. Wait until we get voter registered event
//  4. Wait until SigningPolicyInitialized
//  5. Call signNewSigningPolicy
//  6. Wait until SigningPolicySigned is emitted (for the voter)

type registrationClient struct {
	epoch *utils.Epoch
	db    *gorm.DB

	flareSystemManager *system.FlareSystemManager
	voterRegistry      *registry.Registry

	voterRegistrationTopic0  string
	voterRegistrationAddress string

	mockableTime utils.TimeProvider
}

func NewRegistratinClient(ctx context.ClientContext) (*registrationClient, error) {
	config := ctx.Config()
	chainConfig := config.ChainConfig()
	cc, err := chainConfig.DialETH()
	if err != nil {
		return nil, err
	}

	fsmContract, err := system.NewFlareSystemManager(config.ContractAddresses.SystemManager, cc)
	if err != nil {
		return nil, err
	}

	regContract, err := registry.NewRegistry(config.ContractAddresses.VoterRegistry, cc)
	if err != nil {
		return nil, err
	}

	return &registrationClient{
		epoch:                    utils.NewEpoch(config.VoterRegistration.EpochStart, config.VoterRegistration.EpochPeriod), // Temp
		db:                       ctx.DB(),
		flareSystemManager:       fsmContract,
		voterRegistry:            regContract,
		voterRegistrationTopic0:  config.VoterRegistration.Topic0,
		voterRegistrationAddress: config.VoterRegistration.Address,
		mockableTime:             utils.RealTimeProvider{},
	}, nil
}

// Run runs the registration client, should be called in a goroutine
func (c *registrationClient) Run() {

	fsmf := c.flareSystemManager.FlareSystemManagerFilterer

	vpbsListenter := NewVotePowerBlockSelectedListener(
		c.db,
		&fsmf,
		c.epoch,
		c.voterRegistrationTopic0,
		c.voterRegistrationAddress,
	)
	for {
		// Wait until VotePowerBlockSelected (enabled voter registration) event is emitted
		powerBlockData := <-vpbsListenter.C
		logger.Info("VotePowerBlockSelected event emitted for epoch %d", powerBlockData.RewardEpoch)

		// Call RegisterVoter function on VoterRegistry

	}
}
