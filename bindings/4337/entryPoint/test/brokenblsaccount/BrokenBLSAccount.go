// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package brokenblsaccount

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

// BrokenBLSAccountMetaData contains all meta data concerning the BrokenBLSAccount contract.
var BrokenBLSAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"anAggregator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"oldPublicKey\",\"type\":\"uint256[4]\"},{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"newPublicKey\",\"type\":\"uint256[4]\"}],\"name\":\"PublicKeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SimpleAccountInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlsPublicKey\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"aPublicKey\",\"type\":\"uint256[4]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e03461016b57611507906001600160401b03601f38849003908101601f191683019082821184831017610170578084916040968794855283398101031261016b5781516001600160a01b0392838216820361016b5760200151928316830361016b573060805260a0527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff82861c1661015a578080831603610116575b50505060c052516113809081610187823960805181818161076801526108a1015260a0518181816102a00152818161051d015281816105bf015281816109b501528181610af401528181610cb8015281816110080152818161113401526111a3015260c0518181816102dd01526103830152f35b6001600160401b0319909116811790915582519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a13880806100a2565b845163f92ee8a960e01b8152600490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c806301ffc9a71461015b578063150b7a021461015657806319822f7c14610151578063245a7bfc1461014c57806347e1da2a146101475780634a58db19146101425780634d44560d1461013d5780634f1ef2861461013857806352d1902d146101335780638da5cb5b1461012e578063ad3cb1cc14610129578063b0d691fe14610124578063b61d27f61461011f578063bc197c811461011a578063c399ec8814610115578063c4d66de814610110578063d087d2881461010b578063e02afbae14610106578063ee472f36146101015763f23a6e610361000e57610e80565b610d84565b610d24565b610c85565b610b55565b610ac8565b610a39565b6109e4565b61099f565b610922565b6108f9565b61088e565b610715565b61058e565b61050e565b6103e2565b61036d565b610266565b61020c565b346101c95760203660031901126101c95760043563ffffffff60e01b81168091036101c957602090630a85bd0160e11b81149081156101b8575b81156101a7575b506040519015158152f35b6301ffc9a760e01b1490503861019c565b630271189760e51b81149150610195565b600080fd5b6001600160a01b038116036101c957565b9181601f840112156101c9578235916001600160401b0383116101c957602083818601950101116101c957565b346101c95760803660031901126101c9576102286004356101ce565b6102336024356101ce565b6064356001600160401b0381116101c9576102529036906004016101df565b5050604051630a85bd0160e11b8152602090f35b346101c9576003196060368201126101c957600435906001600160401b0382116101c95761012091360301126101c9576001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811633036103285761032490600060408051926102db84610639565b7f000000000000000000000000000000000000000000000000000000000000000016928381528260208201520152610314604435610fe0565b6040519081529081906020820190565b0390f35b60405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e74000000006044820152606490fd5b346101c95760003660031901126101c9576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b9181601f840112156101c9578235916001600160401b0383116101c9576020808501948460051b0101116101c957565b346101c95760603660031901126101c9576001600160401b036004358181116101c9576104139036906004016103b2565b6024929192358281116101c95761042e9036906004016103b2565b926044359081116101c9576104479036906004016103b2565b939091610452610ffe565b848414806104fd575b61046490610eda565b816104b157505060005b82811061047757005b806104ab61049061048b600194878a610f32565b610f47565b6104a561049e848988610f54565b36916106de565b9061108c565b0161046e565b91909460009493945b8581106104c357005b806104f76104d761048b6001948a87610f32565b6104e2838b89610f32565b356104f161049e858b8a610f54565b916110b4565b016104ba565b5081158061045b575081851461045b565b60008060031936011261058b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681813b1561058b5760405163b760faf960e01b8152306004820152918290602490829034905af180156105865761057a575080f35b61058390610659565b80f35b610f95565b80fd5b346101c9576000604036600319011261058b576004356105ad816101ce565b6105b56110cb565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b1561061f5760449083604051958694859363040b850f60e31b855216600484015260243560248401525af180156105865761057a575080f35b8280fd5b634e487b7160e01b600052604160045260246000fd5b606081019081106001600160401b0382111761065457604052565b610623565b6001600160401b03811161065457604052565b604081019081106001600160401b0382111761065457604052565b608081019081106001600160401b0382111761065457604052565b90601f801991011681019081106001600160401b0382111761065457604052565b6001600160401b03811161065457601f01601f191660200190565b9291926106ea826106c3565b916106f860405193846106a2565b8294818452818301116101c9578281602093846000960137010152565b60403660031901126101c957600480359061072f826101ce565b6024356001600160401b0381116101c957366023820112156101c95761075e90369060248185013591016106de565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116308114908115610872575b506108615790602083926107a66110cb565b6040516352d1902d60e01b8152938491829088165afa60009281610830575b506107f3575050604051634c9c8ce360e01b81526001600160a01b0390921690820190815281906020010390fd5b838360008051602061130b83398151915284036108145761001983836111ea565b604051632a87526960e21b815290810184815281906020010390fd5b61085391935060203d60201161085a575b61084b81836106a2565b810190610fa1565b91386107c5565b503d610841565b60405163703e46dd60e11b81528390fd5b90508160008051602061130b8339815191525416141538610794565b346101c95760003660031901126101c9577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036108e757602060405160008051602061130b8339815191528152f35b60405163703e46dd60e11b8152600490fd5b346101c95760003660031901126101c9576000546040516001600160a01b039091168152602090f35b346101c95760003660031901126101c95760408051906109418261066c565b60058252602090640352e302e360dc1b6020840152604051916020835283519182602085015260005b83811061098c5784604081866000838284010152601f80199101168101030190f35b858101830151858201830152820161096a565b346101c95760003660031901126101c9576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101c95760603660031901126101c957600435610a01816101ce565b604435906001600160401b0382116101c957610a2f610a276100199336906004016101df565b61049e610ffe565b90602435906110b4565b346101c95760a03660031901126101c957610a556004356101ce565b610a606024356101ce565b6001600160401b036044358181116101c957610a809036906004016103b2565b50506064358181116101c957610a9a9036906004016103b2565b50506084359081116101c957610ab49036906004016101df565b505060405163bc197c8160e01b8152602090f35b346101c95760003660031901126101c9576040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561058657602091600091610b38575b50604051908152f35b610b4f9150823d841161085a5761084b81836106a2565b38610b2f565b346101c95760203660031901126101c957600435610b72816101ce565b60008051602061132b83398151915254906001600160401b0360ff8360401c1615921680159081610c7d575b6001149081610c73575b159081610c6a575b50610c585760008051602061132b833981519152805467ffffffffffffffff19166001179055610be49082610c3357611182565b610bea57005b60008051602061132b833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b60008051602061132b833981519152805460ff60401b1916600160401b179055611182565b60405163f92ee8a960e01b8152600490fd5b90501538610bb0565b303b159150610ba8565b839150610b9e565b346101c95760003660031901126101c957604051631aab3f0d60e11b8152306004820152600060248201526020816044817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156105865761032491600091610d05575b506040519081529081906020820190565b610d1e915060203d60201161085a5761084b81836106a2565b38610cf4565b346101c95760003660031901126101c9576080604051610d4381610687565b369037604051610d5281610687565b608036823760405190600090825b60048310610d6d57608084f35b600190825181526020809101920192019190610d60565b346101c95760803660031901126101c95736602312156101c957604051610daa81610687565b6084366084116101c9576004915b818310610e705760008051602061132b833981519152546001600160401b0360ff8260401c1615911680159081610e68575b6001149081610e5e575b159081610e55575b50610c585760008051602061132b833981519152805467ffffffffffffffff1916600117905580610e30575b610be4611122565b60008051602061132b833981519152805460ff60401b1916600160401b179055610e28565b90501582610dfc565b303b159150610df4565b829150610dea565b8235815260209283019201610db8565b346101c95760a03660031901126101c957610e9c6004356101ce565b610ea76024356101ce565b6084356001600160401b0381116101c957610ec69036906004016101df565b505060405163f23a6e6160e01b8152602090f35b15610ee157565b60405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b6044820152606490fd5b634e487b7160e01b600052603260045260246000fd5b9190811015610f425760051b0190565b610f1c565b35610f51816101ce565b90565b9190811015610f425760051b81013590601e19813603018212156101c95701908135916001600160401b0383116101c95760200182360381136101c9579190565b6040513d6000823e3d90fd5b908160209103126101c9575190565b3d15610fdb573d90610fc1826106c3565b91610fcf60405193846106a2565b82523d6000602084013e565b606090565b80610fe85750565b600080808093338219f150610ffb610fb0565b50565b60018060a01b03807f000000000000000000000000000000000000000000000000000000000000000016331490811561107e575b501561103a57565b606460405162461bcd60e51b815260206004820152602060248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152fd5b905060005416331438611032565b600091829182602083519301915af16110a3610fb0565b90156110ac5750565b602081519101fd5b916000928392602083519301915af16110a3610fb0565b6000546001600160a01b031633148015611119575b156110e757565b60405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b6044820152606490fd5b503033146110e0565b600080546001600160a01b03191681557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de8280a3565b600080546001600160a01b0319166001600160a01b039283169081178255917f000000000000000000000000000000000000000000000000000000000000000016907f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de9080a3565b90813b1561126d5760008051602061130b83398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a280511561125257610ffb9161128e565b50503461125b57565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b600080610f5193602081519101845af46112a6610fb0565b91906112ce57508051156112bc57805190602001fd5b60405163d6bda27560e01b8152600490fd5b81511580611301575b6112df575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b156112d756fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212203c2da3c5641277907952b8f68fe7b1bfb2476bf609bb31900585a098d8726e2364736f6c63430008190033",
}

// BrokenBLSAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use BrokenBLSAccountMetaData.ABI instead.
var BrokenBLSAccountABI = BrokenBLSAccountMetaData.ABI

// BrokenBLSAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BrokenBLSAccountMetaData.Bin instead.
var BrokenBLSAccountBin = BrokenBLSAccountMetaData.Bin

// DeployBrokenBLSAccount deploys a new Ethereum contract, binding an instance of BrokenBLSAccount to it.
func DeployBrokenBLSAccount(auth *bind.TransactOpts, backend bind.ContractBackend, anEntryPoint common.Address, anAggregator common.Address) (common.Address, *types.Transaction, *BrokenBLSAccount, error) {
	parsed, err := BrokenBLSAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BrokenBLSAccountBin), backend, anEntryPoint, anAggregator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BrokenBLSAccount{BrokenBLSAccountCaller: BrokenBLSAccountCaller{contract: contract}, BrokenBLSAccountTransactor: BrokenBLSAccountTransactor{contract: contract}, BrokenBLSAccountFilterer: BrokenBLSAccountFilterer{contract: contract}}, nil
}

// BrokenBLSAccount is an auto generated Go binding around an Ethereum contract.
type BrokenBLSAccount struct {
	BrokenBLSAccountCaller     // Read-only binding to the contract
	BrokenBLSAccountTransactor // Write-only binding to the contract
	BrokenBLSAccountFilterer   // Log filterer for contract events
}

// BrokenBLSAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrokenBLSAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokenBLSAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrokenBLSAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokenBLSAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrokenBLSAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokenBLSAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrokenBLSAccountSession struct {
	Contract     *BrokenBLSAccount // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokenBLSAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrokenBLSAccountCallerSession struct {
	Contract *BrokenBLSAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BrokenBLSAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrokenBLSAccountTransactorSession struct {
	Contract     *BrokenBLSAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BrokenBLSAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrokenBLSAccountRaw struct {
	Contract *BrokenBLSAccount // Generic contract binding to access the raw methods on
}

// BrokenBLSAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrokenBLSAccountCallerRaw struct {
	Contract *BrokenBLSAccountCaller // Generic read-only contract binding to access the raw methods on
}

// BrokenBLSAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrokenBLSAccountTransactorRaw struct {
	Contract *BrokenBLSAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrokenBLSAccount creates a new instance of BrokenBLSAccount, bound to a specific deployed contract.
func NewBrokenBLSAccount(address common.Address, backend bind.ContractBackend) (*BrokenBLSAccount, error) {
	contract, err := bindBrokenBLSAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BrokenBLSAccount{BrokenBLSAccountCaller: BrokenBLSAccountCaller{contract: contract}, BrokenBLSAccountTransactor: BrokenBLSAccountTransactor{contract: contract}, BrokenBLSAccountFilterer: BrokenBLSAccountFilterer{contract: contract}}, nil
}

