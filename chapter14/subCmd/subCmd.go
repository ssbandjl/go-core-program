package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	/*
		期望效果:
		./tools -h
		./tools getServiceId -s sssss
		./tools updateFlume -rollback
		./tools updateFlume -update -f /home/icksys/config/flume/config.yml -serviceName bc-data
	*/
	helpText := `
DevOps工具箱@2020

Usage:  ./tools COMMAND [args...]

Options:
	--config string      Location of client config files (default "/home/icksys/.docker")
	-D, --debug              Enable debug mode
	-H, --host list          Daemon socket(s) to connect to
	-l, --log-level string   Set the logging level ("debug"|"info"|"warn"|"error"|"fatal") (default "info")
	-v, --version            Print version information and quit

COMMAND:
	getServiceId     从本地consul获取服务ID，可以用于注销服务
	addService2Flume	添加服务日志收集

RUN:
./tools -h 查看帮助
./tools getServiceId -h  查看子命令帮助信息
	`

	//注册子命令
	getServiceIdCmd := flag.NewFlagSet("getServiceId", flag.ExitOnError)

	//子命令选项
	service := getServiceIdCmd.String("s", "", "service 服务名,默认为空")
	api := getServiceIdCmd.String("a", "", "api地址，默认为空")

	// 对于不同的子命令，我们可以定义不同的 flag
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) == 2 && os.Args[1] == "-h" {
		fmt.Println("帮助信息")
		fmt.Println(helpText)
		os.Exit(1)
	}

	// 期望前面定义的子命令作为第一个参数传入。
	if len(os.Args) < 2 {
		fmt.Println(helpText)
		os.Exit(1)
	}

	// 检查哪一个子命令被调用了。
	switch os.Args[1] {

	// 每个子命令，都会解析自己的 flag 并允许它访问后续的参数。
	case "getServiceId":
		getServiceIdCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'getServiceIdCmd'")
		fmt.Println("  service:", *service)
		fmt.Println("  api:", *api)
		fmt.Println("  tail:", getServiceIdCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println(helpText)
		os.Exit(1)
	}
}
