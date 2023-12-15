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
	ABI: "[{\"inputs\":[],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitFastPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Commit is a paid mutator transaction binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() returns()
func (_Submission *SubmissionTransactor) Commit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "commit")
}

// Commit is a paid mutator transaction binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() returns()
func (_Submission *SubmissionSession) Commit() (*types.Transaction, error) {
	return _Submission.Contract.Commit(&_Submission.TransactOpts)
}

// Commit is a paid mutator transaction binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() returns()
func (_Submission *SubmissionTransactorSession) Commit() (*types.Transaction, error) {
	return _Submission.Contract.Commit(&_Submission.TransactOpts)
}

// Finalise is a paid mutator transaction binding the contract method 0xa4399263.
//
// Solidity: function finalise() returns()
func (_Submission *SubmissionTransactor) Finalise(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "finalise")
}

// Finalise is a paid mutator transaction binding the contract method 0xa4399263.
//
// Solidity: function finalise() returns()
func (_Submission *SubmissionSession) Finalise() (*types.Transaction, error) {
	return _Submission.Contract.Finalise(&_Submission.TransactOpts)
}

// Finalise is a paid mutator transaction binding the contract method 0xa4399263.
//
// Solidity: function finalise() returns()
func (_Submission *SubmissionTransactorSession) Finalise() (*types.Transaction, error) {
	return _Submission.Contract.Finalise(&_Submission.TransactOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() returns()
func (_Submission *SubmissionTransactor) Reveal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "reveal")
}

// Reveal is a paid mutator transaction binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() returns()
func (_Submission *SubmissionSession) Reveal() (*types.Transaction, error) {
	return _Submission.Contract.Reveal(&_Submission.TransactOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() returns()
func (_Submission *SubmissionTransactorSession) Reveal() (*types.Transaction, error) {
	return _Submission.Contract.Reveal(&_Submission.TransactOpts)
}

// Sign is a paid mutator transaction binding the contract method 0x2ca15122.
//
// Solidity: function sign() returns()
func (_Submission *SubmissionTransactor) Sign(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "sign")
}

// Sign is a paid mutator transaction binding the contract method 0x2ca15122.
//
// Solidity: function sign() returns()
func (_Submission *SubmissionSession) Sign() (*types.Transaction, error) {
	return _Submission.Contract.Sign(&_Submission.TransactOpts)
}

// Sign is a paid mutator transaction binding the contract method 0x2ca15122.
//
// Solidity: function sign() returns()
func (_Submission *SubmissionTransactorSession) Sign() (*types.Transaction, error) {
	return _Submission.Contract.Sign(&_Submission.TransactOpts)
}

// SubmitFastPrices is a paid mutator transaction binding the contract method 0x5d2417cc.
//
// Solidity: function submitFastPrices() returns()
func (_Submission *SubmissionTransactor) SubmitFastPrices(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submission.contract.Transact(opts, "submitFastPrices")
}

// SubmitFastPrices is a paid mutator transaction binding the contract method 0x5d2417cc.
//
// Solidity: function submitFastPrices() returns()
func (_Submission *SubmissionSession) SubmitFastPrices() (*types.Transaction, error) {
	return _Submission.Contract.SubmitFastPrices(&_Submission.TransactOpts)
}

// SubmitFastPrices is a paid mutator transaction binding the contract method 0x5d2417cc.
//
// Solidity: function submitFastPrices() returns()
func (_Submission *SubmissionTransactorSession) SubmitFastPrices() (*types.Transaction, error) {
	return _Submission.Contract.SubmitFastPrices(&_Submission.TransactOpts)
}
