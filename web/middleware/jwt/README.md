适用于gin-gonic框架的JWT中间件.

JSON Web Token (JWT) 详见: http://self-issued.info/docs/draft-ietf-oauth-json-web-token.html

EDIT: 下面是测试代码, 详见: [christopherL91/Go-API](https://github.com/christopherL91/Go-API)

```go
package jwt_test

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Token string `json:"token"`
}

func createNewsUser(username, password string) *User {
	return &User{username, password}
}

func TestLogin(t *testing.T) {
	Convey("Should be able to login", t, func() {
		user := createNewsUser("jonas", "1234")
		jsondata, _ := json.Marshal(user)
		post_data := strings.NewReader(string(jsondata))
		req, _ := http.NewRequest("POST", "http://localhost:3000/api/login", post_data)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		res, _ := client.Do(req)
		So(res.StatusCode, ShouldEqual, 200)

		Convey("Should be able to parse body", func() {
			body, err := ioutil.ReadAll(res.Body)
			defer res.Body.Close()
			So(err, ShouldBeNil)
			Convey("Should be able to get json back", func() {
				responseData := new(Response)
				err := json.Unmarshal(body, responseData)
				So(err, ShouldBeNil)

				Convey("Should be able to be authorized", func() {
					token := responseData.Token
					req, _ := http.NewRequest("GET", "http://localhost:3000/api/auth/testAuth", nil)
					req.Header.Set("Authorization", "Bearer "+token)
					client = &http.Client{}
					res, _ := client.Do(req)
					So(res.StatusCode, ShouldEqual, 200)
				})
			})
		})
	})
	Convey("Should not be able to login with false credentials", t, func() {
		user := createNewsUser("jnwfkjnkfneknvjwenv", "wenknfkwnfknfknkfjnwkfenw")
		jsondata, _ := json.Marshal(user)
		post_data := strings.NewReader(string(jsondata))
		req, _ := http.NewRequest("POST", "http://localhost:3000/api/login", post_data)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		res, _ := client.Do(req)
		So(res.StatusCode, ShouldEqual, 401)
	})

	Convey("Should not be able to authorize with false credentials", t, func() {
		token := ""
		req, _ := http.NewRequest("GET", "http://localhost:3000/api/auth/testAuth", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		client := &http.Client{}
		res, _ := client.Do(req)
		So(res.StatusCode, ShouldEqual, 401)
	})
}
```



# 测试步骤

- 运行gin服务器: go run main.go
- 运行单元测试: go test



```
package jwt

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// JWT认证中间件
func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			c.AbortWithError(401, err)
		}
	}
}

```

解析请求参数中的JWT

```
// Extract and parse a JWT token from an HTTP request. 从HTTP请求中解析JWT令牌
// This behaves the same as Parse, but accepts a request and an extractor
// instead of a token string.  The Extractor interface allows you to define
// the logic for extracting a token.  Several useful implementations are provided.
//
// You can provide options to modify parsing behavior
func ParseFromRequest(req *http.Request, extractor Extractor, keyFunc jwt.Keyfunc, options ...ParseFromRequestOption) (token *jwt.Token, err error) {
	// Create basic parser struct
	p := &fromRequestParser{req, extractor, nil, nil}

	// Handle options
	for _, option := range options {
		option(p)
	}

	// Set defaults
	if p.claims == nil {
		p.claims = jwt.MapClaims{}
	}
	if p.parser == nil {
		p.parser = &jwt.Parser{}
	}

	// perform extract
	tokenString, err := p.extractor.ExtractToken(req)
	if err != nil {
		return nil, err
	}

	// perform parse
	return p.parser.ParseWithClaims(tokenString, p.claims, keyFunc)
}

```



# 前后端分离必备, Golang Gin中如何使用JWT(JsonWebToken)中间件?



## 什么是JWT?

JSON Web Token（缩写 JWT）是目前最流行的跨域认证解决方案，也是目前前后端分离项目中普遍使用的认证技术. 本文介绍如何在Golang Gin Web框架中使用JWT认证中间件以及模拟测试, 以供参考, 关于JWT详细原理可以参考:

- JWT RFC: https://tools.ietf.org/html/rfc7519
- JWT IETF: http://self-issued.info/docs/draft-ietf-oauth-json-web-token.html
- JSON Web Token入门教程: http://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html

## 主要流程

- 初始化Gin引擎
- 定义获取Token的接口, 访问该接口, 内部自动生成JWT令牌, 并返回给前端
- 定义需要认证的路由接口, 使用JWT中间件进行认证, 中间件由
- 利用GoConvey(Golang的测试框架,集成go test, 支持终端和浏览器模式), 构造客户端, 填写Token, 模拟前端访问
- JWT中间件进行认证, 认证通过则返回消息体, 否则直接返回401或其他错误

## 流程图

![image-20201212151536687](/Users/xb/Library/Application Support/typora-user-images/image-20201212151536687.png)

该流程图描述了服务端代码中的Token构造, 以及认证流程.

## 服务端代码

```
package main

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var (
	mysupersecretpassword = "unicornsAreAwesome"
)



func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//log.Printf("Request:\n%+v", c.Request)
		// ParseFromRequest方法提取路径请求中的JWT令牌, 并进行验证
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			//log.Printf("b:%+v", b)
			return b, nil
		})

		log.Printf("token:%+v", token)
		if err != nil {
			c.AbortWithError(401, err)
		}
	}
}


func main() {
	r := gin.Default()

	public := r.Group("/api")

	// 定义根路由, 访问http://locahost:8080/api/可以获取到token
	public.GET("/", func(c *gin.Context) {
		// Create the token New方法接受一个签名方法的接口类型(SigningMethod)参数, 返回一个Token结构指针
		// GetSigningMethod(签名算法algorithm)
		token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256")) //默认是签名算法是HMAC SHA256（写成 HS256）
		log.Printf("token:%+v", token)
		//2020/12/10 22:32:02 token:&{Raw: Method:0xc00000e2a0 Header:map[alg:HS256 typ:JWT] Claims:map[] Signature: Valid:false}

		// Set some claims 设置Id和过期时间字段, MapClaims实现了Clainms接口
		token.Claims = jwt_lib.MapClaims{
			"Id":  "Christopher",
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		}
		// Sign and get the complete encoded token as a string // 签名并得到完整编码后的Token字符串
		tokenString, err := token.SignedString([]byte(mysupersecretpassword))
		//{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IkNocmlzdG9waGVyIiwiZXhwIjoxNjA3NjE0MzIyfQ.eQd7ztDn3706GrpitgnikKgOtzx-RHnq7cr2eqUlsZo"}
		if err != nil {
			c.JSON(500, gin.H{"message": "Could not generate token"})
		}
		c.JSON(200, gin.H{"token": tokenString})
	})


	// 定义需要Token验证通过才能访问的私有接口组http://localhost:8080/api/private
	private := r.Group("/api/private")
	private.Use(Auth(mysupersecretpassword))  // 使用JWT认证中间件(带参数)

	/*
		Set this header in your request to get here.
		Authorization: Bearer `token`
	*/

	// 定义具体的私有根接口:http://localhost:8080/api/private/
	private.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from private"})
	})

	r.Run("localhost:8080")
}
```



## 客户端代码

```
package test_test

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"  //https://github.com/smartystreets/goconvey GoConvey是Golang的测试框架,集成go test, 支持终端和浏览器模式.
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Token string `json:"token"`
}

func createNewsUser(username, password string) *User {
	return &User{username, password}
}

func TestLogin(t *testing.T) {
	Convey("Should be able to login", t, func() {
		user := createNewsUser("jonas", "1234")
		jsondata, _ := json.Marshal(user)
		userData := strings.NewReader(string(jsondata))
		log.Printf("userData:%+v", userData)
		// 这里模拟用户登录, 实际上后台没有使用用户名和密码, 该接口直接返回内部生成的Token
		req, _ := http.NewRequest("GET", "http://localhost:8080/api/", userData)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		res, _ := client.Do(req)
		//log.Printf("res:%+v", res)
		So(res.StatusCode, ShouldEqual, 200) //对响应码进行断言, 期望得到状态码为200

		Convey("Should be able to parse body", func() { //解析响应体
			body, err := ioutil.ReadAll(res.Body)
			defer res.Body.Close()
			So(err, ShouldBeNil)
			Convey("Should be able to get json back", func() {
				responseData := new(Response)
				err := json.Unmarshal(body, responseData)
				So(err, ShouldBeNil)
				log.Printf("responseData:%s", responseData)
				Convey("Should be able to be authorized", func() {
					token := responseData.Token //提取Token
					log.Printf("token:%s", token)
					// 构造带Token的请求
					req, _ := http.NewRequest("GET", "http://localhost:8080/api/private", nil)
					req.Header.Set("Authorization", "Bearer "+token) //设置认证头
					client = &http.Client{}
					res, _ := client.Do(req)
					body, err := ioutil.ReadAll(res.Body)
					if err != nil {
						log.Printf("Read body failed, %s", err.Error())
					}
					log.Printf("Body:%s", string(body))
					So(res.StatusCode, ShouldEqual, 200)
				})
			})
		})
	})
}
```



## 参考文档

gin-gonic/contrib/jwt中间件: https://github.com/gin-gonic/contrib/tree/master/jwt