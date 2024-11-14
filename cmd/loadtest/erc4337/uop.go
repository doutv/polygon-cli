package erc4337loadtest

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	// "github.com/ethereum/go-ethereum/accounts/abi"
	polyabi "github.com/0xPolygon/polygon-cli/abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// UserOperation represents an ERC-4337 User Operation
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

const modeType = "0x0000000000000000000000000000000000000000000000000000000000000000"

var pubKeyX, _ = new(big.Int).SetString("640c5cacef387563d0b105c7724c45ee19f8a952cb583de494a6a7ce5ed16760", 16)
var pubKeyY, _ = new(big.Int).SetString("142b33cbf8255e9f0628ab9e250e179a3e7e8e24e0a2a4340f0b9fdeb29a1b48", 16)

// UserOperationRequest represents the request format for RPC calls
type UserOperationRequest struct {
	UserOperation
	EntryPoint  common.Address
	BlockNumber *big.Int
	BlockHash   common.Hash
}

// Add these ABI constants
// const smartAccountABI = `[
//     {"inputs":[{"internalType":"address","name":"recoveryModule","type":"address"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"installRecoveryModule","outputs":[],"stateMutability":"nonpayable","type":"function"},
//     {"inputs":[{"internalType":"uint256","name":"typeId","type":"uint256"},{"internalType":"address","name":"module","type":"address"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"installModule","outputs":[],"stateMutability":"nonpayable","type":"function"}
// ]`

func smartAccountInstallRecoveryModule(module common.Address, data []byte) ([]byte, error) {
	calldataStr, err := polyabi.AbiEncode("installRecoveryModule(address,bytes)", []string{module.String(), string(data)})
	if err != nil {
		return nil, err
	}
	return []byte(calldataStr), nil
}

func smartAccountInstallFallbackModule(id int64, module common.Address, data []byte) ([]byte, error) {
	calldataStr, err := polyabi.AbiEncode("installModule(uint256,address,bytes)", []string{strconv.FormatInt(id, 10), module.String(), string(data)})
	if err != nil {
		return nil, err
	}
	return []byte(calldataStr), nil
}

func InitcodeCalldata(
	tops *bind.TransactOpts,
	cfg *ERC4337Config,
	salt *big.Int,
) (sender common.Address, initCode []byte, err error) {
	cops := &bind.CallOpts{}
	// Calculate the sender address (counterfactual address)
	sender, err = cfg.AccountFactory.Contract.ComputeAddress(
		cops,
		cfg.PayableAccount.Address,
		salt,
	)
	if err != nil {
		return
	}

	// Create the initialization data for the account
	installRecoveryModuleCalldata, err := smartAccountInstallRecoveryModule(
		tops.From,
		[]byte{},
	)
	if err != nil {
		return
	}

	installFallbackModuleCalldata, err := smartAccountInstallFallbackModule(
		3,
		cfg.TokenReceiver.Address,
		[]byte{},
	)
	if err != nil {
		return
	}

	// Create the initcode
	initializer0, _, err := cfg.Helper.Contract.GetAccountInitializer2(
		cops,
		pubKeyX,
		pubKeyY,
		cfg.WebAuthnValidator.Address,
		tops.From,
		tops.From,
		installRecoveryModuleCalldata,
		installFallbackModuleCalldata,
	)

	calldataStr, err := polyabi.AbiEncode("createAccount(address,bytes,bytes)", []string{sender.String(), string(initializer0), salt.String()})
	if err != nil {
		return
	}
	calldata := []byte(calldataStr)
	initcode, err := cfg.Helper.Contract.EncodePacked(cops, cfg.AccountFactory.Address, calldata)
	if err != nil {
		return
	}
	return sender, initcode, nil
}

var salt *big.Int
func init() {
	salt = big.NewInt(2)
}

func SendInitUop(
	tops *bind.TransactOpts,
	cfg *ERC4337Config,
) (err error) {
	sender, initCode, err := InitcodeCalldata(tops, cfg, salt)
	if err != nil {
		return
	}

	executionCalldata, err := packExecutionCalldata(tops.From, big.NewInt(1), []byte{})
	if err != nil {
		return
	}
	calldata, err := polyabi.AbiEncode("execute(bytes32,bytes)", []string{modeType, string(executionCalldata)})
	if err != nil {
		return
	}

	return nil
}

func GenerateUops(
	client *ethclient.Client,
	ctx context.Context,
	tops *bind.TransactOpts,
	cfg *ERC4337Config,
	sender common.Address,
	count int64,
	callData []byte,
	initCode []byte,
) ([]UserOperation, error) {
	if count == 0 {
		return nil, nil
	}

	// Check if account needs initialization
	code, err := client.CodeAt(ctx, sender, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get code at address: %v", err)
	}
	if code 

	needInit := len(code) == 0
	userOps := make([]UserOperation, 0, count)

	for i := int64(0); i < count; i++ {
		nonce := new(big.Int).Add(currentNonce, big.NewInt(i))

		// Create base UserOperation
		userOp := types.UserOperation{
			Sender:               sender,
			Nonce:                nonce,
			InitCode:             []byte{},
			CallData:             callData,
			CallGasLimit:         big.NewInt(100000),
			VerificationGasLimit: big.NewInt(2000000),
			PreVerificationGas:   big.NewInt(0),
			MaxFeePerGas:         big.NewInt(1e9),
			MaxPriorityFeePerGas: big.NewInt(1e9),
			PaymasterAndData:     []byte{},
			Signature:            []byte{},
		}

		// Add initCode only for first operation if needed
		if i == 0 && needInit {
			userOp.InitCode = initCode
		}

		// Generate signature
		sig, err := g.generateSignatureForUop(tops, userOp, entryPoint, helper)
		if err != nil {
			return nil, fmt.Errorf("failed to generate signature: %v", err)
		}

		userOp.Signature = sig
		userOps = append(userOps, userOp)
	}

	return userOps, nil
}

func (g *CallDataGenerator) generateSignatureForUop(
	auth *bind.TransactOpts,
	userOp types.UserOperation,
	entryPoint *EntryPoint,
	helper *Helper,
) ([]byte, error) {
	// Get the user operation hash
	userOpHash, err := helper.GetUserOpHash(userOp, entryPoint.Address())
	if err != nil {
		return nil, fmt.Errorf("failed to get user op hash: %v", err)
	}

	// Sign the hash
	signature, err := crypto.Sign(userOpHash[:], auth.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign user op: %v", err)
	}

	return signature, nil
}

func packExecutionCalldata(address common.Address, value *big.Int, data []byte) ([]byte, error) {
	arguments := abi.Arguments{
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.UintTy, Size: 256}},
		{Type: abi.Type{T: abi.BytesTy}},
	}
	
	packed, err := arguments.Pack(address, value, data)
	if err != nil {
		return nil, err
	}
	return packed, nil
}
