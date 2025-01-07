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

	// TEMP CHANGE for upgrading Relay contract, can be removed after 20 Jan 2025

	// If using new Coston2 Relay, query the old one as well.
	// Note: this won't have any effect on other networks as we currently have unique Relay addresses for each network.
	if r.address == common.HexToAddress("0x97702e350CaEda540935d92aAf213307e9069784") {
		logsOld, err := db.FetchLogsByAddressAndTopic0(common.HexToAddress("0x4087D4B5E009Af9FF41db910205439F82C3dc63c"), r.topic0SPI, from, to)
		if err != nil {
			return nil, err
		}
		allLogs = append(allLogs, logsOld...)
	}
	// END TEMP CHANGE

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

	execStatusChan := shared.ExecuteWithRetryChan(func() (string, error) {
		err := r.chainClient.SendRawTx(r.privateKey, r.address, input, r.gasConfig, chain.DefaultTxTimeout, dryRun)
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

	select {
	case execStatus := <-execStatusChan:
		if execStatus.Success {
			logger.Infof("Relaying finished for protocol %d with %s", protocolID, execStatus.Value)
		} else {
			logger.Warnf("Relaying failed with: %v", execStatus.Message)
		}

	case <-ctx.Done():
		return
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
