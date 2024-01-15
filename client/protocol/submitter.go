package protocol

import (
	"bytes"
	"flare-tlc/client/config"
	"flare-tlc/utils"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type Submitter struct {
	SubmitterBase
}

type SignatureSubmitter struct {
	SubmitterBase
}

func newSubmitter(
	ethClient *ethclient.Client,
	pc *protocolContext,
	epoch *utils.Epoch,
	submitCfg *config.SubmitConfig,
	selector []byte,
	subProtocols []*SubProtocol,
	epochOffset int64,
	name string,
) *Submitter {
	submitter := &Submitter{
		SubmitterBase: SubmitterBase{
			ethClient:       ethClient,
			protocolContext: pc,
			epoch:           epoch,
			startOffset:     submitCfg.StartOffset,
			nRetries:        1,
			selector:        selector,
			subProtocols:    subProtocols,
			epochOffset:     epochOffset,
			name:            name,
		},
	}
	submitter.payloadProvider = submitter
	return submitter
}

func (s *Submitter) WritePayload(buffer *bytes.Buffer, currentEpoch int64, data []byte) error {
	buffer.Write(data)
	return nil
}

func newSignatureSubmitter(
	ethClient *ethclient.Client,
	pc *protocolContext,
	epoch *utils.Epoch,
	submitCfg *config.SubmitConfig,
	selector []byte,
	subProtocols []*SubProtocol,
) *SignatureSubmitter {
	submitter := &SignatureSubmitter{
		SubmitterBase: SubmitterBase{
			ethClient:       ethClient,
			protocolContext: pc,
			epoch:           epoch,
			startOffset:     submitCfg.StartOffset,
			nRetries:        4,
			selector:        selector,
			subProtocols:    subProtocols,
			epochOffset:     -1,
			name:            "submitSignatures",
		},
	}
	submitter.payloadProvider = submitter
	return submitter
}

func (s *SignatureSubmitter) WritePayload(buffer *bytes.Buffer, currentEpoch int64, data []byte) error {
	dataHash := accounts.TextHash(crypto.Keccak256(data))
	signature, err := crypto.Sign(dataHash, s.protocolContext.signerPrivateKey)
	if err != nil {
		return errors.Wrap(err, "error signing submitSignatures data")
	}

	epochBytes := uint32toBytes(uint32(currentEpoch - 1))
	lengthBytes := uint16toBytes(104) // 104 + length of additional data

	buffer.WriteByte(100)        // Protocol ID (1 byte)
	buffer.Write(epochBytes[:])  // Epoch (4 bytes)
	buffer.Write(lengthBytes[:]) // Length (2 bytes)

	buffer.WriteByte(0) // Type (1 byte)
	buffer.Write(data)  // Message (38 bytes)

	buffer.WriteByte(signature[64] + 27) // V (1 byte)
	buffer.Write(signature[0:32])        // R (32 bytes)
	buffer.Write(signature[32:64])       // S (32 bytes)

	buffer.Write(data)

	// Todo: append additional data
	return nil
}
