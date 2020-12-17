package main

import (
	"fmt"
	"log"
)

type T int

func IsClosed(ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

//违反了通道关闭原则, 不建议使用, 这样的粗鲁方法不仅违反了通道关闭原则，而且Go白皮书和标准编译器不保证它的实现中不存在数据竞争
func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// 一个函数的返回结果可以在defer调用中修改。
			justClosed = false
		}
	}()

	// 假设ch != nil。
	close(ch)   // 如果ch已关闭，则产生一个恐慌。
	return true // <=> justClosed = true; return
}

func main() {
	c := make(chan T)
	fmt.Println(IsClosed(c)) // false
	//close(c)
	log.Printf("第一次安全关闭:%v", SafeClose(c))
	log.Printf("第二次安全关闭:%v", SafeClose(c))
	fmt.Println(IsClosed(c)) // true
}
