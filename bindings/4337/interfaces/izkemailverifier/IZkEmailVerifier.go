// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izkemailverifier

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

// IZkEmailVerifierMetaData contains all meta data concerning the IZkEmailVerifier contract.
var IZkEmailVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_newPubKeyHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"emailHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dkimKeyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IZkEmailVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IZkEmailVerifierMetaData.ABI instead.
var IZkEmailVerifierABI = IZkEmailVerifierMetaData.ABI

// IZkEmailVerifier is an auto generated Go binding around an Ethereum contract.
type IZkEmailVerifier struct {
	IZkEmailVerifierCaller     // Read-only binding to the contract
	IZkEmailVerifierTransactor // Write-only binding to the contract
	IZkEmailVerifierFilterer   // Log filterer for contract events
}

// IZkEmailVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZkEmailVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkEmailVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZkEmailVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkEmailVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZkEmailVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkEmailVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZkEmailVerifierSession struct {
	Contract     *IZkEmailVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IZkEmailVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZkEmailVerifierCallerSession struct {
	Contract *IZkEmailVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IZkEmailVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZkEmailVerifierTransactorSession struct {
	Contract     *IZkEmailVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IZkEmailVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZkEmailVerifierRaw struct {
	Contract *IZkEmailVerifier // Generic contract binding to access the raw methods on
}

// IZkEmailVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZkEmailVerifierCallerRaw struct {
	Contract *IZkEmailVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IZkEmailVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZkEmailVerifierTransactorRaw struct {
	Contract *IZkEmailVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZkEmailVerifier creates a new instance of IZkEmailVerifier, bound to a specific deployed contract.
func NewIZkEmailVerifier(address common.Address, backend bind.ContractBackend) (*IZkEmailVerifier, error) {
	contract, err := bindIZkEmailVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZkEmailVerifier{IZkEmailVerifierCaller: IZkEmailVerifierCaller{contract: contract}, IZkEmailVerifierTransactor: IZkEmailVerifierTransactor{contract: contract}, IZkEmailVerifierFilterer: IZkEmailVerifierFilterer{contract: contract}}, nil
}

// NewIZkEmailVerifierCaller creates a new read-only instance of IZkEmailVerifier, bound to a specific deployed contract.
func NewIZkEmailVerifierCaller(address common.Address, caller bind.ContractCaller) (*IZkEmailVerifierCaller, error) {
	contract, err := bindIZkEmailVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZkEmailVerifierCaller{contract: contract}, nil
}

// NewIZkEmailVerifierTransactor creates a new write-only instance of IZkEmailVerifier, bound to a specific deployed contract.
func NewIZkEmailVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IZkEmailVerifierTransactor, error) {
	contract, err := bindIZkEmailVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZkEmailVerifierTransactor{contract: contract}, nil
}

// NewIZkEmailVerifierFilterer creates a new log filterer instance of IZkEmailVerifier, bound to a specific deployed contract.
func NewIZkEmailVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IZkEmailVerifierFilterer, error) {
	contract, err := bindIZkEmailVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZkEmailVerifierFilterer{contract: contract}, nil
}

// bindIZkEmailVerifier binds a generic wrapper to an already deployed contract.
func bindIZkEmailVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZkEmailVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZkEmailVerifier *IZkEmailVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZkEmailVerifier.Contract.IZkEmailVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZkEmailVerifier *IZkEmailVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkEmailVerifier.Contract.IZkEmailVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZkEmailVerifier *IZkEmailVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZkEmailVerifier.Contract.IZkEmailVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZkEmailVerifier *IZkEmailVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZkEmailVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZkEmailVerifier *IZkEmailVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkEmailVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZkEmailVerifier *IZkEmailVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZkEmailVerifier.Contract.contract.Transact(opts, method, params...)
}

// Verify is a free data retrieval call binding the contract method 0x9967af67.
//
// Solidity: function verify(address _account, address _validator, bytes32 _newPubKeyHash, bytes _proof) view returns(bool success, bytes32 emailHash, bytes32 dkimKeyHash, uint256 timestamp, string domain)
func (_IZkEmailVerifier *IZkEmailVerifierCaller) Verify(opts *bind.CallOpts, _account common.Address, _validator common.Address, _newPubKeyHash [32]byte, _proof []byte) (struct {
	Success     bool
	EmailHash   [32]byte
	DkimKeyHash [32]byte
	Timestamp   *big.Int
	Domain      string
}, error) {
	var out []interface{}
	err := _IZkEmailVerifier.contract.Call(opts, &out, "verify", _account, _validator, _newPubKeyHash, _proof)

	outstruct := new(struct {
		Success     bool
		EmailHash   [32]byte
		DkimKeyHash [32]byte
		Timestamp   *big.Int
		Domain      string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Success = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.EmailHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.DkimKeyHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Domain = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// Verify is a free data retrieval call binding the contract method 0x9967af67.
//
// Solidity: function verify(address _account, address _validator, bytes32 _newPubKeyHash, bytes _proof) view returns(bool success, bytes32 emailHash, bytes32 dkimKeyHash, uint256 timestamp, string domain)
func (_IZkEmailVerifier *IZkEmailVerifierSession) Verify(_account common.Address, _validator common.Address, _newPubKeyHash [32]byte, _proof []byte) (struct {
	Success     bool
	EmailHash   [32]byte
	DkimKeyHash [32]byte
	Timestamp   *big.Int
	Domain      string
}, error) {
	return _IZkEmailVerifier.Contract.Verify(&_IZkEmailVerifier.CallOpts, _account, _validator, _newPubKeyHash, _proof)
}

// Verify is a free data retrieval call binding the contract method 0x9967af67.
//
// Solidity: function verify(address _account, address _validator, bytes32 _newPubKeyHash, bytes _proof) view returns(bool success, bytes32 emailHash, bytes32 dkimKeyHash, uint256 timestamp, string domain)
func (_IZkEmailVerifier *IZkEmailVerifierCallerSession) Verify(_account common.Address, _validator common.Address, _newPubKeyHash [32]byte, _proof []byte) (struct {
	Success     bool
	EmailHash   [32]byte
	DkimKeyHash [32]byte
	Timestamp   *big.Int
	Domain      string
}, error) {
	return _IZkEmailVerifier.Contract.Verify(&_IZkEmailVerifier.CallOpts, _account, _validator, _newPubKeyHash, _proof)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_IZkEmailVerifier *IZkEmailVerifierTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _IZkEmailVerifier.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_IZkEmailVerifier *IZkEmailVerifierSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _IZkEmailVerifier.Contract.Initialize(&_IZkEmailVerifier.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_IZkEmailVerifier *IZkEmailVerifierTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _IZkEmailVerifier.Contract.Initialize(&_IZkEmailVerifier.TransactOpts, _owner)
}
