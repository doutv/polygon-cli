// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blssignatureaggregator

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

// BLSSignatureAggregatorMetaData contains all meta data concerning the BLSSignatureAggregator contract.
var BLSSignatureAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BLS_DOMAIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"N\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"}],\"name\":\"aggregateSignatures\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"aggregatedSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getTrailingPublicKey\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"publicKey\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpPublicKey\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"publicKey\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"userOpToMessage\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validateSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"validateUserOpSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"sigForUserOp\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a034607457601f6116af38819003918201601f19168301916001600160401b03831184841017607957808492602094604052833981010312607457516001600160a01b038116810360745760805260405161161f908161009082396080518181816101270152818161078901526110e40152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c80629d9250146100c65780630396cb60146100c1578063062a422b146100bc57806322cdde4c146100b75780632dd81133146100b25780639b2004b5146100ad578063ae574a43146100a8578063b0d691fe146100a3578063b7620eb41461009e578063c9e525df146100995763d4fedb4d1461009457600080fd5b61081c565b6107f3565b6107b8565b610773565b61065e565b61063e565b6105a0565b61053e565b6101fb565b610106565b346101015760003660031901126101015760206040517fd84c4373167c517e9ccd66803f86d8a4f49e7e1315a7a73b516affea7428f82b8152f35b600080fd5b600060203660031901126101a45760043563ffffffff81168091036101a0577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316908290823b156101a057602460405180948193621cb65b60e51b8352600483015234905af1801561019b57610182575080f35b6001600160401b0381116101965760405280f35b61032f565b61087c565b5080fd5b80fd5b919082519283825260005b8481106101d3575050826000602080949584010152601f8019910116010190565b6020818301810151848301820152016101b2565b9060206101f89281815201906101a7565b90565b3461010157600319602036820112610101576004356001600160401b0381116101015761012081600401928236030112610101576102a1916020916102b561025361024b61010480940185610888565b8101906108ba565b926102ab61028661026c610267368561044b565b610d57565b9261028161027985610eef565b91369061044b565b610f18565b9160405197889663ebbdac9160e01b8852600488019061092c565b6044860190610954565b60c484019061092c565b8173__$39ff535e9424a736d9ed48ea59cbf93b79$__5af4801561019b576102e591600091610300575b5061097c565b6102fc6102f06109b9565b604051918291826101e7565b0390f35b610322915060203d602011610328575b61031a8183610397565b810190610914565b386102df565b503d610310565b634e487b7160e01b600052604160045260246000fd5b61012081019081106001600160401b0382111761019657604052565b604081019081106001600160401b0382111761019657604052565b608081019081106001600160401b0382111761019657604052565b90601f801991011681019081106001600160401b0382111761019657604052565b604051906103c582610345565b565b604051906103c582610361565b604051906103c58261037c565b35906001600160a01b038216820361010157565b81601f82011215610101578035906001600160401b0382116101965760405192610429601f8401601f191660200185610397565b8284526020838301011161010157816000926020809301838601378301015290565b919061012083820312610101576104606103b8565b9261046a816103e1565b8452602081013560208501526040810135916001600160401b039283811161010157816104989184016103f5565b6040860152606082013583811161010157816104b59184016103f5565b60608601526080820135608086015260a082013560a086015260c082013560c086015260e082013583811161010157816104f09184016103f5565b60e086015261010092838301359081116101015761050e92016103f5565b90830152565b602060031982011261010157600435906001600160401b038211610101576101f89160040161044b565b3461010157602061056861055136610514565b61056261055d82610d57565b610eef565b90611034565b604051908152f35b9181601f84011215610101578235916001600160401b038311610101576020808501948460051b01011161010157565b34610101576040366003190112610101576001600160401b03600435818111610101576105d1903690600401610570565b919060243592828411610101573660238501121561010157836004013592831161010157366024848601011161010157602461060e940191610c4b565b005b6080810192916000915b6004831061062757505050565b60019082518152602080910192019201919061061a565b34610101576102fc61065261026736610514565b60405191829182610610565b3461010157602080600319360112610101576004356001600160401b0381116101015761068f903690600401610570565b9161069983610a21565b926106a76040519485610397565b808452601f1992836106b883610a21565b018360005b82811061075d5750505060005b855181101561072557806106fa6106f26106e76001948787610b18565b610100810190610888565b810190610e54565b6107026103c7565b918252868201526107138289610b40565b5261071e8188610b40565b50016106ca565b6102fc846102f0876107368a611113565b80519084015160408051958601928352602083019190915284910103908101835282610397565b610765610e2b565b82828a0101520184906106bd565b34610101576000366003190112610101576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610101576020366003190112610101576004356001600160401b038111610101576106526107ee6102fc9236906004016103f5565b610e6a565b346101015760003660031901126101015760206040516000805160206115ca8339815191528152f35b346101015761084f61082d36610514565b6040805161083a81610361565b36903761084961055d82610d57565b90610f18565b60405190600090825b6002831061086557604084f35b600190825181526020809101920192019190610858565b6040513d6000823e3d90fd5b903590601e198136030182121561010157018035906001600160401b0382116101015760200191813603831361010157565b906040828203126101015780601f8301121561010157604051916108dd83610361565b82906040810192831161010157905b8282106108f95750505090565b81358152602091820191016108ec565b6101f890369061044b565b90816020910312610101575180151581036101015790565b6000915b6002831061093d57505050565b600190825181526020809101920192019190610930565b6000915b6004831061096557505050565b600190825181526020809101920192019190610958565b1561098357565b60405162461bcd60e51b815260206004820152600e60248201526d424c533a2077726f6e672073696760901b6044820152606490fd5b60405190602082018281106001600160401b038211176101965760405260008252565b156109e357565b60405162461bcd60e51b8152602060048201526016602482015275424c533a20696e76616c6964207369676e617475726560501b6044820152606490fd5b6001600160401b0381116101965760051b60200190565b60405190610a458261037c565b6080368337565b90610a5682610a21565b604090610a666040519182610397565b8381528093610a77601f1991610a21565b019160005b838110610a895750505050565b6020908251610a978161037c565b608036823782828601015201610a7c565b90610ab282610a21565b6040610ac16040519283610397565b8382528193610ad2601f1991610a21565b019160005b838110610ae45750505050565b6020908351610af281610361565b8436823782828501015201610ad7565b634e487b7160e01b600052603260045260246000fd5b9190811015610b3b5760051b8101359061011e1981360301821215610101570190565b610b02565b8051821015610b3b5760209160051b010190565b91608092610b6681608081019461092c565b60409360806040830152825180945260a08201936020809401916000905b828210610bd957505050506060818403910152602080855193848152019401926000905b838210610bb757505050505090565b9091929394838282610bcc6001948a5161092c565b0196019493920190610ba8565b909192969495868282610bef6001948c51610954565b9799970197960193920190610b84565b15610c0657565b60405162461bcd60e51b815260206004820152601e60248201527f424c533a2076616c69646174655369676e617475726573206661696c656400006044820152606490fd5b91610c5d919361024b604082146109dc565b610c6683610a4c565b91610c7084610aa8565b9360005b818110610cf05750505090602091610ca06040519485938493639141376360e01b855260048501610b54565b038173__$39ff535e9424a736d9ed48ea59cbf93b79$__5af4801561019b576103c591600091610cd1575b50610bff565b610cea915060203d6020116103285761031a8183610397565b38610ccb565b80610d3b610d09610d046001948688610b18565b610909565b610d1281610d57565b610d1c848a610b40565b52610d278389610b40565b50610849610d35848a610b40565b51610eef565b610d458289610b40565b52610d508188610b40565b5001610c74565b60405190610d648261037c565b608080923690376040810151805190919015610d8557506101f89150610e6a565b516040516370157dd760e11b815291508290829060049082906001600160a01b031661c350fa91821561019b57600092610dbe57505090565b803d8211610e24575b610dd18184610397565b8201918181840312610e205782601f82011215610e205760405193610df58561037c565b849282019384116101a45750905b828210610e105750505090565b8151815260209182019101610e03565b8380fd5b503d610dc7565b60405190604082018281106001600160401b038211176101965760405260006020838281520152565b9190826040910312610101576020823592013590565b90604051610e778161037c565b6080368237809280516080811115610eb15760609101605f198101518352603f198101516020840152601f19810151604084015251910152565b60405162461bcd60e51b81526020600482015260166024820152756461746120746f6f2073686f727420666f722073696760501b6044820152606490fd5b604051610f1281610f04602082019485610610565b03601f198101835282610397565b51902090565b90610f3490604092838051610f2c81610361565b369037611034565b610f95828051602093602082015260208152610f4f81610361565b81518093819263a850a90960e01b83527fd84c4373167c517e9ccd66803f86d8a4f49e7e1315a7a73b516affea7428f82b600484015284602484015260448301906101a7565b038173__$39ff535e9424a736d9ed48ea59cbf93b79$__5af492831561019b57600093610fc3575b50505090565b803d821161102d575b610fd68184610397565b82019181818403126110295782601f8201121561102957815194610ff986610361565b859282019384116101a45750905b82821061101a5750505050388080610fbd565b81518152908301908301611007565b8480fd5b503d610fcc565b610f1260018060a01b0382511660209283810151906040810151858151910120906060810151868151910120608082015160a08301519160e060c085015194015189815191012094604051968a880198895260408801526060870152608086015260a085015260c084015260e0830152610100908183015281526110b781610345565b51902060408051938401918252602082019490945230938101939093524660608401526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660808401528160a08401610f04565b9061111c610e2b565b825115610b3b57602090602084019182515192855115610b3b57516020015160019390845b87518610156111825791611174916001949361115d888b610b40565b5151918561116b8a8d610b40565b51015193611315565b959093019491929091611141565b9250935091945080801515908161126c575b5080611264575b1561122e5760009060016000805160206115ca833981519152825b6111e4575050506000805160206115ca83398151915292918184808281950980930987520990096020830152565b806112206000805160206115ca833981519152611208866112269598969798611596565b94611215828888096112e7565b9008949580946115b6565b90611308565b9190826111b6565b60405162461bcd60e51b815260206004820152600e60248201526d24b73b30b634b210373ab6b132b960911b6044820152606490fd5b50600161119b565b6000805160206115ca8339815191529150141538611194565b1561128c57565b60405162461bcd60e51b815260206004820152601e60248201527f557365206a6163446f75626c652066756e6374696f6e20696e737465616400006044820152606490fd5b634e487b7160e01b600052601160045260246000fd5b906000805160206115ca83398151915291820391821161130357565b6112d1565b9190820391821161130357565b94929091939480158061158e575b6115835781158061157b575b611572576113d761133e610a38565b966000805160206115ca83398151915280888009808a52880960208901526000805160206115ca8339815191526040890195600187528160608b0195600187526113866103d4565b9851900994858852519009906000805160206115ca8339815191528060208801968488528b519009916020604089019b848d52015190099060608701948286521491821592611567575b5050611285565b6113df610a38565b955183516113ec906112e7565b6000805160206115ca8339815191529108865251815161140b906112e7565b6000805160206115ca8339815191529108602086019281845286516000805160206115ca83398151915281800960408901918183526000805160206115ca8339815191529109926060890193808552611463906112e7565b90806000805160206115ca8339815191529109906000805160206115ca8339815191529108825182516000805160206115ca83398151915291096000805160206115ca833981519152906002096114b9906112e7565b6000805160206115ca83398151915291089451915190516000805160206115ca83398151915291096114ea856112e7565b6000805160206115ca83398151915291086000805160206115ca8339815191529109915190516000805160206115ca833981519152910961152a906112e7565b6000805160206115ca8339815191529108935190939260016000805160206115ca83398151915291096000805160206115ca833981519152910990565b1415905038806113d0565b91945050929190565b50851561132f565b509392506001919050565b508215611323565b81156115a0570490565b634e487b7160e01b600052601260045260246000fd5b818102929181159184041417156113035756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a264697066735822122015944ac6667c52e2e812cb901baf1db8255c204db2550bb9dc97fc9231d1c44f64736f6c63430008190033",
}

// BLSSignatureAggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use BLSSignatureAggregatorMetaData.ABI instead.
var BLSSignatureAggregatorABI = BLSSignatureAggregatorMetaData.ABI

// BLSSignatureAggregatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BLSSignatureAggregatorMetaData.Bin instead.
var BLSSignatureAggregatorBin = BLSSignatureAggregatorMetaData.Bin

// DeployBLSSignatureAggregator deploys a new Ethereum contract, binding an instance of BLSSignatureAggregator to it.
func DeployBLSSignatureAggregator(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *BLSSignatureAggregator, error) {
	parsed, err := BLSSignatureAggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BLSSignatureAggregatorBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BLSSignatureAggregator{BLSSignatureAggregatorCaller: BLSSignatureAggregatorCaller{contract: contract}, BLSSignatureAggregatorTransactor: BLSSignatureAggregatorTransactor{contract: contract}, BLSSignatureAggregatorFilterer: BLSSignatureAggregatorFilterer{contract: contract}}, nil
}

// BLSSignatureAggregator is an auto generated Go binding around an Ethereum contract.
type BLSSignatureAggregator struct {
	BLSSignatureAggregatorCaller     // Read-only binding to the contract
	BLSSignatureAggregatorTransactor // Write-only binding to the contract
	BLSSignatureAggregatorFilterer   // Log filterer for contract events
}

// BLSSignatureAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BLSSignatureAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSSignatureAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BLSSignatureAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSSignatureAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BLSSignatureAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSSignatureAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BLSSignatureAggregatorSession struct {
	Contract     *BLSSignatureAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BLSSignatureAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BLSSignatureAggregatorCallerSession struct {
	Contract *BLSSignatureAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// BLSSignatureAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BLSSignatureAggregatorTransactorSession struct {
	Contract     *BLSSignatureAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// BLSSignatureAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BLSSignatureAggregatorRaw struct {
	Contract *BLSSignatureAggregator // Generic contract binding to access the raw methods on
}

// BLSSignatureAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BLSSignatureAggregatorCallerRaw struct {
	Contract *BLSSignatureAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// BLSSignatureAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BLSSignatureAggregatorTransactorRaw struct {
	Contract *BLSSignatureAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBLSSignatureAggregator creates a new instance of BLSSignatureAggregator, bound to a specific deployed contract.
func NewBLSSignatureAggregator(address common.Address, backend bind.ContractBackend) (*BLSSignatureAggregator, error) {
	contract, err := bindBLSSignatureAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BLSSignatureAggregator{BLSSignatureAggregatorCaller: BLSSignatureAggregatorCaller{contract: contract}, BLSSignatureAggregatorTransactor: BLSSignatureAggregatorTransactor{contract: contract}, BLSSignatureAggregatorFilterer: BLSSignatureAggregatorFilterer{contract: contract}}, nil
}

