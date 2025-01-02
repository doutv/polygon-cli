package erc4337loadtest

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/0xPolygon/polygon-cli/bindings/4337/accountfactory"
	"github.com/0xPolygon/polygon-cli/bindings/4337/entryPoint/core/entrypoint"
	"github.com/0xPolygon/polygon-cli/bindings/4337/payableaccount"
	"github.com/0xPolygon/polygon-cli/bindings/4337/test/helper"
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

var passkeyPubX, _ = new(big.Int).SetString("640c5cacef387563d0b105c7724c45ee19f8a952cb583de494a6a7ce5ed16760", 16)
var passkeyPubY, _ = new(big.Int).SetString("142b33cbf8255e9f0628ab9e250e179a3e7e8e24e0a2a4340f0b9fdeb29a1b48", 16)
var passkeyPrivKey, _ = new(big.Int).SetString("42d2bd030a8a71ff2f9043adcfb46138a5d87287cefff37d18a638f956c33449", 16)

var salt *big.Int
var modeType [32]byte
var emptyCalldata []byte

func init() {
	salt = big.NewInt(2)
	// empty execution calldata
	accountAbi, err := payableaccount.PayableAccountMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	emptyCalldata, err = accountAbi.Pack("execute", modeType, []byte{})
	if err != nil {
		panic(err)
	}
}

