// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package groth16verifier

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

// Groth16VerifierMetaData contains all meta data concerning the Groth16Verifier contract.
var Groth16VerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[17]\",\"name\":\"_pubSignals\",\"type\":\"uint256[17]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557610e9a908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b6000803560e01c63c1940a361461002857600080fd5b3461014957610320366003190112610149576100433661014c565b3660c411610145576100543661015e565b903661032411610141576020929161013d916104006040526100786101043561016c565b6100846101243561016c565b6100906101443561016c565b61009c6101643561016c565b6100a86101843561016c565b6100b46101a43561016c565b6100c06101c43561016c565b6100cc6101e43561016c565b6100d86102043561016c565b6100e46102243561016c565b6100f06102443561016c565b6100fc6102643561016c565b6101086102843561016c565b6101146102a43561016c565b6101206102c43561016c565b61012c6102e43561016c565b6101386103043561016c565b610ac8565b8152f35b8280fd5b5080fd5b80fd5b9060049160441161015957565b600080fd5b9060c4916101041161015957565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001111561019557565b6000805260206000f35b604051907f2102dc0100b48329323c53280ae877da738da899f1a8f56586b9f3db415c16dd82527f1e6ca1ca9b627b68a4ed21432296be6e47d01e5286f1f802e73933fdf20d59986020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f1ad26e20ca34ed5728c869bd385109c9d88cff8f70d597e58b24cdc5589d979382527f1a66523292584ffbb4e973b11e188aa265bff1eb9d1558c0395f7bec4218b1206020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f0b55aab344f6d78726a0386f705bfbf629e33633a1789a5998efe885bc169dbf82527f169a6763ca5222ce7419afc3477d52d5e349f3d42f215750223017891926359d6020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f267c94cb005b72e7eb6930a764358cbf86c60f3e2f6157f9555db70e409a1e7982527f10247c59bb633d5bee15924d64a496c74f57ceabcb7fd3532250d1f35ae2720b6020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f195d487b09fccf451a06d28284c22339ff6ef986aa12d22ffe90581610af896682527f06f47d6692d9be4c4621a06b7daabccf08d25267c57e93c2c3fea3737e1a6e946020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f22a4c17784e66a4c344cdfc3aa057c27fc9bebed331f8ba303319b3ba928de0d82527f15faba32b8caec4ebe23c31d173d7482673888e98eee9d6c58d3e710f4f8dde56020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f0bbe89bf4208ea0dc6ff2e46e45f6a8f37e425819879a8d5994b0bde9094c80d82527f10ce8e20eab9ebaf20bf9c0e37d0bf7221c9ccf6cad9daea9f3c443de3a637556020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f14eb893ac09c3d3aee54418c040dbe88bf783d348bd3b5c4518c7bf1d8aec26282527f141e361dad6e277d65be816699cefb0ca2d923479778144ee00f29c13213d77f6020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f0a605a76e88bf31a11123fc3e369e0df3963c9c028634c361f9119b8de1b990482527f0738cbcecd43296bbc96f1edda74f47d60e5c39c0c78490d85575cd0654f97cf6020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f229a614ec1f2b488436bdff722d3181ab858756cbd0d8e6998ba7ee69402eea982527f1d1fe191fd8e3a87259f813bc6d129e33c7103fd2a5581c5a5a653ebd58560fa6020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f11e0dfaa09b5f711e02bdc2aeceae6087ec2f51a67360136620e4d362569c57f82527f2b175f9bd2ea8a198e48bf884f6122c6cd5f158f65c81f0a87dfc36a9c8ad0806020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907e38a2dd1123eed9c46fc32c6d7b7bc76501a796451b8ef21a610685bd3a3d0e82527f0565ecea2cabb990246ef3f76d2c15c29534aceb54268944d08f376f2d3d95536020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f2ed0df32674716396b7a5c3d9bdba4fea3696946998ff49fdbcf5281a169279382527f163dd0255b92a820b32d59e099f73e201bdb8e571790d34ac6ba28349d7fca166020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f17261d4ffaffa7335907d4dbf6afa1645476e27e6d7c86aabb9ec21b2550a64882527f179e93baab96c3d37700ab1317864b5d71fef6e1dfbeb1aae743d8f0b5d76fff6020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f270cd23e87b08d7306de146775961bccded1850a4ec9ae5e09a51754ceed342582527f08990b55dad36cee232ba6496fb3200c96afda7dbc9bfdb4491889acec80710b6020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f051aef5ce2a503e20b64f08cbb8d40da692577307163775eea9cf47e96fc290082527f2c8dba9a0a180b5a808d12e7a5b3bd52ffaa313d443dfeba235e323e9d569a4c6020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b604051907f0db46c2a8077c2b9ee01ae292a3847bf9eb990cee9604de06518b7555071304282527f113a153506ad9fca85f4cf4e321d65d164535e506a77fb8662a86a09525352b76020830152604082019081526107cf196040836060816007855a01fa156101955760409260066080939284938451905260a05160608401525a01fa1561019557565b6020907f2b0fbae590d705323add300533843f3de53e877c016e08593f7c28fe75e890056080527f185f574e3ac0601a486d097729c8cefef767665b4d23fedb0642b4ed25ed729960a052610b1f6101043561019f565b610b2b61012435610229565b610b37610144356102b3565b610b436101643561033d565b610b4f610184356103c7565b610b5b6101a435610451565b610b676101c4356104db565b610b736101e435610565565b610b7f610204356105ef565b610b8b61022435610679565b610b9761024435610703565b610ba36102643561078d565b610baf61028435610816565b610bbb6102a4356108a0565b610bc76102c43561092a565b610bd36102e4356109b4565b610bdf61030435610a3e565b61010092818492358352837f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd479101358103066101205260443561014052606435610160526084356101805260a4356101a0527f2d4d9aa7e302d9df41749d5507949d05dbea33fbb16c643b22f599a2be6df2e26101c0527f14bedd503c37ceb061d8ec60209fe345ce89830a19230301f076caff004d19266101e0527f0967032fcbf776d1afc985f88877f182d38480a653f2decaa9794cbc3bf3060c610200527f0e187847ad4c798374d0d6732bf501847dd68bc0e071241e0213bc7fc13db7ab610220527f304cfbd1e08a704a99f5e847d93f8c3caafddec46b7a0d379da69a4d112346a7610240527f1739c1b1a457a8c7313123d24d2f9192f896b7c63eea05a9d57f06547ad0cec8610260526080516102805260a0516102a0527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26102c0527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6102e05282610300917f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b83527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa610320528035610340520135610360527f1d0156ab95798e5dd2b89e834e57a3d846902599b894244708ecc68dec978609610380527f05ff1213d7efbc6706d81737de79aac39ebc4405c92f4dc8d893c1bdb1c6778d6103a0527f10bdcc6e8a53c19d893744f44fa672bdae7a5cf0b9bf8772fd29a9fdfde8c00e6103c0527f183072ee893a7fd87845d43df7aa78bcf0bc80acb9c57eea66f415795ac34a556103e0528160086107cf195a01fa9051169056fea2646970667358221220ae399e57d197881e864b021b6fbbf5606875cddaabc28a38501381a7967ed3be64736f6c63430008190033",
}

