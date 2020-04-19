package main

/*
功能说明:适配器函数bridge用作统一处理接口,可以接受不同的参数进行反射，并调用对应的函数
执行方式:go test -v
*/

import (
	"fmt"
	"testing"
	"reflect"
)

func TestReflectFunc(t *testing.T){
	call1 := func(v1 int, v2 int){
		t.Log(v1, v2)
		fmt.Println(v1, v2)
	}

	call2 := func(v1 int, v2 int, s string){
		t.Log(v1, v2, s)
		fmt.Println(v1, v2, s)
	}

	var(
		function reflect.Value
		inValue []reflect.Value
		n int
	)

	//适配器函数：接受任意函数和任意个数参数
	bridge := func(call interface{}, args ...interface{}){
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i:=0; i<n; i++ {
			inValue[i] = reflect.ValueOf(args[i]) //切片赋值
		}
		function = reflect.ValueOf(call) //官方接口: func ValueOf(i interface{}) Value
		function.Call(inValue) //传入参数切片[]reflect.Value   官方接口: func (v Value) Call(in []Value) []Value
	}

	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")
}