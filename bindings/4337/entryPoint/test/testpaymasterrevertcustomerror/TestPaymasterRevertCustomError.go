// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testpaymasterrevertcustomerror

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

// TestPaymasterRevertCustomErrorMetaData contains all meta data concerning the TestPaymasterRevertCustomError contract.
var TestPaymasterRevertCustomErrorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"customReason\",\"type\":\"string\"}],\"name\":\"CustomError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTestPaymasterRevertCustomError.RevertType\",\"name\":\"_revertType\",\"type\":\"uint8\"}],\"name\":\"setRevertType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604090808252346101b35780610aed803803809161001f82856101b8565b83396020928391810103126101b35751906001600160a01b03808316908184036101b357331561019b5760008054336001600160a01b031982168117835587519492938693869360249385939091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08880a36301ffc9a760e01b825263122a0e9b60e31b60048301525afa91821561018f57819261014e575b50501561010b5750608052516108fb90816101f28239608051818181610110015281816101ca015281816104510152818161049d01528181610525015281816105aa0152818161063701526108580152f35b60649083519062461bcd60e51b82526004820152601e60248201527f49456e747279506f696e7420696e74657266616365206d69736d6174636800006044820152fd5b9091508281813d8311610188575b61016681836101b8565b810103126101845751908115158203610181575038806100b9565b80fd5b5080fd5b503d61015c565b508451903d90823e3d90fd5b8451631e4fbdf760e01b815260006004820152602490fd5b600080fd5b601f909101601f19168101906001600160401b038211908210176101db57604052565b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c80630396cb60146100e7578063205c2878146100e25780632cd10339146100dd57806352b7512c146100d8578063715018a6146100d35780637c627b21146100ce5780638da5cb5b146100c9578063b0d691fe146100c4578063bb9fe6bf146100bf578063c23a5cea146100ba578063c399ec88146100b5578063d0e30db0146100b05763f2fde38b146100ab57600080fd5b610694565b610628565b61057d565b6104f4565b610480565b61043b565b610412565b6103b7565b61035c565b6102b8565b61022e565b610199565b600060203660031901126101805760043563ffffffff811680910361017c5761010e61082a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316908290823b1561017c57602460405180948193621cb65b60e51b8352600483015234905af180156101775761016b575080f35b61017490610720565b80f35b610756565b5080fd5b80fd5b6001600160a01b0381160361019457565b600080fd5b346101945760006040366003190112610180576004356101b881610183565b6101c061082a565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b1561022a5760449083604051958694859363040b850f60e31b855216600484015260243560248401525af180156101775761016b575080f35b8280fd5b34610194576020366003190112610194576004356002811015610194576000805460ff60a01b191660a09290921b60ff60a01b16919091179055005b91906040835280519081604085015260005b8281106102a25750506020606082600082819588010152601f8019910116850101930152565b806020809284010151606082880101520161027c565b34610194576003196060368201126101945760043567ffffffffffffffff9182821161019457610120908236030112610194576102f3610856565b6004013561030081610183565b6040805160609290921b6bffffffffffffffffffffffff19166020830152601482528101918211818310176103465781604052603f1990610341818461026a565b030190f35b634e487b7160e01b600052604160045260246000fd5b34610194576000806003193601126101805761037661082a565b80546001600160a01b03198116825581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b3461019457608036600319011261019457600360043510156101945760243567ffffffffffffffff8082116101945736602383011215610194578160040135908111610194573691016024011161019457610410610782565b005b34610194576000366003190112610194576000546040516001600160a01b039091168152602090f35b34610194576000366003190112610194576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610194576000806003193601126101805761049a61082a565b807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b156104f157819060046040518094819363bb9fe6bf60e01b83525af180156101775761016b575080f35b50fd5b3461019457600060203660031901126101805760043561051381610183565b61051b61082a565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b1561022a5760249083604051958694859363611d2e7560e11b85521660048401525af180156101775761016b575080f35b3461019457600080600319360112610180576040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9081156101775782916105f1575b604051828152602090f35b0390f35b90506020813d602011610620575b8161060c60209383610734565b8101031261017c576105ed915051386105e2565b3d91506105ff565b600080600319360112610180577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681813b156101805760405163b760faf960e01b8152306004820152918290602490829034905af180156101775761016b575080f35b34610194576020366003190112610194576004356106b181610183565b6106b961082a565b6001600160a01b03908116801561070757600080546001600160a01b03198116831782559092167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b604051631e4fbdf760e01b815260006004820152602490fd5b67ffffffffffffffff811161034657604052565b90601f8019910116810190811067ffffffffffffffff82111761034657604052565b6040513d6000823e3d90fd5b6002111561076c57565b634e487b7160e01b600052602160045260246000fd5b61078a610856565b60ff60005460a01c1661079c81610762565b80610806576040516346b7545f60e11b815260206004820152603660248201527f746869732069732061206c6f6e672072657665727420726561736f6e2073747260448201527534b733903bb29030b932903637b7b5b4b733903337b960511b6064820152608490fd5b80610812600192610762565b1461081957565b63deaddead60e01b60005260206000fd5b6000546001600160a01b0316330361083e57565b60405163118cdaa760e01b8152336004820152602490fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361088857565b60405162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08115b9d1c9e541bda5b9d605a1b6044820152606490fdfea264697066735822122032792d1a5b8dd844960178bc3ad8630181e7346d30fb2cfe3edce9c4804155da64736f6c63430008190033",
}

