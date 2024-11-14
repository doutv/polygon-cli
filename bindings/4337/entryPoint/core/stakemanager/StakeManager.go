// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakemanager

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

// StakeManagerMetaData contains all meta data concerning the StakeManager contract.
var StakeManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposit\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawTime\",\"type\":\"uint256\"}],\"name\":\"StakeUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"internalType\":\"structIStakeManager.DepositInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StakeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeManagerMetaData.ABI instead.
var StakeManagerABI = StakeManagerMetaData.ABI

// StakeManager is an auto generated Go binding around an Ethereum contract.
type StakeManager struct {
	StakeManagerCaller     // Read-only binding to the contract
	StakeManagerTransactor // Write-only binding to the contract
	StakeManagerFilterer   // Log filterer for contract events
}

// StakeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeManagerSession struct {
	Contract     *StakeManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeManagerCallerSession struct {
	Contract *StakeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StakeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeManagerTransactorSession struct {
	Contract     *StakeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeManagerRaw struct {
	Contract *StakeManager // Generic contract binding to access the raw methods on
}

// StakeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeManagerCallerRaw struct {
	Contract *StakeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeManagerTransactorRaw struct {
	Contract *StakeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeManager creates a new instance of StakeManager, bound to a specific deployed contract.
func NewStakeManager(address common.Address, backend bind.ContractBackend) (*StakeManager, error) {
	contract, err := bindStakeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeManager{StakeManagerCaller: StakeManagerCaller{contract: contract}, StakeManagerTransactor: StakeManagerTransactor{contract: contract}, StakeManagerFilterer: StakeManagerFilterer{contract: contract}}, nil
}

// NewStakeManagerCaller creates a new read-only instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerCaller(address common.Address, caller bind.ContractCaller) (*StakeManagerCaller, error) {
	contract, err := bindStakeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeManagerCaller{contract: contract}, nil
}

// NewStakeManagerTransactor creates a new write-only instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeManagerTransactor, error) {
	contract, err := bindStakeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeManagerTransactor{contract: contract}, nil
}

// NewStakeManagerFilterer creates a new log filterer instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeManagerFilterer, error) {
	contract, err := bindStakeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeManagerFilterer{contract: contract}, nil
}

