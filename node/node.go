package node

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func RunNode() {
	client, err := ethclient.Dial("http://172.30.64.1:8545")
	if err != nil {
		log.Println(err)
	}
	log.Println("we have a connection to ethereum")
	_ = client // we'll use this in the upcoming sections
}
