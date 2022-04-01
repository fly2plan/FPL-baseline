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

// IACKVerifierMetaData contains all meta data concerning the IACKVerifier contract.
var IACKVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[12]\",\"name\":\"input\",\"type\":\"uint256[12]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8cbac0fa": "verifyProof(uint256[2],uint256[2][2],uint256[2],uint256[12])",
	},
}

// IACKVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IACKVerifierMetaData.ABI instead.
var IACKVerifierABI = IACKVerifierMetaData.ABI

// Deprecated: Use IACKVerifierMetaData.Sigs instead.
// IACKVerifierFuncSigs maps the 4-byte function signature to its string representation.
var IACKVerifierFuncSigs = IACKVerifierMetaData.Sigs

// IACKVerifier is an auto generated Go binding around an Ethereum contract.
type IACKVerifier struct {
	IACKVerifierCaller     // Read-only binding to the contract
	IACKVerifierTransactor // Write-only binding to the contract
	IACKVerifierFilterer   // Log filterer for contract events
}

// IACKVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IACKVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IACKVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IACKVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IACKVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IACKVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IACKVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IACKVerifierSession struct {
	Contract     *IACKVerifier     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IACKVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IACKVerifierCallerSession struct {
	Contract *IACKVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IACKVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IACKVerifierTransactorSession struct {
	Contract     *IACKVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IACKVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IACKVerifierRaw struct {
	Contract *IACKVerifier // Generic contract binding to access the raw methods on
}

// IACKVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IACKVerifierCallerRaw struct {
	Contract *IACKVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IACKVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IACKVerifierTransactorRaw struct {
	Contract *IACKVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIACKVerifier creates a new instance of IACKVerifier, bound to a specific deployed contract.
func NewIACKVerifier(address common.Address, backend bind.ContractBackend) (*IACKVerifier, error) {
	contract, err := bindIACKVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IACKVerifier{IACKVerifierCaller: IACKVerifierCaller{contract: contract}, IACKVerifierTransactor: IACKVerifierTransactor{contract: contract}, IACKVerifierFilterer: IACKVerifierFilterer{contract: contract}}, nil
}

// NewIACKVerifierCaller creates a new read-only instance of IACKVerifier, bound to a specific deployed contract.
func NewIACKVerifierCaller(address common.Address, caller bind.ContractCaller) (*IACKVerifierCaller, error) {
	contract, err := bindIACKVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IACKVerifierCaller{contract: contract}, nil
}

// NewIACKVerifierTransactor creates a new write-only instance of IACKVerifier, bound to a specific deployed contract.
func NewIACKVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IACKVerifierTransactor, error) {
	contract, err := bindIACKVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IACKVerifierTransactor{contract: contract}, nil
}

// NewIACKVerifierFilterer creates a new log filterer instance of IACKVerifier, bound to a specific deployed contract.
func NewIACKVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IACKVerifierFilterer, error) {
	contract, err := bindIACKVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IACKVerifierFilterer{contract: contract}, nil
}

// bindIACKVerifier binds a generic wrapper to an already deployed contract.
func bindIACKVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IACKVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IACKVerifier *IACKVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IACKVerifier.Contract.IACKVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IACKVerifier *IACKVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IACKVerifier.Contract.IACKVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IACKVerifier *IACKVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IACKVerifier.Contract.IACKVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IACKVerifier *IACKVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IACKVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IACKVerifier *IACKVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IACKVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IACKVerifier *IACKVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IACKVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a paid mutator transaction binding the contract method 0x8cbac0fa.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[12] input) returns(bool r)
func (_IACKVerifier *IACKVerifierTransactor) VerifyProof(opts *bind.TransactOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [12]*big.Int) (*types.Transaction, error) {
	return _IACKVerifier.contract.Transact(opts, "verifyProof", a, b, c, input)
}

// VerifyProof is a paid mutator transaction binding the contract method 0x8cbac0fa.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[12] input) returns(bool r)
func (_IACKVerifier *IACKVerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [12]*big.Int) (*types.Transaction, error) {
	return _IACKVerifier.Contract.VerifyProof(&_IACKVerifier.TransactOpts, a, b, c, input)
}

// VerifyProof is a paid mutator transaction binding the contract method 0x8cbac0fa.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[12] input) returns(bool r)
func (_IACKVerifier *IACKVerifierTransactorSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [12]*big.Int) (*types.Transaction, error) {
	return _IACKVerifier.Contract.VerifyProof(&_IACKVerifier.TransactOpts, a, b, c, input)
}

// IFPLVerifierMetaData contains all meta data concerning the IFPLVerifier contract.
var IFPLVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[6]\",\"name\":\"input\",\"type\":\"uint256[6]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f398789b": "verifyProof(uint256[2],uint256[2][2],uint256[2],uint256[6])",
	},
}

// IFPLVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IFPLVerifierMetaData.ABI instead.
var IFPLVerifierABI = IFPLVerifierMetaData.ABI

// Deprecated: Use IFPLVerifierMetaData.Sigs instead.
// IFPLVerifierFuncSigs maps the 4-byte function signature to its string representation.
var IFPLVerifierFuncSigs = IFPLVerifierMetaData.Sigs

// IFPLVerifier is an auto generated Go binding around an Ethereum contract.
type IFPLVerifier struct {
	IFPLVerifierCaller     // Read-only binding to the contract
	IFPLVerifierTransactor // Write-only binding to the contract
	IFPLVerifierFilterer   // Log filterer for contract events
}

// IFPLVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFPLVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFPLVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFPLVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFPLVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFPLVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFPLVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFPLVerifierSession struct {
	Contract     *IFPLVerifier     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFPLVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFPLVerifierCallerSession struct {
	Contract *IFPLVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IFPLVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFPLVerifierTransactorSession struct {
	Contract     *IFPLVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IFPLVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFPLVerifierRaw struct {
	Contract *IFPLVerifier // Generic contract binding to access the raw methods on
}

// IFPLVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFPLVerifierCallerRaw struct {
	Contract *IFPLVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IFPLVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFPLVerifierTransactorRaw struct {
	Contract *IFPLVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFPLVerifier creates a new instance of IFPLVerifier, bound to a specific deployed contract.
func NewIFPLVerifier(address common.Address, backend bind.ContractBackend) (*IFPLVerifier, error) {
	contract, err := bindIFPLVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFPLVerifier{IFPLVerifierCaller: IFPLVerifierCaller{contract: contract}, IFPLVerifierTransactor: IFPLVerifierTransactor{contract: contract}, IFPLVerifierFilterer: IFPLVerifierFilterer{contract: contract}}, nil
}

// NewIFPLVerifierCaller creates a new read-only instance of IFPLVerifier, bound to a specific deployed contract.
func NewIFPLVerifierCaller(address common.Address, caller bind.ContractCaller) (*IFPLVerifierCaller, error) {
	contract, err := bindIFPLVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFPLVerifierCaller{contract: contract}, nil
}

// NewIFPLVerifierTransactor creates a new write-only instance of IFPLVerifier, bound to a specific deployed contract.
func NewIFPLVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IFPLVerifierTransactor, error) {
	contract, err := bindIFPLVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFPLVerifierTransactor{contract: contract}, nil
}

// NewIFPLVerifierFilterer creates a new log filterer instance of IFPLVerifier, bound to a specific deployed contract.
func NewIFPLVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IFPLVerifierFilterer, error) {
	contract, err := bindIFPLVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFPLVerifierFilterer{contract: contract}, nil
}

// bindIFPLVerifier binds a generic wrapper to an already deployed contract.
func bindIFPLVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFPLVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFPLVerifier *IFPLVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFPLVerifier.Contract.IFPLVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFPLVerifier *IFPLVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFPLVerifier.Contract.IFPLVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFPLVerifier *IFPLVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFPLVerifier.Contract.IFPLVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFPLVerifier *IFPLVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFPLVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFPLVerifier *IFPLVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFPLVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFPLVerifier *IFPLVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFPLVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a paid mutator transaction binding the contract method 0xf398789b.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[6] input) returns(bool r)
func (_IFPLVerifier *IFPLVerifierTransactor) VerifyProof(opts *bind.TransactOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [6]*big.Int) (*types.Transaction, error) {
	return _IFPLVerifier.contract.Transact(opts, "verifyProof", a, b, c, input)
}

// VerifyProof is a paid mutator transaction binding the contract method 0xf398789b.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[6] input) returns(bool r)
func (_IFPLVerifier *IFPLVerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [6]*big.Int) (*types.Transaction, error) {
	return _IFPLVerifier.Contract.VerifyProof(&_IFPLVerifier.TransactOpts, a, b, c, input)
}

