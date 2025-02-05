// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pay

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

// IPayChequeParams is an auto generated low-level Go binding around an user-defined struct.
type IPayChequeParams struct {
	ChequeID     *big.Int
	To           common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Expiration   *big.Int
}

// PayMetaData contains all meta data concerning the Pay contract.
var PayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AccessError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CancleError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ChequeExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ChequeExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ChequeNotExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"},{\"internalType\":\"enumChequeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ChequeStatusError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ClaimAddressError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"ETHAmountError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedHalt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedHalt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"refundTime\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"cancelTime\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"max\",\"type\":\"uint128\"}],\"name\":\"ExpirationTimeNotValid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"InvalidAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Rejected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldExpirationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExpirationTime\",\"type\":\"uint256\"}],\"name\":\"ChangeMaxExpirationTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumChequeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ChequeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"ChequeSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"HaltStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"WhitelistTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"WhitelistTokenRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accountFactory\",\"outputs\":[{\"internalType\":\"contractIAccountFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"addWhitelistedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"chequeIDs\",\"type\":\"uint256[]\"}],\"name\":\"batchRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"expirationTime\",\"type\":\"uint128\"}],\"name\":\"changeMaxExpirationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cheques\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"enumChequeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"halt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"halted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelistedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxExpirationTime\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"removeWhitelistedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"internalType\":\"structIPay.ChequeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chequeID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"sharelinkClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unhalt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c03461014957601f611cf038819003918201601f19168301916001600160401b0383118484101761014e578084926060946040528339810103126101495761004781610164565b61005f604061005860208501610164565b9301610164565b600054926001600160a01b039290919083821680156101305760018593620100008260b01b039060101b16818060b01b0319881617176000558260016040519761ffff19161760101c167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3600180551660a05216608052600480546001600160801b03191662278d00179055611b77908161017982396080518181816103f60152610bc2015260a05181818161027f015281816107c001528181610c6501528181610e1b01526110c20152f35b604051631e4fbdf760e01b815260006004820152602490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101495756fe60406080815260049081361015610029575b361561001c57600080fd5b5163efde6dcf60e01b8152fd5b6000803560e01c806305424c541461123b57806306fdde03146111fc578063278ecde1146111d7578063379607f5146110885780633f4ba83a1461101c57806340e58ee514610dde57806351208fe414610d6857806354fd4d5014610d275780635c975abb14610d025780635ed7ca5b14610c94578063687cd9c114610c50578063715018a614610bf157806379502c5514610bad5780638456cb5914610b525780638da5cb5b14610b2657806396e1274d1461078c578063ab37f4861461074e578063b9b8af0b1461072b578063bcec454f146106a7578063cb3e64fd14610639578063cd698e69146105ac578063ddf2ff7e14610515578063e4884f8d14610210578063f2fde38b146101705763f541b382146101485750610011565b3461016d578060031936011261016d57506001600160801b0360209254169051908152f35b80fd5b50913461020c57602036600319011261020c5761018b6113e7565b610193611872565b6001600160a01b038181169390929084156101f6575050835462010000600160b01b03198116601092831b62010000600160b01b0316178555901c167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b51631e4fbdf760e01b8152908101859052602490fd5b8280fd5b5090346105115780600319360112610511578235602491823567ffffffffffffffff9384821161050d573660238301121561050d578187013594851161050d573681868401011161050d5782516308f329a760e21b815233888201526020926001600160a01b039291848184817f000000000000000000000000000000000000000000000000000000000000000088165afa9081156105035789916104e6575b50156104d2576102be61173b565b6102c661156d565b8588526002845284882096838651986102de8a6112f3565b818154168a528160018201541690878b0191825282600282015416898c0152600381015460608c01528c81015460808c01526005015460ff1660a08b0190610325916114f0565b5116806104a657508085116104a25782820135428110610485576103d58786948c8997958c6103de968551908b8201924684523088840152606083015233608083015260a082015260a0815261037a816112f3565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008352601c52603c8220926044601f198301916103c56103bc846114bc565b97519788611341565b828752018a86013783015261199e565b909291926119da565b8651948593849263729d3f5760e11b8452168c8301527f0000000000000000000000000000000000000000000000000000000000000000165afa91821561047b57869261044e575b50501561043f5750906104389161175c565b6001805580f35b51631eb49d6d60e11b81528490fd5b61046d9250803d10610474575b6104658183611341565b8101906114d8565b3880610426565b503d61045b565b83513d88823e3d90fd5b865163aa2fd92560e01b8152808c01919091524281850152604490fd5b8880fd5b8651633586e9ab60e01b81526001600160a01b03909116818c0190815233602082015281906040010390fd5b5083516325abcd9160e11b815233818a0152fd5b6104fd9150853d8711610474576104658183611341565b386102b0565b86513d8b823e3d90fd5b8580fd5b5080fd5b5091903461020c57602036600319011261020c578135906001600160801b03908183168093036105a857610547611872565b8215610599577fd764d769918360f43f93e481dd41732ac1e5a40e2407dd7c569651126d82160f9084549281519084168152846020820152a16fffffffffffffffffffffffffffffffff191617905580f35b5163162908e360e11b81528390fd5b8480fd5b509034610511576105bc36611402565b906105c5611872565b8151835b81811061060757847f9f9d9f44bc581f34670150874e4db46bb05da1753d75c6b5c7b952fcdaa96f4a61060186865191829182611528565b0390a180f35b6001600160a01b036106198286611492565b511685526003602052828520805460ff19166001908117909155016105c9565b50913461020c578260031936011261020c57610653611872565b825490600160ff83161461069957507fd57639a4f3c4c1059ca3fdd5175e8ca8ebf6f1eee17c2962389afa8bfd9b34e691600160209260ff19161784555160018152a180f35b825163b63be3c160e01b8152fd5b509034610511576106b736611402565b906106c0611872565b8151835b8181106106fc57847fa50261031b405a768d9301886abe947b7a26918759bc9b5c89a4398b7dcff16e61060186865191829182611528565b6001906001600160a01b036107118287611492565b511686526003602052838620805460ff19169055016106c4565b509034610511578160031936011261051157600260ff6020935416149051908152f35b5090346105115760203660031901126105115760209160ff9082906001600160a01b036107796113e7565b1681526003855220541690519015158152f35b509160a036600319011261020c5781516308f329a760e21b815233828201526024926020926001600160a01b0392848187817f000000000000000000000000000000000000000000000000000000000000000088165afa908115610b1c578791610aff575b5015610aea57826108006114fc565b168087526003855260ff838820541615610ad7575061081d61173b565b600260ff87541614610ac95761083161156d565b606435918215610abc5773eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee846108596114fc565b1603610a3457823403610a18575b608435918260801c6001600160801b0390818516918084541681421601818111610a0657169182811180156109fd575b6109df57505050803596878952600287526003838a2001546109cb579060019160058986888b8e8c8a6108c8611512565b9260026108d36114fc565b958351996108e08b6112f3565b338b5285828c019716875285858c019816885260608b0198895260808b01998a5260a08b019e8f5283525220955116906bffffffffffffffffffffffff60a01b91828754161786558d60018701915116828254161790558c600286019251169082541617905551600383015551848201550192519160058310156109ba57505060ff80198354169116179055610974611512565b9261097d6114fc565b958583519716875286015284015216917fd18cb5c9f5533f163598578532d03f4b1fb74decf1fb9a23f7d16defc21dda0560603393a46001805580f35b634e487b7160e01b8b526021905289fd5b915163e9860ee760e01b8152908101879052fd5b89906064955194632326787360e11b86528501528301526044820152fd5b50828211610897565b634e487b7160e01b8c52601185528a8cfd5b51634dcbc7ab60e01b8152908101919091523481850152604490fd5b34610aa157610a416114fc565b81516323b872dd60e01b878201523388820152306044820152606480820186905281529060a0820167ffffffffffffffff811183821017610a8f578352610a8a91908616611a49565b610867565b634e487b7160e01b8a5260418552888afd5b51634dcbc7ab60e01b81529081018690523481860152604490fd5b51634ff64a9f60e01b8152fd5b9051630248e51360e61b8152fd5b8592519163961c9a4f60e01b8352820152fd5b849151906325abcd9160e11b82523390820152fd5b610b169150853d8711610474576104658183611341565b386107f1565b83513d89823e3d90fd5b5090346105115781600319360112610511579054905160109190911c6001600160a01b03168152602090f35b50903461051157816003193601126105115760207f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25891610b90611872565b610b9861173b565b835461ff00191661010017845551338152a180f35b509034610511578160031936011261051157517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b503461016d578060031936011261016d57610c0a611872565b805462010000600160b01b031981168255819060101c6001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b509034610511578160031936011261051157517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50913461020c578260031936011261020c57610cae611872565b825490600260ff831614610cf457507fd57639a4f3c4c1059ca3fdd5175e8ca8ebf6f1eee17c2962389afa8bfd9b34e691600260209260ff19161784555160028152a180f35b8251630248e51360e61b8152fd5b50903461051157816003193601126105115760ff6020925460081c1690519015158152f35b5090346105115781600319360112610511578051610d6491610d4882611325565b60058252640312e302e360dc1b6020830152519182918261137b565b0390f35b50913461020c57602036600319011261020c57808260c094610ddc9335815260026020522060018060a01b0391828254169483600184015416936002840154169060ff60056003860154948601549501541694815197885260208801528601526060850152608084015260a08301906113c4565bf35b5091903461020c57602090816003193601126110185780516308f329a760e21b81523381850152602492843592916001600160a01b0390828187817f000000000000000000000000000000000000000000000000000000000000000086165afa90811561100e578891610ff1575b5015610fdc57610e5a61156d565b8387526002825282872094835196610e71886112f3565b82875416885282600188015416918489019283528360028901541698868101998a5260038901549860608201998a5260ff6005858301549260808501938452015416610ec160a0840191826114f0565b80516005811015610fca57600103610f8557505160801c4210610f7557843391511603610f5857505090806002949392610f02828a511689519033906118a1565b868a528584526005858b20018660ff19825416179055511696511694519082519586528501528301527f4617b59973e052c439b84edb2deaedd5c5dad563af6952409e72c65194ba4a4c60603393a46001805580f35b855163627d1d5960e01b8152918201879052339082015260449150fd5b865163c8a0f51f60e01b81528390fd5b848d8b8b879451926005841015610fb85750516352cdd5b160e11b815292830152604492610fb691908301906113c4565bfd5b634e487b7160e01b8152602185528590fd5b634e487b7160e01b8e5260218652848efd5b82516325abcd9160e11b815233818801528590fd5b6110089150833d8511610474576104658183611341565b38610e4c565b84513d8a823e3d90fd5b8380fd5b50913461020c578260031936011261020c57611036611872565b82549060ff8260081c161561107a575061ff0019168255513381527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa90602090a180f35b8251638dfc202b60e01b8152fd5b509034610511576020908160031936011261020c5780516308f329a760e21b81523381860152843591906001600160a01b039084816024817f000000000000000000000000000000000000000000000000000000000000000086165afa90811561047b5786916111ba575b50156111a45761110161173b565b61110961156d565b828552600284528185209361116460ff6005855197611127896112f3565b8581541689528560018201541694890194855285600282015416878a0152600381015460608a01528a81015460808a015201541660a087016114f0565b5116903382036111795784610438858561175c565b51633586e9ab60e01b81526001600160a01b0390911681860190815233602082015281906040010390fd5b81516325abcd9160e11b81523381880152602490fd5b6111d19150853d8711610474576104658183611341565b386110f3565b50823461051157602036600319011261051157610438906111f661156d565b35611590565b5090346105115781600319360112610511578051610d649161121d82611325565b60078252666f6b782e70617960c81b6020830152519182918261137b565b50913461020c5760209060206003193601126110185780359067ffffffffffffffff82116105a857366023830112156105a85781013561128661127d82611363565b94519485611341565b8084526024602085019160051b8301019136831161050d57602401905b8282106112e45785856112b461156d565b805190825b8281106112c857836001805580f35b806112de6112d860019385611492565b51611590565b016112b9565b813581529083019083016112a3565b60c0810190811067ffffffffffffffff82111761130f57604052565b634e487b7160e01b600052604160045260246000fd5b6040810190811067ffffffffffffffff82111761130f57604052565b90601f8019910116810190811067ffffffffffffffff82111761130f57604052565b67ffffffffffffffff811161130f5760051b60200190565b6020808252825181830181905290939260005b8281106113b057505060409293506000838284010152601f8019910116010190565b81810186015184820160400152850161138e565b9060058210156113d15752565b634e487b7160e01b600052602160045260246000fd5b600435906001600160a01b03821682036113fd57565b600080fd5b6020806003198301126113fd576004359167ffffffffffffffff83116113fd57806023840112156113fd57826004013561143b81611363565b936114496040519586611341565b8185526024602086019260051b8201019283116113fd57602401905b828210611473575050505090565b81356001600160a01b03811681036113fd578152908301908301611465565b80518210156114a65760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b67ffffffffffffffff811161130f57601f01601f191660200190565b908160209103126113fd575180151581036113fd5790565b60058210156113d15752565b6044356001600160a01b03811681036113fd5790565b6024356001600160a01b03811681036113fd5790565b602090602060408183019282815285518094520193019160005b828110611550575050505090565b83516001600160a01b031685529381019392810192600101611542565b60026001541461157e576002600155565b604051633ee5aeb560e01b8152600490fd5b6000908082526002602052604090818320928251936115ae856112f3565b60018060a01b039384825416865284600183015416946020870195865280600284015416968281019788526003840154946060820195865260049460ff600587830154926080860193845201541661160a60a0850191826114f0565b80516005811015611728576001036116e45750516001600160801b0316428111806116d7575b6116a55750918080859360057f4617b59973e052c439b84edb2deaedd5c5dad563af6952409e72c65194ba4a4c9997606099976116778f86905116868651168c51916118a1565b8c8152600260205220018660ff198254161790555116985116985116925181519384526020840152820152a4565b8451632cf66e4960e01b81528087018981526001600160801b0390921660208301524260408301529081906060010390fd5b5083895116331415611630565b51869089878560058410156117155781516352cdd5b160e11b8152808601849052604490610fb660248201876113c4565b634e487b7160e01b815260218552602490fd5b634e487b7160e01b845260218852602484fd5b60ff60005460081c1661174a57565b60405163d93c066560e01b8152600490fd5b9060a08101805160058110156113d15760010361184257506001600160801b036080820151164281106118135750604081018051606080840180519395937f4617b59973e052c439b84edb2deaedd5c5dad563af6952409e72c65194ba4a4c936001600160a01b03916117d291339084166118a1565b8460005260026020526005604060002001600360ff1982541617905580602081885116970151169651169051604051918252602082015260036040820152a4565b60405163020fb0ed60e41b815260048101939093526001600160801b0316602483015250426044820152606490fd5b9050519060058210156113d157610fb6604492604051926352cdd5b160e11b8452600484015260248301906113c4565b60005460101c6001600160a01b0316330361188957565b60405163118cdaa760e01b8152336004820152602490fd5b6001600160a01b039081169073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee82036119555750506040516020810181811067ffffffffffffffff82111761130f5760405260008080948194828095525af1903d1561194f573d90611906826114bc565b916119146040519384611341565b825260203d92013e5b1561192457565b60405162461bcd60e51b815260206004820152600360248201526253544560e81b6044820152606490fd5b5061191d565b60405163a9059cbb60e01b602082015292166024830152604480830193909352918152608081019167ffffffffffffffff83118284101761130f5761199c92604052611a49565b565b81519190604183036119cf576119c892506020820151906060604084015193015160001a90611ab1565b9192909190565b505060009160029190565b60048110156113d157806119ec575050565b60018103611a065760405163f645eedf60e01b8152600490fd5b60028103611a275760405163fce698f760e01b815260048101839052602490fd5b600314611a315750565b602490604051906335e2f38360e21b82526004820152fd5b906000602091828151910182855af115611aa5576000513d611a9c57506001600160a01b0381163b155b611a7a5750565b604051635274afe760e01b81526001600160a01b039091166004820152602490fd5b60011415611a73565b6040513d6000823e3d90fd5b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411611b3557926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa15611b295780516001600160a01b03811615611b2057918190565b50809160019190565b604051903d90823e3d90fd5b5050506000916003919056fea26469706673582212201256e4031fea1732a37c3a50c8da5f778f8debc49c565cee20bfe75073980d8364736f6c63430008190033",
}

