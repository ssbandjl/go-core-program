package main

import (
	"log"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2) //2的平方根
	v := Vertex{3, 4}         //3*3+4*4 开根号
	log.Printf("f:%f, v:%f", f, v)
	a = f // a MyFloat implements Abser  将实现了该接口的结构体赋值给该空接口
	log.Printf("a的绝对值:%f", a.Abs())

	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	log.Printf("a的绝对值:%f", a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
