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

// trafficStationupchainTrafficInfo is an auto generated low-level Go binding around an user-defined struct.
type trafficStationupchainTrafficInfo struct {
	TrafficID   *big.Int
	SourceAddr  common.Address
	TrafficInfo string
}

// trafficStationvoteInfo is an auto generated low-level Go binding around an user-defined struct.
type trafficStationvoteInfo struct {
	SourceAddr common.Address
	VoteAddr   common.Address
	TrafficID  *big.Int
	State      bool
}

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"trafficID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"trafficInfo\",\"type\":\"string\"}],\"name\":\"trafficTrans\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voteAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"trafficID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"voteTrans\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"trafficID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sourceAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"trafficInfo\",\"type\":\"string\"}],\"internalType\":\"structtrafficStation.upchainTrafficInfo\",\"name\":\"uti\",\"type\":\"tuple\"}],\"name\":\"emitTrafficTrans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voteAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"trafficID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"internalType\":\"structtrafficStation.voteInfo\",\"name\":\"vi\",\"type\":\"tuple\"}],\"name\":\"emitVoteTrans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"trafficID\",\"type\":\"uint256\"}],\"name\":\"getTrafficInfoByID\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"trafficID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sourceAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"trafficInfo\",\"type\":\"string\"}],\"internalType\":\"structtrafficStation.upchainTrafficInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"trafficID\",\"type\":\"uint256\"}],\"name\":\"getVoteInfoByID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x60806040523480156100115760006000fd5b50610017565b610c34806100266000396000f3fe60806040523480156100115760006000fd5b50600436106100515760003560e01c806314f6c7401461005757806367aad5fb14610087578063ab680e45146100b7578063e56816ee146100d357610051565b60006000fd5b610071600480360381019061006c919061072b565b6100ef565b60405161007e91906108d4565b60405180910390f35b6100a1600480360381019061009c919061072b565b610157565b6040516100ae91906108b1565b60405180910390f35b6100d160048036038101906100cc9190610700565b6102ca565b005b6100ed60048036038101906100e891906106bc565b6103e8565b005b6000600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000838152602001908152602001600020600050549050610152565b919050565b61015f6104dd565b600160005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000838152602001908152602001600020600050604051806060016040529081600082016000505481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201600050805461023990610aa0565b80601f016020809104026020016040519081016040528092919081815260200182805461026590610aa0565b80156102b25780601f10610287576101008083540402835291602001916102b2565b820191906000526020600020905b81548152906001019060200180831161029557829003601f168201915b50505050508152602001505090506102c5565b919050565b600073ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff161415151561030b5760006000fd5b6001151581606001511515141561039757600160006000506000836000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000836040015181526020019081526020016000206000828282505461038d91906109b8565b9250508190909055505b7fe6d2d2c21284e671da4a773c628d519f81d37af2ae3950c34d207f567bb08f4e81600001518260200151836040015184606001516040516103dc949392919061086b565b60405180910390a15b50565b806020015173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff16141515156104295760006000fd5b600060006000506000836020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000836000015181526020019081526020016000206000508190909055507f82a0cd03c34dff9e0879fd6d40cea94cdf58b4b4bbf4b2c8a3ebb91ca80075778160000151826020015183604001516040516104d1939291906108f0565b60405180910390a15b50565b604051806060016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001509056610bfd565b600061052e61052984610956565b61092f565b9050828152602081018484840111156105475760006000fd5b610552848285610a5b565b505b9392505050565b60008135905061056a81610bac565b5b92915050565b60008135905061058081610bc7565b5b92915050565b600082601f830112151561059b5760006000fd5b81356105ab84826020860161051b565b9150505b92915050565b6000606082840312156105c85760006000fd5b6105d2606061092f565b905060006105e2848285016106a6565b60008301525060206105f68482850161055b565b602083015250604082013567ffffffffffffffff8111156106175760006000fd5b61062384828501610587565b6040830152505b92915050565b6000608082840312156106435760006000fd5b61064d608061092f565b9050600061065d8482850161055b565b60008301525060206106718482850161055b565b6020830152506040610685848285016106a6565b604083015250606061069984828501610571565b6060830152505b92915050565b6000813590506106b581610be2565b5b92915050565b6000602082840312156106cf5760006000fd5b600082013567ffffffffffffffff8111156106ea5760006000fd5b6106f6848285016105b5565b9150505b92915050565b6000608082840312156107135760006000fd5b600061072184828501610630565b9150505b92915050565b60006020828403121561073e5760006000fd5b600061074c848285016106a6565b9150505b92915050565b61075f81610a0f565b82525b5050565b61076f81610a0f565b82525b5050565b61077f81610a22565b82525b5050565b600061079182610988565b61079b8185610994565b93506107ab818560208601610a6b565b6107b481610b9a565b84019150505b92915050565b60006107cb82610988565b6107d581856109a6565b93506107e5818560208601610a6b565b6107ee81610b9a565b84019150505b92915050565b6000606083016000830151610812600086018261084b565b5060208301516108256020860182610756565b506040830151848203604086015261083d8282610786565b915050809150505b92915050565b61085481610a50565b82525b5050565b61086481610a50565b82525b5050565b60006080820190506108806000830187610766565b61088d6020830186610766565b61089a604083018561085b565b6108a76060830184610776565b5b95945050505050565b600060208201905081810360008301526108cb81846107fa565b90505b92915050565b60006020820190506108e9600083018461085b565b5b92915050565b6000606082019050610905600083018661085b565b6109126020830185610766565b818103604083015261092481846107c0565b90505b949350505050565b600061093961094b565b90506109458282610ad5565b5b919050565b600060405190505b90565b600067ffffffffffffffff82111561097157610970610b69565b5b61097a82610b9a565b90506020810190505b919050565b6000815190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60006109c382610a50565b91506109ce83610a50565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610a0357610a02610b07565b5b82820190505b92915050565b6000610a1a82610a2f565b90505b919050565b600081151590505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b828183376000838301525b505050565b60005b83811015610a8a5780820151818401525b602081019050610a6e565b83811115610a99576000848401525b505b505050565b600060028204905060018216801515610aba57607f821691505b60208210811415610ace57610acd610b38565b5b505b919050565b610ade82610b9a565b810181811067ffffffffffffffff82111715610afd57610afc610b69565b5b80604052505b5050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b565b6000601f19601f83011690505b919050565b610bb581610a0f565b81141515610bc35760006000fd5b5b50565b610bd081610a22565b81141515610bde5760006000fd5b5b50565b610beb81610a50565b81141515610bf95760006000fd5b5b50565bfea26469706673582212209ca6957911d6a0f253e641fff16042a1177000de656c150e8480b1465c0dab1964736f6c63430008020033"

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

