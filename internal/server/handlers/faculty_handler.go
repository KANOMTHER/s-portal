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

func (h *FacultyHandler) GetFacultyByID(context *gin.Context) {
	id := context.Param("id")
	faculty, err := h.facultyService.GetFacultyByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": "Faculty not found",
		})
		return
	}
	// Return faculty
	context.JSON(200, gin.H{
		"message": faculty,
	})
}

func (h *FacultyHandler) UpdateFacultyByID(context *gin.Context) {
	id := context.Param("id")
	err := h.facultyService.UpdateFacultyByID(context, id)
	if err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Faculty updated successfully",
	})
}
