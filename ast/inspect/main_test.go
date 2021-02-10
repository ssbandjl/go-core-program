package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"testing"
)

func TestInspectAST(t *testing.T) {
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "./example/test1.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	// golang提供了ast.Inspect方法供我们遍历整个AST树，比如下面例子遍历整个example/test1.go文件寻找所有return返回的地方
	ast.Inspect(f, func(n ast.Node) bool {
		// Find Return Statements 查找返回语句
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			fmt.Printf("return statement found on line %v:\n", fset.Position(ret.Pos()))
			printer.Fprint(os.Stdout, fset, ret)
			fmt.Printf("\n")
			return true
		}
		return true
	})
}
