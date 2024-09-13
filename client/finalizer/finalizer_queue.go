package finalizer

import (
	"context"
	"flare-fsc/logger"
	"flare-fsc/utils"
	"flare-fsc/utils/contracts/relay"
	"fmt"
	"math/big"
	"sync"
	"time"
)

var relayFunctionSelector []byte

const (
	finalizerQueueProcessorInterval = 100 * time.Millisecond
)

func init() {
	relayABI, err := relay.RelayMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	relayFunctionSelector = relayABI.Methods["relay"].ID
}

type queueItemV2 struct {
	seed          *big.Int
	votingRoundID uint32
	protocolID    uint8
}

func (i *queueItemV2) String() string {
	return fmt.Sprintf("seed=%v, votingRoundID=%v, protocolID=%v", i.seed, i.votingRoundID, i.protocolID)
}

type finalizerQueueV2 struct {
	queue []*queueItemV2

	sync.Mutex
}

type finalizerQueueProcessorV2 struct {
	db            finalizerDB
	queue         *finalizerQueueV2
	delayedQueues *utils.DelayedQueueManager[*queueItemV2]

	finalizationStorage *finalizationStorage
	relayClient         *relayContractClient
	finalizerContext    *finalizerContext
}

func newFinalizerQueueProcessorV2(
	db finalizerDB,
	finalizationStorage *finalizationStorage,
	relayClient *relayContractClient,
	finalizerContext *finalizerContext,
) *finalizerQueueProcessorV2 {
	qp := &finalizerQueueProcessorV2{
		db:                  db,
		finalizationStorage: finalizationStorage,
		relayClient:         relayClient,
		queue:               newFinalizerQueueV2(),

		finalizerContext: finalizerContext,
	}
	qp.delayedQueues = utils.NewDelayedQueueManager[*queueItemV2](qp.processDelayedQueue)
	return qp
}

func newFinalizerQueueV2() *finalizerQueueV2 {
	return &finalizerQueueV2{
		queue: make([]*queueItemV2, 0, 256),
	}
}

func (q *finalizerQueueV2) Add(item *queueItemV2) {
	q.Lock()
	defer q.Unlock()

	q.queue = append(q.queue, item)
}

func (q *finalizerQueueV2) Pop() *queueItemV2 {
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

// Add adds a finalizationItem to the finalization queue
func (p *finalizerQueueProcessorV2) Add(item *FinalizationReady, seed *big.Int) {
	p.queue.Add(&queueItemV2{
		seed:          seed,
		votingRoundID: item.votingRoundID,
		protocolID:    item.protocolID,
	})
}

// Run runs the infinite loops that handles finalization queue.
//
// Should be run in a goroutine.
func (p *finalizerQueueProcessorV2) Run(ctx context.Context) error {
	ticker := time.NewTicker(finalizerQueueProcessorInterval)
	for {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			logger.Info("Finalizer queue processor stopped")
			return ctx.Err()
		}

		item := p.queue.Pop()

		if item == nil {
			continue
		}

		if p.isVoterForCurrentEpoch(item) {
			logger.Info("Finalizer with address %v was selected for item %v", p.relayClient.senderAddress, item)

			p.processItem(ctx, item, false)
		} else {
			logger.Info("Finalizer with address %v will send outside grace period for item %v", p.relayClient.senderAddress, item)

			_, exists := p.finalizationStorage.Get(item.votingRoundID, item.protocolID)
			if exists {
				// Finalization for a votingRoundID should happen in the following voting round votingRoundID + 1
				votingRoundStartTime := p.finalizerContext.votingEpoch.StartTime(int64(item.votingRoundID + 1))
				st := votingRoundStartTime.Add(p.finalizerContext.gracePeriodEndOffset)
				logger.Info("Finalizer will send item %v at %v", item, st)
				p.delayedQueues.Add(st, item)
			}
		}
	}
}

func (p *finalizerQueueProcessorV2) isVoterForCurrentEpoch(item *queueItemV2) bool {
	if item == nil {
		return false
	}
	data, exists := p.finalizationStorage.Get(item.votingRoundID, item.protocolID)
	if !exists {
		return false
	}

	voters, err := data.signingPolicy.voters.SelectVoters(item.seed, item.protocolID, item.votingRoundID, p.finalizerContext.voterThresholdBIPS)
	if err != nil {
		return false
	}

	logger.Debug("Finalizer voters for item %v: %v", item, voters)

	return voters[p.relayClient.senderAddress]
}

func (p *finalizerQueueProcessorV2) processItem(ctx context.Context, item *queueItemV2, isDelayed bool) {
	if item == nil {
		return
	}

	data, exists := p.finalizationStorage.Get(item.votingRoundID, item.protocolID)
	if !exists {
		logger.Warn("finalization data for protocol %d for round %d missing", item.protocolID, item.votingRoundID)
		return
	}

	finalizationData, err := data.PrepareFinalizationResults()
	if err != nil {
		logger.Warn("finalization data preparation for protocol %d for round %d failed - %v", item.protocolID, item.votingRoundID, err)
		return
	}

	txInput, err := finalizationData.PrepareFinalizationTxInput()
	if err != nil {
		logger.Warn("finalization tx input preparation for protocol %d for round %d failed - %v", item.protocolID, item.votingRoundID, err)
		return
	}

	logger.Info("Relaying for round %d for protocol %d", item.votingRoundID, item.protocolID)
	p.relayClient.SubmitPayloadsV2(ctx, txInput, isDelayed)
}

func (p *finalizerQueueProcessorV2) processDelayedQueue(items []*queueItemV2) error {
	now := time.Now()
	currentEpoch := p.finalizerContext.votingEpoch.EpochIndex(now)
	startTime := p.finalizerContext.votingEpoch.StartTime(currentEpoch)

	relayedItems, err := p.relayClient.ProtocolMessageRelayedV2(p.db, startTime, now)
	if err != nil {
		return err
	}

	for _, item := range items {
		if relayedItems[*item] {
			continue
		}
		logger.Info("Finalizer processes delayed queue item %v", item)
		p.processItem(context.TODO(), item, true)
	}
	return nil
}
