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

// FlareSystemManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type FlareSystemManagerSettings struct {
	FirstVotingRoundStartTs                          uint64
	VotingEpochDurationSeconds                       uint64
	FirstRewardEpochStartVotingRoundId               uint64
	RewardEpochDurationInVotingEpochs                uint64
	NewSigningPolicyInitializationStartSeconds       uint64
	NonPunishableRandomAcquisitionMinDurationSeconds uint64
	NonPunishableRandomAcquisitionMinDurationBlocks  uint64
	VoterRegistrationMinDurationSeconds              uint64
	VoterRegistrationMinDurationBlocks               uint64
	NonPunishableSigningPolicySignMinDurationSeconds uint64
	NonPunishableSigningPolicySignMinDurationBlocks  uint64
	SigningPolicyThresholdPPM                        uint64
	SigningPolicyMinNumberOfVoters                   uint64
}

// FlareSystemManagerSignature is an auto generated low-level Go binding around an user-defined struct.
type FlareSystemManagerSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// FlareSystemManagerMetaData contains all meta data concerning the FlareSystemManager contract.
var FlareSystemManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_flareDaemon\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"firstVotingRoundStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEpochDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstRewardEpochStartVotingRoundId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rewardEpochDurationInVotingEpochs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newSigningPolicyInitializationStartSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonPunishableRandomAcquisitionMinDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonPunishableRandomAcquisitionMinDurationBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"voterRegistrationMinDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"voterRegistrationMinDurationBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonPunishableSigningPolicySignMinDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonPunishableSigningPolicySignMinDurationBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"signingPolicyThresholdPPM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"signingPolicyMinNumberOfVoters\",\"type\":\"uint64\"}],\"internalType\":\"structFlareSystemManager.Settings\",\"name\":\"_settings\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_firstRandomAcquisitionNumberOfBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"_firstRewardEpochId\",\"type\":\"uint24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"RandomAcquisitionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signingAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rewardsHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noOfWeightBasedClaims\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"RewardsSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startVotingRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"threshold\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint16[]\",\"name\":\"weights\",\"type\":\"uint16[]\"}],\"name\":\"SigningPolicyInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signingAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"SigningPolicySigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signingAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"uptimeVoteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"UptimeVoteSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"votePowerBlock\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"VotePowerBlockSelected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_usePriceSubmitter\",\"type\":\"bool\"}],\"name\":\"changeRandomProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_signingPolicyThresholdPPM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicyMinNumberOfVoters\",\"type\":\"uint64\"}],\"name\":\"changeSigningPolicySettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daemonize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstVotingRoundStartTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareDaemon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentRandom\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRandomWithQuality\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentRandom\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_goodRandom\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRewardEpochId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"_currentRewardEpochId\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_rewardEpoch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_rewardOwner\",\"type\":\"address\"}],\"name\":\"getRewardsFeeBurnFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpoch\",\"type\":\"uint256\"}],\"name\":\"getVotePowerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_votePowerBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpoch\",\"type\":\"uint256\"}],\"name\":\"getVoterRegistrationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_votePowerBlock\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newSigningPolicyInitializationStartSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"noOfWeightBasedClaims\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonPunishableRandomAcquisitionMinDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonPunishableRandomAcquisitionMinDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonPunishableSigningPolicySignMinDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonPunishableSigningPolicySignMinDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceSubmitter\",\"outputs\":[{\"internalType\":\"contractIPriceSubmitter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relay\",\"outputs\":[{\"internalType\":\"contractRelay\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardEpochsStartTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardsHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"_newSigningPolicyHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structFlareSystemManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signNewSigningPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_noOfWeightBasedClaims\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_rewardsHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structFlareSystemManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"_uptimeVoteHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structFlareSystemManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signUptimeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingPolicyMinNumberOfVoters\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingPolicyThresholdPPM\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submission\",\"outputs\":[{\"internalType\":\"contractSubmission\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToFallbackMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uptimeVoteHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usePriceSubmitterAsRandomProvider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistrationMinDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistrationMinDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistry\",\"outputs\":[{\"internalType\":\"contractVoterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetCurrentRandom is a free data retrieval call binding the contract method 0xd89601fd.
//
// Solidity: function getCurrentRandom() view returns(uint256 _currentRandom)
func (_FlareSystemManager *FlareSystemManagerCaller) GetCurrentRandom(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getCurrentRandom")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRandom is a free data retrieval call binding the contract method 0xd89601fd.
//
// Solidity: function getCurrentRandom() view returns(uint256 _currentRandom)
func (_FlareSystemManager *FlareSystemManagerSession) GetCurrentRandom() (*big.Int, error) {
	return _FlareSystemManager.Contract.GetCurrentRandom(&_FlareSystemManager.CallOpts)
}

// GetCurrentRandom is a free data retrieval call binding the contract method 0xd89601fd.
//
// Solidity: function getCurrentRandom() view returns(uint256 _currentRandom)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetCurrentRandom() (*big.Int, error) {
	return _FlareSystemManager.Contract.GetCurrentRandom(&_FlareSystemManager.CallOpts)
}