func generateInitcode(
	tops *bind.TransactOpts,
	cops *bind.CallOpts,
	cfg *ERC4337Config,
	salt *big.Int,
) (initCode []byte, err error) {
	accountAbi, err := payableaccount.PayableAccountMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	installRecoveryModuleCalldata, err := accountAbi.Pack("installRecoveryModule", tops.From, []byte{})
	if err != nil {
		panic(err)
	}
	installFallbackModuleCalldata, err := accountAbi.Pack("installModule", big.NewInt(3), cfg.TokenReceiver.Address, []byte{})
	if err != nil {
		panic(err)
	}

	// Create the initcode
	initializer0, _, err := cfg.Helper.Contract.GetAccountInitializer2(
		cops,
		passkeyPubX,
		passkeyPubY,
		cfg.Validator.Address,
		tops.From,
		tops.From,
		installRecoveryModuleCalldata,
		installFallbackModuleCalldata,
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
	return initcode, nil
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
) ([]entrypoint.IEntryPointUserOpsPerAggregator, error) {
	if batchSize == 0 {
		return nil, fmt.Errorf("no uops to generate")
	}

	// Check if account needs initialization
	code, err := client.CodeAt(ctx, cfg.Sender, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get code at address: %v", err)
	}

	userOps := make([]entrypoint.PackedUserOperation, 0, batchSize)

	for i := uint32(0); i < batchSize; i++ {
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
			CallData:             emptyCalldata,
			CallGasLimit:         big.NewInt(100000),
			VerificationGasLimit: big.NewInt(2000000),
			PreVerificationGas:   big.NewInt(0),
			MaxFeePerGas:         big.NewInt(1e9),
			MaxPriorityFeePerGas: big.NewInt(1e9),
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

	proof := getFixedZKProof()
	encodedProof, err := encodeZKProof(proof)
	if err != nil {
		panic(err)
	}

	aggUserOps := []entrypoint.IEntryPointUserOpsPerAggregator{
		{
			UserOps:    userOps,
			Aggregator: cfg.Aggregator.Address,
			Signature:  encodedProof,
		},
	}
	return aggUserOps, nil
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
	eoaSignature, err := personalSign(uopHash[:], eoaPrivateKey)
	if err != nil {
		panic(fmt.Errorf("failed to sign message: %v", err))
	}
	// Verify EOA signature
	recoverAddr, err := helper.RecoverAddress(cops, eoaSignature, uopHash)
	if err != nil {
		panic(fmt.Errorf("failed to verify EOA signature: %v", err))
	}
	if recoverAddr != tops.From {
		panic(fmt.Sprintf("eoa recover address failed: expected %s, got %s", tops.From.Hex(), recoverAddr.Hex()))
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
	r, s, err := ecdsa.Sign(rand.Reader, passkeyPrivateKey, passkeyMsgHash[:])
	if err != nil {
		panic(fmt.Errorf("failed to sign message: %v", err))
	}
	if !ecdsa.Verify(&passkeyPrivateKey.PublicKey, passkeyMsgHash[:], r, s) {
		panic(fmt.Errorf("failed to verify passkey signature"))
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

// ZKProof represents a fixed zero-knowledge proof structure for testing
type ZKProof struct {
	Proof          [8]*big.Int
	Commitments    [2]*big.Int
	CommitmentPok  [2]*big.Int
}

// getFixedZKProof returns a hardcoded ZK proof for testing
func getFixedZKProof() *ZKProof {
	proof := [8]*big.Int{
		hexToBigInt("0x068466E96FE3E43A47A45A219D48EBFFC731102FBF921BB681D929F8CCD7DCCD"),
		hexToBigInt("0x18CA295C5AA34118E0D51A10514B1FEBDDD32C5655296B434AA8731A324081EA"),
		hexToBigInt("0x124C3E2CBC52722EE1827E22B984BEC788D57BED83D02FB45E88A5F8DF402FC9"),
		hexToBigInt("0x224878AB963945288E9EC5305622E6E9E4DDA3B8B859D0630B6095D6D67755DA"),
		hexToBigInt("0x1AC0CC9A2C98ECB5DEDE94825D0FF08B8767E9B5FE13C843D16F423A18586484"),
		hexToBigInt("0x14FF00434FB50F3716B5F2011C341175EEF7C0E8E478E9FAC850D10E63E97DB9"),
		hexToBigInt("0x01F608E7E6620BDED6A2472EC329D0AFFE52ECD5E4AF62299476EEF15DD769B4"),
		hexToBigInt("0x23E4C88D2AC7F2062AA77082A4BD9967A69590BF61A0AB657EA02EA35A10A97B"),
	}
	
	commitments := [2]*big.Int{
		hexToBigInt("0x032afc4fa24014e503d42052a52ff814628c7ee24cf555f1c08f09b365f98d42"),
		hexToBigInt("0x157a2732c40aba0a752ba9b09c236b146f8166763dc6e596c021ede0650b034a"),
	}
	
	commitmentPok := [2]*big.Int{
		hexToBigInt("0x2d73a3ae1e7d1365ae44c36204fb9280dedcc250db8d4f24858a8969ed542b1d"),
		hexToBigInt("0x228f0592b8f77f4d8ad72ff574af1b6e1afdcad2324cbbf1280d45a60a22226d"),
	}

	return &ZKProof{
		Proof:         proof,
		Commitments:   commitments,
		CommitmentPok: commitmentPok,
	}
}

// encodeZKProof encodes the ZK proof into ABI-encoded bytes
func encodeZKProof(zkp *ZKProof) ([]byte, error) {
	uint256Ty, _ := abi.NewType("uint256", "", nil)
	
	arguments := abi.Arguments{
		{Type: arrayType(uint256Ty, 8)},
		{Type: arrayType(uint256Ty, 2)},
		{Type: arrayType(uint256Ty, 2)},
	}

	return arguments.Pack(
		zkp.Proof,
		zkp.Commitments,
		zkp.CommitmentPok,
	)
}

// Helper function to create fixed-size array types
func arrayType(t abi.Type, size int) abi.Type {
	return abi.Type{
		T:          abi.ArrayTy,
		Size:       size,
		Elem:       &t,
	}
}

// Helper function to convert hex string to big.Int
func hexToBigInt(hex string) *big.Int {
	n := new(big.Int)
	n.SetString(hex[2:], 16) // Remove "0x" prefix
	return n
}
