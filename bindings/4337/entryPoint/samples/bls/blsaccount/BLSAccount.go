// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blsaccount

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

// BLSAccountMetaData contains all meta data concerning the BLSAccount contract.
var BLSAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"anAggregator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"oldPublicKey\",\"type\":\"uint256[4]\"},{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"newPublicKey\",\"type\":\"uint256[4]\"}],\"name\":\"PublicKeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SimpleAccountInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlsPublicKey\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"aPublicKey\",\"type\":\"uint256[4]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"newPublicKey\",\"type\":\"uint256[4]\"}],\"name\":\"setBlsPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e03461016b57611731906001600160401b03601f38849003908101601f191683019082821184831017610170578084916040968794855283398101031261016b5781516001600160a01b0392838216820361016b5760200151928316830361016b573060805260a0527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff82861c1661015a578080831603610116575b50505060c052516115aa9081610187823960805181818161089d01526109d6015260a0518181816103d7015281816106f20152818161079401528181610aea01528181610c2901528181610ded0152818161110e015281816112e901526113cd015260c051818181610432015261055f0152f35b6001600160401b0319909116811790915582519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a13880806100a2565b845163f92ee8a960e01b8152600490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c806301ffc9a71461016b578063150b7a021461016657806318fc5c441461016157806319822f7c1461015c578063245a7bfc1461015757806347e1da2a146101525780634a58db191461014d5780634d44560d146101485780634f1ef2861461014357806352d1902d1461013e5780638da5cb5b14610139578063ad3cb1cc14610134578063b0d691fe1461012f578063b61d27f61461012a578063bc197c8114610125578063c399ec8814610120578063c4d66de81461011b578063d087d28814610116578063e02afbae14610111578063ee472f361461010c5763f23a6e610361000e57610f67565b610eab565b610e81565b610dba565b610c8a565b610bfd565b610b6e565b610b19565b610ad4565b610a57565b610a2e565b6109c3565b61084a565b610763565b6106e3565b6105be565b610549565b6103a1565b610383565b61021c565b346101d95760203660031901126101d95760043563ffffffff60e01b81168091036101d957602090630a85bd0160e11b81149081156101c8575b81156101b7575b506040519015158152f35b6301ffc9a760e01b149050386101ac565b630271189760e51b811491506101a5565b600080fd5b6001600160a01b038116036101d957565b9181601f840112156101d9578235916001600160401b0383116101d957602083818601950101116101d957565b346101d95760803660031901126101d9576102386004356101de565b6102436024356101de565b6064356001600160401b0381116101d9576102629036906004016101ef565b5050604051630a85bd0160e11b8152602090f35b634e487b7160e01b600052604160045260246000fd5b6001600160401b03811161029f57604052565b610276565b608081019081106001600160401b0382111761029f57604052565b60a081019081106001600160401b0382111761029f57604052565b604081019081106001600160401b0382111761029f57604052565b90601f801991011681019081106001600160401b0382111761029f57604052565b60405190606082018281106001600160401b0382111761029f57604052565b60806003198201126101d95780602312156101d95760405190610357826102a4565b816084916084116101d9576004905b8282106103735750505090565b8135815260209182019101610366565b346101d95761001961039436610335565b61039c611161565b6111b8565b346101d9576003196060368201126101d9576004356001600160401b0381116101d957610120816004019282360301126101d9577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036105045760440190610413828261103b565b905061048e575b61048a61046f610428610316565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001681526000602082018190526040820152516001600160a01b031690565b61047a6044356112c1565b6040519081529081906020820190565b0390f35b6104f06104e96104fd936104a06110d1565b6040516104b1602082018093610e59565b608081526104be816102bf565b519020936104dc6104e36104d2848461103b565b948593919461103b565b9050611219565b9161123e565b3691610813565b6020815191012014611256565b388061041a565b60405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e74000000006044820152606490fd5b346101d95760003660031901126101d9576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b9181601f840112156101d9578235916001600160401b0383116101d9576020808501948460051b0101116101d957565b346101d95760603660031901126101d9576001600160401b036004358181116101d9576105ef90369060040161058e565b6024929192358281116101d95761060a90369060040161058e565b926044359081116101d95761062390369060040161058e565b93909161062e6112df565b848414806106d2575b61064090610fc1565b8161068657505060005b82811061065357005b8061068061066c610667600194878a611019565b61102e565b61067a6104e984898861106d565b9061136d565b0161064a565b91909460009493945b85811061069857005b806106cc6106ac6106676001948a87611019565b6106b7838b89611019565b356106c66104e9858b8a61106d565b91611395565b0161068f565b508115806106375750818514610637565b600080600319360112610760577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681813b156107605760405163b760faf960e01b8152306004820152918290602490829034905af1801561075b5761074f575080f35b6107589061028c565b80f35b611088565b80fd5b346101d9576000604036600319011261076057600435610782816101de565b61078a611161565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b156107f45760449083604051958694859363040b850f60e31b855216600484015260243560248401525af1801561075b5761074f575080f35b8280fd5b6001600160401b03811161029f57601f01601f191660200190565b92919261081f826107f8565b9161082d60405193846102f5565b8294818452818301116101d9578281602093846000960137010152565b60403660031901126101d9576004803590610864826101de565b6024356001600160401b0381116101d957366023820112156101d9576108939036906024818501359101610813565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081163081149081156109a7575b506109965790602083926108db611161565b6040516352d1902d60e01b8152938491829088165afa60009281610965575b50610928575050604051634c9c8ce360e01b81526001600160a01b0390921690820190815281906020010390fd5b83836000805160206115358339815191528403610949576100198383611414565b604051632a87526960e21b815290810184815281906020010390fd5b61098891935060203d60201161098f575b61098081836102f5565b810190611094565b91386108fa565b503d610976565b60405163703e46dd60e11b81528390fd5b90508160008051602061153583398151915254161415386108c9565b346101d95760003660031901126101d9577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610a1c5760206040516000805160206115358339815191528152f35b60405163703e46dd60e11b8152600490fd5b346101d95760003660031901126101d9576000546040516001600160a01b039091168152602090f35b346101d95760003660031901126101d9576040805190610a76826102da565b60058252602090640352e302e360dc1b6020840152604051916020835283519182602085015260005b838110610ac15784604081866000838284010152601f80199101168101030190f35b8581018301518582018301528201610a9f565b346101d95760003660031901126101d9576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101d95760603660031901126101d957600435610b36816101de565b604435906001600160401b0382116101d957610b64610b5c6100199336906004016101ef565b6104e96112df565b9060243590611395565b346101d95760a03660031901126101d957610b8a6004356101de565b610b956024356101de565b6001600160401b036044358181116101d957610bb590369060040161058e565b50506064358181116101d957610bcf90369060040161058e565b50506084359081116101d957610be99036906004016101ef565b505060405163bc197c8160e01b8152602090f35b346101d95760003660031901126101d9576040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561075b57602091600091610c6d575b50604051908152f35b610c849150823d841161098f5761098081836102f5565b38610c64565b346101d95760203660031901126101d957600435610ca7816101de565b60008051602061155583398151915254906001600160401b0360ff8360401c1615921680159081610db2575b6001149081610da8575b159081610d9f575b50610d8d57600080516020611555833981519152805467ffffffffffffffff19166001179055610d199082610d68576113ac565b610d1f57005b600080516020611555833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b600080516020611555833981519152805460ff60401b1916600160401b1790556113ac565b60405163f92ee8a960e01b8152600490fd5b90501538610ce5565b303b159150610cdd565b839150610cd3565b346101d95760003660031901126101d957604051631aab3f0d60e11b8152306004820152600060248201526020816044817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561075b5761048a91600091610e3a575b506040519081529081906020820190565b610e53915060203d60201161098f5761098081836102f5565b38610e29565b6000915b60048310610e6a57505050565b600190825181526020809101920192019190610e5d565b346101d95760003660031901126101d9576080610e9c6110d1565b610ea96040518092610e59565bf35b346101d957610eb936610335565b60008051602061155583398151915254906001600160401b0360ff8360401c1615921680159081610f5f575b6001149081610f55575b159081610f4c575b50610d8d57600080516020611555833981519152805467ffffffffffffffff19166001179055610d199082156110f757600080516020611555833981519152805460ff60401b1916600160401b1790556110f7565b90501538610ef7565b303b159150610eef565b839150610ee5565b346101d95760a03660031901126101d957610f836004356101de565b610f8e6024356101de565b6084356001600160401b0381116101d957610fad9036906004016101ef565b505060405163f23a6e6160e01b8152602090f35b15610fc857565b60405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b6044820152606490fd5b634e487b7160e01b600052603260045260246000fd5b91908110156110295760051b0190565b611003565b35611038816101de565b90565b903590601e19813603018212156101d957018035906001600160401b0382116101d9576020019181360383136101d957565b90821015611029576110849160051b81019061103b565b9091565b6040513d6000823e3d90fd5b908160209103126101d9575190565b906001916001906000905b600482106110bd575050509050565b8254815291840191908401906020016110ae565b60806040516110df816102a4565b3690376040516110ee816110a3565b611038816102a4565b600080546001600160a01b031916815561115f91907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de8280a36111b8565b565b6000546001600160a01b0316331480156111af575b1561117d57565b60405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b6044820152606490fd5b50303314611176565b7f42e4c4ce1432650f17e41c4ea77ed12c0ab20b229d3ffd84a2ebc9f8abb25a836101006040516111e8816110a3565b6111f56080820185610e59565ba160005b60048110611205575050565b6001906020835193019281830155016111f9565b607f1981019190821161122857565b634e487b7160e01b600052601160045260246000fd5b909392938483116101d95784116101d9578101920390565b1561125d57565b60405162461bcd60e51b815260206004820152600c60248201526b77726f6e67207075626b657960a01b6044820152606490fd5b3d156112bc573d906112a2826107f8565b916112b060405193846102f5565b82523d6000602084013e565b606090565b806112c95750565b600080808093338219f1506112dc611291565b50565b60018060a01b03807f000000000000000000000000000000000000000000000000000000000000000016331490811561135f575b501561131b57565b606460405162461bcd60e51b815260206004820152602060248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152fd5b905060005416331438611313565b600091829182602083519301915af1611384611291565b901561138d5750565b602081519101fd5b916000928392602083519301915af1611384611291565b600080546001600160a01b0319166001600160a01b039283169081178255917f000000000000000000000000000000000000000000000000000000000000000016907f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de9080a3565b90813b156114975760008051602061153583398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a280511561147c576112dc916114b8565b50503461148557565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b60008061103893602081519101845af46114d0611291565b91906114f857508051156114e657805190602001fd5b604051630a12f52160e11b8152600490fd5b8151158061152b575b611509575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561150156fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212205712c32d2e13f6980e2a2049096e8afa6aee105da0771d0e2ee2c89646c87c2a64736f6c63430008190033",
}

// BLSAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use BLSAccountMetaData.ABI instead.
var BLSAccountABI = BLSAccountMetaData.ABI

// BLSAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BLSAccountMetaData.Bin instead.
var BLSAccountBin = BLSAccountMetaData.Bin

// DeployBLSAccount deploys a new Ethereum contract, binding an instance of BLSAccount to it.
func DeployBLSAccount(auth *bind.TransactOpts, backend bind.ContractBackend, anEntryPoint common.Address, anAggregator common.Address) (common.Address, *types.Transaction, *BLSAccount, error) {
	parsed, err := BLSAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BLSAccountBin), backend, anEntryPoint, anAggregator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BLSAccount{BLSAccountCaller: BLSAccountCaller{contract: contract}, BLSAccountTransactor: BLSAccountTransactor{contract: contract}, BLSAccountFilterer: BLSAccountFilterer{contract: contract}}, nil
}

// BLSAccount is an auto generated Go binding around an Ethereum contract.
type BLSAccount struct {
	BLSAccountCaller     // Read-only binding to the contract
	BLSAccountTransactor // Write-only binding to the contract
	BLSAccountFilterer   // Log filterer for contract events
}

// BLSAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type BLSAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BLSAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BLSAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BLSAccountSession struct {
	Contract     *BLSAccount       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BLSAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BLSAccountCallerSession struct {
	Contract *BLSAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BLSAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BLSAccountTransactorSession struct {
	Contract     *BLSAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BLSAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type BLSAccountRaw struct {
	Contract *BLSAccount // Generic contract binding to access the raw methods on
}

// BLSAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BLSAccountCallerRaw struct {
	Contract *BLSAccountCaller // Generic read-only contract binding to access the raw methods on
}

// BLSAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BLSAccountTransactorRaw struct {
	Contract *BLSAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBLSAccount creates a new instance of BLSAccount, bound to a specific deployed contract.
func NewBLSAccount(address common.Address, backend bind.ContractBackend) (*BLSAccount, error) {
	contract, err := bindBLSAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BLSAccount{BLSAccountCaller: BLSAccountCaller{contract: contract}, BLSAccountTransactor: BLSAccountTransactor{contract: contract}, BLSAccountFilterer: BLSAccountFilterer{contract: contract}}, nil
}

// NewBLSAccountCaller creates a new read-only instance of BLSAccount, bound to a specific deployed contract.
func NewBLSAccountCaller(address common.Address, caller bind.ContractCaller) (*BLSAccountCaller, error) {
	contract, err := bindBLSAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BLSAccountCaller{contract: contract}, nil
}

// NewBLSAccountTransactor creates a new write-only instance of BLSAccount, bound to a specific deployed contract.
func NewBLSAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*BLSAccountTransactor, error) {
	contract, err := bindBLSAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BLSAccountTransactor{contract: contract}, nil
}

