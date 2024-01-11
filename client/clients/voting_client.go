package clients

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"flare-tlc/client/config"
	clientContext "flare-tlc/client/context"
	"flare-tlc/client/shared"
	globalConfig "flare-tlc/config"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/submission"
	"flare-tlc/utils/contracts/system"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type VotingClient struct {
	subProtocols      []SubProtocol
	eth               *ethclient.Client
	submitKey         string
	signingKey        string
	contractAddresses globalConfig.ContractAddresses
	contractSelectors contractSelectors
	epochSettings     shared.EpochSettings
	systemManager     *system.FlareSystemManager
}

type SubProtocol struct {
	config config.ProtocolConfig
}

func NewSubProtocol(config config.ProtocolConfig) SubProtocol {
	return SubProtocol{
		config: config,
	}
}

func NewVotingClient(ctx clientContext.ClientContext) (*VotingClient, error) {
	config := ctx.Config()

	if !config.Voting.EnabledProtocolVoting {
		return nil, nil
	}

	subProtocols := []SubProtocol{
		NewSubProtocol(config.Protocol),
	}

	chainConfig := config.ChainConfig()
	cl, err := chainConfig.DialETH()

	if err != nil {
		return nil, err
	}

	bn, err := cl.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to chain, current block number: %d", bn)

	submitKey, err := chainConfig.GetSubmissionKey()
	if err != nil {
		return nil, err
	}

	signingKey, err := chainConfig.GetSigningKey()
	if err != nil {
		return nil, err
	}
	systemManager, _ := system.NewFlareSystemManager(ctx.Config().ContractAddresses.SystemManager, cl)

	rewardEpochStart, _ := systemManager.RewardEpochsStartTs(nil)
	rewardEpochDuration, _ := systemManager.RewardEpochDurationSeconds(nil)
	firstVotingEpochStart, _ := systemManager.FirstVotingRoundStartTs(nil)
	votingEpochDuration, _ := systemManager.VotingEpochDurationSeconds(nil)

	epochSettings := shared.EpochSettings{
		RewardEpochStartSec:      rewardEpochStart,
		RewardEpochDurationSec:   rewardEpochDuration,
		FirstVotingEpochStartSec: firstVotingEpochStart,
		VotingEpochDurationSec:   votingEpochDuration,
	}

	selectors := newSelectors()

	fmt.Println("Selectors :", hex.EncodeToString(selectors.commit), hex.EncodeToString(selectors.reveal), hex.EncodeToString(selectors.sign))

	return &VotingClient{
		eth:               cl,
		subProtocols:      subProtocols,
		submitKey:         submitKey,
		signingKey:        signingKey,
		contractAddresses: config.ContractAddresses,
		contractSelectors: selectors,
		epochSettings:     epochSettings,
		systemManager:     systemManager,
	}, nil
}

type contractSelectors struct {
	commit []byte
	reveal []byte
	sign   []byte
}

func newSelectors() contractSelectors {
	submissionABI, err := abi.JSON(strings.NewReader(submission.SubmissionMetaData.ABI))
	if err != nil {
		panic(err)
	}
	return contractSelectors{
		commit: submissionABI.Methods["submit1"].ID,
		reveal: submissionABI.Methods["submit2"].ID,
		sign:   submissionABI.Methods["submitSignatures"].ID,
	}
}

func (c *VotingClient) Run() error {
	signingAddress, err := chain.PrivateKeyToEthAddress(c.signingKey)
	if err != nil {
		return errors.Wrap(err, "error getting signing address")
	}

	for {
		votingEpochStart := c.epochSettings.NextVotingEpochStart(time.Now())
		fmt.Println("Next epoch starts at:", votingEpochStart)
		time.Sleep(time.Until(votingEpochStart))

		currentVotingEpoch := int(c.epochSettings.VotingEpochForTime(time.Now()))
		fmt.Println("Current voting epoch:", currentVotingEpoch, "time", time.Now())
		previousVotingEpoch := currentVotingEpoch - 1

		fmt.Println("Processing epoch:", currentVotingEpoch)
		if err := processCommits(c, currentVotingEpoch, signingAddress); err != nil {
			return errors.Wrap(err, "error processing commits")
		}

		fmt.Println("Processing reveals:", previousVotingEpoch, time.Now())
		if err := processReveals(c, previousVotingEpoch, signingAddress); err != nil {
			logger.Info("Error processing reveals, aborting voting round: %v", err)
			continue
		}

		// Wait until reveal deadline
		revealDeadline := votingEpochStart.Add(time.Duration(c.epochSettings.VotingEpochDurationSec / 2))
		time.Sleep(time.Until(revealDeadline))

		// Get results
		if err := processResults(c, previousVotingEpoch, signingAddress); err != nil {
			return errors.Wrap(err, "error processing results")
		}
	}
}

// Calldata format:
//
// - 4 bytes: Solidity function selector
// Followed by for each sub-protocol:
// - 1 byte: protocol id
// - 2 bytes: length of data
// - n bytes: data
func processCommits(c *VotingClient, currentPriceEpoch int, signingAddress common.Address) error {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.contractSelectors.commit)

	for _, protocol := range c.subProtocols {
		commitData, err := getCommitData(currentPriceEpoch, protocol, signingAddress.Hex())
		if err != nil {
			return err
		}
		buffer.Write(commitData)

		fmt.Println("Total encoded:", buffer.Len())
	}
	commitPayload := buffer.Bytes()
	fmt.Println("Submitting commit payload:", len(commitPayload))

	return chain.SendRawTx(*c.eth, c.submitKey, c.contractAddresses.Submission, commitPayload)
}

