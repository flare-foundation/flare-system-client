package finalizer

import (
	"flare-tlc/client/shared"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/relay"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

const (
	listenerBufferSize = 10
)

type relayContractClient struct {
	address    common.Address
	relay      *relay.Relay
	txOpts     *bind.TransactOpts
	txVerifier *chain.TxVerifier
}

type signingPolicyListenerResponse struct {
	policyData *relay.RelaySigningPolicyInitialized
	timestamp  int64
}

func NewRelayContractClient(
	ethClient *ethclient.Client,
	address common.Address,
	txOpts *bind.TransactOpts,
) (*relayContractClient, error) {
	relay, err := relay.NewRelay(address, ethClient)
	if err != nil {
		return nil, err
	}

	return &relayContractClient{
		address:    address,
		relay:      relay,
		txOpts:     txOpts,
		txVerifier: chain.NewTxVerifier(ethClient),
	}, nil
}

func (r *relayContractClient) SigningPolicyInitializedListener(db *gorm.DB, startTime time.Time) <-chan signingPolicyListenerResponse {
	out := make(chan signingPolicyListenerResponse, listenerBufferSize)
	topic0, err := chain.EventIDFromMetadata(relay.RelayMetaData, "SigningPolicyInitialized")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	go func() {
		ticker := time.NewTicker(shared.ListenerInterval)
		eventRangeStart := startTime.Unix()
		for {
			<-ticker.C
			now := time.Now().Unix()
			logs, err := database.FetchLogsByAddressAndTopic0(db, r.address.Hex(), topic0, eventRangeStart, now)
			if err != nil {
				logger.Error("Error fetching logs %v", err)
				continue
			}
			for _, log := range logs {
				policyData, err := shared.ParseSigningPolicyInitializedEvent(r.relay, log)
				if err != nil {
					logger.Error("Error parsing SigningPolicyInitialized event %v", err)
					break
				}
				out <- signingPolicyListenerResponse{policyData, int64(log.Timestamp)}
				// continue with timestamps > log.Timestamp,
				// there should be only one such log per timestamp
				eventRangeStart = int64(log.Timestamp)
			}
		}
	}()
	return out
}

func (r *relayContractClient) Submit() {

}
