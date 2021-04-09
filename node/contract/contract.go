// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"msgConn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"events\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"}],\"name\":\"indexRconn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sp\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"rconn\",\"type\":\"string[]\"}],\"name\":\"insertRconn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"}],\"name\":\"insertRddos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodesAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rconnSpeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"reCheckDDos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x6080604052348015620000125760006000fd5b505b5b33600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160006000506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b6000600090505b600160005080549050811015620001ca5760006006600050600060016000508481548110151562000135577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b8080620001c190620001e2565b915050620000db565b505b620002635662000262565b60008190505b919050565b6000620001ef82620001d7565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141562000225576200022462000231565b5b6001820190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b5b6115fe80620002736000396000f3fe60806040523480156100115760006000fd5b50600436106100825760003560e01c80634420e4861161005c5780634420e4861461010457806346db51a914610120578063b8b7db1b1461013c578063f495ae8c1461015857610082565b806306db1fd4146100885780630b791430146100a45780631161840f146100d457610082565b60006000fd5b6100a2600480360381019061009d9190610f00565b610188565b005b6100be60048036038101906100b99190610ed5565b6104ac565b6040516100cb919061117b565b60405180910390f35b6100ee60048036038101906100e99190610ed5565b6104d5565b6040516100fb91906110db565b60405180910390f35b61011e60048036038101906101199190610eaa565b61051b565b005b61013a60048036038101906101359190610ed5565b6106d2565b005b61015660048036038101906101519190610f80565b610a21565b005b610172600480360381019061016d9190610ed5565b610b79565b60405161017f919061117b565b60405180910390f35b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156101e45760006000fd5b6003600050600085815260200190815260200160002060005060050160009054906101000a900460ff16151515610250576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102479061115a565b60405180910390fd5b826003600050600086815260200190815260200160002060005060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816003600050600086815260200190815260200160002060005060020160005081909090555080600081518110151561030d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151600360005060008681526020019081526020016000206000506001016000509080519060200190610347929190610c04565b50806001815181101515610384577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516003600050600086815260200190815260200160002060005060030160005090805190602001906103be929190610c04565b508060028151811015156103fb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151600360005060008681526020019081526020016000206000506004016000509080519060200190610435929190610c04565b5060016003600050600086815260200190815260200160002060005060050160006101000a81548160ff0219169083151502179055507f15e0c0bad37064a057c88d56e6913f20bacef4e5264629992fe52ad599d1c4a4848460405161049c9291906111e7565b60405180910390a15b5b50505050565b600560005081815481106104bf57600080fd5b906000526020600020900160005b915090505481565b600160005081815481106104e857600080fd5b906000526020600020900160005b9150909054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156105785760006000fd5b600060005060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151561060a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610601906110f7565b60405180910390fd5b6001600060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160005081908060018154018082558091505060019003906000526020600020900160005b9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b50565b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561072e5760006000fd5b6003600050600082815260200190815260200160002060005060050160009054906101000a900460ff1615151561079a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079190611139565b60405180910390fd5b600560005081908060018154018082558091505060019003906000526020600020900160005b9091909190915090905560036000506000828152602001908152602001600020600050600460005060008381526020019081526020016000206000506000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600182016000508160010160005090805461087d9061131e565b610888929190610c8f565b5060028201600050548160020160005090905560038201600050816003016000509080546108b59061131e565b6108c0929190610c8f565b5060048201600050816004016000509080546108db9061131e565b6108e6929190610c8f565b506005820160009054906101000a900460ff168160050160006101000a81548160ff0219169083151502179055509050506003600050600082815260200190815260200160002060006000820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006109669190610d1c565b60028201600050600090556003820160006109819190610d1c565b6004820160006109919190610d1c565b6005820160006101000a81549060ff021916905550507f15e0c0bad37064a057c88d56e6913f20bacef4e5264629992fe52ad599d1c4a4816004600050600084815260200190815260200160002060005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051610a14929190611197565b60405180910390a15b5b50565b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610a7d5760006000fd5b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610b0e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0590611118565b60405180910390fd5b8082101515610b73576001600660005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b5b5b5050565b6000600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610bd75760006000fd5b60036000506000838152602001908152602001600020600050600201600050549050610bfe565b5b919050565b828054610c109061131e565b90600052602060002090601f016020900481019282610c325760008555610c7e565b82601f10610c4b57805160ff1916838001178555610c7e565b82800160010185558215610c7e579182015b82811115610c7d5782518260005090905591602001919060010190610c5d565b5b509050610c8b9190610d5c565b5090565b828054610c9b9061131e565b90600052602060002090601f016020900481019282610cbd5760008555610d0b565b82601f10610cce5780548555610d0b565b82800160010185558215610d0b57600052602060002091601f016020900482015b82811115610d0a578254825591600101919060010190610cef565b5b509050610d189190610d5c565b5090565b508054610d289061131e565b6000825580601f10610d3a5750610d59565b601f016020900490600052602060002090810190610d589190610d5c565b5b50565b610d61565b80821115610d7b5760008181506000905550600101610d61565b5090566115c7565b6000610d96610d918461125e565b611237565b9050808382526020820190508260005b85811015610dd75781358501610dbc8882610e66565b8452602084019350602083019250505b600181019050610da6565b5050505b9392505050565b6000610df5610df08461128b565b611237565b905082815260208101848484011115610e0e5760006000fd5b610e1984828561130e565b505b9392505050565b600081359050610e3181611591565b5b92915050565b600082601f8301121515610e4c5760006000fd5b8135610e5c848260208601610d83565b9150505b92915050565b600082601f8301121515610e7a5760006000fd5b8135610e8a848260208601610de2565b9150505b92915050565b600081359050610ea3816115ac565b5b92915050565b600060208284031215610ebd5760006000fd5b6000610ecb84828501610e22565b9150505b92915050565b600060208284031215610ee85760006000fd5b6000610ef684828501610e94565b9150505b92915050565b600060006000600060808587031215610f195760006000fd5b6000610f2787828801610e94565b9450506020610f3887828801610e22565b9350506040610f4987828801610e94565b925050606085013567ffffffffffffffff811115610f675760006000fd5b610f7387828801610e38565b9150505b92959194509250565b6000600060408385031215610f955760006000fd5b6000610fa385828601610e94565b9250506020610fb485828601610e94565b9150505b9250929050565b610fc8816112cf565b82525b5050565b6000610fdc6005836112bd565b9150610fe7826113f9565b6020820190505b919050565b60006110006005836112bd565b915061100b82611423565b6020820190505b919050565b6000611024601b836112bd565b915061102f8261144d565b6020820190505b919050565b60006110486023836112bd565b915061105382611477565b6040820190505b919050565b600061106c6037836112bd565b9150611077826114c7565b6040820190505b919050565b60006110906035836112bd565b915061109b82611517565b6040820190505b919050565b60006110b46006836112bd565b91506110bf82611567565b6020820190505b919050565b6110d481611303565b82525b5050565b60006020820190506110f06000830184610fbf565b5b92915050565b6000602082019050818103600083015261111081611017565b90505b919050565b600060208201905081810360008301526111318161103b565b90505b919050565b600060208201905081810360008301526111528161105f565b90505b919050565b6000602082019050818103600083015261117381611083565b90505b919050565b600060208201905061119060008301846110cb565b5b92915050565b60006080820190506111ac60008301856110cb565b6111b96020830184610fbf565b81810360408301526111ca816110a7565b905081810360608301526111dd81610fcf565b90505b9392505050565b60006080820190506111fc60008301856110cb565b6112096020830184610fbf565b818103604083015261121a816110a7565b9050818103606083015261122d81610ff3565b90505b9392505050565b6000611241611253565b905061124d8282611353565b5b919050565b600060405190505b90565b600067ffffffffffffffff821115611279576112786113b6565b5b6020820290506020810190505b919050565b600067ffffffffffffffff8211156112a6576112a56113b6565b5b6112af826113e7565b90506020810190505b919050565b60008282526020820190505b92915050565b60006112da826112e2565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b828183376000838301525b505050565b60006002820490506001821680151561133857607f821691505b6020821081141561134c5761134b611385565b5b505b919050565b61135c826113e7565b810181811067ffffffffffffffff8211171561137b5761137a6113b6565b5b80604052505b5050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b565b6000601f19601f83011690505b919050565b7f5264646f7300000000000000000000000000000000000000000000000000000060008201525b50565b7f52636f6e6e00000000000000000000000000000000000000000000000000000060008201525b50565b7f546865206e6f646520686173206265656e20726567697374656421000000000060008201525b50565b7f546865206e6f646527732061646472657373206973206e6f742072656769737460008201527f656421000000000000000000000000000000000000000000000000000000000060208201525b50565b7f54686973206576656e74494420686173206265656e2075736564202c2063616e60008201527f277420696e7365727420696e746f207264646f7365732100000000000000000060208201525b50565b7f54686973206576656e74494420686173206265656e2075736564202c2063616e60008201527f277420696e7365727420696e746f2072636f6e6e73000000000000000000000060208201525b50565b7f696e73657274000000000000000000000000000000000000000000000000000060008201525b50565b61159a816112cf565b811415156115a85760006000fd5b5b50565b6115b581611303565b811415156115c35760006000fd5b5b50565bfea264697066735822122080afafa881436471b26192bb1369175c6601827739860ada4b68f0ad8da99c1964736f6c63430008020033"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) Events(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "events", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint256)
func (_Contract *ContractSession) Events(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.Events(&_Contract.CallOpts, arg0)
}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) Events(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.Events(&_Contract.CallOpts, arg0)
}

