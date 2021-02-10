package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	src := []byte(`package main
import "fmt"
func main() {
  fmt.Println("Hello, world!")
}
`)

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(file, func(n ast.Node) bool {
		//这边我们所做的就是从所有节点中找出所有类型为 *ast.CallExpr 的节点，这些节点就代表了函数调用。我们会通过使用 printer 包，传入 Fun 成员变量，来打印函数的名称, 输出参考:fmt.Println
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		printer.Fprint(os.Stdout, fset, call.Fun)
		fmt.Println()

		return false
	})
}
