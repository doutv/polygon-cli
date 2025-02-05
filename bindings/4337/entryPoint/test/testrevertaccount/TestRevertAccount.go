// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testrevertaccount

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

// TestRevertAccountMetaData contains all meta data concerning the TestRevertAccount contract.
var TestRevertAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_ep\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"revertLong\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080601f6101a738819003918201601f19168301916001600160401b03831184841017607157808492602094604052833981010312606c57516001600160a01b03811690819003606c57600080546001600160a01b03191691909117905560405161011f90816100888239f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe60806040526004361015601157600080fd5b6000803560e01c90816319822f7c146036575063be76c6ef14603257600080fd5b60ce565b3460c35760031960603682011260ca576004359067ffffffffffffffff821160c657610120913603011260c357805481906001600160a01b0316803b1560c05760246080809263b760faf960e01b825230608452604435905af190811560b4575060a7575b60405160008152602090f35b60b06080604052565b609b565b604051903d90823e3d90fd5b50fd5b80fd5b8280fd5b5080fd5b3460e457602036600319011260e4576004356000fd5b600080fdfea2646970667358221220e351f8ab0fd6db992776771a1ca14ea9eb54f8e005f8b533080e92c1b40879e964736f6c63430008190033",
}

// TestRevertAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use TestRevertAccountMetaData.ABI instead.
var TestRevertAccountABI = TestRevertAccountMetaData.ABI

// TestRevertAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestRevertAccountMetaData.Bin instead.
var TestRevertAccountBin = TestRevertAccountMetaData.Bin

// DeployTestRevertAccount deploys a new Ethereum contract, binding an instance of TestRevertAccount to it.
func DeployTestRevertAccount(auth *bind.TransactOpts, backend bind.ContractBackend, _ep common.Address) (common.Address, *types.Transaction, *TestRevertAccount, error) {
	parsed, err := TestRevertAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestRevertAccountBin), backend, _ep)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestRevertAccount{TestRevertAccountCaller: TestRevertAccountCaller{contract: contract}, TestRevertAccountTransactor: TestRevertAccountTransactor{contract: contract}, TestRevertAccountFilterer: TestRevertAccountFilterer{contract: contract}}, nil
}

// TestRevertAccount is an auto generated Go binding around an Ethereum contract.
type TestRevertAccount struct {
	TestRevertAccountCaller     // Read-only binding to the contract
	TestRevertAccountTransactor // Write-only binding to the contract
	TestRevertAccountFilterer   // Log filterer for contract events
}

// TestRevertAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestRevertAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestRevertAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestRevertAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestRevertAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestRevertAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestRevertAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestRevertAccountSession struct {
	Contract     *TestRevertAccount // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestRevertAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestRevertAccountCallerSession struct {
	Contract *TestRevertAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TestRevertAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestRevertAccountTransactorSession struct {
	Contract     *TestRevertAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TestRevertAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRevertAccountRaw struct {
	Contract *TestRevertAccount // Generic contract binding to access the raw methods on
}

// TestRevertAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestRevertAccountCallerRaw struct {
	Contract *TestRevertAccountCaller // Generic read-only contract binding to access the raw methods on
}

// TestRevertAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestRevertAccountTransactorRaw struct {
	Contract *TestRevertAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestRevertAccount creates a new instance of TestRevertAccount, bound to a specific deployed contract.
func NewTestRevertAccount(address common.Address, backend bind.ContractBackend) (*TestRevertAccount, error) {
	contract, err := bindTestRevertAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestRevertAccount{TestRevertAccountCaller: TestRevertAccountCaller{contract: contract}, TestRevertAccountTransactor: TestRevertAccountTransactor{contract: contract}, TestRevertAccountFilterer: TestRevertAccountFilterer{contract: contract}}, nil
}

// NewTestRevertAccountCaller creates a new read-only instance of TestRevertAccount, bound to a specific deployed contract.
func NewTestRevertAccountCaller(address common.Address, caller bind.ContractCaller) (*TestRevertAccountCaller, error) {
	contract, err := bindTestRevertAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestRevertAccountCaller{contract: contract}, nil
}

// NewTestRevertAccountTransactor creates a new write-only instance of TestRevertAccount, bound to a specific deployed contract.
func NewTestRevertAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*TestRevertAccountTransactor, error) {
	contract, err := bindTestRevertAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestRevertAccountTransactor{contract: contract}, nil
}

// NewTestRevertAccountFilterer creates a new log filterer instance of TestRevertAccount, bound to a specific deployed contract.
func NewTestRevertAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*TestRevertAccountFilterer, error) {
	contract, err := bindTestRevertAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestRevertAccountFilterer{contract: contract}, nil
}

// bindTestRevertAccount binds a generic wrapper to an already deployed contract.
func bindTestRevertAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestRevertAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestRevertAccount *TestRevertAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestRevertAccount.Contract.TestRevertAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestRevertAccount *TestRevertAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestRevertAccount.Contract.TestRevertAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestRevertAccount *TestRevertAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestRevertAccount.Contract.TestRevertAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestRevertAccount *TestRevertAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestRevertAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestRevertAccount *TestRevertAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestRevertAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestRevertAccount *TestRevertAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestRevertAccount.Contract.contract.Transact(opts, method, params...)
}

// RevertLong is a free data retrieval call binding the contract method 0xbe76c6ef.
//
// Solidity: function revertLong(uint256 length) pure returns()
func (_TestRevertAccount *TestRevertAccountCaller) RevertLong(opts *bind.CallOpts, length *big.Int) error {
	var out []interface{}
	err := _TestRevertAccount.contract.Call(opts, &out, "revertLong", length)

	if err != nil {
		return err
	}

	return err

}

// RevertLong is a free data retrieval call binding the contract method 0xbe76c6ef.
//
// Solidity: function revertLong(uint256 length) pure returns()
func (_TestRevertAccount *TestRevertAccountSession) RevertLong(length *big.Int) error {
	return _TestRevertAccount.Contract.RevertLong(&_TestRevertAccount.CallOpts, length)
}

// RevertLong is a free data retrieval call binding the contract method 0xbe76c6ef.
//
// Solidity: function revertLong(uint256 length) pure returns()
func (_TestRevertAccount *TestRevertAccountCallerSession) RevertLong(length *big.Int) error {
	return _TestRevertAccount.Contract.RevertLong(&_TestRevertAccount.CallOpts, length)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) , bytes32 , uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestRevertAccount *TestRevertAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, arg0 PackedUserOperation, arg1 [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestRevertAccount.contract.Transact(opts, "validateUserOp", arg0, arg1, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) , bytes32 , uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestRevertAccount *TestRevertAccountSession) ValidateUserOp(arg0 PackedUserOperation, arg1 [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestRevertAccount.Contract.ValidateUserOp(&_TestRevertAccount.TransactOpts, arg0, arg1, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) , bytes32 , uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestRevertAccount *TestRevertAccountTransactorSession) ValidateUserOp(arg0 PackedUserOperation, arg1 [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestRevertAccount.Contract.ValidateUserOp(&_TestRevertAccount.TransactOpts, arg0, arg1, missingAccountFunds)
}
