// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testaggregatedaccount

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

// TestAggregatedAccountMetaData contains all meta data concerning the TestAggregatedAccount contract.
var TestAggregatedAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"anAggregator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SimpleAccountInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e03461016457611304906001600160401b03601f38849003908101601f19168301908282118483101761016957808491604096879485528339810103126101645781516001600160a01b0392838216820361016457602001519283168303610164573060805260a0527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff82861c1661015357808083160361010f575b50505060c052516111849081610180823960805181818161072d0152610866015260a051818181610280015281816104fd0152818161059f0152818161097a01528181610ab901528181610c8001528181610e2e0152610ed4015260c0518181816102bd01526103630152f35b6001600160401b0319909116811790915582519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a13880806100a2565b845163f92ee8a960e01b8152600490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c806301ffc9a71461013b578063150b7a021461013657806319822f7c14610131578063245a7bfc1461012c57806347e1da2a146101275780634a58db19146101225780634d44560d1461011d5780634f1ef2861461011857806352d1902d146101135780638da5cb5b1461010e578063ad3cb1cc14610109578063b0d691fe14610104578063b61d27f6146100ff578063bc197c81146100fa578063c399ec88146100f5578063c4d66de8146100f0578063d087d288146100eb5763f23a6e610361000e57610cec565b610c4d565b610b1a565b610a8d565b6109fe565b6109a9565b610964565b6108e7565b6108be565b610853565b6106da565b61056e565b6104ee565b6103c2565b61034d565b610246565b6101ec565b346101a95760203660031901126101a95760043563ffffffff60e01b81168091036101a957602090630a85bd0160e11b8114908115610198575b8115610187575b506040519015158152f35b6301ffc9a760e01b1490503861017c565b630271189760e51b81149150610175565b600080fd5b6001600160a01b038116036101a957565b9181601f840112156101a9578235916001600160401b0383116101a957602083818601950101116101a957565b346101a95760803660031901126101a9576102086004356101ae565b6102136024356101ae565b6064356001600160401b0381116101a9576102329036906004016101bf565b5050604051630a85bd0160e11b8152602090f35b346101a9576003196060368201126101a957600435906001600160401b0382116101a95761012091360301126101a9576001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811633036103085761030490600060408051926102bb84610619565b7f0000000000000000000000000000000000000000000000000000000000000000169283815282602082015201526102f4604435610eac565b6040519081529081906020820190565b0390f35b60405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e74000000006044820152606490fd5b346101a95760003660031901126101a9576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b9181601f840112156101a9578235916001600160401b0383116101a9576020808501948460051b0101116101a957565b346101a95760603660031901126101a9576001600160401b036004358181116101a9576103f3903690600401610392565b6024929192358281116101a95761040e903690600401610392565b926044359081116101a957610427903690600401610392565b939091610432610eca565b848414806104dd575b61044490610d46565b8161049157505060005b82811061045757005b8061048b61047061046b600194878a610d9e565b610db3565b61048561047e848988610dc0565b36916106a3565b90610f58565b0161044e565b91909460009493945b8581106104a357005b806104d76104b761046b6001948a87610d9e565b6104c2838b89610d9e565b356104d161047e858b8a610dc0565b91610f80565b0161049a565b5081158061043b575081851461043b565b60008060031936011261056b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681813b1561056b5760405163b760faf960e01b8152306004820152918290602490829034905af180156105665761055a575080f35b61056390610639565b80f35b610e01565b80fd5b346101a9576000604036600319011261056b5760043561058d816101ae565b610595610f97565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b156105ff5760449083604051958694859363040b850f60e31b855216600484015260243560248401525af180156105665761055a575080f35b8280fd5b634e487b7160e01b600052604160045260246000fd5b606081019081106001600160401b0382111761063457604052565b610603565b6001600160401b03811161063457604052565b604081019081106001600160401b0382111761063457604052565b90601f801991011681019081106001600160401b0382111761063457604052565b6001600160401b03811161063457601f01601f191660200190565b9291926106af82610688565b916106bd6040519384610667565b8294818452818301116101a9578281602093846000960137010152565b60403660031901126101a95760048035906106f4826101ae565b6024356001600160401b0381116101a957366023820112156101a95761072390369060248185013591016106a3565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116308114908115610837575b5061082657906020839261076b610f97565b6040516352d1902d60e01b8152938491829088165afa600092816107f5575b506107b8575050604051634c9c8ce360e01b81526001600160a01b0390921690820190815281906020010390fd5b838360008051602061110f83398151915284036107d9576100198383610fee565b604051632a87526960e21b815290810184815281906020010390fd5b61081891935060203d60201161081f575b6108108183610667565b810190610e0d565b913861078a565b503d610806565b60405163703e46dd60e11b81528390fd5b90508160008051602061110f8339815191525416141538610759565b346101a95760003660031901126101a9577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036108ac57602060405160008051602061110f8339815191528152f35b60405163703e46dd60e11b8152600490fd5b346101a95760003660031901126101a9576000546040516001600160a01b039091168152602090f35b346101a95760003660031901126101a95760408051906109068261064c565b60058252602090640352e302e360dc1b6020840152604051916020835283519182602085015260005b8381106109515784604081866000838284010152601f80199101168101030190f35b858101830151858201830152820161092f565b346101a95760003660031901126101a9576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101a95760603660031901126101a9576004356109c6816101ae565b604435906001600160401b0382116101a9576109f46109ec6100199336906004016101bf565b61047e610eca565b9060243590610f80565b346101a95760a03660031901126101a957610a1a6004356101ae565b610a256024356101ae565b6001600160401b036044358181116101a957610a45903690600401610392565b50506064358181116101a957610a5f903690600401610392565b50506084359081116101a957610a799036906004016101bf565b505060405163bc197c8160e01b8152602090f35b346101a95760003660031901126101a9576040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561056657602091600091610afd575b50604051908152f35b610b149150823d841161081f576108108183610667565b38610af4565b346101a95760203660031901126101a957610b366004356101ae565b60008051602061112f833981519152546001600160401b0360ff8260401c1615911680159081610c45575b6001149081610c3b575b159081610c32575b50610c205760008051602061112f833981519152805467ffffffffffffffff1916600117905580610bf6575b610ba7610e1c565b610bad57005b60008051602061112f833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b60008051602061112f833981519152805460ff60401b191668010000000000000000179055610b9f565b60405163f92ee8a960e01b8152600490fd5b90501538610b73565b303b159150610b6b565b829150610b61565b346101a95760003660031901126101a957604051631aab3f0d60e11b8152306004820152600060248201526020816044817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156105665761030491600091610ccd575b506040519081529081906020820190565b610ce6915060203d60201161081f576108108183610667565b38610cbc565b346101a95760a03660031901126101a957610d086004356101ae565b610d136024356101ae565b6084356001600160401b0381116101a957610d329036906004016101bf565b505060405163f23a6e6160e01b8152602090f35b15610d4d57565b60405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b6044820152606490fd5b634e487b7160e01b600052603260045260246000fd5b9190811015610dae5760051b0190565b610d88565b35610dbd816101ae565b90565b9190811015610dae5760051b81013590601e19813603018212156101a95701908135916001600160401b0383116101a95760200182360381136101a9579190565b6040513d6000823e3d90fd5b908160209103126101a9575190565b600080546001600160a01b03191681557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de8280a3565b3d15610ea7573d90610e8d82610688565b91610e9b6040519384610667565b82523d6000602084013e565b606090565b80610eb45750565b600080808093338219f150610ec7610e7c565b50565b60018060a01b03807f0000000000000000000000000000000000000000000000000000000000000000163314908115610f4a575b5015610f0657565b606460405162461bcd60e51b815260206004820152602060248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152fd5b905060005416331438610efe565b600091829182602083519301915af1610f6f610e7c565b9015610f785750565b602081519101fd5b916000928392602083519301915af1610f6f610e7c565b6000546001600160a01b031633148015610fe5575b15610fb357565b60405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b6044820152606490fd5b50303314610fac565b90813b156110715760008051602061110f83398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a280511561105657610ec791611092565b50503461105f57565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b600080610dbd93602081519101845af46110aa610e7c565b91906110d257508051156110c057805190602001fd5b60405163d6bda27560e01b8152600490fd5b81511580611105575b6110e3575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b156110db56fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122062d5b31babf6efc3b6725d8577215fd8f922d9b7b6ef93d8a1156c740f4c562e64736f6c63430008190033",
}

// TestAggregatedAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use TestAggregatedAccountMetaData.ABI instead.
var TestAggregatedAccountABI = TestAggregatedAccountMetaData.ABI

// TestAggregatedAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestAggregatedAccountMetaData.Bin instead.
var TestAggregatedAccountBin = TestAggregatedAccountMetaData.Bin

// DeployTestAggregatedAccount deploys a new Ethereum contract, binding an instance of TestAggregatedAccount to it.
func DeployTestAggregatedAccount(auth *bind.TransactOpts, backend bind.ContractBackend, anEntryPoint common.Address, anAggregator common.Address) (common.Address, *types.Transaction, *TestAggregatedAccount, error) {
	parsed, err := TestAggregatedAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestAggregatedAccountBin), backend, anEntryPoint, anAggregator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestAggregatedAccount{TestAggregatedAccountCaller: TestAggregatedAccountCaller{contract: contract}, TestAggregatedAccountTransactor: TestAggregatedAccountTransactor{contract: contract}, TestAggregatedAccountFilterer: TestAggregatedAccountFilterer{contract: contract}}, nil
}

// TestAggregatedAccount is an auto generated Go binding around an Ethereum contract.
type TestAggregatedAccount struct {
	TestAggregatedAccountCaller     // Read-only binding to the contract
	TestAggregatedAccountTransactor // Write-only binding to the contract
	TestAggregatedAccountFilterer   // Log filterer for contract events
}

// TestAggregatedAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestAggregatedAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAggregatedAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestAggregatedAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAggregatedAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestAggregatedAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAggregatedAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestAggregatedAccountSession struct {
	Contract     *TestAggregatedAccount // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestAggregatedAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestAggregatedAccountCallerSession struct {
	Contract *TestAggregatedAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TestAggregatedAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestAggregatedAccountTransactorSession struct {
	Contract     *TestAggregatedAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TestAggregatedAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestAggregatedAccountRaw struct {
	Contract *TestAggregatedAccount // Generic contract binding to access the raw methods on
}

// TestAggregatedAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestAggregatedAccountCallerRaw struct {
	Contract *TestAggregatedAccountCaller // Generic read-only contract binding to access the raw methods on
}

// TestAggregatedAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestAggregatedAccountTransactorRaw struct {
	Contract *TestAggregatedAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestAggregatedAccount creates a new instance of TestAggregatedAccount, bound to a specific deployed contract.
func NewTestAggregatedAccount(address common.Address, backend bind.ContractBackend) (*TestAggregatedAccount, error) {
	contract, err := bindTestAggregatedAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestAggregatedAccount{TestAggregatedAccountCaller: TestAggregatedAccountCaller{contract: contract}, TestAggregatedAccountTransactor: TestAggregatedAccountTransactor{contract: contract}, TestAggregatedAccountFilterer: TestAggregatedAccountFilterer{contract: contract}}, nil
}

// NewTestAggregatedAccountCaller creates a new read-only instance of TestAggregatedAccount, bound to a specific deployed contract.
func NewTestAggregatedAccountCaller(address common.Address, caller bind.ContractCaller) (*TestAggregatedAccountCaller, error) {
	contract, err := bindTestAggregatedAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestAggregatedAccountCaller{contract: contract}, nil
}

// NewTestAggregatedAccountTransactor creates a new write-only instance of TestAggregatedAccount, bound to a specific deployed contract.
func NewTestAggregatedAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*TestAggregatedAccountTransactor, error) {
	contract, err := bindTestAggregatedAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestAggregatedAccountTransactor{contract: contract}, nil
}

// NewTestAggregatedAccountFilterer creates a new log filterer instance of TestAggregatedAccount, bound to a specific deployed contract.
func NewTestAggregatedAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*TestAggregatedAccountFilterer, error) {
	contract, err := bindTestAggregatedAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestAggregatedAccountFilterer{contract: contract}, nil
}

