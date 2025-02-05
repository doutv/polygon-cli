// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testexpirepaymaster

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

// TestExpirePaymasterMetaData contains all meta data concerning the TestExpirePaymaster contract.
var TestExpirePaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604090808252346101b35780610a6a803803809161001f82856101b8565b83396020928391810103126101b35751906001600160a01b03808316908184036101b357331561019b5760008054336001600160a01b031982168117835587519492938693869360249385939091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08880a36301ffc9a760e01b825263122a0e9b60e31b60048301525afa91821561018f57819261014e575b50501561010b57506080525161087890816101f28239608051818181610154015281816101ff015281816102e101528181610357015281816103c201528181610664015281816106e901526107c20152f35b60649083519062461bcd60e51b82526004820152601e60248201527f49456e747279506f696e7420696e74657266616365206d69736d6174636800006044820152fd5b9091508281813d8311610188575b61016681836101b8565b810103126101845751908115158203610181575038806100b9565b80fd5b5080fd5b503d61015c565b508451903d90823e3d90fd5b8451631e4fbdf760e01b815260006004820152602490fd5b600080fd5b601f909101601f19168101906001600160401b038211908210176101db57604052565b634e487b7160e01b600052604160045260246000fdfe60406080815260048036101561001457600080fd5b600091823560e01c80630396cb60146106c157838163205c2878146106375750806352b7512c14610503578063715018a6146104a95780637c627b211461041d5780638da5cb5b146103f5578063b0d691fe146103ad57838163bb9fe6bf1461033a578163c23a5cea146102b357508063c399ec88146101d557838163d0e30db014610144575063f2fde38b146100aa57600080fd5b34610140576020366003190112610140576001600160a01b0382358181169391929084900361013c576100db610794565b8315610126575050600054826bffffffffffffffffffffffff60a01b821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a380f35b51631e4fbdf760e01b8152908101849052602490fd5b8480fd5b8280fd5b808484826003193601126101d1577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156101cc578390602483518095819363b760faf960e01b8352309083015234905af19081156101c357506101b05750f35b6101b99061076a565b6101c05780f35b80fd5b513d84823e3d90fd5b505050fd5b5050fd5b503461014057826003193601126101405780516370a0823160e01b815230838201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9283156102a9578493610241575b6020848451908152f35b909192506020903d6020116102a1575b601f8201601f191683019067ffffffffffffffff82118483101761028c57506020918391855281010312610140576020925051903880610237565b604190634e487b7160e01b6000525260246000fd5b3d9150610251565b82513d86823e3d90fd5b808484346101d15760203660031901126101d1576102cf61074f565b6102d7610794565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116803b15610336578592836024928651978895869463611d2e7560e11b865216908401525af19081156101c357506101b05750f35b8580fd5b808484346101d157826003193601126101d157610355610794565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156101cc57815163bb9fe6bf60e01b81529284918491829084905af19081156101c357506101b05750f35b8382346103f157816003193601126103f157517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b8382346103f157816003193601126103f157905490516001600160a01b039091168152602090f35b5091346101c05760803660031901126101c0576003823510156101c05760243567ffffffffffffffff8082116101405736602383011215610140578184013590811161014057369101602401116101c05750602060649261047c6107c0565b5162461bcd60e51b815291820152600d60248201526c6d757374206f7665727269646560981b6044820152fd5b83346101c057806003193601126101c0576104c2610794565b600080546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b50903461014057606092600319906060823601126101c05767ffffffffffffffff928035848111610140578036039061012085830112610633576105456107c0565b60e48101359160221901821215610633570192818401359085821161063357813603602486011361063357816034116106335790840184900301602f190185136103f1576105956058840161082f565b9265ffffffffffff60a01b906105ad9060780161082f565b60a01b169385519060209260208301918383109083111761028c5750858589939589938452808752835196879585875281518096880152825b86811061061d575050506060809650848601015265ffffffffffff60d01b9060d01b16176020830152601f80199101168101030190f35b8083018901518a820183015289975088016105e6565b8380fd5b808484346101d157806003193601126101d15761065261074f565b61065a610794565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116803b15610336578592836044928651978895869463040b850f60e31b8652169084015260243560248401525af19081156101c357506101b05750f35b5060203660031901126101405782823563ffffffff81168091036103f1576106e7610794565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031693843b156101405760249084519586938492621cb65b60e51b845283015234905af19081156101c35750610743575080f35b61074c9061076a565b80f35b600435906001600160a01b038216820361076557565b600080fd5b67ffffffffffffffff811161077e57604052565b634e487b7160e01b600052604160045260246000fd5b6000546001600160a01b031633036107a857565b60405163118cdaa760e01b8152336004820152602490fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036107f257565b60405162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08115b9d1c9e541bda5b9d605a1b6044820152606490fd5b359065ffffffffffff821682036107655756fea264697066735822122032d5afa963e4a28bc6fa342d1677d3157d07d9b1fdd76bde31279f29f776e74264736f6c63430008190033",
}

// TestExpirePaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use TestExpirePaymasterMetaData.ABI instead.
var TestExpirePaymasterABI = TestExpirePaymasterMetaData.ABI

// TestExpirePaymasterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestExpirePaymasterMetaData.Bin instead.
var TestExpirePaymasterBin = TestExpirePaymasterMetaData.Bin

// DeployTestExpirePaymaster deploys a new Ethereum contract, binding an instance of TestExpirePaymaster to it.
func DeployTestExpirePaymaster(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *TestExpirePaymaster, error) {
	parsed, err := TestExpirePaymasterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestExpirePaymasterBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestExpirePaymaster{TestExpirePaymasterCaller: TestExpirePaymasterCaller{contract: contract}, TestExpirePaymasterTransactor: TestExpirePaymasterTransactor{contract: contract}, TestExpirePaymasterFilterer: TestExpirePaymasterFilterer{contract: contract}}, nil
}

// TestExpirePaymaster is an auto generated Go binding around an Ethereum contract.
type TestExpirePaymaster struct {
	TestExpirePaymasterCaller     // Read-only binding to the contract
	TestExpirePaymasterTransactor // Write-only binding to the contract
	TestExpirePaymasterFilterer   // Log filterer for contract events
}

// TestExpirePaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestExpirePaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExpirePaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestExpirePaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExpirePaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestExpirePaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExpirePaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestExpirePaymasterSession struct {
	Contract     *TestExpirePaymaster // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestExpirePaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestExpirePaymasterCallerSession struct {
	Contract *TestExpirePaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TestExpirePaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestExpirePaymasterTransactorSession struct {
	Contract     *TestExpirePaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TestExpirePaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestExpirePaymasterRaw struct {
	Contract *TestExpirePaymaster // Generic contract binding to access the raw methods on
}

// TestExpirePaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestExpirePaymasterCallerRaw struct {
	Contract *TestExpirePaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// TestExpirePaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestExpirePaymasterTransactorRaw struct {
	Contract *TestExpirePaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestExpirePaymaster creates a new instance of TestExpirePaymaster, bound to a specific deployed contract.
func NewTestExpirePaymaster(address common.Address, backend bind.ContractBackend) (*TestExpirePaymaster, error) {
	contract, err := bindTestExpirePaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestExpirePaymaster{TestExpirePaymasterCaller: TestExpirePaymasterCaller{contract: contract}, TestExpirePaymasterTransactor: TestExpirePaymasterTransactor{contract: contract}, TestExpirePaymasterFilterer: TestExpirePaymasterFilterer{contract: contract}}, nil
}

// NewTestExpirePaymasterCaller creates a new read-only instance of TestExpirePaymaster, bound to a specific deployed contract.
func NewTestExpirePaymasterCaller(address common.Address, caller bind.ContractCaller) (*TestExpirePaymasterCaller, error) {
	contract, err := bindTestExpirePaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestExpirePaymasterCaller{contract: contract}, nil
}

// NewTestExpirePaymasterTransactor creates a new write-only instance of TestExpirePaymaster, bound to a specific deployed contract.
func NewTestExpirePaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*TestExpirePaymasterTransactor, error) {
	contract, err := bindTestExpirePaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestExpirePaymasterTransactor{contract: contract}, nil
}

// NewTestExpirePaymasterFilterer creates a new log filterer instance of TestExpirePaymaster, bound to a specific deployed contract.
func NewTestExpirePaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*TestExpirePaymasterFilterer, error) {
	contract, err := bindTestExpirePaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestExpirePaymasterFilterer{contract: contract}, nil
}