// NewBLSAccountFilterer creates a new log filterer instance of BLSAccount, bound to a specific deployed contract.
func NewBLSAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*BLSAccountFilterer, error) {
	contract, err := bindBLSAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BLSAccountFilterer{contract: contract}, nil
}

// bindBLSAccount binds a generic wrapper to an already deployed contract.
func bindBLSAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BLSAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSAccount *BLSAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSAccount.Contract.BLSAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSAccount *BLSAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSAccount.Contract.BLSAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSAccount *BLSAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSAccount.Contract.BLSAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSAccount *BLSAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSAccount *BLSAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSAccount *BLSAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSAccount.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BLSAccount *BLSAccountCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BLSAccount.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BLSAccount *BLSAccountSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BLSAccount.Contract.UPGRADEINTERFACEVERSION(&_BLSAccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BLSAccount *BLSAccountCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BLSAccount.Contract.UPGRADEINTERFACEVERSION(&_BLSAccount.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_BLSAccount *BLSAccountCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSAccount.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_BLSAccount *BLSAccountSession) Aggregator() (common.Address, error) {
	return _BLSAccount.Contract.Aggregator(&_BLSAccount.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_BLSAccount *BLSAccountCallerSession) Aggregator() (common.Address, error) {
	return _BLSAccount.Contract.Aggregator(&_BLSAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_BLSAccount *BLSAccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSAccount.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_BLSAccount *BLSAccountSession) EntryPoint() (common.Address, error) {
	return _BLSAccount.Contract.EntryPoint(&_BLSAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_BLSAccount *BLSAccountCallerSession) EntryPoint() (common.Address, error) {
	return _BLSAccount.Contract.EntryPoint(&_BLSAccount.CallOpts)
}

// GetBlsPublicKey is a free data retrieval call binding the contract method 0xe02afbae.
//
// Solidity: function getBlsPublicKey() view returns(uint256[4])
func (_BLSAccount *BLSAccountCaller) GetBlsPublicKey(opts *bind.CallOpts) ([4]*big.Int, error) {
	var out []interface{}
	err := _BLSAccount.contract.Call(opts, &out, "getBlsPublicKey")

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetBlsPublicKey is a free data retrieval call binding the contract method 0xe02afbae.
//
// Solidity: function getBlsPublicKey() view returns(uint256[4])
func (_BLSAccount *BLSAccountSession) GetBlsPublicKey() ([4]*big.Int, error) {
	return _BLSAccount.Contract.GetBlsPublicKey(&_BLSAccount.CallOpts)
}

// GetBlsPublicKey is a free data retrieval call binding the contract method 0xe02afbae.
//
// Solidity: function getBlsPublicKey() view returns(uint256[4])
func (_BLSAccount *BLSAccountCallerSession) GetBlsPublicKey() ([4]*big.Int, error) {
	return _BLSAccount.Contract.GetBlsPublicKey(&_BLSAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_BLSAccount *BLSAccountCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BLSAccount.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_BLSAccount *BLSAccountSession) GetDeposit() (*big.Int, error) {
	return _BLSAccount.Contract.GetDeposit(&_BLSAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_BLSAccount *BLSAccountCallerSession) GetDeposit() (*big.Int, error) {
	return _BLSAccount.Contract.GetDeposit(&_BLSAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_BLSAccount *BLSAccountCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BLSAccount.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_BLSAccount *BLSAccountSession) GetNonce() (*big.Int, error) {
	return _BLSAccount.Contract.GetNonce(&_BLSAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_BLSAccount *BLSAccountCallerSession) GetNonce() (*big.Int, error) {
	return _BLSAccount.Contract.GetNonce(&_BLSAccount.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_BLSAccount *BLSAccountCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _BLSAccount.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_BLSAccount *BLSAccountSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _BLSAccount.Contract.OnERC1155BatchReceived(&_BLSAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_BLSAccount *BLSAccountCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _BLSAccount.Contract.OnERC1155BatchReceived(&_BLSAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_BLSAccount *BLSAccountCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _BLSAccount.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_BLSAccount *BLSAccountSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _BLSAccount.Contract.OnERC1155Received(&_BLSAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_BLSAccount *BLSAccountCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _BLSAccount.Contract.OnERC1155Received(&_BLSAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_BLSAccount *BLSAccountCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _BLSAccount.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_BLSAccount *BLSAccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _BLSAccount.Contract.OnERC721Received(&_BLSAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_BLSAccount *BLSAccountCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _BLSAccount.Contract.OnERC721Received(&_BLSAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BLSAccount *BLSAccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSAccount.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BLSAccount *BLSAccountSession) Owner() (common.Address, error) {
	return _BLSAccount.Contract.Owner(&_BLSAccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BLSAccount *BLSAccountCallerSession) Owner() (common.Address, error) {
	return _BLSAccount.Contract.Owner(&_BLSAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BLSAccount *BLSAccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BLSAccount.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BLSAccount *BLSAccountSession) ProxiableUUID() ([32]byte, error) {
	return _BLSAccount.Contract.ProxiableUUID(&_BLSAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BLSAccount *BLSAccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BLSAccount.Contract.ProxiableUUID(&_BLSAccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BLSAccount *BLSAccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BLSAccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BLSAccount *BLSAccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BLSAccount.Contract.SupportsInterface(&_BLSAccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BLSAccount *BLSAccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BLSAccount.Contract.SupportsInterface(&_BLSAccount.CallOpts, interfaceId)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_BLSAccount *BLSAccountTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSAccount.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_BLSAccount *BLSAccountSession) AddDeposit() (*types.Transaction, error) {
	return _BLSAccount.Contract.AddDeposit(&_BLSAccount.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_BLSAccount *BLSAccountTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _BLSAccount.Contract.AddDeposit(&_BLSAccount.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_BLSAccount *BLSAccountTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _BLSAccount.contract.Transact(opts, "execute", dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_BLSAccount *BLSAccountSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _BLSAccount.Contract.Execute(&_BLSAccount.TransactOpts, dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_BLSAccount *BLSAccountTransactorSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _BLSAccount.Contract.Execute(&_BLSAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_BLSAccount *BLSAccountTransactor) ExecuteBatch(opts *bind.TransactOpts, dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _BLSAccount.contract.Transact(opts, "executeBatch", dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_BLSAccount *BLSAccountSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _BLSAccount.Contract.ExecuteBatch(&_BLSAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_BLSAccount *BLSAccountTransactorSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _BLSAccount.Contract.ExecuteBatch(&_BLSAccount.TransactOpts, dest, value, arg2)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_BLSAccount *BLSAccountTransactor) Initialize(opts *bind.TransactOpts, anOwner common.Address) (*types.Transaction, error) {
	return _BLSAccount.contract.Transact(opts, "initialize", anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_BLSAccount *BLSAccountSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _BLSAccount.Contract.Initialize(&_BLSAccount.TransactOpts, anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_BLSAccount *BLSAccountTransactorSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _BLSAccount.Contract.Initialize(&_BLSAccount.TransactOpts, anOwner)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xee472f36.
//
// Solidity: function initialize(uint256[4] aPublicKey) returns()
func (_BLSAccount *BLSAccountTransactor) Initialize0(opts *bind.TransactOpts, aPublicKey [4]*big.Int) (*types.Transaction, error) {
	return _BLSAccount.contract.Transact(opts, "initialize0", aPublicKey)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xee472f36.
//
// Solidity: function initialize(uint256[4] aPublicKey) returns()
func (_BLSAccount *BLSAccountSession) Initialize0(aPublicKey [4]*big.Int) (*types.Transaction, error) {
	return _BLSAccount.Contract.Initialize0(&_BLSAccount.TransactOpts, aPublicKey)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xee472f36.
//
// Solidity: function initialize(uint256[4] aPublicKey) returns()
func (_BLSAccount *BLSAccountTransactorSession) Initialize0(aPublicKey [4]*big.Int) (*types.Transaction, error) {
	return _BLSAccount.Contract.Initialize0(&_BLSAccount.TransactOpts, aPublicKey)
}

// SetBlsPublicKey is a paid mutator transaction binding the contract method 0x18fc5c44.
//
// Solidity: function setBlsPublicKey(uint256[4] newPublicKey) returns()
func (_BLSAccount *BLSAccountTransactor) SetBlsPublicKey(opts *bind.TransactOpts, newPublicKey [4]*big.Int) (*types.Transaction, error) {
	return _BLSAccount.contract.Transact(opts, "setBlsPublicKey", newPublicKey)
}

// SetBlsPublicKey is a paid mutator transaction binding the contract method 0x18fc5c44.
//
// Solidity: function setBlsPublicKey(uint256[4] newPublicKey) returns()
func (_BLSAccount *BLSAccountSession) SetBlsPublicKey(newPublicKey [4]*big.Int) (*types.Transaction, error) {
	return _BLSAccount.Contract.SetBlsPublicKey(&_BLSAccount.TransactOpts, newPublicKey)
}

// SetBlsPublicKey is a paid mutator transaction binding the contract method 0x18fc5c44.
//
// Solidity: function setBlsPublicKey(uint256[4] newPublicKey) returns()
func (_BLSAccount *BLSAccountTransactorSession) SetBlsPublicKey(newPublicKey [4]*big.Int) (*types.Transaction, error) {
	return _BLSAccount.Contract.SetBlsPublicKey(&_BLSAccount.TransactOpts, newPublicKey)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BLSAccount *BLSAccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BLSAccount.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BLSAccount *BLSAccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BLSAccount.Contract.UpgradeToAndCall(&_BLSAccount.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BLSAccount *BLSAccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BLSAccount.Contract.UpgradeToAndCall(&_BLSAccount.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_BLSAccount *BLSAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _BLSAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_BLSAccount *BLSAccountSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _BLSAccount.Contract.ValidateUserOp(&_BLSAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_BLSAccount *BLSAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _BLSAccount.Contract.ValidateUserOp(&_BLSAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_BLSAccount *BLSAccountTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BLSAccount.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_BLSAccount *BLSAccountSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BLSAccount.Contract.WithdrawDepositTo(&_BLSAccount.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_BLSAccount *BLSAccountTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BLSAccount.Contract.WithdrawDepositTo(&_BLSAccount.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BLSAccount *BLSAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BLSAccount *BLSAccountSession) Receive() (*types.Transaction, error) {
	return _BLSAccount.Contract.Receive(&_BLSAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BLSAccount *BLSAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _BLSAccount.Contract.Receive(&_BLSAccount.TransactOpts)
}

// BLSAccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BLSAccount contract.
type BLSAccountInitializedIterator struct {
	Event *BLSAccountInitialized // Event containing the contract specifics and raw log

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
func (it *BLSAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSAccountInitialized)
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
		it.Event = new(BLSAccountInitialized)
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
func (it *BLSAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSAccountInitialized represents a Initialized event raised by the BLSAccount contract.
type BLSAccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BLSAccount *BLSAccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*BLSAccountInitializedIterator, error) {

	logs, sub, err := _BLSAccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BLSAccountInitializedIterator{contract: _BLSAccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BLSAccount *BLSAccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BLSAccountInitialized) (event.Subscription, error) {

	logs, sub, err := _BLSAccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSAccountInitialized)
				if err := _BLSAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BLSAccount *BLSAccountFilterer) ParseInitialized(log types.Log) (*BLSAccountInitialized, error) {
	event := new(BLSAccountInitialized)
	if err := _BLSAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSAccountPublicKeyChangedIterator is returned from FilterPublicKeyChanged and is used to iterate over the raw logs and unpacked data for PublicKeyChanged events raised by the BLSAccount contract.
type BLSAccountPublicKeyChangedIterator struct {
	Event *BLSAccountPublicKeyChanged // Event containing the contract specifics and raw log

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
func (it *BLSAccountPublicKeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSAccountPublicKeyChanged)
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
		it.Event = new(BLSAccountPublicKeyChanged)
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
func (it *BLSAccountPublicKeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSAccountPublicKeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSAccountPublicKeyChanged represents a PublicKeyChanged event raised by the BLSAccount contract.
type BLSAccountPublicKeyChanged struct {
	OldPublicKey [4]*big.Int
	NewPublicKey [4]*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyChanged is a free log retrieval operation binding the contract event 0x42e4c4ce1432650f17e41c4ea77ed12c0ab20b229d3ffd84a2ebc9f8abb25a83.
//
// Solidity: event PublicKeyChanged(uint256[4] oldPublicKey, uint256[4] newPublicKey)
func (_BLSAccount *BLSAccountFilterer) FilterPublicKeyChanged(opts *bind.FilterOpts) (*BLSAccountPublicKeyChangedIterator, error) {

	logs, sub, err := _BLSAccount.contract.FilterLogs(opts, "PublicKeyChanged")
	if err != nil {
		return nil, err
	}
	return &BLSAccountPublicKeyChangedIterator{contract: _BLSAccount.contract, event: "PublicKeyChanged", logs: logs, sub: sub}, nil
}

// WatchPublicKeyChanged is a free log subscription operation binding the contract event 0x42e4c4ce1432650f17e41c4ea77ed12c0ab20b229d3ffd84a2ebc9f8abb25a83.
//
// Solidity: event PublicKeyChanged(uint256[4] oldPublicKey, uint256[4] newPublicKey)
func (_BLSAccount *BLSAccountFilterer) WatchPublicKeyChanged(opts *bind.WatchOpts, sink chan<- *BLSAccountPublicKeyChanged) (event.Subscription, error) {

	logs, sub, err := _BLSAccount.contract.WatchLogs(opts, "PublicKeyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSAccountPublicKeyChanged)
				if err := _BLSAccount.contract.UnpackLog(event, "PublicKeyChanged", log); err != nil {
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

// ParsePublicKeyChanged is a log parse operation binding the contract event 0x42e4c4ce1432650f17e41c4ea77ed12c0ab20b229d3ffd84a2ebc9f8abb25a83.
//
// Solidity: event PublicKeyChanged(uint256[4] oldPublicKey, uint256[4] newPublicKey)
func (_BLSAccount *BLSAccountFilterer) ParsePublicKeyChanged(log types.Log) (*BLSAccountPublicKeyChanged, error) {
	event := new(BLSAccountPublicKeyChanged)
	if err := _BLSAccount.contract.UnpackLog(event, "PublicKeyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSAccountSimpleAccountInitializedIterator is returned from FilterSimpleAccountInitialized and is used to iterate over the raw logs and unpacked data for SimpleAccountInitialized events raised by the BLSAccount contract.
type BLSAccountSimpleAccountInitializedIterator struct {
	Event *BLSAccountSimpleAccountInitialized // Event containing the contract specifics and raw log

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
func (it *BLSAccountSimpleAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSAccountSimpleAccountInitialized)
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
		it.Event = new(BLSAccountSimpleAccountInitialized)
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
func (it *BLSAccountSimpleAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSAccountSimpleAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSAccountSimpleAccountInitialized represents a SimpleAccountInitialized event raised by the BLSAccount contract.
type BLSAccountSimpleAccountInitialized struct {
	EntryPoint common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSimpleAccountInitialized is a free log retrieval operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_BLSAccount *BLSAccountFilterer) FilterSimpleAccountInitialized(opts *bind.FilterOpts, entryPoint []common.Address, owner []common.Address) (*BLSAccountSimpleAccountInitializedIterator, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BLSAccount.contract.FilterLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BLSAccountSimpleAccountInitializedIterator{contract: _BLSAccount.contract, event: "SimpleAccountInitialized", logs: logs, sub: sub}, nil
}

// WatchSimpleAccountInitialized is a free log subscription operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_BLSAccount *BLSAccountFilterer) WatchSimpleAccountInitialized(opts *bind.WatchOpts, sink chan<- *BLSAccountSimpleAccountInitialized, entryPoint []common.Address, owner []common.Address) (event.Subscription, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BLSAccount.contract.WatchLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSAccountSimpleAccountInitialized)
				if err := _BLSAccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
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

// ParseSimpleAccountInitialized is a log parse operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_BLSAccount *BLSAccountFilterer) ParseSimpleAccountInitialized(log types.Log) (*BLSAccountSimpleAccountInitialized, error) {
	event := new(BLSAccountSimpleAccountInitialized)
	if err := _BLSAccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSAccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BLSAccount contract.
type BLSAccountUpgradedIterator struct {
	Event *BLSAccountUpgraded // Event containing the contract specifics and raw log

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
func (it *BLSAccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSAccountUpgraded)
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
		it.Event = new(BLSAccountUpgraded)
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
func (it *BLSAccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSAccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSAccountUpgraded represents a Upgraded event raised by the BLSAccount contract.
type BLSAccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BLSAccount *BLSAccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BLSAccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BLSAccount.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BLSAccountUpgradedIterator{contract: _BLSAccount.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BLSAccount *BLSAccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BLSAccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BLSAccount.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSAccountUpgraded)
				if err := _BLSAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BLSAccount *BLSAccountFilterer) ParseUpgraded(log types.Log) (*BLSAccountUpgraded, error) {
	event := new(BLSAccountUpgraded)
	if err := _BLSAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
