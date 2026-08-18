package main

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

type dataProviderResponse struct {
	Status           string `json:"status"`
	Data             string `json:"data"`
	AdditionalData   string `json:"additionalData"`
	FinalizationData string `json:"finalizationData,omitempty"`
}

func NewMockServer(port int, protocolID uint8) *http.Server {
	server := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/submit1/{votingRoundID}/{submitAddress}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		votingRound, err := strconv.Atoi(params["votingRoundID"])
		if err != nil {
			http.Error(w, fmt.Sprintf("writing response: %s", err), http.StatusInternalServerError)
		}
		data := buildMessage(protocolID, uint32(votingRound), []byte("bla"))
		resp := dataProviderResponse{Status: "OK", Data: data}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&resp)
		if err != nil {
			http.Error(w, fmt.Sprintf("writing response: %s", err), http.StatusInternalServerError)
		}
		logger.Infof("handled a submit1 request for voting round %d", votingRound)
	})

	muxRouter.HandleFunc("/submit2/{votingRoundID}/{submitAddress}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		votingRound, err := strconv.Atoi(params["votingRoundID"])
		if err != nil {
			http.Error(w, fmt.Sprintf("writing response: %s", err), http.StatusInternalServerError)
		}
		data := buildMessage(protocolID, uint32(votingRound), []byte("bla"))
		resp := dataProviderResponse{Status: "OK", Data: data}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&resp)
		if err != nil {
			http.Error(w, fmt.Sprintf("writing response: %s", err), http.StatusInternalServerError)
		}
		logger.Infof("handled a submit2 request for voting round %d", votingRound)
	})

	muxRouter.HandleFunc("/submitSignatures/{votingRoundID}/{submitAddress}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		votingRound, err := strconv.Atoi(params["votingRoundID"])
		if err != nil {
			http.Error(w, fmt.Sprintf("writing response: %s", err), http.StatusInternalServerError)
		}
		// a single-leaf tree: the root is the random leaf itself, so the finalization
		// data is the value alone and folds against the message it comes with
		random := crypto.Keccak256Hash([]byte("random"), binary.BigEndian.AppendUint32(nil, uint32(votingRound)))
		merkleRoot := randomLeaf(uint32(votingRound), random)
		data := buildMessageForSigning(protocolID, uint32(votingRound), merkleRoot)
		resp := dataProviderResponse{Status: "OK", Data: data, FinalizationData: random.Hex()}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&resp)
		if err != nil {
			http.Error(w, fmt.Sprintf("writing response: %s", err), http.StatusInternalServerError)
		}
		logger.Infof("handled a submitSignatures request for voting round %d", votingRound)
	})

	server.Handler = muxRouter

	return server
}

func buildMessage(protocolID uint8, votingRoundID uint32, payload []byte) string {
	message := make([]byte, 7, 7+len(payload))
	message[0] = protocolID

	binary.BigEndian.PutUint32(message[1:5], votingRoundID)
	binary.BigEndian.PutUint16(message[5:7], uint16(len(payload)))

	message = append(message, payload...)

	return "0x" + hex.EncodeToString(message)
}

// randomLeaf is the Relay's keccak256(abi.encode(uint256 votingRoundId, uint256 value,
// uint256 isSecure)) for the secure random the message claims.
func randomLeaf(votingRoundID uint32, value common.Hash) []byte {
	var buf [96]byte
	binary.BigEndian.PutUint32(buf[28:32], votingRoundID)
	copy(buf[32:64], value[:])
	buf[95] = 1
	return crypto.Keccak256(buf[:])
}

func buildMessageForSigning(protocolID uint8, roundID uint32, merkleRoot []byte) string {
	data := make([]byte, 38)

	data[0] = protocolID
	binary.BigEndian.PutUint32(data[1:5], roundID)
	data[5] = 1 // claim secure random
	copy(data[6:38], merkleRoot[:])

	return "0x" + hex.EncodeToString(data)
}

func main() {
	logger.Set(logger.Config{Console: true})

	port := 3100
	protocolID := uint8(101)
	srv := NewMockServer(port, protocolID)

	logger.Infof("Server mocking protocol %d running on localhost:%d", protocolID, port)
	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		logger.Error("Server shutdown failed:", err)
	}
}
