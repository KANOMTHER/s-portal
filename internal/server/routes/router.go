package routes

import (
	"s-portal/internal/domain/service"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, services *service.Service) *gin.Engine {
	api := router.Group("/api")
	{
		ExampleRoutes(api)
		StudentRoutes(api, services.StudentService, services.AuthService)
		FacultyRoutes(api, services.FacultyService, services.AuthService)
		CourseRoutes(api, services.CourseService, services.AuthService)
		ProgramRoutes(api, services.ProgramService, services.AuthService)
		TimeTableRoutes(api, services.TimeTableService, services.AuthService)
		ClassRoutes(api, services.ClassService, services.AuthService)
		ProfessorRoutes(api, services.ProfessorService, services.AuthService)
		InstructorRoutes(api, services.InstructorService, services.AuthService)
		ClassRegisterRoutes(api, services.ClassRegisterService, services.AuthService)
		PaymentRoutes(api, services.PaymentService, services.AuthService)
		TARoute(api, services.TAService, services.AuthService)
		AuthRoutes(api, services.AuthService)
		GradeRoutes(api, services.GradeService, services.AuthService)
	}
	return router
}
