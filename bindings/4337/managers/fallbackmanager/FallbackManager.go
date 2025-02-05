// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fallbackmanager

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

// FallbackManagerMetaData contains all meta data concerning the FallbackManager contract.
var FallbackManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"fallbackHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608080604052346013576082908160198239f35b600080fdfe6080806040526004361015601257600080fd5b600090813560e01c63eed2f25214602857600080fd5b346048578160031936011260485790546001600160a01b03168152602090f35b5080fdfea2646970667358221220940d714b19c384c71387ffbfef20969e3a98f80388e4b351c498d8e346780fe864736f6c63430008190033",
}

// FallbackManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use FallbackManagerMetaData.ABI instead.
var FallbackManagerABI = FallbackManagerMetaData.ABI

// FallbackManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FallbackManagerMetaData.Bin instead.
var FallbackManagerBin = FallbackManagerMetaData.Bin

// DeployFallbackManager deploys a new Ethereum contract, binding an instance of FallbackManager to it.
func DeployFallbackManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FallbackManager, error) {
	parsed, err := FallbackManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FallbackManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FallbackManager{FallbackManagerCaller: FallbackManagerCaller{contract: contract}, FallbackManagerTransactor: FallbackManagerTransactor{contract: contract}, FallbackManagerFilterer: FallbackManagerFilterer{contract: contract}}, nil
}

// FallbackManager is an auto generated Go binding around an Ethereum contract.
type FallbackManager struct {
	FallbackManagerCaller     // Read-only binding to the contract
	FallbackManagerTransactor // Write-only binding to the contract
	FallbackManagerFilterer   // Log filterer for contract events
}

// FallbackManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FallbackManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FallbackManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FallbackManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FallbackManagerSession struct {
	Contract     *FallbackManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FallbackManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FallbackManagerCallerSession struct {
	Contract *FallbackManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FallbackManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FallbackManagerTransactorSession struct {
	Contract     *FallbackManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FallbackManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FallbackManagerRaw struct {
	Contract *FallbackManager // Generic contract binding to access the raw methods on
}

// FallbackManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FallbackManagerCallerRaw struct {
	Contract *FallbackManagerCaller // Generic read-only contract binding to access the raw methods on
}

// FallbackManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FallbackManagerTransactorRaw struct {
	Contract *FallbackManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFallbackManager creates a new instance of FallbackManager, bound to a specific deployed contract.
func NewFallbackManager(address common.Address, backend bind.ContractBackend) (*FallbackManager, error) {
	contract, err := bindFallbackManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FallbackManager{FallbackManagerCaller: FallbackManagerCaller{contract: contract}, FallbackManagerTransactor: FallbackManagerTransactor{contract: contract}, FallbackManagerFilterer: FallbackManagerFilterer{contract: contract}}, nil
}

// NewFallbackManagerCaller creates a new read-only instance of FallbackManager, bound to a specific deployed contract.
func NewFallbackManagerCaller(address common.Address, caller bind.ContractCaller) (*FallbackManagerCaller, error) {
	contract, err := bindFallbackManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FallbackManagerCaller{contract: contract}, nil
}

// NewFallbackManagerTransactor creates a new write-only instance of FallbackManager, bound to a specific deployed contract.
func NewFallbackManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FallbackManagerTransactor, error) {
	contract, err := bindFallbackManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FallbackManagerTransactor{contract: contract}, nil
}

// NewFallbackManagerFilterer creates a new log filterer instance of FallbackManager, bound to a specific deployed contract.
func NewFallbackManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FallbackManagerFilterer, error) {
	contract, err := bindFallbackManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FallbackManagerFilterer{contract: contract}, nil
}

// bindFallbackManager binds a generic wrapper to an already deployed contract.
func bindFallbackManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FallbackManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FallbackManager *FallbackManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FallbackManager.Contract.FallbackManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FallbackManager *FallbackManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FallbackManager.Contract.FallbackManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FallbackManager *FallbackManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FallbackManager.Contract.FallbackManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FallbackManager *FallbackManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FallbackManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FallbackManager *FallbackManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FallbackManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FallbackManager *FallbackManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FallbackManager.Contract.contract.Transact(opts, method, params...)
}

// FallbackHandler is a free data retrieval call binding the contract method 0xeed2f252.
//
// Solidity: function fallbackHandler() view returns(address)
func (_FallbackManager *FallbackManagerCaller) FallbackHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FallbackManager.contract.Call(opts, &out, "fallbackHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FallbackHandler is a free data retrieval call binding the contract method 0xeed2f252.
//
// Solidity: function fallbackHandler() view returns(address)
func (_FallbackManager *FallbackManagerSession) FallbackHandler() (common.Address, error) {
	return _FallbackManager.Contract.FallbackHandler(&_FallbackManager.CallOpts)
}

// FallbackHandler is a free data retrieval call binding the contract method 0xeed2f252.
//
// Solidity: function fallbackHandler() view returns(address)
func (_FallbackManager *FallbackManagerCallerSession) FallbackHandler() (common.Address, error) {
	return _FallbackManager.Contract.FallbackHandler(&_FallbackManager.CallOpts)
}
