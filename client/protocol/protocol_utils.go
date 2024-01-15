package protocol

import (
	"encoding/hex"
	"encoding/json"
	"flare-tlc/client/config"
	"flare-tlc/client/shared"
	"flare-tlc/logger"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type DataVerifier interface {
	VerifyResponse(*SubProtocolResponse) error
}

type SubProtocol struct {
	Id          uint8
	ApiEndpoint string
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
	}
}

func (sp *SubProtocol) getData(votingRound int64, submitName string, signingAddress string, timeout time.Duration) (*SubProtocolResponse, error) {
	url, err := getUrl(votingRound, sp, submitName, signingAddress)
	if err != nil {
		return nil, errors.Wrap(err, "error getting url")
	}

	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url.String())
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
	signingAddress string,
	nRetries int,
	timeout time.Duration,
	dataVerifier DataVerifier,
) <-chan shared.ExecuteStatus[*SubProtocolResponse] {
	return shared.ExecuteWithRetry(func() (*SubProtocolResponse, error) {
		data, err := sp.getData(votingRound, endpoint, signingAddress, timeout)
		if err == nil {
			err = dataVerifier.VerifyResponse(data)
		}
		if err != nil {
			logger.Info("Error getting data from protocol client with id %d, endpoint %s, voting round %d: %v",
				sp.Id, sp.ApiEndpoint, votingRound, err)
			return nil, err
		}
		return data, nil
	}, nRetries, 0)
}

func getUrl(votingRound int64, protocol *SubProtocol, endpoint string, signingAddress string) (*url.URL, error) {
	baseURL, err := url.JoinPath(
		protocol.ApiEndpoint,
		endpoint,
		strconv.FormatInt(votingRound, 10),
	)
	if err != nil {
		return nil, errors.Wrap(err, "error creating url path")
	}
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.Wrap(err, "error creating url")
	}
	if len(signingAddress) > 0 {
		query := url.Query()
		query.Set("signingAddress", signingAddress)
		url.RawQuery = query.Encode()
	}
	return url, nil
}