// GetCurrentRandomWithQuality is a free data retrieval call binding the contract method 0xa978fb6b.
//
// Solidity: function getCurrentRandomWithQuality() view returns(uint256 _currentRandom, bool _goodRandom)
func (_FlareSystemManager *FlareSystemManagerCaller) GetCurrentRandomWithQuality(opts *bind.CallOpts) (struct {
	CurrentRandom *big.Int
	GoodRandom    bool
}, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getCurrentRandomWithQuality")

	outstruct := new(struct {
		CurrentRandom *big.Int
		GoodRandom    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentRandom = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GoodRandom = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetCurrentRandomWithQuality is a free data retrieval call binding the contract method 0xa978fb6b.
//
// Solidity: function getCurrentRandomWithQuality() view returns(uint256 _currentRandom, bool _goodRandom)
func (_FlareSystemManager *FlareSystemManagerSession) GetCurrentRandomWithQuality() (struct {
	CurrentRandom *big.Int
	GoodRandom    bool
}, error) {
	return _FlareSystemManager.Contract.GetCurrentRandomWithQuality(&_FlareSystemManager.CallOpts)
}

// GetCurrentRandomWithQuality is a free data retrieval call binding the contract method 0xa978fb6b.
//
// Solidity: function getCurrentRandomWithQuality() view returns(uint256 _currentRandom, bool _goodRandom)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetCurrentRandomWithQuality() (struct {
	CurrentRandom *big.Int
	GoodRandom    bool
}, error) {
	return _FlareSystemManager.Contract.GetCurrentRandomWithQuality(&_FlareSystemManager.CallOpts)
}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24 _currentRewardEpochId)
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
// Solidity: function getCurrentRewardEpochId() view returns(uint24 _currentRewardEpochId)
func (_FlareSystemManager *FlareSystemManagerSession) GetCurrentRewardEpochId() (*big.Int, error) {
	return _FlareSystemManager.Contract.GetCurrentRewardEpochId(&_FlareSystemManager.CallOpts)
}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24 _currentRewardEpochId)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetCurrentRewardEpochId() (*big.Int, error) {
	return _FlareSystemManager.Contract.GetCurrentRewardEpochId(&_FlareSystemManager.CallOpts)
}

// GetRewardsFeeBurnFactor is a free data retrieval call binding the contract method 0x31db9891.
//
// Solidity: function getRewardsFeeBurnFactor(uint64 _rewardEpoch, address _rewardOwner) view returns(uint256)
func (_FlareSystemManager *FlareSystemManagerCaller) GetRewardsFeeBurnFactor(opts *bind.CallOpts, _rewardEpoch uint64, _rewardOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getRewardsFeeBurnFactor", _rewardEpoch, _rewardOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardsFeeBurnFactor is a free data retrieval call binding the contract method 0x31db9891.
//
// Solidity: function getRewardsFeeBurnFactor(uint64 _rewardEpoch, address _rewardOwner) view returns(uint256)
func (_FlareSystemManager *FlareSystemManagerSession) GetRewardsFeeBurnFactor(_rewardEpoch uint64, _rewardOwner common.Address) (*big.Int, error) {
	return _FlareSystemManager.Contract.GetRewardsFeeBurnFactor(&_FlareSystemManager.CallOpts, _rewardEpoch, _rewardOwner)
}

// GetRewardsFeeBurnFactor is a free data retrieval call binding the contract method 0x31db9891.
//
// Solidity: function getRewardsFeeBurnFactor(uint64 _rewardEpoch, address _rewardOwner) view returns(uint256)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetRewardsFeeBurnFactor(_rewardEpoch uint64, _rewardOwner common.Address) (*big.Int, error) {
	return _FlareSystemManager.Contract.GetRewardsFeeBurnFactor(&_FlareSystemManager.CallOpts, _rewardEpoch, _rewardOwner)
}

// GetVotePowerBlock is a free data retrieval call binding the contract method 0xc2632216.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpoch) view returns(uint256 _votePowerBlock)
func (_FlareSystemManager *FlareSystemManagerCaller) GetVotePowerBlock(opts *bind.CallOpts, _rewardEpoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getVotePowerBlock", _rewardEpoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotePowerBlock is a free data retrieval call binding the contract method 0xc2632216.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpoch) view returns(uint256 _votePowerBlock)
func (_FlareSystemManager *FlareSystemManagerSession) GetVotePowerBlock(_rewardEpoch *big.Int) (*big.Int, error) {
	return _FlareSystemManager.Contract.GetVotePowerBlock(&_FlareSystemManager.CallOpts, _rewardEpoch)
}

// GetVotePowerBlock is a free data retrieval call binding the contract method 0xc2632216.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpoch) view returns(uint256 _votePowerBlock)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetVotePowerBlock(_rewardEpoch *big.Int) (*big.Int, error) {
	return _FlareSystemManager.Contract.GetVotePowerBlock(&_FlareSystemManager.CallOpts, _rewardEpoch)
}

