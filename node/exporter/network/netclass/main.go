package main

import (
	"fmt"
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/procfs/sysfs" // 软件包 sysfs 提供了从伪文件系统 sys 中检索系统和内核指标的功能。
)

var sysPath = "/sys"

type netClassCollector struct {
	fs                    sysfs.FS
	subsystem             string
	ignoredDevicesPattern *regexp.Regexp
	metricDescs           map[string]*prometheus.Desc
}

func main() {
	netclassIgnoredDevices := "^$"
	fs, err := sysfs.NewFS(sysPath)
	if err != nil {
		fmt.Errorf("failed to open sysfs: %w", err)
	}
	pattern := regexp.MustCompile(netclassIgnoredDevices)

	c := netClassCollector{
		fs:                    fs,
		subsystem:             "network",
		ignoredDevicesPattern: pattern,
		metricDescs:           map[string]*prometheus.Desc{},
	}
	fmt.Println(c.getNetClassInfo())
}

func (c *netClassCollector) getNetClassInfo() (sysfs.NetClass, error) {
	netClass := sysfs.NetClass{}
	netDevices, err := c.fs.NetClassDevices()
	if err != nil {
		return netClass, err
	}

	for _, device := range netDevices {
		if c.ignoredDevicesPattern.MatchString(device) {
			continue
		}
		interfaceClass, err := c.fs.NetClassByIface(device)
		if err != nil {
			return netClass, err
		}
		netClass[device] = *interfaceClass
		// fmt.Println("netClass")
	}

	return netClass, nil
}