// NewBLSSignatureAggregatorCaller creates a new read-only instance of BLSSignatureAggregator, bound to a specific deployed contract.
func NewBLSSignatureAggregatorCaller(address common.Address, caller bind.ContractCaller) (*BLSSignatureAggregatorCaller, error) {
	contract, err := bindBLSSignatureAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BLSSignatureAggregatorCaller{contract: contract}, nil
}

// NewBLSSignatureAggregatorTransactor creates a new write-only instance of BLSSignatureAggregator, bound to a specific deployed contract.
func NewBLSSignatureAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*BLSSignatureAggregatorTransactor, error) {
	contract, err := bindBLSSignatureAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BLSSignatureAggregatorTransactor{contract: contract}, nil
}

// NewBLSSignatureAggregatorFilterer creates a new log filterer instance of BLSSignatureAggregator, bound to a specific deployed contract.
func NewBLSSignatureAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*BLSSignatureAggregatorFilterer, error) {
	contract, err := bindBLSSignatureAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BLSSignatureAggregatorFilterer{contract: contract}, nil
}

// bindBLSSignatureAggregator binds a generic wrapper to an already deployed contract.
func bindBLSSignatureAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BLSSignatureAggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSSignatureAggregator *BLSSignatureAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSSignatureAggregator.Contract.BLSSignatureAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSSignatureAggregator *BLSSignatureAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSSignatureAggregator.Contract.BLSSignatureAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSSignatureAggregator *BLSSignatureAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSSignatureAggregator.Contract.BLSSignatureAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSSignatureAggregator *BLSSignatureAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSSignatureAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSSignatureAggregator *BLSSignatureAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSSignatureAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSSignatureAggregator *BLSSignatureAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSSignatureAggregator.Contract.contract.Transact(opts, method, params...)
}

