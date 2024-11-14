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
	Bin: "0x60a080604052346100ff57600160025561015f8181016001600160401b038111838210176100e95782916142c3833903906000f080156100dd57608052610044610104565b6000815260208101906000825280602061005c610104565b600081520152600380546001600160a01b0319169055516004555160055560644310156100985760405161419f90816101248239608051815050f35b60405162461bcd60e51b815260206004820152601660248201527f73686f756c64206e6f74206265206465706c6f796564000000000000000000006044820152606490fd5b6040513d6000823e3d90fd5b634e487b7160e01b600052604160045260246000fd5b600080fd5b60408051919082016001600160401b038111838210176100e95760405256fe60806040526004361015610024575b361561001957600080fd5b610022336125ba565b005b60003560e01c806242dc5314611bc357806301ffc9a714611b255780630396cb60146118d05780630bd28e3b146118825780631b2e01b81461182c578063205c28781461170e57806322cdde4c146116ee57806335567e1a146116895780635287ce12146115ae57806370a0823114611574578063765e827f1461135b578063850aaf62146112d5578063957122ab1461119b57806397b2dcb914610d3d5780639b249f6914610c65578063b760faf914610c49578063bb9fe6bf14610b22578063c23a5cea14610951578063c3bce00914610653578063dbed18e0146101905763fc7e286d0361000e573461018b57602036600319011261018b576001600160a01b036101306121e7565b16600052600060205260a0604060002065ffffffffffff6001825492015460405192835260ff8116151560208401526001600160701b038160081c16604084015263ffffffff8160781c16606084015260981c166080820152f35b600080fd5b3461018b5761019e36612272565b906101a7612881565b60009160005b8281106104ac57506101bf84936124dd565b6000805b8481106102de5750507fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972600080a16000809360005b8181106102365761022f868660007f575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d8180a2613ce0565b6001600255005b61028461024482848a612643565b6001600160a01b036102586020830161269a565b167f575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d600080a280612665565b906000915b80831061029b575050506001016101f8565b909194976102d56102cf6001926102c98c8b6102c2826102bc8e8b8d612557565b9261252d565b5191613a33565b9061231b565b99612328565b95019190610289565b60206102eb828789612643565b6103016102f88280612665565b9390920161269a565b600092906001600160a01b039081165b8285106103255750505050506001016101c3565b90919293956103548361034d61033b848c61252d565b516103478b898b612557565b85613153565b92906140dd565b91168403610469576104235761036a84916140dd565b91166103e05761038a5761037f600191612328565b960193929190610311565b60a48760405190631101335b60e11b8252600482015260406024820152602160448201527f41413332207061796d61737465722065787069726564206f72206e6f742064756064820152606560f81b6084820152fd5b60848860405190631101335b60e11b8252600482015260406024820152601460448201527320a0999a1039b4b3b730ba3ab9329032b93937b960611b6064820152fd5b60848860405190631101335b60e11b82526004820152604060248201526017604482015276414132322065787069726564206f72206e6f742064756560481b6064820152fd5b60848960405190631101335b60e11b8252600482015260406024820152601460448201527320a0991a1039b4b3b730ba3ab9329032b93937b960611b6064820152fd5b6104b7818487612643565b936104c28580612665565b909590919060206001600160a01b036104dc82840161269a565b1697600192838a1461060e578961050c575b5050505060019293949550906105039161231b565b939291016101ad565b80604061051a920190612611565b918a3b1561018b5792939190604051948593632dd8113360e01b855288604486016040600488015252606490818601918a60051b8701019680936000915b8c83106105c957505050505083850360031901602485015250600093839261057f92612599565b03818a5afa90816105ba575b506105a95760405163086a9f7560e41b815260048101879052602490fd5b9394508493610503600189806104ee565b6105c390612130565b8861058b565b9193959697509193976063198a8203018652883561011e198336030181121561018b57836105fb8793858394016126df565b9a01960193019091899796959492610558565b60405162461bcd60e51b815260048101849052601760248201527f4141393620696e76616c69642061676772656761746f720000000000000000006044820152606490fd5b3461018b5761066136612240565b60405161066d816120ad565b604051610679816120ad565b600081526000602082015260006040820152600060608201526060608082015281526040516106a781612115565b600081526000602082015260208201526040516106c381612115565b600081526000602082015260408201526040516106df81612115565b600081526000602082015260608201526080604051916106fe83612115565b6000835260405161070e81612115565b60008152600060208201526020840152015261072861246c565b9061073281613e00565b61073c82826128a2565b835160e00151919291610757906001600160a01b0316613f78565b84515161079090610770906001600160a01b0316613f78565b936000602060405161078181612115565b82815201526040810190612611565b909590601481106109425760141161018b576107b16080963560601c613f78565b9060018060a01b0386169387820151966060604084015193015192604051986107d98a6120ad565b8952602089015260408801526060870152868601526040516107fa81612115565b6003546001600160a01b0316815260405161081481612115565b6004548152600554602082015260208201529280151580610937575b610910575b5060405194610843866120ad565b8552602085019384526040850152606084015283830152602080604051948594828652826108b1838351610140848b015280516101608b0152838101516101808b015260408101516101a08b015260608101516101c08b0152015160a06101e08a01526102008901906122f6565b95518051604089015201516060870152826040820151805184890152015160a0870152826060820151805160c0890152015160e0870152015160018060a01b038151166101008601520151805161012085015201516101408301520390f35b90925061091c81613f78565b6040519161092983612115565b825260208201529186610835565b506001811415610830565b50608095506107b16000613f78565b3461018b5760208060031936011261018b5761096b6121e7565b33600052600082526001604060002001908154916001600160701b038360081c16928315610ae65765ffffffffffff8160981c168015610aa1574210610a5c57610100600160c81b0319169055604080516001600160a01b03831681526020810184905260009384938493849333917fb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda391a26001600160a01b03165af1610a10612344565b5015610a1857005b6064906040519062461bcd60e51b82526004820152601860248201527f6661696c656420746f207769746864726177207374616b6500000000000000006044820152fd5b60405162461bcd60e51b815260048101869052601b60248201527f5374616b65207769746864726177616c206973206e6f742064756500000000006044820152606490fd5b60405162461bcd60e51b815260048101879052601d60248201527f6d7573742063616c6c20756e6c6f636b5374616b6528292066697273740000006044820152606490fd5b60405162461bcd60e51b81526004810186905260146024820152734e6f207374616b6520746f20776974686472617760601b6044820152606490fd5b3461018b57600036600319011261018b573360005260006020526001604060002001805463ffffffff8160781c16908115610c175760ff1615610bde5765ffffffffffff908142160191818311610bc857805460ff65ffffffffffff60981b01191665ffffffffffff60981b609885901b161790556040519116815233907ffa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a90602090a2005b634e487b7160e01b600052601160045260246000fd5b60405162461bcd60e51b8152602060048201526011602482015270616c726561647920756e7374616b696e6760781b6044820152606490fd5b60405162461bcd60e51b815260206004820152600a6024820152691b9bdd081cdd185ad95960b21b6044820152606490fd5b602036600319011261018b57610022610c606121e7565b6125ba565b3461018b57602036600319011261018b576004356001600160401b03811161018b576020610c9a610cd39236906004016121fd565b600654604051632b870d1b60e11b815260048101859052946001600160a01b039493869392861692849260009284926024840191612599565b03925af1908115610d3157602492600092610d00575b50604051633653dc0360e11b815291166004820152fd5b610d2391925060203d602011610d2a575b610d1b818361215e565b81019061257a565b9083610ce9565b503d610d11565b6040513d6000823e3d90fd5b3461018b5760031960603682011261018b576001600160401b0390816004351161018b57610120906004353603011261018b57610d786121d1565b9060443590811161018b57610d919036906004016121fd565b606060a0604051610da1816120fa565b600081526000602082015260006040820152600083820152600060808201520152610dca612881565b610dd261246c565b91610de1600435600401613e00565b610df0836004356004016128a2565b9092905a9460608101516040519687610e13606460043501600435600401612611565b919060009260038111611192575b638dd7712f60e01b936001600160e01b031916840361112e575050506000610eb1610edd60209384880151906040519182918783015260406024830152610e70606483016004356004016126df565b9060448301520390610e8a601f199283810183528261215e565b610ed16040519485926242dc5360e01b8985015261020060248501526102248401906122f6565b610ebe604484018c61399b565b82810360231901610204840152896122f6565b0390810183528261215e565b828151910182305af16000519860405215610fc5575b509596949550610f8f94606094600094906001600160a01b038316610f93575b505050608001519460405195610f28876120fa565b865260208601968752604086019081526060860191825260808601921515835260a0860193845260016002556040519687966020885251602088015251604087015251606086015251608085015251151560a08401525160c08084015260e08301906122f6565b0390f35b6000949650849395508390826040519384928337810182815203925af1916080610fbb612344565b9392908880610f13565b909596506000953d60201461111d575b63deaddead60e01b8703611021576084604051631101335b60e11b81526000600482015260406024820152600f60448201526e41413935206f7574206f662067617360881b6064820152fd5b610f8f9663deadaa5160e01b03611070575061104161104c915a90612337565b60808301519061231b565b61106560408301519161105e84614079565b8284614018565b965b96959488610ef3565b6111086110fd61110f9360208601518651907ff62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f4792602060018060a01b038451169301513d90610800808311611115575b50604051916020818401016040528083526000602084013e6110f360405192839283526040602084015260408301906122f6565b0390a35a90612337565b60808501519061231b565b90836127a4565b96611067565b9150386110bf565b955060206000803e60005195610fd5565b6242dc5360e01b60208401526102006024840152611188935061117a9161115a91610224850191612599565b611167604484018861399b565b82810360231901610204840152856122f6565b03601f1981018a528961215e565b6020600089610edd565b81359350610e21565b3461018b57606036600319011261018b576001600160401b0360043581811161018b576111cc9036906004016121fd565b9190506111d76121d1565b9060443590811161018b576111f09036906004016121fd565b9190921590816112cb575b50611286576014811015611229575b60405162461bcd60e51b81526020600482015260006024820152604490fd5b60141161018b573560601c3b1561124157808061120a565b60405162461bcd60e51b815260206004820152601b60248201527f41413330207061796d6173746572206e6f74206465706c6f79656400000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601960248201527f41413230206163636f756e74206e6f74206465706c6f796564000000000000006044820152606490fd5b90503b15836111fb565b3461018b57604036600319011261018b576112ee6121e7565b6024356001600160401b03811161018b57600091611311839236906004016121fd565b90816040519283928337810184815203915af461132c612344565b60408051632650415560e21b81529215156004840152602483015281906113579060448301906122f6565b0390fd5b3461018b5761136936612272565b611374929192612881565b61137d836124dd565b60005b8481106113f557506000927fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972600080a16000915b8583106113c55761022f8585613ce0565b9091936001906113eb6113d9878987612557565b6113e3888661252d565b519088613a33565b01940191906113b4565b6114206114196114078385979561252d565b51611413848987612557565b84613153565b91906140dd565b6001600160a01b0392918316611531576114eb5761143d906140dd565b91166114a85761145257600101929092611380565b60a49060405190631101335b60e11b8252600482015260406024820152602160448201527f41413332207061796d61737465722065787069726564206f72206e6f742064756064820152606560f81b6084820152fd5b60848260405190631101335b60e11b8252600482015260406024820152601460448201527320a0999a1039b4b3b730ba3ab9329032b93937b960611b6064820152fd5b60848360405190631101335b60e11b82526004820152604060248201526017604482015276414132322065787069726564206f72206e6f742064756560481b6064820152fd5b60848460405190631101335b60e11b8252600482015260406024820152601460448201527320a0991a1039b4b3b730ba3ab9329032b93937b960611b6064820152fd5b3461018b57602036600319011261018b576001600160a01b036115956121e7565b1660005260006020526020604060002054604051908152f35b3461018b57602036600319011261018b576115c76121e7565b600060806040516115d7816120ad565b828152826020820152826040820152826060820152015260018060a01b0316600052600060205260a060406000206080604051611613816120ad565b6001835493848352015490602081019060ff8316151582526001600160701b0360408201818560081c16815263ffffffff936060840193858760781c16855265ffffffffffff978891019660981c1686526040519788525115156020880152511660408601525116606084015251166080820152f35b3461018b57604036600319011261018b5760206116a46121e7565b6116ac61222a565b6001600160a01b0390911660009081526001835260408082206001600160c01b03841683528452908190205481519290911b67ffffffffffffffff1916178152f35b3461018b57602061170661170136612240565b612374565b604051908152f35b3461018b57604036600319011261018b576117276121e7565b60243590336000526000602052604060002080548084116117e75760008481948294611754838596612337565b9055604080516001600160a01b03831681526020810184905233917fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb91a26001600160a01b03165af16117a5612344565b50156117ad57005b60405162461bcd60e51b81526020600482015260126024820152716661696c656420746f20776974686472617760701b6044820152606490fd5b60405162461bcd60e51b815260206004820152601960248201527f576974686472617720616d6f756e7420746f6f206c61726765000000000000006044820152606490fd5b3461018b57604036600319011261018b576118456121e7565b61184d61222a565b9060018060a01b0316600052600160205260406000209060018060c01b03166000526020526020604060002054604051908152f35b3461018b57602036600319011261018b576004356001600160c01b0381169081900361018b5733600052600160205260406000209060005260205260406000206118cc8154612328565b9055005b60208060031936011261018b5760043563ffffffff9182821680920361018b5733600052600081526040600020928215611ae0576001840154908160781c168310611a9b5761192d6001600160701b039182349160081c1661231b565b938415611a6157818511611a2b579065ffffffffffff6119fa925460405190611955826120ad565b81528481019260018452604082019088168152606082018781526001608084019360008552336000526000895260406000209051815501945115159160ff6effffffffffffffffffffffffffff008754925160081b169263ffffffff60781b905160781b169316906cffffffffffffffffffffffffff60981b161717178355511681549065ffffffffffff60981b9060981b169065ffffffffffff60981b1916179055565b6040519283528201527fa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c0160403392a2005b60405162461bcd60e51b815260048101849052600e60248201526d7374616b65206f766572666c6f7760901b6044820152606490fd5b60405162461bcd60e51b81526004810184905260126024820152711b9bc81cdd185ad9481cdc1958da599a595960721b6044820152606490fd5b60405162461bcd60e51b815260048101839052601c60248201527f63616e6e6f7420646563726561736520756e7374616b652074696d65000000006044820152606490fd5b60405162461bcd60e51b815260048101839052601a60248201527f6d757374207370656369667920756e7374616b652064656c61790000000000006044820152606490fd5b3461018b57602036600319011261018b5760043563ffffffff60e01b811680910361018b5760209063307e35b760e11b8114908115611bb2575b8115611ba1575b8115611b90575b8115611b7f575b506040519015158152f35b6301ffc9a760e01b14905082611b74565b633e84f02160e01b81149150611b6d565b63cf28ef9760e01b81149150611b66565b63122a0e9b60e31b81149150611b5f565b3461018b5761020036600319011261018b576001600160401b0360043581811161018b573660238201121561018b57611c0690369060248160040135910161219a565b903660231901906101c0821261018b5761014060405192611c26846120ad565b1261018b57604051611c37816120de565b611c3f6121d1565b815260443560208201526064356040820152608435606082015260a435608082015260c43560a082015260e43560c0820152610104356001600160a01b038116810361018b5760e0820152610124356101008201526101443561012082015282526101643560208301526101843560408301526101a43560608301526101c43560808301526101e43590811161018b57611cdd9036906004016121fd565b91905a303303612068578251606081015195603f5a0260061c61271060a0840151890101116120575760009681519182611f8f575b5050505090611d2c915a900360808401510193369161219a565b915a600093835192611d3d84613fc8565b60e085015190926001600160a01b039091169081611e8f57505083516001600160a01b0316925b5a9003019260a06060820151910151016080850151840390818111611e7b575b50508202604084015191818310600014611de65750506003851015611dd257602094600203611dc3576117069293508093611dbe81614079565b614018565b505063deadaa5160e01b825250fd5b634e487b7160e01b84526021600452602484fd5b81611df692969795930390613ff2565b506003831015611e6757602094507f49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f60808683015192519460018060a01b03865116948860018060a01b0360e0890151169701519160405192835215898301528760408301526060820152a4611706565b634e487b7160e01b85526021600452602485fd5b6064919003600a0204909201918680611d84565b819491948151611ea1575b5050611d64565b60038a1015611f795760028a0315611e9a5760a087015191813b1561018b57611ef2928b60008094604051809781968295637c627b2160e01b845260048401526080602484015260848301906122f6565b8b8b0260448301528b60648301520393f19081611f65575b50611f5e57611357873d610800808211611f56575b5060405191602082840101604052818352602083013e604051632b5e552f60e21b81526020600482015291829160248301906122f6565b905083611f1f565b8880611e9a565b611f70919850612130565b60009689611f0a565b634e487b7160e01b600052602160045260246000fd5b9160009291838093602060018060a01b03885116910192f115611fb5575b808080611d12565b611d2c929195503d61080080821161204f575b50604051906020818301016040528082526000602083013e8051611ff3575b50506001949091611fad565b7f1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201602086015191602060018060a01b0385511694015161204560405192839283526040602084015260408301906122f6565b0390a38580611fe7565b905087611fc8565b63deaddead60e01b60005260206000fd5b60405162461bcd60e51b815260206004820152601760248201527f4141393220696e7465726e616c2063616c6c206f6e6c790000000000000000006044820152606490fd5b60a081019081106001600160401b038211176120c857604052565b634e487b7160e01b600052604160045260246000fd5b61014081019081106001600160401b038211176120c857604052565b60c081019081106001600160401b038211176120c857604052565b604081019081106001600160401b038211176120c857604052565b6001600160401b0381116120c857604052565b606081019081106001600160401b038211176120c857604052565b90601f801991011681019081106001600160401b038211176120c857604052565b6001600160401b0381116120c857601f01601f191660200190565b9291926121a68261217f565b916121b4604051938461215e565b82948184528183011161018b578281602093846000960137010152565b602435906001600160a01b038216820361018b57565b600435906001600160a01b038216820361018b57565b9181601f8401121561018b578235916001600160401b03831161018b576020838186019501011161018b57565b602435906001600160c01b038216820361018b57565b6003199060208183011261018b57600435916001600160401b03831161018b57826101209203011261018b5760040190565b90604060031983011261018b576004356001600160401b039283821161018b578060238301121561018b57816004013593841161018b5760248460051b8301011161018b576024908101929190356001600160a01b038116810361018b5790565b60005b8381106122e65750506000910152565b81810151838201526020016122d6565b9060209161230f815180928185528580860191016122d3565b601f01601f1916010190565b91908201809211610bc857565b6000198114610bc85760010190565b91908203918211610bc857565b3d1561236f573d906123558261217f565b91612363604051938461215e565b82523d6000602084013e565b606090565b604061238281830183612611565b90818351918237206123976060840184612611565b90818451918237209260c06123af60e0830183612611565b908186519182372091845195602087019460018060a01b03833516865260208301358789015260608801526080870152608081013560a087015260a081013582870152013560e0850152610100908185015283526101208301916001600160401b0391848410838511176120c857838252845190206101408501908152306101608601524661018086015260608452936101a001918211838310176120c8575251902090565b6001600160401b0381116120c85760051b60200190565b60405190612479826120ad565b604051608083612488836120de565b60009283815283602082015283604082015283606082015283838201528360a08201528360c08201528360e0820152836101008201528361012082015281528260208201528260408201528260608201520152565b906124e782612455565b6124f4604051918261215e565b8281528092612505601f1991612455565b019060005b82811061251657505050565b60209061252161246c565b8282850101520161250a565b80518210156125415760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b91908110156125415760051b8101359061011e198136030182121561018b570190565b9081602091031261018b57516001600160a01b038116810361018b5790565b908060209392818452848401376000828201840152601f01601f1916010190565b6001805b6005811061260a5750507f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c460206125f53484613ff2565b6040519081526001600160a01b0390931692a2565b81016125be565b903590601e198136030182121561018b57018035906001600160401b03821161018b5760200191813603831361018b57565b91908110156125415760051b81013590605e198136030182121561018b570190565b903590601e198136030182121561018b57018035906001600160401b03821161018b57602001918160051b3603831361018b57565b356001600160a01b038116810361018b5790565b9035601e198236030181121561018b5701602081359101916001600160401b03821161018b57813603831361018b57565b6101209181356001600160a01b0381169081900361018b5761274961272e6127a1956127829385526020860135602086015261271e60408701876126ae565b9091806040880152860191612599565b61273b60608601866126ae565b908583036060870152612599565b6080840135608084015260a084013560a084015260c084013560c084015261277460e08501856126ae565b9084830360e0860152612599565b9161279361010091828101906126ae565b929091818503910152612599565b90565b9291905a906000908551936127b885613fc8565b60e086015190926001600160a01b039182168061286d575050855116935b5a9003019360a06060820151910151016080870151850390818111612859575b5050830260408601519281841060001461283857505080612824575090816128229294611dbe81614079565b565b634e487b7160e01b81526021600452602490fd5b6128489082849398950390613ff2565b506128245750908361282292614018565b6064919003600a02049093019238806127f6565b909591509451156127d657600093506127d6565b60028054146128905760028055565b604051633ee5aeb560e01b8152600490fd5b916000915a8151946001600160a01b036128bb8261269a565b1686526020810135602087015260808101356001600160801b038116606088015260801c604087015260a081013560c087015260c08101356001600160801b03811661010088015260801c61012087015261291960e0820182612611565b801561313757603481106130f2578060141161018b578060241161018b5760341161018b57602481013560801c60a0880152601481013560801c60808801523560601c60e08701525b61296b81612374565b60208401526040860151956001600160781b038760c08301511760608301511760808301511760a0830151176101008301511761012083015117116130ad57604081015160608201510160808201510160a08201510160c08201510161010082015102845160018060a01b03815116906129e86040860186612611565b80612e98575b505060e001516001600160a01b03169060008215612e60575b6020612a33918c828b01516000868b604051978896879586936306608bdf60e21b8552600485016140bb565b0393f160009181612e2c575b50612abf576113573d610800808211612ab7575b50604051906020818301016040528082526000602083013e6040516365c8fd4d60e01b81526000600482015260606024820152600d60648201526c10504c8cc81c995d995c9d1959609a1b608482015260a0604482015291829160a48301906122f6565b905082612a53565b9115612dc5575b509760018060a01b03835116602084015160401c9060005260016020526040600020906000526020526040600020612afe8154612328565b90555a850311612d795760e090910151606091906001600160a01b0316612b3d575b6040850152606084015260809160a090910135905a900301910152565b92919550505a835160e08101516001600160a01b031660008181526020819052604090208054989290868a10612d2d576080600092889283612ba29d039055015192602089015183604051809d819582946314add44b60e21b84528a600485016140bb565b039286f1978860009160009a612ca5575b50612c33576113573d610800808211612c2b575b50604051906020818301016040528082526000602083013e6040516365c8fd4d60e01b81526000600482015260606024820152600d60648201526c10504cccc81c995d995c9d1959609a1b608482015260a0604482015291829160a48301906122f6565b905082612bc7565b97925a900311612c495790959192909190612b20565b60a4604051631101335b60e11b81526000600482015260406024820152602760448201527f41413336206f766572207061796d6173746572566572696669636174696f6e47606482015266185cd31a5b5a5d60ca1b6084820152fd5b915098503d806000833e612cb9818361215e565b81019060408183031261018b578051906001600160401b03821161018b5782601f83830101121561018b578181015191612cf28361217f565b93612d00604051958661215e565b83855260208483850101011161018b57602092612d249184808701918501016122d3565b01519838612bb3565b6084604051631101335b60e11b81526000600482015260406024820152601e60448201527f41413331207061796d6173746572206465706f73697420746f6f206c6f7700006064820152fd5b6084604051631101335b60e11b81526000600482015260406024820152601e60448201527f41413236206f76657220766572696669636174696f6e4761734c696d697400006064820152fd5b600052600060205260406000208054808411612de657839003905538612ac6565b6084604051631101335b60e11b81526000600482015260406024820152601760448201527610504c8c48191a591b89dd081c185e481c1c99599d5b99604a1b6064820152fd5b9091506020813d602011612e58575b81612e486020938361215e565b8101031261018b57519038612a3f565b3d9150612e3b565b50806000526000602052604060002054838111600014612e8b5750612a33602060005b915050612a07565b6020612a33918503612e83565b833b613061576000602060018060a01b036006541660408c51015160405180948193632b870d1b60e11b835285600484015282612ed9602482018a8c612599565b0393f1908115610d3157600091613042575b506001600160a01b0381168015612ff6578503612faa573b15612f5e5760141161018b5782907fd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d604060208b01519260018060a01b0360e08d510151168251913560601c82526020820152a338806129ee565b6084604051631101335b60e11b81526000600482015260406024820152602060448201527f4141313520696e6974436f6465206d757374206372656174652073656e6465726064820152fd5b6084604051631101335b60e11b81526000600482015260406024820152602060448201527f4141313420696e6974436f6465206d7573742072657475726e2073656e6465726064820152fd5b6084604051631101335b60e11b81526000600482015260406024820152601b60448201527f4141313320696e6974436f6465206661696c6564206f72204f4f4700000000006064820152fd5b61305b915060203d602011610d2a57610d1b818361215e565b38612eeb565b6084604051631101335b60e11b81526000600482015260406024820152601f60448201527f414131302073656e64657220616c726561647920636f6e7374727563746564006064820152fd5b60405162461bcd60e51b815260206004820152601860248201527f41413934206761732076616c756573206f766572666c6f7700000000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601d60248201527f4141393320696e76616c6964207061796d6173746572416e64446174610000006044820152606490fd5b5050600060e087015260006080870152600060a0870152612962565b926000905a845194906001600160a01b0361316d8661269a565b1686526020850135602087015260808501356001600160801b038116606088015260801c604087015260a085013560c087015260c08501356001600160801b03811661010088015260801c6101208701526131cb60e0860186612611565b801561397f57603481106130f2578060141161018b578060241161018b5760341161018b57602481013560801c60a0880152601481013560801c60808801523560601c60e08701525b61321d85612374565b60208301526040860151946001600160781b038660c08901511760608901511760808901511760a0890151176101008901511761012089015117116130ad57604087015160608801510160808801510160a08801510160c0880151016101008801510296835160018060a01b038151169061329b6040850185612611565b8061376a575b505060e001516001600160a01b03169060008215613732575b60206132e6918b828a01516000868a604051978896879586936306608bdf60e21b8552600485016140bb565b0393f1600091816136fe575b50613371573d8c610800808311613369575b50604051916020818401016040528083526000602084013e6113576040519283926365c8fd4d60e01b8452600484015260606024840152600d60648401526c10504c8cc81c995d995c9d1959609a1b608484015260a0604484015260a48301906122f6565b915082613304565b9a92939495969798999a9115613697575b509760018060a01b03835116602084015160401c90600052600160205260406000209060005260205260406000206133ba8154612328565b90555a85031161364b5760e090910151606091906001600160a01b03166133fe575b509060a09184959697986040608096015260608601520135905a900301910152565b969550505a9683519760018060a01b0360e08a015116806000526000602052604060002080548481106135ff5760806134629a9b9c600093878094039055015192602089015183604051809d819582946314add44b60e21b84528c600485016140bb565b039286f1978860009160009a613573575b506134f2573d8b6108008083116134ea575b50604051916020818401016040528083526000602084013e6113576040519283926365c8fd4d60e01b8452600484015260606024840152600d60648401526c10504cccc81c995d995c9d1959609a1b608484015260a0604484015260a48301906122f6565b915082613485565b9991929394959697989998925a900311613517575090969590949392919060806133dc565b60a49060405190631101335b60e11b8252600482015260406024820152602760448201527f41413336206f766572207061796d6173746572566572696669636174696f6e47606482015266185cd31a5b5a5d60ca1b6084820152fd5b915098503d90816000823e613588828261215e565b604081838101031261018b578051906001600160401b03821161018b57828101601f83830101121561018b5781810151916135c28361217f565b936135d0604051958661215e565b838552820160208483850101011161018b576020926135f69184808701918501016122d3565b01519838613473565b60848b60405190631101335b60e11b8252600482015260406024820152601e60448201527f41413331207061796d6173746572206465706f73697420746f6f206c6f7700006064820152fd5b60849060405190631101335b60e11b8252600482015260406024820152601e60448201527f41413236206f76657220766572696669636174696f6e4761734c696d697400006064820152fd5b600052600060205260406000208054808c116136b8578b9003905538613382565b60848460405190631101335b60e11b8252600482015260406024820152601760448201527610504c8c48191a591b89dd081c185e481c1c99599d5b99604a1b6064820152fd5b9091506020813d60201161372a575b8161371a6020938361215e565b8101031261018b575190386132f2565b3d915061370d565b508060005260006020526040600020548a811160001461375d57506132e6602060005b9150506132ba565b60206132e6918c03613755565b833b613933576000602060018060a01b036006541660408b51015160405180948193632b870d1b60e11b8352856004840152826137ab602482018a8c612599565b0393f1908115610d3157600091613914575b506001600160a01b03811680156138c857850361387c573b156138305760141161018b5782907fd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d604060208a01519260018060a01b0360e08c510151168251913560601c82526020820152a338806132a1565b60848d60405190631101335b60e11b8252600482015260406024820152602060448201527f4141313520696e6974436f6465206d757374206372656174652073656e6465726064820152fd5b60848e60405190631101335b60e11b8252600482015260406024820152602060448201527f4141313420696e6974436f6465206d7573742072657475726e2073656e6465726064820152fd5b60848f60405190631101335b60e11b8252600482015260406024820152601b60448201527f4141313320696e6974436f6465206661696c6564206f72204f4f4700000000006064820152fd5b61392d915060203d602011610d2a57610d1b818361215e565b386137bd565b60848d60405190631101335b60e11b8252600482015260406024820152601f60448201527f414131302073656e64657220616c726561647920636f6e7374727563746564006064820152fd5b5050600060e087015260006080870152600060a0870152613214565b60806101a091805160018060a01b03808251168652602082015160208701526040820151604087015260608201516060870152838201518487015260a082015160a087015260c082015160c087015260e08201511660e0860152610100808201519086015261012080910151908501526020810151610140850152604081015161016085015260608101516101808501520151910152565b9092915a60608201519160409283519687613a516060830183612611565b919060009260038111613cd7575b638dd7712f60e01b936001600160e01b0319168403613c7157505050613b08600092613ae8926020870151613aa78a5193849360208501528b602485015260648401906126df565b9060448301520390613ac1601f199283810183528261215e565b610ed189519485926242dc5360e01b602085015261020060248501526102248401906122f6565b613af5604484018a61399b565b82810360231901610204840152876122f6565b6020918183809351910182305af160005198865215613b2a575b505050505050565b909192939495965060003d8214613c67575b63deaddead60e01b8103613b865760848787805191631101335b60e11b835260048301526024820152600f60448201526e41413935206f7574206f662067617360881b6064820152fd5b93955091939290919063deadaa5160e01b03613bcc575050613bb06110fd613bbf93945a90612337565b9083015183611dbe8295614079565b905b388080808080613b22565b613c5994613c539293827ff62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f4792613c4894880151918851936110f38260018060a01b03875116960151913d90610800808311613c5f575b508051918581840101825280835260008684013e80805195869586528501528301906122f6565b60808401519061231b565b916127a4565b90613bc1565b915038613c21565b8181803e51613b3c565b6242dc5360e01b60208401526102006024840152613ccf9450919250613cc191613ca19161022485019190612599565b613cae604484018761399b565b82810360231901610204840152846122f6565b03601f19810189528861215e565b600087613b08565b81359350613a5f565b6001600160a01b03168015613d4d57600080809381935af1613d00612344565b5015613d0857565b60405162461bcd60e51b815260206004820152601f60248201527f41413931206661696c65642073656e6420746f2062656e6566696369617279006044820152606490fd5b60405162461bcd60e51b815260206004820152601860248201527f4141393020696e76616c69642062656e656669636961727900000000000000006044820152606490fd5b600060443d106127a157604051600319913d83016004833e81516001600160401b03918282113d602484011117613def57818401948551938411613df7573d85010160208487010111613def57506127a19291016020019061215e565b949350505050565b50949350505050565b604080516135a560f21b602082019081523060601b6022830152600160f81b6036830152601782526000939091613e3681612115565b51909120600680546001600160a01b0319166001600160a01b0392831617905590613e6381840182612611565b919092613e7d613e728361269a565b9260e0810190612611565b91303b15613f745792879492613ecc92613eb497958951988997889763957122ab60e01b8952606060048a01526064890191612599565b93166024860152848303600319016044860152612599565b0381305afa9081613f61575b50613f5d5760018260033d11613f4d575b6308c379a014613f06575b613efc575050565b51903d90823e3d90fd5b613f0e613d92565b80613f1a575b50613ef4565b805184925015613f14578361135784928351938493631101335b60e11b85526004850152602484015260448301906122f6565b50600483803e825160e01c613ee9565b5050565b613f6d90939193612130565b9138613ed8565b8780fd5b90604051613f8581612115565b600080825260208083018281526001600160a01b0390951682528190526040902060010154600881901c6001600160701b0316825260781c63ffffffff16909252565b610120610100820151910151808214613fee57480180821015613fe9575090565b905090565b5090565b60018060a01b03166000526000602052614012604060002091825461231b565b80915590565b9190917f49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f6080602083015192519460018060a01b03946020868851169660e089015116970151916040519283526000602084015260408301526060820152a4565b60208101519051907f67b4fa9642f42120bf031f3051d1824b0fe25627945b27b8a6a65d5761d5482e60208060018060a01b03855116940151604051908152a3565b6140d3604092959493956060835260608301906126df565b9460208201520152565b8015614160576000604080516140f281612143565b8281526020810183905201526001600160a01b0381169065ffffffffffff9060409060a081901c8316908115614158575b60d01c9282519161413383612143565b858352846020840152169182910152421190811561415057509091565b905042109091565b839150614123565b5060009060009056fea2646970667358221220020f618d6e660ecf6dca88ad284db671aa60bef10584f0838147de19ca5ae56364736f6c6343000819003360808060405234601557610144908161001b8239f35b600080fdfe6080600436101561000f57600080fd5b6000803560e01c63570e1a361461002557600080fd5b3461010b57602036600319011261010b576004359167ffffffffffffffff9081841161010757366023850112156101075783600401358281116101035736602482870101116101035780601411610103576013198101928084116100ef57600b8201601f19908116603f01168301908111838210176100ef5792846024819482600c60209a968b9960405286845289840196603889018837830101525193013560601c5af190805191156100e7575b506040516001600160a01b039091168152f35b9050386100d4565b634e487b7160e01b85526041600452602485fd5b8380fd5b8280fd5b80fdfea26469706673582212206a1b9f618c552b5e5a558cda5cc2bafd851f85b1d0662028a1684dd5741f8d9464736f6c63430008190033",
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
