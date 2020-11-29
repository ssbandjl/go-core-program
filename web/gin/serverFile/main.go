package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/local/file", func(c *gin.Context) {
		c.File("./main.go")
	})

	// A FileSystem implements access to a collection of named files.
	// The elements in a file path are separated by slash ('/', U+002F)
	// characters, regardless of host operating system convention.
	// FileSystem接口, 要求实现文件的访问的方法, 提供文件访问服务根路径的HTTP处理器
	var fs http.FileSystem = http.Dir("./") //将本地目录作为文件服务根路径
	router.GET("/fs/file", func(c *gin.Context) {
		c.FileFromFS("main.go", fs) //将文件服务系统下的文件数据返回
	})
	router.Run(":8080")
}

/*
模拟访问文件数据:
curl http://localhost:8080/local/file

模拟访问文件系统下的文件数据:
curl http://localhost:8080/fs/file
*/
