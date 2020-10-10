package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"time"
)

// ...
func main() {
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

	//生成预签名URL,允许临时分享文件
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("dms"),
		Key:    aws.String("vbn-cloudpepper-v3-uit-mcs-mysql/20201009234312/vbn-cloudpepper-v3-uit-mcs-mysql.binlog-on.dump.gz"),
	})
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		log.Println("签名失败", err)
	}
	log.Println("签名成功,URL:", urlStr)
}
