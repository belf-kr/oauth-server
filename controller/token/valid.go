package token

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Valid(c *gin.Context) {
	var msg struct {
		Id       int       `json:"id"`
		Message  string    `json:"message"`
		DateTime time.Time `json:"dateTime"`
	}
	msg.Id = 1
	msg.Message = "토근 값이 정상입니다."
	msg.DateTime = time.Now()
	c.JSON(http.StatusOK, msg)
}
