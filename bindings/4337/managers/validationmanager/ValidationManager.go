// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validationmanager

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

// ValidationManagerMetaData contains all meta data concerning the ValidationManager contract.
var ValidationManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MAXIMUM_KEYS_PER_ACCOUNT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"validFrom\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"validUntil\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523460135760de908160198239f35b600080fdfe6080806040526004361015601257600080fd5b600090813560e01c9081639bdafcb3146050575063d2c1934f14603457600080fd5b34604d5780600319360112604d57602060405160148152f35b80fd5b90503460a457602036600319011260a457604060609260043581526002602052206001808060a01b038254169101549082526fffffffffffffffffffffffffffffffff8116602083015260801c6040820152f35b5080fdfea264697066735822122026e99b63ac3d34da06e3a787fc1fbded143000009182da487541d167865f6a8c64736f6c63430008190033",
}

// ValidationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidationManagerMetaData.ABI instead.
var ValidationManagerABI = ValidationManagerMetaData.ABI

// ValidationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidationManagerMetaData.Bin instead.
var ValidationManagerBin = ValidationManagerMetaData.Bin

// DeployValidationManager deploys a new Ethereum contract, binding an instance of ValidationManager to it.
func DeployValidationManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidationManager, error) {
	parsed, err := ValidationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidationManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidationManager{ValidationManagerCaller: ValidationManagerCaller{contract: contract}, ValidationManagerTransactor: ValidationManagerTransactor{contract: contract}, ValidationManagerFilterer: ValidationManagerFilterer{contract: contract}}, nil
}

// ValidationManager is an auto generated Go binding around an Ethereum contract.
type ValidationManager struct {
	ValidationManagerCaller     // Read-only binding to the contract
	ValidationManagerTransactor // Write-only binding to the contract
	ValidationManagerFilterer   // Log filterer for contract events
}

// ValidationManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidationManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidationManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidationManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidationManagerSession struct {
	Contract     *ValidationManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidationManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidationManagerCallerSession struct {
	Contract *ValidationManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidationManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidationManagerTransactorSession struct {
	Contract     *ValidationManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidationManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidationManagerRaw struct {
	Contract *ValidationManager // Generic contract binding to access the raw methods on
}

// ValidationManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidationManagerCallerRaw struct {
	Contract *ValidationManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ValidationManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidationManagerTransactorRaw struct {
	Contract *ValidationManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidationManager creates a new instance of ValidationManager, bound to a specific deployed contract.
func NewValidationManager(address common.Address, backend bind.ContractBackend) (*ValidationManager, error) {
	contract, err := bindValidationManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidationManager{ValidationManagerCaller: ValidationManagerCaller{contract: contract}, ValidationManagerTransactor: ValidationManagerTransactor{contract: contract}, ValidationManagerFilterer: ValidationManagerFilterer{contract: contract}}, nil
}

// NewValidationManagerCaller creates a new read-only instance of ValidationManager, bound to a specific deployed contract.
func NewValidationManagerCaller(address common.Address, caller bind.ContractCaller) (*ValidationManagerCaller, error) {
	contract, err := bindValidationManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidationManagerCaller{contract: contract}, nil
}

// NewValidationManagerTransactor creates a new write-only instance of ValidationManager, bound to a specific deployed contract.
func NewValidationManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidationManagerTransactor, error) {
	contract, err := bindValidationManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidationManagerTransactor{contract: contract}, nil
}

// NewValidationManagerFilterer creates a new log filterer instance of ValidationManager, bound to a specific deployed contract.
func NewValidationManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidationManagerFilterer, error) {
	contract, err := bindValidationManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidationManagerFilterer{contract: contract}, nil
}

// bindValidationManager binds a generic wrapper to an already deployed contract.
func bindValidationManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidationManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidationManager *ValidationManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidationManager.Contract.ValidationManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidationManager *ValidationManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationManager.Contract.ValidationManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidationManager *ValidationManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidationManager.Contract.ValidationManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidationManager *ValidationManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidationManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidationManager *ValidationManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidationManager *ValidationManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidationManager.Contract.contract.Transact(opts, method, params...)
}

// MAXIMUMKEYSPERACCOUNT is a free data retrieval call binding the contract method 0xd2c1934f.
//
// Solidity: function MAXIMUM_KEYS_PER_ACCOUNT() view returns(uint8)
func (_ValidationManager *ValidationManagerCaller) MAXIMUMKEYSPERACCOUNT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ValidationManager.contract.Call(opts, &out, "MAXIMUM_KEYS_PER_ACCOUNT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMKEYSPERACCOUNT is a free data retrieval call binding the contract method 0xd2c1934f.
//
// Solidity: function MAXIMUM_KEYS_PER_ACCOUNT() view returns(uint8)
func (_ValidationManager *ValidationManagerSession) MAXIMUMKEYSPERACCOUNT() (uint8, error) {
	return _ValidationManager.Contract.MAXIMUMKEYSPERACCOUNT(&_ValidationManager.CallOpts)
}

// MAXIMUMKEYSPERACCOUNT is a free data retrieval call binding the contract method 0xd2c1934f.
//
// Solidity: function MAXIMUM_KEYS_PER_ACCOUNT() view returns(uint8)
func (_ValidationManager *ValidationManagerCallerSession) MAXIMUMKEYSPERACCOUNT() (uint8, error) {
	return _ValidationManager.Contract.MAXIMUMKEYSPERACCOUNT(&_ValidationManager.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x9bdafcb3.
//
// Solidity: function validators(bytes32 ) view returns(address validator, uint128 validFrom, uint128 validUntil)
func (_ValidationManager *ValidationManagerCaller) Validators(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Validator  common.Address
	ValidFrom  *big.Int
	ValidUntil *big.Int
}, error) {
	var out []interface{}
	err := _ValidationManager.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Validator  common.Address
		ValidFrom  *big.Int
		ValidUntil *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ValidFrom = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ValidUntil = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0x9bdafcb3.
//
// Solidity: function validators(bytes32 ) view returns(address validator, uint128 validFrom, uint128 validUntil)
func (_ValidationManager *ValidationManagerSession) Validators(arg0 [32]byte) (struct {
	Validator  common.Address
	ValidFrom  *big.Int
	ValidUntil *big.Int
}, error) {
	return _ValidationManager.Contract.Validators(&_ValidationManager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x9bdafcb3.
//
// Solidity: function validators(bytes32 ) view returns(address validator, uint128 validFrom, uint128 validUntil)
func (_ValidationManager *ValidationManagerCallerSession) Validators(arg0 [32]byte) (struct {
	Validator  common.Address
	ValidFrom  *big.Int
	ValidUntil *big.Int
}, error) {
	return _ValidationManager.Contract.Validators(&_ValidationManager.CallOpts, arg0)
}
