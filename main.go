package main

import (
	"main/capture"
	"main/mitm"
	"main/node"
	"main/web"
)

func main() {
	go mitm.Run()
	go capture.CapturePacket()
	go node.RunNode()
	go node.WatchMessage()
	web.RunWeb()
}

