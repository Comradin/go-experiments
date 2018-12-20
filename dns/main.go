package main

import (
	"fmt"
	"net"
)

func main() {

	ipAddr := "127.0.0.1"
	hostName := "localhost"
	dnsName, _ := net.LookupAddr(ipAddr)
	ip, _ := net.LookupIP(hostName)
	fmt.Println("Forward lookup of 'localhost' resolves to: ", ip)
	fmt.Println("Reverse lookup of '127.0.0.1' resolves to: ", dnsName)
}
