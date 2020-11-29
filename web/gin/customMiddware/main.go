package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//自定义日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable 在gin上下文中设置键值对
		c.Set("example", "12345")

		// before request

		//Next方法只能用于中间件中,在当前中间件中, 从方法链执行挂起的处理器
		c.Next()

		// after request  打印中间件执行耗时
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending  打印本中间件的状态码
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
