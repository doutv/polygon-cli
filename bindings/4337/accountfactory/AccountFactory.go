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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"config\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocktime\",\"type\":\"uint256\"}],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"}],\"name\":\"InvalidSingleton\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_initializer\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONFIG\",\"outputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"computeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"createAccountWithSignature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isValidAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60e06040908082523461019057806116a7803803809161001f8285610195565b833960209283918101031261019057516001600160a01b03811690819003610190573060805260a052815161027b61005983820183610195565b80825261142c83830139825190828282519260005b84811061017957505061008e928101600083820152038084520182610195565b81815191012060c0527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549160ff83851c16610168576001600160401b03906002600160401b031984831601610126575b845161125d90816101cf8239608051818181610596015261064b015260a051818181610208015281816102d6015281816109450152610bb3015260c051816108790152f35b6001600160401b0319909316811790925582519182527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d291a1388080806100e1565b835163f92ee8a960e01b8152600490fd5b81810184015186820185015286938693500161006e565b600080fd5b601f909101601f19168101906001600160401b038211908210176101b857604052565b634e487b7160e01b600052604160045260246000fdfe608060409080825260048036101561001657600080fd5b600092833560e01c92836323cca69c146109c557508263296601cd146108e557826336b5aa2d1461083c5782634f1ef286146105fa57826352d1902d14610582578263715018a6146105185782638da5cb5b146104e2578263ad3cb1cc14610443578263c4d66de814610305578263d92e82e4146102c1578263ddede0a1146100d957505063f2fde38b146100aa57600080fd5b346100d65760203660031901126100d6576100d36100c66109fe565b6100ce610d5e565b610b0b565b80f35b80fd5b8390346102b55760803660031901126102b5576100f46109fe565b67ffffffffffffffff906024358281116102bd576101159036908701610a19565b91604435936064359081116102b9576101319036908901610a19565b9790966020978989116102b5578035428110610299576101d78a9b6101dd938b809d8c8c966101e69f8d61017f9151998a95860196468852309087015260a0606087015260c0860191610aea565b91608084015260a0830152039461019e601f1996878101835282610a47565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008752601c528d603c8720943693019101610a9b565b90610dd8565b90989198610e14565b8751636e786d3960e11b81526001600160a01b039788168382015289816024817f00000000000000000000000000000000000000000000000000000000000000008c165afa91821561028e5791610261575b501561025357509061024b939291610b7f565b915191168152f35b8651638baa579f60e01b8152fd5b6102819150893d8b11610287575b6102798183610a47565b810190610ad2565b89610238565b503d61026f565b8951903d90823e3d90fd5b836044918a519163aa2fd92560e01b8352820152426024820152fd5b5080fd5b8680fd5b8480fd5b8390346102b557816003193601126102b557517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b91503461043f57602036600319011261043f576103206109fe565b907ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0091825460ff81861c16159267ffffffffffffffff821680159081610437575b600114908161042d575b159081610424575b50610416575067ffffffffffffffff19811660011784556103a69190836103f7575b5061039e610d97565b6100ce610d97565b6103ae610d97565b6103b6578280f35b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808280f35b68ffffffffffffffffff19166801000000000000000117845538610395565b855163f92ee8a960e01b8152fd5b90501538610373565b303b15915061036b565b859150610361565b8280fd5b91503461043f578260031936011261043f578151908282019082821067ffffffffffffffff8311176104cf5750825260058152602090640352e302e360dc1b6020820152825193849260208452825192836020860152825b8481106104b957505050828201840152601f01601f19168101030190f35b818101830151888201880152879550820161049b565b634e487b7160e01b855260419052602484fd5b8390346102b557816003193601126102b5576000805160206112088339815191525490516001600160a01b039091168152602090f35b83346100d657806003193601126100d657610531610d5e565b60008051602061120883398151915280546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b83346100d657806003193601126100d657507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036105ed57602090517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b5163703e46dd60e11b8152fd5b8091925060031936011261043f576106106109fe565b9060243567ffffffffffffffff81116102bd57366023820112156102bd576106419036906024818701359101610a9b565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811630811490811561080e575b506107fe57610684610d5e565b82516352d1902d60e01b81529084169360209182818881895afa8891816107cb575b506106c2578451634c9c8ce360e01b8152808801879052602490fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc969295919396908181036107b55750833b1561079e5780546001600160a01b0319168217905583518792917fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8480a286511561078157505080858561077397519101845af4913d15610777573d61076561075c82610a7f565b92519283610a47565b81528581943d92013e610f29565b5080f35b5060609250610f29565b945094505050503461079257505080f35b63b398979f60e01b8152fd5b8451634c9c8ce360e01b8152808401839052602490fd5b83602491875191632a87526960e21b8352820152fd5b9091508381813d83116107f7575b6107e38183610a47565b810103126107f3575190386106a6565b8880fd5b503d6107d9565b825163703e46dd60e11b81528590fd5b9050817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416141538610677565b8390346102b557806003193601126102b5576108566109fe565b50805190602082019060ff60f81b82523060601b602184015260243560358401527f0000000000000000000000000000000000000000000000000000000000000000605584015260558352608083019383851067ffffffffffffffff8611176108d25750839052905190206001600160a01b0316815260209150f35b634e487b7160e01b815260418652602490fd5b83346100d65760603660031901126100d6576108ff6109fe565b9260243567ffffffffffffffff811161043f5761091f9036908301610a19565b8451636e786d3960e11b815233818501526001600160a01b0394919391906020816024817f00000000000000000000000000000000000000000000000000000000000000008a165afa9182156109ba579161099b575b501561098d57509361024b9160209560443592610b7f565b8451631eb49d6d60e11b8152fd5b6109b4915060203d602011610287576102798183610a47565b87610975565b8751903d90823e3d90fd5b84913461043f57602036600319011261043f5760209260ff91906001600160a01b036109ef6109fe565b16815280855220541615158152f35b600435906001600160a01b0382168203610a1457565b600080fd5b9181601f84011215610a145782359167ffffffffffffffff8311610a145760208381860195010111610a1457565b90601f8019910116810190811067ffffffffffffffff821117610a6957604052565b634e487b7160e01b600052604160045260246000fd5b67ffffffffffffffff8111610a6957601f01601f191660200190565b929192610aa782610a7f565b91610ab56040519384610a47565b829481845281830111610a14578281602093846000960137010152565b90816020910312610a1457518015158103610a145790565b908060209392818452848401376000828201840152601f01601f1916010190565b6001600160a01b03908116908115610b665760008051602061120883398151915280546001600160a01b031981168417909155167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b604051631e4fbdf760e01b815260006004820152602490fd5b604080516318914b1360e31b81526001600160a01b03928316600482018190529296959394600094909392916020816024817f00000000000000000000000000000000000000000000000000000000000000008d165afa908115610d09578691610d3f575b5015610d275781519761027b9889810167ffffffffffffffff9a8282108c831117610d13579180918593610f8d8339039088f58015610d09571695863b15610d0557825163347d5e2560e21b8152846004820152836024820152868180610c4f604482018a87610aea565b0381838c5af1998a15610cfb5788999a989798610cc6575b5050610cbc908388887fa4ec333d142e947b3345528c6cbc210be703d984f8df2c3d589f2b3ea39f7437999a528060205220600160ff1982541617905583519586958652606060208701526060860191610aea565b918301520390a290565b9080929496989395975011610ce75784529486949093909290918083610c67565b634e487b7160e01b82526041600452602482fd5b84513d89823e3d90fd5b8580fd5b83513d88823e3d90fd5b634e487b7160e01b89526041600452602489fd5b8151630f00f1a160e11b815260048101849052602490fd5b610d58915060203d602011610287576102798183610a47565b38610be4565b600080516020611208833981519152546001600160a01b03163303610d7f57565b60405163118cdaa760e01b8152336004820152602490fd5b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c1615610dc657565b604051631afcd79f60e31b8152600490fd5b8151919060418303610e0957610e0292506020820151906060604084015193015160001a90610e99565b9192909190565b505060009160029190565b6004811015610e835780610e26575050565b60018103610e405760405163f645eedf60e01b8152600490fd5b60028103610e615760405163fce698f760e01b815260048101839052602490fd5b600314610e6b5750565b602490604051906335e2f38360e21b82526004820152fd5b634e487b7160e01b600052602160045260246000fd5b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411610f1d57926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa15610f115780516001600160a01b03811615610f0857918190565b50809160019190565b604051903d90823e3d90fd5b50505060009160039190565b90610f505750805115610f3e57805190602001fd5b60405163d6bda27560e01b8152600490fd5b81511580610f83575b610f61575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b15610f5956fe60808060405234601557610260908161001b8239f35b600080fdfe60806040526004361015610024575b361561001f5734156101eb57600080fd5b6101eb565b6000803560e01c63d1f578941461003b575061000e565b346100af5760403660031901126100af576004356001600160a01b03811681036100ab576024359067ffffffffffffffff908183116100a757366023840112156100a75782600401359182116100a75736602483850101116100a75760246100a4930190610111565b80f35b8380fd5b5080fd5b80fd5b634e487b7160e01b600052604160045260246000fd5b6020808252825181830181905290939260005b8281106100fd57505060409293506000838284010152601f8019910116010190565b8181018601518482016040015285016100db565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8054929390926001600160a01b03166101d95760009382859455816040519283928337810184815203915af43d156101d15767ffffffffffffffff903d8281116101cc5760405192601f8201601f19908116603f01168401908111848210176101cc5760405282523d6000602084013e5b156101ab5750565b604051633018224d60e21b81529081906101c890600483016100c8565b0390fd5b6100b2565b6060906101a3565b60405163c28d69c760e01b8152600490fd5b600036818037808036817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc545af43d82803e15610226573d90f35b3d90fdfea2646970667358221220dbc2ad734b73b998278ed7bcd9ef361791ffe94471f015bd78d961860b387f4664736f6c634300081900339016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300a2646970667358221220b67e3ad7adf4c65e3a32b36fc292197c915c934254343b787164436983d6448964736f6c6343000819003360808060405234601557610260908161001b8239f35b600080fdfe60806040526004361015610024575b361561001f5734156101eb57600080fd5b6101eb565b6000803560e01c63d1f578941461003b575061000e565b346100af5760403660031901126100af576004356001600160a01b03811681036100ab576024359067ffffffffffffffff908183116100a757366023840112156100a75782600401359182116100a75736602483850101116100a75760246100a4930190610111565b80f35b8380fd5b5080fd5b80fd5b634e487b7160e01b600052604160045260246000fd5b6020808252825181830181905290939260005b8281106100fd57505060409293506000838284010152601f8019910116010190565b8181018601518482016040015285016100db565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8054929390926001600160a01b03166101d95760009382859455816040519283928337810184815203915af43d156101d15767ffffffffffffffff903d8281116101cc5760405192601f8201601f19908116603f01168401908111848210176101cc5760405282523d6000602084013e5b156101ab5750565b604051633018224d60e21b81529081906101c890600483016100c8565b0390fd5b6100b2565b6060906101a3565b60405163c28d69c760e01b8152600490fd5b600036818037808036817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc545af43d82803e15610226573d90f35b3d90fdfea2646970667358221220dbc2ad734b73b998278ed7bcd9ef361791ffe94471f015bd78d961860b387f4664736f6c63430008190033",
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

