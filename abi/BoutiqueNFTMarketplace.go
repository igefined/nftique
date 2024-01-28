// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// BoutiqueNFTMarketplaceMetaData contains all meta data concerning the BoutiqueNFTMarketplace contract.
var BoutiqueNFTMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"NFTAdded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161086238038061086283398101604081905261002e91610060565b60018054336001600160a01b0319918216179091555f80549091166001600160a01b039290921691909117905561008d565b5f60208284031215610070575f80fd5b81516001600160a01b0381168114610086575f80fd5b9392505050565b6107c88061009a5f395ff3fe60806040526004361061006e575f3560e01c8063b688a3631161004c578063b688a36314610103578063bfb231d214610117578063c57981b514610150578063d96a094a14610164575f80fd5b8063025e7c2714610072578063771602f7146100c357806396131049146100e4575b5f80fd5b34801561007d575f80fd5b506100a661008c366004610685565b60036020525f90815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156100ce575f80fd5b506100e26100dd36600461069c565b610177565b005b3480156100ef575f80fd5b506100e26100fe3660046106d3565b6103a5565b34801561010e575f80fd5b506100e2610424565b348015610122575f80fd5b50610142610131366004610685565b60026020525f908152604090205481565b6040519081526020016100ba565b34801561015b575f80fd5b50610142600181565b6100e2610172366004610685565b6104b6565b5f80546040516331a9108f60e11b8152600481018590526001600160a01b0390911690636352211e90602401602060405180830381865afa1580156101be573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101e291906106f5565b90506001600160a01b03811633146101f8575f80fd5b5f5460405163e985e9c560e01b81526001600160a01b0383811660048301523060248301529091169063e985e9c590604401602060405180830381865afa158015610245573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102699190610710565b806102e457505f5460405163020604bf60e21b81526004810185905230916001600160a01b03169063081812fc90602401602060405180830381865afa1580156102b5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102d991906106f5565b6001600160a01b0316145b6103285760405162461bcd60e51b815260206004820152601060248201526f1a5b9d985b1a5908185c1c1c9bdd985b60821b60448201526064015b60405180910390fd5b5f8211610333575f80fd5b5f838152600260209081526040808320859055600382529182902080546001600160a01b0319166001600160a01b03851690811790915582518581529182015284917f90cb21d5235a673aa1ffd63bf09c73d28454d541af736a13618a551f5774ca6e910160405180910390a2505050565b6001546001600160a01b031633146103ee5760405162461bcd60e51b815260206004820152600c60248201526b3737ba1030b71037bbb732b960a11b604482015260640161031f565b6040516001600160a01b038216904780156108fc02915f818181858888f19350505050158015610420573d5f803e3d5ffd5b5050565b335f818152600460205260409020541561049b5760405162461bcd60e51b815260206004820152603260248201527f426f7574697175654e46544d61726b6574706c6163653a207061727469636970604482015271616e7420616c72656164792065786973747360701b606482015260840161031f565b6001600160a01b03165f908152600460205260409020429055565b5f80546040516331a9108f60e11b8152600481018490526001600160a01b0390911690636352211e90602401602060405180830381865afa1580156104fd573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061052191906106f5565b6001600160a01b031603610533575f80fd5b5f81815260026020526040902054610549575f80fd5b5f818152600360209081526040808320546002909252909120546001600160a01b039091169034811461057a575f80fd5b5f54604051632142170760e11b81526001600160a01b03848116600483015233602483015260448201869052909116906342842e0e906064015f604051808303815f87803b1580156105ca575f80fd5b505af11580156105dc573d5f803e3d5ffd5b5050505f8481526002602090815260408083208390556003909152812080546001600160a01b031916905590506001600160a01b0383166064610620600185610743565b61062a9190610760565b610634908461077f565b6040515f81818185875af1925050503d805f811461066d576040519150601f19603f3d011682016040523d82523d5f602084013e610672565b606091505b505090508061067f575f80fd5b50505050565b5f60208284031215610695575f80fd5b5035919050565b5f80604083850312156106ad575f80fd5b50508035926020909101359150565b6001600160a01b03811681146106d0575f80fd5b50565b5f602082840312156106e3575f80fd5b81356106ee816106bc565b9392505050565b5f60208284031215610705575f80fd5b81516106ee816106bc565b5f60208284031215610720575f80fd5b815180151581146106ee575f80fd5b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761075a5761075a61072f565b92915050565b5f8261077a57634e487b7160e01b5f52601260045260245ffd5b500490565b8181038181111561075a5761075a61072f56fea2646970667358221220ffb2cf06ba1a8542929f6bb8d14251c3885dee189b813eb1b506020e780fe60764736f6c63430008180033",
}

// BoutiqueNFTMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use BoutiqueNFTMarketplaceMetaData.ABI instead.
var BoutiqueNFTMarketplaceABI = BoutiqueNFTMarketplaceMetaData.ABI

// BoutiqueNFTMarketplaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BoutiqueNFTMarketplaceMetaData.Bin instead.
var BoutiqueNFTMarketplaceBin = BoutiqueNFTMarketplaceMetaData.Bin

// DeployBoutiqueNFTMarketplace deploys a new Ethereum contract, binding an instance of BoutiqueNFTMarketplace to it.
func DeployBoutiqueNFTMarketplace(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *BoutiqueNFTMarketplace, error) {
	parsed, err := BoutiqueNFTMarketplaceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BoutiqueNFTMarketplaceBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BoutiqueNFTMarketplace{BoutiqueNFTMarketplaceCaller: BoutiqueNFTMarketplaceCaller{contract: contract}, BoutiqueNFTMarketplaceTransactor: BoutiqueNFTMarketplaceTransactor{contract: contract}, BoutiqueNFTMarketplaceFilterer: BoutiqueNFTMarketplaceFilterer{contract: contract}}, nil
}

// BoutiqueNFTMarketplace is an auto generated Go binding around an Ethereum contract.
type BoutiqueNFTMarketplace struct {
	BoutiqueNFTMarketplaceCaller     // Read-only binding to the contract
	BoutiqueNFTMarketplaceTransactor // Write-only binding to the contract
	BoutiqueNFTMarketplaceFilterer   // Log filterer for contract events
}

// BoutiqueNFTMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoutiqueNFTMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoutiqueNFTMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoutiqueNFTMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoutiqueNFTMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoutiqueNFTMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoutiqueNFTMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoutiqueNFTMarketplaceSession struct {
	Contract     *BoutiqueNFTMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BoutiqueNFTMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoutiqueNFTMarketplaceCallerSession struct {
	Contract *BoutiqueNFTMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// BoutiqueNFTMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoutiqueNFTMarketplaceTransactorSession struct {
	Contract     *BoutiqueNFTMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// BoutiqueNFTMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoutiqueNFTMarketplaceRaw struct {
	Contract *BoutiqueNFTMarketplace // Generic contract binding to access the raw methods on
}

// BoutiqueNFTMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoutiqueNFTMarketplaceCallerRaw struct {
	Contract *BoutiqueNFTMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// BoutiqueNFTMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoutiqueNFTMarketplaceTransactorRaw struct {
	Contract *BoutiqueNFTMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBoutiqueNFTMarketplace creates a new instance of BoutiqueNFTMarketplace, bound to a specific deployed contract.
func NewBoutiqueNFTMarketplace(address common.Address, backend bind.ContractBackend) (*BoutiqueNFTMarketplace, error) {
	contract, err := bindBoutiqueNFTMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BoutiqueNFTMarketplace{BoutiqueNFTMarketplaceCaller: BoutiqueNFTMarketplaceCaller{contract: contract}, BoutiqueNFTMarketplaceTransactor: BoutiqueNFTMarketplaceTransactor{contract: contract}, BoutiqueNFTMarketplaceFilterer: BoutiqueNFTMarketplaceFilterer{contract: contract}}, nil
}

// NewBoutiqueNFTMarketplaceCaller creates a new read-only instance of BoutiqueNFTMarketplace, bound to a specific deployed contract.
func NewBoutiqueNFTMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*BoutiqueNFTMarketplaceCaller, error) {
	contract, err := bindBoutiqueNFTMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoutiqueNFTMarketplaceCaller{contract: contract}, nil
}

