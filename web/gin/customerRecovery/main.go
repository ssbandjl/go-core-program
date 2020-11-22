package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			//if r := recover(); r != nil {
			//	log.Printf("崩溃信息:%s", r)
			//}

			if err, ok := recover().(string); ok {
				log.Printf("您可以在这里完成告警任务,邮件,微信等告警")
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
			}
			c.AbortWithStatus(http.StatusInternalServerError)
		}()
		c.Next()
	}
}

func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	//r.Use(CustomRecovery())  //使用自定义中间件处理程序崩溃

	//使用匿名函数组成中间件,处理程序崩溃
	r.Use(func(c *gin.Context) {
		defer func() {
			if err, ok := recover().(string); ok {
				log.Printf("您可以在这里完成告警任务,邮件,微信等告警")
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
			}
			c.AbortWithStatus(http.StatusInternalServerError)
		}()
		c.Next()
	})

	r.GET("/panic", func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		panic("程序崩溃")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ohai")
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

//模拟程序崩溃: curl http://localhost:8080/panic
