package handlers

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
)

type ClassHandler struct {
	classService *service.ClassService
}

func NewClassHandler(classService *service.ClassService) *ClassHandler {
	return &ClassHandler{
		classService: classService,
	}
}

func (h *ClassHandler) CreateClass(context *gin.Context) {
	class := model.Class{}

	if err := context.ShouldBindJSON(&class); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.classService.CreateClass(&class); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return success message
	context.JSON(200, gin.H{
		"message": "Class created successfully",
	})
}

func (h *ClassHandler) GetClassByID(context *gin.Context) {
	id := context.Param("id")
	class, err := h.classService.GetClassByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return classs
	context.JSON(200, gin.H{
		"message": class,
	})
}

func (h *ClassHandler) GetClassByCourseID(context *gin.Context) {
	id := context.Param("id")
	class, err := h.classService.GetClassByCourseID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return classs
	context.JSON(200, gin.H{
		"message": class,
	})
}
