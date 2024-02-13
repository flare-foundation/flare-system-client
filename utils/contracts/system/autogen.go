// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package system

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FlareSystemsManagerInitialSettings is an auto generated low-level Go binding around an user-defined struct.
type FlareSystemsManagerInitialSettings struct {
	InitialRandomVotePowerBlockSelectionSize uint16
	InitialRewardEpochId                     *big.Int
	InitialRewardEpochThreshold              uint16
}

// FlareSystemsManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type FlareSystemsManagerSettings struct {
	RandomAcquisitionMaxDurationSeconds          uint16
	RandomAcquisitionMaxDurationBlocks           uint16
	NewSigningPolicyInitializationStartSeconds   uint16
	NewSigningPolicyMinNumberOfVotingRoundsDelay uint8
	VoterRegistrationMinDurationSeconds          uint16
	VoterRegistrationMinDurationBlocks           uint16
	SubmitUptimeVoteMinDurationSeconds           uint16
	SubmitUptimeVoteMinDurationBlocks            uint16
	SigningPolicyThresholdPPM                    *big.Int
	SigningPolicyMinNumberOfVoters               uint16
	RewardExpiryOffsetSeconds                    uint32
}

// IFlareSystemsManagerSignature is an auto generated low-level Go binding around an user-defined struct.
type IFlareSystemsManagerSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// FlareSystemsManagerMetaData contains all meta data concerning the FlareSystemsManager contract.
var FlareSystemsManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_flareDaemon\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"randomAcquisitionMaxDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"randomAcquisitionMaxDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newSigningPolicyInitializationStartSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"newSigningPolicyMinNumberOfVotingRoundsDelay\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"voterRegistrationMinDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"voterRegistrationMinDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"submitUptimeVoteMinDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"submitUptimeVoteMinDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"signingPolicyThresholdPPM\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"signingPolicyMinNumberOfVoters\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"rewardExpiryOffsetSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structFlareSystemsManager.Settings\",\"name\":\"_settings\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"_firstVotingRoundStartTs\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_votingEpochDurationSeconds\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_firstRewardEpochStartVotingRoundId\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_rewardEpochDurationInVotingEpochs\",\"type\":\"uint16\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"initialRandomVotePowerBlockSelectionSize\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"initialRewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"initialRewardEpochThreshold\",\"type\":\"uint16\"}],\"internalType\":\"structFlareSystemsManager.InitialSettings\",\"name\":\"_initialSettings\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"ClosingExpiredRewardEpochFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"RandomAcquisitionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startVotingRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"RewardEpochStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rewardsHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noOfWeightBasedClaims\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"RewardsSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"name\":\"SettingCleanUpBlockNumberFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"SignUptimeVoteEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"SigningPolicySigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"TriggeringVoterRegistrationFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"uptimeVoteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"UptimeVoteSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes20[]\",\"name\":\"nodeIds\",\"type\":\"bytes20[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"UptimeVoteSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"votePowerBlock\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"VotePowerBlockSelected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_signingPolicyThresholdPPM\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"_signingPolicyMinNumberOfVoters\",\"type\":\"uint16\"}],\"name\":\"changeSigningPolicySettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanupBlockNumberManager\",\"outputs\":[{\"internalType\":\"contractIICleanupBlockNumberManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewardEpochExpectedEndTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daemonize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstRewardEpochStartTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstVotingRoundStartTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareDaemon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRewardEpochId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getRandomAcquisitionInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_randomAcquisitionStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_randomAcquisitionStartBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_randomAcquisitionEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_randomAcquisitionEndBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getRewardEpochStartInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_rewardEpochStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardEpochStartBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardEpochSwitchoverTriggerContracts\",\"outputs\":[{\"internalType\":\"contractIIRewardEpochSwitchoverTrigger[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getRewardsSignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_rewardsSignStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardsSignStartBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardsSignEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardsSignEndBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getSeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getSigningPolicySignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignStartBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignEndBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getStartVotingRoundId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getUptimeVoteSignStartInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSignStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSignStartBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getVotePowerBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_votePowerBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getVoterRegistrationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_votePowerBlock\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterRewardsSignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_rewardsSignTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardsSignBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterSigningPolicySignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterUptimeVoteSignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSignTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSignBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterUptimeVoteSubmitInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSubmitTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSubmitBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialRandomVotePowerBlockSelectionSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isVoterRegistrationEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInitializedVotingRoundId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newSigningPolicyInitializationStartSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newSigningPolicyMinNumberOfVotingRoundsDelay\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"noOfWeightBasedClaims\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomAcquisitionMaxDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomAcquisitionMaxDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relay\",\"outputs\":[{\"internalType\":\"contractIIRelay\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardEpochIdToExpireNext\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardExpiryOffsetSeconds\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"rewardsHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIIRewardEpochSwitchoverTrigger[]\",\"name\":\"_contracts\",\"type\":\"address[]\"}],\"name\":\"setRewardEpochSwitchoverTriggerContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_noOfWeightBasedClaims\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_rewardsHash\",\"type\":\"bytes32\"}],\"name\":\"setRewardsData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_submit3Aligned\",\"type\":\"bool\"}],\"name\":\"setSubmit3Aligned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_triggerExpirationAndCleanup\",\"type\":\"bool\"}],\"name\":\"setTriggerExpirationAndCleanup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIIVoterRegistrationTrigger\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setVoterRegistrationTriggerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"_newSigningPolicyHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemsManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signNewSigningPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_noOfWeightBasedClaims\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_rewardsHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemsManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"_uptimeVoteHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemsManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signUptimeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingPolicyMinNumberOfVoters\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingPolicyThresholdPPM\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submission\",\"outputs\":[{\"internalType\":\"contractIISubmission\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submit3Aligned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes20[]\",\"name\":\"_nodeIds\",\"type\":\"bytes20[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemsManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"submitUptimeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitUptimeVoteMinDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitUptimeVoteMinDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToFallbackMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerExpirationAndCleanup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"randomAcquisitionMaxDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"randomAcquisitionMaxDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newSigningPolicyInitializationStartSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"newSigningPolicyMinNumberOfVotingRoundsDelay\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"voterRegistrationMinDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"voterRegistrationMinDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"submitUptimeVoteMinDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"submitUptimeVoteMinDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"signingPolicyThresholdPPM\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"signingPolicyMinNumberOfVoters\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"rewardExpiryOffsetSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structFlareSystemsManager.Settings\",\"name\":\"_settings\",\"type\":\"tuple\"}],\"name\":\"updateSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"uptimeVoteHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistrationMinDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistrationMinDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistrationTriggerContract\",\"outputs\":[{\"internalType\":\"contractIIVoterRegistrationTrigger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistry\",\"outputs\":[{\"internalType\":\"contractIIVoterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FlareSystemsManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use FlareSystemsManagerMetaData.ABI instead.
var FlareSystemsManagerABI = FlareSystemsManagerMetaData.ABI

// FlareSystemsManager is an auto generated Go binding around an Ethereum contract.
type FlareSystemsManager struct {
	FlareSystemsManagerCaller     // Read-only binding to the contract
	FlareSystemsManagerTransactor // Write-only binding to the contract
	FlareSystemsManagerFilterer   // Log filterer for contract events
}

// FlareSystemsManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlareSystemsManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlareSystemsManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlareSystemsManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlareSystemsManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlareSystemsManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlareSystemsManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlareSystemsManagerSession struct {
	Contract     *FlareSystemsManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FlareSystemsManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlareSystemsManagerCallerSession struct {
	Contract *FlareSystemsManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// FlareSystemsManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlareSystemsManagerTransactorSession struct {
	Contract     *FlareSystemsManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// FlareSystemsManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlareSystemsManagerRaw struct {
	Contract *FlareSystemsManager // Generic contract binding to access the raw methods on
}

// FlareSystemsManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlareSystemsManagerCallerRaw struct {
	Contract *FlareSystemsManagerCaller // Generic read-only contract binding to access the raw methods on
}

// FlareSystemsManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlareSystemsManagerTransactorRaw struct {
	Contract *FlareSystemsManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlareSystemsManager creates a new instance of FlareSystemsManager, bound to a specific deployed contract.
func NewFlareSystemsManager(address common.Address, backend bind.ContractBackend) (*FlareSystemsManager, error) {
	contract, err := bindFlareSystemsManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManager{FlareSystemsManagerCaller: FlareSystemsManagerCaller{contract: contract}, FlareSystemsManagerTransactor: FlareSystemsManagerTransactor{contract: contract}, FlareSystemsManagerFilterer: FlareSystemsManagerFilterer{contract: contract}}, nil
}

// NewFlareSystemsManagerCaller creates a new read-only instance of FlareSystemsManager, bound to a specific deployed contract.
func NewFlareSystemsManagerCaller(address common.Address, caller bind.ContractCaller) (*FlareSystemsManagerCaller, error) {
	contract, err := bindFlareSystemsManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerCaller{contract: contract}, nil
}

// NewFlareSystemsManagerTransactor creates a new write-only instance of FlareSystemsManager, bound to a specific deployed contract.
func NewFlareSystemsManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FlareSystemsManagerTransactor, error) {
	contract, err := bindFlareSystemsManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerTransactor{contract: contract}, nil
}

// NewFlareSystemsManagerFilterer creates a new log filterer instance of FlareSystemsManager, bound to a specific deployed contract.
func NewFlareSystemsManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FlareSystemsManagerFilterer, error) {
	contract, err := bindFlareSystemsManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerFilterer{contract: contract}, nil
}

// bindFlareSystemsManager binds a generic wrapper to an already deployed contract.
func bindFlareSystemsManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FlareSystemsManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlareSystemsManager *FlareSystemsManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlareSystemsManager.Contract.FlareSystemsManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlareSystemsManager *FlareSystemsManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.FlareSystemsManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlareSystemsManager *FlareSystemsManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.FlareSystemsManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlareSystemsManager *FlareSystemsManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlareSystemsManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlareSystemsManager *FlareSystemsManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlareSystemsManager *FlareSystemsManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.contract.Transact(opts, method, params...)
}

// CleanupBlockNumberManager is a free data retrieval call binding the contract method 0x4eac870f.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCaller) CleanupBlockNumberManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "cleanupBlockNumberManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CleanupBlockNumberManager is a free data retrieval call binding the contract method 0x4eac870f.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerSession) CleanupBlockNumberManager() (common.Address, error) {
	return _FlareSystemsManager.Contract.CleanupBlockNumberManager(&_FlareSystemsManager.CallOpts)
}

// CleanupBlockNumberManager is a free data retrieval call binding the contract method 0x4eac870f.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) CleanupBlockNumberManager() (common.Address, error) {
	return _FlareSystemsManager.Contract.CleanupBlockNumberManager(&_FlareSystemsManager.CallOpts)
}

// CurrentRewardEpochExpectedEndTs is a free data retrieval call binding the contract method 0xed54fd63.
//
// Solidity: function currentRewardEpochExpectedEndTs() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCaller) CurrentRewardEpochExpectedEndTs(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "currentRewardEpochExpectedEndTs")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrentRewardEpochExpectedEndTs is a free data retrieval call binding the contract method 0xed54fd63.
//
// Solidity: function currentRewardEpochExpectedEndTs() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerSession) CurrentRewardEpochExpectedEndTs() (uint64, error) {
	return _FlareSystemsManager.Contract.CurrentRewardEpochExpectedEndTs(&_FlareSystemsManager.CallOpts)
}

// CurrentRewardEpochExpectedEndTs is a free data retrieval call binding the contract method 0xed54fd63.
//
// Solidity: function currentRewardEpochExpectedEndTs() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) CurrentRewardEpochExpectedEndTs() (uint64, error) {
	return _FlareSystemsManager.Contract.CurrentRewardEpochExpectedEndTs(&_FlareSystemsManager.CallOpts)
}

// FirstRewardEpochStartTs is a free data retrieval call binding the contract method 0x79e047ed.
//
// Solidity: function firstRewardEpochStartTs() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCaller) FirstRewardEpochStartTs(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "firstRewardEpochStartTs")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstRewardEpochStartTs is a free data retrieval call binding the contract method 0x79e047ed.
//
// Solidity: function firstRewardEpochStartTs() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerSession) FirstRewardEpochStartTs() (uint64, error) {
	return _FlareSystemsManager.Contract.FirstRewardEpochStartTs(&_FlareSystemsManager.CallOpts)
}

// FirstRewardEpochStartTs is a free data retrieval call binding the contract method 0x79e047ed.
//
// Solidity: function firstRewardEpochStartTs() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) FirstRewardEpochStartTs() (uint64, error) {
	return _FlareSystemsManager.Contract.FirstRewardEpochStartTs(&_FlareSystemsManager.CallOpts)
}

// FirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCaller) FirstVotingRoundStartTs(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "firstVotingRoundStartTs")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerSession) FirstVotingRoundStartTs() (uint64, error) {
	return _FlareSystemsManager.Contract.FirstVotingRoundStartTs(&_FlareSystemsManager.CallOpts)
}

// FirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) FirstVotingRoundStartTs() (uint64, error) {
	return _FlareSystemsManager.Contract.FirstVotingRoundStartTs(&_FlareSystemsManager.CallOpts)
}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCaller) FlareDaemon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "flareDaemon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerSession) FlareDaemon() (common.Address, error) {
	return _FlareSystemsManager.Contract.FlareDaemon(&_FlareSystemsManager.CallOpts)
}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) FlareDaemon() (common.Address, error) {
	return _FlareSystemsManager.Contract.FlareDaemon(&_FlareSystemsManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetAddressUpdater() (common.Address, error) {
	return _FlareSystemsManager.Contract.GetAddressUpdater(&_FlareSystemsManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetAddressUpdater() (common.Address, error) {
	return _FlareSystemsManager.Contract.GetAddressUpdater(&_FlareSystemsManager.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetContractName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getContractName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetContractName() (string, error) {
	return _FlareSystemsManager.Contract.GetContractName(&_FlareSystemsManager.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetContractName() (string, error) {
	return _FlareSystemsManager.Contract.GetContractName(&_FlareSystemsManager.CallOpts)
}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetCurrentRewardEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getCurrentRewardEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetCurrentRewardEpochId() (*big.Int, error) {
	return _FlareSystemsManager.Contract.GetCurrentRewardEpochId(&_FlareSystemsManager.CallOpts)
}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetCurrentRewardEpochId() (*big.Int, error) {
	return _FlareSystemsManager.Contract.GetCurrentRewardEpochId(&_FlareSystemsManager.CallOpts)
}

// GetRandomAcquisitionInfo is a free data retrieval call binding the contract method 0x8f8f9f3a.
//
// Solidity: function getRandomAcquisitionInfo(uint24 _rewardEpochId) view returns(uint64 _randomAcquisitionStartTs, uint64 _randomAcquisitionStartBlock, uint64 _randomAcquisitionEndTs, uint64 _randomAcquisitionEndBlock)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetRandomAcquisitionInfo(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	RandomAcquisitionStartTs    uint64
	RandomAcquisitionStartBlock uint64
	RandomAcquisitionEndTs      uint64
	RandomAcquisitionEndBlock   uint64
}, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getRandomAcquisitionInfo", _rewardEpochId)

	outstruct := new(struct {
		RandomAcquisitionStartTs    uint64
		RandomAcquisitionStartBlock uint64
		RandomAcquisitionEndTs      uint64
		RandomAcquisitionEndBlock   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RandomAcquisitionStartTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.RandomAcquisitionStartBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.RandomAcquisitionEndTs = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.RandomAcquisitionEndBlock = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetRandomAcquisitionInfo is a free data retrieval call binding the contract method 0x8f8f9f3a.
//
// Solidity: function getRandomAcquisitionInfo(uint24 _rewardEpochId) view returns(uint64 _randomAcquisitionStartTs, uint64 _randomAcquisitionStartBlock, uint64 _randomAcquisitionEndTs, uint64 _randomAcquisitionEndBlock)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetRandomAcquisitionInfo(_rewardEpochId *big.Int) (struct {
	RandomAcquisitionStartTs    uint64
	RandomAcquisitionStartBlock uint64
	RandomAcquisitionEndTs      uint64
	RandomAcquisitionEndBlock   uint64
}, error) {
	return _FlareSystemsManager.Contract.GetRandomAcquisitionInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetRandomAcquisitionInfo is a free data retrieval call binding the contract method 0x8f8f9f3a.
//
// Solidity: function getRandomAcquisitionInfo(uint24 _rewardEpochId) view returns(uint64 _randomAcquisitionStartTs, uint64 _randomAcquisitionStartBlock, uint64 _randomAcquisitionEndTs, uint64 _randomAcquisitionEndBlock)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetRandomAcquisitionInfo(_rewardEpochId *big.Int) (struct {
	RandomAcquisitionStartTs    uint64
	RandomAcquisitionStartBlock uint64
	RandomAcquisitionEndTs      uint64
	RandomAcquisitionEndBlock   uint64
}, error) {
	return _FlareSystemsManager.Contract.GetRandomAcquisitionInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetRewardEpochStartInfo is a free data retrieval call binding the contract method 0x00ddae53.
//
// Solidity: function getRewardEpochStartInfo(uint24 _rewardEpochId) view returns(uint64 _rewardEpochStartTs, uint64 _rewardEpochStartBlock)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetRewardEpochStartInfo(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	RewardEpochStartTs    uint64
	RewardEpochStartBlock uint64
}, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getRewardEpochStartInfo", _rewardEpochId)

	outstruct := new(struct {
		RewardEpochStartTs    uint64
		RewardEpochStartBlock uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardEpochStartTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.RewardEpochStartBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetRewardEpochStartInfo is a free data retrieval call binding the contract method 0x00ddae53.
//
// Solidity: function getRewardEpochStartInfo(uint24 _rewardEpochId) view returns(uint64 _rewardEpochStartTs, uint64 _rewardEpochStartBlock)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetRewardEpochStartInfo(_rewardEpochId *big.Int) (struct {
	RewardEpochStartTs    uint64
	RewardEpochStartBlock uint64
}, error) {
	return _FlareSystemsManager.Contract.GetRewardEpochStartInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetRewardEpochStartInfo is a free data retrieval call binding the contract method 0x00ddae53.
//
// Solidity: function getRewardEpochStartInfo(uint24 _rewardEpochId) view returns(uint64 _rewardEpochStartTs, uint64 _rewardEpochStartBlock)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetRewardEpochStartInfo(_rewardEpochId *big.Int) (struct {
	RewardEpochStartTs    uint64
	RewardEpochStartBlock uint64
}, error) {
	return _FlareSystemsManager.Contract.GetRewardEpochStartInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetRewardEpochSwitchoverTriggerContracts is a free data retrieval call binding the contract method 0x46831531.
//
// Solidity: function getRewardEpochSwitchoverTriggerContracts() view returns(address[])
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetRewardEpochSwitchoverTriggerContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getRewardEpochSwitchoverTriggerContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRewardEpochSwitchoverTriggerContracts is a free data retrieval call binding the contract method 0x46831531.
//
// Solidity: function getRewardEpochSwitchoverTriggerContracts() view returns(address[])
func (_FlareSystemsManager *FlareSystemsManagerSession) GetRewardEpochSwitchoverTriggerContracts() ([]common.Address, error) {
	return _FlareSystemsManager.Contract.GetRewardEpochSwitchoverTriggerContracts(&_FlareSystemsManager.CallOpts)
}

// GetRewardEpochSwitchoverTriggerContracts is a free data retrieval call binding the contract method 0x46831531.
//
// Solidity: function getRewardEpochSwitchoverTriggerContracts() view returns(address[])
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetRewardEpochSwitchoverTriggerContracts() ([]common.Address, error) {
	return _FlareSystemsManager.Contract.GetRewardEpochSwitchoverTriggerContracts(&_FlareSystemsManager.CallOpts)
}

// GetRewardsSignInfo is a free data retrieval call binding the contract method 0xb6c25af0.
//
// Solidity: function getRewardsSignInfo(uint24 _rewardEpochId) view returns(uint64 _rewardsSignStartTs, uint64 _rewardsSignStartBlock, uint64 _rewardsSignEndTs, uint64 _rewardsSignEndBlock)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetRewardsSignInfo(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	RewardsSignStartTs    uint64
	RewardsSignStartBlock uint64
	RewardsSignEndTs      uint64
	RewardsSignEndBlock   uint64
}, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getRewardsSignInfo", _rewardEpochId)

	outstruct := new(struct {
		RewardsSignStartTs    uint64
		RewardsSignStartBlock uint64
		RewardsSignEndTs      uint64
		RewardsSignEndBlock   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardsSignStartTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.RewardsSignStartBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.RewardsSignEndTs = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.RewardsSignEndBlock = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetRewardsSignInfo is a free data retrieval call binding the contract method 0xb6c25af0.
//
// Solidity: function getRewardsSignInfo(uint24 _rewardEpochId) view returns(uint64 _rewardsSignStartTs, uint64 _rewardsSignStartBlock, uint64 _rewardsSignEndTs, uint64 _rewardsSignEndBlock)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetRewardsSignInfo(_rewardEpochId *big.Int) (struct {
	RewardsSignStartTs    uint64
	RewardsSignStartBlock uint64
	RewardsSignEndTs      uint64
	RewardsSignEndBlock   uint64
}, error) {
	return _FlareSystemsManager.Contract.GetRewardsSignInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetRewardsSignInfo is a free data retrieval call binding the contract method 0xb6c25af0.
//
// Solidity: function getRewardsSignInfo(uint24 _rewardEpochId) view returns(uint64 _rewardsSignStartTs, uint64 _rewardsSignStartBlock, uint64 _rewardsSignEndTs, uint64 _rewardsSignEndBlock)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetRewardsSignInfo(_rewardEpochId *big.Int) (struct {
	RewardsSignStartTs    uint64
	RewardsSignStartBlock uint64
	RewardsSignEndTs      uint64
	RewardsSignEndBlock   uint64
}, error) {
	return _FlareSystemsManager.Contract.GetRewardsSignInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetSeed is a free data retrieval call binding the contract method 0xe0d4ea37.
//
// Solidity: function getSeed(uint256 _rewardEpochId) view returns(uint256)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetSeed(opts *bind.CallOpts, _rewardEpochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getSeed", _rewardEpochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSeed is a free data retrieval call binding the contract method 0xe0d4ea37.
//
// Solidity: function getSeed(uint256 _rewardEpochId) view returns(uint256)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetSeed(_rewardEpochId *big.Int) (*big.Int, error) {
	return _FlareSystemsManager.Contract.GetSeed(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetSeed is a free data retrieval call binding the contract method 0xe0d4ea37.
//
// Solidity: function getSeed(uint256 _rewardEpochId) view returns(uint256)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetSeed(_rewardEpochId *big.Int) (*big.Int, error) {
	return _FlareSystemsManager.Contract.GetSeed(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetSigningPolicySignInfo is a free data retrieval call binding the contract method 0xd2e9ad71.
//
// Solidity: function getSigningPolicySignInfo(uint24 _rewardEpochId) view returns(uint64 _signingPolicySignStartTs, uint64 _signingPolicySignStartBlock, uint64 _signingPolicySignEndTs, uint64 _signingPolicySignEndBlock)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetSigningPolicySignInfo(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	SigningPolicySignStartTs    uint64
	SigningPolicySignStartBlock uint64
	SigningPolicySignEndTs      uint64
	SigningPolicySignEndBlock   uint64
}, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getSigningPolicySignInfo", _rewardEpochId)

	outstruct := new(struct {
		SigningPolicySignStartTs    uint64
		SigningPolicySignStartBlock uint64
		SigningPolicySignEndTs      uint64
		SigningPolicySignEndBlock   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SigningPolicySignStartTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.SigningPolicySignStartBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.SigningPolicySignEndTs = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.SigningPolicySignEndBlock = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetSigningPolicySignInfo is a free data retrieval call binding the contract method 0xd2e9ad71.
//
// Solidity: function getSigningPolicySignInfo(uint24 _rewardEpochId) view returns(uint64 _signingPolicySignStartTs, uint64 _signingPolicySignStartBlock, uint64 _signingPolicySignEndTs, uint64 _signingPolicySignEndBlock)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetSigningPolicySignInfo(_rewardEpochId *big.Int) (struct {
	SigningPolicySignStartTs    uint64
	SigningPolicySignStartBlock uint64
	SigningPolicySignEndTs      uint64
	SigningPolicySignEndBlock   uint64
}, error) {
	return _FlareSystemsManager.Contract.GetSigningPolicySignInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetSigningPolicySignInfo is a free data retrieval call binding the contract method 0xd2e9ad71.
//
// Solidity: function getSigningPolicySignInfo(uint24 _rewardEpochId) view returns(uint64 _signingPolicySignStartTs, uint64 _signingPolicySignStartBlock, uint64 _signingPolicySignEndTs, uint64 _signingPolicySignEndBlock)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetSigningPolicySignInfo(_rewardEpochId *big.Int) (struct {
	SigningPolicySignStartTs    uint64
	SigningPolicySignStartBlock uint64
	SigningPolicySignEndTs      uint64
	SigningPolicySignEndBlock   uint64
}, error) {
	return _FlareSystemsManager.Contract.GetSigningPolicySignInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetStartVotingRoundId is a free data retrieval call binding the contract method 0x75d2187a.
//
// Solidity: function getStartVotingRoundId(uint256 _rewardEpochId) view returns(uint32)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetStartVotingRoundId(opts *bind.CallOpts, _rewardEpochId *big.Int) (uint32, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getStartVotingRoundId", _rewardEpochId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetStartVotingRoundId is a free data retrieval call binding the contract method 0x75d2187a.
//
// Solidity: function getStartVotingRoundId(uint256 _rewardEpochId) view returns(uint32)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetStartVotingRoundId(_rewardEpochId *big.Int) (uint32, error) {
	return _FlareSystemsManager.Contract.GetStartVotingRoundId(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetStartVotingRoundId is a free data retrieval call binding the contract method 0x75d2187a.
//
// Solidity: function getStartVotingRoundId(uint256 _rewardEpochId) view returns(uint32)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetStartVotingRoundId(_rewardEpochId *big.Int) (uint32, error) {
	return _FlareSystemsManager.Contract.GetStartVotingRoundId(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetThreshold is a free data retrieval call binding the contract method 0x4615d5e9.
//
// Solidity: function getThreshold(uint256 _rewardEpochId) view returns(uint16)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetThreshold(opts *bind.CallOpts, _rewardEpochId *big.Int) (uint16, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getThreshold", _rewardEpochId)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0x4615d5e9.
//
// Solidity: function getThreshold(uint256 _rewardEpochId) view returns(uint16)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetThreshold(_rewardEpochId *big.Int) (uint16, error) {
	return _FlareSystemsManager.Contract.GetThreshold(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetThreshold is a free data retrieval call binding the contract method 0x4615d5e9.
//
// Solidity: function getThreshold(uint256 _rewardEpochId) view returns(uint16)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetThreshold(_rewardEpochId *big.Int) (uint16, error) {
	return _FlareSystemsManager.Contract.GetThreshold(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetUptimeVoteSignStartInfo is a free data retrieval call binding the contract method 0xc9f1d2aa.
//
// Solidity: function getUptimeVoteSignStartInfo(uint24 _rewardEpochId) view returns(uint64 _uptimeVoteSignStartTs, uint64 _uptimeVoteSignStartBlock)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetUptimeVoteSignStartInfo(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	UptimeVoteSignStartTs    uint64
	UptimeVoteSignStartBlock uint64
}, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getUptimeVoteSignStartInfo", _rewardEpochId)

	outstruct := new(struct {
		UptimeVoteSignStartTs    uint64
		UptimeVoteSignStartBlock uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UptimeVoteSignStartTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.UptimeVoteSignStartBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetUptimeVoteSignStartInfo is a free data retrieval call binding the contract method 0xc9f1d2aa.
//
// Solidity: function getUptimeVoteSignStartInfo(uint24 _rewardEpochId) view returns(uint64 _uptimeVoteSignStartTs, uint64 _uptimeVoteSignStartBlock)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetUptimeVoteSignStartInfo(_rewardEpochId *big.Int) (struct {
	UptimeVoteSignStartTs    uint64
	UptimeVoteSignStartBlock uint64
}, error) {
	return _FlareSystemsManager.Contract.GetUptimeVoteSignStartInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetUptimeVoteSignStartInfo is a free data retrieval call binding the contract method 0xc9f1d2aa.
//
// Solidity: function getUptimeVoteSignStartInfo(uint24 _rewardEpochId) view returns(uint64 _uptimeVoteSignStartTs, uint64 _uptimeVoteSignStartBlock)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetUptimeVoteSignStartInfo(_rewardEpochId *big.Int) (struct {
	UptimeVoteSignStartTs    uint64
	UptimeVoteSignStartBlock uint64
}, error) {
	return _FlareSystemsManager.Contract.GetUptimeVoteSignStartInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetVotePowerBlock is a free data retrieval call binding the contract method 0xc2632216.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpochId) view returns(uint64 _votePowerBlock)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetVotePowerBlock(opts *bind.CallOpts, _rewardEpochId *big.Int) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getVotePowerBlock", _rewardEpochId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetVotePowerBlock is a free data retrieval call binding the contract method 0xc2632216.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpochId) view returns(uint64 _votePowerBlock)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetVotePowerBlock(_rewardEpochId *big.Int) (uint64, error) {
	return _FlareSystemsManager.Contract.GetVotePowerBlock(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetVotePowerBlock is a free data retrieval call binding the contract method 0xc2632216.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpochId) view returns(uint64 _votePowerBlock)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetVotePowerBlock(_rewardEpochId *big.Int) (uint64, error) {
	return _FlareSystemsManager.Contract.GetVotePowerBlock(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetVoterRegistrationData is a free data retrieval call binding the contract method 0x1703a788.
//
// Solidity: function getVoterRegistrationData(uint256 _rewardEpochId) view returns(uint256 _votePowerBlock, bool _enabled)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetVoterRegistrationData(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	VotePowerBlock *big.Int
	Enabled        bool
}, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getVoterRegistrationData", _rewardEpochId)

	outstruct := new(struct {
		VotePowerBlock *big.Int
		Enabled        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VotePowerBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Enabled = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetVoterRegistrationData is a free data retrieval call binding the contract method 0x1703a788.
//
// Solidity: function getVoterRegistrationData(uint256 _rewardEpochId) view returns(uint256 _votePowerBlock, bool _enabled)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetVoterRegistrationData(_rewardEpochId *big.Int) (struct {
	VotePowerBlock *big.Int
	Enabled        bool
}, error) {
	return _FlareSystemsManager.Contract.GetVoterRegistrationData(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetVoterRegistrationData is a free data retrieval call binding the contract method 0x1703a788.
//
// Solidity: function getVoterRegistrationData(uint256 _rewardEpochId) view returns(uint256 _votePowerBlock, bool _enabled)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetVoterRegistrationData(_rewardEpochId *big.Int) (struct {
	VotePowerBlock *big.Int
	Enabled        bool
}, error) {
	return _FlareSystemsManager.Contract.GetVoterRegistrationData(&_FlareSystemsManager.CallOpts, _rewardEpochId)
}

// GetVoterRewardsSignInfo is a free data retrieval call binding the contract method 0x1916e915.
//
// Solidity: function getVoterRewardsSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _rewardsSignTs, uint64 _rewardsSignBlock)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetVoterRewardsSignInfo(opts *bind.CallOpts, _rewardEpochId *big.Int, _voter common.Address) (struct {
	RewardsSignTs    uint64
	RewardsSignBlock uint64
}, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getVoterRewardsSignInfo", _rewardEpochId, _voter)

	outstruct := new(struct {
		RewardsSignTs    uint64
		RewardsSignBlock uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardsSignTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.RewardsSignBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetVoterRewardsSignInfo is a free data retrieval call binding the contract method 0x1916e915.
//
// Solidity: function getVoterRewardsSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _rewardsSignTs, uint64 _rewardsSignBlock)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetVoterRewardsSignInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	RewardsSignTs    uint64
	RewardsSignBlock uint64
}, error) {
	return _FlareSystemsManager.Contract.GetVoterRewardsSignInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterRewardsSignInfo is a free data retrieval call binding the contract method 0x1916e915.
//
// Solidity: function getVoterRewardsSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _rewardsSignTs, uint64 _rewardsSignBlock)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetVoterRewardsSignInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	RewardsSignTs    uint64
	RewardsSignBlock uint64
}, error) {
	return _FlareSystemsManager.Contract.GetVoterRewardsSignInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterSigningPolicySignInfo is a free data retrieval call binding the contract method 0xdac4319d.
//
// Solidity: function getVoterSigningPolicySignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _signingPolicySignTs, uint64 _signingPolicySignBlock)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetVoterSigningPolicySignInfo(opts *bind.CallOpts, _rewardEpochId *big.Int, _voter common.Address) (struct {
	SigningPolicySignTs    uint64
	SigningPolicySignBlock uint64
}, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getVoterSigningPolicySignInfo", _rewardEpochId, _voter)

	outstruct := new(struct {
		SigningPolicySignTs    uint64
		SigningPolicySignBlock uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SigningPolicySignTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.SigningPolicySignBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetVoterSigningPolicySignInfo is a free data retrieval call binding the contract method 0xdac4319d.
//
// Solidity: function getVoterSigningPolicySignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _signingPolicySignTs, uint64 _signingPolicySignBlock)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetVoterSigningPolicySignInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	SigningPolicySignTs    uint64
	SigningPolicySignBlock uint64
}, error) {
	return _FlareSystemsManager.Contract.GetVoterSigningPolicySignInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterSigningPolicySignInfo is a free data retrieval call binding the contract method 0xdac4319d.
//
// Solidity: function getVoterSigningPolicySignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _signingPolicySignTs, uint64 _signingPolicySignBlock)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetVoterSigningPolicySignInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	SigningPolicySignTs    uint64
	SigningPolicySignBlock uint64
}, error) {
	return _FlareSystemsManager.Contract.GetVoterSigningPolicySignInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterUptimeVoteSignInfo is a free data retrieval call binding the contract method 0x41c05ad5.
//
// Solidity: function getVoterUptimeVoteSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSignTs, uint64 _uptimeVoteSignBlock)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetVoterUptimeVoteSignInfo(opts *bind.CallOpts, _rewardEpochId *big.Int, _voter common.Address) (struct {
	UptimeVoteSignTs    uint64
	UptimeVoteSignBlock uint64
}, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getVoterUptimeVoteSignInfo", _rewardEpochId, _voter)

	outstruct := new(struct {
		UptimeVoteSignTs    uint64
		UptimeVoteSignBlock uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UptimeVoteSignTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.UptimeVoteSignBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetVoterUptimeVoteSignInfo is a free data retrieval call binding the contract method 0x41c05ad5.
//
// Solidity: function getVoterUptimeVoteSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSignTs, uint64 _uptimeVoteSignBlock)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetVoterUptimeVoteSignInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	UptimeVoteSignTs    uint64
	UptimeVoteSignBlock uint64
}, error) {
	return _FlareSystemsManager.Contract.GetVoterUptimeVoteSignInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterUptimeVoteSignInfo is a free data retrieval call binding the contract method 0x41c05ad5.
//
// Solidity: function getVoterUptimeVoteSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSignTs, uint64 _uptimeVoteSignBlock)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetVoterUptimeVoteSignInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	UptimeVoteSignTs    uint64
	UptimeVoteSignBlock uint64
}, error) {
	return _FlareSystemsManager.Contract.GetVoterUptimeVoteSignInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterUptimeVoteSubmitInfo is a free data retrieval call binding the contract method 0x59db0e2f.
//
// Solidity: function getVoterUptimeVoteSubmitInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSubmitTs, uint64 _uptimeVoteSubmitBlock)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GetVoterUptimeVoteSubmitInfo(opts *bind.CallOpts, _rewardEpochId *big.Int, _voter common.Address) (struct {
	UptimeVoteSubmitTs    uint64
	UptimeVoteSubmitBlock uint64
}, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "getVoterUptimeVoteSubmitInfo", _rewardEpochId, _voter)

	outstruct := new(struct {
		UptimeVoteSubmitTs    uint64
		UptimeVoteSubmitBlock uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UptimeVoteSubmitTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.UptimeVoteSubmitBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetVoterUptimeVoteSubmitInfo is a free data retrieval call binding the contract method 0x59db0e2f.
//
// Solidity: function getVoterUptimeVoteSubmitInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSubmitTs, uint64 _uptimeVoteSubmitBlock)
func (_FlareSystemsManager *FlareSystemsManagerSession) GetVoterUptimeVoteSubmitInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	UptimeVoteSubmitTs    uint64
	UptimeVoteSubmitBlock uint64
}, error) {
	return _FlareSystemsManager.Contract.GetVoterUptimeVoteSubmitInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterUptimeVoteSubmitInfo is a free data retrieval call binding the contract method 0x59db0e2f.
//
// Solidity: function getVoterUptimeVoteSubmitInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSubmitTs, uint64 _uptimeVoteSubmitBlock)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GetVoterUptimeVoteSubmitInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	UptimeVoteSubmitTs    uint64
	UptimeVoteSubmitBlock uint64
}, error) {
	return _FlareSystemsManager.Contract.GetVoterUptimeVoteSubmitInfo(&_FlareSystemsManager.CallOpts, _rewardEpochId, _voter)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerSession) Governance() (common.Address, error) {
	return _FlareSystemsManager.Contract.Governance(&_FlareSystemsManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) Governance() (common.Address, error) {
	return _FlareSystemsManager.Contract.Governance(&_FlareSystemsManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerSession) GovernanceSettings() (common.Address, error) {
	return _FlareSystemsManager.Contract.GovernanceSettings(&_FlareSystemsManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) GovernanceSettings() (common.Address, error) {
	return _FlareSystemsManager.Contract.GovernanceSettings(&_FlareSystemsManager.CallOpts)
}

// InitialRandomVotePowerBlockSelectionSize is a free data retrieval call binding the contract method 0xded7c4b8.
//
// Solidity: function initialRandomVotePowerBlockSelectionSize() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCaller) InitialRandomVotePowerBlockSelectionSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "initialRandomVotePowerBlockSelectionSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InitialRandomVotePowerBlockSelectionSize is a free data retrieval call binding the contract method 0xded7c4b8.
//
// Solidity: function initialRandomVotePowerBlockSelectionSize() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerSession) InitialRandomVotePowerBlockSelectionSize() (uint64, error) {
	return _FlareSystemsManager.Contract.InitialRandomVotePowerBlockSelectionSize(&_FlareSystemsManager.CallOpts)
}

// InitialRandomVotePowerBlockSelectionSize is a free data retrieval call binding the contract method 0xded7c4b8.
//
// Solidity: function initialRandomVotePowerBlockSelectionSize() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) InitialRandomVotePowerBlockSelectionSize() (uint64, error) {
	return _FlareSystemsManager.Contract.InitialRandomVotePowerBlockSelectionSize(&_FlareSystemsManager.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerSession) IsExecutor(_address common.Address) (bool, error) {
	return _FlareSystemsManager.Contract.IsExecutor(&_FlareSystemsManager.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _FlareSystemsManager.Contract.IsExecutor(&_FlareSystemsManager.CallOpts, _address)
}

// IsVoterRegistrationEnabled is a free data retrieval call binding the contract method 0x09505d25.
//
// Solidity: function isVoterRegistrationEnabled() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerCaller) IsVoterRegistrationEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "isVoterRegistrationEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoterRegistrationEnabled is a free data retrieval call binding the contract method 0x09505d25.
//
// Solidity: function isVoterRegistrationEnabled() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerSession) IsVoterRegistrationEnabled() (bool, error) {
	return _FlareSystemsManager.Contract.IsVoterRegistrationEnabled(&_FlareSystemsManager.CallOpts)
}

// IsVoterRegistrationEnabled is a free data retrieval call binding the contract method 0x09505d25.
//
// Solidity: function isVoterRegistrationEnabled() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) IsVoterRegistrationEnabled() (bool, error) {
	return _FlareSystemsManager.Contract.IsVoterRegistrationEnabled(&_FlareSystemsManager.CallOpts)
}

// LastInitializedVotingRoundId is a free data retrieval call binding the contract method 0x4f923d37.
//
// Solidity: function lastInitializedVotingRoundId() view returns(uint32)
func (_FlareSystemsManager *FlareSystemsManagerCaller) LastInitializedVotingRoundId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "lastInitializedVotingRoundId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastInitializedVotingRoundId is a free data retrieval call binding the contract method 0x4f923d37.
//
// Solidity: function lastInitializedVotingRoundId() view returns(uint32)
func (_FlareSystemsManager *FlareSystemsManagerSession) LastInitializedVotingRoundId() (uint32, error) {
	return _FlareSystemsManager.Contract.LastInitializedVotingRoundId(&_FlareSystemsManager.CallOpts)
}

// LastInitializedVotingRoundId is a free data retrieval call binding the contract method 0x4f923d37.
//
// Solidity: function lastInitializedVotingRoundId() view returns(uint32)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) LastInitializedVotingRoundId() (uint32, error) {
	return _FlareSystemsManager.Contract.LastInitializedVotingRoundId(&_FlareSystemsManager.CallOpts)
}

// NewSigningPolicyInitializationStartSeconds is a free data retrieval call binding the contract method 0x6aeffddc.
//
// Solidity: function newSigningPolicyInitializationStartSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCaller) NewSigningPolicyInitializationStartSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "newSigningPolicyInitializationStartSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NewSigningPolicyInitializationStartSeconds is a free data retrieval call binding the contract method 0x6aeffddc.
//
// Solidity: function newSigningPolicyInitializationStartSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerSession) NewSigningPolicyInitializationStartSeconds() (uint64, error) {
	return _FlareSystemsManager.Contract.NewSigningPolicyInitializationStartSeconds(&_FlareSystemsManager.CallOpts)
}

// NewSigningPolicyInitializationStartSeconds is a free data retrieval call binding the contract method 0x6aeffddc.
//
// Solidity: function newSigningPolicyInitializationStartSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) NewSigningPolicyInitializationStartSeconds() (uint64, error) {
	return _FlareSystemsManager.Contract.NewSigningPolicyInitializationStartSeconds(&_FlareSystemsManager.CallOpts)
}

// NewSigningPolicyMinNumberOfVotingRoundsDelay is a free data retrieval call binding the contract method 0xa733d54b.
//
// Solidity: function newSigningPolicyMinNumberOfVotingRoundsDelay() view returns(uint32)
func (_FlareSystemsManager *FlareSystemsManagerCaller) NewSigningPolicyMinNumberOfVotingRoundsDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "newSigningPolicyMinNumberOfVotingRoundsDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NewSigningPolicyMinNumberOfVotingRoundsDelay is a free data retrieval call binding the contract method 0xa733d54b.
//
// Solidity: function newSigningPolicyMinNumberOfVotingRoundsDelay() view returns(uint32)
func (_FlareSystemsManager *FlareSystemsManagerSession) NewSigningPolicyMinNumberOfVotingRoundsDelay() (uint32, error) {
	return _FlareSystemsManager.Contract.NewSigningPolicyMinNumberOfVotingRoundsDelay(&_FlareSystemsManager.CallOpts)
}

// NewSigningPolicyMinNumberOfVotingRoundsDelay is a free data retrieval call binding the contract method 0xa733d54b.
//
// Solidity: function newSigningPolicyMinNumberOfVotingRoundsDelay() view returns(uint32)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) NewSigningPolicyMinNumberOfVotingRoundsDelay() (uint32, error) {
	return _FlareSystemsManager.Contract.NewSigningPolicyMinNumberOfVotingRoundsDelay(&_FlareSystemsManager.CallOpts)
}