// Groth16VerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use Groth16VerifierMetaData.ABI instead.
var Groth16VerifierABI = Groth16VerifierMetaData.ABI

// Groth16VerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Groth16VerifierMetaData.Bin instead.
var Groth16VerifierBin = Groth16VerifierMetaData.Bin

// DeployGroth16Verifier deploys a new Ethereum contract, binding an instance of Groth16Verifier to it.
func DeployGroth16Verifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Groth16Verifier, error) {
	parsed, err := Groth16VerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Groth16VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Groth16Verifier{Groth16VerifierCaller: Groth16VerifierCaller{contract: contract}, Groth16VerifierTransactor: Groth16VerifierTransactor{contract: contract}, Groth16VerifierFilterer: Groth16VerifierFilterer{contract: contract}}, nil
}

// Groth16Verifier is an auto generated Go binding around an Ethereum contract.
type Groth16Verifier struct {
	Groth16VerifierCaller     // Read-only binding to the contract
	Groth16VerifierTransactor // Write-only binding to the contract
	Groth16VerifierFilterer   // Log filterer for contract events
}

// Groth16VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type Groth16VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Groth16VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Groth16VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Groth16VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Groth16VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Groth16VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Groth16VerifierSession struct {
	Contract     *Groth16Verifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Groth16VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Groth16VerifierCallerSession struct {
	Contract *Groth16VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Groth16VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Groth16VerifierTransactorSession struct {
	Contract     *Groth16VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Groth16VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type Groth16VerifierRaw struct {
	Contract *Groth16Verifier // Generic contract binding to access the raw methods on
}

// Groth16VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Groth16VerifierCallerRaw struct {
	Contract *Groth16VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// Groth16VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Groth16VerifierTransactorRaw struct {
	Contract *Groth16VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGroth16Verifier creates a new instance of Groth16Verifier, bound to a specific deployed contract.
func NewGroth16Verifier(address common.Address, backend bind.ContractBackend) (*Groth16Verifier, error) {
	contract, err := bindGroth16Verifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Groth16Verifier{Groth16VerifierCaller: Groth16VerifierCaller{contract: contract}, Groth16VerifierTransactor: Groth16VerifierTransactor{contract: contract}, Groth16VerifierFilterer: Groth16VerifierFilterer{contract: contract}}, nil
}

// NewGroth16VerifierCaller creates a new read-only instance of Groth16Verifier, bound to a specific deployed contract.
func NewGroth16VerifierCaller(address common.Address, caller bind.ContractCaller) (*Groth16VerifierCaller, error) {
	contract, err := bindGroth16Verifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Groth16VerifierCaller{contract: contract}, nil
}

// NewGroth16VerifierTransactor creates a new write-only instance of Groth16Verifier, bound to a specific deployed contract.
func NewGroth16VerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*Groth16VerifierTransactor, error) {
	contract, err := bindGroth16Verifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Groth16VerifierTransactor{contract: contract}, nil
}

// NewGroth16VerifierFilterer creates a new log filterer instance of Groth16Verifier, bound to a specific deployed contract.
func NewGroth16VerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*Groth16VerifierFilterer, error) {
	contract, err := bindGroth16Verifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Groth16VerifierFilterer{contract: contract}, nil
}

