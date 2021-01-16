package main

import (
	"log"
	"sync"
	"sync/atomic"
)

var (
	id      int64
	syncMap sync.Map
)

func main() {
	id = atomic.AddInt64(&id, 1)
	id = atomic.AddInt64(&id, 1)
	id = atomic.AddInt64(&id, 1)
	id = atomic.AddInt64(&id, 1)
	//重置id
	id = 0
	id = atomic.AddInt64(&id, 1)
	id = atomic.AddInt64(&id, 1)

	log.Printf("自增后的id:%d", id)
}
