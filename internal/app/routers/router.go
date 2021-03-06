package routers

import (
	"github.com/belf-kr/oauth-server/internal/app/handlers"
	"github.com/belf-kr/oauth-server/internal/app/middlewares"
	"github.com/gin-gonic/gin"
)

func Use(api *gin.RouterGroup) {
	api.GET("/", handlers.AppName)
	api.GET("/ping", handlers.Ping)
	api.GET("/version", handlers.AppVersion)

	users := api.Group("/users")
	{
		users.POST("/login", handlers.UserLogin)
		users.GET("/login/kakao", handlers.UserKakaoLoginCallBack)
		users.POST("/logout", middlewares.TokenAuthMiddleware(), handlers.UserLogout)

		users.GET("/token/valid", middlewares.TokenAuthMiddleware(), handlers.UserTokenValid)
		users.POST("/token/refresh", handlers.UserTokenRefresh)

		users.POST("", handlers.UserSignup)
		users.GET("", middlewares.TokenAuthMiddleware(), handlers.UserInfoTokenQuey)
		users.DELETE("", middlewares.TokenAuthMiddleware(), handlers.UserWithdrawal)
		users.GET(":userKey", handlers.UserInfoQuey)
		users.POST("/avatar", middlewares.TokenAuthMiddleware(), handlers.UploadAvatar)
		users.DELETE("/avatar", middlewares.TokenAuthMiddleware(), handlers.DeleteAvatar)
	}

	configs := api.Group("/configs")
	{
		configs.GET("", handlers.GetConfigs)
	}
}
