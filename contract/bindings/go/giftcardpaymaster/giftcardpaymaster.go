// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package giftcardpaymaster

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

// GiftcardpaymasterMetaData contains all meta data concerning the Giftcardpaymaster contract.
var GiftcardpaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifyingSigner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"}],\"name\":\"VerifyingSignerUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SIGNATURE_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALID_AFTER_OFFSET\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALID_UNTIL_OFFSET\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newSigner\",\"type\":\"address\"}],\"name\":\"setVerifyingSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyingSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161135b38038061135b83398101604081905261002e91610222565b81338061005557604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61005e816100ea565b5061006881610106565b6001600160a01b0390811660805281166100c45760405162461bcd60e51b815260206004820152601860248201527f496e76616c696420766572696679696e67207369676e65720000000000000000604482015260640161004c565b600280546001600160a01b0319166001600160a01b039290921691909117905550610280565b600180546001600160a01b0319169055610103816101bf565b50565b6040516301ffc9a760e01b8152631313998b60e31b60048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa15801561014f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610173919061025a565b6101035760405162461bcd60e51b815260206004820152601e60248201527f49456e747279506f696e7420696e74657266616365206d69736d617463680000604482015260640161004c565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0381168114610103575f5ffd5b5f5f60408385031215610233575f5ffd5b825161023e8161020e565b602084015190925061024f8161020e565b809150509250929050565b5f6020828403121561026a575f5ffd5b81518015158114610279575f5ffd5b9392505050565b60805161108b6102d05f395f81816101230152818161032d0152818161041b015281816104ae0152818161058f015281816106240152818161068e0152818161071901526108d2015261108b5ff3fe608060405260043610610113575f3560e01c80638da5cb5b1161009d578063c399ec8811610062578063c399ec8814610382578063d0e30db014610396578063e30c39781461039e578063f2fde38b146103bb578063f5cba98c146103da575f5ffd5b80638da5cb5b146102ed5780639310001a14610309578063b0d691fe1461031c578063bb9fe6bf1461034f578063c23a5cea14610363575f5ffd5b806352b7512c116100e357806352b7512c146102575780636c7e9b3c14610284578063715018a6146102a657806379ba5097146102ba5780637c627b21146102ce575f5ffd5b80630396cb60146101be578063053c5e59146101d1578063205c28781461020157806323d9ac9b14610220575f5ffd5b366101ba57336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061015857505f546001600160a01b031633145b6101b85760405162461bcd60e51b815260206004820152602660248201527f47696674436172645061796d61737465723a20756e617574686f72697a65642060448201526539b2b73232b960d11b60648201526084015b60405180910390fd5b005b5f5ffd5b6101b86101cc366004610ddf565b6103f9565b3480156101dc575f5ffd5b506101e5603081565b60405165ffffffffffff90911681526020015b60405180910390f35b34801561020c575f5ffd5b506101b861021b366004610e1d565b610480565b34801561022b575f5ffd5b5060025461023f906001600160a01b031681565b6040516001600160a01b0390911681526020016101f8565b348015610262575f5ffd5b50610276610271366004610e47565b6104ef565b6040516101f8929190610e96565b34801561028f575f5ffd5b50610298606081565b6040519081526020016101f8565b3480156102b1575f5ffd5b506101b8610512565b3480156102c5575f5ffd5b506101b8610525565b3480156102d9575f5ffd5b506101b86102e8366004610ed2565b610569565b3480156102f8575f5ffd5b505f546001600160a01b031661023f565b348015610314575f5ffd5b506101e55f81565b348015610327575f5ffd5b5061023f7f000000000000000000000000000000000000000000000000000000000000000081565b34801561035a575f5ffd5b506101b8610585565b34801561036e575f5ffd5b506101b861037d366004610f66565b6105fd565b34801561038d575f5ffd5b50610298610677565b6101b8610704565b3480156103a9575f5ffd5b506001546001600160a01b031661023f565b3480156103c6575f5ffd5b506101b86103d5366004610f66565b610764565b3480156103e5575f5ffd5b506101b86103f4366004610f66565b6107d4565b61040161089b565b604051621cb65b60e51b815263ffffffff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690630396cb609034906024015f604051808303818588803b158015610466575f5ffd5b505af1158015610478573d5f5f3e3d5ffd5b505050505050565b61048861089b565b60405163040b850f60e31b81526001600160a01b038381166004830152602482018390527f0000000000000000000000000000000000000000000000000000000000000000169063205c2878906044015f604051808303815f87803b158015610466575f5ffd5b60605f6104fa6108c7565b610505858585610937565b915091505b935093915050565b61051a61089b565b6105235f610b17565b565b60015433906001600160a01b0316811461055d5760405163118cdaa760e01b81526001600160a01b03821660048201526024016101af565b61056681610b17565b50565b6105716108c7565b61057e8585858585610b30565b5050505050565b61058d61089b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bb9fe6bf6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156105e5575f5ffd5b505af11580156105f7573d5f5f3e3d5ffd5b50505050565b61060561089b565b60405163611d2e7560e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c23a5cea906024015f604051808303815f87803b158015610665575f5ffd5b505af115801561057e573d5f5f3e3d5ffd5b6040516370a0823160e01b81523060048201525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa1580156106db573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106ff9190610f81565b905090565b60405163b760faf960e01b81523060048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063b760faf99034906024015f604051808303818588803b158015610665575f5ffd5b61076c61089b565b600180546001600160a01b0383166001600160a01b0319909116811790915561079c5f546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6107dc61089b565b6001600160a01b0381166108325760405162461bcd60e51b815260206004820152601860248201527f496e76616c696420766572696679696e67207369676e6572000000000000000060448201526064016101af565b600254604080516001600160a01b03928316815291831660208301527fdeee0b46266dc0734b009b2ca66e2ac28d5fd72d38506a3fe462865f00c5df4e910160405180910390a1600280546001600160a01b0319166001600160a01b0392909216919091179055565b5f546001600160a01b031633146105235760405163118cdaa760e01b81523360048201526024016101af565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105235760405162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08115b9d1c9e541bda5b9d605a1b60448201526064016101af565b60605f368161094960e0880188610f98565b9092509050606181101561099f5760405162461bcd60e51b815260206004820152601f60248201527f496e76616c6964207061796d6173746572416e6444617461206c656e6774680060448201526064016101af565b5f6109ae601a60148486610fe2565b6109b791611009565b60d01c90505f6109cb6020601a8587610fe2565b6109d491611009565b60d01c9050365f6109e88560208189610fe2565b915091505f8a8585604051602001610a1b9392919092835265ffffffffffff918216602084015216604082015260600190565b6040516020818303038152906040528051906020012090505f610a6a827f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f908152601c91909152603c902090565b90505f610aae85858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508693925050610b689050565b6002549091506001600160a01b03808316911614610aea57600160405180602001604052805f815250909a509a5050505050505050505061050a565b610af55f8888610b90565b60408051602081019091525f81529f909e509c50505050505050505050505050565b600180546001600160a01b031916905561056681610bc3565b60405162461bcd60e51b815260206004820152600d60248201526c6d757374206f7665727269646560981b60448201526064016101af565b5f5f5f5f610b768686610c12565b925092509250610b868282610c5b565b5090949350505050565b5f60d08265ffffffffffff16901b60a08465ffffffffffff16901b85610bb6575f610bb9565b60015b1717949350505050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f5f5f8351604103610c49576020840151604085015160608601515f1a610c3b88828585610d17565b955095509550505050610c54565b505081515f91506002905b9250925092565b5f826003811115610c6e57610c6e611041565b03610c77575050565b6001826003811115610c8b57610c8b611041565b03610ca95760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115610cbd57610cbd611041565b03610cde5760405163fce698f760e01b8152600481018290526024016101af565b6003826003811115610cf257610cf2611041565b03610d13576040516335e2f38360e21b8152600481018290526024016101af565b5050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115610d5057505f91506003905082610dd5565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610da1573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610dcc57505f925060019150829050610dd5565b92505f91508190505b9450945094915050565b5f60208284031215610def575f5ffd5b813563ffffffff81168114610e02575f5ffd5b9392505050565b6001600160a01b0381168114610566575f5ffd5b5f5f60408385031215610e2e575f5ffd5b8235610e3981610e09565b946020939093013593505050565b5f5f5f60608486031215610e59575f5ffd5b833567ffffffffffffffff811115610e6f575f5ffd5b84016101208187031215610e81575f5ffd5b95602085013595506040909401359392505050565b604081525f83518060408401528060208601606085015e5f606082850101526060601f19601f8301168401019150508260208301529392505050565b5f5f5f5f5f60808688031215610ee6575f5ffd5b853560038110610ef4575f5ffd5b9450602086013567ffffffffffffffff811115610f0f575f5ffd5b8601601f81018813610f1f575f5ffd5b803567ffffffffffffffff811115610f35575f5ffd5b886020828401011115610f46575f5ffd5b959860209190910197509495604081013595606090910135945092505050565b5f60208284031215610f76575f5ffd5b8135610e0281610e09565b5f60208284031215610f91575f5ffd5b5051919050565b5f5f8335601e19843603018112610fad575f5ffd5b83018035915067ffffffffffffffff821115610fc7575f5ffd5b602001915036819003821315610fdb575f5ffd5b9250929050565b5f5f85851115610ff0575f5ffd5b83861115610ffc575f5ffd5b5050820193919092039150565b80356001600160d01b0319811690600684101561103a576001600160d01b0319600685900360031b81901b82161691505b5092915050565b634e487b7160e01b5f52602160045260245ffdfea26469706673582212200c14f009882b0498f229d3f31e388e20f1cf581e74808a51d8cd3871870fe63c64736f6c634300081c0033",
}

// GiftcardpaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use GiftcardpaymasterMetaData.ABI instead.
var GiftcardpaymasterABI = GiftcardpaymasterMetaData.ABI

// GiftcardpaymasterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GiftcardpaymasterMetaData.Bin instead.
var GiftcardpaymasterBin = GiftcardpaymasterMetaData.Bin

// DeployGiftcardpaymaster deploys a new Ethereum contract, binding an instance of Giftcardpaymaster to it.
func DeployGiftcardpaymaster(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address, _verifyingSigner common.Address) (common.Address, *types.Transaction, *Giftcardpaymaster, error) {
	parsed, err := GiftcardpaymasterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GiftcardpaymasterBin), backend, _entryPoint, _verifyingSigner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Giftcardpaymaster{GiftcardpaymasterCaller: GiftcardpaymasterCaller{contract: contract}, GiftcardpaymasterTransactor: GiftcardpaymasterTransactor{contract: contract}, GiftcardpaymasterFilterer: GiftcardpaymasterFilterer{contract: contract}}, nil
}

// Giftcardpaymaster is an auto generated Go binding around an Ethereum contract.
type Giftcardpaymaster struct {
	GiftcardpaymasterCaller     // Read-only binding to the contract
	GiftcardpaymasterTransactor // Write-only binding to the contract
	GiftcardpaymasterFilterer   // Log filterer for contract events
}

// GiftcardpaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type GiftcardpaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GiftcardpaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GiftcardpaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GiftcardpaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GiftcardpaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GiftcardpaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GiftcardpaymasterSession struct {
	Contract     *Giftcardpaymaster // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GiftcardpaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GiftcardpaymasterCallerSession struct {
	Contract *GiftcardpaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GiftcardpaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GiftcardpaymasterTransactorSession struct {
	Contract     *GiftcardpaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GiftcardpaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type GiftcardpaymasterRaw struct {
	Contract *Giftcardpaymaster // Generic contract binding to access the raw methods on
}

// GiftcardpaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GiftcardpaymasterCallerRaw struct {
	Contract *GiftcardpaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// GiftcardpaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GiftcardpaymasterTransactorRaw struct {
	Contract *GiftcardpaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGiftcardpaymaster creates a new instance of Giftcardpaymaster, bound to a specific deployed contract.
func NewGiftcardpaymaster(address common.Address, backend bind.ContractBackend) (*Giftcardpaymaster, error) {
	contract, err := bindGiftcardpaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Giftcardpaymaster{GiftcardpaymasterCaller: GiftcardpaymasterCaller{contract: contract}, GiftcardpaymasterTransactor: GiftcardpaymasterTransactor{contract: contract}, GiftcardpaymasterFilterer: GiftcardpaymasterFilterer{contract: contract}}, nil
}

// NewGiftcardpaymasterCaller creates a new read-only instance of Giftcardpaymaster, bound to a specific deployed contract.
func NewGiftcardpaymasterCaller(address common.Address, caller bind.ContractCaller) (*GiftcardpaymasterCaller, error) {
	contract, err := bindGiftcardpaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GiftcardpaymasterCaller{contract: contract}, nil
}

// NewGiftcardpaymasterTransactor creates a new write-only instance of Giftcardpaymaster, bound to a specific deployed contract.
func NewGiftcardpaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*GiftcardpaymasterTransactor, error) {
	contract, err := bindGiftcardpaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GiftcardpaymasterTransactor{contract: contract}, nil
}

