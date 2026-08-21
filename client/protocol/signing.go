package protocol

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

// SignSignaturePayload signs the protocol message for a submitSignatures payload.
//
// The digest form follows the voting round embedded in the message bytes — the
// same bytes the Relay parses — so the signature covers what the contract will
// actually verify, whatever round label the payload travels under.
//
// A scheduled switch whose round boundary is not known yet signs the pre-switch
// way. That is right while the breaking epoch's policy does not exist, and only
// stale for as long as this node has failed to read one that does — the round tick
// re-resolves the boundary before any submitter of that round runs, so the window
// is a failed lookup, not an unbounded gap.
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
