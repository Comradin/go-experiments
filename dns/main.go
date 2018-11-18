package main

import (
	"fmt"
	"net"
)

func main() {

	ipAddr := "127.0.0.1"
	hostName := "localhost"
	dnsName, _ := net.LookupAddr(ipAddr)
	fmt.Println(dnsName)
	ip, _ := net.LookupIP(hostName)
	fmt.Println(ip)
}
