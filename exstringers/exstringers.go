package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var res = ""
	for i := 0; i < 3; i++ {
		res += ip[i] + "."
	}
	return res + ip[3]
}

func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v \n", name, ip)
	}
}
