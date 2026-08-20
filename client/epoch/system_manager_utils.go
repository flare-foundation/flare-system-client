package epoch

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"
	"github.com/flare-foundation/flare-system-client/utils/chain"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/system"
	"github.com/flare-foundation/go-flare-common/pkg/database"
	"github.com/flare-foundation/go-flare-common/pkg/events"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
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

	flareSystemManagerAbi *abi.ABI
)

func init() {
	var err error
	flareSystemManagerAbi, err = system.FlareSystemsManagerMetaData.GetAbi()
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
}

type systemsManagerContractClient interface {
	RewardEpochTimingFromChain() (*utils.EpochTimingConfig, error)

	RewardEpochStartedListener(ctx context.Context, db epochClientDB, config *utils.EpochTimingConfig) <-chan *system.FlareSystemsManagerRewardEpochStarted

	VotePowerBlockSelectedListener(ctx context.Context, db epochClientDB, config *utils.EpochTimingConfig) <-chan *system.FlareSystemsManagerVotePowerBlockSelected
	SignNewSigningPolicy(ctx context.Context, epochID *big.Int, policy []byte) <-chan shared.ExecuteStatus[any]

	SignUptimeVoteEnabledListener(ctx context.Context, db epochClientDB, config *utils.EpochTimingConfig) <-chan *system.FlareSystemsManagerSignUptimeVoteEnabled
	SignUptimeVote(ctx context.Context, epochID *big.Int) <-chan shared.ExecuteStatus[any]

	UptimeVoteSignedListener(ctx context.Context, db epochClientDB, config *utils.EpochTimingConfig) <-chan *system.FlareSystemsManagerUptimeVoteSigned
	SignRewards(ctx context.Context, epochID *big.Int, rewardHash *common.Hash, weightClaims int) <-chan shared.ExecuteStatus[any]
	IsRewardHashSigned(*big.Int) bool

	GetCurrentRewardEpochID() <-chan shared.ExecuteStatus[*big.Int]
}

type systemsManagerContractClientImpl struct {
	address             common.Address
	flareSystemsManager *system.FlareSystemsManager
	senderTxOpts        *bind.TransactOpts
	gasCfg              *config.Gas
	txVerifier          *chain.TxVerifier
	signerPrivateKey    *ecdsa.PrivateKey
	chainID             int64
	ethClient           *ethclient.Client
	relayCutover        *shared.RelayCutover
}

func NewSystemsManagerClient(
	ethClient *ethclient.Client,
	gasCfg *config.Gas,
	address common.Address,
	senderTxOpts *bind.TransactOpts,
	signerPrivateKey *ecdsa.PrivateKey,
	chainID int64,
	relayCutover *shared.RelayCutover) (*systemsManagerContractClientImpl, error) {
	flareSystemsManager, err := system.NewFlareSystemsManager(address, ethClient)
	if err != nil {
		return nil, err
	}

	return &systemsManagerContractClientImpl{
		address:             address,
		flareSystemsManager: flareSystemsManager,
		senderTxOpts:        senderTxOpts,
		gasCfg:              gasCfg,
		txVerifier:          chain.NewTxVerifier(ethClient),
		signerPrivateKey:    signerPrivateKey,
		chainID:             chainID,
		ethClient:           ethClient,
		relayCutover:        relayCutover,
	}, nil
}

func (s *systemsManagerContractClientImpl) SignNewSigningPolicy(ctx context.Context, rewardEpochId *big.Int, signingPolicy []byte) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetryChan(ctx, func() (any, error) {
		err := s.sendSignNewSigningPolicy(ctx, rewardEpochId, signingPolicy)
		if err != nil {
			return nil, fmt.Errorf("sending sign new signing policy: %w", err)
		}
		return nil, nil
	}, shared.MaxTxSendRetriesLong, shared.TxRetryIntervalLong)
}

