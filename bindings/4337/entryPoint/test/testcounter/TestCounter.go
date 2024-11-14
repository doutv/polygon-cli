// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testcounter

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

// TestCounterMetaData contains all meta data concerning the TestCounter contract.
var TestCounterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"CalledFrom\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countFail\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"counters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repeat\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"gasWaster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"justemit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"xxx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608080604052346015576102ad908161001b8239f35b600080fdfe6040608081526004908136101561001557600080fd5b600091823560e01c90816306661abd14610207578163278ddd3c146101c9578163a1b468901461013e578163a5e9585f14610116578163be65ab8c146100d9578163caece69314610091575063d55565441461007057600080fd5b3461008d578160031936011261008d576020906002549051908152f35b5080fd5b905082346100d657806003193601126100d657506020606492519162461bcd60e51b8352820152600c60248201526b18dbdd5b9d0819985a5b195960a21b6044820152fd5b80fd5b90503461011257602036600319011261011257356001600160a01b03811690819003610112578282916020945280845220549051908152f35b8280fd5b9050346101125760203660031901126101125760209282913581526001845220549051908152f35b919050346101125780600319360112610112578135916024359067ffffffffffffffff908183116101c557366023840112156101c5578201359081116101c15736910160240111610112576001805b83811115610199578480f35b6101bc9060026101a98154610252565b8091558652826020528084872055610252565b61018d565b8480fd5b8580fd5b50503461008d578160031936011261008d5760207ffb3b4d6258432a9a3d78dd9bffbcb6cfb1bd94f58da35fd530d08da7d1d058329151338152a180f35b919050346101125782600319360112610112573383528260205280832054916001830180931161023f57503383528260205282205580f35b634e487b7160e01b845260119052602483fd5b60001981146102615760010190565b634e487b7160e01b600052601160045260246000fdfea26469706673582212201d17e9b110de1152a0c188beed89acbabc8d3fd30f9adef3bb4d10412f199c0f64736f6c63430008190033",
}

// TestCounterABI is the input ABI used to generate the binding from.
// Deprecated: Use TestCounterMetaData.ABI instead.
var TestCounterABI = TestCounterMetaData.ABI

// TestCounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestCounterMetaData.Bin instead.
var TestCounterBin = TestCounterMetaData.Bin

// DeployTestCounter deploys a new Ethereum contract, binding an instance of TestCounter to it.
func DeployTestCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestCounter, error) {
	parsed, err := TestCounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestCounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestCounter{TestCounterCaller: TestCounterCaller{contract: contract}, TestCounterTransactor: TestCounterTransactor{contract: contract}, TestCounterFilterer: TestCounterFilterer{contract: contract}}, nil
}

// TestCounter is an auto generated Go binding around an Ethereum contract.
type TestCounter struct {
	TestCounterCaller     // Read-only binding to the contract
	TestCounterTransactor // Write-only binding to the contract
	TestCounterFilterer   // Log filterer for contract events
}

// TestCounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestCounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestCounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestCounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestCounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestCounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestCounterSession struct {
	Contract     *TestCounter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCounterCallerSession struct {
	Contract *TestCounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestCounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestCounterTransactorSession struct {
	Contract     *TestCounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestCounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestCounterRaw struct {
	Contract *TestCounter // Generic contract binding to access the raw methods on
}

// TestCounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCounterCallerRaw struct {
	Contract *TestCounterCaller // Generic read-only contract binding to access the raw methods on
}

// TestCounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestCounterTransactorRaw struct {
	Contract *TestCounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestCounter creates a new instance of TestCounter, bound to a specific deployed contract.
func NewTestCounter(address common.Address, backend bind.ContractBackend) (*TestCounter, error) {
	contract, err := bindTestCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestCounter{TestCounterCaller: TestCounterCaller{contract: contract}, TestCounterTransactor: TestCounterTransactor{contract: contract}, TestCounterFilterer: TestCounterFilterer{contract: contract}}, nil
}

// NewTestCounterCaller creates a new read-only instance of TestCounter, bound to a specific deployed contract.
func NewTestCounterCaller(address common.Address, caller bind.ContractCaller) (*TestCounterCaller, error) {
	contract, err := bindTestCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCounterCaller{contract: contract}, nil
}

// NewTestCounterTransactor creates a new write-only instance of TestCounter, bound to a specific deployed contract.
func NewTestCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*TestCounterTransactor, error) {
	contract, err := bindTestCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestCounterTransactor{contract: contract}, nil
}

// NewTestCounterFilterer creates a new log filterer instance of TestCounter, bound to a specific deployed contract.
func NewTestCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*TestCounterFilterer, error) {
	contract, err := bindTestCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestCounterFilterer{contract: contract}, nil
}

