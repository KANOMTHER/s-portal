package handlers

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
)

type TimeTableHandler struct {
	timeTableService *service.TimeTableService
}

func NewTimeTableHandler(timeTableService *service.TimeTableService) *TimeTableHandler {
	return &TimeTableHandler{
		timeTableService: timeTableService,
	}
}

func (h *TimeTableHandler) CreateTimeTable(context *gin.Context) {
	timeTable := model.Timetable{}

	if err := context.ShouldBindJSON(&timeTable); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.timeTableService.CreateTimeTable(&timeTable); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return success message
	context.JSON(200, gin.H{
		"message": "Timetable created successfully",
	})
}

func (h *TimeTableHandler) CountTimeTable(context *gin.Context) {
	queryParams := context.Request.URL.Query()
	count, err := h.timeTableService.CountTimeTable(queryParams)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// // Return success message
	context.JSON(200, gin.H{
		"message": count,
	})
}

func (h *TimeTableHandler) DeleteTimeTableByID(context *gin.Context) {
	id := context.Param("id")
	err := h.timeTableService.DeleteTimeTableByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Timetable deleted successfully",
	})
}