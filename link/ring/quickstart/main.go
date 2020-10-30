package main

import (
	"container/ring"
	"log"
)

//环形链表快速入门

var r ring.Ring

func main() {
	log.Printf("环形链表长度:%d", r.Len())
	log.Printf("环形链表当前指针指向的值(零值):%v", r.Value)
}
