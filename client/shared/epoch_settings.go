package shared

import (
	"time"

	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/system"
)

func RewardEpochTimingFromChain(fsm *system.FlareSystemsManager) (*utils.EpochTimingConfig, error) {
	epochStart, err := fsm.FirstRewardEpochStartTs(nil)
	if err != nil {
		return nil, err
	}
	epochPeriod, err := fsm.RewardEpochDurationSeconds(nil)
	if err != nil {
		return nil, err
	}
	return &utils.EpochTimingConfig{
		Start:  time.Unix(int64(epochStart), 0),
		Period: time.Duration(epochPeriod) * time.Second,
	}, nil
}

func VotingRoundTimingFromChain(fsm *system.FlareSystemsManager) (*utils.EpochTimingConfig, error) {
	epochStart, err := fsm.FirstVotingRoundStartTs(nil)
	if err != nil {
		return nil, err
	}
	epochPeriod, err := fsm.VotingEpochDurationSeconds(nil)
	if err != nil {
		return nil, err
	}
	return utils.NewEpochConfig(
		time.Unix(int64(epochStart), 0),
		time.Duration(epochPeriod)*time.Second,
	), nil
}

// Returns the voting round timing and reward epochs properties from the relay contract.
func EpochsFromChain(relay *relay.Relay) (*utils.EpochTimingConfig, *utils.RewardEpochConfig, error) {
	sd, err := relay.StateData(nil)
	if err != nil {
		return nil, nil, err
	}
	return utils.NewEpochConfig(
			time.Unix(int64(sd.FirstVotingRoundStartTs), 0),
			time.Duration(sd.VotingEpochDurationSeconds)*time.Second,
		), utils.NewRewardEpochConfig(
			int64(sd.FirstRewardEpochStartVotingRoundId),
			int64(sd.RewardEpochDurationInVotingEpochs),
		), nil
}
