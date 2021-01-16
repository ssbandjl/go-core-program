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
	// 零时给接口增加方法, 接口转换
	h.(interface{ TempFunc() }).TempFunc()

}
