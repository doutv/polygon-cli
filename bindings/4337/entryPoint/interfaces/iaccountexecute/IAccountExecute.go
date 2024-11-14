// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaccountexecute

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

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	Sender             common.Address
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

// IAccountExecuteMetaData contains all meta data concerning the IAccountExecute contract.
var IAccountExecuteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"name\":\"executeUserOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAccountExecuteABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccountExecuteMetaData.ABI instead.
var IAccountExecuteABI = IAccountExecuteMetaData.ABI

// IAccountExecute is an auto generated Go binding around an Ethereum contract.
type IAccountExecute struct {
	IAccountExecuteCaller     // Read-only binding to the contract
	IAccountExecuteTransactor // Write-only binding to the contract
	IAccountExecuteFilterer   // Log filterer for contract events
}

// IAccountExecuteCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccountExecuteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountExecuteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccountExecuteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountExecuteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccountExecuteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountExecuteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccountExecuteSession struct {
	Contract     *IAccountExecute  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccountExecuteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccountExecuteCallerSession struct {
	Contract *IAccountExecuteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IAccountExecuteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccountExecuteTransactorSession struct {
	Contract     *IAccountExecuteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IAccountExecuteRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccountExecuteRaw struct {
	Contract *IAccountExecute // Generic contract binding to access the raw methods on
}

// IAccountExecuteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccountExecuteCallerRaw struct {
	Contract *IAccountExecuteCaller // Generic read-only contract binding to access the raw methods on
}

// IAccountExecuteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccountExecuteTransactorRaw struct {
	Contract *IAccountExecuteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccountExecute creates a new instance of IAccountExecute, bound to a specific deployed contract.
func NewIAccountExecute(address common.Address, backend bind.ContractBackend) (*IAccountExecute, error) {
	contract, err := bindIAccountExecute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccountExecute{IAccountExecuteCaller: IAccountExecuteCaller{contract: contract}, IAccountExecuteTransactor: IAccountExecuteTransactor{contract: contract}, IAccountExecuteFilterer: IAccountExecuteFilterer{contract: contract}}, nil
}

// NewIAccountExecuteCaller creates a new read-only instance of IAccountExecute, bound to a specific deployed contract.
func NewIAccountExecuteCaller(address common.Address, caller bind.ContractCaller) (*IAccountExecuteCaller, error) {
	contract, err := bindIAccountExecute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountExecuteCaller{contract: contract}, nil
}

// NewIAccountExecuteTransactor creates a new write-only instance of IAccountExecute, bound to a specific deployed contract.
func NewIAccountExecuteTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccountExecuteTransactor, error) {
	contract, err := bindIAccountExecute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountExecuteTransactor{contract: contract}, nil
}

// NewIAccountExecuteFilterer creates a new log filterer instance of IAccountExecute, bound to a specific deployed contract.
func NewIAccountExecuteFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccountExecuteFilterer, error) {
	contract, err := bindIAccountExecute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccountExecuteFilterer{contract: contract}, nil
}

// bindIAccountExecute binds a generic wrapper to an already deployed contract.
func bindIAccountExecute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccountExecuteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountExecute *IAccountExecuteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountExecute.Contract.IAccountExecuteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountExecute *IAccountExecuteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountExecute.Contract.IAccountExecuteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountExecute *IAccountExecuteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountExecute.Contract.IAccountExecuteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountExecute *IAccountExecuteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountExecute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountExecute *IAccountExecuteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountExecute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountExecute *IAccountExecuteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountExecute.Contract.contract.Transact(opts, method, params...)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x8dd7712f.
//
// Solidity: function executeUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash) returns()
func (_IAccountExecute *IAccountExecuteTransactor) ExecuteUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte) (*types.Transaction, error) {
	return _IAccountExecute.contract.Transact(opts, "executeUserOp", userOp, userOpHash)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x8dd7712f.
//
// Solidity: function executeUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash) returns()
func (_IAccountExecute *IAccountExecuteSession) ExecuteUserOp(userOp PackedUserOperation, userOpHash [32]byte) (*types.Transaction, error) {
	return _IAccountExecute.Contract.ExecuteUserOp(&_IAccountExecute.TransactOpts, userOp, userOpHash)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x8dd7712f.
//
// Solidity: function executeUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash) returns()
func (_IAccountExecute *IAccountExecuteTransactorSession) ExecuteUserOp(userOp PackedUserOperation, userOpHash [32]byte) (*types.Transaction, error) {
	return _IAccountExecute.Contract.ExecuteUserOp(&_IAccountExecute.TransactOpts, userOp, userOpHash)
}
