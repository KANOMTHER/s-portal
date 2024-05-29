package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func GradeRoutes(route *gin.RouterGroup, service *service.GradeService) {
	GradeHandler := handlers.NewGradeHandler(service)

	grade := route.Group("/grade")
	{
		grade.GET("/:id", GradeHandler.GetStudentFromClassID)
		grade.POST("/initAll", GradeHandler.InitialAll)
		grade.PUT("/update/:id", GradeHandler.EditGradeMultiple)
	}
}
