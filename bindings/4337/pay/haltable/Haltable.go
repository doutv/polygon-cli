// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package haltable

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

// HaltableMetaData contains all meta data concerning the Haltable contract.
var HaltableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"HaltStatus\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"halted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234602057600160ff196000541617600055607e908160268239f35b600080fdfe6080806040526004361015601257600080fd5b600090813560e01c63b9b8af0b14602857600080fd5b3460445781600319360112604457600260ff6020935416148152f35b5080fdfea26469706673582212208127bb39e8c80f40ca33b87310fab6666533d1bcabf695dce34012b95f8f85de64736f6c63430008190033",
}

// HaltableABI is the input ABI used to generate the binding from.
// Deprecated: Use HaltableMetaData.ABI instead.
var HaltableABI = HaltableMetaData.ABI

// HaltableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HaltableMetaData.Bin instead.
var HaltableBin = HaltableMetaData.Bin

// DeployHaltable deploys a new Ethereum contract, binding an instance of Haltable to it.
func DeployHaltable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Haltable, error) {
	parsed, err := HaltableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HaltableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Haltable{HaltableCaller: HaltableCaller{contract: contract}, HaltableTransactor: HaltableTransactor{contract: contract}, HaltableFilterer: HaltableFilterer{contract: contract}}, nil
}

// Haltable is an auto generated Go binding around an Ethereum contract.
type Haltable struct {
	HaltableCaller     // Read-only binding to the contract
	HaltableTransactor // Write-only binding to the contract
	HaltableFilterer   // Log filterer for contract events
}

// HaltableCaller is an auto generated read-only Go binding around an Ethereum contract.
type HaltableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HaltableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HaltableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HaltableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HaltableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HaltableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HaltableSession struct {
	Contract     *Haltable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HaltableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HaltableCallerSession struct {
	Contract *HaltableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HaltableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HaltableTransactorSession struct {
	Contract     *HaltableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HaltableRaw is an auto generated low-level Go binding around an Ethereum contract.
type HaltableRaw struct {
	Contract *Haltable // Generic contract binding to access the raw methods on
}

// HaltableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HaltableCallerRaw struct {
	Contract *HaltableCaller // Generic read-only contract binding to access the raw methods on
}

// HaltableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HaltableTransactorRaw struct {
	Contract *HaltableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHaltable creates a new instance of Haltable, bound to a specific deployed contract.
func NewHaltable(address common.Address, backend bind.ContractBackend) (*Haltable, error) {
	contract, err := bindHaltable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Haltable{HaltableCaller: HaltableCaller{contract: contract}, HaltableTransactor: HaltableTransactor{contract: contract}, HaltableFilterer: HaltableFilterer{contract: contract}}, nil
}

// NewHaltableCaller creates a new read-only instance of Haltable, bound to a specific deployed contract.
func NewHaltableCaller(address common.Address, caller bind.ContractCaller) (*HaltableCaller, error) {
	contract, err := bindHaltable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HaltableCaller{contract: contract}, nil
}

// NewHaltableTransactor creates a new write-only instance of Haltable, bound to a specific deployed contract.
func NewHaltableTransactor(address common.Address, transactor bind.ContractTransactor) (*HaltableTransactor, error) {
	contract, err := bindHaltable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HaltableTransactor{contract: contract}, nil
}

// NewHaltableFilterer creates a new log filterer instance of Haltable, bound to a specific deployed contract.
func NewHaltableFilterer(address common.Address, filterer bind.ContractFilterer) (*HaltableFilterer, error) {
	contract, err := bindHaltable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HaltableFilterer{contract: contract}, nil
}

// bindHaltable binds a generic wrapper to an already deployed contract.
func bindHaltable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HaltableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Haltable *HaltableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Haltable.Contract.HaltableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Haltable *HaltableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Haltable.Contract.HaltableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Haltable *HaltableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Haltable.Contract.HaltableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Haltable *HaltableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Haltable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Haltable *HaltableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Haltable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Haltable *HaltableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Haltable.Contract.contract.Transact(opts, method, params...)
}

// Halted is a free data retrieval call binding the contract method 0xb9b8af0b.
//
// Solidity: function halted() view returns(bool)
func (_Haltable *HaltableCaller) Halted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Haltable.contract.Call(opts, &out, "halted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Halted is a free data retrieval call binding the contract method 0xb9b8af0b.
//
// Solidity: function halted() view returns(bool)
func (_Haltable *HaltableSession) Halted() (bool, error) {
	return _Haltable.Contract.Halted(&_Haltable.CallOpts)
}

// Halted is a free data retrieval call binding the contract method 0xb9b8af0b.
//
// Solidity: function halted() view returns(bool)
func (_Haltable *HaltableCallerSession) Halted() (bool, error) {
	return _Haltable.Contract.Halted(&_Haltable.CallOpts)
}

// HaltableHaltStatusIterator is returned from FilterHaltStatus and is used to iterate over the raw logs and unpacked data for HaltStatus events raised by the Haltable contract.
type HaltableHaltStatusIterator struct {
	Event *HaltableHaltStatus // Event containing the contract specifics and raw log

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
func (it *HaltableHaltStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HaltableHaltStatus)
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
		it.Event = new(HaltableHaltStatus)
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
func (it *HaltableHaltStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HaltableHaltStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HaltableHaltStatus represents a HaltStatus event raised by the Haltable contract.
type HaltableHaltStatus struct {
	Status *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHaltStatus is a free log retrieval operation binding the contract event 0xd57639a4f3c4c1059ca3fdd5175e8ca8ebf6f1eee17c2962389afa8bfd9b34e6.
//
// Solidity: event HaltStatus(uint256 status)
func (_Haltable *HaltableFilterer) FilterHaltStatus(opts *bind.FilterOpts) (*HaltableHaltStatusIterator, error) {

	logs, sub, err := _Haltable.contract.FilterLogs(opts, "HaltStatus")
	if err != nil {
		return nil, err
	}
	return &HaltableHaltStatusIterator{contract: _Haltable.contract, event: "HaltStatus", logs: logs, sub: sub}, nil
}

// WatchHaltStatus is a free log subscription operation binding the contract event 0xd57639a4f3c4c1059ca3fdd5175e8ca8ebf6f1eee17c2962389afa8bfd9b34e6.
//
// Solidity: event HaltStatus(uint256 status)
func (_Haltable *HaltableFilterer) WatchHaltStatus(opts *bind.WatchOpts, sink chan<- *HaltableHaltStatus) (event.Subscription, error) {

	logs, sub, err := _Haltable.contract.WatchLogs(opts, "HaltStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HaltableHaltStatus)
				if err := _Haltable.contract.UnpackLog(event, "HaltStatus", log); err != nil {
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

// ParseHaltStatus is a log parse operation binding the contract event 0xd57639a4f3c4c1059ca3fdd5175e8ca8ebf6f1eee17c2962389afa8bfd9b34e6.
//
// Solidity: event HaltStatus(uint256 status)
func (_Haltable *HaltableFilterer) ParseHaltStatus(log types.Log) (*HaltableHaltStatus, error) {
	event := new(HaltableHaltStatus)
	if err := _Haltable.contract.UnpackLog(event, "HaltStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
