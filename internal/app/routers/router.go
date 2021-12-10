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
	api.GET("/env", handlers.AppEnv)
	users := api.Group("/users")
	{
		users.POST("/signup", handlers.UserSignup)
		users.POST("/login", handlers.UserLogin)
		users.GET("/login/kakao", handlers.UserKakaoLoginCallBack)
		users.POST("/logout", middlewares.TokenAuthMiddleware(), handlers.UserLogout)

		users.GET("/token/valid", middlewares.TokenAuthMiddleware(), handlers.UserTokenValid)
		users.POST("/token/refresh", handlers.UserTokenRefresh)

		users.GET("", middlewares.TokenAuthMiddleware(), handlers.UserInfoTokenQuey)
		users.GET(":userEmail", handlers.UserInfoEmailQuey)
		users.POST("/avatar", middlewares.TokenAuthMiddleware(), handlers.UploadAvatar)
	}
	configs := api.Group("/configs")
	{
		configs.GET("", handlers.GetConfigs)
	}
}
