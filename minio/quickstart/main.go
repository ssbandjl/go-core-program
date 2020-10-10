package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"log"
)

func main() {
	endpoint := "https://s3.harix.iamidata.com"
	accessKeyID := "5373OR9D1ZA5UD6FWE6O"
	secretAccessKey := "zuf+xPfIfXBjqMnt62dZA9c2wXXCmLVPaMUOmMBt3M6H"
	useSSL := true //是否使用https
	//endpoint := "play.min.io"
	//accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	//secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	//useSSL := true //是否使用https

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		fmt.Printf("初始化minio客户端出错:\n%s", err.Error())
		return
		//log.Fatalln(err.Error())
	}
	log.Printf("初始化minioClient成功:\n%#v\n", minioClient) // minioClient初使化成功
	//log.Printf("初始化minioClient成功:\n%+v", util.Data2Json(minioClient)) // minioClient初使化成功

	//查看所有存储桶列表
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		log.Fatalln("查看存储桶失败:\n", err)
	}

	log.Printf("所有存储桶列表:\n")
	for _, bucket := range buckets {
		log.Println(bucket)
	}

}
