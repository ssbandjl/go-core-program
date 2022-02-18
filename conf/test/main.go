package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/pelletier/go-toml"
)

var (
	monitorConfFile = "./monitor.conf"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	prometheus_ips := "192.168.1.111"
	config, err := toml.LoadFile(monitorConfFile)
	HandleError(err)
	// log.Printf("config:%v", config)
	path := []string{"prometheus", "ip"}
	if config.HasPath(path) {
		pos := config.GetPosition("prometheus.ip")
		// sed -i '28c ip="192.168.1.111"' ./monitor.conf 替换28行内容
		tmp := fmt.Sprintf("sed -i '%dc %s=\"%s\"' %s", pos.Line, "ip", prometheus_ips, monitorConfFile)
		log.Printf("cmd:%s", tmp)
		cmd := exec.Command("/bin/bash", "-c", tmp)
		_, err = cmd.CombinedOutput()
		HandleError(err)
	}
}
