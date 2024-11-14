package erc4337loadtest

// import (
// 	"context"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/common"
// )

// type UserOperation struct {
// 	Sender               common.Address
// 	Nonce               *big.Int
// 	InitCode            []byte
// 	CallData            []byte
// 	CallGasLimit       *big.Int
// 	VerificationGasLimit *big.Int
// 	PreVerificationGas  *big.Int
// 	MaxFeePerGas       *big.Int
// 	MaxPriorityFeePerGas *big.Int
// 	PaymasterAndData    []byte
// 	Signature           []byte
// }

// func GenerateUserOperation(
// 	config *ERC4337Config,
// 	sender common.Address,
// 	nonce *big.Int,
// 	initCode []byte,
// 	callData []byte,
// ) (*UserOperation, error) {
// 	// Similar to calldataGenerator.generateUopsOfSender in TypeScript
// 	op := &UserOperation{
// 		Sender:    sender,
// 		Nonce:     nonce,
// 		InitCode:  initCode,
// 		CallData:  callData,
// 	}

// 	// Estimate gas limits
// 	callGasLimit, err := estimateCallGasLimit(config, op)
// 	if err != nil {
// 		return nil, err
// 	}
// 	op.CallGasLimit = callGasLimit

// 	// Set other gas parameters
// 	op.VerificationGasLimit = new(big.Int).SetUint64(150000) // Example value
// 	op.PreVerificationGas = new(big.Int).SetUint64(21000)    // Example value

// 	// Get current gas prices from network
// 	gasPrice, err := config.Client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}
// 	op.MaxFeePerGas = new(big.Int).Mul(gasPrice, big.NewInt(2))
// 	op.MaxPriorityFeePerGas = gasPrice

// 	return op, nil
// }
