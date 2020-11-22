package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.New()
	//
	////默认Gin引擎Engine使用了Logger(记录请求信息)Recovery(从Panic中恢复)中间件: engine.Use(Logger(), Recovery())
	////Logger 支持记录一个API请求发生时间，返回Status Code，Latency(耗时)，远端IP，请求方法, 请求路径(URL Path)
	//r := gin.Default()

	//自定义日志输出
	//2020-11-21T13:49:54+08:00 [INFO] "GET /ping HTTP/1.1 200 90.317µs "curl/7.64.1" "
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s [INFO] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Response": "OK"})
	})

	r.Run(":8082")
}
