package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"time"
)

// ...

func main() {
	//初始化日志记录器
	myLogger := MyLogger{}
	myLogger.GetLogFile()
	myLogger.GetLogger()
	defer myLogger.LogFile.Close()

	//默认区域和凭证
	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	SharedConfigState: session.SharedConfigEnable,
	//}))

	//设置区域
	//sess, err := session.NewSession(&aws.Config{
	//	Region: aws.String("us-west-2")},
	//)
	//log.Printf("session:%+v", sess)

	//初始化一个会话session对象
	access_key := "5373OR9D1ZA5UD6FWE6O"
	secret_key := "xPfIfXBjqMnt62dZA9c2wXXCmLVPaMUOmMBt3M6H"
	end_point := "s3.harix.iamidata.com"
	bucket := aws.String("dms")
	sess, _ := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(access_key, secret_key, ""),
		Endpoint:         aws.String(end_point),
		Region:           aws.String("default"),
		DisableSSL:       aws.Bool(false),
		S3ForcePathStyle: aws.Bool(true),
	})
	//log.Printf("会话:%+v", sess)

	//创建S3服务客户端
	svc := s3.New(sess)

	//查看桶下所有对象
	i := 0
	err := svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: bucket,
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		myLogger.Log.Printf("第%d页", i)
		i += 1
		for _, obj := range p.Contents {
			//增加过滤代码
			//if strings.Contains(*obj.Key, "mongo") {
			//	fmt.Println("找到该对象:", *obj.Key, *obj.Size)
			//}

			//if *obj.Key != "vbn-cloudpepper-v3-uit-mcs-mysql/20201009234312/vbn-cloudpepper-v3-uit-mcs-mysql.binlog-on.dump.gz" {
			//	fmt.Println("找到该对象:", *obj.Key)
			//}
			//log.Printf("对象:%s, 大小:%d字节, %dMB", *obj.Key, *obj.Size, *obj.Size/1024/1024)
			//myLogger.Log.Printf("对象:%s, 创建时间:%s, 大小:%d字节, %dMB", *obj.Key, *obj.LastModified, *obj.Size, *obj.Size/1024/1024)
			//删除LastModified时间, 为30天前的所有数据
			if obj.LastModified.Before(time.Now().AddDate(0, 0, -30)) {
				log.Printf("删除对象:%s, 最后编辑时间:%s, 大小:%d字节, %dMB", *obj.Key, *obj.LastModified, *obj.Size, *obj.Size/1024/1024)
				myLogger.Log.Printf("删除对象:%s, 最后编辑时间:%s, 大小:%d字节, %dMB", *obj.Key, *obj.LastModified, *obj.Size, *obj.Size/1024/1024)
				//删除
				input := &s3.DeleteObjectInput{
					Bucket: bucket,
					Key:    aws.String(*obj.Key),
				}
				_, err := svc.DeleteObject(input) //忽略result
				if err != nil {
					if aerr, ok := err.(awserr.Error); ok {
						switch aerr.Code() {
						default:
							//Log4Zap(zap.InfoLevel).Info(fmt.Sprintf("删除对象失败,桶:%s, 对象:%s, 错误:%s", bucket, object, aerr.Error()))
							log.Printf("删除失败, 对象%s, 错误:%s", *obj.Key, aerr.Error())
							myLogger.Log.Printf("删除失败, 对象%s, 错误:%s", *obj.Key, aerr.Error())
							continue
						}
					} else {
						// Print the error, cast err to awserr.Error to get the Code and
						// Message from an error.
						//Log4Zap(zap.InfoLevel).Info(fmt.Sprintf("删除对象失败,桶:%s, 对象:%s, 错误:%s", bucket, object, err.Error()))
						log.Printf("删除失败, 对象%s, 错误:%s", *obj.Key, err.Error())
						myLogger.Log.Printf("删除失败, 对象%s, 错误:%s", *obj.Key, err.Error())
						continue
					}
				}

			}
		}
		return true
	})
	if err != nil {
		fmt.Println("列出对象失败:", err)
	}
}
