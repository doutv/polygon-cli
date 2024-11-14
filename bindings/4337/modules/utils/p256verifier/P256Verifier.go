// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package p256verifier

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

// P256VerifierMetaData contains all meta data concerning the P256Verifier contract.
var P256VerifierMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]",
	Bin: "0x60808060405234601557610b32908161001b8239f35b600080fdfe60c06040523461001a576100123661008f565b602081519101f35b600080fd5b6040810190811067ffffffffffffffff82111761003b57604052565b634e487b7160e01b600052604160045260246000fd5b60e0810190811067ffffffffffffffff82111761003b57604052565b90601f8019910116810190811067ffffffffffffffff82111761003b57604052565b60a08103610142578060201161001a57600060409180831161013e578060601161013e578060801161013e5760a01161013b57815182810181811067ffffffffffffffff82111761012757906100fa918452606035815260803560208201528335602035843561015a565b1561011e575060ff6001915b51911660208201526020815261011b8161001f565b90565b60ff9091610106565b634e487b7160e01b83526041600452602483fd5b80fd5b5080fd5b50604051600060208201526020815261011b8161001f565b9092831580156102ec575b80156102e4575b80156102bd575b6102b457805190602061018b81830193845190610313565b156102a957604051948186019082825282604088015282606088015260808701527fffffffff00000000ffffffffffffffffbce6faada7179e84f3b9cac2fc63254f60a08701527bffffffff00000000000000004319055258e8617b0c46353d039cdaae19958660c082015260c0815261020481610051565b600080928192519060055afa903d156102a1573d9167ffffffffffffffff831161028d576040519261023f601f8201601f191686018561006d565b83523d828585013e5b1561027957828280518101031261013b57500151905161027493929185908181890994099151906103cd565b061490565b634e487b7160e01b81526001600452602490fd5b634e487b7160e01b82526041600452602482fd5b606091610248565b505050505050600090565b50505050600090565b507bffffffff00000000000000004319055258e8617b0c46353d039cdaae19831015610173565b50821561016c565b507bffffffff00000000000000004319055258e8617b0c46353d039cdaae19841015610165565b600160601b63ffffffff60c01b03199081811080159061039b575b801561038a575b610382577f5ac635d8aa3a93e7b3ebbd55769886bc651d06b0cc53b0f63bce3c3e27d2604b8282818080956003600160601b0363ffffffff60c01b03190991818180090908089180091490565b505050600090565b508015801561033557508215610335565b508183101561032e565b600160ff1b81146103b7576000190190565b634e487b7160e01b600052601160045260246000fd5b90926080526000906001938493600092811580610586575b61057a576103f3838261064a565b95909460ff60a05260005b600060a0511215610551575b808a0361052057505050600080516020610add83398151915295600080516020610abd833981519152969593919492955b600060a051121561046a5750505050505050610465600163ffffffff60601b0360601b1992610590565b900990565b61047b9395979891929496986106b7565b80989398979291979188998192819a6104996080518960a05161069c565b6104a460a0516103a5565b60a052806104bd5750505050505b95939194929561043b565b92955092959a5092506104fd959a50600181146000146105075750600080516020610add83398151915293600080516020610abd833981519152936108b3565b90979296916104b2565b60020361051757859389936108b3565b879387936108b3565b6002810361053a575050508295819695939194929561043b565b919791600219016104b257509550849584966104b2565b506105616080518560a05161069c565b61056c60a0516103a5565b60a05280156103fe5761040a565b50505050505050600090565b50608051156103e5565b604051906020918281019183835283604083015283606083015260808201526002600160601b0363ffffffff60c01b031960a0820152600163ffffffff60601b0360601b1960c082015260c081526105e781610051565b600080928192519060055afa903d15610642573d9167ffffffffffffffff831161028d5760405192610622601f8201601f191686018561006d565b83523d828585013e5b1561027957828280518101031261013b5750015190565b60609161062b565b801580610694575b6106705761066c916106639161076b565b929190916109c9565b9091565b5050600080516020610abd83398151915290600080516020610add83398151915290565b508115610652565b91906002600192841c831b16921c1681018091116103b75790565b939092821580610763575b61075457600160601b63ffffffff60c01b0319908185600209948280878009809709948380888a0998818080808680096003600160601b0363ffffffff60c01b0319099280096003090884808a6002600160601b0363ffffffff60c01b031909818380090898898603918683116103b757888703908782116103b7578780969481809681950994089009089609930990565b50600093508392508291508190565b5080156106c2565b91908215806108ab575b156107a35750600080516020610abd8339815191529150600080516020610add833981519152906001908190565b7fb01cbd1c01e58065711814b583f061e9d431cca994cea1313449bf97c840ae0a91600160601b63ffffffff60c01b0319808481600186090894817f94e82e0c1ed3bdb90743191a9c5bbf0d88fc827fd214cc5f0b5ec6ba27673d6981600184090893841561087c57505080808480099384099481846001099482808860010995600080516020610abd83398151915209918784038481116103b75784908180866002600160601b0363ffffffff60c01b03190991818580090808978885038581116103b7578580949281930994080908935b93929190565b9350935050921560001461089c5761089391610a17565b91939092610876565b50506000806000926000610876565b508015610775565b919495929390958115806109c1575b156108e35750508315806108db575b6107545793929190565b5082156108d1565b859192949515806109b9575b6109a957600160601b63ffffffff60c01b0319968703918783116103b75787838189850908938689038981116103b7578990818484090892831561098e575050818880959493928180848196099b8c9485099b8c920999099609918784038481116103b75784908180866002600160601b0363ffffffff60c01b03190991818580090808978885038581116103b7578580949281930994080908929190565b965096505050509093501560001461089c5761089391610a17565b9550509150915091906001908190565b5085156108ef565b5080156108c2565b909392821580610a0f575b610a02576109e190610590565b918291600163ffffffff60601b0360601b1980809581940980099009930990565b5050509050600090600090565b5080156109d4565b919091801580610ab4575b610aa757600160601b63ffffffff60c01b031990818460020991808084800980940991816003600160601b0363ffffffff60c01b031981808088860994800960030908958280836002600160601b0363ffffffff60c01b031909818980090896878403918483116103b757858503928584116103b75785809492819309940890090892565b5060009150819081908190565b508215610a2256fe6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c2964fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f5a2646970667358221220366ec21c275fde0bf3f0f1b2321e4518ba6d8a71cd937785d624a4b95ee5d10964736f6c63430008190033",
}

