// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaccountfactory

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

// IAccountFactoryMetaData contains all meta data concerning the IAccountFactory contract.
var IAccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_initializer\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"computeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"createAccountWithSignature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isValidAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IAccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccountFactoryMetaData.ABI instead.
var IAccountFactoryABI = IAccountFactoryMetaData.ABI

// IAccountFactory is an auto generated Go binding around an Ethereum contract.
type IAccountFactory struct {
	IAccountFactoryCaller     // Read-only binding to the contract
	IAccountFactoryTransactor // Write-only binding to the contract
	IAccountFactoryFilterer   // Log filterer for contract events
}

// IAccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccountFactorySession struct {
	Contract     *IAccountFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccountFactoryCallerSession struct {
	Contract *IAccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IAccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccountFactoryTransactorSession struct {
	Contract     *IAccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IAccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccountFactoryRaw struct {
	Contract *IAccountFactory // Generic contract binding to access the raw methods on
}

// IAccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccountFactoryCallerRaw struct {
	Contract *IAccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IAccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccountFactoryTransactorRaw struct {
	Contract *IAccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccountFactory creates a new instance of IAccountFactory, bound to a specific deployed contract.
func NewIAccountFactory(address common.Address, backend bind.ContractBackend) (*IAccountFactory, error) {
	contract, err := bindIAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccountFactory{IAccountFactoryCaller: IAccountFactoryCaller{contract: contract}, IAccountFactoryTransactor: IAccountFactoryTransactor{contract: contract}, IAccountFactoryFilterer: IAccountFactoryFilterer{contract: contract}}, nil
}

// NewIAccountFactoryCaller creates a new read-only instance of IAccountFactory, bound to a specific deployed contract.
func NewIAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*IAccountFactoryCaller, error) {
	contract, err := bindIAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountFactoryCaller{contract: contract}, nil
}

// NewIAccountFactoryTransactor creates a new write-only instance of IAccountFactory, bound to a specific deployed contract.
func NewIAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccountFactoryTransactor, error) {
	contract, err := bindIAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountFactoryTransactor{contract: contract}, nil
}

// NewIAccountFactoryFilterer creates a new log filterer instance of IAccountFactory, bound to a specific deployed contract.
func NewIAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccountFactoryFilterer, error) {
	contract, err := bindIAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccountFactoryFilterer{contract: contract}, nil
}

// bindIAccountFactory binds a generic wrapper to an already deployed contract.
func bindIAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountFactory *IAccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountFactory.Contract.IAccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountFactory *IAccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountFactory.Contract.IAccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountFactory *IAccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountFactory.Contract.IAccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountFactory *IAccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountFactory *IAccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountFactory *IAccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountFactory.Contract.contract.Transact(opts, method, params...)
}

