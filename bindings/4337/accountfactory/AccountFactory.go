// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accountfactory

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

// AccountFactoryMetaData contains all meta data concerning the AccountFactory contract.
var AccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"config\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidBundler\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"}],\"name\":\"InvalidSingleton\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_initializer\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONFIG\",\"outputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"computeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0346100e657601f610b1838819003918201601f19168301916001600160401b038311848410176100eb5780849260409485528339810103126100e657602061004882610101565b916001600160a01b0391829161005e9101610101565b169182156100cd57600080546001600160a01b03198116851782556040519491908416907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a316608052610a029081610116823960805181818161015c0152818161025901526103b50152f35b604051631e4fbdf760e01b815260006004820152602490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036100e65756fe608060408181526004918236101561001657600080fd5b600092833560e01c918263296601cd1461033a5750816336b5aa2d14610210578163715018a6146101b35781638da5cb5b1461018b578163d92e82e414610147578163f2fde38b146100b2575063fbcbc0f11461007257600080fd5b346100ae5760203660031901126100ae5760209160ff9082906001600160a01b0361009b61060a565b1681526001855220541690519015158152f35b5080fd5b905034610143576020366003190112610143576100cd61060a565b906100d66106c1565b6001600160a01b0391821692831561012d575050600054826bffffffffffffffffffffffff60a01b821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a380f35b51631e4fbdf760e01b8152908101849052602490fd5b8280fd5b5050346100ae57816003193601126100ae57517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5050346100ae57816003193601126100ae57905490516001600160a01b039091168152602090f35b833461020d578060031936011261020d576101cc6106c1565b600080546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b9050823461020d578260031936011261020d575061022c61060a565b5081516020926102df61024185820184610625565b8083526106ee8584013980519260018060a01b0392837f0000000000000000000000000000000000000000000000000000000000000000168686015285855282850167ffffffffffffffff9580821087831117610325576102c8908286526102b66102b0606083018096610675565b82610675565b03605f1981018352603f190182610625565b519020918051918683019360ff60f81b85523060601b602185015260243560358501526055840152605583526080830195838710908711176103105750849052519020168152f35b604190634e487b7160e01b6000525260246000fd5b604184634e487b7160e01b6000525260246000fd5b848483346101435760603660031901126101435761035661060a565b60249081359267ffffffffffffffff9586851161020d573660238601121561020d5784820135938785116100ae5780860195818636920101116100ae57634ef3c6d960e11b8952328984015260209860443593906001600160a01b03907f00000000000000000000000000000000000000000000000000000000000000008216908c818681855afa9081156105d35786916105ed575b50156105dd578951635eb43cb560e01b8152968216838801819052968c818681855afa9081156105d35786916105a6575b5015610590578951906102df80830191908d83118484101761057e579280928f928a956106ee8439815203019086f58015610574571698893b15610570579087929184888c8c838b61048d83519a8b968795869463347d5e2560e21b86528c8601528c85015260448401916106a0565b03925af1801561056657918b9897969593918b959361050b575b50505050956104fd9187877fa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f743798995260018c5220600160ff198254161790558751948594855260608b86015260608501916106a0565b90868301520390a251908152f35b91939597999698509180945011610555575050865290938693909290918690817fa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f74376104fd8d6104a7565b634e487b7160e01b84526041905282fd5b8a513d87823e3d90fd5b8380fd5b89513d86823e3d90fd5b634e487b7160e01b8852604186528688fd5b50508751630f00f1a160e11b8152908101859052fd5b6105c691508d803d106105cc575b6105be8183610625565b81019061065d565b8d61041d565b503d6105b4565b8b513d88823e3d90fd5b895163f8a0c54b60e01b81528390fd5b61060491508d803d106105cc576105be8183610625565b8d6103ec565b600435906001600160a01b038216820361062057565b600080fd5b90601f8019910116810190811067ffffffffffffffff82111761064757604052565b634e487b7160e01b600052604160045260246000fd5b90816020910312610620575180151581036106205790565b9081519160005b83811061068d575050016000815290565b806020809284010151818501520161067c565b908060209392818452848401376000828201840152601f01601f1916010190565b6000546001600160a01b031633036106d557565b60405163118cdaa760e01b8152336004820152602490fdfe60a034606357601f6102df38819003918201601f19168301916001600160401b03831184841017606857808492602094604052833981010312606357516001600160a01b0381168103606357608052604051610260908161007f8239608051815050f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe60806040526004361015610024575b361561001f5734156101eb57600080fd5b6101eb565b6000803560e01c63d1f578941461003b575061000e565b346100af5760403660031901126100af576004356001600160a01b03811681036100ab576024359067ffffffffffffffff908183116100a757366023840112156100a75782600401359182116100a75736602483850101116100a75760246100a4930190610111565b80f35b8380fd5b5080fd5b80fd5b634e487b7160e01b600052604160045260246000fd5b6020808252825181830181905290939260005b8281106100fd57505060409293506000838284010152601f8019910116010190565b8181018601518482016040015285016100db565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8054929390926001600160a01b03166101d95760009382859455816040519283928337810184815203915af43d156101d15767ffffffffffffffff903d8281116101cc5760405192601f8201601f19908116603f01168401908111848210176101cc5760405282523d6000602084013e5b156101ab5750565b604051633018224d60e21b81529081906101c890600483016100c8565b0390fd5b6100b2565b6060906101a3565b60405163c28d69c760e01b8152600490fd5b600036818037808036817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc545af43d82803e15610226573d90f35b3d90fdfea264697066735822122058f0f8eb4a924f70bd7688d8b719e0c015fc0aae4f06b8a7f34ce91ee0d9998864736f6c63430008190033a26469706673582212206c92d915e50ccf7c60a05e36965feb19793020f966a80230310216e26c763dc264736f6c63430008190033",
}

// AccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountFactoryMetaData.ABI instead.
var AccountFactoryABI = AccountFactoryMetaData.ABI

// AccountFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccountFactoryMetaData.Bin instead.
var AccountFactoryBin = AccountFactoryMetaData.Bin

// DeployAccountFactory deploys a new Ethereum contract, binding an instance of AccountFactory to it.
func DeployAccountFactory(auth *bind.TransactOpts, backend bind.ContractBackend, config common.Address, initialOwner common.Address) (common.Address, *types.Transaction, *AccountFactory, error) {
	parsed, err := AccountFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccountFactoryBin), backend, config, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountFactory{AccountFactoryCaller: AccountFactoryCaller{contract: contract}, AccountFactoryTransactor: AccountFactoryTransactor{contract: contract}, AccountFactoryFilterer: AccountFactoryFilterer{contract: contract}}, nil
}

// AccountFactory is an auto generated Go binding around an Ethereum contract.
type AccountFactory struct {
	AccountFactoryCaller     // Read-only binding to the contract
	AccountFactoryTransactor // Write-only binding to the contract
	AccountFactoryFilterer   // Log filterer for contract events
}

// AccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountFactorySession struct {
	Contract     *AccountFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountFactoryCallerSession struct {
	Contract *AccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountFactoryTransactorSession struct {
	Contract     *AccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountFactoryRaw struct {
	Contract *AccountFactory // Generic contract binding to access the raw methods on
}

// AccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountFactoryCallerRaw struct {
	Contract *AccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountFactoryTransactorRaw struct {
	Contract *AccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountFactory creates a new instance of AccountFactory, bound to a specific deployed contract.
func NewAccountFactory(address common.Address, backend bind.ContractBackend) (*AccountFactory, error) {
	contract, err := bindAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountFactory{AccountFactoryCaller: AccountFactoryCaller{contract: contract}, AccountFactoryTransactor: AccountFactoryTransactor{contract: contract}, AccountFactoryFilterer: AccountFactoryFilterer{contract: contract}}, nil
}

// NewAccountFactoryCaller creates a new read-only instance of AccountFactory, bound to a specific deployed contract.
func NewAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*AccountFactoryCaller, error) {
	contract, err := bindAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryCaller{contract: contract}, nil
}

// NewAccountFactoryTransactor creates a new write-only instance of AccountFactory, bound to a specific deployed contract.
func NewAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountFactoryTransactor, error) {
	contract, err := bindAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryTransactor{contract: contract}, nil
}

// NewAccountFactoryFilterer creates a new log filterer instance of AccountFactory, bound to a specific deployed contract.
func NewAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountFactoryFilterer, error) {
	contract, err := bindAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryFilterer{contract: contract}, nil
}

