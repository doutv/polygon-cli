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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBundler\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"FactorySignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"PaySignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"RecoverySignerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"RecoverySignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetSafeSingleton\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetSenderSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumVerifierType\",\"name\":\"verifier\",\"type\":\"uint8\"}],\"name\":\"VerifierTypeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumVerifierType\",\"name\":\"verifier\",\"type\":\"uint8\"}],\"name\":\"VerifierTypeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"WhitelistBundlerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"WhitelistBundlerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moduleType\",\"type\":\"uint256\"}],\"name\":\"WhitelistModuleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"WhitelistModuleRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"addRecoverySigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"}],\"name\":\"addSafeSingleton\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"addWhitelistedBundlers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"moduleTypes\",\"type\":\"uint256[]\"}],\"name\":\"addWhitelistedModules\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumVerifierType\",\"name\":\"verifier\",\"type\":\"uint8\"}],\"name\":\"addWhitelistedVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"wallets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"}],\"name\":\"batchResetWalletSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factorySigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"initializeWalletSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isFactorySigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isPaySigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRecoverySigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSafeSingleton\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelistedBundler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumVerifierType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"isWhitelistedVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paySigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"removeRecoverySigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"}],\"name\":\"removeSafeSingleton\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bundlers\",\"type\":\"address[]\"}],\"name\":\"removeWhitelistedBundlers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"name\":\"removeWhitelistedModules\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumVerifierType\",\"name\":\"verifier\",\"type\":\"uint8\"}],\"name\":\"removeWhitelistedVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"setFactorySigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"setPaySigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"walletSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedModuleType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0806040523460c857306080527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c1660b957506001600160401b036002600160401b0319828216016075575b6040516115ac90816100ce8239608051818181610a140152610af20152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a13880806056565b63f92ee8a960e01b8152600490fd5b600080fdfe60406080815260048036101561001457600080fd5b600091823560e01c806303087a571461118d578063158211281461106157806316b97e9a14610fbd57806325d480ad14610f8357806332b0cc3214610f0f578063485cc95514610dc25780634b04ad5114610ce15780634f1ef28614610a7857806352d1902d146109ff578063562d2e2714610995578063589597cb1461095d5780635aa64413146109345780637063908f146108b1578063715018a61461084757806381e36802146107dd5780638da5cb5b146107a75780638eaa17d91461077e578063936587831461074057806396a5d35b14610703578063a884c15314610691578063ad3cb1cc146105f2578063b55b7066146105bb578063b80767421461051d578063c22305cf146104b6578063c48a589814610478578063d4c4c2b014610348578063dcf0da7214610314578063e53a7eae146102dc578063eb6810be1461025c578063f2fde38b1461022c5763fcdc47271461017557600080fd5b346102285760209160206003193601126102245780359067ffffffffffffffff8211610220576101a7913691016112a4565b916101b0611479565b8251845b8181106101f257857f49bd286efd7e26d327db1b6f433560c5f6156363198289a65a6ca9904338f4f46101ec878751918291826113ad565b0390a180f35b6001600160a01b036102048287611399565b511686526001808452848720805460ff191682179055016101b4565b8480fd5b8380fd5b8280fd5b833461025957602036600319011261025957610256610249611208565b610251611479565b611405565b80f35b80fd5b503461022857602036600319011261022857610276611208565b61027e611479565b6001600160a01b03169182156102ce5750600780546001600160a01b03191683179055519081527f322933b0e9c97ef2d23835a0a1835c057a9601a32b8227690193a42c57bacf3d90602090a180f35b905163e6c4247b60e01b8152fd5b838234610310576020366003190112610310576020906102fa611208565b60018060a01b0380600754169116149051908152f35b5080fd5b83823461031057602036600319011261031057602090610332611208565b60018060a01b0380600654169116149051908152f35b503461022857806003193601126102285767ffffffffffffffff82358181116102205761037890369085016112a4565b602435918211610220573660238301121561022057818401359061039b8261128c565b926103a885519485611254565b8284526020926024602086019160051b83010191368311610474576024859101915b83831061046457505050506103dd611479565b805194835186036104565750855b8581106103f6578680f35b6001907f51a45eb94bbc0121a4cdcf5f28257162a61221f0d6f8897adffe076012457865866001600160a01b0361042d8487611399565b51166104398489611399565b51818c526002885280838d2055825191825287820152a1016103eb565b8451631985132360e31b8152fd5b82358152918101918591016103ca565b8880fd5b8382346103105760203660031901126103105760209160ff9082906001600160a01b036104a3611208565b1681526003855220541690519015158152f35b5090346102285760203660031901126102285735906003821015610228577ff7ad1892e8c359fbe99be323cb12129c63748ec6c859ad3dd7f22f427742ebab916101ec91610502611479565b61050b82611330565b805460ff1916905551918291826113f2565b5090346102285760209060206003193601126102245780359067ffffffffffffffff821161022057610551913691016112a4565b9161055a611479565b825192845b84811061056a578580f35b6001907f312c33106e0d74000c54d884caee742f0a33f8143ed96561cc1c80d3351b2cd2856001600160a01b036105a18487611399565b5116808a526002825289878120558651908152a10161055f565b5090346102285760203660031901126102285735916003831015610259575060ff6105e7602093611330565b541690519015158152f35b5090346102285782600319360112610228578151908282019082821067ffffffffffffffff83111761067e5750825260058152602090640352e302e360dc1b6020820152825193849260208452825192836020860152825b84811061066857505050828201840152601f01601f19168101030190f35b818101830151888201880152879550820161064a565b634e487b7160e01b855260419052602484fd5b5034610228576020366003190112610228576106ab611208565b6106b3611479565b6001600160a01b03169182156102ce5750600680546001600160a01b03191683179055519081527f789d9446447ccefd52005113ddeccb6cfdcafd57a9f5229050f0b6c3dcd10f7990602090a180f35b8382346103105760203660031901126103105760209160ff9082906001600160a01b0361072e611208565b16815280855220541690519015158152f35b8382346103105760203660031901126103105760209160ff9082906001600160a01b0361076b611208565b1681526001855220541690519015158152f35b83823461031057816003193601126103105760075490516001600160a01b039091168152602090f35b8382346103105781600319360112610310576000805160206115578339815191525490516001600160a01b039091168152602090f35b5090346102285760203660031901126102285735906003821015610228577fc007db510c05b52ed8aa21436b678ffa49a3f5de71cb7ba2c81363eda6f9a734916101ec91610829611479565b61083282611330565b805460ff1916600117905551918291826113f2565b8334610259578060031936011261025957610860611479565b60008051602061155783398151915280546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b5034610228576020366003190112610228576108cb611208565b6108d3611479565b6001600160a01b03169182156102ce5750816060917ff80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f804061919385526003602052808520600160ff198254161790558051918252600160208301524290820152a180f35b83823461031057816003193601126103105760065490516001600160a01b039091168152602090f35b8382346103105760203660031901126103105760209181906001600160a01b03610985611208565b1681526002845220549051908152f35b8382346103105760203660031901126103105760207fadd78f873e935721c08cc4e66ec8de5f0bd111d6a9966eaa93b5e22b417d23a4916109d4611208565b6109dc611479565b6001600160a01b0316808552848352818520805460ff191690559051908152a180f35b508234610259578060031936011261025957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610a6b57602090517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b5163703e46dd60e11b8152fd5b508060031936011261022857610a8c611208565b90602493843567ffffffffffffffff811161031057366023820112156103105780850135610ab981611314565b94610ac685519687611254565b81865260209182870193368a8383010111610cdd578186928b8693018737880101526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116308114908115610caf575b50610c9f57610b2b611479565b81169585516352d1902d60e01b815283818a818b5afa869181610c6c575b50610b65575050505050505191634c9c8ce360e01b8352820152fd5b9088888894938c7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc91828103610c575750853b15610c43575080546001600160a01b031916821790558451889392917fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8580a2825115610c25575050610c179582915190845af4913d15610c1b573d610c09610c0082611314565b92519283611254565b81528581943d92013e6114f3565b5080f35b50606092506114f3565b955095505050505034610c3757505080f35b63b398979f60e01b8152fd5b8651634c9c8ce360e01b8152808501849052fd5b8751632a87526960e21b815280860191909152fd5b9091508481813d8311610c98575b610c848183611254565b81010312610c9457519038610b49565b8680fd5b503d610c7a565b855163703e46dd60e11b81528890fd5b9050817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416141538610b1e565b8580fd5b503461022857602036600319011261022857610cfb611208565b916001600160a01b0380841691908215908115610dac575b50610d9e57328552600160205260ff838620541615610d905733808652602091825283862080546001600160a01b03191690931790925591519081526001600160a01b03909216908201524260408201527f03b72cb2fffc620d2a9c4f6bce0636faa82c948488a45f2f4b0684bb17cfdc729080606081016101ec565b825163f8a0c54b60e01b8152fd5b825163e6c4247b60e01b8152fd5b9050338652816020528386205416151538610d13565b509034610228578160031936011261022857610ddc611208565b50602435906001600160a01b0382168203610224577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0091825460ff81861c16159267ffffffffffffffff821680159081610f07575b6001149081610efd575b159081610ef4575b50610ee6575067ffffffffffffffff1981166001178455610e76919083610ec7575b50610e6e6114b2565b6102516114b2565b610e7e6114b2565b610e86578280f35b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808280f35b68ffffffffffffffffff19166801000000000000000117845538610e65565b855163f92ee8a960e01b8152fd5b90501538610e43565b303b159150610e3b565b859150610e31565b503461022857602036600319011261022857610f29611208565b610f31611479565b6001600160a01b03169182156102ce57508183526020838152818420805460ff1916600117905590519182527f22dbed02a51e94f91fc3a6d1d2014ed6bc570834e7f3536162d626c429d90e5491a180f35b509034610228576020366003190112610228576020926001600160a01b039183919083610fae611208565b16825285522054169051908152f35b50346102285760209160206003193601126102245780359067ffffffffffffffff821161022057610ff0913691016112a4565b91610ff9611479565b8251845b81811061103557857f623f7acab9936b83f17999b4d96bfda34b1d3a749bd474d39d0413cc977044ed6101ec878751918291826113ad565b6001600160a01b036110478287611399565b511686526001808452848720805460ff1916905501610ffd565b503461022857806003193601126102285767ffffffffffffffff8235818111610220576110919036908501611223565b929091602435908111610cdd576110ab9036908601611223565b6110b6949194611479565b81810361117d57865b8181106110ca578780f35b6110dd6110d882848961135f565b611385565b6110eb6110d883868961135f565b906001600160a01b0380821690811561116d5783168b5260208a8152878c2080546001600160a01b03191690921790915586516001600160a01b0393841681529290911690820152426040820152600191907f03b72cb2fffc620d2a9c4f6bce0636faa82c948488a45f2f4b0684bb17cfdc729080606081010390a1016110bf565b875163e6c4247b60e01b81528b90fd5b8251631985132360e31b81528690fd5b8382346103105760203660031901126103105760607ff80dbc7f86c01c1d15ec0c595ff86fcfe195a0a197a73e1a45f1d52f80406191916111cc611208565b6111d4611479565b6001600160a01b031680855260036020908152828620805460ff19169055825191825281018590524291810191909152a180f35b600435906001600160a01b038216820361121e57565b600080fd5b9181601f8401121561121e5782359167ffffffffffffffff831161121e576020808501948460051b01011161121e57565b90601f8019910116810190811067ffffffffffffffff82111761127657604052565b634e487b7160e01b600052604160045260246000fd5b67ffffffffffffffff81116112765760051b60200190565b9080601f8301121561121e5760209082356112be8161128c565b936112cc6040519586611254565b81855260208086019260051b82010192831161121e57602001905b8282106112f5575050505090565b81356001600160a01b038116810361121e5781529083019083016112e7565b67ffffffffffffffff811161127657601f01601f191660200190565b6003811015611349576000526005602052604060002090565b634e487b7160e01b600052602160045260246000fd5b919081101561136f5760051b0190565b634e487b7160e01b600052603260045260246000fd5b356001600160a01b038116810361121e5790565b805182101561136f5760209160051b010190565b602090602060408183019282815285518094520193019160005b8281106113d5575050505090565b83516001600160a01b0316855293810193928101926001016113c7565b9190602083019260038210156113495752565b6001600160a01b039081169081156114605760008051602061155783398151915280546001600160a01b031981168417909155167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b604051631e4fbdf760e01b815260006004820152602490fd5b600080516020611557833981519152546001600160a01b0316330361149a57565b60405163118cdaa760e01b8152336004820152602490fd5b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c16156114e157565b604051631afcd79f60e31b8152600490fd5b9061151a575080511561150857805190602001fd5b60405163d6bda27560e01b8152600490fd5b8151158061154d575b61152b575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561152356fe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300a2646970667358221220b54d21970aa88563086c93db9a73555a3beb275da508207cad96c99112bd648464736f6c63430008190033",
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

