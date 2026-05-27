package finalizer

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
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
