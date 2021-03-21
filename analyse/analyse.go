package main

// Use tcpdump to create a test file
// tcpdump -w test.pcap
// or use the example above for writing pcap files
import (
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket/layers"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	pcapFile string = "test.pcap"
	handle   *pcap.Handle
	err      error
	localip  string
)

func main() {
	// 打开pacp文件获得一个句柄
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	// 从句柄中读取数据包源
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	// 循环遍历数据包源内的数据包
	ipLists := make(map[string]int)
	//获取本机 ip
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localip = ipnet.IP.String()
			}
		}
	}
	ipTime := make(map[string]int64)
	for packet := range packetSource.Packets() {
		// now := time.Now()
		// 读取数据包中的ip层
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			// fmt.Printf("[+] %s ---------->%s\n", ip.SrcIP, ip.DstIP)
			if ip.SrcIP.String() != localip {
				_, ok := ipLists[ip.SrcIP.String()]
				if !ok {
					ipLists[ip.SrcIP.String()] = 1
					ipTime[ip.SrcIP.String()] = packet.Metadata().Timestamp.Unix()
				} else {
					ipLists[ip.SrcIP.String()] += 1
					if ipLists[ip.SrcIP.String()] > 5 {
						time := packet.Metadata().Timestamp.Unix()
						if time-ipTime[ip.SrcIP.String()] >= 60 {
							fmt.Println("检测到ddos攻击")
						}
					}
				}
			}
		}
	}
	for k, v := range ipLists {
		fmt.Printf("%s<------>%d", k, v)
	}
	for k, v := range ipTime {
		fmt.Printf("%s<------>%d", k, v)
	}
}