// IndexRconn is a free data retrieval call binding the contract method 0xf495ae8c.
//
// Solidity: function indexRconn(uint256 eventID) view returns(uint256)
func (_Contract *ContractCaller) IndexRconn(opts *bind.CallOpts, eventID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "indexRconn", eventID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexRconn is a free data retrieval call binding the contract method 0xf495ae8c.
//
// Solidity: function indexRconn(uint256 eventID) view returns(uint256)
func (_Contract *ContractSession) IndexRconn(eventID *big.Int) (*big.Int, error) {
	return _Contract.Contract.IndexRconn(&_Contract.CallOpts, eventID)
}

// IndexRconn is a free data retrieval call binding the contract method 0xf495ae8c.
//
// Solidity: function indexRconn(uint256 eventID) view returns(uint256)
func (_Contract *ContractCallerSession) IndexRconn(eventID *big.Int) (*big.Int, error) {
	return _Contract.Contract.IndexRconn(&_Contract.CallOpts, eventID)
}

// NodesAddr is a free data retrieval call binding the contract method 0x1161840f.
//
// Solidity: function nodesAddr(uint256 ) view returns(address)
func (_Contract *ContractCaller) NodesAddr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nodesAddr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodesAddr is a free data retrieval call binding the contract method 0x1161840f.
//
// Solidity: function nodesAddr(uint256 ) view returns(address)
func (_Contract *ContractSession) NodesAddr(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.NodesAddr(&_Contract.CallOpts, arg0)
}

// NodesAddr is a free data retrieval call binding the contract method 0x1161840f.
//
// Solidity: function nodesAddr(uint256 ) view returns(address)
func (_Contract *ContractCallerSession) NodesAddr(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.NodesAddr(&_Contract.CallOpts, arg0)
}

// InsertRconn is a paid mutator transaction binding the contract method 0x06db1fd4.
//
// Solidity: function insertRconn(uint256 eventID, address addr, uint256 sp, string[] rconn) returns()
func (_Contract *ContractTransactor) InsertRconn(opts *bind.TransactOpts, eventID *big.Int, addr common.Address, sp *big.Int, rconn []string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "insertRconn", eventID, addr, sp, rconn)
}

