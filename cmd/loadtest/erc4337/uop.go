package erc4337loadtest

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	// "github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/0xPolygon/polygon-cli/bindings/4337/accountfactory"
	"github.com/0xPolygon/polygon-cli/bindings/4337/entryPoint/core/entrypoint"
	"github.com/0xPolygon/polygon-cli/bindings/4337/payableaccount"
	"github.com/0xPolygon/polygon-cli/bindings/4337/test/helper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
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

var passkeyPubX, _ = new(big.Int).SetString("640c5cacef387563d0b105c7724c45ee19f8a952cb583de494a6a7ce5ed16760", 16)
var passkeyPubY, _ = new(big.Int).SetString("142b33cbf8255e9f0628ab9e250e179a3e7e8e24e0a2a4340f0b9fdeb29a1b48", 16)
var passkeyPrivKey, _ = new(big.Int).SetString("42d2bd030a8a71ff2f9043adcfb46138a5d87287cefff37d18a638f956c33449", 16)

var salt *big.Int
var modeType [32]byte
func init() {
	salt = big.NewInt(2)
}

func SendUops(
	client *ethclient.Client,
	ctx context.Context,
	tops *bind.TransactOpts,
	cops *bind.CallOpts,
	eoaPrivateKey *ecdsa.PrivateKey,
	cfg *ERC4337Config,
) (err error) {
	// If account is not created, initcode is used
	sender, initCode, err := generateInitcode(tops, cops, cfg, salt)
	if err != nil {
		return
	}

	// empty execution calldata
	accountAbi, err := payableaccount.PayableAccountMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	calldata, err := accountAbi.Pack("execute", modeType, []byte{})
	if err != nil {
		panic(err)
	}

	userOps, err := generateUops(client, ctx, tops, cops, eoaPrivateKey, cfg, sender, 1, []byte(calldata), initCode)
	if err != nil {
		return
	}
	log.Trace().Interface("userOps", userOps).Msg("Sending user operations")
	_, err = cfg.EntryPoint.Contract.HandleOps(tops, userOps, tops.From)
	if err != nil {
		return
	}

	return nil
}

func generateInitcode(
	tops *bind.TransactOpts,
	cops *bind.CallOpts,
	cfg *ERC4337Config,
	salt *big.Int,
) (sender common.Address, initCode []byte, err error) {
	// Calculate the sender address (counterfactual address)
	sender, err = cfg.AccountFactory.Contract.ComputeAddress(
		cops,
		cfg.PayableAccount.Address,
		salt,
	)
	if err != nil {
		panic(err)
	}

	// Create the initialization data for the account
	// tx, err := cfg.PayableAccount.Contract.InstallRecoveryModule(topsNoSend, topsNoSend.From, []byte{})
	// if err != nil {
	// 	panic(err)
	// }
	// installRecoveryModuleCalldata := tx.Data()
	// tx, err = cfg.PayableAccount.Contract.InstallModule(topsNoSend, big.NewInt(3), cfg.TokenReceiver.Address, nil)
	// if err != nil {
	// 	panic(err)
	// }
	// installFallbackModuleCalldata := tx.Data()

	// Create the initcode
	initializer0, _, err := cfg.Helper.Contract.GetAccountInitializer2(
		cops,
		passkeyPubX,
		passkeyPubY,
		cfg.WebAuthnValidator.Address,
		tops.From,
		tops.From,
		[]byte{},
		[]byte{},
	)
	if err != nil {
		panic(err)
	}

	factoryAbi, err := accountfactory.AccountFactoryMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	calldata, err := factoryAbi.Pack("createAccount", cfg.PayableAccount.Address, initializer0, salt)
	if err != nil {
		panic(err)
	}
	initcode, err := cfg.Helper.Contract.EncodePacked(cops, cfg.AccountFactory.Address, calldata)
	if err != nil {
		panic(err)
	}
	return sender, initcode, nil
}