// BLSDOMAIN is a free data retrieval call binding the contract method 0x009d9250.
//
// Solidity: function BLS_DOMAIN() view returns(bytes32)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCaller) BLSDOMAIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BLSSignatureAggregator.contract.Call(opts, &out, "BLS_DOMAIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLSDOMAIN is a free data retrieval call binding the contract method 0x009d9250.
//
// Solidity: function BLS_DOMAIN() view returns(bytes32)
func (_BLSSignatureAggregator *BLSSignatureAggregatorSession) BLSDOMAIN() ([32]byte, error) {
	return _BLSSignatureAggregator.Contract.BLSDOMAIN(&_BLSSignatureAggregator.CallOpts)
}

// BLSDOMAIN is a free data retrieval call binding the contract method 0x009d9250.
//
// Solidity: function BLS_DOMAIN() view returns(bytes32)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCallerSession) BLSDOMAIN() ([32]byte, error) {
	return _BLSSignatureAggregator.Contract.BLSDOMAIN(&_BLSSignatureAggregator.CallOpts)
}

// N is a free data retrieval call binding the contract method 0xc9e525df.
//
// Solidity: function N() view returns(uint256)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BLSSignatureAggregator.contract.Call(opts, &out, "N")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// N is a free data retrieval call binding the contract method 0xc9e525df.
//
// Solidity: function N() view returns(uint256)
func (_BLSSignatureAggregator *BLSSignatureAggregatorSession) N() (*big.Int, error) {
	return _BLSSignatureAggregator.Contract.N(&_BLSSignatureAggregator.CallOpts)
}

// N is a free data retrieval call binding the contract method 0xc9e525df.
//
// Solidity: function N() view returns(uint256)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCallerSession) N() (*big.Int, error) {
	return _BLSSignatureAggregator.Contract.N(&_BLSSignatureAggregator.CallOpts)
}

// AggregateSignatures is a free data retrieval call binding the contract method 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) pure returns(bytes aggregatedSignature)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCaller) AggregateSignatures(opts *bind.CallOpts, userOps []PackedUserOperation) ([]byte, error) {
	var out []interface{}
	err := _BLSSignatureAggregator.contract.Call(opts, &out, "aggregateSignatures", userOps)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AggregateSignatures is a free data retrieval call binding the contract method 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) pure returns(bytes aggregatedSignature)
func (_BLSSignatureAggregator *BLSSignatureAggregatorSession) AggregateSignatures(userOps []PackedUserOperation) ([]byte, error) {
	return _BLSSignatureAggregator.Contract.AggregateSignatures(&_BLSSignatureAggregator.CallOpts, userOps)
}

