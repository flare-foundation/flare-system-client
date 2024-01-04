package clients

import (
	"crypto/ecdsa"
	"flare-tlc/client/context"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/registry"
	"flare-tlc/utils/contracts/relay"
	"flare-tlc/utils/contracts/system"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
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
	epoch *utils.Epoch
	db    *gorm.DB

	flareSystemManager *system.FlareSystemManager
	voterRegistry      *registry.Registry
	relay              *relay.Relay

	txOpts     *bind.TransactOpts
	privateKey *ecdsa.PrivateKey
	txVerifier *chain.TxVerifier

	voterRegistrationAddress string
	voterRegistrationTopic0  string

	signingPolicyAddress string
	signingPolicyTopic0  string

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

	relayContract, err := relay.NewRelay(config.ContractAddresses.Relay, cc)
	if err != nil {
		return nil, err
	}

	pkString, err := chainConfig.GetPrivateKey()
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(pkString, "0x"))
	if err != nil {
		return nil, err
	}

	txOpts, err := chain.TransactOptsFromPrivateKey(pkString, chainConfig.ChainID)
	if err != nil {
		return nil, err
	}

	// regContract.RegisterVoter()

	return &registrationClient{
		epoch:                    utils.NewEpoch(config.VoterRegistration.EpochStart, config.VoterRegistration.EpochPeriod), // Temp
		db:                       ctx.DB(),
		flareSystemManager:       fsmContract,
		voterRegistry:            regContract,
		relay:                    relayContract,
		txOpts:                   txOpts,
		voterRegistrationTopic0:  config.VoterRegistration.Topic0,
		voterRegistrationAddress: config.VoterRegistration.Address,
		signingPolicyTopic0:      config.SigningPolicy.Topic0,
		signingPolicyAddress:     config.SigningPolicy.Address,
		mockableTime:             utils.RealTimeProvider{},
		privateKey:               privateKey,
		txVerifier:               chain.NewTxVerifier(cc),
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

	voterRegistrator := NewVoterRegistrator(
		c.voterRegistry,
		c.txOpts,
		c.txVerifier,
		c.privateKey,
	)

	signingPolicyHandler := NewSigningPolicyHandler(
		c.db,
		c.relay,
		c.flareSystemManager,
		c.txVerifier,
		c.txOpts,
		c.signingPolicyAddress,
		c.signingPolicyTopic0,
	)

	for {
		// Wait until VotePowerBlockSelected (enabled voter registration) event is emitted
		powerBlockData := <-vpbsListenter.C
		logger.Info("VotePowerBlockSelected event emitted for epoch %d", powerBlockData.RewardEpochId)

		// Call RegisterVoter function on VoterRegistry
		registered := <-voterRegistrator.RegisterVoter(powerBlockData.RewardEpochId, c.voterRegistrationAddress)
		if !registered {
			logger.Error("RegisterVoter failed")
			continue
		}

		// Wait until we get voter registered event
		// Already in RegisterVoter

		// Wait until SigningPolicyInitialized
		signingPolicy := <-signingPolicyHandler.signingPolicyInitializedListener(powerBlockData.Timestamp)

		// Call signNewSigningPolicy
		signed := <-signingPolicyHandler.SignNewSigningPolicy(signingPolicy.RewardEpochId, signingPolicy.SigningPolicyBytes)
		if !signed {
			logger.Error("SignNewSigningPolicy failed")
			continue
		}

	}
}
