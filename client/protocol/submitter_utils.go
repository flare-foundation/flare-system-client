package protocol

import (
	"flare-tlc/logger"
	"flare-tlc/utils"
	"github.com/ethereum/go-ethereum/common/math"
)

type EpochRunner interface {
	GetEpochTicker() *utils.EpochTicker
	RunEpoch(currentEpoch int64)
}

func Run(r EpochRunner, stopAt <-chan int64, lastEpoch chan<- int64) {
	ticker := r.GetEpochTicker()
	var epoch int64
	stopAfterEpoch := int64(math.MaxInt64)

	for {
		if epoch >= stopAfterEpoch {
			break
		}
		select {
		case stopAfterEpoch = <-stopAt:
			logger.Info("Stopping submitter after epoch %d", stopAfterEpoch)
		case epoch = <-ticker.C:
			lastEpoch <- epoch
			r.RunEpoch(epoch)
		}
	}
	logger.Info("Submitter stopped")
}
