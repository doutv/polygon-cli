// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibasepaymaster

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

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	Sender             common.Address
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

// IBasePaymasterMetaData contains all meta data concerning the IBasePaymaster contract.
var IBasePaymasterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"config\",\"type\":\"address\"}],\"name\":\"SetConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasManager\",\"type\":\"address\"}],\"name\":\"SetGasManager\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"postOpMode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sigTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBasePaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use IBasePaymasterMetaData.ABI instead.
var IBasePaymasterABI = IBasePaymasterMetaData.ABI

// IBasePaymaster is an auto generated Go binding around an Ethereum contract.
type IBasePaymaster struct {
	IBasePaymasterCaller     // Read-only binding to the contract
	IBasePaymasterTransactor // Write-only binding to the contract
	IBasePaymasterFilterer   // Log filterer for contract events
}

// IBasePaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBasePaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBasePaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBasePaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBasePaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBasePaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBasePaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBasePaymasterSession struct {
	Contract     *IBasePaymaster   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBasePaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBasePaymasterCallerSession struct {
	Contract *IBasePaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IBasePaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBasePaymasterTransactorSession struct {
	Contract     *IBasePaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IBasePaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBasePaymasterRaw struct {
	Contract *IBasePaymaster // Generic contract binding to access the raw methods on
}

// IBasePaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBasePaymasterCallerRaw struct {
	Contract *IBasePaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// IBasePaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBasePaymasterTransactorRaw struct {
	Contract *IBasePaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBasePaymaster creates a new instance of IBasePaymaster, bound to a specific deployed contract.
func NewIBasePaymaster(address common.Address, backend bind.ContractBackend) (*IBasePaymaster, error) {
	contract, err := bindIBasePaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBasePaymaster{IBasePaymasterCaller: IBasePaymasterCaller{contract: contract}, IBasePaymasterTransactor: IBasePaymasterTransactor{contract: contract}, IBasePaymasterFilterer: IBasePaymasterFilterer{contract: contract}}, nil
}

// NewIBasePaymasterCaller creates a new read-only instance of IBasePaymaster, bound to a specific deployed contract.
func NewIBasePaymasterCaller(address common.Address, caller bind.ContractCaller) (*IBasePaymasterCaller, error) {
	contract, err := bindIBasePaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBasePaymasterCaller{contract: contract}, nil
}

// NewIBasePaymasterTransactor creates a new write-only instance of IBasePaymaster, bound to a specific deployed contract.
func NewIBasePaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*IBasePaymasterTransactor, error) {
	contract, err := bindIBasePaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBasePaymasterTransactor{contract: contract}, nil
}

// NewIBasePaymasterFilterer creates a new log filterer instance of IBasePaymaster, bound to a specific deployed contract.
func NewIBasePaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*IBasePaymasterFilterer, error) {
	contract, err := bindIBasePaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBasePaymasterFilterer{contract: contract}, nil
}

