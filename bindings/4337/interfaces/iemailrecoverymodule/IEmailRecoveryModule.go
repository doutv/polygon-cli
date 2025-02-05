// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iemailrecoverymodule

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

// IEmailRecoveryModuleMetaData contains all meta data concerning the IEmailRecoveryModule contract.
var IEmailRecoveryModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDkimKeyHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEmailHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recoveringAccount\",\"type\":\"address\"}],\"name\":\"InvalidRecoveringAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"emailHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"EmailHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"OracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"}],\"name\":\"SignerUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getEmailHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lastUpdatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"}],\"name\":\"isModuleType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_emailHash\",\"type\":\"bytes32\"}],\"name\":\"updateEmailHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IEmailRecoveryModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use IEmailRecoveryModuleMetaData.ABI instead.
var IEmailRecoveryModuleABI = IEmailRecoveryModuleMetaData.ABI

// IEmailRecoveryModule is an auto generated Go binding around an Ethereum contract.
type IEmailRecoveryModule struct {
	IEmailRecoveryModuleCaller     // Read-only binding to the contract
	IEmailRecoveryModuleTransactor // Write-only binding to the contract
	IEmailRecoveryModuleFilterer   // Log filterer for contract events
}

// IEmailRecoveryModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEmailRecoveryModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEmailRecoveryModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEmailRecoveryModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEmailRecoveryModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEmailRecoveryModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEmailRecoveryModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEmailRecoveryModuleSession struct {
	Contract     *IEmailRecoveryModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IEmailRecoveryModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEmailRecoveryModuleCallerSession struct {
	Contract *IEmailRecoveryModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IEmailRecoveryModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEmailRecoveryModuleTransactorSession struct {
	Contract     *IEmailRecoveryModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IEmailRecoveryModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEmailRecoveryModuleRaw struct {
	Contract *IEmailRecoveryModule // Generic contract binding to access the raw methods on
}

// IEmailRecoveryModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEmailRecoveryModuleCallerRaw struct {
	Contract *IEmailRecoveryModuleCaller // Generic read-only contract binding to access the raw methods on
}

// IEmailRecoveryModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEmailRecoveryModuleTransactorRaw struct {
	Contract *IEmailRecoveryModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEmailRecoveryModule creates a new instance of IEmailRecoveryModule, bound to a specific deployed contract.
func NewIEmailRecoveryModule(address common.Address, backend bind.ContractBackend) (*IEmailRecoveryModule, error) {
	contract, err := bindIEmailRecoveryModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEmailRecoveryModule{IEmailRecoveryModuleCaller: IEmailRecoveryModuleCaller{contract: contract}, IEmailRecoveryModuleTransactor: IEmailRecoveryModuleTransactor{contract: contract}, IEmailRecoveryModuleFilterer: IEmailRecoveryModuleFilterer{contract: contract}}, nil
}

// NewIEmailRecoveryModuleCaller creates a new read-only instance of IEmailRecoveryModule, bound to a specific deployed contract.
func NewIEmailRecoveryModuleCaller(address common.Address, caller bind.ContractCaller) (*IEmailRecoveryModuleCaller, error) {
	contract, err := bindIEmailRecoveryModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEmailRecoveryModuleCaller{contract: contract}, nil
}

// NewIEmailRecoveryModuleTransactor creates a new write-only instance of IEmailRecoveryModule, bound to a specific deployed contract.
func NewIEmailRecoveryModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*IEmailRecoveryModuleTransactor, error) {
	contract, err := bindIEmailRecoveryModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEmailRecoveryModuleTransactor{contract: contract}, nil
}

// NewIEmailRecoveryModuleFilterer creates a new log filterer instance of IEmailRecoveryModule, bound to a specific deployed contract.
func NewIEmailRecoveryModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*IEmailRecoveryModuleFilterer, error) {
	contract, err := bindIEmailRecoveryModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEmailRecoveryModuleFilterer{contract: contract}, nil
}

// bindIEmailRecoveryModule binds a generic wrapper to an already deployed contract.
func bindIEmailRecoveryModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEmailRecoveryModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEmailRecoveryModule *IEmailRecoveryModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEmailRecoveryModule.Contract.IEmailRecoveryModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEmailRecoveryModule *IEmailRecoveryModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.IEmailRecoveryModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEmailRecoveryModule *IEmailRecoveryModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.IEmailRecoveryModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEmailRecoveryModule *IEmailRecoveryModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEmailRecoveryModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEmailRecoveryModule *IEmailRecoveryModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEmailRecoveryModule *IEmailRecoveryModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.contract.Transact(opts, method, params...)
}

