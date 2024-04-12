package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func TARoute(route *gin.RouterGroup, service *service.TAService) {
	taHandler := handlers.NewTAHandler(service)

	students := route.Group("/ta")
	{
		students.GET("/yo", taHandler.GetHello)
		students.POST("/create", taHandler.CreateClassTA)
		students.PUT("/update", taHandler.UpdateClassTA)
	}
}