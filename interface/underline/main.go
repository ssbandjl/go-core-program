package main

type I interface {
	Sing()
}

type T struct {
}

// T自动实现 fun(t *T) Sing(){}方法
func (t T) Sing() {
}

type T2 struct {
}

func (t *T2) Sing() {
}

// Compiler passes
var _ I = T{}

// Compiler passes
var _ I = &T{}

// Compilation failure
var _ I = T2{}

// Compiler passes
var _ I = &T2{}

//接口在底层有两个成员,类型和值, 零值(nil,nil), 指针的零值是nil
var _ I = (*T2)(nil)

func main() {
	//在底层，interface 作为两个成员来实现：一个类型和一个值 (type, value)。value 被称为接口的动态值，它是一个任意的具体值，而该 type 则为该值的类型。对于 int 值 3， 一个接口值示意性地包含 (int, 3)。
	//接口的零值是 (nil, nil)。换句话说。当一个接口和 nil 比较时，只有该接口内部的值和类型都是 nil 时它才等于 nil。比如我们在一个接口值 i 中存储一个 *int 类型的指针 p，则接口 i 的内部类型将为 *int。无论指针 p 是否为 nil，i != nil 将永远返回 true。
	var i interface{}
	var p *int = nil //*int类型不为nil
	i = p
	println(i != nil) //true
	//指针的零值是 nil。因此，可以先将一个接口值转为一个指针类型，然后再与 nil 比较，从而判断接口内部的值是否为 nil。举例
	println(i.(*int) != nil) //false
}
