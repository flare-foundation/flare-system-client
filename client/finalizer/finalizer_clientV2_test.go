package finalizer

import (
	"context"
	"flare-fsc/client/shared"
	"flare-fsc/config"
	"flare-fsc/logger"
	"flare-fsc/utils"
	"os"
	"testing"
	"time"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

const (
	testPrivateKeyHex            = "4f65bffe3c8ed6c0b812e84d35402e949feea042061cc1635fe6ae83ed84df4a"
	relayContractAddressHex      = "0xb849b93B585eFfb7cE4B522Ff88d9b3B24955f24"
	submissionContractAddressHex = "0x2F79Dce2375571207a7976148D4468195F89a73e"
	topicSPIHex                  = "0x91d0280e969157fc6c5b8f952f237b03d934b18534dafcac839075bbc33522f8"
)

var (
	relayContractAddress      = common.HexToAddress(relayContractAddressHex)
	submissionContractAddress = common.HexToAddress(submissionContractAddressHex)
)

func TestMain(m *testing.M) {
	logger.Configure(config.LoggerConfig{
		Level:   "DEBUG",
		Console: true,
	})

	os.Exit(m.Run())
}

func TestFinalizerClientMainlineV2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	clients, err := setupTestV2(t)
	require.NoError(t, err)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	require.Eventually(
		t, clients.eth.hasAnyCalls, 20*time.Second, 100*time.Millisecond,
	)

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Len(t, clients.eth.sentTxs, 1)
	cupaloy.SnapshotT(t, clients.eth.sentTxs[0])
}

func TestFinalizerClientSendTxErrV2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	clients, err := setupTestV2(t)
	require.NoError(t, err)

	clients.eth.sendTxErr = errors.New("sendRawTx error")

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	require.Eventually(
		t, clients.eth.hasAnyCalls, 10*time.Second, 100*time.Millisecond,
	)

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Empty(t, clients.eth.sentTxs)
}

func TestFinalizerClientFetchTxsErrV2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	clients, err := setupTestV2(t)
	require.NoError(t, err)

	clients.db.fetchTxsErr = errors.New("fetchTxs error")

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	require.Eventually(
		t, clients.db.hasAnyFetchTxsCalls, 10*time.Second, 100*time.Millisecond,
	)

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Empty(t, clients.eth.sentTxs)
}

func TestFinalizerClientFetchLogsErrV2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clients, err := setupTestV2(t)
	require.NoError(t, err)

	clients.db.fetchLogsErr = errors.New("fetchLogs error")

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	err = eg.Wait()
	require.True(t, errors.Is(err, clients.db.fetchLogsErr), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Empty(t, clients.eth.sentTxs)
}

type testClientsV2 struct {
	db        *testDB
	eth       *testEthClient
	finalizer *finalizerClientV2
}

