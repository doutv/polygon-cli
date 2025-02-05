// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracletimelockcontroller

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

// OracleTimelockControllerMetaData contains all meta data concerning the OracleTimelockController contract.
var OracleTimelockControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minDelay\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"proposers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"executors\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDelay\",\"type\":\"uint256\"}],\"name\":\"TimelockInsufficientDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payloads\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"values\",\"type\":\"uint256\"}],\"name\":\"TimelockInvalidOperationLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TimelockUnauthorizedCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"predecessorId\",\"type\":\"bytes32\"}],\"name\":\"TimelockUnexecutedPredecessor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operationId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStates\",\"type\":\"bytes32\"}],\"name\":\"TimelockUnexpectedOperationState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateDelayUnsupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"CallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"CallSalt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"CallScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"Cancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"MinDelayChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CANCELLER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROPOSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getOperationState\",\"outputs\":[{\"internalType\":\"enumTimelockController.OperationState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"hashOperation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"hashOperationBatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperationDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperationPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperationReady\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"schedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"scheduleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"updateDelay\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234610151576119b48038038061001981610156565b92833981019060808183031261015157805160208201516001600160401b0390818111610151578461004c9185016101a5565b9360408401519182116101515761006a6060916100719386016101a5565b9301610191565b9061007b30610239565b506001600160a01b0391808316610141575b5060005b84518110156100cf57806100b2846100ab6001948961020f565b51166102b7565b506100c8846100c1838961020f565b5116610357565b5001610091565b50925060005b82518110156100fd57806100f6836100ef6001948761020f565b51166103f2565b50016100d5565b7f11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d560408580600255815190600082526020820152a1604051611506908161048e8239f35b61014a90610239565b503861008d565b600080fd5b6040519190601f01601f191682016001600160401b0381118382101761017b57604052565b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b038216820361015157565b81601f82011215610151578051916020916001600160401b03841161017b578360051b9083806101d6818501610156565b809781520192820101928311610151578301905b8282106101f8575050505090565b83809161020484610191565b8152019101906101ea565b80518210156102235760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b031660008181527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604081205490919060ff166102b357818052816020526040822081835260205260408220600160ff1982541617905533916000805160206119948339815191528180a4600190565b5090565b6001600160a01b031660008181527f3412d5605ac6cd444957cedb533e5dacad6378b4bc819ebe3652188a665066d560205260408120549091907fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc19060ff1661035257808352826020526040832082845260205260408320600160ff19825416179055600080516020611994833981519152339380a4600190565b505090565b6001600160a01b031660008181527fc3ad33e20b0c56a223ad5104fff154aa010f8715b9c981fd38fdc60a4d1a52fb60205260408120549091907ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f7839060ff1661035257808352826020526040832082845260205260408320600160ff19825416179055600080516020611994833981519152339380a4600190565b6001600160a01b031660008181527fdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d706960205260408120549091907fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e639060ff1661035257808352826020526040832082845260205260408320600160ff19825416179055600080516020611994833981519152339380a460019056fe60406080815260049081361015610020575b5050361561001e57600080fd5b005b600091823560e01c90816301d5062a14610abf57816301ffc9a714610a4c57816307bd026514610a2357838263134008d3146109795750816313bc9f2014610959578163150b7a0214610904578163248a9ca3146108da5781632ab0f529146108ba5781632f2ff15d1461089057816331d507501461087057816336568abe1461082a578163584b153e1461080157816364d62353146107e05781637958004c1461079d5781638065657f1461077b5781638f2a0bb0146105db5781638f61f4f5146105a057816391d148541461055a578163a217fddf1461053f578163b08e51c014610504578163b1c5f427146104d8578163bc197c8114610452578163c4d252f514610383578163d45c44351461035b578163d547741f14610316578163e38335e5146101d8578163f23a6e6114610180575063f27a0c9203610011573461017c578160031936011261017c576020906002549051908152f35b5080fd5b8284346101d55760a03660031901126101d55761019b610b9d565b506101a4610bb8565b50608435906001600160401b0382116101d557506020926101c791369101610c9b565b505163f23a6e6160e01b8152f35b80fd5b90506101e336610d12565b9098949591939296976000805160206114b18339815191528b528a602052858b208b805260205260ff868c20541615610308575b8383148015906102fe575b6102d0575061023a610241918a868a878b888f610feb565b98896112b7565b885b81811061025757896102548a611372565b80f35b80808a7fc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b588a8a6102c76102af8f988c6102a8828e6102a28f60019f61029d918591610f70565b610f96565b97610f70565b3595610faa565b906102bc8282878761131d565b8d5194859485610e32565b0390a301610243565b85516001624fcdef60e01b031981529081019283526020830185905260408301849052918291506060010390fd5b5084831415610222565b61031133611233565b610217565b91905034610357578060031936011261035757610353913561034e600161033b610bb8565b9383875286602052862001543390611286565b611410565b5080f35b8280fd5b9050346103575760203660031901126103575760209282913581526001845220549051908152f35b91905034610357576020366003190112610357578135917ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f7838085528460205282852033865260205260ff83862054161561043757506103e183610eb6565b1561041b5750829082825260016020528120557fbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb708280a280f35b826044925191635ead8eb560e01b835282015260066024820152fd5b604492519163e2517d3f60e01b835233908301526024820152fd5b8284346101d55760a03660031901126101d55761046d610b9d565b50610476610bb8565b506001600160401b039060443582811161017c576104979036908601610d73565b5060643582811161017c576104af9036908601610d73565b506084359182116101d557506020926104ca91369101610c9b565b505163bc197c8160e01b8152f35b50503461017c576020906104fd6104ee36610d12565b96959095949194939293610feb565b9051908152f35b50503461017c578160031936011261017c57602090517ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f7838152f35b50503461017c578160031936011261017c5751908152602090f35b9050346103575781600319360112610357578160209360ff9261057b610bb8565b903582528186528282206001600160a01b039091168252855220549151911615158152f35b50503461017c578160031936011261017c57602090517fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc18152f35b919050346103575760c0366003190112610357576001600160401b039082358281116107775761060e9036908501610ce2565b93602435848111610773576106269036908301610ce2565b9460443590811161076f5761063e9036908401610ce2565b606493919335906084359760a43593610656336111b0565b818b14801590610765575b610737575061067789848489858f8b908e610feb565b99610682858c611129565b8a8c5b8a8382106106cb578e838e838161069a578380f35b7f20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d03879160209151908152a28180808380f35b6001927f4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca8b8b61072c8f8c88978f92898f8f8f61071a9161071461029d86809461072199610f70565b9a610f70565b3598610faa565b915196879687610dfa565b0390a3018b90610685565b88516001624fcdef60e01b031981529081018b81526020810184905260408101929092529081906060010390fd5b50828b1415610661565b8780fd5b8680fd5b8480fd5b50503461017c576020906104fd61079136610bfb565b94939093929192610f1b565b83833461017c57602036600319011261017c576107ba8335610edf565b905191838210156107cd57602083838152f35b634e487b7160e01b815260218452602490fd5b8284346101d55760203660031901126101d55750516303871ef160e51b8152fd5b8284346101d55760203660031901126101d5575061082160209235610eb6565b90519015158152f35b83833461017c578060031936011261017c57610844610bb8565b90336001600160a01b038316036108615750610353919235611410565b5163334bd91960e11b81528390fd5b8284346101d55760203660031901126101d5575061082160209235610e9f565b9190503461035757806003193601126103575761035391356108b5600161033b610bb8565b611392565b8284346101d55760203660031901126101d5575061082160209235610e87565b90503461035757602036600319011261035757816020936001923581528085522001549051908152f35b8284346101d55760803660031901126101d55761091f610b9d565b50610928610bb8565b50606435906001600160401b0382116101d5575060209261094b91369101610c9b565b5051630a85bd0160e11b8152f35b8284346101d55760203660031901126101d5575061082160209235610e59565b6102546109f782610a0d7fc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b586109ee896109b136610bfb565b6000805160206114b18339815191528b9a9697939598929a528a602052828b208b805260205260ff838c20541615610a15575b8985858a8a610f1b565b998a98896112b7565b610a038383888861131d565b5194859485610e32565b0390a3611372565b610a1e33611233565b6109e4565b50503461017c578160031936011261017c57602090516000805160206114b18339815191528152f35b90503461035757602036600319011261035757359063ffffffff60e01b82168092036103575760209250630271189760e51b8214918215610a91575b50519015158152f35b909150637965db0b60e01b8114908115610aae575b509038610a88565b6301ffc9a760e01b14905038610aa6565b919050346103575760c036600319011261035757610adb610b9d565b90836024356044356001600160401b038111610357577f4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca95610b1f91369101610bce565b95909160643595610b606084359760a43590610b3a336111b0565b610b488a828d8a8989610f1b565b9a8b97610b55848a611129565b8a5196879687610dfa565b0390a381610b6c578380f35b7f20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d03879160209151908152a23880808380f35b600435906001600160a01b0382168203610bb357565b600080fd5b602435906001600160a01b0382168203610bb357565b9181601f84011215610bb3578235916001600160401b038311610bb35760208381860195010111610bb357565b60a0600319820112610bb3576004356001600160a01b0381168103610bb3579160243591604435906001600160401b038211610bb357610c3d91600401610bce565b90916064359060843590565b90601f801991011681019081106001600160401b03821117610c6a57604052565b634e487b7160e01b600052604160045260246000fd5b6001600160401b038111610c6a57601f01601f191660200190565b81601f82011215610bb357803590610cb282610c80565b92610cc06040519485610c49565b82845260208383010111610bb357816000926020809301838601378301015290565b9181601f84011215610bb3578235916001600160401b038311610bb3576020808501948460051b010111610bb357565b9060a0600319830112610bb3576001600160401b03600435818111610bb35783610d3e91600401610ce2565b93909392602435838111610bb35782610d5991600401610ce2565b93909392604435918211610bb357610c3d91600401610ce2565b81601f82011215610bb3578035916020916001600160401b038411610c6a578360051b9060405194610da785840187610c49565b85528380860192820101928311610bb3578301905b828210610dca575050505090565b81358152908301908301610dbc565b908060209392818452848401376000828201840152601f01601f1916010190565b929093610e28926080959897969860018060a01b03168552602085015260a0604085015260a0840191610dd9565b9460608201520152565b610e56949260609260018060a01b0316825260208201528160408201520191610dd9565b90565b610e6290610edf565b6004811015610e715760021490565b634e487b7160e01b600052602160045260246000fd5b610e9090610edf565b6004811015610e715760031490565b610ea890610edf565b6004811015610e7157151590565b610ebf90610edf565b6004811015610e715760018114908115610ed7575090565b600291501490565b60005260016020526040600020548015600014610efc5750600090565b60018103610f0a5750600390565b421015610f1657600190565b600290565b94610f51610f6a94959293604051968795602087019960018060a01b03168a52604087015260a0606087015260c0860191610dd9565b91608084015260a083015203601f198101835282610c49565b51902090565b9190811015610f805760051b0190565b634e487b7160e01b600052603260045260246000fd5b356001600160a01b0381168103610bb35790565b9190811015610f805760051b81013590601e1981360301821215610bb35701908135916001600160401b038311610bb3576020018236038113610bb3579190565b969294909695919560405196602091828901998060c08b0160a08d525260e08a01919060005b81811061110157505050601f19898203810160408b0152888252976001600160fb1b038111610bb3579089969495939897929160051b80928a830137019380888601878703606089015252604085019460408260051b82010195836000925b84841061109857505050505050610f6a9550608084015260a083015203908101835282610c49565b9193969850919398999496603f198282030184528935601e1984360301811215610bb35783018681019190356001600160401b038111610bb3578036038313610bb3576110ea88928392600195610dd9565b9b0194019401918b98969394919a9997959a611070565b90919283359060018060a01b038216809203610bb35790815285019285019190600101611011565b9061113382610e9f565b61119057600254808210611172575042019081421161115c576000526001602052604060002055565b634e487b7160e01b600052601160045260246000fd5b6044925060405191635433660960e01b835260048301526024820152fd5b604051635ead8eb560e01b81526004810183905260016024820152604490fd5b6001600160a01b031660008181527f3412d5605ac6cd444957cedb533e5dacad6378b4bc819ebe3652188a665066d560205260409020547fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc19060ff1615611215575050565b604492506040519163e2517d3f60e01b835260048301526024820152fd5b6001600160a01b031660008181527fdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d706960205260409020546000805160206114b18339815191529060ff1615611215575050565b80600052600060205260406000209160018060a01b0316918260005260205260ff6040600020541615611215575050565b6112c081610e59565b156112fe5750801515806112ee575b6112d65750565b6024906040519063121534c360e31b82526004820152fd5b506112f881610e87565b156112cf565b60449060405190635ead8eb560e01b8252600482015260046024820152fd5b61136793600093928493826040519384928337810185815203925af13d1561136a573d9061134a82610c80565b916113586040519384610c49565b82523d6000602084013e611485565b50565b606090611485565b61137b81610e59565b156112fe5760005260016020526001604060002055565b9060009180835282602052604083209160018060a01b03169182845260205260ff6040842054161560001461140b57808352826020526040832082845260205260408320600160ff198254161790557f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d339380a4600190565b505090565b9060009180835282602052604083209160018060a01b03169182845260205260ff60408420541660001461140b5780835282602052604083208284526020526040832060ff1981541690557ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b339380a4600190565b9091906114ae575080511561149c57805190602001fd5b60405163d6bda27560e01b8152600490fd5b56fed8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63a26469706673582212209d1063261723e30bc4568fa20758687d6d004b350c63187e6e04af746bbacdea64736f6c634300081900332f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d",
}

// OracleTimelockControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleTimelockControllerMetaData.ABI instead.
var OracleTimelockControllerABI = OracleTimelockControllerMetaData.ABI

// OracleTimelockControllerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleTimelockControllerMetaData.Bin instead.
var OracleTimelockControllerBin = OracleTimelockControllerMetaData.Bin

// DeployOracleTimelockController deploys a new Ethereum contract, binding an instance of OracleTimelockController to it.
func DeployOracleTimelockController(auth *bind.TransactOpts, backend bind.ContractBackend, minDelay *big.Int, proposers []common.Address, executors []common.Address, admin common.Address) (common.Address, *types.Transaction, *OracleTimelockController, error) {
	parsed, err := OracleTimelockControllerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleTimelockControllerBin), backend, minDelay, proposers, executors, admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OracleTimelockController{OracleTimelockControllerCaller: OracleTimelockControllerCaller{contract: contract}, OracleTimelockControllerTransactor: OracleTimelockControllerTransactor{contract: contract}, OracleTimelockControllerFilterer: OracleTimelockControllerFilterer{contract: contract}}, nil
}

// OracleTimelockController is an auto generated Go binding around an Ethereum contract.
type OracleTimelockController struct {
	OracleTimelockControllerCaller     // Read-only binding to the contract
	OracleTimelockControllerTransactor // Write-only binding to the contract
	OracleTimelockControllerFilterer   // Log filterer for contract events
}

// OracleTimelockControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleTimelockControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTimelockControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTimelockControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTimelockControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleTimelockControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTimelockControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleTimelockControllerSession struct {
	Contract     *OracleTimelockController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OracleTimelockControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleTimelockControllerCallerSession struct {
	Contract *OracleTimelockControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// OracleTimelockControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTimelockControllerTransactorSession struct {
	Contract     *OracleTimelockControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// OracleTimelockControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleTimelockControllerRaw struct {
	Contract *OracleTimelockController // Generic contract binding to access the raw methods on
}

// OracleTimelockControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleTimelockControllerCallerRaw struct {
	Contract *OracleTimelockControllerCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTimelockControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTimelockControllerTransactorRaw struct {
	Contract *OracleTimelockControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleTimelockController creates a new instance of OracleTimelockController, bound to a specific deployed contract.
func NewOracleTimelockController(address common.Address, backend bind.ContractBackend) (*OracleTimelockController, error) {
	contract, err := bindOracleTimelockController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleTimelockController{OracleTimelockControllerCaller: OracleTimelockControllerCaller{contract: contract}, OracleTimelockControllerTransactor: OracleTimelockControllerTransactor{contract: contract}, OracleTimelockControllerFilterer: OracleTimelockControllerFilterer{contract: contract}}, nil
}

// NewOracleTimelockControllerCaller creates a new read-only instance of OracleTimelockController, bound to a specific deployed contract.
func NewOracleTimelockControllerCaller(address common.Address, caller bind.ContractCaller) (*OracleTimelockControllerCaller, error) {
	contract, err := bindOracleTimelockController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTimelockControllerCaller{contract: contract}, nil
}

// NewOracleTimelockControllerTransactor creates a new write-only instance of OracleTimelockController, bound to a specific deployed contract.
func NewOracleTimelockControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTimelockControllerTransactor, error) {
	contract, err := bindOracleTimelockController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTimelockControllerTransactor{contract: contract}, nil
}