// NoOfWeightBasedClaims is a free data retrieval call binding the contract method 0x388a4c3d.
//
// Solidity: function noOfWeightBasedClaims(uint256 rewardEpochId) view returns(uint256)
func (_FlareSystemsManager *FlareSystemsManagerCaller) NoOfWeightBasedClaims(opts *bind.CallOpts, rewardEpochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "noOfWeightBasedClaims", rewardEpochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoOfWeightBasedClaims is a free data retrieval call binding the contract method 0x388a4c3d.
//
// Solidity: function noOfWeightBasedClaims(uint256 rewardEpochId) view returns(uint256)
func (_FlareSystemsManager *FlareSystemsManagerSession) NoOfWeightBasedClaims(rewardEpochId *big.Int) (*big.Int, error) {
	return _FlareSystemsManager.Contract.NoOfWeightBasedClaims(&_FlareSystemsManager.CallOpts, rewardEpochId)
}

// NoOfWeightBasedClaims is a free data retrieval call binding the contract method 0x388a4c3d.
//
// Solidity: function noOfWeightBasedClaims(uint256 rewardEpochId) view returns(uint256)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) NoOfWeightBasedClaims(rewardEpochId *big.Int) (*big.Int, error) {
	return _FlareSystemsManager.Contract.NoOfWeightBasedClaims(&_FlareSystemsManager.CallOpts, rewardEpochId)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerSession) ProductionMode() (bool, error) {
	return _FlareSystemsManager.Contract.ProductionMode(&_FlareSystemsManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) ProductionMode() (bool, error) {
	return _FlareSystemsManager.Contract.ProductionMode(&_FlareSystemsManager.CallOpts)
}

// RandomAcquisitionMaxDurationBlocks is a free data retrieval call binding the contract method 0x490344f4.
//
// Solidity: function randomAcquisitionMaxDurationBlocks() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCaller) RandomAcquisitionMaxDurationBlocks(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "randomAcquisitionMaxDurationBlocks")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RandomAcquisitionMaxDurationBlocks is a free data retrieval call binding the contract method 0x490344f4.
//
// Solidity: function randomAcquisitionMaxDurationBlocks() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerSession) RandomAcquisitionMaxDurationBlocks() (uint64, error) {
	return _FlareSystemsManager.Contract.RandomAcquisitionMaxDurationBlocks(&_FlareSystemsManager.CallOpts)
}

