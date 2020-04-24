package main

import (
	_ "fmt"
	"os"
	"text/template"
)

//下面是一个简单的例子，可以打印"17 of wool", 17件羊毛制品
type Inventory struct {
	Material string
	Count    uint
}

func main() {

	// 模板定义和解析模板
	tmpl, err := template.New("test").Parse("{{.Count}} of {{.Material}}")
	if err != nil {
		panic(err)
	}

	// 数据驱动模板
	var sweaters = Inventory{"wool", 17}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}
