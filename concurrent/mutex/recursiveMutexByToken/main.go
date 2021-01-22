package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

// Token方式的递归锁
type TokenRecursiveMutex struct {
	sync.Mutex
	token     int64
	recursion int32
}

// 请求锁，需要传入token
func (m *TokenRecursiveMutex) Lock(token int64) {
	if atomic.LoadInt64(&m.token) == token { //如果传入的token和持有锁的token一致，说明是递归调用
		m.recursion++
		return
	}
	m.Mutex.Lock() // 传入的token不一致，说明不是递归调用
	// 抢到锁之后记录这个token
	atomic.StoreInt64(&m.token, token)
	m.recursion = 1
}

// 释放锁
func (m *TokenRecursiveMutex) Unlock(token int64) {
	if atomic.LoadInt64(&m.token) != token { // 释放其它token持有的锁
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.token, token))
	}
	m.recursion--         // 当前持有这个锁的token释放锁
	if m.recursion != 0 { // 还没有回退到最初的递归调用
		return
	}
	atomic.StoreInt64(&m.token, 0) // 没有递归调用了，释放锁
	m.Mutex.Unlock()
}

func main() {
	var recursiveMutex TokenRecursiveMutex
	log.Printf("重复上锁")
	var token int64
	token = 123456
	recursiveMutex.Lock(token)
	recursiveMutex.Lock(token)
	recursiveMutex.Lock(token)
	log.Printf("多次释放锁, 当前协程标识:%d, 当前锁次数:%d", recursiveMutex.token, recursiveMutex.recursion)
	recursiveMutex.Unlock(token)
	log.Printf("多次释放锁, 当前协程标识:%d, 当前锁次数:%d", recursiveMutex.token, recursiveMutex.recursion)
	recursiveMutex.Unlock(token)
	log.Printf("多次释放锁, 当前协程标识:%d, 当前锁次数:%d", recursiveMutex.token, recursiveMutex.recursion)
	recursiveMutex.Unlock(token)
	log.Printf("多次释放锁, 当前协程标识:%d, 当前锁次数:%d", recursiveMutex.token, recursiveMutex.recursion)
}
