package middlewares

import (
	"net/http"

	"github.com/belf-kr/oauth-server/internal/app/models"
	"github.com/belf-kr/oauth-server/internal/pkg/auth"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 토큰 만료 여부 및 서명 검사
		if err := auth.TokenValid(c.Request); err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrResponse{
				Message: err.Error(),
			})
			c.Abort()
			return
		}
		// redis에 해당 토큰이 존재하는지 검사
		tokenAuth, err := auth.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrResponse{
				Message: "unauthorized",
			})
			c.Abort()
			return
		}
		if _, err := auth.FetchAuth(tokenAuth); err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrResponse{
				Message: "unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
