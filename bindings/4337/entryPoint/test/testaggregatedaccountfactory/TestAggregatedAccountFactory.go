// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testaggregatedaccountfactory

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

// TestAggregatedAccountFactoryMetaData contains all meta data concerning the TestAggregatedAccountFactory contract.
var TestAggregatedAccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"anAggregator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"accountImplementation\",\"outputs\":[{\"internalType\":\"contractTestAggregatedAccount\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"contractTestAggregatedAccount\",\"name\":\"ret\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a03460d857611a44906001600160401b0390601f38849003908101601f19168201908382118383101760c2578083916040968794855283398101031260d85780516001600160a01b0380821692909183900360d8576020015190811680910360d857835192611304918285019182118583101760c257859385936107408539825260208201520301906000f0801560b7576080525161066290816100de8239608051818181609e0152818161021b01526102da0152f35b50513d6000823e3d90fd5b634e487b7160e01b600052604160045260246000fd5b600080fdfe608080604052600436101561001357600080fd5b600090813560e01c90816311464fbe1461008a575080635fbfb9cf1461007057638cb84e181461004257600080fd5b3461006d57602061005b610055366100d1565b9061026b565b6040516001600160a01b039091168152f35b80fd5b503461006d57602061005b610084366100d1565b906101ae565b9050346100cd57816003193601126100cd577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b60409060031901126100f7576004356001600160a01b03811681036100f7579060243590565b600080fd5b6060810190811067ffffffffffffffff82111761011857604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff82111761011857604052565b60005b8381106101635750506000910152565b8181015183820152602001610153565b909160609260018060a01b03168252604060208301526101a28151809281604086015260208686019101610150565b601f01601f1916010190565b906101b9818361026b565b803b61025c575060405163189acdbd60e31b60208201526001600160a01b0392831660248083019190915281526101ef816100fc565b604051906102d38083019183831067ffffffffffffffff8411176101185783926102419261035a8539867f00000000000000000000000000000000000000000000000000000000000000001690610173565b03906000f58015610250571690565b6040513d6000823e3d90fd5b6001600160a01b031692915050565b600b60559161030e936102d3602090610334610340836040968751906102938387018361012e565b8582528282019561035a873961030060018060a01b039c8d92838c519163189acdbd60e31b88840152166024820152602481526102cf816100fc565b8b51928391878301957f00000000000000000000000000000000000000000000000000000000000000001686610173565b03601f19810183528261012e565b8951958693610325868601998a9251928391610150565b84019151809386840190610150565b0103808452018261012e565b5190208351938401528201523081520160ff815320169056fe60806040526102d38038038061001481610194565b92833981019060408183031261018f5780516001600160a01b03811680820361018f5760208381015190936001600160401b03821161018f570184601f8201121561018f5780519061006d610068836101cf565b610194565b9582875285838301011161018f57849060005b83811061017b57505060009186010152813b15610163577f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b03191682179055604051907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a28351156101455750600080848461012c96519101845af4903d1561013c573d61011c610068826101cf565b908152600081943d92013e6101ea565b505b6040516085908161024e8239f35b606092506101ea565b9250505034610154575061012e565b63b398979f60e01b8152600490fd5b60249060405190634c9c8ce360e01b82526004820152fd5b818101830151888201840152869201610080565b600080fd5b6040519190601f01601f191682016001600160401b038111838210176101b957604052565b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116101b957601f01601f191660200190565b9061021157508051156101ff57805190602001fd5b60405163d6bda27560e01b8152600490fd5b81511580610244575b610222575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561021a56fe60806040527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc54600090819081906001600160a01b0316368280378136915af43d82803e15604b573d90f35b3d90fdfea2646970667358221220fe8f61e3fa58bb8add11a833212ef9684476c5dbbd688eea5bcc0b0062abf4cc64736f6c63430008190033a264697066735822122081e437c9f4cd21662440c528e476bbd1a7b304d3454a0d2b007b7cf617e7245e64736f6c6343000819003360e03461016457611304906001600160401b03601f38849003908101601f19168301908282118483101761016957808491604096879485528339810103126101645781516001600160a01b0392838216820361016457602001519283168303610164573060805260a0527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff82861c1661015357808083160361010f575b50505060c052516111849081610180823960805181818161072d0152610866015260a051818181610280015281816104fd0152818161059f0152818161097a01528181610ab901528181610c8001528181610e2e0152610ed4015260c0518181816102bd01526103630152f35b6001600160401b0319909116811790915582519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a13880806100a2565b845163f92ee8a960e01b8152600490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c806301ffc9a71461013b578063150b7a021461013657806319822f7c14610131578063245a7bfc1461012c57806347e1da2a146101275780634a58db19146101225780634d44560d1461011d5780634f1ef2861461011857806352d1902d146101135780638da5cb5b1461010e578063ad3cb1cc14610109578063b0d691fe14610104578063b61d27f6146100ff578063bc197c81146100fa578063c399ec88146100f5578063c4d66de8146100f0578063d087d288146100eb5763f23a6e610361000e57610cec565b610c4d565b610b1a565b610a8d565b6109fe565b6109a9565b610964565b6108e7565b6108be565b610853565b6106da565b61056e565b6104ee565b6103c2565b61034d565b610246565b6101ec565b346101a95760203660031901126101a95760043563ffffffff60e01b81168091036101a957602090630a85bd0160e11b8114908115610198575b8115610187575b506040519015158152f35b6301ffc9a760e01b1490503861017c565b630271189760e51b81149150610175565b600080fd5b6001600160a01b038116036101a957565b9181601f840112156101a9578235916001600160401b0383116101a957602083818601950101116101a957565b346101a95760803660031901126101a9576102086004356101ae565b6102136024356101ae565b6064356001600160401b0381116101a9576102329036906004016101bf565b5050604051630a85bd0160e11b8152602090f35b346101a9576003196060368201126101a957600435906001600160401b0382116101a95761012091360301126101a9576001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811633036103085761030490600060408051926102bb84610619565b7f0000000000000000000000000000000000000000000000000000000000000000169283815282602082015201526102f4604435610eac565b6040519081529081906020820190565b0390f35b60405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e74000000006044820152606490fd5b346101a95760003660031901126101a9576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b9181601f840112156101a9578235916001600160401b0383116101a9576020808501948460051b0101116101a957565b346101a95760603660031901126101a9576001600160401b036004358181116101a9576103f3903690600401610392565b6024929192358281116101a95761040e903690600401610392565b926044359081116101a957610427903690600401610392565b939091610432610eca565b848414806104dd575b61044490610d46565b8161049157505060005b82811061045757005b8061048b61047061046b600194878a610d9e565b610db3565b61048561047e848988610dc0565b36916106a3565b90610f58565b0161044e565b91909460009493945b8581106104a357005b806104d76104b761046b6001948a87610d9e565b6104c2838b89610d9e565b356104d161047e858b8a610dc0565b91610f80565b0161049a565b5081158061043b575081851461043b565b60008060031936011261056b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681813b1561056b5760405163b760faf960e01b8152306004820152918290602490829034905af180156105665761055a575080f35b61056390610639565b80f35b610e01565b80fd5b346101a9576000604036600319011261056b5760043561058d816101ae565b610595610f97565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b156105ff5760449083604051958694859363040b850f60e31b855216600484015260243560248401525af180156105665761055a575080f35b8280fd5b634e487b7160e01b600052604160045260246000fd5b606081019081106001600160401b0382111761063457604052565b610603565b6001600160401b03811161063457604052565b604081019081106001600160401b0382111761063457604052565b90601f801991011681019081106001600160401b0382111761063457604052565b6001600160401b03811161063457601f01601f191660200190565b9291926106af82610688565b916106bd6040519384610667565b8294818452818301116101a9578281602093846000960137010152565b60403660031901126101a95760048035906106f4826101ae565b6024356001600160401b0381116101a957366023820112156101a95761072390369060248185013591016106a3565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116308114908115610837575b5061082657906020839261076b610f97565b6040516352d1902d60e01b8152938491829088165afa600092816107f5575b506107b8575050604051634c9c8ce360e01b81526001600160a01b0390921690820190815281906020010390fd5b838360008051602061110f83398151915284036107d9576100198383610fee565b604051632a87526960e21b815290810184815281906020010390fd5b61081891935060203d60201161081f575b6108108183610667565b810190610e0d565b913861078a565b503d610806565b60405163703e46dd60e11b81528390fd5b90508160008051602061110f8339815191525416141538610759565b346101a95760003660031901126101a9577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036108ac57602060405160008051602061110f8339815191528152f35b60405163703e46dd60e11b8152600490fd5b346101a95760003660031901126101a9576000546040516001600160a01b039091168152602090f35b346101a95760003660031901126101a95760408051906109068261064c565b60058252602090640352e302e360dc1b6020840152604051916020835283519182602085015260005b8381106109515784604081866000838284010152601f80199101168101030190f35b858101830151858201830152820161092f565b346101a95760003660031901126101a9576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101a95760603660031901126101a9576004356109c6816101ae565b604435906001600160401b0382116101a9576109f46109ec6100199336906004016101bf565b61047e610eca565b9060243590610f80565b346101a95760a03660031901126101a957610a1a6004356101ae565b610a256024356101ae565b6001600160401b036044358181116101a957610a45903690600401610392565b50506064358181116101a957610a5f903690600401610392565b50506084359081116101a957610a799036906004016101bf565b505060405163bc197c8160e01b8152602090f35b346101a95760003660031901126101a9576040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561056657602091600091610afd575b50604051908152f35b610b149150823d841161081f576108108183610667565b38610af4565b346101a95760203660031901126101a957610b366004356101ae565b60008051602061112f833981519152546001600160401b0360ff8260401c1615911680159081610c45575b6001149081610c3b575b159081610c32575b50610c205760008051602061112f833981519152805467ffffffffffffffff1916600117905580610bf6575b610ba7610e1c565b610bad57005b60008051602061112f833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b60008051602061112f833981519152805460ff60401b191668010000000000000000179055610b9f565b60405163f92ee8a960e01b8152600490fd5b90501538610b73565b303b159150610b6b565b829150610b61565b346101a95760003660031901126101a957604051631aab3f0d60e11b8152306004820152600060248201526020816044817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156105665761030491600091610ccd575b506040519081529081906020820190565b610ce6915060203d60201161081f576108108183610667565b38610cbc565b346101a95760a03660031901126101a957610d086004356101ae565b610d136024356101ae565b6084356001600160401b0381116101a957610d329036906004016101bf565b505060405163f23a6e6160e01b8152602090f35b15610d4d57565b60405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b6044820152606490fd5b634e487b7160e01b600052603260045260246000fd5b9190811015610dae5760051b0190565b610d88565b35610dbd816101ae565b90565b9190811015610dae5760051b81013590601e19813603018212156101a95701908135916001600160401b0383116101a95760200182360381136101a9579190565b6040513d6000823e3d90fd5b908160209103126101a9575190565b600080546001600160a01b03191681557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de8280a3565b3d15610ea7573d90610e8d82610688565b91610e9b6040519384610667565b82523d6000602084013e565b606090565b80610eb45750565b600080808093338219f150610ec7610e7c565b50565b60018060a01b03807f0000000000000000000000000000000000000000000000000000000000000000163314908115610f4a575b5015610f0657565b606460405162461bcd60e51b815260206004820152602060248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152fd5b905060005416331438610efe565b600091829182602083519301915af1610f6f610e7c565b9015610f785750565b602081519101fd5b916000928392602083519301915af1610f6f610e7c565b6000546001600160a01b031633148015610fe5575b15610fb357565b60405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b6044820152606490fd5b50303314610fac565b90813b156110715760008051602061110f83398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a280511561105657610ec791611092565b50503461105f57565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b600080610dbd93602081519101845af46110aa610e7c565b91906110d257508051156110c057805190602001fd5b60405163d6bda27560e01b8152600490fd5b81511580611105575b6110e3575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b156110db56fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122062d5b31babf6efc3b6725d8577215fd8f922d9b7b6ef93d8a1156c740f4c562e64736f6c63430008190033",
}

// TestAggregatedAccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use TestAggregatedAccountFactoryMetaData.ABI instead.
var TestAggregatedAccountFactoryABI = TestAggregatedAccountFactoryMetaData.ABI

// TestAggregatedAccountFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestAggregatedAccountFactoryMetaData.Bin instead.
var TestAggregatedAccountFactoryBin = TestAggregatedAccountFactoryMetaData.Bin

// DeployTestAggregatedAccountFactory deploys a new Ethereum contract, binding an instance of TestAggregatedAccountFactory to it.
func DeployTestAggregatedAccountFactory(auth *bind.TransactOpts, backend bind.ContractBackend, anEntryPoint common.Address, anAggregator common.Address) (common.Address, *types.Transaction, *TestAggregatedAccountFactory, error) {
	parsed, err := TestAggregatedAccountFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestAggregatedAccountFactoryBin), backend, anEntryPoint, anAggregator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestAggregatedAccountFactory{TestAggregatedAccountFactoryCaller: TestAggregatedAccountFactoryCaller{contract: contract}, TestAggregatedAccountFactoryTransactor: TestAggregatedAccountFactoryTransactor{contract: contract}, TestAggregatedAccountFactoryFilterer: TestAggregatedAccountFactoryFilterer{contract: contract}}, nil
}

