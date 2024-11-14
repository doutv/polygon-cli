// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testsignatureaggregator

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

// TestSignatureAggregatorMetaData contains all meta data concerning the TestSignatureAggregator contract.
var TestSignatureAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"delay\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"}],\"name\":\"aggregateSignatures\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"aggregatedSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validateSignatures\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"validateUserOpSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608080604052346015576104af908161001b8239f35b600080fdfe60806040908082526004918236101561001757600080fd5b600091823560e01c908163062a422b1461032d575080632dd81133146101ab578063451711591461010f5763ae574a431461005157600080fd5b3461010b5760209260206003193601126101075767ffffffffffffffff9381358581116101035761008590369084016103e7565b9185925b8084106100db575050505081519260208401526020835281830193838510908511176100c65750829052603f19906100c1818461039e565b030190f35b604190634e487b7160e01b6000525260246000fd5b909192956100f9600191846100f18a868961041d565b013590610456565b9601929190610089565b8480fd5b8280fd5b5080fd5b508083916003193601126101075781356001600160a01b03811692908390036101a7578360243563ffffffff811680910361010b57843b1561010b576024845180968193621cb65b60e51b83528683015234905af1801561019d57610172578380f35b67ffffffffffffffff831161018a5750528180808380f35b634e487b7160e01b845260419052602483fd5b82513d86823e3d90fd5b8380fd5b5082903461010757806003193601126101075767ffffffffffffffff8235818111610103576101dd90369085016103e7565b9190602435918083116103295736602384011215610329578286013590811161032957808301913660248401116103255793879488915b808310610306575050506020036102af5781602091031261010357602401350361023c578280f35b90602060a492519162461bcd60e51b8352820152604160248201527f546573745369676e617475726556616c696461746f723a20616767726567617460448201527f6564207369676e6174757265206d69736d6174636820286e6f6e63652073756d6064820152602960f81b6084820152fd5b835162461bcd60e51b8152602081870152602b60248201527f546573745369676e617475726556616c696461746f723a20736967206d75737460448201526a103132903ab4b73a191a9b60a91b6064820152608490fd5b90919561031c60019160206100f18a868861041d565b96019190610214565b8780fd5b8680fd5b9290503461039b5760031960203682011261010b5767ffffffffffffffff9085358281116101a75790610120913603011261010b57602084019081118482101761038857610384945082528252519182918261039e565b0390f35b634e487b7160e01b825260418552602482fd5b80fd5b6020808252825181830181905290939260005b8281106103d357505060409293506000838284010152601f8019910116010190565b8181018601518482016040015285016103b1565b9181601f840112156104185782359167ffffffffffffffff8311610418576020808501948460051b01011161041857565b600080fd5b91908110156104405760051b8101359061011e1981360301821215610418570190565b634e487b7160e01b600052603260045260246000fd5b9190820180921161046357565b634e487b7160e01b600052601160045260246000fdfea26469706673582212209be58a2299f3de4b3d983f30e342532e93d4dadca530d837baf5b69a3d58439d64736f6c63430008190033",
}

// TestSignatureAggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use TestSignatureAggregatorMetaData.ABI instead.
var TestSignatureAggregatorABI = TestSignatureAggregatorMetaData.ABI

// TestSignatureAggregatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestSignatureAggregatorMetaData.Bin instead.
var TestSignatureAggregatorBin = TestSignatureAggregatorMetaData.Bin

// DeployTestSignatureAggregator deploys a new Ethereum contract, binding an instance of TestSignatureAggregator to it.
func DeployTestSignatureAggregator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestSignatureAggregator, error) {
	parsed, err := TestSignatureAggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestSignatureAggregatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestSignatureAggregator{TestSignatureAggregatorCaller: TestSignatureAggregatorCaller{contract: contract}, TestSignatureAggregatorTransactor: TestSignatureAggregatorTransactor{contract: contract}, TestSignatureAggregatorFilterer: TestSignatureAggregatorFilterer{contract: contract}}, nil
}

// TestSignatureAggregator is an auto generated Go binding around an Ethereum contract.
type TestSignatureAggregator struct {
	TestSignatureAggregatorCaller     // Read-only binding to the contract
	TestSignatureAggregatorTransactor // Write-only binding to the contract
	TestSignatureAggregatorFilterer   // Log filterer for contract events
}

// TestSignatureAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestSignatureAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSignatureAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestSignatureAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSignatureAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestSignatureAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSignatureAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSignatureAggregatorSession struct {
	Contract     *TestSignatureAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestSignatureAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestSignatureAggregatorCallerSession struct {
	Contract *TestSignatureAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// TestSignatureAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestSignatureAggregatorTransactorSession struct {
	Contract     *TestSignatureAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// TestSignatureAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestSignatureAggregatorRaw struct {
	Contract *TestSignatureAggregator // Generic contract binding to access the raw methods on
}

// TestSignatureAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestSignatureAggregatorCallerRaw struct {
	Contract *TestSignatureAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// TestSignatureAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestSignatureAggregatorTransactorRaw struct {
	Contract *TestSignatureAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestSignatureAggregator creates a new instance of TestSignatureAggregator, bound to a specific deployed contract.
func NewTestSignatureAggregator(address common.Address, backend bind.ContractBackend) (*TestSignatureAggregator, error) {
	contract, err := bindTestSignatureAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestSignatureAggregator{TestSignatureAggregatorCaller: TestSignatureAggregatorCaller{contract: contract}, TestSignatureAggregatorTransactor: TestSignatureAggregatorTransactor{contract: contract}, TestSignatureAggregatorFilterer: TestSignatureAggregatorFilterer{contract: contract}}, nil
}

// NewTestSignatureAggregatorCaller creates a new read-only instance of TestSignatureAggregator, bound to a specific deployed contract.
func NewTestSignatureAggregatorCaller(address common.Address, caller bind.ContractCaller) (*TestSignatureAggregatorCaller, error) {
	contract, err := bindTestSignatureAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestSignatureAggregatorCaller{contract: contract}, nil
}

// NewTestSignatureAggregatorTransactor creates a new write-only instance of TestSignatureAggregator, bound to a specific deployed contract.
func NewTestSignatureAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*TestSignatureAggregatorTransactor, error) {
	contract, err := bindTestSignatureAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestSignatureAggregatorTransactor{contract: contract}, nil
}

// NewTestSignatureAggregatorFilterer creates a new log filterer instance of TestSignatureAggregator, bound to a specific deployed contract.
func NewTestSignatureAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*TestSignatureAggregatorFilterer, error) {
	contract, err := bindTestSignatureAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestSignatureAggregatorFilterer{contract: contract}, nil
}

// bindTestSignatureAggregator binds a generic wrapper to an already deployed contract.
func bindTestSignatureAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestSignatureAggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestSignatureAggregator *TestSignatureAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestSignatureAggregator.Contract.TestSignatureAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestSignatureAggregator *TestSignatureAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSignatureAggregator.Contract.TestSignatureAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestSignatureAggregator *TestSignatureAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestSignatureAggregator.Contract.TestSignatureAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestSignatureAggregator *TestSignatureAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestSignatureAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestSignatureAggregator *TestSignatureAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSignatureAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestSignatureAggregator *TestSignatureAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestSignatureAggregator.Contract.contract.Transact(opts, method, params...)
}

// AggregateSignatures is a free data retrieval call binding the contract method 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) pure returns(bytes aggregatedSignature)
func (_TestSignatureAggregator *TestSignatureAggregatorCaller) AggregateSignatures(opts *bind.CallOpts, userOps []PackedUserOperation) ([]byte, error) {
	var out []interface{}
	err := _TestSignatureAggregator.contract.Call(opts, &out, "aggregateSignatures", userOps)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AggregateSignatures is a free data retrieval call binding the contract method 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) pure returns(bytes aggregatedSignature)
func (_TestSignatureAggregator *TestSignatureAggregatorSession) AggregateSignatures(userOps []PackedUserOperation) ([]byte, error) {
	return _TestSignatureAggregator.Contract.AggregateSignatures(&_TestSignatureAggregator.CallOpts, userOps)
}

