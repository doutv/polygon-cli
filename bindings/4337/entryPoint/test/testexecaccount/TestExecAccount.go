// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testexecaccount

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

// TestExecAccountMetaData contains all meta data concerning the TestExecAccount contract.
var TestExecAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"innerCallRet\",\"type\":\"bytes\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SimpleAccountInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executeUserOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c034610142576001600160401b0390601f61170f38819003918201601f1916830191848311848410176101475780849260209460405283398101031261014257516001600160a01b0381168103610142573060805260a0527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166101305780808316036100eb575b6040516115b1908161015e82396080518181816106c201526107fb015260a051818181610297015281816104bb0152818161055d01528181610a4d01528181610b8c01528181610d5901528181611124015261125f0152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1388080610092565b60405163f92ee8a960e01b8152600490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c806301ffc9a71461013b578063150b7a021461013657806319822f7c1461013157806347e1da2a1461012c5780634a58db19146101275780634d44560d146101225780634f1ef2861461011d57806352d1902d146101185780638da5cb5b146101135780638dd7712f1461010e578063ad3cb1cc14610109578063b0d691fe14610104578063b61d27f6146100ff578063bc197c81146100fa578063c399ec88146100f5578063c4d66de8146100f0578063d087d288146100eb5763f23a6e610361000e57610dc4565b610d26565b610bed565b610b60565b610ad1565b610a7c565b610a37565b6109df565b61087c565b610853565b6107e8565b610680565b61052c565b6104ac565b610380565b610262565b6101f9565b346101a95760203660031901126101a95760043563ffffffff60e01b81168091036101a957602090630a85bd0160e11b8114908115610198575b8115610187575b506040519015158152f35b6301ffc9a760e01b1490503861017c565b630271189760e51b81149150610175565b600080fd5b6001600160a01b038116036101a957565b35906101ca826101ae565b565b9181601f840112156101a9578235916001600160401b0383116101a957602083818601950101116101a957565b346101a95760803660031901126101a9576102156004356101ae565b6102206024356101ae565b6064356001600160401b0381116101a95761023f9036906004016101cc565b5050604051630a85bd0160e11b8152602090f35b90816101209103126101a95790565b346101a95760603660031901126101a9576004356001600160401b0381116101a957610292903690600401610253565b6044357f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361030b576102d56102ed92602435906110a9565b90806102f1575b506040519081529081906020820190565b0390f35b600080808093338219f150610304610eee565b50386102dc565b60405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e74000000006044820152606490fd5b9181601f840112156101a9578235916001600160401b0383116101a9576020808501948460051b0101116101a957565b346101a95760603660031901126101a9576001600160401b036004358181116101a9576103b1903690600401610350565b6024929192358281116101a9576103cc903690600401610350565b926044359081116101a9576103e5903690600401610350565b9390916103f061111a565b8484148061049b575b61040290610e1e565b8161044f57505060005b82811061041557005b8061044961042e610429600194878a610e76565b610e8b565b61044361043c848988610ec7565b369161062b565b906111a8565b0161040c565b91909460009493945b85811061046157005b806104956104756104296001948a87610e76565b610480838b89610e76565b3561048f61043c858b8a610ec7565b916111d0565b01610458565b508115806103f957508185146103f9565b600080600319360112610529577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681813b156105295760405163b760faf960e01b8152306004820152918290602490829034905af1801561052457610518575080f35b610521906105d7565b80f35b610ee2565b80fd5b346101a957600060403660031901126105295760043561054b816101ae565b6105536111e7565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b156105bd5760449083604051958694859363040b850f60e31b855216600484015260243560248401525af1801561052457610518575080f35b8280fd5b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116105ea57604052565b6105c1565b90601f801991011681019081106001600160401b038211176105ea57604052565b6001600160401b0381116105ea57601f01601f191660200190565b92919261063782610610565b9161064560405193846105ef565b8294818452818301116101a9578281602093846000960137010152565b9080601f830112156101a95781602061067d9335910161062b565b90565b60403660031901126101a957600480359061069a826101ae565b6024356001600160401b0381116101a9576106b89036908301610662565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081163081149081156107cc575b506107bb5790602083926107006111e7565b6040516352d1902d60e01b8152938491829088165afa6000928161078a575b5061074d575050604051634c9c8ce360e01b81526001600160a01b0390921690820190815281906020010390fd5b838360008051602061153c833981519152840361076e5761001983836112a6565b604051632a87526960e21b815290810184815281906020010390fd5b6107ad91935060203d6020116107b4575b6107a581836105ef565b81019061109a565b913861071f565b503d61079b565b60405163703e46dd60e11b81528390fd5b90508160008051602061153c83398151915254161415386106ee565b346101a95760003660031901126101a9577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316300361084157602060405160008051602061153c8339815191528152f35b60405163703e46dd60e11b8152600490fd5b346101a95760003660031901126101a9576000546040516001600160a01b039091168152602090f35b346101a9576003196040368201126101a9576001600160401b03906004358281116101a9576108af903690600401610253565b6108b761111a565b6108c46060820182610e95565b9290836004116101a95760609382810161090f575b6040517fd3fddfd1276d1cc278f10907710a44474a32f917b2fcfa198f46ca7689215e2f908061090a888883610fb0565b0390a1005b8160409293949550019282840301126101a9576004810135610930816101ae565b60248201359485116101a95761096f60009360047fd3fddfd1276d1cc278f10907710a44474a32f917b2fcfa198f46ca7689215e2f9786950101610662565b80519160209091019083906001600160a01b03165af190610997610991610eee565b92610f1e565b3880806108d9565b919082519283825260005b8481106109cb575050826000602080949584010152601f8019910116010190565b6020818301810151848301820152016109aa565b346101a95760003660031901126101a95760405160408101908082106001600160401b038311176105ea576102ed9160405260058152640352e302e360dc1b602082015260405191829160208352602083019061099f565b346101a95760003660031901126101a9576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101a95760603660031901126101a957600435610a99816101ae565b604435906001600160401b0382116101a957610ac7610abf6100199336906004016101cc565b61043c61111a565b90602435906111d0565b346101a95760a03660031901126101a957610aed6004356101ae565b610af86024356101ae565b6001600160401b036044358181116101a957610b18903690600401610350565b50506064358181116101a957610b32903690600401610350565b50506084359081116101a957610b4c9036906004016101cc565b505060405163bc197c8160e01b8152602090f35b346101a95760003660031901126101a9576040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561052457602091600091610bd0575b50604051908152f35b610be79150823d84116107b4576107a581836105ef565b38610bc7565b346101a95760203660031901126101a957600435610c0a816101ae565b60008051602061155c83398151915254906001600160401b0360ff8360401c1615921680159081610d1e575b6001149081610d14575b159081610d0b575b50610cf95760008051602061155c833981519152805467ffffffffffffffff19166001179055610c7c9082610ccf5761123e565b610c8257005b60008051602061155c833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290806020810161090a565b60008051602061155c833981519152805460ff60401b19166801000000000000000017905561123e565b60405163f92ee8a960e01b8152600490fd5b90501538610c48565b303b159150610c40565b839150610c36565b346101a95760003660031901126101a957604051631aab3f0d60e11b8152306004820152600060248201526020816044817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015610524576102ed91600091610da557506040519081529081906020820190565b610dbe915060203d6020116107b4576107a581836105ef565b386102dc565b346101a95760a03660031901126101a957610de06004356101ae565b610deb6024356101ae565b6084356001600160401b0381116101a957610e0a9036906004016101cc565b505060405163f23a6e6160e01b8152602090f35b15610e2557565b60405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b6044820152606490fd5b634e487b7160e01b600052603260045260246000fd5b9190811015610e865760051b0190565b610e60565b3561067d816101ae565b903590601e19813603018212156101a957018035906001600160401b0382116101a9576020019181360383136101a957565b90821015610e8657610ede9160051b810190610e95565b9091565b6040513d6000823e3d90fd5b3d15610f19573d90610eff82610610565b91610f0d60405193846105ef565b82523d6000602084013e565b606090565b15610f2557565b60405162461bcd60e51b81526020600482015260116024820152701a5b9b995c8818d85b1b0819985a5b1959607a1b6044820152606490fd5b9035601e19823603018112156101a95701602081359101916001600160401b0382116101a95781360383136101a957565b908060209392818452848401376000828201840152601f01601f1916010190565b909161108c61100e61067d9460408552610fdd60408601610fd0836101bf565b6001600160a01b03169052565b6020810135606086015261107c610ff76040830183610f5e565b9390610120948560808a0152610160890191610f8f565b916110736110366110226060840184610f5e565b603f198b8803810160a08d01529691610f8f565b608083013560c08a015260a083013560e08a01526101009560c0840135878b015261106460e0850185610f5e565b91878c850301908c0152610f8f565b93810190610f5e565b9186840301610140870152610f8f565b91602081840391015261099f565b908160209103126101a9575190565b907f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52603c60002061110a61110160018060a01b03926110fb61043c856000541696610100810190610e95565b9061134d565b909291926113a9565b160361111557600090565b600190565b60018060a01b03807f000000000000000000000000000000000000000000000000000000000000000016331490811561119a575b501561115657565b606460405162461bcd60e51b815260206004820152602060248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152fd5b90506000541633143861114e565b600091829182602083519301915af16111bf610eee565b90156111c85750565b602081519101fd5b916000928392602083519301915af16111bf610eee565b6000546001600160a01b031633148015611235575b1561120357565b60405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b6044820152606490fd5b503033146111fc565b600080546001600160a01b0319166001600160a01b039283169081178255917f000000000000000000000000000000000000000000000000000000000000000016907f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de9080a3565b90813b1561132c5760008051602061153c83398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a28051156113115761130e91611436565b50565b50503461131a57565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b815191906041830361137e5761137792506020820151906060604084015193015160001a90611454565b9192909190565b505060009160029190565b6004111561139357565b634e487b7160e01b600052602160045260246000fd5b6113b281611389565b806113bb575050565b6113c481611389565b600181036113de5760405163f645eedf60e01b8152600490fd5b6113e781611389565b600281036114085760405163fce698f760e01b815260048101839052602490fd5b80611414600392611389565b1461141c5750565b6040516335e2f38360e21b81526004810191909152602490fd5b60008061067d93602081519101845af461144e610eee565b916114d8565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084116114cc57926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa156105245780516001600160a01b038116156114c357918190565b50809160019190565b50505060009160039190565b906114ff57508051156114ed57805190602001fd5b60405163d6bda27560e01b8152600490fd5b81511580611532575b611510575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561150856fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122096cbc01db3264920b2819047e9ea0f7bcde2fec41bba2fb0a697c3459794e82d64736f6c63430008190033",
}

// TestExecAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use TestExecAccountMetaData.ABI instead.
var TestExecAccountABI = TestExecAccountMetaData.ABI

// TestExecAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestExecAccountMetaData.Bin instead.
var TestExecAccountBin = TestExecAccountMetaData.Bin

// DeployTestExecAccount deploys a new Ethereum contract, binding an instance of TestExecAccount to it.
func DeployTestExecAccount(auth *bind.TransactOpts, backend bind.ContractBackend, anEntryPoint common.Address) (common.Address, *types.Transaction, *TestExecAccount, error) {
	parsed, err := TestExecAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestExecAccountBin), backend, anEntryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestExecAccount{TestExecAccountCaller: TestExecAccountCaller{contract: contract}, TestExecAccountTransactor: TestExecAccountTransactor{contract: contract}, TestExecAccountFilterer: TestExecAccountFilterer{contract: contract}}, nil
}

// TestExecAccount is an auto generated Go binding around an Ethereum contract.
type TestExecAccount struct {
	TestExecAccountCaller     // Read-only binding to the contract
	TestExecAccountTransactor // Write-only binding to the contract
	TestExecAccountFilterer   // Log filterer for contract events
}

// TestExecAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestExecAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExecAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestExecAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExecAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestExecAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExecAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestExecAccountSession struct {
	Contract     *TestExecAccount  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestExecAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestExecAccountCallerSession struct {
	Contract *TestExecAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TestExecAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestExecAccountTransactorSession struct {
	Contract     *TestExecAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TestExecAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestExecAccountRaw struct {
	Contract *TestExecAccount // Generic contract binding to access the raw methods on
}

// TestExecAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestExecAccountCallerRaw struct {
	Contract *TestExecAccountCaller // Generic read-only contract binding to access the raw methods on
}

// TestExecAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestExecAccountTransactorRaw struct {
	Contract *TestExecAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestExecAccount creates a new instance of TestExecAccount, bound to a specific deployed contract.
func NewTestExecAccount(address common.Address, backend bind.ContractBackend) (*TestExecAccount, error) {
	contract, err := bindTestExecAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestExecAccount{TestExecAccountCaller: TestExecAccountCaller{contract: contract}, TestExecAccountTransactor: TestExecAccountTransactor{contract: contract}, TestExecAccountFilterer: TestExecAccountFilterer{contract: contract}}, nil
}

// NewTestExecAccountCaller creates a new read-only instance of TestExecAccount, bound to a specific deployed contract.
func NewTestExecAccountCaller(address common.Address, caller bind.ContractCaller) (*TestExecAccountCaller, error) {
	contract, err := bindTestExecAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestExecAccountCaller{contract: contract}, nil
}

// NewTestExecAccountTransactor creates a new write-only instance of TestExecAccount, bound to a specific deployed contract.
func NewTestExecAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*TestExecAccountTransactor, error) {
	contract, err := bindTestExecAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestExecAccountTransactor{contract: contract}, nil
}

