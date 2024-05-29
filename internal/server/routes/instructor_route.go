package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func InstructorRoutes(route *gin.RouterGroup, service *service.InstructorService, authService *service.AuthService) {
	instructorHandler := handlers.NewInstructorHandler(service, authService)

	instructor := route.Group("/instructor")
	{
		instructor.GET("", instructorHandler.GetAllInstructors)
		instructor.GET("/:id", instructorHandler.GetInstructorByID)
		instructor.POST("", instructorHandler.CreateInstructor)
		instructor.PUT("/update/:id", instructorHandler.UpdateInstructorByID)
		instructor.DELETE("/delete/:id", instructorHandler.DeleteInstructorByID)
	}
}
