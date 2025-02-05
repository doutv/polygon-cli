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
	ABI: "[{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"okxSignatureData\",\"type\":\"bytes\"}],\"name\":\"PasskeyFormatDemo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"name\":\"checkValidationDate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"computePredictionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"dealData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"okxSignatureData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"}],\"name\":\"encodePacked\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"enumVerifierType\",\"name\":\"verifyType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"clientDataJSON\",\"type\":\"string\"}],\"name\":\"encodePasskeySig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uopHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"}],\"name\":\"encodeUopHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pubKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyY\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eoaSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"targetData1\",\"type\":\"bytes\"}],\"name\":\"getAccountInitializer2\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlocktimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientDataJSONPre\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"clientDataJSONPost\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"name\":\"getClientJson\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"clientDataJSON\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"config\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eoaSigner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"}],\"name\":\"getConfigSetSignerHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_initializer\",\"type\":\"bytes\"}],\"name\":\"getFactoryCreateAccountHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getHash1\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"getPackedSig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pubKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyY\",\"type\":\"uint256\"}],\"name\":\"getPubkeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientAddress\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"}],\"name\":\"getSharelinkHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pubKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyY\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"passkeySig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"eoaSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"name\":\"getSignature2\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"getSignature2Decode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHashWithEntryPoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"name\":\"getValidationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"okxSignatureData\",\"type\":\"bytes\"}],\"name\":\"mockRecover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validAfter\",\"type\":\"uint256\"}],\"name\":\"packValidationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"name\":\"parseValidationData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"}],\"internalType\":\"structValidationData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"okxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"enumVerifierType\",\"name\":\"verifyType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"clientDataJSON\",\"type\":\"string\"}],\"name\":\"passkeyVerify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"enumVerifierType\",\"name\":\"verifier\",\"type\":\"uint8\"}],\"name\":\"passkeyVerify1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"recoverAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"recoverAddress1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"splitSig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"challenge\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"authenticatorData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"clientDataJSON\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"enumVerifierType\",\"name\":\"verifier\",\"type\":\"uint8\"}],\"name\":\"verifyPasskeySignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557612510908161001b8239f35b600080fdfe608060408181526004918236101561001657600080fd5b60009260e08435811c928363019b696e14611693575082630aeaa733146115e05782630f4d65e1146115be5782630f559e0614611580578263179e90fd146114ec57826318c82e08146114b557826322de96b1146114945782634f2a4c7b1461144b57826353530b181461132157826353ffc3b0146110015782635b6beeb914610e265782635db1bc8814610dd2578263627041f514610d2c57826372e8eeea14610c95578263780dfbf614610bbc578263899413f614610b895782638bb9241114610b2f57826390ea056014610ab257826398a1941f146109d85782639bffea2414610990578263a4b2282e14610940578263ba2f6e5e14610924578263c3c51fc91461088d578263c6e8a99e146107ec578263caeacff61461075d578263e585b437146104be57508163ecaffa771461043a578163f0dc3df514610266578163f29ad216146101d5575063f5f3b08e1461017157600080fd5b346101d15760803660031901126101d15760209061018d6117a5565b81518381019146835260018060a01b0316838201526024356060820152604435608082015260643560a082015260a081526101c781611838565b5190209051908152f35b5080fd5b9050823461025b57602036600319011261025b576001600160401b0391803583811161026257610207913691016118e1565b819391019060608483031261026257833581811161025e578261022b9186016118c6565b92602085013591821161025b575061025792916102499185016118c6565b84519485940135918461190e565b0390f35b80fd5b8380fd5b8280fd5b91905034610262576060366003190112610262576001600160401b0382358181116104365761029890369085016118c6565b92602435918211610436576102af913691016118c6565b9161032e61037a83516020928391604435838201528281526102d0816117d1565b8651906102dc82611802565b8782527f4142434445464748494a4b4c4d4e4f505152535455565758595a616263646566848301527f6768696a6b6c6d6e6f707172737475767778797a303132333435363738392d5f8883015261218e565b8551968791836103478185019889815193849201611708565b830161035b82518093878085019101611708565b0161036e82518093868085019101611708565b01038087520185611853565b80856103986103876119ba565b948651809281928a51928391611708565b8101039060025afa1561042c5780856103ed8151946103dd8780518098876103c98184019687815193849201611708565b820190888201520386810189520187611853565b8651809281928851928391611708565b8101039060025afa1561042c5761042490610417955192845196879660608852606088019061172b565b918683039087015261172b565b918301520390f35b82513d86823e3d90fd5b8480fd5b90503461026257816003193601126102625780356001600160401b03811161025e57602093603c6104756104ae95946104a5943691016118c6565b917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152602435601c5220611fa4565b90939193611ebb565b516001600160a01b039091168152f35b909291503461025e578160031936011261025e576024356001600160401b038111610436576104f090369085016118c6565b9282519060209485808401528483016008905260609081840167686f6f6b4461746160c01b90528184526105238461181d565b61052b6119ba565b86519561053787611838565b608687528887017f7b2274797065223a22776562617574686e2e676574222c226368616c6c656e6790528787017f65223a226354534253556f675f374d503155455177744b70794f474c6965534b90528387017f307043644f6b2d47396a514f6b5f45222c226f726967696e223a22687474703a9052608087017f2f2f6c6f63616c686f73743a38303030222c2263726f73734f726967696e223a905260a087016566616c73657d60d01b90528751998a9261010090818c86015261012085016106009161172b565b98601f19998a868203018c8701526106179161172b565b6001858801527f7fde2a6a50c8b06dfa8869b0c217665828fd6c9bbba6e1e16034bf9b9695899960808601527f21ede5fef42f27196a2f42167c8c4428a3f592b587d780c69eeac62ad5ddc44960a08601527f13e5e3c1d310f8515c2e28f499755c5f9433cb70ba1d7ecf4f276ff4b63a635260c08601527f19ae7d5e4f7341511fef9fc70c4e447b499505e1612dc2e792b272f9d83c6e4f938501939093528301520385810189526106ca9089611853565b85519788916106db9189840161177d565b0384810188526106eb9088611853565b84519135868301528582526106ff826117d1565b84519687938288860152608085016107169161172b565b8585820301878601526107289161172b565b9084848303019084015261073b9161172b565b03908101845261074b9084611853565b5191818392835282016102579161172b565b5082843461025b57608036600319011261025b576107796117a5565b90606435906001600160401b03821161025b57506101c76107a26020956107ce933691016118c6565b8451468782019081526001600160a01b039095168187015260a0606082015292839160c083019061172b565b602435608083015260443560a083015203601f198101835282611853565b84919250346101d15736600319011261025b5760a43590600382101561025b5760c435906001600160401b03821161025b575061086c6102579261083661087493369087016118c6565b945a956108416119ba565b88519235602084015260208352610857836117d1565b60843592606435926044359260243592611a6d565b925a90611964565b9251911515825260208201929092529081906040820190565b508383346101d15760803660031901126101d1576108a96117a5565b906024356001600160a01b038181169291839003610436576108c96117bb565b8251938260208601961686528385015216606083015260643560808301526080825260a08201938285106001600160401b038611176109115750602094508390525190208152f35b634e487b7160e01b815260418652602490fd5b505050346101d157816003193601126101d15760209051428152f35b5082843461025b57602036600319011261025b5750610961606092356119fd565b9080519160018060a01b0381511683528160208201519165ffffffffffff809316602086015201511690820152f35b509050346102625760c03660031901126102625760a43592600384101561025b57506020926109cf916084359060643590604435906024359035611ccb565b90519015158152f35b508383346101d15760a03660031901126101d1576001600160401b0360443581811161025e57610a0b90369086016118c6565b9260643591821161025b575090610aa3610a2b61025793369087016118c6565b825195610a37876117d1565b35865260209460243586880152610a538451978789019061193c565b838752610a5f87611802565b610a6f845192839288840161177d565b0390610a83601f1992838101835282611853565b610a9784519788926084359189850161190e565b03908101865285611853565b5192828493845283019061172b565b8382863461025b578260031936011261025b57602435906001600160401b03821161025b5750610b1b92610aa3610aef61025793369086016118c6565b938251602095869283830190358152838352610b0a836117d1565b855198899351809286860190611708565b820161036e82518093868085019101611708565b505050346101d157816003193601126101d15760209061027b8151610b5684830182611853565b8181528381019161226083396101c78484518093610b7c83830196879251928391611708565b8101038084520182611853565b5090503461026257816003193601126102625760209250815183810191358252602435838201528281526101c781611802565b84828534610262576020918260031936011261025e5780356001600160401b0391828211610c9157610bf0913691016118c6565b80518101918382868501940312610c9157848201519184810151918211610c8d570182603f82011215610c91578481015190610c2b82611874565b93610c3886519586611853565b828552858383010111610c8d57958392610c61603c938899886104ae9a6104a599019101611708565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c5220611fa4565b8680fd5b8580fd5b508383346101d157806003193601126101d157610cb06117a5565b91602435906001600160401b03821161025b57506034610cd961025795610d1a933691016118c6565b9383519485916bffffffffffffffffffffffff199060601b166020830152610d0a8151809260208686019101611708565b8101036014810185520183611853565b5191829160208352602083019061172b565b838591346102625760803660031901126102625760443592600384101561025b57606435906001600160401b03821161025b575093610d1a91610d7561025796369084016118c6565b91610d7e6119ba565b95610dab610d99875198899560c0602088015286019061172b565b601f199586868303018987015261172b565b916001606085015235608084015260243560a084015260c083015203908101845283611853565b508383346101d15760203660031901126101d15782356001600160401b03811161026257606093610e05913691016118c6565b906020820151928482840151930151901a9181519384526020840152820152f35b5082843461025b57602092836003193601126101d15780356001600160401b03811161026257610e5990369083016118c6565b908151809184907a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000080841015610ff4575b50876d04ee2d6d415b85acef810000000080861015610fe5575b5050662386f26fc1000080851015610fd6575b506305f5e10080851015610fc7575b5061271080851015610fbb575b50506064831015610fad575b600a80931015610fa5575b9060219160019281610efa6001869401611f72565b9750870101905b610f6f575b505050506101c7603a84518093878201957f19457468657265756d205369676e6564204d6573736167653a0a0000000000008752610f4c815180928b8787019101611708565b8201610f60825180938b8785019101611708565b0103601a810184520182611853565b600019019083906f181899199a1a9b1b9c1cb0b131b232b360811b8282061a835304918215610fa057919082610f01565b610f06565b600101610ee5565b916064600291049201610eda565b90930492018780610ece565b60089192940493019088610ec1565b60109192940493019088610eb2565b90919294049301908789610e9f565b8304935086915088610e85565b9091503461025e573660031901126102625761101b6117bb565b916064359360018060a01b038086168096036101d15760843595818716809703610262576001600160401b039660a4358881116104365761105f90369088016118e1565b93909160c435998a11610c915761107b87939a36908a016118e1565b93908151996110898b6117d1565b358a526020996024358b82015282516110a58c8201809361193c565b8381526110b181611802565b519020948a83519d8e97828901528a858901526110e76060986fffffffffffffffffffffffffffffffff8a82015289815261181d565b84519e8f906110f5826117d1565b600182528c5b8a858210611305575050906111199161111382611987565b52611987565b508451996111268b611802565b60028b528b805b8a8882106112d457505090611159929161118b97519461114c86611802565b898652850152369161188f565b8b82015261116689611987565b5261117088611987565b5089519361117d85611802565b8452888b850152369161188f565b87820152611198856119aa565b526111a2846119aa565b5085519760a089016080898b01528a5180915260c08a01908960c08260051b8d01019c019188905b8282106112a75750505050601f1995868a8c0301888b01528551808c528b8a80808301928460051b0101980192905b82821061126157505050501690870152608086015284900380820185526102579291611257916112299087611853565b82519463439fab9160e01b8187015260248601528461124b604482018861172b565b03908101855284611853565b519283928361177d565b909192978b808f8d8a858f8f8d859185936112999860019c03018c52519081511684528781015188850152015193820152019061172b565b9a01920192019092916111f9565b9091929c8b808f8f936112c69160019560bf199083030187525161172b565b9f01920192019092916111ca565b909294939596978051926112e784611802565b8084528584015282015282828d01015201918c918b8d96959461112d565b919493959697509182849101015201918e8d928d9695946110fb565b9150923461025b57600319916060368401126101d15761133f6117a5565b92604435906001600160401b039384831161025b5782880191610120809185360301126101d1576001600160a01b039361137c6044820185611f40565b90818b51918237209560c46113946064840187611f40565b90818d5191823720928b6113ab60e4830189611f40565b80925191823720938c51998960208c0199351689528d6024840135908c015260608b015260808a0152608481013560a08a015260a481013560c08a015201359087015261010090818701528552840194848610908611176114385750602096508386528251902093610140830194855216610160820152610180602435910152606081526101c78161181d565b634e487b7160e01b815260418852602490fd5b509050346102625760203660031901126102625782603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020955235601c52209051908152f35b83853461025b5760206104ae836104a56114ad36611750565b929092611e2b565b8382863461025b57602036600319011261025b57506114d49035611c7c565b82516001600160a01b03909216825215156020820152f35b8382863461025b5761010036600319011261025b576001600160401b039180358381116102625761152090369083016118c6565b9060243584811161025e5761153890369083016118c6565b9360443590811161025e5761154f913691016118c6565b60e43592600384101561025b5750610257926108749261086c925a9660c4359260a435926084359260643592611a6d565b505050346101d15760603660031901126101d15760209061159f6117a5565b90516001600160a01b0390911660243560a01b1760443560d01b178152f35b84828534610262576020366003190112610262576020925051903560a01b8152f35b83853461025b57506112576102576115f736611750565b9095929491610a9784519661166b8861162f838c876020850191604193918352602083015260ff60f81b9060f81b1660408201520190565b0394611643601f19968781018c528b611853565b61165f88519a8b9260208401528980840152606083019061172b565b038581018a5289611853565b85519889936020850191604193918352602083015260ff60f81b9060f81b1660408201520190565b925050833461025b578360031936011261025b57506116b1826117d1565b358152602435602082015281516116cc60208201809361193c565b8281526116d881611802565b5190206102578251826020820152602081526116f3816117d1565b8351938493845280602085015283019061172b565b60005b83811061171b5750506000910152565b818101518382015260200161170b565b9060209161174481518092818552858086019101611708565b601f01601f1916010190565b60809060031901126117785760043560ff811681036117785790602435906044359060643590565b600080fd5b90916117946117a29360408452604084019061172b565b91602081840391015261172b565b90565b600435906001600160a01b038216820361177857565b604435906001600160a01b038216820361177857565b604081019081106001600160401b038211176117ec57604052565b634e487b7160e01b600052604160045260246000fd5b606081019081106001600160401b038211176117ec57604052565b608081019081106001600160401b038211176117ec57604052565b60c081019081106001600160401b038211176117ec57604052565b90601f801991011681019081106001600160401b038211176117ec57604052565b6001600160401b0381116117ec57601f01601f191660200190565b92919261189b82611874565b916118a96040519384611853565b829481845281830111611778578281602093846000960137010152565b9080601f83011215611778578160206117a29335910161188f565b9181601f84011215611778578235916001600160401b038311611778576020838186019501011161177857565b9392916119379061192960409360608852606088019061172b565b90868203602088015261172b565b930152565b6000915b6002831061194d57505050565b600190825181526020809101920192019190611940565b9190820391821161197157565b634e487b7160e01b600052601160045260246000fd5b8051156119945760200190565b634e487b7160e01b600052603260045260246000fd5b8051600110156119945760400190565b604051906119c782611802565b602582527f49960de5880e8c687434170f6476605b8fe4aeb9a28632c7995cf3ba831d97636020830152601960f81b6040830152565b600060408051611a0c81611802565b828152826020820152015265ffffffffffff808260a01c168015611a55575b60405192611a3884611802565b6001600160a01b038116845260d01c602084015216604082015290565b5080611a2b565b908151811015611994570160200190565b959391979694929060019681519960258b108015611c55575b611ba19b50611c4c575b604092835192611a9f846117d1565b60158452611acc8360209574113a3cb832911d113bb2b130baba34371733b2ba1160591b87820152611d56565b15611c43575b82611b38611b8892875190611ae682611802565b8882527f4142434445464748494a4b4c4d4e4f505152535455565758595a616263646566888301527f6768696a6b6c6d6e6f707172737475767778797a303132333435363738392d5f8983015261218e565b611b83602e885180936c1131b430b63632b733b2911d1160991b8a830152611b69815180928c602d86019101611708565b8101601160f91b602d82015203600e810184520182611853565b611dc0565b15611c3a575b828085519d848f80965193849201611708565b60009d808f9592869301039060025afa15611c2e57611bfc908251611bec868051809388611bd88184019788815193849201611708565b820190898201520387810184520182611853565b8551928392839251928391611708565b8101039060025afa15611c255750611c1695969751611ccb565b9081611c20575090565b905090565b513d89823e3d90fd5b508251903d90823e3d90fd5b60009950611b8e565b60009a50611ad2565b60009850611a90565b50996020101561199457611ba199611c7660ff60f81b604085015116611d0b565b15611a86565b8015611cc257611c8b906119fd565b65ffffffffffff806040830151164211908115611cb2575b5090516001600160a01b031691565b9050602082015116421038611ca3565b50600090600090565b94939291907f7fffffff800000007fffffffffffffffde737d56d38bcf4279dce5617e3192a88211611d00576117a295612010565b505050505050600090565b6001600160f81b0319600160f81b821601611d5057601f60fb1b600160fb1b821601611d38575b50600190565b600160fc1b90811614611d4b5738611d32565b600090565b50600090565b80519082519260005b838110611d70575050505050600190565b6001908082018083116119715786811090811591611d94575b50611d005701611d5f565b90506001600160f81b0319611db681611dad8589611a5c565b51169286611a5c565b5116141538611d89565b805191805160005b848110611dd9575050505050600190565b60178181018091116119715782811090811591611e08575b50611dfe57600101611dc8565b5050505050600090565b90506001600160f81b0319611e2181611dad8589611a5c565b5116141538611df1565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411611eaf57926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa15611ea35780516001600160a01b03811615611e9a57918190565b50809160019190565b604051903d90823e3d90fd5b50505060009160039190565b6004811015611f2a5780611ecd575050565b60018103611ee75760405163f645eedf60e01b8152600490fd5b60028103611f085760405163fce698f760e01b815260048101839052602490fd5b600314611f125750565b602490604051906335e2f38360e21b82526004820152fd5b634e487b7160e01b600052602160045260246000fd5b903590601e198136030182121561177857018035906001600160401b0382116117785760200191813603831361177857565b90611f7c82611874565b611f896040519182611853565b8281528092611f9a601f1991611874565b0190602036910137565b8151919060418303611fd557611fce92506020820151906060604084015193015160001a90611e2b565b9192909190565b505060009160029190565b3d1561200b573d90611ff182611874565b91611fff6040519384611853565b82523d6000602084013e565b606090565b9493929091946040918251946020978887019784895286868901528160608901528360808901528260a089015260a0885261204a88611838565b6003811015611f2a57806120a55750505050505050600091829151906101005afa90612074611fe0565b9115801561209c575b61209557808280518101031261177857015160011490565b5050600090565b5081511561207d565b6001036120ee575050505050506000918291519073c2b78104907f722dabac4c69f826a522b2754de45afa506120d9611fe0565b90808280518101031261177857015160011490565b9091929496508361214b9496505196612106886117d1565b875287870152845191612118836117d1565b825286820152612141845195878701946304e960d760e01b86526024880152604487019061193c565b608485019061193c565b60a4835260e08301928084106001600160401b038511176117ec576000938493525190735ef35bc6ddb6425a0d43408d7810c3b3180ab58a5afa506120d9611fe0565b9081511561223a578151600281811b93916001600160fe1b0381160361197157600284018094116119715790916121c86003809504611f72565b93602085019284805101926020840194855196600087525b8581106121f1575050505050505290565b8460049101918251600190603f9082828260121c16880101518453828282600c1c16880101518385015382828260061c16880101518885015316850101518682015301906121e0565b5050604051602081018181106001600160401b038211176117ec57604052600081529056fe60808060405234601557610260908161001b8239f35b600080fdfe60806040526004361015610024575b361561001f5734156101eb57600080fd5b6101eb565b6000803560e01c63d1f578941461003b575061000e565b346100af5760403660031901126100af576004356001600160a01b03811681036100ab576024359067ffffffffffffffff908183116100a757366023840112156100a75782600401359182116100a75736602483850101116100a75760246100a4930190610111565b80f35b8380fd5b5080fd5b80fd5b634e487b7160e01b600052604160045260246000fd5b6020808252825181830181905290939260005b8281106100fd57505060409293506000838284010152601f8019910116010190565b8181018601518482016040015285016100db565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8054929390926001600160a01b03166101d95760009382859455816040519283928337810184815203915af43d156101d15767ffffffffffffffff903d8281116101cc5760405192601f8201601f19908116603f01168401908111848210176101cc5760405282523d6000602084013e5b156101ab5750565b604051633018224d60e21b81529081906101c890600483016100c8565b0390fd5b6100b2565b6060906101a3565b60405163c28d69c760e01b8152600490fd5b600036818037808036817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc545af43d82803e15610226573d90f35b3d90fdfea2646970667358221220dbc2ad734b73b998278ed7bcd9ef361791ffe94471f015bd78d961860b387f4664736f6c63430008190033a2646970667358221220e319194c1d0a744414ba3bd5b20764b429173cbe1e768c3e6bd42a5c395ef3eb64736f6c63430008190033",
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

