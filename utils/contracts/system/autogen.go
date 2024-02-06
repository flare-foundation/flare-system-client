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

// FlareSystemManagerInitialSettings is an auto generated low-level Go binding around an user-defined struct.
type FlareSystemManagerInitialSettings struct {
	InitialRandomVotePowerBlockSelectionSize uint16
	InitialRewardEpochId                     *big.Int
	InitialRewardEpochThreshold              uint16
}

// FlareSystemManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type FlareSystemManagerSettings struct {
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

// IFlareSystemManagerSignature is an auto generated low-level Go binding around an user-defined struct.
type IFlareSystemManagerSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// FlareSystemManagerMetaData contains all meta data concerning the FlareSystemManager contract.
var FlareSystemManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_flareDaemon\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"randomAcquisitionMaxDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"randomAcquisitionMaxDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newSigningPolicyInitializationStartSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"newSigningPolicyMinNumberOfVotingRoundsDelay\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"voterRegistrationMinDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"voterRegistrationMinDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"submitUptimeVoteMinDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"submitUptimeVoteMinDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"signingPolicyThresholdPPM\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"signingPolicyMinNumberOfVoters\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"rewardExpiryOffsetSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structFlareSystemManager.Settings\",\"name\":\"_settings\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"_firstVotingRoundStartTs\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_votingEpochDurationSeconds\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_firstRewardEpochStartVotingRoundId\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_rewardEpochDurationInVotingEpochs\",\"type\":\"uint16\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"initialRandomVotePowerBlockSelectionSize\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"initialRewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"initialRewardEpochThreshold\",\"type\":\"uint16\"}],\"internalType\":\"structFlareSystemManager.InitialSettings\",\"name\":\"_initialSettings\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"ClosingExpiredRewardEpochFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"RandomAcquisitionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startVotingRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"RewardEpochStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rewardsHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noOfWeightBasedClaims\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"RewardsSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"name\":\"SettingCleanUpBlockNumberFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"SigningPolicySigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"SingUptimeVoteEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"TriggeringVoterRegistrationFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"uptimeVoteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"UptimeVoteSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes20[]\",\"name\":\"nodeIds\",\"type\":\"bytes20[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"UptimeVoteSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"votePowerBlock\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"VotePowerBlockSelected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_signingPolicyThresholdPPM\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"_signingPolicyMinNumberOfVoters\",\"type\":\"uint16\"}],\"name\":\"changeSigningPolicySettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanupBlockNumberManager\",\"outputs\":[{\"internalType\":\"contractIICleanupBlockNumberManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewardEpochExpectedEndTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daemonize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstRewardEpochStartTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstVotingRoundStartTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareDaemon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRewardEpochId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getRandomAcquisitionInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_randomAcquisitionStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_randomAcquisitionStartBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_randomAcquisitionEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_randomAcquisitionEndBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getRewardEpochStartInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_rewardEpochStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardEpochStartBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardEpochSwitchoverTriggerContracts\",\"outputs\":[{\"internalType\":\"contractIIRewardEpochSwitchoverTrigger[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getRewardsSignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_rewardsSignStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardsSignStartBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardsSignEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardsSignEndBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getSeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getSigningPolicySignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignStartBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignEndBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getStartVotingRoundId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getUptimeVoteSignStartInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSignStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSignStartBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getVotePowerBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_votePowerBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getVoterRegistrationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_votePowerBlock\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterRewardsSignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_rewardsSignTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardsSignBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterSigningPolicySignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterUptimeVoteSignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSignTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSignBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterUptimeVoteSubmitInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSubmitTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSubmitBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialRandomVotePowerBlockSelectionSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isVoterRegistrationEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInitializedVotingRoundId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newSigningPolicyInitializationStartSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newSigningPolicyMinNumberOfVotingRoundsDelay\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"noOfWeightBasedClaims\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomAcquisitionMaxDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomAcquisitionMaxDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relay\",\"outputs\":[{\"internalType\":\"contractIIRelay\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardEpochIdToExpireNext\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardExpiryOffsetSeconds\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardsHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIIRewardEpochSwitchoverTrigger[]\",\"name\":\"_contracts\",\"type\":\"address[]\"}],\"name\":\"setRewardEpochSwitchoverTriggerContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_noOfWeightBasedClaims\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_rewardsHash\",\"type\":\"bytes32\"}],\"name\":\"setRewardsData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_submit3Aligned\",\"type\":\"bool\"}],\"name\":\"setSubmit3Aligned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_triggerExpirationAndCleanup\",\"type\":\"bool\"}],\"name\":\"setTriggerExpirationAndCleanup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIIVoterRegistrationTrigger\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setVoterRegistrationTriggerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"_newSigningPolicyHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signNewSigningPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_noOfWeightBasedClaims\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_rewardsHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"_uptimeVoteHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signUptimeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingPolicyMinNumberOfVoters\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingPolicyThresholdPPM\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submission\",\"outputs\":[{\"internalType\":\"contractIISubmission\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submit3Aligned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes20[]\",\"name\":\"_nodeIds\",\"type\":\"bytes20[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"submitUptimeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitUptimeVoteMinDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitUptimeVoteMinDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToFallbackMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerExpirationAndCleanup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"randomAcquisitionMaxDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"randomAcquisitionMaxDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newSigningPolicyInitializationStartSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"newSigningPolicyMinNumberOfVotingRoundsDelay\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"voterRegistrationMinDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"voterRegistrationMinDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"submitUptimeVoteMinDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"submitUptimeVoteMinDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"signingPolicyThresholdPPM\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"signingPolicyMinNumberOfVoters\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"rewardExpiryOffsetSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structFlareSystemManager.Settings\",\"name\":\"_settings\",\"type\":\"tuple\"}],\"name\":\"updateSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uptimeVoteHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistrationMinDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistrationMinDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistrationTriggerContract\",\"outputs\":[{\"internalType\":\"contractIIVoterRegistrationTrigger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistry\",\"outputs\":[{\"internalType\":\"contractIIVoterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FlareSystemManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use FlareSystemManagerMetaData.ABI instead.
var FlareSystemManagerABI = FlareSystemManagerMetaData.ABI

// FlareSystemManager is an auto generated Go binding around an Ethereum contract.
type FlareSystemManager struct {
	FlareSystemManagerCaller     // Read-only binding to the contract
	FlareSystemManagerTransactor // Write-only binding to the contract
	FlareSystemManagerFilterer   // Log filterer for contract events
}

// FlareSystemManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlareSystemManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlareSystemManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlareSystemManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlareSystemManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlareSystemManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlareSystemManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlareSystemManagerSession struct {
	Contract     *FlareSystemManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FlareSystemManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlareSystemManagerCallerSession struct {
	Contract *FlareSystemManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// FlareSystemManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlareSystemManagerTransactorSession struct {
	Contract     *FlareSystemManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// FlareSystemManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlareSystemManagerRaw struct {
	Contract *FlareSystemManager // Generic contract binding to access the raw methods on
}

// FlareSystemManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlareSystemManagerCallerRaw struct {
	Contract *FlareSystemManagerCaller // Generic read-only contract binding to access the raw methods on
}

// FlareSystemManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlareSystemManagerTransactorRaw struct {
	Contract *FlareSystemManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlareSystemManager creates a new instance of FlareSystemManager, bound to a specific deployed contract.
func NewFlareSystemManager(address common.Address, backend bind.ContractBackend) (*FlareSystemManager, error) {
	contract, err := bindFlareSystemManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlareSystemManager{FlareSystemManagerCaller: FlareSystemManagerCaller{contract: contract}, FlareSystemManagerTransactor: FlareSystemManagerTransactor{contract: contract}, FlareSystemManagerFilterer: FlareSystemManagerFilterer{contract: contract}}, nil
}

// NewFlareSystemManagerCaller creates a new read-only instance of FlareSystemManager, bound to a specific deployed contract.
func NewFlareSystemManagerCaller(address common.Address, caller bind.ContractCaller) (*FlareSystemManagerCaller, error) {
	contract, err := bindFlareSystemManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerCaller{contract: contract}, nil
}

// NewFlareSystemManagerTransactor creates a new write-only instance of FlareSystemManager, bound to a specific deployed contract.
func NewFlareSystemManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FlareSystemManagerTransactor, error) {
	contract, err := bindFlareSystemManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerTransactor{contract: contract}, nil
}

// NewFlareSystemManagerFilterer creates a new log filterer instance of FlareSystemManager, bound to a specific deployed contract.
func NewFlareSystemManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FlareSystemManagerFilterer, error) {
	contract, err := bindFlareSystemManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerFilterer{contract: contract}, nil
}

// bindFlareSystemManager binds a generic wrapper to an already deployed contract.
func bindFlareSystemManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FlareSystemManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlareSystemManager *FlareSystemManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlareSystemManager.Contract.FlareSystemManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlareSystemManager *FlareSystemManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.FlareSystemManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlareSystemManager *FlareSystemManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.FlareSystemManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlareSystemManager *FlareSystemManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlareSystemManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlareSystemManager *FlareSystemManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlareSystemManager *FlareSystemManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.contract.Transact(opts, method, params...)
}

// CleanupBlockNumberManager is a free data retrieval call binding the contract method 0x4eac870f.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCaller) CleanupBlockNumberManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "cleanupBlockNumberManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CleanupBlockNumberManager is a free data retrieval call binding the contract method 0x4eac870f.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (_FlareSystemManager *FlareSystemManagerSession) CleanupBlockNumberManager() (common.Address, error) {
	return _FlareSystemManager.Contract.CleanupBlockNumberManager(&_FlareSystemManager.CallOpts)
}

// CleanupBlockNumberManager is a free data retrieval call binding the contract method 0x4eac870f.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCallerSession) CleanupBlockNumberManager() (common.Address, error) {
	return _FlareSystemManager.Contract.CleanupBlockNumberManager(&_FlareSystemManager.CallOpts)
}

// CurrentRewardEpochExpectedEndTs is a free data retrieval call binding the contract method 0xed54fd63.
//
// Solidity: function currentRewardEpochExpectedEndTs() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) CurrentRewardEpochExpectedEndTs(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "currentRewardEpochExpectedEndTs")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrentRewardEpochExpectedEndTs is a free data retrieval call binding the contract method 0xed54fd63.
//
// Solidity: function currentRewardEpochExpectedEndTs() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) CurrentRewardEpochExpectedEndTs() (uint64, error) {
	return _FlareSystemManager.Contract.CurrentRewardEpochExpectedEndTs(&_FlareSystemManager.CallOpts)
}

// CurrentRewardEpochExpectedEndTs is a free data retrieval call binding the contract method 0xed54fd63.
//
// Solidity: function currentRewardEpochExpectedEndTs() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) CurrentRewardEpochExpectedEndTs() (uint64, error) {
	return _FlareSystemManager.Contract.CurrentRewardEpochExpectedEndTs(&_FlareSystemManager.CallOpts)
}

// FirstRewardEpochStartTs is a free data retrieval call binding the contract method 0x79e047ed.
//
// Solidity: function firstRewardEpochStartTs() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) FirstRewardEpochStartTs(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "firstRewardEpochStartTs")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstRewardEpochStartTs is a free data retrieval call binding the contract method 0x79e047ed.
//
// Solidity: function firstRewardEpochStartTs() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) FirstRewardEpochStartTs() (uint64, error) {
	return _FlareSystemManager.Contract.FirstRewardEpochStartTs(&_FlareSystemManager.CallOpts)
}

// FirstRewardEpochStartTs is a free data retrieval call binding the contract method 0x79e047ed.
//
// Solidity: function firstRewardEpochStartTs() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) FirstRewardEpochStartTs() (uint64, error) {
	return _FlareSystemManager.Contract.FirstRewardEpochStartTs(&_FlareSystemManager.CallOpts)
}

// FirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) FirstVotingRoundStartTs(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "firstVotingRoundStartTs")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) FirstVotingRoundStartTs() (uint64, error) {
	return _FlareSystemManager.Contract.FirstVotingRoundStartTs(&_FlareSystemManager.CallOpts)
}

// FirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) FirstVotingRoundStartTs() (uint64, error) {
	return _FlareSystemManager.Contract.FirstVotingRoundStartTs(&_FlareSystemManager.CallOpts)
}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCaller) FlareDaemon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "flareDaemon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_FlareSystemManager *FlareSystemManagerSession) FlareDaemon() (common.Address, error) {
	return _FlareSystemManager.Contract.FlareDaemon(&_FlareSystemManager.CallOpts)
}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCallerSession) FlareDaemon() (common.Address, error) {
	return _FlareSystemManager.Contract.FlareDaemon(&_FlareSystemManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FlareSystemManager *FlareSystemManagerCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FlareSystemManager *FlareSystemManagerSession) GetAddressUpdater() (common.Address, error) {
	return _FlareSystemManager.Contract.GetAddressUpdater(&_FlareSystemManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetAddressUpdater() (common.Address, error) {
	return _FlareSystemManager.Contract.GetAddressUpdater(&_FlareSystemManager.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FlareSystemManager *FlareSystemManagerCaller) GetContractName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getContractName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FlareSystemManager *FlareSystemManagerSession) GetContractName() (string, error) {
	return _FlareSystemManager.Contract.GetContractName(&_FlareSystemManager.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetContractName() (string, error) {
	return _FlareSystemManager.Contract.GetContractName(&_FlareSystemManager.CallOpts)
}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (_FlareSystemManager *FlareSystemManagerCaller) GetCurrentRewardEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getCurrentRewardEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (_FlareSystemManager *FlareSystemManagerSession) GetCurrentRewardEpochId() (*big.Int, error) {
	return _FlareSystemManager.Contract.GetCurrentRewardEpochId(&_FlareSystemManager.CallOpts)
}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetCurrentRewardEpochId() (*big.Int, error) {
	return _FlareSystemManager.Contract.GetCurrentRewardEpochId(&_FlareSystemManager.CallOpts)
}

// GetRandomAcquisitionInfo is a free data retrieval call binding the contract method 0x8f8f9f3a.
//
// Solidity: function getRandomAcquisitionInfo(uint24 _rewardEpochId) view returns(uint64 _randomAcquisitionStartTs, uint64 _randomAcquisitionStartBlock, uint64 _randomAcquisitionEndTs, uint64 _randomAcquisitionEndBlock)
func (_FlareSystemManager *FlareSystemManagerCaller) GetRandomAcquisitionInfo(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	RandomAcquisitionStartTs    uint64
	RandomAcquisitionStartBlock uint64
	RandomAcquisitionEndTs      uint64
	RandomAcquisitionEndBlock   uint64
}, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getRandomAcquisitionInfo", _rewardEpochId)

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
func (_FlareSystemManager *FlareSystemManagerSession) GetRandomAcquisitionInfo(_rewardEpochId *big.Int) (struct {
	RandomAcquisitionStartTs    uint64
	RandomAcquisitionStartBlock uint64
	RandomAcquisitionEndTs      uint64
	RandomAcquisitionEndBlock   uint64
}, error) {
	return _FlareSystemManager.Contract.GetRandomAcquisitionInfo(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetRandomAcquisitionInfo is a free data retrieval call binding the contract method 0x8f8f9f3a.
//
// Solidity: function getRandomAcquisitionInfo(uint24 _rewardEpochId) view returns(uint64 _randomAcquisitionStartTs, uint64 _randomAcquisitionStartBlock, uint64 _randomAcquisitionEndTs, uint64 _randomAcquisitionEndBlock)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetRandomAcquisitionInfo(_rewardEpochId *big.Int) (struct {
	RandomAcquisitionStartTs    uint64
	RandomAcquisitionStartBlock uint64
	RandomAcquisitionEndTs      uint64
	RandomAcquisitionEndBlock   uint64
}, error) {
	return _FlareSystemManager.Contract.GetRandomAcquisitionInfo(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetRewardEpochStartInfo is a free data retrieval call binding the contract method 0x00ddae53.
//
// Solidity: function getRewardEpochStartInfo(uint24 _rewardEpochId) view returns(uint64 _rewardEpochStartTs, uint64 _rewardEpochStartBlock)
func (_FlareSystemManager *FlareSystemManagerCaller) GetRewardEpochStartInfo(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	RewardEpochStartTs    uint64
	RewardEpochStartBlock uint64
}, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getRewardEpochStartInfo", _rewardEpochId)

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
func (_FlareSystemManager *FlareSystemManagerSession) GetRewardEpochStartInfo(_rewardEpochId *big.Int) (struct {
	RewardEpochStartTs    uint64
	RewardEpochStartBlock uint64
}, error) {
	return _FlareSystemManager.Contract.GetRewardEpochStartInfo(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetRewardEpochStartInfo is a free data retrieval call binding the contract method 0x00ddae53.
//
// Solidity: function getRewardEpochStartInfo(uint24 _rewardEpochId) view returns(uint64 _rewardEpochStartTs, uint64 _rewardEpochStartBlock)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetRewardEpochStartInfo(_rewardEpochId *big.Int) (struct {
	RewardEpochStartTs    uint64
	RewardEpochStartBlock uint64
}, error) {
	return _FlareSystemManager.Contract.GetRewardEpochStartInfo(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetRewardEpochSwitchoverTriggerContracts is a free data retrieval call binding the contract method 0x46831531.
//
// Solidity: function getRewardEpochSwitchoverTriggerContracts() view returns(address[])
func (_FlareSystemManager *FlareSystemManagerCaller) GetRewardEpochSwitchoverTriggerContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getRewardEpochSwitchoverTriggerContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRewardEpochSwitchoverTriggerContracts is a free data retrieval call binding the contract method 0x46831531.
//
// Solidity: function getRewardEpochSwitchoverTriggerContracts() view returns(address[])
func (_FlareSystemManager *FlareSystemManagerSession) GetRewardEpochSwitchoverTriggerContracts() ([]common.Address, error) {
	return _FlareSystemManager.Contract.GetRewardEpochSwitchoverTriggerContracts(&_FlareSystemManager.CallOpts)
}

// GetRewardEpochSwitchoverTriggerContracts is a free data retrieval call binding the contract method 0x46831531.
//
// Solidity: function getRewardEpochSwitchoverTriggerContracts() view returns(address[])
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetRewardEpochSwitchoverTriggerContracts() ([]common.Address, error) {
	return _FlareSystemManager.Contract.GetRewardEpochSwitchoverTriggerContracts(&_FlareSystemManager.CallOpts)
}

// GetRewardsSignInfo is a free data retrieval call binding the contract method 0xb6c25af0.
//
// Solidity: function getRewardsSignInfo(uint24 _rewardEpochId) view returns(uint64 _rewardsSignStartTs, uint64 _rewardsSignStartBlock, uint64 _rewardsSignEndTs, uint64 _rewardsSignEndBlock)
func (_FlareSystemManager *FlareSystemManagerCaller) GetRewardsSignInfo(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	RewardsSignStartTs    uint64
	RewardsSignStartBlock uint64
	RewardsSignEndTs      uint64
	RewardsSignEndBlock   uint64
}, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getRewardsSignInfo", _rewardEpochId)

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
func (_FlareSystemManager *FlareSystemManagerSession) GetRewardsSignInfo(_rewardEpochId *big.Int) (struct {
	RewardsSignStartTs    uint64
	RewardsSignStartBlock uint64
	RewardsSignEndTs      uint64
	RewardsSignEndBlock   uint64
}, error) {
	return _FlareSystemManager.Contract.GetRewardsSignInfo(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetRewardsSignInfo is a free data retrieval call binding the contract method 0xb6c25af0.
//
// Solidity: function getRewardsSignInfo(uint24 _rewardEpochId) view returns(uint64 _rewardsSignStartTs, uint64 _rewardsSignStartBlock, uint64 _rewardsSignEndTs, uint64 _rewardsSignEndBlock)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetRewardsSignInfo(_rewardEpochId *big.Int) (struct {
	RewardsSignStartTs    uint64
	RewardsSignStartBlock uint64
	RewardsSignEndTs      uint64
	RewardsSignEndBlock   uint64
}, error) {
	return _FlareSystemManager.Contract.GetRewardsSignInfo(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetSeed is a free data retrieval call binding the contract method 0xe0d4ea37.
//
// Solidity: function getSeed(uint256 _rewardEpochId) view returns(uint256)
func (_FlareSystemManager *FlareSystemManagerCaller) GetSeed(opts *bind.CallOpts, _rewardEpochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getSeed", _rewardEpochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSeed is a free data retrieval call binding the contract method 0xe0d4ea37.
//
// Solidity: function getSeed(uint256 _rewardEpochId) view returns(uint256)
func (_FlareSystemManager *FlareSystemManagerSession) GetSeed(_rewardEpochId *big.Int) (*big.Int, error) {
	return _FlareSystemManager.Contract.GetSeed(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetSeed is a free data retrieval call binding the contract method 0xe0d4ea37.
//
// Solidity: function getSeed(uint256 _rewardEpochId) view returns(uint256)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetSeed(_rewardEpochId *big.Int) (*big.Int, error) {
	return _FlareSystemManager.Contract.GetSeed(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetSigningPolicySignInfo is a free data retrieval call binding the contract method 0xd2e9ad71.
//
// Solidity: function getSigningPolicySignInfo(uint24 _rewardEpochId) view returns(uint64 _signingPolicySignStartTs, uint64 _signingPolicySignStartBlock, uint64 _signingPolicySignEndTs, uint64 _signingPolicySignEndBlock)
func (_FlareSystemManager *FlareSystemManagerCaller) GetSigningPolicySignInfo(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	SigningPolicySignStartTs    uint64
	SigningPolicySignStartBlock uint64
	SigningPolicySignEndTs      uint64
	SigningPolicySignEndBlock   uint64
}, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getSigningPolicySignInfo", _rewardEpochId)

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
func (_FlareSystemManager *FlareSystemManagerSession) GetSigningPolicySignInfo(_rewardEpochId *big.Int) (struct {
	SigningPolicySignStartTs    uint64
	SigningPolicySignStartBlock uint64
	SigningPolicySignEndTs      uint64
	SigningPolicySignEndBlock   uint64
}, error) {
	return _FlareSystemManager.Contract.GetSigningPolicySignInfo(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetSigningPolicySignInfo is a free data retrieval call binding the contract method 0xd2e9ad71.
//
// Solidity: function getSigningPolicySignInfo(uint24 _rewardEpochId) view returns(uint64 _signingPolicySignStartTs, uint64 _signingPolicySignStartBlock, uint64 _signingPolicySignEndTs, uint64 _signingPolicySignEndBlock)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetSigningPolicySignInfo(_rewardEpochId *big.Int) (struct {
	SigningPolicySignStartTs    uint64
	SigningPolicySignStartBlock uint64
	SigningPolicySignEndTs      uint64
	SigningPolicySignEndBlock   uint64
}, error) {
	return _FlareSystemManager.Contract.GetSigningPolicySignInfo(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetStartVotingRoundId is a free data retrieval call binding the contract method 0x75d2187a.
//
// Solidity: function getStartVotingRoundId(uint256 _rewardEpochId) view returns(uint32)
func (_FlareSystemManager *FlareSystemManagerCaller) GetStartVotingRoundId(opts *bind.CallOpts, _rewardEpochId *big.Int) (uint32, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getStartVotingRoundId", _rewardEpochId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetStartVotingRoundId is a free data retrieval call binding the contract method 0x75d2187a.
//
// Solidity: function getStartVotingRoundId(uint256 _rewardEpochId) view returns(uint32)
func (_FlareSystemManager *FlareSystemManagerSession) GetStartVotingRoundId(_rewardEpochId *big.Int) (uint32, error) {
	return _FlareSystemManager.Contract.GetStartVotingRoundId(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetStartVotingRoundId is a free data retrieval call binding the contract method 0x75d2187a.
//
// Solidity: function getStartVotingRoundId(uint256 _rewardEpochId) view returns(uint32)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetStartVotingRoundId(_rewardEpochId *big.Int) (uint32, error) {
	return _FlareSystemManager.Contract.GetStartVotingRoundId(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetThreshold is a free data retrieval call binding the contract method 0x4615d5e9.
//
// Solidity: function getThreshold(uint256 _rewardEpochId) view returns(uint16)
func (_FlareSystemManager *FlareSystemManagerCaller) GetThreshold(opts *bind.CallOpts, _rewardEpochId *big.Int) (uint16, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getThreshold", _rewardEpochId)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0x4615d5e9.
//
// Solidity: function getThreshold(uint256 _rewardEpochId) view returns(uint16)
func (_FlareSystemManager *FlareSystemManagerSession) GetThreshold(_rewardEpochId *big.Int) (uint16, error) {
	return _FlareSystemManager.Contract.GetThreshold(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetThreshold is a free data retrieval call binding the contract method 0x4615d5e9.
//
// Solidity: function getThreshold(uint256 _rewardEpochId) view returns(uint16)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetThreshold(_rewardEpochId *big.Int) (uint16, error) {
	return _FlareSystemManager.Contract.GetThreshold(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetUptimeVoteSignStartInfo is a free data retrieval call binding the contract method 0xc9f1d2aa.
//
// Solidity: function getUptimeVoteSignStartInfo(uint24 _rewardEpochId) view returns(uint64 _uptimeVoteSignStartTs, uint64 _uptimeVoteSignStartBlock)
func (_FlareSystemManager *FlareSystemManagerCaller) GetUptimeVoteSignStartInfo(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	UptimeVoteSignStartTs    uint64
	UptimeVoteSignStartBlock uint64
}, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getUptimeVoteSignStartInfo", _rewardEpochId)

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
func (_FlareSystemManager *FlareSystemManagerSession) GetUptimeVoteSignStartInfo(_rewardEpochId *big.Int) (struct {
	UptimeVoteSignStartTs    uint64
	UptimeVoteSignStartBlock uint64
}, error) {
	return _FlareSystemManager.Contract.GetUptimeVoteSignStartInfo(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetUptimeVoteSignStartInfo is a free data retrieval call binding the contract method 0xc9f1d2aa.
//
// Solidity: function getUptimeVoteSignStartInfo(uint24 _rewardEpochId) view returns(uint64 _uptimeVoteSignStartTs, uint64 _uptimeVoteSignStartBlock)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetUptimeVoteSignStartInfo(_rewardEpochId *big.Int) (struct {
	UptimeVoteSignStartTs    uint64
	UptimeVoteSignStartBlock uint64
}, error) {
	return _FlareSystemManager.Contract.GetUptimeVoteSignStartInfo(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetVotePowerBlock is a free data retrieval call binding the contract method 0xc2632216.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpochId) view returns(uint64 _votePowerBlock)
func (_FlareSystemManager *FlareSystemManagerCaller) GetVotePowerBlock(opts *bind.CallOpts, _rewardEpochId *big.Int) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getVotePowerBlock", _rewardEpochId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetVotePowerBlock is a free data retrieval call binding the contract method 0xc2632216.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpochId) view returns(uint64 _votePowerBlock)
func (_FlareSystemManager *FlareSystemManagerSession) GetVotePowerBlock(_rewardEpochId *big.Int) (uint64, error) {
	return _FlareSystemManager.Contract.GetVotePowerBlock(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetVotePowerBlock is a free data retrieval call binding the contract method 0xc2632216.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpochId) view returns(uint64 _votePowerBlock)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetVotePowerBlock(_rewardEpochId *big.Int) (uint64, error) {
	return _FlareSystemManager.Contract.GetVotePowerBlock(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetVoterRegistrationData is a free data retrieval call binding the contract method 0x1703a788.
//
// Solidity: function getVoterRegistrationData(uint256 _rewardEpochId) view returns(uint256 _votePowerBlock, bool _enabled)
func (_FlareSystemManager *FlareSystemManagerCaller) GetVoterRegistrationData(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	VotePowerBlock *big.Int
	Enabled        bool
}, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getVoterRegistrationData", _rewardEpochId)

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
func (_FlareSystemManager *FlareSystemManagerSession) GetVoterRegistrationData(_rewardEpochId *big.Int) (struct {
	VotePowerBlock *big.Int
	Enabled        bool
}, error) {
	return _FlareSystemManager.Contract.GetVoterRegistrationData(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetVoterRegistrationData is a free data retrieval call binding the contract method 0x1703a788.
//
// Solidity: function getVoterRegistrationData(uint256 _rewardEpochId) view returns(uint256 _votePowerBlock, bool _enabled)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetVoterRegistrationData(_rewardEpochId *big.Int) (struct {
	VotePowerBlock *big.Int
	Enabled        bool
}, error) {
	return _FlareSystemManager.Contract.GetVoterRegistrationData(&_FlareSystemManager.CallOpts, _rewardEpochId)
}

// GetVoterRewardsSignInfo is a free data retrieval call binding the contract method 0x1916e915.
//
// Solidity: function getVoterRewardsSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _rewardsSignTs, uint64 _rewardsSignBlock)
func (_FlareSystemManager *FlareSystemManagerCaller) GetVoterRewardsSignInfo(opts *bind.CallOpts, _rewardEpochId *big.Int, _voter common.Address) (struct {
	RewardsSignTs    uint64
	RewardsSignBlock uint64
}, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getVoterRewardsSignInfo", _rewardEpochId, _voter)

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
func (_FlareSystemManager *FlareSystemManagerSession) GetVoterRewardsSignInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	RewardsSignTs    uint64
	RewardsSignBlock uint64
}, error) {
	return _FlareSystemManager.Contract.GetVoterRewardsSignInfo(&_FlareSystemManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterRewardsSignInfo is a free data retrieval call binding the contract method 0x1916e915.
//
// Solidity: function getVoterRewardsSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _rewardsSignTs, uint64 _rewardsSignBlock)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetVoterRewardsSignInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	RewardsSignTs    uint64
	RewardsSignBlock uint64
}, error) {
	return _FlareSystemManager.Contract.GetVoterRewardsSignInfo(&_FlareSystemManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterSigningPolicySignInfo is a free data retrieval call binding the contract method 0xdac4319d.
//
// Solidity: function getVoterSigningPolicySignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _signingPolicySignTs, uint64 _signingPolicySignBlock)
func (_FlareSystemManager *FlareSystemManagerCaller) GetVoterSigningPolicySignInfo(opts *bind.CallOpts, _rewardEpochId *big.Int, _voter common.Address) (struct {
	SigningPolicySignTs    uint64
	SigningPolicySignBlock uint64
}, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getVoterSigningPolicySignInfo", _rewardEpochId, _voter)

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
func (_FlareSystemManager *FlareSystemManagerSession) GetVoterSigningPolicySignInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	SigningPolicySignTs    uint64
	SigningPolicySignBlock uint64
}, error) {
	return _FlareSystemManager.Contract.GetVoterSigningPolicySignInfo(&_FlareSystemManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterSigningPolicySignInfo is a free data retrieval call binding the contract method 0xdac4319d.
//
// Solidity: function getVoterSigningPolicySignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _signingPolicySignTs, uint64 _signingPolicySignBlock)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetVoterSigningPolicySignInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	SigningPolicySignTs    uint64
	SigningPolicySignBlock uint64
}, error) {
	return _FlareSystemManager.Contract.GetVoterSigningPolicySignInfo(&_FlareSystemManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterUptimeVoteSignInfo is a free data retrieval call binding the contract method 0x41c05ad5.
//
// Solidity: function getVoterUptimeVoteSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSignTs, uint64 _uptimeVoteSignBlock)
func (_FlareSystemManager *FlareSystemManagerCaller) GetVoterUptimeVoteSignInfo(opts *bind.CallOpts, _rewardEpochId *big.Int, _voter common.Address) (struct {
	UptimeVoteSignTs    uint64
	UptimeVoteSignBlock uint64
}, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getVoterUptimeVoteSignInfo", _rewardEpochId, _voter)

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
func (_FlareSystemManager *FlareSystemManagerSession) GetVoterUptimeVoteSignInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	UptimeVoteSignTs    uint64
	UptimeVoteSignBlock uint64
}, error) {
	return _FlareSystemManager.Contract.GetVoterUptimeVoteSignInfo(&_FlareSystemManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterUptimeVoteSignInfo is a free data retrieval call binding the contract method 0x41c05ad5.
//
// Solidity: function getVoterUptimeVoteSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSignTs, uint64 _uptimeVoteSignBlock)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetVoterUptimeVoteSignInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	UptimeVoteSignTs    uint64
	UptimeVoteSignBlock uint64
}, error) {
	return _FlareSystemManager.Contract.GetVoterUptimeVoteSignInfo(&_FlareSystemManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterUptimeVoteSubmitInfo is a free data retrieval call binding the contract method 0x59db0e2f.
//
// Solidity: function getVoterUptimeVoteSubmitInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSubmitTs, uint64 _uptimeVoteSubmitBlock)
func (_FlareSystemManager *FlareSystemManagerCaller) GetVoterUptimeVoteSubmitInfo(opts *bind.CallOpts, _rewardEpochId *big.Int, _voter common.Address) (struct {
	UptimeVoteSubmitTs    uint64
	UptimeVoteSubmitBlock uint64
}, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getVoterUptimeVoteSubmitInfo", _rewardEpochId, _voter)

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
func (_FlareSystemManager *FlareSystemManagerSession) GetVoterUptimeVoteSubmitInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	UptimeVoteSubmitTs    uint64
	UptimeVoteSubmitBlock uint64
}, error) {
	return _FlareSystemManager.Contract.GetVoterUptimeVoteSubmitInfo(&_FlareSystemManager.CallOpts, _rewardEpochId, _voter)
}

// GetVoterUptimeVoteSubmitInfo is a free data retrieval call binding the contract method 0x59db0e2f.
//
// Solidity: function getVoterUptimeVoteSubmitInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSubmitTs, uint64 _uptimeVoteSubmitBlock)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetVoterUptimeVoteSubmitInfo(_rewardEpochId *big.Int, _voter common.Address) (struct {
	UptimeVoteSubmitTs    uint64
	UptimeVoteSubmitBlock uint64
}, error) {
	return _FlareSystemManager.Contract.GetVoterUptimeVoteSubmitInfo(&_FlareSystemManager.CallOpts, _rewardEpochId, _voter)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FlareSystemManager *FlareSystemManagerSession) Governance() (common.Address, error) {
	return _FlareSystemManager.Contract.Governance(&_FlareSystemManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCallerSession) Governance() (common.Address, error) {
	return _FlareSystemManager.Contract.Governance(&_FlareSystemManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FlareSystemManager *FlareSystemManagerSession) GovernanceSettings() (common.Address, error) {
	return _FlareSystemManager.Contract.GovernanceSettings(&_FlareSystemManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GovernanceSettings() (common.Address, error) {
	return _FlareSystemManager.Contract.GovernanceSettings(&_FlareSystemManager.CallOpts)
}

// InitialRandomVotePowerBlockSelectionSize is a free data retrieval call binding the contract method 0xded7c4b8.
//
// Solidity: function initialRandomVotePowerBlockSelectionSize() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) InitialRandomVotePowerBlockSelectionSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "initialRandomVotePowerBlockSelectionSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InitialRandomVotePowerBlockSelectionSize is a free data retrieval call binding the contract method 0xded7c4b8.
//
// Solidity: function initialRandomVotePowerBlockSelectionSize() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) InitialRandomVotePowerBlockSelectionSize() (uint64, error) {
	return _FlareSystemManager.Contract.InitialRandomVotePowerBlockSelectionSize(&_FlareSystemManager.CallOpts)
}

// InitialRandomVotePowerBlockSelectionSize is a free data retrieval call binding the contract method 0xded7c4b8.
//
// Solidity: function initialRandomVotePowerBlockSelectionSize() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) InitialRandomVotePowerBlockSelectionSize() (uint64, error) {
	return _FlareSystemManager.Contract.InitialRandomVotePowerBlockSelectionSize(&_FlareSystemManager.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FlareSystemManager *FlareSystemManagerSession) IsExecutor(_address common.Address) (bool, error) {
	return _FlareSystemManager.Contract.IsExecutor(&_FlareSystemManager.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _FlareSystemManager.Contract.IsExecutor(&_FlareSystemManager.CallOpts, _address)
}

// IsVoterRegistrationEnabled is a free data retrieval call binding the contract method 0x09505d25.
//
// Solidity: function isVoterRegistrationEnabled() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCaller) IsVoterRegistrationEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "isVoterRegistrationEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoterRegistrationEnabled is a free data retrieval call binding the contract method 0x09505d25.
//
// Solidity: function isVoterRegistrationEnabled() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerSession) IsVoterRegistrationEnabled() (bool, error) {
	return _FlareSystemManager.Contract.IsVoterRegistrationEnabled(&_FlareSystemManager.CallOpts)
}

// IsVoterRegistrationEnabled is a free data retrieval call binding the contract method 0x09505d25.
//
// Solidity: function isVoterRegistrationEnabled() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCallerSession) IsVoterRegistrationEnabled() (bool, error) {
	return _FlareSystemManager.Contract.IsVoterRegistrationEnabled(&_FlareSystemManager.CallOpts)
}

// LastInitializedVotingRoundId is a free data retrieval call binding the contract method 0x4f923d37.
//
// Solidity: function lastInitializedVotingRoundId() view returns(uint32)
func (_FlareSystemManager *FlareSystemManagerCaller) LastInitializedVotingRoundId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "lastInitializedVotingRoundId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastInitializedVotingRoundId is a free data retrieval call binding the contract method 0x4f923d37.
//
// Solidity: function lastInitializedVotingRoundId() view returns(uint32)
func (_FlareSystemManager *FlareSystemManagerSession) LastInitializedVotingRoundId() (uint32, error) {
	return _FlareSystemManager.Contract.LastInitializedVotingRoundId(&_FlareSystemManager.CallOpts)
}

// LastInitializedVotingRoundId is a free data retrieval call binding the contract method 0x4f923d37.
//
// Solidity: function lastInitializedVotingRoundId() view returns(uint32)
func (_FlareSystemManager *FlareSystemManagerCallerSession) LastInitializedVotingRoundId() (uint32, error) {
	return _FlareSystemManager.Contract.LastInitializedVotingRoundId(&_FlareSystemManager.CallOpts)
}

// NewSigningPolicyInitializationStartSeconds is a free data retrieval call binding the contract method 0x6aeffddc.
//
// Solidity: function newSigningPolicyInitializationStartSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) NewSigningPolicyInitializationStartSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "newSigningPolicyInitializationStartSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NewSigningPolicyInitializationStartSeconds is a free data retrieval call binding the contract method 0x6aeffddc.
//
// Solidity: function newSigningPolicyInitializationStartSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) NewSigningPolicyInitializationStartSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.NewSigningPolicyInitializationStartSeconds(&_FlareSystemManager.CallOpts)
}

// NewSigningPolicyInitializationStartSeconds is a free data retrieval call binding the contract method 0x6aeffddc.
//
// Solidity: function newSigningPolicyInitializationStartSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) NewSigningPolicyInitializationStartSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.NewSigningPolicyInitializationStartSeconds(&_FlareSystemManager.CallOpts)
}

// NewSigningPolicyMinNumberOfVotingRoundsDelay is a free data retrieval call binding the contract method 0xa733d54b.
//
// Solidity: function newSigningPolicyMinNumberOfVotingRoundsDelay() view returns(uint32)
func (_FlareSystemManager *FlareSystemManagerCaller) NewSigningPolicyMinNumberOfVotingRoundsDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "newSigningPolicyMinNumberOfVotingRoundsDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NewSigningPolicyMinNumberOfVotingRoundsDelay is a free data retrieval call binding the contract method 0xa733d54b.
//
// Solidity: function newSigningPolicyMinNumberOfVotingRoundsDelay() view returns(uint32)
func (_FlareSystemManager *FlareSystemManagerSession) NewSigningPolicyMinNumberOfVotingRoundsDelay() (uint32, error) {
	return _FlareSystemManager.Contract.NewSigningPolicyMinNumberOfVotingRoundsDelay(&_FlareSystemManager.CallOpts)
}

// NewSigningPolicyMinNumberOfVotingRoundsDelay is a free data retrieval call binding the contract method 0xa733d54b.
//
// Solidity: function newSigningPolicyMinNumberOfVotingRoundsDelay() view returns(uint32)
func (_FlareSystemManager *FlareSystemManagerCallerSession) NewSigningPolicyMinNumberOfVotingRoundsDelay() (uint32, error) {
	return _FlareSystemManager.Contract.NewSigningPolicyMinNumberOfVotingRoundsDelay(&_FlareSystemManager.CallOpts)
}

// NoOfWeightBasedClaims is a free data retrieval call binding the contract method 0x388a4c3d.
//
// Solidity: function noOfWeightBasedClaims(uint256 ) view returns(uint256)
func (_FlareSystemManager *FlareSystemManagerCaller) NoOfWeightBasedClaims(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "noOfWeightBasedClaims", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoOfWeightBasedClaims is a free data retrieval call binding the contract method 0x388a4c3d.
//
// Solidity: function noOfWeightBasedClaims(uint256 ) view returns(uint256)
func (_FlareSystemManager *FlareSystemManagerSession) NoOfWeightBasedClaims(arg0 *big.Int) (*big.Int, error) {
	return _FlareSystemManager.Contract.NoOfWeightBasedClaims(&_FlareSystemManager.CallOpts, arg0)
}

// NoOfWeightBasedClaims is a free data retrieval call binding the contract method 0x388a4c3d.
//
// Solidity: function noOfWeightBasedClaims(uint256 ) view returns(uint256)
func (_FlareSystemManager *FlareSystemManagerCallerSession) NoOfWeightBasedClaims(arg0 *big.Int) (*big.Int, error) {
	return _FlareSystemManager.Contract.NoOfWeightBasedClaims(&_FlareSystemManager.CallOpts, arg0)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerSession) ProductionMode() (bool, error) {
	return _FlareSystemManager.Contract.ProductionMode(&_FlareSystemManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCallerSession) ProductionMode() (bool, error) {
	return _FlareSystemManager.Contract.ProductionMode(&_FlareSystemManager.CallOpts)
}

// RandomAcquisitionMaxDurationBlocks is a free data retrieval call binding the contract method 0x490344f4.
//
// Solidity: function randomAcquisitionMaxDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) RandomAcquisitionMaxDurationBlocks(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "randomAcquisitionMaxDurationBlocks")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RandomAcquisitionMaxDurationBlocks is a free data retrieval call binding the contract method 0x490344f4.
//
// Solidity: function randomAcquisitionMaxDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) RandomAcquisitionMaxDurationBlocks() (uint64, error) {
	return _FlareSystemManager.Contract.RandomAcquisitionMaxDurationBlocks(&_FlareSystemManager.CallOpts)
}

// RandomAcquisitionMaxDurationBlocks is a free data retrieval call binding the contract method 0x490344f4.
//
// Solidity: function randomAcquisitionMaxDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) RandomAcquisitionMaxDurationBlocks() (uint64, error) {
	return _FlareSystemManager.Contract.RandomAcquisitionMaxDurationBlocks(&_FlareSystemManager.CallOpts)
}

// RandomAcquisitionMaxDurationSeconds is a free data retrieval call binding the contract method 0x098e7ff6.
//
// Solidity: function randomAcquisitionMaxDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) RandomAcquisitionMaxDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "randomAcquisitionMaxDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RandomAcquisitionMaxDurationSeconds is a free data retrieval call binding the contract method 0x098e7ff6.
//
// Solidity: function randomAcquisitionMaxDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) RandomAcquisitionMaxDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.RandomAcquisitionMaxDurationSeconds(&_FlareSystemManager.CallOpts)
}

// RandomAcquisitionMaxDurationSeconds is a free data retrieval call binding the contract method 0x098e7ff6.
//
// Solidity: function randomAcquisitionMaxDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) RandomAcquisitionMaxDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.RandomAcquisitionMaxDurationSeconds(&_FlareSystemManager.CallOpts)
}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCaller) Relay(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "relay")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_FlareSystemManager *FlareSystemManagerSession) Relay() (common.Address, error) {
	return _FlareSystemManager.Contract.Relay(&_FlareSystemManager.CallOpts)
}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCallerSession) Relay() (common.Address, error) {
	return _FlareSystemManager.Contract.Relay(&_FlareSystemManager.CallOpts)
}

// RewardEpochDurationSeconds is a free data retrieval call binding the contract method 0x85f3c9c9.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) RewardEpochDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "rewardEpochDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RewardEpochDurationSeconds is a free data retrieval call binding the contract method 0x85f3c9c9.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) RewardEpochDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.RewardEpochDurationSeconds(&_FlareSystemManager.CallOpts)
}

// RewardEpochDurationSeconds is a free data retrieval call binding the contract method 0x85f3c9c9.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) RewardEpochDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.RewardEpochDurationSeconds(&_FlareSystemManager.CallOpts)
}

// RewardEpochIdToExpireNext is a free data retrieval call binding the contract method 0xaec84ab6.
//
// Solidity: function rewardEpochIdToExpireNext() view returns(uint24)
func (_FlareSystemManager *FlareSystemManagerCaller) RewardEpochIdToExpireNext(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "rewardEpochIdToExpireNext")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardEpochIdToExpireNext is a free data retrieval call binding the contract method 0xaec84ab6.
//
// Solidity: function rewardEpochIdToExpireNext() view returns(uint24)
func (_FlareSystemManager *FlareSystemManagerSession) RewardEpochIdToExpireNext() (*big.Int, error) {
	return _FlareSystemManager.Contract.RewardEpochIdToExpireNext(&_FlareSystemManager.CallOpts)
}

// RewardEpochIdToExpireNext is a free data retrieval call binding the contract method 0xaec84ab6.
//
// Solidity: function rewardEpochIdToExpireNext() view returns(uint24)
func (_FlareSystemManager *FlareSystemManagerCallerSession) RewardEpochIdToExpireNext() (*big.Int, error) {
	return _FlareSystemManager.Contract.RewardEpochIdToExpireNext(&_FlareSystemManager.CallOpts)
}

// RewardExpiryOffsetSeconds is a free data retrieval call binding the contract method 0x4eaee307.
//
// Solidity: function rewardExpiryOffsetSeconds() view returns(uint32)
func (_FlareSystemManager *FlareSystemManagerCaller) RewardExpiryOffsetSeconds(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "rewardExpiryOffsetSeconds")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RewardExpiryOffsetSeconds is a free data retrieval call binding the contract method 0x4eaee307.
//
// Solidity: function rewardExpiryOffsetSeconds() view returns(uint32)
func (_FlareSystemManager *FlareSystemManagerSession) RewardExpiryOffsetSeconds() (uint32, error) {
	return _FlareSystemManager.Contract.RewardExpiryOffsetSeconds(&_FlareSystemManager.CallOpts)
}

// RewardExpiryOffsetSeconds is a free data retrieval call binding the contract method 0x4eaee307.
//
// Solidity: function rewardExpiryOffsetSeconds() view returns(uint32)
func (_FlareSystemManager *FlareSystemManagerCallerSession) RewardExpiryOffsetSeconds() (uint32, error) {
	return _FlareSystemManager.Contract.RewardExpiryOffsetSeconds(&_FlareSystemManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FlareSystemManager *FlareSystemManagerSession) RewardManager() (common.Address, error) {
	return _FlareSystemManager.Contract.RewardManager(&_FlareSystemManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCallerSession) RewardManager() (common.Address, error) {
	return _FlareSystemManager.Contract.RewardManager(&_FlareSystemManager.CallOpts)
}

// RewardsHash is a free data retrieval call binding the contract method 0x647006e2.
//
// Solidity: function rewardsHash(uint256 ) view returns(bytes32)
func (_FlareSystemManager *FlareSystemManagerCaller) RewardsHash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "rewardsHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RewardsHash is a free data retrieval call binding the contract method 0x647006e2.
//
// Solidity: function rewardsHash(uint256 ) view returns(bytes32)
func (_FlareSystemManager *FlareSystemManagerSession) RewardsHash(arg0 *big.Int) ([32]byte, error) {
	return _FlareSystemManager.Contract.RewardsHash(&_FlareSystemManager.CallOpts, arg0)
}

// RewardsHash is a free data retrieval call binding the contract method 0x647006e2.
//
// Solidity: function rewardsHash(uint256 ) view returns(bytes32)
func (_FlareSystemManager *FlareSystemManagerCallerSession) RewardsHash(arg0 *big.Int) ([32]byte, error) {
	return _FlareSystemManager.Contract.RewardsHash(&_FlareSystemManager.CallOpts, arg0)
}

// SigningPolicyMinNumberOfVoters is a free data retrieval call binding the contract method 0x2e3645f8.
//
// Solidity: function signingPolicyMinNumberOfVoters() view returns(uint16)
func (_FlareSystemManager *FlareSystemManagerCaller) SigningPolicyMinNumberOfVoters(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "signingPolicyMinNumberOfVoters")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SigningPolicyMinNumberOfVoters is a free data retrieval call binding the contract method 0x2e3645f8.
//
// Solidity: function signingPolicyMinNumberOfVoters() view returns(uint16)
func (_FlareSystemManager *FlareSystemManagerSession) SigningPolicyMinNumberOfVoters() (uint16, error) {
	return _FlareSystemManager.Contract.SigningPolicyMinNumberOfVoters(&_FlareSystemManager.CallOpts)
}

// SigningPolicyMinNumberOfVoters is a free data retrieval call binding the contract method 0x2e3645f8.
//
// Solidity: function signingPolicyMinNumberOfVoters() view returns(uint16)
func (_FlareSystemManager *FlareSystemManagerCallerSession) SigningPolicyMinNumberOfVoters() (uint16, error) {
	return _FlareSystemManager.Contract.SigningPolicyMinNumberOfVoters(&_FlareSystemManager.CallOpts)
}

// SigningPolicyThresholdPPM is a free data retrieval call binding the contract method 0xf21d6304.
//
// Solidity: function signingPolicyThresholdPPM() view returns(uint24)
func (_FlareSystemManager *FlareSystemManagerCaller) SigningPolicyThresholdPPM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "signingPolicyThresholdPPM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SigningPolicyThresholdPPM is a free data retrieval call binding the contract method 0xf21d6304.
//
// Solidity: function signingPolicyThresholdPPM() view returns(uint24)
func (_FlareSystemManager *FlareSystemManagerSession) SigningPolicyThresholdPPM() (*big.Int, error) {
	return _FlareSystemManager.Contract.SigningPolicyThresholdPPM(&_FlareSystemManager.CallOpts)
}

// SigningPolicyThresholdPPM is a free data retrieval call binding the contract method 0xf21d6304.
//
// Solidity: function signingPolicyThresholdPPM() view returns(uint24)
func (_FlareSystemManager *FlareSystemManagerCallerSession) SigningPolicyThresholdPPM() (*big.Int, error) {
	return _FlareSystemManager.Contract.SigningPolicyThresholdPPM(&_FlareSystemManager.CallOpts)
}

// Submission is a free data retrieval call binding the contract method 0x9a759097.
//
// Solidity: function submission() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCaller) Submission(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "submission")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Submission is a free data retrieval call binding the contract method 0x9a759097.
//
// Solidity: function submission() view returns(address)
func (_FlareSystemManager *FlareSystemManagerSession) Submission() (common.Address, error) {
	return _FlareSystemManager.Contract.Submission(&_FlareSystemManager.CallOpts)
}

// Submission is a free data retrieval call binding the contract method 0x9a759097.
//
// Solidity: function submission() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCallerSession) Submission() (common.Address, error) {
	return _FlareSystemManager.Contract.Submission(&_FlareSystemManager.CallOpts)
}

// Submit3Aligned is a free data retrieval call binding the contract method 0x107d8ffb.
//
// Solidity: function submit3Aligned() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCaller) Submit3Aligned(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "submit3Aligned")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Submit3Aligned is a free data retrieval call binding the contract method 0x107d8ffb.
//
// Solidity: function submit3Aligned() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerSession) Submit3Aligned() (bool, error) {
	return _FlareSystemManager.Contract.Submit3Aligned(&_FlareSystemManager.CallOpts)
}

// Submit3Aligned is a free data retrieval call binding the contract method 0x107d8ffb.
//
// Solidity: function submit3Aligned() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCallerSession) Submit3Aligned() (bool, error) {
	return _FlareSystemManager.Contract.Submit3Aligned(&_FlareSystemManager.CallOpts)
}

// SubmitUptimeVoteMinDurationBlocks is a free data retrieval call binding the contract method 0xd8a01a0a.
//
// Solidity: function submitUptimeVoteMinDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) SubmitUptimeVoteMinDurationBlocks(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "submitUptimeVoteMinDurationBlocks")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SubmitUptimeVoteMinDurationBlocks is a free data retrieval call binding the contract method 0xd8a01a0a.
//
// Solidity: function submitUptimeVoteMinDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) SubmitUptimeVoteMinDurationBlocks() (uint64, error) {
	return _FlareSystemManager.Contract.SubmitUptimeVoteMinDurationBlocks(&_FlareSystemManager.CallOpts)
}

// SubmitUptimeVoteMinDurationBlocks is a free data retrieval call binding the contract method 0xd8a01a0a.
//
// Solidity: function submitUptimeVoteMinDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) SubmitUptimeVoteMinDurationBlocks() (uint64, error) {
	return _FlareSystemManager.Contract.SubmitUptimeVoteMinDurationBlocks(&_FlareSystemManager.CallOpts)
}

// SubmitUptimeVoteMinDurationSeconds is a free data retrieval call binding the contract method 0x4c528765.
//
// Solidity: function submitUptimeVoteMinDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) SubmitUptimeVoteMinDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "submitUptimeVoteMinDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SubmitUptimeVoteMinDurationSeconds is a free data retrieval call binding the contract method 0x4c528765.
//
// Solidity: function submitUptimeVoteMinDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) SubmitUptimeVoteMinDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.SubmitUptimeVoteMinDurationSeconds(&_FlareSystemManager.CallOpts)
}

// SubmitUptimeVoteMinDurationSeconds is a free data retrieval call binding the contract method 0x4c528765.
//
// Solidity: function submitUptimeVoteMinDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) SubmitUptimeVoteMinDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.SubmitUptimeVoteMinDurationSeconds(&_FlareSystemManager.CallOpts)
}

// SwitchToFallbackMode is a free data retrieval call binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCaller) SwitchToFallbackMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "switchToFallbackMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwitchToFallbackMode is a free data retrieval call binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerSession) SwitchToFallbackMode() (bool, error) {
	return _FlareSystemManager.Contract.SwitchToFallbackMode(&_FlareSystemManager.CallOpts)
}

// SwitchToFallbackMode is a free data retrieval call binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCallerSession) SwitchToFallbackMode() (bool, error) {
	return _FlareSystemManager.Contract.SwitchToFallbackMode(&_FlareSystemManager.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FlareSystemManager *FlareSystemManagerCaller) TimelockedCalls(opts *bind.CallOpts, arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "timelockedCalls", arg0)

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
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FlareSystemManager *FlareSystemManagerSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FlareSystemManager.Contract.TimelockedCalls(&_FlareSystemManager.CallOpts, arg0)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FlareSystemManager *FlareSystemManagerCallerSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FlareSystemManager.Contract.TimelockedCalls(&_FlareSystemManager.CallOpts, arg0)
}

// TriggerExpirationAndCleanup is a free data retrieval call binding the contract method 0x9b760d13.
//
// Solidity: function triggerExpirationAndCleanup() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCaller) TriggerExpirationAndCleanup(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "triggerExpirationAndCleanup")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TriggerExpirationAndCleanup is a free data retrieval call binding the contract method 0x9b760d13.
//
// Solidity: function triggerExpirationAndCleanup() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerSession) TriggerExpirationAndCleanup() (bool, error) {
	return _FlareSystemManager.Contract.TriggerExpirationAndCleanup(&_FlareSystemManager.CallOpts)
}

// TriggerExpirationAndCleanup is a free data retrieval call binding the contract method 0x9b760d13.
//
// Solidity: function triggerExpirationAndCleanup() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCallerSession) TriggerExpirationAndCleanup() (bool, error) {
	return _FlareSystemManager.Contract.TriggerExpirationAndCleanup(&_FlareSystemManager.CallOpts)
}

// UptimeVoteHash is a free data retrieval call binding the contract method 0xd3466911.
//
// Solidity: function uptimeVoteHash(uint256 ) view returns(bytes32)
func (_FlareSystemManager *FlareSystemManagerCaller) UptimeVoteHash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "uptimeVoteHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UptimeVoteHash is a free data retrieval call binding the contract method 0xd3466911.
//
// Solidity: function uptimeVoteHash(uint256 ) view returns(bytes32)
func (_FlareSystemManager *FlareSystemManagerSession) UptimeVoteHash(arg0 *big.Int) ([32]byte, error) {
	return _FlareSystemManager.Contract.UptimeVoteHash(&_FlareSystemManager.CallOpts, arg0)
}

// UptimeVoteHash is a free data retrieval call binding the contract method 0xd3466911.
//
// Solidity: function uptimeVoteHash(uint256 ) view returns(bytes32)
func (_FlareSystemManager *FlareSystemManagerCallerSession) UptimeVoteHash(arg0 *big.Int) ([32]byte, error) {
	return _FlareSystemManager.Contract.UptimeVoteHash(&_FlareSystemManager.CallOpts, arg0)
}

// VoterRegistrationMinDurationBlocks is a free data retrieval call binding the contract method 0xd10e807f.
//
// Solidity: function voterRegistrationMinDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) VoterRegistrationMinDurationBlocks(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "voterRegistrationMinDurationBlocks")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VoterRegistrationMinDurationBlocks is a free data retrieval call binding the contract method 0xd10e807f.
//
// Solidity: function voterRegistrationMinDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) VoterRegistrationMinDurationBlocks() (uint64, error) {
	return _FlareSystemManager.Contract.VoterRegistrationMinDurationBlocks(&_FlareSystemManager.CallOpts)
}

// VoterRegistrationMinDurationBlocks is a free data retrieval call binding the contract method 0xd10e807f.
//
// Solidity: function voterRegistrationMinDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) VoterRegistrationMinDurationBlocks() (uint64, error) {
	return _FlareSystemManager.Contract.VoterRegistrationMinDurationBlocks(&_FlareSystemManager.CallOpts)
}

// VoterRegistrationMinDurationSeconds is a free data retrieval call binding the contract method 0xa219fe02.
//
// Solidity: function voterRegistrationMinDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) VoterRegistrationMinDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "voterRegistrationMinDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VoterRegistrationMinDurationSeconds is a free data retrieval call binding the contract method 0xa219fe02.
//
// Solidity: function voterRegistrationMinDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) VoterRegistrationMinDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.VoterRegistrationMinDurationSeconds(&_FlareSystemManager.CallOpts)
}

// VoterRegistrationMinDurationSeconds is a free data retrieval call binding the contract method 0xa219fe02.
//
// Solidity: function voterRegistrationMinDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) VoterRegistrationMinDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.VoterRegistrationMinDurationSeconds(&_FlareSystemManager.CallOpts)
}

// VoterRegistrationTriggerContract is a free data retrieval call binding the contract method 0x88e49ac7.
//
// Solidity: function voterRegistrationTriggerContract() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCaller) VoterRegistrationTriggerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "voterRegistrationTriggerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoterRegistrationTriggerContract is a free data retrieval call binding the contract method 0x88e49ac7.
//
// Solidity: function voterRegistrationTriggerContract() view returns(address)
func (_FlareSystemManager *FlareSystemManagerSession) VoterRegistrationTriggerContract() (common.Address, error) {
	return _FlareSystemManager.Contract.VoterRegistrationTriggerContract(&_FlareSystemManager.CallOpts)
}

// VoterRegistrationTriggerContract is a free data retrieval call binding the contract method 0x88e49ac7.
//
// Solidity: function voterRegistrationTriggerContract() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCallerSession) VoterRegistrationTriggerContract() (common.Address, error) {
	return _FlareSystemManager.Contract.VoterRegistrationTriggerContract(&_FlareSystemManager.CallOpts)
}

// VoterRegistry is a free data retrieval call binding the contract method 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCaller) VoterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "voterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoterRegistry is a free data retrieval call binding the contract method 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (_FlareSystemManager *FlareSystemManagerSession) VoterRegistry() (common.Address, error) {
	return _FlareSystemManager.Contract.VoterRegistry(&_FlareSystemManager.CallOpts)
}

// VoterRegistry is a free data retrieval call binding the contract method 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCallerSession) VoterRegistry() (common.Address, error) {
	return _FlareSystemManager.Contract.VoterRegistry(&_FlareSystemManager.CallOpts)
}

// VotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) VotingEpochDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "votingEpochDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) VotingEpochDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.VotingEpochDurationSeconds(&_FlareSystemManager.CallOpts)
}

// VotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) VotingEpochDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.VotingEpochDurationSeconds(&_FlareSystemManager.CallOpts)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FlareSystemManager *FlareSystemManagerSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.CancelGovernanceCall(&_FlareSystemManager.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.CancelGovernanceCall(&_FlareSystemManager.TransactOpts, _selector)
}

// ChangeSigningPolicySettings is a paid mutator transaction binding the contract method 0x898805c1.
//
// Solidity: function changeSigningPolicySettings(uint24 _signingPolicyThresholdPPM, uint16 _signingPolicyMinNumberOfVoters) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) ChangeSigningPolicySettings(opts *bind.TransactOpts, _signingPolicyThresholdPPM *big.Int, _signingPolicyMinNumberOfVoters uint16) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "changeSigningPolicySettings", _signingPolicyThresholdPPM, _signingPolicyMinNumberOfVoters)
}

// ChangeSigningPolicySettings is a paid mutator transaction binding the contract method 0x898805c1.
//
// Solidity: function changeSigningPolicySettings(uint24 _signingPolicyThresholdPPM, uint16 _signingPolicyMinNumberOfVoters) returns()
func (_FlareSystemManager *FlareSystemManagerSession) ChangeSigningPolicySettings(_signingPolicyThresholdPPM *big.Int, _signingPolicyMinNumberOfVoters uint16) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.ChangeSigningPolicySettings(&_FlareSystemManager.TransactOpts, _signingPolicyThresholdPPM, _signingPolicyMinNumberOfVoters)
}

// ChangeSigningPolicySettings is a paid mutator transaction binding the contract method 0x898805c1.
//
// Solidity: function changeSigningPolicySettings(uint24 _signingPolicyThresholdPPM, uint16 _signingPolicyMinNumberOfVoters) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) ChangeSigningPolicySettings(_signingPolicyThresholdPPM *big.Int, _signingPolicyMinNumberOfVoters uint16) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.ChangeSigningPolicySettings(&_FlareSystemManager.TransactOpts, _signingPolicyThresholdPPM, _signingPolicyMinNumberOfVoters)
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_FlareSystemManager *FlareSystemManagerTransactor) Daemonize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "daemonize")
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_FlareSystemManager *FlareSystemManagerSession) Daemonize() (*types.Transaction, error) {
	return _FlareSystemManager.Contract.Daemonize(&_FlareSystemManager.TransactOpts)
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_FlareSystemManager *FlareSystemManagerTransactorSession) Daemonize() (*types.Transaction, error) {
	return _FlareSystemManager.Contract.Daemonize(&_FlareSystemManager.TransactOpts)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FlareSystemManager *FlareSystemManagerSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.ExecuteGovernanceCall(&_FlareSystemManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.ExecuteGovernanceCall(&_FlareSystemManager.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FlareSystemManager *FlareSystemManagerSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.Initialise(&_FlareSystemManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.Initialise(&_FlareSystemManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// SetRewardEpochSwitchoverTriggerContracts is a paid mutator transaction binding the contract method 0x06886f41.
//
// Solidity: function setRewardEpochSwitchoverTriggerContracts(address[] _contracts) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) SetRewardEpochSwitchoverTriggerContracts(opts *bind.TransactOpts, _contracts []common.Address) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "setRewardEpochSwitchoverTriggerContracts", _contracts)
}

// SetRewardEpochSwitchoverTriggerContracts is a paid mutator transaction binding the contract method 0x06886f41.
//
// Solidity: function setRewardEpochSwitchoverTriggerContracts(address[] _contracts) returns()
func (_FlareSystemManager *FlareSystemManagerSession) SetRewardEpochSwitchoverTriggerContracts(_contracts []common.Address) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SetRewardEpochSwitchoverTriggerContracts(&_FlareSystemManager.TransactOpts, _contracts)
}

// SetRewardEpochSwitchoverTriggerContracts is a paid mutator transaction binding the contract method 0x06886f41.
//
// Solidity: function setRewardEpochSwitchoverTriggerContracts(address[] _contracts) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) SetRewardEpochSwitchoverTriggerContracts(_contracts []common.Address) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SetRewardEpochSwitchoverTriggerContracts(&_FlareSystemManager.TransactOpts, _contracts)
}

// SetRewardsData is a paid mutator transaction binding the contract method 0x2a576f8d.
//
// Solidity: function setRewardsData(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) SetRewardsData(opts *bind.TransactOpts, _rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "setRewardsData", _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash)
}

// SetRewardsData is a paid mutator transaction binding the contract method 0x2a576f8d.
//
// Solidity: function setRewardsData(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash) returns()
func (_FlareSystemManager *FlareSystemManagerSession) SetRewardsData(_rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SetRewardsData(&_FlareSystemManager.TransactOpts, _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash)
}

// SetRewardsData is a paid mutator transaction binding the contract method 0x2a576f8d.
//
// Solidity: function setRewardsData(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) SetRewardsData(_rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SetRewardsData(&_FlareSystemManager.TransactOpts, _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash)
}

