// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipayableaccount

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

// IPayableAccountMetaData contains all meta data concerning the IPayableAccount contract.
var IPayableAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"PassKeyError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"oldValidFrom\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"newValidFrom\",\"type\":\"uint128\"}],\"name\":\"RecoveryExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RecoveryFeePaymentFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"subject\",\"type\":\"bytes\"}],\"name\":\"AccountRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ModuleInstalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"typeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ModuleInstalledWithData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ModuleUninstalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"typeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ModuleUninstalledWithData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoveryFeeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recoveryModule\",\"type\":\"address\"}],\"name\":\"RecoveryModuleInstalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recoveryModule\",\"type\":\"address\"}],\"name\":\"RecoveryModuleUninstalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeReceived\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accountId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"accountImplementationId\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"claimRecoveryFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"executionCalldata\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"executionCalldata\",\"type\":\"bytes\"}],\"name\":\"executeFromExecutor\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"typeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"installModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recoveryModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"installRecoveryModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"additionalContext\",\"type\":\"bytes\"}],\"name\":\"isModuleInstalled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"}],\"name\":\"supportsModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"typeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uninstallModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recoveryModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uninstallRecoveryModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPayableAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use IPayableAccountMetaData.ABI instead.
var IPayableAccountABI = IPayableAccountMetaData.ABI

// IPayableAccount is an auto generated Go binding around an Ethereum contract.
type IPayableAccount struct {
	IPayableAccountCaller     // Read-only binding to the contract
	IPayableAccountTransactor // Write-only binding to the contract
	IPayableAccountFilterer   // Log filterer for contract events
}

// IPayableAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPayableAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPayableAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPayableAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPayableAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPayableAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPayableAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPayableAccountSession struct {
	Contract     *IPayableAccount  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPayableAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPayableAccountCallerSession struct {
	Contract *IPayableAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IPayableAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPayableAccountTransactorSession struct {
	Contract     *IPayableAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IPayableAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPayableAccountRaw struct {
	Contract *IPayableAccount // Generic contract binding to access the raw methods on
}

// IPayableAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPayableAccountCallerRaw struct {
	Contract *IPayableAccountCaller // Generic read-only contract binding to access the raw methods on
}

// IPayableAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPayableAccountTransactorRaw struct {
	Contract *IPayableAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPayableAccount creates a new instance of IPayableAccount, bound to a specific deployed contract.
func NewIPayableAccount(address common.Address, backend bind.ContractBackend) (*IPayableAccount, error) {
	contract, err := bindIPayableAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPayableAccount{IPayableAccountCaller: IPayableAccountCaller{contract: contract}, IPayableAccountTransactor: IPayableAccountTransactor{contract: contract}, IPayableAccountFilterer: IPayableAccountFilterer{contract: contract}}, nil
}

// NewIPayableAccountCaller creates a new read-only instance of IPayableAccount, bound to a specific deployed contract.
func NewIPayableAccountCaller(address common.Address, caller bind.ContractCaller) (*IPayableAccountCaller, error) {
	contract, err := bindIPayableAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPayableAccountCaller{contract: contract}, nil
}

// NewIPayableAccountTransactor creates a new write-only instance of IPayableAccount, bound to a specific deployed contract.
func NewIPayableAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*IPayableAccountTransactor, error) {
	contract, err := bindIPayableAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPayableAccountTransactor{contract: contract}, nil
}

// NewIPayableAccountFilterer creates a new log filterer instance of IPayableAccount, bound to a specific deployed contract.
func NewIPayableAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*IPayableAccountFilterer, error) {
	contract, err := bindIPayableAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPayableAccountFilterer{contract: contract}, nil
}

// bindIPayableAccount binds a generic wrapper to an already deployed contract.
func bindIPayableAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPayableAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPayableAccount *IPayableAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPayableAccount.Contract.IPayableAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPayableAccount *IPayableAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPayableAccount.Contract.IPayableAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPayableAccount *IPayableAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPayableAccount.Contract.IPayableAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPayableAccount *IPayableAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPayableAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPayableAccount *IPayableAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPayableAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPayableAccount *IPayableAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPayableAccount.Contract.contract.Transact(opts, method, params...)
}

