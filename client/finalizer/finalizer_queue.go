package finalizer

import (
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

type finalizerQueue struct {
	queue []*queueItem

	sync.Mutex
}

type finalizerQueueProcessor struct {
	queue finalizerQueue

	submissionStorage *submissionStorage
	relayClient       *relayContractClient
	finalizerContext  *finalizerContext
}

func newFinalizerQueueProcessor(
	submissionStorage *submissionStorage,
	relayClient *relayContractClient,
	finalizerContext *finalizerContext,
) *finalizerQueueProcessor {
	return &finalizerQueueProcessor{
		submissionStorage: submissionStorage,
		relayClient:       relayClient,
		queue:             newFinalizerQueue(),
		finalizerContext:  finalizerContext,
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
func (p *finalizerQueueProcessor) Run() {
	ticker := time.NewTicker(finalizerQueueProcessorInterval)
	for {
		<-ticker.C

		item := p.queue.Pop()

		// TODO:!!!!
		// Check if we are among selected voters for the voting epoch
		// If not call process item, but add to delayed queue
		if p.IsVoterForCurrentEpoch(item) {
			selectedPayloads, signingPolicy := p.processItem(item)
			p.relayClient.SubmitPayloads(selectedPayloads, signingPolicy)
		} else {
			// TODO: add to delayed queue
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
		return false // TODO: log error or panic? -- wrong bips value
	}
	return voters.Contains(p.relayClient.senderAddress)
}

func (p *finalizerQueueProcessor) processItem(item *queueItem) ([]*signedPayload, *signingPolicy) {
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
	return selected, data.signingPolicy
}
