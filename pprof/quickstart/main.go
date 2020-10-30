package main

//访问地址:http://localhost:6060/debug/pprof/

//通过交互式终端使用:go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=60
//top10:查看资源开销排名前十的函数

//下载profil文件用于离线分析:wget http://localhost:6060/debug/pprof/profile

//查看可视化界面,函数调用流程,Top等:go tool pprof -http=:6001 profile   MAC检查插件安装:dot -h 安装:brew install graphviz
import (
	"log"
	"net/http"
	_ "net/http/pprof" //以HTTP SERVER运行,采集运行时的性能数据
	"time"
)

var datas []string

func main() {
	go func() {
		for {
			log.Printf("len: %d", Add("go-programming-tour-book"))
			time.Sleep(time.Millisecond * 10)
		}
	}()
	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}
