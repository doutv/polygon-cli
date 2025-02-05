// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package istatelessvalidator

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

// IStatelessValidatorMetaData contains all meta data concerning the IStatelessValidator contract.
var IStatelessValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"}],\"name\":\"isModuleType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"validateSignatureWithData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IStatelessValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use IStatelessValidatorMetaData.ABI instead.
var IStatelessValidatorABI = IStatelessValidatorMetaData.ABI

// IStatelessValidator is an auto generated Go binding around an Ethereum contract.
type IStatelessValidator struct {
	IStatelessValidatorCaller     // Read-only binding to the contract
	IStatelessValidatorTransactor // Write-only binding to the contract
	IStatelessValidatorFilterer   // Log filterer for contract events
}

// IStatelessValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStatelessValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStatelessValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStatelessValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStatelessValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStatelessValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStatelessValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStatelessValidatorSession struct {
	Contract     *IStatelessValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IStatelessValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStatelessValidatorCallerSession struct {
	Contract *IStatelessValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IStatelessValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStatelessValidatorTransactorSession struct {
	Contract     *IStatelessValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IStatelessValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStatelessValidatorRaw struct {
	Contract *IStatelessValidator // Generic contract binding to access the raw methods on
}

// IStatelessValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStatelessValidatorCallerRaw struct {
	Contract *IStatelessValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// IStatelessValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStatelessValidatorTransactorRaw struct {
	Contract *IStatelessValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStatelessValidator creates a new instance of IStatelessValidator, bound to a specific deployed contract.
func NewIStatelessValidator(address common.Address, backend bind.ContractBackend) (*IStatelessValidator, error) {
	contract, err := bindIStatelessValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStatelessValidator{IStatelessValidatorCaller: IStatelessValidatorCaller{contract: contract}, IStatelessValidatorTransactor: IStatelessValidatorTransactor{contract: contract}, IStatelessValidatorFilterer: IStatelessValidatorFilterer{contract: contract}}, nil
}

// NewIStatelessValidatorCaller creates a new read-only instance of IStatelessValidator, bound to a specific deployed contract.
func NewIStatelessValidatorCaller(address common.Address, caller bind.ContractCaller) (*IStatelessValidatorCaller, error) {
	contract, err := bindIStatelessValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStatelessValidatorCaller{contract: contract}, nil
}

// NewIStatelessValidatorTransactor creates a new write-only instance of IStatelessValidator, bound to a specific deployed contract.
func NewIStatelessValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IStatelessValidatorTransactor, error) {
	contract, err := bindIStatelessValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStatelessValidatorTransactor{contract: contract}, nil
}

// NewIStatelessValidatorFilterer creates a new log filterer instance of IStatelessValidator, bound to a specific deployed contract.
func NewIStatelessValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IStatelessValidatorFilterer, error) {
	contract, err := bindIStatelessValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStatelessValidatorFilterer{contract: contract}, nil
}

// bindIStatelessValidator binds a generic wrapper to an already deployed contract.
func bindIStatelessValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStatelessValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStatelessValidator *IStatelessValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStatelessValidator.Contract.IStatelessValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStatelessValidator *IStatelessValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStatelessValidator.Contract.IStatelessValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStatelessValidator *IStatelessValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStatelessValidator.Contract.IStatelessValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStatelessValidator *IStatelessValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStatelessValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStatelessValidator *IStatelessValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStatelessValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStatelessValidator *IStatelessValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStatelessValidator.Contract.contract.Transact(opts, method, params...)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) view returns(bool)
func (_IStatelessValidator *IStatelessValidatorCaller) IsModuleType(opts *bind.CallOpts, moduleTypeId *big.Int) (bool, error) {
	var out []interface{}
	err := _IStatelessValidator.contract.Call(opts, &out, "isModuleType", moduleTypeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) view returns(bool)
func (_IStatelessValidator *IStatelessValidatorSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _IStatelessValidator.Contract.IsModuleType(&_IStatelessValidator.CallOpts, moduleTypeId)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) view returns(bool)
func (_IStatelessValidator *IStatelessValidatorCallerSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _IStatelessValidator.Contract.IsModuleType(&_IStatelessValidator.CallOpts, moduleTypeId)
}

// ValidateSignatureWithData is a free data retrieval call binding the contract method 0x940d3840.
//
// Solidity: function validateSignatureWithData(bytes32 hash, bytes signature, bytes data) view returns(bool)
func (_IStatelessValidator *IStatelessValidatorCaller) ValidateSignatureWithData(opts *bind.CallOpts, hash [32]byte, signature []byte, data []byte) (bool, error) {
	var out []interface{}
	err := _IStatelessValidator.contract.Call(opts, &out, "validateSignatureWithData", hash, signature, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateSignatureWithData is a free data retrieval call binding the contract method 0x940d3840.
//
// Solidity: function validateSignatureWithData(bytes32 hash, bytes signature, bytes data) view returns(bool)
func (_IStatelessValidator *IStatelessValidatorSession) ValidateSignatureWithData(hash [32]byte, signature []byte, data []byte) (bool, error) {
	return _IStatelessValidator.Contract.ValidateSignatureWithData(&_IStatelessValidator.CallOpts, hash, signature, data)
}

// ValidateSignatureWithData is a free data retrieval call binding the contract method 0x940d3840.
//
// Solidity: function validateSignatureWithData(bytes32 hash, bytes signature, bytes data) view returns(bool)
func (_IStatelessValidator *IStatelessValidatorCallerSession) ValidateSignatureWithData(hash [32]byte, signature []byte, data []byte) (bool, error) {
	return _IStatelessValidator.Contract.ValidateSignatureWithData(&_IStatelessValidator.CallOpts, hash, signature, data)
}
