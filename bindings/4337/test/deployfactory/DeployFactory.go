// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deployfactory

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

// DeployFactoryMetaData contains all meta data concerning the DeployFactory contract.
var DeployFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"createdContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608080604052346015576101a8908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806348aac3921461006257634af63f021461003257600080fd5b3461005d57602061004b61004536610091565b9061013f565b6040516001600160a01b039091168152f35b600080fd5b3461005d57602061004b61007536610091565b90610117565b634e487b7160e01b600052604160045260246000fd5b604060031982011261005d5767ffffffffffffffff9060043582811161005d578160238201121561005d578060040135918383116101125760405193601f8401601f19908116603f0116850190811185821017610112576040528284526024838301011161005d578160009260246020930183860137830101529060243590565b61007b565b605591600b91602081519101209060405191604083015260208201523081520160ff81532090565b6101498282610117565b803b61016357506020815191016000f590813b1561005d57565b6001600160a01b03169291505056fea2646970667358221220c217f099f64ff80db1e7fc06335bb842fc67765c1ab53658a2dbe68b6e0e3c3e64736f6c63430008190033",
}

// DeployFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use DeployFactoryMetaData.ABI instead.
var DeployFactoryABI = DeployFactoryMetaData.ABI

// DeployFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DeployFactoryMetaData.Bin instead.
var DeployFactoryBin = DeployFactoryMetaData.Bin

// DeployDeployFactory deploys a new Ethereum contract, binding an instance of DeployFactory to it.
func DeployDeployFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DeployFactory, error) {
	parsed, err := DeployFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DeployFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DeployFactory{DeployFactoryCaller: DeployFactoryCaller{contract: contract}, DeployFactoryTransactor: DeployFactoryTransactor{contract: contract}, DeployFactoryFilterer: DeployFactoryFilterer{contract: contract}}, nil
}

// DeployFactory is an auto generated Go binding around an Ethereum contract.
type DeployFactory struct {
	DeployFactoryCaller     // Read-only binding to the contract
	DeployFactoryTransactor // Write-only binding to the contract
	DeployFactoryFilterer   // Log filterer for contract events
}

// DeployFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeployFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeployFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeployFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeployFactorySession struct {
	Contract     *DeployFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeployFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeployFactoryCallerSession struct {
	Contract *DeployFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DeployFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeployFactoryTransactorSession struct {
	Contract     *DeployFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DeployFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeployFactoryRaw struct {
	Contract *DeployFactory // Generic contract binding to access the raw methods on
}

// DeployFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeployFactoryCallerRaw struct {
	Contract *DeployFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// DeployFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeployFactoryTransactorRaw struct {
	Contract *DeployFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeployFactory creates a new instance of DeployFactory, bound to a specific deployed contract.
func NewDeployFactory(address common.Address, backend bind.ContractBackend) (*DeployFactory, error) {
	contract, err := bindDeployFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeployFactory{DeployFactoryCaller: DeployFactoryCaller{contract: contract}, DeployFactoryTransactor: DeployFactoryTransactor{contract: contract}, DeployFactoryFilterer: DeployFactoryFilterer{contract: contract}}, nil
}

// NewDeployFactoryCaller creates a new read-only instance of DeployFactory, bound to a specific deployed contract.
func NewDeployFactoryCaller(address common.Address, caller bind.ContractCaller) (*DeployFactoryCaller, error) {
	contract, err := bindDeployFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeployFactoryCaller{contract: contract}, nil
}

// NewDeployFactoryTransactor creates a new write-only instance of DeployFactory, bound to a specific deployed contract.
func NewDeployFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*DeployFactoryTransactor, error) {
	contract, err := bindDeployFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeployFactoryTransactor{contract: contract}, nil
}

// NewDeployFactoryFilterer creates a new log filterer instance of DeployFactory, bound to a specific deployed contract.
func NewDeployFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*DeployFactoryFilterer, error) {
	contract, err := bindDeployFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeployFactoryFilterer{contract: contract}, nil
}

// bindDeployFactory binds a generic wrapper to an already deployed contract.
func bindDeployFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DeployFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeployFactory *DeployFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeployFactory.Contract.DeployFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeployFactory *DeployFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeployFactory.Contract.DeployFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeployFactory *DeployFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeployFactory.Contract.DeployFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeployFactory *DeployFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeployFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeployFactory *DeployFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeployFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeployFactory *DeployFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeployFactory.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x48aac392.
//
// Solidity: function getAddress(bytes _initCode, bytes32 _salt) view returns(address)
func (_DeployFactory *DeployFactoryCaller) GetAddress(opts *bind.CallOpts, _initCode []byte, _salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _DeployFactory.contract.Call(opts, &out, "getAddress", _initCode, _salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x48aac392.
//
// Solidity: function getAddress(bytes _initCode, bytes32 _salt) view returns(address)
func (_DeployFactory *DeployFactorySession) GetAddress(_initCode []byte, _salt [32]byte) (common.Address, error) {
	return _DeployFactory.Contract.GetAddress(&_DeployFactory.CallOpts, _initCode, _salt)
}

// GetAddress is a free data retrieval call binding the contract method 0x48aac392.
//
// Solidity: function getAddress(bytes _initCode, bytes32 _salt) view returns(address)
func (_DeployFactory *DeployFactoryCallerSession) GetAddress(_initCode []byte, _salt [32]byte) (common.Address, error) {
	return _DeployFactory.Contract.GetAddress(&_DeployFactory.CallOpts, _initCode, _salt)
}

// Deploy is a paid mutator transaction binding the contract method 0x4af63f02.
//
// Solidity: function deploy(bytes _initCode, bytes32 _salt) returns(address createdContract)
func (_DeployFactory *DeployFactoryTransactor) Deploy(opts *bind.TransactOpts, _initCode []byte, _salt [32]byte) (*types.Transaction, error) {
	return _DeployFactory.contract.Transact(opts, "deploy", _initCode, _salt)
}

// Deploy is a paid mutator transaction binding the contract method 0x4af63f02.
//
// Solidity: function deploy(bytes _initCode, bytes32 _salt) returns(address createdContract)
func (_DeployFactory *DeployFactorySession) Deploy(_initCode []byte, _salt [32]byte) (*types.Transaction, error) {
	return _DeployFactory.Contract.Deploy(&_DeployFactory.TransactOpts, _initCode, _salt)
}

// Deploy is a paid mutator transaction binding the contract method 0x4af63f02.
//
// Solidity: function deploy(bytes _initCode, bytes32 _salt) returns(address createdContract)
func (_DeployFactory *DeployFactoryTransactorSession) Deploy(_initCode []byte, _salt [32]byte) (*types.Transaction, error) {
	return _DeployFactory.Contract.Deploy(&_DeployFactory.TransactOpts, _initCode, _salt)
}
