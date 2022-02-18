package main

import (
	"log"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Failed to query interfaces: %s", err.Error())
	}
	for _, i := range interfaces {
		netAddr, _ := i.Addrs()
		log.Printf("%v", netAddr)
		// for _, addr := range i.Addrs() {
		// 	log.Printf("i.Addrs:%v", addr.Addr)

		// }
	}
}
