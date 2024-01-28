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

// BoutiqueNFTMetaData contains all meta data concerning the BoutiqueNFT contract.
var BoutiqueNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"safeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600160065534801562000015575f80fd5b506040518060400160405280600b81526020016a109bdd5d1a5c5d5953919560aa1b8152506040518060400160405280600381526020016210919560ea1b815250815f90816200006691906200011c565b5060016200007582826200011c565b505050620001e8565b634e487b7160e01b5f52604160045260245ffd5b600181811c90821680620000a757607f821691505b602082108103620000c657634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200011757805f5260205f20601f840160051c81016020851015620000f35750805b601f840160051c820191505b8181101562000114575f8155600101620000ff565b50505b505050565b81516001600160401b038111156200013857620001386200007e565b620001508162000149845462000092565b84620000cc565b602080601f83116001811462000186575f84156200016e5750858301515b5f19600386901b1c1916600185901b178555620001e0565b5f85815260208120601f198616915b82811015620001b65788860151825594840194600190910190840162000195565b5085821015620001d457878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b61102a80620001f65f395ff3fe608060405234801561000f575f80fd5b50600436106100f0575f3560e01c806342966c6811610093578063a22cb46511610063578063a22cb465146101f9578063b88d4fde1461020c578063c87b56dd1461021f578063e985e9c514610232575f80fd5b806342966c68146101aa5780636352211e146101bd57806370a08231146101d057806395d89b41146101f1575f80fd5b8063095ea7b3116100ce578063095ea7b31461015c57806323b872dd1461017157806340d097c31461018457806342842e0e14610197575f80fd5b806301ffc9a7146100f457806306fdde031461011c578063081812fc14610131575b5f80fd5b610107610102366004610c73565b610245565b60405190151581526020015b60405180910390f35b610124610296565b6040516101139190610cdb565b61014461013f366004610ced565b610325565b6040516001600160a01b039091168152602001610113565b61016f61016a366004610d1f565b61034c565b005b61016f61017f366004610d47565b61035b565b61016f610192366004610d80565b6103e9565b61016f6101a5366004610d47565b610409565b61016f6101b8366004610ced565b610428565b6101446101cb366004610ced565b610433565b6101e36101de366004610d80565b61043d565b604051908152602001610113565b610124610482565b61016f610207366004610d99565b610491565b61016f61021a366004610de6565b61049c565b61012461022d366004610ced565b6104b3565b610107610240366004610ebb565b6104e9565b5f6001600160e01b031982166380ac58cd60e01b148061027557506001600160e01b03198216635b5e139f60e01b145b8061029057506301ffc9a760e01b6001600160e01b03198316145b92915050565b60605f80546102a490610eec565b80601f01602080910402602001604051908101604052809291908181526020018280546102d090610eec565b801561031b5780601f106102f25761010080835404028352916020019161031b565b820191905f5260205f20905b8154815290600101906020018083116102fe57829003601f168201915b5050505050905090565b5f61032f82610516565b505f828152600460205260409020546001600160a01b0316610290565b61035782823361054e565b5050565b6001600160a01b03821661038957604051633250574960e11b81525f60048201526024015b60405180910390fd5b5f61039583833361055b565b9050836001600160a01b0316816001600160a01b0316146103e3576040516364283d7b60e01b81526001600160a01b0380861660048301526024820184905282166044820152606401610380565b50505050565b600680545f91826103f983610f24565b919050559050610357828261064d565b61042383838360405180602001604052805f81525061049c565b505050565b6103575f823361055b565b5f61029082610516565b5f6001600160a01b038216610467576040516322718ad960e21b81525f6004820152602401610380565b506001600160a01b03165f9081526003602052604090205490565b6060600180546102a490610eec565b610357338383610666565b6104a784848461035b565b6103e384848484610704565b60605f6104bf8361082a565b9050806040516020016104d29190610f48565b604051602081830303815290604052915050919050565b6001600160a01b039182165f90815260056020908152604080832093909416825291909152205460ff1690565b5f818152600260205260408120546001600160a01b03168061029057604051637e27328960e01b815260048101849052602401610380565b61042383838360016108b1565b5f828152600260205260408120546001600160a01b0390811690831615610587576105878184866109b5565b6001600160a01b038116156105c1576105a25f855f806108b1565b6001600160a01b0381165f90815260036020526040902080545f190190555b6001600160a01b038516156105ef576001600160a01b0385165f908152600360205260409020805460010190555b5f8481526002602052604080822080546001600160a01b0319166001600160a01b0389811691821790925591518793918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4949350505050565b610357828260405180602001604052805f815250610a19565b6001600160a01b03821661069857604051630b61174360e31b81526001600160a01b0383166004820152602401610380565b6001600160a01b038381165f81815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b0383163b156103e357604051630a85bd0160e11b81526001600160a01b0384169063150b7a0290610746903390889087908790600401610f6f565b6020604051808303815f875af1925050508015610780575060408051601f3d908101601f1916820190925261077d91810190610fab565b60015b6107e7573d8080156107ad576040519150601f19603f3d011682016040523d82523d5f602084013e6107b2565b606091505b5080515f036107df57604051633250574960e11b81526001600160a01b0385166004820152602401610380565b805181602001fd5b6001600160e01b03198116630a85bd0160e11b1461082357604051633250574960e11b81526001600160a01b0385166004820152602401610380565b5050505050565b606061083582610516565b505f6108716040805180820190915260208082527f68747470733a2f2f6e6674697175652e73332e616d617a6f6e6177732e636f6d9082015290565b90505f81511161088f5760405180602001604052805f8152506108aa565b8061089984610a2f565b6040516020016104d2929190610fc6565b9392505050565b80806108c557506001600160a01b03821615155b15610986575f6108d484610516565b90506001600160a01b038316158015906109005750826001600160a01b0316816001600160a01b031614155b8015610913575061091181846104e9565b155b1561093c5760405163a9fbf51f60e01b81526001600160a01b0384166004820152602401610380565b81156109845783856001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b50505f90815260046020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b6109c0838383610abf565b610423576001600160a01b0383166109ee57604051637e27328960e01b815260048101829052602401610380565b60405163177e802f60e01b81526001600160a01b038316600482015260248101829052604401610380565b610a238383610b23565b6104235f848484610704565b60605f610a3b83610b84565b60010190505f8167ffffffffffffffff811115610a5a57610a5a610dd2565b6040519080825280601f01601f191660200182016040528015610a84576020820181803683370190505b5090508181016020015b5f19016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a8504945084610a8e57509392505050565b5f6001600160a01b03831615801590610b1b5750826001600160a01b0316846001600160a01b03161480610af85750610af884846104e9565b80610b1b57505f828152600460205260409020546001600160a01b038481169116145b949350505050565b6001600160a01b038216610b4c57604051633250574960e11b81525f6004820152602401610380565b5f610b5883835f61055b565b90506001600160a01b03811615610423576040516339e3563760e11b81525f6004820152602401610380565b5f8072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b8310610bc25772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310610bee576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310610c0c57662386f26fc10000830492506010015b6305f5e1008310610c24576305f5e100830492506008015b6127108310610c3857612710830492506004015b60648310610c4a576064830492506002015b600a83106102905760010192915050565b6001600160e01b031981168114610c70575f80fd5b50565b5f60208284031215610c83575f80fd5b81356108aa81610c5b565b5f5b83811015610ca8578181015183820152602001610c90565b50505f910152565b5f8151808452610cc7816020860160208601610c8e565b601f01601f19169290920160200192915050565b602081525f6108aa6020830184610cb0565b5f60208284031215610cfd575f80fd5b5035919050565b80356001600160a01b0381168114610d1a575f80fd5b919050565b5f8060408385031215610d30575f80fd5b610d3983610d04565b946020939093013593505050565b5f805f60608486031215610d59575f80fd5b610d6284610d04565b9250610d7060208501610d04565b9150604084013590509250925092565b5f60208284031215610d90575f80fd5b6108aa82610d04565b5f8060408385031215610daa575f80fd5b610db383610d04565b915060208301358015158114610dc7575f80fd5b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b5f805f8060808587031215610df9575f80fd5b610e0285610d04565b9350610e1060208601610d04565b925060408501359150606085013567ffffffffffffffff80821115610e33575f80fd5b818701915087601f830112610e46575f80fd5b813581811115610e5857610e58610dd2565b604051601f8201601f19908116603f01168101908382118183101715610e8057610e80610dd2565b816040528281528a6020848701011115610e98575f80fd5b826020860160208301375f60208483010152809550505050505092959194509250565b5f8060408385031215610ecc575f80fd5b610ed583610d04565b9150610ee360208401610d04565b90509250929050565b600181811c90821680610f0057607f821691505b602082108103610f1e57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60018201610f4157634e487b7160e01b5f52601160045260245ffd5b5060010190565b5f8251610f59818460208701610c8e565b632e6a706760e01b920191825250600401919050565b6001600160a01b03858116825284166020820152604081018390526080606082018190525f90610fa190830184610cb0565b9695505050505050565b5f60208284031215610fbb575f80fd5b81516108aa81610c5b565b5f8351610fd7818460208801610c8e565b835190830190610feb818360208801610c8e565b0194935050505056fea2646970667358221220e0b6aad774be56a9877020add0232d62368dafc3de14e08f93c834fa69f8951664736f6c63430008180033",
}

// BoutiqueNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use BoutiqueNFTMetaData.ABI instead.
var BoutiqueNFTABI = BoutiqueNFTMetaData.ABI

// BoutiqueNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BoutiqueNFTMetaData.Bin instead.
var BoutiqueNFTBin = BoutiqueNFTMetaData.Bin

// DeployBoutiqueNFT deploys a new Ethereum contract, binding an instance of BoutiqueNFT to it.
func DeployBoutiqueNFT(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BoutiqueNFT, error) {
	parsed, err := BoutiqueNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BoutiqueNFTBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BoutiqueNFT{BoutiqueNFTCaller: BoutiqueNFTCaller{contract: contract}, BoutiqueNFTTransactor: BoutiqueNFTTransactor{contract: contract}, BoutiqueNFTFilterer: BoutiqueNFTFilterer{contract: contract}}, nil
}

// BoutiqueNFT is an auto generated Go binding around an Ethereum contract.
type BoutiqueNFT struct {
	BoutiqueNFTCaller     // Read-only binding to the contract
	BoutiqueNFTTransactor // Write-only binding to the contract
	BoutiqueNFTFilterer   // Log filterer for contract events
}

// BoutiqueNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoutiqueNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoutiqueNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoutiqueNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoutiqueNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoutiqueNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoutiqueNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoutiqueNFTSession struct {
	Contract     *BoutiqueNFT      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BoutiqueNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoutiqueNFTCallerSession struct {
	Contract *BoutiqueNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BoutiqueNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoutiqueNFTTransactorSession struct {
	Contract     *BoutiqueNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BoutiqueNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoutiqueNFTRaw struct {
	Contract *BoutiqueNFT // Generic contract binding to access the raw methods on
}

// BoutiqueNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoutiqueNFTCallerRaw struct {
	Contract *BoutiqueNFTCaller // Generic read-only contract binding to access the raw methods on
}

// BoutiqueNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoutiqueNFTTransactorRaw struct {
	Contract *BoutiqueNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBoutiqueNFT creates a new instance of BoutiqueNFT, bound to a specific deployed contract.
func NewBoutiqueNFT(address common.Address, backend bind.ContractBackend) (*BoutiqueNFT, error) {
	contract, err := bindBoutiqueNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BoutiqueNFT{BoutiqueNFTCaller: BoutiqueNFTCaller{contract: contract}, BoutiqueNFTTransactor: BoutiqueNFTTransactor{contract: contract}, BoutiqueNFTFilterer: BoutiqueNFTFilterer{contract: contract}}, nil
}

// NewBoutiqueNFTCaller creates a new read-only instance of BoutiqueNFT, bound to a specific deployed contract.
func NewBoutiqueNFTCaller(address common.Address, caller bind.ContractCaller) (*BoutiqueNFTCaller, error) {
	contract, err := bindBoutiqueNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoutiqueNFTCaller{contract: contract}, nil
}

// NewBoutiqueNFTTransactor creates a new write-only instance of BoutiqueNFT, bound to a specific deployed contract.
func NewBoutiqueNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*BoutiqueNFTTransactor, error) {
	contract, err := bindBoutiqueNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoutiqueNFTTransactor{contract: contract}, nil
}

