// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bnpairingprecompilecostestimator

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

// BNPairingPrecompileCostEstimatorMetaData contains all meta data concerning the BNPairingPrecompileCostEstimator contract.
var BNPairingPrecompileCostEstimatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"baseCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pairCount\",\"type\":\"uint256\"}],\"name\":\"getGasCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"perPairCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"run\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608080604052346015576106d3908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b6000803560e01c9081634e79f8ca1461005a575080639382255714610055578063c0406226146100505763ebfd94b21461004b57600080fd5b6101f0565b6100c2565b61009f565b3461009c57602036600319011261009c57600154600435818102929181159184041417156100975754810180911161009757602090604051908152f35b61020e565b80fd5b346100bd5760003660031901126100bd576020600054604051908152f35b600080fd5b346100bd5760008060031936011261009c576101eb6100df610257565b60018152600260208201527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c260408201527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed60608201527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b60808201527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa60a08201526101c961018d61029d565b6101c26101bc60208360c06101a15a610224565b976101af6107d05a116102c2565b60085a99fa945a90610234565b93610322565b5115610382565b6101e26101dd826101d861051a565b610234565b600155565b60015490610234565b815580f35b346100bd5760003660031901126100bd576020600154604051908152f35b634e487b7160e01b600052601160045260246000fd5b6107cf1981019190821161009757565b9190820391821161009757565b634e487b7160e01b600052604160045260246000fd5b6040519060c0820182811067ffffffffffffffff82111761027757604052565b610241565b60405190610180820182811067ffffffffffffffff82111761027757604052565b604051906020820182811067ffffffffffffffff821117610277576040526020368337565b156102c957565b60405162461bcd60e51b815260206004820152603d602482015260008051602061067e83398151915260448201527f3a206e6f7420656e6f756768206761732c2073696e676c6520706169720000006064820152608490fd5b1561032957565b60405162461bcd60e51b815260206004820152603c602482015260008051602061067e83398151915260448201527f3a2073696e676c6520706169722063616c6c206973206661696c6564000000006064820152608490fd5b1561038957565b60405162461bcd60e51b8152602060048201526043602482015260008051602061067e83398151915260448201527f3a2073696e676c6520706169722063616c6c20726573756c74206d757374206260648201526206520360ec1b608482015260a490fd5b156103f557565b60405162461bcd60e51b815260206004820152603d602482015260008051602061067e83398151915260448201527f3a206e6f7420656e6f756768206761732c20636f75706c6520706169720000006064820152608490fd5b1561045557565b60405162461bcd60e51b815260206004820152603c602482015260008051602061067e83398151915260448201527f3a20636f75706c6520706169722063616c6c206973206661696c6564000000006064820152608490fd5b156104b557565b60405162461bcd60e51b8152602060048201526043602482015260008051602061067e83398151915260448201527f3a20636f75706c6520706169722063616c6c20726573756c74206d757374206260648201526265203160e81b608482015260a490fd5b61052261027c565b60018152600260208201527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28060408301527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed908160608401527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b60808401527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa60a0840152600160c0840152600260e08401526101008301526101208201527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec6101408201527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d61016082015261067a600161063d61029d565b61067361066d6020836101806106525a610224565b986106606107d05a116103ee565b60085a9afa955a90610234565b9461044e565b51146104ae565b9056fe424e50616972696e67507265636f6d70696c65436f7374457374696d61746f72a2646970667358221220e5dd058ce2a4e339cfe36e5fa0e8455394f9498a130361fbaf90d29a762fae3964736f6c63430008190033",
}

// BNPairingPrecompileCostEstimatorABI is the input ABI used to generate the binding from.
// Deprecated: Use BNPairingPrecompileCostEstimatorMetaData.ABI instead.
var BNPairingPrecompileCostEstimatorABI = BNPairingPrecompileCostEstimatorMetaData.ABI

// BNPairingPrecompileCostEstimatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BNPairingPrecompileCostEstimatorMetaData.Bin instead.
var BNPairingPrecompileCostEstimatorBin = BNPairingPrecompileCostEstimatorMetaData.Bin

// DeployBNPairingPrecompileCostEstimator deploys a new Ethereum contract, binding an instance of BNPairingPrecompileCostEstimator to it.
func DeployBNPairingPrecompileCostEstimator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BNPairingPrecompileCostEstimator, error) {
	parsed, err := BNPairingPrecompileCostEstimatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BNPairingPrecompileCostEstimatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BNPairingPrecompileCostEstimator{BNPairingPrecompileCostEstimatorCaller: BNPairingPrecompileCostEstimatorCaller{contract: contract}, BNPairingPrecompileCostEstimatorTransactor: BNPairingPrecompileCostEstimatorTransactor{contract: contract}, BNPairingPrecompileCostEstimatorFilterer: BNPairingPrecompileCostEstimatorFilterer{contract: contract}}, nil
}

