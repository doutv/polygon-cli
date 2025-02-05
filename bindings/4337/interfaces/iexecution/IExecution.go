// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iexecution

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

// IExecutionMetaData contains all meta data concerning the IExecution contract.
var IExecutionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"executionCalldata\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"executionCalldata\",\"type\":\"bytes\"}],\"name\":\"executeFromExecutor\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IExecutionABI is the input ABI used to generate the binding from.
// Deprecated: Use IExecutionMetaData.ABI instead.
var IExecutionABI = IExecutionMetaData.ABI

// IExecution is an auto generated Go binding around an Ethereum contract.
type IExecution struct {
	IExecutionCaller     // Read-only binding to the contract
	IExecutionTransactor // Write-only binding to the contract
	IExecutionFilterer   // Log filterer for contract events
}

// IExecutionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExecutionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExecutionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExecutionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExecutionSession struct {
	Contract     *IExecution       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IExecutionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExecutionCallerSession struct {
	Contract *IExecutionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IExecutionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExecutionTransactorSession struct {
	Contract     *IExecutionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IExecutionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExecutionRaw struct {
	Contract *IExecution // Generic contract binding to access the raw methods on
}

// IExecutionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExecutionCallerRaw struct {
	Contract *IExecutionCaller // Generic read-only contract binding to access the raw methods on
}

// IExecutionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExecutionTransactorRaw struct {
	Contract *IExecutionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExecution creates a new instance of IExecution, bound to a specific deployed contract.
func NewIExecution(address common.Address, backend bind.ContractBackend) (*IExecution, error) {
	contract, err := bindIExecution(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExecution{IExecutionCaller: IExecutionCaller{contract: contract}, IExecutionTransactor: IExecutionTransactor{contract: contract}, IExecutionFilterer: IExecutionFilterer{contract: contract}}, nil
}

// NewIExecutionCaller creates a new read-only instance of IExecution, bound to a specific deployed contract.
func NewIExecutionCaller(address common.Address, caller bind.ContractCaller) (*IExecutionCaller, error) {
	contract, err := bindIExecution(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionCaller{contract: contract}, nil
}

// NewIExecutionTransactor creates a new write-only instance of IExecution, bound to a specific deployed contract.
func NewIExecutionTransactor(address common.Address, transactor bind.ContractTransactor) (*IExecutionTransactor, error) {
	contract, err := bindIExecution(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionTransactor{contract: contract}, nil
}

// NewIExecutionFilterer creates a new log filterer instance of IExecution, bound to a specific deployed contract.
func NewIExecutionFilterer(address common.Address, filterer bind.ContractFilterer) (*IExecutionFilterer, error) {
	contract, err := bindIExecution(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExecutionFilterer{contract: contract}, nil
}

// bindIExecution binds a generic wrapper to an already deployed contract.
func bindIExecution(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IExecutionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecution *IExecutionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecution.Contract.IExecutionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecution *IExecutionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecution.Contract.IExecutionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecution *IExecutionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecution.Contract.IExecutionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecution *IExecutionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecution.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecution *IExecutionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecution.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecution *IExecutionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecution.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0xe9ae5c53.
//
// Solidity: function execute(bytes32 mode, bytes executionCalldata) returns()
func (_IExecution *IExecutionTransactor) Execute(opts *bind.TransactOpts, mode [32]byte, executionCalldata []byte) (*types.Transaction, error) {
	return _IExecution.contract.Transact(opts, "execute", mode, executionCalldata)
}

// Execute is a paid mutator transaction binding the contract method 0xe9ae5c53.
//
// Solidity: function execute(bytes32 mode, bytes executionCalldata) returns()
func (_IExecution *IExecutionSession) Execute(mode [32]byte, executionCalldata []byte) (*types.Transaction, error) {
	return _IExecution.Contract.Execute(&_IExecution.TransactOpts, mode, executionCalldata)
}

// Execute is a paid mutator transaction binding the contract method 0xe9ae5c53.
//
// Solidity: function execute(bytes32 mode, bytes executionCalldata) returns()
func (_IExecution *IExecutionTransactorSession) Execute(mode [32]byte, executionCalldata []byte) (*types.Transaction, error) {
	return _IExecution.Contract.Execute(&_IExecution.TransactOpts, mode, executionCalldata)
}

// ExecuteFromExecutor is a paid mutator transaction binding the contract method 0xd691c964.
//
// Solidity: function executeFromExecutor(bytes32 mode, bytes executionCalldata) returns(bytes[] returnData)
func (_IExecution *IExecutionTransactor) ExecuteFromExecutor(opts *bind.TransactOpts, mode [32]byte, executionCalldata []byte) (*types.Transaction, error) {
	return _IExecution.contract.Transact(opts, "executeFromExecutor", mode, executionCalldata)
}

// ExecuteFromExecutor is a paid mutator transaction binding the contract method 0xd691c964.
//
// Solidity: function executeFromExecutor(bytes32 mode, bytes executionCalldata) returns(bytes[] returnData)
func (_IExecution *IExecutionSession) ExecuteFromExecutor(mode [32]byte, executionCalldata []byte) (*types.Transaction, error) {
	return _IExecution.Contract.ExecuteFromExecutor(&_IExecution.TransactOpts, mode, executionCalldata)
}

// ExecuteFromExecutor is a paid mutator transaction binding the contract method 0xd691c964.
//
// Solidity: function executeFromExecutor(bytes32 mode, bytes executionCalldata) returns(bytes[] returnData)
func (_IExecution *IExecutionTransactorSession) ExecuteFromExecutor(mode [32]byte, executionCalldata []byte) (*types.Transaction, error) {
	return _IExecution.Contract.ExecuteFromExecutor(&_IExecution.TransactOpts, mode, executionCalldata)
}
