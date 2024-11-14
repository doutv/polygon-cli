// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blshelper

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

// BLSHelperMetaData contains all meta data concerning the BLSHelper contract.
var BLSHelperMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212208d2ec0d4d6df77d7db0c60445e9a37fe239f4a5380b226afea10c7101969b1fb64736f6c63430008190033",
}

// BLSHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use BLSHelperMetaData.ABI instead.
var BLSHelperABI = BLSHelperMetaData.ABI

// BLSHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BLSHelperMetaData.Bin instead.
var BLSHelperBin = BLSHelperMetaData.Bin

// DeployBLSHelper deploys a new Ethereum contract, binding an instance of BLSHelper to it.
func DeployBLSHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BLSHelper, error) {
	parsed, err := BLSHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BLSHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BLSHelper{BLSHelperCaller: BLSHelperCaller{contract: contract}, BLSHelperTransactor: BLSHelperTransactor{contract: contract}, BLSHelperFilterer: BLSHelperFilterer{contract: contract}}, nil
}

// BLSHelper is an auto generated Go binding around an Ethereum contract.
type BLSHelper struct {
	BLSHelperCaller     // Read-only binding to the contract
	BLSHelperTransactor // Write-only binding to the contract
	BLSHelperFilterer   // Log filterer for contract events
}

// BLSHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type BLSHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BLSHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BLSHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BLSHelperSession struct {
	Contract     *BLSHelper        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BLSHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BLSHelperCallerSession struct {
	Contract *BLSHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BLSHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BLSHelperTransactorSession struct {
	Contract     *BLSHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BLSHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type BLSHelperRaw struct {
	Contract *BLSHelper // Generic contract binding to access the raw methods on
}

// BLSHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BLSHelperCallerRaw struct {
	Contract *BLSHelperCaller // Generic read-only contract binding to access the raw methods on
}

// BLSHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BLSHelperTransactorRaw struct {
	Contract *BLSHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBLSHelper creates a new instance of BLSHelper, bound to a specific deployed contract.
func NewBLSHelper(address common.Address, backend bind.ContractBackend) (*BLSHelper, error) {
	contract, err := bindBLSHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BLSHelper{BLSHelperCaller: BLSHelperCaller{contract: contract}, BLSHelperTransactor: BLSHelperTransactor{contract: contract}, BLSHelperFilterer: BLSHelperFilterer{contract: contract}}, nil
}

// NewBLSHelperCaller creates a new read-only instance of BLSHelper, bound to a specific deployed contract.
func NewBLSHelperCaller(address common.Address, caller bind.ContractCaller) (*BLSHelperCaller, error) {
	contract, err := bindBLSHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BLSHelperCaller{contract: contract}, nil
}

// NewBLSHelperTransactor creates a new write-only instance of BLSHelper, bound to a specific deployed contract.
func NewBLSHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*BLSHelperTransactor, error) {
	contract, err := bindBLSHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BLSHelperTransactor{contract: contract}, nil
}

// NewBLSHelperFilterer creates a new log filterer instance of BLSHelper, bound to a specific deployed contract.
func NewBLSHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*BLSHelperFilterer, error) {
	contract, err := bindBLSHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BLSHelperFilterer{contract: contract}, nil
}

// bindBLSHelper binds a generic wrapper to an already deployed contract.
func bindBLSHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BLSHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSHelper *BLSHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSHelper.Contract.BLSHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSHelper *BLSHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSHelper.Contract.BLSHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSHelper *BLSHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSHelper.Contract.BLSHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSHelper *BLSHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSHelper *BLSHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSHelper *BLSHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSHelper.Contract.contract.Transact(opts, method, params...)
}
