package finalizer

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"flare-tlc/client/config"
	"flare-tlc/client/shared"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/relay"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
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

	relaySelector []byte // for relay method
	topic0SPI     string // for SigningPolicyInitialized event
	topic0PMR     string // for ProtocolMessageRelayed event
}

type relayEthClient interface {
	SendRawTx(*ecdsa.PrivateKey, common.Address, []byte, bool) error
}

type relayEthClientImpl struct {
	client *ethclient.Client
}

func (eth relayEthClientImpl) SendRawTx(privateKey *ecdsa.PrivateKey, to common.Address, data []byte, dryRun bool) error {
	return chain.SendRawTx(eth.client, privateKey, to, data, dryRun, &config.GasConfig{GasPriceFixed: common.Big0}, chain.DefaultTxTimeout)
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
	var allLogs []database.Log

	// TEMP CHANGE for upgrading Relay contract, should be removed after 17 Oct 2024

	// If using new Songbird Relay, query the old one as well.
	// Note: this won't have any effect on other networks as we currently have unique Relay addresses for each network.
	if r.address == common.HexToAddress("0x0D462d2Fec11554D64F52D7c5A5C269d748037aD") {
		logsOld, err := db.FetchLogsByAddressAndTopic0(common.HexToAddress("0xbA35e39D01A3f5710d1e43FC61dbb738B68641c4"), r.topic0SPI, from, to)
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
			logger.Error("Error parsing SigningPolicyInitialized event %v", err)
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

func (r *relayContractClient) SubmitPayloads(ctx context.Context, payloads []*signedPayload, signingPolicy *signingPolicy, dryRun bool) {
	if len(payloads) == 0 || signingPolicy == nil {
		return
	}

	buffer := bytes.NewBuffer(nil)
	buffer.Write(r.relaySelector)
	buffer.Write(signingPolicy.rawBytes)
	buffer.Write(payloads[0].rawMessage)
	signatureBytes, err := EncodeForRelay(payloads)
	if err != nil {
		logger.Error("Error encoding payloads %v", err)
		return
	}
	buffer.Write(signatureBytes)
	payload := buffer.Bytes()

	execStatusChan := shared.ExecuteWithRetry(func() (any, error) {
		err := r.ethClient.SendRawTx(r.privateKey, r.address, payload, dryRun)
		if err != nil {
			if shared.ExistsAsSubstring(nonFatalRelayErrors, err.Error()) {
				logger.Info("Non fatal error sending relay tx: %v", err)
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

func (r *relayContractClient) ProtocolMessageRelayed(db finalizerDB, from time.Time, to time.Time) (mapset.Set[queueItem], error) {
	logs, err := db.FetchLogsByAddressAndTopic0(r.address, r.topic0PMR, from.Unix(), to.Unix())
	if err != nil {
		return nil, err
	}

	result := mapset.NewSet[queueItem]()
	for _, log := range logs {
		data, err := shared.ParseProtocolMessageRelayedEvent(r.relay, log)
		if err != nil {
			return nil, err
		}
		result.Add(queueItem{
			protocolId:    data.ProtocolId,
			votingRoundId: data.VotingRoundId,
			messageHash:   data.MerkleRoot,
		})
	}
	return result, nil
}
