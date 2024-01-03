package clients

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"flare-tlc/client/config"
	globalConfig "flare-tlc/config"
	"flare-tlc/logger"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/submission"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type VotingClient struct {
	subProtocols      []SubProtocol
	eth               *ethclient.Client
	privateKey        string
	contractAddresses globalConfig.ContractAddresses
	contractSelectors contractSelectors
}

type SubProtocol struct {
	config config.ProtocolConfig
}

func NewSubProtocol(config config.ProtocolConfig) SubProtocol {
	return SubProtocol{
		config: config,
	}
}

var currentPriceEpoch = 1452560

func NewVotingClient(config *config.ClientConfig) (*VotingClient, error) {
	subProtocols := []SubProtocol{
		NewSubProtocol(config.Ftso),
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

	privateKey, err := chainConfig.GetPrivateKey()
	if err != nil {
		return nil, err
	}

	return &VotingClient{
		eth:               cl,
		subProtocols:      subProtocols,
		privateKey:        privateKey,
		contractAddresses: config.ContractAddresses,
		contractSelectors: newSelectors(),
	}, nil
}

type contractSelectors struct {
	commit []byte
	reveal []byte
	result []byte
}

func newSelectors() contractSelectors {
	submissionABI, err := abi.JSON(strings.NewReader(submission.SubmissionMetaData.ABI))
	if err != nil {
		panic(err)
	}
	return contractSelectors{
		commit: submissionABI.Methods["commit"].ID,
		reveal: submissionABI.Methods["reveal"].ID,
		result: submissionABI.Methods["result"].ID,
	}
}

// TODO: Read from smart contract or config
const epochDuration = 30 * time.Second
const revealPeriod = 15 * time.Second

func (c *VotingClient) Run() error {
	ticker := time.NewTicker(epochDuration)

	// TODO: Make sure to start at the start of a voting round.
	// 		 Requires getting epoch settings from smart contracts.
	for startTime := time.Now(); true; startTime = <-ticker.C {
		fmt.Println("Starting voting round: ", currentPriceEpoch)

		if err := processCommits(c); err != nil {
			return errors.Wrap(err, "error processing commits")
		}

		if err := processReveals(c); err != nil {
			return errors.Wrap(err, "error processing reveals")
		}

		// Wait until reveal deadline
		revealDeadline := startTime.Add(revealPeriod)
		// TODO: Probably a better way of achieving this than sleeping?
		time.Sleep(time.Until(revealDeadline))

		// Get results
		if err := processResults(c); err != nil {
			return errors.Wrap(err, "error processing results")
		}
		currentPriceEpoch++
	}

	return nil
}

// Calldata format:
//
// - 4 bytes: Solidity function selector
// Followed by for each sub-protocol:
// - 1 byte: protocol id
// - 2 bytes: length of data
// - n bytes: data
func processCommits(c *VotingClient) error {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.contractSelectors.commit)

	for _, protocol := range c.subProtocols {
		commitData, err := getCommitData(currentPriceEpoch, protocol)
		if err != nil {
			return err
		}

		buffer.WriteByte(protocol.config.Id)
		// TODO: Handle overflow errors
		lengthBytes := uint16toBytes(uint16(len(commitData)))
		buffer.Write(lengthBytes[:])
		buffer.Write(commitData)

		fmt.Println("Total encoded:", buffer.Len())
	}
	commitPayload := buffer.Bytes()
	fmt.Println("Submitting commit payload:", len(commitPayload))

	return chain.SendRawTx(*c.eth, c.privateKey, c.contractAddresses.Submission, commitPayload)
}

func processReveals(c *VotingClient) error {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.contractSelectors.reveal)
	for _, protocol := range c.subProtocols {
		revealData, err := getRevealData(currentPriceEpoch, protocol)
		if err != nil {
			return errors.Wrap(err, "processReveals: error getting reveal data")
		}

		buffer.WriteByte(protocol.config.Id)
		length := uint16toBytes(uint16(len(revealData)))
		buffer.Write(length[:])
		buffer.Write(revealData)

		fmt.Println("Total encoded:", buffer.Len())
	}
	revealPayload := buffer.Bytes()
	fmt.Println("Submitting reveal payload:", len(revealPayload))
	return chain.SendRawTx(*c.eth, c.privateKey, c.contractAddresses.Submission, revealPayload)
}

func processResults(c *VotingClient) error {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.contractSelectors.result)
	for _, protocol := range c.subProtocols {
		resultData, err := getResultsData(currentPriceEpoch, protocol)
		if err != nil {
			return errors.Wrap(err, "processResults: error getting result data")
		}

		buffer.WriteByte(protocol.config.Id)
		length := uint16toBytes(uint16(len(resultData)))
		buffer.Write(length[:])
		buffer.Write(resultData)

		fmt.Println("Total encoded:", buffer.Len())
	}
	resultPayload := buffer.Bytes()
	fmt.Println("Submitting result payload:", len(resultPayload))
	return chain.SendRawTx(*c.eth, c.privateKey, c.contractAddresses.Submission, resultPayload)
}

// TODO: Handle error status codes
func getCommitData(votingRound int, protocol SubProtocol) ([]byte, error) {
	url, err := url.JoinPath(protocol.config.ApiEndpoint, fmt.Sprintf("commit/%d", votingRound))
	if err != nil {
		return nil, errors.Wrap(err, "error joining url path")
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error calling commit API")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading commit response")
	}

	bodyString := strings.TrimPrefix(string(body), "0x")
	commitHash, _ := hex.DecodeString(bodyString)

	if len(commitHash) != 32 {
		return nil, errors.New("merkle root is not 32 bytes long")
	}
	fmt.Println("Commit data: ", string(body), len(body), len(commitHash))

	return commitHash, nil
}

// TODO: Handle error status codes
func getRevealData(votingRound int, protocol SubProtocol) ([]byte, error) {
	url := fmt.Sprint(protocol.config.ApiEndpoint, "reveal/", votingRound)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error calling reveal API")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading reveal response")
	}

	bodyString := strings.TrimPrefix(string(body), "0x")
	revealData, err := hex.DecodeString(bodyString)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding reveal data")
	}

	fmt.Println("Reveal data: ", bodyString)

	return revealData, nil
}

func getResultsData(votingRound int, protocol SubProtocol) ([]byte, error) {
	url, err := url.JoinPath(protocol.config.ApiEndpoint, fmt.Sprintf("result/%d", votingRound))
	if err != nil {
		return nil, errors.Wrap(err, "error joining url path")
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error calling results API")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading results response")
	}

	bodyString := strings.TrimPrefix(string(body), "0x")
	merkleRoot, _ := hex.DecodeString(bodyString)

	fmt.Println("Result data: ", bodyString)

	if len(merkleRoot) != 32 {
		return nil, errors.New("merkle root is not 32 bytes long")
	}
	fmt.Println("Result data: ", hex.EncodeToString(merkleRoot))

	return merkleRoot, nil
}

func uint16toBytes(i uint16) (arr [2]byte) {
	binary.LittleEndian.PutUint16(arr[0:2], i)
	return
}
