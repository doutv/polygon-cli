// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ecdsavalidator

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

// ECDSAValidatorMetaData contains all meta data concerning the ECDSAValidator contract.
var ECDSAValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"}],\"name\":\"isModuleType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signatureHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"isValidSignatureWithData\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557610175908161001b8239f35b600080fdfe608080604052600436101561001357600080fd5b600090813560e01c90816354c015f314610075575080636d61fe70146100705780638a91b0e3146100705763ecd059611461004d57600080fd5b3461006d57602036600319011261006d5760206040516001600435148152f35b80fd5b61010c565b9050346100d55760603660031901126100d55767ffffffffffffffff916024358381116100d5576100aa9036906004016100d9565b505060443592831161006d57506100c760209236906004016100d9565b5050630b135d3f60e11b8152f35b5080fd5b9181601f840112156101075782359167ffffffffffffffff8311610107576020838186019501011161010757565b600080fd5b346101075760203660031901126101075760043567ffffffffffffffff81116101075761013d9036906004016100d9565b00fea264697066735822122033deb05e15b9dd624640689ff27296b076e0f8988395743eaf23b103cd3ba24164736f6c63430008190033",
}

// ECDSAValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAValidatorMetaData.ABI instead.
var ECDSAValidatorABI = ECDSAValidatorMetaData.ABI

// ECDSAValidatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAValidatorMetaData.Bin instead.
var ECDSAValidatorBin = ECDSAValidatorMetaData.Bin

// DeployECDSAValidator deploys a new Ethereum contract, binding an instance of ECDSAValidator to it.
func DeployECDSAValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSAValidator, error) {
	parsed, err := ECDSAValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSAValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSAValidator{ECDSAValidatorCaller: ECDSAValidatorCaller{contract: contract}, ECDSAValidatorTransactor: ECDSAValidatorTransactor{contract: contract}, ECDSAValidatorFilterer: ECDSAValidatorFilterer{contract: contract}}, nil
}

// ECDSAValidator is an auto generated Go binding around an Ethereum contract.
type ECDSAValidator struct {
	ECDSAValidatorCaller     // Read-only binding to the contract
	ECDSAValidatorTransactor // Write-only binding to the contract
	ECDSAValidatorFilterer   // Log filterer for contract events
}

// ECDSAValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSAValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSAValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSAValidatorSession struct {
	Contract     *ECDSAValidator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSAValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSAValidatorCallerSession struct {
	Contract *ECDSAValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ECDSAValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSAValidatorTransactorSession struct {
	Contract     *ECDSAValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ECDSAValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSAValidatorRaw struct {
	Contract *ECDSAValidator // Generic contract binding to access the raw methods on
}

// ECDSAValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSAValidatorCallerRaw struct {
	Contract *ECDSAValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ECDSAValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSAValidatorTransactorRaw struct {
	Contract *ECDSAValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSAValidator creates a new instance of ECDSAValidator, bound to a specific deployed contract.
func NewECDSAValidator(address common.Address, backend bind.ContractBackend) (*ECDSAValidator, error) {
	contract, err := bindECDSAValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSAValidator{ECDSAValidatorCaller: ECDSAValidatorCaller{contract: contract}, ECDSAValidatorTransactor: ECDSAValidatorTransactor{contract: contract}, ECDSAValidatorFilterer: ECDSAValidatorFilterer{contract: contract}}, nil
}

// NewECDSAValidatorCaller creates a new read-only instance of ECDSAValidator, bound to a specific deployed contract.
func NewECDSAValidatorCaller(address common.Address, caller bind.ContractCaller) (*ECDSAValidatorCaller, error) {
	contract, err := bindECDSAValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSAValidatorCaller{contract: contract}, nil
}

// NewECDSAValidatorTransactor creates a new write-only instance of ECDSAValidator, bound to a specific deployed contract.
func NewECDSAValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSAValidatorTransactor, error) {
	contract, err := bindECDSAValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSAValidatorTransactor{contract: contract}, nil
}

// NewECDSAValidatorFilterer creates a new log filterer instance of ECDSAValidator, bound to a specific deployed contract.
func NewECDSAValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAValidatorFilterer, error) {
	contract, err := bindECDSAValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAValidatorFilterer{contract: contract}, nil
}

// bindECDSAValidator binds a generic wrapper to an already deployed contract.
func bindECDSAValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ECDSAValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSAValidator *ECDSAValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSAValidator.Contract.ECDSAValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSAValidator *ECDSAValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSAValidator.Contract.ECDSAValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSAValidator *ECDSAValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSAValidator.Contract.ECDSAValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSAValidator *ECDSAValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSAValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSAValidator *ECDSAValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSAValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSAValidator *ECDSAValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSAValidator.Contract.contract.Transact(opts, method, params...)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_ECDSAValidator *ECDSAValidatorCaller) IsModuleType(opts *bind.CallOpts, moduleTypeId *big.Int) (bool, error) {
	var out []interface{}
	err := _ECDSAValidator.contract.Call(opts, &out, "isModuleType", moduleTypeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_ECDSAValidator *ECDSAValidatorSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _ECDSAValidator.Contract.IsModuleType(&_ECDSAValidator.CallOpts, moduleTypeId)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_ECDSAValidator *ECDSAValidatorCallerSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _ECDSAValidator.Contract.IsModuleType(&_ECDSAValidator.CallOpts, moduleTypeId)
}

// IsValidSignatureWithData is a free data retrieval call binding the contract method 0x54c015f3.
//
// Solidity: function isValidSignatureWithData(bytes32 signatureHash, bytes signature, bytes data) view returns(bytes4)
func (_ECDSAValidator *ECDSAValidatorCaller) IsValidSignatureWithData(opts *bind.CallOpts, signatureHash [32]byte, signature []byte, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _ECDSAValidator.contract.Call(opts, &out, "isValidSignatureWithData", signatureHash, signature, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignatureWithData is a free data retrieval call binding the contract method 0x54c015f3.
//
// Solidity: function isValidSignatureWithData(bytes32 signatureHash, bytes signature, bytes data) view returns(bytes4)
func (_ECDSAValidator *ECDSAValidatorSession) IsValidSignatureWithData(signatureHash [32]byte, signature []byte, data []byte) ([4]byte, error) {
	return _ECDSAValidator.Contract.IsValidSignatureWithData(&_ECDSAValidator.CallOpts, signatureHash, signature, data)
}

// IsValidSignatureWithData is a free data retrieval call binding the contract method 0x54c015f3.
//
// Solidity: function isValidSignatureWithData(bytes32 signatureHash, bytes signature, bytes data) view returns(bytes4)
func (_ECDSAValidator *ECDSAValidatorCallerSession) IsValidSignatureWithData(signatureHash [32]byte, signature []byte, data []byte) ([4]byte, error) {
	return _ECDSAValidator.Contract.IsValidSignatureWithData(&_ECDSAValidator.CallOpts, signatureHash, signature, data)
}

// OnInstall is a free data retrieval call binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) pure returns()
func (_ECDSAValidator *ECDSAValidatorCaller) OnInstall(opts *bind.CallOpts, data []byte) error {
	var out []interface{}
	err := _ECDSAValidator.contract.Call(opts, &out, "onInstall", data)

	if err != nil {
		return err
	}

	return err

}

// OnInstall is a free data retrieval call binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) pure returns()
func (_ECDSAValidator *ECDSAValidatorSession) OnInstall(data []byte) error {
	return _ECDSAValidator.Contract.OnInstall(&_ECDSAValidator.CallOpts, data)
}

// OnInstall is a free data retrieval call binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) pure returns()
func (_ECDSAValidator *ECDSAValidatorCallerSession) OnInstall(data []byte) error {
	return _ECDSAValidator.Contract.OnInstall(&_ECDSAValidator.CallOpts, data)
}

// OnUninstall is a free data retrieval call binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) pure returns()
func (_ECDSAValidator *ECDSAValidatorCaller) OnUninstall(opts *bind.CallOpts, data []byte) error {
	var out []interface{}
	err := _ECDSAValidator.contract.Call(opts, &out, "onUninstall", data)

	if err != nil {
		return err
	}

	return err

}

// OnUninstall is a free data retrieval call binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) pure returns()
func (_ECDSAValidator *ECDSAValidatorSession) OnUninstall(data []byte) error {
	return _ECDSAValidator.Contract.OnUninstall(&_ECDSAValidator.CallOpts, data)
}

// OnUninstall is a free data retrieval call binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) pure returns()
func (_ECDSAValidator *ECDSAValidatorCallerSession) OnUninstall(data []byte) error {
	return _ECDSAValidator.Contract.OnUninstall(&_ECDSAValidator.CallOpts, data)
}