// BNPairingPrecompileCostEstimator is an auto generated Go binding around an Ethereum contract.
type BNPairingPrecompileCostEstimator struct {
	BNPairingPrecompileCostEstimatorCaller     // Read-only binding to the contract
	BNPairingPrecompileCostEstimatorTransactor // Write-only binding to the contract
	BNPairingPrecompileCostEstimatorFilterer   // Log filterer for contract events
}

// BNPairingPrecompileCostEstimatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BNPairingPrecompileCostEstimatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BNPairingPrecompileCostEstimatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BNPairingPrecompileCostEstimatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BNPairingPrecompileCostEstimatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BNPairingPrecompileCostEstimatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BNPairingPrecompileCostEstimatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BNPairingPrecompileCostEstimatorSession struct {
	Contract     *BNPairingPrecompileCostEstimator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// BNPairingPrecompileCostEstimatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BNPairingPrecompileCostEstimatorCallerSession struct {
	Contract *BNPairingPrecompileCostEstimatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// BNPairingPrecompileCostEstimatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BNPairingPrecompileCostEstimatorTransactorSession struct {
	Contract     *BNPairingPrecompileCostEstimatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// BNPairingPrecompileCostEstimatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BNPairingPrecompileCostEstimatorRaw struct {
	Contract *BNPairingPrecompileCostEstimator // Generic contract binding to access the raw methods on
}

// BNPairingPrecompileCostEstimatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BNPairingPrecompileCostEstimatorCallerRaw struct {
	Contract *BNPairingPrecompileCostEstimatorCaller // Generic read-only contract binding to access the raw methods on
}

// BNPairingPrecompileCostEstimatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BNPairingPrecompileCostEstimatorTransactorRaw struct {
	Contract *BNPairingPrecompileCostEstimatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBNPairingPrecompileCostEstimator creates a new instance of BNPairingPrecompileCostEstimator, bound to a specific deployed contract.
func NewBNPairingPrecompileCostEstimator(address common.Address, backend bind.ContractBackend) (*BNPairingPrecompileCostEstimator, error) {
	contract, err := bindBNPairingPrecompileCostEstimator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BNPairingPrecompileCostEstimator{BNPairingPrecompileCostEstimatorCaller: BNPairingPrecompileCostEstimatorCaller{contract: contract}, BNPairingPrecompileCostEstimatorTransactor: BNPairingPrecompileCostEstimatorTransactor{contract: contract}, BNPairingPrecompileCostEstimatorFilterer: BNPairingPrecompileCostEstimatorFilterer{contract: contract}}, nil
}

// NewBNPairingPrecompileCostEstimatorCaller creates a new read-only instance of BNPairingPrecompileCostEstimator, bound to a specific deployed contract.
func NewBNPairingPrecompileCostEstimatorCaller(address common.Address, caller bind.ContractCaller) (*BNPairingPrecompileCostEstimatorCaller, error) {
	contract, err := bindBNPairingPrecompileCostEstimator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BNPairingPrecompileCostEstimatorCaller{contract: contract}, nil
}

// NewBNPairingPrecompileCostEstimatorTransactor creates a new write-only instance of BNPairingPrecompileCostEstimator, bound to a specific deployed contract.
func NewBNPairingPrecompileCostEstimatorTransactor(address common.Address, transactor bind.ContractTransactor) (*BNPairingPrecompileCostEstimatorTransactor, error) {
	contract, err := bindBNPairingPrecompileCostEstimator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BNPairingPrecompileCostEstimatorTransactor{contract: contract}, nil
}

// NewBNPairingPrecompileCostEstimatorFilterer creates a new log filterer instance of BNPairingPrecompileCostEstimator, bound to a specific deployed contract.
func NewBNPairingPrecompileCostEstimatorFilterer(address common.Address, filterer bind.ContractFilterer) (*BNPairingPrecompileCostEstimatorFilterer, error) {
	contract, err := bindBNPairingPrecompileCostEstimator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BNPairingPrecompileCostEstimatorFilterer{contract: contract}, nil
}

