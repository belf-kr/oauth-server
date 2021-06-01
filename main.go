package main

import (
	"fmt"
	"oauth-server/controller"
	"oauth-server/controller/token"

	"github.com/gin-gonic/gin"
)

const (
	version = "v0.1.0"
)

func init() {
	fmt.Printf("version: %s\n", version)
}

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
