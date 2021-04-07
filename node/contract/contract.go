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
	TargetIP     string
	Speed        *big.Int
	Timestamp    string
	Descrptions  string
	IsUsed       bool
}

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"msgConn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"events\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"}],\"name\":\"indexRconn\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commiterAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"targetIP\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"speed\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"descrptions\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isUsed\",\"type\":\"bool\"}],\"internalType\":\"structDDoS.Rconn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sp\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"rconn\",\"type\":\"string[]\"}],\"name\":\"insertRconn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"}],\"name\":\"insertRddos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodesAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commiterAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"targetIP\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"speed\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"descrptions\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isUsed\",\"type\":\"bool\"}],\"internalType\":\"structDDoS.Rconn\",\"name\":\"rconn\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"reCheckDDos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x6080604052348015620000125760006000fd5b505b5b33600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160006000506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b6000600090505b600160005080549050811015620001ca5760006006600050600060016000508481548110151562000135577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b8080620001c190620001e2565b915050620000db565b505b620002635662000262565b60008190505b919050565b6000620001ef82620001d7565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141562000225576200022462000231565b5b6001820190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b5b611df480620002736000396000f3fe60806040523480156100115760006000fd5b50600436106100825760003560e01c80631161840f1161005c5780631161840f146100f05780634420e4861461012057806346db51a91461013c578063f495ae8c1461015857610082565b806306db1fd4146100885780630965fbee146100a45780630b791430146100c057610082565b60006000fd5b6100a2600480360381019061009d9190611485565b610188565b005b6100be60048036038101906100b99190611402565b6104ac565b005b6100da60048036038101906100d5919061145a565b61069d565b6040516100e7919061182b565b60405180910390f35b61010a6004803603810190610105919061145a565b6106c6565b6040516101179190611747565b60405180910390f35b61013a600480360381019061013591906113d7565b61070c565b005b6101566004803603810190610151919061145a565b6108c3565b005b610172600480360381019061016d919061145a565b610d0c565b60405161017f9190611808565b60405180910390f35b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156101e45760006000fd5b6003600050600085815260200190815260200160002060005060050160009054906101000a900460ff16151515610250576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610247906117e7565b60405180910390fd5b826003600050600086815260200190815260200160002060005060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816003600050600086815260200190815260200160002060005060020160005081909090555080600081518110151561030d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151600360005060008681526020019081526020016000206000506001016000509080519060200190610347929190610fe1565b50806001815181101515610384577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516003600050600086815260200190815260200160002060005060030160005090805190602001906103be929190610fe1565b508060028151811015156103fb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151600360005060008681526020019081526020016000206000506004016000509080519060200190610435929190610fe1565b5060016003600050600086815260200190815260200160002060005060050160006101000a81548160ff0219169083151502179055507f15e0c0bad37064a057c88d56e6913f20bacef4e5264629992fe52ad599d1c4a4848460405161049c929190611897565b60405180910390a15b5b50505050565b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156105085760006000fd5b60006000506000836000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561059d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610594906117c6565b60405180910390fd5b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561062e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062590611784565b60405180910390fd5b808260400151101515610697576001600660005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b5b5b5050565b600560005081815481106106b057600080fd5b906000526020600020900160005b915090505481565b600160005081815481106106d957600080fd5b906000526020600020900160005b9150909054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156107695760006000fd5b600060005060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515156107fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f290611763565b60405180910390fd5b6001600060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160005081908060018154018082558091505060019003906000526020600020900160005b9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b50565b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561091f5760006000fd5b6000600090505b600160005080549050811015610a17576000151560066000506000600160005084815481101515610980577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151415610a035760006000fd5b5b8080610a0f90611a95565b915050610926565b506003600050600082815260200190815260200160002060005060050160009054906101000a900460ff16151515610a84576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7b906117a5565b60405180910390fd5b600560005081908060018154018082558091505060019003906000526020600020900160005b9091909190915090905560036000506000828152602001908152602001600020600050600460005060008381526020019081526020016000206000506000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001820160005081600101600050908054610b6790611a2e565b610b7292919061106c565b506002820160005054816002016000509090556003820160005081600301600050908054610b9f90611a2e565b610baa92919061106c565b506004820160005081600401600050908054610bc590611a2e565b610bd092919061106c565b506005820160009054906101000a900460ff168160050160006101000a81548160ff0219169083151502179055509050506003600050600082815260200190815260200160002060006000820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182016000610c5091906110f9565b6002820160005060009055600382016000610c6b91906110f9565b600482016000610c7b91906110f9565b6005820160006101000a81549060ff021916905550507f15e0c0bad37064a057c88d56e6913f20bacef4e5264629992fe52ad599d1c4a4816004600050600084815260200190815260200160002060005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051610cfe929190611847565b60405180910390a15b5b5b50565b610d14611139565b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610d705760006000fd5b600360005060008381526020019081526020016000206000506040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182016000508054610dfd90611a2e565b80601f0160208091040260200160405190810160405280929190818152602001828054610e2990611a2e565b8015610e765780601f10610e4b57610100808354040283529160200191610e76565b820191906000526020600020905b815481529060010190602001808311610e5957829003601f168201915b5050505050815260200160028201600050548152602001600382016000508054610e9f90611a2e565b80601f0160208091040260200160405190810160405280929190818152602001828054610ecb90611a2e565b8015610f185780601f10610eed57610100808354040283529160200191610f18565b820191906000526020600020905b815481529060010190602001808311610efb57829003601f168201915b50505050508152602001600482016000508054610f3490611a2e565b80601f0160208091040260200160405190810160405280929190818152602001828054610f6090611a2e565b8015610fad5780601f10610f8257610100808354040283529160200191610fad565b820191906000526020600020905b815481529060010190602001808311610f9057829003601f168201915b505050505081526020016005820160009054906101000a900460ff1615151515815260200150509050610fdb565b5b919050565b828054610fed90611a2e565b90600052602060002090601f01602090048101928261100f576000855561105b565b82601f1061102857805160ff191683800117855561105b565b8280016001018555821561105b579182015b8281111561105a578251826000509090559160200191906001019061103a565b5b509050611068919061118a565b5090565b82805461107890611a2e565b90600052602060002090601f01602090048101928261109a57600085556110e8565b82601f106110ab57805485556110e8565b828001600101855582156110e857600052602060002091601f016020900482015b828111156110e75782548255916001019190600101906110cc565b5b5090506110f5919061118a565b5090565b50805461110590611a2e565b6000825580601f106111175750611136565b601f016020900490600052602060002090810190611135919061118a565b5b50565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016000815260200160608152602001606081526020016000151581526020015090565b61118f565b808211156111a9576000818150600090555060010161118f565b509056611dbd565b60006111c46111bf8461190e565b6118e7565b9050808382526020820190508260005b8581101561120557813585016111ea88826112aa565b8452602084019350602083019250505b6001810190506111d4565b5050505b9392505050565b600061122361121e8461193b565b6118e7565b90508281526020810184848401111561123c5760006000fd5b6112478482856119e9565b505b9392505050565b60008135905061125f81611d6c565b5b92915050565b600082601f830112151561127a5760006000fd5b813561128a8482602086016111b1565b9150505b92915050565b6000813590506112a381611d87565b5b92915050565b600082601f83011215156112be5760006000fd5b81356112ce848260208601611210565b9150505b92915050565b600060c082840312156112eb5760006000fd5b6112f560c06118e7565b9050600061130584828501611250565b600083015250602082013567ffffffffffffffff8111156113265760006000fd5b611332848285016112aa565b6020830152506040611346848285016113c1565b604083015250606082013567ffffffffffffffff8111156113675760006000fd5b611373848285016112aa565b606083015250608082013567ffffffffffffffff8111156113945760006000fd5b6113a0848285016112aa565b60808301525060a06113b484828501611294565b60a0830152505b92915050565b6000813590506113d081611da2565b5b92915050565b6000602082840312156113ea5760006000fd5b60006113f884828501611250565b9150505b92915050565b60006000604083850312156114175760006000fd5b600083013567ffffffffffffffff8111156114325760006000fd5b61143e858286016112d8565b925050602061144f858286016113c1565b9150505b9250929050565b60006020828403121561146d5760006000fd5b600061147b848285016113c1565b9150505b92915050565b60006000600060006080858703121561149e5760006000fd5b60006114ac878288016113c1565b94505060206114bd87828801611250565b93505060406114ce878288016113c1565b925050606085013567ffffffffffffffff8111156114ec5760006000fd5b6114f887828801611266565b9150505b92959194509250565b61150e8161199d565b82525b5050565b61151e8161199d565b82525b5050565b61152e816119b0565b82525b5050565b60006115408261196d565b61154a8185611979565b935061155a8185602086016119f9565b61156381611b72565b84019150505b92915050565b600061157c60058361198b565b915061158782611b84565b6020820190505b919050565b60006115a060058361198b565b91506115ab82611bae565b6020820190505b919050565b60006115c4601b8361198b565b91506115cf82611bd8565b6020820190505b919050565b60006115e860238361198b565b91506115f382611c02565b6040820190505b919050565b600061160c60378361198b565b915061161782611c52565b6040820190505b919050565b600061163060278361198b565b915061163b82611ca2565b6040820190505b919050565b600061165460358361198b565b915061165f82611cf2565b6040820190505b919050565b600061167860068361198b565b915061168382611d42565b6020820190505b919050565b600060c0830160008301516116a76000860182611505565b50602083015184820360208601526116bf8282611535565b91505060408301516116d46040860182611727565b50606083015184820360608601526116ec8282611535565b915050608083015184820360808601526117068282611535565b91505060a083015161171b60a0860182611525565b50809150505b92915050565b611730816119de565b82525b5050565b611740816119de565b82525b5050565b600060208201905061175c6000830184611515565b5b92915050565b6000602082019050818103600083015261177c816115b7565b90505b919050565b6000602082019050818103600083015261179d816115db565b90505b919050565b600060208201905081810360008301526117be816115ff565b90505b919050565b600060208201905081810360008301526117df81611623565b90505b919050565b6000602082019050818103600083015261180081611647565b90505b919050565b60006020820190508181036000830152611822818461168f565b90505b92915050565b60006020820190506118406000830184611737565b5b92915050565b600060808201905061185c6000830185611737565b6118696020830184611515565b818103604083015261187a8161166b565b9050818103606083015261188d8161156f565b90505b9392505050565b60006080820190506118ac6000830185611737565b6118b96020830184611515565b81810360408301526118ca8161166b565b905081810360608301526118dd81611593565b90505b9392505050565b60006118f1611903565b90506118fd8282611a63565b5b919050565b600060405190505b90565b600067ffffffffffffffff82111561192957611928611b41565b5b6020820290506020810190505b919050565b600067ffffffffffffffff82111561195657611955611b41565b5b61195f82611b72565b90506020810190505b919050565b6000815190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60006119a8826119bd565b90505b919050565b600081151590505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b828183376000838301525b505050565b60005b83811015611a185780820151818401525b6020810190506119fc565b83811115611a27576000848401525b505b505050565b600060028204905060018216801515611a4857607f821691505b60208210811415611a5c57611a5b611b10565b5b505b919050565b611a6c82611b72565b810181811067ffffffffffffffff82111715611a8b57611a8a611b41565b5b80604052505b5050565b6000611aa0826119de565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611ad357611ad2611adf565b5b6001820190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b565b6000601f19601f83011690505b919050565b7f5264646f7300000000000000000000000000000000000000000000000000000060008201525b50565b7f52636f6e6e00000000000000000000000000000000000000000000000000000060008201525b50565b7f546865206e6f646520686173206265656e20726567697374656421000000000060008201525b50565b7f546865206e6f646527732061646472657373206973206e6f742072656769737460008201527f656421000000000000000000000000000000000000000000000000000000000060208201525b50565b7f54686973206576656e74494420686173206265656e2075736564202c2063616e60008201527f277420696e7365727420696e746f207264646f7365732100000000000000000060208201525b50565b7f54686520636f6d6d6974657227732061646472657373206973206e6f7420726560008201527f676973746564210000000000000000000000000000000000000000000000000060208201525b50565b7f54686973206576656e74494420686173206265656e2075736564202c2063616e60008201527f277420696e7365727420696e746f2072636f6e6e73000000000000000000000060208201525b50565b7f696e73657274000000000000000000000000000000000000000000000000000060008201525b50565b611d758161199d565b81141515611d835760006000fd5b5b50565b611d90816119b0565b81141515611d9e5760006000fd5b5b50565b611dab816119de565b81141515611db95760006000fd5b5b50565bfea26469706673582212209774f46f934750b3b280d3317706a6951001277121440b88886f961e2902844464736f6c63430008020033"

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
// Solidity: function indexRconn(uint256 eventID) view returns((address,string,uint256,string,string,bool))
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
// Solidity: function indexRconn(uint256 eventID) view returns((address,string,uint256,string,string,bool))
func (_Contract *ContractSession) IndexRconn(eventID *big.Int) (DDoSRconn, error) {
	return _Contract.Contract.IndexRconn(&_Contract.CallOpts, eventID)
}