func (s *systemsManagerContractClientImpl) sendSignNewSigningPolicy(ctx context.Context, rewardEpochId *big.Int, signingPolicy []byte) error {
	newSigningPolicyHash, err := s.signingPolicyHash(rewardEpochId, signingPolicy)
	if err != nil {
		return err
	}

	hashSignature, err := crypto.Sign(accounts.TextHash(newSigningPolicyHash), s.signerPrivateKey)
	if err != nil {
		return err
	}

	signature := system.IFlareSystemsManagerSignature{
		R: [32]byte(hashSignature[0:32]),
		S: [32]byte(hashSignature[32:64]),
		V: hashSignature[64] + 27,
	}

	// senderTxOpts is shared between concurrent send paths, so gas settings are applied to a copy
	txOpts := chain.CopyTxOpts(s.senderTxOpts)

	estimatedGasLimit, err := chain.DryRunTxAbi(
		ctx,
		s.ethClient,
		chain.DefaultTxTimeout,
		txOpts.From,
		s.address,
		common.Big0,
		flareSystemManagerAbi,
		"signNewSigningPolicy",
		rewardEpochId,
		[32]byte(newSigningPolicyHash),
		signature,
	)
	if err != nil {
		if shared.ExistsAsSubstring(nonFatalSignNewSigningPolicyErrors, err.Error()) {
			logger.Debugf("Non fatal error dry run sign new signing policy: %v", err)
			return nil
		}
		logger.Warnf("Dry run fail: %v", err)
		return err
	}
	txOpts.GasLimit = estimatedGasLimit

	err = SetGas(ctx, txOpts, s.ethClient, s.gasCfg)
	if err != nil {
		return err
	}

	tx, err := s.flareSystemsManager.SignNewSigningPolicy(txOpts, rewardEpochId, [32]byte(newSigningPolicyHash), signature)
	if err != nil {
		if shared.ExistsAsSubstring(nonFatalSignNewSigningPolicyErrors, err.Error()) {
			logger.Debugf("Non fatal error sending sign new signing policy: %v", err)
			return nil
		}
		return err
	}
	err = s.txVerifier.WaitUntilMined(ctx, txOpts.From, tx, chain.DefaultTxTimeout)
	if err != nil {
		return err
	}
	logger.Infof("New signing policy sent for epoch %v", rewardEpochId)
	return nil
}

// signingPolicyHash returns the hash of signingPolicy that FlareSystemsManager
// will accept for rewardEpochId: the chain-bound scheme from the breaking epoch on,
// the old fold before it. Which Relay the manager points at is part of the answer,
// not a substitute for the epoch — the new Relay delegates epochs below the breaking
// one to the old Relay, and until governance repoints, the old Relay answers for
// every epoch including the breaking one. The other scheme stays a checked fallback,
// so a table that disagrees with the chain costs a warning, not the signature. Only a
// hash the Relay agrees with is signed, and only one derived from the event's bytes.
func (s *systemsManagerContractClientImpl) signingPolicyHash(rewardEpochId *big.Int, signingPolicy []byte) ([]byte, error) {
	relayAddress, err := s.flareSystemsManager.Relay(nil)
	if err != nil {
		return nil, fmt.Errorf("reading the manager's relay address: %w", err)
	}
	relayContract, err := relay.NewRelay(relayAddress, s.ethClient)
	if err != nil {
		return nil, fmt.Errorf("creating relay contract: %w", err)
	}
	stored, err := relayContract.ToSigningPolicyHash(nil, rewardEpochId)
	if err != nil {
		return nil, fmt.Errorf("reading the signing policy hash of epoch %v: %w", rewardEpochId, err)
	}

	chainBound := s.relayCutover.NewRelayFromRewardEpoch(rewardEpochId.Int64()) &&
		relayAddress == s.relayCutover.NewAddress

	expected, other := SigningPolicyHash(signingPolicy), ChainBoundSigningPolicyHash(signingPolicy, s.chainID)
	expectedName, otherName := "legacy", "chain-bound"
	if chainBound {
		expected, other = other, expected
		expectedName, otherName = otherName, expectedName
	}

	if bytes.Equal(expected, stored[:]) {
		return expected, nil
	}
	if bytes.Equal(other, stored[:]) {
		logger.Warnf(
			"Signing policy of epoch %v is hashed %s by relay %s, not %s as the cutover table implies: the table disagrees with the chain",
			rewardEpochId, otherName, relayAddress, expectedName)
		return other, nil
	}
	return nil, fmt.Errorf("no supported hash of the signing policy of epoch %v matches relay %s hash %s",
		rewardEpochId, relayAddress, common.Hash(stored))
}

// ChainBoundSigningPolicyHash is the hash the new Relay stores: one keccak over
// the 32-byte source chain id followed by the raw encoded policy, unpadded.
func ChainBoundSigningPolicyHash(signingPolicy []byte, chainID int64) []byte {
	return crypto.Keccak256(shared.ChainIDWord(chainID), signingPolicy)
}

