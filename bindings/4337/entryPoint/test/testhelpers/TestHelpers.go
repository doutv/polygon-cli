// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testhelpers

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

// ValidationData is an auto generated low-level Go binding around an user-defined struct.
type ValidationData struct {
	Aggregator common.Address
	ValidAfter *big.Int
	ValidUntil *big.Int
}

// TestHelpersMetaData contains all meta data concerning the TestHelpers contract.
var TestHelpersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"sigFailed\",\"type\":\"bool\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"}],\"name\":\"packValidationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"}],\"internalType\":\"structValidationData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"packValidationDataStruct\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"name\":\"parseValidationData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"}],\"internalType\":\"structValidationData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557610277908161001b8239f35b600080fdfe604060808152600436101561001357600080fd5b600090813560e01c806324d3cde614610142578063a4b2282e146100ba5763b059e2fa1461004057600080fd5b346100b65760603660031901126100b6576004359081151582036100b2579160209261006a6101f5565b9161007361020f565b93156100a9575060ff6001915b519365ffffffffffff60d01b9060d01b169265ffffffffffff60a01b9060a01b16911617178152f35b60ff9091610080565b8280fd5b5080fd5b50346100b65760203660031901126100b657606090600435906100db610224565b506100e4610224565b506001600160a01b0382169165ffffffffffff9060a081901c82169082908215610139575b81856101136101bf565b888152602081019360d01c84520193168352845195865251166020850152511690820152f35b91508091610109565b50346100b65760603660031901126100b65761015c6101bf565b906004356001600160a01b038116908181036101bb57602094508352816101816101f5565b93848682015261018f61020f565b9384910152519265ffffffffffff60d01b9060d01b169165ffffffffffff60a01b9060a01b1617178152f35b8480fd5b604051906060820182811067ffffffffffffffff8211176101df57604052565b634e487b7160e01b600052604160045260246000fd5b6024359065ffffffffffff8216820361020a57565b600080fd5b6044359065ffffffffffff8216820361020a57565b61022c6101bf565b9060008252600060208301526000604083015256fea26469706673582212208738e9297a20de6f5e8f45d1f7a72a0308b7bcf1211fe6c39f996e6a9ac2abd264736f6c63430008190033",
}

// TestHelpersABI is the input ABI used to generate the binding from.
// Deprecated: Use TestHelpersMetaData.ABI instead.
var TestHelpersABI = TestHelpersMetaData.ABI

// TestHelpersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestHelpersMetaData.Bin instead.
var TestHelpersBin = TestHelpersMetaData.Bin

// DeployTestHelpers deploys a new Ethereum contract, binding an instance of TestHelpers to it.
func DeployTestHelpers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestHelpers, error) {
	parsed, err := TestHelpersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestHelpersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestHelpers{TestHelpersCaller: TestHelpersCaller{contract: contract}, TestHelpersTransactor: TestHelpersTransactor{contract: contract}, TestHelpersFilterer: TestHelpersFilterer{contract: contract}}, nil
}

// TestHelpers is an auto generated Go binding around an Ethereum contract.
type TestHelpers struct {
	TestHelpersCaller     // Read-only binding to the contract
	TestHelpersTransactor // Write-only binding to the contract
	TestHelpersFilterer   // Log filterer for contract events
}

// TestHelpersCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestHelpersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestHelpersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestHelpersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestHelpersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestHelpersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestHelpersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestHelpersSession struct {
	Contract     *TestHelpers      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestHelpersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestHelpersCallerSession struct {
	Contract *TestHelpersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestHelpersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestHelpersTransactorSession struct {
	Contract     *TestHelpersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestHelpersRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestHelpersRaw struct {
	Contract *TestHelpers // Generic contract binding to access the raw methods on
}

// TestHelpersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestHelpersCallerRaw struct {
	Contract *TestHelpersCaller // Generic read-only contract binding to access the raw methods on
}

// TestHelpersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestHelpersTransactorRaw struct {
	Contract *TestHelpersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestHelpers creates a new instance of TestHelpers, bound to a specific deployed contract.
func NewTestHelpers(address common.Address, backend bind.ContractBackend) (*TestHelpers, error) {
	contract, err := bindTestHelpers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestHelpers{TestHelpersCaller: TestHelpersCaller{contract: contract}, TestHelpersTransactor: TestHelpersTransactor{contract: contract}, TestHelpersFilterer: TestHelpersFilterer{contract: contract}}, nil
}

// NewTestHelpersCaller creates a new read-only instance of TestHelpers, bound to a specific deployed contract.
func NewTestHelpersCaller(address common.Address, caller bind.ContractCaller) (*TestHelpersCaller, error) {
	contract, err := bindTestHelpers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestHelpersCaller{contract: contract}, nil
}

// NewTestHelpersTransactor creates a new write-only instance of TestHelpers, bound to a specific deployed contract.
func NewTestHelpersTransactor(address common.Address, transactor bind.ContractTransactor) (*TestHelpersTransactor, error) {
	contract, err := bindTestHelpers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestHelpersTransactor{contract: contract}, nil
}

// NewTestHelpersFilterer creates a new log filterer instance of TestHelpers, bound to a specific deployed contract.
func NewTestHelpersFilterer(address common.Address, filterer bind.ContractFilterer) (*TestHelpersFilterer, error) {
	contract, err := bindTestHelpers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestHelpersFilterer{contract: contract}, nil
}

// bindTestHelpers binds a generic wrapper to an already deployed contract.
func bindTestHelpers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestHelpersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestHelpers *TestHelpersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestHelpers.Contract.TestHelpersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestHelpers *TestHelpersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestHelpers.Contract.TestHelpersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestHelpers *TestHelpersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestHelpers.Contract.TestHelpersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestHelpers *TestHelpersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestHelpers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestHelpers *TestHelpersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestHelpers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestHelpers *TestHelpersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestHelpers.Contract.contract.Transact(opts, method, params...)
}

