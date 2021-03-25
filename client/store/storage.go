// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IPFSStorageABI is the input ABI used to generate the binding from.
const IPFSStorageABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getEntry\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"entry\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"setEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IPFSStorageFuncSigs maps the 4-byte function signature to its string representation.
var IPFSStorageFuncSigs = map[string]string{
	"7db6a4e4": "getEntry(address)",
	"4b369820": "setEntry(string)",
}

// IPFSStorageBin is the compiled bytecode used for deploying new contracts.
var IPFSStorageBin = "0x608060405234801561001057600080fd5b50610394806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634b3698201461003b5780637db6a4e414610050575b600080fd5b61004e610049366004610210565b610079565b005b61006361005e3660046101e2565b61009d565b60405161007091906102ba565b60405180910390f35b336000908152600160209081526040909120825161009992840190610149565b5050565b6001600160a01b03811660009081526001602052604090208054606091906100c49061030d565b80601f01602080910402602001604051908101604052809291908181526020018280546100f09061030d565b801561013d5780601f106101125761010080835404028352916020019161013d565b820191906000526020600020905b81548152906001019060200180831161012057829003601f168201915b50505050509050919050565b8280546101559061030d565b90600052602060002090601f01602090048101928261017757600085556101bd565b82601f1061019057805160ff19168380011785556101bd565b828001600101855582156101bd579182015b828111156101bd5782518255916020019190600101906101a2565b506101c99291506101cd565b5090565b5b808211156101c957600081556001016101ce565b6000602082840312156101f3578081fd5b81356001600160a01b0381168114610209578182fd5b9392505050565b600060208284031215610221578081fd5b813567ffffffffffffffff80821115610238578283fd5b818401915084601f83011261024b578283fd5b81358181111561025d5761025d610348565b604051601f8201601f19908116603f0116810190838211818310171561028557610285610348565b8160405282815287602084870101111561029d578586fd5b826020860160208301379182016020019490945295945050505050565b6000602080835283518082850152825b818110156102e6578581018301518582016040015282016102ca565b818111156102f75783604083870101525b50601f01601f1916929092016040019392505050565b600181811c9082168061032157607f821691505b6020821081141561034257634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea26469706673582212204aed9197fd33119b13416a1a3e0f96874499832286e5bb5d75ca7b12324bf9a064736f6c63430008030033"

// DeployIPFSStorage deploys a new Ethereum contract, binding an instance of IPFSStorage to it.
func DeployIPFSStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IPFSStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(IPFSStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IPFSStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IPFSStorage{IPFSStorageCaller: IPFSStorageCaller{contract: contract}, IPFSStorageTransactor: IPFSStorageTransactor{contract: contract}, IPFSStorageFilterer: IPFSStorageFilterer{contract: contract}}, nil
}

// IPFSStorage is an auto generated Go binding around an Ethereum contract.
type IPFSStorage struct {
	IPFSStorageCaller     // Read-only binding to the contract
	IPFSStorageTransactor // Write-only binding to the contract
	IPFSStorageFilterer   // Log filterer for contract events
}

// IPFSStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPFSStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPFSStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPFSStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPFSStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPFSStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPFSStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPFSStorageSession struct {
	Contract     *IPFSStorage      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPFSStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPFSStorageCallerSession struct {
	Contract *IPFSStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IPFSStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPFSStorageTransactorSession struct {
	Contract     *IPFSStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IPFSStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPFSStorageRaw struct {
	Contract *IPFSStorage // Generic contract binding to access the raw methods on
}

// IPFSStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPFSStorageCallerRaw struct {
	Contract *IPFSStorageCaller // Generic read-only contract binding to access the raw methods on
}

// IPFSStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPFSStorageTransactorRaw struct {
	Contract *IPFSStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPFSStorage creates a new instance of IPFSStorage, bound to a specific deployed contract.
func NewIPFSStorage(address common.Address, backend bind.ContractBackend) (*IPFSStorage, error) {
	contract, err := bindIPFSStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPFSStorage{IPFSStorageCaller: IPFSStorageCaller{contract: contract}, IPFSStorageTransactor: IPFSStorageTransactor{contract: contract}, IPFSStorageFilterer: IPFSStorageFilterer{contract: contract}}, nil
}

// NewIPFSStorageCaller creates a new read-only instance of IPFSStorage, bound to a specific deployed contract.
func NewIPFSStorageCaller(address common.Address, caller bind.ContractCaller) (*IPFSStorageCaller, error) {
	contract, err := bindIPFSStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPFSStorageCaller{contract: contract}, nil
}

// NewIPFSStorageTransactor creates a new write-only instance of IPFSStorage, bound to a specific deployed contract.
func NewIPFSStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*IPFSStorageTransactor, error) {
	contract, err := bindIPFSStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPFSStorageTransactor{contract: contract}, nil
}

// NewIPFSStorageFilterer creates a new log filterer instance of IPFSStorage, bound to a specific deployed contract.
func NewIPFSStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*IPFSStorageFilterer, error) {
	contract, err := bindIPFSStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPFSStorageFilterer{contract: contract}, nil
}

// bindIPFSStorage binds a generic wrapper to an already deployed contract.
func bindIPFSStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPFSStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPFSStorage *IPFSStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPFSStorage.Contract.IPFSStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPFSStorage *IPFSStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPFSStorage.Contract.IPFSStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPFSStorage *IPFSStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPFSStorage.Contract.IPFSStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPFSStorage *IPFSStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPFSStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPFSStorage *IPFSStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPFSStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPFSStorage *IPFSStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPFSStorage.Contract.contract.Transact(opts, method, params...)
}

// GetEntry is a free data retrieval call binding the contract method 0x7db6a4e4.
//
// Solidity: function getEntry(address _address) view returns(string entry)
func (_IPFSStorage *IPFSStorageCaller) GetEntry(opts *bind.CallOpts, _address common.Address) (string, error) {
	var out []interface{}
	err := _IPFSStorage.contract.Call(opts, &out, "getEntry", _address)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEntry is a free data retrieval call binding the contract method 0x7db6a4e4.
//
// Solidity: function getEntry(address _address) view returns(string entry)
func (_IPFSStorage *IPFSStorageSession) GetEntry(_address common.Address) (string, error) {
	return _IPFSStorage.Contract.GetEntry(&_IPFSStorage.CallOpts, _address)
}

// GetEntry is a free data retrieval call binding the contract method 0x7db6a4e4.
//
// Solidity: function getEntry(address _address) view returns(string entry)
func (_IPFSStorage *IPFSStorageCallerSession) GetEntry(_address common.Address) (string, error) {
	return _IPFSStorage.Contract.GetEntry(&_IPFSStorage.CallOpts, _address)
}

// SetEntry is a paid mutator transaction binding the contract method 0x4b369820.
//
// Solidity: function setEntry(string _hash) returns()
func (_IPFSStorage *IPFSStorageTransactor) SetEntry(opts *bind.TransactOpts, _hash string) (*types.Transaction, error) {
	return _IPFSStorage.contract.Transact(opts, "setEntry", _hash)
}

// SetEntry is a paid mutator transaction binding the contract method 0x4b369820.
//
// Solidity: function setEntry(string _hash) returns()
func (_IPFSStorage *IPFSStorageSession) SetEntry(_hash string) (*types.Transaction, error) {
	return _IPFSStorage.Contract.SetEntry(&_IPFSStorage.TransactOpts, _hash)
}

// SetEntry is a paid mutator transaction binding the contract method 0x4b369820.
//
// Solidity: function setEntry(string _hash) returns()
func (_IPFSStorage *IPFSStorageTransactorSession) SetEntry(_hash string) (*types.Transaction, error) {
	return _IPFSStorage.Contract.SetEntry(&_IPFSStorage.TransactOpts, _hash)
}
