// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ientrypointsimulations

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

// IEntryPointAggregatorStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type IEntryPointAggregatorStakeInfo struct {
	Aggregator common.Address
	StakeInfo  IStakeManagerStakeInfo
}

// IEntryPointReturnInfo is an auto generated low-level Go binding around an user-defined struct.
type IEntryPointReturnInfo struct {
	PreOpGas                *big.Int
	Prefund                 *big.Int
	AccountValidationData   *big.Int
	PaymasterValidationData *big.Int
	PaymasterContext        []byte
}

// IEntryPointSimulationsExecutionResult is an auto generated low-level Go binding around an user-defined struct.
type IEntryPointSimulationsExecutionResult struct {
	PreOpGas                *big.Int
	Paid                    *big.Int
	AccountValidationData   *big.Int
	PaymasterValidationData *big.Int
	TargetSuccess           bool
	TargetResult            []byte
}

// IEntryPointSimulationsValidationResult is an auto generated low-level Go binding around an user-defined struct.
type IEntryPointSimulationsValidationResult struct {
	ReturnInfo     IEntryPointReturnInfo
	SenderInfo     IStakeManagerStakeInfo
	FactoryInfo    IStakeManagerStakeInfo
	PaymasterInfo  IStakeManagerStakeInfo
	AggregatorInfo IEntryPointAggregatorStakeInfo
}

// IEntryPointUserOpsPerAggregator is an auto generated low-level Go binding around an user-defined struct.
type IEntryPointUserOpsPerAggregator struct {
	UserOps    []PackedUserOperation
	Aggregator common.Address
	Signature  []byte
}

// IStakeManagerDepositInfo is an auto generated low-level Go binding around an user-defined struct.
type IStakeManagerDepositInfo struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}

// IStakeManagerStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type IStakeManagerStakeInfo struct {
	Stake           *big.Int
	UnstakeDelaySec *big.Int
}

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

// IEntryPointSimulationsMetaData contains all meta data concerning the IEntryPointSimulations contract.
var IEntryPointSimulationsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"name\":\"DelegateAndRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FailedOp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"inner\",\"type\":\"bytes\"}],\"name\":\"FailedOpWithRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"PostOpReverted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderAddressResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureValidationFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposit\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureAggregatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawTime\",\"type\":\"uint256\"}],\"name\":\"StakeUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"internalType\":\"structIStakeManager.DepositInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"name\":\"getSenderAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAggregator\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.UserOpsPerAggregator[]\",\"name\":\"opsPerAggregator\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleAggregatedOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"op\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetCallData\",\"type\":\"bytes\"}],\"name\":\"simulateHandleOp\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountValidationData\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymasterValidationData\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"targetSuccess\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"targetResult\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPointSimulations.ExecutionResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"simulateValidation\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountValidationData\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymasterValidationData\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterContext\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.ReturnInfo\",\"name\":\"returnInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"senderInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"factoryInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"paymasterInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"stakeInfo\",\"type\":\"tuple\"}],\"internalType\":\"structIEntryPoint.AggregatorStakeInfo\",\"name\":\"aggregatorInfo\",\"type\":\"tuple\"}],\"internalType\":\"structIEntryPointSimulations.ValidationResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IEntryPointSimulationsABI is the input ABI used to generate the binding from.
// Deprecated: Use IEntryPointSimulationsMetaData.ABI instead.
var IEntryPointSimulationsABI = IEntryPointSimulationsMetaData.ABI

// IEntryPointSimulations is an auto generated Go binding around an Ethereum contract.
type IEntryPointSimulations struct {
	IEntryPointSimulationsCaller     // Read-only binding to the contract
	IEntryPointSimulationsTransactor // Write-only binding to the contract
	IEntryPointSimulationsFilterer   // Log filterer for contract events
}

// IEntryPointSimulationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEntryPointSimulationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEntryPointSimulationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEntryPointSimulationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEntryPointSimulationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEntryPointSimulationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEntryPointSimulationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEntryPointSimulationsSession struct {
	Contract     *IEntryPointSimulations // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IEntryPointSimulationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEntryPointSimulationsCallerSession struct {
	Contract *IEntryPointSimulationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IEntryPointSimulationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEntryPointSimulationsTransactorSession struct {
	Contract     *IEntryPointSimulationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IEntryPointSimulationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEntryPointSimulationsRaw struct {
	Contract *IEntryPointSimulations // Generic contract binding to access the raw methods on
}

// IEntryPointSimulationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEntryPointSimulationsCallerRaw struct {
	Contract *IEntryPointSimulationsCaller // Generic read-only contract binding to access the raw methods on
}

// IEntryPointSimulationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEntryPointSimulationsTransactorRaw struct {
	Contract *IEntryPointSimulationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEntryPointSimulations creates a new instance of IEntryPointSimulations, bound to a specific deployed contract.
func NewIEntryPointSimulations(address common.Address, backend bind.ContractBackend) (*IEntryPointSimulations, error) {
	contract, err := bindIEntryPointSimulations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulations{IEntryPointSimulationsCaller: IEntryPointSimulationsCaller{contract: contract}, IEntryPointSimulationsTransactor: IEntryPointSimulationsTransactor{contract: contract}, IEntryPointSimulationsFilterer: IEntryPointSimulationsFilterer{contract: contract}}, nil
}

// NewIEntryPointSimulationsCaller creates a new read-only instance of IEntryPointSimulations, bound to a specific deployed contract.
func NewIEntryPointSimulationsCaller(address common.Address, caller bind.ContractCaller) (*IEntryPointSimulationsCaller, error) {
	contract, err := bindIEntryPointSimulations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsCaller{contract: contract}, nil
}

// NewIEntryPointSimulationsTransactor creates a new write-only instance of IEntryPointSimulations, bound to a specific deployed contract.
func NewIEntryPointSimulationsTransactor(address common.Address, transactor bind.ContractTransactor) (*IEntryPointSimulationsTransactor, error) {
	contract, err := bindIEntryPointSimulations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsTransactor{contract: contract}, nil
}

// NewIEntryPointSimulationsFilterer creates a new log filterer instance of IEntryPointSimulations, bound to a specific deployed contract.
func NewIEntryPointSimulationsFilterer(address common.Address, filterer bind.ContractFilterer) (*IEntryPointSimulationsFilterer, error) {
	contract, err := bindIEntryPointSimulations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsFilterer{contract: contract}, nil
}

// bindIEntryPointSimulations binds a generic wrapper to an already deployed contract.
func bindIEntryPointSimulations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEntryPointSimulationsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEntryPointSimulations *IEntryPointSimulationsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEntryPointSimulations.Contract.IEntryPointSimulationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEntryPointSimulations *IEntryPointSimulationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.IEntryPointSimulationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEntryPointSimulations *IEntryPointSimulationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.IEntryPointSimulationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEntryPointSimulations *IEntryPointSimulationsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEntryPointSimulations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IEntryPointSimulations *IEntryPointSimulationsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IEntryPointSimulations.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IEntryPointSimulations *IEntryPointSimulationsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IEntryPointSimulations.Contract.BalanceOf(&_IEntryPointSimulations.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IEntryPointSimulations *IEntryPointSimulationsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IEntryPointSimulations.Contract.BalanceOf(&_IEntryPointSimulations.CallOpts, account)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_IEntryPointSimulations *IEntryPointSimulationsCaller) GetDepositInfo(opts *bind.CallOpts, account common.Address) (IStakeManagerDepositInfo, error) {
	var out []interface{}
	err := _IEntryPointSimulations.contract.Call(opts, &out, "getDepositInfo", account)

	if err != nil {
		return *new(IStakeManagerDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeManagerDepositInfo)).(*IStakeManagerDepositInfo)

	return out0, err

}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_IEntryPointSimulations *IEntryPointSimulationsSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _IEntryPointSimulations.Contract.GetDepositInfo(&_IEntryPointSimulations.CallOpts, account)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_IEntryPointSimulations *IEntryPointSimulationsCallerSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _IEntryPointSimulations.Contract.GetDepositInfo(&_IEntryPointSimulations.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_IEntryPointSimulations *IEntryPointSimulationsCaller) GetNonce(opts *bind.CallOpts, sender common.Address, key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IEntryPointSimulations.contract.Call(opts, &out, "getNonce", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_IEntryPointSimulations *IEntryPointSimulationsSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _IEntryPointSimulations.Contract.GetNonce(&_IEntryPointSimulations.CallOpts, sender, key)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_IEntryPointSimulations *IEntryPointSimulationsCallerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _IEntryPointSimulations.Contract.GetNonce(&_IEntryPointSimulations.CallOpts, sender, key)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_IEntryPointSimulations *IEntryPointSimulationsCaller) GetUserOpHash(opts *bind.CallOpts, userOp PackedUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _IEntryPointSimulations.contract.Call(opts, &out, "getUserOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHash is a free data retrieval call binding the contract method 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_IEntryPointSimulations *IEntryPointSimulationsSession) GetUserOpHash(userOp PackedUserOperation) ([32]byte, error) {
	return _IEntryPointSimulations.Contract.GetUserOpHash(&_IEntryPointSimulations.CallOpts, userOp)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_IEntryPointSimulations *IEntryPointSimulationsCallerSession) GetUserOpHash(userOp PackedUserOperation) ([32]byte, error) {
	return _IEntryPointSimulations.Contract.GetUserOpHash(&_IEntryPointSimulations.CallOpts, userOp)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 _unstakeDelaySec) payable returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactor) AddStake(opts *bind.TransactOpts, _unstakeDelaySec uint32) (*types.Transaction, error) {
	return _IEntryPointSimulations.contract.Transact(opts, "addStake", _unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 _unstakeDelaySec) payable returns()
func (_IEntryPointSimulations *IEntryPointSimulationsSession) AddStake(_unstakeDelaySec uint32) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.AddStake(&_IEntryPointSimulations.TransactOpts, _unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 _unstakeDelaySec) payable returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorSession) AddStake(_unstakeDelaySec uint32) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.AddStake(&_IEntryPointSimulations.TransactOpts, _unstakeDelaySec)
}

