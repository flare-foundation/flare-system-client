package registration

import (
	"crypto/ecdsa"
	"encoding/hex"
	"flare-tlc/client/shared"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/system"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"math/big"
	"math/rand"
	"time"
)

var (
	nonFatalSignNewSigningPolicyErrors = []string{
		"new signing policy already signed",
	}
	nonFatalSignUptimeVoteErrors = []string{
		"submit uptime vote already ended", "voter already signed", "uptime vote hash already signed",
	}
	nonFatalSignRewardsErrors = []string{
		"rewards hash already signed", "voter already signed",
	}
	int64Ty, _          = abi.NewType("int64", "int64", nil)
	bytes32Ty, _        = abi.NewType("bytes32", "bytes32", nil)
	uptimeVoteArguments = abi.Arguments{
		{ // reward epoch id
			Type: int64Ty,
		},
		{ // hash
			Type: bytes32Ty,
		},
	}
)

type systemsManagerContractClient interface {
	RewardEpochFromChain() (*utils.Epoch, error)
	VotePowerBlockSelectedListener(
		registrationClientDB, *utils.Epoch,
	) <-chan *system.FlareSystemsManagerVotePowerBlockSelected
	SignNewSigningPolicy(*big.Int, []byte) <-chan shared.ExecuteStatus[any]
	SignUptimeVoteEnabledListener(registrationClientDB, *utils.Epoch, int64) <-chan *system.FlareSystemsManagerSignUptimeVoteEnabled
	SignUptimeVote(*big.Int) <-chan shared.ExecuteStatus[any]
	UptimeVoteSignedListener(registrationClientDB, *utils.Epoch, int64) <-chan *system.FlareSystemsManagerUptimeVoteSigned
	SignRewards(*big.Int, *common.Hash, int) <-chan shared.ExecuteStatus[any]
	GetCurrentRewardEpochId() <-chan shared.ExecuteStatus[*big.Int]
}

type systemsManagerContractClientImpl struct {
	address             common.Address
	flareSystemsManager *system.FlareSystemsManager
	senderTxOpts        *bind.TransactOpts
	txVerifier          *chain.TxVerifier
	signerPrivateKey    *ecdsa.PrivateKey
	chainId             int
}

