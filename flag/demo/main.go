package main

import (
	"flag"
	"log"
)

// go run main.go -h
func main() {
	var name string
	flag.StringVar(&name, "name", "Go flag 默认信息, 长参数", "长参数帮助信息")
	flag.StringVar(&name, "n", "Go flag 默认信息, 短参数", "短参数帮助信息")
	flag.Parse()
	log.Printf("name:%s", name)
}
