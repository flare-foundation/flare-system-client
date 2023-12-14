package clients

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"time"
)

type VotingClient struct {
	subProtocols []subProtocol
}

type subProtocol struct {
	// TODO: Should use 2 bytes for id?
	id byte
	// E.g. localhost:3000/ftso/price-controller/
	apiEndpoint string
}

var currentPriceEpoch = 1452560

func NewVotingClient() *VotingClient {
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

	return &VotingClient{
		subProtocols: subProtocols,
	}
}

func (c *VotingClient) Run() error {
	ticker := time.NewTicker(30 * time.Second)

	// TODO: Make sure to start at the start of a voting round.
	// 		 Requires getting voting settings from smart contract.
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
	// TODO: Submit to smart contract
}

func processReveals(c *VotingClient) {
	buffer := bytes.NewBuffer(nil)
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
