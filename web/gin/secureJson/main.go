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

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// Will output  :   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

/*
)]}',
["lena","austin","foo"]%
*/
