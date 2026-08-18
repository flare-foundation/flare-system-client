package finalizer

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// randomData is a voting round's random number together with the Merkle proof that
// it is a leaf of the tree whose root that round's providers signed.
//
// From the Relay cutover on, relay() requires both for the random protocol: it
// rebuilds the leaf, folds the proof, and reverts unless the result is the signed
// root (Relay.sol processRandomMerkleProof).
type randomData struct {
	Value common.Hash
	Proof []common.Hash
}

// parseRandomData reads the finalization data the random protocol's provider serves with
// its message: value(32) ‖ proof(32×d), the trailer bytes themselves. A whole number of
// words, at least one; zero proof nodes is a single-leaf tree.
func parseRandomData(raw []byte) (randomData, error) {
	if len(raw) == 0 || len(raw)%common.HashLength != 0 {
		return randomData{}, fmt.Errorf("finalization data is %d bytes, want value(32) followed by whole proof nodes", len(raw))
	}
	d := randomData{Value: common.Hash(raw[:common.HashLength])}
	for i := common.HashLength; i < len(raw); i += common.HashLength {
		d.Proof = append(d.Proof, common.Hash(raw[i:i+common.HashLength]))
	}
	return d, nil
}

// randomLeaf mirrors the Relay's leaf:
// keccak256(abi.encode(uint256 votingRoundId, uint256 value, uint256 isSecure)).
func randomLeaf(votingRoundID uint32, value common.Hash, isSecure bool) common.Hash {
	var buf [96]byte
	binary.BigEndian.PutUint32(buf[28:32], votingRoundID)
	copy(buf[32:64], value[:])
	if isSecure {
		buf[95] = 1
	}
	return crypto.Keccak256Hash(buf[:])
}

// sortedPair is OpenZeppelin's parent hash, as the Relay folds it: operands are
// ordered as unsigned 32-byte integers, so a proof carries no left/right bit.
func sortedPair(a, b common.Hash) common.Hash {
	if bytes.Compare(a[:], b[:]) > 0 {
		a, b = b, a
	}
	return crypto.Keccak256Hash(a[:], b[:])
}

// foldProof walks the proof bottom-up. An empty proof yields the leaf itself, which
// the Relay accepts for a single-leaf tree.
func foldProof(leaf common.Hash, proof []common.Hash) common.Hash {
	hash := leaf
	for _, node := range proof {
		hash = sortedPair(hash, node)
	}
	return hash
}

// verify repeats the check relay() performs, so a proof that cannot possibly
// finalize is never broadcast.
//
// The round and the secure flag are taken from the signed message rather than from
// the response, so a wrong or hostile source cannot steer the leaf: only the value
// and the proof are taken on trust, and the fold pins both to the signed root.
func (d randomData) verify(message shared.RelayMessage) error {
	leaf := randomLeaf(message.VotingRoundID, d.Value, message.IsSecureRandom)
	if root := foldProof(leaf, d.Proof); root != message.MerkleRoot {
		return fmt.Errorf("random proof folds to %s, the signed merkle root is %s", root, message.MerkleRoot)
	}
	return nil
}

// trailer encodes randomNumber(32) ‖ proof(32×d), which relay() reads from the
// calldata after the signatures.
func (d randomData) trailer() []byte {
	buf := make([]byte, 0, common.HashLength*(1+len(d.Proof)))
	buf = append(buf, d.Value[:]...)
	for _, node := range d.Proof {
		buf = append(buf, node[:]...)
	}
	return buf
}