func NewSystemsManagerClient(ethClient *ethclient.Client, address common.Address, senderTxOpts *bind.TransactOpts, signerPrivateKey *ecdsa.PrivateKey, chainId int) (*systemsManagerContractClientImpl, error) {
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
		chainId:             chainId,
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
		randomDelay()
		ticker := time.NewTicker(shared.EventListenerInterval)
		eventRangeStart := epoch.StartTime(epoch.EpochIndex(time.Now()) - 1).Unix()
		for {
			<-ticker.C
			now := time.Now().Unix()
			logs, err := db.FetchLogsByAddressAndTopic0(s.address, topic0, eventRangeStart, now)
			if err != nil {
				logger.Error("Error fetching logs %v", err)
				continue
			}
			if len(logs) > 0 {
				powerBlockData, err := s.parseVotePowerBlockSelectedEvent(logs[len(logs)-1])
				if err != nil {
					logger.Error("Error parsing VotePowerBlockSelected event %v", err)
					continue
				}
				out <- powerBlockData
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

func (s *systemsManagerContractClientImpl) SignUptimeVoteEnabledListener(db registrationClientDB, epoch *utils.Epoch, window int64) <-chan *system.FlareSystemsManagerSignUptimeVoteEnabled {
	out := make(chan *system.FlareSystemsManagerSignUptimeVoteEnabled)
	topic0, err := chain.EventIDFromMetadata(system.FlareSystemsManagerMetaData, "SignUptimeVoteEnabled")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	go func() {
		randomDelay()
		ticker := time.NewTicker(shared.EventListenerInterval)
		currentEpoch := epoch.EpochIndex(time.Now())
		eventRangeStart := epoch.StartTime(currentEpoch - window + 1).Unix()
		logger.Info("Current epoch %d", currentEpoch)
		for {
			<-ticker.C
			now := time.Now().Unix()
			logs, err := db.FetchLogsByAddressAndTopic0(s.address, topic0, eventRangeStart, now)
			if err != nil {
				logger.Error("Error fetching logs %v", err)
				continue
			}
			for _, log := range logs {
				uptimeVoteEnabled, err := s.parseSignUptimeVoteEnabledEvent(log)
				if err != nil {
					logger.Error("Error parsing SignUptimeVoteEnabled event %v", err)
					continue
				}
				if uptimeVoteEnabled.RewardEpochId.Int64() >= (currentEpoch - window) {
					out <- uptimeVoteEnabled
				}
				eventRangeStart = int64(uptimeVoteEnabled.Timestamp)
			}
		}
	}()
	return out
}

func (s *systemsManagerContractClientImpl) parseSignUptimeVoteEnabledEvent(dbLog database.Log) (*system.FlareSystemsManagerSignUptimeVoteEnabled, error) {
	contractLog, err := shared.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, err
	}
	return s.flareSystemsManager.FlareSystemsManagerFilterer.ParseSignUptimeVoteEnabled(*contractLog)
}

func (s *systemsManagerContractClientImpl) SignUptimeVote(rewardEpochId *big.Int) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetry(func() (any, error) {
		err := s.sendSignUptimeVote(rewardEpochId)
		if err != nil {
			return nil, errors.Wrap(err, "error sending sign uptime vote")
		}
		return nil, nil
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)
}

func (s *systemsManagerContractClientImpl) sendSignUptimeVote(rewardEpochId *big.Int) error {
	zero := [32]byte{}
	zeroHash := crypto.Keccak256(zero[:])
	logger.Info("Signing uptime vote for epoch %v: %s", rewardEpochId, hex.EncodeToString(zeroHash[:]))

	toSign, err := uptimeVoteArguments.Pack(rewardEpochId.Int64(), [32]byte(zeroHash))
	if err != nil {
		return errors.Wrapf(err, "error packing uptime vote arguments: %v, %v", rewardEpochId, zeroHash)
	}

	hashSignature, err := crypto.Sign(accounts.TextHash(crypto.Keccak256(toSign)), s.signerPrivateKey)
	if err != nil {
		return err
	}

	signature := system.IFlareSystemsManagerSignature{
		R: [32]byte(hashSignature[0:32]),
		S: [32]byte(hashSignature[32:64]),
		V: hashSignature[64] + 27,
	}

	tx, err := s.flareSystemsManager.SignUptimeVote(s.senderTxOpts, rewardEpochId, [32]byte(zeroHash), signature)
	if err != nil {
		if shared.ExistsAsSubstring(nonFatalSignUptimeVoteErrors, err.Error()) {
			logger.Info("Non fatal error sending sign uptime vote: %v", err)
			return nil
		}
		return err
	}
	err = s.txVerifier.WaitUntilMined(s.senderTxOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Info("Uptime vote sent for epoch %v", rewardEpochId)
	return nil
}

func (s *systemsManagerContractClientImpl) UptimeVoteSignedListener(db registrationClientDB, epoch *utils.Epoch, window int64) <-chan *system.FlareSystemsManagerUptimeVoteSigned {
	out := make(chan *system.FlareSystemsManagerUptimeVoteSigned)
	topic0, err := chain.EventIDFromMetadata(system.FlareSystemsManagerMetaData, "UptimeVoteSigned")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	go func() {
		randomDelay()
		ticker := time.NewTicker(shared.EventListenerInterval)
		currentEpoch := epoch.EpochIndex(time.Now())
		eventRangeStart := epoch.StartTime(currentEpoch - window + 1).Unix()
		for {
			<-ticker.C
			now := time.Now().Unix()
			logs, err := db.FetchLogsByAddressAndTopic0(s.address, topic0, eventRangeStart, now)
			if err != nil {
				logger.Error("Error fetching logs %v", err)
				continue
			}

			for _, log := range logs {
				contractLog, err := shared.ConvertDatabaseLogToChainLog(log)
				if err != nil {
					logger.Error("Error parsing UptimeVoteSigned database log %v", err)
					continue
				}
				uptimeVoteSigned, err := s.flareSystemsManager.FlareSystemsManagerFilterer.ParseUptimeVoteSigned(*contractLog)
				if err != nil {
					logger.Error("Error parsing UptimeVoteSigned event %v", err)
					continue
				}
				if uptimeVoteSigned.ThresholdReached && uptimeVoteSigned.RewardEpochId.Int64() >= (currentEpoch-window) {
					out <- uptimeVoteSigned
				}
				eventRangeStart = int64(uptimeVoteSigned.Timestamp)
			}
		}
	}()
	return out
}

func (s *systemsManagerContractClientImpl) SignRewards(epochId *big.Int, rewardHash *common.Hash, weightClaims int) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetry(func() (any, error) {
		err := s.sendSignRewards(epochId, rewardHash, weightClaims)
		if err != nil {
			return nil, errors.Wrap(err, "error sending sign rewards")
		}
		return nil, nil
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)
}

func (s *systemsManagerContractClientImpl) sendSignRewards(epochId *big.Int, rewardHash *common.Hash, weightClaims int) error {
	logger.Info("Signing rewards for epoch %v, hash: %s", epochId, rewardHash.Hex())
	packed := encodeRewardsData(epochId, s.chainId, rewardHash, weightClaims)

	hashSignature, err := crypto.Sign(accounts.TextHash(crypto.Keccak256(packed)), s.signerPrivateKey)
	if err != nil {
		return err
	}

	signature := system.IFlareSystemsManagerSignature{
		R: [32]byte(hashSignature[0:32]),
		S: [32]byte(hashSignature[32:64]),
		V: hashSignature[64] + 27,
	}

	tx, err := s.flareSystemsManager.SignRewards(s.senderTxOpts, epochId, []system.IFlareSystemsManagerNumberOfWeightBasedClaims{
		{
			RewardManagerId:       big.NewInt(int64(s.chainId)),
			NoOfWeightBasedClaims: big.NewInt(int64(weightClaims)),
		},
	}, *rewardHash, signature)
	if err != nil {
		if shared.ExistsAsSubstring(nonFatalSignRewardsErrors, err.Error()) {
			logger.Info("Non fatal error sending reward signature: %v", err)
			return nil
		}
		return err
	}
	err = s.txVerifier.WaitUntilMined(s.senderTxOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Info("Rewards signed for epoch %v", epochId)

	return nil
}

// sleep for a random duration between 0 and 1 second
func randomDelay() {
	randomDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(randomDuration)
}
