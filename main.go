package main

import (
	"fmt"
	"net/http"
	"oauth-server/controller"
	"oauth-server/controller/token"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	name    = "oauth-server"
	version = "0.1.0"
)

func init() {
	fmt.Printf("version: %s\n", version)
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		resData := []byte(name)
		c.Data(http.StatusOK, "text/html; charset=utf-8", resData)
	})
	r.GET("/ping", func(c *gin.Context) {
		resData := []byte(`OK`)
		c.Data(http.StatusOK, "text/html; charset=utf-8", resData)
	})
	r.GET("/version", func(c *gin.Context) {
		resData := []byte(version)
		c.Data(http.StatusOK, "text/html; charset=utf-8", resData)
	})
	r.GET("/env", func(c *gin.Context) {
		resData := os.Environ()
		c.JSON(http.StatusOK, resData)
	})

	r.GET("/login", controller.Login)
	r.GET("/token/valid", token.Valid)

	r.Run(":3000")
}
