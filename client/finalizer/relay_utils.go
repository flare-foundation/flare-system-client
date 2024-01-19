package finalizer

import (
	"bytes"
	"crypto/ecdsa"
	"flare-tlc/client/shared"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/relay"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	listenerBufferSize = 10
)

type relayContractClient struct {
	address common.Address

	ethClient  *ethclient.Client
	relay      *relay.Relay
	privateKey *ecdsa.PrivateKey

	relaySelector []byte
}

type signingPolicyListenerResponse struct {
	policyData *relay.RelaySigningPolicyInitialized
	timestamp  int64
}

func NewRelayContractClient(
	ethClient *ethclient.Client,
	address common.Address,
	privateKey *ecdsa.PrivateKey,
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

	return &relayContractClient{
		address:       address,
		relay:         relayContract,
		privateKey:    privateKey,
		relaySelector: relaySelectorBytes,
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

func (r *relayContractClient) SubmitPayloads(payloads []*signedPayload, signingPolicy *signingPolicy) {
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

	execStatus := <-shared.ExecuteWithRetry(func() (any, error) {
		err := chain.SendRawTx(r.ethClient, r.privateKey, r.address, payload)
		if err != nil {
			return nil, errors.Wrap(err, "Error sending relay tx")
		}
		return nil, nil
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)

	// TODO: what happens if not successful
	if execStatus.Success {
		logger.Info("Relay tx sent")
	}
}
