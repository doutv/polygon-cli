// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package modexpsqrt

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

// ModexpSqrtMetaData contains all meta data concerning the ModexpSqrt contract.
var ModexpSqrtMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220416c49f418825a73f2d252e90eae538437303e35fd44bc7f2376c3e2ef2a1e8464736f6c63430008190033",
}

// ModexpSqrtABI is the input ABI used to generate the binding from.
// Deprecated: Use ModexpSqrtMetaData.ABI instead.
var ModexpSqrtABI = ModexpSqrtMetaData.ABI

// ModexpSqrtBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ModexpSqrtMetaData.Bin instead.
var ModexpSqrtBin = ModexpSqrtMetaData.Bin

// DeployModexpSqrt deploys a new Ethereum contract, binding an instance of ModexpSqrt to it.
func DeployModexpSqrt(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ModexpSqrt, error) {
	parsed, err := ModexpSqrtMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ModexpSqrtBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ModexpSqrt{ModexpSqrtCaller: ModexpSqrtCaller{contract: contract}, ModexpSqrtTransactor: ModexpSqrtTransactor{contract: contract}, ModexpSqrtFilterer: ModexpSqrtFilterer{contract: contract}}, nil
}

// ModexpSqrt is an auto generated Go binding around an Ethereum contract.
type ModexpSqrt struct {
	ModexpSqrtCaller     // Read-only binding to the contract
	ModexpSqrtTransactor // Write-only binding to the contract
	ModexpSqrtFilterer   // Log filterer for contract events
}

// ModexpSqrtCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModexpSqrtCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModexpSqrtTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModexpSqrtTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModexpSqrtFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModexpSqrtFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModexpSqrtSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModexpSqrtSession struct {
	Contract     *ModexpSqrt       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModexpSqrtCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModexpSqrtCallerSession struct {
	Contract *ModexpSqrtCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ModexpSqrtTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModexpSqrtTransactorSession struct {
	Contract     *ModexpSqrtTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ModexpSqrtRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModexpSqrtRaw struct {
	Contract *ModexpSqrt // Generic contract binding to access the raw methods on
}

// ModexpSqrtCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModexpSqrtCallerRaw struct {
	Contract *ModexpSqrtCaller // Generic read-only contract binding to access the raw methods on
}

// ModexpSqrtTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModexpSqrtTransactorRaw struct {
	Contract *ModexpSqrtTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModexpSqrt creates a new instance of ModexpSqrt, bound to a specific deployed contract.
func NewModexpSqrt(address common.Address, backend bind.ContractBackend) (*ModexpSqrt, error) {
	contract, err := bindModexpSqrt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ModexpSqrt{ModexpSqrtCaller: ModexpSqrtCaller{contract: contract}, ModexpSqrtTransactor: ModexpSqrtTransactor{contract: contract}, ModexpSqrtFilterer: ModexpSqrtFilterer{contract: contract}}, nil
}

// NewModexpSqrtCaller creates a new read-only instance of ModexpSqrt, bound to a specific deployed contract.
func NewModexpSqrtCaller(address common.Address, caller bind.ContractCaller) (*ModexpSqrtCaller, error) {
	contract, err := bindModexpSqrt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModexpSqrtCaller{contract: contract}, nil
}

// NewModexpSqrtTransactor creates a new write-only instance of ModexpSqrt, bound to a specific deployed contract.
func NewModexpSqrtTransactor(address common.Address, transactor bind.ContractTransactor) (*ModexpSqrtTransactor, error) {
	contract, err := bindModexpSqrt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModexpSqrtTransactor{contract: contract}, nil
}

// NewModexpSqrtFilterer creates a new log filterer instance of ModexpSqrt, bound to a specific deployed contract.
func NewModexpSqrtFilterer(address common.Address, filterer bind.ContractFilterer) (*ModexpSqrtFilterer, error) {
	contract, err := bindModexpSqrt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModexpSqrtFilterer{contract: contract}, nil
}

// bindModexpSqrt binds a generic wrapper to an already deployed contract.
func bindModexpSqrt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ModexpSqrtMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModexpSqrt *ModexpSqrtRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModexpSqrt.Contract.ModexpSqrtCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModexpSqrt *ModexpSqrtRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModexpSqrt.Contract.ModexpSqrtTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModexpSqrt *ModexpSqrtRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModexpSqrt.Contract.ModexpSqrtTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModexpSqrt *ModexpSqrtCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModexpSqrt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModexpSqrt *ModexpSqrtTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModexpSqrt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModexpSqrt *ModexpSqrtTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModexpSqrt.Contract.contract.Transact(opts, method, params...)
}
