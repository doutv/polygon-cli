// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package config

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

// ConfigMetaData contains all meta data concerning the Config contract.
var ConfigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultCallbackHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBundler\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"singer\",\"type\":\"address\"}],\"name\":\"RecoverySignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetSafeSingleton\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetSenderSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"WhitelistBundlerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"WhitelistBundlerRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_CALLBACK_HANDLER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"}],\"name\":\"addSafeSingleton\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"addWhitelistedBundlers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverySigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"}],\"name\":\"removeSafeSingleton\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"removeWhitelistedBundlers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"safeSingleton\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"setRecoverySigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eoaSigner\",\"type\":\"address\"}],\"name\":\"setWalletSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"walletSigners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedBundler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080346100d557601f6108d338819003918201601f19168301916001600160401b038311848410176100da5780849260409485528339810103126100d5576020610048826100f0565b916001600160a01b0391829161005e91016100f0565b169081156100bc576000549260018060a01b03199280848616176000558260405195167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3169060015416176001556107ce90816101058239f35b604051631e4fbdf760e01b815260006004820152602490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036100d55756fe6040608081526004908136101561001557600080fd5b600091823560e01c90816303087a57146105b357816316b97e9a1461052f5781635eb43cb5146104f25781637063908f1461046f578163715018a6146104125781637910c00b146103e95781637a1c8449146103ae57816387ab75e8146102e85781638da5cb5b146102c05781639de78db214610282578163be08382114610259578163ee3a5fab146101d7578163f2fde38b1461014c575063fcdc4727146100bd57600080fd5b34610148576100cb36610647565b906100d461076c565b8151835b81811061011657847f49bd286efd7e26d327db1b6f433560c5f6156363198289a65a6ca9904338f4f461011086865191829182610727565b0390a180f35b6001600160a01b0361012882866106fd565b511685526002602052828520805460ff19166001908117909155016100d8565b5080fd5b9050346101d35760203660031901126101d35761016761062c565b9061017061076c565b6001600160a01b039182169283156101bd57505082546001600160a01b0319811683178455167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b51631e4fbdf760e01b8152908101849052602490fd5b8280fd5b919050346101d35760203660031901126101d3576101f361062c565b6101fb61076c565b6001600160a01b031691821561024b5750600580546001600160a01b03191683179055519081527f75da6d4a6a549915e7b28de22af1418486fc73dc2359d7a6851a3ceca2982a6b90602090a180f35b905163e6c4247b60e01b8152fd5b50503461014857816003193601126101485760055490516001600160a01b039091168152602090f35b5050346101485760203660031901126101485760209160ff9082906001600160a01b036102ad61062c565b1681526002855220541690519015158152f35b505034610148578160031936011261014857905490516001600160a01b039091168152602090f35b919050346101d357806003193601126101d35761030361062c565b6024356001600160a01b03818116949092918590036103aa57841561039c57328652600260205260ff84872054161561038e5750168084526003602090815282852080546001600160a01b03191685179055825191825281019290925242908201527f03b72cb2fffc620d2a9c4f6bce0636faa82c948488a45f2f4b0684bb17cfdc7290606090a180f35b835163f8a0c54b60e01b8152fd5b835163e6c4247b60e01b8152fd5b8580fd5b505034610148576020366003190112610148576020916001600160a01b03908290826103d861062c565b168152600385522054169051908152f35b50503461014857816003193601126101485760015490516001600160a01b039091168152602090f35b833461046c578060031936011261046c5761042b61076c565b80546001600160a01b03198116825581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b919050346101d35760203660031901126101d35761048b61062c565b61049361076c565b6001600160a01b031691821561024b57916060917ff80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f8040619193828652602052808520600160ff198254161790558051918252600160208301524290820152a180f35b9050346101d35760203660031901126101d35760209260ff918391906001600160a01b0361051e61062c565b168252855220541690519015158152f35b5050346101485761053f36610647565b9061054861076c565b8151835b81811061058457847f623f7acab9936b83f17999b4d96bfda34b1d3a749bd474d39d0413cc977044ed61011086865191829182610727565b6001906001600160a01b0361059982876106fd565b511686526002602052838620805460ff191690550161054c565b919050346101d35760203660031901126101d3577ff80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f80406191916060916105f461062c565b6105fc61076c565b6001600160a01b03168086526020928352818620805460ff1916905581519081529182018590524290820152a180f35b600435906001600160a01b038216820361064257565b600080fd5b6020806003198301126106425767ffffffffffffffff916004358381116106425781602382011215610642578060040135908482116106e7578160051b9160405195601f19603f850116870190878210908211176106e75760405285526024602086019282010192831161064257602401905b8282106106c8575050505090565b81356001600160a01b03811681036106425781529083019083016106ba565b634e487b7160e01b600052604160045260246000fd5b80518210156107115760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b602090602060408183019282815285518094520193019160005b82811061074f575050505090565b83516001600160a01b031685529381019392810192600101610741565b6000546001600160a01b0316330361078057565b60405163118cdaa760e01b8152336004820152602490fdfea2646970667358221220f2d294bc808728b27b345d7c1cb476d3b42a0c324192f687c879ce2b981939e164736f6c63430008190033",
}

// ConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use ConfigMetaData.ABI instead.
var ConfigABI = ConfigMetaData.ABI

// ConfigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConfigMetaData.Bin instead.
var ConfigBin = ConfigMetaData.Bin

// DeployConfig deploys a new Ethereum contract, binding an instance of Config to it.
func DeployConfig(auth *bind.TransactOpts, backend bind.ContractBackend, defaultCallbackHandler common.Address, initialOwner common.Address) (common.Address, *types.Transaction, *Config, error) {
	parsed, err := ConfigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConfigBin), backend, defaultCallbackHandler, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Config{ConfigCaller: ConfigCaller{contract: contract}, ConfigTransactor: ConfigTransactor{contract: contract}, ConfigFilterer: ConfigFilterer{contract: contract}}, nil
}

// Config is an auto generated Go binding around an Ethereum contract.
type Config struct {
	ConfigCaller     // Read-only binding to the contract
	ConfigTransactor // Write-only binding to the contract
	ConfigFilterer   // Log filterer for contract events
}

// ConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConfigSession struct {
	Contract     *Config           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConfigCallerSession struct {
	Contract *ConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConfigTransactorSession struct {
	Contract     *ConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConfigRaw struct {
	Contract *Config // Generic contract binding to access the raw methods on
}

// ConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConfigCallerRaw struct {
	Contract *ConfigCaller // Generic read-only contract binding to access the raw methods on
}

// ConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConfigTransactorRaw struct {
	Contract *ConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConfig creates a new instance of Config, bound to a specific deployed contract.
func NewConfig(address common.Address, backend bind.ContractBackend) (*Config, error) {
	contract, err := bindConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Config{ConfigCaller: ConfigCaller{contract: contract}, ConfigTransactor: ConfigTransactor{contract: contract}, ConfigFilterer: ConfigFilterer{contract: contract}}, nil
}

// NewConfigCaller creates a new read-only instance of Config, bound to a specific deployed contract.
func NewConfigCaller(address common.Address, caller bind.ContractCaller) (*ConfigCaller, error) {
	contract, err := bindConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigCaller{contract: contract}, nil
}

// NewConfigTransactor creates a new write-only instance of Config, bound to a specific deployed contract.
func NewConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfigTransactor, error) {
	contract, err := bindConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigTransactor{contract: contract}, nil
}

// NewConfigFilterer creates a new log filterer instance of Config, bound to a specific deployed contract.
func NewConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfigFilterer, error) {
	contract, err := bindConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfigFilterer{contract: contract}, nil
}

