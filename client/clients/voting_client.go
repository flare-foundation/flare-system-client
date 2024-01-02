package clients

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	localContext "flare-tlc/client/context"
	"flare-tlc/client/shared"
	"flare-tlc/config"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/submission"
	"flare-tlc/utils/contracts/system"
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
	epochSettings     shared.EpochSettings
	systemManager     *system.FlareSystemManager
}

type subProtocol struct {
	// TODO: Should use 2 bytes for id?
	id uint8
	// E.g. localhost:3000/ftso/price-controller/
	apiEndpoint string
}

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
		privateKey:        privateKey,
		contractAddresses: ctx.Config().ContractAddresses,
		contractSelectors: selectors,
		epochSettings:     epochSettings,
		systemManager:     systemManager,
	}
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
		commit: submissionABI.Methods["commit"].ID,
		reveal: submissionABI.Methods["reveal"].ID,
		sign:   submissionABI.Methods["sign"].ID,
	}
}

func (c *VotingClient) Run() error {

	for {
		currentTime := time.Now()
		votingEpochStart := c.epochSettings.NextVotingEpochStart(currentTime)
		fmt.Println("Next epoch starts at:", votingEpochStart)
		time.Sleep(time.Until(votingEpochStart))

		currentVotingEpoch := int(c.epochSettings.VotingEpochForTime(time.Now()))

		fmt.Println("Starting voting round: ", currentVotingEpoch)

		processCommits(currentVotingEpoch, c)
		processReveals(currentVotingEpoch-1, c)

		// Wait until reveal deadline
		revealDeadline := votingEpochStart.Add(time.Duration(c.epochSettings.VotingEpochDurationSec / 2))
		time.Sleep(time.Until(revealDeadline))

		// Get results
		processResults(currentVotingEpoch-1, c)
	}
}

// Calldata format:
//
// - 4 bytes: Solidity function selector
// Followed by for each sub-protocol:
// - 1 byte: protocol id
// - 2 bytes: length of data
// - n bytes: data
func processCommits(epoch int, c *VotingClient) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.contractSelectors.commit)

	for _, protocol := range c.subProtocols {
		commitData := getCommitData(epoch, protocol)

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

func processReveals(epoch int, c *VotingClient) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.contractSelectors.reveal)
	for _, protocol := range c.subProtocols {
		revealData := getRevealData(epoch, protocol)

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

func processResults(epoch int, c *VotingClient) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.contractSelectors.sign)
	for _, protocol := range c.subProtocols {
		resultData := getResultsData(epoch, protocol)

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
