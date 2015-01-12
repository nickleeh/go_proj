package main

import (
	"fmt"
	"net"
)

func main() {

	// Get local IP addresses
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}

	list, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	// If you want to know interface names too, use net.Interfaces() to get a list of interfaces first.

	for i, iface := range list {
		fmt.Printf("%d name=%s %v\n", i, iface.Name, iface)
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		for j, addr := range addrs {
			fmt.Printf(" %d %v\n", j, addr)
		}
	}

}
