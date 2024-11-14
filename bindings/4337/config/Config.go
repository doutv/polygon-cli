// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package config

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

// ConfigMetaData contains all meta data concerning the Config contract.
var ConfigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBundler\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"singer\",\"type\":\"address\"}],\"name\":\"RecoverySignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetSafeSingleton\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetSenderSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"WhitelistBundlerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"WhitelistBundlerRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_CALLBACK_HANDLER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"}],\"name\":\"addSafeSingleton\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"addWhitelistedBundlers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultCallbackHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverySigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"}],\"name\":\"removeSafeSingleton\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"removeWhitelistedBundlers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"safeSingleton\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"setRecoverySigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eoaSigner\",\"type\":\"address\"}],\"name\":\"setWalletSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"walletSigners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedBundler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0806040523460c857306080527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c1660b957506001600160401b036002600160401b0319828216016075575b604051610e1e90816100ce82396080518181816105c601526106a60152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a13880806056565b63f92ee8a960e01b8152600490fd5b600080fdfe6040608081526004908136101561001557600080fd5b600091823560e01c90816303087a5714610a7557816316b97e9a146109f2578163485cc955146108955781634f1ef2861461062a57816352d1902d146105b15781635eb43cb5146105735781637063908f146104ee578163715018a6146104845781637910c00b1461045c5781637a1c84491461042157816387ab75e8146103645781638da5cb5b1461032e5781639de78db2146102f0578163ad3cb1cc14610251578163be08382114610229578163ee3a5fab146101a657508063f2fde38b146101765763fcdc4727146100e957600080fd5b34610172576100f736610b59565b90610100610ceb565b8151835b81811061014257847f49bd286efd7e26d327db1b6f433560c5f6156363198289a65a6ca9904338f4f461013c86865191829182610c32565b0390a180f35b6001600160a01b036101548286610c08565b5116855260016020819052838620805460ff19168217905501610104565b5080fd5b82346101a35760203660031901126101a3576101a0610193610af0565b61019b610ceb565b610c77565b80f35b80fd5b91905034610225576020366003190112610225576101c2610af0565b6101ca610ceb565b6001600160a01b03169182156102175780546001600160a01b03191683179055519081527f75da6d4a6a549915e7b28de22af1418486fc73dc2359d7a6851a3ceca2982a6b90602090a180f35b905163e6c4247b60e01b8152fd5b8280fd5b9050346102255782600319360112610225575490516001600160a01b03909116815260209150f35b9050346102255782600319360112610225578151908282019082821067ffffffffffffffff8311176102dd5750825260058152602090640352e302e360dc1b6020820152825193849260208452825192836020860152825b8481106102c757505050828201840152601f01601f19168101030190f35b81810183015188820188015287955082016102a9565b634e487b7160e01b855260419052602484fd5b5050346101725760203660031901126101725760209160ff9082906001600160a01b0361031b610af0565b1681526001855220541690519015158152f35b505034610172578160031936011261017257600080516020610dc98339815191525490516001600160a01b039091168152602090f35b9190503461022557806003193601126102255761037f610af0565b6001600160a01b039081610391610b0b565b1693841561041357328652600160205260ff8487205416156104055750168084526002602090815282852080546001600160a01b03191685179055825191825281019290925242908201527f03b72cb2fffc620d2a9c4f6bce0636faa82c948488a45f2f4b0684bb17cfdc7290606090a180f35b835163f8a0c54b60e01b8152fd5b835163e6c4247b60e01b8152fd5b505034610172576020366003190112610172576020916001600160a01b039082908261044b610af0565b168152600285522054169051908152f35b505034610172578160031936011261017257905490516001600160a01b039091168152602090f35b83346101a357806003193601126101a35761049d610ceb565b600080516020610dc983398151915280546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b919050346102255760203660031901126102255761050a610af0565b610512610ceb565b6001600160a01b03169182156102175750816060917ff80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f804061919385526003602052808520600160ff198254161790558051918252600160208301524290820152a180f35b5050346101725760203660031901126101725760209160ff9082906001600160a01b0361059e610af0565b1681526003855220541690519015158152f35b8284346101a357806003193601126101a357507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316300361061d57602090517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b5163703e46dd60e11b8152fd5b9180915060031936011261022557610640610af0565b90602493843567ffffffffffffffff81116101725736602382011215610172578085013561066d81610bec565b9461067a85519687610b21565b81865260209182870193368a8383010111610891578186928b8693018737880101526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116308114908115610863575b50610853576106df610ceb565b81169585516352d1902d60e01b815283818a818b5afa869181610820575b50610719575050505050505191634c9c8ce360e01b8352820152fd5b9088888894938c7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc9182810361080b5750853b156107f7575080546001600160a01b031916821790558451889392917fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8580a28251156107d95750506107cb9582915190845af4913d156107cf573d6107bd6107b482610bec565b92519283610b21565b81528581943d92013e610d65565b5080f35b5060609250610d65565b9550955050505050346107eb57505080f35b63b398979f60e01b8152fd5b8651634c9c8ce360e01b8152808501849052fd5b8751632a87526960e21b815280860191909152fd5b9091508481813d831161084c575b6108388183610b21565b81010312610848575190386106fd565b8680fd5b503d61082e565b855163703e46dd60e11b81528890fd5b9050817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc54161415386106d2565b8580fd5b9050346102255781600319360112610225576108af610af0565b906108b8610b0b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0092835460ff81871c16159367ffffffffffffffff8216801590816109ea575b60011490816109e0575b1590816109d7575b506109c9575067ffffffffffffffff1981166001178555610959929190846109aa575b5086546001600160a01b0319166001600160a01b0391909116178655610951610d24565b61019b610d24565b610961610d24565b610969578280f35b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808280f35b68ffffffffffffffffff1916680100000000000000011785553861092d565b865163f92ee8a960e01b8152fd5b9050153861090a565b303b159150610902565b8691506108f8565b50503461017257610a0236610b59565b90610a0b610ceb565b8151835b818110610a4757847f623f7acab9936b83f17999b4d96bfda34b1d3a749bd474d39d0413cc977044ed61013c86865191829182610c32565b6001600160a01b03610a598286610c08565b5116855260016020819052838620805460ff1916905501610a0f565b5050346101725760203660031901126101725760607ff80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f8040619191610ab4610af0565b610abc610ceb565b6001600160a01b031680855260036020908152828620805460ff19169055825191825281018590524291810191909152a180f35b600435906001600160a01b0382168203610b0657565b600080fd5b602435906001600160a01b0382168203610b0657565b90601f8019910116810190811067ffffffffffffffff821117610b4357604052565b634e487b7160e01b600052604160045260246000fd5b602080600319830112610b065767ffffffffffffffff91600435838111610b065781602382011215610b06578060040135938411610b43578360051b9060405194610ba76020840187610b21565b855260246020860192820101928311610b0657602401905b828210610bcd575050505090565b81356001600160a01b0381168103610b06578152908301908301610bbf565b67ffffffffffffffff8111610b4357601f01601f191660200190565b8051821015610c1c5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b602090602060408183019282815285518094520193019160005b828110610c5a575050505090565b83516001600160a01b031685529381019392810192600101610c4c565b6001600160a01b03908116908115610cd257600080516020610dc983398151915280546001600160a01b031981168417909155167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b604051631e4fbdf760e01b815260006004820152602490fd5b600080516020610dc9833981519152546001600160a01b03163303610d0c57565b60405163118cdaa760e01b8152336004820152602490fd5b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c1615610d5357565b604051631afcd79f60e31b8152600490fd5b90610d8c5750805115610d7a57805190602001fd5b604051630a12f52160e11b8152600490fd5b81511580610dbf575b610d9d575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b15610d9556fe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300a2646970667358221220bff2ed44b38298e6351a5de4e1b24f06f99b0b867a45c534c89baa6d8e3fe30664736f6c63430008190033",
}

// ConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use ConfigMetaData.ABI instead.
var ConfigABI = ConfigMetaData.ABI

// ConfigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConfigMetaData.Bin instead.
var ConfigBin = ConfigMetaData.Bin

// DeployConfig deploys a new Ethereum contract, binding an instance of Config to it.
func DeployConfig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Config, error) {
	parsed, err := ConfigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConfigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Config{ConfigCaller: ConfigCaller{contract: contract}, ConfigTransactor: ConfigTransactor{contract: contract}, ConfigFilterer: ConfigFilterer{contract: contract}}, nil
}

// Config is an auto generated Go binding around an Ethereum contract.
type Config struct {
	ConfigCaller     // Read-only binding to the contract
	ConfigTransactor // Write-only binding to the contract
	ConfigFilterer   // Log filterer for contract events
}

// ConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConfigSession struct {
	Contract     *Config           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConfigCallerSession struct {
	Contract *ConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConfigTransactorSession struct {
	Contract     *ConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConfigRaw struct {
	Contract *Config // Generic contract binding to access the raw methods on
}

// ConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConfigCallerRaw struct {
	Contract *ConfigCaller // Generic read-only contract binding to access the raw methods on
}

// ConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConfigTransactorRaw struct {
	Contract *ConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConfig creates a new instance of Config, bound to a specific deployed contract.
func NewConfig(address common.Address, backend bind.ContractBackend) (*Config, error) {
	contract, err := bindConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Config{ConfigCaller: ConfigCaller{contract: contract}, ConfigTransactor: ConfigTransactor{contract: contract}, ConfigFilterer: ConfigFilterer{contract: contract}}, nil
}

// NewConfigCaller creates a new read-only instance of Config, bound to a specific deployed contract.
func NewConfigCaller(address common.Address, caller bind.ContractCaller) (*ConfigCaller, error) {
	contract, err := bindConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigCaller{contract: contract}, nil
}

// NewConfigTransactor creates a new write-only instance of Config, bound to a specific deployed contract.
func NewConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfigTransactor, error) {
	contract, err := bindConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigTransactor{contract: contract}, nil
}

