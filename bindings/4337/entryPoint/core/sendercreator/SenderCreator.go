// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sendercreator

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

// SenderCreatorMetaData contains all meta data concerning the SenderCreator contract.
var SenderCreatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"name\":\"createSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557610144908161001b8239f35b600080fdfe6080600436101561000f57600080fd5b6000803560e01c63570e1a361461002557600080fd5b3461010b57602036600319011261010b576004359167ffffffffffffffff9081841161010757366023850112156101075783600401358281116101035736602482870101116101035780601411610103576013198101928084116100ef57600b8201601f19908116603f01168301908111838210176100ef5792846024819482600c60209a968b9960405286845289840196603889018837830101525193013560601c5af190805191156100e7575b506040516001600160a01b039091168152f35b9050386100d4565b634e487b7160e01b85526041600452602485fd5b8380fd5b8280fd5b80fdfea26469706673582212206a1b9f618c552b5e5a558cda5cc2bafd851f85b1d0662028a1684dd5741f8d9464736f6c63430008190033",
}

// SenderCreatorABI is the input ABI used to generate the binding from.
// Deprecated: Use SenderCreatorMetaData.ABI instead.
var SenderCreatorABI = SenderCreatorMetaData.ABI

// SenderCreatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SenderCreatorMetaData.Bin instead.
var SenderCreatorBin = SenderCreatorMetaData.Bin

// DeploySenderCreator deploys a new Ethereum contract, binding an instance of SenderCreator to it.
func DeploySenderCreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SenderCreator, error) {
	parsed, err := SenderCreatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SenderCreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SenderCreator{SenderCreatorCaller: SenderCreatorCaller{contract: contract}, SenderCreatorTransactor: SenderCreatorTransactor{contract: contract}, SenderCreatorFilterer: SenderCreatorFilterer{contract: contract}}, nil
}

// SenderCreator is an auto generated Go binding around an Ethereum contract.
type SenderCreator struct {
	SenderCreatorCaller     // Read-only binding to the contract
	SenderCreatorTransactor // Write-only binding to the contract
	SenderCreatorFilterer   // Log filterer for contract events
}

// SenderCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenderCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenderCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenderCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenderCreatorSession struct {
	Contract     *SenderCreator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenderCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenderCreatorCallerSession struct {
	Contract *SenderCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SenderCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenderCreatorTransactorSession struct {
	Contract     *SenderCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SenderCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenderCreatorRaw struct {
	Contract *SenderCreator // Generic contract binding to access the raw methods on
}

// SenderCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenderCreatorCallerRaw struct {
	Contract *SenderCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// SenderCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenderCreatorTransactorRaw struct {
	Contract *SenderCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSenderCreator creates a new instance of SenderCreator, bound to a specific deployed contract.
func NewSenderCreator(address common.Address, backend bind.ContractBackend) (*SenderCreator, error) {
	contract, err := bindSenderCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SenderCreator{SenderCreatorCaller: SenderCreatorCaller{contract: contract}, SenderCreatorTransactor: SenderCreatorTransactor{contract: contract}, SenderCreatorFilterer: SenderCreatorFilterer{contract: contract}}, nil
}

// NewSenderCreatorCaller creates a new read-only instance of SenderCreator, bound to a specific deployed contract.
func NewSenderCreatorCaller(address common.Address, caller bind.ContractCaller) (*SenderCreatorCaller, error) {
	contract, err := bindSenderCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenderCreatorCaller{contract: contract}, nil
}

// NewSenderCreatorTransactor creates a new write-only instance of SenderCreator, bound to a specific deployed contract.
func NewSenderCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*SenderCreatorTransactor, error) {
	contract, err := bindSenderCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenderCreatorTransactor{contract: contract}, nil
}

// NewSenderCreatorFilterer creates a new log filterer instance of SenderCreator, bound to a specific deployed contract.
func NewSenderCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*SenderCreatorFilterer, error) {
	contract, err := bindSenderCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenderCreatorFilterer{contract: contract}, nil
}

// bindSenderCreator binds a generic wrapper to an already deployed contract.
func bindSenderCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SenderCreatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenderCreator *SenderCreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenderCreator.Contract.SenderCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenderCreator *SenderCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenderCreator.Contract.SenderCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenderCreator *SenderCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenderCreator.Contract.SenderCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenderCreator *SenderCreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenderCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenderCreator *SenderCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenderCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenderCreator *SenderCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenderCreator.Contract.contract.Transact(opts, method, params...)
}

// CreateSender is a paid mutator transaction binding the contract method 0x570e1a36.
//
// Solidity: function createSender(bytes initCode) returns(address sender)
func (_SenderCreator *SenderCreatorTransactor) CreateSender(opts *bind.TransactOpts, initCode []byte) (*types.Transaction, error) {
	return _SenderCreator.contract.Transact(opts, "createSender", initCode)
}

// CreateSender is a paid mutator transaction binding the contract method 0x570e1a36.
//
// Solidity: function createSender(bytes initCode) returns(address sender)
func (_SenderCreator *SenderCreatorSession) CreateSender(initCode []byte) (*types.Transaction, error) {
	return _SenderCreator.Contract.CreateSender(&_SenderCreator.TransactOpts, initCode)
}

// CreateSender is a paid mutator transaction binding the contract method 0x570e1a36.
//
// Solidity: function createSender(bytes initCode) returns(address sender)
func (_SenderCreator *SenderCreatorTransactorSession) CreateSender(initCode []byte) (*types.Transaction, error) {
	return _SenderCreator.Contract.CreateSender(&_SenderCreator.TransactOpts, initCode)
}
