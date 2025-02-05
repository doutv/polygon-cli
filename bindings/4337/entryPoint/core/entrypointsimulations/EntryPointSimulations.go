// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package entrypointsimulations

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

// EntryPointMemoryUserOp is an auto generated low-level Go binding around an user-defined struct.
type EntryPointMemoryUserOp struct {
	Sender                        common.Address
	Nonce                         *big.Int
	VerificationGasLimit          *big.Int
	CallGasLimit                  *big.Int
	PaymasterVerificationGasLimit *big.Int
	PaymasterPostOpGasLimit       *big.Int
	PreVerificationGas            *big.Int
	Paymaster                     common.Address
	MaxFeePerGas                  *big.Int
	MaxPriorityFeePerGas          *big.Int
}

// EntryPointUserOpInfo is an auto generated low-level Go binding around an user-defined struct.
type EntryPointUserOpInfo struct {
	MUserOp       EntryPointMemoryUserOp
	UserOpHash    [32]byte
	Prefund       *big.Int
	ContextOffset *big.Int
	PreOpGas      *big.Int
}

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

// EntryPointSimulationsMetaData contains all meta data concerning the EntryPointSimulations contract.
var EntryPointSimulationsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"name\":\"DelegateAndRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FailedOp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"inner\",\"type\":\"bytes\"}],\"name\":\"FailedOpWithRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"PostOpReverted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderAddressResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureValidationFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposit\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureAggregatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawTime\",\"type\":\"uint256\"}],\"name\":\"StakeUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"}],\"name\":\"_validateSenderAndPaymaster\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"internalType\":\"structIStakeManager.DepositInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"name\":\"getSenderAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAggregator\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.UserOpsPerAggregator[]\",\"name\":\"opsPerAggregator\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleAggregatedOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymasterVerificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymasterPostOpGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"}],\"internalType\":\"structEntryPoint.MemoryUserOp\",\"name\":\"mUserOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contextOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"}],\"internalType\":\"structEntryPoint.UserOpInfo\",\"name\":\"opInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"innerHandleOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"name\":\"nonceSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"op\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetCallData\",\"type\":\"bytes\"}],\"name\":\"simulateHandleOp\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountValidationData\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymasterValidationData\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"targetSuccess\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"targetResult\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPointSimulations.ExecutionResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"simulateValidation\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountValidationData\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymasterValidationData\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterContext\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.ReturnInfo\",\"name\":\"returnInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"senderInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"factoryInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"paymasterInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"stakeInfo\",\"type\":\"tuple\"}],\"internalType\":\"structIEntryPoint.AggregatorStakeInfo\",\"name\":\"aggregatorInfo\",\"type\":\"tuple\"}],\"internalType\":\"structIEntryPointSimulations.ValidationResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a080604052346100ff57600160025561015f8181016001600160401b038111838210176100e957829161485e833903906000f080156100dd57608052610044610104565b6000815260208101906000825280602061005c610104565b600081520152600380546001600160a01b0319169055516004555160055560644310156100985760405161473a90816101248239608051815050f35b60405162461bcd60e51b815260206004820152601660248201527f73686f756c64206e6f74206265206465706c6f796564000000000000000000006044820152606490fd5b6040513d6000823e3d90fd5b634e487b7160e01b600052604160045260246000fd5b600080fd5b60408051919082016001600160401b038111838210176100e95760405256fe60806040526004361015610024575b361561001957600080fd5b61002233613406565b005b60003560e01c806242dc5314612a0f57806301ffc9a7146129715780630396cb601461271c5780630bd28e3b146126ce5780631b2e01b814612678578063205c28781461255a57806322cdde4c1461253a57806335567e1a146124d55780635287ce12146123fa57806370a08231146123c0578063765e827f146121a7578063850aaf6214612125578063957122ab14611feb57806397b2dcb9146115de5780639b249f6914611522578063b760faf914611506578063bb9fe6bf146113df578063c23a5cea1461120e578063c3bce00914610653578063dbed18e0146101905763fc7e286d0361000e573461018b57602036600319011261018b576001600160a01b03610130613033565b16600052600060205260a0604060002065ffffffffffff6001825492015460405192835260ff8116151560208401526001600160701b038160081c16604084015263ffffffff8160781c16606084015260981c166080820152f35b600080fd5b3461018b5761019e366130be565b906101a76136cd565b60009160005b8281106104ac57506101bf8493613329565b6000805b8481106102de5750507fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972600080a16000809360005b8181106102365761022f868660007f575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d8180a261427b565b6001600255005b61028461024482848a61348f565b6001600160a01b03610258602083016134e6565b167f575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d600080a2806134b1565b906000915b80831061029b575050506001016101f8565b909194976102d56102cf6001926102c98c8b6102c2826102bc8e8b8d6133a3565b92613379565b5191613fce565b90613167565b99613174565b95019190610289565b60206102eb82878961348f565b6103016102f882806134b1565b939092016134e6565b600092906001600160a01b039081165b8285106103255750505050506001016101c3565b90919293956103548361034d61033b848c613379565b516103478b898b6133a3565b856136ee565b9290614678565b91168403610469576104235761036a8491614678565b91166103e05761038a5761037f600191613174565b960193929190610311565b60a48760405190631101335b60e11b8252600482015260406024820152602160448201527f41413332207061796d61737465722065787069726564206f72206e6f742064756064820152606560f81b6084820152fd5b60848860405190631101335b60e11b8252600482015260406024820152601460448201527320a0999a1039b4b3b730ba3ab9329032b93937b960611b6064820152fd5b60848860405190631101335b60e11b82526004820152604060248201526017604482015276414132322065787069726564206f72206e6f742064756560481b6064820152fd5b60848960405190631101335b60e11b8252600482015260406024820152601460448201527320a0991a1039b4b3b730ba3ab9329032b93937b960611b6064820152fd5b6104b781848761348f565b936104c285806134b1565b909590919060206001600160a01b036104dc8284016134e6565b1697600192838a1461060e578961050c575b50505050600192939495509061050391613167565b939291016101ad565b80604061051a92019061345d565b918a3b1561018b5792939190604051948593632dd8113360e01b855288604486016040600488015252606490818601918a60051b8701019680936000915b8c83106105c957505050505083850360031901602485015250600093839261057f926133e5565b03818a5afa90816105ba575b506105a95760405163086a9f7560e41b815260048101879052602490fd5b9394508493610503600189806104ee565b6105c390612f7c565b8861058b565b9193959697509193976063198a8203018652883561011e198336030181121561018b57836105fb87938583940161352b565b9a01960193019091899796959492610558565b60405162461bcd60e51b815260048101849052601760248201527f4141393620696e76616c69642061676772656761746f720000000000000000006044820152606490fd5b3461018b576106613661308c565b60405161066d81612ef9565b60405161067981612ef9565b600081526000602082015260006040820152600060608201526060608082015281526040516106a781612f61565b600081526000602082015260208201526040516106c381612f61565b600081526000602082015260408201526040516106df81612f61565b600081526000602082015260608201526080604051916106fe83612f61565b6000835260405161070e81612f61565b6000815260006020820152602084015201526107286132b8565b906107328161439b565b6000905a835193906001600160a01b0361074b846134e6565b1685526020830135602086015260808301356001600160801b038116606087015260801c604086015260a083013560c08601526001600160801b0360c08401351661010086015260c083013560801c6101208601526107ad60e084018461345d565b80156111f257603481106111ad578060141161018b578060241161018b5760341161018b57602481013560801c60a0870152601481013560801c60808701523560601c60e08601525b6107ff836131c0565b60208301526040850151946001600160781b038660c08301511760608301511760808301511760a08301511761010083015117610120830151171161116857604081015160608201510160808201510160a08201510160c0820151016101008201510283519660018060a01b038851169761087d604088018861345d565b80610f37575b505060e001516001600160a01b03169760008915610eff575b6020878101516040516306608bdf60e21b815292839182916108c2918d60048501614656565b038160008688f160009181610ecb575b506109565761094a3d61080080821161094e575b50604051906020818301016040528082526000602083013e6040516365c8fd4d60e01b81526000600482015260606024820152600d60648201526c10504c8cc81c995d995c9d1959609a1b608482015260a0604482015291829160a4830190613142565b0390fd5b9050826108e6565b9815610e64575b5060018060a01b03835116602084015160401c90600052600160205260406000209060005260205260406000206109948154613174565b90555a840311610e185760e090910151606091906001600160a01b0316610bdd575b6040840152606083015260a0830135905a9003016080820152805160e001516109e7906001600160a01b0316614513565b91610a1e6109fe60018060a01b0384515116614513565b9160006020604051610a0f81612f61565b8281520152604081019061345d565b60148110610bd15760141161018b57610a3a903560601c614513565b9160018060a01b03861694608082015196606060408401519301519260405198610a638a612ef9565b89526020890152604088015260608701526080860152604051610a8581612f61565b6003546001600160a01b03168152604051610a9f81612f61565b6004548152600554602082015260208201529380151580610bc6575b610b9f575b5060405194610ace86612ef9565b855260208501526040840152606083015260808201526040518091602082526020806080610b3c8185516101408589015280516101608901528481015161018089015260408101516101a089015260608101516101c0890152015160a06101e0880152610200870190613142565b9382808201518051604089015201516060870152826040820151805184890152015160a0870152826060820151805160c0890152015160e0870152015160018060a01b038151166101008601520151805161012085015201516101408301520390f35b909350610bab81614513565b60405191610bb883612f61565b825260208201529285610ac0565b506001811415610abb565b5050610a3a6000614513565b9450505a825160e08101516001600160a01b03166000818152602081905260409020805497939290848910610dcc576080600092869283610c429c0390550151926020880151838a6040519c8d95869485936314add44b60e21b855260048501614656565b039286f19687600091600099610d40575b50610cd35761094a3d610800808211610ccb575b50604051906020818301016040528082526000602083013e6040516365c8fd4d60e01b81526000600482015260606024820152600d60648201526c10504cccc81c995d995c9d1959609a1b608482015260a0604482015291829160a4830190613142565b905082610c67565b96915a900311610ce45794906109b6565b60a4604051631101335b60e11b81526000600482015260406024820152602760448201527f41413336206f766572207061796d6173746572566572696669636174696f6e47606482015266185cd31a5b5a5d60ca1b6084820152fd5b915097503d90816000823e610d558282612faa565b604081838101031261018b578051906001600160401b03821161018b57828101601f83830101121561018b578181015191610d8f83612fcb565b93610d9d6040519586612faa565b838552820160208483850101011161018b57602092610dc391848087019185010161311f565b0151978a610c53565b6084604051631101335b60e11b81526000600482015260406024820152601e60448201527f41413331207061796d6173746572206465706f73697420746f6f206c6f7700006064820152fd5b6084604051631101335b60e11b81526000600482015260406024820152601e60448201527f41413236206f76657220766572696669636174696f6e4761734c696d697400006064820152fd5b600052600060205260406000208054808411610e855783900390558861095d565b6084604051631101335b60e11b81526000600482015260406024820152601760448201527610504c8c48191a591b89dd081c185e481c1c99599d5b99604a1b6064820152fd5b9091506020813d602011610ef7575b81610ee760209383612faa565b8101031261018b5751908b6108d2565b3d9150610eda565b50806000526000602052604060002054838111600014610f2a57506108c2602060005b91505061089c565b60206108c2918503610f22565b8a3b61111c576000602060018060a01b036006541660408b51015160405180948193632b870d1b60e11b835285600484015282610f78602482018a8c6133e5565b0393f1908115611110576000916110e1575b506001600160a01b0381168015611095578c03611049573b15610ffd5760141161018b5789907fd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d604060208a01519260018060a01b0360e08c510151168251913560601c82526020820152a38980610883565b6084604051631101335b60e11b81526000600482015260406024820152602060448201527f4141313520696e6974436f6465206d757374206372656174652073656e6465726064820152fd5b6084604051631101335b60e11b81526000600482015260406024820152602060448201527f4141313420696e6974436f6465206d7573742072657475726e2073656e6465726064820152fd5b6084604051631101335b60e11b81526000600482015260406024820152601b60448201527f4141313320696e6974436f6465206661696c6564206f72204f4f4700000000006064820152fd5b611103915060203d602011611109575b6110fb8183612faa565b8101906133c6565b8c610f8a565b503d6110f1565b6040513d6000823e3d90fd5b6084604051631101335b60e11b81526000600482015260406024820152601f60448201527f414131302073656e64657220616c726561647920636f6e7374727563746564006064820152fd5b60405162461bcd60e51b815260206004820152601860248201527f41413934206761732076616c756573206f766572666c6f7700000000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601d60248201527f4141393320696e76616c6964207061796d6173746572416e64446174610000006044820152606490fd5b5050600060e086015260006080860152600060a08601526107f6565b3461018b5760208060031936011261018b57611228613033565b33600052600082526001604060002001908154916001600160701b038360081c169283156113a35765ffffffffffff8160981c16801561135e57421061131957610100600160c81b0319169055604080516001600160a01b03831681526020810184905260009384938493849333917fb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda391a26001600160a01b03165af16112cd613190565b50156112d557005b6064906040519062461bcd60e51b82526004820152601860248201527f6661696c656420746f207769746864726177207374616b6500000000000000006044820152fd5b60405162461bcd60e51b815260048101869052601b60248201527f5374616b65207769746864726177616c206973206e6f742064756500000000006044820152606490fd5b60405162461bcd60e51b815260048101879052601d60248201527f6d7573742063616c6c20756e6c6f636b5374616b6528292066697273740000006044820152606490fd5b60405162461bcd60e51b81526004810186905260146024820152734e6f207374616b6520746f20776974686472617760601b6044820152606490fd5b3461018b57600036600319011261018b573360005260006020526001604060002001805463ffffffff8160781c169081156114d45760ff161561149b5765ffffffffffff90814216019181831161148557805460ff65ffffffffffff60981b01191665ffffffffffff60981b609885901b161790556040519116815233907ffa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a90602090a2005b634e487b7160e01b600052601160045260246000fd5b60405162461bcd60e51b8152602060048201526011602482015270616c726561647920756e7374616b696e6760781b6044820152606490fd5b60405162461bcd60e51b815260206004820152600a6024820152691b9bdd081cdd185ad95960b21b6044820152606490fd5b602036600319011261018b5761002261151d613033565b613406565b3461018b57602036600319011261018b576004356001600160401b03811161018b576020611557611590923690600401613049565b600654604051632b870d1b60e11b815260048101859052946001600160a01b0394938693928616928492600092849260248401916133e5565b03925af1908115611110576024926000926115bd575b50604051633653dc0360e11b815291166004820152fd5b6115d791925060203d602011611109576110fb8183612faa565b90836115a6565b3461018b5760031960603682011261018b576001600160401b036004351161018b57610120906004353603011261018b5761161761301d565b6044356001600160401b03811161018b57611636903690600401613049565b606060a060405161164681612f46565b60008152600060208201526000604082015260008382015260006080820152015261166f6136cd565b6116776132b8565b9161168660043560040161439b565b6000915a8451949093906001600160a01b036116a560048035016134e6565b16865260048035602481013560208901526001600160801b03608482013581811660608b0152608090811c60408b015260a483013560c08b015260c48301359182166101008b01521c6101208901526117039160e48201910161345d565b8015611fcf57603481106111ad578060141161018b578060241161018b5760341161018b57602481013560801c60a0880152601481013560801c60808801523560601c60e08701525b61175a6004356004016131c0565b60208201526040860151946001600160781b038660c08901511760608901511760808901511760a08901511761010089015117610120890151171161116857604087015160608801510160808801510160a08801510160c0880151016101008801510282519660018060a01b03885116976117df60446004350160043560040161345d565b80611eea575b505060e001516001600160a01b03169760008915611eb2575b6020868101516040516306608bdf60e21b8152928391829161182891600480358101908501614656565b038160008688f160009181611e7e575b506118af5761094a3d61080080821161094e5750604051906020818301016040528082526000602083013e6040516365c8fd4d60e01b81526000600482015260606024820152600d60648201526c10504c8cc81c995d995c9d1959609a1b608482015260a0604482015291829160a4830190613142565b9815611e5d575b5060018060a01b0389511660208a015160401c90600052600160205260406000209060005260205260406000206118ed8154613174565b90555a830311610e185760e090970151606097906001600160a01b0316611cce575b604083015286606083015260a46004350135905a90030160808201525a95604051968761194660646004350160043560040161345d565b919060009260038111611cc5575b638dd7712f60e01b936001600160e01b0319168403611c615750505060006119e4611a10602093848801519060405191829187830152604060248301526119a36064830160043560040161352b565b90604483015203906119bd601f1992838101835282612faa565b611a046040519485926242dc5360e01b898501526102006024850152610224840190613142565b6119f1604484018c613f36565b8281036023190161020484015289613142565b03908101835282612faa565b828151910182305af16000519860405215611af8575b509596949550611ac294606094600094906001600160a01b038316611ac6575b505050608001519460405195611a5b87612f46565b865260208601968752604086019081526060860191825260808601921515835260a0860193845260016002556040519687966020885251602088015251604087015251606086015251608085015251151560a08401525160c08084015260e0830190613142565b0390f35b6000949650849395508390826040519384928337810182815203925af1916080611aee613190565b9392908880611a46565b909596506000953d602014611c50575b63deaddead60e01b8703611b54576084604051631101335b60e11b81526000600482015260406024820152600f60448201526e41413935206f7574206f662067617360881b6064820152fd5b611ac29663deadaa5160e01b03611ba35750611b74611b7f915a90613183565b608083015190613167565b611b98604083015191611b9184614614565b82846145b3565b965b96959488611a26565b611c3b611c30611c429360208601518651907ff62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f4792602060018060a01b038451169301513d90610800808311611c48575b50604051916020818401016040528083526000602084013e611c266040519283928352604060208401526040830190613142565b0390a35a90613183565b608085015190613167565b90836135f0565b96611b9a565b915038611bf2565b955060206000803e60005195611b08565b6242dc5360e01b60208401526102006024840152611cbb9350611cad91611c8d916102248501916133e5565b611c9a6044840188613f36565b8281036023190161020484015285613142565b03601f1981018a5289612faa565b6020600089611a10565b81359350611954565b965091505a815160e08101516001600160a01b031660008181526020819052604090208054959392908a8710610dcc5760806000928c9283611d389a039055015192602087015183604051809a819582946314add44b60e21b845260043560040160048501614656565b039286f19485600091600097611dd1575b50611dc05761094a3d610800808211610ccb5750604051906020818301016040528082526000602083013e6040516365c8fd4d60e01b81526000600482015260606024820152600d60648201526c10504cccc81c995d995c9d1959609a1b608482015260a0604482015291829160a4830190613142565b94915a900311610ce457929661190f565b915095503d90816000823e611de68282612faa565b604081838101031261018b578051906001600160401b03821161018b57828101601f83830101121561018b578181015191611e2083612fcb565b93611e2e6040519586612faa565b838552820160208483850101011161018b57602092611e5491848087019185010161311f565b0151958c611d49565b600052600060205260406000208054808411610e855783900390558a6118b6565b9091506020813d602011611eaa575b81611e9a60209383612faa565b8101031261018b5751908d611838565b3d9150611e8d565b50806000526000602052604060002054838111600014611edd5750611828602060005b9150506117fe565b6020611828918503611ed5565b8a3b61111c576000602060018060a01b036006541660408a51015160405180948193632b870d1b60e11b835285600484015282611f2b602482018a8c6133e5565b0393f190811561111057600091611fb0575b506001600160a01b0381168015611095578c03611049573b15610ffd5760141161018b5789907fd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d604060208901519260018060a01b0360e08b510151168251913560601c82526020820152a38b806117e5565b611fc9915060203d602011611109576110fb8183612faa565b8e611f3d565b5050600060e087015260006080870152600060a087015261174c565b3461018b57606036600319011261018b576001600160401b0360043581811161018b5761201c903690600401613049565b91905061202761301d565b9060443590811161018b57612040903690600401613049565b91909215908161211b575b506120d6576014811015612079575b60405162461bcd60e51b81526020600482015260006024820152604490fd5b60141161018b573560601c3b1561209157808061205a565b60405162461bcd60e51b815260206004820152601b60248201527f41413330207061796d6173746572206e6f74206465706c6f79656400000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601960248201527f41413230206163636f756e74206e6f74206465706c6f796564000000000000006044820152606490fd5b90503b158361204b565b3461018b57604036600319011261018b5761213e613033565b6024356001600160401b03811161018b5760009161216183923690600401613049565b90816040519283928337810184815203915af461217c613190565b60408051632650415560e21b815292151560048401526024830152819061094a906044830190613142565b3461018b576121b5366130be565b6121c09291926136cd565b6121c983613329565b60005b84811061224157506000927fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972600080a16000915b8583106122115761022f858561427b565b9091936001906122376122258789876133a3565b61222f8886613379565b519088613fce565b0194019190612200565b61226c61226561225383859795613379565b5161225f8489876133a3565b846136ee565b9190614678565b6001600160a01b039291831661237d576123375761228990614678565b91166122f45761229e576001019290926121cc565b60a49060405190631101335b60e11b8252600482015260406024820152602160448201527f41413332207061796d61737465722065787069726564206f72206e6f742064756064820152606560f81b6084820152fd5b60848260405190631101335b60e11b8252600482015260406024820152601460448201527320a0999a1039b4b3b730ba3ab9329032b93937b960611b6064820152fd5b60848360405190631101335b60e11b82526004820152604060248201526017604482015276414132322065787069726564206f72206e6f742064756560481b6064820152fd5b60848460405190631101335b60e11b8252600482015260406024820152601460448201527320a0991a1039b4b3b730ba3ab9329032b93937b960611b6064820152fd5b3461018b57602036600319011261018b576001600160a01b036123e1613033565b1660005260006020526020604060002054604051908152f35b3461018b57602036600319011261018b57612413613033565b6000608060405161242381612ef9565b828152826020820152826040820152826060820152015260018060a01b0316600052600060205260a06040600020608060405161245f81612ef9565b6001835493848352015490602081019060ff8316151582526001600160701b0360408201818560081c16815263ffffffff936060840193858760781c16855265ffffffffffff978891019660981c1686526040519788525115156020880152511660408601525116606084015251166080820152f35b3461018b57604036600319011261018b5760206124f0613033565b6124f8613076565b6001600160a01b0390911660009081526001835260408082206001600160c01b03841683528452908190205481519290911b67ffffffffffffffff1916178152f35b3461018b57602061255261254d3661308c565b6131c0565b604051908152f35b3461018b57604036600319011261018b57612573613033565b602435903360005260006020526040600020805480841161263357600084819482946125a0838596613183565b9055604080516001600160a01b03831681526020810184905233917fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb91a26001600160a01b03165af16125f1613190565b50156125f957005b60405162461bcd60e51b81526020600482015260126024820152716661696c656420746f20776974686472617760701b6044820152606490fd5b60405162461bcd60e51b815260206004820152601960248201527f576974686472617720616d6f756e7420746f6f206c61726765000000000000006044820152606490fd5b3461018b57604036600319011261018b57612691613033565b612699613076565b9060018060a01b0316600052600160205260406000209060018060c01b03166000526020526020604060002054604051908152f35b3461018b57602036600319011261018b576004356001600160c01b0381169081900361018b5733600052600160205260406000209060005260205260406000206127188154613174565b9055005b60208060031936011261018b5760043563ffffffff9182821680920361018b573360005260008152604060002092821561292c576001840154908160781c1683106128e7576127796001600160701b039182349160081c16613167565b9384156128ad57818511612877579065ffffffffffff6128469254604051906127a182612ef9565b81528481019260018452604082019088168152606082018781526001608084019360008552336000526000895260406000209051815501945115159160ff6effffffffffffffffffffffffffff008754925160081b169263ffffffff60781b905160781b169316906cffffffffffffffffffffffffff60981b161717178355511681549065ffffffffffff60981b9060981b169065ffffffffffff60981b1916179055565b6040519283528201527fa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c0160403392a2005b60405162461bcd60e51b815260048101849052600e60248201526d7374616b65206f766572666c6f7760901b6044820152606490fd5b60405162461bcd60e51b81526004810184905260126024820152711b9bc81cdd185ad9481cdc1958da599a595960721b6044820152606490fd5b60405162461bcd60e51b815260048101839052601c60248201527f63616e6e6f7420646563726561736520756e7374616b652074696d65000000006044820152606490fd5b60405162461bcd60e51b815260048101839052601a60248201527f6d757374207370656369667920756e7374616b652064656c61790000000000006044820152606490fd5b3461018b57602036600319011261018b5760043563ffffffff60e01b811680910361018b5760209063307e35b760e11b81149081156129fe575b81156129ed575b81156129dc575b81156129cb575b506040519015158152f35b6301ffc9a760e01b149050826129c0565b633e84f02160e01b811491506129b9565b63cf28ef9760e01b811491506129b2565b63122a0e9b60e31b811491506129ab565b3461018b5761020036600319011261018b576001600160401b0360043581811161018b573660238201121561018b57612a52903690602481600401359101612fe6565b903660231901906101c0821261018b5761014060405192612a7284612ef9565b1261018b57604051612a8381612f2a565b612a8b61301d565b815260443560208201526064356040820152608435606082015260a435608082015260c43560a082015260e43560c0820152610104356001600160a01b038116810361018b5760e0820152610124356101008201526101443561012082015282526101643560208301526101843560408301526101a43560608301526101c43560808301526101e43590811161018b57612b29903690600401613049565b91905a303303612eb4578251606081015195603f5a0260061c61271060a084015189010111612ea35760009681519182612ddb575b5050505090612b78915a9003608084015101933691612fe6565b915a600093835192612b8984614563565b60e085015190926001600160a01b039091169081612cdb57505083516001600160a01b0316925b5a9003019260a06060820151910151016080850151840390818111612cc7575b50508202604084015191818310600014612c325750506003851015612c1e57602094600203612c0f576125529293508093612c0a81614614565b6145b3565b505063deadaa5160e01b825250fd5b634e487b7160e01b84526021600452602484fd5b81612c429296979593039061458d565b506003831015612cb357602094507f49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f60808683015192519460018060a01b03865116948860018060a01b0360e0890151169701519160405192835215898301528760408301526060820152a4612552565b634e487b7160e01b85526021600452602485fd5b6064919003600a0204909201918680612bd0565b819491948151612ced575b5050612bb0565b60038a1015612dc55760028a0315612ce65760a087015191813b1561018b57612d3e928b60008094604051809781968295637c627b2160e01b84526004840152608060248401526084830190613142565b8b8b0260448301528b60648301520393f19081612db1575b50612daa5761094a873d610800808211612da2575b5060405191602082840101604052818352602083013e604051632b5e552f60e21b8152602060048201529182916024830190613142565b905083612d6b565b8880612ce6565b612dbc919850612f7c565b60009689612d56565b634e487b7160e01b600052602160045260246000fd5b9160009291838093602060018060a01b03885116910192f115612e01575b808080612b5e565b612b78929195503d610800808211612e9b575b50604051906020818301016040528082526000602083013e8051612e3f575b50506001949091612df9565b7f1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201602086015191602060018060a01b03855116940151612e916040519283928352604060208401526040830190613142565b0390a38580612e33565b905087612e14565b63deaddead60e01b60005260206000fd5b60405162461bcd60e51b815260206004820152601760248201527f4141393220696e7465726e616c2063616c6c206f6e6c790000000000000000006044820152606490fd5b60a081019081106001600160401b03821117612f1457604052565b634e487b7160e01b600052604160045260246000fd5b61014081019081106001600160401b03821117612f1457604052565b60c081019081106001600160401b03821117612f1457604052565b604081019081106001600160401b03821117612f1457604052565b6001600160401b038111612f1457604052565b606081019081106001600160401b03821117612f1457604052565b90601f801991011681019081106001600160401b03821117612f1457604052565b6001600160401b038111612f1457601f01601f191660200190565b929192612ff282612fcb565b916130006040519384612faa565b82948184528183011161018b578281602093846000960137010152565b602435906001600160a01b038216820361018b57565b600435906001600160a01b038216820361018b57565b9181601f8401121561018b578235916001600160401b03831161018b576020838186019501011161018b57565b602435906001600160c01b038216820361018b57565b6003199060208183011261018b57600435916001600160401b03831161018b57826101209203011261018b5760040190565b90604060031983011261018b576004356001600160401b039283821161018b578060238301121561018b57816004013593841161018b5760248460051b8301011161018b576024908101929190356001600160a01b038116810361018b5790565b60005b8381106131325750506000910152565b8181015183820152602001613122565b9060209161315b8151809281855285808601910161311f565b601f01601f1916010190565b9190820180921161148557565b60001981146114855760010190565b9190820391821161148557565b3d156131bb573d906131a182612fcb565b916131af6040519384612faa565b82523d6000602084013e565b606090565b60406131ce8183018361345d565b90818351918237206131e3606084018461345d565b90818451918237209260c06131fb60e083018361345d565b908186519182372091845195602087019460018060a01b03833516865260208301358789015260608801526080870152608081013560a087015260a081013582870152013560e0850152610100908185015283526101208301916001600160401b039184841083851117612f1457838252845190206101408501908152306101608601524661018086015260608452936101a00191821183831017612f14575251902090565b6001600160401b038111612f145760051b60200190565b604051906132c582612ef9565b6040516080836132d483612f2a565b60009283815283602082015283604082015283606082015283838201528360a08201528360c08201528360e0820152836101008201528361012082015281528260208201528260408201528260608201520152565b90613333826132a1565b6133406040519182612faa565b8281528092613351601f19916132a1565b019060005b82811061336257505050565b60209061336d6132b8565b82828501015201613356565b805182101561338d5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b919081101561338d5760051b8101359061011e198136030182121561018b570190565b9081602091031261018b57516001600160a01b038116810361018b5790565b908060209392818452848401376000828201840152601f01601f1916010190565b6001805b600581106134565750507f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c46020613441348461458d565b6040519081526001600160a01b0390931692a2565b810161340a565b903590601e198136030182121561018b57018035906001600160401b03821161018b5760200191813603831361018b57565b919081101561338d5760051b81013590605e198136030182121561018b570190565b903590601e198136030182121561018b57018035906001600160401b03821161018b57602001918160051b3603831361018b57565b356001600160a01b038116810361018b5790565b9035601e198236030181121561018b5701602081359101916001600160401b03821161018b57813603831361018b57565b6101209181356001600160a01b0381169081900361018b5761359561357a6135ed956135ce9385526020860135602086015261356a60408701876134fa565b90918060408801528601916133e5565b61358760608601866134fa565b9085830360608701526133e5565b6080840135608084015260a084013560a084015260c084013560c08401526135c060e08501856134fa565b9084830360e08601526133e5565b916135df61010091828101906134fa565b9290918185039101526133e5565b90565b9291905a9060009085519361360485614563565b60e086015190926001600160a01b03918216806136b9575050855116935b5a9003019360a060608201519101510160808701518503908181116136a5575b50508302604086015192818410600014613684575050806136705750908161366e9294612c0a81614614565b565b634e487b7160e01b81526021600452602490fd5b613694908284939895039061458d565b506136705750908361366e926145b3565b6064919003600a0204909301923880613642565b909591509451156136225760009350613622565b60028054146136dc5760028055565b604051633ee5aeb560e01b8152600490fd5b926000905a845194906001600160a01b03613708866134e6565b1686526020850135602087015260808501356001600160801b038116606088015260801c604087015260a085013560c087015260c08501356001600160801b03811661010088015260801c61012087015261376660e086018661345d565b8015613f1a57603481106111ad578060141161018b578060241161018b5760341161018b57602481013560801c60a0880152601481013560801c60808801523560601c60e08701525b6137b8856131c0565b60208301526040860151946001600160781b038660c08901511760608901511760808901511760a08901511761010089015117610120890151171161116857604087015160608801510160808801510160a08801510160c0880151016101008801510296835160018060a01b0381511690613836604085018561345d565b80613d05575b505060e001516001600160a01b03169060008215613ccd575b6020613881918b828a01516000868a604051978896879586936306608bdf60e21b855260048501614656565b0393f160009181613c99575b5061390c573d8c610800808311613904575b50604051916020818401016040528083526000602084013e61094a6040519283926365c8fd4d60e01b8452600484015260606024840152600d60648401526c10504c8cc81c995d995c9d1959609a1b608484015260a0604484015260a4830190613142565b91508261389f565b9a92939495969798999a9115613c32575b509760018060a01b03835116602084015160401c90600052600160205260406000209060005260205260406000206139558154613174565b90555a850311613be65760e090910151606091906001600160a01b0316613999575b509060a09184959697986040608096015260608601520135905a900301910152565b969550505a9683519760018060a01b0360e08a01511680600052600060205260406000208054848110613b9a5760806139fd9a9b9c600093878094039055015192602089015183604051809d819582946314add44b60e21b84528c60048501614656565b039286f1978860009160009a613b0e575b50613a8d573d8b610800808311613a85575b50604051916020818401016040528083526000602084013e61094a6040519283926365c8fd4d60e01b8452600484015260606024840152600d60648401526c10504cccc81c995d995c9d1959609a1b608484015260a0604484015260a4830190613142565b915082613a20565b9991929394959697989998925a900311613ab257509096959094939291906080613977565b60a49060405190631101335b60e11b8252600482015260406024820152602760448201527f41413336206f766572207061796d6173746572566572696669636174696f6e47606482015266185cd31a5b5a5d60ca1b6084820152fd5b915098503d90816000823e613b238282612faa565b604081838101031261018b578051906001600160401b03821161018b57828101601f83830101121561018b578181015191613b5d83612fcb565b93613b6b6040519586612faa565b838552820160208483850101011161018b57602092613b9191848087019185010161311f565b01519838613a0e565b60848b60405190631101335b60e11b8252600482015260406024820152601e60448201527f41413331207061796d6173746572206465706f73697420746f6f206c6f7700006064820152fd5b60849060405190631101335b60e11b8252600482015260406024820152601e60448201527f41413236206f76657220766572696669636174696f6e4761734c696d697400006064820152fd5b600052600060205260406000208054808c11613c53578b900390553861391d565b60848460405190631101335b60e11b8252600482015260406024820152601760448201527610504c8c48191a591b89dd081c185e481c1c99599d5b99604a1b6064820152fd5b9091506020813d602011613cc5575b81613cb560209383612faa565b8101031261018b5751903861388d565b3d9150613ca8565b508060005260006020526040600020548a8111600014613cf85750613881602060005b915050613855565b6020613881918c03613cf0565b833b613ece576000602060018060a01b036006541660408b51015160405180948193632b870d1b60e11b835285600484015282613d46602482018a8c6133e5565b0393f190811561111057600091613eaf575b506001600160a01b0381168015613e63578503613e17573b15613dcb5760141161018b5782907fd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d604060208a01519260018060a01b0360e08c510151168251913560601c82526020820152a3388061383c565b60848d60405190631101335b60e11b8252600482015260406024820152602060448201527f4141313520696e6974436f6465206d757374206372656174652073656e6465726064820152fd5b60848e60405190631101335b60e11b8252600482015260406024820152602060448201527f4141313420696e6974436f6465206d7573742072657475726e2073656e6465726064820152fd5b60848f60405190631101335b60e11b8252600482015260406024820152601b60448201527f4141313320696e6974436f6465206661696c6564206f72204f4f4700000000006064820152fd5b613ec8915060203d602011611109576110fb8183612faa565b38613d58565b60848d60405190631101335b60e11b8252600482015260406024820152601f60448201527f414131302073656e64657220616c726561647920636f6e7374727563746564006064820152fd5b5050600060e087015260006080870152600060a08701526137af565b60806101a091805160018060a01b03808251168652602082015160208701526040820151604087015260608201516060870152838201518487015260a082015160a087015260c082015160c087015260e08201511660e0860152610100808201519086015261012080910151908501526020810151610140850152604081015161016085015260608101516101808501520151910152565b9092915a60608201519160409283519687613fec606083018361345d565b919060009260038111614272575b638dd7712f60e01b936001600160e01b031916840361420c575050506140a36000926140839260208701516140428a5193849360208501528b6024850152606484019061352b565b906044830152039061405c601f1992838101835282612faa565b611a0489519485926242dc5360e01b60208501526102006024850152610224840190613142565b614090604484018a613f36565b8281036023190161020484015287613142565b6020918183809351910182305af1600051988652156140c5575b505050505050565b909192939495965060003d8214614202575b63deaddead60e01b81036141215760848787805191631101335b60e11b835260048301526024820152600f60448201526e41413935206f7574206f662067617360881b6064820152fd5b93955091939290919063deadaa5160e01b0361416757505061414b611c3061415a93945a90613183565b9083015183612c0a8295614614565b905b3880808080806140bd565b6141f4946141ee9293827ff62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f47926141e39488015191885193611c268260018060a01b03875116960151913d906108008083116141fa575b508051918581840101825280835260008684013e8080519586958652850152830190613142565b608084015190613167565b916135f0565b9061415c565b9150386141bc565b8181803e516140d7565b6242dc5360e01b6020840152610200602484015261426a945091925061425c9161423c91610224850191906133e5565b6142496044840187613f36565b8281036023190161020484015284613142565b03601f198101895288612faa565b6000876140a3565b81359350613ffa565b6001600160a01b031680156142e857600080809381935af161429b613190565b50156142a357565b60405162461bcd60e51b815260206004820152601f60248201527f41413931206661696c65642073656e6420746f2062656e6566696369617279006044820152606490fd5b60405162461bcd60e51b815260206004820152601860248201527f4141393020696e76616c69642062656e656669636961727900000000000000006044820152606490fd5b600060443d106135ed57604051600319913d83016004833e81516001600160401b03918282113d60248401111761438a57818401948551938411614392573d8501016020848701011161438a57506135ed92910160200190612faa565b949350505050565b50949350505050565b604080516135a560f21b602082019081523060601b6022830152600160f81b60368301526017825260009390916143d181612f61565b51909120600680546001600160a01b0319166001600160a01b03928316179055906143fe8184018261345d565b91909261441861440d836134e6565b9260e081019061345d565b91303b1561450f57928794926144679261444f97958951988997889763957122ab60e01b8952606060048a015260648901916133e5565b931660248601528483036003190160448601526133e5565b0381305afa90816144fc575b506144f85760018260033d116144e8575b6308c379a0146144a1575b614497575050565b51903d90823e3d90fd5b6144a961432d565b806144b5575b5061448f565b8051849250156144af578361094a84928351938493631101335b60e11b8552600485015260248401526044830190613142565b50600483803e825160e01c614484565b5050565b61450890939193612f7c565b9138614473565b8780fd5b9060405161452081612f61565b600080825260208083018281526001600160a01b0390951682528190526040902060010154600881901c6001600160701b0316825260781c63ffffffff16909252565b61012061010082015191015180821461458957480180821015614584575090565b905090565b5090565b60018060a01b031660005260006020526145ad6040600020918254613167565b80915590565b9190917f49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f6080602083015192519460018060a01b03946020868851169660e089015116970151916040519283526000602084015260408301526060820152a4565b60208101519051907f67b4fa9642f42120bf031f3051d1824b0fe25627945b27b8a6a65d5761d5482e60208060018060a01b03855116940151604051908152a3565b61466e6040929594939560608352606083019061352b565b9460208201520152565b80156146fb5760006040805161468d81612f8f565b8281526020810183905201526001600160a01b0381169065ffffffffffff9060409060a081901c83169081156146f3575b60d01c928251916146ce83612f8f565b85835284602084015216918291015242119081156146eb57509091565b905042109091565b8391506146be565b5060009060009056fea26469706673582212206d4fee6a58f19c465518402c6dc3e9fbdae7ad7f2db345a21f9340a4ca4b7a1f64736f6c6343000819003360808060405234601557610144908161001b8239f35b600080fdfe6080600436101561000f57600080fd5b6000803560e01c63570e1a361461002557600080fd5b3461010b57602036600319011261010b576004359167ffffffffffffffff9081841161010757366023850112156101075783600401358281116101035736602482870101116101035780601411610103576013198101928084116100ef57600b8201601f19908116603f01168301908111838210176100ef5792846024819482600c60209a968b9960405286845289840196603889018837830101525193013560601c5af190805191156100e7575b506040516001600160a01b039091168152f35b9050386100d4565b634e487b7160e01b85526041600452602485fd5b8380fd5b8280fd5b80fdfea26469706673582212206a1b9f618c552b5e5a558cda5cc2bafd851f85b1d0662028a1684dd5741f8d9464736f6c63430008190033",
}

