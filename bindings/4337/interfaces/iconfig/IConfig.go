// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iconfig

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

// IConfigMetaData contains all meta data concerning the IConfig contract.
var IConfigMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"FactorySignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"PaySignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"RecoverySignerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"RecoverySignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetSafeSingleton\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetSenderSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumVerifierType\",\"name\":\"verifier\",\"type\":\"uint8\"}],\"name\":\"VerifierTypeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumVerifierType\",\"name\":\"verifier\",\"type\":\"uint8\"}],\"name\":\"VerifierTypeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"WhitelistBundlerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"WhitelistBundlerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moduleType\",\"type\":\"uint256\"}],\"name\":\"WhitelistModuleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"WhitelistModuleRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"addRecoverySigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"addWhitelistedBundlers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"moduleTypes\",\"type\":\"uint256[]\"}],\"name\":\"addWhitelistedModules\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumVerifierType\",\"name\":\"verifier\",\"type\":\"uint8\"}],\"name\":\"addWhitelistedVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"initializeWalletSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isFactorySigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isPaySigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isRecoverySigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"}],\"name\":\"isSafeSingleton\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bundler\",\"type\":\"address\"}],\"name\":\"isWhitelistedBundler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumVerifierType\",\"name\":\"verifier\",\"type\":\"uint8\"}],\"name\":\"isWhitelistedVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"removeRecoverySigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"removeWhitelistedBundlers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"name\":\"removeWhitelistedModules\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumVerifierType\",\"name\":\"verifier\",\"type\":\"uint8\"}],\"name\":\"removeWhitelistedVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"walletSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"whitelistedModuleType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleType\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use IConfigMetaData.ABI instead.
var IConfigABI = IConfigMetaData.ABI

// IConfig is an auto generated Go binding around an Ethereum contract.
type IConfig struct {
	IConfigCaller     // Read-only binding to the contract
	IConfigTransactor // Write-only binding to the contract
	IConfigFilterer   // Log filterer for contract events
}

// IConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type IConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IConfigSession struct {
	Contract     *IConfig          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IConfigCallerSession struct {
	Contract *IConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IConfigTransactorSession struct {
	Contract     *IConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type IConfigRaw struct {
	Contract *IConfig // Generic contract binding to access the raw methods on
}

// IConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IConfigCallerRaw struct {
	Contract *IConfigCaller // Generic read-only contract binding to access the raw methods on
}

// IConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IConfigTransactorRaw struct {
	Contract *IConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIConfig creates a new instance of IConfig, bound to a specific deployed contract.
func NewIConfig(address common.Address, backend bind.ContractBackend) (*IConfig, error) {
	contract, err := bindIConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IConfig{IConfigCaller: IConfigCaller{contract: contract}, IConfigTransactor: IConfigTransactor{contract: contract}, IConfigFilterer: IConfigFilterer{contract: contract}}, nil
}

// NewIConfigCaller creates a new read-only instance of IConfig, bound to a specific deployed contract.
func NewIConfigCaller(address common.Address, caller bind.ContractCaller) (*IConfigCaller, error) {
	contract, err := bindIConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IConfigCaller{contract: contract}, nil
}

// NewIConfigTransactor creates a new write-only instance of IConfig, bound to a specific deployed contract.
func NewIConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*IConfigTransactor, error) {
	contract, err := bindIConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IConfigTransactor{contract: contract}, nil
}

// NewIConfigFilterer creates a new log filterer instance of IConfig, bound to a specific deployed contract.
func NewIConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*IConfigFilterer, error) {
	contract, err := bindIConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IConfigFilterer{contract: contract}, nil
}

// bindIConfig binds a generic wrapper to an already deployed contract.
func bindIConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IConfig *IConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IConfig.Contract.IConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IConfig *IConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IConfig.Contract.IConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IConfig *IConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IConfig.Contract.IConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IConfig *IConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IConfig *IConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IConfig *IConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IConfig.Contract.contract.Transact(opts, method, params...)
}

