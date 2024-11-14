// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iblsaccount

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

// IBLSAccountMetaData contains all meta data concerning the IBLSAccount contract.
var IBLSAccountMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"oldPublicKey\",\"type\":\"uint256[4]\"},{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"newPublicKey\",\"type\":\"uint256[4]\"}],\"name\":\"PublicKeyChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getBlsPublicKey\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBLSAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use IBLSAccountMetaData.ABI instead.
var IBLSAccountABI = IBLSAccountMetaData.ABI

// IBLSAccount is an auto generated Go binding around an Ethereum contract.
type IBLSAccount struct {
	IBLSAccountCaller     // Read-only binding to the contract
	IBLSAccountTransactor // Write-only binding to the contract
	IBLSAccountFilterer   // Log filterer for contract events
}

// IBLSAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBLSAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBLSAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBLSAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBLSAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBLSAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBLSAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBLSAccountSession struct {
	Contract     *IBLSAccount      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBLSAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBLSAccountCallerSession struct {
	Contract *IBLSAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IBLSAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBLSAccountTransactorSession struct {
	Contract     *IBLSAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IBLSAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBLSAccountRaw struct {
	Contract *IBLSAccount // Generic contract binding to access the raw methods on
}

// IBLSAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBLSAccountCallerRaw struct {
	Contract *IBLSAccountCaller // Generic read-only contract binding to access the raw methods on
}

// IBLSAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBLSAccountTransactorRaw struct {
	Contract *IBLSAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBLSAccount creates a new instance of IBLSAccount, bound to a specific deployed contract.
func NewIBLSAccount(address common.Address, backend bind.ContractBackend) (*IBLSAccount, error) {
	contract, err := bindIBLSAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBLSAccount{IBLSAccountCaller: IBLSAccountCaller{contract: contract}, IBLSAccountTransactor: IBLSAccountTransactor{contract: contract}, IBLSAccountFilterer: IBLSAccountFilterer{contract: contract}}, nil
}

// NewIBLSAccountCaller creates a new read-only instance of IBLSAccount, bound to a specific deployed contract.
func NewIBLSAccountCaller(address common.Address, caller bind.ContractCaller) (*IBLSAccountCaller, error) {
	contract, err := bindIBLSAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBLSAccountCaller{contract: contract}, nil
}

// NewIBLSAccountTransactor creates a new write-only instance of IBLSAccount, bound to a specific deployed contract.
func NewIBLSAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*IBLSAccountTransactor, error) {
	contract, err := bindIBLSAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBLSAccountTransactor{contract: contract}, nil
}

// NewIBLSAccountFilterer creates a new log filterer instance of IBLSAccount, bound to a specific deployed contract.
func NewIBLSAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*IBLSAccountFilterer, error) {
	contract, err := bindIBLSAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBLSAccountFilterer{contract: contract}, nil
}

// bindIBLSAccount binds a generic wrapper to an already deployed contract.
func bindIBLSAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBLSAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBLSAccount *IBLSAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBLSAccount.Contract.IBLSAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBLSAccount *IBLSAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBLSAccount.Contract.IBLSAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBLSAccount *IBLSAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBLSAccount.Contract.IBLSAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBLSAccount *IBLSAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBLSAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBLSAccount *IBLSAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBLSAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBLSAccount *IBLSAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBLSAccount.Contract.contract.Transact(opts, method, params...)
}