// RandomAcquisitionMaxDurationBlocks is a free data retrieval call binding the contract method 0x490344f4.
//
// Solidity: function randomAcquisitionMaxDurationBlocks() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) RandomAcquisitionMaxDurationBlocks() (uint64, error) {
	return _FlareSystemsManager.Contract.RandomAcquisitionMaxDurationBlocks(&_FlareSystemsManager.CallOpts)
}

// RandomAcquisitionMaxDurationSeconds is a free data retrieval call binding the contract method 0x098e7ff6.
//
// Solidity: function randomAcquisitionMaxDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCaller) RandomAcquisitionMaxDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "randomAcquisitionMaxDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RandomAcquisitionMaxDurationSeconds is a free data retrieval call binding the contract method 0x098e7ff6.
//
// Solidity: function randomAcquisitionMaxDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerSession) RandomAcquisitionMaxDurationSeconds() (uint64, error) {
	return _FlareSystemsManager.Contract.RandomAcquisitionMaxDurationSeconds(&_FlareSystemsManager.CallOpts)
}

// RandomAcquisitionMaxDurationSeconds is a free data retrieval call binding the contract method 0x098e7ff6.
//
// Solidity: function randomAcquisitionMaxDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) RandomAcquisitionMaxDurationSeconds() (uint64, error) {
	return _FlareSystemsManager.Contract.RandomAcquisitionMaxDurationSeconds(&_FlareSystemsManager.CallOpts)
}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCaller) Relay(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "relay")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerSession) Relay() (common.Address, error) {
	return _FlareSystemsManager.Contract.Relay(&_FlareSystemsManager.CallOpts)
}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) Relay() (common.Address, error) {
	return _FlareSystemsManager.Contract.Relay(&_FlareSystemsManager.CallOpts)
}

// RewardEpochDurationSeconds is a free data retrieval call binding the contract method 0x85f3c9c9.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCaller) RewardEpochDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "rewardEpochDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RewardEpochDurationSeconds is a free data retrieval call binding the contract method 0x85f3c9c9.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerSession) RewardEpochDurationSeconds() (uint64, error) {
	return _FlareSystemsManager.Contract.RewardEpochDurationSeconds(&_FlareSystemsManager.CallOpts)
}

// RewardEpochDurationSeconds is a free data retrieval call binding the contract method 0x85f3c9c9.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) RewardEpochDurationSeconds() (uint64, error) {
	return _FlareSystemsManager.Contract.RewardEpochDurationSeconds(&_FlareSystemsManager.CallOpts)
}

// RewardEpochIdToExpireNext is a free data retrieval call binding the contract method 0xaec84ab6.
//
// Solidity: function rewardEpochIdToExpireNext() view returns(uint24)
func (_FlareSystemsManager *FlareSystemsManagerCaller) RewardEpochIdToExpireNext(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "rewardEpochIdToExpireNext")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardEpochIdToExpireNext is a free data retrieval call binding the contract method 0xaec84ab6.
//
// Solidity: function rewardEpochIdToExpireNext() view returns(uint24)
func (_FlareSystemsManager *FlareSystemsManagerSession) RewardEpochIdToExpireNext() (*big.Int, error) {
	return _FlareSystemsManager.Contract.RewardEpochIdToExpireNext(&_FlareSystemsManager.CallOpts)
}

// RewardEpochIdToExpireNext is a free data retrieval call binding the contract method 0xaec84ab6.
//
// Solidity: function rewardEpochIdToExpireNext() view returns(uint24)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) RewardEpochIdToExpireNext() (*big.Int, error) {
	return _FlareSystemsManager.Contract.RewardEpochIdToExpireNext(&_FlareSystemsManager.CallOpts)
}

// RewardExpiryOffsetSeconds is a free data retrieval call binding the contract method 0x4eaee307.
//
// Solidity: function rewardExpiryOffsetSeconds() view returns(uint32)
func (_FlareSystemsManager *FlareSystemsManagerCaller) RewardExpiryOffsetSeconds(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "rewardExpiryOffsetSeconds")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RewardExpiryOffsetSeconds is a free data retrieval call binding the contract method 0x4eaee307.
//
// Solidity: function rewardExpiryOffsetSeconds() view returns(uint32)
func (_FlareSystemsManager *FlareSystemsManagerSession) RewardExpiryOffsetSeconds() (uint32, error) {
	return _FlareSystemsManager.Contract.RewardExpiryOffsetSeconds(&_FlareSystemsManager.CallOpts)
}

// RewardExpiryOffsetSeconds is a free data retrieval call binding the contract method 0x4eaee307.
//
// Solidity: function rewardExpiryOffsetSeconds() view returns(uint32)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) RewardExpiryOffsetSeconds() (uint32, error) {
	return _FlareSystemsManager.Contract.RewardExpiryOffsetSeconds(&_FlareSystemsManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerSession) RewardManager() (common.Address, error) {
	return _FlareSystemsManager.Contract.RewardManager(&_FlareSystemsManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) RewardManager() (common.Address, error) {
	return _FlareSystemsManager.Contract.RewardManager(&_FlareSystemsManager.CallOpts)
}

// RewardsHash is a free data retrieval call binding the contract method 0x647006e2.
//
// Solidity: function rewardsHash(uint256 rewardEpochId) view returns(bytes32)
func (_FlareSystemsManager *FlareSystemsManagerCaller) RewardsHash(opts *bind.CallOpts, rewardEpochId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "rewardsHash", rewardEpochId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RewardsHash is a free data retrieval call binding the contract method 0x647006e2.
//
// Solidity: function rewardsHash(uint256 rewardEpochId) view returns(bytes32)
func (_FlareSystemsManager *FlareSystemsManagerSession) RewardsHash(rewardEpochId *big.Int) ([32]byte, error) {
	return _FlareSystemsManager.Contract.RewardsHash(&_FlareSystemsManager.CallOpts, rewardEpochId)
}

// RewardsHash is a free data retrieval call binding the contract method 0x647006e2.
//
// Solidity: function rewardsHash(uint256 rewardEpochId) view returns(bytes32)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) RewardsHash(rewardEpochId *big.Int) ([32]byte, error) {
	return _FlareSystemsManager.Contract.RewardsHash(&_FlareSystemsManager.CallOpts, rewardEpochId)
}

// SigningPolicyMinNumberOfVoters is a free data retrieval call binding the contract method 0x2e3645f8.
//
// Solidity: function signingPolicyMinNumberOfVoters() view returns(uint16)
func (_FlareSystemsManager *FlareSystemsManagerCaller) SigningPolicyMinNumberOfVoters(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "signingPolicyMinNumberOfVoters")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SigningPolicyMinNumberOfVoters is a free data retrieval call binding the contract method 0x2e3645f8.
//
// Solidity: function signingPolicyMinNumberOfVoters() view returns(uint16)
func (_FlareSystemsManager *FlareSystemsManagerSession) SigningPolicyMinNumberOfVoters() (uint16, error) {
	return _FlareSystemsManager.Contract.SigningPolicyMinNumberOfVoters(&_FlareSystemsManager.CallOpts)
}

// SigningPolicyMinNumberOfVoters is a free data retrieval call binding the contract method 0x2e3645f8.
//
// Solidity: function signingPolicyMinNumberOfVoters() view returns(uint16)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) SigningPolicyMinNumberOfVoters() (uint16, error) {
	return _FlareSystemsManager.Contract.SigningPolicyMinNumberOfVoters(&_FlareSystemsManager.CallOpts)
}

