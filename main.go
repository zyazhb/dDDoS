package main

import (
	"main/capture"
	"main/mitm"
)

func main() {
	go mitm.Run()
	capture.CapturePacket()

}
