package protocol

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

// SignSignaturePayload signs the protocol message of votingRoundID for a
// submitSignatures payload.
//
// The digest form follows the Relay that will finalize the round, which the
// finalizer picks by the signing policy's reward epoch.
//
// A scheduled switch whose round boundary is not known yet signs the pre-switch
// way. That is right while the breaking epoch's policy does not exist, and only
// stale for as long as this node has failed to read one that does — the round tick
// re-resolves the boundary before any submitter of that round runs, so the window
// is a failed lookup, not an unbounded gap.
func SignSignaturePayload(cutover *shared.RelayCutover, votingRoundID uint32, data []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	chainBound, known := cutover.NewRelayFromVotingRound(votingRoundID)
	if !known {
		logger.Debugf("Relay cutover: reward epoch %d has no known start round yet, signing round %d the old way",
			cutover.BreakingRewardEpoch, votingRoundID)
	}

	signature, err := crypto.Sign(shared.MessageDigest(data, cutover.ChainID, chainBound), privateKey)
	if err != nil {
		return nil, fmt.Errorf("signing message digest: %w", err)
	}

	return signature, nil
}