// IsFactorySigner is a free data retrieval call binding the contract method 0xdcf0da72.
//
// Solidity: function isFactorySigner(address sender) view returns(bool)
func (_IConfig *IConfigCaller) IsFactorySigner(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _IConfig.contract.Call(opts, &out, "isFactorySigner", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFactorySigner is a free data retrieval call binding the contract method 0xdcf0da72.
//
// Solidity: function isFactorySigner(address sender) view returns(bool)
func (_IConfig *IConfigSession) IsFactorySigner(sender common.Address) (bool, error) {
	return _IConfig.Contract.IsFactorySigner(&_IConfig.CallOpts, sender)
}

// IsFactorySigner is a free data retrieval call binding the contract method 0xdcf0da72.
//
// Solidity: function isFactorySigner(address sender) view returns(bool)
func (_IConfig *IConfigCallerSession) IsFactorySigner(sender common.Address) (bool, error) {
	return _IConfig.Contract.IsFactorySigner(&_IConfig.CallOpts, sender)
}

// IsPaySigner is a free data retrieval call binding the contract method 0xe53a7eae.
//
// Solidity: function isPaySigner(address sender) view returns(bool)
func (_IConfig *IConfigCaller) IsPaySigner(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _IConfig.contract.Call(opts, &out, "isPaySigner", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaySigner is a free data retrieval call binding the contract method 0xe53a7eae.
//
// Solidity: function isPaySigner(address sender) view returns(bool)
func (_IConfig *IConfigSession) IsPaySigner(sender common.Address) (bool, error) {
	return _IConfig.Contract.IsPaySigner(&_IConfig.CallOpts, sender)
}

// IsPaySigner is a free data retrieval call binding the contract method 0xe53a7eae.
//
// Solidity: function isPaySigner(address sender) view returns(bool)
func (_IConfig *IConfigCallerSession) IsPaySigner(sender common.Address) (bool, error) {
	return _IConfig.Contract.IsPaySigner(&_IConfig.CallOpts, sender)
}

// IsRecoverySigner is a free data retrieval call binding the contract method 0x96a5d35b.
//
// Solidity: function isRecoverySigner(address signer) view returns(bool)
func (_IConfig *IConfigCaller) IsRecoverySigner(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _IConfig.contract.Call(opts, &out, "isRecoverySigner", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRecoverySigner is a free data retrieval call binding the contract method 0x96a5d35b.
//
// Solidity: function isRecoverySigner(address signer) view returns(bool)
func (_IConfig *IConfigSession) IsRecoverySigner(signer common.Address) (bool, error) {
	return _IConfig.Contract.IsRecoverySigner(&_IConfig.CallOpts, signer)
}

// IsRecoverySigner is a free data retrieval call binding the contract method 0x96a5d35b.
//
// Solidity: function isRecoverySigner(address signer) view returns(bool)
func (_IConfig *IConfigCallerSession) IsRecoverySigner(signer common.Address) (bool, error) {
	return _IConfig.Contract.IsRecoverySigner(&_IConfig.CallOpts, signer)
}

// IsSafeSingleton is a free data retrieval call binding the contract method 0xc48a5898.
//
// Solidity: function isSafeSingleton(address singleton) view returns(bool)
func (_IConfig *IConfigCaller) IsSafeSingleton(opts *bind.CallOpts, singleton common.Address) (bool, error) {
	var out []interface{}
	err := _IConfig.contract.Call(opts, &out, "isSafeSingleton", singleton)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSafeSingleton is a free data retrieval call binding the contract method 0xc48a5898.
//
// Solidity: function isSafeSingleton(address singleton) view returns(bool)
func (_IConfig *IConfigSession) IsSafeSingleton(singleton common.Address) (bool, error) {
	return _IConfig.Contract.IsSafeSingleton(&_IConfig.CallOpts, singleton)
}

// IsSafeSingleton is a free data retrieval call binding the contract method 0xc48a5898.
//
// Solidity: function isSafeSingleton(address singleton) view returns(bool)
func (_IConfig *IConfigCallerSession) IsSafeSingleton(singleton common.Address) (bool, error) {
	return _IConfig.Contract.IsSafeSingleton(&_IConfig.CallOpts, singleton)
}

// IsWhitelistedBundler is a free data retrieval call binding the contract method 0x93658783.
//
// Solidity: function isWhitelistedBundler(address bundler) view returns(bool)
func (_IConfig *IConfigCaller) IsWhitelistedBundler(opts *bind.CallOpts, bundler common.Address) (bool, error) {
	var out []interface{}
	err := _IConfig.contract.Call(opts, &out, "isWhitelistedBundler", bundler)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedBundler is a free data retrieval call binding the contract method 0x93658783.
//
// Solidity: function isWhitelistedBundler(address bundler) view returns(bool)
func (_IConfig *IConfigSession) IsWhitelistedBundler(bundler common.Address) (bool, error) {
	return _IConfig.Contract.IsWhitelistedBundler(&_IConfig.CallOpts, bundler)
}

// IsWhitelistedBundler is a free data retrieval call binding the contract method 0x93658783.
//
// Solidity: function isWhitelistedBundler(address bundler) view returns(bool)
func (_IConfig *IConfigCallerSession) IsWhitelistedBundler(bundler common.Address) (bool, error) {
	return _IConfig.Contract.IsWhitelistedBundler(&_IConfig.CallOpts, bundler)
}

// IsWhitelistedVerifier is a free data retrieval call binding the contract method 0xb55b7066.
//
// Solidity: function isWhitelistedVerifier(uint8 verifier) view returns(bool)
func (_IConfig *IConfigCaller) IsWhitelistedVerifier(opts *bind.CallOpts, verifier uint8) (bool, error) {
	var out []interface{}
	err := _IConfig.contract.Call(opts, &out, "isWhitelistedVerifier", verifier)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedVerifier is a free data retrieval call binding the contract method 0xb55b7066.
//
// Solidity: function isWhitelistedVerifier(uint8 verifier) view returns(bool)
func (_IConfig *IConfigSession) IsWhitelistedVerifier(verifier uint8) (bool, error) {
	return _IConfig.Contract.IsWhitelistedVerifier(&_IConfig.CallOpts, verifier)
}

// IsWhitelistedVerifier is a free data retrieval call binding the contract method 0xb55b7066.
//
// Solidity: function isWhitelistedVerifier(uint8 verifier) view returns(bool)
func (_IConfig *IConfigCallerSession) IsWhitelistedVerifier(verifier uint8) (bool, error) {
	return _IConfig.Contract.IsWhitelistedVerifier(&_IConfig.CallOpts, verifier)
}

// WalletSigner is a free data retrieval call binding the contract method 0x25d480ad.
//
// Solidity: function walletSigner(address sender) view returns(address)
func (_IConfig *IConfigCaller) WalletSigner(opts *bind.CallOpts, sender common.Address) (common.Address, error) {
	var out []interface{}
	err := _IConfig.contract.Call(opts, &out, "walletSigner", sender)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletSigner is a free data retrieval call binding the contract method 0x25d480ad.
//
// Solidity: function walletSigner(address sender) view returns(address)
func (_IConfig *IConfigSession) WalletSigner(sender common.Address) (common.Address, error) {
	return _IConfig.Contract.WalletSigner(&_IConfig.CallOpts, sender)
}

// WalletSigner is a free data retrieval call binding the contract method 0x25d480ad.
//
// Solidity: function walletSigner(address sender) view returns(address)
func (_IConfig *IConfigCallerSession) WalletSigner(sender common.Address) (common.Address, error) {
	return _IConfig.Contract.WalletSigner(&_IConfig.CallOpts, sender)
}

// WhitelistedModuleType is a free data retrieval call binding the contract method 0x589597cb.
//
// Solidity: function whitelistedModuleType(address module) view returns(uint256 moduleType)
func (_IConfig *IConfigCaller) WhitelistedModuleType(opts *bind.CallOpts, module common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IConfig.contract.Call(opts, &out, "whitelistedModuleType", module)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistedModuleType is a free data retrieval call binding the contract method 0x589597cb.
//
// Solidity: function whitelistedModuleType(address module) view returns(uint256 moduleType)
func (_IConfig *IConfigSession) WhitelistedModuleType(module common.Address) (*big.Int, error) {
	return _IConfig.Contract.WhitelistedModuleType(&_IConfig.CallOpts, module)
}

// WhitelistedModuleType is a free data retrieval call binding the contract method 0x589597cb.
//
// Solidity: function whitelistedModuleType(address module) view returns(uint256 moduleType)
func (_IConfig *IConfigCallerSession) WhitelistedModuleType(module common.Address) (*big.Int, error) {
	return _IConfig.Contract.WhitelistedModuleType(&_IConfig.CallOpts, module)
}

// AddRecoverySigner is a paid mutator transaction binding the contract method 0x32b0cc32.
//
// Solidity: function addRecoverySigner(address signer) returns()
func (_IConfig *IConfigTransactor) AddRecoverySigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _IConfig.contract.Transact(opts, "addRecoverySigner", signer)
}

// AddRecoverySigner is a paid mutator transaction binding the contract method 0x32b0cc32.
//
// Solidity: function addRecoverySigner(address signer) returns()
func (_IConfig *IConfigSession) AddRecoverySigner(signer common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.AddRecoverySigner(&_IConfig.TransactOpts, signer)
}

// AddRecoverySigner is a paid mutator transaction binding the contract method 0x32b0cc32.
//
// Solidity: function addRecoverySigner(address signer) returns()
func (_IConfig *IConfigTransactorSession) AddRecoverySigner(signer common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.AddRecoverySigner(&_IConfig.TransactOpts, signer)
}

// AddWhitelistedBundlers is a paid mutator transaction binding the contract method 0xfcdc4727.
//
// Solidity: function addWhitelistedBundlers(address[] bundlers) returns()
func (_IConfig *IConfigTransactor) AddWhitelistedBundlers(opts *bind.TransactOpts, bundlers []common.Address) (*types.Transaction, error) {
	return _IConfig.contract.Transact(opts, "addWhitelistedBundlers", bundlers)
}

// AddWhitelistedBundlers is a paid mutator transaction binding the contract method 0xfcdc4727.
//
// Solidity: function addWhitelistedBundlers(address[] bundlers) returns()
func (_IConfig *IConfigSession) AddWhitelistedBundlers(bundlers []common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.AddWhitelistedBundlers(&_IConfig.TransactOpts, bundlers)
}

// AddWhitelistedBundlers is a paid mutator transaction binding the contract method 0xfcdc4727.
//
// Solidity: function addWhitelistedBundlers(address[] bundlers) returns()
func (_IConfig *IConfigTransactorSession) AddWhitelistedBundlers(bundlers []common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.AddWhitelistedBundlers(&_IConfig.TransactOpts, bundlers)
}

// AddWhitelistedModules is a paid mutator transaction binding the contract method 0xd4c4c2b0.
//
// Solidity: function addWhitelistedModules(address[] modules, uint256[] moduleTypes) returns()
func (_IConfig *IConfigTransactor) AddWhitelistedModules(opts *bind.TransactOpts, modules []common.Address, moduleTypes []*big.Int) (*types.Transaction, error) {
	return _IConfig.contract.Transact(opts, "addWhitelistedModules", modules, moduleTypes)
}

// AddWhitelistedModules is a paid mutator transaction binding the contract method 0xd4c4c2b0.
//
// Solidity: function addWhitelistedModules(address[] modules, uint256[] moduleTypes) returns()
func (_IConfig *IConfigSession) AddWhitelistedModules(modules []common.Address, moduleTypes []*big.Int) (*types.Transaction, error) {
	return _IConfig.Contract.AddWhitelistedModules(&_IConfig.TransactOpts, modules, moduleTypes)
}

// AddWhitelistedModules is a paid mutator transaction binding the contract method 0xd4c4c2b0.
//
// Solidity: function addWhitelistedModules(address[] modules, uint256[] moduleTypes) returns()
func (_IConfig *IConfigTransactorSession) AddWhitelistedModules(modules []common.Address, moduleTypes []*big.Int) (*types.Transaction, error) {
	return _IConfig.Contract.AddWhitelistedModules(&_IConfig.TransactOpts, modules, moduleTypes)
}

// AddWhitelistedVerifier is a paid mutator transaction binding the contract method 0x81e36802.
//
// Solidity: function addWhitelistedVerifier(uint8 verifier) returns()
func (_IConfig *IConfigTransactor) AddWhitelistedVerifier(opts *bind.TransactOpts, verifier uint8) (*types.Transaction, error) {
	return _IConfig.contract.Transact(opts, "addWhitelistedVerifier", verifier)
}

// AddWhitelistedVerifier is a paid mutator transaction binding the contract method 0x81e36802.
//
// Solidity: function addWhitelistedVerifier(uint8 verifier) returns()
func (_IConfig *IConfigSession) AddWhitelistedVerifier(verifier uint8) (*types.Transaction, error) {
	return _IConfig.Contract.AddWhitelistedVerifier(&_IConfig.TransactOpts, verifier)
}

// AddWhitelistedVerifier is a paid mutator transaction binding the contract method 0x81e36802.
//
// Solidity: function addWhitelistedVerifier(uint8 verifier) returns()
func (_IConfig *IConfigTransactorSession) AddWhitelistedVerifier(verifier uint8) (*types.Transaction, error) {
	return _IConfig.Contract.AddWhitelistedVerifier(&_IConfig.TransactOpts, verifier)
}

// InitializeWalletSigner is a paid mutator transaction binding the contract method 0x4b04ad51.
//
// Solidity: function initializeWalletSigner(address signer) returns()
func (_IConfig *IConfigTransactor) InitializeWalletSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _IConfig.contract.Transact(opts, "initializeWalletSigner", signer)
}

// InitializeWalletSigner is a paid mutator transaction binding the contract method 0x4b04ad51.
//
// Solidity: function initializeWalletSigner(address signer) returns()
func (_IConfig *IConfigSession) InitializeWalletSigner(signer common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.InitializeWalletSigner(&_IConfig.TransactOpts, signer)
}

// InitializeWalletSigner is a paid mutator transaction binding the contract method 0x4b04ad51.
//
// Solidity: function initializeWalletSigner(address signer) returns()
func (_IConfig *IConfigTransactorSession) InitializeWalletSigner(signer common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.InitializeWalletSigner(&_IConfig.TransactOpts, signer)
}

// RemoveRecoverySigner is a paid mutator transaction binding the contract method 0x562d2e27.
//
// Solidity: function removeRecoverySigner(address signer) returns()
func (_IConfig *IConfigTransactor) RemoveRecoverySigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _IConfig.contract.Transact(opts, "removeRecoverySigner", signer)
}

// RemoveRecoverySigner is a paid mutator transaction binding the contract method 0x562d2e27.
//
// Solidity: function removeRecoverySigner(address signer) returns()
func (_IConfig *IConfigSession) RemoveRecoverySigner(signer common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.RemoveRecoverySigner(&_IConfig.TransactOpts, signer)
}

// RemoveRecoverySigner is a paid mutator transaction binding the contract method 0x562d2e27.
//
// Solidity: function removeRecoverySigner(address signer) returns()
func (_IConfig *IConfigTransactorSession) RemoveRecoverySigner(signer common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.RemoveRecoverySigner(&_IConfig.TransactOpts, signer)
}

// RemoveWhitelistedBundlers is a paid mutator transaction binding the contract method 0x16b97e9a.
//
// Solidity: function removeWhitelistedBundlers(address[] bundlers) returns()
func (_IConfig *IConfigTransactor) RemoveWhitelistedBundlers(opts *bind.TransactOpts, bundlers []common.Address) (*types.Transaction, error) {
	return _IConfig.contract.Transact(opts, "removeWhitelistedBundlers", bundlers)
}

// RemoveWhitelistedBundlers is a paid mutator transaction binding the contract method 0x16b97e9a.
//
// Solidity: function removeWhitelistedBundlers(address[] bundlers) returns()
func (_IConfig *IConfigSession) RemoveWhitelistedBundlers(bundlers []common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.RemoveWhitelistedBundlers(&_IConfig.TransactOpts, bundlers)
}

// RemoveWhitelistedBundlers is a paid mutator transaction binding the contract method 0x16b97e9a.
//
// Solidity: function removeWhitelistedBundlers(address[] bundlers) returns()
func (_IConfig *IConfigTransactorSession) RemoveWhitelistedBundlers(bundlers []common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.RemoveWhitelistedBundlers(&_IConfig.TransactOpts, bundlers)
}

// RemoveWhitelistedModules is a paid mutator transaction binding the contract method 0xb8076742.
//
// Solidity: function removeWhitelistedModules(address[] modules) returns()
func (_IConfig *IConfigTransactor) RemoveWhitelistedModules(opts *bind.TransactOpts, modules []common.Address) (*types.Transaction, error) {
	return _IConfig.contract.Transact(opts, "removeWhitelistedModules", modules)
}

// RemoveWhitelistedModules is a paid mutator transaction binding the contract method 0xb8076742.
//
// Solidity: function removeWhitelistedModules(address[] modules) returns()
func (_IConfig *IConfigSession) RemoveWhitelistedModules(modules []common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.RemoveWhitelistedModules(&_IConfig.TransactOpts, modules)
}

// RemoveWhitelistedModules is a paid mutator transaction binding the contract method 0xb8076742.
//
// Solidity: function removeWhitelistedModules(address[] modules) returns()
func (_IConfig *IConfigTransactorSession) RemoveWhitelistedModules(modules []common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.RemoveWhitelistedModules(&_IConfig.TransactOpts, modules)
}

// RemoveWhitelistedVerifier is a paid mutator transaction binding the contract method 0xc22305cf.
//
// Solidity: function removeWhitelistedVerifier(uint8 verifier) returns()
func (_IConfig *IConfigTransactor) RemoveWhitelistedVerifier(opts *bind.TransactOpts, verifier uint8) (*types.Transaction, error) {
	return _IConfig.contract.Transact(opts, "removeWhitelistedVerifier", verifier)
}

// RemoveWhitelistedVerifier is a paid mutator transaction binding the contract method 0xc22305cf.
//
// Solidity: function removeWhitelistedVerifier(uint8 verifier) returns()
func (_IConfig *IConfigSession) RemoveWhitelistedVerifier(verifier uint8) (*types.Transaction, error) {
	return _IConfig.Contract.RemoveWhitelistedVerifier(&_IConfig.TransactOpts, verifier)
}

// RemoveWhitelistedVerifier is a paid mutator transaction binding the contract method 0xc22305cf.
//
// Solidity: function removeWhitelistedVerifier(uint8 verifier) returns()
func (_IConfig *IConfigTransactorSession) RemoveWhitelistedVerifier(verifier uint8) (*types.Transaction, error) {
	return _IConfig.Contract.RemoveWhitelistedVerifier(&_IConfig.TransactOpts, verifier)
}

// IConfigFactorySignerUpdatedIterator is returned from FilterFactorySignerUpdated and is used to iterate over the raw logs and unpacked data for FactorySignerUpdated events raised by the IConfig contract.
type IConfigFactorySignerUpdatedIterator struct {
	Event *IConfigFactorySignerUpdated // Event containing the contract specifics and raw log

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
func (it *IConfigFactorySignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfigFactorySignerUpdated)
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
		it.Event = new(IConfigFactorySignerUpdated)
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
func (it *IConfigFactorySignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfigFactorySignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfigFactorySignerUpdated represents a FactorySignerUpdated event raised by the IConfig contract.
type IConfigFactorySignerUpdated struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFactorySignerUpdated is a free log retrieval operation binding the contract event 0x789d9446447ccefd52005113ddeccb6cfdcafd57a9f5229050f0b6c3dcd10f79.
//
// Solidity: event FactorySignerUpdated(address signer)
func (_IConfig *IConfigFilterer) FilterFactorySignerUpdated(opts *bind.FilterOpts) (*IConfigFactorySignerUpdatedIterator, error) {

	logs, sub, err := _IConfig.contract.FilterLogs(opts, "FactorySignerUpdated")
	if err != nil {
		return nil, err
	}
	return &IConfigFactorySignerUpdatedIterator{contract: _IConfig.contract, event: "FactorySignerUpdated", logs: logs, sub: sub}, nil
}

// WatchFactorySignerUpdated is a free log subscription operation binding the contract event 0x789d9446447ccefd52005113ddeccb6cfdcafd57a9f5229050f0b6c3dcd10f79.
//
// Solidity: event FactorySignerUpdated(address signer)
func (_IConfig *IConfigFilterer) WatchFactorySignerUpdated(opts *bind.WatchOpts, sink chan<- *IConfigFactorySignerUpdated) (event.Subscription, error) {

	logs, sub, err := _IConfig.contract.WatchLogs(opts, "FactorySignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfigFactorySignerUpdated)
				if err := _IConfig.contract.UnpackLog(event, "FactorySignerUpdated", log); err != nil {
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

// ParseFactorySignerUpdated is a log parse operation binding the contract event 0x789d9446447ccefd52005113ddeccb6cfdcafd57a9f5229050f0b6c3dcd10f79.
//
// Solidity: event FactorySignerUpdated(address signer)
func (_IConfig *IConfigFilterer) ParseFactorySignerUpdated(log types.Log) (*IConfigFactorySignerUpdated, error) {
	event := new(IConfigFactorySignerUpdated)
	if err := _IConfig.contract.UnpackLog(event, "FactorySignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IConfigPaySignerUpdatedIterator is returned from FilterPaySignerUpdated and is used to iterate over the raw logs and unpacked data for PaySignerUpdated events raised by the IConfig contract.
type IConfigPaySignerUpdatedIterator struct {
	Event *IConfigPaySignerUpdated // Event containing the contract specifics and raw log

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
func (it *IConfigPaySignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfigPaySignerUpdated)
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
		it.Event = new(IConfigPaySignerUpdated)
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
func (it *IConfigPaySignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfigPaySignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfigPaySignerUpdated represents a PaySignerUpdated event raised by the IConfig contract.
type IConfigPaySignerUpdated struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaySignerUpdated is a free log retrieval operation binding the contract event 0x322933b0e9c97ef2d23835a0a1835c057a9601a32b8227690193a42c57bacf3d.
//
// Solidity: event PaySignerUpdated(address signer)
func (_IConfig *IConfigFilterer) FilterPaySignerUpdated(opts *bind.FilterOpts) (*IConfigPaySignerUpdatedIterator, error) {

	logs, sub, err := _IConfig.contract.FilterLogs(opts, "PaySignerUpdated")
	if err != nil {
		return nil, err
	}
	return &IConfigPaySignerUpdatedIterator{contract: _IConfig.contract, event: "PaySignerUpdated", logs: logs, sub: sub}, nil
}

// WatchPaySignerUpdated is a free log subscription operation binding the contract event 0x322933b0e9c97ef2d23835a0a1835c057a9601a32b8227690193a42c57bacf3d.
//
// Solidity: event PaySignerUpdated(address signer)
func (_IConfig *IConfigFilterer) WatchPaySignerUpdated(opts *bind.WatchOpts, sink chan<- *IConfigPaySignerUpdated) (event.Subscription, error) {

	logs, sub, err := _IConfig.contract.WatchLogs(opts, "PaySignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfigPaySignerUpdated)
				if err := _IConfig.contract.UnpackLog(event, "PaySignerUpdated", log); err != nil {
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

// ParsePaySignerUpdated is a log parse operation binding the contract event 0x322933b0e9c97ef2d23835a0a1835c057a9601a32b8227690193a42c57bacf3d.
//
// Solidity: event PaySignerUpdated(address signer)
func (_IConfig *IConfigFilterer) ParsePaySignerUpdated(log types.Log) (*IConfigPaySignerUpdated, error) {
	event := new(IConfigPaySignerUpdated)
	if err := _IConfig.contract.UnpackLog(event, "PaySignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IConfigRecoverySignerAddedIterator is returned from FilterRecoverySignerAdded and is used to iterate over the raw logs and unpacked data for RecoverySignerAdded events raised by the IConfig contract.
type IConfigRecoverySignerAddedIterator struct {
	Event *IConfigRecoverySignerAdded // Event containing the contract specifics and raw log

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
func (it *IConfigRecoverySignerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfigRecoverySignerAdded)
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
		it.Event = new(IConfigRecoverySignerAdded)
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
func (it *IConfigRecoverySignerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfigRecoverySignerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfigRecoverySignerAdded represents a RecoverySignerAdded event raised by the IConfig contract.
type IConfigRecoverySignerAdded struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverySignerAdded is a free log retrieval operation binding the contract event 0x22dbed02a51e94f91fc3a6d1d2014ed6bc570834e7f3536162d626c429d90e54.
//
// Solidity: event RecoverySignerAdded(address signer)
func (_IConfig *IConfigFilterer) FilterRecoverySignerAdded(opts *bind.FilterOpts) (*IConfigRecoverySignerAddedIterator, error) {

	logs, sub, err := _IConfig.contract.FilterLogs(opts, "RecoverySignerAdded")
	if err != nil {
		return nil, err
	}
	return &IConfigRecoverySignerAddedIterator{contract: _IConfig.contract, event: "RecoverySignerAdded", logs: logs, sub: sub}, nil
}

// WatchRecoverySignerAdded is a free log subscription operation binding the contract event 0x22dbed02a51e94f91fc3a6d1d2014ed6bc570834e7f3536162d626c429d90e54.
//
// Solidity: event RecoverySignerAdded(address signer)
func (_IConfig *IConfigFilterer) WatchRecoverySignerAdded(opts *bind.WatchOpts, sink chan<- *IConfigRecoverySignerAdded) (event.Subscription, error) {

	logs, sub, err := _IConfig.contract.WatchLogs(opts, "RecoverySignerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfigRecoverySignerAdded)
				if err := _IConfig.contract.UnpackLog(event, "RecoverySignerAdded", log); err != nil {
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

// ParseRecoverySignerAdded is a log parse operation binding the contract event 0x22dbed02a51e94f91fc3a6d1d2014ed6bc570834e7f3536162d626c429d90e54.
//
// Solidity: event RecoverySignerAdded(address signer)
func (_IConfig *IConfigFilterer) ParseRecoverySignerAdded(log types.Log) (*IConfigRecoverySignerAdded, error) {
	event := new(IConfigRecoverySignerAdded)
	if err := _IConfig.contract.UnpackLog(event, "RecoverySignerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IConfigRecoverySignerRemovedIterator is returned from FilterRecoverySignerRemoved and is used to iterate over the raw logs and unpacked data for RecoverySignerRemoved events raised by the IConfig contract.
type IConfigRecoverySignerRemovedIterator struct {
	Event *IConfigRecoverySignerRemoved // Event containing the contract specifics and raw log

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
func (it *IConfigRecoverySignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfigRecoverySignerRemoved)
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
		it.Event = new(IConfigRecoverySignerRemoved)
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
func (it *IConfigRecoverySignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfigRecoverySignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfigRecoverySignerRemoved represents a RecoverySignerRemoved event raised by the IConfig contract.
type IConfigRecoverySignerRemoved struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverySignerRemoved is a free log retrieval operation binding the contract event 0xadd78f873e935721c08cc4e66ec8de5f0bd111d6a9966eaa93b5e22b417d23a4.
//
// Solidity: event RecoverySignerRemoved(address signer)
func (_IConfig *IConfigFilterer) FilterRecoverySignerRemoved(opts *bind.FilterOpts) (*IConfigRecoverySignerRemovedIterator, error) {

	logs, sub, err := _IConfig.contract.FilterLogs(opts, "RecoverySignerRemoved")
	if err != nil {
		return nil, err
	}
	return &IConfigRecoverySignerRemovedIterator{contract: _IConfig.contract, event: "RecoverySignerRemoved", logs: logs, sub: sub}, nil
}

// WatchRecoverySignerRemoved is a free log subscription operation binding the contract event 0xadd78f873e935721c08cc4e66ec8de5f0bd111d6a9966eaa93b5e22b417d23a4.
//
// Solidity: event RecoverySignerRemoved(address signer)
func (_IConfig *IConfigFilterer) WatchRecoverySignerRemoved(opts *bind.WatchOpts, sink chan<- *IConfigRecoverySignerRemoved) (event.Subscription, error) {

	logs, sub, err := _IConfig.contract.WatchLogs(opts, "RecoverySignerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfigRecoverySignerRemoved)
				if err := _IConfig.contract.UnpackLog(event, "RecoverySignerRemoved", log); err != nil {
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

// ParseRecoverySignerRemoved is a log parse operation binding the contract event 0xadd78f873e935721c08cc4e66ec8de5f0bd111d6a9966eaa93b5e22b417d23a4.
//
// Solidity: event RecoverySignerRemoved(address signer)
func (_IConfig *IConfigFilterer) ParseRecoverySignerRemoved(log types.Log) (*IConfigRecoverySignerRemoved, error) {
	event := new(IConfigRecoverySignerRemoved)
	if err := _IConfig.contract.UnpackLog(event, "RecoverySignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IConfigSetSafeSingletonIterator is returned from FilterSetSafeSingleton and is used to iterate over the raw logs and unpacked data for SetSafeSingleton events raised by the IConfig contract.
type IConfigSetSafeSingletonIterator struct {
	Event *IConfigSetSafeSingleton // Event containing the contract specifics and raw log

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
func (it *IConfigSetSafeSingletonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfigSetSafeSingleton)
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
		it.Event = new(IConfigSetSafeSingleton)
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
func (it *IConfigSetSafeSingletonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfigSetSafeSingletonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfigSetSafeSingleton represents a SetSafeSingleton event raised by the IConfig contract.
type IConfigSetSafeSingleton struct {
	Singleton common.Address
	Status    bool
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetSafeSingleton is a free log retrieval operation binding the contract event 0xf80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f80406191.
//
// Solidity: event SetSafeSingleton(address singleton, bool status, uint256 timestamp)
func (_IConfig *IConfigFilterer) FilterSetSafeSingleton(opts *bind.FilterOpts) (*IConfigSetSafeSingletonIterator, error) {

	logs, sub, err := _IConfig.contract.FilterLogs(opts, "SetSafeSingleton")
	if err != nil {
		return nil, err
	}
	return &IConfigSetSafeSingletonIterator{contract: _IConfig.contract, event: "SetSafeSingleton", logs: logs, sub: sub}, nil
}

// WatchSetSafeSingleton is a free log subscription operation binding the contract event 0xf80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f80406191.
//
// Solidity: event SetSafeSingleton(address singleton, bool status, uint256 timestamp)
func (_IConfig *IConfigFilterer) WatchSetSafeSingleton(opts *bind.WatchOpts, sink chan<- *IConfigSetSafeSingleton) (event.Subscription, error) {

	logs, sub, err := _IConfig.contract.WatchLogs(opts, "SetSafeSingleton")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfigSetSafeSingleton)
				if err := _IConfig.contract.UnpackLog(event, "SetSafeSingleton", log); err != nil {
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

// ParseSetSafeSingleton is a log parse operation binding the contract event 0xf80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f80406191.
//
// Solidity: event SetSafeSingleton(address singleton, bool status, uint256 timestamp)
func (_IConfig *IConfigFilterer) ParseSetSafeSingleton(log types.Log) (*IConfigSetSafeSingleton, error) {
	event := new(IConfigSetSafeSingleton)
	if err := _IConfig.contract.UnpackLog(event, "SetSafeSingleton", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IConfigSetSenderSignerIterator is returned from FilterSetSenderSigner and is used to iterate over the raw logs and unpacked data for SetSenderSigner events raised by the IConfig contract.
type IConfigSetSenderSignerIterator struct {
	Event *IConfigSetSenderSigner // Event containing the contract specifics and raw log

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
func (it *IConfigSetSenderSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfigSetSenderSigner)
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
		it.Event = new(IConfigSetSenderSigner)
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
func (it *IConfigSetSenderSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfigSetSenderSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfigSetSenderSigner represents a SetSenderSigner event raised by the IConfig contract.
type IConfigSetSenderSigner struct {
	Sender    common.Address
	Signer    common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetSenderSigner is a free log retrieval operation binding the contract event 0x03b72cb2fffc620d2a9c4f6bce0636faa82c948488a45f2f4b0684bb17cfdc72.
//
// Solidity: event SetSenderSigner(address sender, address signer, uint256 timestamp)
func (_IConfig *IConfigFilterer) FilterSetSenderSigner(opts *bind.FilterOpts) (*IConfigSetSenderSignerIterator, error) {

	logs, sub, err := _IConfig.contract.FilterLogs(opts, "SetSenderSigner")
	if err != nil {
		return nil, err
	}
	return &IConfigSetSenderSignerIterator{contract: _IConfig.contract, event: "SetSenderSigner", logs: logs, sub: sub}, nil
}

// WatchSetSenderSigner is a free log subscription operation binding the contract event 0x03b72cb2fffc620d2a9c4f6bce0636faa82c948488a45f2f4b0684bb17cfdc72.
//
// Solidity: event SetSenderSigner(address sender, address signer, uint256 timestamp)
func (_IConfig *IConfigFilterer) WatchSetSenderSigner(opts *bind.WatchOpts, sink chan<- *IConfigSetSenderSigner) (event.Subscription, error) {

	logs, sub, err := _IConfig.contract.WatchLogs(opts, "SetSenderSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfigSetSenderSigner)
				if err := _IConfig.contract.UnpackLog(event, "SetSenderSigner", log); err != nil {
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

// ParseSetSenderSigner is a log parse operation binding the contract event 0x03b72cb2fffc620d2a9c4f6bce0636faa82c948488a45f2f4b0684bb17cfdc72.
//
// Solidity: event SetSenderSigner(address sender, address signer, uint256 timestamp)
func (_IConfig *IConfigFilterer) ParseSetSenderSigner(log types.Log) (*IConfigSetSenderSigner, error) {
	event := new(IConfigSetSenderSigner)
	if err := _IConfig.contract.UnpackLog(event, "SetSenderSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IConfigVerifierTypeAddedIterator is returned from FilterVerifierTypeAdded and is used to iterate over the raw logs and unpacked data for VerifierTypeAdded events raised by the IConfig contract.
type IConfigVerifierTypeAddedIterator struct {
	Event *IConfigVerifierTypeAdded // Event containing the contract specifics and raw log

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
func (it *IConfigVerifierTypeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfigVerifierTypeAdded)
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
		it.Event = new(IConfigVerifierTypeAdded)
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
func (it *IConfigVerifierTypeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfigVerifierTypeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfigVerifierTypeAdded represents a VerifierTypeAdded event raised by the IConfig contract.
type IConfigVerifierTypeAdded struct {
	Verifier uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierTypeAdded is a free log retrieval operation binding the contract event 0xc007db510c05b52ed8aa21436b678ffa49a3f5de71cb7ba2c81363eda6f9a734.
//
// Solidity: event VerifierTypeAdded(uint8 verifier)
func (_IConfig *IConfigFilterer) FilterVerifierTypeAdded(opts *bind.FilterOpts) (*IConfigVerifierTypeAddedIterator, error) {

	logs, sub, err := _IConfig.contract.FilterLogs(opts, "VerifierTypeAdded")
	if err != nil {
		return nil, err
	}
	return &IConfigVerifierTypeAddedIterator{contract: _IConfig.contract, event: "VerifierTypeAdded", logs: logs, sub: sub}, nil
}

// WatchVerifierTypeAdded is a free log subscription operation binding the contract event 0xc007db510c05b52ed8aa21436b678ffa49a3f5de71cb7ba2c81363eda6f9a734.
//
// Solidity: event VerifierTypeAdded(uint8 verifier)
func (_IConfig *IConfigFilterer) WatchVerifierTypeAdded(opts *bind.WatchOpts, sink chan<- *IConfigVerifierTypeAdded) (event.Subscription, error) {

	logs, sub, err := _IConfig.contract.WatchLogs(opts, "VerifierTypeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfigVerifierTypeAdded)
				if err := _IConfig.contract.UnpackLog(event, "VerifierTypeAdded", log); err != nil {
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

// ParseVerifierTypeAdded is a log parse operation binding the contract event 0xc007db510c05b52ed8aa21436b678ffa49a3f5de71cb7ba2c81363eda6f9a734.
//
// Solidity: event VerifierTypeAdded(uint8 verifier)
func (_IConfig *IConfigFilterer) ParseVerifierTypeAdded(log types.Log) (*IConfigVerifierTypeAdded, error) {
	event := new(IConfigVerifierTypeAdded)
	if err := _IConfig.contract.UnpackLog(event, "VerifierTypeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IConfigVerifierTypeRemovedIterator is returned from FilterVerifierTypeRemoved and is used to iterate over the raw logs and unpacked data for VerifierTypeRemoved events raised by the IConfig contract.
type IConfigVerifierTypeRemovedIterator struct {
	Event *IConfigVerifierTypeRemoved // Event containing the contract specifics and raw log

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
func (it *IConfigVerifierTypeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfigVerifierTypeRemoved)
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
		it.Event = new(IConfigVerifierTypeRemoved)
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
func (it *IConfigVerifierTypeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfigVerifierTypeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfigVerifierTypeRemoved represents a VerifierTypeRemoved event raised by the IConfig contract.
type IConfigVerifierTypeRemoved struct {
	Verifier uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierTypeRemoved is a free log retrieval operation binding the contract event 0xf7ad1892e8c359fbe99be323cb12129c63748ec6c859ad3dd7f22f427742ebab.
//
// Solidity: event VerifierTypeRemoved(uint8 verifier)
func (_IConfig *IConfigFilterer) FilterVerifierTypeRemoved(opts *bind.FilterOpts) (*IConfigVerifierTypeRemovedIterator, error) {

	logs, sub, err := _IConfig.contract.FilterLogs(opts, "VerifierTypeRemoved")
	if err != nil {
		return nil, err
	}
	return &IConfigVerifierTypeRemovedIterator{contract: _IConfig.contract, event: "VerifierTypeRemoved", logs: logs, sub: sub}, nil
}

// WatchVerifierTypeRemoved is a free log subscription operation binding the contract event 0xf7ad1892e8c359fbe99be323cb12129c63748ec6c859ad3dd7f22f427742ebab.
//
// Solidity: event VerifierTypeRemoved(uint8 verifier)
func (_IConfig *IConfigFilterer) WatchVerifierTypeRemoved(opts *bind.WatchOpts, sink chan<- *IConfigVerifierTypeRemoved) (event.Subscription, error) {

	logs, sub, err := _IConfig.contract.WatchLogs(opts, "VerifierTypeRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfigVerifierTypeRemoved)
				if err := _IConfig.contract.UnpackLog(event, "VerifierTypeRemoved", log); err != nil {
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

// ParseVerifierTypeRemoved is a log parse operation binding the contract event 0xf7ad1892e8c359fbe99be323cb12129c63748ec6c859ad3dd7f22f427742ebab.
//
// Solidity: event VerifierTypeRemoved(uint8 verifier)
func (_IConfig *IConfigFilterer) ParseVerifierTypeRemoved(log types.Log) (*IConfigVerifierTypeRemoved, error) {
	event := new(IConfigVerifierTypeRemoved)
	if err := _IConfig.contract.UnpackLog(event, "VerifierTypeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IConfigWhitelistBundlerAddedIterator is returned from FilterWhitelistBundlerAdded and is used to iterate over the raw logs and unpacked data for WhitelistBundlerAdded events raised by the IConfig contract.
type IConfigWhitelistBundlerAddedIterator struct {
	Event *IConfigWhitelistBundlerAdded // Event containing the contract specifics and raw log

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
func (it *IConfigWhitelistBundlerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfigWhitelistBundlerAdded)
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
		it.Event = new(IConfigWhitelistBundlerAdded)
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
func (it *IConfigWhitelistBundlerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfigWhitelistBundlerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfigWhitelistBundlerAdded represents a WhitelistBundlerAdded event raised by the IConfig contract.
type IConfigWhitelistBundlerAdded struct {
	Bundlers []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistBundlerAdded is a free log retrieval operation binding the contract event 0x49bd286efd7e26d327db1b6f433560c5f6156363198289a65a6ca9904338f4f4.
//
// Solidity: event WhitelistBundlerAdded(address[] bundlers)
func (_IConfig *IConfigFilterer) FilterWhitelistBundlerAdded(opts *bind.FilterOpts) (*IConfigWhitelistBundlerAddedIterator, error) {

	logs, sub, err := _IConfig.contract.FilterLogs(opts, "WhitelistBundlerAdded")
	if err != nil {
		return nil, err
	}
	return &IConfigWhitelistBundlerAddedIterator{contract: _IConfig.contract, event: "WhitelistBundlerAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistBundlerAdded is a free log subscription operation binding the contract event 0x49bd286efd7e26d327db1b6f433560c5f6156363198289a65a6ca9904338f4f4.
//
// Solidity: event WhitelistBundlerAdded(address[] bundlers)
func (_IConfig *IConfigFilterer) WatchWhitelistBundlerAdded(opts *bind.WatchOpts, sink chan<- *IConfigWhitelistBundlerAdded) (event.Subscription, error) {

	logs, sub, err := _IConfig.contract.WatchLogs(opts, "WhitelistBundlerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfigWhitelistBundlerAdded)
				if err := _IConfig.contract.UnpackLog(event, "WhitelistBundlerAdded", log); err != nil {
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

// ParseWhitelistBundlerAdded is a log parse operation binding the contract event 0x49bd286efd7e26d327db1b6f433560c5f6156363198289a65a6ca9904338f4f4.
//
// Solidity: event WhitelistBundlerAdded(address[] bundlers)
func (_IConfig *IConfigFilterer) ParseWhitelistBundlerAdded(log types.Log) (*IConfigWhitelistBundlerAdded, error) {
	event := new(IConfigWhitelistBundlerAdded)
	if err := _IConfig.contract.UnpackLog(event, "WhitelistBundlerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IConfigWhitelistBundlerRemovedIterator is returned from FilterWhitelistBundlerRemoved and is used to iterate over the raw logs and unpacked data for WhitelistBundlerRemoved events raised by the IConfig contract.
type IConfigWhitelistBundlerRemovedIterator struct {
	Event *IConfigWhitelistBundlerRemoved // Event containing the contract specifics and raw log

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
func (it *IConfigWhitelistBundlerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfigWhitelistBundlerRemoved)
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
		it.Event = new(IConfigWhitelistBundlerRemoved)
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
func (it *IConfigWhitelistBundlerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfigWhitelistBundlerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfigWhitelistBundlerRemoved represents a WhitelistBundlerRemoved event raised by the IConfig contract.
type IConfigWhitelistBundlerRemoved struct {
	Bundlers []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistBundlerRemoved is a free log retrieval operation binding the contract event 0x623f7acab9936b83f17999b4d96bfda34b1d3a749bd474d39d0413cc977044ed.
//
// Solidity: event WhitelistBundlerRemoved(address[] bundlers)
func (_IConfig *IConfigFilterer) FilterWhitelistBundlerRemoved(opts *bind.FilterOpts) (*IConfigWhitelistBundlerRemovedIterator, error) {

	logs, sub, err := _IConfig.contract.FilterLogs(opts, "WhitelistBundlerRemoved")
	if err != nil {
		return nil, err
	}
	return &IConfigWhitelistBundlerRemovedIterator{contract: _IConfig.contract, event: "WhitelistBundlerRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistBundlerRemoved is a free log subscription operation binding the contract event 0x623f7acab9936b83f17999b4d96bfda34b1d3a749bd474d39d0413cc977044ed.
//
// Solidity: event WhitelistBundlerRemoved(address[] bundlers)
func (_IConfig *IConfigFilterer) WatchWhitelistBundlerRemoved(opts *bind.WatchOpts, sink chan<- *IConfigWhitelistBundlerRemoved) (event.Subscription, error) {

	logs, sub, err := _IConfig.contract.WatchLogs(opts, "WhitelistBundlerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfigWhitelistBundlerRemoved)
				if err := _IConfig.contract.UnpackLog(event, "WhitelistBundlerRemoved", log); err != nil {
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

// ParseWhitelistBundlerRemoved is a log parse operation binding the contract event 0x623f7acab9936b83f17999b4d96bfda34b1d3a749bd474d39d0413cc977044ed.
//
// Solidity: event WhitelistBundlerRemoved(address[] bundlers)
func (_IConfig *IConfigFilterer) ParseWhitelistBundlerRemoved(log types.Log) (*IConfigWhitelistBundlerRemoved, error) {
	event := new(IConfigWhitelistBundlerRemoved)
	if err := _IConfig.contract.UnpackLog(event, "WhitelistBundlerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IConfigWhitelistModuleAddedIterator is returned from FilterWhitelistModuleAdded and is used to iterate over the raw logs and unpacked data for WhitelistModuleAdded events raised by the IConfig contract.
type IConfigWhitelistModuleAddedIterator struct {
	Event *IConfigWhitelistModuleAdded // Event containing the contract specifics and raw log

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
func (it *IConfigWhitelistModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfigWhitelistModuleAdded)
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
		it.Event = new(IConfigWhitelistModuleAdded)
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
func (it *IConfigWhitelistModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfigWhitelistModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfigWhitelistModuleAdded represents a WhitelistModuleAdded event raised by the IConfig contract.
type IConfigWhitelistModuleAdded struct {
	Module     common.Address
	ModuleType *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWhitelistModuleAdded is a free log retrieval operation binding the contract event 0x51a45eb94bbc0121a4cdcf5f28257162a61221f0d6f8897adffe076012457865.
//
// Solidity: event WhitelistModuleAdded(address module, uint256 moduleType)
func (_IConfig *IConfigFilterer) FilterWhitelistModuleAdded(opts *bind.FilterOpts) (*IConfigWhitelistModuleAddedIterator, error) {

	logs, sub, err := _IConfig.contract.FilterLogs(opts, "WhitelistModuleAdded")
	if err != nil {
		return nil, err
	}
	return &IConfigWhitelistModuleAddedIterator{contract: _IConfig.contract, event: "WhitelistModuleAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistModuleAdded is a free log subscription operation binding the contract event 0x51a45eb94bbc0121a4cdcf5f28257162a61221f0d6f8897adffe076012457865.
//
// Solidity: event WhitelistModuleAdded(address module, uint256 moduleType)
func (_IConfig *IConfigFilterer) WatchWhitelistModuleAdded(opts *bind.WatchOpts, sink chan<- *IConfigWhitelistModuleAdded) (event.Subscription, error) {

	logs, sub, err := _IConfig.contract.WatchLogs(opts, "WhitelistModuleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfigWhitelistModuleAdded)
				if err := _IConfig.contract.UnpackLog(event, "WhitelistModuleAdded", log); err != nil {
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

// ParseWhitelistModuleAdded is a log parse operation binding the contract event 0x51a45eb94bbc0121a4cdcf5f28257162a61221f0d6f8897adffe076012457865.
//
// Solidity: event WhitelistModuleAdded(address module, uint256 moduleType)
func (_IConfig *IConfigFilterer) ParseWhitelistModuleAdded(log types.Log) (*IConfigWhitelistModuleAdded, error) {
	event := new(IConfigWhitelistModuleAdded)
	if err := _IConfig.contract.UnpackLog(event, "WhitelistModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IConfigWhitelistModuleRemovedIterator is returned from FilterWhitelistModuleRemoved and is used to iterate over the raw logs and unpacked data for WhitelistModuleRemoved events raised by the IConfig contract.
type IConfigWhitelistModuleRemovedIterator struct {
	Event *IConfigWhitelistModuleRemoved // Event containing the contract specifics and raw log

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
func (it *IConfigWhitelistModuleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfigWhitelistModuleRemoved)
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
		it.Event = new(IConfigWhitelistModuleRemoved)
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
func (it *IConfigWhitelistModuleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfigWhitelistModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfigWhitelistModuleRemoved represents a WhitelistModuleRemoved event raised by the IConfig contract.
type IConfigWhitelistModuleRemoved struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWhitelistModuleRemoved is a free log retrieval operation binding the contract event 0x312c33106e0d74000c54d884caee742f0a33f8143ed96561cc1c80d3351b2cd2.
//
// Solidity: event WhitelistModuleRemoved(address module)
func (_IConfig *IConfigFilterer) FilterWhitelistModuleRemoved(opts *bind.FilterOpts) (*IConfigWhitelistModuleRemovedIterator, error) {

	logs, sub, err := _IConfig.contract.FilterLogs(opts, "WhitelistModuleRemoved")
	if err != nil {
		return nil, err
	}
	return &IConfigWhitelistModuleRemovedIterator{contract: _IConfig.contract, event: "WhitelistModuleRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistModuleRemoved is a free log subscription operation binding the contract event 0x312c33106e0d74000c54d884caee742f0a33f8143ed96561cc1c80d3351b2cd2.
//
// Solidity: event WhitelistModuleRemoved(address module)
func (_IConfig *IConfigFilterer) WatchWhitelistModuleRemoved(opts *bind.WatchOpts, sink chan<- *IConfigWhitelistModuleRemoved) (event.Subscription, error) {

	logs, sub, err := _IConfig.contract.WatchLogs(opts, "WhitelistModuleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfigWhitelistModuleRemoved)
				if err := _IConfig.contract.UnpackLog(event, "WhitelistModuleRemoved", log); err != nil {
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

// ParseWhitelistModuleRemoved is a log parse operation binding the contract event 0x312c33106e0d74000c54d884caee742f0a33f8143ed96561cc1c80d3351b2cd2.
//
// Solidity: event WhitelistModuleRemoved(address module)
func (_IConfig *IConfigFilterer) ParseWhitelistModuleRemoved(log types.Log) (*IConfigWhitelistModuleRemoved, error) {
	event := new(IConfigWhitelistModuleRemoved)
	if err := _IConfig.contract.UnpackLog(event, "WhitelistModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
