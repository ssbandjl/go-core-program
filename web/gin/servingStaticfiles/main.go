package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	cwd, _ := os.Getwd() //获取当前文件目录
	log.Printf("当前项目路径:%s", cwd)
	router.Static("/static", cwd)                          //提供静态文件服务器, 第一个参数为相对路径,第二个参数为根路径, 这个路径一般放置css,js,fonts等静态文件,前端html中采用/static/js/xxx或/static/css/xxx等相对路径的方式引用
	router.StaticFS("/more_static", http.Dir("./"))        //将本地文件树结构映射到前端, 通过浏览器可以访问本地文件系统, 模拟访问:http://localhost:8080/more_static
	router.StaticFile("/logo.png", "./resources/logo.png") //StaticFile提供单静态单文件服务, 模拟访问:http://localhost:8080/log.png

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
