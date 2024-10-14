package protocol

import (
	"bytes"
	"crypto/ecdsa"
	"flare-fsc/client/config"
	"flare-fsc/client/shared"
	"flare-fsc/utils"
	"flare-fsc/utils/chain"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/logger"
)

type SubmitterBase struct {
	chainClient chain.Client
	gasConfig   *config.Gas

	protocolContext *protocolContext

	epoch        *utils.EpochConfig
	selector     []byte
	subProtocols []*SubProtocol

	startOffset      time.Duration
	submitRetries    int    // number of retries for submitting tx
	name             string // e.g., "submit1", "submit2", "submit3", "signatureSubmitter"
	submitPrivateKey *ecdsa.PrivateKey

	dataFetchRetries int           // number of retries for fetching data of each provider
	dataFetchTimeout time.Duration // timeout for fetching data of each provider
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
		err := s.chainClient.SendRawTx(s.submitPrivateKey, s.protocolContext.submitContractAddress, payload, s.gasConfig)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error sending submit tx for submitter %s tx", s.name))
		}
		return nil, nil
	}, s.submitRetries, shared.TxRetryInterval)
	if sendResult.Success {
		logger.Infof("Submitter %s successfully sent tx", s.name)
	}
	return sendResult.Success
}

func newSubmitter(
	ethClient *ethclient.Client,
	pc *protocolContext,
	epoch *utils.EpochConfig,
	submitCfg *config.Submit,
	gasCfg *config.Gas,
	selector []byte,
	subProtocols []*SubProtocol,
	epochOffset int64,
	name string,
) *Submitter {
	return &Submitter{
		SubmitterBase: SubmitterBase{
			chainClient:      chain.ClientImpl{EthClient: ethClient},
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
		channels[i] = protocol.fetchDataWithRetry(
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
			logger.Errorf("Error getting data for submitter %s: %s", s.name, data.Message)
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
	logger.Infof("Submitter %s running for epoch %d [%v, %v]", s.name, currentEpoch, s.epoch.StartTime(currentEpoch), s.epoch.EndTime(currentEpoch))

	payload, err := s.GetPayload(currentEpoch)

	if err != nil {
		logger.Errorf("Error getting payload for submitter %s: %v", s.name, err)
		return
	}
	if payload != nil {
		s.submit(payload)
	} else {
		logger.Infof("Submitter %s did not get any data, skipping submission", s.name)
	}
}

func newSignatureSubmitter(
	ethClient *ethclient.Client,
	pc *protocolContext,
	epoch *utils.EpochConfig,
	submitCfg *config.SubmitSignatures,
	gasCfg *config.Gas,
	selector []byte,
	subProtocols []*SubProtocol,
	messageChannel chan<- shared.ProtocolMessage,
) *SignatureSubmitter {
	return &SignatureSubmitter{
		SubmitterBase: SubmitterBase{
			chainClient:      chain.ClientImpl{EthClient: ethClient},
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

// WritePayload encodes payload to buffer.
// Payload data should be valid (data length 38, additional data length <= maxuint16 - 66)
func (s *SignatureSubmitter) WritePayload(
	buffer *bytes.Buffer, currentEpoch int64, data *SubProtocolResponse, protocolID, protocolType uint8,
) error {
	var dataLength int
	switch protocolType {
	case 0:
		dataLength = 104
	case 1:
		dataLength = 66
	default:
		return errors.New("unrecognized protocol type")
	}

	dataHash := accounts.TextHash(crypto.Keccak256(data.Data))
	signature, err := crypto.Sign(dataHash, s.protocolContext.signerPrivateKey)
	if err != nil {
		return errors.Wrap(err, "error signing submitSignatures data")
	}

	epochBytes := shared.Uint32toBytes(uint32(currentEpoch - 1))
	lengthBytes := shared.Uint16toBytes(uint16(dataLength + len(data.AdditionalData)))

	err = buffer.WriteByte(protocolID) // Protocol ID (1 byte)
	if err != nil {
		return errors.Wrap(err, "error writing Payload")
	}
	_, err = buffer.Write(epochBytes[:]) // Epoch (4 bytes)
	if err != nil {
		return errors.Wrap(err, "error writing Payload")
	}

	_, err = buffer.Write(lengthBytes[:]) // Length (2 bytes)
	if err != nil {
		return errors.Wrap(err, "error writing Payload")
	}

	err = buffer.WriteByte(protocolType) // Type (1 byte)
	if err != nil {
		return errors.Wrap(err, "error writing Payload")
	}
	if protocolType == 0 {
		n, err := buffer.Write(data.Data) // Message (38 bytes)
		if err != nil {
			return errors.Wrap(err, "error writing Payload")
		}
		if n != 38 {
			return errors.New("message not 38 bytes")
		}
	}

	if len(signature) != 65 {
		return errors.New("signature sanity check, this should not happen")
	}
	_, err = buffer.Write(utils.TransformSignatureRSVtoVRS(signature))
	if err != nil {
		return errors.Wrap(err, "error writing Payload")
	}

	_, err = buffer.Write(data.AdditionalData)
	if err != nil {
		return errors.Wrap(err, "error writing Payload")
	}

	return nil
}

// RunEpoch gets the submitSignature messages from the subprotocols providers.
//  1. run every sub-protocol provider with delay of 1 second at most five times.
//  2. repeat 1 for each sub-protocol provider not giving valid answer.
//
// Repeat 1 and 2 until all sub-protocol providers give valid answer or we did maxRounds attempts.
func (s *SignatureSubmitter) RunEpoch(currentEpoch int64) {
	logger.Info("Submitter %s running for epoch %d [%v, %v]", s.name, currentEpoch, s.epoch.StartTime(currentEpoch), s.epoch.EndTime(currentEpoch))

	protocolsToSend := make(map[int]bool)
	for i := range s.subProtocols {
		protocolsToSend[i] = true
	}
	channels := make([]<-chan shared.ExecuteStatus[*SubProtocolResponse], len(s.subProtocols))
	for i := 0; i < s.maxRounds && len(protocolsToSend) > 0; i++ {
		for i, protocol := range s.subProtocols {
			if !protocolsToSend[i] {
				continue
			}
			channels[i] = protocol.fetchDataWithRetry(
				currentEpoch-1,
				"submitSignatures",
				s.protocolContext.submitSignaturesAddress.Hex(),
				s.dataFetchRetries,
				s.dataFetchTimeout,
				SignatureSubmitterDataVerifier,
			)
		}

		protocolsSent := []int{}

		buffer := bytes.NewBuffer(nil)
		buffer.Write(s.selector)
		for i := range s.subProtocols {
			if !protocolsToSend[i] {
				continue
			}

			data := <-channels[i]
			if !data.Success {
				logger.Errorf("Error getting data for submitter %s: %s", s.name, data.Message)
				continue
			}
			err := s.WritePayload(buffer, currentEpoch, data.Value, s.subProtocols[i].ID, s.subProtocols[i].Type)
			if err != nil {
				logger.Errorf("Error writing payload for submitter %s: %v", s.name, err)
				continue
			}

			// send message to finalizer
			s.messageChannel <- shared.ProtocolMessage{ProtocolID: s.subProtocols[i].ID, VotingRoundID: uint32(currentEpoch - 1), Message: data.Value.Data}

			protocolsSent = append(protocolsSent, i)
		}
		if len(protocolsSent) > 0 {
			if s.submit(buffer.Bytes()) {
				for _, i := range protocolsSent {
					delete(protocolsToSend, i)
				}
			}
		} else {
			logger.Infof("Submitter %s did not get any new data", s.name)
		}
	}
}
