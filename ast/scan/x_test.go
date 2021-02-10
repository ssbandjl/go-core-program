/*
golang深入源代码系列之一AST的遍历:https://studygolang.com/articles/19353
何为语法树:https://huang-jerryc.com/2016/03/15/%E4%BD%95%E4%B8%BA%E8%AF%AD%E6%B3%95%E6%A0%91/
*/

package scan

import (
	"fmt"
	"go/token"
	"testing"
	"text/scanner"
)

//  我们先创建源码字符串，然后初始化 scanner.Scanner 来扫描我们的源码。我们可以通过不停地调用 Scan() 方法来获取 token 的位置，类型和字面量，直到遇到 EOF 标记
func TestScanner(t *testing.T) {
	src := []byte(`package main
import "fmt"
//comment
func main() {
  fmt.Println("Hello, world!")
}
`)

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))

	s.Init(file, src, nil, 0)

	for {
		pos, tok, lit := s.Scan()
		fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

		if tok == token.EOF {
			break
		}
	}
}