// NewOracleTimelockControllerFilterer creates a new log filterer instance of OracleTimelockController, bound to a specific deployed contract.
func NewOracleTimelockControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleTimelockControllerFilterer, error) {
	contract, err := bindOracleTimelockController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleTimelockControllerFilterer{contract: contract}, nil
}

// bindOracleTimelockController binds a generic wrapper to an already deployed contract.
func bindOracleTimelockController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleTimelockControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleTimelockController *OracleTimelockControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleTimelockController.Contract.OracleTimelockControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleTimelockController *OracleTimelockControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.OracleTimelockControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleTimelockController *OracleTimelockControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.OracleTimelockControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleTimelockController *OracleTimelockControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleTimelockController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleTimelockController *OracleTimelockControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleTimelockController *OracleTimelockControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.contract.Transact(opts, method, params...)
}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCaller) CANCELLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "CANCELLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerSession) CANCELLERROLE() ([32]byte, error) {
	return _OracleTimelockController.Contract.CANCELLERROLE(&_OracleTimelockController.CallOpts)
}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) CANCELLERROLE() ([32]byte, error) {
	return _OracleTimelockController.Contract.CANCELLERROLE(&_OracleTimelockController.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OracleTimelockController.Contract.DEFAULTADMINROLE(&_OracleTimelockController.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OracleTimelockController.Contract.DEFAULTADMINROLE(&_OracleTimelockController.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCaller) EXECUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "EXECUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerSession) EXECUTORROLE() ([32]byte, error) {
	return _OracleTimelockController.Contract.EXECUTORROLE(&_OracleTimelockController.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) EXECUTORROLE() ([32]byte, error) {
	return _OracleTimelockController.Contract.EXECUTORROLE(&_OracleTimelockController.CallOpts)
}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCaller) PROPOSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "PROPOSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerSession) PROPOSERROLE() ([32]byte, error) {
	return _OracleTimelockController.Contract.PROPOSERROLE(&_OracleTimelockController.CallOpts)
}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) PROPOSERROLE() ([32]byte, error) {
	return _OracleTimelockController.Contract.PROPOSERROLE(&_OracleTimelockController.CallOpts)
}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (_OracleTimelockController *OracleTimelockControllerCaller) GetMinDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "getMinDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (_OracleTimelockController *OracleTimelockControllerSession) GetMinDelay() (*big.Int, error) {
	return _OracleTimelockController.Contract.GetMinDelay(&_OracleTimelockController.CallOpts)
}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) GetMinDelay() (*big.Int, error) {
	return _OracleTimelockController.Contract.GetMinDelay(&_OracleTimelockController.CallOpts)
}

// GetOperationState is a free data retrieval call binding the contract method 0x7958004c.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (_OracleTimelockController *OracleTimelockControllerCaller) GetOperationState(opts *bind.CallOpts, id [32]byte) (uint8, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "getOperationState", id)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperationState is a free data retrieval call binding the contract method 0x7958004c.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (_OracleTimelockController *OracleTimelockControllerSession) GetOperationState(id [32]byte) (uint8, error) {
	return _OracleTimelockController.Contract.GetOperationState(&_OracleTimelockController.CallOpts, id)
}

