package main
import (
	"fmt"
	"time"
)


//write Data
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//放入数据
		intChan<- i //
		fmt.Println("writeData ", i)
		//time.Sleep(time.Second)
	}
	close(intChan) //关闭
}

//read data
func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v) 
	}
	//readData 读取完数据后，即任务完成
	exitChan<- true
	close(exitChan)

}

func main() {

	//创建两个管道，这里的通道大小为10，但是我们定义的50个数据，编译器只要检测要通道有读取，就不会死锁，即使写得快读得慢；如果只有写，没有读，则通道会阻塞（死锁deadlock）
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	
	go writeData(intChan)
	go readData(intChan, exitChan)

	//time.Sleep(time.Second * 10)
	for {
		_, ok := <-exitChan
		if !ok {
			fmt.Println("协程执行完毕，退出通道exitChan已关闭，主线程检测到协程退出标识也退出")
			break
		}
	}

}