// TestAggregatedAccountFactory is an auto generated Go binding around an Ethereum contract.
type TestAggregatedAccountFactory struct {
	TestAggregatedAccountFactoryCaller     // Read-only binding to the contract
	TestAggregatedAccountFactoryTransactor // Write-only binding to the contract
	TestAggregatedAccountFactoryFilterer   // Log filterer for contract events
}

// TestAggregatedAccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestAggregatedAccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAggregatedAccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestAggregatedAccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAggregatedAccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestAggregatedAccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAggregatedAccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestAggregatedAccountFactorySession struct {
	Contract     *TestAggregatedAccountFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TestAggregatedAccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestAggregatedAccountFactoryCallerSession struct {
	Contract *TestAggregatedAccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// TestAggregatedAccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestAggregatedAccountFactoryTransactorSession struct {
	Contract     *TestAggregatedAccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// TestAggregatedAccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestAggregatedAccountFactoryRaw struct {
	Contract *TestAggregatedAccountFactory // Generic contract binding to access the raw methods on
}

// TestAggregatedAccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestAggregatedAccountFactoryCallerRaw struct {
	Contract *TestAggregatedAccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TestAggregatedAccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestAggregatedAccountFactoryTransactorRaw struct {
	Contract *TestAggregatedAccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestAggregatedAccountFactory creates a new instance of TestAggregatedAccountFactory, bound to a specific deployed contract.
func NewTestAggregatedAccountFactory(address common.Address, backend bind.ContractBackend) (*TestAggregatedAccountFactory, error) {
	contract, err := bindTestAggregatedAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestAggregatedAccountFactory{TestAggregatedAccountFactoryCaller: TestAggregatedAccountFactoryCaller{contract: contract}, TestAggregatedAccountFactoryTransactor: TestAggregatedAccountFactoryTransactor{contract: contract}, TestAggregatedAccountFactoryFilterer: TestAggregatedAccountFactoryFilterer{contract: contract}}, nil
}

// NewTestAggregatedAccountFactoryCaller creates a new read-only instance of TestAggregatedAccountFactory, bound to a specific deployed contract.
func NewTestAggregatedAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*TestAggregatedAccountFactoryCaller, error) {
	contract, err := bindTestAggregatedAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestAggregatedAccountFactoryCaller{contract: contract}, nil
}