// SetSubmit3Aligned is a paid mutator transaction binding the contract method 0xa72b826e.
//
// Solidity: function setSubmit3Aligned(bool _submit3Aligned) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) SetSubmit3Aligned(opts *bind.TransactOpts, _submit3Aligned bool) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "setSubmit3Aligned", _submit3Aligned)
}

// SetSubmit3Aligned is a paid mutator transaction binding the contract method 0xa72b826e.
//
// Solidity: function setSubmit3Aligned(bool _submit3Aligned) returns()
func (_FlareSystemManager *FlareSystemManagerSession) SetSubmit3Aligned(_submit3Aligned bool) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SetSubmit3Aligned(&_FlareSystemManager.TransactOpts, _submit3Aligned)
}

// SetSubmit3Aligned is a paid mutator transaction binding the contract method 0xa72b826e.
//
// Solidity: function setSubmit3Aligned(bool _submit3Aligned) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) SetSubmit3Aligned(_submit3Aligned bool) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SetSubmit3Aligned(&_FlareSystemManager.TransactOpts, _submit3Aligned)
}

// SetTriggerExpirationAndCleanup is a paid mutator transaction binding the contract method 0x67daec89.
//
// Solidity: function setTriggerExpirationAndCleanup(bool _triggerExpirationAndCleanup) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) SetTriggerExpirationAndCleanup(opts *bind.TransactOpts, _triggerExpirationAndCleanup bool) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "setTriggerExpirationAndCleanup", _triggerExpirationAndCleanup)
}