// PayABI is the input ABI used to generate the binding from.
// Deprecated: Use PayMetaData.ABI instead.
var PayABI = PayMetaData.ABI

// PayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PayMetaData.Bin instead.
var PayBin = PayMetaData.Bin

// DeployPay deploys a new Ethereum contract, binding an instance of Pay to it.
func DeployPay(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _config common.Address, _initialOwner common.Address) (common.Address, *types.Transaction, *Pay, error) {
	parsed, err := PayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PayBin), backend, _factory, _config, _initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pay{PayCaller: PayCaller{contract: contract}, PayTransactor: PayTransactor{contract: contract}, PayFilterer: PayFilterer{contract: contract}}, nil
}

// Pay is an auto generated Go binding around an Ethereum contract.
type Pay struct {
	PayCaller     // Read-only binding to the contract
	PayTransactor // Write-only binding to the contract
	PayFilterer   // Log filterer for contract events
}

// PayCaller is an auto generated read-only Go binding around an Ethereum contract.
type PayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaySession struct {
	Contract     *Pay              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PayCallerSession struct {
	Contract *PayCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PayTransactorSession struct {
	Contract     *PayTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PayRaw is an auto generated low-level Go binding around an Ethereum contract.
type PayRaw struct {
	Contract *Pay // Generic contract binding to access the raw methods on
}

// PayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PayCallerRaw struct {
	Contract *PayCaller // Generic read-only contract binding to access the raw methods on
}

// PayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PayTransactorRaw struct {
	Contract *PayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPay creates a new instance of Pay, bound to a specific deployed contract.
func NewPay(address common.Address, backend bind.ContractBackend) (*Pay, error) {
	contract, err := bindPay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pay{PayCaller: PayCaller{contract: contract}, PayTransactor: PayTransactor{contract: contract}, PayFilterer: PayFilterer{contract: contract}}, nil
}

// NewPayCaller creates a new read-only instance of Pay, bound to a specific deployed contract.
func NewPayCaller(address common.Address, caller bind.ContractCaller) (*PayCaller, error) {
	contract, err := bindPay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PayCaller{contract: contract}, nil
}

// NewPayTransactor creates a new write-only instance of Pay, bound to a specific deployed contract.
func NewPayTransactor(address common.Address, transactor bind.ContractTransactor) (*PayTransactor, error) {
	contract, err := bindPay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PayTransactor{contract: contract}, nil
}

// NewPayFilterer creates a new log filterer instance of Pay, bound to a specific deployed contract.
func NewPayFilterer(address common.Address, filterer bind.ContractFilterer) (*PayFilterer, error) {
	contract, err := bindPay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PayFilterer{contract: contract}, nil
}

// bindPay binds a generic wrapper to an already deployed contract.
func bindPay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pay *PayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pay.Contract.PayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pay *PayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pay.Contract.PayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pay *PayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pay.Contract.PayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pay *PayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pay *PayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pay *PayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pay.Contract.contract.Transact(opts, method, params...)
}

// AccountFactory is a free data retrieval call binding the contract method 0x687cd9c1.
//
// Solidity: function accountFactory() view returns(address)
func (_Pay *PayCaller) AccountFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "accountFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountFactory is a free data retrieval call binding the contract method 0x687cd9c1.
//
// Solidity: function accountFactory() view returns(address)
func (_Pay *PaySession) AccountFactory() (common.Address, error) {
	return _Pay.Contract.AccountFactory(&_Pay.CallOpts)
}

// AccountFactory is a free data retrieval call binding the contract method 0x687cd9c1.
//
// Solidity: function accountFactory() view returns(address)
func (_Pay *PayCallerSession) AccountFactory() (common.Address, error) {
	return _Pay.Contract.AccountFactory(&_Pay.CallOpts)
}

// Cheques is a free data retrieval call binding the contract method 0x51208fe4.
//
// Solidity: function cheques(uint256 ) view returns(address from, address to, address tokenAddress, uint256 amount, uint256 expiration, uint8 status)
func (_Pay *PayCaller) Cheques(opts *bind.CallOpts, arg0 *big.Int) (struct {
	From         common.Address
	To           common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Expiration   *big.Int
	Status       uint8
}, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "cheques", arg0)

	outstruct := new(struct {
		From         common.Address
		To           common.Address
		TokenAddress common.Address
		Amount       *big.Int
		Expiration   *big.Int
		Status       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.From = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.To = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Expiration = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// Cheques is a free data retrieval call binding the contract method 0x51208fe4.
//
// Solidity: function cheques(uint256 ) view returns(address from, address to, address tokenAddress, uint256 amount, uint256 expiration, uint8 status)
func (_Pay *PaySession) Cheques(arg0 *big.Int) (struct {
	From         common.Address
	To           common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Expiration   *big.Int
	Status       uint8
}, error) {
	return _Pay.Contract.Cheques(&_Pay.CallOpts, arg0)
}

// Cheques is a free data retrieval call binding the contract method 0x51208fe4.
//
// Solidity: function cheques(uint256 ) view returns(address from, address to, address tokenAddress, uint256 amount, uint256 expiration, uint8 status)
func (_Pay *PayCallerSession) Cheques(arg0 *big.Int) (struct {
	From         common.Address
	To           common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Expiration   *big.Int
	Status       uint8
}, error) {
	return _Pay.Contract.Cheques(&_Pay.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_Pay *PayCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "config")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_Pay *PaySession) Config() (common.Address, error) {
	return _Pay.Contract.Config(&_Pay.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_Pay *PayCallerSession) Config() (common.Address, error) {
	return _Pay.Contract.Config(&_Pay.CallOpts)
}

// Halted is a free data retrieval call binding the contract method 0xb9b8af0b.
//
// Solidity: function halted() view returns(bool)
func (_Pay *PayCaller) Halted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "halted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Halted is a free data retrieval call binding the contract method 0xb9b8af0b.
//
// Solidity: function halted() view returns(bool)
func (_Pay *PaySession) Halted() (bool, error) {
	return _Pay.Contract.Halted(&_Pay.CallOpts)
}

// Halted is a free data retrieval call binding the contract method 0xb9b8af0b.
//
// Solidity: function halted() view returns(bool)
func (_Pay *PayCallerSession) Halted() (bool, error) {
	return _Pay.Contract.Halted(&_Pay.CallOpts)
}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address ) view returns(bool)
func (_Pay *PayCaller) IsWhitelistedToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "isWhitelistedToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address ) view returns(bool)
func (_Pay *PaySession) IsWhitelistedToken(arg0 common.Address) (bool, error) {
	return _Pay.Contract.IsWhitelistedToken(&_Pay.CallOpts, arg0)
}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address ) view returns(bool)
func (_Pay *PayCallerSession) IsWhitelistedToken(arg0 common.Address) (bool, error) {
	return _Pay.Contract.IsWhitelistedToken(&_Pay.CallOpts, arg0)
}

// MaxExpirationTime is a free data retrieval call binding the contract method 0xf541b382.
//
// Solidity: function maxExpirationTime() view returns(uint128)
func (_Pay *PayCaller) MaxExpirationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "maxExpirationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExpirationTime is a free data retrieval call binding the contract method 0xf541b382.
//
// Solidity: function maxExpirationTime() view returns(uint128)
func (_Pay *PaySession) MaxExpirationTime() (*big.Int, error) {
	return _Pay.Contract.MaxExpirationTime(&_Pay.CallOpts)
}

// MaxExpirationTime is a free data retrieval call binding the contract method 0xf541b382.
//
// Solidity: function maxExpirationTime() view returns(uint128)
func (_Pay *PayCallerSession) MaxExpirationTime() (*big.Int, error) {
	return _Pay.Contract.MaxExpirationTime(&_Pay.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Pay *PayCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Pay *PaySession) Name() (string, error) {
	return _Pay.Contract.Name(&_Pay.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Pay *PayCallerSession) Name() (string, error) {
	return _Pay.Contract.Name(&_Pay.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pay *PayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pay *PaySession) Owner() (common.Address, error) {
	return _Pay.Contract.Owner(&_Pay.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pay *PayCallerSession) Owner() (common.Address, error) {
	return _Pay.Contract.Owner(&_Pay.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pay *PayCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pay *PaySession) Paused() (bool, error) {
	return _Pay.Contract.Paused(&_Pay.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pay *PayCallerSession) Paused() (bool, error) {
	return _Pay.Contract.Paused(&_Pay.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Pay *PayCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Pay *PaySession) Version() (string, error) {
	return _Pay.Contract.Version(&_Pay.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Pay *PayCallerSession) Version() (string, error) {
	return _Pay.Contract.Version(&_Pay.CallOpts)
}

// AddWhitelistedTokens is a paid mutator transaction binding the contract method 0xcd698e69.
//
// Solidity: function addWhitelistedTokens(address[] tokens) returns()
func (_Pay *PayTransactor) AddWhitelistedTokens(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "addWhitelistedTokens", tokens)
}

// AddWhitelistedTokens is a paid mutator transaction binding the contract method 0xcd698e69.
//
// Solidity: function addWhitelistedTokens(address[] tokens) returns()
func (_Pay *PaySession) AddWhitelistedTokens(tokens []common.Address) (*types.Transaction, error) {
	return _Pay.Contract.AddWhitelistedTokens(&_Pay.TransactOpts, tokens)
}

// AddWhitelistedTokens is a paid mutator transaction binding the contract method 0xcd698e69.
//
// Solidity: function addWhitelistedTokens(address[] tokens) returns()
func (_Pay *PayTransactorSession) AddWhitelistedTokens(tokens []common.Address) (*types.Transaction, error) {
	return _Pay.Contract.AddWhitelistedTokens(&_Pay.TransactOpts, tokens)
}

// BatchRefund is a paid mutator transaction binding the contract method 0x05424c54.
//
// Solidity: function batchRefund(uint256[] chequeIDs) returns()
func (_Pay *PayTransactor) BatchRefund(opts *bind.TransactOpts, chequeIDs []*big.Int) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "batchRefund", chequeIDs)
}

// BatchRefund is a paid mutator transaction binding the contract method 0x05424c54.
//
// Solidity: function batchRefund(uint256[] chequeIDs) returns()
func (_Pay *PaySession) BatchRefund(chequeIDs []*big.Int) (*types.Transaction, error) {
	return _Pay.Contract.BatchRefund(&_Pay.TransactOpts, chequeIDs)
}

// BatchRefund is a paid mutator transaction binding the contract method 0x05424c54.
//
// Solidity: function batchRefund(uint256[] chequeIDs) returns()
func (_Pay *PayTransactorSession) BatchRefund(chequeIDs []*big.Int) (*types.Transaction, error) {
	return _Pay.Contract.BatchRefund(&_Pay.TransactOpts, chequeIDs)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 chequeID) returns()
func (_Pay *PayTransactor) Cancel(opts *bind.TransactOpts, chequeID *big.Int) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "cancel", chequeID)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 chequeID) returns()
func (_Pay *PaySession) Cancel(chequeID *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.Cancel(&_Pay.TransactOpts, chequeID)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 chequeID) returns()
func (_Pay *PayTransactorSession) Cancel(chequeID *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.Cancel(&_Pay.TransactOpts, chequeID)
}

// ChangeMaxExpirationTime is a paid mutator transaction binding the contract method 0xddf2ff7e.
//
// Solidity: function changeMaxExpirationTime(uint128 expirationTime) returns()
func (_Pay *PayTransactor) ChangeMaxExpirationTime(opts *bind.TransactOpts, expirationTime *big.Int) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "changeMaxExpirationTime", expirationTime)
}

// ChangeMaxExpirationTime is a paid mutator transaction binding the contract method 0xddf2ff7e.
//
// Solidity: function changeMaxExpirationTime(uint128 expirationTime) returns()
func (_Pay *PaySession) ChangeMaxExpirationTime(expirationTime *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.ChangeMaxExpirationTime(&_Pay.TransactOpts, expirationTime)
}

// ChangeMaxExpirationTime is a paid mutator transaction binding the contract method 0xddf2ff7e.
//
// Solidity: function changeMaxExpirationTime(uint128 expirationTime) returns()
func (_Pay *PayTransactorSession) ChangeMaxExpirationTime(expirationTime *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.ChangeMaxExpirationTime(&_Pay.TransactOpts, expirationTime)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 chequeID) returns()
func (_Pay *PayTransactor) Claim(opts *bind.TransactOpts, chequeID *big.Int) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "claim", chequeID)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 chequeID) returns()
func (_Pay *PaySession) Claim(chequeID *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.Claim(&_Pay.TransactOpts, chequeID)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 chequeID) returns()
func (_Pay *PayTransactorSession) Claim(chequeID *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.Claim(&_Pay.TransactOpts, chequeID)
}

// Halt is a paid mutator transaction binding the contract method 0x5ed7ca5b.
//
// Solidity: function halt() returns()
func (_Pay *PayTransactor) Halt(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "halt")
}