func generateUops(
	client *ethclient.Client,
	ctx context.Context,
	tops *bind.TransactOpts,
	cops *bind.CallOpts,
	eoaPrivateKey *ecdsa.PrivateKey,
	cfg *ERC4337Config,
	sender common.Address,
	count uint32,
	callData []byte,
	initCode []byte,
) ([]entrypoint.PackedUserOperation, error) {
	if count == 0 {
		return nil, fmt.Errorf("no uops to generate")
	}

	// Check if account needs initialization
	code, err := client.CodeAt(ctx, sender, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get code at address: %v", err)
	}

	userOps := make([]entrypoint.PackedUserOperation, 0, count)

	for i := uint32(0); i < count; i++ {
		// Create base UserOperation
		userOp := PackUserOp(UserOperation{
			Sender:               sender,
			Nonce:                tops.Nonce,
			InitCode:             func() []byte { if code == nil { return initCode }; return nil }(),
			CallData:             callData,
			CallGasLimit:         big.NewInt(100000),
			VerificationGasLimit: big.NewInt(2000000),
			PreVerificationGas:   big.NewInt(0),
			MaxFeePerGas:         big.NewInt(1e9),
			MaxPriorityFeePerGas: big.NewInt(1e9),
			PaymasterAndData:     nil,
			Signature:            nil,
		})
		// nonce++
		tops.Nonce = new(big.Int).Add(tops.Nonce, big.NewInt(1))

		// Generate signature
		sig, err := generateSignatureForUop(tops, cops, userOp, cfg.EntryPoint.Contract, cfg.Helper.Contract, eoaPrivateKey)
		if err != nil {
			return nil, fmt.Errorf("failed to generate signature: %v", err)
		}

		userOp.Signature = sig
		userOps = append(userOps, userOp)
	}

	return userOps, nil
}

// PackAccountGasLimits packs two gas limits into a single hex string
func PackAccountGasLimits(verificationGasLimit, callGasLimit *big.Int) [32]byte {
    // Pad both values to 16 bytes (128 bits) each
    verificationGasHex := common.LeftPadBytes(verificationGasLimit.Bytes(), 16)
    callGasHex := common.LeftPadBytes(callGasLimit.Bytes(), 16)
    
    // Concatenate the byte slices
    return [32]byte(append(verificationGasHex, callGasHex...))
}

// PackUserOp packs a user operation into its compact form
func PackUserOp(userOp UserOperation) entrypoint.PackedUserOperation {
    accountGasLimits := PackAccountGasLimits(
        userOp.VerificationGasLimit,
        userOp.CallGasLimit,
    )
    gasFees := PackAccountGasLimits(
        userOp.MaxPriorityFeePerGas,
        userOp.MaxFeePerGas,
    )

    return entrypoint.PackedUserOperation{
        Sender:            userOp.Sender,
        Nonce:            userOp.Nonce,
        CallData:         userOp.CallData,
        AccountGasLimits: accountGasLimits,
        InitCode:         userOp.InitCode,
        PreVerificationGas: userOp.PreVerificationGas,
        GasFees:          gasFees,
        PaymasterAndData: userOp.PaymasterAndData,
    }
}

