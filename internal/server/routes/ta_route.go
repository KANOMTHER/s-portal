package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func TARoute(route *gin.RouterGroup, service *service.TAService, authService *service.AuthService) {
	taHandler := handlers.NewTAHandler(service, authService)

	students := route.Group("/ta")
	{
		students.GET("/get", taHandler.GetClassTA)
		students.GET("/get/:id", taHandler.GetClassTAByClassID)
		students.POST("/create", taHandler.CreateClassTA)
		students.PUT("/update", taHandler.UpdateClassTA)
		students.DELETE("/delete/:id", taHandler.DeleteClassTA)
	}
}
