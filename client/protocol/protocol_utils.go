package protocol

import (
	"encoding/hex"
	"encoding/json"
	"flare-tlc/client/config"
	"flare-tlc/client/shared"
	"flare-tlc/logger"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type DataVerifier func(*SubProtocolResponse) error

type SubProtocol struct {
	Id          uint8
	ApiEndpoint string
	XApiKey     string
}

type SubProtocolResponse struct {
	Status         string `json:"status"`
	Data           []byte `json:"data"`
	AdditionalData []byte `json:"additionalData"`
}

type dataProviderResponse struct {
	Status         string `json:"status"`
	Data           string `json:"data"`
	AdditionalData string `json:"additionalData"`
}

func NewSubProtocol(config config.ProtocolConfig) *SubProtocol {
	return &SubProtocol{
		Id:          config.Id,
		ApiEndpoint: config.ApiEndpoint,
		XApiKey:     config.XApiKey(),
	}
}

func (sp *SubProtocol) getData(votingRound int64, submitName string, submitAddress string, timeout time.Duration) (*SubProtocolResponse, error) {
	url, err := getUrl(votingRound, sp, submitName, submitAddress)
	if err != nil {
		return nil, errors.Wrap(err, "error getting url")
	}

	logger.Info("Calling protocol client API: %s", url.String())
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
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("protocol client returned http status %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading protocol client response")
	}

	var response dataProviderResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.Wrap(err, "cannot parse protocol client response body")
	}

	if response.Status != "OK" {
		return &SubProtocolResponse{Status: resp.Status}, nil
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

func (sp *SubProtocol) getDataWithRetry(
	votingRound int64,
	endpoint string,
	submitAddress string,
	nRetries int,
	timeout time.Duration,
	dataVerifier DataVerifier,
) <-chan shared.ExecuteStatus[*SubProtocolResponse] {
	return shared.ExecuteWithRetry(func() (*SubProtocolResponse, error) {
		data, err := sp.getData(votingRound, endpoint, submitAddress, timeout)
		if err == nil {
			err = dataVerifier(data)
		}
		if err != nil {
			logger.Error("Error getting data from protocol client with id %d, endpoint %s, voting round %d: %v",
				sp.Id, sp.ApiEndpoint, votingRound, err)
			return nil, err
		}
		return data, nil
	}, nRetries, 0)
}

func SignatureSubmitterDataVerifier(data *SubProtocolResponse) error {
	if data.Status != "OK" {
		return fmt.Errorf("status %s", data.Status)
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

func IdentityDataVerifier(data *SubProtocolResponse) error {
	return nil
}

func getUrl(votingRound int64, protocol *SubProtocol, endpoint string, signingAddress string) (*url.URL, error) {
	baseURL, err := url.JoinPath(
		protocol.ApiEndpoint,
		endpoint,
		strconv.FormatInt(votingRound, 10),
		signingAddress,
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
