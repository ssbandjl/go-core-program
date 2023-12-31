package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()

	// gin.H is a shortcut for map[string]interface{}
	// gin.H对象是一个map映射,键名为字符串类型, 键值是接口,所以可以传递所有的类型
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}

		//JSON serializes the given struct as JSON into the response body. It also sets the Content-Type as "application/json".
		//JSON方法将给定的结构序列化为JSON到响应体, 并设置内容类型Content-Type为:"application/json"
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	//Protocol buffers are a language-neutral, platform-neutral extensible mechanism for serializing structured data.
	//Protocol buffers(简称ProtoBuf)是来自Google的一个跨语言,跨平台,用于将结构化数据序列化的可扩展机制,
	//详见:https://developers.google.com/protocol-buffers
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// The specific definition of protobuf is written in the testdata/protoexample file.
		// 使用protoexample.Test这个特别的protobuf结构来定义测试数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// Note that data becomes binary data in the response  //将data序列化为二进制的响应数据
		// Will output protoexample.Test protobuf serialized data
		// ProtoBuf serializes the given struct as ProtoBuf into the response body.
		// ProtoBuf方法将给定的结构序列化为ProtoBuf响应体
		c.ProtoBuf(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

/*
模拟测试
curl http://localhost:8080/someJSON
{"message":"hey","status":200}

curl http://localhost:8080/moreJSON
{"user":"Lena","Message":"hey","Number":123}

curl http://localhost:8080/someXML
<map><message>hey</message><status>200</status></map>

curl http://localhost:8080/someYAML
message: hey
status: 200

curl http://localhost:8080/someProtoBuf
test
*/
