package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/belf-kr/oauth-server/internal/pkg/project"
)

// @Summary 앱 이름
// @Description 앱 이름을 응답합니다.
// @Tags App
// @Success 200 {string} string	"oauth-server"
// @Router / [get]
func AppName(c *gin.Context) {
	resData := []byte(project.AppName)
	c.Data(http.StatusOK, "text/html; charset=utf-8", resData)
}

// @Summary server 헬스 체크
// @Description server의 헬스를 체크합니다.
// @Tags App
// @Success 200
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.Status(http.StatusOK)
}

// @Summary 앱 버전
// @Description 앱 버전을 응답합니다.
// @Tags App
// @Success 200 {string} string	"0.1.0"
// @Router /version [get]
func AppVersion(c *gin.Context) {
	resData := []byte(project.AppVersion)
	c.Data(http.StatusOK, "text/html; charset=utf-8", resData)
}

// @Summary 앱 환경 변수
// @Description 앱의 환경변수를 응답합니다.
// @Tags App
// @Produce json
// @Success 200 {array} []string
// @Router /env [get]
func AppEnv(c *gin.Context) {
	resData := os.Environ()
	c.JSON(http.StatusOK, resData)
}
