// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer_1_0_20\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1002,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1003,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_uint256)1004_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1004_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[47]\",\"numberOfBytes\":\"1504\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L2StandardBridgeStorageLayout = new(solc.StorageLayout)

var L2StandardBridgeDeployedBin = "0x6080604052600436106100ec5760003560e01c806354fd4d501161008a5780638f601f66116100595780638f601f661461034e578063927ede2d14610394578063a3a79548146103c8578063e11013dd146103db57600080fd5b806354fd4d50146102c5578063662a633a146102e75780637f46ddb2146102fa578063870876231461032e57600080fd5b806332b7006d116100c657806332b7006d1461020657806336c717c1146102195780633cb747bf14610272578063540abf73146102a557600080fd5b80630166a07a146101c057806309fc8843146101e05780631635f5fd146101f357600080fd5b366101bb57333b15610185576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b6101b973deaddeaddeaddeaddeaddeaddeaddeaddead000033333462030d40604051806020016040528060008152506103ee565b005b600080fd5b3480156101cc57600080fd5b506101b96101db366004612103565b610548565b6101b96101ee3660046121b4565b61096b565b6101b9610201366004612207565b610a42565b6101b961021436600461227a565b610f3e565b34801561022557600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561027e57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610248565b3480156102b157600080fd5b506101b96102c03660046122ce565b611018565b3480156102d157600080fd5b506102da611066565b60405161026991906123bb565b6101b96102f5366004612103565b611109565b34801561030657600080fd5b506102487f000000000000000000000000000000000000000000000000000000000000000081565b34801561033a57600080fd5b506101b96103493660046123ce565b6111f6565b34801561035a57600080fd5b50610386610369366004612451565b600260209081526000928352604080842090915290825290205481565b604051908152602001610269565b3480156103a057600080fd5b506102487f000000000000000000000000000000000000000000000000000000000000000081565b6101b96103d63660046123ce565b6112ca565b6101b96103e936600461248a565b61130e565b60008673ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa15801561043b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061045f91906124ed565b90507fffffffffffffffffffffffff215221522152215221522152215221522153000073ffffffffffffffffffffffffffffffffffffffff8816016104b0576104ab8686868686611357565b6104bf565b6104bf87828888888888611596565b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e8888876040516105379392919061250a565b60405180910390a450505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561066657507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561062a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064e91906124ed565b73ffffffffffffffffffffffffffffffffffffffff16145b610718576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a40161017c565b6107218761194f565b1561086f5761073087876119b1565b6107e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a40161017c565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b15801561085257600080fd5b505af1158015610866573d6000803e3d6000fd5b505050506108f1565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a16835292905220546108ad908490612577565b73ffffffffffffffffffffffffffffffffffffffff8089166000818152600260209081526040808320948c16835293905291909120919091556108f1908585611a58565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd8787878760405161053794939291906125d7565b333b156109fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606482015260840161017c565b610a3d3333348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061135792505050565b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148015610b6057507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b4891906124ed565b73ffffffffffffffffffffffffffffffffffffffff16145b610c12576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a40161017c565b823414610ca1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e74207265717569726564000000000000606482015260840161017c565b3073ffffffffffffffffffffffffffffffffffffffff851603610d46576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c660000000000000000000000000000000000000000000000000000000000606482015260840161017c565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610e21576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201527f657373656e676572000000000000000000000000000000000000000000000000606482015260840161017c565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d858585604051610e8293929190612617565b60405180910390a36000610ea7855a8660405180602001604052806000815250611b2c565b905080610f36576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c65640000000000000000000000000000000000000000000000000000000000606482015260840161017c565b505050505050565b333b15610fcd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606482015260840161017c565b611011853333878787878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506103ee92505050565b5050505050565b61105d87873388888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061159692505050565b50505050505050565b60606110917f0000000000000000000000000000000000000000000000000000000000000000611b46565b6110ba7f0000000000000000000000000000000000000000000000000000000000000000611b46565b6110e37f0000000000000000000000000000000000000000000000000000000000000000611b46565b6040516020016110f593929190612631565b604051602081830303815290604052905090565b73ffffffffffffffffffffffffffffffffffffffff8716158015611156575073ffffffffffffffffffffffffffffffffffffffff861673deaddeaddeaddeaddeaddeaddeaddeaddead0000145b1561116d576111688585858585610a42565b61117c565b61117c86888787878787610548565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd898787878760405161053794939291906125d7565b333b15611285576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606482015260840161017c565b610f3686863333888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061159692505050565b610f36863387878787878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506103ee92505050565b6113513385348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061135792505050565b50505050565b8234146113e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e6720455448206d757360448201527f7420696e636c7564652073756666696369656e74204554482076616c75650000606482015260840161017c565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af585846040516114459291906126a7565b60405180910390a37f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b847f0000000000000000000000000000000000000000000000000000000000000000631635f5fd60e01b898989886040516024016114ca94939291906126c0565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b909216825261155d929188906004016126ff565b6000604051808303818588803b15801561157657600080fd5b505af115801561158a573d6000803e3d6000fd5b50505050505050505050565b61159f8761194f565b156116ed576115ae87876119b1565b611660576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a40161017c565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b1580156116d057600080fd5b505af11580156116e4573d6000803e3d6000fd5b50505050611781565b61170f73ffffffffffffffffffffffffffffffffffffffff8816863086611c83565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a168352929052205461174d908490612744565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b16835292905220555b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf8787866040516117f99392919061250a565b60405180910390a47f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b7f0000000000000000000000000000000000000000000000000000000000000000630166a07a60e01b898b8a8a8a896040516024016118819695949392919061275c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b9092168252611914929187906004016126ff565b600060405180830381600087803b15801561192e57600080fd5b505af1158015611942573d6000803e3d6000fd5b5050505050505050505050565b600061197b827f1d1d8b6300000000000000000000000000000000000000000000000000000000611ce1565b806119ab57506119ab827fec4fc8e300000000000000000000000000000000000000000000000000000000611ce1565b92915050565b60008273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119fe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a2291906124ed565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614905092915050565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610a3d9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611d04565b600080600080845160208601878a8af19695505050505050565b606081600003611b8957505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611bb35780611b9d816127b7565b9150611bac9050600a8361281e565b9150611b8d565b60008167ffffffffffffffff811115611bce57611bce612832565b6040519080825280601f01601f191660200182016040528015611bf8576020820181803683370190505b5090505b8415611c7b57611c0d600183612577565b9150611c1a600a86612861565b611c25906030612744565b60f81b818381518110611c3a57611c3a612875565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611c74600a8661281e565b9450611bfc565b949350505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526113519085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611aaa565b6000611cec83611e10565b8015611cfd5750611cfd8383611e74565b9392505050565b6000611d66826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611f439092919063ffffffff16565b805190915015610a3d5780806020019051810190611d8491906128a4565b610a3d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161017c565b6000611e3c827f01ffc9a700000000000000000000000000000000000000000000000000000000611e74565b80156119ab5750611e6d827fffffffff00000000000000000000000000000000000000000000000000000000611e74565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015611f2c575060208210155b8015611f385750600081115b979650505050505050565b6060611c7b84846000858573ffffffffffffffffffffffffffffffffffffffff85163b611fcc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161017c565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611ff591906128c6565b60006040518083038185875af1925050503d8060008114612032576040519150601f19603f3d011682016040523d82523d6000602084013e612037565b606091505b5091509150611f3882828660608315612051575081611cfd565b8251156120615782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017c91906123bb565b73ffffffffffffffffffffffffffffffffffffffff811681146120b757600080fd5b50565b60008083601f8401126120cc57600080fd5b50813567ffffffffffffffff8111156120e457600080fd5b6020830191508360208285010111156120fc57600080fd5b9250929050565b600080600080600080600060c0888a03121561211e57600080fd5b873561212981612095565b9650602088013561213981612095565b9550604088013561214981612095565b9450606088013561215981612095565b93506080880135925060a088013567ffffffffffffffff81111561217c57600080fd5b6121888a828b016120ba565b989b979a50959850939692959293505050565b803563ffffffff811681146121af57600080fd5b919050565b6000806000604084860312156121c957600080fd5b6121d28461219b565b9250602084013567ffffffffffffffff8111156121ee57600080fd5b6121fa868287016120ba565b9497909650939450505050565b60008060008060006080868803121561221f57600080fd5b853561222a81612095565b9450602086013561223a81612095565b935060408601359250606086013567ffffffffffffffff81111561225d57600080fd5b612269888289016120ba565b969995985093965092949392505050565b60008060008060006080868803121561229257600080fd5b853561229d81612095565b9450602086013593506122b26040870161219b565b9250606086013567ffffffffffffffff81111561225d57600080fd5b600080600080600080600060c0888a0312156122e957600080fd5b87356122f481612095565b9650602088013561230481612095565b9550604088013561231481612095565b9450606088013593506123296080890161219b565b925060a088013567ffffffffffffffff81111561217c57600080fd5b60005b83811015612360578181015183820152602001612348565b838111156113515750506000910152565b60008151808452612389816020860160208601612345565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611cfd6020830184612371565b60008060008060008060a087890312156123e757600080fd5b86356123f281612095565b9550602087013561240281612095565b9450604087013593506124176060880161219b565b9250608087013567ffffffffffffffff81111561243357600080fd5b61243f89828a016120ba565b979a9699509497509295939492505050565b6000806040838503121561246457600080fd5b823561246f81612095565b9150602083013561247f81612095565b809150509250929050565b600080600080606085870312156124a057600080fd5b84356124ab81612095565b93506124b96020860161219b565b9250604085013567ffffffffffffffff8111156124d557600080fd5b6124e1878288016120ba565b95989497509550505050565b6000602082840312156124ff57600080fd5b8151611cfd81612095565b73ffffffffffffffffffffffffffffffffffffffff8416815282602082015260606040820152600061253f6060830184612371565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561258957612589612548565b500390565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260606040820152600061260d60608301848661258e565b9695505050505050565b83815260406020820152600061253f60408301848661258e565b60008451612643818460208901612345565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855161267f816001850160208a01612345565b6001920191820152835161269a816002840160208801612345565b0160020195945050505050565b828152604060208201526000611c7b6040830184612371565b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152508360408301526080606083015261260d6080830184612371565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600061272e6060830185612371565b905063ffffffff83166040830152949350505050565b6000821982111561275757612757612548565b500190565b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a08301526127ab60c0830184612371565b98975050505050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036127e8576127e8612548565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261282d5761282d6127ef565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082612870576128706127ef565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156128b657600080fd5b81518015158114611cfd57600080fd5b600082516128d8818460208701612345565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2StandardBridgeStorageLayoutJSON), L2StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2StandardBridge"] = L2StandardBridgeStorageLayout
	deployedBytecodes["L2StandardBridge"] = L2StandardBridgeDeployedBin
}
