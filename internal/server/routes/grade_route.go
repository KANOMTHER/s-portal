package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func GradeRoutes(route *gin.RouterGroup, service *service.GradeService, authService *service.AuthService) {
	GradeHandler := handlers.NewGradeHandler(service, authService)

	grade := route.Group("/grade")
	{
		grade.GET("/:id", GradeHandler.GetStudentFromClassID)
		grade.POST("/initAll", GradeHandler.InitialAll)
		grade.PUT("/update/:id", GradeHandler.EditGradeMultiple)
	}
}