// SigningPolicyHash is the hash the old Relay stores: the encoded policy is
// zero-padded to a multiple of 32 bytes and its chunks are folded left to right.
func SigningPolicyHash(signingPolicy []byte) []byte {
	if rest := len(signingPolicy) % 32; rest != 0 {
		// copy — appending could write past the caller's slice into the same allocation
		padded := make([]byte, len(signingPolicy)+32-rest)
		copy(padded, signingPolicy)
		signingPolicy = padded
	}
	hash := crypto.Keccak256(signingPolicy[:32], signingPolicy[32:64])
	for i := 2; i < len(signingPolicy)/32; i++ {
		hash = crypto.Keccak256(hash, signingPolicy[i*32:(i+1)*32])
	}
	return hash
}

func (s *systemsManagerContractClientImpl) GetCurrentRewardEpochID() <-chan shared.ExecuteStatus[*big.Int] {
	return shared.ExecuteWithRetryChan(context.Background(), func() (*big.Int, error) {
		id, err := s.flareSystemsManager.GetCurrentRewardEpochId(nil)
		if err != nil {
			return nil, err
		}
		return id, nil
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)
}

func (s *systemsManagerContractClientImpl) RewardEpochStartedListener(ctx context.Context, db epochClientDB, rewardEpochTiming *utils.EpochTimingConfig) <-chan *system.FlareSystemsManagerRewardEpochStarted {
	out := make(chan *system.FlareSystemsManagerRewardEpochStarted)
	topic0, err := chain.EventIDFromMetadata(system.FlareSystemsManagerMetaData, "RewardEpochStarted")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	go func() {
		randomDelay()
		ticker := time.NewTicker(shared.EventListenerInterval)
		eventRangeStart := rewardEpochTiming.StartTime(rewardEpochTiming.EpochIndex(time.Now())).Unix() - 60*60 // Expected epoch start - 1h
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}
			now := time.Now().Unix()
			logs, err := db.FetchLogsByAddressAndTopic0Timestamp(ctx, s.address, topic0, eventRangeStart, now)
			if err != nil {
				logger.Errorf("Error fetching logs %v", err)
				continue
			}
			if len(logs) > 0 {
				rewardEpochStarted, err := s.parseRewardEpochStartedEvent(logs[len(logs)-1])
				if err != nil {
					logger.Errorf("Error parsing RewardEpochStarted event %v", err)
					continue
				}
				out <- rewardEpochStarted
				eventRangeStart = int64(rewardEpochStarted.Timestamp)
			}
		}
	}()
	return out
}

func (s *systemsManagerContractClientImpl) parseRewardEpochStartedEvent(dbLog database.Log) (*system.FlareSystemsManagerRewardEpochStarted, error) {
	contractLog, err := events.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, err
	}
	return s.flareSystemsManager.ParseRewardEpochStarted(*contractLog)
}

func (s *systemsManagerContractClientImpl) VotePowerBlockSelectedListener(ctx context.Context, db epochClientDB, rewardEpochTiming *utils.EpochTimingConfig) <-chan *system.FlareSystemsManagerVotePowerBlockSelected {
	out := make(chan *system.FlareSystemsManagerVotePowerBlockSelected)
	topic0, err := chain.EventIDFromMetadata(system.FlareSystemsManagerMetaData, "VotePowerBlockSelected")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	go func() {
		randomDelay()
		ticker := time.NewTicker(shared.EventListenerInterval)
		eventRangeStart := rewardEpochTiming.StartTime(rewardEpochTiming.EpochIndex(time.Now()) - 1).Unix()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}
			now := time.Now().Unix()
			logs, err := db.FetchLogsByAddressAndTopic0Timestamp(ctx, s.address, topic0, eventRangeStart, now)
			if err != nil {
				logger.Errorf("Error fetching logs %v", err)
				continue
			}
			if len(logs) > 0 {
				powerBlockData, err := s.parseVotePowerBlockSelectedEvent(logs[len(logs)-1])
				if err != nil {
					logger.Errorf("Error parsing VotePowerBlockSelected event %v", err)
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
	contractLog, err := events.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, err
	}
	return s.flareSystemsManager.ParseVotePowerBlockSelected(*contractLog)
}

func (s *systemsManagerContractClientImpl) RewardEpochTimingFromChain() (*utils.EpochTimingConfig, error) {
	return shared.RewardEpochTimingFromChain(s.flareSystemsManager)
}

func (s *systemsManagerContractClientImpl) SignUptimeVoteEnabledListener(ctx context.Context, db epochClientDB, epoch *utils.EpochTimingConfig) <-chan *system.FlareSystemsManagerSignUptimeVoteEnabled {
	out := make(chan *system.FlareSystemsManagerSignUptimeVoteEnabled)
	topic0, err := chain.EventIDFromMetadata(system.FlareSystemsManagerMetaData, "SignUptimeVoteEnabled")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	go func() {
		randomDelay()
		ticker := time.NewTicker(shared.EventListenerInterval)
		startEpoch := epoch.EpochIndex(time.Now())
		eventRangeStart := epoch.StartTime(startEpoch).Unix()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}
			now := time.Now()
			currentEpoch := epoch.EpochIndex(now)

			logs, err := db.FetchLogsByAddressAndTopic0Timestamp(ctx, s.address, topic0, eventRangeStart, now.Unix())
			if err != nil {
				logger.Errorf("Error fetching logs %v", err)
				continue
			}
			for _, log := range logs {
				uptimeVoteEnabled, err := s.parseSignUptimeVoteEnabledEvent(log)
				if err != nil {
					logger.Errorf("Error parsing SignUptimeVoteEnabled event %v", err)
					continue
				}
				if uptimeVoteEnabled.RewardEpochId.Int64() == currentEpoch-1 {
					out <- uptimeVoteEnabled
				}
				eventRangeStart = int64(uptimeVoteEnabled.Timestamp)
			}
		}
	}()
	return out
}

