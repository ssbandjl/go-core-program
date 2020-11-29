package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// 一行式LetsEncrypt证书, https服务
	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}