// DelegateAndRevert is a paid mutator transaction binding the contract method 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactor) DelegateAndRevert(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _IEntryPointSimulations.contract.Transact(opts, "delegateAndRevert", target, data)
}

// DelegateAndRevert is a paid mutator transaction binding the contract method 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsSession) DelegateAndRevert(target common.Address, data []byte) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.DelegateAndRevert(&_IEntryPointSimulations.TransactOpts, target, data)
}

// DelegateAndRevert is a paid mutator transaction binding the contract method 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorSession) DelegateAndRevert(target common.Address, data []byte) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.DelegateAndRevert(&_IEntryPointSimulations.TransactOpts, target, data)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactor) DepositTo(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IEntryPointSimulations.contract.Transact(opts, "depositTo", account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_IEntryPointSimulations *IEntryPointSimulationsSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.DepositTo(&_IEntryPointSimulations.TransactOpts, account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.DepositTo(&_IEntryPointSimulations.TransactOpts, account)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactor) GetSenderAddress(opts *bind.TransactOpts, initCode []byte) (*types.Transaction, error) {
	return _IEntryPointSimulations.contract.Transact(opts, "getSenderAddress", initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsSession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.GetSenderAddress(&_IEntryPointSimulations.TransactOpts, initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorSession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.GetSenderAddress(&_IEntryPointSimulations.TransactOpts, initCode)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0xdbed18e0.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactor) HandleAggregatedOps(opts *bind.TransactOpts, opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _IEntryPointSimulations.contract.Transact(opts, "handleAggregatedOps", opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0xdbed18e0.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsSession) HandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.HandleAggregatedOps(&_IEntryPointSimulations.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0xdbed18e0.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorSession) HandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.HandleAggregatedOps(&_IEntryPointSimulations.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x765e827f.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address beneficiary) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactor) HandleOps(opts *bind.TransactOpts, ops []PackedUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _IEntryPointSimulations.contract.Transact(opts, "handleOps", ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x765e827f.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address beneficiary) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsSession) HandleOps(ops []PackedUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.HandleOps(&_IEntryPointSimulations.TransactOpts, ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x765e827f.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address beneficiary) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorSession) HandleOps(ops []PackedUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.HandleOps(&_IEntryPointSimulations.TransactOpts, ops, beneficiary)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactor) IncrementNonce(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _IEntryPointSimulations.contract.Transact(opts, "incrementNonce", key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.IncrementNonce(&_IEntryPointSimulations.TransactOpts, key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.IncrementNonce(&_IEntryPointSimulations.TransactOpts, key)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0x97b2dcb9.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) op, address target, bytes targetCallData) returns((uint256,uint256,uint256,uint256,bool,bytes))
func (_IEntryPointSimulations *IEntryPointSimulationsTransactor) SimulateHandleOp(opts *bind.TransactOpts, op PackedUserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _IEntryPointSimulations.contract.Transact(opts, "simulateHandleOp", op, target, targetCallData)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0x97b2dcb9.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) op, address target, bytes targetCallData) returns((uint256,uint256,uint256,uint256,bool,bytes))
func (_IEntryPointSimulations *IEntryPointSimulationsSession) SimulateHandleOp(op PackedUserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.SimulateHandleOp(&_IEntryPointSimulations.TransactOpts, op, target, targetCallData)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0x97b2dcb9.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) op, address target, bytes targetCallData) returns((uint256,uint256,uint256,uint256,bool,bytes))
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorSession) SimulateHandleOp(op PackedUserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.SimulateHandleOp(&_IEntryPointSimulations.TransactOpts, op, target, targetCallData)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xc3bce009.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) returns(((uint256,uint256,uint256,uint256,bytes),(uint256,uint256),(uint256,uint256),(uint256,uint256),(address,(uint256,uint256))))
func (_IEntryPointSimulations *IEntryPointSimulationsTransactor) SimulateValidation(opts *bind.TransactOpts, userOp PackedUserOperation) (*types.Transaction, error) {
	return _IEntryPointSimulations.contract.Transact(opts, "simulateValidation", userOp)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xc3bce009.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) returns(((uint256,uint256,uint256,uint256,bytes),(uint256,uint256),(uint256,uint256),(uint256,uint256),(address,(uint256,uint256))))
