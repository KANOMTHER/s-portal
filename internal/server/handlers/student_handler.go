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
