package handlers

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
)

type InstructorHandler struct {
	instructorService *service.InstructorService
}

func NewInstructorHandler(instructorService *service.InstructorService) *InstructorHandler {
	return &InstructorHandler{
		instructorService: instructorService,
	}
}

func (h *InstructorHandler) GetAllInstructors(context *gin.Context) {
	instructors, err := h.instructorService.GetAllInstructors()
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": "No instructors found",
		})
	}
	// Return instructors
	context.JSON(200, gin.H{
		"message": instructors,
	})
}

func (h *InstructorHandler) CreateInstructor(context *gin.Context) {
	instructor := model.Instructor{}

	if err := context.ShouldBindJSON(&instructor); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.instructorService.CreateInstructor(&instructor); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return success message
	context.JSON(200, gin.H{
		"message": "Instructor created successfully",
	})
}

func (h *InstructorHandler) GetInstructorByID(context *gin.Context) {
	id := context.Param("id")
	instructor, err := h.instructorService.GetInstructorByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return instructors
	context.JSON(200, gin.H{
		"message": instructor,
	})
}

func (h *InstructorHandler) UpdateInstructorByID(context *gin.Context) {
	id := context.Param("id")
	err := h.instructorService.UpdateInstructorByID(context, id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Instructor updated successfully",
	})
}

func (h *InstructorHandler) DeleteInstructorByID(context *gin.Context) {
	id := context.Param("id")
	err := h.instructorService.DeleteInstructorByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Instructor deleted successfully",
	})
}
