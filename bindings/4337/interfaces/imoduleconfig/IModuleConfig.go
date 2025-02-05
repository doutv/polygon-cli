// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package imoduleconfig

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

// IModuleConfigMetaData contains all meta data concerning the IModuleConfig contract.
var IModuleConfigMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ModuleInstalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ModuleUninstalled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"installModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"additionalContext\",\"type\":\"bytes\"}],\"name\":\"isModuleInstalled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"}],\"name\":\"supportsModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"deInitData\",\"type\":\"bytes\"}],\"name\":\"uninstallModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IModuleConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use IModuleConfigMetaData.ABI instead.
var IModuleConfigABI = IModuleConfigMetaData.ABI

// IModuleConfig is an auto generated Go binding around an Ethereum contract.
type IModuleConfig struct {
	IModuleConfigCaller     // Read-only binding to the contract
	IModuleConfigTransactor // Write-only binding to the contract
	IModuleConfigFilterer   // Log filterer for contract events
}

// IModuleConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type IModuleConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IModuleConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IModuleConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IModuleConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IModuleConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IModuleConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IModuleConfigSession struct {
	Contract     *IModuleConfig    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IModuleConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IModuleConfigCallerSession struct {
	Contract *IModuleConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IModuleConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IModuleConfigTransactorSession struct {
	Contract     *IModuleConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IModuleConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type IModuleConfigRaw struct {
	Contract *IModuleConfig // Generic contract binding to access the raw methods on
}

// IModuleConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IModuleConfigCallerRaw struct {
	Contract *IModuleConfigCaller // Generic read-only contract binding to access the raw methods on
}

// IModuleConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IModuleConfigTransactorRaw struct {
	Contract *IModuleConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIModuleConfig creates a new instance of IModuleConfig, bound to a specific deployed contract.
func NewIModuleConfig(address common.Address, backend bind.ContractBackend) (*IModuleConfig, error) {
	contract, err := bindIModuleConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IModuleConfig{IModuleConfigCaller: IModuleConfigCaller{contract: contract}, IModuleConfigTransactor: IModuleConfigTransactor{contract: contract}, IModuleConfigFilterer: IModuleConfigFilterer{contract: contract}}, nil
}

// NewIModuleConfigCaller creates a new read-only instance of IModuleConfig, bound to a specific deployed contract.
func NewIModuleConfigCaller(address common.Address, caller bind.ContractCaller) (*IModuleConfigCaller, error) {
	contract, err := bindIModuleConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IModuleConfigCaller{contract: contract}, nil
}

// NewIModuleConfigTransactor creates a new write-only instance of IModuleConfig, bound to a specific deployed contract.
func NewIModuleConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*IModuleConfigTransactor, error) {
	contract, err := bindIModuleConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IModuleConfigTransactor{contract: contract}, nil
}

// NewIModuleConfigFilterer creates a new log filterer instance of IModuleConfig, bound to a specific deployed contract.
func NewIModuleConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*IModuleConfigFilterer, error) {
	contract, err := bindIModuleConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IModuleConfigFilterer{contract: contract}, nil
}

// bindIModuleConfig binds a generic wrapper to an already deployed contract.
func bindIModuleConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IModuleConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IModuleConfig *IModuleConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IModuleConfig.Contract.IModuleConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IModuleConfig *IModuleConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IModuleConfig.Contract.IModuleConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IModuleConfig *IModuleConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IModuleConfig.Contract.IModuleConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IModuleConfig *IModuleConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IModuleConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IModuleConfig *IModuleConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IModuleConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IModuleConfig *IModuleConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IModuleConfig.Contract.contract.Transact(opts, method, params...)
}

