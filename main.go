package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/parkgang/modern-board/configs"
	"github.com/parkgang/modern-board/docs"
	_ "github.com/parkgang/modern-board/internal/app/data/orm"
	"github.com/parkgang/modern-board/internal/app/middlewares"
	apiRouter "github.com/parkgang/modern-board/internal/app/routers"
	"github.com/parkgang/modern-board/internal/pkg/project"
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

	// SPA
	{
		const spaPath string = "../webapp/build"
		// 정적파일 응답에 헤더 추가 ("/" 경로에 헤더를 추가하기 위해서는 정적파일 서빙 미들웨어보다 위에 선언되어야 합니다)
		router.Use(middlewares.SpaResponseHeaders())
		// 정적파일 서빙
		router.Use(static.Serve("/", static.LocalFile(spaPath, true)))
		// [CSR Router를 위함](https://github.com/gin-gonic/contrib/issues/90#issuecomment-286924994)
		router.NoRoute(func(c *gin.Context) {
			// CRA에서 "homepage": "./" 와 같이 경로를 지정하면 index.html 파일 경로까지 지정해야합니다.
			c.File(spaPath)
		})
	}

	// Swagger
	{
		docs.SwaggerInfo.Title = "Modern Board API"
		docs.SwaggerInfo.Description = "게시판 데이터를 관리할 수 있는 API를 제공합니다."
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