// GetTrafficInfoByID is a free data retrieval call binding the contract method 0x67aad5fb.
//
// Solidity: function getTrafficInfoByID(uint256 trafficID) view returns((uint256,address,string))
func (_Contract *ContractCaller) GetTrafficInfoByID(opts *bind.CallOpts, trafficID *big.Int) (trafficStationupchainTrafficInfo, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTrafficInfoByID", trafficID)

	if err != nil {
		return *new(trafficStationupchainTrafficInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(trafficStationupchainTrafficInfo)).(*trafficStationupchainTrafficInfo)

	return out0, err

}

// GetTrafficInfoByID is a free data retrieval call binding the contract method 0x67aad5fb.
//
// Solidity: function getTrafficInfoByID(uint256 trafficID) view returns((uint256,address,string))
func (_Contract *ContractSession) GetTrafficInfoByID(trafficID *big.Int) (trafficStationupchainTrafficInfo, error) {
	return _Contract.Contract.GetTrafficInfoByID(&_Contract.CallOpts, trafficID)
}

// GetTrafficInfoByID is a free data retrieval call binding the contract method 0x67aad5fb.
//
// Solidity: function getTrafficInfoByID(uint256 trafficID) view returns((uint256,address,string))
func (_Contract *ContractCallerSession) GetTrafficInfoByID(trafficID *big.Int) (trafficStationupchainTrafficInfo, error) {
	return _Contract.Contract.GetTrafficInfoByID(&_Contract.CallOpts, trafficID)
}

// GetVoteInfoByID is a free data retrieval call binding the contract method 0x14f6c740.
//
// Solidity: function getVoteInfoByID(uint256 trafficID) view returns(uint256)
func (_Contract *ContractCaller) GetVoteInfoByID(opts *bind.CallOpts, trafficID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getVoteInfoByID", trafficID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoteInfoByID is a free data retrieval call binding the contract method 0x14f6c740.
//
// Solidity: function getVoteInfoByID(uint256 trafficID) view returns(uint256)
func (_Contract *ContractSession) GetVoteInfoByID(trafficID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetVoteInfoByID(&_Contract.CallOpts, trafficID)
}

// GetVoteInfoByID is a free data retrieval call binding the contract method 0x14f6c740.
//
// Solidity: function getVoteInfoByID(uint256 trafficID) view returns(uint256)
func (_Contract *ContractCallerSession) GetVoteInfoByID(trafficID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetVoteInfoByID(&_Contract.CallOpts, trafficID)
}

// EmitTrafficTrans is a paid mutator transaction binding the contract method 0xe56816ee.
//
// Solidity: function emitTrafficTrans((uint256,address,string) uti) returns()
func (_Contract *ContractTransactor) EmitTrafficTrans(opts *bind.TransactOpts, uti trafficStationupchainTrafficInfo) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "emitTrafficTrans", uti)
}

