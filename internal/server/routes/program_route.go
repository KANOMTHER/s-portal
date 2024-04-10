package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func ProgramRoutes(route *gin.RouterGroup, service *service.ProgramService) {
	programHandler := handlers.NewProgramHandler(service)

	program := route.Group("/program")
	{
		program.GET("/", programHandler.GetAllPrograms)
	}
}
