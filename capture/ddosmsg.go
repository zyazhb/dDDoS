package capture

import (
	"log"
	"time"
)

func DetectedDDoS(Reason string) {
	log.Println("--------------------------")
	log.Println("[+] detected DDoS attack!")
	log.Println("Time:", time.Now())
	log.Println("Reason:", Reason)
	log.Println("--------------------------")
	//TO-DO 区块链交互
}
