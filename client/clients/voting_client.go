package clients

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	localContext "flare-tlc/client/context"
	"flare-tlc/config"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/submission"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type VotingClient struct {
	subProtocols      []subProtocol
	eth               *ethclient.Client
	privateKey        string
	contractAddresses config.ContractAddresses
	selectors         selectors
}

type subProtocol struct {
	// TODO: Should use 2 bytes for id?
	id byte
	// E.g. localhost:3000/ftso/price-controller/
	apiEndpoint string
}

var currentPriceEpoch = 1452560

func NewVotingClient(ctx localContext.ClientContext) *VotingClient {
	// TODO: Read from config
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
		panic(err)
	}

	bn, err := cl.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Current block number: ", bn)

	privateKey, err := ctx.Config().ChainConfig().GetPrivateKey()
	if err != nil {
		panic(err)
	}

	return &VotingClient{
		eth:               cl,
		subProtocols:      subProtocols,
		privateKey:        privateKey,
		contractAddresses: ctx.Config().ContractAddresses,
		selectors:         newSelectors(),
	}
}

type selectors struct {
	commitSelector []byte
	revealSelector []byte
	resultSelector []byte
}

func newSelectors() selectors {
	submissionABI, err := abi.JSON(strings.NewReader(submission.SubmissionABI))
	if err != nil {
		panic(err)
	}
	return selectors{
		commitSelector: submissionABI.Methods["commit"].ID,
		revealSelector: submissionABI.Methods["reveal"].ID,
		resultSelector: submissionABI.Methods["result"].ID,
	}
}

func (c *VotingClient) Run() error {
	ticker := time.NewTicker(30 * time.Second)

	// TODO: Make sure to start at the start of a voting round.
	// 		 Requires getting epoch settings from smart contracts.
	for startTime := time.Now(); true; startTime = <-ticker.C {
		fmt.Println("Starting voting round: ", currentPriceEpoch)

		processCommits(c)
		processReveals(c)

		// Wait until reveal deadline
		revealDeadline := startTime.Add(15 * time.Second)
		time.Sleep(time.Until(revealDeadline))

		// Get results
		processResults(c)

		currentPriceEpoch++
	}

	return nil
}

func processCommits(c *VotingClient) {
	buffer := bytes.NewBuffer(nil)

	buffer.Write(c.selectors.commitSelector)

	for _, protocol := range c.subProtocols {
		commitData := getCommitData(currentPriceEpoch, protocol)

		buffer.WriteByte(protocol.id)
		// TODO: Probablty don't need 4 bytes for length
		length := intToBytes(len(commitData))
		buffer.Write(length[:])
		buffer.Write(commitData)

		fmt.Println("Total encoded:", buffer.Len())
	}
	commitPayload := buffer.Bytes()
	fmt.Println("Submitting commit payload:", len(commitPayload))

	sendRawTx(*c.eth, c.privateKey, c.contractAddresses.Submission, commitPayload)
}

func processReveals(c *VotingClient) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.selectors.revealSelector)
	for _, protocol := range c.subProtocols {
		revealData := getRevealData(currentPriceEpoch, protocol)

		buffer.WriteByte(protocol.id)
		length := intToBytes(len(revealData))
		buffer.Write(length[:])
		buffer.Write(revealData)

		fmt.Println("Total encoded:", buffer.Len())
	}
	revealPayload := buffer.Bytes()
	fmt.Println("Submitting reveal payload:", len(revealPayload))
	// TODO: Submit to smart contract
}

func processResults(c *VotingClient) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(c.selectors.resultSelector)
	for _, protocol := range c.subProtocols {
		revealData := getResultsData(currentPriceEpoch, protocol)

		buffer.WriteByte(protocol.id)
		length := intToBytes(len(revealData))
		buffer.Write(length[:])
		buffer.Write(revealData)

		fmt.Println("Total encoded:", buffer.Len())
	}
	revealPayload := buffer.Bytes()
	fmt.Println("Submitting result payload:", len(revealPayload))
	// TODO: Submit to smart contract
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

	fmt.Println("Commit data: ", string(body))
	return body
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
	fmt.Println("Reveal data: ", string(body))

	return body
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

	fmt.Println("Result data: ", string(body))
	return body
}

func intToBytes(i int) (arr [4]byte) {
	binary.BigEndian.PutUint32(arr[0:4], uint32(i))
	return
}

func sendRawTx(client ethclient.Client, pk string, toAddress common.Address, data []byte) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(210000)               // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sending signed tx: ", signedTx.Hash().Hex())

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err)
	}

	rec, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		panic(err)
	}
	fmt.Println("Receipt: ", rec.Status)

	verifier := chain.NewTxVerifier(&client)

	fmt.Println("Waiting for tx to be mined...", time.Now())
	err = verifier.WaitUntilMined(fromAddress, signedTx, 10*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tx mined ", time.Now())
}
