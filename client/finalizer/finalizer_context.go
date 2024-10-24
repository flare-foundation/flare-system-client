package finalizer

import (
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
)

// Finalizer client settings
type finalizerContext struct {
	startingRewardEpoch int64
	startingVotingRound uint32
	startTimeOffset     time.Duration // offset for fetching reward epochs at the start of the client

	voterThresholdBIPS   uint16
	gracePeriodEndOffset time.Duration

	votingRoundTiming *utils.EpochTimingConfig
	rewardEpoch       *utils.RewardEpochConfig
}

// func newFinalizerContext(cfg *config.ClientConfig, systemsManager *system.FlareSystemsManager) (*finalizerContext, error) {
func newFinalizerContext(cfg *config.Client, relay *relay.Relay) (*finalizerContext, error) {
	votingRoundTiming, rewardEpoch, err := shared.EpochsFromChain(relay)
	if err != nil {
		return nil, err
	}
	startingVotingRound := cfg.Finalizer.StartingVotingRound
	if startingVotingRound == 0 {
		startingVotingRound = uint32(votingRoundTiming.EpochIndex(time.Now()))
	}
	return &finalizerContext{
		startingRewardEpoch:  cfg.Finalizer.StartingRewardEpoch,
		startingVotingRound:  startingVotingRound,
		startTimeOffset:      cfg.Finalizer.StartOffset,
		voterThresholdBIPS:   cfg.Finalizer.VoterThresholdBIPS,
		gracePeriodEndOffset: cfg.Finalizer.GracePeriodEndOffset,
		votingRoundTiming:    votingRoundTiming,
		rewardEpoch:          rewardEpoch,
	}, nil
}
