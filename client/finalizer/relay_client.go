package finalizer

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils/chain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/flare-foundation/go-flare-common/pkg/database"
	"github.com/flare-foundation/go-flare-common/pkg/logger"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
)

const (
	listenerBufferSize = 10
)

var (
	nonFatalRelayErrors = []string{
		"Already relayed",
		"nonce too low",
	}
)

type relayContractClient struct {
	address common.Address

	chainClient chain.Client
	gasConfig   *config.Gas

	relay         *relay.Relay
	privateKey    *ecdsa.PrivateKey
	senderAddress common.Address

	relaySelector []byte      // for relay method
	topic0SPI     common.Hash // for SigningPolicyInitialized event
	topic0PMR     common.Hash // for ProtocolMessageRelayed event
}

type signingPolicyListenerResponse struct {
	policyData *relay.RelaySigningPolicyInitialized
	timestamp  int64
}

func NewRelayContractClient(
	ethClient *ethclient.Client,
	address common.Address,
	privateKey *ecdsa.PrivateKey,
	senderAddress common.Address,
	gasConfig *config.Gas,
) (*relayContractClient, error) {
	relayContract, err := relay.NewRelay(address, ethClient)
	if err != nil {
		return nil, err
	}

	relayABI, err := relay.RelayMetaData.GetAbi()
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	relaySelectorBytes := relayABI.Methods["relay"].ID
	topic0SPI, err := chain.EventIDFromMetadata(relay.RelayMetaData, "SigningPolicyInitialized")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	topic0PMR, err := chain.EventIDFromMetadata(relay.RelayMetaData, "ProtocolMessageRelayed")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}

	return &relayContractClient{
		chainClient:   chain.ClientImpl{EthClient: ethClient},
		address:       address,
		relay:         relayContract,
		privateKey:    privateKey,
		senderAddress: senderAddress,
		relaySelector: relaySelectorBytes,
		topic0SPI:     topic0SPI,
		topic0PMR:     topic0PMR,
		gasConfig:     gasConfig,
	}, nil
}

func (r *relayContractClient) FetchSigningPolicies(db finalizerDB, from, to int64) ([]signingPolicyListenerResponse, error) {
	var allLogs []database.Log

	logs, err := db.FetchLogsByAddressAndTopic0(r.address, r.topic0SPI, from, to)
	if err != nil {
		return nil, err
	}

	allLogs = append(allLogs, logs...)

	result := make([]signingPolicyListenerResponse, 0, len(allLogs))
	for _, log := range allLogs {
		policyData, err := shared.ParseSigningPolicyInitializedEvent(r.relay, log)
		if err != nil {
			logger.Errorf("Error parsing SigningPolicyInitialized event %v", err)
			return nil, err
		}
		result = append(result, signingPolicyListenerResponse{policyData, int64(log.Timestamp)})
	}
	return result, nil
}

func (r *relayContractClient) SigningPolicyInitializedListener(db finalizerDB, startTime time.Time) <-chan signingPolicyListenerResponse {
	out := make(chan signingPolicyListenerResponse, listenerBufferSize)
	go func() {
		ticker := time.NewTicker(shared.EventListenerInterval)
		eventRangeStart := startTime.Unix()
		for {
			<-ticker.C
			now := time.Now().Unix()
			logs, err := db.FetchLogsByAddressAndTopic0(r.address, r.topic0SPI, eventRangeStart, now)
			if err != nil {
				logger.Errorf("Error fetching logs %v", err)
				continue
			}
			for _, log := range logs {
				policyData, err := shared.ParseSigningPolicyInitializedEvent(r.relay, log)
				if err != nil {
					logger.Errorf("Error parsing SigningPolicyInitialized event %v", err)
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

// SubmitPayloads sends a transaction with input to Relay contract.
func (r *relayContractClient) SubmitPayloads(ctx context.Context, input []byte, dryRun bool, protocolID uint8) {
	if len(input) == 0 {
		return
	}

	nonce, err := r.chainClient.Nonce(r.privateKey, 2*time.Second)
	if err != nil {
		logger.Error("error getting nonce: %v", err)
		return
	}

	sendResult := <-shared.ExecuteWithRetryAttempts(func(ri int) (string, error) {
		gasConfig := chain.GasConfigForAttempt(r.gasConfig, ri)

		err := r.chainClient.SendRawTx(r.privateKey, nonce, r.address, input, gasConfig, chain.DefaultTxTimeout, dryRun)
		if err != nil {
			if shared.ExistsAsSubstring(nonFatalRelayErrors, err.Error()) {
				logger.Debugf("Non fatal error sending relay tx for protocol %d: %v", protocolID, err)
				return "non fatal error", nil
			} else {
				return "", errors.Wrap(err, "Error sending relay tx")
			}
		}
		return "success", nil
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)

	if sendResult.Success {
		logger.Infof("Relaying finished for protocol %d with %s", protocolID, sendResult.Value)
	} else {
		logger.Warnf("Relaying failed with: %v", sendResult.Message)
	}

}

// ProtocolMessageRelayed returns a set of pairs of protocol and round that have been finalized.
func (r *relayContractClient) ProtocolMessageRelayed(db finalizerDB, from time.Time, to time.Time) (map[queueItem]bool, error) {
	logs, err := db.FetchLogsByAddressAndTopic0(r.address, r.topic0PMR, from.Unix(), to.Unix())
	if err != nil {
		return nil, err
	}

	result := make(map[queueItem]bool)
	for _, log := range logs {
		data, err := shared.ParseProtocolMessageRelayedEvent(r.relay, log)
		if err != nil {
			return nil, err
		}
		result[queueItem{
			protocolID:    data.ProtocolId,
			votingRoundID: data.VotingRoundId,
		}] = true
	}
	return result, nil
}
