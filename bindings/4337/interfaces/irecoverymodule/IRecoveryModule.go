// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package irecoverymodule

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

// IRecoveryModuleMetaData contains all meta data concerning the IRecoveryModule contract.
var IRecoveryModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"}],\"name\":\"isModuleType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IRecoveryModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use IRecoveryModuleMetaData.ABI instead.
var IRecoveryModuleABI = IRecoveryModuleMetaData.ABI

// IRecoveryModule is an auto generated Go binding around an Ethereum contract.
type IRecoveryModule struct {
	IRecoveryModuleCaller     // Read-only binding to the contract
	IRecoveryModuleTransactor // Write-only binding to the contract
	IRecoveryModuleFilterer   // Log filterer for contract events
}

// IRecoveryModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRecoveryModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRecoveryModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRecoveryModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRecoveryModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRecoveryModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRecoveryModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRecoveryModuleSession struct {
	Contract     *IRecoveryModule  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRecoveryModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRecoveryModuleCallerSession struct {
	Contract *IRecoveryModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IRecoveryModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRecoveryModuleTransactorSession struct {
	Contract     *IRecoveryModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IRecoveryModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRecoveryModuleRaw struct {
	Contract *IRecoveryModule // Generic contract binding to access the raw methods on
}

// IRecoveryModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRecoveryModuleCallerRaw struct {
	Contract *IRecoveryModuleCaller // Generic read-only contract binding to access the raw methods on
}

// IRecoveryModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRecoveryModuleTransactorRaw struct {
	Contract *IRecoveryModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRecoveryModule creates a new instance of IRecoveryModule, bound to a specific deployed contract.
func NewIRecoveryModule(address common.Address, backend bind.ContractBackend) (*IRecoveryModule, error) {
	contract, err := bindIRecoveryModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRecoveryModule{IRecoveryModuleCaller: IRecoveryModuleCaller{contract: contract}, IRecoveryModuleTransactor: IRecoveryModuleTransactor{contract: contract}, IRecoveryModuleFilterer: IRecoveryModuleFilterer{contract: contract}}, nil
}

// NewIRecoveryModuleCaller creates a new read-only instance of IRecoveryModule, bound to a specific deployed contract.
func NewIRecoveryModuleCaller(address common.Address, caller bind.ContractCaller) (*IRecoveryModuleCaller, error) {
	contract, err := bindIRecoveryModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRecoveryModuleCaller{contract: contract}, nil
}

// NewIRecoveryModuleTransactor creates a new write-only instance of IRecoveryModule, bound to a specific deployed contract.
func NewIRecoveryModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*IRecoveryModuleTransactor, error) {
	contract, err := bindIRecoveryModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRecoveryModuleTransactor{contract: contract}, nil
}

// NewIRecoveryModuleFilterer creates a new log filterer instance of IRecoveryModule, bound to a specific deployed contract.
func NewIRecoveryModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*IRecoveryModuleFilterer, error) {
	contract, err := bindIRecoveryModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRecoveryModuleFilterer{contract: contract}, nil
}

// bindIRecoveryModule binds a generic wrapper to an already deployed contract.
func bindIRecoveryModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRecoveryModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRecoveryModule *IRecoveryModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRecoveryModule.Contract.IRecoveryModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRecoveryModule *IRecoveryModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRecoveryModule.Contract.IRecoveryModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRecoveryModule *IRecoveryModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRecoveryModule.Contract.IRecoveryModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRecoveryModule *IRecoveryModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRecoveryModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRecoveryModule *IRecoveryModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRecoveryModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRecoveryModule *IRecoveryModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRecoveryModule.Contract.contract.Transact(opts, method, params...)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) view returns(bool)
func (_IRecoveryModule *IRecoveryModuleCaller) IsModuleType(opts *bind.CallOpts, moduleTypeId *big.Int) (bool, error) {
	var out []interface{}
	err := _IRecoveryModule.contract.Call(opts, &out, "isModuleType", moduleTypeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) view returns(bool)
func (_IRecoveryModule *IRecoveryModuleSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _IRecoveryModule.Contract.IsModuleType(&_IRecoveryModule.CallOpts, moduleTypeId)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) view returns(bool)
func (_IRecoveryModule *IRecoveryModuleCallerSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _IRecoveryModule.Contract.IsModuleType(&_IRecoveryModule.CallOpts, moduleTypeId)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IRecoveryModule *IRecoveryModuleTransactor) OnInstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IRecoveryModule.contract.Transact(opts, "onInstall", data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IRecoveryModule *IRecoveryModuleSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IRecoveryModule.Contract.OnInstall(&_IRecoveryModule.TransactOpts, data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IRecoveryModule *IRecoveryModuleTransactorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IRecoveryModule.Contract.OnInstall(&_IRecoveryModule.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IRecoveryModule *IRecoveryModuleTransactor) OnUninstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IRecoveryModule.contract.Transact(opts, "onUninstall", data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IRecoveryModule *IRecoveryModuleSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IRecoveryModule.Contract.OnUninstall(&_IRecoveryModule.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IRecoveryModule *IRecoveryModuleTransactorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IRecoveryModule.Contract.OnUninstall(&_IRecoveryModule.TransactOpts, data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address _account, bytes _data) returns()
func (_IRecoveryModule *IRecoveryModuleTransactor) Recover(opts *bind.TransactOpts, _account common.Address, _data []byte) (*types.Transaction, error) {
	return _IRecoveryModule.contract.Transact(opts, "recover", _account, _data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address _account, bytes _data) returns()
func (_IRecoveryModule *IRecoveryModuleSession) Recover(_account common.Address, _data []byte) (*types.Transaction, error) {
	return _IRecoveryModule.Contract.Recover(&_IRecoveryModule.TransactOpts, _account, _data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address _account, bytes _data) returns()
func (_IRecoveryModule *IRecoveryModuleTransactorSession) Recover(_account common.Address, _data []byte) (*types.Transaction, error) {
	return _IRecoveryModule.Contract.Recover(&_IRecoveryModule.TransactOpts, _account, _data)
}
