package finalizer

import (
	"encoding/binary"
	"errors"
	"flare-fsc/client/shared/voters"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type payloadMessage struct {
	protocolID    uint8
	votingRoundID uint32
	payload       []byte
}

func ExtractPayloads(data []byte) ([]payloadMessage, error) {
	messages := []payloadMessage{}

	data = data[4:] // trim function selector
	for len(data) > 0 {
		if len(data) < 7 { // 7 = 1 + 4 + 2
			return nil, errors.New("wrongly formatted tx input, too short")
		}

		protocol := data[0] // 1 byte protocol ID

		votingRound := binary.BigEndian.Uint32(data[1:5]) // 4 bytes votingRoundID

		length := binary.BigEndian.Uint16(data[5:7]) // 2 bytes length of payload in bytes

		end := 7 + length

		if len(data) < int(end) {
			return nil, errors.New("wrongly formatted tx input")
		}

		payload := data[7:end]

		message := payloadMessage{
			protocolID:    protocol,
			votingRoundID: votingRound,
			payload:       payload,
		}

		messages = append(messages, message)

		data = data[end:] // trim the extracted payload
	}
	return messages, nil
}

type submitSignaturesPayload struct {
	protocolID    uint8
	votingRoundID uint32
	typeID        uint8
	signature     []byte

	signer     common.Address
	voterIndex int
	weight     uint16
}

func decodeSignedPayloadV2(payloadMsg payloadMessage) (submitSignaturesPayload, error) {
	typeID := payloadMsg.payload[0]

	var signatureStart, signatureEnd int

	switch typeID {
	case 0:
		signatureStart = 1 + 38
		signatureEnd = signatureStart + 1 + 2*32

	case 1:
		signatureStart = 1
		signatureEnd = signatureStart + 1 + 2*32
	default:
		return submitSignaturesPayload{}, fmt.Errorf("invalid typeID %d", typeID)
	}

	if len(payloadMsg.payload) < signatureEnd {
		return submitSignaturesPayload{}, fmt.Errorf("payload to short got %d, should be at least %d", len(payloadMsg.payload), signatureEnd)
	}

	signature := payloadMsg.payload[signatureStart:signatureEnd]

	signedPayload := submitSignaturesPayload{
		protocolID:    payloadMsg.protocolID,
		votingRoundID: payloadMsg.votingRoundID,
		typeID:        typeID,
		signature:     signature,

		voterIndex: -1, // 0 is a valid index, we use -1 before assigning the proper value
	}

	return signedPayload, nil
}

func (pld *submitSignaturesPayload) AddSigner(messageHash []byte, voterSet *voters.VoterSet) error {
	transformedSignature := transformSignature(pld.signature)

	pk, err := crypto.SigToPub(messageHash, transformedSignature[:])
	if err != nil {
		return err
	}

	signer := crypto.PubkeyToAddress(*pk)

	pld.signer = signer

	pld.voterIndex = voterSet.VoterIndex(signer)

	if pld.voterIndex < 0 {
		return fmt.Errorf("signer %s is not a registered voter in the current reward epoch", signer.Hex())
	}

	pld.weight = voterSet.VoterWeight(pld.voterIndex)

	return nil
}
