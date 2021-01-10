package main

import (
	"fmt"
	"reflect"
	"time"
)

//发送同一个请求到多个微服务节点，只要任意一个微服务节点返回结果，就算成功
// 在 chan 数量比较多的情况下，递归并不是一个很好的解决方式，根据这一讲最开始介绍的反射的方法，我们也可以实现 Or-Done 模式
//这是递归和反射两种方法实现 Or-Done 模式的代码。反射方式避免了深层递归的情况，可以处理有大量 chan 的情况。其实最笨的一种方法就是为每一个 Channel 启动一个 goroutine，不过这会启动非常多的 goroutine，太多的 goroutine 会影响性能，所以不太常用。你只要知道这种用法就行了，不用重点掌
func or(channels ...<-chan interface{}) <-chan interface{} {
	//特殊情况，只有0个或者1个
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		// 利用反射构建SelectCase
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		// 随机选择一个可用的case
		reflect.Select(cases)
	}()

	return orDone
}

// 返回只读chan
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()

	<-or(
		sig(10*time.Second),
		sig(20*time.Second),
		sig(30*time.Second),
		sig(40*time.Second),
		sig(50*time.Second),
		sig(01*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}