// SigningPolicyThresholdPPM is a free data retrieval call binding the contract method 0xf21d6304.
//
// Solidity: function signingPolicyThresholdPPM() view returns(uint24)
func (_FlareSystemsManager *FlareSystemsManagerCaller) SigningPolicyThresholdPPM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "signingPolicyThresholdPPM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SigningPolicyThresholdPPM is a free data retrieval call binding the contract method 0xf21d6304.
//
// Solidity: function signingPolicyThresholdPPM() view returns(uint24)
func (_FlareSystemsManager *FlareSystemsManagerSession) SigningPolicyThresholdPPM() (*big.Int, error) {
	return _FlareSystemsManager.Contract.SigningPolicyThresholdPPM(&_FlareSystemsManager.CallOpts)
}

// SigningPolicyThresholdPPM is a free data retrieval call binding the contract method 0xf21d6304.
//
// Solidity: function signingPolicyThresholdPPM() view returns(uint24)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) SigningPolicyThresholdPPM() (*big.Int, error) {
	return _FlareSystemsManager.Contract.SigningPolicyThresholdPPM(&_FlareSystemsManager.CallOpts)
}

// Submission is a free data retrieval call binding the contract method 0x9a759097.
//
// Solidity: function submission() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCaller) Submission(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "submission")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Submission is a free data retrieval call binding the contract method 0x9a759097.
//
// Solidity: function submission() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerSession) Submission() (common.Address, error) {
	return _FlareSystemsManager.Contract.Submission(&_FlareSystemsManager.CallOpts)
}

// Submission is a free data retrieval call binding the contract method 0x9a759097.
//
// Solidity: function submission() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) Submission() (common.Address, error) {
	return _FlareSystemsManager.Contract.Submission(&_FlareSystemsManager.CallOpts)
}

// Submit3Aligned is a free data retrieval call binding the contract method 0x107d8ffb.
//
// Solidity: function submit3Aligned() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerCaller) Submit3Aligned(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "submit3Aligned")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Submit3Aligned is a free data retrieval call binding the contract method 0x107d8ffb.
//
// Solidity: function submit3Aligned() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerSession) Submit3Aligned() (bool, error) {
	return _FlareSystemsManager.Contract.Submit3Aligned(&_FlareSystemsManager.CallOpts)
}

// Submit3Aligned is a free data retrieval call binding the contract method 0x107d8ffb.
//
// Solidity: function submit3Aligned() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) Submit3Aligned() (bool, error) {
	return _FlareSystemsManager.Contract.Submit3Aligned(&_FlareSystemsManager.CallOpts)
}

// SubmitUptimeVoteMinDurationBlocks is a free data retrieval call binding the contract method 0xd8a01a0a.
//
// Solidity: function submitUptimeVoteMinDurationBlocks() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCaller) SubmitUptimeVoteMinDurationBlocks(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "submitUptimeVoteMinDurationBlocks")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SubmitUptimeVoteMinDurationBlocks is a free data retrieval call binding the contract method 0xd8a01a0a.
//
// Solidity: function submitUptimeVoteMinDurationBlocks() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerSession) SubmitUptimeVoteMinDurationBlocks() (uint64, error) {
	return _FlareSystemsManager.Contract.SubmitUptimeVoteMinDurationBlocks(&_FlareSystemsManager.CallOpts)
}

// SubmitUptimeVoteMinDurationBlocks is a free data retrieval call binding the contract method 0xd8a01a0a.
//
// Solidity: function submitUptimeVoteMinDurationBlocks() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) SubmitUptimeVoteMinDurationBlocks() (uint64, error) {
	return _FlareSystemsManager.Contract.SubmitUptimeVoteMinDurationBlocks(&_FlareSystemsManager.CallOpts)
}

// SubmitUptimeVoteMinDurationSeconds is a free data retrieval call binding the contract method 0x4c528765.
//
// Solidity: function submitUptimeVoteMinDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCaller) SubmitUptimeVoteMinDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "submitUptimeVoteMinDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SubmitUptimeVoteMinDurationSeconds is a free data retrieval call binding the contract method 0x4c528765.
//
// Solidity: function submitUptimeVoteMinDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerSession) SubmitUptimeVoteMinDurationSeconds() (uint64, error) {
	return _FlareSystemsManager.Contract.SubmitUptimeVoteMinDurationSeconds(&_FlareSystemsManager.CallOpts)
}

// SubmitUptimeVoteMinDurationSeconds is a free data retrieval call binding the contract method 0x4c528765.
//
// Solidity: function submitUptimeVoteMinDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) SubmitUptimeVoteMinDurationSeconds() (uint64, error) {
	return _FlareSystemsManager.Contract.SubmitUptimeVoteMinDurationSeconds(&_FlareSystemsManager.CallOpts)
}

// SwitchToFallbackMode is a free data retrieval call binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerCaller) SwitchToFallbackMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "switchToFallbackMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwitchToFallbackMode is a free data retrieval call binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerSession) SwitchToFallbackMode() (bool, error) {
	return _FlareSystemsManager.Contract.SwitchToFallbackMode(&_FlareSystemsManager.CallOpts)
}

// SwitchToFallbackMode is a free data retrieval call binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) SwitchToFallbackMode() (bool, error) {
	return _FlareSystemsManager.Contract.SwitchToFallbackMode(&_FlareSystemsManager.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FlareSystemsManager *FlareSystemsManagerCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "timelockedCalls", selector)

	outstruct := new(struct {
		AllowedAfterTimestamp *big.Int
		EncodedCall           []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowedAfterTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EncodedCall = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FlareSystemsManager *FlareSystemsManagerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FlareSystemsManager.Contract.TimelockedCalls(&_FlareSystemsManager.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FlareSystemsManager.Contract.TimelockedCalls(&_FlareSystemsManager.CallOpts, selector)
}

// TriggerExpirationAndCleanup is a free data retrieval call binding the contract method 0x9b760d13.
//
// Solidity: function triggerExpirationAndCleanup() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerCaller) TriggerExpirationAndCleanup(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "triggerExpirationAndCleanup")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TriggerExpirationAndCleanup is a free data retrieval call binding the contract method 0x9b760d13.
//
// Solidity: function triggerExpirationAndCleanup() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerSession) TriggerExpirationAndCleanup() (bool, error) {
	return _FlareSystemsManager.Contract.TriggerExpirationAndCleanup(&_FlareSystemsManager.CallOpts)
}

// TriggerExpirationAndCleanup is a free data retrieval call binding the contract method 0x9b760d13.
//
// Solidity: function triggerExpirationAndCleanup() view returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) TriggerExpirationAndCleanup() (bool, error) {
	return _FlareSystemsManager.Contract.TriggerExpirationAndCleanup(&_FlareSystemsManager.CallOpts)
}

// UptimeVoteHash is a free data retrieval call binding the contract method 0xd3466911.
//
// Solidity: function uptimeVoteHash(uint256 rewardEpochId) view returns(bytes32)
func (_FlareSystemsManager *FlareSystemsManagerCaller) UptimeVoteHash(opts *bind.CallOpts, rewardEpochId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "uptimeVoteHash", rewardEpochId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UptimeVoteHash is a free data retrieval call binding the contract method 0xd3466911.
//
// Solidity: function uptimeVoteHash(uint256 rewardEpochId) view returns(bytes32)
func (_FlareSystemsManager *FlareSystemsManagerSession) UptimeVoteHash(rewardEpochId *big.Int) ([32]byte, error) {
	return _FlareSystemsManager.Contract.UptimeVoteHash(&_FlareSystemsManager.CallOpts, rewardEpochId)
}

// UptimeVoteHash is a free data retrieval call binding the contract method 0xd3466911.
//
// Solidity: function uptimeVoteHash(uint256 rewardEpochId) view returns(bytes32)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) UptimeVoteHash(rewardEpochId *big.Int) ([32]byte, error) {
	return _FlareSystemsManager.Contract.UptimeVoteHash(&_FlareSystemsManager.CallOpts, rewardEpochId)
}

// VoterRegistrationMinDurationBlocks is a free data retrieval call binding the contract method 0xd10e807f.
//
// Solidity: function voterRegistrationMinDurationBlocks() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCaller) VoterRegistrationMinDurationBlocks(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "voterRegistrationMinDurationBlocks")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VoterRegistrationMinDurationBlocks is a free data retrieval call binding the contract method 0xd10e807f.
//
// Solidity: function voterRegistrationMinDurationBlocks() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerSession) VoterRegistrationMinDurationBlocks() (uint64, error) {
	return _FlareSystemsManager.Contract.VoterRegistrationMinDurationBlocks(&_FlareSystemsManager.CallOpts)
}

// VoterRegistrationMinDurationBlocks is a free data retrieval call binding the contract method 0xd10e807f.
//
// Solidity: function voterRegistrationMinDurationBlocks() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) VoterRegistrationMinDurationBlocks() (uint64, error) {
	return _FlareSystemsManager.Contract.VoterRegistrationMinDurationBlocks(&_FlareSystemsManager.CallOpts)
}

// VoterRegistrationMinDurationSeconds is a free data retrieval call binding the contract method 0xa219fe02.
//
// Solidity: function voterRegistrationMinDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCaller) VoterRegistrationMinDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "voterRegistrationMinDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VoterRegistrationMinDurationSeconds is a free data retrieval call binding the contract method 0xa219fe02.
//
// Solidity: function voterRegistrationMinDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerSession) VoterRegistrationMinDurationSeconds() (uint64, error) {
	return _FlareSystemsManager.Contract.VoterRegistrationMinDurationSeconds(&_FlareSystemsManager.CallOpts)
}

// VoterRegistrationMinDurationSeconds is a free data retrieval call binding the contract method 0xa219fe02.
//
// Solidity: function voterRegistrationMinDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) VoterRegistrationMinDurationSeconds() (uint64, error) {
	return _FlareSystemsManager.Contract.VoterRegistrationMinDurationSeconds(&_FlareSystemsManager.CallOpts)
}

// VoterRegistrationTriggerContract is a free data retrieval call binding the contract method 0x88e49ac7.
//
// Solidity: function voterRegistrationTriggerContract() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCaller) VoterRegistrationTriggerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "voterRegistrationTriggerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoterRegistrationTriggerContract is a free data retrieval call binding the contract method 0x88e49ac7.
//
// Solidity: function voterRegistrationTriggerContract() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerSession) VoterRegistrationTriggerContract() (common.Address, error) {
	return _FlareSystemsManager.Contract.VoterRegistrationTriggerContract(&_FlareSystemsManager.CallOpts)
}

// VoterRegistrationTriggerContract is a free data retrieval call binding the contract method 0x88e49ac7.
//
// Solidity: function voterRegistrationTriggerContract() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) VoterRegistrationTriggerContract() (common.Address, error) {
	return _FlareSystemsManager.Contract.VoterRegistrationTriggerContract(&_FlareSystemsManager.CallOpts)
}

// VoterRegistry is a free data retrieval call binding the contract method 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCaller) VoterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "voterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoterRegistry is a free data retrieval call binding the contract method 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerSession) VoterRegistry() (common.Address, error) {
	return _FlareSystemsManager.Contract.VoterRegistry(&_FlareSystemsManager.CallOpts)
}

// VoterRegistry is a free data retrieval call binding the contract method 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) VoterRegistry() (common.Address, error) {
	return _FlareSystemsManager.Contract.VoterRegistry(&_FlareSystemsManager.CallOpts)
}

// VotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCaller) VotingEpochDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemsManager.contract.Call(opts, &out, "votingEpochDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerSession) VotingEpochDurationSeconds() (uint64, error) {
	return _FlareSystemsManager.Contract.VotingEpochDurationSeconds(&_FlareSystemsManager.CallOpts)
}

// VotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (_FlareSystemsManager *FlareSystemsManagerCallerSession) VotingEpochDurationSeconds() (uint64, error) {
	return _FlareSystemsManager.Contract.VotingEpochDurationSeconds(&_FlareSystemsManager.CallOpts)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.CancelGovernanceCall(&_FlareSystemsManager.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.CancelGovernanceCall(&_FlareSystemsManager.TransactOpts, _selector)
}

// ChangeSigningPolicySettings is a paid mutator transaction binding the contract method 0x898805c1.
//
// Solidity: function changeSigningPolicySettings(uint24 _signingPolicyThresholdPPM, uint16 _signingPolicyMinNumberOfVoters) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) ChangeSigningPolicySettings(opts *bind.TransactOpts, _signingPolicyThresholdPPM *big.Int, _signingPolicyMinNumberOfVoters uint16) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "changeSigningPolicySettings", _signingPolicyThresholdPPM, _signingPolicyMinNumberOfVoters)
}

// ChangeSigningPolicySettings is a paid mutator transaction binding the contract method 0x898805c1.
//
// Solidity: function changeSigningPolicySettings(uint24 _signingPolicyThresholdPPM, uint16 _signingPolicyMinNumberOfVoters) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) ChangeSigningPolicySettings(_signingPolicyThresholdPPM *big.Int, _signingPolicyMinNumberOfVoters uint16) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.ChangeSigningPolicySettings(&_FlareSystemsManager.TransactOpts, _signingPolicyThresholdPPM, _signingPolicyMinNumberOfVoters)
}