// IsValidAccount is a free data retrieval call binding the contract method 0x23cca69c.
//
// Solidity: function isValidAccount(address ) view returns(bool)
func (_AccountFactory *AccountFactoryCaller) IsValidAccount(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "isValidAccount", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAccount is a free data retrieval call binding the contract method 0x23cca69c.
//
// Solidity: function isValidAccount(address ) view returns(bool)
func (_AccountFactory *AccountFactorySession) IsValidAccount(arg0 common.Address) (bool, error) {
	return _AccountFactory.Contract.IsValidAccount(&_AccountFactory.CallOpts, arg0)
}

// IsValidAccount is a free data retrieval call binding the contract method 0x23cca69c.
//
// Solidity: function isValidAccount(address ) view returns(bool)
func (_AccountFactory *AccountFactoryCallerSession) IsValidAccount(arg0 common.Address) (bool, error) {
	return _AccountFactory.Contract.IsValidAccount(&_AccountFactory.CallOpts, arg0)
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

// CreateAccountWithSignature is a paid mutator transaction binding the contract method 0xddede0a1.
//
// Solidity: function createAccountWithSignature(address _implementation, bytes _initializer, uint256 _salt, bytes _signature) returns(address)
func (_AccountFactory *AccountFactoryTransactor) CreateAccountWithSignature(opts *bind.TransactOpts, _implementation common.Address, _initializer []byte, _salt *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "createAccountWithSignature", _implementation, _initializer, _salt, _signature)
}

// CreateAccountWithSignature is a paid mutator transaction binding the contract method 0xddede0a1.
//
// Solidity: function createAccountWithSignature(address _implementation, bytes _initializer, uint256 _salt, bytes _signature) returns(address)
func (_AccountFactory *AccountFactorySession) CreateAccountWithSignature(_implementation common.Address, _initializer []byte, _salt *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountFactory.Contract.CreateAccountWithSignature(&_AccountFactory.TransactOpts, _implementation, _initializer, _salt, _signature)
}

// CreateAccountWithSignature is a paid mutator transaction binding the contract method 0xddede0a1.
//
// Solidity: function createAccountWithSignature(address _implementation, bytes _initializer, uint256 _salt, bytes _signature) returns(address)
func (_AccountFactory *AccountFactoryTransactorSession) CreateAccountWithSignature(_implementation common.Address, _initializer []byte, _salt *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountFactory.Contract.CreateAccountWithSignature(&_AccountFactory.TransactOpts, _implementation, _initializer, _salt, _signature)
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
