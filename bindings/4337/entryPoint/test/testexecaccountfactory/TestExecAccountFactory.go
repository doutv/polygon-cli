// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testexecaccountfactory

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

// TestExecAccountFactoryMetaData contains all meta data concerning the TestExecAccountFactory contract.
var TestExecAccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"accountImplementation\",\"outputs\":[{\"internalType\":\"contractTestExecAccount\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ret\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a03460c5576001600160401b0390601f611e3338819003918201601f19168301918483118484101760af5780849260209460405283398101031260c557516001600160a01b0381169081900360c5576040519161170f908184019081118482101760af576020928492610724843981520301906000f0801560a35760805260405161065990816100cb8239608051818181609e0152818161021b01526102d10152f35b6040513d6000823e3d90fd5b634e487b7160e01b600052604160045260246000fd5b600080fdfe608080604052600436101561001357600080fd5b600090813560e01c90816311464fbe1461008a575080635fbfb9cf1461007057638cb84e181461004257600080fd5b3461006d57602061005b610055366100d1565b90610262565b6040516001600160a01b039091168152f35b80fd5b503461006d57602061005b610084366100d1565b906101ae565b9050346100cd57816003193601126100cd577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b60409060031901126100f7576004356001600160a01b03811681036100f7579060243590565b600080fd5b6060810190811067ffffffffffffffff82111761011857604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff82111761011857604052565b60005b8381106101635750506000910152565b8181015183820152602001610153565b909160609260018060a01b03168252604060208301526101a28151809281604086015260208686019101610150565b601f01601f1916010190565b906101b98183610262565b803b61025c575060405163189acdbd60e31b60208201526001600160a01b0392831660248083019190915281526101ef816100fc565b604051906102d38083019183831067ffffffffffffffff841117610118578392610241926103518539867f00000000000000000000000000000000000000000000000000000000000000001690610173565b03906000f58015610250571690565b6040513d6000823e3d90fd5b91505090565b600b605591610305936102d360209061032b6103378360409687519061028a8387018361012e565b8582528282019561035187396102f760018060a01b039c8d92838c519163189acdbd60e31b88840152166024820152602481526102c6816100fc565b8b51928391878301957f00000000000000000000000000000000000000000000000000000000000000001686610173565b03601f19810183528261012e565b895195869361031c868601998a9251928391610150565b84019151809386840190610150565b0103808452018261012e565b5190208351938401528201523081520160ff815320169056fe60806040526102d38038038061001481610194565b92833981019060408183031261018f5780516001600160a01b03811680820361018f5760208381015190936001600160401b03821161018f570184601f8201121561018f5780519061006d610068836101cf565b610194565b9582875285838301011161018f57849060005b83811061017b57505060009186010152813b15610163577f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b03191682179055604051907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a28351156101455750600080848461012c96519101845af4903d1561013c573d61011c610068826101cf565b908152600081943d92013e6101ea565b505b6040516085908161024e8239f35b606092506101ea565b9250505034610154575061012e565b63b398979f60e01b8152600490fd5b60249060405190634c9c8ce360e01b82526004820152fd5b818101830151888201840152869201610080565b600080fd5b6040519190601f01601f191682016001600160401b038111838210176101b957604052565b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116101b957601f01601f191660200190565b9061021157508051156101ff57805190602001fd5b60405163d6bda27560e01b8152600490fd5b81511580610244575b610222575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561021a56fe60806040527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc54600090819081906001600160a01b0316368280378136915af43d82803e15604b573d90f35b3d90fdfea2646970667358221220fe8f61e3fa58bb8add11a833212ef9684476c5dbbd688eea5bcc0b0062abf4cc64736f6c63430008190033a2646970667358221220b6451f58b75930d4275a629c9f0104bd70ad2a26ca1ccbe89cacfd31682e39a264736f6c6343000819003360c034610142576001600160401b0390601f61170f38819003918201601f1916830191848311848410176101475780849260209460405283398101031261014257516001600160a01b0381168103610142573060805260a0527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166101305780808316036100eb575b6040516115b1908161015e82396080518181816106c201526107fb015260a051818181610297015281816104bb0152818161055d01528181610a4d01528181610b8c01528181610d5901528181611124015261125f0152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1388080610092565b60405163f92ee8a960e01b8152600490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c806301ffc9a71461013b578063150b7a021461013657806319822f7c1461013157806347e1da2a1461012c5780634a58db19146101275780634d44560d146101225780634f1ef2861461011d57806352d1902d146101185780638da5cb5b146101135780638dd7712f1461010e578063ad3cb1cc14610109578063b0d691fe14610104578063b61d27f6146100ff578063bc197c81146100fa578063c399ec88146100f5578063c4d66de8146100f0578063d087d288146100eb5763f23a6e610361000e57610dc4565b610d26565b610bed565b610b60565b610ad1565b610a7c565b610a37565b6109df565b61087c565b610853565b6107e8565b610680565b61052c565b6104ac565b610380565b610262565b6101f9565b346101a95760203660031901126101a95760043563ffffffff60e01b81168091036101a957602090630a85bd0160e11b8114908115610198575b8115610187575b506040519015158152f35b6301ffc9a760e01b1490503861017c565b630271189760e51b81149150610175565b600080fd5b6001600160a01b038116036101a957565b35906101ca826101ae565b565b9181601f840112156101a9578235916001600160401b0383116101a957602083818601950101116101a957565b346101a95760803660031901126101a9576102156004356101ae565b6102206024356101ae565b6064356001600160401b0381116101a95761023f9036906004016101cc565b5050604051630a85bd0160e11b8152602090f35b90816101209103126101a95790565b346101a95760603660031901126101a9576004356001600160401b0381116101a957610292903690600401610253565b6044357f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361030b576102d56102ed92602435906110a9565b90806102f1575b506040519081529081906020820190565b0390f35b600080808093338219f150610304610eee565b50386102dc565b60405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e74000000006044820152606490fd5b9181601f840112156101a9578235916001600160401b0383116101a9576020808501948460051b0101116101a957565b346101a95760603660031901126101a9576001600160401b036004358181116101a9576103b1903690600401610350565b6024929192358281116101a9576103cc903690600401610350565b926044359081116101a9576103e5903690600401610350565b9390916103f061111a565b8484148061049b575b61040290610e1e565b8161044f57505060005b82811061041557005b8061044961042e610429600194878a610e76565b610e8b565b61044361043c848988610ec7565b369161062b565b906111a8565b0161040c565b91909460009493945b85811061046157005b806104956104756104296001948a87610e76565b610480838b89610e76565b3561048f61043c858b8a610ec7565b916111d0565b01610458565b508115806103f957508185146103f9565b600080600319360112610529577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681813b156105295760405163b760faf960e01b8152306004820152918290602490829034905af1801561052457610518575080f35b610521906105d7565b80f35b610ee2565b80fd5b346101a957600060403660031901126105295760043561054b816101ae565b6105536111e7565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b156105bd5760449083604051958694859363040b850f60e31b855216600484015260243560248401525af1801561052457610518575080f35b8280fd5b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116105ea57604052565b6105c1565b90601f801991011681019081106001600160401b038211176105ea57604052565b6001600160401b0381116105ea57601f01601f191660200190565b92919261063782610610565b9161064560405193846105ef565b8294818452818301116101a9578281602093846000960137010152565b9080601f830112156101a95781602061067d9335910161062b565b90565b60403660031901126101a957600480359061069a826101ae565b6024356001600160401b0381116101a9576106b89036908301610662565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081163081149081156107cc575b506107bb5790602083926107006111e7565b6040516352d1902d60e01b8152938491829088165afa6000928161078a575b5061074d575050604051634c9c8ce360e01b81526001600160a01b0390921690820190815281906020010390fd5b838360008051602061153c833981519152840361076e5761001983836112a6565b604051632a87526960e21b815290810184815281906020010390fd5b6107ad91935060203d6020116107b4575b6107a581836105ef565b81019061109a565b913861071f565b503d61079b565b60405163703e46dd60e11b81528390fd5b90508160008051602061153c83398151915254161415386106ee565b346101a95760003660031901126101a9577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316300361084157602060405160008051602061153c8339815191528152f35b60405163703e46dd60e11b8152600490fd5b346101a95760003660031901126101a9576000546040516001600160a01b039091168152602090f35b346101a9576003196040368201126101a9576001600160401b03906004358281116101a9576108af903690600401610253565b6108b761111a565b6108c46060820182610e95565b9290836004116101a95760609382810161090f575b6040517fd3fddfd1276d1cc278f10907710a44474a32f917b2fcfa198f46ca7689215e2f908061090a888883610fb0565b0390a1005b8160409293949550019282840301126101a9576004810135610930816101ae565b60248201359485116101a95761096f60009360047fd3fddfd1276d1cc278f10907710a44474a32f917b2fcfa198f46ca7689215e2f9786950101610662565b80519160209091019083906001600160a01b03165af190610997610991610eee565b92610f1e565b3880806108d9565b919082519283825260005b8481106109cb575050826000602080949584010152601f8019910116010190565b6020818301810151848301820152016109aa565b346101a95760003660031901126101a95760405160408101908082106001600160401b038311176105ea576102ed9160405260058152640352e302e360dc1b602082015260405191829160208352602083019061099f565b346101a95760003660031901126101a9576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101a95760603660031901126101a957600435610a99816101ae565b604435906001600160401b0382116101a957610ac7610abf6100199336906004016101cc565b61043c61111a565b90602435906111d0565b346101a95760a03660031901126101a957610aed6004356101ae565b610af86024356101ae565b6001600160401b036044358181116101a957610b18903690600401610350565b50506064358181116101a957610b32903690600401610350565b50506084359081116101a957610b4c9036906004016101cc565b505060405163bc197c8160e01b8152602090f35b346101a95760003660031901126101a9576040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561052457602091600091610bd0575b50604051908152f35b610be79150823d84116107b4576107a581836105ef565b38610bc7565b346101a95760203660031901126101a957600435610c0a816101ae565b60008051602061155c83398151915254906001600160401b0360ff8360401c1615921680159081610d1e575b6001149081610d14575b159081610d0b575b50610cf95760008051602061155c833981519152805467ffffffffffffffff19166001179055610c7c9082610ccf5761123e565b610c8257005b60008051602061155c833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290806020810161090a565b60008051602061155c833981519152805460ff60401b19166801000000000000000017905561123e565b60405163f92ee8a960e01b8152600490fd5b90501538610c48565b303b159150610c40565b839150610c36565b346101a95760003660031901126101a957604051631aab3f0d60e11b8152306004820152600060248201526020816044817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015610524576102ed91600091610da557506040519081529081906020820190565b610dbe915060203d6020116107b4576107a581836105ef565b386102dc565b346101a95760a03660031901126101a957610de06004356101ae565b610deb6024356101ae565b6084356001600160401b0381116101a957610e0a9036906004016101cc565b505060405163f23a6e6160e01b8152602090f35b15610e2557565b60405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b6044820152606490fd5b634e487b7160e01b600052603260045260246000fd5b9190811015610e865760051b0190565b610e60565b3561067d816101ae565b903590601e19813603018212156101a957018035906001600160401b0382116101a9576020019181360383136101a957565b90821015610e8657610ede9160051b810190610e95565b9091565b6040513d6000823e3d90fd5b3d15610f19573d90610eff82610610565b91610f0d60405193846105ef565b82523d6000602084013e565b606090565b15610f2557565b60405162461bcd60e51b81526020600482015260116024820152701a5b9b995c8818d85b1b0819985a5b1959607a1b6044820152606490fd5b9035601e19823603018112156101a95701602081359101916001600160401b0382116101a95781360383136101a957565b908060209392818452848401376000828201840152601f01601f1916010190565b909161108c61100e61067d9460408552610fdd60408601610fd0836101bf565b6001600160a01b03169052565b6020810135606086015261107c610ff76040830183610f5e565b9390610120948560808a0152610160890191610f8f565b916110736110366110226060840184610f5e565b603f198b8803810160a08d01529691610f8f565b608083013560c08a015260a083013560e08a01526101009560c0840135878b015261106460e0850185610f5e565b91878c850301908c0152610f8f565b93810190610f5e565b9186840301610140870152610f8f565b91602081840391015261099f565b908160209103126101a9575190565b907f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52603c60002061110a61110160018060a01b03926110fb61043c856000541696610100810190610e95565b9061134d565b909291926113a9565b160361111557600090565b600190565b60018060a01b03807f000000000000000000000000000000000000000000000000000000000000000016331490811561119a575b501561115657565b606460405162461bcd60e51b815260206004820152602060248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152fd5b90506000541633143861114e565b600091829182602083519301915af16111bf610eee565b90156111c85750565b602081519101fd5b916000928392602083519301915af16111bf610eee565b6000546001600160a01b031633148015611235575b1561120357565b60405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b6044820152606490fd5b503033146111fc565b600080546001600160a01b0319166001600160a01b039283169081178255917f000000000000000000000000000000000000000000000000000000000000000016907f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de9080a3565b90813b1561132c5760008051602061153c83398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a28051156113115761130e91611436565b50565b50503461131a57565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b815191906041830361137e5761137792506020820151906060604084015193015160001a90611454565b9192909190565b505060009160029190565b6004111561139357565b634e487b7160e01b600052602160045260246000fd5b6113b281611389565b806113bb575050565b6113c481611389565b600181036113de5760405163f645eedf60e01b8152600490fd5b6113e781611389565b600281036114085760405163fce698f760e01b815260048101839052602490fd5b80611414600392611389565b1461141c5750565b6040516335e2f38360e21b81526004810191909152602490fd5b60008061067d93602081519101845af461144e610eee565b916114d8565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084116114cc57926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa156105245780516001600160a01b038116156114c357918190565b50809160019190565b50505060009160039190565b906114ff57508051156114ed57805190602001fd5b60405163d6bda27560e01b8152600490fd5b81511580611532575b611510575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561150856fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122096cbc01db3264920b2819047e9ea0f7bcde2fec41bba2fb0a697c3459794e82d64736f6c63430008190033",
}

// TestExecAccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use TestExecAccountFactoryMetaData.ABI instead.
var TestExecAccountFactoryABI = TestExecAccountFactoryMetaData.ABI

// TestExecAccountFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestExecAccountFactoryMetaData.Bin instead.
var TestExecAccountFactoryBin = TestExecAccountFactoryMetaData.Bin

// DeployTestExecAccountFactory deploys a new Ethereum contract, binding an instance of TestExecAccountFactory to it.
func DeployTestExecAccountFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *TestExecAccountFactory, error) {
	parsed, err := TestExecAccountFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestExecAccountFactoryBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestExecAccountFactory{TestExecAccountFactoryCaller: TestExecAccountFactoryCaller{contract: contract}, TestExecAccountFactoryTransactor: TestExecAccountFactoryTransactor{contract: contract}, TestExecAccountFactoryFilterer: TestExecAccountFactoryFilterer{contract: contract}}, nil
}

// TestExecAccountFactory is an auto generated Go binding around an Ethereum contract.
type TestExecAccountFactory struct {
	TestExecAccountFactoryCaller     // Read-only binding to the contract
	TestExecAccountFactoryTransactor // Write-only binding to the contract
	TestExecAccountFactoryFilterer   // Log filterer for contract events
}

// TestExecAccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestExecAccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExecAccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestExecAccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExecAccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestExecAccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExecAccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestExecAccountFactorySession struct {
	Contract     *TestExecAccountFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TestExecAccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestExecAccountFactoryCallerSession struct {
	Contract *TestExecAccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// TestExecAccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestExecAccountFactoryTransactorSession struct {
	Contract     *TestExecAccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// TestExecAccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestExecAccountFactoryRaw struct {
	Contract *TestExecAccountFactory // Generic contract binding to access the raw methods on
}

// TestExecAccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestExecAccountFactoryCallerRaw struct {
	Contract *TestExecAccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TestExecAccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestExecAccountFactoryTransactorRaw struct {
	Contract *TestExecAccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestExecAccountFactory creates a new instance of TestExecAccountFactory, bound to a specific deployed contract.
func NewTestExecAccountFactory(address common.Address, backend bind.ContractBackend) (*TestExecAccountFactory, error) {
	contract, err := bindTestExecAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestExecAccountFactory{TestExecAccountFactoryCaller: TestExecAccountFactoryCaller{contract: contract}, TestExecAccountFactoryTransactor: TestExecAccountFactoryTransactor{contract: contract}, TestExecAccountFactoryFilterer: TestExecAccountFactoryFilterer{contract: contract}}, nil
}

// NewTestExecAccountFactoryCaller creates a new read-only instance of TestExecAccountFactory, bound to a specific deployed contract.
func NewTestExecAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*TestExecAccountFactoryCaller, error) {
	contract, err := bindTestExecAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestExecAccountFactoryCaller{contract: contract}, nil
}

// NewTestExecAccountFactoryTransactor creates a new write-only instance of TestExecAccountFactory, bound to a specific deployed contract.
func NewTestExecAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TestExecAccountFactoryTransactor, error) {
	contract, err := bindTestExecAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestExecAccountFactoryTransactor{contract: contract}, nil
}

