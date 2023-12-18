package clients

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	localContext "flare-tlc/client/context"
	"flare-tlc/config"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/submission"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/ethclient"
)

type VotingClient struct {
	subProtocols      []subProtocol
	eth               *ethclient.Client
	privateKey        string
	contractAddresses config.ContractAddresses
	contractSelectors contractSelectors
}

type subProtocol struct {
	// TODO: Should use 2 bytes for id?
	id uint8
	// E.g. localhost:3000/ftso/price-controller/
	apiEndpoint string
}

var currentPriceEpoch = 1452560

func NewVotingClient(ctx localContext.ClientContext) *VotingClient {
	// TODO: Move to config
	subProtocols := []subProtocol{
		{
			id:          1,
			apiEndpoint: "http://localhost:3000/ftso/price-controller/",
		},
		{
			id:          2,
			apiEndpoint: "http://localhost:3000/ftso/price-controller/",
		},
	}

	config := ctx.Config().ChainConfig()
	cl, err := config.DialETH()

	if err != nil {
		// TODO: repalce all panic(error) with proper error handling
		panic(err)
	}

	bn, err := cl.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	// TODO: Use a logger everywhere instead of fmt.Println
	fmt.Println("Connected to chain, current block number: ", bn)

	privateKey, err := ctx.Config().ChainConfig().GetPrivateKey()
	if err != nil {
		panic(err)
	}

	return &VotingClient{
		eth:               cl,
		subProtocols:      subProtocols,
		privateKey:        privateKey,
		contractAddresses: ctx.Config().ContractAddresses,
		contractSelectors: newSelectors(),
	}
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

		processCommits(c)
		processReveals(c)

		// Wait until reveal deadline
		revealDeadline := startTime.Add(revealPeriod)
		// TODO: Probably a better way of achieving this than sleeping?
		time.Sleep(time.Until(revealDeadline))

		// Get results
		processResults(c)
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
func processCommits(c *VotingClient) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.contractSelectors.commit)

	for _, protocol := range c.subProtocols {
		commitData := getCommitData(currentPriceEpoch, protocol)

		buffer.WriteByte(protocol.id)
		// TODO: Handle overflow errors
		length := uint16toBytes(uint16(len(commitData)))
		buffer.Write(length[:])
		buffer.Write(commitData)

		fmt.Println("Total encoded:", buffer.Len())
	}
	commitPayload := buffer.Bytes()
	fmt.Println("Submitting commit payload:", len(commitPayload))

	chain.SendRawTx(*c.eth, c.privateKey, c.contractAddresses.Submission, commitPayload)
}

func processReveals(c *VotingClient) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.contractSelectors.reveal)
	for _, protocol := range c.subProtocols {
		revealData := getRevealData(currentPriceEpoch, protocol)

		buffer.WriteByte(protocol.id)
		length := uint16toBytes(uint16(len(revealData)))
		buffer.Write(length[:])
		buffer.Write(revealData)

		fmt.Println("Total encoded:", buffer.Len())
	}
	revealPayload := buffer.Bytes()
	fmt.Println("Submitting reveal payload:", len(revealPayload))
	chain.SendRawTx(*c.eth, c.privateKey, c.contractAddresses.Submission, revealPayload)
}

func processResults(c *VotingClient) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.contractSelectors.result)
	for _, protocol := range c.subProtocols {
		resultData := getResultsData(currentPriceEpoch, protocol)

		buffer.WriteByte(protocol.id)
		length := uint16toBytes(uint16(len(resultData)))
		buffer.Write(length[:])
		buffer.Write(resultData)

		fmt.Println("Total encoded:", buffer.Len())
	}
	resultPayload := buffer.Bytes()
	fmt.Println("Submitting result payload:", len(resultPayload))
	chain.SendRawTx(*c.eth, c.privateKey, c.contractAddresses.Submission, resultPayload)
}

// TODO: Handle error status codes
func getCommitData(votingRound int, protocol subProtocol) []byte {
	url := fmt.Sprint(protocol.apiEndpoint, "commit/", votingRound)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error calling commit API:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading commit response:", err)
		return nil
	}

	bodyString := strings.TrimPrefix(string(body), "0x")
	commitHash, _ := hex.DecodeString(bodyString)

	if len(commitHash) != 32 {
		panic(errors.New("merkle root is not 32 bytes long"))
	}
	fmt.Println("Commit data: ", string(body), len(body), len(commitHash))

	return commitHash
}

// TODO: Handle error status codes
func getRevealData(votingRound int, protocol subProtocol) []byte {
	url := fmt.Sprint(protocol.apiEndpoint, "reveal/", votingRound)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error calling reveal API:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading reveal response:", err)
		return nil
	}

	bodyString := strings.TrimPrefix(string(body), "0x")
	revealData, _ := hex.DecodeString(bodyString)

	fmt.Println("Reveal data: ", bodyString)

	return revealData
}

// TODO: Handle error status codes
func getResultsData(votingRound int, protocol subProtocol) []byte {
	url := fmt.Sprint(protocol.apiEndpoint, "result/", votingRound)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error calling results API:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading results response:", err)
		return nil
	}

	bodyString := strings.TrimPrefix(string(body), "0x")
	merkleRoot, _ := hex.DecodeString(bodyString)

	fmt.Println("Result data: ", bodyString)

	if len(merkleRoot) != 32 {
		panic(errors.New("merkle root is not 32 bytes long"))
	}
	fmt.Println("Result data: ", hex.EncodeToString(merkleRoot))

	return merkleRoot
}

func uint16toBytes(i uint16) (arr [2]byte) {
	binary.LittleEndian.PutUint16(arr[0:2], i)
	return
}
