package shared

import (
	"flare-tlc/utils"
	"flare-tlc/utils/contracts/system"
	"time"
)

func RewardEpochFromChain(fsm *system.FlareSystemManager) (*utils.Epoch, error) {
	epochStart, err := fsm.RewardEpochsStartTs(nil)
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

func VotingEpochFromChain(fsm *system.FlareSystemManager) (*utils.Epoch, error) {
	epochStart, err := fsm.FirstVotingRoundStartTs(nil)
	if err != nil {
		return nil, err
	}
	epochPeriod, err := fsm.VotingEpochDurationSeconds(nil)
	if err != nil {
		return nil, err
	}
	return &utils.Epoch{
		Start:  time.Unix(int64(epochStart), 0),
		Period: time.Duration(epochPeriod) * time.Second,
	}, nil
}
