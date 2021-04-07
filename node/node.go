package node

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/ethclient"

	"main/node/contract"
)

var (
	Client *ethclient.Client

	Instance   *contract.Contract
	Auth       *bind.TransactOpts
	SenderAddr common.Address

	FromAddress common.Address
)

const (
	ContractAddr = "0xB3785f7f25d7eCd60D30539CeB0096B137D0EE07"
)

func RunNode() {
	// client, err := ethclient.Dial("http://172.30.64.1:8545")
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws/v3/41357bf55fa74db5b4fed634eee82886")
	if err != nil {
		log.Println(err)
	}
	log.Println("we have a connection to ethereum")

	Client = client // we'll use this in the upcoming sections
	Auth = consultWithNode("b74629fb3ea5cbd5b1dabb170080ba752d16bc5151f107b60dc891025f62a0ce") // msg.sender private key
	Instance = connectToContract(ContractAddr)
	SenderAddr = common.HexToAddress("0x8548D0fA2250aE400Ec2aA31Dbd6294239FB97D0")
}

// ConsultWithNode 获取身份认证
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

func connectToContract(contractAddr string) *contract.Contract {
	address := common.HexToAddress(contractAddr)
	instances, err := contract.NewContract(address, Client)
	if err != nil {
		log.Fatal(err)
	}

	return instances
}

func UpdateNonce() (uint64, error) {
	return Client.PendingNonceAt(context.Background(), FromAddress)
}

func WatchMessage() {
	contractAddress := common.HexToAddress(ContractAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(contract.ContractABI)))
	if err != nil {
        log.Fatal(err)
    }

	receiveMap := map[string]interface{}{}

	for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
			err := contractAbi.UnpackIntoMap(receiveMap, "msgConn", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			eventID := big.NewInt(receiveMap["eventID"].(int64))

			if receiveMap["name"].(string) == "Rconn" {
				IntanceRconn, err := Instance.IndexRconn(nil, eventID)
				if err != nil {
					log.Fatalln(err)
				}

				_, err = Instance.ReCheckDDos(Auth, IntanceRconn, big.NewInt(100))
				if err != nil {
					log.Fatalln(err)
				}
			} else if receiveMap["name"].(string) == "Rddos" {
				_, err = Instance.InsertRddos(Auth, eventID)
				if err != nil {
					log.Fatalln(err)
				}
			} else {
				log.Fatalln("Invaild message!")
			}
        }
    }
}
