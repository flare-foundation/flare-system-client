package finalizer

import (
	"bytes"
	"flare-tlc/client/shared"
	"flare-tlc/utils/contracts/relay"
	"fmt"
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
	rawBytes           []byte
}

func newSigningPolicy(r *relay.RelaySigningPolicyInitialized) *signingPolicy {
	return &signingPolicy{
		rewardEpochId:      r.RewardEpochId.Int64(),
		startVotingRoundId: r.StartVotingRoundId,
		threshold:          r.Threshold,
		seed:               r.Seed,
		voters:             r.Voters,
		weights:            r.Weights,
		rawBytes:           r.SigningPolicyBytes,
	}
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
		minRewardEpochId: -1,
		maxRewardEpochId: -1,
		spMap:            make(map[int64]*signingPolicy),
		voterMap:         make(map[int64]map[common.Address]voterData),
	}
}

func (s *signingPolicyStorage) Add(sp *signingPolicy) error {
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

	if s.minRewardEpochId < 0 {
		s.minRewardEpochId = sp.rewardEpochId
	}
	s.maxRewardEpochId = sp.rewardEpochId
	return nil
}

func (s *signingPolicyStorage) GetForVotingRound(votingRoundId uint32) *signingPolicy {
	s.Lock()
	defer s.Unlock()

	if len(s.spMap) == 0 {
		return nil
	}
	for i := s.maxRewardEpochId; i >= s.minRewardEpochId; i-- {
		sp := s.spMap[i]
		if sp.startVotingRoundId <= votingRoundId {
			return sp
		}
	}
	return nil
}

func (s *signingPolicy) Encode() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)

	size := len(s.voters)
	// TODO: size, epoch, voting round size checks?

	sizeBytes := shared.Uint16toBytes(uint16(size))
	epochBytes := shared.Uint32toBytes(uint32(s.rewardEpochId))
	startVotingRoundBytes := shared.Uint32toBytes(s.startVotingRoundId)
	thresholdBytes := shared.Uint16toBytes(s.threshold)

	buffer.Write(sizeBytes[:])
	buffer.Write(epochBytes[:])
	buffer.Write(startVotingRoundBytes[:])
	buffer.Write(thresholdBytes[:])
	buffer.Write(s.seed.Bytes())

	// voters and weights
	for i := 0; i < size; i++ {
		weightBytes := shared.Uint16toBytes(s.weights[i])
		buffer.Write(s.voters[i].Bytes())
		buffer.Write(weightBytes[:])
	}
	return buffer.Bytes(), nil
}
