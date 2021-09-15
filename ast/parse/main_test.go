/*
golang深入源代码系列之一：AST的遍历:https://studygolang.com/articles/19353
*/
package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// go test
func TestParserAST(t *testing.T) { //测试解析抽象语法树
	src := []byte(`/*comment0*/
package main
import "fmt"
//comment1
/*comment2*/
func main() {
  fmt.Println("Hello, world!")
}
`)

	// Create the AST by parsing src.
	// NewFileSet 创建一个新的文件集
	fset := token.NewFileSet() // positions are relative to fset
	// func parser.ParseFile(fset *token.FileSet, filename string, src interface{}, mode parser.Mode) (f *ast.File, err error)
	//ParseFile 解析单个源代码文件,并返回对应的抽象语法树文件节点ast.File node, 源代码可以通过filename或者src参数提供
	//同样注意没有扫描出注释，需要的话要将parser.ParseFile的最后一个参数改为parser.ParseComments
	// mode参数控制已解析的源文本和其他可选解析器功能的总数。位置信息记录在文件集fset中，它不能为空
	f, err := parser.ParseFile(fset, "", src, 0)

	if err != nil {
		panic(err)
	}

	// Print the AST.
	// 打印抽象语法树
	ast.Print(fset, f)
}
