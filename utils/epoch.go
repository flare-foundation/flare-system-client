package utils

import "time"

type EpochTimingConfig struct {
	Start  time.Time     // start of the epoch with index 0
	Period time.Duration // length of one epoch
}

func NewEpochConfig(start time.Time, duration time.Duration) *EpochTimingConfig {
	return &EpochTimingConfig{
		Start:  start,
		Period: duration,
	}
}

// EpochIndex returns returns the consecutive index of the epoch that takes place at time t.
func (e EpochTimingConfig) EpochIndex(t time.Time) int64 {
	return int64(t.Sub(e.Start) / e.Period)
}

// StartTime returns the start time of the epoch.
func (e EpochTimingConfig) StartTime(epoch int64) time.Time {
	return e.Start.Add(time.Duration(epoch) * e.Period)
}

// EndTime returns the end time of the epoch.
func (e EpochTimingConfig) EndTime(epoch int64) time.Time {
	return e.StartTime(epoch + 1)
}

type RewardEpochConfig struct {
	Start  int64 // index of the voting round in which the reward epoch with index 0 starts
	Period int64 // length of one reward epoch in voting rounds
}

func NewRewardEpochConfig(start int64, period int64) *RewardEpochConfig {
	return &RewardEpochConfig{
		Start:  start,
		Period: period,
	}
}

// StartEpoch returns the index of the voting round in which the given reward epoch is expected to start.
func (e RewardEpochConfig) StartEpoch(rewardEpoch int64) int64 {
	return e.Start + e.Period*rewardEpoch
}

// EndEpoch returns the index of the voting round in which the given reward epoch is expected to end.
func (e RewardEpochConfig) EndEpoch(rewardEpoch int64) int64 {
	return e.Start + e.Period*(rewardEpoch+1)
}