// bindTestCounter binds a generic wrapper to an already deployed contract.
func bindTestCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestCounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestCounter *TestCounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestCounter.Contract.TestCounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestCounter *TestCounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestCounter.Contract.TestCounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestCounter *TestCounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestCounter.Contract.TestCounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestCounter *TestCounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestCounter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestCounter *TestCounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestCounter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestCounter *TestCounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestCounter.Contract.contract.Transact(opts, method, params...)
}

// CountFail is a free data retrieval call binding the contract method 0xcaece693.
//
// Solidity: function countFail() pure returns()
func (_TestCounter *TestCounterCaller) CountFail(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestCounter.contract.Call(opts, &out, "countFail")

	if err != nil {
		return err
	}

	return err

}

// CountFail is a free data retrieval call binding the contract method 0xcaece693.
//
// Solidity: function countFail() pure returns()
func (_TestCounter *TestCounterSession) CountFail() error {
	return _TestCounter.Contract.CountFail(&_TestCounter.CallOpts)
}

// CountFail is a free data retrieval call binding the contract method 0xcaece693.
//
// Solidity: function countFail() pure returns()
func (_TestCounter *TestCounterCallerSession) CountFail() error {
	return _TestCounter.Contract.CountFail(&_TestCounter.CallOpts)
}

