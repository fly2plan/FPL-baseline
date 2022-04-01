// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solidity

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
)

// ACKVerifierMetaData contains all meta data concerning the ACKVerifier contract.
var ACKVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[12]\",\"name\":\"input\",\"type\":\"uint256[12]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8cbac0fa": "verifyProof(uint256[2],uint256[2][2],uint256[2],uint256[12])",
	},
	Bin: "0x608060405234801561001057600080fd5b506114d8806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80638cbac0fa14610030575b600080fd5b61004361003e3660046112ef565b610057565b604051901515815260200160405180910390f35b60006100616110c8565b6040805180820182528751815260208089015181830152908352815160808101835287515181840190815288518301516060830152815282518084018452888301805151825251830151818401528183015283820152815180830183528651815286820151918101919091529082015260006100db61055c565b604080518082019091526000808252602082015283515191925090600080516020611483833981519152116101575760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61582d6774652d7072696d652d7100000000000000000060448201526064015b60405180910390fd5b825160200151600080516020611483833981519152116101b95760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61592d6774652d7072696d652d71000000000000000000604482015260640161014e565b602083015151516000805160206114838339815191521161021c5760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258302d6774652d7072696d652d710000000000000000604482015260640161014e565b602083810151015151600080516020611483833981519152116102815760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259302d6774652d7072696d652d710000000000000000604482015260640161014e565b602083810151510151600080516020611483833981519152116102e65760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258312d6774652d7072696d652d710000000000000000604482015260640161014e565b60208381015181015101516000805160206114838339815191521161034d5760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259312d6774652d7072696d652d710000000000000000604482015260640161014e565b604083015151600080516020611483833981519152116103af5760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63582d6774652d7072696d652d71000000000000000000604482015260640161014e565b600080516020611483833981519152836040015160200151106104145760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63592d6774652d7072696d652d71000000000000000000604482015260640161014e565b60005b600c811015610508577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018682600c8110610453576104536113cd565b6020020151106104a55760405162461bcd60e51b815260206004820152601f60248201527f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c6400604482015260640161014e565b6104f4826104ef85608001518460016104be91906113f9565b600d81106104ce576104ce6113cd565b60200201518985600c81106104e5576104e56113cd565b6020020151610c2c565b610cc2565b91508061050081611411565b915050610417565b5060808201515161051a908290610cc2565b905061055061052c8460000151610d5b565b84602001518460000151856020015185876040015189604001518960600151610df1565b98975050505050505050565b610564611119565b6040805180820182527f089c00453e4fcd1d631816f4483d6ca9902e54ad783c652868a5d14d5ea56e5e81527f045e28b8f55a3798d9f43822d2badf4d53de40d73f69a886bea839c3c620e7026020808301919091529083528151608080820184527f22c6be2d2b0cda274a18acbfac6be37c132bb59ce0d2d93d44075c73d82999958285019081527f177297c2c9f7c3636cbc33ab53a37168df5129dd36b2fd14ead87c6ab456b3c2606080850191909152908352845180860186527f11257668b989a06ae285b2bc8476945a8e4cf691bedc65c7bff277e76e59a46a81527f095ff7a5f28e9c6ad8ecc2514bbda0fdfd2f5665194c7bdc46a06c1145e63411818601528385015285840192909252835180820185527f1345d8bfb8ff3c4a88a5ce024b931b6977cd5f81b3c0a8aa9a28a6894e7698848186019081527f3013c53c6734a262ab12f58c43b2a0c7ad88c0cd870a7eeba14c998a52317192828501528152845180860186527f2d49afd96218c692dd8fcd32b6de1f31d31e2fefe516a2ce5646385c3ab74f9c81527f2d029bd938901741db4da7a74ead28def2dabfb3921d38ca32e1574e694acdad818601528185015285850152835180820185527f2ca44582807962dd0468c4b429ade4059a61c1787546b42e9d95b2304ad78fa28186019081527f28fdcbea5677e3169b784740b0c8dd9d5bd37e906fa9e44826664858e8f6b6c0828501528152845180860186527f0d592138056c9f63e0e89338246339f6de751f8f85508d26b4786e023701e2e581527f2fbdb846824238af35ff627a2a4bc4d5bb912cbcf31a4f0389f3895861eb5676818601528185015282860152835180850185527f0f3cd11d103fc84730b5f9b0e1af8ac9a630cb609acf9d5717d61964087ec18981527f11ede8ba1a59832efe48c5a591e35b87b7618a5a409e4bfaf3b523c1d037d1fe81850152818601805191909152845180860186527f24579bc43a924ac0a98194306bf0b7eac8db043ab3b76d7171ca64dadc90607c81527f20ff1b6360770023dd735ad3b7a04f5f84da2b53eb7215623ca18aca0df0e2e0818601528151850152845180860186527f2c1b3de426a4d75cfca0a5ec47001bce85e5bbc538c6b2150a0b9aa49fc080c481527f10cefaa05c3d38149f0cc6594ed771328676bb88553b0e54a9bc0363efafb341818601528151860152845180860186527f0ab1d28c884a1c3c05b93828f08f069fbd9ad970005eb424af6127607eb3711381527f187290aa5457d20ea9806535c24a43a6b30802ce3703de31c49ffe624a69605b81860152815190930192909252835180850185527f1415430542dc8292258a2cc4c9fdb7bb630e93e279d9896603aab3471cb2c8f281527f0d1cae16f40dc91aa88747d75cf4961c130ad9eec77687117985f061cec6ee6481850152825190910152825180840184527f20ca68bf1fa677f9712f8efee1ceb92e79d53eba8f381fb4de2212033a659f4f81527f082b8f3b94036afc78202a62f154172681027d24a2f7f2212f2cb9665cdc175081840152815160a00152825180840184527f0acec4f80fdbddad15ae7d8ad2bb900d3ca43b645aeea8e4519e2e223357ead181527f20ff5f5f41ed656a9db36a756c6a866ee44652c3c9cb039e850b818ec9bb6a5e81840152815160c00152825180840184527f265df053874ee47bef62434144b5ac1dcb43cfbcee338b9fb9582ef46c85cdb581527f1b80610a000eeeddb608e24ef75156147720157df2f5f3b74786f1c742a0231c81840152815160e00152825180840184527f2d86050f52b5f20eb685e4916edcb95b74d7d7eb6d18a3ff224f34f8bbc3d5a881527f26ac9891d170de1b19552cdd3ef9f7beec1b1fa1a1b82be0e477369ca217e8248184015281516101000152825180840184527f2f99c00aa274a310312f69872d2a9645e901c8fa881b00dd48d3af7d93d4f52281527f0a6f4b784076aa1b20773ad2a0deb22eb28997f6260324ac736a36c7a0b5b4128184015281516101200152825180840184527f0bef8cec7737ff340d49cc9c8f6e00f5ff321207685df71460a20cdfba7ec73a81527f2258876a26b5c8d972842471f685bdb5617a10d79bc53664649b67fa69a3b7e98184015281516101400152825180840184527f0c3cb96a63eb726da5c43c797486480e7bb8f9b3e6fd29829f9d19d05282cab081527f2707444b02739a9c7743f0cd32aff51fdf01e6523dfc5c20828481a541770070818401528151610160015282518084019093527f09b5c8714877253dfc0ca4974687060db03b02c564588ae7ea302db499f5317283527f22c4f0d12dd1a889f3090600ad78fa2eb76ee0d5b75ac4398f40ce95546f28b49183019190915251610180015290565b6040805180820190915260008082526020820152610c4861116a565b835181526020808501519082015260408101839052600060608360808460076107d05a03fa90508080610c7757fe5b5080610cba5760405162461bcd60e51b81526020600482015260126024820152711c185a5c9a5b99cb5b5d5b0b59985a5b195960721b604482015260640161014e565b505092915050565b6040805180820190915260008082526020820152610cde611188565b8351815260208085015181830152835160408301528301516060808301919091526000908360c08460066107d05a03fa90508080610d1857fe5b5080610cba5760405162461bcd60e51b81526020600482015260126024820152711c185a5c9a5b99cb5859190b59985a5b195960721b604482015260640161014e565b60408051808201909152600080825260208201528151158015610d8057506020820151155b15610d9e575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206114838339815191528460200151610dd1919061142a565b610de99060008051602061148383398151915261144c565b905292915050565b60408051608080820183528a825260208083018a90528284018890526060808401879052845192830185528b83528282018a9052828501889052820185905283516018808252610320820190955260009491859190839082016103008036833701905050905060005b6004811015611045576000610e70826006611463565b9050858260048110610e8457610e846113cd565b60200201515183610e968360006113f9565b81518110610ea657610ea66113cd565b602002602001018181525050858260048110610ec457610ec46113cd565b60200201516020015183826001610edb91906113f9565b81518110610eeb57610eeb6113cd565b602002602001018181525050848260048110610f0957610f096113cd565b6020020151515183610f1c8360026113f9565b81518110610f2c57610f2c6113cd565b602002602001018181525050848260048110610f4a57610f4a6113cd565b6020020151516001602002015183610f638360036113f9565b81518110610f7357610f736113cd565b602002602001018181525050848260048110610f9157610f916113cd565b602002015160200151600060028110610fac57610fac6113cd565b602002015183610fbd8360046113f9565b81518110610fcd57610fcd6113cd565b602002602001018181525050848260048110610feb57610feb6113cd565b602002015160200151600160028110611006576110066113cd565b6020020151836110178360056113f9565b81518110611027576110276113cd565b6020908102919091010152508061103d81611411565b915050610e5a565b5061104e6111a6565b6000602082602086026020860160086107d05a03fa9050808061106d57fe5b50806110b35760405162461bcd60e51b81526020600482015260156024820152741c185a5c9a5b99cb5bdc18dbd9194b59985a5b1959605a1b604482015260640161014e565b505115159d9c50505050505050505050505050565b6040805160a0810190915260006060820181815260808301919091528152602081016110f26111c4565b8152602001611114604051806040016040528060008152602001600081525090565b905290565b6040805160e08101909152600060a0820181815260c08301919091528152602081016111436111c4565b81526020016111506111c4565b815260200161115d6111c4565b81526020016111146111e4565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806111d761121e565b815260200161111461121e565b604051806101a00160405280600d905b60408051808201909152600080825260208201528152602001906001900390816111f45790505090565b60405180604001604052806002906020820280368337509192915050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156112755761127561123c565b60405290565b604051610180810167ffffffffffffffff811182821017156112755761127561123c565b600082601f8301126112b057600080fd5b6112b8611252565b8060408401858111156112ca57600080fd5b845b818110156112e45780358452602093840193016112cc565b509095945050505050565b60008060008061028080868803121561130757600080fd5b611311878761129f565b9450604087605f88011261132457600080fd5b61132c611252565b8060c089018a81111561133e57600080fd5b838a015b81811015611363576113548c8261129f565b84526020909301928401611342565b508197506113718b8261129f565b9650505050508661011f87011261138757600080fd5b61138f61127b565b9086019080888311156113a157600080fd5b61010088015b838110156113bf5780358352602092830192016113a7565b509598949750929550505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000821982111561140c5761140c6113e3565b500190565b600060018201611423576114236113e3565b5060010190565b60008261144757634e487b7160e01b600052601260045260246000fd5b500690565b60008282101561145e5761145e6113e3565b500390565b600081600019048311821515161561147d5761147d6113e3565b50029056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a26469706673582212200b6ce314e227fe31a03305c4a0607c3eed145082f8cc4f7af281b7f4ed5b089564736f6c634300080d0033",
}

// ACKVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ACKVerifierMetaData.ABI instead.
var ACKVerifierABI = ACKVerifierMetaData.ABI

// Deprecated: Use ACKVerifierMetaData.Sigs instead.
// ACKVerifierFuncSigs maps the 4-byte function signature to its string representation.
var ACKVerifierFuncSigs = ACKVerifierMetaData.Sigs

// ACKVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ACKVerifierMetaData.Bin instead.
var ACKVerifierBin = ACKVerifierMetaData.Bin

// DeployACKVerifier deploys a new Ethereum contract, binding an instance of ACKVerifier to it.
func DeployACKVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ACKVerifier, error) {
	parsed, err := ACKVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ACKVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ACKVerifier{ACKVerifierCaller: ACKVerifierCaller{contract: contract}, ACKVerifierTransactor: ACKVerifierTransactor{contract: contract}, ACKVerifierFilterer: ACKVerifierFilterer{contract: contract}}, nil
}

// ACKVerifier is an auto generated Go binding around an Ethereum contract.
type ACKVerifier struct {
	ACKVerifierCaller     // Read-only binding to the contract
	ACKVerifierTransactor // Write-only binding to the contract
	ACKVerifierFilterer   // Log filterer for contract events
}

// ACKVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ACKVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACKVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ACKVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACKVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ACKVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACKVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ACKVerifierSession struct {
	Contract     *ACKVerifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ACKVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ACKVerifierCallerSession struct {
	Contract *ACKVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ACKVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ACKVerifierTransactorSession struct {
	Contract     *ACKVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ACKVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ACKVerifierRaw struct {
	Contract *ACKVerifier // Generic contract binding to access the raw methods on
}

// ACKVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ACKVerifierCallerRaw struct {
	Contract *ACKVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ACKVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ACKVerifierTransactorRaw struct {
	Contract *ACKVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewACKVerifier creates a new instance of ACKVerifier, bound to a specific deployed contract.
func NewACKVerifier(address common.Address, backend bind.ContractBackend) (*ACKVerifier, error) {
	contract, err := bindACKVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ACKVerifier{ACKVerifierCaller: ACKVerifierCaller{contract: contract}, ACKVerifierTransactor: ACKVerifierTransactor{contract: contract}, ACKVerifierFilterer: ACKVerifierFilterer{contract: contract}}, nil
}

// NewACKVerifierCaller creates a new read-only instance of ACKVerifier, bound to a specific deployed contract.
func NewACKVerifierCaller(address common.Address, caller bind.ContractCaller) (*ACKVerifierCaller, error) {
	contract, err := bindACKVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ACKVerifierCaller{contract: contract}, nil
}

// NewACKVerifierTransactor creates a new write-only instance of ACKVerifier, bound to a specific deployed contract.
func NewACKVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ACKVerifierTransactor, error) {
	contract, err := bindACKVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ACKVerifierTransactor{contract: contract}, nil
}

// NewACKVerifierFilterer creates a new log filterer instance of ACKVerifier, bound to a specific deployed contract.
func NewACKVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ACKVerifierFilterer, error) {
	contract, err := bindACKVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ACKVerifierFilterer{contract: contract}, nil
}

// bindACKVerifier binds a generic wrapper to an already deployed contract.
func bindACKVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ACKVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACKVerifier *ACKVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACKVerifier.Contract.ACKVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACKVerifier *ACKVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACKVerifier.Contract.ACKVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACKVerifier *ACKVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACKVerifier.Contract.ACKVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACKVerifier *ACKVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACKVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACKVerifier *ACKVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACKVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACKVerifier *ACKVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACKVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x8cbac0fa.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[12] input) view returns(bool r)
func (_ACKVerifier *ACKVerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [12]*big.Int) (bool, error) {
	var out []interface{}
	err := _ACKVerifier.contract.Call(opts, &out, "verifyProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x8cbac0fa.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[12] input) view returns(bool r)
func (_ACKVerifier *ACKVerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [12]*big.Int) (bool, error) {
	return _ACKVerifier.Contract.VerifyProof(&_ACKVerifier.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x8cbac0fa.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[12] input) view returns(bool r)
func (_ACKVerifier *ACKVerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [12]*big.Int) (bool, error) {
	return _ACKVerifier.Contract.VerifyProof(&_ACKVerifier.CallOpts, a, b, c, input)
}
