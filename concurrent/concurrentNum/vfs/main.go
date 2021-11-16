package main

import (
	"log"

	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/gatefs"
)

// fatefs子包控制访问该虚拟文件系统的最大并发数
func main() {
	fs := gatefs.New(vfs.OS("/path"), make(chan bool, 8))
	log.Printf("文件系统:%+v", fs)
}
