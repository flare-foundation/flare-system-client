package utils

import "time"

type Epoch struct {
	Start  time.Time
	Period time.Duration
}

func NewEpoch(start time.Time, duration time.Duration) *Epoch {
	return &Epoch{
		Start:  start,
		Period: duration,
	}
}

func (e Epoch) EpochIndex(t time.Time) int64 {
	return int64(t.Sub(e.Start) / e.Period)
}

func (e Epoch) StartTime(epoch int64) time.Time {
	return e.Start.Add(time.Duration(epoch) * e.Period)
}

func (e Epoch) EndTime(epoch int64) time.Time {
	return e.StartTime(epoch + 1)
}

func (e Epoch) TimeRange(epoch int64) (time.Time, time.Time) {
	return e.StartTime(epoch), e.EndTime(epoch)
}

type IntEpoch struct {
	Start  int64
	Period int64
}

func NewIntEpoch(start int64, period int64) *IntEpoch {
	return &IntEpoch{
		Start:  start,
		Period: period,
	}
}

func (e IntEpoch) EpochIndex(n int64) int64 {
	return (n - e.Start) / e.Period
}

func (e IntEpoch) StartEpoch(n int64) int64 {
	return e.Start + e.Period*e.EpochIndex(n)
}

func (e IntEpoch) EndEpoch(n int64) int64 {
	return e.Start + e.Period*(e.EpochIndex(n)+1)
}
