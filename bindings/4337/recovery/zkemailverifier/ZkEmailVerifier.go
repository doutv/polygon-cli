// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zkemailverifier

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

// ZkEmailVerifierMetaData contains all meta data concerning the ZkEmailVerifier contract.
var ZkEmailVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"split32To16\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_newPubKeyHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[17]\",\"name\":\"_pubSignals\",\"type\":\"uint256[17]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0806040523460c757306080527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c1660b857506001600160401b036002600160401b0319828216016074575b604051611eca90816100cd823960805181818160fe01526103510152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a13880806056565b63f92ee8a960e01b8152600490fd5b600080fdfe60406080815260048036101561001457600080fd5b6000803560e01c9283634f1ef286146100be57505050806352d1902d146100b9578063715018a6146100b45780638da5cb5b146100af5780639967af67146100aa578063ad3cb1cc146100a5578063c1940a36146100a0578063c4d66de81461009b578063ea237576146100965763f2fde38b1461009157600080fd5b61081e565b6107ed565b6106bb565b61059b565b61053a565b6104b8565b610417565b6103a9565b61033e565b80600319360112610222576100d1610226565b9160243567ffffffffffffffff811161022257906100f485939236908501610320565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116308114908115610206575b506101f65760209061013a611857565b84516352d1902d60e01b8152978891829089165afa809683976101c5575b5061018457505051634c9c8ce360e01b81526001600160a01b0390921690820190815281906020010390fd5b92909391600080516020611e5583398151915286036101ab57846101a88585611a8f565b80f35b51632a87526960e21b815290810185815281906020010390fd5b6101e891975060203d6020116101ef575b6101e081836102ab565b810190611848565b9587610158565b503d6101d6565b5050505163703e46dd60e11b8152fd5b905081600080516020611e55833981519152541614158861012a565b5080fd5b600435906001600160a01b038216820361023c57565b600080fd5b602435906001600160a01b038216820361023c57565b634e487b7160e01b600052604160045260246000fd5b6040810190811067ffffffffffffffff82111761028957604052565b610257565b610140810190811067ffffffffffffffff82111761028957604052565b90601f8019910116810190811067ffffffffffffffff82111761028957604052565b67ffffffffffffffff811161028957601f01601f191660200190565b9291926102f5826102cd565b9161030360405193846102ab565b82948184528183011161023c578281602093846000960137010152565b9080601f8301121561023c5781602061033b933591016102e9565b90565b3461023c57600036600319011261023c577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610397576020604051600080516020611e558339815191528152f35b60405163703e46dd60e11b8152600490fd5b3461023c57600080600319360112610414576103c3611857565b600080516020611e3583398151915280546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b3461023c57600036600319011261023c57600080516020611e35833981519152546040516001600160a01b039091168152602090f35b919082519283825260005b848110610479575050826000602080949584010152601f8019910116010190565b602081830181015184830182015201610458565b919260a09361033b96959215158452602084015260408301526060820152816080820152019061044d565b3461023c57608036600319011261023c576104d1610226565b6104d9610241565b6064359067ffffffffffffffff9081831161023c573660238401121561023c57826004013591821161023c57366024838501011161023c57610536936024610526940191604435916109ce565b916040959395519586958661048d565b0390f35b3461023c57600036600319011261023c5761053660405161055a8161026d565b60058152640352e302e360dc1b602082015260405191829160208352602083019061044d565b9060049160441161023c57565b9060c4916101041161023c57565b3461023c5761032036600319011261023c576105b636610580565b3660c41161023c576105c73661058d565b90366103241161023c576106b2916040519161038083016040526105ed61010435610ad3565b6105f961012435610ad3565b61060561014435610ad3565b61061161016435610ad3565b61061d61018435610ad3565b6106296101a435610ad3565b6106356101c435610ad3565b6106416101e435610ad3565b61064d61020435610ad3565b61065961022435610ad3565b61066561024435610ad3565b61067161026435610ad3565b61067d61028435610ad3565b6106896102a435610ad3565b6106956102c435610ad3565b6106a16102e435610ad3565b6106ad61030435610ad3565b6113fc565b60005260206000f35b3461023c57602036600319011261023c576106d4610226565b600080516020611e75833981519152549067ffffffffffffffff60ff8360401c16159216801590816107e5575b60011490816107db575b1590816107d2575b506107c057600080516020611e75833981519152805467ffffffffffffffff19166001179055610747908261079657611a79565b61074d57005b600080516020611e75833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b600080516020611e75833981519152805460ff60401b191668010000000000000000179055611a79565b60405163f92ee8a960e01b8152600490fd5b90501538610713565b303b15915061070b565b839150610701565b3461023c57602036600319011261023c57604080516001600160801b0319600435818116835260801b166020820152f35b3461023c57602036600319011261023c5761084761083a610226565b610842611857565b6117d4565b005b9080601f8301121561023c57604051916108628361026d565b82906040810192831161023c57905b82821061087e5750505090565b8151815260209182019101610871565b90916101008284031261023c576108a58383610849565b92604081605f8501121561023c57604051906108c08261026d565b8160c086019584871161023c57604001905b8682106108e7575050509061033b9193610849565b602083916108f58785610849565b8152019101906108d2565b9081602091031261023c5751801515810361023c5790565b6000915b6002831061092957505050565b60019082518152602080910192019201919061091c565b91949392909461095583610320810197610918565b600060408481015b600283106109a557505050509061097b6101009260c0830190610918565b01906000915b6011831061098e57505050565b600190825181526020809101920192019190610981565b602082826109b66001948851610918565b0194019201919261095d565b6040513d6000823e3d90fd5b929394909484019160c08584031261023c578435956020918287013596606081013596608082013567ffffffffffffffff9081811161023c57830188601f8201121561023c57888188610a23933591016102e9565b9760a084013591821161023c57886040938b8d8f8b99610a5c610a4e610a6998610a849a8d01610320565b8c808251830101910161088e565b9d919a909b013595611902565b6040516360ca051b60e11b8152958694859460048601610940565b0381305afa918215610ace57600092610aa1575b50509493929190565b610ac09250803d10610ac7575b610ab881836102ab565b810190610900565b3880610a98565b503d610aae565b6109c2565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000011115610afc57565b6000805260206000f35b604080517f2102dc0100b48329323c53280ae877da738da899f1a8f56586b9f3db415c16dd81527f1e6ca1ca9b627b68a4ed21432296be6e47d01e5286f1f802e73933fdf20d599860208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f1ad26e20ca34ed5728c869bd385109c9d88cff8f70d597e58b24cdc5589d979381527f1a66523292584ffbb4e973b11e188aa265bff1eb9d1558c0395f7bec4218b12060208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f0b55aab344f6d78726a0386f705bfbf629e33633a1789a5998efe885bc169dbf81527f169a6763ca5222ce7419afc3477d52d5e349f3d42f215750223017891926359d60208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f267c94cb005b72e7eb6930a764358cbf86c60f3e2f6157f9555db70e409a1e7981527f10247c59bb633d5bee15924d64a496c74f57ceabcb7fd3532250d1f35ae2720b60208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f195d487b09fccf451a06d28284c22339ff6ef986aa12d22ffe90581610af896681527f06f47d6692d9be4c4621a06b7daabccf08d25267c57e93c2c3fea3737e1a6e9460208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f22a4c17784e66a4c344cdfc3aa057c27fc9bebed331f8ba303319b3ba928de0d81527f15faba32b8caec4ebe23c31d173d7482673888e98eee9d6c58d3e710f4f8dde560208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f0bbe89bf4208ea0dc6ff2e46e45f6a8f37e425819879a8d5994b0bde9094c80d81527f10ce8e20eab9ebaf20bf9c0e37d0bf7221c9ccf6cad9daea9f3c443de3a6375560208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f14eb893ac09c3d3aee54418c040dbe88bf783d348bd3b5c4518c7bf1d8aec26281527f141e361dad6e277d65be816699cefb0ca2d923479778144ee00f29c13213d77f60208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f0a605a76e88bf31a11123fc3e369e0df3963c9c028634c361f9119b8de1b990481527f0738cbcecd43296bbc96f1edda74f47d60e5c39c0c78490d85575cd0654f97cf60208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f229a614ec1f2b488436bdff722d3181ab858756cbd0d8e6998ba7ee69402eea981527f1d1fe191fd8e3a87259f813bc6d129e33c7103fd2a5581c5a5a653ebd58560fa60208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f11e0dfaa09b5f711e02bdc2aeceae6087ec2f51a67360136620e4d362569c57f81527f2b175f9bd2ea8a198e48bf884f6122c6cd5f158f65c81f0a87dfc36a9c8ad08060208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517e38a2dd1123eed9c46fc32c6d7b7bc76501a796451b8ef21a610685bd3a3d0e81527f0565ecea2cabb990246ef3f76d2c15c29534aceb54268944d08f376f2d3d955360208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f2ed0df32674716396b7a5c3d9bdba4fea3696946998ff49fdbcf5281a169279381527f163dd0255b92a820b32d59e099f73e201bdb8e571790d34ac6ba28349d7fca1660208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f17261d4ffaffa7335907d4dbf6afa1645476e27e6d7c86aabb9ec21b2550a64881527f179e93baab96c3d37700ab1317864b5d71fef6e1dfbeb1aae743d8f0b5d76fff60208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f270cd23e87b08d7306de146775961bccded1850a4ec9ae5e09a51754ceed342581527f08990b55dad36cee232ba6496fb3200c96afda7dbc9bfdb4491889acec80710b60208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f051aef5ce2a503e20b64f08cbb8d40da692577307163775eea9cf47e96fc290081527f2c8dba9a0a180b5a808d12e7a5b3bd52ffaa313d443dfeba235e323e9d569a4c60208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b604080517f0db46c2a8077c2b9ee01ae292a3847bf9eb990cee9604de06518b7555071304281527f113a153506ad9fca85f4cf4e321d65d164535e506a77fb8662a86a09525352b760208201528082019384526107cf1991816060816007865a01fa15610afc57600660809260409585519052602085015160608401525a01fa15610afc57565b6020919282608082019485937f2b0fbae590d705323add300533843f3de53e877c016e08593f7c28fe75e890058452828401907f185f574e3ac0601a486d097729c8cefef767665b4d23fedb0642b4ed25ed7299825261145f6101043586610b06565b61146c6101243586610b8d565b6114796101443586610c14565b6114866101643586610c9b565b6114936101843586610d22565b6114a06101a43586610da9565b6114ad6101c43586610e30565b6114ba6101e43586610eb7565b6114c76102043586610f3e565b6114d46102243586610fc5565b6114e1610244358661104c565b6114ee61026435866110d3565b6114fb6102843586611159565b6115086102a435866111e0565b6115156102c43586611267565b6115226102e435866112ee565b61152f6103043586611375565b80358652837f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4791013581030660a085015260443560c085015260643560e085015260843561010085015260a4356101208501527f2d4d9aa7e302d9df41749d5507949d05dbea33fbb16c643b22f599a2be6df2e26101408501527f14bedd503c37ceb061d8ec60209fe345ce89830a19230301f076caff004d19266101608501527f0967032fcbf776d1afc985f88877f182d38480a653f2decaa9794cbc3bf3060c6101808501527f0e187847ad4c798374d0d6732bf501847dd68bc0e071241e0213bc7fc13db7ab6101a08501527f304cfbd1e08a704a99f5e847d93f8c3caafddec46b7a0d379da69a4d112346a76101c08501527f1739c1b1a457a8c7313123d24d2f9192f896b7c63eea05a9d57f06547ad0cec86101e08501528351610200850152516102208401527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26102408401527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6102608401527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b6102808401527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa6102a084015280356102c084015201356102e08201527f183072ee893a7fd87845d43df7aa78bcf0bc80acb9c57eea66f415795ac34a55610360610300927f1d0156ab95798e5dd2b89e834e57a3d846902599b894244708ecc68dec978609848201527f05ff1213d7efbc6706d81737de79aac39ebc4405c92f4dc8d893c1bdb1c6778d6103208201527f10bdcc6e8a53c19d893744f44fa672bdae7a5cf0b9bf8772fd29a9fdfde8c00e61034082015201528160086107cf195a01fa90511690565b6001600160a01b0390811690811561182f57600080516020611e3583398151915280546001600160a01b031981168417909155167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b604051631e4fbdf760e01b815260006004820152602490fd5b9081602091031261023c575190565b600080516020611e35833981519152546001600160a01b0316330361187857565b60405163118cdaa760e01b8152336004820152602490fd5b6040519061022080830183811067ffffffffffffffff82111761028957604052368337565b634e487b7160e01b600052603260045260246000fd5b8051156118d85760200190565b6118b5565b80518210156118d85760209160051b010190565b9060118110156118d85760051b0190565b94959092919397969761191c611916611890565b99611b61565b60005b60098110611a5957505061012089015261014088015261016087015260405160609390931b6bffffffffffffffffffffffff1916602084015260148352601f199261196b6034826102ab565b61197490611c24565b61197d906118cb565b5161018087015260405160609190911b6bffffffffffffffffffffffff19166020820152603483810182526119b290826102ab565b6119bb90611c24565b6119c4906118cb565b516101a08601526040516001600160801b031982811660208301526030808501835260809390931b1692916119f990826102ab565b611a0290611cc6565b611a0b906118cb565b516101c08601526040516001600160801b031992909216602083015260309081018252611a3890826102ab565b611a4190611cc6565b611a4a906118cb565b516101e0840152610200830152565b808b611a7282611a6b600195876118dd565b51926118f1565b520161191f565b611a8d90611a85611d60565b610842611d60565b565b90813b15611b1557600080516020611e5583398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a2805115611afa57611af791611d8f565b50565b505034611b0357565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b60405190611b438261026d565b6001825260203681840137565b9081518110156118d8570160200190565b604051611b6d8161028e565b6009808252610120366020840137600090815b818110611b8f57505050905090565b600092919291825b60f8811080611c1a575b15611bfc5786518210611bbd575b600860019101910190611b97565b926008600191611bf0611bea611be4611bd6878d611b50565b516001600160f81b03191690565b60f81c90565b60ff1690565b861b1794915050611baf565b50929190600191959495611c1082886118dd565b5201939293611b80565b5060ff8210611ba1565b90611c2d611b36565b91600091825b600180821015611cbe5760009081815b611c5f575b505090600191611c5882886118dd565b5201611c33565b60f896919680821080611cb4575b15611cab5790879187518410611c8d575b50918101969160080190611c43565b829194600891611c9d868b611b50565b51901c861b17949150611c7e565b50819650611c48565b5060148310611c6d565b509092505050565b90611ccf611b36565b91600091825b600180821015611cbe5760009081815b611d01575b505090600191611cfa82886118dd565b5201611cd5565b60f896919680821080611d56575b15611d4d5790879187518410611d2f575b50918101969160080190611ce5565b829194600891611d3f868b611b50565b51901c861b17949150611d20565b50819650611cea565b5060108310611d0f565b60ff600080516020611e758339815191525460401c1615611d7d57565b604051631afcd79f60e31b8152600490fd5b60008061033b93602081519101845af43d15611dcd573d91611db0836102cd565b92611dbe60405194856102ab565b83523d6000602085013e611dd1565b6060915b90611df85750805115611de657805190602001fd5b60405163d6bda27560e01b8152600490fd5b81511580611e2b575b611e09575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b15611e0156fe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212203f583093c9a686c6e4e2e5ee0552f3da8f6020590c7e8f449ab321e303b371f964736f6c63430008190033",
}

// ZkEmailVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkEmailVerifierMetaData.ABI instead.
var ZkEmailVerifierABI = ZkEmailVerifierMetaData.ABI

// ZkEmailVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZkEmailVerifierMetaData.Bin instead.
var ZkEmailVerifierBin = ZkEmailVerifierMetaData.Bin

// DeployZkEmailVerifier deploys a new Ethereum contract, binding an instance of ZkEmailVerifier to it.
func DeployZkEmailVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZkEmailVerifier, error) {
	parsed, err := ZkEmailVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZkEmailVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZkEmailVerifier{ZkEmailVerifierCaller: ZkEmailVerifierCaller{contract: contract}, ZkEmailVerifierTransactor: ZkEmailVerifierTransactor{contract: contract}, ZkEmailVerifierFilterer: ZkEmailVerifierFilterer{contract: contract}}, nil
}

// ZkEmailVerifier is an auto generated Go binding around an Ethereum contract.
type ZkEmailVerifier struct {
	ZkEmailVerifierCaller     // Read-only binding to the contract
	ZkEmailVerifierTransactor // Write-only binding to the contract
	ZkEmailVerifierFilterer   // Log filterer for contract events
}

// ZkEmailVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkEmailVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkEmailVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkEmailVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkEmailVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkEmailVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkEmailVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkEmailVerifierSession struct {
	Contract     *ZkEmailVerifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkEmailVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkEmailVerifierCallerSession struct {
	Contract *ZkEmailVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ZkEmailVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkEmailVerifierTransactorSession struct {
	Contract     *ZkEmailVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ZkEmailVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkEmailVerifierRaw struct {
	Contract *ZkEmailVerifier // Generic contract binding to access the raw methods on
}

// ZkEmailVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkEmailVerifierCallerRaw struct {
	Contract *ZkEmailVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ZkEmailVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkEmailVerifierTransactorRaw struct {
	Contract *ZkEmailVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkEmailVerifier creates a new instance of ZkEmailVerifier, bound to a specific deployed contract.
func NewZkEmailVerifier(address common.Address, backend bind.ContractBackend) (*ZkEmailVerifier, error) {
	contract, err := bindZkEmailVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZkEmailVerifier{ZkEmailVerifierCaller: ZkEmailVerifierCaller{contract: contract}, ZkEmailVerifierTransactor: ZkEmailVerifierTransactor{contract: contract}, ZkEmailVerifierFilterer: ZkEmailVerifierFilterer{contract: contract}}, nil
}

// NewZkEmailVerifierCaller creates a new read-only instance of ZkEmailVerifier, bound to a specific deployed contract.
func NewZkEmailVerifierCaller(address common.Address, caller bind.ContractCaller) (*ZkEmailVerifierCaller, error) {
	contract, err := bindZkEmailVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkEmailVerifierCaller{contract: contract}, nil
}

// NewZkEmailVerifierTransactor creates a new write-only instance of ZkEmailVerifier, bound to a specific deployed contract.
func NewZkEmailVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkEmailVerifierTransactor, error) {
	contract, err := bindZkEmailVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkEmailVerifierTransactor{contract: contract}, nil
}

// NewZkEmailVerifierFilterer creates a new log filterer instance of ZkEmailVerifier, bound to a specific deployed contract.
func NewZkEmailVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkEmailVerifierFilterer, error) {
	contract, err := bindZkEmailVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkEmailVerifierFilterer{contract: contract}, nil
}

// bindZkEmailVerifier binds a generic wrapper to an already deployed contract.
func bindZkEmailVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZkEmailVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkEmailVerifier *ZkEmailVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkEmailVerifier.Contract.ZkEmailVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkEmailVerifier *ZkEmailVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkEmailVerifier.Contract.ZkEmailVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkEmailVerifier *ZkEmailVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkEmailVerifier.Contract.ZkEmailVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkEmailVerifier *ZkEmailVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkEmailVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkEmailVerifier *ZkEmailVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkEmailVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkEmailVerifier *ZkEmailVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkEmailVerifier.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZkEmailVerifier *ZkEmailVerifierCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZkEmailVerifier.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZkEmailVerifier *ZkEmailVerifierSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ZkEmailVerifier.Contract.UPGRADEINTERFACEVERSION(&_ZkEmailVerifier.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZkEmailVerifier *ZkEmailVerifierCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ZkEmailVerifier.Contract.UPGRADEINTERFACEVERSION(&_ZkEmailVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZkEmailVerifier *ZkEmailVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkEmailVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZkEmailVerifier *ZkEmailVerifierSession) Owner() (common.Address, error) {
	return _ZkEmailVerifier.Contract.Owner(&_ZkEmailVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZkEmailVerifier *ZkEmailVerifierCallerSession) Owner() (common.Address, error) {
	return _ZkEmailVerifier.Contract.Owner(&_ZkEmailVerifier.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZkEmailVerifier *ZkEmailVerifierCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkEmailVerifier.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZkEmailVerifier *ZkEmailVerifierSession) ProxiableUUID() ([32]byte, error) {
	return _ZkEmailVerifier.Contract.ProxiableUUID(&_ZkEmailVerifier.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZkEmailVerifier *ZkEmailVerifierCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ZkEmailVerifier.Contract.ProxiableUUID(&_ZkEmailVerifier.CallOpts)
}

// Split32To16 is a free data retrieval call binding the contract method 0xea237576.
//
// Solidity: function split32To16(bytes32 _data) pure returns(bytes16, bytes16)
func (_ZkEmailVerifier *ZkEmailVerifierCaller) Split32To16(opts *bind.CallOpts, _data [32]byte) ([16]byte, [16]byte, error) {
	var out []interface{}
	err := _ZkEmailVerifier.contract.Call(opts, &out, "split32To16", _data)

	if err != nil {
		return *new([16]byte), *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)
	out1 := *abi.ConvertType(out[1], new([16]byte)).(*[16]byte)

	return out0, out1, err

}

// Split32To16 is a free data retrieval call binding the contract method 0xea237576.
//
// Solidity: function split32To16(bytes32 _data) pure returns(bytes16, bytes16)
func (_ZkEmailVerifier *ZkEmailVerifierSession) Split32To16(_data [32]byte) ([16]byte, [16]byte, error) {
	return _ZkEmailVerifier.Contract.Split32To16(&_ZkEmailVerifier.CallOpts, _data)
}

// Split32To16 is a free data retrieval call binding the contract method 0xea237576.
//
// Solidity: function split32To16(bytes32 _data) pure returns(bytes16, bytes16)
func (_ZkEmailVerifier *ZkEmailVerifierCallerSession) Split32To16(_data [32]byte) ([16]byte, [16]byte, error) {
	return _ZkEmailVerifier.Contract.Split32To16(&_ZkEmailVerifier.CallOpts, _data)
}

// Verify is a free data retrieval call binding the contract method 0x9967af67.
//
// Solidity: function verify(address _account, address _validator, bytes32 _newPubKeyHash, bytes _proof) view returns(bool, bytes32, bytes32, uint256, string)
func (_ZkEmailVerifier *ZkEmailVerifierCaller) Verify(opts *bind.CallOpts, _account common.Address, _validator common.Address, _newPubKeyHash [32]byte, _proof []byte) (bool, [32]byte, [32]byte, *big.Int, string, error) {
	var out []interface{}
	err := _ZkEmailVerifier.contract.Call(opts, &out, "verify", _account, _validator, _newPubKeyHash, _proof)

	if err != nil {
		return *new(bool), *new([32]byte), *new([32]byte), *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	out2 := *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)

	return out0, out1, out2, out3, out4, err

}

// Verify is a free data retrieval call binding the contract method 0x9967af67.
//
// Solidity: function verify(address _account, address _validator, bytes32 _newPubKeyHash, bytes _proof) view returns(bool, bytes32, bytes32, uint256, string)
func (_ZkEmailVerifier *ZkEmailVerifierSession) Verify(_account common.Address, _validator common.Address, _newPubKeyHash [32]byte, _proof []byte) (bool, [32]byte, [32]byte, *big.Int, string, error) {
	return _ZkEmailVerifier.Contract.Verify(&_ZkEmailVerifier.CallOpts, _account, _validator, _newPubKeyHash, _proof)
}

// Verify is a free data retrieval call binding the contract method 0x9967af67.
//
// Solidity: function verify(address _account, address _validator, bytes32 _newPubKeyHash, bytes _proof) view returns(bool, bytes32, bytes32, uint256, string)
func (_ZkEmailVerifier *ZkEmailVerifierCallerSession) Verify(_account common.Address, _validator common.Address, _newPubKeyHash [32]byte, _proof []byte) (bool, [32]byte, [32]byte, *big.Int, string, error) {
	return _ZkEmailVerifier.Contract.Verify(&_ZkEmailVerifier.CallOpts, _account, _validator, _newPubKeyHash, _proof)
}

// VerifyProof is a free data retrieval call binding the contract method 0xc1940a36.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[17] _pubSignals) view returns(bool)
func (_ZkEmailVerifier *ZkEmailVerifierCaller) VerifyProof(opts *bind.CallOpts, _pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [17]*big.Int) (bool, error) {
	var out []interface{}
	err := _ZkEmailVerifier.contract.Call(opts, &out, "verifyProof", _pA, _pB, _pC, _pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0xc1940a36.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[17] _pubSignals) view returns(bool)
func (_ZkEmailVerifier *ZkEmailVerifierSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [17]*big.Int) (bool, error) {
	return _ZkEmailVerifier.Contract.VerifyProof(&_ZkEmailVerifier.CallOpts, _pA, _pB, _pC, _pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0xc1940a36.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[17] _pubSignals) view returns(bool)
func (_ZkEmailVerifier *ZkEmailVerifierCallerSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [17]*big.Int) (bool, error) {
	return _ZkEmailVerifier.Contract.VerifyProof(&_ZkEmailVerifier.CallOpts, _pA, _pB, _pC, _pubSignals)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_ZkEmailVerifier *ZkEmailVerifierTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _ZkEmailVerifier.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_ZkEmailVerifier *ZkEmailVerifierSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _ZkEmailVerifier.Contract.Initialize(&_ZkEmailVerifier.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_ZkEmailVerifier *ZkEmailVerifierTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _ZkEmailVerifier.Contract.Initialize(&_ZkEmailVerifier.TransactOpts, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZkEmailVerifier *ZkEmailVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkEmailVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZkEmailVerifier *ZkEmailVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZkEmailVerifier.Contract.RenounceOwnership(&_ZkEmailVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZkEmailVerifier *ZkEmailVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZkEmailVerifier.Contract.RenounceOwnership(&_ZkEmailVerifier.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZkEmailVerifier *ZkEmailVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZkEmailVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZkEmailVerifier *ZkEmailVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZkEmailVerifier.Contract.TransferOwnership(&_ZkEmailVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZkEmailVerifier *ZkEmailVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZkEmailVerifier.Contract.TransferOwnership(&_ZkEmailVerifier.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZkEmailVerifier *ZkEmailVerifierTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZkEmailVerifier.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZkEmailVerifier *ZkEmailVerifierSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZkEmailVerifier.Contract.UpgradeToAndCall(&_ZkEmailVerifier.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZkEmailVerifier *ZkEmailVerifierTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZkEmailVerifier.Contract.UpgradeToAndCall(&_ZkEmailVerifier.TransactOpts, newImplementation, data)
}

// ZkEmailVerifierInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZkEmailVerifier contract.
type ZkEmailVerifierInitializedIterator struct {
	Event *ZkEmailVerifierInitialized // Event containing the contract specifics and raw log

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
func (it *ZkEmailVerifierInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEmailVerifierInitialized)
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
		it.Event = new(ZkEmailVerifierInitialized)
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
func (it *ZkEmailVerifierInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEmailVerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEmailVerifierInitialized represents a Initialized event raised by the ZkEmailVerifier contract.
type ZkEmailVerifierInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZkEmailVerifier *ZkEmailVerifierFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZkEmailVerifierInitializedIterator, error) {

	logs, sub, err := _ZkEmailVerifier.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZkEmailVerifierInitializedIterator{contract: _ZkEmailVerifier.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZkEmailVerifier *ZkEmailVerifierFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZkEmailVerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _ZkEmailVerifier.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEmailVerifierInitialized)
				if err := _ZkEmailVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ZkEmailVerifier *ZkEmailVerifierFilterer) ParseInitialized(log types.Log) (*ZkEmailVerifierInitialized, error) {
	event := new(ZkEmailVerifierInitialized)
	if err := _ZkEmailVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEmailVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZkEmailVerifier contract.
type ZkEmailVerifierOwnershipTransferredIterator struct {
	Event *ZkEmailVerifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZkEmailVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEmailVerifierOwnershipTransferred)
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
		it.Event = new(ZkEmailVerifierOwnershipTransferred)
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
func (it *ZkEmailVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEmailVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEmailVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the ZkEmailVerifier contract.
type ZkEmailVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZkEmailVerifier *ZkEmailVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZkEmailVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZkEmailVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZkEmailVerifierOwnershipTransferredIterator{contract: _ZkEmailVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZkEmailVerifier *ZkEmailVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZkEmailVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZkEmailVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEmailVerifierOwnershipTransferred)
				if err := _ZkEmailVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZkEmailVerifier *ZkEmailVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*ZkEmailVerifierOwnershipTransferred, error) {
	event := new(ZkEmailVerifierOwnershipTransferred)
	if err := _ZkEmailVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEmailVerifierUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ZkEmailVerifier contract.
type ZkEmailVerifierUpgradedIterator struct {
	Event *ZkEmailVerifierUpgraded // Event containing the contract specifics and raw log

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
func (it *ZkEmailVerifierUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEmailVerifierUpgraded)
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
		it.Event = new(ZkEmailVerifierUpgraded)
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
func (it *ZkEmailVerifierUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEmailVerifierUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEmailVerifierUpgraded represents a Upgraded event raised by the ZkEmailVerifier contract.
type ZkEmailVerifierUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZkEmailVerifier *ZkEmailVerifierFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ZkEmailVerifierUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ZkEmailVerifier.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ZkEmailVerifierUpgradedIterator{contract: _ZkEmailVerifier.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZkEmailVerifier *ZkEmailVerifierFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ZkEmailVerifierUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ZkEmailVerifier.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEmailVerifierUpgraded)
				if err := _ZkEmailVerifier.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ZkEmailVerifier *ZkEmailVerifierFilterer) ParseUpgraded(log types.Log) (*ZkEmailVerifierUpgraded, error) {
	event := new(ZkEmailVerifierUpgraded)
	if err := _ZkEmailVerifier.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