// GetVoterRegistrationData is a free data retrieval call binding the contract method 0x1703a788.
//
// Solidity: function getVoterRegistrationData(uint256 _rewardEpoch) view returns(uint256 _votePowerBlock, bool _enabled)
func (_FlareSystemManager *FlareSystemManagerCaller) GetVoterRegistrationData(opts *bind.CallOpts, _rewardEpoch *big.Int) (struct {
	VotePowerBlock *big.Int
	Enabled        bool
}, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "getVoterRegistrationData", _rewardEpoch)

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
// Solidity: function getVoterRegistrationData(uint256 _rewardEpoch) view returns(uint256 _votePowerBlock, bool _enabled)
func (_FlareSystemManager *FlareSystemManagerSession) GetVoterRegistrationData(_rewardEpoch *big.Int) (struct {
	VotePowerBlock *big.Int
	Enabled        bool
}, error) {
	return _FlareSystemManager.Contract.GetVoterRegistrationData(&_FlareSystemManager.CallOpts, _rewardEpoch)
}

// GetVoterRegistrationData is a free data retrieval call binding the contract method 0x1703a788.
//
// Solidity: function getVoterRegistrationData(uint256 _rewardEpoch) view returns(uint256 _votePowerBlock, bool _enabled)
func (_FlareSystemManager *FlareSystemManagerCallerSession) GetVoterRegistrationData(_rewardEpoch *big.Int) (struct {
	VotePowerBlock *big.Int
	Enabled        bool
}, error) {
	return _FlareSystemManager.Contract.GetVoterRegistrationData(&_FlareSystemManager.CallOpts, _rewardEpoch)
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

// NonPunishableRandomAcquisitionMinDurationBlocks is a free data retrieval call binding the contract method 0x060afa37.
//
// Solidity: function nonPunishableRandomAcquisitionMinDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) NonPunishableRandomAcquisitionMinDurationBlocks(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "nonPunishableRandomAcquisitionMinDurationBlocks")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NonPunishableRandomAcquisitionMinDurationBlocks is a free data retrieval call binding the contract method 0x060afa37.
//
// Solidity: function nonPunishableRandomAcquisitionMinDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) NonPunishableRandomAcquisitionMinDurationBlocks() (uint64, error) {
	return _FlareSystemManager.Contract.NonPunishableRandomAcquisitionMinDurationBlocks(&_FlareSystemManager.CallOpts)
}

// NonPunishableRandomAcquisitionMinDurationBlocks is a free data retrieval call binding the contract method 0x060afa37.
//
// Solidity: function nonPunishableRandomAcquisitionMinDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) NonPunishableRandomAcquisitionMinDurationBlocks() (uint64, error) {
	return _FlareSystemManager.Contract.NonPunishableRandomAcquisitionMinDurationBlocks(&_FlareSystemManager.CallOpts)
}

// NonPunishableRandomAcquisitionMinDurationSeconds is a free data retrieval call binding the contract method 0xa4583262.
//
// Solidity: function nonPunishableRandomAcquisitionMinDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) NonPunishableRandomAcquisitionMinDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "nonPunishableRandomAcquisitionMinDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NonPunishableRandomAcquisitionMinDurationSeconds is a free data retrieval call binding the contract method 0xa4583262.
//
// Solidity: function nonPunishableRandomAcquisitionMinDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) NonPunishableRandomAcquisitionMinDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.NonPunishableRandomAcquisitionMinDurationSeconds(&_FlareSystemManager.CallOpts)
}

// NonPunishableRandomAcquisitionMinDurationSeconds is a free data retrieval call binding the contract method 0xa4583262.
//
// Solidity: function nonPunishableRandomAcquisitionMinDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) NonPunishableRandomAcquisitionMinDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.NonPunishableRandomAcquisitionMinDurationSeconds(&_FlareSystemManager.CallOpts)
}

// NonPunishableSigningPolicySignMinDurationBlocks is a free data retrieval call binding the contract method 0x47426bda.
//
// Solidity: function nonPunishableSigningPolicySignMinDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) NonPunishableSigningPolicySignMinDurationBlocks(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "nonPunishableSigningPolicySignMinDurationBlocks")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NonPunishableSigningPolicySignMinDurationBlocks is a free data retrieval call binding the contract method 0x47426bda.
//
// Solidity: function nonPunishableSigningPolicySignMinDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) NonPunishableSigningPolicySignMinDurationBlocks() (uint64, error) {
	return _FlareSystemManager.Contract.NonPunishableSigningPolicySignMinDurationBlocks(&_FlareSystemManager.CallOpts)
}

