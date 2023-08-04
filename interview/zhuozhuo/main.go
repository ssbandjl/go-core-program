package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

/**
你需要实现的目标函数 target

@param id 是一个随机字符串，例如 6A10A467-2842-A460-5353-DBE7D41986B7；
@param job 函数是一个耗时操作，例如：去数据库 query 数据，可能耗时 500ms；
@return count 表示在执行本次 job 期间有多少相同的 id 调用过 target

关键特性：相同 id 并发调用 target，target 只执行一次 job 函数，举例来说：
第一个线程传入 id 为 "id-123" 调用 target，job 函数开始执行，在此期间，又有其他 4 个线程以 id 为 "id-123" 调用了 target；
在此期间，只有一个 job 函数执行，等它执行完成后，上述 5 个线程均收到返回值 count=5，表示这段时间有 5 个相同 id 进行了调用；
*/

func target(id string, job func()) (count int) {
	var ret int
	var mutex sync.Mutex
	cond := sync.Cond{L: &mutex}
	idCallNumLock.Lock()
	idCount = atomic.AddInt64(&idCount, 1)
	idCallNumCount, ok := idCallNum[id]
	if ok {
		log.Printf("already exec")
	} else {
		job()
	}
	idCallNumCount++
	idCallNum[id] = idCallNumCount
	idCallNumLock.Unlock()
	//TODO implement this
	log.Printf("id:%s, idCallNumCount:%d\n", id, idCallNumCount)
	go func() {
		for {
			log.Printf("idCount:%d, idCount:%d", idCount, idCount)
			if idCount == 5 {
				// cond.L.Lock()
				cond.Broadcast()
				// cond.L.Unlock()
				ret = int(idCount)
				goto Loop
			}
			time.Sleep(30 * time.Millisecond)
		}
	Loop:
		return
	}()
	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()
	return int(ret)
}

//用来模拟 job 函数的变量
//不要修改
var (
	counter     int
	counterLock sync.Mutex
	// add
	idCallNum     = make(map[string]int, 50)
	idCallNumLock sync.Mutex
	// ch            = make(chan int64, 50)
	// arr     = make([]int, 10)
	idCount int64
)

//用来模拟耗时，时间不固定，实现 target 时不能依赖此时间
//不要修改
const (
	mockJobTimeout = 300 * time.Millisecond
	tolerate       = 30 * time.Millisecond
)

//测试用的 job 函数，是一个计数器，用来模拟耗时操作
//不要修改
func mockJob() {
	time.Sleep(mockJobTimeout)
	counterLock.Lock()
	counter++
	counterLock.Unlock()
}

//相同 id 并行调用
//不要修改
func testCaseSampleIdParallel() {
	counter = 0 //重置计数器
	const (
		id     = "CBD225E1-B7D9-BE76-9735-1D0A9B62EE4D"
		repeat = 5 //用来模拟相同 id 的多次重复调用，调用次数不固定，实现 target 时不能依赖此调用次数
	)
	wg := sync.WaitGroup{}
	wg.Add(repeat)
	tStart := time.Now()
	for i := 0; i < repeat; i++ {
		go func() {
			count := target(id, mockJob)
			wg.Done()
			if count != repeat {
				panic(fmt.Sprintln("[testCaseSampleIdConcurrence] count:", count, "!= repeat:", repeat))
			}
		}()
	}
	wg.Wait()
	if counter != 1 { //应该只调用了一次 job 函数
		panic(fmt.Sprintln("[testCaseSampleIdConcurrence] counter:", counter, "!= 1"))
	}
	var (
		tDelta  = time.Now().Sub(tStart)
		tExpect = mockJobTimeout + tolerate
	)
	if tDelta > tExpect {
		panic(fmt.Sprintln("[testCaseRandomId] timeout", tDelta, ">", tExpect))
	}
}

//相同 id 串行调用
//不要修改
func testCaseSampleIdSerial() {
	counter = 0
	const (
		id     = "3E5A5C8D-B254-383B-4F33-F6927578FD11"
		repeat = 2
	)
	tStart := time.Now()
	for i := 0; i < repeat; i++ {
		count := target(id, mockJob)
		if count != 1 {
			panic(fmt.Sprintln("[testCaseSampleIdSerial] count:", count, "!= 1"))
		}
	}
	if counter != repeat { //虽然是相同 id，但因为是串行调用，应该执行 repeat 次 job 函数
		panic(fmt.Sprintln("[testCaseSampleIdSerial] counter:", counter, "!= repeat:", repeat))
	}
	var (
		tDelta  = time.Now().Sub(tStart)
		tExpect = repeat*mockJobTimeout + tolerate
	)
	if tDelta > tExpect {
		panic(fmt.Sprintln("[testCaseSampleIdSerial] timeout", tDelta, ">", tExpect))
	}
}

//不同 id 并行调用
//不要修改
func testCaseRandomId() {
	counter = 0 //重置计数器
	ids := []string{
		"id-3",
		"id-3",
		"id-3",

		"id-2",
		"id-2",

		"id-1",
	}
	wg := sync.WaitGroup{}
	wg.Add(len(ids))
	tStart := time.Now()
	for _, id := range ids {
		id := id
		go func() {
			count := target(id, mockJob)
			wg.Done()
			var expectedCount int
			switch id {
			case "id-1":
				expectedCount = 1
			case "id-2":
				expectedCount = 2
			case "id-3":
				expectedCount = 3
			}
			if count != expectedCount {
				panic(fmt.Sprintln("[testCaseRandomId] count:", count, "!= expectedCount:", expectedCount, "id:", id))
			}
		}()
	}
	wg.Wait()
	if counter != 3 { //3个不同的 id 同时并发调用，job 函数应该执行 3 次
		panic(fmt.Sprintln("[testCaseSampleIdConcurrence] counter:", counter, "!= 3"))
	}
	var (
		tDelta  = time.Now().Sub(tStart)
		tExpect = 3*mockJobTimeout + tolerate
	)
	if tDelta > tExpect {
		panic(fmt.Sprintln("[testCaseRandomId] timeout", tDelta, ">", tExpect))
	}
}

//不要修改
func main() {
	const repeat = 50
	for i := 0; i < repeat; i++ {
		testCaseSampleIdParallel()
		testCaseSampleIdSerial()
		testCaseRandomId()
		fmt.Print("\r", i+1, "/", repeat, " ✔ ")
	}
	fmt.Println("🎉 All Tests Passed!")
}
