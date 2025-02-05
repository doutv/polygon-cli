// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package webauthnandecdsavalidator

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

// WebAuthnAndECDSAValidatorMetaData contains all meta data concerning the WebAuthnAndECDSAValidator contract.
var WebAuthnAndECDSAValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"config\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CONFIG\",\"outputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"}],\"name\":\"isModuleType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signatureHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"validateSignatureWithData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a034606c57601f610da938819003918201601f19168301916001600160401b03831184841017607157808492602094604052833981010312606c57516001600160a01b0381168103606c57608052604051610d2190816100888239608051818181607701526101e60152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe60406080815260048036101561001457600080fd5b6000803560e01c8063940d3840146100aa578063d92e82e4146100625763ecd059611461004057600080fd5b3461005f57602036600319011261005f57506007602092519135148152f35b80fd5b5082346100a657816003193601126100a657517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b5082346100a65760603660031901126100a65767ffffffffffffffff906024358281116105ab576100de90369086016105b7565b6044959195358481116105b3576100f890369084016105b7565b9590916020968088116105af5785116100a657870184888203126100a65787358681116105af578161012b918a0161065a565b9787810135908782116105ab5761014392910161065a565b96805181019360c0828987019603126105af57878201518781116105ab57820185603f820112156105ab578581888b61017f94015191016106c4565b94868301518881116105a757830181603f820112156105a757898101516101a8929189016106c4565b9460608301519560808401519260c060a08601519501519760038910156105a3578951635aadb83360e11b81528281018a90526001600160a01b039d7f00000000000000000000000000000000000000000000000000000000000000008f16918e81602481865afa908115610599578f8b9261055e575b50501561054e578d60249261026261026b937f19457468657265756d205369676e6564204d6573736167653a0a3332000000008d528735601c52603c8d20610722565b9094919461075e565b8d516325d480ad60e01b8152338782015293849182905afa918215610544578992610506575b509d808e9f9e9c9d9e169116149b8b519183358c8401528b83526102b4836105ea565b60019b8d8751602581109081156104ce575b506104c4575b6103049291879151906102de826105ea565b6015825274113a3cb832911d113bb2b130baba34371733b2ba1160591b908201526108c5565b156104bb575b8b519260608401918211848310176104a85750928b6103ca8f9561038484966103cf96602e9386528582527f4142434445464748494a4b4c4d4e4f505152535455565758595a6162636465668a8301527f6768696a6b6c6d6e6f707172737475767778797a303132333435363738392d5f86830152610a14565b925180936c1131b430b63632b733b2911d1160991b898301526103b0815180928b602d860191016106a1565b8101601160f91b602d82015203600e81018452018261061c565b610851565b1561049e575b6103ec8691838c51928284809451938492016106a1565b8101039060025afa1561049457836104418b9282516104318c805180938861041d81840197888151938492016106a1565b82019089820152038781018452018261061c565b8b519283928392519283916106a1565b8101039060025afa1561048a5790839291896104639695013593359251610941565b9081610482575b508261047a575b50519015158152f35b915083610471565b90508461046a565b86513d84823e3d90fd5b87513d85823e3d90fd5b94975087946103d5565b634e487b7160e01b895260419052602488fd5b9950869961030a565b999c508c996102cc565b9291505010156104f3578e8d6104ec60ff60f81b828a0151166107e3565b15386102c6565b634e487b7160e01b895260328452602489fd5b9091508d81813d831161053d575b61051e818361061c565b8101031261053957519d8e81811603610539579d908e610291565b8880fd5b503d610514565b8c513d8b823e3d90fd5b8b5163baa3de5f60e01b81528490fd5b90809250813d8311610592575b610575818361061c565b8101031261058e5751801515810361058e57388f61021f565b8980fd5b503d61056b565b8d513d8c823e3d90fd5b8680fd5b8480fd5b8380fd5b8280fd5b8580fd5b9181601f840112156105e55782359167ffffffffffffffff83116105e557602083818601950101116105e557565b600080fd5b6040810190811067ffffffffffffffff82111761060657604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff82111761060657604052565b67ffffffffffffffff811161060657601f01601f191660200190565b81601f820112156105e5578035906106718261063e565b9261067f604051948561061c565b828452602083830101116105e557816000926020809301838601378301015290565b60005b8381106106b45750506000910152565b81810151838201526020016106a4565b909291926106d18161063e565b916106df604051938461061c565b8294828452828201116105e55760206106f99301906106a1565b565b90815181101561070c570160200190565b634e487b7160e01b600052603260045260246000fd5b81519190604183036107535761074c92506020820151906060604084015193015160001a90610984565b9192909190565b505060009160029190565b60048110156107cd5780610770575050565b6001810361078a5760405163f645eedf60e01b8152600490fd5b600281036107ab5760405163fce698f760e01b815260048101839052602490fd5b6003146107b55750565b602490604051906335e2f38360e21b82526004820152fd5b634e487b7160e01b600052602160045260246000fd5b6001600160f81b0319600160f81b82160161082857601f60fb1b600160fb1b821601610810575b50600190565b600160fc1b90811614610823573861080a565b600090565b50600090565b9190820180921161083b57565b634e487b7160e01b600052601160045260246000fd5b805191805160005b84811061086a575050505050600190565b601781810180911161083b5782811090811591610899575b5061088f57600101610859565b5050505050600090565b90506001600160f81b03196108bb816108b285896106fb565b511692866106fb565b5116141538610882565b805192825160009360005b8681106108e35750505050505050600190565b826108ee828661082e565b1080159061090d575b610903576001016108d0565b5050505050905090565b506001600160f81b03198061092283886106fb565b511690610938610932848861082e565b856106fb565b511614156108f7565b94939291907f7fffffff800000007fffffffffffffffde737d56d38bcf4279dce5617e3192a882116109795761097695610b60565b90565b505050505050600090565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411610a0857926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa156109fc5780516001600160a01b038116156109f357918190565b50809160019190565b604051903d90823e3d90fd5b50505060009160039190565b90815115610ae2578151600281811b93916001600160fe1b0381160361083b576002840180941161083b579091600380940493610a69610a538661063e565b95610a61604051978861061c565b80875261063e565b6020860190601f19013682379284805101926020840194855196600087525b858110610a99575050505050505290565b8460049101918251600190603f9082828260121c16880101518453828282600c1c16880101518385015382828260061c1688010151888501531685010151868201530190610a88565b50506040516020810181811067ffffffffffffffff821117610606576040526000815290565b6000915b60028310610b1957505050565b600190825181526020809101920192019190610b0c565b3d15610b5b573d90610b418261063e565b91610b4f604051938461061c565b82523d6000602084013e565b606090565b949193946040918251946020978887019784895283868901528060608901528260808901528160a089015260a0885260c088019367ffffffffffffffff97898610898711176106065785885260038110156107cd5780610c0957505050505050505050600091829151906101005afa90610bd8610b30565b91158015610c00575b610bf95780828051810103126105e557015160011490565b5050600090565b50815115610be1565b600103610c545750505050505050506000918291519073c2b78104907f722dabac4c69f826a522b2754de45afa50610c3f610b30565b9080828051810103126105e557015160011490565b610cb09597995097849698610c6d60e0939495966105ea565b87520152855191610c7d836105ea565b825287820152610ca6855193888501956304e960d760e01b875260248601526044850190610b08565b6084830190610b08565b60a4815260e0810193841181851017610606576000938493525190735ef35bc6ddb6425a0d43408d7810c3b3180ab58a5afa50610c3f610b3056fea26469706673582212202e149457ec85b3d5f152e1adaf6d9c084350f85042810e457f4448da11a6137964736f6c63430008190033",
}

// WebAuthnAndECDSAValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use WebAuthnAndECDSAValidatorMetaData.ABI instead.
var WebAuthnAndECDSAValidatorABI = WebAuthnAndECDSAValidatorMetaData.ABI

// WebAuthnAndECDSAValidatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WebAuthnAndECDSAValidatorMetaData.Bin instead.
var WebAuthnAndECDSAValidatorBin = WebAuthnAndECDSAValidatorMetaData.Bin

// DeployWebAuthnAndECDSAValidator deploys a new Ethereum contract, binding an instance of WebAuthnAndECDSAValidator to it.
func DeployWebAuthnAndECDSAValidator(auth *bind.TransactOpts, backend bind.ContractBackend, config common.Address) (common.Address, *types.Transaction, *WebAuthnAndECDSAValidator, error) {
	parsed, err := WebAuthnAndECDSAValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WebAuthnAndECDSAValidatorBin), backend, config)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WebAuthnAndECDSAValidator{WebAuthnAndECDSAValidatorCaller: WebAuthnAndECDSAValidatorCaller{contract: contract}, WebAuthnAndECDSAValidatorTransactor: WebAuthnAndECDSAValidatorTransactor{contract: contract}, WebAuthnAndECDSAValidatorFilterer: WebAuthnAndECDSAValidatorFilterer{contract: contract}}, nil
}

