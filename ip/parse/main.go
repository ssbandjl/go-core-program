package main

import (
	"log"
	"net"
)

var MinNumberLeadingBits = 16

func main() {
	monitorIp := "192.168.1.100/15"
	if netIp, net, err := net.ParseCIDR(monitorIp); err == nil {
		log.Printf("netIp:%v, net:%v", netIp, net)
		if n, _ := net.Mask.Size(); n < MinNumberLeadingBits {
			log.Printf("掩码错误")
		}
	} else {
		log.Printf(err.Error())
	}
}
