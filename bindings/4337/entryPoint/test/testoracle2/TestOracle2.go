// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testoracle2

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

// TestOracle2MetaData contains all meta data concerning the TestOracle2 contract.
var TestOracle2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_price\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"setDecimals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_price\",\"type\":\"int256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608034606957601f6101e038819003918201601f19168301916001600160401b03831184841017606e578084926040948552833981010312606957602081519101519060ff821680920360695760005560ff19600154161760015560405161015b90816100858239f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b600090813560e01c908163313ce56714610106575080637a1395aa146100d4578063a035b1fe146100b7578063f7a308061461009d5763feaf968c1461005857600080fd5b3461009a578060031936011261009a5760a09054604051906804000000000000247a90818352602083015263642a887b60408301524260608301526080820152f35b80fd5b503461009a57602036600319011261009a57600435815580f35b503461009a578060031936011261009a5760209054604051908152f35b503461009a57602036600319011261009a5760043560ff81168091036101025760ff19600154161760015580f35b5080fd5b90503461010257816003193601126101025760209060ff600154168152f3fea26469706673582212201b636613e7775b6376dc34816e22aa876a61c03d4df2bd60425732d707163cd364736f6c63430008190033",
}

// TestOracle2ABI is the input ABI used to generate the binding from.
// Deprecated: Use TestOracle2MetaData.ABI instead.
var TestOracle2ABI = TestOracle2MetaData.ABI

// TestOracle2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestOracle2MetaData.Bin instead.
var TestOracle2Bin = TestOracle2MetaData.Bin

// DeployTestOracle2 deploys a new Ethereum contract, binding an instance of TestOracle2 to it.
func DeployTestOracle2(auth *bind.TransactOpts, backend bind.ContractBackend, _price *big.Int, _decimals uint8) (common.Address, *types.Transaction, *TestOracle2, error) {
	parsed, err := TestOracle2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestOracle2Bin), backend, _price, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestOracle2{TestOracle2Caller: TestOracle2Caller{contract: contract}, TestOracle2Transactor: TestOracle2Transactor{contract: contract}, TestOracle2Filterer: TestOracle2Filterer{contract: contract}}, nil
}

// TestOracle2 is an auto generated Go binding around an Ethereum contract.
type TestOracle2 struct {
	TestOracle2Caller     // Read-only binding to the contract
	TestOracle2Transactor // Write-only binding to the contract
	TestOracle2Filterer   // Log filterer for contract events
}

