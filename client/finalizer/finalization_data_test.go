package finalizer

import (
	"bytes"
	"context"
	"encoding/binary"
	"testing"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/policy"
	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

const randomProtocolID = uint8(100)

var (
	randomValue = common.HexToHash("0x00000000000000000000000000000000000000000000000000000000deadbeef")
	proofNodeA  = common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")
)

// words is the shape the provider serves and relay() consumes: value(32) ‖ proof(32×d).
func words(hashes ...common.Hash) []byte {
	buf := make([]byte, 0, common.HashLength*len(hashes))
	for _, h := range hashes {
		buf = append(buf, h[:]...)
	}
	return buf
}

// buildMessage assembles the signed protocol message:
// protocolID(1) ‖ votingRoundID(4) ‖ isSecureRandom(1) ‖ merkleRoot(32).
func buildMessage(protocolID uint8, round uint32, root common.Hash) shared.Message {
	msg := make(shared.Message, shared.RelayMessageLength)
	msg[0] = protocolID
	binary.BigEndian.PutUint32(msg[1:5], round)
	msg[5] = 1
	copy(msg[6:], root[:])
	return msg
}

// The random protocol must be one the submitter queries, or startup fails; the same set gates
// which peer payloads are stored, so an unset one admits nothing.
func TestServesOnlyConfiguredProtocols(t *testing.T) {
	fc := &finalizerContext{protocolIDs: configuredProtocolIDs(map[string]config.ProtocolConfig{
		"ftso": {ID: randomProtocolID, APIURL: "https://ftso.example/api"},
		"fdc":  {ID: 200, APIURL: "https://fdc.example"},
	})}
	require.True(t, fc.serves(randomProtocolID))
	require.True(t, fc.serves(200))
	require.False(t, fc.serves(42))

	require.False(t, (&finalizerContext{protocolIDs: configuredProtocolIDs(nil)}).serves(randomProtocolID))
	require.False(t, (&finalizerContext{}).serves(randomProtocolID), "an unset set must fail closed")
}

// Only the random protocol's rounds on the new Relay keep what the message carries,
// and only in the shape relay() reads it: whole 32-byte words.
func TestFinalizationDataToStore(t *testing.T) {
	full := words(randomValue, proofNodeA)
	atCap := make([]byte, maxFinalizationDataLength)
	overCap := make([]byte, maxFinalizationDataLength+common.HashLength)
	policyAt := func(epoch int64) *policy.SigningPolicy { return &policy.SigningPolicy{RewardEpochID: epoch} }

	cases := []struct {
		name       string
		protocolID uint8
		epoch      int64
		data       []byte
		want       []byte
	}{
		{"kept for the random protocol from the breaking epoch", randomProtocolID, testBreakingEpoch, full, full},
		{"a single word is a single-leaf tree", randomProtocolID, testBreakingEpoch, randomValue.Bytes(), randomValue.Bytes()},
		{"missing data is reported, not stored", randomProtocolID, testBreakingEpoch, nil, nil},
		{"a truncated word is not stored", randomProtocolID, testBreakingEpoch, full[:len(full)-1], nil},
		{"exactly the cap is kept", randomProtocolID, testBreakingEpoch, atCap, atCap},
		{"one word over the cap is not stored", randomProtocolID, testBreakingEpoch, overCap, nil},
		{"ignored before the breaking epoch", randomProtocolID, testBreakingEpoch - 1, full, nil},
		{"ignored for another protocol", 200, testBreakingEpoch, full, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cl := &client{
				finalizerContext: &finalizerContext{randomNumberProtocolID: randomProtocolID},
				relayCutover:     cutoverForTest(),
			}
			m := &shared.ProtocolMessage{
				ProtocolID:       c.protocolID,
				VotingRoundID:    7,
				Message:          buildMessage(c.protocolID, 7, common.Hash{}),
				FinalizationData: c.data,
			}
			require.Equal(t, c.want, cl.finalizationDataToStore(m, policyAt(c.epoch)))
		})
	}

	// with no cutover scheduled nothing is ever stored
	plain := &client{
		finalizerContext: &finalizerContext{randomNumberProtocolID: randomProtocolID},
		relayCutover:     shared.NewRelayCutover(testChainID, common.Address{}, 0),
	}
	require.Nil(t, plain.finalizationDataToStore(
		&shared.ProtocolMessage{ProtocolID: randomProtocolID, VotingRoundID: 7, FinalizationData: full},
		policyAt(testBreakingEpoch)))
}

