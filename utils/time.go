package utils

import (
	"math/rand"
	"time"
)

type TimeProvider interface {
	Now() time.Time
}

type RealTimeProvider struct{}

func (RealTimeProvider) Now() time.Time {
	return time.Now()
}

type FixedTimeProvider struct {
	Time time.Time
}

func (f FixedTimeProvider) Now() time.Time {
	return f.Time
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

type EpochTicker struct {
	epoch        *Epoch
	timeProvider TimeProvider

	// C is the channel on which the epoch index is sent
	C <-chan int64
}

// NewEpochTicker creates a ticker that sends the epoch index on the channel C
// at the start of the epoch + offset
func NewEpochTicker(offset time.Duration, epoch *Epoch) *EpochTicker {
	c := make(chan int64)
	ticker := &EpochTicker{
		epoch:        epoch,
		timeProvider: RealTimeProvider{},
		C:            c,
	}
	ticker.start(c, offset)
	return ticker
}

func (t *EpochTicker) start(c chan int64, offset time.Duration) {
	go func() {
		now := t.timeProvider.Now()
		epoch := t.epoch.EpochIndex(now)
		epochStart := t.epoch.StartTime(epoch)
		delayToStart := epochStart.Add(offset).Sub(now)

		for {
			if delayToStart < 0 {
				delayToStart = 0
			}
			<-time.NewTimer(delayToStart).C
			c <- epoch
			epoch += 1
			epochStart = t.epoch.StartTime(epoch)
			delayToStart = epochStart.Add(offset).Sub(t.timeProvider.Now())
		}
	}()
}

type RetriableEpochTicker struct {
	epoch *Epoch

	C <-chan int64
}

// RetriableEpochTicker is an EpochTicker that sends the epoch index on the channel C
// at the start of the epoch + offset, and then retries every retryPeriod until the end of the epoch
func NewRetriableEpochTicker(offset time.Duration, retryPeriod time.Duration, epoch *Epoch) *RetriableEpochTicker {
	c := make(chan int64)
	ticker := &RetriableEpochTicker{
		epoch: epoch,
		C:     c,
	}
	ticker.start(c, offset, retryPeriod)
	return ticker
}

func (ret *RetriableEpochTicker) start(c chan int64, offset time.Duration, retryPeriod time.Duration) {
	go func() {
		epochTicker := NewEpochTicker(offset, ret.epoch)
		for {
			epoch := <-epochTicker.C
			c <- epoch

			go func() {
				currentEpoch := epoch
				retryTicker := time.NewTicker(retryPeriod)
				for {
					t := <-retryTicker.C
					if !t.Before(ret.epoch.EndTime(currentEpoch)) {
						retryTicker.Stop()
						break
					} else {
						c <- currentEpoch
					}
				}
			}()
		}
	}()
}
