package main

import (
	"log"
	"os"
	"path/filepath"
)

func GetDirFileSize(path string) (int64, error) {
	var size int64
	// Walk 从根开始遍历文件数, 对每个文件或目录执行WalkFunc遍历方法
	// type WalkFunc func(path string, info os.FileInfo, err error) error
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() { //如果是文件, 则获取该文件字节
			size += info.Size()
		}
		return err
	})
	return size, err
}

func RemoveDirFile(path string) error {
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filePath := path + "/" + info.Name()
			err = os.Remove(filePath)
			if err != nil {
				log.Printf("删除文件失败, 文件:%s", filePath)
				return err
			}
			log.Printf("删除文件成功, 文件:%s, 大小:%d", filePath, info.Size())

		}
		return err
	})
	return err
}

func main() {
	size, err := GetDirFileSize("/Users/xb/gitlab/go/go_core_program/os")
	if err != nil {
		log.Printf("获取目录大小失败, %s", err.Error())
	}
	log.Printf("目录:/Users/xb/gitlab/go/go_core_program, 大小:%d", size)

	err = RemoveDirFile("/Users/xb/gitlab/go/go_core_program/tmp")
	if err != nil {
		log.Printf("清理目录下的文件失败, %s", err.Error())
	}
}