// Only the new Relay verifies and stores the data; a pre-cutover round goes to the old one.
func TestNeedsFinalizationData(t *testing.T) {
	p := &finalizerQueueProcessor{
		finalizerContext: &finalizerContext{randomNumberProtocolID: randomProtocolID},
		relayClient:      &relayContractClient{address: oldRelayAddress, relayCutover: cutoverForTest()},
	}

	require.False(t, p.needsFinalizationData(randomProtocolID, testBreakingEpoch-1))
	require.True(t, p.needsFinalizationData(randomProtocolID, testBreakingEpoch))
	require.True(t, p.needsFinalizationData(randomProtocolID, testBreakingEpoch+1))
	require.False(t, p.needsFinalizationData(200, testBreakingEpoch))

	// with no cutover scheduled at all, no round ever needs one
	plain := &finalizerQueueProcessor{
		finalizerContext: &finalizerContext{randomNumberProtocolID: randomProtocolID},
		relayClient: &relayContractClient{
			address:      oldRelayAddress,
			relayCutover: shared.NewRelayCutover(testChainID, common.Address{}, 0),
		},
	}
	require.False(t, plain.needsFinalizationData(randomProtocolID, testBreakingEpoch+1))
}

// The data lives with the message it arrived with.
func TestAddMessageStoresTheDataWithItsOwnMessage(t *testing.T) {
	sp := &policy.SigningPolicy{
		RewardEpochID: testBreakingEpoch,
		Voters:        voters.NewSet([]common.Address{{}}, []uint16{1}, nil),
	}
	message := buildMessage(randomProtocolID, 9, randomValue)
	data := words(randomValue, proofNodeA)

	storage := newFinalizationStorage(cutoverForTest())
	_, err := storage.AddMessage(&shared.ProtocolMessage{
		ProtocolID: randomProtocolID, VotingRoundID: 9, Message: message, FinalizationData: data,
	}, sp, 1)
	require.NoError(t, err)

	sc, exists := storage.get(9, randomProtocolID)
	require.True(t, exists)
	require.Equal(t, message, sc.message)
	require.Equal(t, data, sc.finalizationData)
}

// relayRound finalizes one already-crossed round of the random protocol and returns what it sent:
// rewardEpochID picks the target Relay, data is what the provider served with the message.
func relayRound(t *testing.T, rewardEpochID int64, data []byte) *testEthClient {
	t.Helper()
	cutover := cutoverForTest()
	key, err := crypto.HexToECDSA(testPrivateKeyHex)
	require.NoError(t, err)
	signer := crypto.PubkeyToAddress(key.PublicKey)

	message := buildMessage(randomProtocolID, 7, randomValue)
	sp := &policy.SigningPolicy{
		RewardEpochID: rewardEpochID,
		Voters:        voters.NewSet([]common.Address{signer}, []uint16{2}, nil),
	}

	storage := newFinalizationStorage(cutover)
	_, err = storage.AddMessage(&shared.ProtocolMessage{
		ProtocolID: randomProtocolID, VotingRoundID: 7, Message: message, FinalizationData: data,
	}, sp, 1)
	require.NoError(t, err)

	// the boundary is unlearned here, so the digest follows the policy's epoch
	ready, err := storage.addPayload(&submitSignaturesPayload{
		sender: signer, protocolID: randomProtocolID, votingRoundID: 7,
		signature: signDigest(t, shared.MessageDigest(message, cutover.ChainID, rewardEpochID >= testBreakingEpoch), key),
	}, sp, 1)
	require.NoError(t, err)
	require.True(t, ready.thresholdReached)

	eth := new(testEthClient)
	relayClient, err := NewRelayContractClient(nil, oldRelayAddress, key, signer, &config.Gas{}, testChainID, cutover)
	require.NoError(t, err)
	relayClient.chainClient = eth
	relayClient.retryDelay = time.Millisecond

	qp := newFinalizerQueueProcessor(&testDB{}, storage, relayClient, &finalizerContext{
		randomNumberProtocolID: randomProtocolID,
		votingRoundTiming:      &utils.EpochTimingConfig{Start: time.Unix(0, 0), Period: time.Hour},
	})
	qp.processItem(context.Background(), &queueItem{protocolID: randomProtocolID, votingRoundID: 7}, true)

	return eth
}