// NonPunishableSigningPolicySignMinDurationBlocks is a free data retrieval call binding the contract method 0x47426bda.
//
// Solidity: function nonPunishableSigningPolicySignMinDurationBlocks() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) NonPunishableSigningPolicySignMinDurationBlocks() (uint64, error) {
	return _FlareSystemManager.Contract.NonPunishableSigningPolicySignMinDurationBlocks(&_FlareSystemManager.CallOpts)
}

// NonPunishableSigningPolicySignMinDurationSeconds is a free data retrieval call binding the contract method 0x702a6514.
//
// Solidity: function nonPunishableSigningPolicySignMinDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) NonPunishableSigningPolicySignMinDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "nonPunishableSigningPolicySignMinDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NonPunishableSigningPolicySignMinDurationSeconds is a free data retrieval call binding the contract method 0x702a6514.
//
// Solidity: function nonPunishableSigningPolicySignMinDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) NonPunishableSigningPolicySignMinDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.NonPunishableSigningPolicySignMinDurationSeconds(&_FlareSystemManager.CallOpts)
}

// NonPunishableSigningPolicySignMinDurationSeconds is a free data retrieval call binding the contract method 0x702a6514.
//
// Solidity: function nonPunishableSigningPolicySignMinDurationSeconds() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) NonPunishableSigningPolicySignMinDurationSeconds() (uint64, error) {
	return _FlareSystemManager.Contract.NonPunishableSigningPolicySignMinDurationSeconds(&_FlareSystemManager.CallOpts)
}

// PriceSubmitter is a free data retrieval call binding the contract method 0xf937d6ad.
//
// Solidity: function priceSubmitter() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCaller) PriceSubmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "priceSubmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceSubmitter is a free data retrieval call binding the contract method 0xf937d6ad.
//
// Solidity: function priceSubmitter() view returns(address)
func (_FlareSystemManager *FlareSystemManagerSession) PriceSubmitter() (common.Address, error) {
	return _FlareSystemManager.Contract.PriceSubmitter(&_FlareSystemManager.CallOpts)
}

// PriceSubmitter is a free data retrieval call binding the contract method 0xf937d6ad.
//
// Solidity: function priceSubmitter() view returns(address)
func (_FlareSystemManager *FlareSystemManagerCallerSession) PriceSubmitter() (common.Address, error) {
	return _FlareSystemManager.Contract.PriceSubmitter(&_FlareSystemManager.CallOpts)
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

// RewardEpochsStartTs is a free data retrieval call binding the contract method 0xa578f55b.
//
// Solidity: function rewardEpochsStartTs() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) RewardEpochsStartTs(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "rewardEpochsStartTs")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RewardEpochsStartTs is a free data retrieval call binding the contract method 0xa578f55b.
//
// Solidity: function rewardEpochsStartTs() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) RewardEpochsStartTs() (uint64, error) {
	return _FlareSystemManager.Contract.RewardEpochsStartTs(&_FlareSystemManager.CallOpts)
}

// RewardEpochsStartTs is a free data retrieval call binding the contract method 0xa578f55b.
//
// Solidity: function rewardEpochsStartTs() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) RewardEpochsStartTs() (uint64, error) {
	return _FlareSystemManager.Contract.RewardEpochsStartTs(&_FlareSystemManager.CallOpts)
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
// Solidity: function signingPolicyMinNumberOfVoters() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) SigningPolicyMinNumberOfVoters(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "signingPolicyMinNumberOfVoters")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SigningPolicyMinNumberOfVoters is a free data retrieval call binding the contract method 0x2e3645f8.
//
// Solidity: function signingPolicyMinNumberOfVoters() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) SigningPolicyMinNumberOfVoters() (uint64, error) {
	return _FlareSystemManager.Contract.SigningPolicyMinNumberOfVoters(&_FlareSystemManager.CallOpts)
}

// SigningPolicyMinNumberOfVoters is a free data retrieval call binding the contract method 0x2e3645f8.
//
// Solidity: function signingPolicyMinNumberOfVoters() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) SigningPolicyMinNumberOfVoters() (uint64, error) {
	return _FlareSystemManager.Contract.SigningPolicyMinNumberOfVoters(&_FlareSystemManager.CallOpts)
}

// SigningPolicyThresholdPPM is a free data retrieval call binding the contract method 0xf21d6304.
//
// Solidity: function signingPolicyThresholdPPM() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCaller) SigningPolicyThresholdPPM(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "signingPolicyThresholdPPM")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SigningPolicyThresholdPPM is a free data retrieval call binding the contract method 0xf21d6304.
//
// Solidity: function signingPolicyThresholdPPM() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerSession) SigningPolicyThresholdPPM() (uint64, error) {
	return _FlareSystemManager.Contract.SigningPolicyThresholdPPM(&_FlareSystemManager.CallOpts)
}