// Halt is a paid mutator transaction binding the contract method 0x5ed7ca5b.
//
// Solidity: function halt() returns()
func (_Pay *PaySession) Halt() (*types.Transaction, error) {
	return _Pay.Contract.Halt(&_Pay.TransactOpts)
}

// Halt is a paid mutator transaction binding the contract method 0x5ed7ca5b.
//
// Solidity: function halt() returns()
func (_Pay *PayTransactorSession) Halt() (*types.Transaction, error) {
	return _Pay.Contract.Halt(&_Pay.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pay *PayTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pay *PaySession) Pause() (*types.Transaction, error) {
	return _Pay.Contract.Pause(&_Pay.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pay *PayTransactorSession) Pause() (*types.Transaction, error) {
	return _Pay.Contract.Pause(&_Pay.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 chequeID) returns()
func (_Pay *PayTransactor) Refund(opts *bind.TransactOpts, chequeID *big.Int) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "refund", chequeID)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 chequeID) returns()
func (_Pay *PaySession) Refund(chequeID *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.Refund(&_Pay.TransactOpts, chequeID)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 chequeID) returns()
func (_Pay *PayTransactorSession) Refund(chequeID *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.Refund(&_Pay.TransactOpts, chequeID)
}

// RemoveWhitelistedTokens is a paid mutator transaction binding the contract method 0xbcec454f.
//
// Solidity: function removeWhitelistedTokens(address[] tokens) returns()
func (_Pay *PayTransactor) RemoveWhitelistedTokens(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "removeWhitelistedTokens", tokens)
}

