package main

import (
	"fmt"
	"net"
)

func main() {
	validIPV4 := "10.40.210.253"

	checkIPAddress(validIPV4)

	invalidIPV4 := "1000.40.210"
	checkIPAddress(invalidIPV4)

}

func checkIPAddress(ip string) {
	if net.ParseIP(ip) == nil {
		fmt.Printf("IP Address: %s - Invalid\n", ip)
	} else {
		fmt.Printf("IP Address: %s - Valid\n", ip)
	}
}
