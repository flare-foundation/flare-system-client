// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mirroring

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

// IPChainStakeMirrorVerifierPChainStake is an auto generated low-level Go binding around an user-defined struct.
type IPChainStakeMirrorVerifierPChainStake struct {
	TxId         [32]byte
	StakingType  uint8
	InputAddress [20]byte
	NodeId       [20]byte
	StartTime    uint64
	EndTime      uint64
	Weight       uint64
}

// MirroringMetaData contains all meta data concerning the Mirroring contract.
var MirroringMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"contractFlareDaemon\",\"name\":\"_flareDaemon\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxUpdatesPerBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"CreatedTotalSupplyCache\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUpdatesPerBlock\",\"type\":\"uint256\"}],\"name\":\"MaxUpdatesPerBlockSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"nodeId\",\"type\":\"bytes20\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pChainTxId\",\"type\":\"bytes32\"}],\"name\":\"StakeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"nodeId\",\"type\":\"bytes20\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWei\",\"type\":\"uint256\"}],\"name\":\"StakeEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes20\",\"name\":\"nodeId\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"VotePowerCacheCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"nodeId\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priorVotePower\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotePower\",\"type\":\"uint256\"}],\"name\":\"VotePowerChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressBinder\",\"outputs\":[{\"internalType\":\"contractIAddressBinder\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"balanceHistoryCleanup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20[]\",\"name\":\"_owners\",\"type\":\"bytes20[]\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"batchVotePowerOfAt\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_votePowers\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanerContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanupBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanupBlockNumberManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daemonize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareDaemon\",\"outputs\":[{\"internalType\":\"contractFlareDaemon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceVotePower\",\"outputs\":[{\"internalType\":\"contractIIGovernanceVotePower\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes20\",\"name\":\"_inputAddress\",\"type\":\"bytes20\"}],\"name\":\"isActiveStakeMirrored\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxUpdatesPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"stakingType\",\"type\":\"uint8\"},{\"internalType\":\"bytes20\",\"name\":\"inputAddress\",\"type\":\"bytes20\"},{\"internalType\":\"bytes20\",\"name\":\"nodeId\",\"type\":\"bytes20\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structIPChainStakeMirrorVerifier.PChainStake\",\"name\":\"_stakeData\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"mirrorStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTimestampToTrigger\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cleanerContract\",\"type\":\"address\"}],\"name\":\"setCleanerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"setCleanupBlockNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxUpdatesPerBlock\",\"type\":\"uint256\"}],\"name\":\"setMaxUpdatesPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"stakesHistoryCleanup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"stakesOf\",\"outputs\":[{\"internalType\":\"bytes20[]\",\"name\":\"_nodeIds\",\"type\":\"bytes20[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"stakesOfAt\",\"outputs\":[{\"internalType\":\"bytes20[]\",\"name\":\"_nodeIds\",\"type\":\"bytes20[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToFallbackMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"totalSupplyCacheCleanup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"totalSupplyHistoryCleanup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVotePower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"totalVotePowerAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"totalVotePowerAtCached\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIIPChainStakeMirrorVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_nodeId\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"votePowerCacheCleanup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes20\",\"name\":\"_nodeId\",\"type\":\"bytes20\"}],\"name\":\"votePowerFromTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_votePower\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes20\",\"name\":\"_nodeId\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"votePowerFromToAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_votePower\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_nodeId\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"votePowerHistoryCleanup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_nodeId\",\"type\":\"bytes20\"}],\"name\":\"votePowerOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_nodeId\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"votePowerOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_nodeId\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"votePowerOfAtCached\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MirroringABI is the input ABI used to generate the binding from.
// Deprecated: Use MirroringMetaData.ABI instead.
var MirroringABI = MirroringMetaData.ABI

// Mirroring is an auto generated Go binding around an Ethereum contract.
type Mirroring struct {
	MirroringCaller     // Read-only binding to the contract
	MirroringTransactor // Write-only binding to the contract
	MirroringFilterer   // Log filterer for contract events
}

