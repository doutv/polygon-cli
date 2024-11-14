// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oraclehelper

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

// OracleHelperMetaData contains all meta data concerning the OracleHelper contract.
var OracleHelperMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cachedPriceTimestamp\",\"type\":\"uint256\"}],\"name\":\"TokenPriceUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cachedPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cachedPriceTimestamp\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"}],\"name\":\"updateCachedPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OracleHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleHelperMetaData.ABI instead.
var OracleHelperABI = OracleHelperMetaData.ABI

// OracleHelper is an auto generated Go binding around an Ethereum contract.
type OracleHelper struct {
	OracleHelperCaller     // Read-only binding to the contract
	OracleHelperTransactor // Write-only binding to the contract
	OracleHelperFilterer   // Log filterer for contract events
}

// OracleHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleHelperSession struct {
	Contract     *OracleHelper     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleHelperCallerSession struct {
	Contract *OracleHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OracleHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleHelperTransactorSession struct {
	Contract     *OracleHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OracleHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleHelperRaw struct {
	Contract *OracleHelper // Generic contract binding to access the raw methods on
}

// OracleHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleHelperCallerRaw struct {
	Contract *OracleHelperCaller // Generic read-only contract binding to access the raw methods on
}

// OracleHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleHelperTransactorRaw struct {
	Contract *OracleHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleHelper creates a new instance of OracleHelper, bound to a specific deployed contract.
func NewOracleHelper(address common.Address, backend bind.ContractBackend) (*OracleHelper, error) {
	contract, err := bindOracleHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleHelper{OracleHelperCaller: OracleHelperCaller{contract: contract}, OracleHelperTransactor: OracleHelperTransactor{contract: contract}, OracleHelperFilterer: OracleHelperFilterer{contract: contract}}, nil
}

// NewOracleHelperCaller creates a new read-only instance of OracleHelper, bound to a specific deployed contract.
func NewOracleHelperCaller(address common.Address, caller bind.ContractCaller) (*OracleHelperCaller, error) {
	contract, err := bindOracleHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleHelperCaller{contract: contract}, nil
}

// NewOracleHelperTransactor creates a new write-only instance of OracleHelper, bound to a specific deployed contract.
func NewOracleHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleHelperTransactor, error) {
	contract, err := bindOracleHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleHelperTransactor{contract: contract}, nil
}

// NewOracleHelperFilterer creates a new log filterer instance of OracleHelper, bound to a specific deployed contract.
func NewOracleHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleHelperFilterer, error) {
	contract, err := bindOracleHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleHelperFilterer{contract: contract}, nil
}

// bindOracleHelper binds a generic wrapper to an already deployed contract.
func bindOracleHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleHelper *OracleHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleHelper.Contract.OracleHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleHelper *OracleHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleHelper.Contract.OracleHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleHelper *OracleHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleHelper.Contract.OracleHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleHelper *OracleHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleHelper *OracleHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleHelper *OracleHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleHelper.Contract.contract.Transact(opts, method, params...)
}

