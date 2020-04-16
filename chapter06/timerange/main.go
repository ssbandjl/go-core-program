package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	//定义几个变量，用于接收命令行的参数值
	var start string
	var end string

	//&user 就是接收用户命令行中输入的 -u 后面的参数值
	//"u" ,就是 -u 指定参数
	//"" , 默认值
	//"用户名,默认为空" 说明
	flag.StringVar(&start, "s", "2006-01-02 15:04:05", "开始时间")
	flag.StringVar(&end, "end", "2006-01-02 15:04:05", "结束时间")
	//这里有一个非常重要的操作,转换， 必须调用该方法
	flag.Parse()

	//输出结果
	fmt.Printf("start=%v end=%v", start, end)

	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串 fmt.Printf(now.Format("2006-01-02 15:04:05"))
	startS, _ := time.Parse("2006-01-02 15:04:05", start)
	// entS, _ := (time.Parse("2006-01-02 15:04:05", end)).Unix()
	fmt.Println(startS.unix())
}
