// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simpleaccountfactory

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

// SimpleAccountFactoryMetaData contains all meta data concerning the SimpleAccountFactory contract.
var SimpleAccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"accountImplementation\",\"outputs\":[{\"internalType\":\"contractSimpleAccount\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"contractSimpleAccount\",\"name\":\"ret\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a03460c5576001600160401b0390601f611b6738819003918201601f19168301918483118484101760af5780849260209460405283398101031260c557516001600160a01b0381169081900360c5576040519161143a908184019081118482101760af57602092849261072d843981520301906000f0801560a35760805260405161066290816100cb8239608051818181609e0152818161021b01526102da0152f35b6040513d6000823e3d90fd5b634e487b7160e01b600052604160045260246000fd5b600080fdfe608080604052600436101561001357600080fd5b600090813560e01c90816311464fbe1461008a575080635fbfb9cf1461007057638cb84e181461004257600080fd5b3461006d57602061005b610055366100d1565b9061026b565b6040516001600160a01b039091168152f35b80fd5b503461006d57602061005b610084366100d1565b906101ae565b9050346100cd57816003193601126100cd577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b60409060031901126100f7576004356001600160a01b03811681036100f7579060243590565b600080fd5b6060810190811067ffffffffffffffff82111761011857604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff82111761011857604052565b60005b8381106101635750506000910152565b8181015183820152602001610153565b909160609260018060a01b03168252604060208301526101a28151809281604086015260208686019101610150565b601f01601f1916010190565b906101b9818361026b565b803b61025c575060405163189acdbd60e31b60208201526001600160a01b0392831660248083019190915281526101ef816100fc565b604051906102d38083019183831067ffffffffffffffff8411176101185783926102419261035a8539867f00000000000000000000000000000000000000000000000000000000000000001690610173565b03906000f58015610250571690565b6040513d6000823e3d90fd5b6001600160a01b031692915050565b600b60559161030e936102d3602090610334610340836040968751906102938387018361012e565b8582528282019561035a873961030060018060a01b039c8d92838c519163189acdbd60e31b88840152166024820152602481526102cf816100fc565b8b51928391878301957f00000000000000000000000000000000000000000000000000000000000000001686610173565b03601f19810183528261012e565b8951958693610325868601998a9251928391610150565b84019151809386840190610150565b0103808452018261012e565b5190208351938401528201523081520160ff815320169056fe60806040526102d38038038061001481610194565b92833981019060408183031261018f5780516001600160a01b03811680820361018f5760208381015190936001600160401b03821161018f570184601f8201121561018f5780519061006d610068836101cf565b610194565b9582875285838301011161018f57849060005b83811061017b57505060009186010152813b15610163577f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b03191682179055604051907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a28351156101455750600080848461012c96519101845af4903d1561013c573d61011c610068826101cf565b908152600081943d92013e6101ea565b505b6040516085908161024e8239f35b606092506101ea565b9250505034610154575061012e565b63b398979f60e01b8152600490fd5b60249060405190634c9c8ce360e01b82526004820152fd5b818101830151888201840152869201610080565b600080fd5b6040519190601f01601f191682016001600160401b038111838210176101b957604052565b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116101b957601f01601f191660200190565b9061021157508051156101ff57805190602001fd5b60405163d6bda27560e01b8152600490fd5b81511580610244575b610222575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561021a56fe60806040527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc54600090819081906001600160a01b0316368280378136915af43d82803e15604b573d90f35b3d90fdfea2646970667358221220fe8f61e3fa58bb8add11a833212ef9684476c5dbbd688eea5bcc0b0062abf4cc64736f6c63430008190033a2646970667358221220d5b1aa686f6b95c25d4d24bd2d79096ec3531126a4357c138c0b2707b0d5d5ae64736f6c6343000819003360c034610142576001600160401b0390601f61143a38819003918201601f1916830191848311848410176101475780849260209460405283398101031261014257516001600160a01b0381168103610142573060805260a0527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166101305780808316036100eb575b6040516112dc908161015e82396080518181816106a801526107e1015260a05181818161026c0152818161049301528181610535015281816108f501528181610a3401528181610bfd01528181610e4f0152610f8a0152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1388080610092565b60405163f92ee8a960e01b8152600490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c806301ffc9a71461012b578063150b7a021461012657806319822f7c1461012157806347e1da2a1461011c5780634a58db19146101175780634d44560d146101125780634f1ef2861461010d57806352d1902d146101085780638da5cb5b14610103578063ad3cb1cc146100fe578063b0d691fe146100f9578063b61d27f6146100f4578063bc197c81146100ef578063c399ec88146100ea578063c4d66de8146100e5578063d087d288146100e05763f23a6e610361000e57610c68565b610bca565b610a95565b610a08565b610979565b610924565b6108df565b610862565b610839565b6107ce565b610655565b610504565b610484565b610358565b610236565b6101dc565b346101995760203660031901126101995760043563ffffffff60e01b811680910361019957602090630a85bd0160e11b8114908115610188575b8115610177575b506040519015158152f35b6301ffc9a760e01b1490503861016c565b630271189760e51b81149150610165565b600080fd5b6001600160a01b0381160361019957565b9181601f84011215610199578235916001600160401b038311610199576020838186019501011161019957565b34610199576080366003190112610199576101f860043561019e565b61020360243561019e565b6064356001600160401b038111610199576102229036906004016101af565b5050604051630a85bd0160e11b8152602090f35b346101995760031960603682011261019957600435906001600160401b03821161019957610120908236030112610199576044357f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036102e3576102ad6102c59260243590600401610da4565b90806102c9575b506040519081529081906020820190565b0390f35b600080808093338219f1506102dc610e15565b50386102b4565b60405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e74000000006044820152606490fd5b9181601f84011215610199578235916001600160401b038311610199576020808501948460051b01011161019957565b34610199576060366003190112610199576001600160401b0360043581811161019957610389903690600401610328565b602492919235828111610199576103a4903690600401610328565b92604435908111610199576103bd903690600401610328565b9390916103c8610e45565b84841480610473575b6103da90610cc2565b8161042757505060005b8281106103ed57005b80610421610406610401600194878a610d1a565b610d2f565b61041b610414848988610d6e565b369161061e565b90610ed3565b016103e4565b91909460009493945b85811061043957005b8061046d61044d6104016001948a87610d1a565b610458838b89610d1a565b35610467610414858b8a610d6e565b91610efb565b01610430565b508115806103d157508185146103d1565b600080600319360112610501577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681813b156105015760405163b760faf960e01b8152306004820152918290602490829034905af180156104fc576104f0575080f35b6104f9906105af565b80f35b610d89565b80fd5b346101995760006040366003190112610501576004356105238161019e565b61052b610f12565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b156105955760449083604051958694859363040b850f60e31b855216600484015260243560248401525af180156104fc576104f0575080f35b8280fd5b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116105c257604052565b610599565b604081019081106001600160401b038211176105c257604052565b90601f801991011681019081106001600160401b038211176105c257604052565b6001600160401b0381116105c257601f01601f191660200190565b92919261062a82610603565b9161063860405193846105e2565b829481845281830111610199578281602093846000960137010152565b604036600319011261019957600480359061066f8261019e565b6024356001600160401b03811161019957366023820112156101995761069e903690602481850135910161061e565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081163081149081156107b2575b506107a15790602083926106e6610f12565b6040516352d1902d60e01b8152938491829088165afa60009281610770575b50610733575050604051634c9c8ce360e01b81526001600160a01b0390921690820190815281906020010390fd5b83836000805160206112678339815191528403610754576100198383610fd1565b604051632a87526960e21b815290810184815281906020010390fd5b61079391935060203d60201161079a575b61078b81836105e2565b810190610d95565b9138610705565b503d610781565b60405163703e46dd60e11b81528390fd5b90508160008051602061126783398151915254161415386106d4565b34610199576000366003190112610199577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036108275760206040516000805160206112678339815191528152f35b60405163703e46dd60e11b8152600490fd5b34610199576000366003190112610199576000546040516001600160a01b039091168152602090f35b34610199576000366003190112610199576040805190610881826105c7565b60058252602090640352e302e360dc1b6020840152604051916020835283519182602085015260005b8381106108cc5784604081866000838284010152601f80199101168101030190f35b85810183015185820183015282016108aa565b34610199576000366003190112610199576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610199576060366003190112610199576004356109418161019e565b604435906001600160401b0382116101995761096f6109676100199336906004016101af565b610414610e45565b9060243590610efb565b346101995760a03660031901126101995761099560043561019e565b6109a060243561019e565b6001600160401b03604435818111610199576109c0903690600401610328565b5050606435818111610199576109da903690600401610328565b5050608435908111610199576109f49036906004016101af565b505060405163bc197c8160e01b8152602090f35b34610199576000366003190112610199576040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156104fc57602091600091610a78575b50604051908152f35b610a8f9150823d841161079a5761078b81836105e2565b38610a6f565b3461019957602036600319011261019957600435610ab28161019e565b60008051602061128783398151915254906001600160401b0360ff8360401c1615921680159081610bc2575b6001149081610bb8575b159081610baf575b50610b9d57600080516020611287833981519152805467ffffffffffffffff19166001179055610b249082610b7357610f69565b610b2a57005b600080516020611287833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b600080516020611287833981519152805460ff60401b191668010000000000000000179055610f69565b60405163f92ee8a960e01b8152600490fd5b90501538610af0565b303b159150610ae8565b839150610ade565b3461019957600036600319011261019957604051631aab3f0d60e11b8152306004820152600060248201526020816044817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156104fc576102c591600091610c4957506040519081529081906020820190565b610c62915060203d60201161079a5761078b81836105e2565b386102b4565b346101995760a036600319011261019957610c8460043561019e565b610c8f60243561019e565b6084356001600160401b03811161019957610cae9036906004016101af565b505060405163f23a6e6160e01b8152602090f35b15610cc957565b60405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b6044820152606490fd5b634e487b7160e01b600052603260045260246000fd5b9190811015610d2a5760051b0190565b610d04565b35610d398161019e565b90565b903590601e198136030182121561019957018035906001600160401b0382116101995760200191813603831361019957565b90821015610d2a57610d859160051b810190610d3c565b9091565b6040513d6000823e3d90fd5b90816020910312610199575190565b907f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52603c600020610e05610dfc60018060a01b0392610df6610414856000541696610100810190610d3c565b90611078565b909291926110d4565b1603610e1057600090565b600190565b3d15610e40573d90610e2682610603565b91610e3460405193846105e2565b82523d6000602084013e565b606090565b60018060a01b03807f0000000000000000000000000000000000000000000000000000000000000000163314908115610ec5575b5015610e8157565b606460405162461bcd60e51b815260206004820152602060248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152fd5b905060005416331438610e79565b600091829182602083519301915af1610eea610e15565b9015610ef35750565b602081519101fd5b916000928392602083519301915af1610eea610e15565b6000546001600160a01b031633148015610f60575b15610f2e57565b60405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b6044820152606490fd5b50303314610f27565b600080546001600160a01b0319166001600160a01b039283169081178255917f000000000000000000000000000000000000000000000000000000000000000016907f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de9080a3565b90813b156110575760008051602061126783398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a280511561103c5761103991611161565b50565b50503461104557565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b81519190604183036110a9576110a292506020820151906060604084015193015160001a9061117f565b9192909190565b505060009160029190565b600411156110be57565b634e487b7160e01b600052602160045260246000fd5b6110dd816110b4565b806110e6575050565b6110ef816110b4565b600181036111095760405163f645eedf60e01b8152600490fd5b611112816110b4565b600281036111335760405163fce698f760e01b815260048101839052602490fd5b8061113f6003926110b4565b146111475750565b6040516335e2f38360e21b81526004810191909152602490fd5b600080610d3993602081519101845af4611179610e15565b91611203565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084116111f757926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa156104fc5780516001600160a01b038116156111ee57918190565b50809160019190565b50505060009160039190565b9061122a575080511561121857805190602001fd5b60405163d6bda27560e01b8152600490fd5b8151158061125d575b61123b575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561123356fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212206c2c7b93a0987919a131528f5af693fd610321f18f91fecd0fc7fecb5a7619ce64736f6c63430008190033",
}

// SimpleAccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleAccountFactoryMetaData.ABI instead.
var SimpleAccountFactoryABI = SimpleAccountFactoryMetaData.ABI

// SimpleAccountFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleAccountFactoryMetaData.Bin instead.
var SimpleAccountFactoryBin = SimpleAccountFactoryMetaData.Bin

// DeploySimpleAccountFactory deploys a new Ethereum contract, binding an instance of SimpleAccountFactory to it.
func DeploySimpleAccountFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *SimpleAccountFactory, error) {
	parsed, err := SimpleAccountFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleAccountFactoryBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleAccountFactory{SimpleAccountFactoryCaller: SimpleAccountFactoryCaller{contract: contract}, SimpleAccountFactoryTransactor: SimpleAccountFactoryTransactor{contract: contract}, SimpleAccountFactoryFilterer: SimpleAccountFactoryFilterer{contract: contract}}, nil
}

// SimpleAccountFactory is an auto generated Go binding around an Ethereum contract.
type SimpleAccountFactory struct {
	SimpleAccountFactoryCaller     // Read-only binding to the contract
	SimpleAccountFactoryTransactor // Write-only binding to the contract
	SimpleAccountFactoryFilterer   // Log filterer for contract events
}

// SimpleAccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleAccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleAccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleAccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleAccountFactorySession struct {
	Contract     *SimpleAccountFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SimpleAccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleAccountFactoryCallerSession struct {
	Contract *SimpleAccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SimpleAccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleAccountFactoryTransactorSession struct {
	Contract     *SimpleAccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SimpleAccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleAccountFactoryRaw struct {
	Contract *SimpleAccountFactory // Generic contract binding to access the raw methods on
}

// SimpleAccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleAccountFactoryCallerRaw struct {
	Contract *SimpleAccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleAccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleAccountFactoryTransactorRaw struct {
	Contract *SimpleAccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleAccountFactory creates a new instance of SimpleAccountFactory, bound to a specific deployed contract.
func NewSimpleAccountFactory(address common.Address, backend bind.ContractBackend) (*SimpleAccountFactory, error) {
	contract, err := bindSimpleAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleAccountFactory{SimpleAccountFactoryCaller: SimpleAccountFactoryCaller{contract: contract}, SimpleAccountFactoryTransactor: SimpleAccountFactoryTransactor{contract: contract}, SimpleAccountFactoryFilterer: SimpleAccountFactoryFilterer{contract: contract}}, nil
}

// NewSimpleAccountFactoryCaller creates a new read-only instance of SimpleAccountFactory, bound to a specific deployed contract.
func NewSimpleAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*SimpleAccountFactoryCaller, error) {
	contract, err := bindSimpleAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAccountFactoryCaller{contract: contract}, nil
}

// NewSimpleAccountFactoryTransactor creates a new write-only instance of SimpleAccountFactory, bound to a specific deployed contract.
func NewSimpleAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleAccountFactoryTransactor, error) {
	contract, err := bindSimpleAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAccountFactoryTransactor{contract: contract}, nil
}

// NewSimpleAccountFactoryFilterer creates a new log filterer instance of SimpleAccountFactory, bound to a specific deployed contract.
func NewSimpleAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleAccountFactoryFilterer, error) {
	contract, err := bindSimpleAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleAccountFactoryFilterer{contract: contract}, nil
}

