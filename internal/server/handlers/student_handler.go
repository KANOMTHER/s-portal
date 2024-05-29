package handlers

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
)

type StudentHandler struct {
	studentService *service.StudentService
	authService    *service.AuthService
}

func NewStudentHandler(studentService *service.StudentService, authService *service.AuthService) *StudentHandler {
	return &StudentHandler{
		studentService: studentService,
		authService:    authService,
	}
}

// @Summary		CreateStudent
// @Description	create a new student
// @Tags			Student
// @Accept			json
// @Produce		json
// @Param			CreateStudentFields	body		model.CreateStudentFields	true	"CreateStudentFields object"
// @Success		200					{object}	string						"Student created successfully"
// @Failure		400					{object}	string						"some error message here (from err.Error())"
// @Router			/student [POST]
func (h *StudentHandler) CreateStudent(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	student := model.CreateStudentFields{}

	if err := context.ShouldBindJSON(&student); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	status, err := h.studentService.CreateStudent(&student)
	if err != nil {
		// Handle error
		context.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return success message
	context.JSON(status, gin.H{
		"message": "Student created successfully",
	})
}

// @Summary		GetDistinctYears
// @Description	get a distinct year of student
// @Tags			Student
// @Accept			json
// @Produce		json
// @Success		200	{array}		uint	"Array of distinct year in dedscending order"
// @Failure		404	{object}	string	"some error message here (from err.Error())"
// @Router			/student/year [GET]
func (h *StudentHandler) GetDistinctYears(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	distinct_year, err := h.studentService.GetDistinctYears()
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": distinct_year,
	})
}

// @Summary		GetStudentsByYear
// @Description	get a student by year
// @Tags			Student
// @Accept			json
// @Produce		json
// @Param			year	path		string	true	"64, 65, 66, 67, ..."
// @Success		200		{array}		uint	"Array of student's year [64, 65, 66, 67, ...]"
// @Failure		404		{object}	string	"some error message here (from err.Error())"
// @Router			/student/getByYear [GET]
func (h *StudentHandler) GetStudentsIDByYear(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	year := context.Param("year")
	distinct_year, err := h.studentService.GetStudentsIDByYear(year)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": distinct_year,
	})
}

// @Summary		GetStudentByID
// @Description	get a student by id
// @Tags			Student
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"student id"
// @Success		200	{object}	model.Student
// @Failure		404	{object}	string
// @Router			/student/{id} [GET]
func (h *StudentHandler) GetStudentByID(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	id := context.Param("id")
	student, err := h.studentService.GetStudentByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": student,
	})
}

// @Summary		UpdateStudentFields
// @Description	update a student by id depending on the user role
// @Tags			Student
// @Accept			json
// @Produce		json
// @Param			id				path		string						true	"student id"
// @Param			updatedField	body		model.UpdateStudentFields	true	"admin can change FName, LName, Graduated, Email, Phone; student can change FName, LName, Email, Phone"
// @Success		200				{object}	string						"Student updated successfully"
// @Failure		400				{object}	string						"some error message here (from err.Error())"
// @Router			/student/update/{id} [PUT]
func (h *StudentHandler) UpdateStudentByID(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	id := context.Param("id")
	status, err := h.studentService.UpdateStudentByID(context, id, h.authService)
	if err != nil {
		// Handle error
		context.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	// // Return success message
	context.JSON(status, gin.H{
		"message": "Student updated successfully",
	})
}

// @Summary		IsTA
// @Description	check if a student is a TA
// @Tags			Student
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"student id"
// @Success		200	{object}	uint	"TA id or null value if not TA"
// @Failure		404	{object}	string	"some error message here (from err.Error())"
// @Router			/student/is-ta/{id} [GET]
func (h *StudentHandler) IsTA(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	id := context.Param("id")
	ta_id, err := h.studentService.IsTA(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": ta_id,
	})
}