// bindConfig binds a generic wrapper to an already deployed contract.
func bindConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Config *ConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Config.Contract.ConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Config *ConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.Contract.ConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Config *ConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Config.Contract.ConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Config *ConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Config.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Config *ConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Config *ConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Config.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTCALLBACKHANDLER is a free data retrieval call binding the contract method 0x7910c00b.
//
// Solidity: function DEFAULT_CALLBACK_HANDLER() view returns(address)
func (_Config *ConfigCaller) DEFAULTCALLBACKHANDLER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "DEFAULT_CALLBACK_HANDLER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEFAULTCALLBACKHANDLER is a free data retrieval call binding the contract method 0x7910c00b.
//
// Solidity: function DEFAULT_CALLBACK_HANDLER() view returns(address)
func (_Config *ConfigSession) DEFAULTCALLBACKHANDLER() (common.Address, error) {
	return _Config.Contract.DEFAULTCALLBACKHANDLER(&_Config.CallOpts)
}

// DEFAULTCALLBACKHANDLER is a free data retrieval call binding the contract method 0x7910c00b.
//
// Solidity: function DEFAULT_CALLBACK_HANDLER() view returns(address)
func (_Config *ConfigCallerSession) DEFAULTCALLBACKHANDLER() (common.Address, error) {
	return _Config.Contract.DEFAULTCALLBACKHANDLER(&_Config.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Config *ConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Config *ConfigSession) Owner() (common.Address, error) {
	return _Config.Contract.Owner(&_Config.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Config *ConfigCallerSession) Owner() (common.Address, error) {
	return _Config.Contract.Owner(&_Config.CallOpts)
}

// RecoverySigner is a free data retrieval call binding the contract method 0xbe083821.
//
// Solidity: function recoverySigner() view returns(address)
func (_Config *ConfigCaller) RecoverySigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "recoverySigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverySigner is a free data retrieval call binding the contract method 0xbe083821.
//
// Solidity: function recoverySigner() view returns(address)
func (_Config *ConfigSession) RecoverySigner() (common.Address, error) {
	return _Config.Contract.RecoverySigner(&_Config.CallOpts)
}

// RecoverySigner is a free data retrieval call binding the contract method 0xbe083821.
//
// Solidity: function recoverySigner() view returns(address)
func (_Config *ConfigCallerSession) RecoverySigner() (common.Address, error) {
	return _Config.Contract.RecoverySigner(&_Config.CallOpts)
}

// SafeSingleton is a free data retrieval call binding the contract method 0x5eb43cb5.
//
// Solidity: function safeSingleton(address ) view returns(bool)
func (_Config *ConfigCaller) SafeSingleton(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "safeSingleton", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SafeSingleton is a free data retrieval call binding the contract method 0x5eb43cb5.
//
// Solidity: function safeSingleton(address ) view returns(bool)
func (_Config *ConfigSession) SafeSingleton(arg0 common.Address) (bool, error) {
	return _Config.Contract.SafeSingleton(&_Config.CallOpts, arg0)
}

// SafeSingleton is a free data retrieval call binding the contract method 0x5eb43cb5.
//
// Solidity: function safeSingleton(address ) view returns(bool)
func (_Config *ConfigCallerSession) SafeSingleton(arg0 common.Address) (bool, error) {
	return _Config.Contract.SafeSingleton(&_Config.CallOpts, arg0)
}

// WalletSigners is a free data retrieval call binding the contract method 0x7a1c8449.
//
// Solidity: function walletSigners(address ) view returns(address)
func (_Config *ConfigCaller) WalletSigners(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "walletSigners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletSigners is a free data retrieval call binding the contract method 0x7a1c8449.
//
// Solidity: function walletSigners(address ) view returns(address)
func (_Config *ConfigSession) WalletSigners(arg0 common.Address) (common.Address, error) {
	return _Config.Contract.WalletSigners(&_Config.CallOpts, arg0)
}

// WalletSigners is a free data retrieval call binding the contract method 0x7a1c8449.
//
// Solidity: function walletSigners(address ) view returns(address)
func (_Config *ConfigCallerSession) WalletSigners(arg0 common.Address) (common.Address, error) {
	return _Config.Contract.WalletSigners(&_Config.CallOpts, arg0)
}

// WhitelistedBundler is a free data retrieval call binding the contract method 0x9de78db2.
//
// Solidity: function whitelistedBundler(address ) view returns(bool)
func (_Config *ConfigCaller) WhitelistedBundler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "whitelistedBundler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedBundler is a free data retrieval call binding the contract method 0x9de78db2.
//
// Solidity: function whitelistedBundler(address ) view returns(bool)
func (_Config *ConfigSession) WhitelistedBundler(arg0 common.Address) (bool, error) {
	return _Config.Contract.WhitelistedBundler(&_Config.CallOpts, arg0)
}

// WhitelistedBundler is a free data retrieval call binding the contract method 0x9de78db2.
//
// Solidity: function whitelistedBundler(address ) view returns(bool)
func (_Config *ConfigCallerSession) WhitelistedBundler(arg0 common.Address) (bool, error) {
	return _Config.Contract.WhitelistedBundler(&_Config.CallOpts, arg0)
}

// AddSafeSingleton is a paid mutator transaction binding the contract method 0x7063908f.
//
// Solidity: function addSafeSingleton(address singleton) returns()
func (_Config *ConfigTransactor) AddSafeSingleton(opts *bind.TransactOpts, singleton common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "addSafeSingleton", singleton)
}

// AddSafeSingleton is a paid mutator transaction binding the contract method 0x7063908f.
//
// Solidity: function addSafeSingleton(address singleton) returns()
func (_Config *ConfigSession) AddSafeSingleton(singleton common.Address) (*types.Transaction, error) {
	return _Config.Contract.AddSafeSingleton(&_Config.TransactOpts, singleton)
}

// AddSafeSingleton is a paid mutator transaction binding the contract method 0x7063908f.
//
// Solidity: function addSafeSingleton(address singleton) returns()
func (_Config *ConfigTransactorSession) AddSafeSingleton(singleton common.Address) (*types.Transaction, error) {
	return _Config.Contract.AddSafeSingleton(&_Config.TransactOpts, singleton)
}

// AddWhitelistedBundlers is a paid mutator transaction binding the contract method 0xfcdc4727.
//
// Solidity: function addWhitelistedBundlers(address[] bundlers) returns()
func (_Config *ConfigTransactor) AddWhitelistedBundlers(opts *bind.TransactOpts, bundlers []common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "addWhitelistedBundlers", bundlers)
}

// AddWhitelistedBundlers is a paid mutator transaction binding the contract method 0xfcdc4727.
//
// Solidity: function addWhitelistedBundlers(address[] bundlers) returns()
func (_Config *ConfigSession) AddWhitelistedBundlers(bundlers []common.Address) (*types.Transaction, error) {
	return _Config.Contract.AddWhitelistedBundlers(&_Config.TransactOpts, bundlers)
}

// AddWhitelistedBundlers is a paid mutator transaction binding the contract method 0xfcdc4727.
//
// Solidity: function addWhitelistedBundlers(address[] bundlers) returns()
func (_Config *ConfigTransactorSession) AddWhitelistedBundlers(bundlers []common.Address) (*types.Transaction, error) {
	return _Config.Contract.AddWhitelistedBundlers(&_Config.TransactOpts, bundlers)
}

// RemoveSafeSingleton is a paid mutator transaction binding the contract method 0x03087a57.
//
// Solidity: function removeSafeSingleton(address singleton) returns()
func (_Config *ConfigTransactor) RemoveSafeSingleton(opts *bind.TransactOpts, singleton common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "removeSafeSingleton", singleton)
}

// RemoveSafeSingleton is a paid mutator transaction binding the contract method 0x03087a57.
//
// Solidity: function removeSafeSingleton(address singleton) returns()
func (_Config *ConfigSession) RemoveSafeSingleton(singleton common.Address) (*types.Transaction, error) {
	return _Config.Contract.RemoveSafeSingleton(&_Config.TransactOpts, singleton)
}

// RemoveSafeSingleton is a paid mutator transaction binding the contract method 0x03087a57.
//
// Solidity: function removeSafeSingleton(address singleton) returns()
func (_Config *ConfigTransactorSession) RemoveSafeSingleton(singleton common.Address) (*types.Transaction, error) {
	return _Config.Contract.RemoveSafeSingleton(&_Config.TransactOpts, singleton)
}

// RemoveWhitelistedBundlers is a paid mutator transaction binding the contract method 0x16b97e9a.
//
// Solidity: function removeWhitelistedBundlers(address[] bundlers) returns()
func (_Config *ConfigTransactor) RemoveWhitelistedBundlers(opts *bind.TransactOpts, bundlers []common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "removeWhitelistedBundlers", bundlers)
}

