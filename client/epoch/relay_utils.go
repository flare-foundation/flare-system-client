package epoch

import (
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
	SigningPolicyInitializedListener(epochClientDB, *utils.EpochTimingConfig) <-chan *relay.RelaySigningPolicyInitialized
}

type relayContractClientImpl struct {
	address    common.Address
	relay      *relay.Relay
	txVerifier *chain.TxVerifier
}

func NewRelayContractClient(
	ethClient *ethclient.Client,
	address common.Address,
) (*relayContractClientImpl, error) {
	relay, err := relay.NewRelay(address, ethClient)
	if err != nil {
		return nil, err
	}
	return &relayContractClientImpl{
		address:    address,
		relay:      relay,
		txVerifier: chain.NewTxVerifier(ethClient),
	}, nil
}

func (r *relayContractClientImpl) SigningPolicyInitializedListener(db epochClientDB, rewardEpochTiming *utils.EpochTimingConfig) <-chan *relay.RelaySigningPolicyInitialized {
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
			<-ticker.C
			now := time.Now().Unix()
			logs, err := db.FetchLogsByAddressAndTopic0Timestamp(r.address, topic0, eventRangeStart, now)
			if err != nil {
				logger.Errorf("Error fetching logs %v", err)
				continue
			}
			if len(logs) > 0 {
				policyData, err := r.parseSigningPolicyInitializedEvent(logs[len(logs)-1])
				if err != nil {
					logger.Errorf("Error parsing SigningPolicyInitialized event %v", err)
					continue
				}
				out <- policyData
				eventRangeStart = int64(policyData.Timestamp)
			}
		}
	}()
	return out
}

func (r *relayContractClientImpl) parseSigningPolicyInitializedEvent(dbLog database.Log) (*relay.RelaySigningPolicyInitialized, error) {
	return shared.ParseSigningPolicyInitializedEvent(r.relay, dbLog)
}
