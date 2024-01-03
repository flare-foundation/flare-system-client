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
