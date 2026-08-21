package finalizer

import (
	"errors"
	"fmt"
	"sync"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/flare-foundation/go-flare-common/pkg/policy"
)

type signaturesCollection struct {
	message          shared.Message
	signatures       [][]byte // expected to be of length len(voters), i-th entry is the signature of i-th voter
	weight           uint16
	thresholdReached bool
	signingPolicy    *policy.SigningPolicy
	threshold        uint16

	mu sync.RWMutex
}

type protocolCollection struct {
	messageAdded        bool
	messageChosenDigest common.Hash
	signatureCollection map[common.Hash]*signaturesCollection
	unprocessedPayloads []*submitSignaturesPayload
	bufferedSenders     map[common.Address]struct{} // senders already buffered pre-message (DOS-01 cap)
	signingPolicy       *policy.SigningPolicy
	threshold           uint16
	relayCutover        *shared.RelayCutover
}

// messageDigest returns the digest the collection's signatures cover, derived from
// the round embedded in the message bytes — the contract's own derivation, and the
// one every signer of these bytes used. Fallback: with the boundary unlearned (a
// restart can fetch only post-breaking policies) or the bytes unparseable, the
// governing policy's epoch decides, as it does for the target Relay.
func (pc *protocolCollection) messageDigest(message shared.Message) []byte {
	if digest, known, err := pc.relayCutover.DigestFromMessage(message); err == nil && known {
		return digest
	}
	return pc.relayCutover.DigestForRewardEpoch(message, pc.signingPolicy.RewardEpochID)
}

// roundCollection maps protocolID to protocolCollection
type roundCollection struct {
	protocolCollections map[uint8]*protocolCollection
}

type finalizationStorage struct {
	stg               map[uint32]*roundCollection // map from roundID to roundCollection
	lowestRoundStored uint32
	relayCutover      *shared.RelayCutover

	// mutex
	sync.RWMutex
}

type FinalizationReady struct {
	thresholdReached bool
	protocolID       uint8
	votingRoundID    uint32
	digest           common.Hash
}

func NewSignatureCollection(message shared.Message, signingPolicy *policy.SigningPolicy, threshold uint16) *signaturesCollection {
	return &signaturesCollection{
		message:       message,
		signatures:    make([][]byte, signingPolicy.Voters.Count()),
		signingPolicy: signingPolicy,
		threshold:     threshold,
	}
}

// addSignature adds signature to the signatures collection.
func (sc *signaturesCollection) addSignature(p *submitSignaturesPayload) (bool, error) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if p.voterIndex < 0 {
		return false, errors.New("voter not recognized")
	}

	if len(sc.signatures[p.voterIndex]) != 0 {
		return false, fmt.Errorf("%w: signature for signer %d with address %s already added", errBadPayload, p.voterIndex, p.signer)
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
		return false, common.Hash{}, errors.New("message added twice")
	}

	digest := common.Hash(pc.messageDigest(message))
	// An existing collection at digest already holds an equal-bytes message
	// (same hash). Do not reassign it: PrepareFinalizationResults reads message
	// under sc.mu only, not the storage lock, so it must stay fixed.
	_, exists := pc.signatureCollection[digest]
	if !exists {
		pc.signatureCollection[digest] = NewSignatureCollection(message, pc.signingPolicy, pc.threshold)
	}

	pc.messageChosenDigest = digest
	pc.messageAdded = true

	thresholdReached := false

	// reset bufferedSenders before re-running addPayload on the buffered set —
	// they were counted on the first pass and would otherwise hit the cap again.
	// The drain repopulates it, and it continues to bound post-message admits.
	pc.bufferedSenders = nil

	for _, up := range pc.unprocessedPayloads {
		tr, digestCheck, err := pc.addPayload(up)
		switch {
		case errors.Is(err, errBadPayload):
			logger.Debugf("Ignoring buffered signature for voting round %d, protocolID %d from sender %s: %v", up.votingRoundID, up.protocolID, up.sender, err)
		case err != nil:
			logger.Errorf("Failed to add buffered signature for voting round %d, protocolID %d from sender %s: %v", up.votingRoundID, up.protocolID, up.sender, err)
		case digestCheck != digest:
			logger.Debug("Unexpected behavior, hashes should match")
		}
		if tr {
			thresholdReached = true
		}
	}

	//clear unprocessedPayloads
	pc.unprocessedPayloads = nil

	return thresholdReached, digest, nil
}

func (pc *protocolCollection) addPayload(payload *submitSignaturesPayload) (bool, common.Hash, error) {
	// DOS-01: cap each sender to one payload per (round, protocol). Applied above
	// the type/message-state branching to bound both attack vectors:
	//   - pre-message TypeID-1 buffer (unprocessedPayloads + ECDSA burst on drain)
	//   - TypeID-0 signatureCollection allocations keyed on attacker-controlled hashes
	if pc.bufferedSenders == nil {
		pc.bufferedSenders = make(map[common.Address]struct{})
	}
	if _, seen := pc.bufferedSenders[payload.sender]; seen {
		return false, common.Hash{}, nil
	}
	pc.bufferedSenders[payload.sender] = struct{}{}

	if !pc.messageAdded && payload.typeID != 0 {
		pc.unprocessedPayloads = append(pc.unprocessedPayloads, payload)

		return false, common.Hash{}, nil
	}

	// key and digest can differ: messageDigest follows the round once the boundary is
	// learned, which can happen after the message was filed under the fallback form.
	// The signature must be recovered under the form its signer used, but the caller
	// gets the key the collection is actually stored under.
	var digest []byte
	var key common.Hash
	var sigCollection *signaturesCollection
	if payload.typeID == 0 {
		digest = pc.messageDigest(payload.message)
		key = common.Hash(digest)
		sc, exists := pc.signatureCollection[key]
		if !exists {
			sc = NewSignatureCollection(payload.message, pc.signingPolicy, pc.threshold)
			pc.signatureCollection[key] = sc
		}
		sigCollection = sc
	} else if pc.messageAdded {
		key = pc.messageChosenDigest
		sigCollection = pc.signatureCollection[key]
		digest = pc.messageDigest(sigCollection.message)
	} else {
		return false, common.Hash{}, errors.New("unexpected behavior, no message")
	}

	err := payload.AddSigner(digest, sigCollection.signingPolicy.Voters)
	if err != nil {
		return false, common.Hash{}, err
	}

	thresholdReached, err := sigCollection.addSignature(payload)

	return thresholdReached, key, err
}

