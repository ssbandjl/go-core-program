package main

import (
	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors[]"` //标签中的colors[]数组切片与html文件中的name="colors[]"对应
}

func main() {
	r := gin.Default()

	//LoadHTMLGlob采用通配符模式匹配HTML文件,并将内容进行渲染,提供给前端访问
	r.LoadHTMLGlob("*.html")
	r.GET("/", indexHandler)
	r.POST("/", formHandler)

	r.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm) //Bind方法根据请求头类型Content-Type, 自动选择合适的绑定引擎,如Json/XML
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

//将html与main.go放到一个目录,执行go run main.go运行后, 访问http://localhost:8080,勾选复选框,然后提交测试
