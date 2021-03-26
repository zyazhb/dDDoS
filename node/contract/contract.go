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

// DDoSRconn is an auto generated low-level Go binding around an user-defined struct.
type DDoSRconn struct {
	CommiterAddr common.Address
	FromIP       string
	TargetIP     string
	Speed        *big.Int
	Timestamp    string
	IsDDoS       bool
	IsUsed       bool
}

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"msgConn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"events\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"}],\"name\":\"indexRconn\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commiterAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"fromIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetIP\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"speed\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isDDoS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isUsed\",\"type\":\"bool\"}],\"internalType\":\"structDDoS.Rconn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sp\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"rconn\",\"type\":\"string[]\"}],\"name\":\"insertRconn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"}],\"name\":\"insertRddos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x60806040523480156100115760006000fd5b505b33600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160006000506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b6100d7565b611650806100e66000396000f3fe60806040523480156100115760006000fd5b506004361061005c5760003560e01c806306db1fd4146100625780630b7914301461007e5780634420e486146100ae57806346db51a9146100ca578063f495ae8c146100e65761005c565b60006000fd5b61007c60048036038101906100779190610f48565b610116565b005b61009860048036038101906100939190610f1d565b61046f565b6040516100a59190611237565b60405180910390f35b6100c860048036038101906100c39190610ef2565b610498565b005b6100e460048036038101906100df9190610f1d565b6105e7565b005b61010060048036038101906100fb9190610f1d565b610902565b60405161010d9190611214565b60405180910390f35b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156101725760006000fd5b6002600050600085815260200190815260200160002060005060050160019054906101000a900460ff161515156101de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d5906111f3565b60405180910390fd5b826002600050600086815260200190815260200160002060005060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000815181101515610275577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516002600050600086815260200190815260200160002060005060010160005090805190602001906102af929190610bf2565b508060018151811015156102ec577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151600260005060008681526020019081526020016000206000506002016000509080519060200190610326929190610bf2565b508160026000506000868152602001908152602001600020600050600301600050819090905550806002815181101515610389577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516002600050600086815260200190815260200160002060005060040160005090805190602001906103c3929190610bf2565b5060006002600050600086815260200190815260200160002060005060050160006101000a81548160ff02191690831515021790555060016002600050600086815260200190815260200160002060005060050160016101000a81548160ff0219169083151502179055507f15e0c0bad37064a057c88d56e6913f20bacef4e5264629992fe52ad599d1c4a4848460405161045f929190611253565b60405180910390a15b5b50505050565b6004600050818154811061048257600080fd5b906000526020600020900160005b915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156104f55760006000fd5b600060005060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151515610587576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057e906111b1565b60405180910390fd5b6001600060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b5b50565b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156106435760006000fd5b6002600050600082815260200190815260200160002060005060050160019054906101000a900460ff161515156106af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a6906111d2565b60405180910390fd5b600460005081908060018154018082558091505060019003906000526020600020900160005b9091909190915090905560026000506000828152602001908152602001600020600050600360005060008381526020019081526020016000206000506000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001820160005081600101600050908054610792906113ea565b61079d929190610c7d565b5060028201600050816002016000509080546107b8906113ea565b6107c3929190610c7d565b5060038201600050548160030160005090905560048201600050816004016000509080546107f0906113ea565b6107fb929190610c7d565b506005820160009054906101000a900460ff168160050160006101000a81548160ff0219169083151502179055506005820160019054906101000a900460ff168160050160016101000a81548160ff0219169083151502179055509050506002600050600082815260200190815260200160002060006000820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006108a89190610d0a565b6002820160006108b89190610d0a565b60038201600050600090556004820160006108d39190610d0a565b6005820160006101000a81549060ff02191690556005820160016101000a81549060ff021916905550505b5b50565b61090a610d4a565b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156109665760006000fd5b600260005060008381526020019081526020016000206000506040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160005080546109f3906113ea565b80601f0160208091040260200160405190810160405280929190818152602001828054610a1f906113ea565b8015610a6c5780601f10610a4157610100808354040283529160200191610a6c565b820191906000526020600020905b815481529060010190602001808311610a4f57829003601f168201915b50505050508152602001600282016000508054610a88906113ea565b80601f0160208091040260200160405190810160405280929190818152602001828054610ab4906113ea565b8015610b015780601f10610ad657610100808354040283529160200191610b01565b820191906000526020600020905b815481529060010190602001808311610ae457829003601f168201915b5050505050815260200160038201600050548152602001600482016000508054610b2a906113ea565b80601f0160208091040260200160405190810160405280929190818152602001828054610b56906113ea565b8015610ba35780601f10610b7857610100808354040283529160200191610ba3565b820191906000526020600020905b815481529060010190602001808311610b8657829003601f168201915b505050505081526020016005820160009054906101000a900460ff161515151581526020016005820160019054906101000a900460ff1615151515815260200150509050610bec565b5b919050565b828054610bfe906113ea565b90600052602060002090601f016020900481019282610c205760008555610c6c565b82601f10610c3957805160ff1916838001178555610c6c565b82800160010185558215610c6c579182015b82811115610c6b5782518260005090905591602001919060010190610c4b565b5b509050610c799190610da4565b5090565b828054610c89906113ea565b90600052602060002090601f016020900481019282610cab5760008555610cf9565b82601f10610cbc5780548555610cf9565b82800160010185558215610cf957600052602060002091601f016020900482015b82811115610cf8578254825591600101919060010190610cdd565b5b509050610d069190610da4565b5090565b508054610d16906113ea565b6000825580601f10610d285750610d47565b601f016020900490600052602060002090810190610d469190610da4565b5b50565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160008152602001606081526020016000151581526020016000151581526020015090565b610da9565b80821115610dc35760008181506000905550600101610da9565b509056611619565b6000610dde610dd9846112ca565b6112a3565b9050808382526020820190508260005b85811015610e1f5781358501610e048882610eae565b8452602084019350602083019250505b600181019050610dee565b5050505b9392505050565b6000610e3d610e38846112f7565b6112a3565b905082815260208101848484011115610e565760006000fd5b610e618482856113a5565b505b9392505050565b600081359050610e79816115e3565b5b92915050565b600082601f8301121515610e945760006000fd5b8135610ea4848260208601610dcb565b9150505b92915050565b600082601f8301121515610ec25760006000fd5b8135610ed2848260208601610e2a565b9150505b92915050565b600081359050610eeb816115fe565b5b92915050565b600060208284031215610f055760006000fd5b6000610f1384828501610e6a565b9150505b92915050565b600060208284031215610f305760006000fd5b6000610f3e84828501610edc565b9150505b92915050565b600060006000600060808587031215610f615760006000fd5b6000610f6f87828801610edc565b9450506020610f8087828801610e6a565b9350506040610f9187828801610edc565b925050606085013567ffffffffffffffff811115610faf5760006000fd5b610fbb87828801610e80565b9150505b92959194509250565b610fd181611359565b82525b5050565b610fe181611359565b82525b5050565b610ff18161136c565b82525b5050565b600061100382611329565b61100d8185611335565b935061101d8185602086016113b5565b611026816114b3565b84019150505b92915050565b600061103f600583611347565b915061104a826114c5565b6020820190505b919050565b6000611063601b83611347565b915061106e826114ef565b6020820190505b919050565b6000611087603783611347565b915061109282611519565b6040820190505b919050565b60006110ab603583611347565b91506110b682611569565b6040820190505b919050565b60006110cf600683611347565b91506110da826115b9565b6020820190505b919050565b600060e0830160008301516110fe6000860182610fc8565b50602083015184820360208601526111168282610ff8565b915050604083015184820360408601526111308282610ff8565b91505060608301516111456060860182611191565b506080830151848203608086015261115d8282610ff8565b91505060a083015161117260a0860182610fe8565b5060c083015161118560c0860182610fe8565b50809150505b92915050565b61119a8161139a565b82525b5050565b6111aa8161139a565b82525b5050565b600060208201905081810360008301526111ca81611056565b90505b919050565b600060208201905081810360008301526111eb8161107a565b90505b919050565b6000602082019050818103600083015261120c8161109e565b90505b919050565b6000602082019050818103600083015261122e81846110e6565b90505b92915050565b600060208201905061124c60008301846111a1565b5b92915050565b600060808201905061126860008301856111a1565b6112756020830184610fd8565b8181036040830152611286816110c2565b9050818103606083015261129981611032565b90505b9392505050565b60006112ad6112bf565b90506112b9828261141f565b5b919050565b600060405190505b90565b600067ffffffffffffffff8211156112e5576112e4611482565b5b6020820290506020810190505b919050565b600067ffffffffffffffff82111561131257611311611482565b5b61131b826114b3565b90506020810190505b919050565b6000815190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b600061136482611379565b90505b919050565b600081151590505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b828183376000838301525b505050565b60005b838110156113d45780820151818401525b6020810190506113b8565b838111156113e3576000848401525b505b505050565b60006002820490506001821680151561140457607f821691505b6020821081141561141857611417611451565b5b505b919050565b611428826114b3565b810181811067ffffffffffffffff8211171561144757611446611482565b5b80604052505b5050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b565b6000601f19601f83011690505b919050565b7f52636f6e6e00000000000000000000000000000000000000000000000000000060008201525b50565b7f546865206e6f646520686173206265656e20726567697374656421000000000060008201525b50565b7f54686973206576656e74494420686173206265656e2075736564202c2063616e60008201527f277420696e7365727420696e746f207264646f7365732100000000000000000060208201525b50565b7f54686973206576656e74494420686173206265656e2075736564202c2063616e60008201527f277420696e7365727420696e746f2072636f6e6e73000000000000000000000060208201525b50565b7f696e73657274000000000000000000000000000000000000000000000000000060008201525b50565b6115ec81611359565b811415156115fa5760006000fd5b5b50565b6116078161139a565b811415156116155760006000fd5b5b50565bfea264697066735822122062faa2ffa7d5e37c19e70b02f342a11f52be11a40f9e01d033612bbab9ce211c64736f6c63430008020033"

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
// Solidity: function indexRconn(uint256 eventID) view returns((address,string,string,uint256,string,bool,bool))
func (_Contract *ContractCaller) IndexRconn(opts *bind.CallOpts, eventID *big.Int) (DDoSRconn, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "indexRconn", eventID)

	if err != nil {
		return *new(DDoSRconn), err
	}

	out0 := *abi.ConvertType(out[0], new(DDoSRconn)).(*DDoSRconn)

	return out0, err

}

// IndexRconn is a free data retrieval call binding the contract method 0xf495ae8c.
//
// Solidity: function indexRconn(uint256 eventID) view returns((address,string,string,uint256,string,bool,bool))
func (_Contract *ContractSession) IndexRconn(eventID *big.Int) (DDoSRconn, error) {
	return _Contract.Contract.IndexRconn(&_Contract.CallOpts, eventID)
}

// IndexRconn is a free data retrieval call binding the contract method 0xf495ae8c.
//
// Solidity: function indexRconn(uint256 eventID) view returns((address,string,string,uint256,string,bool,bool))
func (_Contract *ContractCallerSession) IndexRconn(eventID *big.Int) (DDoSRconn, error) {
	return _Contract.Contract.IndexRconn(&_Contract.CallOpts, eventID)
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
