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
	router.Use(cors.Default())

	apiRoutes := router.Group("/api/v1/auth")
	{
		apiRoutes.POST("/login", api.AuthLogin)
	}
	authedApiRoutes := router.Group("/api/v1").Use(middleware.AuthMiddleware())
	{
		authedApiRoutes.GET("/users", api.GetUsers)
		authedApiRoutes.POST("/users", api.CreateUser)
	}
	router.Run(":8080")
}
