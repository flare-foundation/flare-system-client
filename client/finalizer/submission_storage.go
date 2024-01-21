package finalizer

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/exp/slices"
)

type messageData struct {
	payload       []*signedPayload
	weight        uint16
	signingPolicy *signingPolicy
}

type storageItemKey struct {
	votingRoundId uint32
	protocolId    byte
	messageHash   common.Hash
}

type submissionStorage struct {
	ssMap map[storageItemKey]*messageData

	// mutex
	sync.Mutex
}

type addPayloadResult struct {
	added   bool
	message *messageData
}

func (m *messageData) thresholdReached() bool {
	return m.weight > m.signingPolicy.threshold
}

func newSubmissionStorage() *submissionStorage {
	return &submissionStorage{
		ssMap: make(map[storageItemKey]*messageData),
	}
}

// Add adds a signed payload to the submission storage
// The provided signing policy must be the signing policy for the voting round
// Returns true if the payload was added, false if it was already added
func (s *submissionStorage) Add(p *signedPayload, sp *signingPolicy) (addPayloadResult, error) {
	s.Lock()
	defer s.Unlock()

	key := storageItemKey{
		votingRoundId: p.message.votingRoundId,
		protocolId:    p.message.protocolId,
		messageHash:   p.messageHash,
	}
	message, ok := s.ssMap[key]
	if !ok {
		message = &messageData{
			payload: make([]*signedPayload, len(sp.voters)),
		}
		s.ssMap[key] = message
	}

	voterIndex := slices.Index(sp.voters, p.signer)
	if voterIndex < 0 {
		return addPayloadResult{}, fmt.Errorf("signer %s is not a voter", p.signer.Hex())
	}
	if message.payload[voterIndex] != nil {
		return addPayloadResult{added: false}, nil // already added
	}
	p.index = voterIndex
	message.payload[voterIndex] = p
	message.weight += sp.weights[voterIndex]
	message.signingPolicy = sp
	return addPayloadResult{
		added:   true,
		message: message,
	}, nil
}

func (s *submissionStorage) Get(
	votingRoundId uint32,
	protocolId byte,
	messageHash common.Hash,
) *messageData {
	s.Lock()
	defer s.Unlock()

	key := storageItemKey{
		votingRoundId: votingRoundId,
		protocolId:    protocolId,
		messageHash:   messageHash,
	}
	if message, ok := s.ssMap[key]; ok {
		return message.Copy()
	}
	return nil
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
