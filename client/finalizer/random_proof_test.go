package finalizer

import (
	"encoding/binary"
	"testing"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

var (
	randomValue = common.HexToHash("0x00000000000000000000000000000000000000000000000000000000deadbeef")
	proofNodeA  = common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")
	proofNodeB  = common.HexToHash("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
)

// buildMessage assembles the signed protocol message:
// protocolID(1) ‖ votingRoundID(4) ‖ isSecureRandom(1) ‖ merkleRoot(32).
func buildMessage(protocolID uint8, votingRoundID uint32, isSecureByte byte, root common.Hash) shared.Message {
	msg := make(shared.Message, shared.RelayMessageLength)
	msg[0] = protocolID
	binary.BigEndian.PutUint32(msg[1:5], votingRoundID)
	msg[5] = isSecureByte
	copy(msg[6:], root[:])
	return msg
}

// The leaf must be keccak256(abi.encode(uint256 votingRoundId, uint256 value,
// uint256 isSecure)) — three left-padded words, as Relay.processRandomMerkleProof
// builds it (mirrored in RelayIsSecureNormFV._randomLeaf).
func TestRandomLeafMatchesAbiEncode(t *testing.T) {
	expected := make([]byte, 96)
	binary.BigEndian.PutUint32(expected[28:32], 4242)
	copy(expected[32:64], randomValue[:])
	expected[95] = 1

	require.Equal(t, crypto.Keccak256Hash(expected), randomLeaf(4242, randomValue, true))

	// isSecure false leaves the third word zero
	expected[95] = 0
	require.Equal(t, crypto.Keccak256Hash(expected), randomLeaf(4242, randomValue, false))
}

// The Relay normalizes the message's isSecureRandom byte with `!= 0`, and the leaf
// commits to the normalized value — so any nonzero byte must give the same leaf.
func TestIsSecureNormalization(t *testing.T) {
	secure := randomLeaf(1, randomValue, true)

	for _, b := range []byte{1, 2, 0xff} {
		parsed, err := buildMessage(100, 1, b, common.Hash{}).Parse()
		require.NoError(t, err)
		require.True(t, parsed.IsSecureRandom, "byte %d", b)
		require.Equal(t, secure, randomLeaf(1, randomValue, parsed.IsSecureRandom), "byte %d", b)
	}

	parsed, err := buildMessage(100, 1, 0, common.Hash{}).Parse()
	require.NoError(t, err)
	require.False(t, parsed.IsSecureRandom)
	require.NotEqual(t, secure, randomLeaf(1, randomValue, parsed.IsSecureRandom))
}

// Sorted-pair hashing orders operands as unsigned 32-byte integers, so a proof
// carries no left/right bit and the fold is order-independent within a step.
func TestSortedPair(t *testing.T) {
	require.Equal(t, sortedPair(proofNodeA, proofNodeB), sortedPair(proofNodeB, proofNodeA))
	require.Equal(t, crypto.Keccak256Hash(proofNodeA[:], proofNodeB[:]), sortedPair(proofNodeA, proofNodeB))
}

// An empty proof folds to the leaf itself; the Relay accepts that for a single-leaf
// tree, so the client must not reject it.
func TestFoldEmptyProofIsTheLeaf(t *testing.T) {
	leaf := randomLeaf(7, randomValue, true)
	require.Equal(t, leaf, foldProof(leaf, nil))
}

func TestVerifyAcceptsAMatchingProof(t *testing.T) {
	leaf := randomLeaf(7, randomValue, true)
	root := sortedPair(sortedPair(leaf, proofNodeA), proofNodeB)

	message, err := buildMessage(100, 7, 1, root).Parse()
	require.NoError(t, err)

	data := randomData{Value: randomValue, Proof: []common.Hash{proofNodeA, proofNodeB}}
	require.NoError(t, data.verify(message))
}

func TestVerifyRejectsBadData(t *testing.T) {
	leaf := randomLeaf(7, randomValue, true)
	root := sortedPair(sortedPair(leaf, proofNodeA), proofNodeB)
	proof := []common.Hash{proofNodeA, proofNodeB}

	message, err := buildMessage(100, 7, 1, root).Parse()
	require.NoError(t, err)

	cases := []struct {
		name string
		data randomData
	}{
		{"wrong value", randomData{Value: common.HexToHash("0x01"), Proof: proof}},
		{"missing proof node", randomData{Value: randomValue, Proof: proof[:1]}},
		{"empty proof", randomData{Value: randomValue, Proof: nil}},
		{"reordered proof", randomData{Value: randomValue, Proof: []common.Hash{proofNodeB, proofNodeA}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Error(t, c.data.verify(message))
		})
	}

	// a proof valid for another round must not verify for this one
	otherLeaf := randomLeaf(8, randomValue, true)
	otherMessage, err := buildMessage(100, 7, 1, sortedPair(otherLeaf, proofNodeA)).Parse()
	require.NoError(t, err)
	require.Error(t, randomData{Value: randomValue, Proof: []common.Hash{proofNodeA}}.verify(otherMessage))

	// the leaf's secure flag comes from the signed message, so a proof the source
	// built for the other flag cannot fold to the signed root
	insecureLeaf := randomLeaf(7, randomValue, false)
	wrongFlagMessage, err := buildMessage(100, 7, 1, sortedPair(insecureLeaf, proofNodeA)).Parse()
	require.NoError(t, err)
	require.Error(t, randomData{Value: randomValue, Proof: []common.Hash{proofNodeA}}.verify(wrongFlagMessage))
}

// The trailer is randomNumber(32) ‖ proof(32×d) — the Relay rejects anything that is
// not a whole number of words.
func TestTrailerEncoding(t *testing.T) {
	data := randomData{Value: randomValue, Proof: []common.Hash{proofNodeA, proofNodeB}}

	trailer := data.trailer()
	require.Len(t, trailer, 96)
	require.Equal(t, randomValue[:], trailer[0:32])
	require.Equal(t, proofNodeA[:], trailer[32:64])
	require.Equal(t, proofNodeB[:], trailer[64:96])

	require.Len(t, randomData{Value: randomValue}.trailer(), 32)
}

func TestParseRejectsAWrongLengthMessage(t *testing.T) {
	_, err := shared.Message(make([]byte, 37)).Parse()
	require.Error(t, err)

	parsed, err := buildMessage(200, 99, 0, proofNodeA).Parse()
	require.NoError(t, err)
	require.Equal(t, uint8(200), parsed.ProtocolID)
	require.Equal(t, uint32(99), parsed.VotingRoundID)
	require.Equal(t, proofNodeA, parsed.MerkleRoot)
}
