package finalizer

import (
	"flare-fsc/client/shared"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"

	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/logger"
)

type signaturesCollection struct {
	message          shared.Message
	signatures       [][]byte
	weight           uint16
	thresholdReached bool
	signingPolicy    *signingPolicy
	threshold        uint16
}

type protocolCollection struct {
	messageAdded        bool
	messageChosenHash   common.Hash
	signatureCollection map[common.Hash]*signaturesCollection
	unprocessedPayloads []*submitSignaturesPayload
	signingPolicy       *signingPolicy
	threshold           uint16
}

// roundCollection maps protocolID to signatureCollection
type roundCollection struct {
	protocolCollections map[uint8]*protocolCollection
}

type finalizationStorage struct {
	stg               map[uint32]*roundCollection // map from roundID to roundCollection
	lowestRoundStored uint32

	// mutex
	sync.RWMutex
}

type FinalizationReady struct {
	thresholdReached bool
	protocolID       uint8
	votingRoundID    uint32
	msgHash          common.Hash
}

func NewSignatureCollection(message shared.Message, signingPolicy *signingPolicy, threshold uint16) *signaturesCollection {
	return &signaturesCollection{
		message:       message,
		signatures:    make([][]byte, signingPolicy.voters.Count()),
		signingPolicy: signingPolicy,
		threshold:     threshold,
	}
}

func (sc *signaturesCollection) addSignature(p *submitSignaturesPayload) (bool, error) {
	if p.voterIndex < 0 {
		return false, fmt.Errorf("voter not recognized")
	}

	if len(sc.signatures[p.voterIndex]) != 0 {
		return false, fmt.Errorf("signature for signer %d with address %s already added", p.voterIndex, p.signer)
	}

	sc.signatures[p.voterIndex] = p.signature

	sc.weight += p.weight
	if !sc.thresholdReached {
		sc.thresholdReached = sc.weight > sc.threshold

		return sc.thresholdReached, nil
	}
	return false, nil
}

func (pc *protocolCollection) addMessage(message shared.Message) (bool, common.Hash, error) {
	if pc.messageAdded {
		return false, common.Hash{}, fmt.Errorf("message added twice")
	}

	msgHsh := common.Hash(message.Hash())
	_, exists := pc.signatureCollection[msgHsh]
	if !exists {
		pc.signatureCollection[msgHsh] = NewSignatureCollection(message, pc.signingPolicy, pc.threshold)
		pc.messageChosenHash = msgHsh
	}

	pc.signatureCollection[msgHsh].message = message
	pc.messageAdded = true

	thresholdReached := false

	for _, up := range pc.unprocessedPayloads {
		tr, msgHashCheck, err := pc.addPayload(up)
		if err != nil {
			logger.Error("Adding payload after message error:", err)
		}
		if msgHashCheck != msgHsh {
			logger.Debug("Unexpected behavior, hashes should match")
		}
		if tr {
			thresholdReached = true
		}
	}

	//clear unprocessedPayloads
	pc.unprocessedPayloads = nil

	return thresholdReached, msgHsh, nil
}

func (pc *protocolCollection) addPayload(payload *submitSignaturesPayload) (bool, common.Hash, error) {
	if !pc.messageAdded && payload.typeID != 0 {
		pc.unprocessedPayloads = append(pc.unprocessedPayloads, payload)

		return false, common.Hash{}, nil
	}

	var msgHash []byte
	var sigCollection *signaturesCollection
	if payload.typeID == 0 {
		msgHash = payload.message.Hash()
		_, exists := pc.signatureCollection[common.Hash(msgHash)]
		if !exists {
			pc.signatureCollection[common.Hash(msgHash)] = NewSignatureCollection(payload.message, pc.signingPolicy, pc.threshold)
		}
		sigCollection = pc.signatureCollection[common.Hash(msgHash)]
	} else if pc.messageAdded {
		sigCollection = pc.signatureCollection[pc.messageChosenHash]
		msgHash = sigCollection.message.Hash()
	} else {
		return false, common.Hash{}, fmt.Errorf("unexpected behavior, no message")
	}

	err := payload.AddSigner(msgHash, sigCollection.signingPolicy.voters)
	if err != nil {
		return false, common.Hash{}, fmt.Errorf("adding payload, %v", err)
	}

	thresholdReached, err := sigCollection.addSignature(payload)

	return thresholdReached, common.Hash(msgHash), err
}

func newFinalizationStorage() *finalizationStorage {
	return &finalizationStorage{
		stg: make(map[uint32]*roundCollection),
	}
}

