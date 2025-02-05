// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testtoken20

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

// TestToken20MetaData contains all meta data concerning the TestToken20 contract.
var TestToken20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080346102ea576040906001600160401b0390808301828111828210176101f6578352600b81526020916a054657374546f6b656e32360ac1b838301528351848101818110838211176101f65785526005815264054535432360dc1b8482015282518281116101f6576003918254916001958684811c941680156102e0575b888510146102ca578190601f94858111610279575b5088908583116001146102175760009261020c575b505060001982861b1c191690861b1783555b80519384116101f65760049586548681811c911680156101ec575b828210146101d757838111610191575b5080928511600114610128575093839491849260009561011d575b50501b92600019911b1c19161790555b5161075690816102f08239f35b015193503880610100565b92919084601f1981168860005285600020956000905b89838310610177575050501061015d575b50505050811b019055610110565b01519060f884600019921b161c191690553880808061014f565b85870151895590970196948501948893509081019061013e565b87600052816000208480880160051c8201928489106101ce575b0160051c019087905b8281106101c25750506100e5565b600081550187906101b4565b925081926101ab565b602288634e487b7160e01b6000525260246000fd5b90607f16906100d5565b634e487b7160e01b600052604160045260246000fd5b0151905038806100a8565b90889350601f19831691876000528a6000209260005b8c828210610263575050841161024b575b505050811b0183556100ba565b015160001983881b60f8161c1916905538808061023e565b8385015186558c9790950194938401930161022d565b90915085600052886000208580850160051c8201928b86106102c1575b918a91869594930160051c01915b8281106102b2575050610093565b600081558594508a91016102a4565b92508192610296565b634e487b7160e01b600052602260045260246000fd5b93607f169361007e565b600080fdfe608060408181526004918236101561001657600080fd5b600092833560e01c91826306fdde03146104ee57508163095ea7b31461044457816318160ddd1461042557816323b872dd1461032e578163313ce5671461031257816340c10f191461026657816370a082311461022f57816395d89b411461010e57508063a9059cbb146100de5763dd62ed3e1461009357600080fd5b346100da57806003193601126100da57806020926100af610611565b6100b761062c565b6001600160a01b0391821683526001865283832091168252845220549051908152f35b5080fd5b50346100da57806003193601126100da576020906101076100fd610611565b6024359033610642565b5160018152f35b8383346100da57816003193601126100da5780519082845460018160011c9060018316928315610225575b6020938484108114610212578388529081156101f657506001146101a1575b505050829003601f01601f191682019267ffffffffffffffff84118385101761018e575082918261018a9252826105c8565b0390f35b634e487b7160e01b815260418552602490fd5b8787529192508591837f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5b8385106101e25750505050830101858080610158565b8054888601830152930192849082016101cc565b60ff1916878501525050151560051b8401019050858080610158565b634e487b7160e01b895260228a52602489fd5b91607f1691610139565b5050346100da5760203660031901126100da5760209181906001600160a01b03610257610611565b16815280845220549051908152f35b9190503461030e578060031936011261030e57610281610611565b6001600160a01b031691602435919083156102f957600254908382018092116102e6575084927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9260209260025585855284835280852082815401905551908152a380f35b634e487b7160e01b865260119052602485fd5b84602492519163ec442f0560e01b8352820152fd5b8280fd5b5050346100da57816003193601126100da576020905160128152f35b905082346104225760603660031901126104225761034a610611565b61035261062c565b916044359360018060a01b03831680835260016020528683203384526020528683205491600019830361038e575b602088610107898989610642565b8683106103f65781156103df5733156103c85750825260016020908152868320338452815291869020908590039055829061010787610380565b8751634a1406b160e11b8152908101849052602490fd5b875163e602df0560e01b8152908101849052602490fd5b8751637dc7a0d960e11b8152339181019182526020820193909352604081018790528291506060010390fd5b80fd5b5050346100da57816003193601126100da576020906002549051908152f35b90503461030e578160031936011261030e5761045e610611565b6024359033156104d7576001600160a01b03169182156104c057508083602095338152600187528181208582528752205582519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925843392a35160018152f35b8351634a1406b160e11b8152908101859052602490fd5b835163e602df0560e01b8152808401869052602490fd5b8490843461030e578260031936011261030e578260035460018160011c90600183169283156105be575b6020938484108114610212578388529081156101f6575060011461056857505050829003601f01601f191682019267ffffffffffffffff84118385101761018e575082918261018a9252826105c8565b600387529192508591837fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b8385106105aa5750505050830101858080610158565b805488860183015293019284908201610594565b91607f1691610518565b6020808252825181830181905290939260005b8281106105fd57505060409293506000838284010152601f8019910116010190565b8181018601518482016040015285016105db565b600435906001600160a01b038216820361062757565b600080fd5b602435906001600160a01b038216820361062757565b916001600160a01b0380841692831561070757169283156106ee57600090838252816020526040822054908382106106bc575091604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815220818154019055604051908152a3565b60405163391434e360e21b81526001600160a01b03919091166004820152602481019190915260448101839052606490fd5b60405163ec442f0560e01b815260006004820152602490fd5b604051634b637e8f60e11b815260006004820152602490fdfea2646970667358221220bb6689ba8b696129fa6c01c5e1bbb8648eede3a79994c9cdae619b4402f8a8d364736f6c63430008190033",
}

