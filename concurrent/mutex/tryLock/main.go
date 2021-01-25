package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// 复制Mutex定义的常量
const (
	mutexLocked      = 1 << iota // 1 加锁标识位置 , iota从0开始, 将1往左, 移位运算0次, 所以结果是指数型递增
	mutexWoken                   // 2 唤醒标识位置  移位1次
	mutexStarving                // 4 锁饥饿标识位置
	mutexWaiterShift = iota      // 3 标识waiter的起始bit位置
)

// 扩展一个Mutex结构
type Mutex struct {
	sync.Mutex
}

// 尝试获取锁
func (m *Mutex) TryLock() bool {
	// 如果能成功抢到锁, fast path 捷径
	log.Printf("当前状态, mutexLocked:%+v", mutexLocked)
	log.Printf("当前状态, mutexWoken:%+v", mutexWoken)
	log.Printf("当前状态, mutexStarving:%+v", mutexStarving)
	log.Printf("当前状态, mutexWaiterShift:%+v", mutexWaiterShift)
	// 对int32的, 旧值和新值进行比较和交换
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) { //加锁状态为1
		return true
	}

	// 如果处于唤醒、加锁或者饥饿状态，这次请求就不参与竞争了，返回false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	// 逻辑与运算
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	// 尝试在竞争的状态下请求锁
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}

func try() {
	var mu Mutex
	// mu.Lock()

	go func() { // 启动一个goroutine持有一段时间的锁
		mu.Lock()
		log.Printf("加锁")
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		time.Sleep(time.Second * 2)
		mu.Unlock()
		log.Printf("等待若干时间后, 解锁")

	}()

	// time.Sleep(time.Second * 1)
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	log.Printf("try方法中, 尝试获取锁")
	ok := mu.TryLock() // 尝试获取到锁
	if ok {            // 获取成功
		log.Printf("got the lock")
		// do something
		mu.Unlock()
		return
	}

	// 没有获取到
	fmt.Println("can't get the lock")
}

func main() {
	var mutex Mutex
	// mutex.Lock()
	log.Printf("尝试获取锁:%+v", mutex.TryLock())

	//以下为测试代码
	try()
}
