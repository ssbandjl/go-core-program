package main

import (
	"os"
	"strings"
	"text/template"
)

//title是函数,将首字母大写
const templateText = `
Output 0: {{title .Name1}}
Output 1: {{title .Name2}}
Output 2: {{.Name3 | title}}
`

func main() {
	funcMap := template.FuncMap{"title": strings.Title} //注册名为title的自定义函数
	tpl, _ := template.New("go-programming-tour").Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "go",
		"Name2": "programming",
		"Name3": "tour",
	}
	_ = tpl.Execute(os.Stdout, data)
}