// EmitTrafficTrans is a paid mutator transaction binding the contract method 0xe56816ee.
//
// Solidity: function emitTrafficTrans((uint256,address,string) uti) returns()
func (_Contract *ContractSession) EmitTrafficTrans(uti trafficStationupchainTrafficInfo) (*types.Transaction, error) {
	return _Contract.Contract.EmitTrafficTrans(&_Contract.TransactOpts, uti)
}

// EmitTrafficTrans is a paid mutator transaction binding the contract method 0xe56816ee.
//
// Solidity: function emitTrafficTrans((uint256,address,string) uti) returns()
func (_Contract *ContractTransactorSession) EmitTrafficTrans(uti trafficStationupchainTrafficInfo) (*types.Transaction, error) {
	return _Contract.Contract.EmitTrafficTrans(&_Contract.TransactOpts, uti)
}

// EmitVoteTrans is a paid mutator transaction binding the contract method 0xab680e45.
//
// Solidity: function emitVoteTrans((address,address,uint256,bool) vi) returns()
func (_Contract *ContractTransactor) EmitVoteTrans(opts *bind.TransactOpts, vi trafficStationvoteInfo) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "emitVoteTrans", vi)
}

// EmitVoteTrans is a paid mutator transaction binding the contract method 0xab680e45.
//
// Solidity: function emitVoteTrans((address,address,uint256,bool) vi) returns()
func (_Contract *ContractSession) EmitVoteTrans(vi trafficStationvoteInfo) (*types.Transaction, error) {
	return _Contract.Contract.EmitVoteTrans(&_Contract.TransactOpts, vi)
}

// EmitVoteTrans is a paid mutator transaction binding the contract method 0xab680e45.
//
// Solidity: function emitVoteTrans((address,address,uint256,bool) vi) returns()
func (_Contract *ContractTransactorSession) EmitVoteTrans(vi trafficStationvoteInfo) (*types.Transaction, error) {
	return _Contract.Contract.EmitVoteTrans(&_Contract.TransactOpts, vi)
}

