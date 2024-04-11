package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func ClassRoutes(route *gin.RouterGroup, service *service.ClassService) {
	classHandler := handlers.NewClassHandler(service)

	class := route.Group("/class")
	{
		class.POST("/", classHandler.CreateClass)
	}
}
