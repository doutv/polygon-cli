// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package useroperationlib

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

// UserOperationLibMetaData contains all meta data concerning the UserOperationLib contract.
var UserOperationLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"PAYMASTER_DATA_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAYMASTER_POSTOP_GAS_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAYMASTER_VALIDATION_GAS_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523460175760bb9081601d823930815050f35b600080fdfe6080806040526004361015601257600080fd5b600090813560e01c90816325093e1b14606c57508063b29a8ff41460565763ede3150214603e57600080fd5b80600319360112605357602060405160348152f35b80fd5b5080600319360112605357602060405160148152f35b90508160031936011260815780602460209252f35b5080fdfea2646970667358221220bd5f31a1912c33b7e028dfc74fa5b16d1d2c478d281719bfebc0c431bf1cc13864736f6c63430008190033",
}

// UserOperationLibABI is the input ABI used to generate the binding from.
// Deprecated: Use UserOperationLibMetaData.ABI instead.
var UserOperationLibABI = UserOperationLibMetaData.ABI

// UserOperationLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UserOperationLibMetaData.Bin instead.
var UserOperationLibBin = UserOperationLibMetaData.Bin

// DeployUserOperationLib deploys a new Ethereum contract, binding an instance of UserOperationLib to it.
func DeployUserOperationLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserOperationLib, error) {
	parsed, err := UserOperationLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UserOperationLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserOperationLib{UserOperationLibCaller: UserOperationLibCaller{contract: contract}, UserOperationLibTransactor: UserOperationLibTransactor{contract: contract}, UserOperationLibFilterer: UserOperationLibFilterer{contract: contract}}, nil
}

// UserOperationLib is an auto generated Go binding around an Ethereum contract.
type UserOperationLib struct {
	UserOperationLibCaller     // Read-only binding to the contract
	UserOperationLibTransactor // Write-only binding to the contract
	UserOperationLibFilterer   // Log filterer for contract events
}

// UserOperationLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserOperationLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserOperationLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserOperationLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserOperationLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserOperationLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserOperationLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserOperationLibSession struct {
	Contract     *UserOperationLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserOperationLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserOperationLibCallerSession struct {
	Contract *UserOperationLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// UserOperationLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserOperationLibTransactorSession struct {
	Contract     *UserOperationLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UserOperationLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserOperationLibRaw struct {
	Contract *UserOperationLib // Generic contract binding to access the raw methods on
}

// UserOperationLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserOperationLibCallerRaw struct {
	Contract *UserOperationLibCaller // Generic read-only contract binding to access the raw methods on
}

// UserOperationLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserOperationLibTransactorRaw struct {
	Contract *UserOperationLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserOperationLib creates a new instance of UserOperationLib, bound to a specific deployed contract.
func NewUserOperationLib(address common.Address, backend bind.ContractBackend) (*UserOperationLib, error) {
	contract, err := bindUserOperationLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserOperationLib{UserOperationLibCaller: UserOperationLibCaller{contract: contract}, UserOperationLibTransactor: UserOperationLibTransactor{contract: contract}, UserOperationLibFilterer: UserOperationLibFilterer{contract: contract}}, nil
}

// NewUserOperationLibCaller creates a new read-only instance of UserOperationLib, bound to a specific deployed contract.
func NewUserOperationLibCaller(address common.Address, caller bind.ContractCaller) (*UserOperationLibCaller, error) {
	contract, err := bindUserOperationLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserOperationLibCaller{contract: contract}, nil
}

// NewUserOperationLibTransactor creates a new write-only instance of UserOperationLib, bound to a specific deployed contract.
func NewUserOperationLibTransactor(address common.Address, transactor bind.ContractTransactor) (*UserOperationLibTransactor, error) {
	contract, err := bindUserOperationLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserOperationLibTransactor{contract: contract}, nil
}

// NewUserOperationLibFilterer creates a new log filterer instance of UserOperationLib, bound to a specific deployed contract.
func NewUserOperationLibFilterer(address common.Address, filterer bind.ContractFilterer) (*UserOperationLibFilterer, error) {
	contract, err := bindUserOperationLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserOperationLibFilterer{contract: contract}, nil
}