// NewTestExecAccountFilterer creates a new log filterer instance of TestExecAccount, bound to a specific deployed contract.
func NewTestExecAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*TestExecAccountFilterer, error) {
	contract, err := bindTestExecAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestExecAccountFilterer{contract: contract}, nil
}

// bindTestExecAccount binds a generic wrapper to an already deployed contract.
func bindTestExecAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestExecAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestExecAccount *TestExecAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestExecAccount.Contract.TestExecAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestExecAccount *TestExecAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExecAccount.Contract.TestExecAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestExecAccount *TestExecAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestExecAccount.Contract.TestExecAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestExecAccount *TestExecAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestExecAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestExecAccount *TestExecAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExecAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestExecAccount *TestExecAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestExecAccount.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TestExecAccount *TestExecAccountCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestExecAccount.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TestExecAccount *TestExecAccountSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TestExecAccount.Contract.UPGRADEINTERFACEVERSION(&_TestExecAccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TestExecAccount *TestExecAccountCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TestExecAccount.Contract.UPGRADEINTERFACEVERSION(&_TestExecAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestExecAccount *TestExecAccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestExecAccount.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestExecAccount *TestExecAccountSession) EntryPoint() (common.Address, error) {
	return _TestExecAccount.Contract.EntryPoint(&_TestExecAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestExecAccount *TestExecAccountCallerSession) EntryPoint() (common.Address, error) {
	return _TestExecAccount.Contract.EntryPoint(&_TestExecAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestExecAccount *TestExecAccountCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestExecAccount.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestExecAccount *TestExecAccountSession) GetDeposit() (*big.Int, error) {
	return _TestExecAccount.Contract.GetDeposit(&_TestExecAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestExecAccount *TestExecAccountCallerSession) GetDeposit() (*big.Int, error) {
	return _TestExecAccount.Contract.GetDeposit(&_TestExecAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_TestExecAccount *TestExecAccountCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestExecAccount.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_TestExecAccount *TestExecAccountSession) GetNonce() (*big.Int, error) {
	return _TestExecAccount.Contract.GetNonce(&_TestExecAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_TestExecAccount *TestExecAccountCallerSession) GetNonce() (*big.Int, error) {
	return _TestExecAccount.Contract.GetNonce(&_TestExecAccount.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_TestExecAccount *TestExecAccountCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TestExecAccount.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_TestExecAccount *TestExecAccountSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _TestExecAccount.Contract.OnERC1155BatchReceived(&_TestExecAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_TestExecAccount *TestExecAccountCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _TestExecAccount.Contract.OnERC1155BatchReceived(&_TestExecAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_TestExecAccount *TestExecAccountCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TestExecAccount.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_TestExecAccount *TestExecAccountSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _TestExecAccount.Contract.OnERC1155Received(&_TestExecAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_TestExecAccount *TestExecAccountCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _TestExecAccount.Contract.OnERC1155Received(&_TestExecAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TestExecAccount *TestExecAccountCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TestExecAccount.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TestExecAccount *TestExecAccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _TestExecAccount.Contract.OnERC721Received(&_TestExecAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TestExecAccount *TestExecAccountCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _TestExecAccount.Contract.OnERC721Received(&_TestExecAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestExecAccount *TestExecAccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestExecAccount.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestExecAccount *TestExecAccountSession) Owner() (common.Address, error) {
	return _TestExecAccount.Contract.Owner(&_TestExecAccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestExecAccount *TestExecAccountCallerSession) Owner() (common.Address, error) {
	return _TestExecAccount.Contract.Owner(&_TestExecAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TestExecAccount *TestExecAccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TestExecAccount.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TestExecAccount *TestExecAccountSession) ProxiableUUID() ([32]byte, error) {
	return _TestExecAccount.Contract.ProxiableUUID(&_TestExecAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TestExecAccount *TestExecAccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TestExecAccount.Contract.ProxiableUUID(&_TestExecAccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestExecAccount *TestExecAccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TestExecAccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestExecAccount *TestExecAccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestExecAccount.Contract.SupportsInterface(&_TestExecAccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestExecAccount *TestExecAccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestExecAccount.Contract.SupportsInterface(&_TestExecAccount.CallOpts, interfaceId)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_TestExecAccount *TestExecAccountTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExecAccount.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_TestExecAccount *TestExecAccountSession) AddDeposit() (*types.Transaction, error) {
	return _TestExecAccount.Contract.AddDeposit(&_TestExecAccount.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_TestExecAccount *TestExecAccountTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _TestExecAccount.Contract.AddDeposit(&_TestExecAccount.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_TestExecAccount *TestExecAccountTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _TestExecAccount.contract.Transact(opts, "execute", dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_TestExecAccount *TestExecAccountSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _TestExecAccount.Contract.Execute(&_TestExecAccount.TransactOpts, dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_TestExecAccount *TestExecAccountTransactorSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _TestExecAccount.Contract.Execute(&_TestExecAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_TestExecAccount *TestExecAccountTransactor) ExecuteBatch(opts *bind.TransactOpts, dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _TestExecAccount.contract.Transact(opts, "executeBatch", dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_TestExecAccount *TestExecAccountSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _TestExecAccount.Contract.ExecuteBatch(&_TestExecAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_TestExecAccount *TestExecAccountTransactorSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _TestExecAccount.Contract.ExecuteBatch(&_TestExecAccount.TransactOpts, dest, value, arg2)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x8dd7712f.
//
// Solidity: function executeUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 ) returns()
func (_TestExecAccount *TestExecAccountTransactor) ExecuteUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, arg1 [32]byte) (*types.Transaction, error) {
	return _TestExecAccount.contract.Transact(opts, "executeUserOp", userOp, arg1)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x8dd7712f.
//
// Solidity: function executeUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 ) returns()
func (_TestExecAccount *TestExecAccountSession) ExecuteUserOp(userOp PackedUserOperation, arg1 [32]byte) (*types.Transaction, error) {
	return _TestExecAccount.Contract.ExecuteUserOp(&_TestExecAccount.TransactOpts, userOp, arg1)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x8dd7712f.
//
// Solidity: function executeUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 ) returns()
func (_TestExecAccount *TestExecAccountTransactorSession) ExecuteUserOp(userOp PackedUserOperation, arg1 [32]byte) (*types.Transaction, error) {
	return _TestExecAccount.Contract.ExecuteUserOp(&_TestExecAccount.TransactOpts, userOp, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_TestExecAccount *TestExecAccountTransactor) Initialize(opts *bind.TransactOpts, anOwner common.Address) (*types.Transaction, error) {
	return _TestExecAccount.contract.Transact(opts, "initialize", anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_TestExecAccount *TestExecAccountSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _TestExecAccount.Contract.Initialize(&_TestExecAccount.TransactOpts, anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_TestExecAccount *TestExecAccountTransactorSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _TestExecAccount.Contract.Initialize(&_TestExecAccount.TransactOpts, anOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TestExecAccount *TestExecAccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TestExecAccount.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TestExecAccount *TestExecAccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TestExecAccount.Contract.UpgradeToAndCall(&_TestExecAccount.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TestExecAccount *TestExecAccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TestExecAccount.Contract.UpgradeToAndCall(&_TestExecAccount.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestExecAccount *TestExecAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestExecAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestExecAccount *TestExecAccountSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestExecAccount.Contract.ValidateUserOp(&_TestExecAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestExecAccount *TestExecAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestExecAccount.Contract.ValidateUserOp(&_TestExecAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_TestExecAccount *TestExecAccountTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestExecAccount.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_TestExecAccount *TestExecAccountSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestExecAccount.Contract.WithdrawDepositTo(&_TestExecAccount.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_TestExecAccount *TestExecAccountTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestExecAccount.Contract.WithdrawDepositTo(&_TestExecAccount.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestExecAccount *TestExecAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExecAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestExecAccount *TestExecAccountSession) Receive() (*types.Transaction, error) {
	return _TestExecAccount.Contract.Receive(&_TestExecAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestExecAccount *TestExecAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _TestExecAccount.Contract.Receive(&_TestExecAccount.TransactOpts)
}

// TestExecAccountExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the TestExecAccount contract.
type TestExecAccountExecutedIterator struct {
	Event *TestExecAccountExecuted // Event containing the contract specifics and raw log

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
func (it *TestExecAccountExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestExecAccountExecuted)
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
		it.Event = new(TestExecAccountExecuted)
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
func (it *TestExecAccountExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestExecAccountExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestExecAccountExecuted represents a Executed event raised by the TestExecAccount contract.
type TestExecAccountExecuted struct {
	UserOp       PackedUserOperation
	InnerCallRet []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xd3fddfd1276d1cc278f10907710a44474a32f917b2fcfa198f46ca7689215e2f.
//
// Solidity: event Executed((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes innerCallRet)
func (_TestExecAccount *TestExecAccountFilterer) FilterExecuted(opts *bind.FilterOpts) (*TestExecAccountExecutedIterator, error) {

	logs, sub, err := _TestExecAccount.contract.FilterLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return &TestExecAccountExecutedIterator{contract: _TestExecAccount.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xd3fddfd1276d1cc278f10907710a44474a32f917b2fcfa198f46ca7689215e2f.
//
// Solidity: event Executed((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes innerCallRet)
func (_TestExecAccount *TestExecAccountFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *TestExecAccountExecuted) (event.Subscription, error) {

	logs, sub, err := _TestExecAccount.contract.WatchLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestExecAccountExecuted)
				if err := _TestExecAccount.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0xd3fddfd1276d1cc278f10907710a44474a32f917b2fcfa198f46ca7689215e2f.
//
// Solidity: event Executed((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes innerCallRet)
func (_TestExecAccount *TestExecAccountFilterer) ParseExecuted(log types.Log) (*TestExecAccountExecuted, error) {
	event := new(TestExecAccountExecuted)
	if err := _TestExecAccount.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestExecAccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TestExecAccount contract.
type TestExecAccountInitializedIterator struct {
	Event *TestExecAccountInitialized // Event containing the contract specifics and raw log

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
func (it *TestExecAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestExecAccountInitialized)
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
		it.Event = new(TestExecAccountInitialized)
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
func (it *TestExecAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestExecAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestExecAccountInitialized represents a Initialized event raised by the TestExecAccount contract.
type TestExecAccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TestExecAccount *TestExecAccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*TestExecAccountInitializedIterator, error) {

	logs, sub, err := _TestExecAccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TestExecAccountInitializedIterator{contract: _TestExecAccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TestExecAccount *TestExecAccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TestExecAccountInitialized) (event.Subscription, error) {

	logs, sub, err := _TestExecAccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestExecAccountInitialized)
				if err := _TestExecAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TestExecAccount *TestExecAccountFilterer) ParseInitialized(log types.Log) (*TestExecAccountInitialized, error) {
	event := new(TestExecAccountInitialized)
	if err := _TestExecAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestExecAccountSimpleAccountInitializedIterator is returned from FilterSimpleAccountInitialized and is used to iterate over the raw logs and unpacked data for SimpleAccountInitialized events raised by the TestExecAccount contract.
type TestExecAccountSimpleAccountInitializedIterator struct {
	Event *TestExecAccountSimpleAccountInitialized // Event containing the contract specifics and raw log

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
func (it *TestExecAccountSimpleAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestExecAccountSimpleAccountInitialized)
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
		it.Event = new(TestExecAccountSimpleAccountInitialized)
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
func (it *TestExecAccountSimpleAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestExecAccountSimpleAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestExecAccountSimpleAccountInitialized represents a SimpleAccountInitialized event raised by the TestExecAccount contract.
type TestExecAccountSimpleAccountInitialized struct {
	EntryPoint common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSimpleAccountInitialized is a free log retrieval operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_TestExecAccount *TestExecAccountFilterer) FilterSimpleAccountInitialized(opts *bind.FilterOpts, entryPoint []common.Address, owner []common.Address) (*TestExecAccountSimpleAccountInitializedIterator, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TestExecAccount.contract.FilterLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TestExecAccountSimpleAccountInitializedIterator{contract: _TestExecAccount.contract, event: "SimpleAccountInitialized", logs: logs, sub: sub}, nil
}

// WatchSimpleAccountInitialized is a free log subscription operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_TestExecAccount *TestExecAccountFilterer) WatchSimpleAccountInitialized(opts *bind.WatchOpts, sink chan<- *TestExecAccountSimpleAccountInitialized, entryPoint []common.Address, owner []common.Address) (event.Subscription, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TestExecAccount.contract.WatchLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestExecAccountSimpleAccountInitialized)
				if err := _TestExecAccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
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
func (_TestExecAccount *TestExecAccountFilterer) ParseSimpleAccountInitialized(log types.Log) (*TestExecAccountSimpleAccountInitialized, error) {
	event := new(TestExecAccountSimpleAccountInitialized)
	if err := _TestExecAccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestExecAccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TestExecAccount contract.
type TestExecAccountUpgradedIterator struct {
	Event *TestExecAccountUpgraded // Event containing the contract specifics and raw log

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
func (it *TestExecAccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestExecAccountUpgraded)
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
		it.Event = new(TestExecAccountUpgraded)
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
func (it *TestExecAccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestExecAccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestExecAccountUpgraded represents a Upgraded event raised by the TestExecAccount contract.
type TestExecAccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TestExecAccount *TestExecAccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TestExecAccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TestExecAccount.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TestExecAccountUpgradedIterator{contract: _TestExecAccount.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TestExecAccount *TestExecAccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TestExecAccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TestExecAccount.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestExecAccountUpgraded)
				if err := _TestExecAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TestExecAccount *TestExecAccountFilterer) ParseUpgraded(log types.Log) (*TestExecAccountUpgraded, error) {
	event := new(TestExecAccountUpgraded)
	if err := _TestExecAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