// bindGroth16Verifier binds a generic wrapper to an already deployed contract.
func bindGroth16Verifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Groth16VerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Groth16Verifier *Groth16VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Groth16Verifier.Contract.Groth16VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Groth16Verifier *Groth16VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Groth16Verifier.Contract.Groth16VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Groth16Verifier *Groth16VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Groth16Verifier.Contract.Groth16VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Groth16Verifier *Groth16VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Groth16Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Groth16Verifier *Groth16VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Groth16Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Groth16Verifier *Groth16VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Groth16Verifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0xc1940a36.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[17] _pubSignals) view returns(bool)
func (_Groth16Verifier *Groth16VerifierCaller) VerifyProof(opts *bind.CallOpts, _pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [17]*big.Int) (bool, error) {
	var out []interface{}
	err := _Groth16Verifier.contract.Call(opts, &out, "verifyProof", _pA, _pB, _pC, _pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0xc1940a36.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[17] _pubSignals) view returns(bool)
func (_Groth16Verifier *Groth16VerifierSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [17]*big.Int) (bool, error) {
	return _Groth16Verifier.Contract.VerifyProof(&_Groth16Verifier.CallOpts, _pA, _pB, _pC, _pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0xc1940a36.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[17] _pubSignals) view returns(bool)
func (_Groth16Verifier *Groth16VerifierCallerSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [17]*big.Int) (bool, error) {
	return _Groth16Verifier.Contract.VerifyProof(&_Groth16Verifier.CallOpts, _pA, _pB, _pC, _pubSignals)
}
