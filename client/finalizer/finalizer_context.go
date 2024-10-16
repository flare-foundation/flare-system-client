package finalizer

import (
	"flare-fsc/client/config"
	"flare-fsc/client/shared"
	"flare-fsc/utils"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
)

// Finalizer client settings
type finalizerContext struct {
	startingRewardEpoch int64
	startingVotingRound uint32
	startTimeOffset     time.Duration // offset for fetching reward epochs at the start of the client

	voterThresholdBIPS   uint16
	gracePeriodEndOffset time.Duration

	votingEpoch *utils.EpochConfig
	rewardEpoch *utils.RewardEpochConfig
}

// func newFinalizerContext(cfg *config.ClientConfig, systemsManager *system.FlareSystemsManager) (*finalizerContext, error) {
func newFinalizerContext(cfg *config.Client, relay *relay.Relay) (*finalizerContext, error) {
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
