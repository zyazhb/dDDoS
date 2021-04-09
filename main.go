package main

import (
	"main/capture"
	"main/mitm"
	"main/node"
	"main/web"
	"time"
)

func main() {
	go node.RunNode()

	time.Sleep(5 * time.Second)

	go mitm.Run()
	go capture.CapturePacket()

	// go node.WatchMessage()
	web.RunWeb()
}
