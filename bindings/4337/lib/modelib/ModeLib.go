// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package modelib

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

// ModeLibMetaData contains all meta data concerning the ModeLib contract.
var ModeLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220feb6a951ab0658e34be6ebb71a535f305ae13c1021096b240ce715b1ba05fd0664736f6c63430008190033",
}

// ModeLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ModeLibMetaData.ABI instead.
var ModeLibABI = ModeLibMetaData.ABI

// ModeLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ModeLibMetaData.Bin instead.
var ModeLibBin = ModeLibMetaData.Bin

// DeployModeLib deploys a new Ethereum contract, binding an instance of ModeLib to it.
func DeployModeLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ModeLib, error) {
	parsed, err := ModeLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ModeLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ModeLib{ModeLibCaller: ModeLibCaller{contract: contract}, ModeLibTransactor: ModeLibTransactor{contract: contract}, ModeLibFilterer: ModeLibFilterer{contract: contract}}, nil
}

// ModeLib is an auto generated Go binding around an Ethereum contract.
type ModeLib struct {
	ModeLibCaller     // Read-only binding to the contract
	ModeLibTransactor // Write-only binding to the contract
	ModeLibFilterer   // Log filterer for contract events
}

// ModeLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModeLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModeLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModeLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModeLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModeLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModeLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModeLibSession struct {
	Contract     *ModeLib          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModeLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModeLibCallerSession struct {
	Contract *ModeLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ModeLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModeLibTransactorSession struct {
	Contract     *ModeLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ModeLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModeLibRaw struct {
	Contract *ModeLib // Generic contract binding to access the raw methods on
}

// ModeLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModeLibCallerRaw struct {
	Contract *ModeLibCaller // Generic read-only contract binding to access the raw methods on
}

// ModeLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModeLibTransactorRaw struct {
	Contract *ModeLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModeLib creates a new instance of ModeLib, bound to a specific deployed contract.
func NewModeLib(address common.Address, backend bind.ContractBackend) (*ModeLib, error) {
	contract, err := bindModeLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ModeLib{ModeLibCaller: ModeLibCaller{contract: contract}, ModeLibTransactor: ModeLibTransactor{contract: contract}, ModeLibFilterer: ModeLibFilterer{contract: contract}}, nil
}

// NewModeLibCaller creates a new read-only instance of ModeLib, bound to a specific deployed contract.
func NewModeLibCaller(address common.Address, caller bind.ContractCaller) (*ModeLibCaller, error) {
	contract, err := bindModeLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModeLibCaller{contract: contract}, nil
}

// NewModeLibTransactor creates a new write-only instance of ModeLib, bound to a specific deployed contract.
func NewModeLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ModeLibTransactor, error) {
	contract, err := bindModeLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModeLibTransactor{contract: contract}, nil
}

// NewModeLibFilterer creates a new log filterer instance of ModeLib, bound to a specific deployed contract.
func NewModeLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ModeLibFilterer, error) {
	contract, err := bindModeLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModeLibFilterer{contract: contract}, nil
}

// bindModeLib binds a generic wrapper to an already deployed contract.
func bindModeLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ModeLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModeLib *ModeLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModeLib.Contract.ModeLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModeLib *ModeLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModeLib.Contract.ModeLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModeLib *ModeLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModeLib.Contract.ModeLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModeLib *ModeLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModeLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModeLib *ModeLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModeLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModeLib *ModeLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModeLib.Contract.contract.Transact(opts, method, params...)
}
