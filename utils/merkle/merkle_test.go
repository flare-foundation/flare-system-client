package merkle_test

import (
	"flare-tlc/utils/merkle"
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmptyTree(t *testing.T) {
	tree := merkle.New(nil)
	_, err := tree.Root()
	assert.Equal(t, err, merkle.ErrEmptyTree)

	treeSlice := tree.Tree()
	assert.Len(t, treeSlice, 0)

	numLeaves := tree.HashCount()
	assert.Equal(t, numLeaves, 0)

	sortedHashes := tree.SortedHashes()
	assert.Len(t, sortedHashes, 0)

	_, err = tree.GetHash(0)
	assert.Equal(t, err, merkle.ErrInvalidIndex)

	_, err = tree.GetProof(0)
	assert.Equal(t, err, merkle.ErrInvalidIndex)

	_, err = tree.GetProofFromHash(common.HexToHash("0x01"))
	assert.Equal(t, err, merkle.ErrHashNotFound)
}

func TestSingleLeafTree(t *testing.T) {
	val := common.HexToHash("0x01")
	tree := merkle.New([]common.Hash{val})

	root, err := tree.Root()
	require.NoError(t, err)
	assert.Equal(t, root, val)

	treeSlice := tree.Tree()
	assert.Len(t, treeSlice, 1)
	assert.Equal(t, treeSlice[0], val)

	numLeaves := tree.HashCount()
	assert.Equal(t, numLeaves, 1)

	sortedHashes := tree.SortedHashes()
	assert.Len(t, sortedHashes, 1)
	assert.Equal(t, sortedHashes[0], val)

	hash, err := tree.GetHash(0)
	require.NoError(t, err)
	assert.Equal(t, hash, val)

	proof, err := tree.GetProof(0)
	require.NoError(t, err)
	require.Len(t, proof, 0)

	verified := merkle.VerifyProof(val, proof, root)
	assert.True(t, verified)

	proof, err = tree.GetProofFromHash(val)
	require.NoError(t, err)
	require.Len(t, proof, 0)

	verified = merkle.VerifyProof(val, proof, root)
	assert.True(t, verified)
}

func TestMultiLeafTree(t *testing.T) {
	vals := []string{
		"0x01", "0x02", "0x03", "0x04", "0x05",
	}

	tree := merkle.BuildFromHex(vals, true)

	root, err := tree.Root()
	require.NoError(t, err)
	cupaloy.SnapshotT(t, root.Hex())

	t.Run("TreeSlice", func(t *testing.T) {
		treeSlice := tree.Tree()
		assert.Len(t, treeSlice, 9)
		cupaloy.SnapshotT(t, treeSlice)
	})

	numLeaves := tree.HashCount()
	assert.Equal(t, numLeaves, 5)

	sortedHashes := tree.SortedHashes()
	t.Run("SortedHashes", func(t *testing.T) {
		assert.Len(t, sortedHashes, 5)
		cupaloy.SnapshotT(t, sortedHashes)
	})

	for i := range sortedHashes {
		hash, err := tree.GetHash(i)
		require.NoError(t, err)
		assert.Equal(t, hash, sortedHashes[i])
	}

	for i, hash := range sortedHashes {
		t.Run(fmt.Sprintf("Proof_%d", i), func(t *testing.T) {
			proof, err := tree.GetProof(i)
			require.NoError(t, err)
			cupaloy.SnapshotT(t, proof)

			verified := merkle.VerifyProof(hash, proof, root)
			assert.True(t, verified)
		})

		t.Run(fmt.Sprintf("ProofFromHash_%d", i), func(t *testing.T) {
			proof, err := tree.GetProofFromHash(hash)
			require.NoError(t, err)
			cupaloy.SnapshotT(t, proof)

			verified := merkle.VerifyProof(hash, proof, root)
			assert.True(t, verified)
		})
	}
}
