// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package submission

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

// SubmissionMetaData contains all meta data concerning the Submission contract.
var SubmissionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_submit3MethodEnabled\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewVotingRoundInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemManager\",\"outputs\":[{\"internalType\":\"contractFlareSystemManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_submit1Addresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_submit2Addresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_submit3Addresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_submitSignaturesAddresses\",\"type\":\"address[]\"}],\"name\":\"initNewVotingRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setSubmit3MethodEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_submitAndPassContract\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_submitAndPassSelector\",\"type\":\"bytes4\"}],\"name\":\"setSubmitAndPassData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submit1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submit2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submit3\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submit3MethodEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"submitAndPass\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitAndPassContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitAndPassSelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitSignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SubmissionABI is the input ABI used to generate the binding from.
// Deprecated: Use SubmissionMetaData.ABI instead.
var SubmissionABI = SubmissionMetaData.ABI

// Submission is an auto generated Go binding around an Ethereum contract.
type Submission struct {
	SubmissionCaller     // Read-only binding to the contract
	SubmissionTransactor // Write-only binding to the contract
	SubmissionFilterer   // Log filterer for contract events
}

// SubmissionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubmissionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmissionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubmissionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmissionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubmissionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmissionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubmissionSession struct {
	Contract     *Submission       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubmissionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubmissionCallerSession struct {
	Contract *SubmissionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SubmissionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubmissionTransactorSession struct {
	Contract     *SubmissionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SubmissionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubmissionRaw struct {
	Contract *Submission // Generic contract binding to access the raw methods on
}

// SubmissionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubmissionCallerRaw struct {
	Contract *SubmissionCaller // Generic read-only contract binding to access the raw methods on
}

// SubmissionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubmissionTransactorRaw struct {
	Contract *SubmissionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubmission creates a new instance of Submission, bound to a specific deployed contract.
func NewSubmission(address common.Address, backend bind.ContractBackend) (*Submission, error) {
	contract, err := bindSubmission(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Submission{SubmissionCaller: SubmissionCaller{contract: contract}, SubmissionTransactor: SubmissionTransactor{contract: contract}, SubmissionFilterer: SubmissionFilterer{contract: contract}}, nil
}

// NewSubmissionCaller creates a new read-only instance of Submission, bound to a specific deployed contract.
func NewSubmissionCaller(address common.Address, caller bind.ContractCaller) (*SubmissionCaller, error) {
	contract, err := bindSubmission(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubmissionCaller{contract: contract}, nil
}

// NewSubmissionTransactor creates a new write-only instance of Submission, bound to a specific deployed contract.
func NewSubmissionTransactor(address common.Address, transactor bind.ContractTransactor) (*SubmissionTransactor, error) {
	contract, err := bindSubmission(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubmissionTransactor{contract: contract}, nil
}

// NewSubmissionFilterer creates a new log filterer instance of Submission, bound to a specific deployed contract.
func NewSubmissionFilterer(address common.Address, filterer bind.ContractFilterer) (*SubmissionFilterer, error) {
	contract, err := bindSubmission(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubmissionFilterer{contract: contract}, nil
}

// bindSubmission binds a generic wrapper to an already deployed contract.
func bindSubmission(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SubmissionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Submission *SubmissionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Submission.Contract.SubmissionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Submission *SubmissionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submission.Contract.SubmissionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Submission *SubmissionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Submission.Contract.SubmissionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Submission *SubmissionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Submission.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Submission *SubmissionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submission.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Submission *SubmissionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Submission.Contract.contract.Transact(opts, method, params...)
}

// FlareSystemManager is a free data retrieval call binding the contract method 0xbb25d5df.
//
// Solidity: function flareSystemManager() view returns(address)
func (_Submission *SubmissionCaller) FlareSystemManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submission.contract.Call(opts, &out, "flareSystemManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemManager is a free data retrieval call binding the contract method 0xbb25d5df.
//
// Solidity: function flareSystemManager() view returns(address)
func (_Submission *SubmissionSession) FlareSystemManager() (common.Address, error) {
	return _Submission.Contract.FlareSystemManager(&_Submission.CallOpts)
}

// FlareSystemManager is a free data retrieval call binding the contract method 0xbb25d5df.
//
// Solidity: function flareSystemManager() view returns(address)
func (_Submission *SubmissionCallerSession) FlareSystemManager() (common.Address, error) {
	return _Submission.Contract.FlareSystemManager(&_Submission.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Submission *SubmissionCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submission.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Submission *SubmissionSession) GetAddressUpdater() (common.Address, error) {
	return _Submission.Contract.GetAddressUpdater(&_Submission.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Submission *SubmissionCallerSession) GetAddressUpdater() (common.Address, error) {
	return _Submission.Contract.GetAddressUpdater(&_Submission.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Submission *SubmissionCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submission.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Submission *SubmissionSession) Governance() (common.Address, error) {
	return _Submission.Contract.Governance(&_Submission.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Submission *SubmissionCallerSession) Governance() (common.Address, error) {
	return _Submission.Contract.Governance(&_Submission.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Submission *SubmissionCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submission.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Submission *SubmissionSession) GovernanceSettings() (common.Address, error) {
	return _Submission.Contract.GovernanceSettings(&_Submission.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Submission *SubmissionCallerSession) GovernanceSettings() (common.Address, error) {
	return _Submission.Contract.GovernanceSettings(&_Submission.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_Submission *SubmissionCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Submission.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_Submission *SubmissionSession) IsExecutor(_address common.Address) (bool, error) {
	return _Submission.Contract.IsExecutor(&_Submission.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_Submission *SubmissionCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _Submission.Contract.IsExecutor(&_Submission.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Submission *SubmissionCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Submission.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Submission *SubmissionSession) ProductionMode() (bool, error) {
	return _Submission.Contract.ProductionMode(&_Submission.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Submission *SubmissionCallerSession) ProductionMode() (bool, error) {
	return _Submission.Contract.ProductionMode(&_Submission.CallOpts)
}

// Submit3MethodEnabled is a free data retrieval call binding the contract method 0x32de7a9f.
//
// Solidity: function submit3MethodEnabled() view returns(bool)
func (_Submission *SubmissionCaller) Submit3MethodEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Submission.contract.Call(opts, &out, "submit3MethodEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Submit3MethodEnabled is a free data retrieval call binding the contract method 0x32de7a9f.
//
// Solidity: function submit3MethodEnabled() view returns(bool)
func (_Submission *SubmissionSession) Submit3MethodEnabled() (bool, error) {
	return _Submission.Contract.Submit3MethodEnabled(&_Submission.CallOpts)
}

// Submit3MethodEnabled is a free data retrieval call binding the contract method 0x32de7a9f.
//
// Solidity: function submit3MethodEnabled() view returns(bool)
func (_Submission *SubmissionCallerSession) Submit3MethodEnabled() (bool, error) {
	return _Submission.Contract.Submit3MethodEnabled(&_Submission.CallOpts)
}

// SubmitAndPassContract is a free data retrieval call binding the contract method 0x93953af1.
//
// Solidity: function submitAndPassContract() view returns(address)
func (_Submission *SubmissionCaller) SubmitAndPassContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submission.contract.Call(opts, &out, "submitAndPassContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SubmitAndPassContract is a free data retrieval call binding the contract method 0x93953af1.
//
// Solidity: function submitAndPassContract() view returns(address)
func (_Submission *SubmissionSession) SubmitAndPassContract() (common.Address, error) {
	return _Submission.Contract.SubmitAndPassContract(&_Submission.CallOpts)
}

// SubmitAndPassContract is a free data retrieval call binding the contract method 0x93953af1.
//
// Solidity: function submitAndPassContract() view returns(address)
func (_Submission *SubmissionCallerSession) SubmitAndPassContract() (common.Address, error) {
	return _Submission.Contract.SubmitAndPassContract(&_Submission.CallOpts)
}

// SubmitAndPassSelector is a free data retrieval call binding the contract method 0xafd7f821.
//
// Solidity: function submitAndPassSelector() view returns(bytes4)
func (_Submission *SubmissionCaller) SubmitAndPassSelector(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Submission.contract.Call(opts, &out, "submitAndPassSelector")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// SubmitAndPassSelector is a free data retrieval call binding the contract method 0xafd7f821.
//
// Solidity: function submitAndPassSelector() view returns(bytes4)
func (_Submission *SubmissionSession) SubmitAndPassSelector() ([4]byte, error) {
	return _Submission.Contract.SubmitAndPassSelector(&_Submission.CallOpts)
}

// SubmitAndPassSelector is a free data retrieval call binding the contract method 0xafd7f821.
//
// Solidity: function submitAndPassSelector() view returns(bytes4)
func (_Submission *SubmissionCallerSession) SubmitAndPassSelector() ([4]byte, error) {
	return _Submission.Contract.SubmitAndPassSelector(&_Submission.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Submission *SubmissionCaller) TimelockedCalls(opts *bind.CallOpts, arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _Submission.contract.Call(opts, &out, "timelockedCalls", arg0)

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
func (_Submission *SubmissionSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _Submission.Contract.TimelockedCalls(&_Submission.CallOpts, arg0)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Submission *SubmissionCallerSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _Submission.Contract.TimelockedCalls(&_Submission.CallOpts, arg0)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Submission *SubmissionTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Submission *SubmissionSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Submission.Contract.CancelGovernanceCall(&_Submission.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Submission *SubmissionTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Submission.Contract.CancelGovernanceCall(&_Submission.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Submission *SubmissionTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Submission *SubmissionSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Submission.Contract.ExecuteGovernanceCall(&_Submission.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Submission *SubmissionTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Submission.Contract.ExecuteGovernanceCall(&_Submission.TransactOpts, _selector)
}

// InitNewVotingRound is a paid mutator transaction binding the contract method 0xf8ae8a2f.
//
// Solidity: function initNewVotingRound(address[] _submit1Addresses, address[] _submit2Addresses, address[] _submit3Addresses, address[] _submitSignaturesAddresses) returns()
func (_Submission *SubmissionTransactor) InitNewVotingRound(opts *bind.TransactOpts, _submit1Addresses []common.Address, _submit2Addresses []common.Address, _submit3Addresses []common.Address, _submitSignaturesAddresses []common.Address) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "initNewVotingRound", _submit1Addresses, _submit2Addresses, _submit3Addresses, _submitSignaturesAddresses)
}

// InitNewVotingRound is a paid mutator transaction binding the contract method 0xf8ae8a2f.
//
// Solidity: function initNewVotingRound(address[] _submit1Addresses, address[] _submit2Addresses, address[] _submit3Addresses, address[] _submitSignaturesAddresses) returns()
func (_Submission *SubmissionSession) InitNewVotingRound(_submit1Addresses []common.Address, _submit2Addresses []common.Address, _submit3Addresses []common.Address, _submitSignaturesAddresses []common.Address) (*types.Transaction, error) {
	return _Submission.Contract.InitNewVotingRound(&_Submission.TransactOpts, _submit1Addresses, _submit2Addresses, _submit3Addresses, _submitSignaturesAddresses)
}

// InitNewVotingRound is a paid mutator transaction binding the contract method 0xf8ae8a2f.
//
// Solidity: function initNewVotingRound(address[] _submit1Addresses, address[] _submit2Addresses, address[] _submit3Addresses, address[] _submitSignaturesAddresses) returns()
func (_Submission *SubmissionTransactorSession) InitNewVotingRound(_submit1Addresses []common.Address, _submit2Addresses []common.Address, _submit3Addresses []common.Address, _submitSignaturesAddresses []common.Address) (*types.Transaction, error) {
	return _Submission.Contract.InitNewVotingRound(&_Submission.TransactOpts, _submit1Addresses, _submit2Addresses, _submit3Addresses, _submitSignaturesAddresses)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_Submission *SubmissionTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_Submission *SubmissionSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Submission.Contract.Initialise(&_Submission.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_Submission *SubmissionTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Submission.Contract.Initialise(&_Submission.TransactOpts, _governanceSettings, _initialGovernance)
}

// SetSubmit3MethodEnabled is a paid mutator transaction binding the contract method 0x941877d0.
//
// Solidity: function setSubmit3MethodEnabled(bool _enabled) returns()
func (_Submission *SubmissionTransactor) SetSubmit3MethodEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "setSubmit3MethodEnabled", _enabled)
}

// SetSubmit3MethodEnabled is a paid mutator transaction binding the contract method 0x941877d0.
//
// Solidity: function setSubmit3MethodEnabled(bool _enabled) returns()
func (_Submission *SubmissionSession) SetSubmit3MethodEnabled(_enabled bool) (*types.Transaction, error) {
	return _Submission.Contract.SetSubmit3MethodEnabled(&_Submission.TransactOpts, _enabled)
}

// SetSubmit3MethodEnabled is a paid mutator transaction binding the contract method 0x941877d0.
//
// Solidity: function setSubmit3MethodEnabled(bool _enabled) returns()
func (_Submission *SubmissionTransactorSession) SetSubmit3MethodEnabled(_enabled bool) (*types.Transaction, error) {
	return _Submission.Contract.SetSubmit3MethodEnabled(&_Submission.TransactOpts, _enabled)
}

// SetSubmitAndPassData is a paid mutator transaction binding the contract method 0x9ee7fe4d.
//
// Solidity: function setSubmitAndPassData(address _submitAndPassContract, bytes4 _submitAndPassSelector) returns()
func (_Submission *SubmissionTransactor) SetSubmitAndPassData(opts *bind.TransactOpts, _submitAndPassContract common.Address, _submitAndPassSelector [4]byte) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "setSubmitAndPassData", _submitAndPassContract, _submitAndPassSelector)
}

// SetSubmitAndPassData is a paid mutator transaction binding the contract method 0x9ee7fe4d.
//
// Solidity: function setSubmitAndPassData(address _submitAndPassContract, bytes4 _submitAndPassSelector) returns()
func (_Submission *SubmissionSession) SetSubmitAndPassData(_submitAndPassContract common.Address, _submitAndPassSelector [4]byte) (*types.Transaction, error) {
	return _Submission.Contract.SetSubmitAndPassData(&_Submission.TransactOpts, _submitAndPassContract, _submitAndPassSelector)
}

// SetSubmitAndPassData is a paid mutator transaction binding the contract method 0x9ee7fe4d.
//
// Solidity: function setSubmitAndPassData(address _submitAndPassContract, bytes4 _submitAndPassSelector) returns()
func (_Submission *SubmissionTransactorSession) SetSubmitAndPassData(_submitAndPassContract common.Address, _submitAndPassSelector [4]byte) (*types.Transaction, error) {
	return _Submission.Contract.SetSubmitAndPassData(&_Submission.TransactOpts, _submitAndPassContract, _submitAndPassSelector)
}

// Submit1 is a paid mutator transaction binding the contract method 0x6c532fae.
//
// Solidity: function submit1() returns(bool)
func (_Submission *SubmissionTransactor) Submit1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "submit1")
}

// Submit1 is a paid mutator transaction binding the contract method 0x6c532fae.
//
// Solidity: function submit1() returns(bool)
func (_Submission *SubmissionSession) Submit1() (*types.Transaction, error) {
	return _Submission.Contract.Submit1(&_Submission.TransactOpts)
}

// Submit1 is a paid mutator transaction binding the contract method 0x6c532fae.
//
// Solidity: function submit1() returns(bool)
func (_Submission *SubmissionTransactorSession) Submit1() (*types.Transaction, error) {
	return _Submission.Contract.Submit1(&_Submission.TransactOpts)
}

// Submit2 is a paid mutator transaction binding the contract method 0x9d00c9fd.
//
// Solidity: function submit2() returns(bool)
func (_Submission *SubmissionTransactor) Submit2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "submit2")
}

// Submit2 is a paid mutator transaction binding the contract method 0x9d00c9fd.
//
// Solidity: function submit2() returns(bool)
func (_Submission *SubmissionSession) Submit2() (*types.Transaction, error) {
	return _Submission.Contract.Submit2(&_Submission.TransactOpts)
}

// Submit2 is a paid mutator transaction binding the contract method 0x9d00c9fd.
//
// Solidity: function submit2() returns(bool)
func (_Submission *SubmissionTransactorSession) Submit2() (*types.Transaction, error) {
	return _Submission.Contract.Submit2(&_Submission.TransactOpts)
}

// Submit3 is a paid mutator transaction binding the contract method 0xe1b157e7.
//
// Solidity: function submit3() returns(bool)
func (_Submission *SubmissionTransactor) Submit3(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "submit3")
}

// Submit3 is a paid mutator transaction binding the contract method 0xe1b157e7.
//
// Solidity: function submit3() returns(bool)
func (_Submission *SubmissionSession) Submit3() (*types.Transaction, error) {
	return _Submission.Contract.Submit3(&_Submission.TransactOpts)
}

// Submit3 is a paid mutator transaction binding the contract method 0xe1b157e7.
//
// Solidity: function submit3() returns(bool)
func (_Submission *SubmissionTransactorSession) Submit3() (*types.Transaction, error) {
	return _Submission.Contract.Submit3(&_Submission.TransactOpts)
}

// SubmitAndPass is a paid mutator transaction binding the contract method 0x833bf6c0.
//
// Solidity: function submitAndPass(bytes _data) returns(bool)
func (_Submission *SubmissionTransactor) SubmitAndPass(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "submitAndPass", _data)
}

// SubmitAndPass is a paid mutator transaction binding the contract method 0x833bf6c0.
//
// Solidity: function submitAndPass(bytes _data) returns(bool)
func (_Submission *SubmissionSession) SubmitAndPass(_data []byte) (*types.Transaction, error) {
	return _Submission.Contract.SubmitAndPass(&_Submission.TransactOpts, _data)
}

// SubmitAndPass is a paid mutator transaction binding the contract method 0x833bf6c0.
//
// Solidity: function submitAndPass(bytes _data) returns(bool)
func (_Submission *SubmissionTransactorSession) SubmitAndPass(_data []byte) (*types.Transaction, error) {
	return _Submission.Contract.SubmitAndPass(&_Submission.TransactOpts, _data)
}

// SubmitSignatures is a paid mutator transaction binding the contract method 0x57eed580.
//
// Solidity: function submitSignatures() returns(bool)
func (_Submission *SubmissionTransactor) SubmitSignatures(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "submitSignatures")
}

// SubmitSignatures is a paid mutator transaction binding the contract method 0x57eed580.
//
// Solidity: function submitSignatures() returns(bool)
func (_Submission *SubmissionSession) SubmitSignatures() (*types.Transaction, error) {
	return _Submission.Contract.SubmitSignatures(&_Submission.TransactOpts)
}

// SubmitSignatures is a paid mutator transaction binding the contract method 0x57eed580.
//
// Solidity: function submitSignatures() returns(bool)
func (_Submission *SubmissionTransactorSession) SubmitSignatures() (*types.Transaction, error) {
	return _Submission.Contract.SubmitSignatures(&_Submission.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Submission *SubmissionTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Submission *SubmissionSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _Submission.Contract.SwitchToProductionMode(&_Submission.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Submission *SubmissionTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _Submission.Contract.SwitchToProductionMode(&_Submission.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Submission *SubmissionTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Submission *SubmissionSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Submission.Contract.UpdateContractAddresses(&_Submission.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Submission *SubmissionTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Submission.Contract.UpdateContractAddresses(&_Submission.TransactOpts, _contractNameHashes, _contractAddresses)
}

// SubmissionGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the Submission contract.
type SubmissionGovernanceCallTimelockedIterator struct {
	Event *SubmissionGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *SubmissionGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmissionGovernanceCallTimelocked)
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
		it.Event = new(SubmissionGovernanceCallTimelocked)
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
func (it *SubmissionGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmissionGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmissionGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Submission contract.
type SubmissionGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Submission *SubmissionFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*SubmissionGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _Submission.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &SubmissionGovernanceCallTimelockedIterator{contract: _Submission.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Submission *SubmissionFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *SubmissionGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _Submission.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmissionGovernanceCallTimelocked)
				if err := _Submission.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_Submission *SubmissionFilterer) ParseGovernanceCallTimelocked(log types.Log) (*SubmissionGovernanceCallTimelocked, error) {
	event := new(SubmissionGovernanceCallTimelocked)
	if err := _Submission.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmissionGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the Submission contract.
type SubmissionGovernanceInitialisedIterator struct {
	Event *SubmissionGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *SubmissionGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmissionGovernanceInitialised)
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
		it.Event = new(SubmissionGovernanceInitialised)
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
func (it *SubmissionGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmissionGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmissionGovernanceInitialised represents a GovernanceInitialised event raised by the Submission contract.
type SubmissionGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_Submission *SubmissionFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*SubmissionGovernanceInitialisedIterator, error) {

	logs, sub, err := _Submission.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &SubmissionGovernanceInitialisedIterator{contract: _Submission.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_Submission *SubmissionFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *SubmissionGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _Submission.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmissionGovernanceInitialised)
				if err := _Submission.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_Submission *SubmissionFilterer) ParseGovernanceInitialised(log types.Log) (*SubmissionGovernanceInitialised, error) {
	event := new(SubmissionGovernanceInitialised)
	if err := _Submission.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmissionGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the Submission contract.
type SubmissionGovernedProductionModeEnteredIterator struct {
	Event *SubmissionGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *SubmissionGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmissionGovernedProductionModeEntered)
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
		it.Event = new(SubmissionGovernedProductionModeEntered)
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
func (it *SubmissionGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmissionGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmissionGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the Submission contract.
type SubmissionGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_Submission *SubmissionFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*SubmissionGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _Submission.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &SubmissionGovernedProductionModeEnteredIterator{contract: _Submission.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_Submission *SubmissionFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *SubmissionGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _Submission.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmissionGovernedProductionModeEntered)
				if err := _Submission.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_Submission *SubmissionFilterer) ParseGovernedProductionModeEntered(log types.Log) (*SubmissionGovernedProductionModeEntered, error) {
	event := new(SubmissionGovernedProductionModeEntered)
	if err := _Submission.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmissionNewVotingRoundInitiatedIterator is returned from FilterNewVotingRoundInitiated and is used to iterate over the raw logs and unpacked data for NewVotingRoundInitiated events raised by the Submission contract.
type SubmissionNewVotingRoundInitiatedIterator struct {
	Event *SubmissionNewVotingRoundInitiated // Event containing the contract specifics and raw log

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
func (it *SubmissionNewVotingRoundInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmissionNewVotingRoundInitiated)
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
		it.Event = new(SubmissionNewVotingRoundInitiated)
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
func (it *SubmissionNewVotingRoundInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmissionNewVotingRoundInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmissionNewVotingRoundInitiated represents a NewVotingRoundInitiated event raised by the Submission contract.
type SubmissionNewVotingRoundInitiated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewVotingRoundInitiated is a free log retrieval operation binding the contract event 0xb74d3a815b816fdb5f14fb14f14bf86e1a87dcbc3f23150f1c32f89cd4622f3d.
//
// Solidity: event NewVotingRoundInitiated()
func (_Submission *SubmissionFilterer) FilterNewVotingRoundInitiated(opts *bind.FilterOpts) (*SubmissionNewVotingRoundInitiatedIterator, error) {

	logs, sub, err := _Submission.contract.FilterLogs(opts, "NewVotingRoundInitiated")
	if err != nil {
		return nil, err
	}
	return &SubmissionNewVotingRoundInitiatedIterator{contract: _Submission.contract, event: "NewVotingRoundInitiated", logs: logs, sub: sub}, nil
}

// WatchNewVotingRoundInitiated is a free log subscription operation binding the contract event 0xb74d3a815b816fdb5f14fb14f14bf86e1a87dcbc3f23150f1c32f89cd4622f3d.
//
// Solidity: event NewVotingRoundInitiated()
func (_Submission *SubmissionFilterer) WatchNewVotingRoundInitiated(opts *bind.WatchOpts, sink chan<- *SubmissionNewVotingRoundInitiated) (event.Subscription, error) {

	logs, sub, err := _Submission.contract.WatchLogs(opts, "NewVotingRoundInitiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmissionNewVotingRoundInitiated)
				if err := _Submission.contract.UnpackLog(event, "NewVotingRoundInitiated", log); err != nil {
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

// ParseNewVotingRoundInitiated is a log parse operation binding the contract event 0xb74d3a815b816fdb5f14fb14f14bf86e1a87dcbc3f23150f1c32f89cd4622f3d.
//
// Solidity: event NewVotingRoundInitiated()
func (_Submission *SubmissionFilterer) ParseNewVotingRoundInitiated(log types.Log) (*SubmissionNewVotingRoundInitiated, error) {
	event := new(SubmissionNewVotingRoundInitiated)
	if err := _Submission.contract.UnpackLog(event, "NewVotingRoundInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmissionTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the Submission contract.
type SubmissionTimelockedGovernanceCallCanceledIterator struct {
	Event *SubmissionTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *SubmissionTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmissionTimelockedGovernanceCallCanceled)
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
		it.Event = new(SubmissionTimelockedGovernanceCallCanceled)
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
func (it *SubmissionTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmissionTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmissionTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the Submission contract.
type SubmissionTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_Submission *SubmissionFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*SubmissionTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _Submission.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &SubmissionTimelockedGovernanceCallCanceledIterator{contract: _Submission.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_Submission *SubmissionFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *SubmissionTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _Submission.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmissionTimelockedGovernanceCallCanceled)
				if err := _Submission.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_Submission *SubmissionFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*SubmissionTimelockedGovernanceCallCanceled, error) {
	event := new(SubmissionTimelockedGovernanceCallCanceled)
	if err := _Submission.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmissionTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the Submission contract.
type SubmissionTimelockedGovernanceCallExecutedIterator struct {
	Event *SubmissionTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *SubmissionTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmissionTimelockedGovernanceCallExecuted)
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
		it.Event = new(SubmissionTimelockedGovernanceCallExecuted)
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
func (it *SubmissionTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmissionTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmissionTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the Submission contract.
type SubmissionTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_Submission *SubmissionFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*SubmissionTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _Submission.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &SubmissionTimelockedGovernanceCallExecutedIterator{contract: _Submission.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_Submission *SubmissionFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *SubmissionTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _Submission.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmissionTimelockedGovernanceCallExecuted)
				if err := _Submission.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_Submission *SubmissionFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*SubmissionTimelockedGovernanceCallExecuted, error) {
	event := new(SubmissionTimelockedGovernanceCallExecuted)
	if err := _Submission.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
