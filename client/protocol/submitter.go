package protocol

import (
	"bytes"
	"crypto/ecdsa"
	"flare-fsc/client/config"
	"flare-fsc/client/shared"
	"flare-fsc/logger"
	"flare-fsc/utils"
	"flare-fsc/utils/chain"
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type SubmitterBase struct {
	ethClient submitterEthClient
	gasConfig *config.GasConfig

	protocolContext *protocolContext

	epoch        *utils.Epoch
	selector     []byte
	subProtocols []*SubProtocol

	startOffset      time.Duration
	submitRetries    int    // number of retries for submitting tx
	name             string // e.g., "submit1", "submit2", "submit3", "signatureSubmitter"
	submitPrivateKey *ecdsa.PrivateKey

	dataFetchRetries int           // number of retries for fetching data of each provider
	dataFetchTimeout time.Duration // timeout for fetching data of each provider
}

type submitterEthClient interface {
	SendRawTx(*ecdsa.PrivateKey, common.Address, []byte, *config.GasConfig) error
}

type submitterEthClientImpl struct {
	ethClient *ethclient.Client
}

func (c submitterEthClientImpl) SendRawTx(privateKey *ecdsa.PrivateKey, to common.Address, payload []byte, gasConfig *config.GasConfig) error {
	return chain.SendRawTx(c.ethClient, privateKey, to, payload, true, gasConfig)
}

type Submitter struct {
	SubmitterBase

	epochOffset int64 // offset from current epoch, e.g., -1, 0
}

type SignatureSubmitter struct {
	SubmitterBase

	messageChannel chan<- shared.ProtocolMessage

	maxRounds int // number of rounds for sending submitSignatures tx
}

func (s *SubmitterBase) submit(payload []byte) bool {
	sendResult := <-shared.ExecuteWithRetry(func() (any, error) {
		err := s.ethClient.SendRawTx(s.submitPrivateKey, s.protocolContext.submitContractAddress, payload, s.gasConfig)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error sending submit tx for submitter %s tx", s.name))
		}
		return nil, nil
	}, s.submitRetries, shared.TxRetryInterval)
	if sendResult.Success {
		logger.Info("Submitter %s successfully sent tx", s.name)
	}
	return sendResult.Success
}

func newSubmitter(
	ethClient *ethclient.Client,
	pc *protocolContext,
	epoch *utils.Epoch,
	submitCfg *config.SubmitConfig,
	gasCfg *config.GasConfig,
	selector []byte,
	subProtocols []*SubProtocol,
	epochOffset int64,
	name string,
) *Submitter {
	return &Submitter{
		SubmitterBase: SubmitterBase{
			ethClient:        submitterEthClientImpl{ethClient: ethClient},
			gasConfig:        gasCfg,
			protocolContext:  pc,
			epoch:            epoch,
			selector:         selector,
			subProtocols:     subProtocols,
			startOffset:      submitCfg.StartOffset,
			submitRetries:    max(1, submitCfg.TxSubmitRetries),
			name:             name,
			submitPrivateKey: pc.submitPrivateKey,
			dataFetchRetries: submitCfg.DataFetchRetries,
			dataFetchTimeout: submitCfg.DataFetchTimeout,
		},
		epochOffset: epochOffset,
	}
}

func (s *Submitter) GetPayload(currentEpoch int64) ([]byte, error) {
	channels := make([]<-chan shared.ExecuteStatus[*SubProtocolResponse], len(s.subProtocols))
	for i, protocol := range s.subProtocols {
		channels[i] = protocol.getDataWithRetry(
			currentEpoch+s.epochOffset,
			s.name,
			s.protocolContext.submitAddress.Hex(),
			s.dataFetchRetries,
			s.dataFetchTimeout,
			IdentityDataVerifier,
		)
	}

	buffer := bytes.NewBuffer(nil)
	buffer.Write(s.selector)

	dataReceived := false
	for _, channel := range channels {
		data := <-channel
		if !data.Success || data.Value.Status != "OK" {
			logger.Error("Error getting data for submitter %s: %s", s.name, data.Message)
			continue
		}
		dataReceived = true
		buffer.Write(data.Value.Data)
	}

	if !dataReceived {
		return nil, nil
	}

	return buffer.Bytes(), nil
}

func (s *Submitter) RunEpoch(currentEpoch int64) {
	logger.Info("Submitter %s running for epoch %d [%v, %v]", s.name, currentEpoch, s.epoch.StartTime(currentEpoch), s.epoch.EndTime(currentEpoch))

	payload, err := s.GetPayload(currentEpoch)

	if err != nil {
		logger.Error("Error getting payload for submitter %s: %v", s.name, err)
		return
	}
	if payload != nil {
		s.submit(payload)
	} else {
		logger.Info("Submitter %s did not get any data, skipping submission", s.name)
	}
}

