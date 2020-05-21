package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {

	//定义几个变量，用于接收命令行的参数值
	var service string

	//&user 就是接收用户命令行中输入的 -u 后面的参数值
	//"u" ,就是 -u 指定参数
	//"" , 默认值
	//"用户名,默认为空" 说明
	flag.StringVar(&service, "s", "", "默认服务名为空")
	//这里有一个非常重要的操作,转换， 必须调用该方法
	flag.Parse()
	//输出结果
	fmt.Printf("服务名:%v", service)
	resp, err := http.Get("http://127.0.0.1:8500/v1/agent/services")
	fmt.Printf("返回值: %v, 错误信息: %v", resp, err)

}
