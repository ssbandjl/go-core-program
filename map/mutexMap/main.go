package main

import "sync"

// MutexMap 是一个简单的 map + sync.Mutex 的并发安全散列表实现
type MutexMap struct {
	data map[interface{}]interface{}
	mu   sync.Mutex
}

func (m *MutexMap) Load(k interface{}) (v interface{}, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok = m.data[k]
	return
}

func (m *MutexMap) Store(k, v interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[k] = v
}
