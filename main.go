package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	ip := getHostIp()
	fmt.Println(ip)

}

// getHostIp1 借助 net.InterfaceAddrs 方法（多网卡时，不推荐）
func getHostIp1() string {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("get current host ip err: ", err)
		return ""
	}
	var ip string
	for _, address := range addrList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				break
			}
		}
	}
	return ip
}

// getHostIp (推荐)使用 udp 不需要关注是否送达，只需要对应的 ip 和 port 正确，即可获取到 IP 地址
func getHostIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println("get current host ip err: ", err)
		return ""
	}
	addr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(addr.String(), ":")[0]
	return ip
}
