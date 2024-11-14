// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testexpiryaccount

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

// TestExpiryAccountMetaData contains all meta data concerning the TestExpiryAccount contract.
var TestExpiryAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SimpleAccountInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"_after\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"_until\",\"type\":\"uint48\"}],\"name\":\"addTemporaryOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownerAfter\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownerUntil\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c03460ab57601f6115e238819003918201601f19168301916001600160401b0383118484101760b05780849260209460405283398101031260ab57516001600160a01b038116810360ab573060805260a05260405161151b90816100c782396080518181816107c401526108fd015260a0518181816102a2015281816105af0152818161065101528181610a1101528181610b5001528181610e1f01528181610feb01526110bb0152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c806301ffc9a71461015b578063150b7a021461015657806319822f7c146101515780633e4769511461014c5780633fb5a7a11461014757806347e1da2a146101425780634a58db191461013d5780634d44560d146101385780634f1ef2861461013357806352d1902d1461012e5780638da5cb5b14610129578063ad3cb1cc14610124578063b0d691fe1461011f578063b61d27f61461011a578063bc197c8114610115578063c399ec8814610110578063c4d66de81461010b578063cf6dca5514610106578063d087d288146101015763f23a6e610361000e57610e8b565b610dec565b610d10565b610bb1565b610b24565b610a95565b610a40565b6109fb565b61097e565b610955565b6108ea565b610771565b610620565b6105a0565b61047b565b610406565b6103c1565b610266565b61020c565b346101c95760203660031901126101c95760043563ffffffff60e01b81168091036101c957602090630a85bd0160e11b81149081156101b8575b81156101a7575b506040519015158152f35b6301ffc9a760e01b1490503861019c565b630271189760e51b81149150610195565b600080fd5b6001600160a01b038116036101c957565b9181601f840112156101c9578235916001600160401b0383116101c957602083818601950101116101c957565b346101c95760803660031901126101c9576102286004356101ce565b6102336024356101ce565b6064356001600160401b0381116101c9576102529036906004016101df565b5050604051630a85bd0160e11b8152602090f35b346101c9576003196060368201126101c957600435906001600160401b0382116101c9576101209082360301126101c9576001600160a01b03907f00000000000000000000000000000000000000000000000000000000000000008216330361037c5761032161032a917f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052602435601c5261031b610314603c60002092610104810190600401610f5f565b369161073a565b906112b7565b90929192611313565b16600052600260205261037861035d65ffffffffffff8060406000205416906001602052604060002054169080156111d5565b610368604435611093565b6040519081529081906020820190565b0390f35b60405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e74000000006044820152606490fd5b346101c95760203660031901126101c9576004356103de816101ce565b60018060a01b03166000526002602052602065ffffffffffff60406000205416604051908152f35b346101c95760203660031901126101c957600435610423816101ce565b60018060a01b03166000526001602052602065ffffffffffff60406000205416604051908152f35b9181601f840112156101c9578235916001600160401b0383116101c9576020808501948460051b0101116101c957565b346101c95760603660031901126101c9576001600160401b036004358181116101c9576104ac90369060040161044b565b6024929192358281116101c9576104c790369060040161044b565b926044359081116101c9576104e090369060040161044b565b9390916104eb6110b1565b8484148061058f575b6104fd90610ee5565b8161054357505060005b82811061051057005b8061053d610529610524600194878a610f3d565b610f52565b610537610314848988610f91565b9061113f565b01610507565b91909460009493945b85811061055557005b806105896105696105246001948a87610f3d565b610574838b89610f3d565b35610583610314858b8a610f91565b91611167565b0161054c565b508115806104f457508185146104f4565b60008060031936011261061d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681813b1561061d5760405163b760faf960e01b8152306004820152918290602490829034905af180156106185761060c575080f35b610615906106cb565b80f35b610fac565b80fd5b346101c9576000604036600319011261061d5760043561063f816101ce565b61064761117e565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b156106b15760449083604051958694859363040b850f60e31b855216600484015260243560248401525af180156106185761060c575080f35b8280fd5b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116106de57604052565b6106b5565b604081019081106001600160401b038211176106de57604052565b90601f801991011681019081106001600160401b038211176106de57604052565b6001600160401b0381116106de57601f01601f191660200190565b9291926107468261071f565b9161075460405193846106fe565b8294818452818301116101c9578281602093846000960137010152565b60403660031901126101c957600480359061078b826101ce565b6024356001600160401b0381116101c957366023820112156101c9576107ba903690602481850135910161073a565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081163081149081156108ce575b506108bd57906020839261080261117e565b6040516352d1902d60e01b8152938491829088165afa6000928161088c575b5061084f575050604051634c9c8ce360e01b81526001600160a01b0390921690820190815281906020010390fd5b83836000805160206114a68339815191528403610870576100198383611213565b604051632a87526960e21b815290810184815281906020010390fd5b6108af91935060203d6020116108b6575b6108a781836106fe565b810190610fb8565b9138610821565b503d61089d565b60405163703e46dd60e11b81528390fd5b9050816000805160206114a683398151915254161415386107f0565b346101c95760003660031901126101c9577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036109435760206040516000805160206114a68339815191528152f35b60405163703e46dd60e11b8152600490fd5b346101c95760003660031901126101c9576000546040516001600160a01b039091168152602090f35b346101c95760003660031901126101c957604080519061099d826106e3565b60058252602090640352e302e360dc1b6020840152604051916020835283519182602085015260005b8381106109e85784604081866000838284010152601f80199101168101030190f35b85810183015185820183015282016109c6565b346101c95760003660031901126101c9576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101c95760603660031901126101c957600435610a5d816101ce565b604435906001600160401b0382116101c957610a8b610a836100199336906004016101df565b6103146110b1565b9060243590611167565b346101c95760a03660031901126101c957610ab16004356101ce565b610abc6024356101ce565b6001600160401b036044358181116101c957610adc90369060040161044b565b50506064358181116101c957610af690369060040161044b565b50506084359081116101c957610b109036906004016101df565b505060405163bc197c8160e01b8152602090f35b346101c95760003660031901126101c9576040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561061857602091600091610b94575b50604051908152f35b610bab9150823d84116108b6576108a781836106fe565b38610b8b565b346101c95760203660031901126101c957600435610bce816101ce565b6000805160206114c683398151915254906001600160401b0360ff8360401c1615921680159081610cde575b6001149081610cd4575b159081610ccb575b50610cb9576000805160206114c6833981519152805467ffffffffffffffff19166001179055610c409082610c8f57610fc7565b610c4657005b6000805160206114c6833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b6000805160206114c6833981519152805460ff60401b191668010000000000000000179055610fc7565b60405163f92ee8a960e01b8152600490fd5b90501538610c0c565b303b159150610c04565b839150610bfa565b6024359065ffffffffffff821682036101c957565b6044359065ffffffffffff821682036101c957565b346101c95760603660031901126101c957600435610d2d816101ce565b610d35610ce6565b90610d3e610cfb565b90610d4761117e565b65ffffffffffff8084169083161115610db35761001992610d909160018060a01b0316600052600160205260406000209065ffffffffffff1665ffffffffffff19825416179055565b600260205260406000209065ffffffffffff1665ffffffffffff19825416179055565b60405162461bcd60e51b81526020600482015260116024820152703bb937b733903ab73a34b617b0b33a32b960791b6044820152606490fd5b346101c95760003660031901126101c957604051631aab3f0d60e11b8152306004820152600060248201526020816044817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156106185761037891600091610e6c575b506040519081529081906020820190565b610e85915060203d6020116108b6576108a781836106fe565b38610e5b565b346101c95760a03660031901126101c957610ea76004356101ce565b610eb26024356101ce565b6084356001600160401b0381116101c957610ed19036906004016101df565b505060405163f23a6e6160e01b8152602090f35b15610eec57565b60405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b6044820152606490fd5b634e487b7160e01b600052603260045260246000fd5b9190811015610f4d5760051b0190565b610f27565b35610f5c816101ce565b90565b903590601e19813603018212156101c957018035906001600160401b0382116101c9576020019181360383136101c957565b90821015610f4d57610fa89160051b810190610f5f565b9091565b6040513d6000823e3d90fd5b908160209103126101c9575190565b600080546001600160a01b0319166001600160a01b039283169081178255909181907f0000000000000000000000000000000000000000000000000000000000000000167f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de8480a361103761117e565b8152600160205265ffffffffffff60408083209282199384815416905560026020522091825416179055565b3d1561108e573d906110748261071f565b9161108260405193846106fe565b82523d6000602084013e565b606090565b8061109b5750565b600080808093338219f1506110ae611063565b50565b60018060a01b03807f0000000000000000000000000000000000000000000000000000000000000000163314908115611131575b50156110ed57565b606460405162461bcd60e51b815260206004820152602060248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152fd5b9050600054163314386110e5565b600091829182602083519301915af1611156611063565b901561115f5750565b602081519101fd5b916000928392602083519301915af1611156611063565b6000546001600160a01b0316331480156111cc575b1561119a57565b60405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b6044820152606490fd5b50303314611193565b909190156112095760ff6001915b65ffffffffffff60d01b9060d01b169265ffffffffffff60a01b9060a01b169116171790565b60ff6000916111e3565b90813b15611296576000805160206114a683398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a280511561127b576110ae916113a0565b50503461128457565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b81519190604183036112e8576112e192506020820151906060604084015193015160001a906113be565b9192909190565b505060009160029190565b600411156112fd57565b634e487b7160e01b600052602160045260246000fd5b61131c816112f3565b80611325575050565b61132e816112f3565b600181036113485760405163f645eedf60e01b8152600490fd5b611351816112f3565b600281036113725760405163fce698f760e01b815260048101839052602490fd5b8061137e6003926112f3565b146113865750565b6040516335e2f38360e21b81526004810191909152602490fd5b600080610f5c93602081519101845af46113b8611063565b91611442565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841161143657926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa156106185780516001600160a01b0381161561142d57918190565b50809160019190565b50505060009160039190565b90611469575080511561145757805190602001fd5b604051630a12f52160e11b8152600490fd5b8151158061149c575b61147a575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561147256fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212205c6fc8fe0e303fa9d9e442107b50930681da8e8d456268a3e2ea0f5b24b28e2464736f6c63430008190033",
}

// TestExpiryAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use TestExpiryAccountMetaData.ABI instead.
var TestExpiryAccountABI = TestExpiryAccountMetaData.ABI

// TestExpiryAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestExpiryAccountMetaData.Bin instead.
var TestExpiryAccountBin = TestExpiryAccountMetaData.Bin

// DeployTestExpiryAccount deploys a new Ethereum contract, binding an instance of TestExpiryAccount to it.
func DeployTestExpiryAccount(auth *bind.TransactOpts, backend bind.ContractBackend, anEntryPoint common.Address) (common.Address, *types.Transaction, *TestExpiryAccount, error) {
	parsed, err := TestExpiryAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestExpiryAccountBin), backend, anEntryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestExpiryAccount{TestExpiryAccountCaller: TestExpiryAccountCaller{contract: contract}, TestExpiryAccountTransactor: TestExpiryAccountTransactor{contract: contract}, TestExpiryAccountFilterer: TestExpiryAccountFilterer{contract: contract}}, nil
}

// TestExpiryAccount is an auto generated Go binding around an Ethereum contract.
type TestExpiryAccount struct {
	TestExpiryAccountCaller     // Read-only binding to the contract
	TestExpiryAccountTransactor // Write-only binding to the contract
	TestExpiryAccountFilterer   // Log filterer for contract events
}

// TestExpiryAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestExpiryAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExpiryAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestExpiryAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExpiryAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestExpiryAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExpiryAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestExpiryAccountSession struct {
	Contract     *TestExpiryAccount // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestExpiryAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestExpiryAccountCallerSession struct {
	Contract *TestExpiryAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TestExpiryAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestExpiryAccountTransactorSession struct {
	Contract     *TestExpiryAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TestExpiryAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestExpiryAccountRaw struct {
	Contract *TestExpiryAccount // Generic contract binding to access the raw methods on
}

// TestExpiryAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestExpiryAccountCallerRaw struct {
	Contract *TestExpiryAccountCaller // Generic read-only contract binding to access the raw methods on
}

// TestExpiryAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestExpiryAccountTransactorRaw struct {
	Contract *TestExpiryAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestExpiryAccount creates a new instance of TestExpiryAccount, bound to a specific deployed contract.
func NewTestExpiryAccount(address common.Address, backend bind.ContractBackend) (*TestExpiryAccount, error) {
	contract, err := bindTestExpiryAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestExpiryAccount{TestExpiryAccountCaller: TestExpiryAccountCaller{contract: contract}, TestExpiryAccountTransactor: TestExpiryAccountTransactor{contract: contract}, TestExpiryAccountFilterer: TestExpiryAccountFilterer{contract: contract}}, nil
}