// SigningPolicyThresholdPPM is a free data retrieval call binding the contract method 0xf21d6304.
//
// Solidity: function signingPolicyThresholdPPM() view returns(uint64)
func (_FlareSystemManager *FlareSystemManagerCallerSession) SigningPolicyThresholdPPM() (uint64, error) {
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

// SwitchToFallbackMode is a free data retrieval call binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() pure returns(bool)
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
// Solidity: function switchToFallbackMode() pure returns(bool)
func (_FlareSystemManager *FlareSystemManagerSession) SwitchToFallbackMode() (bool, error) {
	return _FlareSystemManager.Contract.SwitchToFallbackMode(&_FlareSystemManager.CallOpts)
}

// SwitchToFallbackMode is a free data retrieval call binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() pure returns(bool)
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

// UsePriceSubmitterAsRandomProvider is a free data retrieval call binding the contract method 0x5c4617b5.
//
// Solidity: function usePriceSubmitterAsRandomProvider() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCaller) UsePriceSubmitterAsRandomProvider(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlareSystemManager.contract.Call(opts, &out, "usePriceSubmitterAsRandomProvider")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsePriceSubmitterAsRandomProvider is a free data retrieval call binding the contract method 0x5c4617b5.
//
// Solidity: function usePriceSubmitterAsRandomProvider() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerSession) UsePriceSubmitterAsRandomProvider() (bool, error) {
	return _FlareSystemManager.Contract.UsePriceSubmitterAsRandomProvider(&_FlareSystemManager.CallOpts)
}

// UsePriceSubmitterAsRandomProvider is a free data retrieval call binding the contract method 0x5c4617b5.
//
// Solidity: function usePriceSubmitterAsRandomProvider() view returns(bool)
func (_FlareSystemManager *FlareSystemManagerCallerSession) UsePriceSubmitterAsRandomProvider() (bool, error) {
	return _FlareSystemManager.Contract.UsePriceSubmitterAsRandomProvider(&_FlareSystemManager.CallOpts)
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

// ChangeRandomProvider is a paid mutator transaction binding the contract method 0x460f6ce9.
//
// Solidity: function changeRandomProvider(bool _usePriceSubmitter) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) ChangeRandomProvider(opts *bind.TransactOpts, _usePriceSubmitter bool) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "changeRandomProvider", _usePriceSubmitter)
}

// ChangeRandomProvider is a paid mutator transaction binding the contract method 0x460f6ce9.
//
// Solidity: function changeRandomProvider(bool _usePriceSubmitter) returns()
func (_FlareSystemManager *FlareSystemManagerSession) ChangeRandomProvider(_usePriceSubmitter bool) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.ChangeRandomProvider(&_FlareSystemManager.TransactOpts, _usePriceSubmitter)
}

// ChangeRandomProvider is a paid mutator transaction binding the contract method 0x460f6ce9.
//
// Solidity: function changeRandomProvider(bool _usePriceSubmitter) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) ChangeRandomProvider(_usePriceSubmitter bool) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.ChangeRandomProvider(&_FlareSystemManager.TransactOpts, _usePriceSubmitter)
}

// ChangeSigningPolicySettings is a paid mutator transaction binding the contract method 0xb556fae3.
//
// Solidity: function changeSigningPolicySettings(uint64 _signingPolicyThresholdPPM, uint64 _signingPolicyMinNumberOfVoters) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) ChangeSigningPolicySettings(opts *bind.TransactOpts, _signingPolicyThresholdPPM uint64, _signingPolicyMinNumberOfVoters uint64) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "changeSigningPolicySettings", _signingPolicyThresholdPPM, _signingPolicyMinNumberOfVoters)
}

// ChangeSigningPolicySettings is a paid mutator transaction binding the contract method 0xb556fae3.
//
// Solidity: function changeSigningPolicySettings(uint64 _signingPolicyThresholdPPM, uint64 _signingPolicyMinNumberOfVoters) returns()
func (_FlareSystemManager *FlareSystemManagerSession) ChangeSigningPolicySettings(_signingPolicyThresholdPPM uint64, _signingPolicyMinNumberOfVoters uint64) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.ChangeSigningPolicySettings(&_FlareSystemManager.TransactOpts, _signingPolicyThresholdPPM, _signingPolicyMinNumberOfVoters)
}

// ChangeSigningPolicySettings is a paid mutator transaction binding the contract method 0xb556fae3.
//
// Solidity: function changeSigningPolicySettings(uint64 _signingPolicyThresholdPPM, uint64 _signingPolicyMinNumberOfVoters) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) ChangeSigningPolicySettings(_signingPolicyThresholdPPM uint64, _signingPolicyMinNumberOfVoters uint64) (*types.Transaction, error) {
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

