// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package brokenblsaccountfactory

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

// BrokenBLSAccountFactoryMetaData contains all meta data concerning the BrokenBLSAccountFactory contract.
var BrokenBLSAccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"accountImplementation\",\"outputs\":[{\"internalType\":\"contractBrokenBLSAccount\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"aPublicKey\",\"type\":\"uint256[4]\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"contractBrokenBLSAccount\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"aPublicKey\",\"type\":\"uint256[4]\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a03460d857611cb2906001600160401b0390601f38849003908101601f19168201908382118383101760c2578083916040968794855283398101031260d85780516001600160a01b0380821692909183900360d8576020015190811680910360d857835192611507918285019182118583101760c257859385936107ab8539825260208201520301906000f0801560b757608052516106cd90816100de8239608051818181609e0152818161027501526103470152f35b50513d6000823e3d90fd5b634e487b7160e01b600052604160045260246000fd5b600080fdfe608080604052600436101561001357600080fd5b600090813560e01c90816311464fbe1461008a5750806319c2a1b2146100705763de3398dd1461004257600080fd5b3461006d57602061005b61005536610109565b906102c5565b6040516001600160a01b039091168152f35b80fd5b503461006d57602061005b61008436610109565b906101fe565b9050346100cd57816003193601126100cd577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b90601f8019910116810190811067ffffffffffffffff8211176100f357604052565b634e487b7160e01b600052604160045260246000fd5b9060a060031983011261016d5760043591806043121561016d57604051906080820182811067ffffffffffffffff8211176100f3576040528160a49160a41161016d576024905b82821061015d5750505090565b8135815260209182019101610150565b600080fd5b6080810192916000915b6004831061018957505050565b60019082518152602080910192019201919061017c565b60005b8381106101b35750506000910152565b81810151838201526020016101a3565b909160609260018060a01b03168252604060208301526101f281518092816040860152602086860191016101a0565b601f01601f1916010190565b61020882826102c5565b803b6102b65750604051637723979b60e11b60208201526001600160a01b039261024990829061023b9060248301610172565b03601f1981018352826100d1565b604051906102d38083019183831067ffffffffffffffff8411176100f357839261029b926103c58539867f000000000000000000000000000000000000000000000000000000000000000016906101c3565b03906000f580156102aa571690565b6040513d6000823e3d90fd5b6001600160a01b031692915050565b600b610328926055926102d360209061039f6103ab836040968751906102ed838701836100d1565b858252828201956103c587398851637723979b60e11b848201526001600160a01b039c8d9261036d9261037992909182919060248301610172565b039061033c601f19928381018352826100d1565b8c51938491888301967f000000000000000000000000000000000000000000000000000000000000000016876101c3565b039081018352826100d1565b8951958693610390868601998a92519283916101a0565b840191518093868401906101a0565b010380845201826100d1565b5190208351938401528201523081520160ff815320169056fe60806040526102d38038038061001481610194565b92833981019060408183031261018f5780516001600160a01b03811680820361018f5760208381015190936001600160401b03821161018f570184601f8201121561018f5780519061006d610068836101cf565b610194565b9582875285838301011161018f57849060005b83811061017b57505060009186010152813b15610163577f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b03191682179055604051907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a28351156101455750600080848461012c96519101845af4903d1561013c573d61011c610068826101cf565b908152600081943d92013e6101ea565b505b6040516085908161024e8239f35b606092506101ea565b9250505034610154575061012e565b63b398979f60e01b8152600490fd5b60249060405190634c9c8ce360e01b82526004820152fd5b818101830151888201840152869201610080565b600080fd5b6040519190601f01601f191682016001600160401b038111838210176101b957604052565b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116101b957601f01601f191660200190565b9061021157508051156101ff57805190602001fd5b60405163d6bda27560e01b8152600490fd5b81511580610244575b610222575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561021a56fe60806040527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc54600090819081906001600160a01b0316368280378136915af43d82803e15604b573d90f35b3d90fdfea2646970667358221220fe8f61e3fa58bb8add11a833212ef9684476c5dbbd688eea5bcc0b0062abf4cc64736f6c63430008190033a2646970667358221220c202f9265cd9bb0efb5071815e0d5f5dd93cad379779e1600d99d6927bbc2cd664736f6c6343000819003360e03461016b57611507906001600160401b03601f38849003908101601f191683019082821184831017610170578084916040968794855283398101031261016b5781516001600160a01b0392838216820361016b5760200151928316830361016b573060805260a0527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff82861c1661015a578080831603610116575b50505060c052516113809081610187823960805181818161076801526108a1015260a0518181816102a00152818161051d015281816105bf015281816109b501528181610af401528181610cb8015281816110080152818161113401526111a3015260c0518181816102dd01526103830152f35b6001600160401b0319909116811790915582519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a13880806100a2565b845163f92ee8a960e01b8152600490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c806301ffc9a71461015b578063150b7a021461015657806319822f7c14610151578063245a7bfc1461014c57806347e1da2a146101475780634a58db19146101425780634d44560d1461013d5780634f1ef2861461013857806352d1902d146101335780638da5cb5b1461012e578063ad3cb1cc14610129578063b0d691fe14610124578063b61d27f61461011f578063bc197c811461011a578063c399ec8814610115578063c4d66de814610110578063d087d2881461010b578063e02afbae14610106578063ee472f36146101015763f23a6e610361000e57610e80565b610d84565b610d24565b610c85565b610b55565b610ac8565b610a39565b6109e4565b61099f565b610922565b6108f9565b61088e565b610715565b61058e565b61050e565b6103e2565b61036d565b610266565b61020c565b346101c95760203660031901126101c95760043563ffffffff60e01b81168091036101c957602090630a85bd0160e11b81149081156101b8575b81156101a7575b506040519015158152f35b6301ffc9a760e01b1490503861019c565b630271189760e51b81149150610195565b600080fd5b6001600160a01b038116036101c957565b9181601f840112156101c9578235916001600160401b0383116101c957602083818601950101116101c957565b346101c95760803660031901126101c9576102286004356101ce565b6102336024356101ce565b6064356001600160401b0381116101c9576102529036906004016101df565b5050604051630a85bd0160e11b8152602090f35b346101c9576003196060368201126101c957600435906001600160401b0382116101c95761012091360301126101c9576001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811633036103285761032490600060408051926102db84610639565b7f000000000000000000000000000000000000000000000000000000000000000016928381528260208201520152610314604435610fe0565b6040519081529081906020820190565b0390f35b60405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e74000000006044820152606490fd5b346101c95760003660031901126101c9576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b9181601f840112156101c9578235916001600160401b0383116101c9576020808501948460051b0101116101c957565b346101c95760603660031901126101c9576001600160401b036004358181116101c9576104139036906004016103b2565b6024929192358281116101c95761042e9036906004016103b2565b926044359081116101c9576104479036906004016103b2565b939091610452610ffe565b848414806104fd575b61046490610eda565b816104b157505060005b82811061047757005b806104ab61049061048b600194878a610f32565b610f47565b6104a561049e848988610f54565b36916106de565b9061108c565b0161046e565b91909460009493945b8581106104c357005b806104f76104d761048b6001948a87610f32565b6104e2838b89610f32565b356104f161049e858b8a610f54565b916110b4565b016104ba565b5081158061045b575081851461045b565b60008060031936011261058b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681813b1561058b5760405163b760faf960e01b8152306004820152918290602490829034905af180156105865761057a575080f35b61058390610659565b80f35b610f95565b80fd5b346101c9576000604036600319011261058b576004356105ad816101ce565b6105b56110cb565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b1561061f5760449083604051958694859363040b850f60e31b855216600484015260243560248401525af180156105865761057a575080f35b8280fd5b634e487b7160e01b600052604160045260246000fd5b606081019081106001600160401b0382111761065457604052565b610623565b6001600160401b03811161065457604052565b604081019081106001600160401b0382111761065457604052565b608081019081106001600160401b0382111761065457604052565b90601f801991011681019081106001600160401b0382111761065457604052565b6001600160401b03811161065457601f01601f191660200190565b9291926106ea826106c3565b916106f860405193846106a2565b8294818452818301116101c9578281602093846000960137010152565b60403660031901126101c957600480359061072f826101ce565b6024356001600160401b0381116101c957366023820112156101c95761075e90369060248185013591016106de565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116308114908115610872575b506108615790602083926107a66110cb565b6040516352d1902d60e01b8152938491829088165afa60009281610830575b506107f3575050604051634c9c8ce360e01b81526001600160a01b0390921690820190815281906020010390fd5b838360008051602061130b83398151915284036108145761001983836111ea565b604051632a87526960e21b815290810184815281906020010390fd5b61085391935060203d60201161085a575b61084b81836106a2565b810190610fa1565b91386107c5565b503d610841565b60405163703e46dd60e11b81528390fd5b90508160008051602061130b8339815191525416141538610794565b346101c95760003660031901126101c9577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036108e757602060405160008051602061130b8339815191528152f35b60405163703e46dd60e11b8152600490fd5b346101c95760003660031901126101c9576000546040516001600160a01b039091168152602090f35b346101c95760003660031901126101c95760408051906109418261066c565b60058252602090640352e302e360dc1b6020840152604051916020835283519182602085015260005b83811061098c5784604081866000838284010152601f80199101168101030190f35b858101830151858201830152820161096a565b346101c95760003660031901126101c9576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101c95760603660031901126101c957600435610a01816101ce565b604435906001600160401b0382116101c957610a2f610a276100199336906004016101df565b61049e610ffe565b90602435906110b4565b346101c95760a03660031901126101c957610a556004356101ce565b610a606024356101ce565b6001600160401b036044358181116101c957610a809036906004016103b2565b50506064358181116101c957610a9a9036906004016103b2565b50506084359081116101c957610ab49036906004016101df565b505060405163bc197c8160e01b8152602090f35b346101c95760003660031901126101c9576040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561058657602091600091610b38575b50604051908152f35b610b4f9150823d841161085a5761084b81836106a2565b38610b2f565b346101c95760203660031901126101c957600435610b72816101ce565b60008051602061132b83398151915254906001600160401b0360ff8360401c1615921680159081610c7d575b6001149081610c73575b159081610c6a575b50610c585760008051602061132b833981519152805467ffffffffffffffff19166001179055610be49082610c3357611182565b610bea57005b60008051602061132b833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b60008051602061132b833981519152805460ff60401b1916600160401b179055611182565b60405163f92ee8a960e01b8152600490fd5b90501538610bb0565b303b159150610ba8565b839150610b9e565b346101c95760003660031901126101c957604051631aab3f0d60e11b8152306004820152600060248201526020816044817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156105865761032491600091610d05575b506040519081529081906020820190565b610d1e915060203d60201161085a5761084b81836106a2565b38610cf4565b346101c95760003660031901126101c9576080604051610d4381610687565b369037604051610d5281610687565b608036823760405190600090825b60048310610d6d57608084f35b600190825181526020809101920192019190610d60565b346101c95760803660031901126101c95736602312156101c957604051610daa81610687565b6084366084116101c9576004915b818310610e705760008051602061132b833981519152546001600160401b0360ff8260401c1615911680159081610e68575b6001149081610e5e575b159081610e55575b50610c585760008051602061132b833981519152805467ffffffffffffffff1916600117905580610e30575b610be4611122565b60008051602061132b833981519152805460ff60401b1916600160401b179055610e28565b90501582610dfc565b303b159150610df4565b829150610dea565b8235815260209283019201610db8565b346101c95760a03660031901126101c957610e9c6004356101ce565b610ea76024356101ce565b6084356001600160401b0381116101c957610ec69036906004016101df565b505060405163f23a6e6160e01b8152602090f35b15610ee157565b60405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b6044820152606490fd5b634e487b7160e01b600052603260045260246000fd5b9190811015610f425760051b0190565b610f1c565b35610f51816101ce565b90565b9190811015610f425760051b81013590601e19813603018212156101c95701908135916001600160401b0383116101c95760200182360381136101c9579190565b6040513d6000823e3d90fd5b908160209103126101c9575190565b3d15610fdb573d90610fc1826106c3565b91610fcf60405193846106a2565b82523d6000602084013e565b606090565b80610fe85750565b600080808093338219f150610ffb610fb0565b50565b60018060a01b03807f000000000000000000000000000000000000000000000000000000000000000016331490811561107e575b501561103a57565b606460405162461bcd60e51b815260206004820152602060248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152fd5b905060005416331438611032565b600091829182602083519301915af16110a3610fb0565b90156110ac5750565b602081519101fd5b916000928392602083519301915af16110a3610fb0565b6000546001600160a01b031633148015611119575b156110e757565b60405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b6044820152606490fd5b503033146110e0565b600080546001600160a01b03191681557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de8280a3565b600080546001600160a01b0319166001600160a01b039283169081178255917f000000000000000000000000000000000000000000000000000000000000000016907f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de9080a3565b90813b1561126d5760008051602061130b83398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a280511561125257610ffb9161128e565b50503461125b57565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b600080610f5193602081519101845af46112a6610fb0565b91906112ce57508051156112bc57805190602001fd5b60405163d6bda27560e01b8152600490fd5b81511580611301575b6112df575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b156112d756fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212203c2da3c5641277907952b8f68fe7b1bfb2476bf609bb31900585a098d8726e2364736f6c63430008190033",
}

// BrokenBLSAccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use BrokenBLSAccountFactoryMetaData.ABI instead.
var BrokenBLSAccountFactoryABI = BrokenBLSAccountFactoryMetaData.ABI

// BrokenBLSAccountFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BrokenBLSAccountFactoryMetaData.Bin instead.
var BrokenBLSAccountFactoryBin = BrokenBLSAccountFactoryMetaData.Bin

// DeployBrokenBLSAccountFactory deploys a new Ethereum contract, binding an instance of BrokenBLSAccountFactory to it.
func DeployBrokenBLSAccountFactory(auth *bind.TransactOpts, backend bind.ContractBackend, entryPoint common.Address, aggregator common.Address) (common.Address, *types.Transaction, *BrokenBLSAccountFactory, error) {
	parsed, err := BrokenBLSAccountFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BrokenBLSAccountFactoryBin), backend, entryPoint, aggregator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BrokenBLSAccountFactory{BrokenBLSAccountFactoryCaller: BrokenBLSAccountFactoryCaller{contract: contract}, BrokenBLSAccountFactoryTransactor: BrokenBLSAccountFactoryTransactor{contract: contract}, BrokenBLSAccountFactoryFilterer: BrokenBLSAccountFactoryFilterer{contract: contract}}, nil
}

// BrokenBLSAccountFactory is an auto generated Go binding around an Ethereum contract.
type BrokenBLSAccountFactory struct {
	BrokenBLSAccountFactoryCaller     // Read-only binding to the contract
	BrokenBLSAccountFactoryTransactor // Write-only binding to the contract
	BrokenBLSAccountFactoryFilterer   // Log filterer for contract events
}

// BrokenBLSAccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrokenBLSAccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokenBLSAccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrokenBLSAccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokenBLSAccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrokenBLSAccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokenBLSAccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrokenBLSAccountFactorySession struct {
	Contract     *BrokenBLSAccountFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BrokenBLSAccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrokenBLSAccountFactoryCallerSession struct {
	Contract *BrokenBLSAccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// BrokenBLSAccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrokenBLSAccountFactoryTransactorSession struct {
	Contract     *BrokenBLSAccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// BrokenBLSAccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrokenBLSAccountFactoryRaw struct {
	Contract *BrokenBLSAccountFactory // Generic contract binding to access the raw methods on
}

// BrokenBLSAccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrokenBLSAccountFactoryCallerRaw struct {
	Contract *BrokenBLSAccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// BrokenBLSAccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrokenBLSAccountFactoryTransactorRaw struct {
	Contract *BrokenBLSAccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrokenBLSAccountFactory creates a new instance of BrokenBLSAccountFactory, bound to a specific deployed contract.
func NewBrokenBLSAccountFactory(address common.Address, backend bind.ContractBackend) (*BrokenBLSAccountFactory, error) {
	contract, err := bindBrokenBLSAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BrokenBLSAccountFactory{BrokenBLSAccountFactoryCaller: BrokenBLSAccountFactoryCaller{contract: contract}, BrokenBLSAccountFactoryTransactor: BrokenBLSAccountFactoryTransactor{contract: contract}, BrokenBLSAccountFactoryFilterer: BrokenBLSAccountFactoryFilterer{contract: contract}}, nil
}

