package capture

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var localip string
var ipLists = make(map[string]int)
var ipTime = make(map[string]int64)

func RunPolicy(packet gopacket.Packet) {
	SameSrcIp(packet)
}

func SameSrcIp(packet gopacket.Packet) {
	//获取本机 ip
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		for _, value := range addrs {
			if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					localip = ipnet.IP.String()
				}
			}
		}
	}

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		srcip := ip.SrcIP.String()
		if srcip != localip {
			_, ok := ipLists[srcip]
			if len(ipLists) > 20 {
				for k := range ipLists {
					if packet.Metadata().Timestamp.Unix()-ipTime[k] > 1 {
						delete(ipLists, k)
						delete(ipTime, k)
					}
				}
			}
			if !ok {
				ipLists[srcip] = 1
				ipTime[srcip] = packet.Metadata().Timestamp.Unix()
			} else {
				ipLists[srcip] += 1
				if ipLists[srcip] > 30 {
					time := packet.Metadata().Timestamp.Unix()
					// fmt.Printf("%v\n", time-ipTime[srcip])
					if time-ipTime[srcip] <= 1 {
						//Detected DDoS
						DetectedDDoS("Too many Same SrcIp")
						delete(ipLists, srcip)
					} else if time-ipTime[srcip] > 3 {
						delete(ipLists, srcip)
					}
				}
			}

		}
	}
}

func SynFlood() {}

func UdpFlood() {}

func Other() {}
