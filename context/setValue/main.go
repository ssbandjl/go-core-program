package main

import (
	"context"
	"fmt"
	"time"
)

var (
	key = "mykey"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, key, "add value")

	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//get value
			fmt.Println(ctx.Value(key), "设置成功，调用ctx.Done退出，参考信息:is cancel")

			return
		default:
			//get value
			fmt.Println(ctx.Value(key), "int goroutine")

			time.Sleep(2 * time.Second)
		}
	}
}
