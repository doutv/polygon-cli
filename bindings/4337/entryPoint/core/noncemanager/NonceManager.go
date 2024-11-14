// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package noncemanager

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

// NonceManagerMetaData contains all meta data concerning the NonceManager contract.
var NonceManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"name\":\"nonceSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NonceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NonceManagerMetaData.ABI instead.
var NonceManagerABI = NonceManagerMetaData.ABI

// NonceManager is an auto generated Go binding around an Ethereum contract.
type NonceManager struct {
	NonceManagerCaller     // Read-only binding to the contract
	NonceManagerTransactor // Write-only binding to the contract
	NonceManagerFilterer   // Log filterer for contract events
}

// NonceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NonceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NonceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NonceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NonceManagerSession struct {
	Contract     *NonceManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NonceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NonceManagerCallerSession struct {
	Contract *NonceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NonceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NonceManagerTransactorSession struct {
	Contract     *NonceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NonceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NonceManagerRaw struct {
	Contract *NonceManager // Generic contract binding to access the raw methods on
}

// NonceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NonceManagerCallerRaw struct {
	Contract *NonceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NonceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NonceManagerTransactorRaw struct {
	Contract *NonceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNonceManager creates a new instance of NonceManager, bound to a specific deployed contract.
func NewNonceManager(address common.Address, backend bind.ContractBackend) (*NonceManager, error) {
	contract, err := bindNonceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NonceManager{NonceManagerCaller: NonceManagerCaller{contract: contract}, NonceManagerTransactor: NonceManagerTransactor{contract: contract}, NonceManagerFilterer: NonceManagerFilterer{contract: contract}}, nil
}

// NewNonceManagerCaller creates a new read-only instance of NonceManager, bound to a specific deployed contract.
func NewNonceManagerCaller(address common.Address, caller bind.ContractCaller) (*NonceManagerCaller, error) {
	contract, err := bindNonceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NonceManagerCaller{contract: contract}, nil
}

// NewNonceManagerTransactor creates a new write-only instance of NonceManager, bound to a specific deployed contract.
func NewNonceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NonceManagerTransactor, error) {
	contract, err := bindNonceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NonceManagerTransactor{contract: contract}, nil
}

// NewNonceManagerFilterer creates a new log filterer instance of NonceManager, bound to a specific deployed contract.
func NewNonceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NonceManagerFilterer, error) {
	contract, err := bindNonceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NonceManagerFilterer{contract: contract}, nil
}

// bindNonceManager binds a generic wrapper to an already deployed contract.
func bindNonceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NonceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonceManager *NonceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonceManager.Contract.NonceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonceManager *NonceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonceManager.Contract.NonceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonceManager *NonceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonceManager.Contract.NonceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonceManager *NonceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonceManager *NonceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonceManager *NonceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonceManager.Contract.contract.Transact(opts, method, params...)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_NonceManager *NonceManagerCaller) GetNonce(opts *bind.CallOpts, sender common.Address, key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NonceManager.contract.Call(opts, &out, "getNonce", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_NonceManager *NonceManagerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _NonceManager.Contract.GetNonce(&_NonceManager.CallOpts, sender, key)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_NonceManager *NonceManagerCallerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _NonceManager.Contract.GetNonce(&_NonceManager.CallOpts, sender, key)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_NonceManager *NonceManagerCaller) NonceSequenceNumber(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NonceManager.contract.Call(opts, &out, "nonceSequenceNumber", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_NonceManager *NonceManagerSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NonceManager.Contract.NonceSequenceNumber(&_NonceManager.CallOpts, arg0, arg1)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_NonceManager *NonceManagerCallerSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NonceManager.Contract.NonceSequenceNumber(&_NonceManager.CallOpts, arg0, arg1)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_NonceManager *NonceManagerTransactor) IncrementNonce(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _NonceManager.contract.Transact(opts, "incrementNonce", key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_NonceManager *NonceManagerSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _NonceManager.Contract.IncrementNonce(&_NonceManager.TransactOpts, key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_NonceManager *NonceManagerTransactorSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _NonceManager.Contract.IncrementNonce(&_NonceManager.TransactOpts, key)
}
