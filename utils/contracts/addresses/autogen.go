// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package addresses

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

// BinderMetaData contains all meta data concerning the Binder contract.
var BinderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes20\",\"name\":\"pAddress\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cAddress\",\"type\":\"address\"}],\"name\":\"AddressesRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cAddressToPAddress\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"name\":\"pAddressToCAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes20\",\"name\":\"_pAddress\",\"type\":\"bytes20\"},{\"internalType\":\"address\",\"name\":\"_cAddress\",\"type\":\"address\"}],\"name\":\"registerAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"registerPublicKey\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"_pAddress\",\"type\":\"bytes20\"},{\"internalType\":\"address\",\"name\":\"_cAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BinderABI is the input ABI used to generate the binding from.
// Deprecated: Use BinderMetaData.ABI instead.
var BinderABI = BinderMetaData.ABI

// Binder is an auto generated Go binding around an Ethereum contract.
type Binder struct {
	BinderCaller     // Read-only binding to the contract
	BinderTransactor // Write-only binding to the contract
	BinderFilterer   // Log filterer for contract events
}

// BinderCaller is an auto generated read-only Go binding around an Ethereum contract.
type BinderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BinderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BinderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BinderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BinderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BinderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BinderSession struct {
	Contract     *Binder           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BinderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BinderCallerSession struct {
	Contract *BinderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BinderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BinderTransactorSession struct {
	Contract     *BinderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BinderRaw is an auto generated low-level Go binding around an Ethereum contract.
type BinderRaw struct {
	Contract *Binder // Generic contract binding to access the raw methods on
}

// BinderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BinderCallerRaw struct {
	Contract *BinderCaller // Generic read-only contract binding to access the raw methods on
}

// BinderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BinderTransactorRaw struct {
	Contract *BinderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBinder creates a new instance of Binder, bound to a specific deployed contract.
func NewBinder(address common.Address, backend bind.ContractBackend) (*Binder, error) {
	contract, err := bindBinder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Binder{BinderCaller: BinderCaller{contract: contract}, BinderTransactor: BinderTransactor{contract: contract}, BinderFilterer: BinderFilterer{contract: contract}}, nil
}

// NewBinderCaller creates a new read-only instance of Binder, bound to a specific deployed contract.
func NewBinderCaller(address common.Address, caller bind.ContractCaller) (*BinderCaller, error) {
	contract, err := bindBinder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BinderCaller{contract: contract}, nil
}

// NewBinderTransactor creates a new write-only instance of Binder, bound to a specific deployed contract.
func NewBinderTransactor(address common.Address, transactor bind.ContractTransactor) (*BinderTransactor, error) {
	contract, err := bindBinder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BinderTransactor{contract: contract}, nil
}

// NewBinderFilterer creates a new log filterer instance of Binder, bound to a specific deployed contract.
func NewBinderFilterer(address common.Address, filterer bind.ContractFilterer) (*BinderFilterer, error) {
	contract, err := bindBinder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BinderFilterer{contract: contract}, nil
}

// bindBinder binds a generic wrapper to an already deployed contract.
func bindBinder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BinderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Binder *BinderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Binder.Contract.BinderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Binder *BinderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Binder.Contract.BinderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Binder *BinderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Binder.Contract.BinderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Binder *BinderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Binder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Binder *BinderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Binder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Binder *BinderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Binder.Contract.contract.Transact(opts, method, params...)
}

// CAddressToPAddress is a free data retrieval call binding the contract method 0xe82199ad.
//
// Solidity: function cAddressToPAddress(address ) view returns(bytes20)
func (_Binder *BinderCaller) CAddressToPAddress(opts *bind.CallOpts, arg0 common.Address) ([20]byte, error) {
	var out []interface{}
	err := _Binder.contract.Call(opts, &out, "cAddressToPAddress", arg0)

	if err != nil {
		return *new([20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([20]byte)).(*[20]byte)

	return out0, err

}

// CAddressToPAddress is a free data retrieval call binding the contract method 0xe82199ad.
//
// Solidity: function cAddressToPAddress(address ) view returns(bytes20)
func (_Binder *BinderSession) CAddressToPAddress(arg0 common.Address) ([20]byte, error) {
	return _Binder.Contract.CAddressToPAddress(&_Binder.CallOpts, arg0)
}

// CAddressToPAddress is a free data retrieval call binding the contract method 0xe82199ad.
//
// Solidity: function cAddressToPAddress(address ) view returns(bytes20)
func (_Binder *BinderCallerSession) CAddressToPAddress(arg0 common.Address) ([20]byte, error) {
	return _Binder.Contract.CAddressToPAddress(&_Binder.CallOpts, arg0)
}

// PAddressToCAddress is a free data retrieval call binding the contract method 0x373e6999.
//
// Solidity: function pAddressToCAddress(bytes20 ) view returns(address)
func (_Binder *BinderCaller) PAddressToCAddress(opts *bind.CallOpts, arg0 [20]byte) (common.Address, error) {
	var out []interface{}
	err := _Binder.contract.Call(opts, &out, "pAddressToCAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PAddressToCAddress is a free data retrieval call binding the contract method 0x373e6999.
//
// Solidity: function pAddressToCAddress(bytes20 ) view returns(address)
func (_Binder *BinderSession) PAddressToCAddress(arg0 [20]byte) (common.Address, error) {
	return _Binder.Contract.PAddressToCAddress(&_Binder.CallOpts, arg0)
}

// PAddressToCAddress is a free data retrieval call binding the contract method 0x373e6999.
//
// Solidity: function pAddressToCAddress(bytes20 ) view returns(address)
func (_Binder *BinderCallerSession) PAddressToCAddress(arg0 [20]byte) (common.Address, error) {
	return _Binder.Contract.PAddressToCAddress(&_Binder.CallOpts, arg0)
}

// RegisterAddresses is a paid mutator transaction binding the contract method 0x5b75d79c.
//
// Solidity: function registerAddresses(bytes _publicKey, bytes20 _pAddress, address _cAddress) returns()
func (_Binder *BinderTransactor) RegisterAddresses(opts *bind.TransactOpts, _publicKey []byte, _pAddress [20]byte, _cAddress common.Address) (*types.Transaction, error) {
	return _Binder.contract.Transact(opts, "registerAddresses", _publicKey, _pAddress, _cAddress)
}

// RegisterAddresses is a paid mutator transaction binding the contract method 0x5b75d79c.
//
// Solidity: function registerAddresses(bytes _publicKey, bytes20 _pAddress, address _cAddress) returns()
func (_Binder *BinderSession) RegisterAddresses(_publicKey []byte, _pAddress [20]byte, _cAddress common.Address) (*types.Transaction, error) {
	return _Binder.Contract.RegisterAddresses(&_Binder.TransactOpts, _publicKey, _pAddress, _cAddress)
}

// RegisterAddresses is a paid mutator transaction binding the contract method 0x5b75d79c.
//
// Solidity: function registerAddresses(bytes _publicKey, bytes20 _pAddress, address _cAddress) returns()
func (_Binder *BinderTransactorSession) RegisterAddresses(_publicKey []byte, _pAddress [20]byte, _cAddress common.Address) (*types.Transaction, error) {
	return _Binder.Contract.RegisterAddresses(&_Binder.TransactOpts, _publicKey, _pAddress, _cAddress)
}

// RegisterPublicKey is a paid mutator transaction binding the contract method 0x85623594.
//
// Solidity: function registerPublicKey(bytes _publicKey) returns(bytes20 _pAddress, address _cAddress)
func (_Binder *BinderTransactor) RegisterPublicKey(opts *bind.TransactOpts, _publicKey []byte) (*types.Transaction, error) {
	return _Binder.contract.Transact(opts, "registerPublicKey", _publicKey)
}

// RegisterPublicKey is a paid mutator transaction binding the contract method 0x85623594.
//
// Solidity: function registerPublicKey(bytes _publicKey) returns(bytes20 _pAddress, address _cAddress)
func (_Binder *BinderSession) RegisterPublicKey(_publicKey []byte) (*types.Transaction, error) {
	return _Binder.Contract.RegisterPublicKey(&_Binder.TransactOpts, _publicKey)
}

// RegisterPublicKey is a paid mutator transaction binding the contract method 0x85623594.
//
// Solidity: function registerPublicKey(bytes _publicKey) returns(bytes20 _pAddress, address _cAddress)
func (_Binder *BinderTransactorSession) RegisterPublicKey(_publicKey []byte) (*types.Transaction, error) {
	return _Binder.Contract.RegisterPublicKey(&_Binder.TransactOpts, _publicKey)
}

// BinderAddressesRegisteredIterator is returned from FilterAddressesRegistered and is used to iterate over the raw logs and unpacked data for AddressesRegistered events raised by the Binder contract.
type BinderAddressesRegisteredIterator struct {
	Event *BinderAddressesRegistered // Event containing the contract specifics and raw log

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
func (it *BinderAddressesRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BinderAddressesRegistered)
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
		it.Event = new(BinderAddressesRegistered)
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
func (it *BinderAddressesRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BinderAddressesRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BinderAddressesRegistered represents a AddressesRegistered event raised by the Binder contract.
type BinderAddressesRegistered struct {
	PublicKey []byte
	PAddress  [20]byte
	CAddress  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddressesRegistered is a free log retrieval operation binding the contract event 0x5438d74d11fdde1d4f54295fb3e6c454489914c68bb311c255399c3328ae3082.
//
// Solidity: event AddressesRegistered(bytes publicKey, bytes20 pAddress, address cAddress)
func (_Binder *BinderFilterer) FilterAddressesRegistered(opts *bind.FilterOpts) (*BinderAddressesRegisteredIterator, error) {

	logs, sub, err := _Binder.contract.FilterLogs(opts, "AddressesRegistered")
	if err != nil {
		return nil, err
	}
	return &BinderAddressesRegisteredIterator{contract: _Binder.contract, event: "AddressesRegistered", logs: logs, sub: sub}, nil
}

// WatchAddressesRegistered is a free log subscription operation binding the contract event 0x5438d74d11fdde1d4f54295fb3e6c454489914c68bb311c255399c3328ae3082.
//
// Solidity: event AddressesRegistered(bytes publicKey, bytes20 pAddress, address cAddress)
func (_Binder *BinderFilterer) WatchAddressesRegistered(opts *bind.WatchOpts, sink chan<- *BinderAddressesRegistered) (event.Subscription, error) {

	logs, sub, err := _Binder.contract.WatchLogs(opts, "AddressesRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BinderAddressesRegistered)
				if err := _Binder.contract.UnpackLog(event, "AddressesRegistered", log); err != nil {
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

// ParseAddressesRegistered is a log parse operation binding the contract event 0x5438d74d11fdde1d4f54295fb3e6c454489914c68bb311c255399c3328ae3082.
//
// Solidity: event AddressesRegistered(bytes publicKey, bytes20 pAddress, address cAddress)
func (_Binder *BinderFilterer) ParseAddressesRegistered(log types.Log) (*BinderAddressesRegistered, error) {
	event := new(BinderAddressesRegistered)
	if err := _Binder.contract.UnpackLog(event, "AddressesRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
