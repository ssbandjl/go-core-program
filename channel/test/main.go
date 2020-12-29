package main
import (
	"log"
)

func ToCh(ch chan int) {
	for i:=0;i<10;i++{
		ch <- i
	}
	close(ch)
}

func main(){
	ch := make(chan int)
	go ToCh(ch)
	for i := range ch {
		log.Printf("%d", i)
	}
}
