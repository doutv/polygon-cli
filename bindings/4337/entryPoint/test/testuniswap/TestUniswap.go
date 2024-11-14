// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testuniswap

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

// ISwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// ISwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// TestUniswapMetaData contains all meta data concerning the TestUniswap contract.
var TestUniswapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractTestWrappedNativeToken\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"StubUniswapExchangeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractTestWrappedNativeToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608034607057601f61060a38819003918201601f19168301916001600160401b03831184841017607557808492602094604052833981010312607057516001600160a01b03811690819003607057600080546001600160a01b03191691909117905560405161057e908161008c8239f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe608060409080825260049182361015610023575b505050361561002157600080fd5b005b600091823560e01c9081633fc8cef31461048d57508063414bf3891461039457806349404b7c1461020d5763db3e2198036100135734610209576101003660031901126102095760c4359160041983019283116101f65760a4357f29a1b057b49189ad91f5d92216fc55e37caf33e398f9e914469523c9aa5f6a036100a66104b1565b6100ae6104cc565b8551878152602081018590526001600160a01b03928316604082015291166060820152608090a16001600160a01b0394856100e76104b1565b85516323b872dd60e01b81523384820190815230602080830191909152604082018a9052999391928a9284928390036060019183918a91165af180156101ec579188939161017f9695936101cf575b5061013f6104cc565b16908461014a6104e2565b885163a9059cbb60e01b81526001600160a01b0390911692810192835260208301949094529295869384929091839160400190565b03925af19081156101c45750610197575b5051908152f35b6101b690843d86116101bd575b6101ae81836104f8565b810190610530565b5038610190565b503d6101a4565b8351903d90823e3d90fd5b6101e590853d87116101bd576101ae81836104f8565b5038610136565b86513d87823e3d90fd5b634e487b7160e01b815260118452602490fd5b5080fd5b508281600319360112610312576024356001600160a01b0381811691829003610390578454168351926370a0823160e01b84523081850152602084602481855afa93841561038657869461034e575b50803584106103165785908415928315610274578280f35b803b156103125782906024885180958193632e1a7d4d60e01b83528a878401525af18015610308576102d4575b50508480938193829383906102cb575bf1156102c1578080808481808280f35b51903d90823e3d90fd5b506108fc6102b1565b67ffffffffffffffff829793949597116102f55750845293919085806102a1565b634e487b7160e01b835260419052602482fd5b86513d89823e3d90fd5b8280fd5b606490602086519162461bcd60e51b83528201526012602482015271496e73756666696369656e7420574554483960701b6044820152fd5b9093506020813d60201161037e575b8161036a602093836104f8565b8101031261037a5751928661025c565b8580fd5b3d915061035d565b85513d88823e3d90fd5b8480fd5b5034610209576101003660031901126102095760c43591600583018093116101f6576104629360a4357f29a1b057b49189ad91f5d92216fc55e37caf33e398f9e914469523c9aa5f6a036103e66104b1565b6103ee6104cc565b8651848152602081018990526001600160a01b03928316604082015291166060820152608090a16001600160a01b0390816104276104b1565b86516323b872dd60e01b815233868201908152306020828101919091526040820195909552939990938a938593169183918991839160600190565b03925af1801561048357928792879261017f96956101cf575061013f6104cc565b85513d86823e3d90fd5b83903461020957816003193601126102095790546001600160a01b03168152602090f35b6004356001600160a01b03811681036104c75790565b600080fd5b6024356001600160a01b03811681036104c75790565b6064356001600160a01b03811681036104c75790565b90601f8019910116810190811067ffffffffffffffff82111761051a57604052565b634e487b7160e01b600052604160045260246000fd5b908160209103126104c7575180151581036104c7579056fea26469706673582212201b50f90176ad40c61e1ef03ba299ca803cfa655d996e211b3616a45187cd4f5264736f6c63430008190033",
}

// TestUniswapABI is the input ABI used to generate the binding from.
// Deprecated: Use TestUniswapMetaData.ABI instead.
var TestUniswapABI = TestUniswapMetaData.ABI

// TestUniswapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestUniswapMetaData.Bin instead.
var TestUniswapBin = TestUniswapMetaData.Bin

