package protocol

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/flare-foundation/go-flare-common/pkg/payload"
	"github.com/pkg/errors"
)

const maxRespSize = 100 * (1 << 20) // 100 MB for maximal response size

type DataVerifier func(*SubProtocolResponse) error

type SubProtocol struct {
	ID      uint8
	APIUrl  string
	XApiKey string
	Type    uint8 //type of submitSignature payload
}

type SubProtocolResponse struct {
	Status         payload.ResponseStatus `json:"status"`
	Data           []byte                 `json:"data"`
	AdditionalData []byte                 `json:"additionalData"`
}

func NewSubProtocol(config config.ProtocolConfig) *SubProtocol {
	apiUrl := config.APIUrl
	if apiUrl == "" {
		apiUrl = config.APIEndpoint
	}
	return &SubProtocol{
		ID:      config.ID,
		APIUrl:  apiUrl,
		XApiKey: config.XApiKey(),
		Type:    config.Type,
	}
}

// fetchData queries the provider for data for SubProtocol.
func (sp *SubProtocol) fetchData(url *url.URL, timeout time.Duration) (*SubProtocolResponse, error) {
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating protocol client API request")
	}
	if len(sp.XApiKey) > 0 {
		req.Header.Set("X-API-KEY", sp.XApiKey)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error calling protocol client API")
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("protocol client returned http status %v", resp.Status)
	}

	respLimited := &io.LimitedReader{R: resp.Body, N: maxRespSize}

	decoder := json.NewDecoder(respLimited)
	decoder.DisallowUnknownFields()

	var response payload.SubprotocolResponse

	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}

	if response.Status != payload.Ok {
		return &SubProtocolResponse{Status: response.Status}, nil
	}

	bodyString := strings.TrimPrefix(response.Data, "0x")
	data, err := hex.DecodeString(bodyString)
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode protocol client response body")
	}

	var addData []byte
	addDataString := strings.TrimPrefix(response.AdditionalData, "0x")
	if len(addDataString) > 0 {
		addData, err = hex.DecodeString(addDataString)
		if err != nil {
			return nil, errors.Wrap(err, "cannot decode protocol client response additional data")
		}
	}

	return &SubProtocolResponse{
		Status:         response.Status,
		Data:           data,
		AdditionalData: addData,
	}, nil
}

func (sp *SubProtocol) fetchWithVerificationFunc(url *url.URL, timeout time.Duration, dataVerifier DataVerifier) func() (*SubProtocolResponse, error) {
	return func() (*SubProtocolResponse, error) {
		data, err := sp.fetchData(url, timeout)
		if err == nil {
			err = dataVerifier(data)
		}
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

// fetchDataWithRetryChan
func (sp *SubProtocol) fetchDataWithRetryChan(
	votingRound int64,
	endpoint string,
	submitAddress string,
	nRetries int,
	timeout time.Duration,
	dataVerifier DataVerifier,
) <-chan shared.ExecuteStatus[*SubProtocolResponse] {
	url, err := submitEndpointUrl(votingRound, sp.APIUrl, endpoint, submitAddress)
	if err != nil {
		logger.Errorf("building url for protocol %v: %s", sp.ID, err)
		out := make(chan shared.ExecuteStatus[*SubProtocolResponse])
		go func() {
			out <- shared.ExecuteStatus[*SubProtocolResponse]{Success: false, Message: fmt.Sprintf("initial error: %s", err)}
		}()

		return out
	}

	logger.Infof("Calling protocol %v client API: %s", sp.ID, url.String())

	return shared.ExecuteWithRetryChan(sp.fetchWithVerificationFunc(url, timeout, dataVerifier), nRetries, 0)
}

// fetchDataWithRetry
func (sp *SubProtocol) fetchDataWithRetry(
	ctx context.Context,
	votingRound int64,
	endpoint string,
	submitAddress string,
	timeout time.Duration,
	dataVerifier DataVerifier,
	minimalRetryDuration time.Duration,
) shared.ExecuteStatus[*SubProtocolResponse] {
	url, err := submitEndpointUrl(votingRound, sp.APIUrl, endpoint, submitAddress)
	if err != nil {
		logger.Errorf("building url for protocol %v: %v", sp.ID, err)
		return shared.ExecuteStatus[*SubProtocolResponse]{Success: false, Message: fmt.Sprintf("initial error: %s", err)}
	}

	logger.Infof("Calling protocol %v client API: %s", sp.ID, url.String())

	return shared.ExecuteWithRetryWithContext(ctx,
		sp.fetchWithVerificationFunc(url, timeout, dataVerifier),
		minimalRetryDuration)
}

func SignatureSubmitterDataVerifier(data *SubProtocolResponse) error {
	switch data.Status {
	case payload.Ok:
	case payload.Retry:
		return errors.New("retry")
	case payload.Empty:
		return nil
	default:
		return fmt.Errorf("unknown status: %v", data.Status)
	}

	if len(data.Data) != 38 {
		return fmt.Errorf("data length %d is not 38", len(data.Data))
	}
	// Check if additional data is too long
	// Length of data without additional data is 104 bytes: 1 (type) + 38 (message) + 65 (signature)
	if len(data.AdditionalData) > math.MaxUint16-104 {
		return errors.New("additional data too long")
	}
	return nil
}

func StatusDataVerifier(data *SubProtocolResponse) error {
	switch data.Status {
	case payload.Ok:
		return nil
	case payload.Retry:
		return errors.New("retry")
	case payload.Empty:
		return nil
	default:
		return fmt.Errorf("unknown status: %v", data.Status)
	}
}

// submitEndpointUrl builds url to be queried for the data for subprotocol for a given votingRound and address.
func submitEndpointUrl(votingRound int64, apiEndpoint string, endpoint string, address string) (*url.URL, error) {
	baseURL, err := url.JoinPath(
		apiEndpoint,
		endpoint,
		strconv.FormatInt(votingRound, 10),
		address,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error creating url path")
	}
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.Wrap(err, "error creating url")
	}
	// if len(signingAddress) > 0 {
	// 	query := url.Query()
	// 	query.Set("signingAddress", signingAddress)
	// 	url.RawQuery = query.Encode()
	// }
	return url, nil
}
