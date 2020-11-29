package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()                        //Date方法返回年,月,日
	return fmt.Sprintf("%d%02d/%02d", year, month, day) //格式化时间
}

func main() {
	router := gin.Default()
	router.Delims("{[{", "}]}") //自定义模板中的左右分隔符
	//SetFuncMap方法用给定的template.FuncMap设置到Gin引擎上, 后面模板渲染时会调用同名方法
	//FuncMap是一个map,键名关联方法名, 键值关联方法, 每个方法必须返回一个值, 或者返回两个值,其中第二个是error类型
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("./testdata/template/raw.tmpl") //加载单个模板文件并与HTML渲染器关联

	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", gin.H{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	router.Run(":8080")
}

/*
模拟测试:
curl http://localhost:8080/raw
*/
