package finalizer

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type messageData struct {
	payload          []*signedPayload
	weight           uint16
	signingPolicy    *signingPolicy
	thresholdReached bool
}

type MessageThresholdProvider interface {
	MessageThreshold(votingRound uint16) bool
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
	message          *messageData
	thresholdReached bool // true if the threshold was reached after adding this payload (and was below before)
}

func newMessageData(sp *signingPolicy) *messageData {
	return &messageData{
		payload:       make([]*signedPayload, sp.voters.Count()),
		signingPolicy: sp,
	}
}

func (m *messageData) addPayload(p *signedPayload, threshold uint16) error {
	voterIndex := m.signingPolicy.voters.VoterIndex(p.signer)
	if voterIndex < 0 {
		return fmt.Errorf("signer %s is not a registered voter in the current reward epoch", p.signer.Hex())
	}
	if m.payload[voterIndex] != nil {
		return nil // already added
	}
	p.index = voterIndex

	m.payload[voterIndex] = p
	m.weight += m.signingPolicy.voters.VoterWeight(voterIndex)
	if !m.thresholdReached {
		m.thresholdReached = m.weight > threshold
	}
	return nil
}

func newSubmissionStorage() *submissionStorage {
	return &submissionStorage{
		vrMap: make(map[uint32]*votingRoundItem),
	}
}

// Add adds a signed payload to the submission storage
// The provided signing policy must be the signing policy for the voting round
// Returns true if the payload was added, false if it was already added
func (s *submissionStorage) Add(p *signedPayload, sp *signingPolicy, threshold uint16) (addPayloadResult, error) {
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
		message = newMessageData(sp)
		vrItem.msgMap[key] = message
	}

	// Add the payload to the message
	thresholdAlreadyReached := message.thresholdReached
	err := message.addPayload(p, threshold)
	if err != nil {
		return addPayloadResult{}, err
	}
	return addPayloadResult{
		message:          message,
		thresholdReached: !thresholdAlreadyReached && message.thresholdReached,
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
