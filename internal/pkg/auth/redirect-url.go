package auth

import (
	"fmt"

	"github.com/belf-kr/oauth-server/internal/app/models"
	"github.com/spf13/viper"
)

// 로그인 성공 시 토큰 정보를 queryString에 조립해줍니다. 현재는 카카오 로그인 성공 시 사용되고 있습니다.
func AuthSuccessRedUrl(ts models.JWTToken) string {
	authRedirectUrl := viper.GetString("AUTH_REDIRECT_URL")

	return fmt.Sprintf("%s?accessToken=%s&refreshToken=%s", authRedirectUrl, ts.AccessToken, ts.RefreshToken)
}