func (s *systemsManagerContractClientImpl) parseSignUptimeVoteEnabledEvent(dbLog database.Log) (*system.FlareSystemsManagerSignUptimeVoteEnabled, error) {
	contractLog, err := events.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, err
	}
	return s.flareSystemsManager.ParseSignUptimeVoteEnabled(*contractLog)
}

func (s *systemsManagerContractClientImpl) SignUptimeVote(ctx context.Context, rewardEpochId *big.Int) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetryChan(ctx, func() (any, error) {
		err := s.sendSignUptimeVote(ctx, rewardEpochId)
		if err != nil {
			return nil, fmt.Errorf("sending sign uptime vote: %w", err)
		}
		return nil, nil
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)
}

func (s *systemsManagerContractClientImpl) sendSignUptimeVote(ctx context.Context, rewardEpochId *big.Int) error {
	hash, signature, err := getUptimeSignature(rewardEpochId, s.signerPrivateKey)
	if err != nil {
		return err
	}

	// senderTxOpts is shared between concurrent send paths, so gas settings are applied to a copy
	txOpts := chain.CopyTxOpts(s.senderTxOpts)

	estimatedGasLimit, err := chain.DryRunTxAbi(
		ctx,
		s.ethClient,
		chain.DefaultTxTimeout,
		txOpts.From,
		s.address,
		common.Big0,
		flareSystemManagerAbi,
		"signUptimeVote",
		rewardEpochId,
		hash,
		*signature,
	)
	if err != nil {
		if shared.ExistsAsSubstring(nonFatalSignUptimeVoteErrors, err.Error()) {
			logger.Debugf("Non fatal error dryRun sign uptime vote: %v", err)
			return nil
		}
		logger.Warnf("Dry run fail: %v", err)
		return err
	}
	txOpts.GasLimit = estimatedGasLimit

	err = SetGas(ctx, txOpts, s.ethClient, s.gasCfg)
	if err != nil {
		return err
	}

	tx, err := s.flareSystemsManager.SignUptimeVote(txOpts, rewardEpochId, hash, *signature)
	if err != nil {
		if shared.ExistsAsSubstring(nonFatalSignUptimeVoteErrors, err.Error()) {
			logger.Debugf("Non fatal error sending sign uptime vote: %v", err)
			return nil
		}
		return err
	}
	// Uptime signing gates reward signing and is not time-sensitive, so wait long for a single
	// modestly-priced tx to mine instead of timing out and re-sending duplicate transactions.
	err = s.txVerifier.WaitUntilMined(ctx, txOpts.From, tx, chain.LongTxTimeout)
	if err != nil {
		return err
	}
	logger.Infof("Uptime vote sent for epoch %v", rewardEpochId)
	return nil
}