// IndexRconn is a free data retrieval call binding the contract method 0xf495ae8c.
//
// Solidity: function indexRconn(uint256 eventID) view returns((address,string,uint256,string,string,bool))
func (_Contract *ContractCallerSession) IndexRconn(eventID *big.Int) (DDoSRconn, error) {
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

// ReCheckDDos is a paid mutator transaction binding the contract method 0x0965fbee.
//
// Solidity: function reCheckDDos((address,string,uint256,string,string,bool) rconn, uint256 threshold) returns()
func (_Contract *ContractTransactor) ReCheckDDos(opts *bind.TransactOpts, rconn DDoSRconn, threshold *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "reCheckDDos", rconn, threshold)
}

// ReCheckDDos is a paid mutator transaction binding the contract method 0x0965fbee.
//
// Solidity: function reCheckDDos((address,string,uint256,string,string,bool) rconn, uint256 threshold) returns()
func (_Contract *ContractSession) ReCheckDDos(rconn DDoSRconn, threshold *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReCheckDDos(&_Contract.TransactOpts, rconn, threshold)
}

// ReCheckDDos is a paid mutator transaction binding the contract method 0x0965fbee.
//
// Solidity: function reCheckDDos((address,string,uint256,string,string,bool) rconn, uint256 threshold) returns()
func (_Contract *ContractTransactorSession) ReCheckDDos(rconn DDoSRconn, threshold *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReCheckDDos(&_Contract.TransactOpts, rconn, threshold)
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
