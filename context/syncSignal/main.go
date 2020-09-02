package main

import (
	"context"
	"fmt"
	"time"
)

/*
使用Context同步信号
我们可以通过一个代码片段了解 context.Context 是如何对信号进行同步的。在这段代码中，我们创建了一个过期时间为 1s 的上下文，并向上下文传入 handle 函数，该方法会使用 500ms 的时间处理传入的『请求』
因为过期时间大于处理时间，所以我们有足够的时间处理该『请求』，运行上述代码会打印出如下所示的内容

handle 函数没有进入超时的 select 分支，但是 main 函数的 select 却会等待 context.Context 的超时并打印出 main context deadline exceeded

相信这两个例子能够帮助各位读者理解 context.Context 的使用方法和设计原理 — 多个 Goroutine 同时订阅 ctx.Done() 管道中的消息，一旦接收到取消信号就立刻停止当前正在执行的工作
*/

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)

	select {
	case <-ctx.Done(): //main函数，正常等待上下文超时之后发送超时信号给上下文
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done(): //超时或取消
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration): //处理时间到了，处理时间，小于上下文超时时间，走此分支
		fmt.Println("process request with", duration)
	}
}
