package finalizer

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"testing"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/encoding"
	"github.com/flare-foundation/go-flare-common/pkg/policy"
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

// Only the random protocol's rounds keep what the message carries,
// and only in the shape relay() reads it: whole 32-byte words.
func TestFinalizationDataToStore(t *testing.T) {
	full := words(randomValue, proofNodeA)
	atCap := make([]byte, maxFinalizationDataLength)
	overCap := make([]byte, maxFinalizationDataLength+common.HashLength)

	cases := []struct {
		name       string
		protocolID uint8
		data       []byte
		want       []byte
	}{
		{"kept for the random protocol", randomProtocolID, full, full},
		{"a single word is a single-leaf tree", randomProtocolID, randomValue.Bytes(), randomValue.Bytes()},
		{"missing data is reported, not stored", randomProtocolID, nil, nil},
		{"a truncated word is not stored", randomProtocolID, full[:len(full)-1], nil},
		{"exactly the cap is kept", randomProtocolID, atCap, atCap},
		{"one word over the cap is not stored", randomProtocolID, overCap, nil},
		{"ignored for another protocol", 200, full, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cl := &client{
				finalizerContext: &finalizerContext{randomNumberProtocolID: randomProtocolID},
			}
			m := &shared.ProtocolMessage{
				ProtocolID:       c.protocolID,
				VotingRoundID:    7,
				Message:          buildMessage(c.protocolID, 7, common.Hash{}),
				FinalizationData: c.data,
			}
			require.Equal(t, c.want, cl.finalizationDataToStore(m))
		})
	}
}

// The data lives with the message it arrived with.
func TestAddMessageStoresTheDataWithItsOwnMessage(t *testing.T) {
	sp := &policy.SigningPolicy{
		Voters: newVoterSet(t, []common.Address{{}}, []uint16{1}),
	}
	message := buildMessage(randomProtocolID, 9, randomValue)
	data := words(randomValue, proofNodeA)

	storage := newFinalizationStorage(testChainID)
	_, err := storage.AddMessage(&shared.ProtocolMessage{
		ProtocolID: randomProtocolID, VotingRoundID: 9, Message: message, FinalizationData: data,
	}, sp, 1)
	require.NoError(t, err)

	sc, exists := storage.get(9, randomProtocolID)
	require.True(t, exists)
	require.Equal(t, message, sc.message)
	require.Equal(t, data, sc.finalizationData)
}

// relayRound finalizes one round of the random protocol and returns what it sent; data is what
// the provider served with the message.
func relayRound(t *testing.T, data []byte) *testEthClient {
	t.Helper()
	key, err := crypto.HexToECDSA(testPrivateKeyHex)
	require.NoError(t, err)
	signer := crypto.PubkeyToAddress(key.PublicKey)

	message := buildMessage(randomProtocolID, 7, randomValue)
	sp := &policy.SigningPolicy{
		Voters: newVoterSet(t, []common.Address{signer}, []uint16{2}),
	}

	storage := newFinalizationStorage(testChainID)
	_, err = storage.AddMessage(&shared.ProtocolMessage{
		ProtocolID: randomProtocolID, VotingRoundID: 7, Message: message, FinalizationData: data,
	}, sp, 1)
	require.NoError(t, err)

	ready, err := storage.addPayload(&submitSignaturesPayload{
		sender: signer, protocolID: randomProtocolID, votingRoundID: 7,
		signature: signDigest(t, shared.MessageDigest(message, testChainID), key),
	}, sp, 1)
	require.NoError(t, err)
	require.True(t, ready.thresholdReached)

	eth := new(testEthClient)
	relayClient, err := NewRelayContractClient(nil, relayContractAddress, key, signer, &config.Gas{}, testChainID)
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

func signDigest(t *testing.T, digest []byte, key *ecdsa.PrivateKey) []byte {
	t.Helper()
	signature, err := crypto.Sign(digest, key)
	require.NoError(t, err)
	vrs, err := encoding.TransformSignatureRSVtoVRS(signature)
	require.NoError(t, err)
	return vrs
}

func TestFinalizationDataIsAppended(t *testing.T) {
	data := words(randomValue, proofNodeA)
	eth := relayRound(t, data)

	require.Len(t, eth.sentTxs, 1)
	require.Equal(t, relayContractAddress, eth.sentTxs[0].to)
	require.True(t, bytes.HasSuffix(eth.sentTxs[0].data, data))
}

// relay() reverts without the words, so a round the provider served none for is not sent at all.
func TestNothingIsSentWithoutTheFinalizationData(t *testing.T) {
	require.Empty(t, relayRound(t, nil).sentTxs)
}

// The data only reaches the tx input when the finalization carries one.
func TestPrepareFinalizationTxInputAppendsTheData(t *testing.T) {
	sp := &policy.SigningPolicy{Voters: newVoterSet(t, []common.Address{{}}, []uint16{1})}
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
