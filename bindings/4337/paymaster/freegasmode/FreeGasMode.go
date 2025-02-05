// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package freegasmode

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

// FreeGasModeMetaData contains all meta data concerning the FreeGasMode contract.
var FreeGasModeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_config\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"config\",\"type\":\"address\"}],\"name\":\"SetConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasManager\",\"type\":\"address\"}],\"name\":\"SetGasManager\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_THIS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIG\",\"outputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ENTRYPOINT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIG_VALIDATION_FAILED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"postOpMode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"requiredPreFund\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e03461011657601f61079438819003918201601f19168301916001600160401b0383118484101761011b578084926060946040528339810103126101165761004781610131565b61005f604061005860208501610131565b9301610131565b6001600160a01b039182169283156100fd57600080546001600160a01b03198116861782556040519591908516907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a33060c05260a0521660805261064e90816101468239608051818181610229015261048b015260a0518181816028015281816101e10152610313015260c0518181816069015261026d0152f35b604051631e4fbdf760e01b815260006004820152602490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101165756fe60806040818152600491823610156100b9575b50361561001e57600080fd5b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811692833b156100b45760246000928451958693849263b760faf960e01b84527f0000000000000000000000000000000000000000000000000000000000000000169083015234905af19081156100aa575061009f57005b6100a8906105a6565b005b513d6000823e3d90fd5b600080fd5b600090813560e01c90816352b7512c1461043857508063715018a6146103de5780637c627b2114610382578063813f3f44146102e05780638da5cb5b146102b85780638f41ec5a1461029c578063cc025f7c14610258578063d92e82e414610214578063e8eb3cc6146101cc5763f2fde38b03610012579190346101c85760203660031901126101c8576001600160a01b038235818116939192908490036101c4576101636105ec565b83156101ae575050600054826bffffffffffffffffffffffff60a01b821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a380f35b51631e4fbdf760e01b8152908101849052602490fd5b8480fd5b8280fd5b509034610210578160031936011261021057517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b509034610210578160031936011261021057517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b509034610210578160031936011261021057517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5090346102105781600319360112610210576020905160018152f35b509034610210578160031936011261021057905490516001600160a01b039091168152602090f35b509190346101c857806003193601126101c8578282356001600160a01b03818116918290036101c8576103116105ec565b7f00000000000000000000000000000000000000000000000000000000000000001693843b156101c857604490838551968794859363040b850f60e31b855284015260243560248401525af1908115610379575061036d575080f35b610376906105a6565b80f35b513d84823e3d90fd5b50823461021057608036600319011261021057600381351015610210576024359067ffffffffffffffff908183116103da57366023840112156103da578201359081116101c857369101602401116103d75780f35b80fd5b8380fd5b50346103d757806003193601126103d7576103f76105ec565b80546001600160a01b03198116825581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b92919050346103d757606092600319906060823601126101c85767ffffffffffffffff9186358381116101c4579061012091360301126101c857639365878360e01b8152328682015260209586826024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa92831561059c57849361053d575b5050506000146105265781516104d8816105d0565b81815281905b83519484865281518095870152825b858110610513578660608188888c89858486010152830152601f80199101168101030190f35b82810188015187820183015287016104ed565b8151610531816105d0565b818152906001916104de565b9091925086913d8811610594575b601f8301601f191684019182118483101761058157508691839186528101031261021057518015158103610210573880806104c3565b634e487b7160e01b855260419052602484fd5b3d925061054b565b85513d86823e3d90fd5b67ffffffffffffffff81116105ba57604052565b634e487b7160e01b600052604160045260246000fd5b6020810190811067ffffffffffffffff8211176105ba57604052565b6000546001600160a01b0316330361060057565b60405163118cdaa760e01b8152336004820152602490fdfea2646970667358221220aafb720f0aa2dbd6cc99adf59e61ce943570faeb18203c1ac4f4351d5494744464736f6c63430008190033",
}

// FreeGasModeABI is the input ABI used to generate the binding from.
// Deprecated: Use FreeGasModeMetaData.ABI instead.
var FreeGasModeABI = FreeGasModeMetaData.ABI

// FreeGasModeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FreeGasModeMetaData.Bin instead.
var FreeGasModeBin = FreeGasModeMetaData.Bin