// TestPaymasterRevertCustomErrorABI is the input ABI used to generate the binding from.
// Deprecated: Use TestPaymasterRevertCustomErrorMetaData.ABI instead.
var TestPaymasterRevertCustomErrorABI = TestPaymasterRevertCustomErrorMetaData.ABI

// TestPaymasterRevertCustomErrorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestPaymasterRevertCustomErrorMetaData.Bin instead.
var TestPaymasterRevertCustomErrorBin = TestPaymasterRevertCustomErrorMetaData.Bin

// DeployTestPaymasterRevertCustomError deploys a new Ethereum contract, binding an instance of TestPaymasterRevertCustomError to it.
func DeployTestPaymasterRevertCustomError(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *TestPaymasterRevertCustomError, error) {
	parsed, err := TestPaymasterRevertCustomErrorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestPaymasterRevertCustomErrorBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestPaymasterRevertCustomError{TestPaymasterRevertCustomErrorCaller: TestPaymasterRevertCustomErrorCaller{contract: contract}, TestPaymasterRevertCustomErrorTransactor: TestPaymasterRevertCustomErrorTransactor{contract: contract}, TestPaymasterRevertCustomErrorFilterer: TestPaymasterRevertCustomErrorFilterer{contract: contract}}, nil
}

// TestPaymasterRevertCustomError is an auto generated Go binding around an Ethereum contract.
type TestPaymasterRevertCustomError struct {
	TestPaymasterRevertCustomErrorCaller     // Read-only binding to the contract
	TestPaymasterRevertCustomErrorTransactor // Write-only binding to the contract
	TestPaymasterRevertCustomErrorFilterer   // Log filterer for contract events
}

// TestPaymasterRevertCustomErrorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestPaymasterRevertCustomErrorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPaymasterRevertCustomErrorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestPaymasterRevertCustomErrorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPaymasterRevertCustomErrorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestPaymasterRevertCustomErrorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPaymasterRevertCustomErrorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestPaymasterRevertCustomErrorSession struct {
	Contract     *TestPaymasterRevertCustomError // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TestPaymasterRevertCustomErrorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestPaymasterRevertCustomErrorCallerSession struct {
	Contract *TestPaymasterRevertCustomErrorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// TestPaymasterRevertCustomErrorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestPaymasterRevertCustomErrorTransactorSession struct {
	Contract     *TestPaymasterRevertCustomErrorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// TestPaymasterRevertCustomErrorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestPaymasterRevertCustomErrorRaw struct {
	Contract *TestPaymasterRevertCustomError // Generic contract binding to access the raw methods on
}

// TestPaymasterRevertCustomErrorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestPaymasterRevertCustomErrorCallerRaw struct {
	Contract *TestPaymasterRevertCustomErrorCaller // Generic read-only contract binding to access the raw methods on
}

// TestPaymasterRevertCustomErrorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestPaymasterRevertCustomErrorTransactorRaw struct {
	Contract *TestPaymasterRevertCustomErrorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestPaymasterRevertCustomError creates a new instance of TestPaymasterRevertCustomError, bound to a specific deployed contract.
func NewTestPaymasterRevertCustomError(address common.Address, backend bind.ContractBackend) (*TestPaymasterRevertCustomError, error) {
	contract, err := bindTestPaymasterRevertCustomError(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestPaymasterRevertCustomError{TestPaymasterRevertCustomErrorCaller: TestPaymasterRevertCustomErrorCaller{contract: contract}, TestPaymasterRevertCustomErrorTransactor: TestPaymasterRevertCustomErrorTransactor{contract: contract}, TestPaymasterRevertCustomErrorFilterer: TestPaymasterRevertCustomErrorFilterer{contract: contract}}, nil
}

// NewTestPaymasterRevertCustomErrorCaller creates a new read-only instance of TestPaymasterRevertCustomError, bound to a specific deployed contract.
func NewTestPaymasterRevertCustomErrorCaller(address common.Address, caller bind.ContractCaller) (*TestPaymasterRevertCustomErrorCaller, error) {
	contract, err := bindTestPaymasterRevertCustomError(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestPaymasterRevertCustomErrorCaller{contract: contract}, nil
}

// NewTestPaymasterRevertCustomErrorTransactor creates a new write-only instance of TestPaymasterRevertCustomError, bound to a specific deployed contract.
func NewTestPaymasterRevertCustomErrorTransactor(address common.Address, transactor bind.ContractTransactor) (*TestPaymasterRevertCustomErrorTransactor, error) {
	contract, err := bindTestPaymasterRevertCustomError(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestPaymasterRevertCustomErrorTransactor{contract: contract}, nil
}

// NewTestPaymasterRevertCustomErrorFilterer creates a new log filterer instance of TestPaymasterRevertCustomError, bound to a specific deployed contract.
func NewTestPaymasterRevertCustomErrorFilterer(address common.Address, filterer bind.ContractFilterer) (*TestPaymasterRevertCustomErrorFilterer, error) {
	contract, err := bindTestPaymasterRevertCustomError(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestPaymasterRevertCustomErrorFilterer{contract: contract}, nil
}

// bindTestPaymasterRevertCustomError binds a generic wrapper to an already deployed contract.
func bindTestPaymasterRevertCustomError(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestPaymasterRevertCustomErrorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestPaymasterRevertCustomError.Contract.TestPaymasterRevertCustomErrorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.TestPaymasterRevertCustomErrorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.TestPaymasterRevertCustomErrorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestPaymasterRevertCustomError.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestPaymasterRevertCustomError.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorSession) EntryPoint() (common.Address, error) {
	return _TestPaymasterRevertCustomError.Contract.EntryPoint(&_TestPaymasterRevertCustomError.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorCallerSession) EntryPoint() (common.Address, error) {
	return _TestPaymasterRevertCustomError.Contract.EntryPoint(&_TestPaymasterRevertCustomError.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPaymasterRevertCustomError.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorSession) GetDeposit() (*big.Int, error) {
	return _TestPaymasterRevertCustomError.Contract.GetDeposit(&_TestPaymasterRevertCustomError.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorCallerSession) GetDeposit() (*big.Int, error) {
	return _TestPaymasterRevertCustomError.Contract.GetDeposit(&_TestPaymasterRevertCustomError.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestPaymasterRevertCustomError.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorSession) Owner() (common.Address, error) {
	return _TestPaymasterRevertCustomError.Contract.Owner(&_TestPaymasterRevertCustomError.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorCallerSession) Owner() (common.Address, error) {
	return _TestPaymasterRevertCustomError.Contract.Owner(&_TestPaymasterRevertCustomError.CallOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.AddStake(&_TestPaymasterRevertCustomError.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.AddStake(&_TestPaymasterRevertCustomError.TransactOpts, unstakeDelaySec)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorSession) Deposit() (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.Deposit(&_TestPaymasterRevertCustomError.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactorSession) Deposit() (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.Deposit(&_TestPaymasterRevertCustomError.TransactOpts)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.contract.Transact(opts, "postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.PostOp(&_TestPaymasterRevertCustomError.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.PostOp(&_TestPaymasterRevertCustomError.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.RenounceOwnership(&_TestPaymasterRevertCustomError.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.RenounceOwnership(&_TestPaymasterRevertCustomError.TransactOpts)
}

// SetRevertType is a paid mutator transaction binding the contract method 0x2cd10339.
//
// Solidity: function setRevertType(uint8 _revertType) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactor) SetRevertType(opts *bind.TransactOpts, _revertType uint8) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.contract.Transact(opts, "setRevertType", _revertType)
}

// SetRevertType is a paid mutator transaction binding the contract method 0x2cd10339.
//
// Solidity: function setRevertType(uint8 _revertType) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorSession) SetRevertType(_revertType uint8) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.SetRevertType(&_TestPaymasterRevertCustomError.TransactOpts, _revertType)
}

// SetRevertType is a paid mutator transaction binding the contract method 0x2cd10339.
//
// Solidity: function setRevertType(uint8 _revertType) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactorSession) SetRevertType(_revertType uint8) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.SetRevertType(&_TestPaymasterRevertCustomError.TransactOpts, _revertType)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.TransferOwnership(&_TestPaymasterRevertCustomError.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.TransferOwnership(&_TestPaymasterRevertCustomError.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorSession) UnlockStake() (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.UnlockStake(&_TestPaymasterRevertCustomError.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.UnlockStake(&_TestPaymasterRevertCustomError.TransactOpts)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.ValidatePaymasterUserOp(&_TestPaymasterRevertCustomError.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactorSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.ValidatePaymasterUserOp(&_TestPaymasterRevertCustomError.TransactOpts, userOp, userOpHash, maxCost)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.WithdrawStake(&_TestPaymasterRevertCustomError.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.WithdrawStake(&_TestPaymasterRevertCustomError.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.WithdrawTo(&_TestPaymasterRevertCustomError.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestPaymasterRevertCustomError.Contract.WithdrawTo(&_TestPaymasterRevertCustomError.TransactOpts, withdrawAddress, amount)
}

// TestPaymasterRevertCustomErrorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TestPaymasterRevertCustomError contract.
type TestPaymasterRevertCustomErrorOwnershipTransferredIterator struct {
	Event *TestPaymasterRevertCustomErrorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TestPaymasterRevertCustomErrorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPaymasterRevertCustomErrorOwnershipTransferred)
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
		it.Event = new(TestPaymasterRevertCustomErrorOwnershipTransferred)
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
func (it *TestPaymasterRevertCustomErrorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPaymasterRevertCustomErrorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPaymasterRevertCustomErrorOwnershipTransferred represents a OwnershipTransferred event raised by the TestPaymasterRevertCustomError contract.
type TestPaymasterRevertCustomErrorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestPaymasterRevertCustomErrorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestPaymasterRevertCustomError.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestPaymasterRevertCustomErrorOwnershipTransferredIterator{contract: _TestPaymasterRevertCustomError.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestPaymasterRevertCustomErrorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestPaymasterRevertCustomError.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPaymasterRevertCustomErrorOwnershipTransferred)
				if err := _TestPaymasterRevertCustomError.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TestPaymasterRevertCustomError *TestPaymasterRevertCustomErrorFilterer) ParseOwnershipTransferred(log types.Log) (*TestPaymasterRevertCustomErrorOwnershipTransferred, error) {
	event := new(TestPaymasterRevertCustomErrorOwnershipTransferred)
	if err := _TestPaymasterRevertCustomError.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
