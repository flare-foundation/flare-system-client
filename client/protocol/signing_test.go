package protocol

import (
	"crypto/ecdsa"
	"testing"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

const (
	testChainID       = int64(114)
	testBreakingEpoch = int64(5236)
	testBreakingRound = uint32(1_000_000)
)

func signingTestKey(t *testing.T) (*ecdsa.PrivateKey, common.Address) {
	t.Helper()
	key, err := crypto.HexToECDSA("4f65bffe3c8ed6c0b812e84d35402e949feea042061cc1635fe6ae83ed84df4a")
	require.NoError(t, err)
	return key, crypto.PubkeyToAddress(key.PublicKey)
}

func scheduledCutover(t *testing.T, observed bool) *shared.RelayCutover {
	t.Helper()
	c := &shared.RelayCutover{
		ChainID:             testChainID,
		NewAddress:          common.HexToAddress("0x00000000000000000000000000000000000000ff"),
		BreakingRewardEpoch: testBreakingEpoch,
	}
	if observed {
		c.ObserveSigningPolicy(testBreakingEpoch, testBreakingRound)
	}
	return c
}

func requireSignedUnder(t *testing.T, signature, data []byte, chainBound bool, signer common.Address) {
	t.Helper()

	pub, err := crypto.SigToPub(shared.MessageDigest(data, testChainID, chainBound), signature)
	require.NoError(t, err)
	require.Equal(t, signer, crypto.PubkeyToAddress(*pub))

	// the other form must not verify, or the cutover would be a no-op
	pub, err = crypto.SigToPub(shared.MessageDigest(data, testChainID, !chainBound), signature)
	require.NoError(t, err)
	require.NotEqual(t, signer, crypto.PubkeyToAddress(*pub))
}

// A submitSignatures signature must recover to the signer under the digest of the
// Relay that will finalize the round, and that Relay changes at the round the
// breaking epoch's signing policy starts on.
func TestSignSignaturePayloadFollowsObservedBoundary(t *testing.T) {
	key, signer := signingTestKey(t)
	cutover := scheduledCutover(t, true)
	data := make([]byte, 38)

	cases := []struct {
		name       string
		round      uint32
		chainBound bool
	}{
		{"round before the boundary", testBreakingRound - 1, false},
		{"first round of the breaking epoch", testBreakingRound, true},
		{"round after the boundary", testBreakingRound + 1, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			signature, err := SignSignaturePayload(cutover, c.round, data, key)
			require.NoError(t, err)
			requireSignedUnder(t, signature, data, c.chainBound, signer)
		})
	}
}

// Before the breaking epoch's policy is seen the boundary is unknown. The round
// cannot be past it yet, so the old digest is the correct answer — and the only
// one peers would accept.
func TestSignSignaturePayloadBeforeBoundaryIsKnown(t *testing.T) {
	key, signer := signingTestKey(t)
	cutover := scheduledCutover(t, false)
	data := make([]byte, 38)

	signature, err := SignSignaturePayload(cutover, testBreakingRound+10, data, key)
	require.NoError(t, err)
	requireSignedUnder(t, signature, data, false, signer)

	// once observed, the same round signs the new way
	cutover.ObserveSigningPolicy(testBreakingEpoch, testBreakingRound)
	signature, err = SignSignaturePayload(cutover, testBreakingRound+10, data, key)
	require.NoError(t, err)
	requireSignedUnder(t, signature, data, true, signer)
}

// With no cutover scheduled the payload keeps the legacy digest whatever the round.
func TestSignSignaturePayloadWithoutCutover(t *testing.T) {
	key, signer := signingTestKey(t)
	data := make([]byte, 38)

	signature, err := SignSignaturePayload(shared.NewRelayCutover(testChainID, common.Address{}, 0), 1<<31, data, key)
	require.NoError(t, err)
	requireSignedUnder(t, signature, data, false, signer)
}
