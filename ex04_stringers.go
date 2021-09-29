package main

import (
	"fmt"
	"strings"
	"strconv"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	address := make([] string, 4)

	for i:= 0; i < len(ip); i++ {
		address[i] = strconv.Itoa(int(ip[i]))
	}
	return strings.Join(address, ".")
}


func main() {
	hosts := map[string]IPAddr {
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
