package finalizer

import (
	"encoding/binary"
	"fmt"
)

var (
	errPayloadTooShort = fmt.Errorf("invalid payload length: too short")
)

// Todo: fix names in the file

type payloadMessage struct {
	protocolId byte
	epochId    uint32
	payload    []byte
}

type signedPayload struct {
	typeId         byte
	message        []byte
	signature      []byte
	additionalData []byte
}

type submittedPayload struct {
	protocolId         byte
	votingRoundId      uint32
	randomQualityScore bool
	merkleRoot         []byte
}

func DecodeSignatureSubmitterPayload(message []byte) ([]*payloadMessage, error) {
	var messages []*payloadMessage
	for i := 0; i < len(message); {
		if len(message)-i < 7 {
			return nil, fmt.Errorf("invalid payload length at index %d of %d", i, len(message))
		}
		protocolId := message[i]
		i += 1
		epochId := binary.BigEndian.Uint32(message[i : i+4])
		i += 4
		payloadLength := int(binary.BigEndian.Uint16(message[i : i+2]))
		i += 2
		if len(message)-i < payloadLength {
			return nil, errPayloadTooShort
		}
		messages = append(messages, &payloadMessage{
			protocolId: protocolId,
			epochId:    epochId,
			payload:    message[i : i+payloadLength],
		})
		i += payloadLength
	}
	return messages, nil
}

func decodeSignedPayload(payload []byte) (*signedPayload, error) {
	if len(payload) < 104 { // 104 = 1 + 38 + 65
		return nil, errPayloadTooShort
	}
	reponse := &signedPayload{
		typeId:    payload[0],
		message:   payload[1:39],
		signature: payload[39:104],
	}
	if len(payload) > 104 {
		reponse.additionalData = payload[104:]
	}
	return reponse, nil
}

func decodeSubmittedPayload(payload []byte) (*submittedPayload, error) {
	if len(payload) < 38 { // 38 = 1 + 4 + 1 + 32
		return nil, errPayloadTooShort
	}
	rqs := payload[5]
	if rqs != 0 && rqs != 1 {
		return nil, fmt.Errorf("invalid random quality score value: %d", rqs)
	}
	return &submittedPayload{
		protocolId:         payload[0],
		votingRoundId:      binary.BigEndian.Uint32(payload[1:5]),
		randomQualityScore: rqs == 1,
		merkleRoot:         payload[6:38],
	}, nil
}
