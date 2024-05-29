package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func InitRoute(route *gin.RouterGroup, service *service.InitService) {
	initHandler := handlers.NewInitHandler(service)

	students := route.Group("/init")
	{
		students.GET("/go", initHandler.GoInit)
	}
}