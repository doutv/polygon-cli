// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package maliciousaccount

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

// MaliciousAccountMetaData contains all meta data concerning the MaliciousAccount contract.
var MaliciousAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_ep\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080601f6102c638819003918201601f19168301916001600160401b03831184841017607157808492602094604052833981010312606c57516001600160a01b03811690819003606c57600080546001600160a01b03191691909117905560405161023e90816100888239f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe60808060405260048036101561001457600080fd5b600091823560e01c6319822f7c1461002b57600080fd5b346101e157600319906060368301126101dd5782359067ffffffffffffffff9081831161019d576101208336039485011261019d578554604435946001600160a01b039091169187833b156101da5781879463b760faf960e01b8252308a830152816024998a925af180156101cf576101a9575b5061010484013590602219018112156101a557830191858301359081116101a5578483019281360384136101a1576020918101031261019d5761010a6100ff60848501356fffffffffffffffffffffffffffffffff81169060801c6101e5565b60a4850135906101e5565b90811561018b579060c4910492013560801c820391821161017957350361013657602083604051908152f35b90601d60649260206040519362461bcd60e51b85528401528201527f5265766572742061667465722066697273742076616c69646174696f6e0000006044820152fd5b634e487b7160e01b8552601184528285fd5b634e487b7160e01b8752601286528487fd5b8580fd5b8780fd5b8680fd5b8381989298116101bd57604052953861009f565b634e487b7160e01b8252604187528582fd5b6040513d8a823e3d90fd5b80fd5b8380fd5b8280fd5b919082018092116101f257565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220075c734a9985998ab5fb2f22874da9680c10e816124510961d7e61e8be70778664736f6c63430008190033",
}

// MaliciousAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use MaliciousAccountMetaData.ABI instead.
var MaliciousAccountABI = MaliciousAccountMetaData.ABI

// MaliciousAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MaliciousAccountMetaData.Bin instead.
var MaliciousAccountBin = MaliciousAccountMetaData.Bin

// DeployMaliciousAccount deploys a new Ethereum contract, binding an instance of MaliciousAccount to it.
func DeployMaliciousAccount(auth *bind.TransactOpts, backend bind.ContractBackend, _ep common.Address) (common.Address, *types.Transaction, *MaliciousAccount, error) {
	parsed, err := MaliciousAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MaliciousAccountBin), backend, _ep)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MaliciousAccount{MaliciousAccountCaller: MaliciousAccountCaller{contract: contract}, MaliciousAccountTransactor: MaliciousAccountTransactor{contract: contract}, MaliciousAccountFilterer: MaliciousAccountFilterer{contract: contract}}, nil
}

// MaliciousAccount is an auto generated Go binding around an Ethereum contract.
type MaliciousAccount struct {
	MaliciousAccountCaller     // Read-only binding to the contract
	MaliciousAccountTransactor // Write-only binding to the contract
	MaliciousAccountFilterer   // Log filterer for contract events
}

// MaliciousAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type MaliciousAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaliciousAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MaliciousAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaliciousAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MaliciousAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaliciousAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MaliciousAccountSession struct {
	Contract     *MaliciousAccount // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MaliciousAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MaliciousAccountCallerSession struct {
	Contract *MaliciousAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MaliciousAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MaliciousAccountTransactorSession struct {
	Contract     *MaliciousAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MaliciousAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type MaliciousAccountRaw struct {
	Contract *MaliciousAccount // Generic contract binding to access the raw methods on
}

// MaliciousAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MaliciousAccountCallerRaw struct {
	Contract *MaliciousAccountCaller // Generic read-only contract binding to access the raw methods on
}

// MaliciousAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MaliciousAccountTransactorRaw struct {
	Contract *MaliciousAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMaliciousAccount creates a new instance of MaliciousAccount, bound to a specific deployed contract.
func NewMaliciousAccount(address common.Address, backend bind.ContractBackend) (*MaliciousAccount, error) {
	contract, err := bindMaliciousAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MaliciousAccount{MaliciousAccountCaller: MaliciousAccountCaller{contract: contract}, MaliciousAccountTransactor: MaliciousAccountTransactor{contract: contract}, MaliciousAccountFilterer: MaliciousAccountFilterer{contract: contract}}, nil
}

// NewMaliciousAccountCaller creates a new read-only instance of MaliciousAccount, bound to a specific deployed contract.
func NewMaliciousAccountCaller(address common.Address, caller bind.ContractCaller) (*MaliciousAccountCaller, error) {
	contract, err := bindMaliciousAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MaliciousAccountCaller{contract: contract}, nil
}

// NewMaliciousAccountTransactor creates a new write-only instance of MaliciousAccount, bound to a specific deployed contract.
func NewMaliciousAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*MaliciousAccountTransactor, error) {
	contract, err := bindMaliciousAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MaliciousAccountTransactor{contract: contract}, nil
}

// NewMaliciousAccountFilterer creates a new log filterer instance of MaliciousAccount, bound to a specific deployed contract.
func NewMaliciousAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*MaliciousAccountFilterer, error) {
	contract, err := bindMaliciousAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MaliciousAccountFilterer{contract: contract}, nil
}

// bindMaliciousAccount binds a generic wrapper to an already deployed contract.
func bindMaliciousAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MaliciousAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MaliciousAccount *MaliciousAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MaliciousAccount.Contract.MaliciousAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MaliciousAccount *MaliciousAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaliciousAccount.Contract.MaliciousAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MaliciousAccount *MaliciousAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MaliciousAccount.Contract.MaliciousAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MaliciousAccount *MaliciousAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MaliciousAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MaliciousAccount *MaliciousAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaliciousAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MaliciousAccount *MaliciousAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MaliciousAccount.Contract.contract.Transact(opts, method, params...)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 , uint256 missingAccountFunds) returns(uint256 validationData)
func (_MaliciousAccount *MaliciousAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, arg1 [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _MaliciousAccount.contract.Transact(opts, "validateUserOp", userOp, arg1, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 , uint256 missingAccountFunds) returns(uint256 validationData)
func (_MaliciousAccount *MaliciousAccountSession) ValidateUserOp(userOp PackedUserOperation, arg1 [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _MaliciousAccount.Contract.ValidateUserOp(&_MaliciousAccount.TransactOpts, userOp, arg1, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 , uint256 missingAccountFunds) returns(uint256 validationData)
func (_MaliciousAccount *MaliciousAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, arg1 [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _MaliciousAccount.Contract.ValidateUserOp(&_MaliciousAccount.TransactOpts, userOp, arg1, missingAccountFunds)
}
