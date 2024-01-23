package finalizer

import (
	"flare-tlc/client/config"
	"flare-tlc/client/shared"
	"flare-tlc/utils"
	"flare-tlc/utils/contracts/system"
)

// Finalizer client settings
type finalizerContext struct {
	startingRewardEpoch int64
	startingVotingRound uint32

	votingEpoch *utils.Epoch
}

func newFinalizerContext(cfg *config.ClientConfig, systemManager *system.FlareSystemManager) (*finalizerContext, error) {
	votingEpoch, err := shared.VotingEpochFromChain(systemManager)
	if err != nil {
		return nil, err
	}
	return &finalizerContext{
		startingRewardEpoch: cfg.Finalizer.StartingRewardEpoch,
		startingVotingRound: cfg.Finalizer.StartingVotingRound,
		votingEpoch:         votingEpoch,
	}, nil
}
