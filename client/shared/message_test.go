package shared

import (
	"encoding/binary"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func protocolMessage(protocolID uint8, round uint32, isSecureByte byte, root common.Hash) Message {
	msg := make(Message, RelayMessageLength)
	msg[0] = protocolID
	binary.BigEndian.PutUint32(msg[1:5], round)
	msg[5] = isSecureByte
	copy(msg[6:], root[:])
	return msg
}

// protocolID(1) ‖ votingRoundID(4) ‖ isSecureRandom(1) ‖ merkleRoot(32); other lengths error.
func TestParseDecodesTheMessage(t *testing.T) {
	root := common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")

	parsed, err := protocolMessage(200, 99, 0, root).Parse()
	require.NoError(t, err)
	require.Equal(t, uint8(200), parsed.ProtocolID)
	require.Equal(t, uint32(99), parsed.VotingRoundID)
	require.Equal(t, root, parsed.MerkleRoot)
	require.False(t, parsed.IsSecureRandom)

	_, err = Message(make([]byte, RelayMessageLength-1)).Parse()
	require.Error(t, err)
	_, err = Message(make([]byte, RelayMessageLength+1)).Parse()
	require.Error(t, err)
}

// The Relay normalizes the isSecureRandom byte with `!= 0`, and the leaf commits to that.
func TestParseNormalizesIsSecureRandom(t *testing.T) {
	for _, b := range []byte{1, 2, 0xff} {
		parsed, err := protocolMessage(100, 1, b, common.Hash{}).Parse()
		require.NoError(t, err)
		require.True(t, parsed.IsSecureRandom, "byte %d", b)
	}

	parsed, err := protocolMessage(100, 1, 0, common.Hash{}).Parse()
	require.NoError(t, err)
	require.False(t, parsed.IsSecureRandom)
}
