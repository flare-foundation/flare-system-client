package protocol

import (
	"encoding/binary"
	"flare-tlc/utils"
)

type EpochRunner interface {
	GetEpochTicker() *utils.EpochTicker
	RunEpoch(currentEpoch int64)
}

func Run(r EpochRunner) {
	ticker := r.GetEpochTicker()
	for {
		currentEpoch := <-ticker.C
		r.RunEpoch(currentEpoch)
	}
}

func uint16toBytes(i uint16) (arr [2]byte) {
	binary.BigEndian.PutUint16(arr[0:2], i)
	return
}

func uint32toBytes(i uint32) (arr [4]byte) {
	binary.BigEndian.PutUint32(arr[0:4], i)
	return
}
