// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package idkimkeyoracle

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

// IDkimKeyOracleMetaData contains all meta data concerning the IDkimKeyOracle contract.
var IDkimKeyOracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"DkimKeyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"}],\"name\":\"DkimKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"PauserUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"}],\"name\":\"getDomainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"}],\"name\":\"removeKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_domain\",\"type\":\"bytes32\"}],\"name\":\"updateKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IDkimKeyOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use IDkimKeyOracleMetaData.ABI instead.
var IDkimKeyOracleABI = IDkimKeyOracleMetaData.ABI

// IDkimKeyOracle is an auto generated Go binding around an Ethereum contract.
type IDkimKeyOracle struct {
	IDkimKeyOracleCaller     // Read-only binding to the contract
	IDkimKeyOracleTransactor // Write-only binding to the contract
	IDkimKeyOracleFilterer   // Log filterer for contract events
}

// IDkimKeyOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDkimKeyOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDkimKeyOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDkimKeyOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDkimKeyOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDkimKeyOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDkimKeyOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDkimKeyOracleSession struct {
	Contract     *IDkimKeyOracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDkimKeyOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDkimKeyOracleCallerSession struct {
	Contract *IDkimKeyOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IDkimKeyOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDkimKeyOracleTransactorSession struct {
	Contract     *IDkimKeyOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IDkimKeyOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDkimKeyOracleRaw struct {
	Contract *IDkimKeyOracle // Generic contract binding to access the raw methods on
}

// IDkimKeyOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDkimKeyOracleCallerRaw struct {
	Contract *IDkimKeyOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IDkimKeyOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDkimKeyOracleTransactorRaw struct {
	Contract *IDkimKeyOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDkimKeyOracle creates a new instance of IDkimKeyOracle, bound to a specific deployed contract.
func NewIDkimKeyOracle(address common.Address, backend bind.ContractBackend) (*IDkimKeyOracle, error) {
	contract, err := bindIDkimKeyOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDkimKeyOracle{IDkimKeyOracleCaller: IDkimKeyOracleCaller{contract: contract}, IDkimKeyOracleTransactor: IDkimKeyOracleTransactor{contract: contract}, IDkimKeyOracleFilterer: IDkimKeyOracleFilterer{contract: contract}}, nil
}

// NewIDkimKeyOracleCaller creates a new read-only instance of IDkimKeyOracle, bound to a specific deployed contract.
func NewIDkimKeyOracleCaller(address common.Address, caller bind.ContractCaller) (*IDkimKeyOracleCaller, error) {
	contract, err := bindIDkimKeyOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDkimKeyOracleCaller{contract: contract}, nil
}

// NewIDkimKeyOracleTransactor creates a new write-only instance of IDkimKeyOracle, bound to a specific deployed contract.
func NewIDkimKeyOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IDkimKeyOracleTransactor, error) {
	contract, err := bindIDkimKeyOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDkimKeyOracleTransactor{contract: contract}, nil
}

// NewIDkimKeyOracleFilterer creates a new log filterer instance of IDkimKeyOracle, bound to a specific deployed contract.
func NewIDkimKeyOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IDkimKeyOracleFilterer, error) {
	contract, err := bindIDkimKeyOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDkimKeyOracleFilterer{contract: contract}, nil
}

// bindIDkimKeyOracle binds a generic wrapper to an already deployed contract.
func bindIDkimKeyOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDkimKeyOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDkimKeyOracle *IDkimKeyOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDkimKeyOracle.Contract.IDkimKeyOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDkimKeyOracle *IDkimKeyOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDkimKeyOracle.Contract.IDkimKeyOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDkimKeyOracle *IDkimKeyOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDkimKeyOracle.Contract.IDkimKeyOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDkimKeyOracle *IDkimKeyOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDkimKeyOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDkimKeyOracle *IDkimKeyOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDkimKeyOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDkimKeyOracle *IDkimKeyOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDkimKeyOracle.Contract.contract.Transact(opts, method, params...)
}

