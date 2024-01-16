package finalizer

import (
	"flare-tlc/utils/contracts/relay"
	"fmt"
	"math"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// Duplicates relay.RelaySigningPolicyInitialized but with fewer fields and
// different types for some fields
type signingPolicy struct {
	rewardEpochId      int64
	startVotingRoundId uint32
	threshold          uint16
	seed               *big.Int
	voters             []common.Address
	weights            []uint16
}

type voterData struct {
	index  int
	weight uint16
}

type signingPolicyStorage struct {
	minRewardEpochId int64
	maxRewardEpochId int64

	// rewardEpochId -> signingPolicy
	spMap map[int64]*signingPolicy

	// rewardEpochId -> voter -> { index, weight }
	voterMap map[int64]map[common.Address]voterData

	// mutex
	sync.Mutex
}

func newSigningPolicyStorage() *signingPolicyStorage {
	return &signingPolicyStorage{
		minRewardEpochId: math.MaxInt64,
		maxRewardEpochId: -1,
		spMap:            make(map[int64]*signingPolicy),
		voterMap:         make(map[int64]map[common.Address]voterData),
	}
}

func newSigningPolicy(r *relay.RelaySigningPolicyInitialized) *signingPolicy {
	return &signingPolicy{
		rewardEpochId:      r.RewardEpochId.Int64(),
		startVotingRoundId: r.StartVotingRoundId,
		threshold:          r.Threshold,
		seed:               r.Seed,
		voters:             r.Voters,
		weights:            r.Weights,
	}
}

func (s *signingPolicyStorage) add(sp *signingPolicy) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.spMap[sp.rewardEpochId]; ok {
		return nil // already added
	}
	if len(s.spMap) > 0 {
		// check consistency, previous epoch should be already added
		if _, ok := s.spMap[sp.rewardEpochId-1]; !ok {
			return fmt.Errorf("missing signing policy for reward epoch id %d", sp.rewardEpochId-1)
		}
	}

	s.spMap[sp.rewardEpochId] = sp
	vMap := make(map[common.Address]voterData)
	s.voterMap[sp.rewardEpochId] = vMap
	for i, voter := range sp.voters {
		if _, ok := vMap[voter]; !ok {
			vMap[voter] = voterData{
				index:  i,
				weight: sp.weights[i],
			}
		}
	}

	if sp.rewardEpochId < 0 {
		s.minRewardEpochId = sp.rewardEpochId
	}
	s.maxRewardEpochId = sp.rewardEpochId
	return nil
}