// DeployTestUniswap deploys a new Ethereum contract, binding an instance of TestUniswap to it.
func DeployTestUniswap(auth *bind.TransactOpts, backend bind.ContractBackend, _weth common.Address) (common.Address, *types.Transaction, *TestUniswap, error) {
	parsed, err := TestUniswapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestUniswapBin), backend, _weth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestUniswap{TestUniswapCaller: TestUniswapCaller{contract: contract}, TestUniswapTransactor: TestUniswapTransactor{contract: contract}, TestUniswapFilterer: TestUniswapFilterer{contract: contract}}, nil
}

// TestUniswap is an auto generated Go binding around an Ethereum contract.
type TestUniswap struct {
	TestUniswapCaller     // Read-only binding to the contract
	TestUniswapTransactor // Write-only binding to the contract
	TestUniswapFilterer   // Log filterer for contract events
}

// TestUniswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestUniswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUniswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestUniswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUniswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestUniswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUniswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestUniswapSession struct {
	Contract     *TestUniswap      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestUniswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestUniswapCallerSession struct {
	Contract *TestUniswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestUniswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestUniswapTransactorSession struct {
	Contract     *TestUniswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestUniswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestUniswapRaw struct {
	Contract *TestUniswap // Generic contract binding to access the raw methods on
}

// TestUniswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestUniswapCallerRaw struct {
	Contract *TestUniswapCaller // Generic read-only contract binding to access the raw methods on
}

// TestUniswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestUniswapTransactorRaw struct {
	Contract *TestUniswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestUniswap creates a new instance of TestUniswap, bound to a specific deployed contract.
func NewTestUniswap(address common.Address, backend bind.ContractBackend) (*TestUniswap, error) {
	contract, err := bindTestUniswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestUniswap{TestUniswapCaller: TestUniswapCaller{contract: contract}, TestUniswapTransactor: TestUniswapTransactor{contract: contract}, TestUniswapFilterer: TestUniswapFilterer{contract: contract}}, nil
}

// NewTestUniswapCaller creates a new read-only instance of TestUniswap, bound to a specific deployed contract.
func NewTestUniswapCaller(address common.Address, caller bind.ContractCaller) (*TestUniswapCaller, error) {
	contract, err := bindTestUniswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestUniswapCaller{contract: contract}, nil
}

// NewTestUniswapTransactor creates a new write-only instance of TestUniswap, bound to a specific deployed contract.
func NewTestUniswapTransactor(address common.Address, transactor bind.ContractTransactor) (*TestUniswapTransactor, error) {
	contract, err := bindTestUniswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestUniswapTransactor{contract: contract}, nil
}

// NewTestUniswapFilterer creates a new log filterer instance of TestUniswap, bound to a specific deployed contract.
func NewTestUniswapFilterer(address common.Address, filterer bind.ContractFilterer) (*TestUniswapFilterer, error) {
	contract, err := bindTestUniswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestUniswapFilterer{contract: contract}, nil
}

// bindTestUniswap binds a generic wrapper to an already deployed contract.
func bindTestUniswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestUniswapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestUniswap *TestUniswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestUniswap.Contract.TestUniswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestUniswap *TestUniswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUniswap.Contract.TestUniswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestUniswap *TestUniswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestUniswap.Contract.TestUniswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestUniswap *TestUniswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestUniswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestUniswap *TestUniswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUniswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestUniswap *TestUniswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestUniswap.Contract.contract.Transact(opts, method, params...)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_TestUniswap *TestUniswapCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestUniswap.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_TestUniswap *TestUniswapSession) Weth() (common.Address, error) {
	return _TestUniswap.Contract.Weth(&_TestUniswap.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_TestUniswap *TestUniswapCallerSession) Weth() (common.Address, error) {
	return _TestUniswap.Contract.Weth(&_TestUniswap.CallOpts)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) returns(uint256)
func (_TestUniswap *TestUniswapTransactor) ExactInputSingle(opts *bind.TransactOpts, params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _TestUniswap.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) returns(uint256)
func (_TestUniswap *TestUniswapSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _TestUniswap.Contract.ExactInputSingle(&_TestUniswap.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) returns(uint256)
func (_TestUniswap *TestUniswapTransactorSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _TestUniswap.Contract.ExactInputSingle(&_TestUniswap.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) returns(uint256)
func (_TestUniswap *TestUniswapTransactor) ExactOutputSingle(opts *bind.TransactOpts, params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _TestUniswap.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) returns(uint256)
func (_TestUniswap *TestUniswapSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _TestUniswap.Contract.ExactOutputSingle(&_TestUniswap.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) returns(uint256)
func (_TestUniswap *TestUniswapTransactorSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _TestUniswap.Contract.ExactOutputSingle(&_TestUniswap.TransactOpts, params)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_TestUniswap *TestUniswapTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _TestUniswap.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_TestUniswap *TestUniswapSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _TestUniswap.Contract.UnwrapWETH9(&_TestUniswap.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_TestUniswap *TestUniswapTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _TestUniswap.Contract.UnwrapWETH9(&_TestUniswap.TransactOpts, amountMinimum, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestUniswap *TestUniswapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUniswap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestUniswap *TestUniswapSession) Receive() (*types.Transaction, error) {
	return _TestUniswap.Contract.Receive(&_TestUniswap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestUniswap *TestUniswapTransactorSession) Receive() (*types.Transaction, error) {
	return _TestUniswap.Contract.Receive(&_TestUniswap.TransactOpts)
}

// TestUniswapStubUniswapExchangeEventIterator is returned from FilterStubUniswapExchangeEvent and is used to iterate over the raw logs and unpacked data for StubUniswapExchangeEvent events raised by the TestUniswap contract.
type TestUniswapStubUniswapExchangeEventIterator struct {
	Event *TestUniswapStubUniswapExchangeEvent // Event containing the contract specifics and raw log

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
func (it *TestUniswapStubUniswapExchangeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestUniswapStubUniswapExchangeEvent)
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
		it.Event = new(TestUniswapStubUniswapExchangeEvent)
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
func (it *TestUniswapStubUniswapExchangeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestUniswapStubUniswapExchangeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestUniswapStubUniswapExchangeEvent represents a StubUniswapExchangeEvent event raised by the TestUniswap contract.
type TestUniswapStubUniswapExchangeEvent struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	TokenIn   common.Address
	TokenOut  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStubUniswapExchangeEvent is a free log retrieval operation binding the contract event 0x29a1b057b49189ad91f5d92216fc55e37caf33e398f9e914469523c9aa5f6a03.
//
// Solidity: event StubUniswapExchangeEvent(uint256 amountIn, uint256 amountOut, address tokenIn, address tokenOut)
func (_TestUniswap *TestUniswapFilterer) FilterStubUniswapExchangeEvent(opts *bind.FilterOpts) (*TestUniswapStubUniswapExchangeEventIterator, error) {

	logs, sub, err := _TestUniswap.contract.FilterLogs(opts, "StubUniswapExchangeEvent")
	if err != nil {
		return nil, err
	}
	return &TestUniswapStubUniswapExchangeEventIterator{contract: _TestUniswap.contract, event: "StubUniswapExchangeEvent", logs: logs, sub: sub}, nil
}

// WatchStubUniswapExchangeEvent is a free log subscription operation binding the contract event 0x29a1b057b49189ad91f5d92216fc55e37caf33e398f9e914469523c9aa5f6a03.
//
// Solidity: event StubUniswapExchangeEvent(uint256 amountIn, uint256 amountOut, address tokenIn, address tokenOut)
func (_TestUniswap *TestUniswapFilterer) WatchStubUniswapExchangeEvent(opts *bind.WatchOpts, sink chan<- *TestUniswapStubUniswapExchangeEvent) (event.Subscription, error) {

	logs, sub, err := _TestUniswap.contract.WatchLogs(opts, "StubUniswapExchangeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestUniswapStubUniswapExchangeEvent)
				if err := _TestUniswap.contract.UnpackLog(event, "StubUniswapExchangeEvent", log); err != nil {
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

// ParseStubUniswapExchangeEvent is a log parse operation binding the contract event 0x29a1b057b49189ad91f5d92216fc55e37caf33e398f9e914469523c9aa5f6a03.
//
// Solidity: event StubUniswapExchangeEvent(uint256 amountIn, uint256 amountOut, address tokenIn, address tokenOut)
func (_TestUniswap *TestUniswapFilterer) ParseStubUniswapExchangeEvent(log types.Log) (*TestUniswapStubUniswapExchangeEvent, error) {
	event := new(TestUniswapStubUniswapExchangeEvent)
	if err := _TestUniswap.contract.UnpackLog(event, "StubUniswapExchangeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
