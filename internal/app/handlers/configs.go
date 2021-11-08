package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parkgang/modern-board/internal/app/models"
	"github.com/spf13/viper"
)

// @Summary 구성 조회
// @Description client가 앱을 실행하기 위해 필요한 구성을 조회합니다.
// @Tags App
// @Accept json
// @Produce json
// @Success 200 {object} models.Config
// @Router /configs [get]
func GetConfigs(c *gin.Context) {
	restApiKey := viper.GetString("kakao.restApiKey")
	redirectUri := viper.GetString("kakao.redirectUri")

	c.JSON(http.StatusOK, models.Config{
		Kakao: models.Kakao{
			RestApiKey:  restApiKey,
			RedirectUri: redirectUri,
		},
	})
}
