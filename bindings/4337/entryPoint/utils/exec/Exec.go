// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exec

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

// ExecMetaData contains all meta data concerning the Exec contract.
var ExecMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212205c667d2c2e9614ea0e0239c9454f82d0bc51fa224afe13d4851073fd718bf26c64736f6c63430008190033",
}

// ExecABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecMetaData.ABI instead.
var ExecABI = ExecMetaData.ABI

// ExecBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExecMetaData.Bin instead.
var ExecBin = ExecMetaData.Bin

// DeployExec deploys a new Ethereum contract, binding an instance of Exec to it.
func DeployExec(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Exec, error) {
	parsed, err := ExecMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Exec{ExecCaller: ExecCaller{contract: contract}, ExecTransactor: ExecTransactor{contract: contract}, ExecFilterer: ExecFilterer{contract: contract}}, nil
}

// Exec is an auto generated Go binding around an Ethereum contract.
type Exec struct {
	ExecCaller     // Read-only binding to the contract
	ExecTransactor // Write-only binding to the contract
	ExecFilterer   // Log filterer for contract events
}

// ExecCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecSession struct {
	Contract     *Exec             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecCallerSession struct {
	Contract *ExecCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ExecTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecTransactorSession struct {
	Contract     *ExecTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecRaw struct {
	Contract *Exec // Generic contract binding to access the raw methods on
}

// ExecCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecCallerRaw struct {
	Contract *ExecCaller // Generic read-only contract binding to access the raw methods on
}

// ExecTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecTransactorRaw struct {
	Contract *ExecTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExec creates a new instance of Exec, bound to a specific deployed contract.
func NewExec(address common.Address, backend bind.ContractBackend) (*Exec, error) {
	contract, err := bindExec(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exec{ExecCaller: ExecCaller{contract: contract}, ExecTransactor: ExecTransactor{contract: contract}, ExecFilterer: ExecFilterer{contract: contract}}, nil
}

// NewExecCaller creates a new read-only instance of Exec, bound to a specific deployed contract.
func NewExecCaller(address common.Address, caller bind.ContractCaller) (*ExecCaller, error) {
	contract, err := bindExec(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecCaller{contract: contract}, nil
}

// NewExecTransactor creates a new write-only instance of Exec, bound to a specific deployed contract.
func NewExecTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecTransactor, error) {
	contract, err := bindExec(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecTransactor{contract: contract}, nil
}

// NewExecFilterer creates a new log filterer instance of Exec, bound to a specific deployed contract.
func NewExecFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecFilterer, error) {
	contract, err := bindExec(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecFilterer{contract: contract}, nil
}

// bindExec binds a generic wrapper to an already deployed contract.
func bindExec(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExecMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exec *ExecRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exec.Contract.ExecCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exec *ExecRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exec.Contract.ExecTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exec *ExecRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exec.Contract.ExecTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exec *ExecCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exec.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exec *ExecTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exec.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exec *ExecTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exec.Contract.contract.Transact(opts, method, params...)
}