// NewGiftcardpaymasterFilterer creates a new log filterer instance of Giftcardpaymaster, bound to a specific deployed contract.
func NewGiftcardpaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*GiftcardpaymasterFilterer, error) {
	contract, err := bindGiftcardpaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GiftcardpaymasterFilterer{contract: contract}, nil
}

// bindGiftcardpaymaster binds a generic wrapper to an already deployed contract.
func bindGiftcardpaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GiftcardpaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Giftcardpaymaster *GiftcardpaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Giftcardpaymaster.Contract.GiftcardpaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Giftcardpaymaster *GiftcardpaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.GiftcardpaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Giftcardpaymaster *GiftcardpaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.GiftcardpaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Giftcardpaymaster *GiftcardpaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Giftcardpaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Giftcardpaymaster *GiftcardpaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Giftcardpaymaster *GiftcardpaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.contract.Transact(opts, method, params...)
}

// SIGNATUREOFFSET is a free data retrieval call binding the contract method 0x6c7e9b3c.
//
// Solidity: function SIGNATURE_OFFSET() view returns(uint256)
func (_Giftcardpaymaster *GiftcardpaymasterCaller) SIGNATUREOFFSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Giftcardpaymaster.contract.Call(opts, &out, "SIGNATURE_OFFSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGNATUREOFFSET is a free data retrieval call binding the contract method 0x6c7e9b3c.
//
// Solidity: function SIGNATURE_OFFSET() view returns(uint256)
func (_Giftcardpaymaster *GiftcardpaymasterSession) SIGNATUREOFFSET() (*big.Int, error) {
	return _Giftcardpaymaster.Contract.SIGNATUREOFFSET(&_Giftcardpaymaster.CallOpts)
}

// SIGNATUREOFFSET is a free data retrieval call binding the contract method 0x6c7e9b3c.
//
// Solidity: function SIGNATURE_OFFSET() view returns(uint256)
func (_Giftcardpaymaster *GiftcardpaymasterCallerSession) SIGNATUREOFFSET() (*big.Int, error) {
	return _Giftcardpaymaster.Contract.SIGNATUREOFFSET(&_Giftcardpaymaster.CallOpts)
}

// VALIDAFTEROFFSET is a free data retrieval call binding the contract method 0x053c5e59.
//
// Solidity: function VALID_AFTER_OFFSET() view returns(uint48)
func (_Giftcardpaymaster *GiftcardpaymasterCaller) VALIDAFTEROFFSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Giftcardpaymaster.contract.Call(opts, &out, "VALID_AFTER_OFFSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VALIDAFTEROFFSET is a free data retrieval call binding the contract method 0x053c5e59.
//
// Solidity: function VALID_AFTER_OFFSET() view returns(uint48)
func (_Giftcardpaymaster *GiftcardpaymasterSession) VALIDAFTEROFFSET() (*big.Int, error) {
	return _Giftcardpaymaster.Contract.VALIDAFTEROFFSET(&_Giftcardpaymaster.CallOpts)
}

// VALIDAFTEROFFSET is a free data retrieval call binding the contract method 0x053c5e59.
//
// Solidity: function VALID_AFTER_OFFSET() view returns(uint48)
func (_Giftcardpaymaster *GiftcardpaymasterCallerSession) VALIDAFTEROFFSET() (*big.Int, error) {
	return _Giftcardpaymaster.Contract.VALIDAFTEROFFSET(&_Giftcardpaymaster.CallOpts)
}

// VALIDUNTILOFFSET is a free data retrieval call binding the contract method 0x9310001a.
//
// Solidity: function VALID_UNTIL_OFFSET() view returns(uint48)
func (_Giftcardpaymaster *GiftcardpaymasterCaller) VALIDUNTILOFFSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Giftcardpaymaster.contract.Call(opts, &out, "VALID_UNTIL_OFFSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VALIDUNTILOFFSET is a free data retrieval call binding the contract method 0x9310001a.
//
// Solidity: function VALID_UNTIL_OFFSET() view returns(uint48)
func (_Giftcardpaymaster *GiftcardpaymasterSession) VALIDUNTILOFFSET() (*big.Int, error) {
	return _Giftcardpaymaster.Contract.VALIDUNTILOFFSET(&_Giftcardpaymaster.CallOpts)
}

// VALIDUNTILOFFSET is a free data retrieval call binding the contract method 0x9310001a.
//
// Solidity: function VALID_UNTIL_OFFSET() view returns(uint48)
func (_Giftcardpaymaster *GiftcardpaymasterCallerSession) VALIDUNTILOFFSET() (*big.Int, error) {
	return _Giftcardpaymaster.Contract.VALIDUNTILOFFSET(&_Giftcardpaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Giftcardpaymaster *GiftcardpaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Giftcardpaymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Giftcardpaymaster *GiftcardpaymasterSession) EntryPoint() (common.Address, error) {
	return _Giftcardpaymaster.Contract.EntryPoint(&_Giftcardpaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Giftcardpaymaster *GiftcardpaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _Giftcardpaymaster.Contract.EntryPoint(&_Giftcardpaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Giftcardpaymaster *GiftcardpaymasterCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Giftcardpaymaster.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Giftcardpaymaster *GiftcardpaymasterSession) GetDeposit() (*big.Int, error) {
	return _Giftcardpaymaster.Contract.GetDeposit(&_Giftcardpaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Giftcardpaymaster *GiftcardpaymasterCallerSession) GetDeposit() (*big.Int, error) {
	return _Giftcardpaymaster.Contract.GetDeposit(&_Giftcardpaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Giftcardpaymaster *GiftcardpaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Giftcardpaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Giftcardpaymaster *GiftcardpaymasterSession) Owner() (common.Address, error) {
	return _Giftcardpaymaster.Contract.Owner(&_Giftcardpaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Giftcardpaymaster *GiftcardpaymasterCallerSession) Owner() (common.Address, error) {
	return _Giftcardpaymaster.Contract.Owner(&_Giftcardpaymaster.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Giftcardpaymaster *GiftcardpaymasterCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Giftcardpaymaster.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Giftcardpaymaster *GiftcardpaymasterSession) PendingOwner() (common.Address, error) {
	return _Giftcardpaymaster.Contract.PendingOwner(&_Giftcardpaymaster.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Giftcardpaymaster *GiftcardpaymasterCallerSession) PendingOwner() (common.Address, error) {
	return _Giftcardpaymaster.Contract.PendingOwner(&_Giftcardpaymaster.CallOpts)
}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_Giftcardpaymaster *GiftcardpaymasterCaller) VerifyingSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Giftcardpaymaster.contract.Call(opts, &out, "verifyingSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_Giftcardpaymaster *GiftcardpaymasterSession) VerifyingSigner() (common.Address, error) {
	return _Giftcardpaymaster.Contract.VerifyingSigner(&_Giftcardpaymaster.CallOpts)
}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_Giftcardpaymaster *GiftcardpaymasterCallerSession) VerifyingSigner() (common.Address, error) {
	return _Giftcardpaymaster.Contract.VerifyingSigner(&_Giftcardpaymaster.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Giftcardpaymaster.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Giftcardpaymaster *GiftcardpaymasterSession) AcceptOwnership() (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.AcceptOwnership(&_Giftcardpaymaster.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.AcceptOwnership(&_Giftcardpaymaster.TransactOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Giftcardpaymaster.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Giftcardpaymaster *GiftcardpaymasterSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.AddStake(&_Giftcardpaymaster.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.AddStake(&_Giftcardpaymaster.TransactOpts, unstakeDelaySec)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Giftcardpaymaster.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Giftcardpaymaster *GiftcardpaymasterSession) Deposit() (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.Deposit(&_Giftcardpaymaster.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactorSession) Deposit() (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.Deposit(&_Giftcardpaymaster.TransactOpts)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _Giftcardpaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_Giftcardpaymaster *GiftcardpaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.PostOp(&_Giftcardpaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.PostOp(&_Giftcardpaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Giftcardpaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Giftcardpaymaster *GiftcardpaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.RenounceOwnership(&_Giftcardpaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.RenounceOwnership(&_Giftcardpaymaster.TransactOpts)
}

// SetVerifyingSigner is a paid mutator transaction binding the contract method 0xf5cba98c.
//
// Solidity: function setVerifyingSigner(address _newSigner) returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactor) SetVerifyingSigner(opts *bind.TransactOpts, _newSigner common.Address) (*types.Transaction, error) {
	return _Giftcardpaymaster.contract.Transact(opts, "setVerifyingSigner", _newSigner)
}

// SetVerifyingSigner is a paid mutator transaction binding the contract method 0xf5cba98c.
//
// Solidity: function setVerifyingSigner(address _newSigner) returns()
func (_Giftcardpaymaster *GiftcardpaymasterSession) SetVerifyingSigner(_newSigner common.Address) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.SetVerifyingSigner(&_Giftcardpaymaster.TransactOpts, _newSigner)
}

// SetVerifyingSigner is a paid mutator transaction binding the contract method 0xf5cba98c.
//
// Solidity: function setVerifyingSigner(address _newSigner) returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactorSession) SetVerifyingSigner(_newSigner common.Address) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.SetVerifyingSigner(&_Giftcardpaymaster.TransactOpts, _newSigner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Giftcardpaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Giftcardpaymaster *GiftcardpaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.TransferOwnership(&_Giftcardpaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.TransferOwnership(&_Giftcardpaymaster.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Giftcardpaymaster.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Giftcardpaymaster *GiftcardpaymasterSession) UnlockStake() (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.UnlockStake(&_Giftcardpaymaster.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.UnlockStake(&_Giftcardpaymaster.TransactOpts)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_Giftcardpaymaster *GiftcardpaymasterTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _Giftcardpaymaster.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_Giftcardpaymaster *GiftcardpaymasterSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.ValidatePaymasterUserOp(&_Giftcardpaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_Giftcardpaymaster *GiftcardpaymasterTransactorSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.ValidatePaymasterUserOp(&_Giftcardpaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _Giftcardpaymaster.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Giftcardpaymaster *GiftcardpaymasterSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.WithdrawStake(&_Giftcardpaymaster.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.WithdrawStake(&_Giftcardpaymaster.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Giftcardpaymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_Giftcardpaymaster *GiftcardpaymasterSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.WithdrawTo(&_Giftcardpaymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.WithdrawTo(&_Giftcardpaymaster.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Giftcardpaymaster.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Giftcardpaymaster *GiftcardpaymasterSession) Receive() (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.Receive(&_Giftcardpaymaster.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Giftcardpaymaster *GiftcardpaymasterTransactorSession) Receive() (*types.Transaction, error) {
	return _Giftcardpaymaster.Contract.Receive(&_Giftcardpaymaster.TransactOpts)
}

// GiftcardpaymasterOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Giftcardpaymaster contract.
type GiftcardpaymasterOwnershipTransferStartedIterator struct {
	Event *GiftcardpaymasterOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *GiftcardpaymasterOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GiftcardpaymasterOwnershipTransferStarted)
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
		it.Event = new(GiftcardpaymasterOwnershipTransferStarted)
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
func (it *GiftcardpaymasterOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GiftcardpaymasterOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GiftcardpaymasterOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Giftcardpaymaster contract.
type GiftcardpaymasterOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Giftcardpaymaster *GiftcardpaymasterFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GiftcardpaymasterOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Giftcardpaymaster.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GiftcardpaymasterOwnershipTransferStartedIterator{contract: _Giftcardpaymaster.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Giftcardpaymaster *GiftcardpaymasterFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *GiftcardpaymasterOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Giftcardpaymaster.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GiftcardpaymasterOwnershipTransferStarted)
				if err := _Giftcardpaymaster.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Giftcardpaymaster *GiftcardpaymasterFilterer) ParseOwnershipTransferStarted(log types.Log) (*GiftcardpaymasterOwnershipTransferStarted, error) {
	event := new(GiftcardpaymasterOwnershipTransferStarted)
	if err := _Giftcardpaymaster.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GiftcardpaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Giftcardpaymaster contract.
type GiftcardpaymasterOwnershipTransferredIterator struct {
	Event *GiftcardpaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GiftcardpaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GiftcardpaymasterOwnershipTransferred)
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
		it.Event = new(GiftcardpaymasterOwnershipTransferred)
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
func (it *GiftcardpaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GiftcardpaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GiftcardpaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the Giftcardpaymaster contract.
type GiftcardpaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Giftcardpaymaster *GiftcardpaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GiftcardpaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Giftcardpaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GiftcardpaymasterOwnershipTransferredIterator{contract: _Giftcardpaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Giftcardpaymaster *GiftcardpaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GiftcardpaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Giftcardpaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GiftcardpaymasterOwnershipTransferred)
				if err := _Giftcardpaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Giftcardpaymaster *GiftcardpaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*GiftcardpaymasterOwnershipTransferred, error) {
	event := new(GiftcardpaymasterOwnershipTransferred)
	if err := _Giftcardpaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GiftcardpaymasterVerifyingSignerUpdatedIterator is returned from FilterVerifyingSignerUpdated and is used to iterate over the raw logs and unpacked data for VerifyingSignerUpdated events raised by the Giftcardpaymaster contract.
type GiftcardpaymasterVerifyingSignerUpdatedIterator struct {
	Event *GiftcardpaymasterVerifyingSignerUpdated // Event containing the contract specifics and raw log

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
func (it *GiftcardpaymasterVerifyingSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GiftcardpaymasterVerifyingSignerUpdated)
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
		it.Event = new(GiftcardpaymasterVerifyingSignerUpdated)
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
func (it *GiftcardpaymasterVerifyingSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GiftcardpaymasterVerifyingSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GiftcardpaymasterVerifyingSignerUpdated represents a VerifyingSignerUpdated event raised by the Giftcardpaymaster contract.
type GiftcardpaymasterVerifyingSignerUpdated struct {
	OldSigner common.Address
	NewSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVerifyingSignerUpdated is a free log retrieval operation binding the contract event 0xdeee0b46266dc0734b009b2ca66e2ac28d5fd72d38506a3fe462865f00c5df4e.
//
// Solidity: event VerifyingSignerUpdated(address oldSigner, address newSigner)
func (_Giftcardpaymaster *GiftcardpaymasterFilterer) FilterVerifyingSignerUpdated(opts *bind.FilterOpts) (*GiftcardpaymasterVerifyingSignerUpdatedIterator, error) {

	logs, sub, err := _Giftcardpaymaster.contract.FilterLogs(opts, "VerifyingSignerUpdated")
	if err != nil {
		return nil, err
	}
	return &GiftcardpaymasterVerifyingSignerUpdatedIterator{contract: _Giftcardpaymaster.contract, event: "VerifyingSignerUpdated", logs: logs, sub: sub}, nil
}

// WatchVerifyingSignerUpdated is a free log subscription operation binding the contract event 0xdeee0b46266dc0734b009b2ca66e2ac28d5fd72d38506a3fe462865f00c5df4e.
//
// Solidity: event VerifyingSignerUpdated(address oldSigner, address newSigner)
func (_Giftcardpaymaster *GiftcardpaymasterFilterer) WatchVerifyingSignerUpdated(opts *bind.WatchOpts, sink chan<- *GiftcardpaymasterVerifyingSignerUpdated) (event.Subscription, error) {

	logs, sub, err := _Giftcardpaymaster.contract.WatchLogs(opts, "VerifyingSignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GiftcardpaymasterVerifyingSignerUpdated)
				if err := _Giftcardpaymaster.contract.UnpackLog(event, "VerifyingSignerUpdated", log); err != nil {
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

// ParseVerifyingSignerUpdated is a log parse operation binding the contract event 0xdeee0b46266dc0734b009b2ca66e2ac28d5fd72d38506a3fe462865f00c5df4e.
//
// Solidity: event VerifyingSignerUpdated(address oldSigner, address newSigner)
func (_Giftcardpaymaster *GiftcardpaymasterFilterer) ParseVerifyingSignerUpdated(log types.Log) (*GiftcardpaymasterVerifyingSignerUpdated, error) {
	event := new(GiftcardpaymasterVerifyingSignerUpdated)
	if err := _Giftcardpaymaster.contract.UnpackLog(event, "VerifyingSignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
