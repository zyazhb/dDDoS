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
		// fmt.Println("[*]--IPv4 layer detected.--")
		fmt.Printf("[-] detect localip %s\n", localip)
		ip, _ := ipLayer.(*layers.IPv4)
		// fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		// fmt.Printf("[+] %s ---------->%s\n", ip.SrcIP, ip.DstIP)
		if ip.SrcIP.String() != localip {
			fmt.Println("[-]--IPv4 layer detected.--")
			fmt.Printf("[-]From %s to %s\n", ip.SrcIP, ip.DstIP)
			_, ok := ipLists[ip.SrcIP.String()]
			fmt.Printf("-------------ok:%v------------\n", ok)
			if len(ipLists) > 20 {
				min := 10
				var key string
				for k, _ := range ipLists {
					fmt.Printf("k:%s\n", k)
					if ipLists[k] < min {
						min = ipLists[k]
						key = k
					}
				}
				delete(ipLists, key)
				delete(ipTime, key)
			}
			if !ok {
				ipLists[ip.SrcIP.String()] = 1
				ipTime[ip.SrcIP.String()] = packet.Metadata().Timestamp.Unix()
			} else {
				ipLists[ip.SrcIP.String()] += 1
				if ipLists[ip.SrcIP.String()] > 100 {
					time := packet.Metadata().Timestamp.Unix()
					fmt.Printf("%v\n", time-ipTime[ip.SrcIP.String()])
					if time-ipTime[ip.SrcIP.String()] <= 10 {
						fmt.Println("[+] detect dddos attack!")
					}
				}
			}
		}
	}
}

func SynFlood() {}

func UdpFlood() {}

func Other() {}
