package finalizer

import (
	"bytes"
	"cmp"
	"flare-fsc/client/shared"
	"flare-fsc/client/shared/voters"
	"flare-fsc/logger"
	"flare-fsc/utils/contracts/relay"
	"fmt"
	"math/big"
	"sort"
	"sync"
)

// Duplicates relay.RelaySigningPolicyInitialized but with different fields and
// different types for some fields.
type signingPolicy struct {
	rewardEpochID      int64
	startVotingRoundID uint32
	threshold          uint16
	seed               *big.Int
	rawBytes           []byte
	blockTimestamp     uint64

	// The set of all voters and their weights
	voters *voters.VoterSet
}

func newSigningPolicy(r *relay.RelaySigningPolicyInitialized) *signingPolicy {
	return &signingPolicy{
		rewardEpochID:      r.RewardEpochId.Int64(),
		startVotingRoundID: r.StartVotingRoundId,
		threshold:          r.Threshold,
		seed:               r.Seed,
		rawBytes:           r.SigningPolicyBytes,
		blockTimestamp:     r.Timestamp,

		voters: voters.NewVoterSet(r.Voters, r.Weights),
	}
}

type signingPolicyStorage struct {
	// sorted list of signing policies, sorted by rewardEpochId (and also by startVotingRoundId)
	spList []*signingPolicy

	// mutex
	sync.Mutex
}

func newSigningPolicyStorage() *signingPolicyStorage {
	return &signingPolicyStorage{
		spList: make([]*signingPolicy, 0, 10),
	}
}

// Does not lock the structure, should be called from a function that does lock.
// We assume that the list is sorted by rewardEpochId and also by startVotingRoundId.
func (s *signingPolicyStorage) findByVotingRoundId(votingRoundId uint32) *signingPolicy {
	i, found := sort.Find(len(s.spList), func(i int) int {
		return cmp.Compare(votingRoundId, s.spList[i].startVotingRoundID)
	})
	if found {
		return s.spList[i]
	}
	if i == 0 {
		return nil
	}
	return s.spList[i-1]
}

func (s *signingPolicyStorage) Add(sp *signingPolicy) error {
	s.Lock()
	defer s.Unlock()

	if len(s.spList) > 0 {
		// check consistency, previous epoch should be already added
		if s.spList[len(s.spList)-1].rewardEpochID != sp.rewardEpochID-1 {
			return fmt.Errorf("missing signing policy for reward epoch id %d", sp.rewardEpochID-1)
		}
		// should be sorted by voting round id, should not happen
		if sp.startVotingRoundID < s.spList[len(s.spList)-1].startVotingRoundID {
			return fmt.Errorf("signing policy for reward epoch id %d has larger start voting round id than previous policy",
				sp.rewardEpochID)
		}
	}

	s.spList = append(s.spList, sp)
	return nil
}

// GetForVotingRound return the signingPolicy for the votingRoundID, or nil if not found.
// Also returns true if the policy is the last one or false otherwise.
func (s *signingPolicyStorage) GetForVotingRound(votingRoundID uint32) (*signingPolicy, bool) {
	s.Lock()
	defer s.Unlock()

	sp := s.findByVotingRoundId(votingRoundID)
	if sp == nil {
		return nil, false
	}
	return sp, sp.rewardEpochID == s.spList[len(s.spList)-1].rewardEpochID
}

func (s *signingPolicyStorage) First() *signingPolicy {
	s.Lock()
	defer s.Unlock()

	if len(s.spList) == 0 {
		return nil
	}
	return s.spList[0]
}

// RemoveBefore removes all signing policies that ended strictly before votingRoundID.
func (s *signingPolicyStorage) RemoveBefore(votingRoundID uint32) {
	s.Lock()
	defer s.Unlock()

	for len(s.spList) > 1 && s.spList[1].startVotingRoundID < votingRoundID {
		logger.Debug("Removing signing policy for reward epoch %d", s.spList[0].rewardEpochID)
		s.spList[0] = nil
		s.spList = s.spList[1:]
	}
}

func (s *signingPolicy) Encode() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)

	size := s.voters.Count()

	sizeBytes := shared.Uint16toBytes(uint16(size))
	epochBytes := shared.Uint32toBytes(uint32(s.rewardEpochID))
	startVotingRoundBytes := shared.Uint32toBytes(s.startVotingRoundID)
	thresholdBytes := shared.Uint16toBytes(s.threshold)

	buffer.Write(sizeBytes[:])
	buffer.Write(epochBytes[:])
	buffer.Write(startVotingRoundBytes[:])
	buffer.Write(thresholdBytes[:])
	buffer.Write(s.seed.Bytes())

	// voters and weights
	for i := 0; i < size; i++ {
		s.voters.WriteVoterRaw(buffer, i)
	}
	return buffer.Bytes(), nil
}
