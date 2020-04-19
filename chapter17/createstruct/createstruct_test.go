package main

/*
功能说明: 使用反射创建并操作结构体，本实例理解有难度，建议画内存图便于理解
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
		st reflect.Type
		elem reflect.Value
	)

	st = reflect.TypeOf(model)
	t.Log("relect.Typeof", st.Kind().String()) //relect.Typeof ptr

	st = st.Elem() //st指向的类型 Elem returns the value that the interface v contains or that the pointer v points to. It panics if v's Kind is not Interface or Ptr. It returns the zero Value if v is nil
	t.Log("relect.TypeOf.Elem", st.Kind().String()) //relect.TypeOf.Elem struct

	elem=reflect.New(st)
	t.Log("reflect.New", elem.Kind().String()) //reflect.New ptr
	t.Log("relect.New.Elem", elem.Elem().Kind().String())

	//此操作使model和elem指向同一个地址空间
	model=elem.Interface().(*user)   //先转成空接口，再用类型断言

	elem = elem.Elem()  //elem直接指向user空间
	elem.FieldByName("UserId").SetString("12345678")
	elem.FieldByName("Name").SetString("nickname")
	t.Log("model model.Name", model, model.Name)
}

//测试日志: 
// PS D:\go\src\go_code\chapter17\createstruct> go test -v
// === RUN   TestReflectStruct
// --- PASS: TestReflectStruct (0.00s)
//         createstruct_test.go:26: relect.Typeof ptr
//         createstruct_test.go:28: relect.TypeOf.Elem struct
//         createstruct_test.go:30: reflect.New ptr
//         createstruct_test.go:31: relect.New.Elem struct
//         createstruct_test.go:36: model model.Name &{12345678 nickname} nickname
// PASS
// ok      go_code/chapter17/createstruct  0.047s
// PS D:\go\src\go_code\chapter17\createstruct>