package protocol

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	clientConfig "github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

const (
	testPrivateKeyHex     = "4f65bffe3c8ed6c0b812e84d35402e949feea042061cc1635fe6ae83ed84df4a"
	submitContractAddress = "0xBB6eae07aD2c5899A081984e31157035b0604106"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSubmitter(t *testing.T) {
	var apiEndpoint testAPIEndpoint
	err := apiEndpoint.Listen()
	require.NoError(t, err)
	defer apiEndpoint.Close()

	apiEndpointURL := apiEndpoint.URL()
	t.Logf("apiEndpointURL: %v", apiEndpointURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error { return apiEndpoint.Run(ctx) })

	chainClient := testChainClient{}

	subProtocol := &SubProtocol{ID: 100, APIUrl: apiEndpointURL, Type: 0}

	privKey, err := crypto.HexToECDSA(testPrivateKeyHex)
	require.NoError(t, err)

	address := crypto.PubkeyToAddress(privKey.PublicKey)

	base := SubmitterBase{
		chainClient: &chainClient,
		gasConfig:   &clientConfig.Gas{GasPriceFixed: common.Big0},
		protocolContext: &protocolContext{
			submitPrivateKey:           privKey,
			signerPrivateKey:           privKey,
			submitSignaturesPrivateKey: privKey,
			submitContractAddress:      common.HexToAddress(submitContractAddress),
			signingAddress:             address,
			submitAddress:              address,
			submitSignaturesAddress:    address,
		},
		votingRoundTiming: &utils.EpochTimingConfig{Start: time.Unix(0, 0), Period: time.Hour},
		subProtocols:      []*SubProtocol{subProtocol},
		submitRetries:     1,
		submitTimeout:     1 * time.Second,
		dataFetchRetries:  1,
		dataFetchTimeout:  1 * time.Second,
		name:              "test",
		submitPrivateKey:  privKey,
	}

	t.Run("Submitter", func(t *testing.T) {
		defer chainClient.reset()

		submitter := Submitter{
			SubmitterBase: base,
		}

		epochID := int64(1)
		submitter.RunEpoch(epochID)

		t.Logf("sentTxs: %v", chainClient.sentTxs)
		require.Len(t, chainClient.sentTxs, 1)
		cupaloy.SnapshotT(t, chainClient.sentTxs[0])
	})

	t.Run("SubmitterError", func(t *testing.T) {
		defer chainClient.reset()

		errorStatus := http.StatusInternalServerError
		apiEndpoint.errorStatus = &errorStatus
		defer func() { apiEndpoint.errorStatus = nil }()

		submitter := Submitter{
			SubmitterBase: base,
		}

		epochID := int64(1)
		submitter.RunEpoch(epochID)

		t.Logf("sentTxs: %v", chainClient.sentTxs)
		require.Empty(t, chainClient.sentTxs)
	})

	t.Run("SignatureSubmitterType0", func(t *testing.T) {
		msgChan := make(chan<- shared.ProtocolMessage, 10)
		defer close(msgChan)

		defer chainClient.reset()

		submitter := SignatureSubmitter{
			SubmitterBase:  base,
			messageChannel: msgChan,
			maxCycles:      1,
			cycleDuration:  time.Second,
		}

		epochID := int64(1)
		submitter.RunEpoch(epochID)

		t.Logf("sentTxs: %v", chainClient.sentTxs)
		require.Len(t, chainClient.sentTxs, 1)

		cupaloy.SnapshotT(t, chainClient.sentTxs[0])
	})

	t.Run("SignatureSubmitterType1", func(t *testing.T) {
		msgChan := make(chan<- shared.ProtocolMessage, 10)
		defer close(msgChan)

		defer chainClient.reset()

		submitter := SignatureSubmitter{
			SubmitterBase:  base,
			messageChannel: msgChan,
			maxCycles:      1,
			cycleDuration:  time.Second,
		}
		subProtocolType1 := &SubProtocol{ID: 100, APIUrl: apiEndpointURL, Type: 1}
		submitter.SubmitterBase.subProtocols = []*SubProtocol{subProtocolType1}

		epochID := int64(1)
		submitter.RunEpoch(epochID)

		t.Logf("sentTxs: %v", chainClient.sentTxs)
		require.Len(t, chainClient.sentTxs, 1)

		cupaloy.SnapshotT(t, chainClient.sentTxs[0])
	})

	t.Run("SignatureSubmitterError", func(t *testing.T) {
		msgChan := make(chan<- shared.ProtocolMessage, 10)
		defer close(msgChan)

		defer chainClient.reset()

		errorStatus := http.StatusInternalServerError
		apiEndpoint.errorStatus = &errorStatus
		defer func() { apiEndpoint.errorStatus = nil }()

		submitter := SignatureSubmitter{
			SubmitterBase:  base,
			messageChannel: msgChan,
			maxCycles:      1,
			cycleDuration:  time.Second,
		}

		epochID := int64(1)
		submitter.RunEpoch(epochID)

		t.Logf("sentTxs: %v", chainClient.sentTxs)
		require.Empty(t, chainClient.sentTxs)
	})

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled))
}

