package protocol

import (
	"bytes"
	"crypto/ecdsa"
	"flare-tlc/client/config"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/chain"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type SignatureSubmitter struct {
	ethClient             *ethclient.Client
	submitPrivateKey      *ecdsa.PrivateKey
	submitContractAddress common.Address

	signingAddress   common.Address
	signerPrivateKey *ecdsa.PrivateKey

	epoch        *utils.Epoch
	startOffset  time.Duration
	selector     []byte
	subProtocols []*SubProtocol
}

func newSignatureSubmitter(
	ethClient *ethclient.Client,
	pc *protocolCredentials,
	pa *protocolAddresses,
	epoch *utils.Epoch,
	submitCfg *config.SubmitConfig,
	selector []byte,
	subProtocols []*SubProtocol,
) *SignatureSubmitter {
	return &SignatureSubmitter{
		ethClient:             ethClient,
		submitPrivateKey:      pc.submitPrivateKey,
		submitContractAddress: pa.SubmitContractAddress,
		signingAddress:        pa.signingAddress,
		signerPrivateKey:      pc.signerPrivateKey,
		epoch:                 epoch,
		startOffset:           submitCfg.StartOffset,
		selector:              selector,
		subProtocols:          subProtocols,
	}
}

func (s *SignatureSubmitter) GetPayload(currentEpoch int64) ([]byte, error) {
	channels := make([]<-chan []byte, len(s.subProtocols))
	previousEpoch := currentEpoch - 1
	for i, protocol := range s.subProtocols {
		// Todo: get data with retry
		channels[i] = protocol.getDataAsync(previousEpoch, "submitSignatures", s.signingAddress.Hex())
	}

	buffer := bytes.NewBuffer(nil)
	buffer.Write(s.selector)
	for _, channel := range channels {
		data := <-channel

		// Todo: handle error
		if len(data) == 0 {
			continue
		}

		dataHash := accounts.TextHash(crypto.Keccak256(data))
		signature, err := crypto.Sign(dataHash, s.signerPrivateKey)
		if err != nil {
			return nil, errors.Wrap(err, "error signing submitSignatures data")
		}

		epochBytes := uint32toBytes(uint32(previousEpoch))
		lengthBytes := uint16toBytes(104) // 104 + length of additional data

		buffer.WriteByte(100)        // Protocol ID (1 byte)
		buffer.Write(epochBytes[:])  // Epoch (4 bytes)
		buffer.Write(lengthBytes[:]) // Length (2 bytes)

		buffer.WriteByte(0) // Type (1 byte)
		buffer.Write(data)  // Message (38 bytes)

		buffer.WriteByte(signature[64] + 27) // V (1 byte)
		buffer.Write(signature[0:32])        // R (32 bytes)
		buffer.Write(signature[32:64])       // S (32 bytes)

		// Todo: append additional data

	}
	return buffer.Bytes(), nil
}

func (s *SignatureSubmitter) submit(payload []byte) {
	// Todo: retry
	err := chain.SendRawTx(s.ethClient, s.submitPrivateKey, s.submitContractAddress, payload)
	if err != nil {
		logger.Error("error sending submit tx by signatureSubmitter: %v", err)
		return
	}
	logger.Info("submitSugnature tx submitted")
}

func (s *SignatureSubmitter) Run() {
	ticker := utils.NewEpochTicker(s.startOffset, s.epoch)
	for {
		currentEpoch := <-ticker.C
		payload, _ := s.GetPayload(currentEpoch)
		s.submit(payload)
	}
}