// RemoveWhitelistedBundlers is a paid mutator transaction binding the contract method 0x16b97e9a.
//
// Solidity: function removeWhitelistedBundlers(address[] bundlers) returns()
func (_Config *ConfigSession) RemoveWhitelistedBundlers(bundlers []common.Address) (*types.Transaction, error) {
	return _Config.Contract.RemoveWhitelistedBundlers(&_Config.TransactOpts, bundlers)
}

// RemoveWhitelistedBundlers is a paid mutator transaction binding the contract method 0x16b97e9a.
//
// Solidity: function removeWhitelistedBundlers(address[] bundlers) returns()
func (_Config *ConfigTransactorSession) RemoveWhitelistedBundlers(bundlers []common.Address) (*types.Transaction, error) {
	return _Config.Contract.RemoveWhitelistedBundlers(&_Config.TransactOpts, bundlers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Config *ConfigTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Config *ConfigSession) RenounceOwnership() (*types.Transaction, error) {
	return _Config.Contract.RenounceOwnership(&_Config.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Config *ConfigTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Config.Contract.RenounceOwnership(&_Config.TransactOpts)
}

// SetRecoverySigner is a paid mutator transaction binding the contract method 0xee3a5fab.
//
// Solidity: function setRecoverySigner(address signer) returns()
func (_Config *ConfigTransactor) SetRecoverySigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setRecoverySigner", signer)
}

// SetRecoverySigner is a paid mutator transaction binding the contract method 0xee3a5fab.
//
// Solidity: function setRecoverySigner(address signer) returns()
func (_Config *ConfigSession) SetRecoverySigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetRecoverySigner(&_Config.TransactOpts, signer)
}

