package main

import (
	"fmt"
	"unicode/utf8"
)

const nihongo = "go语言"

func main() {
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString((nihongo[i:]))
		fmt.Printf("%#U starts at byte positon %d\n", runeValue, i)
		w = width
	}
}
