package finalizer

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	finalizerQueueProcessorInterval = 100 * time.Millisecond
)

type queueItem struct {
	votingRoundId uint32
	protocolId    byte
	messageHash   common.Hash
}

type finalizerQueue struct {
	queue []*queueItem

	sync.Mutex
}

type finalizerQueueProcessor struct {
	queue finalizerQueue
}

func newFinalizerQueueProcessor() *finalizerQueueProcessor {
	return &finalizerQueueProcessor{
		queue: newFinalizerQueue(),
	}
}

func newFinalizerQueue() finalizerQueue {
	return finalizerQueue{
		queue: make([]*queueItem, 0, 256),
	}
}

func (q *finalizerQueue) Add(item *queueItem) {
	q.Lock()
	defer q.Unlock()

	q.queue = append(q.queue, item)
}

func (q *finalizerQueue) Pop() *queueItem {
	q.Lock()
	defer q.Unlock()

	if len(q.queue) == 0 {
		return nil
	}

	item := q.queue[0]
	q.queue[0] = nil
	q.queue = q.queue[1:]
	return item
}

func (p *finalizerQueueProcessor) Add(item *submitterPayloadItem) {
	p.queue.Add(&queueItem{
		votingRoundId: item.votingRoundId,
		protocolId:    item.protocolId,
		messageHash:   item.payload.messageHash,
	})
}

// Infinite loop, should be run in a goroutine
func (p *finalizerQueueProcessor) Process() {
	ticker := time.NewTicker(finalizerQueueProcessorInterval)
	for {
		<-ticker.C

		item := p.queue.Pop()
		if item != nil {
		}
	}
}
