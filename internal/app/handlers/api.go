package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/belf-kr/oauth-server/internal/app/models"
	"github.com/belf-kr/oauth-server/internal/pkg/project"
)

func AppName(c *gin.Context) {
	resData := []byte(project.AppName)
	c.Data(http.StatusOK, "text/html; charset=utf-8", resData)
}

// @Summary Server Health Check
// @Description gin server의 헬스를 체크합니다.
// @Produce json
// @Success 200 {object} models.Pong
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, models.Pong{
		Message: "pong",
	})
}

func AppVersion(c *gin.Context) {
	resData := []byte(project.AppVersion)
	c.Data(http.StatusOK, "text/html; charset=utf-8", resData)
}

func AppEnv(c *gin.Context) {
	resData := os.Environ()
	c.JSON(http.StatusOK, resData)
}
