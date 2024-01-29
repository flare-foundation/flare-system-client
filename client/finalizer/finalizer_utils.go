package finalizer

import (
	"bytes"
	"encoding/binary"
	"flare-tlc/client/shared"
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	errPayloadTooShort = fmt.Errorf("invalid payload length: too short")
)

type submitterPayloadItem struct {
	protocolId    byte
	votingRoundId uint32
	payload       *signedPayload
}

type signedPayload struct {
	typeId         byte
	message        *submittedPayload
	rawMessage     []byte
	signature      []byte
	additionalData []byte

	// calculated from the fields above
	signer      common.Address
	messageHash common.Hash

	// index of voter in signing policy, updated when inserting it into storage
	index int
}

type submittedPayload struct {
	protocolId         byte
	votingRoundId      uint32
	randomQualityScore bool
	merkleRoot         []byte
}

func DecodeSubmitterPayload(message []byte) ([]*submitterPayloadItem, error) {
	if len(message) == 0 {
		return nil, nil
	}
	var messages []*submitterPayloadItem
	for i := 4; i < len(message); {
		if len(message)-i < 7 {
			return nil, fmt.Errorf("invalid payload length at index %d of %d", i, len(message))
		}
		protocolId := message[i]
		i += 1
		votingRoundId := binary.BigEndian.Uint32(message[i : i+4])
		i += 4
		payloadLength := int(binary.BigEndian.Uint16(message[i : i+2]))
		i += 2
		if len(message)-i < payloadLength {
			return nil, errPayloadTooShort
		}
		payload, err := decodeSignedPayload(message[i : i+payloadLength])
		if err != nil {
			return nil, err
		}
		messages = append(messages, &submitterPayloadItem{
			protocolId:    protocolId,
			votingRoundId: votingRoundId,
			payload:       payload,
		})
		i += payloadLength
	}
	return messages, nil
}

// Transform signature to be used by go-ethereum crypto.SigToPub:
// transforms [V || R || S] to [R || S || V - 27]
// No checks are performed, we assume that signature array has length 65
func transformSignature(signature []byte) (RSV [65]byte) {
	copy(RSV[:], signature[1:33])
	copy(RSV[32:], signature[33:65])
	RSV[64] = signature[0] - 27
	return RSV
}

func decodeSignedPayload(payload []byte) (*signedPayload, error) {
	if len(payload) < 104 { // 104 = 1 + 38 + 65
		return nil, errPayloadTooShort
	}
	rawMessage := payload[1:39]
	message, err := decodeSubmittedPayload(rawMessage)
	if err != nil {
		return nil, err
	}
	signature := payload[39:104]

	messageHash := accounts.TextHash(crypto.Keccak256(rawMessage))
	transformedSignature := transformSignature(signature)
	pk, err := crypto.SigToPub(messageHash, transformedSignature[:])
	if err != nil {
		return nil, err
	}
	signer := crypto.PubkeyToAddress(*pk)
	reponse := &signedPayload{
		typeId:     payload[0],
		message:    message,
		rawMessage: rawMessage,
		signature:  signature,

		messageHash: common.BytesToHash(messageHash),
		signer:      signer,

		index: -1,
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

func EncodeForRelay(payloads []*signedPayload) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if len(payloads) > math.MaxUint16 {
		return nil, fmt.Errorf("too many payloads: %d", len(payloads))
	}

	sizeBytes := shared.Uint16toBytes(uint16(len(payloads)))
	buffer.Write(sizeBytes[:])
	prevIndex := -1
	for _, payload := range payloads {
		if payload.index < 0 {
			return nil, fmt.Errorf("payload index not set")
		}
		if prevIndex >= payload.index {
			return nil, fmt.Errorf("payloads not sorted by index")
		}

		indexBytes := shared.Uint16toBytes(uint16(payload.index))
		buffer.Write(payload.signature)
		buffer.Write(indexBytes[:])
		prevIndex = payload.index
	}
	return buffer.Bytes(), nil
}
