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
	chainID int,
	ethClient *ethclient.Client,
	address common.Address,
	privateKey string,
) (*RelayContractClient, error) {
	txOpts, _, err := chain.CredentialsFromPrivateKey(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	relay, err := relay.NewRelay(address, ethClient)
	if err != nil {
		return nil, err
	}

	return &RelayContractClient{
		address:    address,
		relay:      relay,
		txOpts:     txOpts,
		txVerifier: chain.NewTxVerifier(ethClient),
	}, nil
}

func (r *RelayContractClient) SigningPolicyInitializedListener(db *gorm.DB, startTimestamp uint64, topic0 string) <-chan *relay.RelaySigningPolicyInitialized {
	out := make(chan *relay.RelaySigningPolicyInitialized)
	go func() {
		ticker := time.NewTicker(ListenerInterval)
		for {
			<-ticker.C
			now := time.Now().Unix()
			logs, err := database.FetchLogsByAddressAndTopic0(db, r.address.Hex(), topic0, int64(startTimestamp), now)
			if err != nil {
				logger.Error("Error fetching logs %w", err)
				continue
			}
			if len(logs) > 0 {
				policyData, err := r.parseSigningPolicyInitializedEvent(logs[len(logs)-1])
				if err != nil {
					logger.Error("Error parsing SigningPolicyInitialized event %w", err)
					continue
				}
				out <- policyData
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
