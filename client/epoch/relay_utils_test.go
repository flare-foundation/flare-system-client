package epoch

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	"github.com/flare-foundation/go-flare-common/pkg/database"
)

const relayTestChainID = int64(114)

func newRelayImplForTest(t *testing.T, cutover *shared.RelayCutover) *relayContractClientImpl {
	t.Helper()
	relayContract, err := relay.NewRelay(common.Address{}, nil)
	require.NoError(t, err)
	return &relayContractClientImpl{relay: relayContract, relayCutover: cutover}
}

// spiLog builds a SigningPolicyInitialized log as the indexer stores it: epoch id in
// topic1, the rest ABI-packed.
func spiLog(t *testing.T, rewardEpochID int64, startVotingRoundID uint32) database.Log {
	t.Helper()

	relayABI, err := relay.RelayMetaData.GetAbi()
	require.NoError(t, err)
	event := relayABI.Events["SigningPolicyInitialized"]

	data, err := event.Inputs.NonIndexed().Pack(
		startVotingRoundID,
		uint16(1),
		big.NewInt(1),
		[]common.Address{{}},
		[]uint16{1},
		[]byte{0x01},
		uint64(rewardEpochID),
	)
	require.NoError(t, err)

	return database.Log{
		Data:   hex.EncodeToString(data),
		Topic0: event.ID.Hex(),
		Topic1: common.BigToHash(big.NewInt(rewardEpochID)).Hex(),
		Topic2: "NULL",
		Topic3: "NULL",
	}
}

// logs of two Relays merge with no order by emitter, so the policy is picked by reward epoch
func TestSelectPolicyPrefersTheAnticipatedEpoch(t *testing.T) {
	r := newRelayImplForTest(t, shared.NewRelayCutover(relayTestChainID, common.Address{}, 0))

	logs := []database.Log{spiLog(t, 12, 120), spiLog(t, 10, 100), spiLog(t, 11, 110)}

	selected := r.selectPolicy(logs, 11)
	require.NotNil(t, selected)
	require.Equal(t, int64(11), selected.RewardEpochId.Int64())
	require.Equal(t, uint32(110), selected.StartVotingRoundId)
}

// a delayed epoch leaves the time-derived anticipation ahead of the chain, so take the highest
func TestSelectPolicyFallsBackToHighestEpoch(t *testing.T) {
	r := newRelayImplForTest(t, shared.NewRelayCutover(relayTestChainID, common.Address{}, 0))

	logs := []database.Log{spiLog(t, 10, 100), spiLog(t, 11, 110)}

	selected := r.selectPolicy(logs, 12)
	require.NotNil(t, selected)
	require.Equal(t, int64(11), selected.RewardEpochId.Int64())

	require.Nil(t, r.selectPolicy(nil, 12))
}

// the listener is where nodes learn the cutover round: the breaking epoch's policy carries it
func TestSelectPolicyDatesTheRelayCutover(t *testing.T) {
	cutover := &shared.RelayCutover{
		ChainID:             relayTestChainID,
		NewAddress:          common.HexToAddress("0x00000000000000000000000000000000000000ff"),
		BreakingRewardEpoch: 11,
	}
	r := newRelayImplForTest(t, cutover)

	// a policy of another epoch says nothing
	r.selectPolicy([]database.Log{spiLog(t, 10, 100)}, 10)
	_, known := cutover.BreakingVotingRound()
	require.False(t, known)

	// the breaking epoch's policy dates the switch even when it is not the one selected
	r.selectPolicy([]database.Log{spiLog(t, 11, 110), spiLog(t, 12, 120)}, 12)

	round, known := cutover.BreakingVotingRound()
	require.True(t, known)
	require.Equal(t, uint32(110), round)
}