// SetRecoverySigner is a paid mutator transaction binding the contract method 0xee3a5fab.
//
// Solidity: function setRecoverySigner(address signer) returns()
func (_Config *ConfigTransactorSession) SetRecoverySigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetRecoverySigner(&_Config.TransactOpts, signer)
}

// SetWalletSigner is a paid mutator transaction binding the contract method 0x87ab75e8.
//
// Solidity: function setWalletSigner(address sender, address eoaSigner) returns()
func (_Config *ConfigTransactor) SetWalletSigner(opts *bind.TransactOpts, sender common.Address, eoaSigner common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setWalletSigner", sender, eoaSigner)
}

// SetWalletSigner is a paid mutator transaction binding the contract method 0x87ab75e8.
//
// Solidity: function setWalletSigner(address sender, address eoaSigner) returns()
func (_Config *ConfigSession) SetWalletSigner(sender common.Address, eoaSigner common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetWalletSigner(&_Config.TransactOpts, sender, eoaSigner)
}

// SetWalletSigner is a paid mutator transaction binding the contract method 0x87ab75e8.
//
// Solidity: function setWalletSigner(address sender, address eoaSigner) returns()
func (_Config *ConfigTransactorSession) SetWalletSigner(sender common.Address, eoaSigner common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetWalletSigner(&_Config.TransactOpts, sender, eoaSigner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Config *ConfigTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Config *ConfigSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Config.Contract.TransferOwnership(&_Config.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Config *ConfigTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Config.Contract.TransferOwnership(&_Config.TransactOpts, newOwner)
}

// ConfigOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Config contract.
type ConfigOwnershipTransferredIterator struct {
	Event *ConfigOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ConfigOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigOwnershipTransferred)
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
		it.Event = new(ConfigOwnershipTransferred)
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
func (it *ConfigOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigOwnershipTransferred represents a OwnershipTransferred event raised by the Config contract.
type ConfigOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Config *ConfigFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ConfigOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Config.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConfigOwnershipTransferredIterator{contract: _Config.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Config *ConfigFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConfigOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Config.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigOwnershipTransferred)
				if err := _Config.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Config *ConfigFilterer) ParseOwnershipTransferred(log types.Log) (*ConfigOwnershipTransferred, error) {
	event := new(ConfigOwnershipTransferred)
	if err := _Config.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigRecoverySignerUpdatedIterator is returned from FilterRecoverySignerUpdated and is used to iterate over the raw logs and unpacked data for RecoverySignerUpdated events raised by the Config contract.
type ConfigRecoverySignerUpdatedIterator struct {
	Event *ConfigRecoverySignerUpdated // Event containing the contract specifics and raw log

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
func (it *ConfigRecoverySignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigRecoverySignerUpdated)
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
		it.Event = new(ConfigRecoverySignerUpdated)
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
func (it *ConfigRecoverySignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigRecoverySignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigRecoverySignerUpdated represents a RecoverySignerUpdated event raised by the Config contract.
type ConfigRecoverySignerUpdated struct {
	Singer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverySignerUpdated is a free log retrieval operation binding the contract event 0x75da6d4a6a549915e7b28de22af1418486fc73dc2359d7a6851a3ceca2982a6b.
//
// Solidity: event RecoverySignerUpdated(address singer)
func (_Config *ConfigFilterer) FilterRecoverySignerUpdated(opts *bind.FilterOpts) (*ConfigRecoverySignerUpdatedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "RecoverySignerUpdated")
	if err != nil {
		return nil, err
	}
	return &ConfigRecoverySignerUpdatedIterator{contract: _Config.contract, event: "RecoverySignerUpdated", logs: logs, sub: sub}, nil
}

// WatchRecoverySignerUpdated is a free log subscription operation binding the contract event 0x75da6d4a6a549915e7b28de22af1418486fc73dc2359d7a6851a3ceca2982a6b.
//
// Solidity: event RecoverySignerUpdated(address singer)
func (_Config *ConfigFilterer) WatchRecoverySignerUpdated(opts *bind.WatchOpts, sink chan<- *ConfigRecoverySignerUpdated) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "RecoverySignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigRecoverySignerUpdated)
				if err := _Config.contract.UnpackLog(event, "RecoverySignerUpdated", log); err != nil {
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

// ParseRecoverySignerUpdated is a log parse operation binding the contract event 0x75da6d4a6a549915e7b28de22af1418486fc73dc2359d7a6851a3ceca2982a6b.
//
// Solidity: event RecoverySignerUpdated(address singer)
func (_Config *ConfigFilterer) ParseRecoverySignerUpdated(log types.Log) (*ConfigRecoverySignerUpdated, error) {
	event := new(ConfigRecoverySignerUpdated)
	if err := _Config.contract.UnpackLog(event, "RecoverySignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigSetSafeSingletonIterator is returned from FilterSetSafeSingleton and is used to iterate over the raw logs and unpacked data for SetSafeSingleton events raised by the Config contract.
type ConfigSetSafeSingletonIterator struct {
	Event *ConfigSetSafeSingleton // Event containing the contract specifics and raw log

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
func (it *ConfigSetSafeSingletonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigSetSafeSingleton)
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
		it.Event = new(ConfigSetSafeSingleton)
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
func (it *ConfigSetSafeSingletonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigSetSafeSingletonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigSetSafeSingleton represents a SetSafeSingleton event raised by the Config contract.
type ConfigSetSafeSingleton struct {
	Singleton common.Address
	Status    bool
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetSafeSingleton is a free log retrieval operation binding the contract event 0xf80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f80406191.
//
// Solidity: event SetSafeSingleton(address singleton, bool status, uint256 timestamp)
func (_Config *ConfigFilterer) FilterSetSafeSingleton(opts *bind.FilterOpts) (*ConfigSetSafeSingletonIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "SetSafeSingleton")
	if err != nil {
		return nil, err
	}
	return &ConfigSetSafeSingletonIterator{contract: _Config.contract, event: "SetSafeSingleton", logs: logs, sub: sub}, nil
}

// WatchSetSafeSingleton is a free log subscription operation binding the contract event 0xf80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f80406191.
//
// Solidity: event SetSafeSingleton(address singleton, bool status, uint256 timestamp)
func (_Config *ConfigFilterer) WatchSetSafeSingleton(opts *bind.WatchOpts, sink chan<- *ConfigSetSafeSingleton) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "SetSafeSingleton")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigSetSafeSingleton)
				if err := _Config.contract.UnpackLog(event, "SetSafeSingleton", log); err != nil {
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

// ParseSetSafeSingleton is a log parse operation binding the contract event 0xf80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f80406191.
//
// Solidity: event SetSafeSingleton(address singleton, bool status, uint256 timestamp)
func (_Config *ConfigFilterer) ParseSetSafeSingleton(log types.Log) (*ConfigSetSafeSingleton, error) {
	event := new(ConfigSetSafeSingleton)
	if err := _Config.contract.UnpackLog(event, "SetSafeSingleton", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigSetSenderSignerIterator is returned from FilterSetSenderSigner and is used to iterate over the raw logs and unpacked data for SetSenderSigner events raised by the Config contract.
type ConfigSetSenderSignerIterator struct {
	Event *ConfigSetSenderSigner // Event containing the contract specifics and raw log

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
func (it *ConfigSetSenderSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigSetSenderSigner)
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
		it.Event = new(ConfigSetSenderSigner)
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
func (it *ConfigSetSenderSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigSetSenderSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigSetSenderSigner represents a SetSenderSigner event raised by the Config contract.
type ConfigSetSenderSigner struct {
	Sender    common.Address
	Signer    common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetSenderSigner is a free log retrieval operation binding the contract event 0x03b72cb2fffc620d2a9c4f6bce0636faa82c948488a45f2f4b0684bb17cfdc72.
//
// Solidity: event SetSenderSigner(address sender, address signer, uint256 timestamp)
func (_Config *ConfigFilterer) FilterSetSenderSigner(opts *bind.FilterOpts) (*ConfigSetSenderSignerIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "SetSenderSigner")
	if err != nil {
		return nil, err
	}
	return &ConfigSetSenderSignerIterator{contract: _Config.contract, event: "SetSenderSigner", logs: logs, sub: sub}, nil
}

// WatchSetSenderSigner is a free log subscription operation binding the contract event 0x03b72cb2fffc620d2a9c4f6bce0636faa82c948488a45f2f4b0684bb17cfdc72.
//
// Solidity: event SetSenderSigner(address sender, address signer, uint256 timestamp)
func (_Config *ConfigFilterer) WatchSetSenderSigner(opts *bind.WatchOpts, sink chan<- *ConfigSetSenderSigner) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "SetSenderSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigSetSenderSigner)
				if err := _Config.contract.UnpackLog(event, "SetSenderSigner", log); err != nil {
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

// ParseSetSenderSigner is a log parse operation binding the contract event 0x03b72cb2fffc620d2a9c4f6bce0636faa82c948488a45f2f4b0684bb17cfdc72.
//
// Solidity: event SetSenderSigner(address sender, address signer, uint256 timestamp)
func (_Config *ConfigFilterer) ParseSetSenderSigner(log types.Log) (*ConfigSetSenderSigner, error) {
	event := new(ConfigSetSenderSigner)
	if err := _Config.contract.UnpackLog(event, "SetSenderSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigWhitelistBundlerAddedIterator is returned from FilterWhitelistBundlerAdded and is used to iterate over the raw logs and unpacked data for WhitelistBundlerAdded events raised by the Config contract.
type ConfigWhitelistBundlerAddedIterator struct {
	Event *ConfigWhitelistBundlerAdded // Event containing the contract specifics and raw log

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
func (it *ConfigWhitelistBundlerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigWhitelistBundlerAdded)
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
		it.Event = new(ConfigWhitelistBundlerAdded)
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
func (it *ConfigWhitelistBundlerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigWhitelistBundlerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigWhitelistBundlerAdded represents a WhitelistBundlerAdded event raised by the Config contract.
type ConfigWhitelistBundlerAdded struct {
	Bundlers []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistBundlerAdded is a free log retrieval operation binding the contract event 0x49bd286efd7e26d327db1b6f433560c5f6156363198289a65a6ca9904338f4f4.
//
// Solidity: event WhitelistBundlerAdded(address[] bundlers)
func (_Config *ConfigFilterer) FilterWhitelistBundlerAdded(opts *bind.FilterOpts) (*ConfigWhitelistBundlerAddedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "WhitelistBundlerAdded")
	if err != nil {
		return nil, err
	}
	return &ConfigWhitelistBundlerAddedIterator{contract: _Config.contract, event: "WhitelistBundlerAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistBundlerAdded is a free log subscription operation binding the contract event 0x49bd286efd7e26d327db1b6f433560c5f6156363198289a65a6ca9904338f4f4.
//
// Solidity: event WhitelistBundlerAdded(address[] bundlers)
func (_Config *ConfigFilterer) WatchWhitelistBundlerAdded(opts *bind.WatchOpts, sink chan<- *ConfigWhitelistBundlerAdded) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "WhitelistBundlerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigWhitelistBundlerAdded)
				if err := _Config.contract.UnpackLog(event, "WhitelistBundlerAdded", log); err != nil {
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

// ParseWhitelistBundlerAdded is a log parse operation binding the contract event 0x49bd286efd7e26d327db1b6f433560c5f6156363198289a65a6ca9904338f4f4.
//
// Solidity: event WhitelistBundlerAdded(address[] bundlers)
func (_Config *ConfigFilterer) ParseWhitelistBundlerAdded(log types.Log) (*ConfigWhitelistBundlerAdded, error) {
	event := new(ConfigWhitelistBundlerAdded)
	if err := _Config.contract.UnpackLog(event, "WhitelistBundlerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigWhitelistBundlerRemovedIterator is returned from FilterWhitelistBundlerRemoved and is used to iterate over the raw logs and unpacked data for WhitelistBundlerRemoved events raised by the Config contract.
type ConfigWhitelistBundlerRemovedIterator struct {
	Event *ConfigWhitelistBundlerRemoved // Event containing the contract specifics and raw log

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
func (it *ConfigWhitelistBundlerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigWhitelistBundlerRemoved)
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
		it.Event = new(ConfigWhitelistBundlerRemoved)
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
func (it *ConfigWhitelistBundlerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigWhitelistBundlerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigWhitelistBundlerRemoved represents a WhitelistBundlerRemoved event raised by the Config contract.
type ConfigWhitelistBundlerRemoved struct {
	Bundlers []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistBundlerRemoved is a free log retrieval operation binding the contract event 0x623f7acab9936b83f17999b4d96bfda34b1d3a749bd474d39d0413cc977044ed.
//
// Solidity: event WhitelistBundlerRemoved(address[] bundlers)
func (_Config *ConfigFilterer) FilterWhitelistBundlerRemoved(opts *bind.FilterOpts) (*ConfigWhitelistBundlerRemovedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "WhitelistBundlerRemoved")
	if err != nil {
		return nil, err
	}
	return &ConfigWhitelistBundlerRemovedIterator{contract: _Config.contract, event: "WhitelistBundlerRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistBundlerRemoved is a free log subscription operation binding the contract event 0x623f7acab9936b83f17999b4d96bfda34b1d3a749bd474d39d0413cc977044ed.
//
// Solidity: event WhitelistBundlerRemoved(address[] bundlers)
func (_Config *ConfigFilterer) WatchWhitelistBundlerRemoved(opts *bind.WatchOpts, sink chan<- *ConfigWhitelistBundlerRemoved) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "WhitelistBundlerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigWhitelistBundlerRemoved)
				if err := _Config.contract.UnpackLog(event, "WhitelistBundlerRemoved", log); err != nil {
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

// ParseWhitelistBundlerRemoved is a log parse operation binding the contract event 0x623f7acab9936b83f17999b4d96bfda34b1d3a749bd474d39d0413cc977044ed.
//
// Solidity: event WhitelistBundlerRemoved(address[] bundlers)
func (_Config *ConfigFilterer) ParseWhitelistBundlerRemoved(log types.Log) (*ConfigWhitelistBundlerRemoved, error) {
	event := new(ConfigWhitelistBundlerRemoved)
	if err := _Config.contract.UnpackLog(event, "WhitelistBundlerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
