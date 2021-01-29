package main

// 起了多个goroutine作为writer，每个writer内部随机生成字符串写进去。唯一的reader读取数据并打印：
// 参考链接:https://studygolang.com/articles/9665
// 官方链接:https://golang.org/pkg/io/#Pipe
// Pipe适用于，产生了一条数据，紧接着就要处理掉这条数据的场景, 比如先写后读, 而且因为其内部是一把大锁，因此是线程安全的, 由于没有用临时存储, 所以减少了内存使用
import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)

//利用当前时间的纳秒作为随机种子, 初始化并得到随机对象Rand
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func generate(id int, writer *io.PipeWriter) {
	// log.Printf("随机写入")
	i := 0
	for {
		if i == 10 {
			break
		}
		data := fmt.Sprintf("%d, 第%d次写入", id, i)
		_, err := writer.Write([]byte(data))
		if nil != err {
			log.Fatal(err)
		}
		// log.Printf("write %d bytes, %s", n, data)
		time.Sleep(200 * time.Millisecond)
		i++
	}
}

func main() {
	//利用pipe包, 初始化一根管道, 得到两个端,读端和写端
	rp, wp := io.Pipe()
	//多个写端并发执行
	for i := 0; i < 2; i++ {
		go generate(i, wp)
	}
	time.Sleep(1 * time.Second)
	data := make([]byte, 64)
	for {
		//一个读端, 从管道中的读取已经写入的数据
		n, err := rp.Read(data)
		if nil != err {
			log.Fatal(err)
		}
		if 0 != n {
			log.Println("read data:", n, string(data))
		}
		time.Sleep(100 * time.Millisecond)
	}
}
