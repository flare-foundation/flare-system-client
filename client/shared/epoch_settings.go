package shared

import (
	"time"
)

type EpochSettings struct {
	RewardEpochStartSec      uint64
	RewardEpochDurationSec   uint64
	FirstVotingEpochStartSec uint64
	VotingEpochDurationSec   uint64
}

func (e *EpochSettings) VotingEpochForTime(t time.Time) uint64 {
	unixSeconds := uint64(t.Unix())
	return (unixSeconds - e.FirstVotingEpochStartSec) / e.VotingEpochDurationSec
}

func (e *EpochSettings) NextVotingEpochStart(t time.Time) time.Time {
	currentEpoch := e.VotingEpochForTime(t)
	nextEpochStartSec := e.FirstVotingEpochStartSec + (currentEpoch+1)*e.VotingEpochDurationSec
	return time.Unix(int64(nextEpochStartSec), 0)
}

func (e *EpochSettings) RewardEpochForTime(t time.Time) uint64 {
	unixSeconds := uint64(t.Unix())
	return (unixSeconds - e.RewardEpochStartSec) / e.RewardEpochDurationSec
}
