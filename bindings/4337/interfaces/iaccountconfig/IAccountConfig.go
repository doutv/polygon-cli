// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaccountconfig

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

// IAccountConfigMetaData contains all meta data concerning the IAccountConfig contract.
var IAccountConfigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"accountId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"accountImplementationId\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IAccountConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccountConfigMetaData.ABI instead.
var IAccountConfigABI = IAccountConfigMetaData.ABI

// IAccountConfig is an auto generated Go binding around an Ethereum contract.
type IAccountConfig struct {
	IAccountConfigCaller     // Read-only binding to the contract
	IAccountConfigTransactor // Write-only binding to the contract
	IAccountConfigFilterer   // Log filterer for contract events
}

// IAccountConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccountConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccountConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccountConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccountConfigSession struct {
	Contract     *IAccountConfig   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccountConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccountConfigCallerSession struct {
	Contract *IAccountConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IAccountConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccountConfigTransactorSession struct {
	Contract     *IAccountConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IAccountConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccountConfigRaw struct {
	Contract *IAccountConfig // Generic contract binding to access the raw methods on
}

// IAccountConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccountConfigCallerRaw struct {
	Contract *IAccountConfigCaller // Generic read-only contract binding to access the raw methods on
}

// IAccountConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccountConfigTransactorRaw struct {
	Contract *IAccountConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccountConfig creates a new instance of IAccountConfig, bound to a specific deployed contract.
func NewIAccountConfig(address common.Address, backend bind.ContractBackend) (*IAccountConfig, error) {
	contract, err := bindIAccountConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccountConfig{IAccountConfigCaller: IAccountConfigCaller{contract: contract}, IAccountConfigTransactor: IAccountConfigTransactor{contract: contract}, IAccountConfigFilterer: IAccountConfigFilterer{contract: contract}}, nil
}

// NewIAccountConfigCaller creates a new read-only instance of IAccountConfig, bound to a specific deployed contract.
func NewIAccountConfigCaller(address common.Address, caller bind.ContractCaller) (*IAccountConfigCaller, error) {
	contract, err := bindIAccountConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountConfigCaller{contract: contract}, nil
}

// NewIAccountConfigTransactor creates a new write-only instance of IAccountConfig, bound to a specific deployed contract.
func NewIAccountConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccountConfigTransactor, error) {
	contract, err := bindIAccountConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountConfigTransactor{contract: contract}, nil
}

// NewIAccountConfigFilterer creates a new log filterer instance of IAccountConfig, bound to a specific deployed contract.
func NewIAccountConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccountConfigFilterer, error) {
	contract, err := bindIAccountConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccountConfigFilterer{contract: contract}, nil
}

// bindIAccountConfig binds a generic wrapper to an already deployed contract.
func bindIAccountConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccountConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountConfig *IAccountConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountConfig.Contract.IAccountConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountConfig *IAccountConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountConfig.Contract.IAccountConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountConfig *IAccountConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountConfig.Contract.IAccountConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountConfig *IAccountConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountConfig *IAccountConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountConfig *IAccountConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountConfig.Contract.contract.Transact(opts, method, params...)
}

// AccountId is a free data retrieval call binding the contract method 0x9cfd7cff.
//
// Solidity: function accountId() view returns(string accountImplementationId)
func (_IAccountConfig *IAccountConfigCaller) AccountId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IAccountConfig.contract.Call(opts, &out, "accountId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AccountId is a free data retrieval call binding the contract method 0x9cfd7cff.
//
// Solidity: function accountId() view returns(string accountImplementationId)
func (_IAccountConfig *IAccountConfigSession) AccountId() (string, error) {
	return _IAccountConfig.Contract.AccountId(&_IAccountConfig.CallOpts)
}

// AccountId is a free data retrieval call binding the contract method 0x9cfd7cff.
//
// Solidity: function accountId() view returns(string accountImplementationId)
func (_IAccountConfig *IAccountConfigCallerSession) AccountId() (string, error) {
	return _IAccountConfig.Contract.AccountId(&_IAccountConfig.CallOpts)
}
