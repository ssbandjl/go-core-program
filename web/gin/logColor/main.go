package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 默认输出到控制台的日志颜色是根据您使用的虚拟终端TTY来着色的
	// Disable log's color 禁用日志颜色
	gin.DisableConsoleColor()

	// Force log's color 强制开启日志颜色
	//gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}

//模拟请求测试: curl http://localhost:8080/ping