// FactorySigner is a free data retrieval call binding the contract method 0x5aa64413.
//
// Solidity: function factorySigner() view returns(address)
func (_Config *ConfigCaller) FactorySigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "factorySigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactorySigner is a free data retrieval call binding the contract method 0x5aa64413.
//
// Solidity: function factorySigner() view returns(address)
func (_Config *ConfigSession) FactorySigner() (common.Address, error) {
	return _Config.Contract.FactorySigner(&_Config.CallOpts)
}

// FactorySigner is a free data retrieval call binding the contract method 0x5aa64413.
//
// Solidity: function factorySigner() view returns(address)
func (_Config *ConfigCallerSession) FactorySigner() (common.Address, error) {
	return _Config.Contract.FactorySigner(&_Config.CallOpts)
}

// IsFactorySigner is a free data retrieval call binding the contract method 0xdcf0da72.
//
// Solidity: function isFactorySigner(address sender) view returns(bool)
func (_Config *ConfigCaller) IsFactorySigner(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "isFactorySigner", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFactorySigner is a free data retrieval call binding the contract method 0xdcf0da72.
//
// Solidity: function isFactorySigner(address sender) view returns(bool)
func (_Config *ConfigSession) IsFactorySigner(sender common.Address) (bool, error) {
	return _Config.Contract.IsFactorySigner(&_Config.CallOpts, sender)
}

// IsFactorySigner is a free data retrieval call binding the contract method 0xdcf0da72.
//
// Solidity: function isFactorySigner(address sender) view returns(bool)
func (_Config *ConfigCallerSession) IsFactorySigner(sender common.Address) (bool, error) {
	return _Config.Contract.IsFactorySigner(&_Config.CallOpts, sender)
}

// IsPaySigner is a free data retrieval call binding the contract method 0xe53a7eae.
//
// Solidity: function isPaySigner(address sender) view returns(bool)
func (_Config *ConfigCaller) IsPaySigner(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "isPaySigner", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaySigner is a free data retrieval call binding the contract method 0xe53a7eae.
//
// Solidity: function isPaySigner(address sender) view returns(bool)
func (_Config *ConfigSession) IsPaySigner(sender common.Address) (bool, error) {
	return _Config.Contract.IsPaySigner(&_Config.CallOpts, sender)
}

// IsPaySigner is a free data retrieval call binding the contract method 0xe53a7eae.
//
// Solidity: function isPaySigner(address sender) view returns(bool)
func (_Config *ConfigCallerSession) IsPaySigner(sender common.Address) (bool, error) {
	return _Config.Contract.IsPaySigner(&_Config.CallOpts, sender)
}

// IsRecoverySigner is a free data retrieval call binding the contract method 0x96a5d35b.
//
// Solidity: function isRecoverySigner(address ) view returns(bool)
func (_Config *ConfigCaller) IsRecoverySigner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "isRecoverySigner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRecoverySigner is a free data retrieval call binding the contract method 0x96a5d35b.
//
// Solidity: function isRecoverySigner(address ) view returns(bool)
func (_Config *ConfigSession) IsRecoverySigner(arg0 common.Address) (bool, error) {
	return _Config.Contract.IsRecoverySigner(&_Config.CallOpts, arg0)
}

// IsRecoverySigner is a free data retrieval call binding the contract method 0x96a5d35b.
//
// Solidity: function isRecoverySigner(address ) view returns(bool)
func (_Config *ConfigCallerSession) IsRecoverySigner(arg0 common.Address) (bool, error) {
	return _Config.Contract.IsRecoverySigner(&_Config.CallOpts, arg0)
}

// IsSafeSingleton is a free data retrieval call binding the contract method 0xc48a5898.
//
// Solidity: function isSafeSingleton(address ) view returns(bool)
func (_Config *ConfigCaller) IsSafeSingleton(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "isSafeSingleton", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSafeSingleton is a free data retrieval call binding the contract method 0xc48a5898.
//
// Solidity: function isSafeSingleton(address ) view returns(bool)
func (_Config *ConfigSession) IsSafeSingleton(arg0 common.Address) (bool, error) {
	return _Config.Contract.IsSafeSingleton(&_Config.CallOpts, arg0)
}

// IsSafeSingleton is a free data retrieval call binding the contract method 0xc48a5898.
//
// Solidity: function isSafeSingleton(address ) view returns(bool)
func (_Config *ConfigCallerSession) IsSafeSingleton(arg0 common.Address) (bool, error) {
	return _Config.Contract.IsSafeSingleton(&_Config.CallOpts, arg0)
}

// IsWhitelistedBundler is a free data retrieval call binding the contract method 0x93658783.
//
// Solidity: function isWhitelistedBundler(address ) view returns(bool)
func (_Config *ConfigCaller) IsWhitelistedBundler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "isWhitelistedBundler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedBundler is a free data retrieval call binding the contract method 0x93658783.
//
// Solidity: function isWhitelistedBundler(address ) view returns(bool)
func (_Config *ConfigSession) IsWhitelistedBundler(arg0 common.Address) (bool, error) {
	return _Config.Contract.IsWhitelistedBundler(&_Config.CallOpts, arg0)
}

// IsWhitelistedBundler is a free data retrieval call binding the contract method 0x93658783.
//
// Solidity: function isWhitelistedBundler(address ) view returns(bool)
func (_Config *ConfigCallerSession) IsWhitelistedBundler(arg0 common.Address) (bool, error) {
	return _Config.Contract.IsWhitelistedBundler(&_Config.CallOpts, arg0)
}

// IsWhitelistedVerifier is a free data retrieval call binding the contract method 0xb55b7066.
//
// Solidity: function isWhitelistedVerifier(uint8 ) view returns(bool)
func (_Config *ConfigCaller) IsWhitelistedVerifier(opts *bind.CallOpts, arg0 uint8) (bool, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "isWhitelistedVerifier", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedVerifier is a free data retrieval call binding the contract method 0xb55b7066.
//
// Solidity: function isWhitelistedVerifier(uint8 ) view returns(bool)
func (_Config *ConfigSession) IsWhitelistedVerifier(arg0 uint8) (bool, error) {
	return _Config.Contract.IsWhitelistedVerifier(&_Config.CallOpts, arg0)
}

// IsWhitelistedVerifier is a free data retrieval call binding the contract method 0xb55b7066.
//
// Solidity: function isWhitelistedVerifier(uint8 ) view returns(bool)
func (_Config *ConfigCallerSession) IsWhitelistedVerifier(arg0 uint8) (bool, error) {
	return _Config.Contract.IsWhitelistedVerifier(&_Config.CallOpts, arg0)
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

// PaySigner is a free data retrieval call binding the contract method 0x8eaa17d9.
//
// Solidity: function paySigner() view returns(address)
func (_Config *ConfigCaller) PaySigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "paySigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaySigner is a free data retrieval call binding the contract method 0x8eaa17d9.
//
// Solidity: function paySigner() view returns(address)
func (_Config *ConfigSession) PaySigner() (common.Address, error) {
	return _Config.Contract.PaySigner(&_Config.CallOpts)
}

// PaySigner is a free data retrieval call binding the contract method 0x8eaa17d9.
//
// Solidity: function paySigner() view returns(address)
func (_Config *ConfigCallerSession) PaySigner() (common.Address, error) {
	return _Config.Contract.PaySigner(&_Config.CallOpts)
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

// WalletSigner is a free data retrieval call binding the contract method 0x25d480ad.
//
// Solidity: function walletSigner(address ) view returns(address)
func (_Config *ConfigCaller) WalletSigner(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "walletSigner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletSigner is a free data retrieval call binding the contract method 0x25d480ad.
//
// Solidity: function walletSigner(address ) view returns(address)
func (_Config *ConfigSession) WalletSigner(arg0 common.Address) (common.Address, error) {
	return _Config.Contract.WalletSigner(&_Config.CallOpts, arg0)
}

// WalletSigner is a free data retrieval call binding the contract method 0x25d480ad.
//
// Solidity: function walletSigner(address ) view returns(address)
func (_Config *ConfigCallerSession) WalletSigner(arg0 common.Address) (common.Address, error) {
	return _Config.Contract.WalletSigner(&_Config.CallOpts, arg0)
}

// WhitelistedModuleType is a free data retrieval call binding the contract method 0x589597cb.
//
// Solidity: function whitelistedModuleType(address ) view returns(uint256)
func (_Config *ConfigCaller) WhitelistedModuleType(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "whitelistedModuleType", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistedModuleType is a free data retrieval call binding the contract method 0x589597cb.
//
// Solidity: function whitelistedModuleType(address ) view returns(uint256)
func (_Config *ConfigSession) WhitelistedModuleType(arg0 common.Address) (*big.Int, error) {
	return _Config.Contract.WhitelistedModuleType(&_Config.CallOpts, arg0)
}

// WhitelistedModuleType is a free data retrieval call binding the contract method 0x589597cb.
//
// Solidity: function whitelistedModuleType(address ) view returns(uint256)
func (_Config *ConfigCallerSession) WhitelistedModuleType(arg0 common.Address) (*big.Int, error) {
	return _Config.Contract.WhitelistedModuleType(&_Config.CallOpts, arg0)
}

// AddRecoverySigner is a paid mutator transaction binding the contract method 0x32b0cc32.
//
// Solidity: function addRecoverySigner(address signer) returns()
func (_Config *ConfigTransactor) AddRecoverySigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "addRecoverySigner", signer)
}

// AddRecoverySigner is a paid mutator transaction binding the contract method 0x32b0cc32.
//
// Solidity: function addRecoverySigner(address signer) returns()
func (_Config *ConfigSession) AddRecoverySigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.AddRecoverySigner(&_Config.TransactOpts, signer)
}

// AddRecoverySigner is a paid mutator transaction binding the contract method 0x32b0cc32.
//
// Solidity: function addRecoverySigner(address signer) returns()
func (_Config *ConfigTransactorSession) AddRecoverySigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.AddRecoverySigner(&_Config.TransactOpts, signer)
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

// AddWhitelistedModules is a paid mutator transaction binding the contract method 0xd4c4c2b0.
//
// Solidity: function addWhitelistedModules(address[] modules, uint256[] moduleTypes) returns()
func (_Config *ConfigTransactor) AddWhitelistedModules(opts *bind.TransactOpts, modules []common.Address, moduleTypes []*big.Int) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "addWhitelistedModules", modules, moduleTypes)
}

// AddWhitelistedModules is a paid mutator transaction binding the contract method 0xd4c4c2b0.
//
// Solidity: function addWhitelistedModules(address[] modules, uint256[] moduleTypes) returns()
func (_Config *ConfigSession) AddWhitelistedModules(modules []common.Address, moduleTypes []*big.Int) (*types.Transaction, error) {
	return _Config.Contract.AddWhitelistedModules(&_Config.TransactOpts, modules, moduleTypes)
}

// AddWhitelistedModules is a paid mutator transaction binding the contract method 0xd4c4c2b0.
//
// Solidity: function addWhitelistedModules(address[] modules, uint256[] moduleTypes) returns()
func (_Config *ConfigTransactorSession) AddWhitelistedModules(modules []common.Address, moduleTypes []*big.Int) (*types.Transaction, error) {
	return _Config.Contract.AddWhitelistedModules(&_Config.TransactOpts, modules, moduleTypes)
}

// AddWhitelistedVerifier is a paid mutator transaction binding the contract method 0x81e36802.
//
// Solidity: function addWhitelistedVerifier(uint8 verifier) returns()
func (_Config *ConfigTransactor) AddWhitelistedVerifier(opts *bind.TransactOpts, verifier uint8) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "addWhitelistedVerifier", verifier)
}

// AddWhitelistedVerifier is a paid mutator transaction binding the contract method 0x81e36802.
//
// Solidity: function addWhitelistedVerifier(uint8 verifier) returns()
func (_Config *ConfigSession) AddWhitelistedVerifier(verifier uint8) (*types.Transaction, error) {
	return _Config.Contract.AddWhitelistedVerifier(&_Config.TransactOpts, verifier)
}

// AddWhitelistedVerifier is a paid mutator transaction binding the contract method 0x81e36802.
//
// Solidity: function addWhitelistedVerifier(uint8 verifier) returns()
func (_Config *ConfigTransactorSession) AddWhitelistedVerifier(verifier uint8) (*types.Transaction, error) {
	return _Config.Contract.AddWhitelistedVerifier(&_Config.TransactOpts, verifier)
}

// BatchResetWalletSigner is a paid mutator transaction binding the contract method 0x15821128.
//
// Solidity: function batchResetWalletSigner(address[] wallets, address[] signers) returns()
func (_Config *ConfigTransactor) BatchResetWalletSigner(opts *bind.TransactOpts, wallets []common.Address, signers []common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "batchResetWalletSigner", wallets, signers)
}

