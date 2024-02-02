package finalizer

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"flare-tlc/config"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/contracts/relay"
	"math/big"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func TestMain(m *testing.M) {
	logger.Configure(config.LoggerConfig{
		Level:   "DEBUG",
		Console: true,
	})

	os.Exit(m.Run())
}

func TestFinalizerClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var ethClient testEthClient

	privateKey, err := crypto.HexToECDSA(testPrivateKeyHex)
	require.NoError(t, err)

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	relayClient, err := NewRelayContractClient(
		nil,
		common.HexToAddress(relayContractAddressHex),
		privateKey,
		fromAddress,
	)
	require.NoError(t, err)
	relayClient.ethClient = &ethClient

	submissionStorage := newSubmissionStorage()

	db, err := newTestDB()
	require.NoError(t, err)

	fCtx := &finalizerContext{
		votingEpoch: &utils.Epoch{
			Start:  time.Unix(0, 0),
			Period: time.Hour,
		},
	}

	client := finalizerClient{
		db:                   db,
		relayClient:          relayClient,
		signingPolicyStorage: newSigningPolicyStorage(),
		submissionStorage:    submissionStorage,
		submissionClient:     NewSubmissionContractClient(common.HexToAddress(submissionContractAddressHex)),
		queueProcessor: newFinalizerQueueProcessor(
			db, submissionStorage, relayClient, fCtx,
		),
		finalizerContext: fCtx,
	}

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return client.RunContext(ctx)
	})

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(ethClient.sentTxs))
}

type testEthClient struct {
	sentTxs []*sentTxInfo
	mu      sync.Mutex
}

type sentTxInfo struct {
	privateKey *ecdsa.PrivateKey
	to         common.Address
	data       []byte
}

func (eth *testEthClient) SendRawTx(privateKey *ecdsa.PrivateKey, to common.Address, data []byte) error {
	eth.mu.Lock()
	defer eth.mu.Unlock()

	eth.sentTxs = append(eth.sentTxs, &sentTxInfo{
		privateKey: privateKey,
		to:         to,
		data:       data,
	})

	return nil
}

type testDB struct{ spiLog *database.Log }

func newTestDB() (*testDB, error) {
	spiLog := relay.RelaySigningPolicyInitialized{
		RewardEpochId:      big.NewInt(1),
		StartVotingRoundId: 1,
		Threshold:          1,
		Seed:               big.NewInt(1),
		Voters:             []common.Address{common.HexToAddress(relayContractAddressHex)},
		Weights:            []uint16{1},
		SigningPolicyBytes: []byte{0x01},
		Timestamp:          0,
	}

	relayABI, err := relay.RelayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topic1 := common.BigToHash(spiLog.RewardEpochId)

	event, ok := relayABI.Events["SigningPolicyInitialized"]
	if !ok {
		return nil, errors.New("event not found")
	}

	inputs := event.Inputs
	logger.Debug("inputs: %+v", inputs)
	packedData, err := inputs[1:].Pack(
		spiLog.StartVotingRoundId,
		spiLog.Threshold,
		spiLog.Seed,
		spiLog.Voters,
		spiLog.Weights,
		spiLog.SigningPolicyBytes,
		spiLog.Timestamp,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack data")
	}

	_, err = inputs.Unpack(packedData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unpack data")
	}

	hexData := hex.EncodeToString(packedData)
	logger.Debug("packed data: %s", hexData)

	log := &database.Log{
		BaseEntity: database.BaseEntity{
			ID: 1,
		},
		TransactionID:   1,
		Address:         relayContractAddressHex,
		Data:            hexData,
		Topic0:          topicSPIHex,
		Topic1:          topic1.Hex(),
		Topic2:          "NULL",
		Topic3:          "NULL",
		TransactionHash: "0x" + strings.Repeat("01", 32),
		LogIndex:        0,
		Timestamp:       0,
	}

	relayContract, err := relay.NewRelay(common.HexToAddress(relayContractAddressHex), nil)
	if err != nil {
		return nil, err
	}

	err = relayABI.UnpackIntoInterface(new(relay.RelaySigningPolicyInitialized), "SigningPolicyInitialized", packedData)
	if err != nil {
		return nil, errors.Wrap(err, "UnpackIntoInterface failed")
	}

	_, err = relayContract.RelayFilterer.ParseSigningPolicyInitialized(types.Log{
		Data: packedData,
		Topics: []common.Hash{
			common.HexToHash(topicSPIHex),
			topic1,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "ParseSigningPolicyInitialized failed")
	}

	return &testDB{spiLog: log}, nil
}

func (db testDB) FetchTransactionsByAddressAndSelector(
	address common.Address, selector []byte, from, to int64,
) ([]database.Transaction, error) {
	logger.Debug("fetching transactions from db: address=%s, selector=%x", address.Hex(), selector)
	return nil, nil
}

func (db testDB) FetchLogsByAddressAndTopic0(
	address common.Address, topic string, from, to int64,
) ([]database.Log, error) {
	logger.Debug("fetching logs from db: address=%s, topic=%s", address.Hex(), topic)
	if topic == topicSPIHex {
		return []database.Log{*db.spiLog}, nil
	}

	return nil, errors.New("unknown topic")
}
