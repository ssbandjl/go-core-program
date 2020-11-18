package main

//S3封装, 参考:https://docs.aws.amazon.com/sdk-for-go/api/service/s3/
//官方案例:https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code/s3
import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"

	"strings"
)

//封装S3结构
type Client struct {
	AK       string
	SK       string
	EP       string
	session  *session.Session
	S3Client *s3.S3
}

func NewClient() *Client {
	var instance = Client{
		AK: "5373OR9D1ZA5UD6FWE6O",
		SK: "xPfIfXBjqMnt62dZA9c2wXXCmLVPaMUOmMBt3M6H",
		EP: "s3.harix.iamidata.com",
	}

	sess, _ := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(instance.AK, instance.SK, ""),
		Endpoint:         aws.String(instance.EP),
		Region:           aws.String("default"),
		DisableSSL:       aws.Bool(false),
		S3ForcePathStyle: aws.Bool(true),
	})
	//log.Printf("会话:%+v", sess)
	//创建S3服务客户端
	instance.S3Client = s3.New(sess)
	instance.session = sess
	return &instance
}

func (this *Client) ListObjectsPages() {
	//查看桶下所有对象
	i := 0
	err := this.S3Client.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: aws.String("dms"),
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		fmt.Println("Page:", i)
		i += 1
		for _, obj := range p.Contents {
			//增加过滤代码
			if strings.Contains(*obj.Key, "mongo") {
				fmt.Println("找到该对象:", *obj.Key, *obj.Size)
			}
			//if *obj.Key != "vbn-cloudpepper-v3-uit-mcs-mysql/20201009234312/vbn-cloudpepper-v3-uit-mcs-mysql.binlog-on.dump.gz" {
			//	fmt.Println("找到该对象:", *obj.Key)
			//}
		}
		return true
	})
	if err != nil {
		fmt.Println("列出对象失败:", err)
	}
}

//下载文件
func (this *Client) S3CopyFile2Local(bucket, object, localFileName string) (bool, error) {
	downloader := s3manager.NewDownloader(this.session)
	// Create a file to write the S3 Object contents to.
	f, err := os.Create(localFileName)
	if err != nil {
		return false, fmt.Errorf("failed to create file %q, %v", localFileName, err)
	}
	// Write the contents of S3 Object to the file
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})
	if err != nil {
		return false, fmt.Errorf("failed to download file, %v", err)
	}
	fmt.Printf("file downloaded, %d bytes\n", n)
	return true, nil
}

func main() {
	S3Client := NewClient()
	//S3Client.ListObjectsPages()
	S3Client.S3CopyFile2Local("dms", "mysql/binlog/201117/cloud-poc-0/default/cloud-mysql/backup-mysql-master-bin.000041", "backup-mysql-master-bin.000041")
}