// bindTestAggregatedAccount binds a generic wrapper to an already deployed contract.
func bindTestAggregatedAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestAggregatedAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAggregatedAccount *TestAggregatedAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAggregatedAccount.Contract.TestAggregatedAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAggregatedAccount *TestAggregatedAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.TestAggregatedAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAggregatedAccount *TestAggregatedAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.TestAggregatedAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAggregatedAccount *TestAggregatedAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAggregatedAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAggregatedAccount *TestAggregatedAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAggregatedAccount *TestAggregatedAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TestAggregatedAccount *TestAggregatedAccountCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestAggregatedAccount.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TestAggregatedAccount *TestAggregatedAccountSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TestAggregatedAccount.Contract.UPGRADEINTERFACEVERSION(&_TestAggregatedAccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TestAggregatedAccount *TestAggregatedAccountCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TestAggregatedAccount.Contract.UPGRADEINTERFACEVERSION(&_TestAggregatedAccount.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_TestAggregatedAccount *TestAggregatedAccountCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestAggregatedAccount.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_TestAggregatedAccount *TestAggregatedAccountSession) Aggregator() (common.Address, error) {
	return _TestAggregatedAccount.Contract.Aggregator(&_TestAggregatedAccount.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_TestAggregatedAccount *TestAggregatedAccountCallerSession) Aggregator() (common.Address, error) {
	return _TestAggregatedAccount.Contract.Aggregator(&_TestAggregatedAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestAggregatedAccount *TestAggregatedAccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestAggregatedAccount.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestAggregatedAccount *TestAggregatedAccountSession) EntryPoint() (common.Address, error) {
	return _TestAggregatedAccount.Contract.EntryPoint(&_TestAggregatedAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TestAggregatedAccount *TestAggregatedAccountCallerSession) EntryPoint() (common.Address, error) {
	return _TestAggregatedAccount.Contract.EntryPoint(&_TestAggregatedAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestAggregatedAccount *TestAggregatedAccountCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestAggregatedAccount.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestAggregatedAccount *TestAggregatedAccountSession) GetDeposit() (*big.Int, error) {
	return _TestAggregatedAccount.Contract.GetDeposit(&_TestAggregatedAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TestAggregatedAccount *TestAggregatedAccountCallerSession) GetDeposit() (*big.Int, error) {
	return _TestAggregatedAccount.Contract.GetDeposit(&_TestAggregatedAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_TestAggregatedAccount *TestAggregatedAccountCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestAggregatedAccount.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_TestAggregatedAccount *TestAggregatedAccountSession) GetNonce() (*big.Int, error) {
	return _TestAggregatedAccount.Contract.GetNonce(&_TestAggregatedAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_TestAggregatedAccount *TestAggregatedAccountCallerSession) GetNonce() (*big.Int, error) {
	return _TestAggregatedAccount.Contract.GetNonce(&_TestAggregatedAccount.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_TestAggregatedAccount *TestAggregatedAccountCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TestAggregatedAccount.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_TestAggregatedAccount *TestAggregatedAccountSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _TestAggregatedAccount.Contract.OnERC1155BatchReceived(&_TestAggregatedAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_TestAggregatedAccount *TestAggregatedAccountCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _TestAggregatedAccount.Contract.OnERC1155BatchReceived(&_TestAggregatedAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_TestAggregatedAccount *TestAggregatedAccountCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TestAggregatedAccount.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_TestAggregatedAccount *TestAggregatedAccountSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _TestAggregatedAccount.Contract.OnERC1155Received(&_TestAggregatedAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_TestAggregatedAccount *TestAggregatedAccountCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _TestAggregatedAccount.Contract.OnERC1155Received(&_TestAggregatedAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TestAggregatedAccount *TestAggregatedAccountCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TestAggregatedAccount.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TestAggregatedAccount *TestAggregatedAccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _TestAggregatedAccount.Contract.OnERC721Received(&_TestAggregatedAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TestAggregatedAccount *TestAggregatedAccountCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _TestAggregatedAccount.Contract.OnERC721Received(&_TestAggregatedAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestAggregatedAccount *TestAggregatedAccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestAggregatedAccount.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestAggregatedAccount *TestAggregatedAccountSession) Owner() (common.Address, error) {
	return _TestAggregatedAccount.Contract.Owner(&_TestAggregatedAccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestAggregatedAccount *TestAggregatedAccountCallerSession) Owner() (common.Address, error) {
	return _TestAggregatedAccount.Contract.Owner(&_TestAggregatedAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TestAggregatedAccount *TestAggregatedAccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TestAggregatedAccount.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TestAggregatedAccount *TestAggregatedAccountSession) ProxiableUUID() ([32]byte, error) {
	return _TestAggregatedAccount.Contract.ProxiableUUID(&_TestAggregatedAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TestAggregatedAccount *TestAggregatedAccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TestAggregatedAccount.Contract.ProxiableUUID(&_TestAggregatedAccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestAggregatedAccount *TestAggregatedAccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TestAggregatedAccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestAggregatedAccount *TestAggregatedAccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestAggregatedAccount.Contract.SupportsInterface(&_TestAggregatedAccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestAggregatedAccount *TestAggregatedAccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestAggregatedAccount.Contract.SupportsInterface(&_TestAggregatedAccount.CallOpts, interfaceId)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAggregatedAccount.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_TestAggregatedAccount *TestAggregatedAccountSession) AddDeposit() (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.AddDeposit(&_TestAggregatedAccount.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.AddDeposit(&_TestAggregatedAccount.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _TestAggregatedAccount.contract.Transact(opts, "execute", dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_TestAggregatedAccount *TestAggregatedAccountSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.Execute(&_TestAggregatedAccount.TransactOpts, dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactorSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.Execute(&_TestAggregatedAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactor) ExecuteBatch(opts *bind.TransactOpts, dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _TestAggregatedAccount.contract.Transact(opts, "executeBatch", dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_TestAggregatedAccount *TestAggregatedAccountSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.ExecuteBatch(&_TestAggregatedAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactorSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.ExecuteBatch(&_TestAggregatedAccount.TransactOpts, dest, value, arg2)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ) returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactor) Initialize(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _TestAggregatedAccount.contract.Transact(opts, "initialize", arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ) returns()
func (_TestAggregatedAccount *TestAggregatedAccountSession) Initialize(arg0 common.Address) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.Initialize(&_TestAggregatedAccount.TransactOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ) returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactorSession) Initialize(arg0 common.Address) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.Initialize(&_TestAggregatedAccount.TransactOpts, arg0)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TestAggregatedAccount.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TestAggregatedAccount *TestAggregatedAccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.UpgradeToAndCall(&_TestAggregatedAccount.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.UpgradeToAndCall(&_TestAggregatedAccount.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestAggregatedAccount *TestAggregatedAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestAggregatedAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestAggregatedAccount *TestAggregatedAccountSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.ValidateUserOp(&_TestAggregatedAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_TestAggregatedAccount *TestAggregatedAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.ValidateUserOp(&_TestAggregatedAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestAggregatedAccount.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_TestAggregatedAccount *TestAggregatedAccountSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.WithdrawDepositTo(&_TestAggregatedAccount.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.WithdrawDepositTo(&_TestAggregatedAccount.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAggregatedAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestAggregatedAccount *TestAggregatedAccountSession) Receive() (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.Receive(&_TestAggregatedAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestAggregatedAccount *TestAggregatedAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _TestAggregatedAccount.Contract.Receive(&_TestAggregatedAccount.TransactOpts)
}

// TestAggregatedAccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TestAggregatedAccount contract.
type TestAggregatedAccountInitializedIterator struct {
	Event *TestAggregatedAccountInitialized // Event containing the contract specifics and raw log

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
func (it *TestAggregatedAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestAggregatedAccountInitialized)
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
		it.Event = new(TestAggregatedAccountInitialized)
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
func (it *TestAggregatedAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestAggregatedAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestAggregatedAccountInitialized represents a Initialized event raised by the TestAggregatedAccount contract.
type TestAggregatedAccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TestAggregatedAccount *TestAggregatedAccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*TestAggregatedAccountInitializedIterator, error) {

	logs, sub, err := _TestAggregatedAccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TestAggregatedAccountInitializedIterator{contract: _TestAggregatedAccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TestAggregatedAccount *TestAggregatedAccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TestAggregatedAccountInitialized) (event.Subscription, error) {

	logs, sub, err := _TestAggregatedAccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestAggregatedAccountInitialized)
				if err := _TestAggregatedAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TestAggregatedAccount *TestAggregatedAccountFilterer) ParseInitialized(log types.Log) (*TestAggregatedAccountInitialized, error) {
	event := new(TestAggregatedAccountInitialized)
	if err := _TestAggregatedAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestAggregatedAccountSimpleAccountInitializedIterator is returned from FilterSimpleAccountInitialized and is used to iterate over the raw logs and unpacked data for SimpleAccountInitialized events raised by the TestAggregatedAccount contract.
type TestAggregatedAccountSimpleAccountInitializedIterator struct {
	Event *TestAggregatedAccountSimpleAccountInitialized // Event containing the contract specifics and raw log

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
func (it *TestAggregatedAccountSimpleAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestAggregatedAccountSimpleAccountInitialized)
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
		it.Event = new(TestAggregatedAccountSimpleAccountInitialized)
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
func (it *TestAggregatedAccountSimpleAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestAggregatedAccountSimpleAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestAggregatedAccountSimpleAccountInitialized represents a SimpleAccountInitialized event raised by the TestAggregatedAccount contract.
type TestAggregatedAccountSimpleAccountInitialized struct {
	EntryPoint common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSimpleAccountInitialized is a free log retrieval operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_TestAggregatedAccount *TestAggregatedAccountFilterer) FilterSimpleAccountInitialized(opts *bind.FilterOpts, entryPoint []common.Address, owner []common.Address) (*TestAggregatedAccountSimpleAccountInitializedIterator, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TestAggregatedAccount.contract.FilterLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TestAggregatedAccountSimpleAccountInitializedIterator{contract: _TestAggregatedAccount.contract, event: "SimpleAccountInitialized", logs: logs, sub: sub}, nil
}

// WatchSimpleAccountInitialized is a free log subscription operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_TestAggregatedAccount *TestAggregatedAccountFilterer) WatchSimpleAccountInitialized(opts *bind.WatchOpts, sink chan<- *TestAggregatedAccountSimpleAccountInitialized, entryPoint []common.Address, owner []common.Address) (event.Subscription, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TestAggregatedAccount.contract.WatchLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestAggregatedAccountSimpleAccountInitialized)
				if err := _TestAggregatedAccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
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
func (_TestAggregatedAccount *TestAggregatedAccountFilterer) ParseSimpleAccountInitialized(log types.Log) (*TestAggregatedAccountSimpleAccountInitialized, error) {
	event := new(TestAggregatedAccountSimpleAccountInitialized)
	if err := _TestAggregatedAccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestAggregatedAccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TestAggregatedAccount contract.
type TestAggregatedAccountUpgradedIterator struct {
	Event *TestAggregatedAccountUpgraded // Event containing the contract specifics and raw log

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
func (it *TestAggregatedAccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestAggregatedAccountUpgraded)
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
		it.Event = new(TestAggregatedAccountUpgraded)
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
func (it *TestAggregatedAccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestAggregatedAccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestAggregatedAccountUpgraded represents a Upgraded event raised by the TestAggregatedAccount contract.
type TestAggregatedAccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TestAggregatedAccount *TestAggregatedAccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TestAggregatedAccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TestAggregatedAccount.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TestAggregatedAccountUpgradedIterator{contract: _TestAggregatedAccount.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TestAggregatedAccount *TestAggregatedAccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TestAggregatedAccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TestAggregatedAccount.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestAggregatedAccountUpgraded)
				if err := _TestAggregatedAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TestAggregatedAccount *TestAggregatedAccountFilterer) ParseUpgraded(log types.Log) (*TestAggregatedAccountUpgraded, error) {
	event := new(TestAggregatedAccountUpgraded)
	if err := _TestAggregatedAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
