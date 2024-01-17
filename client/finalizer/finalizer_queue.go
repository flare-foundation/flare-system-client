package finalizer

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
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

func newFinalizerQueue() *finalizerQueue {
	return &finalizerQueue{
		queue: make([]*queueItem, 0),
	}
}

func (q *finalizerQueue) Add(votingRoundId uint32, protocolId byte, messageHash common.Hash) {
	q.Lock()
	defer q.Unlock()

	q.queue = append(q.queue, &queueItem{
		votingRoundId: votingRoundId,
		protocolId:    protocolId,
		messageHash:   messageHash,
	})
}
