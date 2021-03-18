package capture

import (
	"github.com/google/gopacket"
)

func RunPolicy(w gopacket.Packet) {
	SameSrcIp(w)
}

func SameSrcIp(w gopacket.Packet) {
	print(w)
}

func SynFlood() {}

func UdpFlood() {}

func Other() {}
