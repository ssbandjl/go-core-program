package main

import (
	"flag"
	"log"
)

var name string

//解析子命令
func main() {
	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError) //创建一个新的空命令集
	goCmd.StringVar(&name, "name", "Go语言", "帮助信息")
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "PHP语言", "帮助信息")

	args := flag.Args()
	log.Printf("参数args:%s", args)
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "php":
		_ = phpCmd.Parse(args[1:])
	}
	log.Printf("name:%s", name)
}