func newFinalizationStorage(relayCutover *shared.RelayCutover) *finalizationStorage {
	return &finalizationStorage{
		stg:          make(map[uint32]*roundCollection),
		relayCutover: relayCutover,
	}
}

// addPayload adds a submitSignature payload to the finalizationStorage.
// The payload is added to the protocolCollection for the protocolID and roundID of the payload.
// An indicator whether the addition has made the protocol reach the threshold for the round is returned.
func (s *finalizationStorage) addPayload(p *submitSignaturesPayload, signingPolicy *policy.SigningPolicy, threshold uint16) (FinalizationReady, error) {
	s.Lock()
	defer s.Unlock()

	if p.votingRoundID < s.lowestRoundStored {
		return FinalizationReady{thresholdReached: false}, fmt.Errorf("%w: round %d before lowest stored round %d", errBadPayload, p.votingRoundID, s.lowestRoundStored)
	}

	rc, exists := s.stg[p.votingRoundID]
	if !exists {
		rc = &roundCollection{protocolCollections: make(map[uint8]*protocolCollection)}

		s.stg[p.votingRoundID] = rc
	}

	pc, exists := rc.protocolCollections[p.protocolID]
	if !exists {
		pc = &protocolCollection{signingPolicy: signingPolicy, signatureCollection: make(map[common.Hash]*signaturesCollection), threshold: threshold, relayCutover: s.relayCutover}
		rc.protocolCollections[p.protocolID] = pc
	}

	thresholdReached, digest, err := pc.addPayload(p)
	if err != nil {
		return FinalizationReady{thresholdReached: false}, err
	}
	if thresholdReached {
		return FinalizationReady{thresholdReached: true, protocolID: p.protocolID, votingRoundID: p.votingRoundID, digest: digest}, nil
	}

	return FinalizationReady{thresholdReached: false}, nil
}

// AddMessage adds a protocol message to the finalizationStorage for the respective protocol and round, and adds all unprocessedPayloads for the respective round and protocol.
// An indicator whether the additions have made the protocol reach the threshold for the round is returned.
func (s *finalizationStorage) AddMessage(p *shared.ProtocolMessage, signingPolicy *policy.SigningPolicy, threshold uint16) (FinalizationReady, error) {
	s.Lock()
	defer s.Unlock()

	if p.VotingRoundID < s.lowestRoundStored {
		return FinalizationReady{thresholdReached: false}, fmt.Errorf("message for round %d before lowest stored round %d", p.VotingRoundID, s.lowestRoundStored)
	}

	rc, exists := s.stg[p.VotingRoundID]
	if !exists {
		rc = &roundCollection{protocolCollections: make(map[uint8]*protocolCollection)}
		s.stg[p.VotingRoundID] = rc
	}

	pc, exists := rc.protocolCollections[p.ProtocolID]
	if !exists {
		pc = &protocolCollection{signatureCollection: make(map[common.Hash]*signaturesCollection), signingPolicy: signingPolicy, threshold: threshold, relayCutover: s.relayCutover}
		rc.protocolCollections[p.ProtocolID] = pc
	}

	thresholdReached, digest, err := pc.addMessage(p.Message)
	if err != nil {
		return FinalizationReady{thresholdReached: false}, err
	}
	if thresholdReached {
		return FinalizationReady{thresholdReached: true, protocolID: p.ProtocolID, votingRoundID: p.VotingRoundID, digest: digest}, nil
	}

	return FinalizationReady{thresholdReached: false}, nil
}

// get returns the signatureCollection for votingRoundID and protocolID.
// A boolean inductor of existence is also returned.
// Access or mutate signatures, weight, and thresholdReached under the mutex;
// the other fields are fixed after creation.
func (fs *finalizationStorage) get(votingRoundID uint32, protocolID uint8, digest common.Hash) (*signaturesCollection, bool) {
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

	sigCollection, exists := pc.signatureCollection[digest]
	if !exists {
		return &signaturesCollection{}, false
	}

	return sigCollection, true
}

// LowestRoundStored returns the lowest round that is still stored.
func (fs *finalizationStorage) LowestRoundStored() uint32 {
	fs.RLock()
	defer fs.RUnlock()

	return fs.lowestRoundStored
}

// RemoveRoundsBefore deletes rounds before votingRoundID.
func (fs *finalizationStorage) RemoveRoundsBefore(votingRoundID uint32) {
	fs.Lock()
	defer fs.Unlock()

	// initial cleanup
	if fs.lowestRoundStored == 0 && votingRoundID > 20 {
		fs.lowestRoundStored = votingRoundID - 20
	}

	if votingRoundID > fs.lowestRoundStored {
		for i := fs.lowestRoundStored; i < votingRoundID; i++ {
			logger.Debugf("Deleting round %d in finalization storage", i)
			delete(fs.stg, i)
		}

		// votingRoundID is the lowest round that remains stored
		fs.lowestRoundStored = votingRoundID
	}
}