// bindSimpleAccountFactory binds a generic wrapper to an already deployed contract.
func bindSimpleAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleAccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAccountFactory *SimpleAccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleAccountFactory.Contract.SimpleAccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAccountFactory *SimpleAccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAccountFactory.Contract.SimpleAccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAccountFactory *SimpleAccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleAccountFactory.Contract.SimpleAccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAccountFactory *SimpleAccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleAccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAccountFactory *SimpleAccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAccountFactory *SimpleAccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleAccountFactory.Contract.contract.Transact(opts, method, params...)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_SimpleAccountFactory *SimpleAccountFactoryCaller) AccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleAccountFactory.contract.Call(opts, &out, "accountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_SimpleAccountFactory *SimpleAccountFactorySession) AccountImplementation() (common.Address, error) {
	return _SimpleAccountFactory.Contract.AccountImplementation(&_SimpleAccountFactory.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_SimpleAccountFactory *SimpleAccountFactoryCallerSession) AccountImplementation() (common.Address, error) {
	return _SimpleAccountFactory.Contract.AccountImplementation(&_SimpleAccountFactory.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_SimpleAccountFactory *SimpleAccountFactoryCaller) GetAddress(opts *bind.CallOpts, owner common.Address, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SimpleAccountFactory.contract.Call(opts, &out, "getAddress", owner, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_SimpleAccountFactory *SimpleAccountFactorySession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _SimpleAccountFactory.Contract.GetAddress(&_SimpleAccountFactory.CallOpts, owner, salt)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_SimpleAccountFactory *SimpleAccountFactoryCallerSession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _SimpleAccountFactory.Contract.GetAddress(&_SimpleAccountFactory.CallOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_SimpleAccountFactory *SimpleAccountFactoryTransactor) CreateAccount(opts *bind.TransactOpts, owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _SimpleAccountFactory.contract.Transact(opts, "createAccount", owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_SimpleAccountFactory *SimpleAccountFactorySession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _SimpleAccountFactory.Contract.CreateAccount(&_SimpleAccountFactory.TransactOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_SimpleAccountFactory *SimpleAccountFactoryTransactorSession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _SimpleAccountFactory.Contract.CreateAccount(&_SimpleAccountFactory.TransactOpts, owner, salt)
}