// bindStakeManager binds a generic wrapper to an already deployed contract.
func bindStakeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeManager *StakeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeManager.Contract.StakeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeManager *StakeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeManager *StakeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeManager *StakeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeManager *StakeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeManager *StakeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeManager.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakeManager *StakeManagerCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakeManager *StakeManagerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StakeManager.Contract.BalanceOf(&_StakeManager.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StakeManager.Contract.BalanceOf(&_StakeManager.CallOpts, account)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_StakeManager *StakeManagerCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		Deposit         *big.Int
		Staked          bool
		Stake           *big.Int
		UnstakeDelaySec uint32
		WithdrawTime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Staked = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Stake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakeDelaySec = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.WithdrawTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_StakeManager *StakeManagerSession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _StakeManager.Contract.Deposits(&_StakeManager.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_StakeManager *StakeManagerCallerSession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _StakeManager.Contract.Deposits(&_StakeManager.CallOpts, arg0)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_StakeManager *StakeManagerCaller) GetDepositInfo(opts *bind.CallOpts, account common.Address) (IStakeManagerDepositInfo, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getDepositInfo", account)

	if err != nil {
		return *new(IStakeManagerDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeManagerDepositInfo)).(*IStakeManagerDepositInfo)

	return out0, err

}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_StakeManager *StakeManagerSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _StakeManager.Contract.GetDepositInfo(&_StakeManager.CallOpts, account)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_StakeManager *StakeManagerCallerSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _StakeManager.Contract.GetDepositInfo(&_StakeManager.CallOpts, account)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_StakeManager *StakeManagerTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_StakeManager *StakeManagerSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _StakeManager.Contract.AddStake(&_StakeManager.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_StakeManager *StakeManagerTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _StakeManager.Contract.AddStake(&_StakeManager.TransactOpts, unstakeDelaySec)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_StakeManager *StakeManagerTransactor) DepositTo(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "depositTo", account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_StakeManager *StakeManagerSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.DepositTo(&_StakeManager.TransactOpts, account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_StakeManager *StakeManagerTransactorSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.DepositTo(&_StakeManager.TransactOpts, account)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_StakeManager *StakeManagerTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_StakeManager *StakeManagerSession) UnlockStake() (*types.Transaction, error) {
	return _StakeManager.Contract.UnlockStake(&_StakeManager.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_StakeManager *StakeManagerTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _StakeManager.Contract.UnlockStake(&_StakeManager.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_StakeManager *StakeManagerTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_StakeManager *StakeManagerSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawStake(&_StakeManager.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_StakeManager *StakeManagerTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawStake(&_StakeManager.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_StakeManager *StakeManagerTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdrawTo", withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_StakeManager *StakeManagerSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawTo(&_StakeManager.TransactOpts, withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_StakeManager *StakeManagerTransactorSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawTo(&_StakeManager.TransactOpts, withdrawAddress, withdrawAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakeManager *StakeManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakeManager *StakeManagerSession) Receive() (*types.Transaction, error) {
	return _StakeManager.Contract.Receive(&_StakeManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakeManager *StakeManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _StakeManager.Contract.Receive(&_StakeManager.TransactOpts)
}

// StakeManagerDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the StakeManager contract.
type StakeManagerDepositedIterator struct {
	Event *StakeManagerDeposited // Event containing the contract specifics and raw log

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
func (it *StakeManagerDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerDeposited)
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
		it.Event = new(StakeManagerDeposited)
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
func (it *StakeManagerDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerDeposited represents a Deposited event raised by the StakeManager contract.
type StakeManagerDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_StakeManager *StakeManagerFilterer) FilterDeposited(opts *bind.FilterOpts, account []common.Address) (*StakeManagerDepositedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerDepositedIterator{contract: _StakeManager.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_StakeManager *StakeManagerFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *StakeManagerDeposited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerDeposited)
				if err := _StakeManager.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_StakeManager *StakeManagerFilterer) ParseDeposited(log types.Log) (*StakeManagerDeposited, error) {
	event := new(StakeManagerDeposited)
	if err := _StakeManager.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerStakeLockedIterator is returned from FilterStakeLocked and is used to iterate over the raw logs and unpacked data for StakeLocked events raised by the StakeManager contract.
type StakeManagerStakeLockedIterator struct {
	Event *StakeManagerStakeLocked // Event containing the contract specifics and raw log

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
func (it *StakeManagerStakeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerStakeLocked)
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
		it.Event = new(StakeManagerStakeLocked)
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
func (it *StakeManagerStakeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerStakeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerStakeLocked represents a StakeLocked event raised by the StakeManager contract.
type StakeManagerStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeLocked is a free log retrieval operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_StakeManager *StakeManagerFilterer) FilterStakeLocked(opts *bind.FilterOpts, account []common.Address) (*StakeManagerStakeLockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerStakeLockedIterator{contract: _StakeManager.contract, event: "StakeLocked", logs: logs, sub: sub}, nil
}

// WatchStakeLocked is a free log subscription operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_StakeManager *StakeManagerFilterer) WatchStakeLocked(opts *bind.WatchOpts, sink chan<- *StakeManagerStakeLocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerStakeLocked)
				if err := _StakeManager.contract.UnpackLog(event, "StakeLocked", log); err != nil {
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
func (_StakeManager *StakeManagerFilterer) ParseStakeLocked(log types.Log) (*StakeManagerStakeLocked, error) {
	event := new(StakeManagerStakeLocked)
	if err := _StakeManager.contract.UnpackLog(event, "StakeLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerStakeUnlockedIterator is returned from FilterStakeUnlocked and is used to iterate over the raw logs and unpacked data for StakeUnlocked events raised by the StakeManager contract.
type StakeManagerStakeUnlockedIterator struct {
	Event *StakeManagerStakeUnlocked // Event containing the contract specifics and raw log

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
func (it *StakeManagerStakeUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerStakeUnlocked)
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
		it.Event = new(StakeManagerStakeUnlocked)
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
func (it *StakeManagerStakeUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerStakeUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerStakeUnlocked represents a StakeUnlocked event raised by the StakeManager contract.
type StakeManagerStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeUnlocked is a free log retrieval operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_StakeManager *StakeManagerFilterer) FilterStakeUnlocked(opts *bind.FilterOpts, account []common.Address) (*StakeManagerStakeUnlockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerStakeUnlockedIterator{contract: _StakeManager.contract, event: "StakeUnlocked", logs: logs, sub: sub}, nil
}

// WatchStakeUnlocked is a free log subscription operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_StakeManager *StakeManagerFilterer) WatchStakeUnlocked(opts *bind.WatchOpts, sink chan<- *StakeManagerStakeUnlocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerStakeUnlocked)
				if err := _StakeManager.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
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
func (_StakeManager *StakeManagerFilterer) ParseStakeUnlocked(log types.Log) (*StakeManagerStakeUnlocked, error) {
	event := new(StakeManagerStakeUnlocked)
	if err := _StakeManager.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the StakeManager contract.
type StakeManagerStakeWithdrawnIterator struct {
	Event *StakeManagerStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakeManagerStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerStakeWithdrawn)
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
		it.Event = new(StakeManagerStakeWithdrawn)
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
func (it *StakeManagerStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerStakeWithdrawn represents a StakeWithdrawn event raised by the StakeManager contract.
type StakeManagerStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_StakeManager *StakeManagerFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, account []common.Address) (*StakeManagerStakeWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerStakeWithdrawnIterator{contract: _StakeManager.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_StakeManager *StakeManagerFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *StakeManagerStakeWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerStakeWithdrawn)
				if err := _StakeManager.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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
func (_StakeManager *StakeManagerFilterer) ParseStakeWithdrawn(log types.Log) (*StakeManagerStakeWithdrawn, error) {
	event := new(StakeManagerStakeWithdrawn)
	if err := _StakeManager.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the StakeManager contract.
type StakeManagerWithdrawnIterator struct {
	Event *StakeManagerWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakeManagerWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerWithdrawn)
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
		it.Event = new(StakeManagerWithdrawn)
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
func (it *StakeManagerWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerWithdrawn represents a Withdrawn event raised by the StakeManager contract.
type StakeManagerWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_StakeManager *StakeManagerFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*StakeManagerWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerWithdrawnIterator{contract: _StakeManager.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_StakeManager *StakeManagerFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *StakeManagerWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerWithdrawn)
				if err := _StakeManager.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_StakeManager *StakeManagerFilterer) ParseWithdrawn(log types.Log) (*StakeManagerWithdrawn, error) {
	event := new(StakeManagerWithdrawn)
	if err := _StakeManager.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