// addPayload adds a submitSignature payload to the finalizationStorage.
// The payload is added to the protocolCollection for the protocolID and roundID of the payload.
// An indicator whether the addition has made the protocol reach the threshold for the round is returned.
func (s *finalizationStorage) addPayload(p *submitSignaturesPayload, signingPolicy *signingPolicy, threshold uint16) (FinalizationReady, error) {
	s.Lock()
	defer s.Unlock()

	if p.votingRoundID < s.lowestRoundStored {
		return FinalizationReady{thresholdReached: false}, fmt.Errorf("payload for an round before lowestRoundStored %d", s.lowestRoundStored)
	}

	rc, exists := s.stg[p.votingRoundID]
	if !exists {
		rc = &roundCollection{protocolCollections: make(map[uint8]*protocolCollection)}

		s.stg[p.votingRoundID] = rc
	}

	pc, exists := rc.protocolCollections[p.protocolID]
	if !exists {
		pc = &protocolCollection{signingPolicy: signingPolicy, signatureCollection: make(map[common.Hash]*signaturesCollection), threshold: threshold}
		rc.protocolCollections[p.protocolID] = pc
	}

	thresholdReached, msgHash, err := pc.addPayload(p)
	if err != nil {
		return FinalizationReady{thresholdReached: false}, err
	}
	if thresholdReached {
		return FinalizationReady{thresholdReached: true, protocolID: p.protocolID, votingRoundID: p.votingRoundID, msgHash: common.Hash(msgHash)}, nil
	}

	return FinalizationReady{thresholdReached: false}, nil
}

// AddMessage adds a protocol message to the finalizationStorage for the respective protocol and round, and adds all unprocessedPayloads for the respective round and protocol.
// An indicator whether the additions have made the protocol reach the threshold for the round is returned.
func (s *finalizationStorage) AddMessage(p *shared.ProtocolMessage, signingPolicy *signingPolicy, threshold uint16) (FinalizationReady, error) {
	s.Lock()
	defer s.Unlock()

	if p.VotingRoundID < s.lowestRoundStored {
		return FinalizationReady{thresholdReached: false}, nil //TODO
	}

	rc, exists := s.stg[p.VotingRoundID]
	if !exists {
		rc = &roundCollection{protocolCollections: make(map[uint8]*protocolCollection)}
		s.stg[p.VotingRoundID] = rc
	}

	pc, exists := rc.protocolCollections[p.ProtocolID]
	if !exists {
		pc = &protocolCollection{signatureCollection: make(map[common.Hash]*signaturesCollection), signingPolicy: signingPolicy, threshold: threshold}
		rc.protocolCollections[p.ProtocolID] = pc
	}

	thresholdReached, msgHash, err := pc.addMessage(p.Message)
	if err != nil {
		return FinalizationReady{thresholdReached: false}, err
	}
	if thresholdReached {
		return FinalizationReady{thresholdReached: true, protocolID: p.ProtocolID, votingRoundID: p.VotingRoundID, msgHash: msgHash}, nil
	}

	return FinalizationReady{thresholdReached: false}, nil
}

// Get returns the signatureCollection for votingRoundID and protocolID.
// A boolean inductor of existence is also returned.
func (fs *finalizationStorage) Get(votingRoundID uint32, protocolID uint8, msgHash common.Hash) (*signaturesCollection, bool) {
	fs.RLock()
	defer fs.RUnlock()
	round, exists := fs.stg[votingRoundID]
	if !exists {
		return &signaturesCollection{}, false
	}

	pc, exists := round.protocolCollections[protocolID]
	if !exists {
		return &signaturesCollection{}, false
	}

	sigCollection, exists := pc.signatureCollection[msgHash]
	if !exists {
		return &signaturesCollection{}, false
	}

	return sigCollection, true
}

// RemoveRoundsBefore deletes rounds before votingRoundID.
func (fs *finalizationStorage) RemoveRoundsBefore(votingRoundID uint32) {
	// initial cleanup
	if fs.lowestRoundStored == 0 && votingRoundID > 20 {
		fs.lowestRoundStored = votingRoundID - 20
	}

	if votingRoundID > fs.lowestRoundStored {
		fs.Lock()
		defer fs.Unlock()

		for i := fs.lowestRoundStored; i < votingRoundID; i++ {
			logger.Infof("Deleting round %d in finalization storage", i)
			delete(fs.stg, i)
		}

		fs.lowestRoundStored = votingRoundID + 1
	}
}
