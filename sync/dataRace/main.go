package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

func main(){
	c := make(chan bool)
	m := make(map[string]string)
	go func(){
		m["1"] = "a"  //运行时,这里有个协程对m进行写操作
		c <- true
	}()

	m["2"] = "b"  //主函数中也会对m进行写操作
	<-c  //根据通道先写后读的原则,这里会等待上面的go func执行完成
	for k, v := range m {
		fmt.Println(k, v)
	}
}



// ParallelWrite writes data to file1 and file2, returns the errors.
func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// This err is shared with the main goroutine, 这里的err变量与主线程共享了
			// so the write races with the write below.
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2") // The second conflicting write to err.这里也在对err进行写操作
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}


...
_, err := f1.Write(data)
...
_, err := f2.Write(data)
...



var service map[string]net.Addr

func RegisterService(name string, addr net.Addr) {
	service[name] = addr
}

func LookupService(name string) net.Addr {
	return service[name]
}




var (
	service   map[string]net.Addr
	serviceMu sync.Mutex
)

func RegisterService(name string, addr net.Addr) {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	service[name] = addr
}

func LookupService(name string) net.Addr {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	return service[name]
}



type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {
	w.last = time.Now().UnixNano() // First conflicting access. 写操作,与下面的写操作存在冲突
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			// Second conflicting access.
			if w.last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}



type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {
	atomic.StoreInt64(&w.last, time.Now().UnixNano()) //使用原子包存储方法
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {  //使用原子包的读取方法
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}

c := make(chan struct{}) // or buffered channel 这里也可以使用带缓冲的通道演示

// The race detector cannot derive the happens before relation
// for the following send and close operations. These two operations 下面的通道发送和关闭操作没有进行同步,导致冲突
// are unsynchronized and happen concurrently.
go func() { c <- struct{}{} }()
close(c)


c := make(chan struct{}) // or buffered channel

go func() { c <- struct{}{} }()
<-c //根据Go语言内存模型,一个通道上写操作发生在读操作之前,所以这里读取时,写已经完成,完成了通道同步
close(c)
