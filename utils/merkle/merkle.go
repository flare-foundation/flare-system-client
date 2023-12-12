package merkle

import (
	"errors"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrEmptyTree    = errors.New("empty tree")
	ErrInvalidIndex = errors.New("invalid index")
	ErrHashNotFound = errors.New("hash not found")
)

// Tree implementation with helper functions.
type Tree struct {
	tree []common.Hash
}

// New creates a new Merkle tree from the given hash values as bytes. It is
// required that the values are sorted by their hex representation.
func New(values []common.Hash) Tree {
	return Tree{tree: values}
}

// NewFromHex creates a new Merkle tree from the given hex values. It is
// required that the values are sorted as strings.
func NewFromHex(hexValues []string) Tree {
	values := make([]common.Hash, len(hexValues))

	for i, hexValue := range hexValues {
		values[i] = common.HexToHash(hexValue)
	}

	return New(values)
}

// Given an array of leaf hashes, builds the Merkle tree.
func Build(hashes []common.Hash, initialHash bool) Tree {
	if initialHash {
		hashes = mapSingleHash(hashes)
	}

	// Hashes must be sorted to enable binary search.
	sort.Slice(hashes, func(i, j int) bool {
		return hashes[i].Hex() < hashes[j].Hex()
	})

	n := len(hashes)
	tree := make([]common.Hash, n-1, (2*n)-1)
	tree = append(tree, hashes...)

	for i := n - 2; i >= 0; i-- {
		tree[i] = SortedHashPair(tree[2*i+1], tree[2*i+2])
	}

	return New(tree)
}

// Given an array of hex-encoded leaf hashes, builds the Merkle tree.
func BuildFromHex(hexValues []string, initialHash bool) Tree {
	var hashes []common.Hash
	for i := range hexValues {
		if i == 0 || hexValues[i] != hexValues[i-1] {
			hashes = append(hashes, common.HexToHash(hexValues[i]))
		}
	}

	return Build(hashes, initialHash)
}

func mapSingleHash(hashes []common.Hash) []common.Hash {
	output := make([]common.Hash, len(hashes))

	for i := range hashes {
		output[i] = crypto.Keccak256Hash(hashes[i].Bytes())
	}

	return output
}

// SortedHashPair returns a sorted hash of two hashes.
func SortedHashPair(x, y common.Hash) common.Hash {
	if x.Hex() <= y.Hex() {
		return crypto.Keccak256Hash(x.Bytes(), y.Bytes())
	}

	return crypto.Keccak256Hash(y.Bytes(), x.Bytes())
}

// Root returns the Merkle root of the tree.
func (t Tree) Root() (common.Hash, error) {
	if len(t.tree) == 0 {
		return common.Hash{}, ErrEmptyTree
	}

	return t.tree[0], nil
}

// Tree returns the a slice representing the full tree.
func (t Tree) Tree() []common.Hash {
	return t.tree
}

// HashCount returns the number of leaves in the tree.
func (t Tree) HashCount() int {
	if len(t.tree) == 0 {
		return 0
	}

	return (len(t.tree) + 1) / 2
}

// SortedHashes returns all leaves in a slice.
func (t Tree) SortedHashes() []common.Hash {
	numLeaves := t.HashCount()
	if numLeaves == 0 {
		return nil
	}

	return t.tree[numLeaves-1:]
}

// GetHash returns the hash of the `i`th leaf.
func (t Tree) GetHash(i int) (common.Hash, error) {
	numLeaves := t.HashCount()
	if numLeaves == 0 || i < 0 || i >= numLeaves {
		return common.Hash{}, ErrInvalidIndex
	}

	pos := len(t.tree) - numLeaves + i
	return t.tree[pos], nil
}

// GetProof returns the Merkle proof for the `i`th leaf.
func (t Tree) GetProof(i int) ([]common.Hash, error) {
	numLeaves := t.HashCount()
	if numLeaves == 0 || i < 0 || i >= numLeaves {
		return nil, ErrInvalidIndex
	}

	var proof []common.Hash

	for pos := len(t.tree) - numLeaves + i; pos > 0; pos = parent(pos) {
		sibling := pos + ((2 * (pos % 2)) - 1)
		proof = append(proof, t.tree[sibling])
	}

	return proof, nil
}

// parent returns the index of the parent node.
func parent(i int) int {
	return (i - 1) / 2
}

func (t Tree) GetProofFromHash(hash common.Hash) ([]common.Hash, error) {
	i, err := t.binarySearch(hash)
	if err != nil {
		return nil, err
	}

	return t.GetProof(i)
}

func (t Tree) binarySearch(hash common.Hash) (int, error) {
	leaves := t.SortedHashes()
	i := sort.Search(len(leaves), func(i int) bool {
		return leaves[i].Hex() >= hash.Hex()
	})

	if i < len(leaves) && leaves[i] == hash {
		return i, nil
	}

	return 0, ErrHashNotFound
}

// VerifyProof verifies a Merkle proof for a given leaf.
func VerifyProof(leaf common.Hash, proof []common.Hash, root common.Hash) bool {
	hash := leaf
	for _, pair := range proof {
		hash = SortedHashPair(pair, hash)
	}

	return hash == root
}
