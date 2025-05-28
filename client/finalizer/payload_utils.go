package finalizer

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

// payloadMessage is a general structure that is used in the submit calls to the chain.
type payloadMessage struct {
	protocolID    uint8
	votingRoundID uint32
	payload       []byte
}

// ExtractPayloads extracts payloads from a transaction input to submission contracts and returns a slice of payloadMessages.
func ExtractPayloads(data []byte) ([]payloadMessage, error) {
	messages := []payloadMessage{}

	data = data[4:] // trim function selector
	for len(data) > 0 {
		if len(data) < 7 { // 7 = 1 + 4 + 2
			return nil, errors.New("wrongly formatted tx input, too short")
		}

		protocol := data[0]                               // 1 byte protocol ID
		votingRound := binary.BigEndian.Uint32(data[1:5]) // 4 bytes votingRoundID
		length := binary.BigEndian.Uint16(data[5:7])      // 2 bytes length of payload in bytes
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

// submitSignaturesPayload is a specialized structure used in submitSignatures calls to the chain.
type submitSignaturesPayload struct {
	protocolID    uint8
	votingRoundID uint32
	typeID        uint8
	signature     []byte

	//assigned after processing
	signer     common.Address
	voterIndex int
	weight     uint16

	message shared.Message // only if type 0
}

func (s *submitSignaturesPayload) FromSignedPayload(payloadMsg payloadMessage) error {
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
		return fmt.Errorf("invalid typeID %d", typeID)
	}

	if len(payloadMsg.payload) < signatureEnd {
		return fmt.Errorf("payload of type %d to short got %d, should be at least %d", typeID, len(payloadMsg.payload), signatureEnd)
	}

	signature := payloadMsg.payload[signatureStart:signatureEnd]

	s.protocolID = payloadMsg.protocolID
	s.votingRoundID = payloadMsg.votingRoundID
	s.typeID = typeID
	s.signature = signature
	s.voterIndex = -1 // 0 is a valid index, we use -1 before assigning the proper value

	if typeID == 0 {
		s.message = payloadMsg.payload[1:39]
	}

	return nil
}

// AddSigner calculates the public key of the signer from the signature and messageHash and adds its voterIndex and weight (if the signer is in the votingSet) to the submitSignaturesPayload.
func (pld *submitSignaturesPayload) AddSigner(messageHash []byte, voterSet *voters.Set) error {
	transformedSignature := utils.TransformSignatureVRStoRSV(pld.signature)

	pk, err := crypto.SigToPub(messageHash, transformedSignature)
	if err != nil {
		return fmt.Errorf("recovering signer for %v", err)
	}

	pld.signer = crypto.PubkeyToAddress(*pk)

	pld.voterIndex = voterSet.VoterIndex(pld.signer)
	if pld.voterIndex < 0 {
		return fmt.Errorf("signer %s is not a registered voter in the current reward epoch", pld.signer.Hex())
	}

	pld.weight = voterSet.VoterWeight(pld.voterIndex)

	return nil
}
