package finalizer

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"flare-tlc/config"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

const (
	testPrivateKeyHex            = "4f65bffe3c8ed6c0b812e84d35402e949feea042061cc1635fe6ae83ed84df4a"
	relayContractAddressHex      = "0xb849b93B585eFfb7cE4B522Ff88d9b3B24955f24"
	submissionContractAddressHex = "0x2F79Dce2375571207a7976148D4468195F89a73e"
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

	var db testDB

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

type testDB struct{}

func (db testDB) FetchTransactionsByAddressAndSelector(
	address common.Address, selector []byte, from, to int64,
) ([]database.Transaction, error) {
	logger.Debug("fetching transactions from db")
	return nil, nil
}

func (db testDB) FetchLogsByAddressAndTopic0(
	address common.Address, topic string, from, to int64,
) ([]database.Log, error) {
	logger.Debug("fetching logs from db")
	return nil, nil
}
