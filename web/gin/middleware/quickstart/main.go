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
	//r := gin.New()

	//默认Gin引擎Engine使用了Logger(记录请求信息)Recovery(从Panic中恢复)中间件: engine.Use(Logger(), Recovery())
	//Logger 支持记录一个API请求发生时间，返回Status Code，Latency(耗时)，远端IP，请求方法, 请求路径(URL Path)
	r := gin.Default()

	r.Use(Logger())

	r.GET("/ping", func(c *gin.Context) {
		example := c.MustGet("example").(string) //使用MustGet方法获取键值
		// 读取中间件在上下文中存储的内容
		log.Printf("中间件在上线文中存储的键值, key:example, value:%s", example)

		// Gin无法恢复gorouting中的panic, 因为任何gorouting中发生了panic，都会panic整个程序。每个gorouting需要自己处理panic
		go func() {
			panic("Panic in Gorouting")
		}()
		// Gin可以恢复Handler中的panic
		//panic("Panic in Handler!")

		// 返回
		c.JSON(http.StatusOK, gin.H{"Response": "OK"})
	})

	r.Run(":8082")
}
