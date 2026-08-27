package protocol

import (
	"crypto/ecdsa"
	"encoding/binary"
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

func testMessage(round uint32) []byte {
	msg := make([]byte, shared.RelayMessageLength)
	msg[0] = 1
	binary.BigEndian.PutUint32(msg[1:5], round)
	return msg
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

// A submitSignatures signature must recover under the digest of the Relay finalizing the
// round, read from the message bytes as the contract reads it; that Relay changes at the
// breaking epoch's first round.
func TestSignSignaturePayloadFollowsObservedBoundary(t *testing.T) {
	key, signer := signingTestKey(t)
	cutover := scheduledCutover(t, true)

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
			data := testMessage(c.round)
			signature, err := SignSignaturePayload(cutover, data, key)
			require.NoError(t, err)
			requireSignedUnder(t, signature, data, c.chainBound, signer)
		})
	}
}

// An unlearned boundary signs the pre-switch way — the only form derivable here, not
// necessarily the right one; see TestFinalizerRecoversSignersUnderThePolicyEpochDigest.
func TestSignSignaturePayloadBeforeBoundaryIsKnown(t *testing.T) {
	key, signer := signingTestKey(t)
	cutover := scheduledCutover(t, false)
	data := testMessage(testBreakingRound + 10)

	signature, err := SignSignaturePayload(cutover, data, key)
	require.NoError(t, err)
	requireSignedUnder(t, signature, data, false, signer)

	// once observed, the same round signs the new way
	cutover.ObserveSigningPolicy(testBreakingEpoch, testBreakingRound)
	signature, err = SignSignaturePayload(cutover, data, key)
	require.NoError(t, err)
	requireSignedUnder(t, signature, data, true, signer)
}

// With no cutover scheduled the payload keeps the legacy digest whatever the round.
func TestSignSignaturePayloadWithoutCutover(t *testing.T) {
	key, signer := signingTestKey(t)
	data := testMessage(1 << 31)

	signature, err := SignSignaturePayload(shared.NewRelayCutover(testChainID, common.Address{}, 0), data, key)
	require.NoError(t, err)
	requireSignedUnder(t, signature, data, false, signer)
}

// Data that is not a protocol message names no round, so no Relay to sign for.
func TestSignSignaturePayloadRejectsMalformedData(t *testing.T) {
	key, _ := signingTestKey(t)

	_, err := SignSignaturePayload(scheduledCutover(t, true), []byte("not a protocol message"), key)
	require.Error(t, err)
}
