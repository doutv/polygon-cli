// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenpaymaster

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

// OracleHelperOracleHelperConfig is an auto generated low-level Go binding around an user-defined struct.
type OracleHelperOracleHelperConfig struct {
	CacheTimeToLive      *big.Int
	MaxOracleRoundAge    *big.Int
	TokenOracle          common.Address
	NativeOracle         common.Address
	TokenToNativeOracle  bool
	TokenOracleReverse   bool
	NativeOracleReverse  bool
	PriceUpdateThreshold *big.Int
}

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

// TokenPaymasterTokenPaymasterConfig is an auto generated low-level Go binding around an user-defined struct.
type TokenPaymasterTokenPaymasterConfig struct {
	PriceMarkup          *big.Int
	MinEntryPointBalance *big.Int
	RefundPostopCost     *big.Int
	PriceMaxAge          *big.Int
}

// UniswapHelperUniswapHelperConfig is an auto generated low-level Go binding around an user-defined struct.
type UniswapHelperUniswapHelperConfig struct {
	MinSwapAmount  *big.Int
	UniswapPoolFee *big.Int
	Slippage       uint8
}

// TokenPaymasterMetaData contains all meta data concerning the TokenPaymaster contract.
var TokenPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_wrappedNative\",\"type\":\"address\"},{\"internalType\":\"contractISwapRouter\",\"name\":\"_uniswap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"priceMarkup\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"minEntryPointBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint48\",\"name\":\"refundPostopCost\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"priceMaxAge\",\"type\":\"uint48\"}],\"internalType\":\"structTokenPaymaster.TokenPaymasterConfig\",\"name\":\"_tokenPaymasterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint48\",\"name\":\"cacheTimeToLive\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"maxOracleRoundAge\",\"type\":\"uint48\"},{\"internalType\":\"contractIOracle\",\"name\":\"tokenOracle\",\"type\":\"address\"},{\"internalType\":\"contractIOracle\",\"name\":\"nativeOracle\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"tokenToNativeOracle\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"tokenOracleReverse\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"nativeOracleReverse\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"priceUpdateThreshold\",\"type\":\"uint256\"}],\"internalType\":\"structOracleHelper.OracleHelperConfig\",\"name\":\"_oracleHelperConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minSwapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"uniswapPoolFee\",\"type\":\"uint24\"},{\"internalType\":\"uint8\",\"name\":\"slippage\",\"type\":\"uint8\"}],\"internalType\":\"structUniswapHelper.UniswapHelperConfig\",\"name\":\"_uniswapHelperConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"priceMarkup\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"minEntryPointBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint48\",\"name\":\"refundPostopCost\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"priceMaxAge\",\"type\":\"uint48\"}],\"indexed\":false,\"internalType\":\"structTokenPaymaster.TokenPaymasterConfig\",\"name\":\"tokenPaymasterConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cachedPriceTimestamp\",\"type\":\"uint256\"}],\"name\":\"TokenPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"name\":\"UniswapReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualTokenCharge\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualTokenPriceWithMarkup\",\"type\":\"uint256\"}],\"name\":\"UserOperationSponsored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cachedPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cachedPriceTimestamp\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"priceMarkup\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"minEntryPointBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint48\",\"name\":\"refundPostopCost\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"priceMaxAge\",\"type\":\"uint48\"}],\"internalType\":\"structTokenPaymaster.TokenPaymasterConfig\",\"name\":\"_tokenPaymasterConfig\",\"type\":\"tuple\"}],\"name\":\"setTokenPaymasterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minSwapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"uniswapPoolFee\",\"type\":\"uint24\"},{\"internalType\":\"uint8\",\"name\":\"slippage\",\"type\":\"uint8\"}],\"internalType\":\"structUniswapHelper.UniswapHelperConfig\",\"name\":\"_uniswapHelperConfig\",\"type\":\"tuple\"}],\"name\":\"setUniswapConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPaymasterConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceMarkup\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"minEntryPointBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint48\",\"name\":\"refundPostopCost\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"priceMaxAge\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"tokenToWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswap\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"}],\"name\":\"updateCachedPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"weiToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedNative\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x610100604052346106bb5761246d803803809161001e82610100610821565b6101003961028081126106bb5761010051906001600160a01b03821682036106bb5761012051906001600160a01b03821682036106bb5761014051926001600160a01b03841684036106bb57610160516001600160a01b03811681036106bb576080607f198401126106bb5760405194608086016001600160401b0381118782101761080b576040526101805186526101a0516001600160801b03811681036106bb5760208701526100d16101c0610844565b60408701526100e16101e0610844565b606087015261010060ff198501126106bb57604051936001600160401b0361010086019081119086111761080b576101008501604052606090610125610200610844565b8652610132610220610844565b6020870152610142610240610857565b6040870152610152610260610857565b8683015261016161028061086b565b60808701526101716102a061086b565b60a08701526101816102c061086b565b60c08701526102e05160e08701526101ff1901126106bb5760405192606084016001600160401b0381118582101761080b576040526103005184526103205162ffffff811681036106bb5760208501526101dc610340610878565b604085015261036051956001600160a01b03871687036106bb576101ff33610886565b6040516301ffc9a760e01b815263122a0e9b60e31b60048201526020816024816001600160a01b0386165afa9081156106c8576000916107d1575b501561078c5760805260405163095ea7b360e01b81526001600160a01b03848116600483015260001960248301526020908290604490829060009087165af180156106c857610753575b5060018060a01b031660c05260e05260a052805160015562ffffff60208201511663ff000000604060025493015160181b169163ffffffff1916171760025560001960035565ffffffffffff8151166bffffffffffff000000000000602083015160301b169060018060601b0319604084015160601b169117178060055560018060a01b03606083015116906a52b7d2dcc80cd2e400000060e060065460ff60a01b6080870151151560a01b1660ff60a81b60a0880151151560a81b16908660ff60b01b60c08a0151151560b01b169360018060b81b0319161717171794856006550151806007551161070e5760206004916040519283809263313ce56760e01b825260601c5afa80156106c8576000906106d4575b60ff91506001600160801b03906103b0906108eb565b600880546001600160801b0319169290911691821790559260a01c16156106415750506006546001600160a01b03166105fc57600880546001600160801b0316600160801b1790555b610401610912565b6aa56fa5b99019a5c80000008251116105b7576a52b7d2dcc80cd2e4000000825110610572577fcd938817f1c47094d43be3d07e8c67e11766db2e11a2b4376e7ee937b15793a260808365ffffffffffff60606104cf96519283600955600180861b03602082015116600a549084871b6040840151881b16908560b01b8585015160b01b169263ffffffff60e01b16171717600a55604051938452600180861b0360208201511660208501528260408201511660408501520151166060820152a16104ca610912565b610886565b604051611b2e908161093f82396080518181816104e3015281816105db0152818161068b015281816107020152818161076d01528181610a6201528181610fc70152818161111101526116d9015260a051818181610ade01528181610f6a0152611942015260c0518181816101780152818161086501528181610a1c01528181610ab701528181610c040152610da2015260e05181818161044a015261191b0152f35b60405162461bcd60e51b815260206004820152601960248201527f54504d3a207072696365206d61726b757020746f6f206c6f77000000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601a60248201527f54504d3a207072696365206d61726b757020746f6f20686967680000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601f60248201527f54504d3a206e6174697665206f7261636c65206d757374206265207a65726f006044820152606490fd5b60206004916040519283809263313ce56760e01b82525afa9081156106c857600091610689575b506001600160801b03199061067c906108eb565b60801b16176008556103f9565b90506020813d6020116106c0575b816106a460209383610821565b810103126106bb576106b590610878565b38610668565b600080fd5b3d9150610697565b6040513d6000823e3d90fd5b506020813d602011610706575b816106ee60209383610821565b810103126106bb5761070160ff91610878565b61039a565b3d91506106e1565b60405162461bcd60e51b815260206004820152601e60248201527f54504d3a20757064617465207468726573686f6c6420746f6f206869676800006044820152606490fd5b6020813d602011610784575b8161076c60209383610821565b810103126106bb5761077d9061086b565b5038610284565b3d915061075f565b60405162461bcd60e51b815260206004820152601e60248201527f49456e747279506f696e7420696e74657266616365206d69736d6174636800006044820152606490fd5b90506020813d602011610803575b816107ec60209383610821565b810103126106bb576107fd9061086b565b3861023a565b3d91506107df565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761080b57604052565b519065ffffffffffff821682036106bb57565b51906001600160a01b03821682036106bb57565b519081151582036106bb57565b519060ff821682036106bb57565b6001600160a01b039081169081156108d2576000548260018060a01b0319821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b604051631e4fbdf760e01b815260006004820152602490fd5b60ff16604d81116108fc57600a0a90565b634e487b7160e01b600052601160045260246000fd5b6000546001600160a01b0316330361092657565b60405163118cdaa760e01b8152336004820152602490fdfe604060808152600480361015610048575b50361561001c57600080fd5b513481527f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f8852587460203392a2005b600090813560e01c9081630396cb60146110e75781631b9a91a414611024578163205c287814610f995781632681f7e414610f555781633ba9290f14610f2a57816352b7512c14610c9a578163715018a614610c405781637c627b211461090c5781637c986aac146108b45781638da5cb5b1461088c5781639e281a981461083b578163a0840fa71461079c578163b0d691fe14610758578163bb9fe6bf146106e4578163c23a5cea1461065c578163c399ec88146105ad578163cb721cfd14610561578163d0e30db0146104d2578163d7a23b3c146104a2578163e1d8153c14610479578163eb6d3a1114610435578163f14d64ed14610260578163f2fde38b146101ca57508063f60fdcb3146101ab5763fc0c546a036100105790346101a757816003193601126101a757517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b5090346101a757816003193601126101a7576020906003549051908152f35b8391503461025c57602036600319011261025c576101e6611174565b906101ef61145f565b6001600160a01b03918216928315610246575050600054826bffffffffffffffffffffffff60a01b821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a380f35b51631e4fbdf760e01b8152908101849052602490fd5b8280fd5b9050346101a75760803660031901126101a75782516080810181811067ffffffffffffffff82111761042057845281358152602435916001600160801b03808416840361041c57602083019384526044359065ffffffffffff9384831683036104175787810192835260643593858516850361041757606082019485526102e561145f565b6aa56fa5b99019a5c80000008251116103d4576a52b7d2dcc80cd2e4000000825110610391575084608095937fcd938817f1c47094d43be3d07e8c67e11766db2e11a2b4376e7ee937b15793a29795938a9351968760095581815116600a549065ffffffffffff8b1b85518c1b169065ffffffffffff60b01b895160b01b169263ffffffff60e01b16171717600a5584519788525116602087015251169084015251166060820152a180f35b60649060208a519162461bcd60e51b8352820152601960248201527f54504d3a207072696365206d61726b757020746f6f206c6f77000000000000006044820152fd5b60649060208a519162461bcd60e51b8352820152601a60248201527f54504d3a207072696365206d61726b757020746f6f20686967680000000000006044820152fd5b600080fd5b8480fd5b604183634e487b7160e01b6000525260246000fd5b8284346101a757816003193601126101a757517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b83833461049f578060031936011261049f575065ffffffffffff60209254169051908152f35b80fd5b8284346101a7576020906a52b7d2dcc80cd2e40000006104ca6104c43661118a565b9061120f565b049051908152f35b905081928260031936011261055d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b15610558578390602483518095819363b760faf960e01b8352309083015234905af190811561054f575061053f5750f35b610548906111a0565b61049f5780f35b513d84823e3d90fd5b505050fd5b5050fd5b8284346101a757816003193601126101a75760809060095490600a5465ffffffffffff9180519384526001600160801b03821660208501528282861c169084015260b01c166060820152f35b8284346101a757816003193601126101a75780516370a0823160e01b815230938101939093526020836024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa918215610651579161061c575b6020925051908152f35b90506020823d602011610649575b81610637602093836111ca565b81010312610417576020915190610612565b3d915061062a565b9051903d90823e3d90fd5b905081923461055d57602036600319011261055d57610679611174565b61068161145f565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116803b156106e0578592836024928651978895869463611d2e7560e11b865216908401525af190811561054f575061053f5750f35b8580fd5b905081923461055d578260031936011261055d5761070061145f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b1561055857815163bb9fe6bf60e01b81529284918491829084905af190811561054f575061053f5750f35b8284346101a757816003193601126101a757517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b8391503461025c57606036600319011261025c578151906060820182811067ffffffffffffffff82111761082657835235815260243562ffffff91828216820361041c57602081019182526044359360ff851685036106e057810193845261080261145f565b51600155511663ff000000600254925160181b169163ffffffff1916171760025580f35b604182634e487b7160e01b6000525260246000fd5b8284346101a75736600319011261049f57610889610857611174565b61085f61145f565b602435907f000000000000000000000000000000000000000000000000000000000000000061178d565b80f35b8284346101a757816003193601126101a757905490516001600160a01b039091168152602090f35b83833461049f57506108c53661118a565b926a52b7d2dcc80cd2e4000000918281029281840414901517156108f757506020926108f091611222565b9051908152f35b601190634e487b7160e01b6000525260246000fd5b8284346101a75760803660031901126101a7576003833510156101a75760249081359367ffffffffffffffff80861161041c573660238701121561041c578582013590811161041c578501368482011161041c576044359061096c6116d7565b83876009549203126106e05760448588013597013560018060a01b03918282168092036104175761099b611242565b936109b46a52b7d2dcc80cd2e400000092838702611222565b9160643565ffffffffffff600a5460801c16028201818102918183041490151715610c2e579982606092610a0a7f46caa0511cf037f06f57a0bf273a2ff04229f5b12fb04675234a6cbe2e7f1a89958d9e611222565b928380821115610bed57610a409103877f000000000000000000000000000000000000000000000000000000000000000061178d565b8951928352602083015288820152a283516370a0823160e01b815230848201527f0000000000000000000000000000000000000000000000000000000000000000821692906020818881875afa908115610be3578891610bae575b506001600160801b03600a541611610ab1578680f35b610adb907f0000000000000000000000000000000000000000000000000000000000000000611890565b907f000000000000000000000000000000000000000000000000000000000000000016803b15610baa57845163125012df60e21b8152848101928352306020840152918791839182908490829060400103925af18015610ba057908691610b8c575b50504793813b156106e05785928451958693849263b760faf960e01b845230908401525af190811561054f5750610b78575b80808080808680f35b610b81906111a0565b61049f578082610b6f565b610b95906111a0565b61041c578487610b3d565b84513d88823e3d90fd5b8680fd5b9750506020873d602011610bdb575b81610bca602093836111ca565b810103126104175787965189610a9b565b3d9150610bbd565b86513d8a823e3d90fd5b8110610bfa575b50610a40565b610c2890840330887f00000000000000000000000000000000000000000000000000000000000000006117df565b8d610bf4565b634e487b7160e01b8a5260118752888afd5b823461049f578060031936011261049f57610c5961145f565b600080546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b9050823461025c576060926003199060608236011261049f57833567ffffffffffffffff81116101a7576101208186019382360301126101a757610cdc6116d7565b6009549060e48101610cee8186611746565b6033190193905083158015610f20575b15610edd5765ffffffffffff9283600a5460801c1690610d1e8489611746565b603411610ed9576024013560801c821015610e965760c46001600160801b03910135160260443501610d606a52b7d2dcc80cd2e4000000928360035402611222565b926020809614610e71575b50818102918183041490151715610e5e5791610dcb610d9082938b989a999795611222565b98610dc68a610d9e83611779565b30907f00000000000000000000000000000000000000000000000000000000000000006117df565b611779565b855160208101999099526001600160a01b03166040808a01919091528852601f1996610df860608a6111ca565b541665ffffffffffff60a01b91600a5460b01c160160a01b1691835196879585875281518096880152835b868110610e4857876060818b601f8b8b8b868387010152602085015201168101030190f35b8281018401518a82018301528997508301610e23565b634e487b7160e01b855260118852602485fd5b610e7b9088611746565b605411610baa576034013583811015610d6b5792508a610d6b565b885162461bcd60e51b81526020818c0152601b60248201527f54504d3a20706f73744f704761734c696d697420746f6f206c6f7700000000006044820152606490fd5b8780fd5b865162461bcd60e51b81526020818a0152601860248201527f54504d3a20696e76616c69642064617461206c656e67746800000000000000006044820152606490fd5b5060208414610cfe565b8391503461025c57602036600319011261025c573591821515830361049f57506108f060209261136a565b8284346101a757816003193601126101a757517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b905081923461055d578060031936011261055d57610fb5611174565b610fbd61145f565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116803b156106e0578592836044928651978895869463040b850f60e31b8652169084015260243560248401525af190811561054f575061053f5750f35b9050823461025c578060031936011261025c5782808080611043611174565b61104b61145f565b602435906001600160a01b03165af13d156110e2573d67ffffffffffffffff81116110cf57825190611087601f8201601f1916602001836111ca565b81528460203d92013e5b1561109a578280f35b906020606492519162461bcd60e51b8352820152600f60248201526e1dda5d1a191c985dc819985a5b1959608a1b6044820152fd5b634e487b7160e01b855260418452602485fd5b611091565b905082602036600319011261025c5782823563ffffffff81168091036101a75761110f61145f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031693843b1561025c5760249084519586938492621cb65b60e51b845283015234905af190811561054f575061116b575080f35b610889906111a0565b600435906001600160a01b038216820361041757565b6040906003190112610417576004359060243590565b67ffffffffffffffff81116111b457604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff8211176111b457604052565b919082039182116111f957565b634e487b7160e01b600052601160045260246000fd5b818102929181159184041417156111f957565b811561122c570490565b634e487b7160e01b600052601260045260246000fd5b60055465ffffffffffff9060045482821661125f848316426111ec565b111561136157600754926112a060065461127e6003549560601c6114a2565b600160ff8360a01c1615611349575b60ff808460b01c169360a81c1691611633565b936a52b7d2dcc80cd2e400000085810281810487036111f957856112c391611222565b8282018083116111f9578111928315611335575b5050501561132e57916060917ed4fe314618b73a96886b87817a53a5ed51433b0234c85a5e9dafe2cb7b884293856003554216809165ffffffffffff1916176004556040519185835260208301526040820152a190565b5050905090565b820392509082116111f957103880806112d7565b5061135c6001600160a01b0383166114a2565b61128d565b50505060035490565b60055460045465ffffffffffff611383828216426111ec565b841580918192611452575b50611447576007546113ab60065461127e6003549760601c6114a2565b956a52b7d2dcc80cd2e40000009187830283810489036111f957876113cf91611222565b9193611433575b8315611335575050501561132e57916060917ed4fe314618b73a96886b87817a53a5ed51433b0234c85a5e9dafe2cb7b884293856003554216809165ffffffffffff1916176004556040519185835260208301526040820152a190565b80935082018083116111f9578111926113d6565b505050505060035490565b905082851610153861138e565b6000546001600160a01b0316330361147357565b60405163118cdaa760e01b8152336004820152602490fd5b519069ffffffffffffffffffff8216820361041757565b604051633fabe5a360e21b81529060a090829060049082906001600160a01b03165afa80156116275760008092819082936115d0575b50600084131561158b576114f965ffffffffffff60055460301c16426111ec565b1161154e5769ffffffffffffffffffff8091169116106115165790565b60405162461bcd60e51b815260206004820152601060248201526f54504d3a205374616c6520707269636560801b6044820152606490fd5b60405162461bcd60e51b81526020600482015260156024820152741514134e88125b98dbdb5c1b195d19481c9bdd5b99605a1b6044820152606490fd5b60405162461bcd60e51b815260206004820152601960248201527f54504d3a20436861696e6c696e6b207072696365203c3d2030000000000000006044820152606490fd5b935050905060a0823d60a01161161f575b816115ee60a093836111ca565b8101031261049f57506116008161148b565b602082015161161660806060850151940161148b565b919092386114d8565b3d91506115e1565b6040513d6000823e3d90fd5b9091156116a3576001600160801b0360085416906a52b7d2dcc80cd2e400000091808302928304036111f95761166891611222565b915b1561168b576116889161167c9161120f565b60085460801c90611222565b90565b61169e6116889260085460801c9061120f565b611222565b6a52b7d2dcc80cd2e400000090808202918204036111f9576116d1906001600160801b036008541690611222565b9161166a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361170957565b60405162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08115b9d1c9e541bda5b9d605a1b6044820152606490fd5b903590601e1981360301821215610417570180359067ffffffffffffffff82116104175760200191813603831361041757565b356001600160a01b03811681036104175790565b60405163a9059cbb60e01b60208201526001600160a01b039092166024830152604480830193909352918152608081019167ffffffffffffffff8311828410176111b4576117dd92604052611834565b565b6040516323b872dd60e01b60208201526001600160a01b03928316602482015292909116604483015260648083019390935291815260a081019181831067ffffffffffffffff8411176111b4576117dd926040525b906000602091828151910182855af115611627576000513d61188757506001600160a01b0381163b155b6118655750565b604051635274afe760e01b81526001600160a01b039091166004820152602490fd5b6001141561185e565b604080516370a0823160e01b81523060048201526001600160a01b03928316939091906020908184602481895afa938415611aed57600094611abc575b506118e46a52b7d2dcc80cd2e4000000918561120f565b0490600254916103e89060ff8460181c1682039061ffff8083116111f95761190d92169061120f565b04946001548610611ab057807f00000000000000000000000000000000000000000000000000000000000000001692600090827f000000000000000000000000000000000000000000000000000000000000000016865191610100830183811067ffffffffffffffff821117611a9c5788528a835282860187815262ffffff918216848a0190815260608501848152426080870190815260a087018d815260c088018f815260e089018a81528e5163414bf38960e01b815299518c1660048b015295518b1660248a015293519095166044880152905188166064870152516084860152915160a4850152905160c48401525190931660e4820152918390839081845a9261010493f190918282611a6a575b5050611a605750825195865285015283015260608201527ff7edd4c6ec425decf715a8b8eaa3b65d3d86e31ad0ff750aa60fa834190f515f90608090a1600090565b9550505050505090565b909192508382813d8311611a95575b611a8381836111ca565b8101031261049f575051903880611a1e565b503d611a79565b634e487b7160e01b85526041600452602485fd5b50505050505050600090565b90938282813d8311611ae6575b611ad381836111ca565b8101031261049f575051926118e46118cd565b503d611ac9565b83513d6000823e3d90fdfea26469706673582212209b6adbc469eae18f3d3eb1fbf36f77b7265feb7c549f5d035917f21d924d088164736f6c63430008190033",
}

// TokenPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenPaymasterMetaData.ABI instead.
var TokenPaymasterABI = TokenPaymasterMetaData.ABI

// TokenPaymasterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenPaymasterMetaData.Bin instead.
var TokenPaymasterBin = TokenPaymasterMetaData.Bin

// DeployTokenPaymaster deploys a new Ethereum contract, binding an instance of TokenPaymaster to it.
func DeployTokenPaymaster(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _entryPoint common.Address, _wrappedNative common.Address, _uniswap common.Address, _tokenPaymasterConfig TokenPaymasterTokenPaymasterConfig, _oracleHelperConfig OracleHelperOracleHelperConfig, _uniswapHelperConfig UniswapHelperUniswapHelperConfig, _owner common.Address) (common.Address, *types.Transaction, *TokenPaymaster, error) {
	parsed, err := TokenPaymasterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenPaymasterBin), backend, _token, _entryPoint, _wrappedNative, _uniswap, _tokenPaymasterConfig, _oracleHelperConfig, _uniswapHelperConfig, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenPaymaster{TokenPaymasterCaller: TokenPaymasterCaller{contract: contract}, TokenPaymasterTransactor: TokenPaymasterTransactor{contract: contract}, TokenPaymasterFilterer: TokenPaymasterFilterer{contract: contract}}, nil
}