// WebAuthnAndECDSAValidator is an auto generated Go binding around an Ethereum contract.
type WebAuthnAndECDSAValidator struct {
	WebAuthnAndECDSAValidatorCaller     // Read-only binding to the contract
	WebAuthnAndECDSAValidatorTransactor // Write-only binding to the contract
	WebAuthnAndECDSAValidatorFilterer   // Log filterer for contract events
}

// WebAuthnAndECDSAValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type WebAuthnAndECDSAValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebAuthnAndECDSAValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WebAuthnAndECDSAValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebAuthnAndECDSAValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WebAuthnAndECDSAValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebAuthnAndECDSAValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WebAuthnAndECDSAValidatorSession struct {
	Contract     *WebAuthnAndECDSAValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WebAuthnAndECDSAValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WebAuthnAndECDSAValidatorCallerSession struct {
	Contract *WebAuthnAndECDSAValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// WebAuthnAndECDSAValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WebAuthnAndECDSAValidatorTransactorSession struct {
	Contract     *WebAuthnAndECDSAValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// WebAuthnAndECDSAValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type WebAuthnAndECDSAValidatorRaw struct {
	Contract *WebAuthnAndECDSAValidator // Generic contract binding to access the raw methods on
}

// WebAuthnAndECDSAValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WebAuthnAndECDSAValidatorCallerRaw struct {
	Contract *WebAuthnAndECDSAValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// WebAuthnAndECDSAValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WebAuthnAndECDSAValidatorTransactorRaw struct {
	Contract *WebAuthnAndECDSAValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWebAuthnAndECDSAValidator creates a new instance of WebAuthnAndECDSAValidator, bound to a specific deployed contract.
func NewWebAuthnAndECDSAValidator(address common.Address, backend bind.ContractBackend) (*WebAuthnAndECDSAValidator, error) {
	contract, err := bindWebAuthnAndECDSAValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WebAuthnAndECDSAValidator{WebAuthnAndECDSAValidatorCaller: WebAuthnAndECDSAValidatorCaller{contract: contract}, WebAuthnAndECDSAValidatorTransactor: WebAuthnAndECDSAValidatorTransactor{contract: contract}, WebAuthnAndECDSAValidatorFilterer: WebAuthnAndECDSAValidatorFilterer{contract: contract}}, nil
}

// NewWebAuthnAndECDSAValidatorCaller creates a new read-only instance of WebAuthnAndECDSAValidator, bound to a specific deployed contract.
func NewWebAuthnAndECDSAValidatorCaller(address common.Address, caller bind.ContractCaller) (*WebAuthnAndECDSAValidatorCaller, error) {
	contract, err := bindWebAuthnAndECDSAValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WebAuthnAndECDSAValidatorCaller{contract: contract}, nil
}

