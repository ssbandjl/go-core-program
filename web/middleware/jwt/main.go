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
	private.Use(Auth(mysupersecretpassword)) // 使用JWT认证中间件(带参数)

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
