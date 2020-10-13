package main

import (
	"log"
	"time"
	_ "net/http/pprof"
)

var datas []string

func main(){
	go func(){
		for {
			log.Printf("len: %d", Add("go-programming-tour-book"))
			time.Sleep(time.Millisecond * 10)
		}
	}()
	_ = http.ListenAndServer("0.0.0.0:6060", nil)
}

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}