// ComputePredictionHash is a free data retrieval call binding the contract method 0x8bb92411.
//
// Solidity: function computePredictionHash() pure returns(bytes32)
func (_Helper *HelperCaller) ComputePredictionHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "computePredictionHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputePredictionHash is a free data retrieval call binding the contract method 0x8bb92411.
//
// Solidity: function computePredictionHash() pure returns(bytes32)
func (_Helper *HelperSession) ComputePredictionHash() ([32]byte, error) {
	return _Helper.Contract.ComputePredictionHash(&_Helper.CallOpts)
}

// ComputePredictionHash is a free data retrieval call binding the contract method 0x8bb92411.
//
// Solidity: function computePredictionHash() pure returns(bytes32)
func (_Helper *HelperCallerSession) ComputePredictionHash() ([32]byte, error) {
	return _Helper.Contract.ComputePredictionHash(&_Helper.CallOpts)
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
// Solidity: function getFactoryCreateAccountHash(address factory, uint256 _salt, uint256 expireTime, bytes _initializer) view returns(bytes32)
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
// Solidity: function getFactoryCreateAccountHash(address factory, uint256 _salt, uint256 expireTime, bytes _initializer) view returns(bytes32)
func (_Helper *HelperSession) GetFactoryCreateAccountHash(factory common.Address, _salt *big.Int, expireTime *big.Int, _initializer []byte) ([32]byte, error) {
	return _Helper.Contract.GetFactoryCreateAccountHash(&_Helper.CallOpts, factory, _salt, expireTime, _initializer)
}

// GetFactoryCreateAccountHash is a free data retrieval call binding the contract method 0xcaeacff6.
//
// Solidity: function getFactoryCreateAccountHash(address factory, uint256 _salt, uint256 expireTime, bytes _initializer) view returns(bytes32)
func (_Helper *HelperCallerSession) GetFactoryCreateAccountHash(factory common.Address, _salt *big.Int, expireTime *big.Int, _initializer []byte) ([32]byte, error) {
	return _Helper.Contract.GetFactoryCreateAccountHash(&_Helper.CallOpts, factory, _salt, expireTime, _initializer)
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

// GetPackedSig is a free data retrieval call binding the contract method 0x90ea0560.
//
// Solidity: function getPackedSig(uint256 expireTime, bytes sig) pure returns(bytes)
func (_Helper *HelperCaller) GetPackedSig(opts *bind.CallOpts, expireTime *big.Int, sig []byte) ([]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getPackedSig", expireTime, sig)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPackedSig is a free data retrieval call binding the contract method 0x90ea0560.
//
// Solidity: function getPackedSig(uint256 expireTime, bytes sig) pure returns(bytes)
func (_Helper *HelperSession) GetPackedSig(expireTime *big.Int, sig []byte) ([]byte, error) {
	return _Helper.Contract.GetPackedSig(&_Helper.CallOpts, expireTime, sig)
}

// GetPackedSig is a free data retrieval call binding the contract method 0x90ea0560.
//
// Solidity: function getPackedSig(uint256 expireTime, bytes sig) pure returns(bytes)
func (_Helper *HelperCallerSession) GetPackedSig(expireTime *big.Int, sig []byte) ([]byte, error) {
	return _Helper.Contract.GetPackedSig(&_Helper.CallOpts, expireTime, sig)
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

// GetSharelinkHash is a free data retrieval call binding the contract method 0xf5f3b08e.
//
// Solidity: function getSharelinkHash(address payAddress, uint256 chequeID, uint256 recipientAddress, uint256 expireTime) view returns(bytes32)
func (_Helper *HelperCaller) GetSharelinkHash(opts *bind.CallOpts, payAddress common.Address, chequeID *big.Int, recipientAddress *big.Int, expireTime *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Helper.contract.Call(opts, &out, "getSharelinkHash", payAddress, chequeID, recipientAddress, expireTime)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSharelinkHash is a free data retrieval call binding the contract method 0xf5f3b08e.
//
// Solidity: function getSharelinkHash(address payAddress, uint256 chequeID, uint256 recipientAddress, uint256 expireTime) view returns(bytes32)
func (_Helper *HelperSession) GetSharelinkHash(payAddress common.Address, chequeID *big.Int, recipientAddress *big.Int, expireTime *big.Int) ([32]byte, error) {
	return _Helper.Contract.GetSharelinkHash(&_Helper.CallOpts, payAddress, chequeID, recipientAddress, expireTime)
}

// GetSharelinkHash is a free data retrieval call binding the contract method 0xf5f3b08e.
//
// Solidity: function getSharelinkHash(address payAddress, uint256 chequeID, uint256 recipientAddress, uint256 expireTime) view returns(bytes32)
func (_Helper *HelperCallerSession) GetSharelinkHash(payAddress common.Address, chequeID *big.Int, recipientAddress *big.Int, expireTime *big.Int) ([32]byte, error) {
	return _Helper.Contract.GetSharelinkHash(&_Helper.CallOpts, payAddress, chequeID, recipientAddress, expireTime)
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
