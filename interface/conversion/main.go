package main

import (
	"fmt"
)

// Getter sets a string
type Getter interface {
	Get() string
}

// Setter sets a string
type Setter interface {
	Set(string)
}

// A is a demo struct
type A struct {
	s string
}

func (a A) Get() string {
	return a.s
}
func (a *A) Set(s string) {
	a.s = s
}

// NewGetter returns new Getter interface
func NewGetter() Getter {
	return A{s: "Hello World!"}
}

func main() {
	a := NewGetter()
	fmt.Println(a.Get()) // Prints Hello World!

	// Following does not work and it is ok.
	// a.(Setter).Set("Hello Earth!") // panic: interface conversion: main.A is not main.Setter: missing method Set

	// Following does not work and it is ok, but weird.
	// (&a).(Setter).Set("Hello Earth!") // build failed: invalid type assertion: (&a).(Setter) (non-interface type *Getter on left)

	// Following does not work, but wait, what works than?
	// var i interface{} = &a
	// i.(Setter).Set("Hello Earth!") // panic: interface conversion: *main.Getter is not main.Setter: missing method Set

	// Following works, but how is it different from the previous example?
	var j interface{} = &A{
		s: "Hello World!",
	}

	// 将j断言为Setter接口, 再调用Set方法
	j.(Setter).Set("Hello Earth!")
	fmt.Println(j)

	// 一个对象实现了一个接口, 利用这个接口调用该对象的其他(非实现接口的方法)方法, 需要临时将该接口转换为该对象其他的方法实现的接口(临时或已存在的接口)来适配(调用该对象的其他方法)
	j.(interface{ Set(s string) }).Set("你好")
	fmt.Println(j)

}
