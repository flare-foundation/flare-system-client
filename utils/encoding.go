package utils

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
)

const (
	hexPrefix             = "0x"
	addressChainSeparator = "-"
)

var (
	errInvalidPrefixError        = errors.New("string does not have hex prefix")
	errInvalidAddressLengthError = errors.New("address length is not 32")
	errInvalidIdLengthError      = errors.New("id length is not 20")
)

// DecodeHexString decodes a string that is prefixed with "0x" into a byte slice
func DecodeHexString(s string) ([]byte, error) {
	if !strings.HasPrefix(s, hexPrefix) {
		return nil, errInvalidPrefixError
	}
	return hex.DecodeString(s[len(hexPrefix):])
}

// Convert node id string to 20 byte hex string
func NodeIDToHex(nodeID string) (string, error) {
	id, err := ids.NodeIDFromString(nodeID)
	if err != nil {
		return "", err
	}
	return hexPrefix + hex.EncodeToString(id.Bytes()), nil
}

// Convert address string to 20 byte hex string
func AddressToHex(addrStr string) (string, error) {
	if !strings.Contains(addrStr, addressChainSeparator) {
		addrStr = addressChainSeparator + addrStr
	}
	id, err := address.ParseToID(addrStr)
	if err != nil {
		return "", err
	}
	return hexPrefix + hex.EncodeToString(id.Bytes()), nil
}

// Convert id string to 20 byte hex string
func IdToHex(idStr string) (string, error) {
	id, err := ids.FromString(idStr)
	if err != nil {
		return "", err
	}
	return hexPrefix + hex.EncodeToString(id[:]), nil
}

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