// GetEmailHash is a free data retrieval call binding the contract method 0x3adde5bf.
//
// Solidity: function getEmailHash(address _account) view returns(bytes32 _hash, uint256 _lastUpdatedAt)
func (_IEmailRecoveryModule *IEmailRecoveryModuleCaller) GetEmailHash(opts *bind.CallOpts, _account common.Address) (struct {
	Hash          [32]byte
	LastUpdatedAt *big.Int
}, error) {
	var out []interface{}
	err := _IEmailRecoveryModule.contract.Call(opts, &out, "getEmailHash", _account)

	outstruct := new(struct {
		Hash          [32]byte
		LastUpdatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.LastUpdatedAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetEmailHash is a free data retrieval call binding the contract method 0x3adde5bf.
//
// Solidity: function getEmailHash(address _account) view returns(bytes32 _hash, uint256 _lastUpdatedAt)
func (_IEmailRecoveryModule *IEmailRecoveryModuleSession) GetEmailHash(_account common.Address) (struct {
	Hash          [32]byte
	LastUpdatedAt *big.Int
}, error) {
	return _IEmailRecoveryModule.Contract.GetEmailHash(&_IEmailRecoveryModule.CallOpts, _account)
}

// GetEmailHash is a free data retrieval call binding the contract method 0x3adde5bf.
//
// Solidity: function getEmailHash(address _account) view returns(bytes32 _hash, uint256 _lastUpdatedAt)
func (_IEmailRecoveryModule *IEmailRecoveryModuleCallerSession) GetEmailHash(_account common.Address) (struct {
	Hash          [32]byte
	LastUpdatedAt *big.Int
}, error) {
	return _IEmailRecoveryModule.Contract.GetEmailHash(&_IEmailRecoveryModule.CallOpts, _account)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_IEmailRecoveryModule *IEmailRecoveryModuleCaller) GetOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEmailRecoveryModule.contract.Call(opts, &out, "getOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_IEmailRecoveryModule *IEmailRecoveryModuleSession) GetOracle() (common.Address, error) {
	return _IEmailRecoveryModule.Contract.GetOracle(&_IEmailRecoveryModule.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_IEmailRecoveryModule *IEmailRecoveryModuleCallerSession) GetOracle() (common.Address, error) {
	return _IEmailRecoveryModule.Contract.GetOracle(&_IEmailRecoveryModule.CallOpts)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) view returns(bool)
func (_IEmailRecoveryModule *IEmailRecoveryModuleCaller) IsModuleType(opts *bind.CallOpts, moduleTypeId *big.Int) (bool, error) {
	var out []interface{}
	err := _IEmailRecoveryModule.contract.Call(opts, &out, "isModuleType", moduleTypeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) view returns(bool)
func (_IEmailRecoveryModule *IEmailRecoveryModuleSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _IEmailRecoveryModule.Contract.IsModuleType(&_IEmailRecoveryModule.CallOpts, moduleTypeId)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) view returns(bool)
func (_IEmailRecoveryModule *IEmailRecoveryModuleCallerSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _IEmailRecoveryModule.Contract.IsModuleType(&_IEmailRecoveryModule.CallOpts, moduleTypeId)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleTransactor) OnInstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IEmailRecoveryModule.contract.Transact(opts, "onInstall", data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.OnInstall(&_IEmailRecoveryModule.TransactOpts, data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleTransactorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.OnInstall(&_IEmailRecoveryModule.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleTransactor) OnUninstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IEmailRecoveryModule.contract.Transact(opts, "onUninstall", data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.OnUninstall(&_IEmailRecoveryModule.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleTransactorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.OnUninstall(&_IEmailRecoveryModule.TransactOpts, data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address _account, bytes _data) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleTransactor) Recover(opts *bind.TransactOpts, _account common.Address, _data []byte) (*types.Transaction, error) {
	return _IEmailRecoveryModule.contract.Transact(opts, "recover", _account, _data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address _account, bytes _data) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleSession) Recover(_account common.Address, _data []byte) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.Recover(&_IEmailRecoveryModule.TransactOpts, _account, _data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address _account, bytes _data) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleTransactorSession) Recover(_account common.Address, _data []byte) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.Recover(&_IEmailRecoveryModule.TransactOpts, _account, _data)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _newOracle) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleTransactor) SetOracle(opts *bind.TransactOpts, _newOracle common.Address) (*types.Transaction, error) {
	return _IEmailRecoveryModule.contract.Transact(opts, "setOracle", _newOracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _newOracle) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleSession) SetOracle(_newOracle common.Address) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.SetOracle(&_IEmailRecoveryModule.TransactOpts, _newOracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _newOracle) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleTransactorSession) SetOracle(_newOracle common.Address) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.SetOracle(&_IEmailRecoveryModule.TransactOpts, _newOracle)
}

// UpdateEmailHash is a paid mutator transaction binding the contract method 0x29a27b1d.
//
// Solidity: function updateEmailHash(bytes32 _emailHash) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleTransactor) UpdateEmailHash(opts *bind.TransactOpts, _emailHash [32]byte) (*types.Transaction, error) {
	return _IEmailRecoveryModule.contract.Transact(opts, "updateEmailHash", _emailHash)
}

// UpdateEmailHash is a paid mutator transaction binding the contract method 0x29a27b1d.
//
// Solidity: function updateEmailHash(bytes32 _emailHash) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleSession) UpdateEmailHash(_emailHash [32]byte) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.UpdateEmailHash(&_IEmailRecoveryModule.TransactOpts, _emailHash)
}

// UpdateEmailHash is a paid mutator transaction binding the contract method 0x29a27b1d.
//
// Solidity: function updateEmailHash(bytes32 _emailHash) returns()
func (_IEmailRecoveryModule *IEmailRecoveryModuleTransactorSession) UpdateEmailHash(_emailHash [32]byte) (*types.Transaction, error) {
	return _IEmailRecoveryModule.Contract.UpdateEmailHash(&_IEmailRecoveryModule.TransactOpts, _emailHash)
}

// IEmailRecoveryModuleEmailHashUpdatedIterator is returned from FilterEmailHashUpdated and is used to iterate over the raw logs and unpacked data for EmailHashUpdated events raised by the IEmailRecoveryModule contract.
type IEmailRecoveryModuleEmailHashUpdatedIterator struct {
	Event *IEmailRecoveryModuleEmailHashUpdated // Event containing the contract specifics and raw log

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
func (it *IEmailRecoveryModuleEmailHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEmailRecoveryModuleEmailHashUpdated)
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
		it.Event = new(IEmailRecoveryModuleEmailHashUpdated)
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
func (it *IEmailRecoveryModuleEmailHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEmailRecoveryModuleEmailHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEmailRecoveryModuleEmailHashUpdated represents a EmailHashUpdated event raised by the IEmailRecoveryModule contract.
type IEmailRecoveryModuleEmailHashUpdated struct {
	Account   common.Address
	EmailHash [32]byte
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEmailHashUpdated is a free log retrieval operation binding the contract event 0xfc50c12be028a09bd30ac0b3f470f1a81c4d72e1043d9358931da4dbf162b401.
//
// Solidity: event EmailHashUpdated(address account, bytes32 emailHash, uint256 updatedAt)
func (_IEmailRecoveryModule *IEmailRecoveryModuleFilterer) FilterEmailHashUpdated(opts *bind.FilterOpts) (*IEmailRecoveryModuleEmailHashUpdatedIterator, error) {

	logs, sub, err := _IEmailRecoveryModule.contract.FilterLogs(opts, "EmailHashUpdated")
	if err != nil {
		return nil, err
	}
	return &IEmailRecoveryModuleEmailHashUpdatedIterator{contract: _IEmailRecoveryModule.contract, event: "EmailHashUpdated", logs: logs, sub: sub}, nil
}

// WatchEmailHashUpdated is a free log subscription operation binding the contract event 0xfc50c12be028a09bd30ac0b3f470f1a81c4d72e1043d9358931da4dbf162b401.
//
// Solidity: event EmailHashUpdated(address account, bytes32 emailHash, uint256 updatedAt)
func (_IEmailRecoveryModule *IEmailRecoveryModuleFilterer) WatchEmailHashUpdated(opts *bind.WatchOpts, sink chan<- *IEmailRecoveryModuleEmailHashUpdated) (event.Subscription, error) {

	logs, sub, err := _IEmailRecoveryModule.contract.WatchLogs(opts, "EmailHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEmailRecoveryModuleEmailHashUpdated)
				if err := _IEmailRecoveryModule.contract.UnpackLog(event, "EmailHashUpdated", log); err != nil {
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

// ParseEmailHashUpdated is a log parse operation binding the contract event 0xfc50c12be028a09bd30ac0b3f470f1a81c4d72e1043d9358931da4dbf162b401.
//
// Solidity: event EmailHashUpdated(address account, bytes32 emailHash, uint256 updatedAt)
func (_IEmailRecoveryModule *IEmailRecoveryModuleFilterer) ParseEmailHashUpdated(log types.Log) (*IEmailRecoveryModuleEmailHashUpdated, error) {
	event := new(IEmailRecoveryModuleEmailHashUpdated)
	if err := _IEmailRecoveryModule.contract.UnpackLog(event, "EmailHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEmailRecoveryModuleOracleUpdatedIterator is returned from FilterOracleUpdated and is used to iterate over the raw logs and unpacked data for OracleUpdated events raised by the IEmailRecoveryModule contract.
type IEmailRecoveryModuleOracleUpdatedIterator struct {
	Event *IEmailRecoveryModuleOracleUpdated // Event containing the contract specifics and raw log

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
func (it *IEmailRecoveryModuleOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEmailRecoveryModuleOracleUpdated)
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
		it.Event = new(IEmailRecoveryModuleOracleUpdated)
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
func (it *IEmailRecoveryModuleOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEmailRecoveryModuleOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEmailRecoveryModuleOracleUpdated represents a OracleUpdated event raised by the IEmailRecoveryModule contract.
type IEmailRecoveryModuleOracleUpdated struct {
	NewOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleUpdated is a free log retrieval operation binding the contract event 0x3df77beb5db05fcdd70a30fc8adf3f83f9501b68579455adbd100b8180940394.
//
// Solidity: event OracleUpdated(address newOracle)
func (_IEmailRecoveryModule *IEmailRecoveryModuleFilterer) FilterOracleUpdated(opts *bind.FilterOpts) (*IEmailRecoveryModuleOracleUpdatedIterator, error) {

	logs, sub, err := _IEmailRecoveryModule.contract.FilterLogs(opts, "OracleUpdated")
	if err != nil {
		return nil, err
	}
	return &IEmailRecoveryModuleOracleUpdatedIterator{contract: _IEmailRecoveryModule.contract, event: "OracleUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleUpdated is a free log subscription operation binding the contract event 0x3df77beb5db05fcdd70a30fc8adf3f83f9501b68579455adbd100b8180940394.
//
// Solidity: event OracleUpdated(address newOracle)
func (_IEmailRecoveryModule *IEmailRecoveryModuleFilterer) WatchOracleUpdated(opts *bind.WatchOpts, sink chan<- *IEmailRecoveryModuleOracleUpdated) (event.Subscription, error) {

	logs, sub, err := _IEmailRecoveryModule.contract.WatchLogs(opts, "OracleUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEmailRecoveryModuleOracleUpdated)
				if err := _IEmailRecoveryModule.contract.UnpackLog(event, "OracleUpdated", log); err != nil {
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

// ParseOracleUpdated is a log parse operation binding the contract event 0x3df77beb5db05fcdd70a30fc8adf3f83f9501b68579455adbd100b8180940394.
//
// Solidity: event OracleUpdated(address newOracle)
func (_IEmailRecoveryModule *IEmailRecoveryModuleFilterer) ParseOracleUpdated(log types.Log) (*IEmailRecoveryModuleOracleUpdated, error) {
	event := new(IEmailRecoveryModuleOracleUpdated)
	if err := _IEmailRecoveryModule.contract.UnpackLog(event, "OracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEmailRecoveryModuleSignerUpdatedIterator is returned from FilterSignerUpdated and is used to iterate over the raw logs and unpacked data for SignerUpdated events raised by the IEmailRecoveryModule contract.
type IEmailRecoveryModuleSignerUpdatedIterator struct {
	Event *IEmailRecoveryModuleSignerUpdated // Event containing the contract specifics and raw log

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
func (it *IEmailRecoveryModuleSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEmailRecoveryModuleSignerUpdated)
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
		it.Event = new(IEmailRecoveryModuleSignerUpdated)
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
func (it *IEmailRecoveryModuleSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEmailRecoveryModuleSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEmailRecoveryModuleSignerUpdated represents a SignerUpdated event raised by the IEmailRecoveryModule contract.
type IEmailRecoveryModuleSignerUpdated struct {
	NewSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignerUpdated is a free log retrieval operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address newSigner)
func (_IEmailRecoveryModule *IEmailRecoveryModuleFilterer) FilterSignerUpdated(opts *bind.FilterOpts) (*IEmailRecoveryModuleSignerUpdatedIterator, error) {

	logs, sub, err := _IEmailRecoveryModule.contract.FilterLogs(opts, "SignerUpdated")
	if err != nil {
		return nil, err
	}
	return &IEmailRecoveryModuleSignerUpdatedIterator{contract: _IEmailRecoveryModule.contract, event: "SignerUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerUpdated is a free log subscription operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address newSigner)
func (_IEmailRecoveryModule *IEmailRecoveryModuleFilterer) WatchSignerUpdated(opts *bind.WatchOpts, sink chan<- *IEmailRecoveryModuleSignerUpdated) (event.Subscription, error) {

	logs, sub, err := _IEmailRecoveryModule.contract.WatchLogs(opts, "SignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEmailRecoveryModuleSignerUpdated)
				if err := _IEmailRecoveryModule.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
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

// ParseSignerUpdated is a log parse operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address newSigner)
func (_IEmailRecoveryModule *IEmailRecoveryModuleFilterer) ParseSignerUpdated(log types.Log) (*IEmailRecoveryModuleSignerUpdated, error) {
	event := new(IEmailRecoveryModuleSignerUpdated)
	if err := _IEmailRecoveryModule.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
