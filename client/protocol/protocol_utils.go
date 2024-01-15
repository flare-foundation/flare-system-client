package protocol

import (
	"encoding/binary"
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

const (
	subProtocolTimeout = 2 * time.Second
)

type SubProtocol struct {
	Id          uint8
	ApiEndpoint string

	client http.Client
}

type DataProviderResponse struct {
	Status         string `json:"status"`
	Data           string `json:"data"`
	AdditionalData string `json:"additionalData"`
}

func NewSubProtocol(config config.ProtocolConfig) *SubProtocol {
	return &SubProtocol{
		Id:          config.Id,
		ApiEndpoint: config.ApiEndpoint,
		client: http.Client{
			Timeout: subProtocolTimeout,
		},
	}
}

func (sp *SubProtocol) getData(votingRound int64, submitName string, signingAddress string) ([]byte, error) {
	url, err := getUrl(votingRound, sp, submitName, signingAddress)
	if err != nil {
		return nil, errors.Wrap(err, "error getting url")
	}

	resp, err := sp.client.Get(url.String())
	if err != nil {
		return nil, errors.Wrap(err, "error calling protocol client API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("protocol client returned status %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading protocol client response")
	}

	var response DataProviderResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.Wrap(err, "cannot parse protocol client response body")
	}

	if response.Status != "OK" {
		return nil, fmt.Errorf("protocol client returned status %v", response.Status)
	}

	bodyString := strings.TrimPrefix(string(response.Data), "0x")
	data, err := hex.DecodeString(bodyString)
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode protocol client response body")
	}

	return data, nil
}

// func (sp *SubProtocol) getDataAsync(votingRound int64, endpoint string, signingAddress string) <-chan []byte {
// 	ch := make(chan []byte)
// 	go func() {
// 		data, err := sp.getData(votingRound, endpoint, signingAddress)
// 		if err != nil {
// 			logger.Info("Error getting data from protocol client with id %d, endpoint %s, voting round %d: %v",
// 				sp.Id, sp.ApiEndpoint, votingRound, err)
// 			ch <- nil
// 			return
// 		}
// 		ch <- data
// 	}()
// 	return ch
// }

func (sp *SubProtocol) getDataWithRetry(votingRound int64, endpoint string, signingAddress string, nRetries int) <-chan shared.ExecuteStatus[[]byte] {
	return shared.ExecuteWithRetry(func() ([]byte, error) {
		data, err := sp.getData(votingRound, endpoint, signingAddress)
		if err != nil {
			logger.Info("Error getting data from protocol client with id %d, endpoint %s, voting round %d: %v",
				sp.Id, sp.ApiEndpoint, votingRound, err)
			return nil, err
		}
		return data, nil
	}, nRetries)
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

func uint16toBytes(i uint16) (arr [2]byte) {
	binary.BigEndian.PutUint16(arr[0:2], i)
	return
}

func uint32toBytes(i uint32) (arr [4]byte) {
	binary.BigEndian.PutUint32(arr[0:4], i)
	return
}