// EntryPointSimulationsABI is the input ABI used to generate the binding from.
// Deprecated: Use EntryPointSimulationsMetaData.ABI instead.
var EntryPointSimulationsABI = EntryPointSimulationsMetaData.ABI

// EntryPointSimulationsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EntryPointSimulationsMetaData.Bin instead.
var EntryPointSimulationsBin = EntryPointSimulationsMetaData.Bin

// DeployEntryPointSimulations deploys a new Ethereum contract, binding an instance of EntryPointSimulations to it.
func DeployEntryPointSimulations(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EntryPointSimulations, error) {
	parsed, err := EntryPointSimulationsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EntryPointSimulationsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EntryPointSimulations{EntryPointSimulationsCaller: EntryPointSimulationsCaller{contract: contract}, EntryPointSimulationsTransactor: EntryPointSimulationsTransactor{contract: contract}, EntryPointSimulationsFilterer: EntryPointSimulationsFilterer{contract: contract}}, nil
}

// EntryPointSimulations is an auto generated Go binding around an Ethereum contract.
type EntryPointSimulations struct {
	EntryPointSimulationsCaller     // Read-only binding to the contract
	EntryPointSimulationsTransactor // Write-only binding to the contract
	EntryPointSimulationsFilterer   // Log filterer for contract events
}

// EntryPointSimulationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EntryPointSimulationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntryPointSimulationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EntryPointSimulationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntryPointSimulationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EntryPointSimulationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntryPointSimulationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EntryPointSimulationsSession struct {
	Contract     *EntryPointSimulations // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EntryPointSimulationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EntryPointSimulationsCallerSession struct {
	Contract *EntryPointSimulationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// EntryPointSimulationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EntryPointSimulationsTransactorSession struct {
	Contract     *EntryPointSimulationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// EntryPointSimulationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EntryPointSimulationsRaw struct {
	Contract *EntryPointSimulations // Generic contract binding to access the raw methods on
}

// EntryPointSimulationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EntryPointSimulationsCallerRaw struct {
	Contract *EntryPointSimulationsCaller // Generic read-only contract binding to access the raw methods on
}

// EntryPointSimulationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EntryPointSimulationsTransactorRaw struct {
	Contract *EntryPointSimulationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEntryPointSimulations creates a new instance of EntryPointSimulations, bound to a specific deployed contract.
func NewEntryPointSimulations(address common.Address, backend bind.ContractBackend) (*EntryPointSimulations, error) {
	contract, err := bindEntryPointSimulations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulations{EntryPointSimulationsCaller: EntryPointSimulationsCaller{contract: contract}, EntryPointSimulationsTransactor: EntryPointSimulationsTransactor{contract: contract}, EntryPointSimulationsFilterer: EntryPointSimulationsFilterer{contract: contract}}, nil
}

