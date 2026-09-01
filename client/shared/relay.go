package shared

import (
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

// RelayCutover tracks a chain's switch to a new Relay contract; one shared instance.
// From BreakingRewardEpoch on, policies and finalizations live at NewAddress, and the
// message digest and policy hash bind the chain id — keccak256(chainID ‖ content).
// The round boundary is that epoch's startVotingRoundId, learned by ObserveSigningPolicy.
type RelayCutover struct {
	ChainID             int64
	NewAddress          common.Address
	BreakingRewardEpoch int64

	// startVotingRoundId of BreakingRewardEpoch, +1 — zero means "not learned" (round 0 is legal)
	breakingVotingRound atomic.Uint32
}

// NewRelayCutover schedules the switch to newAddress from startingRewardEpoch on, both
// [relay_cutover] config values; zero means no switch. The round it takes effect on is learned
// later because a reward epoch's start can be delayed — see ObserveSigningPolicy.
func NewRelayCutover(chainID int64, newAddress common.Address, startingRewardEpoch int64) *RelayCutover {
	return &RelayCutover{
		ChainID:             chainID,
		NewAddress:          newAddress,
		BreakingRewardEpoch: startingRewardEpoch,
	}
}

// Scheduled reports whether a switch is configured for the chain.
func (c *RelayCutover) Scheduled() bool {
	return c.NewAddress != (common.Address{}) && c.BreakingRewardEpoch > 0
}

// ObserveSigningPolicy records the round boundary from BreakingRewardEpoch's
// SigningPolicyInitialized event or an equivalent chain read. First value wins: a later
// disagreeing one would silently re-date the switch, so it is refused and logged.
func (c *RelayCutover) ObserveSigningPolicy(rewardEpochID int64, startVotingRoundID uint32) {
	if !c.Scheduled() || rewardEpochID != c.BreakingRewardEpoch {
		return
	}
	if c.breakingVotingRound.CompareAndSwap(0, startVotingRoundID+1) {
		logger.Infof("Relay cutover: reward epoch %d starts at voting round %d, switching to Relay %s from there",
			c.BreakingRewardEpoch, startVotingRoundID, c.NewAddress)
		return
	}
	if known := c.breakingVotingRound.Load() - 1; known != startVotingRoundID {
		logger.Errorf("Relay cutover: reward epoch %d reported starting at voting round %d, keeping %d",
			c.BreakingRewardEpoch, startVotingRoundID, known)
	}
}

// BreakingVotingRound returns the first round the new Relay serves and whether it is known.
func (c *RelayCutover) BreakingVotingRound() (uint32, bool) {
	v := c.breakingVotingRound.Load()
	if v == 0 {
		return 0, false
	}
	return v - 1, true
}

// NewRelayFromRewardEpoch reports whether the new Relay holds rewardEpochID's signing policy.
func (c *RelayCutover) NewRelayFromRewardEpoch(rewardEpochID int64) bool {
	return c.Scheduled() && rewardEpochID >= c.BreakingRewardEpoch
}

// NewRelayFromVotingRound reports whether votingRoundID falls in a reward epoch served by
// the new Relay. known is false while a scheduled switch has no learned boundary; the caller
// must then fall back to the pre-switch behaviour.
func (c *RelayCutover) NewRelayFromVotingRound(votingRoundID uint32) (useNew, known bool) {
	if !c.Scheduled() {
		return false, true
	}
	breaking, ok := c.BreakingVotingRound()
	if !ok {
		return false, false
	}
	return votingRoundID >= breaking, true
}

// DigestForRewardEpoch returns the digest of msg signed under rewardEpochID's signing policy.
func (c *RelayCutover) DigestForRewardEpoch(msg []byte, rewardEpochID int64) []byte {
	return MessageDigest(msg, c.ChainID, c.NewRelayFromRewardEpoch(rewardEpochID))
}

// DigestFromMessage derives the digest as the Relay does, from the voting round in msg, so
// components hashing the same bytes agree without sharing context. known is false while a
// scheduled switch has no learned boundary: the digest is then the pre-switch form, and a
// caller that knows the governing policy should prefer DigestForRewardEpoch.
func (c *RelayCutover) DigestFromMessage(msg Message) (digest []byte, known bool, err error) {
	m, err := msg.Parse()
	if err != nil {
		return nil, false, err
	}
	useNew, known := c.NewRelayFromVotingRound(m.VotingRoundID)
	return MessageDigest(msg, c.ChainID, useNew), known, nil
}