func generateSignatureForUop(
	tops *bind.TransactOpts,
	cops *bind.CallOpts,
	userOp entrypoint.PackedUserOperation,
	entryPoint *entrypoint.EntryPoint,
	helper *helper.Helper,
	eoaPrivateKey *ecdsa.PrivateKey,
) ([]byte, error) {
	// Calculate expiration time (10 minutes from now)
	expireTime := big.NewInt(time.Now().Unix() + 600)
	validationData, err := helper.GetValidationData(cops, expireTime)
	if err != nil {
		panic(err)
	}
	// Verify validation date
	_, outOfTimeRange, err := helper.CheckValidationDate(cops, validationData)
	if err != nil || outOfTimeRange {
		panic(fmt.Errorf("validation date check failed: %v", err))
	}

	// Get UserOp hash from EntryPoint
	entrypointUopHash, err := entryPoint.GetUserOpHash(cops, userOp)
	if err != nil {
		panic(err)
	}
	// Encode UOP hash
	uopHash, err := helper.EncodeUopHash(cops, entrypointUopHash, validationData)
	if err != nil {
		panic(err)
	}

	// Generate EOA signature ECDSA-Secp256k1
	eoaSignature, err := personalSign(hexutil.Encode(uopHash[:]), eoaPrivateKey)
	if err != nil {
		panic(fmt.Errorf("failed to sign message: %v", err))
	}
	log.Info().Msgf("eoaSignature: %s", hexutil.Encode(eoaSignature))
	// Verify EOA signature
	recoverAddr, err := helper.RecoverAddress(cops, eoaSignature, uopHash)
	if err != nil {
		panic(fmt.Errorf("failed to verify EOA signature: %v", err))
	}
	if recoverAddr != tops.From {
		panic(fmt.Sprintf("eoa signature verification failed: expected %s, got %s", tops.From.Hex(), recoverAddr.Hex()))
	}

	// Generate Passkey Signature ECDSA-Secp256r1 / P-256
	clientDataJSONPre := `{"type":"webauthn.get","challenge":"`
	clientDataJSONPost := `","origin":"http://localhost:8000","crossOrigin":false}`
	clientJson, err := helper.GetClientJson(
		cops,
		clientDataJSONPre,
		clientDataJSONPost,
		uopHash,
	)
	if err != nil {
		panic(fmt.Errorf("failed to get client json: %v", err))
	}
	passkeyMsgHash := clientJson.MessageHash
	passkeyPrivateKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     passkeyPubX,
			Y:     passkeyPubY,
		},
		D: passkeyPrivKey,
	}
	passkeySignature, err := passkeyPrivateKey.Sign(rand.Reader, passkeyMsgHash[:], nil)
	if err != nil {
		panic(fmt.Errorf("failed to sign message: %v", err))
	}
	r := new(big.Int).SetBytes(passkeySignature[:32])
	s := new(big.Int).SetBytes(passkeySignature[32:])
	// Verify passkey signature
	/// 0 : PRECOMPILED_VERIFIER, helper validation passes, but handleops fails.
	/// 1 : DAIMO_VERIFIER success
	/// 2 : ELLIPTIC_CURVE success(local fork node)
	verifyType := uint8(0)
	isVerifySuccess, _, err := helper.PasskeyVerify(
		cops,
		passkeyMsgHash,
		r,
		s,
		passkeyPubX,
		passkeyPubY,
		verifyType,
		clientJson.ClientDataJSON,
	)
	if err != nil || !isVerifySuccess {
		panic(fmt.Errorf("passkey verification failed: %v", err))
	}

	// Encode passkey signature
	passkeySig, err := helper.EncodePasskeySig(
		cops,
		r,
		s,
		verifyType,
		clientJson.ClientDataJSON,
	)
	if err != nil {
		panic(fmt.Errorf("failed to encode passkey signature: %v", err))
	}

	// Get final signature
	signature, err := helper.GetSignature2(
		cops,
		passkeyPubX,
		passkeyPubY,
		passkeySig,
		eoaSignature,
		validationData,
	)
	if err != nil {
		panic(fmt.Errorf("failed to get final signature: %v", err))
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

// func smartAccountInstallRecoveryModule(module common.Address, data []byte) ([]byte, error) {
// 	calldataStr, err := polyabi.AbiEncode("installRecoveryModule(address,bytes)", []string{module.String(), string(data)})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return []byte(calldataStr), nil
// }

// func smartAccountInstallFallbackModule(id int64, module common.Address, data []byte) ([]byte, error) {
// 	calldataStr, err := polyabi.AbiEncode("installModule(uint256,address,bytes)", []string{strconv.FormatInt(id, 10), module.String(), string(data)})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return []byte(calldataStr), nil
// }


// Source: https://github.com/etaaa/Golang-Ethereum-Personal-Sign/
func personalSign(message string, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	signatureBytes[64] += 27
	return signatureBytes, nil
}