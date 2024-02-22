package registration

import (
	"flare-tlc/client/shared"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/relay"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type relayContractClient interface {
	SigningPolicyInitializedListener(registrationClientDB, uint64) <-chan *relay.RelaySigningPolicyInitialized
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

func (r *relayContractClientImpl) SigningPolicyInitializedListener(db registrationClientDB, startTimestamp uint64) <-chan *relay.RelaySigningPolicyInitialized {
	topic0, err := chain.EventIDFromMetadata(relay.RelayMetaData, "SigningPolicyInitialized")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	out := make(chan *relay.RelaySigningPolicyInitialized)
	go func() {
		ticker := time.NewTicker(shared.ListenerInterval)
		for {
			<-ticker.C
			now := time.Now().Unix()
			logs, err := db.FetchLogsByAddressAndTopic0(r.address, topic0, int64(startTimestamp), now)
			if err != nil {
				logger.Error("Error fetching logs %v", err)
				continue
			}
			if len(logs) > 0 {
				policyData, err := r.parseSigningPolicyInitializedEvent(logs[len(logs)-1])
				if err != nil {
					logger.Error("Error parsing SigningPolicyInitialized event %v", err)
					continue
				}
				out <- policyData
				break
			}
		}
	}()
	return out
}

func (r *relayContractClientImpl) parseSigningPolicyInitializedEvent(dbLog database.Log) (*relay.RelaySigningPolicyInitialized, error) {
	return shared.ParseSigningPolicyInitializedEvent(r.relay, dbLog)
}
