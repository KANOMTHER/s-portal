package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func ClassRegisterRoutes(route *gin.RouterGroup, service *service.ClassRegisterService) {
	classRegisterHandler := handlers.NewClassRegisterHandler(service)

	class := route.Group("/register")
	{
		class.POST("/get", classRegisterHandler.GetRegisterClassByID)
		class.POST("/create", classRegisterHandler.CreateRegisterClass)
		class.PUT("/update", classRegisterHandler.UpdateRegisterClass)
		class.DELETE("/delete", classRegisterHandler.DeleteRegisterClass)
	}
}
