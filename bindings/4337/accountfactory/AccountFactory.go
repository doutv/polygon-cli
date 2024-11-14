// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accountfactory

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

// AccountFactoryMetaData contains all meta data concerning the AccountFactory contract.
var AccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"config\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBundler\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"}],\"name\":\"InvalidSingleton\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_initializer\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONFIG\",\"outputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"computeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60c034610121576001600160401b0390601f61110338819003918201601f1916830191848311848410176101265780849260209460405283398101031261012157516001600160a01b03811690819003610121573060805260a0527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c1661010f5780808316036100ca575b604051610fc6908161013d82396080518181816103a10152610481015260a051818181610121015281816106c501526108050152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1388080610094565b60405163f92ee8a960e01b8152600490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe608060408181526004918236101561001657600080fd5b600092833560e01c918263296601cd1461078a5750816336b5aa2d146106705781634f1ef2861461040557816352d1902d1461038c578163715018a6146103215781638da5cb5b146102eb578163ad3cb1cc1461028e578163c4d66de81461015057508063d92e82e41461010d578063f2fde38b146100dd5763fbcbc0f11461009e57600080fd5b346100d95760203660031901126100d95760209160ff9082906001600160a01b036100c7610a59565b16815280855220541690519015158152f35b5080fd5b823461010a57602036600319011261010a576101076100fa610a59565b610102610bb4565b610b1f565b80f35b80fd5b50346100d957816003193601126100d957517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b90503461028a57602036600319011261028a5761016b610a59565b907ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0091825460ff81861c16159267ffffffffffffffff821680159081610282575b6001149081610278575b15908161026f575b50610261575067ffffffffffffffff19811660011784556101f1919083610242575b506101e9610bed565b610102610bed565b6101f9610bed565b610201578280f35b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808280f35b68ffffffffffffffffff191668010000000000000001178455386101e0565b855163f92ee8a960e01b8152fd5b905015386101be565b303b1591506101b6565b8591506101ac565b8280fd5b5050346100d957816003193601126100d957806102dd91516102af81610a74565b6005815260208101640352e302e360dc1b815282519384926020845251809281602086015285850190610ae4565b601f01601f19168101030190f35b5050346100d957816003193601126100d957600080516020610f718339815191525490516001600160a01b039091168152602090f35b833461010a578060031936011261010a5761033a610bb4565b600080516020610f7183398151915280546001600160a01b031981169091556000906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b82843461010a578060031936011261010a57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036103f857602090517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b5163703e46dd60e11b8152fd5b9180915060031936011261028a5761041b610a59565b90602493843567ffffffffffffffff81116100d957366023820112156100d9578085013561044881610ac8565b9461045585519687610aa6565b81865260209182870193368a838301011161066c578186928b8693018737880101526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811630811490811561063e575b5061062e576104ba610bb4565b81169585516352d1902d60e01b815283818a818b5afa8691816105fb575b506104f4575050505050505191634c9c8ce360e01b8352820152fd5b9088888894938c7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc918281036105e65750853b156105d2575080546001600160a01b031916821790558451889392917fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8580a28251156105b45750506105a69582915190845af4913d156105aa573d61059861058f82610ac8565b92519283610aa6565b81528581943d92013e610c2e565b5080f35b5060609250610c2e565b9550955050505050346105c657505080f35b63b398979f60e01b8152fd5b8651634c9c8ce360e01b8152808501849052fd5b8751632a87526960e21b815280860191909152fd5b9091508481813d8311610627575b6106138183610aa6565b81010312610623575190386104d8565b8680fd5b503d610609565b855163703e46dd60e11b81528890fd5b9050817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc54161415386104ad565b8580fd5b9050823461010a578260031936011261010a575061068c610a59565b5081516020926102df6106a185820184610aa6565b80835284830190610c9282396107196107258684519581870160018060a01b0397887f00000000000000000000000000000000000000000000000000000000000000001682528381526106f381610a74565b875195869361070a868601998a9251928391610ae4565b84019151809386840190610ae4565b01038084520182610aa6565b519020908051908582019260ff60f81b84523060601b60218401526024356035840152605583015260558252608082019482861067ffffffffffffffff8711176107755750849052519020168152f35b604190634e487b7160e01b6000525260246000fd5b8484833461028a57606036600319011261028a576107a6610a59565b60249081359267ffffffffffffffff9586851161010a573660238601121561010a5784820135938785116100d95780860195818636920101116100d957634ef3c6d960e11b8952328984015260209860443593906001600160a01b03907f00000000000000000000000000000000000000000000000000000000000000008216908c818681855afa908115610a22578691610a3c575b5015610a2c578951635eb43cb560e01b8152968216838801819052968c818681855afa908115610a225786916109f5575b50156109df578951906102df80830191908d8311848410176109cd579280928f928a95610c928439815203019086f580156109c3571698893b156109bf579087929184888c8c838b6108dd83519a8b968795869463347d5e2560e21b86528c8601528c8501526044840191610b93565b03925af180156109b557918b9897969593918b959361095a575b505050509561094c9187877fa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f7437989952808c5220600160ff198254161790558751948594855260608b8601526060850191610b93565b90868301520390a251908152f35b919395979996985091809450116109a4575050865290938693909290918690817fa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f743761094c8d6108f7565b634e487b7160e01b84526041905282fd5b8a513d87823e3d90fd5b8380fd5b89513d86823e3d90fd5b634e487b7160e01b8852604186528688fd5b50508751630f00f1a160e11b8152908101859052fd5b610a1591508d803d10610a1b575b610a0d8183610aa6565b810190610b07565b8d61086d565b503d610a03565b8b513d88823e3d90fd5b895163f8a0c54b60e01b81528390fd5b610a5391508d803d10610a1b57610a0d8183610aa6565b8d61083c565b600435906001600160a01b0382168203610a6f57565b600080fd5b6040810190811067ffffffffffffffff821117610a9057604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff821117610a9057604052565b67ffffffffffffffff8111610a9057601f01601f191660200190565b60005b838110610af75750506000910152565b8181015183820152602001610ae7565b90816020910312610a6f57518015158103610a6f5790565b6001600160a01b03908116908115610b7a57600080516020610f7183398151915280546001600160a01b031981168417909155167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b604051631e4fbdf760e01b815260006004820152602490fd5b908060209392818452848401376000828201840152601f01601f1916010190565b600080516020610f71833981519152546001600160a01b03163303610bd557565b60405163118cdaa760e01b8152336004820152602490fd5b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c1615610c1c57565b604051631afcd79f60e31b8152600490fd5b90610c555750805115610c4357805190602001fd5b604051630a12f52160e11b8152600490fd5b81511580610c88575b610c66575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b15610c5e56fe60a034606357601f6102df38819003918201601f19168301916001600160401b03831184841017606857808492602094604052833981010312606357516001600160a01b0381168103606357608052604051610260908161007f8239608051815050f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe60806040526004361015610024575b361561001f5734156101eb57600080fd5b6101eb565b6000803560e01c63d1f578941461003b575061000e565b346100af5760403660031901126100af576004356001600160a01b03811681036100ab576024359067ffffffffffffffff908183116100a757366023840112156100a75782600401359182116100a75736602483850101116100a75760246100a4930190610111565b80f35b8380fd5b5080fd5b80fd5b634e487b7160e01b600052604160045260246000fd5b6020808252825181830181905290939260005b8281106100fd57505060409293506000838284010152601f8019910116010190565b8181018601518482016040015285016100db565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8054929390926001600160a01b03166101d95760009382859455816040519283928337810184815203915af43d156101d15767ffffffffffffffff903d8281116101cc5760405192601f8201601f19908116603f01168401908111848210176101cc5760405282523d6000602084013e5b156101ab5750565b604051633018224d60e21b81529081906101c890600483016100c8565b0390fd5b6100b2565b6060906101a3565b60405163c28d69c760e01b8152600490fd5b600036818037808036817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc545af43d82803e15610226573d90f35b3d90fdfea264697066735822122058f0f8eb4a924f70bd7688d8b719e0c015fc0aae4f06b8a7f34ce91ee0d9998864736f6c634300081900339016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300a2646970667358221220972bf4e84137aff4ebed42e63948c3162782159ed9915a123ae2570f0023160d64736f6c63430008190033",
}

// AccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountFactoryMetaData.ABI instead.
var AccountFactoryABI = AccountFactoryMetaData.ABI

// AccountFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccountFactoryMetaData.Bin instead.
var AccountFactoryBin = AccountFactoryMetaData.Bin

// DeployAccountFactory deploys a new Ethereum contract, binding an instance of AccountFactory to it.
func DeployAccountFactory(auth *bind.TransactOpts, backend bind.ContractBackend, config common.Address) (common.Address, *types.Transaction, *AccountFactory, error) {
	parsed, err := AccountFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccountFactoryBin), backend, config)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountFactory{AccountFactoryCaller: AccountFactoryCaller{contract: contract}, AccountFactoryTransactor: AccountFactoryTransactor{contract: contract}, AccountFactoryFilterer: AccountFactoryFilterer{contract: contract}}, nil
}

// AccountFactory is an auto generated Go binding around an Ethereum contract.
type AccountFactory struct {
	AccountFactoryCaller     // Read-only binding to the contract
	AccountFactoryTransactor // Write-only binding to the contract
	AccountFactoryFilterer   // Log filterer for contract events
}

// AccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountFactorySession struct {
	Contract     *AccountFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountFactoryCallerSession struct {
	Contract *AccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountFactoryTransactorSession struct {
	Contract     *AccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountFactoryRaw struct {
	Contract *AccountFactory // Generic contract binding to access the raw methods on
}

// AccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountFactoryCallerRaw struct {
	Contract *AccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountFactoryTransactorRaw struct {
	Contract *AccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountFactory creates a new instance of AccountFactory, bound to a specific deployed contract.
func NewAccountFactory(address common.Address, backend bind.ContractBackend) (*AccountFactory, error) {
	contract, err := bindAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountFactory{AccountFactoryCaller: AccountFactoryCaller{contract: contract}, AccountFactoryTransactor: AccountFactoryTransactor{contract: contract}, AccountFactoryFilterer: AccountFactoryFilterer{contract: contract}}, nil
}

// NewAccountFactoryCaller creates a new read-only instance of AccountFactory, bound to a specific deployed contract.
func NewAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*AccountFactoryCaller, error) {
	contract, err := bindAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryCaller{contract: contract}, nil
}

// NewAccountFactoryTransactor creates a new write-only instance of AccountFactory, bound to a specific deployed contract.
func NewAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountFactoryTransactor, error) {
	contract, err := bindAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryTransactor{contract: contract}, nil
}

// NewAccountFactoryFilterer creates a new log filterer instance of AccountFactory, bound to a specific deployed contract.
func NewAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountFactoryFilterer, error) {
	contract, err := bindAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryFilterer{contract: contract}, nil
}

