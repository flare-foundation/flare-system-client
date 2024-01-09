package clients

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
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

	for {
		currentTime := time.Now()
		votingEpochStart := c.epochSettings.NextVotingEpochStart(currentTime)
		fmt.Println("Next epoch starts at:", votingEpochStart)
		time.Sleep(time.Until(votingEpochStart))

		currentPriceEpoch := int(c.epochSettings.VotingEpochForTime(currentTime))

		fmt.Println("Processing epoch:", currentPriceEpoch)
		if err := processCommits(c, currentPriceEpoch); err != nil {
			return errors.Wrap(err, "error processing commits")
		}

		if err := processReveals(c, currentPriceEpoch); err != nil {
			return errors.Wrap(err, "error processing reveals")
		}

		// Wait until reveal deadline
		revealDeadline := votingEpochStart.Add(time.Duration(c.epochSettings.VotingEpochDurationSec / 2))
		time.Sleep(time.Until(revealDeadline))

		// Get results
		if err := processResults(c, currentPriceEpoch); err != nil {
			return errors.Wrap(err, "error processing results")
		}
		currentPriceEpoch++
	}
}

// Calldata format:
//
// - 4 bytes: Solidity function selector
// Followed by for each sub-protocol:
// - 1 byte: protocol id
// - 2 bytes: length of data
// - n bytes: data
func processCommits(c *VotingClient, currentPriceEpoch int) error {
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

func processReveals(c *VotingClient, currentPriceEpoch int) error {
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

func processResults(c *VotingClient, currentPriceEpoch int) error {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.contractSelectors.sign)
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