// NewBoutiqueNFTMarketplaceTransactor creates a new write-only instance of BoutiqueNFTMarketplace, bound to a specific deployed contract.
func NewBoutiqueNFTMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*BoutiqueNFTMarketplaceTransactor, error) {
	contract, err := bindBoutiqueNFTMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoutiqueNFTMarketplaceTransactor{contract: contract}, nil
}

// NewBoutiqueNFTMarketplaceFilterer creates a new log filterer instance of BoutiqueNFTMarketplace, bound to a specific deployed contract.
func NewBoutiqueNFTMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*BoutiqueNFTMarketplaceFilterer, error) {
	contract, err := bindBoutiqueNFTMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoutiqueNFTMarketplaceFilterer{contract: contract}, nil
}

// bindBoutiqueNFTMarketplace binds a generic wrapper to an already deployed contract.
func bindBoutiqueNFTMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BoutiqueNFTMarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BoutiqueNFTMarketplace.Contract.BoutiqueNFTMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.Contract.BoutiqueNFTMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.Contract.BoutiqueNFTMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BoutiqueNFTMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.Contract.contract.Transact(opts, method, params...)
}

// FEE is a free data retrieval call binding the contract method 0xc57981b5.
//
// Solidity: function FEE() view returns(uint256)
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceCaller) FEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoutiqueNFTMarketplace.contract.Call(opts, &out, "FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEE is a free data retrieval call binding the contract method 0xc57981b5.
//
// Solidity: function FEE() view returns(uint256)
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceSession) FEE() (*big.Int, error) {
	return _BoutiqueNFTMarketplace.Contract.FEE(&_BoutiqueNFTMarketplace.CallOpts)
}

// FEE is a free data retrieval call binding the contract method 0xc57981b5.
//
// Solidity: function FEE() view returns(uint256)
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceCallerSession) FEE() (*big.Int, error) {
	return _BoutiqueNFTMarketplace.Contract.FEE(&_BoutiqueNFTMarketplace.CallOpts)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 tokenId) view returns(uint256 price)
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceCaller) Items(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BoutiqueNFTMarketplace.contract.Call(opts, &out, "items", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 tokenId) view returns(uint256 price)
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceSession) Items(tokenId *big.Int) (*big.Int, error) {
	return _BoutiqueNFTMarketplace.Contract.Items(&_BoutiqueNFTMarketplace.CallOpts, tokenId)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 tokenId) view returns(uint256 price)
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceCallerSession) Items(tokenId *big.Int) (*big.Int, error) {
	return _BoutiqueNFTMarketplace.Contract.Items(&_BoutiqueNFTMarketplace.CallOpts, tokenId)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 tokenId) view returns(address owner)
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceCaller) Owners(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BoutiqueNFTMarketplace.contract.Call(opts, &out, "owners", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 tokenId) view returns(address owner)
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceSession) Owners(tokenId *big.Int) (common.Address, error) {
	return _BoutiqueNFTMarketplace.Contract.Owners(&_BoutiqueNFTMarketplace.CallOpts, tokenId)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 tokenId) view returns(address owner)
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceCallerSession) Owners(tokenId *big.Int) (common.Address, error) {
	return _BoutiqueNFTMarketplace.Contract.Owners(&_BoutiqueNFTMarketplace.CallOpts, tokenId)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 _tokenId, uint256 _price) returns()
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceTransactor) Add(opts *bind.TransactOpts, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.contract.Transact(opts, "add", _tokenId, _price)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 _tokenId, uint256 _price) returns()
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceSession) Add(_tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.Contract.Add(&_BoutiqueNFTMarketplace.TransactOpts, _tokenId, _price)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 _tokenId, uint256 _price) returns()
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceTransactorSession) Add(_tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.Contract.Add(&_BoutiqueNFTMarketplace.TransactOpts, _tokenId, _price)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 _tokenId) payable returns()
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceTransactor) Buy(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.contract.Transact(opts, "buy", _tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 _tokenId) payable returns()
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceSession) Buy(_tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.Contract.Buy(&_BoutiqueNFTMarketplace.TransactOpts, _tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 _tokenId) payable returns()
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceTransactorSession) Buy(_tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.Contract.Buy(&_BoutiqueNFTMarketplace.TransactOpts, _tokenId)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceTransactor) Join(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.contract.Transact(opts, "join")
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceSession) Join() (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.Contract.Join(&_BoutiqueNFTMarketplace.TransactOpts)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceTransactorSession) Join() (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.Contract.Join(&_BoutiqueNFTMarketplace.TransactOpts)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x96131049.
//
// Solidity: function withdrawal(address _to) returns()
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceTransactor) Withdrawal(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.contract.Transact(opts, "withdrawal", _to)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x96131049.
//
// Solidity: function withdrawal(address _to) returns()
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceSession) Withdrawal(_to common.Address) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.Contract.Withdrawal(&_BoutiqueNFTMarketplace.TransactOpts, _to)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x96131049.
//
// Solidity: function withdrawal(address _to) returns()
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceTransactorSession) Withdrawal(_to common.Address) (*types.Transaction, error) {
	return _BoutiqueNFTMarketplace.Contract.Withdrawal(&_BoutiqueNFTMarketplace.TransactOpts, _to)
}

