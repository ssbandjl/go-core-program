package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// You can also use your own secure json prefix
	// 你也可以自定义安全Json的前缀
	r.SecureJsonPrefix(")]}',\n")

	//使用SecureJSON方法保护Json不被劫持, 如果响应体是一个数组, 该方法会默认添加`while(1)`前缀到响应头,  这样的死循环可以防止后面的代码被恶意执行, 也可以自定义安全JSON的前缀.
	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		//names := map[string]string{
		//	"hello": "world",
		//}

		// Will output  :   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

/*
模拟请求:curl http://localhost:8080/someJSON
)]}',
["lena","austin","foo"]%
*/