// VerifyProof is a paid mutator transaction binding the contract method 0xf398789b.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[6] input) returns(bool r)
func (_IFPLVerifier *IFPLVerifierTransactorSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [6]*big.Int) (*types.Transaction, error) {
	return _IFPLVerifier.Contract.VerifyProof(&_IFPLVerifier.TransactOpts, a, b, c, input)
}

// IShieldMetaData contains all meta data concerning the IShield contract.
var IShieldMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getACKVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFPLVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[12]\",\"name\":\"input\",\"type\":\"uint256[12]\"},{\"internalType\":\"bytes32\",\"name\":\"_newCommitment\",\"type\":\"bytes32\"}],\"name\":\"verifyAndPushACK\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[6]\",\"name\":\"input\",\"type\":\"uint256[6]\"},{\"internalType\":\"bytes32\",\"name\":\"_newCommitment\",\"type\":\"bytes32\"}],\"name\":\"verifyAndPushFPL\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"76a7d786": "getACKVerifier()",
		"9c233ca1": "getFPLVerifier()",
		"741569fc": "verifyAndPushACK(uint256[2],uint256[2][2],uint256[2],uint256[12],bytes32)",
		"9f0efffd": "verifyAndPushFPL(uint256[2],uint256[2][2],uint256[2],uint256[6],bytes32)",
	},
}

// IShieldABI is the input ABI used to generate the binding from.
// Deprecated: Use IShieldMetaData.ABI instead.
var IShieldABI = IShieldMetaData.ABI

// Deprecated: Use IShieldMetaData.Sigs instead.
// IShieldFuncSigs maps the 4-byte function signature to its string representation.
var IShieldFuncSigs = IShieldMetaData.Sigs

// IShield is an auto generated Go binding around an Ethereum contract.
type IShield struct {
	IShieldCaller     // Read-only binding to the contract
	IShieldTransactor // Write-only binding to the contract
	IShieldFilterer   // Log filterer for contract events
}

