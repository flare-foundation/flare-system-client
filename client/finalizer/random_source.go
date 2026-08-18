package finalizer

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
)

// randomDataWait bounds how long a send waits for a round's data when peers reached the
// threshold before the local provider answered; the delayed queue retries later, by
// which time the message has arrived or the round is someone else's.
const randomDataWait = 10 * time.Second

var errNoRandomSource = errors.New("no random source configured")

// randomSource supplies the random number and Merkle proof that relay() requires to
// finalize the random protocol on the new Relay. The random protocol's provider serves
// them with its submitSignatures message, so they exist exactly when the round's
// message does; the submitter forwards them and the message listener stores them.
type randomSource interface {
	// Put stores a round's data; a second Put for the round is ignored.
	Put(votingRoundID uint32, data randomData)

	// Get returns the round's data, waiting a bounded time for a message still on its way.
	Get(ctx context.Context, votingRoundID uint32) (randomData, error)

	// RemoveBefore drops rounds that can no longer be finalized.
	RemoveBefore(votingRoundID uint32)
}

// randomEntry is one round's data, or the promise of it: done is closed by Put.
type randomEntry struct {
	done chan struct{}
	data randomData
}

type randomStore struct {
	wait time.Duration

	mu      sync.Mutex
	entries map[uint32]*randomEntry
}

func newRandomStore() *randomStore {
	return &randomStore{
		wait:    randomDataWait,
		entries: make(map[uint32]*randomEntry),
	}
}

// randomProtocolConfigured reports whether the submitter queries the protocol whose
// provider serves the random data — otherwise the finalizer could never see it.
func randomProtocolConfigured(protocols map[string]config.ProtocolConfig, protocolID uint8) bool {
	for _, protocol := range protocols {
		if protocol.ID == protocolID {
			return true
		}
	}
	return false
}

// entry returns the round's entry, pending if nothing has been stored yet.
func (s *randomStore) entry(votingRoundID uint32) *randomEntry {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry, ok := s.entries[votingRoundID]
	if !ok {
		entry = &randomEntry{done: make(chan struct{})}
		s.entries[votingRoundID] = entry
	}
	return entry
}

func (s *randomStore) Put(votingRoundID uint32, data randomData) {
	entry := s.entry(votingRoundID)

	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-entry.done:
		return
	default:
	}
	entry.data = data
	close(entry.done)
}

func (s *randomStore) Get(ctx context.Context, votingRoundID uint32) (randomData, error) {
	entry := s.entry(votingRoundID)

	timer := time.NewTimer(s.wait)
	defer timer.Stop()

	select {
	case <-entry.done:
		return entry.data, nil
	case <-timer.C:
		return randomData{}, fmt.Errorf("no finalization data for round %d arrived from the provider within %s", votingRoundID, s.wait)
	case <-ctx.Done():
		return randomData{}, fmt.Errorf("waiting for the finalization data of round %d: %w", votingRoundID, ctx.Err())
	}
}

func (s *randomStore) RemoveBefore(votingRoundID uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for round := range s.entries {
		if round < votingRoundID {
			delete(s.entries, round)
		}
	}
}

// unconfiguredRandomSource stands in while no Relay cutover is scheduled, when the
// trailer is never needed.
type unconfiguredRandomSource struct{}

func (unconfiguredRandomSource) Put(uint32, randomData) {}

func (unconfiguredRandomSource) Get(context.Context, uint32) (randomData, error) {
	return randomData{}, errNoRandomSource
}

func (unconfiguredRandomSource) RemoveBefore(uint32) {}