func newSignatureSubmitter(
	ethClient *ethclient.Client,
	pc *protocolContext,
	epoch *utils.Epoch,
	submitCfg *config.SubmitSignaturesConfig,
	gasCfg *config.GasConfig,
	selector []byte,
	subProtocols []*SubProtocol,
	messageChannel chan<- shared.ProtocolMessage,
) *SignatureSubmitter {
	return &SignatureSubmitter{
		SubmitterBase: SubmitterBase{
			ethClient:        submitterEthClientImpl{ethClient: ethClient},
			gasConfig:        gasCfg,
			protocolContext:  pc,
			epoch:            epoch,
			startOffset:      submitCfg.StartOffset,
			selector:         selector,
			subProtocols:     subProtocols,
			submitRetries:    max(1, submitCfg.TxSubmitRetries),
			name:             "submitSignatures",
			submitPrivateKey: pc.submitSignaturesPrivateKey,
			dataFetchTimeout: submitCfg.DataFetchTimeout,
			dataFetchRetries: submitCfg.DataFetchRetries,
		},
		maxRounds:      submitCfg.MaxRounds,
		messageChannel: messageChannel,
	}
}

// Payload data should be valid (data length 38, additional data length <= maxuint16 - 104)
func (s *SignatureSubmitter) WritePayload(
	buffer *bytes.Buffer, currentEpoch int64, data *SubProtocolResponse, protocolID uint8,
) error {
	dataHash := accounts.TextHash(crypto.Keccak256(data.Data))
	signature, err := crypto.Sign(dataHash, s.protocolContext.signerPrivateKey)
	if err != nil {
		return errors.Wrap(err, "error signing submitSignatures data")
	}

	epochBytes := shared.Uint32toBytes(uint32(currentEpoch - 1))
	lengthBytes := shared.Uint16toBytes(uint16(104 + len(data.AdditionalData)))

	buffer.WriteByte(protocolID) // Protocol ID (1 byte)
	buffer.Write(epochBytes[:])  // Epoch (4 bytes)
	buffer.Write(lengthBytes[:]) // Length (2 bytes)

	buffer.WriteByte(0)     // Type (1 byte)
	buffer.Write(data.Data) // Message (38 bytes)

	buffer.WriteByte(signature[64] + 27) // V (1 byte)
	buffer.Write(signature[0:32])        // R (32 bytes)
	buffer.Write(signature[32:64])       // S (32 bytes)

	buffer.Write(data.AdditionalData)
	return nil
}

//  1. Run every sub-protocol provider with delay of 1 second at most five times.
//  2. repeat 1 for each sub-protocol provider not giving valid answer.
//
// Repeat 1 and 2 until all sub-protocol providers give valid answer or we did 10 rounds.
func (s *SignatureSubmitter) RunEpoch(currentEpoch int64) {
	logger.Info("Submitter %s running for epoch %d [%v, %v]", s.name, currentEpoch, s.epoch.StartTime(currentEpoch), s.epoch.EndTime(currentEpoch))

	protocolsToSend := mapset.NewSet[int]()
	for i := range s.subProtocols {
		protocolsToSend.Add(i)
	}
	channels := make([]<-chan shared.ExecuteStatus[*SubProtocolResponse], len(s.subProtocols))
	for i := 0; i < s.maxRounds && protocolsToSend.Cardinality() > 0; i++ {
		for i, protocol := range s.subProtocols {
			if !protocolsToSend.Contains(i) {
				continue
			}
			channels[i] = protocol.getDataWithRetry(
				currentEpoch-1,
				"submitSignatures",
				s.protocolContext.submitSignaturesAddress.Hex(),
				s.dataFetchRetries,
				s.dataFetchTimeout,
				SignatureSubmitterDataVerifier,
			)
		}

		protocolsToSendCopy := protocolsToSend.Clone() // copy in case of submit failure

		buffer := bytes.NewBuffer(nil)
		buffer.Write(s.selector)
		for i := range s.subProtocols {
			if !protocolsToSend.Contains(i) {
				continue
			}

			data := <-channels[i]
			if !data.Success {
				logger.Error("Error getting data for submitter %s: %s", s.name, data.Message)
				continue
			}
			err := s.WritePayload(buffer, currentEpoch, data.Value, s.subProtocols[i].ID)

			s.messageChannel <- shared.ProtocolMessage{ProtocolID: s.subProtocols[i].ID, VotingRoundID: uint32(currentEpoch - 1), Message: data.Value.Data}

			if err != nil {
				logger.Error("Error writing payload for submitter %s: %v", s.name, err)
				continue
			}
			protocolsToSend.Remove(i)
		}
		if protocolsToSendCopy.Cardinality() > protocolsToSend.Cardinality() {
			if !s.submit(buffer.Bytes()) {
				protocolsToSend = protocolsToSendCopy
			}
		} else {
			logger.Info("Submitter %s did not get any new data", s.name)
		}
	}
}
