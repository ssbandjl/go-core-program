package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,                                     //Prompt指定一个回调函数有条件的接受证书机构CA的TOS服务, 使用AcceptTOS总是接受服务条款
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"), //HostPolicy用于控制指定哪些域名, 管理器将检索新证书
		Cache:      autocert.DirCache("/var/www/.cache"),                   //缓存证书和其他状态
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
