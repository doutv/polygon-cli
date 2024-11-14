// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package decodelib

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

// DecodeLibMetaData contains all meta data concerning the DecodeLib contract.
var DecodeLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220e0ff84e017c2e83732d7b3c2bd0208e22c94d7aab8283e503a32656cbbee5ae964736f6c63430008190033",
}

// DecodeLibABI is the input ABI used to generate the binding from.
// Deprecated: Use DecodeLibMetaData.ABI instead.
var DecodeLibABI = DecodeLibMetaData.ABI

// DecodeLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DecodeLibMetaData.Bin instead.
var DecodeLibBin = DecodeLibMetaData.Bin

// DeployDecodeLib deploys a new Ethereum contract, binding an instance of DecodeLib to it.
func DeployDecodeLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DecodeLib, error) {
	parsed, err := DecodeLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DecodeLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DecodeLib{DecodeLibCaller: DecodeLibCaller{contract: contract}, DecodeLibTransactor: DecodeLibTransactor{contract: contract}, DecodeLibFilterer: DecodeLibFilterer{contract: contract}}, nil
}

// DecodeLib is an auto generated Go binding around an Ethereum contract.
type DecodeLib struct {
	DecodeLibCaller     // Read-only binding to the contract
	DecodeLibTransactor // Write-only binding to the contract
	DecodeLibFilterer   // Log filterer for contract events
}

// DecodeLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type DecodeLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecodeLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DecodeLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecodeLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DecodeLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecodeLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DecodeLibSession struct {
	Contract     *DecodeLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DecodeLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DecodeLibCallerSession struct {
	Contract *DecodeLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DecodeLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DecodeLibTransactorSession struct {
	Contract     *DecodeLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DecodeLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type DecodeLibRaw struct {
	Contract *DecodeLib // Generic contract binding to access the raw methods on
}

// DecodeLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DecodeLibCallerRaw struct {
	Contract *DecodeLibCaller // Generic read-only contract binding to access the raw methods on
}

// DecodeLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DecodeLibTransactorRaw struct {
	Contract *DecodeLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDecodeLib creates a new instance of DecodeLib, bound to a specific deployed contract.
func NewDecodeLib(address common.Address, backend bind.ContractBackend) (*DecodeLib, error) {
	contract, err := bindDecodeLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DecodeLib{DecodeLibCaller: DecodeLibCaller{contract: contract}, DecodeLibTransactor: DecodeLibTransactor{contract: contract}, DecodeLibFilterer: DecodeLibFilterer{contract: contract}}, nil
}

// NewDecodeLibCaller creates a new read-only instance of DecodeLib, bound to a specific deployed contract.
func NewDecodeLibCaller(address common.Address, caller bind.ContractCaller) (*DecodeLibCaller, error) {
	contract, err := bindDecodeLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DecodeLibCaller{contract: contract}, nil
}

// NewDecodeLibTransactor creates a new write-only instance of DecodeLib, bound to a specific deployed contract.
func NewDecodeLibTransactor(address common.Address, transactor bind.ContractTransactor) (*DecodeLibTransactor, error) {
	contract, err := bindDecodeLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DecodeLibTransactor{contract: contract}, nil
}

// NewDecodeLibFilterer creates a new log filterer instance of DecodeLib, bound to a specific deployed contract.
func NewDecodeLibFilterer(address common.Address, filterer bind.ContractFilterer) (*DecodeLibFilterer, error) {
	contract, err := bindDecodeLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DecodeLibFilterer{contract: contract}, nil
}

// bindDecodeLib binds a generic wrapper to an already deployed contract.
func bindDecodeLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DecodeLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DecodeLib *DecodeLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DecodeLib.Contract.DecodeLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DecodeLib *DecodeLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecodeLib.Contract.DecodeLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DecodeLib *DecodeLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DecodeLib.Contract.DecodeLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DecodeLib *DecodeLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DecodeLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DecodeLib *DecodeLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecodeLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DecodeLib *DecodeLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DecodeLib.Contract.contract.Transact(opts, method, params...)
}