// MirroringCaller is an auto generated read-only Go binding around an Ethereum contract.
type MirroringCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirroringTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MirroringTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirroringFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MirroringFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirroringSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MirroringSession struct {
	Contract     *Mirroring        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MirroringCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MirroringCallerSession struct {
	Contract *MirroringCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MirroringTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MirroringTransactorSession struct {
	Contract     *MirroringTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MirroringRaw is an auto generated low-level Go binding around an Ethereum contract.
type MirroringRaw struct {
	Contract *Mirroring // Generic contract binding to access the raw methods on
}

// MirroringCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MirroringCallerRaw struct {
	Contract *MirroringCaller // Generic read-only contract binding to access the raw methods on
}

// MirroringTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MirroringTransactorRaw struct {
	Contract *MirroringTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMirroring creates a new instance of Mirroring, bound to a specific deployed contract.
func NewMirroring(address common.Address, backend bind.ContractBackend) (*Mirroring, error) {
	contract, err := bindMirroring(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mirroring{MirroringCaller: MirroringCaller{contract: contract}, MirroringTransactor: MirroringTransactor{contract: contract}, MirroringFilterer: MirroringFilterer{contract: contract}}, nil
}

// NewMirroringCaller creates a new read-only instance of Mirroring, bound to a specific deployed contract.
func NewMirroringCaller(address common.Address, caller bind.ContractCaller) (*MirroringCaller, error) {
	contract, err := bindMirroring(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MirroringCaller{contract: contract}, nil
}

// NewMirroringTransactor creates a new write-only instance of Mirroring, bound to a specific deployed contract.
func NewMirroringTransactor(address common.Address, transactor bind.ContractTransactor) (*MirroringTransactor, error) {
	contract, err := bindMirroring(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MirroringTransactor{contract: contract}, nil
}

// NewMirroringFilterer creates a new log filterer instance of Mirroring, bound to a specific deployed contract.
func NewMirroringFilterer(address common.Address, filterer bind.ContractFilterer) (*MirroringFilterer, error) {
	contract, err := bindMirroring(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MirroringFilterer{contract: contract}, nil
}

// bindMirroring binds a generic wrapper to an already deployed contract.
func bindMirroring(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MirroringMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mirroring *MirroringRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mirroring.Contract.MirroringCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mirroring *MirroringRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mirroring.Contract.MirroringTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mirroring *MirroringRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mirroring.Contract.MirroringTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mirroring *MirroringCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mirroring.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mirroring *MirroringTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mirroring.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mirroring *MirroringTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mirroring.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Mirroring *MirroringCaller) Active(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "active")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Mirroring *MirroringSession) Active() (bool, error) {
	return _Mirroring.Contract.Active(&_Mirroring.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Mirroring *MirroringCallerSession) Active() (bool, error) {
	return _Mirroring.Contract.Active(&_Mirroring.CallOpts)
}

// AddressBinder is a free data retrieval call binding the contract method 0xa546f6ac.
//
// Solidity: function addressBinder() view returns(address)
func (_Mirroring *MirroringCaller) AddressBinder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "addressBinder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressBinder is a free data retrieval call binding the contract method 0xa546f6ac.
//
// Solidity: function addressBinder() view returns(address)
func (_Mirroring *MirroringSession) AddressBinder() (common.Address, error) {
	return _Mirroring.Contract.AddressBinder(&_Mirroring.CallOpts)
}

// AddressBinder is a free data retrieval call binding the contract method 0xa546f6ac.
//
// Solidity: function addressBinder() view returns(address)
func (_Mirroring *MirroringCallerSession) AddressBinder() (common.Address, error) {
	return _Mirroring.Contract.AddressBinder(&_Mirroring.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Mirroring *MirroringCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Mirroring *MirroringSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Mirroring.Contract.BalanceOf(&_Mirroring.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Mirroring *MirroringCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Mirroring.Contract.BalanceOf(&_Mirroring.CallOpts, _owner)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_Mirroring *MirroringCaller) BalanceOfAt(opts *bind.CallOpts, _owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "balanceOfAt", _owner, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_Mirroring *MirroringSession) BalanceOfAt(_owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _Mirroring.Contract.BalanceOfAt(&_Mirroring.CallOpts, _owner, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_Mirroring *MirroringCallerSession) BalanceOfAt(_owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _Mirroring.Contract.BalanceOfAt(&_Mirroring.CallOpts, _owner, _blockNumber)
}

// BatchVotePowerOfAt is a free data retrieval call binding the contract method 0xa9e70199.
//
// Solidity: function batchVotePowerOfAt(bytes20[] _owners, uint256 _blockNumber) view returns(uint256[] _votePowers)
func (_Mirroring *MirroringCaller) BatchVotePowerOfAt(opts *bind.CallOpts, _owners [][20]byte, _blockNumber *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "batchVotePowerOfAt", _owners, _blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BatchVotePowerOfAt is a free data retrieval call binding the contract method 0xa9e70199.
//
// Solidity: function batchVotePowerOfAt(bytes20[] _owners, uint256 _blockNumber) view returns(uint256[] _votePowers)
func (_Mirroring *MirroringSession) BatchVotePowerOfAt(_owners [][20]byte, _blockNumber *big.Int) ([]*big.Int, error) {
	return _Mirroring.Contract.BatchVotePowerOfAt(&_Mirroring.CallOpts, _owners, _blockNumber)
}

// BatchVotePowerOfAt is a free data retrieval call binding the contract method 0xa9e70199.
//
// Solidity: function batchVotePowerOfAt(bytes20[] _owners, uint256 _blockNumber) view returns(uint256[] _votePowers)
func (_Mirroring *MirroringCallerSession) BatchVotePowerOfAt(_owners [][20]byte, _blockNumber *big.Int) ([]*big.Int, error) {
	return _Mirroring.Contract.BatchVotePowerOfAt(&_Mirroring.CallOpts, _owners, _blockNumber)
}

// CleanerContract is a free data retrieval call binding the contract method 0x9ca20e4e.
//
// Solidity: function cleanerContract() view returns(address)
func (_Mirroring *MirroringCaller) CleanerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "cleanerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CleanerContract is a free data retrieval call binding the contract method 0x9ca20e4e.
//
// Solidity: function cleanerContract() view returns(address)
func (_Mirroring *MirroringSession) CleanerContract() (common.Address, error) {
	return _Mirroring.Contract.CleanerContract(&_Mirroring.CallOpts)
}

// CleanerContract is a free data retrieval call binding the contract method 0x9ca20e4e.
//
// Solidity: function cleanerContract() view returns(address)
func (_Mirroring *MirroringCallerSession) CleanerContract() (common.Address, error) {
	return _Mirroring.Contract.CleanerContract(&_Mirroring.CallOpts)
}

// CleanupBlockNumber is a free data retrieval call binding the contract method 0xdeea13e7.
//
// Solidity: function cleanupBlockNumber() view returns(uint256)
func (_Mirroring *MirroringCaller) CleanupBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "cleanupBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CleanupBlockNumber is a free data retrieval call binding the contract method 0xdeea13e7.
//
// Solidity: function cleanupBlockNumber() view returns(uint256)
func (_Mirroring *MirroringSession) CleanupBlockNumber() (*big.Int, error) {
	return _Mirroring.Contract.CleanupBlockNumber(&_Mirroring.CallOpts)
}

// CleanupBlockNumber is a free data retrieval call binding the contract method 0xdeea13e7.
//
// Solidity: function cleanupBlockNumber() view returns(uint256)
func (_Mirroring *MirroringCallerSession) CleanupBlockNumber() (*big.Int, error) {
	return _Mirroring.Contract.CleanupBlockNumber(&_Mirroring.CallOpts)
}

// CleanupBlockNumberManager is a free data retrieval call binding the contract method 0x4eac870f.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (_Mirroring *MirroringCaller) CleanupBlockNumberManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "cleanupBlockNumberManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CleanupBlockNumberManager is a free data retrieval call binding the contract method 0x4eac870f.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (_Mirroring *MirroringSession) CleanupBlockNumberManager() (common.Address, error) {
	return _Mirroring.Contract.CleanupBlockNumberManager(&_Mirroring.CallOpts)
}

// CleanupBlockNumberManager is a free data retrieval call binding the contract method 0x4eac870f.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (_Mirroring *MirroringCallerSession) CleanupBlockNumberManager() (common.Address, error) {
	return _Mirroring.Contract.CleanupBlockNumberManager(&_Mirroring.CallOpts)
}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_Mirroring *MirroringCaller) FlareDaemon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "flareDaemon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_Mirroring *MirroringSession) FlareDaemon() (common.Address, error) {
	return _Mirroring.Contract.FlareDaemon(&_Mirroring.CallOpts)
}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_Mirroring *MirroringCallerSession) FlareDaemon() (common.Address, error) {
	return _Mirroring.Contract.FlareDaemon(&_Mirroring.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Mirroring *MirroringCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Mirroring *MirroringSession) GetAddressUpdater() (common.Address, error) {
	return _Mirroring.Contract.GetAddressUpdater(&_Mirroring.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Mirroring *MirroringCallerSession) GetAddressUpdater() (common.Address, error) {
	return _Mirroring.Contract.GetAddressUpdater(&_Mirroring.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_Mirroring *MirroringCaller) GetContractName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "getContractName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_Mirroring *MirroringSession) GetContractName() (string, error) {
	return _Mirroring.Contract.GetContractName(&_Mirroring.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_Mirroring *MirroringCallerSession) GetContractName() (string, error) {
	return _Mirroring.Contract.GetContractName(&_Mirroring.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Mirroring *MirroringCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Mirroring *MirroringSession) Governance() (common.Address, error) {
	return _Mirroring.Contract.Governance(&_Mirroring.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Mirroring *MirroringCallerSession) Governance() (common.Address, error) {
	return _Mirroring.Contract.Governance(&_Mirroring.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Mirroring *MirroringCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Mirroring *MirroringSession) GovernanceSettings() (common.Address, error) {
	return _Mirroring.Contract.GovernanceSettings(&_Mirroring.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Mirroring *MirroringCallerSession) GovernanceSettings() (common.Address, error) {
	return _Mirroring.Contract.GovernanceSettings(&_Mirroring.CallOpts)
}

// GovernanceVotePower is a free data retrieval call binding the contract method 0x8c2b8ae1.
//
// Solidity: function governanceVotePower() view returns(address)
func (_Mirroring *MirroringCaller) GovernanceVotePower(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "governanceVotePower")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceVotePower is a free data retrieval call binding the contract method 0x8c2b8ae1.
//
// Solidity: function governanceVotePower() view returns(address)
func (_Mirroring *MirroringSession) GovernanceVotePower() (common.Address, error) {
	return _Mirroring.Contract.GovernanceVotePower(&_Mirroring.CallOpts)
}

// GovernanceVotePower is a free data retrieval call binding the contract method 0x8c2b8ae1.
//
// Solidity: function governanceVotePower() view returns(address)
func (_Mirroring *MirroringCallerSession) GovernanceVotePower() (common.Address, error) {
	return _Mirroring.Contract.GovernanceVotePower(&_Mirroring.CallOpts)
}

// IsActiveStakeMirrored is a free data retrieval call binding the contract method 0xd9ab4dfe.
//
// Solidity: function isActiveStakeMirrored(bytes32 _txId, bytes20 _inputAddress) view returns(bool)
func (_Mirroring *MirroringCaller) IsActiveStakeMirrored(opts *bind.CallOpts, _txId [32]byte, _inputAddress [20]byte) (bool, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "isActiveStakeMirrored", _txId, _inputAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveStakeMirrored is a free data retrieval call binding the contract method 0xd9ab4dfe.
//
// Solidity: function isActiveStakeMirrored(bytes32 _txId, bytes20 _inputAddress) view returns(bool)
func (_Mirroring *MirroringSession) IsActiveStakeMirrored(_txId [32]byte, _inputAddress [20]byte) (bool, error) {
	return _Mirroring.Contract.IsActiveStakeMirrored(&_Mirroring.CallOpts, _txId, _inputAddress)
}

// IsActiveStakeMirrored is a free data retrieval call binding the contract method 0xd9ab4dfe.
//
// Solidity: function isActiveStakeMirrored(bytes32 _txId, bytes20 _inputAddress) view returns(bool)
func (_Mirroring *MirroringCallerSession) IsActiveStakeMirrored(_txId [32]byte, _inputAddress [20]byte) (bool, error) {
	return _Mirroring.Contract.IsActiveStakeMirrored(&_Mirroring.CallOpts, _txId, _inputAddress)
}

// MaxUpdatesPerBlock is a free data retrieval call binding the contract method 0xca4af538.
//
// Solidity: function maxUpdatesPerBlock() view returns(uint256)
func (_Mirroring *MirroringCaller) MaxUpdatesPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "maxUpdatesPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxUpdatesPerBlock is a free data retrieval call binding the contract method 0xca4af538.
//
// Solidity: function maxUpdatesPerBlock() view returns(uint256)
func (_Mirroring *MirroringSession) MaxUpdatesPerBlock() (*big.Int, error) {
	return _Mirroring.Contract.MaxUpdatesPerBlock(&_Mirroring.CallOpts)
}

// MaxUpdatesPerBlock is a free data retrieval call binding the contract method 0xca4af538.
//
// Solidity: function maxUpdatesPerBlock() view returns(uint256)
func (_Mirroring *MirroringCallerSession) MaxUpdatesPerBlock() (*big.Int, error) {
	return _Mirroring.Contract.MaxUpdatesPerBlock(&_Mirroring.CallOpts)
}

// NextTimestampToTrigger is a free data retrieval call binding the contract method 0xe5e1ddc6.
//
// Solidity: function nextTimestampToTrigger() view returns(uint256)
func (_Mirroring *MirroringCaller) NextTimestampToTrigger(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "nextTimestampToTrigger")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTimestampToTrigger is a free data retrieval call binding the contract method 0xe5e1ddc6.
//
// Solidity: function nextTimestampToTrigger() view returns(uint256)
func (_Mirroring *MirroringSession) NextTimestampToTrigger() (*big.Int, error) {
	return _Mirroring.Contract.NextTimestampToTrigger(&_Mirroring.CallOpts)
}

// NextTimestampToTrigger is a free data retrieval call binding the contract method 0xe5e1ddc6.
//
// Solidity: function nextTimestampToTrigger() view returns(uint256)
func (_Mirroring *MirroringCallerSession) NextTimestampToTrigger() (*big.Int, error) {
	return _Mirroring.Contract.NextTimestampToTrigger(&_Mirroring.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Mirroring *MirroringCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Mirroring *MirroringSession) ProductionMode() (bool, error) {
	return _Mirroring.Contract.ProductionMode(&_Mirroring.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Mirroring *MirroringCallerSession) ProductionMode() (bool, error) {
	return _Mirroring.Contract.ProductionMode(&_Mirroring.CallOpts)
}

// StakesOf is a free data retrieval call binding the contract method 0x33b69c4c.
//
// Solidity: function stakesOf(address _owner) view returns(bytes20[] _nodeIds, uint256[] _amounts)
func (_Mirroring *MirroringCaller) StakesOf(opts *bind.CallOpts, _owner common.Address) (struct {
	NodeIds [][20]byte
	Amounts []*big.Int
}, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "stakesOf", _owner)

	outstruct := new(struct {
		NodeIds [][20]byte
		Amounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NodeIds = *abi.ConvertType(out[0], new([][20]byte)).(*[][20]byte)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// StakesOf is a free data retrieval call binding the contract method 0x33b69c4c.
//
// Solidity: function stakesOf(address _owner) view returns(bytes20[] _nodeIds, uint256[] _amounts)
func (_Mirroring *MirroringSession) StakesOf(_owner common.Address) (struct {
	NodeIds [][20]byte
	Amounts []*big.Int
}, error) {
	return _Mirroring.Contract.StakesOf(&_Mirroring.CallOpts, _owner)
}

// StakesOf is a free data retrieval call binding the contract method 0x33b69c4c.
//
// Solidity: function stakesOf(address _owner) view returns(bytes20[] _nodeIds, uint256[] _amounts)
func (_Mirroring *MirroringCallerSession) StakesOf(_owner common.Address) (struct {
	NodeIds [][20]byte
	Amounts []*big.Int
}, error) {
	return _Mirroring.Contract.StakesOf(&_Mirroring.CallOpts, _owner)
}

// StakesOfAt is a free data retrieval call binding the contract method 0x4be91f32.
//
// Solidity: function stakesOfAt(address _owner, uint256 _blockNumber) view returns(bytes20[] _nodeIds, uint256[] _amounts)
func (_Mirroring *MirroringCaller) StakesOfAt(opts *bind.CallOpts, _owner common.Address, _blockNumber *big.Int) (struct {
	NodeIds [][20]byte
	Amounts []*big.Int
}, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "stakesOfAt", _owner, _blockNumber)

	outstruct := new(struct {
		NodeIds [][20]byte
		Amounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NodeIds = *abi.ConvertType(out[0], new([][20]byte)).(*[][20]byte)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// StakesOfAt is a free data retrieval call binding the contract method 0x4be91f32.
//
// Solidity: function stakesOfAt(address _owner, uint256 _blockNumber) view returns(bytes20[] _nodeIds, uint256[] _amounts)
func (_Mirroring *MirroringSession) StakesOfAt(_owner common.Address, _blockNumber *big.Int) (struct {
	NodeIds [][20]byte
	Amounts []*big.Int
}, error) {
	return _Mirroring.Contract.StakesOfAt(&_Mirroring.CallOpts, _owner, _blockNumber)
}

// StakesOfAt is a free data retrieval call binding the contract method 0x4be91f32.
//
// Solidity: function stakesOfAt(address _owner, uint256 _blockNumber) view returns(bytes20[] _nodeIds, uint256[] _amounts)
func (_Mirroring *MirroringCallerSession) StakesOfAt(_owner common.Address, _blockNumber *big.Int) (struct {
	NodeIds [][20]byte
	Amounts []*big.Int
}, error) {
	return _Mirroring.Contract.StakesOfAt(&_Mirroring.CallOpts, _owner, _blockNumber)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Mirroring *MirroringCaller) TimelockedCalls(opts *bind.CallOpts, arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "timelockedCalls", arg0)

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
func (_Mirroring *MirroringSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _Mirroring.Contract.TimelockedCalls(&_Mirroring.CallOpts, arg0)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Mirroring *MirroringCallerSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _Mirroring.Contract.TimelockedCalls(&_Mirroring.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mirroring *MirroringCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mirroring *MirroringSession) TotalSupply() (*big.Int, error) {
	return _Mirroring.Contract.TotalSupply(&_Mirroring.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mirroring *MirroringCallerSession) TotalSupply() (*big.Int, error) {
	return _Mirroring.Contract.TotalSupply(&_Mirroring.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _blockNumber) view returns(uint256)
func (_Mirroring *MirroringCaller) TotalSupplyAt(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "totalSupplyAt", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _blockNumber) view returns(uint256)
func (_Mirroring *MirroringSession) TotalSupplyAt(_blockNumber *big.Int) (*big.Int, error) {
	return _Mirroring.Contract.TotalSupplyAt(&_Mirroring.CallOpts, _blockNumber)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _blockNumber) view returns(uint256)
func (_Mirroring *MirroringCallerSession) TotalSupplyAt(_blockNumber *big.Int) (*big.Int, error) {
	return _Mirroring.Contract.TotalSupplyAt(&_Mirroring.CallOpts, _blockNumber)
}

// TotalVotePower is a free data retrieval call binding the contract method 0xf5f3d4f7.
//
// Solidity: function totalVotePower() view returns(uint256)
func (_Mirroring *MirroringCaller) TotalVotePower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "totalVotePower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVotePower is a free data retrieval call binding the contract method 0xf5f3d4f7.
//
// Solidity: function totalVotePower() view returns(uint256)
func (_Mirroring *MirroringSession) TotalVotePower() (*big.Int, error) {
	return _Mirroring.Contract.TotalVotePower(&_Mirroring.CallOpts)
}

// TotalVotePower is a free data retrieval call binding the contract method 0xf5f3d4f7.
//
// Solidity: function totalVotePower() view returns(uint256)
func (_Mirroring *MirroringCallerSession) TotalVotePower() (*big.Int, error) {
	return _Mirroring.Contract.TotalVotePower(&_Mirroring.CallOpts)
}

// TotalVotePowerAt is a free data retrieval call binding the contract method 0x3e5aa26a.
//
// Solidity: function totalVotePowerAt(uint256 _blockNumber) view returns(uint256)
func (_Mirroring *MirroringCaller) TotalVotePowerAt(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "totalVotePowerAt", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVotePowerAt is a free data retrieval call binding the contract method 0x3e5aa26a.
//
// Solidity: function totalVotePowerAt(uint256 _blockNumber) view returns(uint256)
func (_Mirroring *MirroringSession) TotalVotePowerAt(_blockNumber *big.Int) (*big.Int, error) {
	return _Mirroring.Contract.TotalVotePowerAt(&_Mirroring.CallOpts, _blockNumber)
}

// TotalVotePowerAt is a free data retrieval call binding the contract method 0x3e5aa26a.
//
// Solidity: function totalVotePowerAt(uint256 _blockNumber) view returns(uint256)
func (_Mirroring *MirroringCallerSession) TotalVotePowerAt(_blockNumber *big.Int) (*big.Int, error) {
	return _Mirroring.Contract.TotalVotePowerAt(&_Mirroring.CallOpts, _blockNumber)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Mirroring *MirroringCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Mirroring *MirroringSession) Verifier() (common.Address, error) {
	return _Mirroring.Contract.Verifier(&_Mirroring.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Mirroring *MirroringCallerSession) Verifier() (common.Address, error) {
	return _Mirroring.Contract.Verifier(&_Mirroring.CallOpts)
}

// VotePowerFromTo is a free data retrieval call binding the contract method 0x59c345f5.
//
// Solidity: function votePowerFromTo(address _owner, bytes20 _nodeId) view returns(uint256 _votePower)
func (_Mirroring *MirroringCaller) VotePowerFromTo(opts *bind.CallOpts, _owner common.Address, _nodeId [20]byte) (*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "votePowerFromTo", _owner, _nodeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotePowerFromTo is a free data retrieval call binding the contract method 0x59c345f5.
//
// Solidity: function votePowerFromTo(address _owner, bytes20 _nodeId) view returns(uint256 _votePower)
func (_Mirroring *MirroringSession) VotePowerFromTo(_owner common.Address, _nodeId [20]byte) (*big.Int, error) {
	return _Mirroring.Contract.VotePowerFromTo(&_Mirroring.CallOpts, _owner, _nodeId)
}

// VotePowerFromTo is a free data retrieval call binding the contract method 0x59c345f5.
//
// Solidity: function votePowerFromTo(address _owner, bytes20 _nodeId) view returns(uint256 _votePower)
func (_Mirroring *MirroringCallerSession) VotePowerFromTo(_owner common.Address, _nodeId [20]byte) (*big.Int, error) {
	return _Mirroring.Contract.VotePowerFromTo(&_Mirroring.CallOpts, _owner, _nodeId)
}

// VotePowerFromToAt is a free data retrieval call binding the contract method 0x1f7ff2c7.
//
// Solidity: function votePowerFromToAt(address _owner, bytes20 _nodeId, uint256 _blockNumber) view returns(uint256 _votePower)
func (_Mirroring *MirroringCaller) VotePowerFromToAt(opts *bind.CallOpts, _owner common.Address, _nodeId [20]byte, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "votePowerFromToAt", _owner, _nodeId, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotePowerFromToAt is a free data retrieval call binding the contract method 0x1f7ff2c7.
//
// Solidity: function votePowerFromToAt(address _owner, bytes20 _nodeId, uint256 _blockNumber) view returns(uint256 _votePower)
func (_Mirroring *MirroringSession) VotePowerFromToAt(_owner common.Address, _nodeId [20]byte, _blockNumber *big.Int) (*big.Int, error) {
	return _Mirroring.Contract.VotePowerFromToAt(&_Mirroring.CallOpts, _owner, _nodeId, _blockNumber)
}

// VotePowerFromToAt is a free data retrieval call binding the contract method 0x1f7ff2c7.
//
// Solidity: function votePowerFromToAt(address _owner, bytes20 _nodeId, uint256 _blockNumber) view returns(uint256 _votePower)
func (_Mirroring *MirroringCallerSession) VotePowerFromToAt(_owner common.Address, _nodeId [20]byte, _blockNumber *big.Int) (*big.Int, error) {
	return _Mirroring.Contract.VotePowerFromToAt(&_Mirroring.CallOpts, _owner, _nodeId, _blockNumber)
}

// VotePowerOf is a free data retrieval call binding the contract method 0xb4eb2a81.
//
// Solidity: function votePowerOf(bytes20 _nodeId) view returns(uint256)
func (_Mirroring *MirroringCaller) VotePowerOf(opts *bind.CallOpts, _nodeId [20]byte) (*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "votePowerOf", _nodeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotePowerOf is a free data retrieval call binding the contract method 0xb4eb2a81.
//
// Solidity: function votePowerOf(bytes20 _nodeId) view returns(uint256)
func (_Mirroring *MirroringSession) VotePowerOf(_nodeId [20]byte) (*big.Int, error) {
	return _Mirroring.Contract.VotePowerOf(&_Mirroring.CallOpts, _nodeId)
}

// VotePowerOf is a free data retrieval call binding the contract method 0xb4eb2a81.
//
// Solidity: function votePowerOf(bytes20 _nodeId) view returns(uint256)
func (_Mirroring *MirroringCallerSession) VotePowerOf(_nodeId [20]byte) (*big.Int, error) {
	return _Mirroring.Contract.VotePowerOf(&_Mirroring.CallOpts, _nodeId)
}

// VotePowerOfAt is a free data retrieval call binding the contract method 0x46431374.
//
// Solidity: function votePowerOfAt(bytes20 _nodeId, uint256 _blockNumber) view returns(uint256)
func (_Mirroring *MirroringCaller) VotePowerOfAt(opts *bind.CallOpts, _nodeId [20]byte, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mirroring.contract.Call(opts, &out, "votePowerOfAt", _nodeId, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotePowerOfAt is a free data retrieval call binding the contract method 0x46431374.
//
// Solidity: function votePowerOfAt(bytes20 _nodeId, uint256 _blockNumber) view returns(uint256)
func (_Mirroring *MirroringSession) VotePowerOfAt(_nodeId [20]byte, _blockNumber *big.Int) (*big.Int, error) {
	return _Mirroring.Contract.VotePowerOfAt(&_Mirroring.CallOpts, _nodeId, _blockNumber)
}

// VotePowerOfAt is a free data retrieval call binding the contract method 0x46431374.
//
// Solidity: function votePowerOfAt(bytes20 _nodeId, uint256 _blockNumber) view returns(uint256)
func (_Mirroring *MirroringCallerSession) VotePowerOfAt(_nodeId [20]byte, _blockNumber *big.Int) (*big.Int, error) {
	return _Mirroring.Contract.VotePowerOfAt(&_Mirroring.CallOpts, _nodeId, _blockNumber)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Mirroring *MirroringTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Mirroring *MirroringSession) Activate() (*types.Transaction, error) {
	return _Mirroring.Contract.Activate(&_Mirroring.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Mirroring *MirroringTransactorSession) Activate() (*types.Transaction, error) {
	return _Mirroring.Contract.Activate(&_Mirroring.TransactOpts)
}

// BalanceHistoryCleanup is a paid mutator transaction binding the contract method 0xf0e292c9.
//
// Solidity: function balanceHistoryCleanup(address _owner, uint256 _count) returns(uint256)
func (_Mirroring *MirroringTransactor) BalanceHistoryCleanup(opts *bind.TransactOpts, _owner common.Address, _count *big.Int) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "balanceHistoryCleanup", _owner, _count)
}

// BalanceHistoryCleanup is a paid mutator transaction binding the contract method 0xf0e292c9.
//
// Solidity: function balanceHistoryCleanup(address _owner, uint256 _count) returns(uint256)
func (_Mirroring *MirroringSession) BalanceHistoryCleanup(_owner common.Address, _count *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.BalanceHistoryCleanup(&_Mirroring.TransactOpts, _owner, _count)
}

// BalanceHistoryCleanup is a paid mutator transaction binding the contract method 0xf0e292c9.
//
// Solidity: function balanceHistoryCleanup(address _owner, uint256 _count) returns(uint256)
func (_Mirroring *MirroringTransactorSession) BalanceHistoryCleanup(_owner common.Address, _count *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.BalanceHistoryCleanup(&_Mirroring.TransactOpts, _owner, _count)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Mirroring *MirroringTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Mirroring *MirroringSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Mirroring.Contract.CancelGovernanceCall(&_Mirroring.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Mirroring *MirroringTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Mirroring.Contract.CancelGovernanceCall(&_Mirroring.TransactOpts, _selector)
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_Mirroring *MirroringTransactor) Daemonize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "daemonize")
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_Mirroring *MirroringSession) Daemonize() (*types.Transaction, error) {
	return _Mirroring.Contract.Daemonize(&_Mirroring.TransactOpts)
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_Mirroring *MirroringTransactorSession) Daemonize() (*types.Transaction, error) {
	return _Mirroring.Contract.Daemonize(&_Mirroring.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Mirroring *MirroringTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Mirroring *MirroringSession) Deactivate() (*types.Transaction, error) {
	return _Mirroring.Contract.Deactivate(&_Mirroring.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Mirroring *MirroringTransactorSession) Deactivate() (*types.Transaction, error) {
	return _Mirroring.Contract.Deactivate(&_Mirroring.TransactOpts)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Mirroring *MirroringTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Mirroring *MirroringSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Mirroring.Contract.ExecuteGovernanceCall(&_Mirroring.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Mirroring *MirroringTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Mirroring.Contract.ExecuteGovernanceCall(&_Mirroring.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _initialGovernance) returns()
func (_Mirroring *MirroringTransactor) Initialise(opts *bind.TransactOpts, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "initialise", _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _initialGovernance) returns()
func (_Mirroring *MirroringSession) Initialise(_initialGovernance common.Address) (*types.Transaction, error) {
	return _Mirroring.Contract.Initialise(&_Mirroring.TransactOpts, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _initialGovernance) returns()
func (_Mirroring *MirroringTransactorSession) Initialise(_initialGovernance common.Address) (*types.Transaction, error) {
	return _Mirroring.Contract.Initialise(&_Mirroring.TransactOpts, _initialGovernance)
}

// MirrorStake is a paid mutator transaction binding the contract method 0x2e335805.
//
// Solidity: function mirrorStake((bytes32,uint8,bytes20,bytes20,uint64,uint64,uint64) _stakeData, bytes32[] _merkleProof) returns()
func (_Mirroring *MirroringTransactor) MirrorStake(opts *bind.TransactOpts, _stakeData IPChainStakeMirrorVerifierPChainStake, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "mirrorStake", _stakeData, _merkleProof)
}

// MirrorStake is a paid mutator transaction binding the contract method 0x2e335805.
//
// Solidity: function mirrorStake((bytes32,uint8,bytes20,bytes20,uint64,uint64,uint64) _stakeData, bytes32[] _merkleProof) returns()
func (_Mirroring *MirroringSession) MirrorStake(_stakeData IPChainStakeMirrorVerifierPChainStake, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Mirroring.Contract.MirrorStake(&_Mirroring.TransactOpts, _stakeData, _merkleProof)
}

// MirrorStake is a paid mutator transaction binding the contract method 0x2e335805.
//
// Solidity: function mirrorStake((bytes32,uint8,bytes20,bytes20,uint64,uint64,uint64) _stakeData, bytes32[] _merkleProof) returns()
func (_Mirroring *MirroringTransactorSession) MirrorStake(_stakeData IPChainStakeMirrorVerifierPChainStake, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Mirroring.Contract.MirrorStake(&_Mirroring.TransactOpts, _stakeData, _merkleProof)
}

// SetCleanerContract is a paid mutator transaction binding the contract method 0xf6a494af.
//
// Solidity: function setCleanerContract(address _cleanerContract) returns()
func (_Mirroring *MirroringTransactor) SetCleanerContract(opts *bind.TransactOpts, _cleanerContract common.Address) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "setCleanerContract", _cleanerContract)
}

// SetCleanerContract is a paid mutator transaction binding the contract method 0xf6a494af.
//
// Solidity: function setCleanerContract(address _cleanerContract) returns()
func (_Mirroring *MirroringSession) SetCleanerContract(_cleanerContract common.Address) (*types.Transaction, error) {
	return _Mirroring.Contract.SetCleanerContract(&_Mirroring.TransactOpts, _cleanerContract)
}

// SetCleanerContract is a paid mutator transaction binding the contract method 0xf6a494af.
//
// Solidity: function setCleanerContract(address _cleanerContract) returns()
func (_Mirroring *MirroringTransactorSession) SetCleanerContract(_cleanerContract common.Address) (*types.Transaction, error) {
	return _Mirroring.Contract.SetCleanerContract(&_Mirroring.TransactOpts, _cleanerContract)
}

// SetCleanupBlockNumber is a paid mutator transaction binding the contract method 0x13de97f5.
//
// Solidity: function setCleanupBlockNumber(uint256 _blockNumber) returns()
func (_Mirroring *MirroringTransactor) SetCleanupBlockNumber(opts *bind.TransactOpts, _blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "setCleanupBlockNumber", _blockNumber)
}

// SetCleanupBlockNumber is a paid mutator transaction binding the contract method 0x13de97f5.
//
// Solidity: function setCleanupBlockNumber(uint256 _blockNumber) returns()
func (_Mirroring *MirroringSession) SetCleanupBlockNumber(_blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.SetCleanupBlockNumber(&_Mirroring.TransactOpts, _blockNumber)
}

// SetCleanupBlockNumber is a paid mutator transaction binding the contract method 0x13de97f5.
//
// Solidity: function setCleanupBlockNumber(uint256 _blockNumber) returns()
func (_Mirroring *MirroringTransactorSession) SetCleanupBlockNumber(_blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.SetCleanupBlockNumber(&_Mirroring.TransactOpts, _blockNumber)
}

// SetMaxUpdatesPerBlock is a paid mutator transaction binding the contract method 0x4dd4ce49.
//
// Solidity: function setMaxUpdatesPerBlock(uint256 _maxUpdatesPerBlock) returns()
func (_Mirroring *MirroringTransactor) SetMaxUpdatesPerBlock(opts *bind.TransactOpts, _maxUpdatesPerBlock *big.Int) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "setMaxUpdatesPerBlock", _maxUpdatesPerBlock)
}

// SetMaxUpdatesPerBlock is a paid mutator transaction binding the contract method 0x4dd4ce49.
//
// Solidity: function setMaxUpdatesPerBlock(uint256 _maxUpdatesPerBlock) returns()
func (_Mirroring *MirroringSession) SetMaxUpdatesPerBlock(_maxUpdatesPerBlock *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.SetMaxUpdatesPerBlock(&_Mirroring.TransactOpts, _maxUpdatesPerBlock)
}

// SetMaxUpdatesPerBlock is a paid mutator transaction binding the contract method 0x4dd4ce49.
//
// Solidity: function setMaxUpdatesPerBlock(uint256 _maxUpdatesPerBlock) returns()
func (_Mirroring *MirroringTransactorSession) SetMaxUpdatesPerBlock(_maxUpdatesPerBlock *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.SetMaxUpdatesPerBlock(&_Mirroring.TransactOpts, _maxUpdatesPerBlock)
}

// StakesHistoryCleanup is a paid mutator transaction binding the contract method 0xc13edf70.
//
// Solidity: function stakesHistoryCleanup(address _owner, uint256 _count) returns(uint256)
func (_Mirroring *MirroringTransactor) StakesHistoryCleanup(opts *bind.TransactOpts, _owner common.Address, _count *big.Int) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "stakesHistoryCleanup", _owner, _count)
}

// StakesHistoryCleanup is a paid mutator transaction binding the contract method 0xc13edf70.
//
// Solidity: function stakesHistoryCleanup(address _owner, uint256 _count) returns(uint256)
func (_Mirroring *MirroringSession) StakesHistoryCleanup(_owner common.Address, _count *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.StakesHistoryCleanup(&_Mirroring.TransactOpts, _owner, _count)
}

// StakesHistoryCleanup is a paid mutator transaction binding the contract method 0xc13edf70.
//
// Solidity: function stakesHistoryCleanup(address _owner, uint256 _count) returns(uint256)
func (_Mirroring *MirroringTransactorSession) StakesHistoryCleanup(_owner common.Address, _count *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.StakesHistoryCleanup(&_Mirroring.TransactOpts, _owner, _count)
}

// SwitchToFallbackMode is a paid mutator transaction binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() returns(bool)
func (_Mirroring *MirroringTransactor) SwitchToFallbackMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "switchToFallbackMode")
}

// SwitchToFallbackMode is a paid mutator transaction binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() returns(bool)
func (_Mirroring *MirroringSession) SwitchToFallbackMode() (*types.Transaction, error) {
	return _Mirroring.Contract.SwitchToFallbackMode(&_Mirroring.TransactOpts)
}

// SwitchToFallbackMode is a paid mutator transaction binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() returns(bool)
func (_Mirroring *MirroringTransactorSession) SwitchToFallbackMode() (*types.Transaction, error) {
	return _Mirroring.Contract.SwitchToFallbackMode(&_Mirroring.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Mirroring *MirroringTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Mirroring *MirroringSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _Mirroring.Contract.SwitchToProductionMode(&_Mirroring.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Mirroring *MirroringTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _Mirroring.Contract.SwitchToProductionMode(&_Mirroring.TransactOpts)
}

// TotalSupplyCacheCleanup is a paid mutator transaction binding the contract method 0x43ea370b.
//
// Solidity: function totalSupplyCacheCleanup(uint256 _blockNumber) returns(uint256)
func (_Mirroring *MirroringTransactor) TotalSupplyCacheCleanup(opts *bind.TransactOpts, _blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "totalSupplyCacheCleanup", _blockNumber)
}

// TotalSupplyCacheCleanup is a paid mutator transaction binding the contract method 0x43ea370b.
//
// Solidity: function totalSupplyCacheCleanup(uint256 _blockNumber) returns(uint256)
func (_Mirroring *MirroringSession) TotalSupplyCacheCleanup(_blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.TotalSupplyCacheCleanup(&_Mirroring.TransactOpts, _blockNumber)
}

// TotalSupplyCacheCleanup is a paid mutator transaction binding the contract method 0x43ea370b.
//
// Solidity: function totalSupplyCacheCleanup(uint256 _blockNumber) returns(uint256)
func (_Mirroring *MirroringTransactorSession) TotalSupplyCacheCleanup(_blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.TotalSupplyCacheCleanup(&_Mirroring.TransactOpts, _blockNumber)
}

// TotalSupplyHistoryCleanup is a paid mutator transaction binding the contract method 0xf62f8f3a.
//
// Solidity: function totalSupplyHistoryCleanup(uint256 _count) returns(uint256)
func (_Mirroring *MirroringTransactor) TotalSupplyHistoryCleanup(opts *bind.TransactOpts, _count *big.Int) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "totalSupplyHistoryCleanup", _count)
}

// TotalSupplyHistoryCleanup is a paid mutator transaction binding the contract method 0xf62f8f3a.
//
// Solidity: function totalSupplyHistoryCleanup(uint256 _count) returns(uint256)
func (_Mirroring *MirroringSession) TotalSupplyHistoryCleanup(_count *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.TotalSupplyHistoryCleanup(&_Mirroring.TransactOpts, _count)
}

// TotalSupplyHistoryCleanup is a paid mutator transaction binding the contract method 0xf62f8f3a.
//
// Solidity: function totalSupplyHistoryCleanup(uint256 _count) returns(uint256)
func (_Mirroring *MirroringTransactorSession) TotalSupplyHistoryCleanup(_count *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.TotalSupplyHistoryCleanup(&_Mirroring.TransactOpts, _count)
}

// TotalVotePowerAtCached is a paid mutator transaction binding the contract method 0xcaeb942b.
//
// Solidity: function totalVotePowerAtCached(uint256 _blockNumber) returns(uint256)
func (_Mirroring *MirroringTransactor) TotalVotePowerAtCached(opts *bind.TransactOpts, _blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "totalVotePowerAtCached", _blockNumber)
}

// TotalVotePowerAtCached is a paid mutator transaction binding the contract method 0xcaeb942b.
//
// Solidity: function totalVotePowerAtCached(uint256 _blockNumber) returns(uint256)
func (_Mirroring *MirroringSession) TotalVotePowerAtCached(_blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.TotalVotePowerAtCached(&_Mirroring.TransactOpts, _blockNumber)
}

// TotalVotePowerAtCached is a paid mutator transaction binding the contract method 0xcaeb942b.
//
// Solidity: function totalVotePowerAtCached(uint256 _blockNumber) returns(uint256)
func (_Mirroring *MirroringTransactorSession) TotalVotePowerAtCached(_blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.TotalVotePowerAtCached(&_Mirroring.TransactOpts, _blockNumber)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Mirroring *MirroringTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Mirroring *MirroringSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Mirroring.Contract.UpdateContractAddresses(&_Mirroring.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Mirroring *MirroringTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Mirroring.Contract.UpdateContractAddresses(&_Mirroring.TransactOpts, _contractNameHashes, _contractAddresses)
}

// VotePowerCacheCleanup is a paid mutator transaction binding the contract method 0xa2e96264.
//
// Solidity: function votePowerCacheCleanup(bytes20 _nodeId, uint256 _blockNumber) returns(uint256)
func (_Mirroring *MirroringTransactor) VotePowerCacheCleanup(opts *bind.TransactOpts, _nodeId [20]byte, _blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "votePowerCacheCleanup", _nodeId, _blockNumber)
}

// VotePowerCacheCleanup is a paid mutator transaction binding the contract method 0xa2e96264.
//
// Solidity: function votePowerCacheCleanup(bytes20 _nodeId, uint256 _blockNumber) returns(uint256)
func (_Mirroring *MirroringSession) VotePowerCacheCleanup(_nodeId [20]byte, _blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.VotePowerCacheCleanup(&_Mirroring.TransactOpts, _nodeId, _blockNumber)
}

// VotePowerCacheCleanup is a paid mutator transaction binding the contract method 0xa2e96264.
//
// Solidity: function votePowerCacheCleanup(bytes20 _nodeId, uint256 _blockNumber) returns(uint256)
func (_Mirroring *MirroringTransactorSession) VotePowerCacheCleanup(_nodeId [20]byte, _blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.VotePowerCacheCleanup(&_Mirroring.TransactOpts, _nodeId, _blockNumber)
}

// VotePowerHistoryCleanup is a paid mutator transaction binding the contract method 0x119b923a.
//
// Solidity: function votePowerHistoryCleanup(bytes20 _nodeId, uint256 _count) returns(uint256)
func (_Mirroring *MirroringTransactor) VotePowerHistoryCleanup(opts *bind.TransactOpts, _nodeId [20]byte, _count *big.Int) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "votePowerHistoryCleanup", _nodeId, _count)
}

// VotePowerHistoryCleanup is a paid mutator transaction binding the contract method 0x119b923a.
//
// Solidity: function votePowerHistoryCleanup(bytes20 _nodeId, uint256 _count) returns(uint256)
func (_Mirroring *MirroringSession) VotePowerHistoryCleanup(_nodeId [20]byte, _count *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.VotePowerHistoryCleanup(&_Mirroring.TransactOpts, _nodeId, _count)
}

// VotePowerHistoryCleanup is a paid mutator transaction binding the contract method 0x119b923a.
//
// Solidity: function votePowerHistoryCleanup(bytes20 _nodeId, uint256 _count) returns(uint256)
func (_Mirroring *MirroringTransactorSession) VotePowerHistoryCleanup(_nodeId [20]byte, _count *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.VotePowerHistoryCleanup(&_Mirroring.TransactOpts, _nodeId, _count)
}

// VotePowerOfAtCached is a paid mutator transaction binding the contract method 0xbd61ffee.
//
// Solidity: function votePowerOfAtCached(bytes20 _nodeId, uint256 _blockNumber) returns(uint256)
func (_Mirroring *MirroringTransactor) VotePowerOfAtCached(opts *bind.TransactOpts, _nodeId [20]byte, _blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.contract.Transact(opts, "votePowerOfAtCached", _nodeId, _blockNumber)
}

// VotePowerOfAtCached is a paid mutator transaction binding the contract method 0xbd61ffee.
//
// Solidity: function votePowerOfAtCached(bytes20 _nodeId, uint256 _blockNumber) returns(uint256)
func (_Mirroring *MirroringSession) VotePowerOfAtCached(_nodeId [20]byte, _blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.VotePowerOfAtCached(&_Mirroring.TransactOpts, _nodeId, _blockNumber)
}

// VotePowerOfAtCached is a paid mutator transaction binding the contract method 0xbd61ffee.
//
// Solidity: function votePowerOfAtCached(bytes20 _nodeId, uint256 _blockNumber) returns(uint256)
func (_Mirroring *MirroringTransactorSession) VotePowerOfAtCached(_nodeId [20]byte, _blockNumber *big.Int) (*types.Transaction, error) {
	return _Mirroring.Contract.VotePowerOfAtCached(&_Mirroring.TransactOpts, _nodeId, _blockNumber)
}

// MirroringCreatedTotalSupplyCacheIterator is returned from FilterCreatedTotalSupplyCache and is used to iterate over the raw logs and unpacked data for CreatedTotalSupplyCache events raised by the Mirroring contract.
type MirroringCreatedTotalSupplyCacheIterator struct {
	Event *MirroringCreatedTotalSupplyCache // Event containing the contract specifics and raw log

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
func (it *MirroringCreatedTotalSupplyCacheIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirroringCreatedTotalSupplyCache)
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
		it.Event = new(MirroringCreatedTotalSupplyCache)
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
func (it *MirroringCreatedTotalSupplyCacheIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirroringCreatedTotalSupplyCacheIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirroringCreatedTotalSupplyCache represents a CreatedTotalSupplyCache event raised by the Mirroring contract.
type MirroringCreatedTotalSupplyCache struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreatedTotalSupplyCache is a free log retrieval operation binding the contract event 0xfec477a10b4fcdfdf1114eb32b3caf6118b2d76d20e1fcb70007274bb4b616be.
//
// Solidity: event CreatedTotalSupplyCache(uint256 _blockNumber)
func (_Mirroring *MirroringFilterer) FilterCreatedTotalSupplyCache(opts *bind.FilterOpts) (*MirroringCreatedTotalSupplyCacheIterator, error) {

	logs, sub, err := _Mirroring.contract.FilterLogs(opts, "CreatedTotalSupplyCache")
	if err != nil {
		return nil, err
	}
	return &MirroringCreatedTotalSupplyCacheIterator{contract: _Mirroring.contract, event: "CreatedTotalSupplyCache", logs: logs, sub: sub}, nil
}

// WatchCreatedTotalSupplyCache is a free log subscription operation binding the contract event 0xfec477a10b4fcdfdf1114eb32b3caf6118b2d76d20e1fcb70007274bb4b616be.
//
// Solidity: event CreatedTotalSupplyCache(uint256 _blockNumber)
func (_Mirroring *MirroringFilterer) WatchCreatedTotalSupplyCache(opts *bind.WatchOpts, sink chan<- *MirroringCreatedTotalSupplyCache) (event.Subscription, error) {

	logs, sub, err := _Mirroring.contract.WatchLogs(opts, "CreatedTotalSupplyCache")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirroringCreatedTotalSupplyCache)
				if err := _Mirroring.contract.UnpackLog(event, "CreatedTotalSupplyCache", log); err != nil {
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

// ParseCreatedTotalSupplyCache is a log parse operation binding the contract event 0xfec477a10b4fcdfdf1114eb32b3caf6118b2d76d20e1fcb70007274bb4b616be.
//
// Solidity: event CreatedTotalSupplyCache(uint256 _blockNumber)
func (_Mirroring *MirroringFilterer) ParseCreatedTotalSupplyCache(log types.Log) (*MirroringCreatedTotalSupplyCache, error) {
	event := new(MirroringCreatedTotalSupplyCache)
	if err := _Mirroring.contract.UnpackLog(event, "CreatedTotalSupplyCache", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirroringGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the Mirroring contract.
type MirroringGovernanceCallTimelockedIterator struct {
	Event *MirroringGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *MirroringGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirroringGovernanceCallTimelocked)
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
		it.Event = new(MirroringGovernanceCallTimelocked)
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
func (it *MirroringGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirroringGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirroringGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Mirroring contract.
type MirroringGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Mirroring *MirroringFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*MirroringGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _Mirroring.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &MirroringGovernanceCallTimelockedIterator{contract: _Mirroring.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Mirroring *MirroringFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *MirroringGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _Mirroring.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirroringGovernanceCallTimelocked)
				if err := _Mirroring.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_Mirroring *MirroringFilterer) ParseGovernanceCallTimelocked(log types.Log) (*MirroringGovernanceCallTimelocked, error) {
	event := new(MirroringGovernanceCallTimelocked)
	if err := _Mirroring.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirroringGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the Mirroring contract.
type MirroringGovernanceInitialisedIterator struct {
	Event *MirroringGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *MirroringGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirroringGovernanceInitialised)
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
		it.Event = new(MirroringGovernanceInitialised)
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
func (it *MirroringGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirroringGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirroringGovernanceInitialised represents a GovernanceInitialised event raised by the Mirroring contract.
type MirroringGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_Mirroring *MirroringFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*MirroringGovernanceInitialisedIterator, error) {

	logs, sub, err := _Mirroring.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &MirroringGovernanceInitialisedIterator{contract: _Mirroring.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_Mirroring *MirroringFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *MirroringGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _Mirroring.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirroringGovernanceInitialised)
				if err := _Mirroring.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_Mirroring *MirroringFilterer) ParseGovernanceInitialised(log types.Log) (*MirroringGovernanceInitialised, error) {
	event := new(MirroringGovernanceInitialised)
	if err := _Mirroring.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirroringGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the Mirroring contract.
type MirroringGovernedProductionModeEnteredIterator struct {
	Event *MirroringGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *MirroringGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirroringGovernedProductionModeEntered)
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
		it.Event = new(MirroringGovernedProductionModeEntered)
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
func (it *MirroringGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirroringGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirroringGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the Mirroring contract.
type MirroringGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_Mirroring *MirroringFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*MirroringGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _Mirroring.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &MirroringGovernedProductionModeEnteredIterator{contract: _Mirroring.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_Mirroring *MirroringFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *MirroringGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _Mirroring.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirroringGovernedProductionModeEntered)
				if err := _Mirroring.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_Mirroring *MirroringFilterer) ParseGovernedProductionModeEntered(log types.Log) (*MirroringGovernedProductionModeEntered, error) {
	event := new(MirroringGovernedProductionModeEntered)
	if err := _Mirroring.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirroringMaxUpdatesPerBlockSetIterator is returned from FilterMaxUpdatesPerBlockSet and is used to iterate over the raw logs and unpacked data for MaxUpdatesPerBlockSet events raised by the Mirroring contract.
type MirroringMaxUpdatesPerBlockSetIterator struct {
	Event *MirroringMaxUpdatesPerBlockSet // Event containing the contract specifics and raw log

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
func (it *MirroringMaxUpdatesPerBlockSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirroringMaxUpdatesPerBlockSet)
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
		it.Event = new(MirroringMaxUpdatesPerBlockSet)
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
func (it *MirroringMaxUpdatesPerBlockSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirroringMaxUpdatesPerBlockSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirroringMaxUpdatesPerBlockSet represents a MaxUpdatesPerBlockSet event raised by the Mirroring contract.
type MirroringMaxUpdatesPerBlockSet struct {
	MaxUpdatesPerBlock *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMaxUpdatesPerBlockSet is a free log retrieval operation binding the contract event 0x45e418145c8aa16b8f8c5a80e9589a7f2e054bac2d90cd0f05e19f40cdd37a75.
//
// Solidity: event MaxUpdatesPerBlockSet(uint256 maxUpdatesPerBlock)
func (_Mirroring *MirroringFilterer) FilterMaxUpdatesPerBlockSet(opts *bind.FilterOpts) (*MirroringMaxUpdatesPerBlockSetIterator, error) {

	logs, sub, err := _Mirroring.contract.FilterLogs(opts, "MaxUpdatesPerBlockSet")
	if err != nil {
		return nil, err
	}
	return &MirroringMaxUpdatesPerBlockSetIterator{contract: _Mirroring.contract, event: "MaxUpdatesPerBlockSet", logs: logs, sub: sub}, nil
}

// WatchMaxUpdatesPerBlockSet is a free log subscription operation binding the contract event 0x45e418145c8aa16b8f8c5a80e9589a7f2e054bac2d90cd0f05e19f40cdd37a75.
//
// Solidity: event MaxUpdatesPerBlockSet(uint256 maxUpdatesPerBlock)
func (_Mirroring *MirroringFilterer) WatchMaxUpdatesPerBlockSet(opts *bind.WatchOpts, sink chan<- *MirroringMaxUpdatesPerBlockSet) (event.Subscription, error) {

	logs, sub, err := _Mirroring.contract.WatchLogs(opts, "MaxUpdatesPerBlockSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirroringMaxUpdatesPerBlockSet)
				if err := _Mirroring.contract.UnpackLog(event, "MaxUpdatesPerBlockSet", log); err != nil {
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

// ParseMaxUpdatesPerBlockSet is a log parse operation binding the contract event 0x45e418145c8aa16b8f8c5a80e9589a7f2e054bac2d90cd0f05e19f40cdd37a75.
//
// Solidity: event MaxUpdatesPerBlockSet(uint256 maxUpdatesPerBlock)
func (_Mirroring *MirroringFilterer) ParseMaxUpdatesPerBlockSet(log types.Log) (*MirroringMaxUpdatesPerBlockSet, error) {
	event := new(MirroringMaxUpdatesPerBlockSet)
	if err := _Mirroring.contract.UnpackLog(event, "MaxUpdatesPerBlockSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirroringStakeConfirmedIterator is returned from FilterStakeConfirmed and is used to iterate over the raw logs and unpacked data for StakeConfirmed events raised by the Mirroring contract.
type MirroringStakeConfirmedIterator struct {
	Event *MirroringStakeConfirmed // Event containing the contract specifics and raw log

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
func (it *MirroringStakeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirroringStakeConfirmed)
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
		it.Event = new(MirroringStakeConfirmed)
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
func (it *MirroringStakeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirroringStakeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirroringStakeConfirmed represents a StakeConfirmed event raised by the Mirroring contract.
type MirroringStakeConfirmed struct {
	Owner      common.Address
	NodeId     [20]byte
	TxHash     [32]byte
	AmountWei  *big.Int
	PChainTxId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStakeConfirmed is a free log retrieval operation binding the contract event 0x4963ffc566fe8fe6af0920b81985e6514afa925fe3408430f490ea3b61e548cd.
//
// Solidity: event StakeConfirmed(address indexed owner, bytes20 indexed nodeId, bytes32 indexed txHash, uint256 amountWei, bytes32 pChainTxId)
func (_Mirroring *MirroringFilterer) FilterStakeConfirmed(opts *bind.FilterOpts, owner []common.Address, nodeId [][20]byte, txHash [][32]byte) (*MirroringStakeConfirmedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _Mirroring.contract.FilterLogs(opts, "StakeConfirmed", ownerRule, nodeIdRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return &MirroringStakeConfirmedIterator{contract: _Mirroring.contract, event: "StakeConfirmed", logs: logs, sub: sub}, nil
}

// WatchStakeConfirmed is a free log subscription operation binding the contract event 0x4963ffc566fe8fe6af0920b81985e6514afa925fe3408430f490ea3b61e548cd.
//
// Solidity: event StakeConfirmed(address indexed owner, bytes20 indexed nodeId, bytes32 indexed txHash, uint256 amountWei, bytes32 pChainTxId)
func (_Mirroring *MirroringFilterer) WatchStakeConfirmed(opts *bind.WatchOpts, sink chan<- *MirroringStakeConfirmed, owner []common.Address, nodeId [][20]byte, txHash [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _Mirroring.contract.WatchLogs(opts, "StakeConfirmed", ownerRule, nodeIdRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirroringStakeConfirmed)
				if err := _Mirroring.contract.UnpackLog(event, "StakeConfirmed", log); err != nil {
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

// ParseStakeConfirmed is a log parse operation binding the contract event 0x4963ffc566fe8fe6af0920b81985e6514afa925fe3408430f490ea3b61e548cd.
//
// Solidity: event StakeConfirmed(address indexed owner, bytes20 indexed nodeId, bytes32 indexed txHash, uint256 amountWei, bytes32 pChainTxId)
func (_Mirroring *MirroringFilterer) ParseStakeConfirmed(log types.Log) (*MirroringStakeConfirmed, error) {
	event := new(MirroringStakeConfirmed)
	if err := _Mirroring.contract.UnpackLog(event, "StakeConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirroringStakeEndedIterator is returned from FilterStakeEnded and is used to iterate over the raw logs and unpacked data for StakeEnded events raised by the Mirroring contract.
type MirroringStakeEndedIterator struct {
	Event *MirroringStakeEnded // Event containing the contract specifics and raw log

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
func (it *MirroringStakeEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirroringStakeEnded)
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
		it.Event = new(MirroringStakeEnded)
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
func (it *MirroringStakeEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirroringStakeEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirroringStakeEnded represents a StakeEnded event raised by the Mirroring contract.
type MirroringStakeEnded struct {
	Owner     common.Address
	NodeId    [20]byte
	TxHash    [32]byte
	AmountWei *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeEnded is a free log retrieval operation binding the contract event 0xd89d725ca28c0d5002ce399c8a15b0a26908a79c36244c807f9f0b9878f7e756.
//
// Solidity: event StakeEnded(address indexed owner, bytes20 indexed nodeId, bytes32 indexed txHash, uint256 amountWei)
func (_Mirroring *MirroringFilterer) FilterStakeEnded(opts *bind.FilterOpts, owner []common.Address, nodeId [][20]byte, txHash [][32]byte) (*MirroringStakeEndedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _Mirroring.contract.FilterLogs(opts, "StakeEnded", ownerRule, nodeIdRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return &MirroringStakeEndedIterator{contract: _Mirroring.contract, event: "StakeEnded", logs: logs, sub: sub}, nil
}

// WatchStakeEnded is a free log subscription operation binding the contract event 0xd89d725ca28c0d5002ce399c8a15b0a26908a79c36244c807f9f0b9878f7e756.
//
// Solidity: event StakeEnded(address indexed owner, bytes20 indexed nodeId, bytes32 indexed txHash, uint256 amountWei)
func (_Mirroring *MirroringFilterer) WatchStakeEnded(opts *bind.WatchOpts, sink chan<- *MirroringStakeEnded, owner []common.Address, nodeId [][20]byte, txHash [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _Mirroring.contract.WatchLogs(opts, "StakeEnded", ownerRule, nodeIdRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirroringStakeEnded)
				if err := _Mirroring.contract.UnpackLog(event, "StakeEnded", log); err != nil {
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

// ParseStakeEnded is a log parse operation binding the contract event 0xd89d725ca28c0d5002ce399c8a15b0a26908a79c36244c807f9f0b9878f7e756.
//
// Solidity: event StakeEnded(address indexed owner, bytes20 indexed nodeId, bytes32 indexed txHash, uint256 amountWei)
func (_Mirroring *MirroringFilterer) ParseStakeEnded(log types.Log) (*MirroringStakeEnded, error) {
	event := new(MirroringStakeEnded)
	if err := _Mirroring.contract.UnpackLog(event, "StakeEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirroringTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the Mirroring contract.
type MirroringTimelockedGovernanceCallCanceledIterator struct {
	Event *MirroringTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *MirroringTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirroringTimelockedGovernanceCallCanceled)
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
		it.Event = new(MirroringTimelockedGovernanceCallCanceled)
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
func (it *MirroringTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirroringTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirroringTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the Mirroring contract.
type MirroringTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_Mirroring *MirroringFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*MirroringTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _Mirroring.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &MirroringTimelockedGovernanceCallCanceledIterator{contract: _Mirroring.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_Mirroring *MirroringFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *MirroringTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _Mirroring.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirroringTimelockedGovernanceCallCanceled)
				if err := _Mirroring.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_Mirroring *MirroringFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*MirroringTimelockedGovernanceCallCanceled, error) {
	event := new(MirroringTimelockedGovernanceCallCanceled)
	if err := _Mirroring.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirroringTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the Mirroring contract.
type MirroringTimelockedGovernanceCallExecutedIterator struct {
	Event *MirroringTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *MirroringTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirroringTimelockedGovernanceCallExecuted)
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
		it.Event = new(MirroringTimelockedGovernanceCallExecuted)
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
func (it *MirroringTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirroringTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirroringTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the Mirroring contract.
type MirroringTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_Mirroring *MirroringFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*MirroringTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _Mirroring.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &MirroringTimelockedGovernanceCallExecutedIterator{contract: _Mirroring.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_Mirroring *MirroringFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *MirroringTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _Mirroring.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirroringTimelockedGovernanceCallExecuted)
				if err := _Mirroring.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_Mirroring *MirroringFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*MirroringTimelockedGovernanceCallExecuted, error) {
	event := new(MirroringTimelockedGovernanceCallExecuted)
	if err := _Mirroring.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirroringVotePowerCacheCreatedIterator is returned from FilterVotePowerCacheCreated and is used to iterate over the raw logs and unpacked data for VotePowerCacheCreated events raised by the Mirroring contract.
type MirroringVotePowerCacheCreatedIterator struct {
	Event *MirroringVotePowerCacheCreated // Event containing the contract specifics and raw log

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
func (it *MirroringVotePowerCacheCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirroringVotePowerCacheCreated)
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
		it.Event = new(MirroringVotePowerCacheCreated)
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
func (it *MirroringVotePowerCacheCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirroringVotePowerCacheCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirroringVotePowerCacheCreated represents a VotePowerCacheCreated event raised by the Mirroring contract.
type MirroringVotePowerCacheCreated struct {
	NodeId      [20]byte
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVotePowerCacheCreated is a free log retrieval operation binding the contract event 0xb34d3f5bc1f5b60002062672fcc0a956db7b59b3416f8c5343e9aef5adc7b971.
//
// Solidity: event VotePowerCacheCreated(bytes20 nodeId, uint256 blockNumber)
func (_Mirroring *MirroringFilterer) FilterVotePowerCacheCreated(opts *bind.FilterOpts) (*MirroringVotePowerCacheCreatedIterator, error) {

	logs, sub, err := _Mirroring.contract.FilterLogs(opts, "VotePowerCacheCreated")
	if err != nil {
		return nil, err
	}
	return &MirroringVotePowerCacheCreatedIterator{contract: _Mirroring.contract, event: "VotePowerCacheCreated", logs: logs, sub: sub}, nil
}

// WatchVotePowerCacheCreated is a free log subscription operation binding the contract event 0xb34d3f5bc1f5b60002062672fcc0a956db7b59b3416f8c5343e9aef5adc7b971.
//
// Solidity: event VotePowerCacheCreated(bytes20 nodeId, uint256 blockNumber)
func (_Mirroring *MirroringFilterer) WatchVotePowerCacheCreated(opts *bind.WatchOpts, sink chan<- *MirroringVotePowerCacheCreated) (event.Subscription, error) {

	logs, sub, err := _Mirroring.contract.WatchLogs(opts, "VotePowerCacheCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirroringVotePowerCacheCreated)
				if err := _Mirroring.contract.UnpackLog(event, "VotePowerCacheCreated", log); err != nil {
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

// ParseVotePowerCacheCreated is a log parse operation binding the contract event 0xb34d3f5bc1f5b60002062672fcc0a956db7b59b3416f8c5343e9aef5adc7b971.
//
// Solidity: event VotePowerCacheCreated(bytes20 nodeId, uint256 blockNumber)
func (_Mirroring *MirroringFilterer) ParseVotePowerCacheCreated(log types.Log) (*MirroringVotePowerCacheCreated, error) {
	event := new(MirroringVotePowerCacheCreated)
	if err := _Mirroring.contract.UnpackLog(event, "VotePowerCacheCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirroringVotePowerChangedIterator is returned from FilterVotePowerChanged and is used to iterate over the raw logs and unpacked data for VotePowerChanged events raised by the Mirroring contract.
type MirroringVotePowerChangedIterator struct {
	Event *MirroringVotePowerChanged // Event containing the contract specifics and raw log

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
func (it *MirroringVotePowerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirroringVotePowerChanged)
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
		it.Event = new(MirroringVotePowerChanged)
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
func (it *MirroringVotePowerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirroringVotePowerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirroringVotePowerChanged represents a VotePowerChanged event raised by the Mirroring contract.
type MirroringVotePowerChanged struct {
	Owner          common.Address
	NodeId         [20]byte
	PriorVotePower *big.Int
	NewVotePower   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVotePowerChanged is a free log retrieval operation binding the contract event 0xe03ab1522dc81fa0410fe7c9668da7d7b8b9be42ea6011d936e90deda8c0aea1.
//
// Solidity: event VotePowerChanged(address indexed owner, bytes20 indexed nodeId, uint256 priorVotePower, uint256 newVotePower)
func (_Mirroring *MirroringFilterer) FilterVotePowerChanged(opts *bind.FilterOpts, owner []common.Address, nodeId [][20]byte) (*MirroringVotePowerChangedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _Mirroring.contract.FilterLogs(opts, "VotePowerChanged", ownerRule, nodeIdRule)
	if err != nil {
		return nil, err
	}
	return &MirroringVotePowerChangedIterator{contract: _Mirroring.contract, event: "VotePowerChanged", logs: logs, sub: sub}, nil
}

// WatchVotePowerChanged is a free log subscription operation binding the contract event 0xe03ab1522dc81fa0410fe7c9668da7d7b8b9be42ea6011d936e90deda8c0aea1.
//
// Solidity: event VotePowerChanged(address indexed owner, bytes20 indexed nodeId, uint256 priorVotePower, uint256 newVotePower)
func (_Mirroring *MirroringFilterer) WatchVotePowerChanged(opts *bind.WatchOpts, sink chan<- *MirroringVotePowerChanged, owner []common.Address, nodeId [][20]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _Mirroring.contract.WatchLogs(opts, "VotePowerChanged", ownerRule, nodeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirroringVotePowerChanged)
				if err := _Mirroring.contract.UnpackLog(event, "VotePowerChanged", log); err != nil {
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

// ParseVotePowerChanged is a log parse operation binding the contract event 0xe03ab1522dc81fa0410fe7c9668da7d7b8b9be42ea6011d936e90deda8c0aea1.
//
// Solidity: event VotePowerChanged(address indexed owner, bytes20 indexed nodeId, uint256 priorVotePower, uint256 newVotePower)
func (_Mirroring *MirroringFilterer) ParseVotePowerChanged(log types.Log) (*MirroringVotePowerChanged, error) {
	event := new(MirroringVotePowerChanged)
	if err := _Mirroring.contract.UnpackLog(event, "VotePowerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
