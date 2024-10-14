package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"net/http"
	"time"

	"flare-fsc/client/config"

	globalConfig "flare-fsc/config"

	"github.com/gorilla/mux"

	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/logger"
)

type dataProviderResponse struct {
	Status         string `json:"status"`
	Data           string `json:"data"`
	AdditionalData string `json:"additionalData"`
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
			http.Error(w, fmt.Sprintf("error writing response: %s", err), http.StatusInternalServerError)
		}
		data := buildMessage(protocolID, uint32(votingRound), []byte("bla"))
		resp := dataProviderResponse{Status: "OK", Data: data}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&resp)
		if err != nil {
			http.Error(w, fmt.Sprintf("error writing response: %s", err), http.StatusInternalServerError)
		}
		logger.Info("handled a submit1 request for voting round %d", votingRound)
	})

	muxRouter.HandleFunc("/submit2/{votingRoundID}/{submitAddress}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		votingRound, err := strconv.Atoi(params["votingRoundID"])
		if err != nil {
			http.Error(w, fmt.Sprintf("error writing response: %s", err), http.StatusInternalServerError)
		}
		data := buildMessage(protocolID, uint32(votingRound), []byte("bla"))
		resp := dataProviderResponse{Status: "OK", Data: data}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&resp)
		if err != nil {
			http.Error(w, fmt.Sprintf("error writing response: %s", err), http.StatusInternalServerError)
		}
		logger.Info("handled a submit2 request for voting round %d", votingRound)
	})

	muxRouter.HandleFunc("/submitSignatures/{votingRoundID}/{submitAddress}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		votingRound, err := strconv.Atoi(params["votingRoundID"])
		if err != nil {
			http.Error(w, fmt.Sprintf("error writing response: %s", err), http.StatusInternalServerError)
		}
		merkleRoot := bytes.Repeat([]byte{0xff}, 32)
		data := buildMessageForSigning(protocolID, uint32(votingRound), merkleRoot)
		resp := dataProviderResponse{Status: "OK", Data: data}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&resp)
		if err != nil {
			http.Error(w, fmt.Sprintf("error writing response: %s", err), http.StatusInternalServerError)
		}
		logger.Info("handled a submitSignatures request for voting round %d", votingRound)
	})

	server.Handler = muxRouter

	return server
}

func buildMessage(protocolID uint8, votingRoundID uint32, payload []byte) string {
	message := make([]byte, 7)
	message[0] = protocolID

	binary.BigEndian.PutUint32(message[1:5], votingRoundID)
	binary.BigEndian.PutUint16(message[5:7], uint16(len(payload)))

	message = append(message, payload...)

	return "0x" + hex.EncodeToString(message)
}

func buildMessageForSigning(protocolID uint8, roundID uint32, merkleRoot []byte) string {
	data := make([]byte, 38)

	data[0] = uint8(protocolID)
	binary.BigEndian.PutUint32(data[1:5], uint32(roundID))
	data[5] = 1 // claim secure random
	copy(data[6:38], merkleRoot[:])

	return "0x" + hex.EncodeToString(data)

}

func main() {
	globalConfig.GlobalConfigCallback.Call(config.Client{Logger: logger.Config{Console: true}})

	port := 3100
	protocolID := uint8(101)
	srv := NewMockServer(port, protocolID)

	logger.Info("Server mocking protocol 100 running on localhost:%d", port)
	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		logger.Error("Server shutdown failed:", err)
	}
}
