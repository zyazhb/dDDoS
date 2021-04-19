package capture

import (
	"log"
	"math/big"
	"time"

	"main/node"
)

type UpInfo struct {
	Now time.Time
	Reason string
}

var (
	eventID = 1700064

	UpChan = make(chan UpInfo, 1024)
)

func DetectedDDoS(reason, srcip string, speed int) {
	info := UpInfo{
		Now: time.Now(),
		Reason: reason,
	}

	UpChan <- info

	log.Println("--------------------------")
	log.Println("[+] detected DDoS attack!")
	log.Println("Time:", info.Now)
	log.Println("Reason:", info.Reason)
	log.Println("--------------------------")
	//TO-DO 区块链交互

	rconn := []string{srcip, info.Now.String(), info.Reason}
	eventID += 1

	_, err := node.Instance.InsertRconn(node.Auth, big.NewInt(int64(eventID)), node.SenderAddr, big.NewInt(int64(speed)), rconn)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Insert into chain successfully!")
	}

	newNonce, err := node.UpdateNonce()
	node.Auth.Nonce = big.NewInt(int64(newNonce))
}
