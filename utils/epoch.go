package utils

import "time"

type EpochConfig struct {
	Start  time.Time     // start of the epoch with index 0
	Period time.Duration // length of one epoch
}

func NewEpochConfig(start time.Time, duration time.Duration) *EpochConfig {
	return &EpochConfig{
		Start:  start,
		Period: duration,
	}
}

// EpochIndex returns returns the consecutive index of the epoch that takes place at time t.
func (e EpochConfig) EpochIndex(t time.Time) int64 {
	return int64(t.Sub(e.Start) / e.Period)
}

// StartTime returns the start time of the epoch.
func (e EpochConfig) StartTime(epoch int64) time.Time {
	return e.Start.Add(time.Duration(epoch) * e.Period)
}

// EndTime returns the end time of the epoch.
func (e EpochConfig) EndTime(epoch int64) time.Time {
	return e.StartTime(epoch + 1)
}

type RewardEpochConfig struct {
	Start  int64 // index of the voting epoch in which the reward epoch with index 0 starts
	Period int64 // length of one reward epoch in voting epochs
}

func NewRewardEpochConfig(start int64, period int64) *RewardEpochConfig {
	return &RewardEpochConfig{
		Start:  start,
		Period: period,
	}
}

// StartEpoch returns the index of the voting epoch in which the n-th reward epoch is expected to start.
func (e RewardEpochConfig) StartEpoch(n int64) int64 {
	return e.Start + e.Period*n
}

// EndEpoch returns the index of the voting epoch in which the n-th reward epoch is expected to end.
func (e RewardEpochConfig) EndEpoch(n int64) int64 {
	return e.Start + e.Period*(n+1)
}
