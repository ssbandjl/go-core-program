package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	// 待词法扫描的表达式
	src := []byte("cos(x) + 2i*sin(x) // Euler 欧拉")

	// 初始化 scanner
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)
	// 扫描, 每个标识符和运算符都被特定的token代替
	fmt.Printf("位置\t符号\t字符串字面量\n")
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
