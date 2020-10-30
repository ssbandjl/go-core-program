package main

//打开垃圾收集器GC运行信息到标准错误流,用于观察GC运行状况:GODEBUG=gctrace=1 go run main.go
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
