package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("运行时中获取Go版本:%s\n", runtime.Version())

	pc, file, line, ok := runtime.Caller(0)
	if ok {
		f := runtime.FuncForPC(pc)
		fmt.Printf("文件:%s: 行数:%d 方法名:%s\n", file, line, f.Name())
	}
}
