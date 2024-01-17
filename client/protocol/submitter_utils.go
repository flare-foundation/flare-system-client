package protocol

import (
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
