//懒汉式单例模式
package lazy

import "sync"

type singleton struct {
	count int
}

var (
	instance *singleton //结构体的默认值不是nil,所以使用指针,指针的默认值是nil
	mutex    sync.Mutex
)

//每次创建都会执行锁检查
func New() *singleton {
	mutex.Lock()
	if instance == nil {
		instance = new(singleton)
	}
	mutex.Unlock()
	return instance
}

//双重检查,instance实例化后,锁永远不执行,提高性能
func New2() *singleton {
	if instance == nil { //第一次检查
		//这里可能有多个goroutine同时到达
		mutex.Lock()
		//这里同时只能有一个goroutine
		if instance == nil { //第二次检查
			instance = new(singleton)
		}
		mutex.Unlock()
	}
	return instance
}

func (s *singleton) Add() int {
	s.count++
	return s.count
}
