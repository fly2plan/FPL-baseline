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

// FPLVerifierMetaData contains all meta data concerning the FPLVerifier contract.
var FPLVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[6]\",\"name\":\"input\",\"type\":\"uint256[6]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f398789b": "verifyProof(uint256[2],uint256[2][2],uint256[2],uint256[6])",
	},
	Bin: "0x608060405234801561001057600080fd5b506112d3806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f398789b14610030575b600080fd5b61004361003e3660046110ea565b610057565b604051901515815260200160405180910390f35b6000610061610ec5565b6040805180820182528751815260208089015181830152908352815160808101835287515181840190815288518301516060830152815282518084018452888301805151825251830151818401528183015283820152815180830183528651815286820151918101919091529082015260006100db61055c565b60408051808201909152600080825260208201528351519192509060008051602061127e833981519152116101575760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61582d6774652d7072696d652d7100000000000000000060448201526064015b60405180910390fd5b82516020015160008051602061127e833981519152116101b95760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61592d6774652d7072696d652d71000000000000000000604482015260640161014e565b6020830151515160008051602061127e8339815191521161021c5760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258302d6774652d7072696d652d710000000000000000604482015260640161014e565b60208381015101515160008051602061127e833981519152116102815760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259302d6774652d7072696d652d710000000000000000604482015260640161014e565b60208381015151015160008051602061127e833981519152116102e65760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258312d6774652d7072696d652d710000000000000000604482015260640161014e565b602083810151810151015160008051602061127e8339815191521161034d5760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259312d6774652d7072696d652d710000000000000000604482015260640161014e565b60408301515160008051602061127e833981519152116103af5760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63582d6774652d7072696d652d71000000000000000000604482015260640161014e565b60008051602061127e833981519152836040015160200151106104145760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63592d6774652d7072696d652d71000000000000000000604482015260640161014e565b60005b6006811015610508577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001868260068110610453576104536111c8565b6020020151106104a55760405162461bcd60e51b815260206004820152601f60248201527f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c6400604482015260640161014e565b6104f4826104ef85608001518460016104be91906111f4565b600781106104ce576104ce6111c8565b60200201518985600681106104e5576104e56111c8565b6020020151610a29565b610abf565b9150806105008161120c565b915050610417565b5060808201515161051a908290610abf565b905061055061052c8460000151610b58565b84602001518460000151856020015185876040015189604001518960600151610bee565b98975050505050505050565b610564610f16565b6040805180820182527f2ab7cd61120b810c88bcee80fc1b83ad799bd26f73cc7ca693a3d5668245e4e581527f09b2d3a18a2b68611b1b311ef66be4cda92b36b9eb92f9ae71e0664a810dc4e66020808301919091529083528151608080820184527f12933edf6a4917512045cee2f83eed916e5527dfb0e39e54552063ea02c445ee8285019081527f1d4aceaa16740963a35b64cfdfbc666372c38d852b7f7e95aa30039904b89de7606080850191909152908352845180860186527f25040a2fd6d3e3d37576430c0b12d8ba23686c4a03a662300846de4440c2b75681527f22dd7c09c68e92fbf23bbdb42b17f563dddcaf6281de2887eba3a7b40eb4fae2818601528385015285840192909252835180820185527f197f404353dd1ef150f483009b5ee38bc8352f5b0c45f1a77bacfacd432c90868186019081527f24d00293bd40398e2dbdfb9332c296fb7b121ca03b64859b5486115cb4a57eca828501528152845180860186527f15347b653cd297cc23e404e7c4a6a632aaec6fb1f943a0f54c8f1b6864ac3e5e81527f2b5c600c9a6798cd8a4c412046087b9f5cb23913d3544f117701f1c499a4ed3b818601528185015285850152835180820185527f0d76fbf41f62fa2750698e8cf219de1e531df5efc1ecf08caead652fe47feacb8186019081527f25ef5ac80d28fb040f3f23eaa23f783420c7226a5a07125fdc0f07285d07424d828501528152845180860186527f2f804ed9c84544f70570afdf9fea4133eebd9c00ce341efd7ddd516f68ce0cc781527f0b44be2e2695baf9d2944868bbf7c0f9a54b7ac2451b0876ea2a6bed700c29c5818601528185015282860152835180850185527f24d6a900ed2c5c4dea3e67b5485236105b9b160a9cd4f90255488c15f5ffccdd81527f260f51ad5aa544fedde893c86b44eb10d8e9282c8a12787f0984e263b2d71ac881850152818601805191909152845180860186527f1116a423a7dd761f41dd67960ecba4a0211b3db783086f59a33405a608e5f08e81527f150211bac5eaa9243ee881f5a54037ef655079776761f5103b6e4d88a0eee928818601528151850152845180860186527f1c75321b7133b3e976204e97dd3cb6896d75d6c661997cb92dbd9d8e6e56322681527f1fb42ad7cfc002b96e4ee50528d7e2761a8dbd9a996db70666e2b32bb395f093818601528151860152845180860186527f2dafb9c064286c039aa6aac05d19c66ace205b712c975bfa581fd439471cefe581527f26a6a160d06c96377c55695434787ced7e3fc87acbe886753767f508f9e5620381860152815190930192909252835180850185527f24504401b241dd415f120692665aa26a64531a04540510abd43db13d2661e6c581527f29ab65b2fe01a7a7e2b7c09be1b005924005e6649887d5cc13eed6ed17d41aee81850152825190910152825180840184527f076758c1bab1f924da54e3a7ce74e66c26a4f941680166bc785df03afdac0d2481527f26ecf6eb5ea77f419cd0f0f0090792a9167d5884a38db29eac5dd19f8f4ab9a681840152815160a0015282518084019093527f0d2c9315f84e57978bcba2a30fa2dd1d545afedeb2a21d48e653220ded4a1c9183527f019bcd552f8bf28aec61f90082802baeb6966e527d93f0fe6241de6999e89f00918301919091525160c0015290565b6040805180820190915260008082526020820152610a45610f67565b835181526020808501519082015260408101839052600060608360808460076107d05a03fa90508080610a7457fe5b5080610ab75760405162461bcd60e51b81526020600482015260126024820152711c185a5c9a5b99cb5b5d5b0b59985a5b195960721b604482015260640161014e565b505092915050565b6040805180820190915260008082526020820152610adb610f85565b8351815260208085015181830152835160408301528301516060808301919091526000908360c08460066107d05a03fa90508080610b1557fe5b5080610ab75760405162461bcd60e51b81526020600482015260126024820152711c185a5c9a5b99cb5859190b59985a5b195960721b604482015260640161014e565b60408051808201909152600080825260208201528151158015610b7d57506020820151155b15610b9b575050604080518082019091526000808252602082015290565b60405180604001604052808360000151815260200160008051602061127e8339815191528460200151610bce9190611225565b610be69060008051602061127e833981519152611247565b905292915050565b60408051608080820183528a825260208083018a90528284018890526060808401879052845192830185528b83528282018a9052828501889052820185905283516018808252610320820190955260009491859190839082016103008036833701905050905060005b6004811015610e42576000610c6d82600661125e565b9050858260048110610c8157610c816111c8565b60200201515183610c938360006111f4565b81518110610ca357610ca36111c8565b602002602001018181525050858260048110610cc157610cc16111c8565b60200201516020015183826001610cd891906111f4565b81518110610ce857610ce86111c8565b602002602001018181525050848260048110610d0657610d066111c8565b6020020151515183610d198360026111f4565b81518110610d2957610d296111c8565b602002602001018181525050848260048110610d4757610d476111c8565b6020020151516001602002015183610d608360036111f4565b81518110610d7057610d706111c8565b602002602001018181525050848260048110610d8e57610d8e6111c8565b602002015160200151600060028110610da957610da96111c8565b602002015183610dba8360046111f4565b81518110610dca57610dca6111c8565b602002602001018181525050848260048110610de857610de86111c8565b602002015160200151600160028110610e0357610e036111c8565b602002015183610e148360056111f4565b81518110610e2457610e246111c8565b60209081029190910101525080610e3a8161120c565b915050610c57565b50610e4b610fa3565b6000602082602086026020860160086107d05a03fa90508080610e6a57fe5b5080610eb05760405162461bcd60e51b81526020600482015260156024820152741c185a5c9a5b99cb5bdc18dbd9194b59985a5b1959605a1b604482015260640161014e565b505115159d9c50505050505050505050505050565b6040805160a081019091526000606082018181526080830191909152815260208101610eef610fc1565b8152602001610f11604051806040016040528060008152602001600081525090565b905290565b6040805160e08101909152600060a0820181815260c0830191909152815260208101610f40610fc1565b8152602001610f4d610fc1565b8152602001610f5a610fc1565b8152602001610f11610fe1565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060400160405280610fd461101a565b8152602001610f1161101a565b6040518060e001604052806007905b6040805180820190915260008082526020820152815260200190600190039081610ff05790505090565b60405180604001604052806002906020820280368337509192915050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561107157611071611038565b60405290565b60405160c0810167ffffffffffffffff8111828210171561107157611071611038565b600082601f8301126110ab57600080fd5b6110b361104e565b8060408401858111156110c557600080fd5b845b818110156110df5780358452602093840193016110c7565b509095945050505050565b6000806000806101c080868803121561110257600080fd5b61110c878761109a565b9450604087605f88011261111f57600080fd5b61112761104e565b8060c089018a81111561113957600080fd5b838a015b8181101561115e5761114f8c8261109a565b8452602090930192840161113d565b5081975061116c8b8261109a565b9650505050508661011f87011261118257600080fd5b61118a611077565b90860190808883111561119c57600080fd5b61010088015b838110156111ba5780358352602092830192016111a2565b509598949750929550505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008219821115611207576112076111de565b500190565b60006001820161121e5761121e6111de565b5060010190565b60008261124257634e487b7160e01b600052601260045260246000fd5b500690565b600082821015611259576112596111de565b500390565b6000816000190483118215151615611278576112786111de565b50029056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220b4c793e264dfc6ef38c3b62bd65f5f794c72ed247953a7d78e4be8e48cddc7ec64736f6c634300080d0033",
}

// FPLVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use FPLVerifierMetaData.ABI instead.
var FPLVerifierABI = FPLVerifierMetaData.ABI

// Deprecated: Use FPLVerifierMetaData.Sigs instead.
// FPLVerifierFuncSigs maps the 4-byte function signature to its string representation.
var FPLVerifierFuncSigs = FPLVerifierMetaData.Sigs

// FPLVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FPLVerifierMetaData.Bin instead.
var FPLVerifierBin = FPLVerifierMetaData.Bin

// DeployFPLVerifier deploys a new Ethereum contract, binding an instance of FPLVerifier to it.
func DeployFPLVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FPLVerifier, error) {
	parsed, err := FPLVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FPLVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FPLVerifier{FPLVerifierCaller: FPLVerifierCaller{contract: contract}, FPLVerifierTransactor: FPLVerifierTransactor{contract: contract}, FPLVerifierFilterer: FPLVerifierFilterer{contract: contract}}, nil
}

// FPLVerifier is an auto generated Go binding around an Ethereum contract.
type FPLVerifier struct {
	FPLVerifierCaller     // Read-only binding to the contract
	FPLVerifierTransactor // Write-only binding to the contract
	FPLVerifierFilterer   // Log filterer for contract events
}

// FPLVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type FPLVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FPLVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FPLVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FPLVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FPLVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FPLVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FPLVerifierSession struct {
	Contract     *FPLVerifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FPLVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FPLVerifierCallerSession struct {
	Contract *FPLVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FPLVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FPLVerifierTransactorSession struct {
	Contract     *FPLVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FPLVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type FPLVerifierRaw struct {
	Contract *FPLVerifier // Generic contract binding to access the raw methods on
}

// FPLVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FPLVerifierCallerRaw struct {
	Contract *FPLVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// FPLVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FPLVerifierTransactorRaw struct {
	Contract *FPLVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFPLVerifier creates a new instance of FPLVerifier, bound to a specific deployed contract.
func NewFPLVerifier(address common.Address, backend bind.ContractBackend) (*FPLVerifier, error) {
	contract, err := bindFPLVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FPLVerifier{FPLVerifierCaller: FPLVerifierCaller{contract: contract}, FPLVerifierTransactor: FPLVerifierTransactor{contract: contract}, FPLVerifierFilterer: FPLVerifierFilterer{contract: contract}}, nil
}

// NewFPLVerifierCaller creates a new read-only instance of FPLVerifier, bound to a specific deployed contract.
func NewFPLVerifierCaller(address common.Address, caller bind.ContractCaller) (*FPLVerifierCaller, error) {
	contract, err := bindFPLVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FPLVerifierCaller{contract: contract}, nil
}

// NewFPLVerifierTransactor creates a new write-only instance of FPLVerifier, bound to a specific deployed contract.
func NewFPLVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*FPLVerifierTransactor, error) {
	contract, err := bindFPLVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FPLVerifierTransactor{contract: contract}, nil
}

// NewFPLVerifierFilterer creates a new log filterer instance of FPLVerifier, bound to a specific deployed contract.
func NewFPLVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*FPLVerifierFilterer, error) {
	contract, err := bindFPLVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FPLVerifierFilterer{contract: contract}, nil
}

// bindFPLVerifier binds a generic wrapper to an already deployed contract.
func bindFPLVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FPLVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FPLVerifier *FPLVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FPLVerifier.Contract.FPLVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FPLVerifier *FPLVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FPLVerifier.Contract.FPLVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FPLVerifier *FPLVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FPLVerifier.Contract.FPLVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FPLVerifier *FPLVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FPLVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FPLVerifier *FPLVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FPLVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FPLVerifier *FPLVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FPLVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf398789b.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[6] input) view returns(bool r)
func (_FPLVerifier *FPLVerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [6]*big.Int) (bool, error) {
	var out []interface{}
	err := _FPLVerifier.contract.Call(opts, &out, "verifyProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0xf398789b.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[6] input) view returns(bool r)
func (_FPLVerifier *FPLVerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [6]*big.Int) (bool, error) {
	return _FPLVerifier.Contract.VerifyProof(&_FPLVerifier.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf398789b.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[6] input) view returns(bool r)
func (_FPLVerifier *FPLVerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [6]*big.Int) (bool, error) {
	return _FPLVerifier.Contract.VerifyProof(&_FPLVerifier.CallOpts, a, b, c, input)
}
