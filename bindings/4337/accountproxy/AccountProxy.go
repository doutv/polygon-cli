// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accountproxy

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

// AccountProxyMetaData contains all meta data concerning the AccountProxy contract.
var AccountProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"name\":\"AccountCreationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletInitialized\",\"type\":\"error\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_initializer\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60808060405234601557610260908161001b8239f35b600080fdfe60806040526004361015610024575b361561001f5734156101eb57600080fd5b6101eb565b6000803560e01c63d1f578941461003b575061000e565b346100af5760403660031901126100af576004356001600160a01b03811681036100ab576024359067ffffffffffffffff908183116100a757366023840112156100a75782600401359182116100a75736602483850101116100a75760246100a4930190610111565b80f35b8380fd5b5080fd5b80fd5b634e487b7160e01b600052604160045260246000fd5b6020808252825181830181905290939260005b8281106100fd57505060409293506000838284010152601f8019910116010190565b8181018601518482016040015285016100db565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8054929390926001600160a01b03166101d95760009382859455816040519283928337810184815203915af43d156101d15767ffffffffffffffff903d8281116101cc5760405192601f8201601f19908116603f01168401908111848210176101cc5760405282523d6000602084013e5b156101ab5750565b604051633018224d60e21b81529081906101c890600483016100c8565b0390fd5b6100b2565b6060906101a3565b60405163c28d69c760e01b8152600490fd5b600036818037808036817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc545af43d82803e15610226573d90f35b3d90fdfea2646970667358221220dbc2ad734b73b998278ed7bcd9ef361791ffe94471f015bd78d961860b387f4664736f6c63430008190033",
}

// AccountProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountProxyMetaData.ABI instead.
var AccountProxyABI = AccountProxyMetaData.ABI

// AccountProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccountProxyMetaData.Bin instead.
var AccountProxyBin = AccountProxyMetaData.Bin

// DeployAccountProxy deploys a new Ethereum contract, binding an instance of AccountProxy to it.
func DeployAccountProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccountProxy, error) {
	parsed, err := AccountProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccountProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountProxy{AccountProxyCaller: AccountProxyCaller{contract: contract}, AccountProxyTransactor: AccountProxyTransactor{contract: contract}, AccountProxyFilterer: AccountProxyFilterer{contract: contract}}, nil
}

// AccountProxy is an auto generated Go binding around an Ethereum contract.
type AccountProxy struct {
	AccountProxyCaller     // Read-only binding to the contract
	AccountProxyTransactor // Write-only binding to the contract
	AccountProxyFilterer   // Log filterer for contract events
}

// AccountProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountProxySession struct {
	Contract     *AccountProxy     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountProxyCallerSession struct {
	Contract *AccountProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AccountProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountProxyTransactorSession struct {
	Contract     *AccountProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AccountProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountProxyRaw struct {
	Contract *AccountProxy // Generic contract binding to access the raw methods on
}

// AccountProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountProxyCallerRaw struct {
	Contract *AccountProxyCaller // Generic read-only contract binding to access the raw methods on
}

// AccountProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountProxyTransactorRaw struct {
	Contract *AccountProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountProxy creates a new instance of AccountProxy, bound to a specific deployed contract.
func NewAccountProxy(address common.Address, backend bind.ContractBackend) (*AccountProxy, error) {
	contract, err := bindAccountProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountProxy{AccountProxyCaller: AccountProxyCaller{contract: contract}, AccountProxyTransactor: AccountProxyTransactor{contract: contract}, AccountProxyFilterer: AccountProxyFilterer{contract: contract}}, nil
}

// NewAccountProxyCaller creates a new read-only instance of AccountProxy, bound to a specific deployed contract.
func NewAccountProxyCaller(address common.Address, caller bind.ContractCaller) (*AccountProxyCaller, error) {
	contract, err := bindAccountProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountProxyCaller{contract: contract}, nil
}

// NewAccountProxyTransactor creates a new write-only instance of AccountProxy, bound to a specific deployed contract.
func NewAccountProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountProxyTransactor, error) {
	contract, err := bindAccountProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountProxyTransactor{contract: contract}, nil
}

// NewAccountProxyFilterer creates a new log filterer instance of AccountProxy, bound to a specific deployed contract.
func NewAccountProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountProxyFilterer, error) {
	contract, err := bindAccountProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountProxyFilterer{contract: contract}, nil
}

// bindAccountProxy binds a generic wrapper to an already deployed contract.
func bindAccountProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountProxy *AccountProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountProxy.Contract.AccountProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountProxy *AccountProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountProxy.Contract.AccountProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountProxy *AccountProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountProxy.Contract.AccountProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountProxy *AccountProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountProxy *AccountProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountProxy *AccountProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountProxy.Contract.contract.Transact(opts, method, params...)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _implementation, bytes _initializer) returns()
func (_AccountProxy *AccountProxyTransactor) Initialize(opts *bind.TransactOpts, _implementation common.Address, _initializer []byte) (*types.Transaction, error) {
	return _AccountProxy.contract.Transact(opts, "initialize", _implementation, _initializer)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _implementation, bytes _initializer) returns()
func (_AccountProxy *AccountProxySession) Initialize(_implementation common.Address, _initializer []byte) (*types.Transaction, error) {
	return _AccountProxy.Contract.Initialize(&_AccountProxy.TransactOpts, _implementation, _initializer)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _implementation, bytes _initializer) returns()
func (_AccountProxy *AccountProxyTransactorSession) Initialize(_implementation common.Address, _initializer []byte) (*types.Transaction, error) {
	return _AccountProxy.Contract.Initialize(&_AccountProxy.TransactOpts, _implementation, _initializer)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_AccountProxy *AccountProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _AccountProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_AccountProxy *AccountProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AccountProxy.Contract.Fallback(&_AccountProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_AccountProxy *AccountProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AccountProxy.Contract.Fallback(&_AccountProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AccountProxy *AccountProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AccountProxy *AccountProxySession) Receive() (*types.Transaction, error) {
	return _AccountProxy.Contract.Receive(&_AccountProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AccountProxy *AccountProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _AccountProxy.Contract.Receive(&_AccountProxy.TransactOpts)
}