// SetTriggerExpirationAndCleanup is a paid mutator transaction binding the contract method 0x67daec89.
//
// Solidity: function setTriggerExpirationAndCleanup(bool _triggerExpirationAndCleanup) returns()
func (_FlareSystemManager *FlareSystemManagerSession) SetTriggerExpirationAndCleanup(_triggerExpirationAndCleanup bool) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SetTriggerExpirationAndCleanup(&_FlareSystemManager.TransactOpts, _triggerExpirationAndCleanup)
}

// SetTriggerExpirationAndCleanup is a paid mutator transaction binding the contract method 0x67daec89.
//
// Solidity: function setTriggerExpirationAndCleanup(bool _triggerExpirationAndCleanup) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) SetTriggerExpirationAndCleanup(_triggerExpirationAndCleanup bool) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SetTriggerExpirationAndCleanup(&_FlareSystemManager.TransactOpts, _triggerExpirationAndCleanup)
}

// SetVoterRegistrationTriggerContract is a paid mutator transaction binding the contract method 0x24eb64de.
//
// Solidity: function setVoterRegistrationTriggerContract(address _contract) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) SetVoterRegistrationTriggerContract(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "setVoterRegistrationTriggerContract", _contract)
}

// SetVoterRegistrationTriggerContract is a paid mutator transaction binding the contract method 0x24eb64de.
//
// Solidity: function setVoterRegistrationTriggerContract(address _contract) returns()
func (_FlareSystemManager *FlareSystemManagerSession) SetVoterRegistrationTriggerContract(_contract common.Address) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SetVoterRegistrationTriggerContract(&_FlareSystemManager.TransactOpts, _contract)
}