// NewTestExecAccountFactoryFilterer creates a new log filterer instance of TestExecAccountFactory, bound to a specific deployed contract.
func NewTestExecAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TestExecAccountFactoryFilterer, error) {
	contract, err := bindTestExecAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestExecAccountFactoryFilterer{contract: contract}, nil
}

// bindTestExecAccountFactory binds a generic wrapper to an already deployed contract.
func bindTestExecAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestExecAccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestExecAccountFactory *TestExecAccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestExecAccountFactory.Contract.TestExecAccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestExecAccountFactory *TestExecAccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExecAccountFactory.Contract.TestExecAccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestExecAccountFactory *TestExecAccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestExecAccountFactory.Contract.TestExecAccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestExecAccountFactory *TestExecAccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestExecAccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestExecAccountFactory *TestExecAccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExecAccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestExecAccountFactory *TestExecAccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestExecAccountFactory.Contract.contract.Transact(opts, method, params...)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_TestExecAccountFactory *TestExecAccountFactoryCaller) AccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestExecAccountFactory.contract.Call(opts, &out, "accountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_TestExecAccountFactory *TestExecAccountFactorySession) AccountImplementation() (common.Address, error) {
	return _TestExecAccountFactory.Contract.AccountImplementation(&_TestExecAccountFactory.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_TestExecAccountFactory *TestExecAccountFactoryCallerSession) AccountImplementation() (common.Address, error) {
	return _TestExecAccountFactory.Contract.AccountImplementation(&_TestExecAccountFactory.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_TestExecAccountFactory *TestExecAccountFactoryCaller) GetAddress(opts *bind.CallOpts, owner common.Address, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TestExecAccountFactory.contract.Call(opts, &out, "getAddress", owner, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_TestExecAccountFactory *TestExecAccountFactorySession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _TestExecAccountFactory.Contract.GetAddress(&_TestExecAccountFactory.CallOpts, owner, salt)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_TestExecAccountFactory *TestExecAccountFactoryCallerSession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _TestExecAccountFactory.Contract.GetAddress(&_TestExecAccountFactory.CallOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_TestExecAccountFactory *TestExecAccountFactoryTransactor) CreateAccount(opts *bind.TransactOpts, owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _TestExecAccountFactory.contract.Transact(opts, "createAccount", owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_TestExecAccountFactory *TestExecAccountFactorySession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _TestExecAccountFactory.Contract.CreateAccount(&_TestExecAccountFactory.TransactOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_TestExecAccountFactory *TestExecAccountFactoryTransactorSession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _TestExecAccountFactory.Contract.CreateAccount(&_TestExecAccountFactory.TransactOpts, owner, salt)
}
