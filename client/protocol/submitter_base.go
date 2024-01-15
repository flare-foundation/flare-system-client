package protocol

import (
	"encoding/binary"
	"flare-tlc/client/shared"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/chain"
	"fmt"
	"math"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type EpochRunner interface {
	RunEpoch(currentEpoch int64)
}

type SubmitterBase struct {
	DataVerifier
	EpochRunner

	ethClient *ethclient.Client

	protocolContext *protocolContext

	epoch        *utils.Epoch
	selector     []byte
	subProtocols []*SubProtocol

	startOffset   time.Duration
	submitRetries int    // number of retries for submitting tx
	name          string // e.g., "submit1", "submit2", "submit3", "signatureSubmitter"
}

func (s *SubmitterBase) submit(payload []byte) bool {
	sendResult := <-shared.ExecuteWithRetry(func() (any, error) {
		err := chain.SendRawTx(s.ethClient, s.protocolContext.submitPrivateKey, s.protocolContext.submitContractAddress, payload)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error sending submit tx for submitter %s tx", s.name))
		}
		return nil, nil
	}, s.submitRetries, shared.TxRetryInterval)
	if sendResult.Success {
		logger.Info("submitter %s submitted tx", s.name)
	}
	return sendResult.Success
}

func (s *SubmitterBase) Run() {
	ticker := utils.NewEpochTicker(s.startOffset, s.epoch)
	for {
		currentEpoch := <-ticker.C
		s.RunEpoch(currentEpoch)
	}
}

func (s *SubmitterBase) VerifyData(data *SubProtocolResponse) error {
	if data.Status != "OK" {
		return fmt.Errorf("status %s", data.Status)
	}
	if len(data.Data) != 38 {
		return fmt.Errorf("data length %d is not 38", len(data.Data))
	}
	// Check if additional data is too long
	// Length of data without additional data is 104 bytes: 1 (type) + 38 (message) + 65 (signature)
	if len(data.AdditionalData) > math.MaxUint16-104 {
		return errors.New("additional data too long")
	}
	return nil
}

func uint16toBytes(i uint16) (arr [2]byte) {
	binary.BigEndian.PutUint16(arr[0:2], i)
	return
}

func uint32toBytes(i uint32) (arr [4]byte) {
	binary.BigEndian.PutUint32(arr[0:4], i)
	return
}
