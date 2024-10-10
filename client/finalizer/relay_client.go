package finalizer

import (
	"context"
	"crypto/ecdsa"
	"flare-fsc/client/config"
	"flare-fsc/client/shared"
	"flare-fsc/utils/chain"
	"flare-fsc/utils/contracts/relay"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/logger"
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

	ethClient     relayEthClient
	relay         *relay.Relay
	privateKey    *ecdsa.PrivateKey
	senderAddress common.Address

	relaySelector []byte      // for relay method
	topic0SPI     common.Hash // for SigningPolicyInitialized event
	topic0PMR     common.Hash // for ProtocolMessageRelayed event
}

type relayEthClient interface {
	SendRawTx(*ecdsa.PrivateKey, common.Address, []byte, bool) error
}

type relayEthClientImpl struct {
	client *ethclient.Client
}

func (eth relayEthClientImpl) SendRawTx(privateKey *ecdsa.PrivateKey, to common.Address, data []byte, dryRun bool) error {
	return chain.SendRawType2Tx(eth.client, privateKey, to, data, dryRun, &config.GasConfig{GasPriceFixed: common.Big0})
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
		ethClient:     relayEthClientImpl{client: ethClient},
		address:       address,
		relay:         relayContract,
		privateKey:    privateKey,
		senderAddress: senderAddress,
		relaySelector: relaySelectorBytes,
		topic0SPI:     topic0SPI,
		topic0PMR:     topic0PMR,
	}, nil
}

func (r *relayContractClient) FetchSigningPolicies(db finalizerDB, from, to int64) ([]signingPolicyListenerResponse, error) {
	logs, err := db.FetchLogsByAddressAndTopic0(r.address, r.topic0SPI, from, to)
	if err != nil {
		return nil, err
	}

	result := make([]signingPolicyListenerResponse, 0, len(logs))
	for _, log := range logs {
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

// SubmitPayloads sends a transaction with input relay contract.
func (r *relayContractClient) SubmitPayloads(ctx context.Context, input []byte, dryRun bool) {
	if len(input) == 0 {
		return
	}

	execStatusChan := shared.ExecuteWithRetry(func() (any, error) {
		err := r.ethClient.SendRawTx(r.privateKey, r.address, input, dryRun)
		if err != nil {
			if shared.ExistsAsSubstring(nonFatalRelayErrors, err.Error()) {
				logger.Infof("Non fatal error sending relay tx: %v", err)
			} else {
				return nil, errors.Wrap(err, "Error sending relay tx")
			}
		}
		return nil, nil
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)

	select {
	case execStatus := <-execStatusChan:
		if execStatus.Success {
			logger.Info("Relaying finished")
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
