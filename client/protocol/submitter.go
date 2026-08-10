package protocol

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"
	"github.com/flare-foundation/flare-system-client/utils/chain"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

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
	retryDelay       time.Duration // backoff between send attempts; tests shrink it
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

// submit signs and broadcasts a tx with input to submitContractAddress, retrying
// with pre-/post-broadcast awareness:
//   - pre-broadcast failures (tx never sent) refresh the nonce and retry;
//   - post-broadcast timeouts keep the nonce and bump gas, so the retry replaces
//     the pending tx rather than duplicating it;
//   - "nonce too low" is reconciled against the hashes broadcast so far: if one
//     was accepted the send succeeded, otherwise the nonce is bumped and resent.
func (s *SubmitterBase) submit(ctx context.Context, input []byte) bool {
	if len(input) <= 4 {
		return false
	}

	nonceResult := <-shared.ExecuteWithRetryChan(ctx, func() (uint64, error) { return s.chainClient.Nonce(ctx, s.submitPrivateKey, 2*time.Second) }, 3, 100*time.Millisecond)
	if !nonceResult.Success {
		logger.Errorf("Submitter %s getting nonce: %v", s.name, nonceResult.Message)
		return false
	}
	nonce := nonceResult.Value

	from := crypto.PubkeyToAddress(s.submitPrivateKey.PublicKey)
	var broadcastHashes []common.Hash
	undetermined := false // an own broadcast's fate was never resolved

	sendResult := <-shared.ExecuteWithRetryAttempts(ctx, func(ri int) (string, error) {
		gasConfig := chain.GasConfigForAttempt(s.gasConfig, ri)
		logger.Debugf("[Attempt %d] Submitter %s sending tx with nonce %d, gas config: %+v, timeout: %s", ri, s.name, nonce, gasConfig, s.submitTimeout)

		res := s.chainClient.SendRawTx(ctx, s.submitPrivateKey, nonce, s.protocolContext.submitContractAddress, input, gasConfig, s.submitTimeout, true)
		if res.Broadcast {
			broadcastHashes = append(broadcastHashes, res.Hash)
		}

		switch {
		case res.Err == nil:
			return res.Hash.Hex(), nil // mined successfully
		case chain.IsNonceTooLow(res.Err):
			// Nonce is consumed by some mined tx. If it was one of ours the payload
			// is submitted; if we cannot tell, do NOT resend (would duplicate);
			// only bump once we know none of ours landed (a foreign tx took it).
			// A tight lookup timeout: the submitter is vote-window sensitive, so it
			// fails fast to an Undetermined retry rather than blocking on a slow RPC.
			h, acc := chain.AnyAccepted(ctx, s.chainClient, from, broadcastHashes, nil, time.Second)
			switch acc {
			case chain.Accepted:
				logger.Infof("Submitter %s: broadcast tx %s accepted, nonce too low is non-fatal", s.name, h.Hex())
				return h.Hex(), nil
			case chain.Undetermined:
				// Outcome unknown: refresh the nonce and resend (duplicate submits are idempotent).
				undetermined = true
				logger.Warnf("Submitter %s: nonce %d too low, prior tx status unknown; resending at a refreshed nonce", s.name, nonce)
				nonce = s.refreshNonce(ctx, nonce)
				return "", res.Err
			default: // NonceConsumed
				logger.Warnf("Submitter %s: nonce %d consumed by another tx, bumping nonce", s.name, nonce)
				nonce = s.refreshNonce(ctx, nonce)
				return "", res.Err
			}
		case res.Broadcast && chain.IsTimeout(res.Err):
			// Post-broadcast timeout: tx may be in the mempool at this nonce. Keep
			// the nonce and bump gas next attempt so it replaces the pending tx.
			logger.Warnf("Submitter %s: timed out awaiting confirmation of tx %s (nonce %d), retrying as replacement", s.name, res.Hash.Hex(), nonce)
			return "", res.Err
		default:
			// Not confirmed. If we have already broadcast a tx it is outstanding at
			// this nonce, so keep it (a retry replaces it); only refresh the nonce
			// when nothing has been broadcast yet.
			if len(broadcastHashes) == 0 {
				nonce = s.refreshNonce(ctx, nonce)
			}
			logger.Warnf("Submitter %s: send failed at nonce %d: %v", s.name, nonce, res.Err)
			return "", res.Err
		}
	}, s.submitRetries, s.retryDelay)

	if sendResult.Success {
		logger.Infof("Submitter %s successfully sent tx %s", s.name, sendResult.Value)
	} else if undetermined {
		logger.Warnf("Submitter %s: submission outcome unknown, a broadcast tx may be on chain: %s", s.name, sendResult.Message)
	} else {
		logger.Errorf("Submitter %s unsuccessful tx: %s", s.name, sendResult.Message)
	}
	return sendResult.Success
}

