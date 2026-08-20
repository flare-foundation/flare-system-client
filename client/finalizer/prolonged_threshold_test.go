package finalizer

import (
	"math"
	"testing"

	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/policy"
	"github.com/flare-foundation/go-flare-common/pkg/voters"
	"github.com/stretchr/testify/require"
)

// clientForThreshold stores one policy for rewardEpochID, so ForVotingRound returns it
// as the last known one and signingPolicyData takes the prolonged-epoch branch.
func clientForThreshold(t *testing.T, threshold, totalWeight, thresholdIncreaseBIPS uint16) *client {
	t.Helper()

	// one voter carrying the whole weight — only Threshold and TotalWeight matter here
	sp := &policy.SigningPolicy{
		RewardEpochID:      1,
		StartVotingRoundID: 100,
		Threshold:          threshold,
		Voters:             voters.NewSet([]common.Address{{1}}, []uint16{totalWeight}, nil),
	}
	require.Equal(t, totalWeight, sp.Voters.TotalWeight)

	storage := policy.NewStorage()
	require.NoError(t, storage.Add(sp))

	return &client{
		signingPolicyStorage: storage,
		finalizerContext: &finalizerContext{
			// epoch 1 is expected to end at round 200, so 200+ is prolonged
			rewardEpoch:           utils.NewRewardEpochConfig(0, 100),
			thresholdIncreaseBIPS: thresholdIncreaseBIPS,
		},
	}
}

// Inside the expected span the policy's own threshold is used unchanged.
func TestSigningPolicyDataUsesThePolicyThresholdBeforeTheExpectedEnd(t *testing.T) {
	c := clientForThreshold(t, 32767, 65533, 12000)

	sp, threshold := c.signingPolicyData(150)
	require.NotNil(t, sp)
	require.Equal(t, uint16(32767), threshold)
}

// Past it, the Relay's own formula: floor(policyThreshold * thresholdIncreaseBIPS / 10000).
// The replaced 60%-of-total-weight rule is one unit low for every odd total weight, which
// the contract rejects with NotEnoughWeight because both gates are strict.
func TestSigningPolicyDataRaisesTheThresholdLikeTheRelay(t *testing.T) {
	tests := []struct {
		name        string
		threshold   uint16 // ceil(W/2) at signingPolicyThresholdPPM = 500000
		totalWeight uint16
		bips        uint16
		want        uint16
		oldFormula  uint16 // floor(W * 60 / 100), for the record
	}{
		{"odd total weight, Coston today", 32767, 65533, 12000, 39320, 39319},
		{"odd total weight, Coston2 today", 32766, 65531, 12000, 39319, 39318},
		{"even total weight", 32766, 65532, 12000, 39319, 39319},
		{"raised factor", 32767, 65533, 15000, 49150, 39319},
		{"no increase", 32767, 65533, 10000, 32767, 39319},
		{"unreachable, clamped not wrapped", 32767, 65533, 40000, math.MaxUint16, 39319},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clientForThreshold(t, tt.threshold, tt.totalWeight, tt.bips)

			sp, threshold := c.signingPolicyData(250)
			require.NotNil(t, sp)
			require.Equal(t, tt.want, threshold)

			old := uint16(uint32(tt.totalWeight) * 60 / 100)
			require.Equal(t, tt.oldFormula, old, "the recorded old value must be the one the code produced")
		})
	}
}

// The threshold is scaled, never derived from total weight: two policies with the same
// total weight and different thresholds must not come out equal.
func TestSigningPolicyDataScalesTheThresholdNotTheTotalWeight(t *testing.T) {
	_, low := clientForThreshold(t, 30000, 65533, 12000).signingPolicyData(250)
	_, high := clientForThreshold(t, 32767, 65533, 12000).signingPolicyData(250)

	require.Equal(t, uint16(36000), low)
	require.Equal(t, uint16(39320), high)
}