// GetOperationState is a free data retrieval call binding the contract method 0x7958004c.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) GetOperationState(id [32]byte) (uint8, error) {
	return _OracleTimelockController.Contract.GetOperationState(&_OracleTimelockController.CallOpts, id)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OracleTimelockController.Contract.GetRoleAdmin(&_OracleTimelockController.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OracleTimelockController.Contract.GetRoleAdmin(&_OracleTimelockController.CallOpts, role)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (_OracleTimelockController *OracleTimelockControllerCaller) GetTimestamp(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "getTimestamp", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (_OracleTimelockController *OracleTimelockControllerSession) GetTimestamp(id [32]byte) (*big.Int, error) {
	return _OracleTimelockController.Contract.GetTimestamp(&_OracleTimelockController.CallOpts, id)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) GetTimestamp(id [32]byte) (*big.Int, error) {
	return _OracleTimelockController.Contract.GetTimestamp(&_OracleTimelockController.CallOpts, id)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OracleTimelockController.Contract.HasRole(&_OracleTimelockController.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OracleTimelockController.Contract.HasRole(&_OracleTimelockController.CallOpts, role, account)
}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCaller) HashOperation(opts *bind.CallOpts, target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "hashOperation", target, value, data, predecessor, salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerSession) HashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _OracleTimelockController.Contract.HashOperation(&_OracleTimelockController.CallOpts, target, value, data, predecessor, salt)
}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) HashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _OracleTimelockController.Contract.HashOperation(&_OracleTimelockController.CallOpts, target, value, data, predecessor, salt)
}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCaller) HashOperationBatch(opts *bind.CallOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "hashOperationBatch", targets, values, payloads, predecessor, salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerSession) HashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _OracleTimelockController.Contract.HashOperationBatch(&_OracleTimelockController.CallOpts, targets, values, payloads, predecessor, salt)
}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) HashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _OracleTimelockController.Contract.HashOperationBatch(&_OracleTimelockController.CallOpts, targets, values, payloads, predecessor, salt)
}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerCaller) IsOperation(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "isOperation", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerSession) IsOperation(id [32]byte) (bool, error) {
	return _OracleTimelockController.Contract.IsOperation(&_OracleTimelockController.CallOpts, id)
}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) IsOperation(id [32]byte) (bool, error) {
	return _OracleTimelockController.Contract.IsOperation(&_OracleTimelockController.CallOpts, id)
}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerCaller) IsOperationDone(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "isOperationDone", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerSession) IsOperationDone(id [32]byte) (bool, error) {
	return _OracleTimelockController.Contract.IsOperationDone(&_OracleTimelockController.CallOpts, id)
}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) IsOperationDone(id [32]byte) (bool, error) {
	return _OracleTimelockController.Contract.IsOperationDone(&_OracleTimelockController.CallOpts, id)
}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerCaller) IsOperationPending(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "isOperationPending", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerSession) IsOperationPending(id [32]byte) (bool, error) {
	return _OracleTimelockController.Contract.IsOperationPending(&_OracleTimelockController.CallOpts, id)
}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) IsOperationPending(id [32]byte) (bool, error) {
	return _OracleTimelockController.Contract.IsOperationPending(&_OracleTimelockController.CallOpts, id)
}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerCaller) IsOperationReady(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "isOperationReady", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerSession) IsOperationReady(id [32]byte) (bool, error) {
	return _OracleTimelockController.Contract.IsOperationReady(&_OracleTimelockController.CallOpts, id)
}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) IsOperationReady(id [32]byte) (bool, error) {
	return _OracleTimelockController.Contract.IsOperationReady(&_OracleTimelockController.CallOpts, id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OracleTimelockController.Contract.SupportsInterface(&_OracleTimelockController.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OracleTimelockController *OracleTimelockControllerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OracleTimelockController.Contract.SupportsInterface(&_OracleTimelockController.CallOpts, interfaceId)
}

// UpdateDelay is a free data retrieval call binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) pure returns()
func (_OracleTimelockController *OracleTimelockControllerCaller) UpdateDelay(opts *bind.CallOpts, newDelay *big.Int) error {
	var out []interface{}
	err := _OracleTimelockController.contract.Call(opts, &out, "updateDelay", newDelay)

	if err != nil {
		return err
	}

	return err

}

// UpdateDelay is a free data retrieval call binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) pure returns()
func (_OracleTimelockController *OracleTimelockControllerSession) UpdateDelay(newDelay *big.Int) error {
	return _OracleTimelockController.Contract.UpdateDelay(&_OracleTimelockController.CallOpts, newDelay)
}

// UpdateDelay is a free data retrieval call binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) pure returns()
func (_OracleTimelockController *OracleTimelockControllerCallerSession) UpdateDelay(newDelay *big.Int) error {
	return _OracleTimelockController.Contract.UpdateDelay(&_OracleTimelockController.CallOpts, newDelay)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_OracleTimelockController *OracleTimelockControllerTransactor) Cancel(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _OracleTimelockController.contract.Transact(opts, "cancel", id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_OracleTimelockController *OracleTimelockControllerSession) Cancel(id [32]byte) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.Cancel(&_OracleTimelockController.TransactOpts, id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_OracleTimelockController *OracleTimelockControllerTransactorSession) Cancel(id [32]byte) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.Cancel(&_OracleTimelockController.TransactOpts, id)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (_OracleTimelockController *OracleTimelockControllerTransactor) Execute(opts *bind.TransactOpts, target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _OracleTimelockController.contract.Transact(opts, "execute", target, value, payload, predecessor, salt)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (_OracleTimelockController *OracleTimelockControllerSession) Execute(target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.Execute(&_OracleTimelockController.TransactOpts, target, value, payload, predecessor, salt)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (_OracleTimelockController *OracleTimelockControllerTransactorSession) Execute(target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.Execute(&_OracleTimelockController.TransactOpts, target, value, payload, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_OracleTimelockController *OracleTimelockControllerTransactor) ExecuteBatch(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _OracleTimelockController.contract.Transact(opts, "executeBatch", targets, values, payloads, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_OracleTimelockController *OracleTimelockControllerSession) ExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.ExecuteBatch(&_OracleTimelockController.TransactOpts, targets, values, payloads, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_OracleTimelockController *OracleTimelockControllerTransactorSession) ExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.ExecuteBatch(&_OracleTimelockController.TransactOpts, targets, values, payloads, predecessor, salt)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OracleTimelockController *OracleTimelockControllerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OracleTimelockController.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OracleTimelockController *OracleTimelockControllerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.GrantRole(&_OracleTimelockController.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OracleTimelockController *OracleTimelockControllerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.GrantRole(&_OracleTimelockController.TransactOpts, role, account)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_OracleTimelockController *OracleTimelockControllerTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OracleTimelockController.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_OracleTimelockController *OracleTimelockControllerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.OnERC1155BatchReceived(&_OracleTimelockController.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_OracleTimelockController *OracleTimelockControllerTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.OnERC1155BatchReceived(&_OracleTimelockController.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_OracleTimelockController *OracleTimelockControllerTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OracleTimelockController.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_OracleTimelockController *OracleTimelockControllerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.OnERC1155Received(&_OracleTimelockController.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_OracleTimelockController *OracleTimelockControllerTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.OnERC1155Received(&_OracleTimelockController.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_OracleTimelockController *OracleTimelockControllerTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _OracleTimelockController.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_OracleTimelockController *OracleTimelockControllerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.OnERC721Received(&_OracleTimelockController.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_OracleTimelockController *OracleTimelockControllerTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.OnERC721Received(&_OracleTimelockController.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OracleTimelockController *OracleTimelockControllerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OracleTimelockController.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OracleTimelockController *OracleTimelockControllerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.RenounceRole(&_OracleTimelockController.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OracleTimelockController *OracleTimelockControllerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.RenounceRole(&_OracleTimelockController.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OracleTimelockController *OracleTimelockControllerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OracleTimelockController.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OracleTimelockController *OracleTimelockControllerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.RevokeRole(&_OracleTimelockController.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OracleTimelockController *OracleTimelockControllerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.RevokeRole(&_OracleTimelockController.TransactOpts, role, account)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_OracleTimelockController *OracleTimelockControllerTransactor) Schedule(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _OracleTimelockController.contract.Transact(opts, "schedule", target, value, data, predecessor, salt, delay)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_OracleTimelockController *OracleTimelockControllerSession) Schedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.Schedule(&_OracleTimelockController.TransactOpts, target, value, data, predecessor, salt, delay)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_OracleTimelockController *OracleTimelockControllerTransactorSession) Schedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.Schedule(&_OracleTimelockController.TransactOpts, target, value, data, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_OracleTimelockController *OracleTimelockControllerTransactor) ScheduleBatch(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _OracleTimelockController.contract.Transact(opts, "scheduleBatch", targets, values, payloads, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_OracleTimelockController *OracleTimelockControllerSession) ScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.ScheduleBatch(&_OracleTimelockController.TransactOpts, targets, values, payloads, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_OracleTimelockController *OracleTimelockControllerTransactorSession) ScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _OracleTimelockController.Contract.ScheduleBatch(&_OracleTimelockController.TransactOpts, targets, values, payloads, predecessor, salt, delay)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OracleTimelockController *OracleTimelockControllerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleTimelockController.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OracleTimelockController *OracleTimelockControllerSession) Receive() (*types.Transaction, error) {
	return _OracleTimelockController.Contract.Receive(&_OracleTimelockController.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OracleTimelockController *OracleTimelockControllerTransactorSession) Receive() (*types.Transaction, error) {
	return _OracleTimelockController.Contract.Receive(&_OracleTimelockController.TransactOpts)
}

// OracleTimelockControllerCallExecutedIterator is returned from FilterCallExecuted and is used to iterate over the raw logs and unpacked data for CallExecuted events raised by the OracleTimelockController contract.
type OracleTimelockControllerCallExecutedIterator struct {
	Event *OracleTimelockControllerCallExecuted // Event containing the contract specifics and raw log

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
func (it *OracleTimelockControllerCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTimelockControllerCallExecuted)
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
		it.Event = new(OracleTimelockControllerCallExecuted)
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
func (it *OracleTimelockControllerCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTimelockControllerCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTimelockControllerCallExecuted represents a CallExecuted event raised by the OracleTimelockController contract.
type OracleTimelockControllerCallExecuted struct {
	Id     [32]byte
	Index  *big.Int
	Target common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallExecuted is a free log retrieval operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_OracleTimelockController *OracleTimelockControllerFilterer) FilterCallExecuted(opts *bind.FilterOpts, id [][32]byte, index []*big.Int) (*OracleTimelockControllerCallExecutedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _OracleTimelockController.contract.FilterLogs(opts, "CallExecuted", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &OracleTimelockControllerCallExecutedIterator{contract: _OracleTimelockController.contract, event: "CallExecuted", logs: logs, sub: sub}, nil
}

// WatchCallExecuted is a free log subscription operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_OracleTimelockController *OracleTimelockControllerFilterer) WatchCallExecuted(opts *bind.WatchOpts, sink chan<- *OracleTimelockControllerCallExecuted, id [][32]byte, index []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _OracleTimelockController.contract.WatchLogs(opts, "CallExecuted", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTimelockControllerCallExecuted)
				if err := _OracleTimelockController.contract.UnpackLog(event, "CallExecuted", log); err != nil {
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

// ParseCallExecuted is a log parse operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_OracleTimelockController *OracleTimelockControllerFilterer) ParseCallExecuted(log types.Log) (*OracleTimelockControllerCallExecuted, error) {
	event := new(OracleTimelockControllerCallExecuted)
	if err := _OracleTimelockController.contract.UnpackLog(event, "CallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleTimelockControllerCallSaltIterator is returned from FilterCallSalt and is used to iterate over the raw logs and unpacked data for CallSalt events raised by the OracleTimelockController contract.
type OracleTimelockControllerCallSaltIterator struct {
	Event *OracleTimelockControllerCallSalt // Event containing the contract specifics and raw log

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
func (it *OracleTimelockControllerCallSaltIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTimelockControllerCallSalt)
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
		it.Event = new(OracleTimelockControllerCallSalt)
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
func (it *OracleTimelockControllerCallSaltIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTimelockControllerCallSaltIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTimelockControllerCallSalt represents a CallSalt event raised by the OracleTimelockController contract.
type OracleTimelockControllerCallSalt struct {
	Id   [32]byte
	Salt [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCallSalt is a free log retrieval operation binding the contract event 0x20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d0387.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (_OracleTimelockController *OracleTimelockControllerFilterer) FilterCallSalt(opts *bind.FilterOpts, id [][32]byte) (*OracleTimelockControllerCallSaltIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _OracleTimelockController.contract.FilterLogs(opts, "CallSalt", idRule)
	if err != nil {
		return nil, err
	}
	return &OracleTimelockControllerCallSaltIterator{contract: _OracleTimelockController.contract, event: "CallSalt", logs: logs, sub: sub}, nil
}

// WatchCallSalt is a free log subscription operation binding the contract event 0x20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d0387.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (_OracleTimelockController *OracleTimelockControllerFilterer) WatchCallSalt(opts *bind.WatchOpts, sink chan<- *OracleTimelockControllerCallSalt, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _OracleTimelockController.contract.WatchLogs(opts, "CallSalt", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTimelockControllerCallSalt)
				if err := _OracleTimelockController.contract.UnpackLog(event, "CallSalt", log); err != nil {
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

// ParseCallSalt is a log parse operation binding the contract event 0x20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d0387.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (_OracleTimelockController *OracleTimelockControllerFilterer) ParseCallSalt(log types.Log) (*OracleTimelockControllerCallSalt, error) {
	event := new(OracleTimelockControllerCallSalt)
	if err := _OracleTimelockController.contract.UnpackLog(event, "CallSalt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleTimelockControllerCallScheduledIterator is returned from FilterCallScheduled and is used to iterate over the raw logs and unpacked data for CallScheduled events raised by the OracleTimelockController contract.
type OracleTimelockControllerCallScheduledIterator struct {
	Event *OracleTimelockControllerCallScheduled // Event containing the contract specifics and raw log

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
func (it *OracleTimelockControllerCallScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTimelockControllerCallScheduled)
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
		it.Event = new(OracleTimelockControllerCallScheduled)
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
func (it *OracleTimelockControllerCallScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTimelockControllerCallScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTimelockControllerCallScheduled represents a CallScheduled event raised by the OracleTimelockController contract.
type OracleTimelockControllerCallScheduled struct {
	Id          [32]byte
	Index       *big.Int
	Target      common.Address
	Value       *big.Int
	Data        []byte
	Predecessor [32]byte
	Delay       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCallScheduled is a free log retrieval operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_OracleTimelockController *OracleTimelockControllerFilterer) FilterCallScheduled(opts *bind.FilterOpts, id [][32]byte, index []*big.Int) (*OracleTimelockControllerCallScheduledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _OracleTimelockController.contract.FilterLogs(opts, "CallScheduled", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &OracleTimelockControllerCallScheduledIterator{contract: _OracleTimelockController.contract, event: "CallScheduled", logs: logs, sub: sub}, nil
}

// WatchCallScheduled is a free log subscription operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_OracleTimelockController *OracleTimelockControllerFilterer) WatchCallScheduled(opts *bind.WatchOpts, sink chan<- *OracleTimelockControllerCallScheduled, id [][32]byte, index []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _OracleTimelockController.contract.WatchLogs(opts, "CallScheduled", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTimelockControllerCallScheduled)
				if err := _OracleTimelockController.contract.UnpackLog(event, "CallScheduled", log); err != nil {
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

// ParseCallScheduled is a log parse operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_OracleTimelockController *OracleTimelockControllerFilterer) ParseCallScheduled(log types.Log) (*OracleTimelockControllerCallScheduled, error) {
	event := new(OracleTimelockControllerCallScheduled)
	if err := _OracleTimelockController.contract.UnpackLog(event, "CallScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleTimelockControllerCancelledIterator is returned from FilterCancelled and is used to iterate over the raw logs and unpacked data for Cancelled events raised by the OracleTimelockController contract.
type OracleTimelockControllerCancelledIterator struct {
	Event *OracleTimelockControllerCancelled // Event containing the contract specifics and raw log

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
func (it *OracleTimelockControllerCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTimelockControllerCancelled)
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
		it.Event = new(OracleTimelockControllerCancelled)
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
func (it *OracleTimelockControllerCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTimelockControllerCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTimelockControllerCancelled represents a Cancelled event raised by the OracleTimelockController contract.
type OracleTimelockControllerCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancelled is a free log retrieval operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_OracleTimelockController *OracleTimelockControllerFilterer) FilterCancelled(opts *bind.FilterOpts, id [][32]byte) (*OracleTimelockControllerCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _OracleTimelockController.contract.FilterLogs(opts, "Cancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &OracleTimelockControllerCancelledIterator{contract: _OracleTimelockController.contract, event: "Cancelled", logs: logs, sub: sub}, nil
}

// WatchCancelled is a free log subscription operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_OracleTimelockController *OracleTimelockControllerFilterer) WatchCancelled(opts *bind.WatchOpts, sink chan<- *OracleTimelockControllerCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _OracleTimelockController.contract.WatchLogs(opts, "Cancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTimelockControllerCancelled)
				if err := _OracleTimelockController.contract.UnpackLog(event, "Cancelled", log); err != nil {
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

// ParseCancelled is a log parse operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_OracleTimelockController *OracleTimelockControllerFilterer) ParseCancelled(log types.Log) (*OracleTimelockControllerCancelled, error) {
	event := new(OracleTimelockControllerCancelled)
	if err := _OracleTimelockController.contract.UnpackLog(event, "Cancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleTimelockControllerMinDelayChangeIterator is returned from FilterMinDelayChange and is used to iterate over the raw logs and unpacked data for MinDelayChange events raised by the OracleTimelockController contract.
type OracleTimelockControllerMinDelayChangeIterator struct {
	Event *OracleTimelockControllerMinDelayChange // Event containing the contract specifics and raw log

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
func (it *OracleTimelockControllerMinDelayChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTimelockControllerMinDelayChange)
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
		it.Event = new(OracleTimelockControllerMinDelayChange)
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
func (it *OracleTimelockControllerMinDelayChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTimelockControllerMinDelayChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTimelockControllerMinDelayChange represents a MinDelayChange event raised by the OracleTimelockController contract.
type OracleTimelockControllerMinDelayChange struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinDelayChange is a free log retrieval operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_OracleTimelockController *OracleTimelockControllerFilterer) FilterMinDelayChange(opts *bind.FilterOpts) (*OracleTimelockControllerMinDelayChangeIterator, error) {

	logs, sub, err := _OracleTimelockController.contract.FilterLogs(opts, "MinDelayChange")
	if err != nil {
		return nil, err
	}
	return &OracleTimelockControllerMinDelayChangeIterator{contract: _OracleTimelockController.contract, event: "MinDelayChange", logs: logs, sub: sub}, nil
}

// WatchMinDelayChange is a free log subscription operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_OracleTimelockController *OracleTimelockControllerFilterer) WatchMinDelayChange(opts *bind.WatchOpts, sink chan<- *OracleTimelockControllerMinDelayChange) (event.Subscription, error) {

	logs, sub, err := _OracleTimelockController.contract.WatchLogs(opts, "MinDelayChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTimelockControllerMinDelayChange)
				if err := _OracleTimelockController.contract.UnpackLog(event, "MinDelayChange", log); err != nil {
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

// ParseMinDelayChange is a log parse operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_OracleTimelockController *OracleTimelockControllerFilterer) ParseMinDelayChange(log types.Log) (*OracleTimelockControllerMinDelayChange, error) {
	event := new(OracleTimelockControllerMinDelayChange)
	if err := _OracleTimelockController.contract.UnpackLog(event, "MinDelayChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleTimelockControllerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the OracleTimelockController contract.
type OracleTimelockControllerRoleAdminChangedIterator struct {
	Event *OracleTimelockControllerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *OracleTimelockControllerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTimelockControllerRoleAdminChanged)
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
		it.Event = new(OracleTimelockControllerRoleAdminChanged)
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
func (it *OracleTimelockControllerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTimelockControllerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTimelockControllerRoleAdminChanged represents a RoleAdminChanged event raised by the OracleTimelockController contract.
type OracleTimelockControllerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OracleTimelockController *OracleTimelockControllerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*OracleTimelockControllerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _OracleTimelockController.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &OracleTimelockControllerRoleAdminChangedIterator{contract: _OracleTimelockController.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OracleTimelockController *OracleTimelockControllerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *OracleTimelockControllerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _OracleTimelockController.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTimelockControllerRoleAdminChanged)
				if err := _OracleTimelockController.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OracleTimelockController *OracleTimelockControllerFilterer) ParseRoleAdminChanged(log types.Log) (*OracleTimelockControllerRoleAdminChanged, error) {
	event := new(OracleTimelockControllerRoleAdminChanged)
	if err := _OracleTimelockController.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleTimelockControllerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the OracleTimelockController contract.
type OracleTimelockControllerRoleGrantedIterator struct {
	Event *OracleTimelockControllerRoleGranted // Event containing the contract specifics and raw log

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
func (it *OracleTimelockControllerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTimelockControllerRoleGranted)
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
		it.Event = new(OracleTimelockControllerRoleGranted)
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
func (it *OracleTimelockControllerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTimelockControllerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTimelockControllerRoleGranted represents a RoleGranted event raised by the OracleTimelockController contract.
type OracleTimelockControllerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OracleTimelockController *OracleTimelockControllerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OracleTimelockControllerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleTimelockController.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OracleTimelockControllerRoleGrantedIterator{contract: _OracleTimelockController.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OracleTimelockController *OracleTimelockControllerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *OracleTimelockControllerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleTimelockController.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTimelockControllerRoleGranted)
				if err := _OracleTimelockController.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OracleTimelockController *OracleTimelockControllerFilterer) ParseRoleGranted(log types.Log) (*OracleTimelockControllerRoleGranted, error) {
	event := new(OracleTimelockControllerRoleGranted)
	if err := _OracleTimelockController.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleTimelockControllerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the OracleTimelockController contract.
type OracleTimelockControllerRoleRevokedIterator struct {
	Event *OracleTimelockControllerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *OracleTimelockControllerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTimelockControllerRoleRevoked)
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
		it.Event = new(OracleTimelockControllerRoleRevoked)
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
func (it *OracleTimelockControllerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTimelockControllerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTimelockControllerRoleRevoked represents a RoleRevoked event raised by the OracleTimelockController contract.
type OracleTimelockControllerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OracleTimelockController *OracleTimelockControllerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OracleTimelockControllerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleTimelockController.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OracleTimelockControllerRoleRevokedIterator{contract: _OracleTimelockController.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OracleTimelockController *OracleTimelockControllerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *OracleTimelockControllerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleTimelockController.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTimelockControllerRoleRevoked)
				if err := _OracleTimelockController.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OracleTimelockController *OracleTimelockControllerFilterer) ParseRoleRevoked(log types.Log) (*OracleTimelockControllerRoleRevoked, error) {
	event := new(OracleTimelockControllerRoleRevoked)
	if err := _OracleTimelockController.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
