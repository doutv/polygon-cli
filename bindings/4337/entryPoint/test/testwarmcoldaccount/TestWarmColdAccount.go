// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testwarmcoldaccount

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

// TestWarmColdAccountMetaData contains all meta data concerning the TestWarmColdAccount contract.
var TestWarmColdAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_ep\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"touchPaymaster\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"touchStorage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080601f61038b38819003918201601f19168301916001600160401b03831184841017607557808492602094604052833981010312607057516001600160a01b0381169081900360705760018055600080546001600160a01b0319169190911790556040516102ff908161008c8239f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe608060408181526004918236101561001657600080fd5b600092833560e01c91826319822f7c146100c557508163c19d93fb146100a6578163f115d40d14610072575063fb0c24251461005157600080fd5b3461006e578160031936011261006e576020906001549051908152f35b5080fd5b9050346100a25760203660031901126100a257356001600160a01b03811681036100a257602092503b9051908152f35b8280fd5b50503461006e578160031936011261006e576020906001549051908152f35b90849250346100a25760031960603682011261025157813567ffffffffffffffff9384821161028d576101208236039384011261028d5785546001600160a01b0316803b156102895781602481899363b760faf960e01b82523089830152604435905af1801561027f57610259575b506024810135600181036101ae575050845163fb0c242560e01b8152925060209183915081306103e8fa80156101a457610174575b506020915b51908152f35b602090813d831161019d575b61018a8183610291565b810103126101985782610169565b600080fd5b503d610180565b83513d84823e3d90fd5b6002146101c2575b5050505060209161016e565b60e4810135916022190182121561025557018181013592831161025157602401918036038313610251576014116100a257835163f115d40d60e01b8152913560601c90820152602081602481306103e8fa80156101a457610226575b8080806101b6565b602090813d831161024a575b61023c8183610291565b81010312610198578261021e565b503d610232565b8380fd5b8480fd5b84819692961161026c5786529386610134565b634e487b7160e01b825260418452602482fd5b87513d88823e3d90fd5b8680fd5b8580fd5b90601f8019910116810190811067ffffffffffffffff8211176102b357604052565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220d8443a148157c2f10daec85f86224dccb36331c4fec7252196024d25f024fdbe64736f6c63430008190033",
}

// TestWarmColdAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use TestWarmColdAccountMetaData.ABI instead.
var TestWarmColdAccountABI = TestWarmColdAccountMetaData.ABI

// TestWarmColdAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestWarmColdAccountMetaData.Bin instead.
var TestWarmColdAccountBin = TestWarmColdAccountMetaData.Bin

// DeployTestWarmColdAccount deploys a new Ethereum contract, binding an instance of TestWarmColdAccount to it.
func DeployTestWarmColdAccount(auth *bind.TransactOpts, backend bind.ContractBackend, _ep common.Address) (common.Address, *types.Transaction, *TestWarmColdAccount, error) {
	parsed, err := TestWarmColdAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestWarmColdAccountBin), backend, _ep)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestWarmColdAccount{TestWarmColdAccountCaller: TestWarmColdAccountCaller{contract: contract}, TestWarmColdAccountTransactor: TestWarmColdAccountTransactor{contract: contract}, TestWarmColdAccountFilterer: TestWarmColdAccountFilterer{contract: contract}}, nil
}

// TestWarmColdAccount is an auto generated Go binding around an Ethereum contract.
type TestWarmColdAccount struct {
	TestWarmColdAccountCaller     // Read-only binding to the contract
	TestWarmColdAccountTransactor // Write-only binding to the contract
	TestWarmColdAccountFilterer   // Log filterer for contract events
}

// TestWarmColdAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestWarmColdAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestWarmColdAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestWarmColdAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestWarmColdAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestWarmColdAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestWarmColdAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestWarmColdAccountSession struct {
	Contract     *TestWarmColdAccount // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestWarmColdAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestWarmColdAccountCallerSession struct {
	Contract *TestWarmColdAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TestWarmColdAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestWarmColdAccountTransactorSession struct {
	Contract     *TestWarmColdAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TestWarmColdAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestWarmColdAccountRaw struct {
	Contract *TestWarmColdAccount // Generic contract binding to access the raw methods on
}

// TestWarmColdAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestWarmColdAccountCallerRaw struct {
	Contract *TestWarmColdAccountCaller // Generic read-only contract binding to access the raw methods on
}

// TestWarmColdAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestWarmColdAccountTransactorRaw struct {
	Contract *TestWarmColdAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestWarmColdAccount creates a new instance of TestWarmColdAccount, bound to a specific deployed contract.
func NewTestWarmColdAccount(address common.Address, backend bind.ContractBackend) (*TestWarmColdAccount, error) {
	contract, err := bindTestWarmColdAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestWarmColdAccount{TestWarmColdAccountCaller: TestWarmColdAccountCaller{contract: contract}, TestWarmColdAccountTransactor: TestWarmColdAccountTransactor{contract: contract}, TestWarmColdAccountFilterer: TestWarmColdAccountFilterer{contract: contract}}, nil
}

