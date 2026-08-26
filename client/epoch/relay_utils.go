package epoch

import (
	"context"
	"time"

	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"
	"github.com/flare-foundation/flare-system-client/utils/chain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/flare-foundation/go-flare-common/pkg/database"
	"github.com/flare-foundation/go-flare-common/pkg/logger"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
)

type relayContractClient interface {
	SigningPolicyInitializedListener(ctx context.Context, db epochClientDB, config *utils.EpochTimingConfig) <-chan *relay.RelaySigningPolicyInitialized
}

type relayContractClientImpl struct {
	addresses    []common.Address // configured Relay and, across the cutover, the new one
	relay        *relay.Relay
	txVerifier   *chain.TxVerifier
	relayCutover *shared.RelayCutover
}

func NewRelayContractClient(
	ethClient *ethclient.Client,
	address common.Address,
	relayCutover *shared.RelayCutover,
) (*relayContractClientImpl, error) {
	relay, err := relay.NewRelay(address, ethClient)
	if err != nil {
		return nil, err
	}

	addresses := []common.Address{address}
	if relayCutover.Scheduled() {
		addresses = append(addresses, relayCutover.NewAddress)
	}

	return &relayContractClientImpl{
		addresses:    addresses,
		relay:        relay,
		txVerifier:   chain.NewTxVerifier(ethClient),
		relayCutover: relayCutover,
	}, nil
}

func (r *relayContractClientImpl) SigningPolicyInitializedListener(ctx context.Context, db epochClientDB, rewardEpochTiming *utils.EpochTimingConfig) <-chan *relay.RelaySigningPolicyInitialized {
	topic0, err := chain.EventIDFromMetadata(relay.RelayMetaData, "SigningPolicyInitialized")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	out := make(chan *relay.RelaySigningPolicyInitialized)

	go func() {
		randomDelay()
		ticker := time.NewTicker(shared.EventListenerInterval)
		eventRangeStart := rewardEpochTiming.StartTime(rewardEpochTiming.EpochIndex(time.Now()) - 1).Unix()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}
			now := time.Now().Unix()

			var logs []database.Log
			failed := false
			for _, address := range r.addresses {
				addressLogs, err := db.FetchLogsByAddressAndTopic0Timestamp(ctx, address, topic0, eventRangeStart, now)
				if err != nil {
					logger.Errorf("Error fetching logs %v", err)
					failed = true
					break
				}
				logs = append(logs, addressLogs...)
			}
			if failed {
				continue
			}

			// Only the anticipated epoch's policy is worth signing; across the switch either
			// contract may have emitted it, so it is picked by reward epoch, not log position.
			anticipated := rewardEpochTiming.EpochIndex(time.Now()) + 1

			policyData := r.selectPolicy(logs, anticipated)
			if policyData != nil {
				out <- policyData
				eventRangeStart = int64(policyData.Timestamp)
			}
		}
	}()
	return out
}

// selectPolicy returns the anticipated epoch's policy, else the highest present — a delayed
// epoch start leaves the anticipated index running ahead of the chain. Every policy seen is
// offered to the Relay cutover, which is how it learns the voting round the switch starts on.
func (r *relayContractClientImpl) selectPolicy(logs []database.Log, anticipated int64) *relay.RelaySigningPolicyInitialized {
	var selected *relay.RelaySigningPolicyInitialized

	for _, log := range logs {
		policyData, err := r.parseSigningPolicyInitializedEvent(log)
		if err != nil {
			logger.Errorf("Error parsing SigningPolicyInitialized event %v", err)
			continue
		}
		epoch := policyData.RewardEpochId.Int64()
		r.relayCutover.ObserveSigningPolicy(epoch, policyData.StartVotingRoundId)

		if epoch == anticipated {
			return policyData
		}
		if selected == nil || epoch > selected.RewardEpochId.Int64() {
			selected = policyData
		}
	}

	if selected != nil && selected.RewardEpochId.Int64() != anticipated {
		logger.Debugf("Expected signing policy for reward epoch %d, taking the highest available %v",
			anticipated, selected.RewardEpochId)
	}
	return selected
}

func (r *relayContractClientImpl) parseSigningPolicyInitializedEvent(dbLog database.Log) (*relay.RelaySigningPolicyInitialized, error) {
	return shared.ParseSigningPolicyInitializedEvent(r.relay, dbLog)
}