// TestOracle2Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestOracle2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOracle2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestOracle2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOracle2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestOracle2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOracle2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestOracle2Session struct {
	Contract     *TestOracle2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestOracle2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestOracle2CallerSession struct {
	Contract *TestOracle2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestOracle2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestOracle2TransactorSession struct {
	Contract     *TestOracle2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestOracle2Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestOracle2Raw struct {
	Contract *TestOracle2 // Generic contract binding to access the raw methods on
}

// TestOracle2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestOracle2CallerRaw struct {
	Contract *TestOracle2Caller // Generic read-only contract binding to access the raw methods on
}

// TestOracle2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestOracle2TransactorRaw struct {
	Contract *TestOracle2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestOracle2 creates a new instance of TestOracle2, bound to a specific deployed contract.
func NewTestOracle2(address common.Address, backend bind.ContractBackend) (*TestOracle2, error) {
	contract, err := bindTestOracle2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestOracle2{TestOracle2Caller: TestOracle2Caller{contract: contract}, TestOracle2Transactor: TestOracle2Transactor{contract: contract}, TestOracle2Filterer: TestOracle2Filterer{contract: contract}}, nil
}

// NewTestOracle2Caller creates a new read-only instance of TestOracle2, bound to a specific deployed contract.
func NewTestOracle2Caller(address common.Address, caller bind.ContractCaller) (*TestOracle2Caller, error) {
	contract, err := bindTestOracle2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestOracle2Caller{contract: contract}, nil
}

// NewTestOracle2Transactor creates a new write-only instance of TestOracle2, bound to a specific deployed contract.
func NewTestOracle2Transactor(address common.Address, transactor bind.ContractTransactor) (*TestOracle2Transactor, error) {
	contract, err := bindTestOracle2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestOracle2Transactor{contract: contract}, nil
}

// NewTestOracle2Filterer creates a new log filterer instance of TestOracle2, bound to a specific deployed contract.
func NewTestOracle2Filterer(address common.Address, filterer bind.ContractFilterer) (*TestOracle2Filterer, error) {
	contract, err := bindTestOracle2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestOracle2Filterer{contract: contract}, nil
}

// bindTestOracle2 binds a generic wrapper to an already deployed contract.
func bindTestOracle2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestOracle2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestOracle2 *TestOracle2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOracle2.Contract.TestOracle2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestOracle2 *TestOracle2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOracle2.Contract.TestOracle2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestOracle2 *TestOracle2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOracle2.Contract.TestOracle2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestOracle2 *TestOracle2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOracle2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestOracle2 *TestOracle2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOracle2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestOracle2 *TestOracle2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOracle2.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestOracle2 *TestOracle2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestOracle2.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestOracle2 *TestOracle2Session) Decimals() (uint8, error) {
	return _TestOracle2.Contract.Decimals(&_TestOracle2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestOracle2 *TestOracle2CallerSession) Decimals() (uint8, error) {
	return _TestOracle2.Contract.Decimals(&_TestOracle2.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_TestOracle2 *TestOracle2Caller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _TestOracle2.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_TestOracle2 *TestOracle2Session) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _TestOracle2.Contract.LatestRoundData(&_TestOracle2.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_TestOracle2 *TestOracle2CallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _TestOracle2.Contract.LatestRoundData(&_TestOracle2.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(int256)
func (_TestOracle2 *TestOracle2Caller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOracle2.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(int256)
func (_TestOracle2 *TestOracle2Session) Price() (*big.Int, error) {
	return _TestOracle2.Contract.Price(&_TestOracle2.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(int256)
func (_TestOracle2 *TestOracle2CallerSession) Price() (*big.Int, error) {
	return _TestOracle2.Contract.Price(&_TestOracle2.CallOpts)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x7a1395aa.
//
// Solidity: function setDecimals(uint8 _decimals) returns()
func (_TestOracle2 *TestOracle2Transactor) SetDecimals(opts *bind.TransactOpts, _decimals uint8) (*types.Transaction, error) {
	return _TestOracle2.contract.Transact(opts, "setDecimals", _decimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x7a1395aa.
//
// Solidity: function setDecimals(uint8 _decimals) returns()
func (_TestOracle2 *TestOracle2Session) SetDecimals(_decimals uint8) (*types.Transaction, error) {
	return _TestOracle2.Contract.SetDecimals(&_TestOracle2.TransactOpts, _decimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x7a1395aa.
//
// Solidity: function setDecimals(uint8 _decimals) returns()
func (_TestOracle2 *TestOracle2TransactorSession) SetDecimals(_decimals uint8) (*types.Transaction, error) {
	return _TestOracle2.Contract.SetDecimals(&_TestOracle2.TransactOpts, _decimals)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7a30806.
//
// Solidity: function setPrice(int256 _price) returns()
func (_TestOracle2 *TestOracle2Transactor) SetPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _TestOracle2.contract.Transact(opts, "setPrice", _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7a30806.
//
// Solidity: function setPrice(int256 _price) returns()
func (_TestOracle2 *TestOracle2Session) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _TestOracle2.Contract.SetPrice(&_TestOracle2.TransactOpts, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7a30806.
//
// Solidity: function setPrice(int256 _price) returns()
func (_TestOracle2 *TestOracle2TransactorSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _TestOracle2.Contract.SetPrice(&_TestOracle2.TransactOpts, _price)
}
