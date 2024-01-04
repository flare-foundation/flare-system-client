// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relay

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

// RelaySigningPolicy is an auto generated low-level Go binding around an user-defined struct.
type RelaySigningPolicy struct {
	RewardEpochId      *big.Int
	StartVotingRoundId uint32
	Threshold          uint16
	Seed               *big.Int
	Voters             []common.Address
	Weights            []uint16
}

// RelayMetaData contains all meta data concerning the Relay contract.
var RelayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signingPolicySetter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialRewardEpochId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_initialSigningPolicyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_randomNumberProtocolId\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_firstVotingRoundStartTs\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_votingEpochDurationSeconds\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_firstRewardEpochStartVotingRoundId\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_rewardEpochDurationInVotingEpochs\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_thresholdIncreaseBIPS\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"protocolId\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"votingRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"randomQualityScore\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"ProtocolMessageRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startVotingRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"threshold\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint16[]\",\"name\":\"weights\",\"type\":\"uint16[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signingPolicyBytes\",\"type\":\"bytes\"}],\"name\":\"SigningPolicyInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"SigningPolicyRelayed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_AND_WEIGHT_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ADDRESS_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ADDRESS_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MD_BOFF_numberOfVoters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MD_BOFF_rewardEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MD_BOFF_startingVotingRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MD_BOFF_threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MD_MASK_numberOfVoters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MD_MASK_rewardEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MD_MASK_startingVotingRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MD_MASK_threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSAGE_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSAGE_NO_MR_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"METADATA_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MSG_NMR_BOFF_protocolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MSG_NMR_BOFF_randomQualityScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MSG_NMR_BOFF_votingRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MSG_NMR_MASK_protocolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MSG_NMR_MASK_randomQualityScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MSG_NMR_MASK_votingRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"M_0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"M_1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"M_2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"M_2_signingPolicyHashTmp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"M_3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"M_3_existingSigningPolicyHashTmp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"M_4\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"M_5_stateData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"M_6_merkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NUMBER_OF_SIGNATURES_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NUMBER_OF_SIGNATURES_MASK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NUMBER_OF_SIGNATURES_RIGHT_SHIFT_BITS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NUMBER_OF_VOTERS_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NUMBER_OF_VOTERS_MASK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROTOCOL_ID_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RANDOM_SEED_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_BOFF_firstRewardEpochStartVotingRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_BOFF_firstVotingRoundStartTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_BOFF_randomNumberProtocolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_BOFF_randomNumberQualityScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_BOFF_randomVotingRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_BOFF_rewardEpochDurationInVotingEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_BOFF_thresholdIncreaseBIPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_BOFF_votingEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_MASK_firstRewardEpochStartVotingRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_MASK_firstVotingRoundStartTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_MASK_randomNumberProtocolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_MASK_randomNumberQualityScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_MASK_randomVotingRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_MASK_rewardEpochDurationInVotingEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_MASK_thresholdIncreaseBIPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_MASK_votingEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SELECTOR_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INDEX_RIGHT_SHIFT_BITS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_V_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_WITH_INDEX_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNING_POLICY_PREFIX_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"THRESHOLD_BIPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WEIGHT_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WEIGHT_MASK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_protocolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_votingRoundId\",\"type\":\"uint256\"}],\"name\":\"getConfirmedMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRandomNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_randomNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_randomNumberQualityScore\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_randomTimestamp\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getVotingRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInitializedRewardEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_result\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"startVotingRoundId\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"threshold\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"internalType\":\"uint16[]\",\"name\":\"weights\",\"type\":\"uint16[]\"}],\"internalType\":\"structRelay.SigningPolicy\",\"name\":\"_signingPolicy\",\"type\":\"tuple\"}],\"name\":\"setSigningPolicy\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingPolicySetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateData\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"randomNumberProtocolId\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"firstVotingRoundStartTs\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"votingEpochDurationSeconds\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"firstRewardEpochStartVotingRoundId\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"rewardEpochDurationInVotingEpochs\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"thresholdIncreaseBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"randomVotingRoundId\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"randomNumberQualityScore\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"toSigningPolicyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RelayABI is the input ABI used to generate the binding from.
// Deprecated: Use RelayMetaData.ABI instead.
var RelayABI = RelayMetaData.ABI

// Relay is an auto generated Go binding around an Ethereum contract.
type Relay struct {
	RelayCaller     // Read-only binding to the contract
	RelayTransactor // Write-only binding to the contract
	RelayFilterer   // Log filterer for contract events
}

// RelayCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelaySession struct {
	Contract     *Relay            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayCallerSession struct {
	Contract *RelayCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RelayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayTransactorSession struct {
	Contract     *RelayTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayRaw struct {
	Contract *Relay // Generic contract binding to access the raw methods on
}

// RelayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayCallerRaw struct {
	Contract *RelayCaller // Generic read-only contract binding to access the raw methods on
}

// RelayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayTransactorRaw struct {
	Contract *RelayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelay creates a new instance of Relay, bound to a specific deployed contract.
func NewRelay(address common.Address, backend bind.ContractBackend) (*Relay, error) {
	contract, err := bindRelay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relay{RelayCaller: RelayCaller{contract: contract}, RelayTransactor: RelayTransactor{contract: contract}, RelayFilterer: RelayFilterer{contract: contract}}, nil
}

// NewRelayCaller creates a new read-only instance of Relay, bound to a specific deployed contract.
func NewRelayCaller(address common.Address, caller bind.ContractCaller) (*RelayCaller, error) {
	contract, err := bindRelay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayCaller{contract: contract}, nil
}

// NewRelayTransactor creates a new write-only instance of Relay, bound to a specific deployed contract.
func NewRelayTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayTransactor, error) {
	contract, err := bindRelay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayTransactor{contract: contract}, nil
}

// NewRelayFilterer creates a new log filterer instance of Relay, bound to a specific deployed contract.
func NewRelayFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayFilterer, error) {
	contract, err := bindRelay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayFilterer{contract: contract}, nil
}

// bindRelay binds a generic wrapper to an already deployed contract.
func bindRelay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RelayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relay *RelayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Relay.Contract.RelayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relay *RelayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relay.Contract.RelayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relay *RelayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relay.Contract.RelayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relay *RelayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Relay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relay *RelayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relay *RelayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relay.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSANDWEIGHTBYTES is a free data retrieval call binding the contract method 0x81e41f0d.
//
// Solidity: function ADDRESS_AND_WEIGHT_BYTES() view returns(uint256)
func (_Relay *RelayCaller) ADDRESSANDWEIGHTBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "ADDRESS_AND_WEIGHT_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ADDRESSANDWEIGHTBYTES is a free data retrieval call binding the contract method 0x81e41f0d.
//
// Solidity: function ADDRESS_AND_WEIGHT_BYTES() view returns(uint256)
func (_Relay *RelaySession) ADDRESSANDWEIGHTBYTES() (*big.Int, error) {
	return _Relay.Contract.ADDRESSANDWEIGHTBYTES(&_Relay.CallOpts)
}

// ADDRESSANDWEIGHTBYTES is a free data retrieval call binding the contract method 0x81e41f0d.
//
// Solidity: function ADDRESS_AND_WEIGHT_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) ADDRESSANDWEIGHTBYTES() (*big.Int, error) {
	return _Relay.Contract.ADDRESSANDWEIGHTBYTES(&_Relay.CallOpts)
}

// ADDRESSBYTES is a free data retrieval call binding the contract method 0x2395c2de.
//
// Solidity: function ADDRESS_BYTES() view returns(uint256)
func (_Relay *RelayCaller) ADDRESSBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "ADDRESS_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ADDRESSBYTES is a free data retrieval call binding the contract method 0x2395c2de.
//
// Solidity: function ADDRESS_BYTES() view returns(uint256)
func (_Relay *RelaySession) ADDRESSBYTES() (*big.Int, error) {
	return _Relay.Contract.ADDRESSBYTES(&_Relay.CallOpts)
}

// ADDRESSBYTES is a free data retrieval call binding the contract method 0x2395c2de.
//
// Solidity: function ADDRESS_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) ADDRESSBYTES() (*big.Int, error) {
	return _Relay.Contract.ADDRESSBYTES(&_Relay.CallOpts)
}

// ADDRESSOFFSET is a free data retrieval call binding the contract method 0xe2f742d8.
//
// Solidity: function ADDRESS_OFFSET() view returns(uint256)
func (_Relay *RelayCaller) ADDRESSOFFSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "ADDRESS_OFFSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ADDRESSOFFSET is a free data retrieval call binding the contract method 0xe2f742d8.
//
// Solidity: function ADDRESS_OFFSET() view returns(uint256)
func (_Relay *RelaySession) ADDRESSOFFSET() (*big.Int, error) {
	return _Relay.Contract.ADDRESSOFFSET(&_Relay.CallOpts)
}

// ADDRESSOFFSET is a free data retrieval call binding the contract method 0xe2f742d8.
//
// Solidity: function ADDRESS_OFFSET() view returns(uint256)
func (_Relay *RelayCallerSession) ADDRESSOFFSET() (*big.Int, error) {
	return _Relay.Contract.ADDRESSOFFSET(&_Relay.CallOpts)
}

// MDBOFFNumberOfVoters is a free data retrieval call binding the contract method 0x21acc681.
//
// Solidity: function MD_BOFF_numberOfVoters() view returns(uint256)
func (_Relay *RelayCaller) MDBOFFNumberOfVoters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MD_BOFF_numberOfVoters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MDBOFFNumberOfVoters is a free data retrieval call binding the contract method 0x21acc681.
//
// Solidity: function MD_BOFF_numberOfVoters() view returns(uint256)
func (_Relay *RelaySession) MDBOFFNumberOfVoters() (*big.Int, error) {
	return _Relay.Contract.MDBOFFNumberOfVoters(&_Relay.CallOpts)
}

// MDBOFFNumberOfVoters is a free data retrieval call binding the contract method 0x21acc681.
//
// Solidity: function MD_BOFF_numberOfVoters() view returns(uint256)
func (_Relay *RelayCallerSession) MDBOFFNumberOfVoters() (*big.Int, error) {
	return _Relay.Contract.MDBOFFNumberOfVoters(&_Relay.CallOpts)
}

// MDBOFFRewardEpochId is a free data retrieval call binding the contract method 0x2a0ffe2e.
//
// Solidity: function MD_BOFF_rewardEpochId() view returns(uint256)
func (_Relay *RelayCaller) MDBOFFRewardEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MD_BOFF_rewardEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MDBOFFRewardEpochId is a free data retrieval call binding the contract method 0x2a0ffe2e.
//
// Solidity: function MD_BOFF_rewardEpochId() view returns(uint256)
func (_Relay *RelaySession) MDBOFFRewardEpochId() (*big.Int, error) {
	return _Relay.Contract.MDBOFFRewardEpochId(&_Relay.CallOpts)
}

// MDBOFFRewardEpochId is a free data retrieval call binding the contract method 0x2a0ffe2e.
//
// Solidity: function MD_BOFF_rewardEpochId() view returns(uint256)
func (_Relay *RelayCallerSession) MDBOFFRewardEpochId() (*big.Int, error) {
	return _Relay.Contract.MDBOFFRewardEpochId(&_Relay.CallOpts)
}