// SignNewSigningPolicy is a paid mutator transaction binding the contract method 0x6b4c7bd6.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) SignNewSigningPolicy(opts *bind.TransactOpts, _rewardEpochId *big.Int, _newSigningPolicyHash [32]byte, _signature FlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "signNewSigningPolicy", _rewardEpochId, _newSigningPolicyHash, _signature)
}

// SignNewSigningPolicy is a paid mutator transaction binding the contract method 0x6b4c7bd6.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerSession) SignNewSigningPolicy(_rewardEpochId *big.Int, _newSigningPolicyHash [32]byte, _signature FlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SignNewSigningPolicy(&_FlareSystemManager.TransactOpts, _rewardEpochId, _newSigningPolicyHash, _signature)
}

// SignNewSigningPolicy is a paid mutator transaction binding the contract method 0x6b4c7bd6.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) SignNewSigningPolicy(_rewardEpochId *big.Int, _newSigningPolicyHash [32]byte, _signature FlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SignNewSigningPolicy(&_FlareSystemManager.TransactOpts, _rewardEpochId, _newSigningPolicyHash, _signature)
}

// SignRewards is a paid mutator transaction binding the contract method 0xab25ac5b.
//
// Solidity: function signRewards(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) SignRewards(opts *bind.TransactOpts, _rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte, _signature FlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "signRewards", _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash, _signature)
}

// SignRewards is a paid mutator transaction binding the contract method 0xab25ac5b.
//
// Solidity: function signRewards(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerSession) SignRewards(_rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte, _signature FlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SignRewards(&_FlareSystemManager.TransactOpts, _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash, _signature)
}

// SignRewards is a paid mutator transaction binding the contract method 0xab25ac5b.
//
// Solidity: function signRewards(uint24 _rewardEpochId, uint64 _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) SignRewards(_rewardEpochId *big.Int, _noOfWeightBasedClaims uint64, _rewardsHash [32]byte, _signature FlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SignRewards(&_FlareSystemManager.TransactOpts, _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash, _signature)
}

// SignUptimeVote is a paid mutator transaction binding the contract method 0xdc5a4225.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactor) SignUptimeVote(opts *bind.TransactOpts, _rewardEpochId *big.Int, _uptimeVoteHash [32]byte, _signature FlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.contract.Transact(opts, "signUptimeVote", _rewardEpochId, _uptimeVoteHash, _signature)
}

// SignUptimeVote is a paid mutator transaction binding the contract method 0xdc5a4225.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerSession) SignUptimeVote(_rewardEpochId *big.Int, _uptimeVoteHash [32]byte, _signature FlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SignUptimeVote(&_FlareSystemManager.TransactOpts, _rewardEpochId, _uptimeVoteHash, _signature)
}

// SignUptimeVote is a paid mutator transaction binding the contract method 0xdc5a4225.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (_FlareSystemManager *FlareSystemManagerTransactorSession) SignUptimeVote(_rewardEpochId *big.Int, _uptimeVoteHash [32]byte, _signature FlareSystemManagerSignature) (*types.Transaction, error) {
	return _FlareSystemManager.Contract.SignUptimeVote(&_FlareSystemManager.TransactOpts, _rewardEpochId, _uptimeVoteHash, _signature)
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
// Solidity: event RandomAcquisitionStarted(uint24 rewardEpochId, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterRandomAcquisitionStarted(opts *bind.FilterOpts) (*FlareSystemManagerRandomAcquisitionStartedIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "RandomAcquisitionStarted")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerRandomAcquisitionStartedIterator{contract: _FlareSystemManager.contract, event: "RandomAcquisitionStarted", logs: logs, sub: sub}, nil
}

// WatchRandomAcquisitionStarted is a free log subscription operation binding the contract event 0xf9991783e5e480e42d9a54d3f35f4321857f8f0ebeb3742d326dce28b1126708.
//
// Solidity: event RandomAcquisitionStarted(uint24 rewardEpochId, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchRandomAcquisitionStarted(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerRandomAcquisitionStarted) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "RandomAcquisitionStarted")
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
// Solidity: event RandomAcquisitionStarted(uint24 rewardEpochId, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseRandomAcquisitionStarted(log types.Log) (*FlareSystemManagerRandomAcquisitionStarted, error) {
	event := new(FlareSystemManagerRandomAcquisitionStarted)
	if err := _FlareSystemManager.contract.UnpackLog(event, "RandomAcquisitionStarted", log); err != nil {
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
	SigningAddress        common.Address
	Voter                 common.Address
	RewardsHash           [32]byte
	NoOfWeightBasedClaims *big.Int
	Timestamp             uint64
	ThresholdReached      bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRewardsSigned is a free log retrieval operation binding the contract event 0x80df246a47153e3604b0026fcf53ec85ea48178efd4f3cad24f9a1e1da2dd52a.
//
// Solidity: event RewardsSigned(uint24 rewardEpochId, address signingAddress, address voter, bytes32 rewardsHash, uint256 noOfWeightBasedClaims, uint64 timestamp, bool thresholdReached)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterRewardsSigned(opts *bind.FilterOpts) (*FlareSystemManagerRewardsSignedIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "RewardsSigned")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerRewardsSignedIterator{contract: _FlareSystemManager.contract, event: "RewardsSigned", logs: logs, sub: sub}, nil
}

