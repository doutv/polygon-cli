// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ihook

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

// IHookMetaData contains all meta data concerning the IHook contract.
var IHookMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"}],\"name\":\"isModuleType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"hookData\",\"type\":\"bytes\"}],\"name\":\"postCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"msgData\",\"type\":\"bytes\"}],\"name\":\"preCheck\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"hookData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IHookABI is the input ABI used to generate the binding from.
// Deprecated: Use IHookMetaData.ABI instead.
var IHookABI = IHookMetaData.ABI

// IHook is an auto generated Go binding around an Ethereum contract.
type IHook struct {
	IHookCaller     // Read-only binding to the contract
	IHookTransactor // Write-only binding to the contract
	IHookFilterer   // Log filterer for contract events
}

// IHookCaller is an auto generated read-only Go binding around an Ethereum contract.
type IHookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IHookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IHookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IHookSession struct {
	Contract     *IHook            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IHookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IHookCallerSession struct {
	Contract *IHookCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IHookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IHookTransactorSession struct {
	Contract     *IHookTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IHookRaw is an auto generated low-level Go binding around an Ethereum contract.
type IHookRaw struct {
	Contract *IHook // Generic contract binding to access the raw methods on
}

// IHookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IHookCallerRaw struct {
	Contract *IHookCaller // Generic read-only contract binding to access the raw methods on
}

// IHookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IHookTransactorRaw struct {
	Contract *IHookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIHook creates a new instance of IHook, bound to a specific deployed contract.
func NewIHook(address common.Address, backend bind.ContractBackend) (*IHook, error) {
	contract, err := bindIHook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IHook{IHookCaller: IHookCaller{contract: contract}, IHookTransactor: IHookTransactor{contract: contract}, IHookFilterer: IHookFilterer{contract: contract}}, nil
}

// NewIHookCaller creates a new read-only instance of IHook, bound to a specific deployed contract.
func NewIHookCaller(address common.Address, caller bind.ContractCaller) (*IHookCaller, error) {
	contract, err := bindIHook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IHookCaller{contract: contract}, nil
}

// NewIHookTransactor creates a new write-only instance of IHook, bound to a specific deployed contract.
func NewIHookTransactor(address common.Address, transactor bind.ContractTransactor) (*IHookTransactor, error) {
	contract, err := bindIHook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IHookTransactor{contract: contract}, nil
}

// NewIHookFilterer creates a new log filterer instance of IHook, bound to a specific deployed contract.
func NewIHookFilterer(address common.Address, filterer bind.ContractFilterer) (*IHookFilterer, error) {
	contract, err := bindIHook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IHookFilterer{contract: contract}, nil
}

// bindIHook binds a generic wrapper to an already deployed contract.
func bindIHook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IHookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHook *IHookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHook.Contract.IHookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHook *IHookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHook.Contract.IHookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHook *IHookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHook.Contract.IHookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHook *IHookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHook *IHookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHook *IHookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHook.Contract.contract.Transact(opts, method, params...)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) view returns(bool)
func (_IHook *IHookCaller) IsModuleType(opts *bind.CallOpts, moduleTypeId *big.Int) (bool, error) {
	var out []interface{}
	err := _IHook.contract.Call(opts, &out, "isModuleType", moduleTypeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) view returns(bool)
func (_IHook *IHookSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _IHook.Contract.IsModuleType(&_IHook.CallOpts, moduleTypeId)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) view returns(bool)
func (_IHook *IHookCallerSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _IHook.Contract.IsModuleType(&_IHook.CallOpts, moduleTypeId)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IHook *IHookTransactor) OnInstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IHook.contract.Transact(opts, "onInstall", data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IHook *IHookSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IHook.Contract.OnInstall(&_IHook.TransactOpts, data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IHook *IHookTransactorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IHook.Contract.OnInstall(&_IHook.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IHook *IHookTransactor) OnUninstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IHook.contract.Transact(opts, "onUninstall", data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IHook *IHookSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IHook.Contract.OnUninstall(&_IHook.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IHook *IHookTransactorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IHook.Contract.OnUninstall(&_IHook.TransactOpts, data)
}

// PostCheck is a paid mutator transaction binding the contract method 0x173bf7da.
//
// Solidity: function postCheck(bytes hookData) returns()
func (_IHook *IHookTransactor) PostCheck(opts *bind.TransactOpts, hookData []byte) (*types.Transaction, error) {
	return _IHook.contract.Transact(opts, "postCheck", hookData)
}

// PostCheck is a paid mutator transaction binding the contract method 0x173bf7da.
//
// Solidity: function postCheck(bytes hookData) returns()
func (_IHook *IHookSession) PostCheck(hookData []byte) (*types.Transaction, error) {
	return _IHook.Contract.PostCheck(&_IHook.TransactOpts, hookData)
}

// PostCheck is a paid mutator transaction binding the contract method 0x173bf7da.
//
// Solidity: function postCheck(bytes hookData) returns()
func (_IHook *IHookTransactorSession) PostCheck(hookData []byte) (*types.Transaction, error) {
	return _IHook.Contract.PostCheck(&_IHook.TransactOpts, hookData)
}

// PreCheck is a paid mutator transaction binding the contract method 0xd68f6025.
//
// Solidity: function preCheck(address msgSender, uint256 value, bytes msgData) returns(bytes hookData)
func (_IHook *IHookTransactor) PreCheck(opts *bind.TransactOpts, msgSender common.Address, value *big.Int, msgData []byte) (*types.Transaction, error) {
	return _IHook.contract.Transact(opts, "preCheck", msgSender, value, msgData)
}

// PreCheck is a paid mutator transaction binding the contract method 0xd68f6025.
//
// Solidity: function preCheck(address msgSender, uint256 value, bytes msgData) returns(bytes hookData)
func (_IHook *IHookSession) PreCheck(msgSender common.Address, value *big.Int, msgData []byte) (*types.Transaction, error) {
	return _IHook.Contract.PreCheck(&_IHook.TransactOpts, msgSender, value, msgData)
}

// PreCheck is a paid mutator transaction binding the contract method 0xd68f6025.
//
// Solidity: function preCheck(address msgSender, uint256 value, bytes msgData) returns(bytes hookData)
func (_IHook *IHookTransactorSession) PreCheck(msgSender common.Address, value *big.Int, msgData []byte) (*types.Transaction, error) {
	return _IHook.Contract.PreCheck(&_IHook.TransactOpts, msgSender, value, msgData)
}
