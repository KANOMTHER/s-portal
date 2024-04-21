package routes

import (
	"s-portal/internal/domain/service"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, services *service.Service) *gin.Engine {
	api := router.Group("/api") 
	{
		ExampleRoutes(api)
		StudentRoutes(api, services.StudentService)
		FacultyRoutes(api, services.FacultyService)
		CourseRoutes(api, services.CourseService)
		ProgramRoutes(api, services.ProgramService)
		TimeTableRoutes(api, services.TimeTableService)
		ClassRoutes(api, services.ClassService)
		ProfessorRoutes(api, services.ProfessorService)
		InstructorRoutes(api, services.InstructorService)
		ClassRegisterRoutes(api, services.ClassRegisterService)
		PaymentRoutes(api, services.PaymentService)
    	TARoute(api, services.TAService)
	}
	return router
}
