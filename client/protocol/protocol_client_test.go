package protocol

import (
	"crypto/ecdsa"
	"encoding/json"
	"flare-tlc/config"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
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

	go apiEndpoint.Run()

	ethClient := testEthClient{}

	subProtocol := &SubProtocol{Id: 0, ApiEndpoint: apiEndpointURL}

	privKey, err := crypto.HexToECDSA(testPrivateKeyHex)
	require.NoError(t, err)

	submitter := Submitter{
		SubmitterBase: SubmitterBase{
			ethClient: &ethClient,
			protocolContext: &protocolContext{
				submitPrivateKey:      privKey,
				submitContractAddress: common.HexToAddress(submitContractAddress),
			},
			epoch:         &utils.Epoch{Start: time.Unix(0, 0), Period: time.Hour},
			subProtocols:  []*SubProtocol{subProtocol},
			submitRetries: 1,
			name:          "test",
		},
	}

	epochID := int64(1)
	submitter.RunEpoch(epochID)

	t.Logf("sentTxs: %v", ethClient.sentTxs)
	require.Len(t, ethClient.sentTxs, 1)
	cupaloy.SnapshotT(t, ethClient.sentTxs[0])
}

type testEthClient struct {
	sentTxs []*sentTxInfo
}

type sentTxInfo struct {
	privateKey *ecdsa.PrivateKey
	to         common.Address
	payload    []byte
}

func (c *testEthClient) SendRawTx(
	privateKey *ecdsa.PrivateKey, to common.Address, payload []byte,
) error {
	c.sentTxs = append(c.sentTxs, &sentTxInfo{
		privateKey: privateKey,
		to:         to,
		payload:    payload,
	})
	return nil
}

type testAPIEndpoint struct {
	listener net.Listener
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

func (ep *testAPIEndpoint) Run() error {
	http.Handle("/", ep)
	return http.Serve(ep.listener, nil)
}

func (ep *testAPIEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Info("test: handling API request: %+v", r)

	rsp := dataProviderResponse{
		Status:         "OK",
		Data:           "0x1234",
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