func processReveals(c *VotingClient, previousVotingEpoch int, signingAddress common.Address) error {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.contractSelectors.reveal)

	for _, protocol := range c.subProtocols {
		revealData, err := getRevealData(previousVotingEpoch, protocol, signingAddress.Hex())
		fmt.Println("Reveal data:", revealData, "error:", err)
		if err != nil {
			return errors.Wrap(err, "processReveals: error getting reveal data")
		}

		buffer.Write(revealData)

		fmt.Println("Total encoded:", buffer.Len())
	}
	revealPayload := buffer.Bytes()
	fmt.Println("Submitting reveal payload:", len(revealPayload))
	return chain.SendRawTx(*c.eth, c.submitKey, c.contractAddresses.Submission, revealPayload)
}

func processResults(c *VotingClient, previousVotingEpoch int, signingAddress common.Address) error {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.contractSelectors.sign)

	for _, protocol := range c.subProtocols {
		resultData, err := getResultsData(previousVotingEpoch, protocol, signingAddress.Hex())
		if err != nil {
			return errors.Wrap(err, "processResults: error getting result data")
		}

		key, err := crypto.HexToECDSA(c.signingKey)
		if err != nil {
			return errors.Wrap(err, "processResults: error getting signing key")
		}

		hash := accounts.TextHash(crypto.Keccak256(resultData))
		signature, err := crypto.Sign(hash, key)
		if err != nil {
			return errors.Wrap(err, "processResults: error signing result data")
		}
		fmt.Println("Hash to sign:", hex.EncodeToString(hash))
		fmt.Println("Signature v:", signature[0])
		fmt.Println("Signature r:", hex.EncodeToString(signature[1:33]))
		fmt.Println("Signature s:", hex.EncodeToString(signature[33:65]))

		epochBytes := uint32toBytes(uint32(previousVotingEpoch))
		lengthBytes := uint16toBytes(104)

		buffer.WriteByte(100)        // Protocol ID (1 byte)
		buffer.Write(epochBytes[:])  // Epoch (4 bytes)
		buffer.Write(lengthBytes[:]) // Length (2 bytes)

		buffer.WriteByte(0)      // Type (1 byte)
		buffer.Write(resultData) // Message (38 bytes)

		buffer.WriteByte(signature[64] + 27) // V (1 byte)
		buffer.Write(signature[0:32])        // R (32 bytes)
		buffer.Write(signature[32:64])       // S (32 bytes)

		fmt.Println("Total encoded:", buffer.Len())
	}
	resultPayload := buffer.Bytes()
	fmt.Println("Submitting result payload:", len(resultPayload))
	return chain.SendRawTx(*c.eth, c.submitKey, c.contractAddresses.Submission, resultPayload)
}

type DataProviderResponse struct {
	Status         string `json:"status"`
	Data           string `json:"data"`
	AdditionalData string `json:"additionalData"`
}

// TODO: Handle error status codes
func getCommitData(votingRound int, protocol SubProtocol, signingAddress string) ([]byte, error) {
	url, err := getUrl(votingRound, protocol, "submit1", signingAddress)
	if err != nil {
		return nil, errors.Wrap(err, "error getting url")
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, errors.Wrap(err, "error calling commit API")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading commit response")
	}

	var response DataProviderResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.Wrap(err, "cannot parse response JSON")
	}

	bodyString := strings.TrimPrefix(string(response.Data), "0x")
	data, _ := hex.DecodeString(bodyString)

	fmt.Println("Commit data: ", string(body), len(body), len(data))

	return data, nil
}

func getUrl(votingRound int, protocol SubProtocol, endpoint string, signingAddress string) (*url.URL, error) {
	path := fmt.Sprintf("%s/%d", endpoint, votingRound)
	baseURL, err := url.JoinPath(protocol.config.ApiEndpoint, path)
	if err != nil {
		return nil, errors.Wrap(err, "error joining url path")
	}
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing url")
	}

	query := url.Query()
	query.Set("signingAddress", signingAddress)
	url.RawQuery = query.Encode()

	return url, nil
}

// TODO: Handle error status codes
func getRevealData(votingRound int, protocol SubProtocol, signingAddress string) ([]byte, error) {
	url, err := getUrl(votingRound, protocol, "submit2", signingAddress)
	if err != nil {
		return nil, errors.Wrap(err, "error getting url")
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, errors.Wrap(err, "error calling commit API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error getting reveal data: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading reveal response")
	}

	var response DataProviderResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.Wrap(err, "cannot parse response JSON")
	}

	bodyString := strings.TrimPrefix(string(response.Data), "0x")
	revealData, err := hex.DecodeString(bodyString)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding reveal data")
	}

	fmt.Println("Reveal data: ", revealData)

	return revealData, nil
}

func getResultsData(votingRound int, protocol SubProtocol, signingAddress string) ([]byte, error) {
	url, err := getUrl(votingRound, protocol, "submitSignatures", signingAddress)
	if err != nil {
		return nil, errors.Wrap(err, "error getting url")
	}

	fmt.Println("Getting results from:", url.String())
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, errors.Wrap(err, "error calling commit API")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading results response")
	}

	var response DataProviderResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.Wrap(err, "cannot parse response JSON")
	}

	bodyString := strings.TrimPrefix(string(response.Data), "0x")
	resultData, _ := hex.DecodeString(bodyString)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding result data")
	}

	fmt.Println("Result data: ", hex.EncodeToString(resultData))
	return resultData, nil
}

func uint16toBytes(i uint16) (arr [2]byte) {
	binary.BigEndian.PutUint16(arr[0:2], i)
	return
}

func uint32toBytes(i uint32) (arr [4]byte) {
	binary.BigEndian.PutUint32(arr[0:4], i)
	return
}
