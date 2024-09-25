package utils

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
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
	binary.LittleEndian.PutUint32(buf, uint32(value))
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

// Transform signature to be used by go-ethereum crypto.SigToPub:
// transforms [V || R || S] to [R || S || V - 27]
// No checks are performed, we assume that signature array has length 65
func TransformSignatureVRStoRSV(vrs []byte) (rsv []byte) {
	rsv = make([]byte, 65)

	copy(rsv[:], vrs[1:33])
	copy(rsv[32:], vrs[33:65])
	rsv[64] = vrs[0] - 27

	return rsv
}

// Transform signature transforms [R || S || V - 27] to [V || R || S].
// No checks are performed, we assume that signature array has length 65
func TransformSignatureRSVtoVRS(rsv []byte) (vrs []byte) {
	vrs = make([]byte, 65)

	vrs[0] = rsv[64] + 27
	copy(vrs[1:], rsv[0:32])
	copy(vrs[33:], rsv[32:64])

	return vrs
}
