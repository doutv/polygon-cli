// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basemode

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

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	Sender             common.Address
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

// BaseModeMetaData contains all meta data concerning the BaseMode contract.
var BaseModeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_config\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"config\",\"type\":\"address\"}],\"name\":\"SetConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasManager\",\"type\":\"address\"}],\"name\":\"SetGasManager\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_THIS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIG\",\"outputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ENTRYPOINT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIG_VALIDATION_FAILED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"postOpMode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"requiredPreFund\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e03461010357601f6105da38819003918201601f19168301916001600160401b0383118484101761010857808492606094604052833981010312610103576100478161011e565b61005f60406100586020850161011e565b930161011e565b6001600160a01b039182169283156100ea57600080546001600160a01b03198116861782556040519591908516907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a33060c05260a052166080526104a79081610133823960805181610192015260a05181818161014a015261027b015260c051816101d60152f35b604051631e4fbdf760e01b815260006004820152602490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101035756fe6080604081815260049081361015610022575b505050361561002057600080fd5b005b600092833560e01c90816352b7512c146103c257508063715018a6146103685780637c627b211461030c57838163813f3f4414610249575080638da5cb5b146102215780638f41ec5a14610205578063cc025f7c146101c1578063d92e82e41461017d578063e8eb3cc6146101355763f2fde38b036100125734610131576020366003190112610131576001600160a01b0382358181169391929084900361012d576100cc610445565b8315610117575050600054826bffffffffffffffffffffffff60a01b821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a380f35b51631e4fbdf760e01b8152908101849052602490fd5b8480fd5b8280fd5b838234610179578160031936011261017957517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b838234610179578160031936011261017957517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b838234610179578160031936011261017957517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b8382346101795781600319360112610179576020905160018152f35b838234610179578160031936011261017957905490516001600160a01b039091168152602090f35b9290503461030857816003193601126103085780356001600160a01b038181169182900361012d57610279610445565b7f000000000000000000000000000000000000000000000000000000000000000016803b1561012d57604485928551968793849263040b850f60e31b84528784015260243560248401525af180156102fe576102d3578380f35b67ffffffffffffffff83116102eb5750523880808380f35b634e487b7160e01b845260419052602483fd5b82513d86823e3d90fd5b5050fd5b50503461017957608036600319011261017957600381351015610179576024359067ffffffffffffffff9081831161036457366023840112156103645782013590811161013157369101602401116103615780f35b80fd5b8380fd5b8334610361578060031936011261036157610381610445565b80546001600160a01b03198116825581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b828585346101795760609060031990606082360112610364573567ffffffffffffffff811161036457906101209136030112610179579083918383526060518094840152815b84811061042f57505091826060938483850101526020830152601f80199101168101030190f35b6080810151868201830152859350602001610408565b6000546001600160a01b0316330361045957565b60405163118cdaa760e01b8152336004820152602490fdfea26469706673582212208ea67df5beb192ef4775d3c59a38879ffdc4a332725150cc881777651fa499f464736f6c63430008190033",
}

// BaseModeABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseModeMetaData.ABI instead.
var BaseModeABI = BaseModeMetaData.ABI

// BaseModeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BaseModeMetaData.Bin instead.
var BaseModeBin = BaseModeMetaData.Bin

// DeployBaseMode deploys a new Ethereum contract, binding an instance of BaseMode to it.
func DeployBaseMode(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _entryPoint common.Address, _config common.Address) (common.Address, *types.Transaction, *BaseMode, error) {
	parsed, err := BaseModeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BaseModeBin), backend, _owner, _entryPoint, _config)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseMode{BaseModeCaller: BaseModeCaller{contract: contract}, BaseModeTransactor: BaseModeTransactor{contract: contract}, BaseModeFilterer: BaseModeFilterer{contract: contract}}, nil
}

// BaseMode is an auto generated Go binding around an Ethereum contract.
type BaseMode struct {
	BaseModeCaller     // Read-only binding to the contract
	BaseModeTransactor // Write-only binding to the contract
	BaseModeFilterer   // Log filterer for contract events
}

// BaseModeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseModeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseModeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseModeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseModeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseModeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseModeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseModeSession struct {
	Contract     *BaseMode         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseModeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseModeCallerSession struct {
	Contract *BaseModeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BaseModeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseModeTransactorSession struct {
	Contract     *BaseModeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BaseModeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseModeRaw struct {
	Contract *BaseMode // Generic contract binding to access the raw methods on
}

// BaseModeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseModeCallerRaw struct {
	Contract *BaseModeCaller // Generic read-only contract binding to access the raw methods on
}

// BaseModeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseModeTransactorRaw struct {
	Contract *BaseModeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseMode creates a new instance of BaseMode, bound to a specific deployed contract.
func NewBaseMode(address common.Address, backend bind.ContractBackend) (*BaseMode, error) {
	contract, err := bindBaseMode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseMode{BaseModeCaller: BaseModeCaller{contract: contract}, BaseModeTransactor: BaseModeTransactor{contract: contract}, BaseModeFilterer: BaseModeFilterer{contract: contract}}, nil
}

// NewBaseModeCaller creates a new read-only instance of BaseMode, bound to a specific deployed contract.
func NewBaseModeCaller(address common.Address, caller bind.ContractCaller) (*BaseModeCaller, error) {
	contract, err := bindBaseMode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseModeCaller{contract: contract}, nil
}

// NewBaseModeTransactor creates a new write-only instance of BaseMode, bound to a specific deployed contract.
func NewBaseModeTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseModeTransactor, error) {
	contract, err := bindBaseMode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseModeTransactor{contract: contract}, nil
}

// NewBaseModeFilterer creates a new log filterer instance of BaseMode, bound to a specific deployed contract.
func NewBaseModeFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseModeFilterer, error) {
	contract, err := bindBaseMode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseModeFilterer{contract: contract}, nil
}

// bindBaseMode binds a generic wrapper to an already deployed contract.
func bindBaseMode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseModeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseMode *BaseModeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseMode.Contract.BaseModeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseMode *BaseModeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMode.Contract.BaseModeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseMode *BaseModeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseMode.Contract.BaseModeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseMode *BaseModeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseMode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseMode *BaseModeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseMode *BaseModeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseMode.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSTHIS is a free data retrieval call binding the contract method 0xcc025f7c.
//
// Solidity: function ADDRESS_THIS() view returns(address)
func (_BaseMode *BaseModeCaller) ADDRESSTHIS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMode.contract.Call(opts, &out, "ADDRESS_THIS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSTHIS is a free data retrieval call binding the contract method 0xcc025f7c.
//
// Solidity: function ADDRESS_THIS() view returns(address)
func (_BaseMode *BaseModeSession) ADDRESSTHIS() (common.Address, error) {
	return _BaseMode.Contract.ADDRESSTHIS(&_BaseMode.CallOpts)
}

// ADDRESSTHIS is a free data retrieval call binding the contract method 0xcc025f7c.
//
// Solidity: function ADDRESS_THIS() view returns(address)
func (_BaseMode *BaseModeCallerSession) ADDRESSTHIS() (common.Address, error) {
	return _BaseMode.Contract.ADDRESSTHIS(&_BaseMode.CallOpts)
}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_BaseMode *BaseModeCaller) CONFIG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMode.contract.Call(opts, &out, "CONFIG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_BaseMode *BaseModeSession) CONFIG() (common.Address, error) {
	return _BaseMode.Contract.CONFIG(&_BaseMode.CallOpts)
}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_BaseMode *BaseModeCallerSession) CONFIG() (common.Address, error) {
	return _BaseMode.Contract.CONFIG(&_BaseMode.CallOpts)
}

