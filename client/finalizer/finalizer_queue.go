package finalizer

import (
	"flare-tlc/logger"
	"flare-tlc/utils"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/exp/slices"
)

const (
	finalizerQueueProcessorInterval = 100 * time.Millisecond
)

type queueItem struct {
	votingRoundId uint32
	protocolId    byte
	messageHash   common.Hash
}

func (i *queueItem) String() string {
	return fmt.Sprintf("votingRoundId=%v, protocolId=%v, messageHash=%v", i.votingRoundId, i.protocolId, i.messageHash.Hex())
}

type finalizerQueue struct {
	queue []*queueItem

	sync.Mutex
}

type finalizerQueueProcessor struct {
	queue         *finalizerQueue
	delayedQueues *utils.DelayedQueueManager[*queueItem]

	submissionStorage *submissionStorage
	relayClient       *relayContractClient
	finalizerContext  *finalizerContext
}

func newFinalizerQueueProcessor(
	submissionStorage *submissionStorage,
	relayClient *relayContractClient,
	finalizerContext *finalizerContext,
) *finalizerQueueProcessor {
	qp := &finalizerQueueProcessor{
		submissionStorage: submissionStorage,
		relayClient:       relayClient,
		queue:             newFinalizerQueue(),

		finalizerContext: finalizerContext,
	}
	qp.delayedQueues = utils.NewDelayedQueueManager[*queueItem](qp.processDelayedQueue)
	return qp
}

func newFinalizerQueue() *finalizerQueue {
	return &finalizerQueue{
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
func (p *finalizerQueueProcessor) Run() {
	ticker := time.NewTicker(finalizerQueueProcessorInterval)
	for {
		<-ticker.C

		item := p.queue.Pop()

		if item == nil {
			continue
		}

		if p.IsVoterForCurrentEpoch(item) {
			logger.Debug("Finalizer with address %v was selected for item %v", p.relayClient.senderAddress, item)

			p.processItem(item)
		} else {
			logger.Debug("Finalizer with address %v will send outside grace period for item %v", p.relayClient.senderAddress, item)

			data := p.submissionStorage.Get(item.votingRoundId, item.protocolId, item.messageHash)
			if data != nil {
				st := p.finalizerContext.votingEpoch.StartTime(int64(item.votingRoundId)).Add(p.finalizerContext.gracePeriodEndOffset)
				p.delayedQueues.Add(st, item)
			}
		}
	}
}

func (p *finalizerQueueProcessor) IsVoterForCurrentEpoch(item *queueItem) bool {
	if item == nil {
		return false
	}
	data := p.submissionStorage.Get(item.votingRoundId, item.protocolId, item.messageHash)
	if data == nil {
		return false
	}
	voters, err := data.signingPolicy.voters.SelectVoters(item.protocolId, item.votingRoundId, p.finalizerContext.voterThresholdBIPS)
	if err != nil {
		return false
	}

	logger.Debug("Finalizer voters for item %v: %v", item, voters)

	return voters.Contains(p.relayClient.senderAddress)
}

func (p *finalizerQueueProcessor) processItem(item *queueItem) {
	if item == nil {
		return
	}
	data := p.submissionStorage.Get(item.votingRoundId, item.protocolId, item.messageHash)
	if data == nil {
		return
	}

	payloads := make([]*signedPayload, 0, len(data.payload))
	for _, payload := range data.payload {
		if payload != nil {
			payloads = append(payloads, payload)
		}
	}

	// (sort descreasing by weight)
	slices.SortFunc(payloads, func(p, q *signedPayload) bool {
		return data.signingPolicy.voters.VoterWeight(p.index) > data.signingPolicy.voters.VoterWeight(q.index)
	})

	// greedy select until threshold is reached
	weight := uint16(0)
	var selected []*signedPayload
	for _, payload := range payloads {
		weight += data.signingPolicy.voters.VoterWeight(payload.index)
		selected = append(selected, payload)
		if weight > data.signingPolicy.threshold {
			break
		}
	}

	// sort selected payloads by index
	slices.SortFunc(selected, func(p, q *signedPayload) bool {
		return p.index < q.index
	})

	p.relayClient.SubmitPayloads(selected, data.signingPolicy)
}

func (p *finalizerQueueProcessor) processDelayedQueue(items []*queueItem) error {
	for _, item := range items {
		logger.Debug("Finalizer processes delayed queue item %v", item)
		p.processItem(item)
	}
	return nil
}
