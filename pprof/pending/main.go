package main

//模拟阻塞,调用通道chan, sync.Mutex同步锁, time.Sleep()都会造成阻塞
//访问地址:http://localhost:6060/debug/pprof/

//用pprof工具进行阻塞分析: go tool pprof http://localhost:6061/debug/pprof/mutex
//top 查看互斥排名
//list 查看指定函数代码情况

import (
	"net/http"
	_ "net/http/pprof" //以HTTP SERVER运行,采集运行时的性能数据
	"runtime"
	"sync"
)

func init() {
	runtime.SetMutexProfileFraction(1) //设置互斥锁采集频率,如果保持默认或设置为小于0,则不进行采集

	//runtime.SetBlockProfileRate(1) //Block Profiling分析
}

func main() {
	var m sync.Mutex
	var datas = make(map[int]struct{})
	for i := 0; i < 999; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			datas[i] = struct{}{}
		}(i)
	}

	_ = http.ListenAndServe(":6061", nil)
}