// ChangeSigningPolicySettings is a paid mutator transaction binding the contract method 0x898805c1.
//
// Solidity: function changeSigningPolicySettings(uint24 _signingPolicyThresholdPPM, uint16 _signingPolicyMinNumberOfVoters) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) ChangeSigningPolicySettings(_signingPolicyThresholdPPM *big.Int, _signingPolicyMinNumberOfVoters uint16) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.ChangeSigningPolicySettings(&_FlareSystemsManager.TransactOpts, _signingPolicyThresholdPPM, _signingPolicyMinNumberOfVoters)
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerTransactor) Daemonize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "daemonize")
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerSession) Daemonize() (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.Daemonize(&_FlareSystemsManager.TransactOpts)
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) Daemonize() (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.Daemonize(&_FlareSystemsManager.TransactOpts)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.ExecuteGovernanceCall(&_FlareSystemsManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.ExecuteGovernanceCall(&_FlareSystemsManager.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.Initialise(&_FlareSystemsManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.Initialise(&_FlareSystemsManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// SetRewardEpochSwitchoverTriggerContracts is a paid mutator transaction binding the contract method 0x06886f41.
//
// Solidity: function setRewardEpochSwitchoverTriggerContracts(address[] _contracts) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) SetRewardEpochSwitchoverTriggerContracts(opts *bind.TransactOpts, _contracts []common.Address) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "setRewardEpochSwitchoverTriggerContracts", _contracts)
}

// SetRewardEpochSwitchoverTriggerContracts is a paid mutator transaction binding the contract method 0x06886f41.
//
// Solidity: function setRewardEpochSwitchoverTriggerContracts(address[] _contracts) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) SetRewardEpochSwitchoverTriggerContracts(_contracts []common.Address) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SetRewardEpochSwitchoverTriggerContracts(&_FlareSystemsManager.TransactOpts, _contracts)
}

// SetRewardEpochSwitchoverTriggerContracts is a paid mutator transaction binding the contract method 0x06886f41.
//
// Solidity: function setRewardEpochSwitchoverTriggerContracts(address[] _contracts) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) SetRewardEpochSwitchoverTriggerContracts(_contracts []common.Address) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SetRewardEpochSwitchoverTriggerContracts(&_FlareSystemsManager.TransactOpts, _contracts)
}

// SetRewardsData is a paid mutator transaction binding the contract method 0x2a576f8d.
//
// Solidity: function setRewardsData(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) SetRewardsData(opts *bind.TransactOpts, _rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "setRewardsData", _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash)
}

// SetRewardsData is a paid mutator transaction binding the contract method 0x2a576f8d.
//
// Solidity: function setRewardsData(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) SetRewardsData(_rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SetRewardsData(&_FlareSystemsManager.TransactOpts, _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash)
}

// SetRewardsData is a paid mutator transaction binding the contract method 0x2a576f8d.
//
// Solidity: function setRewardsData(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) SetRewardsData(_rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SetRewardsData(&_FlareSystemsManager.TransactOpts, _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash)
}

// SetSubmit3Aligned is a paid mutator transaction binding the contract method 0xa72b826e.
//
// Solidity: function setSubmit3Aligned(bool _submit3Aligned) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) SetSubmit3Aligned(opts *bind.TransactOpts, _submit3Aligned bool) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "setSubmit3Aligned", _submit3Aligned)
}

// SetSubmit3Aligned is a paid mutator transaction binding the contract method 0xa72b826e.
//
// Solidity: function setSubmit3Aligned(bool _submit3Aligned) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) SetSubmit3Aligned(_submit3Aligned bool) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SetSubmit3Aligned(&_FlareSystemsManager.TransactOpts, _submit3Aligned)
}

// SetSubmit3Aligned is a paid mutator transaction binding the contract method 0xa72b826e.
//
// Solidity: function setSubmit3Aligned(bool _submit3Aligned) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) SetSubmit3Aligned(_submit3Aligned bool) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SetSubmit3Aligned(&_FlareSystemsManager.TransactOpts, _submit3Aligned)
}

// SetTriggerExpirationAndCleanup is a paid mutator transaction binding the contract method 0x67daec89.
//
// Solidity: function setTriggerExpirationAndCleanup(bool _triggerExpirationAndCleanup) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) SetTriggerExpirationAndCleanup(opts *bind.TransactOpts, _triggerExpirationAndCleanup bool) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "setTriggerExpirationAndCleanup", _triggerExpirationAndCleanup)
}

// SetTriggerExpirationAndCleanup is a paid mutator transaction binding the contract method 0x67daec89.
//
// Solidity: function setTriggerExpirationAndCleanup(bool _triggerExpirationAndCleanup) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) SetTriggerExpirationAndCleanup(_triggerExpirationAndCleanup bool) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SetTriggerExpirationAndCleanup(&_FlareSystemsManager.TransactOpts, _triggerExpirationAndCleanup)
}

// SetTriggerExpirationAndCleanup is a paid mutator transaction binding the contract method 0x67daec89.
//
// Solidity: function setTriggerExpirationAndCleanup(bool _triggerExpirationAndCleanup) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) SetTriggerExpirationAndCleanup(_triggerExpirationAndCleanup bool) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SetTriggerExpirationAndCleanup(&_FlareSystemsManager.TransactOpts, _triggerExpirationAndCleanup)
}

// SetVoterRegistrationTriggerContract is a paid mutator transaction binding the contract method 0x24eb64de.
//
// Solidity: function setVoterRegistrationTriggerContract(address _contract) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) SetVoterRegistrationTriggerContract(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "setVoterRegistrationTriggerContract", _contract)
}

// SetVoterRegistrationTriggerContract is a paid mutator transaction binding the contract method 0x24eb64de.
//
// Solidity: function setVoterRegistrationTriggerContract(address _contract) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) SetVoterRegistrationTriggerContract(_contract common.Address) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SetVoterRegistrationTriggerContract(&_FlareSystemsManager.TransactOpts, _contract)
}

// SetVoterRegistrationTriggerContract is a paid mutator transaction binding the contract method 0x24eb64de.
//
// Solidity: function setVoterRegistrationTriggerContract(address _contract) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) SetVoterRegistrationTriggerContract(_contract common.Address) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SetVoterRegistrationTriggerContract(&_FlareSystemsManager.TransactOpts, _contract)
}

// SignNewSigningPolicy is a paid mutator transaction binding the contract method 0x6b4c7bd6.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) SignNewSigningPolicy(opts *bind.TransactOpts, _rewardEpochId *big.Int, _newSigningPolicyHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "signNewSigningPolicy", _rewardEpochId, _newSigningPolicyHash, _signature)
}

// SignNewSigningPolicy is a paid mutator transaction binding the contract method 0x6b4c7bd6.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) SignNewSigningPolicy(_rewardEpochId *big.Int, _newSigningPolicyHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SignNewSigningPolicy(&_FlareSystemsManager.TransactOpts, _rewardEpochId, _newSigningPolicyHash, _signature)
}

// SignNewSigningPolicy is a paid mutator transaction binding the contract method 0x6b4c7bd6.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) SignNewSigningPolicy(_rewardEpochId *big.Int, _newSigningPolicyHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SignNewSigningPolicy(&_FlareSystemsManager.TransactOpts, _rewardEpochId, _newSigningPolicyHash, _signature)
}

// SignRewards is a paid mutator transaction binding the contract method 0xab25ac5b.
//
// Solidity: function signRewards(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) SignRewards(opts *bind.TransactOpts, _rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "signRewards", _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash, _signature)
}

// SignRewards is a paid mutator transaction binding the contract method 0xab25ac5b.
//
// Solidity: function signRewards(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) SignRewards(_rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SignRewards(&_FlareSystemsManager.TransactOpts, _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash, _signature)
}

// SignRewards is a paid mutator transaction binding the contract method 0xab25ac5b.
//
// Solidity: function signRewards(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) SignRewards(_rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SignRewards(&_FlareSystemsManager.TransactOpts, _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash, _signature)
}

// SignUptimeVote is a paid mutator transaction binding the contract method 0xdc5a4225.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) SignUptimeVote(opts *bind.TransactOpts, _rewardEpochId *big.Int, _uptimeVoteHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "signUptimeVote", _rewardEpochId, _uptimeVoteHash, _signature)
}

// SignUptimeVote is a paid mutator transaction binding the contract method 0xdc5a4225.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) SignUptimeVote(_rewardEpochId *big.Int, _uptimeVoteHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SignUptimeVote(&_FlareSystemsManager.TransactOpts, _rewardEpochId, _uptimeVoteHash, _signature)
}

// SignUptimeVote is a paid mutator transaction binding the contract method 0xdc5a4225.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) SignUptimeVote(_rewardEpochId *big.Int, _uptimeVoteHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SignUptimeVote(&_FlareSystemsManager.TransactOpts, _rewardEpochId, _uptimeVoteHash, _signature)
}

// SubmitUptimeVote is a paid mutator transaction binding the contract method 0x9dd6850f.
//
// Solidity: function submitUptimeVote(uint24 _rewardEpochId, bytes20[] _nodeIds, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) SubmitUptimeVote(opts *bind.TransactOpts, _rewardEpochId *big.Int, _nodeIds [][20]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "submitUptimeVote", _rewardEpochId, _nodeIds, _signature)
}

// SubmitUptimeVote is a paid mutator transaction binding the contract method 0x9dd6850f.
//
// Solidity: function submitUptimeVote(uint24 _rewardEpochId, bytes20[] _nodeIds, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) SubmitUptimeVote(_rewardEpochId *big.Int, _nodeIds [][20]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SubmitUptimeVote(&_FlareSystemsManager.TransactOpts, _rewardEpochId, _nodeIds, _signature)
}

// SubmitUptimeVote is a paid mutator transaction binding the contract method 0x9dd6850f.
//
// Solidity: function submitUptimeVote(uint24 _rewardEpochId, bytes20[] _nodeIds, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) SubmitUptimeVote(_rewardEpochId *big.Int, _nodeIds [][20]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SubmitUptimeVote(&_FlareSystemsManager.TransactOpts, _rewardEpochId, _nodeIds, _signature)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SwitchToProductionMode(&_FlareSystemsManager.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.SwitchToProductionMode(&_FlareSystemsManager.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.UpdateContractAddresses(&_FlareSystemsManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.UpdateContractAddresses(&_FlareSystemsManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x8fbaf860.
//
// Solidity: function updateSettings((uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint24,uint16,uint32) _settings) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactor) UpdateSettings(opts *bind.TransactOpts, _settings FlareSystemsManagerSettings) (*types.Transaction, error) {
	return _FlareSystemsManager.contract.Transact(opts, "updateSettings", _settings)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x8fbaf860.
//
// Solidity: function updateSettings((uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint24,uint16,uint32) _settings) returns()
func (_FlareSystemsManager *FlareSystemsManagerSession) UpdateSettings(_settings FlareSystemsManagerSettings) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.UpdateSettings(&_FlareSystemsManager.TransactOpts, _settings)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x8fbaf860.
//
// Solidity: function updateSettings((uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint24,uint16,uint32) _settings) returns()
func (_FlareSystemsManager *FlareSystemsManagerTransactorSession) UpdateSettings(_settings FlareSystemsManagerSettings) (*types.Transaction, error) {
	return _FlareSystemsManager.Contract.UpdateSettings(&_FlareSystemsManager.TransactOpts, _settings)
}

// FlareSystemsManagerClosingExpiredRewardEpochFailedIterator is returned from FilterClosingExpiredRewardEpochFailed and is used to iterate over the raw logs and unpacked data for ClosingExpiredRewardEpochFailed events raised by the FlareSystemsManager contract.
type FlareSystemsManagerClosingExpiredRewardEpochFailedIterator struct {
	Event *FlareSystemsManagerClosingExpiredRewardEpochFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerClosingExpiredRewardEpochFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerClosingExpiredRewardEpochFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerClosingExpiredRewardEpochFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerClosingExpiredRewardEpochFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerClosingExpiredRewardEpochFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerClosingExpiredRewardEpochFailed represents a ClosingExpiredRewardEpochFailed event raised by the FlareSystemsManager contract.
type FlareSystemsManagerClosingExpiredRewardEpochFailed struct {
	RewardEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterClosingExpiredRewardEpochFailed is a free log retrieval operation binding the contract event 0xc0cded1f60001401da804c8a7703c1e8dc60521fca0f0f9853e2f1984b5410ba.
//
// Solidity: event ClosingExpiredRewardEpochFailed(uint24 rewardEpochId)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterClosingExpiredRewardEpochFailed(opts *bind.FilterOpts) (*FlareSystemsManagerClosingExpiredRewardEpochFailedIterator, error) {

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "ClosingExpiredRewardEpochFailed")
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerClosingExpiredRewardEpochFailedIterator{contract: _FlareSystemsManager.contract, event: "ClosingExpiredRewardEpochFailed", logs: logs, sub: sub}, nil
}

// WatchClosingExpiredRewardEpochFailed is a free log subscription operation binding the contract event 0xc0cded1f60001401da804c8a7703c1e8dc60521fca0f0f9853e2f1984b5410ba.
//
// Solidity: event ClosingExpiredRewardEpochFailed(uint24 rewardEpochId)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchClosingExpiredRewardEpochFailed(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerClosingExpiredRewardEpochFailed) (event.Subscription, error) {

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "ClosingExpiredRewardEpochFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerClosingExpiredRewardEpochFailed)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "ClosingExpiredRewardEpochFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClosingExpiredRewardEpochFailed is a log parse operation binding the contract event 0xc0cded1f60001401da804c8a7703c1e8dc60521fca0f0f9853e2f1984b5410ba.
//
// Solidity: event ClosingExpiredRewardEpochFailed(uint24 rewardEpochId)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseClosingExpiredRewardEpochFailed(log types.Log) (*FlareSystemsManagerClosingExpiredRewardEpochFailed, error) {
	event := new(FlareSystemsManagerClosingExpiredRewardEpochFailed)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "ClosingExpiredRewardEpochFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the FlareSystemsManager contract.
