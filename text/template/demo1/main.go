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
	var sweaters = Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}
