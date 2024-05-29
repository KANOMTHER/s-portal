package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func ProfessorRoutes(route *gin.RouterGroup, service *service.ProfessorService, authService *service.AuthService) {
	professorHandler := handlers.NewProfessorHandler(service, authService)

	professor := route.Group("/professor")
	{
		professor.GET("", professorHandler.GetAllProfessors)
		professor.GET("/:id", professorHandler.GetProfessorByID)
		professor.POST("", professorHandler.CreateProfessor)
		professor.PUT("/update/:id", professorHandler.UpdateProfessorByID)
		professor.DELETE("/delete/:id", professorHandler.DeleteProfessorByID)
	}
}