// ComputeAddress is a free data retrieval call binding the contract method 0x36b5aa2d.
//
// Solidity: function computeAddress(address _implementation, uint256 _salt) view returns(address)
func (_IAccountFactory *IAccountFactoryCaller) ComputeAddress(opts *bind.CallOpts, _implementation common.Address, _salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IAccountFactory.contract.Call(opts, &out, "computeAddress", _implementation, _salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeAddress is a free data retrieval call binding the contract method 0x36b5aa2d.
//
// Solidity: function computeAddress(address _implementation, uint256 _salt) view returns(address)
func (_IAccountFactory *IAccountFactorySession) ComputeAddress(_implementation common.Address, _salt *big.Int) (common.Address, error) {
	return _IAccountFactory.Contract.ComputeAddress(&_IAccountFactory.CallOpts, _implementation, _salt)
}

// ComputeAddress is a free data retrieval call binding the contract method 0x36b5aa2d.
//
// Solidity: function computeAddress(address _implementation, uint256 _salt) view returns(address)
func (_IAccountFactory *IAccountFactoryCallerSession) ComputeAddress(_implementation common.Address, _salt *big.Int) (common.Address, error) {
	return _IAccountFactory.Contract.ComputeAddress(&_IAccountFactory.CallOpts, _implementation, _salt)
}

// IsValidAccount is a free data retrieval call binding the contract method 0x23cca69c.
//
// Solidity: function isValidAccount(address account) view returns(bool)
func (_IAccountFactory *IAccountFactoryCaller) IsValidAccount(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAccountFactory.contract.Call(opts, &out, "isValidAccount", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAccount is a free data retrieval call binding the contract method 0x23cca69c.
//
// Solidity: function isValidAccount(address account) view returns(bool)
func (_IAccountFactory *IAccountFactorySession) IsValidAccount(account common.Address) (bool, error) {
	return _IAccountFactory.Contract.IsValidAccount(&_IAccountFactory.CallOpts, account)
}

// IsValidAccount is a free data retrieval call binding the contract method 0x23cca69c.
//
// Solidity: function isValidAccount(address account) view returns(bool)
func (_IAccountFactory *IAccountFactoryCallerSession) IsValidAccount(account common.Address) (bool, error) {
	return _IAccountFactory.Contract.IsValidAccount(&_IAccountFactory.CallOpts, account)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x296601cd.
//
// Solidity: function createAccount(address _implementation, bytes _initializer, uint256 _salt) returns(address)
func (_IAccountFactory *IAccountFactoryTransactor) CreateAccount(opts *bind.TransactOpts, _implementation common.Address, _initializer []byte, _salt *big.Int) (*types.Transaction, error) {
	return _IAccountFactory.contract.Transact(opts, "createAccount", _implementation, _initializer, _salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x296601cd.
//
// Solidity: function createAccount(address _implementation, bytes _initializer, uint256 _salt) returns(address)
func (_IAccountFactory *IAccountFactorySession) CreateAccount(_implementation common.Address, _initializer []byte, _salt *big.Int) (*types.Transaction, error) {
	return _IAccountFactory.Contract.CreateAccount(&_IAccountFactory.TransactOpts, _implementation, _initializer, _salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x296601cd.
//
// Solidity: function createAccount(address _implementation, bytes _initializer, uint256 _salt) returns(address)
func (_IAccountFactory *IAccountFactoryTransactorSession) CreateAccount(_implementation common.Address, _initializer []byte, _salt *big.Int) (*types.Transaction, error) {
	return _IAccountFactory.Contract.CreateAccount(&_IAccountFactory.TransactOpts, _implementation, _initializer, _salt)
}

// CreateAccountWithSignature is a paid mutator transaction binding the contract method 0xddede0a1.
//
// Solidity: function createAccountWithSignature(address _implementation, bytes _initializer, uint256 _salt, bytes _signature) returns(address)
func (_IAccountFactory *IAccountFactoryTransactor) CreateAccountWithSignature(opts *bind.TransactOpts, _implementation common.Address, _initializer []byte, _salt *big.Int, _signature []byte) (*types.Transaction, error) {
	return _IAccountFactory.contract.Transact(opts, "createAccountWithSignature", _implementation, _initializer, _salt, _signature)
}

// CreateAccountWithSignature is a paid mutator transaction binding the contract method 0xddede0a1.
//
// Solidity: function createAccountWithSignature(address _implementation, bytes _initializer, uint256 _salt, bytes _signature) returns(address)
func (_IAccountFactory *IAccountFactorySession) CreateAccountWithSignature(_implementation common.Address, _initializer []byte, _salt *big.Int, _signature []byte) (*types.Transaction, error) {
	return _IAccountFactory.Contract.CreateAccountWithSignature(&_IAccountFactory.TransactOpts, _implementation, _initializer, _salt, _signature)
}

// CreateAccountWithSignature is a paid mutator transaction binding the contract method 0xddede0a1.
//
// Solidity: function createAccountWithSignature(address _implementation, bytes _initializer, uint256 _salt, bytes _signature) returns(address)
func (_IAccountFactory *IAccountFactoryTransactorSession) CreateAccountWithSignature(_implementation common.Address, _initializer []byte, _salt *big.Int, _signature []byte) (*types.Transaction, error) {
	return _IAccountFactory.Contract.CreateAccountWithSignature(&_IAccountFactory.TransactOpts, _implementation, _initializer, _salt, _signature)
}

// IAccountFactoryAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the IAccountFactory contract.
type IAccountFactoryAccountCreatedIterator struct {
	Event *IAccountFactoryAccountCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAccountFactoryAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountFactoryAccountCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAccountFactoryAccountCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAccountFactoryAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountFactoryAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountFactoryAccountCreated represents a AccountCreated event raised by the IAccountFactory contract.
type IAccountFactoryAccountCreated struct {
	Account        common.Address
	Implementation common.Address
	Initializer    []byte
	Salt           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0xa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f7437.
//
// Solidity: event AccountCreated(address indexed account, address _implementation, bytes _initializer, uint256 _salt)
func (_IAccountFactory *IAccountFactoryFilterer) FilterAccountCreated(opts *bind.FilterOpts, account []common.Address) (*IAccountFactoryAccountCreatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAccountFactory.contract.FilterLogs(opts, "AccountCreated", accountRule)
	if err != nil {
		return nil, err
	}
	return &IAccountFactoryAccountCreatedIterator{contract: _IAccountFactory.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0xa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f7437.
//
// Solidity: event AccountCreated(address indexed account, address _implementation, bytes _initializer, uint256 _salt)
func (_IAccountFactory *IAccountFactoryFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *IAccountFactoryAccountCreated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAccountFactory.contract.WatchLogs(opts, "AccountCreated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountFactoryAccountCreated)
				if err := _IAccountFactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountCreated is a log parse operation binding the contract event 0xa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f7437.
//
// Solidity: event AccountCreated(address indexed account, address _implementation, bytes _initializer, uint256 _salt)
func (_IAccountFactory *IAccountFactoryFilterer) ParseAccountCreated(log types.Log) (*IAccountFactoryAccountCreated, error) {
	event := new(IAccountFactoryAccountCreated)
	if err := _IAccountFactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
