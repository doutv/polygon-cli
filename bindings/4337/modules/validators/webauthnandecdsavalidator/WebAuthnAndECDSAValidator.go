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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"config\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CONFIG\",\"outputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"}],\"name\":\"isModuleType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signatureHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"isValidSignatureWithData\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a034606c57601f610f7238819003918201601f19168301916001600160401b03831184841017607157808492602094604052833981010312606c57516001600160a01b0381168103606c57608052604051610eea90816100888239608051818181608d01526102470152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe60406080815260048036101561001457600080fd5b6000803560e01c806354c015f3146100c55780636d61fe70146100c05780638a91b0e3146100c0578063d92e82e4146100785763ecd059611461005657600080fd5b3461007557602036600319011261007557506001602092519135148152f35b80fd5b5082346100bc57816003193601126100bc57517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b610379565b5082346100bc5760603660031901126100bc5782359167ffffffffffffffff906024358281116100bc576100fc9036908701610346565b604496919635848111610342576101169036908401610346565b96909160209780891161033a57871161033e578801868982031261033e57883586811161033a5781610149918b01610438565b98888101359087821161033657610161929101610438565b978051810160c0828a830192031261033a578882015187811161033657820181603f820112156103365781818a8c61019c94015191016104a2565b968883015190811161033657820181603f8201121561033657898101516101c592918a016104a2565b606082015160808301519160c060a08501519401519660ff88168803610332578b9c7f19457468657265756d205369676e6564204d6573736167653a0a3332000000008a5286601c52603c8a209061021c9161092a565b61022591610966565b600160a01b600190039d8e928d51938491637a1c844960e01b835233908301527f000000000000000000000000000000000000000000000000000000000000000016815a91602492fa918215610328578a926102ea575b509c806102ab9a9b9c9d9e16911614998b51958d8701528c865261029f866103ac565b8c870135963595610500565b826102e2575b5050156102d357630b135d3f60e11b905b516001600160e01b03199091168152f35b6001600160e01b0319906102c2565b9150846102b1565b9091508c81813d8311610321575b61030281836103fa565b8101031261031d57519c8d8181160361031d579c908d61027c565b8980fd5b503d6102f8565b8c513d8c823e3d90fd5b8880fd5b8680fd5b8580fd5b8480fd5b8380fd5b9181601f840112156103745782359167ffffffffffffffff8311610374576020838186019501011161037457565b600080fd5b346103745760203660031901126103745760043567ffffffffffffffff8111610374576103aa903690600401610346565b005b6040810190811067ffffffffffffffff8211176103c857604052565b634e487b7160e01b600052604160045260246000fd5b60c0810190811067ffffffffffffffff8211176103c857604052565b90601f8019910116810190811067ffffffffffffffff8211176103c857604052565b67ffffffffffffffff81116103c857601f01601f191660200190565b81601f820112156103745780359061044f8261041c565b9261045d60405194856103fa565b8284526020838301011161037457816000926020809301838601378301015290565b60005b8381106104925750506000910152565b8181015183820152602001610482565b909291926104af8161041c565b916104bd60405193846103fa565b8294828452828201116103745760206104d793019061047f565b565b9081518110156104ea570160200190565b634e487b7160e01b600052603260045260246000fd5b979593919694929096600197805160258110908115610901575b506108f8575b60409182519a61052f8c6103ac565b60158c5261055e858360209e8f74113a3cb832911d113bb2b130baba34371733b2ba1160591b90820152610ab7565b156108ef575b8351606081019181831067ffffffffffffffff8411176103c8577f4142434445464748494a4b4c4d4e4f505152535455565758595a6162636465668e6105d69488528784528301527f6768696a6b6c6d6e6f707172737475767778797a303132333435363738392b2f86830152610d7f565b600081516002811190816108bf575b5015610860575060025b815190810390811161084a578c92916106078261041c565b91610614875193846103fa565b8083526106208161041c565b8386019290601f190136843760005b8281106107bf5750505061066e610688602e61068d9486948a519485926c1131b430b63632b733b2911d1160991b8b850152518092602d85019061047f565b8101601160f91b602d82015203600e8101845201826103fa565b610a43565b156107b6575b6106ab6000918386519282848094519384920161047f565b8101039060025afa156107ab5760006107018b9282516106f18680518093886106dd818401978881519384920161047f565b8201908982015203878101845201826103fa565b855192839283925192839161047f565b8101039060025afa156107a15760005191600019146107355750610726969750610c1b565b9081610730575090565b905090565b939160009750879650849593919551948986019687528501526060840152608083015260a082015260a0815261076a816103de565b519073c2b78104907f722dabac4c69f826a522b2754de45afa5061078c610beb565b90808280518101031261037457015160011490565b513d6000823e3d90fd5b50513d6000823e3d90fd5b60009a50610693565b9495509293919290916001906001600160f81b0319602b60f81b816107e484876104d9565b5116036108075750602d6107f882886104d9565b535b01908f959493929161062f565b602f60f81b8161081784876104d9565b5116036108315750605f61082b82886104d9565b536107fa565b61083b82856104d9565b511660001a61082b82886104d9565b634e487b7160e01b600052601160045260246000fd5b815160018111908161087b575b50156105ef575060016105ef565b6000198101915081116108ab57603d60f81b906001600160f81b0319906108a290856104d9565b5116143861086d565b634e487b7160e01b82526011600452602482fd5b6001198101915081116108ab57603d60f81b906001600160f81b0319906108e690856104d9565b511614386105e5565b60009a50610564565b60009850610520565b9050602010156104ea576040810151610923906001600160f81b0319166109eb565b153861051a565b815191906041830361095b5761095492506020820151906060604084015193015160001a90610b33565b9192909190565b505060009160029190565b60048110156109d55780610978575050565b600181036109925760405163f645eedf60e01b8152600490fd5b600281036109b35760405163fce698f760e01b815260048101839052602490fd5b6003146109bd5750565b602490604051906335e2f38360e21b82526004820152fd5b634e487b7160e01b600052602160045260246000fd5b6001600160f81b0319600160f81b821601610a3057601f60fb1b600160fb1b821601610a18575b50600190565b600160fc1b90811614610a2b5738610a12565b600090565b50600090565b9190820180921161084a57565b805191805160005b848110610a5c575050505050600190565b601781810180911161084a5782811090811591610a8b575b50610a8157600101610a4b565b5050505050600090565b90506001600160f81b0319610aad81610aa485896104d9565b511692866104d9565b5116141538610a74565b805192825160009360005b868110610ad55750505050505050600190565b82610ae08286610a36565b10801590610aff575b610af557600101610ac2565b5050505050905090565b506001600160f81b031980610b1483886104d9565b511690610b2a610b248488610a36565b856104d9565b51161415610ae9565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411610bb757926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa15610bab5780516001600160a01b03811615610ba257918190565b50809160019190565b604051903d90823e3d90fd5b50505060009160039190565b6000915b60028310610bd457505050565b600190825181526020809101920192019190610bc7565b3d15610c16573d90610bfc8261041c565b91610c0a60405193846103fa565b82523d6000602084013e565b606090565b94939290919460409182519460209760ff89880198858a5287878a01528260608a01528460808a01528360a08a015260a08952610c57896103de565b1680610caa5750505050505050600091829151906101005afa90610c79610beb565b91158015610ca1575b610c9a57808280518101031261037457015160011490565b5050600090565b50815115610c82565b600103610cde575050505050506000918291519073c2b78104907f722dabac4c69f826a522b2754de45afa5061078c610beb565b90919294965083610d3b9496505196610cf6886103ac565b875287870152845191610d08836103ac565b825286820152610d31845195878701946304e960d760e01b865260248801526044870190610bc3565b6084850190610bc3565b60a4835260e083019280841067ffffffffffffffff8511176103c8576000938493525190731704af085ca24d18100d3c98d87c96f4dd5557c65afa5061078c610beb565b919091805115610e8d578051926002916002850180951161084a57600394859004600281901b93906001600160fe1b0381160361084a5794610dd9610dc38561041c565b94610dd160405196876103fa565b80865261041c565b6020850190601f190136823792829183518401976020890192835194600085525b8a8110610e4057505050506003939495965052510680600114610e2d57600214610e22575090565b603d90600019015390565b50603d9081600019820153600119015390565b836004919b989b019a8b51600190603f9082828260121c16870101518453828282600c1c16870101518385015382828260061c168701015187850153168401015185820153019699610dfa565b5090506040516020810181811067ffffffffffffffff8211176103c857604052600081529056fea26469706673582212208e6fbc388eb8498fcee74d348b0d7c7610b9de20598310316305c4bfc4bd98b064736f6c63430008190033",
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

