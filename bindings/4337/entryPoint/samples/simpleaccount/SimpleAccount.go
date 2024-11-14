// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simpleaccount

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

// SimpleAccountMetaData contains all meta data concerning the SimpleAccount contract.
var SimpleAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SimpleAccountInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c034610142576001600160401b0390601f61143a38819003918201601f1916830191848311848410176101475780849260209460405283398101031261014257516001600160a01b0381168103610142573060805260a0527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166101305780808316036100eb575b6040516112dc908161015e82396080518181816106a801526107e1015260a05181818161026c0152818161049301528181610535015281816108f501528181610a3401528181610bfd01528181610e4f0152610f8a0152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1388080610092565b60405163f92ee8a960e01b8152600490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c806301ffc9a71461012b578063150b7a021461012657806319822f7c1461012157806347e1da2a1461011c5780634a58db19146101175780634d44560d146101125780634f1ef2861461010d57806352d1902d146101085780638da5cb5b14610103578063ad3cb1cc146100fe578063b0d691fe146100f9578063b61d27f6146100f4578063bc197c81146100ef578063c399ec88146100ea578063c4d66de8146100e5578063d087d288146100e05763f23a6e610361000e57610c68565b610bca565b610a95565b610a08565b610979565b610924565b6108df565b610862565b610839565b6107ce565b610655565b610504565b610484565b610358565b610236565b6101dc565b346101995760203660031901126101995760043563ffffffff60e01b811680910361019957602090630a85bd0160e11b8114908115610188575b8115610177575b506040519015158152f35b6301ffc9a760e01b1490503861016c565b630271189760e51b81149150610165565b600080fd5b6001600160a01b0381160361019957565b9181601f84011215610199578235916001600160401b038311610199576020838186019501011161019957565b34610199576080366003190112610199576101f860043561019e565b61020360243561019e565b6064356001600160401b038111610199576102229036906004016101af565b5050604051630a85bd0160e11b8152602090f35b346101995760031960603682011261019957600435906001600160401b03821161019957610120908236030112610199576044357f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036102e3576102ad6102c59260243590600401610da4565b90806102c9575b506040519081529081906020820190565b0390f35b600080808093338219f1506102dc610e15565b50386102b4565b60405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e74000000006044820152606490fd5b9181601f84011215610199578235916001600160401b038311610199576020808501948460051b01011161019957565b34610199576060366003190112610199576001600160401b0360043581811161019957610389903690600401610328565b602492919235828111610199576103a4903690600401610328565b92604435908111610199576103bd903690600401610328565b9390916103c8610e45565b84841480610473575b6103da90610cc2565b8161042757505060005b8281106103ed57005b80610421610406610401600194878a610d1a565b610d2f565b61041b610414848988610d6e565b369161061e565b90610ed3565b016103e4565b91909460009493945b85811061043957005b8061046d61044d6104016001948a87610d1a565b610458838b89610d1a565b35610467610414858b8a610d6e565b91610efb565b01610430565b508115806103d157508185146103d1565b600080600319360112610501577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681813b156105015760405163b760faf960e01b8152306004820152918290602490829034905af180156104fc576104f0575080f35b6104f9906105af565b80f35b610d89565b80fd5b346101995760006040366003190112610501576004356105238161019e565b61052b610f12565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b156105955760449083604051958694859363040b850f60e31b855216600484015260243560248401525af180156104fc576104f0575080f35b8280fd5b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116105c257604052565b610599565b604081019081106001600160401b038211176105c257604052565b90601f801991011681019081106001600160401b038211176105c257604052565b6001600160401b0381116105c257601f01601f191660200190565b92919261062a82610603565b9161063860405193846105e2565b829481845281830111610199578281602093846000960137010152565b604036600319011261019957600480359061066f8261019e565b6024356001600160401b03811161019957366023820112156101995761069e903690602481850135910161061e565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081163081149081156107b2575b506107a15790602083926106e6610f12565b6040516352d1902d60e01b8152938491829088165afa60009281610770575b50610733575050604051634c9c8ce360e01b81526001600160a01b0390921690820190815281906020010390fd5b83836000805160206112678339815191528403610754576100198383610fd1565b604051632a87526960e21b815290810184815281906020010390fd5b61079391935060203d60201161079a575b61078b81836105e2565b810190610d95565b9138610705565b503d610781565b60405163703e46dd60e11b81528390fd5b90508160008051602061126783398151915254161415386106d4565b34610199576000366003190112610199577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036108275760206040516000805160206112678339815191528152f35b60405163703e46dd60e11b8152600490fd5b34610199576000366003190112610199576000546040516001600160a01b039091168152602090f35b34610199576000366003190112610199576040805190610881826105c7565b60058252602090640352e302e360dc1b6020840152604051916020835283519182602085015260005b8381106108cc5784604081866000838284010152601f80199101168101030190f35b85810183015185820183015282016108aa565b34610199576000366003190112610199576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610199576060366003190112610199576004356109418161019e565b604435906001600160401b0382116101995761096f6109676100199336906004016101af565b610414610e45565b9060243590610efb565b346101995760a03660031901126101995761099560043561019e565b6109a060243561019e565b6001600160401b03604435818111610199576109c0903690600401610328565b5050606435818111610199576109da903690600401610328565b5050608435908111610199576109f49036906004016101af565b505060405163bc197c8160e01b8152602090f35b34610199576000366003190112610199576040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156104fc57602091600091610a78575b50604051908152f35b610a8f9150823d841161079a5761078b81836105e2565b38610a6f565b3461019957602036600319011261019957600435610ab28161019e565b60008051602061128783398151915254906001600160401b0360ff8360401c1615921680159081610bc2575b6001149081610bb8575b159081610baf575b50610b9d57600080516020611287833981519152805467ffffffffffffffff19166001179055610b249082610b7357610f69565b610b2a57005b600080516020611287833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b600080516020611287833981519152805460ff60401b191668010000000000000000179055610f69565b60405163f92ee8a960e01b8152600490fd5b90501538610af0565b303b159150610ae8565b839150610ade565b3461019957600036600319011261019957604051631aab3f0d60e11b8152306004820152600060248201526020816044817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156104fc576102c591600091610c4957506040519081529081906020820190565b610c62915060203d60201161079a5761078b81836105e2565b386102b4565b346101995760a036600319011261019957610c8460043561019e565b610c8f60243561019e565b6084356001600160401b03811161019957610cae9036906004016101af565b505060405163f23a6e6160e01b8152602090f35b15610cc957565b60405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b6044820152606490fd5b634e487b7160e01b600052603260045260246000fd5b9190811015610d2a5760051b0190565b610d04565b35610d398161019e565b90565b903590601e198136030182121561019957018035906001600160401b0382116101995760200191813603831361019957565b90821015610d2a57610d859160051b810190610d3c565b9091565b6040513d6000823e3d90fd5b90816020910312610199575190565b907f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52603c600020610e05610dfc60018060a01b0392610df6610414856000541696610100810190610d3c565b90611078565b909291926110d4565b1603610e1057600090565b600190565b3d15610e40573d90610e2682610603565b91610e3460405193846105e2565b82523d6000602084013e565b606090565b60018060a01b03807f0000000000000000000000000000000000000000000000000000000000000000163314908115610ec5575b5015610e8157565b606460405162461bcd60e51b815260206004820152602060248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152fd5b905060005416331438610e79565b600091829182602083519301915af1610eea610e15565b9015610ef35750565b602081519101fd5b916000928392602083519301915af1610eea610e15565b6000546001600160a01b031633148015610f60575b15610f2e57565b60405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b6044820152606490fd5b50303314610f27565b600080546001600160a01b0319166001600160a01b039283169081178255917f000000000000000000000000000000000000000000000000000000000000000016907f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de9080a3565b90813b156110575760008051602061126783398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a280511561103c5761103991611161565b50565b50503461104557565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b81519190604183036110a9576110a292506020820151906060604084015193015160001a9061117f565b9192909190565b505060009160029190565b600411156110be57565b634e487b7160e01b600052602160045260246000fd5b6110dd816110b4565b806110e6575050565b6110ef816110b4565b600181036111095760405163f645eedf60e01b8152600490fd5b611112816110b4565b600281036111335760405163fce698f760e01b815260048101839052602490fd5b8061113f6003926110b4565b146111475750565b6040516335e2f38360e21b81526004810191909152602490fd5b600080610d3993602081519101845af4611179610e15565b91611203565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084116111f757926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa156104fc5780516001600160a01b038116156111ee57918190565b50809160019190565b50505060009160039190565b9061122a575080511561121857805190602001fd5b604051630a12f52160e11b8152600490fd5b8151158061125d575b61123b575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561123356fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122082f2fa65e2876854fdf00961c65919efe811a0639d34590117dcfb3e334fddc264736f6c63430008190033",
}

// SimpleAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleAccountMetaData.ABI instead.
var SimpleAccountABI = SimpleAccountMetaData.ABI

// SimpleAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleAccountMetaData.Bin instead.
var SimpleAccountBin = SimpleAccountMetaData.Bin

// DeploySimpleAccount deploys a new Ethereum contract, binding an instance of SimpleAccount to it.
func DeploySimpleAccount(auth *bind.TransactOpts, backend bind.ContractBackend, anEntryPoint common.Address) (common.Address, *types.Transaction, *SimpleAccount, error) {
	parsed, err := SimpleAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleAccountBin), backend, anEntryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleAccount{SimpleAccountCaller: SimpleAccountCaller{contract: contract}, SimpleAccountTransactor: SimpleAccountTransactor{contract: contract}, SimpleAccountFilterer: SimpleAccountFilterer{contract: contract}}, nil
}

// SimpleAccount is an auto generated Go binding around an Ethereum contract.
type SimpleAccount struct {
	SimpleAccountCaller     // Read-only binding to the contract
	SimpleAccountTransactor // Write-only binding to the contract
	SimpleAccountFilterer   // Log filterer for contract events
}

// SimpleAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleAccountSession struct {
	Contract     *SimpleAccount    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleAccountCallerSession struct {
	Contract *SimpleAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleAccountTransactorSession struct {
	Contract     *SimpleAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleAccountRaw struct {
	Contract *SimpleAccount // Generic contract binding to access the raw methods on
}

// SimpleAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleAccountCallerRaw struct {
	Contract *SimpleAccountCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleAccountTransactorRaw struct {
	Contract *SimpleAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleAccount creates a new instance of SimpleAccount, bound to a specific deployed contract.
func NewSimpleAccount(address common.Address, backend bind.ContractBackend) (*SimpleAccount, error) {
	contract, err := bindSimpleAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleAccount{SimpleAccountCaller: SimpleAccountCaller{contract: contract}, SimpleAccountTransactor: SimpleAccountTransactor{contract: contract}, SimpleAccountFilterer: SimpleAccountFilterer{contract: contract}}, nil
}

// NewSimpleAccountCaller creates a new read-only instance of SimpleAccount, bound to a specific deployed contract.
func NewSimpleAccountCaller(address common.Address, caller bind.ContractCaller) (*SimpleAccountCaller, error) {
	contract, err := bindSimpleAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAccountCaller{contract: contract}, nil
}

// NewSimpleAccountTransactor creates a new write-only instance of SimpleAccount, bound to a specific deployed contract.
func NewSimpleAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleAccountTransactor, error) {
	contract, err := bindSimpleAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAccountTransactor{contract: contract}, nil
}

// NewSimpleAccountFilterer creates a new log filterer instance of SimpleAccount, bound to a specific deployed contract.
func NewSimpleAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleAccountFilterer, error) {
	contract, err := bindSimpleAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleAccountFilterer{contract: contract}, nil
}

// bindSimpleAccount binds a generic wrapper to an already deployed contract.
func bindSimpleAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAccount *SimpleAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleAccount.Contract.SimpleAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAccount *SimpleAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAccount.Contract.SimpleAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAccount *SimpleAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleAccount.Contract.SimpleAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAccount *SimpleAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAccount *SimpleAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAccount *SimpleAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleAccount.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SimpleAccount *SimpleAccountCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SimpleAccount.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SimpleAccount *SimpleAccountSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SimpleAccount.Contract.UPGRADEINTERFACEVERSION(&_SimpleAccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SimpleAccount *SimpleAccountCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SimpleAccount.Contract.UPGRADEINTERFACEVERSION(&_SimpleAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_SimpleAccount *SimpleAccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleAccount.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_SimpleAccount *SimpleAccountSession) EntryPoint() (common.Address, error) {
	return _SimpleAccount.Contract.EntryPoint(&_SimpleAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_SimpleAccount *SimpleAccountCallerSession) EntryPoint() (common.Address, error) {
	return _SimpleAccount.Contract.EntryPoint(&_SimpleAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_SimpleAccount *SimpleAccountCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleAccount.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_SimpleAccount *SimpleAccountSession) GetDeposit() (*big.Int, error) {
	return _SimpleAccount.Contract.GetDeposit(&_SimpleAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_SimpleAccount *SimpleAccountCallerSession) GetDeposit() (*big.Int, error) {
	return _SimpleAccount.Contract.GetDeposit(&_SimpleAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_SimpleAccount *SimpleAccountCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleAccount.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_SimpleAccount *SimpleAccountSession) GetNonce() (*big.Int, error) {
	return _SimpleAccount.Contract.GetNonce(&_SimpleAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_SimpleAccount *SimpleAccountCallerSession) GetNonce() (*big.Int, error) {
	return _SimpleAccount.Contract.GetNonce(&_SimpleAccount.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_SimpleAccount *SimpleAccountCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _SimpleAccount.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_SimpleAccount *SimpleAccountSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _SimpleAccount.Contract.OnERC1155BatchReceived(&_SimpleAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_SimpleAccount *SimpleAccountCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _SimpleAccount.Contract.OnERC1155BatchReceived(&_SimpleAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_SimpleAccount *SimpleAccountCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _SimpleAccount.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_SimpleAccount *SimpleAccountSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _SimpleAccount.Contract.OnERC1155Received(&_SimpleAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_SimpleAccount *SimpleAccountCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _SimpleAccount.Contract.OnERC1155Received(&_SimpleAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_SimpleAccount *SimpleAccountCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _SimpleAccount.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_SimpleAccount *SimpleAccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _SimpleAccount.Contract.OnERC721Received(&_SimpleAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_SimpleAccount *SimpleAccountCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _SimpleAccount.Contract.OnERC721Received(&_SimpleAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SimpleAccount *SimpleAccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleAccount.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SimpleAccount *SimpleAccountSession) Owner() (common.Address, error) {
	return _SimpleAccount.Contract.Owner(&_SimpleAccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SimpleAccount *SimpleAccountCallerSession) Owner() (common.Address, error) {
	return _SimpleAccount.Contract.Owner(&_SimpleAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SimpleAccount *SimpleAccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SimpleAccount.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SimpleAccount *SimpleAccountSession) ProxiableUUID() ([32]byte, error) {
	return _SimpleAccount.Contract.ProxiableUUID(&_SimpleAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SimpleAccount *SimpleAccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SimpleAccount.Contract.ProxiableUUID(&_SimpleAccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SimpleAccount *SimpleAccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SimpleAccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SimpleAccount *SimpleAccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SimpleAccount.Contract.SupportsInterface(&_SimpleAccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SimpleAccount *SimpleAccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SimpleAccount.Contract.SupportsInterface(&_SimpleAccount.CallOpts, interfaceId)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_SimpleAccount *SimpleAccountTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAccount.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_SimpleAccount *SimpleAccountSession) AddDeposit() (*types.Transaction, error) {
	return _SimpleAccount.Contract.AddDeposit(&_SimpleAccount.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_SimpleAccount *SimpleAccountTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _SimpleAccount.Contract.AddDeposit(&_SimpleAccount.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_SimpleAccount *SimpleAccountTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SimpleAccount.contract.Transact(opts, "execute", dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_SimpleAccount *SimpleAccountSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SimpleAccount.Contract.Execute(&_SimpleAccount.TransactOpts, dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_SimpleAccount *SimpleAccountTransactorSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SimpleAccount.Contract.Execute(&_SimpleAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_SimpleAccount *SimpleAccountTransactor) ExecuteBatch(opts *bind.TransactOpts, dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _SimpleAccount.contract.Transact(opts, "executeBatch", dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_SimpleAccount *SimpleAccountSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _SimpleAccount.Contract.ExecuteBatch(&_SimpleAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_SimpleAccount *SimpleAccountTransactorSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _SimpleAccount.Contract.ExecuteBatch(&_SimpleAccount.TransactOpts, dest, value, arg2)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_SimpleAccount *SimpleAccountTransactor) Initialize(opts *bind.TransactOpts, anOwner common.Address) (*types.Transaction, error) {
	return _SimpleAccount.contract.Transact(opts, "initialize", anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_SimpleAccount *SimpleAccountSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _SimpleAccount.Contract.Initialize(&_SimpleAccount.TransactOpts, anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_SimpleAccount *SimpleAccountTransactorSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _SimpleAccount.Contract.Initialize(&_SimpleAccount.TransactOpts, anOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SimpleAccount *SimpleAccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SimpleAccount.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SimpleAccount *SimpleAccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SimpleAccount.Contract.UpgradeToAndCall(&_SimpleAccount.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SimpleAccount *SimpleAccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SimpleAccount.Contract.UpgradeToAndCall(&_SimpleAccount.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_SimpleAccount *SimpleAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _SimpleAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_SimpleAccount *SimpleAccountSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _SimpleAccount.Contract.ValidateUserOp(&_SimpleAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_SimpleAccount *SimpleAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _SimpleAccount.Contract.ValidateUserOp(&_SimpleAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_SimpleAccount *SimpleAccountTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleAccount.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_SimpleAccount *SimpleAccountSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleAccount.Contract.WithdrawDepositTo(&_SimpleAccount.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_SimpleAccount *SimpleAccountTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleAccount.Contract.WithdrawDepositTo(&_SimpleAccount.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleAccount *SimpleAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleAccount *SimpleAccountSession) Receive() (*types.Transaction, error) {
	return _SimpleAccount.Contract.Receive(&_SimpleAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleAccount *SimpleAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _SimpleAccount.Contract.Receive(&_SimpleAccount.TransactOpts)
}

// SimpleAccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SimpleAccount contract.
type SimpleAccountInitializedIterator struct {
	Event *SimpleAccountInitialized // Event containing the contract specifics and raw log

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
func (it *SimpleAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleAccountInitialized)
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
		it.Event = new(SimpleAccountInitialized)
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
func (it *SimpleAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleAccountInitialized represents a Initialized event raised by the SimpleAccount contract.
type SimpleAccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SimpleAccount *SimpleAccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*SimpleAccountInitializedIterator, error) {

	logs, sub, err := _SimpleAccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SimpleAccountInitializedIterator{contract: _SimpleAccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SimpleAccount *SimpleAccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SimpleAccountInitialized) (event.Subscription, error) {

	logs, sub, err := _SimpleAccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleAccountInitialized)
				if err := _SimpleAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SimpleAccount *SimpleAccountFilterer) ParseInitialized(log types.Log) (*SimpleAccountInitialized, error) {
	event := new(SimpleAccountInitialized)
	if err := _SimpleAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleAccountSimpleAccountInitializedIterator is returned from FilterSimpleAccountInitialized and is used to iterate over the raw logs and unpacked data for SimpleAccountInitialized events raised by the SimpleAccount contract.
type SimpleAccountSimpleAccountInitializedIterator struct {
	Event *SimpleAccountSimpleAccountInitialized // Event containing the contract specifics and raw log

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
func (it *SimpleAccountSimpleAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleAccountSimpleAccountInitialized)
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
		it.Event = new(SimpleAccountSimpleAccountInitialized)
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
func (it *SimpleAccountSimpleAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleAccountSimpleAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleAccountSimpleAccountInitialized represents a SimpleAccountInitialized event raised by the SimpleAccount contract.
type SimpleAccountSimpleAccountInitialized struct {
	EntryPoint common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSimpleAccountInitialized is a free log retrieval operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_SimpleAccount *SimpleAccountFilterer) FilterSimpleAccountInitialized(opts *bind.FilterOpts, entryPoint []common.Address, owner []common.Address) (*SimpleAccountSimpleAccountInitializedIterator, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SimpleAccount.contract.FilterLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SimpleAccountSimpleAccountInitializedIterator{contract: _SimpleAccount.contract, event: "SimpleAccountInitialized", logs: logs, sub: sub}, nil
}

// WatchSimpleAccountInitialized is a free log subscription operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_SimpleAccount *SimpleAccountFilterer) WatchSimpleAccountInitialized(opts *bind.WatchOpts, sink chan<- *SimpleAccountSimpleAccountInitialized, entryPoint []common.Address, owner []common.Address) (event.Subscription, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SimpleAccount.contract.WatchLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleAccountSimpleAccountInitialized)
				if err := _SimpleAccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
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
func (_SimpleAccount *SimpleAccountFilterer) ParseSimpleAccountInitialized(log types.Log) (*SimpleAccountSimpleAccountInitialized, error) {
	event := new(SimpleAccountSimpleAccountInitialized)
	if err := _SimpleAccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleAccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SimpleAccount contract.
type SimpleAccountUpgradedIterator struct {
	Event *SimpleAccountUpgraded // Event containing the contract specifics and raw log

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
func (it *SimpleAccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleAccountUpgraded)
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
		it.Event = new(SimpleAccountUpgraded)
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
func (it *SimpleAccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleAccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleAccountUpgraded represents a Upgraded event raised by the SimpleAccount contract.
type SimpleAccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SimpleAccount *SimpleAccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SimpleAccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SimpleAccount.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SimpleAccountUpgradedIterator{contract: _SimpleAccount.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SimpleAccount *SimpleAccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SimpleAccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SimpleAccount.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleAccountUpgraded)
				if err := _SimpleAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_SimpleAccount *SimpleAccountFilterer) ParseUpgraded(log types.Log) (*SimpleAccountUpgraded, error) {
	event := new(SimpleAccountUpgraded)
	if err := _SimpleAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
