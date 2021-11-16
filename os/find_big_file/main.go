package main

import (
	"log"
	"os"
	"path/filepath"
	"sort"
)

const ScanDir = "/Users/xb/Downloads" // 	"/Users/xb/Downloads/often"
const MaxFileMB = 100

var fileSizeMap map[string]int64

func GetDirFileSize(path string) error {
	err := filepath.Walk(path, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			log.Printf("erro:%s", err.Error())
			// return err
		}
		if !fileInfo.IsDir() { //如果是文件, 则获取该文件字节
			fileSize := fileInfo.Size() / 1024 / 1024
			if fileSize > MaxFileMB {
				fileSizeMap[filePath] = fileSize
			}
		}
		return nil
	})
	log.Printf("error:%+v", err.Error())
	return nil
}

//要对golang map按照value进行排序，思路是直接不用map，用struct存放key和value，实现sort接口，就可以调用sort.Sort进行排序了。
// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int64
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int64) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func main() {
	fileSizeMap = make(map[string]int64, 10000000)
	err := GetDirFileSize(ScanDir)
	if err != nil {
		log.Printf("获取目录大小失败, %s", err.Error())
	}
	pairList := sortMapByValue(fileSizeMap)
	for index, fileSizeObj := range pairList {
		log.Printf("%d %s %d", index, fileSizeObj.Key, fileSizeObj.Value)
	}
}