// MDBOFFStartingVotingRoundId is a free data retrieval call binding the contract method 0xca55ae81.
//
// Solidity: function MD_BOFF_startingVotingRoundId() view returns(uint256)
func (_Relay *RelayCaller) MDBOFFStartingVotingRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MD_BOFF_startingVotingRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MDBOFFStartingVotingRoundId is a free data retrieval call binding the contract method 0xca55ae81.
//
// Solidity: function MD_BOFF_startingVotingRoundId() view returns(uint256)
func (_Relay *RelaySession) MDBOFFStartingVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.MDBOFFStartingVotingRoundId(&_Relay.CallOpts)
}

// MDBOFFStartingVotingRoundId is a free data retrieval call binding the contract method 0xca55ae81.
//
// Solidity: function MD_BOFF_startingVotingRoundId() view returns(uint256)
func (_Relay *RelayCallerSession) MDBOFFStartingVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.MDBOFFStartingVotingRoundId(&_Relay.CallOpts)
}

// MDBOFFThreshold is a free data retrieval call binding the contract method 0x8c05389b.
//
// Solidity: function MD_BOFF_threshold() view returns(uint256)
func (_Relay *RelayCaller) MDBOFFThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MD_BOFF_threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MDBOFFThreshold is a free data retrieval call binding the contract method 0x8c05389b.
//
// Solidity: function MD_BOFF_threshold() view returns(uint256)
func (_Relay *RelaySession) MDBOFFThreshold() (*big.Int, error) {
	return _Relay.Contract.MDBOFFThreshold(&_Relay.CallOpts)
}

// MDBOFFThreshold is a free data retrieval call binding the contract method 0x8c05389b.
//
// Solidity: function MD_BOFF_threshold() view returns(uint256)
func (_Relay *RelayCallerSession) MDBOFFThreshold() (*big.Int, error) {
	return _Relay.Contract.MDBOFFThreshold(&_Relay.CallOpts)
}

// MDMASKNumberOfVoters is a free data retrieval call binding the contract method 0x96bfc0e2.
//
// Solidity: function MD_MASK_numberOfVoters() view returns(uint256)
func (_Relay *RelayCaller) MDMASKNumberOfVoters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MD_MASK_numberOfVoters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MDMASKNumberOfVoters is a free data retrieval call binding the contract method 0x96bfc0e2.
//
// Solidity: function MD_MASK_numberOfVoters() view returns(uint256)
func (_Relay *RelaySession) MDMASKNumberOfVoters() (*big.Int, error) {
	return _Relay.Contract.MDMASKNumberOfVoters(&_Relay.CallOpts)
}

// MDMASKNumberOfVoters is a free data retrieval call binding the contract method 0x96bfc0e2.
//
// Solidity: function MD_MASK_numberOfVoters() view returns(uint256)
func (_Relay *RelayCallerSession) MDMASKNumberOfVoters() (*big.Int, error) {
	return _Relay.Contract.MDMASKNumberOfVoters(&_Relay.CallOpts)
}

// MDMASKRewardEpochId is a free data retrieval call binding the contract method 0x14ee9464.
//
// Solidity: function MD_MASK_rewardEpochId() view returns(uint256)
func (_Relay *RelayCaller) MDMASKRewardEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MD_MASK_rewardEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MDMASKRewardEpochId is a free data retrieval call binding the contract method 0x14ee9464.
//
// Solidity: function MD_MASK_rewardEpochId() view returns(uint256)
func (_Relay *RelaySession) MDMASKRewardEpochId() (*big.Int, error) {
	return _Relay.Contract.MDMASKRewardEpochId(&_Relay.CallOpts)
}

// MDMASKRewardEpochId is a free data retrieval call binding the contract method 0x14ee9464.
//
// Solidity: function MD_MASK_rewardEpochId() view returns(uint256)
func (_Relay *RelayCallerSession) MDMASKRewardEpochId() (*big.Int, error) {
	return _Relay.Contract.MDMASKRewardEpochId(&_Relay.CallOpts)
}

// MDMASKStartingVotingRoundId is a free data retrieval call binding the contract method 0xe7b1e0e2.
//
// Solidity: function MD_MASK_startingVotingRoundId() view returns(uint256)
func (_Relay *RelayCaller) MDMASKStartingVotingRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MD_MASK_startingVotingRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MDMASKStartingVotingRoundId is a free data retrieval call binding the contract method 0xe7b1e0e2.
//
// Solidity: function MD_MASK_startingVotingRoundId() view returns(uint256)
func (_Relay *RelaySession) MDMASKStartingVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.MDMASKStartingVotingRoundId(&_Relay.CallOpts)
}

// MDMASKStartingVotingRoundId is a free data retrieval call binding the contract method 0xe7b1e0e2.
//
// Solidity: function MD_MASK_startingVotingRoundId() view returns(uint256)
func (_Relay *RelayCallerSession) MDMASKStartingVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.MDMASKStartingVotingRoundId(&_Relay.CallOpts)
}

// MDMASKThreshold is a free data retrieval call binding the contract method 0xfc4e1654.
//
// Solidity: function MD_MASK_threshold() view returns(uint256)
func (_Relay *RelayCaller) MDMASKThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MD_MASK_threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MDMASKThreshold is a free data retrieval call binding the contract method 0xfc4e1654.
//
// Solidity: function MD_MASK_threshold() view returns(uint256)
func (_Relay *RelaySession) MDMASKThreshold() (*big.Int, error) {
	return _Relay.Contract.MDMASKThreshold(&_Relay.CallOpts)
}

// MDMASKThreshold is a free data retrieval call binding the contract method 0xfc4e1654.
//
// Solidity: function MD_MASK_threshold() view returns(uint256)
func (_Relay *RelayCallerSession) MDMASKThreshold() (*big.Int, error) {
	return _Relay.Contract.MDMASKThreshold(&_Relay.CallOpts)
}

// MESSAGEBYTES is a free data retrieval call binding the contract method 0xddceb34f.
//
// Solidity: function MESSAGE_BYTES() view returns(uint256)
func (_Relay *RelayCaller) MESSAGEBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MESSAGE_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MESSAGEBYTES is a free data retrieval call binding the contract method 0xddceb34f.
//
// Solidity: function MESSAGE_BYTES() view returns(uint256)
func (_Relay *RelaySession) MESSAGEBYTES() (*big.Int, error) {
	return _Relay.Contract.MESSAGEBYTES(&_Relay.CallOpts)
}

// MESSAGEBYTES is a free data retrieval call binding the contract method 0xddceb34f.
//
// Solidity: function MESSAGE_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) MESSAGEBYTES() (*big.Int, error) {
	return _Relay.Contract.MESSAGEBYTES(&_Relay.CallOpts)
}

// MESSAGENOMRBYTES is a free data retrieval call binding the contract method 0xfcce02f5.
//
// Solidity: function MESSAGE_NO_MR_BYTES() view returns(uint256)
func (_Relay *RelayCaller) MESSAGENOMRBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MESSAGE_NO_MR_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MESSAGENOMRBYTES is a free data retrieval call binding the contract method 0xfcce02f5.
//
// Solidity: function MESSAGE_NO_MR_BYTES() view returns(uint256)
func (_Relay *RelaySession) MESSAGENOMRBYTES() (*big.Int, error) {
	return _Relay.Contract.MESSAGENOMRBYTES(&_Relay.CallOpts)
}

// MESSAGENOMRBYTES is a free data retrieval call binding the contract method 0xfcce02f5.
//
// Solidity: function MESSAGE_NO_MR_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) MESSAGENOMRBYTES() (*big.Int, error) {
	return _Relay.Contract.MESSAGENOMRBYTES(&_Relay.CallOpts)
}

// METADATABYTES is a free data retrieval call binding the contract method 0x078a3450.
//
// Solidity: function METADATA_BYTES() view returns(uint256)
func (_Relay *RelayCaller) METADATABYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "METADATA_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// METADATABYTES is a free data retrieval call binding the contract method 0x078a3450.
//
// Solidity: function METADATA_BYTES() view returns(uint256)
func (_Relay *RelaySession) METADATABYTES() (*big.Int, error) {
	return _Relay.Contract.METADATABYTES(&_Relay.CallOpts)
}

// METADATABYTES is a free data retrieval call binding the contract method 0x078a3450.
//
// Solidity: function METADATA_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) METADATABYTES() (*big.Int, error) {
	return _Relay.Contract.METADATABYTES(&_Relay.CallOpts)
}

// MSGNMRBOFFProtocolId is a free data retrieval call binding the contract method 0x25ab6ba6.
//
// Solidity: function MSG_NMR_BOFF_protocolId() view returns(uint256)
func (_Relay *RelayCaller) MSGNMRBOFFProtocolId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MSG_NMR_BOFF_protocolId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MSGNMRBOFFProtocolId is a free data retrieval call binding the contract method 0x25ab6ba6.
//
// Solidity: function MSG_NMR_BOFF_protocolId() view returns(uint256)
func (_Relay *RelaySession) MSGNMRBOFFProtocolId() (*big.Int, error) {
	return _Relay.Contract.MSGNMRBOFFProtocolId(&_Relay.CallOpts)
}

// MSGNMRBOFFProtocolId is a free data retrieval call binding the contract method 0x25ab6ba6.
//
// Solidity: function MSG_NMR_BOFF_protocolId() view returns(uint256)
func (_Relay *RelayCallerSession) MSGNMRBOFFProtocolId() (*big.Int, error) {
	return _Relay.Contract.MSGNMRBOFFProtocolId(&_Relay.CallOpts)
}

// MSGNMRBOFFRandomQualityScore is a free data retrieval call binding the contract method 0xf0623116.
//
// Solidity: function MSG_NMR_BOFF_randomQualityScore() view returns(uint256)
func (_Relay *RelayCaller) MSGNMRBOFFRandomQualityScore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MSG_NMR_BOFF_randomQualityScore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MSGNMRBOFFRandomQualityScore is a free data retrieval call binding the contract method 0xf0623116.
//
// Solidity: function MSG_NMR_BOFF_randomQualityScore() view returns(uint256)
func (_Relay *RelaySession) MSGNMRBOFFRandomQualityScore() (*big.Int, error) {
	return _Relay.Contract.MSGNMRBOFFRandomQualityScore(&_Relay.CallOpts)
}

// MSGNMRBOFFRandomQualityScore is a free data retrieval call binding the contract method 0xf0623116.
//
// Solidity: function MSG_NMR_BOFF_randomQualityScore() view returns(uint256)
func (_Relay *RelayCallerSession) MSGNMRBOFFRandomQualityScore() (*big.Int, error) {
	return _Relay.Contract.MSGNMRBOFFRandomQualityScore(&_Relay.CallOpts)
}

// MSGNMRBOFFVotingRoundId is a free data retrieval call binding the contract method 0x0478b476.
//
// Solidity: function MSG_NMR_BOFF_votingRoundId() view returns(uint256)
func (_Relay *RelayCaller) MSGNMRBOFFVotingRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MSG_NMR_BOFF_votingRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MSGNMRBOFFVotingRoundId is a free data retrieval call binding the contract method 0x0478b476.
//
// Solidity: function MSG_NMR_BOFF_votingRoundId() view returns(uint256)
func (_Relay *RelaySession) MSGNMRBOFFVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.MSGNMRBOFFVotingRoundId(&_Relay.CallOpts)
}

