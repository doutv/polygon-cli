// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testutil

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

// TestUtilMetaData contains all meta data concerning the TestUtil contract.
var TestUtilMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"op\",\"type\":\"tuple\"}],\"name\":\"encodeUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608080604052346015576101d8908161001b8239f35b600080fdfe60806040818152600436101561001457600080fd5b600091823560e01c63a124062e1461002b57600080fd5b34610163576020600319936020853601126101675767ffffffffffffffff946004358681116101635780600401946101208093833603011261015f5760c490610077604484018861016a565b8291819337209161008b606482018861016a565b90818a51918237206100a060e483018961016a565b90818b5191823720895198356001600160a01b031660208a01526024830135898b015260608901949094526080880152608481013560a088015260a481013560c0880152013560e086015261010080860191909152845283019485118386101761014b578484526020855282519182610140850152815b8381106101365750508282016101600152601f01601f19168101030190f35b80829186016101608382015191015201610117565b634e487b7160e01b81526041600452602490fd5b8380fd5b8280fd5b80fd5b903590601e198136030182121561019d570180359067ffffffffffffffff821161019d5760200191813603831361019d57565b600080fdfea2646970667358221220b3a56026a9ab41a5c124e612c62c194af0883ae5c9802820d953ce94812b7e3f64736f6c63430008190033",
}

// TestUtilABI is the input ABI used to generate the binding from.
// Deprecated: Use TestUtilMetaData.ABI instead.
var TestUtilABI = TestUtilMetaData.ABI

// TestUtilBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestUtilMetaData.Bin instead.
var TestUtilBin = TestUtilMetaData.Bin

// DeployTestUtil deploys a new Ethereum contract, binding an instance of TestUtil to it.
func DeployTestUtil(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestUtil, error) {
	parsed, err := TestUtilMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestUtilBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestUtil{TestUtilCaller: TestUtilCaller{contract: contract}, TestUtilTransactor: TestUtilTransactor{contract: contract}, TestUtilFilterer: TestUtilFilterer{contract: contract}}, nil
}

// TestUtil is an auto generated Go binding around an Ethereum contract.
type TestUtil struct {
	TestUtilCaller     // Read-only binding to the contract
	TestUtilTransactor // Write-only binding to the contract
	TestUtilFilterer   // Log filterer for contract events
}

// TestUtilCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestUtilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUtilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestUtilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUtilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestUtilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUtilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestUtilSession struct {
	Contract     *TestUtil         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestUtilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestUtilCallerSession struct {
	Contract *TestUtilCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TestUtilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestUtilTransactorSession struct {
	Contract     *TestUtilTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TestUtilRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestUtilRaw struct {
	Contract *TestUtil // Generic contract binding to access the raw methods on
}

// TestUtilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestUtilCallerRaw struct {
	Contract *TestUtilCaller // Generic read-only contract binding to access the raw methods on
}

// TestUtilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestUtilTransactorRaw struct {
	Contract *TestUtilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestUtil creates a new instance of TestUtil, bound to a specific deployed contract.
func NewTestUtil(address common.Address, backend bind.ContractBackend) (*TestUtil, error) {
	contract, err := bindTestUtil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestUtil{TestUtilCaller: TestUtilCaller{contract: contract}, TestUtilTransactor: TestUtilTransactor{contract: contract}, TestUtilFilterer: TestUtilFilterer{contract: contract}}, nil
}

// NewTestUtilCaller creates a new read-only instance of TestUtil, bound to a specific deployed contract.
func NewTestUtilCaller(address common.Address, caller bind.ContractCaller) (*TestUtilCaller, error) {
	contract, err := bindTestUtil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestUtilCaller{contract: contract}, nil
}

// NewTestUtilTransactor creates a new write-only instance of TestUtil, bound to a specific deployed contract.
func NewTestUtilTransactor(address common.Address, transactor bind.ContractTransactor) (*TestUtilTransactor, error) {
	contract, err := bindTestUtil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestUtilTransactor{contract: contract}, nil
}

// NewTestUtilFilterer creates a new log filterer instance of TestUtil, bound to a specific deployed contract.
func NewTestUtilFilterer(address common.Address, filterer bind.ContractFilterer) (*TestUtilFilterer, error) {
	contract, err := bindTestUtil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestUtilFilterer{contract: contract}, nil
}

// bindTestUtil binds a generic wrapper to an already deployed contract.
func bindTestUtil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestUtilMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestUtil *TestUtilRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestUtil.Contract.TestUtilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestUtil *TestUtilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUtil.Contract.TestUtilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestUtil *TestUtilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestUtil.Contract.TestUtilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestUtil *TestUtilCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestUtil.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestUtil *TestUtilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUtil.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestUtil *TestUtilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestUtil.Contract.contract.Transact(opts, method, params...)
}

// EncodeUserOp is a free data retrieval call binding the contract method 0xa124062e.
//
// Solidity: function encodeUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) op) pure returns(bytes)
func (_TestUtil *TestUtilCaller) EncodeUserOp(opts *bind.CallOpts, op PackedUserOperation) ([]byte, error) {
	var out []interface{}
	err := _TestUtil.contract.Call(opts, &out, "encodeUserOp", op)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeUserOp is a free data retrieval call binding the contract method 0xa124062e.
//
// Solidity: function encodeUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) op) pure returns(bytes)
func (_TestUtil *TestUtilSession) EncodeUserOp(op PackedUserOperation) ([]byte, error) {
	return _TestUtil.Contract.EncodeUserOp(&_TestUtil.CallOpts, op)
}

// EncodeUserOp is a free data retrieval call binding the contract method 0xa124062e.
//
// Solidity: function encodeUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) op) pure returns(bytes)
func (_TestUtil *TestUtilCallerSession) EncodeUserOp(op PackedUserOperation) ([]byte, error) {
	return _TestUtil.Contract.EncodeUserOp(&_TestUtil.CallOpts, op)
}
