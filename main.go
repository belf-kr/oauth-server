package main

import (
	"oauth-server/controller"
	"oauth-server/controller/token"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/login", controller.Login)
	r.GET("/token/valid", token.Valid)
	r.Run(":3000")
}