// bindAccountFactory binds a generic wrapper to an already deployed contract.
func bindAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountFactory *AccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountFactory.Contract.AccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountFactory *AccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountFactory.Contract.AccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountFactory *AccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountFactory.Contract.AccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountFactory *AccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountFactory *AccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountFactory *AccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountFactory.Contract.contract.Transact(opts, method, params...)
}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_AccountFactory *AccountFactoryCaller) CONFIG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "CONFIG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_AccountFactory *AccountFactorySession) CONFIG() (common.Address, error) {
	return _AccountFactory.Contract.CONFIG(&_AccountFactory.CallOpts)
}

// CONFIG is a free data retrieval call binding the contract method 0xd92e82e4.
//
// Solidity: function CONFIG() view returns(address)
func (_AccountFactory *AccountFactoryCallerSession) CONFIG() (common.Address, error) {
	return _AccountFactory.Contract.CONFIG(&_AccountFactory.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AccountFactory *AccountFactoryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AccountFactory *AccountFactorySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AccountFactory.Contract.UPGRADEINTERFACEVERSION(&_AccountFactory.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AccountFactory *AccountFactoryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AccountFactory.Contract.UPGRADEINTERFACEVERSION(&_AccountFactory.CallOpts)
}

// ComputeAddress is a free data retrieval call binding the contract method 0x36b5aa2d.
//
// Solidity: function computeAddress(address , uint256 _salt) view returns(address)
func (_AccountFactory *AccountFactoryCaller) ComputeAddress(opts *bind.CallOpts, arg0 common.Address, _salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "computeAddress", arg0, _salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeAddress is a free data retrieval call binding the contract method 0x36b5aa2d.
//
// Solidity: function computeAddress(address , uint256 _salt) view returns(address)
func (_AccountFactory *AccountFactorySession) ComputeAddress(arg0 common.Address, _salt *big.Int) (common.Address, error) {
	return _AccountFactory.Contract.ComputeAddress(&_AccountFactory.CallOpts, arg0, _salt)
}

// ComputeAddress is a free data retrieval call binding the contract method 0x36b5aa2d.
//
// Solidity: function computeAddress(address , uint256 _salt) view returns(address)
func (_AccountFactory *AccountFactoryCallerSession) ComputeAddress(arg0 common.Address, _salt *big.Int) (common.Address, error) {
	return _AccountFactory.Contract.ComputeAddress(&_AccountFactory.CallOpts, arg0, _salt)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address ) view returns(bool)
func (_AccountFactory *AccountFactoryCaller) GetAccount(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "getAccount", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address ) view returns(bool)
func (_AccountFactory *AccountFactorySession) GetAccount(arg0 common.Address) (bool, error) {
	return _AccountFactory.Contract.GetAccount(&_AccountFactory.CallOpts, arg0)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address ) view returns(bool)
func (_AccountFactory *AccountFactoryCallerSession) GetAccount(arg0 common.Address) (bool, error) {
	return _AccountFactory.Contract.GetAccount(&_AccountFactory.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountFactory *AccountFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountFactory *AccountFactorySession) Owner() (common.Address, error) {
	return _AccountFactory.Contract.Owner(&_AccountFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountFactory *AccountFactoryCallerSession) Owner() (common.Address, error) {
	return _AccountFactory.Contract.Owner(&_AccountFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AccountFactory *AccountFactoryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AccountFactory *AccountFactorySession) ProxiableUUID() ([32]byte, error) {
	return _AccountFactory.Contract.ProxiableUUID(&_AccountFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AccountFactory *AccountFactoryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AccountFactory.Contract.ProxiableUUID(&_AccountFactory.CallOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x296601cd.
//
// Solidity: function createAccount(address _implementation, bytes _initializer, uint256 _salt) returns(address)
func (_AccountFactory *AccountFactoryTransactor) CreateAccount(opts *bind.TransactOpts, _implementation common.Address, _initializer []byte, _salt *big.Int) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "createAccount", _implementation, _initializer, _salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x296601cd.
//
// Solidity: function createAccount(address _implementation, bytes _initializer, uint256 _salt) returns(address)
func (_AccountFactory *AccountFactorySession) CreateAccount(_implementation common.Address, _initializer []byte, _salt *big.Int) (*types.Transaction, error) {
	return _AccountFactory.Contract.CreateAccount(&_AccountFactory.TransactOpts, _implementation, _initializer, _salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x296601cd.
//
// Solidity: function createAccount(address _implementation, bytes _initializer, uint256 _salt) returns(address)
func (_AccountFactory *AccountFactoryTransactorSession) CreateAccount(_implementation common.Address, _initializer []byte, _salt *big.Int) (*types.Transaction, error) {
	return _AccountFactory.Contract.CreateAccount(&_AccountFactory.TransactOpts, _implementation, _initializer, _salt)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_AccountFactory *AccountFactoryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_AccountFactory *AccountFactorySession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _AccountFactory.Contract.Initialize(&_AccountFactory.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_AccountFactory *AccountFactoryTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _AccountFactory.Contract.Initialize(&_AccountFactory.TransactOpts, initialOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountFactory *AccountFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountFactory *AccountFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountFactory.Contract.RenounceOwnership(&_AccountFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountFactory *AccountFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountFactory.Contract.RenounceOwnership(&_AccountFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountFactory *AccountFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountFactory *AccountFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountFactory.Contract.TransferOwnership(&_AccountFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountFactory *AccountFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountFactory.Contract.TransferOwnership(&_AccountFactory.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AccountFactory *AccountFactoryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AccountFactory *AccountFactorySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AccountFactory.Contract.UpgradeToAndCall(&_AccountFactory.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AccountFactory *AccountFactoryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AccountFactory.Contract.UpgradeToAndCall(&_AccountFactory.TransactOpts, newImplementation, data)
}

// AccountFactoryAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the AccountFactory contract.
type AccountFactoryAccountCreatedIterator struct {
	Event *AccountFactoryAccountCreated // Event containing the contract specifics and raw log

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
func (it *AccountFactoryAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountFactoryAccountCreated)
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
		it.Event = new(AccountFactoryAccountCreated)
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
func (it *AccountFactoryAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountFactoryAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountFactoryAccountCreated represents a AccountCreated event raised by the AccountFactory contract.
type AccountFactoryAccountCreated struct {
	Account        common.Address
	Implementation common.Address
	Initializer    []byte
	Salt           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0xa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f7437.
//
// Solidity: event AccountCreated(address indexed account, address _implementation, bytes _initializer, uint256 _salt)
func (_AccountFactory *AccountFactoryFilterer) FilterAccountCreated(opts *bind.FilterOpts, account []common.Address) (*AccountFactoryAccountCreatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccountFactory.contract.FilterLogs(opts, "AccountCreated", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryAccountCreatedIterator{contract: _AccountFactory.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0xa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f7437.
//
// Solidity: event AccountCreated(address indexed account, address _implementation, bytes _initializer, uint256 _salt)
func (_AccountFactory *AccountFactoryFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *AccountFactoryAccountCreated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccountFactory.contract.WatchLogs(opts, "AccountCreated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountFactoryAccountCreated)
				if err := _AccountFactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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

// ParseAccountCreated is a log parse operation binding the contract event 0xa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f7437.
//
// Solidity: event AccountCreated(address indexed account, address _implementation, bytes _initializer, uint256 _salt)
func (_AccountFactory *AccountFactoryFilterer) ParseAccountCreated(log types.Log) (*AccountFactoryAccountCreated, error) {
	event := new(AccountFactoryAccountCreated)
	if err := _AccountFactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AccountFactory contract.
type AccountFactoryInitializedIterator struct {
	Event *AccountFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *AccountFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountFactoryInitialized)
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
		it.Event = new(AccountFactoryInitialized)
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
func (it *AccountFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountFactoryInitialized represents a Initialized event raised by the AccountFactory contract.
type AccountFactoryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AccountFactory *AccountFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*AccountFactoryInitializedIterator, error) {

	logs, sub, err := _AccountFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AccountFactoryInitializedIterator{contract: _AccountFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AccountFactory *AccountFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AccountFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _AccountFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountFactoryInitialized)
				if err := _AccountFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AccountFactory *AccountFactoryFilterer) ParseInitialized(log types.Log) (*AccountFactoryInitialized, error) {
	event := new(AccountFactoryInitialized)
	if err := _AccountFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AccountFactory contract.
type AccountFactoryOwnershipTransferredIterator struct {
	Event *AccountFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountFactoryOwnershipTransferred)
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
		it.Event = new(AccountFactoryOwnershipTransferred)
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
func (it *AccountFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the AccountFactory contract.
type AccountFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountFactory *AccountFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryOwnershipTransferredIterator{contract: _AccountFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountFactory *AccountFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountFactoryOwnershipTransferred)
				if err := _AccountFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AccountFactory *AccountFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*AccountFactoryOwnershipTransferred, error) {
	event := new(AccountFactoryOwnershipTransferred)
	if err := _AccountFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountFactoryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AccountFactory contract.
type AccountFactoryUpgradedIterator struct {
	Event *AccountFactoryUpgraded // Event containing the contract specifics and raw log

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
func (it *AccountFactoryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountFactoryUpgraded)
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
		it.Event = new(AccountFactoryUpgraded)
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
func (it *AccountFactoryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountFactoryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountFactoryUpgraded represents a Upgraded event raised by the AccountFactory contract.
type AccountFactoryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AccountFactory *AccountFactoryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AccountFactoryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AccountFactory.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryUpgradedIterator{contract: _AccountFactory.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AccountFactory *AccountFactoryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AccountFactoryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AccountFactory.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountFactoryUpgraded)
				if err := _AccountFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AccountFactory *AccountFactoryFilterer) ParseUpgraded(log types.Log) (*AccountFactoryUpgraded, error) {
	event := new(AccountFactoryUpgraded)
	if err := _AccountFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
