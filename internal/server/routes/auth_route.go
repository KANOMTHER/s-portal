package routes

import (
	"github.com/gin-gonic/gin"
	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func AuthRoutes(route *gin.RouterGroup, service *service.AuthService) {
	authHandler := handlers.NewAuthHandler(service)

	auth := route.Group("/auth")
	{
		auth.GET("/", authHandler.Status)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
	}
}
