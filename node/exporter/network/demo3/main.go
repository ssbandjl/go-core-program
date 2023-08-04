package main

import (
	"encoding/json"
	"fmt"

	"github.com/prometheus/procfs/sysfs"
)

var (
	netClass map[string]sysfs.NetClassIface
)

type netClassCollector struct {
	fs sysfs.FS
}

func (c *netClassCollector) getNetClassInfo() (sysfs.NetClass, error) {
	netClass := sysfs.NetClass{}
	netDevices, err := c.fs.NetClassDevices()
	if err != nil {
		return netClass, err
	}
	for _, device := range netDevices {
		interfaceClass, err := c.fs.NetClassByIface(device)
		if err != nil {
			return netClass, err
		}
		netClass[device] = *interfaceClass
		// fmt.Println("netClass")
	}

	return netClass, nil
}

func (c *netClassCollector) Update() error {
	netClass, err := c.getNetClassInfo()
	if err != nil {
		return err
	}
	// fmt.Println(netClass)
	data, _ := json.Marshal(netClass)
	fmt.Println(string(data))
	return fmt.Errorf("could not get net class info: %w", err)
}

func main() {
	fs, err := sysfs.NewFS("/sys")
	if err != nil {
		fmt.Println(err.Error())
	}
	var netClassCollect = &netClassCollector{fs: fs}
	netClassCollect.Update()

}
