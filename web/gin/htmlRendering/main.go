package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	//LoadHTMLGlob方法以glob模式加载匹配的HTML文件, 并与HTML渲染器结合
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		//HTML方法设置响应码, 模板文件名, 渲染替换模板中的值, 设置响应内容类型Content-Type "text/html"
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}

/*
模拟测试:
curl http://localhost:8080/index
*/
