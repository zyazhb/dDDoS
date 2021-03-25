package capture

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var ipLists = make(map[string]int)
var ipTime = make(map[string]int64)

func RunPolicy(packet gopacket.Packet) {
	SameSrcIp(packet, 1, 10, 3)
}

// WT最小检测时间 CON最大连接数 CT最短清空时间
func SameSrcIp(packet gopacket.Packet, WT int64, CON int, CT int64) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		srcip := ip.SrcIP.String()
		if srcip != localip && srcip != "127.0.0.1" {
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
				if ipLists[srcip] > CON {
					time := packet.Metadata().Timestamp.Unix()
					if time-ipTime[srcip] <= WT {
						//Detected DDoS
						reason := "Too many Same SrcIp: " + srcip
						DetectedDDoS(reason)
						delete(ipLists, srcip)
					} else if time-ipTime[srcip] > CT {
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