// NewTestExpiryAccountCaller creates a new read-only instance of TestExpiryAccount, bound to a specific deployed contract.
func NewTestExpiryAccountCaller(address common.Address, caller bind.ContractCaller) (*TestExpiryAccountCaller, error) {
	contract, err := bindTestExpiryAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestExpiryAccountCaller{contract: contract}, nil
}

// NewTestExpiryAccountTransactor creates a new write-only instance of TestExpiryAccount, bound to a specific deployed contract.
func NewTestExpiryAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*TestExpiryAccountTransactor, error) {
	contract, err := bindTestExpiryAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestExpiryAccountTransactor{contract: contract}, nil
}

// NewTestExpiryAccountFilterer creates a new log filterer instance of TestExpiryAccount, bound to a specific deployed contract.
func NewTestExpiryAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*TestExpiryAccountFilterer, error) {
	contract, err := bindTestExpiryAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestExpiryAccountFilterer{contract: contract}, nil
}

// bindTestExpiryAccount binds a generic wrapper to an already deployed contract.
func bindTestExpiryAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestExpiryAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestExpiryAccount *TestExpiryAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestExpiryAccount.Contract.TestExpiryAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestExpiryAccount *TestExpiryAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.TestExpiryAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestExpiryAccount *TestExpiryAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.TestExpiryAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestExpiryAccount *TestExpiryAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestExpiryAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestExpiryAccount *TestExpiryAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestExpiryAccount *TestExpiryAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TestExpiryAccount *TestExpiryAccountCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestExpiryAccount.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TestExpiryAccount *TestExpiryAccountSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TestExpiryAccount.Contract.UPGRADEINTERFACEVERSION(&_TestExpiryAccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TestExpiryAccount *TestExpiryAccountCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TestExpiryAccount.Contract.UPGRADEINTERFACEVERSION(&_TestExpiryAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestExpiryAccount *TestExpiryAccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestExpiryAccount.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestExpiryAccount *TestExpiryAccountSession) EntryPoint() (common.Address, error) {
	return _TestExpiryAccount.Contract.EntryPoint(&_TestExpiryAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestExpiryAccount *TestExpiryAccountCallerSession) EntryPoint() (common.Address, error) {
	return _TestExpiryAccount.Contract.EntryPoint(&_TestExpiryAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestExpiryAccount *TestExpiryAccountCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestExpiryAccount.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestExpiryAccount *TestExpiryAccountSession) GetDeposit() (*big.Int, error) {
	return _TestExpiryAccount.Contract.GetDeposit(&_TestExpiryAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestExpiryAccount *TestExpiryAccountCallerSession) GetDeposit() (*big.Int, error) {
	return _TestExpiryAccount.Contract.GetDeposit(&_TestExpiryAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_TestExpiryAccount *TestExpiryAccountCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestExpiryAccount.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_TestExpiryAccount *TestExpiryAccountSession) GetNonce() (*big.Int, error) {
	return _TestExpiryAccount.Contract.GetNonce(&_TestExpiryAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_TestExpiryAccount *TestExpiryAccountCallerSession) GetNonce() (*big.Int, error) {
	return _TestExpiryAccount.Contract.GetNonce(&_TestExpiryAccount.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_TestExpiryAccount *TestExpiryAccountCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TestExpiryAccount.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_TestExpiryAccount *TestExpiryAccountSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _TestExpiryAccount.Contract.OnERC1155BatchReceived(&_TestExpiryAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_TestExpiryAccount *TestExpiryAccountCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _TestExpiryAccount.Contract.OnERC1155BatchReceived(&_TestExpiryAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_TestExpiryAccount *TestExpiryAccountCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TestExpiryAccount.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_TestExpiryAccount *TestExpiryAccountSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _TestExpiryAccount.Contract.OnERC1155Received(&_TestExpiryAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_TestExpiryAccount *TestExpiryAccountCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _TestExpiryAccount.Contract.OnERC1155Received(&_TestExpiryAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TestExpiryAccount *TestExpiryAccountCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TestExpiryAccount.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TestExpiryAccount *TestExpiryAccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _TestExpiryAccount.Contract.OnERC721Received(&_TestExpiryAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TestExpiryAccount *TestExpiryAccountCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _TestExpiryAccount.Contract.OnERC721Received(&_TestExpiryAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestExpiryAccount *TestExpiryAccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestExpiryAccount.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestExpiryAccount *TestExpiryAccountSession) Owner() (common.Address, error) {
	return _TestExpiryAccount.Contract.Owner(&_TestExpiryAccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestExpiryAccount *TestExpiryAccountCallerSession) Owner() (common.Address, error) {
	return _TestExpiryAccount.Contract.Owner(&_TestExpiryAccount.CallOpts)
}

// OwnerAfter is a free data retrieval call binding the contract method 0x3fb5a7a1.
//
// Solidity: function ownerAfter(address ) view returns(uint48)
func (_TestExpiryAccount *TestExpiryAccountCaller) OwnerAfter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestExpiryAccount.contract.Call(opts, &out, "ownerAfter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerAfter is a free data retrieval call binding the contract method 0x3fb5a7a1.
//
// Solidity: function ownerAfter(address ) view returns(uint48)
func (_TestExpiryAccount *TestExpiryAccountSession) OwnerAfter(arg0 common.Address) (*big.Int, error) {
	return _TestExpiryAccount.Contract.OwnerAfter(&_TestExpiryAccount.CallOpts, arg0)
}

// OwnerAfter is a free data retrieval call binding the contract method 0x3fb5a7a1.
//
// Solidity: function ownerAfter(address ) view returns(uint48)
func (_TestExpiryAccount *TestExpiryAccountCallerSession) OwnerAfter(arg0 common.Address) (*big.Int, error) {
	return _TestExpiryAccount.Contract.OwnerAfter(&_TestExpiryAccount.CallOpts, arg0)
}

// OwnerUntil is a free data retrieval call binding the contract method 0x3e476951.
//
// Solidity: function ownerUntil(address ) view returns(uint48)
func (_TestExpiryAccount *TestExpiryAccountCaller) OwnerUntil(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestExpiryAccount.contract.Call(opts, &out, "ownerUntil", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerUntil is a free data retrieval call binding the contract method 0x3e476951.
//
// Solidity: function ownerUntil(address ) view returns(uint48)
func (_TestExpiryAccount *TestExpiryAccountSession) OwnerUntil(arg0 common.Address) (*big.Int, error) {
	return _TestExpiryAccount.Contract.OwnerUntil(&_TestExpiryAccount.CallOpts, arg0)
}

// OwnerUntil is a free data retrieval call binding the contract method 0x3e476951.
//
// Solidity: function ownerUntil(address ) view returns(uint48)
func (_TestExpiryAccount *TestExpiryAccountCallerSession) OwnerUntil(arg0 common.Address) (*big.Int, error) {
	return _TestExpiryAccount.Contract.OwnerUntil(&_TestExpiryAccount.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TestExpiryAccount *TestExpiryAccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TestExpiryAccount.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TestExpiryAccount *TestExpiryAccountSession) ProxiableUUID() ([32]byte, error) {
	return _TestExpiryAccount.Contract.ProxiableUUID(&_TestExpiryAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TestExpiryAccount *TestExpiryAccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TestExpiryAccount.Contract.ProxiableUUID(&_TestExpiryAccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestExpiryAccount *TestExpiryAccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TestExpiryAccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestExpiryAccount *TestExpiryAccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestExpiryAccount.Contract.SupportsInterface(&_TestExpiryAccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestExpiryAccount *TestExpiryAccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestExpiryAccount.Contract.SupportsInterface(&_TestExpiryAccount.CallOpts, interfaceId)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_TestExpiryAccount *TestExpiryAccountTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExpiryAccount.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_TestExpiryAccount *TestExpiryAccountSession) AddDeposit() (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.AddDeposit(&_TestExpiryAccount.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_TestExpiryAccount *TestExpiryAccountTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.AddDeposit(&_TestExpiryAccount.TransactOpts)
}

// AddTemporaryOwner is a paid mutator transaction binding the contract method 0xcf6dca55.
//
// Solidity: function addTemporaryOwner(address owner, uint48 _after, uint48 _until) returns()
func (_TestExpiryAccount *TestExpiryAccountTransactor) AddTemporaryOwner(opts *bind.TransactOpts, owner common.Address, _after *big.Int, _until *big.Int) (*types.Transaction, error) {
	return _TestExpiryAccount.contract.Transact(opts, "addTemporaryOwner", owner, _after, _until)
}

// AddTemporaryOwner is a paid mutator transaction binding the contract method 0xcf6dca55.
//
// Solidity: function addTemporaryOwner(address owner, uint48 _after, uint48 _until) returns()
func (_TestExpiryAccount *TestExpiryAccountSession) AddTemporaryOwner(owner common.Address, _after *big.Int, _until *big.Int) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.AddTemporaryOwner(&_TestExpiryAccount.TransactOpts, owner, _after, _until)
}

// AddTemporaryOwner is a paid mutator transaction binding the contract method 0xcf6dca55.
//
// Solidity: function addTemporaryOwner(address owner, uint48 _after, uint48 _until) returns()
func (_TestExpiryAccount *TestExpiryAccountTransactorSession) AddTemporaryOwner(owner common.Address, _after *big.Int, _until *big.Int) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.AddTemporaryOwner(&_TestExpiryAccount.TransactOpts, owner, _after, _until)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_TestExpiryAccount *TestExpiryAccountTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _TestExpiryAccount.contract.Transact(opts, "execute", dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_TestExpiryAccount *TestExpiryAccountSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.Execute(&_TestExpiryAccount.TransactOpts, dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_TestExpiryAccount *TestExpiryAccountTransactorSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.Execute(&_TestExpiryAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_TestExpiryAccount *TestExpiryAccountTransactor) ExecuteBatch(opts *bind.TransactOpts, dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _TestExpiryAccount.contract.Transact(opts, "executeBatch", dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_TestExpiryAccount *TestExpiryAccountSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.ExecuteBatch(&_TestExpiryAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_TestExpiryAccount *TestExpiryAccountTransactorSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.ExecuteBatch(&_TestExpiryAccount.TransactOpts, dest, value, arg2)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_TestExpiryAccount *TestExpiryAccountTransactor) Initialize(opts *bind.TransactOpts, anOwner common.Address) (*types.Transaction, error) {
	return _TestExpiryAccount.contract.Transact(opts, "initialize", anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_TestExpiryAccount *TestExpiryAccountSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.Initialize(&_TestExpiryAccount.TransactOpts, anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_TestExpiryAccount *TestExpiryAccountTransactorSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.Initialize(&_TestExpiryAccount.TransactOpts, anOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TestExpiryAccount *TestExpiryAccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TestExpiryAccount.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TestExpiryAccount *TestExpiryAccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.UpgradeToAndCall(&_TestExpiryAccount.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TestExpiryAccount *TestExpiryAccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.UpgradeToAndCall(&_TestExpiryAccount.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestExpiryAccount *TestExpiryAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestExpiryAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestExpiryAccount *TestExpiryAccountSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.ValidateUserOp(&_TestExpiryAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestExpiryAccount *TestExpiryAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.ValidateUserOp(&_TestExpiryAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_TestExpiryAccount *TestExpiryAccountTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestExpiryAccount.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_TestExpiryAccount *TestExpiryAccountSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.WithdrawDepositTo(&_TestExpiryAccount.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_TestExpiryAccount *TestExpiryAccountTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.WithdrawDepositTo(&_TestExpiryAccount.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestExpiryAccount *TestExpiryAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExpiryAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestExpiryAccount *TestExpiryAccountSession) Receive() (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.Receive(&_TestExpiryAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestExpiryAccount *TestExpiryAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _TestExpiryAccount.Contract.Receive(&_TestExpiryAccount.TransactOpts)
}

// TestExpiryAccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TestExpiryAccount contract.
type TestExpiryAccountInitializedIterator struct {
	Event *TestExpiryAccountInitialized // Event containing the contract specifics and raw log

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
func (it *TestExpiryAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestExpiryAccountInitialized)
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
		it.Event = new(TestExpiryAccountInitialized)
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
func (it *TestExpiryAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestExpiryAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestExpiryAccountInitialized represents a Initialized event raised by the TestExpiryAccount contract.
type TestExpiryAccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TestExpiryAccount *TestExpiryAccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*TestExpiryAccountInitializedIterator, error) {

	logs, sub, err := _TestExpiryAccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TestExpiryAccountInitializedIterator{contract: _TestExpiryAccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TestExpiryAccount *TestExpiryAccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TestExpiryAccountInitialized) (event.Subscription, error) {

	logs, sub, err := _TestExpiryAccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestExpiryAccountInitialized)
				if err := _TestExpiryAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TestExpiryAccount *TestExpiryAccountFilterer) ParseInitialized(log types.Log) (*TestExpiryAccountInitialized, error) {
	event := new(TestExpiryAccountInitialized)
	if err := _TestExpiryAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestExpiryAccountSimpleAccountInitializedIterator is returned from FilterSimpleAccountInitialized and is used to iterate over the raw logs and unpacked data for SimpleAccountInitialized events raised by the TestExpiryAccount contract.
type TestExpiryAccountSimpleAccountInitializedIterator struct {
	Event *TestExpiryAccountSimpleAccountInitialized // Event containing the contract specifics and raw log

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
func (it *TestExpiryAccountSimpleAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestExpiryAccountSimpleAccountInitialized)
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
		it.Event = new(TestExpiryAccountSimpleAccountInitialized)
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
func (it *TestExpiryAccountSimpleAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestExpiryAccountSimpleAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestExpiryAccountSimpleAccountInitialized represents a SimpleAccountInitialized event raised by the TestExpiryAccount contract.
type TestExpiryAccountSimpleAccountInitialized struct {
	EntryPoint common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSimpleAccountInitialized is a free log retrieval operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_TestExpiryAccount *TestExpiryAccountFilterer) FilterSimpleAccountInitialized(opts *bind.FilterOpts, entryPoint []common.Address, owner []common.Address) (*TestExpiryAccountSimpleAccountInitializedIterator, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TestExpiryAccount.contract.FilterLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TestExpiryAccountSimpleAccountInitializedIterator{contract: _TestExpiryAccount.contract, event: "SimpleAccountInitialized", logs: logs, sub: sub}, nil
}

// WatchSimpleAccountInitialized is a free log subscription operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_TestExpiryAccount *TestExpiryAccountFilterer) WatchSimpleAccountInitialized(opts *bind.WatchOpts, sink chan<- *TestExpiryAccountSimpleAccountInitialized, entryPoint []common.Address, owner []common.Address) (event.Subscription, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TestExpiryAccount.contract.WatchLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestExpiryAccountSimpleAccountInitialized)
				if err := _TestExpiryAccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
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
func (_TestExpiryAccount *TestExpiryAccountFilterer) ParseSimpleAccountInitialized(log types.Log) (*TestExpiryAccountSimpleAccountInitialized, error) {
	event := new(TestExpiryAccountSimpleAccountInitialized)
	if err := _TestExpiryAccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestExpiryAccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TestExpiryAccount contract.
type TestExpiryAccountUpgradedIterator struct {
	Event *TestExpiryAccountUpgraded // Event containing the contract specifics and raw log

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
func (it *TestExpiryAccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestExpiryAccountUpgraded)
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
		it.Event = new(TestExpiryAccountUpgraded)
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
func (it *TestExpiryAccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestExpiryAccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestExpiryAccountUpgraded represents a Upgraded event raised by the TestExpiryAccount contract.
type TestExpiryAccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TestExpiryAccount *TestExpiryAccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TestExpiryAccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TestExpiryAccount.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TestExpiryAccountUpgradedIterator{contract: _TestExpiryAccount.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TestExpiryAccount *TestExpiryAccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TestExpiryAccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TestExpiryAccount.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestExpiryAccountUpgraded)
				if err := _TestExpiryAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TestExpiryAccount *TestExpiryAccountFilterer) ParseUpgraded(log types.Log) (*TestExpiryAccountUpgraded, error) {
	event := new(TestExpiryAccountUpgraded)
	if err := _TestExpiryAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
