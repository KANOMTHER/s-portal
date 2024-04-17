package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func StudentRoutes(route *gin.RouterGroup, service *service.StudentService) {
	studentHandler := handlers.NewStudentHandler(service)

	student := route.Group("/student")
	{
		student.POST("/", studentHandler.CreateStudent)
		student.GET("/:id", studentHandler.GetStudentByID)
		student.GET("/year", studentHandler.GetDistinctYears)
		student.GET("/year/:year", studentHandler.GetStudentsIDByYear)
		student.PUT("/update/:id", studentHandler.UpdateStudentByID)
		student.GET("/is-ta/:id", studentHandler.IsTA)
	}
}