// IShieldCaller is an auto generated read-only Go binding around an Ethereum contract.
type IShieldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IShieldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IShieldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IShieldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IShieldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IShieldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IShieldSession struct {
	Contract     *IShield          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IShieldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IShieldCallerSession struct {
	Contract *IShieldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IShieldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IShieldTransactorSession struct {
	Contract     *IShieldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IShieldRaw is an auto generated low-level Go binding around an Ethereum contract.
type IShieldRaw struct {
	Contract *IShield // Generic contract binding to access the raw methods on
}

// IShieldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IShieldCallerRaw struct {
	Contract *IShieldCaller // Generic read-only contract binding to access the raw methods on
}

// IShieldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IShieldTransactorRaw struct {
	Contract *IShieldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIShield creates a new instance of IShield, bound to a specific deployed contract.
func NewIShield(address common.Address, backend bind.ContractBackend) (*IShield, error) {
	contract, err := bindIShield(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IShield{IShieldCaller: IShieldCaller{contract: contract}, IShieldTransactor: IShieldTransactor{contract: contract}, IShieldFilterer: IShieldFilterer{contract: contract}}, nil
}

// NewIShieldCaller creates a new read-only instance of IShield, bound to a specific deployed contract.
func NewIShieldCaller(address common.Address, caller bind.ContractCaller) (*IShieldCaller, error) {
	contract, err := bindIShield(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IShieldCaller{contract: contract}, nil
}

// NewIShieldTransactor creates a new write-only instance of IShield, bound to a specific deployed contract.
func NewIShieldTransactor(address common.Address, transactor bind.ContractTransactor) (*IShieldTransactor, error) {
	contract, err := bindIShield(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IShieldTransactor{contract: contract}, nil
}

// NewIShieldFilterer creates a new log filterer instance of IShield, bound to a specific deployed contract.
func NewIShieldFilterer(address common.Address, filterer bind.ContractFilterer) (*IShieldFilterer, error) {
	contract, err := bindIShield(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IShieldFilterer{contract: contract}, nil
}

// bindIShield binds a generic wrapper to an already deployed contract.
func bindIShield(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IShieldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IShield *IShieldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IShield.Contract.IShieldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IShield *IShieldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IShield.Contract.IShieldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IShield *IShieldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IShield.Contract.IShieldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IShield *IShieldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IShield.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IShield *IShieldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IShield.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IShield *IShieldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IShield.Contract.contract.Transact(opts, method, params...)
}

// GetACKVerifier is a free data retrieval call binding the contract method 0x76a7d786.
//
// Solidity: function getACKVerifier() view returns(address)
func (_IShield *IShieldCaller) GetACKVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IShield.contract.Call(opts, &out, "getACKVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetACKVerifier is a free data retrieval call binding the contract method 0x76a7d786.
//
// Solidity: function getACKVerifier() view returns(address)
func (_IShield *IShieldSession) GetACKVerifier() (common.Address, error) {
	return _IShield.Contract.GetACKVerifier(&_IShield.CallOpts)
}

// GetACKVerifier is a free data retrieval call binding the contract method 0x76a7d786.
//
// Solidity: function getACKVerifier() view returns(address)
func (_IShield *IShieldCallerSession) GetACKVerifier() (common.Address, error) {
	return _IShield.Contract.GetACKVerifier(&_IShield.CallOpts)
}

// GetFPLVerifier is a free data retrieval call binding the contract method 0x9c233ca1.
//
// Solidity: function getFPLVerifier() view returns(address)
func (_IShield *IShieldCaller) GetFPLVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IShield.contract.Call(opts, &out, "getFPLVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFPLVerifier is a free data retrieval call binding the contract method 0x9c233ca1.
//
// Solidity: function getFPLVerifier() view returns(address)
func (_IShield *IShieldSession) GetFPLVerifier() (common.Address, error) {
	return _IShield.Contract.GetFPLVerifier(&_IShield.CallOpts)
}

// GetFPLVerifier is a free data retrieval call binding the contract method 0x9c233ca1.
//
// Solidity: function getFPLVerifier() view returns(address)
func (_IShield *IShieldCallerSession) GetFPLVerifier() (common.Address, error) {
	return _IShield.Contract.GetFPLVerifier(&_IShield.CallOpts)
}

// VerifyAndPushACK is a paid mutator transaction binding the contract method 0x741569fc.
//
// Solidity: function verifyAndPushACK(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[12] input, bytes32 _newCommitment) returns(bool)
func (_IShield *IShieldTransactor) VerifyAndPushACK(opts *bind.TransactOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [12]*big.Int, _newCommitment [32]byte) (*types.Transaction, error) {
	return _IShield.contract.Transact(opts, "verifyAndPushACK", a, b, c, input, _newCommitment)
}

// VerifyAndPushACK is a paid mutator transaction binding the contract method 0x741569fc.
//
// Solidity: function verifyAndPushACK(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[12] input, bytes32 _newCommitment) returns(bool)
func (_IShield *IShieldSession) VerifyAndPushACK(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [12]*big.Int, _newCommitment [32]byte) (*types.Transaction, error) {
	return _IShield.Contract.VerifyAndPushACK(&_IShield.TransactOpts, a, b, c, input, _newCommitment)
}

// VerifyAndPushACK is a paid mutator transaction binding the contract method 0x741569fc.
//
// Solidity: function verifyAndPushACK(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[12] input, bytes32 _newCommitment) returns(bool)
func (_IShield *IShieldTransactorSession) VerifyAndPushACK(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [12]*big.Int, _newCommitment [32]byte) (*types.Transaction, error) {
	return _IShield.Contract.VerifyAndPushACK(&_IShield.TransactOpts, a, b, c, input, _newCommitment)
}

// VerifyAndPushFPL is a paid mutator transaction binding the contract method 0x9f0efffd.
//
// Solidity: function verifyAndPushFPL(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[6] input, bytes32 _newCommitment) returns(bool)
func (_IShield *IShieldTransactor) VerifyAndPushFPL(opts *bind.TransactOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [6]*big.Int, _newCommitment [32]byte) (*types.Transaction, error) {
	return _IShield.contract.Transact(opts, "verifyAndPushFPL", a, b, c, input, _newCommitment)
}

// VerifyAndPushFPL is a paid mutator transaction binding the contract method 0x9f0efffd.
//
// Solidity: function verifyAndPushFPL(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[6] input, bytes32 _newCommitment) returns(bool)
func (_IShield *IShieldSession) VerifyAndPushFPL(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [6]*big.Int, _newCommitment [32]byte) (*types.Transaction, error) {
	return _IShield.Contract.VerifyAndPushFPL(&_IShield.TransactOpts, a, b, c, input, _newCommitment)
}

// VerifyAndPushFPL is a paid mutator transaction binding the contract method 0x9f0efffd.
//
// Solidity: function verifyAndPushFPL(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[6] input, bytes32 _newCommitment) returns(bool)
func (_IShield *IShieldTransactorSession) VerifyAndPushFPL(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [6]*big.Int, _newCommitment [32]byte) (*types.Transaction, error) {
	return _IShield.Contract.VerifyAndPushFPL(&_IShield.TransactOpts, a, b, c, input, _newCommitment)
}

// MerkleTreeSHA256MetaData contains all meta data concerning the MerkleTreeSHA256 contract.
var MerkleTreeSHA256MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_treeHeight\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"leafValue\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"NewLeaf\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minLeafIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"leafValues\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"NewLeaves\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes27\",\"name\":\"leftInput\",\"type\":\"bytes27\"},{\"indexed\":false,\"internalType\":\"bytes27\",\"name\":\"rightInput\",\"type\":\"bytes27\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"output\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodeIndex\",\"type\":\"uint256\"}],\"name\":\"Output\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"frontier\",\"outputs\":[{\"internalType\":\"bytes27\",\"name\":\"\",\"type\":\"bytes27\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leafCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeWidth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zero\",\"outputs\":[{\"internalType\":\"bytes27\",\"name\":\"\",\"type\":\"bytes27\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6e0c3fee": "frontier(uint256)",
		"d7b0fef1": "latestRoot()",
		"30e69fc3": "leafCount()",
		"01e3e915": "treeHeight()",
		"76c601b1": "treeWidth()",
		"bc1b392d": "zero()",
	},
	Bin: "0x6080604052600480546001600160d81b031916905534801561002057600080fd5b506040516102b93803806102b983398101604081905261003f91610058565b600081905561004f81600261016d565b60015550610180565b60006020828403121561006a57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600181815b808511156100c25781600019048211156100a8576100a8610071565b808516156100b557918102915b93841c939080029061008c565b509250929050565b6000826100d957506001610167565b816100e657506000610167565b81600181146100fc576002811461010657610122565b6001915050610167565b60ff84111561011757610117610071565b50506001821b610167565b5060208310610133831016604e8410600b8410161715610145575081810a610167565b61014f8383610087565b806000190482111561016357610163610071565b0290505b92915050565b600061017983836100ca565b9392505050565b61012a8061018f6000396000f3fe6080604052348015600f57600080fd5b5060043610605a5760003560e01c806301e3e91514605f57806330e69fc314607a5780636e0c3fee14608257806376c601b11460a7578063bc1b392d1460af578063d7b0fef11460bb575b600080fd5b606760005481565b6040519081526020015b60405180910390f35b606760035481565b6091608d36600460dc565b60c3565b60405164ffffffffff1990911681526020016071565b606760015481565b60045460919060281b81565b606760025481565b6005816021811060d257600080fd5b015460281b905081565b60006020828403121560ed57600080fd5b503591905056fea264697066735822122098def4d00c13e638b9e02895cfb9b76423f2aa4532a7a22184b6274744724f1464736f6c634300080d0033",
}

// MerkleTreeSHA256ABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleTreeSHA256MetaData.ABI instead.
var MerkleTreeSHA256ABI = MerkleTreeSHA256MetaData.ABI

// Deprecated: Use MerkleTreeSHA256MetaData.Sigs instead.
// MerkleTreeSHA256FuncSigs maps the 4-byte function signature to its string representation.
var MerkleTreeSHA256FuncSigs = MerkleTreeSHA256MetaData.Sigs

// MerkleTreeSHA256Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleTreeSHA256MetaData.Bin instead.
var MerkleTreeSHA256Bin = MerkleTreeSHA256MetaData.Bin

// DeployMerkleTreeSHA256 deploys a new Ethereum contract, binding an instance of MerkleTreeSHA256 to it.
func DeployMerkleTreeSHA256(auth *bind.TransactOpts, backend bind.ContractBackend, _treeHeight *big.Int) (common.Address, *types.Transaction, *MerkleTreeSHA256, error) {
	parsed, err := MerkleTreeSHA256MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleTreeSHA256Bin), backend, _treeHeight)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleTreeSHA256{MerkleTreeSHA256Caller: MerkleTreeSHA256Caller{contract: contract}, MerkleTreeSHA256Transactor: MerkleTreeSHA256Transactor{contract: contract}, MerkleTreeSHA256Filterer: MerkleTreeSHA256Filterer{contract: contract}}, nil
}

// MerkleTreeSHA256 is an auto generated Go binding around an Ethereum contract.
type MerkleTreeSHA256 struct {
	MerkleTreeSHA256Caller     // Read-only binding to the contract
	MerkleTreeSHA256Transactor // Write-only binding to the contract
	MerkleTreeSHA256Filterer   // Log filterer for contract events
}

// MerkleTreeSHA256Caller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleTreeSHA256Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeSHA256Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTreeSHA256Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeSHA256Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleTreeSHA256Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeSHA256Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleTreeSHA256Session struct {
	Contract     *MerkleTreeSHA256 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleTreeSHA256CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleTreeSHA256CallerSession struct {
	Contract *MerkleTreeSHA256Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MerkleTreeSHA256TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTreeSHA256TransactorSession struct {
	Contract     *MerkleTreeSHA256Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MerkleTreeSHA256Raw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleTreeSHA256Raw struct {
	Contract *MerkleTreeSHA256 // Generic contract binding to access the raw methods on
}

// MerkleTreeSHA256CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleTreeSHA256CallerRaw struct {
	Contract *MerkleTreeSHA256Caller // Generic read-only contract binding to access the raw methods on
}

// MerkleTreeSHA256TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTreeSHA256TransactorRaw struct {
	Contract *MerkleTreeSHA256Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleTreeSHA256 creates a new instance of MerkleTreeSHA256, bound to a specific deployed contract.
func NewMerkleTreeSHA256(address common.Address, backend bind.ContractBackend) (*MerkleTreeSHA256, error) {
	contract, err := bindMerkleTreeSHA256(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeSHA256{MerkleTreeSHA256Caller: MerkleTreeSHA256Caller{contract: contract}, MerkleTreeSHA256Transactor: MerkleTreeSHA256Transactor{contract: contract}, MerkleTreeSHA256Filterer: MerkleTreeSHA256Filterer{contract: contract}}, nil
}

// NewMerkleTreeSHA256Caller creates a new read-only instance of MerkleTreeSHA256, bound to a specific deployed contract.
func NewMerkleTreeSHA256Caller(address common.Address, caller bind.ContractCaller) (*MerkleTreeSHA256Caller, error) {
	contract, err := bindMerkleTreeSHA256(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeSHA256Caller{contract: contract}, nil
}

// NewMerkleTreeSHA256Transactor creates a new write-only instance of MerkleTreeSHA256, bound to a specific deployed contract.
func NewMerkleTreeSHA256Transactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTreeSHA256Transactor, error) {
	contract, err := bindMerkleTreeSHA256(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeSHA256Transactor{contract: contract}, nil
}

// NewMerkleTreeSHA256Filterer creates a new log filterer instance of MerkleTreeSHA256, bound to a specific deployed contract.
func NewMerkleTreeSHA256Filterer(address common.Address, filterer bind.ContractFilterer) (*MerkleTreeSHA256Filterer, error) {
	contract, err := bindMerkleTreeSHA256(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeSHA256Filterer{contract: contract}, nil
}

// bindMerkleTreeSHA256 binds a generic wrapper to an already deployed contract.
func bindMerkleTreeSHA256(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleTreeSHA256ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTreeSHA256 *MerkleTreeSHA256Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTreeSHA256.Contract.MerkleTreeSHA256Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTreeSHA256 *MerkleTreeSHA256Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTreeSHA256.Contract.MerkleTreeSHA256Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTreeSHA256 *MerkleTreeSHA256Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTreeSHA256.Contract.MerkleTreeSHA256Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTreeSHA256 *MerkleTreeSHA256CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTreeSHA256.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTreeSHA256 *MerkleTreeSHA256TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTreeSHA256.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTreeSHA256 *MerkleTreeSHA256TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTreeSHA256.Contract.contract.Transact(opts, method, params...)
}

// Frontier is a free data retrieval call binding the contract method 0x6e0c3fee.
//
// Solidity: function frontier(uint256 ) view returns(bytes27)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Caller) Frontier(opts *bind.CallOpts, arg0 *big.Int) ([27]byte, error) {
	var out []interface{}
	err := _MerkleTreeSHA256.contract.Call(opts, &out, "frontier", arg0)

	if err != nil {
		return *new([27]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([27]byte)).(*[27]byte)

	return out0, err

}

// Frontier is a free data retrieval call binding the contract method 0x6e0c3fee.
//
// Solidity: function frontier(uint256 ) view returns(bytes27)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Session) Frontier(arg0 *big.Int) ([27]byte, error) {
	return _MerkleTreeSHA256.Contract.Frontier(&_MerkleTreeSHA256.CallOpts, arg0)
}

// Frontier is a free data retrieval call binding the contract method 0x6e0c3fee.
//
// Solidity: function frontier(uint256 ) view returns(bytes27)
func (_MerkleTreeSHA256 *MerkleTreeSHA256CallerSession) Frontier(arg0 *big.Int) ([27]byte, error) {
	return _MerkleTreeSHA256.Contract.Frontier(&_MerkleTreeSHA256.CallOpts, arg0)
}

// LatestRoot is a free data retrieval call binding the contract method 0xd7b0fef1.
//
// Solidity: function latestRoot() view returns(bytes32)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Caller) LatestRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MerkleTreeSHA256.contract.Call(opts, &out, "latestRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestRoot is a free data retrieval call binding the contract method 0xd7b0fef1.
//
// Solidity: function latestRoot() view returns(bytes32)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Session) LatestRoot() ([32]byte, error) {
	return _MerkleTreeSHA256.Contract.LatestRoot(&_MerkleTreeSHA256.CallOpts)
}

// LatestRoot is a free data retrieval call binding the contract method 0xd7b0fef1.
//
// Solidity: function latestRoot() view returns(bytes32)
func (_MerkleTreeSHA256 *MerkleTreeSHA256CallerSession) LatestRoot() ([32]byte, error) {
	return _MerkleTreeSHA256.Contract.LatestRoot(&_MerkleTreeSHA256.CallOpts)
}

// LeafCount is a free data retrieval call binding the contract method 0x30e69fc3.
//
// Solidity: function leafCount() view returns(uint256)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Caller) LeafCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTreeSHA256.contract.Call(opts, &out, "leafCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LeafCount is a free data retrieval call binding the contract method 0x30e69fc3.
//
// Solidity: function leafCount() view returns(uint256)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Session) LeafCount() (*big.Int, error) {
	return _MerkleTreeSHA256.Contract.LeafCount(&_MerkleTreeSHA256.CallOpts)
}

// LeafCount is a free data retrieval call binding the contract method 0x30e69fc3.
//
// Solidity: function leafCount() view returns(uint256)
func (_MerkleTreeSHA256 *MerkleTreeSHA256CallerSession) LeafCount() (*big.Int, error) {
	return _MerkleTreeSHA256.Contract.LeafCount(&_MerkleTreeSHA256.CallOpts)
}

// TreeHeight is a free data retrieval call binding the contract method 0x01e3e915.
//
// Solidity: function treeHeight() view returns(uint256)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Caller) TreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTreeSHA256.contract.Call(opts, &out, "treeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreeHeight is a free data retrieval call binding the contract method 0x01e3e915.
//
// Solidity: function treeHeight() view returns(uint256)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Session) TreeHeight() (*big.Int, error) {
	return _MerkleTreeSHA256.Contract.TreeHeight(&_MerkleTreeSHA256.CallOpts)
}

// TreeHeight is a free data retrieval call binding the contract method 0x01e3e915.
//
// Solidity: function treeHeight() view returns(uint256)
func (_MerkleTreeSHA256 *MerkleTreeSHA256CallerSession) TreeHeight() (*big.Int, error) {
	return _MerkleTreeSHA256.Contract.TreeHeight(&_MerkleTreeSHA256.CallOpts)
}

// TreeWidth is a free data retrieval call binding the contract method 0x76c601b1.
//
// Solidity: function treeWidth() view returns(uint256)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Caller) TreeWidth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTreeSHA256.contract.Call(opts, &out, "treeWidth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreeWidth is a free data retrieval call binding the contract method 0x76c601b1.
//
// Solidity: function treeWidth() view returns(uint256)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Session) TreeWidth() (*big.Int, error) {
	return _MerkleTreeSHA256.Contract.TreeWidth(&_MerkleTreeSHA256.CallOpts)
}

// TreeWidth is a free data retrieval call binding the contract method 0x76c601b1.
//
// Solidity: function treeWidth() view returns(uint256)
func (_MerkleTreeSHA256 *MerkleTreeSHA256CallerSession) TreeWidth() (*big.Int, error) {
	return _MerkleTreeSHA256.Contract.TreeWidth(&_MerkleTreeSHA256.CallOpts)
}

// Zero is a free data retrieval call binding the contract method 0xbc1b392d.
//
// Solidity: function zero() view returns(bytes27)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Caller) Zero(opts *bind.CallOpts) ([27]byte, error) {
	var out []interface{}
	err := _MerkleTreeSHA256.contract.Call(opts, &out, "zero")

	if err != nil {
		return *new([27]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([27]byte)).(*[27]byte)

	return out0, err

}

// Zero is a free data retrieval call binding the contract method 0xbc1b392d.
//
// Solidity: function zero() view returns(bytes27)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Session) Zero() ([27]byte, error) {
	return _MerkleTreeSHA256.Contract.Zero(&_MerkleTreeSHA256.CallOpts)
}

// Zero is a free data retrieval call binding the contract method 0xbc1b392d.
//
// Solidity: function zero() view returns(bytes27)
func (_MerkleTreeSHA256 *MerkleTreeSHA256CallerSession) Zero() ([27]byte, error) {
	return _MerkleTreeSHA256.Contract.Zero(&_MerkleTreeSHA256.CallOpts)
}

// MerkleTreeSHA256NewLeafIterator is returned from FilterNewLeaf and is used to iterate over the raw logs and unpacked data for NewLeaf events raised by the MerkleTreeSHA256 contract.
type MerkleTreeSHA256NewLeafIterator struct {
	Event *MerkleTreeSHA256NewLeaf // Event containing the contract specifics and raw log

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
func (it *MerkleTreeSHA256NewLeafIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleTreeSHA256NewLeaf)
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
		it.Event = new(MerkleTreeSHA256NewLeaf)
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
func (it *MerkleTreeSHA256NewLeafIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleTreeSHA256NewLeafIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleTreeSHA256NewLeaf represents a NewLeaf event raised by the MerkleTreeSHA256 contract.
type MerkleTreeSHA256NewLeaf struct {
	LeafIndex *big.Int
	LeafValue [32]byte
	Root      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewLeaf is a free log retrieval operation binding the contract event 0x6a82ba2aa1d2c039c41e6e2b5a5a1090d09906f060d32af9c1ac0beff7af75c0.
//
// Solidity: event NewLeaf(uint256 leafIndex, bytes32 leafValue, bytes32 root)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Filterer) FilterNewLeaf(opts *bind.FilterOpts) (*MerkleTreeSHA256NewLeafIterator, error) {

	logs, sub, err := _MerkleTreeSHA256.contract.FilterLogs(opts, "NewLeaf")
	if err != nil {
		return nil, err
	}
	return &MerkleTreeSHA256NewLeafIterator{contract: _MerkleTreeSHA256.contract, event: "NewLeaf", logs: logs, sub: sub}, nil
}

// WatchNewLeaf is a free log subscription operation binding the contract event 0x6a82ba2aa1d2c039c41e6e2b5a5a1090d09906f060d32af9c1ac0beff7af75c0.
//
// Solidity: event NewLeaf(uint256 leafIndex, bytes32 leafValue, bytes32 root)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Filterer) WatchNewLeaf(opts *bind.WatchOpts, sink chan<- *MerkleTreeSHA256NewLeaf) (event.Subscription, error) {

	logs, sub, err := _MerkleTreeSHA256.contract.WatchLogs(opts, "NewLeaf")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleTreeSHA256NewLeaf)
				if err := _MerkleTreeSHA256.contract.UnpackLog(event, "NewLeaf", log); err != nil {
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

// ParseNewLeaf is a log parse operation binding the contract event 0x6a82ba2aa1d2c039c41e6e2b5a5a1090d09906f060d32af9c1ac0beff7af75c0.
//
// Solidity: event NewLeaf(uint256 leafIndex, bytes32 leafValue, bytes32 root)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Filterer) ParseNewLeaf(log types.Log) (*MerkleTreeSHA256NewLeaf, error) {
	event := new(MerkleTreeSHA256NewLeaf)
	if err := _MerkleTreeSHA256.contract.UnpackLog(event, "NewLeaf", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleTreeSHA256NewLeavesIterator is returned from FilterNewLeaves and is used to iterate over the raw logs and unpacked data for NewLeaves events raised by the MerkleTreeSHA256 contract.
type MerkleTreeSHA256NewLeavesIterator struct {
	Event *MerkleTreeSHA256NewLeaves // Event containing the contract specifics and raw log

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
func (it *MerkleTreeSHA256NewLeavesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleTreeSHA256NewLeaves)
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
		it.Event = new(MerkleTreeSHA256NewLeaves)
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
func (it *MerkleTreeSHA256NewLeavesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleTreeSHA256NewLeavesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleTreeSHA256NewLeaves represents a NewLeaves event raised by the MerkleTreeSHA256 contract.
type MerkleTreeSHA256NewLeaves struct {
	MinLeafIndex *big.Int
	LeafValues   [][32]byte
	Root         [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewLeaves is a free log retrieval operation binding the contract event 0x8ec50f97970775682a68d3c6f9caedf60fd82448ea40706b8b65d6c03648b922.
//
// Solidity: event NewLeaves(uint256 minLeafIndex, bytes32[] leafValues, bytes32 root)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Filterer) FilterNewLeaves(opts *bind.FilterOpts) (*MerkleTreeSHA256NewLeavesIterator, error) {

	logs, sub, err := _MerkleTreeSHA256.contract.FilterLogs(opts, "NewLeaves")
	if err != nil {
		return nil, err
	}
	return &MerkleTreeSHA256NewLeavesIterator{contract: _MerkleTreeSHA256.contract, event: "NewLeaves", logs: logs, sub: sub}, nil
}

// WatchNewLeaves is a free log subscription operation binding the contract event 0x8ec50f97970775682a68d3c6f9caedf60fd82448ea40706b8b65d6c03648b922.
//
// Solidity: event NewLeaves(uint256 minLeafIndex, bytes32[] leafValues, bytes32 root)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Filterer) WatchNewLeaves(opts *bind.WatchOpts, sink chan<- *MerkleTreeSHA256NewLeaves) (event.Subscription, error) {

	logs, sub, err := _MerkleTreeSHA256.contract.WatchLogs(opts, "NewLeaves")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleTreeSHA256NewLeaves)
				if err := _MerkleTreeSHA256.contract.UnpackLog(event, "NewLeaves", log); err != nil {
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

// ParseNewLeaves is a log parse operation binding the contract event 0x8ec50f97970775682a68d3c6f9caedf60fd82448ea40706b8b65d6c03648b922.
//
// Solidity: event NewLeaves(uint256 minLeafIndex, bytes32[] leafValues, bytes32 root)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Filterer) ParseNewLeaves(log types.Log) (*MerkleTreeSHA256NewLeaves, error) {
	event := new(MerkleTreeSHA256NewLeaves)
	if err := _MerkleTreeSHA256.contract.UnpackLog(event, "NewLeaves", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleTreeSHA256OutputIterator is returned from FilterOutput and is used to iterate over the raw logs and unpacked data for Output events raised by the MerkleTreeSHA256 contract.
type MerkleTreeSHA256OutputIterator struct {
	Event *MerkleTreeSHA256Output // Event containing the contract specifics and raw log

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
func (it *MerkleTreeSHA256OutputIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleTreeSHA256Output)
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
		it.Event = new(MerkleTreeSHA256Output)
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
func (it *MerkleTreeSHA256OutputIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleTreeSHA256OutputIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleTreeSHA256Output represents a Output event raised by the MerkleTreeSHA256 contract.
type MerkleTreeSHA256Output struct {
	LeftInput  [27]byte
	RightInput [27]byte
	Output     [32]byte
	NodeIndex  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOutput is a free log retrieval operation binding the contract event 0xc84edbd4893d9f5e7f1a0e9cef31bc608e285b8683036b7b0c62346c9785b60d.
//
// Solidity: event Output(bytes27 leftInput, bytes27 rightInput, bytes32 output, uint256 nodeIndex)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Filterer) FilterOutput(opts *bind.FilterOpts) (*MerkleTreeSHA256OutputIterator, error) {

	logs, sub, err := _MerkleTreeSHA256.contract.FilterLogs(opts, "Output")
	if err != nil {
		return nil, err
	}
	return &MerkleTreeSHA256OutputIterator{contract: _MerkleTreeSHA256.contract, event: "Output", logs: logs, sub: sub}, nil
}

// WatchOutput is a free log subscription operation binding the contract event 0xc84edbd4893d9f5e7f1a0e9cef31bc608e285b8683036b7b0c62346c9785b60d.
//
// Solidity: event Output(bytes27 leftInput, bytes27 rightInput, bytes32 output, uint256 nodeIndex)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Filterer) WatchOutput(opts *bind.WatchOpts, sink chan<- *MerkleTreeSHA256Output) (event.Subscription, error) {

	logs, sub, err := _MerkleTreeSHA256.contract.WatchLogs(opts, "Output")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleTreeSHA256Output)
				if err := _MerkleTreeSHA256.contract.UnpackLog(event, "Output", log); err != nil {
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

// ParseOutput is a log parse operation binding the contract event 0xc84edbd4893d9f5e7f1a0e9cef31bc608e285b8683036b7b0c62346c9785b60d.
//
// Solidity: event Output(bytes27 leftInput, bytes27 rightInput, bytes32 output, uint256 nodeIndex)
func (_MerkleTreeSHA256 *MerkleTreeSHA256Filterer) ParseOutput(log types.Log) (*MerkleTreeSHA256Output, error) {
	event := new(MerkleTreeSHA256Output)
	if err := _MerkleTreeSHA256.contract.UnpackLog(event, "Output", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShieldMetaData contains all meta data concerning the Shield contract.
var ShieldMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_FPLverifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ACKverifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_treeHeight\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"leafValue\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"NewLeaf\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minLeafIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"leafValues\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"NewLeaves\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes27\",\"name\":\"leftInput\",\"type\":\"bytes27\"},{\"indexed\":false,\"internalType\":\"bytes27\",\"name\":\"rightInput\",\"type\":\"bytes27\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"output\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodeIndex\",\"type\":\"uint256\"}],\"name\":\"Output\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"frontier\",\"outputs\":[{\"internalType\":\"bytes27\",\"name\":\"\",\"type\":\"bytes27\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getACKVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFPLVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leafCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeWidth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[12]\",\"name\":\"input\",\"type\":\"uint256[12]\"},{\"internalType\":\"bytes32\",\"name\":\"_newCommitment\",\"type\":\"bytes32\"}],\"name\":\"verifyAndPushACK\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[6]\",\"name\":\"input\",\"type\":\"uint256[6]\"},{\"internalType\":\"bytes32\",\"name\":\"_newCommitment\",\"type\":\"bytes32\"}],\"name\":\"verifyAndPushFPL\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zero\",\"outputs\":[{\"internalType\":\"bytes27\",\"name\":\"\",\"type\":\"bytes27\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6e0c3fee": "frontier(uint256)",
		"76a7d786": "getACKVerifier()",
		"9c233ca1": "getFPLVerifier()",
		"d7b0fef1": "latestRoot()",
		"30e69fc3": "leafCount()",
		"01e3e915": "treeHeight()",
		"76c601b1": "treeWidth()",
		"741569fc": "verifyAndPushACK(uint256[2],uint256[2][2],uint256[2],uint256[12],bytes32)",
		"9f0efffd": "verifyAndPushFPL(uint256[2],uint256[2][2],uint256[2],uint256[6],bytes32)",
		"bc1b392d": "zero()",
	},
	Bin: "0x6080604052600480546001600160d81b03191690553480156200002157600080fd5b5060405162000c4938038062000c498339810160408190526200004491620000ab565b6000819055806200005781600262000201565b6001555050602680546001600160a01b039384166001600160a01b0319918216179091556027805492909316911617905562000216565b80516001600160a01b0381168114620000a657600080fd5b919050565b600080600060608486031215620000c157600080fd5b620000cc846200008e565b9250620000dc602085016200008e565b9150604084015190509250925092565b634e487b7160e01b600052601160045260246000fd5b600181815b8085111562000143578160001904821115620001275762000127620000ec565b808516156200013557918102915b93841c939080029062000107565b509250929050565b6000826200015c57506001620001fb565b816200016b57506000620001fb565b81600181146200018457600281146200018f57620001af565b6001915050620001fb565b60ff841115620001a357620001a3620000ec565b50506001821b620001fb565b5060208310610133831016604e8410600b8410161715620001d4575081810a620001fb565b620001e0838362000102565b8060001904821115620001f757620001f7620000ec565b0290505b92915050565b60006200020f83836200014b565b9392505050565b610a2380620002266000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806376c601b11161006657806376c601b11461013a5780639c233ca1146101435780639f0efffd14610154578063bc1b392d14610167578063d7b0fef11461017457600080fd5b806301e3e915146100a357806330e69fc3146100bf5780636e0c3fee146100c8578063741569fc146100f257806376a7d78614610115575b600080fd5b6100ac60005481565b6040519081526020015b60405180910390f35b6100ac60035481565b6100db6100d6366004610598565b61017d565b60405164ffffffffff1990911681526020016100b6565b6101056101003660046106d6565b610197565b60405190151581526020016100b6565b6027546001600160a01b03165b6040516001600160a01b0390911681526020016100b6565b6100ac60015481565b6026546001600160a01b0316610122565b610105610162366004610775565b61029f565b6004546100db9060281b81565b6100ac60025481565b6005816021811061018d57600080fd5b015460281b905081565b60275460405163465d607d60e11b815260009182916001600160a01b0390911690638cbac0fa906101d2908a908a908a908a90600401610864565b6020604051808303816000875af11580156101f1573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061021591906108c1565b9050806102885760405162461bcd60e51b815260206004820152603660248201527f5468652070726f6f66206661696c656420766572696669636174696f6e20696e604482015275081d1a19481d995c9a599a595c8818dbdb9d1c9858dd60521b60648201526084015b60405180910390fd5b610291836102da565b506001979650505050505050565b60265460405163f398789b60e01b815260009182916001600160a01b039091169063f398789b906101d2908a908a908a908a906004016108ea565b60006003546001541161033b5760405162461bcd60e51b815260206004820152602360248201527f5468657265206973206e6f207370616365206c65667420696e2074686520747260448201526232b29760e91b606482015260840161027f565b6000610348600354610505565b905060006001805460035461035d9190610951565b6103679190610969565b9050602884901b60008061037961057a565b6000805b600054811015610494578781036103bc5785600589602181106103a2576103a2610980565b0180546001600160d81b03191660289290921c9190911790555b6103c76002886109ac565b60000361043957600581602181106103e1576103e1610980565b015460405160289190911b808252601b820188905295508694506020846036836002600019fa92508290508061041357fe5b50825160281b95506002610428600189610969565b61043291906109c0565b9650610482565b60045460405187815260289190911b601b820181905287965094506020846036836002600019fa92508290508061046c57fe5b50825160281b955061047f6002886109c0565b96505b8061048c816109d4565b91505061037d565b508151600281905560035460408051918252602082018c90528101919091527f6a82ba2aa1d2c039c41e6e2b5a5a1090d09906f060d32af9c1ac0beff7af75c09060600160405180910390a1600380549060006104f0836109d4565b90915550506002549998505050505050505050565b60006105126002836109ac565b600103610575576001600260045b83600003610571578082610535876001610951565b61053f9190610969565b61054991906109ac565b60000361055857829350610520565b9050600181901b82610569816109d4565b935050610520565b5050505b919050565b60405180602001604052806001906020820280368337509192915050565b6000602082840312156105aa57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156105ea576105ea6105b1565b60405290565b604051610180810167ffffffffffffffff811182821017156105ea576105ea6105b1565b60405160c0810167ffffffffffffffff811182821017156105ea576105ea6105b1565b600082601f83011261064857600080fd5b6106506105c7565b80604084018581111561066257600080fd5b845b8181101561067c578035845260209384019301610664565b509095945050505050565b600082601f83011261069857600080fd5b6106a06105c7565b8060808401858111156106b257600080fd5b845b8181101561067c576106c68782610637565b84526020909301926040016106b4565b60008060008060006102a086880312156106ef57600080fd5b6106f98787610637565b94506107088760408801610687565b93506107178760c08801610637565b92508661011f87011261072957600080fd5b6107316105f0565b8061028088018981111561074457600080fd5b61010089015b8181101561076257803584526020938401930161074a565b5096999598509396509294359392505050565b60008060008060006101e0868803121561078e57600080fd5b6107988787610637565b94506107a78760408801610687565b93506107b68760c08801610637565b92508661011f8701126107c857600080fd5b6107d0610614565b806101c088018981111561074457600080fd5b8060005b60028110156108065781518452602093840193909101906001016107e7565b50505050565b806000805b6002808210610820575061085d565b835186845b83811015610843578251825260209283019290910190600101610825565b505050604095909501945060209290920191600101610811565b5050505050565b610280810161087382876107e3565b610880604083018661080c565b61088d60c08301856107e3565b61010082018360005b600c8110156108b5578151835260209283019290910190600101610896565b50505095945050505050565b6000602082840312156108d357600080fd5b815180151581146108e357600080fd5b9392505050565b6101c081016108f982876107e3565b610906604083018661080c565b61091360c08301856107e3565b61010082018360005b60068110156108b557815183526020928301929091019060010161091c565b634e487b7160e01b600052601160045260246000fd5b600082198211156109645761096461093b565b500190565b60008282101561097b5761097b61093b565b500390565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601260045260246000fd5b6000826109bb576109bb610996565b500690565b6000826109cf576109cf610996565b500490565b6000600182016109e6576109e661093b565b506001019056fea264697066735822122019fbcf75132c986754c58881aaafb82ffc29bc03e9084225bce71c3473b4578d64736f6c634300080d0033",
}

// ShieldABI is the input ABI used to generate the binding from.
// Deprecated: Use ShieldMetaData.ABI instead.
var ShieldABI = ShieldMetaData.ABI

// Deprecated: Use ShieldMetaData.Sigs instead.
// ShieldFuncSigs maps the 4-byte function signature to its string representation.
var ShieldFuncSigs = ShieldMetaData.Sigs

// ShieldBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ShieldMetaData.Bin instead.
var ShieldBin = ShieldMetaData.Bin

// DeployShield deploys a new Ethereum contract, binding an instance of Shield to it.
func DeployShield(auth *bind.TransactOpts, backend bind.ContractBackend, _FPLverifier common.Address, _ACKverifier common.Address, _treeHeight *big.Int) (common.Address, *types.Transaction, *Shield, error) {
	parsed, err := ShieldMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ShieldBin), backend, _FPLverifier, _ACKverifier, _treeHeight)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Shield{ShieldCaller: ShieldCaller{contract: contract}, ShieldTransactor: ShieldTransactor{contract: contract}, ShieldFilterer: ShieldFilterer{contract: contract}}, nil
}

// Shield is an auto generated Go binding around an Ethereum contract.
type Shield struct {
	ShieldCaller     // Read-only binding to the contract
	ShieldTransactor // Write-only binding to the contract
	ShieldFilterer   // Log filterer for contract events
}

// ShieldCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShieldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShieldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShieldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShieldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShieldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShieldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShieldSession struct {
	Contract     *Shield           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShieldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShieldCallerSession struct {
	Contract *ShieldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ShieldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShieldTransactorSession struct {
	Contract     *ShieldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShieldRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShieldRaw struct {
	Contract *Shield // Generic contract binding to access the raw methods on
}

// ShieldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShieldCallerRaw struct {
	Contract *ShieldCaller // Generic read-only contract binding to access the raw methods on
}

// ShieldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShieldTransactorRaw struct {
	Contract *ShieldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShield creates a new instance of Shield, bound to a specific deployed contract.
func NewShield(address common.Address, backend bind.ContractBackend) (*Shield, error) {
	contract, err := bindShield(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Shield{ShieldCaller: ShieldCaller{contract: contract}, ShieldTransactor: ShieldTransactor{contract: contract}, ShieldFilterer: ShieldFilterer{contract: contract}}, nil
}

// NewShieldCaller creates a new read-only instance of Shield, bound to a specific deployed contract.
func NewShieldCaller(address common.Address, caller bind.ContractCaller) (*ShieldCaller, error) {
	contract, err := bindShield(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShieldCaller{contract: contract}, nil
}

// NewShieldTransactor creates a new write-only instance of Shield, bound to a specific deployed contract.
func NewShieldTransactor(address common.Address, transactor bind.ContractTransactor) (*ShieldTransactor, error) {
	contract, err := bindShield(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShieldTransactor{contract: contract}, nil
}

// NewShieldFilterer creates a new log filterer instance of Shield, bound to a specific deployed contract.
func NewShieldFilterer(address common.Address, filterer bind.ContractFilterer) (*ShieldFilterer, error) {
	contract, err := bindShield(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShieldFilterer{contract: contract}, nil
}

// bindShield binds a generic wrapper to an already deployed contract.
func bindShield(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShieldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shield *ShieldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shield.Contract.ShieldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shield *ShieldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shield.Contract.ShieldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shield *ShieldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shield.Contract.ShieldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shield *ShieldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shield.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shield *ShieldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shield.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shield *ShieldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shield.Contract.contract.Transact(opts, method, params...)
}

// Frontier is a free data retrieval call binding the contract method 0x6e0c3fee.
//
// Solidity: function frontier(uint256 ) view returns(bytes27)
func (_Shield *ShieldCaller) Frontier(opts *bind.CallOpts, arg0 *big.Int) ([27]byte, error) {
	var out []interface{}
	err := _Shield.contract.Call(opts, &out, "frontier", arg0)

	if err != nil {
		return *new([27]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([27]byte)).(*[27]byte)

	return out0, err

}

// Frontier is a free data retrieval call binding the contract method 0x6e0c3fee.
//
// Solidity: function frontier(uint256 ) view returns(bytes27)
func (_Shield *ShieldSession) Frontier(arg0 *big.Int) ([27]byte, error) {
	return _Shield.Contract.Frontier(&_Shield.CallOpts, arg0)
}

// Frontier is a free data retrieval call binding the contract method 0x6e0c3fee.
//
// Solidity: function frontier(uint256 ) view returns(bytes27)
func (_Shield *ShieldCallerSession) Frontier(arg0 *big.Int) ([27]byte, error) {
	return _Shield.Contract.Frontier(&_Shield.CallOpts, arg0)
}

// GetACKVerifier is a free data retrieval call binding the contract method 0x76a7d786.
//
// Solidity: function getACKVerifier() view returns(address)
func (_Shield *ShieldCaller) GetACKVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Shield.contract.Call(opts, &out, "getACKVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetACKVerifier is a free data retrieval call binding the contract method 0x76a7d786.
//
// Solidity: function getACKVerifier() view returns(address)
func (_Shield *ShieldSession) GetACKVerifier() (common.Address, error) {
	return _Shield.Contract.GetACKVerifier(&_Shield.CallOpts)
}

// GetACKVerifier is a free data retrieval call binding the contract method 0x76a7d786.
//
// Solidity: function getACKVerifier() view returns(address)
func (_Shield *ShieldCallerSession) GetACKVerifier() (common.Address, error) {
	return _Shield.Contract.GetACKVerifier(&_Shield.CallOpts)
}

// GetFPLVerifier is a free data retrieval call binding the contract method 0x9c233ca1.
//
// Solidity: function getFPLVerifier() view returns(address)
func (_Shield *ShieldCaller) GetFPLVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Shield.contract.Call(opts, &out, "getFPLVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFPLVerifier is a free data retrieval call binding the contract method 0x9c233ca1.
//
// Solidity: function getFPLVerifier() view returns(address)
func (_Shield *ShieldSession) GetFPLVerifier() (common.Address, error) {
	return _Shield.Contract.GetFPLVerifier(&_Shield.CallOpts)
}

// GetFPLVerifier is a free data retrieval call binding the contract method 0x9c233ca1.
//
// Solidity: function getFPLVerifier() view returns(address)
func (_Shield *ShieldCallerSession) GetFPLVerifier() (common.Address, error) {
	return _Shield.Contract.GetFPLVerifier(&_Shield.CallOpts)
}

// LatestRoot is a free data retrieval call binding the contract method 0xd7b0fef1.
//
// Solidity: function latestRoot() view returns(bytes32)
func (_Shield *ShieldCaller) LatestRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Shield.contract.Call(opts, &out, "latestRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestRoot is a free data retrieval call binding the contract method 0xd7b0fef1.
//
// Solidity: function latestRoot() view returns(bytes32)
func (_Shield *ShieldSession) LatestRoot() ([32]byte, error) {
	return _Shield.Contract.LatestRoot(&_Shield.CallOpts)
}

// LatestRoot is a free data retrieval call binding the contract method 0xd7b0fef1.
//
// Solidity: function latestRoot() view returns(bytes32)
func (_Shield *ShieldCallerSession) LatestRoot() ([32]byte, error) {
	return _Shield.Contract.LatestRoot(&_Shield.CallOpts)
}

// LeafCount is a free data retrieval call binding the contract method 0x30e69fc3.
//
// Solidity: function leafCount() view returns(uint256)
func (_Shield *ShieldCaller) LeafCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Shield.contract.Call(opts, &out, "leafCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LeafCount is a free data retrieval call binding the contract method 0x30e69fc3.
//
// Solidity: function leafCount() view returns(uint256)
func (_Shield *ShieldSession) LeafCount() (*big.Int, error) {
	return _Shield.Contract.LeafCount(&_Shield.CallOpts)
}

// LeafCount is a free data retrieval call binding the contract method 0x30e69fc3.
//
// Solidity: function leafCount() view returns(uint256)
func (_Shield *ShieldCallerSession) LeafCount() (*big.Int, error) {
	return _Shield.Contract.LeafCount(&_Shield.CallOpts)
}

// TreeHeight is a free data retrieval call binding the contract method 0x01e3e915.
//
// Solidity: function treeHeight() view returns(uint256)
func (_Shield *ShieldCaller) TreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Shield.contract.Call(opts, &out, "treeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreeHeight is a free data retrieval call binding the contract method 0x01e3e915.
//
// Solidity: function treeHeight() view returns(uint256)
func (_Shield *ShieldSession) TreeHeight() (*big.Int, error) {
	return _Shield.Contract.TreeHeight(&_Shield.CallOpts)
}

// TreeHeight is a free data retrieval call binding the contract method 0x01e3e915.
//
// Solidity: function treeHeight() view returns(uint256)
func (_Shield *ShieldCallerSession) TreeHeight() (*big.Int, error) {
	return _Shield.Contract.TreeHeight(&_Shield.CallOpts)
}

// TreeWidth is a free data retrieval call binding the contract method 0x76c601b1.
//
// Solidity: function treeWidth() view returns(uint256)
func (_Shield *ShieldCaller) TreeWidth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Shield.contract.Call(opts, &out, "treeWidth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreeWidth is a free data retrieval call binding the contract method 0x76c601b1.
//
// Solidity: function treeWidth() view returns(uint256)
func (_Shield *ShieldSession) TreeWidth() (*big.Int, error) {
	return _Shield.Contract.TreeWidth(&_Shield.CallOpts)
}

// TreeWidth is a free data retrieval call binding the contract method 0x76c601b1.
//
// Solidity: function treeWidth() view returns(uint256)
func (_Shield *ShieldCallerSession) TreeWidth() (*big.Int, error) {
	return _Shield.Contract.TreeWidth(&_Shield.CallOpts)
}

// Zero is a free data retrieval call binding the contract method 0xbc1b392d.
//
// Solidity: function zero() view returns(bytes27)
func (_Shield *ShieldCaller) Zero(opts *bind.CallOpts) ([27]byte, error) {
	var out []interface{}
	err := _Shield.contract.Call(opts, &out, "zero")

	if err != nil {
		return *new([27]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([27]byte)).(*[27]byte)

	return out0, err

}

// Zero is a free data retrieval call binding the contract method 0xbc1b392d.
//
// Solidity: function zero() view returns(bytes27)
func (_Shield *ShieldSession) Zero() ([27]byte, error) {
	return _Shield.Contract.Zero(&_Shield.CallOpts)
}

// Zero is a free data retrieval call binding the contract method 0xbc1b392d.
//
// Solidity: function zero() view returns(bytes27)
func (_Shield *ShieldCallerSession) Zero() ([27]byte, error) {
	return _Shield.Contract.Zero(&_Shield.CallOpts)
}

// VerifyAndPushACK is a paid mutator transaction binding the contract method 0x741569fc.
//
// Solidity: function verifyAndPushACK(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[12] input, bytes32 _newCommitment) returns(bool)
func (_Shield *ShieldTransactor) VerifyAndPushACK(opts *bind.TransactOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [12]*big.Int, _newCommitment [32]byte) (*types.Transaction, error) {
	return _Shield.contract.Transact(opts, "verifyAndPushACK", a, b, c, input, _newCommitment)
}

// VerifyAndPushACK is a paid mutator transaction binding the contract method 0x741569fc.
//
// Solidity: function verifyAndPushACK(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[12] input, bytes32 _newCommitment) returns(bool)
func (_Shield *ShieldSession) VerifyAndPushACK(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [12]*big.Int, _newCommitment [32]byte) (*types.Transaction, error) {
	return _Shield.Contract.VerifyAndPushACK(&_Shield.TransactOpts, a, b, c, input, _newCommitment)
}

// VerifyAndPushACK is a paid mutator transaction binding the contract method 0x741569fc.
//
// Solidity: function verifyAndPushACK(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[12] input, bytes32 _newCommitment) returns(bool)
func (_Shield *ShieldTransactorSession) VerifyAndPushACK(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [12]*big.Int, _newCommitment [32]byte) (*types.Transaction, error) {
	return _Shield.Contract.VerifyAndPushACK(&_Shield.TransactOpts, a, b, c, input, _newCommitment)
}

// VerifyAndPushFPL is a paid mutator transaction binding the contract method 0x9f0efffd.
//
// Solidity: function verifyAndPushFPL(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[6] input, bytes32 _newCommitment) returns(bool)
func (_Shield *ShieldTransactor) VerifyAndPushFPL(opts *bind.TransactOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [6]*big.Int, _newCommitment [32]byte) (*types.Transaction, error) {
	return _Shield.contract.Transact(opts, "verifyAndPushFPL", a, b, c, input, _newCommitment)
}

// VerifyAndPushFPL is a paid mutator transaction binding the contract method 0x9f0efffd.
//
// Solidity: function verifyAndPushFPL(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[6] input, bytes32 _newCommitment) returns(bool)
func (_Shield *ShieldSession) VerifyAndPushFPL(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [6]*big.Int, _newCommitment [32]byte) (*types.Transaction, error) {
	return _Shield.Contract.VerifyAndPushFPL(&_Shield.TransactOpts, a, b, c, input, _newCommitment)
}

// VerifyAndPushFPL is a paid mutator transaction binding the contract method 0x9f0efffd.
//
// Solidity: function verifyAndPushFPL(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[6] input, bytes32 _newCommitment) returns(bool)
func (_Shield *ShieldTransactorSession) VerifyAndPushFPL(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [6]*big.Int, _newCommitment [32]byte) (*types.Transaction, error) {
	return _Shield.Contract.VerifyAndPushFPL(&_Shield.TransactOpts, a, b, c, input, _newCommitment)
}

// ShieldNewLeafIterator is returned from FilterNewLeaf and is used to iterate over the raw logs and unpacked data for NewLeaf events raised by the Shield contract.
type ShieldNewLeafIterator struct {
	Event *ShieldNewLeaf // Event containing the contract specifics and raw log

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
func (it *ShieldNewLeafIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShieldNewLeaf)
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
		it.Event = new(ShieldNewLeaf)
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
func (it *ShieldNewLeafIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShieldNewLeafIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShieldNewLeaf represents a NewLeaf event raised by the Shield contract.
type ShieldNewLeaf struct {
	LeafIndex *big.Int
	LeafValue [32]byte
	Root      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewLeaf is a free log retrieval operation binding the contract event 0x6a82ba2aa1d2c039c41e6e2b5a5a1090d09906f060d32af9c1ac0beff7af75c0.
//
// Solidity: event NewLeaf(uint256 leafIndex, bytes32 leafValue, bytes32 root)
func (_Shield *ShieldFilterer) FilterNewLeaf(opts *bind.FilterOpts) (*ShieldNewLeafIterator, error) {

	logs, sub, err := _Shield.contract.FilterLogs(opts, "NewLeaf")
	if err != nil {
		return nil, err
	}
	return &ShieldNewLeafIterator{contract: _Shield.contract, event: "NewLeaf", logs: logs, sub: sub}, nil
}

// WatchNewLeaf is a free log subscription operation binding the contract event 0x6a82ba2aa1d2c039c41e6e2b5a5a1090d09906f060d32af9c1ac0beff7af75c0.
//
// Solidity: event NewLeaf(uint256 leafIndex, bytes32 leafValue, bytes32 root)
func (_Shield *ShieldFilterer) WatchNewLeaf(opts *bind.WatchOpts, sink chan<- *ShieldNewLeaf) (event.Subscription, error) {

	logs, sub, err := _Shield.contract.WatchLogs(opts, "NewLeaf")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShieldNewLeaf)
				if err := _Shield.contract.UnpackLog(event, "NewLeaf", log); err != nil {
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

// ParseNewLeaf is a log parse operation binding the contract event 0x6a82ba2aa1d2c039c41e6e2b5a5a1090d09906f060d32af9c1ac0beff7af75c0.
//
// Solidity: event NewLeaf(uint256 leafIndex, bytes32 leafValue, bytes32 root)
func (_Shield *ShieldFilterer) ParseNewLeaf(log types.Log) (*ShieldNewLeaf, error) {
	event := new(ShieldNewLeaf)
	if err := _Shield.contract.UnpackLog(event, "NewLeaf", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShieldNewLeavesIterator is returned from FilterNewLeaves and is used to iterate over the raw logs and unpacked data for NewLeaves events raised by the Shield contract.
type ShieldNewLeavesIterator struct {
	Event *ShieldNewLeaves // Event containing the contract specifics and raw log

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
func (it *ShieldNewLeavesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShieldNewLeaves)
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
		it.Event = new(ShieldNewLeaves)
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
func (it *ShieldNewLeavesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShieldNewLeavesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShieldNewLeaves represents a NewLeaves event raised by the Shield contract.
type ShieldNewLeaves struct {
	MinLeafIndex *big.Int
	LeafValues   [][32]byte
	Root         [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewLeaves is a free log retrieval operation binding the contract event 0x8ec50f97970775682a68d3c6f9caedf60fd82448ea40706b8b65d6c03648b922.
//
// Solidity: event NewLeaves(uint256 minLeafIndex, bytes32[] leafValues, bytes32 root)
func (_Shield *ShieldFilterer) FilterNewLeaves(opts *bind.FilterOpts) (*ShieldNewLeavesIterator, error) {

	logs, sub, err := _Shield.contract.FilterLogs(opts, "NewLeaves")
	if err != nil {
		return nil, err
	}
	return &ShieldNewLeavesIterator{contract: _Shield.contract, event: "NewLeaves", logs: logs, sub: sub}, nil
}

// WatchNewLeaves is a free log subscription operation binding the contract event 0x8ec50f97970775682a68d3c6f9caedf60fd82448ea40706b8b65d6c03648b922.
//
// Solidity: event NewLeaves(uint256 minLeafIndex, bytes32[] leafValues, bytes32 root)
func (_Shield *ShieldFilterer) WatchNewLeaves(opts *bind.WatchOpts, sink chan<- *ShieldNewLeaves) (event.Subscription, error) {

	logs, sub, err := _Shield.contract.WatchLogs(opts, "NewLeaves")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShieldNewLeaves)
				if err := _Shield.contract.UnpackLog(event, "NewLeaves", log); err != nil {
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

// ParseNewLeaves is a log parse operation binding the contract event 0x8ec50f97970775682a68d3c6f9caedf60fd82448ea40706b8b65d6c03648b922.
//
// Solidity: event NewLeaves(uint256 minLeafIndex, bytes32[] leafValues, bytes32 root)
func (_Shield *ShieldFilterer) ParseNewLeaves(log types.Log) (*ShieldNewLeaves, error) {
	event := new(ShieldNewLeaves)
	if err := _Shield.contract.UnpackLog(event, "NewLeaves", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShieldOutputIterator is returned from FilterOutput and is used to iterate over the raw logs and unpacked data for Output events raised by the Shield contract.
type ShieldOutputIterator struct {
	Event *ShieldOutput // Event containing the contract specifics and raw log

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
func (it *ShieldOutputIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShieldOutput)
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
		it.Event = new(ShieldOutput)
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
func (it *ShieldOutputIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShieldOutputIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShieldOutput represents a Output event raised by the Shield contract.
type ShieldOutput struct {
	LeftInput  [27]byte
	RightInput [27]byte
	Output     [32]byte
	NodeIndex  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOutput is a free log retrieval operation binding the contract event 0xc84edbd4893d9f5e7f1a0e9cef31bc608e285b8683036b7b0c62346c9785b60d.
//
// Solidity: event Output(bytes27 leftInput, bytes27 rightInput, bytes32 output, uint256 nodeIndex)
func (_Shield *ShieldFilterer) FilterOutput(opts *bind.FilterOpts) (*ShieldOutputIterator, error) {

	logs, sub, err := _Shield.contract.FilterLogs(opts, "Output")
	if err != nil {
		return nil, err
	}
	return &ShieldOutputIterator{contract: _Shield.contract, event: "Output", logs: logs, sub: sub}, nil
}

// WatchOutput is a free log subscription operation binding the contract event 0xc84edbd4893d9f5e7f1a0e9cef31bc608e285b8683036b7b0c62346c9785b60d.
//
// Solidity: event Output(bytes27 leftInput, bytes27 rightInput, bytes32 output, uint256 nodeIndex)
func (_Shield *ShieldFilterer) WatchOutput(opts *bind.WatchOpts, sink chan<- *ShieldOutput) (event.Subscription, error) {

	logs, sub, err := _Shield.contract.WatchLogs(opts, "Output")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShieldOutput)
				if err := _Shield.contract.UnpackLog(event, "Output", log); err != nil {
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

// ParseOutput is a log parse operation binding the contract event 0xc84edbd4893d9f5e7f1a0e9cef31bc608e285b8683036b7b0c62346c9785b60d.
//
// Solidity: event Output(bytes27 leftInput, bytes27 rightInput, bytes32 output, uint256 nodeIndex)
func (_Shield *ShieldFilterer) ParseOutput(log types.Log) (*ShieldOutput, error) {
	event := new(ShieldOutput)
	if err := _Shield.contract.UnpackLog(event, "Output", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
