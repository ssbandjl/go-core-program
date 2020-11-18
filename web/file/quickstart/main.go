package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	router.GET("/api/v1/attachments/:file", DownloadAttachmentHandler)
	router.POST("/api/v1/attachment/files", UploadAttachmentFileHandler)

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

// UploadAttachmentFileHandler ...
func UploadAttachmentFileHandler(c *gin.Context) {
	//文件后缀
	fileSuffix := ""
	//文件路径
	fileDir := "./data" //文件上传保存的路径eg:./data
	//文件全名
	fileName := ""
	formdata, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, HTTPGenericResponse{
			Code:    http.StatusInternalServerError,
			Message: "文件读取失败:" + err.Error(),
		})
		return
	}
	fileHeaders := formdata.File["file"]
	attachs := ""
	name := ""
	for index, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusOK, HTTPGenericResponse{
				Code:    http.StatusInternalServerError,
				Message: "文件读取失败:" + err.Error(),
			})
			return
		}
		defer file.Close()

		fileSuffix = path.Ext(fileHeader.Filename)
		//文件全名
		fileName = fileSuffix

		if _, err = os.Stat(fileDir); os.IsNotExist(err) {
			err = os.MkdirAll(fileDir, FileMode)
			if err != nil {
				c.JSON(http.StatusOK, HTTPGenericResponse{
					Code:    http.StatusInternalServerError,
					Message: "文件创建失败:" + err.Error(),
				})
				return
			}
		}

		fW, err := os.Create(filepath.Join(fileDir, fileName))
		if err != nil {
			c.JSON(http.StatusOK, HTTPGenericResponse{
				Code:    http.StatusInternalServerError,
				Message: "文件创建失败:" + err.Error(),
			})
			return
		}
		defer fW.Close()

		_, err = io.Copy(fW, file)
		if err != nil {
			c.JSON(http.StatusOK, HTTPGenericResponse{
				Code:    http.StatusInternalServerError,
				Message: "文件保存失败:" + err.Error(),
			})
			return
		}

	}
	c.JSON(http.StatusOK, HTTPGenericResponse{
		Code:    http.StatusOk,
		Message: "文件保存失败:" + err.Error(),
	})
	return
}

// DownloadAttachmentHandler ...
func DownloadAttachmentHandler(c *gin.Context) {
	var request api.DownloadAttachmentRequest
	filePath = c.Param("file")
	file, err := os.Open(filePath) //Create a file
	if err != nil {
		c.JSON(http.StatusNotFound, HTTPGenericResponse{
			Code:    http.StatusInternalServerError,
			Message: "文件加载失败:" + err.Error(),
		})
		return
	}
	defer file.Close()
	c.Writer.Header().Add("Content-type", "application/octet-stream")
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.JSON(http.StatusNotFound, HTTPGenericResponse{
			Code:    http.StatusInternalServerError,
			Message: "文件加载失败:" + err.Error(),
		})
		return
	}
}
