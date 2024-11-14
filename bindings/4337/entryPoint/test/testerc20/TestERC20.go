// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testerc20

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

// TestERC20MetaData contains all meta data concerning the TestERC20 contract.
var TestERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sudoApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sudoMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"sudoTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a03461039f576001600160401b0390601f90601f19610be13881900384810183168401868111858210176102af5784928291604052833960209384918101031261039f57519360ff8516850361039f576100586103a4565b93600985526805465737445524332360bc1b848601526100766103a4565b946003918287526205432360ec1b8688015281518481116102af5783546001938482811c92168015610395575b8983101461037f57818484931161032e575b5088908483116001146102d0576000926102c5575b505060001982861b1c191690831b1783555b86519384116102af5760049687548381811c911680156102a5575b888210146102905782811161024a575b50869185116001146101e657849550908492916000956101db575b50501b92600019911b1c19161782555b33156101c35760025469d3c21bcecceda1000000928382018092116101ae57506000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91600255338352828152604083208481540190556040519384523393a360805260405161081d90816103c48239608051816102d80152f35b601190634e487b7160e01b6000525260246000fd5b60405163ec442f0560e01b8152600081840152602490fd5b015193503880610122565b9291948416928760005284876000209460005b89898383106102335750505010610219575b50505050811b018255610132565b01519060f884600019921b161c191690553880808061020b565b8686015189559097019694850194889350016101f9565b88600052876000208380880160051c8201928a8910610287575b0160051c019084905b82811061027b575050610107565b6000815501849061026d565b92508192610264565b602289634e487b7160e01b6000525260246000fd5b90607f16906100f7565b634e487b7160e01b600052604160045260246000fd5b0151905038806100ca565b908886941691876000528a6000209260005b8c8282106103185750508411610300575b505050811b0183556100dc565b015160001983881b60f8161c191690553880806102f3565b83850151865589979095019493840193016102e2565b90915085600052886000208480850160051c8201928b8610610376575b918791869594930160051c01915b8281106103675750506100b5565b60008155859450879101610359565b9250819261034b565b634e487b7160e01b600052602260045260246000fd5b91607f16916100a3565b600080fd5b60408051919082016001600160401b038111838210176102af5760405256fe6040608081526004908136101561001557600080fd5b600091823560e01c90816305ea5c22146105ac57816306fdde03146104d0578163095ea7b3146104a657816318160ddd1461048757816323b872dd146103a85781632d688ca8146102fc578163313ce567146102be57816370a082311461028757816395d89b411461016657508063a9059cbb14610136578063dd62ed3e146100ee5763fb4dcefa146100a757600080fd5b346100ea57806003193601126100ea576100e7906100c36105c7565b906100cc6105e2565b6001600160a01b038316855260208590529084205491610676565b80f35b5080fd5b50346100ea57806003193601126100ea578060209261010b6105c7565b6101136105e2565b6001600160a01b0391821683526001865283832091168252845220549051908152f35b50346100ea57806003193601126100ea5760209061015f6101556105c7565b6024359033610676565b5160018152f35b8383346100ea57816003193601126100ea5780519082845460018160011c906001831692831561027d575b602093848410811461026a5783885290811561024e57506001146101f9575b505050829003601f01601f191682019267ffffffffffffffff8411838510176101e657508291826101e292528261062d565b0390f35b634e487b7160e01b815260418552602490fd5b8787529192508591837f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5b83851061023a57505050508301018580806101b0565b805488860183015293019284908201610224565b60ff1916878501525050151560051b84010190508580806101b0565b634e487b7160e01b895260228a52602489fd5b91607f1691610191565b5050346100ea5760203660031901126100ea5760209181906001600160a01b036102af6105c7565b16815280845220549051908152f35b5050346100ea57816003193601126100ea576020905160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b919050346103a457806003193601126103a4576103176105c7565b6001600160a01b0316916024359190831561038f576002549083820180921161037c575084927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9260209260025585855284835280852082815401905551908152a380f35b634e487b7160e01b865260119052602485fd5b84602492519163ec442f0560e01b8352820152fd5b8280fd5b839150346100ea576103b9366105f8565b6001600160a01b03831680865260016020818152888820338952905287872054929693959083016103f3575b60208861015f898989610676565b86831061045b57811561044457331561042d5750825260016020908152868320338452815291869020908590039055829061015f876103e5565b8751634a1406b160e11b8152908101849052602490fd5b875163e602df0560e01b8152908101849052602490fd5b8751637dc7a0d960e11b8152339181019182526020820193909352604081018790528291506060010390fd5b5050346100ea57816003193601126100ea576020906002549051908152f35b5050346100ea57806003193601126100ea5760209061015f6104c66105c7565b6024359033610754565b8383346100ea57816003193601126100ea578051908260035460018160011c90600183169283156105a2575b602093848410811461026a5783885290811561024e575060011461054c57505050829003601f01601f191682019267ffffffffffffffff8411838510176101e657508291826101e292528261062d565b600387529192508591837fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b83851061058e57505050508301018580806101b0565b805488860183015293019284908201610578565b91607f16916104fc565b83346105c4576100e76105be366105f8565b91610754565b80fd5b600435906001600160a01b03821682036105dd57565b600080fd5b602435906001600160a01b03821682036105dd57565b60609060031901126105dd576001600160a01b039060043582811681036105dd579160243590811681036105dd579060443590565b6020808252825181830181905290939260005b82811061066257505060409293506000838284010152601f8019910116010190565b818101860151848201604001528501610640565b916001600160a01b0380841692831561073b571692831561072257600090838252816020526040822054908382106106f0575091604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815220818154019055604051908152a3565b60405163391434e360e21b81526001600160a01b03919091166004820152602481019190915260448101839052606490fd5b60405163ec442f0560e01b815260006004820152602490fd5b604051634b637e8f60e11b815260006004820152602490fd5b6001600160a01b039081169182156107ce57169182156107b55760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925918360005260018252604060002085600052825280604060002055604051908152a3565b604051634a1406b160e11b815260006004820152602490fd5b60405163e602df0560e01b815260006004820152602490fdfea2646970667358221220c499755dc0e13b9888aae599ab1c16a88dafd75211a8705d181f57952f571be864736f6c63430008190033",
}

// TestERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use TestERC20MetaData.ABI instead.
var TestERC20ABI = TestERC20MetaData.ABI

// TestERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestERC20MetaData.Bin instead.
var TestERC20Bin = TestERC20MetaData.Bin

// DeployTestERC20 deploys a new Ethereum contract, binding an instance of TestERC20 to it.
func DeployTestERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _decimals uint8) (common.Address, *types.Transaction, *TestERC20, error) {
	parsed, err := TestERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestERC20Bin), backend, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestERC20{TestERC20Caller: TestERC20Caller{contract: contract}, TestERC20Transactor: TestERC20Transactor{contract: contract}, TestERC20Filterer: TestERC20Filterer{contract: contract}}, nil
}

// TestERC20 is an auto generated Go binding around an Ethereum contract.
type TestERC20 struct {
	TestERC20Caller     // Read-only binding to the contract
	TestERC20Transactor // Write-only binding to the contract
	TestERC20Filterer   // Log filterer for contract events
}

// TestERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestERC20Session struct {
	Contract     *TestERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestERC20CallerSession struct {
	Contract *TestERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TestERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestERC20TransactorSession struct {
	Contract     *TestERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestERC20Raw struct {
	Contract *TestERC20 // Generic contract binding to access the raw methods on
}

// TestERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestERC20CallerRaw struct {
	Contract *TestERC20Caller // Generic read-only contract binding to access the raw methods on
}

// TestERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestERC20TransactorRaw struct {
	Contract *TestERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestERC20 creates a new instance of TestERC20, bound to a specific deployed contract.
func NewTestERC20(address common.Address, backend bind.ContractBackend) (*TestERC20, error) {
	contract, err := bindTestERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestERC20{TestERC20Caller: TestERC20Caller{contract: contract}, TestERC20Transactor: TestERC20Transactor{contract: contract}, TestERC20Filterer: TestERC20Filterer{contract: contract}}, nil
}

// NewTestERC20Caller creates a new read-only instance of TestERC20, bound to a specific deployed contract.
func NewTestERC20Caller(address common.Address, caller bind.ContractCaller) (*TestERC20Caller, error) {
	contract, err := bindTestERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestERC20Caller{contract: contract}, nil
}

// NewTestERC20Transactor creates a new write-only instance of TestERC20, bound to a specific deployed contract.
func NewTestERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*TestERC20Transactor, error) {
	contract, err := bindTestERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestERC20Transactor{contract: contract}, nil
}

