package main

import (
	"fmt"
	"log"
	"reflect"
)

type TestEchoACK struct {
	Msg   string
	Value int32
}

func main() {
	// var a * TesEchoACK {} = nil
	log.Printf("获取空类型:%+v", (*TestEchoACK)(nil))
	myType := reflect.TypeOf((*TestEchoACK)(nil)).Elem() //类型为TestEchoACK零值的反射类型
	log.Printf("myType:%+v", myType)

	/*
		a：=(* interface {})(nil)等于var a * interface {} = nil.
		但是var b interface {} =(* interface {})(nil),mean b是type interface {},而interface {}变量只有nil,当它的类型和值都是nil时,显然type * interface {}不是nil .
	*/
	a := (*interface{})(nil)
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a))
	var b interface{} = (*interface{})(nil)
	fmt.Println(reflect.TypeOf(b), reflect.ValueOf(b))
	fmt.Println(a == nil, b == nil)
}