// SetVoterRegistrationTriggerContract is a paid mutator transaction binding the contract method 0x24eb64de.
//
// Solidity: function setVoterRegistrationTriggerContract(address _contract) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) SetVoterRegistrationTriggerContract(_contract common.Address) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SetVoterRegistrationTriggerContract(&_FlareSystemManager.TransactOpts, _contract)
}

// SignNewSigningPolicy is a paid mutator transaction binding the contract method 0x6b4c7bd6.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) SignNewSigningPolicy(opts *bind.TransactOpts, _rewardEpochId *big.Int, _newSigningPolicyHash [32]byte, _signature IFlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "signNewSigningPolicy", _rewardEpochId, _newSigningPolicyHash, _signature)
}

// SignNewSigningPolicy is a paid mutator transaction binding the contract method 0x6b4c7bd6.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerSession) SignNewSigningPolicy(_rewardEpochId *big.Int, _newSigningPolicyHash [32]byte, _signature IFlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SignNewSigningPolicy(&_FlareSystemManager.TransactOpts, _rewardEpochId, _newSigningPolicyHash, _signature)
}

// SignNewSigningPolicy is a paid mutator transaction binding the contract method 0x6b4c7bd6.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) SignNewSigningPolicy(_rewardEpochId *big.Int, _newSigningPolicyHash [32]byte, _signature IFlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SignNewSigningPolicy(&_FlareSystemManager.TransactOpts, _rewardEpochId, _newSigningPolicyHash, _signature)
}

