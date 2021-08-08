package node

import (
	"context"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/ethclient"

	"main/node/contract"
)

// TODO: 改成struct，减少依赖关系
var (
	Client *ethclient.Client

	Instance *contract.Contract
	Auth     *bind.TransactOpts

	FullURL string
)

// RunNode 连接到geth节点，完成配置初始化
func RunNode() error {
	FullURL := "ws://" + Conf.ChainAddress + ":" + Conf.ChainPort

	client, err := ethclient.Dial(FullURL)
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

	WriteJsonMessage(Message{
		TypeName: DETECTTYPE,
		Content: nil,
	})
}

func SendVote(trafficID *big.Int, sourceAddr common.Address, voteState bool) {
	auth := consultWithNode(Conf.Client.ClientPrivateAddr)

	nonce, err := Client.PendingNonceAt(context.Background(), common.HexToAddress(Conf.Client.ClientPublicAddr))
	if err != nil {
		log.Fatalln("[x] Error to pending nonce!")
	}

	auth.Nonce = new(big.Int).SetUint64(nonce)

	tmp, err := Instance.EmitVoteTrans(auth, contract.TrafficStationvoteInfo{
		SourceAddr: sourceAddr,
		VoteAddr:   common.HexToAddress(Conf.Client.ClientPublicAddr),
		TrafficID:  trafficID,
		State:      voteState,
	})
	if err != nil {
		log.Fatalf("[x] Send vote transaction with error message: %s\n", err)
		return
	}

	WriteJsonMessage(Message{
		TypeName: VOTETYPE,
		Content: map[string]MessageType{
			"sourceAddr": MessageType(sourceAddr.String()),
			"voteAddr": MessageType(Conf.Client.ClientPublicAddr),
			"trafficID": MessageType(tmp.Hash().String()),
			"state": MessageType(strconv.FormatBool(voteState)),
		},
	})
}