// IsValidSignatureWithData is a free data retrieval call binding the contract method 0x54c015f3.
//
// Solidity: function isValidSignatureWithData(bytes32 signatureHash, bytes signature, bytes data) view returns(bytes4)
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorCaller) IsValidSignatureWithData(opts *bind.CallOpts, signatureHash [32]byte, signature []byte, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _WebAuthnAndECDSAValidator.contract.Call(opts, &out, "isValidSignatureWithData", signatureHash, signature, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignatureWithData is a free data retrieval call binding the contract method 0x54c015f3.
//
// Solidity: function isValidSignatureWithData(bytes32 signatureHash, bytes signature, bytes data) view returns(bytes4)
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorSession) IsValidSignatureWithData(signatureHash [32]byte, signature []byte, data []byte) ([4]byte, error) {
	return _WebAuthnAndECDSAValidator.Contract.IsValidSignatureWithData(&_WebAuthnAndECDSAValidator.CallOpts, signatureHash, signature, data)
}

// IsValidSignatureWithData is a free data retrieval call binding the contract method 0x54c015f3.
//
// Solidity: function isValidSignatureWithData(bytes32 signatureHash, bytes signature, bytes data) view returns(bytes4)
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorCallerSession) IsValidSignatureWithData(signatureHash [32]byte, signature []byte, data []byte) ([4]byte, error) {
	return _WebAuthnAndECDSAValidator.Contract.IsValidSignatureWithData(&_WebAuthnAndECDSAValidator.CallOpts, signatureHash, signature, data)
}

