// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package legacytokenpaymaster

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

// LegacyTokenPaymasterMetaData contains all meta data concerning the LegacyTokenPaymaster contract.
var LegacyTokenPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountFactory\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COST_OF_POST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"theFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c0604090808252346105535761181c803803809161001e8285610585565b83398101606082820312610553578151916001600160a01b038084168403610553576020828101519094906001600160401b039081811161055357840190601f9486868401121561055357825182811161056f57601f19938a51986100898b878b860116018b610585565b828a528a8383010111610553578a92918a9160005b828110610558575050906000918a0101520151948486169586810361055357331561053b576000543360018060a01b03198216176000558a51963391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a36301ffc9a760e01b865285898160049963122a0e9b60e31b8b8301526024998a915afa908115610530576000916104f3575b50156104b1576080528651968288116103e3578654976001988981811c911680156104a7575b8b8210146103c557838111610461575b50808a848211600114610402576000916103f7575b50600019600383901b1c191690891b1787555b80519283116103e35760059384548981811c911680156103d9575b8b8210146103c5579081848695949311610372575b508a9284116001146103125750600092610307575b5050600019600383901b1c191690861b1790555b60a05230156102f257600354918383018093116102df5750506003553060005280825282600020818154019055825190815260007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef833093a330600052600281528160002033600052815260001980836000205582519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92533923092a35161127390816105a9823960805181818161037501528181610423015281816104d50152818161054b015281816105b601528181610cf201528181610f3d0152611166015260a05181818161062b01526109b40152f35b601190634e487b7160e01b600052526000fd5b90600085519163ec442f0560e01b8352820152fd5b0151905038806101d5565b89949291921691856000528a6000209260005b8c82821061035c5750508411610343575b505050811b0190556101e9565b015160001960f88460031b161c19169055388080610336565b8385015186558c97909501949384019301610325565b9091929350856000528a60002084808701881c8201928d88106103bc575b9187968d929695949301891c01915b8281106103ad5750506101c0565b600081558796508c910161039f565b92508192610390565b8760228a634e487b7160e01b600052526000fd5b90607f16906101ab565b85604188634e487b7160e01b600052526000fd5b90508201513861017d565b868b9316908a6000528c600020918d6000905b82821061044a5750508311610431575b5050811b018755610190565b84015160001960f88460031b161c191690553880610425565b8388015185558e969094019392830192018e610415565b886000528a6000208480840160051c8201928d851061049e575b0160051c01908a905b828110610492575050610168565b60008155018a90610484565b9250819261047b565b90607f1690610158565b895162461bcd60e51b81528088018a9052601e818801527f49456e747279506f696e7420696e74657266616365206d69736d6174636800006044820152606490fd5b8a81813d8311610529575b6105088183610585565b810103126105255751908115158203610522575038610132565b80fd5b5080fd5b503d6104fe565b8b513d6000823e3d90fd5b8951631e4fbdf760e01b815260006004820152602490fd5b600080fd5b8181018401518c82018501528d95508c930161009e565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761056f5760405256fe60406080815260048036101561001457600080fd5b600091823560e01c80630396cb6014610f1557806306fdde0314610e1e578063095ea7b314610d6e57806318160ddd14610d4f57838163205c287814610cc55750806323b872dd14610bcc578063313ce56714610bb057806352b7512c1461091d57806370a08231146108da578063715018a614610880578063796d4371146108635780637c627b211461079f5780638da5cb5b1461077757806395d89b411461065a5780639f5ca22114610616578063a9059cbb146105e5578063b0d691fe146105a157838163bb9fe6bf1461052e578163c23a5cea146104a757508063c399ec88146103f657838163d0e30db01461036557508063dd62ed3e14610313578063f0dda65c146102795763f2fde38b1461012e57600080fd5b346102755760208060031936011261027157610148610fe0565b90610151611058565b84546001600160a01b0390811692301561025a578315610243578190306000526002845285600020856000528452600086812055855194600086527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259586863092a36101bb611058565b1694851561022c57508490600054826bffffffffffffffffffffffff60a01b821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a330600052600281528260002084600052815260001992838160002055519283523092a380f35b8451631e4fbdf760e01b8152908101879052602490fd5b8451634a1406b160e11b8152600081880152602490fd5b845163e602df0560e01b8152600081880152602490fd5b8380fd5b8280fd5b5034610275578060031936011261027557610292610fe0565b906024359161029f611058565b6001600160a01b03169283156102fe57506020827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef926102e360009560035461121a565b6003558585526001835280852082815401905551908152a380f35b84602492519163ec442f0560e01b8352820152fd5b838234610361578060031936011261036157602091610330610fe0565b82610339610ff6565b6001600160a01b03928316845260028652922091166000908152908352819020549051908152f35b5080fd5b808484826003193601126103f2577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156103ed578390602483518095819363b760faf960e01b8352309083015234905af19081156103e457506103d15750f35b6103da9061100c565b6103e15780f35b80fd5b513d84823e3d90fd5b505050fd5b5050fd5b503461027557826003193601126102755780516370a0823160e01b815230928101929092526020826024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa91821561049d578392610465575b6020838351908152f35b9091506020813d602011610495575b8161048160209383611036565b81010312610275576020925051903861045b565b3d9150610474565b81513d85823e3d90fd5b808484346103f25760203660031901126103f2576104c3610fe0565b6104cb611058565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116803b1561052a578592836024928651978895869463611d2e7560e11b865216908401525af19081156103e457506103d15750f35b8580fd5b808484346103f257826003193601126103f257610549611058565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156103ed57815163bb9fe6bf60e01b81529284918491829084905af19081156103e457506103d15750f35b838234610361578160031936011261036157517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b83823461036157806003193601126103615760209061060f610605610fe0565b6024359033611084565b5160018152f35b838234610361578160031936011261036157517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5082346103e157806003193601126103e157815191816005549260018460011c916001861695861561076d575b602096878510811461075a578899509688969785829a5291826000146107335750506001146106d7575b5050506106d392916106c4910385611036565b51928284938452830190610fa0565b0390f35b9190869350600583527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db05b82841061071b57505050820101816106c46106d36106b1565b8054848a018601528895508794909301928101610702565b60ff19168782015293151560051b860190930193508492506106c491506106d390506106b1565b634e487b7160e01b835260228a52602483fd5b92607f1692610687565b838234610361578160031936011261036157905490516001600160a01b039091168152602090f35b8284346103e15760803660031901126103e1576003823510156103e1576024359167ffffffffffffffff80841161027557366023850112156102755783820135908111610275578301923660248501116102755760208160643595610802611164565b031261027557602401359060018060a01b03821680920361085e57613a98938085029485040361084b5750606461083f610848939460443561121a565b04903090611084565b80f35b634e487b7160e01b835260119052602482fd5b600080fd5b83823461036157816003193601126103615760209051613a988152f35b83346103e157806003193601126103e157610899611058565b600080546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b838234610361576020366003190112610361576020906109166108fb610fe0565b6001600160a01b031660009081526001602052604090205490565b9051908152f35b50919034610361576003199260603685011261027557813567ffffffffffffffff94858211610bac57610120828501918336030112610bac576044610960611164565b606481350461097260e48501846111d3565b603495919511610afb57613a9860248096013560801c1115610b5c57820161099a81856111d3565b159050610aff576109ab90846111d3565b601411610afb577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316903560601c03610aac576109f26108fb84611206565b10610a5b5750610a0190611206565b9282519360018060a01b03166020850152602084528284019584871090871117610a485750508381528352610a396080820182610fa0565b916060820152603f1991030190f35b604190634e487b7160e01b600052526000fd5b847f546f6b656e5061796d61737465723a206e6f2062616c616e636520287072652d608492602786602089519562461bcd60e51b8752860152840152820152666372656174652960c81b6064820152fd5b845162461bcd60e51b81526020818801526025818601527f546f6b656e5061796d61737465723a2077726f6e67206163636f756e74206661818401526463746f727960d81b6064820152608490fd5b8780fd5b50610b0c6108fb84611206565b10610b1b5750610a0190611206565b847f546f6b656e5061796d61737465723a206e6f2062616c616e6365000000000000606492601a86602089519562461bcd60e51b8752860152840152820152fd5b855162461bcd60e51b81526020818901526026818701527f546f6b656e5061796d61737465723a2067617320746f6f206c6f7720666f722081850152650706f73744f760d41b6064820152608490fd5b8480fd5b8382346103615781600319360112610361576020905160128152f35b5091903461036157606036600319011261036157610be8610fe0565b610bf0610ff6565b91604435938560018060a01b03841691828152600260205220336000526020528560002054916000198303610c2e575b60208761060f888888611084565b858310610c99578115610c82573315610c6b575060009081526002602090815286822033835281529086902091859003909155829061060f610c20565b6024906000885191634a1406b160e11b8352820152fd5b602490600088519163e602df0560e01b8352820152fd5b8651637dc7a0d960e11b8152339181019182526020820193909352604081018690528291506060010390fd5b808484346103f257806003193601126103f257610ce0610fe0565b610ce8611058565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116803b1561052a578592836044928651978895869463040b850f60e31b8652169084015260243560248401525af19081156103e457506103d15750f35b8382346103615781600319360112610361576020906003549051908152f35b5082346103e157816003193601126103e15750610d89610fe0565b602435903315610e07576001600160a01b0316908115610df0576020935033600052600284528260002082600052845280836000205582519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925843392a35160018152f35b8251634a1406b160e11b8152600081860152602490fd5b825163e602df0560e01b8152600081860152602490fd5b5082346103e157806003193601126103e1578151918184549260018460011c9160018616958615610f0b575b602096878510811461075a579087899a92868b999a9b529182600014610ee1575050600114610e86575b85886106d3896106c4848a0385611036565b815286935091907f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5b828410610ec957505050820101816106c46106d388610e74565b8054848a018601528895508794909301928101610eaf565b60ff19168882015294151560051b870190940194508593506106c492506106d39150899050610e74565b92607f1692610e4a565b5060203660031901126102755782823563ffffffff811680910361036157610f3b611058565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031693843b156102755760249084519586938492621cb65b60e51b845283015234905af19081156103e45750610f97575080f35b6108489061100c565b919082519283825260005b848110610fcc575050826000602080949584010152601f8019910116010190565b602081830181015184830182015201610fab565b600435906001600160a01b038216820361085e57565b602435906001600160a01b038216820361085e57565b67ffffffffffffffff811161102057604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff82111761102057604052565b6000546001600160a01b0316330361106c57565b60405163118cdaa760e01b8152336004820152602490fd5b916001600160a01b0380841692831561114b5716928315611132576000908382526001602052604082205490838210611100575091604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef95876020965260018652038282205586815220818154019055604051908152a3565b60405163391434e360e21b81526001600160a01b03919091166004820152602481019190915260448101839052606490fd5b60405163ec442f0560e01b815260006004820152602490fd5b604051634b637e8f60e11b815260006004820152602490fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361119657565b60405162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08115b9d1c9e541bda5b9d605a1b6044820152606490fd5b903590601e198136030182121561085e570180359067ffffffffffffffff821161085e5760200191813603831361085e57565b356001600160a01b038116810361085e5790565b9190820180921161122757565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220ee697a3fad7ae8d54c8d89bffa53ca4e698add9179d5a14cbf3c5687ebd9b22a64736f6c63430008190033",
}

// LegacyTokenPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use LegacyTokenPaymasterMetaData.ABI instead.
var LegacyTokenPaymasterABI = LegacyTokenPaymasterMetaData.ABI

// LegacyTokenPaymasterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LegacyTokenPaymasterMetaData.Bin instead.
var LegacyTokenPaymasterBin = LegacyTokenPaymasterMetaData.Bin

// DeployLegacyTokenPaymaster deploys a new Ethereum contract, binding an instance of LegacyTokenPaymaster to it.
func DeployLegacyTokenPaymaster(auth *bind.TransactOpts, backend bind.ContractBackend, accountFactory common.Address, _symbol string, _entryPoint common.Address) (common.Address, *types.Transaction, *LegacyTokenPaymaster, error) {
	parsed, err := LegacyTokenPaymasterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LegacyTokenPaymasterBin), backend, accountFactory, _symbol, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LegacyTokenPaymaster{LegacyTokenPaymasterCaller: LegacyTokenPaymasterCaller{contract: contract}, LegacyTokenPaymasterTransactor: LegacyTokenPaymasterTransactor{contract: contract}, LegacyTokenPaymasterFilterer: LegacyTokenPaymasterFilterer{contract: contract}}, nil
}

// LegacyTokenPaymaster is an auto generated Go binding around an Ethereum contract.
type LegacyTokenPaymaster struct {
	LegacyTokenPaymasterCaller     // Read-only binding to the contract
	LegacyTokenPaymasterTransactor // Write-only binding to the contract
	LegacyTokenPaymasterFilterer   // Log filterer for contract events
}

// LegacyTokenPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type LegacyTokenPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyTokenPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LegacyTokenPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyTokenPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LegacyTokenPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyTokenPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LegacyTokenPaymasterSession struct {
	Contract     *LegacyTokenPaymaster // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LegacyTokenPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LegacyTokenPaymasterCallerSession struct {
	Contract *LegacyTokenPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// LegacyTokenPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LegacyTokenPaymasterTransactorSession struct {
	Contract     *LegacyTokenPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// LegacyTokenPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type LegacyTokenPaymasterRaw struct {
	Contract *LegacyTokenPaymaster // Generic contract binding to access the raw methods on
}

// LegacyTokenPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LegacyTokenPaymasterCallerRaw struct {
	Contract *LegacyTokenPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// LegacyTokenPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LegacyTokenPaymasterTransactorRaw struct {
	Contract *LegacyTokenPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLegacyTokenPaymaster creates a new instance of LegacyTokenPaymaster, bound to a specific deployed contract.
func NewLegacyTokenPaymaster(address common.Address, backend bind.ContractBackend) (*LegacyTokenPaymaster, error) {
	contract, err := bindLegacyTokenPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LegacyTokenPaymaster{LegacyTokenPaymasterCaller: LegacyTokenPaymasterCaller{contract: contract}, LegacyTokenPaymasterTransactor: LegacyTokenPaymasterTransactor{contract: contract}, LegacyTokenPaymasterFilterer: LegacyTokenPaymasterFilterer{contract: contract}}, nil
}

