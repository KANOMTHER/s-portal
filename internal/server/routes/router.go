package routes

import (
	"s-portal/internal/domain/service"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, services *service.Service) *gin.Engine {
	api := router.Group("/api") 
	{
		ExampleRoutes(api)
		StudentRoute(api)
		FacultyRoutes(api, services.FacultyService)
		CourseRoutes(api, services.CourseService)
		ProgramRoutes(api, services.ProgramService)
		TARoute(api, services.TAService)
	}
	return router
}