// NewBrokenBLSAccountFactoryCaller creates a new read-only instance of BrokenBLSAccountFactory, bound to a specific deployed contract.
func NewBrokenBLSAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*BrokenBLSAccountFactoryCaller, error) {
	contract, err := bindBrokenBLSAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrokenBLSAccountFactoryCaller{contract: contract}, nil
}

// NewBrokenBLSAccountFactoryTransactor creates a new write-only instance of BrokenBLSAccountFactory, bound to a specific deployed contract.
func NewBrokenBLSAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*BrokenBLSAccountFactoryTransactor, error) {
	contract, err := bindBrokenBLSAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrokenBLSAccountFactoryTransactor{contract: contract}, nil
}

// NewBrokenBLSAccountFactoryFilterer creates a new log filterer instance of BrokenBLSAccountFactory, bound to a specific deployed contract.
func NewBrokenBLSAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*BrokenBLSAccountFactoryFilterer, error) {
	contract, err := bindBrokenBLSAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrokenBLSAccountFactoryFilterer{contract: contract}, nil
}

// bindBrokenBLSAccountFactory binds a generic wrapper to an already deployed contract.
func bindBrokenBLSAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BrokenBLSAccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrokenBLSAccountFactory.Contract.BrokenBLSAccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrokenBLSAccountFactory.Contract.BrokenBLSAccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrokenBLSAccountFactory.Contract.BrokenBLSAccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrokenBLSAccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrokenBLSAccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrokenBLSAccountFactory.Contract.contract.Transact(opts, method, params...)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactoryCaller) AccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrokenBLSAccountFactory.contract.Call(opts, &out, "accountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactorySession) AccountImplementation() (common.Address, error) {
	return _BrokenBLSAccountFactory.Contract.AccountImplementation(&_BrokenBLSAccountFactory.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactoryCallerSession) AccountImplementation() (common.Address, error) {
	return _BrokenBLSAccountFactory.Contract.AccountImplementation(&_BrokenBLSAccountFactory.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0xde3398dd.
//
// Solidity: function getAddress(uint256 salt, uint256[4] aPublicKey) view returns(address)
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactoryCaller) GetAddress(opts *bind.CallOpts, salt *big.Int, aPublicKey [4]*big.Int) (common.Address, error) {
	var out []interface{}
	err := _BrokenBLSAccountFactory.contract.Call(opts, &out, "getAddress", salt, aPublicKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0xde3398dd.
//
// Solidity: function getAddress(uint256 salt, uint256[4] aPublicKey) view returns(address)
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactorySession) GetAddress(salt *big.Int, aPublicKey [4]*big.Int) (common.Address, error) {
	return _BrokenBLSAccountFactory.Contract.GetAddress(&_BrokenBLSAccountFactory.CallOpts, salt, aPublicKey)
}

// GetAddress is a free data retrieval call binding the contract method 0xde3398dd.
//
// Solidity: function getAddress(uint256 salt, uint256[4] aPublicKey) view returns(address)
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactoryCallerSession) GetAddress(salt *big.Int, aPublicKey [4]*big.Int) (common.Address, error) {
	return _BrokenBLSAccountFactory.Contract.GetAddress(&_BrokenBLSAccountFactory.CallOpts, salt, aPublicKey)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x19c2a1b2.
//
// Solidity: function createAccount(uint256 salt, uint256[4] aPublicKey) returns(address)
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactoryTransactor) CreateAccount(opts *bind.TransactOpts, salt *big.Int, aPublicKey [4]*big.Int) (*types.Transaction, error) {
	return _BrokenBLSAccountFactory.contract.Transact(opts, "createAccount", salt, aPublicKey)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x19c2a1b2.
//
// Solidity: function createAccount(uint256 salt, uint256[4] aPublicKey) returns(address)
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactorySession) CreateAccount(salt *big.Int, aPublicKey [4]*big.Int) (*types.Transaction, error) {
	return _BrokenBLSAccountFactory.Contract.CreateAccount(&_BrokenBLSAccountFactory.TransactOpts, salt, aPublicKey)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x19c2a1b2.
//
// Solidity: function createAccount(uint256 salt, uint256[4] aPublicKey) returns(address)
func (_BrokenBLSAccountFactory *BrokenBLSAccountFactoryTransactorSession) CreateAccount(salt *big.Int, aPublicKey [4]*big.Int) (*types.Transaction, error) {
	return _BrokenBLSAccountFactory.Contract.CreateAccount(&_BrokenBLSAccountFactory.TransactOpts, salt, aPublicKey)
}