// bindBNPairingPrecompileCostEstimator binds a generic wrapper to an already deployed contract.
func bindBNPairingPrecompileCostEstimator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BNPairingPrecompileCostEstimatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BNPairingPrecompileCostEstimator.Contract.BNPairingPrecompileCostEstimatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BNPairingPrecompileCostEstimator.Contract.BNPairingPrecompileCostEstimatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BNPairingPrecompileCostEstimator.Contract.BNPairingPrecompileCostEstimatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BNPairingPrecompileCostEstimator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BNPairingPrecompileCostEstimator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BNPairingPrecompileCostEstimator.Contract.contract.Transact(opts, method, params...)
}

// BaseCost is a free data retrieval call binding the contract method 0x93822557.
//
// Solidity: function baseCost() view returns(uint256)
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorCaller) BaseCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNPairingPrecompileCostEstimator.contract.Call(opts, &out, "baseCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseCost is a free data retrieval call binding the contract method 0x93822557.
//
// Solidity: function baseCost() view returns(uint256)
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorSession) BaseCost() (*big.Int, error) {
	return _BNPairingPrecompileCostEstimator.Contract.BaseCost(&_BNPairingPrecompileCostEstimator.CallOpts)
}

// BaseCost is a free data retrieval call binding the contract method 0x93822557.
//
// Solidity: function baseCost() view returns(uint256)
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorCallerSession) BaseCost() (*big.Int, error) {
	return _BNPairingPrecompileCostEstimator.Contract.BaseCost(&_BNPairingPrecompileCostEstimator.CallOpts)
}

// GetGasCost is a free data retrieval call binding the contract method 0x4e79f8ca.
//
// Solidity: function getGasCost(uint256 pairCount) view returns(uint256)
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorCaller) GetGasCost(opts *bind.CallOpts, pairCount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BNPairingPrecompileCostEstimator.contract.Call(opts, &out, "getGasCost", pairCount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGasCost is a free data retrieval call binding the contract method 0x4e79f8ca.
//
// Solidity: function getGasCost(uint256 pairCount) view returns(uint256)
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorSession) GetGasCost(pairCount *big.Int) (*big.Int, error) {
	return _BNPairingPrecompileCostEstimator.Contract.GetGasCost(&_BNPairingPrecompileCostEstimator.CallOpts, pairCount)
}

// GetGasCost is a free data retrieval call binding the contract method 0x4e79f8ca.
//
// Solidity: function getGasCost(uint256 pairCount) view returns(uint256)
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorCallerSession) GetGasCost(pairCount *big.Int) (*big.Int, error) {
	return _BNPairingPrecompileCostEstimator.Contract.GetGasCost(&_BNPairingPrecompileCostEstimator.CallOpts, pairCount)
}

// PerPairCost is a free data retrieval call binding the contract method 0xebfd94b2.
//
// Solidity: function perPairCost() view returns(uint256)
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorCaller) PerPairCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNPairingPrecompileCostEstimator.contract.Call(opts, &out, "perPairCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerPairCost is a free data retrieval call binding the contract method 0xebfd94b2.
//
// Solidity: function perPairCost() view returns(uint256)
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorSession) PerPairCost() (*big.Int, error) {
	return _BNPairingPrecompileCostEstimator.Contract.PerPairCost(&_BNPairingPrecompileCostEstimator.CallOpts)
}

// PerPairCost is a free data retrieval call binding the contract method 0xebfd94b2.
//
// Solidity: function perPairCost() view returns(uint256)
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorCallerSession) PerPairCost() (*big.Int, error) {
	return _BNPairingPrecompileCostEstimator.Contract.PerPairCost(&_BNPairingPrecompileCostEstimator.CallOpts)
}

// Run is a paid mutator transaction binding the contract method 0xc0406226.
//
// Solidity: function run() returns()
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorTransactor) Run(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BNPairingPrecompileCostEstimator.contract.Transact(opts, "run")
}

// Run is a paid mutator transaction binding the contract method 0xc0406226.
//
// Solidity: function run() returns()
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorSession) Run() (*types.Transaction, error) {
	return _BNPairingPrecompileCostEstimator.Contract.Run(&_BNPairingPrecompileCostEstimator.TransactOpts)
}

// Run is a paid mutator transaction binding the contract method 0xc0406226.
//
// Solidity: function run() returns()
func (_BNPairingPrecompileCostEstimator *BNPairingPrecompileCostEstimatorTransactorSession) Run() (*types.Transaction, error) {
	return _BNPairingPrecompileCostEstimator.Contract.Run(&_BNPairingPrecompileCostEstimator.TransactOpts)
}
