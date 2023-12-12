// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voting

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

// IPChainStakeMirrorMultiSigVotingPChainVotes is an auto generated low-level Go binding around an user-defined struct.
type IPChainStakeMirrorMultiSigVotingPChainVotes struct {
	MerkleRoot [32]byte
	Votes      []common.Address
}

// VotingMetaData contains all meta data concerning the Voting contract.
var VotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_firstEpochStartTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epochDurationSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_votingThreshold\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes20[]\",\"name\":\"nodeIds\",\"type\":\"bytes20[]\"}],\"name\":\"PChainStakeMirrorValidatorUptimeVoteSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"PChainStakeMirrorVoteSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"}],\"name\":\"PChainStakeMirrorVotersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"PChainStakeMirrorVotingFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"PChainStakeMirrorVotingReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votingThreshold\",\"type\":\"uint256\"}],\"name\":\"PChainStakeMirrorVotingThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_newVotersList\",\"type\":\"address[]\"}],\"name\":\"changeVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpochConfiguration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_firstEpochStartTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epochDurationSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVoters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"votes\",\"type\":\"address[]\"}],\"internalType\":\"structIPChainStakeMirrorMultiSigVoting.PChainVotes[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotingThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"resetVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_votingThreshold\",\"type\":\"uint256\"}],\"name\":\"setVotingThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"shouldVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"},{\"internalType\":\"bytes20[]\",\"name\":\"_nodeIds\",\"type\":\"bytes20[]\"}],\"name\":\"submitValidatorUptimeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"submitVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VotingABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingMetaData.ABI instead.
var VotingABI = VotingMetaData.ABI

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentEpochId is a free data retrieval call binding the contract method 0xa29a839f.
//
// Solidity: function getCurrentEpochId() view returns(uint256)
func (_Voting *VotingCaller) GetCurrentEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getCurrentEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpochId is a free data retrieval call binding the contract method 0xa29a839f.
//
// Solidity: function getCurrentEpochId() view returns(uint256)
func (_Voting *VotingSession) GetCurrentEpochId() (*big.Int, error) {
	return _Voting.Contract.GetCurrentEpochId(&_Voting.CallOpts)
}

// GetCurrentEpochId is a free data retrieval call binding the contract method 0xa29a839f.
//
// Solidity: function getCurrentEpochId() view returns(uint256)
func (_Voting *VotingCallerSession) GetCurrentEpochId() (*big.Int, error) {
	return _Voting.Contract.GetCurrentEpochId(&_Voting.CallOpts)
}

