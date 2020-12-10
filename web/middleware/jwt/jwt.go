package main

import (
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	mysupersecretpassword = "unicornsAreAwesome"
)

func main() {
	r := gin.Default()

	public := r.Group("/api")

	public.GET("/", func(c *gin.Context) {
		// Create the token
		token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
		log.Printf("token:%+v", token)
		//2020/12/10 22:32:02 token:&{Raw: Method:0xc00000e2a0 Header:map[alg:HS256 typ:JWT] Claims:map[] Signature: Valid:false}

		// Set some claims
		token.Claims = jwt_lib.MapClaims{
			"Id":  "Christopher",
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		}
		// Sign and get the complete encoded token as a string
		tokenString, err := token.SignedString([]byte(mysupersecretpassword))
		//{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IkNocmlzdG9waGVyIiwiZXhwIjoxNjA3NjE0MzIyfQ.eQd7ztDn3706GrpitgnikKgOtzx-RHnq7cr2eqUlsZo"}%
		if err != nil {
			c.JSON(500, gin.H{"message": "Could not generate token"})
		}
		c.JSON(200, gin.H{"token": tokenString})
	})

	private := r.Group("/api/private")
	private.Use(jwt.Auth(mysupersecretpassword))

	/*
		Set this header in your request to get here.
		Authorization: Bearer `token`
	*/

	private.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from private"})
	})

	r.Run("localhost:8080")
}