// BoutiqueNFTMarketplaceNFTAddedIterator is returned from FilterNFTAdded and is used to iterate over the raw logs and unpacked data for NFTAdded events raised by the BoutiqueNFTMarketplace contract.
type BoutiqueNFTMarketplaceNFTAddedIterator struct {
	Event *BoutiqueNFTMarketplaceNFTAdded // Event containing the contract specifics and raw log

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
func (it *BoutiqueNFTMarketplaceNFTAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoutiqueNFTMarketplaceNFTAdded)
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
		it.Event = new(BoutiqueNFTMarketplaceNFTAdded)
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
func (it *BoutiqueNFTMarketplaceNFTAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoutiqueNFTMarketplaceNFTAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoutiqueNFTMarketplaceNFTAdded represents a NFTAdded event raised by the BoutiqueNFTMarketplace contract.
type BoutiqueNFTMarketplaceNFTAdded struct {
	TokenId *big.Int
	Price   *big.Int
	Seller  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNFTAdded is a free log retrieval operation binding the contract event 0x90cb21d5235a673aa1ffd63bf09c73d28454d541af736a13618a551f5774ca6e.
//
// Solidity: event NFTAdded(uint256 indexed tokenId, uint256 price, address seller)
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceFilterer) FilterNFTAdded(opts *bind.FilterOpts, tokenId []*big.Int) (*BoutiqueNFTMarketplaceNFTAddedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BoutiqueNFTMarketplace.contract.FilterLogs(opts, "NFTAdded", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BoutiqueNFTMarketplaceNFTAddedIterator{contract: _BoutiqueNFTMarketplace.contract, event: "NFTAdded", logs: logs, sub: sub}, nil
}

// WatchNFTAdded is a free log subscription operation binding the contract event 0x90cb21d5235a673aa1ffd63bf09c73d28454d541af736a13618a551f5774ca6e.
//
// Solidity: event NFTAdded(uint256 indexed tokenId, uint256 price, address seller)
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceFilterer) WatchNFTAdded(opts *bind.WatchOpts, sink chan<- *BoutiqueNFTMarketplaceNFTAdded, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BoutiqueNFTMarketplace.contract.WatchLogs(opts, "NFTAdded", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoutiqueNFTMarketplaceNFTAdded)
				if err := _BoutiqueNFTMarketplace.contract.UnpackLog(event, "NFTAdded", log); err != nil {
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

// ParseNFTAdded is a log parse operation binding the contract event 0x90cb21d5235a673aa1ffd63bf09c73d28454d541af736a13618a551f5774ca6e.
//
// Solidity: event NFTAdded(uint256 indexed tokenId, uint256 price, address seller)
func (_BoutiqueNFTMarketplace *BoutiqueNFTMarketplaceFilterer) ParseNFTAdded(log types.Log) (*BoutiqueNFTMarketplaceNFTAdded, error) {
	event := new(BoutiqueNFTMarketplaceNFTAdded)
	if err := _BoutiqueNFTMarketplace.contract.UnpackLog(event, "NFTAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
