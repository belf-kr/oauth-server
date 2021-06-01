package token

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Valid(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Lena"
	c.JSON(http.StatusOK, msg)
}
