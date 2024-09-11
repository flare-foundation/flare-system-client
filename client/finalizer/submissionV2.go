package finalizer

import (
	"context"
	"encoding/hex"
	"flare-fsc/client/shared"
	"flare-fsc/logger"
	"flare-fsc/utils/contracts/submission"
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/slices"
)

type signaturesCollection struct {
	message          shared.Message
	signatures       [][]byte
	weight           uint16
	thresholdReached bool
	signingPolicy    *signingPolicy
}

type protocolCollection struct {
	messageAdded        bool
	signatureCollection signaturesCollection
	unprocessedPayloads []*submitSignaturesPayload
	signingPolicy       *signingPolicy
}

// roundCollection maps protocolID to signatureCollection
type roundCollection struct {
	protocolCollections map[uint8]*protocolCollection
}

type finalizationStorage struct {
	stg               map[uint32]*roundCollection // map from roundID to roundCollection
	lowestRoundStored uint32

	// mutex
	sync.Mutex
}

type FinalizationReady struct {
	thresholdReached bool
	protocolID       uint8
	votingRoundID    uint32
}

func NewSignatureCollection(message []byte, signingPolicy *signingPolicy) *signaturesCollection {
	return &signaturesCollection{
		message:       message,
		signatures:    make([][]byte, signingPolicy.voters.Count()),
		signingPolicy: signingPolicy,
	}
}

func (sc *signaturesCollection) addSignature(p *submitSignaturesPayload) (bool, error) {
	if p.voterIndex < 0 {
		return false, fmt.Errorf("voter not recognized")
	}

	if len(sc.signatures[p.voterIndex]) != 0 {
		return false, fmt.Errorf("signature for %s already added", p.signer)
	}

	sc.signatures[p.voterIndex] = p.signature

	sc.weight += p.weight
	if !sc.thresholdReached {
		sc.thresholdReached = sc.weight > sc.signingPolicy.threshold

		return sc.thresholdReached, nil
	}
	return false, nil
}

func (pc *protocolCollection) addMessage(message []byte) (bool, error) {
	if pc.messageAdded {
		return false, fmt.Errorf("message added twice")
	}

	pc.signatureCollection = *NewSignatureCollection(message, pc.signingPolicy)
	pc.messageAdded = true

	for _, up := range pc.unprocessedPayloads {
		thresholdReached, err := pc.addPayload(up)

		if err != nil {
			logger.Debug("%v", err)
		}

		if thresholdReached {
			return true, nil
		}
	}

	//clear unprocessedPayloads
	pc.unprocessedPayloads = nil

	return false, nil
}

func (pc *protocolCollection) addPayload(payload *submitSignaturesPayload) (bool, error) {
	if !pc.messageAdded {
		pc.unprocessedPayloads = append(pc.unprocessedPayloads, payload)

		return false, nil
	}

	err := payload.AddSigner(pc.signatureCollection.message.Hash(), pc.signatureCollection.signingPolicy.voters)
	if err != nil {
		return false, fmt.Errorf("adding payload, %v", err)
	}

	return pc.signatureCollection.addSignature(payload)
}

func newFinalizationStorage() *finalizationStorage {
	return &finalizationStorage{
		stg: make(map[uint32]*roundCollection),
	}
}