// BatchResetWalletSigner is a paid mutator transaction binding the contract method 0x15821128.
//
// Solidity: function batchResetWalletSigner(address[] wallets, address[] signers) returns()
func (_Config *ConfigSession) BatchResetWalletSigner(wallets []common.Address, signers []common.Address) (*types.Transaction, error) {
	return _Config.Contract.BatchResetWalletSigner(&_Config.TransactOpts, wallets, signers)
}

// BatchResetWalletSigner is a paid mutator transaction binding the contract method 0x15821128.
//
// Solidity: function batchResetWalletSigner(address[] wallets, address[] signers) returns()
func (_Config *ConfigTransactorSession) BatchResetWalletSigner(wallets []common.Address, signers []common.Address) (*types.Transaction, error) {
	return _Config.Contract.BatchResetWalletSigner(&_Config.TransactOpts, wallets, signers)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address _initialOwner) returns()
func (_Config *ConfigTransactor) Initialize(opts *bind.TransactOpts, arg0 common.Address, _initialOwner common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "initialize", arg0, _initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address _initialOwner) returns()
func (_Config *ConfigSession) Initialize(arg0 common.Address, _initialOwner common.Address) (*types.Transaction, error) {
	return _Config.Contract.Initialize(&_Config.TransactOpts, arg0, _initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address _initialOwner) returns()
func (_Config *ConfigTransactorSession) Initialize(arg0 common.Address, _initialOwner common.Address) (*types.Transaction, error) {
	return _Config.Contract.Initialize(&_Config.TransactOpts, arg0, _initialOwner)
}

// InitializeWalletSigner is a paid mutator transaction binding the contract method 0x4b04ad51.
//
// Solidity: function initializeWalletSigner(address signer) returns()
func (_Config *ConfigTransactor) InitializeWalletSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "initializeWalletSigner", signer)
}