// AccountId is a free data retrieval call binding the contract method 0x9cfd7cff.
//
// Solidity: function accountId() view returns(string accountImplementationId)
func (_IPayableAccount *IPayableAccountCaller) AccountId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IPayableAccount.contract.Call(opts, &out, "accountId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AccountId is a free data retrieval call binding the contract method 0x9cfd7cff.
//
// Solidity: function accountId() view returns(string accountImplementationId)
func (_IPayableAccount *IPayableAccountSession) AccountId() (string, error) {
	return _IPayableAccount.Contract.AccountId(&_IPayableAccount.CallOpts)
}

// AccountId is a free data retrieval call binding the contract method 0x9cfd7cff.
//
// Solidity: function accountId() view returns(string accountImplementationId)
func (_IPayableAccount *IPayableAccountCallerSession) AccountId() (string, error) {
	return _IPayableAccount.Contract.AccountId(&_IPayableAccount.CallOpts)
}

// IsModuleInstalled is a free data retrieval call binding the contract method 0x112d3a7d.
//
// Solidity: function isModuleInstalled(uint256 moduleTypeId, address module, bytes additionalContext) view returns(bool)
func (_IPayableAccount *IPayableAccountCaller) IsModuleInstalled(opts *bind.CallOpts, moduleTypeId *big.Int, module common.Address, additionalContext []byte) (bool, error) {
	var out []interface{}
	err := _IPayableAccount.contract.Call(opts, &out, "isModuleInstalled", moduleTypeId, module, additionalContext)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleInstalled is a free data retrieval call binding the contract method 0x112d3a7d.
//
// Solidity: function isModuleInstalled(uint256 moduleTypeId, address module, bytes additionalContext) view returns(bool)
func (_IPayableAccount *IPayableAccountSession) IsModuleInstalled(moduleTypeId *big.Int, module common.Address, additionalContext []byte) (bool, error) {
	return _IPayableAccount.Contract.IsModuleInstalled(&_IPayableAccount.CallOpts, moduleTypeId, module, additionalContext)
}

// IsModuleInstalled is a free data retrieval call binding the contract method 0x112d3a7d.
//
// Solidity: function isModuleInstalled(uint256 moduleTypeId, address module, bytes additionalContext) view returns(bool)
func (_IPayableAccount *IPayableAccountCallerSession) IsModuleInstalled(moduleTypeId *big.Int, module common.Address, additionalContext []byte) (bool, error) {
	return _IPayableAccount.Contract.IsModuleInstalled(&_IPayableAccount.CallOpts, moduleTypeId, module, additionalContext)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_IPayableAccount *IPayableAccountCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _IPayableAccount.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_IPayableAccount *IPayableAccountSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _IPayableAccount.Contract.IsValidSignature(&_IPayableAccount.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_IPayableAccount *IPayableAccountCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _IPayableAccount.Contract.IsValidSignature(&_IPayableAccount.CallOpts, _hash, _signature)
}

// SupportsModule is a free data retrieval call binding the contract method 0xf2dc691d.
//
// Solidity: function supportsModule(uint256 moduleTypeId) view returns(bool)
func (_IPayableAccount *IPayableAccountCaller) SupportsModule(opts *bind.CallOpts, moduleTypeId *big.Int) (bool, error) {
	var out []interface{}
	err := _IPayableAccount.contract.Call(opts, &out, "supportsModule", moduleTypeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsModule is a free data retrieval call binding the contract method 0xf2dc691d.
//
// Solidity: function supportsModule(uint256 moduleTypeId) view returns(bool)
func (_IPayableAccount *IPayableAccountSession) SupportsModule(moduleTypeId *big.Int) (bool, error) {
	return _IPayableAccount.Contract.SupportsModule(&_IPayableAccount.CallOpts, moduleTypeId)
}

// SupportsModule is a free data retrieval call binding the contract method 0xf2dc691d.
//
// Solidity: function supportsModule(uint256 moduleTypeId) view returns(bool)
func (_IPayableAccount *IPayableAccountCallerSession) SupportsModule(moduleTypeId *big.Int) (bool, error) {
	return _IPayableAccount.Contract.SupportsModule(&_IPayableAccount.CallOpts, moduleTypeId)
}

// ClaimRecoveryFee is a paid mutator transaction binding the contract method 0x3eb1765a.
//
// Solidity: function claimRecoveryFee(address receiver, uint256 value) returns()
func (_IPayableAccount *IPayableAccountTransactor) ClaimRecoveryFee(opts *bind.TransactOpts, receiver common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPayableAccount.contract.Transact(opts, "claimRecoveryFee", receiver, value)
}

// ClaimRecoveryFee is a paid mutator transaction binding the contract method 0x3eb1765a.
//
// Solidity: function claimRecoveryFee(address receiver, uint256 value) returns()
func (_IPayableAccount *IPayableAccountSession) ClaimRecoveryFee(receiver common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPayableAccount.Contract.ClaimRecoveryFee(&_IPayableAccount.TransactOpts, receiver, value)
}

// ClaimRecoveryFee is a paid mutator transaction binding the contract method 0x3eb1765a.
//
// Solidity: function claimRecoveryFee(address receiver, uint256 value) returns()
func (_IPayableAccount *IPayableAccountTransactorSession) ClaimRecoveryFee(receiver common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPayableAccount.Contract.ClaimRecoveryFee(&_IPayableAccount.TransactOpts, receiver, value)
}

// Execute is a paid mutator transaction binding the contract method 0xe9ae5c53.
//
// Solidity: function execute(bytes32 mode, bytes executionCalldata) returns()
func (_IPayableAccount *IPayableAccountTransactor) Execute(opts *bind.TransactOpts, mode [32]byte, executionCalldata []byte) (*types.Transaction, error) {
	return _IPayableAccount.contract.Transact(opts, "execute", mode, executionCalldata)
}

// Execute is a paid mutator transaction binding the contract method 0xe9ae5c53.
//
// Solidity: function execute(bytes32 mode, bytes executionCalldata) returns()
func (_IPayableAccount *IPayableAccountSession) Execute(mode [32]byte, executionCalldata []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.Execute(&_IPayableAccount.TransactOpts, mode, executionCalldata)
}

// Execute is a paid mutator transaction binding the contract method 0xe9ae5c53.
//
// Solidity: function execute(bytes32 mode, bytes executionCalldata) returns()
func (_IPayableAccount *IPayableAccountTransactorSession) Execute(mode [32]byte, executionCalldata []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.Execute(&_IPayableAccount.TransactOpts, mode, executionCalldata)
}

// ExecuteFromExecutor is a paid mutator transaction binding the contract method 0xd691c964.
//
// Solidity: function executeFromExecutor(bytes32 mode, bytes executionCalldata) returns(bytes[] returnData)
func (_IPayableAccount *IPayableAccountTransactor) ExecuteFromExecutor(opts *bind.TransactOpts, mode [32]byte, executionCalldata []byte) (*types.Transaction, error) {
	return _IPayableAccount.contract.Transact(opts, "executeFromExecutor", mode, executionCalldata)
}

// ExecuteFromExecutor is a paid mutator transaction binding the contract method 0xd691c964.
//
// Solidity: function executeFromExecutor(bytes32 mode, bytes executionCalldata) returns(bytes[] returnData)
func (_IPayableAccount *IPayableAccountSession) ExecuteFromExecutor(mode [32]byte, executionCalldata []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.ExecuteFromExecutor(&_IPayableAccount.TransactOpts, mode, executionCalldata)
}

// ExecuteFromExecutor is a paid mutator transaction binding the contract method 0xd691c964.
//
// Solidity: function executeFromExecutor(bytes32 mode, bytes executionCalldata) returns(bytes[] returnData)
func (_IPayableAccount *IPayableAccountTransactorSession) ExecuteFromExecutor(mode [32]byte, executionCalldata []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.ExecuteFromExecutor(&_IPayableAccount.TransactOpts, mode, executionCalldata)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_IPayableAccount *IPayableAccountTransactor) Initialize(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.contract.Transact(opts, "initialize", data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_IPayableAccount *IPayableAccountSession) Initialize(data []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.Initialize(&_IPayableAccount.TransactOpts, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_IPayableAccount *IPayableAccountTransactorSession) Initialize(data []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.Initialize(&_IPayableAccount.TransactOpts, data)
}

// InstallModule is a paid mutator transaction binding the contract method 0x9517e29f.
//
// Solidity: function installModule(uint256 typeId, address module, bytes data) returns()
func (_IPayableAccount *IPayableAccountTransactor) InstallModule(opts *bind.TransactOpts, typeId *big.Int, module common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.contract.Transact(opts, "installModule", typeId, module, data)
}

// InstallModule is a paid mutator transaction binding the contract method 0x9517e29f.
//
// Solidity: function installModule(uint256 typeId, address module, bytes data) returns()
func (_IPayableAccount *IPayableAccountSession) InstallModule(typeId *big.Int, module common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.InstallModule(&_IPayableAccount.TransactOpts, typeId, module, data)
}

// InstallModule is a paid mutator transaction binding the contract method 0x9517e29f.
//
// Solidity: function installModule(uint256 typeId, address module, bytes data) returns()
func (_IPayableAccount *IPayableAccountTransactorSession) InstallModule(typeId *big.Int, module common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.InstallModule(&_IPayableAccount.TransactOpts, typeId, module, data)
}

// InstallRecoveryModule is a paid mutator transaction binding the contract method 0x73f15c02.
//
// Solidity: function installRecoveryModule(address recoveryModule, bytes data) returns()
func (_IPayableAccount *IPayableAccountTransactor) InstallRecoveryModule(opts *bind.TransactOpts, recoveryModule common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.contract.Transact(opts, "installRecoveryModule", recoveryModule, data)
}

// InstallRecoveryModule is a paid mutator transaction binding the contract method 0x73f15c02.
//
// Solidity: function installRecoveryModule(address recoveryModule, bytes data) returns()
func (_IPayableAccount *IPayableAccountSession) InstallRecoveryModule(recoveryModule common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.InstallRecoveryModule(&_IPayableAccount.TransactOpts, recoveryModule, data)
}

// InstallRecoveryModule is a paid mutator transaction binding the contract method 0x73f15c02.
//
// Solidity: function installRecoveryModule(address recoveryModule, bytes data) returns()
func (_IPayableAccount *IPayableAccountTransactorSession) InstallRecoveryModule(recoveryModule common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.InstallRecoveryModule(&_IPayableAccount.TransactOpts, recoveryModule, data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address validator, bytes data) returns()
func (_IPayableAccount *IPayableAccountTransactor) Recover(opts *bind.TransactOpts, validator common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.contract.Transact(opts, "recover", validator, data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address validator, bytes data) returns()
func (_IPayableAccount *IPayableAccountSession) Recover(validator common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.Recover(&_IPayableAccount.TransactOpts, validator, data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address validator, bytes data) returns()
func (_IPayableAccount *IPayableAccountTransactorSession) Recover(validator common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.Recover(&_IPayableAccount.TransactOpts, validator, data)
}

// UninstallModule is a paid mutator transaction binding the contract method 0xa71763a8.
//
// Solidity: function uninstallModule(uint256 typeId, address module, bytes data) returns()
func (_IPayableAccount *IPayableAccountTransactor) UninstallModule(opts *bind.TransactOpts, typeId *big.Int, module common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.contract.Transact(opts, "uninstallModule", typeId, module, data)
}

// UninstallModule is a paid mutator transaction binding the contract method 0xa71763a8.
//
// Solidity: function uninstallModule(uint256 typeId, address module, bytes data) returns()
func (_IPayableAccount *IPayableAccountSession) UninstallModule(typeId *big.Int, module common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.UninstallModule(&_IPayableAccount.TransactOpts, typeId, module, data)
}

// UninstallModule is a paid mutator transaction binding the contract method 0xa71763a8.
//
// Solidity: function uninstallModule(uint256 typeId, address module, bytes data) returns()
func (_IPayableAccount *IPayableAccountTransactorSession) UninstallModule(typeId *big.Int, module common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.UninstallModule(&_IPayableAccount.TransactOpts, typeId, module, data)
}

// UninstallRecoveryModule is a paid mutator transaction binding the contract method 0x2a74b088.
//
// Solidity: function uninstallRecoveryModule(address recoveryModule, bytes data) returns()
func (_IPayableAccount *IPayableAccountTransactor) UninstallRecoveryModule(opts *bind.TransactOpts, recoveryModule common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.contract.Transact(opts, "uninstallRecoveryModule", recoveryModule, data)
}

// UninstallRecoveryModule is a paid mutator transaction binding the contract method 0x2a74b088.
//
// Solidity: function uninstallRecoveryModule(address recoveryModule, bytes data) returns()
func (_IPayableAccount *IPayableAccountSession) UninstallRecoveryModule(recoveryModule common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.UninstallRecoveryModule(&_IPayableAccount.TransactOpts, recoveryModule, data)
}

// UninstallRecoveryModule is a paid mutator transaction binding the contract method 0x2a74b088.
//
// Solidity: function uninstallRecoveryModule(address recoveryModule, bytes data) returns()
func (_IPayableAccount *IPayableAccountTransactorSession) UninstallRecoveryModule(recoveryModule common.Address, data []byte) (*types.Transaction, error) {
	return _IPayableAccount.Contract.UninstallRecoveryModule(&_IPayableAccount.TransactOpts, recoveryModule, data)
}

// IPayableAccountAccountRecoveredIterator is returned from FilterAccountRecovered and is used to iterate over the raw logs and unpacked data for AccountRecovered events raised by the IPayableAccount contract.
type IPayableAccountAccountRecoveredIterator struct {
	Event *IPayableAccountAccountRecovered // Event containing the contract specifics and raw log

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
func (it *IPayableAccountAccountRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayableAccountAccountRecovered)
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
		it.Event = new(IPayableAccountAccountRecovered)
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
func (it *IPayableAccountAccountRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayableAccountAccountRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayableAccountAccountRecovered represents a AccountRecovered event raised by the IPayableAccount contract.
type IPayableAccountAccountRecovered struct {
	Validator common.Address
	Subject   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountRecovered is a free log retrieval operation binding the contract event 0xe3a4fcd6e998a33443a7a770eb4a2aab0c5da9e88b35cdde1f222a143f447e33.
//
// Solidity: event AccountRecovered(address validator, bytes subject)
func (_IPayableAccount *IPayableAccountFilterer) FilterAccountRecovered(opts *bind.FilterOpts) (*IPayableAccountAccountRecoveredIterator, error) {

	logs, sub, err := _IPayableAccount.contract.FilterLogs(opts, "AccountRecovered")
	if err != nil {
		return nil, err
	}
	return &IPayableAccountAccountRecoveredIterator{contract: _IPayableAccount.contract, event: "AccountRecovered", logs: logs, sub: sub}, nil
}

// WatchAccountRecovered is a free log subscription operation binding the contract event 0xe3a4fcd6e998a33443a7a770eb4a2aab0c5da9e88b35cdde1f222a143f447e33.
//
// Solidity: event AccountRecovered(address validator, bytes subject)
func (_IPayableAccount *IPayableAccountFilterer) WatchAccountRecovered(opts *bind.WatchOpts, sink chan<- *IPayableAccountAccountRecovered) (event.Subscription, error) {

	logs, sub, err := _IPayableAccount.contract.WatchLogs(opts, "AccountRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayableAccountAccountRecovered)
				if err := _IPayableAccount.contract.UnpackLog(event, "AccountRecovered", log); err != nil {
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

// ParseAccountRecovered is a log parse operation binding the contract event 0xe3a4fcd6e998a33443a7a770eb4a2aab0c5da9e88b35cdde1f222a143f447e33.
//
// Solidity: event AccountRecovered(address validator, bytes subject)
func (_IPayableAccount *IPayableAccountFilterer) ParseAccountRecovered(log types.Log) (*IPayableAccountAccountRecovered, error) {
	event := new(IPayableAccountAccountRecovered)
	if err := _IPayableAccount.contract.UnpackLog(event, "AccountRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPayableAccountModuleInstalledIterator is returned from FilterModuleInstalled and is used to iterate over the raw logs and unpacked data for ModuleInstalled events raised by the IPayableAccount contract.
type IPayableAccountModuleInstalledIterator struct {
	Event *IPayableAccountModuleInstalled // Event containing the contract specifics and raw log

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
func (it *IPayableAccountModuleInstalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayableAccountModuleInstalled)
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
		it.Event = new(IPayableAccountModuleInstalled)
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
func (it *IPayableAccountModuleInstalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayableAccountModuleInstalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayableAccountModuleInstalled represents a ModuleInstalled event raised by the IPayableAccount contract.
type IPayableAccountModuleInstalled struct {
	ModuleTypeId *big.Int
	Module       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterModuleInstalled is a free log retrieval operation binding the contract event 0xd21d0b289f126c4b473ea641963e766833c2f13866e4ff480abd787c100ef123.
//
// Solidity: event ModuleInstalled(uint256 moduleTypeId, address module)
func (_IPayableAccount *IPayableAccountFilterer) FilterModuleInstalled(opts *bind.FilterOpts) (*IPayableAccountModuleInstalledIterator, error) {

	logs, sub, err := _IPayableAccount.contract.FilterLogs(opts, "ModuleInstalled")
	if err != nil {
		return nil, err
	}
	return &IPayableAccountModuleInstalledIterator{contract: _IPayableAccount.contract, event: "ModuleInstalled", logs: logs, sub: sub}, nil
}

// WatchModuleInstalled is a free log subscription operation binding the contract event 0xd21d0b289f126c4b473ea641963e766833c2f13866e4ff480abd787c100ef123.
//
// Solidity: event ModuleInstalled(uint256 moduleTypeId, address module)
func (_IPayableAccount *IPayableAccountFilterer) WatchModuleInstalled(opts *bind.WatchOpts, sink chan<- *IPayableAccountModuleInstalled) (event.Subscription, error) {

	logs, sub, err := _IPayableAccount.contract.WatchLogs(opts, "ModuleInstalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayableAccountModuleInstalled)
				if err := _IPayableAccount.contract.UnpackLog(event, "ModuleInstalled", log); err != nil {
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

// ParseModuleInstalled is a log parse operation binding the contract event 0xd21d0b289f126c4b473ea641963e766833c2f13866e4ff480abd787c100ef123.
//
// Solidity: event ModuleInstalled(uint256 moduleTypeId, address module)
func (_IPayableAccount *IPayableAccountFilterer) ParseModuleInstalled(log types.Log) (*IPayableAccountModuleInstalled, error) {
	event := new(IPayableAccountModuleInstalled)
	if err := _IPayableAccount.contract.UnpackLog(event, "ModuleInstalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPayableAccountModuleInstalledWithDataIterator is returned from FilterModuleInstalledWithData and is used to iterate over the raw logs and unpacked data for ModuleInstalledWithData events raised by the IPayableAccount contract.
type IPayableAccountModuleInstalledWithDataIterator struct {
	Event *IPayableAccountModuleInstalledWithData // Event containing the contract specifics and raw log

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
func (it *IPayableAccountModuleInstalledWithDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayableAccountModuleInstalledWithData)
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
		it.Event = new(IPayableAccountModuleInstalledWithData)
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
func (it *IPayableAccountModuleInstalledWithDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayableAccountModuleInstalledWithDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayableAccountModuleInstalledWithData represents a ModuleInstalledWithData event raised by the IPayableAccount contract.
type IPayableAccountModuleInstalledWithData struct {
	TypeId *big.Int
	Module common.Address
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterModuleInstalledWithData is a free log retrieval operation binding the contract event 0x2ee31983ec5cbb71ac0f5b0cee054a7bd1771e18ba75efa5b30b8e62696bc9da.
//
// Solidity: event ModuleInstalledWithData(uint256 typeId, address module, bytes data)
func (_IPayableAccount *IPayableAccountFilterer) FilterModuleInstalledWithData(opts *bind.FilterOpts) (*IPayableAccountModuleInstalledWithDataIterator, error) {

	logs, sub, err := _IPayableAccount.contract.FilterLogs(opts, "ModuleInstalledWithData")
	if err != nil {
		return nil, err
	}
	return &IPayableAccountModuleInstalledWithDataIterator{contract: _IPayableAccount.contract, event: "ModuleInstalledWithData", logs: logs, sub: sub}, nil
}

// WatchModuleInstalledWithData is a free log subscription operation binding the contract event 0x2ee31983ec5cbb71ac0f5b0cee054a7bd1771e18ba75efa5b30b8e62696bc9da.
//
// Solidity: event ModuleInstalledWithData(uint256 typeId, address module, bytes data)
func (_IPayableAccount *IPayableAccountFilterer) WatchModuleInstalledWithData(opts *bind.WatchOpts, sink chan<- *IPayableAccountModuleInstalledWithData) (event.Subscription, error) {

	logs, sub, err := _IPayableAccount.contract.WatchLogs(opts, "ModuleInstalledWithData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayableAccountModuleInstalledWithData)
				if err := _IPayableAccount.contract.UnpackLog(event, "ModuleInstalledWithData", log); err != nil {
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

// ParseModuleInstalledWithData is a log parse operation binding the contract event 0x2ee31983ec5cbb71ac0f5b0cee054a7bd1771e18ba75efa5b30b8e62696bc9da.
//
// Solidity: event ModuleInstalledWithData(uint256 typeId, address module, bytes data)
func (_IPayableAccount *IPayableAccountFilterer) ParseModuleInstalledWithData(log types.Log) (*IPayableAccountModuleInstalledWithData, error) {
	event := new(IPayableAccountModuleInstalledWithData)
	if err := _IPayableAccount.contract.UnpackLog(event, "ModuleInstalledWithData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPayableAccountModuleUninstalledIterator is returned from FilterModuleUninstalled and is used to iterate over the raw logs and unpacked data for ModuleUninstalled events raised by the IPayableAccount contract.
type IPayableAccountModuleUninstalledIterator struct {
	Event *IPayableAccountModuleUninstalled // Event containing the contract specifics and raw log

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
func (it *IPayableAccountModuleUninstalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayableAccountModuleUninstalled)
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
		it.Event = new(IPayableAccountModuleUninstalled)
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
func (it *IPayableAccountModuleUninstalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayableAccountModuleUninstalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayableAccountModuleUninstalled represents a ModuleUninstalled event raised by the IPayableAccount contract.
type IPayableAccountModuleUninstalled struct {
	ModuleTypeId *big.Int
	Module       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterModuleUninstalled is a free log retrieval operation binding the contract event 0x341347516a9de374859dfda710fa4828b2d48cb57d4fbe4c1149612b8e02276e.
//
// Solidity: event ModuleUninstalled(uint256 moduleTypeId, address module)
func (_IPayableAccount *IPayableAccountFilterer) FilterModuleUninstalled(opts *bind.FilterOpts) (*IPayableAccountModuleUninstalledIterator, error) {

	logs, sub, err := _IPayableAccount.contract.FilterLogs(opts, "ModuleUninstalled")
	if err != nil {
		return nil, err
	}
	return &IPayableAccountModuleUninstalledIterator{contract: _IPayableAccount.contract, event: "ModuleUninstalled", logs: logs, sub: sub}, nil
}

// WatchModuleUninstalled is a free log subscription operation binding the contract event 0x341347516a9de374859dfda710fa4828b2d48cb57d4fbe4c1149612b8e02276e.
//
// Solidity: event ModuleUninstalled(uint256 moduleTypeId, address module)
func (_IPayableAccount *IPayableAccountFilterer) WatchModuleUninstalled(opts *bind.WatchOpts, sink chan<- *IPayableAccountModuleUninstalled) (event.Subscription, error) {

	logs, sub, err := _IPayableAccount.contract.WatchLogs(opts, "ModuleUninstalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayableAccountModuleUninstalled)
				if err := _IPayableAccount.contract.UnpackLog(event, "ModuleUninstalled", log); err != nil {
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

// ParseModuleUninstalled is a log parse operation binding the contract event 0x341347516a9de374859dfda710fa4828b2d48cb57d4fbe4c1149612b8e02276e.
//
// Solidity: event ModuleUninstalled(uint256 moduleTypeId, address module)
func (_IPayableAccount *IPayableAccountFilterer) ParseModuleUninstalled(log types.Log) (*IPayableAccountModuleUninstalled, error) {
	event := new(IPayableAccountModuleUninstalled)
	if err := _IPayableAccount.contract.UnpackLog(event, "ModuleUninstalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPayableAccountModuleUninstalledWithDataIterator is returned from FilterModuleUninstalledWithData and is used to iterate over the raw logs and unpacked data for ModuleUninstalledWithData events raised by the IPayableAccount contract.
type IPayableAccountModuleUninstalledWithDataIterator struct {
	Event *IPayableAccountModuleUninstalledWithData // Event containing the contract specifics and raw log

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
func (it *IPayableAccountModuleUninstalledWithDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayableAccountModuleUninstalledWithData)
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
		it.Event = new(IPayableAccountModuleUninstalledWithData)
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
func (it *IPayableAccountModuleUninstalledWithDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayableAccountModuleUninstalledWithDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayableAccountModuleUninstalledWithData represents a ModuleUninstalledWithData event raised by the IPayableAccount contract.
type IPayableAccountModuleUninstalledWithData struct {
	TypeId *big.Int
	Module common.Address
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterModuleUninstalledWithData is a free log retrieval operation binding the contract event 0xe0f0a339f8c9b3ec44c5d7a3f23d82baf6e9d8b332372fde4b79d0a4dd89ada0.
//
// Solidity: event ModuleUninstalledWithData(uint256 typeId, address module, bytes data)
func (_IPayableAccount *IPayableAccountFilterer) FilterModuleUninstalledWithData(opts *bind.FilterOpts) (*IPayableAccountModuleUninstalledWithDataIterator, error) {

	logs, sub, err := _IPayableAccount.contract.FilterLogs(opts, "ModuleUninstalledWithData")
	if err != nil {
		return nil, err
	}
	return &IPayableAccountModuleUninstalledWithDataIterator{contract: _IPayableAccount.contract, event: "ModuleUninstalledWithData", logs: logs, sub: sub}, nil
}

// WatchModuleUninstalledWithData is a free log subscription operation binding the contract event 0xe0f0a339f8c9b3ec44c5d7a3f23d82baf6e9d8b332372fde4b79d0a4dd89ada0.
//
// Solidity: event ModuleUninstalledWithData(uint256 typeId, address module, bytes data)
func (_IPayableAccount *IPayableAccountFilterer) WatchModuleUninstalledWithData(opts *bind.WatchOpts, sink chan<- *IPayableAccountModuleUninstalledWithData) (event.Subscription, error) {

	logs, sub, err := _IPayableAccount.contract.WatchLogs(opts, "ModuleUninstalledWithData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayableAccountModuleUninstalledWithData)
				if err := _IPayableAccount.contract.UnpackLog(event, "ModuleUninstalledWithData", log); err != nil {
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

// ParseModuleUninstalledWithData is a log parse operation binding the contract event 0xe0f0a339f8c9b3ec44c5d7a3f23d82baf6e9d8b332372fde4b79d0a4dd89ada0.
//
// Solidity: event ModuleUninstalledWithData(uint256 typeId, address module, bytes data)
func (_IPayableAccount *IPayableAccountFilterer) ParseModuleUninstalledWithData(log types.Log) (*IPayableAccountModuleUninstalledWithData, error) {
	event := new(IPayableAccountModuleUninstalledWithData)
	if err := _IPayableAccount.contract.UnpackLog(event, "ModuleUninstalledWithData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPayableAccountRecoveryFeeClaimedIterator is returned from FilterRecoveryFeeClaimed and is used to iterate over the raw logs and unpacked data for RecoveryFeeClaimed events raised by the IPayableAccount contract.
type IPayableAccountRecoveryFeeClaimedIterator struct {
	Event *IPayableAccountRecoveryFeeClaimed // Event containing the contract specifics and raw log

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
func (it *IPayableAccountRecoveryFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayableAccountRecoveryFeeClaimed)
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
		it.Event = new(IPayableAccountRecoveryFeeClaimed)
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
func (it *IPayableAccountRecoveryFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayableAccountRecoveryFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayableAccountRecoveryFeeClaimed represents a RecoveryFeeClaimed event raised by the IPayableAccount contract.
type IPayableAccountRecoveryFeeClaimed struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoveryFeeClaimed is a free log retrieval operation binding the contract event 0xad157ab5a41b214c289a98351f7e5dadda83054e32879678fa7926fa904c0c87.
//
// Solidity: event RecoveryFeeClaimed(address to, uint256 amount)
func (_IPayableAccount *IPayableAccountFilterer) FilterRecoveryFeeClaimed(opts *bind.FilterOpts) (*IPayableAccountRecoveryFeeClaimedIterator, error) {

	logs, sub, err := _IPayableAccount.contract.FilterLogs(opts, "RecoveryFeeClaimed")
	if err != nil {
		return nil, err
	}
	return &IPayableAccountRecoveryFeeClaimedIterator{contract: _IPayableAccount.contract, event: "RecoveryFeeClaimed", logs: logs, sub: sub}, nil
}

// WatchRecoveryFeeClaimed is a free log subscription operation binding the contract event 0xad157ab5a41b214c289a98351f7e5dadda83054e32879678fa7926fa904c0c87.
//
// Solidity: event RecoveryFeeClaimed(address to, uint256 amount)
func (_IPayableAccount *IPayableAccountFilterer) WatchRecoveryFeeClaimed(opts *bind.WatchOpts, sink chan<- *IPayableAccountRecoveryFeeClaimed) (event.Subscription, error) {

	logs, sub, err := _IPayableAccount.contract.WatchLogs(opts, "RecoveryFeeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayableAccountRecoveryFeeClaimed)
				if err := _IPayableAccount.contract.UnpackLog(event, "RecoveryFeeClaimed", log); err != nil {
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

// ParseRecoveryFeeClaimed is a log parse operation binding the contract event 0xad157ab5a41b214c289a98351f7e5dadda83054e32879678fa7926fa904c0c87.
//
// Solidity: event RecoveryFeeClaimed(address to, uint256 amount)
func (_IPayableAccount *IPayableAccountFilterer) ParseRecoveryFeeClaimed(log types.Log) (*IPayableAccountRecoveryFeeClaimed, error) {
	event := new(IPayableAccountRecoveryFeeClaimed)
	if err := _IPayableAccount.contract.UnpackLog(event, "RecoveryFeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPayableAccountRecoveryModuleInstalledIterator is returned from FilterRecoveryModuleInstalled and is used to iterate over the raw logs and unpacked data for RecoveryModuleInstalled events raised by the IPayableAccount contract.
type IPayableAccountRecoveryModuleInstalledIterator struct {
	Event *IPayableAccountRecoveryModuleInstalled // Event containing the contract specifics and raw log

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
func (it *IPayableAccountRecoveryModuleInstalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayableAccountRecoveryModuleInstalled)
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
		it.Event = new(IPayableAccountRecoveryModuleInstalled)
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
func (it *IPayableAccountRecoveryModuleInstalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayableAccountRecoveryModuleInstalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayableAccountRecoveryModuleInstalled represents a RecoveryModuleInstalled event raised by the IPayableAccount contract.
type IPayableAccountRecoveryModuleInstalled struct {
	RecoveryModule common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRecoveryModuleInstalled is a free log retrieval operation binding the contract event 0x6cfe7eb4ddea48bed2901168a3eda97a7c9ab17b5ce9230670b19cbdd632af44.
//
// Solidity: event RecoveryModuleInstalled(address recoveryModule)
func (_IPayableAccount *IPayableAccountFilterer) FilterRecoveryModuleInstalled(opts *bind.FilterOpts) (*IPayableAccountRecoveryModuleInstalledIterator, error) {

	logs, sub, err := _IPayableAccount.contract.FilterLogs(opts, "RecoveryModuleInstalled")
	if err != nil {
		return nil, err
	}
	return &IPayableAccountRecoveryModuleInstalledIterator{contract: _IPayableAccount.contract, event: "RecoveryModuleInstalled", logs: logs, sub: sub}, nil
}

// WatchRecoveryModuleInstalled is a free log subscription operation binding the contract event 0x6cfe7eb4ddea48bed2901168a3eda97a7c9ab17b5ce9230670b19cbdd632af44.
//
// Solidity: event RecoveryModuleInstalled(address recoveryModule)
func (_IPayableAccount *IPayableAccountFilterer) WatchRecoveryModuleInstalled(opts *bind.WatchOpts, sink chan<- *IPayableAccountRecoveryModuleInstalled) (event.Subscription, error) {

	logs, sub, err := _IPayableAccount.contract.WatchLogs(opts, "RecoveryModuleInstalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayableAccountRecoveryModuleInstalled)
				if err := _IPayableAccount.contract.UnpackLog(event, "RecoveryModuleInstalled", log); err != nil {
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

// ParseRecoveryModuleInstalled is a log parse operation binding the contract event 0x6cfe7eb4ddea48bed2901168a3eda97a7c9ab17b5ce9230670b19cbdd632af44.
//
// Solidity: event RecoveryModuleInstalled(address recoveryModule)
func (_IPayableAccount *IPayableAccountFilterer) ParseRecoveryModuleInstalled(log types.Log) (*IPayableAccountRecoveryModuleInstalled, error) {
	event := new(IPayableAccountRecoveryModuleInstalled)
	if err := _IPayableAccount.contract.UnpackLog(event, "RecoveryModuleInstalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPayableAccountRecoveryModuleUninstalledIterator is returned from FilterRecoveryModuleUninstalled and is used to iterate over the raw logs and unpacked data for RecoveryModuleUninstalled events raised by the IPayableAccount contract.
type IPayableAccountRecoveryModuleUninstalledIterator struct {
	Event *IPayableAccountRecoveryModuleUninstalled // Event containing the contract specifics and raw log

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
func (it *IPayableAccountRecoveryModuleUninstalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayableAccountRecoveryModuleUninstalled)
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
		it.Event = new(IPayableAccountRecoveryModuleUninstalled)
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
func (it *IPayableAccountRecoveryModuleUninstalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayableAccountRecoveryModuleUninstalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayableAccountRecoveryModuleUninstalled represents a RecoveryModuleUninstalled event raised by the IPayableAccount contract.
type IPayableAccountRecoveryModuleUninstalled struct {
	RecoveryModule common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRecoveryModuleUninstalled is a free log retrieval operation binding the contract event 0xe4c9c992ee2b4d5eee8cc0f0f25ad7ce7868ea3c76d06421e476f6adf240fc6d.
//
// Solidity: event RecoveryModuleUninstalled(address recoveryModule)
func (_IPayableAccount *IPayableAccountFilterer) FilterRecoveryModuleUninstalled(opts *bind.FilterOpts) (*IPayableAccountRecoveryModuleUninstalledIterator, error) {

	logs, sub, err := _IPayableAccount.contract.FilterLogs(opts, "RecoveryModuleUninstalled")
	if err != nil {
		return nil, err
	}
	return &IPayableAccountRecoveryModuleUninstalledIterator{contract: _IPayableAccount.contract, event: "RecoveryModuleUninstalled", logs: logs, sub: sub}, nil
}

// WatchRecoveryModuleUninstalled is a free log subscription operation binding the contract event 0xe4c9c992ee2b4d5eee8cc0f0f25ad7ce7868ea3c76d06421e476f6adf240fc6d.
//
// Solidity: event RecoveryModuleUninstalled(address recoveryModule)
func (_IPayableAccount *IPayableAccountFilterer) WatchRecoveryModuleUninstalled(opts *bind.WatchOpts, sink chan<- *IPayableAccountRecoveryModuleUninstalled) (event.Subscription, error) {

	logs, sub, err := _IPayableAccount.contract.WatchLogs(opts, "RecoveryModuleUninstalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayableAccountRecoveryModuleUninstalled)
				if err := _IPayableAccount.contract.UnpackLog(event, "RecoveryModuleUninstalled", log); err != nil {
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

// ParseRecoveryModuleUninstalled is a log parse operation binding the contract event 0xe4c9c992ee2b4d5eee8cc0f0f25ad7ce7868ea3c76d06421e476f6adf240fc6d.
//
// Solidity: event RecoveryModuleUninstalled(address recoveryModule)
func (_IPayableAccount *IPayableAccountFilterer) ParseRecoveryModuleUninstalled(log types.Log) (*IPayableAccountRecoveryModuleUninstalled, error) {
	event := new(IPayableAccountRecoveryModuleUninstalled)
	if err := _IPayableAccount.contract.UnpackLog(event, "RecoveryModuleUninstalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPayableAccountSafeReceivedIterator is returned from FilterSafeReceived and is used to iterate over the raw logs and unpacked data for SafeReceived events raised by the IPayableAccount contract.
type IPayableAccountSafeReceivedIterator struct {
	Event *IPayableAccountSafeReceived // Event containing the contract specifics and raw log

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
func (it *IPayableAccountSafeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPayableAccountSafeReceived)
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
		it.Event = new(IPayableAccountSafeReceived)
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
func (it *IPayableAccountSafeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPayableAccountSafeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPayableAccountSafeReceived represents a SafeReceived event raised by the IPayableAccount contract.
type IPayableAccountSafeReceived struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSafeReceived is a free log retrieval operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address sender, uint256 value)
func (_IPayableAccount *IPayableAccountFilterer) FilterSafeReceived(opts *bind.FilterOpts) (*IPayableAccountSafeReceivedIterator, error) {

	logs, sub, err := _IPayableAccount.contract.FilterLogs(opts, "SafeReceived")
	if err != nil {
		return nil, err
	}
	return &IPayableAccountSafeReceivedIterator{contract: _IPayableAccount.contract, event: "SafeReceived", logs: logs, sub: sub}, nil
}

// WatchSafeReceived is a free log subscription operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address sender, uint256 value)
func (_IPayableAccount *IPayableAccountFilterer) WatchSafeReceived(opts *bind.WatchOpts, sink chan<- *IPayableAccountSafeReceived) (event.Subscription, error) {

	logs, sub, err := _IPayableAccount.contract.WatchLogs(opts, "SafeReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPayableAccountSafeReceived)
				if err := _IPayableAccount.contract.UnpackLog(event, "SafeReceived", log); err != nil {
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

// ParseSafeReceived is a log parse operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address sender, uint256 value)
func (_IPayableAccount *IPayableAccountFilterer) ParseSafeReceived(log types.Log) (*IPayableAccountSafeReceived, error) {
	event := new(IPayableAccountSafeReceived)
	if err := _IPayableAccount.contract.UnpackLog(event, "SafeReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
