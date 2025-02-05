// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zkpasskeyverifier

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

// ZkPasskeyVerifierMetaData contains all meta data concerning the ZkPasskeyVerifier contract.
var ZkPasskeyVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CommitmentInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[10]\",\"name\":\"pubkeyX\",\"type\":\"uint256[10]\"},{\"internalType\":\"uint256[10]\",\"name\":\"pubkeyY\",\"type\":\"uint256[10]\"},{\"internalType\":\"uint256[10]\",\"name\":\"txHash\",\"type\":\"uint256[10]\"}],\"name\":\"callVerify10Signatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[10]\",\"name\":\"pubkeyX\",\"type\":\"uint256[10]\"},{\"internalType\":\"uint256[10]\",\"name\":\"pubkeyY\",\"type\":\"uint256[10]\"},{\"internalType\":\"uint256[10]\",\"name\":\"txHash\",\"type\":\"uint256[10]\"}],\"name\":\"hash10Signatures\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[1]\",\"name\":\"input\",\"type\":\"uint256[1]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557610a4b908161001b8239f35b600080fdfe60806040818152600436101561001457600080fd5b600091823560e01c90816330a3a20e146102465750806343db3c72146101c05763d588b4591461004357600080fd5b346101bc576103c0806003193601126101b8576101443681116101b457610284913683116101b057366103c4116101b05760209450835192858401926004358452358585015235606084015260243560808401526101643560a08401526102a43560c084015260443560e0840152610184356101008401526102c4356101208401526064356101408401526101a4356101608401526102e4356101808401526084356101a08401526101c4356101c0840152610304356101e084015260a4356102008401526101e4356102208401526103243561024084015260c43561026084015261020435610280840152610344356102a084015260e4356102c0840152610224356102e08401526103643561030084015261010435610320840152610244356103408401526103843561036084015261012435610380840152610264356103a08401526103a43581840152825261019b82610409565b9051902090516001600160f81b039091168152f35b8480fd5b8380fd5b8280fd5b5080fd5b50346101bc576101a03660031901126101bc5736610104116101bc5736610144116101bc576101843681116101b857366101a312156101b8578151906102058261043c565b816101a4913683116102425760209550905b8282106102335750505061022a90610458565b90519015158152f35b81358152908501908501610217565b8580fd5b9050346101b8576105403660031901126101b85736610104116101b85736610144116101b8576101843681116101b4576102c43681116101b0576104049036821161024257366105441161024257602084019235835235848401523560608301526101a43560808301526102e43560a08301526104243560c08301526101c43560e083015261030435610100830152610444356101208301526101e4356101408301526103243561016083015261046435610180830152610204356101a0830152610344356101c0830152610484356101e083015261022435610200830152610364356102208301526104a43561024083015261024435610260830152610384356102808301526104c4356102a0830152610264356102c08301526103a4356102e08301526104e435610300830152610284356103208301526103c435610340830152610504356103608301526102a4356103808301526103e4356103a08301526103c0610524358184015282526103bd82610409565b60018060f81b039151902016815192602084019084821067ffffffffffffffff8311176103f557508252825260209161022a90610458565b634e487b7160e01b81526041600452602490fd5b6103e0810190811067ffffffffffffffff82111761042657604052565b634e487b7160e01b600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761042657604052565b906040908151926104688461043c565b60209081368637610104948451958387019281358452610124358789015260609388858101865160809060005b8a8282106109fe575050505003601f8019918281018c52011689019867ffffffffffffffff998181108b8211176104265789527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001918291519020068452875193888486377f032826efd3da14888756b1889f120cf1711b82a77c8d1422e6bcf0436ac1ce5e898601527f01662a6d31bc7ffa43d1a466cf157dd76ff27680f3936151df88a31e3d1b795c868601527f2011f1ea76efd601e525ca6f68d2bb07385711a85d135ce5410bd9a1a9f0da7060808601527f11c0cc9c7b6a10e9e9cc696945597e175321a256f42d571d0ccd602175908c8f60a086015261014492898460c0880137610100917f0a21766bc1fe523f08b885b4011fbae89777991c318ffd6127cd45ef256d3f7a83880152610120937f03d6ebd8b1d35d15507a3e0fc9ae20e20e7dc2f3ad0dde07b4c1b9c353b3cad685890152610140957f1d379ec314086e52b27f935e089d0f689b8035108d6e26cde35ec4212a406280878a0152610160977f1bccaa80fa28cc639c7f288bd6dc8eca093585ce54f82b92acb5928d6ee545e8898b0152610180998c818c8160085afa905116156109ed578d9e8e9b9c9d9e519b8c01908c821090821117610426578f528a903683116109e8578d90915b8383106109d8575090508e809250519b8c9481848701937f184932331b713e5635bb7e0443166753b1a3b9cb8a6cdb59e89bcc6f567e15c888528088019e8f7f2f535d862892b75efbb060b71ac0d80de3ed7709fbe9f07a6d435b6933a2fe92905280518652015196019586528d848082805a906006608092fa986000865260008952516080830199818b5288838888815a600790fa9210161691805a906006608092fa16957f0e79c9a617690548a2bd3f098c025014522c02817c98f4abe46ce803b31ebd2e84527f199155d53afcaf3058138da8194650172b48259bea110fce7a6571c583e0d048905251809652815a600790fa921016168a88805a906006608092fa166001169651955196156109c7577f13918ac57d51930fbbde9ef4edadd0f8452d95a48cc8d15a88a7ffc85a5d98e3927f1e87dfbc0e15687d9b98cad686ad8a3b2fbadb4e8bcb25ac146b46633ab16ac87f05b5a6e0242a1be7b69ea0aec06b80141100e461f7509299e101f6de6b4109559695937f021b8bfa93c806690419599de5f2014a99adef3c695823a7b3a0c3ea7a14d5c9847f236ad71ff63cf1e83eaf8f2a3b255c4acedccd5bb4cc40d08925bf0367b748569560049e9f519e8f378d01528b01528901528701528501527f276ba5104aa5e5d103ad8c64d1b4524bfb65c0bd0d61428347bd630f9647857b6101a08501527f1999c9b3857b55a30a5a9d57ff60ed8cc4672bbd8bd777325fd26c746fe628906101c08501527f144c1cf6a462d5494e52cc331eecf618d67eca3ed935df9592e32f3c13dd86eb6101e08501527f1c1bf941d3a1df4a0c7ab35e7cde95db0d8045262be4666879b07dd2754aff2a6102008501527f289aceba05f5758f8d639a9e1b78f53c18efe25acbbcb110c71d26a2c5ce1f7d6102208501526102408401526102608301527f0f81a3fbbf1fae092e1c314b75bb154515e82213991b95f2248cb0a8c968ceda6102808301527f2ef50caab497754de5e44560eb6413a585935b4ed7f0f9552ab49c25aa3fa5106102a08301527f294009d0a3889406bc9ba32f49694ed6e519f85bbbba77c5b7e22071b19d3ba76102c08301527f1b1f82c5746904927f255ffeedeeb55dcacbd0a176bc59290cc383dc62d4fe216102e0830152816103008160085afa90511690565b895163a54f8e2760e01b8152600490fd5b82358152918101918e9101610678565b600080fd5b8d516351d49ff760e11b8152600490fd5b835185528e9550938401939092019160010161049556fea2646970667358221220a27515001fc58a5d4c9137d4a761151399ff03e228872fe0ed2485de418f156c64736f6c63430008190033",
}

// ZkPasskeyVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkPasskeyVerifierMetaData.ABI instead.
var ZkPasskeyVerifierABI = ZkPasskeyVerifierMetaData.ABI

// ZkPasskeyVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZkPasskeyVerifierMetaData.Bin instead.
var ZkPasskeyVerifierBin = ZkPasskeyVerifierMetaData.Bin

// DeployZkPasskeyVerifier deploys a new Ethereum contract, binding an instance of ZkPasskeyVerifier to it.
func DeployZkPasskeyVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZkPasskeyVerifier, error) {
	parsed, err := ZkPasskeyVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZkPasskeyVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZkPasskeyVerifier{ZkPasskeyVerifierCaller: ZkPasskeyVerifierCaller{contract: contract}, ZkPasskeyVerifierTransactor: ZkPasskeyVerifierTransactor{contract: contract}, ZkPasskeyVerifierFilterer: ZkPasskeyVerifierFilterer{contract: contract}}, nil
}

// ZkPasskeyVerifier is an auto generated Go binding around an Ethereum contract.
type ZkPasskeyVerifier struct {
	ZkPasskeyVerifierCaller     // Read-only binding to the contract
	ZkPasskeyVerifierTransactor // Write-only binding to the contract
	ZkPasskeyVerifierFilterer   // Log filterer for contract events
}

// ZkPasskeyVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkPasskeyVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkPasskeyVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkPasskeyVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkPasskeyVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkPasskeyVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkPasskeyVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkPasskeyVerifierSession struct {
	Contract     *ZkPasskeyVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZkPasskeyVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkPasskeyVerifierCallerSession struct {
	Contract *ZkPasskeyVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ZkPasskeyVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkPasskeyVerifierTransactorSession struct {
	Contract     *ZkPasskeyVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ZkPasskeyVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkPasskeyVerifierRaw struct {
	Contract *ZkPasskeyVerifier // Generic contract binding to access the raw methods on
}

// ZkPasskeyVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkPasskeyVerifierCallerRaw struct {
	Contract *ZkPasskeyVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ZkPasskeyVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkPasskeyVerifierTransactorRaw struct {
	Contract *ZkPasskeyVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkPasskeyVerifier creates a new instance of ZkPasskeyVerifier, bound to a specific deployed contract.
func NewZkPasskeyVerifier(address common.Address, backend bind.ContractBackend) (*ZkPasskeyVerifier, error) {
	contract, err := bindZkPasskeyVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZkPasskeyVerifier{ZkPasskeyVerifierCaller: ZkPasskeyVerifierCaller{contract: contract}, ZkPasskeyVerifierTransactor: ZkPasskeyVerifierTransactor{contract: contract}, ZkPasskeyVerifierFilterer: ZkPasskeyVerifierFilterer{contract: contract}}, nil
}

// NewZkPasskeyVerifierCaller creates a new read-only instance of ZkPasskeyVerifier, bound to a specific deployed contract.
func NewZkPasskeyVerifierCaller(address common.Address, caller bind.ContractCaller) (*ZkPasskeyVerifierCaller, error) {
	contract, err := bindZkPasskeyVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkPasskeyVerifierCaller{contract: contract}, nil
}

// NewZkPasskeyVerifierTransactor creates a new write-only instance of ZkPasskeyVerifier, bound to a specific deployed contract.
func NewZkPasskeyVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkPasskeyVerifierTransactor, error) {
	contract, err := bindZkPasskeyVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkPasskeyVerifierTransactor{contract: contract}, nil
}

// NewZkPasskeyVerifierFilterer creates a new log filterer instance of ZkPasskeyVerifier, bound to a specific deployed contract.
func NewZkPasskeyVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkPasskeyVerifierFilterer, error) {
	contract, err := bindZkPasskeyVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkPasskeyVerifierFilterer{contract: contract}, nil
}

// bindZkPasskeyVerifier binds a generic wrapper to an already deployed contract.
func bindZkPasskeyVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZkPasskeyVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkPasskeyVerifier *ZkPasskeyVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkPasskeyVerifier.Contract.ZkPasskeyVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkPasskeyVerifier *ZkPasskeyVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkPasskeyVerifier.Contract.ZkPasskeyVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkPasskeyVerifier *ZkPasskeyVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkPasskeyVerifier.Contract.ZkPasskeyVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkPasskeyVerifier *ZkPasskeyVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkPasskeyVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkPasskeyVerifier *ZkPasskeyVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkPasskeyVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkPasskeyVerifier *ZkPasskeyVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkPasskeyVerifier.Contract.contract.Transact(opts, method, params...)
}

