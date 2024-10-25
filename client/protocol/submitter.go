package protocol

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"math"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"
	"github.com/flare-foundation/flare-system-client/utils/chain"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

type SubmitterBase struct {
	chainClient chain.Client
	gasConfig   *config.Gas

	protocolContext *protocolContext

	votingRoundTiming *utils.EpochTimingConfig
	selector          []byte
	subProtocols      []*SubProtocol

	startOffset      time.Duration
	submitRetries    int           // number of retries for submitting tx
	submitTimeout    time.Duration // timeout for waiting for tx to be mined
	name             string        // e.g., "submit1", "submit2", "submit3", "signatureSubmitter"
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

	maxRounds int           // number of rounds for sending submitSignatures tx
	delay     time.Duration // minimal duration of one round
}

func (s *SubmitterBase) submit(payload []byte) bool {
	sendResult := <-shared.ExecuteWithRetryAttempts(func(ri int) (any, error) {
		gasConfig := gasConfigForAttempt(s.gasConfig, ri)
		logger.Debugf("[Attempt %d] Submitter %s sending tx with gas config: %+v, timeout: %s", ri, s.name, gasConfig, s.submitTimeout)
		err := s.chainClient.SendRawTx(s.submitPrivateKey, s.protocolContext.submitContractAddress, payload, gasConfig, s.submitTimeout)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error sending submit tx for submitter %s tx", s.name))
		}
		return nil, nil
	}, s.submitRetries, 1*time.Second)
	if sendResult.Success {
		logger.Infof("Submitter %s successfully sent tx", s.name)
	}
	return sendResult.Success
}

func newSubmitter(
	ethClient *ethclient.Client,
	pc *protocolContext,
	votingRoundTiming *utils.EpochTimingConfig,
	submitCfg *config.Submit,
	gasCfg *config.Gas,
	selector []byte,
	subProtocols []*SubProtocol,
	epochOffset int64,
	name string,
) *Submitter {
	return &Submitter{
		SubmitterBase: SubmitterBase{
			chainClient:       chain.ClientImpl{EthClient: ethClient},
			gasConfig:         gasCfg,
			protocolContext:   pc,
			votingRoundTiming: votingRoundTiming,
			selector:          selector,
			subProtocols:      subProtocols,
			startOffset:       submitCfg.StartOffset,
			submitRetries:     max(1, submitCfg.TxSubmitRetries),
			submitTimeout:     max(1*time.Second, submitCfg.TxSubmitTimeout),
			name:              name,
			submitPrivateKey:  pc.submitPrivateKey,
			dataFetchRetries:  submitCfg.DataFetchRetries,
			dataFetchTimeout:  submitCfg.DataFetchTimeout,
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
	logger.Infof("Submitter %s running for epoch %d [%v, %v]", s.name, currentEpoch, s.votingRoundTiming.StartTime(currentEpoch), s.votingRoundTiming.EndTime(currentEpoch))

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
	votingRoundTiming *utils.EpochTimingConfig,
	submitCfg *config.SubmitSignatures,
	gasCfg *config.Gas,
	selector []byte,
	subProtocols []*SubProtocol,
	messageChannel chan<- shared.ProtocolMessage,
) *SignatureSubmitter {

	delay := submitCfg.Delay

	if delay <= 0 {
		delay = time.Second
	}

	return &SignatureSubmitter{
		SubmitterBase: SubmitterBase{
			chainClient:       chain.ClientImpl{EthClient: ethClient},
			gasConfig:         gasCfg,
			protocolContext:   pc,
			votingRoundTiming: votingRoundTiming,
			startOffset:       submitCfg.StartOffset,
			selector:          selector,
			subProtocols:      subProtocols,
			submitRetries:     max(1, submitCfg.TxSubmitRetries),
			submitTimeout:     max(1*time.Second, submitCfg.TxSubmitTimeout),
			name:              "submitSignatures",
			submitPrivateKey:  pc.submitSignaturesPrivateKey,
			dataFetchTimeout:  submitCfg.DataFetchTimeout,
			dataFetchRetries:  submitCfg.DataFetchRetries,
		},
		maxRounds:      submitCfg.MaxRounds,
		delay:          delay,
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

	tempBuffer := bytes.NewBuffer(nil)

	tempBuffer.WriteByte(protocolID)   // Protocol ID (1 byte)
	tempBuffer.Write(epochBytes[:])    // Epoch (4 bytes)
	tempBuffer.Write(lengthBytes[:])   // Length (2 bytes)
	tempBuffer.WriteByte(protocolType) // Type (1 byte)

	if protocolType == 0 {
		n, _ := tempBuffer.Write(data.Data) // Message (38 bytes)
		if n != 38 {
			return errors.New("message not 38 bytes")
		}
	}

	if len(signature) != 65 {
		return errors.New("signature sanity check, this should not happen")
	}
	tempBuffer.Write(utils.TransformSignatureRSVtoVRS(signature))
	tempBuffer.Write(data.AdditionalData)

	buffer.Write(tempBuffer.Bytes())

	return nil
}

// RunEpoch gets the submitSignature messages from the subprotocols providers.
//  1. query every sub-protocol provider at most
//  2. repeat query for protocols that did not give a valid response at most maxRounds time
func (s *SignatureSubmitter) RunEpoch(currentEpoch int64) {
	logger.Infof("Submitter %s running for epoch %d [%v, %v]", s.name, currentEpoch, s.votingRoundTiming.StartTime(currentEpoch), s.votingRoundTiming.EndTime(currentEpoch))

	protocolsToSend := make(map[int]bool)
	for i := range s.subProtocols {
		protocolsToSend[i] = true
	}
	channels := make([]<-chan shared.ExecuteStatus[*SubProtocolResponse], len(s.subProtocols))

	ticker := time.NewTicker(s.delay)

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

		<-ticker.C
	}
}

// gasConfigForAttempt bumps up the gas price multiplier for each retry attempt by 50%,
// up to a maximum of 10x the original value.
//
// Note: If GasPriceFixed is used, the retry multiplier will not be applied.
func gasConfigForAttempt(cfg *config.Gas, ri int) *config.Gas {
	if cfg.GasPriceFixed.Cmp(common.Big0) != 0 {
		return cfg
	}

	retryMultiplier := min(10.0, math.Pow(1.5, float64(ri)))

	return &config.Gas{
		GasPriceMultiplier: max(1.0, cfg.GasPriceMultiplier) * float32(retryMultiplier),
		GasPriceFixed:      cfg.GasPriceFixed,
		GasLimit:           cfg.GasLimit,
	}
}
