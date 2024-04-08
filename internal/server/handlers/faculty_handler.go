package handlers

import (
	"github.com/gin-gonic/gin"

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
