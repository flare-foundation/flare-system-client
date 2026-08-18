package shared

import (
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

// Parse decodes the protocol message:
// protocolID(1) ‖ votingRoundID(4) ‖ isSecureRandom(1) ‖ merkleRoot(32).
func (msg Message) Parse() (RelayMessage, error) {
	if len(msg) != RelayMessageLength {
		return RelayMessage{}, fmt.Errorf("protocol message is %d bytes, expected %d", len(msg), RelayMessageLength)
	}
	return RelayMessage{
		ProtocolID:    msg[0],
		VotingRoundID: binary.BigEndian.Uint32(msg[1:5]),
		// the Relay normalizes any nonzero byte to 1, and the random Merkle leaf
		// commits to the normalized value — matching it is not optional
		IsSecureRandom: msg[5] != 0,
		MerkleRoot:     common.Hash(msg[6:RelayMessageLength]),
	}, nil
}

// MessageDigest returns the digest a protocol-message signature covers, as the
// Relay computes it before ecrecover. The new Relay binds the digest to its
// source chain — keccak256(chainID ‖ msg) — so a signature minted for another
// network is rejected even under an overlapping voter set; the old one hashes
// the message alone. Signing and signer recovery must pick the same form.
func MessageDigest(msg []byte, chainID int64, chainBound bool) []byte {
	if chainBound {
		return accounts.TextHash(crypto.Keccak256(ChainIDWord(chainID), msg))
	}
	return accounts.TextHash(crypto.Keccak256(msg))
}

// ChainIDWord left-pads chainID into the 32-byte word the Relay prepends to
// chain-bound hash preimages.
func ChainIDWord(chainID int64) []byte {
	var w [32]byte
	binary.BigEndian.PutUint64(w[24:], uint64(chainID))
	return w[:]
}

type ProtocolMessage struct {
	ProtocolID    uint8
	VotingRoundID uint32
	Message       Message
	// Served by the provider with the message: what relay() needs appended for this
	// protocol, if anything — the random number and its Merkle proof for the random
	// protocol on the new Relay.
	FinalizationData []byte
}
