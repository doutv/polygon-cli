// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ellipticcurve

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

// EllipticCurveMetaData contains all meta data concerning the EllipticCurve contract.
var EllipticCurveMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes32\"},{\"name\":\"rs\",\"type\":\"uint256[2]\"},{\"name\":\"Q\",\"type\":\"uint256[2]\"}],\"name\":\"validateSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"}],\"name\":\"twice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"multipleGeneratorByScalar\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"isOnCurve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zeroProj\",\"outputs\":[{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"},{\"name\":\"exp\",\"type\":\"uint256\"}],\"name\":\"multiplyPowerBase2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"},{\"name\":\"z0\",\"type\":\"uint256\"}],\"name\":\"toAffinePoint\",\"outputs\":[{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"},{\"name\":\"x2\",\"type\":\"uint256\"},{\"name\":\"y2\",\"type\":\"uint256\"}],\"name\":\"addAndReturnProjectivePoint\",\"outputs\":[{\"name\":\"P\",\"type\":\"uint256[3]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"},{\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"multiplyScalar\",\"outputs\":[{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"}],\"name\":\"toProjectivePoint\",\"outputs\":[{\"name\":\"P\",\"type\":\"uint256[3]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zeroAffine\",\"outputs\":[{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"}],\"name\":\"isZeroCurve\",\"outputs\":[{\"name\":\"isZero\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"},{\"name\":\"z0\",\"type\":\"uint256\"}],\"name\":\"twiceProj\",\"outputs\":[{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"},{\"name\":\"z1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"},{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"},{\"name\":\"z0\",\"type\":\"uint256\"},{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"},{\"name\":\"z1\",\"type\":\"uint256\"}],\"name\":\"addProj\",\"outputs\":[{\"name\":\"x2\",\"type\":\"uint256\"},{\"name\":\"y2\",\"type\":\"uint256\"},{\"name\":\"z2\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611016806100206000396000f3fe6080604052600436106100d4577c0100000000000000000000000000000000000000000000000000000000600035046304e960d781146100d957806309d3ef31146101765780630afb4ddc146101bf5780630b0dbcfa146101e957806314c6706014610219578063322b24aa1461024c578063675ca04314610282578063713eca28146102b857806372fb4a141461032c5780637ec8da8d1461036257806384dfba4614610392578063c30cfa2d146103a7578063c80edca4146103d7578063e022d77c1461040d578063f214aba014610449575b600080fd5b3480156100e557600080fd5b50610162600480360360a08110156100fc57600080fd5b60408051808201825283359392830192916060830191906020840190600290839083908082843760009201919091525050604080518082018252929594938181019392509060029083908390808284376000920191909152509194506104919350505050565b604080519115158252519081900360200190f35b34801561018257600080fd5b506101a66004803603604081101561019957600080fd5b5080359060200135610630565b6040805192835260208301919091528051918290030190f35b3480156101cb57600080fd5b506101a6600480360360208110156101e257600080fd5b503561065f565b3480156101f557600080fd5b506101626004803603604081101561020c57600080fd5b50803590602001356106b6565b34801561022557600080fd5b5061022e6107af565b60408051938452602084019290925282820152519081900360600190f35b34801561025857600080fd5b506101a66004803603606081101561026f57600080fd5b50803590602081013590604001356107b9565b34801561028e57600080fd5b506101a6600480360360608110156102a557600080fd5b5080359060208101359060400135610800565b3480156102c457600080fd5b506102f4600480360360808110156102db57600080fd5b5080359060208101359060408101359060600135610850565b6040518082606080838360005b83811015610319578181015183820152602001610301565b5050505090500191505060405180910390f35b34801561033857600080fd5b506101a66004803603606081101561034f57600080fd5b5080359060208101359060400135610881565b34801561036e57600080fd5b506102f46004803603604081101561038557600080fd5b5080359060200135610948565b34801561039e57600080fd5b506101a66109a1565b3480156103b357600080fd5b50610162600480360360408110156103ca57600080fd5b50803590602001356109a8565b3480156103e357600080fd5b5061022e600480360360608110156103fa57600080fd5b50803590602081013590604001356109cc565b34801561041957600080fd5b506101a66004803603608081101561043057600080fd5b5080359060208101359060408101359060600135610c15565b34801561045557600080fd5b5061022e600480360360c081101561046c57600080fd5b5080359060208101359060408101359060608101359060808101359060a00135610c4a565b815160009015806104b157508251600080516020610fab83398151915211155b806104be57506020830151155b156104cb57506000610629565b815160208301516104dc91906106b6565b15156104ea57506000610629565b60008080808061050f8860016020020151600080516020610fab833981519152610d4f565b905061056d7f6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c2967f4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f5600080516020610fab833981519152848d09610881565b885160208a01518b5193985091955061059b92909190600080516020610fab83398151915290859009610881565b90945091506105a8610f8b565b6105b486858786610850565b604081015190915015156105d15760009650505050505050610629565b60006105f28260026020020151600080516020610fcb833981519152610d4f565b9050600080516020610fcb833981519152808283098351098a51600080516020610fab833981519152909106149750505050505050505b9392505050565b6000806000610641858560016109cc565b91965094509050610653858583610800565b92509250509250929050565b6000806106ad7f6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c2967f4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f585610881565b91509150915091565b60008215806106d25750600080516020610fcb83398151915283145b806106db575081155b806106f35750600080516020610fcb83398151915282145b15610700575060006107a9565b6000600080516020610fcb83398151915283840990506000600080516020610fcb83398151915285600080516020610fcb833981519152878809099050600080516020610fcb833981519152807bfffffffeffffffffffffffffffffffff00000000000000000000000319870982089050600080516020610fcb8339815191527f5ac635d8aa3a93e7b3ebbd55769886bc651d06b0cc53b0f63bce3c3e27d2604b820890501490505b92915050565b6000600181909192565b60008084846001835b868110156107e4576107d58484846109cc565b919550935091506001016107c2565b506107f0838383610800565b945094505050505b935093915050565b600080600061081d84600080516020610fcb833981519152610d4f565b9050600080516020610fcb8339815191528187099250600080516020610fcb833981519152818609915050935093915050565b610858610f8b565b60008061086787878787610c15565b90925090506108768282610948565b979650505050505050565b60008082151561089c576108936109a1565b915091506107f8565b82600114156108af5750839050826107f8565b82600214156108c2576108938585610630565b508390508281816001806002870615156108de57600094508495505b6002909604955b600087111561092d576108f98484846109cc565b9195509350915060028706600114156109225761091a848484898986610c4a565b919750955090505b6002909604956108e5565b610938868683610800565b9550955050505050935093915050565b610950610f8b565b600080516020610fcb833981519152600160000860408201819052600080516020610fcb83398151915290840981526040810151600080516020610fcb833981519152908309602082015292915050565b6000809091565b6000821580156109b6575081155b156109c3575060016107a9565b50600092915050565b60008060008060008060006109e18a8a6109a8565b156109fd576109ee6107af565b96509650965050505050610c0c565b600080516020610fcb833981519152888a099250600080516020610fcb833981519152600284099250600080516020610fcb8339815191528a84099150600080516020610fcb8339815191528983099150600080516020610fcb833981519152600283099150600080516020610fcb8339815191528a8b099950600080516020610fcb83398151915260038b099350600080516020610fcb8339815191528889099750600080516020610fcb8339815191527bfffffffeffffffffffffffffffffffff0000000000000000000000031989099750600080516020610fcb8339815191528885089350600080516020610fcb8339815191528485099050600080516020610fcb833981519152826002099950600080516020610fcb8339815191528a600080516020610fcb8339815191520382089050600080516020610fcb83398151915281600080516020610fcb8339815191520383089950600080516020610fcb8339815191528a85099950600080516020610fcb833981519152838a099850600080516020610fcb833981519152898a099850600080516020610fcb833981519152896002099850600080516020610fcb83398151915289600080516020610fcb833981519152038b089550600080516020610fcb8339815191528184099650600080516020610fcb8339815191528384099450600080516020610fcb8339815191528386099450505050505b93509350939050565b6000806000610c2a8787600188886001610c4a565b91985096509050610c3c878783610800565b925092505094509492505050565b6000806000806000806000610c5f8d8d6109a8565b15610c765789898996509650965050505050610d43565b610c808a8a6109a8565b15610c97578c8c8c96509650965050505050610d43565b600080516020610fcb833981519152888d099350600080516020610fcb8339815191528b8a099250600080516020610fcb833981519152888e099150600080516020610fcb8339815191528b8b09905080821415610d195782841415610d1157610d028d8d8d6109cc565b96509650965050505050610d43565b610d026107af565b610d37600080516020610fcb833981519152898d0983838688610de2565b91985096509450505050505b96509650969350505050565b6000821580610d5d57508183145b80610d66575081155b15610d73575060006107a9565b81831115610d8a578183811515610d8657fe5b0692505b600060018385835b8115610dbd578183811515610da357fe5b949594048581029094039391928383029003919050610d92565b6000851215610dd65750505050600003820390506107a9565b50929695505050505050565b600080808080808080600080516020610fcb8339815191528a600080516020610fcb833981519152038a089050600080516020610fcb8339815191528b600080516020610fcb833981519152038d089450600080516020610fcb8339815191528586099350600080516020610fcb8339815191528182099150600080516020610fcb8339815191528d83099150600080516020610fcb8339815191528c8c089a50600080516020610fcb833981519152848c099a50600080516020610fcb8339815191528b600080516020610fcb8339815191520383089150600080516020610fcb8339815191528286099750600080516020610fcb8339815191528585099250600080516020610fcb833981519152848d099b50600080516020610fcb83398151915282600080516020610fcb833981519152038d089b50600080516020610fcb8339815191528c82099050600080516020610fcb833981519152838a099850600080516020610fcb83398151915289600080516020610fcb8339815191520382089650600080516020610fcb8339815191528d840995505050505050955095509592505050565b606060405190810160405280600390602082028038833950919291505056feffffffff00000000ffffffffffffffffbce6faada7179e84f3b9cac2fc632551ffffffff00000001000000000000000000000000ffffffffffffffffffffffffa165627a7a7230582062ca6c092380f4e3919bcf8b9ea76337908c7baf41dd786c62e184545315e6420029",
}