// NewConfigFilterer creates a new log filterer instance of Config, bound to a specific deployed contract.
func NewConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfigFilterer, error) {
	contract, err := bindConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfigFilterer{contract: contract}, nil
}

// bindConfig binds a generic wrapper to an already deployed contract.
func bindConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Config *ConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Config.Contract.ConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Config *ConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.Contract.ConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Config *ConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Config.Contract.ConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Config *ConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Config.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Config *ConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Config *ConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Config.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTCALLBACKHANDLER is a free data retrieval call binding the contract method 0x7910c00b.
//
// Solidity: function DEFAULT_CALLBACK_HANDLER() view returns(address)
func (_Config *ConfigCaller) DEFAULTCALLBACKHANDLER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "DEFAULT_CALLBACK_HANDLER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEFAULTCALLBACKHANDLER is a free data retrieval call binding the contract method 0x7910c00b.
//
// Solidity: function DEFAULT_CALLBACK_HANDLER() view returns(address)
func (_Config *ConfigSession) DEFAULTCALLBACKHANDLER() (common.Address, error) {
	return _Config.Contract.DEFAULTCALLBACKHANDLER(&_Config.CallOpts)
}

// DEFAULTCALLBACKHANDLER is a free data retrieval call binding the contract method 0x7910c00b.
//
// Solidity: function DEFAULT_CALLBACK_HANDLER() view returns(address)
func (_Config *ConfigCallerSession) DEFAULTCALLBACKHANDLER() (common.Address, error) {
	return _Config.Contract.DEFAULTCALLBACKHANDLER(&_Config.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Config *ConfigCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Config *ConfigSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Config.Contract.UPGRADEINTERFACEVERSION(&_Config.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Config *ConfigCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Config.Contract.UPGRADEINTERFACEVERSION(&_Config.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Config *ConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Config *ConfigSession) Owner() (common.Address, error) {
	return _Config.Contract.Owner(&_Config.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Config *ConfigCallerSession) Owner() (common.Address, error) {
	return _Config.Contract.Owner(&_Config.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Config *ConfigCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Config *ConfigSession) ProxiableUUID() ([32]byte, error) {
	return _Config.Contract.ProxiableUUID(&_Config.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Config *ConfigCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Config.Contract.ProxiableUUID(&_Config.CallOpts)
}

// RecoverySigner is a free data retrieval call binding the contract method 0xbe083821.
//
// Solidity: function recoverySigner() view returns(address)
func (_Config *ConfigCaller) RecoverySigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "recoverySigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverySigner is a free data retrieval call binding the contract method 0xbe083821.
//
// Solidity: function recoverySigner() view returns(address)
func (_Config *ConfigSession) RecoverySigner() (common.Address, error) {
	return _Config.Contract.RecoverySigner(&_Config.CallOpts)
}

// RecoverySigner is a free data retrieval call binding the contract method 0xbe083821.
//
// Solidity: function recoverySigner() view returns(address)
func (_Config *ConfigCallerSession) RecoverySigner() (common.Address, error) {
	return _Config.Contract.RecoverySigner(&_Config.CallOpts)
}

// SafeSingleton is a free data retrieval call binding the contract method 0x5eb43cb5.
//
// Solidity: function safeSingleton(address ) view returns(bool)
func (_Config *ConfigCaller) SafeSingleton(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "safeSingleton", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SafeSingleton is a free data retrieval call binding the contract method 0x5eb43cb5.
//
// Solidity: function safeSingleton(address ) view returns(bool)
func (_Config *ConfigSession) SafeSingleton(arg0 common.Address) (bool, error) {
	return _Config.Contract.SafeSingleton(&_Config.CallOpts, arg0)
}

// SafeSingleton is a free data retrieval call binding the contract method 0x5eb43cb5.
//
// Solidity: function safeSingleton(address ) view returns(bool)
func (_Config *ConfigCallerSession) SafeSingleton(arg0 common.Address) (bool, error) {
	return _Config.Contract.SafeSingleton(&_Config.CallOpts, arg0)
}

// WalletSigners is a free data retrieval call binding the contract method 0x7a1c8449.
//
// Solidity: function walletSigners(address ) view returns(address)
func (_Config *ConfigCaller) WalletSigners(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "walletSigners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletSigners is a free data retrieval call binding the contract method 0x7a1c8449.
//
// Solidity: function walletSigners(address ) view returns(address)
func (_Config *ConfigSession) WalletSigners(arg0 common.Address) (common.Address, error) {
	return _Config.Contract.WalletSigners(&_Config.CallOpts, arg0)
}

// WalletSigners is a free data retrieval call binding the contract method 0x7a1c8449.
//
// Solidity: function walletSigners(address ) view returns(address)
func (_Config *ConfigCallerSession) WalletSigners(arg0 common.Address) (common.Address, error) {
	return _Config.Contract.WalletSigners(&_Config.CallOpts, arg0)
}

// WhitelistedBundler is a free data retrieval call binding the contract method 0x9de78db2.
//
// Solidity: function whitelistedBundler(address ) view returns(bool)
func (_Config *ConfigCaller) WhitelistedBundler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "whitelistedBundler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedBundler is a free data retrieval call binding the contract method 0x9de78db2.
//
// Solidity: function whitelistedBundler(address ) view returns(bool)
func (_Config *ConfigSession) WhitelistedBundler(arg0 common.Address) (bool, error) {
	return _Config.Contract.WhitelistedBundler(&_Config.CallOpts, arg0)
}

// WhitelistedBundler is a free data retrieval call binding the contract method 0x9de78db2.
//
// Solidity: function whitelistedBundler(address ) view returns(bool)
func (_Config *ConfigCallerSession) WhitelistedBundler(arg0 common.Address) (bool, error) {
	return _Config.Contract.WhitelistedBundler(&_Config.CallOpts, arg0)
}

// AddSafeSingleton is a paid mutator transaction binding the contract method 0x7063908f.
//
// Solidity: function addSafeSingleton(address singleton) returns()
func (_Config *ConfigTransactor) AddSafeSingleton(opts *bind.TransactOpts, singleton common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "addSafeSingleton", singleton)
}

// AddSafeSingleton is a paid mutator transaction binding the contract method 0x7063908f.
//
// Solidity: function addSafeSingleton(address singleton) returns()
func (_Config *ConfigSession) AddSafeSingleton(singleton common.Address) (*types.Transaction, error) {
	return _Config.Contract.AddSafeSingleton(&_Config.TransactOpts, singleton)
}

// AddSafeSingleton is a paid mutator transaction binding the contract method 0x7063908f.
//
// Solidity: function addSafeSingleton(address singleton) returns()
func (_Config *ConfigTransactorSession) AddSafeSingleton(singleton common.Address) (*types.Transaction, error) {
	return _Config.Contract.AddSafeSingleton(&_Config.TransactOpts, singleton)
}

// AddWhitelistedBundlers is a paid mutator transaction binding the contract method 0xfcdc4727.
//
// Solidity: function addWhitelistedBundlers(address[] bundlers) returns()
func (_Config *ConfigTransactor) AddWhitelistedBundlers(opts *bind.TransactOpts, bundlers []common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "addWhitelistedBundlers", bundlers)
}

// AddWhitelistedBundlers is a paid mutator transaction binding the contract method 0xfcdc4727.
//
// Solidity: function addWhitelistedBundlers(address[] bundlers) returns()
func (_Config *ConfigSession) AddWhitelistedBundlers(bundlers []common.Address) (*types.Transaction, error) {
	return _Config.Contract.AddWhitelistedBundlers(&_Config.TransactOpts, bundlers)
}

// AddWhitelistedBundlers is a paid mutator transaction binding the contract method 0xfcdc4727.
//
// Solidity: function addWhitelistedBundlers(address[] bundlers) returns()
func (_Config *ConfigTransactorSession) AddWhitelistedBundlers(bundlers []common.Address) (*types.Transaction, error) {
	return _Config.Contract.AddWhitelistedBundlers(&_Config.TransactOpts, bundlers)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address defaultCallbackHandler, address initialOwner) returns()
func (_Config *ConfigTransactor) Initialize(opts *bind.TransactOpts, defaultCallbackHandler common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "initialize", defaultCallbackHandler, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address defaultCallbackHandler, address initialOwner) returns()
func (_Config *ConfigSession) Initialize(defaultCallbackHandler common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Config.Contract.Initialize(&_Config.TransactOpts, defaultCallbackHandler, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address defaultCallbackHandler, address initialOwner) returns()
func (_Config *ConfigTransactorSession) Initialize(defaultCallbackHandler common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Config.Contract.Initialize(&_Config.TransactOpts, defaultCallbackHandler, initialOwner)
}

// RemoveSafeSingleton is a paid mutator transaction binding the contract method 0x03087a57.
//
// Solidity: function removeSafeSingleton(address singleton) returns()
func (_Config *ConfigTransactor) RemoveSafeSingleton(opts *bind.TransactOpts, singleton common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "removeSafeSingleton", singleton)
}

// RemoveSafeSingleton is a paid mutator transaction binding the contract method 0x03087a57.
//
// Solidity: function removeSafeSingleton(address singleton) returns()
func (_Config *ConfigSession) RemoveSafeSingleton(singleton common.Address) (*types.Transaction, error) {
	return _Config.Contract.RemoveSafeSingleton(&_Config.TransactOpts, singleton)
}

// RemoveSafeSingleton is a paid mutator transaction binding the contract method 0x03087a57.
//
// Solidity: function removeSafeSingleton(address singleton) returns()
func (_Config *ConfigTransactorSession) RemoveSafeSingleton(singleton common.Address) (*types.Transaction, error) {
	return _Config.Contract.RemoveSafeSingleton(&_Config.TransactOpts, singleton)
}

// RemoveWhitelistedBundlers is a paid mutator transaction binding the contract method 0x16b97e9a.
//
// Solidity: function removeWhitelistedBundlers(address[] bundlers) returns()
func (_Config *ConfigTransactor) RemoveWhitelistedBundlers(opts *bind.TransactOpts, bundlers []common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "removeWhitelistedBundlers", bundlers)
}

// RemoveWhitelistedBundlers is a paid mutator transaction binding the contract method 0x16b97e9a.
//
// Solidity: function removeWhitelistedBundlers(address[] bundlers) returns()
func (_Config *ConfigSession) RemoveWhitelistedBundlers(bundlers []common.Address) (*types.Transaction, error) {
	return _Config.Contract.RemoveWhitelistedBundlers(&_Config.TransactOpts, bundlers)
}

// RemoveWhitelistedBundlers is a paid mutator transaction binding the contract method 0x16b97e9a.
//
// Solidity: function removeWhitelistedBundlers(address[] bundlers) returns()
func (_Config *ConfigTransactorSession) RemoveWhitelistedBundlers(bundlers []common.Address) (*types.Transaction, error) {
	return _Config.Contract.RemoveWhitelistedBundlers(&_Config.TransactOpts, bundlers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Config *ConfigTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Config *ConfigSession) RenounceOwnership() (*types.Transaction, error) {
	return _Config.Contract.RenounceOwnership(&_Config.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Config *ConfigTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Config.Contract.RenounceOwnership(&_Config.TransactOpts)
}

// SetRecoverySigner is a paid mutator transaction binding the contract method 0xee3a5fab.
//
// Solidity: function setRecoverySigner(address signer) returns()
func (_Config *ConfigTransactor) SetRecoverySigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setRecoverySigner", signer)
}

// SetRecoverySigner is a paid mutator transaction binding the contract method 0xee3a5fab.
//
// Solidity: function setRecoverySigner(address signer) returns()
func (_Config *ConfigSession) SetRecoverySigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetRecoverySigner(&_Config.TransactOpts, signer)
}

// SetRecoverySigner is a paid mutator transaction binding the contract method 0xee3a5fab.
//
// Solidity: function setRecoverySigner(address signer) returns()
func (_Config *ConfigTransactorSession) SetRecoverySigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetRecoverySigner(&_Config.TransactOpts, signer)
}

// SetWalletSigner is a paid mutator transaction binding the contract method 0x87ab75e8.
//
// Solidity: function setWalletSigner(address sender, address eoaSigner) returns()
func (_Config *ConfigTransactor) SetWalletSigner(opts *bind.TransactOpts, sender common.Address, eoaSigner common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setWalletSigner", sender, eoaSigner)
}

// SetWalletSigner is a paid mutator transaction binding the contract method 0x87ab75e8.
//
// Solidity: function setWalletSigner(address sender, address eoaSigner) returns()
func (_Config *ConfigSession) SetWalletSigner(sender common.Address, eoaSigner common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetWalletSigner(&_Config.TransactOpts, sender, eoaSigner)
}

// SetWalletSigner is a paid mutator transaction binding the contract method 0x87ab75e8.
//
// Solidity: function setWalletSigner(address sender, address eoaSigner) returns()
func (_Config *ConfigTransactorSession) SetWalletSigner(sender common.Address, eoaSigner common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetWalletSigner(&_Config.TransactOpts, sender, eoaSigner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Config *ConfigTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Config *ConfigSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Config.Contract.TransferOwnership(&_Config.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Config *ConfigTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Config.Contract.TransferOwnership(&_Config.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Config *ConfigTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Config *ConfigSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Config.Contract.UpgradeToAndCall(&_Config.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Config *ConfigTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Config.Contract.UpgradeToAndCall(&_Config.TransactOpts, newImplementation, data)
}

// ConfigInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Config contract.
type ConfigInitializedIterator struct {
	Event *ConfigInitialized // Event containing the contract specifics and raw log

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
func (it *ConfigInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigInitialized)
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
		it.Event = new(ConfigInitialized)
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
func (it *ConfigInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigInitialized represents a Initialized event raised by the Config contract.
type ConfigInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Config *ConfigFilterer) FilterInitialized(opts *bind.FilterOpts) (*ConfigInitializedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ConfigInitializedIterator{contract: _Config.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Config *ConfigFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ConfigInitialized) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigInitialized)
				if err := _Config.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Config *ConfigFilterer) ParseInitialized(log types.Log) (*ConfigInitialized, error) {
	event := new(ConfigInitialized)
	if err := _Config.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Config contract.
type ConfigOwnershipTransferredIterator struct {
	Event *ConfigOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ConfigOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigOwnershipTransferred)
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
		it.Event = new(ConfigOwnershipTransferred)
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
func (it *ConfigOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigOwnershipTransferred represents a OwnershipTransferred event raised by the Config contract.
type ConfigOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Config *ConfigFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ConfigOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Config.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConfigOwnershipTransferredIterator{contract: _Config.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Config *ConfigFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConfigOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Config.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigOwnershipTransferred)
				if err := _Config.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Config *ConfigFilterer) ParseOwnershipTransferred(log types.Log) (*ConfigOwnershipTransferred, error) {
	event := new(ConfigOwnershipTransferred)
	if err := _Config.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigRecoverySignerUpdatedIterator is returned from FilterRecoverySignerUpdated and is used to iterate over the raw logs and unpacked data for RecoverySignerUpdated events raised by the Config contract.
type ConfigRecoverySignerUpdatedIterator struct {
	Event *ConfigRecoverySignerUpdated // Event containing the contract specifics and raw log

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
func (it *ConfigRecoverySignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigRecoverySignerUpdated)
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
		it.Event = new(ConfigRecoverySignerUpdated)
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
func (it *ConfigRecoverySignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigRecoverySignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigRecoverySignerUpdated represents a RecoverySignerUpdated event raised by the Config contract.
type ConfigRecoverySignerUpdated struct {
	Singer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverySignerUpdated is a free log retrieval operation binding the contract event 0x75da6d4a6a549915e7b28de22af1418486fc73dc2359d7a6851a3ceca2982a6b.
//
// Solidity: event RecoverySignerUpdated(address singer)
func (_Config *ConfigFilterer) FilterRecoverySignerUpdated(opts *bind.FilterOpts) (*ConfigRecoverySignerUpdatedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "RecoverySignerUpdated")
	if err != nil {
		return nil, err
	}
	return &ConfigRecoverySignerUpdatedIterator{contract: _Config.contract, event: "RecoverySignerUpdated", logs: logs, sub: sub}, nil
}

// WatchRecoverySignerUpdated is a free log subscription operation binding the contract event 0x75da6d4a6a549915e7b28de22af1418486fc73dc2359d7a6851a3ceca2982a6b.
//
// Solidity: event RecoverySignerUpdated(address singer)
func (_Config *ConfigFilterer) WatchRecoverySignerUpdated(opts *bind.WatchOpts, sink chan<- *ConfigRecoverySignerUpdated) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "RecoverySignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigRecoverySignerUpdated)
				if err := _Config.contract.UnpackLog(event, "RecoverySignerUpdated", log); err != nil {
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

// ParseRecoverySignerUpdated is a log parse operation binding the contract event 0x75da6d4a6a549915e7b28de22af1418486fc73dc2359d7a6851a3ceca2982a6b.
//
// Solidity: event RecoverySignerUpdated(address singer)
func (_Config *ConfigFilterer) ParseRecoverySignerUpdated(log types.Log) (*ConfigRecoverySignerUpdated, error) {
	event := new(ConfigRecoverySignerUpdated)
	if err := _Config.contract.UnpackLog(event, "RecoverySignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigSetSafeSingletonIterator is returned from FilterSetSafeSingleton and is used to iterate over the raw logs and unpacked data for SetSafeSingleton events raised by the Config contract.
type ConfigSetSafeSingletonIterator struct {
	Event *ConfigSetSafeSingleton // Event containing the contract specifics and raw log

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
func (it *ConfigSetSafeSingletonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigSetSafeSingleton)
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
		it.Event = new(ConfigSetSafeSingleton)
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
func (it *ConfigSetSafeSingletonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigSetSafeSingletonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigSetSafeSingleton represents a SetSafeSingleton event raised by the Config contract.
type ConfigSetSafeSingleton struct {
	Singleton common.Address
	Status    bool
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetSafeSingleton is a free log retrieval operation binding the contract event 0xf80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f80406191.
//
// Solidity: event SetSafeSingleton(address singleton, bool status, uint256 timestamp)
func (_Config *ConfigFilterer) FilterSetSafeSingleton(opts *bind.FilterOpts) (*ConfigSetSafeSingletonIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "SetSafeSingleton")
	if err != nil {
		return nil, err
	}
	return &ConfigSetSafeSingletonIterator{contract: _Config.contract, event: "SetSafeSingleton", logs: logs, sub: sub}, nil
}

// WatchSetSafeSingleton is a free log subscription operation binding the contract event 0xf80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f80406191.
//
// Solidity: event SetSafeSingleton(address singleton, bool status, uint256 timestamp)
func (_Config *ConfigFilterer) WatchSetSafeSingleton(opts *bind.WatchOpts, sink chan<- *ConfigSetSafeSingleton) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "SetSafeSingleton")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigSetSafeSingleton)
				if err := _Config.contract.UnpackLog(event, "SetSafeSingleton", log); err != nil {
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

// ParseSetSafeSingleton is a log parse operation binding the contract event 0xf80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f80406191.
//
// Solidity: event SetSafeSingleton(address singleton, bool status, uint256 timestamp)
func (_Config *ConfigFilterer) ParseSetSafeSingleton(log types.Log) (*ConfigSetSafeSingleton, error) {
	event := new(ConfigSetSafeSingleton)
	if err := _Config.contract.UnpackLog(event, "SetSafeSingleton", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigSetSenderSignerIterator is returned from FilterSetSenderSigner and is used to iterate over the raw logs and unpacked data for SetSenderSigner events raised by the Config contract.
type ConfigSetSenderSignerIterator struct {
	Event *ConfigSetSenderSigner // Event containing the contract specifics and raw log

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
func (it *ConfigSetSenderSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigSetSenderSigner)
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
		it.Event = new(ConfigSetSenderSigner)
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
func (it *ConfigSetSenderSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigSetSenderSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigSetSenderSigner represents a SetSenderSigner event raised by the Config contract.
type ConfigSetSenderSigner struct {
	Sender    common.Address
	Signer    common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetSenderSigner is a free log retrieval operation binding the contract event 0x03b72cb2fffc620d2a9c4f6bce0636faa82c948488a45f2f4b0684bb17cfdc72.
//
// Solidity: event SetSenderSigner(address sender, address signer, uint256 timestamp)
func (_Config *ConfigFilterer) FilterSetSenderSigner(opts *bind.FilterOpts) (*ConfigSetSenderSignerIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "SetSenderSigner")
	if err != nil {
		return nil, err
	}
	return &ConfigSetSenderSignerIterator{contract: _Config.contract, event: "SetSenderSigner", logs: logs, sub: sub}, nil
}

// WatchSetSenderSigner is a free log subscription operation binding the contract event 0x03b72cb2fffc620d2a9c4f6bce0636faa82c948488a45f2f4b0684bb17cfdc72.
//
// Solidity: event SetSenderSigner(address sender, address signer, uint256 timestamp)
func (_Config *ConfigFilterer) WatchSetSenderSigner(opts *bind.WatchOpts, sink chan<- *ConfigSetSenderSigner) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "SetSenderSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigSetSenderSigner)
				if err := _Config.contract.UnpackLog(event, "SetSenderSigner", log); err != nil {
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

// ParseSetSenderSigner is a log parse operation binding the contract event 0x03b72cb2fffc620d2a9c4f6bce0636faa82c948488a45f2f4b0684bb17cfdc72.
//
// Solidity: event SetSenderSigner(address sender, address signer, uint256 timestamp)
func (_Config *ConfigFilterer) ParseSetSenderSigner(log types.Log) (*ConfigSetSenderSigner, error) {
	event := new(ConfigSetSenderSigner)
	if err := _Config.contract.UnpackLog(event, "SetSenderSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Config contract.
type ConfigUpgradedIterator struct {
	Event *ConfigUpgraded // Event containing the contract specifics and raw log

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
func (it *ConfigUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigUpgraded)
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
		it.Event = new(ConfigUpgraded)
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
func (it *ConfigUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigUpgraded represents a Upgraded event raised by the Config contract.
type ConfigUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Config *ConfigFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ConfigUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Config.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ConfigUpgradedIterator{contract: _Config.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Config *ConfigFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ConfigUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Config.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigUpgraded)
				if err := _Config.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Config *ConfigFilterer) ParseUpgraded(log types.Log) (*ConfigUpgraded, error) {
	event := new(ConfigUpgraded)
	if err := _Config.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigWhitelistBundlerAddedIterator is returned from FilterWhitelistBundlerAdded and is used to iterate over the raw logs and unpacked data for WhitelistBundlerAdded events raised by the Config contract.
type ConfigWhitelistBundlerAddedIterator struct {
	Event *ConfigWhitelistBundlerAdded // Event containing the contract specifics and raw log

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
func (it *ConfigWhitelistBundlerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigWhitelistBundlerAdded)
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
		it.Event = new(ConfigWhitelistBundlerAdded)
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
func (it *ConfigWhitelistBundlerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigWhitelistBundlerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigWhitelistBundlerAdded represents a WhitelistBundlerAdded event raised by the Config contract.
type ConfigWhitelistBundlerAdded struct {
	Bundlers []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistBundlerAdded is a free log retrieval operation binding the contract event 0x49bd286efd7e26d327db1b6f433560c5f6156363198289a65a6ca9904338f4f4.
//
// Solidity: event WhitelistBundlerAdded(address[] bundlers)
func (_Config *ConfigFilterer) FilterWhitelistBundlerAdded(opts *bind.FilterOpts) (*ConfigWhitelistBundlerAddedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "WhitelistBundlerAdded")
	if err != nil {
		return nil, err
	}
	return &ConfigWhitelistBundlerAddedIterator{contract: _Config.contract, event: "WhitelistBundlerAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistBundlerAdded is a free log subscription operation binding the contract event 0x49bd286efd7e26d327db1b6f433560c5f6156363198289a65a6ca9904338f4f4.
//
// Solidity: event WhitelistBundlerAdded(address[] bundlers)
func (_Config *ConfigFilterer) WatchWhitelistBundlerAdded(opts *bind.WatchOpts, sink chan<- *ConfigWhitelistBundlerAdded) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "WhitelistBundlerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigWhitelistBundlerAdded)
				if err := _Config.contract.UnpackLog(event, "WhitelistBundlerAdded", log); err != nil {
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

// ParseWhitelistBundlerAdded is a log parse operation binding the contract event 0x49bd286efd7e26d327db1b6f433560c5f6156363198289a65a6ca9904338f4f4.
//
// Solidity: event WhitelistBundlerAdded(address[] bundlers)
func (_Config *ConfigFilterer) ParseWhitelistBundlerAdded(log types.Log) (*ConfigWhitelistBundlerAdded, error) {
	event := new(ConfigWhitelistBundlerAdded)
	if err := _Config.contract.UnpackLog(event, "WhitelistBundlerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigWhitelistBundlerRemovedIterator is returned from FilterWhitelistBundlerRemoved and is used to iterate over the raw logs and unpacked data for WhitelistBundlerRemoved events raised by the Config contract.
type ConfigWhitelistBundlerRemovedIterator struct {
	Event *ConfigWhitelistBundlerRemoved // Event containing the contract specifics and raw log

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
func (it *ConfigWhitelistBundlerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigWhitelistBundlerRemoved)
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
		it.Event = new(ConfigWhitelistBundlerRemoved)
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
func (it *ConfigWhitelistBundlerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigWhitelistBundlerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigWhitelistBundlerRemoved represents a WhitelistBundlerRemoved event raised by the Config contract.
type ConfigWhitelistBundlerRemoved struct {
	Bundlers []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistBundlerRemoved is a free log retrieval operation binding the contract event 0x623f7acab9936b83f17999b4d96bfda34b1d3a749bd474d39d0413cc977044ed.
//
// Solidity: event WhitelistBundlerRemoved(address[] bundlers)
func (_Config *ConfigFilterer) FilterWhitelistBundlerRemoved(opts *bind.FilterOpts) (*ConfigWhitelistBundlerRemovedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "WhitelistBundlerRemoved")
	if err != nil {
		return nil, err
	}
	return &ConfigWhitelistBundlerRemovedIterator{contract: _Config.contract, event: "WhitelistBundlerRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistBundlerRemoved is a free log subscription operation binding the contract event 0x623f7acab9936b83f17999b4d96bfda34b1d3a749bd474d39d0413cc977044ed.
//
// Solidity: event WhitelistBundlerRemoved(address[] bundlers)
func (_Config *ConfigFilterer) WatchWhitelistBundlerRemoved(opts *bind.WatchOpts, sink chan<- *ConfigWhitelistBundlerRemoved) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "WhitelistBundlerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigWhitelistBundlerRemoved)
				if err := _Config.contract.UnpackLog(event, "WhitelistBundlerRemoved", log); err != nil {
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

// ParseWhitelistBundlerRemoved is a log parse operation binding the contract event 0x623f7acab9936b83f17999b4d96bfda34b1d3a749bd474d39d0413cc977044ed.
//
// Solidity: event WhitelistBundlerRemoved(address[] bundlers)
func (_Config *ConfigFilterer) ParseWhitelistBundlerRemoved(log types.Log) (*ConfigWhitelistBundlerRemoved, error) {
	event := new(ConfigWhitelistBundlerRemoved)
	if err := _Config.contract.UnpackLog(event, "WhitelistBundlerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
