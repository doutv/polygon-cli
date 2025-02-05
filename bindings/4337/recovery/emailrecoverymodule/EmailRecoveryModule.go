// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package emailrecoverymodule

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

// EmailRecoveryModuleMetaData contains all meta data concerning the EmailRecoveryModule contract.
var EmailRecoveryModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zkEmailVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDkimKeyHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEmailHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recoveringAccount\",\"type\":\"address\"}],\"name\":\"InvalidRecoveringAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"emailHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"EmailHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"OracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"}],\"name\":\"SignerUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RECOVERY_FEE_BASE_COST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZK_EMAIL_VERIFIER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getEmailHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_emailHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lastUpdatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleTypeId\",\"type\":\"uint256\"}],\"name\":\"isModuleType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_emailHash\",\"type\":\"bytes32\"}],\"name\":\"updateEmailHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a03461010d57601f610cea38819003918201601f19168301916001600160401b038311848410176101125780849260809460405283398101031261010d5761004781610128565b9061005460208201610128565b606061006260408401610128565b926001600160a01b039182916100789101610128565b169182156100f457816000549560018060a01b03199480868916176000558260405198167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a31683600154161760015516906002541617600255608052610bad908161013d823960805181818161019201526105bd0152f35b604051631e4fbdf760e01b815260006004820152602490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b038216820361010d5756fe6040608081526004908136101561001557600080fd5b600091823560e01c806329a27b1d1461097f578063379f4e66146104835780633adde5bf14610443578063417d127a14610426578063580f6da2146103ee5780636d61fe701461039c578063715018a61461034257806379502c55146103195780637adbf97314610298578063833b1fce1461026f5780638a91b0e3146101ed5780638da5cb5b146101c5578063b4b5ca8d1461017d578063ecd05961146101595763f2fde38b146100c657600080fd5b34610155576020366003190112610155576100df61099d565b906100e8610b4b565b6001600160a01b0391821692831561013f575050600054826bffffffffffffffffffffffff60a01b821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a380f35b51631e4fbdf760e01b8152908101849052602490fd5b8280fd5b50823461017a57602036600319011261017a57506002602092519135148152f35b80fd5b5050346101c157816003193601126101c157517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b5050346101c157816003193601126101c157905490516001600160a01b039091168152602090f35b5090346101555760203660031901126101555781359067ffffffffffffffff821161026b577ffc50c12be028a09bd30ac0b3f470f1a81c4d72e1043d9358931da4dbf162b40192610243606093369083016109b3565b505033855260205260006001828620828155015583815191338352816020840152820152a180f35b8380fd5b5050346101c157816003193601126101c15760025490516001600160a01b039091168152602090f35b509034610155576020366003190112610155576102b361099d565b6102bb610b4b565b6001600160a01b031691821561030b5750600280546001600160a01b03191683179055519081527f3df77beb5db05fcdd70a30fc8adf3f83f9501b68579455adbd100b818094039490602090a180f35b905163e6c4247b60e01b8152fd5b5050346101c157816003193601126101c15760015490516001600160a01b039091168152602090f35b833461017a578060031936011261017a5761035b610b4b565b600080546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b8382346101c15760203660031901126101c157803567ffffffffffffffff8111610155576020916103cf913691016109b3565b90809291810103126103e9576103e6903533610ab2565b80f35b600080fd5b5050346101c15760203660031901126101c15760209181906001600160a01b0361041661099d565b1681526003845220549051908152f35b5050346101c157816003193601126101c15760209051619b778152f35b503461015557602036600319011261015557909182916001600160a01b0361046961099d565b168252602052206001815491015482519182526020820152f35b508290346101c157826003193601126101c15761049e61099d565b9267ffffffffffffffff60243581811161097b576104bf90369085016109b3565b919060018060a01b0391826001541691855180936396a5d35b60e01b8252338983015281602460209687935afa9081156107dc578991610946575b5015610936575a9481016080828203126109325761051782610a3a565b9084830135928881013585811161092457810182601f820112156109245780359061054182610a4e565b9361054e8c519586610a0b565b8285528883830101116108335761057e60608998979695948f8f9b8682978f978301838b01378801015201610a3a565b9d6105b88c519485938493639967af6760e01b855216809b8401528b87166024840152876044840152608060648401526084830190610a8d565b03818a7f0000000000000000000000000000000000000000000000000000000000000000165afa801561092857908b959493929186879088938993610878575b501561086857878c8f8b815260038a5220541015610858578b8e8e8b8252895220540361084857858c9260248d93845161064d8161063f8782019488865289830190610a8d565b03601f198101835282610a0b565b519020948c6002541694519485938492631b450dbb60e11b84528301525afa90811561083e578d9161080d575b50036107fd578851918483015284898301526fffffffffffffffffffffffffffffffff6060830152606082526080820192828410908411176107ea57828952853b156107e657631bcfa73360e11b83528616608482015260a4810188905289908290607f19906106ed60c4820182610a8d565b03018183885af180156107dc576107c8575b506003908389525284872055818716610716578580f35b619b77808401938481116107b5575a9003019283116107a2573a8302928084043a14901517156107a2578596813b1561079e5786604492819587519889968795631f58bb2d60e11b8752169085015260248401525af19081156107955750610781575b808080808580f35b61078a906109e1565b61017a578082610779565b513d84823e3d90fd5b8680fd5b634e487b7160e01b865260118552602486fd5b634e487b7160e01b885260118752602488fd5b976107d5600392996109e1565b97906106ff565b87513d8b823e3d90fd5b8a80fd5b634e487b7160e01b8b5260418a5260248bfd5b8851636628d45760e01b81528a90fd5b90508581813d8311610837575b6108248183610a0b565b8101031261083357518e61067a565b8c80fd5b503d61081a565b8b513d8f823e3d90fd5b8a51637e02616560e11b81528c90fd5b8b51635376c04960e01b81528d90fd5b8b516309bde33960e01b81528d90fd5b949596979850505050503d808c833e6108918183610a0b565b810160a082820312610924576108a682610a2d565b9086830151918b8401519160608501519460808101519089821161091b570181601f820112156109205790818e9a999897969594939251916108f36108ea84610a4e565b9c519c8d610a0b565b828c528a838301011161091b5790610910918a808d019101610a6a565b9190929791386105f8565b508f80fd5b8f80fd5b8b80fd5b89513d8d823e3d90fd5b8880fd5b8551632057875960e21b81528790fd5b90508381813d8311610974575b61095d8183610a0b565b810103126109325761096e90610a2d565b8a6104fa565b503d610953565b8480fd5b8382346101c15760203660031901126101c1576103e6903533610ab2565b600435906001600160a01b03821682036103e957565b9181601f840112156103e95782359167ffffffffffffffff83116103e957602083818601950101116103e957565b67ffffffffffffffff81116109f557604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff8211176109f557604052565b519081151582036103e957565b35906001600160a01b03821682036103e957565b67ffffffffffffffff81116109f557601f01601f191660200190565b60005b838110610a7d5750506000910152565b8181015183820152602001610a6d565b90602091610aa681518092818552858086019101610a6a565b601f01601f1916010190565b908015610b39576040918251928084019284841067ffffffffffffffff8511176109f5577ffc50c12be028a09bd30ac0b3f470f1a81c4d72e1043d9358931da4dbf162b40194606094835281815260016020820194428652818060a01b031694856000526004602052846000209251835551910155815192835260208301524290820152a1565b604051637e02616560e11b8152600490fd5b6000546001600160a01b03163303610b5f57565b60405163118cdaa760e01b8152336004820152602490fdfea2646970667358221220b386be834b9e55afc522319f689311fbff4560289fd8a759b177a26cc1c40eee64736f6c63430008190033",
}

// EmailRecoveryModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use EmailRecoveryModuleMetaData.ABI instead.
var EmailRecoveryModuleABI = EmailRecoveryModuleMetaData.ABI

// EmailRecoveryModuleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EmailRecoveryModuleMetaData.Bin instead.
var EmailRecoveryModuleBin = EmailRecoveryModuleMetaData.Bin

// DeployEmailRecoveryModule deploys a new Ethereum contract, binding an instance of EmailRecoveryModule to it.
func DeployEmailRecoveryModule(auth *bind.TransactOpts, backend bind.ContractBackend, _config common.Address, _oracle common.Address, _zkEmailVerifier common.Address, _owner common.Address) (common.Address, *types.Transaction, *EmailRecoveryModule, error) {
	parsed, err := EmailRecoveryModuleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EmailRecoveryModuleBin), backend, _config, _oracle, _zkEmailVerifier, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EmailRecoveryModule{EmailRecoveryModuleCaller: EmailRecoveryModuleCaller{contract: contract}, EmailRecoveryModuleTransactor: EmailRecoveryModuleTransactor{contract: contract}, EmailRecoveryModuleFilterer: EmailRecoveryModuleFilterer{contract: contract}}, nil
}

// EmailRecoveryModule is an auto generated Go binding around an Ethereum contract.
type EmailRecoveryModule struct {
	EmailRecoveryModuleCaller     // Read-only binding to the contract
	EmailRecoveryModuleTransactor // Write-only binding to the contract
	EmailRecoveryModuleFilterer   // Log filterer for contract events
}

// EmailRecoveryModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type EmailRecoveryModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmailRecoveryModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EmailRecoveryModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmailRecoveryModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EmailRecoveryModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmailRecoveryModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EmailRecoveryModuleSession struct {
	Contract     *EmailRecoveryModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EmailRecoveryModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EmailRecoveryModuleCallerSession struct {
	Contract *EmailRecoveryModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// EmailRecoveryModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EmailRecoveryModuleTransactorSession struct {
	Contract     *EmailRecoveryModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// EmailRecoveryModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type EmailRecoveryModuleRaw struct {
	Contract *EmailRecoveryModule // Generic contract binding to access the raw methods on
}

// EmailRecoveryModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EmailRecoveryModuleCallerRaw struct {
	Contract *EmailRecoveryModuleCaller // Generic read-only contract binding to access the raw methods on
}

// EmailRecoveryModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EmailRecoveryModuleTransactorRaw struct {
	Contract *EmailRecoveryModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEmailRecoveryModule creates a new instance of EmailRecoveryModule, bound to a specific deployed contract.
func NewEmailRecoveryModule(address common.Address, backend bind.ContractBackend) (*EmailRecoveryModule, error) {
	contract, err := bindEmailRecoveryModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EmailRecoveryModule{EmailRecoveryModuleCaller: EmailRecoveryModuleCaller{contract: contract}, EmailRecoveryModuleTransactor: EmailRecoveryModuleTransactor{contract: contract}, EmailRecoveryModuleFilterer: EmailRecoveryModuleFilterer{contract: contract}}, nil
}

// NewEmailRecoveryModuleCaller creates a new read-only instance of EmailRecoveryModule, bound to a specific deployed contract.
func NewEmailRecoveryModuleCaller(address common.Address, caller bind.ContractCaller) (*EmailRecoveryModuleCaller, error) {
	contract, err := bindEmailRecoveryModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EmailRecoveryModuleCaller{contract: contract}, nil
}

// NewEmailRecoveryModuleTransactor creates a new write-only instance of EmailRecoveryModule, bound to a specific deployed contract.
func NewEmailRecoveryModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*EmailRecoveryModuleTransactor, error) {
	contract, err := bindEmailRecoveryModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EmailRecoveryModuleTransactor{contract: contract}, nil
}

// NewEmailRecoveryModuleFilterer creates a new log filterer instance of EmailRecoveryModule, bound to a specific deployed contract.
func NewEmailRecoveryModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*EmailRecoveryModuleFilterer, error) {
	contract, err := bindEmailRecoveryModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EmailRecoveryModuleFilterer{contract: contract}, nil
}

// bindEmailRecoveryModule binds a generic wrapper to an already deployed contract.
func bindEmailRecoveryModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EmailRecoveryModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmailRecoveryModule *EmailRecoveryModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmailRecoveryModule.Contract.EmailRecoveryModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmailRecoveryModule *EmailRecoveryModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.EmailRecoveryModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmailRecoveryModule *EmailRecoveryModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.EmailRecoveryModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmailRecoveryModule *EmailRecoveryModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmailRecoveryModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmailRecoveryModule *EmailRecoveryModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmailRecoveryModule *EmailRecoveryModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.contract.Transact(opts, method, params...)
}

