/*
golang深入源代码系列之一AST的遍历:https://studygolang.com/articles/19353
何为语法树:https://huang-jerryc.com/2016/03/15/%E4%BD%95%E4%B8%BA%E8%AF%AD%E6%B3%95%E6%A0%91/
*/

package main

import (
	"fmt"
	"go/token"
	"testing"
	"text/scanner"
)

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
