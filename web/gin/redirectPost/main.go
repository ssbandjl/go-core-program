package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true

	r.GET("/login", func(c *gin.Context) {
		form := "<form method='POST' action='/login/do'><input type='submit' /></form>"
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, form)
	})

	r.POST("/login/do", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/welcome")
	})

	r.GET("/welcome", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome")
	})

	r.Run(":8888")
}
