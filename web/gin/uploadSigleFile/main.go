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
		// single file
		file, _ := c.FormFile("file")           //FormFile从表单中返回第一个匹配到的文件对象(结构)
		log.Printf("获取到的文件名:%s", file.Filename) //文件名必须是安全可信耐的,需要去掉路径信息,保留文件名即可

		// Upload the file to specific dst.
		currentPath, _ := os.Getwd() //获取当前文件路径
		dst := currentPath + "/" + file.Filename
		log.Printf("保存文件绝对路径:%s", dst)
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}

//模拟单文件上传:
//curl -X POST http://localhost:8080/upload  -H "Content-Type: multipart/form-data" -F "file=@pycharm-professional-2020.1.4.dmg"