// InitializeWalletSigner is a paid mutator transaction binding the contract method 0x4b04ad51.
//
// Solidity: function initializeWalletSigner(address signer) returns()
func (_Config *ConfigSession) InitializeWalletSigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.InitializeWalletSigner(&_Config.TransactOpts, signer)
}

// InitializeWalletSigner is a paid mutator transaction binding the contract method 0x4b04ad51.
//
// Solidity: function initializeWalletSigner(address signer) returns()
func (_Config *ConfigTransactorSession) InitializeWalletSigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.InitializeWalletSigner(&_Config.TransactOpts, signer)
}

// RemoveRecoverySigner is a paid mutator transaction binding the contract method 0x562d2e27.
//
// Solidity: function removeRecoverySigner(address signer) returns()
func (_Config *ConfigTransactor) RemoveRecoverySigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "removeRecoverySigner", signer)
}

// RemoveRecoverySigner is a paid mutator transaction binding the contract method 0x562d2e27.
//
// Solidity: function removeRecoverySigner(address signer) returns()
func (_Config *ConfigSession) RemoveRecoverySigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.RemoveRecoverySigner(&_Config.TransactOpts, signer)
}

// RemoveRecoverySigner is a paid mutator transaction binding the contract method 0x562d2e27.
//
// Solidity: function removeRecoverySigner(address signer) returns()
func (_Config *ConfigTransactorSession) RemoveRecoverySigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.RemoveRecoverySigner(&_Config.TransactOpts, signer)
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