// MSGNMRBOFFVotingRoundId is a free data retrieval call binding the contract method 0x0478b476.
//
// Solidity: function MSG_NMR_BOFF_votingRoundId() view returns(uint256)
func (_Relay *RelayCallerSession) MSGNMRBOFFVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.MSGNMRBOFFVotingRoundId(&_Relay.CallOpts)
}

// MSGNMRMASKProtocolId is a free data retrieval call binding the contract method 0xa3c14334.
//
// Solidity: function MSG_NMR_MASK_protocolId() view returns(uint256)
func (_Relay *RelayCaller) MSGNMRMASKProtocolId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MSG_NMR_MASK_protocolId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MSGNMRMASKProtocolId is a free data retrieval call binding the contract method 0xa3c14334.
//
// Solidity: function MSG_NMR_MASK_protocolId() view returns(uint256)
func (_Relay *RelaySession) MSGNMRMASKProtocolId() (*big.Int, error) {
	return _Relay.Contract.MSGNMRMASKProtocolId(&_Relay.CallOpts)
}

// MSGNMRMASKProtocolId is a free data retrieval call binding the contract method 0xa3c14334.
//
// Solidity: function MSG_NMR_MASK_protocolId() view returns(uint256)
func (_Relay *RelayCallerSession) MSGNMRMASKProtocolId() (*big.Int, error) {
	return _Relay.Contract.MSGNMRMASKProtocolId(&_Relay.CallOpts)
}

// MSGNMRMASKRandomQualityScore is a free data retrieval call binding the contract method 0x80c13889.
//
// Solidity: function MSG_NMR_MASK_randomQualityScore() view returns(uint256)
func (_Relay *RelayCaller) MSGNMRMASKRandomQualityScore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MSG_NMR_MASK_randomQualityScore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MSGNMRMASKRandomQualityScore is a free data retrieval call binding the contract method 0x80c13889.
//
// Solidity: function MSG_NMR_MASK_randomQualityScore() view returns(uint256)
func (_Relay *RelaySession) MSGNMRMASKRandomQualityScore() (*big.Int, error) {
	return _Relay.Contract.MSGNMRMASKRandomQualityScore(&_Relay.CallOpts)
}

// MSGNMRMASKRandomQualityScore is a free data retrieval call binding the contract method 0x80c13889.
//
// Solidity: function MSG_NMR_MASK_randomQualityScore() view returns(uint256)
func (_Relay *RelayCallerSession) MSGNMRMASKRandomQualityScore() (*big.Int, error) {
	return _Relay.Contract.MSGNMRMASKRandomQualityScore(&_Relay.CallOpts)
}

// MSGNMRMASKVotingRoundId is a free data retrieval call binding the contract method 0x7a1a530b.
//
// Solidity: function MSG_NMR_MASK_votingRoundId() view returns(uint256)
func (_Relay *RelayCaller) MSGNMRMASKVotingRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "MSG_NMR_MASK_votingRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MSGNMRMASKVotingRoundId is a free data retrieval call binding the contract method 0x7a1a530b.
//
// Solidity: function MSG_NMR_MASK_votingRoundId() view returns(uint256)
func (_Relay *RelaySession) MSGNMRMASKVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.MSGNMRMASKVotingRoundId(&_Relay.CallOpts)
}

// MSGNMRMASKVotingRoundId is a free data retrieval call binding the contract method 0x7a1a530b.
//
// Solidity: function MSG_NMR_MASK_votingRoundId() view returns(uint256)
func (_Relay *RelayCallerSession) MSGNMRMASKVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.MSGNMRMASKVotingRoundId(&_Relay.CallOpts)
}

// M0 is a free data retrieval call binding the contract method 0x67e4b8c9.
//
// Solidity: function M_0() view returns(uint256)
func (_Relay *RelayCaller) M0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "M_0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// M0 is a free data retrieval call binding the contract method 0x67e4b8c9.
//
// Solidity: function M_0() view returns(uint256)
func (_Relay *RelaySession) M0() (*big.Int, error) {
	return _Relay.Contract.M0(&_Relay.CallOpts)
}

// M0 is a free data retrieval call binding the contract method 0x67e4b8c9.
//
// Solidity: function M_0() view returns(uint256)
func (_Relay *RelayCallerSession) M0() (*big.Int, error) {
	return _Relay.Contract.M0(&_Relay.CallOpts)
}

// M1 is a free data retrieval call binding the contract method 0x0f996d90.
//
// Solidity: function M_1() view returns(uint256)
func (_Relay *RelayCaller) M1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "M_1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// M1 is a free data retrieval call binding the contract method 0x0f996d90.
//
// Solidity: function M_1() view returns(uint256)
func (_Relay *RelaySession) M1() (*big.Int, error) {
	return _Relay.Contract.M1(&_Relay.CallOpts)
}

// M1 is a free data retrieval call binding the contract method 0x0f996d90.
//
// Solidity: function M_1() view returns(uint256)
func (_Relay *RelayCallerSession) M1() (*big.Int, error) {
	return _Relay.Contract.M1(&_Relay.CallOpts)
}

// M2 is a free data retrieval call binding the contract method 0x57ac3f3e.
//
// Solidity: function M_2() view returns(uint256)
func (_Relay *RelayCaller) M2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "M_2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// M2 is a free data retrieval call binding the contract method 0x57ac3f3e.
//
// Solidity: function M_2() view returns(uint256)
func (_Relay *RelaySession) M2() (*big.Int, error) {
	return _Relay.Contract.M2(&_Relay.CallOpts)
}

// M2 is a free data retrieval call binding the contract method 0x57ac3f3e.
//
// Solidity: function M_2() view returns(uint256)
func (_Relay *RelayCallerSession) M2() (*big.Int, error) {
	return _Relay.Contract.M2(&_Relay.CallOpts)
}

// M2SigningPolicyHashTmp is a free data retrieval call binding the contract method 0x63392842.
//
// Solidity: function M_2_signingPolicyHashTmp() view returns(uint256)
func (_Relay *RelayCaller) M2SigningPolicyHashTmp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "M_2_signingPolicyHashTmp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// M2SigningPolicyHashTmp is a free data retrieval call binding the contract method 0x63392842.
//
// Solidity: function M_2_signingPolicyHashTmp() view returns(uint256)
func (_Relay *RelaySession) M2SigningPolicyHashTmp() (*big.Int, error) {
	return _Relay.Contract.M2SigningPolicyHashTmp(&_Relay.CallOpts)
}

// M2SigningPolicyHashTmp is a free data retrieval call binding the contract method 0x63392842.
//
// Solidity: function M_2_signingPolicyHashTmp() view returns(uint256)
func (_Relay *RelayCallerSession) M2SigningPolicyHashTmp() (*big.Int, error) {
	return _Relay.Contract.M2SigningPolicyHashTmp(&_Relay.CallOpts)
}

// M3 is a free data retrieval call binding the contract method 0x56c069c3.
//
// Solidity: function M_3() view returns(uint256)
func (_Relay *RelayCaller) M3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "M_3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// M3 is a free data retrieval call binding the contract method 0x56c069c3.
//
// Solidity: function M_3() view returns(uint256)
func (_Relay *RelaySession) M3() (*big.Int, error) {
	return _Relay.Contract.M3(&_Relay.CallOpts)
}

// M3 is a free data retrieval call binding the contract method 0x56c069c3.
//
// Solidity: function M_3() view returns(uint256)
func (_Relay *RelayCallerSession) M3() (*big.Int, error) {
	return _Relay.Contract.M3(&_Relay.CallOpts)
}

// M3ExistingSigningPolicyHashTmp is a free data retrieval call binding the contract method 0xde819567.
//
// Solidity: function M_3_existingSigningPolicyHashTmp() view returns(uint256)
func (_Relay *RelayCaller) M3ExistingSigningPolicyHashTmp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "M_3_existingSigningPolicyHashTmp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// M3ExistingSigningPolicyHashTmp is a free data retrieval call binding the contract method 0xde819567.
//
// Solidity: function M_3_existingSigningPolicyHashTmp() view returns(uint256)
func (_Relay *RelaySession) M3ExistingSigningPolicyHashTmp() (*big.Int, error) {
	return _Relay.Contract.M3ExistingSigningPolicyHashTmp(&_Relay.CallOpts)
}

// M3ExistingSigningPolicyHashTmp is a free data retrieval call binding the contract method 0xde819567.
//
// Solidity: function M_3_existingSigningPolicyHashTmp() view returns(uint256)
func (_Relay *RelayCallerSession) M3ExistingSigningPolicyHashTmp() (*big.Int, error) {
	return _Relay.Contract.M3ExistingSigningPolicyHashTmp(&_Relay.CallOpts)
}

// M4 is a free data retrieval call binding the contract method 0xbf5615a3.
//
// Solidity: function M_4() view returns(uint256)
func (_Relay *RelayCaller) M4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "M_4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// M4 is a free data retrieval call binding the contract method 0xbf5615a3.
//
// Solidity: function M_4() view returns(uint256)
func (_Relay *RelaySession) M4() (*big.Int, error) {
	return _Relay.Contract.M4(&_Relay.CallOpts)
}

// M4 is a free data retrieval call binding the contract method 0xbf5615a3.
//
// Solidity: function M_4() view returns(uint256)
func (_Relay *RelayCallerSession) M4() (*big.Int, error) {
	return _Relay.Contract.M4(&_Relay.CallOpts)
}

// M5StateData is a free data retrieval call binding the contract method 0x65bde92b.
//
// Solidity: function M_5_stateData() view returns(uint256)
func (_Relay *RelayCaller) M5StateData(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "M_5_stateData")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// M5StateData is a free data retrieval call binding the contract method 0x65bde92b.
//
// Solidity: function M_5_stateData() view returns(uint256)
func (_Relay *RelaySession) M5StateData() (*big.Int, error) {
	return _Relay.Contract.M5StateData(&_Relay.CallOpts)
}

// M5StateData is a free data retrieval call binding the contract method 0x65bde92b.
//
// Solidity: function M_5_stateData() view returns(uint256)
func (_Relay *RelayCallerSession) M5StateData() (*big.Int, error) {
	return _Relay.Contract.M5StateData(&_Relay.CallOpts)
}

// M6MerkleRoot is a free data retrieval call binding the contract method 0x4c233a88.
//
// Solidity: function M_6_merkleRoot() view returns(uint256)
func (_Relay *RelayCaller) M6MerkleRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "M_6_merkleRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// M6MerkleRoot is a free data retrieval call binding the contract method 0x4c233a88.
//
// Solidity: function M_6_merkleRoot() view returns(uint256)
func (_Relay *RelaySession) M6MerkleRoot() (*big.Int, error) {
	return _Relay.Contract.M6MerkleRoot(&_Relay.CallOpts)
}

// M6MerkleRoot is a free data retrieval call binding the contract method 0x4c233a88.
//
// Solidity: function M_6_merkleRoot() view returns(uint256)
func (_Relay *RelayCallerSession) M6MerkleRoot() (*big.Int, error) {
	return _Relay.Contract.M6MerkleRoot(&_Relay.CallOpts)
}

// NUMBEROFSIGNATURESBYTES is a free data retrieval call binding the contract method 0x7482919e.
//
// Solidity: function NUMBER_OF_SIGNATURES_BYTES() view returns(uint256)
func (_Relay *RelayCaller) NUMBEROFSIGNATURESBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "NUMBER_OF_SIGNATURES_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NUMBEROFSIGNATURESBYTES is a free data retrieval call binding the contract method 0x7482919e.
//
// Solidity: function NUMBER_OF_SIGNATURES_BYTES() view returns(uint256)
func (_Relay *RelaySession) NUMBEROFSIGNATURESBYTES() (*big.Int, error) {
	return _Relay.Contract.NUMBEROFSIGNATURESBYTES(&_Relay.CallOpts)
}

