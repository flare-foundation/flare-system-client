package utils

import (
	"flare-fsc/logger"
	"sync"
	"time"
)

type QueueProcessorFunc[T any] func([]T) error

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

func (l *DelayedQueueManager[T]) Add(t time.Time, item T) {
	if t.Before(time.Now()) {
		return
	}

	l.Lock()
	defer l.Unlock()

	if _, ok := l.timeMap[t]; !ok {
		l.createTimer(t)
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

func (l *DelayedQueueManager[T]) createTimer(t time.Time) {
	go func() {
		timer := time.NewTimer(time.Until(t))
		<-timer.C
		items := l.Get(t)
		if err := l.processor(items); err != nil {
			logger.Error("DelayedQueueManager processor error: %s", err)
		}
	}()
}
