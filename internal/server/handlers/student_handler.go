package handlers

import (

	"github.com/gin-gonic/gin"

	// "s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
)

type StudentHandler struct {
	studentService *service.StudentService
}

func NewStudentHandler(studentService *service.StudentService) *StudentHandler {
	return &StudentHandler{
		studentService: studentService,
	}
}

//	@Summary		CreateStudent
//	@Description	create a new student
//	@Tags			Student
//	@Accept			json
//	@Produce		json
//	@Param			CreateStudentFields	body		service.CreateStudentFields	true	"CreateStudentFields object"
//	@Success		200					{object}	string						"Student created successfully"
//	@Failure		400					{object}	string						"some error message here (from err.Error())"
//	@Router			/student [POST]
func (h *StudentHandler) CreateStudent(context *gin.Context) {
	student := service.CreateStudentFields{}

	if err := context.ShouldBindJSON(&student); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := h.studentService.CreateStudent(&student)
	if err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return success message
	context.JSON(200, gin.H{
		"message": "Student created successfully",
	})
}

//	@Summary		GetDistinctYears
//	@Description	get a distinct year of student
//	@Tags			Student
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		uint	"Array of distinct year in dedscending order"
//	@Failure		404	{object}	string	"some error message here (from err.Error())"
//	@Router			/student/year [GET]
func (h *StudentHandler) GetDistinctYears(context *gin.Context) {
	distinct_year, err := h.studentService.GetDistinctYears()
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": distinct_year,
	})
}

//	@Summary		GetStudentsByYear
//	@Description	get a student by year
//	@Tags			Student
//	@Accept			json
//	@Produce		json
//	@Param			year	query		string	true	"64, 65, 66, 67, ..."
//	@Success		200		{array}		uint	"Array of student's year [64, 65, 66, 67, ...]"
//	@Failure		404		{object}	string	"some error message here (from err.Error())"
//	@Router			/student/getByYear [GET]
func (h *StudentHandler) GetStudentsIDByYear(context *gin.Context) {
	year := context.Query("year")
	distinct_year, err := h.studentService.GetStudentsIDByYear(year)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": distinct_year,
	})
}

//	@Summary		GetStudentByID
//	@Description	get a student by id
//	@Tags			Student
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"student id"
//	@Success		200	{object}	model.Student
//	@Failure		404	{object}	string
//	@Router			/student/{id} [GET]
func (h *StudentHandler) GetStudentByID(context *gin.Context) {
	id := context.Param("id")
	student, err := h.studentService.GetStudentByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": student,
	})
}