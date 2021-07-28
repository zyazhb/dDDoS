package main

import (
	"log"

	"main/node"
	"main/capture"
	"main/mitm"
)

func main() {
	// 连接到geth节点
	done := make(chan bool)
	err := node.RunNode()
	if err != nil {
		log.Fatalln("Failed to connect to geth node!")
		panic(err)
	} else {
		done <- true
	}
	<- done

	// 捕获流量 -> 插入链中
	go mitm.Run()
	go capture.CapturePacket()

	var wmh node.WatchMessageHandler
	wmh.WatchMessage()
}
