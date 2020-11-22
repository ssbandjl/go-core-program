package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// 设置请求表单最大内存限制,默认是30MB
	//内部调用http请求的ParseMultipartForm方法,该方法要求传入一个字节数, 要取MultipartForm字段的数据，先使用ParseMultipartForm()方法解析Form，解析时会读取所有数据，但需要指定保存在内存中的最大字节数，剩余的字节数会保存在临时磁盘文件中
	maxMultipartMemory := int64(8 << 20)
	log.Printf("解析文件到内存的最大字节:%d", maxMultipartMemory)
	router.MaxMultipartMemory = maxMultipartMemory // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Upload the file to specific dst.
		currentPath, _ := os.Getwd() //获取当前文件路径
		// Multipart form
		form, _ := c.MultipartForm()   //多文件表单
		files := form.File["upload[]"] //通过前端提供的键名获取文件数组
		for _, file := range files {
			dst := currentPath + "/" + file.Filename
			log.Printf("保存文件绝对路径:%s", dst)
			// Upload the file to specific dst.
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8080")
}

//模拟多文件上传
//curl -X POST http://localhost:8080/upload -H "Content-Type: multipart/form-data" -F "upload[]=@xinfracloud179.sql" -F "upload[]=@web-k8s-exec-master.zip"
