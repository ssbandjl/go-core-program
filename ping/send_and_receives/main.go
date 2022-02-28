package main

import (
	"log"

	"github.com/go-ping/ping"
)

func main() {
	// pinger, err := ping.NewPinger("www.google.com")
	pinger, err := ping.NewPinger("baidu.com")
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	log.Printf("ping run")

	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	log.Printf("ping done")
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	log.Printf("stats:%v", stats)
}