// EllipticCurveABI is the input ABI used to generate the binding from.
// Deprecated: Use EllipticCurveMetaData.ABI instead.
var EllipticCurveABI = EllipticCurveMetaData.ABI

// EllipticCurveBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EllipticCurveMetaData.Bin instead.
var EllipticCurveBin = EllipticCurveMetaData.Bin

// DeployEllipticCurve deploys a new Ethereum contract, binding an instance of EllipticCurve to it.
func DeployEllipticCurve(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EllipticCurve, error) {
	parsed, err := EllipticCurveMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EllipticCurveBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EllipticCurve{EllipticCurveCaller: EllipticCurveCaller{contract: contract}, EllipticCurveTransactor: EllipticCurveTransactor{contract: contract}, EllipticCurveFilterer: EllipticCurveFilterer{contract: contract}}, nil
}

// EllipticCurve is an auto generated Go binding around an Ethereum contract.
type EllipticCurve struct {
	EllipticCurveCaller     // Read-only binding to the contract
	EllipticCurveTransactor // Write-only binding to the contract
	EllipticCurveFilterer   // Log filterer for contract events
}

// EllipticCurveCaller is an auto generated read-only Go binding around an Ethereum contract.
type EllipticCurveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EllipticCurveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EllipticCurveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EllipticCurveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EllipticCurveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EllipticCurveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EllipticCurveSession struct {
	Contract     *EllipticCurve    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EllipticCurveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EllipticCurveCallerSession struct {
	Contract *EllipticCurveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EllipticCurveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EllipticCurveTransactorSession struct {
	Contract     *EllipticCurveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EllipticCurveRaw is an auto generated low-level Go binding around an Ethereum contract.
type EllipticCurveRaw struct {
	Contract *EllipticCurve // Generic contract binding to access the raw methods on
}

// EllipticCurveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EllipticCurveCallerRaw struct {
	Contract *EllipticCurveCaller // Generic read-only contract binding to access the raw methods on
}

// EllipticCurveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EllipticCurveTransactorRaw struct {
	Contract *EllipticCurveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEllipticCurve creates a new instance of EllipticCurve, bound to a specific deployed contract.
func NewEllipticCurve(address common.Address, backend bind.ContractBackend) (*EllipticCurve, error) {
	contract, err := bindEllipticCurve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EllipticCurve{EllipticCurveCaller: EllipticCurveCaller{contract: contract}, EllipticCurveTransactor: EllipticCurveTransactor{contract: contract}, EllipticCurveFilterer: EllipticCurveFilterer{contract: contract}}, nil
}

// NewEllipticCurveCaller creates a new read-only instance of EllipticCurve, bound to a specific deployed contract.
func NewEllipticCurveCaller(address common.Address, caller bind.ContractCaller) (*EllipticCurveCaller, error) {
	contract, err := bindEllipticCurve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EllipticCurveCaller{contract: contract}, nil
}

// NewEllipticCurveTransactor creates a new write-only instance of EllipticCurve, bound to a specific deployed contract.
func NewEllipticCurveTransactor(address common.Address, transactor bind.ContractTransactor) (*EllipticCurveTransactor, error) {
	contract, err := bindEllipticCurve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EllipticCurveTransactor{contract: contract}, nil
}

// NewEllipticCurveFilterer creates a new log filterer instance of EllipticCurve, bound to a specific deployed contract.
func NewEllipticCurveFilterer(address common.Address, filterer bind.ContractFilterer) (*EllipticCurveFilterer, error) {
	contract, err := bindEllipticCurve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EllipticCurveFilterer{contract: contract}, nil
}

// bindEllipticCurve binds a generic wrapper to an already deployed contract.
func bindEllipticCurve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EllipticCurveMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EllipticCurve *EllipticCurveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EllipticCurve.Contract.EllipticCurveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EllipticCurve *EllipticCurveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EllipticCurve.Contract.EllipticCurveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EllipticCurve *EllipticCurveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EllipticCurve.Contract.EllipticCurveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EllipticCurve *EllipticCurveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EllipticCurve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EllipticCurve *EllipticCurveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EllipticCurve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EllipticCurve *EllipticCurveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EllipticCurve.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0xe022d77c.
//
// Solidity: function add(uint256 x0, uint256 y0, uint256 x1, uint256 y1) pure returns(uint256, uint256)
func (_EllipticCurve *EllipticCurveCaller) Add(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int, x1 *big.Int, y1 *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "add", x0, y0, x1, y1)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// Add is a free data retrieval call binding the contract method 0xe022d77c.
//
// Solidity: function add(uint256 x0, uint256 y0, uint256 x1, uint256 y1) pure returns(uint256, uint256)
func (_EllipticCurve *EllipticCurveSession) Add(x0 *big.Int, y0 *big.Int, x1 *big.Int, y1 *big.Int) (*big.Int, *big.Int, error) {
	return _EllipticCurve.Contract.Add(&_EllipticCurve.CallOpts, x0, y0, x1, y1)
}

// Add is a free data retrieval call binding the contract method 0xe022d77c.
//
// Solidity: function add(uint256 x0, uint256 y0, uint256 x1, uint256 y1) pure returns(uint256, uint256)
func (_EllipticCurve *EllipticCurveCallerSession) Add(x0 *big.Int, y0 *big.Int, x1 *big.Int, y1 *big.Int) (*big.Int, *big.Int, error) {
	return _EllipticCurve.Contract.Add(&_EllipticCurve.CallOpts, x0, y0, x1, y1)
}

// AddAndReturnProjectivePoint is a free data retrieval call binding the contract method 0x713eca28.
//
// Solidity: function addAndReturnProjectivePoint(uint256 x1, uint256 y1, uint256 x2, uint256 y2) pure returns(uint256[3] P)
func (_EllipticCurve *EllipticCurveCaller) AddAndReturnProjectivePoint(opts *bind.CallOpts, x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) ([3]*big.Int, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "addAndReturnProjectivePoint", x1, y1, x2, y2)

	if err != nil {
		return *new([3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)

	return out0, err

}

// AddAndReturnProjectivePoint is a free data retrieval call binding the contract method 0x713eca28.
//
// Solidity: function addAndReturnProjectivePoint(uint256 x1, uint256 y1, uint256 x2, uint256 y2) pure returns(uint256[3] P)
func (_EllipticCurve *EllipticCurveSession) AddAndReturnProjectivePoint(x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) ([3]*big.Int, error) {
	return _EllipticCurve.Contract.AddAndReturnProjectivePoint(&_EllipticCurve.CallOpts, x1, y1, x2, y2)
}

// AddAndReturnProjectivePoint is a free data retrieval call binding the contract method 0x713eca28.
//
// Solidity: function addAndReturnProjectivePoint(uint256 x1, uint256 y1, uint256 x2, uint256 y2) pure returns(uint256[3] P)
func (_EllipticCurve *EllipticCurveCallerSession) AddAndReturnProjectivePoint(x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) ([3]*big.Int, error) {
	return _EllipticCurve.Contract.AddAndReturnProjectivePoint(&_EllipticCurve.CallOpts, x1, y1, x2, y2)
}

// AddProj is a free data retrieval call binding the contract method 0xf214aba0.
//
// Solidity: function addProj(uint256 x0, uint256 y0, uint256 z0, uint256 x1, uint256 y1, uint256 z1) pure returns(uint256 x2, uint256 y2, uint256 z2)
func (_EllipticCurve *EllipticCurveCaller) AddProj(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int, z0 *big.Int, x1 *big.Int, y1 *big.Int, z1 *big.Int) (struct {
	X2 *big.Int
	Y2 *big.Int
	Z2 *big.Int
}, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "addProj", x0, y0, z0, x1, y1, z1)

	outstruct := new(struct {
		X2 *big.Int
		Y2 *big.Int
		Z2 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X2 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y2 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Z2 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AddProj is a free data retrieval call binding the contract method 0xf214aba0.
//
// Solidity: function addProj(uint256 x0, uint256 y0, uint256 z0, uint256 x1, uint256 y1, uint256 z1) pure returns(uint256 x2, uint256 y2, uint256 z2)
func (_EllipticCurve *EllipticCurveSession) AddProj(x0 *big.Int, y0 *big.Int, z0 *big.Int, x1 *big.Int, y1 *big.Int, z1 *big.Int) (struct {
	X2 *big.Int
	Y2 *big.Int
	Z2 *big.Int
}, error) {
	return _EllipticCurve.Contract.AddProj(&_EllipticCurve.CallOpts, x0, y0, z0, x1, y1, z1)
}

// AddProj is a free data retrieval call binding the contract method 0xf214aba0.
//
// Solidity: function addProj(uint256 x0, uint256 y0, uint256 z0, uint256 x1, uint256 y1, uint256 z1) pure returns(uint256 x2, uint256 y2, uint256 z2)
func (_EllipticCurve *EllipticCurveCallerSession) AddProj(x0 *big.Int, y0 *big.Int, z0 *big.Int, x1 *big.Int, y1 *big.Int, z1 *big.Int) (struct {
	X2 *big.Int
	Y2 *big.Int
	Z2 *big.Int
}, error) {
	return _EllipticCurve.Contract.AddProj(&_EllipticCurve.CallOpts, x0, y0, z0, x1, y1, z1)
}

// IsOnCurve is a free data retrieval call binding the contract method 0x0b0dbcfa.
//
// Solidity: function isOnCurve(uint256 x, uint256 y) pure returns(bool)
func (_EllipticCurve *EllipticCurveCaller) IsOnCurve(opts *bind.CallOpts, x *big.Int, y *big.Int) (bool, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "isOnCurve", x, y)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOnCurve is a free data retrieval call binding the contract method 0x0b0dbcfa.
//
// Solidity: function isOnCurve(uint256 x, uint256 y) pure returns(bool)
func (_EllipticCurve *EllipticCurveSession) IsOnCurve(x *big.Int, y *big.Int) (bool, error) {
	return _EllipticCurve.Contract.IsOnCurve(&_EllipticCurve.CallOpts, x, y)
}

// IsOnCurve is a free data retrieval call binding the contract method 0x0b0dbcfa.
//
// Solidity: function isOnCurve(uint256 x, uint256 y) pure returns(bool)
func (_EllipticCurve *EllipticCurveCallerSession) IsOnCurve(x *big.Int, y *big.Int) (bool, error) {
	return _EllipticCurve.Contract.IsOnCurve(&_EllipticCurve.CallOpts, x, y)
}

// IsZeroCurve is a free data retrieval call binding the contract method 0xc30cfa2d.
//
// Solidity: function isZeroCurve(uint256 x0, uint256 y0) pure returns(bool isZero)
func (_EllipticCurve *EllipticCurveCaller) IsZeroCurve(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int) (bool, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "isZeroCurve", x0, y0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsZeroCurve is a free data retrieval call binding the contract method 0xc30cfa2d.
//
// Solidity: function isZeroCurve(uint256 x0, uint256 y0) pure returns(bool isZero)
func (_EllipticCurve *EllipticCurveSession) IsZeroCurve(x0 *big.Int, y0 *big.Int) (bool, error) {
	return _EllipticCurve.Contract.IsZeroCurve(&_EllipticCurve.CallOpts, x0, y0)
}

// IsZeroCurve is a free data retrieval call binding the contract method 0xc30cfa2d.
//
// Solidity: function isZeroCurve(uint256 x0, uint256 y0) pure returns(bool isZero)
func (_EllipticCurve *EllipticCurveCallerSession) IsZeroCurve(x0 *big.Int, y0 *big.Int) (bool, error) {
	return _EllipticCurve.Contract.IsZeroCurve(&_EllipticCurve.CallOpts, x0, y0)
}

// MultipleGeneratorByScalar is a free data retrieval call binding the contract method 0x0afb4ddc.
//
// Solidity: function multipleGeneratorByScalar(uint256 scalar) pure returns(uint256, uint256)
func (_EllipticCurve *EllipticCurveCaller) MultipleGeneratorByScalar(opts *bind.CallOpts, scalar *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "multipleGeneratorByScalar", scalar)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// MultipleGeneratorByScalar is a free data retrieval call binding the contract method 0x0afb4ddc.
//
// Solidity: function multipleGeneratorByScalar(uint256 scalar) pure returns(uint256, uint256)
func (_EllipticCurve *EllipticCurveSession) MultipleGeneratorByScalar(scalar *big.Int) (*big.Int, *big.Int, error) {
	return _EllipticCurve.Contract.MultipleGeneratorByScalar(&_EllipticCurve.CallOpts, scalar)
}

// MultipleGeneratorByScalar is a free data retrieval call binding the contract method 0x0afb4ddc.
//
// Solidity: function multipleGeneratorByScalar(uint256 scalar) pure returns(uint256, uint256)
func (_EllipticCurve *EllipticCurveCallerSession) MultipleGeneratorByScalar(scalar *big.Int) (*big.Int, *big.Int, error) {
	return _EllipticCurve.Contract.MultipleGeneratorByScalar(&_EllipticCurve.CallOpts, scalar)
}

// MultiplyPowerBase2 is a free data retrieval call binding the contract method 0x322b24aa.
//
// Solidity: function multiplyPowerBase2(uint256 x0, uint256 y0, uint256 exp) pure returns(uint256, uint256)
func (_EllipticCurve *EllipticCurveCaller) MultiplyPowerBase2(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int, exp *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "multiplyPowerBase2", x0, y0, exp)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// MultiplyPowerBase2 is a free data retrieval call binding the contract method 0x322b24aa.
//
// Solidity: function multiplyPowerBase2(uint256 x0, uint256 y0, uint256 exp) pure returns(uint256, uint256)
func (_EllipticCurve *EllipticCurveSession) MultiplyPowerBase2(x0 *big.Int, y0 *big.Int, exp *big.Int) (*big.Int, *big.Int, error) {
	return _EllipticCurve.Contract.MultiplyPowerBase2(&_EllipticCurve.CallOpts, x0, y0, exp)
}

// MultiplyPowerBase2 is a free data retrieval call binding the contract method 0x322b24aa.
//
// Solidity: function multiplyPowerBase2(uint256 x0, uint256 y0, uint256 exp) pure returns(uint256, uint256)
func (_EllipticCurve *EllipticCurveCallerSession) MultiplyPowerBase2(x0 *big.Int, y0 *big.Int, exp *big.Int) (*big.Int, *big.Int, error) {
	return _EllipticCurve.Contract.MultiplyPowerBase2(&_EllipticCurve.CallOpts, x0, y0, exp)
}

// MultiplyScalar is a free data retrieval call binding the contract method 0x72fb4a14.
//
// Solidity: function multiplyScalar(uint256 x0, uint256 y0, uint256 scalar) pure returns(uint256 x1, uint256 y1)
func (_EllipticCurve *EllipticCurveCaller) MultiplyScalar(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int, scalar *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
}, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "multiplyScalar", x0, y0, scalar)

	outstruct := new(struct {
		X1 *big.Int
		Y1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X1 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MultiplyScalar is a free data retrieval call binding the contract method 0x72fb4a14.
//
// Solidity: function multiplyScalar(uint256 x0, uint256 y0, uint256 scalar) pure returns(uint256 x1, uint256 y1)
func (_EllipticCurve *EllipticCurveSession) MultiplyScalar(x0 *big.Int, y0 *big.Int, scalar *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
}, error) {
	return _EllipticCurve.Contract.MultiplyScalar(&_EllipticCurve.CallOpts, x0, y0, scalar)
}

// MultiplyScalar is a free data retrieval call binding the contract method 0x72fb4a14.
//
// Solidity: function multiplyScalar(uint256 x0, uint256 y0, uint256 scalar) pure returns(uint256 x1, uint256 y1)
func (_EllipticCurve *EllipticCurveCallerSession) MultiplyScalar(x0 *big.Int, y0 *big.Int, scalar *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
}, error) {
	return _EllipticCurve.Contract.MultiplyScalar(&_EllipticCurve.CallOpts, x0, y0, scalar)
}

// ToAffinePoint is a free data retrieval call binding the contract method 0x675ca043.
//
// Solidity: function toAffinePoint(uint256 x0, uint256 y0, uint256 z0) pure returns(uint256 x1, uint256 y1)
func (_EllipticCurve *EllipticCurveCaller) ToAffinePoint(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int, z0 *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
}, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "toAffinePoint", x0, y0, z0)

	outstruct := new(struct {
		X1 *big.Int
		Y1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X1 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ToAffinePoint is a free data retrieval call binding the contract method 0x675ca043.
//
// Solidity: function toAffinePoint(uint256 x0, uint256 y0, uint256 z0) pure returns(uint256 x1, uint256 y1)
func (_EllipticCurve *EllipticCurveSession) ToAffinePoint(x0 *big.Int, y0 *big.Int, z0 *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
}, error) {
	return _EllipticCurve.Contract.ToAffinePoint(&_EllipticCurve.CallOpts, x0, y0, z0)
}

// ToAffinePoint is a free data retrieval call binding the contract method 0x675ca043.
//
// Solidity: function toAffinePoint(uint256 x0, uint256 y0, uint256 z0) pure returns(uint256 x1, uint256 y1)
func (_EllipticCurve *EllipticCurveCallerSession) ToAffinePoint(x0 *big.Int, y0 *big.Int, z0 *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
}, error) {
	return _EllipticCurve.Contract.ToAffinePoint(&_EllipticCurve.CallOpts, x0, y0, z0)
}

// ToProjectivePoint is a free data retrieval call binding the contract method 0x7ec8da8d.
//
// Solidity: function toProjectivePoint(uint256 x0, uint256 y0) pure returns(uint256[3] P)
func (_EllipticCurve *EllipticCurveCaller) ToProjectivePoint(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int) ([3]*big.Int, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "toProjectivePoint", x0, y0)

	if err != nil {
		return *new([3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)

	return out0, err

}

// ToProjectivePoint is a free data retrieval call binding the contract method 0x7ec8da8d.
//
// Solidity: function toProjectivePoint(uint256 x0, uint256 y0) pure returns(uint256[3] P)
func (_EllipticCurve *EllipticCurveSession) ToProjectivePoint(x0 *big.Int, y0 *big.Int) ([3]*big.Int, error) {
	return _EllipticCurve.Contract.ToProjectivePoint(&_EllipticCurve.CallOpts, x0, y0)
}

// ToProjectivePoint is a free data retrieval call binding the contract method 0x7ec8da8d.
//
// Solidity: function toProjectivePoint(uint256 x0, uint256 y0) pure returns(uint256[3] P)
func (_EllipticCurve *EllipticCurveCallerSession) ToProjectivePoint(x0 *big.Int, y0 *big.Int) ([3]*big.Int, error) {
	return _EllipticCurve.Contract.ToProjectivePoint(&_EllipticCurve.CallOpts, x0, y0)
}

// Twice is a free data retrieval call binding the contract method 0x09d3ef31.
//
// Solidity: function twice(uint256 x0, uint256 y0) pure returns(uint256, uint256)
func (_EllipticCurve *EllipticCurveCaller) Twice(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "twice", x0, y0)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// Twice is a free data retrieval call binding the contract method 0x09d3ef31.
//
// Solidity: function twice(uint256 x0, uint256 y0) pure returns(uint256, uint256)
func (_EllipticCurve *EllipticCurveSession) Twice(x0 *big.Int, y0 *big.Int) (*big.Int, *big.Int, error) {
	return _EllipticCurve.Contract.Twice(&_EllipticCurve.CallOpts, x0, y0)
}

// Twice is a free data retrieval call binding the contract method 0x09d3ef31.
//
// Solidity: function twice(uint256 x0, uint256 y0) pure returns(uint256, uint256)
func (_EllipticCurve *EllipticCurveCallerSession) Twice(x0 *big.Int, y0 *big.Int) (*big.Int, *big.Int, error) {
	return _EllipticCurve.Contract.Twice(&_EllipticCurve.CallOpts, x0, y0)
}

// TwiceProj is a free data retrieval call binding the contract method 0xc80edca4.
//
// Solidity: function twiceProj(uint256 x0, uint256 y0, uint256 z0) pure returns(uint256 x1, uint256 y1, uint256 z1)
func (_EllipticCurve *EllipticCurveCaller) TwiceProj(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int, z0 *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
	Z1 *big.Int
}, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "twiceProj", x0, y0, z0)

	outstruct := new(struct {
		X1 *big.Int
		Y1 *big.Int
		Z1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X1 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Z1 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TwiceProj is a free data retrieval call binding the contract method 0xc80edca4.
//
// Solidity: function twiceProj(uint256 x0, uint256 y0, uint256 z0) pure returns(uint256 x1, uint256 y1, uint256 z1)
func (_EllipticCurve *EllipticCurveSession) TwiceProj(x0 *big.Int, y0 *big.Int, z0 *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
	Z1 *big.Int
}, error) {
	return _EllipticCurve.Contract.TwiceProj(&_EllipticCurve.CallOpts, x0, y0, z0)
}

// TwiceProj is a free data retrieval call binding the contract method 0xc80edca4.
//
// Solidity: function twiceProj(uint256 x0, uint256 y0, uint256 z0) pure returns(uint256 x1, uint256 y1, uint256 z1)
func (_EllipticCurve *EllipticCurveCallerSession) TwiceProj(x0 *big.Int, y0 *big.Int, z0 *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
	Z1 *big.Int
}, error) {
	return _EllipticCurve.Contract.TwiceProj(&_EllipticCurve.CallOpts, x0, y0, z0)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x04e960d7.
//
// Solidity: function validateSignature(bytes32 message, uint256[2] rs, uint256[2] Q) pure returns(bool)
func (_EllipticCurve *EllipticCurveCaller) ValidateSignature(opts *bind.CallOpts, message [32]byte, rs [2]*big.Int, Q [2]*big.Int) (bool, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "validateSignature", message, rs, Q)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateSignature is a free data retrieval call binding the contract method 0x04e960d7.
//
// Solidity: function validateSignature(bytes32 message, uint256[2] rs, uint256[2] Q) pure returns(bool)
func (_EllipticCurve *EllipticCurveSession) ValidateSignature(message [32]byte, rs [2]*big.Int, Q [2]*big.Int) (bool, error) {
	return _EllipticCurve.Contract.ValidateSignature(&_EllipticCurve.CallOpts, message, rs, Q)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x04e960d7.
//
// Solidity: function validateSignature(bytes32 message, uint256[2] rs, uint256[2] Q) pure returns(bool)
func (_EllipticCurve *EllipticCurveCallerSession) ValidateSignature(message [32]byte, rs [2]*big.Int, Q [2]*big.Int) (bool, error) {
	return _EllipticCurve.Contract.ValidateSignature(&_EllipticCurve.CallOpts, message, rs, Q)
}

// ZeroAffine is a free data retrieval call binding the contract method 0x84dfba46.
//
// Solidity: function zeroAffine() pure returns(uint256 x, uint256 y)
func (_EllipticCurve *EllipticCurveCaller) ZeroAffine(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "zeroAffine")

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ZeroAffine is a free data retrieval call binding the contract method 0x84dfba46.
//
// Solidity: function zeroAffine() pure returns(uint256 x, uint256 y)
func (_EllipticCurve *EllipticCurveSession) ZeroAffine() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _EllipticCurve.Contract.ZeroAffine(&_EllipticCurve.CallOpts)
}

// ZeroAffine is a free data retrieval call binding the contract method 0x84dfba46.
//
// Solidity: function zeroAffine() pure returns(uint256 x, uint256 y)
func (_EllipticCurve *EllipticCurveCallerSession) ZeroAffine() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _EllipticCurve.Contract.ZeroAffine(&_EllipticCurve.CallOpts)
}

// ZeroProj is a free data retrieval call binding the contract method 0x14c67060.
//
// Solidity: function zeroProj() pure returns(uint256 x, uint256 y, uint256 z)
func (_EllipticCurve *EllipticCurveCaller) ZeroProj(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
	Z *big.Int
}, error) {
	var out []interface{}
	err := _EllipticCurve.contract.Call(opts, &out, "zeroProj")

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
		Z *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Z = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ZeroProj is a free data retrieval call binding the contract method 0x14c67060.
//
// Solidity: function zeroProj() pure returns(uint256 x, uint256 y, uint256 z)
func (_EllipticCurve *EllipticCurveSession) ZeroProj() (struct {
	X *big.Int
	Y *big.Int
	Z *big.Int
}, error) {
	return _EllipticCurve.Contract.ZeroProj(&_EllipticCurve.CallOpts)
}

// ZeroProj is a free data retrieval call binding the contract method 0x14c67060.
//
// Solidity: function zeroProj() pure returns(uint256 x, uint256 y, uint256 z)
func (_EllipticCurve *EllipticCurveCallerSession) ZeroProj() (struct {
	X *big.Int
	Y *big.Int
	Z *big.Int
}, error) {
	return _EllipticCurve.Contract.ZeroProj(&_EllipticCurve.CallOpts)
}
