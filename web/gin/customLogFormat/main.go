package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	// type LogFormatter func(params LogFormatterParams) string 这里的LogFormatterParams是一个格式化日志参数的结构体
	// LoggerWithFormatter instance a Logger middleware with the specified log format function.
	// LoggerWithFormatter方法实例化一个日志器Logger中间件,并带有特殊的日志格式
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		// 127.0.0.1 - [Sun, 22 Nov 2020 17:09:53 CST] "GET /ping HTTP/1.1 200 56.113µs "curl/7.64.1" "
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,                       //请求客户端的IP地址
			param.TimeStamp.Format(time.RFC1123), //请求时间
			param.Method,                         //请求方法
			param.Path,                           //路由路径
			param.Request.Proto,                  //请求协议
			param.StatusCode,                     //http响应码
			param.Latency,                        //请求到响应的延时
			param.Request.UserAgent(),            //客户端代理程序
			param.ErrorMessage,                   //如果有错误,也打印错误信息
		)
	}))
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}

//模拟请求测试: curl http://localhost:8080/ping
