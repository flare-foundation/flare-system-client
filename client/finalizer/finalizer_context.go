package finalizer

import (
	"flare-tlc/client/config"
	"flare-tlc/client/shared"
	"flare-tlc/utils"
	"flare-tlc/utils/contracts/system"
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
}

func newFinalizerContext(cfg *config.ClientConfig, systemManager *system.FlareSystemManager) (*finalizerContext, error) {
	votingEpoch, err := shared.VotingEpochFromChain(systemManager)
	if err != nil {
		return nil, err
	}
	return &finalizerContext{
		startingRewardEpoch:  cfg.Finalizer.StartingRewardEpoch,
		startingVotingRound:  cfg.Finalizer.StartingVotingRound,
		startTimeOffset:      cfg.Finalizer.StartOffset,
		voterThresholdBIPS:   cfg.Finalizer.VoterThresholdBIPS,
		gracePeriodEndOffset: cfg.Finalizer.GracePeriodEndOffset,
		votingEpoch:          votingEpoch,
	}, nil
}