// NewTestAggregatedAccountFactoryTransactor creates a new write-only instance of TestAggregatedAccountFactory, bound to a specific deployed contract.
func NewTestAggregatedAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TestAggregatedAccountFactoryTransactor, error) {
	contract, err := bindTestAggregatedAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestAggregatedAccountFactoryTransactor{contract: contract}, nil
}

// NewTestAggregatedAccountFactoryFilterer creates a new log filterer instance of TestAggregatedAccountFactory, bound to a specific deployed contract.
func NewTestAggregatedAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TestAggregatedAccountFactoryFilterer, error) {
	contract, err := bindTestAggregatedAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestAggregatedAccountFactoryFilterer{contract: contract}, nil
}

// bindTestAggregatedAccountFactory binds a generic wrapper to an already deployed contract.
func bindTestAggregatedAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestAggregatedAccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAggregatedAccountFactory.Contract.TestAggregatedAccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAggregatedAccountFactory.Contract.TestAggregatedAccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAggregatedAccountFactory.Contract.TestAggregatedAccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAggregatedAccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAggregatedAccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAggregatedAccountFactory.Contract.contract.Transact(opts, method, params...)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactoryCaller) AccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestAggregatedAccountFactory.contract.Call(opts, &out, "accountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactorySession) AccountImplementation() (common.Address, error) {
	return _TestAggregatedAccountFactory.Contract.AccountImplementation(&_TestAggregatedAccountFactory.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactoryCallerSession) AccountImplementation() (common.Address, error) {
	return _TestAggregatedAccountFactory.Contract.AccountImplementation(&_TestAggregatedAccountFactory.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactoryCaller) GetAddress(opts *bind.CallOpts, owner common.Address, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TestAggregatedAccountFactory.contract.Call(opts, &out, "getAddress", owner, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactorySession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _TestAggregatedAccountFactory.Contract.GetAddress(&_TestAggregatedAccountFactory.CallOpts, owner, salt)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactoryCallerSession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _TestAggregatedAccountFactory.Contract.GetAddress(&_TestAggregatedAccountFactory.CallOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactoryTransactor) CreateAccount(opts *bind.TransactOpts, owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _TestAggregatedAccountFactory.contract.Transact(opts, "createAccount", owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactorySession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _TestAggregatedAccountFactory.Contract.CreateAccount(&_TestAggregatedAccountFactory.TransactOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_TestAggregatedAccountFactory *TestAggregatedAccountFactoryTransactorSession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _TestAggregatedAccountFactory.Contract.CreateAccount(&_TestAggregatedAccountFactory.TransactOpts, owner, salt)
}
