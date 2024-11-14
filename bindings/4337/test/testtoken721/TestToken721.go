// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testtoken721

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

// TestToken721MetaData contains all meta data concerning the TestToken721 contract.
var TestToken721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080346102ea576001600160401b0390604090808201838111828210176102d4578252600c81526020926b54657374546f6b656e37323160a01b848301528251838101818110838211176102d4578452600681526554535437323160d01b858201528251908282116102d45760008054926001958685811c951680156102ca575b898610146102b6578190601f95868111610268575b5089908683116001146102095784926101fe575b5050600019600383901b1c191690861b1781555b81519384116101ea5784548581811c911680156101e0575b888210146101cc57838111610189575b50869284116001146101285783949596509261011d575b5050600019600383901b1c191690821b1790555b51610d7990816102f08239f35b0151905038806100fc565b9190601f1984169685845280842093905b8882106101725750508385969710610159575b505050811b019055610110565b015160001960f88460031b161c1916905538808061014c565b808785968294968601518155019501930190610139565b8582528782208480870160051c8201928a88106101c3575b0160051c019086905b8281106101b85750506100e5565b8381550186906101aa565b925081926101a1565b634e487b7160e01b82526022600452602482fd5b90607f16906100d5565b634e487b7160e01b81526041600452602490fd5b0151905038806100a9565b8480528a85208994509190601f198416865b8d8282106102525750508411610239575b505050811b0181556100bd565b015160001960f88460031b161c1916905538808061022c565b8385015186558c9790950194938401930161021b565b9091508380528984208680850160051c8201928c86106102ad575b918a91869594930160051c01915b82811061029f575050610095565b8681558594508a9101610291565b92508192610283565b634e487b7160e01b83526022600452602483fd5b94607f1694610080565b634e487b7160e01b600052604160045260246000fd5b600080fdfe608060408181526004918236101561001657600080fd5b600092833560e01c91826301ffc9a7146108a45750816306fdde03146107d7578163081812fc1461079b578163095ea7b3146106be57816323b872dd146106a657816340c10f191461046e57816342842e0e1461043f5781636352211e1461040e57816370a08231146103b957816395d89b411461029c578163a22cb465146101fa578163b88d4fde14610161578163c87b56dd1461010f575063e985e9c5146100bf57600080fd5b3461010b578060031936011261010b5760ff816020936100dd61094f565b6100e561096a565b6001600160a01b0391821683526005875283832091168252855220549151911615158152f35b5080fd5b83833461010b57602036600319011261010b5761012f61015d9335610bc4565b5081815161013c816109b5565b52805191610149836109b5565b82525191829160208352602083019061090f565b0390f35b919050346101f65760803660031901126101f65761017d61094f565b61018561096a565b60443591856064359567ffffffffffffffff871161010b573660238801121561010b57860135956101c16101b888610a09565b965196876109e7565b868652366024888301011161010b57866101f39760246020930183890137860101526101ee838383610a25565b610bff565b80f35b8280fd5b919050346101f657806003193601126101f65761021561094f565b9060243591821515809303610298576001600160a01b03169283156102835750338452600560205280842083855260205280842060ff1981541660ff8416179055519081527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3160203392a380f35b836024925191630b61174360e31b8352820152fd5b8480fd5b8284346103b657806003193601126103b65781519181600192600154938460011c91600186169586156103ac575b6020968785108114610399578899509688969785829a529182600014610372575050600114610316575b50505061015d92916103079103856109e7565b5192828493845283019061090f565b9190869350600183527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf65b82841061035a575050508201018161030761015d6102f4565b8054848a018601528895508794909301928101610341565b60ff19168782015293151560051b86019093019350849250610307915061015d90506102f4565b634e487b7160e01b835260228a52602483fd5b92607f16926102ca565b80fd5b8284346103b65760203660031901126103b6576001600160a01b036103dc61094f565b169283156103f95750806020938392526003845220549051908152f35b91516322718ad960e21b815291820152602490fd5b8284346103b65760203660031901126103b6575061042e60209235610bc4565b90516001600160a01b039091168152f35b50503461010b576101f39061045336610980565b91925192610460846109b5565b8584526101ee838383610a25565b9050346101f657816003193601126101f65761048861094f565b6024928335815190610499826109b5565b8682526001600160a01b038481169490929085156106905782895260209360028552858a2054168387821515928361065d575b818d5260038852888d2080546001019055828d5260028852888d2080546001600160a01b031916831790557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8d80a4610647573b610528578780f35b90828888610568899a9b969798999589519586948594630a85bd0160e11b998a87523390870152850152604484015260806064840152608483019061090f565b0381868a5af1839181610603575b506105cb5750503d156105c3573d61058d81610a09565b9061059a855192836109e7565b81528091833d92013e5b805191826105c0575050505191633250574960e11b8352820152fd5b01fd5b5060606105a4565b919695949392506001600160e01b0319909116036105f25750505050388080808080808780f35b51633250574960e11b815291820152fd5b9091508481813d8311610640575b61061b81836109e7565b8101031261063c57516001600160e01b03198116810361063c579038610576565b8380fd5b503d610611565b84516339e3563760e11b81528088018a90528890fd5b600083815260046020526040902080546001600160a01b0319169055808d5260038852888d2080546000190190556104cc565b8451633250574960e11b81528088018a90528890fd5b83346103b6576101f36106b836610980565b91610a25565b919050346101f657806003193601126101f6576106d961094f565b916024356106e681610bc4565b33151580610788575b8061075f575b610749576001600160a01b039485169482918691167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258880a48452602052822080546001600160a01b031916909117905580f35b835163a9fbf51f60e01b81523381850152602490fd5b506001600160a01b03811686526005602090815284872033885290528386205460ff16156106f5565b506001600160a01b0381163314156106ef565b9050346101f65760203660031901126101f6579182602093356107bd81610bc4565b50825283528190205490516001600160a01b039091168152f35b8284346103b657806003193601126103b6578151918182549260018460011c916001861695861561089a575b6020968785108114610399578899509688969785829a52918260001461037257505060011461083f5750505061015d92916103079103856109e7565b91908693508280527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5635b828410610882575050508201018161030761015d6102f4565b8054848a018601528895508794909301928101610869565b92607f1692610803565b8491346101f65760203660031901126101f6573563ffffffff60e01b81168091036101f657602092506380ac58cd60e01b81149081156108fe575b81156108ed575b5015158152f35b6301ffc9a760e01b149050836108e6565b635b5e139f60e01b811491506108df565b919082519283825260005b84811061093b575050826000602080949584010152601f8019910116010190565b60208183018101518483018201520161091a565b600435906001600160a01b038216820361096557565b600080fd5b602435906001600160a01b038216820361096557565b6060906003190112610965576001600160a01b0390600435828116810361096557916024359081168103610965579060443590565b6020810190811067ffffffffffffffff8211176109d157604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff8211176109d157604052565b67ffffffffffffffff81116109d157601f01601f191660200190565b6001600160a01b039182169290918315610bab57600092828452826020956002875260409684888820541696879133151580610b15575b509060027fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9284610ae2575b858352600381528b8320805460010190558683525289812080546001600160a01b0319168517905580a41692838303610ac15750505050565b6064945051926364283d7b60e01b8452600484015260248301526044820152fd5b600087815260046020526040902080546001600160a01b0319169055848352600381528b83208054600019019055610a88565b91939450915080610b6a575b15610b3157859291879138610a5c565b878688610b4e576024915190637e27328960e01b82526004820152fd5b604491519063177e802f60e01b82523360048301526024820152fd5b503387148015610b8f575b80610b215750858252600481523385898420541614610b21565b5086825260058152878220338352815260ff8883205416610b75565b604051633250574960e11b815260006004820152602490fd5b6000818152600260205260409020546001600160a01b0316908115610be7575090565b60249060405190637e27328960e01b82526004820152fd5b813b610c0c575b50505050565b604051630a85bd0160e11b8082523360048301526001600160a01b03928316602483015260448201949094526080606482015260209592909116939092908390610c5a90608483019061090f565b039285816000958187895af1849181610d03575b50610cce575050503d600014610cc6573d610c8881610a09565b90610c9660405192836109e7565b81528091843d92013e5b80519283610cc157604051633250574960e11b815260048101849052602490fd5b019050fd5b506060610ca0565b919450915063ffffffff60e01b1603610ceb575038808080610c06565b60249060405190633250574960e11b82526004820152fd5b9091508681813d8311610d3c575b610d1b81836109e7565b8101031261029857516001600160e01b031981168103610298579038610c6e565b503d610d1156fea26469706673582212208e7033a894d86a407ec529843fb0997dd45f60352f659ca79fb2a779155f6dbe64736f6c63430008190033",
}

