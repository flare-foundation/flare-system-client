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

	// what relay() must have appended for this message, served with it by the provider
	finalizationData []byte

	mu sync.RWMutex
}

type protocolCollection struct {
	collection          *signaturesCollection // nil until the local message arrives
	unprocessedPayloads []*submitSignaturesPayload
	bufferedSenders     map[common.Address]struct{} // senders already buffered pre-message (DOS-01 cap)
	signingPolicy       *policy.SigningPolicy
	threshold           uint16
	relayCutover        *shared.RelayCutover
}

// messageDigest returns the digest the signatures cover, derived from the round in the
// message bytes — the contract's own derivation, and the one every signer used. Falls back to
// the governing policy's epoch, as the target Relay does, when the boundary is unlearned (a
// restart may fetch only post-breaking policies) or the bytes do not parse.
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

// addMessage creates the round's only collection and drains the buffered payloads into it.
func (pc *protocolCollection) addMessage(message shared.Message, finalizationData []byte) (bool, error) {
	if pc.collection != nil {
		return false, errors.New("message added twice")
	}

	sc := NewSignatureCollection(message, pc.signingPolicy, pc.threshold)
	sc.finalizationData = finalizationData // not yet visible to a send, no lock needed
	pc.collection = sc

	thresholdReached := false

	// reset bufferedSenders before re-running addPayload on the buffered set —
	// they were counted on the first pass and would otherwise hit the cap again.
	// The drain repopulates it, and it continues to bound post-message admits.
	pc.bufferedSenders = nil

	for _, up := range pc.unprocessedPayloads {
		tr, err := pc.addPayload(up)
		switch {
		case errors.Is(err, errBadPayload):
			logger.Debugf("Ignoring buffered signature for voting round %d, protocolID %d from sender %s: %v", up.votingRoundID, up.protocolID, up.sender, err)
		case err != nil:
			logger.Errorf("Failed to add buffered signature for voting round %d, protocolID %d from sender %s: %v", up.votingRoundID, up.protocolID, up.sender, err)
		}
		if tr {
			thresholdReached = true
		}
	}

	//clear unprocessedPayloads
	pc.unprocessedPayloads = nil

	return thresholdReached, nil
}

// addPayload buffers payload until the message arrives, then verifies it against that message.
func (pc *protocolCollection) addPayload(payload *submitSignaturesPayload) (bool, error) {
	// DOS-01: one payload per sender per (round, protocol) — bounds the buffer and its ECDSA drain
	if pc.bufferedSenders == nil {
		pc.bufferedSenders = make(map[common.Address]struct{})
	}
	if _, seen := pc.bufferedSenders[payload.sender]; seen {
		return false, nil
	}
	pc.bufferedSenders[payload.sender] = struct{}{}

	if pc.collection == nil {
		pc.unprocessedPayloads = append(pc.unprocessedPayloads, payload)

		return false, nil
	}

	// per payload, not per message: the digest form can change once the cutover boundary is learned
	digest := pc.messageDigest(pc.collection.message)

	err := payload.AddSigner(digest, pc.collection.signingPolicy.Voters)
	if err != nil {
		return false, err
	}

	return pc.collection.addSignature(payload)
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
		pc = &protocolCollection{signingPolicy: signingPolicy, threshold: threshold, relayCutover: s.relayCutover}
		rc.protocolCollections[p.protocolID] = pc
	}

	thresholdReached, err := pc.addPayload(p)
	if err != nil {
		return FinalizationReady{thresholdReached: false}, err
	}
	if thresholdReached {
		return FinalizationReady{thresholdReached: true, protocolID: p.protocolID, votingRoundID: p.votingRoundID}, nil
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
		pc = &protocolCollection{signingPolicy: signingPolicy, threshold: threshold, relayCutover: s.relayCutover}
		rc.protocolCollections[p.ProtocolID] = pc
	}

	thresholdReached, err := pc.addMessage(p.Message, p.FinalizationData)
	if err != nil {
		return FinalizationReady{thresholdReached: false}, err
	}
	if thresholdReached {
		return FinalizationReady{thresholdReached: true, protocolID: p.ProtocolID, votingRoundID: p.VotingRoundID}, nil
	}

	return FinalizationReady{thresholdReached: false}, nil
}

// get returns the collection for votingRoundID and protocolID; none exists before the local message.
// signatures, weight and thresholdReached need the mutex; the rest is fixed before the collection is published.
func (fs *finalizationStorage) get(votingRoundID uint32, protocolID uint8) (*signaturesCollection, bool) {
	fs.RLock()
	defer fs.RUnlock()
	round, exists := fs.stg[votingRoundID]
	if !exists {
		return &signaturesCollection{}, false
	}

	pc, exists := round.protocolCollections[protocolID]
	if !exists || pc.collection == nil {
		return &signaturesCollection{}, false
	}

	return pc.collection, true
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
			fs.warnUnfinalized(i)
			delete(fs.stg, i)
		}

		// votingRoundID is the lowest round that remains stored
		fs.lowestRoundStored = votingRoundID
	}
}

// warnUnfinalized logs a round dropped short of the threshold — the only non-Debug sign that this
// node stopped finalizing. Protocols the local provider never served are skipped.
func (fs *finalizationStorage) warnUnfinalized(votingRoundID uint32) {
	rc, exists := fs.stg[votingRoundID]
	if !exists {
		return
	}

	for protocolID, pc := range rc.protocolCollections {
		sc := pc.collection
		if sc == nil {
			continue
		}

		sc.mu.RLock()
		reached, weight := sc.thresholdReached, sc.weight
		sc.mu.RUnlock()

		if !reached {
			logger.Warnf("Discarding round %d for protocol %d unfinalized: signed weight %d, need more than %d",
				votingRoundID, protocolID, weight, sc.threshold)
		}
	}
}
