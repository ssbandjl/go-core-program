package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var count uint32
	// define call back
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i { // spinning, 让count变量成为一个信号，它的值总是下一个可以调用打印函数的go函数的序号
				fn()
				atomic.AddUint32(&count, 1) // trigger函数会被多个 goroutine 并发地调用, 原子操作函数对被操作的数值的类型有约束
				break
			}
			// time.Sleep(time.Nanosecond)
		}
	}
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn) //call
		}(i)
	}
	trigger(10, func() {}) // when for end i=10, count=10, 主 goroutine 最后一个运行完毕
	// time.Sleep(time.Nanosecond * 100000)
}

/*
golang并发时，go程序需要有启动延时，需要让main函数休眠，才能让goroutine程序在main函数退出前有机会运行完毕
goroutinue程序的启动和for循环执行完毕是同时的，想上述代码，一般情况下gorountine开始执行时，for循环已经结束，因此i是10了。因此需要把i通过闭包封到goroutinue里去
goroutine是随机的，需要控制。
*/

/*
g fifo
https://studygolang.com/articles/26028
纵观count变量、trigger函数以及改造后的for语句和go函数，我要做的是，让count变量成为一个信号，它的值总是下一个可以调用打印函数的go函数的序号。

这个序号其实就是启用 goroutine 时，那个当次迭代的序号。也正因为如此，go函数实际的执行顺序才会与go语句的执行顺序完全一致。此外，这里的trigger函数实现了一种自旋（spinning）。除非发现条件已满足，否则它会不断地进行检查。

最后要说的是，因为我依然想让主 goroutine 最后一个运行完毕，所以还需要加一行代码。不过既然有了trigger函数，我就没有再使用通道。

调用trigger函数完全可以达到相同的效果。由于当所有我手动启用的 goroutine 都运行完毕之后，count的值一定会是10，所以我就把10作为了第一个参数值。又由于我并不想打印这个10，所以我把一个什么都不做的函数作为了第二个参数值
*/