// NUMBEROFSIGNATURESBYTES is a free data retrieval call binding the contract method 0x7482919e.
//
// Solidity: function NUMBER_OF_SIGNATURES_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) NUMBEROFSIGNATURESBYTES() (*big.Int, error) {
	return _Relay.Contract.NUMBEROFSIGNATURESBYTES(&_Relay.CallOpts)
}

// NUMBEROFSIGNATURESMASK is a free data retrieval call binding the contract method 0x4538345e.
//
// Solidity: function NUMBER_OF_SIGNATURES_MASK() view returns(uint256)
func (_Relay *RelayCaller) NUMBEROFSIGNATURESMASK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "NUMBER_OF_SIGNATURES_MASK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NUMBEROFSIGNATURESMASK is a free data retrieval call binding the contract method 0x4538345e.
//
// Solidity: function NUMBER_OF_SIGNATURES_MASK() view returns(uint256)
func (_Relay *RelaySession) NUMBEROFSIGNATURESMASK() (*big.Int, error) {
	return _Relay.Contract.NUMBEROFSIGNATURESMASK(&_Relay.CallOpts)
}

// NUMBEROFSIGNATURESMASK is a free data retrieval call binding the contract method 0x4538345e.
//
// Solidity: function NUMBER_OF_SIGNATURES_MASK() view returns(uint256)
func (_Relay *RelayCallerSession) NUMBEROFSIGNATURESMASK() (*big.Int, error) {
	return _Relay.Contract.NUMBEROFSIGNATURESMASK(&_Relay.CallOpts)
}

// NUMBEROFSIGNATURESRIGHTSHIFTBITS is a free data retrieval call binding the contract method 0x4e5a0927.
//
// Solidity: function NUMBER_OF_SIGNATURES_RIGHT_SHIFT_BITS() view returns(uint256)
func (_Relay *RelayCaller) NUMBEROFSIGNATURESRIGHTSHIFTBITS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "NUMBER_OF_SIGNATURES_RIGHT_SHIFT_BITS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NUMBEROFSIGNATURESRIGHTSHIFTBITS is a free data retrieval call binding the contract method 0x4e5a0927.
//
// Solidity: function NUMBER_OF_SIGNATURES_RIGHT_SHIFT_BITS() view returns(uint256)
func (_Relay *RelaySession) NUMBEROFSIGNATURESRIGHTSHIFTBITS() (*big.Int, error) {
	return _Relay.Contract.NUMBEROFSIGNATURESRIGHTSHIFTBITS(&_Relay.CallOpts)
}

// NUMBEROFSIGNATURESRIGHTSHIFTBITS is a free data retrieval call binding the contract method 0x4e5a0927.
//
// Solidity: function NUMBER_OF_SIGNATURES_RIGHT_SHIFT_BITS() view returns(uint256)
func (_Relay *RelayCallerSession) NUMBEROFSIGNATURESRIGHTSHIFTBITS() (*big.Int, error) {
	return _Relay.Contract.NUMBEROFSIGNATURESRIGHTSHIFTBITS(&_Relay.CallOpts)
}

// NUMBEROFVOTERSBYTES is a free data retrieval call binding the contract method 0xb87127ec.
//
// Solidity: function NUMBER_OF_VOTERS_BYTES() view returns(uint256)
func (_Relay *RelayCaller) NUMBEROFVOTERSBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "NUMBER_OF_VOTERS_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NUMBEROFVOTERSBYTES is a free data retrieval call binding the contract method 0xb87127ec.
//
// Solidity: function NUMBER_OF_VOTERS_BYTES() view returns(uint256)
func (_Relay *RelaySession) NUMBEROFVOTERSBYTES() (*big.Int, error) {
	return _Relay.Contract.NUMBEROFVOTERSBYTES(&_Relay.CallOpts)
}

// NUMBEROFVOTERSBYTES is a free data retrieval call binding the contract method 0xb87127ec.
//
// Solidity: function NUMBER_OF_VOTERS_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) NUMBEROFVOTERSBYTES() (*big.Int, error) {
	return _Relay.Contract.NUMBEROFVOTERSBYTES(&_Relay.CallOpts)
}

// NUMBEROFVOTERSMASK is a free data retrieval call binding the contract method 0x8602a984.
//
// Solidity: function NUMBER_OF_VOTERS_MASK() view returns(uint256)
func (_Relay *RelayCaller) NUMBEROFVOTERSMASK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "NUMBER_OF_VOTERS_MASK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NUMBEROFVOTERSMASK is a free data retrieval call binding the contract method 0x8602a984.
//
// Solidity: function NUMBER_OF_VOTERS_MASK() view returns(uint256)
func (_Relay *RelaySession) NUMBEROFVOTERSMASK() (*big.Int, error) {
	return _Relay.Contract.NUMBEROFVOTERSMASK(&_Relay.CallOpts)
}

// NUMBEROFVOTERSMASK is a free data retrieval call binding the contract method 0x8602a984.
//
// Solidity: function NUMBER_OF_VOTERS_MASK() view returns(uint256)
func (_Relay *RelayCallerSession) NUMBEROFVOTERSMASK() (*big.Int, error) {
	return _Relay.Contract.NUMBEROFVOTERSMASK(&_Relay.CallOpts)
}

// PROTOCOLIDBYTES is a free data retrieval call binding the contract method 0xb31b4617.
//
// Solidity: function PROTOCOL_ID_BYTES() view returns(uint256)
func (_Relay *RelayCaller) PROTOCOLIDBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "PROTOCOL_ID_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROTOCOLIDBYTES is a free data retrieval call binding the contract method 0xb31b4617.
//
// Solidity: function PROTOCOL_ID_BYTES() view returns(uint256)
func (_Relay *RelaySession) PROTOCOLIDBYTES() (*big.Int, error) {
	return _Relay.Contract.PROTOCOLIDBYTES(&_Relay.CallOpts)
}

// PROTOCOLIDBYTES is a free data retrieval call binding the contract method 0xb31b4617.
//
// Solidity: function PROTOCOL_ID_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) PROTOCOLIDBYTES() (*big.Int, error) {
	return _Relay.Contract.PROTOCOLIDBYTES(&_Relay.CallOpts)
}

// RANDOMSEEDBYTES is a free data retrieval call binding the contract method 0x101810f2.
//
// Solidity: function RANDOM_SEED_BYTES() view returns(uint256)
func (_Relay *RelayCaller) RANDOMSEEDBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "RANDOM_SEED_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RANDOMSEEDBYTES is a free data retrieval call binding the contract method 0x101810f2.
//
// Solidity: function RANDOM_SEED_BYTES() view returns(uint256)
func (_Relay *RelaySession) RANDOMSEEDBYTES() (*big.Int, error) {
	return _Relay.Contract.RANDOMSEEDBYTES(&_Relay.CallOpts)
}

// RANDOMSEEDBYTES is a free data retrieval call binding the contract method 0x101810f2.
//
// Solidity: function RANDOM_SEED_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) RANDOMSEEDBYTES() (*big.Int, error) {
	return _Relay.Contract.RANDOMSEEDBYTES(&_Relay.CallOpts)
}

// SDBOFFFirstRewardEpochStartVotingRoundId is a free data retrieval call binding the contract method 0x6e599939.
//
// Solidity: function SD_BOFF_firstRewardEpochStartVotingRoundId() view returns(uint256)
func (_Relay *RelayCaller) SDBOFFFirstRewardEpochStartVotingRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_BOFF_firstRewardEpochStartVotingRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDBOFFFirstRewardEpochStartVotingRoundId is a free data retrieval call binding the contract method 0x6e599939.
//
// Solidity: function SD_BOFF_firstRewardEpochStartVotingRoundId() view returns(uint256)
func (_Relay *RelaySession) SDBOFFFirstRewardEpochStartVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.SDBOFFFirstRewardEpochStartVotingRoundId(&_Relay.CallOpts)
}

// SDBOFFFirstRewardEpochStartVotingRoundId is a free data retrieval call binding the contract method 0x6e599939.
//
// Solidity: function SD_BOFF_firstRewardEpochStartVotingRoundId() view returns(uint256)
func (_Relay *RelayCallerSession) SDBOFFFirstRewardEpochStartVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.SDBOFFFirstRewardEpochStartVotingRoundId(&_Relay.CallOpts)
}

// SDBOFFFirstVotingRoundStartTs is a free data retrieval call binding the contract method 0x86d93c59.
//
// Solidity: function SD_BOFF_firstVotingRoundStartTs() view returns(uint256)
func (_Relay *RelayCaller) SDBOFFFirstVotingRoundStartTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_BOFF_firstVotingRoundStartTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDBOFFFirstVotingRoundStartTs is a free data retrieval call binding the contract method 0x86d93c59.
//
// Solidity: function SD_BOFF_firstVotingRoundStartTs() view returns(uint256)
func (_Relay *RelaySession) SDBOFFFirstVotingRoundStartTs() (*big.Int, error) {
	return _Relay.Contract.SDBOFFFirstVotingRoundStartTs(&_Relay.CallOpts)
}

// SDBOFFFirstVotingRoundStartTs is a free data retrieval call binding the contract method 0x86d93c59.
//
// Solidity: function SD_BOFF_firstVotingRoundStartTs() view returns(uint256)
func (_Relay *RelayCallerSession) SDBOFFFirstVotingRoundStartTs() (*big.Int, error) {
	return _Relay.Contract.SDBOFFFirstVotingRoundStartTs(&_Relay.CallOpts)
}

// SDBOFFRandomNumberProtocolId is a free data retrieval call binding the contract method 0x73787fdc.
//
// Solidity: function SD_BOFF_randomNumberProtocolId() view returns(uint256)
func (_Relay *RelayCaller) SDBOFFRandomNumberProtocolId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_BOFF_randomNumberProtocolId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDBOFFRandomNumberProtocolId is a free data retrieval call binding the contract method 0x73787fdc.
//
// Solidity: function SD_BOFF_randomNumberProtocolId() view returns(uint256)
func (_Relay *RelaySession) SDBOFFRandomNumberProtocolId() (*big.Int, error) {
	return _Relay.Contract.SDBOFFRandomNumberProtocolId(&_Relay.CallOpts)
}

// SDBOFFRandomNumberProtocolId is a free data retrieval call binding the contract method 0x73787fdc.
//
// Solidity: function SD_BOFF_randomNumberProtocolId() view returns(uint256)
func (_Relay *RelayCallerSession) SDBOFFRandomNumberProtocolId() (*big.Int, error) {
	return _Relay.Contract.SDBOFFRandomNumberProtocolId(&_Relay.CallOpts)
}

// SDBOFFRandomNumberQualityScore is a free data retrieval call binding the contract method 0x71d3f1a7.
//
// Solidity: function SD_BOFF_randomNumberQualityScore() view returns(uint256)
func (_Relay *RelayCaller) SDBOFFRandomNumberQualityScore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_BOFF_randomNumberQualityScore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDBOFFRandomNumberQualityScore is a free data retrieval call binding the contract method 0x71d3f1a7.
//
// Solidity: function SD_BOFF_randomNumberQualityScore() view returns(uint256)
func (_Relay *RelaySession) SDBOFFRandomNumberQualityScore() (*big.Int, error) {
	return _Relay.Contract.SDBOFFRandomNumberQualityScore(&_Relay.CallOpts)
}

