package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	p := filepath.FromSlash("path/to/file")
	fmt.Println("Path: " + p)
}
