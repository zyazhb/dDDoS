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

	UpChan = make(chan UpInfo, 1024)
)

func DetectedDDoS(reason, srcip string, speed int) {
	info := UpInfo{
		Now: time.Now(),
		Reason: reason,
	}

	log.Println("--------------------------")
	log.Println("[+] detected DDoS attack!")
	log.Println("Time:", info.Now)
	log.Println("Reason:", info.Reason)
	log.Println("--------------------------")

	node.SendMessage(info.Reason)
}