// WatchRewardsSigned is a free log subscription operation binding the contract event 0x80df246a47153e3604b0026fcf53ec85ea48178efd4f3cad24f9a1e1da2dd52a.
//
// Solidity: event RewardsSigned(uint24 rewardEpochId, address signingAddress, address voter, bytes32 rewardsHash, uint256 noOfWeightBasedClaims, uint64 timestamp, bool thresholdReached)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchRewardsSigned(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerRewardsSigned) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "RewardsSigned")
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
// Solidity: event RewardsSigned(uint24 rewardEpochId, address signingAddress, address voter, bytes32 rewardsHash, uint256 noOfWeightBasedClaims, uint64 timestamp, bool thresholdReached)
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseRewardsSigned(log types.Log) (*FlareSystemManagerRewardsSigned, error) {
	event := new(FlareSystemManagerRewardsSigned)
	if err := _FlareSystemManager.contract.UnpackLog(event, "RewardsSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareSystemManagerSigningPolicyInitializedIterator is returned from FilterSigningPolicyInitialized and is used to iterate over the raw logs and unpacked data for SigningPolicyInitialized events raised by the FlareSystemManager contract.
type FlareSystemManagerSigningPolicyInitializedIterator struct {
	Event *FlareSystemManagerSigningPolicyInitialized // Event containing the contract specifics and raw log

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
func (it *FlareSystemManagerSigningPolicyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareSystemManagerSigningPolicyInitialized)
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
		it.Event = new(FlareSystemManagerSigningPolicyInitialized)
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
func (it *FlareSystemManagerSigningPolicyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareSystemManagerSigningPolicyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareSystemManagerSigningPolicyInitialized represents a SigningPolicyInitialized event raised by the FlareSystemManager contract.
type FlareSystemManagerSigningPolicyInitialized struct {
	RewardEpochId      *big.Int
	StartVotingRoundId uint32
	Threshold          uint16
	Seed               *big.Int
	Voters             []common.Address
	Weights            []uint16
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSigningPolicyInitialized is a free log retrieval operation binding the contract event 0x1e4381bbecb29293043991163c05f6055838c66670a47b415b81f107f5c49755.
//
// Solidity: event SigningPolicyInitialized(uint24 rewardEpochId, uint32 startVotingRoundId, uint16 threshold, uint256 seed, address[] voters, uint16[] weights)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterSigningPolicyInitialized(opts *bind.FilterOpts) (*FlareSystemManagerSigningPolicyInitializedIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "SigningPolicyInitialized")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerSigningPolicyInitializedIterator{contract: _FlareSystemManager.contract, event: "SigningPolicyInitialized", logs: logs, sub: sub}, nil
}

// WatchSigningPolicyInitialized is a free log subscription operation binding the contract event 0x1e4381bbecb29293043991163c05f6055838c66670a47b415b81f107f5c49755.
//
// Solidity: event SigningPolicyInitialized(uint24 rewardEpochId, uint32 startVotingRoundId, uint16 threshold, uint256 seed, address[] voters, uint16[] weights)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchSigningPolicyInitialized(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerSigningPolicyInitialized) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "SigningPolicyInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareSystemManagerSigningPolicyInitialized)
				if err := _FlareSystemManager.contract.UnpackLog(event, "SigningPolicyInitialized", log); err != nil {
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

// ParseSigningPolicyInitialized is a log parse operation binding the contract event 0x1e4381bbecb29293043991163c05f6055838c66670a47b415b81f107f5c49755.
//
// Solidity: event SigningPolicyInitialized(uint24 rewardEpochId, uint32 startVotingRoundId, uint16 threshold, uint256 seed, address[] voters, uint16[] weights)
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseSigningPolicyInitialized(log types.Log) (*FlareSystemManagerSigningPolicyInitialized, error) {
	event := new(FlareSystemManagerSigningPolicyInitialized)
	if err := _FlareSystemManager.contract.UnpackLog(event, "SigningPolicyInitialized", log); err != nil {
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
	RewardEpochId    *big.Int
	SigningAddress   common.Address
	Voter            common.Address
	Timestamp        uint64
	ThresholdReached bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSigningPolicySigned is a free log retrieval operation binding the contract event 0x154b0214ae62d8a5548c1eac25fabd87c38b04932a217732e1022f3118da67f3.
//
// Solidity: event SigningPolicySigned(uint24 rewardEpochId, address signingAddress, address voter, uint64 timestamp, bool thresholdReached)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterSigningPolicySigned(opts *bind.FilterOpts) (*FlareSystemManagerSigningPolicySignedIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "SigningPolicySigned")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerSigningPolicySignedIterator{contract: _FlareSystemManager.contract, event: "SigningPolicySigned", logs: logs, sub: sub}, nil
}

// WatchSigningPolicySigned is a free log subscription operation binding the contract event 0x154b0214ae62d8a5548c1eac25fabd87c38b04932a217732e1022f3118da67f3.
//
// Solidity: event SigningPolicySigned(uint24 rewardEpochId, address signingAddress, address voter, uint64 timestamp, bool thresholdReached)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchSigningPolicySigned(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerSigningPolicySigned) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "SigningPolicySigned")
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
// Solidity: event SigningPolicySigned(uint24 rewardEpochId, address signingAddress, address voter, uint64 timestamp, bool thresholdReached)
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseSigningPolicySigned(log types.Log) (*FlareSystemManagerSigningPolicySigned, error) {
	event := new(FlareSystemManagerSigningPolicySigned)
	if err := _FlareSystemManager.contract.UnpackLog(event, "SigningPolicySigned", log); err != nil {
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
	RewardEpochId    *big.Int
	SigningAddress   common.Address
	Voter            common.Address
	UptimeVoteHash   [32]byte
	Timestamp        uint64
	ThresholdReached bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUptimeVoteSigned is a free log retrieval operation binding the contract event 0x5506337d1266599f8b64675a1c8321701657ca2f2f70be0e0c58302b6c22e797.
//
// Solidity: event UptimeVoteSigned(uint24 rewardEpochId, address signingAddress, address voter, bytes32 uptimeVoteHash, uint64 timestamp, bool thresholdReached)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterUptimeVoteSigned(opts *bind.FilterOpts) (*FlareSystemManagerUptimeVoteSignedIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "UptimeVoteSigned")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerUptimeVoteSignedIterator{contract: _FlareSystemManager.contract, event: "UptimeVoteSigned", logs: logs, sub: sub}, nil
}

// WatchUptimeVoteSigned is a free log subscription operation binding the contract event 0x5506337d1266599f8b64675a1c8321701657ca2f2f70be0e0c58302b6c22e797.
//
// Solidity: event UptimeVoteSigned(uint24 rewardEpochId, address signingAddress, address voter, bytes32 uptimeVoteHash, uint64 timestamp, bool thresholdReached)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchUptimeVoteSigned(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerUptimeVoteSigned) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "UptimeVoteSigned")
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
// Solidity: event UptimeVoteSigned(uint24 rewardEpochId, address signingAddress, address voter, bytes32 uptimeVoteHash, uint64 timestamp, bool thresholdReached)
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseUptimeVoteSigned(log types.Log) (*FlareSystemManagerUptimeVoteSigned, error) {
	event := new(FlareSystemManagerUptimeVoteSigned)
	if err := _FlareSystemManager.contract.UnpackLog(event, "UptimeVoteSigned", log); err != nil {
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
// Solidity: event VotePowerBlockSelected(uint24 rewardEpochId, uint64 votePowerBlock, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) FilterVotePowerBlockSelected(opts *bind.FilterOpts) (*FlareSystemManagerVotePowerBlockSelectedIterator, error) {

	logs, sub, err := _FlareSystemManager.contract.FilterLogs(opts, "VotePowerBlockSelected")
	if err != nil {
		return nil, err
	}
	return &FlareSystemManagerVotePowerBlockSelectedIterator{contract: _FlareSystemManager.contract, event: "VotePowerBlockSelected", logs: logs, sub: sub}, nil
}

// WatchVotePowerBlockSelected is a free log subscription operation binding the contract event 0xf21722dbe044a7cea0f6d81c871cae750971e36c9dd10999e46f2b46f26ac7ff.
//
// Solidity: event VotePowerBlockSelected(uint24 rewardEpochId, uint64 votePowerBlock, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) WatchVotePowerBlockSelected(opts *bind.WatchOpts, sink chan<- *FlareSystemManagerVotePowerBlockSelected) (event.Subscription, error) {

	logs, sub, err := _FlareSystemManager.contract.WatchLogs(opts, "VotePowerBlockSelected")
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
// Solidity: event VotePowerBlockSelected(uint24 rewardEpochId, uint64 votePowerBlock, uint64 timestamp)
func (_FlareSystemManager *FlareSystemManagerFilterer) ParseVotePowerBlockSelected(log types.Log) (*FlareSystemManagerVotePowerBlockSelected, error) {
	event := new(FlareSystemManagerVotePowerBlockSelected)
	if err := _FlareSystemManager.contract.UnpackLog(event, "VotePowerBlockSelected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
