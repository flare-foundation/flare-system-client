package clients

import (
	"encoding/hex"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/relay"
	"flare-tlc/utils/contracts/system"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"gorm.io/gorm"
)

type signingPolicyHandler struct {
	address string
	topic0  string

	relay              *relay.Relay
	flareSystemManager *system.FlareSystemManager

	txVerifier *chain.TxVerifier
	txOpts     *bind.TransactOpts

	db           *gorm.DB
	mockableTime utils.TimeProvider
}

func NewSigningPolicyHandler(
	db *gorm.DB,
	relay *relay.Relay,
	flareSystemManager *system.FlareSystemManager,
	txVerifier *chain.TxVerifier,
	txOpts *bind.TransactOpts,
	address string,
	topic0 string,
) *signingPolicyHandler {
	return &signingPolicyHandler{
		address:            address,
		topic0:             topic0,
		relay:              relay,
		flareSystemManager: flareSystemManager,
		txOpts:             txOpts,
		db:                 db,
		mockableTime:       &utils.RealTimeProvider{},
	}
}

func (s *signingPolicyHandler) signingPolicyInitializedListener(startTimestamp uint64) <-chan *relay.RelaySigningPolicyInitialized {
	out := make(chan *relay.RelaySigningPolicyInitialized)
	go func() {
		ticker := time.NewTicker(10 * time.Second) // read from config
		for {
			<-ticker.C
			now := s.mockableTime.Now().Unix()
			logs, err := database.FetchLogsByAddressAndTopic0(s.db, s.address, s.topic0, int64(startTimestamp), now)
			if err != nil {
				logger.Error("Error fetching logs %w", err)
				continue
			}
			if len(logs) > 0 {
				policyData, err := s.parseSigningPolicyInitializedEvent(logs[len(logs)-1])
				if err != nil {
					logger.Error("Error parsing SigningPolicyInitialized event %w", err)
					continue
				}
				out <- policyData
			}
		}
	}()
	return out
}

func (s *signingPolicyHandler) parseSigningPolicyInitializedEvent(dbLog database.Log) (*relay.RelaySigningPolicyInitialized, error) {
	data, err := hex.DecodeString(dbLog.Data)
	if err != nil {
		return nil, err
	}
	contractLog := types.Log{
		Topics: []common.Hash{
			chain.ParseTopic(dbLog.Topic0),
			chain.ParseTopic(dbLog.Topic1),
			chain.ParseTopic(dbLog.Topic2),
			chain.ParseTopic(dbLog.Topic3),
		},
		Data: data,
		// Other fields are not used by log decoder
	}

	return s.relay.RelayFilterer.ParseSigningPolicyInitialized(contractLog)
}

func (s *signingPolicyHandler) SignNewSigningPolicy(rewardEpochId *big.Int, signingPolicy []byte) <-chan bool {
	out := make(chan bool)
	go func() {
		for retry := 0; retry < maxRetries; retry++ {
			err := s.SendSignNewSigningPolicy(rewardEpochId, signingPolicy)
			if err != nil {
				logger.Error("SendSignNewSigningPolicy: %w", err)
			} else {
				out <- true
				return
			}
		}
		logger.Error("SendSignNewSigningPolicy: max retries reached")
		out <- false
	}()
	return out
}

func (s *signingPolicyHandler) SendSignNewSigningPolicy(rewardEpochId *big.Int, signingPolicy []byte) error {
	newSigningPolicyHash := SigningPolicyHash(signingPolicy)
	hashSignature := crypto.Keccak256(newSigningPolicyHash)
	signature := system.FlareSystemManagerSignature{
		V: hashSignature[0],
		R: [32]byte(hashSignature[1:33]),
		S: [32]byte(hashSignature[33:65]),
	}

	tx, err := s.flareSystemManager.SignNewSigningPolicy(s.txOpts, rewardEpochId, [32]byte(newSigningPolicyHash), signature)
	if err != nil {
		return err
	}
	err = s.txVerifier.WaitUntilMined(s.txOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Info("New signing policy sent for epoch %v", rewardEpochId)
	return nil
}

func SigningPolicyHash(signingPolicy []byte) []byte {
	if len(signingPolicy)%32 != 0 {
		signingPolicy = append(signingPolicy, make([]byte, 32-len(signingPolicy)%32)...)
	}
	hash := crypto.Keccak256(signingPolicy[:32], signingPolicy[32:64])
	for i := 2; i < len(signingPolicy)/32; i++ {
		hash = crypto.Keccak256(hash, signingPolicy[i*32:(i+1)*32])
	}
	return hash
}