// NewWebAuthnAndECDSAValidatorTransactor creates a new write-only instance of WebAuthnAndECDSAValidator, bound to a specific deployed contract.
func NewWebAuthnAndECDSAValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*WebAuthnAndECDSAValidatorTransactor, error) {
	contract, err := bindWebAuthnAndECDSAValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WebAuthnAndECDSAValidatorTransactor{contract: contract}, nil
}

// NewWebAuthnAndECDSAValidatorFilterer creates a new log filterer instance of WebAuthnAndECDSAValidator, bound to a specific deployed contract.
func NewWebAuthnAndECDSAValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*WebAuthnAndECDSAValidatorFilterer, error) {
	contract, err := bindWebAuthnAndECDSAValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WebAuthnAndECDSAValidatorFilterer{contract: contract}, nil
}

// bindWebAuthnAndECDSAValidator binds a generic wrapper to an already deployed contract.
func bindWebAuthnAndECDSAValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WebAuthnAndECDSAValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WebAuthnAndECDSAValidator.Contract.WebAuthnAndECDSAValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WebAuthnAndECDSAValidator.Contract.WebAuthnAndECDSAValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WebAuthnAndECDSAValidator.Contract.WebAuthnAndECDSAValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WebAuthnAndECDSAValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WebAuthnAndECDSAValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WebAuthnAndECDSAValidator.Contract.contract.Transact(opts, method, params...)
}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorCaller) CONFIG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WebAuthnAndECDSAValidator.contract.Call(opts, &out, "CONFIG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorSession) CONFIG() (common.Address, error) {
	return _WebAuthnAndECDSAValidator.Contract.CONFIG(&_WebAuthnAndECDSAValidator.CallOpts)
}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorCallerSession) CONFIG() (common.Address, error) {
	return _WebAuthnAndECDSAValidator.Contract.CONFIG(&_WebAuthnAndECDSAValidator.CallOpts)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorCaller) IsModuleType(opts *bind.CallOpts, moduleTypeId *big.Int) (bool, error) {
	var out []interface{}
	err := _WebAuthnAndECDSAValidator.contract.Call(opts, &out, "isModuleType", moduleTypeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _WebAuthnAndECDSAValidator.Contract.IsModuleType(&_WebAuthnAndECDSAValidator.CallOpts, moduleTypeId)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorCallerSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _WebAuthnAndECDSAValidator.Contract.IsModuleType(&_WebAuthnAndECDSAValidator.CallOpts, moduleTypeId)
}

// ValidateSignatureWithData is a free data retrieval call binding the contract method 0x940d3840.
//
// Solidity: function validateSignatureWithData(bytes32 signatureHash, bytes signature, bytes data) view returns(bool)
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorCaller) ValidateSignatureWithData(opts *bind.CallOpts, signatureHash [32]byte, signature []byte, data []byte) (bool, error) {
	var out []interface{}
	err := _WebAuthnAndECDSAValidator.contract.Call(opts, &out, "validateSignatureWithData", signatureHash, signature, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateSignatureWithData is a free data retrieval call binding the contract method 0x940d3840.
//
// Solidity: function validateSignatureWithData(bytes32 signatureHash, bytes signature, bytes data) view returns(bool)
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorSession) ValidateSignatureWithData(signatureHash [32]byte, signature []byte, data []byte) (bool, error) {
	return _WebAuthnAndECDSAValidator.Contract.ValidateSignatureWithData(&_WebAuthnAndECDSAValidator.CallOpts, signatureHash, signature, data)
}

// ValidateSignatureWithData is a free data retrieval call binding the contract method 0x940d3840.
//
// Solidity: function validateSignatureWithData(bytes32 signatureHash, bytes signature, bytes data) view returns(bool)
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorCallerSession) ValidateSignatureWithData(signatureHash [32]byte, signature []byte, data []byte) (bool, error) {
	return _WebAuthnAndECDSAValidator.Contract.ValidateSignatureWithData(&_WebAuthnAndECDSAValidator.CallOpts, signatureHash, signature, data)
}
