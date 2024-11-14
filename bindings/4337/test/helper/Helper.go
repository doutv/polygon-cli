// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package helper

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

// ValidationData is an auto generated low-level Go binding around an user-defined struct.
type ValidationData struct {
	Aggregator common.Address
	ValidAfter *big.Int
	ValidUntil *big.Int
}

// HelperMetaData contains all meta data concerning the Helper contract.
var HelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"okxSignatureData\",\"type\":\"bytes\"}],\"name\":\"PasskeyFormatDemo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"name\":\"checkValidationDate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"computePredictionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"dealData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"okxSignatureData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"}],\"name\":\"encodePacked\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"verifyType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"clientDataJSON\",\"type\":\"string\"}],\"name\":\"encodePasskeySig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uopHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"}],\"name\":\"encodeUopHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pubKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyY\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eoaSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"targetData1\",\"type\":\"bytes\"}],\"name\":\"getAccountInitializer2\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlocktimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientDataJSONPre\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"clientDataJSONPost\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"name\":\"getClientJson\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"clientDataJSON\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"config\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eoaSigner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"}],\"name\":\"getConfigSetSignerHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_initializer\",\"type\":\"bytes\"}],\"name\":\"getFactoryCreateAccountHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"getFactorySig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getHash1\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pubKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyY\",\"type\":\"uint256\"}],\"name\":\"getPubkeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pubKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyY\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"passkeySig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"eoaSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"name\":\"getSignature2\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"getSignature2Decode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHashWithEntryPoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"name\":\"getValidationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"okxSignatureData\",\"type\":\"bytes\"}],\"name\":\"mockRecover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validAfter\",\"type\":\"uint256\"}],\"name\":\"packValidationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"name\":\"parseValidationData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"}],\"internalType\":\"structValidationData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"okxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"verifyType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"clientDataJSON\",\"type\":\"string\"}],\"name\":\"passkeyVerify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"verifier\",\"type\":\"uint8\"}],\"name\":\"passkeyVerify1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"recoverAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"recoverAddress1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"splitSig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"challenge\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"authenticatorData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"clientDataJSON\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"verifier\",\"type\":\"uint8\"}],\"name\":\"verifyPasskeySignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608080604052346015576126ac908161001b8239f35b600080fdfe60806040908082526004908136101561001757600080fd5b60009060e08235811c918263019b696e14611635575081630aeaa733146115825781630f4d65e1146115605781630f559e0614611523578163179e90fd1461148f57816318c82e081461145957816322de96b1146114385781634f2a4c7b146113ef57816353530b18146112c457816353ffc3b014610fa35781635b6beeb914610dc85781635db1bc8814610d74578163627041f514610ccc57816363e53f6114610c0357816372e8eeea14610b6c578163780dfbf614610a93578163899413f614610a6057816398a1941f146109955781639bffea2414610952578163a4b2282e14610902578163ba2f6e5e146108e7578163c3c51fc914610850578163c6e8a99e146107b1578163caeacff61461071d578163d7d0679f14610692578163e585b437146103f357508063ecaffa771461036e578063f0dc3df5146101f45763f29ad2161461016657600080fd5b346101e95760203660031901126101e9576001600160401b039180358381116101f05761019591369101611893565b81939101906060848303126101f05783358181116101ec57826101b9918601611878565b9260208501359182116101e957506101e592916101d7918501611878565b8451948594013591846118c0565b0390f35b80fd5b8380fd5b8280fd5b5091346101f05760603660031901126101f0576001600160401b03823581811161036a576102259036908501611878565b9260243591821161036a5761023c91369101611878565b916102626102ae835160209283916044358382015282815261025d81611783565b611c24565b85519687918361027b81850198898151938492016116aa565b830161028f825180938780850191016116aa565b016102a2825180938680850191016116aa565b01038087520185611805565b80856102cc6102bb61196c565b948651809281928a519283916116aa565b8101039060025afa156103605780856103218151946103118780518098876102fd81840196878151938492016116aa565b820190888201520386810189520187611805565b86518092819288519283916116aa565b8101039060025afa15610360576103589061034b95519284519687966060885260608801906116cd565b91868303908701526116cd565b918301520390f35b82513d86823e3d90fd5b8480fd5b509190346101f057816003193601126101f05780356001600160401b0381116101ec57602093603c6103aa6103e395946103da94369101611878565b917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152602435601c52206120a0565b90939193611fb7565b516001600160a01b039091168152f35b83858492346101ec57816003193601126101ec576024356001600160401b03811161036a576104259036908501611878565b9282519060209485808401528483016008905260609081840167686f6f6b4461746160c01b9052818452610458846117cf565b61046061196c565b86519561046c876117ea565b608687528887017f7b2274797065223a22776562617574686e2e676574222c226368616c6c656e6790528787017f65223a226354534253556f675f374d503155455177744b70794f474c6965534b90528387017f307043644f6b2d47396a514f6b5f45222c226f726967696e223a22687474703a9052608087017f2f2f6c6f63616c686f73743a38303030222c2263726f73734f726967696e223a905260a087016566616c73657d60d01b90528751998a9261010090818c8601526101208501610535916116cd565b98601f19998a868203018c87015261054c916116cd565b6001858801527f7fde2a6a50c8b06dfa8869b0c217665828fd6c9bbba6e1e16034bf9b9695899960808601527f21ede5fef42f27196a2f42167c8c4428a3f592b587d780c69eeac62ad5ddc44960a08601527f13e5e3c1d310f8515c2e28f499755c5f9433cb70ba1d7ecf4f276ff4b63a635260c08601527f19ae7d5e4f7341511fef9fc70c4e447b499505e1612dc2e792b272f9d83c6e4f938501939093528301520385810189526105ff9089611805565b85519788916106109189840161172f565b0384810188526106209088611805565b845191358683015285825261063482611783565b845196879382888601526080850161064b916116cd565b85858203018786015261065d916116cd565b90848483030190840152610670916116cd565b0390810184526106809084611805565b5191818392835282016101e5916116cd565b5050346101e957826003193601126101e957602435906001600160401b0382116101e957506106fa9261070e6106ce6101e59336908601611878565b9382516020958692838301903581528383526106e983611783565b8551988993518092868601906116aa565b82016102a2825180938680850191016116aa565b519282849384528301906116cd565b838584346101e95760803660031901126101e957610739611757565b90606435906001600160401b0382116101e957506107a761076260209561078993369101611878565b84516001600160a01b039094168685019081526080858701529392839160a08301906116cd565b6024356060830152604435608083015203601f198101835282611805565b5190209051908152f35b90503461084c573660031901126101e9576107ca6116f2565b9060c435906001600160401b0382116101e9575061082b6101e5926107f56108339336908701611878565b945a9561080061196c565b8851923560208401526020835261081683611783565b60843592606435926044359260243592611a1f565b925a90611916565b9251911515825260208201929092529081906040820190565b5080fd5b5050823461084c57608036600319011261084c5761086c611757565b906024356001600160a01b03818116929183900361036a5761088c61176d565b8251938260208601961686528385015216606083015260643560808301526080825260a08201938285106001600160401b038611176108d45750602094508390525190208152f35b634e487b7160e01b815260418652602490fd5b82853461084c578160031936011261084c5760209051428152f35b838584346101e95760203660031901126101e95750610923606092356119af565b9080519160018060a01b0381511683528160208201519165ffffffffffff809316602086015201511690820152f35b838584346101e95760c03660031901126101e9575061098c6020926109756116f2565b90608435906064359060443590602435903561210c565b90519015158152f35b5050823461084c5760a036600319011261084c576001600160401b036044358181116101ec576109c89036908601611878565b926064359182116101e957509061070e6109e86101e59336908701611878565b8251956109f487611783565b35865260209460243586880152610a10845197878901906118ee565b838752610a1c876117b4565b610a2c845192839288840161172f565b0390610a40601f1992838101835282611805565b610a548451978892608435918985016118c0565b03908101865285611805565b828585346101f057816003193601126101f05760209250815183810191358252602435838201528281526107a7816117b4565b505091346101f057602091826003193601126101ec5780356001600160401b0391828211610b6857610ac791369101611878565b80518101918382868501940312610b6857848201519184810151918211610b64570182603f82011215610b68578481015190610b0282611826565b93610b0f86519586611805565b828552858383010111610b6457958392610b38603c938899886103e39a6103da990191016116aa565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c52206120a0565b8680fd5b8580fd5b5050823461084c578060031936011261084c57610b87611757565b91602435906001600160401b0382116101e957506034610bb06101e595610bf193369101611878565b9383519485916bffffffffffffffffffffffff199060601b166020830152610be181518092602086860191016116aa565b8101036014810185520183611805565b519182916020835260208301906116cd565b828585346101f057816003193601126101f057356001600160a01b03811692908390036101e957506102df610c8d610c9960208095855190610c4783870183611805565b8582528282019561239887398651838101918252838152610c6781611783565b8751958693610c7e868601998a92519283916116aa565b840191518093868401906116aa565b01038084520182611805565b51902081518381019160ff60f81b83523060601b602183015260243560358301526055820152605581526107a7816117cf565b849150346101f05760803660031901126101f0576044359260ff84168094036101e957606435906001600160401b0382116101e9575093610bf191610d176101e59636908401611878565b91610d2061196c565b95610d4d610d3b875198899560c060208801528601906116cd565b601f19958686830301898701526116cd565b916001606085015235608084015260243560a084015260c083015203908101845283611805565b5050823461084c57602036600319011261084c5782356001600160401b0381116101f057606093610da791369101611878565b906020820151928482840151930151901a9181519384526020840152820152f35b838584346101e9576020928360031936011261084c5780356001600160401b0381116101f057610dfb9036908301611878565b908151809184907a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000080841015610f96575b50876d04ee2d6d415b85acef810000000080861015610f87575b5050662386f26fc1000080851015610f78575b506305f5e10080851015610f69575b5061271080851015610f5d575b50506064831015610f4f575b600a80931015610f47575b9060219160019281610e9c600186940161206e565b9750870101905b610f11575b505050506107a7603a84518093878201957f19457468657265756d205369676e6564204d6573736167653a0a0000000000008752610eee815180928b87870191016116aa565b8201610f02825180938b87850191016116aa565b0103601a810184520182611805565b600019019083906f181899199a1a9b1b9c1cb0b131b232b360811b8282061a835304918215610f4257919082610ea3565b610ea8565b600101610e87565b916064600291049201610e7c565b90930492018780610e70565b60089192940493019088610e63565b60109192940493019088610e54565b90919294049301908789610e41565b8304935086915088610e27565b84848492346101ec573660031901126101f057610fbe61176d565b916064359360018060a01b0380861680960361084c57608435958187168097036101f0576001600160401b039660a43588811161036a576110029036908801611893565b93909160c435998a11610b685761101e87939a36908a01611893565b939081519961102c8b611783565b358a526020996024358b82015282516110488c820180936118ee565b838152611054816117b4565b519020948a83519d8e97828901528a8589015261108a6060986fffffffffffffffffffffffffffffffff8a8201528981526117cf565b84519e8f9061109882611783565b600182528c5b8a8582106112a8575050906110bc916110b682611939565b52611939565b508451996110c98b6117b4565b60028b528b805b8a888210611277575050906110fc929161112e9751946110ef866117b4565b8986528501523691611841565b8b82015261110989611939565b5261111388611939565b50895193611120856117b4565b8452888b8501523691611841565b8782015261113b8561195c565b526111458461195c565b5085519760a089016080898b01528a5180915260c08a01908960c08260051b8d01019c019188905b82821061124a5750505050601f1995868a8c0301888b01528551808c528b8a80808301928460051b0101980192905b82821061120457505050501690870152608086015284900380820185526101e592916111fa916111cc9087611805565b82519463439fab9160e01b818701526024860152846111ee60448201886116cd565b03908101855284611805565b519283928361172f565b909192978b808f8d8a858f8f8d8591859361123c9860019c03018c5251908151168452878101518885015201519382015201906116cd565b9a019201920190929161119c565b9091929c8b808f8f936112699160019560bf19908303018752516116cd565b9f019201920190929161116d565b9092949395969780519261128a846117b4565b8084528584015282015282828d01015201918c918b8d9695946110d0565b919493959697509182849101015201918e8d928d96959461109e565b90508391346101e9576003199160603684011261084c576112e3611757565b92604435906001600160401b03938483116101e957828801916101208091853603011261084c576001600160a01b0393611320604482018561203c565b90818b51918237209560c4611338606484018761203c565b90818d5191823720928b61134f60e483018961203c565b80925191823720938c51998960208c0199351689528d6024840135908c015260608b015260808a0152608481013560a08a015260a481013560c08a015201359087015261010090818701528552840194848610908611176113dc5750602096508386528251902093610140830194855216610160820152610180602435910152606081526107a7816117cf565b634e487b7160e01b815260418852602490fd5b828585346101f05760203660031901126101f05782603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020955235601c52209051908152f35b8483346101e95760206103e3836103da61145136611707565b929092611f27565b5050346101e95760203660031901126101e957506114779035611bd5565b82516001600160a01b03909216825215156020820152f35b5050346101e9576101003660031901126101e9576001600160401b039180358381116101f0576114c29036908301611878565b906024358481116101ec576114da9036908301611878565b936044359081116101ec576114f191369101611878565b60e4359260ff841684036101e957506101e5926108339261082b925a9660c4359260a435926084359260643592611a1f565b82853461084c57606036600319011261084c57602090611541611757565b90516001600160a01b0390911660243560a01b1760443560d01b178152f35b505091346101f05760203660031901126101f0576020925051903560a01b8152f35b8483346101e957506111fa6101e561159936611707565b9095929491610a5484519661160d886115d1838c876020850191604193918352602083015260ff60f81b9060f81b1660408201520190565b03946115e5601f19968781018c528b611805565b61160188519a8b926020840152898084015260608301906116cd565b038581018a5289611805565b85519889936020850191604193918352602083015260ff60f81b9060f81b1660408201520190565b85908585346101e957836003193601126101e9575061165382611783565b3581526024356020820152815161166e6020820180936118ee565b82815261167a816117b4565b5190206101e582518260208201526020815261169581611783565b835193849384528060208501528301906116cd565b60005b8381106116bd5750506000910152565b81810151838201526020016116ad565b906020916116e6815180928185528580860191016116aa565b601f01601f1916010190565b60a4359060ff8216820361170257565b600080fd5b60809060031901126117025760043560ff811681036117025790602435906044359060643590565b9091611746611754936040845260408401906116cd565b9160208184039101526116cd565b90565b600435906001600160a01b038216820361170257565b604435906001600160a01b038216820361170257565b604081019081106001600160401b0382111761179e57604052565b634e487b7160e01b600052604160045260246000fd5b606081019081106001600160401b0382111761179e57604052565b608081019081106001600160401b0382111761179e57604052565b60c081019081106001600160401b0382111761179e57604052565b90601f801991011681019081106001600160401b0382111761179e57604052565b6001600160401b03811161179e57601f01601f191660200190565b92919261184d82611826565b9161185b6040519384611805565b829481845281830111611702578281602093846000960137010152565b9080601f830112156117025781602061175493359101611841565b9181601f84011215611702578235916001600160401b038311611702576020838186019501011161170257565b9392916118e9906118db6040936060885260608801906116cd565b9086820360208801526116cd565b930152565b6000915b600283106118ff57505050565b6001908251815260208091019201920191906118f2565b9190820391821161192357565b634e487b7160e01b600052601160045260246000fd5b8051156119465760200190565b634e487b7160e01b600052603260045260246000fd5b8051600110156119465760400190565b60405190611979826117b4565b602582527f49960de5880e8c687434170f6476605b8fe4aeb9a28632c7995cf3ba831d97636020830152601960f81b6040830152565b6000604080516119be816117b4565b828152826020820152015265ffffffffffff808260a01c168015611a07575b604051926119ea846117b4565b6001600160a01b038116845260d01c602084015216604082015290565b50806119dd565b908151811015611946570160200190565b959391979694929060019681519960258b108015611bae575b611afa9b50611ba5575b604092835192611a5184611783565b60158452611a7e8360209574113a3cb832911d113bb2b130baba34371733b2ba1160591b87820152611e47565b15611b9c575b82611a91611ae192611c24565b611adc602e885180936c1131b430b63632b733b2911d1160991b8a830152611ac2815180928c602d860191016116aa565b8101601160f91b602d82015203600e810184520182611805565b611ebc565b15611b93575b828085519d848f809651938492016116aa565b60009d808f9592869301039060025afa15611b8757611b55908251611b45868051809388611b3181840197888151938492016116aa565b820190898201520387810184520182611805565b85519283928392519283916116aa565b8101039060025afa15611b7e5750611b6f9596975161210c565b9081611b79575090565b905090565b513d89823e3d90fd5b508251903d90823e3d90fd5b60009950611ae7565b60009a50611a84565b60009850611a42565b50996020101561194657611afa99611bcf60ff60f81b604085015116611dfc565b15611a38565b8015611c1b57611be4906119af565b65ffffffffffff806040830151164211908115611c0b575b5090516001600160a01b031691565b9050602082015116421038611bfc565b50600090600090565b604051611c8991611c34826117b4565b604082527f4142434445464748494a4b4c4d4e4f505152535455565758595a61626364656660208301527f6768696a6b6c6d6e6f707172737475767778797a303132333435363738392b2f6040830152612284565b60008151600281119081611dcc575b5015611d645750611cac60025b8251611916565b611cb581611826565b91611cc36040519384611805565b818352601f19611cd283611826565b0136602085013760005b828110611ce95750505090565b6001906001600160f81b0319602b60f81b81611d058487611a0e565b511603611d215750602d611d198287611a0e565b535b01611cdc565b602f60f81b81611d318487611a0e565b511603611d4b5750605f611d458287611a0e565b53611d1b565b611d558285611a0e565b511660001a611d458287611a0e565b8151600181119081611d88575b50611d80575b611cac90611ca5565b506001611d77565b600019810191508111611db857603d60f81b906001600160f81b031990611daf9085611a0e565b51161438611d71565b634e487b7160e01b82526011600452602482fd5b600119810191508111611db857603d60f81b906001600160f81b031990611df39085611a0e565b51161438611c98565b6001600160f81b0319600160f81b821601611e4157601f60fb1b600160fb1b821601611e29575b50600190565b600160fc1b90811614611e3c5738611e23565b600090565b50600090565b80519082519260005b838110611e61575050505050600190565b6001908082018083116119235786811090811591611e90575b50611e855701611e50565b505050505050600090565b90506001600160f81b0319611eb281611ea98589611a0e565b51169286611a0e565b5116141538611e7a565b805191805160005b848110611ed5575050505050600190565b60178181018091116119235782811090811591611f04575b50611efa57600101611ec4565b5050505050600090565b90506001600160f81b0319611f1d81611ea98589611a0e565b5116141538611eed565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411611fab57926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa15611f9f5780516001600160a01b03811615611f9657918190565b50809160019190565b604051903d90823e3d90fd5b50505060009160039190565b60048110156120265780611fc9575050565b60018103611fe35760405163f645eedf60e01b8152600490fd5b600281036120045760405163fce698f760e01b815260048101839052602490fd5b60031461200e5750565b602490604051906335e2f38360e21b82526004820152fd5b634e487b7160e01b600052602160045260246000fd5b903590601e198136030182121561170257018035906001600160401b0382116117025760200191813603831361170257565b9061207882611826565b6120856040519182611805565b8281528092612096601f1991611826565b0190602036910137565b81519190604183036120d1576120ca92506020820151906060604084015193015160001a90611f27565b9192909190565b505060009160029190565b3d15612107573d906120ed82611826565b916120fb6040519384611805565b82523d6000602084013e565b606090565b94939290919460409182519460209760ff89880198858a5287878a01528260608a01528460808a01528360a08a015260a08952612148896117ea565b168061219b5750505050505050600091829151906101005afa9061216a6120dc565b91158015612192575b61218b57808280518101031261170257015160011490565b5050600090565b50815115612173565b6001036121e4575050505050506000918291519073c2b78104907f722dabac4c69f826a522b2754de45afa506121cf6120dc565b90808280518101031261170257015160011490565b9091929496508361224194965051966121fc88611783565b87528787015284519161220e83611783565b825286820152612237845195878701946304e960d760e01b8652602488015260448701906118ee565b60848501906118ee565b60a4835260e08301928084106001600160401b0385111761179e576000938493525190731704af085ca24d18100d3c98d87c96f4dd5557c65afa506121cf6120dc565b9190918051156123715780519260029160028501809511611923576003948590046001600160fe1b0381168103611923576122c49060029694961b61206e565b926020840192829183518401976020890192835194600085525b8a81106123245750505050600393949596505251068060011461231157600214612306575090565b603d90600019015390565b50603d9081600019820153600119015390565b836004919b989b019a8b51600190603f9082828260121c16870101518453828282600c1c16870101518385015382828260061c1687010151878501531684010151858201530196996122de565b509050604051602081018181106001600160401b0382111761179e57604052600081529056fe60a034606357601f6102df38819003918201601f19168301916001600160401b03831184841017606857808492602094604052833981010312606357516001600160a01b0381168103606357608052604051610260908161007f8239608051815050f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe60806040526004361015610024575b361561001f5734156101eb57600080fd5b6101eb565b6000803560e01c63d1f578941461003b575061000e565b346100af5760403660031901126100af576004356001600160a01b03811681036100ab576024359067ffffffffffffffff908183116100a757366023840112156100a75782600401359182116100a75736602483850101116100a75760246100a4930190610111565b80f35b8380fd5b5080fd5b80fd5b634e487b7160e01b600052604160045260246000fd5b6020808252825181830181905290939260005b8281106100fd57505060409293506000838284010152601f8019910116010190565b8181018601518482016040015285016100db565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8054929390926001600160a01b03166101d95760009382859455816040519283928337810184815203915af43d156101d15767ffffffffffffffff903d8281116101cc5760405192601f8201601f19908116603f01168401908111848210176101cc5760405282523d6000602084013e5b156101ab5750565b604051633018224d60e21b81529081906101c890600483016100c8565b0390fd5b6100b2565b6060906101a3565b60405163c28d69c760e01b8152600490fd5b600036818037808036817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc545af43d82803e15610226573d90f35b3d90fdfea264697066735822122058f0f8eb4a924f70bd7688d8b719e0c015fc0aae4f06b8a7f34ce91ee0d9998864736f6c63430008190033a2646970667358221220b4842d9104b5db82737b8c2d58936c1237cc10a2409835f354ca829bfa82dacf64736f6c63430008190033",
}