// ENTRYPOINT is a free data retrieval call binding the contract method 0xe8eb3cc6.
//
// Solidity: function ENTRYPOINT() view returns(address)
func (_BaseMode *BaseModeCaller) ENTRYPOINT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMode.contract.Call(opts, &out, "ENTRYPOINT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ENTRYPOINT is a free data retrieval call binding the contract method 0xe8eb3cc6.
//
// Solidity: function ENTRYPOINT() view returns(address)
func (_BaseMode *BaseModeSession) ENTRYPOINT() (common.Address, error) {
	return _BaseMode.Contract.ENTRYPOINT(&_BaseMode.CallOpts)
}

// ENTRYPOINT is a free data retrieval call binding the contract method 0xe8eb3cc6.
//
// Solidity: function ENTRYPOINT() view returns(address)
func (_BaseMode *BaseModeCallerSession) ENTRYPOINT() (common.Address, error) {
	return _BaseMode.Contract.ENTRYPOINT(&_BaseMode.CallOpts)
}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_BaseMode *BaseModeCaller) SIGVALIDATIONFAILED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseMode.contract.Call(opts, &out, "SIG_VALIDATION_FAILED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_BaseMode *BaseModeSession) SIGVALIDATIONFAILED() (*big.Int, error) {
	return _BaseMode.Contract.SIGVALIDATIONFAILED(&_BaseMode.CallOpts)
}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_BaseMode *BaseModeCallerSession) SIGVALIDATIONFAILED() (*big.Int, error) {
	return _BaseMode.Contract.SIGVALIDATIONFAILED(&_BaseMode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseMode *BaseModeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMode.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseMode *BaseModeSession) Owner() (common.Address, error) {
	return _BaseMode.Contract.Owner(&_BaseMode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseMode *BaseModeCallerSession) Owner() (common.Address, error) {
	return _BaseMode.Contract.Owner(&_BaseMode.CallOpts)
}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 requiredPreFund) view returns(bytes, uint256)
func (_BaseMode *BaseModeCaller) ValidatePaymasterUserOp(opts *bind.CallOpts, userOp PackedUserOperation, userOpHash [32]byte, requiredPreFund *big.Int) ([]byte, *big.Int, error) {
	var out []interface{}
	err := _BaseMode.contract.Call(opts, &out, "validatePaymasterUserOp", userOp, userOpHash, requiredPreFund)

	if err != nil {
		return *new([]byte), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 requiredPreFund) view returns(bytes, uint256)
func (_BaseMode *BaseModeSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, requiredPreFund *big.Int) ([]byte, *big.Int, error) {
	return _BaseMode.Contract.ValidatePaymasterUserOp(&_BaseMode.CallOpts, userOp, userOpHash, requiredPreFund)
}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 requiredPreFund) view returns(bytes, uint256)
func (_BaseMode *BaseModeCallerSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, requiredPreFund *big.Int) ([]byte, *big.Int, error) {
	return _BaseMode.Contract.ValidatePaymasterUserOp(&_BaseMode.CallOpts, userOp, userOpHash, requiredPreFund)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 postOpMode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_BaseMode *BaseModeTransactor) PostOp(opts *bind.TransactOpts, postOpMode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _BaseMode.contract.Transact(opts, "postOp", postOpMode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 postOpMode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_BaseMode *BaseModeSession) PostOp(postOpMode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _BaseMode.Contract.PostOp(&_BaseMode.TransactOpts, postOpMode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 postOpMode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_BaseMode *BaseModeTransactorSession) PostOp(postOpMode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _BaseMode.Contract.PostOp(&_BaseMode.TransactOpts, postOpMode, context, actualGasCost, actualUserOpFeePerGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaseMode *BaseModeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMode.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaseMode *BaseModeSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaseMode.Contract.RenounceOwnership(&_BaseMode.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaseMode *BaseModeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaseMode.Contract.RenounceOwnership(&_BaseMode.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseMode *BaseModeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseMode.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseMode *BaseModeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseMode.Contract.TransferOwnership(&_BaseMode.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseMode *BaseModeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseMode.Contract.TransferOwnership(&_BaseMode.TransactOpts, newOwner)
}

// WithdrawGas is a paid mutator transaction binding the contract method 0x813f3f44.
//
// Solidity: function withdrawGas(address recipient, uint256 amount) returns()
func (_BaseMode *BaseModeTransactor) WithdrawGas(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseMode.contract.Transact(opts, "withdrawGas", recipient, amount)
}

// WithdrawGas is a paid mutator transaction binding the contract method 0x813f3f44.
//
// Solidity: function withdrawGas(address recipient, uint256 amount) returns()
func (_BaseMode *BaseModeSession) WithdrawGas(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseMode.Contract.WithdrawGas(&_BaseMode.TransactOpts, recipient, amount)
}

// WithdrawGas is a paid mutator transaction binding the contract method 0x813f3f44.
//
// Solidity: function withdrawGas(address recipient, uint256 amount) returns()
func (_BaseMode *BaseModeTransactorSession) WithdrawGas(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseMode.Contract.WithdrawGas(&_BaseMode.TransactOpts, recipient, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BaseMode *BaseModeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMode.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BaseMode *BaseModeSession) Receive() (*types.Transaction, error) {
	return _BaseMode.Contract.Receive(&_BaseMode.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BaseMode *BaseModeTransactorSession) Receive() (*types.Transaction, error) {
	return _BaseMode.Contract.Receive(&_BaseMode.TransactOpts)
}

// BaseModeGasDepositedIterator is returned from FilterGasDeposited and is used to iterate over the raw logs and unpacked data for GasDeposited events raised by the BaseMode contract.
type BaseModeGasDepositedIterator struct {
	Event *BaseModeGasDeposited // Event containing the contract specifics and raw log

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
func (it *BaseModeGasDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseModeGasDeposited)
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
		it.Event = new(BaseModeGasDeposited)
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
func (it *BaseModeGasDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseModeGasDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseModeGasDeposited represents a GasDeposited event raised by the BaseMode contract.
type BaseModeGasDeposited struct {
	EntryPoint common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasDeposited is a free log retrieval operation binding the contract event 0x1dbbf474736d6415d6a265fabee708fe6e988f6fd0c9d870ded36cab380898dd.
//
// Solidity: event GasDeposited(address entryPoint, uint256 amount)
func (_BaseMode *BaseModeFilterer) FilterGasDeposited(opts *bind.FilterOpts) (*BaseModeGasDepositedIterator, error) {

	logs, sub, err := _BaseMode.contract.FilterLogs(opts, "GasDeposited")
	if err != nil {
		return nil, err
	}
	return &BaseModeGasDepositedIterator{contract: _BaseMode.contract, event: "GasDeposited", logs: logs, sub: sub}, nil
}

// WatchGasDeposited is a free log subscription operation binding the contract event 0x1dbbf474736d6415d6a265fabee708fe6e988f6fd0c9d870ded36cab380898dd.
//
// Solidity: event GasDeposited(address entryPoint, uint256 amount)
func (_BaseMode *BaseModeFilterer) WatchGasDeposited(opts *bind.WatchOpts, sink chan<- *BaseModeGasDeposited) (event.Subscription, error) {

	logs, sub, err := _BaseMode.contract.WatchLogs(opts, "GasDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseModeGasDeposited)
				if err := _BaseMode.contract.UnpackLog(event, "GasDeposited", log); err != nil {
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

// ParseGasDeposited is a log parse operation binding the contract event 0x1dbbf474736d6415d6a265fabee708fe6e988f6fd0c9d870ded36cab380898dd.
//
// Solidity: event GasDeposited(address entryPoint, uint256 amount)
func (_BaseMode *BaseModeFilterer) ParseGasDeposited(log types.Log) (*BaseModeGasDeposited, error) {
	event := new(BaseModeGasDeposited)
	if err := _BaseMode.contract.UnpackLog(event, "GasDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseModeGasWithdrawnIterator is returned from FilterGasWithdrawn and is used to iterate over the raw logs and unpacked data for GasWithdrawn events raised by the BaseMode contract.
type BaseModeGasWithdrawnIterator struct {
	Event *BaseModeGasWithdrawn // Event containing the contract specifics and raw log

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
func (it *BaseModeGasWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseModeGasWithdrawn)
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
		it.Event = new(BaseModeGasWithdrawn)
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
func (it *BaseModeGasWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseModeGasWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseModeGasWithdrawn represents a GasWithdrawn event raised by the BaseMode contract.
type BaseModeGasWithdrawn struct {
	EntryPoint common.Address
	Recipient  common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasWithdrawn is a free log retrieval operation binding the contract event 0x926a144b6fffc1d73f115b81af7ec66a7c12aed0ff73197c39a683753fc1d925.
//
// Solidity: event GasWithdrawn(address entryPoint, address recipient, uint256 amount)
func (_BaseMode *BaseModeFilterer) FilterGasWithdrawn(opts *bind.FilterOpts) (*BaseModeGasWithdrawnIterator, error) {

	logs, sub, err := _BaseMode.contract.FilterLogs(opts, "GasWithdrawn")
	if err != nil {
		return nil, err
	}
	return &BaseModeGasWithdrawnIterator{contract: _BaseMode.contract, event: "GasWithdrawn", logs: logs, sub: sub}, nil
}

// WatchGasWithdrawn is a free log subscription operation binding the contract event 0x926a144b6fffc1d73f115b81af7ec66a7c12aed0ff73197c39a683753fc1d925.
//
// Solidity: event GasWithdrawn(address entryPoint, address recipient, uint256 amount)
func (_BaseMode *BaseModeFilterer) WatchGasWithdrawn(opts *bind.WatchOpts, sink chan<- *BaseModeGasWithdrawn) (event.Subscription, error) {

	logs, sub, err := _BaseMode.contract.WatchLogs(opts, "GasWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseModeGasWithdrawn)
				if err := _BaseMode.contract.UnpackLog(event, "GasWithdrawn", log); err != nil {
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

// ParseGasWithdrawn is a log parse operation binding the contract event 0x926a144b6fffc1d73f115b81af7ec66a7c12aed0ff73197c39a683753fc1d925.
//
// Solidity: event GasWithdrawn(address entryPoint, address recipient, uint256 amount)
func (_BaseMode *BaseModeFilterer) ParseGasWithdrawn(log types.Log) (*BaseModeGasWithdrawn, error) {
	event := new(BaseModeGasWithdrawn)
	if err := _BaseMode.contract.UnpackLog(event, "GasWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseModeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BaseMode contract.
type BaseModeOwnershipTransferredIterator struct {
	Event *BaseModeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BaseModeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseModeOwnershipTransferred)
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
		it.Event = new(BaseModeOwnershipTransferred)
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
func (it *BaseModeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseModeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseModeOwnershipTransferred represents a OwnershipTransferred event raised by the BaseMode contract.
type BaseModeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BaseMode *BaseModeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaseModeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaseMode.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaseModeOwnershipTransferredIterator{contract: _BaseMode.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BaseMode *BaseModeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaseModeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaseMode.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseModeOwnershipTransferred)
				if err := _BaseMode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BaseMode *BaseModeFilterer) ParseOwnershipTransferred(log types.Log) (*BaseModeOwnershipTransferred, error) {
	event := new(BaseModeOwnershipTransferred)
	if err := _BaseMode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseModeSetConfigIterator is returned from FilterSetConfig and is used to iterate over the raw logs and unpacked data for SetConfig events raised by the BaseMode contract.
type BaseModeSetConfigIterator struct {
	Event *BaseModeSetConfig // Event containing the contract specifics and raw log

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
func (it *BaseModeSetConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseModeSetConfig)
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
		it.Event = new(BaseModeSetConfig)
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
func (it *BaseModeSetConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseModeSetConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseModeSetConfig represents a SetConfig event raised by the BaseMode contract.
type BaseModeSetConfig struct {
	Config common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetConfig is a free log retrieval operation binding the contract event 0xc5618716db99966ac0bedb011a55472827d54343d73b50c3118c0b03cdf1c75f.
//
// Solidity: event SetConfig(address config)
func (_BaseMode *BaseModeFilterer) FilterSetConfig(opts *bind.FilterOpts) (*BaseModeSetConfigIterator, error) {

	logs, sub, err := _BaseMode.contract.FilterLogs(opts, "SetConfig")
	if err != nil {
		return nil, err
	}
	return &BaseModeSetConfigIterator{contract: _BaseMode.contract, event: "SetConfig", logs: logs, sub: sub}, nil
}

// WatchSetConfig is a free log subscription operation binding the contract event 0xc5618716db99966ac0bedb011a55472827d54343d73b50c3118c0b03cdf1c75f.
//
// Solidity: event SetConfig(address config)
func (_BaseMode *BaseModeFilterer) WatchSetConfig(opts *bind.WatchOpts, sink chan<- *BaseModeSetConfig) (event.Subscription, error) {

	logs, sub, err := _BaseMode.contract.WatchLogs(opts, "SetConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseModeSetConfig)
				if err := _BaseMode.contract.UnpackLog(event, "SetConfig", log); err != nil {
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

// ParseSetConfig is a log parse operation binding the contract event 0xc5618716db99966ac0bedb011a55472827d54343d73b50c3118c0b03cdf1c75f.
//
// Solidity: event SetConfig(address config)
func (_BaseMode *BaseModeFilterer) ParseSetConfig(log types.Log) (*BaseModeSetConfig, error) {
	event := new(BaseModeSetConfig)
	if err := _BaseMode.contract.UnpackLog(event, "SetConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseModeSetGasManagerIterator is returned from FilterSetGasManager and is used to iterate over the raw logs and unpacked data for SetGasManager events raised by the BaseMode contract.
type BaseModeSetGasManagerIterator struct {
	Event *BaseModeSetGasManager // Event containing the contract specifics and raw log

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
func (it *BaseModeSetGasManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseModeSetGasManager)
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
		it.Event = new(BaseModeSetGasManager)
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
func (it *BaseModeSetGasManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseModeSetGasManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseModeSetGasManager represents a SetGasManager event raised by the BaseMode contract.
type BaseModeSetGasManager struct {
	GasManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetGasManager is a free log retrieval operation binding the contract event 0x9daebb79792c29c1b590ef48a1d6cb59a9150c1d281fc94b12724321615e203a.
//
// Solidity: event SetGasManager(address gasManager)
func (_BaseMode *BaseModeFilterer) FilterSetGasManager(opts *bind.FilterOpts) (*BaseModeSetGasManagerIterator, error) {

	logs, sub, err := _BaseMode.contract.FilterLogs(opts, "SetGasManager")
	if err != nil {
		return nil, err
	}
	return &BaseModeSetGasManagerIterator{contract: _BaseMode.contract, event: "SetGasManager", logs: logs, sub: sub}, nil
}

// WatchSetGasManager is a free log subscription operation binding the contract event 0x9daebb79792c29c1b590ef48a1d6cb59a9150c1d281fc94b12724321615e203a.
//
// Solidity: event SetGasManager(address gasManager)
func (_BaseMode *BaseModeFilterer) WatchSetGasManager(opts *bind.WatchOpts, sink chan<- *BaseModeSetGasManager) (event.Subscription, error) {

	logs, sub, err := _BaseMode.contract.WatchLogs(opts, "SetGasManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseModeSetGasManager)
				if err := _BaseMode.contract.UnpackLog(event, "SetGasManager", log); err != nil {
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

// ParseSetGasManager is a log parse operation binding the contract event 0x9daebb79792c29c1b590ef48a1d6cb59a9150c1d281fc94b12724321615e203a.
//
// Solidity: event SetGasManager(address gasManager)
func (_BaseMode *BaseModeFilterer) ParseSetGasManager(log types.Log) (*BaseModeSetGasManager, error) {
	event := new(BaseModeSetGasManager)
	if err := _BaseMode.contract.UnpackLog(event, "SetGasManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