// TestToken20ABI is the input ABI used to generate the binding from.
// Deprecated: Use TestToken20MetaData.ABI instead.
var TestToken20ABI = TestToken20MetaData.ABI

// TestToken20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestToken20MetaData.Bin instead.
var TestToken20Bin = TestToken20MetaData.Bin

// DeployTestToken20 deploys a new Ethereum contract, binding an instance of TestToken20 to it.
func DeployTestToken20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestToken20, error) {
	parsed, err := TestToken20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestToken20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestToken20{TestToken20Caller: TestToken20Caller{contract: contract}, TestToken20Transactor: TestToken20Transactor{contract: contract}, TestToken20Filterer: TestToken20Filterer{contract: contract}}, nil
}

// TestToken20 is an auto generated Go binding around an Ethereum contract.
type TestToken20 struct {
	TestToken20Caller     // Read-only binding to the contract
	TestToken20Transactor // Write-only binding to the contract
	TestToken20Filterer   // Log filterer for contract events
}

// TestToken20Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestToken20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestToken20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestToken20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestToken20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestToken20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestToken20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestToken20Session struct {
	Contract     *TestToken20      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestToken20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestToken20CallerSession struct {
	Contract *TestToken20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestToken20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestToken20TransactorSession struct {
	Contract     *TestToken20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestToken20Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestToken20Raw struct {
	Contract *TestToken20 // Generic contract binding to access the raw methods on
}

// TestToken20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestToken20CallerRaw struct {
	Contract *TestToken20Caller // Generic read-only contract binding to access the raw methods on
}

// TestToken20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestToken20TransactorRaw struct {
	Contract *TestToken20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestToken20 creates a new instance of TestToken20, bound to a specific deployed contract.
func NewTestToken20(address common.Address, backend bind.ContractBackend) (*TestToken20, error) {
	contract, err := bindTestToken20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestToken20{TestToken20Caller: TestToken20Caller{contract: contract}, TestToken20Transactor: TestToken20Transactor{contract: contract}, TestToken20Filterer: TestToken20Filterer{contract: contract}}, nil
}

// NewTestToken20Caller creates a new read-only instance of TestToken20, bound to a specific deployed contract.
func NewTestToken20Caller(address common.Address, caller bind.ContractCaller) (*TestToken20Caller, error) {
	contract, err := bindTestToken20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestToken20Caller{contract: contract}, nil
}

// NewTestToken20Transactor creates a new write-only instance of TestToken20, bound to a specific deployed contract.
func NewTestToken20Transactor(address common.Address, transactor bind.ContractTransactor) (*TestToken20Transactor, error) {
	contract, err := bindTestToken20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestToken20Transactor{contract: contract}, nil
}

// NewTestToken20Filterer creates a new log filterer instance of TestToken20, bound to a specific deployed contract.
func NewTestToken20Filterer(address common.Address, filterer bind.ContractFilterer) (*TestToken20Filterer, error) {
	contract, err := bindTestToken20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestToken20Filterer{contract: contract}, nil
}

// bindTestToken20 binds a generic wrapper to an already deployed contract.
func bindTestToken20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestToken20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestToken20 *TestToken20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestToken20.Contract.TestToken20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestToken20 *TestToken20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestToken20.Contract.TestToken20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestToken20 *TestToken20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestToken20.Contract.TestToken20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestToken20 *TestToken20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestToken20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestToken20 *TestToken20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestToken20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestToken20 *TestToken20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestToken20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestToken20 *TestToken20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestToken20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestToken20 *TestToken20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestToken20.Contract.Allowance(&_TestToken20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestToken20 *TestToken20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestToken20.Contract.Allowance(&_TestToken20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestToken20 *TestToken20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestToken20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestToken20 *TestToken20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestToken20.Contract.BalanceOf(&_TestToken20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestToken20 *TestToken20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestToken20.Contract.BalanceOf(&_TestToken20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestToken20 *TestToken20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestToken20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestToken20 *TestToken20Session) Decimals() (uint8, error) {
	return _TestToken20.Contract.Decimals(&_TestToken20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestToken20 *TestToken20CallerSession) Decimals() (uint8, error) {
	return _TestToken20.Contract.Decimals(&_TestToken20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken20 *TestToken20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestToken20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken20 *TestToken20Session) Name() (string, error) {
	return _TestToken20.Contract.Name(&_TestToken20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken20 *TestToken20CallerSession) Name() (string, error) {
	return _TestToken20.Contract.Name(&_TestToken20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken20 *TestToken20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestToken20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken20 *TestToken20Session) Symbol() (string, error) {
	return _TestToken20.Contract.Symbol(&_TestToken20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken20 *TestToken20CallerSession) Symbol() (string, error) {
	return _TestToken20.Contract.Symbol(&_TestToken20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestToken20 *TestToken20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestToken20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestToken20 *TestToken20Session) TotalSupply() (*big.Int, error) {
	return _TestToken20.Contract.TotalSupply(&_TestToken20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestToken20 *TestToken20CallerSession) TotalSupply() (*big.Int, error) {
	return _TestToken20.Contract.TotalSupply(&_TestToken20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_TestToken20 *TestToken20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestToken20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_TestToken20 *TestToken20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestToken20.Contract.Approve(&_TestToken20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_TestToken20 *TestToken20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestToken20.Contract.Approve(&_TestToken20.TransactOpts, spender, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_TestToken20 *TestToken20Transactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken20.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_TestToken20 *TestToken20Session) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken20.Contract.Mint(&_TestToken20.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_TestToken20 *TestToken20TransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken20.Contract.Mint(&_TestToken20.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_TestToken20 *TestToken20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestToken20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_TestToken20 *TestToken20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestToken20.Contract.Transfer(&_TestToken20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_TestToken20 *TestToken20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestToken20.Contract.Transfer(&_TestToken20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_TestToken20 *TestToken20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestToken20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_TestToken20 *TestToken20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestToken20.Contract.TransferFrom(&_TestToken20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_TestToken20 *TestToken20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestToken20.Contract.TransferFrom(&_TestToken20.TransactOpts, from, to, value)
}

// TestToken20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestToken20 contract.
type TestToken20ApprovalIterator struct {
	Event *TestToken20Approval // Event containing the contract specifics and raw log

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
func (it *TestToken20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestToken20Approval)
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
		it.Event = new(TestToken20Approval)
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
func (it *TestToken20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestToken20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestToken20Approval represents a Approval event raised by the TestToken20 contract.
type TestToken20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestToken20 *TestToken20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TestToken20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestToken20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TestToken20ApprovalIterator{contract: _TestToken20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestToken20 *TestToken20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestToken20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestToken20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestToken20Approval)
				if err := _TestToken20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestToken20 *TestToken20Filterer) ParseApproval(log types.Log) (*TestToken20Approval, error) {
	event := new(TestToken20Approval)
	if err := _TestToken20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestToken20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestToken20 contract.
type TestToken20TransferIterator struct {
	Event *TestToken20Transfer // Event containing the contract specifics and raw log

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
func (it *TestToken20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestToken20Transfer)
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
		it.Event = new(TestToken20Transfer)
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
func (it *TestToken20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestToken20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestToken20Transfer represents a Transfer event raised by the TestToken20 contract.
type TestToken20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestToken20 *TestToken20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestToken20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestToken20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestToken20TransferIterator{contract: _TestToken20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestToken20 *TestToken20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestToken20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestToken20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestToken20Transfer)
				if err := _TestToken20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestToken20 *TestToken20Filterer) ParseTransfer(log types.Log) (*TestToken20Transfer, error) {
	event := new(TestToken20Transfer)
	if err := _TestToken20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
