package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//router.GET("/local/file", func(c *gin.Context) {
	router.POST("/local/file", func(c *gin.Context) {
		c.File("main.go")
	})

	//var fs http.FileSystem = // ...
	//router.GET("/fs/file", func(c *gin.Context) {
	//c.FileFromFS("fs/file.go", fs)
	//})
	router.Run(":9999")
}