// TestToken721ABI is the input ABI used to generate the binding from.
// Deprecated: Use TestToken721MetaData.ABI instead.
var TestToken721ABI = TestToken721MetaData.ABI

// TestToken721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestToken721MetaData.Bin instead.
var TestToken721Bin = TestToken721MetaData.Bin

// DeployTestToken721 deploys a new Ethereum contract, binding an instance of TestToken721 to it.
func DeployTestToken721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestToken721, error) {
	parsed, err := TestToken721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestToken721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestToken721{TestToken721Caller: TestToken721Caller{contract: contract}, TestToken721Transactor: TestToken721Transactor{contract: contract}, TestToken721Filterer: TestToken721Filterer{contract: contract}}, nil
}

// TestToken721 is an auto generated Go binding around an Ethereum contract.
type TestToken721 struct {
	TestToken721Caller     // Read-only binding to the contract
	TestToken721Transactor // Write-only binding to the contract
	TestToken721Filterer   // Log filterer for contract events
}

// TestToken721Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestToken721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestToken721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestToken721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestToken721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestToken721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestToken721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestToken721Session struct {
	Contract     *TestToken721     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestToken721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestToken721CallerSession struct {
	Contract *TestToken721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TestToken721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestToken721TransactorSession struct {
	Contract     *TestToken721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TestToken721Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestToken721Raw struct {
	Contract *TestToken721 // Generic contract binding to access the raw methods on
}

// TestToken721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestToken721CallerRaw struct {
	Contract *TestToken721Caller // Generic read-only contract binding to access the raw methods on
}

// TestToken721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestToken721TransactorRaw struct {
	Contract *TestToken721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestToken721 creates a new instance of TestToken721, bound to a specific deployed contract.
func NewTestToken721(address common.Address, backend bind.ContractBackend) (*TestToken721, error) {
	contract, err := bindTestToken721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestToken721{TestToken721Caller: TestToken721Caller{contract: contract}, TestToken721Transactor: TestToken721Transactor{contract: contract}, TestToken721Filterer: TestToken721Filterer{contract: contract}}, nil
}

// NewTestToken721Caller creates a new read-only instance of TestToken721, bound to a specific deployed contract.
func NewTestToken721Caller(address common.Address, caller bind.ContractCaller) (*TestToken721Caller, error) {
	contract, err := bindTestToken721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestToken721Caller{contract: contract}, nil
}

// NewTestToken721Transactor creates a new write-only instance of TestToken721, bound to a specific deployed contract.
func NewTestToken721Transactor(address common.Address, transactor bind.ContractTransactor) (*TestToken721Transactor, error) {
	contract, err := bindTestToken721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestToken721Transactor{contract: contract}, nil
}

// NewTestToken721Filterer creates a new log filterer instance of TestToken721, bound to a specific deployed contract.
func NewTestToken721Filterer(address common.Address, filterer bind.ContractFilterer) (*TestToken721Filterer, error) {
	contract, err := bindTestToken721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestToken721Filterer{contract: contract}, nil
}

// bindTestToken721 binds a generic wrapper to an already deployed contract.
func bindTestToken721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestToken721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestToken721 *TestToken721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestToken721.Contract.TestToken721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestToken721 *TestToken721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestToken721.Contract.TestToken721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestToken721 *TestToken721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestToken721.Contract.TestToken721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestToken721 *TestToken721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestToken721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestToken721 *TestToken721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestToken721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestToken721 *TestToken721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestToken721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TestToken721 *TestToken721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestToken721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TestToken721 *TestToken721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TestToken721.Contract.BalanceOf(&_TestToken721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TestToken721 *TestToken721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TestToken721.Contract.BalanceOf(&_TestToken721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TestToken721 *TestToken721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TestToken721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TestToken721 *TestToken721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TestToken721.Contract.GetApproved(&_TestToken721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TestToken721 *TestToken721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TestToken721.Contract.GetApproved(&_TestToken721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TestToken721 *TestToken721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TestToken721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TestToken721 *TestToken721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TestToken721.Contract.IsApprovedForAll(&_TestToken721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TestToken721 *TestToken721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TestToken721.Contract.IsApprovedForAll(&_TestToken721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken721 *TestToken721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestToken721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken721 *TestToken721Session) Name() (string, error) {
	return _TestToken721.Contract.Name(&_TestToken721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken721 *TestToken721CallerSession) Name() (string, error) {
	return _TestToken721.Contract.Name(&_TestToken721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TestToken721 *TestToken721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TestToken721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TestToken721 *TestToken721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TestToken721.Contract.OwnerOf(&_TestToken721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TestToken721 *TestToken721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TestToken721.Contract.OwnerOf(&_TestToken721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestToken721 *TestToken721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TestToken721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestToken721 *TestToken721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestToken721.Contract.SupportsInterface(&_TestToken721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TestToken721 *TestToken721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestToken721.Contract.SupportsInterface(&_TestToken721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken721 *TestToken721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestToken721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken721 *TestToken721Session) Symbol() (string, error) {
	return _TestToken721.Contract.Symbol(&_TestToken721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken721 *TestToken721CallerSession) Symbol() (string, error) {
	return _TestToken721.Contract.Symbol(&_TestToken721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TestToken721 *TestToken721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _TestToken721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TestToken721 *TestToken721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _TestToken721.Contract.TokenURI(&_TestToken721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TestToken721 *TestToken721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _TestToken721.Contract.TokenURI(&_TestToken721.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TestToken721 *TestToken721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestToken721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TestToken721 *TestToken721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestToken721.Contract.Approve(&_TestToken721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TestToken721 *TestToken721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestToken721.Contract.Approve(&_TestToken721.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_TestToken721 *TestToken721Transactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestToken721.contract.Transact(opts, "mint", to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_TestToken721 *TestToken721Session) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestToken721.Contract.Mint(&_TestToken721.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_TestToken721 *TestToken721TransactorSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestToken721.Contract.Mint(&_TestToken721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TestToken721 *TestToken721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestToken721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TestToken721 *TestToken721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestToken721.Contract.SafeTransferFrom(&_TestToken721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TestToken721 *TestToken721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestToken721.Contract.SafeTransferFrom(&_TestToken721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TestToken721 *TestToken721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TestToken721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TestToken721 *TestToken721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TestToken721.Contract.SafeTransferFrom0(&_TestToken721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TestToken721 *TestToken721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TestToken721.Contract.SafeTransferFrom0(&_TestToken721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TestToken721 *TestToken721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TestToken721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TestToken721 *TestToken721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TestToken721.Contract.SetApprovalForAll(&_TestToken721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TestToken721 *TestToken721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TestToken721.Contract.SetApprovalForAll(&_TestToken721.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TestToken721 *TestToken721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestToken721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TestToken721 *TestToken721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestToken721.Contract.TransferFrom(&_TestToken721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TestToken721 *TestToken721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestToken721.Contract.TransferFrom(&_TestToken721.TransactOpts, from, to, tokenId)
}

// TestToken721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestToken721 contract.
type TestToken721ApprovalIterator struct {
	Event *TestToken721Approval // Event containing the contract specifics and raw log

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
func (it *TestToken721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestToken721Approval)
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
		it.Event = new(TestToken721Approval)
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
func (it *TestToken721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestToken721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestToken721Approval represents a Approval event raised by the TestToken721 contract.
type TestToken721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TestToken721 *TestToken721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TestToken721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TestToken721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TestToken721ApprovalIterator{contract: _TestToken721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TestToken721 *TestToken721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestToken721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TestToken721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestToken721Approval)
				if err := _TestToken721.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TestToken721 *TestToken721Filterer) ParseApproval(log types.Log) (*TestToken721Approval, error) {
	event := new(TestToken721Approval)
	if err := _TestToken721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestToken721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TestToken721 contract.
type TestToken721ApprovalForAllIterator struct {
	Event *TestToken721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TestToken721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestToken721ApprovalForAll)
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
		it.Event = new(TestToken721ApprovalForAll)
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
func (it *TestToken721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestToken721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestToken721ApprovalForAll represents a ApprovalForAll event raised by the TestToken721 contract.
type TestToken721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TestToken721 *TestToken721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TestToken721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TestToken721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TestToken721ApprovalForAllIterator{contract: _TestToken721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TestToken721 *TestToken721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TestToken721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TestToken721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestToken721ApprovalForAll)
				if err := _TestToken721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TestToken721 *TestToken721Filterer) ParseApprovalForAll(log types.Log) (*TestToken721ApprovalForAll, error) {
	event := new(TestToken721ApprovalForAll)
	if err := _TestToken721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestToken721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestToken721 contract.
type TestToken721TransferIterator struct {
	Event *TestToken721Transfer // Event containing the contract specifics and raw log

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
func (it *TestToken721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestToken721Transfer)
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
		it.Event = new(TestToken721Transfer)
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
func (it *TestToken721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestToken721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestToken721Transfer represents a Transfer event raised by the TestToken721 contract.
type TestToken721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TestToken721 *TestToken721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TestToken721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TestToken721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TestToken721TransferIterator{contract: _TestToken721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TestToken721 *TestToken721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestToken721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TestToken721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestToken721Transfer)
				if err := _TestToken721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TestToken721 *TestToken721Filterer) ParseTransfer(log types.Log) (*TestToken721Transfer, error) {
	event := new(TestToken721Transfer)
	if err := _TestToken721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