// P256VerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use P256VerifierMetaData.ABI instead.
var P256VerifierABI = P256VerifierMetaData.ABI

// P256VerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use P256VerifierMetaData.Bin instead.
var P256VerifierBin = P256VerifierMetaData.Bin

// DeployP256Verifier deploys a new Ethereum contract, binding an instance of P256Verifier to it.
func DeployP256Verifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *P256Verifier, error) {
	parsed, err := P256VerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(P256VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &P256Verifier{P256VerifierCaller: P256VerifierCaller{contract: contract}, P256VerifierTransactor: P256VerifierTransactor{contract: contract}, P256VerifierFilterer: P256VerifierFilterer{contract: contract}}, nil
}

// P256Verifier is an auto generated Go binding around an Ethereum contract.
type P256Verifier struct {
	P256VerifierCaller     // Read-only binding to the contract
	P256VerifierTransactor // Write-only binding to the contract
	P256VerifierFilterer   // Log filterer for contract events
}

// P256VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type P256VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// P256VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type P256VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// P256VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type P256VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// P256VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type P256VerifierSession struct {
	Contract     *P256Verifier     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// P256VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type P256VerifierCallerSession struct {
	Contract *P256VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// P256VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type P256VerifierTransactorSession struct {
	Contract     *P256VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// P256VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type P256VerifierRaw struct {
	Contract *P256Verifier // Generic contract binding to access the raw methods on
}

// P256VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type P256VerifierCallerRaw struct {
	Contract *P256VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// P256VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type P256VerifierTransactorRaw struct {
	Contract *P256VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewP256Verifier creates a new instance of P256Verifier, bound to a specific deployed contract.
func NewP256Verifier(address common.Address, backend bind.ContractBackend) (*P256Verifier, error) {
	contract, err := bindP256Verifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &P256Verifier{P256VerifierCaller: P256VerifierCaller{contract: contract}, P256VerifierTransactor: P256VerifierTransactor{contract: contract}, P256VerifierFilterer: P256VerifierFilterer{contract: contract}}, nil
}

// NewP256VerifierCaller creates a new read-only instance of P256Verifier, bound to a specific deployed contract.
func NewP256VerifierCaller(address common.Address, caller bind.ContractCaller) (*P256VerifierCaller, error) {
	contract, err := bindP256Verifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &P256VerifierCaller{contract: contract}, nil
}

// NewP256VerifierTransactor creates a new write-only instance of P256Verifier, bound to a specific deployed contract.
func NewP256VerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*P256VerifierTransactor, error) {
	contract, err := bindP256Verifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &P256VerifierTransactor{contract: contract}, nil
}

// NewP256VerifierFilterer creates a new log filterer instance of P256Verifier, bound to a specific deployed contract.
func NewP256VerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*P256VerifierFilterer, error) {
	contract, err := bindP256Verifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &P256VerifierFilterer{contract: contract}, nil
}

// bindP256Verifier binds a generic wrapper to an already deployed contract.
func bindP256Verifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := P256VerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_P256Verifier *P256VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _P256Verifier.Contract.P256VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_P256Verifier *P256VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _P256Verifier.Contract.P256VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_P256Verifier *P256VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _P256Verifier.Contract.P256VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_P256Verifier *P256VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _P256Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_P256Verifier *P256VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _P256Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_P256Verifier *P256VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _P256Verifier.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_P256Verifier *P256VerifierTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _P256Verifier.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_P256Verifier *P256VerifierSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _P256Verifier.Contract.Fallback(&_P256Verifier.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_P256Verifier *P256VerifierTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _P256Verifier.Contract.Fallback(&_P256Verifier.TransactOpts, calldata)
}
