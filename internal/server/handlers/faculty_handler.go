package handlers

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
)

type FacultyHandler struct {
	facultyService *service.FacultyService
}

func NewFacultyHandler(facultyService *service.FacultyService) *FacultyHandler {
	return &FacultyHandler{
		facultyService: facultyService,
	}
}

func (h *FacultyHandler) GetAllFaculties(c *gin.Context) {
	faculties, err := h.facultyService.GetAllFaculties()
	if err != nil {
		// Handle error
		c.JSON(404, gin.H{
			"message": "No faculties found",
		})
	}
	// Return faculties
	c.JSON(200, gin.H{
		"message": faculties,
	})
}

func (h *FacultyHandler) CreateFaculty(context *gin.Context) {
	faculty := model.Faculty{}

	if err := context.ShouldBindJSON(&faculty); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.facultyService.CreateFaculty(context, &faculty); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return success message
	context.JSON(200, gin.H{
		"message": "Faculty created successfully",
	})
}
