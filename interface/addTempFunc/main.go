package main

import "log"

type Hello interface {
	Print()
}

type HelloWorld struct {
	S string
}

func (self HelloWorld) Print() {
	log.Printf(self.S)
}

func main() {
	var h Hello
	para := "我是参数"
	// 零时给接口增加方法, 接口转换
	h.(interface{ TempFunc(para string) }).TempFunc(para)

}
