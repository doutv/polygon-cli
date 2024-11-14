// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifyingpaymaster

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

// VerifyingPaymasterMetaData contains all meta data concerning the VerifyingPaymaster contract.
var VerifyingPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifyingSigner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"}],\"name\":\"parsePaymasterAndData\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyingSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c0604090808252346101d2578181610ffe803803809161002082856101d7565b8339810103126101d25780516001600160a01b03918282168083036101d2576020809201519380851685036101d25733156101ba5760008054336001600160a01b031982168117835588519492938693869360249385939091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08880a36301ffc9a760e01b825263122a0e9b60e31b60048301525afa9182156101ae57819261016d575b50501561012a575060805260a05251610ded9081610211823960805181818161017501528181610223015281816102d50152818161034b015281816103b6015281816106e1015281816107660152610a65015260a0518181816106850152610b5b0152f35b60649084519062461bcd60e51b82526004820152601e60248201527f49456e747279506f696e7420696e74657266616365206d69736d6174636800006044820152fd5b9091508281813d83116101a7575b61018581836101d7565b810103126101a357519081151582036101a0575038806100c5565b80fd5b5080fd5b503d61017b565b508551903d90823e3d90fd5b8551631e4fbdf760e01b815260006004820152602490fd5b600080fd5b601f909101601f19168101906001600160401b038211908210176101fa57604052565b634e487b7160e01b600052604160045260246000fdfe60406080815260048036101561001457600080fd5b600091823560e01c80630396cb601461073e57838163205c2878146106b45750806323d9ac9b1461067057806352b7512c146105d05780635829c5f514610563578063715018a6146105095780637c627b211461048f5780638da5cb5b1461046757806394d4ad60146103e9578063b0d691fe146103a157838163bb9fe6bf1461032e578163c23a5cea146102a757508063c399ec88146101f657838163d0e30db014610165575063f2fde38b146100cb57600080fd5b34610161576020366003190112610161576001600160a01b0382358181169391929084900361015d576100fc610a37565b8315610147575050600054826bffffffffffffffffffffffff60a01b821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a380f35b51631e4fbdf760e01b8152908101849052602490fd5b8480fd5b8280fd5b808484826003193601126101f2577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156101ed578390602483518095819363b760faf960e01b8352309083015234905af19081156101e457506101d15750f35b6101da90610823565b6101e15780f35b80fd5b513d84823e3d90fd5b505050fd5b5050fd5b503461016157826003193601126101615780516370a0823160e01b815230928101929092526020826024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa91821561029d578392610265575b6020838351908152f35b9091506020813d602011610295575b8161028160209383610869565b81010312610161576020925051903861025b565b3d9150610274565b81513d85823e3d90fd5b808484346101f25760203660031901126101f2576102c36107cc565b6102cb610a37565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116803b1561032a578592836024928651978895869463611d2e7560e11b865216908401525af19081156101e457506101d15750f35b8580fd5b808484346101f257826003193601126101f257610349610a37565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156101ed57815163bb9fe6bf60e01b81529284918491829084905af19081156101e457506101d15750f35b8382346103e557816003193601126103e557517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b509190346103e55760203660031901126103e557803567ffffffffffffffff811161016157610429610423608094936060933691016107f5565b906109ef565b938785939694985198899765ffffffffffff809216895216602088015286015281606086015285850137828201840152601f01601f19168101030190f35b8382346103e557816003193601126103e557905490516001600160a01b039091168152602090f35b5091346101e15760803660031901126101e1576003823510156101e1576024359067ffffffffffffffff82116101e157506064926104d2602092369085016107f5565b50506104dc610a63565b5162461bcd60e51b815291820152600d60248201526c6d757374206f7665727269646560981b6044820152fd5b83346101e157806003193601126101e157610522610a37565b600080546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b50903461016157600319926060368501126101e15781359367ffffffffffffffff85116103e5576101209085360301126101e1575065ffffffffffff60243581811681036105cb5760443591821682036105cb576020946105c49301610905565b9051908152f35b600080fd5b5082346101e157606092600319906060823601126101615780359167ffffffffffffffff831161066c576101209083360301126101615761061c84928692610616610a63565b01610ad2565b8391935194859383855285518094860152815b848110610655575050606080955083850101526020830152601f80199101168101030190f35b60208782018101518983018401528896500161062f565b8380fd5b8382346103e557816003193601126103e557517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b808484346101f257806003193601126101f2576106cf6107cc565b6106d7610a37565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116803b1561032a578592836044928651978895869463040b850f60e31b8652169084015260243560248401525af19081156101e457506101d15750f35b5060203660031901126101615782823563ffffffff81168091036103e557610764610a37565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031693843b156101615760249084519586938492621cb65b60e51b845283015234905af19081156101e457506107c0575080f35b6107c990610823565b80f35b600435906001600160a01b03821682036105cb57565b359065ffffffffffff821682036105cb57565b9181601f840112156105cb5782359167ffffffffffffffff83116105cb57602083818601950101116105cb57565b67ffffffffffffffff811161083757604052565b634e487b7160e01b600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761083757604052565b90601f8019910116810190811067ffffffffffffffff82111761083757604052565b903590601e19813603018212156105cb570180359067ffffffffffffffff82116105cb576020019181360383136105cb57565b92919267ffffffffffffffff821161083757604051916108e8601f8201601f191660200184610869565b8294818452818301116105cb578281602093846000960137010152565b909161091e610917604084018461088b565b36916108be565b6020815191012092610936610917606085018561088b565b602081519101209261094b60e082018261088b565b6034929192116105cb57601460c09260405196602088019860018060a01b038535168a52602085013560408a015260608901526080880152608083013560a088015201358286015260a081013560e08601520135610100840152466101208401523061014084015265ffffffffffff80911661016084015261018091168183015281526101a0810181811067ffffffffffffffff8211176108375760405251902090565b9190806034116105cb576040603319848381010301126105cb57610a15603484016107e2565b91610a22605485016107e2565b9293826074116105cb57607401916073190190565b6000546001600160a01b03163303610a4b57565b60405163118cdaa760e01b8152336004820152602490fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610a9557565b60405162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08115b9d1c9e541bda5b9d605a1b6044820152606490fd5b90610ae361042360e084018461088b565b604081969396148015610c5c575b15610bf257610b3f610b4592610b0b8887610b4e97610905565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52603c6000209236916108be565b90610c66565b90929192610ca2565b6001600160a01b039081167f000000000000000000000000000000000000000000000000000000000000000090911603610bb857604051610b8e8161084d565b600081529260a09190911b65ffffffffffff60a01b1660d09190911b6001600160d01b0319161790565b600190604051610bc78161084d565b600081529360a09190911b65ffffffffffff60a01b1660d09190911b6001600160d01b031916171790565b608460405162461bcd60e51b815260206004820152604060248201527f566572696679696e675061796d61737465723a20696e76616c6964207369676e60448201527f6174757265206c656e67746820696e207061796d6173746572416e64446174616064820152fd5b5060418114610af1565b8151919060418303610c9757610c9092506020820151906060604084015193015160001a90610d27565b9192909190565b505060009160029190565b6004811015610d115780610cb4575050565b60018103610cce5760405163f645eedf60e01b8152600490fd5b60028103610cef5760405163fce698f760e01b815260048101839052602490fd5b600314610cf95750565b602490604051906335e2f38360e21b82526004820152fd5b634e487b7160e01b600052602160045260246000fd5b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411610dab57926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa15610d9f5780516001600160a01b03811615610d9657918190565b50809160019190565b604051903d90823e3d90fd5b5050506000916003919056fea2646970667358221220c5d5c4b1e491020df4f8f5ee9fedde997aa7cbc202b3efdb89c69566d834e11764736f6c63430008190033",
}

// VerifyingPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifyingPaymasterMetaData.ABI instead.
var VerifyingPaymasterABI = VerifyingPaymasterMetaData.ABI

// VerifyingPaymasterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifyingPaymasterMetaData.Bin instead.
var VerifyingPaymasterBin = VerifyingPaymasterMetaData.Bin

// DeployVerifyingPaymaster deploys a new Ethereum contract, binding an instance of VerifyingPaymaster to it.
func DeployVerifyingPaymaster(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address, _verifyingSigner common.Address) (common.Address, *types.Transaction, *VerifyingPaymaster, error) {
	parsed, err := VerifyingPaymasterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifyingPaymasterBin), backend, _entryPoint, _verifyingSigner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifyingPaymaster{VerifyingPaymasterCaller: VerifyingPaymasterCaller{contract: contract}, VerifyingPaymasterTransactor: VerifyingPaymasterTransactor{contract: contract}, VerifyingPaymasterFilterer: VerifyingPaymasterFilterer{contract: contract}}, nil
}

// VerifyingPaymaster is an auto generated Go binding around an Ethereum contract.
type VerifyingPaymaster struct {
	VerifyingPaymasterCaller     // Read-only binding to the contract
	VerifyingPaymasterTransactor // Write-only binding to the contract
	VerifyingPaymasterFilterer   // Log filterer for contract events
}

// VerifyingPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifyingPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyingPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifyingPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyingPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifyingPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyingPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifyingPaymasterSession struct {
	Contract     *VerifyingPaymaster // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifyingPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifyingPaymasterCallerSession struct {
	Contract *VerifyingPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// VerifyingPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifyingPaymasterTransactorSession struct {
	Contract     *VerifyingPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// VerifyingPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifyingPaymasterRaw struct {
	Contract *VerifyingPaymaster // Generic contract binding to access the raw methods on
}

// VerifyingPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifyingPaymasterCallerRaw struct {
	Contract *VerifyingPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// VerifyingPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifyingPaymasterTransactorRaw struct {
	Contract *VerifyingPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifyingPaymaster creates a new instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymaster(address common.Address, backend bind.ContractBackend) (*VerifyingPaymaster, error) {
	contract, err := bindVerifyingPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymaster{VerifyingPaymasterCaller: VerifyingPaymasterCaller{contract: contract}, VerifyingPaymasterTransactor: VerifyingPaymasterTransactor{contract: contract}, VerifyingPaymasterFilterer: VerifyingPaymasterFilterer{contract: contract}}, nil
}

// NewVerifyingPaymasterCaller creates a new read-only instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymasterCaller(address common.Address, caller bind.ContractCaller) (*VerifyingPaymasterCaller, error) {
	contract, err := bindVerifyingPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterCaller{contract: contract}, nil
}

// NewVerifyingPaymasterTransactor creates a new write-only instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifyingPaymasterTransactor, error) {
	contract, err := bindVerifyingPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterTransactor{contract: contract}, nil
}

// NewVerifyingPaymasterFilterer creates a new log filterer instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifyingPaymasterFilterer, error) {
	contract, err := bindVerifyingPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterFilterer{contract: contract}, nil
}

