// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// IVoterRegistrySignature is an auto generated low-level Go binding around an user-defined struct.
type IVoterRegistrySignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialRewardEpochId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_initialVoters\",\"type\":\"address[]\"},{\"internalType\":\"uint16[]\",\"name\":\"_initialNormalisedWeights\",\"type\":\"uint16[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"untilRewardEpochId\",\"type\":\"uint256\"}],\"name\":\"VoterChilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"submitAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"submitSignaturesAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"registrationWeight\",\"type\":\"uint256\"}],\"name\":\"VoterRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"VoterRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_noOfRewardEpochs\",\"type\":\"uint256\"}],\"name\":\"chillVoter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_untilRewardEpochId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"chilledUntilRewardEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"createSigningPolicySnapshot\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_signingPolicyAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint16[]\",\"name\":\"_normalisedWeights\",\"type\":\"uint16[]\"},{\"internalType\":\"uint16\",\"name\":\"_normalisedWeightsSum\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entityManager\",\"outputs\":[{\"internalType\":\"contractIIEntityManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemCalculator\",\"outputs\":[{\"internalType\":\"contractIIFlareSystemCalculator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getNumberOfRegisteredVoters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signingPolicyAddress\",\"type\":\"address\"}],\"name\":\"getPublicKeyAndNormalisedWeight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_publicKeyPart1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_publicKeyPart2\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"_normalisedWeight\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_normalisedWeightsSumOfVotersWithPublicKeys\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getRegisteredDelegationAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_delegationAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getRegisteredNodeIds\",\"outputs\":[{\"internalType\":\"bytes20[][]\",\"name\":\"_nodeIds\",\"type\":\"bytes20[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getRegisteredPublicKeys\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_parts1\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_parts2\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getRegisteredSigningPolicyAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_signingPolicyAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getRegisteredSubmitAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getRegisteredSubmitSignaturesAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_signingPolicyAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getRegisteredVoters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signingPolicyAddress\",\"type\":\"address\"}],\"name\":\"getVoterWithNormalisedWeight\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_normalisedWeight\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getWeightsSums\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"_weightsSum\",\"type\":\"uint128\"},{\"internalType\":\"uint16\",\"name\":\"_normalisedWeightsSum\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_normalisedWeightsSumOfVotersWithPublicKeys\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"isVoterRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxVoters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newSigningPolicyInitializationStartBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIVoterRegistry.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"registerVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxVoters\",\"type\":\"uint256\"}],\"name\":\"setMaxVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"setNewSigningPolicyInitializationStartBlockNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_systemRegistrationContractAddress\",\"type\":\"address\"}],\"name\":\"setSystemRegistrationContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"systemRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRegistrationContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// ChilledUntilRewardEpochId is a free data retrieval call binding the contract method 0x1e3f047e.
//
// Solidity: function chilledUntilRewardEpochId(address ) view returns(uint256)
func (_Registry *RegistryCaller) ChilledUntilRewardEpochId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "chilledUntilRewardEpochId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChilledUntilRewardEpochId is a free data retrieval call binding the contract method 0x1e3f047e.
//
// Solidity: function chilledUntilRewardEpochId(address ) view returns(uint256)
func (_Registry *RegistrySession) ChilledUntilRewardEpochId(arg0 common.Address) (*big.Int, error) {
	return _Registry.Contract.ChilledUntilRewardEpochId(&_Registry.CallOpts, arg0)
}

// ChilledUntilRewardEpochId is a free data retrieval call binding the contract method 0x1e3f047e.
//
// Solidity: function chilledUntilRewardEpochId(address ) view returns(uint256)
func (_Registry *RegistryCallerSession) ChilledUntilRewardEpochId(arg0 common.Address) (*big.Int, error) {
	return _Registry.Contract.ChilledUntilRewardEpochId(&_Registry.CallOpts, arg0)
}

// EntityManager is a free data retrieval call binding the contract method 0x50b1d61b.
//
// Solidity: function entityManager() view returns(address)
func (_Registry *RegistryCaller) EntityManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "entityManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntityManager is a free data retrieval call binding the contract method 0x50b1d61b.
//
// Solidity: function entityManager() view returns(address)
func (_Registry *RegistrySession) EntityManager() (common.Address, error) {
	return _Registry.Contract.EntityManager(&_Registry.CallOpts)
}

// EntityManager is a free data retrieval call binding the contract method 0x50b1d61b.
//
// Solidity: function entityManager() view returns(address)
func (_Registry *RegistryCallerSession) EntityManager() (common.Address, error) {
	return _Registry.Contract.EntityManager(&_Registry.CallOpts)
}

// FlareSystemCalculator is a free data retrieval call binding the contract method 0x87ad4ecf.
//
// Solidity: function flareSystemCalculator() view returns(address)
func (_Registry *RegistryCaller) FlareSystemCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "flareSystemCalculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemCalculator is a free data retrieval call binding the contract method 0x87ad4ecf.
//
// Solidity: function flareSystemCalculator() view returns(address)
func (_Registry *RegistrySession) FlareSystemCalculator() (common.Address, error) {
	return _Registry.Contract.FlareSystemCalculator(&_Registry.CallOpts)
}

// FlareSystemCalculator is a free data retrieval call binding the contract method 0x87ad4ecf.
//
// Solidity: function flareSystemCalculator() view returns(address)
func (_Registry *RegistryCallerSession) FlareSystemCalculator() (common.Address, error) {
	return _Registry.Contract.FlareSystemCalculator(&_Registry.CallOpts)
}

// FlareSystemManager is a free data retrieval call binding the contract method 0xbb25d5df.
//
// Solidity: function flareSystemManager() view returns(address)
func (_Registry *RegistryCaller) FlareSystemManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "flareSystemManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemManager is a free data retrieval call binding the contract method 0xbb25d5df.
//
// Solidity: function flareSystemManager() view returns(address)
func (_Registry *RegistrySession) FlareSystemManager() (common.Address, error) {
	return _Registry.Contract.FlareSystemManager(&_Registry.CallOpts)
}

