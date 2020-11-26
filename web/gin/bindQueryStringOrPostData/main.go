package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
	route := gin.Default()
	//route.GET("/testing", startPage)           //使用GET
	route.POST("/testing", startPage) //使用POST
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.  如果路由是GET方法,则只使用查询字符串引擎绑定
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	//如果是POST方式, ShouldBind方法检查请求类型头Content-Type来自动选择绑定引擎,比如Json/XML
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.CreateTime)
		log.Println(person.UnixTime)
	}

	//if c.BindJSON(&person) == nil {
	//	log.Println("====== Bind By JSON ======")
	//	log.Println(person.Name)
	//	log.Println(person.Address)
	//}

	c.String(200, "Success")
}

//模拟查询字符串参数请求:
//curl -X GET "localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033"

//模拟Post Json请求
//curl -X POST localhost:8085/testing --data '{"name":"JJ", "address":"xyz"}' -H "Content-Type:application/json"