// InsertRconn is a paid mutator transaction binding the contract method 0x06db1fd4.
//
// Solidity: function insertRconn(uint256 eventID, address addr, uint256 sp, string[] rconn) returns()
func (_Contract *ContractSession) InsertRconn(eventID *big.Int, addr common.Address, sp *big.Int, rconn []string) (*types.Transaction, error) {
	return _Contract.Contract.InsertRconn(&_Contract.TransactOpts, eventID, addr, sp, rconn)
}

// InsertRconn is a paid mutator transaction binding the contract method 0x06db1fd4.
//
// Solidity: function insertRconn(uint256 eventID, address addr, uint256 sp, string[] rconn) returns()
func (_Contract *ContractTransactorSession) InsertRconn(eventID *big.Int, addr common.Address, sp *big.Int, rconn []string) (*types.Transaction, error) {
	return _Contract.Contract.InsertRconn(&_Contract.TransactOpts, eventID, addr, sp, rconn)
}

// InsertRddos is a paid mutator transaction binding the contract method 0x46db51a9.
//
// Solidity: function insertRddos(uint256 eventID) returns()
func (_Contract *ContractTransactor) InsertRddos(opts *bind.TransactOpts, eventID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "insertRddos", eventID)
}

// InsertRddos is a paid mutator transaction binding the contract method 0x46db51a9.
//
// Solidity: function insertRddos(uint256 eventID) returns()
func (_Contract *ContractSession) InsertRddos(eventID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.InsertRddos(&_Contract.TransactOpts, eventID)
}

