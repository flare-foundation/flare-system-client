package finalizer

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

var relayFunctionSelector []byte

const (
	finalizerQueueProcessorInterval = 100 * time.Millisecond

	// Bounds how long one item's send may hold Run's loop (a full retry budget is
	// ~11 min); grace items cut off by it retry via the delayed queue.
	queueSendTimeout = 50 * time.Second

	// Retry delay for an item whose grace-period-end target has already passed:
	// the delayed queue drops past-time targets, so schedule slightly ahead
	// instead, or a send cut off by queueSendTimeout would lose its retry.
	delayedRetryDelay = 5 * time.Second

	// A bounded send divides its remaining deadline into this many attempts so
	// gas-bumped replacements still fire inside queueSendTimeout; must stay ≥2.
	boundedSendAttempts = 3

	// Floor for a derived per-attempt timeout when little deadline remains.
	minAttemptTimeout = 2 * time.Second

	// Part of a bounded send's window reserved for the send attempts: the nonce
	// prefetch may consume the rest, so a flaky fetch cannot starve the sends.
	sendPhaseReserve = 30 * time.Second
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
	digest        common.Hash
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
	randomSource        randomSource
}

func newFinalizerQueueProcessor(
	db finalizerDB,
	finalizationStorage *finalizationStorage,
	relayClient *relayContractClient,
	finalizerContext *finalizerContext,
	randomSource randomSource,
) *finalizerQueueProcessor {
	qp := &finalizerQueueProcessor{
		db:                  db,
		finalizationStorage: finalizationStorage,
		relayClient:         relayClient,
		queue:               newFinalizerQueue(),
		randomSource:        randomSource,

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
	queued := &queueItem{
		seed:          seed,
		votingRoundID: item.votingRoundID,
		protocolID:    item.protocolID,
		digest:        item.digest,
	}

	p.queue.Add(queued)
}

// needsRandomTrailer reports whether item's finalization must carry the random
// number and its Merkle proof: only the random protocol, and only from the breaking
// reward epoch on, since only the new Relay verifies and stores them.
func (p *finalizerQueueProcessor) needsRandomTrailer(item *queueItem) bool {
	if item.protocolID != p.finalizerContext.randomNumberProtocolID {
		return false
	}
	data, exists := p.finalizationStorage.get(item.votingRoundID, item.protocolID, item.digest)
	if !exists {
		return false
	}
	return p.relayClient.relayCutover.NewRelayFromRewardEpoch(data.signingPolicy.RewardEpochID)
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

			p.processItemBounded(ctx, item, false)

			// Retry via the delayed queue; it skips rounds already relayed.
			p.delayedQueues.Add(ctx, p.delayedRetryTime(item.votingRoundID, time.Now()), item)
		} else {
			logger.Infof("Finalizer with address %v will send outside grace period for voting round %v for protocol %v", p.relayClient.senderAddress, item.votingRoundID, item.protocolID)

			_, exists := p.finalizationStorage.get(item.votingRoundID, item.protocolID, item.digest)
			if exists {
				// Finalization for a votingRoundID should happen in the following voting round votingRoundID + 1
				votingRoundStartTime := p.finalizerContext.votingRoundTiming.StartTime(int64(item.votingRoundID + 1))
				st := votingRoundStartTime.Add(p.finalizerContext.gracePeriodEndOffset)

				if st.Before(time.Now()) {
					logger.Debugf("Finalizer will send now for voting round %v for protocol %v", item.votingRoundID, item.protocolID)
					p.processItemBounded(ctx, item, true)
				}
				p.delayedQueues.Add(ctx, p.delayedRetryTime(item.votingRoundID, time.Now()), item)
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
	data, exists := p.finalizationStorage.get(item.votingRoundID, item.protocolID, item.digest)
	if !exists {
		return false
	}

	voters, err := data.signingPolicy.Voters.SelectVoters(item.seed, item.protocolID, item.votingRoundID, p.finalizerContext.voterThresholdBIPS)
	if err != nil {
		return false
	}

	return voters[p.relayClient.senderAddress]
}

// delayedRetryTime returns when item's delayed-queue retry should run: at the
// grace-period end of the following round, or delayedRetryDelay ahead of now
// when that has already passed.
func (p *finalizerQueueProcessor) delayedRetryTime(votingRoundID uint32, now time.Time) time.Time {
	st := p.finalizerContext.votingRoundTiming.StartTime(int64(votingRoundID + 1)).Add(p.finalizerContext.gracePeriodEndOffset)
	if st.After(now) {
		return st
	}
	// truncate: timeMap keys compare exactly; sub-second/monotonic parts would defeat batching
	return now.Add(delayedRetryDelay).Truncate(time.Second)
}

// processItemBounded runs processItem under queueSendTimeout so one item cannot
// head-of-line-block the queue; delayed-queue sends stay unbounded.
func (p *finalizerQueueProcessor) processItemBounded(ctx context.Context, item *queueItem, isDelayed bool) {
	sendCtx, cancel := context.WithTimeout(ctx, queueSendTimeout)
	defer cancel()
	p.processItem(sendCtx, item, isDelayed)
}

// processItem prepares and sends finalization transaction for item.
func (p *finalizerQueueProcessor) processItem(ctx context.Context, item *queueItem, isDelayed bool) {
	if item == nil {
		return
	}

	data, exists := p.finalizationStorage.get(item.votingRoundID, item.protocolID, item.digest)
	if !exists {
		logger.Warnf("finalization data for protocol %d for round %d missing", item.protocolID, item.votingRoundID)
		return
	}

	finalizationData, err := PrepareFinalizationResults(data)
	if err != nil {
		logger.Warnf("finalization data preparation for protocol %d for round %d failed - %v", item.protocolID, item.votingRoundID, err)
		return
	}

	if p.needsRandomTrailer(item) {
		// Without a proof the Relay reverts, so there is nothing to send: log the
		// missed duty and leave the retry to the delayed queue, by which time the
		// message may have arrived or another finalizer may have relayed the round.
		trailer, err := p.randomTrailer(ctx, finalizationData.message, item.votingRoundID)
		if err != nil {
			logger.Errorf("no random proof to finalize protocol %d for round %d, not sending: %v",
				item.protocolID, item.votingRoundID, err)
			return
		}
		finalizationData.randomTrailer = trailer
	}

	txInput, err := finalizationData.PrepareFinalizationTxInput()
	if err != nil {
		logger.Warnf("finalization tx input preparation for protocol %d for round %d failed - %v", item.protocolID, item.votingRoundID, err)
		return
	}

	// the tx carries the policy bytes, so it must go to the Relay storing that policy's hash
	address := p.relayClient.addressForRewardEpoch(data.signingPolicy.RewardEpochID)

	logger.Infof("Relaying for round %d for protocol %d to %s (delayed=%t)", item.votingRoundID, item.protocolID, address, isDelayed)
	p.relayClient.SubmitPayloads(ctx, address, txInput, isDelayed, item.protocolID, item.votingRoundID)
}

// randomTrailer takes the round's random number and Merkle proof and checks them
// against the signed message before they can reach the chain.
func (p *finalizerQueueProcessor) randomTrailer(ctx context.Context, message shared.Message, votingRoundID uint32) ([]byte, error) {
	parsed, err := message.Parse()
	if err != nil {
		return nil, err
	}
	if parsed.VotingRoundID != votingRoundID {
		return nil, fmt.Errorf("message is for round %d, not %d", parsed.VotingRoundID, votingRoundID)
	}

	data, err := p.randomSource.Get(ctx, votingRoundID)
	if err != nil {
		return nil, err
	}
	// checked again against the finalized message: the local provider's message,
	// which the data was checked against on arrival, need not be the one that won
	if err := data.verify(parsed); err != nil {
		return nil, err
	}

	return data.trailer(), nil
}

func (p *finalizerQueueProcessor) processDelayedQueue(ctx context.Context, items []*queueItem) error {
	now := time.Now()
	currentEpoch := p.finalizerContext.votingRoundTiming.EpochIndex(now)
	startTime := p.finalizerContext.votingRoundTiming.StartTime(currentEpoch)

	// best-effort dedup — the batch is already off the queue, a DB error must not drop it
	relayedItems, err := p.relayClient.ProtocolMessageRelayed(ctx, p.db, startTime, now)
	if err != nil {
		logger.Warnf("Finalizer delayed queue: proceeding without the already-relayed check: %v", err)
	}

	for _, item := range items {
		// skip only when the item's own target Relay already has it
		if data, exists := p.finalizationStorage.get(item.votingRoundID, item.protocolID, item.digest); exists {
			address := p.relayClient.addressForRewardEpoch(data.signingPolicy.RewardEpochID)
			if relayedItems.has(address, relayedKey{protocolID: item.protocolID, votingRoundID: item.votingRoundID}) {
				continue
			}
		}
		logger.Infof("Finalizer processes delayed queue item for round %v for protocol %v", item.votingRoundID, item.protocolID)
		p.processItem(ctx, item, true)
	}
	return nil
}
