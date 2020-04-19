package main

/*
功能说明: 使用反射操作任意结构体
执行方式: go test -v
*/

import (
	"testing"
	"reflect"
)

type user struct {
	UserId string
	Name string
}

func TestReflectStruct(t *testing.T){
	var (
		model *user
		sv reflect.Value
	)

	model = &user{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.ValueOf", sv.Kind().String()) //reflect.ValueOf ptr
	
	//通过地址找到对应的变量值，并做修改
	sv = sv.Elem()
	t.Log("reflect.ValueOf.Elem", sv.Kind().String()) //reflect.ValueOf.Elem struct
	sv.FieldByName("UserId").SetString("12345678")
	sv.FieldByName("Name").SetString("nickname")
	t.Log("model", model) //model &{12345678 nickname}
}