// NewLegacyTokenPaymasterCaller creates a new read-only instance of LegacyTokenPaymaster, bound to a specific deployed contract.
func NewLegacyTokenPaymasterCaller(address common.Address, caller bind.ContractCaller) (*LegacyTokenPaymasterCaller, error) {
	contract, err := bindLegacyTokenPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyTokenPaymasterCaller{contract: contract}, nil
}

// NewLegacyTokenPaymasterTransactor creates a new write-only instance of LegacyTokenPaymaster, bound to a specific deployed contract.
func NewLegacyTokenPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*LegacyTokenPaymasterTransactor, error) {
	contract, err := bindLegacyTokenPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyTokenPaymasterTransactor{contract: contract}, nil
}

// NewLegacyTokenPaymasterFilterer creates a new log filterer instance of LegacyTokenPaymaster, bound to a specific deployed contract.
func NewLegacyTokenPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*LegacyTokenPaymasterFilterer, error) {
	contract, err := bindLegacyTokenPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacyTokenPaymasterFilterer{contract: contract}, nil
}

// bindLegacyTokenPaymaster binds a generic wrapper to an already deployed contract.
func bindLegacyTokenPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LegacyTokenPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyTokenPaymaster *LegacyTokenPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyTokenPaymaster.Contract.LegacyTokenPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyTokenPaymaster *LegacyTokenPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.LegacyTokenPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyTokenPaymaster *LegacyTokenPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.LegacyTokenPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyTokenPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.contract.Transact(opts, method, params...)
}

// COSTOFPOST is a free data retrieval call binding the contract method 0x796d4371.
//
// Solidity: function COST_OF_POST() view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCaller) COSTOFPOST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LegacyTokenPaymaster.contract.Call(opts, &out, "COST_OF_POST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COSTOFPOST is a free data retrieval call binding the contract method 0x796d4371.
//
// Solidity: function COST_OF_POST() view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) COSTOFPOST() (*big.Int, error) {
	return _LegacyTokenPaymaster.Contract.COSTOFPOST(&_LegacyTokenPaymaster.CallOpts)
}

