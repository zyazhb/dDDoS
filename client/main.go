package main

import (
	"log"

	"main/node"
	"main/capture"
	"main/mitm"
)

func main() {
	// 连接到geth节点
	err := node.RunNode()
	if err != nil {
		log.Fatalln("Failed to connect to geth node!")
		panic(err)
	}

	// 捕获流量 -> 插入链中
	go mitm.Run()
	go capture.CapturePacket()

	node.WatchMessage()
}
