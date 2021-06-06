package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var msg struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
	msg.Id = "2d63c985-6745-9d20-7470-fe025dc2a6b0"
	msg.Name = "Kyungeun"
	c.JSON(http.StatusOK, msg)
}
