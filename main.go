package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/belf-kr/oauth-server/configs"
	"github.com/belf-kr/oauth-server/docs"
	_ "github.com/belf-kr/oauth-server/internal/app/data/orm"
	apiRouter "github.com/belf-kr/oauth-server/internal/app/routers"
	"github.com/belf-kr/oauth-server/internal/pkg/project"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	fmt.Printf("version: %s\n", project.AppVersion)
}

func main() {
	port := viper.GetString("server.port")

	router := gin.Default()

	// CORS: 원래는 "router.Use(cors.Default())" 으로 대부분 해결이 되었는데 사용자 인증을 위해서 "Authorization" header를 사용해야하는데 Default가 그부분 까지는 코딩되어 있지 않아 아래와 같이 커스텀 합니다.
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	// API
	api := router.Group("/api")
	{
		apiRouter.Use(api)
	}

	// Swagger
	{
		docs.SwaggerInfo.Title = "OAuth Server API"
		docs.SwaggerInfo.Description = "사용자 정보를 다룰 수 있는 API를 제공하며 JWT 방법으로 로그인할 수 있는 기능을 제공합니다."
		docs.SwaggerInfo.Version = project.AppVersion
		docs.SwaggerInfo.Host = fmt.Sprintf("localhost%s", port)
		docs.SwaggerInfo.BasePath = "/api"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}

		url := ginSwagger.URL(fmt.Sprintf("http://localhost%s/swagger/doc.json", port))
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	log.Printf("gin server listening at http://localhost%s\n", port)
	if err := router.Run(port); err != nil {
		log.Fatal("gin server에서 예상하지 못한 에러가 발생하였습니다.\n\t" + err.Error())
	}
	log.Println("gin server close", port)
}