// RemoveWhitelistedModules is a paid mutator transaction binding the contract method 0xb8076742.
//
// Solidity: function removeWhitelistedModules(address[] modules) returns()
func (_Config *ConfigTransactor) RemoveWhitelistedModules(opts *bind.TransactOpts, modules []common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "removeWhitelistedModules", modules)
}

// RemoveWhitelistedModules is a paid mutator transaction binding the contract method 0xb8076742.
//
// Solidity: function removeWhitelistedModules(address[] modules) returns()
func (_Config *ConfigSession) RemoveWhitelistedModules(modules []common.Address) (*types.Transaction, error) {
	return _Config.Contract.RemoveWhitelistedModules(&_Config.TransactOpts, modules)
}

// RemoveWhitelistedModules is a paid mutator transaction binding the contract method 0xb8076742.
//
// Solidity: function removeWhitelistedModules(address[] modules) returns()
func (_Config *ConfigTransactorSession) RemoveWhitelistedModules(modules []common.Address) (*types.Transaction, error) {
	return _Config.Contract.RemoveWhitelistedModules(&_Config.TransactOpts, modules)
}

// RemoveWhitelistedVerifier is a paid mutator transaction binding the contract method 0xc22305cf.
//
// Solidity: function removeWhitelistedVerifier(uint8 verifier) returns()
func (_Config *ConfigTransactor) RemoveWhitelistedVerifier(opts *bind.TransactOpts, verifier uint8) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "removeWhitelistedVerifier", verifier)
}

// RemoveWhitelistedVerifier is a paid mutator transaction binding the contract method 0xc22305cf.
//
// Solidity: function removeWhitelistedVerifier(uint8 verifier) returns()
func (_Config *ConfigSession) RemoveWhitelistedVerifier(verifier uint8) (*types.Transaction, error) {
	return _Config.Contract.RemoveWhitelistedVerifier(&_Config.TransactOpts, verifier)
}

// RemoveWhitelistedVerifier is a paid mutator transaction binding the contract method 0xc22305cf.
//
// Solidity: function removeWhitelistedVerifier(uint8 verifier) returns()
func (_Config *ConfigTransactorSession) RemoveWhitelistedVerifier(verifier uint8) (*types.Transaction, error) {
	return _Config.Contract.RemoveWhitelistedVerifier(&_Config.TransactOpts, verifier)
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

// SetFactorySigner is a paid mutator transaction binding the contract method 0xa884c153.
//
// Solidity: function setFactorySigner(address signer) returns()
func (_Config *ConfigTransactor) SetFactorySigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setFactorySigner", signer)
}

// SetFactorySigner is a paid mutator transaction binding the contract method 0xa884c153.
//
// Solidity: function setFactorySigner(address signer) returns()
func (_Config *ConfigSession) SetFactorySigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetFactorySigner(&_Config.TransactOpts, signer)
}

// SetFactorySigner is a paid mutator transaction binding the contract method 0xa884c153.
//
// Solidity: function setFactorySigner(address signer) returns()
func (_Config *ConfigTransactorSession) SetFactorySigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetFactorySigner(&_Config.TransactOpts, signer)
}

// SetPaySigner is a paid mutator transaction binding the contract method 0xeb6810be.
//
// Solidity: function setPaySigner(address signer) returns()
func (_Config *ConfigTransactor) SetPaySigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setPaySigner", signer)
}

// SetPaySigner is a paid mutator transaction binding the contract method 0xeb6810be.
//
// Solidity: function setPaySigner(address signer) returns()
func (_Config *ConfigSession) SetPaySigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetPaySigner(&_Config.TransactOpts, signer)
}

