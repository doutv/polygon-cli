// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testwrappednativetoken

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

// TestWrappedNativeTokenMetaData contains all meta data concerning the TestWrappedNativeToken contract.
var TestWrappedNativeTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080346102fc576040906001600160401b039080830182811182821017610208578352601481526020917f57726170706564204e617469766520546f6b656e000000000000000000000000838301528351848101818110838211176102085785526005815264776e546f6b60d81b848201528251828111610208576003918254916001958684811c941680156102f2575b888510146102dc578190601f9485811161028b575b5088908583116001146102295760009261021e575b505060001982861b1c191690861b1783555b80519384116102085760049586548681811c911680156101fe575b828210146101e9578381116101a3575b508092851160011461013a575093839491849260009561012f575b50501b92600019911b1c19161790555b516108cf90816103028239f35b015193503880610112565b92919084601f1981168860005285600020956000905b89838310610189575050501061016f575b50505050811b019055610122565b01519060f884600019921b161c1916905538808080610161565b858701518955909701969485019488935090810190610150565b87600052816000208480880160051c8201928489106101e0575b0160051c019087905b8281106101d45750506100f7565b600081550187906101c6565b925081926101bd565b602288634e487b7160e01b6000525260246000fd5b90607f16906100e7565b634e487b7160e01b600052604160045260246000fd5b0151905038806100ba565b90889350601f19831691876000528a6000209260005b8c828210610275575050841161025d575b505050811b0183556100cc565b015160001983881b60f8161c19169055388080610250565b8385015186558c9790950194938401930161023f565b90915085600052886000208580850160051c8201928b86106102d3575b918a91869594930160051c01915b8281106102c45750506100a5565b600081558594508a91016102b6565b925081926102a8565b634e487b7160e01b600052602260045260246000fd5b93607f1693610090565b600080fdfe608060408181526004918236101561002a575b505050361561002057600080fd5b61002861074b565b005b600092833560e01c91826306fdde03146105a857508163095ea7b3146104fe57816318160ddd146104df57816323b872dd146103eb5781632e1a7d4d146102a0578163313ce5671461028457816370a082311461024d57816395d89b411461014857508063a9059cbb14610118578063d0e30db0146100fe5763dd62ed3e146100b35780610012565b346100fa57806003193601126100fa57806020926100cf6106e2565b6100d76106fd565b6001600160a01b0391821683526001865283832091168252845220549051908152f35b5080fd5b82806003193601126101155761011261074b565b80f35b80fd5b50346100fa57806003193601126100fa576020906101416101376106e2565b60243590336107d4565b5160018152f35b8383346100fa57816003193601126100fa5780519180938054916001908360011c9260018516948515610243575b60209586861081146102305785895290811561020c57506001146101b4575b6101b087876101a6828c0383610713565b5191829182610699565b0390f35b81529295507f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5b8284106101f957505050826101b0946101a692820101948680610195565b80548685018801529286019281016101db565b60ff19168887015250505050151560051b83010192506101a6826101b08680610195565b634e487b7160e01b845260228352602484fd5b93607f1693610176565b5050346100fa5760203660031901126100fa5760209181906001600160a01b036102756106e2565b16815280845220549051908152f35b5050346100fa57816003193601126100fa576020905160128152f35b919050346103e757602090816003193601126103e357823533156103cc57338552848352818520548181106103a05785808381948294338452838952038683205580600254036002558186518281527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef893392a3335af13d1561039b573d67ffffffffffffffff811161038857825190610343601f8201601f1916860183610713565b815285843d92013e5b15610355578380f35b5162461bcd60e51b815291820152600f60248201526e1d1c985b9cd9995c8819985a5b1959608a1b604482015260649150fd5b634e487b7160e01b865260418552602486fd5b61034c565b825163391434e360e21b8152338187019081526020810192909252604082018390529081906060010390fd5b8151634b637e8f60e11b8152808501869052602490fd5b8380fd5b8280fd5b90508234610115576060366003190112610115576104076106e2565b61040f6106fd565b916044359360018060a01b03831680835260016020528683203384526020528683205491600019830361044b575b6020886101418989896107d4565b8683106104b357811561049c573315610485575082526001602090815286832033845281529186902090859003905582906101418761043d565b8751634a1406b160e11b8152908101849052602490fd5b875163e602df0560e01b8152908101849052602490fd5b8751637dc7a0d960e11b8152339181019182526020820193909352604081018790528291506060010390fd5b5050346100fa57816003193601126100fa576020906002549051908152f35b9050346103e757816003193601126103e7576105186106e2565b602435903315610591576001600160a01b031691821561057a57508083602095338152600187528181208582528752205582519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925843392a35160018152f35b8351634a1406b160e11b8152908101859052602490fd5b835163e602df0560e01b8152808401869052602490fd5b929150346103e357836003193601126103e357600354600181811c918690828116801561068f575b602095868610821461067c575084885290811561065a5750600114610601575b6101b086866101a6828b0383610713565b929550600383527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b82841061064757505050826101b0946101a69282010194386105f0565b805486850188015292860192810161062a565b60ff191687860152505050151560051b83010192506101a6826101b0386105f0565b634e487b7160e01b845260229052602483fd5b93607f16936105d0565b6020808252825181830181905290939260005b8281106106ce57505060409293506000838284010152601f8019910116010190565b8181018601518482016040015285016106ac565b600435906001600160a01b03821682036106f857565b600080fd5b602435906001600160a01b03821682036106f857565b90601f8019910116810190811067ffffffffffffffff82111761073557604052565b634e487b7160e01b600052604160045260246000fd5b33156107bb576002543481018091116107a557600255336000526000602052604060002034815401905560405134815260007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60203393a3565b634e487b7160e01b600052601160045260246000fd5b60405163ec442f0560e01b815260006004820152602490fd5b916001600160a01b0380841692831561088057169283156107bb576000908382528160205260408220549083821061084e575091604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815220818154019055604051908152a3565b60405163391434e360e21b81526001600160a01b03919091166004820152602481019190915260448101839052606490fd5b604051634b637e8f60e11b815260006004820152602490fdfea2646970667358221220001e47d89a1d4d526cba723352581fdfc7d1f29901a0fa07c201a151afd7197b64736f6c63430008190033",
}

// TestWrappedNativeTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TestWrappedNativeTokenMetaData.ABI instead.
var TestWrappedNativeTokenABI = TestWrappedNativeTokenMetaData.ABI

// TestWrappedNativeTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestWrappedNativeTokenMetaData.Bin instead.
var TestWrappedNativeTokenBin = TestWrappedNativeTokenMetaData.Bin

// DeployTestWrappedNativeToken deploys a new Ethereum contract, binding an instance of TestWrappedNativeToken to it.
func DeployTestWrappedNativeToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestWrappedNativeToken, error) {
	parsed, err := TestWrappedNativeTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestWrappedNativeTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestWrappedNativeToken{TestWrappedNativeTokenCaller: TestWrappedNativeTokenCaller{contract: contract}, TestWrappedNativeTokenTransactor: TestWrappedNativeTokenTransactor{contract: contract}, TestWrappedNativeTokenFilterer: TestWrappedNativeTokenFilterer{contract: contract}}, nil
}

// TestWrappedNativeToken is an auto generated Go binding around an Ethereum contract.
type TestWrappedNativeToken struct {
	TestWrappedNativeTokenCaller     // Read-only binding to the contract
	TestWrappedNativeTokenTransactor // Write-only binding to the contract
	TestWrappedNativeTokenFilterer   // Log filterer for contract events
}

// TestWrappedNativeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestWrappedNativeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestWrappedNativeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestWrappedNativeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestWrappedNativeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestWrappedNativeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestWrappedNativeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestWrappedNativeTokenSession struct {
	Contract     *TestWrappedNativeToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TestWrappedNativeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestWrappedNativeTokenCallerSession struct {
	Contract *TestWrappedNativeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// TestWrappedNativeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestWrappedNativeTokenTransactorSession struct {
	Contract     *TestWrappedNativeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// TestWrappedNativeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestWrappedNativeTokenRaw struct {
	Contract *TestWrappedNativeToken // Generic contract binding to access the raw methods on
}

// TestWrappedNativeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestWrappedNativeTokenCallerRaw struct {
	Contract *TestWrappedNativeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TestWrappedNativeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestWrappedNativeTokenTransactorRaw struct {
	Contract *TestWrappedNativeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestWrappedNativeToken creates a new instance of TestWrappedNativeToken, bound to a specific deployed contract.
func NewTestWrappedNativeToken(address common.Address, backend bind.ContractBackend) (*TestWrappedNativeToken, error) {
	contract, err := bindTestWrappedNativeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestWrappedNativeToken{TestWrappedNativeTokenCaller: TestWrappedNativeTokenCaller{contract: contract}, TestWrappedNativeTokenTransactor: TestWrappedNativeTokenTransactor{contract: contract}, TestWrappedNativeTokenFilterer: TestWrappedNativeTokenFilterer{contract: contract}}, nil
}

// NewTestWrappedNativeTokenCaller creates a new read-only instance of TestWrappedNativeToken, bound to a specific deployed contract.
func NewTestWrappedNativeTokenCaller(address common.Address, caller bind.ContractCaller) (*TestWrappedNativeTokenCaller, error) {
	contract, err := bindTestWrappedNativeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestWrappedNativeTokenCaller{contract: contract}, nil
}

// NewTestWrappedNativeTokenTransactor creates a new write-only instance of TestWrappedNativeToken, bound to a specific deployed contract.
func NewTestWrappedNativeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TestWrappedNativeTokenTransactor, error) {
	contract, err := bindTestWrappedNativeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestWrappedNativeTokenTransactor{contract: contract}, nil
}

// NewTestWrappedNativeTokenFilterer creates a new log filterer instance of TestWrappedNativeToken, bound to a specific deployed contract.
func NewTestWrappedNativeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TestWrappedNativeTokenFilterer, error) {
	contract, err := bindTestWrappedNativeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestWrappedNativeTokenFilterer{contract: contract}, nil
}