func (s *systemsManagerContractClientImpl) UptimeVoteSignedListener(ctx context.Context, db epochClientDB, epoch *utils.EpochTimingConfig) <-chan *system.FlareSystemsManagerUptimeVoteSigned {
	out := make(chan *system.FlareSystemsManagerUptimeVoteSigned)
	topic0, err := chain.EventIDFromMetadata(system.FlareSystemsManagerMetaData, "UptimeVoteSigned")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	go func() {
		randomDelay()
		ticker := time.NewTicker(shared.EventListenerInterval)
		startEpoch := epoch.EpochIndex(time.Now())
		eventRangeStart := epoch.StartTime(startEpoch).Unix()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}
			now := time.Now()
			currentEpoch := epoch.EpochIndex(now)

			logs, err := db.FetchLogsByAddressAndTopic0Timestamp(ctx, s.address, topic0, eventRangeStart, now.Unix())
			if err != nil {
				logger.Errorf("Error fetching logs %v", err)
				continue
			}

			for _, log := range logs {
				contractLog, err := events.ConvertDatabaseLogToChainLog(log)
				if err != nil {
					logger.Errorf("Error parsing UptimeVoteSigned database log %v", err)
					continue
				}
				uptimeVoteSigned, err := s.flareSystemsManager.ParseUptimeVoteSigned(*contractLog)
				if err != nil {
					logger.Errorf("Error parsing UptimeVoteSigned event %v", err)
					continue
				}
				if uptimeVoteSigned.ThresholdReached && uptimeVoteSigned.RewardEpochId.Int64() == currentEpoch-1 {
					out <- uptimeVoteSigned
				}
				eventRangeStart = int64(uptimeVoteSigned.Timestamp)
			}
		}
	}()
	return out
}

func (s *systemsManagerContractClientImpl) IsRewardHashSigned(epochId *big.Int) bool {
	hash, err := s.flareSystemsManager.RewardsHash(nil, epochId)
	if err != nil {
		logger.Warnf("Error fetching rewards hash for epoch %v: %v", epochId, err)
		return false
	}

	return hash != [32]byte{}
}

func (s *systemsManagerContractClientImpl) SignRewards(ctx context.Context, epochId *big.Int, rewardHash *common.Hash, weightClaims int) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetryChan(ctx, func() (any, error) {
		err := s.sendSignRewards(ctx, epochId, rewardHash, weightClaims)
		if err != nil {
			return nil, fmt.Errorf("sending sign rewards: %w", err)
		}
		return nil, nil
	}, shared.MaxTxSendRetriesLong, shared.TxRetryIntervalLong)
}

func (s *systemsManagerContractClientImpl) sendSignRewards(ctx context.Context, epochId *big.Int, rewardHash *common.Hash, weightClaims int) error {
	logger.Infof("Signing rewards for epoch %v, hash: %s", epochId, rewardHash.Hex())
	packed := encodeRewardsData(epochId, s.chainID, rewardHash, weightClaims)

	hashSignature, err := crypto.Sign(accounts.TextHash(crypto.Keccak256(packed)), s.signerPrivateKey)
	if err != nil {
		return err
	}

	signature := system.IFlareSystemsManagerSignature{
		R: [32]byte(hashSignature[0:32]),
		S: [32]byte(hashSignature[32:64]),
		V: hashSignature[64] + 27,
	}

	numberOfWeightBasedClaims := []system.IFlareSystemsManagerNumberOfWeightBasedClaims{
		{
			RewardManagerId:       big.NewInt(s.chainID),
			NoOfWeightBasedClaims: big.NewInt(int64(weightClaims)),
		},
	}

	// senderTxOpts is shared between concurrent send paths, so gas settings are applied to a copy
	txOpts := chain.CopyTxOpts(s.senderTxOpts)

	estimatedGasLimit, err := chain.DryRunTxAbi(
		ctx,
		s.ethClient,
		chain.DefaultTxTimeout,
		txOpts.From,
		s.address,
		common.Big0,
		flareSystemManagerAbi,
		"signRewards",
		epochId,
		numberOfWeightBasedClaims,
		*rewardHash,
		signature,
	)
	if err != nil {
		if shared.ExistsAsSubstring(nonFatalSignRewardsErrors, err.Error()) {
			logger.Debugf("Non fatal error dry run reward signature: %v", err)
			return nil
		}
		logger.Warnf("Dry run fail: %v", err)
		return err
	}
	txOpts.GasLimit = estimatedGasLimit

	err = SetGas(ctx, txOpts, s.ethClient, s.gasCfg)
	if err != nil {
		return err
	}

	tx, err := s.flareSystemsManager.SignRewards(txOpts, epochId, numberOfWeightBasedClaims, *rewardHash, signature)
	if err != nil {
		if shared.ExistsAsSubstring(nonFatalSignRewardsErrors, err.Error()) {
			logger.Debugf("Non fatal error sending reward signature: %v", err)
			return nil
		}
		return err
	}
	// Reward signing is not time-sensitive, so wait long for a single modestly-priced tx to
	// mine instead of timing out and re-sending duplicate transactions.
	err = s.txVerifier.WaitUntilMined(ctx, txOpts.From, tx, chain.LongTxTimeout)
	if err != nil {
		return err
	}
	logger.Infof("Rewards signed for epoch %v", epochId)

	return nil
}

// sleep for a random duration between 0 and 1 second
func randomDelay() {
	randomDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(randomDuration)
}
