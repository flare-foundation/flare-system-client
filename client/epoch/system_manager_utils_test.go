package epoch

import (
	"bytes"
	"testing"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// The new Relay stores keccak256(sourceChainId ‖ encoded policy) over the raw,
// unpadded bytes — the ones the SigningPolicyInitialized event carries.
func TestChainBoundSigningPolicyHash(t *testing.T) {
	policy := bytes.Repeat([]byte{0xab}, 70) // not a multiple of 32

	require.Equal(t,
		crypto.Keccak256(shared.ChainIDWord(114), policy),
		ChainBoundSigningPolicyHash(policy, 114))

	// another source chain must not produce the same hash
	require.NotEqual(t,
		ChainBoundSigningPolicyHash(policy, 14),
		ChainBoundSigningPolicyHash(policy, 114))

	// and it is not the old Relay's hash, or the switch would be undetectable
	require.NotEqual(t, SigningPolicyHash(policy), ChainBoundSigningPolicyHash(policy, 114))
}

// The old Relay's hash folds 32-byte chunks of the zero-padded policy; it must not touch
// the caller's slice — the event's policy bytes get hashed under both schemes.
func TestSigningPolicyHashDoesNotMutateInput(t *testing.T) {
	policy := make([]byte, 70, 128) // spare capacity — an append would write into it
	for i := range policy {
		policy[i] = 0xab
	}
	original := bytes.Clone(policy)

	first := SigningPolicyHash(policy)

	require.Equal(t, original, policy)
	require.Equal(t, original, policy[:70:cap(policy)][:70])
	require.Equal(t, first, SigningPolicyHash(policy))
	require.Equal(t, first, SigningPolicyHash(append(bytes.Clone(policy), make([]byte, 26)...)),
		"padding to 96 bytes must match hashing the 70-byte policy")
}