// GetBlsPublicKey is a free data retrieval call binding the contract method 0xe02afbae.
//
// Solidity: function getBlsPublicKey() view returns(uint256[4])
func (_IBLSAccount *IBLSAccountCaller) GetBlsPublicKey(opts *bind.CallOpts) ([4]*big.Int, error) {
	var out []interface{}
	err := _IBLSAccount.contract.Call(opts, &out, "getBlsPublicKey")

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetBlsPublicKey is a free data retrieval call binding the contract method 0xe02afbae.
//
// Solidity: function getBlsPublicKey() view returns(uint256[4])
func (_IBLSAccount *IBLSAccountSession) GetBlsPublicKey() ([4]*big.Int, error) {
	return _IBLSAccount.Contract.GetBlsPublicKey(&_IBLSAccount.CallOpts)
}

// GetBlsPublicKey is a free data retrieval call binding the contract method 0xe02afbae.
//
// Solidity: function getBlsPublicKey() view returns(uint256[4])
func (_IBLSAccount *IBLSAccountCallerSession) GetBlsPublicKey() ([4]*big.Int, error) {
	return _IBLSAccount.Contract.GetBlsPublicKey(&_IBLSAccount.CallOpts)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_IBLSAccount *IBLSAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _IBLSAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_IBLSAccount *IBLSAccountSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _IBLSAccount.Contract.ValidateUserOp(&_IBLSAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_IBLSAccount *IBLSAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _IBLSAccount.Contract.ValidateUserOp(&_IBLSAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// IBLSAccountPublicKeyChangedIterator is returned from FilterPublicKeyChanged and is used to iterate over the raw logs and unpacked data for PublicKeyChanged events raised by the IBLSAccount contract.
type IBLSAccountPublicKeyChangedIterator struct {
	Event *IBLSAccountPublicKeyChanged // Event containing the contract specifics and raw log

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
func (it *IBLSAccountPublicKeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBLSAccountPublicKeyChanged)
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
		it.Event = new(IBLSAccountPublicKeyChanged)
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
func (it *IBLSAccountPublicKeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBLSAccountPublicKeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBLSAccountPublicKeyChanged represents a PublicKeyChanged event raised by the IBLSAccount contract.
type IBLSAccountPublicKeyChanged struct {
	OldPublicKey [4]*big.Int
	NewPublicKey [4]*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyChanged is a free log retrieval operation binding the contract event 0x42e4c4ce1432650f17e41c4ea77ed12c0ab20b229d3ffd84a2ebc9f8abb25a83.
//
// Solidity: event PublicKeyChanged(uint256[4] oldPublicKey, uint256[4] newPublicKey)
func (_IBLSAccount *IBLSAccountFilterer) FilterPublicKeyChanged(opts *bind.FilterOpts) (*IBLSAccountPublicKeyChangedIterator, error) {

	logs, sub, err := _IBLSAccount.contract.FilterLogs(opts, "PublicKeyChanged")
	if err != nil {
		return nil, err
	}
	return &IBLSAccountPublicKeyChangedIterator{contract: _IBLSAccount.contract, event: "PublicKeyChanged", logs: logs, sub: sub}, nil
}

// WatchPublicKeyChanged is a free log subscription operation binding the contract event 0x42e4c4ce1432650f17e41c4ea77ed12c0ab20b229d3ffd84a2ebc9f8abb25a83.
//
// Solidity: event PublicKeyChanged(uint256[4] oldPublicKey, uint256[4] newPublicKey)
func (_IBLSAccount *IBLSAccountFilterer) WatchPublicKeyChanged(opts *bind.WatchOpts, sink chan<- *IBLSAccountPublicKeyChanged) (event.Subscription, error) {

	logs, sub, err := _IBLSAccount.contract.WatchLogs(opts, "PublicKeyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBLSAccountPublicKeyChanged)
				if err := _IBLSAccount.contract.UnpackLog(event, "PublicKeyChanged", log); err != nil {
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

// ParsePublicKeyChanged is a log parse operation binding the contract event 0x42e4c4ce1432650f17e41c4ea77ed12c0ab20b229d3ffd84a2ebc9f8abb25a83.
//
// Solidity: event PublicKeyChanged(uint256[4] oldPublicKey, uint256[4] newPublicKey)
func (_IBLSAccount *IBLSAccountFilterer) ParsePublicKeyChanged(log types.Log) (*IBLSAccountPublicKeyChanged, error) {
	event := new(IBLSAccountPublicKeyChanged)
	if err := _IBLSAccount.contract.UnpackLog(event, "PublicKeyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
