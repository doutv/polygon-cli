// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testpaymasterwithpostop

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

// TestPaymasterWithPostOpMetaData contains all meta data concerning the TestPaymasterWithPostOp contract.
var TestPaymasterWithPostOpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040908082523461018d57806109cd803803809161001f8285610192565b833960209283918101031261018d5751906001600160a01b03821680830361018d5733156101755781602491610054336101cb565b85516301ffc9a760e01b815263122a0e9b60e31b600482015292839182905afa90811561016a5760009161012d575b50156100ea57506080523332036100dc575b516107ba90816102138239608051818181610154015281816101ff015281816102e101528181610357015281816103c2015281816105bc0152818161064101526107170152f35b6100e5326101cb565b610095565b60649083519062461bcd60e51b82526004820152601e60248201527f49456e747279506f696e7420696e74657266616365206d69736d6174636800006044820152fd5b8281813d8311610163575b6101428183610192565b8101031261015f575190811515820361015c575038610083565b80fd5b5080fd5b503d610138565b84513d6000823e3d90fd5b8351631e4fbdf760e01b815260006004820152602490fd5b600080fd5b601f909101601f19168101906001600160401b038211908210176101b557604052565b634e487b7160e01b600052604160045260246000fd5b600080546001600160a01b039283166001600160a01b03198216811783559216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a356fe60406080815260048036101561001457600080fd5b600091823560e01c80630396cb601461061957838163205c28781461058f5750806352b7512c146104d8578063715018a61461047e5780637c627b211461041d5780638da5cb5b146103f5578063b0d691fe146103ad57838163bb9fe6bf1461033a578163c23a5cea146102b357508063c399ec88146101d557838163d0e30db014610144575063f2fde38b146100aa57600080fd5b34610140576020366003190112610140576001600160a01b0382358181169391929084900361013c576100db6106e9565b8315610126575050600054826bffffffffffffffffffffffff60a01b821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a380f35b51631e4fbdf760e01b8152908101849052602490fd5b8480fd5b8280fd5b808484826003193601126101d1577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156101cc578390602483518095819363b760faf960e01b8352309083015234905af19081156101c357506101b05750f35b6101b9906106bf565b6101c05780f35b80fd5b513d84823e3d90fd5b505050fd5b5050fd5b503461014057826003193601126101405780516370a0823160e01b815230838201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9283156102a9578493610241575b6020848451908152f35b909192506020903d6020116102a1575b601f8201601f191683019067ffffffffffffffff82118483101761028c57506020918391855281010312610140576020925051903880610237565b604190634e487b7160e01b6000525260246000fd5b3d9150610251565b82513d86823e3d90fd5b808484346101d15760203660031901126101d1576102cf6106a4565b6102d76106e9565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116803b15610336578592836024928651978895869463611d2e7560e11b865216908401525af19081156101c357506101b05750f35b8580fd5b808484346101d157826003193601126101d1576103556106e9565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156101cc57815163bb9fe6bf60e01b81529284918491829084905af19081156101c357506101b05750f35b8382346103f157816003193601126103f157517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b8382346103f157816003193601126103f157905490516001600160a01b039091168152602090f35b5050346103f15760803660031901126103f1576003813510156103f1576024359067ffffffffffffffff9081831161047a573660238401121561047a5782013590811161014057369101602401116101c057610477610715565b80f35b8380fd5b83346101c057806003193601126101c0576104976106e9565b600080546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b5082346101c0576060926003196060813601126101405767ffffffffffffffff90823582811161013c5790610120913603011261014057610517610715565b8351918285019182118383101761028c575083919492935260018452602092603160f81b6020860152815194859383855281518094860152825b8481106105795750505060609350808483850101526020830152601f80199101168101030190f35b8083018701518882018301528795508601610551565b808484346101d157806003193601126101d1576105aa6106a4565b6105b26106e9565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116803b15610336578592836044928651978895869463040b850f60e31b8652169084015260243560248401525af19081156101c357506101b05750f35b5060203660031901126101405782823563ffffffff81168091036103f15761063f6106e9565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031693843b156101405760249084519586938492621cb65b60e51b845283015234905af19081156101c3575061069b575080f35b610477906106bf565b600435906001600160a01b03821682036106ba57565b600080fd5b67ffffffffffffffff81116106d357604052565b634e487b7160e01b600052604160045260246000fd5b6000546001600160a01b031633036106fd57565b60405163118cdaa760e01b8152336004820152602490fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361074757565b60405162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08115b9d1c9e541bda5b9d605a1b6044820152606490fdfea264697066735822122080e2f34f6c1f0538d7bac6b5e71ca0bea9c1e50a827e7548b7f35a0623cea50664736f6c63430008190033",
}