// SignRewards is a paid mutator transaction binding the contract method 0xab25ac5b.
//
// Solidity: function signRewards(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) SignRewards(opts *bind.TransactOpts, _rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte, _signature IFlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "signRewards", _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash, _signature)
}

// SignRewards is a paid mutator transaction binding the contract method 0xab25ac5b.
//
// Solidity: function signRewards(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerSession) SignRewards(_rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte, _signature IFlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SignRewards(&_FlareSystemManager.TransactOpts, _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash, _signature)
}

// SignRewards is a paid mutator transaction binding the contract method 0xab25ac5b.
//
// Solidity: function signRewards(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) SignRewards(_rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte, _signature IFlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SignRewards(&_FlareSystemManager.TransactOpts, _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash, _signature)
}

// SignUptimeVote is a paid mutator transaction binding the contract method 0xdc5a4225.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) SignUptimeVote(opts *bind.TransactOpts, _rewardEpochId *big.Int, _uptimeVoteHash [32]byte, _signature IFlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "signUptimeVote", _rewardEpochId, _uptimeVoteHash, _signature)
}

// SignUptimeVote is a paid mutator transaction binding the contract method 0xdc5a4225.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerSession) SignUptimeVote(_rewardEpochId *big.Int, _uptimeVoteHash [32]byte, _signature IFlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SignUptimeVote(&_FlareSystemManager.TransactOpts, _rewardEpochId, _uptimeVoteHash, _signature)
}

// SignUptimeVote is a paid mutator transaction binding the contract method 0xdc5a4225.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) SignUptimeVote(_rewardEpochId *big.Int, _uptimeVoteHash [32]byte, _signature IFlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SignUptimeVote(&_FlareSystemManager.TransactOpts, _rewardEpochId, _uptimeVoteHash, _signature)
}

// SubmitUptimeVote is a paid mutator transaction binding the contract method 0x9dd6850f.
//
// Solidity: function submitUptimeVote(uint24 _rewardEpochId, bytes20[] _nodeIds, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) SubmitUptimeVote(opts *bind.TransactOpts, _rewardEpochId *big.Int, _nodeIds [][20]byte, _signature IFlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "submitUptimeVote", _rewardEpochId, _nodeIds, _signature)
}

// SubmitUptimeVote is a paid mutator transaction binding the contract method 0x9dd6850f.
//
// Solidity: function submitUptimeVote(uint24 _rewardEpochId, bytes20[] _nodeIds, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerSession) SubmitUptimeVote(_rewardEpochId *big.Int, _nodeIds [][20]byte, _signature IFlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SubmitUptimeVote(&_FlareSystemManager.TransactOpts, _rewardEpochId, _nodeIds, _signature)
}

// SubmitUptimeVote is a paid mutator transaction binding the contract method 0x9dd6850f.
//
// Solidity: function submitUptimeVote(uint24 _rewardEpochId, bytes20[] _nodeIds, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) SubmitUptimeVote(_rewardEpochId *big.Int, _nodeIds [][20]byte, _signature IFlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SubmitUptimeVote(&_FlareSystemManager.TransactOpts, _rewardEpochId, _nodeIds, _signature)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FlareSystemManager *FlareSystemManagerSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SwitchToProductionMode(&_FlareSystemManager.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SwitchToProductionMode(&_FlareSystemManager.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FlareSystemManager *FlareSystemManagerSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.UpdateContractAddresses(&_FlareSystemManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.UpdateContractAddresses(&_FlareSystemManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x8fbaf860.
//
// Solidity: function updateSettings((uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint24,uint16,uint32) _settings) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) UpdateSettings(opts *bind.TransactOpts, _settings FlareSystemManagerSettings) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "updateSettings", _settings)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x8fbaf860.
//
// Solidity: function updateSettings((uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint24,uint16,uint32) _settings) returns()
func (_FlareSystemManager *FlareSystemManagerSession) UpdateSettings(_settings FlareSystemManagerSettings) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.UpdateSettings(&_FlareSystemManager.TransactOpts, _settings)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x8fbaf860.
//
// Solidity: function updateSettings((uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint24,uint16,uint32) _settings) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) UpdateSettings(_settings FlareSystemManagerSettings) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.UpdateSettings(&_FlareSystemManager.TransactOpts, _settings)
}