// AggregateSignatures is a free data retrieval call binding the contract method 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) pure returns(bytes aggregatedSignature)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCallerSession) AggregateSignatures(userOps []PackedUserOperation) ([]byte, error) {
	return _BLSSignatureAggregator.Contract.AggregateSignatures(&_BLSSignatureAggregator.CallOpts, userOps)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSSignatureAggregator.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_BLSSignatureAggregator *BLSSignatureAggregatorSession) EntryPoint() (common.Address, error) {
	return _BLSSignatureAggregator.Contract.EntryPoint(&_BLSSignatureAggregator.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCallerSession) EntryPoint() (common.Address, error) {
	return _BLSSignatureAggregator.Contract.EntryPoint(&_BLSSignatureAggregator.CallOpts)
}

// GetTrailingPublicKey is a free data retrieval call binding the contract method 0xb7620eb4.
//
// Solidity: function getTrailingPublicKey(bytes data) pure returns(uint256[4] publicKey)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCaller) GetTrailingPublicKey(opts *bind.CallOpts, data []byte) ([4]*big.Int, error) {
	var out []interface{}
	err := _BLSSignatureAggregator.contract.Call(opts, &out, "getTrailingPublicKey", data)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetTrailingPublicKey is a free data retrieval call binding the contract method 0xb7620eb4.
//
// Solidity: function getTrailingPublicKey(bytes data) pure returns(uint256[4] publicKey)
func (_BLSSignatureAggregator *BLSSignatureAggregatorSession) GetTrailingPublicKey(data []byte) ([4]*big.Int, error) {
	return _BLSSignatureAggregator.Contract.GetTrailingPublicKey(&_BLSSignatureAggregator.CallOpts, data)
}

// GetTrailingPublicKey is a free data retrieval call binding the contract method 0xb7620eb4.
//
// Solidity: function getTrailingPublicKey(bytes data) pure returns(uint256[4] publicKey)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCallerSession) GetTrailingPublicKey(data []byte) ([4]*big.Int, error) {
	return _BLSSignatureAggregator.Contract.GetTrailingPublicKey(&_BLSSignatureAggregator.CallOpts, data)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCaller) GetUserOpHash(opts *bind.CallOpts, userOp PackedUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _BLSSignatureAggregator.contract.Call(opts, &out, "getUserOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHash is a free data retrieval call binding the contract method 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_BLSSignatureAggregator *BLSSignatureAggregatorSession) GetUserOpHash(userOp PackedUserOperation) ([32]byte, error) {
	return _BLSSignatureAggregator.Contract.GetUserOpHash(&_BLSSignatureAggregator.CallOpts, userOp)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCallerSession) GetUserOpHash(userOp PackedUserOperation) ([32]byte, error) {
	return _BLSSignatureAggregator.Contract.GetUserOpHash(&_BLSSignatureAggregator.CallOpts, userOp)
}

// GetUserOpPublicKey is a free data retrieval call binding the contract method 0x9b2004b5.
//
// Solidity: function getUserOpPublicKey((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(uint256[4] publicKey)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCaller) GetUserOpPublicKey(opts *bind.CallOpts, userOp PackedUserOperation) ([4]*big.Int, error) {
	var out []interface{}
	err := _BLSSignatureAggregator.contract.Call(opts, &out, "getUserOpPublicKey", userOp)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetUserOpPublicKey is a free data retrieval call binding the contract method 0x9b2004b5.
//
// Solidity: function getUserOpPublicKey((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(uint256[4] publicKey)
func (_BLSSignatureAggregator *BLSSignatureAggregatorSession) GetUserOpPublicKey(userOp PackedUserOperation) ([4]*big.Int, error) {
	return _BLSSignatureAggregator.Contract.GetUserOpPublicKey(&_BLSSignatureAggregator.CallOpts, userOp)
}

// GetUserOpPublicKey is a free data retrieval call binding the contract method 0x9b2004b5.
//
// Solidity: function getUserOpPublicKey((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(uint256[4] publicKey)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCallerSession) GetUserOpPublicKey(userOp PackedUserOperation) ([4]*big.Int, error) {
	return _BLSSignatureAggregator.Contract.GetUserOpPublicKey(&_BLSSignatureAggregator.CallOpts, userOp)
}

