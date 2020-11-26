package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	// ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, binding.Query)
	// ShouldBindQuery是c.ShouldBindWith(obj, binding.Query)方法的一个快捷绑定方法, 该方法只绑定请求字符串query string,而忽略Post提交的表单数据
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}

//only bind query 模拟查询字符串请求
//curl -X GET "localhost:8085/testing?name=eason&address=xyz"

//only bind query string, ignore form data 模拟查询字符串请求和Post表单,这里的表单会被忽略
//curl -X POST "localhost:8085/testing?name=eason&address=xyz" --data 'name=ignore&address=ignore' -H "Content-Type:application/x-www-form-urlencoded"
