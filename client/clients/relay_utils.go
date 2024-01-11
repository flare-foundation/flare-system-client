package clients

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

type RelayContractClient struct {
	address    common.Address
	relay      *relay.Relay
	txOpts     *bind.TransactOpts
	txVerifier *chain.TxVerifier
}

func NewRelayContractClient(
	ethClient *ethclient.Client,
	address common.Address,
) (*RelayContractClient, error) {
	relay, err := relay.NewRelay(address, ethClient)
	if err != nil {
		return nil, err
	}

	return &RelayContractClient{
		address:    address,
		relay:      relay,
		txVerifier: chain.NewTxVerifier(ethClient),
	}, nil
}

func (r *RelayContractClient) SigningPolicyInitializedListener(db *gorm.DB, startTimestamp uint64) <-chan *relay.RelaySigningPolicyInitialized {
	topic0, err := chain.EventIDFromMetadata(relay.RelayMetaData, "SigningPolicyInitialized")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	out := make(chan *relay.RelaySigningPolicyInitialized)
	go func() {
		ticker := time.NewTicker(ListenerInterval)
		for {
			<-ticker.C
			now := time.Now().Unix()
			logs, err := database.FetchLogsByAddressAndTopic0(db, r.address.Hex(), topic0, int64(startTimestamp), now)
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

func (r *RelayContractClient) parseSigningPolicyInitializedEvent(dbLog database.Log) (*relay.RelaySigningPolicyInitialized, error) {
	contractLog, err := shared.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, err
	}
	return r.relay.RelayFilterer.ParseSigningPolicyInitialized(*contractLog)
}