// A pre-cutover round goes to the old Relay, which never reads the appended words. The collection can
// still hold them: the arrival filter judges the message-time policy, the send the collection's.
func TestFinalizationDataIsNotAppendedForTheOldRelay(t *testing.T) {
	data := words(randomValue, proofNodeA)
	eth := relayRound(t, testBreakingEpoch-1, data)

	require.Len(t, eth.sentTxs, 1)
	require.Equal(t, oldRelayAddress, eth.sentTxs[0].to)
	require.False(t, bytes.HasSuffix(eth.sentTxs[0].data, data), "the old Relay never reads them")
}

// From the breaking epoch on the words ride the tx, and only the new Relay holds that policy's hash.
func TestFinalizationDataIsAppendedForTheNewRelay(t *testing.T) {
	data := words(randomValue, proofNodeA)
	eth := relayRound(t, testBreakingEpoch, data)

	require.Len(t, eth.sentTxs, 1)
	require.Equal(t, newRelayAddress, eth.sentTxs[0].to)
	require.True(t, bytes.HasSuffix(eth.sentTxs[0].data, data))
}

// relay() reverts without the words, so a round the provider served none for is not sent at all.
func TestNothingIsSentWithoutTheFinalizationData(t *testing.T) {
	require.Empty(t, relayRound(t, testBreakingEpoch, nil).sentTxs)
}

// The data only reaches the tx input when the finalization carries one.
func TestPrepareFinalizationTxInputAppendsTheData(t *testing.T) {
	sp := &policy.SigningPolicy{Voters: voters.NewSet([]common.Address{{}}, []uint16{1}, nil)}
	message := buildMessage(randomProtocolID, 7, common.Hash{})
	data := words(randomValue, proofNodeA)

	withData := FinalizationResult{
		message:          message,
		signingPolicy:    sp,
		signatures:       []IndexedSignature{{index: 0, signature: make([]byte, 65)}},
		finalizationData: data,
	}
	without := FinalizationResult{
		message:       message,
		signingPolicy: sp,
		signatures:    []IndexedSignature{{index: 0, signature: make([]byte, 65)}},
	}

	withInput, err := withData.PrepareFinalizationTxInput()
	require.NoError(t, err)
	withoutInput, err := without.PrepareFinalizationTxInput()
	require.NoError(t, err)

	require.Equal(t, append(append([]byte{}, withoutInput...), data...), withInput)

	// data that is not a whole number of words would revert on chain
	broken := withData
	broken.finalizationData = data[:len(data)-1]
	_, err = broken.PrepareFinalizationTxInput()
	require.ErrorContains(t, err, "not a multiple of")

	atCap := withData
	atCap.finalizationData = make([]byte, maxFinalizationDataLength)
	_, err = atCap.PrepareFinalizationTxInput()
	require.NoError(t, err)

	oversized := withData
	oversized.finalizationData = make([]byte, maxFinalizationDataLength+common.HashLength)
	_, err = oversized.PrepareFinalizationTxInput()
	require.ErrorContains(t, err, "over the")
}
