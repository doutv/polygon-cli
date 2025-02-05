// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipay

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

// IPayChequeParams is an auto generated low-level Go binding around an user-defined struct.
type IPayChequeParams struct {
	ChequeID     *big.Int
	To           common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Expiration   *big.Int
}

// IPayMetaData contains all meta data concerning the IPay contract.
var IPayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldExpirationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExpirationTime\",\"type\":\"uint256\"}],\"name\":\"ChangeMaxExpirationTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumChequeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ChequeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"ChequeSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"WhitelistTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"WhitelistTokenRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"addWhitelistedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"chequeIDs\",\"type\":\"uint256[]\"}],\"name\":\"batchRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isWhitelistedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"removeWhitelistedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"internalType\":\"structIPay.ChequeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IPayABI is the input ABI used to generate the binding from.
// Deprecated: Use IPayMetaData.ABI instead.
var IPayABI = IPayMetaData.ABI

// IPay is an auto generated Go binding around an Ethereum contract.
type IPay struct {
	IPayCaller     // Read-only binding to the contract
	IPayTransactor // Write-only binding to the contract
	IPayFilterer   // Log filterer for contract events
}

// IPayCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPaySession struct {
	Contract     *IPay             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPayCallerSession struct {
	Contract *IPayCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IPayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPayTransactorSession struct {
	Contract     *IPayTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPayRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPayRaw struct {
	Contract *IPay // Generic contract binding to access the raw methods on
}

// IPayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPayCallerRaw struct {
	Contract *IPayCaller // Generic read-only contract binding to access the raw methods on
}

// IPayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPayTransactorRaw struct {
	Contract *IPayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPay creates a new instance of IPay, bound to a specific deployed contract.
func NewIPay(address common.Address, backend bind.ContractBackend) (*IPay, error) {
	contract, err := bindIPay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPay{IPayCaller: IPayCaller{contract: contract}, IPayTransactor: IPayTransactor{contract: contract}, IPayFilterer: IPayFilterer{contract: contract}}, nil
}

// NewIPayCaller creates a new read-only instance of IPay, bound to a specific deployed contract.
func NewIPayCaller(address common.Address, caller bind.ContractCaller) (*IPayCaller, error) {
	contract, err := bindIPay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPayCaller{contract: contract}, nil
}

// NewIPayTransactor creates a new write-only instance of IPay, bound to a specific deployed contract.
func NewIPayTransactor(address common.Address, transactor bind.ContractTransactor) (*IPayTransactor, error) {
	contract, err := bindIPay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPayTransactor{contract: contract}, nil
}

// NewIPayFilterer creates a new log filterer instance of IPay, bound to a specific deployed contract.
func NewIPayFilterer(address common.Address, filterer bind.ContractFilterer) (*IPayFilterer, error) {
	contract, err := bindIPay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPayFilterer{contract: contract}, nil
}

// bindIPay binds a generic wrapper to an already deployed contract.
func bindIPay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPay *IPayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPay.Contract.IPayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPay *IPayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPay.Contract.IPayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPay *IPayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPay.Contract.IPayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPay *IPayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPay *IPayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPay *IPayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPay.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IPay *IPayCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IPay.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IPay *IPaySession) Name() (string, error) {
	return _IPay.Contract.Name(&_IPay.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IPay *IPayCallerSession) Name() (string, error) {
	return _IPay.Contract.Name(&_IPay.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IPay *IPayCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IPay.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IPay *IPaySession) Version() (string, error) {
	return _IPay.Contract.Version(&_IPay.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IPay *IPayCallerSession) Version() (string, error) {
	return _IPay.Contract.Version(&_IPay.CallOpts)
}

// AddWhitelistedTokens is a paid mutator transaction binding the contract method 0xcd698e69.
//
// Solidity: function addWhitelistedTokens(address[] tokens) returns()
func (_IPay *IPayTransactor) AddWhitelistedTokens(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _IPay.contract.Transact(opts, "addWhitelistedTokens", tokens)
}

// AddWhitelistedTokens is a paid mutator transaction binding the contract method 0xcd698e69.
//
// Solidity: function addWhitelistedTokens(address[] tokens) returns()
func (_IPay *IPaySession) AddWhitelistedTokens(tokens []common.Address) (*types.Transaction, error) {
	return _IPay.Contract.AddWhitelistedTokens(&_IPay.TransactOpts, tokens)
}

// AddWhitelistedTokens is a paid mutator transaction binding the contract method 0xcd698e69.
//
// Solidity: function addWhitelistedTokens(address[] tokens) returns()
func (_IPay *IPayTransactorSession) AddWhitelistedTokens(tokens []common.Address) (*types.Transaction, error) {
	return _IPay.Contract.AddWhitelistedTokens(&_IPay.TransactOpts, tokens)
}

// BatchRefund is a paid mutator transaction binding the contract method 0x05424c54.
//
// Solidity: function batchRefund(uint256[] chequeIDs) returns()
func (_IPay *IPayTransactor) BatchRefund(opts *bind.TransactOpts, chequeIDs []*big.Int) (*types.Transaction, error) {
	return _IPay.contract.Transact(opts, "batchRefund", chequeIDs)
}

// BatchRefund is a paid mutator transaction binding the contract method 0x05424c54.
//
// Solidity: function batchRefund(uint256[] chequeIDs) returns()
func (_IPay *IPaySession) BatchRefund(chequeIDs []*big.Int) (*types.Transaction, error) {
	return _IPay.Contract.BatchRefund(&_IPay.TransactOpts, chequeIDs)
}

// BatchRefund is a paid mutator transaction binding the contract method 0x05424c54.
//
// Solidity: function batchRefund(uint256[] chequeIDs) returns()
func (_IPay *IPayTransactorSession) BatchRefund(chequeIDs []*big.Int) (*types.Transaction, error) {
	return _IPay.Contract.BatchRefund(&_IPay.TransactOpts, chequeIDs)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 chequeID) returns()
func (_IPay *IPayTransactor) Cancel(opts *bind.TransactOpts, chequeID *big.Int) (*types.Transaction, error) {
	return _IPay.contract.Transact(opts, "cancel", chequeID)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 chequeID) returns()
func (_IPay *IPaySession) Cancel(chequeID *big.Int) (*types.Transaction, error) {
	return _IPay.Contract.Cancel(&_IPay.TransactOpts, chequeID)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 chequeID) returns()
func (_IPay *IPayTransactorSession) Cancel(chequeID *big.Int) (*types.Transaction, error) {
	return _IPay.Contract.Cancel(&_IPay.TransactOpts, chequeID)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 chequeID) returns()
func (_IPay *IPayTransactor) Claim(opts *bind.TransactOpts, chequeID *big.Int) (*types.Transaction, error) {
	return _IPay.contract.Transact(opts, "claim", chequeID)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 chequeID) returns()
func (_IPay *IPaySession) Claim(chequeID *big.Int) (*types.Transaction, error) {
	return _IPay.Contract.Claim(&_IPay.TransactOpts, chequeID)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 chequeID) returns()
func (_IPay *IPayTransactorSession) Claim(chequeID *big.Int) (*types.Transaction, error) {
	return _IPay.Contract.Claim(&_IPay.TransactOpts, chequeID)
}

// IsWhitelistedToken is a paid mutator transaction binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address token) returns(bool)
func (_IPay *IPayTransactor) IsWhitelistedToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _IPay.contract.Transact(opts, "isWhitelistedToken", token)
}

// IsWhitelistedToken is a paid mutator transaction binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address token) returns(bool)
func (_IPay *IPaySession) IsWhitelistedToken(token common.Address) (*types.Transaction, error) {
	return _IPay.Contract.IsWhitelistedToken(&_IPay.TransactOpts, token)
}

// IsWhitelistedToken is a paid mutator transaction binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address token) returns(bool)
func (_IPay *IPayTransactorSession) IsWhitelistedToken(token common.Address) (*types.Transaction, error) {
	return _IPay.Contract.IsWhitelistedToken(&_IPay.TransactOpts, token)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 chequeID) returns()
func (_IPay *IPayTransactor) Refund(opts *bind.TransactOpts, chequeID *big.Int) (*types.Transaction, error) {
	return _IPay.contract.Transact(opts, "refund", chequeID)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 chequeID) returns()
func (_IPay *IPaySession) Refund(chequeID *big.Int) (*types.Transaction, error) {
	return _IPay.Contract.Refund(&_IPay.TransactOpts, chequeID)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 chequeID) returns()
func (_IPay *IPayTransactorSession) Refund(chequeID *big.Int) (*types.Transaction, error) {
	return _IPay.Contract.Refund(&_IPay.TransactOpts, chequeID)
}

// RemoveWhitelistedTokens is a paid mutator transaction binding the contract method 0xbcec454f.
//
// Solidity: function removeWhitelistedTokens(address[] tokens) returns()
func (_IPay *IPayTransactor) RemoveWhitelistedTokens(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _IPay.contract.Transact(opts, "removeWhitelistedTokens", tokens)
}

// RemoveWhitelistedTokens is a paid mutator transaction binding the contract method 0xbcec454f.
//
// Solidity: function removeWhitelistedTokens(address[] tokens) returns()
func (_IPay *IPaySession) RemoveWhitelistedTokens(tokens []common.Address) (*types.Transaction, error) {
	return _IPay.Contract.RemoveWhitelistedTokens(&_IPay.TransactOpts, tokens)
}

// RemoveWhitelistedTokens is a paid mutator transaction binding the contract method 0xbcec454f.
//
// Solidity: function removeWhitelistedTokens(address[] tokens) returns()
func (_IPay *IPayTransactorSession) RemoveWhitelistedTokens(tokens []common.Address) (*types.Transaction, error) {
	return _IPay.Contract.RemoveWhitelistedTokens(&_IPay.TransactOpts, tokens)
}

// Send is a paid mutator transaction binding the contract method 0x96e1274d.
//
// Solidity: function send((uint256,address,address,uint256,uint256) params) payable returns()
func (_IPay *IPayTransactor) Send(opts *bind.TransactOpts, params IPayChequeParams) (*types.Transaction, error) {
	return _IPay.contract.Transact(opts, "send", params)
}

// Send is a paid mutator transaction binding the contract method 0x96e1274d.
//
// Solidity: function send((uint256,address,address,uint256,uint256) params) payable returns()
func (_IPay *IPaySession) Send(params IPayChequeParams) (*types.Transaction, error) {
	return _IPay.Contract.Send(&_IPay.TransactOpts, params)
}

// Send is a paid mutator transaction binding the contract method 0x96e1274d.
//
// Solidity: function send((uint256,address,address,uint256,uint256) params) payable returns()
func (_IPay *IPayTransactorSession) Send(params IPayChequeParams) (*types.Transaction, error) {
	return _IPay.Contract.Send(&_IPay.TransactOpts, params)
}

// IPayChangeMaxExpirationTimeIterator is returned from FilterChangeMaxExpirationTime and is used to iterate over the raw logs and unpacked data for ChangeMaxExpirationTime events raised by the IPay contract.
type IPayChangeMaxExpirationTimeIterator struct {
	Event *IPayChangeMaxExpirationTime // Event containing the contract specifics and raw log

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
func (it *IPayChangeMaxExpirationTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayChangeMaxExpirationTime)
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
		it.Event = new(IPayChangeMaxExpirationTime)
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
func (it *IPayChangeMaxExpirationTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayChangeMaxExpirationTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayChangeMaxExpirationTime represents a ChangeMaxExpirationTime event raised by the IPay contract.
type IPayChangeMaxExpirationTime struct {
	OldExpirationTime *big.Int
	NewExpirationTime *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChangeMaxExpirationTime is a free log retrieval operation binding the contract event 0xd764d769918360f43f93e481dd41732ac1e5a40e2407dd7c569651126d82160f.
//
// Solidity: event ChangeMaxExpirationTime(uint256 oldExpirationTime, uint256 newExpirationTime)
func (_IPay *IPayFilterer) FilterChangeMaxExpirationTime(opts *bind.FilterOpts) (*IPayChangeMaxExpirationTimeIterator, error) {

	logs, sub, err := _IPay.contract.FilterLogs(opts, "ChangeMaxExpirationTime")
	if err != nil {
		return nil, err
	}
	return &IPayChangeMaxExpirationTimeIterator{contract: _IPay.contract, event: "ChangeMaxExpirationTime", logs: logs, sub: sub}, nil
}

// WatchChangeMaxExpirationTime is a free log subscription operation binding the contract event 0xd764d769918360f43f93e481dd41732ac1e5a40e2407dd7c569651126d82160f.
//
// Solidity: event ChangeMaxExpirationTime(uint256 oldExpirationTime, uint256 newExpirationTime)
func (_IPay *IPayFilterer) WatchChangeMaxExpirationTime(opts *bind.WatchOpts, sink chan<- *IPayChangeMaxExpirationTime) (event.Subscription, error) {

	logs, sub, err := _IPay.contract.WatchLogs(opts, "ChangeMaxExpirationTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayChangeMaxExpirationTime)
				if err := _IPay.contract.UnpackLog(event, "ChangeMaxExpirationTime", log); err != nil {
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

// ParseChangeMaxExpirationTime is a log parse operation binding the contract event 0xd764d769918360f43f93e481dd41732ac1e5a40e2407dd7c569651126d82160f.
//
// Solidity: event ChangeMaxExpirationTime(uint256 oldExpirationTime, uint256 newExpirationTime)
func (_IPay *IPayFilterer) ParseChangeMaxExpirationTime(log types.Log) (*IPayChangeMaxExpirationTime, error) {
	event := new(IPayChangeMaxExpirationTime)
	if err := _IPay.contract.UnpackLog(event, "ChangeMaxExpirationTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPayChequeEventIterator is returned from FilterChequeEvent and is used to iterate over the raw logs and unpacked data for ChequeEvent events raised by the IPay contract.
type IPayChequeEventIterator struct {
	Event *IPayChequeEvent // Event containing the contract specifics and raw log

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
func (it *IPayChequeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayChequeEvent)
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
		it.Event = new(IPayChequeEvent)
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
func (it *IPayChequeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayChequeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayChequeEvent represents a ChequeEvent event raised by the IPay contract.
type IPayChequeEvent struct {
	ChequeID     *big.Int
	From         common.Address
	To           common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChequeEvent is a free log retrieval operation binding the contract event 0x4617b59973e052c439b84edb2deaedd5c5dad563af6952409e72c65194ba4a4c.
//
// Solidity: event ChequeEvent(uint256 indexed chequeID, address indexed from, address indexed to, address tokenAddress, uint256 amount, uint8 status)
func (_IPay *IPayFilterer) FilterChequeEvent(opts *bind.FilterOpts, chequeID []*big.Int, from []common.Address, to []common.Address) (*IPayChequeEventIterator, error) {

	var chequeIDRule []interface{}
	for _, chequeIDItem := range chequeID {
		chequeIDRule = append(chequeIDRule, chequeIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPay.contract.FilterLogs(opts, "ChequeEvent", chequeIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IPayChequeEventIterator{contract: _IPay.contract, event: "ChequeEvent", logs: logs, sub: sub}, nil
}

// WatchChequeEvent is a free log subscription operation binding the contract event 0x4617b59973e052c439b84edb2deaedd5c5dad563af6952409e72c65194ba4a4c.
//
// Solidity: event ChequeEvent(uint256 indexed chequeID, address indexed from, address indexed to, address tokenAddress, uint256 amount, uint8 status)
func (_IPay *IPayFilterer) WatchChequeEvent(opts *bind.WatchOpts, sink chan<- *IPayChequeEvent, chequeID []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var chequeIDRule []interface{}
	for _, chequeIDItem := range chequeID {
		chequeIDRule = append(chequeIDRule, chequeIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPay.contract.WatchLogs(opts, "ChequeEvent", chequeIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayChequeEvent)
				if err := _IPay.contract.UnpackLog(event, "ChequeEvent", log); err != nil {
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

// ParseChequeEvent is a log parse operation binding the contract event 0x4617b59973e052c439b84edb2deaedd5c5dad563af6952409e72c65194ba4a4c.
//
// Solidity: event ChequeEvent(uint256 indexed chequeID, address indexed from, address indexed to, address tokenAddress, uint256 amount, uint8 status)
func (_IPay *IPayFilterer) ParseChequeEvent(log types.Log) (*IPayChequeEvent, error) {
	event := new(IPayChequeEvent)
	if err := _IPay.contract.UnpackLog(event, "ChequeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPayChequeSentIterator is returned from FilterChequeSent and is used to iterate over the raw logs and unpacked data for ChequeSent events raised by the IPay contract.
type IPayChequeSentIterator struct {
	Event *IPayChequeSent // Event containing the contract specifics and raw log

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
func (it *IPayChequeSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayChequeSent)
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
		it.Event = new(IPayChequeSent)
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
func (it *IPayChequeSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayChequeSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayChequeSent represents a ChequeSent event raised by the IPay contract.
type IPayChequeSent struct {
	ChequeID     *big.Int
	From         common.Address
	To           common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Expiration   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChequeSent is a free log retrieval operation binding the contract event 0xd18cb5c9f5533f163598578532d03f4b1fb74decf1fb9a23f7d16defc21dda05.
//
// Solidity: event ChequeSent(uint256 indexed chequeID, address indexed from, address indexed to, address tokenAddress, uint256 amount, uint256 expiration)
func (_IPay *IPayFilterer) FilterChequeSent(opts *bind.FilterOpts, chequeID []*big.Int, from []common.Address, to []common.Address) (*IPayChequeSentIterator, error) {

	var chequeIDRule []interface{}
	for _, chequeIDItem := range chequeID {
		chequeIDRule = append(chequeIDRule, chequeIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPay.contract.FilterLogs(opts, "ChequeSent", chequeIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IPayChequeSentIterator{contract: _IPay.contract, event: "ChequeSent", logs: logs, sub: sub}, nil
}

// WatchChequeSent is a free log subscription operation binding the contract event 0xd18cb5c9f5533f163598578532d03f4b1fb74decf1fb9a23f7d16defc21dda05.
//
// Solidity: event ChequeSent(uint256 indexed chequeID, address indexed from, address indexed to, address tokenAddress, uint256 amount, uint256 expiration)
func (_IPay *IPayFilterer) WatchChequeSent(opts *bind.WatchOpts, sink chan<- *IPayChequeSent, chequeID []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var chequeIDRule []interface{}
	for _, chequeIDItem := range chequeID {
		chequeIDRule = append(chequeIDRule, chequeIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPay.contract.WatchLogs(opts, "ChequeSent", chequeIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayChequeSent)
				if err := _IPay.contract.UnpackLog(event, "ChequeSent", log); err != nil {
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

// ParseChequeSent is a log parse operation binding the contract event 0xd18cb5c9f5533f163598578532d03f4b1fb74decf1fb9a23f7d16defc21dda05.
//
// Solidity: event ChequeSent(uint256 indexed chequeID, address indexed from, address indexed to, address tokenAddress, uint256 amount, uint256 expiration)
func (_IPay *IPayFilterer) ParseChequeSent(log types.Log) (*IPayChequeSent, error) {
	event := new(IPayChequeSent)
	if err := _IPay.contract.UnpackLog(event, "ChequeSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPayWhitelistTokenAddedIterator is returned from FilterWhitelistTokenAdded and is used to iterate over the raw logs and unpacked data for WhitelistTokenAdded events raised by the IPay contract.
type IPayWhitelistTokenAddedIterator struct {
	Event *IPayWhitelistTokenAdded // Event containing the contract specifics and raw log

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
func (it *IPayWhitelistTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayWhitelistTokenAdded)
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
		it.Event = new(IPayWhitelistTokenAdded)
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
func (it *IPayWhitelistTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayWhitelistTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayWhitelistTokenAdded represents a WhitelistTokenAdded event raised by the IPay contract.
type IPayWhitelistTokenAdded struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWhitelistTokenAdded is a free log retrieval operation binding the contract event 0x9f9d9f44bc581f34670150874e4db46bb05da1753d75c6b5c7b952fcdaa96f4a.
//
// Solidity: event WhitelistTokenAdded(address[] tokens)
func (_IPay *IPayFilterer) FilterWhitelistTokenAdded(opts *bind.FilterOpts) (*IPayWhitelistTokenAddedIterator, error) {

	logs, sub, err := _IPay.contract.FilterLogs(opts, "WhitelistTokenAdded")
	if err != nil {
		return nil, err
	}
	return &IPayWhitelistTokenAddedIterator{contract: _IPay.contract, event: "WhitelistTokenAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistTokenAdded is a free log subscription operation binding the contract event 0x9f9d9f44bc581f34670150874e4db46bb05da1753d75c6b5c7b952fcdaa96f4a.
//
// Solidity: event WhitelistTokenAdded(address[] tokens)
func (_IPay *IPayFilterer) WatchWhitelistTokenAdded(opts *bind.WatchOpts, sink chan<- *IPayWhitelistTokenAdded) (event.Subscription, error) {

	logs, sub, err := _IPay.contract.WatchLogs(opts, "WhitelistTokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayWhitelistTokenAdded)
				if err := _IPay.contract.UnpackLog(event, "WhitelistTokenAdded", log); err != nil {
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

// ParseWhitelistTokenAdded is a log parse operation binding the contract event 0x9f9d9f44bc581f34670150874e4db46bb05da1753d75c6b5c7b952fcdaa96f4a.
//
// Solidity: event WhitelistTokenAdded(address[] tokens)
func (_IPay *IPayFilterer) ParseWhitelistTokenAdded(log types.Log) (*IPayWhitelistTokenAdded, error) {
	event := new(IPayWhitelistTokenAdded)
	if err := _IPay.contract.UnpackLog(event, "WhitelistTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPayWhitelistTokenRemovedIterator is returned from FilterWhitelistTokenRemoved and is used to iterate over the raw logs and unpacked data for WhitelistTokenRemoved events raised by the IPay contract.
type IPayWhitelistTokenRemovedIterator struct {
	Event *IPayWhitelistTokenRemoved // Event containing the contract specifics and raw log

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
func (it *IPayWhitelistTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayWhitelistTokenRemoved)
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
		it.Event = new(IPayWhitelistTokenRemoved)
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
func (it *IPayWhitelistTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayWhitelistTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayWhitelistTokenRemoved represents a WhitelistTokenRemoved event raised by the IPay contract.
type IPayWhitelistTokenRemoved struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWhitelistTokenRemoved is a free log retrieval operation binding the contract event 0xa50261031b405a768d9301886abe947b7a26918759bc9b5c89a4398b7dcff16e.
//
// Solidity: event WhitelistTokenRemoved(address[] tokens)
func (_IPay *IPayFilterer) FilterWhitelistTokenRemoved(opts *bind.FilterOpts) (*IPayWhitelistTokenRemovedIterator, error) {

	logs, sub, err := _IPay.contract.FilterLogs(opts, "WhitelistTokenRemoved")
	if err != nil {
		return nil, err
	}
	return &IPayWhitelistTokenRemovedIterator{contract: _IPay.contract, event: "WhitelistTokenRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistTokenRemoved is a free log subscription operation binding the contract event 0xa50261031b405a768d9301886abe947b7a26918759bc9b5c89a4398b7dcff16e.
//
// Solidity: event WhitelistTokenRemoved(address[] tokens)
func (_IPay *IPayFilterer) WatchWhitelistTokenRemoved(opts *bind.WatchOpts, sink chan<- *IPayWhitelistTokenRemoved) (event.Subscription, error) {

	logs, sub, err := _IPay.contract.WatchLogs(opts, "WhitelistTokenRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayWhitelistTokenRemoved)
				if err := _IPay.contract.UnpackLog(event, "WhitelistTokenRemoved", log); err != nil {
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

// ParseWhitelistTokenRemoved is a log parse operation binding the contract event 0xa50261031b405a768d9301886abe947b7a26918759bc9b5c89a4398b7dcff16e.
//
// Solidity: event WhitelistTokenRemoved(address[] tokens)
func (_IPay *IPayFilterer) ParseWhitelistTokenRemoved(log types.Log) (*IPayWhitelistTokenRemoved, error) {
	event := new(IPayWhitelistTokenRemoved)
	if err := _IPay.contract.UnpackLog(event, "WhitelistTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
