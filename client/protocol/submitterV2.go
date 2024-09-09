package protocol

import (
	"bytes"
	"flare-fsc/client/shared"
	"flare-fsc/logger"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// WritePayloadV2 encodes payload to buffer.
// Payload data should be valid (data length 38, additional data length <= maxuint16 - 66)
func (s *SignatureSubmitter) WritePayloadV2(
	buffer *bytes.Buffer, currentEpoch int64, data *SubProtocolResponse, protocolID uint8,
) error {
	dataHash := accounts.TextHash(crypto.Keccak256(data.Data))
	signature, err := crypto.Sign(dataHash, s.protocolContext.signerPrivateKey)
	if err != nil {
		return errors.Wrap(err, "error signing submitSignatures data")
	}

	epochBytes := shared.Uint32toBytes(uint32(currentEpoch - 1))
	lengthBytes := shared.Uint16toBytes(uint16(66 + len(data.AdditionalData)))

	buffer.WriteByte(protocolID) // Protocol ID (1 byte)
	buffer.Write(epochBytes[:])  // Epoch (4 bytes)
	buffer.Write(lengthBytes[:]) // Length (2 bytes)
	buffer.WriteByte(0)          // Type (1 byte)

	buffer.WriteByte(signature[64] + 27) // V (1 byte)
	buffer.Write(signature[0:32])        // R (32 bytes)
	buffer.Write(signature[32:64])       // S (32 bytes)

	buffer.Write(data.AdditionalData)
	return nil
}

//  1. Run every sub-protocol provider with delay of 1 second at most five times.
//  2. repeat 1 for each sub-protocol provider not giving valid answer.
//
// Repeat 1 and 2 until all sub-protocol providers give valid answer or we did maxRounds attempts.
func (s *SignatureSubmitter) RunEpochV2(currentEpoch int64) {
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
			channels[i] = protocol.getDataWithRetry(
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
				logger.Error("Error getting data for submitter %s: %s", s.name, data.Message)
				continue
			}
			err := s.WritePayloadV2(buffer, currentEpoch, data.Value, s.subProtocols[i].Id)
			if err != nil {
				logger.Error("Error writing payload for submitter %s: %v", s.name, err)
				continue
			}

			// send message to finalizer
			s.messageChannel <- shared.ProtocolMessage{ProtocolID: s.subProtocols[i].Id, VotingRoundID: uint32(currentEpoch - 1), Message: data.Value.Data}

			protocolsSent = append(protocolsSent, i)
		}
		if len(protocolsSent) > 0 {
			if s.submit(buffer.Bytes()) {
				for _, i := range protocolsSent {
					delete(protocolsToSend, i)
				}
			}
		} else {
			logger.Info("Submitter %s did not get any new data", s.name)
		}
	}
}