// bindTestWrappedNativeToken binds a generic wrapper to an already deployed contract.
func bindTestWrappedNativeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestWrappedNativeTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestWrappedNativeToken *TestWrappedNativeTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestWrappedNativeToken.Contract.TestWrappedNativeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestWrappedNativeToken *TestWrappedNativeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.TestWrappedNativeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestWrappedNativeToken *TestWrappedNativeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.TestWrappedNativeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestWrappedNativeToken *TestWrappedNativeTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestWrappedNativeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestWrappedNativeToken *TestWrappedNativeTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestWrappedNativeToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestWrappedNativeToken *TestWrappedNativeTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestWrappedNativeToken.Contract.Allowance(&_TestWrappedNativeToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestWrappedNativeToken *TestWrappedNativeTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestWrappedNativeToken.Contract.Allowance(&_TestWrappedNativeToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestWrappedNativeToken *TestWrappedNativeTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestWrappedNativeToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestWrappedNativeToken *TestWrappedNativeTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestWrappedNativeToken.Contract.BalanceOf(&_TestWrappedNativeToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestWrappedNativeToken *TestWrappedNativeTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestWrappedNativeToken.Contract.BalanceOf(&_TestWrappedNativeToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestWrappedNativeToken *TestWrappedNativeTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestWrappedNativeToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestWrappedNativeToken *TestWrappedNativeTokenSession) Decimals() (uint8, error) {
	return _TestWrappedNativeToken.Contract.Decimals(&_TestWrappedNativeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestWrappedNativeToken *TestWrappedNativeTokenCallerSession) Decimals() (uint8, error) {
	return _TestWrappedNativeToken.Contract.Decimals(&_TestWrappedNativeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestWrappedNativeToken *TestWrappedNativeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestWrappedNativeToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestWrappedNativeToken *TestWrappedNativeTokenSession) Name() (string, error) {
	return _TestWrappedNativeToken.Contract.Name(&_TestWrappedNativeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestWrappedNativeToken *TestWrappedNativeTokenCallerSession) Name() (string, error) {
	return _TestWrappedNativeToken.Contract.Name(&_TestWrappedNativeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestWrappedNativeToken *TestWrappedNativeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestWrappedNativeToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestWrappedNativeToken *TestWrappedNativeTokenSession) Symbol() (string, error) {
	return _TestWrappedNativeToken.Contract.Symbol(&_TestWrappedNativeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestWrappedNativeToken *TestWrappedNativeTokenCallerSession) Symbol() (string, error) {
	return _TestWrappedNativeToken.Contract.Symbol(&_TestWrappedNativeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestWrappedNativeToken *TestWrappedNativeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestWrappedNativeToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestWrappedNativeToken *TestWrappedNativeTokenSession) TotalSupply() (*big.Int, error) {
	return _TestWrappedNativeToken.Contract.TotalSupply(&_TestWrappedNativeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestWrappedNativeToken *TestWrappedNativeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _TestWrappedNativeToken.Contract.TotalSupply(&_TestWrappedNativeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestWrappedNativeToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_TestWrappedNativeToken *TestWrappedNativeTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.Approve(&_TestWrappedNativeToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.Approve(&_TestWrappedNativeToken.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestWrappedNativeToken.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestWrappedNativeToken *TestWrappedNativeTokenSession) Deposit() (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.Deposit(&_TestWrappedNativeToken.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactorSession) Deposit() (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.Deposit(&_TestWrappedNativeToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestWrappedNativeToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_TestWrappedNativeToken *TestWrappedNativeTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.Transfer(&_TestWrappedNativeToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.Transfer(&_TestWrappedNativeToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestWrappedNativeToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_TestWrappedNativeToken *TestWrappedNativeTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.TransferFrom(&_TestWrappedNativeToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.TransferFrom(&_TestWrappedNativeToken.TransactOpts, from, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TestWrappedNativeToken.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_TestWrappedNativeToken *TestWrappedNativeTokenSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.Withdraw(&_TestWrappedNativeToken.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.Withdraw(&_TestWrappedNativeToken.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestWrappedNativeToken.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestWrappedNativeToken *TestWrappedNativeTokenSession) Receive() (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.Receive(&_TestWrappedNativeToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestWrappedNativeToken *TestWrappedNativeTokenTransactorSession) Receive() (*types.Transaction, error) {
	return _TestWrappedNativeToken.Contract.Receive(&_TestWrappedNativeToken.TransactOpts)
}

// TestWrappedNativeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestWrappedNativeToken contract.
type TestWrappedNativeTokenApprovalIterator struct {
	Event *TestWrappedNativeTokenApproval // Event containing the contract specifics and raw log

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
func (it *TestWrappedNativeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestWrappedNativeTokenApproval)
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
		it.Event = new(TestWrappedNativeTokenApproval)
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
func (it *TestWrappedNativeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestWrappedNativeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestWrappedNativeTokenApproval represents a Approval event raised by the TestWrappedNativeToken contract.
type TestWrappedNativeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestWrappedNativeToken *TestWrappedNativeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TestWrappedNativeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestWrappedNativeToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TestWrappedNativeTokenApprovalIterator{contract: _TestWrappedNativeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestWrappedNativeToken *TestWrappedNativeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestWrappedNativeTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestWrappedNativeToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestWrappedNativeTokenApproval)
				if err := _TestWrappedNativeToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestWrappedNativeToken *TestWrappedNativeTokenFilterer) ParseApproval(log types.Log) (*TestWrappedNativeTokenApproval, error) {
	event := new(TestWrappedNativeTokenApproval)
	if err := _TestWrappedNativeToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestWrappedNativeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestWrappedNativeToken contract.
type TestWrappedNativeTokenTransferIterator struct {
	Event *TestWrappedNativeTokenTransfer // Event containing the contract specifics and raw log

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
func (it *TestWrappedNativeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestWrappedNativeTokenTransfer)
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
		it.Event = new(TestWrappedNativeTokenTransfer)
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
func (it *TestWrappedNativeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestWrappedNativeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestWrappedNativeTokenTransfer represents a Transfer event raised by the TestWrappedNativeToken contract.
type TestWrappedNativeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestWrappedNativeToken *TestWrappedNativeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestWrappedNativeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestWrappedNativeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestWrappedNativeTokenTransferIterator{contract: _TestWrappedNativeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestWrappedNativeToken *TestWrappedNativeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestWrappedNativeTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestWrappedNativeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestWrappedNativeTokenTransfer)
				if err := _TestWrappedNativeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestWrappedNativeToken *TestWrappedNativeTokenFilterer) ParseTransfer(log types.Log) (*TestWrappedNativeTokenTransfer, error) {
	event := new(TestWrappedNativeTokenTransfer)
	if err := _TestWrappedNativeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
