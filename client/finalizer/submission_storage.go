package finalizer

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type messageData struct {
	payload       []*signedPayload
	weight        uint16
	signingPolicy *signingPolicy
}

type votingRoundKey struct {
	protocolId  byte
	messageHash common.Hash
}

type votingRoundItem struct {
	msgMap map[votingRoundKey]*messageData
}

type submissionStorage struct {
	// Map from voting round id to voting round item, a map from (protocol id, message hash) to message data
	// We use two maps instead of one to make it easier to remove a voting round
	vrMap map[uint32]*votingRoundItem

	// mutex
	sync.Mutex
}

type addPayloadResult struct {
	added            bool
	message          *messageData
	thresholdReached bool // true if the threshold was reached after adding this payload (and was below before)
}

func (m *messageData) thresholdReached() bool {
	return m.weight > m.signingPolicy.threshold
}

func newSubmissionStorage() *submissionStorage {
	return &submissionStorage{
		vrMap: make(map[uint32]*votingRoundItem),
	}
}

// Add adds a signed payload to the submission storage
// The provided signing policy must be the signing policy for the voting round
// Returns true if the payload was added, false if it was already added
func (s *submissionStorage) Add(p *signedPayload, sp *signingPolicy) (addPayloadResult, error) {
	s.Lock()
	defer s.Unlock()

	vrItem, ok := s.vrMap[p.message.votingRoundId]
	if !ok {
		vrItem = &votingRoundItem{
			msgMap: make(map[votingRoundKey]*messageData),
		}
		s.vrMap[p.message.votingRoundId] = vrItem
	}

	key := votingRoundKey{
		protocolId:  p.message.protocolId,
		messageHash: p.messageHash,
	}
	message, ok := vrItem.msgMap[key]
	if !ok {
		message = &messageData{
			payload:       make([]*signedPayload, sp.voters.Count()),
			signingPolicy: sp,
		}
		vrItem.msgMap[key] = message
	}

	voterIndex := sp.voters.VoterIndex(p.signer)
	if voterIndex < 0 {
		return addPayloadResult{}, fmt.Errorf("signer %s is not a voter", p.signer.Hex())
	}
	if message.payload[voterIndex] != nil {
		return addPayloadResult{added: false}, nil // already added
	}
	p.index = voterIndex
	thresholdAlreadyReached := message.thresholdReached()

	message.payload[voterIndex] = p
	message.weight += sp.voters.VoterWeight(voterIndex)
	return addPayloadResult{
		added:            true,
		message:          message,
		thresholdReached: !thresholdAlreadyReached && message.thresholdReached(),
	}, nil
}

func (s *submissionStorage) Get(
	votingRoundId uint32,
	protocolId byte,
	messageHash common.Hash,
) *messageData {
	s.Lock()
	defer s.Unlock()

	if vrItem, ok := s.vrMap[votingRoundId]; ok {
		key := votingRoundKey{
			protocolId:  protocolId,
			messageHash: messageHash,
		}
		if message, ok := vrItem.msgMap[key]; ok {
			return message.Copy()
		}
	}
	return nil
}

func (s *submissionStorage) RemoveVotingRoundIds(votingRoundIds []uint32) {
	s.Lock()
	defer s.Unlock()

	for _, votingRoundId := range votingRoundIds {
		delete(s.vrMap, votingRoundId)
	}
}

func (d *messageData) Copy() *messageData {
	payload := make([]*signedPayload, len(d.payload))
	copy(payload, d.payload)
	return &messageData{
		payload:       payload,
		weight:        d.weight,
		signingPolicy: d.signingPolicy,
	}
}