// RECOVERYFEEBASECOST is a free data retrieval call binding the contract method 0x417d127a.
//
// Solidity: function RECOVERY_FEE_BASE_COST() view returns(uint256)
func (_EmailRecoveryModule *EmailRecoveryModuleCaller) RECOVERYFEEBASECOST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmailRecoveryModule.contract.Call(opts, &out, "RECOVERY_FEE_BASE_COST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RECOVERYFEEBASECOST is a free data retrieval call binding the contract method 0x417d127a.
//
// Solidity: function RECOVERY_FEE_BASE_COST() view returns(uint256)
func (_EmailRecoveryModule *EmailRecoveryModuleSession) RECOVERYFEEBASECOST() (*big.Int, error) {
	return _EmailRecoveryModule.Contract.RECOVERYFEEBASECOST(&_EmailRecoveryModule.CallOpts)
}

// RECOVERYFEEBASECOST is a free data retrieval call binding the contract method 0x417d127a.
//
// Solidity: function RECOVERY_FEE_BASE_COST() view returns(uint256)
func (_EmailRecoveryModule *EmailRecoveryModuleCallerSession) RECOVERYFEEBASECOST() (*big.Int, error) {
	return _EmailRecoveryModule.Contract.RECOVERYFEEBASECOST(&_EmailRecoveryModule.CallOpts)
}

// ZKEMAILVERIFIER is a free data retrieval call binding the contract method 0xb4b5ca8d.
//
// Solidity: function ZK_EMAIL_VERIFIER() view returns(address)
func (_EmailRecoveryModule *EmailRecoveryModuleCaller) ZKEMAILVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EmailRecoveryModule.contract.Call(opts, &out, "ZK_EMAIL_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZKEMAILVERIFIER is a free data retrieval call binding the contract method 0xb4b5ca8d.
//
// Solidity: function ZK_EMAIL_VERIFIER() view returns(address)
func (_EmailRecoveryModule *EmailRecoveryModuleSession) ZKEMAILVERIFIER() (common.Address, error) {
	return _EmailRecoveryModule.Contract.ZKEMAILVERIFIER(&_EmailRecoveryModule.CallOpts)
}

// ZKEMAILVERIFIER is a free data retrieval call binding the contract method 0xb4b5ca8d.
//
// Solidity: function ZK_EMAIL_VERIFIER() view returns(address)
func (_EmailRecoveryModule *EmailRecoveryModuleCallerSession) ZKEMAILVERIFIER() (common.Address, error) {
	return _EmailRecoveryModule.Contract.ZKEMAILVERIFIER(&_EmailRecoveryModule.CallOpts)
}

// AccountNonces is a free data retrieval call binding the contract method 0x580f6da2.
//
// Solidity: function accountNonces(address ) view returns(uint256)
func (_EmailRecoveryModule *EmailRecoveryModuleCaller) AccountNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EmailRecoveryModule.contract.Call(opts, &out, "accountNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountNonces is a free data retrieval call binding the contract method 0x580f6da2.
//
// Solidity: function accountNonces(address ) view returns(uint256)
func (_EmailRecoveryModule *EmailRecoveryModuleSession) AccountNonces(arg0 common.Address) (*big.Int, error) {
	return _EmailRecoveryModule.Contract.AccountNonces(&_EmailRecoveryModule.CallOpts, arg0)
}

// AccountNonces is a free data retrieval call binding the contract method 0x580f6da2.
//
// Solidity: function accountNonces(address ) view returns(uint256)
func (_EmailRecoveryModule *EmailRecoveryModuleCallerSession) AccountNonces(arg0 common.Address) (*big.Int, error) {
	return _EmailRecoveryModule.Contract.AccountNonces(&_EmailRecoveryModule.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_EmailRecoveryModule *EmailRecoveryModuleCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EmailRecoveryModule.contract.Call(opts, &out, "config")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_EmailRecoveryModule *EmailRecoveryModuleSession) Config() (common.Address, error) {
	return _EmailRecoveryModule.Contract.Config(&_EmailRecoveryModule.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_EmailRecoveryModule *EmailRecoveryModuleCallerSession) Config() (common.Address, error) {
	return _EmailRecoveryModule.Contract.Config(&_EmailRecoveryModule.CallOpts)
}

// GetEmailHash is a free data retrieval call binding the contract method 0x3adde5bf.
//
// Solidity: function getEmailHash(address _account) view returns(bytes32 _emailHash, uint256 _lastUpdatedAt)
func (_EmailRecoveryModule *EmailRecoveryModuleCaller) GetEmailHash(opts *bind.CallOpts, _account common.Address) (struct {
	EmailHash     [32]byte
	LastUpdatedAt *big.Int
}, error) {
	var out []interface{}
	err := _EmailRecoveryModule.contract.Call(opts, &out, "getEmailHash", _account)

	outstruct := new(struct {
		EmailHash     [32]byte
		LastUpdatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EmailHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.LastUpdatedAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetEmailHash is a free data retrieval call binding the contract method 0x3adde5bf.
//
// Solidity: function getEmailHash(address _account) view returns(bytes32 _emailHash, uint256 _lastUpdatedAt)
func (_EmailRecoveryModule *EmailRecoveryModuleSession) GetEmailHash(_account common.Address) (struct {
	EmailHash     [32]byte
	LastUpdatedAt *big.Int
}, error) {
	return _EmailRecoveryModule.Contract.GetEmailHash(&_EmailRecoveryModule.CallOpts, _account)
}

// GetEmailHash is a free data retrieval call binding the contract method 0x3adde5bf.
//
// Solidity: function getEmailHash(address _account) view returns(bytes32 _emailHash, uint256 _lastUpdatedAt)
func (_EmailRecoveryModule *EmailRecoveryModuleCallerSession) GetEmailHash(_account common.Address) (struct {
	EmailHash     [32]byte
	LastUpdatedAt *big.Int
}, error) {
	return _EmailRecoveryModule.Contract.GetEmailHash(&_EmailRecoveryModule.CallOpts, _account)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_EmailRecoveryModule *EmailRecoveryModuleCaller) GetOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EmailRecoveryModule.contract.Call(opts, &out, "getOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_EmailRecoveryModule *EmailRecoveryModuleSession) GetOracle() (common.Address, error) {
	return _EmailRecoveryModule.Contract.GetOracle(&_EmailRecoveryModule.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_EmailRecoveryModule *EmailRecoveryModuleCallerSession) GetOracle() (common.Address, error) {
	return _EmailRecoveryModule.Contract.GetOracle(&_EmailRecoveryModule.CallOpts)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_EmailRecoveryModule *EmailRecoveryModuleCaller) IsModuleType(opts *bind.CallOpts, moduleTypeId *big.Int) (bool, error) {
	var out []interface{}
	err := _EmailRecoveryModule.contract.Call(opts, &out, "isModuleType", moduleTypeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_EmailRecoveryModule *EmailRecoveryModuleSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _EmailRecoveryModule.Contract.IsModuleType(&_EmailRecoveryModule.CallOpts, moduleTypeId)
}

// IsModuleType is a free data retrieval call binding the contract method 0xecd05961.
//
// Solidity: function isModuleType(uint256 moduleTypeId) pure returns(bool)
func (_EmailRecoveryModule *EmailRecoveryModuleCallerSession) IsModuleType(moduleTypeId *big.Int) (bool, error) {
	return _EmailRecoveryModule.Contract.IsModuleType(&_EmailRecoveryModule.CallOpts, moduleTypeId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EmailRecoveryModule *EmailRecoveryModuleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EmailRecoveryModule.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EmailRecoveryModule *EmailRecoveryModuleSession) Owner() (common.Address, error) {
	return _EmailRecoveryModule.Contract.Owner(&_EmailRecoveryModule.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EmailRecoveryModule *EmailRecoveryModuleCallerSession) Owner() (common.Address, error) {
	return _EmailRecoveryModule.Contract.Owner(&_EmailRecoveryModule.CallOpts)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactor) OnInstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _EmailRecoveryModule.contract.Transact(opts, "onInstall", data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.OnInstall(&_EmailRecoveryModule.TransactOpts, data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.OnInstall(&_EmailRecoveryModule.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes ) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactor) OnUninstall(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _EmailRecoveryModule.contract.Transact(opts, "onUninstall", arg0)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes ) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleSession) OnUninstall(arg0 []byte) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.OnUninstall(&_EmailRecoveryModule.TransactOpts, arg0)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes ) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactorSession) OnUninstall(arg0 []byte) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.OnUninstall(&_EmailRecoveryModule.TransactOpts, arg0)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address account, bytes data) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactor) Recover(opts *bind.TransactOpts, account common.Address, data []byte) (*types.Transaction, error) {
	return _EmailRecoveryModule.contract.Transact(opts, "recover", account, data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address account, bytes data) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleSession) Recover(account common.Address, data []byte) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.Recover(&_EmailRecoveryModule.TransactOpts, account, data)
}

// Recover is a paid mutator transaction binding the contract method 0x379f4e66.
//
// Solidity: function recover(address account, bytes data) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactorSession) Recover(account common.Address, data []byte) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.Recover(&_EmailRecoveryModule.TransactOpts, account, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmailRecoveryModule.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EmailRecoveryModule *EmailRecoveryModuleSession) RenounceOwnership() (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.RenounceOwnership(&_EmailRecoveryModule.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.RenounceOwnership(&_EmailRecoveryModule.TransactOpts)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _newOracle) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactor) SetOracle(opts *bind.TransactOpts, _newOracle common.Address) (*types.Transaction, error) {
	return _EmailRecoveryModule.contract.Transact(opts, "setOracle", _newOracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _newOracle) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleSession) SetOracle(_newOracle common.Address) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.SetOracle(&_EmailRecoveryModule.TransactOpts, _newOracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _newOracle) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactorSession) SetOracle(_newOracle common.Address) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.SetOracle(&_EmailRecoveryModule.TransactOpts, _newOracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EmailRecoveryModule.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.TransferOwnership(&_EmailRecoveryModule.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.TransferOwnership(&_EmailRecoveryModule.TransactOpts, newOwner)
}

// UpdateEmailHash is a paid mutator transaction binding the contract method 0x29a27b1d.
//
// Solidity: function updateEmailHash(bytes32 _emailHash) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactor) UpdateEmailHash(opts *bind.TransactOpts, _emailHash [32]byte) (*types.Transaction, error) {
	return _EmailRecoveryModule.contract.Transact(opts, "updateEmailHash", _emailHash)
}

// UpdateEmailHash is a paid mutator transaction binding the contract method 0x29a27b1d.
//
// Solidity: function updateEmailHash(bytes32 _emailHash) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleSession) UpdateEmailHash(_emailHash [32]byte) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.UpdateEmailHash(&_EmailRecoveryModule.TransactOpts, _emailHash)
}

// UpdateEmailHash is a paid mutator transaction binding the contract method 0x29a27b1d.
//
// Solidity: function updateEmailHash(bytes32 _emailHash) returns()
func (_EmailRecoveryModule *EmailRecoveryModuleTransactorSession) UpdateEmailHash(_emailHash [32]byte) (*types.Transaction, error) {
	return _EmailRecoveryModule.Contract.UpdateEmailHash(&_EmailRecoveryModule.TransactOpts, _emailHash)
}

// EmailRecoveryModuleEmailHashUpdatedIterator is returned from FilterEmailHashUpdated and is used to iterate over the raw logs and unpacked data for EmailHashUpdated events raised by the EmailRecoveryModule contract.
type EmailRecoveryModuleEmailHashUpdatedIterator struct {
	Event *EmailRecoveryModuleEmailHashUpdated // Event containing the contract specifics and raw log

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
func (it *EmailRecoveryModuleEmailHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmailRecoveryModuleEmailHashUpdated)
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
		it.Event = new(EmailRecoveryModuleEmailHashUpdated)
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
func (it *EmailRecoveryModuleEmailHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmailRecoveryModuleEmailHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmailRecoveryModuleEmailHashUpdated represents a EmailHashUpdated event raised by the EmailRecoveryModule contract.
type EmailRecoveryModuleEmailHashUpdated struct {
	Account   common.Address
	EmailHash [32]byte
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEmailHashUpdated is a free log retrieval operation binding the contract event 0xfc50c12be028a09bd30ac0b3f470f1a81c4d72e1043d9358931da4dbf162b401.
//
// Solidity: event EmailHashUpdated(address account, bytes32 emailHash, uint256 updatedAt)
func (_EmailRecoveryModule *EmailRecoveryModuleFilterer) FilterEmailHashUpdated(opts *bind.FilterOpts) (*EmailRecoveryModuleEmailHashUpdatedIterator, error) {

	logs, sub, err := _EmailRecoveryModule.contract.FilterLogs(opts, "EmailHashUpdated")
	if err != nil {
		return nil, err
	}
	return &EmailRecoveryModuleEmailHashUpdatedIterator{contract: _EmailRecoveryModule.contract, event: "EmailHashUpdated", logs: logs, sub: sub}, nil
}

// WatchEmailHashUpdated is a free log subscription operation binding the contract event 0xfc50c12be028a09bd30ac0b3f470f1a81c4d72e1043d9358931da4dbf162b401.
//
// Solidity: event EmailHashUpdated(address account, bytes32 emailHash, uint256 updatedAt)
func (_EmailRecoveryModule *EmailRecoveryModuleFilterer) WatchEmailHashUpdated(opts *bind.WatchOpts, sink chan<- *EmailRecoveryModuleEmailHashUpdated) (event.Subscription, error) {

	logs, sub, err := _EmailRecoveryModule.contract.WatchLogs(opts, "EmailHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmailRecoveryModuleEmailHashUpdated)
				if err := _EmailRecoveryModule.contract.UnpackLog(event, "EmailHashUpdated", log); err != nil {
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

// ParseEmailHashUpdated is a log parse operation binding the contract event 0xfc50c12be028a09bd30ac0b3f470f1a81c4d72e1043d9358931da4dbf162b401.
//
// Solidity: event EmailHashUpdated(address account, bytes32 emailHash, uint256 updatedAt)
func (_EmailRecoveryModule *EmailRecoveryModuleFilterer) ParseEmailHashUpdated(log types.Log) (*EmailRecoveryModuleEmailHashUpdated, error) {
	event := new(EmailRecoveryModuleEmailHashUpdated)
	if err := _EmailRecoveryModule.contract.UnpackLog(event, "EmailHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmailRecoveryModuleOracleUpdatedIterator is returned from FilterOracleUpdated and is used to iterate over the raw logs and unpacked data for OracleUpdated events raised by the EmailRecoveryModule contract.
type EmailRecoveryModuleOracleUpdatedIterator struct {
	Event *EmailRecoveryModuleOracleUpdated // Event containing the contract specifics and raw log

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
func (it *EmailRecoveryModuleOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmailRecoveryModuleOracleUpdated)
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
		it.Event = new(EmailRecoveryModuleOracleUpdated)
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
func (it *EmailRecoveryModuleOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmailRecoveryModuleOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmailRecoveryModuleOracleUpdated represents a OracleUpdated event raised by the EmailRecoveryModule contract.
type EmailRecoveryModuleOracleUpdated struct {
	NewOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleUpdated is a free log retrieval operation binding the contract event 0x3df77beb5db05fcdd70a30fc8adf3f83f9501b68579455adbd100b8180940394.
//
// Solidity: event OracleUpdated(address newOracle)
func (_EmailRecoveryModule *EmailRecoveryModuleFilterer) FilterOracleUpdated(opts *bind.FilterOpts) (*EmailRecoveryModuleOracleUpdatedIterator, error) {

	logs, sub, err := _EmailRecoveryModule.contract.FilterLogs(opts, "OracleUpdated")
	if err != nil {
		return nil, err
	}
	return &EmailRecoveryModuleOracleUpdatedIterator{contract: _EmailRecoveryModule.contract, event: "OracleUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleUpdated is a free log subscription operation binding the contract event 0x3df77beb5db05fcdd70a30fc8adf3f83f9501b68579455adbd100b8180940394.
//
// Solidity: event OracleUpdated(address newOracle)
func (_EmailRecoveryModule *EmailRecoveryModuleFilterer) WatchOracleUpdated(opts *bind.WatchOpts, sink chan<- *EmailRecoveryModuleOracleUpdated) (event.Subscription, error) {

	logs, sub, err := _EmailRecoveryModule.contract.WatchLogs(opts, "OracleUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmailRecoveryModuleOracleUpdated)
				if err := _EmailRecoveryModule.contract.UnpackLog(event, "OracleUpdated", log); err != nil {
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

// ParseOracleUpdated is a log parse operation binding the contract event 0x3df77beb5db05fcdd70a30fc8adf3f83f9501b68579455adbd100b8180940394.
//
// Solidity: event OracleUpdated(address newOracle)
func (_EmailRecoveryModule *EmailRecoveryModuleFilterer) ParseOracleUpdated(log types.Log) (*EmailRecoveryModuleOracleUpdated, error) {
	event := new(EmailRecoveryModuleOracleUpdated)
	if err := _EmailRecoveryModule.contract.UnpackLog(event, "OracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmailRecoveryModuleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EmailRecoveryModule contract.
type EmailRecoveryModuleOwnershipTransferredIterator struct {
	Event *EmailRecoveryModuleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EmailRecoveryModuleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmailRecoveryModuleOwnershipTransferred)
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
		it.Event = new(EmailRecoveryModuleOwnershipTransferred)
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
func (it *EmailRecoveryModuleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmailRecoveryModuleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmailRecoveryModuleOwnershipTransferred represents a OwnershipTransferred event raised by the EmailRecoveryModule contract.
type EmailRecoveryModuleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EmailRecoveryModule *EmailRecoveryModuleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EmailRecoveryModuleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EmailRecoveryModule.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EmailRecoveryModuleOwnershipTransferredIterator{contract: _EmailRecoveryModule.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EmailRecoveryModule *EmailRecoveryModuleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EmailRecoveryModuleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EmailRecoveryModule.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmailRecoveryModuleOwnershipTransferred)
				if err := _EmailRecoveryModule.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EmailRecoveryModule *EmailRecoveryModuleFilterer) ParseOwnershipTransferred(log types.Log) (*EmailRecoveryModuleOwnershipTransferred, error) {
	event := new(EmailRecoveryModuleOwnershipTransferred)
	if err := _EmailRecoveryModule.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmailRecoveryModuleSignerUpdatedIterator is returned from FilterSignerUpdated and is used to iterate over the raw logs and unpacked data for SignerUpdated events raised by the EmailRecoveryModule contract.
type EmailRecoveryModuleSignerUpdatedIterator struct {
	Event *EmailRecoveryModuleSignerUpdated // Event containing the contract specifics and raw log

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
func (it *EmailRecoveryModuleSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmailRecoveryModuleSignerUpdated)
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
		it.Event = new(EmailRecoveryModuleSignerUpdated)
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
func (it *EmailRecoveryModuleSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmailRecoveryModuleSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmailRecoveryModuleSignerUpdated represents a SignerUpdated event raised by the EmailRecoveryModule contract.
type EmailRecoveryModuleSignerUpdated struct {
	NewSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignerUpdated is a free log retrieval operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address newSigner)
func (_EmailRecoveryModule *EmailRecoveryModuleFilterer) FilterSignerUpdated(opts *bind.FilterOpts) (*EmailRecoveryModuleSignerUpdatedIterator, error) {

	logs, sub, err := _EmailRecoveryModule.contract.FilterLogs(opts, "SignerUpdated")
	if err != nil {
		return nil, err
	}
	return &EmailRecoveryModuleSignerUpdatedIterator{contract: _EmailRecoveryModule.contract, event: "SignerUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerUpdated is a free log subscription operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address newSigner)
func (_EmailRecoveryModule *EmailRecoveryModuleFilterer) WatchSignerUpdated(opts *bind.WatchOpts, sink chan<- *EmailRecoveryModuleSignerUpdated) (event.Subscription, error) {

	logs, sub, err := _EmailRecoveryModule.contract.WatchLogs(opts, "SignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmailRecoveryModuleSignerUpdated)
				if err := _EmailRecoveryModule.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
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

// ParseSignerUpdated is a log parse operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address newSigner)
func (_EmailRecoveryModule *EmailRecoveryModuleFilterer) ParseSignerUpdated(log types.Log) (*EmailRecoveryModuleSignerUpdated, error) {
	event := new(EmailRecoveryModuleSignerUpdated)
	if err := _EmailRecoveryModule.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
