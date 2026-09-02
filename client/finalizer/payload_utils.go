package finalizer

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

// errBadPayload marks rejections caused by the submitted payload itself, as
// opposed to internal invariant violations. Expected in normal operation —
// log at Debug, not Error.
var errBadPayload = errors.New("bad payload")

// payloadMessage is a general structure that is used in the submit calls to the chain.
type payloadMessage struct {
	protocolID    uint8
	votingRoundID uint32
	payload       []byte
}

// ExtractPayloads extracts payloads from a transaction input to submission contracts and returns a slice of payloadMessages.
func ExtractPayloads(data []byte) ([]payloadMessage, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("wrongly formatted tx input: function selector needs 4 bytes, got %d", len(data))
	}

	messages := []payloadMessage{}

	data = data[4:] // trim function selector
	for len(data) > 0 {
		if len(data) < 7 { // 7 = 1 + 4 + 2
			return nil, fmt.Errorf("wrongly formatted tx input: header needs 7 bytes, %d remaining", len(data))
		}

		protocol := data[0]                               // 1 byte protocol ID
		votingRound := binary.BigEndian.Uint32(data[1:5]) // 4 bytes votingRoundID
		length := binary.BigEndian.Uint16(data[5:7])      // 2 bytes length of payload in bytes
		end := 7 + int(length)
		if len(data) < end {
			return nil, fmt.Errorf("wrongly formatted tx input: declared payload length %d exceeds %d remaining bytes", length, len(data)-7)
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
	sender common.Address // tx sender (submitSignatures address)

	protocolID    uint8
	votingRoundID uint32
	signature     []byte

	//assigned after processing
	signer     common.Address
	voterIndex int
	weight     uint16
}

// FromSignedPayload reads the signature out of a payload; a type-0 payload's inline message is skipped.
func (s *submitSignaturesPayload) FromSignedPayload(payloadMsg payloadMessage) error {
	if len(payloadMsg.payload) < 1 {
		return errors.New("empty payload")
	}
	typeID := payloadMsg.payload[0]

	var signatureStart int
	switch typeID {
	case 0:
		signatureStart = 1 + shared.RelayMessageLength
	case 1:
		signatureStart = 1
	default:
		return fmt.Errorf("invalid typeID %d", typeID)
	}
	signatureEnd := signatureStart + utils.SignatureLength

	if len(payloadMsg.payload) < signatureEnd {
		return fmt.Errorf("payload of type %d to short got %d, should be at least %d", typeID, len(payloadMsg.payload), signatureEnd)
	}

	s.protocolID = payloadMsg.protocolID
	s.votingRoundID = payloadMsg.votingRoundID
	s.signature = payloadMsg.payload[signatureStart:signatureEnd]
	s.voterIndex = -1 // 0 is a valid index, we use -1 before assigning the proper value

	return nil
}

var (
	// secp256k1N is the group order; secp256k1HalfN is the EIP-2 low-s bound, matching the
	// Relay's hardcoded literal (Relay.sol, ERR_BAD_S).
	secp256k1N     = crypto.S256().Params().N
	secp256k1HalfN = new(big.Int).Rsh(secp256k1N, 1)
)

// canonicalSignature returns the [V || R || S] signature in the low-s form the Relay
// demands, and whether it had to be normalized. relay() reverts the whole call on v outside
// {27,28} or s above n/2 (ERR_BAD_V / ERR_BAD_S, EIP-2); the previously deployed Relay checked
// neither, so one such signature counted toward the local threshold reverts that round on
// every finalizer. (r, s, v) and (r, n-s, v^1) recover the same signer, so high-s is
// normalized, not dropped — dropping would cost that voter's weight and could put the round
// below threshold.
func canonicalSignature(vrs []byte) ([]byte, bool, error) {
	if len(vrs) != utils.SignatureLength {
		return nil, false, fmt.Errorf("%w: signature is %d bytes, expected %d",
			errBadPayload, len(vrs), utils.SignatureLength)
	}

	v := vrs[0]
	if v != 27 && v != 28 {
		return nil, false, fmt.Errorf("%w: signature v is %d, expected 27 or 28", errBadPayload, v)
	}

	r := new(big.Int).SetBytes(vrs[1:33])
	s := new(big.Int).SetBytes(vrs[33:65])
	if !crypto.ValidateSignatureValues(v-27, r, s, false) {
		return nil, false, fmt.Errorf("%w: signature r or s is outside [1, n)", errBadPayload)
	}
	if s.Cmp(secp256k1HalfN) <= 0 {
		return vrs, false, nil
	}

	normalized := make([]byte, utils.SignatureLength)
	if v == 27 { // flip the recovery bit to match n-s
		normalized[0] = 28
	} else {
		normalized[0] = 27
	}
	copy(normalized[1:33], vrs[1:33])
	new(big.Int).Sub(secp256k1N, s).FillBytes(normalized[33:65])

	return normalized, true, nil
}

// AddSigner recovers the signer from the signature over digest and adds its voterIndex and
// weight, if the signer is in voterSet. digest must be shared.MessageDigest, not
// keccak256(message): the latter recovers a stranger, rejected as an unregistered voter.
// Canonicalizing first keeps Relay-rejected signature forms out of the finalization calldata.
func (pld *submitSignaturesPayload) AddSigner(digest []byte, voterSet *voters.Set) error {
	signature, normalized, err := canonicalSignature(pld.signature)
	if err != nil {
		return err
	}
	if normalized {
		logger.Warnf("Normalized a high-s signature from %s for protocol %d round %d; the Relay rejects that form",
			pld.sender, pld.protocolID, pld.votingRoundID)
	}
	pld.signature = signature

	transformedSignature, err := utils.TransformSignatureVRStoRSV(pld.signature)
	if err != nil {
		return fmt.Errorf("transforming signature: %w", err)
	}

	pk, err := crypto.SigToPub(digest, transformedSignature)
	if err != nil {
		return fmt.Errorf("%w: recovering signer: %w", errBadPayload, err)
	}

	pld.signer = crypto.PubkeyToAddress(*pk)

	pld.voterIndex = voterSet.VoterIndex(pld.signer)
	if pld.voterIndex < 0 {
		return fmt.Errorf("%w: signer %s is not a registered voter in the current reward epoch", errBadPayload, pld.signer.Hex())
	}

	pld.weight = voterSet.VoterWeight(pld.voterIndex)

	return nil
}