// PackValidationData is a free data retrieval call binding the contract method 0xb059e2fa.
//
// Solidity: function packValidationData(bool sigFailed, uint48 validUntil, uint48 validAfter) pure returns(uint256)
func (_TestHelpers *TestHelpersCaller) PackValidationData(opts *bind.CallOpts, sigFailed bool, validUntil *big.Int, validAfter *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestHelpers.contract.Call(opts, &out, "packValidationData", sigFailed, validUntil, validAfter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PackValidationData is a free data retrieval call binding the contract method 0xb059e2fa.
//
// Solidity: function packValidationData(bool sigFailed, uint48 validUntil, uint48 validAfter) pure returns(uint256)
func (_TestHelpers *TestHelpersSession) PackValidationData(sigFailed bool, validUntil *big.Int, validAfter *big.Int) (*big.Int, error) {
	return _TestHelpers.Contract.PackValidationData(&_TestHelpers.CallOpts, sigFailed, validUntil, validAfter)
}

// PackValidationData is a free data retrieval call binding the contract method 0xb059e2fa.
//
// Solidity: function packValidationData(bool sigFailed, uint48 validUntil, uint48 validAfter) pure returns(uint256)
func (_TestHelpers *TestHelpersCallerSession) PackValidationData(sigFailed bool, validUntil *big.Int, validAfter *big.Int) (*big.Int, error) {
	return _TestHelpers.Contract.PackValidationData(&_TestHelpers.CallOpts, sigFailed, validUntil, validAfter)
}

// PackValidationDataStruct is a free data retrieval call binding the contract method 0x24d3cde6.
//
// Solidity: function packValidationDataStruct((address,uint48,uint48) data) pure returns(uint256)
func (_TestHelpers *TestHelpersCaller) PackValidationDataStruct(opts *bind.CallOpts, data ValidationData) (*big.Int, error) {
	var out []interface{}
	err := _TestHelpers.contract.Call(opts, &out, "packValidationDataStruct", data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PackValidationDataStruct is a free data retrieval call binding the contract method 0x24d3cde6.
//
// Solidity: function packValidationDataStruct((address,uint48,uint48) data) pure returns(uint256)
func (_TestHelpers *TestHelpersSession) PackValidationDataStruct(data ValidationData) (*big.Int, error) {
	return _TestHelpers.Contract.PackValidationDataStruct(&_TestHelpers.CallOpts, data)
}

// PackValidationDataStruct is a free data retrieval call binding the contract method 0x24d3cde6.
//
// Solidity: function packValidationDataStruct((address,uint48,uint48) data) pure returns(uint256)
func (_TestHelpers *TestHelpersCallerSession) PackValidationDataStruct(data ValidationData) (*big.Int, error) {
	return _TestHelpers.Contract.PackValidationDataStruct(&_TestHelpers.CallOpts, data)
}

// ParseValidationData is a free data retrieval call binding the contract method 0xa4b2282e.
//
// Solidity: function parseValidationData(uint256 validationData) pure returns((address,uint48,uint48))
func (_TestHelpers *TestHelpersCaller) ParseValidationData(opts *bind.CallOpts, validationData *big.Int) (ValidationData, error) {
	var out []interface{}
	err := _TestHelpers.contract.Call(opts, &out, "parseValidationData", validationData)

	if err != nil {
		return *new(ValidationData), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidationData)).(*ValidationData)

	return out0, err

}

// ParseValidationData is a free data retrieval call binding the contract method 0xa4b2282e.
//
// Solidity: function parseValidationData(uint256 validationData) pure returns((address,uint48,uint48))
func (_TestHelpers *TestHelpersSession) ParseValidationData(validationData *big.Int) (ValidationData, error) {
	return _TestHelpers.Contract.ParseValidationData(&_TestHelpers.CallOpts, validationData)
}

// ParseValidationData is a free data retrieval call binding the contract method 0xa4b2282e.
//
// Solidity: function parseValidationData(uint256 validationData) pure returns((address,uint48,uint48))
func (_TestHelpers *TestHelpersCallerSession) ParseValidationData(validationData *big.Int) (ValidationData, error) {
	return _TestHelpers.Contract.ParseValidationData(&_TestHelpers.CallOpts, validationData)
}