func setupTestV2(t *testing.T) (*testClientsV2, error) {
	ethClient := new(testEthClient)

	privateKey, err := crypto.HexToECDSA(testPrivateKeyHex)
	if err != nil {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	relayClient, err := NewRelayContractClient(
		nil,
		relayContractAddress,
		privateKey,
		fromAddress,
	)
	if err != nil {
		return nil, err
	}

	relayClient.ethClient = ethClient

	finalizationStorage := newFinalizationStorage()

	db, err := newTestDB(privateKey)
	if err != nil {
		return nil, err
	}

	fCtx := &finalizerContext{
		votingEpoch: &utils.Epoch{
			Start:  time.Unix(0, 0),
			Period: time.Hour,
		},
		rewardEpoch: &utils.IntEpoch{
			Start:  0,
			Period: 100,
		},
		voterThresholdBIPS: 5000,
	}

	msgChan := make(chan shared.ProtocolMessage, 15)

	payload, err := encodeSubmittedPayload(msgSubmit)
	require.NoError(t, err)

	client := &finalizerClientV2{
		db:                   db,
		relayClient:          relayClient,
		signingPolicyStorage: newSigningPolicyStorage(),
		finalizationStorage:  finalizationStorage,
		messages:             msgChan,
		submissionClient:     NewSubmissionContractClient(submissionContractAddress),
		queueProcessor: newFinalizerQueueProcessorV2(
			db, finalizationStorage, relayClient, fCtx,
		),
		finalizerContext: fCtx,
	}

	msgChan <- shared.ProtocolMessage{ProtocolID: msgSubmit.protocolId, VotingRoundID: msgSubmit.votingRoundId, Message: payload}

	return &testClientsV2{
		db:        db,
		eth:       ethClient,
		finalizer: client,
	}, nil
}

// type testEthClient struct {
// 	calls     int
// 	sentTxs   []*sentTxInfo
// 	mu        sync.RWMutex
// 	sendTxErr error
// }

// type sentTxInfo struct {
// 	privateKey *ecdsa.PrivateKey
// 	to         common.Address
// 	data       []byte
// }

// func (eth *testEthClient) SendRawTx(privateKey *ecdsa.PrivateKey, to common.Address, data []byte, dryRun bool) error {
// 	eth.mu.Lock()
// 	defer eth.mu.Unlock()

// 	eth.calls++

// 	if eth.sendTxErr != nil {
// 		return eth.sendTxErr
// 	}

// 	eth.sentTxs = append(eth.sentTxs, &sentTxInfo{
// 		privateKey: privateKey,
// 		to:         to,
// 		data:       data,
// 	})

// 	return nil
// }

// func (eth *testEthClient) hasAnyCalls() bool {
// 	eth.mu.RLock()
// 	defer eth.mu.RUnlock()
// 	return eth.calls > 0
// }

// type testDB struct {
// 	fetchLogsErr     error
// 	fetchTxsCalls    int
// 	fetchTxsErr      error
// 	mu               sync.Mutex
// 	spiLog           *database.Log
// 	submitterPayload []byte
// }

// func newTestDB(privateKey *ecdsa.PrivateKey) (*testDB, error) {
// 	spiLog, err := newSPILog(privateKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	submitterPayload, err := encodeSubmitterPayload(privateKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &testDB{
// 		spiLog:           spiLog,
// 		submitterPayload: submitterPayload,
// 	}, nil
// }

// func (db *testDB) hasAnyFetchTxsCalls() bool {
// 	db.mu.Lock()
// 	defer db.mu.Unlock()
// 	return db.fetchTxsCalls > 0
// }

// func newSPILog(privateKey *ecdsa.PrivateKey) (*database.Log, error) {
// 	voterAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

// 	spiLog := relay.RelaySigningPolicyInitialized{
// 		RewardEpochId:      big.NewInt(1),
// 		StartVotingRoundId: 1,
// 		Threshold:          1,
// 		Seed:               big.NewInt(1),
// 		Voters:             []common.Address{voterAddress},
// 		Weights:            []uint16{2}, // Weight of 2 > threshold of 1
// 		SigningPolicyBytes: []byte{0x01},
// 		Timestamp:          0,
// 	}

// 	relayABI, err := relay.RelayMetaData.GetAbi()
// 	if err != nil {
// 		return nil, err
// 	}

// 	event, ok := relayABI.Events["SigningPolicyInitialized"]
// 	if !ok {
// 		return nil, errors.New("event not found")
// 	}

// 	inputs := event.Inputs

// 	// Only non-indexed inputs are packed into the log data, indexed inputs are
// 	// stored in the topics.
// 	packedData, err := inputs.NonIndexed().Pack(
// 		spiLog.StartVotingRoundId,
// 		spiLog.Threshold,
// 		spiLog.Seed,
// 		spiLog.Voters,
// 		spiLog.Weights,
// 		spiLog.SigningPolicyBytes,
// 		spiLog.Timestamp,
// 	)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to pack data")
// 	}

// 	hexData := hex.EncodeToString(packedData)

// 	// Integers are serialized as their byte representation, padded to 32 bytes.
// 	topic1 := common.BigToHash(spiLog.RewardEpochId)

// 	log := &database.Log{
// 		BaseEntity: database.BaseEntity{
// 			ID: 1,
// 		},
// 		TransactionID:   1,
// 		Address:         relayContractAddressHex,
// 		Data:            hexData,
// 		Topic0:          topicSPIHex,
// 		Topic1:          topic1.Hex(),
// 		Topic2:          "NULL", // NULL is required for unused topics.
// 		Topic3:          "NULL",
// 		TransactionHash: "0x" + strings.Repeat("ff", 32),
// 		LogIndex:        0,
// 		Timestamp:       0,
// 	}

// 	// Sanity check that the log can be parsed.
// 	relayContract, err := relay.NewRelay(relayContractAddress, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	_, err = relayContract.RelayFilterer.ParseSigningPolicyInitialized(types.Log{
// 		Data: packedData,
// 		Topics: []common.Hash{
// 			common.HexToHash(topicSPIHex),
// 			topic1,
// 		},
// 	})
// 	if err != nil {
// 		return nil, errors.Wrap(err, "ParseSigningPolicyInitialized failed")
// 	}

// 	return log, nil
// }

// func encodeSubmitterPayload(privateKey *ecdsa.PrivateKey) ([]byte, error) {
// 	item := submitterPayloadItem{
// 		protocolId:    0x1,
// 		votingRoundId: 1,
// 		payload: &signedPayload{
// 			typeId: 0x1,
// 			message: &submittedPayload{
// 				protocolId:         0x1,
// 				votingRoundId:      1,
// 				randomQualityScore: true,
// 				merkleRoot:         bytes.Repeat([]byte{0xff}, 32),
// 			},
// 		},
// 	}

// 	buf := new(bytes.Buffer)

// 	// Start with 4 random bytes, these are not parsed by the implementation.
// 	if _, err := buf.Write([]byte{0xde, 0xad, 0xbe, 0xef}); err != nil {
// 		return nil, err
// 	}

// 	if err := buf.WriteByte(item.protocolId); err != nil {
// 		return nil, err
// 	}

// 	if err := binary.Write(buf, binary.BigEndian, item.votingRoundId); err != nil {
// 		return nil, err
// 	}

// 	payloadBytes, err := encodeSignedPayload(privateKey, item.payload)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(payloadBytes) < 104 {
// 		return nil, errors.Errorf("unexpected payload length %d", len(payloadBytes))
// 	}

// 	if err := binary.Write(buf, binary.BigEndian, uint16(len(payloadBytes))); err != nil {
// 		return nil, err
// 	}

// 	_, err = buf.Write(payloadBytes)
// 	return buf.Bytes(), err
// }

// func encodeSignedPayload(privateKey *ecdsa.PrivateKey, payload *signedPayload) ([]byte, error) {
// 	buf := new(bytes.Buffer)
// 	if err := buf.WriteByte(payload.typeId); err != nil {
// 		return nil, err
// 	}

// 	submittedPayloadBytes, err := encodeSubmittedPayload(payload.message)
// 	if err != nil {
// 		return nil, err
// 	}

// 	_, err = buf.Write(submittedPayloadBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	signature, err := signPayload(privateKey, submittedPayloadBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	_, err = buf.Write(signature)
// 	return buf.Bytes(), err
// }

// func signPayload(privateKey *ecdsa.PrivateKey, submittedPayloadBytes []byte) ([]byte, error) {
// 	hash := accounts.TextHash(crypto.Keccak256(submittedPayloadBytes))
// 	signature, err := crypto.Sign(hash, privateKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	logger.Info("signature: %x", signature)

// 	// Signature is in the format [R | S | V], where V is 0 or 1.
// 	// Need to transform to [V | R | S] where V is 27 or 28.

// 	var retSignature [65]byte
// 	retSignature[0] = signature[64] + 27
// 	copy(retSignature[1:], signature[:64])

// 	return retSignature[:], nil
// }

// func encodeSubmittedPayload(payload *submittedPayload) ([]byte, error) {
// 	buf := new(bytes.Buffer)
// 	if err := buf.WriteByte(payload.protocolId); err != nil {
// 		return nil, err
// 	}

// 	if err := binary.Write(buf, binary.BigEndian, payload.votingRoundId); err != nil {
// 		return nil, err
// 	}

// 	if payload.randomQualityScore {
// 		buf.WriteByte(1)
// 	} else {
// 		buf.WriteByte(0)
// 	}

// 	_, err := buf.Write(payload.merkleRoot)
// 	return buf.Bytes(), err
// }

// func (db *testDB) FetchTransactionsByAddressAndSelector(
// 	address common.Address, selector []byte, from, to int64,
// ) ([]database.Transaction, error) {
// 	db.mu.Lock()
// 	defer db.mu.Unlock()

// 	db.fetchTxsCalls++

// 	if db.fetchTxsErr != nil {
// 		return nil, db.fetchTxsErr
// 	}

// 	logger.Debug("fetching transactions from db: address=%s, selector=%x", address.Hex(), selector)
// 	if address != submissionContractAddress {
// 		return nil, errors.New("unknown address")
// 	}

// 	if db.submitterPayload == nil {
// 		return nil, nil
// 	}

// 	submitterPayloadHex := hex.EncodeToString(db.submitterPayload)
// 	db.submitterPayload = nil

// 	return []database.Transaction{{
// 		Input: submitterPayloadHex,
// 	}}, nil
// }

// func (db *testDB) FetchLogsByAddressAndTopic0(
// 	address common.Address, topic string, from, to int64,
// ) ([]database.Log, error) {
// 	db.mu.Lock()
// 	defer db.mu.Unlock()

// 	if db.fetchLogsErr != nil {
// 		return nil, db.fetchLogsErr
// 	}

// 	logger.Debug("fetching logs from db: address=%s, topic=%s", address.Hex(), topic)
// 	if topic != topicSPIHex {
// 		return nil, errors.New("unknown topic")
// 	}

// 	if db.spiLog == nil {
// 		return nil, nil
// 	}

// 	log := *db.spiLog
// 	db.spiLog = nil
// 	return []database.Log{log}, nil
// }
