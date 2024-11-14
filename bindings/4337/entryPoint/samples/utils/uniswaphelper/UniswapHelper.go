// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswaphelper

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

// UniswapHelperMetaData contains all meta data concerning the UniswapHelper contract.
var UniswapHelperMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"name\":\"UniswapReverted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"tokenToWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswap\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"weiToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedNative\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapHelperMetaData.ABI instead.
var UniswapHelperABI = UniswapHelperMetaData.ABI

// UniswapHelper is an auto generated Go binding around an Ethereum contract.
type UniswapHelper struct {
	UniswapHelperCaller     // Read-only binding to the contract
	UniswapHelperTransactor // Write-only binding to the contract
	UniswapHelperFilterer   // Log filterer for contract events
}

// UniswapHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapHelperSession struct {
	Contract     *UniswapHelper    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapHelperCallerSession struct {
	Contract *UniswapHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UniswapHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapHelperTransactorSession struct {
	Contract     *UniswapHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UniswapHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapHelperRaw struct {
	Contract *UniswapHelper // Generic contract binding to access the raw methods on
}

// UniswapHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapHelperCallerRaw struct {
	Contract *UniswapHelperCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapHelperTransactorRaw struct {
	Contract *UniswapHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapHelper creates a new instance of UniswapHelper, bound to a specific deployed contract.
func NewUniswapHelper(address common.Address, backend bind.ContractBackend) (*UniswapHelper, error) {
	contract, err := bindUniswapHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapHelper{UniswapHelperCaller: UniswapHelperCaller{contract: contract}, UniswapHelperTransactor: UniswapHelperTransactor{contract: contract}, UniswapHelperFilterer: UniswapHelperFilterer{contract: contract}}, nil
}

// NewUniswapHelperCaller creates a new read-only instance of UniswapHelper, bound to a specific deployed contract.
func NewUniswapHelperCaller(address common.Address, caller bind.ContractCaller) (*UniswapHelperCaller, error) {
	contract, err := bindUniswapHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapHelperCaller{contract: contract}, nil
}

// NewUniswapHelperTransactor creates a new write-only instance of UniswapHelper, bound to a specific deployed contract.
func NewUniswapHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapHelperTransactor, error) {
	contract, err := bindUniswapHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapHelperTransactor{contract: contract}, nil
}

// NewUniswapHelperFilterer creates a new log filterer instance of UniswapHelper, bound to a specific deployed contract.
func NewUniswapHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapHelperFilterer, error) {
	contract, err := bindUniswapHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapHelperFilterer{contract: contract}, nil
}

// bindUniswapHelper binds a generic wrapper to an already deployed contract.
func bindUniswapHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapHelper *UniswapHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapHelper.Contract.UniswapHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapHelper *UniswapHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapHelper.Contract.UniswapHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapHelper *UniswapHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapHelper.Contract.UniswapHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapHelper *UniswapHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapHelper *UniswapHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapHelper *UniswapHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapHelper.Contract.contract.Transact(opts, method, params...)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_UniswapHelper *UniswapHelperCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapHelper.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_UniswapHelper *UniswapHelperSession) Token() (common.Address, error) {
	return _UniswapHelper.Contract.Token(&_UniswapHelper.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_UniswapHelper *UniswapHelperCallerSession) Token() (common.Address, error) {
	return _UniswapHelper.Contract.Token(&_UniswapHelper.CallOpts)
}

// TokenToWei is a free data retrieval call binding the contract method 0xd7a23b3c.
//
// Solidity: function tokenToWei(uint256 amount, uint256 price) pure returns(uint256)
func (_UniswapHelper *UniswapHelperCaller) TokenToWei(opts *bind.CallOpts, amount *big.Int, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapHelper.contract.Call(opts, &out, "tokenToWei", amount, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToWei is a free data retrieval call binding the contract method 0xd7a23b3c.
//
// Solidity: function tokenToWei(uint256 amount, uint256 price) pure returns(uint256)
func (_UniswapHelper *UniswapHelperSession) TokenToWei(amount *big.Int, price *big.Int) (*big.Int, error) {
	return _UniswapHelper.Contract.TokenToWei(&_UniswapHelper.CallOpts, amount, price)
}

// TokenToWei is a free data retrieval call binding the contract method 0xd7a23b3c.
//
// Solidity: function tokenToWei(uint256 amount, uint256 price) pure returns(uint256)
func (_UniswapHelper *UniswapHelperCallerSession) TokenToWei(amount *big.Int, price *big.Int) (*big.Int, error) {
	return _UniswapHelper.Contract.TokenToWei(&_UniswapHelper.CallOpts, amount, price)
}

// Uniswap is a free data retrieval call binding the contract method 0x2681f7e4.
//
// Solidity: function uniswap() view returns(address)
func (_UniswapHelper *UniswapHelperCaller) Uniswap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapHelper.contract.Call(opts, &out, "uniswap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uniswap is a free data retrieval call binding the contract method 0x2681f7e4.
//
// Solidity: function uniswap() view returns(address)
func (_UniswapHelper *UniswapHelperSession) Uniswap() (common.Address, error) {
	return _UniswapHelper.Contract.Uniswap(&_UniswapHelper.CallOpts)
}

// Uniswap is a free data retrieval call binding the contract method 0x2681f7e4.
//
// Solidity: function uniswap() view returns(address)
func (_UniswapHelper *UniswapHelperCallerSession) Uniswap() (common.Address, error) {
	return _UniswapHelper.Contract.Uniswap(&_UniswapHelper.CallOpts)
}

// WeiToToken is a free data retrieval call binding the contract method 0x7c986aac.
//
// Solidity: function weiToToken(uint256 amount, uint256 price) pure returns(uint256)
func (_UniswapHelper *UniswapHelperCaller) WeiToToken(opts *bind.CallOpts, amount *big.Int, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapHelper.contract.Call(opts, &out, "weiToToken", amount, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiToToken is a free data retrieval call binding the contract method 0x7c986aac.
//
// Solidity: function weiToToken(uint256 amount, uint256 price) pure returns(uint256)
func (_UniswapHelper *UniswapHelperSession) WeiToToken(amount *big.Int, price *big.Int) (*big.Int, error) {
	return _UniswapHelper.Contract.WeiToToken(&_UniswapHelper.CallOpts, amount, price)
}

// WeiToToken is a free data retrieval call binding the contract method 0x7c986aac.
//
// Solidity: function weiToToken(uint256 amount, uint256 price) pure returns(uint256)
func (_UniswapHelper *UniswapHelperCallerSession) WeiToToken(amount *big.Int, price *big.Int) (*big.Int, error) {
	return _UniswapHelper.Contract.WeiToToken(&_UniswapHelper.CallOpts, amount, price)
}

// WrappedNative is a free data retrieval call binding the contract method 0xeb6d3a11.
//
// Solidity: function wrappedNative() view returns(address)
func (_UniswapHelper *UniswapHelperCaller) WrappedNative(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapHelper.contract.Call(opts, &out, "wrappedNative")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedNative is a free data retrieval call binding the contract method 0xeb6d3a11.
//
// Solidity: function wrappedNative() view returns(address)
func (_UniswapHelper *UniswapHelperSession) WrappedNative() (common.Address, error) {
	return _UniswapHelper.Contract.WrappedNative(&_UniswapHelper.CallOpts)
}

// WrappedNative is a free data retrieval call binding the contract method 0xeb6d3a11.
//
// Solidity: function wrappedNative() view returns(address)
func (_UniswapHelper *UniswapHelperCallerSession) WrappedNative() (common.Address, error) {
	return _UniswapHelper.Contract.WrappedNative(&_UniswapHelper.CallOpts)
}

// UniswapHelperUniswapRevertedIterator is returned from FilterUniswapReverted and is used to iterate over the raw logs and unpacked data for UniswapReverted events raised by the UniswapHelper contract.
type UniswapHelperUniswapRevertedIterator struct {
	Event *UniswapHelperUniswapReverted // Event containing the contract specifics and raw log

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
func (it *UniswapHelperUniswapRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapHelperUniswapReverted)
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
		it.Event = new(UniswapHelperUniswapReverted)
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
func (it *UniswapHelperUniswapRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapHelperUniswapRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapHelperUniswapReverted represents a UniswapReverted event raised by the UniswapHelper contract.
type UniswapHelperUniswapReverted struct {
	TokenIn      common.Address
	TokenOut     common.Address
	AmountIn     *big.Int
	AmountOutMin *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUniswapReverted is a free log retrieval operation binding the contract event 0xf7edd4c6ec425decf715a8b8eaa3b65d3d86e31ad0ff750aa60fa834190f515f.
//
// Solidity: event UniswapReverted(address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOutMin)
func (_UniswapHelper *UniswapHelperFilterer) FilterUniswapReverted(opts *bind.FilterOpts) (*UniswapHelperUniswapRevertedIterator, error) {

	logs, sub, err := _UniswapHelper.contract.FilterLogs(opts, "UniswapReverted")
	if err != nil {
		return nil, err
	}
	return &UniswapHelperUniswapRevertedIterator{contract: _UniswapHelper.contract, event: "UniswapReverted", logs: logs, sub: sub}, nil
}

// WatchUniswapReverted is a free log subscription operation binding the contract event 0xf7edd4c6ec425decf715a8b8eaa3b65d3d86e31ad0ff750aa60fa834190f515f.
//
// Solidity: event UniswapReverted(address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOutMin)
func (_UniswapHelper *UniswapHelperFilterer) WatchUniswapReverted(opts *bind.WatchOpts, sink chan<- *UniswapHelperUniswapReverted) (event.Subscription, error) {

	logs, sub, err := _UniswapHelper.contract.WatchLogs(opts, "UniswapReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapHelperUniswapReverted)
				if err := _UniswapHelper.contract.UnpackLog(event, "UniswapReverted", log); err != nil {
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

// ParseUniswapReverted is a log parse operation binding the contract event 0xf7edd4c6ec425decf715a8b8eaa3b65d3d86e31ad0ff750aa60fa834190f515f.
//
// Solidity: event UniswapReverted(address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOutMin)
func (_UniswapHelper *UniswapHelperFilterer) ParseUniswapReverted(log types.Log) (*UniswapHelperUniswapReverted, error) {
	event := new(UniswapHelperUniswapReverted)
	if err := _UniswapHelper.contract.UnpackLog(event, "UniswapReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
