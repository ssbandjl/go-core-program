package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/mem"
)

func getMetric() {
	memory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get metric failed, err:%v", err)
	}
	log.Printf("Available:%v", memory.Available)
	log.Printf("Used:%v", memory.Used)
	log.Printf("Free:%v", memory.Free)
	log.Printf("Buffers:%v", memory.Buffers)
	log.Printf("Cached:%v", memory.Cached)
	log.Printf("Total:%v", memory.Total)
	log.Printf("Slab:%v", memory.Slab)
}
func main() {
	getMetric()
}
