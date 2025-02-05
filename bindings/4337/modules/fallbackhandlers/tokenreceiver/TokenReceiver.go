// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenreceiver

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

// TokenReceiverMetaData contains all meta data concerning the TokenReceiver contract.
var TokenReceiverMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"}],\"name\":\"isModuleType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608080604052346015576101b2908161001b8239f35b600080fdfe60806040526004361015610015575b3661014d57005b6000803560e01c90816301ffc9a714610053575080636d61fe701461004e5780638a91b0e31461004e5763ecd059610361000e5761012d565b6100e0565b346100dd5760203660031901126100dd576004359063ffffffff60e01b82168092036100dd5750630a85bd0160e11b81149081156100cc575b81156100bb575b81156100aa575b501515608052607f1960a0016080f35b6301ffc9a760e01b1490508161009a565b63bc197c8160e01b81149150610093565b63f23a6e6160e01b8114915061008c565b80fd5b346101285760203660031901126101285760043567ffffffffffffffff8082116101285736602383011215610128578160040135908111610128573691016024011161012857005b600080fd5b346101285760203660031901126101285760206040516003600435148152f35b60003560e01c63bc197c81811463f23a6e6182141763150b7a0282141761017357600080fd5b6020526020603cf3fea2646970667358221220701f05114e2bb5597f830b590f6fed7a7fb8438374cf05ab35394f0cbd94a2bf64736f6c63430008190033",
}

// TokenReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenReceiverMetaData.ABI instead.
var TokenReceiverABI = TokenReceiverMetaData.ABI

// TokenReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenReceiverMetaData.Bin instead.
var TokenReceiverBin = TokenReceiverMetaData.Bin

// DeployTokenReceiver deploys a new Ethereum contract, binding an instance of TokenReceiver to it.
func DeployTokenReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenReceiver, error) {
	parsed, err := TokenReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenReceiver{TokenReceiverCaller: TokenReceiverCaller{contract: contract}, TokenReceiverTransactor: TokenReceiverTransactor{contract: contract}, TokenReceiverFilterer: TokenReceiverFilterer{contract: contract}}, nil
}

// TokenReceiver is an auto generated Go binding around an Ethereum contract.
type TokenReceiver struct {
	TokenReceiverCaller     // Read-only binding to the contract
	TokenReceiverTransactor // Write-only binding to the contract
	TokenReceiverFilterer   // Log filterer for contract events
}

// TokenReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenReceiverSession struct {
	Contract     *TokenReceiver    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenReceiverCallerSession struct {
	Contract *TokenReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenReceiverTransactorSession struct {
	Contract     *TokenReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenReceiverRaw struct {
	Contract *TokenReceiver // Generic contract binding to access the raw methods on
}

// TokenReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenReceiverCallerRaw struct {
	Contract *TokenReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// TokenReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenReceiverTransactorRaw struct {
	Contract *TokenReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenReceiver creates a new instance of TokenReceiver, bound to a specific deployed contract.
func NewTokenReceiver(address common.Address, backend bind.ContractBackend) (*TokenReceiver, error) {
	contract, err := bindTokenReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenReceiver{TokenReceiverCaller: TokenReceiverCaller{contract: contract}, TokenReceiverTransactor: TokenReceiverTransactor{contract: contract}, TokenReceiverFilterer: TokenReceiverFilterer{contract: contract}}, nil
}

// NewTokenReceiverCaller creates a new read-only instance of TokenReceiver, bound to a specific deployed contract.
func NewTokenReceiverCaller(address common.Address, caller bind.ContractCaller) (*TokenReceiverCaller, error) {
	contract, err := bindTokenReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenReceiverCaller{contract: contract}, nil
}

// NewTokenReceiverTransactor creates a new write-only instance of TokenReceiver, bound to a specific deployed contract.
func NewTokenReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenReceiverTransactor, error) {
	contract, err := bindTokenReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenReceiverTransactor{contract: contract}, nil
}

// NewTokenReceiverFilterer creates a new log filterer instance of TokenReceiver, bound to a specific deployed contract.
func NewTokenReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenReceiverFilterer, error) {
	contract, err := bindTokenReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenReceiverFilterer{contract: contract}, nil
}