// bindAccountFactory binds a generic wrapper to an already deployed contract.
func bindAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountFactory *AccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountFactory.Contract.AccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountFactory *AccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountFactory.Contract.AccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountFactory *AccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountFactory.Contract.AccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountFactory *AccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountFactory *AccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountFactory *AccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountFactory.Contract.contract.Transact(opts, method, params...)
}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_AccountFactory *AccountFactoryCaller) CONFIG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "CONFIG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_AccountFactory *AccountFactorySession) CONFIG() (common.Address, error) {
	return _AccountFactory.Contract.CONFIG(&_AccountFactory.CallOpts)
}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_AccountFactory *AccountFactoryCallerSession) CONFIG() (common.Address, error) {
	return _AccountFactory.Contract.CONFIG(&_AccountFactory.CallOpts)
}

// ComputeAddress is a free data retrieval call binding the contract method 0x36b5aa2d.
//
// Solidity: function computeAddress(address , uint256 _salt) view returns(address)
func (_AccountFactory *AccountFactoryCaller) ComputeAddress(opts *bind.CallOpts, arg0 common.Address, _salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "computeAddress", arg0, _salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeAddress is a free data retrieval call binding the contract method 0x36b5aa2d.
//
// Solidity: function computeAddress(address , uint256 _salt) view returns(address)
func (_AccountFactory *AccountFactorySession) ComputeAddress(arg0 common.Address, _salt *big.Int) (common.Address, error) {
	return _AccountFactory.Contract.ComputeAddress(&_AccountFactory.CallOpts, arg0, _salt)
}

// ComputeAddress is a free data retrieval call binding the contract method 0x36b5aa2d.
//
// Solidity: function computeAddress(address , uint256 _salt) view returns(address)
func (_AccountFactory *AccountFactoryCallerSession) ComputeAddress(arg0 common.Address, _salt *big.Int) (common.Address, error) {
	return _AccountFactory.Contract.ComputeAddress(&_AccountFactory.CallOpts, arg0, _salt)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address ) view returns(bool)
func (_AccountFactory *AccountFactoryCaller) GetAccount(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "getAccount", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address ) view returns(bool)
func (_AccountFactory *AccountFactorySession) GetAccount(arg0 common.Address) (bool, error) {
	return _AccountFactory.Contract.GetAccount(&_AccountFactory.CallOpts, arg0)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address ) view returns(bool)
func (_AccountFactory *AccountFactoryCallerSession) GetAccount(arg0 common.Address) (bool, error) {
	return _AccountFactory.Contract.GetAccount(&_AccountFactory.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountFactory *AccountFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountFactory *AccountFactorySession) Owner() (common.Address, error) {
	return _AccountFactory.Contract.Owner(&_AccountFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountFactory *AccountFactoryCallerSession) Owner() (common.Address, error) {
	return _AccountFactory.Contract.Owner(&_AccountFactory.CallOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x296601cd.
//
// Solidity: function createAccount(address _implementation, bytes _initializer, uint256 _salt) returns(address)
func (_AccountFactory *AccountFactoryTransactor) CreateAccount(opts *bind.TransactOpts, _implementation common.Address, _initializer []byte, _salt *big.Int) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "createAccount", _implementation, _initializer, _salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x296601cd.
//
// Solidity: function createAccount(address _implementation, bytes _initializer, uint256 _salt) returns(address)
func (_AccountFactory *AccountFactorySession) CreateAccount(_implementation common.Address, _initializer []byte, _salt *big.Int) (*types.Transaction, error) {
	return _AccountFactory.Contract.CreateAccount(&_AccountFactory.TransactOpts, _implementation, _initializer, _salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x296601cd.
//
// Solidity: function createAccount(address _implementation, bytes _initializer, uint256 _salt) returns(address)
func (_AccountFactory *AccountFactoryTransactorSession) CreateAccount(_implementation common.Address, _initializer []byte, _salt *big.Int) (*types.Transaction, error) {
	return _AccountFactory.Contract.CreateAccount(&_AccountFactory.TransactOpts, _implementation, _initializer, _salt)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountFactory *AccountFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountFactory *AccountFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountFactory.Contract.RenounceOwnership(&_AccountFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountFactory *AccountFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountFactory.Contract.RenounceOwnership(&_AccountFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountFactory *AccountFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountFactory *AccountFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountFactory.Contract.TransferOwnership(&_AccountFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountFactory *AccountFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountFactory.Contract.TransferOwnership(&_AccountFactory.TransactOpts, newOwner)
}

// AccountFactoryAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the AccountFactory contract.
type AccountFactoryAccountCreatedIterator struct {
	Event *AccountFactoryAccountCreated // Event containing the contract specifics and raw log

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
func (it *AccountFactoryAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountFactoryAccountCreated)
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
		it.Event = new(AccountFactoryAccountCreated)
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
func (it *AccountFactoryAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountFactoryAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountFactoryAccountCreated represents a AccountCreated event raised by the AccountFactory contract.
type AccountFactoryAccountCreated struct {
	Account        common.Address
	Implementation common.Address
	Initializer    []byte
	Salt           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0xa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f7437.
//
// Solidity: event AccountCreated(address indexed account, address _implementation, bytes _initializer, uint256 _salt)
func (_AccountFactory *AccountFactoryFilterer) FilterAccountCreated(opts *bind.FilterOpts, account []common.Address) (*AccountFactoryAccountCreatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccountFactory.contract.FilterLogs(opts, "AccountCreated", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryAccountCreatedIterator{contract: _AccountFactory.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0xa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f7437.
//
// Solidity: event AccountCreated(address indexed account, address _implementation, bytes _initializer, uint256 _salt)
func (_AccountFactory *AccountFactoryFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *AccountFactoryAccountCreated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccountFactory.contract.WatchLogs(opts, "AccountCreated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountFactoryAccountCreated)
				if err := _AccountFactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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

// ParseAccountCreated is a log parse operation binding the contract event 0xa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f7437.
//
// Solidity: event AccountCreated(address indexed account, address _implementation, bytes _initializer, uint256 _salt)
func (_AccountFactory *AccountFactoryFilterer) ParseAccountCreated(log types.Log) (*AccountFactoryAccountCreated, error) {
	event := new(AccountFactoryAccountCreated)
	if err := _AccountFactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AccountFactory contract.
type AccountFactoryOwnershipTransferredIterator struct {
	Event *AccountFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountFactoryOwnershipTransferred)
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
		it.Event = new(AccountFactoryOwnershipTransferred)
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
func (it *AccountFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the AccountFactory contract.
type AccountFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountFactory *AccountFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryOwnershipTransferredIterator{contract: _AccountFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountFactory *AccountFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountFactoryOwnershipTransferred)
				if err := _AccountFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AccountFactory *AccountFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*AccountFactoryOwnershipTransferred, error) {
	event := new(AccountFactoryOwnershipTransferred)
	if err := _AccountFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