// TokenPaymaster is an auto generated Go binding around an Ethereum contract.
type TokenPaymaster struct {
	TokenPaymasterCaller     // Read-only binding to the contract
	TokenPaymasterTransactor // Write-only binding to the contract
	TokenPaymasterFilterer   // Log filterer for contract events
}

// TokenPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenPaymasterSession struct {
	Contract     *TokenPaymaster   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenPaymasterCallerSession struct {
	Contract *TokenPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TokenPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenPaymasterTransactorSession struct {
	Contract     *TokenPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TokenPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenPaymasterRaw struct {
	Contract *TokenPaymaster // Generic contract binding to access the raw methods on
}

// TokenPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenPaymasterCallerRaw struct {
	Contract *TokenPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// TokenPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenPaymasterTransactorRaw struct {
	Contract *TokenPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenPaymaster creates a new instance of TokenPaymaster, bound to a specific deployed contract.
func NewTokenPaymaster(address common.Address, backend bind.ContractBackend) (*TokenPaymaster, error) {
	contract, err := bindTokenPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenPaymaster{TokenPaymasterCaller: TokenPaymasterCaller{contract: contract}, TokenPaymasterTransactor: TokenPaymasterTransactor{contract: contract}, TokenPaymasterFilterer: TokenPaymasterFilterer{contract: contract}}, nil
}

// NewTokenPaymasterCaller creates a new read-only instance of TokenPaymaster, bound to a specific deployed contract.
func NewTokenPaymasterCaller(address common.Address, caller bind.ContractCaller) (*TokenPaymasterCaller, error) {
	contract, err := bindTokenPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPaymasterCaller{contract: contract}, nil
}

// NewTokenPaymasterTransactor creates a new write-only instance of TokenPaymaster, bound to a specific deployed contract.
func NewTokenPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenPaymasterTransactor, error) {
	contract, err := bindTokenPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPaymasterTransactor{contract: contract}, nil
}

// NewTokenPaymasterFilterer creates a new log filterer instance of TokenPaymaster, bound to a specific deployed contract.
func NewTokenPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenPaymasterFilterer, error) {
	contract, err := bindTokenPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenPaymasterFilterer{contract: contract}, nil
}