// NewBrokenBLSAccountCaller creates a new read-only instance of BrokenBLSAccount, bound to a specific deployed contract.
func NewBrokenBLSAccountCaller(address common.Address, caller bind.ContractCaller) (*BrokenBLSAccountCaller, error) {
	contract, err := bindBrokenBLSAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrokenBLSAccountCaller{contract: contract}, nil
}

// NewBrokenBLSAccountTransactor creates a new write-only instance of BrokenBLSAccount, bound to a specific deployed contract.
func NewBrokenBLSAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*BrokenBLSAccountTransactor, error) {
	contract, err := bindBrokenBLSAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrokenBLSAccountTransactor{contract: contract}, nil
}

// NewBrokenBLSAccountFilterer creates a new log filterer instance of BrokenBLSAccount, bound to a specific deployed contract.
func NewBrokenBLSAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*BrokenBLSAccountFilterer, error) {
	contract, err := bindBrokenBLSAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrokenBLSAccountFilterer{contract: contract}, nil
}

// bindBrokenBLSAccount binds a generic wrapper to an already deployed contract.
func bindBrokenBLSAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BrokenBLSAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrokenBLSAccount *BrokenBLSAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrokenBLSAccount.Contract.BrokenBLSAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrokenBLSAccount *BrokenBLSAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.BrokenBLSAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrokenBLSAccount *BrokenBLSAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.BrokenBLSAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrokenBLSAccount *BrokenBLSAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrokenBLSAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrokenBLSAccount *BrokenBLSAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrokenBLSAccount *BrokenBLSAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BrokenBLSAccount *BrokenBLSAccountCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BrokenBLSAccount.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BrokenBLSAccount *BrokenBLSAccountSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BrokenBLSAccount.Contract.UPGRADEINTERFACEVERSION(&_BrokenBLSAccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BrokenBLSAccount *BrokenBLSAccountCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BrokenBLSAccount.Contract.UPGRADEINTERFACEVERSION(&_BrokenBLSAccount.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_BrokenBLSAccount *BrokenBLSAccountCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrokenBLSAccount.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_BrokenBLSAccount *BrokenBLSAccountSession) Aggregator() (common.Address, error) {
	return _BrokenBLSAccount.Contract.Aggregator(&_BrokenBLSAccount.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_BrokenBLSAccount *BrokenBLSAccountCallerSession) Aggregator() (common.Address, error) {
	return _BrokenBLSAccount.Contract.Aggregator(&_BrokenBLSAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_BrokenBLSAccount *BrokenBLSAccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrokenBLSAccount.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_BrokenBLSAccount *BrokenBLSAccountSession) EntryPoint() (common.Address, error) {
	return _BrokenBLSAccount.Contract.EntryPoint(&_BrokenBLSAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_BrokenBLSAccount *BrokenBLSAccountCallerSession) EntryPoint() (common.Address, error) {
	return _BrokenBLSAccount.Contract.EntryPoint(&_BrokenBLSAccount.CallOpts)
}

// GetBlsPublicKey is a free data retrieval call binding the contract method 0xe02afbae.
//
// Solidity: function getBlsPublicKey() pure returns(uint256[4])
func (_BrokenBLSAccount *BrokenBLSAccountCaller) GetBlsPublicKey(opts *bind.CallOpts) ([4]*big.Int, error) {
	var out []interface{}
	err := _BrokenBLSAccount.contract.Call(opts, &out, "getBlsPublicKey")

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetBlsPublicKey is a free data retrieval call binding the contract method 0xe02afbae.
//
// Solidity: function getBlsPublicKey() pure returns(uint256[4])
func (_BrokenBLSAccount *BrokenBLSAccountSession) GetBlsPublicKey() ([4]*big.Int, error) {
	return _BrokenBLSAccount.Contract.GetBlsPublicKey(&_BrokenBLSAccount.CallOpts)
}

// GetBlsPublicKey is a free data retrieval call binding the contract method 0xe02afbae.
//
// Solidity: function getBlsPublicKey() pure returns(uint256[4])
func (_BrokenBLSAccount *BrokenBLSAccountCallerSession) GetBlsPublicKey() ([4]*big.Int, error) {
	return _BrokenBLSAccount.Contract.GetBlsPublicKey(&_BrokenBLSAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_BrokenBLSAccount *BrokenBLSAccountCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrokenBLSAccount.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_BrokenBLSAccount *BrokenBLSAccountSession) GetDeposit() (*big.Int, error) {
	return _BrokenBLSAccount.Contract.GetDeposit(&_BrokenBLSAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_BrokenBLSAccount *BrokenBLSAccountCallerSession) GetDeposit() (*big.Int, error) {
	return _BrokenBLSAccount.Contract.GetDeposit(&_BrokenBLSAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_BrokenBLSAccount *BrokenBLSAccountCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrokenBLSAccount.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_BrokenBLSAccount *BrokenBLSAccountSession) GetNonce() (*big.Int, error) {
	return _BrokenBLSAccount.Contract.GetNonce(&_BrokenBLSAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_BrokenBLSAccount *BrokenBLSAccountCallerSession) GetNonce() (*big.Int, error) {
	return _BrokenBLSAccount.Contract.GetNonce(&_BrokenBLSAccount.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_BrokenBLSAccount *BrokenBLSAccountCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _BrokenBLSAccount.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_BrokenBLSAccount *BrokenBLSAccountSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _BrokenBLSAccount.Contract.OnERC1155BatchReceived(&_BrokenBLSAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_BrokenBLSAccount *BrokenBLSAccountCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _BrokenBLSAccount.Contract.OnERC1155BatchReceived(&_BrokenBLSAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_BrokenBLSAccount *BrokenBLSAccountCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _BrokenBLSAccount.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_BrokenBLSAccount *BrokenBLSAccountSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _BrokenBLSAccount.Contract.OnERC1155Received(&_BrokenBLSAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_BrokenBLSAccount *BrokenBLSAccountCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _BrokenBLSAccount.Contract.OnERC1155Received(&_BrokenBLSAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_BrokenBLSAccount *BrokenBLSAccountCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _BrokenBLSAccount.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_BrokenBLSAccount *BrokenBLSAccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _BrokenBLSAccount.Contract.OnERC721Received(&_BrokenBLSAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_BrokenBLSAccount *BrokenBLSAccountCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _BrokenBLSAccount.Contract.OnERC721Received(&_BrokenBLSAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrokenBLSAccount *BrokenBLSAccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrokenBLSAccount.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrokenBLSAccount *BrokenBLSAccountSession) Owner() (common.Address, error) {
	return _BrokenBLSAccount.Contract.Owner(&_BrokenBLSAccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrokenBLSAccount *BrokenBLSAccountCallerSession) Owner() (common.Address, error) {
	return _BrokenBLSAccount.Contract.Owner(&_BrokenBLSAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BrokenBLSAccount *BrokenBLSAccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BrokenBLSAccount.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BrokenBLSAccount *BrokenBLSAccountSession) ProxiableUUID() ([32]byte, error) {
	return _BrokenBLSAccount.Contract.ProxiableUUID(&_BrokenBLSAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BrokenBLSAccount *BrokenBLSAccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BrokenBLSAccount.Contract.ProxiableUUID(&_BrokenBLSAccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BrokenBLSAccount *BrokenBLSAccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BrokenBLSAccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BrokenBLSAccount *BrokenBLSAccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BrokenBLSAccount.Contract.SupportsInterface(&_BrokenBLSAccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BrokenBLSAccount *BrokenBLSAccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BrokenBLSAccount.Contract.SupportsInterface(&_BrokenBLSAccount.CallOpts, interfaceId)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrokenBLSAccount.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_BrokenBLSAccount *BrokenBLSAccountSession) AddDeposit() (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.AddDeposit(&_BrokenBLSAccount.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.AddDeposit(&_BrokenBLSAccount.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _BrokenBLSAccount.contract.Transact(opts, "execute", dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_BrokenBLSAccount *BrokenBLSAccountSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.Execute(&_BrokenBLSAccount.TransactOpts, dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactorSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.Execute(&_BrokenBLSAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactor) ExecuteBatch(opts *bind.TransactOpts, dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _BrokenBLSAccount.contract.Transact(opts, "executeBatch", dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_BrokenBLSAccount *BrokenBLSAccountSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.ExecuteBatch(&_BrokenBLSAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactorSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.ExecuteBatch(&_BrokenBLSAccount.TransactOpts, dest, value, arg2)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactor) Initialize(opts *bind.TransactOpts, anOwner common.Address) (*types.Transaction, error) {
	return _BrokenBLSAccount.contract.Transact(opts, "initialize", anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_BrokenBLSAccount *BrokenBLSAccountSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.Initialize(&_BrokenBLSAccount.TransactOpts, anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactorSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.Initialize(&_BrokenBLSAccount.TransactOpts, anOwner)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xee472f36.
//
// Solidity: function initialize(uint256[4] aPublicKey) returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactor) Initialize0(opts *bind.TransactOpts, aPublicKey [4]*big.Int) (*types.Transaction, error) {
	return _BrokenBLSAccount.contract.Transact(opts, "initialize0", aPublicKey)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xee472f36.
//
// Solidity: function initialize(uint256[4] aPublicKey) returns()
func (_BrokenBLSAccount *BrokenBLSAccountSession) Initialize0(aPublicKey [4]*big.Int) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.Initialize0(&_BrokenBLSAccount.TransactOpts, aPublicKey)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xee472f36.
//
// Solidity: function initialize(uint256[4] aPublicKey) returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactorSession) Initialize0(aPublicKey [4]*big.Int) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.Initialize0(&_BrokenBLSAccount.TransactOpts, aPublicKey)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BrokenBLSAccount.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BrokenBLSAccount *BrokenBLSAccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.UpgradeToAndCall(&_BrokenBLSAccount.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.UpgradeToAndCall(&_BrokenBLSAccount.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_BrokenBLSAccount *BrokenBLSAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _BrokenBLSAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_BrokenBLSAccount *BrokenBLSAccountSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.ValidateUserOp(&_BrokenBLSAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_BrokenBLSAccount *BrokenBLSAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.ValidateUserOp(&_BrokenBLSAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BrokenBLSAccount.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_BrokenBLSAccount *BrokenBLSAccountSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.WithdrawDepositTo(&_BrokenBLSAccount.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.WithdrawDepositTo(&_BrokenBLSAccount.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrokenBLSAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BrokenBLSAccount *BrokenBLSAccountSession) Receive() (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.Receive(&_BrokenBLSAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BrokenBLSAccount *BrokenBLSAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _BrokenBLSAccount.Contract.Receive(&_BrokenBLSAccount.TransactOpts)
}

// BrokenBLSAccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BrokenBLSAccount contract.
type BrokenBLSAccountInitializedIterator struct {
	Event *BrokenBLSAccountInitialized // Event containing the contract specifics and raw log

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
func (it *BrokenBLSAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokenBLSAccountInitialized)
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
		it.Event = new(BrokenBLSAccountInitialized)
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
func (it *BrokenBLSAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokenBLSAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokenBLSAccountInitialized represents a Initialized event raised by the BrokenBLSAccount contract.
type BrokenBLSAccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BrokenBLSAccount *BrokenBLSAccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*BrokenBLSAccountInitializedIterator, error) {

	logs, sub, err := _BrokenBLSAccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BrokenBLSAccountInitializedIterator{contract: _BrokenBLSAccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BrokenBLSAccount *BrokenBLSAccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BrokenBLSAccountInitialized) (event.Subscription, error) {

	logs, sub, err := _BrokenBLSAccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokenBLSAccountInitialized)
				if err := _BrokenBLSAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BrokenBLSAccount *BrokenBLSAccountFilterer) ParseInitialized(log types.Log) (*BrokenBLSAccountInitialized, error) {
	event := new(BrokenBLSAccountInitialized)
	if err := _BrokenBLSAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokenBLSAccountPublicKeyChangedIterator is returned from FilterPublicKeyChanged and is used to iterate over the raw logs and unpacked data for PublicKeyChanged events raised by the BrokenBLSAccount contract.
type BrokenBLSAccountPublicKeyChangedIterator struct {
	Event *BrokenBLSAccountPublicKeyChanged // Event containing the contract specifics and raw log

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
func (it *BrokenBLSAccountPublicKeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokenBLSAccountPublicKeyChanged)
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
		it.Event = new(BrokenBLSAccountPublicKeyChanged)
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
func (it *BrokenBLSAccountPublicKeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokenBLSAccountPublicKeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokenBLSAccountPublicKeyChanged represents a PublicKeyChanged event raised by the BrokenBLSAccount contract.
type BrokenBLSAccountPublicKeyChanged struct {
	OldPublicKey [4]*big.Int
	NewPublicKey [4]*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyChanged is a free log retrieval operation binding the contract event 0x42e4c4ce1432650f17e41c4ea77ed12c0ab20b229d3ffd84a2ebc9f8abb25a83.
//
// Solidity: event PublicKeyChanged(uint256[4] oldPublicKey, uint256[4] newPublicKey)
func (_BrokenBLSAccount *BrokenBLSAccountFilterer) FilterPublicKeyChanged(opts *bind.FilterOpts) (*BrokenBLSAccountPublicKeyChangedIterator, error) {

	logs, sub, err := _BrokenBLSAccount.contract.FilterLogs(opts, "PublicKeyChanged")
	if err != nil {
		return nil, err
	}
	return &BrokenBLSAccountPublicKeyChangedIterator{contract: _BrokenBLSAccount.contract, event: "PublicKeyChanged", logs: logs, sub: sub}, nil
}

// WatchPublicKeyChanged is a free log subscription operation binding the contract event 0x42e4c4ce1432650f17e41c4ea77ed12c0ab20b229d3ffd84a2ebc9f8abb25a83.
//
// Solidity: event PublicKeyChanged(uint256[4] oldPublicKey, uint256[4] newPublicKey)
func (_BrokenBLSAccount *BrokenBLSAccountFilterer) WatchPublicKeyChanged(opts *bind.WatchOpts, sink chan<- *BrokenBLSAccountPublicKeyChanged) (event.Subscription, error) {

	logs, sub, err := _BrokenBLSAccount.contract.WatchLogs(opts, "PublicKeyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokenBLSAccountPublicKeyChanged)
				if err := _BrokenBLSAccount.contract.UnpackLog(event, "PublicKeyChanged", log); err != nil {
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
func (_BrokenBLSAccount *BrokenBLSAccountFilterer) ParsePublicKeyChanged(log types.Log) (*BrokenBLSAccountPublicKeyChanged, error) {
	event := new(BrokenBLSAccountPublicKeyChanged)
	if err := _BrokenBLSAccount.contract.UnpackLog(event, "PublicKeyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokenBLSAccountSimpleAccountInitializedIterator is returned from FilterSimpleAccountInitialized and is used to iterate over the raw logs and unpacked data for SimpleAccountInitialized events raised by the BrokenBLSAccount contract.
type BrokenBLSAccountSimpleAccountInitializedIterator struct {
	Event *BrokenBLSAccountSimpleAccountInitialized // Event containing the contract specifics and raw log

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
func (it *BrokenBLSAccountSimpleAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokenBLSAccountSimpleAccountInitialized)
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
		it.Event = new(BrokenBLSAccountSimpleAccountInitialized)
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
func (it *BrokenBLSAccountSimpleAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokenBLSAccountSimpleAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokenBLSAccountSimpleAccountInitialized represents a SimpleAccountInitialized event raised by the BrokenBLSAccount contract.
type BrokenBLSAccountSimpleAccountInitialized struct {
	EntryPoint common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSimpleAccountInitialized is a free log retrieval operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_BrokenBLSAccount *BrokenBLSAccountFilterer) FilterSimpleAccountInitialized(opts *bind.FilterOpts, entryPoint []common.Address, owner []common.Address) (*BrokenBLSAccountSimpleAccountInitializedIterator, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BrokenBLSAccount.contract.FilterLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BrokenBLSAccountSimpleAccountInitializedIterator{contract: _BrokenBLSAccount.contract, event: "SimpleAccountInitialized", logs: logs, sub: sub}, nil
}

// WatchSimpleAccountInitialized is a free log subscription operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_BrokenBLSAccount *BrokenBLSAccountFilterer) WatchSimpleAccountInitialized(opts *bind.WatchOpts, sink chan<- *BrokenBLSAccountSimpleAccountInitialized, entryPoint []common.Address, owner []common.Address) (event.Subscription, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BrokenBLSAccount.contract.WatchLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokenBLSAccountSimpleAccountInitialized)
				if err := _BrokenBLSAccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
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
func (_BrokenBLSAccount *BrokenBLSAccountFilterer) ParseSimpleAccountInitialized(log types.Log) (*BrokenBLSAccountSimpleAccountInitialized, error) {
	event := new(BrokenBLSAccountSimpleAccountInitialized)
	if err := _BrokenBLSAccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokenBLSAccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BrokenBLSAccount contract.
type BrokenBLSAccountUpgradedIterator struct {
	Event *BrokenBLSAccountUpgraded // Event containing the contract specifics and raw log

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
func (it *BrokenBLSAccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokenBLSAccountUpgraded)
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
		it.Event = new(BrokenBLSAccountUpgraded)
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
func (it *BrokenBLSAccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokenBLSAccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokenBLSAccountUpgraded represents a Upgraded event raised by the BrokenBLSAccount contract.
type BrokenBLSAccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BrokenBLSAccount *BrokenBLSAccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BrokenBLSAccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BrokenBLSAccount.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BrokenBLSAccountUpgradedIterator{contract: _BrokenBLSAccount.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BrokenBLSAccount *BrokenBLSAccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BrokenBLSAccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BrokenBLSAccount.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokenBLSAccountUpgraded)
				if err := _BrokenBLSAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BrokenBLSAccount *BrokenBLSAccountFilterer) ParseUpgraded(log types.Log) (*BrokenBLSAccountUpgraded, error) {
	event := new(BrokenBLSAccountUpgraded)
	if err := _BrokenBLSAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
