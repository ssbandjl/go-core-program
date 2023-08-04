package main

import (
	"log"

	"github.com/shirou/gopsutil/net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Failed to query interfaces: %s", err.Error())
	}
	for _, i := range interfaces {
		netAddr := i.Addrs
		log.Printf("%v", netAddr)
		// for _, addr := range i.Addrs() {
		// 	log.Printf("i.Addrs:%v", addr.Addr)

		// }
	}
}
