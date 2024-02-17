package handlers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/api"
	"main.go/middleware"
)

func InitAPI() {
	fmt.Print("api start")
	router := gin.Default()
	// Настройка CORS middleware для разрешения всего
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"*"}
	router.Use(cors.New(config))

	apiRoutes := router.Group("/api/v1/auth")
	{
		apiRoutes.POST("/login", api.AuthLogin)
	}
	authedApiRoutes := router.Group("/api/v1").Use(middleware.AuthMiddleware())
	{
		authedApiRoutes.GET("/users", api.GetUsers)
		authedApiRoutes.POST("/users", api.CreateUser)
		authedApiRoutes.GET("/stats", api.GetStats)
	}
	router.Run(":8080")
}
