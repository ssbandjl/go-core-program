package main

import (
	"flag"
	"fmt"
	"os"
	"getserviceid"
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
	service := getServiceIdCmd.String("s", "", "service 默认服务名为空")
	consulApiPrefix := getServiceIdCmd.String("a", "http://127.0.0.1:8500/", "默认API前缀")
	


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
		fmt.Println("子命令 'getServiceIdCmd'")
		fmt.Println("  service:", *service)
		fmt.Println("  api:", *consulApiPrefix)
		fmt.Println("  tail:", getServiceIdCmd.Args())
		getserviceid.getserviceid(*service, *consulApiPrefix)
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
