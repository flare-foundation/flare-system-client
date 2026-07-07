package utils

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

var (
	errInvalidAddressLengthError = errors.New("address length is not 32")
	errInvalidIdLengthError      = errors.New("id length is not 20")
)

func UInt64ToHex(value uint64) string {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, value)
	return hex.EncodeToString(buf)
}

func UInt32ToHex(value uint32) string {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, value)
	return hex.EncodeToString(buf)
}

func UInt16ToHex(value uint16) string {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, value)
	return hex.EncodeToString(buf)
}

// Checks if the string is a valid hex string and pad it to
// the given length (even number)
func PadHexString(value string, length int) (string, error) {
	if length%2 != 0 {
		return "", errors.New("length must be even")
	}
	value = strings.TrimPrefix(value, "0x")
	if len(value)%2 != 0 {
		value = "0" + value
	}
	_, err := hex.DecodeString(value)
	if err != nil {
		return "", err
	}
	if len(value) > length {
		return "", errors.New("string too long")
	}
	return strings.Repeat("0", length-len(value)) + value, nil
}

func TransactionHexToBytes32(address string) (result [32]byte, err error) {
	address = strings.TrimPrefix(address, "0x")
	addressBytes, err := hex.DecodeString(address)
	if err != nil {
		return result, err
	}
	if len(addressBytes) != 32 {
		return result, errInvalidAddressLengthError
	}
	copy(result[:], addressBytes)
	return
}

func Hex20ToBytes20(str string) (result [20]byte, err error) {
	str = strings.TrimPrefix(str, "0x")
	strBytes, err := hex.DecodeString(str)
	if err != nil {
		return result, err
	}
	if len(strBytes) != 20 {
		return result, errInvalidIdLengthError
	}
	copy(result[:], strBytes)
	return
}

// SignatureLength is the length in bytes of an ECDSA signature in both [V || R || S] and [R || S || V] forms.
const SignatureLength = 65

// TransformSignatureVRStoRSV transforms a signature to be used by go-ethereum crypto.SigToPub:
// transforms [V || R || S] to [R || S || V - 27].
func TransformSignatureVRStoRSV(vrs []byte) ([]byte, error) {
	if len(vrs) != SignatureLength {
		return nil, fmt.Errorf("invalid signature length %d, expected %d", len(vrs), SignatureLength)
	}

	rsv := make([]byte, SignatureLength)

	copy(rsv[:], vrs[1:33])
	copy(rsv[32:], vrs[33:65])
	rsv[64] = vrs[0] - 27

	return rsv, nil
}

// TransformSignatureRSVtoVRS transforms [R || S || V - 27] to [V || R || S].
func TransformSignatureRSVtoVRS(rsv []byte) ([]byte, error) {
	if len(rsv) != SignatureLength {
		return nil, fmt.Errorf("invalid signature length %d, expected %d", len(rsv), SignatureLength)
	}

	vrs := make([]byte, SignatureLength)

	vrs[0] = rsv[64] + 27
	copy(vrs[1:], rsv[0:32])
	copy(vrs[33:], rsv[32:64])

	return vrs, nil
}