// AggregateSignatures is a free data retrieval call binding the contract method 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) pure returns(bytes aggregatedSignature)
func (_TestSignatureAggregator *TestSignatureAggregatorCallerSession) AggregateSignatures(userOps []PackedUserOperation) ([]byte, error) {
	return _TestSignatureAggregator.Contract.AggregateSignatures(&_TestSignatureAggregator.CallOpts, userOps)
}

// ValidateSignatures is a free data retrieval call binding the contract method 0x2dd81133.
//
// Solidity: function validateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, bytes signature) pure returns()
func (_TestSignatureAggregator *TestSignatureAggregatorCaller) ValidateSignatures(opts *bind.CallOpts, userOps []PackedUserOperation, signature []byte) error {
	var out []interface{}
	err := _TestSignatureAggregator.contract.Call(opts, &out, "validateSignatures", userOps, signature)

	if err != nil {
		return err
	}

	return err

}

// ValidateSignatures is a free data retrieval call binding the contract method 0x2dd81133.
//
// Solidity: function validateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, bytes signature) pure returns()
func (_TestSignatureAggregator *TestSignatureAggregatorSession) ValidateSignatures(userOps []PackedUserOperation, signature []byte) error {
	return _TestSignatureAggregator.Contract.ValidateSignatures(&_TestSignatureAggregator.CallOpts, userOps, signature)
}

// ValidateSignatures is a free data retrieval call binding the contract method 0x2dd81133.
//
// Solidity: function validateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, bytes signature) pure returns()
func (_TestSignatureAggregator *TestSignatureAggregatorCallerSession) ValidateSignatures(userOps []PackedUserOperation, signature []byte) error {
	return _TestSignatureAggregator.Contract.ValidateSignatures(&_TestSignatureAggregator.CallOpts, userOps, signature)
}

// ValidateUserOpSignature is a free data retrieval call binding the contract method 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) ) pure returns(bytes)
func (_TestSignatureAggregator *TestSignatureAggregatorCaller) ValidateUserOpSignature(opts *bind.CallOpts, arg0 PackedUserOperation) ([]byte, error) {
	var out []interface{}
	err := _TestSignatureAggregator.contract.Call(opts, &out, "validateUserOpSignature", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ValidateUserOpSignature is a free data retrieval call binding the contract method 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) ) pure returns(bytes)
func (_TestSignatureAggregator *TestSignatureAggregatorSession) ValidateUserOpSignature(arg0 PackedUserOperation) ([]byte, error) {
	return _TestSignatureAggregator.Contract.ValidateUserOpSignature(&_TestSignatureAggregator.CallOpts, arg0)
}

// ValidateUserOpSignature is a free data retrieval call binding the contract method 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) ) pure returns(bytes)
func (_TestSignatureAggregator *TestSignatureAggregatorCallerSession) ValidateUserOpSignature(arg0 PackedUserOperation) ([]byte, error) {
	return _TestSignatureAggregator.Contract.ValidateUserOpSignature(&_TestSignatureAggregator.CallOpts, arg0)
}

// AddStake is a paid mutator transaction binding the contract method 0x45171159.
//
// Solidity: function addStake(address entryPoint, uint32 delay) payable returns()
func (_TestSignatureAggregator *TestSignatureAggregatorTransactor) AddStake(opts *bind.TransactOpts, entryPoint common.Address, delay uint32) (*types.Transaction, error) {
	return _TestSignatureAggregator.contract.Transact(opts, "addStake", entryPoint, delay)
}

// AddStake is a paid mutator transaction binding the contract method 0x45171159.
//
// Solidity: function addStake(address entryPoint, uint32 delay) payable returns()
func (_TestSignatureAggregator *TestSignatureAggregatorSession) AddStake(entryPoint common.Address, delay uint32) (*types.Transaction, error) {
	return _TestSignatureAggregator.Contract.AddStake(&_TestSignatureAggregator.TransactOpts, entryPoint, delay)
}

// AddStake is a paid mutator transaction binding the contract method 0x45171159.
//
// Solidity: function addStake(address entryPoint, uint32 delay) payable returns()
func (_TestSignatureAggregator *TestSignatureAggregatorTransactorSession) AddStake(entryPoint common.Address, delay uint32) (*types.Transaction, error) {
	return _TestSignatureAggregator.Contract.AddStake(&_TestSignatureAggregator.TransactOpts, entryPoint, delay)
}