// bindVerifyingPaymaster binds a generic wrapper to an already deployed contract.
func bindVerifyingPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifyingPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifyingPaymaster *VerifyingPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifyingPaymaster.Contract.VerifyingPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifyingPaymaster *VerifyingPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.VerifyingPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifyingPaymaster *VerifyingPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.VerifyingPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifyingPaymaster *VerifyingPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifyingPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifyingPaymaster *VerifyingPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifyingPaymaster *VerifyingPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterSession) EntryPoint() (common.Address, error) {
	return _VerifyingPaymaster.Contract.EntryPoint(&_VerifyingPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _VerifyingPaymaster.Contract.EntryPoint(&_VerifyingPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VerifyingPaymaster *VerifyingPaymasterSession) GetDeposit() (*big.Int, error) {
	return _VerifyingPaymaster.Contract.GetDeposit(&_VerifyingPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) GetDeposit() (*big.Int, error) {
	return _VerifyingPaymaster.Contract.GetDeposit(&_VerifyingPaymaster.CallOpts)
}

// GetHash is a free data retrieval call binding the contract method 0x5829c5f5.
//
// Solidity: function getHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, uint48 validUntil, uint48 validAfter) view returns(bytes32)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) GetHash(opts *bind.CallOpts, userOp PackedUserOperation, validUntil *big.Int, validAfter *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "getHash", userOp, validUntil, validAfter)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0x5829c5f5.
//
// Solidity: function getHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, uint48 validUntil, uint48 validAfter) view returns(bytes32)
func (_VerifyingPaymaster *VerifyingPaymasterSession) GetHash(userOp PackedUserOperation, validUntil *big.Int, validAfter *big.Int) ([32]byte, error) {
	return _VerifyingPaymaster.Contract.GetHash(&_VerifyingPaymaster.CallOpts, userOp, validUntil, validAfter)
}

// GetHash is a free data retrieval call binding the contract method 0x5829c5f5.
//
// Solidity: function getHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, uint48 validUntil, uint48 validAfter) view returns(bytes32)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) GetHash(userOp PackedUserOperation, validUntil *big.Int, validAfter *big.Int) ([32]byte, error) {
	return _VerifyingPaymaster.Contract.GetHash(&_VerifyingPaymaster.CallOpts, userOp, validUntil, validAfter)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterSession) Owner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.Owner(&_VerifyingPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) Owner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.Owner(&_VerifyingPaymaster.CallOpts)
}

// ParsePaymasterAndData is a free data retrieval call binding the contract method 0x94d4ad60.
//
// Solidity: function parsePaymasterAndData(bytes paymasterAndData) pure returns(uint48 validUntil, uint48 validAfter, bytes signature)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) ParsePaymasterAndData(opts *bind.CallOpts, paymasterAndData []byte) (struct {
	ValidUntil *big.Int
	ValidAfter *big.Int
	Signature  []byte
}, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "parsePaymasterAndData", paymasterAndData)

	outstruct := new(struct {
		ValidUntil *big.Int
		ValidAfter *big.Int
		Signature  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidUntil = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidAfter = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Signature = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ParsePaymasterAndData is a free data retrieval call binding the contract method 0x94d4ad60.
//
// Solidity: function parsePaymasterAndData(bytes paymasterAndData) pure returns(uint48 validUntil, uint48 validAfter, bytes signature)
func (_VerifyingPaymaster *VerifyingPaymasterSession) ParsePaymasterAndData(paymasterAndData []byte) (struct {
	ValidUntil *big.Int
	ValidAfter *big.Int
	Signature  []byte
}, error) {
	return _VerifyingPaymaster.Contract.ParsePaymasterAndData(&_VerifyingPaymaster.CallOpts, paymasterAndData)
}

// ParsePaymasterAndData is a free data retrieval call binding the contract method 0x94d4ad60.
//
// Solidity: function parsePaymasterAndData(bytes paymasterAndData) pure returns(uint48 validUntil, uint48 validAfter, bytes signature)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) ParsePaymasterAndData(paymasterAndData []byte) (struct {
	ValidUntil *big.Int
	ValidAfter *big.Int
	Signature  []byte
}, error) {
	return _VerifyingPaymaster.Contract.ParsePaymasterAndData(&_VerifyingPaymaster.CallOpts, paymasterAndData)
}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) VerifyingSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "verifyingSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterSession) VerifyingSigner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.VerifyingSigner(&_VerifyingPaymaster.CallOpts)
}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) VerifyingSigner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.VerifyingSigner(&_VerifyingPaymaster.CallOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.AddStake(&_VerifyingPaymaster.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.AddStake(&_VerifyingPaymaster.TransactOpts, unstakeDelaySec)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) Deposit() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.Deposit(&_VerifyingPaymaster.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) Deposit() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.Deposit(&_VerifyingPaymaster.TransactOpts)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.PostOp(&_VerifyingPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.PostOp(&_VerifyingPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.RenounceOwnership(&_VerifyingPaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.RenounceOwnership(&_VerifyingPaymaster.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.TransferOwnership(&_VerifyingPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.TransferOwnership(&_VerifyingPaymaster.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) UnlockStake() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.UnlockStake(&_VerifyingPaymaster.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.UnlockStake(&_VerifyingPaymaster.TransactOpts)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_VerifyingPaymaster *VerifyingPaymasterSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.ValidatePaymasterUserOp(&_VerifyingPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.ValidatePaymasterUserOp(&_VerifyingPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawStake(&_VerifyingPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawStake(&_VerifyingPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawTo(&_VerifyingPaymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawTo(&_VerifyingPaymaster.TransactOpts, withdrawAddress, amount)
}

// VerifyingPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterOwnershipTransferredIterator struct {
	Event *VerifyingPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterOwnershipTransferred)
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
		it.Event = new(VerifyingPaymasterOwnershipTransferred)
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
func (it *VerifyingPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VerifyingPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterOwnershipTransferredIterator{contract: _VerifyingPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterOwnershipTransferred)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*VerifyingPaymasterOwnershipTransferred, error) {
	event := new(VerifyingPaymasterOwnershipTransferred)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
