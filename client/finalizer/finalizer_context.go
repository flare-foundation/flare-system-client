package finalizer

import (
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
)

// Finalizer client settings
type finalizerContext struct {
	startingRewardEpoch int64
	startingVotingRound uint32
	startTimeOffset     time.Duration // offset for fetching reward epochs at the start of the client

	voterThresholdBIPS   uint16
	gracePeriodEndOffset time.Duration

	votingRoundTiming *utils.EpochTimingConfig
	rewardEpoch       *utils.RewardEpochConfig

	// the protocol whose finalization the new Relay requires a random trailer for
	randomNumberProtocolID uint8

	// ids of the [protocol.*] sections — the only ones a local message can arrive for
	protocolIDs map[uint8]struct{}

	// factor the Relay scales a prolonged reward epoch's threshold by; >= 10000 by
	// contract (Relay.sol:327), so it can only raise it
	thresholdIncreaseBIPS uint16
}

// func newFinalizerContext(cfg *config.ClientConfig, systemsManager *system.FlareSystemsManager) (*finalizerContext, error) {
func newFinalizerContext(cfg *config.Client, relay *relay.Relay) (*finalizerContext, error) {
	votingRoundTiming, rewardEpoch, randomNumberProtocolID, thresholdIncreaseBIPS, err := shared.EpochsFromChain(relay)
	if err != nil {
		return nil, err
	}
	startingVotingRound := cfg.Finalizer.StartingVotingRound
	if startingVotingRound == 0 {
		startingVotingRound = uint32(votingRoundTiming.EpochIndex(time.Now()))
	}
	return &finalizerContext{
		startingRewardEpoch:    cfg.Finalizer.StartingRewardEpoch,
		startingVotingRound:    startingVotingRound,
		startTimeOffset:        cfg.Finalizer.StartOffset,
		voterThresholdBIPS:     cfg.Finalizer.VoterThresholdBIPS,
		gracePeriodEndOffset:   cfg.Finalizer.GracePeriodEndOffset,
		votingRoundTiming:      votingRoundTiming,
		rewardEpoch:            rewardEpoch,
		randomNumberProtocolID: randomNumberProtocolID,
		thresholdIncreaseBIPS:  thresholdIncreaseBIPS,
		protocolIDs:            configuredProtocolIDs(cfg.Protocol),
	}, nil
}

// configuredProtocolIDs collects the ids of the [protocol.*] sections.
func configuredProtocolIDs(protocols map[string]config.ProtocolConfig) map[uint8]struct{} {
	ids := make(map[uint8]struct{}, len(protocols))
	for _, protocol := range protocols {
		ids[protocol.ID] = struct{}{}
	}
	return ids
}

// serves reports whether the submitter queries protocolID; an empty set serves nothing.
func (fc *finalizerContext) serves(protocolID uint8) bool {
	_, ok := fc.protocolIDs[protocolID]
	return ok
}