// RemoveWhitelistedTokens is a paid mutator transaction binding the contract method 0xbcec454f.
//
// Solidity: function removeWhitelistedTokens(address[] tokens) returns()
func (_Pay *PaySession) RemoveWhitelistedTokens(tokens []common.Address) (*types.Transaction, error) {
	return _Pay.Contract.RemoveWhitelistedTokens(&_Pay.TransactOpts, tokens)
}

// RemoveWhitelistedTokens is a paid mutator transaction binding the contract method 0xbcec454f.
//
// Solidity: function removeWhitelistedTokens(address[] tokens) returns()
func (_Pay *PayTransactorSession) RemoveWhitelistedTokens(tokens []common.Address) (*types.Transaction, error) {
	return _Pay.Contract.RemoveWhitelistedTokens(&_Pay.TransactOpts, tokens)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pay *PayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pay *PaySession) RenounceOwnership() (*types.Transaction, error) {
	return _Pay.Contract.RenounceOwnership(&_Pay.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pay *PayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pay.Contract.RenounceOwnership(&_Pay.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x96e1274d.
//
// Solidity: function send((uint256,address,address,uint256,uint256) params) payable returns()
func (_Pay *PayTransactor) Send(opts *bind.TransactOpts, params IPayChequeParams) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "send", params)
}

// Send is a paid mutator transaction binding the contract method 0x96e1274d.
//
// Solidity: function send((uint256,address,address,uint256,uint256) params) payable returns()
func (_Pay *PaySession) Send(params IPayChequeParams) (*types.Transaction, error) {
	return _Pay.Contract.Send(&_Pay.TransactOpts, params)
}

// Send is a paid mutator transaction binding the contract method 0x96e1274d.
//
// Solidity: function send((uint256,address,address,uint256,uint256) params) payable returns()
func (_Pay *PayTransactorSession) Send(params IPayChequeParams) (*types.Transaction, error) {
	return _Pay.Contract.Send(&_Pay.TransactOpts, params)
}

// SharelinkClaim is a paid mutator transaction binding the contract method 0xe4884f8d.
//
// Solidity: function sharelinkClaim(uint256 chequeID, bytes signature) returns()
func (_Pay *PayTransactor) SharelinkClaim(opts *bind.TransactOpts, chequeID *big.Int, signature []byte) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "sharelinkClaim", chequeID, signature)
}

