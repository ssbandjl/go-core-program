package main

//这些实现确实比上一节中的方法礼貌一些，但是它们不能完全有效地避免数据竞争。 目前的Go白皮书并不保证发生在一个通道上的并发关闭操作和发送操纵不会产生数据竞争。
//如果一个SafeClose函数和同一个通道上的发送操作同时运行，则数据竞争可能发生（虽然这样的数据竞争一般并不会带来什么危害）。

import (
	"log"
	"sync"
)

type T struct {
	t string
}

type MyChannel struct {
	C      chan T
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	log.Printf("安全关闭通道")
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
}

func (mc *MyChannel) IsClosed() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}
func main() {
	c := NewMyChannel()
	c.SafeClose()
	c.SafeClose()
}
