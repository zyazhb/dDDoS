package node

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/ethclient"

	"main/node/contract"
)

var (
	Client *ethclient.Client

	Instance   *contract.Contract
	Auth       *bind.TransactOpts

	FromAddress common.Address
)

// RunNode 连接到geth节点，完成配置初始化
func RunNode() error {
	fullURL := fmt.Sprintf("ws://%s:%d", Conf.ChainAddress, Conf.ChainPort)

	client, err := ethclient.Dial(fullURL)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("We have a connection to ethereum")

	Client = client
	Auth = consultWithNode(Conf.Client.ClientPrivateAddr)
	Instance = connectToContract(Conf.Server.ContractAddr)

	return nil
}

// ConsultWithNode 依据用户私钥获取身份认证
func consultWithNode(privateKeys string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(privateKeys)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = nil
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	return auth
}

// ConnectToContract 依据合约地址连接到私链上的合约
func connectToContract(contractAddr string) *contract.Contract {
	address := common.HexToAddress(contractAddr)
	instances, err := contract.NewContract(address, Client)
	if err != nil {
		log.Fatal(err)
	}

	return instances
}

func SendMessage(trafficInfo string) {
	// read -> public key | write/event -> private key
	auth := consultWithNode(Conf.Client.ClientPrivateAddr)

	trafficID, err := pendingTrafficIDAt(context.Background(), Conf.Client.ClientPublicAddr)
	if err != nil {
		log.Fatalf("Initial trafficID with error: %v\n", err)
	}

	_, err = Instance.EmitTrafficTrans(auth, contract.TrafficStationupchainTrafficInfo{
		TrafficID:   trafficID,
		SourceAddr:  common.HexToAddress(Conf.Client.ClientPublicAddr),
		TrafficInfo: trafficInfo,
	})
	if err != nil {
		log.Fatalf("[x] Send transaction with error message: %s\n", err)
		return
	}
}

func SendVote(trafficID *big.Int, sourceAddr common.Address, voteState bool) {
	auth := consultWithNode(Conf.Client.ClientPrivateAddr)

	_, err := Instance.EmitVoteTrans(auth, contract.TrafficStationvoteInfo{
		SourceAddr: sourceAddr,
		VoteAddr: common.HexToAddress(Conf.Client.ClientPublicAddr),
		TrafficID: trafficID,
		State: voteState,
	})
	if err != nil {
		log.Fatalf("[x] Send vote transaction with error message: %s\n", err)
		return
	}
}