// OnInstall is a free data retrieval call binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) pure returns()
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorCaller) OnInstall(opts *bind.CallOpts, data []byte) error {
	var out []interface{}
	err := _WebAuthnAndECDSAValidator.contract.Call(opts, &out, "onInstall", data)

	if err != nil {
		return err
	}

	return err

}

// OnInstall is a free data retrieval call binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) pure returns()
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorSession) OnInstall(data []byte) error {
	return _WebAuthnAndECDSAValidator.Contract.OnInstall(&_WebAuthnAndECDSAValidator.CallOpts, data)
}

// OnInstall is a free data retrieval call binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) pure returns()
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorCallerSession) OnInstall(data []byte) error {
	return _WebAuthnAndECDSAValidator.Contract.OnInstall(&_WebAuthnAndECDSAValidator.CallOpts, data)
}

// OnUninstall is a free data retrieval call binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) pure returns()
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorCaller) OnUninstall(opts *bind.CallOpts, data []byte) error {
	var out []interface{}
	err := _WebAuthnAndECDSAValidator.contract.Call(opts, &out, "onUninstall", data)

	if err != nil {
		return err
	}

	return err

}

// OnUninstall is a free data retrieval call binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) pure returns()
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorSession) OnUninstall(data []byte) error {
	return _WebAuthnAndECDSAValidator.Contract.OnUninstall(&_WebAuthnAndECDSAValidator.CallOpts, data)
}

// OnUninstall is a free data retrieval call binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) pure returns()
func (_WebAuthnAndECDSAValidator *WebAuthnAndECDSAValidatorCallerSession) OnUninstall(data []byte) error {
	return _WebAuthnAndECDSAValidator.Contract.OnUninstall(&_WebAuthnAndECDSAValidator.CallOpts, data)
}