type FlareSystemsManagerGovernanceCallTimelockedIterator struct {
	Event *FlareSystemsManagerGovernanceCallTimelocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerGovernanceCallTimelocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerGovernanceCallTimelocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FlareSystemsManager contract.
type FlareSystemsManagerGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*FlareSystemsManagerGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerGovernanceCallTimelockedIterator{contract: _FlareSystemsManager.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerGovernanceCallTimelocked)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseGovernanceCallTimelocked(log types.Log) (*FlareSystemsManagerGovernanceCallTimelocked, error) {
	event := new(FlareSystemsManagerGovernanceCallTimelocked)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the FlareSystemsManager contract.
type FlareSystemsManagerGovernanceInitialisedIterator struct {
	Event *FlareSystemsManagerGovernanceInitialised // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerGovernanceInitialised)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerGovernanceInitialised)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerGovernanceInitialised represents a GovernanceInitialised event raised by the FlareSystemsManager contract.
type FlareSystemsManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*FlareSystemsManagerGovernanceInitialisedIterator, error) {

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerGovernanceInitialisedIterator{contract: _FlareSystemsManager.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerGovernanceInitialised)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceInitialised is a log parse operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseGovernanceInitialised(log types.Log) (*FlareSystemsManagerGovernanceInitialised, error) {
	event := new(FlareSystemsManagerGovernanceInitialised)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the FlareSystemsManager contract.
type FlareSystemsManagerGovernedProductionModeEnteredIterator struct {
	Event *FlareSystemsManagerGovernedProductionModeEntered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerGovernedProductionModeEntered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerGovernedProductionModeEntered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FlareSystemsManager contract.
type FlareSystemsManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*FlareSystemsManagerGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerGovernedProductionModeEnteredIterator{contract: _FlareSystemsManager.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerGovernedProductionModeEntered)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernedProductionModeEntered is a log parse operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseGovernedProductionModeEntered(log types.Log) (*FlareSystemsManagerGovernedProductionModeEntered, error) {
	event := new(FlareSystemsManagerGovernedProductionModeEntered)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerRandomAcquisitionStartedIterator is returned from FilterRandomAcquisitionStarted and is used to iterate over the raw logs and unpacked data for RandomAcquisitionStarted events raised by the FlareSystemsManager contract.
type FlareSystemsManagerRandomAcquisitionStartedIterator struct {
	Event *FlareSystemsManagerRandomAcquisitionStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerRandomAcquisitionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerRandomAcquisitionStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerRandomAcquisitionStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerRandomAcquisitionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerRandomAcquisitionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerRandomAcquisitionStarted represents a RandomAcquisitionStarted event raised by the FlareSystemsManager contract.
type FlareSystemsManagerRandomAcquisitionStarted struct {
	RewardEpochId *big.Int
	Timestamp     uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRandomAcquisitionStarted is a free log retrieval operation binding the contract event 0xf9991783e5e480e42d9a54d3f35f4321857f8f0ebeb3742d326dce28b1126708.
//
// Solidity: event RandomAcquisitionStarted(uint24 indexed rewardEpochId, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterRandomAcquisitionStarted(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*FlareSystemsManagerRandomAcquisitionStartedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "RandomAcquisitionStarted", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerRandomAcquisitionStartedIterator{contract: _FlareSystemsManager.contract, event: "RandomAcquisitionStarted", logs: logs, sub: sub}, nil
}

// WatchRandomAcquisitionStarted is a free log subscription operation binding the contract event 0xf9991783e5e480e42d9a54d3f35f4321857f8f0ebeb3742d326dce28b1126708.
//
// Solidity: event RandomAcquisitionStarted(uint24 indexed rewardEpochId, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchRandomAcquisitionStarted(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerRandomAcquisitionStarted, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "RandomAcquisitionStarted", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerRandomAcquisitionStarted)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "RandomAcquisitionStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRandomAcquisitionStarted is a log parse operation binding the contract event 0xf9991783e5e480e42d9a54d3f35f4321857f8f0ebeb3742d326dce28b1126708.
//
// Solidity: event RandomAcquisitionStarted(uint24 indexed rewardEpochId, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseRandomAcquisitionStarted(log types.Log) (*FlareSystemsManagerRandomAcquisitionStarted, error) {
	event := new(FlareSystemsManagerRandomAcquisitionStarted)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "RandomAcquisitionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerRewardEpochStartedIterator is returned from FilterRewardEpochStarted and is used to iterate over the raw logs and unpacked data for RewardEpochStarted events raised by the FlareSystemsManager contract.
type FlareSystemsManagerRewardEpochStartedIterator struct {
	Event *FlareSystemsManagerRewardEpochStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerRewardEpochStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerRewardEpochStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerRewardEpochStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerRewardEpochStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerRewardEpochStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerRewardEpochStarted represents a RewardEpochStarted event raised by the FlareSystemsManager contract.
type FlareSystemsManagerRewardEpochStarted struct {
	RewardEpochId      *big.Int
	StartVotingRoundId uint32
	Timestamp          uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewardEpochStarted is a free log retrieval operation binding the contract event 0x4abb62ab1e4c42a11b90e4e45b92af1274f74cc634b759518e8c99e000d8be6d.
//
// Solidity: event RewardEpochStarted(uint24 indexed rewardEpochId, uint32 startVotingRoundId, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterRewardEpochStarted(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*FlareSystemsManagerRewardEpochStartedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "RewardEpochStarted", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerRewardEpochStartedIterator{contract: _FlareSystemsManager.contract, event: "RewardEpochStarted", logs: logs, sub: sub}, nil
}

// WatchRewardEpochStarted is a free log subscription operation binding the contract event 0x4abb62ab1e4c42a11b90e4e45b92af1274f74cc634b759518e8c99e000d8be6d.
//
// Solidity: event RewardEpochStarted(uint24 indexed rewardEpochId, uint32 startVotingRoundId, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchRewardEpochStarted(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerRewardEpochStarted, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "RewardEpochStarted", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerRewardEpochStarted)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "RewardEpochStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardEpochStarted is a log parse operation binding the contract event 0x4abb62ab1e4c42a11b90e4e45b92af1274f74cc634b759518e8c99e000d8be6d.
//
// Solidity: event RewardEpochStarted(uint24 indexed rewardEpochId, uint32 startVotingRoundId, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseRewardEpochStarted(log types.Log) (*FlareSystemsManagerRewardEpochStarted, error) {
	event := new(FlareSystemsManagerRewardEpochStarted)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "RewardEpochStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerRewardsSignedIterator is returned from FilterRewardsSigned and is used to iterate over the raw logs and unpacked data for RewardsSigned events raised by the FlareSystemsManager contract.
type FlareSystemsManagerRewardsSignedIterator struct {
	Event *FlareSystemsManagerRewardsSigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerRewardsSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerRewardsSigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerRewardsSigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerRewardsSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerRewardsSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerRewardsSigned represents a RewardsSigned event raised by the FlareSystemsManager contract.
type FlareSystemsManagerRewardsSigned struct {
	RewardEpochId         *big.Int
	SigningPolicyAddress  common.Address
	Voter                 common.Address
	RewardsHash           [32]byte
	NoOfWeightBasedClaims *big.Int
	Timestamp             uint64
	ThresholdReached      bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRewardsSigned is a free log retrieval operation binding the contract event 0x80df246a47153e3604b0026fcf53ec85ea48178efd4f3cad24f9a1e1da2dd52a.
//
// Solidity: event RewardsSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 rewardsHash, uint256 noOfWeightBasedClaims, uint64 timestamp, bool thresholdReached)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterRewardsSigned(opts *bind.FilterOpts, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (*FlareSystemsManagerRewardsSignedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "RewardsSigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerRewardsSignedIterator{contract: _FlareSystemsManager.contract, event: "RewardsSigned", logs: logs, sub: sub}, nil
}

// WatchRewardsSigned is a free log subscription operation binding the contract event 0x80df246a47153e3604b0026fcf53ec85ea48178efd4f3cad24f9a1e1da2dd52a.
//
// Solidity: event RewardsSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 rewardsHash, uint256 noOfWeightBasedClaims, uint64 timestamp, bool thresholdReached)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchRewardsSigned(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerRewardsSigned, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "RewardsSigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerRewardsSigned)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "RewardsSigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsSigned is a log parse operation binding the contract event 0x80df246a47153e3604b0026fcf53ec85ea48178efd4f3cad24f9a1e1da2dd52a.
//
// Solidity: event RewardsSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 rewardsHash, uint256 noOfWeightBasedClaims, uint64 timestamp, bool thresholdReached)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseRewardsSigned(log types.Log) (*FlareSystemsManagerRewardsSigned, error) {
	event := new(FlareSystemsManagerRewardsSigned)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "RewardsSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerSettingCleanUpBlockNumberFailedIterator is returned from FilterSettingCleanUpBlockNumberFailed and is used to iterate over the raw logs and unpacked data for SettingCleanUpBlockNumberFailed events raised by the FlareSystemsManager contract.
type FlareSystemsManagerSettingCleanUpBlockNumberFailedIterator struct {
	Event *FlareSystemsManagerSettingCleanUpBlockNumberFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerSettingCleanUpBlockNumberFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerSettingCleanUpBlockNumberFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerSettingCleanUpBlockNumberFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerSettingCleanUpBlockNumberFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerSettingCleanUpBlockNumberFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerSettingCleanUpBlockNumberFailed represents a SettingCleanUpBlockNumberFailed event raised by the FlareSystemsManager contract.
type FlareSystemsManagerSettingCleanUpBlockNumberFailed struct {
	BlockNumber uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSettingCleanUpBlockNumberFailed is a free log retrieval operation binding the contract event 0xe9a7be2e41a6b0b36d253d56488c6844e611be2bffd8dd4b69b89a078f41fecc.
//
// Solidity: event SettingCleanUpBlockNumberFailed(uint64 blockNumber)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterSettingCleanUpBlockNumberFailed(opts *bind.FilterOpts) (*FlareSystemsManagerSettingCleanUpBlockNumberFailedIterator, error) {

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "SettingCleanUpBlockNumberFailed")
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerSettingCleanUpBlockNumberFailedIterator{contract: _FlareSystemsManager.contract, event: "SettingCleanUpBlockNumberFailed", logs: logs, sub: sub}, nil
}

// WatchSettingCleanUpBlockNumberFailed is a free log subscription operation binding the contract event 0xe9a7be2e41a6b0b36d253d56488c6844e611be2bffd8dd4b69b89a078f41fecc.
//
// Solidity: event SettingCleanUpBlockNumberFailed(uint64 blockNumber)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchSettingCleanUpBlockNumberFailed(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerSettingCleanUpBlockNumberFailed) (event.Subscription, error) {

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "SettingCleanUpBlockNumberFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerSettingCleanUpBlockNumberFailed)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "SettingCleanUpBlockNumberFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettingCleanUpBlockNumberFailed is a log parse operation binding the contract event 0xe9a7be2e41a6b0b36d253d56488c6844e611be2bffd8dd4b69b89a078f41fecc.
//
// Solidity: event SettingCleanUpBlockNumberFailed(uint64 blockNumber)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseSettingCleanUpBlockNumberFailed(log types.Log) (*FlareSystemsManagerSettingCleanUpBlockNumberFailed, error) {
	event := new(FlareSystemsManagerSettingCleanUpBlockNumberFailed)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "SettingCleanUpBlockNumberFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerSignUptimeVoteEnabledIterator is returned from FilterSignUptimeVoteEnabled and is used to iterate over the raw logs and unpacked data for SignUptimeVoteEnabled events raised by the FlareSystemsManager contract.
type FlareSystemsManagerSignUptimeVoteEnabledIterator struct {
	Event *FlareSystemsManagerSignUptimeVoteEnabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerSignUptimeVoteEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerSignUptimeVoteEnabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerSignUptimeVoteEnabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerSignUptimeVoteEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerSignUptimeVoteEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerSignUptimeVoteEnabled represents a SignUptimeVoteEnabled event raised by the FlareSystemsManager contract.
type FlareSystemsManagerSignUptimeVoteEnabled struct {
	RewardEpochId *big.Int
	Timestamp     uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSignUptimeVoteEnabled is a free log retrieval operation binding the contract event 0x235cef7d085c1e59545613282d239e56eb0cd056135aa46b8c658cf54a078561.
//
// Solidity: event SignUptimeVoteEnabled(uint24 indexed rewardEpochId, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterSignUptimeVoteEnabled(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*FlareSystemsManagerSignUptimeVoteEnabledIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "SignUptimeVoteEnabled", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerSignUptimeVoteEnabledIterator{contract: _FlareSystemsManager.contract, event: "SignUptimeVoteEnabled", logs: logs, sub: sub}, nil
}

// WatchSignUptimeVoteEnabled is a free log subscription operation binding the contract event 0x235cef7d085c1e59545613282d239e56eb0cd056135aa46b8c658cf54a078561.
//
// Solidity: event SignUptimeVoteEnabled(uint24 indexed rewardEpochId, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchSignUptimeVoteEnabled(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerSignUptimeVoteEnabled, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "SignUptimeVoteEnabled", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerSignUptimeVoteEnabled)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "SignUptimeVoteEnabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignUptimeVoteEnabled is a log parse operation binding the contract event 0x235cef7d085c1e59545613282d239e56eb0cd056135aa46b8c658cf54a078561.
//
// Solidity: event SignUptimeVoteEnabled(uint24 indexed rewardEpochId, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseSignUptimeVoteEnabled(log types.Log) (*FlareSystemsManagerSignUptimeVoteEnabled, error) {
	event := new(FlareSystemsManagerSignUptimeVoteEnabled)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "SignUptimeVoteEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerSigningPolicySignedIterator is returned from FilterSigningPolicySigned and is used to iterate over the raw logs and unpacked data for SigningPolicySigned events raised by the FlareSystemsManager contract.
type FlareSystemsManagerSigningPolicySignedIterator struct {
	Event *FlareSystemsManagerSigningPolicySigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerSigningPolicySignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerSigningPolicySigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerSigningPolicySigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerSigningPolicySignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerSigningPolicySignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerSigningPolicySigned represents a SigningPolicySigned event raised by the FlareSystemsManager contract.
type FlareSystemsManagerSigningPolicySigned struct {
	RewardEpochId        *big.Int
	SigningPolicyAddress common.Address
	Voter                common.Address
	Timestamp            uint64
	ThresholdReached     bool
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSigningPolicySigned is a free log retrieval operation binding the contract event 0x154b0214ae62d8a5548c1eac25fabd87c38b04932a217732e1022f3118da67f3.
//
// Solidity: event SigningPolicySigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, uint64 timestamp, bool thresholdReached)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterSigningPolicySigned(opts *bind.FilterOpts, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (*FlareSystemsManagerSigningPolicySignedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "SigningPolicySigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerSigningPolicySignedIterator{contract: _FlareSystemsManager.contract, event: "SigningPolicySigned", logs: logs, sub: sub}, nil
}

// WatchSigningPolicySigned is a free log subscription operation binding the contract event 0x154b0214ae62d8a5548c1eac25fabd87c38b04932a217732e1022f3118da67f3.
//
// Solidity: event SigningPolicySigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, uint64 timestamp, bool thresholdReached)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchSigningPolicySigned(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerSigningPolicySigned, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "SigningPolicySigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerSigningPolicySigned)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "SigningPolicySigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSigningPolicySigned is a log parse operation binding the contract event 0x154b0214ae62d8a5548c1eac25fabd87c38b04932a217732e1022f3118da67f3.
//
// Solidity: event SigningPolicySigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, uint64 timestamp, bool thresholdReached)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseSigningPolicySigned(log types.Log) (*FlareSystemsManagerSigningPolicySigned, error) {
	event := new(FlareSystemsManagerSigningPolicySigned)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "SigningPolicySigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the FlareSystemsManager contract.
type FlareSystemsManagerTimelockedGovernanceCallCanceledIterator struct {
	Event *FlareSystemsManagerTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerTimelockedGovernanceCallCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerTimelockedGovernanceCallCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FlareSystemsManager contract.
type FlareSystemsManagerTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*FlareSystemsManagerTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerTimelockedGovernanceCallCanceledIterator{contract: _FlareSystemsManager.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerTimelockedGovernanceCallCanceled)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimelockedGovernanceCallCanceled is a log parse operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*FlareSystemsManagerTimelockedGovernanceCallCanceled, error) {
	event := new(FlareSystemsManagerTimelockedGovernanceCallCanceled)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the FlareSystemsManager contract.
type FlareSystemsManagerTimelockedGovernanceCallExecutedIterator struct {
	Event *FlareSystemsManagerTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerTimelockedGovernanceCallExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerTimelockedGovernanceCallExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FlareSystemsManager contract.
type FlareSystemsManagerTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*FlareSystemsManagerTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerTimelockedGovernanceCallExecutedIterator{contract: _FlareSystemsManager.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerTimelockedGovernanceCallExecuted)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimelockedGovernanceCallExecuted is a log parse operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*FlareSystemsManagerTimelockedGovernanceCallExecuted, error) {
	event := new(FlareSystemsManagerTimelockedGovernanceCallExecuted)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerTriggeringVoterRegistrationFailedIterator is returned from FilterTriggeringVoterRegistrationFailed and is used to iterate over the raw logs and unpacked data for TriggeringVoterRegistrationFailed events raised by the FlareSystemsManager contract.
type FlareSystemsManagerTriggeringVoterRegistrationFailedIterator struct {
	Event *FlareSystemsManagerTriggeringVoterRegistrationFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerTriggeringVoterRegistrationFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerTriggeringVoterRegistrationFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerTriggeringVoterRegistrationFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerTriggeringVoterRegistrationFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerTriggeringVoterRegistrationFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerTriggeringVoterRegistrationFailed represents a TriggeringVoterRegistrationFailed event raised by the FlareSystemsManager contract.
type FlareSystemsManagerTriggeringVoterRegistrationFailed struct {
	RewardEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTriggeringVoterRegistrationFailed is a free log retrieval operation binding the contract event 0x449d255b9c487823db86822a857f218d40682abada12acee2483788dc2fa975a.
//
// Solidity: event TriggeringVoterRegistrationFailed(uint24 rewardEpochId)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterTriggeringVoterRegistrationFailed(opts *bind.FilterOpts) (*FlareSystemsManagerTriggeringVoterRegistrationFailedIterator, error) {

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "TriggeringVoterRegistrationFailed")
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerTriggeringVoterRegistrationFailedIterator{contract: _FlareSystemsManager.contract, event: "TriggeringVoterRegistrationFailed", logs: logs, sub: sub}, nil
}

// WatchTriggeringVoterRegistrationFailed is a free log subscription operation binding the contract event 0x449d255b9c487823db86822a857f218d40682abada12acee2483788dc2fa975a.
//
// Solidity: event TriggeringVoterRegistrationFailed(uint24 rewardEpochId)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchTriggeringVoterRegistrationFailed(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerTriggeringVoterRegistrationFailed) (event.Subscription, error) {

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "TriggeringVoterRegistrationFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerTriggeringVoterRegistrationFailed)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "TriggeringVoterRegistrationFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTriggeringVoterRegistrationFailed is a log parse operation binding the contract event 0x449d255b9c487823db86822a857f218d40682abada12acee2483788dc2fa975a.
//
// Solidity: event TriggeringVoterRegistrationFailed(uint24 rewardEpochId)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseTriggeringVoterRegistrationFailed(log types.Log) (*FlareSystemsManagerTriggeringVoterRegistrationFailed, error) {
	event := new(FlareSystemsManagerTriggeringVoterRegistrationFailed)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "TriggeringVoterRegistrationFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerUptimeVoteSignedIterator is returned from FilterUptimeVoteSigned and is used to iterate over the raw logs and unpacked data for UptimeVoteSigned events raised by the FlareSystemsManager contract.
type FlareSystemsManagerUptimeVoteSignedIterator struct {
	Event *FlareSystemsManagerUptimeVoteSigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerUptimeVoteSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerUptimeVoteSigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerUptimeVoteSigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerUptimeVoteSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerUptimeVoteSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerUptimeVoteSigned represents a UptimeVoteSigned event raised by the FlareSystemsManager contract.
type FlareSystemsManagerUptimeVoteSigned struct {
	RewardEpochId        *big.Int
	SigningPolicyAddress common.Address
	Voter                common.Address
	UptimeVoteHash       [32]byte
	Timestamp            uint64
	ThresholdReached     bool
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUptimeVoteSigned is a free log retrieval operation binding the contract event 0x5506337d1266599f8b64675a1c8321701657ca2f2f70be0e0c58302b6c22e797.
//
// Solidity: event UptimeVoteSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 uptimeVoteHash, uint64 timestamp, bool thresholdReached)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterUptimeVoteSigned(opts *bind.FilterOpts, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (*FlareSystemsManagerUptimeVoteSignedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "UptimeVoteSigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerUptimeVoteSignedIterator{contract: _FlareSystemsManager.contract, event: "UptimeVoteSigned", logs: logs, sub: sub}, nil
}

// WatchUptimeVoteSigned is a free log subscription operation binding the contract event 0x5506337d1266599f8b64675a1c8321701657ca2f2f70be0e0c58302b6c22e797.
//
// Solidity: event UptimeVoteSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 uptimeVoteHash, uint64 timestamp, bool thresholdReached)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchUptimeVoteSigned(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerUptimeVoteSigned, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "UptimeVoteSigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerUptimeVoteSigned)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "UptimeVoteSigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUptimeVoteSigned is a log parse operation binding the contract event 0x5506337d1266599f8b64675a1c8321701657ca2f2f70be0e0c58302b6c22e797.
//
// Solidity: event UptimeVoteSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 uptimeVoteHash, uint64 timestamp, bool thresholdReached)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseUptimeVoteSigned(log types.Log) (*FlareSystemsManagerUptimeVoteSigned, error) {
	event := new(FlareSystemsManagerUptimeVoteSigned)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "UptimeVoteSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerUptimeVoteSubmittedIterator is returned from FilterUptimeVoteSubmitted and is used to iterate over the raw logs and unpacked data for UptimeVoteSubmitted events raised by the FlareSystemsManager contract.
type FlareSystemsManagerUptimeVoteSubmittedIterator struct {
	Event *FlareSystemsManagerUptimeVoteSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerUptimeVoteSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerUptimeVoteSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerUptimeVoteSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerUptimeVoteSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerUptimeVoteSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerUptimeVoteSubmitted represents a UptimeVoteSubmitted event raised by the FlareSystemsManager contract.
type FlareSystemsManagerUptimeVoteSubmitted struct {
	RewardEpochId        *big.Int
	SigningPolicyAddress common.Address
	Voter                common.Address
	NodeIds              [][20]byte
	Timestamp            uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUptimeVoteSubmitted is a free log retrieval operation binding the contract event 0xed370d61eb315e1d46d979894585530b99f94dab64c0d40366685aebe39e3db0.
//
// Solidity: event UptimeVoteSubmitted(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes20[] nodeIds, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterUptimeVoteSubmitted(opts *bind.FilterOpts, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (*FlareSystemsManagerUptimeVoteSubmittedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "UptimeVoteSubmitted", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerUptimeVoteSubmittedIterator{contract: _FlareSystemsManager.contract, event: "UptimeVoteSubmitted", logs: logs, sub: sub}, nil
}

// WatchUptimeVoteSubmitted is a free log subscription operation binding the contract event 0xed370d61eb315e1d46d979894585530b99f94dab64c0d40366685aebe39e3db0.
//
// Solidity: event UptimeVoteSubmitted(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes20[] nodeIds, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchUptimeVoteSubmitted(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerUptimeVoteSubmitted, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "UptimeVoteSubmitted", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerUptimeVoteSubmitted)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "UptimeVoteSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUptimeVoteSubmitted is a log parse operation binding the contract event 0xed370d61eb315e1d46d979894585530b99f94dab64c0d40366685aebe39e3db0.
//
// Solidity: event UptimeVoteSubmitted(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes20[] nodeIds, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseUptimeVoteSubmitted(log types.Log) (*FlareSystemsManagerUptimeVoteSubmitted, error) {
	event := new(FlareSystemsManagerUptimeVoteSubmitted)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "UptimeVoteSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemsManagerVotePowerBlockSelectedIterator is returned from FilterVotePowerBlockSelected and is used to iterate over the raw logs and unpacked data for VotePowerBlockSelected events raised by the FlareSystemsManager contract.
type FlareSystemsManagerVotePowerBlockSelectedIterator struct {
	Event *FlareSystemsManagerVotePowerBlockSelected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlareSystemsManagerVotePowerBlockSelectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemsManagerVotePowerBlockSelected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlareSystemsManagerVotePowerBlockSelected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlareSystemsManagerVotePowerBlockSelectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemsManagerVotePowerBlockSelectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemsManagerVotePowerBlockSelected represents a VotePowerBlockSelected event raised by the FlareSystemsManager contract.
type FlareSystemsManagerVotePowerBlockSelected struct {
	RewardEpochId  *big.Int
	VotePowerBlock uint64
	Timestamp      uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVotePowerBlockSelected is a free log retrieval operation binding the contract event 0xf21722dbe044a7cea0f6d81c871cae750971e36c9dd10999e46f2b46f26ac7ff.
//
// Solidity: event VotePowerBlockSelected(uint24 indexed rewardEpochId, uint64 votePowerBlock, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) FilterVotePowerBlockSelected(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*FlareSystemsManagerVotePowerBlockSelectedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.FilterLogs(opts, "VotePowerBlockSelected", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemsManagerVotePowerBlockSelectedIterator{contract: _FlareSystemsManager.contract, event: "VotePowerBlockSelected", logs: logs, sub: sub}, nil
}

// WatchVotePowerBlockSelected is a free log subscription operation binding the contract event 0xf21722dbe044a7cea0f6d81c871cae750971e36c9dd10999e46f2b46f26ac7ff.
//
// Solidity: event VotePowerBlockSelected(uint24 indexed rewardEpochId, uint64 votePowerBlock, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) WatchVotePowerBlockSelected(opts *bind.WatchOpts, sink chan<- *FlareSystemsManagerVotePowerBlockSelected, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemsManager.contract.WatchLogs(opts, "VotePowerBlockSelected", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemsManagerVotePowerBlockSelected)
				if err := _FlareSystemsManager.contract.UnpackLog(event, "VotePowerBlockSelected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVotePowerBlockSelected is a log parse operation binding the contract event 0xf21722dbe044a7cea0f6d81c871cae750971e36c9dd10999e46f2b46f26ac7ff.
//
// Solidity: event VotePowerBlockSelected(uint24 indexed rewardEpochId, uint64 votePowerBlock, uint64 timestamp)
func (_FlareSystemsManager *FlareSystemsManagerFilterer) ParseVotePowerBlockSelected(log types.Log) (*FlareSystemsManagerVotePowerBlockSelected, error) {
	event := new(FlareSystemsManagerVotePowerBlockSelected)
	if err := _FlareSystemsManager.contract.UnpackLog(event, "VotePowerBlockSelected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
