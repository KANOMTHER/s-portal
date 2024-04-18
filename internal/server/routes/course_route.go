package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func CourseRoutes(route *gin.RouterGroup, service *service.CourseService) {
	courseHandler := handlers.NewCourseHandler(service)

	course := route.Group("/course")
	{
		course.GET("/", courseHandler.GetAllCourses)
		course.GET("/:id", courseHandler.GetCourseByID)
		course.GET("/semester", courseHandler.GetAllDistinctSemester)
		course.GET("/section", courseHandler.GetSectionByClassID)
		course.POST("/", courseHandler.CreateCourse)
		course.PUT("/update/:id", courseHandler.UpdateCourseByID)
		course.DELETE("/delete/:id", courseHandler.DeleteCourseByID)

		//for class register added by ParnNoi
		course.GET("/sem-year", courseHandler.GetCourseBySemesterAndYear)
	}
}