// TestPaymasterWithPostOpABI is the input ABI used to generate the binding from.
// Deprecated: Use TestPaymasterWithPostOpMetaData.ABI instead.
var TestPaymasterWithPostOpABI = TestPaymasterWithPostOpMetaData.ABI

// TestPaymasterWithPostOpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestPaymasterWithPostOpMetaData.Bin instead.
var TestPaymasterWithPostOpBin = TestPaymasterWithPostOpMetaData.Bin

// DeployTestPaymasterWithPostOp deploys a new Ethereum contract, binding an instance of TestPaymasterWithPostOp to it.
func DeployTestPaymasterWithPostOp(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *TestPaymasterWithPostOp, error) {
	parsed, err := TestPaymasterWithPostOpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestPaymasterWithPostOpBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestPaymasterWithPostOp{TestPaymasterWithPostOpCaller: TestPaymasterWithPostOpCaller{contract: contract}, TestPaymasterWithPostOpTransactor: TestPaymasterWithPostOpTransactor{contract: contract}, TestPaymasterWithPostOpFilterer: TestPaymasterWithPostOpFilterer{contract: contract}}, nil
}

// TestPaymasterWithPostOp is an auto generated Go binding around an Ethereum contract.
type TestPaymasterWithPostOp struct {
	TestPaymasterWithPostOpCaller     // Read-only binding to the contract
	TestPaymasterWithPostOpTransactor // Write-only binding to the contract
	TestPaymasterWithPostOpFilterer   // Log filterer for contract events
}

// TestPaymasterWithPostOpCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestPaymasterWithPostOpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPaymasterWithPostOpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestPaymasterWithPostOpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPaymasterWithPostOpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestPaymasterWithPostOpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPaymasterWithPostOpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestPaymasterWithPostOpSession struct {
	Contract     *TestPaymasterWithPostOp // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestPaymasterWithPostOpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestPaymasterWithPostOpCallerSession struct {
	Contract *TestPaymasterWithPostOpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// TestPaymasterWithPostOpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestPaymasterWithPostOpTransactorSession struct {
	Contract     *TestPaymasterWithPostOpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// TestPaymasterWithPostOpRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestPaymasterWithPostOpRaw struct {
	Contract *TestPaymasterWithPostOp // Generic contract binding to access the raw methods on
}

// TestPaymasterWithPostOpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestPaymasterWithPostOpCallerRaw struct {
	Contract *TestPaymasterWithPostOpCaller // Generic read-only contract binding to access the raw methods on
}

// TestPaymasterWithPostOpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestPaymasterWithPostOpTransactorRaw struct {
	Contract *TestPaymasterWithPostOpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestPaymasterWithPostOp creates a new instance of TestPaymasterWithPostOp, bound to a specific deployed contract.
func NewTestPaymasterWithPostOp(address common.Address, backend bind.ContractBackend) (*TestPaymasterWithPostOp, error) {
	contract, err := bindTestPaymasterWithPostOp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestPaymasterWithPostOp{TestPaymasterWithPostOpCaller: TestPaymasterWithPostOpCaller{contract: contract}, TestPaymasterWithPostOpTransactor: TestPaymasterWithPostOpTransactor{contract: contract}, TestPaymasterWithPostOpFilterer: TestPaymasterWithPostOpFilterer{contract: contract}}, nil
}

// NewTestPaymasterWithPostOpCaller creates a new read-only instance of TestPaymasterWithPostOp, bound to a specific deployed contract.
func NewTestPaymasterWithPostOpCaller(address common.Address, caller bind.ContractCaller) (*TestPaymasterWithPostOpCaller, error) {
	contract, err := bindTestPaymasterWithPostOp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestPaymasterWithPostOpCaller{contract: contract}, nil
}

// NewTestPaymasterWithPostOpTransactor creates a new write-only instance of TestPaymasterWithPostOp, bound to a specific deployed contract.
func NewTestPaymasterWithPostOpTransactor(address common.Address, transactor bind.ContractTransactor) (*TestPaymasterWithPostOpTransactor, error) {
	contract, err := bindTestPaymasterWithPostOp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestPaymasterWithPostOpTransactor{contract: contract}, nil
}

