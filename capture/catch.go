package capture

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

var (
	deviceName  string = "eth0"
	snapshotLen int32  = 1024
	promiscuous bool   = false
	err         error
	timeout     time.Duration = -1 * time.Second
	handle      *pcap.Handle
	packetCount int = 0
)

func CapturePacket() {
	// Open output pcap file and write header
	f, _ := os.Create("test.pcap")
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(snapshotLen), layers.LinkTypeEthernet)
	defer f.Close()
	// Open the device for capturing
	handle, err = pcap.OpenLive(deviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		fmt.Printf("Error opening device %s: %v", deviceName, err)
		os.Exit(1)
	}
	defer handle.Close()

	// Set filter
	filter := "tcp and port 8080"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Only capturing TCP port 8080 packets.")

	// Start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		// fmt.Println(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++

		// Only capture 100 and then stop
		if packetCount > 100 {
			break
		}
	}
}
