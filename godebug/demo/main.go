package main

//GODUBUG运行:GODEBUG=schedtrace=1000 go run main.go
//schedtrace,单位毫秒,表示可以在运行时没schedtrace毫秒发出一行调度器摘要信息到标准err输出中
//scheddetail=1表示多行详细信息: GODEBUG=schedtrace=1000,scheddetail=1 go run main.go
//参考输出:SCHED 1009ms: gomaxprocs=8 idleprocs=8 threads=18 spinningthreads=0 idlethreads=11 runqueue=0 [0 0 0 0 0 0 0 0]

import "sync"

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			var counter int
			for i := 0; i < 1e10; i++ {
				counter++
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
