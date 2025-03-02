package erc4337loadtest

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/0xPolygon/polygon-cli/bindings/4337/entryPoint/core/entrypoint"
	"github.com/0xPolygon/polygon-cli/bindings/4337/payableaccount"
	"github.com/0xPolygon/polygon-cli/bindings/4337/test/helper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"math/big"
	mRand "math/rand"
	"strings"
	"time"
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

var (
	passkeyPubX, _    = new(big.Int).SetString("640c5cacef387563d0b105c7724c45ee19f8a952cb583de494a6a7ce5ed16760", 16)
	passkeyPubY, _    = new(big.Int).SetString("142b33cbf8255e9f0628ab9e250e179a3e7e8e24e0a2a4340f0b9fdeb29a1b48", 16)
	passkeyPrivKey, _ = new(big.Int).SetString("42d2bd030a8a71ff2f9043adcfb46138a5d87287cefff37d18a638f956c33449", 16)

	salt     *big.Int
	modeType [32]byte // 0x0000000000000000000000000000000000000000000000000000000000000000
)

func init() {
	salt = big.NewInt(2)
	// Hardcoded
	//     const modeType = mode.encodeModeType(
	//   mode.CallType.Batch,
	//   mode.ExecType.Default,
	//   mode.ModeSelector.Default,
	//   "0x"
	// );
	modeType = [32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
}

func GenerateUops(
	client *ethclient.Client,
	ctx context.Context,
	tops *bind.TransactOpts,
	cops *bind.CallOpts,
	eoaPrivateKey *ecdsa.PrivateKey,
	cfg *ERC4337Config,
	initCode []byte,
	batchSize uint32,
) ([]entrypoint.PackedUserOperation, error) {
	if batchSize == 0 {
		return nil, fmt.Errorf("no uops to generate")
	}

	// Check if account needs initialization
	code, err := client.CodeAt(ctx, cfg.Sender, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get code at address: %v", err)
	}

	//uPrivateKey, err := ethcrypto.HexToECDSA("42d2bd030a8a71ff2f9043adcfb46138a5d87287cefff37d18a638f956c33449")
	//if err != nil {
	//	log.Info().Err(err).Msg("Failed to get private key")
	//	return nil, err
	//}
	//uAddress := ethcrypto.PubkeyToAddress(uPrivateKey.PublicKey)
	userOps := make([]entrypoint.PackedUserOperation, 0, batchSize)

	for i := uint32(0); i < batchSize; i++ {
		// Use a predefined calldata for ERC20 token operations (approve and transfer 1 token)
		// This matches the calldata from the TypeScript implementation
		mRand.Seed(time.Now().UnixNano())
		randomIndex := mRand.Intn(len(cfg.CallDataList))
		rawCallData := cfg.CallDataList[randomIndex]
		calldata, err := hex.DecodeString(strings.TrimPrefix(rawCallData, "0x"))
		log.Info().Msg("calldata: " + rawCallData)
		if err != nil {
			return nil, fmt.Errorf("failed to generate random transfer calldata: %v", err)
		}
		// Create base UserOperation
		userOp := PackUserOp(UserOperation{
			Sender: cfg.Sender,
			// it is safe to use the same nonce for all UOPs, since uop nonce validation is disabled
			Nonce: tops.Nonce,
			InitCode: func() []byte {
				if code == nil {
					return initCode
				}
				return nil
			}(),
			CallData:             calldata,
			CallGasLimit:         big.NewInt(1000000),
			VerificationGasLimit: big.NewInt(20000000),
			PreVerificationGas:   big.NewInt(0),
			MaxFeePerGas:         big.NewInt(1),
			MaxPriorityFeePerGas: big.NewInt(1),
			PaymasterAndData:     nil,
			Signature:            nil,
		})

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
		Sender:             userOp.Sender,
		Nonce:              userOp.Nonce,
		CallData:           userOp.CallData,
		AccountGasLimits:   accountGasLimits,
		InitCode:           userOp.InitCode,
		PreVerificationGas: userOp.PreVerificationGas,
		GasFees:            gasFees,
		PaymasterAndData:   userOp.PaymasterAndData,
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
	// Calculate expiration time (1 day from now)
	expireTime := big.NewInt(time.Now().Unix() + 86400)
	validationData, err := helper.GetValidationData(cops, expireTime)
	if err != nil {
		return nil, err
	}
	// Verify validation date
	_, outOfTimeRange, err := helper.CheckValidationDate(cops, validationData)
	if err != nil || outOfTimeRange {
		return nil, fmt.Errorf("validation date check failed: %v", err)
	}

	// Get UserOp hash from EntryPoint
	entrypointUopHash, err := entryPoint.GetUserOpHash(cops, userOp)
	if err != nil {
		return nil, err
	}
	// Encode UOP hash
	uopHash, err := helper.EncodeUopHash(cops, entrypointUopHash, validationData)
	if err != nil {
		return nil, err
	}

	// Generate EOA signature ECDSA-Secp256k1
	eoaSignature, err := personalSign(uopHash[:], eoaPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %v", err)
	}
	// Verify EOA signature
	recoverAddr, err := helper.RecoverAddress(cops, eoaSignature, uopHash)
	if err != nil {
		return nil, fmt.Errorf("failed to verify EOA signature: %v", err)
	}
	if recoverAddr != tops.From {
		return nil, fmt.Errorf("eoa recover address failed: expected %s, got %s", tops.From.Hex(), recoverAddr.Hex())
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
		return nil, fmt.Errorf("failed to get client json: %v", err)
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
	r, s, err := ecdsa.Sign(rand.Reader, passkeyPrivateKey, passkeyMsgHash[:])
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %v", err)
	}

	// Normalize s value to be in lower half of curve order for compatibility
	nDiv2 := new(big.Int)
	nDiv2.SetString("57896044605178124381348723474703786764998477612067880171211129530534256022184", 10)
	if s.Cmp(nDiv2) > 0 {
		n := new(big.Int)
		n.SetString("115792089210356248762697446949407573529996955224135760342422259061068512044369", 10)
		s = new(big.Int).Sub(n, s)
	}

	if !ecdsa.Verify(&passkeyPrivateKey.PublicKey, passkeyMsgHash[:], r, s) {
		return nil, fmt.Errorf("failed to verify passkey signature")
	}
	// Verify passkey signature
	/// 0 : PRECOMPILED_VERIFIER, helper validation passes, but handleops fails.
	/// 1 : DAIMO_VERIFIER success
	/// 2 : ELLIPTIC_CURVE success(local fork node)
	verifyType := uint8(0)
	isVerifySuccess, _, err := helper.PasskeyVerify(
		cops,
		uopHash,
		r,
		s,
		passkeyPubX,
		passkeyPubY,
		verifyType,
		clientJson.ClientDataJSON,
	)
	if err != nil || !isVerifySuccess {
		return nil, fmt.Errorf("passkey verification failed: %v", err)
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
		return nil, fmt.Errorf("failed to encode passkey signature: %v", err)
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
		return nil, fmt.Errorf("failed to get final signature: %v", err)
	}

	return signature, nil
}

// Source: https://github.com/etaaa/Golang-Ethereum-Personal-Sign/
func personalSign(msgHash []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msgHash), msgHash)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	signatureBytes[64] += 27
	return signatureBytes, nil
}

