package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// 模拟提交表单:curl -XPOST http://localhost:8080/form_post -d "message=消息&nick=昵称"
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}