// SDBOFFRandomNumberQualityScore is a free data retrieval call binding the contract method 0x71d3f1a7.
//
// Solidity: function SD_BOFF_randomNumberQualityScore() view returns(uint256)
func (_Relay *RelayCallerSession) SDBOFFRandomNumberQualityScore() (*big.Int, error) {
	return _Relay.Contract.SDBOFFRandomNumberQualityScore(&_Relay.CallOpts)
}

// SDBOFFRandomVotingRoundId is a free data retrieval call binding the contract method 0xe73bb1eb.
//
// Solidity: function SD_BOFF_randomVotingRoundId() view returns(uint256)
func (_Relay *RelayCaller) SDBOFFRandomVotingRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_BOFF_randomVotingRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDBOFFRandomVotingRoundId is a free data retrieval call binding the contract method 0xe73bb1eb.
//
// Solidity: function SD_BOFF_randomVotingRoundId() view returns(uint256)
func (_Relay *RelaySession) SDBOFFRandomVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.SDBOFFRandomVotingRoundId(&_Relay.CallOpts)
}

// SDBOFFRandomVotingRoundId is a free data retrieval call binding the contract method 0xe73bb1eb.
//
// Solidity: function SD_BOFF_randomVotingRoundId() view returns(uint256)
func (_Relay *RelayCallerSession) SDBOFFRandomVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.SDBOFFRandomVotingRoundId(&_Relay.CallOpts)
}

// SDBOFFRewardEpochDurationInVotingEpochs is a free data retrieval call binding the contract method 0x4975240d.
//
// Solidity: function SD_BOFF_rewardEpochDurationInVotingEpochs() view returns(uint256)
func (_Relay *RelayCaller) SDBOFFRewardEpochDurationInVotingEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_BOFF_rewardEpochDurationInVotingEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDBOFFRewardEpochDurationInVotingEpochs is a free data retrieval call binding the contract method 0x4975240d.
//
// Solidity: function SD_BOFF_rewardEpochDurationInVotingEpochs() view returns(uint256)
func (_Relay *RelaySession) SDBOFFRewardEpochDurationInVotingEpochs() (*big.Int, error) {
	return _Relay.Contract.SDBOFFRewardEpochDurationInVotingEpochs(&_Relay.CallOpts)
}

// SDBOFFRewardEpochDurationInVotingEpochs is a free data retrieval call binding the contract method 0x4975240d.
//
// Solidity: function SD_BOFF_rewardEpochDurationInVotingEpochs() view returns(uint256)
func (_Relay *RelayCallerSession) SDBOFFRewardEpochDurationInVotingEpochs() (*big.Int, error) {
	return _Relay.Contract.SDBOFFRewardEpochDurationInVotingEpochs(&_Relay.CallOpts)
}

// SDBOFFThresholdIncreaseBIPS is a free data retrieval call binding the contract method 0xdfa543f9.
//
// Solidity: function SD_BOFF_thresholdIncreaseBIPS() view returns(uint256)
func (_Relay *RelayCaller) SDBOFFThresholdIncreaseBIPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_BOFF_thresholdIncreaseBIPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDBOFFThresholdIncreaseBIPS is a free data retrieval call binding the contract method 0xdfa543f9.
//
// Solidity: function SD_BOFF_thresholdIncreaseBIPS() view returns(uint256)
func (_Relay *RelaySession) SDBOFFThresholdIncreaseBIPS() (*big.Int, error) {
	return _Relay.Contract.SDBOFFThresholdIncreaseBIPS(&_Relay.CallOpts)
}

// SDBOFFThresholdIncreaseBIPS is a free data retrieval call binding the contract method 0xdfa543f9.
//
// Solidity: function SD_BOFF_thresholdIncreaseBIPS() view returns(uint256)
func (_Relay *RelayCallerSession) SDBOFFThresholdIncreaseBIPS() (*big.Int, error) {
	return _Relay.Contract.SDBOFFThresholdIncreaseBIPS(&_Relay.CallOpts)
}

// SDBOFFVotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x4225816a.
//
// Solidity: function SD_BOFF_votingEpochDurationSeconds() view returns(uint256)
func (_Relay *RelayCaller) SDBOFFVotingEpochDurationSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_BOFF_votingEpochDurationSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDBOFFVotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x4225816a.
//
// Solidity: function SD_BOFF_votingEpochDurationSeconds() view returns(uint256)
func (_Relay *RelaySession) SDBOFFVotingEpochDurationSeconds() (*big.Int, error) {
	return _Relay.Contract.SDBOFFVotingEpochDurationSeconds(&_Relay.CallOpts)
}

// SDBOFFVotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x4225816a.
//
// Solidity: function SD_BOFF_votingEpochDurationSeconds() view returns(uint256)
func (_Relay *RelayCallerSession) SDBOFFVotingEpochDurationSeconds() (*big.Int, error) {
	return _Relay.Contract.SDBOFFVotingEpochDurationSeconds(&_Relay.CallOpts)
}

// SDMASKFirstRewardEpochStartVotingRoundId is a free data retrieval call binding the contract method 0x70b71b41.
//
// Solidity: function SD_MASK_firstRewardEpochStartVotingRoundId() view returns(uint256)
func (_Relay *RelayCaller) SDMASKFirstRewardEpochStartVotingRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_MASK_firstRewardEpochStartVotingRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDMASKFirstRewardEpochStartVotingRoundId is a free data retrieval call binding the contract method 0x70b71b41.
//
// Solidity: function SD_MASK_firstRewardEpochStartVotingRoundId() view returns(uint256)
func (_Relay *RelaySession) SDMASKFirstRewardEpochStartVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.SDMASKFirstRewardEpochStartVotingRoundId(&_Relay.CallOpts)
}

// SDMASKFirstRewardEpochStartVotingRoundId is a free data retrieval call binding the contract method 0x70b71b41.
//
// Solidity: function SD_MASK_firstRewardEpochStartVotingRoundId() view returns(uint256)
func (_Relay *RelayCallerSession) SDMASKFirstRewardEpochStartVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.SDMASKFirstRewardEpochStartVotingRoundId(&_Relay.CallOpts)
}

// SDMASKFirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xf74cd56a.
//
// Solidity: function SD_MASK_firstVotingRoundStartTs() view returns(uint256)
func (_Relay *RelayCaller) SDMASKFirstVotingRoundStartTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_MASK_firstVotingRoundStartTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDMASKFirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xf74cd56a.
//
// Solidity: function SD_MASK_firstVotingRoundStartTs() view returns(uint256)
func (_Relay *RelaySession) SDMASKFirstVotingRoundStartTs() (*big.Int, error) {
	return _Relay.Contract.SDMASKFirstVotingRoundStartTs(&_Relay.CallOpts)
}

// SDMASKFirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xf74cd56a.
//
// Solidity: function SD_MASK_firstVotingRoundStartTs() view returns(uint256)
func (_Relay *RelayCallerSession) SDMASKFirstVotingRoundStartTs() (*big.Int, error) {
	return _Relay.Contract.SDMASKFirstVotingRoundStartTs(&_Relay.CallOpts)
}

// SDMASKRandomNumberProtocolId is a free data retrieval call binding the contract method 0xd2f714eb.
//
// Solidity: function SD_MASK_randomNumberProtocolId() view returns(uint256)
func (_Relay *RelayCaller) SDMASKRandomNumberProtocolId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_MASK_randomNumberProtocolId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDMASKRandomNumberProtocolId is a free data retrieval call binding the contract method 0xd2f714eb.
//
// Solidity: function SD_MASK_randomNumberProtocolId() view returns(uint256)
func (_Relay *RelaySession) SDMASKRandomNumberProtocolId() (*big.Int, error) {
	return _Relay.Contract.SDMASKRandomNumberProtocolId(&_Relay.CallOpts)
}

// SDMASKRandomNumberProtocolId is a free data retrieval call binding the contract method 0xd2f714eb.
//
// Solidity: function SD_MASK_randomNumberProtocolId() view returns(uint256)
func (_Relay *RelayCallerSession) SDMASKRandomNumberProtocolId() (*big.Int, error) {
	return _Relay.Contract.SDMASKRandomNumberProtocolId(&_Relay.CallOpts)
}

// SDMASKRandomNumberQualityScore is a free data retrieval call binding the contract method 0x9a5e42d5.
//
// Solidity: function SD_MASK_randomNumberQualityScore() view returns(uint256)
func (_Relay *RelayCaller) SDMASKRandomNumberQualityScore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_MASK_randomNumberQualityScore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDMASKRandomNumberQualityScore is a free data retrieval call binding the contract method 0x9a5e42d5.
//
// Solidity: function SD_MASK_randomNumberQualityScore() view returns(uint256)
func (_Relay *RelaySession) SDMASKRandomNumberQualityScore() (*big.Int, error) {
	return _Relay.Contract.SDMASKRandomNumberQualityScore(&_Relay.CallOpts)
}

// SDMASKRandomNumberQualityScore is a free data retrieval call binding the contract method 0x9a5e42d5.
//
// Solidity: function SD_MASK_randomNumberQualityScore() view returns(uint256)
func (_Relay *RelayCallerSession) SDMASKRandomNumberQualityScore() (*big.Int, error) {
	return _Relay.Contract.SDMASKRandomNumberQualityScore(&_Relay.CallOpts)
}

// SDMASKRandomVotingRoundId is a free data retrieval call binding the contract method 0x28d41041.
//
// Solidity: function SD_MASK_randomVotingRoundId() view returns(uint256)
func (_Relay *RelayCaller) SDMASKRandomVotingRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_MASK_randomVotingRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDMASKRandomVotingRoundId is a free data retrieval call binding the contract method 0x28d41041.
//
// Solidity: function SD_MASK_randomVotingRoundId() view returns(uint256)
func (_Relay *RelaySession) SDMASKRandomVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.SDMASKRandomVotingRoundId(&_Relay.CallOpts)
}

// SDMASKRandomVotingRoundId is a free data retrieval call binding the contract method 0x28d41041.
//
// Solidity: function SD_MASK_randomVotingRoundId() view returns(uint256)
func (_Relay *RelayCallerSession) SDMASKRandomVotingRoundId() (*big.Int, error) {
	return _Relay.Contract.SDMASKRandomVotingRoundId(&_Relay.CallOpts)
}

// SDMASKRewardEpochDurationInVotingEpochs is a free data retrieval call binding the contract method 0xeae6a84a.
//
// Solidity: function SD_MASK_rewardEpochDurationInVotingEpochs() view returns(uint256)
func (_Relay *RelayCaller) SDMASKRewardEpochDurationInVotingEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_MASK_rewardEpochDurationInVotingEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDMASKRewardEpochDurationInVotingEpochs is a free data retrieval call binding the contract method 0xeae6a84a.
//
// Solidity: function SD_MASK_rewardEpochDurationInVotingEpochs() view returns(uint256)
func (_Relay *RelaySession) SDMASKRewardEpochDurationInVotingEpochs() (*big.Int, error) {
	return _Relay.Contract.SDMASKRewardEpochDurationInVotingEpochs(&_Relay.CallOpts)
}

