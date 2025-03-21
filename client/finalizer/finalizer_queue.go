package finalizer

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
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

type queueItem struct {
	seed          *big.Int
	votingRoundID uint32
	protocolID    uint8
	msgHash       common.Hash
}

func (i *queueItem) String() string {
	return fmt.Sprintf("seed=%v, votingRoundID=%v, protocolID=%v", i.seed, i.votingRoundID, i.protocolID)
}

type finalizerQueue struct {
	queue []*queueItem

	sync.Mutex
}

type finalizerQueueProcessor struct {
	db            finalizerDB
	queue         *finalizerQueue
	delayedQueues *utils.DelayedQueueManager[*queueItem]

	finalizationStorage *finalizationStorage
	relayClient         *relayContractClient
	finalizerContext    *finalizerContext
}

func newFinalizerQueueProcessor(
	db finalizerDB,
	finalizationStorage *finalizationStorage,
	relayClient *relayContractClient,
	finalizerContext *finalizerContext,
) *finalizerQueueProcessor {
	qp := &finalizerQueueProcessor{
		db:                  db,
		finalizationStorage: finalizationStorage,
		relayClient:         relayClient,
		queue:               newFinalizerQueue(),

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

// Add adds a finalizationItem to the finalization queue
func (p *finalizerQueueProcessor) Add(item *FinalizationReady, seed *big.Int) {
	p.queue.Add(&queueItem{
		seed:          seed,
		votingRoundID: item.votingRoundID,
		protocolID:    item.protocolID,
		msgHash:       item.msgHash,
	})
}

// Run runs the infinite loops that handles finalization queue.
//
// Should be run in a goroutine.
func (p *finalizerQueueProcessor) Run(ctx context.Context) error {
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
			logger.Infof("Finalizer with address %v was selected for voting round %v for protocol %v", p.relayClient.senderAddress, item.votingRoundID, item.protocolID)

			p.processItem(ctx, item, false)
		} else {
			logger.Infof("Finalizer with address %v will send outside grace period for voting round %v for protocol %v", p.relayClient.senderAddress, item.votingRoundID, item.protocolID)

			_, exists := p.finalizationStorage.Get(item.votingRoundID, item.protocolID, item.msgHash)
			if exists {
				// Finalization for a votingRoundID should happen in the following voting round votingRoundID + 1
				votingRoundStartTime := p.finalizerContext.votingRoundTiming.StartTime(int64(item.votingRoundID + 1))
				st := votingRoundStartTime.Add(p.finalizerContext.gracePeriodEndOffset)

				if st.Before(time.Now()) {
					logger.Debugf("Finalizer will send now for voting round %v for protocol %v", item.votingRoundID, item.protocolID)
					p.processItem(ctx, item, true)
				}
				p.delayedQueues.Add(st, item)
			} else {
				logger.Errorf("Finalizer missing finalization data for protocol %v in votingRound %v", item.protocolID, item.votingRoundID)
			}
		}
	}
}

// isVoterForCurrentEpoch checks whether the voter is selected to finalize an item in the prioritized period.
func (p *finalizerQueueProcessor) isVoterForCurrentEpoch(item *queueItem) bool {
	if item == nil {
		return false
	}
	data, exists := p.finalizationStorage.Get(item.votingRoundID, item.protocolID, item.msgHash)
	if !exists {
		return false
	}

	voters, err := data.signingPolicy.Voters.SelectVoters(item.seed, item.protocolID, item.votingRoundID, p.finalizerContext.voterThresholdBIPS)
	if err != nil {
		return false
	}

	return voters[p.relayClient.senderAddress]
}

// processItem prepares and sends finalization transaction for item.
func (p *finalizerQueueProcessor) processItem(ctx context.Context, item *queueItem, isDelayed bool) {
	if item == nil {
		return
	}

	data, exists := p.finalizationStorage.Get(item.votingRoundID, item.protocolID, item.msgHash)
	if !exists {
		logger.Warnf("finalization data for protocol %d for round %d missing", item.protocolID, item.votingRoundID)
		return
	}

	finalizationData, err := PrepareFinalizationResults(data)
	if err != nil {
		logger.Warnf("finalization data preparation for protocol %d for round %d failed - %v", item.protocolID, item.votingRoundID, err)
		return
	}

	txInput, err := finalizationData.PrepareFinalizationTxInput()
	if err != nil {
		logger.Warnf("finalization tx input preparation for protocol %d for round %d failed - %v", item.protocolID, item.votingRoundID, err)
		return
	}

	logger.Infof("Relaying for round %d for protocol %d", item.votingRoundID, item.protocolID)
	p.relayClient.SubmitPayloads(ctx, txInput, isDelayed, item.protocolID)
}

func (p *finalizerQueueProcessor) processDelayedQueue(items []*queueItem) error {
	now := time.Now()
	currentEpoch := p.finalizerContext.votingRoundTiming.EpochIndex(now)
	startTime := p.finalizerContext.votingRoundTiming.StartTime(currentEpoch)

	relayedItems, err := p.relayClient.ProtocolMessageRelayed(p.db, startTime, now)
	if err != nil {
		return err
	}

	for _, item := range items {
		if relayedItems[*item] {
			continue
		}
		logger.Infof("Finalizer processes delayed queue item for round %v for protocol %v", item.votingRoundID, item.protocolID)
		p.processItem(context.TODO(), item, true)
	}
	return nil
}
