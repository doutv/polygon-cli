// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package baseaccount

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

// BaseAccountMetaData contains all meta data concerning the BaseAccount contract.
var BaseAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BaseAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseAccountMetaData.ABI instead.
var BaseAccountABI = BaseAccountMetaData.ABI

// BaseAccount is an auto generated Go binding around an Ethereum contract.
type BaseAccount struct {
	BaseAccountCaller     // Read-only binding to the contract
	BaseAccountTransactor // Write-only binding to the contract
	BaseAccountFilterer   // Log filterer for contract events
}

// BaseAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseAccountSession struct {
	Contract     *BaseAccount      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseAccountCallerSession struct {
	Contract *BaseAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BaseAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseAccountTransactorSession struct {
	Contract     *BaseAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BaseAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseAccountRaw struct {
	Contract *BaseAccount // Generic contract binding to access the raw methods on
}

// BaseAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseAccountCallerRaw struct {
	Contract *BaseAccountCaller // Generic read-only contract binding to access the raw methods on
}

// BaseAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseAccountTransactorRaw struct {
	Contract *BaseAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseAccount creates a new instance of BaseAccount, bound to a specific deployed contract.
func NewBaseAccount(address common.Address, backend bind.ContractBackend) (*BaseAccount, error) {
	contract, err := bindBaseAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseAccount{BaseAccountCaller: BaseAccountCaller{contract: contract}, BaseAccountTransactor: BaseAccountTransactor{contract: contract}, BaseAccountFilterer: BaseAccountFilterer{contract: contract}}, nil
}

// NewBaseAccountCaller creates a new read-only instance of BaseAccount, bound to a specific deployed contract.
func NewBaseAccountCaller(address common.Address, caller bind.ContractCaller) (*BaseAccountCaller, error) {
	contract, err := bindBaseAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseAccountCaller{contract: contract}, nil
}

// NewBaseAccountTransactor creates a new write-only instance of BaseAccount, bound to a specific deployed contract.
func NewBaseAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseAccountTransactor, error) {
	contract, err := bindBaseAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseAccountTransactor{contract: contract}, nil
}

// NewBaseAccountFilterer creates a new log filterer instance of BaseAccount, bound to a specific deployed contract.
func NewBaseAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseAccountFilterer, error) {
	contract, err := bindBaseAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseAccountFilterer{contract: contract}, nil
}

// bindBaseAccount binds a generic wrapper to an already deployed contract.
func bindBaseAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseAccount *BaseAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseAccount.Contract.BaseAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseAccount *BaseAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseAccount.Contract.BaseAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseAccount *BaseAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseAccount.Contract.BaseAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseAccount *BaseAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseAccount *BaseAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseAccount *BaseAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseAccount.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_BaseAccount *BaseAccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseAccount.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_BaseAccount *BaseAccountSession) EntryPoint() (common.Address, error) {
	return _BaseAccount.Contract.EntryPoint(&_BaseAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_BaseAccount *BaseAccountCallerSession) EntryPoint() (common.Address, error) {
	return _BaseAccount.Contract.EntryPoint(&_BaseAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_BaseAccount *BaseAccountCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseAccount.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_BaseAccount *BaseAccountSession) GetNonce() (*big.Int, error) {
	return _BaseAccount.Contract.GetNonce(&_BaseAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_BaseAccount *BaseAccountCallerSession) GetNonce() (*big.Int, error) {
	return _BaseAccount.Contract.GetNonce(&_BaseAccount.CallOpts)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_BaseAccount *BaseAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _BaseAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_BaseAccount *BaseAccountSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _BaseAccount.Contract.ValidateUserOp(&_BaseAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_BaseAccount *BaseAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _BaseAccount.Contract.ValidateUserOp(&_BaseAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}
