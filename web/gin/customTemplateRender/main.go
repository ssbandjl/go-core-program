package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	router := gin.Default()
	//template.ParseFiles(文件1,文件2...)创建一个模板对象, 然后解析一组模板，使用文件名作为模板的名字
	// Must方法将模板和错误进行包裹, 返回模板的内存地址 一般用于变量初始化,比如:var t = template.Must(template.New("name").Parse("html"))
	html := template.Must(template.ParseFiles("file1", "file2"))
	router.SetHTMLTemplate(html) //关联模板和HTML渲染器

	router.GET("/index", func(c *gin.Context) {
		//HTML方法设置响应码, 模板文件名, 渲染替换模板中的值, 设置响应内容类型Content-Type "text/html"
		c.HTML(http.StatusOK, "file1", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
