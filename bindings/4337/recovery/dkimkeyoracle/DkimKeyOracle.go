// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dkimkeyoracle

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

// DkimKeyOracleMetaData contains all meta data concerning the DkimKeyOracle contract.
var DkimKeyOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"DkimKeyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"}],\"name\":\"DkimKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"PauserUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getDomainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"}],\"name\":\"removeKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_domain\",\"type\":\"bytes32\"}],\"name\":\"updateKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"updatePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080346100d457601f61048538819003918201601f19168301916001600160401b038311848410176100d95780849260409485528339810103126100d457610052602061004b836100ef565b92016100ef565b906001600160a01b039081169081156100bb576000549260018060a01b03199280848616176000558260405195167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a31690600254161760025561038190816101048239f35b604051631e4fbdf760e01b815260006004820152602490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036100d45756fe608060408181526004908136101561001657600080fd5b600092833560e01c908163368a1b76146102f657508063554bab3c14610288578063715018a61461022b57806373af6745146101d1578063862642f5146101615780638da5cb5b146101395780639fd0506d1461010c5763f2fde38b1461007c57600080fd5b34610108576020366003190112610108576001600160a01b03823581811693919290849003610104576100ad61031f565b83156100ee57505082546001600160a01b0319811683178455167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b51631e4fbdf760e01b8152908101849052602490fd5b8480fd5b8280fd5b83823461013557816003193601126101355760025490516001600160a01b039091168152602090f35b5080fd5b838234610135578160031936011261013557905490516001600160a01b039091168152602090f35b503461010857602036600319011261010857600254823592906001600160a01b031633036101c35750816020917fa54c86c49f4e97ce7b0c5af2e11859918d86f47f03793f03f33db5b7c349d0d193855260018352848181205551908152a180f35b9051634ca8886760e01b8152fd5b50346101085780600319360112610108577fe780335b851f95261f942f1897645485db8bd8c5ff0fc4dff7b5cc2de72d2d46913560243561021061031f565b8185526001602052808386205582519182526020820152a180f35b833461028557806003193601126102855761024461031f565b80546001600160a01b03198116825581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b50903461010857602036600319011261010857356001600160a01b0381169190829003610108577fa4336c0cb1e245b95ad204faed7e940d6dc999684fd8b5e1ff597a0c4efca8ab916020916102dc61031f565b600280546001600160a01b0319168317905551908152a180f35b9290503461031b57602036600319011261031b57926020933581526001845220548152f35b8380fd5b6000546001600160a01b0316330361033357565b60405163118cdaa760e01b8152336004820152602490fdfea264697066735822122013018826ea2f8bfef8838a64ec31ff02380b8ff937a7e1f6e59df9c74f9d990c64736f6c63430008190033",
}

// DkimKeyOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use DkimKeyOracleMetaData.ABI instead.
var DkimKeyOracleABI = DkimKeyOracleMetaData.ABI

// DkimKeyOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DkimKeyOracleMetaData.Bin instead.
var DkimKeyOracleBin = DkimKeyOracleMetaData.Bin

// DeployDkimKeyOracle deploys a new Ethereum contract, binding an instance of DkimKeyOracle to it.
func DeployDkimKeyOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _pauser common.Address) (common.Address, *types.Transaction, *DkimKeyOracle, error) {
	parsed, err := DkimKeyOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DkimKeyOracleBin), backend, _owner, _pauser)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DkimKeyOracle{DkimKeyOracleCaller: DkimKeyOracleCaller{contract: contract}, DkimKeyOracleTransactor: DkimKeyOracleTransactor{contract: contract}, DkimKeyOracleFilterer: DkimKeyOracleFilterer{contract: contract}}, nil
}

// DkimKeyOracle is an auto generated Go binding around an Ethereum contract.
type DkimKeyOracle struct {
	DkimKeyOracleCaller     // Read-only binding to the contract
	DkimKeyOracleTransactor // Write-only binding to the contract
	DkimKeyOracleFilterer   // Log filterer for contract events
}

// DkimKeyOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type DkimKeyOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DkimKeyOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DkimKeyOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DkimKeyOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DkimKeyOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DkimKeyOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DkimKeyOracleSession struct {
	Contract     *DkimKeyOracle    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DkimKeyOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DkimKeyOracleCallerSession struct {
	Contract *DkimKeyOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DkimKeyOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DkimKeyOracleTransactorSession struct {
	Contract     *DkimKeyOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DkimKeyOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type DkimKeyOracleRaw struct {
	Contract *DkimKeyOracle // Generic contract binding to access the raw methods on
}

// DkimKeyOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DkimKeyOracleCallerRaw struct {
	Contract *DkimKeyOracleCaller // Generic read-only contract binding to access the raw methods on
}

// DkimKeyOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DkimKeyOracleTransactorRaw struct {
	Contract *DkimKeyOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDkimKeyOracle creates a new instance of DkimKeyOracle, bound to a specific deployed contract.
func NewDkimKeyOracle(address common.Address, backend bind.ContractBackend) (*DkimKeyOracle, error) {
	contract, err := bindDkimKeyOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DkimKeyOracle{DkimKeyOracleCaller: DkimKeyOracleCaller{contract: contract}, DkimKeyOracleTransactor: DkimKeyOracleTransactor{contract: contract}, DkimKeyOracleFilterer: DkimKeyOracleFilterer{contract: contract}}, nil
}

// NewDkimKeyOracleCaller creates a new read-only instance of DkimKeyOracle, bound to a specific deployed contract.
func NewDkimKeyOracleCaller(address common.Address, caller bind.ContractCaller) (*DkimKeyOracleCaller, error) {
	contract, err := bindDkimKeyOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DkimKeyOracleCaller{contract: contract}, nil
}

// NewDkimKeyOracleTransactor creates a new write-only instance of DkimKeyOracle, bound to a specific deployed contract.
func NewDkimKeyOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*DkimKeyOracleTransactor, error) {
	contract, err := bindDkimKeyOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DkimKeyOracleTransactor{contract: contract}, nil
}

// NewDkimKeyOracleFilterer creates a new log filterer instance of DkimKeyOracle, bound to a specific deployed contract.
func NewDkimKeyOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*DkimKeyOracleFilterer, error) {
	contract, err := bindDkimKeyOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DkimKeyOracleFilterer{contract: contract}, nil
}

// bindDkimKeyOracle binds a generic wrapper to an already deployed contract.
func bindDkimKeyOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DkimKeyOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DkimKeyOracle *DkimKeyOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DkimKeyOracle.Contract.DkimKeyOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DkimKeyOracle *DkimKeyOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.DkimKeyOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DkimKeyOracle *DkimKeyOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.DkimKeyOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DkimKeyOracle *DkimKeyOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DkimKeyOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DkimKeyOracle *DkimKeyOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DkimKeyOracle *DkimKeyOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.contract.Transact(opts, method, params...)
}

