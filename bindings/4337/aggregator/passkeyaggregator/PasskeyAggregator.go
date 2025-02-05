// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package passkeyaggregator

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

// PasskeyAggregatorMetaData contains all meta data concerning the PasskeyAggregator contract.
var PasskeyAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entrypoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CommitmentInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ENTRYPOINT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"}],\"name\":\"aggregateSignatures\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"aggregatedSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[10]\",\"name\":\"pubkeyX\",\"type\":\"uint256[10]\"},{\"internalType\":\"uint256[10]\",\"name\":\"pubkeyY\",\"type\":\"uint256[10]\"},{\"internalType\":\"uint256[10]\",\"name\":\"txHash\",\"type\":\"uint256[10]\"}],\"name\":\"callVerify10Signatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[10]\",\"name\":\"pubkeyX\",\"type\":\"uint256[10]\"},{\"internalType\":\"uint256[10]\",\"name\":\"pubkeyY\",\"type\":\"uint256[10]\"},{\"internalType\":\"uint256[10]\",\"name\":\"txHash\",\"type\":\"uint256[10]\"}],\"name\":\"hash10Signatures\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validateSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"validateUserOpSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"sigForUserOp\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[1]\",\"name\":\"input\",\"type\":\"uint256[1]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a034606c57601f6113bf38819003918201601f19168301916001600160401b03831184841017607157808492602094604052833981010312606c57516001600160a01b0381168103606c576080526040516113379081610088823960805181818160820152610f2b0152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c908163062a422b14611031575080632dd8113314610c2257806330a3a20e1461087357806343db3c721461027a578063ae574a4314610225578063d588b459146100b65763e8eb3cc61461006c57600080fd5b346100b15760003660031901126100b1576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b600080fd5b346100b1576103c0806003193601126100b1576101443681116100b157610284903682116100b157366103c4116100b1576020926040519284840192600435845235604085015235606084015260243560808401526101643560a08401526102a43560c084015260443560e0840152610184356101008401526102c4356101208401526064356101408401526101a4356101608401526102e4356101808401526084356101a08401526101c4356101c0840152610304356101e084015260a4356102008401526101e4356102208401526103243561024084015260c43561026084015261020435610280840152610344356102a084015260e4356102c0840152610224356102e08401526103643561030084015261010435610320840152610244356103408401526103843561036084015261012435610380840152610264356103a08401526103a43581840152825261020f8261111e565b905190206040516001600160f81b039091168152f35b346100b15760203660031901126100b1576004356001600160401b0381116100b1576102559036906004016110b7565b5050610276604051610266816110e7565b600081526040519182918261106e565b0390f35b346100b1576101a03660031901126100b1576101043681116100b157610144903682116100b157610184903682116100b157366101a312156100b157604051916102c3836110e7565b826101a4913683116100b157905b82821061086357505050604051926102e8846110e7565b60203685376040516020810190833582526101243560408201526060810160605160809060005b81811061084d5750505091816103406000805160206112e2833981519152936040969503601f198101835282611155565b5190200685528180519384377f032826efd3da14888756b1889f120cf1711b82a77c8d1422e6bcf0436ac1ce5e828401527f01662a6d31bc7ffa43d1a466cf157dd76ff27680f3936151df88a31e3d1b795c60608401527f2011f1ea76efd601e525ca6f68d2bb07385711a85d135ce5410bd9a1a9f0da7060808401527f11c0cc9c7b6a10e9e9cc696945597e175321a256f42d571d0ccd602175908c8f60a084015260c08301377f0a21766bc1fe523f08b885b4011fbae89777991c318ffd6127cd45ef256d3f7a6101008201527f03d6ebd8b1d35d15507a3e0fc9ae20e20e7dc2f3ad0dde07b4c1b9c353b3cad66101208201527f1d379ec314086e52b27f935e089d0f689b8035108d6e26cde35ec4212a4062806101408201527f1bccaa80fa28cc639c7f288bd6dc8eca093585ce54f82b92acb5928d6ee545e86101608201526020816101808160085afa9051161561083b576001906104a336611176565b926020604051947f184932331b713e5635bb7e0443166753b1a3b9cb8a6cdb59e89bcc6f567e15c886527f2f535d862892b75efbb060b71ac0d80de3ed7709fbe9f07a6d435b6933a2fe928287015280516040870152015190606085019182527f199155d53afcaf3058138da8194650172b48259bea110fce7a6571c583e0d04860408660808160065afa9360006040880152600084525160808701948186526000805160206112e28339815191526040808a016060828c0160075afa9210161660408760808160065afa16927f0e79c9a617690548a2bd3f098c025014522c02817c98f4abe46ce803b31ebd2e604088015252518092526000805160206112e28339815191526040808601606082880160075afa9210161660408360808160065afa1616906020815191015191156108295760209160405191610100600484377f021b8bfa93c806690419599de5f2014a99adef3c695823a7b3a0c3ea7a14d5c96101008401527f1e87dfbc0e15687d9b98cad686ad8a3b2fbadb4e8bcb25ac146b46633ab16ac86101208401527f236ad71ff63cf1e83eaf8f2a3b255c4acedccd5bb4cc40d08925bf0367b748566101408401527f13918ac57d51930fbbde9ef4edadd0f8452d95a48cc8d15a88a7ffc85a5d98e36101608401527f05b5a6e0242a1be7b69ea0aec06b80141100e461f7509299e101f6de6b4109556101808401527f276ba5104aa5e5d103ad8c64d1b4524bfb65c0bd0d61428347bd630f9647857b6101a08401527f1999c9b3857b55a30a5a9d57ff60ed8cc4672bbd8bd777325fd26c746fe628906101c08401527f144c1cf6a462d5494e52cc331eecf618d67eca3ed935df9592e32f3c13dd86eb6101e08401527f1c1bf941d3a1df4a0c7ab35e7cde95db0d8045262be4666879b07dd2754aff2a6102008401527f289aceba05f5758f8d639a9e1b78f53c18efe25acbbcb110c71d26a2c5ce1f7d6102208401526102408301526102608201527f0f81a3fbbf1fae092e1c314b75bb154515e82213991b95f2248cb0a8c968ceda6102808201527f2ef50caab497754de5e44560eb6413a585935b4ed7f0f9552ab49c25aa3fa5106102a08201527f294009d0a3889406bc9ba32f49694ed6e519f85bbbba77c5b7e22071b19d3ba76102c08201527f1b1f82c5746904927f255ffeedeeb55dcacbd0a176bc59290cc383dc62d4fe216102e082015281816103008160085afa9051166040519015158152f35b60405163a54f8e2760e01b8152600490fd5b6040516351d49ff760e11b8152600490fd5b825184526020938401939092019160010161030f565b81358152602091820191016102d1565b346100b1576105403660031901126100b15736610104116100b15736610144116100b15736610184116100b157366102c4116100b15736610404116100b15736610544116100b1576040516101843560208201526102c43560408201526104043560608201526101a43560808201526102e43560a08201526104243560c08201526101c43560e082015261030435610100820152610444356101208201526101e4356101408201526103243561016082015261046435610180820152610204356101a0820152610344356101c0820152610484356101e082015261022435610200820152610364356102208201526104a43561024082015261024435610260820152610384356102808201526104c4356102a0820152610264356102c08201526103a4356102e08201526104e435610300820152610284356103208201526103c435610340820152610504356103608201526102a4356103808201526103e4356103a08201526103c0610524358183015281526109ef8161111e565b60018060f81b03906020815191012016604051908160208101106001600160401b03602084011117610c0c5760208201604052815260405190610a31826110e7565b60203683376040516101043560208201526101243560408201526060810160605160809060005b818110610bf65750505090610a85816000805160206112e28339815191529303601f198101835282611155565b60208151910120068252604051604061010482377f032826efd3da14888756b1889f120cf1711b82a77c8d1422e6bcf0436ac1ce5e60408201527f01662a6d31bc7ffa43d1a466cf157dd76ff27680f3936151df88a31e3d1b795c60608201527f2011f1ea76efd601e525ca6f68d2bb07385711a85d135ce5410bd9a1a9f0da7060808201527f11c0cc9c7b6a10e9e9cc696945597e175321a256f42d571d0ccd602175908c8f60a0820152604061014460c08301377f0a21766bc1fe523f08b885b4011fbae89777991c318ffd6127cd45ef256d3f7a6101008201527f03d6ebd8b1d35d15507a3e0fc9ae20e20e7dc2f3ad0dde07b4c1b9c353b3cad66101208201527f1d379ec314086e52b27f935e089d0f689b8035108d6e26cde35ec4212a4062806101408201527f1bccaa80fa28cc639c7f288bd6dc8eca093585ce54f82b92acb5928d6ee545e86101608201526020816101808160085afa9051161561083b576001906104a336611176565b8251845260209384019390920191600101610a58565b634e487b7160e01b600052604160045260246000fd5b346100b15760403660031901126100b1576004356001600160401b0381116100b157610c529036906004016110b7565b602435916001600160401b0383116100b157366023840112156100b15782600401356001600160401b0381116100b15783019260248401933685116100b157816101809103126100b157836043820112156100b1576040519361010085018581106001600160401b03821117610c0c576040526101248201858282116100b15760248401905b82821061102157505090610164610cf282610cf9946111b0565b93016111b0565b60405191610d0683611102565b610140948536853760405194610d1b86611102565b8636873760405196610d2c88611102565b36883760005b828110610e36575050604051631851d10760e11b81529660009150600488015b60088310610e1f5788610da589610d9a8a610d8f8b610d848c610d796101048a018e611291565b610144890190611291565b6101848701906112b9565b6102c48501906112b9565b6104048301906112b9565b60208161054481305afa908115610e1357600091610dd8575b5015610dc657005b604051638baa579f60e01b8152600490fd5b90506020813d602011610e0b575b81610df360209383611155565b810103126100b1575180151581036100b15781610dbe565b3d9150610de6565b6040513d6000823e3d90fd5b600190825181526020809101920192019190610d52565b610e418184846111f5565b6040516308b3779360e21b815260206004820152919080356001600160a01b038116908190036100b15783610f2760209382936024840152848101356044840152610ea5610e92604083018361122e565b610120606487015261014486019161125f565b610f17610f0b610ecf610ebb606086018661122e565b602319898703810160848b0152959161125f565b608085013560a488015260a085013560c488015260c085013560e4880152610efa60e086018661122e565b9085898403016101048a015261125f565b9261010081019061122e565b918584030161012486015261125f565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa918215610e1357600092610fed575b50610f708185856111f5565b61010081013590601e19813603018212156100b1570180356001600160401b0381116100b1576020820190360381136100b1573501916020830135806020116100b1576040116100b15760608360406001950135610fce858c611280565b520135610fdb838b611280565b52610fe6828b611280565b5201610d32565b9091506020813d602011611019575b8161100960209383611155565b810103126100b15751908a610f64565b3d9150610ffc565b8135815260209182019101610cd8565b346100b1576003196020368201126100b157600435906001600160401b0382116100b15761012091360301126100b15780610266610276926110e7565b6020808252825181830181905290939260005b8281106110a357505060409293506000838284010152601f8019910116010190565b818101860151848201604001528501611081565b9181601f840112156100b1578235916001600160401b0383116100b1576020808501948460051b0101116100b157565b602081019081106001600160401b03821117610c0c57604052565b61014081019081106001600160401b03821117610c0c57604052565b6103e081019081106001600160401b03821117610c0c57604052565b604081019081106001600160401b03821117610c0c57604052565b90601f801991011681019081106001600160401b03821117610c0c57604052565b90604051916111848361113a565b826101449182116100b157610104905b8282106111a057505050565b8135815260209182019101611194565b9080601f830112156100b157604051916111c98361113a565b8290604081019283116100b157905b8282106111e55750505090565b81358152602091820191016111d8565b91908110156112185760051b8101359061011e19813603018212156100b1570190565b634e487b7160e01b600052603260045260246000fd5b9035601e19823603018112156100b15701602081359101916001600160401b0382116100b15781360383136100b157565b908060209392818452848401376000828201840152601f01601f1916010190565b90600a8110156112185760051b0190565b6000915b600283106112a257505050565b600190825181526020809101920192019190611295565b6000915b600a83106112ca57505050565b6001908251815260208091019201920191906112bd56fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220958886651b327342daa0d99281b32ee01589f2af299549e0855bbfe0e2e707c564736f6c63430008190033",
}

// PasskeyAggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use PasskeyAggregatorMetaData.ABI instead.
var PasskeyAggregatorABI = PasskeyAggregatorMetaData.ABI

// PasskeyAggregatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PasskeyAggregatorMetaData.Bin instead.
var PasskeyAggregatorBin = PasskeyAggregatorMetaData.Bin

// DeployPasskeyAggregator deploys a new Ethereum contract, binding an instance of PasskeyAggregator to it.
func DeployPasskeyAggregator(auth *bind.TransactOpts, backend bind.ContractBackend, _entrypoint common.Address) (common.Address, *types.Transaction, *PasskeyAggregator, error) {
	parsed, err := PasskeyAggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PasskeyAggregatorBin), backend, _entrypoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PasskeyAggregator{PasskeyAggregatorCaller: PasskeyAggregatorCaller{contract: contract}, PasskeyAggregatorTransactor: PasskeyAggregatorTransactor{contract: contract}, PasskeyAggregatorFilterer: PasskeyAggregatorFilterer{contract: contract}}, nil
}

// PasskeyAggregator is an auto generated Go binding around an Ethereum contract.
type PasskeyAggregator struct {
	PasskeyAggregatorCaller     // Read-only binding to the contract
	PasskeyAggregatorTransactor // Write-only binding to the contract
	PasskeyAggregatorFilterer   // Log filterer for contract events
}

// PasskeyAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type PasskeyAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PasskeyAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PasskeyAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PasskeyAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PasskeyAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PasskeyAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PasskeyAggregatorSession struct {
	Contract     *PasskeyAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PasskeyAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PasskeyAggregatorCallerSession struct {
	Contract *PasskeyAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PasskeyAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PasskeyAggregatorTransactorSession struct {
	Contract     *PasskeyAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PasskeyAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type PasskeyAggregatorRaw struct {
	Contract *PasskeyAggregator // Generic contract binding to access the raw methods on
}

// PasskeyAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PasskeyAggregatorCallerRaw struct {
	Contract *PasskeyAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// PasskeyAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PasskeyAggregatorTransactorRaw struct {
	Contract *PasskeyAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPasskeyAggregator creates a new instance of PasskeyAggregator, bound to a specific deployed contract.
func NewPasskeyAggregator(address common.Address, backend bind.ContractBackend) (*PasskeyAggregator, error) {
	contract, err := bindPasskeyAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PasskeyAggregator{PasskeyAggregatorCaller: PasskeyAggregatorCaller{contract: contract}, PasskeyAggregatorTransactor: PasskeyAggregatorTransactor{contract: contract}, PasskeyAggregatorFilterer: PasskeyAggregatorFilterer{contract: contract}}, nil
}

// NewPasskeyAggregatorCaller creates a new read-only instance of PasskeyAggregator, bound to a specific deployed contract.
func NewPasskeyAggregatorCaller(address common.Address, caller bind.ContractCaller) (*PasskeyAggregatorCaller, error) {
	contract, err := bindPasskeyAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PasskeyAggregatorCaller{contract: contract}, nil
}

// NewPasskeyAggregatorTransactor creates a new write-only instance of PasskeyAggregator, bound to a specific deployed contract.
func NewPasskeyAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*PasskeyAggregatorTransactor, error) {
	contract, err := bindPasskeyAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PasskeyAggregatorTransactor{contract: contract}, nil
}

// NewPasskeyAggregatorFilterer creates a new log filterer instance of PasskeyAggregator, bound to a specific deployed contract.
func NewPasskeyAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*PasskeyAggregatorFilterer, error) {
	contract, err := bindPasskeyAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PasskeyAggregatorFilterer{contract: contract}, nil
}