// bindUserOperationLib binds a generic wrapper to an already deployed contract.
func bindUserOperationLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UserOperationLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserOperationLib *UserOperationLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserOperationLib.Contract.UserOperationLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserOperationLib *UserOperationLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserOperationLib.Contract.UserOperationLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserOperationLib *UserOperationLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserOperationLib.Contract.UserOperationLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserOperationLib *UserOperationLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserOperationLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserOperationLib *UserOperationLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserOperationLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserOperationLib *UserOperationLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserOperationLib.Contract.contract.Transact(opts, method, params...)
}

// PAYMASTERDATAOFFSET is a free data retrieval call binding the contract method 0xede31502.
//
// Solidity: function PAYMASTER_DATA_OFFSET() view returns(uint256)
func (_UserOperationLib *UserOperationLibCaller) PAYMASTERDATAOFFSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserOperationLib.contract.Call(opts, &out, "PAYMASTER_DATA_OFFSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAYMASTERDATAOFFSET is a free data retrieval call binding the contract method 0xede31502.
//
// Solidity: function PAYMASTER_DATA_OFFSET() view returns(uint256)
func (_UserOperationLib *UserOperationLibSession) PAYMASTERDATAOFFSET() (*big.Int, error) {
	return _UserOperationLib.Contract.PAYMASTERDATAOFFSET(&_UserOperationLib.CallOpts)
}

// PAYMASTERDATAOFFSET is a free data retrieval call binding the contract method 0xede31502.
//
// Solidity: function PAYMASTER_DATA_OFFSET() view returns(uint256)
func (_UserOperationLib *UserOperationLibCallerSession) PAYMASTERDATAOFFSET() (*big.Int, error) {
	return _UserOperationLib.Contract.PAYMASTERDATAOFFSET(&_UserOperationLib.CallOpts)
}

// PAYMASTERPOSTOPGASOFFSET is a free data retrieval call binding the contract method 0x25093e1b.
//
// Solidity: function PAYMASTER_POSTOP_GAS_OFFSET() view returns(uint256)
func (_UserOperationLib *UserOperationLibCaller) PAYMASTERPOSTOPGASOFFSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserOperationLib.contract.Call(opts, &out, "PAYMASTER_POSTOP_GAS_OFFSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAYMASTERPOSTOPGASOFFSET is a free data retrieval call binding the contract method 0x25093e1b.
//
// Solidity: function PAYMASTER_POSTOP_GAS_OFFSET() view returns(uint256)
func (_UserOperationLib *UserOperationLibSession) PAYMASTERPOSTOPGASOFFSET() (*big.Int, error) {
	return _UserOperationLib.Contract.PAYMASTERPOSTOPGASOFFSET(&_UserOperationLib.CallOpts)
}

// PAYMASTERPOSTOPGASOFFSET is a free data retrieval call binding the contract method 0x25093e1b.
//
// Solidity: function PAYMASTER_POSTOP_GAS_OFFSET() view returns(uint256)
func (_UserOperationLib *UserOperationLibCallerSession) PAYMASTERPOSTOPGASOFFSET() (*big.Int, error) {
	return _UserOperationLib.Contract.PAYMASTERPOSTOPGASOFFSET(&_UserOperationLib.CallOpts)
}

// PAYMASTERVALIDATIONGASOFFSET is a free data retrieval call binding the contract method 0xb29a8ff4.
//
// Solidity: function PAYMASTER_VALIDATION_GAS_OFFSET() view returns(uint256)
func (_UserOperationLib *UserOperationLibCaller) PAYMASTERVALIDATIONGASOFFSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserOperationLib.contract.Call(opts, &out, "PAYMASTER_VALIDATION_GAS_OFFSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAYMASTERVALIDATIONGASOFFSET is a free data retrieval call binding the contract method 0xb29a8ff4.
//
// Solidity: function PAYMASTER_VALIDATION_GAS_OFFSET() view returns(uint256)
func (_UserOperationLib *UserOperationLibSession) PAYMASTERVALIDATIONGASOFFSET() (*big.Int, error) {
	return _UserOperationLib.Contract.PAYMASTERVALIDATIONGASOFFSET(&_UserOperationLib.CallOpts)
}

// PAYMASTERVALIDATIONGASOFFSET is a free data retrieval call binding the contract method 0xb29a8ff4.
//
// Solidity: function PAYMASTER_VALIDATION_GAS_OFFSET() view returns(uint256)
func (_UserOperationLib *UserOperationLibCallerSession) PAYMASTERVALIDATIONGASOFFSET() (*big.Int, error) {
	return _UserOperationLib.Contract.PAYMASTERVALIDATIONGASOFFSET(&_UserOperationLib.CallOpts)
}
