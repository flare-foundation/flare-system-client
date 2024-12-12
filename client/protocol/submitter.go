package protocol

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"
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
	"github.com/flare-foundation/go-flare-common/pkg/payload"
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

	deadline time.Duration // deadline for sending submitSignatures tx before end of the grace period

	// config for sending after the deadline
	maxCycles     int           // number of cycles for sending submitSignature transaction after the deadline
	cycleDuration time.Duration // minimal duration of one cycle
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

// GetPayload
func (s *Submitter) GetPayload(currentEpoch int64) []byte {
	channels := make([]<-chan shared.ExecuteStatus[*SubProtocolResponse], len(s.subProtocols))
	for i, protocol := range s.subProtocols {
		channels[i] = protocol.fetchDataWithRetryChan(
			currentEpoch+s.epochOffset,
			s.name,
			s.protocolContext.submitAddress.Hex(),
			s.dataFetchRetries,
			s.dataFetchTimeout,
			StatusDataVerifier,
		)
	}

	buffer := bytes.NewBuffer(nil)
	buffer.Write(s.selector)

	dataReceived := false
	for j, channel := range channels {
		data := <-channel
		if !data.Success {
			logger.Warnf("Error getting data for submitter %s for protocol %v: %s", s.name, s.subProtocols[j].ID, data.Message)
			continue
		} else if data.Value.Status == payload.Empty {
			logger.Debugf("Empty data for submitter %s for protocol %v", s.name, s.subProtocols[j].ID)
		}

		logger.Debugf("%s received data for round %d for protocol %d", s.name, currentEpoch+s.epochOffset, s.subProtocols[j].ID)
		dataReceived = true
		buffer.Write(data.Value.Data)
	}

	if !dataReceived {
		return nil
	}

	return buffer.Bytes()
}

func (s *Submitter) RunEpoch(currentEpoch int64) {
	logger.Debugf("Submitter %s running for epoch %d")

	payload := s.GetPayload(currentEpoch)

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

	deadline := submitCfg.Deadline - submitCfg.StartOffset
	if deadline <= 0 {
		deadline = 0
	}

	delay := submitCfg.CycleDuration
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
		maxCycles:      submitCfg.MaxCycles,
		cycleDuration:  delay,
		messageChannel: messageChannel,
		deadline:       deadline,
	}
}

