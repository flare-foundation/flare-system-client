package shared

import (
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/flare-foundation/go-flare-common/pkg/convert"
)

type Message []byte

// RelayMessageLength is the fixed size of a protocol message.
const RelayMessageLength = 38

// RelayMessage is the decoded protocol message that signatures cover and relay() consumes.
type RelayMessage struct {
	ProtocolID     uint8
	VotingRoundID  uint32
	IsSecureRandom bool
	MerkleRoot     common.Hash
}

// Parse decodes protocolID(1) ‖ votingRoundID(4) ‖ isSecureRandom(1) ‖ merkleRoot(32).
func (msg Message) Parse() (RelayMessage, error) {
	if len(msg) != RelayMessageLength {
		return RelayMessage{}, fmt.Errorf("protocol message is %d bytes, expected %d", len(msg), RelayMessageLength)
	}
	return RelayMessage{
		ProtocolID:    msg[0],
		VotingRoundID: binary.BigEndian.Uint32(msg[1:5]),
		// the Relay normalizes any nonzero byte to 1 and the random Merkle leaf commits to
		// the normalized value — must match
		IsSecureRandom: msg[5] != 0,
		MerkleRoot:     common.Hash(msg[6:RelayMessageLength]),
	}, nil
}

// MessageDigest returns the digest the Relay computes before ecrecover. It binds the source
// chain — keccak256(chainID ‖ msg) — so a signature minted for another network with an
// overlapping voter set is rejected.
func MessageDigest(msg []byte, chainID int64) []byte {
	return accounts.TextHash(crypto.Keccak256(convert.Uint64ToHash(uint64(chainID)).Bytes(), msg))
}

type ProtocolMessage struct {
	ProtocolID    uint8
	VotingRoundID uint32
	Message       Message
	// served by the provider with the message: what relay() needs appended, if anything — the
	// random number and its Merkle proof for the random protocol
	FinalizationData []byte
}