// HelperABI is the input ABI used to generate the binding from.
// Deprecated: Use HelperMetaData.ABI instead.
var HelperABI = HelperMetaData.ABI

// HelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HelperMetaData.Bin instead.
var HelperBin = HelperMetaData.Bin

// DeployHelper deploys a new Ethereum contract, binding an instance of Helper to it.
func DeployHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Helper, error) {
	parsed, err := HelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Helper{HelperCaller: HelperCaller{contract: contract}, HelperTransactor: HelperTransactor{contract: contract}, HelperFilterer: HelperFilterer{contract: contract}}, nil
}

// Helper is an auto generated Go binding around an Ethereum contract.
type Helper struct {
	HelperCaller     // Read-only binding to the contract
	HelperTransactor // Write-only binding to the contract
	HelperFilterer   // Log filterer for contract events
}

// HelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelperSession struct {
	Contract     *Helper           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelperCallerSession struct {
	Contract *HelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelperTransactorSession struct {
	Contract     *HelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelperRaw struct {
	Contract *Helper // Generic contract binding to access the raw methods on
}

// HelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelperCallerRaw struct {
	Contract *HelperCaller // Generic read-only contract binding to access the raw methods on
}

// HelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelperTransactorRaw struct {
	Contract *HelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelper creates a new instance of Helper, bound to a specific deployed contract.
func NewHelper(address common.Address, backend bind.ContractBackend) (*Helper, error) {
	contract, err := bindHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Helper{HelperCaller: HelperCaller{contract: contract}, HelperTransactor: HelperTransactor{contract: contract}, HelperFilterer: HelperFilterer{contract: contract}}, nil
}

// NewHelperCaller creates a new read-only instance of Helper, bound to a specific deployed contract.
func NewHelperCaller(address common.Address, caller bind.ContractCaller) (*HelperCaller, error) {
	contract, err := bindHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelperCaller{contract: contract}, nil
}

// NewHelperTransactor creates a new write-only instance of Helper, bound to a specific deployed contract.
func NewHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*HelperTransactor, error) {
	contract, err := bindHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelperTransactor{contract: contract}, nil
}

