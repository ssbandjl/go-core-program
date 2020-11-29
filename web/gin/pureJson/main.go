package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Serves unicode entities
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

/*
模拟请求,得到将HTML标签转义后的JSON字符串
curl http://localhost:8080/json
{"html":"\u003cb\u003eHello, world!\u003c/b\u003e"}
得到原始JSON字符串
curl http://localhost:8080/purejson
{"html":"<b>Hello, world!</b>"}
*/