// NewTestPaymasterWithPostOpFilterer creates a new log filterer instance of TestPaymasterWithPostOp, bound to a specific deployed contract.
func NewTestPaymasterWithPostOpFilterer(address common.Address, filterer bind.ContractFilterer) (*TestPaymasterWithPostOpFilterer, error) {
	contract, err := bindTestPaymasterWithPostOp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestPaymasterWithPostOpFilterer{contract: contract}, nil
}

// bindTestPaymasterWithPostOp binds a generic wrapper to an already deployed contract.
func bindTestPaymasterWithPostOp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestPaymasterWithPostOpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestPaymasterWithPostOp.Contract.TestPaymasterWithPostOpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.TestPaymasterWithPostOpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.TestPaymasterWithPostOpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestPaymasterWithPostOp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestPaymasterWithPostOp.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpSession) EntryPoint() (common.Address, error) {
	return _TestPaymasterWithPostOp.Contract.EntryPoint(&_TestPaymasterWithPostOp.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpCallerSession) EntryPoint() (common.Address, error) {
	return _TestPaymasterWithPostOp.Contract.EntryPoint(&_TestPaymasterWithPostOp.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPaymasterWithPostOp.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpSession) GetDeposit() (*big.Int, error) {
	return _TestPaymasterWithPostOp.Contract.GetDeposit(&_TestPaymasterWithPostOp.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpCallerSession) GetDeposit() (*big.Int, error) {
	return _TestPaymasterWithPostOp.Contract.GetDeposit(&_TestPaymasterWithPostOp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestPaymasterWithPostOp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpSession) Owner() (common.Address, error) {
	return _TestPaymasterWithPostOp.Contract.Owner(&_TestPaymasterWithPostOp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpCallerSession) Owner() (common.Address, error) {
	return _TestPaymasterWithPostOp.Contract.Owner(&_TestPaymasterWithPostOp.CallOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.AddStake(&_TestPaymasterWithPostOp.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.AddStake(&_TestPaymasterWithPostOp.TransactOpts, unstakeDelaySec)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpSession) Deposit() (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.Deposit(&_TestPaymasterWithPostOp.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactorSession) Deposit() (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.Deposit(&_TestPaymasterWithPostOp.TransactOpts)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.contract.Transact(opts, "postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.PostOp(&_TestPaymasterWithPostOp.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.PostOp(&_TestPaymasterWithPostOp.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.RenounceOwnership(&_TestPaymasterWithPostOp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.RenounceOwnership(&_TestPaymasterWithPostOp.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.TransferOwnership(&_TestPaymasterWithPostOp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.TransferOwnership(&_TestPaymasterWithPostOp.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpSession) UnlockStake() (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.UnlockStake(&_TestPaymasterWithPostOp.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.UnlockStake(&_TestPaymasterWithPostOp.TransactOpts)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.ValidatePaymasterUserOp(&_TestPaymasterWithPostOp.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactorSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.ValidatePaymasterUserOp(&_TestPaymasterWithPostOp.TransactOpts, userOp, userOpHash, maxCost)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.WithdrawStake(&_TestPaymasterWithPostOp.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.WithdrawStake(&_TestPaymasterWithPostOp.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.WithdrawTo(&_TestPaymasterWithPostOp.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestPaymasterWithPostOp.Contract.WithdrawTo(&_TestPaymasterWithPostOp.TransactOpts, withdrawAddress, amount)
}

// TestPaymasterWithPostOpOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TestPaymasterWithPostOp contract.
type TestPaymasterWithPostOpOwnershipTransferredIterator struct {
	Event *TestPaymasterWithPostOpOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TestPaymasterWithPostOpOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPaymasterWithPostOpOwnershipTransferred)
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
		it.Event = new(TestPaymasterWithPostOpOwnershipTransferred)
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
func (it *TestPaymasterWithPostOpOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPaymasterWithPostOpOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPaymasterWithPostOpOwnershipTransferred represents a OwnershipTransferred event raised by the TestPaymasterWithPostOp contract.
type TestPaymasterWithPostOpOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestPaymasterWithPostOpOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestPaymasterWithPostOp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestPaymasterWithPostOpOwnershipTransferredIterator{contract: _TestPaymasterWithPostOp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestPaymasterWithPostOpOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestPaymasterWithPostOp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPaymasterWithPostOpOwnershipTransferred)
				if err := _TestPaymasterWithPostOp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TestPaymasterWithPostOp *TestPaymasterWithPostOpFilterer) ParseOwnershipTransferred(log types.Log) (*TestPaymasterWithPostOpOwnershipTransferred, error) {
	event := new(TestPaymasterWithPostOpOwnershipTransferred)
	if err := _TestPaymasterWithPostOp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
