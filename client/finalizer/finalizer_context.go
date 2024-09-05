package finalizer

import (
	"flare-fsc/client/config"
	"flare-fsc/client/shared"
	"flare-fsc/utils"
	"flare-fsc/utils/contracts/relay"
	"time"
)

// Finalizer client settings
type finalizerContext struct {
	startingRewardEpoch int64
	startingVotingRound uint32
	startTimeOffset     time.Duration // offset for fetching reward epochs at the start of the client

	voterThresholdBIPS   uint16
	gracePeriodEndOffset time.Duration

	votingEpoch *utils.Epoch
	rewardEpoch *utils.IntEpoch
}

// func newFinalizerContext(cfg *config.ClientConfig, systemsManager *system.FlareSystemsManager) (*finalizerContext, error) {
func newFinalizerContext(cfg *config.ClientConfig, relay *relay.Relay) (*finalizerContext, error) {
	votingEpoch, rewardEpoch, err := shared.EpochsFromChain(relay)
	if err != nil {
		return nil, err
	}
	startingVotingRound := cfg.Finalizer.StartingVotingRound
	if startingVotingRound == 0 {
		startingVotingRound = uint32(votingEpoch.EpochIndex(time.Now()))
	}
	return &finalizerContext{
		startingRewardEpoch:  cfg.Finalizer.StartingRewardEpoch,
		startingVotingRound:  startingVotingRound,
		startTimeOffset:      cfg.Finalizer.StartOffset,
		voterThresholdBIPS:   cfg.Finalizer.VoterThresholdBIPS,
		gracePeriodEndOffset: cfg.Finalizer.GracePeriodEndOffset,
		votingEpoch:          votingEpoch,
		rewardEpoch:          rewardEpoch,
	}, nil
}
