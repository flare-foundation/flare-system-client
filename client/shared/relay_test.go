package shared

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

const (
	testChainID       = int64(114)
	testBreakingEpoch = int64(5236)
	testBreakingRound = uint32(1_000_000)
)

var testNewRelay = common.HexToAddress("0x00000000000000000000000000000000000000ff")

func scheduledCutover() *RelayCutover {
	return &RelayCutover{
		ChainID:             testChainID,
		NewAddress:          testNewRelay,
		BreakingRewardEpoch: testBreakingEpoch,
	}
}

// An unconfigured cutover — and a half-configured one, which config validation
// rejects before it can get here — must leave every gate closed: a zero breaking
// epoch would otherwise read as "already switched".
func TestUnscheduledCutoverNeverSwitches(t *testing.T) {
	for name, c := range map[string]*RelayCutover{
		"unconfigured": NewRelayCutover(testChainID, common.Address{}, 0),
		"address only": NewRelayCutover(testChainID, testNewRelay, 0),
		"epoch only":   NewRelayCutover(testChainID, common.Address{}, testBreakingEpoch),
	} {
		require.Equal(t, testChainID, c.ChainID, name)
		require.False(t, c.Scheduled(), name)
		require.False(t, c.NewRelayFromRewardEpoch(1<<40), name)

		// unscheduled is a decided answer, not an unknown one
		useNew, known := c.NewRelayFromVotingRound(1 << 31)
		require.False(t, useNew, name)
		require.True(t, known, name)

		require.Equal(t, MessageDigest([]byte("m"), testChainID, false), c.DigestForRewardEpoch([]byte("m"), 1<<40), name)
	}
}

func TestRewardEpochBoundary(t *testing.T) {
	c := scheduledCutover()

	require.False(t, c.NewRelayFromRewardEpoch(testBreakingEpoch-1))
	require.True(t, c.NewRelayFromRewardEpoch(testBreakingEpoch))
	require.True(t, c.NewRelayFromRewardEpoch(testBreakingEpoch+1))
}

// Until the breaking epoch's policy is seen the round boundary is unknown, and the
// submitter must be told so rather than being handed a false "not yet switched".
func TestVotingRoundBoundaryIsUnknownUntilObserved(t *testing.T) {
	c := scheduledCutover()

	_, ok := c.BreakingVotingRound()
	require.False(t, ok)
	useNew, known := c.NewRelayFromVotingRound(testBreakingRound + 10)
	require.False(t, useNew)
	require.False(t, known)

	// policies of other epochs say nothing about the boundary
	c.ObserveSigningPolicy(testBreakingEpoch-1, 500)
	_, ok = c.BreakingVotingRound()
	require.False(t, ok)

	c.ObserveSigningPolicy(testBreakingEpoch, testBreakingRound)

	round, ok := c.BreakingVotingRound()
	require.True(t, ok)
	require.Equal(t, testBreakingRound, round)

	for _, tc := range []struct {
		round  uint32
		useNew bool
	}{
		{testBreakingRound - 1, false},
		{testBreakingRound, true},
		{testBreakingRound + 1, true},
	} {
		useNew, known := c.NewRelayFromVotingRound(tc.round)
		require.True(t, known)
		require.Equal(t, tc.useNew, useNew, "round %d", tc.round)
	}
}

// Round 0 is a legal start round, so it must not read back as "not learned yet".
func TestObserveSigningPolicyAcceptsRoundZero(t *testing.T) {
	c := scheduledCutover()
	c.ObserveSigningPolicy(testBreakingEpoch, 0)

	round, ok := c.BreakingVotingRound()
	require.True(t, ok)
	require.Zero(t, round)

	useNew, known := c.NewRelayFromVotingRound(0)
	require.True(t, known)
	require.True(t, useNew)
}

// The boundary is fixed on chain once the policy exists; a later disagreeing
// report must not silently re-date the switch.
func TestObserveSigningPolicyKeepsTheFirstBoundary(t *testing.T) {
	c := scheduledCutover()
	c.ObserveSigningPolicy(testBreakingEpoch, testBreakingRound)
	c.ObserveSigningPolicy(testBreakingEpoch, testBreakingRound+50)

	round, _ := c.BreakingVotingRound()
	require.Equal(t, testBreakingRound, round)
}

// An unscheduled chain must ignore observations entirely.
func TestObserveSigningPolicyIgnoredWithoutSchedule(t *testing.T) {
	c := NewRelayCutover(testChainID, common.Address{}, 0)
	c.ObserveSigningPolicy(testBreakingEpoch, testBreakingRound)

	_, ok := c.BreakingVotingRound()
	require.False(t, ok)
}

// The two gates must name the same instant: the submitter signs by voting round,
// the finalizer verifies by the policy's reward epoch.
func TestDigestGatesAgreeAcrossTheBoundary(t *testing.T) {
	c := scheduledCutover()
	c.ObserveSigningPolicy(testBreakingEpoch, testBreakingRound)
	msg := []byte("protocol message")

	before, known := c.NewRelayFromVotingRound(testBreakingRound - 1)
	require.True(t, known)
	require.Equal(t, c.NewRelayFromRewardEpoch(testBreakingEpoch-1), before)

	at, known := c.NewRelayFromVotingRound(testBreakingRound)
	require.True(t, known)
	require.Equal(t, c.NewRelayFromRewardEpoch(testBreakingEpoch), at)

	require.NotEqual(t,
		MessageDigest(msg, c.ChainID, before),
		MessageDigest(msg, c.ChainID, at))
}

// The chain-bound preimage is the Relay's abi.encodePacked(uint256 sourceChainId,
// content): a 32-byte big-endian id, then the raw content.
func TestChainIDWord(t *testing.T) {
	require.Equal(t,
		"0000000000000000000000000000000000000000000000000000000000000072",
		hex.EncodeToString(ChainIDWord(114)))
	require.Equal(t,
		"000000000000000000000000000000000000000000000000000000000000000e",
		hex.EncodeToString(ChainIDWord(14)))
}

func TestMessageDigestForms(t *testing.T) {
	msg := []byte("protocol message")

	require.Equal(t,
		accounts.TextHash(crypto.Keccak256(msg)),
		MessageDigest(msg, testChainID, false))
	require.Equal(t,
		accounts.TextHash(crypto.Keccak256(ChainIDWord(testChainID), msg)),
		MessageDigest(msg, testChainID, true))

	// the binding is what makes a foreign-chain signature unusable
	require.NotEqual(t, MessageDigest(msg, 14, true), MessageDigest(msg, 114, true))
}