// WritePayload encodes payload to buffer.
// Payload data should be valid (data length 38, additional data length <= maxuint16 - 66).
// If an error is returned, the buffer is unchanged.
func (s *SignatureSubmitter) WritePayload(
	buffer *bytes.Buffer, epoch int64, data *SubProtocolResponse, protocolID, protocolType uint8,
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

	epochBytes := shared.Uint32toBytes(uint32(epoch))
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

// RunEpoch gets the submitSignature messages from the subprotocols providers, relays them to the finalizer, aggregated them and submits them to submission contract.
//
// The first cycle of queries happens before the deadline:
// Each subprotocol is repeatedly queried until success or deadline.
// The successful queries are included in the payloads of the transaction.
// The transaction with payloads is sent once all queries are successful or at the deadline (whichever happens first).
//
// After the first cycle, if not all subprotocols are queried successfully, there is another cycle.
// In this cycle, the not yet successful queries are retried at most dataFetchRetries times.
// If any subprotocols are queried successfully, their payload is submitted at the end of the cycle.
// The processes is repeated until all subprotocols are queried successfully or at most maxCycles of times.
func (s *SignatureSubmitter) RunEpoch(currentEpoch int64) {
	logger.Infof("%s fetching data for %d", s.name, currentEpoch-1)
	protocolsToRetry, err := s.RunEpochBeforeDeadline(currentEpoch, s.deadline)

	if err != nil {
		logger.Errorf("error before Deadline %v", err)
	}

	if len(protocolsToRetry) > 0 {
		logger.Debugf("running %s for %v after deadline for protocols", s.name, currentEpoch-1)
		s.RunEpochAfterDeadline(currentEpoch, protocolsToRetry)
	}
}

// RunEpochBeforeDeadline queries subprotocol providers for signed messages until success or deadline.
// Messages are written to transaction input in form of signature payload.
// Messages are also passed to finalizers.
// Once all subprotocols are successfully queried or the deadline has passed the transaction with constructed input is sent to the submission contract.
// A set of indexes of unsuccessfully queried subprotocols is returned.
func (s *SignatureSubmitter) RunEpochBeforeDeadline(currentEpoch int64, deadline time.Duration) (map[int]bool, error) {
	protocolsToQuery := make(map[int]bool)

	results := make([]*SubProtocolResponse, len(s.subProtocols))

	finished := make(chan int, len(s.subProtocols))

	ctx, cancel := context.WithTimeout(context.Background(), deadline)

	for i, protocol := range s.subProtocols {
		protocolsToQuery[i] = true

		go func() {
			response := protocol.fetchDataWithRetry(
				ctx,
				currentEpoch-1,
				"submitSignatures",
				s.protocolContext.submitSignaturesAddress.Hex(),
				s.dataFetchTimeout,
				SignatureSubmitterDataVerifier,
				time.Second, // TODO make it configurable
			)

			if response.Success {
				results[i] = response.Value
				finished <- i
			} else {
				logger.Debugf("unsuccessful data for round %d for protocol %d: %v", currentEpoch-1, protocol.ID, response.Message)
			}
		}()

	}
	buffer := bytes.NewBuffer(nil)
	buffer.Write(s.selector)

	protocolsDone := 0
	readyToSend := false

	// in a case when both finished and ctx.Done() are ready we first process finished
	for {
		select {
		case i := <-finished:
			switch results[i].Status {
			case payload.Ok:
				err := s.WritePayload(buffer, currentEpoch-1, results[i], s.subProtocols[i].ID, s.subProtocols[i].Type)
				if err != nil {
					logger.Errorf("Error writing payload for submitter %s: %v", s.name, err)
				} else {
					logger.Debugf("%s received data for round %d for protocol %d", s.name, currentEpoch-1, s.subProtocols[i].ID)
					s.messageChannel <- shared.ProtocolMessage{ProtocolID: s.subProtocols[i].ID, VotingRoundID: uint32(currentEpoch - 1), Message: results[i].Data}
				}
			case payload.Empty:
				logger.Warnf("Empty payload for for submitter %s for round %d for protocol %d", s.name, currentEpoch-1, s.subProtocols[i].ID)
			default:
				logger.Warnf("Unknown success result status %s for submitter %s for round %d for protocol %d", results[i].Status, s.name, currentEpoch-1, s.subProtocols[i].ID)
			}
			delete(protocolsToQuery, i)
			protocolsDone++
			if protocolsDone == len(s.subProtocols) {
				readyToSend = true
				logger.Debugf("Payloads for submitter %s for round %v collected before deadline", s.name, currentEpoch-1)
			}
		default:
			select {
			case <-ctx.Done():
				logger.Debugf("Tx for submitter %s for round %v triggered by the deadline", s.name, currentEpoch-1)
				readyToSend = true
			case i := <-finished:
				finished <- i
			}

		}
		if readyToSend {
			cancel()

			if protocolsDone > 0 {
				txSent := s.submit(buffer.Bytes())
				if !txSent {
					for i := range s.subProtocols {
						protocolsToQuery[i] = true
					}
					return protocolsToQuery, errors.Errorf("submitSignatures tx failed")
				}
			}
			return protocolsToQuery, nil
		}
	}
}

func (s *SignatureSubmitter) RunEpochAfterDeadline(currentEpoch int64, protocolsToQuery map[int]bool) {
	channels := make([]<-chan shared.ExecuteStatus[*SubProtocolResponse], len(s.subProtocols))

	ticker := time.NewTicker(s.cycleDuration)

	for j := 0; j < s.maxCycles && len(protocolsToQuery) > 0; j++ {
		for i, protocol := range s.subProtocols {
			if !protocolsToQuery[i] {
				continue
			}
			channels[i] = protocol.fetchDataWithRetryChan(
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
			if !protocolsToQuery[i] {
				continue
			}

			data := <-channels[i]
			if !data.Success {
				logger.Warnf("Error getting data for submitter %s: %s", s.name, data.Message)
				continue
			}

			switch data.Value.Status {
			case payload.Ok:
				err := s.WritePayload(buffer, currentEpoch-1, data.Value, s.subProtocols[i].ID, s.subProtocols[i].Type)
				if err != nil {
					logger.Errorf("Error writing payload for submitter %s: %v", s.name, err)
				} else {
					// send message to finalizer
					s.messageChannel <- shared.ProtocolMessage{ProtocolID: s.subProtocols[i].ID, VotingRoundID: uint32(currentEpoch - 1), Message: data.Value.Data}
				}
			case payload.Empty:
				logger.Warnf("Empty submit signature payload for submitter %s for protocol %v ", s.name, s.subProtocols[i].ID)
			default:
				logger.Errorf("Unknown success status %s for submitter %s for protocol %v ", data.Value.Status, s.name, s.subProtocols[i].ID)
			}

			protocolsSent = append(protocolsSent, i)
		}
		if len(protocolsSent) > 0 && buffer.Len() > 0 {
			if s.submit(buffer.Bytes()) {
				for _, i := range protocolsSent {
					delete(protocolsToQuery, i)
				}
			}
		} else {
			logger.Debugf("Submitter %s did not get any new data", s.name)
		}

		<-ticker.C
	}
}

// gasConfigForAttempt sets gas config for a retry attempt.
//
// For type 0 transaction, it bumps up GasPriceMultiplier for each retry attempt by 50%,
// up to a maximum of 10x the original value.
// If GasPriceFixed is used, the retry multiplier will not be applied.
//
// For type 2 transaction, MaxPriorityFeePerGas on the n-the attempt is n times the MaxPriorityFeePerGas of the initial attempt.
func gasConfigForAttempt(cfg *config.Gas, ri int) *config.Gas {
	if cfg.TxType == 0 {
		if cfg.GasPriceFixed.Cmp(common.Big0) != 0 {
			return cfg
		}

		retryMultiplier := min(10.0, math.Pow(1.5, float64(ri)))

		return &config.Gas{
			TxType:   0,
			GasLimit: cfg.GasLimit,

			GasPriceMultiplier: max(1.0, cfg.GasPriceMultiplier) * float32(retryMultiplier),
			GasPriceFixed:      cfg.GasPriceFixed,
		}
	} else if cfg.TxType == 2 {
		tipCap := new(big.Int)
		if cfg.MaxPriorityFeePerGas != nil && cfg.MaxPriorityFeePerGas.Cmp(big.NewInt(0)) == 1 {
			tipCap.Set(cfg.MaxPriorityFeePerGas)
		} else {
			tipCap.Set(chain.DefaultTipCap)
		}

		retryMultiplier := int64(1 + ri)
		tipCap = tipCap.Mul(tipCap, big.NewInt(retryMultiplier))

		return &config.Gas{
			TxType:   2,
			GasLimit: cfg.GasLimit,

			MaxPriorityFeePerGas: tipCap,
			BaseFeePerGasCap:     cfg.BaseFeePerGasCap,
		}
	} else {
		return cfg
	}
}
