package protocol

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

// SignSignaturePayload signs the protocol message for a submitSignatures payload.
// The digest form follows the voting round in the message bytes — the same bytes the
// Relay parses — so the signature covers what the contract verifies, whatever round
// label the payload travels under. An unknown boundary signs the pre-switch way; the
// round tick re-resolves it before any submitter of that round runs.
func SignSignaturePayload(cutover *shared.RelayCutover, data []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	digest, known, err := cutover.DigestFromMessage(data)
	if err != nil {
		return nil, fmt.Errorf("deriving the message digest: %w", err)
	}
	if !known {
		logger.Debugf("Relay cutover: reward epoch %d has no known start round yet, signing the old way",
			cutover.BreakingRewardEpoch)
	}

	signature, err := crypto.Sign(digest, privateKey)
	if err != nil {
		return nil, fmt.Errorf("signing message digest: %w", err)
	}

	return signature, nil
}