// SharelinkClaim is a paid mutator transaction binding the contract method 0xe4884f8d.
//
// Solidity: function sharelinkClaim(uint256 chequeID, bytes signature) returns()
func (_Pay *PaySession) SharelinkClaim(chequeID *big.Int, signature []byte) (*types.Transaction, error) {
	return _Pay.Contract.SharelinkClaim(&_Pay.TransactOpts, chequeID, signature)
}

// SharelinkClaim is a paid mutator transaction binding the contract method 0xe4884f8d.
//
// Solidity: function sharelinkClaim(uint256 chequeID, bytes signature) returns()
func (_Pay *PayTransactorSession) SharelinkClaim(chequeID *big.Int, signature []byte) (*types.Transaction, error) {
	return _Pay.Contract.SharelinkClaim(&_Pay.TransactOpts, chequeID, signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pay *PayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pay *PaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pay.Contract.TransferOwnership(&_Pay.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pay *PayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pay.Contract.TransferOwnership(&_Pay.TransactOpts, newOwner)
}

// Unhalt is a paid mutator transaction binding the contract method 0xcb3e64fd.
//
// Solidity: function unhalt() returns()
func (_Pay *PayTransactor) Unhalt(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "unhalt")
}

// Unhalt is a paid mutator transaction binding the contract method 0xcb3e64fd.
//
// Solidity: function unhalt() returns()
func (_Pay *PaySession) Unhalt() (*types.Transaction, error) {
	return _Pay.Contract.Unhalt(&_Pay.TransactOpts)
}

// Unhalt is a paid mutator transaction binding the contract method 0xcb3e64fd.
//
// Solidity: function unhalt() returns()
func (_Pay *PayTransactorSession) Unhalt() (*types.Transaction, error) {
	return _Pay.Contract.Unhalt(&_Pay.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pay *PayTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pay *PaySession) Unpause() (*types.Transaction, error) {
	return _Pay.Contract.Unpause(&_Pay.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pay *PayTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pay.Contract.Unpause(&_Pay.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pay *PayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pay.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pay *PaySession) Receive() (*types.Transaction, error) {
	return _Pay.Contract.Receive(&_Pay.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pay *PayTransactorSession) Receive() (*types.Transaction, error) {
	return _Pay.Contract.Receive(&_Pay.TransactOpts)
}

// PayChangeMaxExpirationTimeIterator is returned from FilterChangeMaxExpirationTime and is used to iterate over the raw logs and unpacked data for ChangeMaxExpirationTime events raised by the Pay contract.
type PayChangeMaxExpirationTimeIterator struct {
	Event *PayChangeMaxExpirationTime // Event containing the contract specifics and raw log

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
func (it *PayChangeMaxExpirationTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayChangeMaxExpirationTime)
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
		it.Event = new(PayChangeMaxExpirationTime)
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
func (it *PayChangeMaxExpirationTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayChangeMaxExpirationTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayChangeMaxExpirationTime represents a ChangeMaxExpirationTime event raised by the Pay contract.
type PayChangeMaxExpirationTime struct {
	OldExpirationTime *big.Int
	NewExpirationTime *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChangeMaxExpirationTime is a free log retrieval operation binding the contract event 0xd764d769918360f43f93e481dd41732ac1e5a40e2407dd7c569651126d82160f.
//
// Solidity: event ChangeMaxExpirationTime(uint256 oldExpirationTime, uint256 newExpirationTime)
func (_Pay *PayFilterer) FilterChangeMaxExpirationTime(opts *bind.FilterOpts) (*PayChangeMaxExpirationTimeIterator, error) {

	logs, sub, err := _Pay.contract.FilterLogs(opts, "ChangeMaxExpirationTime")
	if err != nil {
		return nil, err
	}
	return &PayChangeMaxExpirationTimeIterator{contract: _Pay.contract, event: "ChangeMaxExpirationTime", logs: logs, sub: sub}, nil
}

// WatchChangeMaxExpirationTime is a free log subscription operation binding the contract event 0xd764d769918360f43f93e481dd41732ac1e5a40e2407dd7c569651126d82160f.
//
// Solidity: event ChangeMaxExpirationTime(uint256 oldExpirationTime, uint256 newExpirationTime)
func (_Pay *PayFilterer) WatchChangeMaxExpirationTime(opts *bind.WatchOpts, sink chan<- *PayChangeMaxExpirationTime) (event.Subscription, error) {

	logs, sub, err := _Pay.contract.WatchLogs(opts, "ChangeMaxExpirationTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayChangeMaxExpirationTime)
				if err := _Pay.contract.UnpackLog(event, "ChangeMaxExpirationTime", log); err != nil {
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

// ParseChangeMaxExpirationTime is a log parse operation binding the contract event 0xd764d769918360f43f93e481dd41732ac1e5a40e2407dd7c569651126d82160f.
//
// Solidity: event ChangeMaxExpirationTime(uint256 oldExpirationTime, uint256 newExpirationTime)
func (_Pay *PayFilterer) ParseChangeMaxExpirationTime(log types.Log) (*PayChangeMaxExpirationTime, error) {
	event := new(PayChangeMaxExpirationTime)
	if err := _Pay.contract.UnpackLog(event, "ChangeMaxExpirationTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayChequeEventIterator is returned from FilterChequeEvent and is used to iterate over the raw logs and unpacked data for ChequeEvent events raised by the Pay contract.
type PayChequeEventIterator struct {
	Event *PayChequeEvent // Event containing the contract specifics and raw log

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
func (it *PayChequeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayChequeEvent)
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
		it.Event = new(PayChequeEvent)
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
func (it *PayChequeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayChequeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayChequeEvent represents a ChequeEvent event raised by the Pay contract.
type PayChequeEvent struct {
	ChequeID     *big.Int
	From         common.Address
	To           common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChequeEvent is a free log retrieval operation binding the contract event 0x4617b59973e052c439b84edb2deaedd5c5dad563af6952409e72c65194ba4a4c.
//
// Solidity: event ChequeEvent(uint256 indexed chequeID, address indexed from, address indexed to, address tokenAddress, uint256 amount, uint8 status)
func (_Pay *PayFilterer) FilterChequeEvent(opts *bind.FilterOpts, chequeID []*big.Int, from []common.Address, to []common.Address) (*PayChequeEventIterator, error) {

	var chequeIDRule []interface{}
	for _, chequeIDItem := range chequeID {
		chequeIDRule = append(chequeIDRule, chequeIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pay.contract.FilterLogs(opts, "ChequeEvent", chequeIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PayChequeEventIterator{contract: _Pay.contract, event: "ChequeEvent", logs: logs, sub: sub}, nil
}

// WatchChequeEvent is a free log subscription operation binding the contract event 0x4617b59973e052c439b84edb2deaedd5c5dad563af6952409e72c65194ba4a4c.
//
// Solidity: event ChequeEvent(uint256 indexed chequeID, address indexed from, address indexed to, address tokenAddress, uint256 amount, uint8 status)
func (_Pay *PayFilterer) WatchChequeEvent(opts *bind.WatchOpts, sink chan<- *PayChequeEvent, chequeID []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var chequeIDRule []interface{}
	for _, chequeIDItem := range chequeID {
		chequeIDRule = append(chequeIDRule, chequeIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pay.contract.WatchLogs(opts, "ChequeEvent", chequeIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayChequeEvent)
				if err := _Pay.contract.UnpackLog(event, "ChequeEvent", log); err != nil {
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

// ParseChequeEvent is a log parse operation binding the contract event 0x4617b59973e052c439b84edb2deaedd5c5dad563af6952409e72c65194ba4a4c.
//
// Solidity: event ChequeEvent(uint256 indexed chequeID, address indexed from, address indexed to, address tokenAddress, uint256 amount, uint8 status)
func (_Pay *PayFilterer) ParseChequeEvent(log types.Log) (*PayChequeEvent, error) {
	event := new(PayChequeEvent)
	if err := _Pay.contract.UnpackLog(event, "ChequeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayChequeSentIterator is returned from FilterChequeSent and is used to iterate over the raw logs and unpacked data for ChequeSent events raised by the Pay contract.
type PayChequeSentIterator struct {
	Event *PayChequeSent // Event containing the contract specifics and raw log

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
func (it *PayChequeSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayChequeSent)
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
		it.Event = new(PayChequeSent)
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
func (it *PayChequeSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayChequeSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayChequeSent represents a ChequeSent event raised by the Pay contract.
type PayChequeSent struct {
	ChequeID     *big.Int
	From         common.Address
	To           common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Expiration   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChequeSent is a free log retrieval operation binding the contract event 0xd18cb5c9f5533f163598578532d03f4b1fb74decf1fb9a23f7d16defc21dda05.
//
// Solidity: event ChequeSent(uint256 indexed chequeID, address indexed from, address indexed to, address tokenAddress, uint256 amount, uint256 expiration)
func (_Pay *PayFilterer) FilterChequeSent(opts *bind.FilterOpts, chequeID []*big.Int, from []common.Address, to []common.Address) (*PayChequeSentIterator, error) {

	var chequeIDRule []interface{}
	for _, chequeIDItem := range chequeID {
		chequeIDRule = append(chequeIDRule, chequeIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pay.contract.FilterLogs(opts, "ChequeSent", chequeIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PayChequeSentIterator{contract: _Pay.contract, event: "ChequeSent", logs: logs, sub: sub}, nil
}

// WatchChequeSent is a free log subscription operation binding the contract event 0xd18cb5c9f5533f163598578532d03f4b1fb74decf1fb9a23f7d16defc21dda05.
//
// Solidity: event ChequeSent(uint256 indexed chequeID, address indexed from, address indexed to, address tokenAddress, uint256 amount, uint256 expiration)
func (_Pay *PayFilterer) WatchChequeSent(opts *bind.WatchOpts, sink chan<- *PayChequeSent, chequeID []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var chequeIDRule []interface{}
	for _, chequeIDItem := range chequeID {
		chequeIDRule = append(chequeIDRule, chequeIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pay.contract.WatchLogs(opts, "ChequeSent", chequeIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayChequeSent)
				if err := _Pay.contract.UnpackLog(event, "ChequeSent", log); err != nil {
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

// ParseChequeSent is a log parse operation binding the contract event 0xd18cb5c9f5533f163598578532d03f4b1fb74decf1fb9a23f7d16defc21dda05.
//
// Solidity: event ChequeSent(uint256 indexed chequeID, address indexed from, address indexed to, address tokenAddress, uint256 amount, uint256 expiration)
func (_Pay *PayFilterer) ParseChequeSent(log types.Log) (*PayChequeSent, error) {
	event := new(PayChequeSent)
	if err := _Pay.contract.UnpackLog(event, "ChequeSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayHaltStatusIterator is returned from FilterHaltStatus and is used to iterate over the raw logs and unpacked data for HaltStatus events raised by the Pay contract.
type PayHaltStatusIterator struct {
	Event *PayHaltStatus // Event containing the contract specifics and raw log

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
func (it *PayHaltStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayHaltStatus)
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
		it.Event = new(PayHaltStatus)
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
func (it *PayHaltStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayHaltStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayHaltStatus represents a HaltStatus event raised by the Pay contract.
type PayHaltStatus struct {
	Status *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHaltStatus is a free log retrieval operation binding the contract event 0xd57639a4f3c4c1059ca3fdd5175e8ca8ebf6f1eee17c2962389afa8bfd9b34e6.
//
// Solidity: event HaltStatus(uint256 status)
func (_Pay *PayFilterer) FilterHaltStatus(opts *bind.FilterOpts) (*PayHaltStatusIterator, error) {

	logs, sub, err := _Pay.contract.FilterLogs(opts, "HaltStatus")
	if err != nil {
		return nil, err
	}
	return &PayHaltStatusIterator{contract: _Pay.contract, event: "HaltStatus", logs: logs, sub: sub}, nil
}

// WatchHaltStatus is a free log subscription operation binding the contract event 0xd57639a4f3c4c1059ca3fdd5175e8ca8ebf6f1eee17c2962389afa8bfd9b34e6.
//
// Solidity: event HaltStatus(uint256 status)
func (_Pay *PayFilterer) WatchHaltStatus(opts *bind.WatchOpts, sink chan<- *PayHaltStatus) (event.Subscription, error) {

	logs, sub, err := _Pay.contract.WatchLogs(opts, "HaltStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayHaltStatus)
				if err := _Pay.contract.UnpackLog(event, "HaltStatus", log); err != nil {
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

// ParseHaltStatus is a log parse operation binding the contract event 0xd57639a4f3c4c1059ca3fdd5175e8ca8ebf6f1eee17c2962389afa8bfd9b34e6.
//
// Solidity: event HaltStatus(uint256 status)
func (_Pay *PayFilterer) ParseHaltStatus(log types.Log) (*PayHaltStatus, error) {
	event := new(PayHaltStatus)
	if err := _Pay.contract.UnpackLog(event, "HaltStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pay contract.
type PayOwnershipTransferredIterator struct {
	Event *PayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayOwnershipTransferred)
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
		it.Event = new(PayOwnershipTransferred)
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
func (it *PayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayOwnershipTransferred represents a OwnershipTransferred event raised by the Pay contract.
type PayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pay *PayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pay.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PayOwnershipTransferredIterator{contract: _Pay.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pay *PayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pay.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayOwnershipTransferred)
				if err := _Pay.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pay *PayFilterer) ParseOwnershipTransferred(log types.Log) (*PayOwnershipTransferred, error) {
	event := new(PayOwnershipTransferred)
	if err := _Pay.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pay contract.
type PayPausedIterator struct {
	Event *PayPaused // Event containing the contract specifics and raw log

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
func (it *PayPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayPaused)
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
		it.Event = new(PayPaused)
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
func (it *PayPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayPaused represents a Paused event raised by the Pay contract.
type PayPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pay *PayFilterer) FilterPaused(opts *bind.FilterOpts) (*PayPausedIterator, error) {

	logs, sub, err := _Pay.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PayPausedIterator{contract: _Pay.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pay *PayFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PayPaused) (event.Subscription, error) {

	logs, sub, err := _Pay.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayPaused)
				if err := _Pay.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pay *PayFilterer) ParsePaused(log types.Log) (*PayPaused, error) {
	event := new(PayPaused)
	if err := _Pay.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pay contract.
type PayUnpausedIterator struct {
	Event *PayUnpaused // Event containing the contract specifics and raw log

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
func (it *PayUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayUnpaused)
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
		it.Event = new(PayUnpaused)
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
func (it *PayUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayUnpaused represents a Unpaused event raised by the Pay contract.
type PayUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pay *PayFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PayUnpausedIterator, error) {

	logs, sub, err := _Pay.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PayUnpausedIterator{contract: _Pay.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pay *PayFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PayUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pay.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayUnpaused)
				if err := _Pay.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pay *PayFilterer) ParseUnpaused(log types.Log) (*PayUnpaused, error) {
	event := new(PayUnpaused)
	if err := _Pay.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayWhitelistTokenAddedIterator is returned from FilterWhitelistTokenAdded and is used to iterate over the raw logs and unpacked data for WhitelistTokenAdded events raised by the Pay contract.
type PayWhitelistTokenAddedIterator struct {
	Event *PayWhitelistTokenAdded // Event containing the contract specifics and raw log

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
func (it *PayWhitelistTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayWhitelistTokenAdded)
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
		it.Event = new(PayWhitelistTokenAdded)
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
func (it *PayWhitelistTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayWhitelistTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayWhitelistTokenAdded represents a WhitelistTokenAdded event raised by the Pay contract.
type PayWhitelistTokenAdded struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWhitelistTokenAdded is a free log retrieval operation binding the contract event 0x9f9d9f44bc581f34670150874e4db46bb05da1753d75c6b5c7b952fcdaa96f4a.
//
// Solidity: event WhitelistTokenAdded(address[] tokens)
func (_Pay *PayFilterer) FilterWhitelistTokenAdded(opts *bind.FilterOpts) (*PayWhitelistTokenAddedIterator, error) {

	logs, sub, err := _Pay.contract.FilterLogs(opts, "WhitelistTokenAdded")
	if err != nil {
		return nil, err
	}
	return &PayWhitelistTokenAddedIterator{contract: _Pay.contract, event: "WhitelistTokenAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistTokenAdded is a free log subscription operation binding the contract event 0x9f9d9f44bc581f34670150874e4db46bb05da1753d75c6b5c7b952fcdaa96f4a.
//
// Solidity: event WhitelistTokenAdded(address[] tokens)
func (_Pay *PayFilterer) WatchWhitelistTokenAdded(opts *bind.WatchOpts, sink chan<- *PayWhitelistTokenAdded) (event.Subscription, error) {

	logs, sub, err := _Pay.contract.WatchLogs(opts, "WhitelistTokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayWhitelistTokenAdded)
				if err := _Pay.contract.UnpackLog(event, "WhitelistTokenAdded", log); err != nil {
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

// ParseWhitelistTokenAdded is a log parse operation binding the contract event 0x9f9d9f44bc581f34670150874e4db46bb05da1753d75c6b5c7b952fcdaa96f4a.
//
// Solidity: event WhitelistTokenAdded(address[] tokens)
func (_Pay *PayFilterer) ParseWhitelistTokenAdded(log types.Log) (*PayWhitelistTokenAdded, error) {
	event := new(PayWhitelistTokenAdded)
	if err := _Pay.contract.UnpackLog(event, "WhitelistTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayWhitelistTokenRemovedIterator is returned from FilterWhitelistTokenRemoved and is used to iterate over the raw logs and unpacked data for WhitelistTokenRemoved events raised by the Pay contract.
type PayWhitelistTokenRemovedIterator struct {
	Event *PayWhitelistTokenRemoved // Event containing the contract specifics and raw log

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
func (it *PayWhitelistTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayWhitelistTokenRemoved)
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
		it.Event = new(PayWhitelistTokenRemoved)
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
func (it *PayWhitelistTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayWhitelistTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayWhitelistTokenRemoved represents a WhitelistTokenRemoved event raised by the Pay contract.
type PayWhitelistTokenRemoved struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWhitelistTokenRemoved is a free log retrieval operation binding the contract event 0xa50261031b405a768d9301886abe947b7a26918759bc9b5c89a4398b7dcff16e.
//
// Solidity: event WhitelistTokenRemoved(address[] tokens)
func (_Pay *PayFilterer) FilterWhitelistTokenRemoved(opts *bind.FilterOpts) (*PayWhitelistTokenRemovedIterator, error) {

	logs, sub, err := _Pay.contract.FilterLogs(opts, "WhitelistTokenRemoved")
	if err != nil {
		return nil, err
	}
	return &PayWhitelistTokenRemovedIterator{contract: _Pay.contract, event: "WhitelistTokenRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistTokenRemoved is a free log subscription operation binding the contract event 0xa50261031b405a768d9301886abe947b7a26918759bc9b5c89a4398b7dcff16e.
//
// Solidity: event WhitelistTokenRemoved(address[] tokens)
func (_Pay *PayFilterer) WatchWhitelistTokenRemoved(opts *bind.WatchOpts, sink chan<- *PayWhitelistTokenRemoved) (event.Subscription, error) {

	logs, sub, err := _Pay.contract.WatchLogs(opts, "WhitelistTokenRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayWhitelistTokenRemoved)
				if err := _Pay.contract.UnpackLog(event, "WhitelistTokenRemoved", log); err != nil {
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

// ParseWhitelistTokenRemoved is a log parse operation binding the contract event 0xa50261031b405a768d9301886abe947b7a26918759bc9b5c89a4398b7dcff16e.
//
// Solidity: event WhitelistTokenRemoved(address[] tokens)
func (_Pay *PayFilterer) ParseWhitelistTokenRemoved(log types.Log) (*PayWhitelistTokenRemoved, error) {
	event := new(PayWhitelistTokenRemoved)
	if err := _Pay.contract.UnpackLog(event, "WhitelistTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
