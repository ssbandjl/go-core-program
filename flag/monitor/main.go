package main

import (
	"flag"
	"log"

	"github.com/pelletier/go-toml"
)

var (
	// ConfigFile monitor config file
	ConfigFile = flag.String("c", "/etc/neonsan/monitor.conf", "Monitor config file path.")
)

// Init init monitor parameter from config file
func Init(confFile string) {
	common.Setlog()
	if err := monitor.LoadConfigFile(confFile); err != nil {
		log.Fatalf("Failed to load config file: %s, %s", confFile, err.Error())
	}

	config, err := toml.LoadFile(confFile)
	if err != nil {
		log.Fatalf("Config file error:%s\n", err.Error())
	}

	monitorStore = config.GetDefault("monitor.monitor_store", false).(bool)
}

func main() {
	flag.Parse()
	log.Printf("monitor.conf:%+v", ConfigFile)
}
