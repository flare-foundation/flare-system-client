package shared

import "encoding/binary"

func Uint16toBytes(i uint16) (arr [2]byte) {
	binary.BigEndian.PutUint16(arr[0:2], i)
	return
}

func Uint32toBytes(i uint32) (arr [4]byte) {
	binary.BigEndian.PutUint32(arr[0:4], i)
	return
}
