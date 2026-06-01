package finalizer

import (
	"crypto/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/flare-foundation/go-flare-common/pkg/policy"
	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

func bufferPayload(t *testing.T, pc *protocolCollection, sender common.Address) {
	t.Helper()
	_, _, err := pc.addPayload(&submitSignaturesPayload{typeID: 1, sender: sender})
	require.NoError(t, err)
}

// Before messageAdded, a sender may buffer at most one payload per (round,
// protocol); extras are dropped to bound the ECDSA-recovery burst on drain (DOS-01).
func TestProtocolCollectionBuffersOnePayloadPerSender(t *testing.T) {
	pc := &protocolCollection{signatureCollection: map[common.Hash]*signaturesCollection{}}
	sender := common.HexToAddress("0x1111111111111111111111111111111111111111")

	for range 100 {
		bufferPayload(t, pc, sender)
	}

	require.Len(t, pc.unprocessedPayloads, 1, "only the first payload per sender should be buffered")
}

func TestProtocolCollectionBuffersPerSenderIndependently(t *testing.T) {
	pc := &protocolCollection{signatureCollection: map[common.Hash]*signaturesCollection{}}
	a := common.HexToAddress("0xaaaa000000000000000000000000000000000001")
	b := common.HexToAddress("0xbbbb000000000000000000000000000000000002")

	bufferPayload(t, pc, a)
	bufferPayload(t, pc, a) // dropped
	bufferPayload(t, pc, b)
	bufferPayload(t, pc, b) // dropped

	require.Len(t, pc.unprocessedPayloads, 2, "one buffered payload per distinct sender")
}

// TypeID-0 payloads carry the message inline, so an attacker controls the
// messageHash key per payload. Without the cap, each payload allocates a fresh
// signaturesCollection before any signer check, giving an unbounded memory
// growth vector. The cap limits allocations to one per sender per (round, protocol).
func TestProtocolCollectionCapsTypeZeroAllocations(t *testing.T) {
	sender := common.HexToAddress("0x1111111111111111111111111111111111111111")
	pc := &protocolCollection{
		signatureCollection: map[common.Hash]*signaturesCollection{},
		signingPolicy: &policy.SigningPolicy{
			Voters: voters.NewSet([]common.Address{sender}, []uint16{1}, nil),
		},
	}

	for range 100 {
		// fresh 38-byte message per call → unique hash per payload
		msg := make(shared.Message, 38)
		_, err := rand.Read(msg)
		require.NoError(t, err)
		// Errors from AddSigner (bad signature) are expected and irrelevant here;
		// what we assert is that no further allocations occur past the cap.
		_, _, _ = pc.addPayload(&submitSignaturesPayload{
			typeID:    0,
			sender:    sender,
			message:   msg,
			signature: make([]byte, 65), // satisfy AddSigner length; recovery will fail
		})
	}

	require.LessOrEqual(t, len(pc.signatureCollection), 1,
		"only the first type-0 payload per sender may allocate a signaturesCollection")
}