// COSTOFPOST is a free data retrieval call binding the contract method 0x796d4371.
//
// Solidity: function COST_OF_POST() view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCallerSession) COSTOFPOST() (*big.Int, error) {
	return _LegacyTokenPaymaster.Contract.COSTOFPOST(&_LegacyTokenPaymaster.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LegacyTokenPaymaster.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LegacyTokenPaymaster.Contract.Allowance(&_LegacyTokenPaymaster.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LegacyTokenPaymaster.Contract.Allowance(&_LegacyTokenPaymaster.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LegacyTokenPaymaster.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LegacyTokenPaymaster.Contract.BalanceOf(&_LegacyTokenPaymaster.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LegacyTokenPaymaster.Contract.BalanceOf(&_LegacyTokenPaymaster.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LegacyTokenPaymaster.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) Decimals() (uint8, error) {
	return _LegacyTokenPaymaster.Contract.Decimals(&_LegacyTokenPaymaster.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCallerSession) Decimals() (uint8, error) {
	return _LegacyTokenPaymaster.Contract.Decimals(&_LegacyTokenPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyTokenPaymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) EntryPoint() (common.Address, error) {
	return _LegacyTokenPaymaster.Contract.EntryPoint(&_LegacyTokenPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _LegacyTokenPaymaster.Contract.EntryPoint(&_LegacyTokenPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LegacyTokenPaymaster.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) GetDeposit() (*big.Int, error) {
	return _LegacyTokenPaymaster.Contract.GetDeposit(&_LegacyTokenPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCallerSession) GetDeposit() (*big.Int, error) {
	return _LegacyTokenPaymaster.Contract.GetDeposit(&_LegacyTokenPaymaster.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LegacyTokenPaymaster.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) Name() (string, error) {
	return _LegacyTokenPaymaster.Contract.Name(&_LegacyTokenPaymaster.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCallerSession) Name() (string, error) {
	return _LegacyTokenPaymaster.Contract.Name(&_LegacyTokenPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyTokenPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) Owner() (common.Address, error) {
	return _LegacyTokenPaymaster.Contract.Owner(&_LegacyTokenPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCallerSession) Owner() (common.Address, error) {
	return _LegacyTokenPaymaster.Contract.Owner(&_LegacyTokenPaymaster.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LegacyTokenPaymaster.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) Symbol() (string, error) {
	return _LegacyTokenPaymaster.Contract.Symbol(&_LegacyTokenPaymaster.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCallerSession) Symbol() (string, error) {
	return _LegacyTokenPaymaster.Contract.Symbol(&_LegacyTokenPaymaster.CallOpts)
}

// TheFactory is a free data retrieval call binding the contract method 0x9f5ca221.
//
// Solidity: function theFactory() view returns(address)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCaller) TheFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyTokenPaymaster.contract.Call(opts, &out, "theFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TheFactory is a free data retrieval call binding the contract method 0x9f5ca221.
//
// Solidity: function theFactory() view returns(address)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) TheFactory() (common.Address, error) {
	return _LegacyTokenPaymaster.Contract.TheFactory(&_LegacyTokenPaymaster.CallOpts)
}

// TheFactory is a free data retrieval call binding the contract method 0x9f5ca221.
//
// Solidity: function theFactory() view returns(address)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCallerSession) TheFactory() (common.Address, error) {
	return _LegacyTokenPaymaster.Contract.TheFactory(&_LegacyTokenPaymaster.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LegacyTokenPaymaster.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) TotalSupply() (*big.Int, error) {
	return _LegacyTokenPaymaster.Contract.TotalSupply(&_LegacyTokenPaymaster.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterCallerSession) TotalSupply() (*big.Int, error) {
	return _LegacyTokenPaymaster.Contract.TotalSupply(&_LegacyTokenPaymaster.CallOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.AddStake(&_LegacyTokenPaymaster.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.AddStake(&_LegacyTokenPaymaster.TransactOpts, unstakeDelaySec)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.Approve(&_LegacyTokenPaymaster.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.Approve(&_LegacyTokenPaymaster.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) Deposit() (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.Deposit(&_LegacyTokenPaymaster.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorSession) Deposit() (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.Deposit(&_LegacyTokenPaymaster.TransactOpts)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(address recipient, uint256 amount) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactor) MintTokens(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.contract.Transact(opts, "mintTokens", recipient, amount)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(address recipient, uint256 amount) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) MintTokens(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.MintTokens(&_LegacyTokenPaymaster.TransactOpts, recipient, amount)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(address recipient, uint256 amount) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorSession) MintTokens(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.MintTokens(&_LegacyTokenPaymaster.TransactOpts, recipient, amount)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.PostOp(&_LegacyTokenPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.PostOp(&_LegacyTokenPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.RenounceOwnership(&_LegacyTokenPaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.RenounceOwnership(&_LegacyTokenPaymaster.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.Transfer(&_LegacyTokenPaymaster.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.Transfer(&_LegacyTokenPaymaster.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.TransferFrom(&_LegacyTokenPaymaster.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.TransferFrom(&_LegacyTokenPaymaster.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.TransferOwnership(&_LegacyTokenPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.TransferOwnership(&_LegacyTokenPaymaster.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) UnlockStake() (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.UnlockStake(&_LegacyTokenPaymaster.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.UnlockStake(&_LegacyTokenPaymaster.TransactOpts)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.ValidatePaymasterUserOp(&_LegacyTokenPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.ValidatePaymasterUserOp(&_LegacyTokenPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.WithdrawStake(&_LegacyTokenPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.WithdrawStake(&_LegacyTokenPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.WithdrawTo(&_LegacyTokenPaymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_LegacyTokenPaymaster *LegacyTokenPaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyTokenPaymaster.Contract.WithdrawTo(&_LegacyTokenPaymaster.TransactOpts, withdrawAddress, amount)
}

// LegacyTokenPaymasterApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LegacyTokenPaymaster contract.
type LegacyTokenPaymasterApprovalIterator struct {
	Event *LegacyTokenPaymasterApproval // Event containing the contract specifics and raw log

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
func (it *LegacyTokenPaymasterApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyTokenPaymasterApproval)
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
		it.Event = new(LegacyTokenPaymasterApproval)
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
func (it *LegacyTokenPaymasterApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyTokenPaymasterApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyTokenPaymasterApproval represents a Approval event raised by the LegacyTokenPaymaster contract.
type LegacyTokenPaymasterApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LegacyTokenPaymasterApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LegacyTokenPaymaster.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LegacyTokenPaymasterApprovalIterator{contract: _LegacyTokenPaymaster.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LegacyTokenPaymasterApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LegacyTokenPaymaster.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyTokenPaymasterApproval)
				if err := _LegacyTokenPaymaster.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LegacyTokenPaymaster *LegacyTokenPaymasterFilterer) ParseApproval(log types.Log) (*LegacyTokenPaymasterApproval, error) {
	event := new(LegacyTokenPaymasterApproval)
	if err := _LegacyTokenPaymaster.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyTokenPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LegacyTokenPaymaster contract.
type LegacyTokenPaymasterOwnershipTransferredIterator struct {
	Event *LegacyTokenPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LegacyTokenPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyTokenPaymasterOwnershipTransferred)
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
		it.Event = new(LegacyTokenPaymasterOwnershipTransferred)
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
func (it *LegacyTokenPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyTokenPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyTokenPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the LegacyTokenPaymaster contract.
type LegacyTokenPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LegacyTokenPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LegacyTokenPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LegacyTokenPaymasterOwnershipTransferredIterator{contract: _LegacyTokenPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LegacyTokenPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LegacyTokenPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyTokenPaymasterOwnershipTransferred)
				if err := _LegacyTokenPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LegacyTokenPaymaster *LegacyTokenPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*LegacyTokenPaymasterOwnershipTransferred, error) {
	event := new(LegacyTokenPaymasterOwnershipTransferred)
	if err := _LegacyTokenPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyTokenPaymasterTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LegacyTokenPaymaster contract.
type LegacyTokenPaymasterTransferIterator struct {
	Event *LegacyTokenPaymasterTransfer // Event containing the contract specifics and raw log

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
func (it *LegacyTokenPaymasterTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyTokenPaymasterTransfer)
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
		it.Event = new(LegacyTokenPaymasterTransfer)
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
func (it *LegacyTokenPaymasterTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyTokenPaymasterTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyTokenPaymasterTransfer represents a Transfer event raised by the LegacyTokenPaymaster contract.
type LegacyTokenPaymasterTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LegacyTokenPaymasterTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LegacyTokenPaymaster.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LegacyTokenPaymasterTransferIterator{contract: _LegacyTokenPaymaster.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LegacyTokenPaymaster *LegacyTokenPaymasterFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LegacyTokenPaymasterTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LegacyTokenPaymaster.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyTokenPaymasterTransfer)
				if err := _LegacyTokenPaymaster.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LegacyTokenPaymaster *LegacyTokenPaymasterFilterer) ParseTransfer(log types.Log) (*LegacyTokenPaymasterTransfer, error) {
	event := new(LegacyTokenPaymasterTransfer)
	if err := _LegacyTokenPaymaster.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
