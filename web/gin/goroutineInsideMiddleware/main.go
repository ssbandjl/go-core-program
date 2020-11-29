package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		// 创建一个Gin上下文的副本, 准备在协程Goroutine中使用
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			// 模拟长时间任务,这里是5秒
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			// 在中间件或者控制器中启动协程时, 不能直接使用原来的上下文, 必须使用一个只读的上线文副本
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		// 没有使用协程时, 可以直接使用Gin上下文
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

/*
模拟同步阻塞访问:http://localhost:8080/long_sync
模拟异步非阻塞访问:http://localhost:8080/long_async
*/
