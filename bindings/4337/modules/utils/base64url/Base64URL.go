// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package base64url

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

// Base64URLMetaData contains all meta data concerning the Base64URL contract.
var Base64URLMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220e3cc9a626d3baf79ff04f773997d88d7f5a19a512a19e3c49ddd0d03869e84cb64736f6c63430008190033",
}

// Base64URLABI is the input ABI used to generate the binding from.
// Deprecated: Use Base64URLMetaData.ABI instead.
var Base64URLABI = Base64URLMetaData.ABI

// Base64URLBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Base64URLMetaData.Bin instead.
var Base64URLBin = Base64URLMetaData.Bin

// DeployBase64URL deploys a new Ethereum contract, binding an instance of Base64URL to it.
func DeployBase64URL(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Base64URL, error) {
	parsed, err := Base64URLMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Base64URLBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Base64URL{Base64URLCaller: Base64URLCaller{contract: contract}, Base64URLTransactor: Base64URLTransactor{contract: contract}, Base64URLFilterer: Base64URLFilterer{contract: contract}}, nil
}

// Base64URL is an auto generated Go binding around an Ethereum contract.
type Base64URL struct {
	Base64URLCaller     // Read-only binding to the contract
	Base64URLTransactor // Write-only binding to the contract
	Base64URLFilterer   // Log filterer for contract events
}

// Base64URLCaller is an auto generated read-only Go binding around an Ethereum contract.
type Base64URLCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64URLTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Base64URLTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64URLFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Base64URLFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64URLSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Base64URLSession struct {
	Contract     *Base64URL        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Base64URLCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Base64URLCallerSession struct {
	Contract *Base64URLCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Base64URLTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Base64URLTransactorSession struct {
	Contract     *Base64URLTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Base64URLRaw is an auto generated low-level Go binding around an Ethereum contract.
type Base64URLRaw struct {
	Contract *Base64URL // Generic contract binding to access the raw methods on
}

// Base64URLCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Base64URLCallerRaw struct {
	Contract *Base64URLCaller // Generic read-only contract binding to access the raw methods on
}

// Base64URLTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Base64URLTransactorRaw struct {
	Contract *Base64URLTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBase64URL creates a new instance of Base64URL, bound to a specific deployed contract.
func NewBase64URL(address common.Address, backend bind.ContractBackend) (*Base64URL, error) {
	contract, err := bindBase64URL(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Base64URL{Base64URLCaller: Base64URLCaller{contract: contract}, Base64URLTransactor: Base64URLTransactor{contract: contract}, Base64URLFilterer: Base64URLFilterer{contract: contract}}, nil
}

// NewBase64URLCaller creates a new read-only instance of Base64URL, bound to a specific deployed contract.
func NewBase64URLCaller(address common.Address, caller bind.ContractCaller) (*Base64URLCaller, error) {
	contract, err := bindBase64URL(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Base64URLCaller{contract: contract}, nil
}

// NewBase64URLTransactor creates a new write-only instance of Base64URL, bound to a specific deployed contract.
func NewBase64URLTransactor(address common.Address, transactor bind.ContractTransactor) (*Base64URLTransactor, error) {
	contract, err := bindBase64URL(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Base64URLTransactor{contract: contract}, nil
}

// NewBase64URLFilterer creates a new log filterer instance of Base64URL, bound to a specific deployed contract.
func NewBase64URLFilterer(address common.Address, filterer bind.ContractFilterer) (*Base64URLFilterer, error) {
	contract, err := bindBase64URL(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Base64URLFilterer{contract: contract}, nil
}

// bindBase64URL binds a generic wrapper to an already deployed contract.
func bindBase64URL(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Base64URLMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base64URL *Base64URLRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Base64URL.Contract.Base64URLCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base64URL *Base64URLRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base64URL.Contract.Base64URLTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base64URL *Base64URLRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base64URL.Contract.Base64URLTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base64URL *Base64URLCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Base64URL.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base64URL *Base64URLTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base64URL.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base64URL *Base64URLTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base64URL.Contract.contract.Transact(opts, method, params...)
}
