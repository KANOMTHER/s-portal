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
		class.GET("/:id", classHandler.GetClassByID)
		class.GET("/course", classHandler.GetClassByCourseID)
		class.GET("/semester", classHandler.GetClassBySemester)
		class.POST("/", classHandler.CreateClass)
		class.DELETE("/delete/:id", classHandler.DeleteClassByID)
	}
}
