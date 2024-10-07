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
	Epoch        *EpochConfig
	timeProvider TimeProvider

	// C is the channel on which the epoch index is sent
	C <-chan int64
}

// NewEpochTicker creates a ticker that sends the epoch index on the channel C
// at the start of the epoch
func NewEpochTicker(epoch *EpochConfig) *EpochTicker {
	c := make(chan int64)
	ticker := &EpochTicker{
		Epoch:        epoch,
		timeProvider: RealTimeProvider{},
		C:            c,
	}
	ticker.start(c)
	return ticker
}

func (t *EpochTicker) start(c chan int64) {
	go func() {
		now := t.timeProvider.Now()
		currentEpoch := t.Epoch.EpochIndex(now)

		epoch := currentEpoch + 1
		epochStart := t.Epoch.StartTime(epoch)

		for {
			<-time.NewTimer(epochStart.Sub(now)).C
			c <- epoch
			epoch += 1
			epochStart = t.Epoch.StartTime(epoch)
			now = t.timeProvider.Now()
		}
	}()
}
