package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func md5sum(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return "", nil
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
		return "", nil
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func main() {
	result, err := md5sum("/Users/xb/gitlab/go/clientgo/kubectl/exec/cpPodFile2Local/output.txt")
	if err != nil {
		log.Printf("md5sum错误,%s", err.Error())
	}
	log.Printf("结果:%s", result)
}
