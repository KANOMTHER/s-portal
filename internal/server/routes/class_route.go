package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func ClassRoutes(route *gin.RouterGroup, service *service.ClassService, authService *service.AuthService) {
	classHandler := handlers.NewClassHandler(service, authService)

	class := route.Group("/class")
	{
		class.GET("/:id", classHandler.GetClassByID)
		class.GET("/course", classHandler.GetClassByCourseID)
		class.GET("/semester-year", classHandler.GetClassBySemesterAndYear)
		class.POST("", classHandler.CreateClass)
		class.DELETE("/delete/:id", classHandler.DeleteClassByID)
	}
}