// CachedPrice is a free data retrieval call binding the contract method 0xf60fdcb3.
//
// Solidity: function cachedPrice() view returns(uint256)
func (_OracleHelper *OracleHelperCaller) CachedPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleHelper.contract.Call(opts, &out, "cachedPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CachedPrice is a free data retrieval call binding the contract method 0xf60fdcb3.
//
// Solidity: function cachedPrice() view returns(uint256)
func (_OracleHelper *OracleHelperSession) CachedPrice() (*big.Int, error) {
	return _OracleHelper.Contract.CachedPrice(&_OracleHelper.CallOpts)
}

// CachedPrice is a free data retrieval call binding the contract method 0xf60fdcb3.
//
// Solidity: function cachedPrice() view returns(uint256)
func (_OracleHelper *OracleHelperCallerSession) CachedPrice() (*big.Int, error) {
	return _OracleHelper.Contract.CachedPrice(&_OracleHelper.CallOpts)
}

// CachedPriceTimestamp is a free data retrieval call binding the contract method 0xe1d8153c.
//
// Solidity: function cachedPriceTimestamp() view returns(uint48)
func (_OracleHelper *OracleHelperCaller) CachedPriceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleHelper.contract.Call(opts, &out, "cachedPriceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CachedPriceTimestamp is a free data retrieval call binding the contract method 0xe1d8153c.
//
// Solidity: function cachedPriceTimestamp() view returns(uint48)
func (_OracleHelper *OracleHelperSession) CachedPriceTimestamp() (*big.Int, error) {
	return _OracleHelper.Contract.CachedPriceTimestamp(&_OracleHelper.CallOpts)
}

// CachedPriceTimestamp is a free data retrieval call binding the contract method 0xe1d8153c.
//
// Solidity: function cachedPriceTimestamp() view returns(uint48)
func (_OracleHelper *OracleHelperCallerSession) CachedPriceTimestamp() (*big.Int, error) {
	return _OracleHelper.Contract.CachedPriceTimestamp(&_OracleHelper.CallOpts)
}

// UpdateCachedPrice is a paid mutator transaction binding the contract method 0x3ba9290f.
//
// Solidity: function updateCachedPrice(bool force) returns(uint256)
func (_OracleHelper *OracleHelperTransactor) UpdateCachedPrice(opts *bind.TransactOpts, force bool) (*types.Transaction, error) {
	return _OracleHelper.contract.Transact(opts, "updateCachedPrice", force)
}

// UpdateCachedPrice is a paid mutator transaction binding the contract method 0x3ba9290f.
//
// Solidity: function updateCachedPrice(bool force) returns(uint256)
func (_OracleHelper *OracleHelperSession) UpdateCachedPrice(force bool) (*types.Transaction, error) {
	return _OracleHelper.Contract.UpdateCachedPrice(&_OracleHelper.TransactOpts, force)
}

// UpdateCachedPrice is a paid mutator transaction binding the contract method 0x3ba9290f.
//
// Solidity: function updateCachedPrice(bool force) returns(uint256)
func (_OracleHelper *OracleHelperTransactorSession) UpdateCachedPrice(force bool) (*types.Transaction, error) {
	return _OracleHelper.Contract.UpdateCachedPrice(&_OracleHelper.TransactOpts, force)
}

// OracleHelperTokenPriceUpdatedIterator is returned from FilterTokenPriceUpdated and is used to iterate over the raw logs and unpacked data for TokenPriceUpdated events raised by the OracleHelper contract.
type OracleHelperTokenPriceUpdatedIterator struct {
	Event *OracleHelperTokenPriceUpdated // Event containing the contract specifics and raw log

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
func (it *OracleHelperTokenPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleHelperTokenPriceUpdated)
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
		it.Event = new(OracleHelperTokenPriceUpdated)
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
func (it *OracleHelperTokenPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleHelperTokenPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleHelperTokenPriceUpdated represents a TokenPriceUpdated event raised by the OracleHelper contract.
type OracleHelperTokenPriceUpdated struct {
	CurrentPrice         *big.Int
	PreviousPrice        *big.Int
	CachedPriceTimestamp *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTokenPriceUpdated is a free log retrieval operation binding the contract event 0x00d4fe314618b73a96886b87817a53a5ed51433b0234c85a5e9dafe2cb7b8842.
//
// Solidity: event TokenPriceUpdated(uint256 currentPrice, uint256 previousPrice, uint256 cachedPriceTimestamp)
func (_OracleHelper *OracleHelperFilterer) FilterTokenPriceUpdated(opts *bind.FilterOpts) (*OracleHelperTokenPriceUpdatedIterator, error) {

	logs, sub, err := _OracleHelper.contract.FilterLogs(opts, "TokenPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &OracleHelperTokenPriceUpdatedIterator{contract: _OracleHelper.contract, event: "TokenPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenPriceUpdated is a free log subscription operation binding the contract event 0x00d4fe314618b73a96886b87817a53a5ed51433b0234c85a5e9dafe2cb7b8842.
//
// Solidity: event TokenPriceUpdated(uint256 currentPrice, uint256 previousPrice, uint256 cachedPriceTimestamp)
func (_OracleHelper *OracleHelperFilterer) WatchTokenPriceUpdated(opts *bind.WatchOpts, sink chan<- *OracleHelperTokenPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _OracleHelper.contract.WatchLogs(opts, "TokenPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleHelperTokenPriceUpdated)
				if err := _OracleHelper.contract.UnpackLog(event, "TokenPriceUpdated", log); err != nil {
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

// ParseTokenPriceUpdated is a log parse operation binding the contract event 0x00d4fe314618b73a96886b87817a53a5ed51433b0234c85a5e9dafe2cb7b8842.
//
// Solidity: event TokenPriceUpdated(uint256 currentPrice, uint256 previousPrice, uint256 cachedPriceTimestamp)
func (_OracleHelper *OracleHelperFilterer) ParseTokenPriceUpdated(log types.Log) (*OracleHelperTokenPriceUpdated, error) {
	event := new(OracleHelperTokenPriceUpdated)
	if err := _OracleHelper.contract.UnpackLog(event, "TokenPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
