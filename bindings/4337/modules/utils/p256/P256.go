// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package p256

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

// P256MetaData contains all meta data concerning the P256 contract.
var P256MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212208f79c1343a449ed5aebd6d684f8f21cfe5788d51f0d4e4c2dd5216870fa38d8f64736f6c63430008190033",
}

// P256ABI is the input ABI used to generate the binding from.
// Deprecated: Use P256MetaData.ABI instead.
var P256ABI = P256MetaData.ABI

// P256Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use P256MetaData.Bin instead.
var P256Bin = P256MetaData.Bin

// DeployP256 deploys a new Ethereum contract, binding an instance of P256 to it.
func DeployP256(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *P256, error) {
	parsed, err := P256MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(P256Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &P256{P256Caller: P256Caller{contract: contract}, P256Transactor: P256Transactor{contract: contract}, P256Filterer: P256Filterer{contract: contract}}, nil
}

// P256 is an auto generated Go binding around an Ethereum contract.
type P256 struct {
	P256Caller     // Read-only binding to the contract
	P256Transactor // Write-only binding to the contract
	P256Filterer   // Log filterer for contract events
}

// P256Caller is an auto generated read-only Go binding around an Ethereum contract.
type P256Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// P256Transactor is an auto generated write-only Go binding around an Ethereum contract.
type P256Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// P256Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type P256Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// P256Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type P256Session struct {
	Contract     *P256             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// P256CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type P256CallerSession struct {
	Contract *P256Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// P256TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type P256TransactorSession struct {
	Contract     *P256Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// P256Raw is an auto generated low-level Go binding around an Ethereum contract.
type P256Raw struct {
	Contract *P256 // Generic contract binding to access the raw methods on
}

// P256CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type P256CallerRaw struct {
	Contract *P256Caller // Generic read-only contract binding to access the raw methods on
}

// P256TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type P256TransactorRaw struct {
	Contract *P256Transactor // Generic write-only contract binding to access the raw methods on
}

// NewP256 creates a new instance of P256, bound to a specific deployed contract.
func NewP256(address common.Address, backend bind.ContractBackend) (*P256, error) {
	contract, err := bindP256(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &P256{P256Caller: P256Caller{contract: contract}, P256Transactor: P256Transactor{contract: contract}, P256Filterer: P256Filterer{contract: contract}}, nil
}

// NewP256Caller creates a new read-only instance of P256, bound to a specific deployed contract.
func NewP256Caller(address common.Address, caller bind.ContractCaller) (*P256Caller, error) {
	contract, err := bindP256(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &P256Caller{contract: contract}, nil
}

// NewP256Transactor creates a new write-only instance of P256, bound to a specific deployed contract.
func NewP256Transactor(address common.Address, transactor bind.ContractTransactor) (*P256Transactor, error) {
	contract, err := bindP256(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &P256Transactor{contract: contract}, nil
}

// NewP256Filterer creates a new log filterer instance of P256, bound to a specific deployed contract.
func NewP256Filterer(address common.Address, filterer bind.ContractFilterer) (*P256Filterer, error) {
	contract, err := bindP256(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &P256Filterer{contract: contract}, nil
}

// bindP256 binds a generic wrapper to an already deployed contract.
func bindP256(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := P256MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_P256 *P256Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _P256.Contract.P256Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_P256 *P256Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _P256.Contract.P256Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_P256 *P256Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _P256.Contract.P256Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_P256 *P256CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _P256.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_P256 *P256TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _P256.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_P256 *P256TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _P256.Contract.contract.Transact(opts, method, params...)
}
