package main

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type ProfileForm struct {
	Name   string                `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`

	// or for multiple files
	// Avatars []*multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/profile", func(c *gin.Context) {
		// you can bind multipart form with explicit binding declaration:  可以使用显示申明的方式,即用ShouldBindWith(&from, binding.Form)方法来绑定多部分类型表单multipart form
		// c.ShouldBindWith(&form, binding.Form)
		// or you can simply use autobinding with ShouldBind method:
		var form ProfileForm
		// in this case proper binding will be automatically selected
		// 这里使用ShouldBind方法自动选择绑定器进行绑定
		if err := c.ShouldBind(&form); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		//保存上传的表单文件到指定的目标文件
		err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
		if err != nil {
			c.String(http.StatusInternalServerError, "unknown error")
			return
		}
		// db.Save(&form)
		c.String(http.StatusOK, "ok")
	})
	router.Run(":8080")
}

//模拟测试:
//curl -X POST -v --form name=user --form "avatar=@./avatar.png" http://localhost:8080/profile
