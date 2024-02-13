package registration

import (
	"crypto/ecdsa"
	"flare-tlc/client/shared"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/system"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

var (
	nonFatalSignNewSigningPolicyErrors = []string{
		"new signing policy already signed",
	}
)

type systemsManagerContractClient interface {
	RewardEpochFromChain() (*utils.Epoch, error)
	VotePowerBlockSelectedListener(
		registrationClientDB, *utils.Epoch,
	) <-chan *system.FlareSystemsManagerVotePowerBlockSelected
	SignNewSigningPolicy(*big.Int, []byte) <-chan shared.ExecuteStatus[any]
	GetCurrentRewardEpochId() <-chan shared.ExecuteStatus[*big.Int]
}

type systemsManagerContractClientImpl struct {
	address             common.Address
	flareSystemsManager *system.FlareSystemsManager
	senderTxOpts        *bind.TransactOpts
	txVerifier          *chain.TxVerifier
	signerPrivateKey    *ecdsa.PrivateKey
}

func NewSystemsManagerClient(
	ethClient *ethclient.Client,
	address common.Address,
	senderTxOpts *bind.TransactOpts,
	signerPrivateKey *ecdsa.PrivateKey,
) (*systemsManagerContractClientImpl, error) {
	flareSystemsManager, err := system.NewFlareSystemsManager(address, ethClient)
	if err != nil {
		return nil, err
	}

	return &systemsManagerContractClientImpl{
		address:             address,
		flareSystemsManager: flareSystemsManager,
		senderTxOpts:        senderTxOpts,
		txVerifier:          chain.NewTxVerifier(ethClient),
		signerPrivateKey:    signerPrivateKey,
	}, nil
}

func (s *systemsManagerContractClientImpl) SignNewSigningPolicy(rewardEpochId *big.Int, signingPolicy []byte) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetry(func() (any, error) {
		err := s.sendSignNewSigningPolicy(rewardEpochId, signingPolicy)
		if err != nil {
			return nil, errors.Wrap(err, "error sending sign new signing policy")
		}
		return nil, nil
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)
}

func (s *systemsManagerContractClientImpl) sendSignNewSigningPolicy(rewardEpochId *big.Int, signingPolicy []byte) error {
	newSigningPolicyHash := SigningPolicyHash(signingPolicy)
	hashSignature, err := crypto.Sign(accounts.TextHash(newSigningPolicyHash), s.signerPrivateKey)
	if err != nil {
		return err
	}

	signature := system.IFlareSystemsManagerSignature{
		R: [32]byte(hashSignature[0:32]),
		S: [32]byte(hashSignature[32:64]),
		V: hashSignature[64] + 27,
	}

	tx, err := s.flareSystemsManager.SignNewSigningPolicy(s.senderTxOpts, rewardEpochId, [32]byte(newSigningPolicyHash), signature)
	if err != nil {
		if shared.ExistsAsSubstring(nonFatalSignNewSigningPolicyErrors, err.Error()) {
			logger.Info("Non fatal error sending sign new signing policy: %v", err)
			return nil
		}
		return err
	}
	err = s.txVerifier.WaitUntilMined(s.senderTxOpts.From, tx, chain.DefaultTxTimeout)
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

func (s *systemsManagerContractClientImpl) GetCurrentRewardEpochId() <-chan shared.ExecuteStatus[*big.Int] {
	return shared.ExecuteWithRetry(func() (*big.Int, error) {
		id, err := s.flareSystemsManager.GetCurrentRewardEpochId(nil)
		if err != nil {
			return nil, err
		}
		return id, nil
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)
}

func (s *systemsManagerContractClientImpl) VotePowerBlockSelectedListener(db registrationClientDB, epoch *utils.Epoch) <-chan *system.FlareSystemsManagerVotePowerBlockSelected {
	out := make(chan *system.FlareSystemsManagerVotePowerBlockSelected)
	topic0, err := chain.EventIDFromMetadata(system.FlareSystemsManagerMetaData, "VotePowerBlockSelected")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	go func() {
		ticker := time.NewTicker(shared.ListenerInterval)
		eventRangeStart := epoch.StartTime(epoch.EpochIndex(time.Now()) - 1).Unix()
		for {
			<-ticker.C
			now := time.Now().Unix()
			// logger.Debug("Fetching logs %v < timestamp <= %v", eventRangeStart, now)
			logs, err := db.FetchLogsByAddressAndTopic0(s.address, topic0, eventRangeStart, now)
			if err != nil {
				logger.Error("Error fetching logs %v", err)
				continue
			}
			if len(logs) > 0 {
				// logger.Debug("Found %v logs", len(logs))
				powerBlockData, err := s.parseVotePowerBlockSelectedEvent(logs[len(logs)-1])
				if err != nil {
					logger.Error("Error parsing VotePowerBlockSelected event %v", err)
					continue
				}
				out <- powerBlockData
				// logger.Debug("Sent VotePowerBlockSelected event")
				eventRangeStart = int64(powerBlockData.Timestamp)
			}
		}
	}()
	return out
}

func (s *systemsManagerContractClientImpl) parseVotePowerBlockSelectedEvent(dbLog database.Log) (*system.FlareSystemsManagerVotePowerBlockSelected, error) {
	contractLog, err := shared.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, err
	}
	return s.flareSystemsManager.FlareSystemsManagerFilterer.ParseVotePowerBlockSelected(*contractLog)
}

func (s *systemsManagerContractClientImpl) RewardEpochFromChain() (*utils.Epoch, error) {
	return shared.RewardEpochFromChain(s.flareSystemsManager)
}
