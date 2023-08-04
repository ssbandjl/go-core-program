package main

import (
	"encoding/json"
	"fmt"

	//	"github.com/prometheus/procfs"
	"github.com/prometheus/procfs/sysfs"
)

var (
	NetClass *sysfs.NetClass
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
	}

	return netClass, nil
}

func getNetClassInfo() error {
	NetClass = make(&sysfs.NetClass)
	fs, err := sysfs.NewFS("/sys")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	var netClassCollect = &netClassCollector{fs: fs}
	NetClass, _ := netClassCollect.getNetClassInfo()
	data, err := json.Marshal(NetClass)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func getNetClassOperStateByInterface(interfaceName string) string {
	if netClassIface, ok := NetClass[interfaceName]; ok {
		return netClassIface.OperState
	} else {
		return "error"
	}
}

// go run -mod=vendor main.go
func main() {
	// fs, err := procfs.NewFS("/proc")
	// fs, err: = sysfs.NewFS("/proc")
	fs, err := sysfs.NewFS("/sys")
	if err != nil {
		fmt.Println(err.Error())
	}
	// stats, err := fs.Stat()
	//	fmt.Println(stats)

	netClass := sysfs.NetClass{}
	fmt.Println(netClass)
	netDevices, err := fs.NetClassDevices()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, device := range netDevices {

		interfaceClass, err := fs.NetClassByIface(device)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(interfaceClass.OperState)
	}
}
