package finalizer

import "flare-tlc/client/config"

// Finalizer client settings
type finalizerContext struct {
	startingRewardEpoch int64
	startingVotingRound uint32
}

func newFinalizerContext(cfg *config.ClientConfig) *finalizerContext {
	return &finalizerContext{
		startingRewardEpoch: cfg.Finalizer.StartingRewardEpoch,
		startingVotingRound: cfg.Finalizer.StartingVotingRound,
	}
}
