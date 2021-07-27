package node

import (
	"context"
	"crypto/ecdsa"
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
	fullURL := fmt.Sprintf("ws://%s:%d", Conf.chainAddress, Conf.chainPort)

	client, err := ethclient.Dial(fullURL)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("We have a connection to ethereum")

	Client = client
	Auth = consultWithNode(Conf.Client.clientAddr)
	Instance = connectToContract(Conf.Server.contractAddr)

	return nil
}

// ConsultWithNode 依据用户私钥获取身份认证
func consultWithNode(privateKeys string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(privateKeys)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	FromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := UpdateNonce()
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
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

// UpdateNonce 更新Nonce高度
func UpdateNonce() (uint64, error) {
	return Client.PendingNonceAt(context.Background(), FromAddress)
}

// WatchMessage 监听链上message信息