// NewHelperFilterer creates a new log filterer instance of Helper, bound to a specific deployed contract.
func NewHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*HelperFilterer, error) {
	contract, err := bindHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelperFilterer{contract: contract}, nil
}

// bindHelper binds a generic wrapper to an already deployed contract.
func bindHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helper *HelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helper.Contract.HelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helper *HelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helper.Contract.HelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helper *HelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helper.Contract.HelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helper *HelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helper *HelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helper *HelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helper.Contract.contract.Transact(opts, method, params...)
}

// PasskeyFormatDemo is a free data retrieval call binding the contract method 0xe585b437.
//
// Solidity: function PasskeyFormatDemo(bytes32 uid, bytes okxSignatureData) pure returns(bytes signature)
func (_Helper *HelperCaller) PasskeyFormatDemo(opts *bind.CallOpts, uid [32]byte, okxSignatureData []byte) ([]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "PasskeyFormatDemo", uid, okxSignatureData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PasskeyFormatDemo is a free data retrieval call binding the contract method 0xe585b437.
//
// Solidity: function PasskeyFormatDemo(bytes32 uid, bytes okxSignatureData) pure returns(bytes signature)
func (_Helper *HelperSession) PasskeyFormatDemo(uid [32]byte, okxSignatureData []byte) ([]byte, error) {
	return _Helper.Contract.PasskeyFormatDemo(&_Helper.CallOpts, uid, okxSignatureData)
}

// PasskeyFormatDemo is a free data retrieval call binding the contract method 0xe585b437.
//
// Solidity: function PasskeyFormatDemo(bytes32 uid, bytes okxSignatureData) pure returns(bytes signature)
func (_Helper *HelperCallerSession) PasskeyFormatDemo(uid [32]byte, okxSignatureData []byte) ([]byte, error) {
	return _Helper.Contract.PasskeyFormatDemo(&_Helper.CallOpts, uid, okxSignatureData)
}

// CheckValidationDate is a free data retrieval call binding the contract method 0x18c82e08.
//
// Solidity: function checkValidationDate(uint256 validationData) view returns(address, bool)
func (_Helper *HelperCaller) CheckValidationDate(opts *bind.CallOpts, validationData *big.Int) (common.Address, bool, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "checkValidationDate", validationData)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// CheckValidationDate is a free data retrieval call binding the contract method 0x18c82e08.
//
// Solidity: function checkValidationDate(uint256 validationData) view returns(address, bool)
func (_Helper *HelperSession) CheckValidationDate(validationData *big.Int) (common.Address, bool, error) {
	return _Helper.Contract.CheckValidationDate(&_Helper.CallOpts, validationData)
}

// CheckValidationDate is a free data retrieval call binding the contract method 0x18c82e08.
//
// Solidity: function checkValidationDate(uint256 validationData) view returns(address, bool)
func (_Helper *HelperCallerSession) CheckValidationDate(validationData *big.Int) (common.Address, bool, error) {
	return _Helper.Contract.CheckValidationDate(&_Helper.CallOpts, validationData)
}

// ComputePredictionHash is a free data retrieval call binding the contract method 0x63e53f61.
//
// Solidity: function computePredictionHash(address _config, uint256 _salt) view returns(bytes32)
func (_Helper *HelperCaller) ComputePredictionHash(opts *bind.CallOpts, _config common.Address, _salt *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "computePredictionHash", _config, _salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputePredictionHash is a free data retrieval call binding the contract method 0x63e53f61.
//
// Solidity: function computePredictionHash(address _config, uint256 _salt) view returns(bytes32)
func (_Helper *HelperSession) ComputePredictionHash(_config common.Address, _salt *big.Int) ([32]byte, error) {
	return _Helper.Contract.ComputePredictionHash(&_Helper.CallOpts, _config, _salt)
}

// ComputePredictionHash is a free data retrieval call binding the contract method 0x63e53f61.
//
// Solidity: function computePredictionHash(address _config, uint256 _salt) view returns(bytes32)
func (_Helper *HelperCallerSession) ComputePredictionHash(_config common.Address, _salt *big.Int) ([32]byte, error) {
	return _Helper.Contract.ComputePredictionHash(&_Helper.CallOpts, _config, _salt)
}

// DealData is a free data retrieval call binding the contract method 0x0aeaa733.
//
// Solidity: function dealData(uint8 v, bytes32 r, bytes32 s, bytes32 dataHash) pure returns(bytes okxSignatureData, bytes sig)
func (_Helper *HelperCaller) DealData(opts *bind.CallOpts, v uint8, r [32]byte, s [32]byte, dataHash [32]byte) (struct {
	OkxSignatureData []byte
	Sig              []byte
}, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "dealData", v, r, s, dataHash)

	outstruct := new(struct {
		OkxSignatureData []byte
		Sig              []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OkxSignatureData = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Sig = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// DealData is a free data retrieval call binding the contract method 0x0aeaa733.
//
// Solidity: function dealData(uint8 v, bytes32 r, bytes32 s, bytes32 dataHash) pure returns(bytes okxSignatureData, bytes sig)
func (_Helper *HelperSession) DealData(v uint8, r [32]byte, s [32]byte, dataHash [32]byte) (struct {
	OkxSignatureData []byte
	Sig              []byte
}, error) {
	return _Helper.Contract.DealData(&_Helper.CallOpts, v, r, s, dataHash)
}

// DealData is a free data retrieval call binding the contract method 0x0aeaa733.
//
// Solidity: function dealData(uint8 v, bytes32 r, bytes32 s, bytes32 dataHash) pure returns(bytes okxSignatureData, bytes sig)
func (_Helper *HelperCallerSession) DealData(v uint8, r [32]byte, s [32]byte, dataHash [32]byte) (struct {
	OkxSignatureData []byte
	Sig              []byte
}, error) {
	return _Helper.Contract.DealData(&_Helper.CallOpts, v, r, s, dataHash)
}

// EncodePacked is a free data retrieval call binding the contract method 0x72e8eeea.
//
// Solidity: function encodePacked(address factory, bytes initializer) pure returns(bytes)
func (_Helper *HelperCaller) EncodePacked(opts *bind.CallOpts, factory common.Address, initializer []byte) ([]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "encodePacked", factory, initializer)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodePacked is a free data retrieval call binding the contract method 0x72e8eeea.
//
// Solidity: function encodePacked(address factory, bytes initializer) pure returns(bytes)
func (_Helper *HelperSession) EncodePacked(factory common.Address, initializer []byte) ([]byte, error) {
	return _Helper.Contract.EncodePacked(&_Helper.CallOpts, factory, initializer)
}

// EncodePacked is a free data retrieval call binding the contract method 0x72e8eeea.
//
// Solidity: function encodePacked(address factory, bytes initializer) pure returns(bytes)
func (_Helper *HelperCallerSession) EncodePacked(factory common.Address, initializer []byte) ([]byte, error) {
	return _Helper.Contract.EncodePacked(&_Helper.CallOpts, factory, initializer)
}

// EncodePasskeySig is a free data retrieval call binding the contract method 0x627041f5.
//
// Solidity: function encodePasskeySig(uint256 r, uint256 s, uint8 verifyType, string clientDataJSON) pure returns(bytes)
func (_Helper *HelperCaller) EncodePasskeySig(opts *bind.CallOpts, r *big.Int, s *big.Int, verifyType uint8, clientDataJSON string) ([]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "encodePasskeySig", r, s, verifyType, clientDataJSON)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodePasskeySig is a free data retrieval call binding the contract method 0x627041f5.
//
// Solidity: function encodePasskeySig(uint256 r, uint256 s, uint8 verifyType, string clientDataJSON) pure returns(bytes)
func (_Helper *HelperSession) EncodePasskeySig(r *big.Int, s *big.Int, verifyType uint8, clientDataJSON string) ([]byte, error) {
	return _Helper.Contract.EncodePasskeySig(&_Helper.CallOpts, r, s, verifyType, clientDataJSON)
}

// EncodePasskeySig is a free data retrieval call binding the contract method 0x627041f5.
//
// Solidity: function encodePasskeySig(uint256 r, uint256 s, uint8 verifyType, string clientDataJSON) pure returns(bytes)
func (_Helper *HelperCallerSession) EncodePasskeySig(r *big.Int, s *big.Int, verifyType uint8, clientDataJSON string) ([]byte, error) {
	return _Helper.Contract.EncodePasskeySig(&_Helper.CallOpts, r, s, verifyType, clientDataJSON)
}

// EncodeUopHash is a free data retrieval call binding the contract method 0x899413f6.
//
// Solidity: function encodeUopHash(bytes32 uopHash, uint256 expireTime) pure returns(bytes32)
func (_Helper *HelperCaller) EncodeUopHash(opts *bind.CallOpts, uopHash [32]byte, expireTime *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "encodeUopHash", uopHash, expireTime)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EncodeUopHash is a free data retrieval call binding the contract method 0x899413f6.
//
// Solidity: function encodeUopHash(bytes32 uopHash, uint256 expireTime) pure returns(bytes32)
func (_Helper *HelperSession) EncodeUopHash(uopHash [32]byte, expireTime *big.Int) ([32]byte, error) {
	return _Helper.Contract.EncodeUopHash(&_Helper.CallOpts, uopHash, expireTime)
}

// EncodeUopHash is a free data retrieval call binding the contract method 0x899413f6.
//
// Solidity: function encodeUopHash(bytes32 uopHash, uint256 expireTime) pure returns(bytes32)
func (_Helper *HelperCallerSession) EncodeUopHash(uopHash [32]byte, expireTime *big.Int) ([32]byte, error) {
	return _Helper.Contract.EncodeUopHash(&_Helper.CallOpts, uopHash, expireTime)
}

// GetAccountInitializer2 is a free data retrieval call binding the contract method 0x53ffc3b0.
//
// Solidity: function getAccountInitializer2(uint256 pubKeyX, uint256 pubKeyY, address validator, address eoaSigner, address target, bytes targetData, bytes targetData1) pure returns(bytes, bytes)
func (_Helper *HelperCaller) GetAccountInitializer2(opts *bind.CallOpts, pubKeyX *big.Int, pubKeyY *big.Int, validator common.Address, eoaSigner common.Address, target common.Address, targetData []byte, targetData1 []byte) ([]byte, []byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getAccountInitializer2", pubKeyX, pubKeyY, validator, eoaSigner, target, targetData, targetData1)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetAccountInitializer2 is a free data retrieval call binding the contract method 0x53ffc3b0.
//
// Solidity: function getAccountInitializer2(uint256 pubKeyX, uint256 pubKeyY, address validator, address eoaSigner, address target, bytes targetData, bytes targetData1) pure returns(bytes, bytes)
func (_Helper *HelperSession) GetAccountInitializer2(pubKeyX *big.Int, pubKeyY *big.Int, validator common.Address, eoaSigner common.Address, target common.Address, targetData []byte, targetData1 []byte) ([]byte, []byte, error) {
	return _Helper.Contract.GetAccountInitializer2(&_Helper.CallOpts, pubKeyX, pubKeyY, validator, eoaSigner, target, targetData, targetData1)
}

// GetAccountInitializer2 is a free data retrieval call binding the contract method 0x53ffc3b0.
//
// Solidity: function getAccountInitializer2(uint256 pubKeyX, uint256 pubKeyY, address validator, address eoaSigner, address target, bytes targetData, bytes targetData1) pure returns(bytes, bytes)
func (_Helper *HelperCallerSession) GetAccountInitializer2(pubKeyX *big.Int, pubKeyY *big.Int, validator common.Address, eoaSigner common.Address, target common.Address, targetData []byte, targetData1 []byte) ([]byte, []byte, error) {
	return _Helper.Contract.GetAccountInitializer2(&_Helper.CallOpts, pubKeyX, pubKeyY, validator, eoaSigner, target, targetData, targetData1)
}

// GetBlocktimeStamp is a free data retrieval call binding the contract method 0xba2f6e5e.
//
// Solidity: function getBlocktimeStamp() view returns(uint256)
func (_Helper *HelperCaller) GetBlocktimeStamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getBlocktimeStamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlocktimeStamp is a free data retrieval call binding the contract method 0xba2f6e5e.
//
// Solidity: function getBlocktimeStamp() view returns(uint256)
func (_Helper *HelperSession) GetBlocktimeStamp() (*big.Int, error) {
	return _Helper.Contract.GetBlocktimeStamp(&_Helper.CallOpts)
}

// GetBlocktimeStamp is a free data retrieval call binding the contract method 0xba2f6e5e.
//
// Solidity: function getBlocktimeStamp() view returns(uint256)
func (_Helper *HelperCallerSession) GetBlocktimeStamp() (*big.Int, error) {
	return _Helper.Contract.GetBlocktimeStamp(&_Helper.CallOpts)
}

// GetClientJson is a free data retrieval call binding the contract method 0xf0dc3df5.
//
// Solidity: function getClientJson(string clientDataJSONPre, string clientDataJSONPost, bytes32 userOpHash) pure returns(string clientDataJSON, bytes message, bytes32 messageHash)
func (_Helper *HelperCaller) GetClientJson(opts *bind.CallOpts, clientDataJSONPre string, clientDataJSONPost string, userOpHash [32]byte) (struct {
	ClientDataJSON string
	Message        []byte
	MessageHash    [32]byte
}, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getClientJson", clientDataJSONPre, clientDataJSONPost, userOpHash)

	outstruct := new(struct {
		ClientDataJSON string
		Message        []byte
		MessageHash    [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ClientDataJSON = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Message = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.MessageHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetClientJson is a free data retrieval call binding the contract method 0xf0dc3df5.
//
// Solidity: function getClientJson(string clientDataJSONPre, string clientDataJSONPost, bytes32 userOpHash) pure returns(string clientDataJSON, bytes message, bytes32 messageHash)
func (_Helper *HelperSession) GetClientJson(clientDataJSONPre string, clientDataJSONPost string, userOpHash [32]byte) (struct {
	ClientDataJSON string
	Message        []byte
	MessageHash    [32]byte
}, error) {
	return _Helper.Contract.GetClientJson(&_Helper.CallOpts, clientDataJSONPre, clientDataJSONPost, userOpHash)
}

// GetClientJson is a free data retrieval call binding the contract method 0xf0dc3df5.
//
// Solidity: function getClientJson(string clientDataJSONPre, string clientDataJSONPost, bytes32 userOpHash) pure returns(string clientDataJSON, bytes message, bytes32 messageHash)
func (_Helper *HelperCallerSession) GetClientJson(clientDataJSONPre string, clientDataJSONPost string, userOpHash [32]byte) (struct {
	ClientDataJSON string
	Message        []byte
	MessageHash    [32]byte
}, error) {
	return _Helper.Contract.GetClientJson(&_Helper.CallOpts, clientDataJSONPre, clientDataJSONPost, userOpHash)
}

// GetConfigSetSignerHash is a free data retrieval call binding the contract method 0xc3c51fc9.
//
// Solidity: function getConfigSetSignerHash(address config, address sender, address eoaSigner, uint256 expireTime) pure returns(bytes32)
func (_Helper *HelperCaller) GetConfigSetSignerHash(opts *bind.CallOpts, config common.Address, sender common.Address, eoaSigner common.Address, expireTime *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getConfigSetSignerHash", config, sender, eoaSigner, expireTime)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConfigSetSignerHash is a free data retrieval call binding the contract method 0xc3c51fc9.
//
// Solidity: function getConfigSetSignerHash(address config, address sender, address eoaSigner, uint256 expireTime) pure returns(bytes32)
func (_Helper *HelperSession) GetConfigSetSignerHash(config common.Address, sender common.Address, eoaSigner common.Address, expireTime *big.Int) ([32]byte, error) {
	return _Helper.Contract.GetConfigSetSignerHash(&_Helper.CallOpts, config, sender, eoaSigner, expireTime)
}

// GetConfigSetSignerHash is a free data retrieval call binding the contract method 0xc3c51fc9.
//
// Solidity: function getConfigSetSignerHash(address config, address sender, address eoaSigner, uint256 expireTime) pure returns(bytes32)
func (_Helper *HelperCallerSession) GetConfigSetSignerHash(config common.Address, sender common.Address, eoaSigner common.Address, expireTime *big.Int) ([32]byte, error) {
	return _Helper.Contract.GetConfigSetSignerHash(&_Helper.CallOpts, config, sender, eoaSigner, expireTime)
}

// GetFactoryCreateAccountHash is a free data retrieval call binding the contract method 0xcaeacff6.
//
// Solidity: function getFactoryCreateAccountHash(address factory, uint256 _salt, uint256 expireTime, bytes _initializer) pure returns(bytes32)
func (_Helper *HelperCaller) GetFactoryCreateAccountHash(opts *bind.CallOpts, factory common.Address, _salt *big.Int, expireTime *big.Int, _initializer []byte) ([32]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getFactoryCreateAccountHash", factory, _salt, expireTime, _initializer)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFactoryCreateAccountHash is a free data retrieval call binding the contract method 0xcaeacff6.
//
// Solidity: function getFactoryCreateAccountHash(address factory, uint256 _salt, uint256 expireTime, bytes _initializer) pure returns(bytes32)
func (_Helper *HelperSession) GetFactoryCreateAccountHash(factory common.Address, _salt *big.Int, expireTime *big.Int, _initializer []byte) ([32]byte, error) {
	return _Helper.Contract.GetFactoryCreateAccountHash(&_Helper.CallOpts, factory, _salt, expireTime, _initializer)
}

// GetFactoryCreateAccountHash is a free data retrieval call binding the contract method 0xcaeacff6.
//
// Solidity: function getFactoryCreateAccountHash(address factory, uint256 _salt, uint256 expireTime, bytes _initializer) pure returns(bytes32)
func (_Helper *HelperCallerSession) GetFactoryCreateAccountHash(factory common.Address, _salt *big.Int, expireTime *big.Int, _initializer []byte) ([32]byte, error) {
	return _Helper.Contract.GetFactoryCreateAccountHash(&_Helper.CallOpts, factory, _salt, expireTime, _initializer)
}

// GetFactorySig is a free data retrieval call binding the contract method 0xd7d0679f.
//
// Solidity: function getFactorySig(uint256 expireTime, bytes sig) pure returns(bytes)
func (_Helper *HelperCaller) GetFactorySig(opts *bind.CallOpts, expireTime *big.Int, sig []byte) ([]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getFactorySig", expireTime, sig)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetFactorySig is a free data retrieval call binding the contract method 0xd7d0679f.
//
// Solidity: function getFactorySig(uint256 expireTime, bytes sig) pure returns(bytes)
func (_Helper *HelperSession) GetFactorySig(expireTime *big.Int, sig []byte) ([]byte, error) {
	return _Helper.Contract.GetFactorySig(&_Helper.CallOpts, expireTime, sig)
}

// GetFactorySig is a free data retrieval call binding the contract method 0xd7d0679f.
//
// Solidity: function getFactorySig(uint256 expireTime, bytes sig) pure returns(bytes)
func (_Helper *HelperCallerSession) GetFactorySig(expireTime *big.Int, sig []byte) ([]byte, error) {
	return _Helper.Contract.GetFactorySig(&_Helper.CallOpts, expireTime, sig)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string data) pure returns(bytes32)
func (_Helper *HelperCaller) GetHash(opts *bind.CallOpts, data string) ([32]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getHash", data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string data) pure returns(bytes32)
func (_Helper *HelperSession) GetHash(data string) ([32]byte, error) {
	return _Helper.Contract.GetHash(&_Helper.CallOpts, data)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string data) pure returns(bytes32)
func (_Helper *HelperCallerSession) GetHash(data string) ([32]byte, error) {
	return _Helper.Contract.GetHash(&_Helper.CallOpts, data)
}

// GetHash1 is a free data retrieval call binding the contract method 0x4f2a4c7b.
//
// Solidity: function getHash1(bytes32 hash) pure returns(bytes32)
func (_Helper *HelperCaller) GetHash1(opts *bind.CallOpts, hash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getHash1", hash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash1 is a free data retrieval call binding the contract method 0x4f2a4c7b.
//
// Solidity: function getHash1(bytes32 hash) pure returns(bytes32)
func (_Helper *HelperSession) GetHash1(hash [32]byte) ([32]byte, error) {
	return _Helper.Contract.GetHash1(&_Helper.CallOpts, hash)
}

// GetHash1 is a free data retrieval call binding the contract method 0x4f2a4c7b.
//
// Solidity: function getHash1(bytes32 hash) pure returns(bytes32)
func (_Helper *HelperCallerSession) GetHash1(hash [32]byte) ([32]byte, error) {
	return _Helper.Contract.GetHash1(&_Helper.CallOpts, hash)
}

// GetPubkeyHash is a free data retrieval call binding the contract method 0x019b696e.
//
// Solidity: function getPubkeyHash(uint256 pubKeyX, uint256 pubKeyY) pure returns(bytes32, bytes)
func (_Helper *HelperCaller) GetPubkeyHash(opts *bind.CallOpts, pubKeyX *big.Int, pubKeyY *big.Int) ([32]byte, []byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getPubkeyHash", pubKeyX, pubKeyY)

	if err != nil {
		return *new([32]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetPubkeyHash is a free data retrieval call binding the contract method 0x019b696e.
//
// Solidity: function getPubkeyHash(uint256 pubKeyX, uint256 pubKeyY) pure returns(bytes32, bytes)
func (_Helper *HelperSession) GetPubkeyHash(pubKeyX *big.Int, pubKeyY *big.Int) ([32]byte, []byte, error) {
	return _Helper.Contract.GetPubkeyHash(&_Helper.CallOpts, pubKeyX, pubKeyY)
}

// GetPubkeyHash is a free data retrieval call binding the contract method 0x019b696e.
//
// Solidity: function getPubkeyHash(uint256 pubKeyX, uint256 pubKeyY) pure returns(bytes32, bytes)
func (_Helper *HelperCallerSession) GetPubkeyHash(pubKeyX *big.Int, pubKeyY *big.Int) ([32]byte, []byte, error) {
	return _Helper.Contract.GetPubkeyHash(&_Helper.CallOpts, pubKeyX, pubKeyY)
}

// GetSignature2 is a free data retrieval call binding the contract method 0x98a1941f.
//
// Solidity: function getSignature2(uint256 pubKeyX, uint256 pubKeyY, bytes passkeySig, bytes eoaSignature, uint256 validationData) pure returns(bytes)
func (_Helper *HelperCaller) GetSignature2(opts *bind.CallOpts, pubKeyX *big.Int, pubKeyY *big.Int, passkeySig []byte, eoaSignature []byte, validationData *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getSignature2", pubKeyX, pubKeyY, passkeySig, eoaSignature, validationData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSignature2 is a free data retrieval call binding the contract method 0x98a1941f.
//
// Solidity: function getSignature2(uint256 pubKeyX, uint256 pubKeyY, bytes passkeySig, bytes eoaSignature, uint256 validationData) pure returns(bytes)
func (_Helper *HelperSession) GetSignature2(pubKeyX *big.Int, pubKeyY *big.Int, passkeySig []byte, eoaSignature []byte, validationData *big.Int) ([]byte, error) {
	return _Helper.Contract.GetSignature2(&_Helper.CallOpts, pubKeyX, pubKeyY, passkeySig, eoaSignature, validationData)
}

// GetSignature2 is a free data retrieval call binding the contract method 0x98a1941f.
//
// Solidity: function getSignature2(uint256 pubKeyX, uint256 pubKeyY, bytes passkeySig, bytes eoaSignature, uint256 validationData) pure returns(bytes)
func (_Helper *HelperCallerSession) GetSignature2(pubKeyX *big.Int, pubKeyY *big.Int, passkeySig []byte, eoaSignature []byte, validationData *big.Int) ([]byte, error) {
	return _Helper.Contract.GetSignature2(&_Helper.CallOpts, pubKeyX, pubKeyY, passkeySig, eoaSignature, validationData)
}

// GetSignature2Decode is a free data retrieval call binding the contract method 0xf29ad216.
//
// Solidity: function getSignature2Decode(bytes sig) pure returns(bytes, bytes, uint256)
func (_Helper *HelperCaller) GetSignature2Decode(opts *bind.CallOpts, sig []byte) ([]byte, []byte, *big.Int, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getSignature2Decode", sig)

	if err != nil {
		return *new([]byte), *new([]byte), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetSignature2Decode is a free data retrieval call binding the contract method 0xf29ad216.
//
// Solidity: function getSignature2Decode(bytes sig) pure returns(bytes, bytes, uint256)
func (_Helper *HelperSession) GetSignature2Decode(sig []byte) ([]byte, []byte, *big.Int, error) {
	return _Helper.Contract.GetSignature2Decode(&_Helper.CallOpts, sig)
}

// GetSignature2Decode is a free data retrieval call binding the contract method 0xf29ad216.
//
// Solidity: function getSignature2Decode(bytes sig) pure returns(bytes, bytes, uint256)
func (_Helper *HelperCallerSession) GetSignature2Decode(sig []byte) ([]byte, []byte, *big.Int, error) {
	return _Helper.Contract.GetSignature2Decode(&_Helper.CallOpts, sig)
}

// GetUserOpHashWithEntryPoint is a free data retrieval call binding the contract method 0x53530b18.
//
// Solidity: function getUserOpHashWithEntryPoint(address entryPoint, uint256 chainid, (address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) pure returns(bytes32)
func (_Helper *HelperCaller) GetUserOpHashWithEntryPoint(opts *bind.CallOpts, entryPoint common.Address, chainid *big.Int, userOp PackedUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getUserOpHashWithEntryPoint", entryPoint, chainid, userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHashWithEntryPoint is a free data retrieval call binding the contract method 0x53530b18.
//
// Solidity: function getUserOpHashWithEntryPoint(address entryPoint, uint256 chainid, (address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) pure returns(bytes32)
func (_Helper *HelperSession) GetUserOpHashWithEntryPoint(entryPoint common.Address, chainid *big.Int, userOp PackedUserOperation) ([32]byte, error) {
	return _Helper.Contract.GetUserOpHashWithEntryPoint(&_Helper.CallOpts, entryPoint, chainid, userOp)
}

// GetUserOpHashWithEntryPoint is a free data retrieval call binding the contract method 0x53530b18.
//
// Solidity: function getUserOpHashWithEntryPoint(address entryPoint, uint256 chainid, (address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) pure returns(bytes32)
func (_Helper *HelperCallerSession) GetUserOpHashWithEntryPoint(entryPoint common.Address, chainid *big.Int, userOp PackedUserOperation) ([32]byte, error) {
	return _Helper.Contract.GetUserOpHashWithEntryPoint(&_Helper.CallOpts, entryPoint, chainid, userOp)
}

// GetValidationData is a free data retrieval call binding the contract method 0x0f4d65e1.
//
// Solidity: function getValidationData(uint256 validationData) pure returns(uint256)
func (_Helper *HelperCaller) GetValidationData(opts *bind.CallOpts, validationData *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getValidationData", validationData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidationData is a free data retrieval call binding the contract method 0x0f4d65e1.
//
// Solidity: function getValidationData(uint256 validationData) pure returns(uint256)
func (_Helper *HelperSession) GetValidationData(validationData *big.Int) (*big.Int, error) {
	return _Helper.Contract.GetValidationData(&_Helper.CallOpts, validationData)
}

// GetValidationData is a free data retrieval call binding the contract method 0x0f4d65e1.
//
// Solidity: function getValidationData(uint256 validationData) pure returns(uint256)
func (_Helper *HelperCallerSession) GetValidationData(validationData *big.Int) (*big.Int, error) {
	return _Helper.Contract.GetValidationData(&_Helper.CallOpts, validationData)
}

// MockRecover is a free data retrieval call binding the contract method 0x780dfbf6.
//
// Solidity: function mockRecover(bytes okxSignatureData) pure returns(address)
func (_Helper *HelperCaller) MockRecover(opts *bind.CallOpts, okxSignatureData []byte) (common.Address, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "mockRecover", okxSignatureData)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MockRecover is a free data retrieval call binding the contract method 0x780dfbf6.
//
// Solidity: function mockRecover(bytes okxSignatureData) pure returns(address)
func (_Helper *HelperSession) MockRecover(okxSignatureData []byte) (common.Address, error) {
	return _Helper.Contract.MockRecover(&_Helper.CallOpts, okxSignatureData)
}

// MockRecover is a free data retrieval call binding the contract method 0x780dfbf6.
//
// Solidity: function mockRecover(bytes okxSignatureData) pure returns(address)
func (_Helper *HelperCallerSession) MockRecover(okxSignatureData []byte) (common.Address, error) {
	return _Helper.Contract.MockRecover(&_Helper.CallOpts, okxSignatureData)
}

// PackValidationData is a free data retrieval call binding the contract method 0x0f559e06.
//
// Solidity: function packValidationData(address aggregator, uint256 validUntil, uint256 validAfter) pure returns(uint256)
func (_Helper *HelperCaller) PackValidationData(opts *bind.CallOpts, aggregator common.Address, validUntil *big.Int, validAfter *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "packValidationData", aggregator, validUntil, validAfter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PackValidationData is a free data retrieval call binding the contract method 0x0f559e06.
//
// Solidity: function packValidationData(address aggregator, uint256 validUntil, uint256 validAfter) pure returns(uint256)
func (_Helper *HelperSession) PackValidationData(aggregator common.Address, validUntil *big.Int, validAfter *big.Int) (*big.Int, error) {
	return _Helper.Contract.PackValidationData(&_Helper.CallOpts, aggregator, validUntil, validAfter)
}

// PackValidationData is a free data retrieval call binding the contract method 0x0f559e06.
//
// Solidity: function packValidationData(address aggregator, uint256 validUntil, uint256 validAfter) pure returns(uint256)
func (_Helper *HelperCallerSession) PackValidationData(aggregator common.Address, validUntil *big.Int, validAfter *big.Int) (*big.Int, error) {
	return _Helper.Contract.PackValidationData(&_Helper.CallOpts, aggregator, validUntil, validAfter)
}

// ParseValidationData is a free data retrieval call binding the contract method 0xa4b2282e.
//
// Solidity: function parseValidationData(uint256 validationData) pure returns((address,uint48,uint48) data)
func (_Helper *HelperCaller) ParseValidationData(opts *bind.CallOpts, validationData *big.Int) (ValidationData, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "parseValidationData", validationData)

	if err != nil {
		return *new(ValidationData), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidationData)).(*ValidationData)

	return out0, err

}

// ParseValidationData is a free data retrieval call binding the contract method 0xa4b2282e.
//
// Solidity: function parseValidationData(uint256 validationData) pure returns((address,uint48,uint48) data)
func (_Helper *HelperSession) ParseValidationData(validationData *big.Int) (ValidationData, error) {
	return _Helper.Contract.ParseValidationData(&_Helper.CallOpts, validationData)
}

// ParseValidationData is a free data retrieval call binding the contract method 0xa4b2282e.
//
// Solidity: function parseValidationData(uint256 validationData) pure returns((address,uint48,uint48) data)
func (_Helper *HelperCallerSession) ParseValidationData(validationData *big.Int) (ValidationData, error) {
	return _Helper.Contract.ParseValidationData(&_Helper.CallOpts, validationData)
}

// PasskeyVerify is a free data retrieval call binding the contract method 0xc6e8a99e.
//
// Solidity: function passkeyVerify(bytes32 okxHash, uint256 r, uint256 s, uint256 x, uint256 y, uint8 verifyType, string clientDataJSON) view returns(bool, uint256)
func (_Helper *HelperCaller) PasskeyVerify(opts *bind.CallOpts, okxHash [32]byte, r *big.Int, s *big.Int, x *big.Int, y *big.Int, verifyType uint8, clientDataJSON string) (bool, *big.Int, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "passkeyVerify", okxHash, r, s, x, y, verifyType, clientDataJSON)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// PasskeyVerify is a free data retrieval call binding the contract method 0xc6e8a99e.
//
// Solidity: function passkeyVerify(bytes32 okxHash, uint256 r, uint256 s, uint256 x, uint256 y, uint8 verifyType, string clientDataJSON) view returns(bool, uint256)
func (_Helper *HelperSession) PasskeyVerify(okxHash [32]byte, r *big.Int, s *big.Int, x *big.Int, y *big.Int, verifyType uint8, clientDataJSON string) (bool, *big.Int, error) {
	return _Helper.Contract.PasskeyVerify(&_Helper.CallOpts, okxHash, r, s, x, y, verifyType, clientDataJSON)
}

// PasskeyVerify is a free data retrieval call binding the contract method 0xc6e8a99e.
//
// Solidity: function passkeyVerify(bytes32 okxHash, uint256 r, uint256 s, uint256 x, uint256 y, uint8 verifyType, string clientDataJSON) view returns(bool, uint256)
func (_Helper *HelperCallerSession) PasskeyVerify(okxHash [32]byte, r *big.Int, s *big.Int, x *big.Int, y *big.Int, verifyType uint8, clientDataJSON string) (bool, *big.Int, error) {
	return _Helper.Contract.PasskeyVerify(&_Helper.CallOpts, okxHash, r, s, x, y, verifyType, clientDataJSON)
}

// PasskeyVerify1 is a free data retrieval call binding the contract method 0x9bffea24.
//
// Solidity: function passkeyVerify1(bytes32 messageHash, uint256 r, uint256 s, uint256 x, uint256 y, uint8 verifier) view returns(bool)
func (_Helper *HelperCaller) PasskeyVerify1(opts *bind.CallOpts, messageHash [32]byte, r *big.Int, s *big.Int, x *big.Int, y *big.Int, verifier uint8) (bool, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "passkeyVerify1", messageHash, r, s, x, y, verifier)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PasskeyVerify1 is a free data retrieval call binding the contract method 0x9bffea24.
//
// Solidity: function passkeyVerify1(bytes32 messageHash, uint256 r, uint256 s, uint256 x, uint256 y, uint8 verifier) view returns(bool)
func (_Helper *HelperSession) PasskeyVerify1(messageHash [32]byte, r *big.Int, s *big.Int, x *big.Int, y *big.Int, verifier uint8) (bool, error) {
	return _Helper.Contract.PasskeyVerify1(&_Helper.CallOpts, messageHash, r, s, x, y, verifier)
}

// PasskeyVerify1 is a free data retrieval call binding the contract method 0x9bffea24.
//
// Solidity: function passkeyVerify1(bytes32 messageHash, uint256 r, uint256 s, uint256 x, uint256 y, uint8 verifier) view returns(bool)
func (_Helper *HelperCallerSession) PasskeyVerify1(messageHash [32]byte, r *big.Int, s *big.Int, x *big.Int, y *big.Int, verifier uint8) (bool, error) {
	return _Helper.Contract.PasskeyVerify1(&_Helper.CallOpts, messageHash, r, s, x, y, verifier)
}

// RecoverAddress is a free data retrieval call binding the contract method 0xecaffa77.
//
// Solidity: function recoverAddress(bytes signature, bytes32 hash) pure returns(address)
func (_Helper *HelperCaller) RecoverAddress(opts *bind.CallOpts, signature []byte, hash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "recoverAddress", signature, hash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverAddress is a free data retrieval call binding the contract method 0xecaffa77.
//
// Solidity: function recoverAddress(bytes signature, bytes32 hash) pure returns(address)
func (_Helper *HelperSession) RecoverAddress(signature []byte, hash [32]byte) (common.Address, error) {
	return _Helper.Contract.RecoverAddress(&_Helper.CallOpts, signature, hash)
}

// RecoverAddress is a free data retrieval call binding the contract method 0xecaffa77.
//
// Solidity: function recoverAddress(bytes signature, bytes32 hash) pure returns(address)
func (_Helper *HelperCallerSession) RecoverAddress(signature []byte, hash [32]byte) (common.Address, error) {
	return _Helper.Contract.RecoverAddress(&_Helper.CallOpts, signature, hash)
}

// RecoverAddress1 is a free data retrieval call binding the contract method 0x22de96b1.
//
// Solidity: function recoverAddress1(uint8 v, bytes32 r, bytes32 s, bytes32 hash) pure returns(address)
func (_Helper *HelperCaller) RecoverAddress1(opts *bind.CallOpts, v uint8, r [32]byte, s [32]byte, hash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "recoverAddress1", v, r, s, hash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverAddress1 is a free data retrieval call binding the contract method 0x22de96b1.
//
// Solidity: function recoverAddress1(uint8 v, bytes32 r, bytes32 s, bytes32 hash) pure returns(address)
func (_Helper *HelperSession) RecoverAddress1(v uint8, r [32]byte, s [32]byte, hash [32]byte) (common.Address, error) {
	return _Helper.Contract.RecoverAddress1(&_Helper.CallOpts, v, r, s, hash)
}

// RecoverAddress1 is a free data retrieval call binding the contract method 0x22de96b1.
//
// Solidity: function recoverAddress1(uint8 v, bytes32 r, bytes32 s, bytes32 hash) pure returns(address)
func (_Helper *HelperCallerSession) RecoverAddress1(v uint8, r [32]byte, s [32]byte, hash [32]byte) (common.Address, error) {
	return _Helper.Contract.RecoverAddress1(&_Helper.CallOpts, v, r, s, hash)
}

// SplitSig is a free data retrieval call binding the contract method 0x5db1bc88.
//
// Solidity: function splitSig(bytes signature) pure returns(bytes32, bytes32, uint8)
func (_Helper *HelperCaller) SplitSig(opts *bind.CallOpts, signature []byte) ([32]byte, [32]byte, uint8, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "splitSig", signature)

	if err != nil {
		return *new([32]byte), *new([32]byte), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	out2 := *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return out0, out1, out2, err

}

// SplitSig is a free data retrieval call binding the contract method 0x5db1bc88.
//
// Solidity: function splitSig(bytes signature) pure returns(bytes32, bytes32, uint8)
func (_Helper *HelperSession) SplitSig(signature []byte) ([32]byte, [32]byte, uint8, error) {
	return _Helper.Contract.SplitSig(&_Helper.CallOpts, signature)
}

// SplitSig is a free data retrieval call binding the contract method 0x5db1bc88.
//
// Solidity: function splitSig(bytes signature) pure returns(bytes32, bytes32, uint8)
func (_Helper *HelperCallerSession) SplitSig(signature []byte) ([32]byte, [32]byte, uint8, error) {
	return _Helper.Contract.SplitSig(&_Helper.CallOpts, signature)
}

// VerifyPasskeySignature is a free data retrieval call binding the contract method 0x179e90fd.
//
// Solidity: function verifyPasskeySignature(bytes challenge, bytes authenticatorData, string clientDataJSON, uint256 r, uint256 s, uint256 x, uint256 y, uint8 verifier) view returns(bool, uint256)
func (_Helper *HelperCaller) VerifyPasskeySignature(opts *bind.CallOpts, challenge []byte, authenticatorData []byte, clientDataJSON string, r *big.Int, s *big.Int, x *big.Int, y *big.Int, verifier uint8) (bool, *big.Int, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "verifyPasskeySignature", challenge, authenticatorData, clientDataJSON, r, s, x, y, verifier)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// VerifyPasskeySignature is a free data retrieval call binding the contract method 0x179e90fd.
//
// Solidity: function verifyPasskeySignature(bytes challenge, bytes authenticatorData, string clientDataJSON, uint256 r, uint256 s, uint256 x, uint256 y, uint8 verifier) view returns(bool, uint256)
func (_Helper *HelperSession) VerifyPasskeySignature(challenge []byte, authenticatorData []byte, clientDataJSON string, r *big.Int, s *big.Int, x *big.Int, y *big.Int, verifier uint8) (bool, *big.Int, error) {
	return _Helper.Contract.VerifyPasskeySignature(&_Helper.CallOpts, challenge, authenticatorData, clientDataJSON, r, s, x, y, verifier)
}

// VerifyPasskeySignature is a free data retrieval call binding the contract method 0x179e90fd.
//
// Solidity: function verifyPasskeySignature(bytes challenge, bytes authenticatorData, string clientDataJSON, uint256 r, uint256 s, uint256 x, uint256 y, uint8 verifier) view returns(bool, uint256)
func (_Helper *HelperCallerSession) VerifyPasskeySignature(challenge []byte, authenticatorData []byte, clientDataJSON string, r *big.Int, s *big.Int, x *big.Int, y *big.Int, verifier uint8) (bool, *big.Int, error) {
	return _Helper.Contract.VerifyPasskeySignature(&_Helper.CallOpts, challenge, authenticatorData, clientDataJSON, r, s, x, y, verifier)
}