// Counters is a free data retrieval call binding the contract method 0xbe65ab8c.
//
// Solidity: function counters(address ) view returns(uint256)
func (_TestCounter *TestCounterCaller) Counters(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestCounter.contract.Call(opts, &out, "counters", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counters is a free data retrieval call binding the contract method 0xbe65ab8c.
//
// Solidity: function counters(address ) view returns(uint256)
func (_TestCounter *TestCounterSession) Counters(arg0 common.Address) (*big.Int, error) {
	return _TestCounter.Contract.Counters(&_TestCounter.CallOpts, arg0)
}

// Counters is a free data retrieval call binding the contract method 0xbe65ab8c.
//
// Solidity: function counters(address ) view returns(uint256)
func (_TestCounter *TestCounterCallerSession) Counters(arg0 common.Address) (*big.Int, error) {
	return _TestCounter.Contract.Counters(&_TestCounter.CallOpts, arg0)
}

// Offset is a free data retrieval call binding the contract method 0xd5556544.
//
// Solidity: function offset() view returns(uint256)
func (_TestCounter *TestCounterCaller) Offset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestCounter.contract.Call(opts, &out, "offset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Offset is a free data retrieval call binding the contract method 0xd5556544.
//
// Solidity: function offset() view returns(uint256)
func (_TestCounter *TestCounterSession) Offset() (*big.Int, error) {
	return _TestCounter.Contract.Offset(&_TestCounter.CallOpts)
}

// Offset is a free data retrieval call binding the contract method 0xd5556544.
//
// Solidity: function offset() view returns(uint256)
func (_TestCounter *TestCounterCallerSession) Offset() (*big.Int, error) {
	return _TestCounter.Contract.Offset(&_TestCounter.CallOpts)
}

// Xxx is a free data retrieval call binding the contract method 0xa5e9585f.
//
// Solidity: function xxx(uint256 ) view returns(uint256)
func (_TestCounter *TestCounterCaller) Xxx(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestCounter.contract.Call(opts, &out, "xxx", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Xxx is a free data retrieval call binding the contract method 0xa5e9585f.
//
// Solidity: function xxx(uint256 ) view returns(uint256)
func (_TestCounter *TestCounterSession) Xxx(arg0 *big.Int) (*big.Int, error) {
	return _TestCounter.Contract.Xxx(&_TestCounter.CallOpts, arg0)
}

// Xxx is a free data retrieval call binding the contract method 0xa5e9585f.
//
// Solidity: function xxx(uint256 ) view returns(uint256)
func (_TestCounter *TestCounterCallerSession) Xxx(arg0 *big.Int) (*big.Int, error) {
	return _TestCounter.Contract.Xxx(&_TestCounter.CallOpts, arg0)
}

// Count is a paid mutator transaction binding the contract method 0x06661abd.
//
// Solidity: function count() returns()
func (_TestCounter *TestCounterTransactor) Count(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestCounter.contract.Transact(opts, "count")
}

// Count is a paid mutator transaction binding the contract method 0x06661abd.
//
// Solidity: function count() returns()
func (_TestCounter *TestCounterSession) Count() (*types.Transaction, error) {
	return _TestCounter.Contract.Count(&_TestCounter.TransactOpts)
}

// Count is a paid mutator transaction binding the contract method 0x06661abd.
//
// Solidity: function count() returns()
func (_TestCounter *TestCounterTransactorSession) Count() (*types.Transaction, error) {
	return _TestCounter.Contract.Count(&_TestCounter.TransactOpts)
}

// GasWaster is a paid mutator transaction binding the contract method 0xa1b46890.
//
// Solidity: function gasWaster(uint256 repeat, string ) returns()
func (_TestCounter *TestCounterTransactor) GasWaster(opts *bind.TransactOpts, repeat *big.Int, arg1 string) (*types.Transaction, error) {
	return _TestCounter.contract.Transact(opts, "gasWaster", repeat, arg1)
}

// GasWaster is a paid mutator transaction binding the contract method 0xa1b46890.
//
// Solidity: function gasWaster(uint256 repeat, string ) returns()
func (_TestCounter *TestCounterSession) GasWaster(repeat *big.Int, arg1 string) (*types.Transaction, error) {
	return _TestCounter.Contract.GasWaster(&_TestCounter.TransactOpts, repeat, arg1)
}

// GasWaster is a paid mutator transaction binding the contract method 0xa1b46890.
//
// Solidity: function gasWaster(uint256 repeat, string ) returns()
func (_TestCounter *TestCounterTransactorSession) GasWaster(repeat *big.Int, arg1 string) (*types.Transaction, error) {
	return _TestCounter.Contract.GasWaster(&_TestCounter.TransactOpts, repeat, arg1)
}

// Justemit is a paid mutator transaction binding the contract method 0x278ddd3c.
//
// Solidity: function justemit() returns()
func (_TestCounter *TestCounterTransactor) Justemit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestCounter.contract.Transact(opts, "justemit")
}

// Justemit is a paid mutator transaction binding the contract method 0x278ddd3c.
//
// Solidity: function justemit() returns()
func (_TestCounter *TestCounterSession) Justemit() (*types.Transaction, error) {
	return _TestCounter.Contract.Justemit(&_TestCounter.TransactOpts)
}

// Justemit is a paid mutator transaction binding the contract method 0x278ddd3c.
//
// Solidity: function justemit() returns()
func (_TestCounter *TestCounterTransactorSession) Justemit() (*types.Transaction, error) {
	return _TestCounter.Contract.Justemit(&_TestCounter.TransactOpts)
}

// TestCounterCalledFromIterator is returned from FilterCalledFrom and is used to iterate over the raw logs and unpacked data for CalledFrom events raised by the TestCounter contract.
type TestCounterCalledFromIterator struct {
	Event *TestCounterCalledFrom // Event containing the contract specifics and raw log

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
func (it *TestCounterCalledFromIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestCounterCalledFrom)
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
		it.Event = new(TestCounterCalledFrom)
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
func (it *TestCounterCalledFromIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestCounterCalledFromIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestCounterCalledFrom represents a CalledFrom event raised by the TestCounter contract.
type TestCounterCalledFrom struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCalledFrom is a free log retrieval operation binding the contract event 0xfb3b4d6258432a9a3d78dd9bffbcb6cfb1bd94f58da35fd530d08da7d1d05832.
//
// Solidity: event CalledFrom(address sender)
func (_TestCounter *TestCounterFilterer) FilterCalledFrom(opts *bind.FilterOpts) (*TestCounterCalledFromIterator, error) {

	logs, sub, err := _TestCounter.contract.FilterLogs(opts, "CalledFrom")
	if err != nil {
		return nil, err
	}
	return &TestCounterCalledFromIterator{contract: _TestCounter.contract, event: "CalledFrom", logs: logs, sub: sub}, nil
}

// WatchCalledFrom is a free log subscription operation binding the contract event 0xfb3b4d6258432a9a3d78dd9bffbcb6cfb1bd94f58da35fd530d08da7d1d05832.
//
// Solidity: event CalledFrom(address sender)
func (_TestCounter *TestCounterFilterer) WatchCalledFrom(opts *bind.WatchOpts, sink chan<- *TestCounterCalledFrom) (event.Subscription, error) {

	logs, sub, err := _TestCounter.contract.WatchLogs(opts, "CalledFrom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestCounterCalledFrom)
				if err := _TestCounter.contract.UnpackLog(event, "CalledFrom", log); err != nil {
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

// ParseCalledFrom is a log parse operation binding the contract event 0xfb3b4d6258432a9a3d78dd9bffbcb6cfb1bd94f58da35fd530d08da7d1d05832.
//
// Solidity: event CalledFrom(address sender)
func (_TestCounter *TestCounterFilterer) ParseCalledFrom(log types.Log) (*TestCounterCalledFrom, error) {
	event := new(TestCounterCalledFrom)
	if err := _TestCounter.contract.UnpackLog(event, "CalledFrom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