// NewEntryPointSimulationsCaller creates a new read-only instance of EntryPointSimulations, bound to a specific deployed contract.
func NewEntryPointSimulationsCaller(address common.Address, caller bind.ContractCaller) (*EntryPointSimulationsCaller, error) {
	contract, err := bindEntryPointSimulations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsCaller{contract: contract}, nil
}

// NewEntryPointSimulationsTransactor creates a new write-only instance of EntryPointSimulations, bound to a specific deployed contract.
func NewEntryPointSimulationsTransactor(address common.Address, transactor bind.ContractTransactor) (*EntryPointSimulationsTransactor, error) {
	contract, err := bindEntryPointSimulations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsTransactor{contract: contract}, nil
}

// NewEntryPointSimulationsFilterer creates a new log filterer instance of EntryPointSimulations, bound to a specific deployed contract.
func NewEntryPointSimulationsFilterer(address common.Address, filterer bind.ContractFilterer) (*EntryPointSimulationsFilterer, error) {
	contract, err := bindEntryPointSimulations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsFilterer{contract: contract}, nil
}

// bindEntryPointSimulations binds a generic wrapper to an already deployed contract.
func bindEntryPointSimulations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EntryPointSimulationsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntryPointSimulations *EntryPointSimulationsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EntryPointSimulations.Contract.EntryPointSimulationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntryPointSimulations *EntryPointSimulationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.EntryPointSimulationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntryPointSimulations *EntryPointSimulationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.EntryPointSimulationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntryPointSimulations *EntryPointSimulationsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EntryPointSimulations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntryPointSimulations *EntryPointSimulationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntryPointSimulations *EntryPointSimulationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.contract.Transact(opts, method, params...)
}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_EntryPointSimulations *EntryPointSimulationsCaller) ValidateSenderAndPaymaster(opts *bind.CallOpts, initCode []byte, sender common.Address, paymasterAndData []byte) error {
	var out []interface{}
	err := _EntryPointSimulations.contract.Call(opts, &out, "_validateSenderAndPaymaster", initCode, sender, paymasterAndData)

	if err != nil {
		return err
	}

	return err

}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_EntryPointSimulations *EntryPointSimulationsSession) ValidateSenderAndPaymaster(initCode []byte, sender common.Address, paymasterAndData []byte) error {
	return _EntryPointSimulations.Contract.ValidateSenderAndPaymaster(&_EntryPointSimulations.CallOpts, initCode, sender, paymasterAndData)
}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_EntryPointSimulations *EntryPointSimulationsCallerSession) ValidateSenderAndPaymaster(initCode []byte, sender common.Address, paymasterAndData []byte) error {
	return _EntryPointSimulations.Contract.ValidateSenderAndPaymaster(&_EntryPointSimulations.CallOpts, initCode, sender, paymasterAndData)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EntryPointSimulations *EntryPointSimulationsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EntryPointSimulations.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EntryPointSimulations *EntryPointSimulationsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EntryPointSimulations.Contract.BalanceOf(&_EntryPointSimulations.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EntryPointSimulations *EntryPointSimulationsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EntryPointSimulations.Contract.BalanceOf(&_EntryPointSimulations.CallOpts, account)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_EntryPointSimulations *EntryPointSimulationsCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	var out []interface{}
	err := _EntryPointSimulations.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		Deposit         *big.Int
		Staked          bool
		Stake           *big.Int
		UnstakeDelaySec uint32
		WithdrawTime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Staked = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Stake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakeDelaySec = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.WithdrawTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_EntryPointSimulations *EntryPointSimulationsSession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _EntryPointSimulations.Contract.Deposits(&_EntryPointSimulations.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_EntryPointSimulations *EntryPointSimulationsCallerSession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _EntryPointSimulations.Contract.Deposits(&_EntryPointSimulations.CallOpts, arg0)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_EntryPointSimulations *EntryPointSimulationsCaller) GetDepositInfo(opts *bind.CallOpts, account common.Address) (IStakeManagerDepositInfo, error) {
	var out []interface{}
	err := _EntryPointSimulations.contract.Call(opts, &out, "getDepositInfo", account)

	if err != nil {
		return *new(IStakeManagerDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeManagerDepositInfo)).(*IStakeManagerDepositInfo)

	return out0, err

}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_EntryPointSimulations *EntryPointSimulationsSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _EntryPointSimulations.Contract.GetDepositInfo(&_EntryPointSimulations.CallOpts, account)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_EntryPointSimulations *EntryPointSimulationsCallerSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _EntryPointSimulations.Contract.GetDepositInfo(&_EntryPointSimulations.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_EntryPointSimulations *EntryPointSimulationsCaller) GetNonce(opts *bind.CallOpts, sender common.Address, key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EntryPointSimulations.contract.Call(opts, &out, "getNonce", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_EntryPointSimulations *EntryPointSimulationsSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _EntryPointSimulations.Contract.GetNonce(&_EntryPointSimulations.CallOpts, sender, key)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_EntryPointSimulations *EntryPointSimulationsCallerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _EntryPointSimulations.Contract.GetNonce(&_EntryPointSimulations.CallOpts, sender, key)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_EntryPointSimulations *EntryPointSimulationsCaller) GetUserOpHash(opts *bind.CallOpts, userOp PackedUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _EntryPointSimulations.contract.Call(opts, &out, "getUserOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHash is a free data retrieval call binding the contract method 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_EntryPointSimulations *EntryPointSimulationsSession) GetUserOpHash(userOp PackedUserOperation) ([32]byte, error) {
	return _EntryPointSimulations.Contract.GetUserOpHash(&_EntryPointSimulations.CallOpts, userOp)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_EntryPointSimulations *EntryPointSimulationsCallerSession) GetUserOpHash(userOp PackedUserOperation) ([32]byte, error) {
	return _EntryPointSimulations.Contract.GetUserOpHash(&_EntryPointSimulations.CallOpts, userOp)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_EntryPointSimulations *EntryPointSimulationsCaller) NonceSequenceNumber(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EntryPointSimulations.contract.Call(opts, &out, "nonceSequenceNumber", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_EntryPointSimulations *EntryPointSimulationsSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _EntryPointSimulations.Contract.NonceSequenceNumber(&_EntryPointSimulations.CallOpts, arg0, arg1)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_EntryPointSimulations *EntryPointSimulationsCallerSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _EntryPointSimulations.Contract.NonceSequenceNumber(&_EntryPointSimulations.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EntryPointSimulations *EntryPointSimulationsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _EntryPointSimulations.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EntryPointSimulations *EntryPointSimulationsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EntryPointSimulations.Contract.SupportsInterface(&_EntryPointSimulations.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EntryPointSimulations *EntryPointSimulationsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EntryPointSimulations.Contract.SupportsInterface(&_EntryPointSimulations.CallOpts, interfaceId)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_EntryPointSimulations *EntryPointSimulationsSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.AddStake(&_EntryPointSimulations.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.AddStake(&_EntryPointSimulations.TransactOpts, unstakeDelaySec)
}

// DelegateAndRevert is a paid mutator transaction binding the contract method 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactor) DelegateAndRevert(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.Transact(opts, "delegateAndRevert", target, data)
}

// DelegateAndRevert is a paid mutator transaction binding the contract method 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (_EntryPointSimulations *EntryPointSimulationsSession) DelegateAndRevert(target common.Address, data []byte) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.DelegateAndRevert(&_EntryPointSimulations.TransactOpts, target, data)
}

// DelegateAndRevert is a paid mutator transaction binding the contract method 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) DelegateAndRevert(target common.Address, data []byte) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.DelegateAndRevert(&_EntryPointSimulations.TransactOpts, target, data)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactor) DepositTo(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.Transact(opts, "depositTo", account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_EntryPointSimulations *EntryPointSimulationsSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.DepositTo(&_EntryPointSimulations.TransactOpts, account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.DepositTo(&_EntryPointSimulations.TransactOpts, account)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactor) GetSenderAddress(opts *bind.TransactOpts, initCode []byte) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.Transact(opts, "getSenderAddress", initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_EntryPointSimulations *EntryPointSimulationsSession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.GetSenderAddress(&_EntryPointSimulations.TransactOpts, initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.GetSenderAddress(&_EntryPointSimulations.TransactOpts, initCode)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0xdbed18e0.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactor) HandleAggregatedOps(opts *bind.TransactOpts, opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.Transact(opts, "handleAggregatedOps", opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0xdbed18e0.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_EntryPointSimulations *EntryPointSimulationsSession) HandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.HandleAggregatedOps(&_EntryPointSimulations.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0xdbed18e0.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) HandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.HandleAggregatedOps(&_EntryPointSimulations.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x765e827f.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address beneficiary) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactor) HandleOps(opts *bind.TransactOpts, ops []PackedUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.Transact(opts, "handleOps", ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x765e827f.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address beneficiary) returns()
func (_EntryPointSimulations *EntryPointSimulationsSession) HandleOps(ops []PackedUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.HandleOps(&_EntryPointSimulations.TransactOpts, ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x765e827f.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address beneficiary) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) HandleOps(ops []PackedUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.HandleOps(&_EntryPointSimulations.TransactOpts, ops, beneficiary)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactor) IncrementNonce(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.Transact(opts, "incrementNonce", key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_EntryPointSimulations *EntryPointSimulationsSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.IncrementNonce(&_EntryPointSimulations.TransactOpts, key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.IncrementNonce(&_EntryPointSimulations.TransactOpts, key)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x0042dc53.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPointSimulations *EntryPointSimulationsTransactor) InnerHandleOp(opts *bind.TransactOpts, callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.Transact(opts, "innerHandleOp", callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x0042dc53.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPointSimulations *EntryPointSimulationsSession) InnerHandleOp(callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.InnerHandleOp(&_EntryPointSimulations.TransactOpts, callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x0042dc53.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) InnerHandleOp(callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.InnerHandleOp(&_EntryPointSimulations.TransactOpts, callData, opInfo, context)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0x97b2dcb9.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) op, address target, bytes targetCallData) returns((uint256,uint256,uint256,uint256,bool,bytes))
func (_EntryPointSimulations *EntryPointSimulationsTransactor) SimulateHandleOp(opts *bind.TransactOpts, op PackedUserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.Transact(opts, "simulateHandleOp", op, target, targetCallData)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0x97b2dcb9.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) op, address target, bytes targetCallData) returns((uint256,uint256,uint256,uint256,bool,bytes))
func (_EntryPointSimulations *EntryPointSimulationsSession) SimulateHandleOp(op PackedUserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.SimulateHandleOp(&_EntryPointSimulations.TransactOpts, op, target, targetCallData)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0x97b2dcb9.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) op, address target, bytes targetCallData) returns((uint256,uint256,uint256,uint256,bool,bytes))
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) SimulateHandleOp(op PackedUserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.SimulateHandleOp(&_EntryPointSimulations.TransactOpts, op, target, targetCallData)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xc3bce009.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) returns(((uint256,uint256,uint256,uint256,bytes),(uint256,uint256),(uint256,uint256),(uint256,uint256),(address,(uint256,uint256))))
func (_EntryPointSimulations *EntryPointSimulationsTransactor) SimulateValidation(opts *bind.TransactOpts, userOp PackedUserOperation) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.Transact(opts, "simulateValidation", userOp)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xc3bce009.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) returns(((uint256,uint256,uint256,uint256,bytes),(uint256,uint256),(uint256,uint256),(uint256,uint256),(address,(uint256,uint256))))
func (_EntryPointSimulations *EntryPointSimulationsSession) SimulateValidation(userOp PackedUserOperation) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.SimulateValidation(&_EntryPointSimulations.TransactOpts, userOp)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xc3bce009.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) returns(((uint256,uint256,uint256,uint256,bytes),(uint256,uint256),(uint256,uint256),(uint256,uint256),(address,(uint256,uint256))))
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) SimulateValidation(userOp PackedUserOperation) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.SimulateValidation(&_EntryPointSimulations.TransactOpts, userOp)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_EntryPointSimulations *EntryPointSimulationsSession) UnlockStake() (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.UnlockStake(&_EntryPointSimulations.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.UnlockStake(&_EntryPointSimulations.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_EntryPointSimulations *EntryPointSimulationsSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.WithdrawStake(&_EntryPointSimulations.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.WithdrawStake(&_EntryPointSimulations.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.Transact(opts, "withdrawTo", withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_EntryPointSimulations *EntryPointSimulationsSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.WithdrawTo(&_EntryPointSimulations.TransactOpts, withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.WithdrawTo(&_EntryPointSimulations.TransactOpts, withdrawAddress, withdrawAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntryPointSimulations.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EntryPointSimulations *EntryPointSimulationsSession) Receive() (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.Receive(&_EntryPointSimulations.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EntryPointSimulations *EntryPointSimulationsTransactorSession) Receive() (*types.Transaction, error) {
	return _EntryPointSimulations.Contract.Receive(&_EntryPointSimulations.TransactOpts)
}

// EntryPointSimulationsAccountDeployedIterator is returned from FilterAccountDeployed and is used to iterate over the raw logs and unpacked data for AccountDeployed events raised by the EntryPointSimulations contract.
type EntryPointSimulationsAccountDeployedIterator struct {
	Event *EntryPointSimulationsAccountDeployed // Event containing the contract specifics and raw log

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
func (it *EntryPointSimulationsAccountDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSimulationsAccountDeployed)
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
		it.Event = new(EntryPointSimulationsAccountDeployed)
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
func (it *EntryPointSimulationsAccountDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSimulationsAccountDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSimulationsAccountDeployed represents a AccountDeployed event raised by the EntryPointSimulations contract.
type EntryPointSimulationsAccountDeployed struct {
	UserOpHash [32]byte
	Sender     common.Address
	Factory    common.Address
	Paymaster  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountDeployed is a free log retrieval operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) FilterAccountDeployed(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*EntryPointSimulationsAccountDeployedIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.FilterLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsAccountDeployedIterator{contract: _EntryPointSimulations.contract, event: "AccountDeployed", logs: logs, sub: sub}, nil
}

// WatchAccountDeployed is a free log subscription operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) WatchAccountDeployed(opts *bind.WatchOpts, sink chan<- *EntryPointSimulationsAccountDeployed, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.WatchLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSimulationsAccountDeployed)
				if err := _EntryPointSimulations.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
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
func (_EntryPointSimulations *EntryPointSimulationsFilterer) ParseAccountDeployed(log types.Log) (*EntryPointSimulationsAccountDeployed, error) {
	event := new(EntryPointSimulationsAccountDeployed)
	if err := _EntryPointSimulations.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointSimulationsBeforeExecutionIterator is returned from FilterBeforeExecution and is used to iterate over the raw logs and unpacked data for BeforeExecution events raised by the EntryPointSimulations contract.
type EntryPointSimulationsBeforeExecutionIterator struct {
	Event *EntryPointSimulationsBeforeExecution // Event containing the contract specifics and raw log

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
func (it *EntryPointSimulationsBeforeExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSimulationsBeforeExecution)
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
		it.Event = new(EntryPointSimulationsBeforeExecution)
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
func (it *EntryPointSimulationsBeforeExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSimulationsBeforeExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSimulationsBeforeExecution represents a BeforeExecution event raised by the EntryPointSimulations contract.
type EntryPointSimulationsBeforeExecution struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBeforeExecution is a free log retrieval operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_EntryPointSimulations *EntryPointSimulationsFilterer) FilterBeforeExecution(opts *bind.FilterOpts) (*EntryPointSimulationsBeforeExecutionIterator, error) {

	logs, sub, err := _EntryPointSimulations.contract.FilterLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsBeforeExecutionIterator{contract: _EntryPointSimulations.contract, event: "BeforeExecution", logs: logs, sub: sub}, nil
}

// WatchBeforeExecution is a free log subscription operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_EntryPointSimulations *EntryPointSimulationsFilterer) WatchBeforeExecution(opts *bind.WatchOpts, sink chan<- *EntryPointSimulationsBeforeExecution) (event.Subscription, error) {

	logs, sub, err := _EntryPointSimulations.contract.WatchLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSimulationsBeforeExecution)
				if err := _EntryPointSimulations.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
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
func (_EntryPointSimulations *EntryPointSimulationsFilterer) ParseBeforeExecution(log types.Log) (*EntryPointSimulationsBeforeExecution, error) {
	event := new(EntryPointSimulationsBeforeExecution)
	if err := _EntryPointSimulations.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointSimulationsDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the EntryPointSimulations contract.
type EntryPointSimulationsDepositedIterator struct {
	Event *EntryPointSimulationsDeposited // Event containing the contract specifics and raw log

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
func (it *EntryPointSimulationsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSimulationsDeposited)
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
		it.Event = new(EntryPointSimulationsDeposited)
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
func (it *EntryPointSimulationsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSimulationsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSimulationsDeposited represents a Deposited event raised by the EntryPointSimulations contract.
type EntryPointSimulationsDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) FilterDeposited(opts *bind.FilterOpts, account []common.Address) (*EntryPointSimulationsDepositedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.FilterLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsDepositedIterator{contract: _EntryPointSimulations.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *EntryPointSimulationsDeposited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.WatchLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSimulationsDeposited)
				if err := _EntryPointSimulations.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_EntryPointSimulations *EntryPointSimulationsFilterer) ParseDeposited(log types.Log) (*EntryPointSimulationsDeposited, error) {
	event := new(EntryPointSimulationsDeposited)
	if err := _EntryPointSimulations.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointSimulationsPostOpRevertReasonIterator is returned from FilterPostOpRevertReason and is used to iterate over the raw logs and unpacked data for PostOpRevertReason events raised by the EntryPointSimulations contract.
type EntryPointSimulationsPostOpRevertReasonIterator struct {
	Event *EntryPointSimulationsPostOpRevertReason // Event containing the contract specifics and raw log

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
func (it *EntryPointSimulationsPostOpRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSimulationsPostOpRevertReason)
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
		it.Event = new(EntryPointSimulationsPostOpRevertReason)
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
func (it *EntryPointSimulationsPostOpRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSimulationsPostOpRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSimulationsPostOpRevertReason represents a PostOpRevertReason event raised by the EntryPointSimulations contract.
type EntryPointSimulationsPostOpRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPostOpRevertReason is a free log retrieval operation binding the contract event 0xf62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f4792.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) FilterPostOpRevertReason(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*EntryPointSimulationsPostOpRevertReasonIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.FilterLogs(opts, "PostOpRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsPostOpRevertReasonIterator{contract: _EntryPointSimulations.contract, event: "PostOpRevertReason", logs: logs, sub: sub}, nil
}

// WatchPostOpRevertReason is a free log subscription operation binding the contract event 0xf62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f4792.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) WatchPostOpRevertReason(opts *bind.WatchOpts, sink chan<- *EntryPointSimulationsPostOpRevertReason, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.WatchLogs(opts, "PostOpRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSimulationsPostOpRevertReason)
				if err := _EntryPointSimulations.contract.UnpackLog(event, "PostOpRevertReason", log); err != nil {
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
func (_EntryPointSimulations *EntryPointSimulationsFilterer) ParsePostOpRevertReason(log types.Log) (*EntryPointSimulationsPostOpRevertReason, error) {
	event := new(EntryPointSimulationsPostOpRevertReason)
	if err := _EntryPointSimulations.contract.UnpackLog(event, "PostOpRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointSimulationsSignatureAggregatorChangedIterator is returned from FilterSignatureAggregatorChanged and is used to iterate over the raw logs and unpacked data for SignatureAggregatorChanged events raised by the EntryPointSimulations contract.
type EntryPointSimulationsSignatureAggregatorChangedIterator struct {
	Event *EntryPointSimulationsSignatureAggregatorChanged // Event containing the contract specifics and raw log

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
func (it *EntryPointSimulationsSignatureAggregatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSimulationsSignatureAggregatorChanged)
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
		it.Event = new(EntryPointSimulationsSignatureAggregatorChanged)
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
func (it *EntryPointSimulationsSignatureAggregatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSimulationsSignatureAggregatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSimulationsSignatureAggregatorChanged represents a SignatureAggregatorChanged event raised by the EntryPointSimulations contract.
type EntryPointSimulationsSignatureAggregatorChanged struct {
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSignatureAggregatorChanged is a free log retrieval operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) FilterSignatureAggregatorChanged(opts *bind.FilterOpts, aggregator []common.Address) (*EntryPointSimulationsSignatureAggregatorChangedIterator, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.FilterLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsSignatureAggregatorChangedIterator{contract: _EntryPointSimulations.contract, event: "SignatureAggregatorChanged", logs: logs, sub: sub}, nil
}

// WatchSignatureAggregatorChanged is a free log subscription operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) WatchSignatureAggregatorChanged(opts *bind.WatchOpts, sink chan<- *EntryPointSimulationsSignatureAggregatorChanged, aggregator []common.Address) (event.Subscription, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.WatchLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSimulationsSignatureAggregatorChanged)
				if err := _EntryPointSimulations.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
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
func (_EntryPointSimulations *EntryPointSimulationsFilterer) ParseSignatureAggregatorChanged(log types.Log) (*EntryPointSimulationsSignatureAggregatorChanged, error) {
	event := new(EntryPointSimulationsSignatureAggregatorChanged)
	if err := _EntryPointSimulations.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointSimulationsStakeLockedIterator is returned from FilterStakeLocked and is used to iterate over the raw logs and unpacked data for StakeLocked events raised by the EntryPointSimulations contract.
type EntryPointSimulationsStakeLockedIterator struct {
	Event *EntryPointSimulationsStakeLocked // Event containing the contract specifics and raw log

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
func (it *EntryPointSimulationsStakeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSimulationsStakeLocked)
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
		it.Event = new(EntryPointSimulationsStakeLocked)
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
func (it *EntryPointSimulationsStakeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSimulationsStakeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSimulationsStakeLocked represents a StakeLocked event raised by the EntryPointSimulations contract.
type EntryPointSimulationsStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeLocked is a free log retrieval operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) FilterStakeLocked(opts *bind.FilterOpts, account []common.Address) (*EntryPointSimulationsStakeLockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.FilterLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsStakeLockedIterator{contract: _EntryPointSimulations.contract, event: "StakeLocked", logs: logs, sub: sub}, nil
}

// WatchStakeLocked is a free log subscription operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) WatchStakeLocked(opts *bind.WatchOpts, sink chan<- *EntryPointSimulationsStakeLocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.WatchLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSimulationsStakeLocked)
				if err := _EntryPointSimulations.contract.UnpackLog(event, "StakeLocked", log); err != nil {
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
func (_EntryPointSimulations *EntryPointSimulationsFilterer) ParseStakeLocked(log types.Log) (*EntryPointSimulationsStakeLocked, error) {
	event := new(EntryPointSimulationsStakeLocked)
	if err := _EntryPointSimulations.contract.UnpackLog(event, "StakeLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointSimulationsStakeUnlockedIterator is returned from FilterStakeUnlocked and is used to iterate over the raw logs and unpacked data for StakeUnlocked events raised by the EntryPointSimulations contract.
type EntryPointSimulationsStakeUnlockedIterator struct {
	Event *EntryPointSimulationsStakeUnlocked // Event containing the contract specifics and raw log

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
func (it *EntryPointSimulationsStakeUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSimulationsStakeUnlocked)
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
		it.Event = new(EntryPointSimulationsStakeUnlocked)
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
func (it *EntryPointSimulationsStakeUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSimulationsStakeUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSimulationsStakeUnlocked represents a StakeUnlocked event raised by the EntryPointSimulations contract.
type EntryPointSimulationsStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeUnlocked is a free log retrieval operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) FilterStakeUnlocked(opts *bind.FilterOpts, account []common.Address) (*EntryPointSimulationsStakeUnlockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.FilterLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsStakeUnlockedIterator{contract: _EntryPointSimulations.contract, event: "StakeUnlocked", logs: logs, sub: sub}, nil
}

// WatchStakeUnlocked is a free log subscription operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) WatchStakeUnlocked(opts *bind.WatchOpts, sink chan<- *EntryPointSimulationsStakeUnlocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.WatchLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSimulationsStakeUnlocked)
				if err := _EntryPointSimulations.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
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
func (_EntryPointSimulations *EntryPointSimulationsFilterer) ParseStakeUnlocked(log types.Log) (*EntryPointSimulationsStakeUnlocked, error) {
	event := new(EntryPointSimulationsStakeUnlocked)
	if err := _EntryPointSimulations.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointSimulationsStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the EntryPointSimulations contract.
type EntryPointSimulationsStakeWithdrawnIterator struct {
	Event *EntryPointSimulationsStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *EntryPointSimulationsStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSimulationsStakeWithdrawn)
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
		it.Event = new(EntryPointSimulationsStakeWithdrawn)
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
func (it *EntryPointSimulationsStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSimulationsStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSimulationsStakeWithdrawn represents a StakeWithdrawn event raised by the EntryPointSimulations contract.
type EntryPointSimulationsStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, account []common.Address) (*EntryPointSimulationsStakeWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.FilterLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsStakeWithdrawnIterator{contract: _EntryPointSimulations.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *EntryPointSimulationsStakeWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.WatchLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSimulationsStakeWithdrawn)
				if err := _EntryPointSimulations.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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
func (_EntryPointSimulations *EntryPointSimulationsFilterer) ParseStakeWithdrawn(log types.Log) (*EntryPointSimulationsStakeWithdrawn, error) {
	event := new(EntryPointSimulationsStakeWithdrawn)
	if err := _EntryPointSimulations.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointSimulationsUserOperationEventIterator is returned from FilterUserOperationEvent and is used to iterate over the raw logs and unpacked data for UserOperationEvent events raised by the EntryPointSimulations contract.
type EntryPointSimulationsUserOperationEventIterator struct {
	Event *EntryPointSimulationsUserOperationEvent // Event containing the contract specifics and raw log

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
func (it *EntryPointSimulationsUserOperationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSimulationsUserOperationEvent)
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
		it.Event = new(EntryPointSimulationsUserOperationEvent)
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
func (it *EntryPointSimulationsUserOperationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSimulationsUserOperationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSimulationsUserOperationEvent represents a UserOperationEvent event raised by the EntryPointSimulations contract.
type EntryPointSimulationsUserOperationEvent struct {
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
func (_EntryPointSimulations *EntryPointSimulationsFilterer) FilterUserOperationEvent(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (*EntryPointSimulationsUserOperationEventIterator, error) {

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

	logs, sub, err := _EntryPointSimulations.contract.FilterLogs(opts, "UserOperationEvent", userOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsUserOperationEventIterator{contract: _EntryPointSimulations.contract, event: "UserOperationEvent", logs: logs, sub: sub}, nil
}

// WatchUserOperationEvent is a free log subscription operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) WatchUserOperationEvent(opts *bind.WatchOpts, sink chan<- *EntryPointSimulationsUserOperationEvent, userOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EntryPointSimulations.contract.WatchLogs(opts, "UserOperationEvent", userOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSimulationsUserOperationEvent)
				if err := _EntryPointSimulations.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
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
func (_EntryPointSimulations *EntryPointSimulationsFilterer) ParseUserOperationEvent(log types.Log) (*EntryPointSimulationsUserOperationEvent, error) {
	event := new(EntryPointSimulationsUserOperationEvent)
	if err := _EntryPointSimulations.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointSimulationsUserOperationPrefundTooLowIterator is returned from FilterUserOperationPrefundTooLow and is used to iterate over the raw logs and unpacked data for UserOperationPrefundTooLow events raised by the EntryPointSimulations contract.
type EntryPointSimulationsUserOperationPrefundTooLowIterator struct {
	Event *EntryPointSimulationsUserOperationPrefundTooLow // Event containing the contract specifics and raw log

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
func (it *EntryPointSimulationsUserOperationPrefundTooLowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSimulationsUserOperationPrefundTooLow)
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
		it.Event = new(EntryPointSimulationsUserOperationPrefundTooLow)
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
func (it *EntryPointSimulationsUserOperationPrefundTooLowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSimulationsUserOperationPrefundTooLowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSimulationsUserOperationPrefundTooLow represents a UserOperationPrefundTooLow event raised by the EntryPointSimulations contract.
type EntryPointSimulationsUserOperationPrefundTooLow struct {
	UserOpHash [32]byte
	Sender     common.Address
	Nonce      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserOperationPrefundTooLow is a free log retrieval operation binding the contract event 0x67b4fa9642f42120bf031f3051d1824b0fe25627945b27b8a6a65d5761d5482e.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender, uint256 nonce)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) FilterUserOperationPrefundTooLow(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*EntryPointSimulationsUserOperationPrefundTooLowIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.FilterLogs(opts, "UserOperationPrefundTooLow", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsUserOperationPrefundTooLowIterator{contract: _EntryPointSimulations.contract, event: "UserOperationPrefundTooLow", logs: logs, sub: sub}, nil
}

// WatchUserOperationPrefundTooLow is a free log subscription operation binding the contract event 0x67b4fa9642f42120bf031f3051d1824b0fe25627945b27b8a6a65d5761d5482e.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender, uint256 nonce)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) WatchUserOperationPrefundTooLow(opts *bind.WatchOpts, sink chan<- *EntryPointSimulationsUserOperationPrefundTooLow, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.WatchLogs(opts, "UserOperationPrefundTooLow", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSimulationsUserOperationPrefundTooLow)
				if err := _EntryPointSimulations.contract.UnpackLog(event, "UserOperationPrefundTooLow", log); err != nil {
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
func (_EntryPointSimulations *EntryPointSimulationsFilterer) ParseUserOperationPrefundTooLow(log types.Log) (*EntryPointSimulationsUserOperationPrefundTooLow, error) {
	event := new(EntryPointSimulationsUserOperationPrefundTooLow)
	if err := _EntryPointSimulations.contract.UnpackLog(event, "UserOperationPrefundTooLow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointSimulationsUserOperationRevertReasonIterator is returned from FilterUserOperationRevertReason and is used to iterate over the raw logs and unpacked data for UserOperationRevertReason events raised by the EntryPointSimulations contract.
type EntryPointSimulationsUserOperationRevertReasonIterator struct {
	Event *EntryPointSimulationsUserOperationRevertReason // Event containing the contract specifics and raw log

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
func (it *EntryPointSimulationsUserOperationRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSimulationsUserOperationRevertReason)
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
		it.Event = new(EntryPointSimulationsUserOperationRevertReason)
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
func (it *EntryPointSimulationsUserOperationRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSimulationsUserOperationRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSimulationsUserOperationRevertReason represents a UserOperationRevertReason event raised by the EntryPointSimulations contract.
type EntryPointSimulationsUserOperationRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserOperationRevertReason is a free log retrieval operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) FilterUserOperationRevertReason(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*EntryPointSimulationsUserOperationRevertReasonIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.FilterLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsUserOperationRevertReasonIterator{contract: _EntryPointSimulations.contract, event: "UserOperationRevertReason", logs: logs, sub: sub}, nil
}

// WatchUserOperationRevertReason is a free log subscription operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) WatchUserOperationRevertReason(opts *bind.WatchOpts, sink chan<- *EntryPointSimulationsUserOperationRevertReason, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.WatchLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSimulationsUserOperationRevertReason)
				if err := _EntryPointSimulations.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
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
func (_EntryPointSimulations *EntryPointSimulationsFilterer) ParseUserOperationRevertReason(log types.Log) (*EntryPointSimulationsUserOperationRevertReason, error) {
	event := new(EntryPointSimulationsUserOperationRevertReason)
	if err := _EntryPointSimulations.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointSimulationsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the EntryPointSimulations contract.
type EntryPointSimulationsWithdrawnIterator struct {
	Event *EntryPointSimulationsWithdrawn // Event containing the contract specifics and raw log

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
func (it *EntryPointSimulationsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSimulationsWithdrawn)
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
		it.Event = new(EntryPointSimulationsWithdrawn)
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
func (it *EntryPointSimulationsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSimulationsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSimulationsWithdrawn represents a Withdrawn event raised by the EntryPointSimulations contract.
type EntryPointSimulationsWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*EntryPointSimulationsWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointSimulationsWithdrawnIterator{contract: _EntryPointSimulations.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPointSimulations *EntryPointSimulationsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *EntryPointSimulationsWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPointSimulations.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSimulationsWithdrawn)
				if err := _EntryPointSimulations.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_EntryPointSimulations *EntryPointSimulationsFilterer) ParseWithdrawn(log types.Log) (*EntryPointSimulationsWithdrawn, error) {
	event := new(EntryPointSimulationsWithdrawn)
	if err := _EntryPointSimulations.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
