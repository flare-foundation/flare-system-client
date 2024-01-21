package finalizer

import (
	"sync"
	"time"

	"golang.org/x/exp/slices"
)

const (
	finalizerQueueProcessorInterval = 100 * time.Millisecond
)

type finalizerQueue struct {
	// waiting to be processed
	queue []*storageItemKey

	sync.Mutex
}

type finalizerQueueProcessor struct {
	queue finalizerQueue

	submissionStorage *submissionStorage
	relayClient       *relayContractClient
}

func newFinalizerQueueProcessor(
	submissionStorage *submissionStorage,
	relayClient *relayContractClient,
) *finalizerQueueProcessor {
	return &finalizerQueueProcessor{
		submissionStorage: submissionStorage,
		relayClient:       relayClient,
		queue:             newFinalizerQueue(),
	}
}

func newFinalizerQueue() finalizerQueue {
	return finalizerQueue{
		queue: make([]*storageItemKey, 0, 256),
	}
}

func (q *finalizerQueue) Add(item *storageItemKey) {
	q.Lock()
	defer q.Unlock()

	q.queue = append(q.queue, item)
}

func (q *finalizerQueue) Pop() *storageItemKey {
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
	p.queue.Add(&storageItemKey{
		votingRoundId: item.votingRoundId,
		protocolId:    item.protocolId,
		messageHash:   item.payload.messageHash,
	})
}

// Infinite loop, should be run in a goroutine
func (p *finalizerQueueProcessor) Run() {
	ticker := time.NewTicker(finalizerQueueProcessorInterval)
	for {
		<-ticker.C

		item := p.queue.Pop()
		selectedPayloads, signingPolicy := p.processItem(item)
		p.relayClient.SubmitPayloads(selectedPayloads, signingPolicy)
	}
}

func (p *finalizerQueueProcessor) processItem(item *storageItemKey) ([]*signedPayload, *signingPolicy) {
	if item == nil {
		return nil, nil
	}
	data := p.submissionStorage.Get(item.votingRoundId, item.protocolId, item.messageHash)
	if data == nil {
		return nil, nil
	}

	payloads := make([]*signedPayload, 0, len(data.payload))
	for _, payload := range data.payload {
		if payload != nil {
			payloads = append(payloads, payload)
		}
	}

	// (sort descreasing by weight)
	slices.SortFunc(payloads, func(p, q *signedPayload) bool {
		return data.signingPolicy.weights[p.index] > data.signingPolicy.weights[q.index]
	})

	// greedy select wuntil threshold is reached
	weight := uint16(0)
	var selected []*signedPayload
	for _, payload := range payloads {
		weight += data.signingPolicy.weights[payload.index]
		selected = append(selected, payload)
		if weight > data.signingPolicy.threshold {
			break
		}
	}

	// sort selected payloads by index
	slices.SortFunc(selected, func(p, q *signedPayload) bool {
		return p.index < q.index
	})
	return selected, data.signingPolicy
}