// refreshNonce best-effort re-fetches the account nonce, keeping current on error.
func (s *SubmitterBase) refreshNonce(ctx context.Context, current uint64) uint64 {
	nonce, err := s.chainClient.Nonce(ctx, s.submitPrivateKey, time.Second)
	if err != nil {
		logger.Warnf("Submitter %s failed to refresh nonce: %v", s.name, err)
		return current
	}
	return nonce
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
			submitTimeout:     max(2*time.Second, submitCfg.TxSubmitTimeout),
			retryDelay:        time.Second,
			name:              name,
			submitPrivateKey:  pc.submitPrivateKey,
			dataFetchRetries:  submitCfg.DataFetchRetries,
			dataFetchTimeout:  submitCfg.DataFetchTimeout,
		},
		epochOffset: epochOffset,
	}
}

// GetPayload
func (s *Submitter) GetPayload(ctx context.Context, currentEpoch int64) []byte {
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
		var data shared.ExecuteStatus[*SubProtocolResponse]
		select {
		case <-ctx.Done():
			return nil
		case data = <-channel:
		}
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

func (s *Submitter) RunEpoch(ctx context.Context, currentEpoch int64) {
	logger.Debugf("Submitter %s running for epoch %d", s.name, currentEpoch)

	payload := s.GetPayload(ctx, currentEpoch)

	if payload != nil {
		s.submit(ctx, payload)
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
			submitTimeout:     max(2*time.Second, submitCfg.TxSubmitTimeout),
			retryDelay:        time.Second,
			name:              "submitSignatures",
			submitPrivateKey:  pc.submitSignaturesPrivateKey,
			dataFetchTimeout:  submitCfg.DataFetchTimeout,
			dataFetchRetries:  submitCfg.DataFetchRetries,
		},
		maxCycles:      submitCfg.MaxCycles,
		cycleDuration:  delay,
		messageChannel: messageChannel,
		deadline:       max(submitCfg.Deadline-submitCfg.StartOffset, 0),
	}
}

// WritePayload encodes payload to buffer using the submitter's signing key.
// Payload data should be valid (data length 38, additional data length <= maxuint16 - 66).
// If an error is returned, the buffer is unchanged.
func (s *SignatureSubmitter) WritePayload(
	buffer *bytes.Buffer, epoch int64, data *SubProtocolResponse, protocolID, protocolType uint8,
) error {
	return EncodePayload(buffer, epoch, data, protocolID, protocolType, s.protocolContext.signerPrivateKey)
}