// FlareSystemManagerClosingExpiredRewardEpochFailedIterator is returned from FilterClosingExpiredRewardEpochFailed and is used to iterate over the raw logs and unpacked data for ClosingExpiredRewardEpochFailed events raised by the FlareSystemManager contract.
type FlareSystemManagerClosingExpiredRewardEpochFailedIterator struct {
	Event *FlareSystemManagerClosingExpiredRewardEpochFailed // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerClosingExpiredRewardEpochFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerClosingExpiredRewardEpochFailed)
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
		it.Event = new(FlareSystemManagerClosingExpiredRewardEpochFailed)
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
func (it *FlareSystemManagerClosingExpiredRewardEpochFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerClosingExpiredRewardEpochFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerClosingExpiredRewardEpochFailed represents a ClosingExpiredRewardEpochFailed event raised by the FlareSystemManager contract.
type FlareSystemManagerClosingExpiredRewardEpochFailed struct {
	RewardEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterClosingExpiredRewardEpochFailed is a free log retrieval operation binding the contract event 0xc0cded1f60001401da804c8a7703c1e8dc60521fca0f0f9853e2f1984b5410ba.
//
// Solidity: event ClosingExpiredRewardEpochFailed(uint24 rewardEpochId)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterClosingExpiredRewardEpochFailed(opts *bind.FilterOpts) (*FlareSystemManagerClosingExpiredRewardEpochFailedIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "ClosingExpiredRewardEpochFailed")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerClosingExpiredRewardEpochFailedIterator{contract: _FlareSystemManager.contract, event: "ClosingExpiredRewardEpochFailed", logs: logs, sub: sub}, nil
}

// WatchClosingExpiredRewardEpochFailed is a free log subscription operation binding the contract event 0xc0cded1f60001401da804c8a7703c1e8dc60521fca0f0f9853e2f1984b5410ba.
//
// Solidity: event ClosingExpiredRewardEpochFailed(uint24 rewardEpochId)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchClosingExpiredRewardEpochFailed(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerClosingExpiredRewardEpochFailed) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "ClosingExpiredRewardEpochFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerClosingExpiredRewardEpochFailed)
				if err := _FlareSystemManager.contract.UnpackLog(event, "ClosingExpiredRewardEpochFailed", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseClosingExpiredRewardEpochFailed(log types.Log) (*FlareSystemManagerClosingExpiredRewardEpochFailed, error) {
	event := new(FlareSystemManagerClosingExpiredRewardEpochFailed)
	if err := _FlareSystemManager.contract.UnpackLog(event, "ClosingExpiredRewardEpochFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the FlareSystemManager contract.
type FlareSystemManagerGovernanceCallTimelockedIterator struct {
	Event *FlareSystemManagerGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerGovernanceCallTimelocked)
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
		it.Event = new(FlareSystemManagerGovernanceCallTimelocked)
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
func (it *FlareSystemManagerGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FlareSystemManager contract.
type FlareSystemManagerGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*FlareSystemManagerGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerGovernanceCallTimelockedIterator{contract: _FlareSystemManager.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerGovernanceCallTimelocked)
				if err := _FlareSystemManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseGovernanceCallTimelocked(log types.Log) (*FlareSystemManagerGovernanceCallTimelocked, error) {
	event := new(FlareSystemManagerGovernanceCallTimelocked)
	if err := _FlareSystemManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the FlareSystemManager contract.
type FlareSystemManagerGovernanceInitialisedIterator struct {
	Event *FlareSystemManagerGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerGovernanceInitialised)
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
		it.Event = new(FlareSystemManagerGovernanceInitialised)
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
func (it *FlareSystemManagerGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerGovernanceInitialised represents a GovernanceInitialised event raised by the FlareSystemManager contract.
type FlareSystemManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*FlareSystemManagerGovernanceInitialisedIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerGovernanceInitialisedIterator{contract: _FlareSystemManager.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerGovernanceInitialised)
				if err := _FlareSystemManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseGovernanceInitialised(log types.Log) (*FlareSystemManagerGovernanceInitialised, error) {
	event := new(FlareSystemManagerGovernanceInitialised)
	if err := _FlareSystemManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the FlareSystemManager contract.
type FlareSystemManagerGovernedProductionModeEnteredIterator struct {
	Event *FlareSystemManagerGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerGovernedProductionModeEntered)
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
		it.Event = new(FlareSystemManagerGovernedProductionModeEntered)
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
func (it *FlareSystemManagerGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FlareSystemManager contract.
type FlareSystemManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*FlareSystemManagerGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerGovernedProductionModeEnteredIterator{contract: _FlareSystemManager.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerGovernedProductionModeEntered)
				if err := _FlareSystemManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseGovernedProductionModeEntered(log types.Log) (*FlareSystemManagerGovernedProductionModeEntered, error) {
	event := new(FlareSystemManagerGovernedProductionModeEntered)
	if err := _FlareSystemManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerRandomAcquisitionStartedIterator is returned from FilterRandomAcquisitionStarted and is used to iterate over the raw logs and unpacked data for RandomAcquisitionStarted events raised by the FlareSystemManager contract.
type FlareSystemManagerRandomAcquisitionStartedIterator struct {
	Event *FlareSystemManagerRandomAcquisitionStarted // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerRandomAcquisitionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerRandomAcquisitionStarted)
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
		it.Event = new(FlareSystemManagerRandomAcquisitionStarted)
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
func (it *FlareSystemManagerRandomAcquisitionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerRandomAcquisitionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerRandomAcquisitionStarted represents a RandomAcquisitionStarted event raised by the FlareSystemManager contract.
type FlareSystemManagerRandomAcquisitionStarted struct {
	RewardEpochId *big.Int
	Timestamp     uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRandomAcquisitionStarted is a free log retrieval operation binding the contract event 0xf9991783e5e480e42d9a54d3f35f4321857f8f0ebeb3742d326dce28b1126708.
//
// Solidity: event RandomAcquisitionStarted(uint24 indexed rewardEpochId, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterRandomAcquisitionStarted(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*FlareSystemManagerRandomAcquisitionStartedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "RandomAcquisitionStarted", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerRandomAcquisitionStartedIterator{contract: _FlareSystemManager.contract, event: "RandomAcquisitionStarted", logs: logs, sub: sub}, nil
}

// WatchRandomAcquisitionStarted is a free log subscription operation binding the contract event 0xf9991783e5e480e42d9a54d3f35f4321857f8f0ebeb3742d326dce28b1126708.
//
// Solidity: event RandomAcquisitionStarted(uint24 indexed rewardEpochId, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchRandomAcquisitionStarted(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerRandomAcquisitionStarted, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "RandomAcquisitionStarted", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerRandomAcquisitionStarted)
				if err := _FlareSystemManager.contract.UnpackLog(event, "RandomAcquisitionStarted", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseRandomAcquisitionStarted(log types.Log) (*FlareSystemManagerRandomAcquisitionStarted, error) {
	event := new(FlareSystemManagerRandomAcquisitionStarted)
	if err := _FlareSystemManager.contract.UnpackLog(event, "RandomAcquisitionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerRewardEpochStartedIterator is returned from FilterRewardEpochStarted and is used to iterate over the raw logs and unpacked data for RewardEpochStarted events raised by the FlareSystemManager contract.
type FlareSystemManagerRewardEpochStartedIterator struct {
	Event *FlareSystemManagerRewardEpochStarted // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerRewardEpochStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerRewardEpochStarted)
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
		it.Event = new(FlareSystemManagerRewardEpochStarted)
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
func (it *FlareSystemManagerRewardEpochStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerRewardEpochStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerRewardEpochStarted represents a RewardEpochStarted event raised by the FlareSystemManager contract.
type FlareSystemManagerRewardEpochStarted struct {
	RewardEpochId      *big.Int
	StartVotingRoundId uint32
	Timestamp          uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewardEpochStarted is a free log retrieval operation binding the contract event 0x4abb62ab1e4c42a11b90e4e45b92af1274f74cc634b759518e8c99e000d8be6d.
//
// Solidity: event RewardEpochStarted(uint24 indexed rewardEpochId, uint32 startVotingRoundId, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterRewardEpochStarted(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*FlareSystemManagerRewardEpochStartedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "RewardEpochStarted", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerRewardEpochStartedIterator{contract: _FlareSystemManager.contract, event: "RewardEpochStarted", logs: logs, sub: sub}, nil
}

// WatchRewardEpochStarted is a free log subscription operation binding the contract event 0x4abb62ab1e4c42a11b90e4e45b92af1274f74cc634b759518e8c99e000d8be6d.
//
// Solidity: event RewardEpochStarted(uint24 indexed rewardEpochId, uint32 startVotingRoundId, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchRewardEpochStarted(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerRewardEpochStarted, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "RewardEpochStarted", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerRewardEpochStarted)
				if err := _FlareSystemManager.contract.UnpackLog(event, "RewardEpochStarted", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseRewardEpochStarted(log types.Log) (*FlareSystemManagerRewardEpochStarted, error) {
	event := new(FlareSystemManagerRewardEpochStarted)
	if err := _FlareSystemManager.contract.UnpackLog(event, "RewardEpochStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerRewardsSignedIterator is returned from FilterRewardsSigned and is used to iterate over the raw logs and unpacked data for RewardsSigned events raised by the FlareSystemManager contract.
type FlareSystemManagerRewardsSignedIterator struct {
	Event *FlareSystemManagerRewardsSigned // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerRewardsSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerRewardsSigned)
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
		it.Event = new(FlareSystemManagerRewardsSigned)
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
func (it *FlareSystemManagerRewardsSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerRewardsSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerRewardsSigned represents a RewardsSigned event raised by the FlareSystemManager contract.
type FlareSystemManagerRewardsSigned struct {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterRewardsSigned(opts *bind.FilterOpts, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (*FlareSystemManagerRewardsSignedIterator, error) {

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

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "RewardsSigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerRewardsSignedIterator{contract: _FlareSystemManager.contract, event: "RewardsSigned", logs: logs, sub: sub}, nil
}

// WatchRewardsSigned is a free log subscription operation binding the contract event 0x80df246a47153e3604b0026fcf53ec85ea48178efd4f3cad24f9a1e1da2dd52a.
//
// Solidity: event RewardsSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 rewardsHash, uint256 noOfWeightBasedClaims, uint64 timestamp, bool thresholdReached)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchRewardsSigned(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerRewardsSigned, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "RewardsSigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerRewardsSigned)
				if err := _FlareSystemManager.contract.UnpackLog(event, "RewardsSigned", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseRewardsSigned(log types.Log) (*FlareSystemManagerRewardsSigned, error) {
	event := new(FlareSystemManagerRewardsSigned)
	if err := _FlareSystemManager.contract.UnpackLog(event, "RewardsSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerSettingCleanUpBlockNumberFailedIterator is returned from FilterSettingCleanUpBlockNumberFailed and is used to iterate over the raw logs and unpacked data for SettingCleanUpBlockNumberFailed events raised by the FlareSystemManager contract.
type FlareSystemManagerSettingCleanUpBlockNumberFailedIterator struct {
	Event *FlareSystemManagerSettingCleanUpBlockNumberFailed // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerSettingCleanUpBlockNumberFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerSettingCleanUpBlockNumberFailed)
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
		it.Event = new(FlareSystemManagerSettingCleanUpBlockNumberFailed)
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
func (it *FlareSystemManagerSettingCleanUpBlockNumberFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerSettingCleanUpBlockNumberFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerSettingCleanUpBlockNumberFailed represents a SettingCleanUpBlockNumberFailed event raised by the FlareSystemManager contract.
type FlareSystemManagerSettingCleanUpBlockNumberFailed struct {
	BlockNumber uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSettingCleanUpBlockNumberFailed is a free log retrieval operation binding the contract event 0xe9a7be2e41a6b0b36d253d56488c6844e611be2bffd8dd4b69b89a078f41fecc.
//
// Solidity: event SettingCleanUpBlockNumberFailed(uint64 blockNumber)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterSettingCleanUpBlockNumberFailed(opts *bind.FilterOpts) (*FlareSystemManagerSettingCleanUpBlockNumberFailedIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "SettingCleanUpBlockNumberFailed")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerSettingCleanUpBlockNumberFailedIterator{contract: _FlareSystemManager.contract, event: "SettingCleanUpBlockNumberFailed", logs: logs, sub: sub}, nil
}

// WatchSettingCleanUpBlockNumberFailed is a free log subscription operation binding the contract event 0xe9a7be2e41a6b0b36d253d56488c6844e611be2bffd8dd4b69b89a078f41fecc.
//
// Solidity: event SettingCleanUpBlockNumberFailed(uint64 blockNumber)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchSettingCleanUpBlockNumberFailed(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerSettingCleanUpBlockNumberFailed) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "SettingCleanUpBlockNumberFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerSettingCleanUpBlockNumberFailed)
				if err := _FlareSystemManager.contract.UnpackLog(event, "SettingCleanUpBlockNumberFailed", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseSettingCleanUpBlockNumberFailed(log types.Log) (*FlareSystemManagerSettingCleanUpBlockNumberFailed, error) {
	event := new(FlareSystemManagerSettingCleanUpBlockNumberFailed)
	if err := _FlareSystemManager.contract.UnpackLog(event, "SettingCleanUpBlockNumberFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerSigningPolicySignedIterator is returned from FilterSigningPolicySigned and is used to iterate over the raw logs and unpacked data for SigningPolicySigned events raised by the FlareSystemManager contract.
type FlareSystemManagerSigningPolicySignedIterator struct {
	Event *FlareSystemManagerSigningPolicySigned // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerSigningPolicySignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerSigningPolicySigned)
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
		it.Event = new(FlareSystemManagerSigningPolicySigned)
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
func (it *FlareSystemManagerSigningPolicySignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerSigningPolicySignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerSigningPolicySigned represents a SigningPolicySigned event raised by the FlareSystemManager contract.
type FlareSystemManagerSigningPolicySigned struct {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterSigningPolicySigned(opts *bind.FilterOpts, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (*FlareSystemManagerSigningPolicySignedIterator, error) {

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

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "SigningPolicySigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerSigningPolicySignedIterator{contract: _FlareSystemManager.contract, event: "SigningPolicySigned", logs: logs, sub: sub}, nil
}

// WatchSigningPolicySigned is a free log subscription operation binding the contract event 0x154b0214ae62d8a5548c1eac25fabd87c38b04932a217732e1022f3118da67f3.
//
// Solidity: event SigningPolicySigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, uint64 timestamp, bool thresholdReached)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchSigningPolicySigned(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerSigningPolicySigned, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "SigningPolicySigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerSigningPolicySigned)
				if err := _FlareSystemManager.contract.UnpackLog(event, "SigningPolicySigned", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseSigningPolicySigned(log types.Log) (*FlareSystemManagerSigningPolicySigned, error) {
	event := new(FlareSystemManagerSigningPolicySigned)
	if err := _FlareSystemManager.contract.UnpackLog(event, "SigningPolicySigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerSingUptimeVoteEnabledIterator is returned from FilterSingUptimeVoteEnabled and is used to iterate over the raw logs and unpacked data for SingUptimeVoteEnabled events raised by the FlareSystemManager contract.
type FlareSystemManagerSingUptimeVoteEnabledIterator struct {
	Event *FlareSystemManagerSingUptimeVoteEnabled // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerSingUptimeVoteEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerSingUptimeVoteEnabled)
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
		it.Event = new(FlareSystemManagerSingUptimeVoteEnabled)
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
func (it *FlareSystemManagerSingUptimeVoteEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerSingUptimeVoteEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerSingUptimeVoteEnabled represents a SingUptimeVoteEnabled event raised by the FlareSystemManager contract.
type FlareSystemManagerSingUptimeVoteEnabled struct {
	RewardEpochId *big.Int
	Timestamp     uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSingUptimeVoteEnabled is a free log retrieval operation binding the contract event 0xf03263754f56959acf7e09464f5d304491e4e5389363e5e12d2ce2ad492632f4.
//
// Solidity: event SingUptimeVoteEnabled(uint24 indexed rewardEpochId, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterSingUptimeVoteEnabled(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*FlareSystemManagerSingUptimeVoteEnabledIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "SingUptimeVoteEnabled", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerSingUptimeVoteEnabledIterator{contract: _FlareSystemManager.contract, event: "SingUptimeVoteEnabled", logs: logs, sub: sub}, nil
}

// WatchSingUptimeVoteEnabled is a free log subscription operation binding the contract event 0xf03263754f56959acf7e09464f5d304491e4e5389363e5e12d2ce2ad492632f4.
//
// Solidity: event SingUptimeVoteEnabled(uint24 indexed rewardEpochId, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchSingUptimeVoteEnabled(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerSingUptimeVoteEnabled, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "SingUptimeVoteEnabled", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerSingUptimeVoteEnabled)
				if err := _FlareSystemManager.contract.UnpackLog(event, "SingUptimeVoteEnabled", log); err != nil {
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

// ParseSingUptimeVoteEnabled is a log parse operation binding the contract event 0xf03263754f56959acf7e09464f5d304491e4e5389363e5e12d2ce2ad492632f4.
//
// Solidity: event SingUptimeVoteEnabled(uint24 indexed rewardEpochId, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseSingUptimeVoteEnabled(log types.Log) (*FlareSystemManagerSingUptimeVoteEnabled, error) {
	event := new(FlareSystemManagerSingUptimeVoteEnabled)
	if err := _FlareSystemManager.contract.UnpackLog(event, "SingUptimeVoteEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the FlareSystemManager contract.
type FlareSystemManagerTimelockedGovernanceCallCanceledIterator struct {
	Event *FlareSystemManagerTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerTimelockedGovernanceCallCanceled)
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
		it.Event = new(FlareSystemManagerTimelockedGovernanceCallCanceled)
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
func (it *FlareSystemManagerTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FlareSystemManager contract.
type FlareSystemManagerTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*FlareSystemManagerTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerTimelockedGovernanceCallCanceledIterator{contract: _FlareSystemManager.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerTimelockedGovernanceCallCanceled)
				if err := _FlareSystemManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*FlareSystemManagerTimelockedGovernanceCallCanceled, error) {
	event := new(FlareSystemManagerTimelockedGovernanceCallCanceled)
	if err := _FlareSystemManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the FlareSystemManager contract.
type FlareSystemManagerTimelockedGovernanceCallExecutedIterator struct {
	Event *FlareSystemManagerTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerTimelockedGovernanceCallExecuted)
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
		it.Event = new(FlareSystemManagerTimelockedGovernanceCallExecuted)
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
func (it *FlareSystemManagerTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FlareSystemManager contract.
type FlareSystemManagerTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*FlareSystemManagerTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerTimelockedGovernanceCallExecutedIterator{contract: _FlareSystemManager.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerTimelockedGovernanceCallExecuted)
				if err := _FlareSystemManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*FlareSystemManagerTimelockedGovernanceCallExecuted, error) {
	event := new(FlareSystemManagerTimelockedGovernanceCallExecuted)
	if err := _FlareSystemManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerTriggeringVoterRegistrationFailedIterator is returned from FilterTriggeringVoterRegistrationFailed and is used to iterate over the raw logs and unpacked data for TriggeringVoterRegistrationFailed events raised by the FlareSystemManager contract.
type FlareSystemManagerTriggeringVoterRegistrationFailedIterator struct {
	Event *FlareSystemManagerTriggeringVoterRegistrationFailed // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerTriggeringVoterRegistrationFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerTriggeringVoterRegistrationFailed)
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
		it.Event = new(FlareSystemManagerTriggeringVoterRegistrationFailed)
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
func (it *FlareSystemManagerTriggeringVoterRegistrationFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerTriggeringVoterRegistrationFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerTriggeringVoterRegistrationFailed represents a TriggeringVoterRegistrationFailed event raised by the FlareSystemManager contract.
type FlareSystemManagerTriggeringVoterRegistrationFailed struct {
	RewardEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTriggeringVoterRegistrationFailed is a free log retrieval operation binding the contract event 0x449d255b9c487823db86822a857f218d40682abada12acee2483788dc2fa975a.
//
// Solidity: event TriggeringVoterRegistrationFailed(uint24 rewardEpochId)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterTriggeringVoterRegistrationFailed(opts *bind.FilterOpts) (*FlareSystemManagerTriggeringVoterRegistrationFailedIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "TriggeringVoterRegistrationFailed")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerTriggeringVoterRegistrationFailedIterator{contract: _FlareSystemManager.contract, event: "TriggeringVoterRegistrationFailed", logs: logs, sub: sub}, nil
}

// WatchTriggeringVoterRegistrationFailed is a free log subscription operation binding the contract event 0x449d255b9c487823db86822a857f218d40682abada12acee2483788dc2fa975a.
//
// Solidity: event TriggeringVoterRegistrationFailed(uint24 rewardEpochId)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchTriggeringVoterRegistrationFailed(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerTriggeringVoterRegistrationFailed) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "TriggeringVoterRegistrationFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerTriggeringVoterRegistrationFailed)
				if err := _FlareSystemManager.contract.UnpackLog(event, "TriggeringVoterRegistrationFailed", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseTriggeringVoterRegistrationFailed(log types.Log) (*FlareSystemManagerTriggeringVoterRegistrationFailed, error) {
	event := new(FlareSystemManagerTriggeringVoterRegistrationFailed)
	if err := _FlareSystemManager.contract.UnpackLog(event, "TriggeringVoterRegistrationFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerUptimeVoteSignedIterator is returned from FilterUptimeVoteSigned and is used to iterate over the raw logs and unpacked data for UptimeVoteSigned events raised by the FlareSystemManager contract.
type FlareSystemManagerUptimeVoteSignedIterator struct {
	Event *FlareSystemManagerUptimeVoteSigned // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerUptimeVoteSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerUptimeVoteSigned)
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
		it.Event = new(FlareSystemManagerUptimeVoteSigned)
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
func (it *FlareSystemManagerUptimeVoteSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerUptimeVoteSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerUptimeVoteSigned represents a UptimeVoteSigned event raised by the FlareSystemManager contract.
type FlareSystemManagerUptimeVoteSigned struct {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterUptimeVoteSigned(opts *bind.FilterOpts, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (*FlareSystemManagerUptimeVoteSignedIterator, error) {

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

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "UptimeVoteSigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerUptimeVoteSignedIterator{contract: _FlareSystemManager.contract, event: "UptimeVoteSigned", logs: logs, sub: sub}, nil
}

// WatchUptimeVoteSigned is a free log subscription operation binding the contract event 0x5506337d1266599f8b64675a1c8321701657ca2f2f70be0e0c58302b6c22e797.
//
// Solidity: event UptimeVoteSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 uptimeVoteHash, uint64 timestamp, bool thresholdReached)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchUptimeVoteSigned(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerUptimeVoteSigned, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "UptimeVoteSigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerUptimeVoteSigned)
				if err := _FlareSystemManager.contract.UnpackLog(event, "UptimeVoteSigned", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseUptimeVoteSigned(log types.Log) (*FlareSystemManagerUptimeVoteSigned, error) {
	event := new(FlareSystemManagerUptimeVoteSigned)
	if err := _FlareSystemManager.contract.UnpackLog(event, "UptimeVoteSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerUptimeVoteSubmittedIterator is returned from FilterUptimeVoteSubmitted and is used to iterate over the raw logs and unpacked data for UptimeVoteSubmitted events raised by the FlareSystemManager contract.
type FlareSystemManagerUptimeVoteSubmittedIterator struct {
	Event *FlareSystemManagerUptimeVoteSubmitted // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerUptimeVoteSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerUptimeVoteSubmitted)
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
		it.Event = new(FlareSystemManagerUptimeVoteSubmitted)
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
func (it *FlareSystemManagerUptimeVoteSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerUptimeVoteSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerUptimeVoteSubmitted represents a UptimeVoteSubmitted event raised by the FlareSystemManager contract.
type FlareSystemManagerUptimeVoteSubmitted struct {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterUptimeVoteSubmitted(opts *bind.FilterOpts, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (*FlareSystemManagerUptimeVoteSubmittedIterator, error) {

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

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "UptimeVoteSubmitted", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerUptimeVoteSubmittedIterator{contract: _FlareSystemManager.contract, event: "UptimeVoteSubmitted", logs: logs, sub: sub}, nil
}

// WatchUptimeVoteSubmitted is a free log subscription operation binding the contract event 0xed370d61eb315e1d46d979894585530b99f94dab64c0d40366685aebe39e3db0.
//
// Solidity: event UptimeVoteSubmitted(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes20[] nodeIds, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchUptimeVoteSubmitted(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerUptimeVoteSubmitted, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "UptimeVoteSubmitted", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerUptimeVoteSubmitted)
				if err := _FlareSystemManager.contract.UnpackLog(event, "UptimeVoteSubmitted", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseUptimeVoteSubmitted(log types.Log) (*FlareSystemManagerUptimeVoteSubmitted, error) {
	event := new(FlareSystemManagerUptimeVoteSubmitted)
	if err := _FlareSystemManager.contract.UnpackLog(event, "UptimeVoteSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerVotePowerBlockSelectedIterator is returned from FilterVotePowerBlockSelected and is used to iterate over the raw logs and unpacked data for VotePowerBlockSelected events raised by the FlareSystemManager contract.
type FlareSystemManagerVotePowerBlockSelectedIterator struct {
	Event *FlareSystemManagerVotePowerBlockSelected // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerVotePowerBlockSelectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerVotePowerBlockSelected)
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
		it.Event = new(FlareSystemManagerVotePowerBlockSelected)
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
func (it *FlareSystemManagerVotePowerBlockSelectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerVotePowerBlockSelectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerVotePowerBlockSelected represents a VotePowerBlockSelected event raised by the FlareSystemManager contract.
type FlareSystemManagerVotePowerBlockSelected struct {
	RewardEpochId  *big.Int
	VotePowerBlock uint64
	Timestamp      uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVotePowerBlockSelected is a free log retrieval operation binding the contract event 0xf21722dbe044a7cea0f6d81c871cae750971e36c9dd10999e46f2b46f26ac7ff.
//
// Solidity: event VotePowerBlockSelected(uint24 indexed rewardEpochId, uint64 votePowerBlock, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterVotePowerBlockSelected(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*FlareSystemManagerVotePowerBlockSelectedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "VotePowerBlockSelected", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerVotePowerBlockSelectedIterator{contract: _FlareSystemManager.contract, event: "VotePowerBlockSelected", logs: logs, sub: sub}, nil
}

// WatchVotePowerBlockSelected is a free log subscription operation binding the contract event 0xf21722dbe044a7cea0f6d81c871cae750971e36c9dd10999e46f2b46f26ac7ff.
//
// Solidity: event VotePowerBlockSelected(uint24 indexed rewardEpochId, uint64 votePowerBlock, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchVotePowerBlockSelected(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerVotePowerBlockSelected, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "VotePowerBlockSelected", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerVotePowerBlockSelected)
				if err := _FlareSystemManager.contract.UnpackLog(event, "VotePowerBlockSelected", log); err != nil {
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
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseVotePowerBlockSelected(log types.Log) (*FlareSystemManagerVotePowerBlockSelected, error) {
	event := new(FlareSystemManagerVotePowerBlockSelected)
	if err := _FlareSystemManager.contract.UnpackLog(event, "VotePowerBlockSelected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
