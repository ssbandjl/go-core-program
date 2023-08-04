package main

import (
	"fmt"
	"strings"
)

func main() {
	// reg := []string{"a", "b", "c"}
	reg := []string{"a", "b", "c"}
	fmt.Println(strings.Join(reg, "|"))
}