// EncodePayload encodes a signed submitSignatures payload to buffer.
// Payload data should be valid (data length 38, additional data length <= maxuint16 - 66).
// If an error is returned, the buffer is unchanged.
func EncodePayload(
	buffer *bytes.Buffer, epoch int64, data *SubProtocolResponse, protocolID, protocolType uint8,
	signerPrivateKey *ecdsa.PrivateKey,
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
	signature, err := crypto.Sign(dataHash, signerPrivateKey)
	if err != nil {
		return fmt.Errorf("signing submitSignatures data: %w", err)
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

	vrsSignature, err := utils.TransformSignatureRSVtoVRS(signature)
	if err != nil {
		return fmt.Errorf("signature sanity check, this should not happen: %w", err)
	}
	tempBuffer.Write(vrsSignature)
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
func (s *SignatureSubmitter) RunEpoch(ctx context.Context, round int64) {
	logger.Infof("%s fetching data for %d", s.name, round)
	protocolsToRetry, err := s.RunEpochBeforeDeadline(ctx, round, s.deadline)
	if err != nil {
		logger.Errorf("before Deadline %v", err)
	}

	if len(protocolsToRetry) > 0 {
		logger.Debugf("running %s for %v after deadline for protocols", s.name, round)
		s.RunEpochAfterDeadline(ctx, round, protocolsToRetry)
	}
}

// RunEpochBeforeDeadline queries subprotocol providers for signed messages until success or deadline.
// Messages are written to transaction input in form of signature payload.
// Messages are also passed to finalizers.
// Once all subprotocols are successfully queried or the deadline has passed the transaction with constructed input is sent to the submission contract.
// A set of indexes of unsuccessfully queried subprotocols is returned.
func (s *SignatureSubmitter) RunEpochBeforeDeadline(ctx context.Context, round int64, deadline time.Duration) (map[int]bool, error) {
	protocolsToQuery := make(map[int]bool)

	results := make([]*SubProtocolResponse, len(s.subProtocols))

	finished := make(chan int, len(s.subProtocols))

	deadlineCtx, cancel := context.WithTimeout(ctx, deadline)

	for i, protocol := range s.subProtocols {
		protocolsToQuery[i] = true

		go func() {
			response := protocol.fetchDataWithRetry(
				deadlineCtx,
				round,
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
				logger.Debugf("unsuccessful data for round %d for protocol %d: %v", round, protocol.ID, response.Message)
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
				err := s.WritePayload(buffer, round, results[i], s.subProtocols[i].ID, s.subProtocols[i].Type)
				if err != nil {
					logger.Errorf("Error writing payload for submitter %s for round %d for protocol %d : %s", s.name, round, s.subProtocols[i].ID, err)
				} else {
					logger.Debugf("%s received data for round %d for protocol %d", s.name, round, s.subProtocols[i].ID)
					if s.messageChannel != nil {
						select {
						case s.messageChannel <- shared.ProtocolMessage{
							ProtocolID:    s.subProtocols[i].ID,
							VotingRoundID: uint32(round),
							Message:       results[i].Data,
						}:
						default:
							logger.Warnf("message channel full. Dropping message round %d for protocol %d after deadline", round, s.subProtocols[i].ID)
						}
					}
				}
			case payload.Empty:
				logger.Warnf("Empty payload for for submitter %s for round %d for protocol %d", s.name, round, s.subProtocols[i].ID)
			default:
				logger.Warnf("Unknown success result status %s for submitter %s for round %d for protocol %d", results[i].Status, s.name, round, s.subProtocols[i].ID)
			}
			delete(protocolsToQuery, i)
			protocolsDone++
			if protocolsDone == len(s.subProtocols) {
				readyToSend = true
				logger.Debugf("Payloads for submitter %s for round %v collected before deadline", s.name, round)
			}
		default:
			select {
			case <-deadlineCtx.Done():
				logger.Debugf("Tx for submitter %s for round %v triggered by the deadline", s.name, round)
				readyToSend = true
			case i := <-finished:
				finished <- i
			}
		}

		if readyToSend {
			cancel()

			if protocolsDone > 0 {
				txSent := s.submit(ctx, buffer.Bytes())
				if !txSent {
					for i := range s.subProtocols {
						protocolsToQuery[i] = true
					}
					return protocolsToQuery, errors.New("submitSignatures tx failed")
				}
			}
			return protocolsToQuery, nil
		}
	}
}

func (s *SignatureSubmitter) RunEpochAfterDeadline(ctx context.Context, round int64, protocolsToQuery map[int]bool) {
	channels := make([]<-chan shared.ExecuteStatus[*SubProtocolResponse], len(s.subProtocols))

	ticker := time.NewTicker(s.cycleDuration)
	defer ticker.Stop()

	for j := 0; j < s.maxCycles && len(protocolsToQuery) > 0; j++ {
		for i, protocol := range s.subProtocols {
			if !protocolsToQuery[i] {
				continue
			}
			channels[i] = protocol.fetchDataWithRetryChan(
				round,
				"submitSignatures",
				s.protocolContext.submitSignaturesAddress.Hex(),
				s.dataFetchRetries,
				s.dataFetchTimeout,
				SignatureSubmitterDataVerifier,
			)
		}

		protocolsSent := make([]int, 0, len(protocolsToQuery))

		buffer := bytes.NewBuffer(nil)
		buffer.Write(s.selector)
		for i := range s.subProtocols {
			if !protocolsToQuery[i] {
				continue
			}

			select {
			case <-ctx.Done():
				return
			case data := <-channels[i]:
				if !data.Success {
					logger.Warnf("Error getting data for submitter %s: %s", s.name, data.Message)
					continue
				}

				switch data.Value.Status {
				case payload.Ok:
					err := s.WritePayload(buffer, round, data.Value, s.subProtocols[i].ID, s.subProtocols[i].Type)
					if err != nil {
						logger.Errorf("Error writing payload for submitter %s for round %d for protocol %d : %s", s.name, round, s.subProtocols[i].ID, err)
					} else {
						logger.Debugf("%s received data for round %d for protocol %d after deadline", s.name, round, s.subProtocols[i].ID)
						if s.messageChannel != nil {
							select {
							case s.messageChannel <- shared.ProtocolMessage{
								ProtocolID:    s.subProtocols[i].ID,
								VotingRoundID: uint32(round),
								Message:       data.Value.Data,
							}:
							default:
								logger.Warnf("message channel full. Dropping message round %d for protocol %d after deadline", round, s.subProtocols[i].ID)
							}
						}
					}
				case payload.Empty:
					logger.Warnf("Empty submit signature payload for submitter %s for protocol %v ", s.name, s.subProtocols[i].ID)
				default:
					logger.Errorf("Unknown success status %s for submitter %s for protocol %v ", data.Value.Status, s.name, s.subProtocols[i].ID)
				}

				protocolsSent = append(protocolsSent, i)
			}
		}
		if len(protocolsSent) > 0 && buffer.Len() > 0 {
			if s.submit(ctx, buffer.Bytes()) {
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
