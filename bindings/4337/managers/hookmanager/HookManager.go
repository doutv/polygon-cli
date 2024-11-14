// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hookmanager

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

// HookManagerMetaData contains all meta data concerning the HookManager contract.
var HookManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"executionHook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608080604052346013576082908160198239f35b600080fdfe6080806040526004361015601257600080fd5b600090813560e01c63a09f1f8514602857600080fd5b346048578160031936011260485790546001600160a01b03168152602090f35b5080fdfea2646970667358221220ed2da71c63c1ce899a851d53f55aca90190351b17d737a9d8f3e687c62e97c5a64736f6c63430008190033",
}

// HookManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use HookManagerMetaData.ABI instead.
var HookManagerABI = HookManagerMetaData.ABI

// HookManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HookManagerMetaData.Bin instead.
var HookManagerBin = HookManagerMetaData.Bin

// DeployHookManager deploys a new Ethereum contract, binding an instance of HookManager to it.
func DeployHookManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HookManager, error) {
	parsed, err := HookManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HookManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HookManager{HookManagerCaller: HookManagerCaller{contract: contract}, HookManagerTransactor: HookManagerTransactor{contract: contract}, HookManagerFilterer: HookManagerFilterer{contract: contract}}, nil
}

// HookManager is an auto generated Go binding around an Ethereum contract.
type HookManager struct {
	HookManagerCaller     // Read-only binding to the contract
	HookManagerTransactor // Write-only binding to the contract
	HookManagerFilterer   // Log filterer for contract events
}

// HookManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type HookManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HookManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HookManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HookManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HookManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HookManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HookManagerSession struct {
	Contract     *HookManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HookManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HookManagerCallerSession struct {
	Contract *HookManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HookManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HookManagerTransactorSession struct {
	Contract     *HookManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HookManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type HookManagerRaw struct {
	Contract *HookManager // Generic contract binding to access the raw methods on
}

// HookManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HookManagerCallerRaw struct {
	Contract *HookManagerCaller // Generic read-only contract binding to access the raw methods on
}

// HookManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HookManagerTransactorRaw struct {
	Contract *HookManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHookManager creates a new instance of HookManager, bound to a specific deployed contract.
func NewHookManager(address common.Address, backend bind.ContractBackend) (*HookManager, error) {
	contract, err := bindHookManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HookManager{HookManagerCaller: HookManagerCaller{contract: contract}, HookManagerTransactor: HookManagerTransactor{contract: contract}, HookManagerFilterer: HookManagerFilterer{contract: contract}}, nil
}

// NewHookManagerCaller creates a new read-only instance of HookManager, bound to a specific deployed contract.
func NewHookManagerCaller(address common.Address, caller bind.ContractCaller) (*HookManagerCaller, error) {
	contract, err := bindHookManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HookManagerCaller{contract: contract}, nil
}

// NewHookManagerTransactor creates a new write-only instance of HookManager, bound to a specific deployed contract.
func NewHookManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*HookManagerTransactor, error) {
	contract, err := bindHookManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HookManagerTransactor{contract: contract}, nil
}

// NewHookManagerFilterer creates a new log filterer instance of HookManager, bound to a specific deployed contract.
func NewHookManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*HookManagerFilterer, error) {
	contract, err := bindHookManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HookManagerFilterer{contract: contract}, nil
}

// bindHookManager binds a generic wrapper to an already deployed contract.
func bindHookManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HookManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HookManager *HookManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HookManager.Contract.HookManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HookManager *HookManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HookManager.Contract.HookManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HookManager *HookManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HookManager.Contract.HookManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HookManager *HookManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HookManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HookManager *HookManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HookManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HookManager *HookManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HookManager.Contract.contract.Transact(opts, method, params...)
}

// ExecutionHook is a free data retrieval call binding the contract method 0xa09f1f85.
//
// Solidity: function executionHook() view returns(address)
func (_HookManager *HookManagerCaller) ExecutionHook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HookManager.contract.Call(opts, &out, "executionHook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutionHook is a free data retrieval call binding the contract method 0xa09f1f85.
//
// Solidity: function executionHook() view returns(address)
func (_HookManager *HookManagerSession) ExecutionHook() (common.Address, error) {
	return _HookManager.Contract.ExecutionHook(&_HookManager.CallOpts)
}

// ExecutionHook is a free data retrieval call binding the contract method 0xa09f1f85.
//
// Solidity: function executionHook() view returns(address)
func (_HookManager *HookManagerCallerSession) ExecutionHook() (common.Address, error) {
	return _HookManager.Contract.ExecutionHook(&_HookManager.CallOpts)
}
