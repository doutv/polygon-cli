// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package webauthn

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

// WebAuthnMetaData contains all meta data concerning the WebAuthn contract.
var WebAuthnMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212202a5cceebd56f9cb1ab782754256ce0b2fd14ec748a9cf5a6d4677a5f97db26dd64736f6c63430008190033",
}

// WebAuthnABI is the input ABI used to generate the binding from.
// Deprecated: Use WebAuthnMetaData.ABI instead.
var WebAuthnABI = WebAuthnMetaData.ABI

// WebAuthnBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WebAuthnMetaData.Bin instead.
var WebAuthnBin = WebAuthnMetaData.Bin

// DeployWebAuthn deploys a new Ethereum contract, binding an instance of WebAuthn to it.
func DeployWebAuthn(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WebAuthn, error) {
	parsed, err := WebAuthnMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WebAuthnBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WebAuthn{WebAuthnCaller: WebAuthnCaller{contract: contract}, WebAuthnTransactor: WebAuthnTransactor{contract: contract}, WebAuthnFilterer: WebAuthnFilterer{contract: contract}}, nil
}

// WebAuthn is an auto generated Go binding around an Ethereum contract.
type WebAuthn struct {
	WebAuthnCaller     // Read-only binding to the contract
	WebAuthnTransactor // Write-only binding to the contract
	WebAuthnFilterer   // Log filterer for contract events
}

// WebAuthnCaller is an auto generated read-only Go binding around an Ethereum contract.
type WebAuthnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebAuthnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WebAuthnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebAuthnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WebAuthnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebAuthnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WebAuthnSession struct {
	Contract     *WebAuthn         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WebAuthnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WebAuthnCallerSession struct {
	Contract *WebAuthnCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// WebAuthnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WebAuthnTransactorSession struct {
	Contract     *WebAuthnTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WebAuthnRaw is an auto generated low-level Go binding around an Ethereum contract.
type WebAuthnRaw struct {
	Contract *WebAuthn // Generic contract binding to access the raw methods on
}

// WebAuthnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WebAuthnCallerRaw struct {
	Contract *WebAuthnCaller // Generic read-only contract binding to access the raw methods on
}

// WebAuthnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WebAuthnTransactorRaw struct {
	Contract *WebAuthnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWebAuthn creates a new instance of WebAuthn, bound to a specific deployed contract.
func NewWebAuthn(address common.Address, backend bind.ContractBackend) (*WebAuthn, error) {
	contract, err := bindWebAuthn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WebAuthn{WebAuthnCaller: WebAuthnCaller{contract: contract}, WebAuthnTransactor: WebAuthnTransactor{contract: contract}, WebAuthnFilterer: WebAuthnFilterer{contract: contract}}, nil
}

// NewWebAuthnCaller creates a new read-only instance of WebAuthn, bound to a specific deployed contract.
func NewWebAuthnCaller(address common.Address, caller bind.ContractCaller) (*WebAuthnCaller, error) {
	contract, err := bindWebAuthn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WebAuthnCaller{contract: contract}, nil
}

// NewWebAuthnTransactor creates a new write-only instance of WebAuthn, bound to a specific deployed contract.
func NewWebAuthnTransactor(address common.Address, transactor bind.ContractTransactor) (*WebAuthnTransactor, error) {
	contract, err := bindWebAuthn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WebAuthnTransactor{contract: contract}, nil
}

// NewWebAuthnFilterer creates a new log filterer instance of WebAuthn, bound to a specific deployed contract.
func NewWebAuthnFilterer(address common.Address, filterer bind.ContractFilterer) (*WebAuthnFilterer, error) {
	contract, err := bindWebAuthn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WebAuthnFilterer{contract: contract}, nil
}

// bindWebAuthn binds a generic wrapper to an already deployed contract.
func bindWebAuthn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WebAuthnMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WebAuthn *WebAuthnRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WebAuthn.Contract.WebAuthnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WebAuthn *WebAuthnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WebAuthn.Contract.WebAuthnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WebAuthn *WebAuthnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WebAuthn.Contract.WebAuthnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WebAuthn *WebAuthnCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WebAuthn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WebAuthn *WebAuthnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WebAuthn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WebAuthn *WebAuthnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WebAuthn.Contract.contract.Transact(opts, method, params...)
}