// bindTestExpirePaymaster binds a generic wrapper to an already deployed contract.
func bindTestExpirePaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestExpirePaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestExpirePaymaster *TestExpirePaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestExpirePaymaster.Contract.TestExpirePaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestExpirePaymaster *TestExpirePaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.TestExpirePaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestExpirePaymaster *TestExpirePaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.TestExpirePaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestExpirePaymaster *TestExpirePaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestExpirePaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestExpirePaymaster *TestExpirePaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestExpirePaymaster *TestExpirePaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestExpirePaymaster *TestExpirePaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestExpirePaymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestExpirePaymaster *TestExpirePaymasterSession) EntryPoint() (common.Address, error) {
	return _TestExpirePaymaster.Contract.EntryPoint(&_TestExpirePaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestExpirePaymaster *TestExpirePaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _TestExpirePaymaster.Contract.EntryPoint(&_TestExpirePaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestExpirePaymaster *TestExpirePaymasterCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestExpirePaymaster.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestExpirePaymaster *TestExpirePaymasterSession) GetDeposit() (*big.Int, error) {
	return _TestExpirePaymaster.Contract.GetDeposit(&_TestExpirePaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestExpirePaymaster *TestExpirePaymasterCallerSession) GetDeposit() (*big.Int, error) {
	return _TestExpirePaymaster.Contract.GetDeposit(&_TestExpirePaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestExpirePaymaster *TestExpirePaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestExpirePaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestExpirePaymaster *TestExpirePaymasterSession) Owner() (common.Address, error) {
	return _TestExpirePaymaster.Contract.Owner(&_TestExpirePaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestExpirePaymaster *TestExpirePaymasterCallerSession) Owner() (common.Address, error) {
	return _TestExpirePaymaster.Contract.Owner(&_TestExpirePaymaster.CallOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _TestExpirePaymaster.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_TestExpirePaymaster *TestExpirePaymasterSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.AddStake(&_TestExpirePaymaster.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.AddStake(&_TestExpirePaymaster.TransactOpts, unstakeDelaySec)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExpirePaymaster.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestExpirePaymaster *TestExpirePaymasterSession) Deposit() (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.Deposit(&_TestExpirePaymaster.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactorSession) Deposit() (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.Deposit(&_TestExpirePaymaster.TransactOpts)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _TestExpirePaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_TestExpirePaymaster *TestExpirePaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.PostOp(&_TestExpirePaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.PostOp(&_TestExpirePaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExpirePaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestExpirePaymaster *TestExpirePaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.RenounceOwnership(&_TestExpirePaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.RenounceOwnership(&_TestExpirePaymaster.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TestExpirePaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestExpirePaymaster *TestExpirePaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.TransferOwnership(&_TestExpirePaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.TransferOwnership(&_TestExpirePaymaster.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExpirePaymaster.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_TestExpirePaymaster *TestExpirePaymasterSession) UnlockStake() (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.UnlockStake(&_TestExpirePaymaster.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.UnlockStake(&_TestExpirePaymaster.TransactOpts)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_TestExpirePaymaster *TestExpirePaymasterTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _TestExpirePaymaster.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_TestExpirePaymaster *TestExpirePaymasterSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.ValidatePaymasterUserOp(&_TestExpirePaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_TestExpirePaymaster *TestExpirePaymasterTransactorSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.ValidatePaymasterUserOp(&_TestExpirePaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _TestExpirePaymaster.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_TestExpirePaymaster *TestExpirePaymasterSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.WithdrawStake(&_TestExpirePaymaster.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.WithdrawStake(&_TestExpirePaymaster.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestExpirePaymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_TestExpirePaymaster *TestExpirePaymasterSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.WithdrawTo(&_TestExpirePaymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_TestExpirePaymaster *TestExpirePaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestExpirePaymaster.Contract.WithdrawTo(&_TestExpirePaymaster.TransactOpts, withdrawAddress, amount)
}

// TestExpirePaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TestExpirePaymaster contract.
type TestExpirePaymasterOwnershipTransferredIterator struct {
	Event *TestExpirePaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TestExpirePaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestExpirePaymasterOwnershipTransferred)
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
		it.Event = new(TestExpirePaymasterOwnershipTransferred)
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
func (it *TestExpirePaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestExpirePaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestExpirePaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the TestExpirePaymaster contract.
type TestExpirePaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestExpirePaymaster *TestExpirePaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestExpirePaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestExpirePaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestExpirePaymasterOwnershipTransferredIterator{contract: _TestExpirePaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestExpirePaymaster *TestExpirePaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestExpirePaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestExpirePaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestExpirePaymasterOwnershipTransferred)
				if err := _TestExpirePaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TestExpirePaymaster *TestExpirePaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*TestExpirePaymasterOwnershipTransferred, error) {
	event := new(TestExpirePaymasterOwnershipTransferred)
	if err := _TestExpirePaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
