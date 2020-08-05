package main

import (
	"fmt"
	"strconv"
	_ "time"
)

// 在主线程(可以理解成进程)中，开启一个goroutine, 该协程每隔1秒输出 "hello,world"
// 在主线程中也每隔一秒输出"hello,golang", 输出10次后，退出程序
// 要求主线程和goroutine同时执行

//编写一个函数，每隔1秒输出 "hello,world"
func test() {
	for i := 1; i <= 1000; i++ {
		fmt.Println("tesst () hello,world " + strconv.Itoa(i))
		//time.Sleep(time.Second)
	}
	fmt.Println("协程程执行完毕")
}

func main() {

	go test() // 开启了一个协程

	for i := 1; i <= 100; i++ {
		fmt.Println(" main() hello,golang" + strconv.Itoa(i))
		//time.Sleep(time.Second)
	}
	fmt.Println("主线程执行完毕，直接结束")
}
