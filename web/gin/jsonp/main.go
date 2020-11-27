package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/JSONP", func(c *gin.Context) {
		data := gin.H{
			"foo": "bar",
		}

		//callback is x
		// Will output  :   x({\"foo\":\"bar\"})
		// 使用JSONP可以实现跨域请求数据, 如果请求中有查询字符串参数callback, 则将返回数据作为参数传递给callback值(前端函数名),整体作为一个响应体,返回给前端
		//JSONP是服务器与客户端跨源通信的常用方法。最大特点就是简单适用，老式浏览器全部支持，服务器改造非常小。
		//它的基本思想是，网页通过添加一个<script>元素，向服务器请求JSON数据，这种做法不受同源政策限制；服务器收到请求后，将数据放在一个指定名字的回调函数里传回来
		c.JSONP(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")

	// 模拟客户端,请求参数中有callback参数,值为x(前端函数名),最后响应内容为x("foo":"bar")
	// curl http://127.0.0.1:8080/JSONP?callback=x
}
