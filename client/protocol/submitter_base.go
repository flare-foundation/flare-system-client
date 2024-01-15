package protocol

import (
	"bytes"
	"flare-tlc/client/shared"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/chain"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type PayloadProvider interface {
	WritePayload(buffer *bytes.Buffer, currentEpoch int64, data []byte) error
}

type SubmitterBase struct {
	ethClient *ethclient.Client

	protocolContext *protocolContext

	epoch        *utils.Epoch
	selector     []byte
	subProtocols []*SubProtocol

	startOffset time.Duration
	epochOffset int64  // 0, -1
	nRetries    int    // number of retries for fetching data and submitting tx
	name        string // e.g., "submit1", "submit2", "submit3"

	payloadProvider PayloadProvider
}

func (s *SubmitterBase) GetPayload(currentEpoch int64) ([]byte, error) {
	channels := make([]<-chan shared.ExecuteStatus[[]byte], len(s.subProtocols))
	for i, protocol := range s.subProtocols {
		channels[i] = protocol.getDataWithRetry(
			currentEpoch+s.epochOffset,
			s.name,
			s.protocolContext.signingAddress.Hex(),
			s.nRetries,
		)
	}

	buffer := bytes.NewBuffer(nil)
	buffer.Write(s.selector)
	for _, channel := range channels {
		data := <-channel
		if !data.Success {
			logger.Error("error getting data for submitter %s: %s", s.name, data.Message)
			continue
		}
		err := s.payloadProvider.WritePayload(buffer, currentEpoch, data.Value)
		if err != nil {
			logger.Error("error writing payload for submitter %s: %v", s.name, err)
			return nil, err
		}
	}
	return buffer.Bytes(), nil
}

func (s *SubmitterBase) submit(payload []byte) {
	sendResult := <-shared.ExecuteWithRetry(func() (any, error) {
		err := chain.SendRawTx(s.ethClient, s.protocolContext.submitPrivateKey, s.protocolContext.submitContractAddress, payload)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error sending submit tx for submitter %s tx", s.name))
		}
		return nil, nil
	}, s.nRetries)
	if sendResult.Success {
		logger.Info("submitter %s submitted tx", s.name)
	}
}

func (s *SubmitterBase) Run() {
	ticker := utils.NewEpochTicker(s.startOffset, s.epoch)
	for {
		currentEpoch := <-ticker.C
		payload, err := s.GetPayload(currentEpoch)
		if err != nil {
			s.submit(payload)
		}
	}
}
