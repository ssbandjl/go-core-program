package main

import (
	"log"
	"strings"
)

func main() {
	filePathArray := strings.Split("https://s3.harix.iamidata.com/dms/dms/mysql/cloud-sit-2/75/cms-mysql/20210922184824/cms-mysql.binlog-on.dump.gz", "/")
	log.Printf("filePathArray:%v", filePathArray)
	filePathSlice := filePathArray[5:]
	log.Printf("filePathSlice:%v", filePathSlice)

	bucket := filePathArray[4]
	log.Printf("bucket:%v", bucket)
	objectPath := strings.Join(filePathSlice, "/")
	log.Printf("objectPath:%v", objectPath)

}
