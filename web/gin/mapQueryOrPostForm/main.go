package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")        //查询参数中的Map
		names := c.PostFormMap("names") //Post表单中的Map

		fmt.Printf("ids: %v; names: %v\n", ids, names)
	})
	router.Run(":8080")
}