// DeployFreeGasMode deploys a new Ethereum contract, binding an instance of FreeGasMode to it.
func DeployFreeGasMode(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _entryPoint common.Address, _config common.Address) (common.Address, *types.Transaction, *FreeGasMode, error) {
	parsed, err := FreeGasModeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FreeGasModeBin), backend, _owner, _entryPoint, _config)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FreeGasMode{FreeGasModeCaller: FreeGasModeCaller{contract: contract}, FreeGasModeTransactor: FreeGasModeTransactor{contract: contract}, FreeGasModeFilterer: FreeGasModeFilterer{contract: contract}}, nil
}

// FreeGasMode is an auto generated Go binding around an Ethereum contract.
type FreeGasMode struct {
	FreeGasModeCaller     // Read-only binding to the contract
	FreeGasModeTransactor // Write-only binding to the contract
	FreeGasModeFilterer   // Log filterer for contract events
}

// FreeGasModeCaller is an auto generated read-only Go binding around an Ethereum contract.
type FreeGasModeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreeGasModeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FreeGasModeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreeGasModeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FreeGasModeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreeGasModeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FreeGasModeSession struct {
	Contract     *FreeGasMode      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FreeGasModeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FreeGasModeCallerSession struct {
	Contract *FreeGasModeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FreeGasModeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FreeGasModeTransactorSession struct {
	Contract     *FreeGasModeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FreeGasModeRaw is an auto generated low-level Go binding around an Ethereum contract.
type FreeGasModeRaw struct {
	Contract *FreeGasMode // Generic contract binding to access the raw methods on
}

// FreeGasModeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FreeGasModeCallerRaw struct {
	Contract *FreeGasModeCaller // Generic read-only contract binding to access the raw methods on
}

// FreeGasModeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FreeGasModeTransactorRaw struct {
	Contract *FreeGasModeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFreeGasMode creates a new instance of FreeGasMode, bound to a specific deployed contract.
func NewFreeGasMode(address common.Address, backend bind.ContractBackend) (*FreeGasMode, error) {
	contract, err := bindFreeGasMode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FreeGasMode{FreeGasModeCaller: FreeGasModeCaller{contract: contract}, FreeGasModeTransactor: FreeGasModeTransactor{contract: contract}, FreeGasModeFilterer: FreeGasModeFilterer{contract: contract}}, nil
}

// NewFreeGasModeCaller creates a new read-only instance of FreeGasMode, bound to a specific deployed contract.
func NewFreeGasModeCaller(address common.Address, caller bind.ContractCaller) (*FreeGasModeCaller, error) {
	contract, err := bindFreeGasMode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FreeGasModeCaller{contract: contract}, nil
}

// NewFreeGasModeTransactor creates a new write-only instance of FreeGasMode, bound to a specific deployed contract.
func NewFreeGasModeTransactor(address common.Address, transactor bind.ContractTransactor) (*FreeGasModeTransactor, error) {
	contract, err := bindFreeGasMode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FreeGasModeTransactor{contract: contract}, nil
}

// NewFreeGasModeFilterer creates a new log filterer instance of FreeGasMode, bound to a specific deployed contract.
func NewFreeGasModeFilterer(address common.Address, filterer bind.ContractFilterer) (*FreeGasModeFilterer, error) {
	contract, err := bindFreeGasMode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FreeGasModeFilterer{contract: contract}, nil
}

// bindFreeGasMode binds a generic wrapper to an already deployed contract.
func bindFreeGasMode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FreeGasModeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreeGasMode *FreeGasModeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreeGasMode.Contract.FreeGasModeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreeGasMode *FreeGasModeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeGasMode.Contract.FreeGasModeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreeGasMode *FreeGasModeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreeGasMode.Contract.FreeGasModeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreeGasMode *FreeGasModeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreeGasMode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreeGasMode *FreeGasModeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeGasMode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreeGasMode *FreeGasModeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreeGasMode.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSTHIS is a free data retrieval call binding the contract method 0xcc025f7c.
//
// Solidity: function ADDRESS_THIS() view returns(address)
func (_FreeGasMode *FreeGasModeCaller) ADDRESSTHIS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FreeGasMode.contract.Call(opts, &out, "ADDRESS_THIS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSTHIS is a free data retrieval call binding the contract method 0xcc025f7c.
//
// Solidity: function ADDRESS_THIS() view returns(address)
func (_FreeGasMode *FreeGasModeSession) ADDRESSTHIS() (common.Address, error) {
	return _FreeGasMode.Contract.ADDRESSTHIS(&_FreeGasMode.CallOpts)
}

// ADDRESSTHIS is a free data retrieval call binding the contract method 0xcc025f7c.
//
// Solidity: function ADDRESS_THIS() view returns(address)
func (_FreeGasMode *FreeGasModeCallerSession) ADDRESSTHIS() (common.Address, error) {
	return _FreeGasMode.Contract.ADDRESSTHIS(&_FreeGasMode.CallOpts)
}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_FreeGasMode *FreeGasModeCaller) CONFIG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FreeGasMode.contract.Call(opts, &out, "CONFIG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_FreeGasMode *FreeGasModeSession) CONFIG() (common.Address, error) {
	return _FreeGasMode.Contract.CONFIG(&_FreeGasMode.CallOpts)
}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_FreeGasMode *FreeGasModeCallerSession) CONFIG() (common.Address, error) {
	return _FreeGasMode.Contract.CONFIG(&_FreeGasMode.CallOpts)
}

// ENTRYPOINT is a free data retrieval call binding the contract method 0xe8eb3cc6.
//
// Solidity: function ENTRYPOINT() view returns(address)
func (_FreeGasMode *FreeGasModeCaller) ENTRYPOINT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FreeGasMode.contract.Call(opts, &out, "ENTRYPOINT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ENTRYPOINT is a free data retrieval call binding the contract method 0xe8eb3cc6.
//
// Solidity: function ENTRYPOINT() view returns(address)
func (_FreeGasMode *FreeGasModeSession) ENTRYPOINT() (common.Address, error) {
	return _FreeGasMode.Contract.ENTRYPOINT(&_FreeGasMode.CallOpts)
}

// ENTRYPOINT is a free data retrieval call binding the contract method 0xe8eb3cc6.
//
// Solidity: function ENTRYPOINT() view returns(address)
func (_FreeGasMode *FreeGasModeCallerSession) ENTRYPOINT() (common.Address, error) {
	return _FreeGasMode.Contract.ENTRYPOINT(&_FreeGasMode.CallOpts)
}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_FreeGasMode *FreeGasModeCaller) SIGVALIDATIONFAILED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeGasMode.contract.Call(opts, &out, "SIG_VALIDATION_FAILED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_FreeGasMode *FreeGasModeSession) SIGVALIDATIONFAILED() (*big.Int, error) {
	return _FreeGasMode.Contract.SIGVALIDATIONFAILED(&_FreeGasMode.CallOpts)
}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_FreeGasMode *FreeGasModeCallerSession) SIGVALIDATIONFAILED() (*big.Int, error) {
	return _FreeGasMode.Contract.SIGVALIDATIONFAILED(&_FreeGasMode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FreeGasMode *FreeGasModeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FreeGasMode.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FreeGasMode *FreeGasModeSession) Owner() (common.Address, error) {
	return _FreeGasMode.Contract.Owner(&_FreeGasMode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FreeGasMode *FreeGasModeCallerSession) Owner() (common.Address, error) {
	return _FreeGasMode.Contract.Owner(&_FreeGasMode.CallOpts)
}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) , bytes32 userOpHash, uint256 requiredPreFund) view returns(bytes context, uint256 validationData)
func (_FreeGasMode *FreeGasModeCaller) ValidatePaymasterUserOp(opts *bind.CallOpts, arg0 PackedUserOperation, userOpHash [32]byte, requiredPreFund *big.Int) (struct {
	Context        []byte
	ValidationData *big.Int
}, error) {
	var out []interface{}
	err := _FreeGasMode.contract.Call(opts, &out, "validatePaymasterUserOp", arg0, userOpHash, requiredPreFund)

	outstruct := new(struct {
		Context        []byte
		ValidationData *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Context = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ValidationData = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) , bytes32 userOpHash, uint256 requiredPreFund) view returns(bytes context, uint256 validationData)
func (_FreeGasMode *FreeGasModeSession) ValidatePaymasterUserOp(arg0 PackedUserOperation, userOpHash [32]byte, requiredPreFund *big.Int) (struct {
	Context        []byte
	ValidationData *big.Int
}, error) {
	return _FreeGasMode.Contract.ValidatePaymasterUserOp(&_FreeGasMode.CallOpts, arg0, userOpHash, requiredPreFund)
}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) , bytes32 userOpHash, uint256 requiredPreFund) view returns(bytes context, uint256 validationData)
func (_FreeGasMode *FreeGasModeCallerSession) ValidatePaymasterUserOp(arg0 PackedUserOperation, userOpHash [32]byte, requiredPreFund *big.Int) (struct {
	Context        []byte
	ValidationData *big.Int
}, error) {
	return _FreeGasMode.Contract.ValidatePaymasterUserOp(&_FreeGasMode.CallOpts, arg0, userOpHash, requiredPreFund)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 postOpMode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_FreeGasMode *FreeGasModeTransactor) PostOp(opts *bind.TransactOpts, postOpMode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _FreeGasMode.contract.Transact(opts, "postOp", postOpMode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 postOpMode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_FreeGasMode *FreeGasModeSession) PostOp(postOpMode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _FreeGasMode.Contract.PostOp(&_FreeGasMode.TransactOpts, postOpMode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 postOpMode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_FreeGasMode *FreeGasModeTransactorSession) PostOp(postOpMode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _FreeGasMode.Contract.PostOp(&_FreeGasMode.TransactOpts, postOpMode, context, actualGasCost, actualUserOpFeePerGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FreeGasMode *FreeGasModeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeGasMode.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FreeGasMode *FreeGasModeSession) RenounceOwnership() (*types.Transaction, error) {
	return _FreeGasMode.Contract.RenounceOwnership(&_FreeGasMode.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FreeGasMode *FreeGasModeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FreeGasMode.Contract.RenounceOwnership(&_FreeGasMode.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FreeGasMode *FreeGasModeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FreeGasMode.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FreeGasMode *FreeGasModeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FreeGasMode.Contract.TransferOwnership(&_FreeGasMode.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FreeGasMode *FreeGasModeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FreeGasMode.Contract.TransferOwnership(&_FreeGasMode.TransactOpts, newOwner)
}

// WithdrawGas is a paid mutator transaction binding the contract method 0x813f3f44.
//
// Solidity: function withdrawGas(address recipient, uint256 amount) returns()
func (_FreeGasMode *FreeGasModeTransactor) WithdrawGas(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FreeGasMode.contract.Transact(opts, "withdrawGas", recipient, amount)
}

// WithdrawGas is a paid mutator transaction binding the contract method 0x813f3f44.
//
// Solidity: function withdrawGas(address recipient, uint256 amount) returns()
func (_FreeGasMode *FreeGasModeSession) WithdrawGas(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FreeGasMode.Contract.WithdrawGas(&_FreeGasMode.TransactOpts, recipient, amount)
}

// WithdrawGas is a paid mutator transaction binding the contract method 0x813f3f44.
//
// Solidity: function withdrawGas(address recipient, uint256 amount) returns()
func (_FreeGasMode *FreeGasModeTransactorSession) WithdrawGas(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FreeGasMode.Contract.WithdrawGas(&_FreeGasMode.TransactOpts, recipient, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FreeGasMode *FreeGasModeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeGasMode.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FreeGasMode *FreeGasModeSession) Receive() (*types.Transaction, error) {
	return _FreeGasMode.Contract.Receive(&_FreeGasMode.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FreeGasMode *FreeGasModeTransactorSession) Receive() (*types.Transaction, error) {
	return _FreeGasMode.Contract.Receive(&_FreeGasMode.TransactOpts)
}

// FreeGasModeGasDepositedIterator is returned from FilterGasDeposited and is used to iterate over the raw logs and unpacked data for GasDeposited events raised by the FreeGasMode contract.
type FreeGasModeGasDepositedIterator struct {
	Event *FreeGasModeGasDeposited // Event containing the contract specifics and raw log

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
func (it *FreeGasModeGasDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreeGasModeGasDeposited)
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
		it.Event = new(FreeGasModeGasDeposited)
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
func (it *FreeGasModeGasDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreeGasModeGasDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreeGasModeGasDeposited represents a GasDeposited event raised by the FreeGasMode contract.
type FreeGasModeGasDeposited struct {
	EntryPoint common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasDeposited is a free log retrieval operation binding the contract event 0x1dbbf474736d6415d6a265fabee708fe6e988f6fd0c9d870ded36cab380898dd.
//
// Solidity: event GasDeposited(address entryPoint, uint256 amount)
func (_FreeGasMode *FreeGasModeFilterer) FilterGasDeposited(opts *bind.FilterOpts) (*FreeGasModeGasDepositedIterator, error) {

	logs, sub, err := _FreeGasMode.contract.FilterLogs(opts, "GasDeposited")
	if err != nil {
		return nil, err
	}
	return &FreeGasModeGasDepositedIterator{contract: _FreeGasMode.contract, event: "GasDeposited", logs: logs, sub: sub}, nil
}

// WatchGasDeposited is a free log subscription operation binding the contract event 0x1dbbf474736d6415d6a265fabee708fe6e988f6fd0c9d870ded36cab380898dd.
//
// Solidity: event GasDeposited(address entryPoint, uint256 amount)
func (_FreeGasMode *FreeGasModeFilterer) WatchGasDeposited(opts *bind.WatchOpts, sink chan<- *FreeGasModeGasDeposited) (event.Subscription, error) {

	logs, sub, err := _FreeGasMode.contract.WatchLogs(opts, "GasDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreeGasModeGasDeposited)
				if err := _FreeGasMode.contract.UnpackLog(event, "GasDeposited", log); err != nil {
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
func (_FreeGasMode *FreeGasModeFilterer) ParseGasDeposited(log types.Log) (*FreeGasModeGasDeposited, error) {
	event := new(FreeGasModeGasDeposited)
	if err := _FreeGasMode.contract.UnpackLog(event, "GasDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreeGasModeGasWithdrawnIterator is returned from FilterGasWithdrawn and is used to iterate over the raw logs and unpacked data for GasWithdrawn events raised by the FreeGasMode contract.
type FreeGasModeGasWithdrawnIterator struct {
	Event *FreeGasModeGasWithdrawn // Event containing the contract specifics and raw log

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
func (it *FreeGasModeGasWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreeGasModeGasWithdrawn)
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
		it.Event = new(FreeGasModeGasWithdrawn)
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
func (it *FreeGasModeGasWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreeGasModeGasWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreeGasModeGasWithdrawn represents a GasWithdrawn event raised by the FreeGasMode contract.
type FreeGasModeGasWithdrawn struct {
	EntryPoint common.Address
	Recipient  common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasWithdrawn is a free log retrieval operation binding the contract event 0x926a144b6fffc1d73f115b81af7ec66a7c12aed0ff73197c39a683753fc1d925.
//
// Solidity: event GasWithdrawn(address entryPoint, address recipient, uint256 amount)
func (_FreeGasMode *FreeGasModeFilterer) FilterGasWithdrawn(opts *bind.FilterOpts) (*FreeGasModeGasWithdrawnIterator, error) {

	logs, sub, err := _FreeGasMode.contract.FilterLogs(opts, "GasWithdrawn")
	if err != nil {
		return nil, err
	}
	return &FreeGasModeGasWithdrawnIterator{contract: _FreeGasMode.contract, event: "GasWithdrawn", logs: logs, sub: sub}, nil
}

// WatchGasWithdrawn is a free log subscription operation binding the contract event 0x926a144b6fffc1d73f115b81af7ec66a7c12aed0ff73197c39a683753fc1d925.
//
// Solidity: event GasWithdrawn(address entryPoint, address recipient, uint256 amount)
func (_FreeGasMode *FreeGasModeFilterer) WatchGasWithdrawn(opts *bind.WatchOpts, sink chan<- *FreeGasModeGasWithdrawn) (event.Subscription, error) {

	logs, sub, err := _FreeGasMode.contract.WatchLogs(opts, "GasWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreeGasModeGasWithdrawn)
				if err := _FreeGasMode.contract.UnpackLog(event, "GasWithdrawn", log); err != nil {
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
func (_FreeGasMode *FreeGasModeFilterer) ParseGasWithdrawn(log types.Log) (*FreeGasModeGasWithdrawn, error) {
	event := new(FreeGasModeGasWithdrawn)
	if err := _FreeGasMode.contract.UnpackLog(event, "GasWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreeGasModeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FreeGasMode contract.
type FreeGasModeOwnershipTransferredIterator struct {
	Event *FreeGasModeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FreeGasModeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreeGasModeOwnershipTransferred)
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
		it.Event = new(FreeGasModeOwnershipTransferred)
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
func (it *FreeGasModeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreeGasModeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreeGasModeOwnershipTransferred represents a OwnershipTransferred event raised by the FreeGasMode contract.
type FreeGasModeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FreeGasMode *FreeGasModeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FreeGasModeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FreeGasMode.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FreeGasModeOwnershipTransferredIterator{contract: _FreeGasMode.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FreeGasMode *FreeGasModeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FreeGasModeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FreeGasMode.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreeGasModeOwnershipTransferred)
				if err := _FreeGasMode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FreeGasMode *FreeGasModeFilterer) ParseOwnershipTransferred(log types.Log) (*FreeGasModeOwnershipTransferred, error) {
	event := new(FreeGasModeOwnershipTransferred)
	if err := _FreeGasMode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreeGasModeSetConfigIterator is returned from FilterSetConfig and is used to iterate over the raw logs and unpacked data for SetConfig events raised by the FreeGasMode contract.
type FreeGasModeSetConfigIterator struct {
	Event *FreeGasModeSetConfig // Event containing the contract specifics and raw log

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
func (it *FreeGasModeSetConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreeGasModeSetConfig)
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
		it.Event = new(FreeGasModeSetConfig)
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
func (it *FreeGasModeSetConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreeGasModeSetConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreeGasModeSetConfig represents a SetConfig event raised by the FreeGasMode contract.
type FreeGasModeSetConfig struct {
	Config common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetConfig is a free log retrieval operation binding the contract event 0xc5618716db99966ac0bedb011a55472827d54343d73b50c3118c0b03cdf1c75f.
//
// Solidity: event SetConfig(address config)
func (_FreeGasMode *FreeGasModeFilterer) FilterSetConfig(opts *bind.FilterOpts) (*FreeGasModeSetConfigIterator, error) {

	logs, sub, err := _FreeGasMode.contract.FilterLogs(opts, "SetConfig")
	if err != nil {
		return nil, err
	}
	return &FreeGasModeSetConfigIterator{contract: _FreeGasMode.contract, event: "SetConfig", logs: logs, sub: sub}, nil
}

// WatchSetConfig is a free log subscription operation binding the contract event 0xc5618716db99966ac0bedb011a55472827d54343d73b50c3118c0b03cdf1c75f.
//
// Solidity: event SetConfig(address config)
func (_FreeGasMode *FreeGasModeFilterer) WatchSetConfig(opts *bind.WatchOpts, sink chan<- *FreeGasModeSetConfig) (event.Subscription, error) {

	logs, sub, err := _FreeGasMode.contract.WatchLogs(opts, "SetConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreeGasModeSetConfig)
				if err := _FreeGasMode.contract.UnpackLog(event, "SetConfig", log); err != nil {
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
func (_FreeGasMode *FreeGasModeFilterer) ParseSetConfig(log types.Log) (*FreeGasModeSetConfig, error) {
	event := new(FreeGasModeSetConfig)
	if err := _FreeGasMode.contract.UnpackLog(event, "SetConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreeGasModeSetGasManagerIterator is returned from FilterSetGasManager and is used to iterate over the raw logs and unpacked data for SetGasManager events raised by the FreeGasMode contract.
type FreeGasModeSetGasManagerIterator struct {
	Event *FreeGasModeSetGasManager // Event containing the contract specifics and raw log

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
func (it *FreeGasModeSetGasManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreeGasModeSetGasManager)
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
		it.Event = new(FreeGasModeSetGasManager)
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
func (it *FreeGasModeSetGasManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreeGasModeSetGasManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreeGasModeSetGasManager represents a SetGasManager event raised by the FreeGasMode contract.
type FreeGasModeSetGasManager struct {
	GasManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetGasManager is a free log retrieval operation binding the contract event 0x9daebb79792c29c1b590ef48a1d6cb59a9150c1d281fc94b12724321615e203a.
//
// Solidity: event SetGasManager(address gasManager)
func (_FreeGasMode *FreeGasModeFilterer) FilterSetGasManager(opts *bind.FilterOpts) (*FreeGasModeSetGasManagerIterator, error) {

	logs, sub, err := _FreeGasMode.contract.FilterLogs(opts, "SetGasManager")
	if err != nil {
		return nil, err
	}
	return &FreeGasModeSetGasManagerIterator{contract: _FreeGasMode.contract, event: "SetGasManager", logs: logs, sub: sub}, nil
}

// WatchSetGasManager is a free log subscription operation binding the contract event 0x9daebb79792c29c1b590ef48a1d6cb59a9150c1d281fc94b12724321615e203a.
//
// Solidity: event SetGasManager(address gasManager)
func (_FreeGasMode *FreeGasModeFilterer) WatchSetGasManager(opts *bind.WatchOpts, sink chan<- *FreeGasModeSetGasManager) (event.Subscription, error) {

	logs, sub, err := _FreeGasMode.contract.WatchLogs(opts, "SetGasManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreeGasModeSetGasManager)
				if err := _FreeGasMode.contract.UnpackLog(event, "SetGasManager", log); err != nil {
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
func (_FreeGasMode *FreeGasModeFilterer) ParseSetGasManager(log types.Log) (*FreeGasModeSetGasManager, error) {
	event := new(FreeGasModeSetGasManager)
	if err := _FreeGasMode.contract.UnpackLog(event, "SetGasManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