// bindPasskeyAggregator binds a generic wrapper to an already deployed contract.
func bindPasskeyAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PasskeyAggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PasskeyAggregator *PasskeyAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PasskeyAggregator.Contract.PasskeyAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PasskeyAggregator *PasskeyAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PasskeyAggregator.Contract.PasskeyAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PasskeyAggregator *PasskeyAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PasskeyAggregator.Contract.PasskeyAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PasskeyAggregator *PasskeyAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PasskeyAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PasskeyAggregator *PasskeyAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PasskeyAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PasskeyAggregator *PasskeyAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PasskeyAggregator.Contract.contract.Transact(opts, method, params...)
}

// ENTRYPOINT is a free data retrieval call binding the contract method 0xe8eb3cc6.
//
// Solidity: function ENTRYPOINT() view returns(address)
func (_PasskeyAggregator *PasskeyAggregatorCaller) ENTRYPOINT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PasskeyAggregator.contract.Call(opts, &out, "ENTRYPOINT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ENTRYPOINT is a free data retrieval call binding the contract method 0xe8eb3cc6.
//
// Solidity: function ENTRYPOINT() view returns(address)
func (_PasskeyAggregator *PasskeyAggregatorSession) ENTRYPOINT() (common.Address, error) {
	return _PasskeyAggregator.Contract.ENTRYPOINT(&_PasskeyAggregator.CallOpts)
}

// ENTRYPOINT is a free data retrieval call binding the contract method 0xe8eb3cc6.
//
// Solidity: function ENTRYPOINT() view returns(address)
func (_PasskeyAggregator *PasskeyAggregatorCallerSession) ENTRYPOINT() (common.Address, error) {
	return _PasskeyAggregator.Contract.ENTRYPOINT(&_PasskeyAggregator.CallOpts)
}

// AggregateSignatures is a free data retrieval call binding the contract method 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) pure returns(bytes aggregatedSignature)
func (_PasskeyAggregator *PasskeyAggregatorCaller) AggregateSignatures(opts *bind.CallOpts, userOps []PackedUserOperation) ([]byte, error) {
	var out []interface{}
	err := _PasskeyAggregator.contract.Call(opts, &out, "aggregateSignatures", userOps)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AggregateSignatures is a free data retrieval call binding the contract method 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) pure returns(bytes aggregatedSignature)
func (_PasskeyAggregator *PasskeyAggregatorSession) AggregateSignatures(userOps []PackedUserOperation) ([]byte, error) {
	return _PasskeyAggregator.Contract.AggregateSignatures(&_PasskeyAggregator.CallOpts, userOps)
}

// AggregateSignatures is a free data retrieval call binding the contract method 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) pure returns(bytes aggregatedSignature)
func (_PasskeyAggregator *PasskeyAggregatorCallerSession) AggregateSignatures(userOps []PackedUserOperation) ([]byte, error) {
	return _PasskeyAggregator.Contract.AggregateSignatures(&_PasskeyAggregator.CallOpts, userOps)
}

// CallVerify10Signatures is a free data retrieval call binding the contract method 0x30a3a20e.
//
// Solidity: function callVerify10Signatures(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[10] pubkeyX, uint256[10] pubkeyY, uint256[10] txHash) view returns(bool success)
func (_PasskeyAggregator *PasskeyAggregatorCaller) CallVerify10Signatures(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, pubkeyX [10]*big.Int, pubkeyY [10]*big.Int, txHash [10]*big.Int) (bool, error) {
	var out []interface{}
	err := _PasskeyAggregator.contract.Call(opts, &out, "callVerify10Signatures", proof, commitments, commitmentPok, pubkeyX, pubkeyY, txHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CallVerify10Signatures is a free data retrieval call binding the contract method 0x30a3a20e.
//
// Solidity: function callVerify10Signatures(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[10] pubkeyX, uint256[10] pubkeyY, uint256[10] txHash) view returns(bool success)
func (_PasskeyAggregator *PasskeyAggregatorSession) CallVerify10Signatures(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, pubkeyX [10]*big.Int, pubkeyY [10]*big.Int, txHash [10]*big.Int) (bool, error) {
	return _PasskeyAggregator.Contract.CallVerify10Signatures(&_PasskeyAggregator.CallOpts, proof, commitments, commitmentPok, pubkeyX, pubkeyY, txHash)
}

// CallVerify10Signatures is a free data retrieval call binding the contract method 0x30a3a20e.
//
// Solidity: function callVerify10Signatures(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[10] pubkeyX, uint256[10] pubkeyY, uint256[10] txHash) view returns(bool success)
func (_PasskeyAggregator *PasskeyAggregatorCallerSession) CallVerify10Signatures(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, pubkeyX [10]*big.Int, pubkeyY [10]*big.Int, txHash [10]*big.Int) (bool, error) {
	return _PasskeyAggregator.Contract.CallVerify10Signatures(&_PasskeyAggregator.CallOpts, proof, commitments, commitmentPok, pubkeyX, pubkeyY, txHash)
}

// Hash10Signatures is a free data retrieval call binding the contract method 0xd588b459.
//
// Solidity: function hash10Signatures(uint256[10] pubkeyX, uint256[10] pubkeyY, uint256[10] txHash) pure returns(bytes32)
func (_PasskeyAggregator *PasskeyAggregatorCaller) Hash10Signatures(opts *bind.CallOpts, pubkeyX [10]*big.Int, pubkeyY [10]*big.Int, txHash [10]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PasskeyAggregator.contract.Call(opts, &out, "hash10Signatures", pubkeyX, pubkeyY, txHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash10Signatures is a free data retrieval call binding the contract method 0xd588b459.
//
// Solidity: function hash10Signatures(uint256[10] pubkeyX, uint256[10] pubkeyY, uint256[10] txHash) pure returns(bytes32)
func (_PasskeyAggregator *PasskeyAggregatorSession) Hash10Signatures(pubkeyX [10]*big.Int, pubkeyY [10]*big.Int, txHash [10]*big.Int) ([32]byte, error) {
	return _PasskeyAggregator.Contract.Hash10Signatures(&_PasskeyAggregator.CallOpts, pubkeyX, pubkeyY, txHash)
}

// Hash10Signatures is a free data retrieval call binding the contract method 0xd588b459.
//
// Solidity: function hash10Signatures(uint256[10] pubkeyX, uint256[10] pubkeyY, uint256[10] txHash) pure returns(bytes32)
func (_PasskeyAggregator *PasskeyAggregatorCallerSession) Hash10Signatures(pubkeyX [10]*big.Int, pubkeyY [10]*big.Int, txHash [10]*big.Int) ([32]byte, error) {
	return _PasskeyAggregator.Contract.Hash10Signatures(&_PasskeyAggregator.CallOpts, pubkeyX, pubkeyY, txHash)
}

// ValidateSignatures is a free data retrieval call binding the contract method 0x2dd81133.
//
// Solidity: function validateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, bytes signature) view returns()
func (_PasskeyAggregator *PasskeyAggregatorCaller) ValidateSignatures(opts *bind.CallOpts, userOps []PackedUserOperation, signature []byte) error {
	var out []interface{}
	err := _PasskeyAggregator.contract.Call(opts, &out, "validateSignatures", userOps, signature)

	if err != nil {
		return err
	}

	return err

}

// ValidateSignatures is a free data retrieval call binding the contract method 0x2dd81133.
//
// Solidity: function validateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, bytes signature) view returns()
func (_PasskeyAggregator *PasskeyAggregatorSession) ValidateSignatures(userOps []PackedUserOperation, signature []byte) error {
	return _PasskeyAggregator.Contract.ValidateSignatures(&_PasskeyAggregator.CallOpts, userOps, signature)
}

// ValidateSignatures is a free data retrieval call binding the contract method 0x2dd81133.
//
// Solidity: function validateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, bytes signature) view returns()
func (_PasskeyAggregator *PasskeyAggregatorCallerSession) ValidateSignatures(userOps []PackedUserOperation, signature []byte) error {
	return _PasskeyAggregator.Contract.ValidateSignatures(&_PasskeyAggregator.CallOpts, userOps, signature)
}

// ValidateUserOpSignature is a free data retrieval call binding the contract method 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes sigForUserOp)
func (_PasskeyAggregator *PasskeyAggregatorCaller) ValidateUserOpSignature(opts *bind.CallOpts, userOp PackedUserOperation) ([]byte, error) {
	var out []interface{}
	err := _PasskeyAggregator.contract.Call(opts, &out, "validateUserOpSignature", userOp)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ValidateUserOpSignature is a free data retrieval call binding the contract method 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes sigForUserOp)
func (_PasskeyAggregator *PasskeyAggregatorSession) ValidateUserOpSignature(userOp PackedUserOperation) ([]byte, error) {
	return _PasskeyAggregator.Contract.ValidateUserOpSignature(&_PasskeyAggregator.CallOpts, userOp)
}

// ValidateUserOpSignature is a free data retrieval call binding the contract method 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes sigForUserOp)
func (_PasskeyAggregator *PasskeyAggregatorCallerSession) ValidateUserOpSignature(userOp PackedUserOperation) ([]byte, error) {
	return _PasskeyAggregator.Contract.ValidateUserOpSignature(&_PasskeyAggregator.CallOpts, userOp)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43db3c72.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[1] input) view returns(bool success)
func (_PasskeyAggregator *PasskeyAggregatorCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [1]*big.Int) (bool, error) {
	var out []interface{}
	err := _PasskeyAggregator.contract.Call(opts, &out, "verifyProof", proof, commitments, commitmentPok, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x43db3c72.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[1] input) view returns(bool success)
func (_PasskeyAggregator *PasskeyAggregatorSession) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _PasskeyAggregator.Contract.VerifyProof(&_PasskeyAggregator.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43db3c72.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[1] input) view returns(bool success)
func (_PasskeyAggregator *PasskeyAggregatorCallerSession) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _PasskeyAggregator.Contract.VerifyProof(&_PasskeyAggregator.CallOpts, proof, commitments, commitmentPok, input)
}