// FlareSystemManager is a free data retrieval call binding the contract method 0xbb25d5df.
//
// Solidity: function flareSystemManager() view returns(address)
func (_Registry *RegistryCallerSession) FlareSystemManager() (common.Address, error) {
	return _Registry.Contract.FlareSystemManager(&_Registry.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Registry *RegistryCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Registry *RegistrySession) GetAddressUpdater() (common.Address, error) {
	return _Registry.Contract.GetAddressUpdater(&_Registry.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Registry *RegistryCallerSession) GetAddressUpdater() (common.Address, error) {
	return _Registry.Contract.GetAddressUpdater(&_Registry.CallOpts)
}

// GetNumberOfRegisteredVoters is a free data retrieval call binding the contract method 0x369e9434.
//
// Solidity: function getNumberOfRegisteredVoters(uint256 _rewardEpochId) view returns(uint256)
func (_Registry *RegistryCaller) GetNumberOfRegisteredVoters(opts *bind.CallOpts, _rewardEpochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getNumberOfRegisteredVoters", _rewardEpochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfRegisteredVoters is a free data retrieval call binding the contract method 0x369e9434.
//
// Solidity: function getNumberOfRegisteredVoters(uint256 _rewardEpochId) view returns(uint256)
func (_Registry *RegistrySession) GetNumberOfRegisteredVoters(_rewardEpochId *big.Int) (*big.Int, error) {
	return _Registry.Contract.GetNumberOfRegisteredVoters(&_Registry.CallOpts, _rewardEpochId)
}

// GetNumberOfRegisteredVoters is a free data retrieval call binding the contract method 0x369e9434.
//
// Solidity: function getNumberOfRegisteredVoters(uint256 _rewardEpochId) view returns(uint256)
func (_Registry *RegistryCallerSession) GetNumberOfRegisteredVoters(_rewardEpochId *big.Int) (*big.Int, error) {
	return _Registry.Contract.GetNumberOfRegisteredVoters(&_Registry.CallOpts, _rewardEpochId)
}

// GetPublicKeyAndNormalisedWeight is a free data retrieval call binding the contract method 0x44b571d9.
//
// Solidity: function getPublicKeyAndNormalisedWeight(uint256 _rewardEpochId, address _signingPolicyAddress) view returns(bytes32 _publicKeyPart1, bytes32 _publicKeyPart2, uint16 _normalisedWeight, uint16 _normalisedWeightsSumOfVotersWithPublicKeys)
func (_Registry *RegistryCaller) GetPublicKeyAndNormalisedWeight(opts *bind.CallOpts, _rewardEpochId *big.Int, _signingPolicyAddress common.Address) (struct {
	PublicKeyPart1                             [32]byte
	PublicKeyPart2                             [32]byte
	NormalisedWeight                           uint16
	NormalisedWeightsSumOfVotersWithPublicKeys uint16
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getPublicKeyAndNormalisedWeight", _rewardEpochId, _signingPolicyAddress)

	outstruct := new(struct {
		PublicKeyPart1                             [32]byte
		PublicKeyPart2                             [32]byte
		NormalisedWeight                           uint16
		NormalisedWeightsSumOfVotersWithPublicKeys uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PublicKeyPart1 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.PublicKeyPart2 = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.NormalisedWeight = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.NormalisedWeightsSumOfVotersWithPublicKeys = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetPublicKeyAndNormalisedWeight is a free data retrieval call binding the contract method 0x44b571d9.
//
// Solidity: function getPublicKeyAndNormalisedWeight(uint256 _rewardEpochId, address _signingPolicyAddress) view returns(bytes32 _publicKeyPart1, bytes32 _publicKeyPart2, uint16 _normalisedWeight, uint16 _normalisedWeightsSumOfVotersWithPublicKeys)
func (_Registry *RegistrySession) GetPublicKeyAndNormalisedWeight(_rewardEpochId *big.Int, _signingPolicyAddress common.Address) (struct {
	PublicKeyPart1                             [32]byte
	PublicKeyPart2                             [32]byte
	NormalisedWeight                           uint16
	NormalisedWeightsSumOfVotersWithPublicKeys uint16
}, error) {
	return _Registry.Contract.GetPublicKeyAndNormalisedWeight(&_Registry.CallOpts, _rewardEpochId, _signingPolicyAddress)
}

// GetPublicKeyAndNormalisedWeight is a free data retrieval call binding the contract method 0x44b571d9.
//
// Solidity: function getPublicKeyAndNormalisedWeight(uint256 _rewardEpochId, address _signingPolicyAddress) view returns(bytes32 _publicKeyPart1, bytes32 _publicKeyPart2, uint16 _normalisedWeight, uint16 _normalisedWeightsSumOfVotersWithPublicKeys)
func (_Registry *RegistryCallerSession) GetPublicKeyAndNormalisedWeight(_rewardEpochId *big.Int, _signingPolicyAddress common.Address) (struct {
	PublicKeyPart1                             [32]byte
	PublicKeyPart2                             [32]byte
	NormalisedWeight                           uint16
	NormalisedWeightsSumOfVotersWithPublicKeys uint16
}, error) {
	return _Registry.Contract.GetPublicKeyAndNormalisedWeight(&_Registry.CallOpts, _rewardEpochId, _signingPolicyAddress)
}

// GetRegisteredDelegationAddresses is a free data retrieval call binding the contract method 0xcc8356ca.
//
// Solidity: function getRegisteredDelegationAddresses(uint256 _rewardEpochId) view returns(address[] _delegationAddresses)
func (_Registry *RegistryCaller) GetRegisteredDelegationAddresses(opts *bind.CallOpts, _rewardEpochId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getRegisteredDelegationAddresses", _rewardEpochId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredDelegationAddresses is a free data retrieval call binding the contract method 0xcc8356ca.
//
// Solidity: function getRegisteredDelegationAddresses(uint256 _rewardEpochId) view returns(address[] _delegationAddresses)
func (_Registry *RegistrySession) GetRegisteredDelegationAddresses(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetRegisteredDelegationAddresses(&_Registry.CallOpts, _rewardEpochId)
}

// GetRegisteredDelegationAddresses is a free data retrieval call binding the contract method 0xcc8356ca.
//
// Solidity: function getRegisteredDelegationAddresses(uint256 _rewardEpochId) view returns(address[] _delegationAddresses)
func (_Registry *RegistryCallerSession) GetRegisteredDelegationAddresses(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetRegisteredDelegationAddresses(&_Registry.CallOpts, _rewardEpochId)
}

// GetRegisteredNodeIds is a free data retrieval call binding the contract method 0x987a4500.
//
// Solidity: function getRegisteredNodeIds(uint256 _rewardEpochId) view returns(bytes20[][] _nodeIds)
func (_Registry *RegistryCaller) GetRegisteredNodeIds(opts *bind.CallOpts, _rewardEpochId *big.Int) ([][][20]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getRegisteredNodeIds", _rewardEpochId)

	if err != nil {
		return *new([][][20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][][20]byte)).(*[][][20]byte)

	return out0, err

}

// GetRegisteredNodeIds is a free data retrieval call binding the contract method 0x987a4500.
//
// Solidity: function getRegisteredNodeIds(uint256 _rewardEpochId) view returns(bytes20[][] _nodeIds)
func (_Registry *RegistrySession) GetRegisteredNodeIds(_rewardEpochId *big.Int) ([][][20]byte, error) {
	return _Registry.Contract.GetRegisteredNodeIds(&_Registry.CallOpts, _rewardEpochId)
}

// GetRegisteredNodeIds is a free data retrieval call binding the contract method 0x987a4500.
//
// Solidity: function getRegisteredNodeIds(uint256 _rewardEpochId) view returns(bytes20[][] _nodeIds)
func (_Registry *RegistryCallerSession) GetRegisteredNodeIds(_rewardEpochId *big.Int) ([][][20]byte, error) {
	return _Registry.Contract.GetRegisteredNodeIds(&_Registry.CallOpts, _rewardEpochId)
}

// GetRegisteredPublicKeys is a free data retrieval call binding the contract method 0x39661f5f.
//
// Solidity: function getRegisteredPublicKeys(uint256 _rewardEpochId) view returns(bytes32[] _parts1, bytes32[] _parts2)
func (_Registry *RegistryCaller) GetRegisteredPublicKeys(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	Parts1 [][32]byte
	Parts2 [][32]byte
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getRegisteredPublicKeys", _rewardEpochId)

	outstruct := new(struct {
		Parts1 [][32]byte
		Parts2 [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Parts1 = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.Parts2 = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

// GetRegisteredPublicKeys is a free data retrieval call binding the contract method 0x39661f5f.
//
// Solidity: function getRegisteredPublicKeys(uint256 _rewardEpochId) view returns(bytes32[] _parts1, bytes32[] _parts2)
func (_Registry *RegistrySession) GetRegisteredPublicKeys(_rewardEpochId *big.Int) (struct {
	Parts1 [][32]byte
	Parts2 [][32]byte
}, error) {
	return _Registry.Contract.GetRegisteredPublicKeys(&_Registry.CallOpts, _rewardEpochId)
}

// GetRegisteredPublicKeys is a free data retrieval call binding the contract method 0x39661f5f.
//
// Solidity: function getRegisteredPublicKeys(uint256 _rewardEpochId) view returns(bytes32[] _parts1, bytes32[] _parts2)
func (_Registry *RegistryCallerSession) GetRegisteredPublicKeys(_rewardEpochId *big.Int) (struct {
	Parts1 [][32]byte
	Parts2 [][32]byte
}, error) {
	return _Registry.Contract.GetRegisteredPublicKeys(&_Registry.CallOpts, _rewardEpochId)
}

// GetRegisteredSigningPolicyAddresses is a free data retrieval call binding the contract method 0x29a2e5ed.
//
// Solidity: function getRegisteredSigningPolicyAddresses(uint256 _rewardEpochId) view returns(address[] _signingPolicyAddresses)
func (_Registry *RegistryCaller) GetRegisteredSigningPolicyAddresses(opts *bind.CallOpts, _rewardEpochId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getRegisteredSigningPolicyAddresses", _rewardEpochId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredSigningPolicyAddresses is a free data retrieval call binding the contract method 0x29a2e5ed.
//
// Solidity: function getRegisteredSigningPolicyAddresses(uint256 _rewardEpochId) view returns(address[] _signingPolicyAddresses)
func (_Registry *RegistrySession) GetRegisteredSigningPolicyAddresses(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetRegisteredSigningPolicyAddresses(&_Registry.CallOpts, _rewardEpochId)
}

// GetRegisteredSigningPolicyAddresses is a free data retrieval call binding the contract method 0x29a2e5ed.
//
// Solidity: function getRegisteredSigningPolicyAddresses(uint256 _rewardEpochId) view returns(address[] _signingPolicyAddresses)
func (_Registry *RegistryCallerSession) GetRegisteredSigningPolicyAddresses(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetRegisteredSigningPolicyAddresses(&_Registry.CallOpts, _rewardEpochId)
}

// GetRegisteredSubmitAddresses is a free data retrieval call binding the contract method 0x7c2cb921.
//
// Solidity: function getRegisteredSubmitAddresses(uint256 _rewardEpochId) view returns(address[])
func (_Registry *RegistryCaller) GetRegisteredSubmitAddresses(opts *bind.CallOpts, _rewardEpochId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getRegisteredSubmitAddresses", _rewardEpochId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredSubmitAddresses is a free data retrieval call binding the contract method 0x7c2cb921.
//
// Solidity: function getRegisteredSubmitAddresses(uint256 _rewardEpochId) view returns(address[])
func (_Registry *RegistrySession) GetRegisteredSubmitAddresses(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetRegisteredSubmitAddresses(&_Registry.CallOpts, _rewardEpochId)
}

// GetRegisteredSubmitAddresses is a free data retrieval call binding the contract method 0x7c2cb921.
//
// Solidity: function getRegisteredSubmitAddresses(uint256 _rewardEpochId) view returns(address[])
func (_Registry *RegistryCallerSession) GetRegisteredSubmitAddresses(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetRegisteredSubmitAddresses(&_Registry.CallOpts, _rewardEpochId)
}

// GetRegisteredSubmitSignaturesAddresses is a free data retrieval call binding the contract method 0xb61b57fa.
//
// Solidity: function getRegisteredSubmitSignaturesAddresses(uint256 _rewardEpochId) view returns(address[] _signingPolicyAddresses)
func (_Registry *RegistryCaller) GetRegisteredSubmitSignaturesAddresses(opts *bind.CallOpts, _rewardEpochId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getRegisteredSubmitSignaturesAddresses", _rewardEpochId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredSubmitSignaturesAddresses is a free data retrieval call binding the contract method 0xb61b57fa.
//
// Solidity: function getRegisteredSubmitSignaturesAddresses(uint256 _rewardEpochId) view returns(address[] _signingPolicyAddresses)
func (_Registry *RegistrySession) GetRegisteredSubmitSignaturesAddresses(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetRegisteredSubmitSignaturesAddresses(&_Registry.CallOpts, _rewardEpochId)
}

// GetRegisteredSubmitSignaturesAddresses is a free data retrieval call binding the contract method 0xb61b57fa.
//
// Solidity: function getRegisteredSubmitSignaturesAddresses(uint256 _rewardEpochId) view returns(address[] _signingPolicyAddresses)
func (_Registry *RegistryCallerSession) GetRegisteredSubmitSignaturesAddresses(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetRegisteredSubmitSignaturesAddresses(&_Registry.CallOpts, _rewardEpochId)
}

// GetRegisteredVoters is a free data retrieval call binding the contract method 0x457c2e47.
//
// Solidity: function getRegisteredVoters(uint256 _rewardEpochId) view returns(address[])
func (_Registry *RegistryCaller) GetRegisteredVoters(opts *bind.CallOpts, _rewardEpochId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getRegisteredVoters", _rewardEpochId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredVoters is a free data retrieval call binding the contract method 0x457c2e47.
//
// Solidity: function getRegisteredVoters(uint256 _rewardEpochId) view returns(address[])
func (_Registry *RegistrySession) GetRegisteredVoters(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetRegisteredVoters(&_Registry.CallOpts, _rewardEpochId)
}

// GetRegisteredVoters is a free data retrieval call binding the contract method 0x457c2e47.
//
// Solidity: function getRegisteredVoters(uint256 _rewardEpochId) view returns(address[])
func (_Registry *RegistryCallerSession) GetRegisteredVoters(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetRegisteredVoters(&_Registry.CallOpts, _rewardEpochId)
}

// GetVoterWithNormalisedWeight is a free data retrieval call binding the contract method 0x8c645728.
//
// Solidity: function getVoterWithNormalisedWeight(uint256 _rewardEpochId, address _signingPolicyAddress) view returns(address _voter, uint16 _normalisedWeight)
func (_Registry *RegistryCaller) GetVoterWithNormalisedWeight(opts *bind.CallOpts, _rewardEpochId *big.Int, _signingPolicyAddress common.Address) (struct {
	Voter            common.Address
	NormalisedWeight uint16
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getVoterWithNormalisedWeight", _rewardEpochId, _signingPolicyAddress)

	outstruct := new(struct {
		Voter            common.Address
		NormalisedWeight uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Voter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NormalisedWeight = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetVoterWithNormalisedWeight is a free data retrieval call binding the contract method 0x8c645728.
//
// Solidity: function getVoterWithNormalisedWeight(uint256 _rewardEpochId, address _signingPolicyAddress) view returns(address _voter, uint16 _normalisedWeight)
func (_Registry *RegistrySession) GetVoterWithNormalisedWeight(_rewardEpochId *big.Int, _signingPolicyAddress common.Address) (struct {
	Voter            common.Address
	NormalisedWeight uint16
}, error) {
	return _Registry.Contract.GetVoterWithNormalisedWeight(&_Registry.CallOpts, _rewardEpochId, _signingPolicyAddress)
}

// GetVoterWithNormalisedWeight is a free data retrieval call binding the contract method 0x8c645728.
//
// Solidity: function getVoterWithNormalisedWeight(uint256 _rewardEpochId, address _signingPolicyAddress) view returns(address _voter, uint16 _normalisedWeight)
func (_Registry *RegistryCallerSession) GetVoterWithNormalisedWeight(_rewardEpochId *big.Int, _signingPolicyAddress common.Address) (struct {
	Voter            common.Address
	NormalisedWeight uint16
}, error) {
	return _Registry.Contract.GetVoterWithNormalisedWeight(&_Registry.CallOpts, _rewardEpochId, _signingPolicyAddress)
}

// GetWeightsSums is a free data retrieval call binding the contract method 0x9508858e.
//
// Solidity: function getWeightsSums(uint256 _rewardEpochId) view returns(uint128 _weightsSum, uint16 _normalisedWeightsSum, uint16 _normalisedWeightsSumOfVotersWithPublicKeys)
func (_Registry *RegistryCaller) GetWeightsSums(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	WeightsSum                                 *big.Int
	NormalisedWeightsSum                       uint16
	NormalisedWeightsSumOfVotersWithPublicKeys uint16
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getWeightsSums", _rewardEpochId)

	outstruct := new(struct {
		WeightsSum                                 *big.Int
		NormalisedWeightsSum                       uint16
		NormalisedWeightsSumOfVotersWithPublicKeys uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WeightsSum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NormalisedWeightsSum = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.NormalisedWeightsSumOfVotersWithPublicKeys = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetWeightsSums is a free data retrieval call binding the contract method 0x9508858e.
//
// Solidity: function getWeightsSums(uint256 _rewardEpochId) view returns(uint128 _weightsSum, uint16 _normalisedWeightsSum, uint16 _normalisedWeightsSumOfVotersWithPublicKeys)
func (_Registry *RegistrySession) GetWeightsSums(_rewardEpochId *big.Int) (struct {
	WeightsSum                                 *big.Int
	NormalisedWeightsSum                       uint16
	NormalisedWeightsSumOfVotersWithPublicKeys uint16
}, error) {
	return _Registry.Contract.GetWeightsSums(&_Registry.CallOpts, _rewardEpochId)
}

// GetWeightsSums is a free data retrieval call binding the contract method 0x9508858e.
//
// Solidity: function getWeightsSums(uint256 _rewardEpochId) view returns(uint128 _weightsSum, uint16 _normalisedWeightsSum, uint16 _normalisedWeightsSumOfVotersWithPublicKeys)
func (_Registry *RegistryCallerSession) GetWeightsSums(_rewardEpochId *big.Int) (struct {
	WeightsSum                                 *big.Int
	NormalisedWeightsSum                       uint16
	NormalisedWeightsSumOfVotersWithPublicKeys uint16
}, error) {
	return _Registry.Contract.GetWeightsSums(&_Registry.CallOpts, _rewardEpochId)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Registry *RegistryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Registry *RegistrySession) Governance() (common.Address, error) {
	return _Registry.Contract.Governance(&_Registry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Registry *RegistryCallerSession) Governance() (common.Address, error) {
	return _Registry.Contract.Governance(&_Registry.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Registry *RegistryCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Registry *RegistrySession) GovernanceSettings() (common.Address, error) {
	return _Registry.Contract.GovernanceSettings(&_Registry.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Registry *RegistryCallerSession) GovernanceSettings() (common.Address, error) {
	return _Registry.Contract.GovernanceSettings(&_Registry.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_Registry *RegistryCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_Registry *RegistrySession) IsExecutor(_address common.Address) (bool, error) {
	return _Registry.Contract.IsExecutor(&_Registry.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_Registry *RegistryCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _Registry.Contract.IsExecutor(&_Registry.CallOpts, _address)
}

// IsVoterRegistered is a free data retrieval call binding the contract method 0x4f5a9968.
//
// Solidity: function isVoterRegistered(address _voter, uint256 _rewardEpochId) view returns(bool)
func (_Registry *RegistryCaller) IsVoterRegistered(opts *bind.CallOpts, _voter common.Address, _rewardEpochId *big.Int) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isVoterRegistered", _voter, _rewardEpochId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoterRegistered is a free data retrieval call binding the contract method 0x4f5a9968.
//
// Solidity: function isVoterRegistered(address _voter, uint256 _rewardEpochId) view returns(bool)
func (_Registry *RegistrySession) IsVoterRegistered(_voter common.Address, _rewardEpochId *big.Int) (bool, error) {
	return _Registry.Contract.IsVoterRegistered(&_Registry.CallOpts, _voter, _rewardEpochId)
}

// IsVoterRegistered is a free data retrieval call binding the contract method 0x4f5a9968.
//
// Solidity: function isVoterRegistered(address _voter, uint256 _rewardEpochId) view returns(bool)
func (_Registry *RegistryCallerSession) IsVoterRegistered(_voter common.Address, _rewardEpochId *big.Int) (bool, error) {
	return _Registry.Contract.IsVoterRegistered(&_Registry.CallOpts, _voter, _rewardEpochId)
}

// MaxVoters is a free data retrieval call binding the contract method 0xd5e50a63.
//
// Solidity: function maxVoters() view returns(uint256)
func (_Registry *RegistryCaller) MaxVoters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "maxVoters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxVoters is a free data retrieval call binding the contract method 0xd5e50a63.
//
// Solidity: function maxVoters() view returns(uint256)
func (_Registry *RegistrySession) MaxVoters() (*big.Int, error) {
	return _Registry.Contract.MaxVoters(&_Registry.CallOpts)
}

// MaxVoters is a free data retrieval call binding the contract method 0xd5e50a63.
//
// Solidity: function maxVoters() view returns(uint256)
func (_Registry *RegistryCallerSession) MaxVoters() (*big.Int, error) {
	return _Registry.Contract.MaxVoters(&_Registry.CallOpts)
}

// NewSigningPolicyInitializationStartBlockNumber is a free data retrieval call binding the contract method 0xfff50753.
//
// Solidity: function newSigningPolicyInitializationStartBlockNumber(uint256 ) view returns(uint256)
func (_Registry *RegistryCaller) NewSigningPolicyInitializationStartBlockNumber(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "newSigningPolicyInitializationStartBlockNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewSigningPolicyInitializationStartBlockNumber is a free data retrieval call binding the contract method 0xfff50753.
//
// Solidity: function newSigningPolicyInitializationStartBlockNumber(uint256 ) view returns(uint256)
func (_Registry *RegistrySession) NewSigningPolicyInitializationStartBlockNumber(arg0 *big.Int) (*big.Int, error) {
	return _Registry.Contract.NewSigningPolicyInitializationStartBlockNumber(&_Registry.CallOpts, arg0)
}

// NewSigningPolicyInitializationStartBlockNumber is a free data retrieval call binding the contract method 0xfff50753.
//
// Solidity: function newSigningPolicyInitializationStartBlockNumber(uint256 ) view returns(uint256)
func (_Registry *RegistryCallerSession) NewSigningPolicyInitializationStartBlockNumber(arg0 *big.Int) (*big.Int, error) {
	return _Registry.Contract.NewSigningPolicyInitializationStartBlockNumber(&_Registry.CallOpts, arg0)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Registry *RegistryCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Registry *RegistrySession) ProductionMode() (bool, error) {
	return _Registry.Contract.ProductionMode(&_Registry.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Registry *RegistryCallerSession) ProductionMode() (bool, error) {
	return _Registry.Contract.ProductionMode(&_Registry.CallOpts)
}

// SystemRegistrationContractAddress is a free data retrieval call binding the contract method 0x2cc2c0f6.
//
// Solidity: function systemRegistrationContractAddress() view returns(address)
func (_Registry *RegistryCaller) SystemRegistrationContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "systemRegistrationContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRegistrationContractAddress is a free data retrieval call binding the contract method 0x2cc2c0f6.
//
// Solidity: function systemRegistrationContractAddress() view returns(address)
func (_Registry *RegistrySession) SystemRegistrationContractAddress() (common.Address, error) {
	return _Registry.Contract.SystemRegistrationContractAddress(&_Registry.CallOpts)
}

// SystemRegistrationContractAddress is a free data retrieval call binding the contract method 0x2cc2c0f6.
//
// Solidity: function systemRegistrationContractAddress() view returns(address)
func (_Registry *RegistryCallerSession) SystemRegistrationContractAddress() (common.Address, error) {
	return _Registry.Contract.SystemRegistrationContractAddress(&_Registry.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Registry *RegistryCaller) TimelockedCalls(opts *bind.CallOpts, arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "timelockedCalls", arg0)

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
func (_Registry *RegistrySession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _Registry.Contract.TimelockedCalls(&_Registry.CallOpts, arg0)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Registry *RegistryCallerSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _Registry.Contract.TimelockedCalls(&_Registry.CallOpts, arg0)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Registry *RegistryTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Registry *RegistrySession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Registry.Contract.CancelGovernanceCall(&_Registry.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Registry *RegistryTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Registry.Contract.CancelGovernanceCall(&_Registry.TransactOpts, _selector)
}

// ChillVoter is a paid mutator transaction binding the contract method 0x201f267e.
//
// Solidity: function chillVoter(address _voter, uint256 _noOfRewardEpochs) returns(uint256 _untilRewardEpochId)
func (_Registry *RegistryTransactor) ChillVoter(opts *bind.TransactOpts, _voter common.Address, _noOfRewardEpochs *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "chillVoter", _voter, _noOfRewardEpochs)
}

// ChillVoter is a paid mutator transaction binding the contract method 0x201f267e.
//
// Solidity: function chillVoter(address _voter, uint256 _noOfRewardEpochs) returns(uint256 _untilRewardEpochId)
func (_Registry *RegistrySession) ChillVoter(_voter common.Address, _noOfRewardEpochs *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.ChillVoter(&_Registry.TransactOpts, _voter, _noOfRewardEpochs)
}

// ChillVoter is a paid mutator transaction binding the contract method 0x201f267e.
//
// Solidity: function chillVoter(address _voter, uint256 _noOfRewardEpochs) returns(uint256 _untilRewardEpochId)
func (_Registry *RegistryTransactorSession) ChillVoter(_voter common.Address, _noOfRewardEpochs *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.ChillVoter(&_Registry.TransactOpts, _voter, _noOfRewardEpochs)
}

// CreateSigningPolicySnapshot is a paid mutator transaction binding the contract method 0xc452e47f.
//
// Solidity: function createSigningPolicySnapshot(uint256 _rewardEpochId) returns(address[] _signingPolicyAddresses, uint16[] _normalisedWeights, uint16 _normalisedWeightsSum)
func (_Registry *RegistryTransactor) CreateSigningPolicySnapshot(opts *bind.TransactOpts, _rewardEpochId *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "createSigningPolicySnapshot", _rewardEpochId)
}

// CreateSigningPolicySnapshot is a paid mutator transaction binding the contract method 0xc452e47f.
//
// Solidity: function createSigningPolicySnapshot(uint256 _rewardEpochId) returns(address[] _signingPolicyAddresses, uint16[] _normalisedWeights, uint16 _normalisedWeightsSum)
func (_Registry *RegistrySession) CreateSigningPolicySnapshot(_rewardEpochId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.CreateSigningPolicySnapshot(&_Registry.TransactOpts, _rewardEpochId)
}

// CreateSigningPolicySnapshot is a paid mutator transaction binding the contract method 0xc452e47f.
//
// Solidity: function createSigningPolicySnapshot(uint256 _rewardEpochId) returns(address[] _signingPolicyAddresses, uint16[] _normalisedWeights, uint16 _normalisedWeightsSum)
func (_Registry *RegistryTransactorSession) CreateSigningPolicySnapshot(_rewardEpochId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.CreateSigningPolicySnapshot(&_Registry.TransactOpts, _rewardEpochId)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Registry *RegistryTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Registry *RegistrySession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Registry.Contract.ExecuteGovernanceCall(&_Registry.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Registry *RegistryTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Registry.Contract.ExecuteGovernanceCall(&_Registry.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_Registry *RegistryTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_Registry *RegistrySession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Initialise(&_Registry.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_Registry *RegistryTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Initialise(&_Registry.TransactOpts, _governanceSettings, _initialGovernance)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x8f7d0957.
//
// Solidity: function registerVoter(address _voter, (uint8,bytes32,bytes32) _signature) returns()
func (_Registry *RegistryTransactor) RegisterVoter(opts *bind.TransactOpts, _voter common.Address, _signature IVoterRegistrySignature) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerVoter", _voter, _signature)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x8f7d0957.
//
// Solidity: function registerVoter(address _voter, (uint8,bytes32,bytes32) _signature) returns()
func (_Registry *RegistrySession) RegisterVoter(_voter common.Address, _signature IVoterRegistrySignature) (*types.Transaction, error) {
	return _Registry.Contract.RegisterVoter(&_Registry.TransactOpts, _voter, _signature)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x8f7d0957.
//
// Solidity: function registerVoter(address _voter, (uint8,bytes32,bytes32) _signature) returns()
func (_Registry *RegistryTransactorSession) RegisterVoter(_voter common.Address, _signature IVoterRegistrySignature) (*types.Transaction, error) {
	return _Registry.Contract.RegisterVoter(&_Registry.TransactOpts, _voter, _signature)
}

// SetMaxVoters is a paid mutator transaction binding the contract method 0xfd587daf.
//
// Solidity: function setMaxVoters(uint256 _maxVoters) returns()
func (_Registry *RegistryTransactor) SetMaxVoters(opts *bind.TransactOpts, _maxVoters *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setMaxVoters", _maxVoters)
}

// SetMaxVoters is a paid mutator transaction binding the contract method 0xfd587daf.
//
// Solidity: function setMaxVoters(uint256 _maxVoters) returns()
func (_Registry *RegistrySession) SetMaxVoters(_maxVoters *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetMaxVoters(&_Registry.TransactOpts, _maxVoters)
}

// SetMaxVoters is a paid mutator transaction binding the contract method 0xfd587daf.
//
// Solidity: function setMaxVoters(uint256 _maxVoters) returns()
func (_Registry *RegistryTransactorSession) SetMaxVoters(_maxVoters *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetMaxVoters(&_Registry.TransactOpts, _maxVoters)
}

// SetNewSigningPolicyInitializationStartBlockNumber is a paid mutator transaction binding the contract method 0x52131823.
//
// Solidity: function setNewSigningPolicyInitializationStartBlockNumber(uint256 _rewardEpochId) returns()
func (_Registry *RegistryTransactor) SetNewSigningPolicyInitializationStartBlockNumber(opts *bind.TransactOpts, _rewardEpochId *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setNewSigningPolicyInitializationStartBlockNumber", _rewardEpochId)
}

// SetNewSigningPolicyInitializationStartBlockNumber is a paid mutator transaction binding the contract method 0x52131823.
//
// Solidity: function setNewSigningPolicyInitializationStartBlockNumber(uint256 _rewardEpochId) returns()
func (_Registry *RegistrySession) SetNewSigningPolicyInitializationStartBlockNumber(_rewardEpochId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetNewSigningPolicyInitializationStartBlockNumber(&_Registry.TransactOpts, _rewardEpochId)
}

// SetNewSigningPolicyInitializationStartBlockNumber is a paid mutator transaction binding the contract method 0x52131823.
//
// Solidity: function setNewSigningPolicyInitializationStartBlockNumber(uint256 _rewardEpochId) returns()
func (_Registry *RegistryTransactorSession) SetNewSigningPolicyInitializationStartBlockNumber(_rewardEpochId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetNewSigningPolicyInitializationStartBlockNumber(&_Registry.TransactOpts, _rewardEpochId)
}

// SetSystemRegistrationContractAddress is a paid mutator transaction binding the contract method 0x448436b1.
//
// Solidity: function setSystemRegistrationContractAddress(address _systemRegistrationContractAddress) returns()
func (_Registry *RegistryTransactor) SetSystemRegistrationContractAddress(opts *bind.TransactOpts, _systemRegistrationContractAddress common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setSystemRegistrationContractAddress", _systemRegistrationContractAddress)
}

// SetSystemRegistrationContractAddress is a paid mutator transaction binding the contract method 0x448436b1.
//
// Solidity: function setSystemRegistrationContractAddress(address _systemRegistrationContractAddress) returns()
func (_Registry *RegistrySession) SetSystemRegistrationContractAddress(_systemRegistrationContractAddress common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetSystemRegistrationContractAddress(&_Registry.TransactOpts, _systemRegistrationContractAddress)
}

// SetSystemRegistrationContractAddress is a paid mutator transaction binding the contract method 0x448436b1.
//
// Solidity: function setSystemRegistrationContractAddress(address _systemRegistrationContractAddress) returns()
func (_Registry *RegistryTransactorSession) SetSystemRegistrationContractAddress(_systemRegistrationContractAddress common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetSystemRegistrationContractAddress(&_Registry.TransactOpts, _systemRegistrationContractAddress)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Registry *RegistryTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Registry *RegistrySession) SwitchToProductionMode() (*types.Transaction, error) {
	return _Registry.Contract.SwitchToProductionMode(&_Registry.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Registry *RegistryTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _Registry.Contract.SwitchToProductionMode(&_Registry.TransactOpts)
}

// SystemRegistration is a paid mutator transaction binding the contract method 0x18d1812d.
//
// Solidity: function systemRegistration(address _voter) returns()
func (_Registry *RegistryTransactor) SystemRegistration(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "systemRegistration", _voter)
}

// SystemRegistration is a paid mutator transaction binding the contract method 0x18d1812d.
//
// Solidity: function systemRegistration(address _voter) returns()
func (_Registry *RegistrySession) SystemRegistration(_voter common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SystemRegistration(&_Registry.TransactOpts, _voter)
}

// SystemRegistration is a paid mutator transaction binding the contract method 0x18d1812d.
//
// Solidity: function systemRegistration(address _voter) returns()
func (_Registry *RegistryTransactorSession) SystemRegistration(_voter common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SystemRegistration(&_Registry.TransactOpts, _voter)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Registry *RegistryTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Registry *RegistrySession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpdateContractAddresses(&_Registry.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Registry *RegistryTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpdateContractAddresses(&_Registry.TransactOpts, _contractNameHashes, _contractAddresses)
}

// RegistryGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the Registry contract.
type RegistryGovernanceCallTimelockedIterator struct {
	Event *RegistryGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *RegistryGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryGovernanceCallTimelocked)
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
		it.Event = new(RegistryGovernanceCallTimelocked)
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
func (it *RegistryGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Registry contract.
type RegistryGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Registry *RegistryFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*RegistryGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &RegistryGovernanceCallTimelockedIterator{contract: _Registry.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Registry *RegistryFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *RegistryGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryGovernanceCallTimelocked)
				if err := _Registry.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseGovernanceCallTimelocked(log types.Log) (*RegistryGovernanceCallTimelocked, error) {
	event := new(RegistryGovernanceCallTimelocked)
	if err := _Registry.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the Registry contract.
type RegistryGovernanceInitialisedIterator struct {
	Event *RegistryGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *RegistryGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryGovernanceInitialised)
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
		it.Event = new(RegistryGovernanceInitialised)
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
func (it *RegistryGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryGovernanceInitialised represents a GovernanceInitialised event raised by the Registry contract.
type RegistryGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_Registry *RegistryFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*RegistryGovernanceInitialisedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &RegistryGovernanceInitialisedIterator{contract: _Registry.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_Registry *RegistryFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *RegistryGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryGovernanceInitialised)
				if err := _Registry.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseGovernanceInitialised(log types.Log) (*RegistryGovernanceInitialised, error) {
	event := new(RegistryGovernanceInitialised)
	if err := _Registry.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the Registry contract.
type RegistryGovernedProductionModeEnteredIterator struct {
	Event *RegistryGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *RegistryGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryGovernedProductionModeEntered)
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
		it.Event = new(RegistryGovernedProductionModeEntered)
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
func (it *RegistryGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the Registry contract.
type RegistryGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_Registry *RegistryFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*RegistryGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &RegistryGovernedProductionModeEnteredIterator{contract: _Registry.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_Registry *RegistryFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *RegistryGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryGovernedProductionModeEntered)
				if err := _Registry.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseGovernedProductionModeEntered(log types.Log) (*RegistryGovernedProductionModeEntered, error) {
	event := new(RegistryGovernedProductionModeEntered)
	if err := _Registry.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the Registry contract.
type RegistryTimelockedGovernanceCallCanceledIterator struct {
	Event *RegistryTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *RegistryTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryTimelockedGovernanceCallCanceled)
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
		it.Event = new(RegistryTimelockedGovernanceCallCanceled)
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
func (it *RegistryTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the Registry contract.
type RegistryTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_Registry *RegistryFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*RegistryTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &RegistryTimelockedGovernanceCallCanceledIterator{contract: _Registry.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_Registry *RegistryFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *RegistryTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryTimelockedGovernanceCallCanceled)
				if err := _Registry.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*RegistryTimelockedGovernanceCallCanceled, error) {
	event := new(RegistryTimelockedGovernanceCallCanceled)
	if err := _Registry.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the Registry contract.
type RegistryTimelockedGovernanceCallExecutedIterator struct {
	Event *RegistryTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *RegistryTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryTimelockedGovernanceCallExecuted)
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
		it.Event = new(RegistryTimelockedGovernanceCallExecuted)
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
func (it *RegistryTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the Registry contract.
type RegistryTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_Registry *RegistryFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*RegistryTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &RegistryTimelockedGovernanceCallExecutedIterator{contract: _Registry.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_Registry *RegistryFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *RegistryTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryTimelockedGovernanceCallExecuted)
				if err := _Registry.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*RegistryTimelockedGovernanceCallExecuted, error) {
	event := new(RegistryTimelockedGovernanceCallExecuted)
	if err := _Registry.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryVoterChilledIterator is returned from FilterVoterChilled and is used to iterate over the raw logs and unpacked data for VoterChilled events raised by the Registry contract.
type RegistryVoterChilledIterator struct {
	Event *RegistryVoterChilled // Event containing the contract specifics and raw log

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
func (it *RegistryVoterChilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryVoterChilled)
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
		it.Event = new(RegistryVoterChilled)
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
func (it *RegistryVoterChilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryVoterChilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryVoterChilled represents a VoterChilled event raised by the Registry contract.
type RegistryVoterChilled struct {
	Voter              common.Address
	UntilRewardEpochId *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterVoterChilled is a free log retrieval operation binding the contract event 0x0c2fcef22ab22997ed46cd27f7f0aa308600145401a7a141065d61c5d87341d2.
//
// Solidity: event VoterChilled(address indexed voter, uint256 untilRewardEpochId)
func (_Registry *RegistryFilterer) FilterVoterChilled(opts *bind.FilterOpts, voter []common.Address) (*RegistryVoterChilledIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "VoterChilled", voterRule)
	if err != nil {
		return nil, err
	}
	return &RegistryVoterChilledIterator{contract: _Registry.contract, event: "VoterChilled", logs: logs, sub: sub}, nil
}

// WatchVoterChilled is a free log subscription operation binding the contract event 0x0c2fcef22ab22997ed46cd27f7f0aa308600145401a7a141065d61c5d87341d2.
//
// Solidity: event VoterChilled(address indexed voter, uint256 untilRewardEpochId)
func (_Registry *RegistryFilterer) WatchVoterChilled(opts *bind.WatchOpts, sink chan<- *RegistryVoterChilled, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "VoterChilled", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryVoterChilled)
				if err := _Registry.contract.UnpackLog(event, "VoterChilled", log); err != nil {
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

// ParseVoterChilled is a log parse operation binding the contract event 0x0c2fcef22ab22997ed46cd27f7f0aa308600145401a7a141065d61c5d87341d2.
//
// Solidity: event VoterChilled(address indexed voter, uint256 untilRewardEpochId)
func (_Registry *RegistryFilterer) ParseVoterChilled(log types.Log) (*RegistryVoterChilled, error) {
	event := new(RegistryVoterChilled)
	if err := _Registry.contract.UnpackLog(event, "VoterChilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryVoterRegisteredIterator is returned from FilterVoterRegistered and is used to iterate over the raw logs and unpacked data for VoterRegistered events raised by the Registry contract.
type RegistryVoterRegisteredIterator struct {
	Event *RegistryVoterRegistered // Event containing the contract specifics and raw log

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
func (it *RegistryVoterRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryVoterRegistered)
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
		it.Event = new(RegistryVoterRegistered)
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
func (it *RegistryVoterRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryVoterRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryVoterRegistered represents a VoterRegistered event raised by the Registry contract.
type RegistryVoterRegistered struct {
	Voter                   common.Address
	RewardEpochId           *big.Int
	SigningPolicyAddress    common.Address
	DelegationAddress       common.Address
	SubmitAddress           common.Address
	SubmitSignaturesAddress common.Address
	RegistrationWeight      *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterVoterRegistered is a free log retrieval operation binding the contract event 0x36e64c7339af9c00a1b00e5da00d8c2159dff48c9c7020a2088b622e11777820.
//
// Solidity: event VoterRegistered(address indexed voter, uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address delegationAddress, address submitAddress, address submitSignaturesAddress, uint256 registrationWeight)
func (_Registry *RegistryFilterer) FilterVoterRegistered(opts *bind.FilterOpts, voter []common.Address, rewardEpochId []*big.Int, signingPolicyAddress []common.Address) (*RegistryVoterRegisteredIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "VoterRegistered", voterRule, rewardEpochIdRule, signingPolicyAddressRule)
	if err != nil {
		return nil, err
	}
	return &RegistryVoterRegisteredIterator{contract: _Registry.contract, event: "VoterRegistered", logs: logs, sub: sub}, nil
}

// WatchVoterRegistered is a free log subscription operation binding the contract event 0x36e64c7339af9c00a1b00e5da00d8c2159dff48c9c7020a2088b622e11777820.
//
// Solidity: event VoterRegistered(address indexed voter, uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address delegationAddress, address submitAddress, address submitSignaturesAddress, uint256 registrationWeight)
func (_Registry *RegistryFilterer) WatchVoterRegistered(opts *bind.WatchOpts, sink chan<- *RegistryVoterRegistered, voter []common.Address, rewardEpochId []*big.Int, signingPolicyAddress []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "VoterRegistered", voterRule, rewardEpochIdRule, signingPolicyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryVoterRegistered)
				if err := _Registry.contract.UnpackLog(event, "VoterRegistered", log); err != nil {
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

// ParseVoterRegistered is a log parse operation binding the contract event 0x36e64c7339af9c00a1b00e5da00d8c2159dff48c9c7020a2088b622e11777820.
//
// Solidity: event VoterRegistered(address indexed voter, uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address delegationAddress, address submitAddress, address submitSignaturesAddress, uint256 registrationWeight)
func (_Registry *RegistryFilterer) ParseVoterRegistered(log types.Log) (*RegistryVoterRegistered, error) {
	event := new(RegistryVoterRegistered)
	if err := _Registry.contract.UnpackLog(event, "VoterRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryVoterRemovedIterator is returned from FilterVoterRemoved and is used to iterate over the raw logs and unpacked data for VoterRemoved events raised by the Registry contract.
type RegistryVoterRemovedIterator struct {
	Event *RegistryVoterRemoved // Event containing the contract specifics and raw log

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
func (it *RegistryVoterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryVoterRemoved)
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
		it.Event = new(RegistryVoterRemoved)
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
func (it *RegistryVoterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryVoterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryVoterRemoved represents a VoterRemoved event raised by the Registry contract.
type RegistryVoterRemoved struct {
	Voter         common.Address
	RewardEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVoterRemoved is a free log retrieval operation binding the contract event 0x98a7f87f8e2aa2f23f43769eff67782bb12946384b142d1ce1e8e38e05d9a3e6.
//
// Solidity: event VoterRemoved(address indexed voter, uint256 indexed rewardEpochId)
func (_Registry *RegistryFilterer) FilterVoterRemoved(opts *bind.FilterOpts, voter []common.Address, rewardEpochId []*big.Int) (*RegistryVoterRemovedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "VoterRemoved", voterRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryVoterRemovedIterator{contract: _Registry.contract, event: "VoterRemoved", logs: logs, sub: sub}, nil
}

// WatchVoterRemoved is a free log subscription operation binding the contract event 0x98a7f87f8e2aa2f23f43769eff67782bb12946384b142d1ce1e8e38e05d9a3e6.
//
// Solidity: event VoterRemoved(address indexed voter, uint256 indexed rewardEpochId)
func (_Registry *RegistryFilterer) WatchVoterRemoved(opts *bind.WatchOpts, sink chan<- *RegistryVoterRemoved, voter []common.Address, rewardEpochId []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "VoterRemoved", voterRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryVoterRemoved)
				if err := _Registry.contract.UnpackLog(event, "VoterRemoved", log); err != nil {
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

// ParseVoterRemoved is a log parse operation binding the contract event 0x98a7f87f8e2aa2f23f43769eff67782bb12946384b142d1ce1e8e38e05d9a3e6.
//
// Solidity: event VoterRemoved(address indexed voter, uint256 indexed rewardEpochId)
func (_Registry *RegistryFilterer) ParseVoterRemoved(log types.Log) (*RegistryVoterRemoved, error) {
	event := new(RegistryVoterRemoved)
	if err := _Registry.contract.UnpackLog(event, "VoterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