// CallVerify10Signatures is a free data retrieval call binding the contract method 0x30a3a20e.
//
// Solidity: function callVerify10Signatures(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[10] pubkeyX, uint256[10] pubkeyY, uint256[10] txHash) view returns(bool success)
func (_ZkPasskeyVerifier *ZkPasskeyVerifierCaller) CallVerify10Signatures(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, pubkeyX [10]*big.Int, pubkeyY [10]*big.Int, txHash [10]*big.Int) (bool, error) {
	var out []interface{}
	err := _ZkPasskeyVerifier.contract.Call(opts, &out, "callVerify10Signatures", proof, commitments, commitmentPok, pubkeyX, pubkeyY, txHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CallVerify10Signatures is a free data retrieval call binding the contract method 0x30a3a20e.
//
// Solidity: function callVerify10Signatures(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[10] pubkeyX, uint256[10] pubkeyY, uint256[10] txHash) view returns(bool success)
func (_ZkPasskeyVerifier *ZkPasskeyVerifierSession) CallVerify10Signatures(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, pubkeyX [10]*big.Int, pubkeyY [10]*big.Int, txHash [10]*big.Int) (bool, error) {
	return _ZkPasskeyVerifier.Contract.CallVerify10Signatures(&_ZkPasskeyVerifier.CallOpts, proof, commitments, commitmentPok, pubkeyX, pubkeyY, txHash)
}

// CallVerify10Signatures is a free data retrieval call binding the contract method 0x30a3a20e.
//
// Solidity: function callVerify10Signatures(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[10] pubkeyX, uint256[10] pubkeyY, uint256[10] txHash) view returns(bool success)
func (_ZkPasskeyVerifier *ZkPasskeyVerifierCallerSession) CallVerify10Signatures(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, pubkeyX [10]*big.Int, pubkeyY [10]*big.Int, txHash [10]*big.Int) (bool, error) {
	return _ZkPasskeyVerifier.Contract.CallVerify10Signatures(&_ZkPasskeyVerifier.CallOpts, proof, commitments, commitmentPok, pubkeyX, pubkeyY, txHash)
}

// Hash10Signatures is a free data retrieval call binding the contract method 0xd588b459.
//
// Solidity: function hash10Signatures(uint256[10] pubkeyX, uint256[10] pubkeyY, uint256[10] txHash) pure returns(bytes32)
func (_ZkPasskeyVerifier *ZkPasskeyVerifierCaller) Hash10Signatures(opts *bind.CallOpts, pubkeyX [10]*big.Int, pubkeyY [10]*big.Int, txHash [10]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ZkPasskeyVerifier.contract.Call(opts, &out, "hash10Signatures", pubkeyX, pubkeyY, txHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash10Signatures is a free data retrieval call binding the contract method 0xd588b459.
//
// Solidity: function hash10Signatures(uint256[10] pubkeyX, uint256[10] pubkeyY, uint256[10] txHash) pure returns(bytes32)
func (_ZkPasskeyVerifier *ZkPasskeyVerifierSession) Hash10Signatures(pubkeyX [10]*big.Int, pubkeyY [10]*big.Int, txHash [10]*big.Int) ([32]byte, error) {
	return _ZkPasskeyVerifier.Contract.Hash10Signatures(&_ZkPasskeyVerifier.CallOpts, pubkeyX, pubkeyY, txHash)
}

// Hash10Signatures is a free data retrieval call binding the contract method 0xd588b459.
//
// Solidity: function hash10Signatures(uint256[10] pubkeyX, uint256[10] pubkeyY, uint256[10] txHash) pure returns(bytes32)
func (_ZkPasskeyVerifier *ZkPasskeyVerifierCallerSession) Hash10Signatures(pubkeyX [10]*big.Int, pubkeyY [10]*big.Int, txHash [10]*big.Int) ([32]byte, error) {
	return _ZkPasskeyVerifier.Contract.Hash10Signatures(&_ZkPasskeyVerifier.CallOpts, pubkeyX, pubkeyY, txHash)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43db3c72.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[1] input) view returns(bool success)
func (_ZkPasskeyVerifier *ZkPasskeyVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [1]*big.Int) (bool, error) {
	var out []interface{}
	err := _ZkPasskeyVerifier.contract.Call(opts, &out, "verifyProof", proof, commitments, commitmentPok, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x43db3c72.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[1] input) view returns(bool success)
func (_ZkPasskeyVerifier *ZkPasskeyVerifierSession) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _ZkPasskeyVerifier.Contract.VerifyProof(&_ZkPasskeyVerifier.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43db3c72.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[1] input) view returns(bool success)
func (_ZkPasskeyVerifier *ZkPasskeyVerifierCallerSession) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _ZkPasskeyVerifier.Contract.VerifyProof(&_ZkPasskeyVerifier.CallOpts, proof, commitments, commitmentPok, input)
}
