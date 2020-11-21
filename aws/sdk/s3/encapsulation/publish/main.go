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

//封装S3结构,内部主要包含凭据,端点
type S3 struct {
	AK      string           //密钥ID
	SK      string           //密钥Key
	EP      string           //端点
	session *session.Session //会话
	Client  *s3.S3           //服务客户端service client
}

func NewClient() *S3 {
	//新建S3结构实例
	var instance = S3{
		AK: "YOUR_ACCESS_KEY_ID",
		SK: "YOUR_SECRET_ACCESS_KEY",
		EP: "s3.xxx.xxx.com",
	}

	//创建会话
	sess, _ := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(instance.AK, instance.SK, ""), //使用静态凭据,硬编码
		Endpoint:         aws.String(instance.EP),                                        //配置端点
		Region:           aws.String("default"),                                          //配置区域
		DisableSSL:       aws.Bool(false),                                                //是否禁用https,这里表示不禁用,即使用HTTPS
		S3ForcePathStyle: aws.Bool(true),                                                 //使用路径样式而非虚拟主机样式,区别请参考:https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html
	})
	//创建S3服务客户端,设置会话字段
	instance.Client = s3.New(sess)
	instance.session = sess
	return &instance
}

//绑定下载文件的方法
func (this *S3) Download(bucket, object, localFileName string) (bool, error) {
	downloader := s3manager.NewDownloader(this.session) //新建下载器
	// Create a file to write the S3 Object contents to.
	f, err := os.Create(localFileName) //创建本地下载文件localFileName
	if err != nil {
		return false, fmt.Errorf("failed to create file %q, %v", localFileName, err)
	}
	// Write the contents of S3 Object to the file
	n, err := downloader.Download(f, &s3.GetObjectInput{ //下载文件
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})
	if err != nil {
		return false, fmt.Errorf("failed to download file, %v", err)
	}
	fmt.Printf("file downloaded, %d bytes\n", n) //打印下载字节数
	return true, nil
}

func main() {
	//初始化服务客户端service client
	S3Client := NewClient()
	//S3Client.ListObjectsPages()
	S3Client.Download("bucket", "object", "localFileName") //将S3中桶名为bucket,对象为object的文件,下载为本地的localFileName
}

//为S3结构绑定方法,
func (this *S3) ListObjectsPages() {
	//查看桶下所有对象
	i := 0 //定义页码号,以分页形式列出所有对象
	err := this.Client.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: aws.String("桶名bucket"),
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		fmt.Println("Page:", i)
		i += 1
		for _, obj := range p.Contents {
			//增加过滤代码,如果
			if strings.Contains(*obj.Key, "mongo") {
				fmt.Println("找到该对象:", *obj.Key, *obj.Size)
			}
		}
		return true
	})
	if err != nil {
		fmt.Println("列出对象失败:", err)
	}
}
