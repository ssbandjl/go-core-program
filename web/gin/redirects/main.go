package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/") //重定向到外部链接
	})

	//重定向到内部链接
	r.GET("/internal", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/home")
	})

	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "这是首页"})
	})
	r.Run(":8080")
}

/*
模拟测试:
curl http://localhost:8080/test
*/
