package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func FacultyRoutes(route *gin.RouterGroup, service *service.FacultyService) {
	facultyHandler := handlers.NewFacultyHandler(service)

	faculty := route.Group("/faculty")
	{
		faculty.GET("/", facultyHandler.GetAllFaculties)
		faculty.GET("/:id", facultyHandler.GetFacultyByID)
		faculty.POST("/", facultyHandler.CreateFaculty)
		faculty.PUT("/update/:id", facultyHandler.UpdateFacultyByID)
		faculty.DELETE("/delete/:id", facultyHandler.DeleteFacultyByID)
	}
}
