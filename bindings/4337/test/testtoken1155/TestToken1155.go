// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testtoken1155

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

// TestToken1155MetaData contains all meta data concerning the TestToken1155 contract.
var TestToken1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60803461010e5760408101906001600160401b038211818310176100f857602091604052601d8152017f68747470733a2f2f6a757374666f72746573742f7b69647d2e6a736f6e000000815260025490600191600181811c911680156100ee575b60208210146100d857601f811161008f575b815162ffffff1916603a176002556040516110fb90816101148239f35b6002600052601f0160051c7f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace908101905b8181106100cd5750610072565b6000815583016100c0565b634e487b7160e01b600052602260045260246000fd5b90607f1690610060565b634e487b7160e01b600052604160045260246000fd5b600080fdfe6080604052600436101561001257600080fd5b6000803560e01c908162fdd58e146100a95750806301ffc9a7146100a45780630e89341c1461009f5780631b2ef1ca1461009a5780632eb2c2d6146100955780634e1273f414610090578063a22cb4651461008b578063e985e9c5146100865763f242432a1461008157600080fd5b610955565b6108f8565b61083b565b610780565b610612565b610319565b610210565b61014e565b346100f45760403660031901126100f4576100eb60209160406100ca6100f7565b916024358152808552209060018060a01b0316600052602052604060002090565b54604051908152f35b80fd5b600435906001600160a01b038216820361010d57565b600080fd5b602435906001600160a01b038216820361010d57565b35906001600160a01b038216820361010d57565b6001600160e01b031981160361010d57565b3461010d57602036600319011261010d57602060043561016d8161013c565b63ffffffff60e01b16636cdb3d1360e11b81149081156101ab575b811561019a575b506040519015158152f35b6301ffc9a760e01b1490503861018f565b6303a24d0760e21b81149150610188565b919082519283825260005b8481106101e8575050826000602080949584010152601f8019910116010190565b6020818301810151848301820152016101c7565b90602061020d9281815201906101bc565b90565b3461010d5760208060031936011261010d57604051906000906002546001918160011c926001831692831561030f575b6020851084146102fb5784875260208701939081156102dc5750600114610282575b61027e8661027281880382610515565b604051918291826101fc565b0390f35b6002600090815294509192917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace5b8386106102cb57505050910190506102728261027e38610262565b8054858701529482019481016102b0565b60ff1916845250505090151560051b0190506102728261027e38610262565b634e487b7160e01b86526022600452602486fd5b93607f1693610240565b3461010d5760408060031936011261010d57805190610337826104f4565b600080835233156104c8576103706024356004359160405192600184526020840152604083019160018352606084015260808301604052565b91909384518351908181036104a7575050815b85518110156103e3578060019160051b6103db6103d3602080848c0101519389010151926103bc33916000526000602052604060002090565b9060018060a01b0316600052602052604060002090565b918254610d55565b905501610383565b5091909392600184511460001461046c578460208501517fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62602085015193518061043d339633958360209093929193604081019481520152565b0390a45b82516001036104615760208061045e940151910151903333611059565b80f35b61045e923333610e5f565b518433917f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb33918061049f878a83610d78565b0390a4610441565b8551635b05999160e01b815260048101919091526024810191909152604490fd5b6024915190632bfa23e760e11b82526004820152fd5b634e487b7160e01b600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761051057604052565b6104de565b90601f8019910116810190811067ffffffffffffffff82111761051057604052565b67ffffffffffffffff81116105105760051b60200190565b9080601f8301121561010d57602090823561056981610537565b936105776040519586610515565b81855260208086019260051b82010192831161010d57602001905b8282106105a0575050505090565b81358152908301908301610592565b67ffffffffffffffff811161051057601f01601f191660200190565b81601f8201121561010d578035906105e2826105af565b926105f06040519485610515565b8284526020838301011161010d57816000926020809301838601378301015290565b3461010d5760a036600319011261010d5761062b6100f7565b610633610112565b906044359167ffffffffffffffff9081841161010d576106586004943690860161054f565b9060643583811161010d57610670903690870161054f565b9260843590811161010d5761068890369087016105cb565b936001600160a01b03808216903382141580610717575b6106ea578316156106d257156106bb576106b99550610b5c565b005b604051626a0d4560e21b8152600081880152602490fd5b604051632bfa23e760e11b8152600081890152602490fd5b6040805163711bec9160e11b815233818b019081526001600160a01b038616602082015290918291010390fd5b50600082815260016020908152604080832033845290915290205460ff161561069f565b90815180825260208080930193019160005b82811061075b575050505090565b83518552938101939281019260010161074d565b90602061020d92818152019061073b565b3461010d57604036600319011261010d5760043567ffffffffffffffff80821161010d573660238301121561010d5781600401356107bd81610537565b926107cb6040519485610515565b8184526020916024602086019160051b8301019136831161010d57602401905b828210610824578560243586811161010d5761027e9161081261081892369060040161054f565b90610a9d565b6040519182918261076f565b83809161083084610128565b8152019101906107eb565b3461010d57604036600319011261010d576108546100f7565b60243580151580820361010d576001600160a01b0383169283156108e0573360009081526001602090815260408083206001600160a01b039094168352929052209060ff801983541691161790557f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31604051806108db339482919091602081019215159052565b0390a3005b60405162ced3e160e81b815260006004820152602490fd5b3461010d57604036600319011261010d57602060ff6109496109186100f7565b610920610112565b6001600160a01b0391821660009081526001865260408082209290931681526020919091522090565b54166040519015158152f35b3461010d5760a036600319011261010d5761096e6100f7565b610976610112565b60843567ffffffffffffffff811161010d576109969036906004016105cb565b906001600160a01b03838116903382141580610a4f575b610a2857821615610a0f57156109f7576106b9926109ef6064356044359160405192600184526020840152604083019160018352606084015260808301604052565b929091610b5c565b604051626a0d4560e21b815260006004820152602490fd5b604051632bfa23e760e11b815260006004820152602490fd5b60405163711bec9160e11b81523360048201526001600160a01b0386166024820152604490fd5b50600082815260016020908152604080832033845290915290205460ff16156109ad565b8051821015610a875760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b91909180518351808203610b3a575050805190610ad2610abc83610537565b92610aca6040519485610515565b808452610537565b60209190601f1901368484013760005b8151811015610b3257600581901b8281018401519087018401516000908152602081815260408083206001600160a01b039094168352929052205460019190610b2b8287610a73565b5201610ae2565b509193505050565b604051635b05999160e01b815260048101919091526024810191909152604490fd5b949190918151845190818103610b3a57505060005b8251811015610c5f57600581901b83810160209081015191870101516001600160a01b03929186908a8516610bd9575b6001948216610bb4575b50505001610b71565b610bcf916103bc6103d3926000526000602052604060002090565b9055388581610bab565b9192939050610bf68a6103bc846000526000602052604060002090565b54838110610c285791879184600196959403610c208d6103bc856000526000602052604060002090565b559450610ba1565b6040516303dee4c560e01b81526001600160a01b038c16600482015260248101919091526044810184905260648101839052608490fd5b509491939290936001855114600014610d0a576020858101518382015160408051928352928201526001600160a01b03838116929086169133917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f6291a45b6001600160a01b038116610cd3575b5050505050565b8451600103610cf957602080610cef960151920151923361108e565b3880808080610ccc565b610d0594919233610f5a565b610cef565b6040516001600160a01b03828116919085169033907f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb9080610d4d888c83610d78565b0390a4610cbd565b91908201809211610d6257565b634e487b7160e01b600052601160045260246000fd5b9091610d8f61020d9360408452604084019061073b565b91602081840391015261073b565b9081602091031261010d575161020d8161013c565b92610de161020d9593610def9360018060a01b031686526000602087015260a0604087015260a086019061073b565b90848203606086015261073b565b9160808184039101526101bc565b939061020d9593610de191610def9460018060a01b03809216885216602087015260a0604087015260a086019061073b565b3d15610e5a573d90610e40826105af565b91610e4e6040519384610515565b82523d6000602084013e565b606090565b9293919093843b610e71575050505050565b602091610e94604051948593849363bc197c8160e01b9889865260048601610db2565b038160006001600160a01b0388165af160009181610f29575b50610eec5782610ebb610e2f565b8051919082610ee557604051632bfa23e760e11b81526001600160a01b0383166004820152602490fd5b9050602001fd5b6001600160e01b03191603610f0657503880808080610ccc565b604051632bfa23e760e11b81526001600160a01b03919091166004820152602490fd5b610f4c91925060203d602011610f53575b610f448183610515565b810190610d9d565b9038610ead565b503d610f3a565b939290949194853b610f6f575b505050505050565b610f92602093604051958694859463bc197c8160e01b998a875260048701610dfd565b038160006001600160a01b0388165af160009181610fd4575b50610fb95782610ebb610e2f565b6001600160e01b03191603610f065750388080808080610f67565b610fee91925060203d602011610f5357610f448183610515565b9038610fab565b909260a09261020d9594600180861b03168352600060208401526040830152606082015281608082015201906101bc565b919261020d95949160a094600180871b0380921685521660208401526040830152606082015281608082015201906101bc565b9293919093843b61106b575050505050565b602091610e94604051948593849363f23a6e6160e01b9889865260048601610ff5565b939290949194853b6110a257505050505050565b610f92602093604051958694859463f23a6e6160e01b998a87526004870161102656fea26469706673582212200adcf53d7adaa8577bc1e764be94f1ff727684a8378056917d238ba342df7e5d64736f6c63430008190033",
}

