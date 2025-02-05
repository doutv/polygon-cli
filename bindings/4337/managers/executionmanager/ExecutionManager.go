// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executionmanager

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

// ExecutionManagerMetaData contains all meta data concerning the ExecutionManager contract.
var ExecutionManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchExecutionindex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"TryExecuteUnsuccessful\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"executors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601357609f908160198239f35b600080fdfe6080806040526004361015601257600080fd5b600090813560e01c639ac2a01114602857600080fd5b3460655760203660031901126065576004356001600160a01b0381169081900360615760408360ff926020955280855220541615158152f35b8280fd5b5080fdfea26469706673582212206cee7e0bb4fd60e43009a41b66d2fdc01b280474d7edf2ab352e412ecbdcbaff64736f6c63430008190033",
}

// ExecutionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutionManagerMetaData.ABI instead.
var ExecutionManagerABI = ExecutionManagerMetaData.ABI

// ExecutionManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExecutionManagerMetaData.Bin instead.
var ExecutionManagerBin = ExecutionManagerMetaData.Bin

// DeployExecutionManager deploys a new Ethereum contract, binding an instance of ExecutionManager to it.
func DeployExecutionManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExecutionManager, error) {
	parsed, err := ExecutionManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutionManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExecutionManager{ExecutionManagerCaller: ExecutionManagerCaller{contract: contract}, ExecutionManagerTransactor: ExecutionManagerTransactor{contract: contract}, ExecutionManagerFilterer: ExecutionManagerFilterer{contract: contract}}, nil
}

// ExecutionManager is an auto generated Go binding around an Ethereum contract.
type ExecutionManager struct {
	ExecutionManagerCaller     // Read-only binding to the contract
	ExecutionManagerTransactor // Write-only binding to the contract
	ExecutionManagerFilterer   // Log filterer for contract events
}

// ExecutionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutionManagerSession struct {
	Contract     *ExecutionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecutionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutionManagerCallerSession struct {
	Contract *ExecutionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ExecutionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutionManagerTransactorSession struct {
	Contract     *ExecutionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExecutionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutionManagerRaw struct {
	Contract *ExecutionManager // Generic contract binding to access the raw methods on
}

// ExecutionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutionManagerCallerRaw struct {
	Contract *ExecutionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutionManagerTransactorRaw struct {
	Contract *ExecutionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutionManager creates a new instance of ExecutionManager, bound to a specific deployed contract.
func NewExecutionManager(address common.Address, backend bind.ContractBackend) (*ExecutionManager, error) {
	contract, err := bindExecutionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionManager{ExecutionManagerCaller: ExecutionManagerCaller{contract: contract}, ExecutionManagerTransactor: ExecutionManagerTransactor{contract: contract}, ExecutionManagerFilterer: ExecutionManagerFilterer{contract: contract}}, nil
}

// NewExecutionManagerCaller creates a new read-only instance of ExecutionManager, bound to a specific deployed contract.
func NewExecutionManagerCaller(address common.Address, caller bind.ContractCaller) (*ExecutionManagerCaller, error) {
	contract, err := bindExecutionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionManagerCaller{contract: contract}, nil
}

// NewExecutionManagerTransactor creates a new write-only instance of ExecutionManager, bound to a specific deployed contract.
func NewExecutionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutionManagerTransactor, error) {
	contract, err := bindExecutionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionManagerTransactor{contract: contract}, nil
}

// NewExecutionManagerFilterer creates a new log filterer instance of ExecutionManager, bound to a specific deployed contract.
func NewExecutionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutionManagerFilterer, error) {
	contract, err := bindExecutionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutionManagerFilterer{contract: contract}, nil
}

// bindExecutionManager binds a generic wrapper to an already deployed contract.
func bindExecutionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExecutionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionManager *ExecutionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionManager.Contract.ExecutionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionManager *ExecutionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionManager.Contract.ExecutionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionManager *ExecutionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionManager.Contract.ExecutionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionManager *ExecutionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionManager *ExecutionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionManager *ExecutionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionManager.Contract.contract.Transact(opts, method, params...)
}