// SetPaySigner is a paid mutator transaction binding the contract method 0xeb6810be.
//
// Solidity: function setPaySigner(address signer) returns()
func (_Config *ConfigTransactorSession) SetPaySigner(signer common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetPaySigner(&_Config.TransactOpts, signer)
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

// ConfigFactorySignerUpdatedIterator is returned from FilterFactorySignerUpdated and is used to iterate over the raw logs and unpacked data for FactorySignerUpdated events raised by the Config contract.
type ConfigFactorySignerUpdatedIterator struct {
	Event *ConfigFactorySignerUpdated // Event containing the contract specifics and raw log

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
func (it *ConfigFactorySignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigFactorySignerUpdated)
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
		it.Event = new(ConfigFactorySignerUpdated)
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
func (it *ConfigFactorySignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigFactorySignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigFactorySignerUpdated represents a FactorySignerUpdated event raised by the Config contract.
type ConfigFactorySignerUpdated struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFactorySignerUpdated is a free log retrieval operation binding the contract event 0x789d9446447ccefd52005113ddeccb6cfdcafd57a9f5229050f0b6c3dcd10f79.
//
// Solidity: event FactorySignerUpdated(address signer)
func (_Config *ConfigFilterer) FilterFactorySignerUpdated(opts *bind.FilterOpts) (*ConfigFactorySignerUpdatedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "FactorySignerUpdated")
	if err != nil {
		return nil, err
	}
	return &ConfigFactorySignerUpdatedIterator{contract: _Config.contract, event: "FactorySignerUpdated", logs: logs, sub: sub}, nil
}

// WatchFactorySignerUpdated is a free log subscription operation binding the contract event 0x789d9446447ccefd52005113ddeccb6cfdcafd57a9f5229050f0b6c3dcd10f79.
//
// Solidity: event FactorySignerUpdated(address signer)
func (_Config *ConfigFilterer) WatchFactorySignerUpdated(opts *bind.WatchOpts, sink chan<- *ConfigFactorySignerUpdated) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "FactorySignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigFactorySignerUpdated)
				if err := _Config.contract.UnpackLog(event, "FactorySignerUpdated", log); err != nil {
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

// ParseFactorySignerUpdated is a log parse operation binding the contract event 0x789d9446447ccefd52005113ddeccb6cfdcafd57a9f5229050f0b6c3dcd10f79.
//
// Solidity: event FactorySignerUpdated(address signer)
func (_Config *ConfigFilterer) ParseFactorySignerUpdated(log types.Log) (*ConfigFactorySignerUpdated, error) {
	event := new(ConfigFactorySignerUpdated)
	if err := _Config.contract.UnpackLog(event, "FactorySignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// ConfigPaySignerUpdatedIterator is returned from FilterPaySignerUpdated and is used to iterate over the raw logs and unpacked data for PaySignerUpdated events raised by the Config contract.
type ConfigPaySignerUpdatedIterator struct {
	Event *ConfigPaySignerUpdated // Event containing the contract specifics and raw log

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
func (it *ConfigPaySignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigPaySignerUpdated)
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
		it.Event = new(ConfigPaySignerUpdated)
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
func (it *ConfigPaySignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigPaySignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigPaySignerUpdated represents a PaySignerUpdated event raised by the Config contract.
type ConfigPaySignerUpdated struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaySignerUpdated is a free log retrieval operation binding the contract event 0x322933b0e9c97ef2d23835a0a1835c057a9601a32b8227690193a42c57bacf3d.
//
// Solidity: event PaySignerUpdated(address signer)
func (_Config *ConfigFilterer) FilterPaySignerUpdated(opts *bind.FilterOpts) (*ConfigPaySignerUpdatedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "PaySignerUpdated")
	if err != nil {
		return nil, err
	}
	return &ConfigPaySignerUpdatedIterator{contract: _Config.contract, event: "PaySignerUpdated", logs: logs, sub: sub}, nil
}

// WatchPaySignerUpdated is a free log subscription operation binding the contract event 0x322933b0e9c97ef2d23835a0a1835c057a9601a32b8227690193a42c57bacf3d.
//
// Solidity: event PaySignerUpdated(address signer)
func (_Config *ConfigFilterer) WatchPaySignerUpdated(opts *bind.WatchOpts, sink chan<- *ConfigPaySignerUpdated) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "PaySignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigPaySignerUpdated)
				if err := _Config.contract.UnpackLog(event, "PaySignerUpdated", log); err != nil {
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

// ParsePaySignerUpdated is a log parse operation binding the contract event 0x322933b0e9c97ef2d23835a0a1835c057a9601a32b8227690193a42c57bacf3d.
//
// Solidity: event PaySignerUpdated(address signer)
func (_Config *ConfigFilterer) ParsePaySignerUpdated(log types.Log) (*ConfigPaySignerUpdated, error) {
	event := new(ConfigPaySignerUpdated)
	if err := _Config.contract.UnpackLog(event, "PaySignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigRecoverySignerAddedIterator is returned from FilterRecoverySignerAdded and is used to iterate over the raw logs and unpacked data for RecoverySignerAdded events raised by the Config contract.
type ConfigRecoverySignerAddedIterator struct {
	Event *ConfigRecoverySignerAdded // Event containing the contract specifics and raw log

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
func (it *ConfigRecoverySignerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigRecoverySignerAdded)
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
		it.Event = new(ConfigRecoverySignerAdded)
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
func (it *ConfigRecoverySignerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigRecoverySignerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigRecoverySignerAdded represents a RecoverySignerAdded event raised by the Config contract.
type ConfigRecoverySignerAdded struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverySignerAdded is a free log retrieval operation binding the contract event 0x22dbed02a51e94f91fc3a6d1d2014ed6bc570834e7f3536162d626c429d90e54.
//
// Solidity: event RecoverySignerAdded(address signer)
func (_Config *ConfigFilterer) FilterRecoverySignerAdded(opts *bind.FilterOpts) (*ConfigRecoverySignerAddedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "RecoverySignerAdded")
	if err != nil {
		return nil, err
	}
	return &ConfigRecoverySignerAddedIterator{contract: _Config.contract, event: "RecoverySignerAdded", logs: logs, sub: sub}, nil
}

// WatchRecoverySignerAdded is a free log subscription operation binding the contract event 0x22dbed02a51e94f91fc3a6d1d2014ed6bc570834e7f3536162d626c429d90e54.
//
// Solidity: event RecoverySignerAdded(address signer)
func (_Config *ConfigFilterer) WatchRecoverySignerAdded(opts *bind.WatchOpts, sink chan<- *ConfigRecoverySignerAdded) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "RecoverySignerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigRecoverySignerAdded)
				if err := _Config.contract.UnpackLog(event, "RecoverySignerAdded", log); err != nil {
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

// ParseRecoverySignerAdded is a log parse operation binding the contract event 0x22dbed02a51e94f91fc3a6d1d2014ed6bc570834e7f3536162d626c429d90e54.
//
// Solidity: event RecoverySignerAdded(address signer)
func (_Config *ConfigFilterer) ParseRecoverySignerAdded(log types.Log) (*ConfigRecoverySignerAdded, error) {
	event := new(ConfigRecoverySignerAdded)
	if err := _Config.contract.UnpackLog(event, "RecoverySignerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigRecoverySignerRemovedIterator is returned from FilterRecoverySignerRemoved and is used to iterate over the raw logs and unpacked data for RecoverySignerRemoved events raised by the Config contract.
type ConfigRecoverySignerRemovedIterator struct {
	Event *ConfigRecoverySignerRemoved // Event containing the contract specifics and raw log

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
func (it *ConfigRecoverySignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigRecoverySignerRemoved)
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
		it.Event = new(ConfigRecoverySignerRemoved)
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
func (it *ConfigRecoverySignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigRecoverySignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigRecoverySignerRemoved represents a RecoverySignerRemoved event raised by the Config contract.
type ConfigRecoverySignerRemoved struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverySignerRemoved is a free log retrieval operation binding the contract event 0xadd78f873e935721c08cc4e66ec8de5f0bd111d6a9966eaa93b5e22b417d23a4.
//
// Solidity: event RecoverySignerRemoved(address signer)
func (_Config *ConfigFilterer) FilterRecoverySignerRemoved(opts *bind.FilterOpts) (*ConfigRecoverySignerRemovedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "RecoverySignerRemoved")
	if err != nil {
		return nil, err
	}
	return &ConfigRecoverySignerRemovedIterator{contract: _Config.contract, event: "RecoverySignerRemoved", logs: logs, sub: sub}, nil
}

// WatchRecoverySignerRemoved is a free log subscription operation binding the contract event 0xadd78f873e935721c08cc4e66ec8de5f0bd111d6a9966eaa93b5e22b417d23a4.
//
// Solidity: event RecoverySignerRemoved(address signer)
func (_Config *ConfigFilterer) WatchRecoverySignerRemoved(opts *bind.WatchOpts, sink chan<- *ConfigRecoverySignerRemoved) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "RecoverySignerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigRecoverySignerRemoved)
				if err := _Config.contract.UnpackLog(event, "RecoverySignerRemoved", log); err != nil {
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

// ParseRecoverySignerRemoved is a log parse operation binding the contract event 0xadd78f873e935721c08cc4e66ec8de5f0bd111d6a9966eaa93b5e22b417d23a4.
//
// Solidity: event RecoverySignerRemoved(address signer)
func (_Config *ConfigFilterer) ParseRecoverySignerRemoved(log types.Log) (*ConfigRecoverySignerRemoved, error) {
	event := new(ConfigRecoverySignerRemoved)
	if err := _Config.contract.UnpackLog(event, "RecoverySignerRemoved", log); err != nil {
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

// ConfigVerifierTypeAddedIterator is returned from FilterVerifierTypeAdded and is used to iterate over the raw logs and unpacked data for VerifierTypeAdded events raised by the Config contract.
type ConfigVerifierTypeAddedIterator struct {
	Event *ConfigVerifierTypeAdded // Event containing the contract specifics and raw log

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
func (it *ConfigVerifierTypeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigVerifierTypeAdded)
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
		it.Event = new(ConfigVerifierTypeAdded)
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
func (it *ConfigVerifierTypeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigVerifierTypeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigVerifierTypeAdded represents a VerifierTypeAdded event raised by the Config contract.
type ConfigVerifierTypeAdded struct {
	Verifier uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierTypeAdded is a free log retrieval operation binding the contract event 0xc007db510c05b52ed8aa21436b678ffa49a3f5de71cb7ba2c81363eda6f9a734.
//
// Solidity: event VerifierTypeAdded(uint8 verifier)
func (_Config *ConfigFilterer) FilterVerifierTypeAdded(opts *bind.FilterOpts) (*ConfigVerifierTypeAddedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "VerifierTypeAdded")
	if err != nil {
		return nil, err
	}
	return &ConfigVerifierTypeAddedIterator{contract: _Config.contract, event: "VerifierTypeAdded", logs: logs, sub: sub}, nil
}

// WatchVerifierTypeAdded is a free log subscription operation binding the contract event 0xc007db510c05b52ed8aa21436b678ffa49a3f5de71cb7ba2c81363eda6f9a734.
//
// Solidity: event VerifierTypeAdded(uint8 verifier)
func (_Config *ConfigFilterer) WatchVerifierTypeAdded(opts *bind.WatchOpts, sink chan<- *ConfigVerifierTypeAdded) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "VerifierTypeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigVerifierTypeAdded)
				if err := _Config.contract.UnpackLog(event, "VerifierTypeAdded", log); err != nil {
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

// ParseVerifierTypeAdded is a log parse operation binding the contract event 0xc007db510c05b52ed8aa21436b678ffa49a3f5de71cb7ba2c81363eda6f9a734.
//
// Solidity: event VerifierTypeAdded(uint8 verifier)
func (_Config *ConfigFilterer) ParseVerifierTypeAdded(log types.Log) (*ConfigVerifierTypeAdded, error) {
	event := new(ConfigVerifierTypeAdded)
	if err := _Config.contract.UnpackLog(event, "VerifierTypeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigVerifierTypeRemovedIterator is returned from FilterVerifierTypeRemoved and is used to iterate over the raw logs and unpacked data for VerifierTypeRemoved events raised by the Config contract.
type ConfigVerifierTypeRemovedIterator struct {
	Event *ConfigVerifierTypeRemoved // Event containing the contract specifics and raw log

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
func (it *ConfigVerifierTypeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigVerifierTypeRemoved)
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
		it.Event = new(ConfigVerifierTypeRemoved)
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
func (it *ConfigVerifierTypeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigVerifierTypeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigVerifierTypeRemoved represents a VerifierTypeRemoved event raised by the Config contract.
type ConfigVerifierTypeRemoved struct {
	Verifier uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierTypeRemoved is a free log retrieval operation binding the contract event 0xf7ad1892e8c359fbe99be323cb12129c63748ec6c859ad3dd7f22f427742ebab.
//
// Solidity: event VerifierTypeRemoved(uint8 verifier)
func (_Config *ConfigFilterer) FilterVerifierTypeRemoved(opts *bind.FilterOpts) (*ConfigVerifierTypeRemovedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "VerifierTypeRemoved")
	if err != nil {
		return nil, err
	}
	return &ConfigVerifierTypeRemovedIterator{contract: _Config.contract, event: "VerifierTypeRemoved", logs: logs, sub: sub}, nil
}

// WatchVerifierTypeRemoved is a free log subscription operation binding the contract event 0xf7ad1892e8c359fbe99be323cb12129c63748ec6c859ad3dd7f22f427742ebab.
//
// Solidity: event VerifierTypeRemoved(uint8 verifier)
func (_Config *ConfigFilterer) WatchVerifierTypeRemoved(opts *bind.WatchOpts, sink chan<- *ConfigVerifierTypeRemoved) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "VerifierTypeRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigVerifierTypeRemoved)
				if err := _Config.contract.UnpackLog(event, "VerifierTypeRemoved", log); err != nil {
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

// ParseVerifierTypeRemoved is a log parse operation binding the contract event 0xf7ad1892e8c359fbe99be323cb12129c63748ec6c859ad3dd7f22f427742ebab.
//
// Solidity: event VerifierTypeRemoved(uint8 verifier)
func (_Config *ConfigFilterer) ParseVerifierTypeRemoved(log types.Log) (*ConfigVerifierTypeRemoved, error) {
	event := new(ConfigVerifierTypeRemoved)
	if err := _Config.contract.UnpackLog(event, "VerifierTypeRemoved", log); err != nil {
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

// ConfigWhitelistModuleAddedIterator is returned from FilterWhitelistModuleAdded and is used to iterate over the raw logs and unpacked data for WhitelistModuleAdded events raised by the Config contract.
type ConfigWhitelistModuleAddedIterator struct {
	Event *ConfigWhitelistModuleAdded // Event containing the contract specifics and raw log

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
func (it *ConfigWhitelistModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigWhitelistModuleAdded)
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
		it.Event = new(ConfigWhitelistModuleAdded)
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
func (it *ConfigWhitelistModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigWhitelistModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigWhitelistModuleAdded represents a WhitelistModuleAdded event raised by the Config contract.
type ConfigWhitelistModuleAdded struct {
	Module     common.Address
	ModuleType *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWhitelistModuleAdded is a free log retrieval operation binding the contract event 0x51a45eb94bbc0121a4cdcf5f28257162a61221f0d6f8897adffe076012457865.
//
// Solidity: event WhitelistModuleAdded(address module, uint256 moduleType)
func (_Config *ConfigFilterer) FilterWhitelistModuleAdded(opts *bind.FilterOpts) (*ConfigWhitelistModuleAddedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "WhitelistModuleAdded")
	if err != nil {
		return nil, err
	}
	return &ConfigWhitelistModuleAddedIterator{contract: _Config.contract, event: "WhitelistModuleAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistModuleAdded is a free log subscription operation binding the contract event 0x51a45eb94bbc0121a4cdcf5f28257162a61221f0d6f8897adffe076012457865.
//
// Solidity: event WhitelistModuleAdded(address module, uint256 moduleType)
func (_Config *ConfigFilterer) WatchWhitelistModuleAdded(opts *bind.WatchOpts, sink chan<- *ConfigWhitelistModuleAdded) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "WhitelistModuleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigWhitelistModuleAdded)
				if err := _Config.contract.UnpackLog(event, "WhitelistModuleAdded", log); err != nil {
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

// ParseWhitelistModuleAdded is a log parse operation binding the contract event 0x51a45eb94bbc0121a4cdcf5f28257162a61221f0d6f8897adffe076012457865.
//
// Solidity: event WhitelistModuleAdded(address module, uint256 moduleType)
func (_Config *ConfigFilterer) ParseWhitelistModuleAdded(log types.Log) (*ConfigWhitelistModuleAdded, error) {
	event := new(ConfigWhitelistModuleAdded)
	if err := _Config.contract.UnpackLog(event, "WhitelistModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigWhitelistModuleRemovedIterator is returned from FilterWhitelistModuleRemoved and is used to iterate over the raw logs and unpacked data for WhitelistModuleRemoved events raised by the Config contract.
type ConfigWhitelistModuleRemovedIterator struct {
	Event *ConfigWhitelistModuleRemoved // Event containing the contract specifics and raw log

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
func (it *ConfigWhitelistModuleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigWhitelistModuleRemoved)
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
		it.Event = new(ConfigWhitelistModuleRemoved)
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
func (it *ConfigWhitelistModuleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigWhitelistModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigWhitelistModuleRemoved represents a WhitelistModuleRemoved event raised by the Config contract.
type ConfigWhitelistModuleRemoved struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWhitelistModuleRemoved is a free log retrieval operation binding the contract event 0x312c33106e0d74000c54d884caee742f0a33f8143ed96561cc1c80d3351b2cd2.
//
// Solidity: event WhitelistModuleRemoved(address module)
func (_Config *ConfigFilterer) FilterWhitelistModuleRemoved(opts *bind.FilterOpts) (*ConfigWhitelistModuleRemovedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "WhitelistModuleRemoved")
	if err != nil {
		return nil, err
	}
	return &ConfigWhitelistModuleRemovedIterator{contract: _Config.contract, event: "WhitelistModuleRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistModuleRemoved is a free log subscription operation binding the contract event 0x312c33106e0d74000c54d884caee742f0a33f8143ed96561cc1c80d3351b2cd2.
//
// Solidity: event WhitelistModuleRemoved(address module)
func (_Config *ConfigFilterer) WatchWhitelistModuleRemoved(opts *bind.WatchOpts, sink chan<- *ConfigWhitelistModuleRemoved) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "WhitelistModuleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigWhitelistModuleRemoved)
				if err := _Config.contract.UnpackLog(event, "WhitelistModuleRemoved", log); err != nil {
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

// ParseWhitelistModuleRemoved is a log parse operation binding the contract event 0x312c33106e0d74000c54d884caee742f0a33f8143ed96561cc1c80d3351b2cd2.
//
// Solidity: event WhitelistModuleRemoved(address module)
func (_Config *ConfigFilterer) ParseWhitelistModuleRemoved(log types.Log) (*ConfigWhitelistModuleRemoved, error) {
	event := new(ConfigWhitelistModuleRemoved)
	if err := _Config.contract.UnpackLog(event, "WhitelistModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
