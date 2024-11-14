// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package inoncemanager

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

// INonceManagerMetaData contains all meta data concerning the INonceManager contract.
var INonceManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// INonceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use INonceManagerMetaData.ABI instead.
var INonceManagerABI = INonceManagerMetaData.ABI

// INonceManager is an auto generated Go binding around an Ethereum contract.
type INonceManager struct {
	INonceManagerCaller     // Read-only binding to the contract
	INonceManagerTransactor // Write-only binding to the contract
	INonceManagerFilterer   // Log filterer for contract events
}

// INonceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type INonceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INonceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INonceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INonceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INonceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INonceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INonceManagerSession struct {
	Contract     *INonceManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INonceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INonceManagerCallerSession struct {
	Contract *INonceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// INonceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INonceManagerTransactorSession struct {
	Contract     *INonceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// INonceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type INonceManagerRaw struct {
	Contract *INonceManager // Generic contract binding to access the raw methods on
}

// INonceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INonceManagerCallerRaw struct {
	Contract *INonceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// INonceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INonceManagerTransactorRaw struct {
	Contract *INonceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINonceManager creates a new instance of INonceManager, bound to a specific deployed contract.
func NewINonceManager(address common.Address, backend bind.ContractBackend) (*INonceManager, error) {
	contract, err := bindINonceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INonceManager{INonceManagerCaller: INonceManagerCaller{contract: contract}, INonceManagerTransactor: INonceManagerTransactor{contract: contract}, INonceManagerFilterer: INonceManagerFilterer{contract: contract}}, nil
}

// NewINonceManagerCaller creates a new read-only instance of INonceManager, bound to a specific deployed contract.
func NewINonceManagerCaller(address common.Address, caller bind.ContractCaller) (*INonceManagerCaller, error) {
	contract, err := bindINonceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INonceManagerCaller{contract: contract}, nil
}

// NewINonceManagerTransactor creates a new write-only instance of INonceManager, bound to a specific deployed contract.
func NewINonceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*INonceManagerTransactor, error) {
	contract, err := bindINonceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INonceManagerTransactor{contract: contract}, nil
}

// NewINonceManagerFilterer creates a new log filterer instance of INonceManager, bound to a specific deployed contract.
func NewINonceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*INonceManagerFilterer, error) {
	contract, err := bindINonceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INonceManagerFilterer{contract: contract}, nil
}

// bindINonceManager binds a generic wrapper to an already deployed contract.
func bindINonceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := INonceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INonceManager *INonceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INonceManager.Contract.INonceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INonceManager *INonceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INonceManager.Contract.INonceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INonceManager *INonceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INonceManager.Contract.INonceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INonceManager *INonceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INonceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INonceManager *INonceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INonceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INonceManager *INonceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INonceManager.Contract.contract.Transact(opts, method, params...)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_INonceManager *INonceManagerCaller) GetNonce(opts *bind.CallOpts, sender common.Address, key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _INonceManager.contract.Call(opts, &out, "getNonce", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_INonceManager *INonceManagerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _INonceManager.Contract.GetNonce(&_INonceManager.CallOpts, sender, key)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_INonceManager *INonceManagerCallerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _INonceManager.Contract.GetNonce(&_INonceManager.CallOpts, sender, key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_INonceManager *INonceManagerTransactor) IncrementNonce(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _INonceManager.contract.Transact(opts, "incrementNonce", key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_INonceManager *INonceManagerSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _INonceManager.Contract.IncrementNonce(&_INonceManager.TransactOpts, key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_INonceManager *INonceManagerTransactorSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _INonceManager.Contract.IncrementNonce(&_INonceManager.TransactOpts, key)
}