// Tranfer to random address with random amount
func getRandomTransferCalldata() ([]byte, error) {
	accountAbi, err := payableaccount.PayableAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	// Generate random address
	randomAddrBytes := make([]byte, 20) // Ethereum addresses are 20 bytes
	_, err = rand.Read(randomAddrBytes)
	if err != nil {
		return nil, err
	}
	randomAddr := common.BytesToAddress(randomAddrBytes)

	// Generate random tranfer amount between [0, 10wei)
	randomAmount, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		return nil, err
	}

	// Pack the parameters
	executeCalldata := append(randomAddr.Bytes(), common.LeftPadBytes(randomAmount.Bytes(), 32)...)
	executeCalldata = append(executeCalldata, []byte{}...) // Empty bytes
	calldata, err := accountAbi.Pack("execute", modeType, executeCalldata)
	if err != nil {
		return nil, err
	}

	return calldata, nil
}

// Tranfer to random address with random amount
//func getRandomTransferERC20Calldata(cfg *ERC4337Config) ([]byte, error) {
//	erc20ABI := testtoken20.TestToken20MetaData.ABI
//	payABI := pay.PayMetaData.ABI
//	payableABI := payableaccount.PayableAccountMetaData.ABI
//	// Parse the ABIs
//	erc20, err := abi.JSON(strings.NewReader(erc20ABI))
//	if err != nil {
//		return nil, fmt.Errorf("failed to parse ERC20 ABI: %v", err)
//	}
//	pay, err := abi.JSON(strings.NewReader(payABI))
//	if err != nil {
//		return nil, fmt.Errorf("failed to parse Pay ABI: %v", err)
//	}
//	payable, err := abi.JSON(strings.NewReader(payableABI))
//	if err != nil {
//		return nil, fmt.Errorf("failed to parse SmartAccount ABI: %v", err)
//	}
//
//	approveData, err := erc20.Pack("approve", randomAddr, randomAmount)
//	if err != nil {
//		return nil, fmt.Errorf("failed to encode approve data: %v", err)
//	}
//
//	return calldata, nil
//}
