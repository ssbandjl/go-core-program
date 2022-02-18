package main

import "C"

// 完整CGO, 执行go build编译和链接会启动gcc编译器
func main() {
	println("hello cgo")
}