// InsertRddos is a paid mutator transaction binding the contract method 0x46db51a9.
//
// Solidity: function insertRddos(uint256 eventID) returns()
func (_Contract *ContractTransactorSession) InsertRddos(eventID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.InsertRddos(&_Contract.TransactOpts, eventID)
}

// ReCheckDDos is a paid mutator transaction binding the contract method 0xb8b7db1b.
//
// Solidity: function reCheckDDos(uint256 rconnSpeed, uint256 threshold) returns()
func (_Contract *ContractTransactor) ReCheckDDos(opts *bind.TransactOpts, rconnSpeed *big.Int, threshold *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "reCheckDDos", rconnSpeed, threshold)
}

// ReCheckDDos is a paid mutator transaction binding the contract method 0xb8b7db1b.
//
// Solidity: function reCheckDDos(uint256 rconnSpeed, uint256 threshold) returns()
func (_Contract *ContractSession) ReCheckDDos(rconnSpeed *big.Int, threshold *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReCheckDDos(&_Contract.TransactOpts, rconnSpeed, threshold)
}

// ReCheckDDos is a paid mutator transaction binding the contract method 0xb8b7db1b.
//
// Solidity: function reCheckDDos(uint256 rconnSpeed, uint256 threshold) returns()
func (_Contract *ContractTransactorSession) ReCheckDDos(rconnSpeed *big.Int, threshold *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReCheckDDos(&_Contract.TransactOpts, rconnSpeed, threshold)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address addr) returns()
func (_Contract *ContractTransactor) Register(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "register", addr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address addr) returns()
func (_Contract *ContractSession) Register(addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Register(&_Contract.TransactOpts, addr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address addr) returns()
func (_Contract *ContractTransactorSession) Register(addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Register(&_Contract.TransactOpts, addr)
}

// ContractMsgConnIterator is returned from FilterMsgConn and is used to iterate over the raw logs and unpacked data for MsgConn events raised by the Contract contract.
type ContractMsgConnIterator struct {
	Event *ContractMsgConn // Event containing the contract specifics and raw log

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
func (it *ContractMsgConnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMsgConn)
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
		it.Event = new(ContractMsgConn)
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
func (it *ContractMsgConnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMsgConnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMsgConn represents a MsgConn event raised by the Contract contract.
type ContractMsgConn struct {
	EventID  *big.Int
	Addr     common.Address
	TypeName string
	Name     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMsgConn is a free log retrieval operation binding the contract event 0x15e0c0bad37064a057c88d56e6913f20bacef4e5264629992fe52ad599d1c4a4.
//
// Solidity: event msgConn(uint256 eventID, address addr, string typeName, string name)
func (_Contract *ContractFilterer) FilterMsgConn(opts *bind.FilterOpts) (*ContractMsgConnIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "msgConn")
	if err != nil {
		return nil, err
	}
	return &ContractMsgConnIterator{contract: _Contract.contract, event: "msgConn", logs: logs, sub: sub}, nil
}

// WatchMsgConn is a free log subscription operation binding the contract event 0x15e0c0bad37064a057c88d56e6913f20bacef4e5264629992fe52ad599d1c4a4.
//
// Solidity: event msgConn(uint256 eventID, address addr, string typeName, string name)
func (_Contract *ContractFilterer) WatchMsgConn(opts *bind.WatchOpts, sink chan<- *ContractMsgConn) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "msgConn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMsgConn)
				if err := _Contract.contract.UnpackLog(event, "msgConn", log); err != nil {
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

// ParseMsgConn is a log parse operation binding the contract event 0x15e0c0bad37064a057c88d56e6913f20bacef4e5264629992fe52ad599d1c4a4.
//
// Solidity: event msgConn(uint256 eventID, address addr, string typeName, string name)
func (_Contract *ContractFilterer) ParseMsgConn(log types.Log) (*ContractMsgConn, error) {
	event := new(ContractMsgConn)
	if err := _Contract.contract.UnpackLog(event, "msgConn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
