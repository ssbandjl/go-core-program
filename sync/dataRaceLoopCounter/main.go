package main

import (
	"fmt"
	"sync"
)

func main() {
	//以下代码存在冲突,i不是看到的i值
	//var wg sync.WaitGroup
	//wg.Add(5)
	//for i := 0; i < 5; i++ {
	//	go func() {
	//		fmt.Println(i) // Not the 'i' you are looking for.
	//		wg.Done()
	//	}()
	//}

	//读取本地拷贝值,解决冲突
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Println(j) // Good. Read local copy of the loop counter.
			wg.Done()
		}(i)
	}
	wg.Wait()

}
