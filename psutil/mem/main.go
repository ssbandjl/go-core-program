package main

import (
	"fmt"
	"log"
	"math"

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
	// used=MemTotal-(MemFree+Buffers+Cached+Slab) ，把MemFree+Buffers+Cached做为available的数据
	log.Printf("used:%v", memory.Total-(memory.Free+memory.Buffers+memory.Cached+memory.Slab))
	log.Printf("available:%v", memory.Free+memory.Cached+memory.Buffers)

	memStat, err := mem.VirtualMemory()

	used := float64(memStat.Total) - float64(memStat.Free) - float64(memStat.Buffers) - float64(memStat.Cached)
	awailable := float64(memStat.Total) - used

	value := (float64(1.0) - awailable/float64(memStat.Total)) * 100
	value = math.Floor(value + 0.5)
	log.Printf("percent:%v, used:%v, available:%v", value, used, awailable)
}
func main() {
	getMetric()
}