// UserOpToMessage is a free data retrieval call binding the contract method 0xd4fedb4d.
//
// Solidity: function userOpToMessage((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(uint256[2])
func (_BLSSignatureAggregator *BLSSignatureAggregatorCaller) UserOpToMessage(opts *bind.CallOpts, userOp PackedUserOperation) ([2]*big.Int, error) {
	var out []interface{}
	err := _BLSSignatureAggregator.contract.Call(opts, &out, "userOpToMessage", userOp)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// UserOpToMessage is a free data retrieval call binding the contract method 0xd4fedb4d.
//
// Solidity: function userOpToMessage((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(uint256[2])
func (_BLSSignatureAggregator *BLSSignatureAggregatorSession) UserOpToMessage(userOp PackedUserOperation) ([2]*big.Int, error) {
	return _BLSSignatureAggregator.Contract.UserOpToMessage(&_BLSSignatureAggregator.CallOpts, userOp)
}

// UserOpToMessage is a free data retrieval call binding the contract method 0xd4fedb4d.
//
// Solidity: function userOpToMessage((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(uint256[2])
func (_BLSSignatureAggregator *BLSSignatureAggregatorCallerSession) UserOpToMessage(userOp PackedUserOperation) ([2]*big.Int, error) {
	return _BLSSignatureAggregator.Contract.UserOpToMessage(&_BLSSignatureAggregator.CallOpts, userOp)
}

// ValidateSignatures is a free data retrieval call binding the contract method 0x2dd81133.
//
// Solidity: function validateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, bytes signature) view returns()
func (_BLSSignatureAggregator *BLSSignatureAggregatorCaller) ValidateSignatures(opts *bind.CallOpts, userOps []PackedUserOperation, signature []byte) error {
	var out []interface{}
	err := _BLSSignatureAggregator.contract.Call(opts, &out, "validateSignatures", userOps, signature)

	if err != nil {
		return err
	}

	return err

}

// ValidateSignatures is a free data retrieval call binding the contract method 0x2dd81133.
//
// Solidity: function validateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, bytes signature) view returns()
func (_BLSSignatureAggregator *BLSSignatureAggregatorSession) ValidateSignatures(userOps []PackedUserOperation, signature []byte) error {
	return _BLSSignatureAggregator.Contract.ValidateSignatures(&_BLSSignatureAggregator.CallOpts, userOps, signature)
}

// ValidateSignatures is a free data retrieval call binding the contract method 0x2dd81133.
//
// Solidity: function validateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, bytes signature) view returns()
func (_BLSSignatureAggregator *BLSSignatureAggregatorCallerSession) ValidateSignatures(userOps []PackedUserOperation, signature []byte) error {
	return _BLSSignatureAggregator.Contract.ValidateSignatures(&_BLSSignatureAggregator.CallOpts, userOps, signature)
}

// ValidateUserOpSignature is a free data retrieval call binding the contract method 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes sigForUserOp)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCaller) ValidateUserOpSignature(opts *bind.CallOpts, userOp PackedUserOperation) ([]byte, error) {
	var out []interface{}
	err := _BLSSignatureAggregator.contract.Call(opts, &out, "validateUserOpSignature", userOp)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ValidateUserOpSignature is a free data retrieval call binding the contract method 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes sigForUserOp)
func (_BLSSignatureAggregator *BLSSignatureAggregatorSession) ValidateUserOpSignature(userOp PackedUserOperation) ([]byte, error) {
	return _BLSSignatureAggregator.Contract.ValidateUserOpSignature(&_BLSSignatureAggregator.CallOpts, userOp)
}

// ValidateUserOpSignature is a free data retrieval call binding the contract method 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes sigForUserOp)
func (_BLSSignatureAggregator *BLSSignatureAggregatorCallerSession) ValidateUserOpSignature(userOp PackedUserOperation) ([]byte, error) {
	return _BLSSignatureAggregator.Contract.ValidateUserOpSignature(&_BLSSignatureAggregator.CallOpts, userOp)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_BLSSignatureAggregator *BLSSignatureAggregatorTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _BLSSignatureAggregator.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_BLSSignatureAggregator *BLSSignatureAggregatorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _BLSSignatureAggregator.Contract.AddStake(&_BLSSignatureAggregator.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_BLSSignatureAggregator *BLSSignatureAggregatorTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _BLSSignatureAggregator.Contract.AddStake(&_BLSSignatureAggregator.TransactOpts, unstakeDelaySec)
}
