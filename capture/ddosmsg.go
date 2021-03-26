package capture

import (
	"log"
	"math/big"
	"time"

	"main/node"
)

const (
	eventID = iota
)

func DetectedDDoS(Reason, srcip string, speed int) {
	now := time.Now()

	log.Println("--------------------------")
	log.Println("[+] detected DDoS attack!")
	log.Println("Time:", now)
	log.Println("Reason:", Reason)
	log.Println("--------------------------")
	//TO-DO 区块链交互

	rconn := []string{srcip, now.String()}

	_, err := node.Instance.InsertRconn(node.Auth, big.NewInt(eventID), node.SenderAddr, big.NewInt(int64(speed)), rconn)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Insert into chain successfully!")
	}

	newNonce, _ := node.UpdateNonce()

	node.Auth.Nonce = big.NewInt(int64(newNonce))
}
