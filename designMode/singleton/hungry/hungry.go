//饿汉式单例模式
package main

type singleton struct {
	count int
}

var Instance = new(singleton)

func (s *singleton) Add() int {
	s.count++
	return s.count
}
