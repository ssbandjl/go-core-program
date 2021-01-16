package main

import (
	"encoding/gob" // /usr/local/go/src/encoding/gob
	"fmt"
	"os"
)

// main 执行godoc后的HTML页面:https://golang.org/pkg/encoding/gob/
func main() {
	info := map[string]string{
		"name":    "C语言中文网",
		"website": "http://c.biancheng.net/golang/",
	}
	name := "demo.gob"
	File, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	defer File.Close()
	enc := gob.NewEncoder(File)
	if err := enc.Encode(info); err != nil {
		fmt.Println(err)
	}
}
