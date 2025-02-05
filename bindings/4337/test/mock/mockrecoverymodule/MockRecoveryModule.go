// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockrecoverymodule

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

// MockRecoveryModuleMetaData contains all meta data concerning the MockRecoveryModule contract.
var MockRecoveryModuleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"OnInstallCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"OnUninstallCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"RecoveryCallbackCalled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"claimRecoveryFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"}],\"name\":\"isModuleType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"recoverByTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523460155761048c908161001b8239f35b600080fdfe60806040908082526004908136101561001757600080fd5b600090813560e01c90816317586005146102a257508063379f4e66146102315780636d61fe70146101da5780638a91b0e314610168578063ad809f69146100895763ecd059611461006757600080fd5b3461008657602036600319011261008657506002602092519135148152f35b80fd5b5080918334610164576060366003190112610164576100a661031f565b60443567ffffffffffffffff8111610160576100c86100d0913690860161033a565b8101906103ae565b93909160018060a01b0380918551966020880152602435868801526001600160801b0360608801526060875261010587610392565b1690813b1561015c57868094610130875198899687958694631bcfa73360e11b8652169084016103d4565b03925af190811561015357506101435750f35b61014c90610368565b6100865780f35b513d84823e3d90fd5b8680fd5b8480fd5b5050fd5b509190346101d65760203660031901126101d65780359167ffffffffffffffff83116101d2576101bf6101cc927fcc464038b9ad6da1fb7abf8926124d82bccae955e2c90296634bd1707e3b19b19436910161033a565b929091519283928361042e565b0390a180f35b8380fd5b8280fd5b509190346101d65760203660031901126101d65780359167ffffffffffffffff83116101d2576101bf6101cc927fe4ef1393240978463584a382073a306d63f735fc24f5446f0df70b5a8b1bbdb49436910161033a565b508091833461016457806003193601126101645761024d61031f565b60243567ffffffffffffffff8111610160576100c861026f913690860161033a565b93909160018060a01b038091855196602088015242868801526001600160801b0360608801526060875261010587610392565b90508382843461031b57606036600319011261031b576102c061031f565b6001600160a01b036024358181169290839003610160571690813b156101d257836044928794928593631f58bb2d60e11b8552840152833560248401525af1908115610153575061030f575080f35b61031890610368565b80f35b5080fd5b600435906001600160a01b038216820361033557565b600080fd5b9181601f840112156103355782359167ffffffffffffffff8311610335576020838186019501011161033557565b67ffffffffffffffff811161037c57604052565b634e487b7160e01b600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761037c57604052565b91908260409103126103355781356001600160a01b038116810361033557916020013590565b6001600160a01b039091168152604060208083018290528351918301829052939260005b82811061041a57505060609293506000838284010152601f8019910116010190565b8181018601518482016060015285016103f8565b90918060409360208452816020850152848401376000828201840152601f01601f191601019056fea2646970667358221220085b17848688e2bd352428afb85a32f5cf3fa499c7134cca0b7a044f22a117dd64736f6c63430008190033",
}

// MockRecoveryModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use MockRecoveryModuleMetaData.ABI instead.
var MockRecoveryModuleABI = MockRecoveryModuleMetaData.ABI

// MockRecoveryModuleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockRecoveryModuleMetaData.Bin instead.
var MockRecoveryModuleBin = MockRecoveryModuleMetaData.Bin

// DeployMockRecoveryModule deploys a new Ethereum contract, binding an instance of MockRecoveryModule to it.
func DeployMockRecoveryModule(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockRecoveryModule, error) {
	parsed, err := MockRecoveryModuleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockRecoveryModuleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockRecoveryModule{MockRecoveryModuleCaller: MockRecoveryModuleCaller{contract: contract}, MockRecoveryModuleTransactor: MockRecoveryModuleTransactor{contract: contract}, MockRecoveryModuleFilterer: MockRecoveryModuleFilterer{contract: contract}}, nil
}

// MockRecoveryModule is an auto generated Go binding around an Ethereum contract.
type MockRecoveryModule struct {
	MockRecoveryModuleCaller     // Read-only binding to the contract
	MockRecoveryModuleTransactor // Write-only binding to the contract
	MockRecoveryModuleFilterer   // Log filterer for contract events
}

// MockRecoveryModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockRecoveryModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockRecoveryModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockRecoveryModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockRecoveryModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockRecoveryModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockRecoveryModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockRecoveryModuleSession struct {
	Contract     *MockRecoveryModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MockRecoveryModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockRecoveryModuleCallerSession struct {
	Contract *MockRecoveryModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MockRecoveryModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockRecoveryModuleTransactorSession struct {
	Contract     *MockRecoveryModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MockRecoveryModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockRecoveryModuleRaw struct {
	Contract *MockRecoveryModule // Generic contract binding to access the raw methods on
}

// MockRecoveryModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockRecoveryModuleCallerRaw struct {
	Contract *MockRecoveryModuleCaller // Generic read-only contract binding to access the raw methods on
}

// MockRecoveryModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockRecoveryModuleTransactorRaw struct {
	Contract *MockRecoveryModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockRecoveryModule creates a new instance of MockRecoveryModule, bound to a specific deployed contract.
func NewMockRecoveryModule(address common.Address, backend bind.ContractBackend) (*MockRecoveryModule, error) {
	contract, err := bindMockRecoveryModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockRecoveryModule{MockRecoveryModuleCaller: MockRecoveryModuleCaller{contract: contract}, MockRecoveryModuleTransactor: MockRecoveryModuleTransactor{contract: contract}, MockRecoveryModuleFilterer: MockRecoveryModuleFilterer{contract: contract}}, nil
}

// NewMockRecoveryModuleCaller creates a new read-only instance of MockRecoveryModule, bound to a specific deployed contract.
func NewMockRecoveryModuleCaller(address common.Address, caller bind.ContractCaller) (*MockRecoveryModuleCaller, error) {
	contract, err := bindMockRecoveryModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockRecoveryModuleCaller{contract: contract}, nil
}

// NewMockRecoveryModuleTransactor creates a new write-only instance of MockRecoveryModule, bound to a specific deployed contract.
func NewMockRecoveryModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*MockRecoveryModuleTransactor, error) {
	contract, err := bindMockRecoveryModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockRecoveryModuleTransactor{contract: contract}, nil
}

// NewMockRecoveryModuleFilterer creates a new log filterer instance of MockRecoveryModule, bound to a specific deployed contract.
func NewMockRecoveryModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*MockRecoveryModuleFilterer, error) {
	contract, err := bindMockRecoveryModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockRecoveryModuleFilterer{contract: contract}, nil
}