// GetDomainHash is a free data retrieval call binding the contract method 0x368a1b76.
//
// Solidity: function getDomainHash(bytes32 _keyHash) view returns(bytes32)
func (_IDkimKeyOracle *IDkimKeyOracleCaller) GetDomainHash(opts *bind.CallOpts, _keyHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IDkimKeyOracle.contract.Call(opts, &out, "getDomainHash", _keyHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainHash is a free data retrieval call binding the contract method 0x368a1b76.
//
// Solidity: function getDomainHash(bytes32 _keyHash) view returns(bytes32)
func (_IDkimKeyOracle *IDkimKeyOracleSession) GetDomainHash(_keyHash [32]byte) ([32]byte, error) {
	return _IDkimKeyOracle.Contract.GetDomainHash(&_IDkimKeyOracle.CallOpts, _keyHash)
}

// GetDomainHash is a free data retrieval call binding the contract method 0x368a1b76.
//
// Solidity: function getDomainHash(bytes32 _keyHash) view returns(bytes32)
func (_IDkimKeyOracle *IDkimKeyOracleCallerSession) GetDomainHash(_keyHash [32]byte) ([32]byte, error) {
	return _IDkimKeyOracle.Contract.GetDomainHash(&_IDkimKeyOracle.CallOpts, _keyHash)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x862642f5.
//
// Solidity: function removeKey(bytes32 _keyHash) returns()
func (_IDkimKeyOracle *IDkimKeyOracleTransactor) RemoveKey(opts *bind.TransactOpts, _keyHash [32]byte) (*types.Transaction, error) {
	return _IDkimKeyOracle.contract.Transact(opts, "removeKey", _keyHash)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x862642f5.
//
// Solidity: function removeKey(bytes32 _keyHash) returns()
func (_IDkimKeyOracle *IDkimKeyOracleSession) RemoveKey(_keyHash [32]byte) (*types.Transaction, error) {
	return _IDkimKeyOracle.Contract.RemoveKey(&_IDkimKeyOracle.TransactOpts, _keyHash)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x862642f5.
//
// Solidity: function removeKey(bytes32 _keyHash) returns()
func (_IDkimKeyOracle *IDkimKeyOracleTransactorSession) RemoveKey(_keyHash [32]byte) (*types.Transaction, error) {
	return _IDkimKeyOracle.Contract.RemoveKey(&_IDkimKeyOracle.TransactOpts, _keyHash)
}

// UpdateKey is a paid mutator transaction binding the contract method 0x73af6745.
//
// Solidity: function updateKey(bytes32 _keyHash, bytes32 _domain) returns()
func (_IDkimKeyOracle *IDkimKeyOracleTransactor) UpdateKey(opts *bind.TransactOpts, _keyHash [32]byte, _domain [32]byte) (*types.Transaction, error) {
	return _IDkimKeyOracle.contract.Transact(opts, "updateKey", _keyHash, _domain)
}

// UpdateKey is a paid mutator transaction binding the contract method 0x73af6745.
//
// Solidity: function updateKey(bytes32 _keyHash, bytes32 _domain) returns()
func (_IDkimKeyOracle *IDkimKeyOracleSession) UpdateKey(_keyHash [32]byte, _domain [32]byte) (*types.Transaction, error) {
	return _IDkimKeyOracle.Contract.UpdateKey(&_IDkimKeyOracle.TransactOpts, _keyHash, _domain)
}

// UpdateKey is a paid mutator transaction binding the contract method 0x73af6745.
//
// Solidity: function updateKey(bytes32 _keyHash, bytes32 _domain) returns()
func (_IDkimKeyOracle *IDkimKeyOracleTransactorSession) UpdateKey(_keyHash [32]byte, _domain [32]byte) (*types.Transaction, error) {
	return _IDkimKeyOracle.Contract.UpdateKey(&_IDkimKeyOracle.TransactOpts, _keyHash, _domain)
}

// IDkimKeyOracleDkimKeyRemovedIterator is returned from FilterDkimKeyRemoved and is used to iterate over the raw logs and unpacked data for DkimKeyRemoved events raised by the IDkimKeyOracle contract.
type IDkimKeyOracleDkimKeyRemovedIterator struct {
	Event *IDkimKeyOracleDkimKeyRemoved // Event containing the contract specifics and raw log

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
func (it *IDkimKeyOracleDkimKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDkimKeyOracleDkimKeyRemoved)
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
		it.Event = new(IDkimKeyOracleDkimKeyRemoved)
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
func (it *IDkimKeyOracleDkimKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDkimKeyOracleDkimKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDkimKeyOracleDkimKeyRemoved represents a DkimKeyRemoved event raised by the IDkimKeyOracle contract.
type IDkimKeyOracleDkimKeyRemoved struct {
	KeyHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDkimKeyRemoved is a free log retrieval operation binding the contract event 0xa54c86c49f4e97ce7b0c5af2e11859918d86f47f03793f03f33db5b7c349d0d1.
//
// Solidity: event DkimKeyRemoved(bytes32 keyHash)
func (_IDkimKeyOracle *IDkimKeyOracleFilterer) FilterDkimKeyRemoved(opts *bind.FilterOpts) (*IDkimKeyOracleDkimKeyRemovedIterator, error) {

	logs, sub, err := _IDkimKeyOracle.contract.FilterLogs(opts, "DkimKeyRemoved")
	if err != nil {
		return nil, err
	}
	return &IDkimKeyOracleDkimKeyRemovedIterator{contract: _IDkimKeyOracle.contract, event: "DkimKeyRemoved", logs: logs, sub: sub}, nil
}

// WatchDkimKeyRemoved is a free log subscription operation binding the contract event 0xa54c86c49f4e97ce7b0c5af2e11859918d86f47f03793f03f33db5b7c349d0d1.
//
// Solidity: event DkimKeyRemoved(bytes32 keyHash)
func (_IDkimKeyOracle *IDkimKeyOracleFilterer) WatchDkimKeyRemoved(opts *bind.WatchOpts, sink chan<- *IDkimKeyOracleDkimKeyRemoved) (event.Subscription, error) {

	logs, sub, err := _IDkimKeyOracle.contract.WatchLogs(opts, "DkimKeyRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDkimKeyOracleDkimKeyRemoved)
				if err := _IDkimKeyOracle.contract.UnpackLog(event, "DkimKeyRemoved", log); err != nil {
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

// ParseDkimKeyRemoved is a log parse operation binding the contract event 0xa54c86c49f4e97ce7b0c5af2e11859918d86f47f03793f03f33db5b7c349d0d1.
//
// Solidity: event DkimKeyRemoved(bytes32 keyHash)
func (_IDkimKeyOracle *IDkimKeyOracleFilterer) ParseDkimKeyRemoved(log types.Log) (*IDkimKeyOracleDkimKeyRemoved, error) {
	event := new(IDkimKeyOracleDkimKeyRemoved)
	if err := _IDkimKeyOracle.contract.UnpackLog(event, "DkimKeyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDkimKeyOracleDkimKeyUpdatedIterator is returned from FilterDkimKeyUpdated and is used to iterate over the raw logs and unpacked data for DkimKeyUpdated events raised by the IDkimKeyOracle contract.
type IDkimKeyOracleDkimKeyUpdatedIterator struct {
	Event *IDkimKeyOracleDkimKeyUpdated // Event containing the contract specifics and raw log

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
func (it *IDkimKeyOracleDkimKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDkimKeyOracleDkimKeyUpdated)
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
		it.Event = new(IDkimKeyOracleDkimKeyUpdated)
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
func (it *IDkimKeyOracleDkimKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDkimKeyOracleDkimKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDkimKeyOracleDkimKeyUpdated represents a DkimKeyUpdated event raised by the IDkimKeyOracle contract.
type IDkimKeyOracleDkimKeyUpdated struct {
	KeyHash    [32]byte
	DomainHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDkimKeyUpdated is a free log retrieval operation binding the contract event 0xe780335b851f95261f942f1897645485db8bd8c5ff0fc4dff7b5cc2de72d2d46.
//
// Solidity: event DkimKeyUpdated(bytes32 keyHash, bytes32 domainHash)
func (_IDkimKeyOracle *IDkimKeyOracleFilterer) FilterDkimKeyUpdated(opts *bind.FilterOpts) (*IDkimKeyOracleDkimKeyUpdatedIterator, error) {

	logs, sub, err := _IDkimKeyOracle.contract.FilterLogs(opts, "DkimKeyUpdated")
	if err != nil {
		return nil, err
	}
	return &IDkimKeyOracleDkimKeyUpdatedIterator{contract: _IDkimKeyOracle.contract, event: "DkimKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchDkimKeyUpdated is a free log subscription operation binding the contract event 0xe780335b851f95261f942f1897645485db8bd8c5ff0fc4dff7b5cc2de72d2d46.
//
// Solidity: event DkimKeyUpdated(bytes32 keyHash, bytes32 domainHash)
func (_IDkimKeyOracle *IDkimKeyOracleFilterer) WatchDkimKeyUpdated(opts *bind.WatchOpts, sink chan<- *IDkimKeyOracleDkimKeyUpdated) (event.Subscription, error) {

	logs, sub, err := _IDkimKeyOracle.contract.WatchLogs(opts, "DkimKeyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDkimKeyOracleDkimKeyUpdated)
				if err := _IDkimKeyOracle.contract.UnpackLog(event, "DkimKeyUpdated", log); err != nil {
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

// ParseDkimKeyUpdated is a log parse operation binding the contract event 0xe780335b851f95261f942f1897645485db8bd8c5ff0fc4dff7b5cc2de72d2d46.
//
// Solidity: event DkimKeyUpdated(bytes32 keyHash, bytes32 domainHash)
func (_IDkimKeyOracle *IDkimKeyOracleFilterer) ParseDkimKeyUpdated(log types.Log) (*IDkimKeyOracleDkimKeyUpdated, error) {
	event := new(IDkimKeyOracleDkimKeyUpdated)
	if err := _IDkimKeyOracle.contract.UnpackLog(event, "DkimKeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDkimKeyOraclePauserUpdatedIterator is returned from FilterPauserUpdated and is used to iterate over the raw logs and unpacked data for PauserUpdated events raised by the IDkimKeyOracle contract.
type IDkimKeyOraclePauserUpdatedIterator struct {
	Event *IDkimKeyOraclePauserUpdated // Event containing the contract specifics and raw log

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
func (it *IDkimKeyOraclePauserUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDkimKeyOraclePauserUpdated)
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
		it.Event = new(IDkimKeyOraclePauserUpdated)
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
func (it *IDkimKeyOraclePauserUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDkimKeyOraclePauserUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDkimKeyOraclePauserUpdated represents a PauserUpdated event raised by the IDkimKeyOracle contract.
type IDkimKeyOraclePauserUpdated struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPauserUpdated is a free log retrieval operation binding the contract event 0xa4336c0cb1e245b95ad204faed7e940d6dc999684fd8b5e1ff597a0c4efca8ab.
//
// Solidity: event PauserUpdated(address pauser)
func (_IDkimKeyOracle *IDkimKeyOracleFilterer) FilterPauserUpdated(opts *bind.FilterOpts) (*IDkimKeyOraclePauserUpdatedIterator, error) {

	logs, sub, err := _IDkimKeyOracle.contract.FilterLogs(opts, "PauserUpdated")
	if err != nil {
		return nil, err
	}
	return &IDkimKeyOraclePauserUpdatedIterator{contract: _IDkimKeyOracle.contract, event: "PauserUpdated", logs: logs, sub: sub}, nil
}

// WatchPauserUpdated is a free log subscription operation binding the contract event 0xa4336c0cb1e245b95ad204faed7e940d6dc999684fd8b5e1ff597a0c4efca8ab.
//
// Solidity: event PauserUpdated(address pauser)
func (_IDkimKeyOracle *IDkimKeyOracleFilterer) WatchPauserUpdated(opts *bind.WatchOpts, sink chan<- *IDkimKeyOraclePauserUpdated) (event.Subscription, error) {

	logs, sub, err := _IDkimKeyOracle.contract.WatchLogs(opts, "PauserUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDkimKeyOraclePauserUpdated)
				if err := _IDkimKeyOracle.contract.UnpackLog(event, "PauserUpdated", log); err != nil {
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

// ParsePauserUpdated is a log parse operation binding the contract event 0xa4336c0cb1e245b95ad204faed7e940d6dc999684fd8b5e1ff597a0c4efca8ab.
//
// Solidity: event PauserUpdated(address pauser)
func (_IDkimKeyOracle *IDkimKeyOracleFilterer) ParsePauserUpdated(log types.Log) (*IDkimKeyOraclePauserUpdated, error) {
	event := new(IDkimKeyOraclePauserUpdated)
	if err := _IDkimKeyOracle.contract.UnpackLog(event, "PauserUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
