package finalizer

import (
	"bytes"
	"flare-fsc/client/shared"
	"fmt"
	"math"
)

type FinalizationResult struct {
	message       shared.Message
	signatures    []IndexedSignature //signatures are ordered by voterIndex of their provider
	signingPolicy *signingPolicy
}

type IndexedSignature struct {
	index     int
	signature []byte
}

// PrepareFinalizationTxInput prepares a tx input needed to finalize.
func (fr FinalizationResult) PrepareFinalizationTxInput() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(relayFunctionSelector)
	buffer.Write(fr.signingPolicy.rawBytes)
	buffer.Write(fr.message)

	encodedSignatures, err := encodeSignatures(fr.signatures)
	if err != nil {
		return nil, err
	}

	buffer.Write(encodedSignatures)

	return buffer.Bytes(), nil
}

// encodeSignatures encodes indexed signature to be used in the finalization transaction.
// Signatures should be ordered by the indexes of their providers.
func encodeSignatures(signatures []IndexedSignature) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if len(signatures) > math.MaxUint16 {
		return nil, fmt.Errorf("too many payloads: %d", len(signatures))
	}

	sizeBytes := shared.Uint16toBytes(uint16(len(signatures)))
	buffer.Write(sizeBytes[:])
	prevIndex := -1
	for _, signature := range signatures {
		if signature.index < 0 {
			return nil, fmt.Errorf("payload index not set")
		}
		if prevIndex >= signature.index {
			return nil, fmt.Errorf("payloads not sorted by index")
		}

		indexBytes := shared.Uint16toBytes(uint16(signature.index))
		buffer.Write(signature.signature)
		buffer.Write(indexBytes[:])
		prevIndex = signature.index
	}
	return buffer.Bytes(), nil
}
