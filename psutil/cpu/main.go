package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	//CPU utilization
	// for {
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("cpu percent:%v\n", percent)
	// }
}
func main() {
	getCpuInfo()
}