// bindTokenPaymaster binds a generic wrapper to an already deployed contract.
func bindTokenPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenPaymaster *TokenPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPaymaster.Contract.TokenPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenPaymaster *TokenPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.TokenPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenPaymaster *TokenPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.TokenPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenPaymaster *TokenPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenPaymaster *TokenPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenPaymaster *TokenPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.contract.Transact(opts, method, params...)
}

// CachedPrice is a free data retrieval call binding the contract method 0xf60fdcb3.
//
// Solidity: function cachedPrice() view returns(uint256)
func (_TokenPaymaster *TokenPaymasterCaller) CachedPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenPaymaster.contract.Call(opts, &out, "cachedPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CachedPrice is a free data retrieval call binding the contract method 0xf60fdcb3.
//
// Solidity: function cachedPrice() view returns(uint256)
func (_TokenPaymaster *TokenPaymasterSession) CachedPrice() (*big.Int, error) {
	return _TokenPaymaster.Contract.CachedPrice(&_TokenPaymaster.CallOpts)
}

// CachedPrice is a free data retrieval call binding the contract method 0xf60fdcb3.
//
// Solidity: function cachedPrice() view returns(uint256)
func (_TokenPaymaster *TokenPaymasterCallerSession) CachedPrice() (*big.Int, error) {
	return _TokenPaymaster.Contract.CachedPrice(&_TokenPaymaster.CallOpts)
}

// CachedPriceTimestamp is a free data retrieval call binding the contract method 0xe1d8153c.
//
// Solidity: function cachedPriceTimestamp() view returns(uint48)
func (_TokenPaymaster *TokenPaymasterCaller) CachedPriceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenPaymaster.contract.Call(opts, &out, "cachedPriceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CachedPriceTimestamp is a free data retrieval call binding the contract method 0xe1d8153c.
//
// Solidity: function cachedPriceTimestamp() view returns(uint48)
func (_TokenPaymaster *TokenPaymasterSession) CachedPriceTimestamp() (*big.Int, error) {
	return _TokenPaymaster.Contract.CachedPriceTimestamp(&_TokenPaymaster.CallOpts)
}

// CachedPriceTimestamp is a free data retrieval call binding the contract method 0xe1d8153c.
//
// Solidity: function cachedPriceTimestamp() view returns(uint48)
func (_TokenPaymaster *TokenPaymasterCallerSession) CachedPriceTimestamp() (*big.Int, error) {
	return _TokenPaymaster.Contract.CachedPriceTimestamp(&_TokenPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TokenPaymaster *TokenPaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenPaymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TokenPaymaster *TokenPaymasterSession) EntryPoint() (common.Address, error) {
	return _TokenPaymaster.Contract.EntryPoint(&_TokenPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TokenPaymaster *TokenPaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _TokenPaymaster.Contract.EntryPoint(&_TokenPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TokenPaymaster *TokenPaymasterCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenPaymaster.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TokenPaymaster *TokenPaymasterSession) GetDeposit() (*big.Int, error) {
	return _TokenPaymaster.Contract.GetDeposit(&_TokenPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_TokenPaymaster *TokenPaymasterCallerSession) GetDeposit() (*big.Int, error) {
	return _TokenPaymaster.Contract.GetDeposit(&_TokenPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenPaymaster *TokenPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenPaymaster *TokenPaymasterSession) Owner() (common.Address, error) {
	return _TokenPaymaster.Contract.Owner(&_TokenPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenPaymaster *TokenPaymasterCallerSession) Owner() (common.Address, error) {
	return _TokenPaymaster.Contract.Owner(&_TokenPaymaster.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TokenPaymaster *TokenPaymasterCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenPaymaster.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TokenPaymaster *TokenPaymasterSession) Token() (common.Address, error) {
	return _TokenPaymaster.Contract.Token(&_TokenPaymaster.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TokenPaymaster *TokenPaymasterCallerSession) Token() (common.Address, error) {
	return _TokenPaymaster.Contract.Token(&_TokenPaymaster.CallOpts)
}

// TokenPaymasterConfig is a free data retrieval call binding the contract method 0xcb721cfd.
//
// Solidity: function tokenPaymasterConfig() view returns(uint256 priceMarkup, uint128 minEntryPointBalance, uint48 refundPostopCost, uint48 priceMaxAge)
func (_TokenPaymaster *TokenPaymasterCaller) TokenPaymasterConfig(opts *bind.CallOpts) (struct {
	PriceMarkup          *big.Int
	MinEntryPointBalance *big.Int
	RefundPostopCost     *big.Int
	PriceMaxAge          *big.Int
}, error) {
	var out []interface{}
	err := _TokenPaymaster.contract.Call(opts, &out, "tokenPaymasterConfig")

	outstruct := new(struct {
		PriceMarkup          *big.Int
		MinEntryPointBalance *big.Int
		RefundPostopCost     *big.Int
		PriceMaxAge          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PriceMarkup = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MinEntryPointBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RefundPostopCost = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PriceMaxAge = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenPaymasterConfig is a free data retrieval call binding the contract method 0xcb721cfd.
//
// Solidity: function tokenPaymasterConfig() view returns(uint256 priceMarkup, uint128 minEntryPointBalance, uint48 refundPostopCost, uint48 priceMaxAge)
func (_TokenPaymaster *TokenPaymasterSession) TokenPaymasterConfig() (struct {
	PriceMarkup          *big.Int
	MinEntryPointBalance *big.Int
	RefundPostopCost     *big.Int
	PriceMaxAge          *big.Int
}, error) {
	return _TokenPaymaster.Contract.TokenPaymasterConfig(&_TokenPaymaster.CallOpts)
}

// TokenPaymasterConfig is a free data retrieval call binding the contract method 0xcb721cfd.
//
// Solidity: function tokenPaymasterConfig() view returns(uint256 priceMarkup, uint128 minEntryPointBalance, uint48 refundPostopCost, uint48 priceMaxAge)
func (_TokenPaymaster *TokenPaymasterCallerSession) TokenPaymasterConfig() (struct {
	PriceMarkup          *big.Int
	MinEntryPointBalance *big.Int
	RefundPostopCost     *big.Int
	PriceMaxAge          *big.Int
}, error) {
	return _TokenPaymaster.Contract.TokenPaymasterConfig(&_TokenPaymaster.CallOpts)
}

// TokenToWei is a free data retrieval call binding the contract method 0xd7a23b3c.
//
// Solidity: function tokenToWei(uint256 amount, uint256 price) pure returns(uint256)
func (_TokenPaymaster *TokenPaymasterCaller) TokenToWei(opts *bind.CallOpts, amount *big.Int, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenPaymaster.contract.Call(opts, &out, "tokenToWei", amount, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToWei is a free data retrieval call binding the contract method 0xd7a23b3c.
//
// Solidity: function tokenToWei(uint256 amount, uint256 price) pure returns(uint256)
func (_TokenPaymaster *TokenPaymasterSession) TokenToWei(amount *big.Int, price *big.Int) (*big.Int, error) {
	return _TokenPaymaster.Contract.TokenToWei(&_TokenPaymaster.CallOpts, amount, price)
}

// TokenToWei is a free data retrieval call binding the contract method 0xd7a23b3c.
//
// Solidity: function tokenToWei(uint256 amount, uint256 price) pure returns(uint256)
func (_TokenPaymaster *TokenPaymasterCallerSession) TokenToWei(amount *big.Int, price *big.Int) (*big.Int, error) {
	return _TokenPaymaster.Contract.TokenToWei(&_TokenPaymaster.CallOpts, amount, price)
}

// Uniswap is a free data retrieval call binding the contract method 0x2681f7e4.
//
// Solidity: function uniswap() view returns(address)
func (_TokenPaymaster *TokenPaymasterCaller) Uniswap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenPaymaster.contract.Call(opts, &out, "uniswap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uniswap is a free data retrieval call binding the contract method 0x2681f7e4.
//
// Solidity: function uniswap() view returns(address)
func (_TokenPaymaster *TokenPaymasterSession) Uniswap() (common.Address, error) {
	return _TokenPaymaster.Contract.Uniswap(&_TokenPaymaster.CallOpts)
}

// Uniswap is a free data retrieval call binding the contract method 0x2681f7e4.
//
// Solidity: function uniswap() view returns(address)
func (_TokenPaymaster *TokenPaymasterCallerSession) Uniswap() (common.Address, error) {
	return _TokenPaymaster.Contract.Uniswap(&_TokenPaymaster.CallOpts)
}

// WeiToToken is a free data retrieval call binding the contract method 0x7c986aac.
//
// Solidity: function weiToToken(uint256 amount, uint256 price) pure returns(uint256)
func (_TokenPaymaster *TokenPaymasterCaller) WeiToToken(opts *bind.CallOpts, amount *big.Int, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenPaymaster.contract.Call(opts, &out, "weiToToken", amount, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiToToken is a free data retrieval call binding the contract method 0x7c986aac.
//
// Solidity: function weiToToken(uint256 amount, uint256 price) pure returns(uint256)
func (_TokenPaymaster *TokenPaymasterSession) WeiToToken(amount *big.Int, price *big.Int) (*big.Int, error) {
	return _TokenPaymaster.Contract.WeiToToken(&_TokenPaymaster.CallOpts, amount, price)
}

// WeiToToken is a free data retrieval call binding the contract method 0x7c986aac.
//
// Solidity: function weiToToken(uint256 amount, uint256 price) pure returns(uint256)
func (_TokenPaymaster *TokenPaymasterCallerSession) WeiToToken(amount *big.Int, price *big.Int) (*big.Int, error) {
	return _TokenPaymaster.Contract.WeiToToken(&_TokenPaymaster.CallOpts, amount, price)
}

// WrappedNative is a free data retrieval call binding the contract method 0xeb6d3a11.
//
// Solidity: function wrappedNative() view returns(address)
func (_TokenPaymaster *TokenPaymasterCaller) WrappedNative(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenPaymaster.contract.Call(opts, &out, "wrappedNative")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedNative is a free data retrieval call binding the contract method 0xeb6d3a11.
//
// Solidity: function wrappedNative() view returns(address)
func (_TokenPaymaster *TokenPaymasterSession) WrappedNative() (common.Address, error) {
	return _TokenPaymaster.Contract.WrappedNative(&_TokenPaymaster.CallOpts)
}

// WrappedNative is a free data retrieval call binding the contract method 0xeb6d3a11.
//
// Solidity: function wrappedNative() view returns(address)
func (_TokenPaymaster *TokenPaymasterCallerSession) WrappedNative() (common.Address, error) {
	return _TokenPaymaster.Contract.WrappedNative(&_TokenPaymaster.CallOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_TokenPaymaster *TokenPaymasterTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_TokenPaymaster *TokenPaymasterSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.AddStake(&_TokenPaymaster.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_TokenPaymaster *TokenPaymasterTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.AddStake(&_TokenPaymaster.TransactOpts, unstakeDelaySec)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TokenPaymaster *TokenPaymasterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TokenPaymaster *TokenPaymasterSession) Deposit() (*types.Transaction, error) {
	return _TokenPaymaster.Contract.Deposit(&_TokenPaymaster.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TokenPaymaster *TokenPaymasterTransactorSession) Deposit() (*types.Transaction, error) {
	return _TokenPaymaster.Contract.Deposit(&_TokenPaymaster.TransactOpts)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_TokenPaymaster *TokenPaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_TokenPaymaster *TokenPaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.PostOp(&_TokenPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_TokenPaymaster *TokenPaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.PostOp(&_TokenPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenPaymaster *TokenPaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenPaymaster *TokenPaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenPaymaster.Contract.RenounceOwnership(&_TokenPaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenPaymaster *TokenPaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenPaymaster.Contract.RenounceOwnership(&_TokenPaymaster.TransactOpts)
}

// SetTokenPaymasterConfig is a paid mutator transaction binding the contract method 0xf14d64ed.
//
// Solidity: function setTokenPaymasterConfig((uint256,uint128,uint48,uint48) _tokenPaymasterConfig) returns()
func (_TokenPaymaster *TokenPaymasterTransactor) SetTokenPaymasterConfig(opts *bind.TransactOpts, _tokenPaymasterConfig TokenPaymasterTokenPaymasterConfig) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "setTokenPaymasterConfig", _tokenPaymasterConfig)
}

// SetTokenPaymasterConfig is a paid mutator transaction binding the contract method 0xf14d64ed.
//
// Solidity: function setTokenPaymasterConfig((uint256,uint128,uint48,uint48) _tokenPaymasterConfig) returns()
func (_TokenPaymaster *TokenPaymasterSession) SetTokenPaymasterConfig(_tokenPaymasterConfig TokenPaymasterTokenPaymasterConfig) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.SetTokenPaymasterConfig(&_TokenPaymaster.TransactOpts, _tokenPaymasterConfig)
}

// SetTokenPaymasterConfig is a paid mutator transaction binding the contract method 0xf14d64ed.
//
// Solidity: function setTokenPaymasterConfig((uint256,uint128,uint48,uint48) _tokenPaymasterConfig) returns()
func (_TokenPaymaster *TokenPaymasterTransactorSession) SetTokenPaymasterConfig(_tokenPaymasterConfig TokenPaymasterTokenPaymasterConfig) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.SetTokenPaymasterConfig(&_TokenPaymaster.TransactOpts, _tokenPaymasterConfig)
}

// SetUniswapConfiguration is a paid mutator transaction binding the contract method 0xa0840fa7.
//
// Solidity: function setUniswapConfiguration((uint256,uint24,uint8) _uniswapHelperConfig) returns()
func (_TokenPaymaster *TokenPaymasterTransactor) SetUniswapConfiguration(opts *bind.TransactOpts, _uniswapHelperConfig UniswapHelperUniswapHelperConfig) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "setUniswapConfiguration", _uniswapHelperConfig)
}

// SetUniswapConfiguration is a paid mutator transaction binding the contract method 0xa0840fa7.
//
// Solidity: function setUniswapConfiguration((uint256,uint24,uint8) _uniswapHelperConfig) returns()
func (_TokenPaymaster *TokenPaymasterSession) SetUniswapConfiguration(_uniswapHelperConfig UniswapHelperUniswapHelperConfig) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.SetUniswapConfiguration(&_TokenPaymaster.TransactOpts, _uniswapHelperConfig)
}

// SetUniswapConfiguration is a paid mutator transaction binding the contract method 0xa0840fa7.
//
// Solidity: function setUniswapConfiguration((uint256,uint24,uint8) _uniswapHelperConfig) returns()
func (_TokenPaymaster *TokenPaymasterTransactorSession) SetUniswapConfiguration(_uniswapHelperConfig UniswapHelperUniswapHelperConfig) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.SetUniswapConfiguration(&_TokenPaymaster.TransactOpts, _uniswapHelperConfig)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenPaymaster *TokenPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenPaymaster *TokenPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.TransferOwnership(&_TokenPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenPaymaster *TokenPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.TransferOwnership(&_TokenPaymaster.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_TokenPaymaster *TokenPaymasterTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_TokenPaymaster *TokenPaymasterSession) UnlockStake() (*types.Transaction, error) {
	return _TokenPaymaster.Contract.UnlockStake(&_TokenPaymaster.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_TokenPaymaster *TokenPaymasterTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _TokenPaymaster.Contract.UnlockStake(&_TokenPaymaster.TransactOpts)
}

// UpdateCachedPrice is a paid mutator transaction binding the contract method 0x3ba9290f.
//
// Solidity: function updateCachedPrice(bool force) returns(uint256)
func (_TokenPaymaster *TokenPaymasterTransactor) UpdateCachedPrice(opts *bind.TransactOpts, force bool) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "updateCachedPrice", force)
}

// UpdateCachedPrice is a paid mutator transaction binding the contract method 0x3ba9290f.
//
// Solidity: function updateCachedPrice(bool force) returns(uint256)
func (_TokenPaymaster *TokenPaymasterSession) UpdateCachedPrice(force bool) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.UpdateCachedPrice(&_TokenPaymaster.TransactOpts, force)
}

// UpdateCachedPrice is a paid mutator transaction binding the contract method 0x3ba9290f.
//
// Solidity: function updateCachedPrice(bool force) returns(uint256)
func (_TokenPaymaster *TokenPaymasterTransactorSession) UpdateCachedPrice(force bool) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.UpdateCachedPrice(&_TokenPaymaster.TransactOpts, force)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_TokenPaymaster *TokenPaymasterTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_TokenPaymaster *TokenPaymasterSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.ValidatePaymasterUserOp(&_TokenPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_TokenPaymaster *TokenPaymasterTransactorSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.ValidatePaymasterUserOp(&_TokenPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address recipient, uint256 amount) returns()
func (_TokenPaymaster *TokenPaymasterTransactor) WithdrawEth(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "withdrawEth", recipient, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address recipient, uint256 amount) returns()
func (_TokenPaymaster *TokenPaymasterSession) WithdrawEth(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.WithdrawEth(&_TokenPaymaster.TransactOpts, recipient, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address recipient, uint256 amount) returns()
func (_TokenPaymaster *TokenPaymasterTransactorSession) WithdrawEth(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.WithdrawEth(&_TokenPaymaster.TransactOpts, recipient, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_TokenPaymaster *TokenPaymasterTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_TokenPaymaster *TokenPaymasterSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.WithdrawStake(&_TokenPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_TokenPaymaster *TokenPaymasterTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.WithdrawStake(&_TokenPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_TokenPaymaster *TokenPaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_TokenPaymaster *TokenPaymasterSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.WithdrawTo(&_TokenPaymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_TokenPaymaster *TokenPaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.WithdrawTo(&_TokenPaymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address to, uint256 amount) returns()
func (_TokenPaymaster *TokenPaymasterTransactor) WithdrawToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.contract.Transact(opts, "withdrawToken", to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address to, uint256 amount) returns()
func (_TokenPaymaster *TokenPaymasterSession) WithdrawToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.WithdrawToken(&_TokenPaymaster.TransactOpts, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address to, uint256 amount) returns()
func (_TokenPaymaster *TokenPaymasterTransactorSession) WithdrawToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenPaymaster.Contract.WithdrawToken(&_TokenPaymaster.TransactOpts, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TokenPaymaster *TokenPaymasterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPaymaster.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TokenPaymaster *TokenPaymasterSession) Receive() (*types.Transaction, error) {
	return _TokenPaymaster.Contract.Receive(&_TokenPaymaster.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TokenPaymaster *TokenPaymasterTransactorSession) Receive() (*types.Transaction, error) {
	return _TokenPaymaster.Contract.Receive(&_TokenPaymaster.TransactOpts)
}

// TokenPaymasterConfigUpdatedIterator is returned from FilterConfigUpdated and is used to iterate over the raw logs and unpacked data for ConfigUpdated events raised by the TokenPaymaster contract.
type TokenPaymasterConfigUpdatedIterator struct {
	Event *TokenPaymasterConfigUpdated // Event containing the contract specifics and raw log

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
func (it *TokenPaymasterConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPaymasterConfigUpdated)
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
		it.Event = new(TokenPaymasterConfigUpdated)
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
func (it *TokenPaymasterConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPaymasterConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPaymasterConfigUpdated represents a ConfigUpdated event raised by the TokenPaymaster contract.
type TokenPaymasterConfigUpdated struct {
	TokenPaymasterConfig TokenPaymasterTokenPaymasterConfig
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterConfigUpdated is a free log retrieval operation binding the contract event 0xcd938817f1c47094d43be3d07e8c67e11766db2e11a2b4376e7ee937b15793a2.
//
// Solidity: event ConfigUpdated((uint256,uint128,uint48,uint48) tokenPaymasterConfig)
func (_TokenPaymaster *TokenPaymasterFilterer) FilterConfigUpdated(opts *bind.FilterOpts) (*TokenPaymasterConfigUpdatedIterator, error) {

	logs, sub, err := _TokenPaymaster.contract.FilterLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenPaymasterConfigUpdatedIterator{contract: _TokenPaymaster.contract, event: "ConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchConfigUpdated is a free log subscription operation binding the contract event 0xcd938817f1c47094d43be3d07e8c67e11766db2e11a2b4376e7ee937b15793a2.
//
// Solidity: event ConfigUpdated((uint256,uint128,uint48,uint48) tokenPaymasterConfig)
func (_TokenPaymaster *TokenPaymasterFilterer) WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *TokenPaymasterConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _TokenPaymaster.contract.WatchLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPaymasterConfigUpdated)
				if err := _TokenPaymaster.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
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

// ParseConfigUpdated is a log parse operation binding the contract event 0xcd938817f1c47094d43be3d07e8c67e11766db2e11a2b4376e7ee937b15793a2.
//
// Solidity: event ConfigUpdated((uint256,uint128,uint48,uint48) tokenPaymasterConfig)
func (_TokenPaymaster *TokenPaymasterFilterer) ParseConfigUpdated(log types.Log) (*TokenPaymasterConfigUpdated, error) {
	event := new(TokenPaymasterConfigUpdated)
	if err := _TokenPaymaster.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenPaymaster contract.
type TokenPaymasterOwnershipTransferredIterator struct {
	Event *TokenPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPaymasterOwnershipTransferred)
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
		it.Event = new(TokenPaymasterOwnershipTransferred)
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
func (it *TokenPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the TokenPaymaster contract.
type TokenPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenPaymaster *TokenPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenPaymasterOwnershipTransferredIterator{contract: _TokenPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenPaymaster *TokenPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPaymasterOwnershipTransferred)
				if err := _TokenPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenPaymaster *TokenPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*TokenPaymasterOwnershipTransferred, error) {
	event := new(TokenPaymasterOwnershipTransferred)
	if err := _TokenPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenPaymasterReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the TokenPaymaster contract.
type TokenPaymasterReceivedIterator struct {
	Event *TokenPaymasterReceived // Event containing the contract specifics and raw log

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
func (it *TokenPaymasterReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPaymasterReceived)
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
		it.Event = new(TokenPaymasterReceived)
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
func (it *TokenPaymasterReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPaymasterReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPaymasterReceived represents a Received event raised by the TokenPaymaster contract.
type TokenPaymasterReceived struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed sender, uint256 value)
func (_TokenPaymaster *TokenPaymasterFilterer) FilterReceived(opts *bind.FilterOpts, sender []common.Address) (*TokenPaymasterReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TokenPaymaster.contract.FilterLogs(opts, "Received", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenPaymasterReceivedIterator{contract: _TokenPaymaster.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed sender, uint256 value)
func (_TokenPaymaster *TokenPaymasterFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *TokenPaymasterReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TokenPaymaster.contract.WatchLogs(opts, "Received", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPaymasterReceived)
				if err := _TokenPaymaster.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed sender, uint256 value)
func (_TokenPaymaster *TokenPaymasterFilterer) ParseReceived(log types.Log) (*TokenPaymasterReceived, error) {
	event := new(TokenPaymasterReceived)
	if err := _TokenPaymaster.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenPaymasterTokenPriceUpdatedIterator is returned from FilterTokenPriceUpdated and is used to iterate over the raw logs and unpacked data for TokenPriceUpdated events raised by the TokenPaymaster contract.
type TokenPaymasterTokenPriceUpdatedIterator struct {
	Event *TokenPaymasterTokenPriceUpdated // Event containing the contract specifics and raw log

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
func (it *TokenPaymasterTokenPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPaymasterTokenPriceUpdated)
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
		it.Event = new(TokenPaymasterTokenPriceUpdated)
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
func (it *TokenPaymasterTokenPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPaymasterTokenPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPaymasterTokenPriceUpdated represents a TokenPriceUpdated event raised by the TokenPaymaster contract.
type TokenPaymasterTokenPriceUpdated struct {
	CurrentPrice         *big.Int
	PreviousPrice        *big.Int
	CachedPriceTimestamp *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTokenPriceUpdated is a free log retrieval operation binding the contract event 0x00d4fe314618b73a96886b87817a53a5ed51433b0234c85a5e9dafe2cb7b8842.
//
// Solidity: event TokenPriceUpdated(uint256 currentPrice, uint256 previousPrice, uint256 cachedPriceTimestamp)
func (_TokenPaymaster *TokenPaymasterFilterer) FilterTokenPriceUpdated(opts *bind.FilterOpts) (*TokenPaymasterTokenPriceUpdatedIterator, error) {

	logs, sub, err := _TokenPaymaster.contract.FilterLogs(opts, "TokenPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenPaymasterTokenPriceUpdatedIterator{contract: _TokenPaymaster.contract, event: "TokenPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenPriceUpdated is a free log subscription operation binding the contract event 0x00d4fe314618b73a96886b87817a53a5ed51433b0234c85a5e9dafe2cb7b8842.
//
// Solidity: event TokenPriceUpdated(uint256 currentPrice, uint256 previousPrice, uint256 cachedPriceTimestamp)
func (_TokenPaymaster *TokenPaymasterFilterer) WatchTokenPriceUpdated(opts *bind.WatchOpts, sink chan<- *TokenPaymasterTokenPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _TokenPaymaster.contract.WatchLogs(opts, "TokenPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPaymasterTokenPriceUpdated)
				if err := _TokenPaymaster.contract.UnpackLog(event, "TokenPriceUpdated", log); err != nil {
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

// ParseTokenPriceUpdated is a log parse operation binding the contract event 0x00d4fe314618b73a96886b87817a53a5ed51433b0234c85a5e9dafe2cb7b8842.
//
// Solidity: event TokenPriceUpdated(uint256 currentPrice, uint256 previousPrice, uint256 cachedPriceTimestamp)
func (_TokenPaymaster *TokenPaymasterFilterer) ParseTokenPriceUpdated(log types.Log) (*TokenPaymasterTokenPriceUpdated, error) {
	event := new(TokenPaymasterTokenPriceUpdated)
	if err := _TokenPaymaster.contract.UnpackLog(event, "TokenPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenPaymasterUniswapRevertedIterator is returned from FilterUniswapReverted and is used to iterate over the raw logs and unpacked data for UniswapReverted events raised by the TokenPaymaster contract.
type TokenPaymasterUniswapRevertedIterator struct {
	Event *TokenPaymasterUniswapReverted // Event containing the contract specifics and raw log

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
func (it *TokenPaymasterUniswapRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPaymasterUniswapReverted)
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
		it.Event = new(TokenPaymasterUniswapReverted)
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
func (it *TokenPaymasterUniswapRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPaymasterUniswapRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPaymasterUniswapReverted represents a UniswapReverted event raised by the TokenPaymaster contract.
type TokenPaymasterUniswapReverted struct {
	TokenIn      common.Address
	TokenOut     common.Address
	AmountIn     *big.Int
	AmountOutMin *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUniswapReverted is a free log retrieval operation binding the contract event 0xf7edd4c6ec425decf715a8b8eaa3b65d3d86e31ad0ff750aa60fa834190f515f.
//
// Solidity: event UniswapReverted(address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOutMin)
func (_TokenPaymaster *TokenPaymasterFilterer) FilterUniswapReverted(opts *bind.FilterOpts) (*TokenPaymasterUniswapRevertedIterator, error) {

	logs, sub, err := _TokenPaymaster.contract.FilterLogs(opts, "UniswapReverted")
	if err != nil {
		return nil, err
	}
	return &TokenPaymasterUniswapRevertedIterator{contract: _TokenPaymaster.contract, event: "UniswapReverted", logs: logs, sub: sub}, nil
}

// WatchUniswapReverted is a free log subscription operation binding the contract event 0xf7edd4c6ec425decf715a8b8eaa3b65d3d86e31ad0ff750aa60fa834190f515f.
//
// Solidity: event UniswapReverted(address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOutMin)
func (_TokenPaymaster *TokenPaymasterFilterer) WatchUniswapReverted(opts *bind.WatchOpts, sink chan<- *TokenPaymasterUniswapReverted) (event.Subscription, error) {

	logs, sub, err := _TokenPaymaster.contract.WatchLogs(opts, "UniswapReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPaymasterUniswapReverted)
				if err := _TokenPaymaster.contract.UnpackLog(event, "UniswapReverted", log); err != nil {
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

// ParseUniswapReverted is a log parse operation binding the contract event 0xf7edd4c6ec425decf715a8b8eaa3b65d3d86e31ad0ff750aa60fa834190f515f.
//
// Solidity: event UniswapReverted(address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOutMin)
func (_TokenPaymaster *TokenPaymasterFilterer) ParseUniswapReverted(log types.Log) (*TokenPaymasterUniswapReverted, error) {
	event := new(TokenPaymasterUniswapReverted)
	if err := _TokenPaymaster.contract.UnpackLog(event, "UniswapReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenPaymasterUserOperationSponsoredIterator is returned from FilterUserOperationSponsored and is used to iterate over the raw logs and unpacked data for UserOperationSponsored events raised by the TokenPaymaster contract.
type TokenPaymasterUserOperationSponsoredIterator struct {
	Event *TokenPaymasterUserOperationSponsored // Event containing the contract specifics and raw log

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
func (it *TokenPaymasterUserOperationSponsoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPaymasterUserOperationSponsored)
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
		it.Event = new(TokenPaymasterUserOperationSponsored)
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
func (it *TokenPaymasterUserOperationSponsoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPaymasterUserOperationSponsoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPaymasterUserOperationSponsored represents a UserOperationSponsored event raised by the TokenPaymaster contract.
type TokenPaymasterUserOperationSponsored struct {
	User                       common.Address
	ActualTokenCharge          *big.Int
	ActualGasCost              *big.Int
	ActualTokenPriceWithMarkup *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterUserOperationSponsored is a free log retrieval operation binding the contract event 0x46caa0511cf037f06f57a0bf273a2ff04229f5b12fb04675234a6cbe2e7f1a89.
//
// Solidity: event UserOperationSponsored(address indexed user, uint256 actualTokenCharge, uint256 actualGasCost, uint256 actualTokenPriceWithMarkup)
func (_TokenPaymaster *TokenPaymasterFilterer) FilterUserOperationSponsored(opts *bind.FilterOpts, user []common.Address) (*TokenPaymasterUserOperationSponsoredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TokenPaymaster.contract.FilterLogs(opts, "UserOperationSponsored", userRule)
	if err != nil {
		return nil, err
	}
	return &TokenPaymasterUserOperationSponsoredIterator{contract: _TokenPaymaster.contract, event: "UserOperationSponsored", logs: logs, sub: sub}, nil
}

// WatchUserOperationSponsored is a free log subscription operation binding the contract event 0x46caa0511cf037f06f57a0bf273a2ff04229f5b12fb04675234a6cbe2e7f1a89.
//
// Solidity: event UserOperationSponsored(address indexed user, uint256 actualTokenCharge, uint256 actualGasCost, uint256 actualTokenPriceWithMarkup)
func (_TokenPaymaster *TokenPaymasterFilterer) WatchUserOperationSponsored(opts *bind.WatchOpts, sink chan<- *TokenPaymasterUserOperationSponsored, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TokenPaymaster.contract.WatchLogs(opts, "UserOperationSponsored", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPaymasterUserOperationSponsored)
				if err := _TokenPaymaster.contract.UnpackLog(event, "UserOperationSponsored", log); err != nil {
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

// ParseUserOperationSponsored is a log parse operation binding the contract event 0x46caa0511cf037f06f57a0bf273a2ff04229f5b12fb04675234a6cbe2e7f1a89.
//
// Solidity: event UserOperationSponsored(address indexed user, uint256 actualTokenCharge, uint256 actualGasCost, uint256 actualTokenPriceWithMarkup)
func (_TokenPaymaster *TokenPaymasterFilterer) ParseUserOperationSponsored(log types.Log) (*TokenPaymasterUserOperationSponsored, error) {
	event := new(TokenPaymasterUserOperationSponsored)
	if err := _TokenPaymaster.contract.UnpackLog(event, "UserOperationSponsored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