// NewTestERC20Filterer creates a new log filterer instance of TestERC20, bound to a specific deployed contract.
func NewTestERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*TestERC20Filterer, error) {
	contract, err := bindTestERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestERC20Filterer{contract: contract}, nil
}

// bindTestERC20 binds a generic wrapper to an already deployed contract.
func bindTestERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestERC20 *TestERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestERC20.Contract.TestERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestERC20 *TestERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestERC20.Contract.TestERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestERC20 *TestERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestERC20.Contract.TestERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestERC20 *TestERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestERC20 *TestERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestERC20 *TestERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestERC20 *TestERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestERC20 *TestERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestERC20.Contract.Allowance(&_TestERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestERC20 *TestERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestERC20.Contract.Allowance(&_TestERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestERC20 *TestERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestERC20 *TestERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestERC20.Contract.BalanceOf(&_TestERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestERC20 *TestERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestERC20.Contract.BalanceOf(&_TestERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestERC20 *TestERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestERC20 *TestERC20Session) Decimals() (uint8, error) {
	return _TestERC20.Contract.Decimals(&_TestERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestERC20 *TestERC20CallerSession) Decimals() (uint8, error) {
	return _TestERC20.Contract.Decimals(&_TestERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestERC20 *TestERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestERC20 *TestERC20Session) Name() (string, error) {
	return _TestERC20.Contract.Name(&_TestERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestERC20 *TestERC20CallerSession) Name() (string, error) {
	return _TestERC20.Contract.Name(&_TestERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestERC20 *TestERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestERC20 *TestERC20Session) Symbol() (string, error) {
	return _TestERC20.Contract.Symbol(&_TestERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestERC20 *TestERC20CallerSession) Symbol() (string, error) {
	return _TestERC20.Contract.Symbol(&_TestERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestERC20 *TestERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestERC20 *TestERC20Session) TotalSupply() (*big.Int, error) {
	return _TestERC20.Contract.TotalSupply(&_TestERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestERC20 *TestERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _TestERC20.Contract.TotalSupply(&_TestERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_TestERC20 *TestERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_TestERC20 *TestERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestERC20.Contract.Approve(&_TestERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_TestERC20 *TestERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestERC20.Contract.Approve(&_TestERC20.TransactOpts, spender, value)
}

// SudoApprove is a paid mutator transaction binding the contract method 0x05ea5c22.
//
// Solidity: function sudoApprove(address _from, address _to, uint256 _amount) returns()
func (_TestERC20 *TestERC20Transactor) SudoApprove(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TestERC20.contract.Transact(opts, "sudoApprove", _from, _to, _amount)
}

// SudoApprove is a paid mutator transaction binding the contract method 0x05ea5c22.
//
// Solidity: function sudoApprove(address _from, address _to, uint256 _amount) returns()
func (_TestERC20 *TestERC20Session) SudoApprove(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TestERC20.Contract.SudoApprove(&_TestERC20.TransactOpts, _from, _to, _amount)
}

// SudoApprove is a paid mutator transaction binding the contract method 0x05ea5c22.
//
// Solidity: function sudoApprove(address _from, address _to, uint256 _amount) returns()
func (_TestERC20 *TestERC20TransactorSession) SudoApprove(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TestERC20.Contract.SudoApprove(&_TestERC20.TransactOpts, _from, _to, _amount)
}

// SudoMint is a paid mutator transaction binding the contract method 0x2d688ca8.
//
// Solidity: function sudoMint(address _to, uint256 _amount) returns()
func (_TestERC20 *TestERC20Transactor) SudoMint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TestERC20.contract.Transact(opts, "sudoMint", _to, _amount)
}

// SudoMint is a paid mutator transaction binding the contract method 0x2d688ca8.
//
// Solidity: function sudoMint(address _to, uint256 _amount) returns()
func (_TestERC20 *TestERC20Session) SudoMint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TestERC20.Contract.SudoMint(&_TestERC20.TransactOpts, _to, _amount)
}

// SudoMint is a paid mutator transaction binding the contract method 0x2d688ca8.
//
// Solidity: function sudoMint(address _to, uint256 _amount) returns()
func (_TestERC20 *TestERC20TransactorSession) SudoMint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TestERC20.Contract.SudoMint(&_TestERC20.TransactOpts, _to, _amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0xfb4dcefa.
//
// Solidity: function sudoTransfer(address _from, address _to) returns()
func (_TestERC20 *TestERC20Transactor) SudoTransfer(opts *bind.TransactOpts, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _TestERC20.contract.Transact(opts, "sudoTransfer", _from, _to)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0xfb4dcefa.
//
// Solidity: function sudoTransfer(address _from, address _to) returns()
func (_TestERC20 *TestERC20Session) SudoTransfer(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _TestERC20.Contract.SudoTransfer(&_TestERC20.TransactOpts, _from, _to)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0xfb4dcefa.
//
// Solidity: function sudoTransfer(address _from, address _to) returns()
func (_TestERC20 *TestERC20TransactorSession) SudoTransfer(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _TestERC20.Contract.SudoTransfer(&_TestERC20.TransactOpts, _from, _to)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_TestERC20 *TestERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_TestERC20 *TestERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestERC20.Contract.Transfer(&_TestERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_TestERC20 *TestERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestERC20.Contract.Transfer(&_TestERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_TestERC20 *TestERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_TestERC20 *TestERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestERC20.Contract.TransferFrom(&_TestERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_TestERC20 *TestERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestERC20.Contract.TransferFrom(&_TestERC20.TransactOpts, from, to, value)
}

// TestERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestERC20 contract.
type TestERC20ApprovalIterator struct {
	Event *TestERC20Approval // Event containing the contract specifics and raw log

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
func (it *TestERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestERC20Approval)
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
		it.Event = new(TestERC20Approval)
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
func (it *TestERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestERC20Approval represents a Approval event raised by the TestERC20 contract.
type TestERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestERC20 *TestERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TestERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TestERC20ApprovalIterator{contract: _TestERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestERC20 *TestERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestERC20Approval)
				if err := _TestERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_TestERC20 *TestERC20Filterer) ParseApproval(log types.Log) (*TestERC20Approval, error) {
	event := new(TestERC20Approval)
	if err := _TestERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestERC20 contract.
type TestERC20TransferIterator struct {
	Event *TestERC20Transfer // Event containing the contract specifics and raw log

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
func (it *TestERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestERC20Transfer)
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
		it.Event = new(TestERC20Transfer)
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
func (it *TestERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestERC20Transfer represents a Transfer event raised by the TestERC20 contract.
type TestERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestERC20 *TestERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestERC20TransferIterator{contract: _TestERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestERC20 *TestERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestERC20Transfer)
				if err := _TestERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_TestERC20 *TestERC20Filterer) ParseTransfer(log types.Log) (*TestERC20Transfer, error) {
	event := new(TestERC20Transfer)
	if err := _TestERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
