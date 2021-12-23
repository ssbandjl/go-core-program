package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond := sync.Cond{L: &mutex}
	condition := false
	go func() {
		time.Sleep(1 * time.Second)
		cond.L.Lock()
		fmt.Println("子协程已经锁定")
		fmt.Println("子协程更改条件, 并发送通知")
		condition = true
		// cond.Signal() //发送通知给单个等待者, 建议在锁定条件下调用
		cond.Broadcast() //发送通知给所有等待者, 唤醒所有等待条件的协程, 建议在锁定条件下调用
		fmt.Println("子协程继续执行")
		time.Sleep(5 * time.Second)
		fmt.Println("子协程解锁")
		cond.L.Unlock()
	}()
	cond.L.Lock()
	fmt.Println("main, 已锁定")
	if !condition {
		fmt.Println("main, 即将等待")
	}
	cond.Wait() // 自行解锁c.L, 并阻塞当前协程, 待线程恢复执行时, Wait()方法在返回前锁定c.L
	fmt.Println("main, 被唤醒")
	fmt.Println("main, 继续")
	fmt.Println("main, 解锁")
	cond.L.Unlock()
}
