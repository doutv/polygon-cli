// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaccount

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

// IAccountMetaData contains all meta data concerning the IAccount contract.
var IAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccountMetaData.ABI instead.
var IAccountABI = IAccountMetaData.ABI

// IAccount is an auto generated Go binding around an Ethereum contract.
type IAccount struct {
	IAccountCaller     // Read-only binding to the contract
	IAccountTransactor // Write-only binding to the contract
	IAccountFilterer   // Log filterer for contract events
}

// IAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccountSession struct {
	Contract     *IAccount         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccountCallerSession struct {
	Contract *IAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccountTransactorSession struct {
	Contract     *IAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccountRaw struct {
	Contract *IAccount // Generic contract binding to access the raw methods on
}

// IAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccountCallerRaw struct {
	Contract *IAccountCaller // Generic read-only contract binding to access the raw methods on
}

// IAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccountTransactorRaw struct {
	Contract *IAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccount creates a new instance of IAccount, bound to a specific deployed contract.
func NewIAccount(address common.Address, backend bind.ContractBackend) (*IAccount, error) {
	contract, err := bindIAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccount{IAccountCaller: IAccountCaller{contract: contract}, IAccountTransactor: IAccountTransactor{contract: contract}, IAccountFilterer: IAccountFilterer{contract: contract}}, nil
}

// NewIAccountCaller creates a new read-only instance of IAccount, bound to a specific deployed contract.
func NewIAccountCaller(address common.Address, caller bind.ContractCaller) (*IAccountCaller, error) {
	contract, err := bindIAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountCaller{contract: contract}, nil
}

// NewIAccountTransactor creates a new write-only instance of IAccount, bound to a specific deployed contract.
func NewIAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccountTransactor, error) {
	contract, err := bindIAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountTransactor{contract: contract}, nil
}

// NewIAccountFilterer creates a new log filterer instance of IAccount, bound to a specific deployed contract.
func NewIAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccountFilterer, error) {
	contract, err := bindIAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccountFilterer{contract: contract}, nil
}

// bindIAccount binds a generic wrapper to an already deployed contract.
func bindIAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccount *IAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccount.Contract.IAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccount *IAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccount.Contract.IAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccount *IAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccount.Contract.IAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccount *IAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccount *IAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccount *IAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccount.Contract.contract.Transact(opts, method, params...)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_IAccount *IAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _IAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_IAccount *IAccountSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _IAccount.Contract.ValidateUserOp(&_IAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_IAccount *IAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _IAccount.Contract.ValidateUserOp(&_IAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}
