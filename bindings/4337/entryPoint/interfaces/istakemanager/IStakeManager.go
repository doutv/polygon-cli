// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package istakemanager

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

// IStakeManagerDepositInfo is an auto generated low-level Go binding around an user-defined struct.
type IStakeManagerDepositInfo struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}

// IStakeManagerMetaData contains all meta data concerning the IStakeManager contract.
var IStakeManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposit\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawTime\",\"type\":\"uint256\"}],\"name\":\"StakeUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"internalType\":\"structIStakeManager.DepositInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IStakeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakeManagerMetaData.ABI instead.
var IStakeManagerABI = IStakeManagerMetaData.ABI

// IStakeManager is an auto generated Go binding around an Ethereum contract.
type IStakeManager struct {
	IStakeManagerCaller     // Read-only binding to the contract
	IStakeManagerTransactor // Write-only binding to the contract
	IStakeManagerFilterer   // Log filterer for contract events
}

// IStakeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakeManagerSession struct {
	Contract     *IStakeManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStakeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakeManagerCallerSession struct {
	Contract *IStakeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IStakeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakeManagerTransactorSession struct {
	Contract     *IStakeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IStakeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakeManagerRaw struct {
	Contract *IStakeManager // Generic contract binding to access the raw methods on
}

// IStakeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakeManagerCallerRaw struct {
	Contract *IStakeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IStakeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakeManagerTransactorRaw struct {
	Contract *IStakeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakeManager creates a new instance of IStakeManager, bound to a specific deployed contract.
func NewIStakeManager(address common.Address, backend bind.ContractBackend) (*IStakeManager, error) {
	contract, err := bindIStakeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakeManager{IStakeManagerCaller: IStakeManagerCaller{contract: contract}, IStakeManagerTransactor: IStakeManagerTransactor{contract: contract}, IStakeManagerFilterer: IStakeManagerFilterer{contract: contract}}, nil
}

// NewIStakeManagerCaller creates a new read-only instance of IStakeManager, bound to a specific deployed contract.
func NewIStakeManagerCaller(address common.Address, caller bind.ContractCaller) (*IStakeManagerCaller, error) {
	contract, err := bindIStakeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerCaller{contract: contract}, nil
}

// NewIStakeManagerTransactor creates a new write-only instance of IStakeManager, bound to a specific deployed contract.
func NewIStakeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakeManagerTransactor, error) {
	contract, err := bindIStakeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerTransactor{contract: contract}, nil
}

// NewIStakeManagerFilterer creates a new log filterer instance of IStakeManager, bound to a specific deployed contract.
func NewIStakeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakeManagerFilterer, error) {
	contract, err := bindIStakeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerFilterer{contract: contract}, nil
}

// bindIStakeManager binds a generic wrapper to an already deployed contract.
func bindIStakeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakeManager *IStakeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakeManager.Contract.IStakeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakeManager *IStakeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakeManager.Contract.IStakeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakeManager *IStakeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakeManager.Contract.IStakeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakeManager *IStakeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakeManager *IStakeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakeManager *IStakeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakeManager.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IStakeManager *IStakeManagerCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStakeManager.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IStakeManager *IStakeManagerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IStakeManager.Contract.BalanceOf(&_IStakeManager.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IStakeManager *IStakeManagerCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IStakeManager.Contract.BalanceOf(&_IStakeManager.CallOpts, account)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_IStakeManager *IStakeManagerCaller) GetDepositInfo(opts *bind.CallOpts, account common.Address) (IStakeManagerDepositInfo, error) {
	var out []interface{}
	err := _IStakeManager.contract.Call(opts, &out, "getDepositInfo", account)

	if err != nil {
		return *new(IStakeManagerDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeManagerDepositInfo)).(*IStakeManagerDepositInfo)

	return out0, err

}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_IStakeManager *IStakeManagerSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _IStakeManager.Contract.GetDepositInfo(&_IStakeManager.CallOpts, account)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_IStakeManager *IStakeManagerCallerSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _IStakeManager.Contract.GetDepositInfo(&_IStakeManager.CallOpts, account)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 _unstakeDelaySec) payable returns()
func (_IStakeManager *IStakeManagerTransactor) AddStake(opts *bind.TransactOpts, _unstakeDelaySec uint32) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "addStake", _unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 _unstakeDelaySec) payable returns()
func (_IStakeManager *IStakeManagerSession) AddStake(_unstakeDelaySec uint32) (*types.Transaction, error) {
	return _IStakeManager.Contract.AddStake(&_IStakeManager.TransactOpts, _unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 _unstakeDelaySec) payable returns()
func (_IStakeManager *IStakeManagerTransactorSession) AddStake(_unstakeDelaySec uint32) (*types.Transaction, error) {
	return _IStakeManager.Contract.AddStake(&_IStakeManager.TransactOpts, _unstakeDelaySec)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_IStakeManager *IStakeManagerTransactor) DepositTo(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "depositTo", account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_IStakeManager *IStakeManagerSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.DepositTo(&_IStakeManager.TransactOpts, account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_IStakeManager *IStakeManagerTransactorSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.DepositTo(&_IStakeManager.TransactOpts, account)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_IStakeManager *IStakeManagerTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_IStakeManager *IStakeManagerSession) UnlockStake() (*types.Transaction, error) {
	return _IStakeManager.Contract.UnlockStake(&_IStakeManager.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_IStakeManager *IStakeManagerTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _IStakeManager.Contract.UnlockStake(&_IStakeManager.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_IStakeManager *IStakeManagerTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_IStakeManager *IStakeManagerSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.WithdrawStake(&_IStakeManager.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_IStakeManager *IStakeManagerTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.WithdrawStake(&_IStakeManager.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_IStakeManager *IStakeManagerTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "withdrawTo", withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_IStakeManager *IStakeManagerSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.WithdrawTo(&_IStakeManager.TransactOpts, withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_IStakeManager *IStakeManagerTransactorSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.WithdrawTo(&_IStakeManager.TransactOpts, withdrawAddress, withdrawAmount)
}

// IStakeManagerDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the IStakeManager contract.
type IStakeManagerDepositedIterator struct {
	Event *IStakeManagerDeposited // Event containing the contract specifics and raw log

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
func (it *IStakeManagerDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeManagerDeposited)
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
		it.Event = new(IStakeManagerDeposited)
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
func (it *IStakeManagerDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeManagerDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeManagerDeposited represents a Deposited event raised by the IStakeManager contract.
type IStakeManagerDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_IStakeManager *IStakeManagerFilterer) FilterDeposited(opts *bind.FilterOpts, account []common.Address) (*IStakeManagerDepositedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IStakeManager.contract.FilterLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerDepositedIterator{contract: _IStakeManager.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_IStakeManager *IStakeManagerFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *IStakeManagerDeposited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IStakeManager.contract.WatchLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeManagerDeposited)
				if err := _IStakeManager.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_IStakeManager *IStakeManagerFilterer) ParseDeposited(log types.Log) (*IStakeManagerDeposited, error) {
	event := new(IStakeManagerDeposited)
	if err := _IStakeManager.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakeManagerStakeLockedIterator is returned from FilterStakeLocked and is used to iterate over the raw logs and unpacked data for StakeLocked events raised by the IStakeManager contract.
type IStakeManagerStakeLockedIterator struct {
	Event *IStakeManagerStakeLocked // Event containing the contract specifics and raw log

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
func (it *IStakeManagerStakeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeManagerStakeLocked)
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
		it.Event = new(IStakeManagerStakeLocked)
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
func (it *IStakeManagerStakeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeManagerStakeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeManagerStakeLocked represents a StakeLocked event raised by the IStakeManager contract.
type IStakeManagerStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeLocked is a free log retrieval operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_IStakeManager *IStakeManagerFilterer) FilterStakeLocked(opts *bind.FilterOpts, account []common.Address) (*IStakeManagerStakeLockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IStakeManager.contract.FilterLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerStakeLockedIterator{contract: _IStakeManager.contract, event: "StakeLocked", logs: logs, sub: sub}, nil
}

// WatchStakeLocked is a free log subscription operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_IStakeManager *IStakeManagerFilterer) WatchStakeLocked(opts *bind.WatchOpts, sink chan<- *IStakeManagerStakeLocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IStakeManager.contract.WatchLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeManagerStakeLocked)
				if err := _IStakeManager.contract.UnpackLog(event, "StakeLocked", log); err != nil {
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

// ParseStakeLocked is a log parse operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_IStakeManager *IStakeManagerFilterer) ParseStakeLocked(log types.Log) (*IStakeManagerStakeLocked, error) {
	event := new(IStakeManagerStakeLocked)
	if err := _IStakeManager.contract.UnpackLog(event, "StakeLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakeManagerStakeUnlockedIterator is returned from FilterStakeUnlocked and is used to iterate over the raw logs and unpacked data for StakeUnlocked events raised by the IStakeManager contract.
type IStakeManagerStakeUnlockedIterator struct {
	Event *IStakeManagerStakeUnlocked // Event containing the contract specifics and raw log

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
func (it *IStakeManagerStakeUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeManagerStakeUnlocked)
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
		it.Event = new(IStakeManagerStakeUnlocked)
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
func (it *IStakeManagerStakeUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeManagerStakeUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeManagerStakeUnlocked represents a StakeUnlocked event raised by the IStakeManager contract.
type IStakeManagerStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeUnlocked is a free log retrieval operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_IStakeManager *IStakeManagerFilterer) FilterStakeUnlocked(opts *bind.FilterOpts, account []common.Address) (*IStakeManagerStakeUnlockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IStakeManager.contract.FilterLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerStakeUnlockedIterator{contract: _IStakeManager.contract, event: "StakeUnlocked", logs: logs, sub: sub}, nil
}

// WatchStakeUnlocked is a free log subscription operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_IStakeManager *IStakeManagerFilterer) WatchStakeUnlocked(opts *bind.WatchOpts, sink chan<- *IStakeManagerStakeUnlocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IStakeManager.contract.WatchLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeManagerStakeUnlocked)
				if err := _IStakeManager.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
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

// ParseStakeUnlocked is a log parse operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_IStakeManager *IStakeManagerFilterer) ParseStakeUnlocked(log types.Log) (*IStakeManagerStakeUnlocked, error) {
	event := new(IStakeManagerStakeUnlocked)
	if err := _IStakeManager.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakeManagerStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the IStakeManager contract.
type IStakeManagerStakeWithdrawnIterator struct {
	Event *IStakeManagerStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *IStakeManagerStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeManagerStakeWithdrawn)
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
		it.Event = new(IStakeManagerStakeWithdrawn)
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
func (it *IStakeManagerStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeManagerStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeManagerStakeWithdrawn represents a StakeWithdrawn event raised by the IStakeManager contract.
type IStakeManagerStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_IStakeManager *IStakeManagerFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, account []common.Address) (*IStakeManagerStakeWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IStakeManager.contract.FilterLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerStakeWithdrawnIterator{contract: _IStakeManager.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_IStakeManager *IStakeManagerFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *IStakeManagerStakeWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IStakeManager.contract.WatchLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeManagerStakeWithdrawn)
				if err := _IStakeManager.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_IStakeManager *IStakeManagerFilterer) ParseStakeWithdrawn(log types.Log) (*IStakeManagerStakeWithdrawn, error) {
	event := new(IStakeManagerStakeWithdrawn)
	if err := _IStakeManager.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakeManagerWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the IStakeManager contract.
type IStakeManagerWithdrawnIterator struct {
	Event *IStakeManagerWithdrawn // Event containing the contract specifics and raw log

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
func (it *IStakeManagerWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeManagerWithdrawn)
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
		it.Event = new(IStakeManagerWithdrawn)
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
func (it *IStakeManagerWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeManagerWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeManagerWithdrawn represents a Withdrawn event raised by the IStakeManager contract.
type IStakeManagerWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_IStakeManager *IStakeManagerFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*IStakeManagerWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IStakeManager.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerWithdrawnIterator{contract: _IStakeManager.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_IStakeManager *IStakeManagerFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *IStakeManagerWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IStakeManager.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeManagerWithdrawn)
				if err := _IStakeManager.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_IStakeManager *IStakeManagerFilterer) ParseWithdrawn(log types.Log) (*IStakeManagerWithdrawn, error) {
	event := new(IStakeManagerWithdrawn)
	if err := _IStakeManager.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
