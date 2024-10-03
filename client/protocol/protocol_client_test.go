package protocol

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	clientConfig "flare-tlc/client/config"
	"flare-tlc/config"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
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
	testPrivateKeyHex     = "4f65bffe3c8ed6c0b812e84d35402e949feea042061cc1635fe6ae83ed84df4a"
	submitContractAddress = "0xBB6eae07aD2c5899A081984e31157035b0604106"
)

func TestMain(m *testing.M) {
	logger.Configure(config.LoggerConfig{
		Level:   "DEBUG",
		Console: true,
	})

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

	ethClient := testEthClient{}

	subProtocol := &SubProtocol{Id: 100, ApiEndpoint: apiEndpointURL}

	privKey, err := crypto.HexToECDSA(testPrivateKeyHex)
	require.NoError(t, err)

	address := crypto.PubkeyToAddress(privKey.PublicKey)

	base := SubmitterBase{
		ethClient: &ethClient,
		gasConfig: &clientConfig.GasConfig{GasPriceFixed: common.Big0},
		protocolContext: &protocolContext{
			submitPrivateKey:           privKey,
			signerPrivateKey:           privKey,
			submitSignaturesPrivateKey: privKey,
			submitContractAddress:      common.HexToAddress(submitContractAddress),
			signingAddress:             address,
			submitAddress:              address,
			submitSignaturesAddress:    address,
		},
		epoch:            &utils.Epoch{Start: time.Unix(0, 0), Period: time.Hour},
		subProtocols:     []*SubProtocol{subProtocol},
		submitRetries:    1,
		submitTimeout:    1 * time.Second,
		dataFetchRetries: 1,
		dataFetchTimeout: 1 * time.Second,
		name:             "test",
		submitPrivateKey: privKey,
	}

	t.Run("Submitter", func(t *testing.T) {
		defer ethClient.reset()

		submitter := Submitter{
			SubmitterBase: base,
		}

		epochID := int64(1)
		submitter.RunEpoch(epochID)

		t.Logf("sentTxs: %v", ethClient.sentTxs)
		require.Len(t, ethClient.sentTxs, 1)
		cupaloy.SnapshotT(t, ethClient.sentTxs[0])
	})

	t.Run("SubmitterError", func(t *testing.T) {
		defer ethClient.reset()

		errorStatus := http.StatusInternalServerError
		apiEndpoint.errorStatus = &errorStatus
		defer func() { apiEndpoint.errorStatus = nil }()

		submitter := Submitter{
			SubmitterBase: base,
		}

		epochID := int64(1)
		submitter.RunEpoch(epochID)

		t.Logf("sentTxs: %v", ethClient.sentTxs)
		require.Empty(t, ethClient.sentTxs)
	})

	t.Run("SignatureSubmitter", func(t *testing.T) {
		defer ethClient.reset()

		submitter := SignatureSubmitter{
			SubmitterBase: base,
			maxRounds:     1,
		}

		epochID := int64(1)
		submitter.RunEpoch(epochID)

		t.Logf("sentTxs: %v", ethClient.sentTxs)
		require.Len(t, ethClient.sentTxs, 1)
		cupaloy.SnapshotT(t, ethClient.sentTxs[0])
	})

	t.Run("SignatureSubmitterError", func(t *testing.T) {
		defer ethClient.reset()

		errorStatus := http.StatusInternalServerError
		apiEndpoint.errorStatus = &errorStatus
		defer func() { apiEndpoint.errorStatus = nil }()

		submitter := SignatureSubmitter{
			SubmitterBase: base,
			maxRounds:     1,
		}

		epochID := int64(1)
		submitter.RunEpoch(epochID)

		t.Logf("sentTxs: %v", ethClient.sentTxs)
		require.Empty(t, ethClient.sentTxs)
	})

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled))
}

type testEthClient struct {
	sentTxs []*sentTxInfo
}

func (c *testEthClient) reset() {
	c.sentTxs = nil
}

type sentTxInfo struct {
	privateKey *ecdsa.PrivateKey
	to         common.Address
	payload    []byte
}

func (c *testEthClient) SendRawTx(
	privateKey *ecdsa.PrivateKey, to common.Address, payload []byte, _ *clientConfig.GasConfig, _ time.Duration,
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
	logger.Info("test: handling API request: %+v", r)

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
		logger.Error("test: failed to marshal response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		logger.Error("test: failed to write response: %v", err)
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
	client := ProtocolClient{
		registry:        &registry,
		rewardEpoch:     utils.NewEpoch(time.Now(), 100*time.Millisecond),
		identityAddress: identityAddress,
	}

	err := client.waitUntilRegistered(ctx)
	require.NoError(t, err)

	currentEpoch := client.rewardEpoch.EpochIndex(time.Now())
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
	client := ProtocolClient{
		registry:        &registry,
		rewardEpoch:     utils.NewEpoch(time.Now(), time.Minute),
		identityAddress: identityAddress,
	}

	err := client.waitUntilRegistered(ctx)
	require.NoError(t, err)

	currentEpoch := client.rewardEpoch.EpochIndex(time.Now())
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
