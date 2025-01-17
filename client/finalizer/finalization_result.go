package finalizer

import (
	"bytes"
	"fmt"
	"math"
	"slices"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/flare-foundation/go-flare-common/pkg/policy"
)

type FinalizationResult struct {
	message       shared.Message
	signatures    []IndexedSignature //signatures are ordered by voterIndex of their provider
	signingPolicy *policy.SigningPolicy
}

type IndexedSignature struct {
	index     int
	signature []byte
}

// PrepareFinalizationResults returns the message and signatures that are needed to construct the transaction input that is needed for the finalization.
//
// The signatures are chosen in a way to minimize the number of signatures needed for finalization.
func PrepareFinalizationResults(sc *signaturesCollection) (FinalizationResult, error) {
	availableSignatures := []IndexedSignature{}
	selectedSignatures := []IndexedSignature{}

	for i := range sc.signatures {
		if len(sc.signatures[i]) > 0 {
			availableSignatures = append(availableSignatures, IndexedSignature{index: i, signature: sc.signatures[i]})
		}
	}

	// sort decreasing by weight
	slices.SortFunc(availableSignatures, func(a, b IndexedSignature) int {
		return int(sc.signingPolicy.Voters.VoterWeight(b.index)) - int(sc.signingPolicy.Voters.VoterWeight(a.index))
	})

	// greedy select until threshold is reached
	weight := uint16(0)
	for i := range availableSignatures {
		selectedSignatures = append(selectedSignatures, availableSignatures[i])
		weight += sc.signingPolicy.Voters.VoterWeight(availableSignatures[i].index)
		if weight > sc.threshold {
			break
		}
	}
	if weight <= sc.threshold {
		return FinalizationResult{}, fmt.Errorf("threshold not reached")
	}

	// sort selected by index
	slices.SortFunc(selectedSignatures, func(p, q IndexedSignature) int {
		return p.index - q.index
	})

	return FinalizationResult{message: sc.message, signatures: selectedSignatures, signingPolicy: sc.signingPolicy}, nil
}

// PrepareFinalizationTxInput prepares a transaction input needed to finalize.
func (fr FinalizationResult) PrepareFinalizationTxInput() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(relayFunctionSelector)
	buffer.Write(fr.signingPolicy.RawBytes())
	buffer.Write(fr.message)

	encodedSignatures, err := encodeSignatures(fr.signatures)
	if err != nil {
		return nil, err
	}

	buffer.Write(encodedSignatures)

	return buffer.Bytes(), nil
}

// encodeSignatures encodes indexed signature to be used in the finalization transaction input.
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
