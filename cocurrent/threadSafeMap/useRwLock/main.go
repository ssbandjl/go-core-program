package main

//利用读写锁实现线程安全的 map[int]int类型
import (
	"log"
	"sync"
)

type RWMap struct { // 一个读写锁保护的线程安全的map
	sync.RWMutex // 读写锁保护下面的map字段
	m            map[int]int
}

// 新建一个RWMap
func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int, n),
	}
}
func (m *RWMap) Get(k int) (int, bool) { //从map中读取一个值
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k] // 在锁的保护下从map中读取
	return v, existed
}

func (m *RWMap) Set(k int, v int) { // 设置一个键值对
	m.Lock() // 锁保护
	defer m.Unlock()
	m.m[k] = v
}

func (m *RWMap) Delete(k int) { //删除一个键
	m.Lock() // 锁保护
	defer m.Unlock()
	delete(m.m, k)
}

func (m *RWMap) Len() int { // map的长度
	m.RLock() // 锁保护
	defer m.RUnlock()
	return len(m.m)
}

func (m *RWMap) Each(f func(k, v int) bool) { // 遍历map
	m.RLock() //遍历期间一直持有读锁
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}

func main() {
	//新建线程安全的map
	m := NewRWMap(100)
	//设置值
	m.Set(1, 100)
	m.Set(2, 200)
	m.Set(3, 300)
	//获取长度
	log.Printf("map长度:%d", m.Len())
	//获取值
	value1, ok := m.Get(1)
	if ok {
		log.Printf("key=1, value=%d", value1)
	}
}