// NewBoutiqueNFTFilterer creates a new log filterer instance of BoutiqueNFT, bound to a specific deployed contract.
func NewBoutiqueNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*BoutiqueNFTFilterer, error) {
	contract, err := bindBoutiqueNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoutiqueNFTFilterer{contract: contract}, nil
}

// bindBoutiqueNFT binds a generic wrapper to an already deployed contract.
func bindBoutiqueNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BoutiqueNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoutiqueNFT *BoutiqueNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BoutiqueNFT.Contract.BoutiqueNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoutiqueNFT *BoutiqueNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.BoutiqueNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoutiqueNFT *BoutiqueNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.BoutiqueNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoutiqueNFT *BoutiqueNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BoutiqueNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoutiqueNFT *BoutiqueNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoutiqueNFT *BoutiqueNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BoutiqueNFT *BoutiqueNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BoutiqueNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BoutiqueNFT *BoutiqueNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BoutiqueNFT.Contract.BalanceOf(&_BoutiqueNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BoutiqueNFT *BoutiqueNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BoutiqueNFT.Contract.BalanceOf(&_BoutiqueNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BoutiqueNFT *BoutiqueNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BoutiqueNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BoutiqueNFT *BoutiqueNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BoutiqueNFT.Contract.GetApproved(&_BoutiqueNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BoutiqueNFT *BoutiqueNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BoutiqueNFT.Contract.GetApproved(&_BoutiqueNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BoutiqueNFT *BoutiqueNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _BoutiqueNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BoutiqueNFT *BoutiqueNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BoutiqueNFT.Contract.IsApprovedForAll(&_BoutiqueNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BoutiqueNFT *BoutiqueNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BoutiqueNFT.Contract.IsApprovedForAll(&_BoutiqueNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BoutiqueNFT *BoutiqueNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BoutiqueNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BoutiqueNFT *BoutiqueNFTSession) Name() (string, error) {
	return _BoutiqueNFT.Contract.Name(&_BoutiqueNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BoutiqueNFT *BoutiqueNFTCallerSession) Name() (string, error) {
	return _BoutiqueNFT.Contract.Name(&_BoutiqueNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BoutiqueNFT *BoutiqueNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BoutiqueNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BoutiqueNFT *BoutiqueNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BoutiqueNFT.Contract.OwnerOf(&_BoutiqueNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BoutiqueNFT *BoutiqueNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BoutiqueNFT.Contract.OwnerOf(&_BoutiqueNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BoutiqueNFT *BoutiqueNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BoutiqueNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BoutiqueNFT *BoutiqueNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BoutiqueNFT.Contract.SupportsInterface(&_BoutiqueNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BoutiqueNFT *BoutiqueNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BoutiqueNFT.Contract.SupportsInterface(&_BoutiqueNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BoutiqueNFT *BoutiqueNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BoutiqueNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BoutiqueNFT *BoutiqueNFTSession) Symbol() (string, error) {
	return _BoutiqueNFT.Contract.Symbol(&_BoutiqueNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BoutiqueNFT *BoutiqueNFTCallerSession) Symbol() (string, error) {
	return _BoutiqueNFT.Contract.Symbol(&_BoutiqueNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BoutiqueNFT *BoutiqueNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _BoutiqueNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BoutiqueNFT *BoutiqueNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _BoutiqueNFT.Contract.TokenURI(&_BoutiqueNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BoutiqueNFT *BoutiqueNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _BoutiqueNFT.Contract.TokenURI(&_BoutiqueNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BoutiqueNFT *BoutiqueNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.Approve(&_BoutiqueNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.Approve(&_BoutiqueNFT.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFT.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_BoutiqueNFT *BoutiqueNFTSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.Burn(&_BoutiqueNFT.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.Burn(&_BoutiqueNFT.TransactOpts, tokenId)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactor) SafeMint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BoutiqueNFT.contract.Transact(opts, "safeMint", to)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns()
func (_BoutiqueNFT *BoutiqueNFTSession) SafeMint(to common.Address) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.SafeMint(&_BoutiqueNFT.TransactOpts, to)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactorSession) SafeMint(to common.Address) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.SafeMint(&_BoutiqueNFT.TransactOpts, to)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BoutiqueNFT *BoutiqueNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.SafeTransferFrom(&_BoutiqueNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.SafeTransferFrom(&_BoutiqueNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BoutiqueNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BoutiqueNFT *BoutiqueNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.SafeTransferFrom0(&_BoutiqueNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.SafeTransferFrom0(&_BoutiqueNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _BoutiqueNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BoutiqueNFT *BoutiqueNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.SetApprovalForAll(&_BoutiqueNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.SetApprovalForAll(&_BoutiqueNFT.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BoutiqueNFT *BoutiqueNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.TransferFrom(&_BoutiqueNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BoutiqueNFT *BoutiqueNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BoutiqueNFT.Contract.TransferFrom(&_BoutiqueNFT.TransactOpts, from, to, tokenId)
}

// BoutiqueNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BoutiqueNFT contract.
type BoutiqueNFTApprovalIterator struct {
	Event *BoutiqueNFTApproval // Event containing the contract specifics and raw log

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
func (it *BoutiqueNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoutiqueNFTApproval)
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
		it.Event = new(BoutiqueNFTApproval)
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
func (it *BoutiqueNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoutiqueNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoutiqueNFTApproval represents a Approval event raised by the BoutiqueNFT contract.
type BoutiqueNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BoutiqueNFT *BoutiqueNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BoutiqueNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BoutiqueNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BoutiqueNFTApprovalIterator{contract: _BoutiqueNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BoutiqueNFT *BoutiqueNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BoutiqueNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BoutiqueNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoutiqueNFTApproval)
				if err := _BoutiqueNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BoutiqueNFT *BoutiqueNFTFilterer) ParseApproval(log types.Log) (*BoutiqueNFTApproval, error) {
	event := new(BoutiqueNFTApproval)
	if err := _BoutiqueNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoutiqueNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the BoutiqueNFT contract.
type BoutiqueNFTApprovalForAllIterator struct {
	Event *BoutiqueNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BoutiqueNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoutiqueNFTApprovalForAll)
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
		it.Event = new(BoutiqueNFTApprovalForAll)
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
func (it *BoutiqueNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoutiqueNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoutiqueNFTApprovalForAll represents a ApprovalForAll event raised by the BoutiqueNFT contract.
type BoutiqueNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BoutiqueNFT *BoutiqueNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BoutiqueNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BoutiqueNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BoutiqueNFTApprovalForAllIterator{contract: _BoutiqueNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BoutiqueNFT *BoutiqueNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BoutiqueNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BoutiqueNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoutiqueNFTApprovalForAll)
				if err := _BoutiqueNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BoutiqueNFT *BoutiqueNFTFilterer) ParseApprovalForAll(log types.Log) (*BoutiqueNFTApprovalForAll, error) {
	event := new(BoutiqueNFTApprovalForAll)
	if err := _BoutiqueNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoutiqueNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BoutiqueNFT contract.
type BoutiqueNFTTransferIterator struct {
	Event *BoutiqueNFTTransfer // Event containing the contract specifics and raw log

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
func (it *BoutiqueNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoutiqueNFTTransfer)
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
		it.Event = new(BoutiqueNFTTransfer)
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
func (it *BoutiqueNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoutiqueNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoutiqueNFTTransfer represents a Transfer event raised by the BoutiqueNFT contract.
type BoutiqueNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BoutiqueNFT *BoutiqueNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BoutiqueNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BoutiqueNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BoutiqueNFTTransferIterator{contract: _BoutiqueNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BoutiqueNFT *BoutiqueNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BoutiqueNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BoutiqueNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoutiqueNFTTransfer)
				if err := _BoutiqueNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BoutiqueNFT *BoutiqueNFTFilterer) ParseTransfer(log types.Log) (*BoutiqueNFTTransfer, error) {
	event := new(BoutiqueNFTTransfer)
	if err := _BoutiqueNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