// Executors is a free data retrieval call binding the contract method 0x9ac2a011.
//
// Solidity: function executors(address ) view returns(bool)
func (_ExecutionManager *ExecutionManagerCaller) Executors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ExecutionManager.contract.Call(opts, &out, "executors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Executors is a free data retrieval call binding the contract method 0x9ac2a011.
//
// Solidity: function executors(address ) view returns(bool)
func (_ExecutionManager *ExecutionManagerSession) Executors(arg0 common.Address) (bool, error) {
	return _ExecutionManager.Contract.Executors(&_ExecutionManager.CallOpts, arg0)
}

// Executors is a free data retrieval call binding the contract method 0x9ac2a011.
//
// Solidity: function executors(address ) view returns(bool)
func (_ExecutionManager *ExecutionManagerCallerSession) Executors(arg0 common.Address) (bool, error) {
	return _ExecutionManager.Contract.Executors(&_ExecutionManager.CallOpts, arg0)
}

// ExecutionManagerTryExecuteUnsuccessfulIterator is returned from FilterTryExecuteUnsuccessful and is used to iterate over the raw logs and unpacked data for TryExecuteUnsuccessful events raised by the ExecutionManager contract.
type ExecutionManagerTryExecuteUnsuccessfulIterator struct {
	Event *ExecutionManagerTryExecuteUnsuccessful // Event containing the contract specifics and raw log

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
func (it *ExecutionManagerTryExecuteUnsuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionManagerTryExecuteUnsuccessful)
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
		it.Event = new(ExecutionManagerTryExecuteUnsuccessful)
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
func (it *ExecutionManagerTryExecuteUnsuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionManagerTryExecuteUnsuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionManagerTryExecuteUnsuccessful represents a TryExecuteUnsuccessful event raised by the ExecutionManager contract.
type ExecutionManagerTryExecuteUnsuccessful struct {
	BatchExecutionindex *big.Int
	Result              []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTryExecuteUnsuccessful is a free log retrieval operation binding the contract event 0xe723f28f104e46b47fd3531f3608374ac226bcf3ddda334a23a266453e0efdb7.
//
// Solidity: event TryExecuteUnsuccessful(uint256 batchExecutionindex, bytes result)
func (_ExecutionManager *ExecutionManagerFilterer) FilterTryExecuteUnsuccessful(opts *bind.FilterOpts) (*ExecutionManagerTryExecuteUnsuccessfulIterator, error) {

	logs, sub, err := _ExecutionManager.contract.FilterLogs(opts, "TryExecuteUnsuccessful")
	if err != nil {
		return nil, err
	}
	return &ExecutionManagerTryExecuteUnsuccessfulIterator{contract: _ExecutionManager.contract, event: "TryExecuteUnsuccessful", logs: logs, sub: sub}, nil
}

// WatchTryExecuteUnsuccessful is a free log subscription operation binding the contract event 0xe723f28f104e46b47fd3531f3608374ac226bcf3ddda334a23a266453e0efdb7.
//
// Solidity: event TryExecuteUnsuccessful(uint256 batchExecutionindex, bytes result)
func (_ExecutionManager *ExecutionManagerFilterer) WatchTryExecuteUnsuccessful(opts *bind.WatchOpts, sink chan<- *ExecutionManagerTryExecuteUnsuccessful) (event.Subscription, error) {

	logs, sub, err := _ExecutionManager.contract.WatchLogs(opts, "TryExecuteUnsuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionManagerTryExecuteUnsuccessful)
				if err := _ExecutionManager.contract.UnpackLog(event, "TryExecuteUnsuccessful", log); err != nil {
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

// ParseTryExecuteUnsuccessful is a log parse operation binding the contract event 0xe723f28f104e46b47fd3531f3608374ac226bcf3ddda334a23a266453e0efdb7.
//
// Solidity: event TryExecuteUnsuccessful(uint256 batchExecutionindex, bytes result)
func (_ExecutionManager *ExecutionManagerFilterer) ParseTryExecuteUnsuccessful(log types.Log) (*ExecutionManagerTryExecuteUnsuccessful, error) {
	event := new(ExecutionManagerTryExecuteUnsuccessful)
	if err := _ExecutionManager.contract.UnpackLog(event, "TryExecuteUnsuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