// TestToken1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use TestToken1155MetaData.ABI instead.
var TestToken1155ABI = TestToken1155MetaData.ABI

// TestToken1155Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestToken1155MetaData.Bin instead.
var TestToken1155Bin = TestToken1155MetaData.Bin

// DeployTestToken1155 deploys a new Ethereum contract, binding an instance of TestToken1155 to it.
func DeployTestToken1155(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestToken1155, error) {
	parsed, err := TestToken1155MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestToken1155Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestToken1155{TestToken1155Caller: TestToken1155Caller{contract: contract}, TestToken1155Transactor: TestToken1155Transactor{contract: contract}, TestToken1155Filterer: TestToken1155Filterer{contract: contract}}, nil
}

// TestToken1155 is an auto generated Go binding around an Ethereum contract.
type TestToken1155 struct {
	TestToken1155Caller     // Read-only binding to the contract
	TestToken1155Transactor // Write-only binding to the contract
	TestToken1155Filterer   // Log filterer for contract events
}

// TestToken1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestToken1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestToken1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestToken1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestToken1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestToken1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestToken1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestToken1155Session struct {
	Contract     *TestToken1155    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestToken1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestToken1155CallerSession struct {
	Contract *TestToken1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TestToken1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestToken1155TransactorSession struct {
	Contract     *TestToken1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestToken1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestToken1155Raw struct {
	Contract *TestToken1155 // Generic contract binding to access the raw methods on
}

// TestToken1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestToken1155CallerRaw struct {
	Contract *TestToken1155Caller // Generic read-only contract binding to access the raw methods on
}

// TestToken1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestToken1155TransactorRaw struct {
	Contract *TestToken1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestToken1155 creates a new instance of TestToken1155, bound to a specific deployed contract.
func NewTestToken1155(address common.Address, backend bind.ContractBackend) (*TestToken1155, error) {
	contract, err := bindTestToken1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestToken1155{TestToken1155Caller: TestToken1155Caller{contract: contract}, TestToken1155Transactor: TestToken1155Transactor{contract: contract}, TestToken1155Filterer: TestToken1155Filterer{contract: contract}}, nil
}

// NewTestToken1155Caller creates a new read-only instance of TestToken1155, bound to a specific deployed contract.
func NewTestToken1155Caller(address common.Address, caller bind.ContractCaller) (*TestToken1155Caller, error) {
	contract, err := bindTestToken1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestToken1155Caller{contract: contract}, nil
}

// NewTestToken1155Transactor creates a new write-only instance of TestToken1155, bound to a specific deployed contract.
func NewTestToken1155Transactor(address common.Address, transactor bind.ContractTransactor) (*TestToken1155Transactor, error) {
	contract, err := bindTestToken1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestToken1155Transactor{contract: contract}, nil
}

// NewTestToken1155Filterer creates a new log filterer instance of TestToken1155, bound to a specific deployed contract.
func NewTestToken1155Filterer(address common.Address, filterer bind.ContractFilterer) (*TestToken1155Filterer, error) {
	contract, err := bindTestToken1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestToken1155Filterer{contract: contract}, nil
}

// bindTestToken1155 binds a generic wrapper to an already deployed contract.
func bindTestToken1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestToken1155MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestToken1155 *TestToken1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestToken1155.Contract.TestToken1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestToken1155 *TestToken1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestToken1155.Contract.TestToken1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestToken1155 *TestToken1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestToken1155.Contract.TestToken1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestToken1155 *TestToken1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestToken1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestToken1155 *TestToken1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestToken1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestToken1155 *TestToken1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestToken1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_TestToken1155 *TestToken1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestToken1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_TestToken1155 *TestToken1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _TestToken1155.Contract.BalanceOf(&_TestToken1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_TestToken1155 *TestToken1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _TestToken1155.Contract.BalanceOf(&_TestToken1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_TestToken1155 *TestToken1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _TestToken1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_TestToken1155 *TestToken1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _TestToken1155.Contract.BalanceOfBatch(&_TestToken1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_TestToken1155 *TestToken1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _TestToken1155.Contract.BalanceOfBatch(&_TestToken1155.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_TestToken1155 *TestToken1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TestToken1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_TestToken1155 *TestToken1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _TestToken1155.Contract.IsApprovedForAll(&_TestToken1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_TestToken1155 *TestToken1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _TestToken1155.Contract.IsApprovedForAll(&_TestToken1155.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestToken1155 *TestToken1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TestToken1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestToken1155 *TestToken1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestToken1155.Contract.SupportsInterface(&_TestToken1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestToken1155 *TestToken1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestToken1155.Contract.SupportsInterface(&_TestToken1155.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_TestToken1155 *TestToken1155Caller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _TestToken1155.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_TestToken1155 *TestToken1155Session) Uri(arg0 *big.Int) (string, error) {
	return _TestToken1155.Contract.Uri(&_TestToken1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_TestToken1155 *TestToken1155CallerSession) Uri(arg0 *big.Int) (string, error) {
	return _TestToken1155.Contract.Uri(&_TestToken1155.CallOpts, arg0)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 tokenId, uint256 amount) returns()
func (_TestToken1155 *TestToken1155Transactor) Mint(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TestToken1155.contract.Transact(opts, "mint", tokenId, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 tokenId, uint256 amount) returns()
func (_TestToken1155 *TestToken1155Session) Mint(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TestToken1155.Contract.Mint(&_TestToken1155.TransactOpts, tokenId, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 tokenId, uint256 amount) returns()
func (_TestToken1155 *TestToken1155TransactorSession) Mint(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TestToken1155.Contract.Mint(&_TestToken1155.TransactOpts, tokenId, amount)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_TestToken1155 *TestToken1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _TestToken1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_TestToken1155 *TestToken1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _TestToken1155.Contract.SafeBatchTransferFrom(&_TestToken1155.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_TestToken1155 *TestToken1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _TestToken1155.Contract.SafeBatchTransferFrom(&_TestToken1155.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_TestToken1155 *TestToken1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _TestToken1155.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_TestToken1155 *TestToken1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _TestToken1155.Contract.SafeTransferFrom(&_TestToken1155.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_TestToken1155 *TestToken1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _TestToken1155.Contract.SafeTransferFrom(&_TestToken1155.TransactOpts, from, to, id, value, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TestToken1155 *TestToken1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TestToken1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TestToken1155 *TestToken1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TestToken1155.Contract.SetApprovalForAll(&_TestToken1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TestToken1155 *TestToken1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TestToken1155.Contract.SetApprovalForAll(&_TestToken1155.TransactOpts, operator, approved)
}

// TestToken1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TestToken1155 contract.
type TestToken1155ApprovalForAllIterator struct {
	Event *TestToken1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TestToken1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestToken1155ApprovalForAll)
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
		it.Event = new(TestToken1155ApprovalForAll)
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
func (it *TestToken1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestToken1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestToken1155ApprovalForAll represents a ApprovalForAll event raised by the TestToken1155 contract.
type TestToken1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_TestToken1155 *TestToken1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*TestToken1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TestToken1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TestToken1155ApprovalForAllIterator{contract: _TestToken1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_TestToken1155 *TestToken1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TestToken1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TestToken1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestToken1155ApprovalForAll)
				if err := _TestToken1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_TestToken1155 *TestToken1155Filterer) ParseApprovalForAll(log types.Log) (*TestToken1155ApprovalForAll, error) {
	event := new(TestToken1155ApprovalForAll)
	if err := _TestToken1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestToken1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the TestToken1155 contract.
type TestToken1155TransferBatchIterator struct {
	Event *TestToken1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *TestToken1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestToken1155TransferBatch)
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
		it.Event = new(TestToken1155TransferBatch)
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
func (it *TestToken1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestToken1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestToken1155TransferBatch represents a TransferBatch event raised by the TestToken1155 contract.
type TestToken1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_TestToken1155 *TestToken1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*TestToken1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestToken1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestToken1155TransferBatchIterator{contract: _TestToken1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_TestToken1155 *TestToken1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *TestToken1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestToken1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestToken1155TransferBatch)
				if err := _TestToken1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_TestToken1155 *TestToken1155Filterer) ParseTransferBatch(log types.Log) (*TestToken1155TransferBatch, error) {
	event := new(TestToken1155TransferBatch)
	if err := _TestToken1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestToken1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the TestToken1155 contract.
type TestToken1155TransferSingleIterator struct {
	Event *TestToken1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *TestToken1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestToken1155TransferSingle)
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
		it.Event = new(TestToken1155TransferSingle)
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
func (it *TestToken1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestToken1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestToken1155TransferSingle represents a TransferSingle event raised by the TestToken1155 contract.
type TestToken1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_TestToken1155 *TestToken1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*TestToken1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestToken1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestToken1155TransferSingleIterator{contract: _TestToken1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_TestToken1155 *TestToken1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *TestToken1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestToken1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestToken1155TransferSingle)
				if err := _TestToken1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_TestToken1155 *TestToken1155Filterer) ParseTransferSingle(log types.Log) (*TestToken1155TransferSingle, error) {
	event := new(TestToken1155TransferSingle)
	if err := _TestToken1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestToken1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the TestToken1155 contract.
type TestToken1155URIIterator struct {
	Event *TestToken1155URI // Event containing the contract specifics and raw log

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
func (it *TestToken1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestToken1155URI)
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
		it.Event = new(TestToken1155URI)
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
func (it *TestToken1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestToken1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestToken1155URI represents a URI event raised by the TestToken1155 contract.
type TestToken1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TestToken1155 *TestToken1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*TestToken1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TestToken1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &TestToken1155URIIterator{contract: _TestToken1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TestToken1155 *TestToken1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *TestToken1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TestToken1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestToken1155URI)
				if err := _TestToken1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TestToken1155 *TestToken1155Filterer) ParseURI(log types.Log) (*TestToken1155URI, error) {
	event := new(TestToken1155URI)
	if err := _TestToken1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
