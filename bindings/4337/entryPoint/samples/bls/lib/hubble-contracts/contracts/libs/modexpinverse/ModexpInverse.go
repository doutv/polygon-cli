// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package modexpinverse

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

// ModexpInverseMetaData contains all meta data concerning the ModexpInverse contract.
var ModexpInverseMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220a89e25e7ff43a1b018f5af06023135634642e35eb0c7bb8a0b4fdf38b5894cc264736f6c63430008190033",
}

// ModexpInverseABI is the input ABI used to generate the binding from.
// Deprecated: Use ModexpInverseMetaData.ABI instead.
var ModexpInverseABI = ModexpInverseMetaData.ABI

// ModexpInverseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ModexpInverseMetaData.Bin instead.
var ModexpInverseBin = ModexpInverseMetaData.Bin

// DeployModexpInverse deploys a new Ethereum contract, binding an instance of ModexpInverse to it.
func DeployModexpInverse(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ModexpInverse, error) {
	parsed, err := ModexpInverseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ModexpInverseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ModexpInverse{ModexpInverseCaller: ModexpInverseCaller{contract: contract}, ModexpInverseTransactor: ModexpInverseTransactor{contract: contract}, ModexpInverseFilterer: ModexpInverseFilterer{contract: contract}}, nil
}

// ModexpInverse is an auto generated Go binding around an Ethereum contract.
type ModexpInverse struct {
	ModexpInverseCaller     // Read-only binding to the contract
	ModexpInverseTransactor // Write-only binding to the contract
	ModexpInverseFilterer   // Log filterer for contract events
}

// ModexpInverseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModexpInverseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModexpInverseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModexpInverseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModexpInverseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModexpInverseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModexpInverseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModexpInverseSession struct {
	Contract     *ModexpInverse    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModexpInverseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModexpInverseCallerSession struct {
	Contract *ModexpInverseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ModexpInverseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModexpInverseTransactorSession struct {
	Contract     *ModexpInverseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ModexpInverseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModexpInverseRaw struct {
	Contract *ModexpInverse // Generic contract binding to access the raw methods on
}

// ModexpInverseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModexpInverseCallerRaw struct {
	Contract *ModexpInverseCaller // Generic read-only contract binding to access the raw methods on
}

// ModexpInverseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModexpInverseTransactorRaw struct {
	Contract *ModexpInverseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModexpInverse creates a new instance of ModexpInverse, bound to a specific deployed contract.
func NewModexpInverse(address common.Address, backend bind.ContractBackend) (*ModexpInverse, error) {
	contract, err := bindModexpInverse(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ModexpInverse{ModexpInverseCaller: ModexpInverseCaller{contract: contract}, ModexpInverseTransactor: ModexpInverseTransactor{contract: contract}, ModexpInverseFilterer: ModexpInverseFilterer{contract: contract}}, nil
}

// NewModexpInverseCaller creates a new read-only instance of ModexpInverse, bound to a specific deployed contract.
func NewModexpInverseCaller(address common.Address, caller bind.ContractCaller) (*ModexpInverseCaller, error) {
	contract, err := bindModexpInverse(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModexpInverseCaller{contract: contract}, nil
}

// NewModexpInverseTransactor creates a new write-only instance of ModexpInverse, bound to a specific deployed contract.
func NewModexpInverseTransactor(address common.Address, transactor bind.ContractTransactor) (*ModexpInverseTransactor, error) {
	contract, err := bindModexpInverse(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModexpInverseTransactor{contract: contract}, nil
}

// NewModexpInverseFilterer creates a new log filterer instance of ModexpInverse, bound to a specific deployed contract.
func NewModexpInverseFilterer(address common.Address, filterer bind.ContractFilterer) (*ModexpInverseFilterer, error) {
	contract, err := bindModexpInverse(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModexpInverseFilterer{contract: contract}, nil
}

// bindModexpInverse binds a generic wrapper to an already deployed contract.
func bindModexpInverse(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ModexpInverseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModexpInverse *ModexpInverseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModexpInverse.Contract.ModexpInverseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModexpInverse *ModexpInverseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModexpInverse.Contract.ModexpInverseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModexpInverse *ModexpInverseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModexpInverse.Contract.ModexpInverseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModexpInverse *ModexpInverseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModexpInverse.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModexpInverse *ModexpInverseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModexpInverse.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModexpInverse *ModexpInverseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModexpInverse.Contract.contract.Transact(opts, method, params...)
}
