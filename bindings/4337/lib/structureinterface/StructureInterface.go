// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package structureinterface

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

// StructureInterfaceMetaData contains all meta data concerning the StructureInterface contract.
var StructureInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StructureInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use StructureInterfaceMetaData.ABI instead.
var StructureInterfaceABI = StructureInterfaceMetaData.ABI

// StructureInterface is an auto generated Go binding around an Ethereum contract.
type StructureInterface struct {
	StructureInterfaceCaller     // Read-only binding to the contract
	StructureInterfaceTransactor // Write-only binding to the contract
	StructureInterfaceFilterer   // Log filterer for contract events
}

// StructureInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type StructureInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructureInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StructureInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructureInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StructureInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructureInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StructureInterfaceSession struct {
	Contract     *StructureInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StructureInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StructureInterfaceCallerSession struct {
	Contract *StructureInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// StructureInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StructureInterfaceTransactorSession struct {
	Contract     *StructureInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// StructureInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type StructureInterfaceRaw struct {
	Contract *StructureInterface // Generic contract binding to access the raw methods on
}

// StructureInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StructureInterfaceCallerRaw struct {
	Contract *StructureInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// StructureInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StructureInterfaceTransactorRaw struct {
	Contract *StructureInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStructureInterface creates a new instance of StructureInterface, bound to a specific deployed contract.
func NewStructureInterface(address common.Address, backend bind.ContractBackend) (*StructureInterface, error) {
	contract, err := bindStructureInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StructureInterface{StructureInterfaceCaller: StructureInterfaceCaller{contract: contract}, StructureInterfaceTransactor: StructureInterfaceTransactor{contract: contract}, StructureInterfaceFilterer: StructureInterfaceFilterer{contract: contract}}, nil
}

// NewStructureInterfaceCaller creates a new read-only instance of StructureInterface, bound to a specific deployed contract.
func NewStructureInterfaceCaller(address common.Address, caller bind.ContractCaller) (*StructureInterfaceCaller, error) {
	contract, err := bindStructureInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StructureInterfaceCaller{contract: contract}, nil
}

// NewStructureInterfaceTransactor creates a new write-only instance of StructureInterface, bound to a specific deployed contract.
func NewStructureInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*StructureInterfaceTransactor, error) {
	contract, err := bindStructureInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StructureInterfaceTransactor{contract: contract}, nil
}

// NewStructureInterfaceFilterer creates a new log filterer instance of StructureInterface, bound to a specific deployed contract.
func NewStructureInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*StructureInterfaceFilterer, error) {
	contract, err := bindStructureInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StructureInterfaceFilterer{contract: contract}, nil
}

// bindStructureInterface binds a generic wrapper to an already deployed contract.
func bindStructureInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StructureInterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StructureInterface *StructureInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StructureInterface.Contract.StructureInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StructureInterface *StructureInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StructureInterface.Contract.StructureInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StructureInterface *StructureInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StructureInterface.Contract.StructureInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StructureInterface *StructureInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StructureInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StructureInterface *StructureInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StructureInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StructureInterface *StructureInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StructureInterface.Contract.contract.Transact(opts, method, params...)
}

// GetValue is a free data retrieval call binding the contract method 0x0ff4c916.
//
// Solidity: function getValue(uint256 _id) view returns(uint256)
func (_StructureInterface *StructureInterfaceCaller) GetValue(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StructureInterface.contract.Call(opts, &out, "getValue", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x0ff4c916.
//
// Solidity: function getValue(uint256 _id) view returns(uint256)
func (_StructureInterface *StructureInterfaceSession) GetValue(_id *big.Int) (*big.Int, error) {
	return _StructureInterface.Contract.GetValue(&_StructureInterface.CallOpts, _id)
}

// GetValue is a free data retrieval call binding the contract method 0x0ff4c916.
//
// Solidity: function getValue(uint256 _id) view returns(uint256)
func (_StructureInterface *StructureInterfaceCallerSession) GetValue(_id *big.Int) (*big.Int, error) {
	return _StructureInterface.Contract.GetValue(&_StructureInterface.CallOpts, _id)
}
