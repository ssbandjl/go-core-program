package main

//static const char* cs = "hello";
import "C"
import "../cgo_helper/"

func main() {
	cgo_helper.PrintCString(C.cs) // C.cs是当前main包引入的虚拟C包下的*cahr类型*main.C.char, 不能传递给*cgo_helper.C.char
}