// bindTokenReceiver binds a generic wrapper to an already deployed contract.
func bindTokenReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenReceiver *TokenReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenReceiver.Contract.TokenReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenReceiver *TokenReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenReceiver.Contract.TokenReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenReceiver *TokenReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenReceiver.Contract.TokenReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenReceiver *TokenReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenReceiver *TokenReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenReceiver *TokenReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenReceiver.Contract.contract.Transact(opts, method, params...)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_TokenReceiver *TokenReceiverCaller) IsModuleType(opts *bind.CallOpts, moduleTypeId *big.Int) (bool, error) {
	var out []interface{}
	err := _TokenReceiver.contract.Call(opts, &out, "isModuleType", moduleTypeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_TokenReceiver *TokenReceiverSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _TokenReceiver.Contract.IsModuleType(&_TokenReceiver.CallOpts, moduleTypeId)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_TokenReceiver *TokenReceiverCallerSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _TokenReceiver.Contract.IsModuleType(&_TokenReceiver.CallOpts, moduleTypeId)
}

// OnInstall is a free data retrieval call binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) pure returns()
func (_TokenReceiver *TokenReceiverCaller) OnInstall(opts *bind.CallOpts, data []byte) error {
	var out []interface{}
	err := _TokenReceiver.contract.Call(opts, &out, "onInstall", data)

	if err != nil {
		return err
	}

	return err

}

// OnInstall is a free data retrieval call binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) pure returns()
func (_TokenReceiver *TokenReceiverSession) OnInstall(data []byte) error {
	return _TokenReceiver.Contract.OnInstall(&_TokenReceiver.CallOpts, data)
}

// OnInstall is a free data retrieval call binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) pure returns()
func (_TokenReceiver *TokenReceiverCallerSession) OnInstall(data []byte) error {
	return _TokenReceiver.Contract.OnInstall(&_TokenReceiver.CallOpts, data)
}

// OnUninstall is a free data retrieval call binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) pure returns()
func (_TokenReceiver *TokenReceiverCaller) OnUninstall(opts *bind.CallOpts, data []byte) error {
	var out []interface{}
	err := _TokenReceiver.contract.Call(opts, &out, "onUninstall", data)

	if err != nil {
		return err
	}

	return err

}

// OnUninstall is a free data retrieval call binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) pure returns()
func (_TokenReceiver *TokenReceiverSession) OnUninstall(data []byte) error {
	return _TokenReceiver.Contract.OnUninstall(&_TokenReceiver.CallOpts, data)
}

// OnUninstall is a free data retrieval call binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) pure returns()
func (_TokenReceiver *TokenReceiverCallerSession) OnUninstall(data []byte) error {
	return _TokenReceiver.Contract.OnUninstall(&_TokenReceiver.CallOpts, data)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenReceiver *TokenReceiverCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TokenReceiver.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenReceiver *TokenReceiverSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenReceiver.Contract.SupportsInterface(&_TokenReceiver.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenReceiver *TokenReceiverCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenReceiver.Contract.SupportsInterface(&_TokenReceiver.CallOpts, interfaceId)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TokenReceiver *TokenReceiverTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TokenReceiver.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TokenReceiver *TokenReceiverSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TokenReceiver.Contract.Fallback(&_TokenReceiver.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TokenReceiver *TokenReceiverTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TokenReceiver.Contract.Fallback(&_TokenReceiver.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TokenReceiver *TokenReceiverTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenReceiver.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TokenReceiver *TokenReceiverSession) Receive() (*types.Transaction, error) {
	return _TokenReceiver.Contract.Receive(&_TokenReceiver.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TokenReceiver *TokenReceiverTransactorSession) Receive() (*types.Transaction, error) {
	return _TokenReceiver.Contract.Receive(&_TokenReceiver.TransactOpts)
}
