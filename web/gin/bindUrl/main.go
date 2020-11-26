package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"` //指定URI标签
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	//下面的URI中的name和id与Person结构中的标签分别对应
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run(":8088")
}

//模拟请求
//curl -v localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
//curl -v localhost:8088/thinkerou/not-uuid
