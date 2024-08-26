package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	result := ""
	for _, n := range ip {
		result = result + strconv.Itoa(int(n)) + "."
	}
	result = result[:len(result) - 1]
	return result
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}