func (_IEntryPointSimulations *IEntryPointSimulationsSession) SimulateValidation(userOp PackedUserOperation) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.SimulateValidation(&_IEntryPointSimulations.TransactOpts, userOp)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xc3bce009.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) returns(((uint256,uint256,uint256,uint256,bytes),(uint256,uint256),(uint256,uint256),(uint256,uint256),(address,(uint256,uint256))))
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorSession) SimulateValidation(userOp PackedUserOperation) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.SimulateValidation(&_IEntryPointSimulations.TransactOpts, userOp)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEntryPointSimulations.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_IEntryPointSimulations *IEntryPointSimulationsSession) UnlockStake() (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.UnlockStake(&_IEntryPointSimulations.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.UnlockStake(&_IEntryPointSimulations.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _IEntryPointSimulations.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.WithdrawStake(&_IEntryPointSimulations.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.WithdrawStake(&_IEntryPointSimulations.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _IEntryPointSimulations.contract.Transact(opts, "withdrawTo", withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.WithdrawTo(&_IEntryPointSimulations.TransactOpts, withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_IEntryPointSimulations *IEntryPointSimulationsTransactorSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _IEntryPointSimulations.Contract.WithdrawTo(&_IEntryPointSimulations.TransactOpts, withdrawAddress, withdrawAmount)
}

// IEntryPointSimulationsAccountDeployedIterator is returned from FilterAccountDeployed and is used to iterate over the raw logs and unpacked data for AccountDeployed events raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsAccountDeployedIterator struct {
	Event *IEntryPointSimulationsAccountDeployed // Event containing the contract specifics and raw log

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
func (it *IEntryPointSimulationsAccountDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointSimulationsAccountDeployed)
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
		it.Event = new(IEntryPointSimulationsAccountDeployed)
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
func (it *IEntryPointSimulationsAccountDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointSimulationsAccountDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointSimulationsAccountDeployed represents a AccountDeployed event raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsAccountDeployed struct {
	UserOpHash [32]byte
	Sender     common.Address
	Factory    common.Address
	Paymaster  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountDeployed is a free log retrieval operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) FilterAccountDeployed(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*IEntryPointSimulationsAccountDeployedIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.FilterLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsAccountDeployedIterator{contract: _IEntryPointSimulations.contract, event: "AccountDeployed", logs: logs, sub: sub}, nil
}

// WatchAccountDeployed is a free log subscription operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) WatchAccountDeployed(opts *bind.WatchOpts, sink chan<- *IEntryPointSimulationsAccountDeployed, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.WatchLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointSimulationsAccountDeployed)
				if err := _IEntryPointSimulations.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
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

// ParseAccountDeployed is a log parse operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) ParseAccountDeployed(log types.Log) (*IEntryPointSimulationsAccountDeployed, error) {
	event := new(IEntryPointSimulationsAccountDeployed)
	if err := _IEntryPointSimulations.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointSimulationsBeforeExecutionIterator is returned from FilterBeforeExecution and is used to iterate over the raw logs and unpacked data for BeforeExecution events raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsBeforeExecutionIterator struct {
	Event *IEntryPointSimulationsBeforeExecution // Event containing the contract specifics and raw log

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
func (it *IEntryPointSimulationsBeforeExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointSimulationsBeforeExecution)
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
		it.Event = new(IEntryPointSimulationsBeforeExecution)
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
func (it *IEntryPointSimulationsBeforeExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointSimulationsBeforeExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointSimulationsBeforeExecution represents a BeforeExecution event raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsBeforeExecution struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBeforeExecution is a free log retrieval operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) FilterBeforeExecution(opts *bind.FilterOpts) (*IEntryPointSimulationsBeforeExecutionIterator, error) {

	logs, sub, err := _IEntryPointSimulations.contract.FilterLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsBeforeExecutionIterator{contract: _IEntryPointSimulations.contract, event: "BeforeExecution", logs: logs, sub: sub}, nil
}

// WatchBeforeExecution is a free log subscription operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) WatchBeforeExecution(opts *bind.WatchOpts, sink chan<- *IEntryPointSimulationsBeforeExecution) (event.Subscription, error) {

	logs, sub, err := _IEntryPointSimulations.contract.WatchLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointSimulationsBeforeExecution)
				if err := _IEntryPointSimulations.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
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

// ParseBeforeExecution is a log parse operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) ParseBeforeExecution(log types.Log) (*IEntryPointSimulationsBeforeExecution, error) {
	event := new(IEntryPointSimulationsBeforeExecution)
	if err := _IEntryPointSimulations.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointSimulationsDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsDepositedIterator struct {
	Event *IEntryPointSimulationsDeposited // Event containing the contract specifics and raw log

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
func (it *IEntryPointSimulationsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointSimulationsDeposited)
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
		it.Event = new(IEntryPointSimulationsDeposited)
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
func (it *IEntryPointSimulationsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointSimulationsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointSimulationsDeposited represents a Deposited event raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) FilterDeposited(opts *bind.FilterOpts, account []common.Address) (*IEntryPointSimulationsDepositedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.FilterLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsDepositedIterator{contract: _IEntryPointSimulations.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *IEntryPointSimulationsDeposited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.WatchLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointSimulationsDeposited)
				if err := _IEntryPointSimulations.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) ParseDeposited(log types.Log) (*IEntryPointSimulationsDeposited, error) {
	event := new(IEntryPointSimulationsDeposited)
	if err := _IEntryPointSimulations.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointSimulationsPostOpRevertReasonIterator is returned from FilterPostOpRevertReason and is used to iterate over the raw logs and unpacked data for PostOpRevertReason events raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsPostOpRevertReasonIterator struct {
	Event *IEntryPointSimulationsPostOpRevertReason // Event containing the contract specifics and raw log

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
func (it *IEntryPointSimulationsPostOpRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointSimulationsPostOpRevertReason)
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
		it.Event = new(IEntryPointSimulationsPostOpRevertReason)
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
func (it *IEntryPointSimulationsPostOpRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointSimulationsPostOpRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointSimulationsPostOpRevertReason represents a PostOpRevertReason event raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsPostOpRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPostOpRevertReason is a free log retrieval operation binding the contract event 0xf62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f4792.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) FilterPostOpRevertReason(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*IEntryPointSimulationsPostOpRevertReasonIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.FilterLogs(opts, "PostOpRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsPostOpRevertReasonIterator{contract: _IEntryPointSimulations.contract, event: "PostOpRevertReason", logs: logs, sub: sub}, nil
}

// WatchPostOpRevertReason is a free log subscription operation binding the contract event 0xf62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f4792.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) WatchPostOpRevertReason(opts *bind.WatchOpts, sink chan<- *IEntryPointSimulationsPostOpRevertReason, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.WatchLogs(opts, "PostOpRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointSimulationsPostOpRevertReason)
				if err := _IEntryPointSimulations.contract.UnpackLog(event, "PostOpRevertReason", log); err != nil {
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

// ParsePostOpRevertReason is a log parse operation binding the contract event 0xf62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f4792.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) ParsePostOpRevertReason(log types.Log) (*IEntryPointSimulationsPostOpRevertReason, error) {
	event := new(IEntryPointSimulationsPostOpRevertReason)
	if err := _IEntryPointSimulations.contract.UnpackLog(event, "PostOpRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointSimulationsSignatureAggregatorChangedIterator is returned from FilterSignatureAggregatorChanged and is used to iterate over the raw logs and unpacked data for SignatureAggregatorChanged events raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsSignatureAggregatorChangedIterator struct {
	Event *IEntryPointSimulationsSignatureAggregatorChanged // Event containing the contract specifics and raw log

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
func (it *IEntryPointSimulationsSignatureAggregatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointSimulationsSignatureAggregatorChanged)
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
		it.Event = new(IEntryPointSimulationsSignatureAggregatorChanged)
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
func (it *IEntryPointSimulationsSignatureAggregatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointSimulationsSignatureAggregatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointSimulationsSignatureAggregatorChanged represents a SignatureAggregatorChanged event raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsSignatureAggregatorChanged struct {
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSignatureAggregatorChanged is a free log retrieval operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) FilterSignatureAggregatorChanged(opts *bind.FilterOpts, aggregator []common.Address) (*IEntryPointSimulationsSignatureAggregatorChangedIterator, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.FilterLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsSignatureAggregatorChangedIterator{contract: _IEntryPointSimulations.contract, event: "SignatureAggregatorChanged", logs: logs, sub: sub}, nil
}

// WatchSignatureAggregatorChanged is a free log subscription operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) WatchSignatureAggregatorChanged(opts *bind.WatchOpts, sink chan<- *IEntryPointSimulationsSignatureAggregatorChanged, aggregator []common.Address) (event.Subscription, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.WatchLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointSimulationsSignatureAggregatorChanged)
				if err := _IEntryPointSimulations.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
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

// ParseSignatureAggregatorChanged is a log parse operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) ParseSignatureAggregatorChanged(log types.Log) (*IEntryPointSimulationsSignatureAggregatorChanged, error) {
	event := new(IEntryPointSimulationsSignatureAggregatorChanged)
	if err := _IEntryPointSimulations.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointSimulationsStakeLockedIterator is returned from FilterStakeLocked and is used to iterate over the raw logs and unpacked data for StakeLocked events raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsStakeLockedIterator struct {
	Event *IEntryPointSimulationsStakeLocked // Event containing the contract specifics and raw log

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
func (it *IEntryPointSimulationsStakeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointSimulationsStakeLocked)
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
		it.Event = new(IEntryPointSimulationsStakeLocked)
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
func (it *IEntryPointSimulationsStakeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointSimulationsStakeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointSimulationsStakeLocked represents a StakeLocked event raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeLocked is a free log retrieval operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) FilterStakeLocked(opts *bind.FilterOpts, account []common.Address) (*IEntryPointSimulationsStakeLockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.FilterLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsStakeLockedIterator{contract: _IEntryPointSimulations.contract, event: "StakeLocked", logs: logs, sub: sub}, nil
}

// WatchStakeLocked is a free log subscription operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) WatchStakeLocked(opts *bind.WatchOpts, sink chan<- *IEntryPointSimulationsStakeLocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.WatchLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointSimulationsStakeLocked)
				if err := _IEntryPointSimulations.contract.UnpackLog(event, "StakeLocked", log); err != nil {
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

// ParseStakeLocked is a log parse operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) ParseStakeLocked(log types.Log) (*IEntryPointSimulationsStakeLocked, error) {
	event := new(IEntryPointSimulationsStakeLocked)
	if err := _IEntryPointSimulations.contract.UnpackLog(event, "StakeLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointSimulationsStakeUnlockedIterator is returned from FilterStakeUnlocked and is used to iterate over the raw logs and unpacked data for StakeUnlocked events raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsStakeUnlockedIterator struct {
	Event *IEntryPointSimulationsStakeUnlocked // Event containing the contract specifics and raw log

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
func (it *IEntryPointSimulationsStakeUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointSimulationsStakeUnlocked)
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
		it.Event = new(IEntryPointSimulationsStakeUnlocked)
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
func (it *IEntryPointSimulationsStakeUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointSimulationsStakeUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointSimulationsStakeUnlocked represents a StakeUnlocked event raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeUnlocked is a free log retrieval operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) FilterStakeUnlocked(opts *bind.FilterOpts, account []common.Address) (*IEntryPointSimulationsStakeUnlockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.FilterLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsStakeUnlockedIterator{contract: _IEntryPointSimulations.contract, event: "StakeUnlocked", logs: logs, sub: sub}, nil
}

// WatchStakeUnlocked is a free log subscription operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) WatchStakeUnlocked(opts *bind.WatchOpts, sink chan<- *IEntryPointSimulationsStakeUnlocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.WatchLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointSimulationsStakeUnlocked)
				if err := _IEntryPointSimulations.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
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

// ParseStakeUnlocked is a log parse operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) ParseStakeUnlocked(log types.Log) (*IEntryPointSimulationsStakeUnlocked, error) {
	event := new(IEntryPointSimulationsStakeUnlocked)
	if err := _IEntryPointSimulations.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointSimulationsStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsStakeWithdrawnIterator struct {
	Event *IEntryPointSimulationsStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *IEntryPointSimulationsStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointSimulationsStakeWithdrawn)
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
		it.Event = new(IEntryPointSimulationsStakeWithdrawn)
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
func (it *IEntryPointSimulationsStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointSimulationsStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointSimulationsStakeWithdrawn represents a StakeWithdrawn event raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, account []common.Address) (*IEntryPointSimulationsStakeWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.FilterLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsStakeWithdrawnIterator{contract: _IEntryPointSimulations.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *IEntryPointSimulationsStakeWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.WatchLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointSimulationsStakeWithdrawn)
				if err := _IEntryPointSimulations.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) ParseStakeWithdrawn(log types.Log) (*IEntryPointSimulationsStakeWithdrawn, error) {
	event := new(IEntryPointSimulationsStakeWithdrawn)
	if err := _IEntryPointSimulations.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointSimulationsUserOperationEventIterator is returned from FilterUserOperationEvent and is used to iterate over the raw logs and unpacked data for UserOperationEvent events raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsUserOperationEventIterator struct {
	Event *IEntryPointSimulationsUserOperationEvent // Event containing the contract specifics and raw log

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
func (it *IEntryPointSimulationsUserOperationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointSimulationsUserOperationEvent)
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
		it.Event = new(IEntryPointSimulationsUserOperationEvent)
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
func (it *IEntryPointSimulationsUserOperationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointSimulationsUserOperationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointSimulationsUserOperationEvent represents a UserOperationEvent event raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsUserOperationEvent struct {
	UserOpHash    [32]byte
	Sender        common.Address
	Paymaster     common.Address
	Nonce         *big.Int
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserOperationEvent is a free log retrieval operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) FilterUserOperationEvent(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (*IEntryPointSimulationsUserOperationEventIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.FilterLogs(opts, "UserOperationEvent", userOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsUserOperationEventIterator{contract: _IEntryPointSimulations.contract, event: "UserOperationEvent", logs: logs, sub: sub}, nil
}

// WatchUserOperationEvent is a free log subscription operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) WatchUserOperationEvent(opts *bind.WatchOpts, sink chan<- *IEntryPointSimulationsUserOperationEvent, userOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.WatchLogs(opts, "UserOperationEvent", userOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointSimulationsUserOperationEvent)
				if err := _IEntryPointSimulations.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
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

// ParseUserOperationEvent is a log parse operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) ParseUserOperationEvent(log types.Log) (*IEntryPointSimulationsUserOperationEvent, error) {
	event := new(IEntryPointSimulationsUserOperationEvent)
	if err := _IEntryPointSimulations.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointSimulationsUserOperationPrefundTooLowIterator is returned from FilterUserOperationPrefundTooLow and is used to iterate over the raw logs and unpacked data for UserOperationPrefundTooLow events raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsUserOperationPrefundTooLowIterator struct {
	Event *IEntryPointSimulationsUserOperationPrefundTooLow // Event containing the contract specifics and raw log

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
func (it *IEntryPointSimulationsUserOperationPrefundTooLowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointSimulationsUserOperationPrefundTooLow)
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
		it.Event = new(IEntryPointSimulationsUserOperationPrefundTooLow)
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
func (it *IEntryPointSimulationsUserOperationPrefundTooLowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointSimulationsUserOperationPrefundTooLowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointSimulationsUserOperationPrefundTooLow represents a UserOperationPrefundTooLow event raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsUserOperationPrefundTooLow struct {
	UserOpHash [32]byte
	Sender     common.Address
	Nonce      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserOperationPrefundTooLow is a free log retrieval operation binding the contract event 0x67b4fa9642f42120bf031f3051d1824b0fe25627945b27b8a6a65d5761d5482e.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender, uint256 nonce)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) FilterUserOperationPrefundTooLow(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*IEntryPointSimulationsUserOperationPrefundTooLowIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.FilterLogs(opts, "UserOperationPrefundTooLow", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsUserOperationPrefundTooLowIterator{contract: _IEntryPointSimulations.contract, event: "UserOperationPrefundTooLow", logs: logs, sub: sub}, nil
}

// WatchUserOperationPrefundTooLow is a free log subscription operation binding the contract event 0x67b4fa9642f42120bf031f3051d1824b0fe25627945b27b8a6a65d5761d5482e.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender, uint256 nonce)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) WatchUserOperationPrefundTooLow(opts *bind.WatchOpts, sink chan<- *IEntryPointSimulationsUserOperationPrefundTooLow, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.WatchLogs(opts, "UserOperationPrefundTooLow", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointSimulationsUserOperationPrefundTooLow)
				if err := _IEntryPointSimulations.contract.UnpackLog(event, "UserOperationPrefundTooLow", log); err != nil {
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

// ParseUserOperationPrefundTooLow is a log parse operation binding the contract event 0x67b4fa9642f42120bf031f3051d1824b0fe25627945b27b8a6a65d5761d5482e.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender, uint256 nonce)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) ParseUserOperationPrefundTooLow(log types.Log) (*IEntryPointSimulationsUserOperationPrefundTooLow, error) {
	event := new(IEntryPointSimulationsUserOperationPrefundTooLow)
	if err := _IEntryPointSimulations.contract.UnpackLog(event, "UserOperationPrefundTooLow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointSimulationsUserOperationRevertReasonIterator is returned from FilterUserOperationRevertReason and is used to iterate over the raw logs and unpacked data for UserOperationRevertReason events raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsUserOperationRevertReasonIterator struct {
	Event *IEntryPointSimulationsUserOperationRevertReason // Event containing the contract specifics and raw log

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
func (it *IEntryPointSimulationsUserOperationRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointSimulationsUserOperationRevertReason)
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
		it.Event = new(IEntryPointSimulationsUserOperationRevertReason)
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
func (it *IEntryPointSimulationsUserOperationRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointSimulationsUserOperationRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointSimulationsUserOperationRevertReason represents a UserOperationRevertReason event raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsUserOperationRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserOperationRevertReason is a free log retrieval operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) FilterUserOperationRevertReason(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*IEntryPointSimulationsUserOperationRevertReasonIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.FilterLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsUserOperationRevertReasonIterator{contract: _IEntryPointSimulations.contract, event: "UserOperationRevertReason", logs: logs, sub: sub}, nil
}

// WatchUserOperationRevertReason is a free log subscription operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) WatchUserOperationRevertReason(opts *bind.WatchOpts, sink chan<- *IEntryPointSimulationsUserOperationRevertReason, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.WatchLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointSimulationsUserOperationRevertReason)
				if err := _IEntryPointSimulations.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
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

// ParseUserOperationRevertReason is a log parse operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) ParseUserOperationRevertReason(log types.Log) (*IEntryPointSimulationsUserOperationRevertReason, error) {
	event := new(IEntryPointSimulationsUserOperationRevertReason)
	if err := _IEntryPointSimulations.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointSimulationsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsWithdrawnIterator struct {
	Event *IEntryPointSimulationsWithdrawn // Event containing the contract specifics and raw log

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
func (it *IEntryPointSimulationsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointSimulationsWithdrawn)
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
		it.Event = new(IEntryPointSimulationsWithdrawn)
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
func (it *IEntryPointSimulationsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointSimulationsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointSimulationsWithdrawn represents a Withdrawn event raised by the IEntryPointSimulations contract.
type IEntryPointSimulationsWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*IEntryPointSimulationsWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointSimulationsWithdrawnIterator{contract: _IEntryPointSimulations.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *IEntryPointSimulationsWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPointSimulations.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointSimulationsWithdrawn)
				if err := _IEntryPointSimulations.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_IEntryPointSimulations *IEntryPointSimulationsFilterer) ParseWithdrawn(log types.Log) (*IEntryPointSimulationsWithdrawn, error) {
	event := new(IEntryPointSimulationsWithdrawn)
	if err := _IEntryPointSimulations.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