// IsModuleInstalled is a free data retrieval call binding the contract method 0x112d3a7d.
//
// Solidity: function isModuleInstalled(uint256 moduleTypeId, address module, bytes additionalContext) view returns(bool)
func (_IModuleConfig *IModuleConfigCaller) IsModuleInstalled(opts *bind.CallOpts, moduleTypeId *big.Int, module common.Address, additionalContext []byte) (bool, error) {
	var out []interface{}
	err := _IModuleConfig.contract.Call(opts, &out, "isModuleInstalled", moduleTypeId, module, additionalContext)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleInstalled is a free data retrieval call binding the contract method 0x112d3a7d.
//
// Solidity: function isModuleInstalled(uint256 moduleTypeId, address module, bytes additionalContext) view returns(bool)
func (_IModuleConfig *IModuleConfigSession) IsModuleInstalled(moduleTypeId *big.Int, module common.Address, additionalContext []byte) (bool, error) {
	return _IModuleConfig.Contract.IsModuleInstalled(&_IModuleConfig.CallOpts, moduleTypeId, module, additionalContext)
}

// IsModuleInstalled is a free data retrieval call binding the contract method 0x112d3a7d.
//
// Solidity: function isModuleInstalled(uint256 moduleTypeId, address module, bytes additionalContext) view returns(bool)
func (_IModuleConfig *IModuleConfigCallerSession) IsModuleInstalled(moduleTypeId *big.Int, module common.Address, additionalContext []byte) (bool, error) {
	return _IModuleConfig.Contract.IsModuleInstalled(&_IModuleConfig.CallOpts, moduleTypeId, module, additionalContext)
}

// SupportsModule is a free data retrieval call binding the contract method 0xf2dc691d.
//
// Solidity: function supportsModule(uint256 moduleTypeId) view returns(bool)
func (_IModuleConfig *IModuleConfigCaller) SupportsModule(opts *bind.CallOpts, moduleTypeId *big.Int) (bool, error) {
	var out []interface{}
	err := _IModuleConfig.contract.Call(opts, &out, "supportsModule", moduleTypeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsModule is a free data retrieval call binding the contract method 0xf2dc691d.
//
// Solidity: function supportsModule(uint256 moduleTypeId) view returns(bool)
func (_IModuleConfig *IModuleConfigSession) SupportsModule(moduleTypeId *big.Int) (bool, error) {
	return _IModuleConfig.Contract.SupportsModule(&_IModuleConfig.CallOpts, moduleTypeId)
}

// SupportsModule is a free data retrieval call binding the contract method 0xf2dc691d.
//
// Solidity: function supportsModule(uint256 moduleTypeId) view returns(bool)
func (_IModuleConfig *IModuleConfigCallerSession) SupportsModule(moduleTypeId *big.Int) (bool, error) {
	return _IModuleConfig.Contract.SupportsModule(&_IModuleConfig.CallOpts, moduleTypeId)
}

// InstallModule is a paid mutator transaction binding the contract method 0x9517e29f.
//
// Solidity: function installModule(uint256 moduleTypeId, address module, bytes initData) returns()
func (_IModuleConfig *IModuleConfigTransactor) InstallModule(opts *bind.TransactOpts, moduleTypeId *big.Int, module common.Address, initData []byte) (*types.Transaction, error) {
	return _IModuleConfig.contract.Transact(opts, "installModule", moduleTypeId, module, initData)
}

// InstallModule is a paid mutator transaction binding the contract method 0x9517e29f.
//
// Solidity: function installModule(uint256 moduleTypeId, address module, bytes initData) returns()
func (_IModuleConfig *IModuleConfigSession) InstallModule(moduleTypeId *big.Int, module common.Address, initData []byte) (*types.Transaction, error) {
	return _IModuleConfig.Contract.InstallModule(&_IModuleConfig.TransactOpts, moduleTypeId, module, initData)
}

// InstallModule is a paid mutator transaction binding the contract method 0x9517e29f.
//
// Solidity: function installModule(uint256 moduleTypeId, address module, bytes initData) returns()
func (_IModuleConfig *IModuleConfigTransactorSession) InstallModule(moduleTypeId *big.Int, module common.Address, initData []byte) (*types.Transaction, error) {
	return _IModuleConfig.Contract.InstallModule(&_IModuleConfig.TransactOpts, moduleTypeId, module, initData)
}

// UninstallModule is a paid mutator transaction binding the contract method 0xa71763a8.
//
// Solidity: function uninstallModule(uint256 moduleTypeId, address module, bytes deInitData) returns()
func (_IModuleConfig *IModuleConfigTransactor) UninstallModule(opts *bind.TransactOpts, moduleTypeId *big.Int, module common.Address, deInitData []byte) (*types.Transaction, error) {
	return _IModuleConfig.contract.Transact(opts, "uninstallModule", moduleTypeId, module, deInitData)
}

// UninstallModule is a paid mutator transaction binding the contract method 0xa71763a8.
//
// Solidity: function uninstallModule(uint256 moduleTypeId, address module, bytes deInitData) returns()
func (_IModuleConfig *IModuleConfigSession) UninstallModule(moduleTypeId *big.Int, module common.Address, deInitData []byte) (*types.Transaction, error) {
	return _IModuleConfig.Contract.UninstallModule(&_IModuleConfig.TransactOpts, moduleTypeId, module, deInitData)
}

// UninstallModule is a paid mutator transaction binding the contract method 0xa71763a8.
//
// Solidity: function uninstallModule(uint256 moduleTypeId, address module, bytes deInitData) returns()
func (_IModuleConfig *IModuleConfigTransactorSession) UninstallModule(moduleTypeId *big.Int, module common.Address, deInitData []byte) (*types.Transaction, error) {
	return _IModuleConfig.Contract.UninstallModule(&_IModuleConfig.TransactOpts, moduleTypeId, module, deInitData)
}

// IModuleConfigModuleInstalledIterator is returned from FilterModuleInstalled and is used to iterate over the raw logs and unpacked data for ModuleInstalled events raised by the IModuleConfig contract.
type IModuleConfigModuleInstalledIterator struct {
	Event *IModuleConfigModuleInstalled // Event containing the contract specifics and raw log

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
func (it *IModuleConfigModuleInstalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IModuleConfigModuleInstalled)
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
		it.Event = new(IModuleConfigModuleInstalled)
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
func (it *IModuleConfigModuleInstalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IModuleConfigModuleInstalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IModuleConfigModuleInstalled represents a ModuleInstalled event raised by the IModuleConfig contract.
type IModuleConfigModuleInstalled struct {
	ModuleTypeId *big.Int
	Module       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterModuleInstalled is a free log retrieval operation binding the contract event 0xd21d0b289f126c4b473ea641963e766833c2f13866e4ff480abd787c100ef123.
//
// Solidity: event ModuleInstalled(uint256 moduleTypeId, address module)
func (_IModuleConfig *IModuleConfigFilterer) FilterModuleInstalled(opts *bind.FilterOpts) (*IModuleConfigModuleInstalledIterator, error) {

	logs, sub, err := _IModuleConfig.contract.FilterLogs(opts, "ModuleInstalled")
	if err != nil {
		return nil, err
	}
	return &IModuleConfigModuleInstalledIterator{contract: _IModuleConfig.contract, event: "ModuleInstalled", logs: logs, sub: sub}, nil
}

// WatchModuleInstalled is a free log subscription operation binding the contract event 0xd21d0b289f126c4b473ea641963e766833c2f13866e4ff480abd787c100ef123.
//
// Solidity: event ModuleInstalled(uint256 moduleTypeId, address module)
func (_IModuleConfig *IModuleConfigFilterer) WatchModuleInstalled(opts *bind.WatchOpts, sink chan<- *IModuleConfigModuleInstalled) (event.Subscription, error) {

	logs, sub, err := _IModuleConfig.contract.WatchLogs(opts, "ModuleInstalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IModuleConfigModuleInstalled)
				if err := _IModuleConfig.contract.UnpackLog(event, "ModuleInstalled", log); err != nil {
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

// ParseModuleInstalled is a log parse operation binding the contract event 0xd21d0b289f126c4b473ea641963e766833c2f13866e4ff480abd787c100ef123.
//
// Solidity: event ModuleInstalled(uint256 moduleTypeId, address module)
func (_IModuleConfig *IModuleConfigFilterer) ParseModuleInstalled(log types.Log) (*IModuleConfigModuleInstalled, error) {
	event := new(IModuleConfigModuleInstalled)
	if err := _IModuleConfig.contract.UnpackLog(event, "ModuleInstalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IModuleConfigModuleUninstalledIterator is returned from FilterModuleUninstalled and is used to iterate over the raw logs and unpacked data for ModuleUninstalled events raised by the IModuleConfig contract.
type IModuleConfigModuleUninstalledIterator struct {
	Event *IModuleConfigModuleUninstalled // Event containing the contract specifics and raw log

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
func (it *IModuleConfigModuleUninstalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IModuleConfigModuleUninstalled)
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
		it.Event = new(IModuleConfigModuleUninstalled)
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
func (it *IModuleConfigModuleUninstalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IModuleConfigModuleUninstalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IModuleConfigModuleUninstalled represents a ModuleUninstalled event raised by the IModuleConfig contract.
type IModuleConfigModuleUninstalled struct {
	ModuleTypeId *big.Int
	Module       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterModuleUninstalled is a free log retrieval operation binding the contract event 0x341347516a9de374859dfda710fa4828b2d48cb57d4fbe4c1149612b8e02276e.
//
// Solidity: event ModuleUninstalled(uint256 moduleTypeId, address module)
func (_IModuleConfig *IModuleConfigFilterer) FilterModuleUninstalled(opts *bind.FilterOpts) (*IModuleConfigModuleUninstalledIterator, error) {

	logs, sub, err := _IModuleConfig.contract.FilterLogs(opts, "ModuleUninstalled")
	if err != nil {
		return nil, err
	}
	return &IModuleConfigModuleUninstalledIterator{contract: _IModuleConfig.contract, event: "ModuleUninstalled", logs: logs, sub: sub}, nil
}

// WatchModuleUninstalled is a free log subscription operation binding the contract event 0x341347516a9de374859dfda710fa4828b2d48cb57d4fbe4c1149612b8e02276e.
//
// Solidity: event ModuleUninstalled(uint256 moduleTypeId, address module)
func (_IModuleConfig *IModuleConfigFilterer) WatchModuleUninstalled(opts *bind.WatchOpts, sink chan<- *IModuleConfigModuleUninstalled) (event.Subscription, error) {

	logs, sub, err := _IModuleConfig.contract.WatchLogs(opts, "ModuleUninstalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IModuleConfigModuleUninstalled)
				if err := _IModuleConfig.contract.UnpackLog(event, "ModuleUninstalled", log); err != nil {
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

// ParseModuleUninstalled is a log parse operation binding the contract event 0x341347516a9de374859dfda710fa4828b2d48cb57d4fbe4c1149612b8e02276e.
//
// Solidity: event ModuleUninstalled(uint256 moduleTypeId, address module)
func (_IModuleConfig *IModuleConfigFilterer) ParseModuleUninstalled(log types.Log) (*IModuleConfigModuleUninstalled, error) {
	event := new(IModuleConfigModuleUninstalled)
	if err := _IModuleConfig.contract.UnpackLog(event, "ModuleUninstalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
