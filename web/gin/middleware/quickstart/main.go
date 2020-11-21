package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 自定义Middleware，返回类型为gin.HandlerFunc
func Logger() gin.HandlerFunc {
	// 输入参数为 *gin.Context
	return func(c *gin.Context) {
		// 实际请求处理函数之前的动作
		t := time.Now()

		// 上下文中可以存贮KV(设置键值对)，后续处理函数中使用
		c.Set("example", "12345")
		log.Printf("设置键值对, key:example, value:12345")

		c.Next()

		// 请求处理函数之后的动作
		latency := time.Since(t)
		log.Printf("执行中间件耗时(延迟):%s", latency)

		// 读取Handler处理的结果
		status := c.Writer.Status()
		log.Printf("执行中间件响应状态码:%d", status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/ping", func(c *gin.Context) {
		example := c.MustGet("example").(string) //使用MustGet方法获取键值
		// 读取中间件在上下文中存储的内容
		log.Printf("中间件在上线文中存储的键值, key:example, value:%s", example)
		// 返回
		c.JSON(http.StatusOK, gin.H{"Response": "OK"})
	})

	r.Run(":8082")
}
