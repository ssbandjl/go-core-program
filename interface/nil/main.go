/*
Nil interface values
A nil interface value holds neither value nor concrete type.
直接对nil空指针调用方法将报错, 但是可以将空结构赋值给接口后调用
Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
*/

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I //i是空接口

	var t *T //t是结构体的指针
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