// NewTestWarmColdAccountCaller creates a new read-only instance of TestWarmColdAccount, bound to a specific deployed contract.
func NewTestWarmColdAccountCaller(address common.Address, caller bind.ContractCaller) (*TestWarmColdAccountCaller, error) {
	contract, err := bindTestWarmColdAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestWarmColdAccountCaller{contract: contract}, nil
}

// NewTestWarmColdAccountTransactor creates a new write-only instance of TestWarmColdAccount, bound to a specific deployed contract.
func NewTestWarmColdAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*TestWarmColdAccountTransactor, error) {
	contract, err := bindTestWarmColdAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestWarmColdAccountTransactor{contract: contract}, nil
}

// NewTestWarmColdAccountFilterer creates a new log filterer instance of TestWarmColdAccount, bound to a specific deployed contract.
func NewTestWarmColdAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*TestWarmColdAccountFilterer, error) {
	contract, err := bindTestWarmColdAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestWarmColdAccountFilterer{contract: contract}, nil
}

// bindTestWarmColdAccount binds a generic wrapper to an already deployed contract.
func bindTestWarmColdAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestWarmColdAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestWarmColdAccount *TestWarmColdAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestWarmColdAccount.Contract.TestWarmColdAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestWarmColdAccount *TestWarmColdAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestWarmColdAccount.Contract.TestWarmColdAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestWarmColdAccount *TestWarmColdAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestWarmColdAccount.Contract.TestWarmColdAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestWarmColdAccount *TestWarmColdAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestWarmColdAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestWarmColdAccount *TestWarmColdAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestWarmColdAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestWarmColdAccount *TestWarmColdAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestWarmColdAccount.Contract.contract.Transact(opts, method, params...)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256)
func (_TestWarmColdAccount *TestWarmColdAccountCaller) State(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestWarmColdAccount.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256)
func (_TestWarmColdAccount *TestWarmColdAccountSession) State() (*big.Int, error) {
	return _TestWarmColdAccount.Contract.State(&_TestWarmColdAccount.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256)
func (_TestWarmColdAccount *TestWarmColdAccountCallerSession) State() (*big.Int, error) {
	return _TestWarmColdAccount.Contract.State(&_TestWarmColdAccount.CallOpts)
}

// TouchPaymaster is a free data retrieval call binding the contract method 0xf115d40d.
//
// Solidity: function touchPaymaster(address paymaster) view returns(uint256)
func (_TestWarmColdAccount *TestWarmColdAccountCaller) TouchPaymaster(opts *bind.CallOpts, paymaster common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestWarmColdAccount.contract.Call(opts, &out, "touchPaymaster", paymaster)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TouchPaymaster is a free data retrieval call binding the contract method 0xf115d40d.
//
// Solidity: function touchPaymaster(address paymaster) view returns(uint256)
func (_TestWarmColdAccount *TestWarmColdAccountSession) TouchPaymaster(paymaster common.Address) (*big.Int, error) {
	return _TestWarmColdAccount.Contract.TouchPaymaster(&_TestWarmColdAccount.CallOpts, paymaster)
}

// TouchPaymaster is a free data retrieval call binding the contract method 0xf115d40d.
//
// Solidity: function touchPaymaster(address paymaster) view returns(uint256)
func (_TestWarmColdAccount *TestWarmColdAccountCallerSession) TouchPaymaster(paymaster common.Address) (*big.Int, error) {
	return _TestWarmColdAccount.Contract.TouchPaymaster(&_TestWarmColdAccount.CallOpts, paymaster)
}

// TouchStorage is a free data retrieval call binding the contract method 0xfb0c2425.
//
// Solidity: function touchStorage() view returns(uint256)
func (_TestWarmColdAccount *TestWarmColdAccountCaller) TouchStorage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestWarmColdAccount.contract.Call(opts, &out, "touchStorage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TouchStorage is a free data retrieval call binding the contract method 0xfb0c2425.
//
// Solidity: function touchStorage() view returns(uint256)
func (_TestWarmColdAccount *TestWarmColdAccountSession) TouchStorage() (*big.Int, error) {
	return _TestWarmColdAccount.Contract.TouchStorage(&_TestWarmColdAccount.CallOpts)
}

// TouchStorage is a free data retrieval call binding the contract method 0xfb0c2425.
//
// Solidity: function touchStorage() view returns(uint256)
func (_TestWarmColdAccount *TestWarmColdAccountCallerSession) TouchStorage() (*big.Int, error) {
	return _TestWarmColdAccount.Contract.TouchStorage(&_TestWarmColdAccount.CallOpts)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 , uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestWarmColdAccount *TestWarmColdAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, arg1 [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestWarmColdAccount.contract.Transact(opts, "validateUserOp", userOp, arg1, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 , uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestWarmColdAccount *TestWarmColdAccountSession) ValidateUserOp(userOp PackedUserOperation, arg1 [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestWarmColdAccount.Contract.ValidateUserOp(&_TestWarmColdAccount.TransactOpts, userOp, arg1, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 , uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestWarmColdAccount *TestWarmColdAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, arg1 [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestWarmColdAccount.Contract.ValidateUserOp(&_TestWarmColdAccount.TransactOpts, userOp, arg1, missingAccountFunds)
}