// SDMASKRewardEpochDurationInVotingEpochs is a free data retrieval call binding the contract method 0xeae6a84a.
//
// Solidity: function SD_MASK_rewardEpochDurationInVotingEpochs() view returns(uint256)
func (_Relay *RelayCallerSession) SDMASKRewardEpochDurationInVotingEpochs() (*big.Int, error) {
	return _Relay.Contract.SDMASKRewardEpochDurationInVotingEpochs(&_Relay.CallOpts)
}

// SDMASKThresholdIncreaseBIPS is a free data retrieval call binding the contract method 0x798f6c57.
//
// Solidity: function SD_MASK_thresholdIncreaseBIPS() view returns(uint256)
func (_Relay *RelayCaller) SDMASKThresholdIncreaseBIPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_MASK_thresholdIncreaseBIPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDMASKThresholdIncreaseBIPS is a free data retrieval call binding the contract method 0x798f6c57.
//
// Solidity: function SD_MASK_thresholdIncreaseBIPS() view returns(uint256)
func (_Relay *RelaySession) SDMASKThresholdIncreaseBIPS() (*big.Int, error) {
	return _Relay.Contract.SDMASKThresholdIncreaseBIPS(&_Relay.CallOpts)
}

// SDMASKThresholdIncreaseBIPS is a free data retrieval call binding the contract method 0x798f6c57.
//
// Solidity: function SD_MASK_thresholdIncreaseBIPS() view returns(uint256)
func (_Relay *RelayCallerSession) SDMASKThresholdIncreaseBIPS() (*big.Int, error) {
	return _Relay.Contract.SDMASKThresholdIncreaseBIPS(&_Relay.CallOpts)
}

// SDMASKVotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x09ff2af2.
//
// Solidity: function SD_MASK_votingEpochDurationSeconds() view returns(uint256)
func (_Relay *RelayCaller) SDMASKVotingEpochDurationSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SD_MASK_votingEpochDurationSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDMASKVotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x09ff2af2.
//
// Solidity: function SD_MASK_votingEpochDurationSeconds() view returns(uint256)
func (_Relay *RelaySession) SDMASKVotingEpochDurationSeconds() (*big.Int, error) {
	return _Relay.Contract.SDMASKVotingEpochDurationSeconds(&_Relay.CallOpts)
}

// SDMASKVotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x09ff2af2.
//
// Solidity: function SD_MASK_votingEpochDurationSeconds() view returns(uint256)
func (_Relay *RelayCallerSession) SDMASKVotingEpochDurationSeconds() (*big.Int, error) {
	return _Relay.Contract.SDMASKVotingEpochDurationSeconds(&_Relay.CallOpts)
}

// SELECTORBYTES is a free data retrieval call binding the contract method 0xcc91c635.
//
// Solidity: function SELECTOR_BYTES() view returns(uint256)
func (_Relay *RelayCaller) SELECTORBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SELECTOR_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SELECTORBYTES is a free data retrieval call binding the contract method 0xcc91c635.
//
// Solidity: function SELECTOR_BYTES() view returns(uint256)
func (_Relay *RelaySession) SELECTORBYTES() (*big.Int, error) {
	return _Relay.Contract.SELECTORBYTES(&_Relay.CallOpts)
}

// SELECTORBYTES is a free data retrieval call binding the contract method 0xcc91c635.
//
// Solidity: function SELECTOR_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) SELECTORBYTES() (*big.Int, error) {
	return _Relay.Contract.SELECTORBYTES(&_Relay.CallOpts)
}

// SIGNATUREINDEXRIGHTSHIFTBITS is a free data retrieval call binding the contract method 0xd4588c24.
//
// Solidity: function SIGNATURE_INDEX_RIGHT_SHIFT_BITS() view returns(uint256)
func (_Relay *RelayCaller) SIGNATUREINDEXRIGHTSHIFTBITS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SIGNATURE_INDEX_RIGHT_SHIFT_BITS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGNATUREINDEXRIGHTSHIFTBITS is a free data retrieval call binding the contract method 0xd4588c24.
//
// Solidity: function SIGNATURE_INDEX_RIGHT_SHIFT_BITS() view returns(uint256)
func (_Relay *RelaySession) SIGNATUREINDEXRIGHTSHIFTBITS() (*big.Int, error) {
	return _Relay.Contract.SIGNATUREINDEXRIGHTSHIFTBITS(&_Relay.CallOpts)
}

// SIGNATUREINDEXRIGHTSHIFTBITS is a free data retrieval call binding the contract method 0xd4588c24.
//
// Solidity: function SIGNATURE_INDEX_RIGHT_SHIFT_BITS() view returns(uint256)
func (_Relay *RelayCallerSession) SIGNATUREINDEXRIGHTSHIFTBITS() (*big.Int, error) {
	return _Relay.Contract.SIGNATUREINDEXRIGHTSHIFTBITS(&_Relay.CallOpts)
}

// SIGNATUREVBYTES is a free data retrieval call binding the contract method 0xd2128d00.
//
// Solidity: function SIGNATURE_V_BYTES() view returns(uint256)
func (_Relay *RelayCaller) SIGNATUREVBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SIGNATURE_V_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGNATUREVBYTES is a free data retrieval call binding the contract method 0xd2128d00.
//
// Solidity: function SIGNATURE_V_BYTES() view returns(uint256)
func (_Relay *RelaySession) SIGNATUREVBYTES() (*big.Int, error) {
	return _Relay.Contract.SIGNATUREVBYTES(&_Relay.CallOpts)
}

// SIGNATUREVBYTES is a free data retrieval call binding the contract method 0xd2128d00.
//
// Solidity: function SIGNATURE_V_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) SIGNATUREVBYTES() (*big.Int, error) {
	return _Relay.Contract.SIGNATUREVBYTES(&_Relay.CallOpts)
}

// SIGNATUREWITHINDEXBYTES is a free data retrieval call binding the contract method 0xd636bfe9.
//
// Solidity: function SIGNATURE_WITH_INDEX_BYTES() view returns(uint256)
func (_Relay *RelayCaller) SIGNATUREWITHINDEXBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SIGNATURE_WITH_INDEX_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGNATUREWITHINDEXBYTES is a free data retrieval call binding the contract method 0xd636bfe9.
//
// Solidity: function SIGNATURE_WITH_INDEX_BYTES() view returns(uint256)
func (_Relay *RelaySession) SIGNATUREWITHINDEXBYTES() (*big.Int, error) {
	return _Relay.Contract.SIGNATUREWITHINDEXBYTES(&_Relay.CallOpts)
}

// SIGNATUREWITHINDEXBYTES is a free data retrieval call binding the contract method 0xd636bfe9.
//
// Solidity: function SIGNATURE_WITH_INDEX_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) SIGNATUREWITHINDEXBYTES() (*big.Int, error) {
	return _Relay.Contract.SIGNATUREWITHINDEXBYTES(&_Relay.CallOpts)
}

// SIGNINGPOLICYPREFIXBYTES is a free data retrieval call binding the contract method 0x14fa9b28.
//
// Solidity: function SIGNING_POLICY_PREFIX_BYTES() view returns(uint256)
func (_Relay *RelayCaller) SIGNINGPOLICYPREFIXBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "SIGNING_POLICY_PREFIX_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGNINGPOLICYPREFIXBYTES is a free data retrieval call binding the contract method 0x14fa9b28.
//
// Solidity: function SIGNING_POLICY_PREFIX_BYTES() view returns(uint256)
func (_Relay *RelaySession) SIGNINGPOLICYPREFIXBYTES() (*big.Int, error) {
	return _Relay.Contract.SIGNINGPOLICYPREFIXBYTES(&_Relay.CallOpts)
}

// SIGNINGPOLICYPREFIXBYTES is a free data retrieval call binding the contract method 0x14fa9b28.
//
// Solidity: function SIGNING_POLICY_PREFIX_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) SIGNINGPOLICYPREFIXBYTES() (*big.Int, error) {
	return _Relay.Contract.SIGNINGPOLICYPREFIXBYTES(&_Relay.CallOpts)
}

// THRESHOLDBIPS is a free data retrieval call binding the contract method 0x6249ffeb.
//
// Solidity: function THRESHOLD_BIPS() view returns(uint256)
func (_Relay *RelayCaller) THRESHOLDBIPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "THRESHOLD_BIPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// THRESHOLDBIPS is a free data retrieval call binding the contract method 0x6249ffeb.
//
// Solidity: function THRESHOLD_BIPS() view returns(uint256)
func (_Relay *RelaySession) THRESHOLDBIPS() (*big.Int, error) {
	return _Relay.Contract.THRESHOLDBIPS(&_Relay.CallOpts)
}

// THRESHOLDBIPS is a free data retrieval call binding the contract method 0x6249ffeb.
//
// Solidity: function THRESHOLD_BIPS() view returns(uint256)
func (_Relay *RelayCallerSession) THRESHOLDBIPS() (*big.Int, error) {
	return _Relay.Contract.THRESHOLDBIPS(&_Relay.CallOpts)
}

// WEIGHTBYTES is a free data retrieval call binding the contract method 0x125041a4.
//
// Solidity: function WEIGHT_BYTES() view returns(uint256)
func (_Relay *RelayCaller) WEIGHTBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "WEIGHT_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEIGHTBYTES is a free data retrieval call binding the contract method 0x125041a4.
//
// Solidity: function WEIGHT_BYTES() view returns(uint256)
func (_Relay *RelaySession) WEIGHTBYTES() (*big.Int, error) {
	return _Relay.Contract.WEIGHTBYTES(&_Relay.CallOpts)
}

// WEIGHTBYTES is a free data retrieval call binding the contract method 0x125041a4.
//
// Solidity: function WEIGHT_BYTES() view returns(uint256)
func (_Relay *RelayCallerSession) WEIGHTBYTES() (*big.Int, error) {
	return _Relay.Contract.WEIGHTBYTES(&_Relay.CallOpts)
}

// WEIGHTMASK is a free data retrieval call binding the contract method 0x236cb8ab.
//
// Solidity: function WEIGHT_MASK() view returns(uint256)
func (_Relay *RelayCaller) WEIGHTMASK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "WEIGHT_MASK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEIGHTMASK is a free data retrieval call binding the contract method 0x236cb8ab.
//
// Solidity: function WEIGHT_MASK() view returns(uint256)
func (_Relay *RelaySession) WEIGHTMASK() (*big.Int, error) {
	return _Relay.Contract.WEIGHTMASK(&_Relay.CallOpts)
}

// WEIGHTMASK is a free data retrieval call binding the contract method 0x236cb8ab.
//
// Solidity: function WEIGHT_MASK() view returns(uint256)
func (_Relay *RelayCallerSession) WEIGHTMASK() (*big.Int, error) {
	return _Relay.Contract.WEIGHTMASK(&_Relay.CallOpts)
}

