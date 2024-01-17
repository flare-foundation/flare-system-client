package finalizer

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/exp/slices"
)

type messageData struct {
	payload []*signedPayload
	weight  uint16
}

type protocolData struct {
	ProtocolId byte
	messageMap map[common.Hash]*messageData
}

type votingRoundData struct {
	votingRoundId uint32
	protocolMap   map[byte]*protocolData
}

type submissionStorage struct {
	// votingRoundId -> votingRound
	vrMap map[uint32]*votingRoundData

	// mutex
	sync.Mutex
}

type addPayloadResult struct {
	added            bool
	thresholdReached bool
}

func newSubmissionStorage() *submissionStorage {
	return &submissionStorage{
		vrMap: make(map[uint32]*votingRoundData),
	}
}

// Add adds a signed payload to the submission storage
// The provided signing policy must be the signing policy for the voting round
// Returns true if the payload was added, false if it was already added
func (s *submissionStorage) Add(p *signedPayload, sp *signingPolicy) (addPayloadResult, error) {
	s.Lock()
	defer s.Unlock()

	vr, ok := s.vrMap[p.message.votingRoundId]
	if !ok {
		vr = &votingRoundData{
			votingRoundId: p.message.votingRoundId,
			protocolMap:   make(map[byte]*protocolData),
		}
		s.vrMap[p.message.votingRoundId] = vr
	}

	protocol, ok := vr.protocolMap[p.message.protocolId]
	if !ok {
		protocol = &protocolData{
			ProtocolId: p.message.protocolId,
			messageMap: make(map[common.Hash]*messageData),
		}
		vr.protocolMap[p.message.protocolId] = protocol
	}

	message, ok := protocol.messageMap[p.messageHash]
	if !ok {
		message = &messageData{
			payload: make([]*signedPayload, len(sp.voters)),
		}
		protocol.messageMap[p.messageHash] = message
	}
	voterIndex := slices.Index(sp.voters, p.signer)
	if voterIndex < 0 {
		return addPayloadResult{}, fmt.Errorf("signer %s is not a voter", p.signer.Hex())
	}
	if message.payload[voterIndex] != nil {
		return addPayloadResult{added: false}, nil // already added
	}
	message.payload[voterIndex] = p
	message.weight += sp.weights[voterIndex]
	return addPayloadResult{
		added:            true,
		thresholdReached: message.weight > sp.threshold,
	}, nil
}
