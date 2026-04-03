package utils

import (
	"context"
	"sync"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

type QueueProcessorFunc[T any] func(context.Context, []T) error

type DelayedQueueManager[T any] struct {
	timeMap map[time.Time][]T

	processor QueueProcessorFunc[T]

	sync.Mutex
}

func NewDelayedQueueManager[T any](processor QueueProcessorFunc[T]) *DelayedQueueManager[T] {
	return &DelayedQueueManager[T]{
		timeMap:   make(map[time.Time][]T),
		processor: processor,
	}
}

func (l *DelayedQueueManager[T]) Add(ctx context.Context, t time.Time, item T) {
	if t.Before(time.Now()) {
		return
	}

	l.Lock()
	defer l.Unlock()

	if _, ok := l.timeMap[t]; !ok {
		l.createTimer(ctx, t)
	}
	l.timeMap[t] = append(l.timeMap[t], item)
}

func (l *DelayedQueueManager[T]) Get(t time.Time) []T {
	l.Lock()
	defer l.Unlock()

	items := l.timeMap[t]
	delete(l.timeMap, t)
	return items
}

func (l *DelayedQueueManager[T]) createTimer(ctx context.Context, t time.Time) {
	go func() {
		timer := time.NewTimer(time.Until(t))
		select {
		case <-timer.C:
			items := l.Get(t)
			if err := l.processor(ctx, items); err != nil {
				logger.Errorf("DelayedQueueManager processor error: %s", err)
			}
		case <-ctx.Done():
			l.Get(t) // remove from map
		}
	}()
}
