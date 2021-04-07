package capture

import (
	// "fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var ipLists = make(map[string]int)
var ipTime = make(map[string]int64)
var synip = make(map[string]int)
var ipseq = make(map[string]int)

func RunPolicy(packet gopacket.Packet) {
	SameSrcIp(packet, 1, 10, 3)
	SynFlood(packet)
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
						DetectedDDoS(reason, srcip, ipLists[srcip])
						delete(ipLists, srcip)
					} else if time-ipTime[srcip] > CT {
						delete(ipLists, srcip)
					}
				}
			}

		}
	}
}

func SynFlood(packet gopacket.Packet) {
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	ip, ipok := ipLayer.(*layers.IPv4)
	tcp, tcpok := tcpLayer.(*layers.TCP)
	srcip := ip.SrcIP.String()
	if ipok && tcpok && srcip != localip && srcip != "127.0.0.1" {
		// fmt.Println("[-]--TCP layer detected.--")
		// fmt.Printf("src:%v seq:%v SYn:%v ack:%v to dst:%v port:%v seq:%v ack:%v\n", ip.SrcIP, tcp.Seq, tcp.SYN, tcp.ACK, ip.DstIP, tcp.DstPort, tcp.Seq, tcp.Ack)
		_, ok := synip[srcip]
		if !ok {
			synip[srcip] = 0
			if tcp.SYN && tcp.ACK {
				ipseq[srcip] = int(tcp.Seq)
			}
		} else {
			if ipseq[srcip]+1 == int(tcp.Seq) {
				synip[srcip] = 1
				delete(synip, srcip)
				delete(ipseq, srcip)
			}
		}
		if len(synip) > 50 {
			reason := "Syn flood: " + srcip
			DetectedDDoS(reason, srcip, len(synip))
		}
	}
}

func UdpFlood() {}

func Other() {}