// GetEpochConfiguration is a free data retrieval call binding the contract method 0xd72a2005.
//
// Solidity: function getEpochConfiguration() view returns(uint256 _firstEpochStartTs, uint256 _epochDurationSeconds)
func (_Voting *VotingCaller) GetEpochConfiguration(opts *bind.CallOpts) (struct {
	FirstEpochStartTs    *big.Int
	EpochDurationSeconds *big.Int
}, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getEpochConfiguration")

	outstruct := new(struct {
		FirstEpochStartTs    *big.Int
		EpochDurationSeconds *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FirstEpochStartTs = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EpochDurationSeconds = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetEpochConfiguration is a free data retrieval call binding the contract method 0xd72a2005.
//
// Solidity: function getEpochConfiguration() view returns(uint256 _firstEpochStartTs, uint256 _epochDurationSeconds)
func (_Voting *VotingSession) GetEpochConfiguration() (struct {
	FirstEpochStartTs    *big.Int
	EpochDurationSeconds *big.Int
}, error) {
	return _Voting.Contract.GetEpochConfiguration(&_Voting.CallOpts)
}

// GetEpochConfiguration is a free data retrieval call binding the contract method 0xd72a2005.
//
// Solidity: function getEpochConfiguration() view returns(uint256 _firstEpochStartTs, uint256 _epochDurationSeconds)
func (_Voting *VotingCallerSession) GetEpochConfiguration() (struct {
	FirstEpochStartTs    *big.Int
	EpochDurationSeconds *big.Int
}, error) {
	return _Voting.Contract.GetEpochConfiguration(&_Voting.CallOpts)
}

// GetEpochId is a free data retrieval call binding the contract method 0x5303548b.
//
// Solidity: function getEpochId(uint256 _timestamp) view returns(uint256)
func (_Voting *VotingCaller) GetEpochId(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getEpochId", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochId is a free data retrieval call binding the contract method 0x5303548b.
//
// Solidity: function getEpochId(uint256 _timestamp) view returns(uint256)
func (_Voting *VotingSession) GetEpochId(_timestamp *big.Int) (*big.Int, error) {
	return _Voting.Contract.GetEpochId(&_Voting.CallOpts, _timestamp)
}

// GetEpochId is a free data retrieval call binding the contract method 0x5303548b.
//
// Solidity: function getEpochId(uint256 _timestamp) view returns(uint256)
func (_Voting *VotingCallerSession) GetEpochId(_timestamp *big.Int) (*big.Int, error) {
	return _Voting.Contract.GetEpochId(&_Voting.CallOpts, _timestamp)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x0aab8ba5.
//
// Solidity: function getMerkleRoot(uint256 _epochId) view returns(bytes32)
func (_Voting *VotingCaller) GetMerkleRoot(opts *bind.CallOpts, _epochId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getMerkleRoot", _epochId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x0aab8ba5.
//
// Solidity: function getMerkleRoot(uint256 _epochId) view returns(bytes32)
func (_Voting *VotingSession) GetMerkleRoot(_epochId *big.Int) ([32]byte, error) {
	return _Voting.Contract.GetMerkleRoot(&_Voting.CallOpts, _epochId)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x0aab8ba5.
//
// Solidity: function getMerkleRoot(uint256 _epochId) view returns(bytes32)
func (_Voting *VotingCallerSession) GetMerkleRoot(_epochId *big.Int) ([32]byte, error) {
	return _Voting.Contract.GetMerkleRoot(&_Voting.CallOpts, _epochId)
}

// GetVoters is a free data retrieval call binding the contract method 0xcdd72253.
//
// Solidity: function getVoters() view returns(address[])
func (_Voting *VotingCaller) GetVoters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getVoters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetVoters is a free data retrieval call binding the contract method 0xcdd72253.
//
// Solidity: function getVoters() view returns(address[])
func (_Voting *VotingSession) GetVoters() ([]common.Address, error) {
	return _Voting.Contract.GetVoters(&_Voting.CallOpts)
}

// GetVoters is a free data retrieval call binding the contract method 0xcdd72253.
//
// Solidity: function getVoters() view returns(address[])
func (_Voting *VotingCallerSession) GetVoters() ([]common.Address, error) {
	return _Voting.Contract.GetVoters(&_Voting.CallOpts)
}

// GetVotes is a free data retrieval call binding the contract method 0xff981099.
//
// Solidity: function getVotes(uint256 _epochId) view returns((bytes32,address[])[])
func (_Voting *VotingCaller) GetVotes(opts *bind.CallOpts, _epochId *big.Int) ([]IPChainStakeMirrorMultiSigVotingPChainVotes, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getVotes", _epochId)

	if err != nil {
		return *new([]IPChainStakeMirrorMultiSigVotingPChainVotes), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPChainStakeMirrorMultiSigVotingPChainVotes)).(*[]IPChainStakeMirrorMultiSigVotingPChainVotes)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0xff981099.
//
// Solidity: function getVotes(uint256 _epochId) view returns((bytes32,address[])[])
func (_Voting *VotingSession) GetVotes(_epochId *big.Int) ([]IPChainStakeMirrorMultiSigVotingPChainVotes, error) {
	return _Voting.Contract.GetVotes(&_Voting.CallOpts, _epochId)
}

// GetVotes is a free data retrieval call binding the contract method 0xff981099.
//
// Solidity: function getVotes(uint256 _epochId) view returns((bytes32,address[])[])
func (_Voting *VotingCallerSession) GetVotes(_epochId *big.Int) ([]IPChainStakeMirrorMultiSigVotingPChainVotes, error) {
	return _Voting.Contract.GetVotes(&_Voting.CallOpts, _epochId)
}

// GetVotingThreshold is a free data retrieval call binding the contract method 0x3d7e4e30.
//
// Solidity: function getVotingThreshold() view returns(uint256)
func (_Voting *VotingCaller) GetVotingThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getVotingThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotingThreshold is a free data retrieval call binding the contract method 0x3d7e4e30.
//
// Solidity: function getVotingThreshold() view returns(uint256)
func (_Voting *VotingSession) GetVotingThreshold() (*big.Int, error) {
	return _Voting.Contract.GetVotingThreshold(&_Voting.CallOpts)
}

// GetVotingThreshold is a free data retrieval call binding the contract method 0x3d7e4e30.
//
// Solidity: function getVotingThreshold() view returns(uint256)
func (_Voting *VotingCallerSession) GetVotingThreshold() (*big.Int, error) {
	return _Voting.Contract.GetVotingThreshold(&_Voting.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Voting *VotingCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Voting *VotingSession) Governance() (common.Address, error) {
	return _Voting.Contract.Governance(&_Voting.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Voting *VotingCallerSession) Governance() (common.Address, error) {
	return _Voting.Contract.Governance(&_Voting.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Voting *VotingCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Voting *VotingSession) GovernanceSettings() (common.Address, error) {
	return _Voting.Contract.GovernanceSettings(&_Voting.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Voting *VotingCallerSession) GovernanceSettings() (common.Address, error) {
	return _Voting.Contract.GovernanceSettings(&_Voting.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Voting *VotingCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Voting *VotingSession) ProductionMode() (bool, error) {
	return _Voting.Contract.ProductionMode(&_Voting.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Voting *VotingCallerSession) ProductionMode() (bool, error) {
	return _Voting.Contract.ProductionMode(&_Voting.CallOpts)
}

// ShouldVote is a free data retrieval call binding the contract method 0x59fa6d59.
//
// Solidity: function shouldVote(uint256 _epochId, address _voter) view returns(bool)
func (_Voting *VotingCaller) ShouldVote(opts *bind.CallOpts, _epochId *big.Int, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "shouldVote", _epochId, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShouldVote is a free data retrieval call binding the contract method 0x59fa6d59.
//
// Solidity: function shouldVote(uint256 _epochId, address _voter) view returns(bool)
func (_Voting *VotingSession) ShouldVote(_epochId *big.Int, _voter common.Address) (bool, error) {
	return _Voting.Contract.ShouldVote(&_Voting.CallOpts, _epochId, _voter)
}

// ShouldVote is a free data retrieval call binding the contract method 0x59fa6d59.
//
// Solidity: function shouldVote(uint256 _epochId, address _voter) view returns(bool)
func (_Voting *VotingCallerSession) ShouldVote(_epochId *big.Int, _voter common.Address) (bool, error) {
	return _Voting.Contract.ShouldVote(&_Voting.CallOpts, _epochId, _voter)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Voting *VotingCaller) TimelockedCalls(opts *bind.CallOpts, arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "timelockedCalls", arg0)

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
func (_Voting *VotingSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _Voting.Contract.TimelockedCalls(&_Voting.CallOpts, arg0)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Voting *VotingCallerSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _Voting.Contract.TimelockedCalls(&_Voting.CallOpts, arg0)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Voting *VotingTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Voting *VotingSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Voting.Contract.CancelGovernanceCall(&_Voting.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Voting *VotingTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Voting.Contract.CancelGovernanceCall(&_Voting.TransactOpts, _selector)
}

// ChangeVoters is a paid mutator transaction binding the contract method 0x1e98e4fa.
//
// Solidity: function changeVoters(address[] _newVotersList) returns()
func (_Voting *VotingTransactor) ChangeVoters(opts *bind.TransactOpts, _newVotersList []common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "changeVoters", _newVotersList)
}

// ChangeVoters is a paid mutator transaction binding the contract method 0x1e98e4fa.
//
// Solidity: function changeVoters(address[] _newVotersList) returns()
func (_Voting *VotingSession) ChangeVoters(_newVotersList []common.Address) (*types.Transaction, error) {
	return _Voting.Contract.ChangeVoters(&_Voting.TransactOpts, _newVotersList)
}

// ChangeVoters is a paid mutator transaction binding the contract method 0x1e98e4fa.
//
// Solidity: function changeVoters(address[] _newVotersList) returns()
func (_Voting *VotingTransactorSession) ChangeVoters(_newVotersList []common.Address) (*types.Transaction, error) {
	return _Voting.Contract.ChangeVoters(&_Voting.TransactOpts, _newVotersList)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Voting *VotingTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Voting *VotingSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Voting.Contract.ExecuteGovernanceCall(&_Voting.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Voting *VotingTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Voting.Contract.ExecuteGovernanceCall(&_Voting.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _initialGovernance) returns()
func (_Voting *VotingTransactor) Initialise(opts *bind.TransactOpts, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "initialise", _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _initialGovernance) returns()
func (_Voting *VotingSession) Initialise(_initialGovernance common.Address) (*types.Transaction, error) {
	return _Voting.Contract.Initialise(&_Voting.TransactOpts, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _initialGovernance) returns()
func (_Voting *VotingTransactorSession) Initialise(_initialGovernance common.Address) (*types.Transaction, error) {
	return _Voting.Contract.Initialise(&_Voting.TransactOpts, _initialGovernance)
}

// ResetVoting is a paid mutator transaction binding the contract method 0x8fae2d0a.
//
// Solidity: function resetVoting(uint256 _epochId) returns()
func (_Voting *VotingTransactor) ResetVoting(opts *bind.TransactOpts, _epochId *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "resetVoting", _epochId)
}

// ResetVoting is a paid mutator transaction binding the contract method 0x8fae2d0a.
//
// Solidity: function resetVoting(uint256 _epochId) returns()
func (_Voting *VotingSession) ResetVoting(_epochId *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.ResetVoting(&_Voting.TransactOpts, _epochId)
}

// ResetVoting is a paid mutator transaction binding the contract method 0x8fae2d0a.
//
// Solidity: function resetVoting(uint256 _epochId) returns()
func (_Voting *VotingTransactorSession) ResetVoting(_epochId *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.ResetVoting(&_Voting.TransactOpts, _epochId)
}

// SetVotingThreshold is a paid mutator transaction binding the contract method 0x836761e0.
//
// Solidity: function setVotingThreshold(uint256 _votingThreshold) returns()
func (_Voting *VotingTransactor) SetVotingThreshold(opts *bind.TransactOpts, _votingThreshold *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "setVotingThreshold", _votingThreshold)
}

// SetVotingThreshold is a paid mutator transaction binding the contract method 0x836761e0.
//
// Solidity: function setVotingThreshold(uint256 _votingThreshold) returns()
func (_Voting *VotingSession) SetVotingThreshold(_votingThreshold *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.SetVotingThreshold(&_Voting.TransactOpts, _votingThreshold)
}

// SetVotingThreshold is a paid mutator transaction binding the contract method 0x836761e0.
//
// Solidity: function setVotingThreshold(uint256 _votingThreshold) returns()
func (_Voting *VotingTransactorSession) SetVotingThreshold(_votingThreshold *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.SetVotingThreshold(&_Voting.TransactOpts, _votingThreshold)
}

// SubmitValidatorUptimeVote is a paid mutator transaction binding the contract method 0x5f9f3fd9.
//
// Solidity: function submitValidatorUptimeVote(uint256 _rewardEpochId, bytes20[] _nodeIds) returns()
func (_Voting *VotingTransactor) SubmitValidatorUptimeVote(opts *bind.TransactOpts, _rewardEpochId *big.Int, _nodeIds [][20]byte) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "submitValidatorUptimeVote", _rewardEpochId, _nodeIds)
}

// SubmitValidatorUptimeVote is a paid mutator transaction binding the contract method 0x5f9f3fd9.
//
// Solidity: function submitValidatorUptimeVote(uint256 _rewardEpochId, bytes20[] _nodeIds) returns()
func (_Voting *VotingSession) SubmitValidatorUptimeVote(_rewardEpochId *big.Int, _nodeIds [][20]byte) (*types.Transaction, error) {
	return _Voting.Contract.SubmitValidatorUptimeVote(&_Voting.TransactOpts, _rewardEpochId, _nodeIds)
}

// SubmitValidatorUptimeVote is a paid mutator transaction binding the contract method 0x5f9f3fd9.
//
// Solidity: function submitValidatorUptimeVote(uint256 _rewardEpochId, bytes20[] _nodeIds) returns()
func (_Voting *VotingTransactorSession) SubmitValidatorUptimeVote(_rewardEpochId *big.Int, _nodeIds [][20]byte) (*types.Transaction, error) {
	return _Voting.Contract.SubmitValidatorUptimeVote(&_Voting.TransactOpts, _rewardEpochId, _nodeIds)
}

// SubmitVote is a paid mutator transaction binding the contract method 0xac8f38c8.
//
// Solidity: function submitVote(uint256 _epochId, bytes32 _merkleRoot) returns()
func (_Voting *VotingTransactor) SubmitVote(opts *bind.TransactOpts, _epochId *big.Int, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "submitVote", _epochId, _merkleRoot)
}

// SubmitVote is a paid mutator transaction binding the contract method 0xac8f38c8.
//
// Solidity: function submitVote(uint256 _epochId, bytes32 _merkleRoot) returns()
func (_Voting *VotingSession) SubmitVote(_epochId *big.Int, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.SubmitVote(&_Voting.TransactOpts, _epochId, _merkleRoot)
}

// SubmitVote is a paid mutator transaction binding the contract method 0xac8f38c8.
//
// Solidity: function submitVote(uint256 _epochId, bytes32 _merkleRoot) returns()
func (_Voting *VotingTransactorSession) SubmitVote(_epochId *big.Int, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.SubmitVote(&_Voting.TransactOpts, _epochId, _merkleRoot)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Voting *VotingTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Voting *VotingSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _Voting.Contract.SwitchToProductionMode(&_Voting.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Voting *VotingTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _Voting.Contract.SwitchToProductionMode(&_Voting.TransactOpts)
}

// VotingGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the Voting contract.
type VotingGovernanceCallTimelockedIterator struct {
	Event *VotingGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *VotingGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingGovernanceCallTimelocked)
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
		it.Event = new(VotingGovernanceCallTimelocked)
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
func (it *VotingGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Voting contract.
type VotingGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Voting *VotingFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*VotingGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &VotingGovernanceCallTimelockedIterator{contract: _Voting.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Voting *VotingFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *VotingGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingGovernanceCallTimelocked)
				if err := _Voting.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_Voting *VotingFilterer) ParseGovernanceCallTimelocked(log types.Log) (*VotingGovernanceCallTimelocked, error) {
	event := new(VotingGovernanceCallTimelocked)
	if err := _Voting.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the Voting contract.
type VotingGovernanceInitialisedIterator struct {
	Event *VotingGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *VotingGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingGovernanceInitialised)
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
		it.Event = new(VotingGovernanceInitialised)
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
func (it *VotingGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingGovernanceInitialised represents a GovernanceInitialised event raised by the Voting contract.
type VotingGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_Voting *VotingFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*VotingGovernanceInitialisedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &VotingGovernanceInitialisedIterator{contract: _Voting.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_Voting *VotingFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *VotingGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingGovernanceInitialised)
				if err := _Voting.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_Voting *VotingFilterer) ParseGovernanceInitialised(log types.Log) (*VotingGovernanceInitialised, error) {
	event := new(VotingGovernanceInitialised)
	if err := _Voting.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the Voting contract.
type VotingGovernedProductionModeEnteredIterator struct {
	Event *VotingGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *VotingGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingGovernedProductionModeEntered)
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
		it.Event = new(VotingGovernedProductionModeEntered)
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
func (it *VotingGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the Voting contract.
type VotingGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_Voting *VotingFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*VotingGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &VotingGovernedProductionModeEnteredIterator{contract: _Voting.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_Voting *VotingFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *VotingGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingGovernedProductionModeEntered)
				if err := _Voting.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_Voting *VotingFilterer) ParseGovernedProductionModeEntered(log types.Log) (*VotingGovernedProductionModeEntered, error) {
	event := new(VotingGovernedProductionModeEntered)
	if err := _Voting.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingPChainStakeMirrorValidatorUptimeVoteSubmittedIterator is returned from FilterPChainStakeMirrorValidatorUptimeVoteSubmitted and is used to iterate over the raw logs and unpacked data for PChainStakeMirrorValidatorUptimeVoteSubmitted events raised by the Voting contract.
type VotingPChainStakeMirrorValidatorUptimeVoteSubmittedIterator struct {
	Event *VotingPChainStakeMirrorValidatorUptimeVoteSubmitted // Event containing the contract specifics and raw log

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
func (it *VotingPChainStakeMirrorValidatorUptimeVoteSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingPChainStakeMirrorValidatorUptimeVoteSubmitted)
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
		it.Event = new(VotingPChainStakeMirrorValidatorUptimeVoteSubmitted)
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
func (it *VotingPChainStakeMirrorValidatorUptimeVoteSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingPChainStakeMirrorValidatorUptimeVoteSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingPChainStakeMirrorValidatorUptimeVoteSubmitted represents a PChainStakeMirrorValidatorUptimeVoteSubmitted event raised by the Voting contract.
type VotingPChainStakeMirrorValidatorUptimeVoteSubmitted struct {
	RewardEpochId *big.Int
	Timestamp     *big.Int
	Voter         common.Address
	NodeIds       [][20]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPChainStakeMirrorValidatorUptimeVoteSubmitted is a free log retrieval operation binding the contract event 0x1dafa85eb81adeb6e48e08365cf90c9ade61034e1e9fa293dc9730c4a8f89b39.
//
// Solidity: event PChainStakeMirrorValidatorUptimeVoteSubmitted(uint256 indexed rewardEpochId, uint256 indexed timestamp, address voter, bytes20[] nodeIds)
func (_Voting *VotingFilterer) FilterPChainStakeMirrorValidatorUptimeVoteSubmitted(opts *bind.FilterOpts, rewardEpochId []*big.Int, timestamp []*big.Int) (*VotingPChainStakeMirrorValidatorUptimeVoteSubmittedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "PChainStakeMirrorValidatorUptimeVoteSubmitted", rewardEpochIdRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return &VotingPChainStakeMirrorValidatorUptimeVoteSubmittedIterator{contract: _Voting.contract, event: "PChainStakeMirrorValidatorUptimeVoteSubmitted", logs: logs, sub: sub}, nil
}

// WatchPChainStakeMirrorValidatorUptimeVoteSubmitted is a free log subscription operation binding the contract event 0x1dafa85eb81adeb6e48e08365cf90c9ade61034e1e9fa293dc9730c4a8f89b39.
//
// Solidity: event PChainStakeMirrorValidatorUptimeVoteSubmitted(uint256 indexed rewardEpochId, uint256 indexed timestamp, address voter, bytes20[] nodeIds)
func (_Voting *VotingFilterer) WatchPChainStakeMirrorValidatorUptimeVoteSubmitted(opts *bind.WatchOpts, sink chan<- *VotingPChainStakeMirrorValidatorUptimeVoteSubmitted, rewardEpochId []*big.Int, timestamp []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "PChainStakeMirrorValidatorUptimeVoteSubmitted", rewardEpochIdRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingPChainStakeMirrorValidatorUptimeVoteSubmitted)
				if err := _Voting.contract.UnpackLog(event, "PChainStakeMirrorValidatorUptimeVoteSubmitted", log); err != nil {
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

// ParsePChainStakeMirrorValidatorUptimeVoteSubmitted is a log parse operation binding the contract event 0x1dafa85eb81adeb6e48e08365cf90c9ade61034e1e9fa293dc9730c4a8f89b39.
//
// Solidity: event PChainStakeMirrorValidatorUptimeVoteSubmitted(uint256 indexed rewardEpochId, uint256 indexed timestamp, address voter, bytes20[] nodeIds)
func (_Voting *VotingFilterer) ParsePChainStakeMirrorValidatorUptimeVoteSubmitted(log types.Log) (*VotingPChainStakeMirrorValidatorUptimeVoteSubmitted, error) {
	event := new(VotingPChainStakeMirrorValidatorUptimeVoteSubmitted)
	if err := _Voting.contract.UnpackLog(event, "PChainStakeMirrorValidatorUptimeVoteSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingPChainStakeMirrorVoteSubmittedIterator is returned from FilterPChainStakeMirrorVoteSubmitted and is used to iterate over the raw logs and unpacked data for PChainStakeMirrorVoteSubmitted events raised by the Voting contract.
type VotingPChainStakeMirrorVoteSubmittedIterator struct {
	Event *VotingPChainStakeMirrorVoteSubmitted // Event containing the contract specifics and raw log

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
func (it *VotingPChainStakeMirrorVoteSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingPChainStakeMirrorVoteSubmitted)
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
		it.Event = new(VotingPChainStakeMirrorVoteSubmitted)
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
func (it *VotingPChainStakeMirrorVoteSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingPChainStakeMirrorVoteSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingPChainStakeMirrorVoteSubmitted represents a PChainStakeMirrorVoteSubmitted event raised by the Voting contract.
type VotingPChainStakeMirrorVoteSubmitted struct {
	EpochId    *big.Int
	Voter      common.Address
	MerkleRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPChainStakeMirrorVoteSubmitted is a free log retrieval operation binding the contract event 0xc20a6d1ff27e89f30517726862e720046602b33a06fbbe3a93aa75687b8419f4.
//
// Solidity: event PChainStakeMirrorVoteSubmitted(uint256 epochId, address voter, bytes32 merkleRoot)
func (_Voting *VotingFilterer) FilterPChainStakeMirrorVoteSubmitted(opts *bind.FilterOpts) (*VotingPChainStakeMirrorVoteSubmittedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "PChainStakeMirrorVoteSubmitted")
	if err != nil {
		return nil, err
	}
	return &VotingPChainStakeMirrorVoteSubmittedIterator{contract: _Voting.contract, event: "PChainStakeMirrorVoteSubmitted", logs: logs, sub: sub}, nil
}

// WatchPChainStakeMirrorVoteSubmitted is a free log subscription operation binding the contract event 0xc20a6d1ff27e89f30517726862e720046602b33a06fbbe3a93aa75687b8419f4.
//
// Solidity: event PChainStakeMirrorVoteSubmitted(uint256 epochId, address voter, bytes32 merkleRoot)
func (_Voting *VotingFilterer) WatchPChainStakeMirrorVoteSubmitted(opts *bind.WatchOpts, sink chan<- *VotingPChainStakeMirrorVoteSubmitted) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "PChainStakeMirrorVoteSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingPChainStakeMirrorVoteSubmitted)
				if err := _Voting.contract.UnpackLog(event, "PChainStakeMirrorVoteSubmitted", log); err != nil {
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

// ParsePChainStakeMirrorVoteSubmitted is a log parse operation binding the contract event 0xc20a6d1ff27e89f30517726862e720046602b33a06fbbe3a93aa75687b8419f4.
//
// Solidity: event PChainStakeMirrorVoteSubmitted(uint256 epochId, address voter, bytes32 merkleRoot)
func (_Voting *VotingFilterer) ParsePChainStakeMirrorVoteSubmitted(log types.Log) (*VotingPChainStakeMirrorVoteSubmitted, error) {
	event := new(VotingPChainStakeMirrorVoteSubmitted)
	if err := _Voting.contract.UnpackLog(event, "PChainStakeMirrorVoteSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingPChainStakeMirrorVotersSetIterator is returned from FilterPChainStakeMirrorVotersSet and is used to iterate over the raw logs and unpacked data for PChainStakeMirrorVotersSet events raised by the Voting contract.
type VotingPChainStakeMirrorVotersSetIterator struct {
	Event *VotingPChainStakeMirrorVotersSet // Event containing the contract specifics and raw log

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
func (it *VotingPChainStakeMirrorVotersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingPChainStakeMirrorVotersSet)
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
		it.Event = new(VotingPChainStakeMirrorVotersSet)
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
func (it *VotingPChainStakeMirrorVotersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingPChainStakeMirrorVotersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingPChainStakeMirrorVotersSet represents a PChainStakeMirrorVotersSet event raised by the Voting contract.
type VotingPChainStakeMirrorVotersSet struct {
	Voters []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPChainStakeMirrorVotersSet is a free log retrieval operation binding the contract event 0xee3c8349f0956bc41369bca45ae7ea8ab85e61bc09b93be0b36b405bd521fc22.
//
// Solidity: event PChainStakeMirrorVotersSet(address[] voters)
func (_Voting *VotingFilterer) FilterPChainStakeMirrorVotersSet(opts *bind.FilterOpts) (*VotingPChainStakeMirrorVotersSetIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "PChainStakeMirrorVotersSet")
	if err != nil {
		return nil, err
	}
	return &VotingPChainStakeMirrorVotersSetIterator{contract: _Voting.contract, event: "PChainStakeMirrorVotersSet", logs: logs, sub: sub}, nil
}

// WatchPChainStakeMirrorVotersSet is a free log subscription operation binding the contract event 0xee3c8349f0956bc41369bca45ae7ea8ab85e61bc09b93be0b36b405bd521fc22.
//
// Solidity: event PChainStakeMirrorVotersSet(address[] voters)
func (_Voting *VotingFilterer) WatchPChainStakeMirrorVotersSet(opts *bind.WatchOpts, sink chan<- *VotingPChainStakeMirrorVotersSet) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "PChainStakeMirrorVotersSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingPChainStakeMirrorVotersSet)
				if err := _Voting.contract.UnpackLog(event, "PChainStakeMirrorVotersSet", log); err != nil {
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

// ParsePChainStakeMirrorVotersSet is a log parse operation binding the contract event 0xee3c8349f0956bc41369bca45ae7ea8ab85e61bc09b93be0b36b405bd521fc22.
//
// Solidity: event PChainStakeMirrorVotersSet(address[] voters)
func (_Voting *VotingFilterer) ParsePChainStakeMirrorVotersSet(log types.Log) (*VotingPChainStakeMirrorVotersSet, error) {
	event := new(VotingPChainStakeMirrorVotersSet)
	if err := _Voting.contract.UnpackLog(event, "PChainStakeMirrorVotersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingPChainStakeMirrorVotingFinalizedIterator is returned from FilterPChainStakeMirrorVotingFinalized and is used to iterate over the raw logs and unpacked data for PChainStakeMirrorVotingFinalized events raised by the Voting contract.
type VotingPChainStakeMirrorVotingFinalizedIterator struct {
	Event *VotingPChainStakeMirrorVotingFinalized // Event containing the contract specifics and raw log

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
func (it *VotingPChainStakeMirrorVotingFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingPChainStakeMirrorVotingFinalized)
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
		it.Event = new(VotingPChainStakeMirrorVotingFinalized)
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
func (it *VotingPChainStakeMirrorVotingFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingPChainStakeMirrorVotingFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingPChainStakeMirrorVotingFinalized represents a PChainStakeMirrorVotingFinalized event raised by the Voting contract.
type VotingPChainStakeMirrorVotingFinalized struct {
	EpochId    *big.Int
	MerkleRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPChainStakeMirrorVotingFinalized is a free log retrieval operation binding the contract event 0x37bd3805c370e0a147273bf48789eabfcbe6b28ad2c3ff62c2f6f332e9cc8692.
//
// Solidity: event PChainStakeMirrorVotingFinalized(uint256 indexed epochId, bytes32 merkleRoot)
func (_Voting *VotingFilterer) FilterPChainStakeMirrorVotingFinalized(opts *bind.FilterOpts, epochId []*big.Int) (*VotingPChainStakeMirrorVotingFinalizedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "PChainStakeMirrorVotingFinalized", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &VotingPChainStakeMirrorVotingFinalizedIterator{contract: _Voting.contract, event: "PChainStakeMirrorVotingFinalized", logs: logs, sub: sub}, nil
}

// WatchPChainStakeMirrorVotingFinalized is a free log subscription operation binding the contract event 0x37bd3805c370e0a147273bf48789eabfcbe6b28ad2c3ff62c2f6f332e9cc8692.
//
// Solidity: event PChainStakeMirrorVotingFinalized(uint256 indexed epochId, bytes32 merkleRoot)
func (_Voting *VotingFilterer) WatchPChainStakeMirrorVotingFinalized(opts *bind.WatchOpts, sink chan<- *VotingPChainStakeMirrorVotingFinalized, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "PChainStakeMirrorVotingFinalized", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingPChainStakeMirrorVotingFinalized)
				if err := _Voting.contract.UnpackLog(event, "PChainStakeMirrorVotingFinalized", log); err != nil {
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

// ParsePChainStakeMirrorVotingFinalized is a log parse operation binding the contract event 0x37bd3805c370e0a147273bf48789eabfcbe6b28ad2c3ff62c2f6f332e9cc8692.
//
// Solidity: event PChainStakeMirrorVotingFinalized(uint256 indexed epochId, bytes32 merkleRoot)
func (_Voting *VotingFilterer) ParsePChainStakeMirrorVotingFinalized(log types.Log) (*VotingPChainStakeMirrorVotingFinalized, error) {
	event := new(VotingPChainStakeMirrorVotingFinalized)
	if err := _Voting.contract.UnpackLog(event, "PChainStakeMirrorVotingFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingPChainStakeMirrorVotingResetIterator is returned from FilterPChainStakeMirrorVotingReset and is used to iterate over the raw logs and unpacked data for PChainStakeMirrorVotingReset events raised by the Voting contract.
type VotingPChainStakeMirrorVotingResetIterator struct {
	Event *VotingPChainStakeMirrorVotingReset // Event containing the contract specifics and raw log

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
func (it *VotingPChainStakeMirrorVotingResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingPChainStakeMirrorVotingReset)
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
		it.Event = new(VotingPChainStakeMirrorVotingReset)
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
func (it *VotingPChainStakeMirrorVotingResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingPChainStakeMirrorVotingResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingPChainStakeMirrorVotingReset represents a PChainStakeMirrorVotingReset event raised by the Voting contract.
type VotingPChainStakeMirrorVotingReset struct {
	EpochId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPChainStakeMirrorVotingReset is a free log retrieval operation binding the contract event 0x72ce53d1ea24486d0db629df4e52d1c99b6a65b1c74f45ff479c23d34628a928.
//
// Solidity: event PChainStakeMirrorVotingReset(uint256 epochId)
func (_Voting *VotingFilterer) FilterPChainStakeMirrorVotingReset(opts *bind.FilterOpts) (*VotingPChainStakeMirrorVotingResetIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "PChainStakeMirrorVotingReset")
	if err != nil {
		return nil, err
	}
	return &VotingPChainStakeMirrorVotingResetIterator{contract: _Voting.contract, event: "PChainStakeMirrorVotingReset", logs: logs, sub: sub}, nil
}

// WatchPChainStakeMirrorVotingReset is a free log subscription operation binding the contract event 0x72ce53d1ea24486d0db629df4e52d1c99b6a65b1c74f45ff479c23d34628a928.
//
// Solidity: event PChainStakeMirrorVotingReset(uint256 epochId)
func (_Voting *VotingFilterer) WatchPChainStakeMirrorVotingReset(opts *bind.WatchOpts, sink chan<- *VotingPChainStakeMirrorVotingReset) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "PChainStakeMirrorVotingReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingPChainStakeMirrorVotingReset)
				if err := _Voting.contract.UnpackLog(event, "PChainStakeMirrorVotingReset", log); err != nil {
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

// ParsePChainStakeMirrorVotingReset is a log parse operation binding the contract event 0x72ce53d1ea24486d0db629df4e52d1c99b6a65b1c74f45ff479c23d34628a928.
//
// Solidity: event PChainStakeMirrorVotingReset(uint256 epochId)
func (_Voting *VotingFilterer) ParsePChainStakeMirrorVotingReset(log types.Log) (*VotingPChainStakeMirrorVotingReset, error) {
	event := new(VotingPChainStakeMirrorVotingReset)
	if err := _Voting.contract.UnpackLog(event, "PChainStakeMirrorVotingReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingPChainStakeMirrorVotingThresholdSetIterator is returned from FilterPChainStakeMirrorVotingThresholdSet and is used to iterate over the raw logs and unpacked data for PChainStakeMirrorVotingThresholdSet events raised by the Voting contract.
type VotingPChainStakeMirrorVotingThresholdSetIterator struct {
	Event *VotingPChainStakeMirrorVotingThresholdSet // Event containing the contract specifics and raw log

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
func (it *VotingPChainStakeMirrorVotingThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingPChainStakeMirrorVotingThresholdSet)
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
		it.Event = new(VotingPChainStakeMirrorVotingThresholdSet)
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
func (it *VotingPChainStakeMirrorVotingThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingPChainStakeMirrorVotingThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingPChainStakeMirrorVotingThresholdSet represents a PChainStakeMirrorVotingThresholdSet event raised by the Voting contract.
type VotingPChainStakeMirrorVotingThresholdSet struct {
	VotingThreshold *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPChainStakeMirrorVotingThresholdSet is a free log retrieval operation binding the contract event 0x79fcd9a44ea76361df1fbf61af58b8402734827d67758ccd4f6d1bf5ec1358b8.
//
// Solidity: event PChainStakeMirrorVotingThresholdSet(uint256 votingThreshold)
func (_Voting *VotingFilterer) FilterPChainStakeMirrorVotingThresholdSet(opts *bind.FilterOpts) (*VotingPChainStakeMirrorVotingThresholdSetIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "PChainStakeMirrorVotingThresholdSet")
	if err != nil {
		return nil, err
	}
	return &VotingPChainStakeMirrorVotingThresholdSetIterator{contract: _Voting.contract, event: "PChainStakeMirrorVotingThresholdSet", logs: logs, sub: sub}, nil
}

// WatchPChainStakeMirrorVotingThresholdSet is a free log subscription operation binding the contract event 0x79fcd9a44ea76361df1fbf61af58b8402734827d67758ccd4f6d1bf5ec1358b8.
//
// Solidity: event PChainStakeMirrorVotingThresholdSet(uint256 votingThreshold)
func (_Voting *VotingFilterer) WatchPChainStakeMirrorVotingThresholdSet(opts *bind.WatchOpts, sink chan<- *VotingPChainStakeMirrorVotingThresholdSet) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "PChainStakeMirrorVotingThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingPChainStakeMirrorVotingThresholdSet)
				if err := _Voting.contract.UnpackLog(event, "PChainStakeMirrorVotingThresholdSet", log); err != nil {
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

// ParsePChainStakeMirrorVotingThresholdSet is a log parse operation binding the contract event 0x79fcd9a44ea76361df1fbf61af58b8402734827d67758ccd4f6d1bf5ec1358b8.
//
// Solidity: event PChainStakeMirrorVotingThresholdSet(uint256 votingThreshold)
func (_Voting *VotingFilterer) ParsePChainStakeMirrorVotingThresholdSet(log types.Log) (*VotingPChainStakeMirrorVotingThresholdSet, error) {
	event := new(VotingPChainStakeMirrorVotingThresholdSet)
	if err := _Voting.contract.UnpackLog(event, "PChainStakeMirrorVotingThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the Voting contract.
type VotingTimelockedGovernanceCallCanceledIterator struct {
	Event *VotingTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *VotingTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingTimelockedGovernanceCallCanceled)
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
		it.Event = new(VotingTimelockedGovernanceCallCanceled)
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
func (it *VotingTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the Voting contract.
type VotingTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_Voting *VotingFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*VotingTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &VotingTimelockedGovernanceCallCanceledIterator{contract: _Voting.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_Voting *VotingFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *VotingTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingTimelockedGovernanceCallCanceled)
				if err := _Voting.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_Voting *VotingFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*VotingTimelockedGovernanceCallCanceled, error) {
	event := new(VotingTimelockedGovernanceCallCanceled)
	if err := _Voting.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the Voting contract.
type VotingTimelockedGovernanceCallExecutedIterator struct {
	Event *VotingTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *VotingTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingTimelockedGovernanceCallExecuted)
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
		it.Event = new(VotingTimelockedGovernanceCallExecuted)
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
func (it *VotingTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the Voting contract.
type VotingTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_Voting *VotingFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*VotingTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &VotingTimelockedGovernanceCallExecutedIterator{contract: _Voting.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_Voting *VotingFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *VotingTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingTimelockedGovernanceCallExecuted)
				if err := _Voting.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_Voting *VotingFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*VotingTimelockedGovernanceCallExecuted, error) {
	event := new(VotingTimelockedGovernanceCallExecuted)
	if err := _Voting.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
