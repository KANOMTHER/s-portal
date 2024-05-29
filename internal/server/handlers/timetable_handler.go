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

// @Summary		GetTimetableByClassID
// @Description	get timetable search by class_id
// @Tags			Timetable
// @Accept			json
// @Produce		json
// @Param			class_id	query		string	true	"class id"
// @Success		200			{array}		service.GetTimetableByClassIDField
// @Failure		404			{object}	string	"No timetables found"
// @Router			/timetable/class [GET]
func (h *TimeTableHandler) GetTimetableByClassID(context *gin.Context) {
	class_id := context.Query("class_id")
	timetables, err := h.timeTableService.GetTimetableByClassID(class_id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": "No timetables found",
		})
		return
	}
	// Return timetables
	context.JSON(200, gin.H{
		"message": timetables,
	})
}

func (h *TimeTableHandler) GetStudentTimetable(context *gin.Context) {
	timetable, err := h.timeTableService.GetStudentTimetable(context)
	if err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return success message
	context.JSON(200, gin.H{
		"message": timetable,
	})
}

func (h *TimeTableHandler) GetTATimetable(context *gin.Context) {
	timetables, err := h.timeTableService.GetTATimetable(context)
	
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message err": err.Error(),
		})
		return;
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": timetables,
	})
}

// @Summary		CreateTimeTable
// @Description	create a new timetable
// @Tags			Timetable
// @Accept			json
// @Produce		json
// @Param			timetable	body		model.Timetable	true	"Timetable object"
// @Success		200			{object}	string			"Timetable created successfully"
// @Failure		400			{object}	string			"some error message here (from err.Error())"
// @Router			/timetable [POST]
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

// @Summary		CountTimeTable
// @Description	get timetable search by class_id
// @Tags			Timetable
// @Accept			json
// @Produce		json
// @Param			class_id	query		string	false	"filter by class_id (original from php version)"	example:"1"
// @Param			Day			query		string	false	"filter by day"										example:"0"
// @Param			StartTime	query		string	false	"filter by StartTime"								example:"2021-08-01T08:00:00Z"
// @Param			EndTime		query		string	false	"filter by EndTime"									example:"2021-08-01T09:00:00Z"
// @Param			Classroom	query		string	false	"filter by Classroom"								example:"CPE1102"
// @Param			Classtype	query		string	false	"filter by ClassType"								example:"Lecture"
// @Success		200			{object}	int64	"count of timetables"
// @Failure		404			{object}	string	"some error message here (from err.Error())"
// @Router			/timetable/count [GET]
func (h *TimeTableHandler) CountTimeTable(context *gin.Context) {
	queryParams := context.Request.URL.Query()
	count, err := h.timeTableService.CountTimeTable(queryParams)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// // Return success message
	context.JSON(200, gin.H{
		"message": count,
	})
}

// @Summary		DeleteTimeTableByID
// @Description	delete a timetable by id
// @Tags			Timetable
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"timetable id"
// @Success		200	{object}	string	"Timetable deleted successfully"
// @Failure		404	{object}	string	"some error message here (from err.Error())"
// @Router			/timetable/delete/{id} [DELETE]
func (h *TimeTableHandler) DeleteTimeTableByID(context *gin.Context) {
	id := context.Param("id")
	err := h.timeTableService.DeleteTimeTableByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Timetable deleted successfully",
	})
}