// bindIBasePaymaster binds a generic wrapper to an already deployed contract.
func bindIBasePaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBasePaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBasePaymaster *IBasePaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBasePaymaster.Contract.IBasePaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBasePaymaster *IBasePaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBasePaymaster.Contract.IBasePaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBasePaymaster *IBasePaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBasePaymaster.Contract.IBasePaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBasePaymaster *IBasePaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBasePaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBasePaymaster *IBasePaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBasePaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBasePaymaster *IBasePaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBasePaymaster.Contract.contract.Transact(opts, method, params...)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 postOpMode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_IBasePaymaster *IBasePaymasterTransactor) PostOp(opts *bind.TransactOpts, postOpMode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _IBasePaymaster.contract.Transact(opts, "postOp", postOpMode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 postOpMode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_IBasePaymaster *IBasePaymasterSession) PostOp(postOpMode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _IBasePaymaster.Contract.PostOp(&_IBasePaymaster.TransactOpts, postOpMode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 postOpMode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_IBasePaymaster *IBasePaymasterTransactorSession) PostOp(postOpMode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _IBasePaymaster.Contract.PostOp(&_IBasePaymaster.TransactOpts, postOpMode, context, actualGasCost, actualUserOpFeePerGas)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 ) returns(bytes context, uint256 sigTime)
func (_IBasePaymaster *IBasePaymasterTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, arg2 *big.Int) (*types.Transaction, error) {
	return _IBasePaymaster.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, arg2)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 ) returns(bytes context, uint256 sigTime)
func (_IBasePaymaster *IBasePaymasterSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, arg2 *big.Int) (*types.Transaction, error) {
	return _IBasePaymaster.Contract.ValidatePaymasterUserOp(&_IBasePaymaster.TransactOpts, userOp, userOpHash, arg2)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 ) returns(bytes context, uint256 sigTime)
func (_IBasePaymaster *IBasePaymasterTransactorSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, arg2 *big.Int) (*types.Transaction, error) {
	return _IBasePaymaster.Contract.ValidatePaymasterUserOp(&_IBasePaymaster.TransactOpts, userOp, userOpHash, arg2)
}

// WithdrawGas is a paid mutator transaction binding the contract method 0x813f3f44.
//
// Solidity: function withdrawGas(address recipient, uint256 amount) returns()
func (_IBasePaymaster *IBasePaymasterTransactor) WithdrawGas(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBasePaymaster.contract.Transact(opts, "withdrawGas", recipient, amount)
}

// WithdrawGas is a paid mutator transaction binding the contract method 0x813f3f44.
//
// Solidity: function withdrawGas(address recipient, uint256 amount) returns()
func (_IBasePaymaster *IBasePaymasterSession) WithdrawGas(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBasePaymaster.Contract.WithdrawGas(&_IBasePaymaster.TransactOpts, recipient, amount)
}

// WithdrawGas is a paid mutator transaction binding the contract method 0x813f3f44.
//
// Solidity: function withdrawGas(address recipient, uint256 amount) returns()
func (_IBasePaymaster *IBasePaymasterTransactorSession) WithdrawGas(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBasePaymaster.Contract.WithdrawGas(&_IBasePaymaster.TransactOpts, recipient, amount)
}

// IBasePaymasterGasDepositedIterator is returned from FilterGasDeposited and is used to iterate over the raw logs and unpacked data for GasDeposited events raised by the IBasePaymaster contract.
type IBasePaymasterGasDepositedIterator struct {
	Event *IBasePaymasterGasDeposited // Event containing the contract specifics and raw log

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
func (it *IBasePaymasterGasDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBasePaymasterGasDeposited)
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
		it.Event = new(IBasePaymasterGasDeposited)
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
func (it *IBasePaymasterGasDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBasePaymasterGasDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBasePaymasterGasDeposited represents a GasDeposited event raised by the IBasePaymaster contract.
type IBasePaymasterGasDeposited struct {
	EntryPoint common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasDeposited is a free log retrieval operation binding the contract event 0x1dbbf474736d6415d6a265fabee708fe6e988f6fd0c9d870ded36cab380898dd.
//
// Solidity: event GasDeposited(address entryPoint, uint256 amount)
func (_IBasePaymaster *IBasePaymasterFilterer) FilterGasDeposited(opts *bind.FilterOpts) (*IBasePaymasterGasDepositedIterator, error) {

	logs, sub, err := _IBasePaymaster.contract.FilterLogs(opts, "GasDeposited")
	if err != nil {
		return nil, err
	}
	return &IBasePaymasterGasDepositedIterator{contract: _IBasePaymaster.contract, event: "GasDeposited", logs: logs, sub: sub}, nil
}

// WatchGasDeposited is a free log subscription operation binding the contract event 0x1dbbf474736d6415d6a265fabee708fe6e988f6fd0c9d870ded36cab380898dd.
//
// Solidity: event GasDeposited(address entryPoint, uint256 amount)
func (_IBasePaymaster *IBasePaymasterFilterer) WatchGasDeposited(opts *bind.WatchOpts, sink chan<- *IBasePaymasterGasDeposited) (event.Subscription, error) {

	logs, sub, err := _IBasePaymaster.contract.WatchLogs(opts, "GasDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBasePaymasterGasDeposited)
				if err := _IBasePaymaster.contract.UnpackLog(event, "GasDeposited", log); err != nil {
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

// ParseGasDeposited is a log parse operation binding the contract event 0x1dbbf474736d6415d6a265fabee708fe6e988f6fd0c9d870ded36cab380898dd.
//
// Solidity: event GasDeposited(address entryPoint, uint256 amount)
func (_IBasePaymaster *IBasePaymasterFilterer) ParseGasDeposited(log types.Log) (*IBasePaymasterGasDeposited, error) {
	event := new(IBasePaymasterGasDeposited)
	if err := _IBasePaymaster.contract.UnpackLog(event, "GasDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBasePaymasterGasWithdrawnIterator is returned from FilterGasWithdrawn and is used to iterate over the raw logs and unpacked data for GasWithdrawn events raised by the IBasePaymaster contract.
type IBasePaymasterGasWithdrawnIterator struct {
	Event *IBasePaymasterGasWithdrawn // Event containing the contract specifics and raw log

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
func (it *IBasePaymasterGasWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBasePaymasterGasWithdrawn)
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
		it.Event = new(IBasePaymasterGasWithdrawn)
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
func (it *IBasePaymasterGasWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBasePaymasterGasWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBasePaymasterGasWithdrawn represents a GasWithdrawn event raised by the IBasePaymaster contract.
type IBasePaymasterGasWithdrawn struct {
	EntryPoint common.Address
	Recipient  common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasWithdrawn is a free log retrieval operation binding the contract event 0x926a144b6fffc1d73f115b81af7ec66a7c12aed0ff73197c39a683753fc1d925.
//
// Solidity: event GasWithdrawn(address entryPoint, address recipient, uint256 amount)
func (_IBasePaymaster *IBasePaymasterFilterer) FilterGasWithdrawn(opts *bind.FilterOpts) (*IBasePaymasterGasWithdrawnIterator, error) {

	logs, sub, err := _IBasePaymaster.contract.FilterLogs(opts, "GasWithdrawn")
	if err != nil {
		return nil, err
	}
	return &IBasePaymasterGasWithdrawnIterator{contract: _IBasePaymaster.contract, event: "GasWithdrawn", logs: logs, sub: sub}, nil
}

// WatchGasWithdrawn is a free log subscription operation binding the contract event 0x926a144b6fffc1d73f115b81af7ec66a7c12aed0ff73197c39a683753fc1d925.
//
// Solidity: event GasWithdrawn(address entryPoint, address recipient, uint256 amount)
func (_IBasePaymaster *IBasePaymasterFilterer) WatchGasWithdrawn(opts *bind.WatchOpts, sink chan<- *IBasePaymasterGasWithdrawn) (event.Subscription, error) {

	logs, sub, err := _IBasePaymaster.contract.WatchLogs(opts, "GasWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBasePaymasterGasWithdrawn)
				if err := _IBasePaymaster.contract.UnpackLog(event, "GasWithdrawn", log); err != nil {
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

// ParseGasWithdrawn is a log parse operation binding the contract event 0x926a144b6fffc1d73f115b81af7ec66a7c12aed0ff73197c39a683753fc1d925.
//
// Solidity: event GasWithdrawn(address entryPoint, address recipient, uint256 amount)
func (_IBasePaymaster *IBasePaymasterFilterer) ParseGasWithdrawn(log types.Log) (*IBasePaymasterGasWithdrawn, error) {
	event := new(IBasePaymasterGasWithdrawn)
	if err := _IBasePaymaster.contract.UnpackLog(event, "GasWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBasePaymasterSetConfigIterator is returned from FilterSetConfig and is used to iterate over the raw logs and unpacked data for SetConfig events raised by the IBasePaymaster contract.
type IBasePaymasterSetConfigIterator struct {
	Event *IBasePaymasterSetConfig // Event containing the contract specifics and raw log

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
func (it *IBasePaymasterSetConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBasePaymasterSetConfig)
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
		it.Event = new(IBasePaymasterSetConfig)
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
func (it *IBasePaymasterSetConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBasePaymasterSetConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBasePaymasterSetConfig represents a SetConfig event raised by the IBasePaymaster contract.
type IBasePaymasterSetConfig struct {
	Config common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetConfig is a free log retrieval operation binding the contract event 0xc5618716db99966ac0bedb011a55472827d54343d73b50c3118c0b03cdf1c75f.
//
// Solidity: event SetConfig(address config)
func (_IBasePaymaster *IBasePaymasterFilterer) FilterSetConfig(opts *bind.FilterOpts) (*IBasePaymasterSetConfigIterator, error) {

	logs, sub, err := _IBasePaymaster.contract.FilterLogs(opts, "SetConfig")
	if err != nil {
		return nil, err
	}
	return &IBasePaymasterSetConfigIterator{contract: _IBasePaymaster.contract, event: "SetConfig", logs: logs, sub: sub}, nil
}

// WatchSetConfig is a free log subscription operation binding the contract event 0xc5618716db99966ac0bedb011a55472827d54343d73b50c3118c0b03cdf1c75f.
//
// Solidity: event SetConfig(address config)
func (_IBasePaymaster *IBasePaymasterFilterer) WatchSetConfig(opts *bind.WatchOpts, sink chan<- *IBasePaymasterSetConfig) (event.Subscription, error) {

	logs, sub, err := _IBasePaymaster.contract.WatchLogs(opts, "SetConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBasePaymasterSetConfig)
				if err := _IBasePaymaster.contract.UnpackLog(event, "SetConfig", log); err != nil {
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

// ParseSetConfig is a log parse operation binding the contract event 0xc5618716db99966ac0bedb011a55472827d54343d73b50c3118c0b03cdf1c75f.
//
// Solidity: event SetConfig(address config)
func (_IBasePaymaster *IBasePaymasterFilterer) ParseSetConfig(log types.Log) (*IBasePaymasterSetConfig, error) {
	event := new(IBasePaymasterSetConfig)
	if err := _IBasePaymaster.contract.UnpackLog(event, "SetConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBasePaymasterSetGasManagerIterator is returned from FilterSetGasManager and is used to iterate over the raw logs and unpacked data for SetGasManager events raised by the IBasePaymaster contract.
type IBasePaymasterSetGasManagerIterator struct {
	Event *IBasePaymasterSetGasManager // Event containing the contract specifics and raw log

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
func (it *IBasePaymasterSetGasManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBasePaymasterSetGasManager)
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
		it.Event = new(IBasePaymasterSetGasManager)
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
func (it *IBasePaymasterSetGasManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBasePaymasterSetGasManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBasePaymasterSetGasManager represents a SetGasManager event raised by the IBasePaymaster contract.
type IBasePaymasterSetGasManager struct {
	GasManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetGasManager is a free log retrieval operation binding the contract event 0x9daebb79792c29c1b590ef48a1d6cb59a9150c1d281fc94b12724321615e203a.
//
// Solidity: event SetGasManager(address gasManager)
func (_IBasePaymaster *IBasePaymasterFilterer) FilterSetGasManager(opts *bind.FilterOpts) (*IBasePaymasterSetGasManagerIterator, error) {

	logs, sub, err := _IBasePaymaster.contract.FilterLogs(opts, "SetGasManager")
	if err != nil {
		return nil, err
	}
	return &IBasePaymasterSetGasManagerIterator{contract: _IBasePaymaster.contract, event: "SetGasManager", logs: logs, sub: sub}, nil
}

// WatchSetGasManager is a free log subscription operation binding the contract event 0x9daebb79792c29c1b590ef48a1d6cb59a9150c1d281fc94b12724321615e203a.
//
// Solidity: event SetGasManager(address gasManager)
func (_IBasePaymaster *IBasePaymasterFilterer) WatchSetGasManager(opts *bind.WatchOpts, sink chan<- *IBasePaymasterSetGasManager) (event.Subscription, error) {

	logs, sub, err := _IBasePaymaster.contract.WatchLogs(opts, "SetGasManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBasePaymasterSetGasManager)
				if err := _IBasePaymaster.contract.UnpackLog(event, "SetGasManager", log); err != nil {
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

// ParseSetGasManager is a log parse operation binding the contract event 0x9daebb79792c29c1b590ef48a1d6cb59a9150c1d281fc94b12724321615e203a.
//
// Solidity: event SetGasManager(address gasManager)
func (_IBasePaymaster *IBasePaymasterFilterer) ParseSetGasManager(log types.Log) (*IBasePaymasterSetGasManager, error) {
	event := new(IBasePaymasterSetGasManager)
	if err := _IBasePaymaster.contract.UnpackLog(event, "SetGasManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
