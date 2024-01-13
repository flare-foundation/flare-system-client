package protocol

import (
	"bytes"
	"crypto/ecdsa"
	"flare-tlc/client/config"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/chain"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Submitter struct {
	ethClient             *ethclient.Client
	submitPrivateKey      *ecdsa.PrivateKey
	submitContractAddress common.Address
	signingAddress        common.Address

	epoch        *utils.Epoch
	startOffset  time.Duration
	selector     []byte
	subProtocols []*SubProtocol
	epochOffset  int64  // 0, -1
	name         string // e.g., "submit1", "submit2", "submit3"
}

func newSubmitter(
	ethClient *ethclient.Client,
	pc *protocolCredentials,
	pa *protocolAddresses,
	epoch *utils.Epoch,
	submitCfg *config.SubmitConfig,
	selector []byte,
	subProtocols []*SubProtocol,
	epochOffset int64,
	name string,
) *Submitter {
	return &Submitter{
		ethClient:             ethClient,
		submitPrivateKey:      pc.submitPrivateKey,
		submitContractAddress: pa.SubmitContractAddress,
		signingAddress:        pa.signingAddress,
		epoch:                 epoch,
		startOffset:           submitCfg.StartOffset,
		selector:              selector,
		subProtocols:          subProtocols,
		epochOffset:           epochOffset,
		name:                  name,
	}
}

func (s *Submitter) GetPayload(currentEpoch int64) []byte {
	channels := make([]<-chan []byte, len(s.subProtocols))
	for i, protocol := range s.subProtocols {
		channels[i] = protocol.getDataAsync(currentEpoch+s.epochOffset, s.name, s.signingAddress.Hex())
	}

	buffer := bytes.NewBuffer(nil)
	buffer.Write(s.selector)
	for _, channel := range channels {
		data := <-channel
		buffer.Write(data)
	}
	return buffer.Bytes()
}

func (s *Submitter) submit(payload []byte) {
	err := chain.SendRawTx(s.ethClient, s.submitPrivateKey, s.submitContractAddress, payload)
	if err != nil {
		logger.Error("error sending submit tx for submitter %s tx: %v", s.name, err)
		return
	}
	logger.Info("submitter %s submitted tx", s.name)
}

func (s *Submitter) Run() {
	ticker := utils.NewEpochTicker(s.startOffset, s.epoch)
	for {
		currentEpoch := <-ticker.C
		payload := s.GetPayload(currentEpoch)
		s.submit(payload)
	}
}