// GetDomainHash is a free data retrieval call binding the contract method 0x368a1b76.
//
// Solidity: function getDomainHash(bytes32 ) view returns(bytes32)
func (_DkimKeyOracle *DkimKeyOracleCaller) GetDomainHash(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DkimKeyOracle.contract.Call(opts, &out, "getDomainHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainHash is a free data retrieval call binding the contract method 0x368a1b76.
//
// Solidity: function getDomainHash(bytes32 ) view returns(bytes32)
func (_DkimKeyOracle *DkimKeyOracleSession) GetDomainHash(arg0 [32]byte) ([32]byte, error) {
	return _DkimKeyOracle.Contract.GetDomainHash(&_DkimKeyOracle.CallOpts, arg0)
}

// GetDomainHash is a free data retrieval call binding the contract method 0x368a1b76.
//
// Solidity: function getDomainHash(bytes32 ) view returns(bytes32)
func (_DkimKeyOracle *DkimKeyOracleCallerSession) GetDomainHash(arg0 [32]byte) ([32]byte, error) {
	return _DkimKeyOracle.Contract.GetDomainHash(&_DkimKeyOracle.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DkimKeyOracle *DkimKeyOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DkimKeyOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DkimKeyOracle *DkimKeyOracleSession) Owner() (common.Address, error) {
	return _DkimKeyOracle.Contract.Owner(&_DkimKeyOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DkimKeyOracle *DkimKeyOracleCallerSession) Owner() (common.Address, error) {
	return _DkimKeyOracle.Contract.Owner(&_DkimKeyOracle.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_DkimKeyOracle *DkimKeyOracleCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DkimKeyOracle.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_DkimKeyOracle *DkimKeyOracleSession) Pauser() (common.Address, error) {
	return _DkimKeyOracle.Contract.Pauser(&_DkimKeyOracle.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_DkimKeyOracle *DkimKeyOracleCallerSession) Pauser() (common.Address, error) {
	return _DkimKeyOracle.Contract.Pauser(&_DkimKeyOracle.CallOpts)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x862642f5.
//
// Solidity: function removeKey(bytes32 _keyHash) returns()
func (_DkimKeyOracle *DkimKeyOracleTransactor) RemoveKey(opts *bind.TransactOpts, _keyHash [32]byte) (*types.Transaction, error) {
	return _DkimKeyOracle.contract.Transact(opts, "removeKey", _keyHash)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x862642f5.
//
// Solidity: function removeKey(bytes32 _keyHash) returns()
func (_DkimKeyOracle *DkimKeyOracleSession) RemoveKey(_keyHash [32]byte) (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.RemoveKey(&_DkimKeyOracle.TransactOpts, _keyHash)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x862642f5.
//
// Solidity: function removeKey(bytes32 _keyHash) returns()
func (_DkimKeyOracle *DkimKeyOracleTransactorSession) RemoveKey(_keyHash [32]byte) (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.RemoveKey(&_DkimKeyOracle.TransactOpts, _keyHash)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DkimKeyOracle *DkimKeyOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DkimKeyOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DkimKeyOracle *DkimKeyOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.RenounceOwnership(&_DkimKeyOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DkimKeyOracle *DkimKeyOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.RenounceOwnership(&_DkimKeyOracle.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DkimKeyOracle *DkimKeyOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DkimKeyOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DkimKeyOracle *DkimKeyOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.TransferOwnership(&_DkimKeyOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DkimKeyOracle *DkimKeyOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.TransferOwnership(&_DkimKeyOracle.TransactOpts, newOwner)
}

// UpdateKey is a paid mutator transaction binding the contract method 0x73af6745.
//
// Solidity: function updateKey(bytes32 _keyHash, bytes32 _domain) returns()
func (_DkimKeyOracle *DkimKeyOracleTransactor) UpdateKey(opts *bind.TransactOpts, _keyHash [32]byte, _domain [32]byte) (*types.Transaction, error) {
	return _DkimKeyOracle.contract.Transact(opts, "updateKey", _keyHash, _domain)
}

// UpdateKey is a paid mutator transaction binding the contract method 0x73af6745.
//
// Solidity: function updateKey(bytes32 _keyHash, bytes32 _domain) returns()
func (_DkimKeyOracle *DkimKeyOracleSession) UpdateKey(_keyHash [32]byte, _domain [32]byte) (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.UpdateKey(&_DkimKeyOracle.TransactOpts, _keyHash, _domain)
}

// UpdateKey is a paid mutator transaction binding the contract method 0x73af6745.
//
// Solidity: function updateKey(bytes32 _keyHash, bytes32 _domain) returns()
func (_DkimKeyOracle *DkimKeyOracleTransactorSession) UpdateKey(_keyHash [32]byte, _domain [32]byte) (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.UpdateKey(&_DkimKeyOracle.TransactOpts, _keyHash, _domain)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _pauser) returns()
func (_DkimKeyOracle *DkimKeyOracleTransactor) UpdatePauser(opts *bind.TransactOpts, _pauser common.Address) (*types.Transaction, error) {
	return _DkimKeyOracle.contract.Transact(opts, "updatePauser", _pauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _pauser) returns()
func (_DkimKeyOracle *DkimKeyOracleSession) UpdatePauser(_pauser common.Address) (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.UpdatePauser(&_DkimKeyOracle.TransactOpts, _pauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _pauser) returns()
func (_DkimKeyOracle *DkimKeyOracleTransactorSession) UpdatePauser(_pauser common.Address) (*types.Transaction, error) {
	return _DkimKeyOracle.Contract.UpdatePauser(&_DkimKeyOracle.TransactOpts, _pauser)
}

// DkimKeyOracleDkimKeyRemovedIterator is returned from FilterDkimKeyRemoved and is used to iterate over the raw logs and unpacked data for DkimKeyRemoved events raised by the DkimKeyOracle contract.
type DkimKeyOracleDkimKeyRemovedIterator struct {
	Event *DkimKeyOracleDkimKeyRemoved // Event containing the contract specifics and raw log

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
func (it *DkimKeyOracleDkimKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DkimKeyOracleDkimKeyRemoved)
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
		it.Event = new(DkimKeyOracleDkimKeyRemoved)
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
func (it *DkimKeyOracleDkimKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DkimKeyOracleDkimKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DkimKeyOracleDkimKeyRemoved represents a DkimKeyRemoved event raised by the DkimKeyOracle contract.
type DkimKeyOracleDkimKeyRemoved struct {
	KeyHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDkimKeyRemoved is a free log retrieval operation binding the contract event 0xa54c86c49f4e97ce7b0c5af2e11859918d86f47f03793f03f33db5b7c349d0d1.
//
// Solidity: event DkimKeyRemoved(bytes32 keyHash)
func (_DkimKeyOracle *DkimKeyOracleFilterer) FilterDkimKeyRemoved(opts *bind.FilterOpts) (*DkimKeyOracleDkimKeyRemovedIterator, error) {

	logs, sub, err := _DkimKeyOracle.contract.FilterLogs(opts, "DkimKeyRemoved")
	if err != nil {
		return nil, err
	}
	return &DkimKeyOracleDkimKeyRemovedIterator{contract: _DkimKeyOracle.contract, event: "DkimKeyRemoved", logs: logs, sub: sub}, nil
}

// WatchDkimKeyRemoved is a free log subscription operation binding the contract event 0xa54c86c49f4e97ce7b0c5af2e11859918d86f47f03793f03f33db5b7c349d0d1.
//
// Solidity: event DkimKeyRemoved(bytes32 keyHash)
func (_DkimKeyOracle *DkimKeyOracleFilterer) WatchDkimKeyRemoved(opts *bind.WatchOpts, sink chan<- *DkimKeyOracleDkimKeyRemoved) (event.Subscription, error) {

	logs, sub, err := _DkimKeyOracle.contract.WatchLogs(opts, "DkimKeyRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DkimKeyOracleDkimKeyRemoved)
				if err := _DkimKeyOracle.contract.UnpackLog(event, "DkimKeyRemoved", log); err != nil {
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

// ParseDkimKeyRemoved is a log parse operation binding the contract event 0xa54c86c49f4e97ce7b0c5af2e11859918d86f47f03793f03f33db5b7c349d0d1.
//
// Solidity: event DkimKeyRemoved(bytes32 keyHash)
func (_DkimKeyOracle *DkimKeyOracleFilterer) ParseDkimKeyRemoved(log types.Log) (*DkimKeyOracleDkimKeyRemoved, error) {
	event := new(DkimKeyOracleDkimKeyRemoved)
	if err := _DkimKeyOracle.contract.UnpackLog(event, "DkimKeyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DkimKeyOracleDkimKeyUpdatedIterator is returned from FilterDkimKeyUpdated and is used to iterate over the raw logs and unpacked data for DkimKeyUpdated events raised by the DkimKeyOracle contract.
type DkimKeyOracleDkimKeyUpdatedIterator struct {
	Event *DkimKeyOracleDkimKeyUpdated // Event containing the contract specifics and raw log

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
func (it *DkimKeyOracleDkimKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DkimKeyOracleDkimKeyUpdated)
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
		it.Event = new(DkimKeyOracleDkimKeyUpdated)
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
func (it *DkimKeyOracleDkimKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DkimKeyOracleDkimKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DkimKeyOracleDkimKeyUpdated represents a DkimKeyUpdated event raised by the DkimKeyOracle contract.
type DkimKeyOracleDkimKeyUpdated struct {
	KeyHash    [32]byte
	DomainHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDkimKeyUpdated is a free log retrieval operation binding the contract event 0xe780335b851f95261f942f1897645485db8bd8c5ff0fc4dff7b5cc2de72d2d46.
//
// Solidity: event DkimKeyUpdated(bytes32 keyHash, bytes32 domainHash)
func (_DkimKeyOracle *DkimKeyOracleFilterer) FilterDkimKeyUpdated(opts *bind.FilterOpts) (*DkimKeyOracleDkimKeyUpdatedIterator, error) {

	logs, sub, err := _DkimKeyOracle.contract.FilterLogs(opts, "DkimKeyUpdated")
	if err != nil {
		return nil, err
	}
	return &DkimKeyOracleDkimKeyUpdatedIterator{contract: _DkimKeyOracle.contract, event: "DkimKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchDkimKeyUpdated is a free log subscription operation binding the contract event 0xe780335b851f95261f942f1897645485db8bd8c5ff0fc4dff7b5cc2de72d2d46.
//
// Solidity: event DkimKeyUpdated(bytes32 keyHash, bytes32 domainHash)
func (_DkimKeyOracle *DkimKeyOracleFilterer) WatchDkimKeyUpdated(opts *bind.WatchOpts, sink chan<- *DkimKeyOracleDkimKeyUpdated) (event.Subscription, error) {

	logs, sub, err := _DkimKeyOracle.contract.WatchLogs(opts, "DkimKeyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DkimKeyOracleDkimKeyUpdated)
				if err := _DkimKeyOracle.contract.UnpackLog(event, "DkimKeyUpdated", log); err != nil {
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

// ParseDkimKeyUpdated is a log parse operation binding the contract event 0xe780335b851f95261f942f1897645485db8bd8c5ff0fc4dff7b5cc2de72d2d46.
//
// Solidity: event DkimKeyUpdated(bytes32 keyHash, bytes32 domainHash)
func (_DkimKeyOracle *DkimKeyOracleFilterer) ParseDkimKeyUpdated(log types.Log) (*DkimKeyOracleDkimKeyUpdated, error) {
	event := new(DkimKeyOracleDkimKeyUpdated)
	if err := _DkimKeyOracle.contract.UnpackLog(event, "DkimKeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DkimKeyOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DkimKeyOracle contract.
type DkimKeyOracleOwnershipTransferredIterator struct {
	Event *DkimKeyOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DkimKeyOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DkimKeyOracleOwnershipTransferred)
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
		it.Event = new(DkimKeyOracleOwnershipTransferred)
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
func (it *DkimKeyOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DkimKeyOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DkimKeyOracleOwnershipTransferred represents a OwnershipTransferred event raised by the DkimKeyOracle contract.
type DkimKeyOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DkimKeyOracle *DkimKeyOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DkimKeyOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DkimKeyOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DkimKeyOracleOwnershipTransferredIterator{contract: _DkimKeyOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DkimKeyOracle *DkimKeyOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DkimKeyOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DkimKeyOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DkimKeyOracleOwnershipTransferred)
				if err := _DkimKeyOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DkimKeyOracle *DkimKeyOracleFilterer) ParseOwnershipTransferred(log types.Log) (*DkimKeyOracleOwnershipTransferred, error) {
	event := new(DkimKeyOracleOwnershipTransferred)
	if err := _DkimKeyOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DkimKeyOraclePauserUpdatedIterator is returned from FilterPauserUpdated and is used to iterate over the raw logs and unpacked data for PauserUpdated events raised by the DkimKeyOracle contract.
type DkimKeyOraclePauserUpdatedIterator struct {
	Event *DkimKeyOraclePauserUpdated // Event containing the contract specifics and raw log

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
func (it *DkimKeyOraclePauserUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DkimKeyOraclePauserUpdated)
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
		it.Event = new(DkimKeyOraclePauserUpdated)
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
func (it *DkimKeyOraclePauserUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DkimKeyOraclePauserUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DkimKeyOraclePauserUpdated represents a PauserUpdated event raised by the DkimKeyOracle contract.
type DkimKeyOraclePauserUpdated struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPauserUpdated is a free log retrieval operation binding the contract event 0xa4336c0cb1e245b95ad204faed7e940d6dc999684fd8b5e1ff597a0c4efca8ab.
//
// Solidity: event PauserUpdated(address pauser)
func (_DkimKeyOracle *DkimKeyOracleFilterer) FilterPauserUpdated(opts *bind.FilterOpts) (*DkimKeyOraclePauserUpdatedIterator, error) {

	logs, sub, err := _DkimKeyOracle.contract.FilterLogs(opts, "PauserUpdated")
	if err != nil {
		return nil, err
	}
	return &DkimKeyOraclePauserUpdatedIterator{contract: _DkimKeyOracle.contract, event: "PauserUpdated", logs: logs, sub: sub}, nil
}

// WatchPauserUpdated is a free log subscription operation binding the contract event 0xa4336c0cb1e245b95ad204faed7e940d6dc999684fd8b5e1ff597a0c4efca8ab.
//
// Solidity: event PauserUpdated(address pauser)
func (_DkimKeyOracle *DkimKeyOracleFilterer) WatchPauserUpdated(opts *bind.WatchOpts, sink chan<- *DkimKeyOraclePauserUpdated) (event.Subscription, error) {

	logs, sub, err := _DkimKeyOracle.contract.WatchLogs(opts, "PauserUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DkimKeyOraclePauserUpdated)
				if err := _DkimKeyOracle.contract.UnpackLog(event, "PauserUpdated", log); err != nil {
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

// ParsePauserUpdated is a log parse operation binding the contract event 0xa4336c0cb1e245b95ad204faed7e940d6dc999684fd8b5e1ff597a0c4efca8ab.
//
// Solidity: event PauserUpdated(address pauser)
func (_DkimKeyOracle *DkimKeyOracleFilterer) ParsePauserUpdated(log types.Log) (*DkimKeyOraclePauserUpdated, error) {
	event := new(DkimKeyOraclePauserUpdated)
	if err := _DkimKeyOracle.contract.UnpackLog(event, "PauserUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