// bindMockRecoveryModule binds a generic wrapper to an already deployed contract.
func bindMockRecoveryModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockRecoveryModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockRecoveryModule *MockRecoveryModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockRecoveryModule.Contract.MockRecoveryModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockRecoveryModule *MockRecoveryModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.MockRecoveryModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockRecoveryModule *MockRecoveryModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.MockRecoveryModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockRecoveryModule *MockRecoveryModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockRecoveryModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockRecoveryModule *MockRecoveryModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockRecoveryModule *MockRecoveryModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.contract.Transact(opts, method, params...)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_MockRecoveryModule *MockRecoveryModuleCaller) IsModuleType(opts *bind.CallOpts, moduleTypeId *big.Int) (bool, error) {
	var out []interface{}
	err := _MockRecoveryModule.contract.Call(opts, &out, "isModuleType", moduleTypeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_MockRecoveryModule *MockRecoveryModuleSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _MockRecoveryModule.Contract.IsModuleType(&_MockRecoveryModule.CallOpts, moduleTypeId)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_MockRecoveryModule *MockRecoveryModuleCallerSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _MockRecoveryModule.Contract.IsModuleType(&_MockRecoveryModule.CallOpts, moduleTypeId)
}

// ClaimRecoveryFee is a paid mutator transaction binding the contract method 0x17586005.
//
// Solidity: function claimRecoveryFee(address _account, address _receiver, uint256 _value) returns()
func (_MockRecoveryModule *MockRecoveryModuleTransactor) ClaimRecoveryFee(opts *bind.TransactOpts, _account common.Address, _receiver common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MockRecoveryModule.contract.Transact(opts, "claimRecoveryFee", _account, _receiver, _value)
}

// ClaimRecoveryFee is a paid mutator transaction binding the contract method 0x17586005.
//
// Solidity: function claimRecoveryFee(address _account, address _receiver, uint256 _value) returns()
func (_MockRecoveryModule *MockRecoveryModuleSession) ClaimRecoveryFee(_account common.Address, _receiver common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.ClaimRecoveryFee(&_MockRecoveryModule.TransactOpts, _account, _receiver, _value)
}

// ClaimRecoveryFee is a paid mutator transaction binding the contract method 0x17586005.
//
// Solidity: function claimRecoveryFee(address _account, address _receiver, uint256 _value) returns()
func (_MockRecoveryModule *MockRecoveryModuleTransactorSession) ClaimRecoveryFee(_account common.Address, _receiver common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.ClaimRecoveryFee(&_MockRecoveryModule.TransactOpts, _account, _receiver, _value)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_MockRecoveryModule *MockRecoveryModuleTransactor) OnInstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _MockRecoveryModule.contract.Transact(opts, "onInstall", data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_MockRecoveryModule *MockRecoveryModuleSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.OnInstall(&_MockRecoveryModule.TransactOpts, data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_MockRecoveryModule *MockRecoveryModuleTransactorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.OnInstall(&_MockRecoveryModule.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_MockRecoveryModule *MockRecoveryModuleTransactor) OnUninstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _MockRecoveryModule.contract.Transact(opts, "onUninstall", data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_MockRecoveryModule *MockRecoveryModuleSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.OnUninstall(&_MockRecoveryModule.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_MockRecoveryModule *MockRecoveryModuleTransactorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.OnUninstall(&_MockRecoveryModule.TransactOpts, data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address _account, bytes _data) returns()
func (_MockRecoveryModule *MockRecoveryModuleTransactor) Recover(opts *bind.TransactOpts, _account common.Address, _data []byte) (*types.Transaction, error) {
	return _MockRecoveryModule.contract.Transact(opts, "recover", _account, _data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address _account, bytes _data) returns()
func (_MockRecoveryModule *MockRecoveryModuleSession) Recover(_account common.Address, _data []byte) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.Recover(&_MockRecoveryModule.TransactOpts, _account, _data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address _account, bytes _data) returns()
func (_MockRecoveryModule *MockRecoveryModuleTransactorSession) Recover(_account common.Address, _data []byte) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.Recover(&_MockRecoveryModule.TransactOpts, _account, _data)
}

// RecoverByTimestamp is a paid mutator transaction binding the contract method 0xad809f69.
//
// Solidity: function recoverByTimestamp(address _account, uint256 _timestamp, bytes _data) returns()
func (_MockRecoveryModule *MockRecoveryModuleTransactor) RecoverByTimestamp(opts *bind.TransactOpts, _account common.Address, _timestamp *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockRecoveryModule.contract.Transact(opts, "recoverByTimestamp", _account, _timestamp, _data)
}

// RecoverByTimestamp is a paid mutator transaction binding the contract method 0xad809f69.
//
// Solidity: function recoverByTimestamp(address _account, uint256 _timestamp, bytes _data) returns()
func (_MockRecoveryModule *MockRecoveryModuleSession) RecoverByTimestamp(_account common.Address, _timestamp *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.RecoverByTimestamp(&_MockRecoveryModule.TransactOpts, _account, _timestamp, _data)
}

// RecoverByTimestamp is a paid mutator transaction binding the contract method 0xad809f69.
//
// Solidity: function recoverByTimestamp(address _account, uint256 _timestamp, bytes _data) returns()
func (_MockRecoveryModule *MockRecoveryModuleTransactorSession) RecoverByTimestamp(_account common.Address, _timestamp *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockRecoveryModule.Contract.RecoverByTimestamp(&_MockRecoveryModule.TransactOpts, _account, _timestamp, _data)
}

// MockRecoveryModuleOnInstallCalledIterator is returned from FilterOnInstallCalled and is used to iterate over the raw logs and unpacked data for OnInstallCalled events raised by the MockRecoveryModule contract.
type MockRecoveryModuleOnInstallCalledIterator struct {
	Event *MockRecoveryModuleOnInstallCalled // Event containing the contract specifics and raw log

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
func (it *MockRecoveryModuleOnInstallCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockRecoveryModuleOnInstallCalled)
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
		it.Event = new(MockRecoveryModuleOnInstallCalled)
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
func (it *MockRecoveryModuleOnInstallCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockRecoveryModuleOnInstallCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockRecoveryModuleOnInstallCalled represents a OnInstallCalled event raised by the MockRecoveryModule contract.
type MockRecoveryModuleOnInstallCalled struct {
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnInstallCalled is a free log retrieval operation binding the contract event 0xe4ef1393240978463584a382073a306d63f735fc24f5446f0df70b5a8b1bbdb4.
//
// Solidity: event OnInstallCalled(bytes _data)
func (_MockRecoveryModule *MockRecoveryModuleFilterer) FilterOnInstallCalled(opts *bind.FilterOpts) (*MockRecoveryModuleOnInstallCalledIterator, error) {

	logs, sub, err := _MockRecoveryModule.contract.FilterLogs(opts, "OnInstallCalled")
	if err != nil {
		return nil, err
	}
	return &MockRecoveryModuleOnInstallCalledIterator{contract: _MockRecoveryModule.contract, event: "OnInstallCalled", logs: logs, sub: sub}, nil
}

// WatchOnInstallCalled is a free log subscription operation binding the contract event 0xe4ef1393240978463584a382073a306d63f735fc24f5446f0df70b5a8b1bbdb4.
//
// Solidity: event OnInstallCalled(bytes _data)
func (_MockRecoveryModule *MockRecoveryModuleFilterer) WatchOnInstallCalled(opts *bind.WatchOpts, sink chan<- *MockRecoveryModuleOnInstallCalled) (event.Subscription, error) {

	logs, sub, err := _MockRecoveryModule.contract.WatchLogs(opts, "OnInstallCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockRecoveryModuleOnInstallCalled)
				if err := _MockRecoveryModule.contract.UnpackLog(event, "OnInstallCalled", log); err != nil {
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

// ParseOnInstallCalled is a log parse operation binding the contract event 0xe4ef1393240978463584a382073a306d63f735fc24f5446f0df70b5a8b1bbdb4.
//
// Solidity: event OnInstallCalled(bytes _data)
func (_MockRecoveryModule *MockRecoveryModuleFilterer) ParseOnInstallCalled(log types.Log) (*MockRecoveryModuleOnInstallCalled, error) {
	event := new(MockRecoveryModuleOnInstallCalled)
	if err := _MockRecoveryModule.contract.UnpackLog(event, "OnInstallCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockRecoveryModuleOnUninstallCalledIterator is returned from FilterOnUninstallCalled and is used to iterate over the raw logs and unpacked data for OnUninstallCalled events raised by the MockRecoveryModule contract.
type MockRecoveryModuleOnUninstallCalledIterator struct {
	Event *MockRecoveryModuleOnUninstallCalled // Event containing the contract specifics and raw log

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
func (it *MockRecoveryModuleOnUninstallCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockRecoveryModuleOnUninstallCalled)
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
		it.Event = new(MockRecoveryModuleOnUninstallCalled)
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
func (it *MockRecoveryModuleOnUninstallCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockRecoveryModuleOnUninstallCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockRecoveryModuleOnUninstallCalled represents a OnUninstallCalled event raised by the MockRecoveryModule contract.
type MockRecoveryModuleOnUninstallCalled struct {
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnUninstallCalled is a free log retrieval operation binding the contract event 0xcc464038b9ad6da1fb7abf8926124d82bccae955e2c90296634bd1707e3b19b1.
//
// Solidity: event OnUninstallCalled(bytes _data)
func (_MockRecoveryModule *MockRecoveryModuleFilterer) FilterOnUninstallCalled(opts *bind.FilterOpts) (*MockRecoveryModuleOnUninstallCalledIterator, error) {

	logs, sub, err := _MockRecoveryModule.contract.FilterLogs(opts, "OnUninstallCalled")
	if err != nil {
		return nil, err
	}
	return &MockRecoveryModuleOnUninstallCalledIterator{contract: _MockRecoveryModule.contract, event: "OnUninstallCalled", logs: logs, sub: sub}, nil
}

// WatchOnUninstallCalled is a free log subscription operation binding the contract event 0xcc464038b9ad6da1fb7abf8926124d82bccae955e2c90296634bd1707e3b19b1.
//
// Solidity: event OnUninstallCalled(bytes _data)
func (_MockRecoveryModule *MockRecoveryModuleFilterer) WatchOnUninstallCalled(opts *bind.WatchOpts, sink chan<- *MockRecoveryModuleOnUninstallCalled) (event.Subscription, error) {

	logs, sub, err := _MockRecoveryModule.contract.WatchLogs(opts, "OnUninstallCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockRecoveryModuleOnUninstallCalled)
				if err := _MockRecoveryModule.contract.UnpackLog(event, "OnUninstallCalled", log); err != nil {
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

// ParseOnUninstallCalled is a log parse operation binding the contract event 0xcc464038b9ad6da1fb7abf8926124d82bccae955e2c90296634bd1707e3b19b1.
//
// Solidity: event OnUninstallCalled(bytes _data)
func (_MockRecoveryModule *MockRecoveryModuleFilterer) ParseOnUninstallCalled(log types.Log) (*MockRecoveryModuleOnUninstallCalled, error) {
	event := new(MockRecoveryModuleOnUninstallCalled)
	if err := _MockRecoveryModule.contract.UnpackLog(event, "OnUninstallCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockRecoveryModuleRecoveryCallbackCalledIterator is returned from FilterRecoveryCallbackCalled and is used to iterate over the raw logs and unpacked data for RecoveryCallbackCalled events raised by the MockRecoveryModule contract.
type MockRecoveryModuleRecoveryCallbackCalledIterator struct {
	Event *MockRecoveryModuleRecoveryCallbackCalled // Event containing the contract specifics and raw log

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
func (it *MockRecoveryModuleRecoveryCallbackCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockRecoveryModuleRecoveryCallbackCalled)
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
		it.Event = new(MockRecoveryModuleRecoveryCallbackCalled)
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
func (it *MockRecoveryModuleRecoveryCallbackCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockRecoveryModuleRecoveryCallbackCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockRecoveryModuleRecoveryCallbackCalled represents a RecoveryCallbackCalled event raised by the MockRecoveryModule contract.
type MockRecoveryModuleRecoveryCallbackCalled struct {
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRecoveryCallbackCalled is a free log retrieval operation binding the contract event 0xa553e8c54a105e80df2e712fd2f1436d50d6f654e10153a072da3164214d9f0d.
//
// Solidity: event RecoveryCallbackCalled(bytes _data)
func (_MockRecoveryModule *MockRecoveryModuleFilterer) FilterRecoveryCallbackCalled(opts *bind.FilterOpts) (*MockRecoveryModuleRecoveryCallbackCalledIterator, error) {

	logs, sub, err := _MockRecoveryModule.contract.FilterLogs(opts, "RecoveryCallbackCalled")
	if err != nil {
		return nil, err
	}
	return &MockRecoveryModuleRecoveryCallbackCalledIterator{contract: _MockRecoveryModule.contract, event: "RecoveryCallbackCalled", logs: logs, sub: sub}, nil
}

// WatchRecoveryCallbackCalled is a free log subscription operation binding the contract event 0xa553e8c54a105e80df2e712fd2f1436d50d6f654e10153a072da3164214d9f0d.
//
// Solidity: event RecoveryCallbackCalled(bytes _data)
func (_MockRecoveryModule *MockRecoveryModuleFilterer) WatchRecoveryCallbackCalled(opts *bind.WatchOpts, sink chan<- *MockRecoveryModuleRecoveryCallbackCalled) (event.Subscription, error) {

	logs, sub, err := _MockRecoveryModule.contract.WatchLogs(opts, "RecoveryCallbackCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockRecoveryModuleRecoveryCallbackCalled)
				if err := _MockRecoveryModule.contract.UnpackLog(event, "RecoveryCallbackCalled", log); err != nil {
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

// ParseRecoveryCallbackCalled is a log parse operation binding the contract event 0xa553e8c54a105e80df2e712fd2f1436d50d6f654e10153a072da3164214d9f0d.
//
// Solidity: event RecoveryCallbackCalled(bytes _data)
func (_MockRecoveryModule *MockRecoveryModuleFilterer) ParseRecoveryCallbackCalled(log types.Log) (*MockRecoveryModuleRecoveryCallbackCalled, error) {
	event := new(MockRecoveryModuleRecoveryCallbackCalled)
	if err := _MockRecoveryModule.contract.UnpackLog(event, "RecoveryCallbackCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
