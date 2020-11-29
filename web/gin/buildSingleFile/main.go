package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	r := gin.New()

	t, err := loadTemplate() //加载go-assets-builder生成的模板
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/index.tmpl", nil)
	})
	r.Run(":8080")
}

// loadTemplate loads templates embedded by go-assets-builder
// 加载go-assets-builder生成的资产文件, 返回模板的地址
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		defer file.Close()
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") { //跳过目录或没有.tmpl后缀的文件
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h)) //新建一个模板, 文件名做为模板名, 文件内容作为模板内容
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

/*
使用方法:
1.下载依赖包
go get github.com/gin-gonic/gin
go get github.com/jessevdk/go-assets-builder

2.将html文件夹(包含html代码)生成为go资产文件assets.go
go-assets-builder html -o assets.go

3.构建服务
go build -o assets-in-binary

4.运行服务
./assets-in-binary
*/