// ContractTrafficTransIterator is returned from FilterTrafficTrans and is used to iterate over the raw logs and unpacked data for TrafficTrans events raised by the Contract contract.
type ContractTrafficTransIterator struct {
	Event *ContractTrafficTrans // Event containing the contract specifics and raw log

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
func (it *ContractTrafficTransIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTrafficTrans)
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
		it.Event = new(ContractTrafficTrans)
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
func (it *ContractTrafficTransIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTrafficTransIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTrafficTrans represents a TrafficTrans event raised by the Contract contract.
type ContractTrafficTrans struct {
	TrafficID   *big.Int
	SourceAddr  common.Address
	TrafficInfo string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTrafficTrans is a free log retrieval operation binding the contract event 0x82a0cd03c34dff9e0879fd6d40cea94cdf58b4b4bbf4b2c8a3ebb91ca8007577.
//
// Solidity: event trafficTrans(uint256 trafficID, address sourceAddr, string trafficInfo)
func (_Contract *ContractFilterer) FilterTrafficTrans(opts *bind.FilterOpts) (*ContractTrafficTransIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "trafficTrans")
	if err != nil {
		return nil, err
	}
	return &ContractTrafficTransIterator{contract: _Contract.contract, event: "trafficTrans", logs: logs, sub: sub}, nil
}

// WatchTrafficTrans is a free log subscription operation binding the contract event 0x82a0cd03c34dff9e0879fd6d40cea94cdf58b4b4bbf4b2c8a3ebb91ca8007577.
//
// Solidity: event trafficTrans(uint256 trafficID, address sourceAddr, string trafficInfo)
func (_Contract *ContractFilterer) WatchTrafficTrans(opts *bind.WatchOpts, sink chan<- *ContractTrafficTrans) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "trafficTrans")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTrafficTrans)
				if err := _Contract.contract.UnpackLog(event, "trafficTrans", log); err != nil {
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

// ParseTrafficTrans is a log parse operation binding the contract event 0x82a0cd03c34dff9e0879fd6d40cea94cdf58b4b4bbf4b2c8a3ebb91ca8007577.
//
// Solidity: event trafficTrans(uint256 trafficID, address sourceAddr, string trafficInfo)
func (_Contract *ContractFilterer) ParseTrafficTrans(log types.Log) (*ContractTrafficTrans, error) {
	event := new(ContractTrafficTrans)
	if err := _Contract.contract.UnpackLog(event, "trafficTrans", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVoteTransIterator is returned from FilterVoteTrans and is used to iterate over the raw logs and unpacked data for VoteTrans events raised by the Contract contract.
type ContractVoteTransIterator struct {
	Event *ContractVoteTrans // Event containing the contract specifics and raw log

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
func (it *ContractVoteTransIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVoteTrans)
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
		it.Event = new(ContractVoteTrans)
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
func (it *ContractVoteTransIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVoteTransIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVoteTrans represents a VoteTrans event raised by the Contract contract.
type ContractVoteTrans struct {
	SourceAddr common.Address
	VoteAddr   common.Address
	TrafficID  *big.Int
	State      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteTrans is a free log retrieval operation binding the contract event 0xe6d2d2c21284e671da4a773c628d519f81d37af2ae3950c34d207f567bb08f4e.
//
// Solidity: event voteTrans(address sourceAddr, address voteAddr, uint256 trafficID, bool state)
func (_Contract *ContractFilterer) FilterVoteTrans(opts *bind.FilterOpts) (*ContractVoteTransIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "voteTrans")
	if err != nil {
		return nil, err
	}
	return &ContractVoteTransIterator{contract: _Contract.contract, event: "voteTrans", logs: logs, sub: sub}, nil
}

// WatchVoteTrans is a free log subscription operation binding the contract event 0xe6d2d2c21284e671da4a773c628d519f81d37af2ae3950c34d207f567bb08f4e.
//
// Solidity: event voteTrans(address sourceAddr, address voteAddr, uint256 trafficID, bool state)
func (_Contract *ContractFilterer) WatchVoteTrans(opts *bind.WatchOpts, sink chan<- *ContractVoteTrans) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "voteTrans")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVoteTrans)
				if err := _Contract.contract.UnpackLog(event, "voteTrans", log); err != nil {
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

// ParseVoteTrans is a log parse operation binding the contract event 0xe6d2d2c21284e671da4a773c628d519f81d37af2ae3950c34d207f567bb08f4e.
//
// Solidity: event voteTrans(address sourceAddr, address voteAddr, uint256 trafficID, bool state)
func (_Contract *ContractFilterer) ParseVoteTrans(log types.Log) (*ContractVoteTrans, error) {
	event := new(ContractVoteTrans)
	if err := _Contract.contract.UnpackLog(event, "voteTrans", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
