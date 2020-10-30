package main

//Timer:计时器类型表示一个事件,当计时器过期时,会将当前时间发送到通道C上(After创建的不会),计时器必须使用NewTimer()或After()创建
import (
	"fmt"
	"time"
)

func main() {
	//创建计时器
	timer1 := time.NewTimer(time.Second * 5)
	fmt.Printf("NewTimer创建对象的类型:%T\n", timer1)
	fmt.Println("当前时间:", time.Now())
	data := <-timer1.C
	fmt.Printf("timer1通道C的类型:%T\n", timer1.C)
	fmt.Printf("timer1通道C中读出一个数据data的类型:%T\n", data)
	fmt.Println("timer1通道C中读出一个数据data:", data)
}