type testChainClient struct {
	sentTxs []*sentTxInfo
}

func (c *testChainClient) reset() {
	c.sentTxs = nil
}

type sentTxInfo struct {
	privateKey *ecdsa.PrivateKey
	to         common.Address
	payload    []byte
}

func (c *testChainClient) SendRawTx(
	privateKey *ecdsa.PrivateKey, to common.Address, payload []byte, _ *clientConfig.Gas, _ time.Duration,
) error {
	c.sentTxs = append(c.sentTxs, &sentTxInfo{
		privateKey: privateKey,
		to:         to,
		payload:    payload,
	})
	return nil
}

type testAPIEndpoint struct {
	listener    net.Listener
	errorStatus *int
}

func (ep *testAPIEndpoint) Listen() error {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return err
	}

	ep.listener = l
	return nil
}

func (ep *testAPIEndpoint) Close() error {
	if ep.listener == nil {
		return nil
	}

	return ep.listener.Close()
}

func (ep *testAPIEndpoint) URL() string {
	u := url.URL{
		Scheme: "http",
		Host:   ep.listener.Addr().String(),
	}

	return u.String()
}

func (ep *testAPIEndpoint) Run(ctx context.Context) error {
	s := http.Server{Handler: ep}

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			return err
		}

		return ctx.Err()
	})

	err := s.Serve(ep.listener)
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return eg.Wait()
}

func (ep *testAPIEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Infof("test: handling API request: %+v", r)

	if ep.errorStatus != nil {
		http.Error(w, "test: error", *ep.errorStatus)
		return
	}

	rsp := dataProviderResponse{
		Status:         "OK",
		Data:           "0x" + strings.Repeat("ff", 38),
		AdditionalData: "0x1234",
	}

	data, err := json.Marshal(rsp)
	if err != nil {
		logger.Errorf("test: failed to marshal response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		logger.Errorf("test: failed to write response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("test: response sent")
}

var identityAddress = common.HexToAddress("0x26B40970948D74d60f37911d1276fF940D8648a4")

func TestWaitUntilRegistered(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	registry := testRegistry{
		expectedAddress: identityAddress,
		registeredEpoch: 3,
	}
	client := client{
		registry:          &registry,
		rewardEpochTiming: utils.NewEpochConfig(time.Now(), 100*time.Millisecond),
		identityAddress:   identityAddress,
	}

	err := client.waitUntilRegistered(ctx)
	require.NoError(t, err)

	currentEpoch := client.rewardEpochTiming.EpochIndex(time.Now())
	require.GreaterOrEqual(t, currentEpoch, registry.registeredEpoch)
}

func TestWaitUntilRegisteredTransientError(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	registry := testRegistry{
		expectedAddress:     identityAddress,
		registeredEpoch:     0,
		transientErrorCount: 3,
	}
	client := client{
		registry:          &registry,
		rewardEpochTiming: utils.NewEpochConfig(time.Now(), time.Minute),
		identityAddress:   identityAddress,
	}

	err := client.waitUntilRegistered(ctx)
	require.NoError(t, err)

	currentEpoch := client.rewardEpochTiming.EpochIndex(time.Now())
	require.GreaterOrEqual(t, currentEpoch, registry.registeredEpoch)
}

type testRegistry struct {
	expectedAddress     common.Address
	registeredEpoch     int64
	transientErrorCount int
}

func (r *testRegistry) IsVoterRegistered(ctx context.Context, address common.Address, epoch int64) (bool, error) {
	if address != r.expectedAddress {
		return false, errors.New("unknown address")
	}

	if r.transientErrorCount > 0 {
		r.transientErrorCount--
		return false, errors.New("transient error")
	}

	return epoch >= r.registeredEpoch, nil

}
