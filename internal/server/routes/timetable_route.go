package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func TimeTableRoutes(route *gin.RouterGroup, service *service.TimeTableService, authService *service.AuthService) {
	timeTableHandler := handlers.NewTimeTableHandler(service, authService)

	timeTable := route.Group("/timetable")
	{
		timeTable.GET("/class", timeTableHandler.GetTimetableByClassID)
		timeTable.POST("/student", timeTableHandler.GetStudentTimetable)
		timeTable.POST("/ta", timeTableHandler.GetTATimetable)
		timeTable.GET("/count", timeTableHandler.CountTimeTable)
		timeTable.POST("", timeTableHandler.CreateTimeTable)
		timeTable.DELETE("/delete/:id", timeTableHandler.DeleteTimeTableByID)
	}
}