// GetConfirmedMerkleRoot is a free data retrieval call binding the contract method 0x22c3f6fa.
//
// Solidity: function getConfirmedMerkleRoot(uint256 _protocolId, uint256 _votingRoundId) view returns(bytes32)
func (_Relay *RelayCaller) GetConfirmedMerkleRoot(opts *bind.CallOpts, _protocolId *big.Int, _votingRoundId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "getConfirmedMerkleRoot", _protocolId, _votingRoundId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConfirmedMerkleRoot is a free data retrieval call binding the contract method 0x22c3f6fa.
//
// Solidity: function getConfirmedMerkleRoot(uint256 _protocolId, uint256 _votingRoundId) view returns(bytes32)
func (_Relay *RelaySession) GetConfirmedMerkleRoot(_protocolId *big.Int, _votingRoundId *big.Int) ([32]byte, error) {
	return _Relay.Contract.GetConfirmedMerkleRoot(&_Relay.CallOpts, _protocolId, _votingRoundId)
}

// GetConfirmedMerkleRoot is a free data retrieval call binding the contract method 0x22c3f6fa.
//
// Solidity: function getConfirmedMerkleRoot(uint256 _protocolId, uint256 _votingRoundId) view returns(bytes32)
func (_Relay *RelayCallerSession) GetConfirmedMerkleRoot(_protocolId *big.Int, _votingRoundId *big.Int) ([32]byte, error) {
	return _Relay.Contract.GetConfirmedMerkleRoot(&_Relay.CallOpts, _protocolId, _votingRoundId)
}

// GetRandomNumber is a free data retrieval call binding the contract method 0xdbdff2c1.
//
// Solidity: function getRandomNumber() view returns(uint256 _randomNumber, bool _randomNumberQualityScore, uint32 _randomTimestamp)
func (_Relay *RelayCaller) GetRandomNumber(opts *bind.CallOpts) (struct {
	RandomNumber             *big.Int
	RandomNumberQualityScore bool
	RandomTimestamp          uint32
}, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "getRandomNumber")

	outstruct := new(struct {
		RandomNumber             *big.Int
		RandomNumberQualityScore bool
		RandomTimestamp          uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RandomNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RandomNumberQualityScore = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RandomTimestamp = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetRandomNumber is a free data retrieval call binding the contract method 0xdbdff2c1.
//
// Solidity: function getRandomNumber() view returns(uint256 _randomNumber, bool _randomNumberQualityScore, uint32 _randomTimestamp)
func (_Relay *RelaySession) GetRandomNumber() (struct {
	RandomNumber             *big.Int
	RandomNumberQualityScore bool
	RandomTimestamp          uint32
}, error) {
	return _Relay.Contract.GetRandomNumber(&_Relay.CallOpts)
}

// GetRandomNumber is a free data retrieval call binding the contract method 0xdbdff2c1.
//
// Solidity: function getRandomNumber() view returns(uint256 _randomNumber, bool _randomNumberQualityScore, uint32 _randomTimestamp)
func (_Relay *RelayCallerSession) GetRandomNumber() (struct {
	RandomNumber             *big.Int
	RandomNumberQualityScore bool
	RandomTimestamp          uint32
}, error) {
	return _Relay.Contract.GetRandomNumber(&_Relay.CallOpts)
}

// GetVotingRoundId is a free data retrieval call binding the contract method 0xab97db37.
//
// Solidity: function getVotingRoundId(uint256 _timestamp) view returns(uint256)
func (_Relay *RelayCaller) GetVotingRoundId(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "getVotingRoundId", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotingRoundId is a free data retrieval call binding the contract method 0xab97db37.
//
// Solidity: function getVotingRoundId(uint256 _timestamp) view returns(uint256)
func (_Relay *RelaySession) GetVotingRoundId(_timestamp *big.Int) (*big.Int, error) {
	return _Relay.Contract.GetVotingRoundId(&_Relay.CallOpts, _timestamp)
}

// GetVotingRoundId is a free data retrieval call binding the contract method 0xab97db37.
//
// Solidity: function getVotingRoundId(uint256 _timestamp) view returns(uint256)
func (_Relay *RelayCallerSession) GetVotingRoundId(_timestamp *big.Int) (*big.Int, error) {
	return _Relay.Contract.GetVotingRoundId(&_Relay.CallOpts, _timestamp)
}

// LastInitializedRewardEpoch is a free data retrieval call binding the contract method 0x1592a087.
//
// Solidity: function lastInitializedRewardEpoch() view returns(uint256)
func (_Relay *RelayCaller) LastInitializedRewardEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "lastInitializedRewardEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastInitializedRewardEpoch is a free data retrieval call binding the contract method 0x1592a087.
//
// Solidity: function lastInitializedRewardEpoch() view returns(uint256)
func (_Relay *RelaySession) LastInitializedRewardEpoch() (*big.Int, error) {
	return _Relay.Contract.LastInitializedRewardEpoch(&_Relay.CallOpts)
}

// LastInitializedRewardEpoch is a free data retrieval call binding the contract method 0x1592a087.
//
// Solidity: function lastInitializedRewardEpoch() view returns(uint256)
func (_Relay *RelayCallerSession) LastInitializedRewardEpoch() (*big.Int, error) {
	return _Relay.Contract.LastInitializedRewardEpoch(&_Relay.CallOpts)
}

// MerkleRoots is a free data retrieval call binding the contract method 0x39436b00.
//
// Solidity: function merkleRoots(uint256 , uint256 ) view returns(bytes32)
func (_Relay *RelayCaller) MerkleRoots(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "merkleRoots", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoots is a free data retrieval call binding the contract method 0x39436b00.
//
// Solidity: function merkleRoots(uint256 , uint256 ) view returns(bytes32)
func (_Relay *RelaySession) MerkleRoots(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _Relay.Contract.MerkleRoots(&_Relay.CallOpts, arg0, arg1)
}

// MerkleRoots is a free data retrieval call binding the contract method 0x39436b00.
//
// Solidity: function merkleRoots(uint256 , uint256 ) view returns(bytes32)
func (_Relay *RelayCallerSession) MerkleRoots(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _Relay.Contract.MerkleRoots(&_Relay.CallOpts, arg0, arg1)
}

// SigningPolicySetter is a free data retrieval call binding the contract method 0xa9dbe8ed.
//
// Solidity: function signingPolicySetter() view returns(address)
func (_Relay *RelayCaller) SigningPolicySetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "signingPolicySetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigningPolicySetter is a free data retrieval call binding the contract method 0xa9dbe8ed.
//
// Solidity: function signingPolicySetter() view returns(address)
func (_Relay *RelaySession) SigningPolicySetter() (common.Address, error) {
	return _Relay.Contract.SigningPolicySetter(&_Relay.CallOpts)
}

// SigningPolicySetter is a free data retrieval call binding the contract method 0xa9dbe8ed.
//
// Solidity: function signingPolicySetter() view returns(address)
func (_Relay *RelayCallerSession) SigningPolicySetter() (common.Address, error) {
	return _Relay.Contract.SigningPolicySetter(&_Relay.CallOpts)
}

// StateData is a free data retrieval call binding the contract method 0x1e8fb36a.
//
// Solidity: function stateData() view returns(uint8 randomNumberProtocolId, uint32 firstVotingRoundStartTs, uint8 votingEpochDurationSeconds, uint32 firstRewardEpochStartVotingRoundId, uint16 rewardEpochDurationInVotingEpochs, uint16 thresholdIncreaseBIPS, uint32 randomVotingRoundId, bool randomNumberQualityScore)
func (_Relay *RelayCaller) StateData(opts *bind.CallOpts) (struct {
	RandomNumberProtocolId             uint8
	FirstVotingRoundStartTs            uint32
	VotingEpochDurationSeconds         uint8
	FirstRewardEpochStartVotingRoundId uint32
	RewardEpochDurationInVotingEpochs  uint16
	ThresholdIncreaseBIPS              uint16
	RandomVotingRoundId                uint32
	RandomNumberQualityScore           bool
}, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "stateData")

	outstruct := new(struct {
		RandomNumberProtocolId             uint8
		FirstVotingRoundStartTs            uint32
		VotingEpochDurationSeconds         uint8
		FirstRewardEpochStartVotingRoundId uint32
		RewardEpochDurationInVotingEpochs  uint16
		ThresholdIncreaseBIPS              uint16
		RandomVotingRoundId                uint32
		RandomNumberQualityScore           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RandomNumberProtocolId = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.FirstVotingRoundStartTs = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.VotingEpochDurationSeconds = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.FirstRewardEpochStartVotingRoundId = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.RewardEpochDurationInVotingEpochs = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.ThresholdIncreaseBIPS = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.RandomVotingRoundId = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.RandomNumberQualityScore = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// StateData is a free data retrieval call binding the contract method 0x1e8fb36a.
//
// Solidity: function stateData() view returns(uint8 randomNumberProtocolId, uint32 firstVotingRoundStartTs, uint8 votingEpochDurationSeconds, uint32 firstRewardEpochStartVotingRoundId, uint16 rewardEpochDurationInVotingEpochs, uint16 thresholdIncreaseBIPS, uint32 randomVotingRoundId, bool randomNumberQualityScore)
func (_Relay *RelaySession) StateData() (struct {
	RandomNumberProtocolId             uint8
	FirstVotingRoundStartTs            uint32
	VotingEpochDurationSeconds         uint8
	FirstRewardEpochStartVotingRoundId uint32
	RewardEpochDurationInVotingEpochs  uint16
	ThresholdIncreaseBIPS              uint16
	RandomVotingRoundId                uint32
	RandomNumberQualityScore           bool
}, error) {
	return _Relay.Contract.StateData(&_Relay.CallOpts)
}

// StateData is a free data retrieval call binding the contract method 0x1e8fb36a.
//
// Solidity: function stateData() view returns(uint8 randomNumberProtocolId, uint32 firstVotingRoundStartTs, uint8 votingEpochDurationSeconds, uint32 firstRewardEpochStartVotingRoundId, uint16 rewardEpochDurationInVotingEpochs, uint16 thresholdIncreaseBIPS, uint32 randomVotingRoundId, bool randomNumberQualityScore)
func (_Relay *RelayCallerSession) StateData() (struct {
	RandomNumberProtocolId             uint8
	FirstVotingRoundStartTs            uint32
	VotingEpochDurationSeconds         uint8
	FirstRewardEpochStartVotingRoundId uint32
	RewardEpochDurationInVotingEpochs  uint16
	ThresholdIncreaseBIPS              uint16
	RandomVotingRoundId                uint32
	RandomNumberQualityScore           bool
}, error) {
	return _Relay.Contract.StateData(&_Relay.CallOpts)
}

// ToSigningPolicyHash is a free data retrieval call binding the contract method 0x0c85bf07.
//
// Solidity: function toSigningPolicyHash(uint256 ) view returns(bytes32)
func (_Relay *RelayCaller) ToSigningPolicyHash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Relay.contract.Call(opts, &out, "toSigningPolicyHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ToSigningPolicyHash is a free data retrieval call binding the contract method 0x0c85bf07.
//
// Solidity: function toSigningPolicyHash(uint256 ) view returns(bytes32)
func (_Relay *RelaySession) ToSigningPolicyHash(arg0 *big.Int) ([32]byte, error) {
	return _Relay.Contract.ToSigningPolicyHash(&_Relay.CallOpts, arg0)
}

// ToSigningPolicyHash is a free data retrieval call binding the contract method 0x0c85bf07.
//
// Solidity: function toSigningPolicyHash(uint256 ) view returns(bytes32)
func (_Relay *RelayCallerSession) ToSigningPolicyHash(arg0 *big.Int) ([32]byte, error) {
	return _Relay.Contract.ToSigningPolicyHash(&_Relay.CallOpts, arg0)
}

// Relay is a paid mutator transaction binding the contract method 0xb59589d1.
//
// Solidity: function relay() returns(uint256 _result)
func (_Relay *RelayTransactor) Relay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "relay")
}

// Relay is a paid mutator transaction binding the contract method 0xb59589d1.
//
// Solidity: function relay() returns(uint256 _result)
func (_Relay *RelaySession) Relay() (*types.Transaction, error) {
	return _Relay.Contract.Relay(&_Relay.TransactOpts)
}

// Relay is a paid mutator transaction binding the contract method 0xb59589d1.
//
// Solidity: function relay() returns(uint256 _result)
func (_Relay *RelayTransactorSession) Relay() (*types.Transaction, error) {
	return _Relay.Contract.Relay(&_Relay.TransactOpts)
}

// SetSigningPolicy is a paid mutator transaction binding the contract method 0x83534125.
//
// Solidity: function setSigningPolicy((uint24,uint32,uint16,uint256,address[],uint16[]) _signingPolicy) returns(bytes32)
func (_Relay *RelayTransactor) SetSigningPolicy(opts *bind.TransactOpts, _signingPolicy RelaySigningPolicy) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "setSigningPolicy", _signingPolicy)
}

// SetSigningPolicy is a paid mutator transaction binding the contract method 0x83534125.
//
// Solidity: function setSigningPolicy((uint24,uint32,uint16,uint256,address[],uint16[]) _signingPolicy) returns(bytes32)
func (_Relay *RelaySession) SetSigningPolicy(_signingPolicy RelaySigningPolicy) (*types.Transaction, error) {
	return _Relay.Contract.SetSigningPolicy(&_Relay.TransactOpts, _signingPolicy)
}

// SetSigningPolicy is a paid mutator transaction binding the contract method 0x83534125.
//
// Solidity: function setSigningPolicy((uint24,uint32,uint16,uint256,address[],uint16[]) _signingPolicy) returns(bytes32)
func (_Relay *RelayTransactorSession) SetSigningPolicy(_signingPolicy RelaySigningPolicy) (*types.Transaction, error) {
	return _Relay.Contract.SetSigningPolicy(&_Relay.TransactOpts, _signingPolicy)
}

// RelayProtocolMessageRelayedIterator is returned from FilterProtocolMessageRelayed and is used to iterate over the raw logs and unpacked data for ProtocolMessageRelayed events raised by the Relay contract.
type RelayProtocolMessageRelayedIterator struct {
	Event *RelayProtocolMessageRelayed // Event containing the contract specifics and raw log

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
func (it *RelayProtocolMessageRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayProtocolMessageRelayed)
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
		it.Event = new(RelayProtocolMessageRelayed)
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
func (it *RelayProtocolMessageRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayProtocolMessageRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayProtocolMessageRelayed represents a ProtocolMessageRelayed event raised by the Relay contract.
type RelayProtocolMessageRelayed struct {
	ProtocolId         uint8
	VotingRoundId      uint32
	RandomQualityScore bool
	MerkleRoot         [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProtocolMessageRelayed is a free log retrieval operation binding the contract event 0x4b781cfef3123d9257ab69e6e8ea36ad75a346d63c5ecf8a46931a0eef48bb9e.
//
// Solidity: event ProtocolMessageRelayed(uint8 indexed protocolId, uint32 indexed votingRoundId, bool randomQualityScore, bytes32 merkleRoot)
func (_Relay *RelayFilterer) FilterProtocolMessageRelayed(opts *bind.FilterOpts, protocolId []uint8, votingRoundId []uint32) (*RelayProtocolMessageRelayedIterator, error) {

	var protocolIdRule []interface{}
	for _, protocolIdItem := range protocolId {
		protocolIdRule = append(protocolIdRule, protocolIdItem)
	}
	var votingRoundIdRule []interface{}
	for _, votingRoundIdItem := range votingRoundId {
		votingRoundIdRule = append(votingRoundIdRule, votingRoundIdItem)
	}

	logs, sub, err := _Relay.contract.FilterLogs(opts, "ProtocolMessageRelayed", protocolIdRule, votingRoundIdRule)
	if err != nil {
		return nil, err
	}
	return &RelayProtocolMessageRelayedIterator{contract: _Relay.contract, event: "ProtocolMessageRelayed", logs: logs, sub: sub}, nil
}

// WatchProtocolMessageRelayed is a free log subscription operation binding the contract event 0x4b781cfef3123d9257ab69e6e8ea36ad75a346d63c5ecf8a46931a0eef48bb9e.
//
// Solidity: event ProtocolMessageRelayed(uint8 indexed protocolId, uint32 indexed votingRoundId, bool randomQualityScore, bytes32 merkleRoot)
func (_Relay *RelayFilterer) WatchProtocolMessageRelayed(opts *bind.WatchOpts, sink chan<- *RelayProtocolMessageRelayed, protocolId []uint8, votingRoundId []uint32) (event.Subscription, error) {

	var protocolIdRule []interface{}
	for _, protocolIdItem := range protocolId {
		protocolIdRule = append(protocolIdRule, protocolIdItem)
	}
	var votingRoundIdRule []interface{}
	for _, votingRoundIdItem := range votingRoundId {
		votingRoundIdRule = append(votingRoundIdRule, votingRoundIdItem)
	}

	logs, sub, err := _Relay.contract.WatchLogs(opts, "ProtocolMessageRelayed", protocolIdRule, votingRoundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayProtocolMessageRelayed)
				if err := _Relay.contract.UnpackLog(event, "ProtocolMessageRelayed", log); err != nil {
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

// ParseProtocolMessageRelayed is a log parse operation binding the contract event 0x4b781cfef3123d9257ab69e6e8ea36ad75a346d63c5ecf8a46931a0eef48bb9e.
//
// Solidity: event ProtocolMessageRelayed(uint8 indexed protocolId, uint32 indexed votingRoundId, bool randomQualityScore, bytes32 merkleRoot)
func (_Relay *RelayFilterer) ParseProtocolMessageRelayed(log types.Log) (*RelayProtocolMessageRelayed, error) {
	event := new(RelayProtocolMessageRelayed)
	if err := _Relay.contract.UnpackLog(event, "ProtocolMessageRelayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelaySigningPolicyInitializedIterator is returned from FilterSigningPolicyInitialized and is used to iterate over the raw logs and unpacked data for SigningPolicyInitialized events raised by the Relay contract.
type RelaySigningPolicyInitializedIterator struct {
	Event *RelaySigningPolicyInitialized // Event containing the contract specifics and raw log

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
func (it *RelaySigningPolicyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelaySigningPolicyInitialized)
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
		it.Event = new(RelaySigningPolicyInitialized)
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
func (it *RelaySigningPolicyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelaySigningPolicyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelaySigningPolicyInitialized represents a SigningPolicyInitialized event raised by the Relay contract.
type RelaySigningPolicyInitialized struct {
	RewardEpochId      *big.Int
	StartVotingRoundId uint32
	Threshold          uint16
	Seed               *big.Int
	Voters             []common.Address
	Weights            []uint16
	SigningPolicyBytes []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSigningPolicyInitialized is a free log retrieval operation binding the contract event 0x514b5927f6785249c3276eacfe1589dba1500895090675eb532fb08c58b0feb4.
//
// Solidity: event SigningPolicyInitialized(uint24 rewardEpochId, uint32 startVotingRoundId, uint16 threshold, uint256 seed, address[] voters, uint16[] weights, bytes signingPolicyBytes)
func (_Relay *RelayFilterer) FilterSigningPolicyInitialized(opts *bind.FilterOpts) (*RelaySigningPolicyInitializedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "SigningPolicyInitialized")
	if err != nil {
		return nil, err
	}
	return &RelaySigningPolicyInitializedIterator{contract: _Relay.contract, event: "SigningPolicyInitialized", logs: logs, sub: sub}, nil
}

// WatchSigningPolicyInitialized is a free log subscription operation binding the contract event 0x514b5927f6785249c3276eacfe1589dba1500895090675eb532fb08c58b0feb4.
//
// Solidity: event SigningPolicyInitialized(uint24 rewardEpochId, uint32 startVotingRoundId, uint16 threshold, uint256 seed, address[] voters, uint16[] weights, bytes signingPolicyBytes)
func (_Relay *RelayFilterer) WatchSigningPolicyInitialized(opts *bind.WatchOpts, sink chan<- *RelaySigningPolicyInitialized) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "SigningPolicyInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelaySigningPolicyInitialized)
				if err := _Relay.contract.UnpackLog(event, "SigningPolicyInitialized", log); err != nil {
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

// ParseSigningPolicyInitialized is a log parse operation binding the contract event 0x514b5927f6785249c3276eacfe1589dba1500895090675eb532fb08c58b0feb4.
//
// Solidity: event SigningPolicyInitialized(uint24 rewardEpochId, uint32 startVotingRoundId, uint16 threshold, uint256 seed, address[] voters, uint16[] weights, bytes signingPolicyBytes)
func (_Relay *RelayFilterer) ParseSigningPolicyInitialized(log types.Log) (*RelaySigningPolicyInitialized, error) {
	event := new(RelaySigningPolicyInitialized)
	if err := _Relay.contract.UnpackLog(event, "SigningPolicyInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelaySigningPolicyRelayedIterator is returned from FilterSigningPolicyRelayed and is used to iterate over the raw logs and unpacked data for SigningPolicyRelayed events raised by the Relay contract.
type RelaySigningPolicyRelayedIterator struct {
	Event *RelaySigningPolicyRelayed // Event containing the contract specifics and raw log

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
func (it *RelaySigningPolicyRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelaySigningPolicyRelayed)
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
		it.Event = new(RelaySigningPolicyRelayed)
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
func (it *RelaySigningPolicyRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelaySigningPolicyRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelaySigningPolicyRelayed represents a SigningPolicyRelayed event raised by the Relay contract.
type RelaySigningPolicyRelayed struct {
	RewardEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSigningPolicyRelayed is a free log retrieval operation binding the contract event 0xe68f222ab8e81b2e0b38a4725817a1846aeee9a4a11f55899e83fc20766175e8.
//
// Solidity: event SigningPolicyRelayed(uint256 indexed rewardEpochId)
func (_Relay *RelayFilterer) FilterSigningPolicyRelayed(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*RelaySigningPolicyRelayedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _Relay.contract.FilterLogs(opts, "SigningPolicyRelayed", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &RelaySigningPolicyRelayedIterator{contract: _Relay.contract, event: "SigningPolicyRelayed", logs: logs, sub: sub}, nil
}

// WatchSigningPolicyRelayed is a free log subscription operation binding the contract event 0xe68f222ab8e81b2e0b38a4725817a1846aeee9a4a11f55899e83fc20766175e8.
//
// Solidity: event SigningPolicyRelayed(uint256 indexed rewardEpochId)
func (_Relay *RelayFilterer) WatchSigningPolicyRelayed(opts *bind.WatchOpts, sink chan<- *RelaySigningPolicyRelayed, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _Relay.contract.WatchLogs(opts, "SigningPolicyRelayed", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelaySigningPolicyRelayed)
				if err := _Relay.contract.UnpackLog(event, "SigningPolicyRelayed", log); err != nil {
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

// ParseSigningPolicyRelayed is a log parse operation binding the contract event 0xe68f222ab8e81b2e0b38a4725817a1846aeee9a4a11f55899e83fc20766175e8.
//
// Solidity: event SigningPolicyRelayed(uint256 indexed rewardEpochId)
func (_Relay *RelayFilterer) ParseSigningPolicyRelayed(log types.Log) (*RelaySigningPolicyRelayed, error) {
	event := new(RelaySigningPolicyRelayed)
	if err := _Relay.contract.UnpackLog(event, "SigningPolicyRelayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
