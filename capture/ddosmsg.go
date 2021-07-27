package capture

import (
	"log"
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
}
