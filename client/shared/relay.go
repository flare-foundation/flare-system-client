package shared

import (
	"fmt"
	"maps"
	"slices"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

// relayCutover is the configured part of a chain's switch to a new Relay: the
// address and the first reward epoch it serves. The voting round the switch takes
// effect on is deliberately absent — a reward epoch's start can be delayed, so it
// is only fixed when that epoch's signing policy is initialized.
type relayCutover struct {
	NewAddress          common.Address
	BreakingRewardEpoch int64
}

// relayCutovers holds the scheduled Relay switch per chain id. A zero entry — or a
// chain absent from the map — means no switch: the configured Relay and the legacy
// unbound signing stay in use.
//
// TODO: fill in the deployed addresses and breaking reward epochs.
var relayCutovers = map[int64]relayCutover{
	14:  {}, // Flare
	114: {}, // Coston2
	19:  {}, // Songbird
	16:  {}, // Coston
}

// RelayCutover tracks a chain's switch to a new Relay contract.
//
// From BreakingRewardEpoch on, signing policies live in NewAddress, finalizations
// go there, and both the protocol-message digest and the signing-policy hash bind
// the chain id (the new Relay hashes keccak256(sourceChainId ‖ content)).
//
// Consumers gate on whichever quantity they hold: the finalizer knows the signing
// policy, so it asks by reward epoch; the submitter only knows the voting round, so
// it asks by round. Both name the same instant, because the round boundary is the
// breaking epoch's own startVotingRoundId — learned at runtime via
// ObserveSigningPolicy, not configured. One instance is shared by all clients.
type RelayCutover struct {
	ChainID             int64
	NewAddress          common.Address
	BreakingRewardEpoch int64

	// breakingVotingRound is startVotingRoundId of BreakingRewardEpoch, +1 so that
	// zero means "not learned yet" (round 0 is a legal value).
	breakingVotingRound atomic.Uint32
}

// NewRelayCutover returns the switch tracked for chainID. The result always carries
// ChainID, so it is the single value a component needs to decide both the digest
// form and the Relay to talk to; the rest is zero when no switch is scheduled.
func NewRelayCutover(chainID int64) *RelayCutover {
	c := relayCutovers[chainID]
	return &RelayCutover{
		ChainID:             chainID,
		NewAddress:          c.NewAddress,
		BreakingRewardEpoch: c.BreakingRewardEpoch,
	}
}

// Scheduled reports whether a switch is configured for the chain.
func (c *RelayCutover) Scheduled() bool {
	return c.NewAddress != (common.Address{}) && c.BreakingRewardEpoch > 0
}

// ObserveSigningPolicy records the round boundary once the breaking epoch's signing
// policy is seen, from its SigningPolicyInitialized event or an equivalent chain
// read. Ignores every other epoch; a later disagreeing value is refused and logged,
// since it would silently re-date the switch.
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

// BreakingVotingRound returns the first voting round served by the new Relay and
// whether it has been learned yet.
func (c *RelayCutover) BreakingVotingRound() (uint32, bool) {
	v := c.breakingVotingRound.Load()
	if v == 0 {
		return 0, false
	}
	return v - 1, true
}

// NewRelayFromRewardEpoch reports whether rewardEpochID's signing policy lives in
// the new Relay.
func (c *RelayCutover) NewRelayFromRewardEpoch(rewardEpochID int64) bool {
	return c.Scheduled() && rewardEpochID >= c.BreakingRewardEpoch
}

// NewRelayFromVotingRound reports whether votingRoundID falls in a reward epoch
// served by the new Relay, and whether that could be decided at all: with a switch
// scheduled but its round boundary not yet learned, the answer is unknown and the
// caller must fall back to the pre-switch behaviour.
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

// DigestForRewardEpoch returns the digest of a message signed under the signing
// policy of rewardEpochID.
func (c *RelayCutover) DigestForRewardEpoch(msg []byte, rewardEpochID int64) []byte {
	return MessageDigest(msg, c.ChainID, c.NewRelayFromRewardEpoch(rewardEpochID))
}

// ValidateRelayCutovers rejects half-filled table entries, which would otherwise
// silently disable the switch on that chain.
func ValidateRelayCutovers() error {
	for _, chainID := range slices.Sorted(maps.Keys(relayCutovers)) {
		c := relayCutovers[chainID]
		if c == (relayCutover{}) || (c.NewAddress != (common.Address{}) && c.BreakingRewardEpoch > 0) {
			continue
		}
		return fmt.Errorf("incomplete relay cutover for chain %d: address %s, reward epoch %d",
			chainID, c.NewAddress, c.BreakingRewardEpoch)
	}
	return nil
}
