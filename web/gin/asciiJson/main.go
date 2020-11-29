package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出结果 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		// AsciiJSON方法返回带有Unicode编码和转义组成的纯ASCII字符串
		c.AsciiJSON(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

/*
模拟请求:curl http://localhost:8080/someJSON
*/