// addPayload add a submitSignature payload to the finalizationStorage.
func (s *finalizationStorage) addPayload(p *submitSignaturesPayload, signingPolicy *signingPolicy) (FinalizationReady, error) {
	s.Lock()
	defer s.Unlock()

	if p.votingRoundID < s.lowestRoundStored {
		return FinalizationReady{thresholdReached: false}, nil //TODO
	}

	rc, exists := s.stg[p.votingRoundID]
	if !exists {
		rc = &roundCollection{}

		s.stg[p.votingRoundID] = rc
	}

	pc, exists := rc.protocolCollections[p.protocolID]
	if !exists {
		pc = &protocolCollection{signingPolicy: signingPolicy}
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

// Add adds a signed payload to the submission storage
// The provided signing policy must be the signing policy for the voting round
// Returns true if the payload was added, false if it was already added
func (s *finalizationStorage) AddMessage(p *shared.ProtocolMessage, signingPolicy *signingPolicy) (FinalizationReady, bool, error) {
	s.Lock()
	defer s.Unlock()

	if p.VotingRoundID < s.lowestRoundStored {
		return FinalizationReady{}, false, nil //TODO
	}

	rc, exists := s.stg[p.VotingRoundID]
	if !exists {
		rc = &roundCollection{protocolCollections: make(map[uint8]*protocolCollection)}
		s.stg[p.VotingRoundID] = rc
	}

	pc, exists := rc.protocolCollections[p.ProtocolID]
	if !exists {
		pc = &protocolCollection{signingPolicy: signingPolicy}
		rc.protocolCollections[p.ProtocolID] = pc
	}

	thresholdReached, err := pc.addMessage(p.Message)
	if err != nil {
		return FinalizationReady{}, false, err
	}
	if thresholdReached {
		return FinalizationReady{protocolID: p.ProtocolID, votingRoundID: p.VotingRoundID}, true, nil
	}

	return FinalizationReady{}, false, nil
}

func (sc *signaturesCollection) PrepareFinalizationResults() (FinalizationResult, error) {
	availableSignatures := []IndexedSignature{}
	selectedSignatures := []IndexedSignature{}

	for i := range sc.signatures {
		if len(sc.signatures[i]) > 0 {
			availableSignatures = append(availableSignatures, IndexedSignature{index: i, signature: sc.signatures[i]})
		}
	}

	// sort decreasing by weight
	slices.SortFunc(availableSignatures, func(a, b IndexedSignature) int {
		return int(sc.signingPolicy.voters.VoterWeight(b.index)) - int(sc.signingPolicy.voters.VoterWeight(a.index))
	})

	// greedy select until threshold is reached
	weight := uint16(0)
	for i := range availableSignatures {
		selectedSignatures = append(selectedSignatures, availableSignatures[i])
		weight += sc.signingPolicy.voters.VoterWeight(availableSignatures[i].index)
		if weight > sc.signingPolicy.threshold {
			break
		}
	}
	if weight <= sc.signingPolicy.threshold {
		return FinalizationResult{}, fmt.Errorf("threshold not reached")
	}

	// sort selected payloads by index
	slices.SortFunc(selectedSignatures, func(p, q IndexedSignature) int {
		return p.index - q.index
	})

	return FinalizationResult{message: sc.message, signatures: selectedSignatures, signingPolicy: sc.signingPolicy}, nil
}

func (fs *finalizationStorage) Get(votingRoundID uint32, protocolID uint8) (*signaturesCollection, bool) {
	fs.Lock()
	defer fs.Unlock()
	round, exists := fs.stg[votingRoundID]

	if !exists {
		return &signaturesCollection{}, false
	}

	pc, exists := round.protocolCollections[protocolID]

	if !exists || !pc.messageAdded {
		return &signaturesCollection{}, false
	}

	return &pc.signatureCollection, true
}

type submissionListenerResponseV2 struct {
	payloads  []*submitSignaturesPayload
	timestamp int64
}

func (fs *finalizationStorage) RemoveRoundsBefore(votingRoundID uint32) {
	if votingRoundID > fs.lowestRoundStored {
		fs.Lock()
		defer fs.Unlock()

		for i := fs.lowestRoundStored; i < votingRoundID; i++ {

			logger.Info("deleting round", i)
			delete(fs.stg, i)
		}

		fs.lowestRoundStored = votingRoundID
	}
}

type submitterProcessorV2 interface {
	// Return error if the submission was not processed and needs a retry
	// Should be able to handle duplicates
	ProcessSubmissionData(submissionListenerResponseV2) error
}

func (s *submissionContractClient) SubmissionTxListenerV2(
	ctx context.Context,
	db finalizerDB,
	startTime time.Time,
	processor submitterProcessorV2,
) error {
	submissionABI, err := submission.SubmissionMetaData.GetAbi()
	if err != nil {
		// Should not happen, unhandled errors will cause a panic further up.
		return err
	}

	selector := submissionABI.Methods["submitSignatures"].ID
	ticker := time.NewTicker(shared.ListenerInterval)
	eventRangeStart := startTime.Unix()
	for {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			logger.Info("Submission tx listener stopped")
			return ctx.Err()
		}
		now := time.Now().Unix()
		txs, err := db.FetchTransactionsByAddressAndSelector(s.address, selector, eventRangeStart, now)
		if err != nil {
			logger.Error("Error fetching transactions %v", err)
			continue
		}
		for _, tx := range txs {
			inputBytes, err := hex.DecodeString(tx.Input)
			if err != nil {
				logger.Info("Invalid submitSignatures tx sent by %s: %v, skipping", tx.FromAddress, err)
			}
			payloads, err := ExtractPayloads(inputBytes)
			if err != nil {
				// if input cannot be decoded, it is not a valid submission and should be skipped
				logger.Info("Invalid submitSignatures input sent by %s: %v, skipping", tx.FromAddress, err)
			}

			signaturePayloads := []*submitSignaturesPayload{}
			for i := range payloads {
				signaturePayload, err := decodeSignedPayloadV2(payloads[i])
				if err != nil {
					// if input cannot be decoded, it is not a valid submission and should be skipped
					logger.Info("Invalid submitSignatures payload sent by %s: %v, skipping", tx.FromAddress, err)

				}
				signaturePayloads = append(signaturePayloads, &signaturePayload)
			}

			if len(signaturePayloads) > 0 {
				err = processor.ProcessSubmissionData(submissionListenerResponseV2{
					payloads:  signaturePayloads,
					timestamp: int64(tx.Timestamp),
				})
				if err != nil {
					// retry the full range, error occurs when the corresponding signing policy
					// is not yet available
					logger.Warn("Error processing submitSignatures payload sent by %s: %v, retrying", tx.FromAddress, err)
					break
				}
			}
			// -1 for overlap in case of an error and retry above
			// processor should be able to handle duplicates
			eventRangeStart = int64(tx.Timestamp) - 1
		}
	}
}
