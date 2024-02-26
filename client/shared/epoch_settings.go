package shared

import (
	"flare-tlc/utils"
	"flare-tlc/utils/contracts/relay"
	"flare-tlc/utils/contracts/system"
	"time"
)

func RewardEpochFromChain(fsm *system.FlareSystemsManager) (*utils.Epoch, error) {
	epochStart, err := fsm.FirstRewardEpochStartTs(nil)
	if err != nil {
		return nil, err
	}
	epochPeriod, err := fsm.RewardEpochDurationSeconds(nil)
	if err != nil {
		return nil, err
	}
	return &utils.Epoch{
		Start:  time.Unix(int64(epochStart), 0),
		Period: time.Duration(epochPeriod) * time.Second,
	}, nil
}

func VotingEpochFromChain(fsm *system.FlareSystemsManager) (*utils.Epoch, error) {
	epochStart, err := fsm.FirstVotingRoundStartTs(nil)
	if err != nil {
		return nil, err
	}
	epochPeriod, err := fsm.VotingEpochDurationSeconds(nil)
	if err != nil {
		return nil, err
	}
	return utils.NewEpoch(
		time.Unix(int64(epochStart), 0),
		time.Duration(epochPeriod)*time.Second,
	), nil
}

// Returns the voting and reward epochs from the relay contract
func EpochsFromChain(relay *relay.Relay) (*utils.Epoch, *utils.IntEpoch, error) {
	sd, err := relay.StateData(nil)
	if err != nil {
		return nil, nil, err
	}
	return utils.NewEpoch(
			time.Unix(int64(sd.FirstVotingRoundStartTs), 0),
			time.Duration(sd.VotingEpochDurationSeconds)*time.Second,
		), utils.NewIntEpoch(
			int64(sd.FirstRewardEpochStartVotingRoundId),
			int64(sd.RewardEpochDurationInVotingEpochs),
		), nil
}
