package utils

import (
	"math/rand"
	"time"
)

type ShiftedTime struct {
	Shift time.Duration
}

func NewShiftedTime(startNow time.Time) *ShiftedTime {
	shift := time.Until(startNow)
	return &ShiftedTime{Shift: shift}
}

func (s *ShiftedTime) SetNow(startNow time.Time) {
	s.Shift = time.Until(startNow)
}

func (s *ShiftedTime) Now() time.Time {
	return time.Now().Add(s.Shift)
}

func (s *ShiftedTime) SetNowUnix(now int64) {
	s.SetNow(time.Unix(now, 0))
}

func (s *ShiftedTime) AdvanceNow(duration time.Duration) {
	s.Shift += duration
}

// Use when s is the correct RFC3339 time (e.g. in tests, error results in panic)
func ParseTime(s string) time.Time {
	time, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return time
}

func NewRandomizedTicker(interval time.Duration, randomDelta time.Duration) <-chan time.Time {
	deltaIntervalMs := int(randomDelta.Milliseconds())
	ch := make(chan time.Time)
	go func() {
		for {
			d := interval + randomDuration(deltaIntervalMs)
			time.Sleep(d)
			ch <- time.Now()
		}
	}()
	return ch
}

func randomDuration(deltaMs int) time.Duration {
	delta := int64(0)
	if deltaMs > 0 {
		delta = int64(rand.Intn(deltaMs))
	}
	return time.Duration(delta * int64(time.Millisecond))
}
